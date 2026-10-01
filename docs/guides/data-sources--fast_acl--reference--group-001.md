---
page_title: "xcsh_fast_acl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl reference."
---

# xcsh_fast_acl reference

<a id="canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-215298096878e71cf18f384b0f29a4dd9c1e109f5f099c9a4ee66a2d11220408"></a>

## Property reference — Property reference / dcef6dfe3243 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- Property reference

<a id="canonical-a7c8f993d7ae11ed42dd41bcb982f67c1e5f1af1a3d2775f18b65292c1f457e4"></a>

## Direct properties — Property reference / dcef6dfe3243 / 3

<a id="canonical-c6d9ee5959c7d23ce15de110992b73a6e06a9bfe363308bceb8d44097ddd8917"></a>

<a id="canonical-f85773a39b2985fa5925094189cf9dc4d5c0b494c0789f92bc4b664bbd58fe6b"></a>

## annotations property — Property reference / dcef6dfe3243 / 4

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

<a id="canonical-d20d345e95b5069def347c53eddf0c5bbb9f1449b754a8fc7c9e5fd97145c2c5"></a>

<a id="canonical-6281b3b2e80f5707473934d3cbd72b495f77fcb337b6d6e3fbe651828092718f"></a>

## description property — Property reference / dcef6dfe3243 / 5

Type: `"string"`. Computed.

Description of the FastACL.

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

<a id="canonical-cae4e9396f906ffe5cef05eeea60629eb32cdde52e34d4b373e6c15cb4f935ab"></a>

<a id="canonical-43eaca15c330afa5455f44daabbc6dbdba81dab9456de0d3b2dde3ebc4559d96"></a>

## id property — Property reference / dcef6dfe3243 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5cb0d53b729fbef3deca3a62ba3a84aa7c64d05472eb1f2c4b9efb0379f512c9"></a>

<a id="canonical-fe6fc9db199e46414dfa7dfaa10a023fbf271ba85f42061ee3d35ee96f7d4130"></a>

## labels property — Property reference / dcef6dfe3243 / 7

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

<a id="canonical-60babab15373a157d147b7b8bc65bf966c13dffeafd980e2aa06941777e5c881"></a>

<a id="canonical-aa659b74873c99d81c617c33f8b86e0aedb01fe7108fe0150c2a22b1da4aa65e"></a>

## name property — Property reference / dcef6dfe3243 / 8

Type: `"string"`. Required.

Name of the FastACL.

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

<a id="canonical-38c76ee5cc73863bb25678b3ab8cf09b73587d93a310365ff7e0d4c3f90e66da"></a>

<a id="canonical-daa97daa95003280789a834f50dba3f7bd900f8663c763d91fa6471fae03981c"></a>

## namespace property — Property reference / dcef6dfe3243 / 9

Type: `"string"`. Optional, Computed.

Namespace where the FastACL exists.

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

- [protocol_policer](data-sources--fast_acl--reference--group-001.md#canonical-35c4a4f70ed7b74b80742bcfe736bc3d1c67e2c806e1668a4df8716003db33e9): complete subsection reference.

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9): complete subsection reference.

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8): complete subsection reference.

<a id="canonical-65182bc1416c8c1c2f0ee28eb06ba9d5df89e87a1470f4e2bac0e1fc06d7b1b5"></a>

## All schema paths — Property reference / dcef6dfe3243 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--fast_acl--reference--group-001.md#canonical-c6d9ee5959c7d23ce15de110992b73a6e06a9bfe363308bceb8d44097ddd8917) |
| `description` | [description](data-sources--fast_acl--reference--group-001.md#canonical-d20d345e95b5069def347c53eddf0c5bbb9f1449b754a8fc7c9e5fd97145c2c5) |
| `id` | [id](data-sources--fast_acl--reference--group-001.md#canonical-cae4e9396f906ffe5cef05eeea60629eb32cdde52e34d4b373e6c15cb4f935ab) |
| `labels` | [labels](data-sources--fast_acl--reference--group-001.md#canonical-5cb0d53b729fbef3deca3a62ba3a84aa7c64d05472eb1f2c4b9efb0379f512c9) |
| `name` | [name](data-sources--fast_acl--reference--group-001.md#canonical-60babab15373a157d147b7b8bc65bf966c13dffeafd980e2aa06941777e5c881) |
| `namespace` | [namespace](data-sources--fast_acl--reference--group-001.md#canonical-38c76ee5cc73863bb25678b3ab8cf09b73587d93a310365ff7e0d4c3f90e66da) |
| `protocol_policer` | [protocol_policer](data-sources--fast_acl--reference--group-001.md#canonical-95128b0796cb5dfdf7a2adcab656324109a1e7035844f2ea061cf30ce201b5a7) |
| `protocol_policer.name` | [protocol_policer.name](data-sources--fast_acl--reference--group-001.md#canonical-866de3bec7c43398f28f7a838d0cc894642c5a182986716497bf0de9ad881ac1) |
| `protocol_policer.namespace` | [protocol_policer.namespace](data-sources--fast_acl--reference--group-001.md#canonical-d934ddee4564a246a76f49ea38d1fd6ee9e02898f2ff0776b054f8f21023a570) |
| `protocol_policer.tenant` | [protocol_policer.tenant](data-sources--fast_acl--reference--group-001.md#canonical-894a5b775c5c85fd374e3f0c9ffa80c1e9c8170e4250b2d03d5dcce2df54db5f) |
| `re_acl` | [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-158166d6616c8974be2ff0f31fd697c5748a8bb3ee53cdeb1bbe13f81bddaed9) |
| `re_acl.all_public_vips` | [re_acl.all_public_vips](data-sources--fast_acl--reference--group-001.md#canonical-9fc95c80667beb4f4068de39137ed85a91c77f7976e0e881e3a9e1459d128281) |
| `re_acl.default_tenant_vip` | [re_acl.default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-9d740f821a465dabf2ec240b82f784ac66e9b53e06450c7181add99f80b66c4b) |
| `re_acl.fast_acl_rules` | [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-10005c380a4332c893a52515c30e9070df3faefb819c2a3e45fd5536f62f346b) |
| `re_acl.fast_acl_rules.action` | [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-4c356425bcaa1d905412b3a994502646f166ab266a252e2280b81a18a3bd2ff2) |
| `re_acl.fast_acl_rules.action.policer_action` | [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-dc5e4e5d174ee4dffb07948d6a024700e745ab98a63dd301c32f117a826651bb) |
| `re_acl.fast_acl_rules.action.policer_action.ref` | [re_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-48aab4c00db83938f63cd9a7d51f7fc2692f62020f5efc255b5c5ec755e815be) |
| `re_acl.fast_acl_rules.action.policer_action.ref.kind` | [re_acl.fast_acl_rules.action.policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-d3e6442e7459d4a127eef5edbc07ef98852d79d5eb2a1c2333638e672b87a6ac) |
| `re_acl.fast_acl_rules.action.policer_action.ref.name` | [re_acl.fast_acl_rules.action.policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-bfe0c573b4dafa8c9848124805b39e5df249ddb4fdfb506ef9cf5dc9d9b5a14b) |
| `re_acl.fast_acl_rules.action.policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-a2747c87a820a33e637d37d1b528c06aa1a4cf7efe33202db7a1ab250f79d5f9) |
| `re_acl.fast_acl_rules.action.policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-9d1673098a5a44c29c02ca5d9d7617b2366b6575a8949778429ecde8c303c3b9) |
| `re_acl.fast_acl_rules.action.policer_action.ref.uid` | [re_acl.fast_acl_rules.action.policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-daa09e441c71869acbd7043d0eef0c321fc5014c6cac8799f2e3dccc90c9ce39) |
| `re_acl.fast_acl_rules.action.protocol_policer_action` | [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-f634f09067346da69a045807ae40a706ac3be3b9306361d908b986e9b209a670) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-841b1e857ca5e0d4669d3bd1be08c0efba5fa537a57e957c2bae8663f015dbc5) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-10f088dd0525af36e78dbfbe1ef2156fd019e82122223e534a9f55dd2fa12493) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-3102c9e2bbb1becb507d1e4245ae176974235127703dd0667a436c0d51bd55ee) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-a7ff6c547cc822baede74fd79eb0ba103cc6be7f558ff5301e9bdf0b07bc9601) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-7716378883921b8154f383c18464babbf2e4d6cd700840c065e9846f804f47a7) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-a79bfdfcf25e9da05ed4de46800a12d8cb3702f0768bd01bf8091c7105addde6) |
| `re_acl.fast_acl_rules.action.simple_action` | [re_acl.fast_acl_rules.action.simple_action](data-sources--fast_acl--reference--group-001.md#canonical-4813d8ce5a101926ce3021a4635ae62235f3d24c0285a69ec72cd2c0483a5373) |
| `re_acl.fast_acl_rules.ip_prefix_set` | [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-49bd7e327a4bc8c9516d296449cc65fe74cc5cb2de838956cf8ce32a92fb601b) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref` | [re_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-e38d00407d76323efdefc5e141d2094b00f7d254bb6cd3cf68db026face44cda) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [re_acl.fast_acl_rules.ip_prefix_set.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-7fbf961d602171f8fbe5ac9df1b945215196fd0571481395a0f1d91fe5f0765d) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.name` | [re_acl.fast_acl_rules.ip_prefix_set.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-0d6cfc24301c34a01ac7fca69666567f357a97b38f671869f730de275e881265) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [re_acl.fast_acl_rules.ip_prefix_set.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-7d677c88de93566cb768ccdd71ec5f78143e241000a20cea3e0652cde436f76a) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [re_acl.fast_acl_rules.ip_prefix_set.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-592be6a16cae1ba68fb7ab444c302eff87a965aa615e0751e461e3bd1829960b) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [re_acl.fast_acl_rules.ip_prefix_set.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-42cd6bdbe5fd9ae4ce02b28cd3876880820d0a3ef59cc486b3dbadf42b638a1a) |
| `re_acl.fast_acl_rules.metadata` | [re_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-772d5a458eeb0998b9d54a842b9bb41fd7409256afabced41f9e5d6001fdf820) |
| `re_acl.fast_acl_rules.metadata.description_spec` | [re_acl.fast_acl_rules.metadata.description_spec](data-sources--fast_acl--reference--group-001.md#canonical-280f8f87690c60986d279243a948cc506c67ca3c7a1bc87f1c15f2d47ba099e4) |
| `re_acl.fast_acl_rules.metadata.name` | [re_acl.fast_acl_rules.metadata.name](data-sources--fast_acl--reference--group-001.md#canonical-4b790954b87017dfba115b1eed6498ba9d55e59a23651e14b746fe1de9d1a7aa) |
| `re_acl.fast_acl_rules.port` | [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-08053f518a38b35be8d4c02af76fcd6ce1d9bc411d86ea5a65e64ae2e20136a3) |
| `re_acl.fast_acl_rules.port.all` | [re_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-fc967209e6b3cec63f74a63f86c2aeb77fa3124acd0de772b27464925d2298d1) |
| `re_acl.fast_acl_rules.port.dns` | [re_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-61257e35b07cca4a73d569c64dd621f2920830b87dd9b23e10ac80b3fa9ce98b) |
| `re_acl.fast_acl_rules.port.user_defined` | [re_acl.fast_acl_rules.port.user_defined](data-sources--fast_acl--reference--group-001.md#canonical-75cbf4dea52bc3640fc239b1cd76a029eaa1be7f84043fb82db7513e99b73b8f) |
| `re_acl.fast_acl_rules.prefix` | [re_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-246d632009911bbd2fbb5a0d4ed91a38123de20e6e43ce3f5acb3855e95205a3) |
| `re_acl.fast_acl_rules.prefix.prefix` | [re_acl.fast_acl_rules.prefix.prefix](data-sources--fast_acl--reference--group-001.md#canonical-95330d744b89b126c87c96ec9e0441510a18c16eb681f26a5388e59e3c73ad61) |
| `re_acl.selected_tenant_vip` | [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-c7c406df3493ef390fdcc9e7d12da22a11c451eff301a9a0241565f3001219f0) |
| `re_acl.selected_tenant_vip.default_tenant_vip` | [re_acl.selected_tenant_vip.default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-d9238949c8747a2730f75d1cada4c24413a31d302a38431c54ce7be73f7f94c4) |
| `re_acl.selected_tenant_vip.public_ip_refs` | [re_acl.selected_tenant_vip.public_ip_refs](data-sources--fast_acl--reference--group-001.md#canonical-8cada62ffabc2421b04cc3a18abe1b5e1bf2c6cb5781a19d29be737e49515032) |
| `re_acl.selected_tenant_vip.public_ip_refs.name` | [re_acl.selected_tenant_vip.public_ip_refs.name](data-sources--fast_acl--reference--group-001.md#canonical-db30b51683b2a0cd1b28b0334ff35d3a5fa477787cc3310b11f16aa7f9058ee0) |
| `re_acl.selected_tenant_vip.public_ip_refs.namespace` | [re_acl.selected_tenant_vip.public_ip_refs.namespace](data-sources--fast_acl--reference--group-001.md#canonical-e81898139a4a129ce82da6b24bc39a109f0fd2b7140ffab1763c9bb988ae9889) |
| `re_acl.selected_tenant_vip.public_ip_refs.tenant` | [re_acl.selected_tenant_vip.public_ip_refs.tenant](data-sources--fast_acl--reference--group-001.md#canonical-9e05fed77ebfdd06be1eadc3dd4e6b0300b884dbb89f295dd433c452a2570b16) |
| `site_acl` | [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-3d945c785f0574b5017b265c0815132febbea769c7750e97a943c3532927aad4) |
| `site_acl.all_services` | [site_acl.all_services](data-sources--fast_acl--reference--group-001.md#canonical-9cffa995690cd60baf94affeff9f4d59b94274dd0293532c1681dec4e29d7536) |
| `site_acl.fast_acl_rules` | [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-d671947fe892c3b54bd68420bdab29552b81ba79eb1caad84331c843f1eff747) |
| `site_acl.fast_acl_rules.action` | [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-ea0176e8cc84c049a33614f81e63e2b79dce09109aa154325dabb61b64add879) |
| `site_acl.fast_acl_rules.action.policer_action` | [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-a2bb86ad1e52b19e8a89296504fa5b4b29a38170590b24f5a33432214066ce70) |
| `site_acl.fast_acl_rules.action.policer_action.ref` | [site_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-3a7cc7842dc16301408689e69fe71f71024d38f10a626c3acec8191c342ae8a0) |
| `site_acl.fast_acl_rules.action.policer_action.ref.kind` | [site_acl.fast_acl_rules.action.policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-8a63145e3b3c65e377a63fe6df4dd348b22c55780fa20fdab4b8c164b4440a92) |
| `site_acl.fast_acl_rules.action.policer_action.ref.name` | [site_acl.fast_acl_rules.action.policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-b3dfd4926254b64068de5a19c112a3c848a4562fe10ea07ead6f2e223990b66f) |
| `site_acl.fast_acl_rules.action.policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-592df430c8ef832495dbb1e959ceda03b951c479c6d9b742f8d9a4835a9d9b69) |
| `site_acl.fast_acl_rules.action.policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-55551eaba6c45d2e0b94549a4fe5866b2cb583c70087ba7e042b823902e30812) |
| `site_acl.fast_acl_rules.action.policer_action.ref.uid` | [site_acl.fast_acl_rules.action.policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-4dc9802c0f648a0a45d4fa70ae25b7361875a1496997af7dcdae2e6775d70b87) |
| `site_acl.fast_acl_rules.action.protocol_policer_action` | [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-bf5239aad5e607d2eb569509afa69cde76e6d1990f44630ca44a7b24ece67fbd) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-1846bdd5bc0ed79e195d1ebd4af21d8a44edf5995bc90234430cd3cb8fdd1c45) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-16bf8d130e2c56fbfc4a6dd1e53b53eb4e81b3b9ff14c3cfff5d0e93cde36699) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-8a4ee362c3bed63861cec36a6057e09199e8b626b45e69e369a075a30bc3e9c4) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-49c1719a4167f1a51f98a189e53d320d721b50b6ca09394d09ee59c371e5a1e8) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-f477c22d26039c715015f9d9b892ba4a2c5d88d933d5f3252d745c0d94451687) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-1426e3df7ec54429cea76e89bb76c8d0f24fb867c0888a8f1c6cb0b1e5a5191c) |
| `site_acl.fast_acl_rules.action.simple_action` | [site_acl.fast_acl_rules.action.simple_action](data-sources--fast_acl--reference--group-001.md#canonical-c796799e68b92ea2eb9535ae891c68dff2af27e1a82ffaa4d4628d923f588542) |
| `site_acl.fast_acl_rules.ip_prefix_set` | [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-b50f3aa986f8f40f78231f460162399b6b0df9ca7b198435c6f840ad57bf68a0) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref` | [site_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-cd7a46ec3d3c49766f74d3aa6d64cdf950d1d0d8aac75540bac91940f515c28d) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [site_acl.fast_acl_rules.ip_prefix_set.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-8a9e1c15cd3f1ccd18016dc1ecb7bf97782dc207807249b8f07e8c76f813e3f9) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.name` | [site_acl.fast_acl_rules.ip_prefix_set.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-e91bf233813c4bb8fce3b833e0789aaa4a820dd731a2e5d9a39c5f1bd64bda2c) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [site_acl.fast_acl_rules.ip_prefix_set.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-24c91d7ed2747721c7c4b29a5908ab4fe76e618bd0be9a1b5631fd6df5563487) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [site_acl.fast_acl_rules.ip_prefix_set.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-fc59fdc47e0233f2994203537d78ff6e972e207aca1876b0e418b22b81bb748b) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [site_acl.fast_acl_rules.ip_prefix_set.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-c3e117c785e2ce616524cb01cef6d3e47161662c34b0c1956534595242e137ce) |
| `site_acl.fast_acl_rules.metadata` | [site_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-87d95569b2bdbb2812d51b60c03d4f568ca8160a99ac7bb271d6065c5e6a2ba3) |
| `site_acl.fast_acl_rules.metadata.description_spec` | [site_acl.fast_acl_rules.metadata.description_spec](data-sources--fast_acl--reference--group-001.md#canonical-c9c30456b16c669a4181e41b9e40bcd1e64fdd104e7ba319d1029d213eb03f1a) |
| `site_acl.fast_acl_rules.metadata.name` | [site_acl.fast_acl_rules.metadata.name](data-sources--fast_acl--reference--group-001.md#canonical-cd850cfc773f778f691aa05eeba38001f679d78e52d4f082d5dd91e5c82c3102) |
| `site_acl.fast_acl_rules.port` | [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-c4bdf8c888cebd718b65fd2127821a75d12193e3beadfad5431c2116b6493318) |
| `site_acl.fast_acl_rules.port.all` | [site_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-656c0c8b74c346feb3a8bc31d56815361a291bec86ce349597a78649d726fbda) |
| `site_acl.fast_acl_rules.port.dns` | [site_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-ceb65f5d04ea3d59542591e739f5ca2a3a0c391fa7007abe4edb0481c248c67c) |
| `site_acl.fast_acl_rules.port.user_defined` | [site_acl.fast_acl_rules.port.user_defined](data-sources--fast_acl--reference--group-001.md#canonical-d59efec33ec8733fb11113b714d8c744fc2fd934241054842affe814ad6e21d4) |
| `site_acl.fast_acl_rules.prefix` | [site_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-2bc3b076936a0e5aca834c1c377a27bf8b5ec398a93c12b1d00b83df96075ac3) |
| `site_acl.fast_acl_rules.prefix.prefix` | [site_acl.fast_acl_rules.prefix.prefix](data-sources--fast_acl--reference--group-001.md#canonical-758740625ef8a197992dbf3a966b459a2e0311d4aa76a924ebe031255a923c59) |
| `site_acl.inside_network` | [site_acl.inside_network](data-sources--fast_acl--reference--group-001.md#canonical-0da8f8ff10f27b1574103990527778965761ffbc2826df0c369d857fcbf9a00e) |
| `site_acl.interface_services` | [site_acl.interface_services](data-sources--fast_acl--reference--group-001.md#canonical-49774e3b517efaa7b41ee3c8f269246d45dd7ff340e02471fbe6cf06674dfebe) |
| `site_acl.outside_network` | [site_acl.outside_network](data-sources--fast_acl--reference--group-001.md#canonical-bbd9fef5f78a1b4da05b925d49f9f46c377cd71aaa3dafdc99fadc39b1faa133) |
| `site_acl.vip_services` | [site_acl.vip_services](data-sources--fast_acl--reference--group-001.md#canonical-ef72401f8fa1c369ffab2989460cf378dcb466322f410fc106835e3424d4c072) |

<a id="canonical-3d6e61876e4b0e6c1ce612ce2638780efed9bf95a42772563bf17c21fe95b7cb"></a>

## Next pages — Property reference / dcef6dfe3243 / 11

- [protocol_policer](data-sources--fast_acl--reference--group-001.md#canonical-35c4a4f70ed7b74b80742bcfe736bc3d1c67e2c806e1668a4df8716003db33e9)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-35c4a4f70ed7b74b80742bcfe736bc3d1c67e2c806e1668a4df8716003db33e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbe23b75351c879f0378059577bddf5e26ff74196b57eef21e38ee19d0402fed"></a>

## protocol_policer — protocol_policer / 69ddee5cb139 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- protocol_policer

<a id="canonical-95128b0796cb5dfdf7a2adcab656324109a1e7035844f2ea061cf30ce201b5a7"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

<a id="canonical-6bb0bcf1bdfa7947ea8e940727aad25012a8e7e83723f28066a95f9ec63f9471"></a>

## Direct properties — protocol_policer / 69ddee5cb139 / 3

<a id="canonical-866de3bec7c43398f28f7a838d0cc894642c5a182986716497bf0de9ad881ac1"></a>

<a id="canonical-b70a65a0c29a2e0d3fe836ba64d1c172e1f233d77192cbfd724c0809e3422fc2"></a>

## name property — protocol_policer / 69ddee5cb139 / 4

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

<a id="canonical-d934ddee4564a246a76f49ea38d1fd6ee9e02898f2ff0776b054f8f21023a570"></a>

<a id="canonical-bdd708da0283bb0f00a36650ae19d91479051560f2607ecab9023f5ee94b0e1c"></a>

## namespace property — protocol_policer / 69ddee5cb139 / 5

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

<a id="canonical-894a5b775c5c85fd374e3f0c9ffa80c1e9c8170e4250b2d03d5dcce2df54db5f"></a>

<a id="canonical-0842505bb081ce32957653b41bcc5b8632ea3c1869dd8fbe92661aa58ceea360"></a>

## tenant property — protocol_policer / 69ddee5cb139 / 6

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

<a id="canonical-377bbcaac02b808b853ae992842c1d1d2f58ac68fedf847b294b288b9edc3e64"></a>

## Next pages — protocol_policer / 69ddee5cb139 / 7

- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d51a6ddad3f3e81d7027873d51c6a33707e085ecef88966d5c38e9ce6719aabe"></a>

## re_acl — re_acl / e1c6b9c9c441 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- re_acl

<a id="canonical-158166d6616c8974be2ff0f31fd697c5748a8bb3ee53cdeb1bbe13f81bddaed9"></a>

Type: `"single"`. Computed.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Upstream description:

Fast ACL definition for RE.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vip_choice": "[\"all_public_vips\",\"default_tenant_vip\",\"selected_tenant_vip\"]"
}
```

OneOf alternatives in this subsection:

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-158166d6616c8974be2ff0f31fd697c5748a8bb3ee53cdeb1bbe13f81bddaed9)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-3d945c785f0574b5017b265c0815132febbea769c7750e97a943c3532927aad4)

Select alternatives according to the provider validators above.

<a id="canonical-63b47d6f2c996abb1020191e966b65edf19b59f86913943dfffc1c464d61ee6e"></a>

## Direct properties — re_acl / e1c6b9c9c441 / 3

- [all_public_vips](data-sources--fast_acl--reference--group-001.md#canonical-a316194a9d57d2da92fbf66d2aec23a10c73d4acbbc664b7c39809e06f4c2d6a): complete subsection reference.

- [default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-8836797a65b0672bfde40e5a92267ad6677ed27498fd1acc4a0830bed293977d): complete subsection reference.

- [fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e): complete subsection reference.

- [selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-64435eb04204e8a306f4cc99cac48f897a85682af06bcb77cc63bbf2c1ea9be1): complete subsection reference.

<a id="canonical-49201ad5853717191c000d9c558b1681b27ede61da7837db0a49048dfbfa56ab"></a>

## Next pages — re_acl / e1c6b9c9c441 / 4

- [re_acl.all_public_vips](data-sources--fast_acl--reference--group-001.md#canonical-a316194a9d57d2da92fbf66d2aec23a10c73d4acbbc664b7c39809e06f4c2d6a)
- [re_acl.default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-8836797a65b0672bfde40e5a92267ad6677ed27498fd1acc4a0830bed293977d)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-64435eb04204e8a306f4cc99cac48f897a85682af06bcb77cc63bbf2c1ea9be1)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-a316194a9d57d2da92fbf66d2aec23a10c73d4acbbc664b7c39809e06f4c2d6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a124ba026b36315e11abef00f81e674ee80658726d0b1ff1f5984fa2cf9b9bb9"></a>

## re_acl.all_public_vips — re_acl.all_public_vips / dd4368160ee5 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- re_acl.all_public_vips

<a id="canonical-9fc95c80667beb4f4068de39137ed85a91c77f7976e0e881e3a9e1459d128281"></a>

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

<a id="canonical-cb6bdd04cbc5c3e5477c3292eb309a25c19c0336bcc2a9711de70f0aefe13244"></a>

## Direct properties — re_acl.all_public_vips / dd4368160ee5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36b6cfbbbb1afc274848a6dcccc7e16f4dceed3b99e5cef397f8bb4265404490"></a>

## Next pages — re_acl.all_public_vips / dd4368160ee5 / 4

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-8836797a65b0672bfde40e5a92267ad6677ed27498fd1acc4a0830bed293977d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6dd5342f97297e8614f0125140e807fcc5b08fcde3876974d0650ab9278b3944"></a>

## re_acl.default_tenant_vip — re_acl.default_tenant_vip / 05a3ce7f5b2e / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- re_acl.default_tenant_vip

<a id="canonical-9d740f821a465dabf2ec240b82f784ac66e9b53e06450c7181add99f80b66c4b"></a>

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

<a id="canonical-e530774af12f353c853466c3e1da5d776a7ce27dda897fbf160accdf3b608ef2"></a>

## Direct properties — re_acl.default_tenant_vip / 05a3ce7f5b2e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3ee6709bbe480fa1a4136a0fb54edb1c2d43cf4de75f5a388f92d04b88fc622"></a>

## Next pages — re_acl.default_tenant_vip / 05a3ce7f5b2e / 4

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3955133daac1d164422656f77d54f8435dd9f3f01397cad2d023459863b3ff85"></a>

## re_acl.fast_acl_rules — re_acl.fast_acl_rules / 296b616cc7fd / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- re_acl.fast_acl_rules

<a id="canonical-10005c380a4332c893a52515c30e9070df3faefb819c2a3e45fd5536f62f346b"></a>

Type: `"list"`. Computed.

Rules. Fast ACL rules to match. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Fast ACL rules to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-c2fee7f1a6a7f20b8e97044c1b582234f2ed7b9e93da670e204f4a12695b5be6"></a>

## Direct properties — re_acl.fast_acl_rules / 296b616cc7fd / 3

- [action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743): complete subsection reference.

- [ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-906c4e33b264ba1fe67afb3b9c6bc43d8ed7b0c84153af7244596e3cc4fa8046): complete subsection reference.

- [metadata](data-sources--fast_acl--reference--group-001.md#canonical-b3405e5155160265fa163b7fa15b02899fd0f76712910895f83471752c314af6): complete subsection reference.

- [port](data-sources--fast_acl--reference--group-001.md#canonical-780c68f973c008f6d087e74bec354e063cef81ed3ededb002a9a19a1403bb0ea): complete subsection reference.

- [prefix](data-sources--fast_acl--reference--group-001.md#canonical-bf6532255eb18c04e83e26c9e85d85b631b1a26ae8dcdc3cf0cf7880dcd02078): complete subsection reference.

<a id="canonical-faf41c89d73a79e188c691d1f8fa0ad4d7a74799a77dc3e6c2b81323846bbc1b"></a>

## Next pages — re_acl.fast_acl_rules / 296b616cc7fd / 4

- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743)
- [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-906c4e33b264ba1fe67afb3b9c6bc43d8ed7b0c84153af7244596e3cc4fa8046)
- [re_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-b3405e5155160265fa163b7fa15b02899fd0f76712910895f83471752c314af6)
- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-780c68f973c008f6d087e74bec354e063cef81ed3ededb002a9a19a1403bb0ea)
- [re_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-bf6532255eb18c04e83e26c9e85d85b631b1a26ae8dcdc3cf0cf7880dcd02078)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5859fdec3d404562fe8a4441e5623b8387762a1464f6d69141da5e56a3a096d4"></a>

## re_acl.fast_acl_rules.action — re_acl.fast_acl_rules.action / 70a77faa7c02 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- re_acl.fast_acl_rules.action

<a id="canonical-4c356425bcaa1d905412b3a994502646f166ab266a252e2280b81a18a3bd2ff2"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

<a id="canonical-0336c055240416d105bba65bda930d4bd1b760e7a8e67ab911b682736e76580a"></a>

## Direct properties — re_acl.fast_acl_rules.action / 70a77faa7c02 / 3

- [policer_action](data-sources--fast_acl--reference--group-001.md#canonical-e0c9536eae58a8b5c4b8f14b156768438f5bd0cb01ee350b1b55c557910de2cb): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-cacff6ff9ca93f00494af31cb5f1036f0943adce93b78001e7471a5014e25e7a): complete subsection reference.

<a id="canonical-4813d8ce5a101926ce3021a4635ae62235f3d24c0285a69ec72cd2c0483a5373"></a>

<a id="canonical-c1d177579aa00c836d1866027e2a8aad4a9bb7448da0034a1b5c88f967801c15"></a>

## simple_action property — re_acl.fast_acl_rules.action / 70a77faa7c02 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b3d6cea47bc92e54c06ca226fb62cf2ca0093cf16516040d37cf8d1f883a38d6"></a>

## Next pages — re_acl.fast_acl_rules.action / 70a77faa7c02 / 5

- [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-e0c9536eae58a8b5c4b8f14b156768438f5bd0cb01ee350b1b55c557910de2cb)
- [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-cacff6ff9ca93f00494af31cb5f1036f0943adce93b78001e7471a5014e25e7a)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-e0c9536eae58a8b5c4b8f14b156768438f5bd0cb01ee350b1b55c557910de2cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86e73a70a1fbbff86b2f5168e3debf3ffd830034923aa337059563d8ba48ea1d"></a>

## re_acl.fast_acl_rules.action.policer_action — re_acl.fast_acl_rules.action.policer_action / 31a85be7642b / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743)
- re_acl.fast_acl_rules.action.policer_action

<a id="canonical-dc5e4e5d174ee4dffb07948d6a024700e745ab98a63dd301c32f117a826651bb"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-463097fb3f2112bbaf94cee3b661853d35c499527cda5c724017216380dfa46a"></a>

## Direct properties — re_acl.fast_acl_rules.action.policer_action / 31a85be7642b / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-7c6d9876ef1dbe4b02ab92c2208acfc7554fa523ed58df31967d688f7cae5fe4): complete subsection reference.

<a id="canonical-0f027036d85241e81f7a9955044fc6c60b6f59b3f7df11cc1da17662e01875fc"></a>

## Next pages — re_acl.fast_acl_rules.action.policer_action / 31a85be7642b / 4

- [re_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-7c6d9876ef1dbe4b02ab92c2208acfc7554fa523ed58df31967d688f7cae5fe4)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-7c6d9876ef1dbe4b02ab92c2208acfc7554fa523ed58df31967d688f7cae5fe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd25f507a9657b0720963e3d9b672d03e5b6bb2417dd17edf25e7ed3e9ed3be1"></a>

## re_acl.fast_acl_rules.action.policer_action.ref — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743)
- [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-e0c9536eae58a8b5c4b8f14b156768438f5bd0cb01ee350b1b55c557910de2cb)
- re_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-48aab4c00db83938f63cd9a7d51f7fc2692f62020f5efc255b5c5ec755e815be"></a>

Type: `"list"`. Computed.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

<a id="canonical-f21df4fc100f5859acebfb9a5929a6b7f535aa39ee93e316b0178749dd1fa6e3"></a>

## Direct properties — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 3

<a id="canonical-d3e6442e7459d4a127eef5edbc07ef98852d79d5eb2a1c2333638e672b87a6ac"></a>

<a id="canonical-a019e5a7c2077d85f6912bf17182839e41c65cf45c61abd8a0e56e7f573a3dc8"></a>

## kind property — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 4

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

<a id="canonical-bfe0c573b4dafa8c9848124805b39e5df249ddb4fdfb506ef9cf5dc9d9b5a14b"></a>

<a id="canonical-735877f10c636ce2296df59efbfe417ed52b1f99dc06accdba8cb72c5c8bc7ef"></a>

## name property — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 5

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

<a id="canonical-a2747c87a820a33e637d37d1b528c06aa1a4cf7efe33202db7a1ab250f79d5f9"></a>

<a id="canonical-a63676ce17bde4d9f20226b9bfba7f6a6f9e972308cce0002dc74e237bccc01c"></a>

## namespace property — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 6

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

<a id="canonical-9d1673098a5a44c29c02ca5d9d7617b2366b6575a8949778429ecde8c303c3b9"></a>

<a id="canonical-5f0b14f08ed633191fa6b5a9861a65a3312ebbaef9e788a0dc699702ee599f71"></a>

## tenant property — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 7

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

<a id="canonical-daa09e441c71869acbd7043d0eef0c321fc5014c6cac8799f2e3dccc90c9ce39"></a>

<a id="canonical-a703b3bb93cac5004abc08f7f539e4eef04e3c13c3a7d4fd3d8868afc961688d"></a>

## uid property — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 8

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

<a id="canonical-c644def0328667bdcb3f724e3f9033c2c23e220c656a8538a2438d2c1a308007"></a>

## Next pages — re_acl.fast_acl_rules.action.policer_action.ref / 68ab557aec09 / 9

- [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-e0c9536eae58a8b5c4b8f14b156768438f5bd0cb01ee350b1b55c557910de2cb)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-cacff6ff9ca93f00494af31cb5f1036f0943adce93b78001e7471a5014e25e7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05231c2d0470770521ef9e6400b6f631f3acc9cfae444a9045626d59787555ca"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action — re_acl.fast_acl_rules.action.protocol_policer_action / 3c28306e3d01 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743)
- re_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-f634f09067346da69a045807ae40a706ac3be3b9306361d908b986e9b209a670"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-7da9c8dda4d946dd5bcbb29548f6467ca1b65c4206eaae995df82178eb4e83ab"></a>

## Direct properties — re_acl.fast_acl_rules.action.protocol_policer_action / 3c28306e3d01 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-ccae775306b69dffd8608da45f2e0f58dd19b076a68968ff6d62017f65ff814f): complete subsection reference.

<a id="canonical-10fea902ca126a70bbf327cb352f77effb2f55e778f51026dee1d9a8916e797d"></a>

## Next pages — re_acl.fast_acl_rules.action.protocol_policer_action / 3c28306e3d01 / 4

- [re_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-ccae775306b69dffd8608da45f2e0f58dd19b076a68968ff6d62017f65ff814f)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-ccae775306b69dffd8608da45f2e0f58dd19b076a68968ff6d62017f65ff814f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4519052cf974743279f418d13c8bade5eac937960b0b4fed2dc423ce34dc87d"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action.ref — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-a472909fe8c2d003bc89c132726b70e190de505d883f593bfe7fe301bc050743)
- [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-cacff6ff9ca93f00494af31cb5f1036f0943adce93b78001e7471a5014e25e7a)
- re_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-841b1e857ca5e0d4669d3bd1be08c0efba5fa537a57e957c2bae8663f015dbc5"></a>

Type: `"list"`. Computed.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

<a id="canonical-a6561063f8ad9f11c4b4d01cb4c1342c289e1b0e95dc4941df0e9002314e79b4"></a>

## Direct properties — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 3

<a id="canonical-10f088dd0525af36e78dbfbe1ef2156fd019e82122223e534a9f55dd2fa12493"></a>

<a id="canonical-5ab898d824aca75e6a74218d0e5c5212b62090f93864e2f3d728996cc61bc766"></a>

## kind property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 4

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

<a id="canonical-3102c9e2bbb1becb507d1e4245ae176974235127703dd0667a436c0d51bd55ee"></a>

<a id="canonical-2b41fe88c4aa9ab1315d10d215be88a93bfaa1e7f0bdb14d3c61cabbfe10eecf"></a>

## name property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 5

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

<a id="canonical-a7ff6c547cc822baede74fd79eb0ba103cc6be7f558ff5301e9bdf0b07bc9601"></a>

<a id="canonical-4e80743fd869e91cd89b70e24bd2e8b14739a142e86bb8a6c91013b461f6366a"></a>

## namespace property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 6

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

<a id="canonical-7716378883921b8154f383c18464babbf2e4d6cd700840c065e9846f804f47a7"></a>

<a id="canonical-3823caa8ead6c3c79347356e571a700d8517e08594ff269ae75d9a0f710b19c5"></a>

## tenant property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 7

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

<a id="canonical-a79bfdfcf25e9da05ed4de46800a12d8cb3702f0768bd01bf8091c7105addde6"></a>

<a id="canonical-814418f0b363543b1d50c6f6a3aef6e86cfd1401a22d9857320c26dd910669b6"></a>

## uid property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 8

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

<a id="canonical-9875b7cdab90bf7a1a2d84defdb67e920c88f7582172f60bdc1c300bfaee45d1"></a>

## Next pages — re_acl.fast_acl_rules.action.protocol_policer_action.ref / 995e6d89e8b7 / 9

- [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-cacff6ff9ca93f00494af31cb5f1036f0943adce93b78001e7471a5014e25e7a)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-906c4e33b264ba1fe67afb3b9c6bc43d8ed7b0c84153af7244596e3cc4fa8046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd415a04218ba2179780cc7495fead434e0c0cb19cc4ed4ab7c51049613ffec1"></a>

## re_acl.fast_acl_rules.ip_prefix_set — re_acl.fast_acl_rules.ip_prefix_set / 6c19e4a5bde2 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- re_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-49bd7e327a4bc8c9516d296449cc65fe74cc5cb2de838956cf8ce32a92fb601b"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-17d2f6e3a566edc9da29785a7251e47ae6550c3eb393af6485aee608a2b14f32"></a>

## Direct properties — re_acl.fast_acl_rules.ip_prefix_set / 6c19e4a5bde2 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-40cb943af78c3a39aa45a4430e6020bd578da816ae827e3e67eabae2a3d4af94): complete subsection reference.

<a id="canonical-2fdc103de3c9c551bd9cc1db218b3715d71d5559a7255cda8a6d8578ee696e56"></a>

## Next pages — re_acl.fast_acl_rules.ip_prefix_set / 6c19e4a5bde2 / 4

- [re_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-40cb943af78c3a39aa45a4430e6020bd578da816ae827e3e67eabae2a3d4af94)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-40cb943af78c3a39aa45a4430e6020bd578da816ae827e3e67eabae2a3d4af94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69ea87be93874873acb19776512c0adf4454de764c236e09dbec32a58360ef55"></a>

## re_acl.fast_acl_rules.ip_prefix_set.ref — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-906c4e33b264ba1fe67afb3b9c6bc43d8ed7b0c84153af7244596e3cc4fa8046)
- re_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-e38d00407d76323efdefc5e141d2094b00f7d254bb6cd3cf68db026face44cda"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-d6b3553901a48d688217e00c0546f45f95c8e0d588de34ee25c083c2fd941b08"></a>

## Direct properties — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 3

<a id="canonical-7fbf961d602171f8fbe5ac9df1b945215196fd0571481395a0f1d91fe5f0765d"></a>

<a id="canonical-9964b8818abfa012c9b8ae4af57172ed627190d4a22c0e1cc104355f3a93c9a6"></a>

## kind property — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 4

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

<a id="canonical-0d6cfc24301c34a01ac7fca69666567f357a97b38f671869f730de275e881265"></a>

<a id="canonical-4f8df3ee14f9b9850e0c35943c59f9b79276d7f2906d4313cbf26a4b875acdfd"></a>

## name property — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 5

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

<a id="canonical-7d677c88de93566cb768ccdd71ec5f78143e241000a20cea3e0652cde436f76a"></a>

<a id="canonical-660d0e131e40d1dc76c55266dad1bf7b828cb53361b2743eca2dedda41da3911"></a>

## namespace property — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 6

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

<a id="canonical-592be6a16cae1ba68fb7ab444c302eff87a965aa615e0751e461e3bd1829960b"></a>

<a id="canonical-3a0eb97063b92eaa50a1963a30e1f4cc8fa18456aa61b0ce23baaefd3e15a508"></a>

## tenant property — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 7

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

<a id="canonical-42cd6bdbe5fd9ae4ce02b28cd3876880820d0a3ef59cc486b3dbadf42b638a1a"></a>

<a id="canonical-0dd420796f440a91c9539e669c4c71406f568d7fb8c03ffa50e16e566648426e"></a>

## uid property — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 8

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

<a id="canonical-8df5b6f7f33465e9d52ba52eca86b7213ddd06df8ef98082a985c8918a3d3e8e"></a>

## Next pages — re_acl.fast_acl_rules.ip_prefix_set.ref / daec8ccc0f51 / 9

- [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-906c4e33b264ba1fe67afb3b9c6bc43d8ed7b0c84153af7244596e3cc4fa8046)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-b3405e5155160265fa163b7fa15b02899fd0f76712910895f83471752c314af6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07bcb84063bae1e29e5616829482eb067e65e98224e3d8982dd9753ee38d8dd2"></a>

## re_acl.fast_acl_rules.metadata — re_acl.fast_acl_rules.metadata / 15c56a6e6016 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- re_acl.fast_acl_rules.metadata

<a id="canonical-772d5a458eeb0998b9d54a842b9bb41fd7409256afabced41f9e5d6001fdf820"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-45a399e4b025d9983bd3510c275b0b54057943ee62c50f36a190c8e7a2c93b7d"></a>

## Direct properties — re_acl.fast_acl_rules.metadata / 15c56a6e6016 / 3

<a id="canonical-280f8f87690c60986d279243a948cc506c67ca3c7a1bc87f1c15f2d47ba099e4"></a>

<a id="canonical-2c4a0cd4dfd45bbcb5e7c941a2a4d2a67c76cb8ff82a56c7df29c6597f3b9381"></a>

## description_spec property — re_acl.fast_acl_rules.metadata / 15c56a6e6016 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-4b790954b87017dfba115b1eed6498ba9d55e59a23651e14b746fe1de9d1a7aa"></a>

<a id="canonical-a5e4cadc95d3dfa240e27109d2e540acf2fd085d2d9a9a9bbe7e5cb636a5f01d"></a>

## name property — re_acl.fast_acl_rules.metadata / 15c56a6e6016 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-a62b07546aa99d3756403802d794ad78118ea970f3760c48e17b1d29f314ac2d"></a>

## Next pages — re_acl.fast_acl_rules.metadata / 15c56a6e6016 / 6

- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-780c68f973c008f6d087e74bec354e063cef81ed3ededb002a9a19a1403bb0ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f96b7dac13075f8cf73c859377fd1d78e40b1dad5d5d148b4f9b1df92c72d19d"></a>

## re_acl.fast_acl_rules.port — re_acl.fast_acl_rules.port / c0d512020f41 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- re_acl.fast_acl_rules.port

<a id="canonical-08053f518a38b35be8d4c02af76fcd6ce1d9bc411d86ea5a65e64ae2e20136a3"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-7a76932f906387cab31b2b7a476fe6f5113a7f2b521526103a73b784bbd305d3"></a>

## Direct properties — re_acl.fast_acl_rules.port / c0d512020f41 / 3

- [all](data-sources--fast_acl--reference--group-001.md#canonical-5003e6d1d549980e5b5f4cc251d4c4e2c04697ecafb71247c7f16a51c8e5f37d): complete subsection reference.

- [dns](data-sources--fast_acl--reference--group-001.md#canonical-d3531ed1cfdadf2fb667e8d7bdf64b29ee77b7dc792bdf625c426d5852b9af32): complete subsection reference.

<a id="canonical-75cbf4dea52bc3640fc239b1cd76a029eaa1be7f84043fb82db7513e99b73b8f"></a>

<a id="canonical-29906189d760bf574516271f4a028bf359590a490677b3dd48fb777caacf2ffb"></a>

## user_defined property — re_acl.fast_acl_rules.port / c0d512020f41 / 4

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-7382a6ba375290b853a6796c6f472e10136833f7dfeb50d2294471a14087b9e7"></a>

## Next pages — re_acl.fast_acl_rules.port / c0d512020f41 / 5

- [re_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-5003e6d1d549980e5b5f4cc251d4c4e2c04697ecafb71247c7f16a51c8e5f37d)
- [re_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-d3531ed1cfdadf2fb667e8d7bdf64b29ee77b7dc792bdf625c426d5852b9af32)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-5003e6d1d549980e5b5f4cc251d4c4e2c04697ecafb71247c7f16a51c8e5f37d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-099f3488ac8fc3eee9ef21eb7068d02d7713f39ee09064bd85e200836d9ae678"></a>

## re_acl.fast_acl_rules.port.all — re_acl.fast_acl_rules.port.all / d419377d2747 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-780c68f973c008f6d087e74bec354e063cef81ed3ededb002a9a19a1403bb0ea)
- re_acl.fast_acl_rules.port.all

<a id="canonical-fc967209e6b3cec63f74a63f86c2aeb77fa3124acd0de772b27464925d2298d1"></a>

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

<a id="canonical-142da9a87687e031b445ffd1584846d0390ea192106a5896b01c4437a1ad514a"></a>

## Direct properties — re_acl.fast_acl_rules.port.all / d419377d2747 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b71ed4de00149ef8af34a577086fe49fcaddfcbdb5b1dc7184b875057dd09a65"></a>

## Next pages — re_acl.fast_acl_rules.port.all / d419377d2747 / 4

- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-780c68f973c008f6d087e74bec354e063cef81ed3ededb002a9a19a1403bb0ea)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-d3531ed1cfdadf2fb667e8d7bdf64b29ee77b7dc792bdf625c426d5852b9af32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dbf02eaafb17daa777326f72ae35897395fadeabd4da371abd6a57307d7d615"></a>

## re_acl.fast_acl_rules.port.dns — re_acl.fast_acl_rules.port.dns / 09e85661ef94 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-780c68f973c008f6d087e74bec354e063cef81ed3ededb002a9a19a1403bb0ea)
- re_acl.fast_acl_rules.port.dns

<a id="canonical-61257e35b07cca4a73d569c64dd621f2920830b87dd9b23e10ac80b3fa9ce98b"></a>

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

<a id="canonical-09dbf698247988b955bc4dbad2aa24673ff9c232cf0dbcfaef79d132455c7969"></a>

## Direct properties — re_acl.fast_acl_rules.port.dns / 09e85661ef94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-887147dc4c81f6c6ab1d63fdb21863cbe817b761d29a9ed11101b0d18ee459a9"></a>

## Next pages — re_acl.fast_acl_rules.port.dns / 09e85661ef94 / 4

- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-780c68f973c008f6d087e74bec354e063cef81ed3ededb002a9a19a1403bb0ea)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-bf6532255eb18c04e83e26c9e85d85b631b1a26ae8dcdc3cf0cf7880dcd02078"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5409b665ba056f7c398af5b5c88c688c61c5fe6cb5eda01e67b7a9c53973059a"></a>

## re_acl.fast_acl_rules.prefix — re_acl.fast_acl_rules.prefix / b29ec5c3d4a1 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- re_acl.fast_acl_rules.prefix

<a id="canonical-246d632009911bbd2fbb5a0d4ed91a38123de20e6e43ce3f5acb3855e95205a3"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-5e548500e9469053d4013f085b315c43d799def1fa279445f35312cfe0c604b7"></a>

## Direct properties — re_acl.fast_acl_rules.prefix / b29ec5c3d4a1 / 3

<a id="canonical-95330d744b89b126c87c96ec9e0441510a18c16eb681f26a5388e59e3c73ad61"></a>

<a id="canonical-59314812f2a205fdca2c4bb8bc32d9f12325383f58c3243cc3aa6bbad780182b"></a>

## prefix property — re_acl.fast_acl_rules.prefix / b29ec5c3d4a1 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-aaa279cc898255777bec5329abd63c031df2d23216cb9ae1ab9d0c151e44b45c"></a>

## Next pages — re_acl.fast_acl_rules.prefix / b29ec5c3d4a1 / 5

- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-45172738c0ab335e0130d3a65b14964aceb7e051c80aee48c523a04df79a9a7e)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-64435eb04204e8a306f4cc99cac48f897a85682af06bcb77cc63bbf2c1ea9be1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fc2c2f22cc8f3c38c6230e565495360eb9f7ff5f65a44348050c91ad8c4b950"></a>

## re_acl.selected_tenant_vip — re_acl.selected_tenant_vip / 6d7f75f77387 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- re_acl.selected_tenant_vip

<a id="canonical-c7c406df3493ef390fdcc9e7d12da22a11c451eff301a9a0241565f3001219f0"></a>

Type: `"single"`. Computed.

Specific Tenant VIP. Select various tenant public VIP(s)

Upstream description:

Select various tenant public VIP(s)

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

<a id="canonical-5e9a216bca780a1ca7f550733c53923e56ccb3c0fc02d050ac2bbc7b4b9bf77c"></a>

## Direct properties — re_acl.selected_tenant_vip / 6d7f75f77387 / 3

<a id="canonical-d9238949c8747a2730f75d1cada4c24413a31d302a38431c54ce7be73f7f94c4"></a>

<a id="canonical-84fc37259397fe7bed0f6dacbae884bdccae89baa663e0275774c5b97cb9ffe1"></a>

## default_tenant_vip property — re_acl.selected_tenant_vip / 6d7f75f77387 / 4

Type: `"bool"`. Computed.

Include tenant VIP in list of specific VIP(s).

Upstream description:

Include tenant VIP in list of specific VIP(s)

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

- [public_ip_refs](data-sources--fast_acl--reference--group-001.md#canonical-f78422a2d83565d740f3592d8ff62c873335019ae69d9512905fa4c4db1e9b29): complete subsection reference.

<a id="canonical-5d129fc83cc6b290f777e97285d1e43f6e62ddb7f57119876e0b9befb9d35ffe"></a>

## Next pages — re_acl.selected_tenant_vip / 6d7f75f77387 / 5

- [re_acl.selected_tenant_vip.public_ip_refs](data-sources--fast_acl--reference--group-001.md#canonical-f78422a2d83565d740f3592d8ff62c873335019ae69d9512905fa4c4db1e9b29)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-f78422a2d83565d740f3592d8ff62c873335019ae69d9512905fa4c4db1e9b29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae9fc818ebcc55c3b73196fc99b4ae99591269986019a79fc3baa088045151c1"></a>

## re_acl.selected_tenant_vip.public_ip_refs — re_acl.selected_tenant_vip.public_ip_refs / 87893ffb90b8 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-6e9f17ccc6ea24fd3d4f6da1b6ab8388a683ea3b1857db094c03e61f35fe97a9)
- [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-64435eb04204e8a306f4cc99cac48f897a85682af06bcb77cc63bbf2c1ea9be1)
- re_acl.selected_tenant_vip.public_ip_refs

<a id="canonical-8cada62ffabc2421b04cc3a18abe1b5e1bf2c6cb5781a19d29be737e49515032"></a>

Type: `"list"`. Computed.

Select Public VIP(s). Select additional public VIP(s)

Upstream description:

Select additional public VIP(s)

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-01d77478d04b8d350673ca9728c516f1e4ef45cbe1f49ba0b591ff0a3d7afdd6"></a>

## Direct properties — re_acl.selected_tenant_vip.public_ip_refs / 87893ffb90b8 / 3

<a id="canonical-db30b51683b2a0cd1b28b0334ff35d3a5fa477787cc3310b11f16aa7f9058ee0"></a>

<a id="canonical-0ef73b0ed802f8c90b13dfc05ad89a2ae675853a275fb004a0dc1d63c707da46"></a>

## name property — re_acl.selected_tenant_vip.public_ip_refs / 87893ffb90b8 / 4

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

<a id="canonical-e81898139a4a129ce82da6b24bc39a109f0fd2b7140ffab1763c9bb988ae9889"></a>

<a id="canonical-4753d99f51a8141cf1f8df3d8fd3108a569f9f229eeecdee0f6523f008549910"></a>

## namespace property — re_acl.selected_tenant_vip.public_ip_refs / 87893ffb90b8 / 5

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

<a id="canonical-9e05fed77ebfdd06be1eadc3dd4e6b0300b884dbb89f295dd433c452a2570b16"></a>

<a id="canonical-c3b372ff76a884900c7a2b41daf170a1e789ecb9443a36d4e13e041b39b27d9e"></a>

## tenant property — re_acl.selected_tenant_vip.public_ip_refs / 87893ffb90b8 / 6

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

<a id="canonical-24d53fcad0d545aa1dc176340d2218c5231d8d440efbc578d45b6769e86770c5"></a>

## Next pages — re_acl.selected_tenant_vip.public_ip_refs / 87893ffb90b8 / 7

- [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-64435eb04204e8a306f4cc99cac48f897a85682af06bcb77cc63bbf2c1ea9be1)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c324979504fdbd680810918f4b66de3af416dcfcb28b982fa1276384f374621"></a>

## site_acl — site_acl / efc574eeed79 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- site_acl

<a id="canonical-3d945c785f0574b5017b265c0815132febbea769c7750e97a943c3532927aad4"></a>

Type: `"single"`. Computed.

Fast ACL for Site. Fast ACL definition for Site.

Upstream description:

Fast ACL definition for Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]",
  "x-ves-oneof-field-vip_choice": "[\"all_services\",\"interface_services\",\"vip_services\"]"
}
```

<a id="canonical-853a56c841bc5298074d257db5fe8d032875dd7dd19bfaffbf407956d57d7314"></a>

## Direct properties — site_acl / efc574eeed79 / 3

- [all_services](data-sources--fast_acl--reference--group-001.md#canonical-1d4d722bc3172a4b4eb81e0dbb4c5f0d7801ee84fdd381c715443d2f73b5c927): complete subsection reference.

- [fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b): complete subsection reference.

- [inside_network](data-sources--fast_acl--reference--group-001.md#canonical-40cd9bb507749d46a1b3ebf382884e81d72fda626acbf09690d8f9cccda71307): complete subsection reference.

- [interface_services](data-sources--fast_acl--reference--group-001.md#canonical-72acc6c00728ec721862f52430fe4b9396a36174ccf7dc1cf19f6a93f1e64b17): complete subsection reference.

- [outside_network](data-sources--fast_acl--reference--group-001.md#canonical-67028f43a61b038a712ee3851d7c34edce7cc11056270c9de22f4c541a8925c8): complete subsection reference.

- [vip_services](data-sources--fast_acl--reference--group-001.md#canonical-bcd187b80c5ce4c0908b0af7f977be294d1f832160c87440bc199c965e9100be): complete subsection reference.

<a id="canonical-02ae5f3c33ee1419b99c90fdb92d0b3b6dec07875bb12784922b5d9f9797a287"></a>

## Next pages — site_acl / efc574eeed79 / 4

- [site_acl.all_services](data-sources--fast_acl--reference--group-001.md#canonical-1d4d722bc3172a4b4eb81e0dbb4c5f0d7801ee84fdd381c715443d2f73b5c927)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.inside_network](data-sources--fast_acl--reference--group-001.md#canonical-40cd9bb507749d46a1b3ebf382884e81d72fda626acbf09690d8f9cccda71307)
- [site_acl.interface_services](data-sources--fast_acl--reference--group-001.md#canonical-72acc6c00728ec721862f52430fe4b9396a36174ccf7dc1cf19f6a93f1e64b17)
- [site_acl.outside_network](data-sources--fast_acl--reference--group-001.md#canonical-67028f43a61b038a712ee3851d7c34edce7cc11056270c9de22f4c541a8925c8)
- [site_acl.vip_services](data-sources--fast_acl--reference--group-001.md#canonical-bcd187b80c5ce4c0908b0af7f977be294d1f832160c87440bc199c965e9100be)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-1d4d722bc3172a4b4eb81e0dbb4c5f0d7801ee84fdd381c715443d2f73b5c927"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5657065f4398ee12609c214413f975388d8ebaf35fb64693270abc0d8eb23aee"></a>

## site_acl.all_services — site_acl.all_services / aa72114ab9d9 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- site_acl.all_services

<a id="canonical-9cffa995690cd60baf94affeff9f4d59b94274dd0293532c1681dec4e29d7536"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all services.

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

<a id="canonical-5156b6d1761f682ce616a155bd33140fad7115d3b38e97e7fea911018b0bb3f1"></a>

## Direct properties — site_acl.all_services / aa72114ab9d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-13c82e9052ba70449179d969c2c381c0cfad5c322f21250395f5d5e0ea746cf9"></a>

## Next pages — site_acl.all_services / aa72114ab9d9 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54d18bda1f5b8cfe4e1ae234eae5c23ec9750906fd6e8fa7ff8740d82da16cc3"></a>

## site_acl.fast_acl_rules — site_acl.fast_acl_rules / 02156c1915df / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- site_acl.fast_acl_rules

<a id="canonical-d671947fe892c3b54bd68420bdab29552b81ba79eb1caad84331c843f1eff747"></a>

Type: `"list"`. Computed.

Rules. Fast ACL rules to match.

Upstream description:

Fast ACL rules to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-9b0d3fb6df721535288223a73ac6058b18e817db3dd98cb3302e134eb26cf262"></a>

## Direct properties — site_acl.fast_acl_rules / 02156c1915df / 3

- [action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e): complete subsection reference.

- [ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-17d8384fb2f7196707c96de88187e5dcadaea73a2c3a53abc745c1043342c5e7): complete subsection reference.

- [metadata](data-sources--fast_acl--reference--group-001.md#canonical-1cc12cf5254918c2a877bb782098ed1f5469870c8f1700814d950d85fcc71197): complete subsection reference.

- [port](data-sources--fast_acl--reference--group-001.md#canonical-79552645af2fbb9c5bd8e555415685db17611aabcb62c3dcade3bbc28da195b6): complete subsection reference.

- [prefix](data-sources--fast_acl--reference--group-001.md#canonical-ceab865d60bd379565ac7a7dce3ecdc0f2fd40c39d4f77cebe5d5bf6d00acb8f): complete subsection reference.

<a id="canonical-28512399627bf5c290b556e0ef820679a80c650019d27ff7c06a9774f0d52dac"></a>

## Next pages — site_acl.fast_acl_rules / 02156c1915df / 4

- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e)
- [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-17d8384fb2f7196707c96de88187e5dcadaea73a2c3a53abc745c1043342c5e7)
- [site_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-1cc12cf5254918c2a877bb782098ed1f5469870c8f1700814d950d85fcc71197)
- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-79552645af2fbb9c5bd8e555415685db17611aabcb62c3dcade3bbc28da195b6)
- [site_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-ceab865d60bd379565ac7a7dce3ecdc0f2fd40c39d4f77cebe5d5bf6d00acb8f)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97273077cc3ee4e7c29e8304457ceacb07f75d4bd58c8ba609bcf35fcce40056"></a>

## site_acl.fast_acl_rules.action — site_acl.fast_acl_rules.action / 1f4fe4f0d0e8 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- site_acl.fast_acl_rules.action

<a id="canonical-ea0176e8cc84c049a33614f81e63e2b79dce09109aa154325dabb61b64add879"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

<a id="canonical-53a28ed89624160aa33db4205ea485a918c22bb1cc743ab20bbe0d604a9d4374"></a>

## Direct properties — site_acl.fast_acl_rules.action / 1f4fe4f0d0e8 / 3

- [policer_action](data-sources--fast_acl--reference--group-001.md#canonical-405033ed1b42798ba7596278416c7180fbd752bf9184e3e1c9c67c9741a3befe): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-a327435eedd33b3bfeb438a38304b96101b3e79d6b7650b88af142982d4e7297): complete subsection reference.

<a id="canonical-c796799e68b92ea2eb9535ae891c68dff2af27e1a82ffaa4d4628d923f588542"></a>

<a id="canonical-55c28d4649f4db25d889bdd5c8ea6a6b6f3689044cc875f4e3d5387ac3e340a4"></a>

## simple_action property — site_acl.fast_acl_rules.action / 1f4fe4f0d0e8 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0bd5a486d1b8b3884ba735cea507e333075e07c7db3d731d2bef92e027d19d0d"></a>

## Next pages — site_acl.fast_acl_rules.action / 1f4fe4f0d0e8 / 5

- [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-405033ed1b42798ba7596278416c7180fbd752bf9184e3e1c9c67c9741a3befe)
- [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-a327435eedd33b3bfeb438a38304b96101b3e79d6b7650b88af142982d4e7297)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-405033ed1b42798ba7596278416c7180fbd752bf9184e3e1c9c67c9741a3befe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ffe8a408470bd756677420f7b288ba268afc6dc096f0df195ddb0ea5c27a762"></a>

## site_acl.fast_acl_rules.action.policer_action — site_acl.fast_acl_rules.action.policer_action / 0c64efe41bc0 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e)
- site_acl.fast_acl_rules.action.policer_action

<a id="canonical-a2bb86ad1e52b19e8a89296504fa5b4b29a38170590b24f5a33432214066ce70"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-03e7ea77683bdb3b276f41b046832bb494a70f49bbc624bd59d18c32cd5a7bcb"></a>

## Direct properties — site_acl.fast_acl_rules.action.policer_action / 0c64efe41bc0 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-8f0e81f7fef10f09d02368c8b5e849295affde67fb21177fd1dec5c3a36e4d08): complete subsection reference.

<a id="canonical-e456dac4f0ae1c4ce52239be8d883c0695533366e6318644f0257d890324a15a"></a>

## Next pages — site_acl.fast_acl_rules.action.policer_action / 0c64efe41bc0 / 4

- [site_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-8f0e81f7fef10f09d02368c8b5e849295affde67fb21177fd1dec5c3a36e4d08)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-8f0e81f7fef10f09d02368c8b5e849295affde67fb21177fd1dec5c3a36e4d08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d72e175ade539572620e2a0d4e42487a5b5aa86419b470aad381c3bec35bb96d"></a>

## site_acl.fast_acl_rules.action.policer_action.ref — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e)
- [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-405033ed1b42798ba7596278416c7180fbd752bf9184e3e1c9c67c9741a3befe)
- site_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-3a7cc7842dc16301408689e69fe71f71024d38f10a626c3acec8191c342ae8a0"></a>

Type: `"list"`. Computed.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

<a id="canonical-6f45eda68756fbdd9c42cde2ca1927a9143a2df11b8201e3f686157ffcc44a45"></a>

## Direct properties — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 3

<a id="canonical-8a63145e3b3c65e377a63fe6df4dd348b22c55780fa20fdab4b8c164b4440a92"></a>

<a id="canonical-a0b8e1e398de3468599cfc74dde3123df55bfafe562cb56d610c198b1ebff5a9"></a>

## kind property — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 4

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

<a id="canonical-b3dfd4926254b64068de5a19c112a3c848a4562fe10ea07ead6f2e223990b66f"></a>

<a id="canonical-9fd8318b706c062e1f01784abc2e1321d4217993062b99ae1cae1a903ce87b07"></a>

## name property — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 5

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

<a id="canonical-592df430c8ef832495dbb1e959ceda03b951c479c6d9b742f8d9a4835a9d9b69"></a>

<a id="canonical-ee5b335bd4c00eee5547a84440bf97a4aa2ee7441eeaf7d32e7c236d697951a7"></a>

## namespace property — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 6

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

<a id="canonical-55551eaba6c45d2e0b94549a4fe5866b2cb583c70087ba7e042b823902e30812"></a>

<a id="canonical-27df4c69c106f45d65db56fe3549e09e9601ae9811728bed65966bdc10f7dad8"></a>

## tenant property — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 7

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

<a id="canonical-4dc9802c0f648a0a45d4fa70ae25b7361875a1496997af7dcdae2e6775d70b87"></a>

<a id="canonical-3720d7e71ffdda0dc6a869279624e69675cb682b632e5f84a46642576224b491"></a>

## uid property — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 8

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

<a id="canonical-94802fd254a8a148776ff6026f1a02b7d5344b2be1521b59d9d992b7ac44ccfa"></a>

## Next pages — site_acl.fast_acl_rules.action.policer_action.ref / 8398a0e59740 / 9

- [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-405033ed1b42798ba7596278416c7180fbd752bf9184e3e1c9c67c9741a3befe)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-a327435eedd33b3bfeb438a38304b96101b3e79d6b7650b88af142982d4e7297"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65b38313397dacafd7d41ca9fa5cd7f79fdaf71331df33d8d58a9c63a2be723c"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action — site_acl.fast_acl_rules.action.protocol_policer_action / f2c8228e5d8f / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e)
- site_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-bf5239aad5e607d2eb569509afa69cde76e6d1990f44630ca44a7b24ece67fbd"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-60f46581b664127137bee1baa7c2272aa2181ad95a30f6b6be7f574f3a6002f7"></a>

## Direct properties — site_acl.fast_acl_rules.action.protocol_policer_action / f2c8228e5d8f / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-36d9831c680f8d2cfbcc7fcf981f254d1d2ce90d0d97d1ab9103e7c158ef8182): complete subsection reference.

<a id="canonical-6b360dda86cbb834d5a5748dcb01be04567671370b93e09afe76818fe3244c01"></a>

## Next pages — site_acl.fast_acl_rules.action.protocol_policer_action / f2c8228e5d8f / 4

- [site_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-36d9831c680f8d2cfbcc7fcf981f254d1d2ce90d0d97d1ab9103e7c158ef8182)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-36d9831c680f8d2cfbcc7fcf981f254d1d2ce90d0d97d1ab9103e7c158ef8182"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cf4368af7fdb0c3f485f2140665118184e32c4d2e8cf891d9e89c3a740f80aa"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action.ref — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0965912cb9fd8e0da9ad458efc8ed8029e21fb8e4cadcd67c6c678ca6a92ab3e)
- [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-a327435eedd33b3bfeb438a38304b96101b3e79d6b7650b88af142982d4e7297)
- site_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-1846bdd5bc0ed79e195d1ebd4af21d8a44edf5995bc90234430cd3cb8fdd1c45"></a>

Type: `"list"`. Computed.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

<a id="canonical-d3c44b188fe15e479894b0ac491b016170ec848d8b095d38a884bf3d61cb1448"></a>

## Direct properties — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 3

<a id="canonical-16bf8d130e2c56fbfc4a6dd1e53b53eb4e81b3b9ff14c3cfff5d0e93cde36699"></a>

<a id="canonical-7c42afb2650a3835924d9a2e84130e1e5b5aa41992a19ea468e03e5dc6f57490"></a>

## kind property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 4

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

<a id="canonical-8a4ee362c3bed63861cec36a6057e09199e8b626b45e69e369a075a30bc3e9c4"></a>

<a id="canonical-e750353430b6c950763e4d1038ea52b908c2b40a30afceda65a1c878bb79b785"></a>

## name property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 5

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

<a id="canonical-49c1719a4167f1a51f98a189e53d320d721b50b6ca09394d09ee59c371e5a1e8"></a>

<a id="canonical-3be5b739905dd7cc6f4654ddb743db46d8c9c832364dc73ff95435b3518fec85"></a>

## namespace property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 6

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

<a id="canonical-f477c22d26039c715015f9d9b892ba4a2c5d88d933d5f3252d745c0d94451687"></a>

<a id="canonical-d5c723eeb2989822ecc2998f6a6c5c2e5f3fff98a35f6e79dc19ca82f4d116b8"></a>

## tenant property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 7

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

<a id="canonical-1426e3df7ec54429cea76e89bb76c8d0f24fb867c0888a8f1c6cb0b1e5a5191c"></a>

<a id="canonical-c71e97174ee2ab9927664081ef1fcc440392c54f8dc5a893a1d814f0271f7321"></a>

## uid property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 8

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

<a id="canonical-11768a71baffcfde7b5f70de9d35ac40ba6b5f86f9a28058afc7e4b88c5a3f85"></a>

## Next pages — site_acl.fast_acl_rules.action.protocol_policer_action.ref / 4b7872dd1f89 / 9

- [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-a327435eedd33b3bfeb438a38304b96101b3e79d6b7650b88af142982d4e7297)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-17d8384fb2f7196707c96de88187e5dcadaea73a2c3a53abc745c1043342c5e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22bde5565d4e28da48d3f181bdba21673795932acf7340e6e4e5f4040568aad1"></a>

## site_acl.fast_acl_rules.ip_prefix_set — site_acl.fast_acl_rules.ip_prefix_set / 97ee3c778d88 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- site_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-b50f3aa986f8f40f78231f460162399b6b0df9ca7b198435c6f840ad57bf68a0"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-6c02190e13c6b18cebacbc88d92bc99e2506c902d70c5d2b92cb5b68ab3009ed"></a>

## Direct properties — site_acl.fast_acl_rules.ip_prefix_set / 97ee3c778d88 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-f0aa54c17b1466e3708a74fbddee9b10f900c155cd7807fea10efa04c762664e): complete subsection reference.

<a id="canonical-3e60846d97ce9ec63b5ff69f7545dd4cbad3778116f45c03878230e3c7ca5eff"></a>

## Next pages — site_acl.fast_acl_rules.ip_prefix_set / 97ee3c778d88 / 4

- [site_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-f0aa54c17b1466e3708a74fbddee9b10f900c155cd7807fea10efa04c762664e)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-f0aa54c17b1466e3708a74fbddee9b10f900c155cd7807fea10efa04c762664e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ed6ca21813e4bcbad14e5a97377ee61b1307290e6434d2050c71d214eebe7ad"></a>

## site_acl.fast_acl_rules.ip_prefix_set.ref — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-17d8384fb2f7196707c96de88187e5dcadaea73a2c3a53abc745c1043342c5e7)
- site_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-cd7a46ec3d3c49766f74d3aa6d64cdf950d1d0d8aac75540bac91940f515c28d"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-ff50ce1f5c23c96c1cc9456b19efced3f2b8479640db187334af69e2376160dd"></a>

## Direct properties — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 3

<a id="canonical-8a9e1c15cd3f1ccd18016dc1ecb7bf97782dc207807249b8f07e8c76f813e3f9"></a>

<a id="canonical-21192e333168f3a2219463fc32d2e6887649f9200c12aa624f5f1907659881c1"></a>

## kind property — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 4

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

<a id="canonical-e91bf233813c4bb8fce3b833e0789aaa4a820dd731a2e5d9a39c5f1bd64bda2c"></a>

<a id="canonical-cb703e6eceaea43b70075d9e30e083ba39d7b0b6251622ff9ed5f41b76f06a2b"></a>

## name property — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 5

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

<a id="canonical-24c91d7ed2747721c7c4b29a5908ab4fe76e618bd0be9a1b5631fd6df5563487"></a>

<a id="canonical-85dfd92f2d967e773af24c45a0fb96cc1b0e951fc0b5e43d37170f6fcab940c2"></a>

## namespace property — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 6

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

<a id="canonical-fc59fdc47e0233f2994203537d78ff6e972e207aca1876b0e418b22b81bb748b"></a>

<a id="canonical-38827bc497843eec7d2daba80da237d1fa41cc83260463e2eda54ddf68fde125"></a>

## tenant property — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 7

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

<a id="canonical-c3e117c785e2ce616524cb01cef6d3e47161662c34b0c1956534595242e137ce"></a>

<a id="canonical-421d903fa15785a834e581646cb9c239a3e2b671a76245e83018f57152b2d052"></a>

## uid property — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 8

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

<a id="canonical-1e817ee08e389b4739fc5ca591282938cad5237fef7f0aa40a80eed8762a33b1"></a>

## Next pages — site_acl.fast_acl_rules.ip_prefix_set.ref / 4929a964124f / 9

- [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-17d8384fb2f7196707c96de88187e5dcadaea73a2c3a53abc745c1043342c5e7)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-1cc12cf5254918c2a877bb782098ed1f5469870c8f1700814d950d85fcc71197"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-661ecce84f83b392a4fea427ee791bcb60121059f3b22c01b2a1dd185b982abf"></a>

## site_acl.fast_acl_rules.metadata — site_acl.fast_acl_rules.metadata / d531881f887d / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- site_acl.fast_acl_rules.metadata

<a id="canonical-87d95569b2bdbb2812d51b60c03d4f568ca8160a99ac7bb271d6065c5e6a2ba3"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-54f6075da292b804f3ce2e60194abfda1b629efbe477579e92e2e013bfe0cc1d"></a>

## Direct properties — site_acl.fast_acl_rules.metadata / d531881f887d / 3

<a id="canonical-c9c30456b16c669a4181e41b9e40bcd1e64fdd104e7ba319d1029d213eb03f1a"></a>

<a id="canonical-a52a23054b14240cab4a66c35fb3637cf10ff4a40a01bf53a49ba2ed3b940707"></a>

## description_spec property — site_acl.fast_acl_rules.metadata / d531881f887d / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-cd850cfc773f778f691aa05eeba38001f679d78e52d4f082d5dd91e5c82c3102"></a>

<a id="canonical-5f41433e65997767e8259546d17e446282ccbf998a8b9b9091c466f193b407f5"></a>

## name property — site_acl.fast_acl_rules.metadata / d531881f887d / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2c24a76d931bae1d2b4f88d220c507f33531b3beee89e589e98ab0ff26c73f5e"></a>

## Next pages — site_acl.fast_acl_rules.metadata / d531881f887d / 6

- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-79552645af2fbb9c5bd8e555415685db17611aabcb62c3dcade3bbc28da195b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69e939492b76b389870cfb6dc1c8cf134a95eab400f7b7f328d7102d00f96753"></a>

## site_acl.fast_acl_rules.port — site_acl.fast_acl_rules.port / b2d7f88c487b / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- site_acl.fast_acl_rules.port

<a id="canonical-c4bdf8c888cebd718b65fd2127821a75d12193e3beadfad5431c2116b6493318"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-44e6a40da1bc26f9713be72fd84dd8babf576a98342901babf756d453fb6f8c1"></a>

## Direct properties — site_acl.fast_acl_rules.port / b2d7f88c487b / 3

- [all](data-sources--fast_acl--reference--group-001.md#canonical-ca99d6019210e47e3a0baaa5b148bdc1b9d965465e83c6ccb00616feb1be1dd5): complete subsection reference.

- [dns](data-sources--fast_acl--reference--group-001.md#canonical-a4a61c1c24b13538844dd3227e9830ad4368db5efe68e9a7f78255f6f39d9e79): complete subsection reference.

<a id="canonical-d59efec33ec8733fb11113b714d8c744fc2fd934241054842affe814ad6e21d4"></a>

<a id="canonical-3e3d5cecce1c49627169e094beb61a53f545f9b1e9b66a061d6ba78fbfd2e520"></a>

## user_defined property — site_acl.fast_acl_rules.port / b2d7f88c487b / 4

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-28de6b3a0784524c167998989e66a4088692bce552239f5bb8265e6a008ab36c"></a>

## Next pages — site_acl.fast_acl_rules.port / b2d7f88c487b / 5

- [site_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-ca99d6019210e47e3a0baaa5b148bdc1b9d965465e83c6ccb00616feb1be1dd5)
- [site_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-a4a61c1c24b13538844dd3227e9830ad4368db5efe68e9a7f78255f6f39d9e79)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-ca99d6019210e47e3a0baaa5b148bdc1b9d965465e83c6ccb00616feb1be1dd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2b7232ed69751ce93f98042bc3cd71b154ab4ac4eac8cf47ccdc4679f0d5bd3"></a>

## site_acl.fast_acl_rules.port.all — site_acl.fast_acl_rules.port.all / a3f785a115a9 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-79552645af2fbb9c5bd8e555415685db17611aabcb62c3dcade3bbc28da195b6)
- site_acl.fast_acl_rules.port.all

<a id="canonical-656c0c8b74c346feb3a8bc31d56815361a291bec86ce349597a78649d726fbda"></a>

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

<a id="canonical-01773ecaaa98cf9a87a72a92caceac8d8207160988473de2749db6189ce5004f"></a>

## Direct properties — site_acl.fast_acl_rules.port.all / a3f785a115a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21469c333999afbe279c0eed6979651e492f5fa31b5d80d2c5be4ebbdb5bb834"></a>

## Next pages — site_acl.fast_acl_rules.port.all / a3f785a115a9 / 4

- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-79552645af2fbb9c5bd8e555415685db17611aabcb62c3dcade3bbc28da195b6)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-a4a61c1c24b13538844dd3227e9830ad4368db5efe68e9a7f78255f6f39d9e79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15623709dd5d4393a009f1a9b830cc3af2aa57f9de5a3f0c468922d0dcdba525"></a>

## site_acl.fast_acl_rules.port.dns — site_acl.fast_acl_rules.port.dns / 4a88ff55ecfe / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-79552645af2fbb9c5bd8e555415685db17611aabcb62c3dcade3bbc28da195b6)
- site_acl.fast_acl_rules.port.dns

<a id="canonical-ceb65f5d04ea3d59542591e739f5ca2a3a0c391fa7007abe4edb0481c248c67c"></a>

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

<a id="canonical-d0c42ff09460e009c5372fc3272a4718f9fe4bbbe2b4f5d3a34439029ef8f16d"></a>

## Direct properties — site_acl.fast_acl_rules.port.dns / 4a88ff55ecfe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b004e8225fe49b9d4f33d3ae6f93403c2566c03f1afc39f1827a47bfa144399"></a>

## Next pages — site_acl.fast_acl_rules.port.dns / 4a88ff55ecfe / 4

- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-79552645af2fbb9c5bd8e555415685db17611aabcb62c3dcade3bbc28da195b6)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-ceab865d60bd379565ac7a7dce3ecdc0f2fd40c39d4f77cebe5d5bf6d00acb8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b697c61c5942fac9086ffa35764bbaf911b24b17026e3837bf6e0b27ea76b75b"></a>

## site_acl.fast_acl_rules.prefix — site_acl.fast_acl_rules.prefix / 4c06382744f8 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- site_acl.fast_acl_rules.prefix

<a id="canonical-2bc3b076936a0e5aca834c1c377a27bf8b5ec398a93c12b1d00b83df96075ac3"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-c88a56073d948d8879533b46d919da9624a1d4ed2c8739c894c4cd5149669f57"></a>

## Direct properties — site_acl.fast_acl_rules.prefix / 4c06382744f8 / 3

<a id="canonical-758740625ef8a197992dbf3a966b459a2e0311d4aa76a924ebe031255a923c59"></a>

<a id="canonical-d945ba0ce87ad9954bfb04489f3f24b77604c296f01efd01cc9802a8fbb47ea1"></a>

## prefix property — site_acl.fast_acl_rules.prefix / 4c06382744f8 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-21c784150c0d95a5e441dfe147bdc12d76c5eeba2285d000528dc8025595cb32"></a>

## Next pages — site_acl.fast_acl_rules.prefix / 4c06382744f8 / 5

- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0f918148a6f435d82dbd1fb1fc3f3ab02ee50aa4af87b7d71d232b996259ac0b)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-40cd9bb507749d46a1b3ebf382884e81d72fda626acbf09690d8f9cccda71307"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81c847a885a7981606e460b2cfa8b032b153e2ecbb8e8d5385ea280eda79aa17"></a>

## site_acl.inside_network — site_acl.inside_network / ea9bfd7d8098 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- site_acl.inside_network

<a id="canonical-0da8f8ff10f27b1574103990527778965761ffbc2826df0c369d857fcbf9a00e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-cb75f6a83060c69fa1fed8fd4c63ec7e2bd93e89f639c522ee414befd6125e47"></a>

## Direct properties — site_acl.inside_network / ea9bfd7d8098 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1aedd4932f5ed50741f969efca767d2cc7e4645c7cb7c9644f90dc00cf67c286"></a>

## Next pages — site_acl.inside_network / ea9bfd7d8098 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-72acc6c00728ec721862f52430fe4b9396a36174ccf7dc1cf19f6a93f1e64b17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2eab47bc636abcd8b264c4d77a4d32950f2c6829fb18b1bc4d4d49b7422ed25"></a>

## site_acl.interface_services — site_acl.interface_services / abecddd69c56 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- site_acl.interface_services

<a id="canonical-49774e3b517efaa7b41ee3c8f269246d45dd7ff340e02471fbe6cf06674dfebe"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for interface services.

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

<a id="canonical-92dc401491399ec3f666921ae27456d7b13d8f46a1fae91e855c29bb5a061e18"></a>

## Direct properties — site_acl.interface_services / abecddd69c56 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-516b9cfd035e2675889adce3ef969fc030eef512767cd9edfda2668283eb6273"></a>

## Next pages — site_acl.interface_services / abecddd69c56 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-67028f43a61b038a712ee3851d7c34edce7cc11056270c9de22f4c541a8925c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e610bde53bf17682e9736aee89b54b5debf44c51a1603f900771dd648d849513"></a>

## site_acl.outside_network — site_acl.outside_network / b33a5348c69f / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- site_acl.outside_network

<a id="canonical-bbd9fef5f78a1b4da05b925d49f9f46c377cd71aaa3dafdc99fadc39b1faa133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-6123cfa05ceea08daf1e0034c20ca2e90fff7d2f5cfe0c5419acfbf3caa717f2"></a>

## Direct properties — site_acl.outside_network / b33a5348c69f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a684a160219d2507fb3a57a098a67ee9a36971d5f4c06e19896618e83ebdb9d6"></a>

## Next pages — site_acl.outside_network / b33a5348c69f / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-bcd187b80c5ce4c0908b0af7f977be294d1f832160c87440bc199c965e9100be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbf0a01d63ef601a9e6e62739ef777944474c35e4332304c28b8e406ee55e628"></a>

## site_acl.vip_services — site_acl.vip_services / 5b3257ea704f / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- site_acl.vip_services

<a id="canonical-ef72401f8fa1c369ffab2989460cf378dcb466322f410fc106835e3424d4c072"></a>

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

<a id="canonical-664b90902a8069fc343e5f41990dbd3f9c5c44bf178d5acff84f96c3cb5dbd92"></a>

## Direct properties — site_acl.vip_services / 5b3257ea704f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a807678a15613765959fc235a3cc947f00445e4052792288789507a494d845b4"></a>

## Next pages — site_acl.vip_services / 5b3257ea704f / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-54161659413efc279699da27bac136e01460c2d03937a44f8b52cf2d9a5dcab8)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
