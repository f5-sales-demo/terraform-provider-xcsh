---
page_title: "xcsh_k8s_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster reference."
---

# xcsh_k8s_cluster reference

<a id="canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c5abeaf2d96d0f8c80a8e047ef5f7725cdad49e0dd36a352fda9c508d9d8a31"></a>

## Property reference — Property reference / 3f99d6be32dc / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- Property reference

<a id="canonical-f6c2d7b2140037b994bdb438edbf5fdfcd2de6c94667b11f4ccc876f85ff8265"></a>

## Direct properties — Property reference / 3f99d6be32dc / 3

<a id="canonical-0f9521d8c92b9c13ab38b1e811f73d500b650ae4b0760f3f3d5dadc7684645f5"></a>

<a id="canonical-ad540ce25df54cb7699b0c6afeaf731bb48782217dbd8a5d8f3bf5412fe7e912"></a>

## annotations property — Property reference / 3f99d6be32dc / 4

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

- [cluster_scoped_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-ed8a65af165a6a1d818048346ad28b12d838f3ac1666a14ebdc8747d934d8279): complete subsection reference.

- [cluster_scoped_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-7ef5b18a3ff71a18f7e46745faefce6e346edb7d752de7ae210c49c447c8b875): complete subsection reference.

- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c): complete subsection reference.

<a id="canonical-d87742567360eada0e153aed156b45ee79c3b9eb1986a50ab61bd8b8201f8e5a"></a>

<a id="canonical-e746466600532aaae6d1014141b45125a86b0eed744bd95c70c8476bb2d51165"></a>

## description property — Property reference / 3f99d6be32dc / 5

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

<a id="canonical-2cc14d0e366996c488d0a93b68b72ba81dba881e429a0daf8602bb385e08e097"></a>

<a id="canonical-c36882b67838cd60fc1d94844375983478e9d42e0d173d3c25d744f4c32e9ea6"></a>

## disable property — Property reference / 3f99d6be32dc / 6

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

- [global_access_enable](resources--k8s_cluster--reference--group-001.md#canonical-543d1f0fc63baae9c66cbe1c5240491656e5aba50f15c39b529158df54111045): complete subsection reference.

<a id="canonical-61fe963f8d38cd1f35d72bb076dfe1df94ab413ebdd7bcc13f42c890743ff400"></a>

<a id="canonical-7d0a05d5eb90269bf6b97f7cb0f40591d6b2ec76dacd383edbf9cbbec14e8e57"></a>

## id property — Property reference / 3f99d6be32dc / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [insecure_registry_list](resources--k8s_cluster--reference--group-001.md#canonical-9f18c90c28a7c9eb9e7bf968d83284e3770cec51707f5d8a7926526aec5a173e): complete subsection reference.

<a id="canonical-b67861a5e73812588aefb6ba55dfe57c0f222d67f759e7cdc811eeb9542a45b3"></a>

<a id="canonical-18b0e0eb726f279ee32e4dd8e1c2433623c4cc5acc821e05b3503cc12a6d141c"></a>

## labels property — Property reference / 3f99d6be32dc / 8

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

- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-bc5e2d2e775ad4b91024bf49cc0342741855cc0e716738813c095743ff937ac5): complete subsection reference.

<a id="canonical-a7e1382581981c4b9ee53534e574feaccbd8aea1b1b0e48f8e0a0f354a25e51d"></a>

<a id="canonical-3d549a450b26c24f14e0acc2814b8f46651034199832cb5a548a5e590057f4c1"></a>

## name property — Property reference / 3f99d6be32dc / 9

Type: `"string"`. Required.

Name of the K8S Cluster. Must be unique within the namespace.

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

<a id="canonical-2d02023e14ce919b744b2d1eda4b0e70c89ea3a712960230c6f105594ad22532"></a>

<a id="canonical-8a1cacd8578117ee80d3592e47a0baf6bc7f3a233e513201d223ca30463ce962"></a>

## namespace property — Property reference / 3f99d6be32dc / 10

Type: `"string"`. Optional, Computed.

Namespace for the K8S Cluster. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [no_cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-a474fd5876b5a65fecad6120997a86db59dba8e6bfb61016e2c699996986887a): complete subsection reference.

- [no_global_access](resources--k8s_cluster--reference--group-001.md#canonical-d2542bc357f403f6d9065ddfa9614cbd79137ec93646a18bd385e3ec4a0d710c): complete subsection reference.

- [no_insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-4e6fa247d2659ca871db09280d7684f2741b47b7fcb533ecb7c5b053fff80c30): complete subsection reference.

- [no_local_access](resources--k8s_cluster--reference--group-001.md#canonical-8077807aedc76aae39bbd3327b76639849bcc659063c9635feb47792e421671c): complete subsection reference.

- [timeouts](resources--k8s_cluster--reference--group-001.md#canonical-ea55ea298379df539e384e4b2c94ccfd7f52a9e1994dd70ac9b7b97d84f02ea9): complete subsection reference.

- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-25cba01023e85b19367de4784d90177d4db294a0e9b3f3fcdf55bff91fb25225): complete subsection reference.

- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-8fdbf543dbe360a9e9867e79272fa399eeb1b85d2ab759ccfc9686e42babce3a): complete subsection reference.

- [use_custom_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-b6339b23e04ef58c510d1d257369dfc5f67ad57e4dbdbe29630a24779be09dea): complete subsection reference.

- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-dd22941e4340fcc6dcf31dfc4f0efc217d4c08c20f27c8ecc7d8cdf105d51ffa): complete subsection reference.

- [use_default_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-cb1e02a99d9a2253aeb7ec284448dab8b4b7765e186c731e6f4d3e5b67a61d65): complete subsection reference.

- [use_default_cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-27beafb6844370155c1cdb6b23faa7ce792ed2d51ed3aa9c275de9791c6ce28f): complete subsection reference.

- [use_default_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-b2433afd2b95a60688052c623f67ebcf347e2fc615ad4124668a111bfa47e6d2): complete subsection reference.

- [use_default_psp](resources--k8s_cluster--reference--group-001.md#canonical-31c81631d89e30dfc747e80cf38eb02550b09013ec08649f35332173ba032a81): complete subsection reference.

- [vk8s_namespace_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-a41fd0b3544156ce047bd06318ca7cb854b541acd7f35cd6ee5e7def4b1ea629): complete subsection reference.

- [vk8s_namespace_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-85347192a1e17fcda18c48e600fa1a7171d635728a5dc28c6a9125eef26e7f2f): complete subsection reference.

<a id="canonical-fc435fa4c64e6d6366ad5dbcfcd26c8a8403cdc2466fb5190710ee8960631a10"></a>

## All schema paths — Property reference / 3f99d6be32dc / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_cluster--reference--group-001.md#canonical-0f9521d8c92b9c13ab38b1e811f73d500b650ae4b0760f3f3d5dadc7684645f5) |
| `cluster_scoped_access_deny` | [cluster_scoped_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-1f8c10536a578170b88e5aaa461bb09dad8120c973d164163603643d6649b81b) |
| `cluster_scoped_access_permit` | [cluster_scoped_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-abdcde6a03943d6a3599a53ff99cae5b8c3a139c6d1da39ee14eec31e1996ed3) |
| `cluster_wide_app_list` | [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-abca740f66fdabc96dbe2f0b31ec973eb0f40a211ddf2f953ea389a55a271464) |
| `cluster_wide_app_list.cluster_wide_apps` | [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-05b054c78bbeb9d492a33889f516da8bacad94519c87a8e4765df4c4ea6d135b) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd` | [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-eff74870a5e59eaaaff0f547ff722af1c6e61d37fa2cdae92ee3f04b3a3cff7b) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-80cdd3b03a60bd0ef5e42e43f5710c82e44ab4ccf874029e3b2655fda7b4d645) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](resources--k8s_cluster--reference--group-001.md#canonical-e632a3e9d2978d73c3d8aa01e59a5fedde96f14a95880414b3961f834f5aa9cc) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-d14bc79e36ea803058bd3ba1335663dfc00ccdeb21acbf18d36915741d85e8a1) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-1ff57e25274a491b7ef86341d33c4434a803e54a8350dc796b5af8a8a4a0c6f1) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-d9a275e4f48c40b80d03f9d1d2ece0f575eeee9552e5576bc794b64be29075b0) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider](resources--k8s_cluster--reference--group-001.md#canonical-e7addae00d27298dd17dd899026875ac63ff0555962ed8a26e744f627b538610) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location](resources--k8s_cluster--reference--group-001.md#canonical-c69ce50de82ad7e4ecd40e3b063072d8d4edebe2c6910bc17aff99bea2ae2ca1) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider](resources--k8s_cluster--reference--group-001.md#canonical-4d597f2f9a45027b19332bd4f1dc164cd2b06a832fc41078d42f3381c09ae309) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-5f92a71ad8f2b0b82cb22bf46c2cbbc5b164975add0c19128624d24f5a68d982) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref](resources--k8s_cluster--reference--group-001.md#canonical-4e08340eccba314f79e6e620335a81b8b63a8520c552746caa75cea812e05c58) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url](resources--k8s_cluster--reference--group-001.md#canonical-14fccf1a49e842a993400b8aee3be202617fa322226d53c813fd768862da7231) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port](resources--k8s_cluster--reference--group-001.md#canonical-e31b57b12630f4ee3e1a2b8b5f03175425bcf95b190c013447c36f727eb66082) |
| `cluster_wide_app_list.cluster_wide_apps.dashboard` | [cluster_wide_app_list.cluster_wide_apps.dashboard](resources--k8s_cluster--reference--group-001.md#canonical-ae3b685ae7a366434b15f8a2ac0d4821f3b43adfcba17a83c2f271fee8fd0e1d) |
| `cluster_wide_app_list.cluster_wide_apps.metrics_server` | [cluster_wide_app_list.cluster_wide_apps.metrics_server](resources--k8s_cluster--reference--group-001.md#canonical-40b1bcf737f4b15ec3da88fe58b62a4c135dfec13afea4b95b98369989409cc7) |
| `cluster_wide_app_list.cluster_wide_apps.prometheus` | [cluster_wide_app_list.cluster_wide_apps.prometheus](resources--k8s_cluster--reference--group-001.md#canonical-6cd3cdf3ec8e886118da85d877af0f5de9e1c001023a451a23e2595e2b9c2054) |
| `description` | [description](resources--k8s_cluster--reference--group-001.md#canonical-d87742567360eada0e153aed156b45ee79c3b9eb1986a50ab61bd8b8201f8e5a) |
| `disable` | [disable](resources--k8s_cluster--reference--group-001.md#canonical-2cc14d0e366996c488d0a93b68b72ba81dba881e429a0daf8602bb385e08e097) |
| `global_access_enable` | [global_access_enable](resources--k8s_cluster--reference--group-001.md#canonical-2752febf2bff790fa0f4c9dfe97670fb7a68437018105b3d5a302c30de34b824) |
| `id` | [id](resources--k8s_cluster--reference--group-001.md#canonical-61fe963f8d38cd1f35d72bb076dfe1df94ab413ebdd7bcc13f42c890743ff400) |
| `insecure_registry_list` | [insecure_registry_list](resources--k8s_cluster--reference--group-001.md#canonical-90283c95455c383b01010892e13b04501fcacb00317c84ca19f6ba9cbdbf7924) |
| `insecure_registry_list.insecure_registries` | [insecure_registry_list.insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-065a5814646df05f1733e81e36c6481850bd9bdbdce0aaf9aeba645c3fbefc59) |
| `labels` | [labels](resources--k8s_cluster--reference--group-001.md#canonical-b67861a5e73812588aefb6ba55dfe57c0f222d67f759e7cdc811eeb9542a45b3) |
| `local_access_config` | [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-9384ecf07868886d336e62b91ef6cd07abea7f2a9e74194c13794531354153eb) |
| `local_access_config.default_port` | [local_access_config.default_port](resources--k8s_cluster--reference--group-001.md#canonical-450d4bf8a45c35327f55fb6723632086452a762b7196e63a7e92be58bf9d16ed) |
| `local_access_config.local_domain` | [local_access_config.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-f0543d8d1d2690d4818d8c84e573dd0e5d90e49578048e3ae8b88e90d2d9612a) |
| `local_access_config.port` | [local_access_config.port](resources--k8s_cluster--reference--group-001.md#canonical-c488b6ef61d0ec580ccd728282f25b88a9138d536abbdff855209b0db41c716e) |
| `name` | [name](resources--k8s_cluster--reference--group-001.md#canonical-a7e1382581981c4b9ee53534e574feaccbd8aea1b1b0e48f8e0a0f354a25e51d) |
| `namespace` | [namespace](resources--k8s_cluster--reference--group-001.md#canonical-2d02023e14ce919b744b2d1eda4b0e70c89ea3a712960230c6f105594ad22532) |
| `no_cluster_wide_apps` | [no_cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-b1d7ae44056ecd50bd713bfb99c661b3e574b7f1056fc6ad684349e5f2bd0066) |
| `no_global_access` | [no_global_access](resources--k8s_cluster--reference--group-001.md#canonical-f7de48be95455070151370c3daa9e756d89e7f99d7f60be31653c3fcd4f82046) |
| `no_insecure_registries` | [no_insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-f9679af264173a1bce23024668298e5faa85e5e1340b6b6273e0d1c165417791) |
| `no_local_access` | [no_local_access](resources--k8s_cluster--reference--group-001.md#canonical-d469e8e32599d7f6fd4660d33891c446a5190b169c512d28b39d2329912142e8) |
| `timeouts` | [timeouts](resources--k8s_cluster--reference--group-001.md#canonical-e4f55491a6cade480b4abae11f38ef997dc7210455b85b945e737b489eb2e47c) |
| `timeouts.create` | [timeouts.create](resources--k8s_cluster--reference--group-001.md#canonical-18c0a60c0b84822358ca7e6f79a2048edbff25217634b774923d4e222498eb50) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_cluster--reference--group-001.md#canonical-7af5d9971e0c21e0c7155111567a76978a6288289a7aec1f4c96f0c73eb42927) |
| `timeouts.read` | [timeouts.read](resources--k8s_cluster--reference--group-001.md#canonical-7673ed8a6e77b36326b1bac0b8b3fe78acc65fd3009ec4d4ecb19d9b147b00c6) |
| `timeouts.update` | [timeouts.update](resources--k8s_cluster--reference--group-001.md#canonical-836a980d4c175a03e2a3787e76976f130bd75f181571abdcc46a608b8f8b4c2d) |
| `use_custom_cluster_role_bindings` | [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-49672d6fcf0f0c7e8956a9a1cf6697ead446d5e75f09d8338ccd6f2080f1ff1a) |
| `use_custom_cluster_role_bindings.cluster_role_bindings` | [use_custom_cluster_role_bindings.cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-754cceb9b17ad166da29308550c8f258bb967291485bb29556df6d95fb0ba7ed) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.name` | [use_custom_cluster_role_bindings.cluster_role_bindings.name](resources--k8s_cluster--reference--group-001.md#canonical-234b6f33ffb2baf717bcd491bfac1ceb6e29718e3c886d580eab9b4e4074e7ea) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.namespace` | [use_custom_cluster_role_bindings.cluster_role_bindings.namespace](resources--k8s_cluster--reference--group-001.md#canonical-d9453689f578b5c9dba7fbee4c2bdfa274d8c3c65a94d40f333040c5a6fb0de9) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.tenant` | [use_custom_cluster_role_bindings.cluster_role_bindings.tenant](resources--k8s_cluster--reference--group-001.md#canonical-318db49c2d0c8f0f401f597b106cdbbf326285b362e9de2c2dc006d386cffe3f) |
| `use_custom_cluster_role_list` | [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-3ace7eebe020f3acc082035789225e98f7b2d66d70b49c18f532ccaaf286658e) |
| `use_custom_cluster_role_list.cluster_roles` | [use_custom_cluster_role_list.cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-2af213205c2c909524d481f4c75ef195730f58961713ee8a8d7ae725fa2b5d6c) |
| `use_custom_cluster_role_list.cluster_roles.name` | [use_custom_cluster_role_list.cluster_roles.name](resources--k8s_cluster--reference--group-001.md#canonical-4dca25a56518b9de0abd3e6acbd852115fe9782861951910c5d06253efb936b1) |
| `use_custom_cluster_role_list.cluster_roles.namespace` | [use_custom_cluster_role_list.cluster_roles.namespace](resources--k8s_cluster--reference--group-001.md#canonical-24ae02e795fa64536d799eb55aca7650e8ade19d5bc0dc56fce0d29a4b10f1db) |
| `use_custom_cluster_role_list.cluster_roles.tenant` | [use_custom_cluster_role_list.cluster_roles.tenant](resources--k8s_cluster--reference--group-001.md#canonical-d2bff11faef4137f0a39a8db7fb176a67625434d1c9a0433d69fef0fa3937697) |
| `use_custom_pod_security_admission` | [use_custom_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-f0bc85b1b56e2fee8d678835b7dca2ac269250c5e5298aa7789fa426ed766893) |
| `use_custom_pod_security_admission.name` | [use_custom_pod_security_admission.name](resources--k8s_cluster--reference--group-001.md#canonical-773b34e1faca9bab3f85a11c036c3e4ca1be826db8ea49d8652ba77b89249194) |
| `use_custom_pod_security_admission.namespace` | [use_custom_pod_security_admission.namespace](resources--k8s_cluster--reference--group-001.md#canonical-8046d79f04e36eda8a56fde6e47abf3c8856aa4b49f8f0b0b372c86782f9fcf3) |
| `use_custom_pod_security_admission.tenant` | [use_custom_pod_security_admission.tenant](resources--k8s_cluster--reference--group-001.md#canonical-fd39695b381bbc7ad05964c44255deb3913112d86ea422f3b29575b537674af6) |
| `use_custom_psp_list` | [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-5d68481131a422668d6284da6845a36fcc93b68f4edeb2391da7061ad3f8a933) |
| `use_custom_psp_list.pod_security_policies` | [use_custom_psp_list.pod_security_policies](resources--k8s_cluster--reference--group-001.md#canonical-ba8dc5d6517cde7231abc99b970d2a6425db2d70628e783875a9099ebfb95f38) |
| `use_custom_psp_list.pod_security_policies.name` | [use_custom_psp_list.pod_security_policies.name](resources--k8s_cluster--reference--group-001.md#canonical-51384c44a7756ca5b6de5b69323b1676d8aa963014169bfc62a2a76bb69fbcf6) |
| `use_custom_psp_list.pod_security_policies.namespace` | [use_custom_psp_list.pod_security_policies.namespace](resources--k8s_cluster--reference--group-001.md#canonical-d697f326bec0b1c218794b2fc5da05dff856dbfec878cab347af4d649068cb37) |
| `use_custom_psp_list.pod_security_policies.tenant` | [use_custom_psp_list.pod_security_policies.tenant](resources--k8s_cluster--reference--group-001.md#canonical-b57adaa4e16242d6e52a73c216d95471540054308262b3f85b7a17ab0d876a58) |
| `use_default_cluster_role_bindings` | [use_default_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-bff446c1302b45e89806f90327befd375f30b72775d823e1487ee152d5d469d1) |
| `use_default_cluster_roles` | [use_default_cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-f9f246984f902f56305118fcecbfbdcf8e1819996337440d8263b0d4be1d9d79) |
| `use_default_pod_security_admission` | [use_default_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-108a63f2ba7c26252d02493b111a44d87844dde79ea8c42e932204bf65c0aa71) |
| `use_default_psp` | [use_default_psp](resources--k8s_cluster--reference--group-001.md#canonical-bc515b1b6bb5c7ff8534111ced546e0c0be86488430bd7365c076751bd2e5586) |
| `vk8s_namespace_access_deny` | [vk8s_namespace_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-6c59a27e79f29e14cc649273cedcc938ed456e1d3263fdbac6d19ef27a248994) |
| `vk8s_namespace_access_permit` | [vk8s_namespace_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-6482823529bab274852064f455879a22c409297a10dad14f7a90b2681d9bc117) |

<a id="canonical-7d07e073b9886bb4fd2e5b8b63bdcba41406b52748f6e80f16d06cfa5d7d2346"></a>

## Next pages — Property reference / 3f99d6be32dc / 12

- [cluster_scoped_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-ed8a65af165a6a1d818048346ad28b12d838f3ac1666a14ebdc8747d934d8279)
- [cluster_scoped_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-7ef5b18a3ff71a18f7e46745faefce6e346edb7d752de7ae210c49c447c8b875)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [global_access_enable](resources--k8s_cluster--reference--group-001.md#canonical-543d1f0fc63baae9c66cbe1c5240491656e5aba50f15c39b529158df54111045)
- [insecure_registry_list](resources--k8s_cluster--reference--group-001.md#canonical-9f18c90c28a7c9eb9e7bf968d83284e3770cec51707f5d8a7926526aec5a173e)
- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-bc5e2d2e775ad4b91024bf49cc0342741855cc0e716738813c095743ff937ac5)
- [no_cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-a474fd5876b5a65fecad6120997a86db59dba8e6bfb61016e2c699996986887a)
- [no_global_access](resources--k8s_cluster--reference--group-001.md#canonical-d2542bc357f403f6d9065ddfa9614cbd79137ec93646a18bd385e3ec4a0d710c)
- [no_insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-4e6fa247d2659ca871db09280d7684f2741b47b7fcb533ecb7c5b053fff80c30)
- [no_local_access](resources--k8s_cluster--reference--group-001.md#canonical-8077807aedc76aae39bbd3327b76639849bcc659063c9635feb47792e421671c)
- [timeouts](resources--k8s_cluster--reference--group-001.md#canonical-ea55ea298379df539e384e4b2c94ccfd7f52a9e1994dd70ac9b7b97d84f02ea9)
- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-25cba01023e85b19367de4784d90177d4db294a0e9b3f3fcdf55bff91fb25225)
- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-8fdbf543dbe360a9e9867e79272fa399eeb1b85d2ab759ccfc9686e42babce3a)
- [use_custom_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-b6339b23e04ef58c510d1d257369dfc5f67ad57e4dbdbe29630a24779be09dea)
- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-dd22941e4340fcc6dcf31dfc4f0efc217d4c08c20f27c8ecc7d8cdf105d51ffa)
- [use_default_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-cb1e02a99d9a2253aeb7ec284448dab8b4b7765e186c731e6f4d3e5b67a61d65)
- [use_default_cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-27beafb6844370155c1cdb6b23faa7ce792ed2d51ed3aa9c275de9791c6ce28f)
- [use_default_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-b2433afd2b95a60688052c623f67ebcf347e2fc615ad4124668a111bfa47e6d2)
- [use_default_psp](resources--k8s_cluster--reference--group-001.md#canonical-31c81631d89e30dfc747e80cf38eb02550b09013ec08649f35332173ba032a81)
- [vk8s_namespace_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-a41fd0b3544156ce047bd06318ca7cb854b541acd7f35cd6ee5e7def4b1ea629)
- [vk8s_namespace_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-85347192a1e17fcda18c48e600fa1a7171d635728a5dc28c6a9125eef26e7f2f)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-ed8a65af165a6a1d818048346ad28b12d838f3ac1666a14ebdc8747d934d8279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afa7e9a975c746045f1885d8484e8b55ffbd37b136ef0000552ef4d786630bda"></a>

## cluster_scoped_access_deny — cluster_scoped_access_deny / c08a7de8f632 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- cluster_scoped_access_deny

<a id="canonical-1f8c10536a578170b88e5aaa461bb09dad8120c973d164163603643d6649b81b"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: cluster\_scoped\_access\_deny, cluster\_scoped\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [cluster_scoped_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-1f8c10536a578170b88e5aaa461bb09dad8120c973d164163603643d6649b81b)
- [cluster_scoped_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-abdcde6a03943d6a3599a53ff99cae5b8c3a139c6d1da39ee14eec31e1996ed3)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cluster_scoped_access_deny = {}
```

<a id="canonical-6d2216d21ee917c3900be5abf5cdf7790128949e8528c7b779ea3d957afe1a08"></a>

## Direct properties — cluster_scoped_access_deny / c08a7de8f632 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-417fe3816d9173cdcb8631c13bfdae13db7b1b1e18e9ea41d6d0a51b211f83a7"></a>

## Next pages — cluster_scoped_access_deny / c08a7de8f632 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-7ef5b18a3ff71a18f7e46745faefce6e346edb7d752de7ae210c49c447c8b875"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de37371854bd352483b81bd7f5815fc402cb0631c5463fa17fc8070fed4f0726"></a>

## cluster_scoped_access_permit — cluster_scoped_access_permit / 9c7d08030fcc / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- cluster_scoped_access_permit

<a id="canonical-abdcde6a03943d6a3599a53ff99cae5b8c3a139c6d1da39ee14eec31e1996ed3"></a>

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
cluster_scoped_access_permit = {}
```

<a id="canonical-bd2c7e7442b396839026147760c98d35a7f375ef9216a80f0ce80241c0330448"></a>

## Direct properties — cluster_scoped_access_permit / 9c7d08030fcc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2fce6c38341ac4669d13875047028124219dbfc31ca6d575f6c9b9af73f93ce6"></a>

## Next pages — cluster_scoped_access_permit / 9c7d08030fcc / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a30f5e54a3cea3964bc0b4700ea4c76814373a746ef46bbaad27b0b82ac881ed"></a>

## cluster_wide_app_list — cluster_wide_app_list / 19cb1254769d / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- cluster_wide_app_list

<a id="canonical-abca740f66fdabc96dbe2f0b31ec973eb0f40a211ddf2f953ea389a55a271464"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cluster\_wide\_app\_list, no\_cluster\_wide\_apps; Default: no\_cluster\_wide\_apps\]
Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_wide_apps")}
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

OneOf alternatives in this subsection:

- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-abca740f66fdabc96dbe2f0b31ec973eb0f40a211ddf2f953ea389a55a271464)
- [no_cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-b1d7ae44056ecd50bd713bfb99c661b3e574b7f1056fc6ad684349e5f2bd0066)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cluster_wide_app_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c3cf10720596f7e9295b258407c3ae90ec8ce8b6f3895508810dc0c6152d3c2"></a>

## Direct properties — cluster_wide_app_list / 19cb1254769d / 3

- [cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88): complete subsection reference.

<a id="canonical-ede85a9c1b4e1db3eefbc56ac81270193687f6352362e5dfd95e7e078181bd2d"></a>

## Next pages — cluster_wide_app_list / 19cb1254769d / 4

- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d11b3c2cc79edbfed8b91b0f73286150886848a3bc7419a224267d8df0303ad2"></a>

## cluster_wide_app_list.cluster_wide_apps — cluster_wide_app_list.cluster_wide_apps / 82db6e776b37 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- cluster_wide_app_list.cluster_wide_apps

<a id="canonical-05b054c78bbeb9d492a33889f516da8bacad94519c87a8e4765df4c4ea6d135b"></a>

Type: `"object"`. list nested block, Optional.

Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("argo_cd",
    "dashboard"),
  validators.ConflictingListObjectAttributes("argo_cd",
    "metrics_server"),
  validators.ConflictingListObjectAttributes("argo_cd",
    "prometheus"),
  validators.ConflictingListObjectAttributes("dashboard",
    "metrics_server"),
  validators.ConflictingListObjectAttributes("dashboard",
    "prometheus"),
  validators.ConflictingListObjectAttributes("metrics_server",
    "prometheus")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
cluster_wide_apps {
  # Configure direct properties listed below.
}
```

<a id="canonical-274304107f0c95d06c8e5a938c583265ffd6289566dbce041399163ab6db98a3"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps / 82db6e776b37 / 3

- [argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac): complete subsection reference.

- [dashboard](resources--k8s_cluster--reference--group-001.md#canonical-d5707d2389d17e25564085fbdb47173e719694eb1b50521402918b8bc436a3c0): complete subsection reference.

- [metrics_server](resources--k8s_cluster--reference--group-001.md#canonical-0d15ffcf7f55332ccee14baada2866a8628c7f8d8f356e7875ca9ed9a8117039): complete subsection reference.

- [prometheus](resources--k8s_cluster--reference--group-001.md#canonical-92bc75ce5872a925b749aadd26d630f3993a273cc018cb4cb0c4f8a4cb772c47): complete subsection reference.

<a id="canonical-6289aacfbad0a3d960417ed84ef1463898f29570b393f54c2214ca3a41b7aba4"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps / 82db6e776b37 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac)
- [cluster_wide_app_list.cluster_wide_apps.dashboard](resources--k8s_cluster--reference--group-001.md#canonical-d5707d2389d17e25564085fbdb47173e719694eb1b50521402918b8bc436a3c0)
- [cluster_wide_app_list.cluster_wide_apps.metrics_server](resources--k8s_cluster--reference--group-001.md#canonical-0d15ffcf7f55332ccee14baada2866a8628c7f8d8f356e7875ca9ed9a8117039)
- [cluster_wide_app_list.cluster_wide_apps.prometheus](resources--k8s_cluster--reference--group-001.md#canonical-92bc75ce5872a925b749aadd26d630f3993a273cc018cb4cb0c4f8a4cb772c47)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97afa20d165800606a48c22ffa5040accd25dcffc8b6f109859ea5f44323a6af"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd — cluster_wide_app_list.cluster_wide_apps.argo_cd / 7e2c8a1090fd / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- cluster_wide_app_list.cluster_wide_apps.argo_cd

<a id="canonical-eff74870a5e59eaaaff0f547ff722af1c6e61d37fa2cdae92ee3f04b3a3cff7b"></a>

Type: `"object"`. single nested block, Optional.

Description Parameters for Argo Continuous Deployment(CD) application.

Upstream description:

Description Parameters for Argo Continuous Deployment(CD) application.

Receipt-pinned upstream constraints:

```json
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
argo_cd {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d06ba726fdde3857d2970044b18445a740e32578a9a32f3b34e267a74437c4d"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd / 7e2c8a1090fd / 3

- [local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731): complete subsection reference.

<a id="canonical-4c6abf4ad237bc54cdf2f1d03aa19e0a172fcf75ed49d16e39fc335b227988ab"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd / 7e2c8a1090fd / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdcfb4e7f32f862e6a3a20bdad8949af7606f965d4bb1fe389babd8372163306"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / e3d9f53ad825 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

<a id="canonical-80cdd3b03a60bd0ef5e42e43f5710c82e44ab4ccf874029e3b2655fda7b4d645"></a>

Type: `"object"`. single nested block, Optional.

Parameters required to enable local access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("local_domain"),
  validators.ConflictingObjectAttributes("default_port",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
local_domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-a353ddd6fa09fb41cb9ba62937fffdf18c2ec3cf4f0cbe6cd071e82832f8c43b"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / e3d9f53ad825 / 3

- [default_port](resources--k8s_cluster--reference--group-001.md#canonical-56ce2f948c684ce86f7f5d505af50c5d3d419f08f6bcf212b28a4dd012e6b825): complete subsection reference.

<a id="canonical-d14bc79e36ea803058bd3ba1335663dfc00ccdeb21acbf18d36915741d85e8a1"></a>

<a id="canonical-561cd5ef800814e386f004517e860414337ea50d6d0e400338f0f1e6a289df33"></a>

## local_domain property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / e3d9f53ad825 / 4

Type: `"string"`. Optional.

ArgoCD will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [password](resources--k8s_cluster--reference--group-001.md#canonical-e9b88cc1c3b73f6f5e917f836597b5f4e3d94a477e688d43223f7d06b7573183): complete subsection reference.

<a id="canonical-e31b57b12630f4ee3e1a2b8b5f03175425bcf95b190c013447c36f727eb66082"></a>

<a id="canonical-b000d94317b873da9639d06225f1c966b605f311617711ddad62bed81e2b4515"></a>

## port property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / e3d9f53ad825 / 5

Type: `"number"`. Optional.

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```

<a id="canonical-a42225e849c76b26cd5679cabcacb07702c241c4b96ac03f8172329f75b4de52"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / e3d9f53ad825 / 6

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](resources--k8s_cluster--reference--group-001.md#canonical-56ce2f948c684ce86f7f5d505af50c5d3d419f08f6bcf212b28a4dd012e6b825)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-e9b88cc1c3b73f6f5e917f836597b5f4e3d94a477e688d43223f7d06b7573183)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-56ce2f948c684ce86f7f5d505af50c5d3d419f08f6bcf212b28a4dd012e6b825"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7615a719ab67c134cb9da74bb2e5b9ff91098fb30cf060fe6b3acf945ee2134b"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port / 898af0cfddff / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port

<a id="canonical-e632a3e9d2978d73c3d8aa01e59a5fedde96f14a95880414b3961f834f5aa9cc"></a>

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
default_port = {}
```

<a id="canonical-529bc385c49aa98299cfdda1649964d9810d35f420c3394c8856c279b6533800"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port / 898af0cfddff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c67de95eb54fe0f37adf25a0cd7d9b6d9cdbc2d8e7390334aac7ffc2cd5cae0"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port / 898af0cfddff / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-e9b88cc1c3b73f6f5e917f836597b5f4e3d94a477e688d43223f7d06b7573183"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a707432ee5e45c27a18dc13d08dabec1885284d4e9cd8750a37048327b80027b"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password / b5122d6be4a3 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password

<a id="canonical-1ff57e25274a491b7ef86341d33c4434a803e54a8350dc796b5af8a8a4a0c6f1"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-52bd49c3a3b0aa41f1aa612c74b9c395cabec50d0746640560f76bf731945440"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password / b5122d6be4a3 / 3

- [blindfold_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-3dabba3e9dd1061bfd739fb071e265841e1a92260d14be6e06f601eaca568c1d): complete subsection reference.

- [clear_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-4b93fa7306577c62b40e6314d1ae6bffe87be3e82f7fbc10ffff5f7a45108d3e): complete subsection reference.

<a id="canonical-283e0d85038dd8792f74e43c7aca9e491b83c241fe42a178735488c1b45cd6e3"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password / b5122d6be4a3 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-3dabba3e9dd1061bfd739fb071e265841e1a92260d14be6e06f601eaca568c1d)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-4b93fa7306577c62b40e6314d1ae6bffe87be3e82f7fbc10ffff5f7a45108d3e)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-3dabba3e9dd1061bfd739fb071e265841e1a92260d14be6e06f601eaca568c1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4495ac28ec6799e2821374b7c15f84595f8af65d587e3a8f2ed09d2b67ad43c"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 0738e4d9d70e / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-e9b88cc1c3b73f6f5e917f836597b5f4e3d94a477e688d43223f7d06b7573183)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info

<a id="canonical-d9a275e4f48c40b80d03f9d1d2ece0f575eeee9552e5576bc794b64be29075b0"></a>

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

<a id="canonical-4852b4187dfe9268b64d79c1d5fe37d5f83c178e6b5954dde77a016cc03b07ec"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 0738e4d9d70e / 3

<a id="canonical-e7addae00d27298dd17dd899026875ac63ff0555962ed8a26e744f627b538610"></a>

<a id="canonical-6ebc719e7ca4c33bad6760a47b5cbd291babc1fc8c88bf0a343de4d6be828a28"></a>

## decryption_provider property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 0738e4d9d70e / 4

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

<a id="canonical-c69ce50de82ad7e4ecd40e3b063072d8d4edebe2c6910bc17aff99bea2ae2ca1"></a>

<a id="canonical-3924a23116e959f0c21ed809c8da0e3a9df3962733f418138b532609d179d8a9"></a>

## location property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 0738e4d9d70e / 5

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

<a id="canonical-4d597f2f9a45027b19332bd4f1dc164cd2b06a832fc41078d42f3381c09ae309"></a>

<a id="canonical-c6d196d3001c54624252d25207f044382360ffcfd30db52a15fc3a73a65c20eb"></a>

## store_provider property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 0738e4d9d70e / 6

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

<a id="canonical-025fabf1cf5416ad78d6f0f88d9a55660371e055770c2af638a1e640ea8a002a"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 0738e4d9d70e / 7

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-e9b88cc1c3b73f6f5e917f836597b5f4e3d94a477e688d43223f7d06b7573183)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-4b93fa7306577c62b40e6314d1ae6bffe87be3e82f7fbc10ffff5f7a45108d3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b20d160a4f70836109900e4ae5f240b857e4a26ef455c7f3c307da5fa2599be3"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / 49c1968df64a / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3e41391cf3270a6e421a5fb88592831530a936a3f0bc4ce32f57392ccf749dac)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-94411200a2074dfad1c5d206d3f095583169199e3bb1b1e9e6b16dbabce30731)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-e9b88cc1c3b73f6f5e917f836597b5f4e3d94a477e688d43223f7d06b7573183)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info

<a id="canonical-5f92a71ad8f2b0b82cb22bf46c2cbbc5b164975add0c19128624d24f5a68d982"></a>

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

<a id="canonical-d35797ec6d4ff1f9e4e09795f4e9b55613615105bd07dfbd8c45458d97d9e67f"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / 49c1968df64a / 3

<a id="canonical-4e08340eccba314f79e6e620335a81b8b63a8520c552746caa75cea812e05c58"></a>

<a id="canonical-e079a9d31e7204667a48120e18f7f41bdb4773d6a659e28c60f2902357665aaa"></a>

## provider_ref property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / 49c1968df64a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-14fccf1a49e842a993400b8aee3be202617fa322226d53c813fd768862da7231"></a>

<a id="canonical-8a769387abebd0edbd3636972f996f76ee2b2958acf6f70895ce36c7b7662181"></a>

## url property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / 49c1968df64a / 5

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

<a id="canonical-585541af29480a24868c0eb28c06841d9f9b2b3ea36c64bcfea419ff195ad04a"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / 49c1968df64a / 6

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-e9b88cc1c3b73f6f5e917f836597b5f4e3d94a477e688d43223f7d06b7573183)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-d5707d2389d17e25564085fbdb47173e719694eb1b50521402918b8bc436a3c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0afb064c867688b6738f8a3cb62361d28386bec1f60acb942656b233ec848096"></a>

## cluster_wide_app_list.cluster_wide_apps.dashboard — cluster_wide_app_list.cluster_wide_apps.dashboard / 45eb8218d665 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- cluster_wide_app_list.cluster_wide_apps.dashboard

<a id="canonical-ae3b685ae7a366434b15f8a2ac0d4821f3b43adfcba17a83c2f271fee8fd0e1d"></a>

Type: `["object", {}]`. Optional.

Description Parameters for K8s dashboard.

Upstream description:

Description Parameters for K8s dashboard.

Receipt-pinned upstream constraints:

```json
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
dashboard = {}
```

<a id="canonical-03316ec96e2f92d092f4e65b59c6ee0bbef239a4aa1abcf6428bc10fa1ef970c"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.dashboard / 45eb8218d665 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-37fd77f6f8d66f1794103132faa9c811b72f96ba3c68418b6b1bb7f5f73cdf76"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.dashboard / 45eb8218d665 / 4

- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-0d15ffcf7f55332ccee14baada2866a8628c7f8d8f356e7875ca9ed9a8117039"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a98c48bcb301d983eac49c77172266a667d37541885bcc4ea371a26261bd258"></a>

## cluster_wide_app_list.cluster_wide_apps.metrics_server — cluster_wide_app_list.cluster_wide_apps.metrics_server / 04c2e0d9b69f / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- cluster_wide_app_list.cluster_wide_apps.metrics_server

<a id="canonical-40b1bcf737f4b15ec3da88fe58b62a4c135dfec13afea4b95b98369989409cc7"></a>

Type: `["object", {}]`. Optional.

Description Parameters for Kubernetes Metrics Server application.

Upstream description:

Description Parameters for Kubernetes Metrics Server application.

Receipt-pinned upstream constraints:

```json
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
metrics_server = {}
```

<a id="canonical-127b6e859f449254a86decf71becd9ff0cc9cdeb4cbdae045f2c65281bc13d0b"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.metrics_server / 04c2e0d9b69f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9bff3a05070dd56b88ac949f0eaa7e0706dc11ba95cee9489f7e7a8fa8175b7"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.metrics_server / 04c2e0d9b69f / 4

- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-92bc75ce5872a925b749aadd26d630f3993a273cc018cb4cb0c4f8a4cb772c47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1de52d2ecf3502b1e923e2fd27a12be8188d9c0125dec9ff29906c9d7468933"></a>

## cluster_wide_app_list.cluster_wide_apps.prometheus — cluster_wide_app_list.cluster_wide_apps.prometheus / 84f6d4a95c24 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-7b86a3ea740ac67324e5a08c0e0f33ca0612c064feda800d9b3053b2f118b24c)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- cluster_wide_app_list.cluster_wide_apps.prometheus

<a id="canonical-6cd3cdf3ec8e886118da85d877af0f5de9e1c001023a451a23e2595e2b9c2054"></a>

Type: `["object", {}]`. Optional.

Description Parameters for Prometheus server access.

Upstream description:

Description Parameters for Prometheus server access.

Receipt-pinned upstream constraints:

```json
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
prometheus = {}
```

<a id="canonical-f85b37f813240fd941f1330e475c9ac2f9272272f6bda60b397ac30832afc664"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.prometheus / 84f6d4a95c24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21cf7742bd884f6592e0fd0ee74eb5bb441ed8111d267f87c9e41f755b3387d6"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.prometheus / 84f6d4a95c24 / 4

- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-20074d6659f74182cf15c8984b08537320f378d913653d2639b918123f37fd88)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-543d1f0fc63baae9c66cbe1c5240491656e5aba50f15c39b529158df54111045"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-425f862bd74c07bc210d50baef5d0ac7d719119113e9cb2421c1a281b773d912"></a>

## global_access_enable — global_access_enable / 369c72054ba6 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- global_access_enable

<a id="canonical-2752febf2bff790fa0f4c9dfe97670fb7a68437018105b3d5a302c30de34b824"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_access\_enable, no\_global\_access; Default: no\_global\_access\] Configuration
parameter for global access enable.

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

OneOf alternatives in this subsection:

- [global_access_enable](resources--k8s_cluster--reference--group-001.md#canonical-2752febf2bff790fa0f4c9dfe97670fb7a68437018105b3d5a302c30de34b824)
- [no_global_access](resources--k8s_cluster--reference--group-001.md#canonical-f7de48be95455070151370c3daa9e756d89e7f99d7f60be31653c3fcd4f82046)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_access_enable = {}
```

<a id="canonical-3fe9c815718203737a5547e40f95d7515c13a2aed944b0dfd7a66ff7bb179cad"></a>

## Direct properties — global_access_enable / 369c72054ba6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8553756fb9ef470c974bbfd597563670d5057f9d9768a3a251a03e774ab58fe"></a>

## Next pages — global_access_enable / 369c72054ba6 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-9f18c90c28a7c9eb9e7bf968d83284e3770cec51707f5d8a7926526aec5a173e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-605c3a36b9801253526d3218cf5e4decdcb105eea04ea7e9595506e19b9ad9c2"></a>

## insecure_registry_list — insecure_registry_list / f65ebd1993d6 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- insecure_registry_list

<a id="canonical-90283c95455c383b01010892e13b04501fcacb00317c84ca19f6ba9cbdbf7924"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: insecure\_registry\_list, no\_insecure\_registries; Default: no\_insecure\_registries\]
Docker Insecure Registry List. List of docker insecure registries.

Upstream description:

List of docker insecure registries.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("insecure_registries")}
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

OneOf alternatives in this subsection:

- [insecure_registry_list](resources--k8s_cluster--reference--group-001.md#canonical-90283c95455c383b01010892e13b04501fcacb00317c84ca19f6ba9cbdbf7924)
- [no_insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-f9679af264173a1bce23024668298e5faa85e5e1340b6b6273e0d1c165417791)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
insecure_registry_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-55a8109109744fa6e4ed1a7f19ea339f4d3dd612a0f53b6978db33f499fe2cb4"></a>

## Direct properties — insecure_registry_list / f65ebd1993d6 / 3

<a id="canonical-065a5814646df05f1733e81e36c6481850bd9bdbdce0aaf9aeba645c3fbefc59"></a>

<a id="canonical-12a9ec2db575c72985c9cfbf4ac3ccb231c04fa0e72f13479f915c8e12814549"></a>

## insecure_registries property — insecure_registry_list / f65ebd1993d6 / 4

Type: `["list", "string"]`. Optional.

List of docker insecure registries in format 'example.com:5000'.

Upstream description:

List of docker insecure registries in format "example.com:5000"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bd8b82dcb91cda4f8cdb3f27412d7c28f4334591f0a87177a258d292d2d34fb1"></a>

## Next pages — insecure_registry_list / f65ebd1993d6 / 5

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-bc5e2d2e775ad4b91024bf49cc0342741855cc0e716738813c095743ff937ac5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b572cf37d0ecba0344e40b6cd8773b3c2d5022e2cc55e51b6b74094e5fef2cb"></a>

## local_access_config — local_access_config / 74df580c20e9 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- local_access_config

<a id="canonical-9384ecf07868886d336e62b91ef6cd07abea7f2a9e74194c13794531354153eb"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: local\_access\_config, no\_local\_access; Default: no\_local\_access\] Parameters required
to enable local access.

Upstream description:

Parameters required to enable local access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("local_domain"),
  validators.ConflictingObjectAttributes("default_port",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

OneOf alternatives in this subsection:

- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-9384ecf07868886d336e62b91ef6cd07abea7f2a9e74194c13794531354153eb)
- [no_local_access](resources--k8s_cluster--reference--group-001.md#canonical-d469e8e32599d7f6fd4660d33891c446a5190b169c512d28b39d2329912142e8)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
local_access_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-831f35712db684c2cc36a39b0becd87cda6533f345e361f6fdf8d53a26edf794"></a>

## Direct properties — local_access_config / 74df580c20e9 / 3

- [default_port](resources--k8s_cluster--reference--group-001.md#canonical-a3fe14c91a0b8728dd5f9c8ff6fadd811571cc51f14634fd58b763f4b1ed7926): complete subsection reference.

<a id="canonical-f0543d8d1d2690d4818d8c84e573dd0e5d90e49578048e3ae8b88e90d2d9612a"></a>

<a id="canonical-a024dd5bf29f02ffe71fc0717189efa27095ce8e5a725b299a7ae28aa9344e40"></a>

## local_domain property — local_access_config / 74df580c20e9 / 4

Type: `"string"`. Optional.

Local K8s API server will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c488b6ef61d0ec580ccd728282f25b88a9138d536abbdff855209b0db41c716e"></a>

<a id="canonical-a14f689321b9f0a8d402ef6563316c455f8de184c85c91a65923b01fde3a3eb4"></a>

## port property — local_access_config / 74df580c20e9 / 5

Type: `"number"`. Optional.

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```

<a id="canonical-fc1b921d009c348f41f1f1357a750e8844ef3c9e19dd211b64e12f168579da0f"></a>

## Next pages — local_access_config / 74df580c20e9 / 6

- [local_access_config.default_port](resources--k8s_cluster--reference--group-001.md#canonical-a3fe14c91a0b8728dd5f9c8ff6fadd811571cc51f14634fd58b763f4b1ed7926)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-a3fe14c91a0b8728dd5f9c8ff6fadd811571cc51f14634fd58b763f4b1ed7926"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f151967a395e565a2854823ebfd97d5d24cecb620490a1e088b2d438c449f6b"></a>

## local_access_config.default_port — local_access_config.default_port / ccb552d05c68 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-bc5e2d2e775ad4b91024bf49cc0342741855cc0e716738813c095743ff937ac5)
- local_access_config.default_port

<a id="canonical-450d4bf8a45c35327f55fb6723632086452a762b7196e63a7e92be58bf9d16ed"></a>

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
default_port = {}
```

<a id="canonical-8dbd3d9f1d94295e6dd3741f8bad8e9da921fc0b8bbea5dbcacb520cf1faeb7c"></a>

## Direct properties — local_access_config.default_port / ccb552d05c68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-395bcf21799eec0311f599d7a14f87bf950635a029243e7c27333f28dd74fc5a"></a>

## Next pages — local_access_config.default_port / ccb552d05c68 / 4

- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-bc5e2d2e775ad4b91024bf49cc0342741855cc0e716738813c095743ff937ac5)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-a474fd5876b5a65fecad6120997a86db59dba8e6bfb61016e2c699996986887a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69ef36d1c0e600d1c7e8ee57e8e92b01206c4833ffb176c538ef591b682ea6d0"></a>

## no_cluster_wide_apps — no_cluster_wide_apps / c2d1bf7e08ea / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- no_cluster_wide_apps

<a id="canonical-b1d7ae44056ecd50bd713bfb99c661b3e574b7f1056fc6ad684349e5f2bd0066"></a>

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
no_cluster_wide_apps = {}
```

<a id="canonical-774ccd6d34f23b69fd89c6740b82b6de2255d8c8d5a73835e0e7eda5dc32fc30"></a>

## Direct properties — no_cluster_wide_apps / c2d1bf7e08ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65b3bd08914b56ca1799f5fa1dae7f9f6f436c3d0d4e4bec95d6db00c194a701"></a>

## Next pages — no_cluster_wide_apps / c2d1bf7e08ea / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-d2542bc357f403f6d9065ddfa9614cbd79137ec93646a18bd385e3ec4a0d710c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9cfb2226fc6739b63c3b740dd6d136e8eebf478306cf1079b00dca3956d91bd5"></a>

## no_global_access — no_global_access / e67a89b9ed82 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- no_global_access

<a id="canonical-f7de48be95455070151370c3daa9e756d89e7f99d7f60be31653c3fcd4f82046"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no global access. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_global_access = {}
```

<a id="canonical-0281ce20ec849230d005f2a5dc63a5e5ba9e651806d1284215b4d53412f05ced"></a>

## Direct properties — no_global_access / e67a89b9ed82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-221236833b4258df614de1c57c5a526048792e9d0873936e0373be05a857a847"></a>

## Next pages — no_global_access / e67a89b9ed82 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-4e6fa247d2659ca871db09280d7684f2741b47b7fcb533ecb7c5b053fff80c30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9ff3319b413084407004d04df1f21eeeb3860703e4a94928fad71360c21e970"></a>

## no_insecure_registries — no_insecure_registries / ec13f985dce0 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- no_insecure_registries

<a id="canonical-f9679af264173a1bce23024668298e5faa85e5e1340b6b6273e0d1c165417791"></a>

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
no_insecure_registries = {}
```

<a id="canonical-cdef9d5c329cf29ca7e7eb7ac6951c14e8e84fc387309d7d7cee67f0c8eeccef"></a>

## Direct properties — no_insecure_registries / ec13f985dce0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4f920e10d5da44d0bccd42e34456c1365d889bb5ba3b1ee1273231e993a60ec"></a>

## Next pages — no_insecure_registries / ec13f985dce0 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-8077807aedc76aae39bbd3327b76639849bcc659063c9635feb47792e421671c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ceb0852ced02b549196de17b819e0d6d697f6810a63b384ed4fac743d21a8637"></a>

## no_local_access — no_local_access / 7ae92793a79d / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- no_local_access

<a id="canonical-d469e8e32599d7f6fd4660d33891c446a5190b169c512d28b39d2329912142e8"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no local access. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_local_access = {}
```

<a id="canonical-f38b60a644faac81109c4871d5da6c5f5e60d333cc2f6abb3dd3f98f1b1f6af2"></a>

## Direct properties — no_local_access / 7ae92793a79d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ad780729e805f738b9743e04fa4058081f83dd7d1edcc40d984d766032594f7"></a>

## Next pages — no_local_access / 7ae92793a79d / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-ea55ea298379df539e384e4b2c94ccfd7f52a9e1994dd70ac9b7b97d84f02ea9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2737f45a11cc06c2896588e9ddd5b243001cad8e3f894ab22cacf9a33cda9df"></a>

## timeouts — timeouts / d9a53264df65 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- timeouts

<a id="canonical-e4f55491a6cade480b4abae11f38ef997dc7210455b85b945e737b489eb2e47c"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-964c98c0c6dd7fc27148b24a1f8c3d4137bd84c9f6f02cdacfde7f75d552a498"></a>

## Direct properties — timeouts / d9a53264df65 / 3

<a id="canonical-18c0a60c0b84822358ca7e6f79a2048edbff25217634b774923d4e222498eb50"></a>

<a id="canonical-8e4a6720f6766142c290d3d291b724a65dbf13d1601b2c0c55935a399e4edb49"></a>

## create property — timeouts / d9a53264df65 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7af5d9971e0c21e0c7155111567a76978a6288289a7aec1f4c96f0c73eb42927"></a>

<a id="canonical-bf97cf343a54c9b356d9abad013f5acb496fbff64f2269dab3a4a6c6b1191e63"></a>

## delete property — timeouts / d9a53264df65 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-7673ed8a6e77b36326b1bac0b8b3fe78acc65fd3009ec4d4ecb19d9b147b00c6"></a>

<a id="canonical-a0a939b762936a54b536eced41483e174261914cf5f00bc41ae745df92e80e9c"></a>

## read property — timeouts / d9a53264df65 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-836a980d4c175a03e2a3787e76976f130bd75f181571abdcc46a608b8f8b4c2d"></a>

<a id="canonical-9b742758752cc67893b1a88ea46f7871cf4d3ba077b483fdd0e684c22c19a9f1"></a>

## update property — timeouts / d9a53264df65 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-19e5ac61c9022ee7055e2683e3a516d7541333f92c8f5c3f568e9c3bf6d2a6d5"></a>

## Next pages — timeouts / d9a53264df65 / 8

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-25cba01023e85b19367de4784d90177d4db294a0e9b3f3fcdf55bff91fb25225"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d621e07c3dc1db8ca5cf4229d924ddb3677171808aecd1e9f36956fb685b23da"></a>

## use_custom_cluster_role_bindings — use_custom_cluster_role_bindings / 0878dfaa1e44 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_custom_cluster_role_bindings

<a id="canonical-49672d6fcf0f0c7e8956a9a1cf6697ead446d5e75f09d8338ccd6f2080f1ff1a"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_bindings, use\_default\_cluster\_role\_bindings; Default:
use\_default\_cluster\_role\_bindings\] List of active cluster role binding list for a K8s cluster.

Upstream description:

List of active cluster role binding list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_role_bindings")}
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

OneOf alternatives in this subsection:

- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-49672d6fcf0f0c7e8956a9a1cf6697ead446d5e75f09d8338ccd6f2080f1ff1a)
- [use_default_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-bff446c1302b45e89806f90327befd375f30b72775d823e1487ee152d5d469d1)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_bindings {
  # Configure direct properties listed below.
}
```

<a id="canonical-1dd41bc33eb10b8dc75c28dbb465e1fdf167f2379f9cf8b5c519ebe4f53d7736"></a>

## Direct properties — use_custom_cluster_role_bindings / 0878dfaa1e44 / 3

- [cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-6af0142df7f14f379371f23ab7acf05961632fab13ce4c219ca113a733ace3e7): complete subsection reference.

<a id="canonical-35aac661688babd694b256488e4cda2a418be95dcca058f41600d58733a44fa9"></a>

## Next pages — use_custom_cluster_role_bindings / 0878dfaa1e44 / 4

- [use_custom_cluster_role_bindings.cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-6af0142df7f14f379371f23ab7acf05961632fab13ce4c219ca113a733ace3e7)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-6af0142df7f14f379371f23ab7acf05961632fab13ce4c219ca113a733ace3e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40b0a01460a23117f6134fa13d3e010f08bda28988f9c4749e55531ec247003a"></a>

## use_custom_cluster_role_bindings.cluster_role_bindings — use_custom_cluster_role_bindings.cluster_role_bindings / 8a234e20a10b / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-25cba01023e85b19367de4784d90177d4db294a0e9b3f3fcdf55bff91fb25225)
- use_custom_cluster_role_bindings.cluster_role_bindings

<a id="canonical-754cceb9b17ad166da29308550c8f258bb967291485bb29556df6d95fb0ba7ed"></a>

Type: `"object"`. list nested block, Optional.

List of active cluster role binding list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cluster_role_bindings {
  # Configure direct properties listed below.
}
```

<a id="canonical-45702676b055a8604a5791efacb36785ef5f946f4ea72a64f7f8e28c06e23346"></a>

## Direct properties — use_custom_cluster_role_bindings.cluster_role_bindings / 8a234e20a10b / 3

<a id="canonical-234b6f33ffb2baf717bcd491bfac1ceb6e29718e3c886d580eab9b4e4074e7ea"></a>

<a id="canonical-828b412c4adfd3597005789aea6e368b52a64cb66063cb596f2273a4c0bc8258"></a>

## name property — use_custom_cluster_role_bindings.cluster_role_bindings / 8a234e20a10b / 4

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

<a id="canonical-d9453689f578b5c9dba7fbee4c2bdfa274d8c3c65a94d40f333040c5a6fb0de9"></a>

<a id="canonical-221bc19a77e1909ba04656b9b2b4f7258529c37ffb3951c91cf9de45ac9ecc60"></a>

## namespace property — use_custom_cluster_role_bindings.cluster_role_bindings / 8a234e20a10b / 5

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

<a id="canonical-318db49c2d0c8f0f401f597b106cdbbf326285b362e9de2c2dc006d386cffe3f"></a>

<a id="canonical-e2d8bc0847faed9138af9269d0f07121fb2c37aed5eacbfd0d3bada9e9c75065"></a>

## tenant property — use_custom_cluster_role_bindings.cluster_role_bindings / 8a234e20a10b / 6

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

<a id="canonical-409979bb31ab597325113ba5a6896f27b2f1c73d0b2fd5f8ebb9610333a729da"></a>

## Next pages — use_custom_cluster_role_bindings.cluster_role_bindings / 8a234e20a10b / 7

- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-25cba01023e85b19367de4784d90177d4db294a0e9b3f3fcdf55bff91fb25225)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-8fdbf543dbe360a9e9867e79272fa399eeb1b85d2ab759ccfc9686e42babce3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e767c64809be9b3cf94d40fd1cc78242e9e7e4b321a4fa12b8f2f5c0fe32ff2"></a>

## use_custom_cluster_role_list — use_custom_cluster_role_list / 160bed7394ae / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_custom_cluster_role_list

<a id="canonical-3ace7eebe020f3acc082035789225e98f7b2d66d70b49c18f532ccaaf286658e"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_list, use\_default\_cluster\_roles; Default:
use\_default\_cluster\_roles\] List of active cluster role list for a K8s cluster.

Upstream description:

List of active cluster role list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_roles")}
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

OneOf alternatives in this subsection:

- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-3ace7eebe020f3acc082035789225e98f7b2d66d70b49c18f532ccaaf286658e)
- [use_default_cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-f9f246984f902f56305118fcecbfbdcf8e1819996337440d8263b0d4be1d9d79)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-c555aa28ce0f58fff3fd8db4d6430a54730e7e252f63813891772eff3bb92be1"></a>

## Direct properties — use_custom_cluster_role_list / 160bed7394ae / 3

- [cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-3c1351c2e5ba59825cebf8bfd11b2e6a0e9af5b3ca348b1841793bba1ac2f0a9): complete subsection reference.

<a id="canonical-9e5e00fc1d7302e5f7d4eb2ea176021c16a8bf6030514b6fc5d33cfd836f299a"></a>

## Next pages — use_custom_cluster_role_list / 160bed7394ae / 4

- [use_custom_cluster_role_list.cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-3c1351c2e5ba59825cebf8bfd11b2e6a0e9af5b3ca348b1841793bba1ac2f0a9)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-3c1351c2e5ba59825cebf8bfd11b2e6a0e9af5b3ca348b1841793bba1ac2f0a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd94fc69db43b1727e3d62038fe25881445e7b97397fe7dae91acb246b02719e"></a>

## use_custom_cluster_role_list.cluster_roles — use_custom_cluster_role_list.cluster_roles / 55cfd345effb / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-8fdbf543dbe360a9e9867e79272fa399eeb1b85d2ab759ccfc9686e42babce3a)
- use_custom_cluster_role_list.cluster_roles

<a id="canonical-2af213205c2c909524d481f4c75ef195730f58961713ee8a8d7ae725fa2b5d6c"></a>

Type: `"object"`. list nested block, Optional.

List of active cluster role list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cluster_roles {
  # Configure direct properties listed below.
}
```

<a id="canonical-95e948a0835627267be8a0a5b9934b1e13ea4437e4b6c400332be0ad4bb53d05"></a>

## Direct properties — use_custom_cluster_role_list.cluster_roles / 55cfd345effb / 3

<a id="canonical-4dca25a56518b9de0abd3e6acbd852115fe9782861951910c5d06253efb936b1"></a>

<a id="canonical-ad25531840163240846767b6308b3a7f5c08ecf972a8ebfb5a82abcefbb21476"></a>

## name property — use_custom_cluster_role_list.cluster_roles / 55cfd345effb / 4

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

<a id="canonical-24ae02e795fa64536d799eb55aca7650e8ade19d5bc0dc56fce0d29a4b10f1db"></a>

<a id="canonical-81d8f3116fafdd4ebe1b890a29bab92825da3acfbe13181b1845191a50f5a678"></a>

## namespace property — use_custom_cluster_role_list.cluster_roles / 55cfd345effb / 5

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

<a id="canonical-d2bff11faef4137f0a39a8db7fb176a67625434d1c9a0433d69fef0fa3937697"></a>

<a id="canonical-1634acea2eda38acb150bc82d80855398811134ad120e5520602567af243a31c"></a>

## tenant property — use_custom_cluster_role_list.cluster_roles / 55cfd345effb / 6

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

<a id="canonical-a2c4a35de6afeec0297175d5714eb442463af534eccb9310c717806101197574"></a>

## Next pages — use_custom_cluster_role_list.cluster_roles / 55cfd345effb / 7

- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-8fdbf543dbe360a9e9867e79272fa399eeb1b85d2ab759ccfc9686e42babce3a)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-b6339b23e04ef58c510d1d257369dfc5f67ad57e4dbdbe29630a24779be09dea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ffc5c1786c6eaf57e401952896427f9197bb3239ddfc9df6a06f74a59590456"></a>

## use_custom_pod_security_admission — use_custom_pod_security_admission / 44af20259d62 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_custom_pod_security_admission

<a id="canonical-f0bc85b1b56e2fee8d678835b7dca2ac269250c5e5298aa7789fa426ed766893"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_pod\_security\_admission, use\_default\_pod\_security\_admission; Default:
use\_default\_pod\_security\_admission\] Type establishes a direct reference from one object(the
referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [use_custom_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-f0bc85b1b56e2fee8d678835b7dca2ac269250c5e5298aa7789fa426ed766893)
- [use_default_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-108a63f2ba7c26252d02493b111a44d87844dde79ea8c42e932204bf65c0aa71)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_pod_security_admission {
  # Configure direct properties listed below.
}
```

<a id="canonical-2485877a5f233ab9c8878ec6d507dcc0d8360abe10917eb205f98775f34dd135"></a>

## Direct properties — use_custom_pod_security_admission / 44af20259d62 / 3

<a id="canonical-773b34e1faca9bab3f85a11c036c3e4ca1be826db8ea49d8652ba77b89249194"></a>

<a id="canonical-915d86295199263f451fa0ea7e992f668b91cca8cb47e08f6ff3d358c2f10a63"></a>

## name property — use_custom_pod_security_admission / 44af20259d62 / 4

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

<a id="canonical-8046d79f04e36eda8a56fde6e47abf3c8856aa4b49f8f0b0b372c86782f9fcf3"></a>

<a id="canonical-aa4a13b0bd8445e10f9f541fa3ad627c1607a2d9669d2c4d8938707ce146c482"></a>

## namespace property — use_custom_pod_security_admission / 44af20259d62 / 5

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

<a id="canonical-fd39695b381bbc7ad05964c44255deb3913112d86ea422f3b29575b537674af6"></a>

<a id="canonical-285e7d5bf0dd1b428dc429f4a87ef93b7c0d12a476f8d2a07d0cb36cb50e2147"></a>

## tenant property — use_custom_pod_security_admission / 44af20259d62 / 6

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

<a id="canonical-4913804109f32b8cfce182dc124638f37b19061dfd427f73f6b7508f5d4f056e"></a>

## Next pages — use_custom_pod_security_admission / 44af20259d62 / 7

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-dd22941e4340fcc6dcf31dfc4f0efc217d4c08c20f27c8ecc7d8cdf105d51ffa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78d7451dfcc10af959c1b741396031e784a9c082f6f1bfcdf8e2b6b2fa65c4e2"></a>

## use_custom_psp_list — use_custom_psp_list / 4e43bd7d9524 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_custom_psp_list

<a id="canonical-5d68481131a422668d6284da6845a36fcc93b68f4edeb2391da7061ad3f8a933"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

Upstream description:

List of active Pod security policies for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("pod_security_policies")}
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

OneOf alternatives in this subsection:

- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-5d68481131a422668d6284da6845a36fcc93b68f4edeb2391da7061ad3f8a933)
- [use_default_psp](resources--k8s_cluster--reference--group-001.md#canonical-bc515b1b6bb5c7ff8534111ced546e0c0be86488430bd7365c076751bd2e5586)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_psp_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb8d1e2af4a7b063feb93f62022b945b7a71aadb8108efc23532ba7295dd1761"></a>

## Direct properties — use_custom_psp_list / 4e43bd7d9524 / 3

- [pod_security_policies](resources--k8s_cluster--reference--group-001.md#canonical-26566989c5eec49a62d01a22b1cf91b7c621b18775098b9b971bd10f39203fc3): complete subsection reference.

<a id="canonical-2dbb6daafc85ebb5614db3e11c978f7424d1971e65ee88eade837406df844fc3"></a>

## Next pages — use_custom_psp_list / 4e43bd7d9524 / 4

- [use_custom_psp_list.pod_security_policies](resources--k8s_cluster--reference--group-001.md#canonical-26566989c5eec49a62d01a22b1cf91b7c621b18775098b9b971bd10f39203fc3)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-26566989c5eec49a62d01a22b1cf91b7c621b18775098b9b971bd10f39203fc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e0c60fd743bdc5e69214faa13814f36fd38dca9769e693ed927200782e70101"></a>

## use_custom_psp_list.pod_security_policies — use_custom_psp_list.pod_security_policies / 8d93d432a8f7 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-dd22941e4340fcc6dcf31dfc4f0efc217d4c08c20f27c8ecc7d8cdf105d51ffa)
- use_custom_psp_list.pod_security_policies

<a id="canonical-ba8dc5d6517cde7231abc99b970d2a6425db2d70628e783875a9099ebfb95f38"></a>

Type: `"object"`. list nested block, Optional.

List of active Pod security policies for a K8s cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pod_security_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-9b3e102ea05fd438d116956312f58860f8f30e1b8abcd33ac835b36337d19792"></a>

## Direct properties — use_custom_psp_list.pod_security_policies / 8d93d432a8f7 / 3

<a id="canonical-51384c44a7756ca5b6de5b69323b1676d8aa963014169bfc62a2a76bb69fbcf6"></a>

<a id="canonical-afc236ca1c4434ff5a29b16227b0cea9f3ea1283edc8e52b14d9b12ba719ae39"></a>

## name property — use_custom_psp_list.pod_security_policies / 8d93d432a8f7 / 4

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

<a id="canonical-d697f326bec0b1c218794b2fc5da05dff856dbfec878cab347af4d649068cb37"></a>

<a id="canonical-98a11cb266c54cb5d5250786058deb71964dbbe5cd58f58bde054851f37a86e9"></a>

## namespace property — use_custom_psp_list.pod_security_policies / 8d93d432a8f7 / 5

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

<a id="canonical-b57adaa4e16242d6e52a73c216d95471540054308262b3f85b7a17ab0d876a58"></a>

<a id="canonical-5bcd627250f87d86027653a571a1baeeae51d78147320b50918f1036dd1a9f92"></a>

## tenant property — use_custom_psp_list.pod_security_policies / 8d93d432a8f7 / 6

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

<a id="canonical-5a48e87ceca129d3abb4048f22f879e01b5eecea940f545d53addfb5afa1bdc0"></a>

## Next pages — use_custom_psp_list.pod_security_policies / 8d93d432a8f7 / 7

- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-dd22941e4340fcc6dcf31dfc4f0efc217d4c08c20f27c8ecc7d8cdf105d51ffa)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-cb1e02a99d9a2253aeb7ec284448dab8b4b7765e186c731e6f4d3e5b67a61d65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23edcac97ef5c97dced3ad01bbebb85e46a49403360ea12482d36e02acc3ed78"></a>

## use_default_cluster_role_bindings — use_default_cluster_role_bindings / 212a7c017031 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_default_cluster_role_bindings

<a id="canonical-bff446c1302b45e89806f90327befd375f30b72775d823e1487ee152d5d469d1"></a>

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
use_default_cluster_role_bindings = {}
```

<a id="canonical-adf1775b24e6701b8c4d5d49c7cb8669cad4632ea8bcb19b408fa1977a2b7ab0"></a>

## Direct properties — use_default_cluster_role_bindings / 212a7c017031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c878acab5d45ad0c156ccae75303dc47467231c5d9159edf5e674af3d7da54e0"></a>

## Next pages — use_default_cluster_role_bindings / 212a7c017031 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-27beafb6844370155c1cdb6b23faa7ce792ed2d51ed3aa9c275de9791c6ce28f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5cf9d6eb9ebdf4417ff3ce110bcc903b3b88b7c88549bc082a273cdf9bb9d2fd"></a>

## use_default_cluster_roles — use_default_cluster_roles / 53dd26dc4b11 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_default_cluster_roles

<a id="canonical-f9f246984f902f56305118fcecbfbdcf8e1819996337440d8263b0d4be1d9d79"></a>

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
use_default_cluster_roles = {}
```

<a id="canonical-79c1cfcf846c5d5dec5e25aa4568ff6fa87ebfa4ffd3ca25e2cbcbc5c392ce5c"></a>

## Direct properties — use_default_cluster_roles / 53dd26dc4b11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90b594ee01a30918e51d043e02de6c47ec11b0da95b909fd21d298f97d0e000c"></a>

## Next pages — use_default_cluster_roles / 53dd26dc4b11 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-b2433afd2b95a60688052c623f67ebcf347e2fc615ad4124668a111bfa47e6d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7b60cf32f2e13a4804f03f71bb492b6d34b72354d62702d2fc990438e7c8a34"></a>

## use_default_pod_security_admission — use_default_pod_security_admission / b599375400e1 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_default_pod_security_admission

<a id="canonical-108a63f2ba7c26252d02493b111a44d87844dde79ea8c42e932204bf65c0aa71"></a>

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
use_default_pod_security_admission = {}
```

<a id="canonical-f8f25bbbb4f5099036f50b8964a9ec76918c1fb2f1365cb4a3b3c34b6bb0dd89"></a>

## Direct properties — use_default_pod_security_admission / b599375400e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b524d13acf71c0662c2f2c8bdad77ed14c52d83b1012b9624e1b466a0474bad6"></a>

## Next pages — use_default_pod_security_admission / b599375400e1 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-31c81631d89e30dfc747e80cf38eb02550b09013ec08649f35332173ba032a81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76694e172d467e8e20841b1c821218eb59a61570b51044928cbc2f8a39887471"></a>

## use_default_psp — use_default_psp / b6d4cbed8a5e / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- use_default_psp

<a id="canonical-bc515b1b6bb5c7ff8534111ced546e0c0be86488430bd7365c076751bd2e5586"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for use default psp. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
use_default_psp = {}
```

<a id="canonical-a81a5145f931446b2f397fc1aeb495eaf613210a367a877887e5149b7af67d77"></a>

## Direct properties — use_default_psp / b6d4cbed8a5e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26d13904979cd065548d738e13e12a1acb92a57ef76075bf8454c6cff986801f"></a>

## Next pages — use_default_psp / b6d4cbed8a5e / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-a41fd0b3544156ce047bd06318ca7cb854b541acd7f35cd6ee5e7def4b1ea629"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9756117dee097f6c6c7c7105d6680c68e97c227688e770146d914b2bead78258"></a>

## vk8s_namespace_access_deny — vk8s_namespace_access_deny / 65d0052d5915 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- vk8s_namespace_access_deny

<a id="canonical-6c59a27e79f29e14cc649273cedcc938ed456e1d3263fdbac6d19ef27a248994"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: vk8s\_namespace\_access\_deny, vk8s\_namespace\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [vk8s_namespace_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-6c59a27e79f29e14cc649273cedcc938ed456e1d3263fdbac6d19ef27a248994)
- [vk8s_namespace_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-6482823529bab274852064f455879a22c409297a10dad14f7a90b2681d9bc117)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
vk8s_namespace_access_deny = {}
```

<a id="canonical-79505c9e9389a8d8677d91063a97176c2f1da642902f4bdf720443ff8bd7ed42"></a>

## Direct properties — vk8s_namespace_access_deny / 65d0052d5915 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0bc8efebb44ddfde10ab8b1f325a341894740608e04dd858908687569fcdcd22"></a>

## Next pages — vk8s_namespace_access_deny / 65d0052d5915 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-85347192a1e17fcda18c48e600fa1a7171d635728a5dc28c6a9125eef26e7f2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b076bb0398e2a92db963b02101191c989e9726183a7b64a9b8a9678b03ac8e2e"></a>

## vk8s_namespace_access_permit — vk8s_namespace_access_permit / 968aca103839 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- vk8s_namespace_access_permit

<a id="canonical-6482823529bab274852064f455879a22c409297a10dad14f7a90b2681d9bc117"></a>

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
vk8s_namespace_access_permit = {}
```

<a id="canonical-891fd0b639397edd01db3bce925e24f1fb2b238c9d1157915aba0041845abfcd"></a>

## Direct properties — vk8s_namespace_access_permit / 968aca103839 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce2a5501f8afa597c63d9a3ac67855ae8572b9764fcd355906f769108e0bdb27"></a>

## Next pages — vk8s_namespace_access_permit / 968aca103839 / 4

- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
