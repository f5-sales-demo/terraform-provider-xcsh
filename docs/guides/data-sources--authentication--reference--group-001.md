---
page_title: "xcsh_authentication reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication reference."
---

# xcsh_authentication reference

<a id="canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0285db351a14d253325d03608a2207c37548ec43a34f368b16da26aea285a2b7"></a>

## Property reference — Property reference / 975af69105bc / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- Property reference

<a id="canonical-6e19730be37004eb9f9680ba1ad1e0dbe662cd5538583d27338d8ed05848c503"></a>

## Direct properties — Property reference / 975af69105bc / 3

<a id="canonical-e63cfef4670586505891f4276b2c0c4fdefdede2ae93b8e66b2503b429621783"></a>

<a id="canonical-5538a91655595b299a81de0e58b48c1f3f09a1aba61a54197e250e07850b5d90"></a>

## annotations property — Property reference / 975af69105bc / 4

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

- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf): complete subsection reference.

<a id="canonical-13cb9cee2c1ae5f2a99d0f6487e47e33346403eebd67373f9371b53d3266ca60"></a>

<a id="canonical-43ec382f2beb3ed1ec43f8d6a276b836e438c17331a48bb735b35b5239ee5d63"></a>

## description property — Property reference / 975af69105bc / 5

Type: `"string"`. Computed.

Description of the Authentication.

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

<a id="canonical-bca984697dc08266175668ed12fe9062759a03b0b47d7b58296a3ed5f357b6c7"></a>

<a id="canonical-3d35c33cbf24fcf37b2b4aa772750a3477b927b0f1b0e0382efe2deca743c561"></a>

## id property — Property reference / 975af69105bc / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0d9258b5ae0e754e065a901acebd9fea1e2a9f315cf4ce3a74b2c2c71661cc48"></a>

<a id="canonical-108aaa3a77222aee56108123214a5e03add38d6c694aa94c2a8ad8a234ffe044"></a>

## labels property — Property reference / 975af69105bc / 7

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

<a id="canonical-dbc8c2eb12ee0c7c717bd6b67a8e83696ceba4c8ee9f14ed03a60c4c5fb5a6b6"></a>

<a id="canonical-2eb34bdf49b8ed6b94902c6271caf99d932e9e79ecc60fdc2d6c1fe63b0d72bd"></a>

## name property — Property reference / 975af69105bc / 8

Type: `"string"`. Required.

Name of the Authentication.

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

<a id="canonical-e2bad3102b6a9d7c07a989cd92b4608bccf55f85a69d88c6cc729d7c65f6732f"></a>

<a id="canonical-1394479b9557340f083e5cb5b75f12d5ed96b941cde0c3016b5aa39c4cd17bba"></a>

## namespace property — Property reference / 975af69105bc / 9

Type: `"string"`. Required.

Namespace where the Authentication exists.

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

- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b): complete subsection reference.

<a id="canonical-3f8efe4481cd92c889e0a4fff10e281762ef1e777d5dcaf40d4ccb80ea8ab39e"></a>

## All schema paths — Property reference / 975af69105bc / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--authentication--reference--group-001.md#canonical-e63cfef4670586505891f4276b2c0c4fdefdede2ae93b8e66b2503b429621783) |
| `cookie_params` | [cookie_params](data-sources--authentication--reference--group-001.md#canonical-beb4839ef3ed95bf27b928e28475ee311c957816814caf1d1979e55cf77ae2f7) |
| `cookie_params.auth_hmac` | [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-0111e869d0486635dbc22dc4b91d20f6278f74ca02e6cb8dbf775eed9881797d) |
| `cookie_params.auth_hmac.prim_key` | [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-a70f236589c5a038bd818553cfe9e395fc1919358e448c9db4d973dd301a1339) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-34caaadd474c2c2e3b3292138a451c02c7ca357869b00c965a8c6422a460bcb7) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](data-sources--authentication--reference--group-001.md#canonical-3d3cdd5a17587ca98429fa686a69a9341846fc52a51a3d0d3bbcac891183fede) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](data-sources--authentication--reference--group-001.md#canonical-a1b59a286253436dab8e0d2fef3c153c2d5c995c1f6d70c1c524e04e0957fdfa) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](data-sources--authentication--reference--group-001.md#canonical-0e2c0118b7785c17017aef179a2ee100e89623c5755de5c1ec6fef7398bbd936) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info` | [cookie_params.auth_hmac.prim_key.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-8af13ccc4d2786d66bf3f6a1071bd6f5d8d695f69c4cf4be3ca3f233a483bb9e) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](data-sources--authentication--reference--group-001.md#canonical-602a76e35f6861dbdbfe829c1605496f71c99f5ab29cadf5900f526de650edd1) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [cookie_params.auth_hmac.prim_key.clear_secret_info.url](data-sources--authentication--reference--group-001.md#canonical-93f1b220e9d64a92bb4457e845b3226fa7d8d4c9a8924f759648fefa6f553c80) |
| `cookie_params.auth_hmac.prim_key_expiry` | [cookie_params.auth_hmac.prim_key_expiry](data-sources--authentication--reference--group-001.md#canonical-37bdeac415a265cb607e89dfb4ba94c805283fbcc495e4682b1ade2ef30a300f) |
| `cookie_params.auth_hmac.sec_key` | [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-83a51fe89691ce6417d211c0e88e82c4b66c3a54a9069a22d7b9430c1d4ef9a4) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-71b57692c19574c34d63312775b1ce20d5f4710949567176ad43b5857cb55b93) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](data-sources--authentication--reference--group-001.md#canonical-5aa6964aa9ce125a41756fabc53215c79cd2c669a77a6ae6d63eca067dbc82aa) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](data-sources--authentication--reference--group-001.md#canonical-04b0bd013863cee8825ee40f17b955cc9ba3f7af289aad3b6c51169014b3d14f) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](data-sources--authentication--reference--group-001.md#canonical-8fc97292a837424382878cfe29d94c4ac50c7b53e0d40a9ac6f60ad587307582) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info` | [cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-48280ca6ba52d9bb7dd303791a009b4a578c12261e308a208148c161b20c804e) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](data-sources--authentication--reference--group-001.md#canonical-2d4be3b1feb1e92985d3bb9ad5914bf4734541c98c822396c655d92d2bbeb40d) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [cookie_params.auth_hmac.sec_key.clear_secret_info.url](data-sources--authentication--reference--group-001.md#canonical-d5778025c0e658e8c64978f0ed0adecf4723e4303f4f4d5b0e1f5c8b74ce2267) |
| `cookie_params.auth_hmac.sec_key_expiry` | [cookie_params.auth_hmac.sec_key_expiry](data-sources--authentication--reference--group-001.md#canonical-05613a838f004a8b0d0201fef006992f310c7ef45aff473050d9028d7ae0b7e5) |
| `cookie_params.cookie_expiry` | [cookie_params.cookie_expiry](data-sources--authentication--reference--group-001.md#canonical-bef566001b0c3a6d9049aca4b9782d43f360cb8d456e0aa9fe657c29b38ec2bd) |
| `cookie_params.cookie_refresh_interval` | [cookie_params.cookie_refresh_interval](data-sources--authentication--reference--group-001.md#canonical-16dddaf8eae795614195627188d52658a6325a9590f156cc40a5dc2ccbab6985) |
| `cookie_params.kms_key_hmac` | [cookie_params.kms_key_hmac](data-sources--authentication--reference--group-001.md#canonical-f63f63686c2cfb6a0af0a070a8205c443c010052f8326bfa521074123dc3f304) |
| `cookie_params.session_expiry` | [cookie_params.session_expiry](data-sources--authentication--reference--group-001.md#canonical-9e9e47a92583a8d87e90eaad206a5f2335bce193ffb065ac9833aea2ed52dda0) |
| `description` | [description](data-sources--authentication--reference--group-001.md#canonical-13cb9cee2c1ae5f2a99d0f6487e47e33346403eebd67373f9371b53d3266ca60) |
| `id` | [id](data-sources--authentication--reference--group-001.md#canonical-bca984697dc08266175668ed12fe9062759a03b0b47d7b58296a3ed5f357b6c7) |
| `labels` | [labels](data-sources--authentication--reference--group-001.md#canonical-0d9258b5ae0e754e065a901acebd9fea1e2a9f315cf4ce3a74b2c2c71661cc48) |
| `name` | [name](data-sources--authentication--reference--group-001.md#canonical-dbc8c2eb12ee0c7c717bd6b67a8e83696ceba4c8ee9f14ed03a60c4c5fb5a6b6) |
| `namespace` | [namespace](data-sources--authentication--reference--group-001.md#canonical-e2bad3102b6a9d7c07a989cd92b4608bccf55f85a69d88c6cc729d7c65f6732f) |
| `oidc_auth` | [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-46725115245c47fd91edf6523e8287b562b6a48e4b4a88400454bc10b354abef) |
| `oidc_auth.client_secret` | [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-a77b48f3c2b4c18dfa3d90c61ed1d46e05fe8d5de13c2653d4fb29d81a1fa232) |
| `oidc_auth.client_secret.blindfold_secret_info` | [oidc_auth.client_secret.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-f4142578ab0560c88fbb71e3d5f0ce0a11e0fe86cc92d0cf7395c307eda4132f) |
| `oidc_auth.client_secret.blindfold_secret_info.decryption_provider` | [oidc_auth.client_secret.blindfold_secret_info.decryption_provider](data-sources--authentication--reference--group-001.md#canonical-ab23c38784f93f6e037831463065a79339f1ffaae4e1c6dceee04dd5b982f59c) |
| `oidc_auth.client_secret.blindfold_secret_info.location` | [oidc_auth.client_secret.blindfold_secret_info.location](data-sources--authentication--reference--group-001.md#canonical-bd492b22b53b9436a9b424a2a38ba7e03f9642e60692b5a1ea1fc9e09aaf6e8f) |
| `oidc_auth.client_secret.blindfold_secret_info.store_provider` | [oidc_auth.client_secret.blindfold_secret_info.store_provider](data-sources--authentication--reference--group-001.md#canonical-746af60ac73855368b6ced7b6ff526ff88d858f6ef43303cfd2176ed1661b661) |
| `oidc_auth.client_secret.clear_secret_info` | [oidc_auth.client_secret.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-ea91951c8a428bbeb9d4b20df09f02d16f7df3881269837b1e00d24324fcb0a0) |
| `oidc_auth.client_secret.clear_secret_info.provider_ref` | [oidc_auth.client_secret.clear_secret_info.provider_ref](data-sources--authentication--reference--group-001.md#canonical-f9893a30eb89a61a5f651fdc64d5afb0bda57fa4676e4ddb10921e69179b786b) |
| `oidc_auth.client_secret.clear_secret_info.url` | [oidc_auth.client_secret.clear_secret_info.url](data-sources--authentication--reference--group-001.md#canonical-1477563430cbb8d14bbac786e8898de7803dffad7c1e30474df9508c4af6490f) |
| `oidc_auth.oidc_auth_params` | [oidc_auth.oidc_auth_params](data-sources--authentication--reference--group-001.md#canonical-23634d39d1c5e50c541df25052f8bae7427a1e98b8fcf1ba59562a3a920d9e4b) |
| `oidc_auth.oidc_auth_params.auth_endpoint_url` | [oidc_auth.oidc_auth_params.auth_endpoint_url](data-sources--authentication--reference--group-001.md#canonical-fbaf232449260b4de1553c8625b648b76a2f64190680dc207462765741a960f7) |
| `oidc_auth.oidc_auth_params.end_session_endpoint_url` | [oidc_auth.oidc_auth_params.end_session_endpoint_url](data-sources--authentication--reference--group-001.md#canonical-5052b78dd3055ccff9b7030e72ed59d31c439ff73ee1afafef06f54a49e8da6d) |
| `oidc_auth.oidc_auth_params.token_endpoint_url` | [oidc_auth.oidc_auth_params.token_endpoint_url](data-sources--authentication--reference--group-001.md#canonical-dfe7c4d8ee665270cacbcd73c4b999211b02aadeae57104f96fec5983bfb90fc) |
| `oidc_auth.oidc_client_id` | [oidc_auth.oidc_client_id](data-sources--authentication--reference--group-001.md#canonical-b5cb02d6209ee8fe970df52bfa40605d82fde063dad022cbd7b0dbc7e8373595) |
| `oidc_auth.oidc_well_known_config_url` | [oidc_auth.oidc_well_known_config_url](data-sources--authentication--reference--group-001.md#canonical-714e183c551c772d61fea1c4a7b1dee920788384b56b8fc98c21d29978c91953) |

<a id="canonical-dc534996169e0f30294921aefa4052863e4469edf42b67d4a4553785aa853a1e"></a>

## Next pages — Property reference / 975af69105bc / 11

- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1289515c80b4bd36a6c67e91df700efbecd92532a72ae6c0e33c931b9786981"></a>

## cookie_params — cookie_params / 0c3dd7e98c0d / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- cookie_params

<a id="canonical-beb4839ef3ed95bf27b928e28475ee311c957816814caf1d1979e55cf77ae2f7"></a>

Type: `"single"`. Computed.

Specifies different cookie related config parameters for authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_choice": "[\"auth_hmac\",\"kms_key_hmac\"]"
}
```

<a id="canonical-fc1dc3d6af2738bb616fd754ed0bd72c6f80b466d0c8f0c87953147779d778b2"></a>

## Direct properties — cookie_params / 0c3dd7e98c0d / 3

- [auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb): complete subsection reference.

<a id="canonical-bef566001b0c3a6d9049aca4b9782d43f360cb8d456e0aa9fe657c29b38ec2bd"></a>

<a id="canonical-62abd00a7d104c29516272cb94bd0df136b497b5888b5356556c1464e4c25233"></a>

## cookie_expiry property — cookie_params / 0c3dd7e98c0d / 4

Type: `"number"`. Computed.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request.

Upstream description:

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-16dddaf8eae795614195627188d52658a6325a9590f156cc40a5dc2ccbab6985"></a>

<a id="canonical-8f89545f73be1ebe204ef2a11403ff089f1b30a6da1c74f47928f8b72415fd5b"></a>

## cookie_refresh_interval property — cookie_params / 0c3dd7e98c0d / 5

Type: `"number"`. Computed.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session..

Upstream description:

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](data-sources--authentication--reference--group-001.md#canonical-5a39e95fe94b3946f40c8b2d88f08bf9b197fdbfa8a8919cd8fb47e65047d4fc): complete subsection reference.

<a id="canonical-9e9e47a92583a8d87e90eaad206a5f2335bce193ffb065ac9833aea2ed52dda0"></a>

<a id="canonical-7aecf57bfc98059241eb8edce671f4c8cb87687e32d681ff60cef17bf7cf8113"></a>

## session_expiry property — cookie_params / 0c3dd7e98c0d / 6

Type: `"number"`. Computed.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Upstream description:

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1296000,
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
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```

<a id="canonical-9d2b3c9e99ce58cc2d2d09f5f167098ee22b958e8317dab65a5d087d807325ee"></a>

## Next pages — cookie_params / 0c3dd7e98c0d / 7

- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- [cookie_params.kms_key_hmac](data-sources--authentication--reference--group-001.md#canonical-5a39e95fe94b3946f40c8b2d88f08bf9b197fdbfa8a8919cd8fb47e65047d4fc)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65488895ff4825d2dd797cf117d6496a186ad0d34fa0d4c9b19d8131b6479c7c"></a>

## cookie_params.auth_hmac — cookie_params.auth_hmac / ca07c710066b / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- cookie_params.auth_hmac

<a id="canonical-0111e869d0486635dbc22dc4b91d20f6278f74ca02e6cb8dbf775eed9881797d"></a>

Type: `"single"`. Computed.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9f491ccc87eed08827176e4b79de08531c6b798be0dcf9d82dca5d4dbee4c640"></a>

## Direct properties — cookie_params.auth_hmac / ca07c710066b / 3

- [prim_key](data-sources--authentication--reference--group-001.md#canonical-f93b1721175307f42114d53f68275e22a3c797392a6d1d7b47d99d4d725a6c05): complete subsection reference.

<a id="canonical-37bdeac415a265cb607e89dfb4ba94c805283fbcc495e4682b1ade2ef30a300f"></a>

<a id="canonical-86a4b9f0f0f98d5d92209520824e609ca3a025ce64a0f7d504b65691072dfe8c"></a>

## prim_key_expiry property — cookie_params.auth_hmac / ca07c710066b / 4

Type: `"string"`. Computed.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

Upstream description:

Primary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [sec_key](data-sources--authentication--reference--group-001.md#canonical-2fea82869c57ba9da95ec5787490c7a7afbb73e12573875bdbe35ea141435aa1): complete subsection reference.

<a id="canonical-05613a838f004a8b0d0201fef006992f310c7ef45aff473050d9028d7ae0b7e5"></a>

<a id="canonical-570df097e9c82ab48a9773ab62e61d0c790ee30ac990ad6ef20351998f5b16bb"></a>

## sec_key_expiry property — cookie_params.auth_hmac / ca07c710066b / 5

Type: `"string"`. Computed.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

Upstream description:

Secondary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-7b28e0451e51dfc4d93475f4f181801fd1cda9912cdfe695fd028d842e6ab2df"></a>

## Next pages — cookie_params.auth_hmac / ca07c710066b / 6

- [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-f93b1721175307f42114d53f68275e22a3c797392a6d1d7b47d99d4d725a6c05)
- [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-2fea82869c57ba9da95ec5787490c7a7afbb73e12573875bdbe35ea141435aa1)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-f93b1721175307f42114d53f68275e22a3c797392a6d1d7b47d99d4d725a6c05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0382eaee589804283b2e021828178ce2a7828550f64add0a13337844ce6cd10"></a>

## cookie_params.auth_hmac.prim_key — cookie_params.auth_hmac.prim_key / d88272213050 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- cookie_params.auth_hmac.prim_key

<a id="canonical-a70f236589c5a038bd818553cfe9e395fc1919358e448c9db4d973dd301a1339"></a>

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

<a id="canonical-56560951736e6a69594170feba51a6e109214e4f963615f77a8ead1c9b38cffe"></a>

## Direct properties — cookie_params.auth_hmac.prim_key / d88272213050 / 3

- [blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-7f4a99c5fe12895cffaa516aa46e31f67f77d62396e3156b355fca5625342190): complete subsection reference.

- [clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-12538acea386124e62d4356a9a98c144e489faf0b8308b7a81687f11a5649933): complete subsection reference.

<a id="canonical-f5c1a4873481ddb8adce4b65ccee686628f52183492952c06c7bca3cce5cb9f6"></a>

## Next pages — cookie_params.auth_hmac.prim_key / d88272213050 / 4

- [cookie_params.auth_hmac.prim_key.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-7f4a99c5fe12895cffaa516aa46e31f67f77d62396e3156b355fca5625342190)
- [cookie_params.auth_hmac.prim_key.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-12538acea386124e62d4356a9a98c144e489faf0b8308b7a81687f11a5649933)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-7f4a99c5fe12895cffaa516aa46e31f67f77d62396e3156b355fca5625342190"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1812291ddbccf6793ae025b6162267a8eef6e332f52cce7a3482a66a72a9906"></a>

## cookie_params.auth_hmac.prim_key.blindfold_secret_info — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 64e7054c5239 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-f93b1721175307f42114d53f68275e22a3c797392a6d1d7b47d99d4d725a6c05)
- cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-34caaadd474c2c2e3b3292138a451c02c7ca357869b00c965a8c6422a460bcb7"></a>

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

<a id="canonical-8b72c02e2d8088999f89013def0b087cc02d11c4d28b5473c30831da90b53f34"></a>

## Direct properties — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 64e7054c5239 / 3

<a id="canonical-3d3cdd5a17587ca98429fa686a69a9341846fc52a51a3d0d3bbcac891183fede"></a>

<a id="canonical-81a44236b4f7eb40459151441d04fbf2d3c869f4fda1b6371a7b00d7705838ad"></a>

## decryption_provider property — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 64e7054c5239 / 4

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

<a id="canonical-a1b59a286253436dab8e0d2fef3c153c2d5c995c1f6d70c1c524e04e0957fdfa"></a>

<a id="canonical-1c814e5943c8de2ab7fafb0a5c0d1a76854d8bf3357a9d38d6804d0a7400f679"></a>

## location property — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 64e7054c5239 / 5

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

<a id="canonical-0e2c0118b7785c17017aef179a2ee100e89623c5755de5c1ec6fef7398bbd936"></a>

<a id="canonical-838614c365fa7c8493503461c19c126f3a972219868b10ed7716fbc9520b298f"></a>

## store_provider property — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 64e7054c5239 / 6

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

<a id="canonical-b5adb3fa5ba4b8038d95bfd3aebe089f5bab2566db5e5a9aee59fac458c92361"></a>

## Next pages — cookie_params.auth_hmac.prim_key.blindfold_secret_info / 64e7054c5239 / 7

- [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-f93b1721175307f42114d53f68275e22a3c797392a6d1d7b47d99d4d725a6c05)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-12538acea386124e62d4356a9a98c144e489faf0b8308b7a81687f11a5649933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd5741e91c39742e33d235d09ddf7929ac060b6f0d57c9dc2a763248993b6519"></a>

## cookie_params.auth_hmac.prim_key.clear_secret_info — cookie_params.auth_hmac.prim_key.clear_secret_info / 3cb4e8153471 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-f93b1721175307f42114d53f68275e22a3c797392a6d1d7b47d99d4d725a6c05)
- cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-8af13ccc4d2786d66bf3f6a1071bd6f5d8d695f69c4cf4be3ca3f233a483bb9e"></a>

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

<a id="canonical-5a1fc1480d1ccdf824d520e52b4cce2f3bd38e35b5141f23df9c600d6097d50d"></a>

## Direct properties — cookie_params.auth_hmac.prim_key.clear_secret_info / 3cb4e8153471 / 3

<a id="canonical-602a76e35f6861dbdbfe829c1605496f71c99f5ab29cadf5900f526de650edd1"></a>

<a id="canonical-d344fafaebd49df1f14fcdbca39dc9be504b21e80ee901dc989fa646c50b930e"></a>

## provider_ref property — cookie_params.auth_hmac.prim_key.clear_secret_info / 3cb4e8153471 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-93f1b220e9d64a92bb4457e845b3226fa7d8d4c9a8924f759648fefa6f553c80"></a>

<a id="canonical-8aaa1453ee028aa25e9b0361384816cc16d8977621b52e5692e7071c6c857207"></a>

## url property — cookie_params.auth_hmac.prim_key.clear_secret_info / 3cb4e8153471 / 5

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

<a id="canonical-8bf4b80749cae05a86b6965f9320956ae739ac5f22536ae16e7f28fdabc7277d"></a>

## Next pages — cookie_params.auth_hmac.prim_key.clear_secret_info / 3cb4e8153471 / 6

- [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-f93b1721175307f42114d53f68275e22a3c797392a6d1d7b47d99d4d725a6c05)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-2fea82869c57ba9da95ec5787490c7a7afbb73e12573875bdbe35ea141435aa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c3a120224855a754f01e1ca55aef3865f1ee0f67378ecd2132d623df8774b77"></a>

## cookie_params.auth_hmac.sec_key — cookie_params.auth_hmac.sec_key / f7c302cac215 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- cookie_params.auth_hmac.sec_key

<a id="canonical-83a51fe89691ce6417d211c0e88e82c4b66c3a54a9069a22d7b9430c1d4ef9a4"></a>

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

<a id="canonical-0abd05bc35252169f75081989f3c7258eb754a8a6463050a6c867f65477ae40c"></a>

## Direct properties — cookie_params.auth_hmac.sec_key / f7c302cac215 / 3

- [blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-901f28d407b9cb6b6487dfd198bc21c9c7358b38634a1fe3480203aae65b12ca): complete subsection reference.

- [clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-c98f24cd44073967579b0621e46caa3a0b44d3994683ac8cdf357843bcf36b83): complete subsection reference.

<a id="canonical-dd7f94e306c5edf320a665d2786d1538bd2c061790a11457a65423013397d697"></a>

## Next pages — cookie_params.auth_hmac.sec_key / f7c302cac215 / 4

- [cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-901f28d407b9cb6b6487dfd198bc21c9c7358b38634a1fe3480203aae65b12ca)
- [cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-c98f24cd44073967579b0621e46caa3a0b44d3994683ac8cdf357843bcf36b83)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-901f28d407b9cb6b6487dfd198bc21c9c7358b38634a1fe3480203aae65b12ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e9250bca42aaec16d071ad15128547755b01a636f2c329a2b5de98ba966dd97"></a>

## cookie_params.auth_hmac.sec_key.blindfold_secret_info — cookie_params.auth_hmac.sec_key.blindfold_secret_info / 7a1ee86c937c / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-2fea82869c57ba9da95ec5787490c7a7afbb73e12573875bdbe35ea141435aa1)
- cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-71b57692c19574c34d63312775b1ce20d5f4710949567176ad43b5857cb55b93"></a>

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

<a id="canonical-a1cbb6362b4f13465a31604869c4fb1a7643e6a4996a5c384ab0bf51a6f1065f"></a>

## Direct properties — cookie_params.auth_hmac.sec_key.blindfold_secret_info / 7a1ee86c937c / 3

<a id="canonical-5aa6964aa9ce125a41756fabc53215c79cd2c669a77a6ae6d63eca067dbc82aa"></a>

<a id="canonical-794520c93159c4209b178a2d3b0a8de5bf955ebc33861e6c16b7a83af60eb0b5"></a>

## decryption_provider property — cookie_params.auth_hmac.sec_key.blindfold_secret_info / 7a1ee86c937c / 4

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

<a id="canonical-04b0bd013863cee8825ee40f17b955cc9ba3f7af289aad3b6c51169014b3d14f"></a>

<a id="canonical-415bdf0d0c3ee71c270966961346f975737317c94efb368c0e5ab4d72b8ef006"></a>

## location property — cookie_params.auth_hmac.sec_key.blindfold_secret_info / 7a1ee86c937c / 5

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

<a id="canonical-8fc97292a837424382878cfe29d94c4ac50c7b53e0d40a9ac6f60ad587307582"></a>

<a id="canonical-db30697ae8342574e83ffefdb6a7efa3b8e59175bf12b832ba8f5c4d9d608965"></a>

## store_provider property — cookie_params.auth_hmac.sec_key.blindfold_secret_info / 7a1ee86c937c / 6

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

<a id="canonical-8deb6d646c2805523ed4c9f53a9399d7b6379f971d2ce1d0ca39a83a1a3be9c6"></a>

## Next pages — cookie_params.auth_hmac.sec_key.blindfold_secret_info / 7a1ee86c937c / 7

- [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-2fea82869c57ba9da95ec5787490c7a7afbb73e12573875bdbe35ea141435aa1)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-c98f24cd44073967579b0621e46caa3a0b44d3994683ac8cdf357843bcf36b83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033c6351e58f86093e2b66b05f550c72ffdc3736c091f6042a8baa6ebe918d0"></a>

## cookie_params.auth_hmac.sec_key.clear_secret_info — cookie_params.auth_hmac.sec_key.clear_secret_info / e0cc538c5b50 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-b3ab273cc6686a00dde018b2bbdc2891c060a1eac4476c619dcfe40dc5a012fb)
- [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-2fea82869c57ba9da95ec5787490c7a7afbb73e12573875bdbe35ea141435aa1)
- cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-48280ca6ba52d9bb7dd303791a009b4a578c12261e308a208148c161b20c804e"></a>

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

<a id="canonical-deecef361769e131d8916701004c737d277660c81d96d9e10afbf7eb5d19cd95"></a>

## Direct properties — cookie_params.auth_hmac.sec_key.clear_secret_info / e0cc538c5b50 / 3

<a id="canonical-2d4be3b1feb1e92985d3bb9ad5914bf4734541c98c822396c655d92d2bbeb40d"></a>

<a id="canonical-9127078ed65ef8054d879cb6a235c248e3f028099d078ef5ce43cb4fcfae9f3c"></a>

## provider_ref property — cookie_params.auth_hmac.sec_key.clear_secret_info / e0cc538c5b50 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-d5778025c0e658e8c64978f0ed0adecf4723e4303f4f4d5b0e1f5c8b74ce2267"></a>

<a id="canonical-39ad9d4fd5441ceef65a63802181f9d6fa0a3ba8b4811a09e5a254a2721de6bf"></a>

## url property — cookie_params.auth_hmac.sec_key.clear_secret_info / e0cc538c5b50 / 5

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

<a id="canonical-d0e37cb85941280905d36aad366a3a83c1799fbd646c25c15055326d7e82c696"></a>

## Next pages — cookie_params.auth_hmac.sec_key.clear_secret_info / e0cc538c5b50 / 6

- [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-2fea82869c57ba9da95ec5787490c7a7afbb73e12573875bdbe35ea141435aa1)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-5a39e95fe94b3946f40c8b2d88f08bf9b197fdbfa8a8919cd8fb47e65047d4fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e9d21f6aa00334b2a33c22e4225091c818ea9d04b5a3acd3be8e5aa381107fa"></a>

## cookie_params.kms_key_hmac — cookie_params.kms_key_hmac / 8aa3192091bc / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- cookie_params.kms_key_hmac

<a id="canonical-f63f63686c2cfb6a0af0a070a8205c443c010052f8326bfa521074123dc3f304"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for kms key hmac.

Upstream description:

Reference to KMS Key Object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-eff038b91bd960fdecae954f4adff204e3becb3812e041624e3c5bdf8ee13322"></a>

## Direct properties — cookie_params.kms_key_hmac / 8aa3192091bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-497494b93eb1ef8f53452e90fa627203475b9c762afc9c528489ebbe73ef8b08"></a>

## Next pages — cookie_params.kms_key_hmac / 8aa3192091bc / 4

- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-70260c1fd00749508caf1cf9b247a750ebab5037979be9fd1621b8ac5867ebbf)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab49612ed0520026cdcb7db1130e9db43170be7a704de6e3a409eac3d4453fc5"></a>

## oidc_auth — oidc_auth / 67fc815ee940 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- oidc_auth

<a id="canonical-46725115245c47fd91edf6523e8287b562b6a48e4b4a88400454bc10b354abef"></a>

Type: `"single"`. Computed.

OIDCAuthType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params_choice": "[\"oidc_auth_params\",\"oidc_well_known_config_url\"]"
}
```

<a id="canonical-68ad8b34d2ee28a3e95836869d8b00ee4262ad7e8c9061d98a1254de0222b2f7"></a>

## Direct properties — oidc_auth / 67fc815ee940 / 3

- [client_secret](data-sources--authentication--reference--group-001.md#canonical-ec16608fed71b84c298a750e54c7e45efc4697994fec05e14f688a5c8c43dc89): complete subsection reference.

- [oidc_auth_params](data-sources--authentication--reference--group-001.md#canonical-69f044829261a1ee626c6cd73631b8b9120d3f4f271fcfb8464b011eae9db786): complete subsection reference.

<a id="canonical-b5cb02d6209ee8fe970df52bfa40605d82fde063dad022cbd7b0dbc7e8373595"></a>

<a id="canonical-3d96b2ff731635fe006897d464c5450b97342b86de863cbdf2531c512ecae8a2"></a>

## oidc_client_id property — oidc_auth / 67fc815ee940 / 4

Type: `"string"`. Computed.

Client ID used while sending the Authorization Request to OIDC server.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-714e183c551c772d61fea1c4a7b1dee920788384b56b8fc98c21d29978c91953"></a>

<a id="canonical-bffa4551457c92724caa5f0d5f45e5453e4668b0088c7888becb5816904f1119"></a>

## oidc_well_known_config_url property — oidc_auth / 67fc815ee940 / 5

Type: `"string"`. Computed.

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Upstream description:

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-92af03a58afb9dd1b95acff431e6632cd3d5bcf5dcbbe0d273766fa01492c73c"></a>

## Next pages — oidc_auth / 67fc815ee940 / 6

- [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-ec16608fed71b84c298a750e54c7e45efc4697994fec05e14f688a5c8c43dc89)
- [oidc_auth.oidc_auth_params](data-sources--authentication--reference--group-001.md#canonical-69f044829261a1ee626c6cd73631b8b9120d3f4f271fcfb8464b011eae9db786)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-ec16608fed71b84c298a750e54c7e45efc4697994fec05e14f688a5c8c43dc89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61968e91bc02fcc36fd21de4b5c4b47e0cf37a0e1790cd45ba5909fb1496c0f8"></a>

## oidc_auth.client_secret — oidc_auth.client_secret / 681347686191 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b)
- oidc_auth.client_secret

<a id="canonical-a77b48f3c2b4c18dfa3d90c61ed1d46e05fe8d5de13c2653d4fb29d81a1fa232"></a>

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

<a id="canonical-f0ca6fe1257c3de3c5a5b23516ef4fdf0f5208dbd764077084da1cc0184c0798"></a>

## Direct properties — oidc_auth.client_secret / 681347686191 / 3

- [blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-c1319a7fd8c8de1b53d3e7232d8d29abe9f42228acd0c66341aed42fe247e1b8): complete subsection reference.

- [clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-614f9aec35af9cc511d5c349574aa59bdd9e28df7d8b1d3391a7cf0da56d1c3f): complete subsection reference.

<a id="canonical-fc6e3665c295e09f6b1a90efbaa36c2cec1de97bbc0fe36d86d187000fab1c32"></a>

## Next pages — oidc_auth.client_secret / 681347686191 / 4

- [oidc_auth.client_secret.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-c1319a7fd8c8de1b53d3e7232d8d29abe9f42228acd0c66341aed42fe247e1b8)
- [oidc_auth.client_secret.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-614f9aec35af9cc511d5c349574aa59bdd9e28df7d8b1d3391a7cf0da56d1c3f)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-c1319a7fd8c8de1b53d3e7232d8d29abe9f42228acd0c66341aed42fe247e1b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-211dd28bc0c44700890421356b7f36a022a7648ce8cc379192c2082feafaab94"></a>

## oidc_auth.client_secret.blindfold_secret_info — oidc_auth.client_secret.blindfold_secret_info / e38d8f2aee14 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b)
- [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-ec16608fed71b84c298a750e54c7e45efc4697994fec05e14f688a5c8c43dc89)
- oidc_auth.client_secret.blindfold_secret_info

<a id="canonical-f4142578ab0560c88fbb71e3d5f0ce0a11e0fe86cc92d0cf7395c307eda4132f"></a>

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

<a id="canonical-08c305feb6680a097b9881cb51755e17f96e9916cd2796544da59586e50ebdf4"></a>

## Direct properties — oidc_auth.client_secret.blindfold_secret_info / e38d8f2aee14 / 3

<a id="canonical-ab23c38784f93f6e037831463065a79339f1ffaae4e1c6dceee04dd5b982f59c"></a>

<a id="canonical-47df1411ce70ffcf13f61d7f16bb115dd94094fde2cf36a00d64f3bfc43d895c"></a>

## decryption_provider property — oidc_auth.client_secret.blindfold_secret_info / e38d8f2aee14 / 4

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

<a id="canonical-bd492b22b53b9436a9b424a2a38ba7e03f9642e60692b5a1ea1fc9e09aaf6e8f"></a>

<a id="canonical-9e11212da79d8bcb645c3e2f587e3402d5879a089d004ad772d3daf23c99a69b"></a>

## location property — oidc_auth.client_secret.blindfold_secret_info / e38d8f2aee14 / 5

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

<a id="canonical-746af60ac73855368b6ced7b6ff526ff88d858f6ef43303cfd2176ed1661b661"></a>

<a id="canonical-52fe738069d03b1d495285bfbd5a0dd22749f9722e7e49ff465c4e0bf5b1af87"></a>

## store_provider property — oidc_auth.client_secret.blindfold_secret_info / e38d8f2aee14 / 6

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

<a id="canonical-55dd7ec9d45847475a3cf0ec4ba191dc800f5f49307fbb08ddb872a07a6e8043"></a>

## Next pages — oidc_auth.client_secret.blindfold_secret_info / e38d8f2aee14 / 7

- [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-ec16608fed71b84c298a750e54c7e45efc4697994fec05e14f688a5c8c43dc89)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-614f9aec35af9cc511d5c349574aa59bdd9e28df7d8b1d3391a7cf0da56d1c3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d123231ca6cee88146d965c192e0dd1ba034ec3ae60086885668c96bfea3fe02"></a>

## oidc_auth.client_secret.clear_secret_info — oidc_auth.client_secret.clear_secret_info / 166b04fd4e44 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b)
- [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-ec16608fed71b84c298a750e54c7e45efc4697994fec05e14f688a5c8c43dc89)
- oidc_auth.client_secret.clear_secret_info

<a id="canonical-ea91951c8a428bbeb9d4b20df09f02d16f7df3881269837b1e00d24324fcb0a0"></a>

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

<a id="canonical-f77d6d96dc09554671d7225e0544ff3fc6b49aaeecfed563e6c2287e4f26f35f"></a>

## Direct properties — oidc_auth.client_secret.clear_secret_info / 166b04fd4e44 / 3

<a id="canonical-f9893a30eb89a61a5f651fdc64d5afb0bda57fa4676e4ddb10921e69179b786b"></a>

<a id="canonical-ff3885b51beae099c746c80b6ab02a91f0149a66a6028fd401d9f4d1c37578e0"></a>

## provider_ref property — oidc_auth.client_secret.clear_secret_info / 166b04fd4e44 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1477563430cbb8d14bbac786e8898de7803dffad7c1e30474df9508c4af6490f"></a>

<a id="canonical-fa5b71713fe497d98aa925fa88e3f67120aea1af803fd9d80cc4b303b7ba15f8"></a>

## url property — oidc_auth.client_secret.clear_secret_info / 166b04fd4e44 / 5

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

<a id="canonical-3eff07d3103ed9639fd22656497cc588027915679bc63f294ba918af7db5be0d"></a>

## Next pages — oidc_auth.client_secret.clear_secret_info / 166b04fd4e44 / 6

- [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-ec16608fed71b84c298a750e54c7e45efc4697994fec05e14f688a5c8c43dc89)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-69f044829261a1ee626c6cd73631b8b9120d3f4f271fcfb8464b011eae9db786"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d9eddc5a25843669b3b27810a8a57bf90b4d4c0dbc989aaeab5f84e17c0eafb"></a>

## oidc_auth.oidc_auth_params — oidc_auth.oidc_auth_params / 990e15d5c887 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b)
- oidc_auth.oidc_auth_params

<a id="canonical-23634d39d1c5e50c541df25052f8bae7427a1e98b8fcf1ba59562a3a920d9e4b"></a>

Type: `"single"`. Computed.

Configuration parameter for oidc auth params.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0ca1677f98fc6a20751455156f801c0e9389a36d0c7bce9f9fa1d82776811191"></a>

## Direct properties — oidc_auth.oidc_auth_params / 990e15d5c887 / 3

<a id="canonical-fbaf232449260b4de1553c8625b648b76a2f64190680dc207462765741a960f7"></a>

<a id="canonical-f56765e586231cd08f8fb259baecd665e07e67349da8aedd7aa9f21775065bdc"></a>

## auth_endpoint_url property — oidc_auth.oidc_auth_params / 990e15d5c887 / 4

Type: `"string"`. Computed.

URL of the authorization server's authorization endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-5052b78dd3055ccff9b7030e72ed59d31c439ff73ee1afafef06f54a49e8da6d"></a>

<a id="canonical-ba395b3c8cdf4d25ad688042a4c15d1720a76f13376df8838b55b5e9bf7be17d"></a>

## end_session_endpoint_url property — oidc_auth.oidc_auth_params / 990e15d5c887 / 5

Type: `"string"`. Computed.

URL of the authorization server's Logout endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-dfe7c4d8ee665270cacbcd73c4b999211b02aadeae57104f96fec5983bfb90fc"></a>

<a id="canonical-c8dc42d04d3f5407bf4834ebb7ade616381035b93eaee37640bc7514cf6aa54b"></a>

## token_endpoint_url property — oidc_auth.oidc_auth_params / 990e15d5c887 / 6

Type: `"string"`. Computed.

URL of the authorization server's Token endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-da23d74ca5d5fdf3ffd8207f0215ba5b711ebfa5331e0415f504eded4653a56f"></a>

## Next pages — oidc_auth.oidc_auth_params / 990e15d5c887 / 7

- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0d50b2bb6675c72d975321b4f58b1dfab6786421bba53020b40f074d29d7715b)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
