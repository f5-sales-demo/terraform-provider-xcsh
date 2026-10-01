---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b60dfb611cdf8a8a9479efc15b302353a74aec54e078c6d848b4622ed4feb864"></a>

## Property reference — Property reference / 56b0abd5bcb9 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- Property reference

<a id="canonical-43119c7fcb2253e518c41a17fcc5ff86e70d3aadb77cd2cfa620000286261a05"></a>

## Direct properties — Property reference / 56b0abd5bcb9 / 3

- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-c96b791db73eb7ed3c59393c5beb519e56f6b98af5a2999dbfbbef69414cbabb): complete subsection reference.

<a id="canonical-553c0bce8453c589c8260fdf1a8ab4dafabaa78c6be5c598e1115d7e3fa904bd"></a>

<a id="canonical-0a44761febe0292bfcc1fc2c8265cf0f6b1e8767ba9cf64f99b1f8633d5157b2"></a>

## annotations property — Property reference / 56b0abd5bcb9 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-29d3d0446d62faa0d39154381930b796864b6b683593f6fa1f025b812be1937e): complete subsection reference.

<a id="canonical-407033a5b28212a9fa2b2d0daa0571bdb68160e34b7bdf0f1c33f656799e5138"></a>

<a id="canonical-2c24b16f4935ac4c99e8ce9af9180a583c2b3eddf10bc595ba9f8d1642a32ac6"></a>

## description property — Property reference / 56b0abd5bcb9 / 5

Type: `"string"`. Computed.

Description of the ApplicationProfiles.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-d2e807f00e2b6d37f74a491dd4da1c2832ccb22ebcc88894dea6b92e5b4e3125"></a>

<a id="canonical-8842747f7cd1717ed962c487862b8015160eea1bdab326e7531f838703518917"></a>

## id property — Property reference / 56b0abd5bcb9 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](data-sources--application_profiles--reference--group-001.md#canonical-b03852e8035d88ad972f06db07b19a393462dd0e06b97aff9ac57a85e610479b): complete subsection reference.

<a id="canonical-e3d6a779dca4cf85451988d9763e5920678e5045b2fc6f9bd70b71cfb2d492a9"></a>

<a id="canonical-f820e38cb6b0f7288f5fb0a65eb0659776775ec15106f527dbde0f6eea35b43b"></a>

## labels property — Property reference / 56b0abd5bcb9 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ea96340a8f3a482f051c465cddc4fb71e54c0d92497a1913415367f4499ca0f2"></a>

<a id="canonical-ab2cc2e75a62bf962d75ea055295a0daea0cd00abe699387a6c85e0f525add98"></a>

## name property — Property reference / 56b0abd5bcb9 / 8

Type: `"string"`. Required.

Name of the ApplicationProfiles.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-506cbfd4b79a664779b05eb167bd9320a9b64f5b876f6c7f71b10e22f656f60d"></a>

<a id="canonical-ed76b0d0bc67857cc2b9d94c30126d778bce0d3dac3b979e513eb3feeec1dbed"></a>

## namespace property — Property reference / 56b0abd5bcb9 / 9

Type: `"string"`. Required.

Namespace where the ApplicationProfiles exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc): complete subsection reference.

<a id="canonical-7f3aa656bdc6030d844ee09b3247ee010850b2779ff86fde339f3d86f04748bc"></a>

## All schema paths — Property reference / 56b0abd5bcb9 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_tcp_profile` | [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-d3767388aea02ec0ca5025809e8b9ea513436b560577f5b21d781ce28158f42b) |
| `advanced_tcp_profile.disable_tcp_advanced_profile` | [advanced_tcp_profile.disable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-0d26b2193786af7aa79e0db3600a2e41886b063eee854d74aa5addef66f9383b) |
| `advanced_tcp_profile.enable_tcp_advanced_profile` | [advanced_tcp_profile.enable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-9e6ebe48e8f5f7a29e5a34daccaba29516ab8384272cdf414e783ef391dba3fd) |
| `annotations` | [annotations](data-sources--application_profiles--reference--group-001.md#canonical-553c0bce8453c589c8260fdf1a8ab4dafabaa78c6be5c598e1115d7e3fa904bd) |
| `ddos_profile` | [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-b8c30b3a65e6d9609c5c15e10affae0439226c346239ecc25fe5608ef8470fc8) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-ba0a9fa3a16fff5595f46101036ae725e701da0f88906e3bfc6cc6b3f5fca749) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-61c54997e65f4ea0911b9bf8f4da246106d1643e36664a7e10d7ee912196e5e6) |
| `description` | [description](data-sources--application_profiles--reference--group-001.md#canonical-407033a5b28212a9fa2b2d0daa0571bdb68160e34b7bdf0f1c33f656799e5138) |
| `id` | [id](data-sources--application_profiles--reference--group-001.md#canonical-d2e807f00e2b6d37f74a491dd4da1c2832ccb22ebcc88894dea6b92e5b4e3125) |
| `irules` | [irules](data-sources--application_profiles--reference--group-001.md#canonical-9f505918e95bbf193101e06dc82e460ae7385da1d3b0e487261980ebfc4c2fa0) |
| `irules.kind` | [irules.kind](data-sources--application_profiles--reference--group-001.md#canonical-ad0d9d62279d6fb30c0456cb22ef8c40aafb7cab14347c374325bfb8d41723c7) |
| `irules.name` | [irules.name](data-sources--application_profiles--reference--group-001.md#canonical-ce7fd1101c941540c77558402346710e11496dbac0f0189034db95a46db22d8a) |
| `irules.namespace` | [irules.namespace](data-sources--application_profiles--reference--group-001.md#canonical-c537b88d8371e4da4a11e7917efcdb6f972e38b8bc3603453241ff6ffe283574) |
| `irules.tenant` | [irules.tenant](data-sources--application_profiles--reference--group-001.md#canonical-175f2bea23fd12acd42a1af0d58283c7ee0a7379d6a7cc91a9398b7bda102f44) |
| `irules.uid` | [irules.uid](data-sources--application_profiles--reference--group-001.md#canonical-a3ef076b67a7821b4a918797dc2fd8ae2de2e9ce3b4d05d01b42f973020f4853) |
| `labels` | [labels](data-sources--application_profiles--reference--group-001.md#canonical-e3d6a779dca4cf85451988d9763e5920678e5045b2fc6f9bd70b71cfb2d492a9) |
| `name` | [name](data-sources--application_profiles--reference--group-001.md#canonical-ea96340a8f3a482f051c465cddc4fb71e54c0d92497a1913415367f4499ca0f2) |
| `namespace` | [namespace](data-sources--application_profiles--reference--group-001.md#canonical-506cbfd4b79a664779b05eb167bd9320a9b64f5b876f6c7f71b10e22f656f60d) |
| `virtual_server` | [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-6d8b283c9385612b65995a6a3cdd81b8bb0f78278a1f7e926c3382872d8b7260) |
| `virtual_server.access_profile` | [virtual_server.access_profile](data-sources--application_profiles--reference--group-001.md#canonical-f18c4d94352b1ed57cb7e55085e3d01eaa7094df7dcead0f0cee184453bb76ad) |
| `virtual_server.access_profile.kind` | [virtual_server.access_profile.kind](data-sources--application_profiles--reference--group-001.md#canonical-addd341b46ac881e8c5b97dd28d7bc7b288efc54b5481b822f0aa7bbfaaa4654) |
| `virtual_server.access_profile.name` | [virtual_server.access_profile.name](data-sources--application_profiles--reference--group-001.md#canonical-41302f6760a20249e1d4869da6deb91fdadc1ae85873a1b2bf35321385d14674) |
| `virtual_server.access_profile.namespace` | [virtual_server.access_profile.namespace](data-sources--application_profiles--reference--group-001.md#canonical-62defd0bb3dbb951810770ff982d02885a94a7f164663eadb76304093a7b33d5) |
| `virtual_server.access_profile.tenant` | [virtual_server.access_profile.tenant](data-sources--application_profiles--reference--group-001.md#canonical-e797019088039352cf9fb282a43ee6b3bf83e601b583333f33fa477989ad25c3) |
| `virtual_server.access_profile.uid` | [virtual_server.access_profile.uid](data-sources--application_profiles--reference--group-001.md#canonical-afd917f455def6751b0b3d22fef529d27d30deea328c2867b4e4729834ec66de) |
| `virtual_server.address_translation` | [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-d5cb7b1fa01499170bf47bf0de2596b184f2a9ac70b0d5ab247153464679d3e0) |
| `virtual_server.address_translation.address_translation_disable` | [virtual_server.address_translation.address_translation_disable](data-sources--application_profiles--reference--group-001.md#canonical-3f4060e9e0e5b960f3b4300906422a8755cfcddcfdce9b0897e0c242ee89b670) |
| `virtual_server.address_translation.address_translation_enable` | [virtual_server.address_translation.address_translation_enable](data-sources--application_profiles--reference--group-001.md#canonical-b40c437c53279a11948161c044f26ba354ac407b7b8fcb23b991e3f8fc692711) |
| `virtual_server.auto_last_hop` | [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-937a6558a263dc0f40a43b93040407bcf1d1bea17e00581999961e95406c6dbc) |
| `virtual_server.auto_last_hop.auto_last_hop_default` | [virtual_server.auto_last_hop.auto_last_hop_default](data-sources--application_profiles--reference--group-001.md#canonical-0f3d26554379edf0f708a65ddd7dac1fc63f4e38b7da543656ca0be810045171) |
| `virtual_server.auto_last_hop.auto_last_hop_disable` | [virtual_server.auto_last_hop.auto_last_hop_disable](data-sources--application_profiles--reference--group-001.md#canonical-f46bae3d4532eb7ff3e67ae4bb710b04e61a785e17d8fb39fa881f9db5432173) |
| `virtual_server.auto_last_hop.auto_last_hop_enable` | [virtual_server.auto_last_hop.auto_last_hop_enable](data-sources--application_profiles--reference--group-002.md#canonical-82206f3b2f1975ec40ed44e55f1960fe96510f1ec9d22b955464cbd70026bafb) |
| `virtual_server.clone_pool_client` | [virtual_server.clone_pool_client](data-sources--application_profiles--reference--group-002.md#canonical-5f21ebe5b4511a050617e0b9479dedcf91e28d53d51ba4590cb00fdeee099a77) |
| `virtual_server.clone_pool_client.kind` | [virtual_server.clone_pool_client.kind](data-sources--application_profiles--reference--group-002.md#canonical-e73990620c630b475a7a61bd2756a3bc2d46a5a009c23dfd8944a3a0c3bd5fc3) |
| `virtual_server.clone_pool_client.name` | [virtual_server.clone_pool_client.name](data-sources--application_profiles--reference--group-002.md#canonical-79cf1ed6eef4c2cbcaff54c4e7ac7dfa31d1669f27278fa779163591c81ec621) |
| `virtual_server.clone_pool_client.namespace` | [virtual_server.clone_pool_client.namespace](data-sources--application_profiles--reference--group-002.md#canonical-a5b6aad587df6b3fac47bb6af430e35a75fe798bc1a35251a10c4f2cfd2f7e6e) |
| `virtual_server.clone_pool_client.tenant` | [virtual_server.clone_pool_client.tenant](data-sources--application_profiles--reference--group-002.md#canonical-84fa3e14e52dbe6ba8b066f9e832492a582421fbc7c017076a2a7a89521f3e7e) |
| `virtual_server.clone_pool_client.uid` | [virtual_server.clone_pool_client.uid](data-sources--application_profiles--reference--group-002.md#canonical-6f564e025e781779fd26e0b7e730307710beb774113d27f266dd6c04f9eb5d90) |
| `virtual_server.clone_pool_server` | [virtual_server.clone_pool_server](data-sources--application_profiles--reference--group-002.md#canonical-29dd50e32aa5ed95979b7dd9e3ac8a4618b69ed68e6d42aabb77c57ed6acb868) |
| `virtual_server.clone_pool_server.kind` | [virtual_server.clone_pool_server.kind](data-sources--application_profiles--reference--group-002.md#canonical-6f59548b35299463534c643bc214049db390c75f2e5c76542479fea6ea0e05bd) |
| `virtual_server.clone_pool_server.name` | [virtual_server.clone_pool_server.name](data-sources--application_profiles--reference--group-002.md#canonical-b335996841b27c0c6c0fed4e8843eb67820ad20564c8a098dd7783410d063a2d) |
| `virtual_server.clone_pool_server.namespace` | [virtual_server.clone_pool_server.namespace](data-sources--application_profiles--reference--group-002.md#canonical-91e361ed2bdcf710b383c7b64e145e31e1fed4713369bd7904216b3ef83d7311) |
| `virtual_server.clone_pool_server.tenant` | [virtual_server.clone_pool_server.tenant](data-sources--application_profiles--reference--group-002.md#canonical-0efca091478e84c6812815566b8b779a2cb3055caab9d4dcc2d068671713a711) |
| `virtual_server.clone_pool_server.uid` | [virtual_server.clone_pool_server.uid](data-sources--application_profiles--reference--group-002.md#canonical-67be68133578d45274404e49ac1e87ac629dfdc3ea234e36ffb60878e38b9ce6) |
| `virtual_server.connection_limit` | [virtual_server.connection_limit](data-sources--application_profiles--reference--group-001.md#canonical-c3d44cfa848788c88ac81d3ccac710087e8bc390d8b3e8e1d113a3e340f5380e) |
| `virtual_server.connection_rate_limit` | [virtual_server.connection_rate_limit](data-sources--application_profiles--reference--group-001.md#canonical-c6b2ef93baf1a91d75e3f5835daea2e407683a962196a7384304944dd3eaebcd) |
| `virtual_server.connection_rate_limit_mode` | [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-ddc4ed26919c16086ceb2bcaefafcb34d867c9cc514e0c9716fc6cec179e319c) |
| `virtual_server.connection_rate_limit_mode.per_destination_address` | [virtual_server.connection_rate_limit_mode.per_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-ee9b3d2df982dc70d72c52f5f7503e8fcfee153548801069800663586acbc4c2) |
| `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask](data-sources--application_profiles--reference--group-002.md#canonical-4119fea0293eb65fbd643d95bd1167fea7b5b3fcb3ebfc2ae29c5c9bbee723eb) |
| `virtual_server.connection_rate_limit_mode.per_source_address` | [virtual_server.connection_rate_limit_mode.per_source_address](data-sources--application_profiles--reference--group-002.md#canonical-2aa3ed29cd305a7b4514b960e06e7ad9ee11027971ccaf319f4a3e5747c81b58) |
| `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_address.source_mask](data-sources--application_profiles--reference--group-002.md#canonical-edd5dd522d8a5793b3182557bc2acc450c04dc2d74a33901cac8299cd9d3c2e0) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-4d87d6d980d30242f54c620326ff6a882c728222afef67d62f9ab0ef52513eeb) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask](data-sources--application_profiles--reference--group-002.md#canonical-86d1bb5653dd131f715aa0998ff0b1544af48f06ed07075fce7fc58e6608696c) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask](data-sources--application_profiles--reference--group-002.md#canonical-8e105434fcd467f29a3406f839675b58d960151b134be74c580dde4ef0a10ac8) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server` | [virtual_server.connection_rate_limit_mode.per_virtual_server](data-sources--application_profiles--reference--group-002.md#canonical-7df149253f9ad02c14e4bf0e12b3efb6e1bb4a42a4ba591896be99d318b25272) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-e83333c5c6c718611130819d4816b640e5fb4ee141c0e12828bd92ed6a744cbe) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask](data-sources--application_profiles--reference--group-002.md#canonical-84072962501041f56f4ef0f34f616b99537a30f6f6449164f1c6dfb798f6756e) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](data-sources--application_profiles--reference--group-002.md#canonical-a0d0a7b40ae0fce1696739b70b3b56491e9d72f723e8965b3465b0eea355cc3b) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask](data-sources--application_profiles--reference--group-002.md#canonical-24c66a1beda73999202cc1ed8465635f646b848e29dae19f8e403291fa8fea0a) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-91572c8df33fe9baa93d7d3ba6f0098a0db67d64471b83069b17772c992a7790) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask](data-sources--application_profiles--reference--group-002.md#canonical-fd7594b3979db6b09048ca2410b77c323fdd2ae7e3e339fe6bdcaddd0f7dac3b) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask](data-sources--application_profiles--reference--group-002.md#canonical-a1c6aa2f3b90685145c49e3d62d0118ee1cebf08d6ee0bc0003496f1e3b278d9) |
| `virtual_server.default_persistence_profile` | [virtual_server.default_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-1e7938358f045952100fcca22fb3c8c5af8d26119330c252db38be6e835badc5) |
| `virtual_server.default_persistence_profile.kind` | [virtual_server.default_persistence_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-433ea78566aeca4d9a219d8be0a14e3ba23d8390a981c6585570e1f37be5b94a) |
| `virtual_server.default_persistence_profile.name` | [virtual_server.default_persistence_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-3b2bf8ce26449710b2874c472ad621e91d86858c80b7a448d6928396045141cf) |
| `virtual_server.default_persistence_profile.namespace` | [virtual_server.default_persistence_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-99388833946b4fc32c505f79c91de91f980d24a5d56112b315e919905b86061e) |
| `virtual_server.default_persistence_profile.tenant` | [virtual_server.default_persistence_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-fc7d6097b501e3cf9d8a8a32f8f7d94b42115e38bb62ab38aa8f3825d9bdf0aa) |
| `virtual_server.default_persistence_profile.uid` | [virtual_server.default_persistence_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-8158350217910ac826077e3c21b7e4dfda3f5414c6d22e670c03d6e0e3560dd3) |
| `virtual_server.default_pool` | [virtual_server.default_pool](data-sources--application_profiles--reference--group-002.md#canonical-0e0dc9da9e9097c5e4ed16e17dab2b180e879646d4c59aec036a91444e298f11) |
| `virtual_server.default_pool.kind` | [virtual_server.default_pool.kind](data-sources--application_profiles--reference--group-002.md#canonical-81ea534fc69e3c233e263ab48a72f3372532fdd1b5786d1d4df426b9a5420c08) |
| `virtual_server.default_pool.name` | [virtual_server.default_pool.name](data-sources--application_profiles--reference--group-002.md#canonical-ce9f8dbb6af102ef239352868dc3273997050e554df2aac90a958b40c4d65c5c) |
| `virtual_server.default_pool.namespace` | [virtual_server.default_pool.namespace](data-sources--application_profiles--reference--group-002.md#canonical-a75c3cedc697cf7abab9925b31c70dcfa71680e2a26c2dcb54a6beb8a15f195a) |
| `virtual_server.default_pool.tenant` | [virtual_server.default_pool.tenant](data-sources--application_profiles--reference--group-002.md#canonical-91b6f0c64af0da8fffb870d3c081db8ae5677377e48a02494b5647247c1a297c) |
| `virtual_server.default_pool.uid` | [virtual_server.default_pool.uid](data-sources--application_profiles--reference--group-002.md#canonical-5254c7e98f59b24713b906c8912f3c18b37f03895e16fc0258230db02332afdf) |
| `virtual_server.fallback_persistence_profile` | [virtual_server.fallback_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-cda90e7c0c5a0f593ed8e46c75cb5c5efdb09d15e5b64f090866401902b39b0d) |
| `virtual_server.fallback_persistence_profile.kind` | [virtual_server.fallback_persistence_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-0cc90ce266fecdd5334b561ee9ca37ae03b208c3d7e98bd434abef684736fef3) |
| `virtual_server.fallback_persistence_profile.name` | [virtual_server.fallback_persistence_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-8557cc993fa464c9b0a66b2a42d60a9b98ca1e4a0b5e1d2eea43618c5838e349) |
| `virtual_server.fallback_persistence_profile.namespace` | [virtual_server.fallback_persistence_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-33cb47cece85343d84201118fd3393b306aec103f92dad1a438a07bdc171f954) |
| `virtual_server.fallback_persistence_profile.tenant` | [virtual_server.fallback_persistence_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-e13c1afcd9264424301c7b8aad7b0b2b7623b2ab7c6e472338a6cb2f8a8cdcd4) |
| `virtual_server.fallback_persistence_profile.uid` | [virtual_server.fallback_persistence_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-6ff9953699b348d69aa3a02c03109a3f17bea350534e9e9b0948231c2c762856) |
| `virtual_server.fix_profile` | [virtual_server.fix_profile](data-sources--application_profiles--reference--group-002.md#canonical-bfc05c64e7b23d9747294460ce2e0e992fc21840ad29934e049e9f688a7d631c) |
| `virtual_server.fix_profile.kind` | [virtual_server.fix_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-5aeff7824d7cb8b2081f461cb2b3eb09d7d3423147fb0377392bbd0d67e55ee8) |
| `virtual_server.fix_profile.name` | [virtual_server.fix_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-629d214d39851e606c254872a7485648b1609df8fb74fb7bfaed1ea5425fceeb) |
| `virtual_server.fix_profile.namespace` | [virtual_server.fix_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-4a9de7edee40ef15dad7c2e132c685fcd6856047c370094dd49da576ee8a3c21) |
| `virtual_server.fix_profile.tenant` | [virtual_server.fix_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-a485e75275af0ae1b06f0fca25fe574be8c9611b2d4f8b1f7ac7eb6c7e3b8b1d) |
| `virtual_server.fix_profile.uid` | [virtual_server.fix_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-31452fb8d1876d15f68b4c9c4c9d5b82062803be61f4c1f197d85eeeda5ff6bd) |
| `virtual_server.http` | [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-3e91f022ac2cc4b9431b39f5f986bd64a35359eb0043e2e380f955e3249e0712) |
| `virtual_server.http.client_ssl_profile` | [virtual_server.http.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-1b99ccfab4dcf31c2422084b09e6dc049efaac7c519638faf63e13a45c25d15e) |
| `virtual_server.http.client_ssl_profile.kind` | [virtual_server.http.client_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-3c9b156b7de194c526a45972f3f772c860f15f7312cd83c1e7d18dbdfee4fbee) |
| `virtual_server.http.client_ssl_profile.name` | [virtual_server.http.client_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-711c38fbd9fd7b68dadeab50257e8e709c2f2c04cf99d18451f7c0c3f88deda5) |
| `virtual_server.http.client_ssl_profile.namespace` | [virtual_server.http.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-e2415dcd4f39a391695b994be809702d4938fe4e544a698dc4f68fcc2a77ff3a) |
| `virtual_server.http.client_ssl_profile.tenant` | [virtual_server.http.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-056b00b50a5c63c710215266e3dada4b7d9394ee3ea2b30c279032d9160b7a6a) |
| `virtual_server.http.client_ssl_profile.uid` | [virtual_server.http.client_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-c364ae7f9dd008ff803a51590348c2c3656d0559676701bf2c7b1fa0473f1f41) |
| `virtual_server.http.http2_client_profile` | [virtual_server.http.http2_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-06c6df1e540c4dd546442e888f4da2c896616ef789c8e966db6ef71e526fa616) |
| `virtual_server.http.http2_client_profile.kind` | [virtual_server.http.http2_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-59e165efb900b4edee109bff124bde6e4da17e3620011aa3b609514e7dfe716a) |
| `virtual_server.http.http2_client_profile.name` | [virtual_server.http.http2_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-3f5e855d20df012a4f858e4c3f12a1055f7a2f9c6b1f019578bc5b3e7ed52c52) |
| `virtual_server.http.http2_client_profile.namespace` | [virtual_server.http.http2_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-c56bb928b0b39278ccd126936a34f226c0531323dbd6225056228cb71ccd10dc) |
| `virtual_server.http.http2_client_profile.tenant` | [virtual_server.http.http2_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-0b7a3703d03d128cafa33019ff4989cc3e244d3d38c7ba5a4b6c2b7f52c629d2) |
| `virtual_server.http.http2_client_profile.uid` | [virtual_server.http.http2_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-f489a3f43f4cece41581224b05b40a327a72fa743dc76252da452f893b6c1b89) |
| `virtual_server.http.http2_server_profile` | [virtual_server.http.http2_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-d0a400e20566b3e910996d9cbe1f9e9d857be08e131bf3bd0fcfd634badcf0a5) |
| `virtual_server.http.http2_server_profile.kind` | [virtual_server.http.http2_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-c38726d21def8650362c2f3527826c47e0ece7f33877dd0bc5b997dc3fdab994) |
| `virtual_server.http.http2_server_profile.name` | [virtual_server.http.http2_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-3b03ff90266efa29472fe8bc7279e2bdf47cc7388353bcecd506743f6b80891f) |
| `virtual_server.http.http2_server_profile.namespace` | [virtual_server.http.http2_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-8b7c4b7eba5c6a33eb4efb02a8d9f429e454606743d635caee9e74b55668cb6e) |
| `virtual_server.http.http2_server_profile.tenant` | [virtual_server.http.http2_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-ac8455c8336f13ff3fb9a1c97679c4257fea278bec09a0f23dc89134ff31b1b6) |
| `virtual_server.http.http2_server_profile.uid` | [virtual_server.http.http2_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-29141dd2a1dfe3b767d738f453a9523e5d0226c49d180f898aa32a4394aef805) |
| `virtual_server.http.http_client_profile` | [virtual_server.http.http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-1a9c78f5056622fe56eb7018efe01d68405639a8d3479713e9551510945d5449) |
| `virtual_server.http.http_client_profile.kind` | [virtual_server.http.http_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-dcac4f14d5d7401edc7d714e03d40e34a11b49fd078f7b37e8103e91928b6cd8) |
| `virtual_server.http.http_client_profile.name` | [virtual_server.http.http_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-8c10e3b32b4c6cd179b0ea7e3dfa98a81cf7ba6d2e9b5cfd898fb6e311ef8f07) |
| `virtual_server.http.http_client_profile.namespace` | [virtual_server.http.http_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-f2fc114ebd2b26c042bb336937faec1a1d9126e951c5db8766061928ebce2272) |
| `virtual_server.http.http_client_profile.tenant` | [virtual_server.http.http_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-9374dac7e768ff91a3e7807769171c05f16c522a624d638166be91b38b892193) |
| `virtual_server.http.http_client_profile.uid` | [virtual_server.http.http_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0a32a30e69f7dc3a77363d45ab11ba22fc0be5bbcd730ebd3e5f162119f0c0aa) |
| `virtual_server.http.http_server_profile` | [virtual_server.http.http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-aa02b94dab5d90b52a777578cff8c404b6dbb4c4b331f51eb01d2d7cb6d078c0) |
| `virtual_server.http.http_server_profile.kind` | [virtual_server.http.http_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-28218f8978ae4c8a86401c2624d3a86ddd6b3e65cdb206bd37b0bebf8385591f) |
| `virtual_server.http.http_server_profile.name` | [virtual_server.http.http_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-3c2a5fd59f8e1ca48b7d3b930358be4f621061d918c342b30a461f9cb507c99b) |
| `virtual_server.http.http_server_profile.namespace` | [virtual_server.http.http_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-295f4d50b395e1549e10dfeee6698fbcdef1a36ca7bd10a10943087eb03a428a) |
| `virtual_server.http.http_server_profile.tenant` | [virtual_server.http.http_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-5bb3dea00e808d20f0d57e25fff056710e7706cbc655d1094dcc318fc2377230) |
| `virtual_server.http.http_server_profile.uid` | [virtual_server.http.http_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-95c824cf1cbf084fb538662682faffdae12fd79a6dc853c809bb2c4d7b5c8b7d) |
| `virtual_server.http.ocsp_profile` | [virtual_server.http.ocsp_profile](data-sources--application_profiles--reference--group-002.md#canonical-3685b833adb363b1381de77d94856efc3319543397d6ead25e84e9b52b5c2d9a) |
| `virtual_server.http.ocsp_profile.kind` | [virtual_server.http.ocsp_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-6ce506af7c30ef5245e5ed626c87f056c87490c2b739245a8f215c6ff4a81fc2) |
| `virtual_server.http.ocsp_profile.name` | [virtual_server.http.ocsp_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-b060b151431754ea49f5c55314dbfbd9a452ffacfb8eb267c5eec6fe57643a17) |
| `virtual_server.http.ocsp_profile.namespace` | [virtual_server.http.ocsp_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-47b08256932f3ee177d5e49867ca4ef96d1e2d7ca9140e7a2f39191695416afe) |
| `virtual_server.http.ocsp_profile.tenant` | [virtual_server.http.ocsp_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-1028207a49e7e5abf5a7d10ac8df3dc80211692249018ac3e699e989b3d0f044) |
| `virtual_server.http.ocsp_profile.uid` | [virtual_server.http.ocsp_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-064e8993dfe95c8fd5511e4f9fa7c1820921a3ab953eece18f2b046dfaa2953a) |
| `virtual_server.http.server_ssl_profile` | [virtual_server.http.server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-4224d0dbf16bde64a9f8a59817118481976c1baa8e977ff924313965b55cb1d2) |
| `virtual_server.http.server_ssl_profile.kind` | [virtual_server.http.server_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-f519bad014b7b52c25d60a78e94dffc1263bbbc49c24fd8acc308671c27993a4) |
| `virtual_server.http.server_ssl_profile.name` | [virtual_server.http.server_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-4c27d911a00ba625bb1f48986327bd1c64f50bbcd834f783bd29c0e21a3051e3) |
| `virtual_server.http.server_ssl_profile.namespace` | [virtual_server.http.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-d76262d1b0fbe848118c8019a1d2b5ad0186a232703c29250ae79ae0da4d8b40) |
| `virtual_server.http.server_ssl_profile.tenant` | [virtual_server.http.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-cdb715882dfb96d730a73fe0711ae6220913e1e1d90c0c78d971ddda21e08ceb) |
| `virtual_server.http.server_ssl_profile.uid` | [virtual_server.http.server_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-5a196d42a153999a79504d8b24f348d948b6755d0acd5c1dad6db3e01c80a0d3) |
| `virtual_server.http.stream_profile` | [virtual_server.http.stream_profile](data-sources--application_profiles--reference--group-002.md#canonical-de9e1be889e52150a81b6d319541f21a17a83cb924b0094b266f36463cce7948) |
| `virtual_server.http.stream_profile.kind` | [virtual_server.http.stream_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-2c024fb5eaf1996d3740b872d9effb256c95aad4045b2a36b738aab2d7536885) |
| `virtual_server.http.stream_profile.name` | [virtual_server.http.stream_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-b7e93084dd7b97c9ba6a8c19bfc6cdf48fa5e2794a08c124a5a0cb348d422367) |
| `virtual_server.http.stream_profile.namespace` | [virtual_server.http.stream_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-167aac626058f3d2b9d4c84e58ebb55b005384154d48e001e99d605dc07e5e56) |
| `virtual_server.http.stream_profile.tenant` | [virtual_server.http.stream_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-25cc2c67df7661e38433d08e796c5c61840c7c1f63d669607ae8773975fc227f) |
| `virtual_server.http.stream_profile.uid` | [virtual_server.http.stream_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-17b7483175589887d8b915dbc7f9a87a1b770030a2dcc83bd63ada74a40b493d) |
| `virtual_server.http.tcp_client_profile` | [virtual_server.http.tcp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-c51b4e45bfd112c205832f158012f507f8410becb123841f1e3c181b17435e41) |
| `virtual_server.http.tcp_client_profile.kind` | [virtual_server.http.tcp_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-66d48c18ebd73937c21b69b14acc7dd7d0268116ea7326f48463a7b03b98be06) |
| `virtual_server.http.tcp_client_profile.name` | [virtual_server.http.tcp_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-fe7aa3cfb10df9f967570bf566b82b64941685c63b0f8026d5c289170a078ddc) |
| `virtual_server.http.tcp_client_profile.namespace` | [virtual_server.http.tcp_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-e8f5eebfecbe1ba7f3ea15b0299f23bf92256098b1f3c44f9dd468c36360c89e) |
| `virtual_server.http.tcp_client_profile.tenant` | [virtual_server.http.tcp_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-5fdfd73ebad5b844f8665936c8d9c947bb187a0dced74a87cae9deb1b483f772) |
| `virtual_server.http.tcp_client_profile.uid` | [virtual_server.http.tcp_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-4f920a894ad3f619bcb1a3f38d09edac59280074a7ee06db65e0b831d5149f2f) |
| `virtual_server.http.tcp_server_profile` | [virtual_server.http.tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-4df66e759912115ff2c98d3352738a3ac31842b64b38ea894abe97cabc3ff691) |
| `virtual_server.http.tcp_server_profile.kind` | [virtual_server.http.tcp_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-7bcfaf6b78e08a7cc9f2e1eabf0c6bb9ad017568837b2ce5ae9c58e054265ef1) |
| `virtual_server.http.tcp_server_profile.name` | [virtual_server.http.tcp_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-f0f224ea108def99ee8ac8ee9d0586ffcae3a965af818d6ac842877e591e7612) |
| `virtual_server.http.tcp_server_profile.namespace` | [virtual_server.http.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-38d3a103eb7ddcc05f995d7aa09f13f0f66fda7cfa3da772c61ed97f13686976) |
| `virtual_server.http.tcp_server_profile.tenant` | [virtual_server.http.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-dbca43842ede3bd74794df567f8bc3485eca859d80fb324708d234886ebc5797) |
| `virtual_server.http.tcp_server_profile.uid` | [virtual_server.http.tcp_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-14afc38dc1eecda70db78953b834f0a448c2c5792b93c25735fbbdfb898b59c6) |
| `virtual_server.http.websocket_client_profile` | [virtual_server.http.websocket_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-99a75be6986d1779943a76b5db2f7896887f33654e5bd138dd25fa01d7bce3d5) |
| `virtual_server.http.websocket_client_profile.kind` | [virtual_server.http.websocket_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-2aa93501420a94f1e282f2d374bb2ad7cbe15a76a659579d6bfa0d3f57934992) |
| `virtual_server.http.websocket_client_profile.name` | [virtual_server.http.websocket_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-751a1dbae2b1543e08985ea76f49a1edf3ce431d052f078e551f2aaa65e87bee) |
| `virtual_server.http.websocket_client_profile.namespace` | [virtual_server.http.websocket_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-e24476187f8804bfb7ec1ce1699e13bff0a0fb9e3596f8570d93294b9f4302c3) |
| `virtual_server.http.websocket_client_profile.tenant` | [virtual_server.http.websocket_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-b7c0714faf9100daae542ca41330b27e96951e29f0bfc16313e0b722046a1d38) |
| `virtual_server.http.websocket_client_profile.uid` | [virtual_server.http.websocket_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-41ab680b7ff1cb85c8c1feb78350640a67d0c7a23bcf1e44eb2d32a7b2610b35) |
| `virtual_server.http.websocket_server_profile` | [virtual_server.http.websocket_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-a1eb4f060c8c0a690908021d81dc0f9a664894d562ec56a50a8e663a8970b9f6) |
| `virtual_server.http.websocket_server_profile.kind` | [virtual_server.http.websocket_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-97cf291a7d43a9b8c2dd6f91153df5220779bc46a5b3582509b89ec8c720b2bb) |
| `virtual_server.http.websocket_server_profile.name` | [virtual_server.http.websocket_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-68f2a25d912a4387ae5ba30943dc5ee579feb41e2053f72bb1f075e0cb7eadce) |
| `virtual_server.http.websocket_server_profile.namespace` | [virtual_server.http.websocket_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-4717e1c31e04546992e0c711608598b7e26928abbb5d94b09920cb8a3de99874) |
| `virtual_server.http.websocket_server_profile.tenant` | [virtual_server.http.websocket_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-5a068417c7abf27dde51081c3f247e510be672746a7775996998f3b24509fc25) |
| `virtual_server.http.websocket_server_profile.uid` | [virtual_server.http.websocket_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-c7f6f90c55f70d2a9bb1200a7daee4fc6b5e7e8a4946a14a801864989e9f7297) |
| `virtual_server.http3` | [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-59c74f35851384576f40767c66a9253e2fd71cb1d116e74ea1357b5355bd03d5) |
| `virtual_server.http3.client_ssl_profile` | [virtual_server.http3.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-2da862f0c0a27f3961b846ec9ac0c9b0fbb93d1c88aa2d038e607b54538e496b) |
| `virtual_server.http3.client_ssl_profile.kind` | [virtual_server.http3.client_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-cb93c4023178aa786455638fbb10943269568649871b417a1e9bf321cbf74090) |
| `virtual_server.http3.client_ssl_profile.name` | [virtual_server.http3.client_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-d215931e8a5ebd7489d85a5b7948e1ee017e66d2ed6051656301f14ae37b1e95) |
| `virtual_server.http3.client_ssl_profile.namespace` | [virtual_server.http3.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-a04f0ce00e8c93d3682dd00c1771fb94ac3649d4ea11d22484a4dd07d4887f66) |
| `virtual_server.http3.client_ssl_profile.tenant` | [virtual_server.http3.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-c792d1e98e463d8b88faf7fe4dc73dad20af8882e640d42228beeb23527acbbf) |
| `virtual_server.http3.client_ssl_profile.uid` | [virtual_server.http3.client_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-2504c99b5fff5fd0e127140bc408f47955d69197007bbd3b6b99fbd7f7877d11) |
| `virtual_server.http3.http3_profile` | [virtual_server.http3.http3_profile](data-sources--application_profiles--reference--group-002.md#canonical-e85b342deb4612b3bfb19a47d01b564d9ba0f221961ff9f5a8f6a55e6d23a1be) |
| `virtual_server.http3.http3_profile.kind` | [virtual_server.http3.http3_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-a697d4db6a32b7ade68b1ef43f8833817c56cef1150f7787ba78f8236ce43693) |
| `virtual_server.http3.http3_profile.name` | [virtual_server.http3.http3_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-78f4c4e2ea0cec354bafe02699029e318b9f190dd6cbf233ea7092d70375a18a) |
| `virtual_server.http3.http3_profile.namespace` | [virtual_server.http3.http3_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-68374b757b28c671cd02be73291e051da9b891a7adc52b474329adb432800d87) |
| `virtual_server.http3.http3_profile.tenant` | [virtual_server.http3.http3_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-00d6ddfdc23b96a84db1a7fd6f561749cf5a08bc8ccee037f93d599dfae00421) |
| `virtual_server.http3.http3_profile.uid` | [virtual_server.http3.http3_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-48123a4b3a4706efb3a7f1aeb8bc539dd169e6ccdb967e97157245ce940b5679) |
| `virtual_server.http3.http_client_profile` | [virtual_server.http3.http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-15ccf63ee6aac4c25cabb0155afb24aa61ccdf950196d80611c37ed485e026b5) |
| `virtual_server.http3.http_client_profile.kind` | [virtual_server.http3.http_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-a08a549efcb756630d46caf0562c8ae3d8ba79e0caf443dc1de73355a4cd98c0) |
| `virtual_server.http3.http_client_profile.name` | [virtual_server.http3.http_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-69c49bdd9ba32b00437f5fe62c1ed0328e28b9e5b28597932e593ccc60c72a33) |
| `virtual_server.http3.http_client_profile.namespace` | [virtual_server.http3.http_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-789e73b40dacf8a18cd57afd26b01f1fcf0fea0403c01c980b24a037ef689236) |
| `virtual_server.http3.http_client_profile.tenant` | [virtual_server.http3.http_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-a65df17694572abe989e93ca8d8e503b4b5c333d3ac39f990902e5fec9b8ab42) |
| `virtual_server.http3.http_client_profile.uid` | [virtual_server.http3.http_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-6ecf1e806d3064c15d60818b0c417637accd448088a2dc3213c1fc9c656a8034) |
| `virtual_server.http3.http_server_profile` | [virtual_server.http3.http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-e9da22cf52ac6df0e313c16c640a2e85b733bb71aa5bcae661ce5d65db9d7424) |
| `virtual_server.http3.http_server_profile.kind` | [virtual_server.http3.http_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-111d9a5e6284823ea2f30caff78e087a42d3c32a46d7a6609eb9f1574fc1bf7e) |
| `virtual_server.http3.http_server_profile.name` | [virtual_server.http3.http_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-07df6918dcdff9ff7f9c956fe7ed3662714af25bf37d8d3f8edad2c87b575e37) |
| `virtual_server.http3.http_server_profile.namespace` | [virtual_server.http3.http_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-a7ad5bdb5bce221058651701445caf32b44bf4519bac3ddec6355e07e8c128c1) |
| `virtual_server.http3.http_server_profile.tenant` | [virtual_server.http3.http_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-8c943699acbd8b65173840a9c950db83106493f52c64898ade6d8e2039ab79da) |
| `virtual_server.http3.http_server_profile.uid` | [virtual_server.http3.http_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-5f885262e3bf5b7029b87b60187a1b4ab2bc337300bf03b3a9302fe711e8efe3) |
| `virtual_server.http3.quic_profile` | [virtual_server.http3.quic_profile](data-sources--application_profiles--reference--group-002.md#canonical-b50e8a2eec57ed79fa1950be4865f9706f3a7d22f5f10988d017df3388e27137) |
| `virtual_server.http3.quic_profile.kind` | [virtual_server.http3.quic_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-63f41adb59f2bdd84f6923bdfc9f41c9563bd19d5531f0833c9d4804ed5b80d5) |
| `virtual_server.http3.quic_profile.name` | [virtual_server.http3.quic_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-15262949d436631a03dc62127a28342377c1d92b53f92355c06deb0b4d4189b9) |
| `virtual_server.http3.quic_profile.namespace` | [virtual_server.http3.quic_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-b857560b8b16d3bc873f12c788047b8f9dc94f76e8d9c1cd1c1520a863a82495) |
| `virtual_server.http3.quic_profile.tenant` | [virtual_server.http3.quic_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-9e9fb74b79a10d0cf7ec602952ba9f8f6e16bed43454c39019acb75fe4057f41) |
| `virtual_server.http3.quic_profile.uid` | [virtual_server.http3.quic_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-8b828f3973545bea6c0fbce252fc3f8f8a9a7f7cb27f83a8be824734836f0ae4) |
| `virtual_server.http3.server_ssl_profile` | [virtual_server.http3.server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-a2e7936f9d7cb4e74f14f2a394bec22d0386db6d5be7d6e0ce7c59a5cff08a43) |
| `virtual_server.http3.server_ssl_profile.kind` | [virtual_server.http3.server_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-06cb7d1d5383b35318ee40487179b9f4a4f3d4605c3166198dbfb0b55917ed76) |
| `virtual_server.http3.server_ssl_profile.name` | [virtual_server.http3.server_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-4e39c3ce0109b220b8f3be8591545016263585f55974626e228b7171a2c3b1d1) |
| `virtual_server.http3.server_ssl_profile.namespace` | [virtual_server.http3.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-a8dc133b17618d8636ce04534f0405a04ad801ff2e469ffcd477fd8a0e981410) |
| `virtual_server.http3.server_ssl_profile.tenant` | [virtual_server.http3.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-19fd4072c5abb1c0b86db8daaeb388446f8019f3a4fa1dcf6f84176ccdf935cf) |
| `virtual_server.http3.server_ssl_profile.uid` | [virtual_server.http3.server_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-f99acca0cf79c1762b86dfc8a536d862a6884db0d93a45c62e4828877122e302) |
| `virtual_server.http3.tcp_server_profile` | [virtual_server.http3.tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-959687a2571b8aa04770a89bcf8a57170286fecdf391e1e4b98ad4642f55a86d) |
| `virtual_server.http3.tcp_server_profile.kind` | [virtual_server.http3.tcp_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-bf7a34d8f639b21838edd1e5bb266a3ae8ae2bdf6485c2a7c4a3462fdedfb8c1) |
| `virtual_server.http3.tcp_server_profile.name` | [virtual_server.http3.tcp_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-b9a75ec583315b8eef735c7224696b3fdfe2762e8c7b2e0d6b19829255c1515b) |
| `virtual_server.http3.tcp_server_profile.namespace` | [virtual_server.http3.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-965b89684d5da78027a927448455e8c737d0bffe4f15f0acc3a380f432da3aaa) |
| `virtual_server.http3.tcp_server_profile.tenant` | [virtual_server.http3.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-b9a6a5b9e027e315db57fb96330cda7242fa2eeff9baef6cf5a59d99b1b7000b) |
| `virtual_server.http3.tcp_server_profile.uid` | [virtual_server.http3.tcp_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-fd14f0436cf609ac4a2d163b734822591deff2fe37ac90b11bce2e681315563a) |
| `virtual_server.http3.udp_client_profile` | [virtual_server.http3.udp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-d455d2e53cb96825d6cd79d9195cf5edd883d522a5a74ae99a2cf7fbd8635269) |
| `virtual_server.http3.udp_client_profile.kind` | [virtual_server.http3.udp_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-a381ad507f416241c6412706080ecc5b2ed2cd13fc012e191ab9eeb4c3672709) |
| `virtual_server.http3.udp_client_profile.name` | [virtual_server.http3.udp_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-970aeb4e827cc02d63bf1975ff1f36d488566c0cf3b2b146b75061f9a15d42e0) |
| `virtual_server.http3.udp_client_profile.namespace` | [virtual_server.http3.udp_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-a34f28c9b3fafef8b345bbd7c0312203c041075eca1187d6c03d817eed9427b1) |
| `virtual_server.http3.udp_client_profile.tenant` | [virtual_server.http3.udp_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-20c4e627d01cb49e1e9884f72ae8fff26f154bd64cd4342954b10a141d5262ed) |
| `virtual_server.http3.udp_client_profile.uid` | [virtual_server.http3.udp_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-df276a148c4bed5388121a127fc955c15fd1d25798c122f4801500b76586eb8d) |
| `virtual_server.http3.udp_server_profile` | [virtual_server.http3.udp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-03e78a67a865975e00df35c6fc6ecc36f4d7901daae8cb33f5ae7f41a6de3463) |
| `virtual_server.http3.udp_server_profile.kind` | [virtual_server.http3.udp_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-1e2bdafdc7f35c1e1f0b980a7b5f10930f597ffdb01291676a8a69d4385b1123) |
| `virtual_server.http3.udp_server_profile.name` | [virtual_server.http3.udp_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-99ad14f06ef68e0b7498c1f579016b8fe0232f645e30b9b76b3f8bf0a90d22e9) |
| `virtual_server.http3.udp_server_profile.namespace` | [virtual_server.http3.udp_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-d71558dd0013a6608cc2386a6916ea10d565cd841aa9e65c671f459f090a52a2) |
| `virtual_server.http3.udp_server_profile.tenant` | [virtual_server.http3.udp_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-978514a2159618d6be71d6d98e55fd16288786fc07a2c7418cd80abf2453bf1d) |
| `virtual_server.http3.udp_server_profile.uid` | [virtual_server.http3.udp_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-6da45e49e79b962745dbefffe9421df37824154fe20f8190d6d328d36d2b4b7b) |
| `virtual_server.https` | [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-d4590d5492772762dcd475e0f257fdb82dd3a054627a0738dc4f0a13978965db) |
| `virtual_server.https.client_ssl_profile` | [virtual_server.https.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-5ee982bf36322fe9f914ead4f15768695c61a7aa42ac737432d767aeac0da4b7) |
| `virtual_server.https.client_ssl_profile.kind` | [virtual_server.https.client_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-5467672c814dd823424b77ddf32537a01c3ad96b70a16cb2742b59a9439f5060) |
| `virtual_server.https.client_ssl_profile.name` | [virtual_server.https.client_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-e73f7703e9c1231810de98bcb5fb85445290849ec06696388d26e138e4a95942) |
| `virtual_server.https.client_ssl_profile.namespace` | [virtual_server.https.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-8fd42f31ad1273c70c2963dafe51b4a9176233c77a0cfb883514b7aed323c460) |
| `virtual_server.https.client_ssl_profile.tenant` | [virtual_server.https.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-16544674d2ac2b82631e500d2cd66b44c9f4456e05f8c72e109514337dbe1104) |
| `virtual_server.https.client_ssl_profile.uid` | [virtual_server.https.client_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-d24e2c915b892193f109ff629814de0b7d403ceffa4a51c1ef0a63adc82182aa) |
| `virtual_server.https.http2_client_profile` | [virtual_server.https.http2_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-b0d6cc993ec4cd1c5c60643897fb53962932f49c554aa1170f2b124fa75ccf14) |
| `virtual_server.https.http2_client_profile.kind` | [virtual_server.https.http2_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-a3336b87d9ba29ea1c83298d7ebc0ceadf9f4dee7822bc1dd3b7f9a59d13bc53) |
| `virtual_server.https.http2_client_profile.name` | [virtual_server.https.http2_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-c818848fcd6fbb9351fbf64df48d33c03b5f70d2ad3e842d58e22a2d09403483) |
| `virtual_server.https.http2_client_profile.namespace` | [virtual_server.https.http2_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-8d40102d4e443c6a5c6634be9efcb7447745dce6ff63d0d7642b93c1b46a367e) |
| `virtual_server.https.http2_client_profile.tenant` | [virtual_server.https.http2_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-92ed7ecef32a7ce4d4161521d2fce018e0bf62c42a8b54b9f9819992bdd6c20e) |
| `virtual_server.https.http2_client_profile.uid` | [virtual_server.https.http2_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-17ad9c332c2f7c621f7dfdd59723fddcc56aedcc8e2e65d7d08fc9ee200510eb) |
| `virtual_server.https.http2_server_profile` | [virtual_server.https.http2_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-cedc1b1efb6532ee0cc2414c5cccb67da37c211833832a853225c077daa91cd1) |
| `virtual_server.https.http2_server_profile.kind` | [virtual_server.https.http2_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-62c5e9fd6324cbeedcaffad5ee49cf2dce181a948c92cc5150753214e14c7b19) |
| `virtual_server.https.http2_server_profile.name` | [virtual_server.https.http2_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-ebf2dc4990ec58933361adaf5c759f97f2915aa634b6e657012990423742c840) |
| `virtual_server.https.http2_server_profile.namespace` | [virtual_server.https.http2_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-0fbb61df015bd9b18f41cc6bcbcd8b017de0656e2ad748c1531a1b73d20a0412) |
| `virtual_server.https.http2_server_profile.tenant` | [virtual_server.https.http2_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-d986d1d0a1ee7a4197d3a2902eecf6da6e7fd8629f0e6cfca9e19f45107dc0df) |
| `virtual_server.https.http2_server_profile.uid` | [virtual_server.https.http2_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-028f4b9d1f7ab90ef983ed22c28e421d233a5926009240cdc5ea151f2db4f805) |
| `virtual_server.https.http_client_profile` | [virtual_server.https.http_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-ddc17c08055c18109993a607e59e76af0393225f3c228b43fdfa08c667e834c3) |
| `virtual_server.https.http_client_profile.kind` | [virtual_server.https.http_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-35b4daf643bd40497ce7c4a1f405ae122fa0282d491803532b0fe4585fee388d) |
| `virtual_server.https.http_client_profile.name` | [virtual_server.https.http_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-e4622718da81b6ab913640e16d2177535859cf25b7136cbc2b649e5a583f9c46) |
| `virtual_server.https.http_client_profile.namespace` | [virtual_server.https.http_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-17e385808a2277afe1ee631ee68a115849edf62d5cf1b365dbb0039d5f149ea5) |
| `virtual_server.https.http_client_profile.tenant` | [virtual_server.https.http_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-c769e3a8f0816dd99fc5da9bd0bd91f7bdb50ec6d1500048cbfd79e6913101c2) |
| `virtual_server.https.http_client_profile.uid` | [virtual_server.https.http_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-26b0d1062156ab0972b57d5edf6dd2dc09641b5b5ec66bce962a18960029ec6c) |
| `virtual_server.https.http_server_profile` | [virtual_server.https.http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-dec9eda488ec8a960180066b26bbca0c038abf43a9f4f014cfcd89592d076588) |
| `virtual_server.https.http_server_profile.kind` | [virtual_server.https.http_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-b76e29ccfd6be22a2d7752b58ca78ce4d53ffa95b14044807368888061bb6eb5) |
| `virtual_server.https.http_server_profile.name` | [virtual_server.https.http_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-4f7de1a837c82b7e3afaef25c7713766dff443c197e97d32f914048d830d1f67) |
| `virtual_server.https.http_server_profile.namespace` | [virtual_server.https.http_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2364178b23216538d2b01b45afca31964db33584e28d302b271f6ed2119a8fda) |
| `virtual_server.https.http_server_profile.tenant` | [virtual_server.https.http_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-a4acbdaaee593c4bc03e5814eb77e408c1cd453d0348661712fe55341692f3ab) |
| `virtual_server.https.http_server_profile.uid` | [virtual_server.https.http_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-46bb3adc75666cc825e48237c0b0d3c51b2d541412251fc10571f28c66d1418c) |
| `virtual_server.https.ocsp_profile` | [virtual_server.https.ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-87091287bd839c53798f2c02171540c3af9a6338920296d53c343792561bb449) |
| `virtual_server.https.ocsp_profile.kind` | [virtual_server.https.ocsp_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-3f78313b58f41f5bf4ea7edc013bd9b05f61dcbbcf677d5825aea5acec2adfe2) |
| `virtual_server.https.ocsp_profile.name` | [virtual_server.https.ocsp_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-ac2f76107cd2fb8478b7cfaf13031b8d1c2d7172ae79687846a381c14eb8d6a6) |
| `virtual_server.https.ocsp_profile.namespace` | [virtual_server.https.ocsp_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-3d68d3a18bcc911bbfd7e8178558d6cef2ddbfd706bde3df4c1b3329e85d883c) |
| `virtual_server.https.ocsp_profile.tenant` | [virtual_server.https.ocsp_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-fabf427876a2f6574836ea4be14cd382e43fc69447e6013b0043a7a2071bfa52) |
| `virtual_server.https.ocsp_profile.uid` | [virtual_server.https.ocsp_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-d4cec6258e0adaf178491b2082030454545e59440cc2a24a6a59f1e07001c360) |
| `virtual_server.https.server_ssl_profile` | [virtual_server.https.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-f3d6bf36cffb58d7a7dbcfadb032655eae87008585c63f86edff665e168906a2) |
| `virtual_server.https.server_ssl_profile.kind` | [virtual_server.https.server_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-83603a426b70080275a6d33cfb1fba108165ea1d12ddbdbe41a3f7c1c97df634) |
| `virtual_server.https.server_ssl_profile.name` | [virtual_server.https.server_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-7bb36a8b523c3bf3cf5e1e83b1da3fa468b39d3a8dfc6fc539bfe6d7d1a33cef) |
| `virtual_server.https.server_ssl_profile.namespace` | [virtual_server.https.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-00628772f6e9a747acf4e45a81a99bb8eb868a55ed6536a1fe29da10b0afb449) |
| `virtual_server.https.server_ssl_profile.tenant` | [virtual_server.https.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-e3f1ebaf5bed7745a5919e2612cd8546e09858fcdafc36dc6b9f92d8ea47395f) |
| `virtual_server.https.server_ssl_profile.uid` | [virtual_server.https.server_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-5dd28854200a7a322569ae81d9550a844910261ca01effcf856977b0b0219e2e) |
| `virtual_server.https.stream_profile` | [virtual_server.https.stream_profile](data-sources--application_profiles--reference--group-003.md#canonical-9159b8d42da7929cd3846360b09881856c3ba471f5eb18c05241202f3da79dc5) |
| `virtual_server.https.stream_profile.kind` | [virtual_server.https.stream_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-99a9ab68c36646f9ca9d0cc2964aa93d1a52e8e188504b07f48b9c4150967fb2) |
| `virtual_server.https.stream_profile.name` | [virtual_server.https.stream_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-6994a2a8b89afcd166c194b0791bd49b031690f34f2cd54edb690662a089713d) |
| `virtual_server.https.stream_profile.namespace` | [virtual_server.https.stream_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-7bc71d00a060e9fc6f82e242d247f5430e5f4493fa1f58437f932043e43194ad) |
| `virtual_server.https.stream_profile.tenant` | [virtual_server.https.stream_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-d616a5c45083d03e307ff5bac81328340a7e78178a2eeeabcbaed5e12523614a) |
| `virtual_server.https.stream_profile.uid` | [virtual_server.https.stream_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-0e6fffec07b17a047715a9fea8bd75f8bbfc0733e26040c5a532bd4dd2c41306) |
| `virtual_server.https.tcp_client_profile` | [virtual_server.https.tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-f542921f0bb847a58417d6670d00afc58412f338750f7d3aa4e6b61e839cf44a) |
| `virtual_server.https.tcp_client_profile.kind` | [virtual_server.https.tcp_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-c01d32c4bc9b4ecffd848aa0b03821594c54ee4cdbbf8f38580e6dfeebf3faed) |
| `virtual_server.https.tcp_client_profile.name` | [virtual_server.https.tcp_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-5316816f398975096f0329db7c3983714e4d64a0e7d526349d06897976b61e99) |
| `virtual_server.https.tcp_client_profile.namespace` | [virtual_server.https.tcp_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-ee10a6746546badae5342c16a921a16babbb7288400cbe2ffe9570565f050d2e) |
| `virtual_server.https.tcp_client_profile.tenant` | [virtual_server.https.tcp_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-4489a866eb442ef5fb1c5cfea10f743c02f678e2247e5e9d879769bf04b477f0) |
| `virtual_server.https.tcp_client_profile.uid` | [virtual_server.https.tcp_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-08e81b4f088414c05c1ed6c9679c5502b7c10049e2ac533243edf52f90197d4e) |
| `virtual_server.https.tcp_server_profile` | [virtual_server.https.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-bf6a3b12fd0251570d2fee396000d411ff7e7e4a1d3c82a7b1af949c75317422) |
| `virtual_server.https.tcp_server_profile.kind` | [virtual_server.https.tcp_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-071bb0b692f5bb69f9a93462eb9915f201a3963077ce3b28d09b99f28e706fab) |
| `virtual_server.https.tcp_server_profile.name` | [virtual_server.https.tcp_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1111f87675cd6a6437454f231949b2e0351c2560e0b6916efcd63fd040e21d91) |
| `virtual_server.https.tcp_server_profile.namespace` | [virtual_server.https.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-e753c24a70ce44f56071ba03abcd57d00a33bc69f38e1b3350a53c0dde321d22) |
| `virtual_server.https.tcp_server_profile.tenant` | [virtual_server.https.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-fee442f391679ed3caa0e8acebb1c571b5d1d381db32ba0fc48290a22387c60a) |
| `virtual_server.https.tcp_server_profile.uid` | [virtual_server.https.tcp_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-6bdc4df7bfa9cd2dcb14a08ae310f0a92158700752c5a23cd535b76c20f32cf9) |
| `virtual_server.https.websocket_client_profile` | [virtual_server.https.websocket_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-d458aaa7cd4c3b57a43d0f2b347b04de2bf6ea2c0b8abc31a7bbff84b0e7d77d) |
| `virtual_server.https.websocket_client_profile.kind` | [virtual_server.https.websocket_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-9d3bd9350990ebcda6514edc832838f02897ff096719f1a6b7eea75260620ee2) |
| `virtual_server.https.websocket_client_profile.name` | [virtual_server.https.websocket_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-fceece888635c30c9763cc902782df97470d8478af7a82e7a70f5b1cdf3544e4) |
| `virtual_server.https.websocket_client_profile.namespace` | [virtual_server.https.websocket_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-148d99bf8efc51daad5728650b6b36459876f3b2d0373266a47a89fc6de81292) |
| `virtual_server.https.websocket_client_profile.tenant` | [virtual_server.https.websocket_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-444dcf30ef004905b057fb6d028cecd43b6ffd4be4dd2c766a5dde7b43220ad8) |
| `virtual_server.https.websocket_client_profile.uid` | [virtual_server.https.websocket_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-73b93473c0acfe0b4edfeebf82deed9ecb421614162c869a4ee722fefedc65b1) |
| `virtual_server.https.websocket_server_profile` | [virtual_server.https.websocket_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-86571b7a55e8c33a958a06494cbd77f1519b55ca11d838c124a00332e4f8400f) |
| `virtual_server.https.websocket_server_profile.kind` | [virtual_server.https.websocket_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-173ee6154284ade44cd5c10b3c2e43f939e2221017d6eb39054b8ba4ed9c4bae) |
| `virtual_server.https.websocket_server_profile.name` | [virtual_server.https.websocket_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-4900917f6ff6e6bc285c2c7742f834071aa0ea224ef8d2427e1ff7c96ed90c4b) |
| `virtual_server.https.websocket_server_profile.namespace` | [virtual_server.https.websocket_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-6da0ae28f2ae011a20cc029d56aade976357b922119b464554786927b396c6cb) |
| `virtual_server.https.websocket_server_profile.tenant` | [virtual_server.https.websocket_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-063da4d3c6bc63e7018adfc431f7941f82618485a66b0d3a9b4ad7424d7fd94e) |
| `virtual_server.https.websocket_server_profile.uid` | [virtual_server.https.websocket_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-20cc66aaa78073189612dcbecb93c669666caa87a095416d3d9ce5fecd7b1681) |
| `virtual_server.immediate_action_on_service_down` | [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-24f59946a3d58f9754f08a2e67462e6b9f8a0985f17241d8c05e8bd5bdfa33d4) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](data-sources--application_profiles--reference--group-003.md#canonical-ef0ed12167543ebba930c555e39a68faba075ef6200abd2eb49872ad6eb8ef1b) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](data-sources--application_profiles--reference--group-003.md#canonical-c0f77183775d8a5679e3fa2bebd2c4effa7f3cbf1ff8c2d2588ea70cff5991dc) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](data-sources--application_profiles--reference--group-003.md#canonical-cddb9284ff3df29dd074fb333c13bf15d54071e1071e619c775fd379c2dee06f) |
| `virtual_server.last_hop_pool` | [virtual_server.last_hop_pool](data-sources--application_profiles--reference--group-003.md#canonical-7dfbdf1f7dc13e226b471d7bf9dbc910b5d21e4e2bc8d7b2d406eac51f0b43a1) |
| `virtual_server.last_hop_pool.kind` | [virtual_server.last_hop_pool.kind](data-sources--application_profiles--reference--group-003.md#canonical-f8b09b3656cad01fdeae4a1f7f9fba938bfa425513fd4e2299c6387b7495fafe) |
| `virtual_server.last_hop_pool.name` | [virtual_server.last_hop_pool.name](data-sources--application_profiles--reference--group-003.md#canonical-5079228e397112c36a01c694a798a009dae5b3a10f91b0373a1626b22343d26c) |
| `virtual_server.last_hop_pool.namespace` | [virtual_server.last_hop_pool.namespace](data-sources--application_profiles--reference--group-003.md#canonical-4aafd1a5c6c0411a5f1b251795ecdb7dffa42e155d119e0bab0f74e9e7a3555e) |
| `virtual_server.last_hop_pool.tenant` | [virtual_server.last_hop_pool.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3a3945ae58fc1137aa54dbebf09b1f882e2df9ff0483c8b30a48c252f6264e60) |
| `virtual_server.last_hop_pool.uid` | [virtual_server.last_hop_pool.uid](data-sources--application_profiles--reference--group-003.md#canonical-2ab5eb34d4b33e195fb97687719c4088f41a22426fb3bd7da2b656df8c1bbba5) |
| `virtual_server.nat64` | [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-d6aa6fe895906e316c1d2f52574869c46bd0aab9012c237b2d86bbcc398455df) |
| `virtual_server.nat64.nat64_disable` | [virtual_server.nat64.nat64_disable](data-sources--application_profiles--reference--group-003.md#canonical-fc240bbed6629c5e62589d6f424a8aeb7b4d6eb9c8f0d23144056c7fe4696e73) |
| `virtual_server.nat64.nat64_enable` | [virtual_server.nat64.nat64_enable](data-sources--application_profiles--reference--group-003.md#canonical-bbf30fd71593225aea4a4a36d80ed3f96b08685aa69ae17c0a2e19b7a70fd17b) |
| `virtual_server.port_translation` | [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-2c1a9a7dc9518624b1ed03a4160943fcc843e743f8e29e9ef81b323d02be6d47) |
| `virtual_server.port_translation.port_translation_disable` | [virtual_server.port_translation.port_translation_disable](data-sources--application_profiles--reference--group-003.md#canonical-6bfc5b2bc3deb4a5e7f9b5c5855f9d7f5d87006eea110759fb60ba8e64d5aa85) |
| `virtual_server.port_translation.port_translation_enable` | [virtual_server.port_translation.port_translation_enable](data-sources--application_profiles--reference--group-003.md#canonical-5b248fabc30abbbd809c4ca12eb25f1375999bf88acb488f7f2b10461ad1d606) |
| `virtual_server.request_logging_profile` | [virtual_server.request_logging_profile](data-sources--application_profiles--reference--group-003.md#canonical-4d530c7ac5195f5df8e87d097dfb37073fb53f5ac6fc8c695f9aa1f4a8b9c182) |
| `virtual_server.request_logging_profile.kind` | [virtual_server.request_logging_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-d516f11b7ab24db2d4a09d761b126baa6926b235afe21cd6e7a41edfaaf4163d) |
| `virtual_server.request_logging_profile.name` | [virtual_server.request_logging_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-0e7499e10e90607a4a3e18de29aefa396cd69047f57b389c42afa59425e4c826) |
| `virtual_server.request_logging_profile.namespace` | [virtual_server.request_logging_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-bfb6d8aac0620447bf6b6cb044c01e0d7603ff2ec87342f6b03177ea7e2656a4) |
| `virtual_server.request_logging_profile.tenant` | [virtual_server.request_logging_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-b9e41430aa0a1d1e36406bca5475c489624d6a2b746154d9768f92633bfdd1f3) |
| `virtual_server.request_logging_profile.uid` | [virtual_server.request_logging_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-6527c9961712f454e771a9a8daa43ccfe80c9b882ca03aa018b65c61d05f5be0) |
| `virtual_server.source_port` | [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-ad363b517828af74b2f7de0607ed43073ab239752d743a549b36f3fcc847d40e) |
| `virtual_server.source_port.source_port_change` | [virtual_server.source_port.source_port_change](data-sources--application_profiles--reference--group-003.md#canonical-4fb102546ee7433007d57646ddee19ebdd648aa79b2b3bcc20f9734116144591) |
| `virtual_server.source_port.source_port_preserve` | [virtual_server.source_port.source_port_preserve](data-sources--application_profiles--reference--group-003.md#canonical-da91e57ce73f7e99acedeeef22892a41b34f21fc63a01bc31fe8b7bfedb4b63b) |
| `virtual_server.source_port.source_port_preserve_strict` | [virtual_server.source_port.source_port_preserve_strict](data-sources--application_profiles--reference--group-003.md#canonical-93b5d4c9ab521cefdb5df5d99c459a4a8c5127b0129049cfa15962003c4177d8) |
| `virtual_server.statistics_profile` | [virtual_server.statistics_profile](data-sources--application_profiles--reference--group-003.md#canonical-8c63ada2758f01b939e0d1e1ef09aecd8f10abb3945d84e973e82b982e1e6a87) |
| `virtual_server.statistics_profile.kind` | [virtual_server.statistics_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-b21123a591a3a0185a49c00c3db2d1938f0fed586514246c9a5f520eef8ac7f1) |
| `virtual_server.statistics_profile.name` | [virtual_server.statistics_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-649caae6df455908facdf6b01b6c75fc83bc5b03dabbaa5d57ddf944a1ce9bdd) |
| `virtual_server.statistics_profile.namespace` | [virtual_server.statistics_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-efc17eb45f5dfbab6c2bb1df657a670797e69834d93ce4a02818770429265fe4) |
| `virtual_server.statistics_profile.tenant` | [virtual_server.statistics_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-e0383c1ef3c9388180333756d0428519fc42d245bf81328b32a6320e8ddfd3b1) |
| `virtual_server.statistics_profile.uid` | [virtual_server.statistics_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-d126ce4776172715ccf9e1d2159ba689e3aae18756cba9f0fcf68bbbaf31bcbc) |
| `virtual_server.tcp` | [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-fed3b36f37e7c02ef01066cef90f506b4200d146691c163d934491714451bed0) |
| `virtual_server.tcp.client_ssl_profile` | [virtual_server.tcp.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-6becd4ce7fc44af63bdd7bd2a5337ac39609fc2edd60384df022d7a0a0ff5489) |
| `virtual_server.tcp.client_ssl_profile.kind` | [virtual_server.tcp.client_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-6811c599e8bf8411c76fe469e97ee3c50a05ec4fa4542281767b7912fe508b35) |
| `virtual_server.tcp.client_ssl_profile.name` | [virtual_server.tcp.client_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-61bbd436ad7fbe7051a82c41306382c2f6d0688cb6f82a89a7ad5b8a5ad01462) |
| `virtual_server.tcp.client_ssl_profile.namespace` | [virtual_server.tcp.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-1a2785e0f2a08687dcf12d69cea302abbf1386a88f4de8043bd3f932a764eebc) |
| `virtual_server.tcp.client_ssl_profile.tenant` | [virtual_server.tcp.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-461fa089a5957a838ea7178ceb6be3cdf509d19770c54f284ed9047001e93d1a) |
| `virtual_server.tcp.client_ssl_profile.uid` | [virtual_server.tcp.client_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-991a15b602076359e1f770748d3b707d8bfe8ddf87e2b840910864f93b12fb9a) |
| `virtual_server.tcp.ocsp_profile` | [virtual_server.tcp.ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-ea5e2eba3da843952d83ef1d58125020784cbe09c83b5687c6e70bb166bf588c) |
| `virtual_server.tcp.ocsp_profile.kind` | [virtual_server.tcp.ocsp_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-9f4cd9f54b82832b309f6c38cf8fba8cc3296ce3fa39f73223980024e4e21a3e) |
| `virtual_server.tcp.ocsp_profile.name` | [virtual_server.tcp.ocsp_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-73c92c8c1002ec35a44af767cbe2dd502ce9f665066ec5286526f54dfe7dec3a) |
| `virtual_server.tcp.ocsp_profile.namespace` | [virtual_server.tcp.ocsp_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-dcb08e05d34e14396334e1031f26405909b454fdefa3214fe6aca20da25bd3c5) |
| `virtual_server.tcp.ocsp_profile.tenant` | [virtual_server.tcp.ocsp_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3a2731fb22e0b952d0473100a46aa47ababed01db36542e068fba7e36abcbb6f) |
| `virtual_server.tcp.ocsp_profile.uid` | [virtual_server.tcp.ocsp_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-ef7352c2aa96b88465b6492a78d95f13ed848c15cfea76ae10b9ea2b79ef2962) |
| `virtual_server.tcp.server_ssl_profile` | [virtual_server.tcp.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-7190c5c4c134028b638b0785b8ae92f8215cc96f575a3e1d394f5720b7cb205e) |
| `virtual_server.tcp.server_ssl_profile.kind` | [virtual_server.tcp.server_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-dba3204329536526c42cb36e102676af2d46948ec422b6f3e1b6847eefb0afd0) |
| `virtual_server.tcp.server_ssl_profile.name` | [virtual_server.tcp.server_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-e82c9f044af52c325cd32957a3f61eb42f952eee539c68ca85b3eb171ea59a1a) |
| `virtual_server.tcp.server_ssl_profile.namespace` | [virtual_server.tcp.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-d5ad219b232cfa683ecc27ed8d124012481411ad520005c5c2f2c7e267863f7b) |
| `virtual_server.tcp.server_ssl_profile.tenant` | [virtual_server.tcp.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-8a523c6fafb0d46270ae4fb0f901b8285de3e9ffa2d5f5e09fbc95260e1cccfd) |
| `virtual_server.tcp.server_ssl_profile.uid` | [virtual_server.tcp.server_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-1380e0d9b387c3c90ee1df25dd0348f3ef8966126de5fe71d2a6c4754368ad0a) |
| `virtual_server.tcp.tcp_client_profile` | [virtual_server.tcp.tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2f709128592e3af6ceb0b049905d4c8b84bc61257894947d89536ef44c66e294) |
| `virtual_server.tcp.tcp_client_profile.kind` | [virtual_server.tcp.tcp_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-cea36b67b0b9992ca66fd77a4de6e25ccbaf8dd789f59f667dd84f658e2a464b) |
| `virtual_server.tcp.tcp_client_profile.name` | [virtual_server.tcp.tcp_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-339cea565067082078abad2ddb0a5a5d6cd74b9bddc6677352a673278dc1df4d) |
| `virtual_server.tcp.tcp_client_profile.namespace` | [virtual_server.tcp.tcp_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-818389ef58313e1e4e58a54de540443cb2fe63cd9fb68044af577396e63bfd21) |
| `virtual_server.tcp.tcp_client_profile.tenant` | [virtual_server.tcp.tcp_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-664320f73d93bf8c61cdb1ac0ca7e8e1c15dec76d13bd42aefaf41d754fbf45a) |
| `virtual_server.tcp.tcp_client_profile.uid` | [virtual_server.tcp.tcp_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-2fa6b3c5662d9355393deff203f45d3f572d91b54890d032692e200c281a4da4) |
| `virtual_server.tcp.tcp_server_profile` | [virtual_server.tcp.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-ab44608225a6ab07a35eaaa98010b5ee0a454d0df672599fd170fe7a969dcfe9) |
| `virtual_server.tcp.tcp_server_profile.kind` | [virtual_server.tcp.tcp_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-bd0bbe56573f934579b9834b4ba824ef8925c0f631584913185142e18266c148) |
| `virtual_server.tcp.tcp_server_profile.name` | [virtual_server.tcp.tcp_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-83347ebd1043c69f339775ce94bdfb66c6de01778d208ea248f9564b6b6609d0) |
| `virtual_server.tcp.tcp_server_profile.namespace` | [virtual_server.tcp.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-faa2112423a167e5586e12c02ddd13d07e54788234b1d63db9d8063292b1250b) |
| `virtual_server.tcp.tcp_server_profile.tenant` | [virtual_server.tcp.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-396c0164845bc25d7d477c33b49184007bae206b5c353cfd6a3add209162152d) |
| `virtual_server.tcp.tcp_server_profile.uid` | [virtual_server.tcp.tcp_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-d68b9e1c328b65fb3cce0050fb12b189a6129d62eb3a6b2e6d2d9dac147050d0) |
| `virtual_server.udp` | [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-ec4acd18442e0e69df2ef446188dfaacef92f9137e352e59ad0b651b89e711ae) |
| `virtual_server.udp.client_ssl_profile` | [virtual_server.udp.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-ab54d95f11fc2ee97cca03519e5ab6eaa168b9cc48e77c05447fd88decc34716) |
| `virtual_server.udp.client_ssl_profile.kind` | [virtual_server.udp.client_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-3ac464f4a5922e72a84223e9dbb59a4caafc7420abf965a45d25b755ed903b9f) |
| `virtual_server.udp.client_ssl_profile.name` | [virtual_server.udp.client_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1e4ecef1b49112cae943af2c88afbe26269b567917a487fdb30d11209acfeb25) |
| `virtual_server.udp.client_ssl_profile.namespace` | [virtual_server.udp.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-644ef6ec0583c430dafb65342f13d2075a0a6591668fee22798a5ec352c691c4) |
| `virtual_server.udp.client_ssl_profile.tenant` | [virtual_server.udp.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-25d5a383ea7155c2d103b444895cb03d9501272995623a85501cf14b2379a9b0) |
| `virtual_server.udp.client_ssl_profile.uid` | [virtual_server.udp.client_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-a5431232b819578f61cd8289689d06a886270d1e2f33fdce859d6dc32dc9b6d4) |
| `virtual_server.udp.server_ssl_profile` | [virtual_server.udp.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-113c6c7fbc1ec2f04c809603d5bd35cc52490f8d9d64deaaa77765bcb122e5d0) |
| `virtual_server.udp.server_ssl_profile.kind` | [virtual_server.udp.server_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-beec88a2b11075f5e18edff61da77733b1059650335da63c1ecdf50288e5c356) |
| `virtual_server.udp.server_ssl_profile.name` | [virtual_server.udp.server_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-a9cd946b1f38f186f702b8d1fda417d9847a812c2ddf5ff37b0a2ef57118749d) |
| `virtual_server.udp.server_ssl_profile.namespace` | [virtual_server.udp.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-508d560d98eeb784b1cc1a82c620c7f767da089524611eabb117624580546145) |
| `virtual_server.udp.server_ssl_profile.tenant` | [virtual_server.udp.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-6b698e52808ff6c37a68bda6c03456d0e6df4b24b6bb024987b4d22e7224ad04) |
| `virtual_server.udp.server_ssl_profile.uid` | [virtual_server.udp.server_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-ad72e40adb053a0d8f9b4934802eae5d00cb27c7a6e984fd246bd016121514d3) |
| `virtual_server.udp.udp_client_profile` | [virtual_server.udp.udp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-4a287d25dd2f51853eac7bfc346f70b2c33101f37d4d142b0275987bab0093d0) |
| `virtual_server.udp.udp_client_profile.kind` | [virtual_server.udp.udp_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-fde67c0e8239a86bdea73fbed61c54aeb34094c964604ce4e33c022f3b9bf08f) |
| `virtual_server.udp.udp_client_profile.name` | [virtual_server.udp.udp_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-78afd2af2586cbd82b0a7420e4620abc9fa017628d91ec37aebce715f874fa22) |
| `virtual_server.udp.udp_client_profile.namespace` | [virtual_server.udp.udp_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-b367ea08a04834bae4bf87c6c934a3b99c4b4ca5f819a7ce6778288df9dbcbdd) |
| `virtual_server.udp.udp_client_profile.tenant` | [virtual_server.udp.udp_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-55dc3f2796ce12cc3ecc8f7a2905aa1cf5d4cd4fed73c9caf139069afb8338eb) |
| `virtual_server.udp.udp_client_profile.uid` | [virtual_server.udp.udp_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-dbee2710dea23677e4d9c68675867c61a63d2d5a8055a2ee513b89d9ee1c846f) |
| `virtual_server.udp.udp_server_profile` | [virtual_server.udp.udp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-6e92e9fae03d12e44a42d646f5ef3e621c3827c8b10da3ffc912cee3cec7a35f) |
| `virtual_server.udp.udp_server_profile.kind` | [virtual_server.udp.udp_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-31ba8ff8262f4af9af428d96b0f6b13385ae283f5f651afb9b60f7982d7865ba) |
| `virtual_server.udp.udp_server_profile.name` | [virtual_server.udp.udp_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-715d9c3e85718431be75de5c94927bb9fdd6ab1fb922b4594f87cfd1f56fb977) |
| `virtual_server.udp.udp_server_profile.namespace` | [virtual_server.udp.udp_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-50e6cd84cb7a6202fe774e173df025ba5e27adc665f392bb22a3eec870dcf1fb) |
| `virtual_server.udp.udp_server_profile.tenant` | [virtual_server.udp.udp_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-08bfd6347868017170671eff6fbd78b8eb195e222dbaa38f938c0ab5466fff99) |
| `virtual_server.udp.udp_server_profile.uid` | [virtual_server.udp.udp_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-f27b48285694654d1bb1b9759d1b5919f4e5c621a16e71645bf9110e6a8a5d2e) |
| `virtual_server.virtual_server_state` | [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-89b19c385bb0df18a103d7cdde9054f96f70b81b27afcc611e2ae38962137b90) |
| `virtual_server.virtual_server_state.state_disabled` | [virtual_server.virtual_server_state.state_disabled](data-sources--application_profiles--reference--group-003.md#canonical-bfb0fbd34a39406765d93e7ea7128ff07984685a2a47757d91a083beb53af121) |
| `virtual_server.virtual_server_state.state_enabled` | [virtual_server.virtual_server_state.state_enabled](data-sources--application_profiles--reference--group-003.md#canonical-9713f599f7c135068f02155ddd77a98abe3da7c62a31f2b43d4462c0527cd32c) |
| `virtual_server.vs_score` | [virtual_server.vs_score](data-sources--application_profiles--reference--group-001.md#canonical-51a437adb20e64820dce32be1cddadfac3c7d009cb042089c829feda5cc50c4f) |

<a id="canonical-8b94e0e4d81a2fbbb616f38c90d01744f84bfe4ed0fdbf48cd5b545619542a6f"></a>

## Next pages — Property reference / 56b0abd5bcb9 / 11

- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-c96b791db73eb7ed3c59393c5beb519e56f6b98af5a2999dbfbbef69414cbabb)
- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-29d3d0446d62faa0d39154381930b796864b6b683593f6fa1f025b812be1937e)
- [irules](data-sources--application_profiles--reference--group-001.md#canonical-b03852e8035d88ad972f06db07b19a393462dd0e06b97aff9ac57a85e610479b)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-c96b791db73eb7ed3c59393c5beb519e56f6b98af5a2999dbfbbef69414cbabb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b825fee9e8e13352ecad687f940221874a88bcbc76ecc96cc5e79f739e81870c"></a>

## advanced_tcp_profile — advanced_tcp_profile / fbb5ca5f5444 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- advanced_tcp_profile

<a id="canonical-d3767388aea02ec0ca5025809e8b9ea513436b560577f5b21d781ce28158f42b"></a>

Type: `"single"`. Computed.

Configuration parameter for advanced tcp profile.

Upstream description:

BIG-IP Advanced TCP Profile.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tcp_advanced_profile_choice": "[\"disable_tcp_advanced_profile\",\"enable_tcp_advanced_profile\"]"
}
```

<a id="canonical-9f934dcf36c7bee3313153478c974d8d05ec4eb1ab72e8e14afd46a1f2a3278b"></a>

## Direct properties — advanced_tcp_profile / fbb5ca5f5444 / 3

- [disable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-65da3288a4fd806d96bc676a3a236ca6d585819053bfaf2ea71855de57509179): complete subsection reference.

- [enable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-6134876561c66f7c3bc0c4c09e3d11b12600bf0abca0d310923b3c05537abe45): complete subsection reference.

<a id="canonical-309523fc67299740b1b18125c238c76236157cfbd3e554b30451737d40bd8333"></a>

## Next pages — advanced_tcp_profile / fbb5ca5f5444 / 4

- [advanced_tcp_profile.disable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-65da3288a4fd806d96bc676a3a236ca6d585819053bfaf2ea71855de57509179)
- [advanced_tcp_profile.enable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-6134876561c66f7c3bc0c4c09e3d11b12600bf0abca0d310923b3c05537abe45)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-65da3288a4fd806d96bc676a3a236ca6d585819053bfaf2ea71855de57509179"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8195f7f9fba7b63d8ea2c956f7043d8d141c6ba8943a664c6e2d1d6a092e6897"></a>

## advanced_tcp_profile.disable_tcp_advanced_profile — advanced_tcp_profile.disable_tcp_advanced_profile / 9880f7f35e8a / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-c96b791db73eb7ed3c59393c5beb519e56f6b98af5a2999dbfbbef69414cbabb)
- advanced_tcp_profile.disable_tcp_advanced_profile

<a id="canonical-0d26b2193786af7aa79e0db3600a2e41886b063eee854d74aa5addef66f9383b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable tcp advanced profile.

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

<a id="canonical-39166a53d47ae877d0395ea6f94fa4bc3eb02cb3d9ed9d9558f46ca311fa8bbd"></a>

## Direct properties — advanced_tcp_profile.disable_tcp_advanced_profile / 9880f7f35e8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-704403d7c9ed5aeecf945adc7d550b2c8da7ed1c85101db06d6baddf03aa57a0"></a>

## Next pages — advanced_tcp_profile.disable_tcp_advanced_profile / 9880f7f35e8a / 4

- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-c96b791db73eb7ed3c59393c5beb519e56f6b98af5a2999dbfbbef69414cbabb)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-6134876561c66f7c3bc0c4c09e3d11b12600bf0abca0d310923b3c05537abe45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7daf043f2e643436d13683cab80cf0bc6d919f5eed731003c498b6b87939b35"></a>

## advanced_tcp_profile.enable_tcp_advanced_profile — advanced_tcp_profile.enable_tcp_advanced_profile / ed56c3e4712b / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-c96b791db73eb7ed3c59393c5beb519e56f6b98af5a2999dbfbbef69414cbabb)
- advanced_tcp_profile.enable_tcp_advanced_profile

<a id="canonical-9e6ebe48e8f5f7a29e5a34daccaba29516ab8384272cdf414e783ef391dba3fd"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable tcp advanced profile.

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

<a id="canonical-ada784e398833685e85b6234a18264b039e32bd7836c33141c7443518cb8e506"></a>

## Direct properties — advanced_tcp_profile.enable_tcp_advanced_profile / ed56c3e4712b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a591d4f566f0bc45cee68a1e9ec50b069598115f3098f7b99aaec440c76a251"></a>

## Next pages — advanced_tcp_profile.enable_tcp_advanced_profile / ed56c3e4712b / 4

- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-c96b791db73eb7ed3c59393c5beb519e56f6b98af5a2999dbfbbef69414cbabb)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-29d3d0446d62faa0d39154381930b796864b6b683593f6fa1f025b812be1937e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9f86f67c5a171d819cdabb2391f80b2469d87d757b9b0a0e2e98834a2cc05f4"></a>

## ddos_profile — ddos_profile / f09077bb779a / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- ddos_profile

<a id="canonical-b8c30b3a65e6d9609c5c15e10affae0439226c346239ecc25fe5608ef8470fc8"></a>

Type: `"single"`. Computed.

Configuration parameter for ddos profile.

Upstream description:

BIG-IP DDoS Protection Rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

<a id="canonical-25d33b385b208782715fe493712cc3da3b6e6510e698e19e6009effba8be998a"></a>

## Direct properties — ddos_profile / f09077bb779a / 3

- [disable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-714f973e1231815bfc6c95de5ab98c2ff63b7c8c49ead62abc7df179876bf94c): complete subsection reference.

- [enable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-9cfe5451f6ddf70cc4c96d1ae89b340bc09920615b6da5ebb63d1b9033df3244): complete subsection reference.

<a id="canonical-e9d28ce051db4a21c13b086755875e54ac75cfaadfe1e12b0158bcc1fe56b9c0"></a>

## Next pages — ddos_profile / f09077bb779a / 4

- [ddos_profile.disable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-714f973e1231815bfc6c95de5ab98c2ff63b7c8c49ead62abc7df179876bf94c)
- [ddos_profile.enable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-9cfe5451f6ddf70cc4c96d1ae89b340bc09920615b6da5ebb63d1b9033df3244)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-714f973e1231815bfc6c95de5ab98c2ff63b7c8c49ead62abc7df179876bf94c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f9ab34877f3e7b868df4b84f0a28fe2e469d2f5d910daa305cb24d49ec40468"></a>

## ddos_profile.disable_ddos_mitigation — ddos_profile.disable_ddos_mitigation / 60da480196a5 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-29d3d0446d62faa0d39154381930b796864b6b683593f6fa1f025b812be1937e)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-ba0a9fa3a16fff5595f46101036ae725e701da0f88906e3bfc6cc6b3f5fca749"></a>

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

<a id="canonical-e2e1432b0acc2c5b7a9d17496c1ed29a4094e92d6451407848551991698842e4"></a>

## Direct properties — ddos_profile.disable_ddos_mitigation / 60da480196a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2086a55994276833745d3f12def13ccf2c54c87c9f3d6b7b46e6713fbdb628dc"></a>

## Next pages — ddos_profile.disable_ddos_mitigation / 60da480196a5 / 4

- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-29d3d0446d62faa0d39154381930b796864b6b683593f6fa1f025b812be1937e)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-9cfe5451f6ddf70cc4c96d1ae89b340bc09920615b6da5ebb63d1b9033df3244"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e62040a1728a3620d3f32c66f7aa606cf394e3a20f45937cabdb1a3d04140f67"></a>

## ddos_profile.enable_ddos_mitigation — ddos_profile.enable_ddos_mitigation / 8c453195a4c2 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-29d3d0446d62faa0d39154381930b796864b6b683593f6fa1f025b812be1937e)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-61c54997e65f4ea0911b9bf8f4da246106d1643e36664a7e10d7ee912196e5e6"></a>

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

<a id="canonical-6f706c9356361b78b71426bf0b9a28c7171d483865f7977ee98f62a7a1994a9b"></a>

## Direct properties — ddos_profile.enable_ddos_mitigation / 8c453195a4c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3efba164a9920f3f5a84dca14bf450247a20f9c3ea57dea43917418b650a124"></a>

## Next pages — ddos_profile.enable_ddos_mitigation / 8c453195a4c2 / 4

- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-29d3d0446d62faa0d39154381930b796864b6b683593f6fa1f025b812be1937e)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-b03852e8035d88ad972f06db07b19a393462dd0e06b97aff9ac57a85e610479b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d14bd434ed347fce245a7330f7e388e9960d61645b49a5d16f878123722808b"></a>

## irules — irules / d3e99e2b2f35 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- irules

<a id="canonical-9f505918e95bbf193101e06dc82e460ae7385da1d3b0e487261980ebfc4c2fa0"></a>

Type: `"list"`. Computed.

OPTIONS for attaching iRules to BIG-IP Proxy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c1e7bf64625d19d50dcea4e278b9b6a221f1f274d8b32e140e8a600d1ddd2e2a"></a>

## Direct properties — irules / d3e99e2b2f35 / 3

<a id="canonical-ad0d9d62279d6fb30c0456cb22ef8c40aafb7cab14347c374325bfb8d41723c7"></a>

<a id="canonical-2322262e668345ec3ffddd2415e0360f254794b2f80908822149335ff5a01e5d"></a>

## kind property — irules / d3e99e2b2f35 / 4

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

<a id="canonical-ce7fd1101c941540c77558402346710e11496dbac0f0189034db95a46db22d8a"></a>

<a id="canonical-c6502cdd524fcac2d813ab48d6ba634e67cb3db3403138f0cbaf7fe76391fb5b"></a>

## name property — irules / d3e99e2b2f35 / 5

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

<a id="canonical-c537b88d8371e4da4a11e7917efcdb6f972e38b8bc3603453241ff6ffe283574"></a>

<a id="canonical-5cb09300708e49f9a2149921eb68bfe12a6a66caf51b7452ab8bfbdc0e32fa4c"></a>

## namespace property — irules / d3e99e2b2f35 / 6

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

<a id="canonical-175f2bea23fd12acd42a1af0d58283c7ee0a7379d6a7cc91a9398b7bda102f44"></a>

<a id="canonical-61d01b21d1604fc3dfcc7a0c2ccd31b2520d4ff779bf3a04b977a8527ce813f7"></a>

## tenant property — irules / d3e99e2b2f35 / 7

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

<a id="canonical-a3ef076b67a7821b4a918797dc2fd8ae2de2e9ce3b4d05d01b42f973020f4853"></a>

<a id="canonical-74963cfea3e91bc370e0cccd57006736f795cfefc1d6ec9cf722c96cd98cec3c"></a>

## uid property — irules / d3e99e2b2f35 / 8

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

<a id="canonical-f9ca7ca09979b8824f66ff12b9f139cf49faf2b726d198054fb6db6ebf9c9287"></a>

## Next pages — irules / d3e99e2b2f35 / 9

- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af8d1cc97b4d0f56b1608aa3b1a1dab1d275581613a0a5833d327e5cbe726736"></a>

## virtual_server — virtual_server / b6ab15ca52a3 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- virtual_server

<a id="canonical-6d8b283c9385612b65995a6a3cdd81b8bb0f78278a1f7e926c3382872d8b7260"></a>

Type: `"single"`. Computed.

Specifies configuration related to virtual server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_server_type": "[\"http\",\"http3\",\"https\",\"tcp\",\"udp\"]"
}
```

<a id="canonical-57cf5d0a666ed882b4d9706ec0fded9ecf033c8c8098d63919aa14c0e242247a"></a>

## Direct properties — virtual_server / b6ab15ca52a3 / 3

- [access_profile](data-sources--application_profiles--reference--group-001.md#canonical-3c56c26bbdeaa9090011d18c377c0e458b10dde91c0d8d6965e15855283ed77e): complete subsection reference.

- [address_translation](data-sources--application_profiles--reference--group-001.md#canonical-7d8ab78d0c9441bc0bd8be5c5ceac2bd7577be661a7b51e0ce84e800a9674dd1): complete subsection reference.

- [auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f): complete subsection reference.

- [clone_pool_client](data-sources--application_profiles--reference--group-002.md#canonical-5d73316e5d127760389f4de33f8f2c93af8832f8e4b48154b679834d4f5408a7): complete subsection reference.

- [clone_pool_server](data-sources--application_profiles--reference--group-002.md#canonical-8ec873d6c734de5225aca2d4908de716882031bc8c0863c1dd26084af348358f): complete subsection reference.

<a id="canonical-c3d44cfa848788c88ac81d3ccac710087e8bc390d8b3e8e1d113a3e340f5380e"></a>

<a id="canonical-6b219730b205e7fb415c1c1356c84f156651c8625beedd1ca4ed3096c5e60b41"></a>

## connection_limit property — virtual_server / b6ab15ca52a3 / 4

Type: `"number"`. Computed.

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The.

Upstream description:

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The default is 0.

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-c6b2ef93baf1a91d75e3f5835daea2e407683a962196a7384304944dd3eaebcd"></a>

<a id="canonical-d8a57ce2b92ea9e25b0184d059080b6fe7cf90378458bed83cb3a5c7060c5c01"></a>

## connection_rate_limit property — virtual_server / b6ab15ca52a3 / 5

Type: `"number"`. Computed.

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where..

Upstream description:

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where
connection requests flood a virtual server. Setting this to 0 turns off connection limits. The
default is 0.

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f): complete subsection reference.

- [default_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-0366d1dd265cc60d0524a3f453cf707e64e9634ca637fbca12d78eb6b8463161): complete subsection reference.

- [default_pool](data-sources--application_profiles--reference--group-002.md#canonical-3ff91b0f7877a94f4912abc43e110504ba594829ab9540f690ae63e862c83ced): complete subsection reference.

- [fallback_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-f7bc990e9c3e8f027c8654d878046f88f3442028146200f0ba190a9e1bf9222c): complete subsection reference.

- [fix_profile](data-sources--application_profiles--reference--group-002.md#canonical-fe1e80eb31e642e61346da9d7773a14ffdb260dee3a620e9e85a918fd920f453): complete subsection reference.

- [http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd): complete subsection reference.

- [http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367): complete subsection reference.

- [https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f): complete subsection reference.

- [immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8): complete subsection reference.

- [last_hop_pool](data-sources--application_profiles--reference--group-003.md#canonical-00c49ee343c420b534167e37120e0f65d3241be0e65da401f8ab8495f7ac87cb): complete subsection reference.

- [nat64](data-sources--application_profiles--reference--group-003.md#canonical-f8055b0439e9d63c4c9239eb747d5b3451270728bf042bcf326f9859bf1d731a): complete subsection reference.

- [port_translation](data-sources--application_profiles--reference--group-003.md#canonical-4f39f6eca33ddb7fb761dd7e59c61a6579e3840ccc7e8c7e76eccd69c2ca2c68): complete subsection reference.

- [request_logging_profile](data-sources--application_profiles--reference--group-003.md#canonical-eee319a7a1bc5aea591a514b66ed5b4f70fbc9ff7220feea6b49aaec53d79ede): complete subsection reference.

- [source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3): complete subsection reference.

- [statistics_profile](data-sources--application_profiles--reference--group-003.md#canonical-2c3db883e913163179c953dff01dfd2f3fdccc8020837b14abe11c6e7a6b105a): complete subsection reference.

- [tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051): complete subsection reference.

- [udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92): complete subsection reference.

- [virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-300638b1ed4116f684a232f38ce18576e112d93f7687ddea795783b9e07a871c): complete subsection reference.

<a id="canonical-51a437adb20e64820dce32be1cddadfac3c7d009cb042089c829feda5cc50c4f"></a>

<a id="canonical-e2a76034de81d57638e92e27b5e435f816f289a1d1f5240db2aac7797ce5ef3d"></a>

## vs_score property — virtual_server / b6ab15ca52a3 / 6

Type: `"number"`. Computed.

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The , meaning that no additional metric is applied
for the virtual server.

Upstream description:

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The default is 0, meaning that no additional
metric is applied for the virtual server.

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-c5bdd04df19855e0d373b849446a9910b02fd6483614d57ecc1371f58e5e4513"></a>

## Next pages — virtual_server / b6ab15ca52a3 / 7

- [virtual_server.access_profile](data-sources--application_profiles--reference--group-001.md#canonical-3c56c26bbdeaa9090011d18c377c0e458b10dde91c0d8d6965e15855283ed77e)
- [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-7d8ab78d0c9441bc0bd8be5c5ceac2bd7577be661a7b51e0ce84e800a9674dd1)
- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f)
- [virtual_server.clone_pool_client](data-sources--application_profiles--reference--group-002.md#canonical-5d73316e5d127760389f4de33f8f2c93af8832f8e4b48154b679834d4f5408a7)
- [virtual_server.clone_pool_server](data-sources--application_profiles--reference--group-002.md#canonical-8ec873d6c734de5225aca2d4908de716882031bc8c0863c1dd26084af348358f)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [virtual_server.default_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-0366d1dd265cc60d0524a3f453cf707e64e9634ca637fbca12d78eb6b8463161)
- [virtual_server.default_pool](data-sources--application_profiles--reference--group-002.md#canonical-3ff91b0f7877a94f4912abc43e110504ba594829ab9540f690ae63e862c83ced)
- [virtual_server.fallback_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-f7bc990e9c3e8f027c8654d878046f88f3442028146200f0ba190a9e1bf9222c)
- [virtual_server.fix_profile](data-sources--application_profiles--reference--group-002.md#canonical-fe1e80eb31e642e61346da9d7773a14ffdb260dee3a620e9e85a918fd920f453)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-b06d67cf7343f5f2ed5be727e404999c334cb21ee5bbf4264c4bb778cb4569f8)
- [virtual_server.last_hop_pool](data-sources--application_profiles--reference--group-003.md#canonical-00c49ee343c420b534167e37120e0f65d3241be0e65da401f8ab8495f7ac87cb)
- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-f8055b0439e9d63c4c9239eb747d5b3451270728bf042bcf326f9859bf1d731a)
- [virtual_server.port_translation](data-sources--application_profiles--reference--group-003.md#canonical-4f39f6eca33ddb7fb761dd7e59c61a6579e3840ccc7e8c7e76eccd69c2ca2c68)
- [virtual_server.request_logging_profile](data-sources--application_profiles--reference--group-003.md#canonical-eee319a7a1bc5aea591a514b66ed5b4f70fbc9ff7220feea6b49aaec53d79ede)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-003.md#canonical-407c4608699a81f4e3bb3778a1b26581ec0c108c04dbd9a3849e526729f9a8e3)
- [virtual_server.statistics_profile](data-sources--application_profiles--reference--group-003.md#canonical-2c3db883e913163179c953dff01dfd2f3fdccc8020837b14abe11c6e7a6b105a)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-003.md#canonical-7e68c1acd62fcb8bd84098a29c63253343102ec83007254f96ec7153f910c051)
- [virtual_server.udp](data-sources--application_profiles--reference--group-003.md#canonical-d51961fd24b168399f71bb2d9281ab5fa247ef8ede7a7b0dee891f5c77b92d92)
- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-003.md#canonical-300638b1ed4116f684a232f38ce18576e112d93f7687ddea795783b9e07a871c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-3c56c26bbdeaa9090011d18c377c0e458b10dde91c0d8d6965e15855283ed77e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dad0a3c34d1105b941e5ddac9ac20708742b1ea18899172fc60ffc3ffe3bc677"></a>

## virtual_server.access_profile — virtual_server.access_profile / 59d3653d3120 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.access_profile

<a id="canonical-f18c4d94352b1ed57cb7e55085e3d01eaa7094df7dcead0f0cee184453bb76ad"></a>

Type: `"list"`. Computed.

Specifies an access policy that determines the authentication rules and access controls applied to
user sessions for this virtual server.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-97bf2f4f9eb0f4f1504c39e39736ffb34281c2351b9ba151b30bd6e3498f5cfe"></a>

## Direct properties — virtual_server.access_profile / 59d3653d3120 / 3

<a id="canonical-addd341b46ac881e8c5b97dd28d7bc7b288efc54b5481b822f0aa7bbfaaa4654"></a>

<a id="canonical-a5f21874fde75bfe5dcc95d556f971b3923f75c302f282b27406977e9595635d"></a>

## kind property — virtual_server.access_profile / 59d3653d3120 / 4

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

<a id="canonical-41302f6760a20249e1d4869da6deb91fdadc1ae85873a1b2bf35321385d14674"></a>

<a id="canonical-1c3052008788aace5c53a3bceee0cbb89d0e95d7ffbfc03a7c7fec6414e25c0c"></a>

## name property — virtual_server.access_profile / 59d3653d3120 / 5

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

<a id="canonical-62defd0bb3dbb951810770ff982d02885a94a7f164663eadb76304093a7b33d5"></a>

<a id="canonical-2a2f0066fdea586b78f06da689fd4ada7044d0f1713bce919e258294822999c1"></a>

## namespace property — virtual_server.access_profile / 59d3653d3120 / 6

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

<a id="canonical-e797019088039352cf9fb282a43ee6b3bf83e601b583333f33fa477989ad25c3"></a>

<a id="canonical-dcd25dd87bed0bbcce50e39a79d2fe61d3b6f9ee58ad780663d098bd348afded"></a>

## tenant property — virtual_server.access_profile / 59d3653d3120 / 7

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

<a id="canonical-afd917f455def6751b0b3d22fef529d27d30deea328c2867b4e4729834ec66de"></a>

<a id="canonical-0407b2252ea46df9825bfda8c3545b3d1eb02de1d62bc704ff8431efd8c5baf5"></a>

## uid property — virtual_server.access_profile / 59d3653d3120 / 8

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

<a id="canonical-3c34051cdd70e3a1d6a995359dd46ff792c8ad13df6d21505f38689d5d33c6ba"></a>

## Next pages — virtual_server.access_profile / 59d3653d3120 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-7d8ab78d0c9441bc0bd8be5c5ceac2bd7577be661a7b51e0ce84e800a9674dd1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48a87a70c197d43ea5548e157dd586a92b1b92d4b5890846d5d436f17362601c"></a>

## virtual_server.address_translation — virtual_server.address_translation / 979d7faf31ac / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.address_translation

<a id="canonical-d5cb7b1fa01499170bf47bf0de2596b184f2a9ac70b0d5ab247153464679d3e0"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address.

Upstream description:

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address. The default is
enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_translation_choice": "[\"address_translation_disable\",\"address_translation_enable\"]"
}
```

<a id="canonical-2ff09887a5f7d5342f74748cd1e65545a930d5654b2609c5bccf61aae3ce5b7a"></a>

## Direct properties — virtual_server.address_translation / 979d7faf31ac / 3

- [address_translation_disable](data-sources--application_profiles--reference--group-001.md#canonical-9516c481f82132ba1f53fa6a6c8c0220dee7ad5b286bc23d812c16bfcc36c19a): complete subsection reference.

- [address_translation_enable](data-sources--application_profiles--reference--group-001.md#canonical-55fe55d10494b3de8691a5a42f198ac8476abad2c29d9c66791bcd13543fce3d): complete subsection reference.

<a id="canonical-f99bb8af018a06fa67568a4600040175b4069aad56f6c2612b498d86bafc847c"></a>

## Next pages — virtual_server.address_translation / 979d7faf31ac / 4

- [virtual_server.address_translation.address_translation_disable](data-sources--application_profiles--reference--group-001.md#canonical-9516c481f82132ba1f53fa6a6c8c0220dee7ad5b286bc23d812c16bfcc36c19a)
- [virtual_server.address_translation.address_translation_enable](data-sources--application_profiles--reference--group-001.md#canonical-55fe55d10494b3de8691a5a42f198ac8476abad2c29d9c66791bcd13543fce3d)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-9516c481f82132ba1f53fa6a6c8c0220dee7ad5b286bc23d812c16bfcc36c19a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4eafed0db7dba1c588039556a7a0011084010b5a9135366cf4e0312a49274b41"></a>

## virtual_server.address_translation.address_translation_disable — virtual_server.address_translation.address_translation_disable / 72e598a7fed4 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-7d8ab78d0c9441bc0bd8be5c5ceac2bd7577be661a7b51e0ce84e800a9674dd1)
- virtual_server.address_translation.address_translation_disable

<a id="canonical-3f4060e9e0e5b960f3b4300906422a8755cfcddcfdce9b0897e0c242ee89b670"></a>

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

<a id="canonical-abc213bfef3ef00b6943ca961f793929eac1758ded2dd53d0cc06bcb86dfde0a"></a>

## Direct properties — virtual_server.address_translation.address_translation_disable / 72e598a7fed4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cdb61dbe68cf83a0a8904f77b924cef5acc19133d241bfa83fcdaadc290ec95d"></a>

## Next pages — virtual_server.address_translation.address_translation_disable / 72e598a7fed4 / 4

- [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-7d8ab78d0c9441bc0bd8be5c5ceac2bd7577be661a7b51e0ce84e800a9674dd1)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-55fe55d10494b3de8691a5a42f198ac8476abad2c29d9c66791bcd13543fce3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73cf956d7f1274b69c5c13af65c2d7010b036292fad91bf68ae44a8cea13ac90"></a>

## virtual_server.address_translation.address_translation_enable — virtual_server.address_translation.address_translation_enable / fc7783286e6a / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-7d8ab78d0c9441bc0bd8be5c5ceac2bd7577be661a7b51e0ce84e800a9674dd1)
- virtual_server.address_translation.address_translation_enable

<a id="canonical-b40c437c53279a11948161c044f26ba354ac407b7b8fcb23b991e3f8fc692711"></a>

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

<a id="canonical-beafc38a00d4ae2cd6a9244722e493d3fcc2e08fe1f6edc81c394948ab8f3646"></a>

## Direct properties — virtual_server.address_translation.address_translation_enable / fc7783286e6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66123cc79a295ba81dad6ed22487e7942d7445f04c8b704c2fba58d0218fc872"></a>

## Next pages — virtual_server.address_translation.address_translation_enable / fc7783286e6a / 4

- [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-7d8ab78d0c9441bc0bd8be5c5ceac2bd7577be661a7b51e0ce84e800a9674dd1)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6292a8629c9712fbc3875533525277c0b79de72c56b46c0021bb060bbd3f5deb"></a>

## virtual_server.auto_last_hop — virtual_server.auto_last_hop / b960e53332ee / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.auto_last_hop

<a id="canonical-937a6558a263dc0f40a43b93040407bcf1d1bea17e00581999961e95406c6dbc"></a>

Type: `"single"`. Computed.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auto_last_hop_choice": "[\"auto_last_hop_default\",\"auto_last_hop_disable\",\"auto_last_hop_enable\"]"
}
```

<a id="canonical-640c5865fcac3e516bd25e99788c4fbc457c8571d93a6f741c796db6e15adb6a"></a>

## Direct properties — virtual_server.auto_last_hop / b960e53332ee / 3

- [auto_last_hop_default](data-sources--application_profiles--reference--group-001.md#canonical-8a883e6200baf4775e2fddc484b69fdfec77a47acf4746cdef6a5659a47b1447): complete subsection reference.

- [auto_last_hop_disable](data-sources--application_profiles--reference--group-001.md#canonical-cbff403cb5e9bc187c648bc34b53a523ec94c8f8ff655f7ebcbe28e1476ae6e8): complete subsection reference.

- [auto_last_hop_enable](data-sources--application_profiles--reference--group-001.md#canonical-84761035cb05ffb88500d877300070d51f218f57a37c9527849dcfbf7248a3a0): complete subsection reference.

<a id="canonical-d00e5d59fcabfc2ca86a1d6edf090418b6af92d1cc089022192c93617917ae05"></a>

## Next pages — virtual_server.auto_last_hop / b960e53332ee / 4

- [virtual_server.auto_last_hop.auto_last_hop_default](data-sources--application_profiles--reference--group-001.md#canonical-8a883e6200baf4775e2fddc484b69fdfec77a47acf4746cdef6a5659a47b1447)
- [virtual_server.auto_last_hop.auto_last_hop_disable](data-sources--application_profiles--reference--group-001.md#canonical-cbff403cb5e9bc187c648bc34b53a523ec94c8f8ff655f7ebcbe28e1476ae6e8)
- [virtual_server.auto_last_hop.auto_last_hop_enable](data-sources--application_profiles--reference--group-001.md#canonical-84761035cb05ffb88500d877300070d51f218f57a37c9527849dcfbf7248a3a0)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-8a883e6200baf4775e2fddc484b69fdfec77a47acf4746cdef6a5659a47b1447"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97932112016403acf77b5910f8552c78880b671622aa90c37954fe55264e8da8"></a>

## virtual_server.auto_last_hop.auto_last_hop_default — virtual_server.auto_last_hop.auto_last_hop_default / 8c92158da912 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f)
- virtual_server.auto_last_hop.auto_last_hop_default

<a id="canonical-0f3d26554379edf0f708a65ddd7dac1fc63f4e38b7da543656ca0be810045171"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for auto last hop default.

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

<a id="canonical-8deba9bea935d146bb1424911d50a4b000daf89707d70ecc406ec7dc42393290"></a>

## Direct properties — virtual_server.auto_last_hop.auto_last_hop_default / 8c92158da912 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e2efdee492d32d534cd112767c5a036b541c949e6ec25f08b994c716ba4b87b"></a>

## Next pages — virtual_server.auto_last_hop.auto_last_hop_default / 8c92158da912 / 4

- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-cbff403cb5e9bc187c648bc34b53a523ec94c8f8ff655f7ebcbe28e1476ae6e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cef0baf2296d1e08380c3d8e310665982b16ed23cccbfadebe4c66c8836af16f"></a>

## virtual_server.auto_last_hop.auto_last_hop_disable — virtual_server.auto_last_hop.auto_last_hop_disable / d56a4587b51d / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f)
- virtual_server.auto_last_hop.auto_last_hop_disable

<a id="canonical-f46bae3d4532eb7ff3e67ae4bb710b04e61a785e17d8fb39fa881f9db5432173"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for auto last hop disable.

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

<a id="canonical-43dc5eeb6d2d2989492c68de38d1c4a41d37d1bd7243609b2d2bd55d31189f00"></a>

## Direct properties — virtual_server.auto_last_hop.auto_last_hop_disable / d56a4587b51d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84317d025a109677adc55b4912a25871c45d55edb7a59723565a2184ba7f972d"></a>

## Next pages — virtual_server.auto_last_hop.auto_last_hop_disable / d56a4587b51d / 4

- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-84761035cb05ffb88500d877300070d51f218f57a37c9527849dcfbf7248a3a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
