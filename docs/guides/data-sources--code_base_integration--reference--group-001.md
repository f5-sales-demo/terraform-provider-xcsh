---
page_title: "xcsh_code_base_integration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration reference."
---

# xcsh_code_base_integration reference

<a id="canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19297ef478c7b1c7d2d84f534312f3899cbcb509064733d2dd4838421af7cbc7"></a>

## Property reference — Property reference / ac6f73b30cf7 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- Property reference

<a id="canonical-fd8ba7b845aef11cc0ef8ddd42e1f03e780aedcad654508bc0463b404b760ce0"></a>

## Direct properties — Property reference / ac6f73b30cf7 / 3

<a id="canonical-d6048cda34c019677863db557257615c08abb06d408ab2810be0dc97305d4393"></a>

<a id="canonical-d7b2f2a3b39c7dbb79bd671023593ddfb9b97135019c4b0ef79e230ff8cddf57"></a>

## annotations property — Property reference / ac6f73b30cf7 / 4

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

- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a): complete subsection reference.

<a id="canonical-cf4270fcec1f31c383312de67391d0963df6a027f30332dacdf677d9c0c6dd4b"></a>

<a id="canonical-67229075cbb8d52d7dd8825a6a9a4698a0dbfb9263a091aa3614cc35e439e464"></a>

## description property — Property reference / ac6f73b30cf7 / 5

Type: `"string"`. Computed.

Description of the CodeBaseIntegration.

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

<a id="canonical-bc4557727c51e6625034ccefdc5b900d3ab9227936bb93d7e8db3036cabf2b40"></a>

<a id="canonical-c60bb65bcbb5e223e4dd7752c5bd5543dfe6f9d3e867bafefea9fb290f57af64"></a>

## id property — Property reference / ac6f73b30cf7 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-654ff693a9645cfdd646f970c8c2e8a7178364cb2af02760190d6bea25329c78"></a>

<a id="canonical-daa33f76b737d1fd94e393a7c32e7a0fe56c4068be3c1dfef589f160d80d9563"></a>

## labels property — Property reference / ac6f73b30cf7 / 7

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

<a id="canonical-70a66c81409c564c393bb7bcda564f77adb5436b886c315211cd0b9e8bfdc9bd"></a>

<a id="canonical-61f42d47b6b95f5bd83fbc3c6f7747854267d7837b5c8bd4c0712ae90bfeea84"></a>

## name property — Property reference / ac6f73b30cf7 / 8

Type: `"string"`. Required.

Name of the CodeBaseIntegration.

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

<a id="canonical-f8090630a7264a7dd9131ca6ab45918c12040bcfcd85035358453b078d298540"></a>

<a id="canonical-a857fbb6c71b310633a6e5b49a98fde2864a31868f3ffc01531c87f329e4c1c2"></a>

## namespace property — Property reference / ac6f73b30cf7 / 9

Type: `"string"`. Required.

Namespace where the CodeBaseIntegration exists.

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

<a id="canonical-f44c63d02ac6933dd30d8bab6369c88c3f376fb371ea5f47b767a93f66b06940"></a>

## All schema paths — Property reference / ac6f73b30cf7 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--code_base_integration--reference--group-001.md#canonical-d6048cda34c019677863db557257615c08abb06d408ab2810be0dc97305d4393) |
| `code_base_integration` | [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-441d0280bdc98543ecdc1bb9fdd77c6b86dc391699d558eab0d8f71ed1f6b0b2) |
| `code_base_integration.azure_repos` | [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-2271d76404bd65b50674f510c03c772cface0f51c615f669e3a31ea80d993645) |
| `code_base_integration.azure_repos.access_token` | [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-ebdd46a0ab5d6e245edb87e17616d3d660e4323c3ebf4cad47eb2678bf886e7a) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info` | [code_base_integration.azure_repos.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-fc792184f5aa7fe056d0b1d6c00ad51054cf4505e5cf9012fa097f8f1a9c01ee) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-cf367c5d1f7b07a3eb5a1b906de51c7cde2b39dcfc1334c1140e20dab3d29d8d) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.location` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-464c70eed4d3f026242ac888f1251f7c49c4f1dc99ce1dffb954bdff9ed97779) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-56be5c4c63e240b86a49987c70b024f3bdd3012e96b556d96d52473d44152c56) |
| `code_base_integration.azure_repos.access_token.clear_secret_info` | [code_base_integration.azure_repos.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-3197e6875337f58d15e62f257af9b1b79df9530a3bb1f6292c7179fe190e17f1) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref` | [code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-114aaae80861d433eb376c952c47307d9555a0e41fa9820baea6a1b2c75bb9a5) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.url` | [code_base_integration.azure_repos.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-ea21e85758e871953ba3db0f50773d50093314c0559eef37dd7033faced8be65) |
| `code_base_integration.bitbucket` | [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-25e5e3240cc93e3baf061905ff23d3ec9b4f26f520b438f1ec0df7b5837e648f) |
| `code_base_integration.bitbucket.passwd` | [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-5388ce00e011e59281a0812d54943bdcfde44855b9fea837bb0eb1b56a9ad698) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info` | [code_base_integration.bitbucket.passwd.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-8a0172840b8ddeccfd69166c5939edb2c6020b25434fbf8b445595c14efd6adc) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-4c8c46c80f48db121db652e83c7ff32b786abbd90e952bb7ce439f6870be995e) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-f2941ececd4c3f81f8eb2ac26afb0dd402a4d92d2278c5c16476ec65fb89f934) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-815b7b6e4264da8fcd71f2c1328fd98621147ce6670d743312f656c86b19c7ff) |
| `code_base_integration.bitbucket.passwd.clear_secret_info` | [code_base_integration.bitbucket.passwd.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-4eaa586b12e7dad7919ab5ebdf65302debe3a86f8a44bd0e866b5cee0792d058) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-cc64a75a08ab735cfb4fcc8e21ea495ed08b16acc3b9c44e7b4eec73ea4bb6c8) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.url` | [code_base_integration.bitbucket.passwd.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-ebcb29828cda01f349b37803ef0180410a3b685eb24c25b04acdd7ff66d7e776) |
| `code_base_integration.bitbucket.username` | [code_base_integration.bitbucket.username](data-sources--code_base_integration--reference--group-001.md#canonical-2691ae77aa7534ddeb855f2775dcefb949fd01900ee3be88b153429d89f2b06d) |
| `code_base_integration.bitbucket_server` | [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-671f777d9eae0dd5989c4ca474cb3c7c53d568d5e677c3624546e2df53457d6e) |
| `code_base_integration.bitbucket_server.passwd` | [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-a73158ded104e94d281f0d8e1104de4a44a7a8e2d5709ab45c4182c853aad368) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-24cff6870d17ad4c985164a7bc84091924f50b581b9e60dd8ffdffebad347655) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-b779ca4a6e846531c6f7f67a748dba0208e7cf5a36cd4dfaeb2fb02d428191d8) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-7ec1ea740c07d61c8fd07b015721f0dd2e0b5676ba792bda2c953f7bbd7008f2) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-4a5dade5087df2ec800e8a5c95cfc876227f5c360be0b360d107547a100db9dc) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info` | [code_base_integration.bitbucket_server.passwd.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-41330652e75e3e2a3274b2d10f53130fb86e6aece2265d00ca91b4711551fee1) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-b7f74149bcf102b9388bd5b386dab220c6df4d4fadde33d5b0b85c05dd4ed053) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.url` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-756316908401132985145c0faf6688cf9529dd383015e62d2409cc9510171963) |
| `code_base_integration.bitbucket_server.url` | [code_base_integration.bitbucket_server.url](data-sources--code_base_integration--reference--group-001.md#canonical-f7ec844f6874e06c638a12bac0d5de78c0bd618b56736facedf2792f83f2e4f6) |
| `code_base_integration.bitbucket_server.username` | [code_base_integration.bitbucket_server.username](data-sources--code_base_integration--reference--group-001.md#canonical-3fffb1506b9001a410681851fc2afa2211b0263c05e4b26ec94fe337c2c10f28) |
| `code_base_integration.bitbucket_server.verify_ssl` | [code_base_integration.bitbucket_server.verify_ssl](data-sources--code_base_integration--reference--group-001.md#canonical-a5cd73cfb3a17694e1726084dc739260cb98aba3e67ec586632c3cbb586f252d) |
| `code_base_integration.github` | [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-cf8ac5c057ac425affd11fa87e987b11119957e384f7c921f606d5a65940a776) |
| `code_base_integration.github.access_token` | [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-19133a832618c762e0b4766b6bca605df36b02ec2043908220bd3641b6d3d965) |
| `code_base_integration.github.access_token.blindfold_secret_info` | [code_base_integration.github.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-d181e91b2fc193e2b93bc12439294a22ca48e9757f98c4c97514e0ea1f5d1317) |
| `code_base_integration.github.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-909bb35e777321db75704ab228c8e8c501e1e165fae7e597f74e7f6438d8eb99) |
| `code_base_integration.github.access_token.blindfold_secret_info.location` | [code_base_integration.github.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-d4e6257055b5cf8a5c83940f4798b0533ada17fbf97b4ffcfa1817a03284f287) |
| `code_base_integration.github.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-de59d9089c1bc6361e987a6132968a980a72efc67f4e3e3e82086ad7d6246037) |
| `code_base_integration.github.access_token.clear_secret_info` | [code_base_integration.github.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-644bc3ef1780b9553542e301953e1f5dfbd632bfe9842eaa7db5978581ce3684) |
| `code_base_integration.github.access_token.clear_secret_info.provider_ref` | [code_base_integration.github.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-7f5727188288a3d8c994e1db2d0fccb8b363728e08bf4cf2586e2eaebfa17f56) |
| `code_base_integration.github.access_token.clear_secret_info.url` | [code_base_integration.github.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-b2638d2980142866dda1a47598bb34a2f8a5ac74c93f52402f1f2adf58d90822) |
| `code_base_integration.github.username` | [code_base_integration.github.username](data-sources--code_base_integration--reference--group-001.md#canonical-d1079c593b1f756c1b9686df1fbb35be6b8559c693766a54156b167d43bf7770) |
| `code_base_integration.github.verify_ssl` | [code_base_integration.github.verify_ssl](data-sources--code_base_integration--reference--group-001.md#canonical-3779a73589498045ba20698c53997f2dc3c148acc5c30d6151958b267b758e83) |
| `code_base_integration.github_enterprise` | [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-3513136be3b0b18d7ff67a66b5119b2611c3dbdf498633010b48ddaea8a446e6) |
| `code_base_integration.github_enterprise.access_token` | [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-036481ddd2ab01edcaf15d8087a190cc02202fbf3d8b18fcbc956262a2f1e66b) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-9c3332da1008c88762b29befe9b5d45d081a535e1018edee61b682e0750b5325) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-770857d8d635d5e05c203b0b2cf00705cfd07d061097046f09423b2dfaaa55f2) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-63c53d354611ba15a4a8203301e564f848a33ea192d0cb44c02e2200ef9f35b5) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-e86fc5eb9e2f69f994093dfd0694212d499cf6c65cc2eba767b8141ea5b2735d) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info` | [code_base_integration.github_enterprise.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-6e7219779df286164a65eb60e31436f5dd2d26d12d63e46f702a38e884782152) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-a28e94c74ac8d655e1710253e69a0a8724470f7658bf8845f1790a6e99b29561) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.url` | [code_base_integration.github_enterprise.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-d04de4cf9216d5f3996956847fcd4bcef2c99564b9f2eda573147b890f50f52b) |
| `code_base_integration.github_enterprise.hostname` | [code_base_integration.github_enterprise.hostname](data-sources--code_base_integration--reference--group-001.md#canonical-1cf7cd8bbb504f4153b71add6cb8aa52affdd2d163706745244930a042fff02d) |
| `code_base_integration.github_enterprise.username` | [code_base_integration.github_enterprise.username](data-sources--code_base_integration--reference--group-001.md#canonical-2414aca3debcc30f1d741aef72d5267e840fd46353459c69fcd2b1e616b6ed54) |
| `code_base_integration.gitlab` | [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-b706bf9d0fdaf100196de89cfaf07baeb63988c7b873311b8285984594f2d2cd) |
| `code_base_integration.gitlab.access_token` | [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-8d302ddc5570f73990fbe407e090bda894a9008ee6e96d82b9d4f67aa17d365b) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info` | [code_base_integration.gitlab.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-c092753bd8915a86b100421da39a01233710168d193a0a0e08d6fc92d83c41b6) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-7fde553a667d93fe30319a8115e63cb90341dc50a4c91b567d30c1607182eb3e) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-11aaebce12530e8de2df6fc05d93c213f5db6bb023a9070123ef4e64be33fb99) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-6344f4fabbb2a26f5fd2d7f58e16ec16af79f46829c0a0ad9e98880909371f14) |
| `code_base_integration.gitlab.access_token.clear_secret_info` | [code_base_integration.gitlab.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-59f0671a9e8894e0cb71effb6f82e324126ce7e8dea836a24e215cb74facf189) |
| `code_base_integration.gitlab.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-63b49b6cc852f148cec80caa5ecda3f0a65b7688db1bf6b6da2034e36d7d7b62) |
| `code_base_integration.gitlab.access_token.clear_secret_info.url` | [code_base_integration.gitlab.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-360647f3e286390a500fca52651a15c5ae683b1b9ccd4324705ecef5f453bb1a) |
| `code_base_integration.gitlab_enterprise` | [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2ff7ce8141b849a71d220e5171e7382c617ceb0b96e2c4b7a91d996e441c467a) |
| `code_base_integration.gitlab_enterprise.access_token` | [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-bbd8377b227c0c9f0cc64a3464ac21d6580b48513fa3d113a524df45e01703fa) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-78834be3849252f6ca318e52d2d5b7340ba8f2c4ffdeb51a5a9584cb3e19aa44) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-17fdf860933f8a11d5871b03ddf1a42ca74ee757349eb95d90c1c1b7f71f5637) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-0002b42b4b1575925aabfe5aab15f7a33d3521299a64ddefe1cfeb8407c09494) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-771b75379dc0c5070b59206ead70bd0f1546ae54b13f58967a74806a42e831c7) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-62589d6b6e7d264d4522ab73999d9ff51aea537b92899196ad6754ab46c5f8e9) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-dc8e75c82e3c0da3335e253d2d03848fdf0117f266aa06891d148ae415a2b60f) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-a23e5c61be52b4a164bdc14b9ffdfcd57b705191cc5e6259a2b1cc323bf943b7) |
| `code_base_integration.gitlab_enterprise.url` | [code_base_integration.gitlab_enterprise.url](data-sources--code_base_integration--reference--group-001.md#canonical-2ec053cbadbbe011ce1db8d54e7be8d82d7d2d800198c590edf01b3b4c204e75) |
| `description` | [description](data-sources--code_base_integration--reference--group-001.md#canonical-cf4270fcec1f31c383312de67391d0963df6a027f30332dacdf677d9c0c6dd4b) |
| `id` | [id](data-sources--code_base_integration--reference--group-001.md#canonical-bc4557727c51e6625034ccefdc5b900d3ab9227936bb93d7e8db3036cabf2b40) |
| `labels` | [labels](data-sources--code_base_integration--reference--group-001.md#canonical-654ff693a9645cfdd646f970c8c2e8a7178364cb2af02760190d6bea25329c78) |
| `name` | [name](data-sources--code_base_integration--reference--group-001.md#canonical-70a66c81409c564c393bb7bcda564f77adb5436b886c315211cd0b9e8bfdc9bd) |
| `namespace` | [namespace](data-sources--code_base_integration--reference--group-001.md#canonical-f8090630a7264a7dd9131ca6ab45918c12040bcfcd85035358453b078d298540) |

<a id="canonical-ac25e6306795eede5960292847bbc7511f3281d2cbd58f77abefca42408d350f"></a>

## Next pages — Property reference / ac6f73b30cf7 / 11

- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e27f74e2b6c6763909739abb97a3dbb500383ea63c45d71d3806ef627e6a3c1e"></a>

## code_base_integration — code_base_integration / b2b7680527e5 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- code_base_integration

<a id="canonical-441d0280bdc98543ecdc1bb9fdd77c6b86dc391699d558eab0d8f71ed1f6b0b2"></a>

Type: `"single"`. Computed.

Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection
details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"azure_repos\",\"bitbucket\",\"bitbucket_server\",\"github\",\"github_enterprise\",\"gitlab\",\"gitlab_enterprise\"]"
}
```

<a id="canonical-9255796ea5ff1fb421e6f0ab0554fa8678e477dcdddbfd387ddd4aa2fe1e52ba"></a>

## Direct properties — code_base_integration / b2b7680527e5 / 3

- [azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-b61d2c542cb4cbf009971d2ad39e7939767ed5064cd24c1566fb7faaec31fea2): complete subsection reference.

- [bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-04a1875e291dfc69a4c1e11cb18be37186bf830af8ced7e493b15dfced49b71b): complete subsection reference.

- [bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-afc28030a1c2560702fb0be8c64583cb272e655832a15deaf2f00cf588b884a8): complete subsection reference.

- [github](data-sources--code_base_integration--reference--group-001.md#canonical-463045aa8e9b205f822124632986b8adaee30fa5a60e062b0778ea3bf890506d): complete subsection reference.

- [github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-9e6712137098c88e9972ddfe960990503f708ce52facf7ba62b9fe68e3d1ddbe): complete subsection reference.

- [gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-c153e4e3b91fb5dc774723280a6e60c47fa195cccb0b0ead5c42c85315b3f6ef): complete subsection reference.

- [gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-bf448306c188c0c01504013875c4737a4140687b4a73c8385e9e265cae53c79b): complete subsection reference.

<a id="canonical-5634b0d46b52d9f09fe44bef5e8378e150e84f7435ec09ab1412300c0e440b68"></a>

## Next pages — code_base_integration / b2b7680527e5 / 4

- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-b61d2c542cb4cbf009971d2ad39e7939767ed5064cd24c1566fb7faaec31fea2)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-04a1875e291dfc69a4c1e11cb18be37186bf830af8ced7e493b15dfced49b71b)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-afc28030a1c2560702fb0be8c64583cb272e655832a15deaf2f00cf588b884a8)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-463045aa8e9b205f822124632986b8adaee30fa5a60e062b0778ea3bf890506d)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-9e6712137098c88e9972ddfe960990503f708ce52facf7ba62b9fe68e3d1ddbe)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-c153e4e3b91fb5dc774723280a6e60c47fa195cccb0b0ead5c42c85315b3f6ef)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-bf448306c188c0c01504013875c4737a4140687b4a73c8385e9e265cae53c79b)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-b61d2c542cb4cbf009971d2ad39e7939767ed5064cd24c1566fb7faaec31fea2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ed887499eb8e44e92140a140058b0d5ec770769f47ef35b6bf9cc42ec7fdfb6"></a>

## code_base_integration.azure_repos — code_base_integration.azure_repos / 05e27cdbf326 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- code_base_integration.azure_repos

<a id="canonical-2271d76404bd65b50674f510c03c772cface0f51c615f669e3a31ea80d993645"></a>

Type: `"single"`. Computed.

Configuration parameter for azure repos.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fc7e4abc74587234c069f17264424997ad5f4b4dcac0c99863c0e240b2552183"></a>

## Direct properties — code_base_integration.azure_repos / 05e27cdbf326 / 3

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-ac2362351a7618c4581571be20b030e55950aa05b78c6bf41d414f9159bd68c0): complete subsection reference.

<a id="canonical-78be975b6f9f4519b786ae346accc8e29a10728517f6b915ef8282583b3234ee"></a>

## Next pages — code_base_integration.azure_repos / 05e27cdbf326 / 4

- [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-ac2362351a7618c4581571be20b030e55950aa05b78c6bf41d414f9159bd68c0)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-ac2362351a7618c4581571be20b030e55950aa05b78c6bf41d414f9159bd68c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b9a7059aeaed20b10a6fa9aa3a7f6dec4c22decac962e1e2474656ae86abe8f"></a>

## code_base_integration.azure_repos.access_token — code_base_integration.azure_repos.access_token / b1223062565c / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-b61d2c542cb4cbf009971d2ad39e7939767ed5064cd24c1566fb7faaec31fea2)
- code_base_integration.azure_repos.access_token

<a id="canonical-ebdd46a0ab5d6e245edb87e17616d3d660e4323c3ebf4cad47eb2678bf886e7a"></a>

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

<a id="canonical-60d5a851c96bb298a9aca683cea573937d553065d9f028d461a9b12d6506c006"></a>

## Direct properties — code_base_integration.azure_repos.access_token / b1223062565c / 3

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-88b01bf9d456bab209e95e186c3c73440f45bedef36658f0559453c6da3743bf): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-90ccd0b554d19a3ecd68fbc69fc89f639cf49091096f164e68537ad0dc30b4d4): complete subsection reference.

<a id="canonical-4ac45a2ff21487e28314c8a95bc3ae14e3761f1391d783304461f6bbf7cc814c"></a>

## Next pages — code_base_integration.azure_repos.access_token / b1223062565c / 4

- [code_base_integration.azure_repos.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-88b01bf9d456bab209e95e186c3c73440f45bedef36658f0559453c6da3743bf)
- [code_base_integration.azure_repos.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-90ccd0b554d19a3ecd68fbc69fc89f639cf49091096f164e68537ad0dc30b4d4)
- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-b61d2c542cb4cbf009971d2ad39e7939767ed5064cd24c1566fb7faaec31fea2)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-88b01bf9d456bab209e95e186c3c73440f45bedef36658f0559453c6da3743bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41aa4dd6438531e068f91061068b352d89ddf37124b05294a98379dff35ac2db"></a>

## code_base_integration.azure_repos.access_token.blindfold_secret_info — code_base_integration.azure_repos.access_token.blindfold_secret_info / 96ae88fe1833 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-b61d2c542cb4cbf009971d2ad39e7939767ed5064cd24c1566fb7faaec31fea2)
- [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-ac2362351a7618c4581571be20b030e55950aa05b78c6bf41d414f9159bd68c0)
- code_base_integration.azure_repos.access_token.blindfold_secret_info

<a id="canonical-fc792184f5aa7fe056d0b1d6c00ad51054cf4505e5cf9012fa097f8f1a9c01ee"></a>

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

<a id="canonical-701963d1483d217b6c803d22c359551dc3f2e9a319087b04075315c84f446b08"></a>

## Direct properties — code_base_integration.azure_repos.access_token.blindfold_secret_info / 96ae88fe1833 / 3

<a id="canonical-cf367c5d1f7b07a3eb5a1b906de51c7cde2b39dcfc1334c1140e20dab3d29d8d"></a>

<a id="canonical-d4e2003d6198601894ad493072aa44cfd4923b76edbf7c20e5db336ddfc687b5"></a>

## decryption_provider property — code_base_integration.azure_repos.access_token.blindfold_secret_info / 96ae88fe1833 / 4

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

<a id="canonical-464c70eed4d3f026242ac888f1251f7c49c4f1dc99ce1dffb954bdff9ed97779"></a>

<a id="canonical-5f95444c1a50d4b267a2561ec56653df4512009c047793519c989f379d058cc9"></a>

## location property — code_base_integration.azure_repos.access_token.blindfold_secret_info / 96ae88fe1833 / 5

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

<a id="canonical-56be5c4c63e240b86a49987c70b024f3bdd3012e96b556d96d52473d44152c56"></a>

<a id="canonical-2004c5ef5ab278ad5010b7313add511b5854bc8d99f6dd5358f37eb276593c8d"></a>

## store_provider property — code_base_integration.azure_repos.access_token.blindfold_secret_info / 96ae88fe1833 / 6

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

<a id="canonical-356c6655977feb6ba7b9e8858d5d531a99fb7a07f369d4cccd4f691238b402d2"></a>

## Next pages — code_base_integration.azure_repos.access_token.blindfold_secret_info / 96ae88fe1833 / 7

- [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-ac2362351a7618c4581571be20b030e55950aa05b78c6bf41d414f9159bd68c0)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-90ccd0b554d19a3ecd68fbc69fc89f639cf49091096f164e68537ad0dc30b4d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5bbd458e471b5515b3b66f2dfe1887d5a2025db77e0f40aba334893cf42ff31"></a>

## code_base_integration.azure_repos.access_token.clear_secret_info — code_base_integration.azure_repos.access_token.clear_secret_info / 77addacdadc7 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-b61d2c542cb4cbf009971d2ad39e7939767ed5064cd24c1566fb7faaec31fea2)
- [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-ac2362351a7618c4581571be20b030e55950aa05b78c6bf41d414f9159bd68c0)
- code_base_integration.azure_repos.access_token.clear_secret_info

<a id="canonical-3197e6875337f58d15e62f257af9b1b79df9530a3bb1f6292c7179fe190e17f1"></a>

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

<a id="canonical-ca201ba6dcf3e0b7db511929065697deee6bf3506cae4e4f723a443182c8b839"></a>

## Direct properties — code_base_integration.azure_repos.access_token.clear_secret_info / 77addacdadc7 / 3

<a id="canonical-114aaae80861d433eb376c952c47307d9555a0e41fa9820baea6a1b2c75bb9a5"></a>

<a id="canonical-95f516ebd608b7a2ebb9ae4f89b7f819b93c4edd998b1c49a35f60f06c2cb6b5"></a>

## provider_ref property — code_base_integration.azure_repos.access_token.clear_secret_info / 77addacdadc7 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ea21e85758e871953ba3db0f50773d50093314c0559eef37dd7033faced8be65"></a>

<a id="canonical-c6eb85e2a563b4ec55d48796137c780a33dda664f533e422ca761fd496ffee79"></a>

## url property — code_base_integration.azure_repos.access_token.clear_secret_info / 77addacdadc7 / 5

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

<a id="canonical-ba8f7886d2cf55817dc48ee7e9243017edfd03815d38915b230141f40d386821"></a>

## Next pages — code_base_integration.azure_repos.access_token.clear_secret_info / 77addacdadc7 / 6

- [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-ac2362351a7618c4581571be20b030e55950aa05b78c6bf41d414f9159bd68c0)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-04a1875e291dfc69a4c1e11cb18be37186bf830af8ced7e493b15dfced49b71b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16c18699055f821ec13e102ef88ac8a2cab5c39efe08b9f16d2853cfdb37bac6"></a>

## code_base_integration.bitbucket — code_base_integration.bitbucket / 6db4b942d6c8 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- code_base_integration.bitbucket

<a id="canonical-25e5e3240cc93e3baf061905ff23d3ec9b4f26f520b438f1ec0df7b5837e648f"></a>

Type: `"single"`. Computed.

BitBucket Cloud Integration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-300ca1ffece973602c6acb79e65e439f583f0a095cc237f9376623c858c5622a"></a>

## Direct properties — code_base_integration.bitbucket / 6db4b942d6c8 / 3

- [passwd](data-sources--code_base_integration--reference--group-001.md#canonical-3c5ea7eaeb15382bc58a1689bf96dd0e58010dc0119c91ff60918034a47b08fc): complete subsection reference.

<a id="canonical-2691ae77aa7534ddeb855f2775dcefb949fd01900ee3be88b153429d89f2b06d"></a>

<a id="canonical-3dfac1bb20d462efddf1a8e3b4cdfe42d9273d2387d4377df8a34cc7a29aa0b4"></a>

## username property — code_base_integration.bitbucket / 6db4b942d6c8 / 4

Type: `"string"`. Computed.

BitBucket Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

<a id="canonical-9780324fd29d034d8ac4e2143f6c253836c6aceaa8618fdb7028bab71fb6b8c0"></a>

## Next pages — code_base_integration.bitbucket / 6db4b942d6c8 / 5

- [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-3c5ea7eaeb15382bc58a1689bf96dd0e58010dc0119c91ff60918034a47b08fc)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-3c5ea7eaeb15382bc58a1689bf96dd0e58010dc0119c91ff60918034a47b08fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76d460591a12a109bcbb65e093380ff42fa3c52a7385b654b0a9cd902e6bbb28"></a>

## code_base_integration.bitbucket.passwd — code_base_integration.bitbucket.passwd / ab28d26482eb / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-04a1875e291dfc69a4c1e11cb18be37186bf830af8ced7e493b15dfced49b71b)
- code_base_integration.bitbucket.passwd

<a id="canonical-5388ce00e011e59281a0812d54943bdcfde44855b9fea837bb0eb1b56a9ad698"></a>

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

<a id="canonical-1defad6865f9aa9d9c5c51f36543c8912174c26393d0a0222274c2b8f1742d24"></a>

## Direct properties — code_base_integration.bitbucket.passwd / ab28d26482eb / 3

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-86fa7e082d9500938680be5b25ed1db656459912ff43c5d41ff95ecffd3d1efb): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-dc523bb39dbd6c155a4bb68c4720a78912891a0bc4116893e3c7fe4bb80af0af): complete subsection reference.

<a id="canonical-5e879f1602a63f1d4a5758f81c22134175679a07b98dacc60860b38fd9bed60e"></a>

## Next pages — code_base_integration.bitbucket.passwd / ab28d26482eb / 4

- [code_base_integration.bitbucket.passwd.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-86fa7e082d9500938680be5b25ed1db656459912ff43c5d41ff95ecffd3d1efb)
- [code_base_integration.bitbucket.passwd.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-dc523bb39dbd6c155a4bb68c4720a78912891a0bc4116893e3c7fe4bb80af0af)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-04a1875e291dfc69a4c1e11cb18be37186bf830af8ced7e493b15dfced49b71b)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-86fa7e082d9500938680be5b25ed1db656459912ff43c5d41ff95ecffd3d1efb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31b0802dbcb2bd33f1e08f75399f04f84ed6b370ab465e1df9df4be747ead113"></a>

## code_base_integration.bitbucket.passwd.blindfold_secret_info — code_base_integration.bitbucket.passwd.blindfold_secret_info / 4a4890954532 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-04a1875e291dfc69a4c1e11cb18be37186bf830af8ced7e493b15dfced49b71b)
- [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-3c5ea7eaeb15382bc58a1689bf96dd0e58010dc0119c91ff60918034a47b08fc)
- code_base_integration.bitbucket.passwd.blindfold_secret_info

<a id="canonical-8a0172840b8ddeccfd69166c5939edb2c6020b25434fbf8b445595c14efd6adc"></a>

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

<a id="canonical-8734fbadbd86bacfffef51c6a7e14317efebf07e03b2411a481e4cc1499dc379"></a>

## Direct properties — code_base_integration.bitbucket.passwd.blindfold_secret_info / 4a4890954532 / 3

<a id="canonical-4c8c46c80f48db121db652e83c7ff32b786abbd90e952bb7ce439f6870be995e"></a>

<a id="canonical-b88e656acddd61d8da91a76730ec1169a1bdad8ec9b6c25e06c530d84fdcbdb5"></a>

## decryption_provider property — code_base_integration.bitbucket.passwd.blindfold_secret_info / 4a4890954532 / 4

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

<a id="canonical-f2941ececd4c3f81f8eb2ac26afb0dd402a4d92d2278c5c16476ec65fb89f934"></a>

<a id="canonical-6daec09aab06ffe6d848633a28ac815ed39869d16dcb44427ddd6fee347050f5"></a>

## location property — code_base_integration.bitbucket.passwd.blindfold_secret_info / 4a4890954532 / 5

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

<a id="canonical-815b7b6e4264da8fcd71f2c1328fd98621147ce6670d743312f656c86b19c7ff"></a>

<a id="canonical-e2a79a12bb2efd9389d2e44158b89687f567d5c84973b225b9c4c68460ef058d"></a>

## store_provider property — code_base_integration.bitbucket.passwd.blindfold_secret_info / 4a4890954532 / 6

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

<a id="canonical-a578425150ad66dbe559ab5795aaa85b01030a94c203f21f1f7824acfe86e786"></a>

## Next pages — code_base_integration.bitbucket.passwd.blindfold_secret_info / 4a4890954532 / 7

- [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-3c5ea7eaeb15382bc58a1689bf96dd0e58010dc0119c91ff60918034a47b08fc)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-dc523bb39dbd6c155a4bb68c4720a78912891a0bc4116893e3c7fe4bb80af0af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed1b8daae9c11ae4ab7199f1a50fe8db8c8e72a377da8fe1fd4e7f760fd1333a"></a>

## code_base_integration.bitbucket.passwd.clear_secret_info — code_base_integration.bitbucket.passwd.clear_secret_info / 1b162904c35a / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-04a1875e291dfc69a4c1e11cb18be37186bf830af8ced7e493b15dfced49b71b)
- [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-3c5ea7eaeb15382bc58a1689bf96dd0e58010dc0119c91ff60918034a47b08fc)
- code_base_integration.bitbucket.passwd.clear_secret_info

<a id="canonical-4eaa586b12e7dad7919ab5ebdf65302debe3a86f8a44bd0e866b5cee0792d058"></a>

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

<a id="canonical-6f74ce15fa69388a6bea43c5ebecfb8c842171f93999bd00dd4c1e59413b8dd6"></a>

## Direct properties — code_base_integration.bitbucket.passwd.clear_secret_info / 1b162904c35a / 3

<a id="canonical-cc64a75a08ab735cfb4fcc8e21ea495ed08b16acc3b9c44e7b4eec73ea4bb6c8"></a>

<a id="canonical-e261d71d7eda2f2f645d6db31e3b58b22d29590f001ee409912fcaf37b4a639c"></a>

## provider_ref property — code_base_integration.bitbucket.passwd.clear_secret_info / 1b162904c35a / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ebcb29828cda01f349b37803ef0180410a3b685eb24c25b04acdd7ff66d7e776"></a>

<a id="canonical-d925b1917ed96b9ded84a8715cc665f9a87c092156d453d8ffb6365858d610c5"></a>

## url property — code_base_integration.bitbucket.passwd.clear_secret_info / 1b162904c35a / 5

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

<a id="canonical-91f2a0f476533499fcfc8bfc9188090b7c23651f562625533c93df424cb4f86c"></a>

## Next pages — code_base_integration.bitbucket.passwd.clear_secret_info / 1b162904c35a / 6

- [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-3c5ea7eaeb15382bc58a1689bf96dd0e58010dc0119c91ff60918034a47b08fc)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-afc28030a1c2560702fb0be8c64583cb272e655832a15deaf2f00cf588b884a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce6f06b8b8522b695782689d027014d6b514b805622d16ee621f3c063abc088d"></a>

## code_base_integration.bitbucket_server — code_base_integration.bitbucket_server / c0d6be6788ba / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- code_base_integration.bitbucket_server

<a id="canonical-671f777d9eae0dd5989c4ca474cb3c7c53d568d5e677c3624546e2df53457d6e"></a>

Type: `"single"`. Computed.

Configuration parameter for bitbucket server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-55bb3e0acb82cd951fb520b01efc64912c810b4926d0d9a287702db56bc11062"></a>

## Direct properties — code_base_integration.bitbucket_server / c0d6be6788ba / 3

- [passwd](data-sources--code_base_integration--reference--group-001.md#canonical-aadee12c6732940841907585416cc6e17ce3c1ae77ca04462e950f82bf8337ad): complete subsection reference.

<a id="canonical-f7ec844f6874e06c638a12bac0d5de78c0bd618b56736facedf2792f83f2e4f6"></a>

<a id="canonical-003a0ae85564c957002d7672595be1c212f52b7599ab7198a66670519770982a"></a>

## url property — code_base_integration.bitbucket_server / c0d6be6788ba / 4

Type: `"string"`. Computed.

BitBucket Server URL. URL or URI reference

Upstream description:

URL or URI reference

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
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
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3fffb1506b9001a410681851fc2afa2211b0263c05e4b26ec94fe337c2c10f28"></a>

<a id="canonical-bec9dabb9a7478a433c6c1698f63afccbe51ed4c2fffbffb3b454a97eb7a2638"></a>

## username property — code_base_integration.bitbucket_server / c0d6be6788ba / 5

Type: `"string"`. Computed.

BitBucket Server Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

<a id="canonical-a5cd73cfb3a17694e1726084dc739260cb98aba3e67ec586632c3cbb586f252d"></a>

<a id="canonical-e3ccee60f3d45ad66760b28c94ae02e1af628a8f762b82bcfe59d10c89518538"></a>

## verify_ssl property — code_base_integration.bitbucket_server / c0d6be6788ba / 6

Type: `"bool"`. Computed.

Verify SSL. Configuration parameter for verify ssl

Upstream description:

Configuration parameter for verify ssl

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b09468fcf2c448aed9f238b2832d2d5aae316693caf8d1702f06ddc03199086c"></a>

## Next pages — code_base_integration.bitbucket_server / c0d6be6788ba / 7

- [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-aadee12c6732940841907585416cc6e17ce3c1ae77ca04462e950f82bf8337ad)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-aadee12c6732940841907585416cc6e17ce3c1ae77ca04462e950f82bf8337ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64a9a81ff811f070b5459410c5a7edb4ee7498deeacf9cd12e40fbf8213a5652"></a>

## code_base_integration.bitbucket_server.passwd — code_base_integration.bitbucket_server.passwd / bff5ab26b99f / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-afc28030a1c2560702fb0be8c64583cb272e655832a15deaf2f00cf588b884a8)
- code_base_integration.bitbucket_server.passwd

<a id="canonical-a73158ded104e94d281f0d8e1104de4a44a7a8e2d5709ab45c4182c853aad368"></a>

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

<a id="canonical-1af0f7475e149ff7b1c279d85f5050d35c866b2029b960278b5fc839ff6f2009"></a>

## Direct properties — code_base_integration.bitbucket_server.passwd / bff5ab26b99f / 3

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-fe142c8fd2b7c6336920192bbe61325bae85a16cc06444533b1125d40cf8c19c): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-4571247bd3bb7c5d5f13681342f8103da50e1fada084e51b290fa314210b73d2): complete subsection reference.

<a id="canonical-e781f1933abac7295397e1b1e67552b265b5b2ff72f003987e5fc7e43dbf45cf"></a>

## Next pages — code_base_integration.bitbucket_server.passwd / bff5ab26b99f / 4

- [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-fe142c8fd2b7c6336920192bbe61325bae85a16cc06444533b1125d40cf8c19c)
- [code_base_integration.bitbucket_server.passwd.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-4571247bd3bb7c5d5f13681342f8103da50e1fada084e51b290fa314210b73d2)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-afc28030a1c2560702fb0be8c64583cb272e655832a15deaf2f00cf588b884a8)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-fe142c8fd2b7c6336920192bbe61325bae85a16cc06444533b1125d40cf8c19c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-380976b2e386fd8b44d4305df682372805c3c51d1a334063352f0a4adb457a14"></a>

## code_base_integration.bitbucket_server.passwd.blindfold_secret_info — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / aca02e495b1a / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-afc28030a1c2560702fb0be8c64583cb272e655832a15deaf2f00cf588b884a8)
- [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-aadee12c6732940841907585416cc6e17ce3c1ae77ca04462e950f82bf8337ad)
- code_base_integration.bitbucket_server.passwd.blindfold_secret_info

<a id="canonical-24cff6870d17ad4c985164a7bc84091924f50b581b9e60dd8ffdffebad347655"></a>

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

<a id="canonical-553877ebd9fd96f37793586235dd5bdff20b9cd335bbdafd1dc105e177df876e"></a>

## Direct properties — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / aca02e495b1a / 3

<a id="canonical-b779ca4a6e846531c6f7f67a748dba0208e7cf5a36cd4dfaeb2fb02d428191d8"></a>

<a id="canonical-1f158c38efa753c37b8d7898224f2d0dab36cab241064a699d26133a2e0388ac"></a>

## decryption_provider property — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / aca02e495b1a / 4

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

<a id="canonical-7ec1ea740c07d61c8fd07b015721f0dd2e0b5676ba792bda2c953f7bbd7008f2"></a>

<a id="canonical-8412962d86695d6c6aa7527cf719b014fbfc4acd7028155e26afe786d095b240"></a>

## location property — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / aca02e495b1a / 5

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

<a id="canonical-4a5dade5087df2ec800e8a5c95cfc876227f5c360be0b360d107547a100db9dc"></a>

<a id="canonical-153e97b22b397ecc5816c65d0b01af1213b1c92b4152a813202070681fadd00c"></a>

## store_provider property — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / aca02e495b1a / 6

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

<a id="canonical-047aacf5063c75c493368a193b711266cc38ac8243a0d61cff3de4a3a14013cd"></a>

## Next pages — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / aca02e495b1a / 7

- [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-aadee12c6732940841907585416cc6e17ce3c1ae77ca04462e950f82bf8337ad)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-4571247bd3bb7c5d5f13681342f8103da50e1fada084e51b290fa314210b73d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f090e208cb1be4629e0f8d41feca967e61a146ede7e036e056a9c4d2f9f11f9c"></a>

## code_base_integration.bitbucket_server.passwd.clear_secret_info — code_base_integration.bitbucket_server.passwd.clear_secret_info / a188d427cf0e / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-afc28030a1c2560702fb0be8c64583cb272e655832a15deaf2f00cf588b884a8)
- [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-aadee12c6732940841907585416cc6e17ce3c1ae77ca04462e950f82bf8337ad)
- code_base_integration.bitbucket_server.passwd.clear_secret_info

<a id="canonical-41330652e75e3e2a3274b2d10f53130fb86e6aece2265d00ca91b4711551fee1"></a>

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

<a id="canonical-fd988edc41e84c13c5e34b6900211d2ab24eadbecf8e6244f2f7cdf6a3dfd381"></a>

## Direct properties — code_base_integration.bitbucket_server.passwd.clear_secret_info / a188d427cf0e / 3

<a id="canonical-b7f74149bcf102b9388bd5b386dab220c6df4d4fadde33d5b0b85c05dd4ed053"></a>

<a id="canonical-59a99d6e7f656ef211a94b37b44369c6bb4ccbba89f637b0180f277b5c4a9caf"></a>

## provider_ref property — code_base_integration.bitbucket_server.passwd.clear_secret_info / a188d427cf0e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-756316908401132985145c0faf6688cf9529dd383015e62d2409cc9510171963"></a>

<a id="canonical-ed635e64490f349cd7d608c1eebeb969bbd4b732ceb1e16851ce641d21ac7bce"></a>

## url property — code_base_integration.bitbucket_server.passwd.clear_secret_info / a188d427cf0e / 5

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

<a id="canonical-52b46dacdffc685a10e4d515353b6d4ae452406a81935a7c3f99ffe0fe5d46ea"></a>

## Next pages — code_base_integration.bitbucket_server.passwd.clear_secret_info / a188d427cf0e / 6

- [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-aadee12c6732940841907585416cc6e17ce3c1ae77ca04462e950f82bf8337ad)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-463045aa8e9b205f822124632986b8adaee30fa5a60e062b0778ea3bf890506d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88143db6834a34bc65a60d0b4dc8c16873aed767d18553133e753127180ec24d"></a>

## code_base_integration.github — code_base_integration.github / 084e80ffb938 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- code_base_integration.github

<a id="canonical-cf8ac5c057ac425affd11fa87e987b11119957e384f7c921f606d5a65940a776"></a>

Type: `"single"`. Computed.

Github Integration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3e1d6241bd2372c1fb1ee7149d7129aab2999e1efc52ee07d224bfe33098569d"></a>

## Direct properties — code_base_integration.github / 084e80ffb938 / 3

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0d8c372bbd5123922c85913bb7bdaa1527ff5cc458bb4f4fcb4df4b1d7d15a12): complete subsection reference.

<a id="canonical-d1079c593b1f756c1b9686df1fbb35be6b8559c693766a54156b167d43bf7770"></a>

<a id="canonical-c9114ea291fa091c53dbb5cc2dd8e9d2f189eac55c8b68c1161d8a0f57a75166"></a>

## username property — code_base_integration.github / 084e80ffb938 / 4

Type: `"string"`. Computed.

GitHub Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

<a id="canonical-3779a73589498045ba20698c53997f2dc3c148acc5c30d6151958b267b758e83"></a>

<a id="canonical-4db6a5224c5932f84b3c9d513558446091a8622aede00c25f77dc4b1efb9c6c9"></a>

## verify_ssl property — code_base_integration.github / 084e80ffb938 / 5

Type: `"bool"`. Computed.

GitHub Verify SSL. Configuration parameter for verify ssl

Upstream description:

Configuration parameter for verify ssl

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-96381c95722dc1965956303a7f7e5a0e1cd065383cf3b12badf19a56078e35da"></a>

## Next pages — code_base_integration.github / 084e80ffb938 / 6

- [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0d8c372bbd5123922c85913bb7bdaa1527ff5cc458bb4f4fcb4df4b1d7d15a12)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-0d8c372bbd5123922c85913bb7bdaa1527ff5cc458bb4f4fcb4df4b1d7d15a12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e618d9a9f628356e5d6205b195644322bc381fd911c758234e0f0499b9cf9f4"></a>

## code_base_integration.github.access_token — code_base_integration.github.access_token / fea3352ad244 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-463045aa8e9b205f822124632986b8adaee30fa5a60e062b0778ea3bf890506d)
- code_base_integration.github.access_token

<a id="canonical-19133a832618c762e0b4766b6bca605df36b02ec2043908220bd3641b6d3d965"></a>

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

<a id="canonical-75d1844129a9e8d9acd11927fa2218ad89686666ea59eb1f6baf5986928bd158"></a>

## Direct properties — code_base_integration.github.access_token / fea3352ad244 / 3

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-54ca053cbd67d78237f31a4c226c95f25e682be23ba14bee1101225ee0f0d350): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-5d278719a96a2ab61568e9a2d7e71f6e97005f0d27e12cdef40370e96a471c82): complete subsection reference.

<a id="canonical-c488bad5c9f96a4f41a8bfc9be0344b6b63a01f57025df65f7c6112b18327e9a"></a>

## Next pages — code_base_integration.github.access_token / fea3352ad244 / 4

- [code_base_integration.github.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-54ca053cbd67d78237f31a4c226c95f25e682be23ba14bee1101225ee0f0d350)
- [code_base_integration.github.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-5d278719a96a2ab61568e9a2d7e71f6e97005f0d27e12cdef40370e96a471c82)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-463045aa8e9b205f822124632986b8adaee30fa5a60e062b0778ea3bf890506d)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-54ca053cbd67d78237f31a4c226c95f25e682be23ba14bee1101225ee0f0d350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc2f56b13c4f301f67221134112707a312d83bee605ab368b9104c2bd06fd4f8"></a>

## code_base_integration.github.access_token.blindfold_secret_info — code_base_integration.github.access_token.blindfold_secret_info / 87b6f7774409 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-463045aa8e9b205f822124632986b8adaee30fa5a60e062b0778ea3bf890506d)
- [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0d8c372bbd5123922c85913bb7bdaa1527ff5cc458bb4f4fcb4df4b1d7d15a12)
- code_base_integration.github.access_token.blindfold_secret_info

<a id="canonical-d181e91b2fc193e2b93bc12439294a22ca48e9757f98c4c97514e0ea1f5d1317"></a>

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

<a id="canonical-a9694d5dcbb4fc4b31d07994dc458a2a7c0f2e83022e59cf1de0c7bae13e946f"></a>

## Direct properties — code_base_integration.github.access_token.blindfold_secret_info / 87b6f7774409 / 3

<a id="canonical-909bb35e777321db75704ab228c8e8c501e1e165fae7e597f74e7f6438d8eb99"></a>

<a id="canonical-d1745aa87e02533544cc508ebee70d0a6a316564cb6af08add62e619cbed0fac"></a>

## decryption_provider property — code_base_integration.github.access_token.blindfold_secret_info / 87b6f7774409 / 4

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

<a id="canonical-d4e6257055b5cf8a5c83940f4798b0533ada17fbf97b4ffcfa1817a03284f287"></a>

<a id="canonical-a75892f49179d3750e6c4b59b08d85c36356ca06df1f0a38827d5ddfa3f71a65"></a>

## location property — code_base_integration.github.access_token.blindfold_secret_info / 87b6f7774409 / 5

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

<a id="canonical-de59d9089c1bc6361e987a6132968a980a72efc67f4e3e3e82086ad7d6246037"></a>

<a id="canonical-1eca29f129082a6fc109e5a42087744774f9dfb6c769edfbf905a456271f29e0"></a>

## store_provider property — code_base_integration.github.access_token.blindfold_secret_info / 87b6f7774409 / 6

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

<a id="canonical-89ecfeb191d25d215d9042fb744d35d743740a11dda25adbe07b459fcf5badba"></a>

## Next pages — code_base_integration.github.access_token.blindfold_secret_info / 87b6f7774409 / 7

- [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0d8c372bbd5123922c85913bb7bdaa1527ff5cc458bb4f4fcb4df4b1d7d15a12)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-5d278719a96a2ab61568e9a2d7e71f6e97005f0d27e12cdef40370e96a471c82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5cbec02863a940969de0f9fc901b2a47986c9ba3a00fbb7cee79be48558aa90"></a>

## code_base_integration.github.access_token.clear_secret_info — code_base_integration.github.access_token.clear_secret_info / 591c35876265 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-463045aa8e9b205f822124632986b8adaee30fa5a60e062b0778ea3bf890506d)
- [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0d8c372bbd5123922c85913bb7bdaa1527ff5cc458bb4f4fcb4df4b1d7d15a12)
- code_base_integration.github.access_token.clear_secret_info

<a id="canonical-644bc3ef1780b9553542e301953e1f5dfbd632bfe9842eaa7db5978581ce3684"></a>

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

<a id="canonical-26c0f408a184279504deb6280aac56e57e945a11c3223b0f5555a0aaf8370217"></a>

## Direct properties — code_base_integration.github.access_token.clear_secret_info / 591c35876265 / 3

<a id="canonical-7f5727188288a3d8c994e1db2d0fccb8b363728e08bf4cf2586e2eaebfa17f56"></a>

<a id="canonical-83eff2acde81d439a286228f8ab563b630be46c872d54cb8cb522ef3a9308593"></a>

## provider_ref property — code_base_integration.github.access_token.clear_secret_info / 591c35876265 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b2638d2980142866dda1a47598bb34a2f8a5ac74c93f52402f1f2adf58d90822"></a>

<a id="canonical-0163485ff5413818d87e2ec7c12e667b3d289e3b34437f970b4a33c874c8b890"></a>

## url property — code_base_integration.github.access_token.clear_secret_info / 591c35876265 / 5

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

<a id="canonical-a35ad1a7f0559bb396b5c05065942a0731572b9510c7f1b4582f6179b52445e1"></a>

## Next pages — code_base_integration.github.access_token.clear_secret_info / 591c35876265 / 6

- [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0d8c372bbd5123922c85913bb7bdaa1527ff5cc458bb4f4fcb4df4b1d7d15a12)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-9e6712137098c88e9972ddfe960990503f708ce52facf7ba62b9fe68e3d1ddbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa7def36e112dbd2a197e1b5a7732b2edf678898cb75a19b010a3e16cfae383d"></a>

## code_base_integration.github_enterprise — code_base_integration.github_enterprise / 97079ac236e2 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- code_base_integration.github_enterprise

<a id="canonical-3513136be3b0b18d7ff67a66b5119b2611c3dbdf498633010b48ddaea8a446e6"></a>

Type: `"single"`. Computed.

Configuration parameter for github enterprise.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9476347bb366444f70217d49d10c12a703378c2b5c127daf881c36c9045ecc3b"></a>

## Direct properties — code_base_integration.github_enterprise / 97079ac236e2 / 3

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-73face7c4b0715bbb828b81ded41ef1de473f33cf3e4c3d3f1f6fc647c7f2771): complete subsection reference.

<a id="canonical-1cf7cd8bbb504f4153b71add6cb8aa52affdd2d163706745244930a042fff02d"></a>

<a id="canonical-0682fdf90e718d25cf33a45bfb173098e7777bf2f6d0205128e0965fed5d0c31"></a>

## hostname property — code_base_integration.github_enterprise / 97079ac236e2 / 4

Type: `"string"`. Computed.

GitHub Hostname. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

<a id="canonical-2414aca3debcc30f1d741aef72d5267e840fd46353459c69fcd2b1e616b6ed54"></a>

<a id="canonical-f94b07be11f503b1cd8c1aa0691d60fc7f6ed5fdeff8dfad5ba3afbfb83ab568"></a>

## username property — code_base_integration.github_enterprise / 97079ac236e2 / 5

Type: `"string"`. Computed.

GitHub Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

<a id="canonical-ae457299dda64a3b1be1338018d57c33d7e0582f1becf9f33ebfb9eef080e76c"></a>

## Next pages — code_base_integration.github_enterprise / 97079ac236e2 / 6

- [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-73face7c4b0715bbb828b81ded41ef1de473f33cf3e4c3d3f1f6fc647c7f2771)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-73face7c4b0715bbb828b81ded41ef1de473f33cf3e4c3d3f1f6fc647c7f2771"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-691685a4cb098bfd2449e74b26a35586a7af7aa81c217b7b5ae3ddd6b0ecbe96"></a>

## code_base_integration.github_enterprise.access_token — code_base_integration.github_enterprise.access_token / 7e32e9a3b380 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-9e6712137098c88e9972ddfe960990503f708ce52facf7ba62b9fe68e3d1ddbe)
- code_base_integration.github_enterprise.access_token

<a id="canonical-036481ddd2ab01edcaf15d8087a190cc02202fbf3d8b18fcbc956262a2f1e66b"></a>

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

<a id="canonical-6b530fb357d66d281409af3e1c64bbf2048e6d41eaffdbe4bcf223da1222da4a"></a>

## Direct properties — code_base_integration.github_enterprise.access_token / 7e32e9a3b380 / 3

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-b15c4e0a825c5f5e1ed60bf22a8442cfd19a6a6c2733aa8078ec68ff185fcdce): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-28056f29f9de0ea40154ff3458662434aa89472ea18eba7327deb5f54b255ee4): complete subsection reference.

<a id="canonical-99fd5c1efb9b065029fe463e396331aca7899174240d6a997eff0103d2d0b6d6"></a>

## Next pages — code_base_integration.github_enterprise.access_token / 7e32e9a3b380 / 4

- [code_base_integration.github_enterprise.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-b15c4e0a825c5f5e1ed60bf22a8442cfd19a6a6c2733aa8078ec68ff185fcdce)
- [code_base_integration.github_enterprise.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-28056f29f9de0ea40154ff3458662434aa89472ea18eba7327deb5f54b255ee4)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-9e6712137098c88e9972ddfe960990503f708ce52facf7ba62b9fe68e3d1ddbe)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-b15c4e0a825c5f5e1ed60bf22a8442cfd19a6a6c2733aa8078ec68ff185fcdce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10b5e065f1f45bd01dc6bc161c419ad1efaa83f34e92b56df09e63ba56dbc44a"></a>

## code_base_integration.github_enterprise.access_token.blindfold_secret_info — code_base_integration.github_enterprise.access_token.blindfold_secret_info / a965e5e44f24 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-9e6712137098c88e9972ddfe960990503f708ce52facf7ba62b9fe68e3d1ddbe)
- [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-73face7c4b0715bbb828b81ded41ef1de473f33cf3e4c3d3f1f6fc647c7f2771)
- code_base_integration.github_enterprise.access_token.blindfold_secret_info

<a id="canonical-9c3332da1008c88762b29befe9b5d45d081a535e1018edee61b682e0750b5325"></a>

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

<a id="canonical-d20c6fc2a6bbcf3cf9fecddd60e391eb0693a63c699d7bd633cfc56d86dbb0df"></a>

## Direct properties — code_base_integration.github_enterprise.access_token.blindfold_secret_info / a965e5e44f24 / 3

<a id="canonical-770857d8d635d5e05c203b0b2cf00705cfd07d061097046f09423b2dfaaa55f2"></a>

<a id="canonical-22a3ad7396e39b56b59733c5bfa41e2f5bf34bce43c5bc1a487cf348cd600ab3"></a>

## decryption_provider property — code_base_integration.github_enterprise.access_token.blindfold_secret_info / a965e5e44f24 / 4

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

<a id="canonical-63c53d354611ba15a4a8203301e564f848a33ea192d0cb44c02e2200ef9f35b5"></a>

<a id="canonical-55455b2b13b8876a26a17e9519606ce8bb89e56cbc7de2775cce356e4f96c652"></a>

## location property — code_base_integration.github_enterprise.access_token.blindfold_secret_info / a965e5e44f24 / 5

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

<a id="canonical-e86fc5eb9e2f69f994093dfd0694212d499cf6c65cc2eba767b8141ea5b2735d"></a>

<a id="canonical-55ab436a9ca6e9aa91b47ede98175380c8794657f44aa493cd29cc7fdfc850e6"></a>

## store_provider property — code_base_integration.github_enterprise.access_token.blindfold_secret_info / a965e5e44f24 / 6

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

<a id="canonical-c20398766b69003e6c193498b0652119deebbf5547587a7cfac86f438cdb4b68"></a>

## Next pages — code_base_integration.github_enterprise.access_token.blindfold_secret_info / a965e5e44f24 / 7

- [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-73face7c4b0715bbb828b81ded41ef1de473f33cf3e4c3d3f1f6fc647c7f2771)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-28056f29f9de0ea40154ff3458662434aa89472ea18eba7327deb5f54b255ee4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aeee7be0c708575695588504990117aaf75af292b42b5e4513163889d3605b77"></a>

## code_base_integration.github_enterprise.access_token.clear_secret_info — code_base_integration.github_enterprise.access_token.clear_secret_info / e1f552d904b5 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-9e6712137098c88e9972ddfe960990503f708ce52facf7ba62b9fe68e3d1ddbe)
- [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-73face7c4b0715bbb828b81ded41ef1de473f33cf3e4c3d3f1f6fc647c7f2771)
- code_base_integration.github_enterprise.access_token.clear_secret_info

<a id="canonical-6e7219779df286164a65eb60e31436f5dd2d26d12d63e46f702a38e884782152"></a>

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

<a id="canonical-99e16fc1e5a9acee05d8325c5c85dc91c2ade040aa5a3886d4e3b670359a2e3a"></a>

## Direct properties — code_base_integration.github_enterprise.access_token.clear_secret_info / e1f552d904b5 / 3

<a id="canonical-a28e94c74ac8d655e1710253e69a0a8724470f7658bf8845f1790a6e99b29561"></a>

<a id="canonical-423a568bae6c2bfee296cd2bb7de57869242c76a39203962558ef4f9b1d9c0d5"></a>

## provider_ref property — code_base_integration.github_enterprise.access_token.clear_secret_info / e1f552d904b5 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-d04de4cf9216d5f3996956847fcd4bcef2c99564b9f2eda573147b890f50f52b"></a>

<a id="canonical-49fbbae8e8acf76667372cff1ac16dde96524777bb76d8f77e2ed6a217c31da1"></a>

## url property — code_base_integration.github_enterprise.access_token.clear_secret_info / e1f552d904b5 / 5

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

<a id="canonical-744c545e01ab42bf5b010ef039539fb3913a1d58e0ac620e049f59dd1fb90803"></a>

## Next pages — code_base_integration.github_enterprise.access_token.clear_secret_info / e1f552d904b5 / 6

- [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-73face7c4b0715bbb828b81ded41ef1de473f33cf3e4c3d3f1f6fc647c7f2771)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-c153e4e3b91fb5dc774723280a6e60c47fa195cccb0b0ead5c42c85315b3f6ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb426ac3d761ccdf692e751222a9d5294b6e219ed5363cfc688b9c4f9734db62"></a>

## code_base_integration.gitlab — code_base_integration.gitlab / e5087c02e8f0 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- code_base_integration.gitlab

<a id="canonical-b706bf9d0fdaf100196de89cfaf07baeb63988c7b873311b8285984594f2d2cd"></a>

Type: `"single"`. Computed.

GitLab Cloud Integration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-036b0cc243b3f9dd76db37546492d775840ef628483da544a8ea653c67aa1de1"></a>

## Direct properties — code_base_integration.gitlab / e5087c02e8f0 / 3

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-d1e449fbeb0776e368fe2f782cd46328a3f82a028831ca0deeb6f0fbb305a809): complete subsection reference.

<a id="canonical-29476544bd9a17d330dbf60462334673b3ca8da8812e59c63e6a18e80758aa18"></a>

## Next pages — code_base_integration.gitlab / e5087c02e8f0 / 4

- [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-d1e449fbeb0776e368fe2f782cd46328a3f82a028831ca0deeb6f0fbb305a809)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-d1e449fbeb0776e368fe2f782cd46328a3f82a028831ca0deeb6f0fbb305a809"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7293c439eab12404fd5a4949d6f3dd8f42c3c65b1b5cd2aa3f2da5aa8e6f7cea"></a>

## code_base_integration.gitlab.access_token — code_base_integration.gitlab.access_token / 8b113b6285c3 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-c153e4e3b91fb5dc774723280a6e60c47fa195cccb0b0ead5c42c85315b3f6ef)
- code_base_integration.gitlab.access_token

<a id="canonical-8d302ddc5570f73990fbe407e090bda894a9008ee6e96d82b9d4f67aa17d365b"></a>

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

<a id="canonical-bf1aabc1ec3f848125c392c0fc7f513193283da636f0956d92251bb2a64d9b41"></a>

## Direct properties — code_base_integration.gitlab.access_token / 8b113b6285c3 / 3

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-5deb54251cd677085e10091003d1e9b96795f95060b361d57e0da3f4f589be7c): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-d87f2e7ef6abebbea1fef9d8a101284e1487e792a6307f64f3a95008f8807786): complete subsection reference.

<a id="canonical-5942db0c4ddac0f789d1843509e12aec92995864d7d916c4cde40073b0234a9a"></a>

## Next pages — code_base_integration.gitlab.access_token / 8b113b6285c3 / 4

- [code_base_integration.gitlab.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-5deb54251cd677085e10091003d1e9b96795f95060b361d57e0da3f4f589be7c)
- [code_base_integration.gitlab.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-d87f2e7ef6abebbea1fef9d8a101284e1487e792a6307f64f3a95008f8807786)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-c153e4e3b91fb5dc774723280a6e60c47fa195cccb0b0ead5c42c85315b3f6ef)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-5deb54251cd677085e10091003d1e9b96795f95060b361d57e0da3f4f589be7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abfe47f63e6c106af31812e914d37f5a7f1910eb97d3e09db2b22bde6bdb48c3"></a>

## code_base_integration.gitlab.access_token.blindfold_secret_info — code_base_integration.gitlab.access_token.blindfold_secret_info / b9ef535d54a1 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-c153e4e3b91fb5dc774723280a6e60c47fa195cccb0b0ead5c42c85315b3f6ef)
- [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-d1e449fbeb0776e368fe2f782cd46328a3f82a028831ca0deeb6f0fbb305a809)
- code_base_integration.gitlab.access_token.blindfold_secret_info

<a id="canonical-c092753bd8915a86b100421da39a01233710168d193a0a0e08d6fc92d83c41b6"></a>

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

<a id="canonical-4b16475048ace233452ec4d9cbb0e04aad0b25384e6d518457047c02b4c3d1fb"></a>

## Direct properties — code_base_integration.gitlab.access_token.blindfold_secret_info / b9ef535d54a1 / 3

<a id="canonical-7fde553a667d93fe30319a8115e63cb90341dc50a4c91b567d30c1607182eb3e"></a>

<a id="canonical-6fb42fafe898a58a7f1e511f38105ce5328646e5356ccd8cc6014d2c3b51db8d"></a>

## decryption_provider property — code_base_integration.gitlab.access_token.blindfold_secret_info / b9ef535d54a1 / 4

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

<a id="canonical-11aaebce12530e8de2df6fc05d93c213f5db6bb023a9070123ef4e64be33fb99"></a>

<a id="canonical-f10532ba3eb65fa481443caa39dac30fe7362d856823afd5d2f54cc7730198ff"></a>

## location property — code_base_integration.gitlab.access_token.blindfold_secret_info / b9ef535d54a1 / 5

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

<a id="canonical-6344f4fabbb2a26f5fd2d7f58e16ec16af79f46829c0a0ad9e98880909371f14"></a>

<a id="canonical-4ebb947fcc20aa8d434a90f3990b667ea167c73ae2ef81e795692d96ca86683b"></a>

## store_provider property — code_base_integration.gitlab.access_token.blindfold_secret_info / b9ef535d54a1 / 6

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

<a id="canonical-f1834c3ad3ad903ad111220d0fd6b7900aeee0bb1d1d02ee730ccd91bbc1e7fb"></a>

## Next pages — code_base_integration.gitlab.access_token.blindfold_secret_info / b9ef535d54a1 / 7

- [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-d1e449fbeb0776e368fe2f782cd46328a3f82a028831ca0deeb6f0fbb305a809)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-d87f2e7ef6abebbea1fef9d8a101284e1487e792a6307f64f3a95008f8807786"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-161044ab8b35cfdbe0ebdd11808af1fff3531a86434fd0cc30d1b2b6a8f790e5"></a>

## code_base_integration.gitlab.access_token.clear_secret_info — code_base_integration.gitlab.access_token.clear_secret_info / 2b785f7643cb / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-c153e4e3b91fb5dc774723280a6e60c47fa195cccb0b0ead5c42c85315b3f6ef)
- [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-d1e449fbeb0776e368fe2f782cd46328a3f82a028831ca0deeb6f0fbb305a809)
- code_base_integration.gitlab.access_token.clear_secret_info

<a id="canonical-59f0671a9e8894e0cb71effb6f82e324126ce7e8dea836a24e215cb74facf189"></a>

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

<a id="canonical-4bfce5022eee3cf5b39f0da2305e6c64c348d94951c10d9aceca3bec9267e59b"></a>

## Direct properties — code_base_integration.gitlab.access_token.clear_secret_info / 2b785f7643cb / 3

<a id="canonical-63b49b6cc852f148cec80caa5ecda3f0a65b7688db1bf6b6da2034e36d7d7b62"></a>

<a id="canonical-9707a0d24497b488d52aab1293bbe1732458662720a007a9257246311abedf4f"></a>

## provider_ref property — code_base_integration.gitlab.access_token.clear_secret_info / 2b785f7643cb / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-360647f3e286390a500fca52651a15c5ae683b1b9ccd4324705ecef5f453bb1a"></a>

<a id="canonical-dccc5b56a02502cda1365f5ef8e3122a585563fa3ce4e888e61e3a8c589c584b"></a>

## url property — code_base_integration.gitlab.access_token.clear_secret_info / 2b785f7643cb / 5

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

<a id="canonical-c569dc72d59e568ed9f199041895ebc44a941666f48da425378a0fe7bab9c088"></a>

## Next pages — code_base_integration.gitlab.access_token.clear_secret_info / 2b785f7643cb / 6

- [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-d1e449fbeb0776e368fe2f782cd46328a3f82a028831ca0deeb6f0fbb305a809)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-bf448306c188c0c01504013875c4737a4140687b4a73c8385e9e265cae53c79b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-761ccd681bde4dc16fcd8c118a6b6d47098d5d30f01fe54edc3b110697f67f3a"></a>

## code_base_integration.gitlab_enterprise — code_base_integration.gitlab_enterprise / 6453eb358440 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- code_base_integration.gitlab_enterprise

<a id="canonical-2ff7ce8141b849a71d220e5171e7382c617ceb0b96e2c4b7a91d996e441c467a"></a>

Type: `"single"`. Computed.

Configuration parameter for gitlab enterprise.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c6f6e7a14167fec2713d6c6f585d4502f95e58cb48b4c4110ac5d2d508bacf7f"></a>

## Direct properties — code_base_integration.gitlab_enterprise / 6453eb358440 / 3

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-c568bfa5a7ef2a938a2f7bafe446df2971614580a1e80a17dddab59303dce00c): complete subsection reference.

<a id="canonical-2ec053cbadbbe011ce1db8d54e7be8d82d7d2d800198c590edf01b3b4c204e75"></a>

<a id="canonical-5e9dc9e70725d97c2a42d25ce3a8ded111ac327d5128a664acf1812b42bbb1b4"></a>

## url property — code_base_integration.gitlab_enterprise / 6453eb358440 / 4

Type: `"string"`. Computed.

GitLab URL. URL or URI reference

Upstream description:

URL or URI reference

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
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
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-7262f2267ba6a669e6ff2820f32b933ea6aeff9aacc0200a41cc4cf754f3fdce"></a>

## Next pages — code_base_integration.gitlab_enterprise / 6453eb358440 / 5

- [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-c568bfa5a7ef2a938a2f7bafe446df2971614580a1e80a17dddab59303dce00c)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-c568bfa5a7ef2a938a2f7bafe446df2971614580a1e80a17dddab59303dce00c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1965d7ff06edd745f4311bd0ef44c05eccc304e7e36a62111b9bff14692aca25"></a>

## code_base_integration.gitlab_enterprise.access_token — code_base_integration.gitlab_enterprise.access_token / 132456758212 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-bf448306c188c0c01504013875c4737a4140687b4a73c8385e9e265cae53c79b)
- code_base_integration.gitlab_enterprise.access_token

<a id="canonical-bbd8377b227c0c9f0cc64a3464ac21d6580b48513fa3d113a524df45e01703fa"></a>

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

<a id="canonical-9dcf9bce1664a9b3779055fcd8f15e597a2723c1b0dc348cad5c7ef15e8a3364"></a>

## Direct properties — code_base_integration.gitlab_enterprise.access_token / 132456758212 / 3

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-acf75d62c0b4fed7918ef43638c991e53a88d0b3fa8bffe524d8074191c51335): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-0bfbe110c0913cf484f6b30384ca1a95454e9a698c9201a07ec38dccc96cfc85): complete subsection reference.

<a id="canonical-91800524185a8dcb3e1069b28482aff2f6c592eb9494e4faa82a3700b5dc7c55"></a>

## Next pages — code_base_integration.gitlab_enterprise.access_token / 132456758212 / 4

- [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-acf75d62c0b4fed7918ef43638c991e53a88d0b3fa8bffe524d8074191c51335)
- [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-0bfbe110c0913cf484f6b30384ca1a95454e9a698c9201a07ec38dccc96cfc85)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-bf448306c188c0c01504013875c4737a4140687b4a73c8385e9e265cae53c79b)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-acf75d62c0b4fed7918ef43638c991e53a88d0b3fa8bffe524d8074191c51335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be0e9cb2724f5f501c2e7b8b865c36ddda4c9e6131db92ed3778890a648eb8ac"></a>

## code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / 2e92e0c0917e / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-bf448306c188c0c01504013875c4737a4140687b4a73c8385e9e265cae53c79b)
- [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-c568bfa5a7ef2a938a2f7bafe446df2971614580a1e80a17dddab59303dce00c)
- code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info

<a id="canonical-78834be3849252f6ca318e52d2d5b7340ba8f2c4ffdeb51a5a9584cb3e19aa44"></a>

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

<a id="canonical-0b47101badf26c4f50cb776e650dcf280dcc4499ee740b3e4f15f86c179cf598"></a>

## Direct properties — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / 2e92e0c0917e / 3

<a id="canonical-17fdf860933f8a11d5871b03ddf1a42ca74ee757349eb95d90c1c1b7f71f5637"></a>

<a id="canonical-e599e4e9869e8e109b07178a4268c053afe347d763e88667a765ab319bc5c157"></a>

## decryption_provider property — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / 2e92e0c0917e / 4

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

<a id="canonical-0002b42b4b1575925aabfe5aab15f7a33d3521299a64ddefe1cfeb8407c09494"></a>

<a id="canonical-858fb2ce8518d40af71764654e57bb07773b6195e405ac5bee0f32a80d84cde9"></a>

## location property — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / 2e92e0c0917e / 5

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

<a id="canonical-771b75379dc0c5070b59206ead70bd0f1546ae54b13f58967a74806a42e831c7"></a>

<a id="canonical-4307cf2698162c5d20499e9db759201c4a63c7d2da7a5ebeb3fd8e6999e497ea"></a>

## store_provider property — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / 2e92e0c0917e / 6

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

<a id="canonical-daa5db49c54446331b27a3d461438e05a7f93eca46f935935f3f7e76b7fc3d8b"></a>

## Next pages — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / 2e92e0c0917e / 7

- [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-c568bfa5a7ef2a938a2f7bafe446df2971614580a1e80a17dddab59303dce00c)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-0bfbe110c0913cf484f6b30384ca1a95454e9a698c9201a07ec38dccc96cfc85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2cc5128e90b69bfa7a16134039f3eabac061d1cda7a2bee824772954474c184"></a>

## code_base_integration.gitlab_enterprise.access_token.clear_secret_info — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 6604318de863 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-537d94efecd4f12f0aead5811512dc50c9109ba3fbc8772343bfd90165c83c9a)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-bf448306c188c0c01504013875c4737a4140687b4a73c8385e9e265cae53c79b)
- [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-c568bfa5a7ef2a938a2f7bafe446df2971614580a1e80a17dddab59303dce00c)
- code_base_integration.gitlab_enterprise.access_token.clear_secret_info

<a id="canonical-62589d6b6e7d264d4522ab73999d9ff51aea537b92899196ad6754ab46c5f8e9"></a>

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

<a id="canonical-155548ae8cac319f22a94deef7d377f31c9da43442e8e5564f2f65479da53b75"></a>

## Direct properties — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 6604318de863 / 3

<a id="canonical-dc8e75c82e3c0da3335e253d2d03848fdf0117f266aa06891d148ae415a2b60f"></a>

<a id="canonical-6d5f2d6f28785210ce60eb374c9b8ba0c7eb87ea9cb9d9e8895eed7a10b481ab"></a>

## provider_ref property — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 6604318de863 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a23e5c61be52b4a164bdc14b9ffdfcd57b705191cc5e6259a2b1cc323bf943b7"></a>

<a id="canonical-e6883d8faed31a138dc900385d0036b293f688c3f8a8765c718a1e4fbaf4bd55"></a>

## url property — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 6604318de863 / 5

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

<a id="canonical-609a2c2fdcad3fff58b932e0589f132bde41aaacb08453becb6ee068bfb0b553"></a>

## Next pages — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 6604318de863 / 6

- [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-c568bfa5a7ef2a938a2f7bafe446df2971614580a1e80a17dddab59303dce00c)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
