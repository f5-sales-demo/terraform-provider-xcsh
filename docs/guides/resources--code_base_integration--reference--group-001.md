---
page_title: "xcsh_code_base_integration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration reference."
---

# xcsh_code_base_integration reference

<a id="canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c61e356ae5de7d8ec3cbdbd479d2f71d7f311e5b5da1f4135bf6365042a860c"></a>

## Property reference — Property reference / afd93695f02e / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- Property reference

<a id="canonical-a73d0f26aaf0cdf7682829025d5989f09bf774c2fd04d209fe1e2ecbf970f9b2"></a>

## Direct properties — Property reference / afd93695f02e / 3

<a id="canonical-b283ed0ab0e134f01c3e0b7cb054f3db5197b189e3cafcf6b0cd823f81516067"></a>

<a id="canonical-b044b654a3df4fc873b2c6a6d495e9b93327abc1da1b82429b97332726a6e241"></a>

## annotations property — Property reference / afd93695f02e / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1): complete subsection reference.

<a id="canonical-79194046c2f16815833ecf7958d7296d3ac656be586ef2e8fa3aa2aa7a47d4b6"></a>

<a id="canonical-417d82fbcd2b3ea5641af85378efac62453a4373bbfef215bcf8f8a37fb64e2e"></a>

## description property — Property reference / afd93695f02e / 5

Type: `"string"`. Optional.

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

<a id="canonical-4f55a03c1fa6a92b9234459e68dc97555e00c064532b3a8b3b46c85c3c345157"></a>

<a id="canonical-93b5a7d1eeab757964817a52312a4160c99593580f461d8ce334c3e2370a2542"></a>

## disable property — Property reference / afd93695f02e / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d236adfe3f606e79bcc5253abcc56ee1ada6687e6d992f725de736f317b45750"></a>

<a id="canonical-d52836a2b624c66680bbeee870651f811b9e4f02b62fb57c74b60e11573c34a3"></a>

## id property — Property reference / afd93695f02e / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-620a3b4beebe118b630336e4e2beac9298e61ee9b57eec56e00ac29a6690a32d"></a>

<a id="canonical-bf909fb7fe3f67c75937951b5751b94e0fb7131c392da90c758062aaebec4182"></a>

## labels property — Property reference / afd93695f02e / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-6be01bc1e6e6476b8be86aa94a3bc2c49e52d5ac7fbb94bc97943c49851a9df2"></a>

<a id="canonical-d8b3ca38afe7b48b762d0c9ff165f77e53d2401e6b44cbe8038f30ca7eebda85"></a>

## name property — Property reference / afd93695f02e / 9

Type: `"string"`. Required.

Name of the Code Base Integration. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-115a5f3f9d84a70a1d054a12fc954d87639509d396e814655d0f024c3af873c6"></a>

<a id="canonical-5c32e17492e7a1fd437ba7da028ff1c8b07fdd9b394a9007bf89f54d11c569ab"></a>

## namespace property — Property reference / afd93695f02e / 10

Type: `"string"`. Required.

Namespace where the Code Base Integration is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--code_base_integration--reference--group-001.md#canonical-41f3c8550b8529fda8a9906a2b929b165847756bf3daba96287289d43c9b3323): complete subsection reference.

<a id="canonical-4b8c47ffce8550c2dd03abbbd25badc07006d28c5f1f75c3c8d3fbbb51457d05"></a>

## All schema paths — Property reference / afd93695f02e / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--code_base_integration--reference--group-001.md#canonical-b283ed0ab0e134f01c3e0b7cb054f3db5197b189e3cafcf6b0cd823f81516067) |
| `code_base_integration` | [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-7db312740ff9369e8db8dad8e0ec1ec538a79022ac2f2870cc40a0d9da22a419) |
| `code_base_integration.azure_repos` | [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-00dc316a4def9f5e509f7fb12a438c18a588a6a49b4601e88f2bb1d931dd895b) |
| `code_base_integration.azure_repos.access_token` | [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-e316c0069ae90f69c111cc652f9520fe0a6aef336c6809ab28a25af23b682e44) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info` | [code_base_integration.azure_repos.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-7ed576f11995961ca2b16dbfa3566a4b2b5698b18681c005be7e605e8f3fb17c) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-9ce84057be61bf9427cfe48ac78a404346f18f022377a41401d61fecbc183a6e) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.location` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-a50d09c74d90e59db639c16750fe25f4ff7720e33a3a4b2698033156f81a23a9) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-800bffad0755ff94354ea30098907f51b8d643e77c5806c4bb120b24f5276dff) |
| `code_base_integration.azure_repos.access_token.clear_secret_info` | [code_base_integration.azure_repos.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-338115c3d8fc70ad7ca3c35d97564d27388bfaa017761a83dbb189450c83fcf3) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref` | [code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-3667fa3bf3b7fcae62fecf364770867f7044ba2e2db394f3194ac000620d5b05) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.url` | [code_base_integration.azure_repos.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-dc1a2d7f7b6a38f739e1d1eeb458064a13a964a979cca71d8dfc664b9f980688) |
| `code_base_integration.bitbucket` | [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-0a4d2fdcc867e8b511f68ee6aa50756c44f039d1a908f8a061d9c8e630da13b7) |
| `code_base_integration.bitbucket.passwd` | [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-0565c8421d5b18964a3baf81494d2974584db89bd37320bef33f928109424491) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info` | [code_base_integration.bitbucket.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-e00ea85babdd0e745c147ed63079d5c215142de358e1c9044e435323670a669e) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-1cf0c0b3f2e5cf20f59a1c8933805ba343f4b61084a126381f8d5ef7c7c6a3c4) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-8927d9f16365743f0d031e9bda91e721b6386739a53845a53652d5227b6be6e6) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-9621adf671a7899327b4cae0c64e3e3e8737d189c9279b0224430d1b1f978645) |
| `code_base_integration.bitbucket.passwd.clear_secret_info` | [code_base_integration.bitbucket.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-deda4d307869154bca2627589a15f36d14a4a885d22e7fc4675fe079afac27dc) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-50a38167d7b443c1569392210d610a987f5247569a5f08b2027aff171c969524) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.url` | [code_base_integration.bitbucket.passwd.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-818c7aeecb6809b08e2d09fff435592c06ff62e647ef341a1d7db86e4150c8d7) |
| `code_base_integration.bitbucket.username` | [code_base_integration.bitbucket.username](resources--code_base_integration--reference--group-001.md#canonical-3560b49c5099811cfa8a7be933964a5a45bc65aad65bd5f8946394dd07a7bd40) |
| `code_base_integration.bitbucket_server` | [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-55aa224de5109ddcf551cc27909a84a97c96158c72a4f97c7f203b3477324901) |
| `code_base_integration.bitbucket_server.passwd` | [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-10d857dfd660ce7c231778fe7d3f9e93f380aa6802de909d753800a3d4700816) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-ab26b469dd65b02e3ca8190378da5e37365f07e08befe54cdfd4c3320ef350f1) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-6008068593c6dfe1563010e5df47d98cbfa7116ceb06fb3fc0e509f8311e3d41) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-ca44b70186af98f14a7f105fc70d5690f0b30134d794b86a8bbdcd53d7adf767) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-ac8fa94793a5dbd08ae3c973991363d103a13ed4b731efe31cf203167ed4d528) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info` | [code_base_integration.bitbucket_server.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-10d4ce9cfa01b64592a154b4153ece8b461a06fe964f4451042d312e87f965f3) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-96d61d87da45340d2493675e4b258e72c314fd0f7e7afe1e5e18b3190153cccc) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.url` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-9eeb2a5eb4af76353b351d6bf16a49cb96c74064315f8638a3755cecbdcf574b) |
| `code_base_integration.bitbucket_server.url` | [code_base_integration.bitbucket_server.url](resources--code_base_integration--reference--group-001.md#canonical-d547d59c022fb383c51e67c5ca2baa97e50401bd5dd5a9806b14d345a40816b3) |
| `code_base_integration.bitbucket_server.username` | [code_base_integration.bitbucket_server.username](resources--code_base_integration--reference--group-001.md#canonical-fe40181c73ad3dc98772a52b6cbda3ab9c15fd2238662dd6f4ecc54395288fa5) |
| `code_base_integration.bitbucket_server.verify_ssl` | [code_base_integration.bitbucket_server.verify_ssl](resources--code_base_integration--reference--group-001.md#canonical-64b6b30a8a0ab638c2b4170bdd0d688ef971a8a47abc65cc6f451c05ed66bc73) |
| `code_base_integration.github` | [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-39612f9ca01618a33329739cd22396068e3cb51ec34ecea747d5821ee6186d01) |
| `code_base_integration.github.access_token` | [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-1a4c71b32a390526d22d92e09a2f5921be8076592030e5d74a8e6419108d3a7a) |
| `code_base_integration.github.access_token.blindfold_secret_info` | [code_base_integration.github.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2c9148543b7af6aa9df2295f43c9c62d7df31d0ff8f6047bb543527e6486f36a) |
| `code_base_integration.github.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-679de269aa5c59046e91e923c1b38184a6f52a5ae4bae5e346dbb5494a784b37) |
| `code_base_integration.github.access_token.blindfold_secret_info.location` | [code_base_integration.github.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-bb58cf548f7d75a5393b8f4825d93f18e9a48743ab59eea60fa331a2a47ea78b) |
| `code_base_integration.github.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-a5fc453bdce912e01d3cd5a1572346059271c0d4392639c293af97655ca4a3bd) |
| `code_base_integration.github.access_token.clear_secret_info` | [code_base_integration.github.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2f36955064d76eb5dac106380eab1b3f0f713d8ea266b6a7d272d21d29a0cfe5) |
| `code_base_integration.github.access_token.clear_secret_info.provider_ref` | [code_base_integration.github.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-a37bfa1e70ca12dfc236f95d635f3467ad7472250afe1ab78ce90951414070f5) |
| `code_base_integration.github.access_token.clear_secret_info.url` | [code_base_integration.github.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-584333d773a93d10c1c8f1b2f92768aebfb723e570b62fc80b335ddb82e53a9c) |
| `code_base_integration.github.username` | [code_base_integration.github.username](resources--code_base_integration--reference--group-001.md#canonical-10342b34b089e207821da834ec64c665b5db1601327af776a78dc0910804ae2a) |
| `code_base_integration.github.verify_ssl` | [code_base_integration.github.verify_ssl](resources--code_base_integration--reference--group-001.md#canonical-38ea48401051647dffcfa62218379eb9f5dc38c7a75489db6c97fe6d0a99be7c) |
| `code_base_integration.github_enterprise` | [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-f7139c43615150249f93ede66a323e27e220d3aec3a9d78776b0eebfb4ca0ced) |
| `code_base_integration.github_enterprise.access_token` | [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-da0efb41ae3d93c98028c6221f9151828b7b570e12b861a8acf4712f97a44804) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-787412a7621de2571985237ede4fc47be231eaec9090aefa46675e2db16a6059) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-b15c899ada42d65f0855715ba1844d190e91a25785bcdf1ec19f40e77fa1b75e) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-f1bc8c0ff5f2114692446983f83364a1030b7078be344aa5fd7478d3c78dc1a8) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-8e900b8b3bfcb75e7efc0e48e946d79f3193fefc942de1a2f831de26ce1f99f9) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info` | [code_base_integration.github_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-adfa34f99f6695a16d6f6770600d0ee663d06579d6fc37a5970dfd2789b9abca) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-b83db40ec772dfb7c8db482f6b943779728fb4b6c51d6fbd4b5b2bd10346757a) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.url` | [code_base_integration.github_enterprise.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-10ece9c8106e87e3886303df8d0875a0ab76897aa1eba5d19a2f02d6ba2ec304) |
| `code_base_integration.github_enterprise.hostname` | [code_base_integration.github_enterprise.hostname](resources--code_base_integration--reference--group-001.md#canonical-b277d28a4f12af2e8a44e4b406dec22fe0c46d8f9ea0d92c238ef20ae372562d) |
| `code_base_integration.github_enterprise.username` | [code_base_integration.github_enterprise.username](resources--code_base_integration--reference--group-001.md#canonical-43e3a25796f80c67276d7a0bc34702760e8f88a848058aa373f8b4f8f98cd415) |
| `code_base_integration.gitlab` | [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-507315619a769bc7e946ac193220669528e070d0491eb5a25860181d9b6f278a) |
| `code_base_integration.gitlab.access_token` | [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-6d3233cca8e47e75a50c83137d9e68627e53b3b8c78fce6aa353992bb694462d) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info` | [code_base_integration.gitlab.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-33a5b3cd2b27e9ffd84ba649d6df7f56538aeec3998e8f972f1f1650c2ac5ccb) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-adad1d06b9bfd2212a73040d624f23880e9dc56b52b5ee272d5ad148f7dc20ad) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-419d8aff7d5caa6882f539a0dfe59828a807656015f0cfbdacd4707f4ac3ee9e) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-3950009eb6e46ea186a7ebb271decba110a7bdf6028825a1f101edfaf442384a) |
| `code_base_integration.gitlab.access_token.clear_secret_info` | [code_base_integration.gitlab.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-250afca179485e4c9d63761f57af3c520a58ad2213587be34e553133d4c2bac6) |
| `code_base_integration.gitlab.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-247e13a062dbdf0b1f70c3613342c76879a34bfaa15d27ac076728b26c5f009a) |
| `code_base_integration.gitlab.access_token.clear_secret_info.url` | [code_base_integration.gitlab.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-358c7913d7bb3ff8dc4fa96f76b52117c87bbc96b0218a7b5ba676b10bcd914c) |
| `code_base_integration.gitlab_enterprise` | [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-fe1b4b49c1e5f0423f59e1ca64edb9ffc864c48141b562d5c8648058affb807d) |
| `code_base_integration.gitlab_enterprise.access_token` | [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-2fe93cc4b13b712c290840055a748c25a8d975380149a1cb4ef07db231975624) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0bc8cce0e2ee6f29537342ae7eb38a68ff92cfdd18b288a6b957645bd307fe21) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-17cbc7dfcf4f653891b548fbdfaaca694dd89ace6b77e640b1330723b52cf02a) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-cc91b29f58b478ccedb5528e936e924a239747c9f44dd8a79bb45f7dc3ba82db) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-31e7fb244cb7fbc6be14b1ed8db4191ff12ef14a3cbc7256dda8d302afa4cf6c) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-d67fe49b5d14290e668e3fa6a414e608a0d8e6a70cd174a56ab09aedc8fe1d88) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-5ce0392955d93dc943965336e80116d1e6798b11958e2dc4bea810aa101422bb) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-107a41ebf8caa57a7eaf0cb33c99c6266d129eb6dee77b6e878316f4592832f9) |
| `code_base_integration.gitlab_enterprise.url` | [code_base_integration.gitlab_enterprise.url](resources--code_base_integration--reference--group-001.md#canonical-2f440c94c2db2d97d70322b5dd8678ce03fc3fee13d2b85901acbcd116587baa) |
| `description` | [description](resources--code_base_integration--reference--group-001.md#canonical-79194046c2f16815833ecf7958d7296d3ac656be586ef2e8fa3aa2aa7a47d4b6) |
| `disable` | [disable](resources--code_base_integration--reference--group-001.md#canonical-4f55a03c1fa6a92b9234459e68dc97555e00c064532b3a8b3b46c85c3c345157) |
| `id` | [id](resources--code_base_integration--reference--group-001.md#canonical-d236adfe3f606e79bcc5253abcc56ee1ada6687e6d992f725de736f317b45750) |
| `labels` | [labels](resources--code_base_integration--reference--group-001.md#canonical-620a3b4beebe118b630336e4e2beac9298e61ee9b57eec56e00ac29a6690a32d) |
| `name` | [name](resources--code_base_integration--reference--group-001.md#canonical-6be01bc1e6e6476b8be86aa94a3bc2c49e52d5ac7fbb94bc97943c49851a9df2) |
| `namespace` | [namespace](resources--code_base_integration--reference--group-001.md#canonical-115a5f3f9d84a70a1d054a12fc954d87639509d396e814655d0f024c3af873c6) |
| `timeouts` | [timeouts](resources--code_base_integration--reference--group-001.md#canonical-7c76f0ab9f8bde9fd78b7728379e134c275eeed89b5af0ece8f2a6773da6fd6a) |
| `timeouts.create` | [timeouts.create](resources--code_base_integration--reference--group-001.md#canonical-e3f5867a27cfa07f86a5abdbe173cc4422d24b217f7b487f74e3b595b7948e4f) |
| `timeouts.delete` | [timeouts.delete](resources--code_base_integration--reference--group-001.md#canonical-49cb18a1583f0713799f86b4b8664e9d4929d100b17557e7c42d4cd4c1f6992d) |
| `timeouts.read` | [timeouts.read](resources--code_base_integration--reference--group-001.md#canonical-0a85cb6596b15c09245f50b64bad4d4c0c5acf1bdc498245ac7b361665631dac) |
| `timeouts.update` | [timeouts.update](resources--code_base_integration--reference--group-001.md#canonical-73af432f6899e4c0545709aed3882b838c9388aad2d2060d14a03cae6769bb75) |

<a id="canonical-9a79f13d0f1a43bdeafdc916e738ab636a83c08fc03c4b70fee1f15f1888ab82"></a>

## Next pages — Property reference / afd93695f02e / 12

- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [timeouts](resources--code_base_integration--reference--group-001.md#canonical-41f3c8550b8529fda8a9906a2b929b165847756bf3daba96287289d43c9b3323)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-949dc7a257250177ecb6c027ed048c735491197ad28e3d0f95fbc752f7942c7f"></a>

## code_base_integration — code_base_integration / 08e5bf6bcff3 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- code_base_integration

<a id="canonical-7db312740ff9369e8db8dad8e0ec1ec538a79022ac2f2870cc40a0d9da22a419"></a>

Type: `"object"`. single nested block, Optional.

Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection
details.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("azure_repos",
    "bitbucket"),
  validators.ConflictingObjectAttributes("azure_repos",
    "bitbucket_server"),
  validators.ConflictingObjectAttributes("azure_repos",
    "github"),
  validators.ConflictingObjectAttributes("azure_repos",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("azure_repos",
    "gitlab"),
  validators.ConflictingObjectAttributes("azure_repos",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket",
    "bitbucket_server"),
  validators.ConflictingObjectAttributes("bitbucket",
    "github"),
  validators.ConflictingObjectAttributes("bitbucket",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket",
    "gitlab"),
  validators.ConflictingObjectAttributes("bitbucket",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "github"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "gitlab"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("github",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("github",
    "gitlab"),
  validators.ConflictingObjectAttributes("github",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("github_enterprise",
    "gitlab"),
  validators.ConflictingObjectAttributes("github_enterprise",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("gitlab",
    "gitlab_enterprise")}
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
  "x-ves-oneof-field-type": "[\"azure_repos\",\"bitbucket\",\"bitbucket_server\",\"github\",\"github_enterprise\",\"gitlab\",\"gitlab_enterprise\"]"
}
```

Terraform syntax:

```terraform
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-6fa58ec2025b99ffca5410275ffef909d40ff569715d4f8dcce618475d171950"></a>

## Direct properties — code_base_integration / 08e5bf6bcff3 / 3

- [azure_repos](resources--code_base_integration--reference--group-001.md#canonical-e115ef605ca30de04d06ea25dac653de5e3c1524583978d011e7789668a94390): complete subsection reference.

- [bitbucket](resources--code_base_integration--reference--group-001.md#canonical-fd278ed80965f4b0fd14f1f01d98bb83dd4b5a268c31e2f83fb0bc021f998615): complete subsection reference.

- [bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-5833b254fea77d976c9314d78cccf7938ee60fd10de9fe24f094780130bdf9ac): complete subsection reference.

- [github](resources--code_base_integration--reference--group-001.md#canonical-18a7423acc4f4bb00e71620d0f4732dc20b8e929976ae6e8d90211e04ab9a28b): complete subsection reference.

- [github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-6eec9228a70e4f0ca2d12e75923873899626927a628f38485a89369e3641f95c): complete subsection reference.

- [gitlab](resources--code_base_integration--reference--group-001.md#canonical-8980a81343c7a2f28adecaa191b8397741bf148b63bfca67b47eddf3d9f72930): complete subsection reference.

- [gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-b046134f2f65097a3b098926c89e7e3bbbeb67dc937a47559cfee09c68c5995b): complete subsection reference.

<a id="canonical-bcc055a8d71e92872b73c1858d4a45a328529ce5896dfdd7cc307de52ba6812f"></a>

## Next pages — code_base_integration / 08e5bf6bcff3 / 4

- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-e115ef605ca30de04d06ea25dac653de5e3c1524583978d011e7789668a94390)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-fd278ed80965f4b0fd14f1f01d98bb83dd4b5a268c31e2f83fb0bc021f998615)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-5833b254fea77d976c9314d78cccf7938ee60fd10de9fe24f094780130bdf9ac)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-18a7423acc4f4bb00e71620d0f4732dc20b8e929976ae6e8d90211e04ab9a28b)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-6eec9228a70e4f0ca2d12e75923873899626927a628f38485a89369e3641f95c)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-8980a81343c7a2f28adecaa191b8397741bf148b63bfca67b47eddf3d9f72930)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-b046134f2f65097a3b098926c89e7e3bbbeb67dc937a47559cfee09c68c5995b)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-e115ef605ca30de04d06ea25dac653de5e3c1524583978d011e7789668a94390"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdbd2b72d91e262b6ac451421bc917f11ea9d386bc7dcf1a929cd30bb476fdc5"></a>

## code_base_integration.azure_repos — code_base_integration.azure_repos / 0b493a630759 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- code_base_integration.azure_repos

<a id="canonical-00dc316a4def9f5e509f7fb12a438c18a588a6a49b4601e88f2bb1d931dd895b"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
azure_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-a8600034fa8997e0b357ab19ae54102e4e57c352f95b851e8827619ed9688711"></a>

## Direct properties — code_base_integration.azure_repos / 0b493a630759 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-6f0e7926fb1dd0eae98eeb89e4e09f09874b4f8c817cd9f22a7c7d80d090a124): complete subsection reference.

<a id="canonical-39e6b63cd6f0a3ee4aa7ac5fce626f9125d3c43c138c1309aadda1fd0d2f27ba"></a>

## Next pages — code_base_integration.azure_repos / 0b493a630759 / 4

- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-6f0e7926fb1dd0eae98eeb89e4e09f09874b4f8c817cd9f22a7c7d80d090a124)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-6f0e7926fb1dd0eae98eeb89e4e09f09874b4f8c817cd9f22a7c7d80d090a124"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b5fa4efb39101e68fc293a9767f1677cb4d58aad2cc827b1ef5aaba64de883a"></a>

## code_base_integration.azure_repos.access_token — code_base_integration.azure_repos.access_token / 063577ef9174 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-e115ef605ca30de04d06ea25dac653de5e3c1524583978d011e7789668a94390)
- code_base_integration.azure_repos.access_token

<a id="canonical-e316c0069ae90f69c111cc652f9520fe0a6aef336c6809ab28a25af23b682e44"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-2788b39c6282516572ae6f4861e5bf4a86e78d2324854cc64842f90e21e143f1"></a>

## Direct properties — code_base_integration.azure_repos.access_token / 063577ef9174 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-c27f54ba745240c5a67373ea4afabd8dfb92df6f1a4f607483262a8e6857dfdb): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-fb12bd0e7dbc11da226d5cbd7cdb82b9f6b4ef4191f76a7a667384af9610d354): complete subsection reference.

<a id="canonical-5e5b028e81727274c6cac6b2920f8912a509c5e25c696b5ce6c196560c67a6cb"></a>

## Next pages — code_base_integration.azure_repos.access_token / 063577ef9174 / 4

- [code_base_integration.azure_repos.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-c27f54ba745240c5a67373ea4afabd8dfb92df6f1a4f607483262a8e6857dfdb)
- [code_base_integration.azure_repos.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-fb12bd0e7dbc11da226d5cbd7cdb82b9f6b4ef4191f76a7a667384af9610d354)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-e115ef605ca30de04d06ea25dac653de5e3c1524583978d011e7789668a94390)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-c27f54ba745240c5a67373ea4afabd8dfb92df6f1a4f607483262a8e6857dfdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e5d326993d432d57347068cb5a8e6b661234ab34d3813ee4410b7f604ff5182"></a>

## code_base_integration.azure_repos.access_token.blindfold_secret_info — code_base_integration.azure_repos.access_token.blindfold_secret_info / 9dc939433154 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-e115ef605ca30de04d06ea25dac653de5e3c1524583978d011e7789668a94390)
- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-6f0e7926fb1dd0eae98eeb89e4e09f09874b4f8c817cd9f22a7c7d80d090a124)
- code_base_integration.azure_repos.access_token.blindfold_secret_info

<a id="canonical-7ed576f11995961ca2b16dbfa3566a4b2b5698b18681c005be7e605e8f3fb17c"></a>

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

<a id="canonical-d7c1f59792519c65f2e4ec6b28fa78a591cff8ccac278a2b68a3b6c5364a4bce"></a>

## Direct properties — code_base_integration.azure_repos.access_token.blindfold_secret_info / 9dc939433154 / 3

<a id="canonical-9ce84057be61bf9427cfe48ac78a404346f18f022377a41401d61fecbc183a6e"></a>

<a id="canonical-aa34db804f345562c1037ff92444b114acc6df6fca4f3c7d734fb1f03ad3df83"></a>

## decryption_provider property — code_base_integration.azure_repos.access_token.blindfold_secret_info / 9dc939433154 / 4

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

<a id="canonical-a50d09c74d90e59db639c16750fe25f4ff7720e33a3a4b2698033156f81a23a9"></a>

<a id="canonical-04dcfe5e897255bb164f13a3c2df55fb3b9e18845426fe37db874f0ac8e3f7ae"></a>

## location property — code_base_integration.azure_repos.access_token.blindfold_secret_info / 9dc939433154 / 5

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

<a id="canonical-800bffad0755ff94354ea30098907f51b8d643e77c5806c4bb120b24f5276dff"></a>

<a id="canonical-a6d103d94bb484c5adb7f800510de74c9a5e65ebe962a73de92d740ff5b1521d"></a>

## store_provider property — code_base_integration.azure_repos.access_token.blindfold_secret_info / 9dc939433154 / 6

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

<a id="canonical-c12a0b663fd284e8afda4c9128cbdc44ce4c19b3d42b5bf4e76caf296bacd1f4"></a>

## Next pages — code_base_integration.azure_repos.access_token.blindfold_secret_info / 9dc939433154 / 7

- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-6f0e7926fb1dd0eae98eeb89e4e09f09874b4f8c817cd9f22a7c7d80d090a124)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-fb12bd0e7dbc11da226d5cbd7cdb82b9f6b4ef4191f76a7a667384af9610d354"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8abeb46977236fc5074851075bcae70240be640eae92212248e4e56052e11621"></a>

## code_base_integration.azure_repos.access_token.clear_secret_info — code_base_integration.azure_repos.access_token.clear_secret_info / 186b692d69d3 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-e115ef605ca30de04d06ea25dac653de5e3c1524583978d011e7789668a94390)
- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-6f0e7926fb1dd0eae98eeb89e4e09f09874b4f8c817cd9f22a7c7d80d090a124)
- code_base_integration.azure_repos.access_token.clear_secret_info

<a id="canonical-338115c3d8fc70ad7ca3c35d97564d27388bfaa017761a83dbb189450c83fcf3"></a>

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

<a id="canonical-889340f2c5cf17f5108e2ec76bc6024e94cbfdabf2724a0746daa9f68ba01dfa"></a>

## Direct properties — code_base_integration.azure_repos.access_token.clear_secret_info / 186b692d69d3 / 3

<a id="canonical-3667fa3bf3b7fcae62fecf364770867f7044ba2e2db394f3194ac000620d5b05"></a>

<a id="canonical-cca6b90d49ae8664b577653428c5cb85b50171a99883bec0843b1158c50d002e"></a>

## provider_ref property — code_base_integration.azure_repos.access_token.clear_secret_info / 186b692d69d3 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-dc1a2d7f7b6a38f739e1d1eeb458064a13a964a979cca71d8dfc664b9f980688"></a>

<a id="canonical-27d1d90483a8c2e2beecb5a6cc717826113e1113b9e20dd5fe8ccde5fcfbbb6e"></a>

## url property — code_base_integration.azure_repos.access_token.clear_secret_info / 186b692d69d3 / 5

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

<a id="canonical-16c4bae392647fec1a39277f408a748b5e66e6fdcab9bb5a4a071bb23372e90d"></a>

## Next pages — code_base_integration.azure_repos.access_token.clear_secret_info / 186b692d69d3 / 6

- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-6f0e7926fb1dd0eae98eeb89e4e09f09874b4f8c817cd9f22a7c7d80d090a124)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-fd278ed80965f4b0fd14f1f01d98bb83dd4b5a268c31e2f83fb0bc021f998615"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfb85ad8f7e6651ad94abd64e052cc6a0328c8fa29244249a7c742988a193442"></a>

## code_base_integration.bitbucket — code_base_integration.bitbucket / 10a207b57aa4 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- code_base_integration.bitbucket

<a id="canonical-0a4d2fdcc867e8b511f68ee6aa50756c44f039d1a908f8a061d9c8e630da13b7"></a>

Type: `"object"`. single nested block, Optional.

BitBucket Cloud Integration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("username")}
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
bitbucket {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f7f306d9d9e345268426c663bc027e16d4c6bc9c77d13ece0965603219f7686"></a>

## Direct properties — code_base_integration.bitbucket / 10a207b57aa4 / 3

- [passwd](resources--code_base_integration--reference--group-001.md#canonical-ee4b40d098d07a539eb718c49b7dc36549444b7b511fb375b85680c06fd52f73): complete subsection reference.

<a id="canonical-3560b49c5099811cfa8a7be933964a5a45bc65aad65bd5f8946394dd07a7bd40"></a>

<a id="canonical-0848c0ba1741db8754fdae46b4101bebaef6b7a6f4edffe958155b810203e3aa"></a>

## username property — code_base_integration.bitbucket / 10a207b57aa4 / 4

Type: `"string"`. Optional.

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

<a id="canonical-0cfd089d8d560285153d9d392e0ae831d32b14f5c9c7df5b91c454f2aee2e6f2"></a>

## Next pages — code_base_integration.bitbucket / 10a207b57aa4 / 5

- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-ee4b40d098d07a539eb718c49b7dc36549444b7b511fb375b85680c06fd52f73)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-ee4b40d098d07a539eb718c49b7dc36549444b7b511fb375b85680c06fd52f73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aeda407ef185d45065f8dd1ce578fcda58388011e54a8a321e4da5e647469d9f"></a>

## code_base_integration.bitbucket.passwd — code_base_integration.bitbucket.passwd / 8b3e5fb63949 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-fd278ed80965f4b0fd14f1f01d98bb83dd4b5a268c31e2f83fb0bc021f998615)
- code_base_integration.bitbucket.passwd

<a id="canonical-0565c8421d5b18964a3baf81494d2974584db89bd37320bef33f928109424491"></a>

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
passwd {
  # Configure direct properties listed below.
}
```

<a id="canonical-a03e0d41ddf941cb37f1cf58a9290af0165437816a44aa9489221024fa330e3a"></a>

## Direct properties — code_base_integration.bitbucket.passwd / 8b3e5fb63949 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-99a58de2962e44b2e37cd834f92776c46fa914b90aeb05c125378edd0a4aa6f0): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-704a0a8702a755ddec506504f0f1554120a9a348127394876f0ed4b206429086): complete subsection reference.

<a id="canonical-a562c3261ddccf8984a26f13fdbf09a177e78e036807ec8b5d16e19c514dd641"></a>

## Next pages — code_base_integration.bitbucket.passwd / 8b3e5fb63949 / 4

- [code_base_integration.bitbucket.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-99a58de2962e44b2e37cd834f92776c46fa914b90aeb05c125378edd0a4aa6f0)
- [code_base_integration.bitbucket.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-704a0a8702a755ddec506504f0f1554120a9a348127394876f0ed4b206429086)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-fd278ed80965f4b0fd14f1f01d98bb83dd4b5a268c31e2f83fb0bc021f998615)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-99a58de2962e44b2e37cd834f92776c46fa914b90aeb05c125378edd0a4aa6f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cff6a5633da4501f10b515468429660820d4f25f378ba6a809677b19d62e1e49"></a>

## code_base_integration.bitbucket.passwd.blindfold_secret_info — code_base_integration.bitbucket.passwd.blindfold_secret_info / 1c872bd6587a / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-fd278ed80965f4b0fd14f1f01d98bb83dd4b5a268c31e2f83fb0bc021f998615)
- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-ee4b40d098d07a539eb718c49b7dc36549444b7b511fb375b85680c06fd52f73)
- code_base_integration.bitbucket.passwd.blindfold_secret_info

<a id="canonical-e00ea85babdd0e745c147ed63079d5c215142de358e1c9044e435323670a669e"></a>

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

<a id="canonical-a8ca20e3003dfdebc153f17763de99032b3f3d206726dab1017c613f9b059700"></a>

## Direct properties — code_base_integration.bitbucket.passwd.blindfold_secret_info / 1c872bd6587a / 3

<a id="canonical-1cf0c0b3f2e5cf20f59a1c8933805ba343f4b61084a126381f8d5ef7c7c6a3c4"></a>

<a id="canonical-ed10c4ce92a972ece281c1a21152f3f3fe8472cec8ab196feaef05423708048a"></a>

## decryption_provider property — code_base_integration.bitbucket.passwd.blindfold_secret_info / 1c872bd6587a / 4

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

<a id="canonical-8927d9f16365743f0d031e9bda91e721b6386739a53845a53652d5227b6be6e6"></a>

<a id="canonical-b24f7005e5ab23e14dcb42510e63c27919286e14ddc9b6ed31670c7446644367"></a>

## location property — code_base_integration.bitbucket.passwd.blindfold_secret_info / 1c872bd6587a / 5

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

<a id="canonical-9621adf671a7899327b4cae0c64e3e3e8737d189c9279b0224430d1b1f978645"></a>

<a id="canonical-a7773ca101ea63e0749097f85103845a6e9405382bd4a404a70ccb759518aa7e"></a>

## store_provider property — code_base_integration.bitbucket.passwd.blindfold_secret_info / 1c872bd6587a / 6

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

<a id="canonical-df07ca43caae2317ee296d35e97460686e1a953368c328f10c5272f1d60117f7"></a>

## Next pages — code_base_integration.bitbucket.passwd.blindfold_secret_info / 1c872bd6587a / 7

- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-ee4b40d098d07a539eb718c49b7dc36549444b7b511fb375b85680c06fd52f73)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-704a0a8702a755ddec506504f0f1554120a9a348127394876f0ed4b206429086"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d875fcd21b3aecc1134da257bfe26166fd1e36dd9d90177abca4bec9cddbc160"></a>

## code_base_integration.bitbucket.passwd.clear_secret_info — code_base_integration.bitbucket.passwd.clear_secret_info / 5b03b79c2e5b / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-fd278ed80965f4b0fd14f1f01d98bb83dd4b5a268c31e2f83fb0bc021f998615)
- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-ee4b40d098d07a539eb718c49b7dc36549444b7b511fb375b85680c06fd52f73)
- code_base_integration.bitbucket.passwd.clear_secret_info

<a id="canonical-deda4d307869154bca2627589a15f36d14a4a885d22e7fc4675fe079afac27dc"></a>

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

<a id="canonical-5d874dfcaf1151cd1c168f928f6b66671497326c874b46774c2564a857914db4"></a>

## Direct properties — code_base_integration.bitbucket.passwd.clear_secret_info / 5b03b79c2e5b / 3

<a id="canonical-50a38167d7b443c1569392210d610a987f5247569a5f08b2027aff171c969524"></a>

<a id="canonical-24d587fc6e52d685b4c2df3553622185ea7134baf2b82f0f4ae9889bd04ddba5"></a>

## provider_ref property — code_base_integration.bitbucket.passwd.clear_secret_info / 5b03b79c2e5b / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-818c7aeecb6809b08e2d09fff435592c06ff62e647ef341a1d7db86e4150c8d7"></a>

<a id="canonical-029c18bbd73f72aa44d8332f090ac1278575f0d25aac7d09e1003c5651862e12"></a>

## url property — code_base_integration.bitbucket.passwd.clear_secret_info / 5b03b79c2e5b / 5

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

<a id="canonical-ac7af26155b448cb64567e394aba57c710469fc24ac27b0104367a5f6c03a02f"></a>

## Next pages — code_base_integration.bitbucket.passwd.clear_secret_info / 5b03b79c2e5b / 6

- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-ee4b40d098d07a539eb718c49b7dc36549444b7b511fb375b85680c06fd52f73)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-5833b254fea77d976c9314d78cccf7938ee60fd10de9fe24f094780130bdf9ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7f81b80e993c2f20767f1c2bd74d2ea7881c0ebedf01dc47a768db39a00ab92"></a>

## code_base_integration.bitbucket_server — code_base_integration.bitbucket_server / 8572a0897e0b / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- code_base_integration.bitbucket_server

<a id="canonical-55aa224de5109ddcf551cc27909a84a97c96158c72a4f97c7f203b3477324901"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bitbucket server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url",
    "username")}
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
bitbucket_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-71f7b52834ff26bc3086e73aded9894c0667f43a98bb441431af3ed7929f296c"></a>

## Direct properties — code_base_integration.bitbucket_server / 8572a0897e0b / 3

- [passwd](resources--code_base_integration--reference--group-001.md#canonical-c21293e9deb5d9efa91842c9f59f741104f262de5d0c0b0fd6ea59f70e0bf030): complete subsection reference.

<a id="canonical-d547d59c022fb383c51e67c5ca2baa97e50401bd5dd5a9806b14d345a40816b3"></a>

<a id="canonical-f2773ec48268cb5825f9f5a0efd1bac0ae992db6bab003ecfe15fbf610a0dd0d"></a>

## url property — code_base_integration.bitbucket_server / 8572a0897e0b / 4

Type: `"string"`. Optional.

BitBucket Server URL. URL or URI reference

Upstream description:

URL or URI reference

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

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

<a id="canonical-fe40181c73ad3dc98772a52b6cbda3ab9c15fd2238662dd6f4ecc54395288fa5"></a>

<a id="canonical-1fd6d08514247cc0893745d4138aabd44aa670bf6488f8c8cd7aee005e64e840"></a>

## username property — code_base_integration.bitbucket_server / 8572a0897e0b / 5

Type: `"string"`. Optional.

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

<a id="canonical-64b6b30a8a0ab638c2b4170bdd0d688ef971a8a47abc65cc6f451c05ed66bc73"></a>

<a id="canonical-8f981ceeef07e240cd7475a3cfaefd21871a8f732a6ab8e66e0a28955c172112"></a>

## verify_ssl property — code_base_integration.bitbucket_server / 8572a0897e0b / 6

Type: `"bool"`. Optional.

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

<a id="canonical-d720cd6a6bd57ac146e2d34268fe9c4058cd2c7c0e469757e12584ba562903e6"></a>

## Next pages — code_base_integration.bitbucket_server / 8572a0897e0b / 7

- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-c21293e9deb5d9efa91842c9f59f741104f262de5d0c0b0fd6ea59f70e0bf030)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-c21293e9deb5d9efa91842c9f59f741104f262de5d0c0b0fd6ea59f70e0bf030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc75ccd1eb55d64a6637c61cb25ba1573663addaedde61577ed8e890cb4d1840"></a>

## code_base_integration.bitbucket_server.passwd — code_base_integration.bitbucket_server.passwd / 21e0f8556d51 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-5833b254fea77d976c9314d78cccf7938ee60fd10de9fe24f094780130bdf9ac)
- code_base_integration.bitbucket_server.passwd

<a id="canonical-10d857dfd660ce7c231778fe7d3f9e93f380aa6802de909d753800a3d4700816"></a>

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
passwd {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f5f0ac892e0ac054c001fc63e0db592b9c87839d8d5d033600e1f9eb5fb5c1a"></a>

## Direct properties — code_base_integration.bitbucket_server.passwd / 21e0f8556d51 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-672bf367a14c2c9d41b6a6845dc8735ad03b8bc1c06e2c79ddc4da24cbc53f8d): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-fd74c10c8cb1275490e40cc8b31862af08de710e08be671fb0e825c1abd49e79): complete subsection reference.

<a id="canonical-08022edddcd1e2ff8132e2f96009e46d728d2a72b9a51c5ff98b5250d7c0ebd3"></a>

## Next pages — code_base_integration.bitbucket_server.passwd / 21e0f8556d51 / 4

- [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-672bf367a14c2c9d41b6a6845dc8735ad03b8bc1c06e2c79ddc4da24cbc53f8d)
- [code_base_integration.bitbucket_server.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-fd74c10c8cb1275490e40cc8b31862af08de710e08be671fb0e825c1abd49e79)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-5833b254fea77d976c9314d78cccf7938ee60fd10de9fe24f094780130bdf9ac)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-672bf367a14c2c9d41b6a6845dc8735ad03b8bc1c06e2c79ddc4da24cbc53f8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9daa89ba542ecc8e808fb7360965921fe69ac7c6d87cf67bfd2cc91bf2ca23d"></a>

## code_base_integration.bitbucket_server.passwd.blindfold_secret_info — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / 1242fc042e70 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-5833b254fea77d976c9314d78cccf7938ee60fd10de9fe24f094780130bdf9ac)
- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-c21293e9deb5d9efa91842c9f59f741104f262de5d0c0b0fd6ea59f70e0bf030)
- code_base_integration.bitbucket_server.passwd.blindfold_secret_info

<a id="canonical-ab26b469dd65b02e3ca8190378da5e37365f07e08befe54cdfd4c3320ef350f1"></a>

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

<a id="canonical-f2fe273350e92f6c155916f34d3f2d2d7821694a5544069cbd5d23f1871170f8"></a>

## Direct properties — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / 1242fc042e70 / 3

<a id="canonical-6008068593c6dfe1563010e5df47d98cbfa7116ceb06fb3fc0e509f8311e3d41"></a>

<a id="canonical-ab409c9c2f16e7f93fa1d88c052977381219ca5c5eb489fcd8f7c2be21afb849"></a>

## decryption_provider property — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / 1242fc042e70 / 4

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

<a id="canonical-ca44b70186af98f14a7f105fc70d5690f0b30134d794b86a8bbdcd53d7adf767"></a>

<a id="canonical-df394bf6fc0089057ac25e0437dd5f46587bf8df056a88353a7a1c02428468da"></a>

## location property — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / 1242fc042e70 / 5

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

<a id="canonical-ac8fa94793a5dbd08ae3c973991363d103a13ed4b731efe31cf203167ed4d528"></a>

<a id="canonical-cc90e25187d1f4bd1596803480bfa0fb3ebe5343853fbfa98915b4d9d8c40a25"></a>

## store_provider property — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / 1242fc042e70 / 6

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

<a id="canonical-94464d85e05c454043e59846e18e1e01f523cbc16567a1c63298b85eaf27cb69"></a>

## Next pages — code_base_integration.bitbucket_server.passwd.blindfold_secret_info / 1242fc042e70 / 7

- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-c21293e9deb5d9efa91842c9f59f741104f262de5d0c0b0fd6ea59f70e0bf030)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-fd74c10c8cb1275490e40cc8b31862af08de710e08be671fb0e825c1abd49e79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-455a3d1ab5e228ab6006d5342cbb86cd039f440c373fae5ea7c047a31dc45b6d"></a>

## code_base_integration.bitbucket_server.passwd.clear_secret_info — code_base_integration.bitbucket_server.passwd.clear_secret_info / 3becddd744ff / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-5833b254fea77d976c9314d78cccf7938ee60fd10de9fe24f094780130bdf9ac)
- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-c21293e9deb5d9efa91842c9f59f741104f262de5d0c0b0fd6ea59f70e0bf030)
- code_base_integration.bitbucket_server.passwd.clear_secret_info

<a id="canonical-10d4ce9cfa01b64592a154b4153ece8b461a06fe964f4451042d312e87f965f3"></a>

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

<a id="canonical-8e6a128095338d5c1dca97bd08d544670eb642d850209511df988ee1ccabe80d"></a>

## Direct properties — code_base_integration.bitbucket_server.passwd.clear_secret_info / 3becddd744ff / 3

<a id="canonical-96d61d87da45340d2493675e4b258e72c314fd0f7e7afe1e5e18b3190153cccc"></a>

<a id="canonical-e7a53e8f950913a2010ee339b77645eae63bd8e6b83615be59f86e2d733b1a97"></a>

## provider_ref property — code_base_integration.bitbucket_server.passwd.clear_secret_info / 3becddd744ff / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9eeb2a5eb4af76353b351d6bf16a49cb96c74064315f8638a3755cecbdcf574b"></a>

<a id="canonical-22a78748da80e5b81978451719689c2970b74ef0920524be126633ec4c43cf9e"></a>

## url property — code_base_integration.bitbucket_server.passwd.clear_secret_info / 3becddd744ff / 5

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

<a id="canonical-e0c6e9fec22d7da10f9b9c50ebbf028d1d13202df66fe4e657cc10530d50128b"></a>

## Next pages — code_base_integration.bitbucket_server.passwd.clear_secret_info / 3becddd744ff / 6

- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-c21293e9deb5d9efa91842c9f59f741104f262de5d0c0b0fd6ea59f70e0bf030)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-18a7423acc4f4bb00e71620d0f4732dc20b8e929976ae6e8d90211e04ab9a28b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aa053982b0829407bf6e58696ab14bafe08a9c8b41e520db2ec942219e1c544"></a>

## code_base_integration.github — code_base_integration.github / 378a283d8265 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- code_base_integration.github

<a id="canonical-39612f9ca01618a33329739cd22396068e3cb51ec34ecea747d5821ee6186d01"></a>

Type: `"object"`. single nested block, Optional.

Github Integration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("username")}
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
github {
  # Configure direct properties listed below.
}
```

<a id="canonical-335bc57659b16ee33342ff329426b6fe38a2c95682693abb29e59f5fed26a402"></a>

## Direct properties — code_base_integration.github / 378a283d8265 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-176c2753d18c9302696fa64dfebd99690dabafec5ca401f83f804a04eb2daf98): complete subsection reference.

<a id="canonical-10342b34b089e207821da834ec64c665b5db1601327af776a78dc0910804ae2a"></a>

<a id="canonical-c6178ca04b31005f53ab3a2fe2fc22d55816b0a48267c1b4b8a23784ce518e2d"></a>

## username property — code_base_integration.github / 378a283d8265 / 4

Type: `"string"`. Optional.

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

<a id="canonical-38ea48401051647dffcfa62218379eb9f5dc38c7a75489db6c97fe6d0a99be7c"></a>

<a id="canonical-56ad8af5095f994201cc28394525a290e96842305e63887839536d3c516b878a"></a>

## verify_ssl property — code_base_integration.github / 378a283d8265 / 5

Type: `"bool"`. Optional.

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

<a id="canonical-d895bebbea7afcbb990336b44015ebfa1245049f220906818b74b69f58967a74"></a>

## Next pages — code_base_integration.github / 378a283d8265 / 6

- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-176c2753d18c9302696fa64dfebd99690dabafec5ca401f83f804a04eb2daf98)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-176c2753d18c9302696fa64dfebd99690dabafec5ca401f83f804a04eb2daf98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c724f5a5b0f7958e7e05762e740cf59ed134d7d516e86c158d82c646aeab1b73"></a>

## code_base_integration.github.access_token — code_base_integration.github.access_token / 9937c31b7cfb / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-18a7423acc4f4bb00e71620d0f4732dc20b8e929976ae6e8d90211e04ab9a28b)
- code_base_integration.github.access_token

<a id="canonical-1a4c71b32a390526d22d92e09a2f5921be8076592030e5d74a8e6419108d3a7a"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-f8c8511e35b6c701d3085d11b4aa94b97b4922e4044707c1b13a1fd4c0dfc30e"></a>

## Direct properties — code_base_integration.github.access_token / 9937c31b7cfb / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3560b7b1fb3908eb18e9a0ecbb1819824759f3f77017e7d9af2e7f23b5ff4cf6): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-77b92e65d02a5c5a880cebedaa98ad3d9a6560070653132fe1f705380207b74b): complete subsection reference.

<a id="canonical-f24c7ba0c1829a4407302c11adf5e24845088d78691ca05f8726e3b1c929619d"></a>

## Next pages — code_base_integration.github.access_token / 9937c31b7cfb / 4

- [code_base_integration.github.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3560b7b1fb3908eb18e9a0ecbb1819824759f3f77017e7d9af2e7f23b5ff4cf6)
- [code_base_integration.github.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-77b92e65d02a5c5a880cebedaa98ad3d9a6560070653132fe1f705380207b74b)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-18a7423acc4f4bb00e71620d0f4732dc20b8e929976ae6e8d90211e04ab9a28b)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-3560b7b1fb3908eb18e9a0ecbb1819824759f3f77017e7d9af2e7f23b5ff4cf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3a7399a8caa58ecc796dc5b040cf76f3230344fd42dec2555e0d318a63d24fc"></a>

## code_base_integration.github.access_token.blindfold_secret_info — code_base_integration.github.access_token.blindfold_secret_info / 16865fec258f / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-18a7423acc4f4bb00e71620d0f4732dc20b8e929976ae6e8d90211e04ab9a28b)
- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-176c2753d18c9302696fa64dfebd99690dabafec5ca401f83f804a04eb2daf98)
- code_base_integration.github.access_token.blindfold_secret_info

<a id="canonical-2c9148543b7af6aa9df2295f43c9c62d7df31d0ff8f6047bb543527e6486f36a"></a>

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

<a id="canonical-f9bcde6c1575456c193d0cdb25c80cb0e098ccde9b8ed0f5f267caa52ac2f79f"></a>

## Direct properties — code_base_integration.github.access_token.blindfold_secret_info / 16865fec258f / 3

<a id="canonical-679de269aa5c59046e91e923c1b38184a6f52a5ae4bae5e346dbb5494a784b37"></a>

<a id="canonical-34289b628512956aeeb6365e6ea07510ee74c0d74dfa85f53b6a9f541371156a"></a>

## decryption_provider property — code_base_integration.github.access_token.blindfold_secret_info / 16865fec258f / 4

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

<a id="canonical-bb58cf548f7d75a5393b8f4825d93f18e9a48743ab59eea60fa331a2a47ea78b"></a>

<a id="canonical-8560ea29eebd66947d485b0d3308596f8ee9b36e9d115bd53fd7e8d14fbecacc"></a>

## location property — code_base_integration.github.access_token.blindfold_secret_info / 16865fec258f / 5

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

<a id="canonical-a5fc453bdce912e01d3cd5a1572346059271c0d4392639c293af97655ca4a3bd"></a>

<a id="canonical-d90f86e47ea938911239b04c33bde800641ded48a1a5069354c92af3f9f59ee7"></a>

## store_provider property — code_base_integration.github.access_token.blindfold_secret_info / 16865fec258f / 6

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

<a id="canonical-7994f202f2bdb3bb85c5d97b3b21a3ac67b239ae79a007c8e32406780a3a8ea9"></a>

## Next pages — code_base_integration.github.access_token.blindfold_secret_info / 16865fec258f / 7

- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-176c2753d18c9302696fa64dfebd99690dabafec5ca401f83f804a04eb2daf98)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-77b92e65d02a5c5a880cebedaa98ad3d9a6560070653132fe1f705380207b74b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbafdc5e528ca19095c3e358877f416f4cff525fc5eeeb98e95a198647790976"></a>

## code_base_integration.github.access_token.clear_secret_info — code_base_integration.github.access_token.clear_secret_info / 64943f696cc0 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-18a7423acc4f4bb00e71620d0f4732dc20b8e929976ae6e8d90211e04ab9a28b)
- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-176c2753d18c9302696fa64dfebd99690dabafec5ca401f83f804a04eb2daf98)
- code_base_integration.github.access_token.clear_secret_info

<a id="canonical-2f36955064d76eb5dac106380eab1b3f0f713d8ea266b6a7d272d21d29a0cfe5"></a>

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

<a id="canonical-4043eed4d9eed50adeb6327d8977a3aa7a4d64a54a8a08829017549b3e940040"></a>

## Direct properties — code_base_integration.github.access_token.clear_secret_info / 64943f696cc0 / 3

<a id="canonical-a37bfa1e70ca12dfc236f95d635f3467ad7472250afe1ab78ce90951414070f5"></a>

<a id="canonical-22c5888f4079fb6e45185449d3704a599d131b9d7743a362e8584df334c8030f"></a>

## provider_ref property — code_base_integration.github.access_token.clear_secret_info / 64943f696cc0 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-584333d773a93d10c1c8f1b2f92768aebfb723e570b62fc80b335ddb82e53a9c"></a>

<a id="canonical-3b780972b8cf6b8217f946759f264c0e43ed3741ea9c83437dd2ee3e1a8b0b55"></a>

## url property — code_base_integration.github.access_token.clear_secret_info / 64943f696cc0 / 5

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

<a id="canonical-1b186cc01a8b490eafcd216c968f990f9a5a9066a0be3710ca9998e47ff8d990"></a>

## Next pages — code_base_integration.github.access_token.clear_secret_info / 64943f696cc0 / 6

- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-176c2753d18c9302696fa64dfebd99690dabafec5ca401f83f804a04eb2daf98)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-6eec9228a70e4f0ca2d12e75923873899626927a628f38485a89369e3641f95c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-547269eadd12629b5267b8efae3536064ebc963961c06117041141fa31905180"></a>

## code_base_integration.github_enterprise — code_base_integration.github_enterprise / 089e8b99c7af / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- code_base_integration.github_enterprise

<a id="canonical-f7139c43615150249f93ede66a323e27e220d3aec3a9d78776b0eebfb4ca0ced"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for github enterprise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hostname",
    "username")}
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
github_enterprise {
  # Configure direct properties listed below.
}
```

<a id="canonical-808b059ca4bea897af7126f47473ee6e480eb090ba4fc4a07026e7d85b40104b"></a>

## Direct properties — code_base_integration.github_enterprise / 089e8b99c7af / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-5054f12db59267eb11dbae78fb8f5485f184c96ec1304578be2b3cb5fc63e7a3): complete subsection reference.

<a id="canonical-b277d28a4f12af2e8a44e4b406dec22fe0c46d8f9ea0d92c238ef20ae372562d"></a>

<a id="canonical-84400f04f57c9befb7f7458007cd33f3597db2c330e36c5ed4e8d0b0ff7396ea"></a>

## hostname property — code_base_integration.github_enterprise / 089e8b99c7af / 4

Type: `"string"`. Optional.

GitHub Hostname. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-43e3a25796f80c67276d7a0bc34702760e8f88a848058aa373f8b4f8f98cd415"></a>

<a id="canonical-b0ac9b4dffac2bc5eebf6838ce03b91f067cdd32b3f0656de2cb1f6ff4c2ad1d"></a>

## username property — code_base_integration.github_enterprise / 089e8b99c7af / 5

Type: `"string"`. Optional.

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

<a id="canonical-17c5b927e859b56a7928f993d2c2202cc807fdd810738f67bb6e7610802626ba"></a>

## Next pages — code_base_integration.github_enterprise / 089e8b99c7af / 6

- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-5054f12db59267eb11dbae78fb8f5485f184c96ec1304578be2b3cb5fc63e7a3)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-5054f12db59267eb11dbae78fb8f5485f184c96ec1304578be2b3cb5fc63e7a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-202b67fd0838a43a6522ed76104a861c1cd5849aa8c15f8ebb268a59973ef6d2"></a>

## code_base_integration.github_enterprise.access_token — code_base_integration.github_enterprise.access_token / 9697c289adc1 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-6eec9228a70e4f0ca2d12e75923873899626927a628f38485a89369e3641f95c)
- code_base_integration.github_enterprise.access_token

<a id="canonical-da0efb41ae3d93c98028c6221f9151828b7b570e12b861a8acf4712f97a44804"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-44620649e55069219873e12d4ad356a7a68529498387283145820273af726417"></a>

## Direct properties — code_base_integration.github_enterprise.access_token / 9697c289adc1 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-575e9429a75e61dcc6067ed51c55435c9badab4798b815e11f1660d2009f9dea): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-86351ada7d39dccf91065cc4f71ab97bb5f8709bc1bd011430b934a015fb5b9e): complete subsection reference.

<a id="canonical-197d015f4f2f7c1b13b6dbddd2fe78e706197858484248039178e82414a3c4d2"></a>

## Next pages — code_base_integration.github_enterprise.access_token / 9697c289adc1 / 4

- [code_base_integration.github_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-575e9429a75e61dcc6067ed51c55435c9badab4798b815e11f1660d2009f9dea)
- [code_base_integration.github_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-86351ada7d39dccf91065cc4f71ab97bb5f8709bc1bd011430b934a015fb5b9e)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-6eec9228a70e4f0ca2d12e75923873899626927a628f38485a89369e3641f95c)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-575e9429a75e61dcc6067ed51c55435c9badab4798b815e11f1660d2009f9dea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b159e513c7321167b042da5d89d671fdacccd243978b5d75d6dba88a6baac30c"></a>

## code_base_integration.github_enterprise.access_token.blindfold_secret_info — code_base_integration.github_enterprise.access_token.blindfold_secret_info / 863b605b74d2 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-6eec9228a70e4f0ca2d12e75923873899626927a628f38485a89369e3641f95c)
- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-5054f12db59267eb11dbae78fb8f5485f184c96ec1304578be2b3cb5fc63e7a3)
- code_base_integration.github_enterprise.access_token.blindfold_secret_info

<a id="canonical-787412a7621de2571985237ede4fc47be231eaec9090aefa46675e2db16a6059"></a>

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

<a id="canonical-792a2c80eb15469298e2bd843401df792d86a8ddd31955fc72e4e14f696e18e5"></a>

## Direct properties — code_base_integration.github_enterprise.access_token.blindfold_secret_info / 863b605b74d2 / 3

<a id="canonical-b15c899ada42d65f0855715ba1844d190e91a25785bcdf1ec19f40e77fa1b75e"></a>

<a id="canonical-46aa89382b87c222afdcace8110ac01a4de79bbb3e37625d9269c9770c412274"></a>

## decryption_provider property — code_base_integration.github_enterprise.access_token.blindfold_secret_info / 863b605b74d2 / 4

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

<a id="canonical-f1bc8c0ff5f2114692446983f83364a1030b7078be344aa5fd7478d3c78dc1a8"></a>

<a id="canonical-8d616a4de55fb2e8349d55cc9ba841fc6b7d8568ed15c37f100971ee8d4da067"></a>

## location property — code_base_integration.github_enterprise.access_token.blindfold_secret_info / 863b605b74d2 / 5

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

<a id="canonical-8e900b8b3bfcb75e7efc0e48e946d79f3193fefc942de1a2f831de26ce1f99f9"></a>

<a id="canonical-771d540238f8826eb0e8bfc4a83baf2c59456d2d96377bbaef9defff1850730a"></a>

## store_provider property — code_base_integration.github_enterprise.access_token.blindfold_secret_info / 863b605b74d2 / 6

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

<a id="canonical-50b780f84c85cee080c694cf2fe3f825af781e7111418c45086397ace66a1326"></a>

## Next pages — code_base_integration.github_enterprise.access_token.blindfold_secret_info / 863b605b74d2 / 7

- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-5054f12db59267eb11dbae78fb8f5485f184c96ec1304578be2b3cb5fc63e7a3)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-86351ada7d39dccf91065cc4f71ab97bb5f8709bc1bd011430b934a015fb5b9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b240f460935c9bd34f4f3e82ebc5d7e1ca7bafff7a18f40aaf90a16ebdeb62d4"></a>

## code_base_integration.github_enterprise.access_token.clear_secret_info — code_base_integration.github_enterprise.access_token.clear_secret_info / a8e39bce5c11 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-6eec9228a70e4f0ca2d12e75923873899626927a628f38485a89369e3641f95c)
- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-5054f12db59267eb11dbae78fb8f5485f184c96ec1304578be2b3cb5fc63e7a3)
- code_base_integration.github_enterprise.access_token.clear_secret_info

<a id="canonical-adfa34f99f6695a16d6f6770600d0ee663d06579d6fc37a5970dfd2789b9abca"></a>

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

<a id="canonical-0cc51a993453ca90f28dad9bf2d5bbb1b6eda195fd15915dc8be4ce9e9ff330b"></a>

## Direct properties — code_base_integration.github_enterprise.access_token.clear_secret_info / a8e39bce5c11 / 3

<a id="canonical-b83db40ec772dfb7c8db482f6b943779728fb4b6c51d6fbd4b5b2bd10346757a"></a>

<a id="canonical-2550566683aba72a2eee7dbdf7303491faf917229414e1c8469c818062490c37"></a>

## provider_ref property — code_base_integration.github_enterprise.access_token.clear_secret_info / a8e39bce5c11 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-10ece9c8106e87e3886303df8d0875a0ab76897aa1eba5d19a2f02d6ba2ec304"></a>

<a id="canonical-a3eb6c99b3eb434368243c28eb2c73baec8f27a54fac4bbb1fa5324ddaece7a9"></a>

## url property — code_base_integration.github_enterprise.access_token.clear_secret_info / a8e39bce5c11 / 5

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

<a id="canonical-f3f3c8f633b8fde2ff9cf1038e9b4b987ee0ad13aeea850493d377b26954ea72"></a>

## Next pages — code_base_integration.github_enterprise.access_token.clear_secret_info / a8e39bce5c11 / 6

- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-5054f12db59267eb11dbae78fb8f5485f184c96ec1304578be2b3cb5fc63e7a3)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-8980a81343c7a2f28adecaa191b8397741bf148b63bfca67b47eddf3d9f72930"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81d04b7eccbc2f2d5a49a36b6d25ef2f54d7fe9c2f090c2dab0a41e89f7c7372"></a>

## code_base_integration.gitlab — code_base_integration.gitlab / e54624bac921 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- code_base_integration.gitlab

<a id="canonical-507315619a769bc7e946ac193220669528e070d0491eb5a25860181d9b6f278a"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
gitlab {
  # Configure direct properties listed below.
}
```

<a id="canonical-ebdb3eabc21cf60ad66f2c51ff80482c57a8142da16c22948e77b1a8dab946be"></a>

## Direct properties — code_base_integration.gitlab / e54624bac921 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-055126b6147efa5712fd297571d33a5b43ddb0cec66d49b66d5c7f5def8ff12f): complete subsection reference.

<a id="canonical-e1bc619241d2e3dd3673c2c72b4241fc1198c18d925d038b8ea3efab6eacd983"></a>

## Next pages — code_base_integration.gitlab / e54624bac921 / 4

- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-055126b6147efa5712fd297571d33a5b43ddb0cec66d49b66d5c7f5def8ff12f)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-055126b6147efa5712fd297571d33a5b43ddb0cec66d49b66d5c7f5def8ff12f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bf645c3a2885bd6d821d56e22abd1c8d656383504bdd699e4ac45804131022b"></a>

## code_base_integration.gitlab.access_token — code_base_integration.gitlab.access_token / 889dc2458d51 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-8980a81343c7a2f28adecaa191b8397741bf148b63bfca67b47eddf3d9f72930)
- code_base_integration.gitlab.access_token

<a id="canonical-6d3233cca8e47e75a50c83137d9e68627e53b3b8c78fce6aa353992bb694462d"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6661774676ef104151b3e97a4162cbe5de08b1fb20b7738ddb9a3a0b1fd932c"></a>

## Direct properties — code_base_integration.gitlab.access_token / 889dc2458d51 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-5bc3d626efda5f5eb6b5d2cc79fa884a797309a3d9335b71486e477ae7002202): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-abb2eb42ce2a78099eb55f323864f5a1cfc02d7098f7621209e0414d8611807f): complete subsection reference.

<a id="canonical-83c0d95c7a120807feb52af1f4ba25273c8290a28d4c56a4ae59b55d8469b73f"></a>

## Next pages — code_base_integration.gitlab.access_token / 889dc2458d51 / 4

- [code_base_integration.gitlab.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-5bc3d626efda5f5eb6b5d2cc79fa884a797309a3d9335b71486e477ae7002202)
- [code_base_integration.gitlab.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-abb2eb42ce2a78099eb55f323864f5a1cfc02d7098f7621209e0414d8611807f)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-8980a81343c7a2f28adecaa191b8397741bf148b63bfca67b47eddf3d9f72930)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-5bc3d626efda5f5eb6b5d2cc79fa884a797309a3d9335b71486e477ae7002202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1647ab72d83cfbf7920506bbd55ce1b1eaf38183414edc5c79af80f1fcbf6bca"></a>

## code_base_integration.gitlab.access_token.blindfold_secret_info — code_base_integration.gitlab.access_token.blindfold_secret_info / 665b636b7d3d / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-8980a81343c7a2f28adecaa191b8397741bf148b63bfca67b47eddf3d9f72930)
- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-055126b6147efa5712fd297571d33a5b43ddb0cec66d49b66d5c7f5def8ff12f)
- code_base_integration.gitlab.access_token.blindfold_secret_info

<a id="canonical-33a5b3cd2b27e9ffd84ba649d6df7f56538aeec3998e8f972f1f1650c2ac5ccb"></a>

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

<a id="canonical-605ea16fdae346786c1c133d70bfa0a4dbb24f09fd299dbf8b889e2b34f7b66e"></a>

## Direct properties — code_base_integration.gitlab.access_token.blindfold_secret_info / 665b636b7d3d / 3

<a id="canonical-adad1d06b9bfd2212a73040d624f23880e9dc56b52b5ee272d5ad148f7dc20ad"></a>

<a id="canonical-746a326187b8c56fcbe604af0b6395c1e1ac56944af2add1a8473b81697025b9"></a>

## decryption_provider property — code_base_integration.gitlab.access_token.blindfold_secret_info / 665b636b7d3d / 4

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

<a id="canonical-419d8aff7d5caa6882f539a0dfe59828a807656015f0cfbdacd4707f4ac3ee9e"></a>

<a id="canonical-3164f9c369d98ae1bdb4621dd3d4c1f5f93423dd65096927569e0dc3be292e49"></a>

## location property — code_base_integration.gitlab.access_token.blindfold_secret_info / 665b636b7d3d / 5

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

<a id="canonical-3950009eb6e46ea186a7ebb271decba110a7bdf6028825a1f101edfaf442384a"></a>

<a id="canonical-d9eb18ab26f4814e0f89c68cdbc5bc87656d12e92a4d26e92084f791b933e074"></a>

## store_provider property — code_base_integration.gitlab.access_token.blindfold_secret_info / 665b636b7d3d / 6

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

<a id="canonical-468dfe39ef54f916bf21c043a9bfe884b7c71825c32d4838f8169be8c40a393f"></a>

## Next pages — code_base_integration.gitlab.access_token.blindfold_secret_info / 665b636b7d3d / 7

- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-055126b6147efa5712fd297571d33a5b43ddb0cec66d49b66d5c7f5def8ff12f)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-abb2eb42ce2a78099eb55f323864f5a1cfc02d7098f7621209e0414d8611807f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46aecb748812846122a55f21706baa5a3ebed043cbbaf843eae4b07ce422704a"></a>

## code_base_integration.gitlab.access_token.clear_secret_info — code_base_integration.gitlab.access_token.clear_secret_info / a755fcbd3919 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-8980a81343c7a2f28adecaa191b8397741bf148b63bfca67b47eddf3d9f72930)
- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-055126b6147efa5712fd297571d33a5b43ddb0cec66d49b66d5c7f5def8ff12f)
- code_base_integration.gitlab.access_token.clear_secret_info

<a id="canonical-250afca179485e4c9d63761f57af3c520a58ad2213587be34e553133d4c2bac6"></a>

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

<a id="canonical-5dbcb0ee1790092e2992c38e786c9b9c6b41fe9984ce9886a0fecd77ce45dfdb"></a>

## Direct properties — code_base_integration.gitlab.access_token.clear_secret_info / a755fcbd3919 / 3

<a id="canonical-247e13a062dbdf0b1f70c3613342c76879a34bfaa15d27ac076728b26c5f009a"></a>

<a id="canonical-c0aaeec843c6fd4bba522c2794bb7c401d66e6acee4dee2576155b455d86dd87"></a>

## provider_ref property — code_base_integration.gitlab.access_token.clear_secret_info / a755fcbd3919 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-358c7913d7bb3ff8dc4fa96f76b52117c87bbc96b0218a7b5ba676b10bcd914c"></a>

<a id="canonical-55105d4b2adb7ee0a731c094936ba85c4386e9a4dc2fd56a56adb996a6da7e3e"></a>

## url property — code_base_integration.gitlab.access_token.clear_secret_info / a755fcbd3919 / 5

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

<a id="canonical-2e0e44e228f07df76b63d8e79ade1d21e87a2074aee40c099d574f6808bcc4bb"></a>

## Next pages — code_base_integration.gitlab.access_token.clear_secret_info / a755fcbd3919 / 6

- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-055126b6147efa5712fd297571d33a5b43ddb0cec66d49b66d5c7f5def8ff12f)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-b046134f2f65097a3b098926c89e7e3bbbeb67dc937a47559cfee09c68c5995b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05d8337ba1bb53938e6be7c759cfc99f925ecf75fb63978b46bd4f869465097a"></a>

## code_base_integration.gitlab_enterprise — code_base_integration.gitlab_enterprise / a10a0968924a / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- code_base_integration.gitlab_enterprise

<a id="canonical-fe1b4b49c1e5f0423f59e1ca64edb9ffc864c48141b562d5c8648058affb807d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for gitlab enterprise.

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
gitlab_enterprise {
  # Configure direct properties listed below.
}
```

<a id="canonical-e9c6eb676b41dd1ae4b1825761114a20d8280db0ae80832b65cef7a90412d94b"></a>

## Direct properties — code_base_integration.gitlab_enterprise / a10a0968924a / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-e63991fa16331693fc1004e0474147afba2b3d9b6cd66faeeac46cfe49e0f34e): complete subsection reference.

<a id="canonical-2f440c94c2db2d97d70322b5dd8678ce03fc3fee13d2b85901acbcd116587baa"></a>

<a id="canonical-58ba2fba157fd9f90de5caf1949486d9f6199685bc4909f4fcb170fa5a4462de"></a>

## url property — code_base_integration.gitlab_enterprise / a10a0968924a / 4

Type: `"string"`. Optional.

GitLab URL. URL or URI reference

Upstream description:

URL or URI reference

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

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

<a id="canonical-d344bc93c2f84e40104e8586291a5c3ec9a1f09be3f69ff3a22cc4954780e48f"></a>

## Next pages — code_base_integration.gitlab_enterprise / a10a0968924a / 5

- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-e63991fa16331693fc1004e0474147afba2b3d9b6cd66faeeac46cfe49e0f34e)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-e63991fa16331693fc1004e0474147afba2b3d9b6cd66faeeac46cfe49e0f34e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90da9271017c4b76b06f00a7859dd8e63c359b9ec376ae4ceef04e855369093e"></a>

## code_base_integration.gitlab_enterprise.access_token — code_base_integration.gitlab_enterprise.access_token / 6e540299e950 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-b046134f2f65097a3b098926c89e7e3bbbeb67dc937a47559cfee09c68c5995b)
- code_base_integration.gitlab_enterprise.access_token

<a id="canonical-2fe93cc4b13b712c290840055a748c25a8d975380149a1cb4ef07db231975624"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-1df2a0723a4d87d5ba0917ab10f47d9b8b131dcf6719427d7ea1b39f16327998"></a>

## Direct properties — code_base_integration.gitlab_enterprise.access_token / 6e540299e950 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-8e6a94d8f85fc4bd918332ae5eeeb44642f7dacea9e588d30d64fbc11b442b65): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-aa834536de3916de66253b33c29353b2d3f548b5c25930a97499a1f1ca043339): complete subsection reference.

<a id="canonical-dcf19c5737bab05578e1d94e0be1615efb1e8e08ae7231bdd654ff9d63f59c00"></a>

## Next pages — code_base_integration.gitlab_enterprise.access_token / 6e540299e950 / 4

- [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-8e6a94d8f85fc4bd918332ae5eeeb44642f7dacea9e588d30d64fbc11b442b65)
- [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-aa834536de3916de66253b33c29353b2d3f548b5c25930a97499a1f1ca043339)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-b046134f2f65097a3b098926c89e7e3bbbeb67dc937a47559cfee09c68c5995b)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-8e6a94d8f85fc4bd918332ae5eeeb44642f7dacea9e588d30d64fbc11b442b65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a64f97925734760c88c671fb7d905a96fda38dfe4980b33898dadea535b4572f"></a>

## code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / ed11eb2216a5 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-b046134f2f65097a3b098926c89e7e3bbbeb67dc937a47559cfee09c68c5995b)
- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-e63991fa16331693fc1004e0474147afba2b3d9b6cd66faeeac46cfe49e0f34e)
- code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info

<a id="canonical-0bc8cce0e2ee6f29537342ae7eb38a68ff92cfdd18b288a6b957645bd307fe21"></a>

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

<a id="canonical-c24e55cafae862122e7b9d935944ce1b0930c7814a7842cb1c4e5ab0e2fafdf7"></a>

## Direct properties — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / ed11eb2216a5 / 3

<a id="canonical-17cbc7dfcf4f653891b548fbdfaaca694dd89ace6b77e640b1330723b52cf02a"></a>

<a id="canonical-459d2a8ca0e715bbb66e7cc812fbe5d0a71024643a59e742fd3450df78dbc7ad"></a>

## decryption_provider property — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / ed11eb2216a5 / 4

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

<a id="canonical-cc91b29f58b478ccedb5528e936e924a239747c9f44dd8a79bb45f7dc3ba82db"></a>

<a id="canonical-cf40a3d566785f94b28da8fea8cfe745bd01b54431299c729ae23007a86db14e"></a>

## location property — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / ed11eb2216a5 / 5

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

<a id="canonical-31e7fb244cb7fbc6be14b1ed8db4191ff12ef14a3cbc7256dda8d302afa4cf6c"></a>

<a id="canonical-89342fe0c3ae2c82f93ed8cd5564978cc64c4cb4d89d6c54d0512885e5db0e54"></a>

## store_provider property — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / ed11eb2216a5 / 6

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

<a id="canonical-5ad5dff34451b17b31dceb35887b16b795d1dfd500c2c4a4477035fa05e47694"></a>

## Next pages — code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info / ed11eb2216a5 / 7

- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-e63991fa16331693fc1004e0474147afba2b3d9b6cd66faeeac46cfe49e0f34e)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-aa834536de3916de66253b33c29353b2d3f548b5c25930a97499a1f1ca043339"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62472853d228bde89b3df413df15245e9db203c0d5409f136e64bb004f25d969"></a>

## code_base_integration.gitlab_enterprise.access_token.clear_secret_info — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 32d7282aecab / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-d631aa7c121a502204d99da03e5856b4ebbb6ceba91bae203fc7e1f5d6ea8ac1)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-b046134f2f65097a3b098926c89e7e3bbbeb67dc937a47559cfee09c68c5995b)
- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-e63991fa16331693fc1004e0474147afba2b3d9b6cd66faeeac46cfe49e0f34e)
- code_base_integration.gitlab_enterprise.access_token.clear_secret_info

<a id="canonical-d67fe49b5d14290e668e3fa6a414e608a0d8e6a70cd174a56ab09aedc8fe1d88"></a>

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

<a id="canonical-857b7f243c9132ad32a68fee62442de9bbefcf54c61dde4f943ae5c74fe4a691"></a>

## Direct properties — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 32d7282aecab / 3

<a id="canonical-5ce0392955d93dc943965336e80116d1e6798b11958e2dc4bea810aa101422bb"></a>

<a id="canonical-93607706bea7e58fcac221f738b8734dba21275b6edfcdd43be45d1a573ca3d8"></a>

## provider_ref property — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 32d7282aecab / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-107a41ebf8caa57a7eaf0cb33c99c6266d129eb6dee77b6e878316f4592832f9"></a>

<a id="canonical-770cfde33bf04764aa10ff3d379503f79c79f4bb913a9c77740903a40bdf5bdb"></a>

## url property — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 32d7282aecab / 5

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

<a id="canonical-b43406a1913d75359c9ea2cb9e353a84b0cd3c9a3e47af171dde8eaff70fffc1"></a>

## Next pages — code_base_integration.gitlab_enterprise.access_token.clear_secret_info / 32d7282aecab / 6

- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-e63991fa16331693fc1004e0474147afba2b3d9b6cd66faeeac46cfe49e0f34e)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-41f3c8550b8529fda8a9906a2b929b165847756bf3daba96287289d43c9b3323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f45e78a10a2ba189cc790284dd04371d6773afc0a46f1c7df7807b68d4e7a5c"></a>

## timeouts — timeouts / 05d85f69b92c / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- timeouts

<a id="canonical-7c76f0ab9f8bde9fd78b7728379e134c275eeed89b5af0ece8f2a6773da6fd6a"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-56c8a84ba90299ceefd2368503e9f7315f8a66a7a93beedd442c86b7cd937029"></a>

## Direct properties — timeouts / 05d85f69b92c / 3

<a id="canonical-e3f5867a27cfa07f86a5abdbe173cc4422d24b217f7b487f74e3b595b7948e4f"></a>

<a id="canonical-fc9da1db3372655d9221c0d1cac738a333a9247d913e35ac7128d63ae9beb2d1"></a>

## create property — timeouts / 05d85f69b92c / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-49cb18a1583f0713799f86b4b8664e9d4929d100b17557e7c42d4cd4c1f6992d"></a>

<a id="canonical-e8ab242e814f355dbe1d6897a7453b3d9dbd7564d255e330161cf024bd0b9941"></a>

## delete property — timeouts / 05d85f69b92c / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0a85cb6596b15c09245f50b64bad4d4c0c5acf1bdc498245ac7b361665631dac"></a>

<a id="canonical-2bb142bfe593c6f3072e648bf7a98f2ee00b864f2d734ff6dc88c6ff16855cb3"></a>

## read property — timeouts / 05d85f69b92c / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-73af432f6899e4c0545709aed3882b838c9388aad2d2060d14a03cae6769bb75"></a>

<a id="canonical-ab7939f68a9e23f99bd70e7c8edb2732a952ce8d19fc23af8518bb0c0b31943d"></a>

## update property — timeouts / 05d85f69b92c / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-b45db0405d483f05d30abcd72c13216f97c5e1a662ce7c2ea0629f114ec900b2"></a>

## Next pages — timeouts / 05d85f69b92c / 8

- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-e196d20a734aa90cd6179e647dcb888644b3ece7c1511b6aa1ea4fe021c002e1)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
