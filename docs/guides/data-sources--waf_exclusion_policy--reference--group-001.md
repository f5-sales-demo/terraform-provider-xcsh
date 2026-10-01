---
page_title: "xcsh_waf_exclusion_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy reference."
---

# xcsh_waf_exclusion_policy reference

<a id="canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ea9d38fdec848af8b625f0fe7c25755b5f42edfa1ef9801fb34d41251e6c763"></a>

## Property reference — Property reference / dc3282ed9a79 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- Property reference

<a id="canonical-20f96a78faec52d1ae5ef8d9cffba9f241413b2601e238091961873f2f90139a"></a>

## Direct properties — Property reference / dc3282ed9a79 / 3

<a id="canonical-1c055e7d733fe539633b468308b01c1efd20f9465f2468bb004403fbdefe37f3"></a>

<a id="canonical-7c903a4078c856a071bfc9a67837417ff0bff83e7db1e492329c4d6099877275"></a>

## annotations property — Property reference / dc3282ed9a79 / 4

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

<a id="canonical-c35c606ed2a89f6265d04c095d12eb7fb432059a59b174265af6c087ad706bd0"></a>

<a id="canonical-e6c4caec0e021dfc4bffcce6610fffdb643bbe132cb7eef9529481f280872dd3"></a>

## description property — Property reference / dc3282ed9a79 / 5

Type: `"string"`. Computed.

Description of the WAFExclusionPolicy.

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

<a id="canonical-0e4b5ebe6d68f6a6a4466c37b8babe695c1d81e2ad1a9843260b493a4cc81b26"></a>

<a id="canonical-423c0c03bbb7703a90ca65dfc21e3f4f684dcf468dcf6d02a1e9eeee517825d7"></a>

## id property — Property reference / dc3282ed9a79 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0bfea7f87661ad4aecb9d5af28e5504dd5841b85e9ad64e7cb8ed4f7ce6e6128"></a>

<a id="canonical-702a3a3b263c9b7ea406e8369301abb3f7d8823716c1a9df1f31524cf7b55a71"></a>

## labels property — Property reference / dc3282ed9a79 / 7

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

<a id="canonical-c1a8947c55607ff6cad0fd953fc16804da99c8ef13f163e503aafaeaf17e3762"></a>

<a id="canonical-31fcc4618f312d97351af3409af95c87ada6c512c1f21ac93dd8ebcd41a94b9a"></a>

## name property — Property reference / dc3282ed9a79 / 8

Type: `"string"`. Required.

Name of the WAFExclusionPolicy.

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

<a id="canonical-c90d9a3281b52b977f226ee625331cf3eea3b3e3a4471d53481079579c00f789"></a>

<a id="canonical-b2e087c14569b05f56a9f8eb71e5fe40b1b30f7c64e9efb7e5cec64b0440d319"></a>

## namespace property — Property reference / dc3282ed9a79 / 9

Type: `"string"`. Required.

Namespace where the WAFExclusionPolicy exists.

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

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec): complete subsection reference.

<a id="canonical-905b3d45628f8430cc80013a8097f345a0fbb8378d98f3ce54174c8ba5b18e4c"></a>

## All schema paths — Property reference / dc3282ed9a79 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1c055e7d733fe539633b468308b01c1efd20f9465f2468bb004403fbdefe37f3) |
| `description` | [description](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c35c606ed2a89f6265d04c095d12eb7fb432059a59b174265af6c087ad706bd0) |
| `id` | [id](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0e4b5ebe6d68f6a6a4466c37b8babe695c1d81e2ad1a9843260b493a4cc81b26) |
| `labels` | [labels](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0bfea7f87661ad4aecb9d5af28e5504dd5841b85e9ad64e7cb8ed4f7ce6e6128) |
| `name` | [name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c1a8947c55607ff6cad0fd953fc16804da99c8ef13f163e503aafaeaf17e3762) |
| `namespace` | [namespace](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c90d9a3281b52b977f226ee625331cf3eea3b3e3a4471d53481079579c00f789) |
| `waf_exclusion_rules` | [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-b02c89ed974fde3069476fbfbb4fbc2d4894884fb4a9a200d53ca67676320627) |
| `waf_exclusion_rules.any_domain` | [waf_exclusion_rules.any_domain](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-e63af021bb7198213101c2cbf0a75698f7c14667d30cf91e3a3f8bbb7e30d196) |
| `waf_exclusion_rules.any_path` | [waf_exclusion_rules.any_path](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-abcd07559fda8097bd474837b49ce2c93b027000478ae0c422c2fd01b94dc3ed) |
| `waf_exclusion_rules.app_firewall_detection_control` | [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-4ce7ec989a7589565d27bf19c21a004ebb8537ceb80617f49456086cc3333223) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fba7cacca492fba3213f7d36015b854b86bd9b95a0d11b4d2c4fd440e45346a8) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-442ba436e734b510b7608593469c3a69a80ec3631a201077ddcf3ce9be414c20) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0119aeb454c999940e22b4698040099e70d670fdb16f35b6682883986a9ad783) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-b9ac661b3c4827b901ff8869830e95e27efb711efdbc5a4e5e59005e74f22213) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-96a66b2c495f6771b62b52e0adb3ff502310bf05356840a5e6e2017a52f74d57) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-4194a3109fba8e922d505bbd74cf28eda22f7937f53459dbc98c84034114695c) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-f92e3703384dc4c4e425f8e631e0a87dbb5ae7968ebe4e53417691383a0777d7) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-5ae44e08a99b0cd6d9b9a945d8029a6efe866cb047ccda6c11db08befcb88bec) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-bdc240c755525bcd4a72a67a65fef3519e8ef2712a869a5b337cc2f2f59cf365) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-f7558fac51c267ce2805c597b6853cf3a9bad0b775c420c012d3a59fe4087885) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-15e335e26ce43f79fc0f16fc79f919280452c07431d6f43e30e75f8ea050d152) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-a407a7c4eea35d5f03069c5a434de6b10d333638b5af004ec51ca2fc832a095b) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-ebd2fcd1019f2bced392bad45d600cf4b59ec8493813e5029c2cfc9f16b02c88) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-f76555425689efee3ed4bcae024879a953b3ecf49520e332bd70a15b73b02161) |
| `waf_exclusion_rules.exact_value` | [waf_exclusion_rules.exact_value](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-ce3308579c755a9686b752fb740af5e52f1a5d459b46a1d8f29ee31b3840950c) |
| `waf_exclusion_rules.expiration_timestamp` | [waf_exclusion_rules.expiration_timestamp](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-707096f3770e36e9265ad12e6546793eec06d3c1139975c10442ebb7a10da402) |
| `waf_exclusion_rules.metadata` | [waf_exclusion_rules.metadata](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-43eb0a2d49baf36e9b78f8556b1a858545e05e536942638f15ebb8820b2961d0) |
| `waf_exclusion_rules.metadata.description_spec` | [waf_exclusion_rules.metadata.description_spec](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-00e7da8bffd44df45d1b4a0ffc2a77a8c20186fc7e81ae3416e4391505477479) |
| `waf_exclusion_rules.metadata.name` | [waf_exclusion_rules.metadata.name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-cffdc33b55d95bc8d2883befaa7d5abbb38022f018171b158c8f9793b0d2f0db) |
| `waf_exclusion_rules.methods` | [waf_exclusion_rules.methods](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-a1000c5c77047a87733e7a8d360b7227f161789bdf6468d8ad9b52946881818c) |
| `waf_exclusion_rules.path_prefix` | [waf_exclusion_rules.path_prefix](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-d62d7153c19c966514b53dd84609cc28b33fa7fd96be64f73ba673cb01a174e9) |
| `waf_exclusion_rules.path_regex` | [waf_exclusion_rules.path_regex](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2e6f26b2c90268219252cbb99f749c84219df110e38f8dae9612ec1e04b0699b) |
| `waf_exclusion_rules.suffix_value` | [waf_exclusion_rules.suffix_value](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-63e51b34ca7308cb2dfe9091c4b9c467235688136554dd7f4a29b81fdc5d59cf) |
| `waf_exclusion_rules.waf_skip_processing` | [waf_exclusion_rules.waf_skip_processing](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-27e94ec923a285520c3c8afd0bc9d619af6d15a12b98567b269982af7f817c80) |

<a id="canonical-37baa7f0f7634ae7e83d67ddd03b2fbc8822fb8d0e6d72f0d2d8d850fecdaef2"></a>

## Next pages — Property reference / dc3282ed9a79 / 11

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f23e7cfdb7854f1c3e7c3e839b94e53bc036d2da296d2f620fe18a6ed208b92"></a>

## waf_exclusion_rules — waf_exclusion_rules / 1222d088f0ed / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- waf_exclusion_rules

<a id="canonical-b02c89ed974fde3069476fbfbb4fbc2d4894884fb4a9a200d53ca67676320627"></a>

Type: `"list"`. Computed.

WAF Exclusion Rules. An ordered list of rules.

Upstream description:

An ordered list of rules.

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-8a6afba0a785150f3e496b03de24d518f7c078e15010eff96dfed7c095aefbba"></a>

## Direct properties — waf_exclusion_rules / 1222d088f0ed / 3

- [any_domain](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-4db9d928a5ca25dbef5527e07c7287c3e816d4896639fa8109e359b8871a9420): complete subsection reference.

- [any_path](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-d9ff10692d99489ae681cd71aa24c3a5d1aa14adf7baeb417852ff04b0d23923): complete subsection reference.

- [app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756): complete subsection reference.

<a id="canonical-ce3308579c755a9686b752fb740af5e52f1a5d459b46a1d8f29ee31b3840950c"></a>

<a id="canonical-71c1642449a0e01ad35c2c9bb7a220ce7481da8266181345b672a7d2f9226eb2"></a>

## exact_value property — waf_exclusion_rules / 1222d088f0ed / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-707096f3770e36e9265ad12e6546793eec06d3c1139975c10442ebb7a10da402"></a>

<a id="canonical-874c9826725c45aa470b02839e6d62fdef49d3a3397b84ea7de579ff589ef041"></a>

## expiration_timestamp property — waf_exclusion_rules / 1222d088f0ed / 5

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-56b1ced51e8d2eea9f97a5b497caee35a43c21b9d3e5ba85ffefa92f87e6922a): complete subsection reference.

<a id="canonical-a1000c5c77047a87733e7a8d360b7227f161789bdf6468d8ad9b52946881818c"></a>

<a id="canonical-39f2c915be81c43bc5e192acf259bad9a54fd21261e540a383bcf57d6117f605"></a>

## methods property — waf_exclusion_rules / 1222d088f0ed / 6

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d62d7153c19c966514b53dd84609cc28b33fa7fd96be64f73ba673cb01a174e9"></a>

<a id="canonical-1b0b73ca13b4107061afd5dd34169a37958f9122d44858ec75d63c713549dc1d"></a>

## path_prefix property — waf_exclusion_rules / 1222d088f0ed / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2e6f26b2c90268219252cbb99f749c84219df110e38f8dae9612ec1e04b0699b"></a>

<a id="canonical-5d666f05ae30042d23b18baa1f75af9dd8283f2bdabfd7a929eb47780ed727db"></a>

## path_regex property — waf_exclusion_rules / 1222d088f0ed / 8

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-63e51b34ca7308cb2dfe9091c4b9c467235688136554dd7f4a29b81fdc5d59cf"></a>

<a id="canonical-1ccecbc289d1ed309ec8775c14dfcc26f06638d319b995d2f02425aa4afbc7a0"></a>

## suffix_value property — waf_exclusion_rules / 1222d088f0ed / 9

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [waf_skip_processing](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0c0209ebdd9e135ad7f510c96773a3b138b8e6bf23ee6ed0a048bffe11b6aed6): complete subsection reference.

<a id="canonical-fa27bd8fd940c85ee8c45ea948edd04bf500a1c2501e266bedfded7ddeae3757"></a>

## Next pages — waf_exclusion_rules / 1222d088f0ed / 10

- [waf_exclusion_rules.any_domain](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-4db9d928a5ca25dbef5527e07c7287c3e816d4896639fa8109e359b8871a9420)
- [waf_exclusion_rules.any_path](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-d9ff10692d99489ae681cd71aa24c3a5d1aa14adf7baeb417852ff04b0d23923)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- [waf_exclusion_rules.metadata](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-56b1ced51e8d2eea9f97a5b497caee35a43c21b9d3e5ba85ffefa92f87e6922a)
- [waf_exclusion_rules.waf_skip_processing](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0c0209ebdd9e135ad7f510c96773a3b138b8e6bf23ee6ed0a048bffe11b6aed6)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-4db9d928a5ca25dbef5527e07c7287c3e816d4896639fa8109e359b8871a9420"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9c0b14d6f1cf2577b7d08a0333cbae889fb3f5f1bd95ce836da26a0dfc082eb"></a>

## waf_exclusion_rules.any_domain — waf_exclusion_rules.any_domain / 1397b9a0d853 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- waf_exclusion_rules.any_domain

<a id="canonical-e63af021bb7198213101c2cbf0a75698f7c14667d30cf91e3a3f8bbb7e30d196"></a>

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

<a id="canonical-55e203a5d4de1fce8f2485f05420a8063cf7efbadc114fcb9db16bdafd0d44c1"></a>

## Direct properties — waf_exclusion_rules.any_domain / 1397b9a0d853 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4cc6e9a54d0cbab6df83b6e6cfbe29f9de68819f5b99d41b1af0748f1c11dc5"></a>

## Next pages — waf_exclusion_rules.any_domain / 1397b9a0d853 / 4

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-d9ff10692d99489ae681cd71aa24c3a5d1aa14adf7baeb417852ff04b0d23923"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edcc5694cc019ba7411360ce95894f4310513a01733c044188bc7e775aeab48e"></a>

## waf_exclusion_rules.any_path — waf_exclusion_rules.any_path / 47c4724ffc45 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- waf_exclusion_rules.any_path

<a id="canonical-abcd07559fda8097bd474837b49ce2c93b027000478ae0c422c2fd01b94dc3ed"></a>

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

<a id="canonical-1caa620e16449c52ea66d2a8903fb0ed4614c552042554028e675fcbafbd7a42"></a>

## Direct properties — waf_exclusion_rules.any_path / 47c4724ffc45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07566f2ca7670020458f7e1058c606ffbb90f7aca2e1d2d2e67294b369558b50"></a>

## Next pages — waf_exclusion_rules.any_path / 47c4724ffc45 / 4

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59e0cfcefd6f8e8c93953859403d1b9d9589ac602d4ef45a84566c968776add6"></a>

## waf_exclusion_rules.app_firewall_detection_control — waf_exclusion_rules.app_firewall_detection_control / 4994a290d76b / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- waf_exclusion_rules.app_firewall_detection_control

<a id="canonical-4ce7ec989a7589565d27bf19c21a004ebb8537ceb80617f49456086cc3333223"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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

<a id="canonical-bf77cfb9872dddb95a7f5d04ee39c3433cb67912e43595c0ae5f8b70e938cb16"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control / 4994a290d76b / 3

- [exclude_attack_type_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-cc6795d0934b09ee3ac1b2eb124fcc8d582af577ea1915c43b7b76420bf48ad7): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-9aeadc11393f12f310c38c4fd0cd0814c96f896dd1c3cc856ce64d8bd8a38aa8): complete subsection reference.

- [exclude_signature_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-db121929ae1fa3d8b21d07cd1e658c1b8f2f9a25ec3671f268a59dd7842ac2b2): complete subsection reference.

- [exclude_violation_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-d54939fdf50f25235ea3a53230d7f7d956ca79d97b9a154ddd8b8b832ee1f71c): complete subsection reference.

<a id="canonical-4c9de90f3fb94484589be3f7c88fc874aa435e4bd595662a97886c8619cfb845"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control / 4994a290d76b / 4

- [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-cc6795d0934b09ee3ac1b2eb124fcc8d582af577ea1915c43b7b76420bf48ad7)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-9aeadc11393f12f310c38c4fd0cd0814c96f896dd1c3cc856ce64d8bd8a38aa8)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-db121929ae1fa3d8b21d07cd1e658c1b8f2f9a25ec3671f268a59dd7842ac2b2)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-d54939fdf50f25235ea3a53230d7f7d956ca79d97b9a154ddd8b8b832ee1f71c)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-cc6795d0934b09ee3ac1b2eb124fcc8d582af577ea1915c43b7b76420bf48ad7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a974a939ff8117b5eb257747be41cd5f31c3a3155ac27e45f9397b128f6a00c"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 1398373c8e56 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-fba7cacca492fba3213f7d36015b854b86bd9b95a0d11b4d2c4fd440e45346a8"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5e4b52a27d11ffd4f4d7131b83bef353e894dc0a8a3438ee7a31fd0ff1270c76"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 1398373c8e56 / 3

<a id="canonical-442ba436e734b510b7608593469c3a69a80ec3631a201077ddcf3ce9be414c20"></a>

<a id="canonical-c2a48c06669d2bb25c9ef8e0f6ec67b2733edb24e16b7aaae71b6c390638c371"></a>

## context property — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 1398373c8e56 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0119aeb454c999940e22b4698040099e70d670fdb16f35b6682883986a9ad783"></a>

<a id="canonical-a83ad147244b660d564b41f17c9da3e363a8ee1e7d737e15c44e3fed8b33224b"></a>

## context_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 1398373c8e56 / 5

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-b9ac661b3c4827b901ff8869830e95e27efb711efdbc5a4e5e59005e74f22213"></a>

<a id="canonical-bdbdb896a0656aa21c59704e32bfe57bdcc51cb39b53e98cce8c362f418cb7c7"></a>

## exclude_attack_type property — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 1398373c8e56 / 6

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7c212ce77fde9f6263c1547748ac6347d42f3b24fb36d2c376469bfb733c82a0"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts / 1398373c8e56 / 7

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-9aeadc11393f12f310c38c4fd0cd0814c96f896dd1c3cc856ce64d8bd8a38aa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-897503b73562fdbe99a9d76738f9aa89fa844a06c67f6d5a62e8bceef89f86c6"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / e010dc6c154e / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-96a66b2c495f6771b62b52e0adb3ff502310bf05356840a5e6e2017a52f74d57"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f74bbecfb83cdd444e91c08116fa38c063233305b0954f74c09714bde65bd478"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / e010dc6c154e / 3

<a id="canonical-4194a3109fba8e922d505bbd74cf28eda22f7937f53459dbc98c84034114695c"></a>

<a id="canonical-29eef2b3bb839fd1985c80568401682c6f55b40d566f7b7e51389b735737be8b"></a>

## bot_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / e010dc6c154e / 4

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-4bfabb177289fa6b45496e95dbe0d6393f5c1ff84989d3b8b9b67a8f829d0bf8"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts / e010dc6c154e / 5

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-db121929ae1fa3d8b21d07cd1e658c1b8f2f9a25ec3671f268a59dd7842ac2b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b876d1aeaba57a56e594a729cb6f9799de80e531cc7cd27dff34ba00147dcbcb"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / ae6491088e63 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-f92e3703384dc4c4e425f8e631e0a87dbb5ae7968ebe4e53417691383a0777d7"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8940d84070eedc803ab70cd56a2e0db40511e620184a237983f30d6b3b63d3f3"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / ae6491088e63 / 3

<a id="canonical-5ae44e08a99b0cd6d9b9a945d8029a6efe866cb047ccda6c11db08befcb88bec"></a>

<a id="canonical-fb00a0dfa96671116614d6de9985f67e5f42166188cae9466422cda87558043e"></a>

## context property — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / ae6491088e63 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bdc240c755525bcd4a72a67a65fef3519e8ef2712a869a5b337cc2f2f59cf365"></a>

<a id="canonical-7f30c7038346215920178da44b7e9332ce0ec49634ea56123a34bec244cf3ceb"></a>

## context_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / ae6491088e63 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-f7558fac51c267ce2805c597b6853cf3a9bad0b775c420c012d3a59fe4087885"></a>

<a id="canonical-3954ddbf09928af4c0363f3bcb5d4817f09f5723ba2d99c42be91193907ca53a"></a>

## signature_id property — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / ae6491088e63 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-30a742f31358e42048f620f28daaa746040200a7dc2aa43600daa16eb7d127d3"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts / ae6491088e63 / 7

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-d54939fdf50f25235ea3a53230d7f7d956ca79d97b9a154ddd8b8b832ee1f71c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1044321b0376acfa603fe5e45a295af901a837b18631c242e6cae3f57340ab76"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 22bfa581df5c / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-15e335e26ce43f79fc0f16fc79f919280452c07431d6f43e30e75f8ea050d152"></a>

Type: `"list"`. Computed.

Violations to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ca2b79d6fcbb37dd5dfdb79231551479f09e26888da14fbbff78cf67296dd97d"></a>

## Direct properties — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 22bfa581df5c / 3

<a id="canonical-a407a7c4eea35d5f03069c5a434de6b10d333638b5af004ec51ca2fc832a095b"></a>

<a id="canonical-cef6eb15e60d8b539d64e7f0035f11d350a007b6d6bf65802e41886b5f70c8e1"></a>

## context property — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 22bfa581df5c / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ebd2fcd1019f2bced392bad45d600cf4b59ec8493813e5029c2cfc9f16b02c88"></a>

<a id="canonical-d9c2ea2a9ad4910fa440b1ef9258b5cc476ed6eded30d422f924d085f4b4efae"></a>

## context_name property — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 22bfa581df5c / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-f76555425689efee3ed4bcae024879a953b3ecf49520e332bd70a15b73b02161"></a>

<a id="canonical-39148e9b119e7e7ef87c935a627e840bd27d193883dc542aa0be384d7f666667"></a>

## exclude_violation property — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 22bfa581df5c / 6

Type: `"string"`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d220b210a94cb4db52f32fe124f35d26b68eb1e8d7b1a32fe4d463b2878394bf"></a>

## Next pages — waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts / 22bfa581df5c / 7

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-fad8212126db92d5d15417731bee2f0c83f4cffb3e39d7776fdb0a9b02eb4756)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-56b1ced51e8d2eea9f97a5b497caee35a43c21b9d3e5ba85ffefa92f87e6922a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74619a3d605e0669d91ea6cfbc3fe231145e95ac7eca9eaa88de116f2b09b9a5"></a>

## waf_exclusion_rules.metadata — waf_exclusion_rules.metadata / 650faf346257 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- waf_exclusion_rules.metadata

<a id="canonical-43eb0a2d49baf36e9b78f8556b1a858545e05e536942638f15ebb8820b2961d0"></a>

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

<a id="canonical-7261ca7e66239162a51c528fabfdba4fd596901df9f17f5a295bcf7e8e384503"></a>

## Direct properties — waf_exclusion_rules.metadata / 650faf346257 / 3

<a id="canonical-00e7da8bffd44df45d1b4a0ffc2a77a8c20186fc7e81ae3416e4391505477479"></a>

<a id="canonical-4a9025a78b4c31a1472ad8d2c3a1bee9b3eb27ccb15192b331ee24bf92730ee2"></a>

## description_spec property — waf_exclusion_rules.metadata / 650faf346257 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-cffdc33b55d95bc8d2883befaa7d5abbb38022f018171b158c8f9793b0d2f0db"></a>

<a id="canonical-299bc14fd6532a1623e46b3fa8945a9ac918ee690d6a5f2bb969eedf22eec034"></a>

## name property — waf_exclusion_rules.metadata / 650faf346257 / 5

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

<a id="canonical-e5e6b0c1b1facdc84d75be81c071acf6eeb1eccd62d48c315adcebe8f1f6d8c2"></a>

## Next pages — waf_exclusion_rules.metadata / 650faf346257 / 6

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-0c0209ebdd9e135ad7f510c96773a3b138b8e6bf23ee6ed0a048bffe11b6aed6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8632c8e79fd8d0cb3ec139e714e4b7443458059cbf25a0d2091e7b13170b218"></a>

## waf_exclusion_rules.waf_skip_processing — waf_exclusion_rules.waf_skip_processing / bad9964a6e91 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-c8d4b20bee0ece317a8f58e588ade2d1ade2e19613853ed5cfdb026a232d1c56)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- waf_exclusion_rules.waf_skip_processing

<a id="canonical-27e94ec923a285520c3c8afd0bc9d619af6d15a12b98567b269982af7f817c80"></a>

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

<a id="canonical-34a0fb769b0dd000bea4df613016dfb6cd5e1f400f5ecc4f5b03250a84cd193f"></a>

## Direct properties — waf_exclusion_rules.waf_skip_processing / bad9964a6e91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb776a8a0fdb7a7edf6fddfe045aaf53c4dc19e10766ce20289c728b6883bcc0"></a>

## Next pages — waf_exclusion_rules.waf_skip_processing / bad9964a6e91 / 4

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-20641a39cccdc99424431c81b7d067b0f9d9f7e4868517ee2a4c5d72e46127ec)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
