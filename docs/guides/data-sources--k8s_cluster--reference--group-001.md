---
page_title: "xcsh_k8s_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster reference."
---

# xcsh_k8s_cluster reference

<a id="canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6216f06b2d993afc443919a54e3575e325c91e1bee853847ee73ae38fc66fcae"></a>

## Property reference — Property reference / 01dbbc76da25 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- Property reference

<a id="canonical-4832ce7143e662a9ba4c2d467e33fac0de0d9f7cfe7b7bf950f837fe3ef3e5ba"></a>

## Direct properties — Property reference / 01dbbc76da25 / 3

<a id="canonical-8cce7a9b3629da85a9b381cf2d38268689e4b61361f3080563f4bc3d2c5a7e8c"></a>

<a id="canonical-7faabac8a12452f1093e063dcffc158303ff54fb5878a2bc71da8e0657038134"></a>

## annotations property — Property reference / 01dbbc76da25 / 4

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

- [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-180bdcc884989526420137adf6d76588eae18458ff269385040e756e71ce8c60): complete subsection reference.

- [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-419ca709b10b3d101be75663dfb5543c8967c02ae0f4fd75e5506540ef11efa2): complete subsection reference.

- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e): complete subsection reference.

<a id="canonical-1a3e750b434bc109ed58f759d4f656a8765cb7dc35e588fcb420aa0416d6f58b"></a>

<a id="canonical-8d5849378dc4dfccbbf6af018560011d5f7f46675b0f6d0d5f02aacce5c4251e"></a>

## description property — Property reference / 01dbbc76da25 / 5

Type: `"string"`. Computed.

Description of the K8SCluster.

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

- [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-f8ab74e4db9b790e1f33596a4418bad9fb7a98a4193dd7694c8a7dada64d60a8): complete subsection reference.

<a id="canonical-2285fce920407dd5de75f2c55cbe757f14a33d7f4bccf8c3c8e7e4fc60913535"></a>

<a id="canonical-efb80ad6c4c1aadc1e6ebcd900322c15cd2ca847b0059b00f2ed132fb5a713b4"></a>

## id property — Property reference / 01dbbc76da25 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-8033d79e8c97d5fce86173ed9a92916d2782f4deb719c478fa22f184112ef7c9): complete subsection reference.

<a id="canonical-08795068e152011049494c604e9cdd5a6aa7c2884c24fa188438ad6c964b6403"></a>

<a id="canonical-4657e0373923da8e48c8a50fb37a62cecd31a0d3ffda690a50d949c59ccd3faa"></a>

## labels property — Property reference / 01dbbc76da25 / 7

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

- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-6e733ff29bc886c146c986108d51341f7f0a592f241e9bff8794f81c9a01367d): complete subsection reference.

<a id="canonical-e0e8f087754f035e1d9102f48f8cd76b47df0c30bd33dde80c62eb9c92f937c7"></a>

<a id="canonical-05bfcbf7404486b7ff91943e1e8400e1ee2849ae98a906a9f2536cb9f9f18c07"></a>

## name property — Property reference / 01dbbc76da25 / 8

Type: `"string"`. Required.

Name of the K8SCluster.

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

<a id="canonical-c13c7bbbc7c4be0c26ec8b9993dea47b1366dabea5f5b40ea682872dcfc5d12d"></a>

<a id="canonical-aeeb0e113848a9a06588882723301bbfdb1e6e42fdd00a7eba9b860b2f0eccb8"></a>

## namespace property — Property reference / 01dbbc76da25 / 9

Type: `"string"`. Optional, Computed.

Namespace where the K8SCluster exists.

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

- [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-cc6606ecbd98efd86a2215cbabfea57755ac50fba68387b147c84ef4517ce150): complete subsection reference.

- [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-20829d31ac0ba309765427eb42d794858936955253d43c3ec2eff9fe3e9a3342): complete subsection reference.

- [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-3c3f200eed62302bcd00e2bdc7aedead002766dc4c5654d4ea9bd4818eb95a5c): complete subsection reference.

- [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-325cb33df7965df05912659d37416b94640a5e1b4ce76ae18e5d9bf486ae93b3): complete subsection reference.

- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-f995f7b1bd68aaf986b57c5348b214c832bf9f45b921c44d8669c67add4083fe): complete subsection reference.

- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-a4d967e0982944c5fedab2419e216b8fdc62c8b1427b2585fc696ffeb68b0ac8): complete subsection reference.

- [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-c6b0becb60105bb78d7eddb05167bf53348b50913b51147940701b9f1292bc38): complete subsection reference.

- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2a4d59d24a327815b184f45187c189a7b8a53806af7a4177c894065d674c943c): complete subsection reference.

- [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-cf9c596d759cdb7a3dc8914e0096acc67a42ba6eb26bf18fddf9eaac0c95061b): complete subsection reference.

- [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-4c6382ce195526069aa4491dcb3a2d1af2260bf99b382c29dc50770f663f570c): complete subsection reference.

- [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-5145aea8589f27ace674e7790f7e106945a8ce7dd91d049e9df2850ec1dfe2a8): complete subsection reference.

- [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-3db9f9e06f28dff62c1ec1d86fb879948c9ba58293936b1e462d04221b426cef): complete subsection reference.

- [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-02fc479f0a14b8d6ed6954b31d30542bef2b918cd2367aff7bc5628ecd804409): complete subsection reference.

- [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-0c3710b30e4ec94fffdfae7def1580a088a9904fd8c57c8d076568509dc47078): complete subsection reference.

<a id="canonical-ef2772c13e2a7eee7b1a3a32199ddb77692076b03bc296868313d4fc90d067fc"></a>

## All schema paths — Property reference / 01dbbc76da25 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_cluster--reference--group-001.md#canonical-8cce7a9b3629da85a9b381cf2d38268689e4b61361f3080563f4bc3d2c5a7e8c) |
| `cluster_scoped_access_deny` | [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-6b13487909ed576d775f70be2c3389cec04957d6517939bf52d18f78e1aa36da) |
| `cluster_scoped_access_permit` | [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-f8fb42e1f62594a7ec698b96aedfa0dbf97fa0d6c7ba21fb3582099c913cbee9) |
| `cluster_wide_app_list` | [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-24aef261c53080e97e9530a6472114ccc05d70678e8c03a334d9274815673275) |
| `cluster_wide_app_list.cluster_wide_apps` | [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-b8097219245fb7b2edf902d4990c161de0b3b0def81c0cd3913ec383a10ffa63) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd` | [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-4c37295f99af822fcd62d0dba30a6446df51751c41bd14ec4eef4a9fa611f3ee) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-297b10ad5eab304192f364f7f376357f7d5855d85add3d674213f77ac32698e0) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-a45cfec45be5f908d5ca35d950cfd8e1a0dd02bea9fdbd64cef7a8afd5bddd8f) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-e2c25d02f3a767d47981a47d53fee4e55b582359ad47e10659086ed50b9497a9) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-4c6f527051b45ff140dc945b3b6b1e13a8fa33f56a2de4733b6e34cef2bc996b) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-bfeb1fa3b02ce7b166a462d6b2160258b83194df1abec3c66f6e33da2de1a7c8) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider](data-sources--k8s_cluster--reference--group-001.md#canonical-a17d2901794ed809b15222c6e6359c1c0757d51d6a90efc89a12571e48bf7cd9) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location](data-sources--k8s_cluster--reference--group-001.md#canonical-b8d457f9645de5645004880e61de193f0cdd702237b27a6331b2a10f612a1238) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider](data-sources--k8s_cluster--reference--group-001.md#canonical-17c583b3dd9a9b80cdd6875b39a715c5558f857322b370ed5ae47851b52a52b8) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-279c5442b5487039bdae7138be07c2bf047db36f30726d659278fb844565cf72) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref](data-sources--k8s_cluster--reference--group-001.md#canonical-b68e5339aab2b197d6ec8f08fa836900e4114d4587cb412c3e335f2ae5a14b19) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url](data-sources--k8s_cluster--reference--group-001.md#canonical-f3fad457a0d3b00da992445ac9971948b85cc92cb2b80588c75750ccb13c0deb) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port](data-sources--k8s_cluster--reference--group-001.md#canonical-da37edc3a7bd98171da11a0ce87d78b88a44d6c62b151467b20ef6f60e71d6be) |
| `cluster_wide_app_list.cluster_wide_apps.dashboard` | [cluster_wide_app_list.cluster_wide_apps.dashboard](data-sources--k8s_cluster--reference--group-001.md#canonical-320fec39b2bf2036432162bac337ec1ddf8de3c2cd09f6bf88ef49189246a2a2) |
| `cluster_wide_app_list.cluster_wide_apps.metrics_server` | [cluster_wide_app_list.cluster_wide_apps.metrics_server](data-sources--k8s_cluster--reference--group-001.md#canonical-b269eaca2398a7574658378cdb1ef61b68c2e2f6e6f7c381862d9e8f682315ce) |
| `cluster_wide_app_list.cluster_wide_apps.prometheus` | [cluster_wide_app_list.cluster_wide_apps.prometheus](data-sources--k8s_cluster--reference--group-001.md#canonical-b96923a71bb1bec53e18b2deb46e642431ed23c6d20f233d2e3570ca36acd480) |
| `description` | [description](data-sources--k8s_cluster--reference--group-001.md#canonical-1a3e750b434bc109ed58f759d4f656a8765cb7dc35e588fcb420aa0416d6f58b) |
| `global_access_enable` | [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-cbf984a7d624393dd776d83156626069a50c0623bcbc4c9fd9efbf9ab11a060b) |
| `id` | [id](data-sources--k8s_cluster--reference--group-001.md#canonical-2285fce920407dd5de75f2c55cbe757f14a33d7f4bccf8c3c8e7e4fc60913535) |
| `insecure_registry_list` | [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-cfdfd7ec6668cf010f5872167ef6a45407407fde1c91e99e72e728d23d8956e6) |
| `insecure_registry_list.insecure_registries` | [insecure_registry_list.insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-bb34f5ea3485fdcc26df615aef4495a3b3ae9022fd81ba4df3770caf94d40409) |
| `labels` | [labels](data-sources--k8s_cluster--reference--group-001.md#canonical-08795068e152011049494c604e9cdd5a6aa7c2884c24fa188438ad6c964b6403) |
| `local_access_config` | [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-9aeef4b9eaa9d0e0c6cd665a7e45481eec85f84f8ab9b1d40f1adaa4a7c923bf) |
| `local_access_config.default_port` | [local_access_config.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-35f145cd8f709c6be6d7b7201adf26fb456c120d0228281051065e6c3cba0033) |
| `local_access_config.local_domain` | [local_access_config.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-4d406acc5e0bbb4d045bf6e86b11319ac01b70c72b0d195d4742044eaee4cd2a) |
| `local_access_config.port` | [local_access_config.port](data-sources--k8s_cluster--reference--group-001.md#canonical-4bbd778bf2b619961a1d791eadf6346ff2e6d8a35241ae3a581f4f7eea962340) |
| `name` | [name](data-sources--k8s_cluster--reference--group-001.md#canonical-e0e8f087754f035e1d9102f48f8cd76b47df0c30bd33dde80c62eb9c92f937c7) |
| `namespace` | [namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-c13c7bbbc7c4be0c26ec8b9993dea47b1366dabea5f5b40ea682872dcfc5d12d) |
| `no_cluster_wide_apps` | [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-9909cbfa9dc3ce712b8671a7e7292781289f3068d03bae372e23613c4875ca23) |
| `no_global_access` | [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-ece2e8ee6bc06ca4835839561774b407a646235c9c17ff7f76e389ce7561ced2) |
| `no_insecure_registries` | [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-598cc5152b8eb7bd1e4a2023880e04ff3b796730bd85be1471640d4f4e96474b) |
| `no_local_access` | [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0d17981c2fed5a324520ef2c6cd9648fba3b75d3cd6ffa1f4b291c59f6a0ea8b) |
| `use_custom_cluster_role_bindings` | [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-ffa5c7967d4c7a14b77c5b43827a1174a031b66374282493d85b49e0c2aa6d5a) |
| `use_custom_cluster_role_bindings.cluster_role_bindings` | [use_custom_cluster_role_bindings.cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-a934007be5765cbfb8af91db30df66b35be0955077d9642c221d3cbb8b7debd0) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.name` | [use_custom_cluster_role_bindings.cluster_role_bindings.name](data-sources--k8s_cluster--reference--group-001.md#canonical-cec138b144dc33774fed30165034371de8bfecf2ef7bda4c0bcc57f85b74cdce) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.namespace` | [use_custom_cluster_role_bindings.cluster_role_bindings.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-f37405628aa82a7e29401261ff27593d16aa173f1eb9607a703def8ae729b933) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.tenant` | [use_custom_cluster_role_bindings.cluster_role_bindings.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-498b55a0e50e52ddbdcc9474ac91df07b90df569c01da4400cff45a4739e4d42) |
| `use_custom_cluster_role_list` | [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-a6bd6244ed93bc499cdfb5894726207631155dd6caea5cd10e730c23a9447fbd) |
| `use_custom_cluster_role_list.cluster_roles` | [use_custom_cluster_role_list.cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-f5e053f457ed7ee5bb9e252e802f545f2c83ee431800688f07556f9a12d5a156) |
| `use_custom_cluster_role_list.cluster_roles.name` | [use_custom_cluster_role_list.cluster_roles.name](data-sources--k8s_cluster--reference--group-001.md#canonical-fb762f006e1f81d517424e06e7c82b0ce35fe76c43f86e04cac8ae727b08230d) |
| `use_custom_cluster_role_list.cluster_roles.namespace` | [use_custom_cluster_role_list.cluster_roles.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-5d2863cdbce829e2c4bcc2558684a12a5debbac273709d11c28965dd40cc23e0) |
| `use_custom_cluster_role_list.cluster_roles.tenant` | [use_custom_cluster_role_list.cluster_roles.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-4bcda5e6b6aab53cb02d8ed89a799888157beb86d26a1038524b7f7dbd662c6e) |
| `use_custom_pod_security_admission` | [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-c7c3a5a682e4a6769b6ec100bede148438ac6e4118d434bf4abb879dca0434dc) |
| `use_custom_pod_security_admission.name` | [use_custom_pod_security_admission.name](data-sources--k8s_cluster--reference--group-001.md#canonical-fe7ba1892de7cb806ef5e708286d11402f1e5ed002c3c15f0cd9fe89136abcde) |
| `use_custom_pod_security_admission.namespace` | [use_custom_pod_security_admission.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-3059ffc0426a7360512af4bf71789b5f1b7ee103bb98bb1c414a42e86aca3a23) |
| `use_custom_pod_security_admission.tenant` | [use_custom_pod_security_admission.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-225da710dc2b88080b8fccd0d642a7d8b05a63c4fbf29e7eb9de5d55e307daa0) |
| `use_custom_psp_list` | [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-f52e46283aa8b946a4c27fe2934931b51dbb39ad463b07c918871af632b51d12) |
| `use_custom_psp_list.pod_security_policies` | [use_custom_psp_list.pod_security_policies](data-sources--k8s_cluster--reference--group-001.md#canonical-0a015eb35fe74aa5761ef242a747d47663a1dab30e4b7214cf714b004eed874f) |
| `use_custom_psp_list.pod_security_policies.name` | [use_custom_psp_list.pod_security_policies.name](data-sources--k8s_cluster--reference--group-001.md#canonical-8eb77aac0ec459438d6c5624e12d04f21f903f9517ac619f88fd2430be2d8ab1) |
| `use_custom_psp_list.pod_security_policies.namespace` | [use_custom_psp_list.pod_security_policies.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-1b5804f719212b36041ca78e4ff8e821bd9a00d465abe00b8393ee21df604b82) |
| `use_custom_psp_list.pod_security_policies.tenant` | [use_custom_psp_list.pod_security_policies.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-4fefe001a3f79ea76a105d6982981e9aeaee19bfac039fa848bd9447f48d4a58) |
| `use_default_cluster_role_bindings` | [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-af713c8b740253d576d41652e94614ccc5d964c104381b6ad30fb6388f06421f) |
| `use_default_cluster_roles` | [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-67c7fd0a40300de8623ff4f650719763ee4d558d3c06dd92862622d68e86ba06) |
| `use_default_pod_security_admission` | [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-8d78b4ca73138786ee3f8e084cf383a43840607f145bb5cd1a7b9a1c23f321e0) |
| `use_default_psp` | [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-8427e54e9d835c49e69c01358b362e06fa1162ef763d265f1cb0f346c9a9e438) |
| `vk8s_namespace_access_deny` | [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-cc02f59fd5ed161634fa19227fe7afb0b65ec345db095353a18495bba11e1adc) |
| `vk8s_namespace_access_permit` | [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-00ed27b0d62b8a7e7c0d969925d44ba79fab620f94b041895a4b461b57828ccf) |

<a id="canonical-49084c4d291513b9405dcaee37cf91b5da2ecb3f6f8baa693ba7ffb5c7e8009d"></a>

## Next pages — Property reference / 01dbbc76da25 / 11

- [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-180bdcc884989526420137adf6d76588eae18458ff269385040e756e71ce8c60)
- [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-419ca709b10b3d101be75663dfb5543c8967c02ae0f4fd75e5506540ef11efa2)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-f8ab74e4db9b790e1f33596a4418bad9fb7a98a4193dd7694c8a7dada64d60a8)
- [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-8033d79e8c97d5fce86173ed9a92916d2782f4deb719c478fa22f184112ef7c9)
- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-6e733ff29bc886c146c986108d51341f7f0a592f241e9bff8794f81c9a01367d)
- [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-cc6606ecbd98efd86a2215cbabfea57755ac50fba68387b147c84ef4517ce150)
- [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-20829d31ac0ba309765427eb42d794858936955253d43c3ec2eff9fe3e9a3342)
- [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-3c3f200eed62302bcd00e2bdc7aedead002766dc4c5654d4ea9bd4818eb95a5c)
- [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-325cb33df7965df05912659d37416b94640a5e1b4ce76ae18e5d9bf486ae93b3)
- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-f995f7b1bd68aaf986b57c5348b214c832bf9f45b921c44d8669c67add4083fe)
- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-a4d967e0982944c5fedab2419e216b8fdc62c8b1427b2585fc696ffeb68b0ac8)
- [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-c6b0becb60105bb78d7eddb05167bf53348b50913b51147940701b9f1292bc38)
- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2a4d59d24a327815b184f45187c189a7b8a53806af7a4177c894065d674c943c)
- [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-cf9c596d759cdb7a3dc8914e0096acc67a42ba6eb26bf18fddf9eaac0c95061b)
- [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-4c6382ce195526069aa4491dcb3a2d1af2260bf99b382c29dc50770f663f570c)
- [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-5145aea8589f27ace674e7790f7e106945a8ce7dd91d049e9df2850ec1dfe2a8)
- [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-3db9f9e06f28dff62c1ec1d86fb879948c9ba58293936b1e462d04221b426cef)
- [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-02fc479f0a14b8d6ed6954b31d30542bef2b918cd2367aff7bc5628ecd804409)
- [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-0c3710b30e4ec94fffdfae7def1580a088a9904fd8c57c8d076568509dc47078)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-180bdcc884989526420137adf6d76588eae18458ff269385040e756e71ce8c60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdb94836105ed33500e6f9961e7b4450ba2d9ac73fd0e7a8ba2ded2ded6c8336"></a>

## cluster_scoped_access_deny — cluster_scoped_access_deny / 99ebdd36bcb5 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- cluster_scoped_access_deny

<a id="canonical-6b13487909ed576d775f70be2c3389cec04957d6517939bf52d18f78e1aa36da"></a>

Type: `["object", {}]`. Computed.

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

- [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-6b13487909ed576d775f70be2c3389cec04957d6517939bf52d18f78e1aa36da)
- [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-f8fb42e1f62594a7ec698b96aedfa0dbf97fa0d6c7ba21fb3582099c913cbee9)

Select alternatives according to the provider validators above.

<a id="canonical-3ec0854627cb2d4af6bc35fb6cd25a571f7b389917f965cbc962334930fcb0e1"></a>

## Direct properties — cluster_scoped_access_deny / 99ebdd36bcb5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e67a0aa854739f87baf4d3ea75a6aa822e84bee4c93b2edbeed6ecda9e37d608"></a>

## Next pages — cluster_scoped_access_deny / 99ebdd36bcb5 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-419ca709b10b3d101be75663dfb5543c8967c02ae0f4fd75e5506540ef11efa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e2bc2135d8255db11a0849a1ea51df8aa8601b9a7e8d10fbc73f41610b7b45c"></a>

## cluster_scoped_access_permit — cluster_scoped_access_permit / c535e07153c8 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- cluster_scoped_access_permit

<a id="canonical-f8fb42e1f62594a7ec698b96aedfa0dbf97fa0d6c7ba21fb3582099c913cbee9"></a>

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

<a id="canonical-22a9f788f5128346960fbcc23a32d1b50cf6c4de24852b6d183b7f3f1b02db08"></a>

## Direct properties — cluster_scoped_access_permit / c535e07153c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b14fce0f4a9888067cfdbfa6b3201a8ab13c4e9847e018922e0b6fc5026048a"></a>

## Next pages — cluster_scoped_access_permit / c535e07153c8 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-600bfdf9cd4b4b2d12bf4729c8f3d815c4481ab1ba5b7de6df0e918d8a1d987d"></a>

## cluster_wide_app_list — cluster_wide_app_list / 08f7c8352946 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- cluster_wide_app_list

<a id="canonical-24aef261c53080e97e9530a6472114ccc05d70678e8c03a334d9274815673275"></a>

Type: `"single"`. Computed.

\[OneOf: cluster\_wide\_app\_list, no\_cluster\_wide\_apps; Default: no\_cluster\_wide\_apps\]
Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

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

- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-24aef261c53080e97e9530a6472114ccc05d70678e8c03a334d9274815673275)
- [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-9909cbfa9dc3ce712b8671a7e7292781289f3068d03bae372e23613c4875ca23)

Select alternatives according to the provider validators above.

<a id="canonical-a978011420ed214e206aebee9ddc74813dbab4c93a7937d763c148e68db90fde"></a>

## Direct properties — cluster_wide_app_list / 08f7c8352946 / 3

- [cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea): complete subsection reference.

<a id="canonical-2c58f4713ee1df1c9aa8bb71a3bf4c7c49abcdde0cd3cc43e8623d6683d36744"></a>

## Next pages — cluster_wide_app_list / 08f7c8352946 / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232dc9c791fa6f5b86f862fd8e69353631dc607986c14710c84fe2ccceebb50"></a>

## cluster_wide_app_list.cluster_wide_apps — cluster_wide_app_list.cluster_wide_apps / 928cc199efdb / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- cluster_wide_app_list.cluster_wide_apps

<a id="canonical-b8097219245fb7b2edf902d4990c161de0b3b0def81c0cd3913ec383a10ffa63"></a>

Type: `"list"`. Computed.

Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

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

<a id="canonical-9f90d8dcace39a11e8ad0bd89f8b2c27cff1fc8b287743fc7ec36b946b78685d"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps / 928cc199efdb / 3

- [argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a): complete subsection reference.

- [dashboard](data-sources--k8s_cluster--reference--group-001.md#canonical-6e6aed4598bd5a1644bdbf6bc6b0fc1494f6b805cea6b719f63b390a279b8bc0): complete subsection reference.

- [metrics_server](data-sources--k8s_cluster--reference--group-001.md#canonical-fbfbdc722d544e7648aa7b02d99350d47936bbdcfb6953fddbb5871ed2bcad81): complete subsection reference.

- [prometheus](data-sources--k8s_cluster--reference--group-001.md#canonical-68822c2dca3b308c5e7cfe13e591dca2725877cf7fdf4cf32d4c3f44331c4197): complete subsection reference.

<a id="canonical-a1e024411ea64f9108cbd0d87d06c98225204ec82ef3040523a4d28554e76d8a"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps / 928cc199efdb / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a)
- [cluster_wide_app_list.cluster_wide_apps.dashboard](data-sources--k8s_cluster--reference--group-001.md#canonical-6e6aed4598bd5a1644bdbf6bc6b0fc1494f6b805cea6b719f63b390a279b8bc0)
- [cluster_wide_app_list.cluster_wide_apps.metrics_server](data-sources--k8s_cluster--reference--group-001.md#canonical-fbfbdc722d544e7648aa7b02d99350d47936bbdcfb6953fddbb5871ed2bcad81)
- [cluster_wide_app_list.cluster_wide_apps.prometheus](data-sources--k8s_cluster--reference--group-001.md#canonical-68822c2dca3b308c5e7cfe13e591dca2725877cf7fdf4cf32d4c3f44331c4197)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d8cd99fbeb3e359f4446df90974d3619813d838ea0d493e747e14f9ad48470f"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd — cluster_wide_app_list.cluster_wide_apps.argo_cd / c17c3c129f40 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- cluster_wide_app_list.cluster_wide_apps.argo_cd

<a id="canonical-4c37295f99af822fcd62d0dba30a6446df51751c41bd14ec4eef4a9fa611f3ee"></a>

Type: `"single"`. Computed.

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

<a id="canonical-ba6232ace97943ed64fe7bb848838211fe5e7adebdb69739603ad1af257b2992"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd / c17c3c129f40 / 3

- [local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb): complete subsection reference.

<a id="canonical-7b8a98bd95943f43983d83aa4008867fb758efcb8ac36990b45a67e38a35adfc"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd / c17c3c129f40 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84011bb9665e90911b82550baa229a18580c4f5b0e256facf5f7c9d3b9809a9d"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / 767b84c8f772 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

<a id="canonical-297b10ad5eab304192f364f7f376357f7d5855d85add3d674213f77ac32698e0"></a>

Type: `"single"`. Computed.

Parameters required to enable local access.

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

<a id="canonical-91f07fd9616a762aa47ca29ecc676fb46d5932c9acfa590bf0fa54a9a479c674"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / 767b84c8f772 / 3

- [default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-bfc940b9558ed09ca2b5d9410581673aa2b059385d8fee0db1ba9a09006a0fd5): complete subsection reference.

<a id="canonical-e2c25d02f3a767d47981a47d53fee4e55b582359ad47e10659086ed50b9497a9"></a>

<a id="canonical-a72030fafd35afd9ebffd3edb15ef1865122886ff4487d3c57cce18768740ad7"></a>

## local_domain property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / 767b84c8f772 / 4

Type: `"string"`. Computed.

ArgoCD will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

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

- [password](data-sources--k8s_cluster--reference--group-001.md#canonical-50140e4ae2bef7e89389fe3c45748061d7c6c358a7cb7dc0813cb3f3bf8e6006): complete subsection reference.

<a id="canonical-da37edc3a7bd98171da11a0ce87d78b88a44d6c62b151467b20ef6f60e71d6be"></a>

<a id="canonical-394f246ea73f60e17c7334e3be08473dcc162adb7a2537e3c671876cc4762331"></a>

## port property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / 767b84c8f772 / 5

Type: `"number"`. Computed.

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

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

<a id="canonical-736126496b14c9dae55c8a7396d653d7df31ed85ea11cf7a4a49ec590ee43b65"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain / 767b84c8f772 / 6

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-bfc940b9558ed09ca2b5d9410581673aa2b059385d8fee0db1ba9a09006a0fd5)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-50140e4ae2bef7e89389fe3c45748061d7c6c358a7cb7dc0813cb3f3bf8e6006)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-bfc940b9558ed09ca2b5d9410581673aa2b059385d8fee0db1ba9a09006a0fd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-319f1361b6908bddd91c13d40e5b59860ced1fce6e9477cd1459a16fef2b4b3f"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port / af796ad8db00 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port

<a id="canonical-a45cfec45be5f908d5ca35d950cfd8e1a0dd02bea9fdbd64cef7a8afd5bddd8f"></a>

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

<a id="canonical-66287641a45634e9cf2d2ca80d7ba0d698a7688d3d75d86bc2e0f927d37b3b46"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port / af796ad8db00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20b5e250094c4d85d89c60d15237195d371aa83fd2df3e5b4a9d96bdd51a2a56"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port / af796ad8db00 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-50140e4ae2bef7e89389fe3c45748061d7c6c358a7cb7dc0813cb3f3bf8e6006"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdf2d3b264ad924ccd7c8213b365f678844ad5871a6f7a9662b0ad6e5c37b87f"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password / e78d21b16c61 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password

<a id="canonical-4c6f527051b45ff140dc945b3b6b1e13a8fa33f56a2de4733b6e34cef2bc996b"></a>

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

<a id="canonical-26e0461d0f185ee99b7a49f204dd93a5eb06cd18c4f675117fff496c926eae1b"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password / e78d21b16c61 / 3

- [blindfold_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-729e43e3964dc7a0fbb1c2ce94bce3ef55b18f498760240c0970c5b7b791f331): complete subsection reference.

- [clear_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-bbe107ad03e135acbc288b7152f73818c8e787fc1da51838283cfe7487f44d26): complete subsection reference.

<a id="canonical-33c7f3addc68a2315cfd2e0ca165df616cb76babd7e624532595b07fd8cf1d08"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password / e78d21b16c61 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-729e43e3964dc7a0fbb1c2ce94bce3ef55b18f498760240c0970c5b7b791f331)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-bbe107ad03e135acbc288b7152f73818c8e787fc1da51838283cfe7487f44d26)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-729e43e3964dc7a0fbb1c2ce94bce3ef55b18f498760240c0970c5b7b791f331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1062fe492d9a31f53714d08f70a420e67bdf8b74d9143aef57899129d31cb5f1"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 992c04e20c9c / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-50140e4ae2bef7e89389fe3c45748061d7c6c358a7cb7dc0813cb3f3bf8e6006)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info

<a id="canonical-bfeb1fa3b02ce7b166a462d6b2160258b83194df1abec3c66f6e33da2de1a7c8"></a>

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

<a id="canonical-17b0a50efcdc2f7761066184cf9710c0e5cd9312f98ab780a9035d92c9edde9e"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 992c04e20c9c / 3

<a id="canonical-a17d2901794ed809b15222c6e6359c1c0757d51d6a90efc89a12571e48bf7cd9"></a>

<a id="canonical-16375a243cf9c87fcedef39e4f478048ef03d2d92cc7186b0a93b2a3f858ffea"></a>

## decryption_provider property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 992c04e20c9c / 4

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

<a id="canonical-b8d457f9645de5645004880e61de193f0cdd702237b27a6331b2a10f612a1238"></a>

<a id="canonical-bfd4a1273b297cf03e3297acca2fd077cd19b76d9b9b48e7a8cb0acc29593918"></a>

## location property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 992c04e20c9c / 5

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

<a id="canonical-17c583b3dd9a9b80cdd6875b39a715c5558f857322b370ed5ae47851b52a52b8"></a>

<a id="canonical-acc23fe055a1711cb819a0f84b37dc2c3c41bdffe3178c273ad3fbf0c667b95e"></a>

## store_provider property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 992c04e20c9c / 6

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

<a id="canonical-d0f5197ba3ebde3c3e9301ee67ab3a319c1967a7210568f63c6bbde8a83bb22a"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_ / 992c04e20c9c / 7

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-50140e4ae2bef7e89389fe3c45748061d7c6c358a7cb7dc0813cb3f3bf8e6006)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-bbe107ad03e135acbc288b7152f73818c8e787fc1da51838283cfe7487f44d26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a51c6e9d2104435d9547eefb380d942672215edb1a6bb9dd9ac735b68b6e89b9"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / fb751f868ae6 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1fee4f61cbfd3c2399aa7ea2f95611d14fa4cabae37c87c57edac70bb270530a)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1aba6b60c68aff6ca2141eec58d7bc16b97fe042b1d06e1987a978cea0f01fcb)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-50140e4ae2bef7e89389fe3c45748061d7c6c358a7cb7dc0813cb3f3bf8e6006)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info

<a id="canonical-279c5442b5487039bdae7138be07c2bf047db36f30726d659278fb844565cf72"></a>

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

<a id="canonical-8062feb98ff14dec37339c7e8a7fd8392307e8f655a4068edf034d4e48cc23ad"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / fb751f868ae6 / 3

<a id="canonical-b68e5339aab2b197d6ec8f08fa836900e4114d4587cb412c3e335f2ae5a14b19"></a>

<a id="canonical-bd5bf266c9df0360dabcfde4829ec73e664f8ecab343719ba86258e031ac671e"></a>

## provider_ref property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / fb751f868ae6 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f3fad457a0d3b00da992445ac9971948b85cc92cb2b80588c75750ccb13c0deb"></a>

<a id="canonical-a009d8b4c3b0aaa1adf513bf874bfb1ea30b4ccc592aa6455704d97bbf08d50a"></a>

## url property — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / fb751f868ae6 / 5

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

<a id="canonical-083c81bb5cc37be8a70a344a1e4cf0bd110f3f70f475a0f98c4bd3af78aee5dd"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secr / fb751f868ae6 / 6

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-50140e4ae2bef7e89389fe3c45748061d7c6c358a7cb7dc0813cb3f3bf8e6006)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-6e6aed4598bd5a1644bdbf6bc6b0fc1494f6b805cea6b719f63b390a279b8bc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-040433a8742b19430f239827eccd024f578ac54e92d03470cd0f2f25a657298d"></a>

## cluster_wide_app_list.cluster_wide_apps.dashboard — cluster_wide_app_list.cluster_wide_apps.dashboard / 61c841a83f08 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- cluster_wide_app_list.cluster_wide_apps.dashboard

<a id="canonical-320fec39b2bf2036432162bac337ec1ddf8de3c2cd09f6bf88ef49189246a2a2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-66596edb145807c22f1065cc01a58ddec5105b6219c73b0dd7a4d145139e5b92"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.dashboard / 61c841a83f08 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-37222fa7f8c6d171c3c59887ef6220eb032590370a6333a5ab87af392bd8f4a3"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.dashboard / 61c841a83f08 / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-fbfbdc722d544e7648aa7b02d99350d47936bbdcfb6953fddbb5871ed2bcad81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fd6e1ef4db4f563f5fe62a77b53ce5795a32a8fc13eb396016976f40416786f"></a>

## cluster_wide_app_list.cluster_wide_apps.metrics_server — cluster_wide_app_list.cluster_wide_apps.metrics_server / 46019e757b0b / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- cluster_wide_app_list.cluster_wide_apps.metrics_server

<a id="canonical-b269eaca2398a7574658378cdb1ef61b68c2e2f6e6f7c381862d9e8f682315ce"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-70273cf3fe51a77d592fa5a296b8cb6d435dd4ce22b9dd4acc1b674ceb79963a"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.metrics_server / 46019e757b0b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64d2c5acea6ab5d9889a591fe61481d7f359dfac205f1bade3f429920ac9f4d1"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.metrics_server / 46019e757b0b / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-68822c2dca3b308c5e7cfe13e591dca2725877cf7fdf4cf32d4c3f44331c4197"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95815dfd7a14591add4bf77c2f45400ff022512d439a18c458ddacef4039e1db"></a>

## cluster_wide_app_list.cluster_wide_apps.prometheus — cluster_wide_app_list.cluster_wide_apps.prometheus / ea4ac95af2ae / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-c60a62e63621a5324d03ac465fc4016c57a8e3ae97180fc6cd4783d88f58709e)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- cluster_wide_app_list.cluster_wide_apps.prometheus

<a id="canonical-b96923a71bb1bec53e18b2deb46e642431ed23c6d20f233d2e3570ca36acd480"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-9a52cc35e0d6159957bfff7e1319c71cb1dbd62727dc3c54192e6465e5032181"></a>

## Direct properties — cluster_wide_app_list.cluster_wide_apps.prometheus / ea4ac95af2ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aea1829a0984863e8d9dae91e0df83be91f46cc691e67657f7afd9ef12458b9c"></a>

## Next pages — cluster_wide_app_list.cluster_wide_apps.prometheus / ea4ac95af2ae / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-757985ed680ef736e05359f5609ab61392efe546f5959b6a9f4157787fcaa8ea)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-f8ab74e4db9b790e1f33596a4418bad9fb7a98a4193dd7694c8a7dada64d60a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d326a61384e93c302d49b5a92b53ad2e29c93acf44226ee076d79ca885d2f7d"></a>

## global_access_enable — global_access_enable / 3e8feffc93c9 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- global_access_enable

<a id="canonical-cbf984a7d624393dd776d83156626069a50c0623bcbc4c9fd9efbf9ab11a060b"></a>

Type: `["object", {}]`. Computed.

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

- [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-cbf984a7d624393dd776d83156626069a50c0623bcbc4c9fd9efbf9ab11a060b)
- [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-ece2e8ee6bc06ca4835839561774b407a646235c9c17ff7f76e389ce7561ced2)

Select alternatives according to the provider validators above.

<a id="canonical-714ab65a79674d5124e379cf2aa83ce05a5c5b59ed1163a9a0f39ef5910295e0"></a>

## Direct properties — global_access_enable / 3e8feffc93c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8ba4c8cd9b92ec750e3770ea8abfecdf0520868d79a6fde2610b0d00fa1274a"></a>

## Next pages — global_access_enable / 3e8feffc93c9 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-8033d79e8c97d5fce86173ed9a92916d2782f4deb719c478fa22f184112ef7c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3872d07d7c482c70c18ca425d13e6cd4e18db39f9b905e841fcb636c7a3f8000"></a>

## insecure_registry_list — insecure_registry_list / ef5e2b86cac8 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- insecure_registry_list

<a id="canonical-cfdfd7ec6668cf010f5872167ef6a45407407fde1c91e99e72e728d23d8956e6"></a>

Type: `"single"`. Computed.

\[OneOf: insecure\_registry\_list, no\_insecure\_registries; Default: no\_insecure\_registries\]
Docker Insecure Registry List. List of docker insecure registries.

Upstream description:

List of docker insecure registries.

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

- [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-cfdfd7ec6668cf010f5872167ef6a45407407fde1c91e99e72e728d23d8956e6)
- [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-598cc5152b8eb7bd1e4a2023880e04ff3b796730bd85be1471640d4f4e96474b)

Select alternatives according to the provider validators above.

<a id="canonical-006604691f195207e7abb6ef874168326fb34f876fcb0fbbe0ea7bda871a94a7"></a>

## Direct properties — insecure_registry_list / ef5e2b86cac8 / 3

<a id="canonical-bb34f5ea3485fdcc26df615aef4495a3b3ae9022fd81ba4df3770caf94d40409"></a>

<a id="canonical-301daf254346afcef1e17e2cfeba043937d53bd6cff3cab71266535c986243a1"></a>

## insecure_registries property — insecure_registry_list / ef5e2b86cac8 / 4

Type: `["list", "string"]`. Computed.

List of docker insecure registries in format 'example.com:5000'.

Upstream description:

List of docker insecure registries in format "example.com:5000"

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

<a id="canonical-21c4f7e552c8d17a2f34608e4ce9eeeefdb58e06a82660d5bfc0a2a670a3476f"></a>

## Next pages — insecure_registry_list / ef5e2b86cac8 / 5

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-6e733ff29bc886c146c986108d51341f7f0a592f241e9bff8794f81c9a01367d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-043a6f3f20de7e5d54ba9d92a52b49f4b66b194419833a5e86b90b15cb1c877f"></a>

## local_access_config — local_access_config / ef55438fe264 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- local_access_config

<a id="canonical-9aeef4b9eaa9d0e0c6cd665a7e45481eec85f84f8ab9b1d40f1adaa4a7c923bf"></a>

Type: `"single"`. Computed.

\[OneOf: local\_access\_config, no\_local\_access; Default: no\_local\_access\] Parameters required
to enable local access.

Upstream description:

Parameters required to enable local access.

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

- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-9aeef4b9eaa9d0e0c6cd665a7e45481eec85f84f8ab9b1d40f1adaa4a7c923bf)
- [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0d17981c2fed5a324520ef2c6cd9648fba3b75d3cd6ffa1f4b291c59f6a0ea8b)

Select alternatives according to the provider validators above.

<a id="canonical-ea6f70aeadf0cf164c818615f3e6992f4d679c13244682212dd54f10907c7cac"></a>

## Direct properties — local_access_config / ef55438fe264 / 3

- [default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-2a21d0f17b075f45029cd91af5c030e7378ec9ab3e0d9dc76c179bd50a31e0d9): complete subsection reference.

<a id="canonical-4d406acc5e0bbb4d045bf6e86b11319ac01b70c72b0d195d4742044eaee4cd2a"></a>

<a id="canonical-7606c93370f1f2ef538e8f4cafea8d46eb94a5a447731bd80f63ad80ece50404"></a>

## local_domain property — local_access_config / ef55438fe264 / 4

Type: `"string"`. Computed.

Local K8s API server will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

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

<a id="canonical-4bbd778bf2b619961a1d791eadf6346ff2e6d8a35241ae3a581f4f7eea962340"></a>

<a id="canonical-84048acdf4aeebcbb56735190148f8a1779be640649fc2a25f4ef8c5f1dadf33"></a>

## port property — local_access_config / ef55438fe264 / 5

Type: `"number"`. Computed.

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

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

<a id="canonical-f2f7b0151058743799d537e031cf266a2e9d882dd8f50bf3174ace2d1b727184"></a>

## Next pages — local_access_config / ef55438fe264 / 6

- [local_access_config.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-2a21d0f17b075f45029cd91af5c030e7378ec9ab3e0d9dc76c179bd50a31e0d9)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-2a21d0f17b075f45029cd91af5c030e7378ec9ab3e0d9dc76c179bd50a31e0d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad77d059d17bfdcd7671269b5851bcc1b11b207969c7abd90e49c9a04d9ed51f"></a>

## local_access_config.default_port — local_access_config.default_port / 76603dd3d1a7 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-6e733ff29bc886c146c986108d51341f7f0a592f241e9bff8794f81c9a01367d)
- local_access_config.default_port

<a id="canonical-35f145cd8f709c6be6d7b7201adf26fb456c120d0228281051065e6c3cba0033"></a>

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

<a id="canonical-0cfe16511b843c4634e39d754a144996e813685541e985132d592cb4c6634774"></a>

## Direct properties — local_access_config.default_port / 76603dd3d1a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40da8094f1c85b147f5cf7cd24df6d88ae42ea5ff163d7ea9c24103dfc0b2cf3"></a>

## Next pages — local_access_config.default_port / 76603dd3d1a7 / 4

- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-6e733ff29bc886c146c986108d51341f7f0a592f241e9bff8794f81c9a01367d)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-cc6606ecbd98efd86a2215cbabfea57755ac50fba68387b147c84ef4517ce150"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c690165e8ff47fd62746a8d0f2dade3744db0b4921679785ca1d78ae12ae711e"></a>

## no_cluster_wide_apps — no_cluster_wide_apps / 9237052a3aaa / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- no_cluster_wide_apps

<a id="canonical-9909cbfa9dc3ce712b8671a7e7292781289f3068d03bae372e23613c4875ca23"></a>

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

<a id="canonical-450193175205689d6aef6838e691c1bc7a487283f6405d6b1f3914b54717d4f4"></a>

## Direct properties — no_cluster_wide_apps / 9237052a3aaa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3ee1791f44841e95a3531011232d4d27a5c2bd4ab49aaa4d4de254a9479a124"></a>

## Next pages — no_cluster_wide_apps / 9237052a3aaa / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-20829d31ac0ba309765427eb42d794858936955253d43c3ec2eff9fe3e9a3342"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bc862090adf4ca3f5e860b571915386253130da10a1c4f597f6378e7b63b06a"></a>

## no_global_access — no_global_access / 0bd129b00a44 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- no_global_access

<a id="canonical-ece2e8ee6bc06ca4835839561774b407a646235c9c17ff7f76e389ce7561ced2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-25df21b6859d8ff09d85f60f6c0db6b787a167e02e1068ea2c90093286ad84a1"></a>

## Direct properties — no_global_access / 0bd129b00a44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6575448f344283bb4857908043db641009082ac6a3969e95a2c1fcef57373a74"></a>

## Next pages — no_global_access / 0bd129b00a44 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-3c3f200eed62302bcd00e2bdc7aedead002766dc4c5654d4ea9bd4818eb95a5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c658872676ffffbf3a1158c0d597294b30f77a80190348bf3b72d2326b713690"></a>

## no_insecure_registries — no_insecure_registries / 405936c99c47 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- no_insecure_registries

<a id="canonical-598cc5152b8eb7bd1e4a2023880e04ff3b796730bd85be1471640d4f4e96474b"></a>

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

<a id="canonical-0f7cb0a61a1c218b2dc8081ce7cbb92fd883b10ab24aa21cd690f759667e25a2"></a>

## Direct properties — no_insecure_registries / 405936c99c47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-637b0136e59a91716a2bd9edf8f559b33064cfcde2400bbe5ad8f44e3e6e6374"></a>

## Next pages — no_insecure_registries / 405936c99c47 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-325cb33df7965df05912659d37416b94640a5e1b4ce76ae18e5d9bf486ae93b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40e6040f88372dae9e5a25c3c23295723c66ecc4eab2c5a2c84a310f8c0cad3a"></a>

## no_local_access — no_local_access / 69e8c24235d3 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- no_local_access

<a id="canonical-0d17981c2fed5a324520ef2c6cd9648fba3b75d3cd6ffa1f4b291c59f6a0ea8b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0c28d2d85659ad23965168bde28d60e6891a84498c1f5f28cc4fa326a8b9cdd8"></a>

## Direct properties — no_local_access / 69e8c24235d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e224edeb1268d901bf12069a92920fa25f8a54cf21d26a62a06402f253838fe6"></a>

## Next pages — no_local_access / 69e8c24235d3 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-f995f7b1bd68aaf986b57c5348b214c832bf9f45b921c44d8669c67add4083fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a01ad1eedeabc678df71e40962f3cbf1829a14acc0cc1c655841bb5cb71c78a2"></a>

## use_custom_cluster_role_bindings — use_custom_cluster_role_bindings / 6790c1e3c6c3 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_custom_cluster_role_bindings

<a id="canonical-ffa5c7967d4c7a14b77c5b43827a1174a031b66374282493d85b49e0c2aa6d5a"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_cluster\_role\_bindings, use\_default\_cluster\_role\_bindings; Default:
use\_default\_cluster\_role\_bindings\] List of active cluster role binding list for a K8s cluster.

Upstream description:

List of active cluster role binding list for a K8s cluster.

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

- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-ffa5c7967d4c7a14b77c5b43827a1174a031b66374282493d85b49e0c2aa6d5a)
- [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-af713c8b740253d576d41652e94614ccc5d964c104381b6ad30fb6388f06421f)

Select alternatives according to the provider validators above.

<a id="canonical-1c41b7424e9b18c567fd137f12b847838f83abcd13e65df580f4bc38a5410ca0"></a>

## Direct properties — use_custom_cluster_role_bindings / 6790c1e3c6c3 / 3

- [cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-9e99160e66aa63a2624803177d8d354e77f8f16bb045e6052be8d5412484ec5b): complete subsection reference.

<a id="canonical-17a70a3399be05f0da1bb4efaf6589faefaa1c466b4ce5e82a311fb5b6685964"></a>

## Next pages — use_custom_cluster_role_bindings / 6790c1e3c6c3 / 4

- [use_custom_cluster_role_bindings.cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-9e99160e66aa63a2624803177d8d354e77f8f16bb045e6052be8d5412484ec5b)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-9e99160e66aa63a2624803177d8d354e77f8f16bb045e6052be8d5412484ec5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8410af2afac5a9e50c8fb9a920b293fbcc557f569859c93b99d4b1d4ada15dea"></a>

## use_custom_cluster_role_bindings.cluster_role_bindings — use_custom_cluster_role_bindings.cluster_role_bindings / 883300ceb223 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-f995f7b1bd68aaf986b57c5348b214c832bf9f45b921c44d8669c67add4083fe)
- use_custom_cluster_role_bindings.cluster_role_bindings

<a id="canonical-a934007be5765cbfb8af91db30df66b35be0955077d9642c221d3cbb8b7debd0"></a>

Type: `"list"`. Computed.

List of active cluster role binding list for a K8s cluster.

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

<a id="canonical-f21b9b059b0d4ebf3bdb20a95f431bede91d8e6d4b62f186c38754c36bc0c10c"></a>

## Direct properties — use_custom_cluster_role_bindings.cluster_role_bindings / 883300ceb223 / 3

<a id="canonical-cec138b144dc33774fed30165034371de8bfecf2ef7bda4c0bcc57f85b74cdce"></a>

<a id="canonical-d23108395eb56580f7833d3b338840a571ad96331e0fb0918982254ad403f299"></a>

## name property — use_custom_cluster_role_bindings.cluster_role_bindings / 883300ceb223 / 4

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

<a id="canonical-f37405628aa82a7e29401261ff27593d16aa173f1eb9607a703def8ae729b933"></a>

<a id="canonical-1cb8bfc02edd5c318bca94c626b07c740880923043244056d2e4a24aaa3d6e6d"></a>

## namespace property — use_custom_cluster_role_bindings.cluster_role_bindings / 883300ceb223 / 5

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

<a id="canonical-498b55a0e50e52ddbdcc9474ac91df07b90df569c01da4400cff45a4739e4d42"></a>

<a id="canonical-b862b95459c1f7e9117a0cf3681a416bdb52b4fa9d21e9ed9cab515f5f74d978"></a>

## tenant property — use_custom_cluster_role_bindings.cluster_role_bindings / 883300ceb223 / 6

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

<a id="canonical-9cabd1c3d494e9e80a65631bed67f722668f41f664a842f47466db47396bbcd8"></a>

## Next pages — use_custom_cluster_role_bindings.cluster_role_bindings / 883300ceb223 / 7

- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-f995f7b1bd68aaf986b57c5348b214c832bf9f45b921c44d8669c67add4083fe)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-a4d967e0982944c5fedab2419e216b8fdc62c8b1427b2585fc696ffeb68b0ac8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ff0c8f9de4edc4fc6d811feac333fc4e29ad1bfad4a07f2fc5779d380ab9ab2"></a>

## use_custom_cluster_role_list — use_custom_cluster_role_list / 831c64e92c43 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_custom_cluster_role_list

<a id="canonical-a6bd6244ed93bc499cdfb5894726207631155dd6caea5cd10e730c23a9447fbd"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_cluster\_role\_list, use\_default\_cluster\_roles; Default:
use\_default\_cluster\_roles\] List of active cluster role list for a K8s cluster.

Upstream description:

List of active cluster role list for a K8s cluster.

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

- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-a6bd6244ed93bc499cdfb5894726207631155dd6caea5cd10e730c23a9447fbd)
- [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-67c7fd0a40300de8623ff4f650719763ee4d558d3c06dd92862622d68e86ba06)

Select alternatives according to the provider validators above.

<a id="canonical-7a7a8d41d4e08777b331260c354a4362aec27e77658db4eb0e965ee395074e9a"></a>

## Direct properties — use_custom_cluster_role_list / 831c64e92c43 / 3

- [cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-539f180a36e99a61ed7fd0c3d9d48b4d4b538358f5ee76527ec702c1df0bca7c): complete subsection reference.

<a id="canonical-c20960cb727ac16626024f92b8d7cf537233dc7f90ade81aa3989f09fae7a907"></a>

## Next pages — use_custom_cluster_role_list / 831c64e92c43 / 4

- [use_custom_cluster_role_list.cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-539f180a36e99a61ed7fd0c3d9d48b4d4b538358f5ee76527ec702c1df0bca7c)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-539f180a36e99a61ed7fd0c3d9d48b4d4b538358f5ee76527ec702c1df0bca7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c2997f7b755459d3d1b88e1c1831868d6fbb5af5ecedb0752f038506fc65948"></a>

## use_custom_cluster_role_list.cluster_roles — use_custom_cluster_role_list.cluster_roles / 7f4f3745d61d / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-a4d967e0982944c5fedab2419e216b8fdc62c8b1427b2585fc696ffeb68b0ac8)
- use_custom_cluster_role_list.cluster_roles

<a id="canonical-f5e053f457ed7ee5bb9e252e802f545f2c83ee431800688f07556f9a12d5a156"></a>

Type: `"list"`. Computed.

List of active cluster role list for a K8s cluster.

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

<a id="canonical-06ed126cc6506b458ca74fa96596f4d689d003abb7514a70e500a45c7fb628a0"></a>

## Direct properties — use_custom_cluster_role_list.cluster_roles / 7f4f3745d61d / 3

<a id="canonical-fb762f006e1f81d517424e06e7c82b0ce35fe76c43f86e04cac8ae727b08230d"></a>

<a id="canonical-be0ebf0d721499e2375d9778c35d181fd37e7e92589691c6bd0ee939447729e0"></a>

## name property — use_custom_cluster_role_list.cluster_roles / 7f4f3745d61d / 4

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

<a id="canonical-5d2863cdbce829e2c4bcc2558684a12a5debbac273709d11c28965dd40cc23e0"></a>

<a id="canonical-00a13082c362630d1543d49f3d7511fff87be2ccd06bf4c9f69cd5b8f0fd5122"></a>

## namespace property — use_custom_cluster_role_list.cluster_roles / 7f4f3745d61d / 5

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

<a id="canonical-4bcda5e6b6aab53cb02d8ed89a799888157beb86d26a1038524b7f7dbd662c6e"></a>

<a id="canonical-cc9af9684af6849b892caffb0d530a59908e89e487c7cd1519e4f4fc14e7d04a"></a>

## tenant property — use_custom_cluster_role_list.cluster_roles / 7f4f3745d61d / 6

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

<a id="canonical-9a8da70e4d9268b8b043c16a230beb76afdde0a25000eb642139beddee256cee"></a>

## Next pages — use_custom_cluster_role_list.cluster_roles / 7f4f3745d61d / 7

- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-a4d967e0982944c5fedab2419e216b8fdc62c8b1427b2585fc696ffeb68b0ac8)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-c6b0becb60105bb78d7eddb05167bf53348b50913b51147940701b9f1292bc38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b481400b6be1488c4506e31f9dff3956796c9c9679be91d7f05a33037512ee7"></a>

## use_custom_pod_security_admission — use_custom_pod_security_admission / 8ca9d72111dc / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_custom_pod_security_admission

<a id="canonical-c7c3a5a682e4a6769b6ec100bede148438ac6e4118d434bf4abb879dca0434dc"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_pod\_security\_admission, use\_default\_pod\_security\_admission; Default:
use\_default\_pod\_security\_admission\] Type establishes a direct reference from one object(the
referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-c7c3a5a682e4a6769b6ec100bede148438ac6e4118d434bf4abb879dca0434dc)
- [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-8d78b4ca73138786ee3f8e084cf383a43840607f145bb5cd1a7b9a1c23f321e0)

Select alternatives according to the provider validators above.

<a id="canonical-531d9494b988ea78c47663d47558c3cd4d1683f746c09da9c9cfab46a53ec66d"></a>

## Direct properties — use_custom_pod_security_admission / 8ca9d72111dc / 3

<a id="canonical-fe7ba1892de7cb806ef5e708286d11402f1e5ed002c3c15f0cd9fe89136abcde"></a>

<a id="canonical-78159777aacc762799cd1ed55f55b02868aa2bd2c90f4d65a14a5341c4ea7521"></a>

## name property — use_custom_pod_security_admission / 8ca9d72111dc / 4

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

<a id="canonical-3059ffc0426a7360512af4bf71789b5f1b7ee103bb98bb1c414a42e86aca3a23"></a>

<a id="canonical-96a05d6d8eb78e197caa5128f920cce16da6ebc02df205a6a5db589e32505861"></a>

## namespace property — use_custom_pod_security_admission / 8ca9d72111dc / 5

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

<a id="canonical-225da710dc2b88080b8fccd0d642a7d8b05a63c4fbf29e7eb9de5d55e307daa0"></a>

<a id="canonical-2adf9513a5342c24420415e19827dc7b46f5d817bb3b7ace25e2db4909ebb79b"></a>

## tenant property — use_custom_pod_security_admission / 8ca9d72111dc / 6

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

<a id="canonical-1f236e9480fd4ea23bce78103b93a64d5b26eb9e56c24000dce1cb22b94f5488"></a>

## Next pages — use_custom_pod_security_admission / 8ca9d72111dc / 7

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-2a4d59d24a327815b184f45187c189a7b8a53806af7a4177c894065d674c943c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-914c8868bbb32922d1b7db345fb67fb501a9e54c7f899657b52067aaadf6b06b"></a>

## use_custom_psp_list — use_custom_psp_list / 291ad20a3a4a / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_custom_psp_list

<a id="canonical-f52e46283aa8b946a4c27fe2934931b51dbb39ad463b07c918871af632b51d12"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

Upstream description:

List of active Pod security policies for a K8s cluster.

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

- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-f52e46283aa8b946a4c27fe2934931b51dbb39ad463b07c918871af632b51d12)
- [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-8427e54e9d835c49e69c01358b362e06fa1162ef763d265f1cb0f346c9a9e438)

Select alternatives according to the provider validators above.

<a id="canonical-e5c1e05686113cd347340c94381988609549f5511e565bf84916b531d02117c7"></a>

## Direct properties — use_custom_psp_list / 291ad20a3a4a / 3

- [pod_security_policies](data-sources--k8s_cluster--reference--group-001.md#canonical-efa95b798949db83b404dd8fc46d512f69f116b4224a8712c453e93c0c19111b): complete subsection reference.

<a id="canonical-2dd83ffe544e99a9f287cd72d31c4e1c331f5c65a99360317ad8acf019cbb731"></a>

## Next pages — use_custom_psp_list / 291ad20a3a4a / 4

- [use_custom_psp_list.pod_security_policies](data-sources--k8s_cluster--reference--group-001.md#canonical-efa95b798949db83b404dd8fc46d512f69f116b4224a8712c453e93c0c19111b)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-efa95b798949db83b404dd8fc46d512f69f116b4224a8712c453e93c0c19111b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46719060193d9258f411f1da1f234949bc86797831acfaab86ae1f29e67560ea"></a>

## use_custom_psp_list.pod_security_policies — use_custom_psp_list.pod_security_policies / 4987c1db5aef / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2a4d59d24a327815b184f45187c189a7b8a53806af7a4177c894065d674c943c)
- use_custom_psp_list.pod_security_policies

<a id="canonical-0a015eb35fe74aa5761ef242a747d47663a1dab30e4b7214cf714b004eed874f"></a>

Type: `"list"`. Computed.

List of active Pod security policies for a K8s cluster.

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

<a id="canonical-e4200a36fd5e9056cfec4ffb9bb703dd7ad36896bc44b8a4ff7e9d54963acd2a"></a>

## Direct properties — use_custom_psp_list.pod_security_policies / 4987c1db5aef / 3

<a id="canonical-8eb77aac0ec459438d6c5624e12d04f21f903f9517ac619f88fd2430be2d8ab1"></a>

<a id="canonical-76a0dab6a7a93355ed3895f045569564a4745c269430cd74caab6fd93c81c83a"></a>

## name property — use_custom_psp_list.pod_security_policies / 4987c1db5aef / 4

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

<a id="canonical-1b5804f719212b36041ca78e4ff8e821bd9a00d465abe00b8393ee21df604b82"></a>

<a id="canonical-ecad99f4f194116526fa400f9a69a6ad90c9bea1d69cc6053bdec6d99af5151a"></a>

## namespace property — use_custom_psp_list.pod_security_policies / 4987c1db5aef / 5

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

<a id="canonical-4fefe001a3f79ea76a105d6982981e9aeaee19bfac039fa848bd9447f48d4a58"></a>

<a id="canonical-aec533c396eab9b9cab8415838b31261615decc57b68fac0a04ea9edc51143b3"></a>

## tenant property — use_custom_psp_list.pod_security_policies / 4987c1db5aef / 6

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

<a id="canonical-f078259b79e4a71daa3f6a430ec6329db4028e67a3681321c5365f2cc4971ff6"></a>

## Next pages — use_custom_psp_list.pod_security_policies / 4987c1db5aef / 7

- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2a4d59d24a327815b184f45187c189a7b8a53806af7a4177c894065d674c943c)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-cf9c596d759cdb7a3dc8914e0096acc67a42ba6eb26bf18fddf9eaac0c95061b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db1f6488678b942d35377b2602499597f8316082f859f9c951e30309023e3254"></a>

## use_default_cluster_role_bindings — use_default_cluster_role_bindings / 2a578973e9f1 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_default_cluster_role_bindings

<a id="canonical-af713c8b740253d576d41652e94614ccc5d964c104381b6ad30fb6388f06421f"></a>

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

<a id="canonical-f791bec1f8758b9eb6ed044482adb6547367beea48a4a3650a47a1b494ab78af"></a>

## Direct properties — use_default_cluster_role_bindings / 2a578973e9f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e5ecf01211013fc158b965b566952b512837d9bb41675ff2a0ef25972c4609b"></a>

## Next pages — use_default_cluster_role_bindings / 2a578973e9f1 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-4c6382ce195526069aa4491dcb3a2d1af2260bf99b382c29dc50770f663f570c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cadfc5273bd471ae0f9591a29eb45a3eda3673fceb4e78fa3b0ca540ff610b0"></a>

## use_default_cluster_roles — use_default_cluster_roles / d885e9df581c / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_default_cluster_roles

<a id="canonical-67c7fd0a40300de8623ff4f650719763ee4d558d3c06dd92862622d68e86ba06"></a>

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

<a id="canonical-874a060bec384a6db8a7eeb69ff52aef69448ec566136233972fd13d7da39824"></a>

## Direct properties — use_default_cluster_roles / d885e9df581c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20f99d7eb07a9a426ee65edfbf7ffb3fa68774ea82e94931c912376553472031"></a>

## Next pages — use_default_cluster_roles / d885e9df581c / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-5145aea8589f27ace674e7790f7e106945a8ce7dd91d049e9df2850ec1dfe2a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cba2397d7a9af3d8ff21e0295e9910f808dfd7d1d6f9e41c8693b0d863ea291"></a>

## use_default_pod_security_admission — use_default_pod_security_admission / 6f260cd0ffa2 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_default_pod_security_admission

<a id="canonical-8d78b4ca73138786ee3f8e084cf383a43840607f145bb5cd1a7b9a1c23f321e0"></a>

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

<a id="canonical-bda1709e9f994278cddf02360a0b1643fcbe9bc109d4f8fcb72e3b37e53cb3fe"></a>

## Direct properties — use_default_pod_security_admission / 6f260cd0ffa2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e3cbe9e4e7d892faa2701ed68cb0b89ae9f0457d4694b9e6d3e0b7de5bd7841"></a>

## Next pages — use_default_pod_security_admission / 6f260cd0ffa2 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-3db9f9e06f28dff62c1ec1d86fb879948c9ba58293936b1e462d04221b426cef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8702e3db791ecf84017925b0eee86c00ea3d9eff4f64960badeddede440d2a18"></a>

## use_default_psp — use_default_psp / 650e0d9c554e / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- use_default_psp

<a id="canonical-8427e54e9d835c49e69c01358b362e06fa1162ef763d265f1cb0f346c9a9e438"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-86f04218464d7a65ca20cf63ec3f1ee0ed5fff50b3c5fcacf6ddd3bff5c3af79"></a>

## Direct properties — use_default_psp / 650e0d9c554e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a57ca081b15e8a1fb3a48706e389fd7afb75cd621d27e421ef3ef666da2c0a8b"></a>

## Next pages — use_default_psp / 650e0d9c554e / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-02fc479f0a14b8d6ed6954b31d30542bef2b918cd2367aff7bc5628ecd804409"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85f8a66a1f53f257596f49e90ff173623c1d84bccb97e3526b540b04e97d5cee"></a>

## vk8s_namespace_access_deny — vk8s_namespace_access_deny / 404db95077d9 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- vk8s_namespace_access_deny

<a id="canonical-cc02f59fd5ed161634fa19227fe7afb0b65ec345db095353a18495bba11e1adc"></a>

Type: `["object", {}]`. Computed.

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

- [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-cc02f59fd5ed161634fa19227fe7afb0b65ec345db095353a18495bba11e1adc)
- [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-00ed27b0d62b8a7e7c0d969925d44ba79fab620f94b041895a4b461b57828ccf)

Select alternatives according to the provider validators above.

<a id="canonical-cb173f0298289bc25a77eef7018ac9e0ab008466494cfea55b847e3b6db04f10"></a>

## Direct properties — vk8s_namespace_access_deny / 404db95077d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62f764a0a4c70cf59f01e2525069c270ead0fde318959f1485baf549bf89570e"></a>

## Next pages — vk8s_namespace_access_deny / 404db95077d9 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-0c3710b30e4ec94fffdfae7def1580a088a9904fd8c57c8d076568509dc47078"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73854b1a027f5f7487a490d15b5938fc0dfdcb14a14afd54a83892d61bad60e4"></a>

## vk8s_namespace_access_permit — vk8s_namespace_access_permit / 8cd2735109fb / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- vk8s_namespace_access_permit

<a id="canonical-00ed27b0d62b8a7e7c0d969925d44ba79fab620f94b041895a4b461b57828ccf"></a>

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

<a id="canonical-d9ff97de32a77b6a7e897710390bc8aadab25a699d5d46d24ab63fb23ff1cb4d"></a>

## Direct properties — vk8s_namespace_access_permit / 8cd2735109fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d240ec724979afccd8b69416e73c049e08fe7a8966aab6e69e3a018e206a98f"></a>

## Next pages — vk8s_namespace_access_permit / 8cd2735109fb / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
