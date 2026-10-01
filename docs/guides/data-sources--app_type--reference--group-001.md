---
page_title: "xcsh_app_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type reference."
---

# xcsh_app_type reference

<a id="canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f320cfa6aa4bde7846c1242d79f3d2dceecbbed697e0cc9a253775b8f728a938"></a>

## Property reference — Property reference / 0166059ffa86 / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- Property reference

<a id="canonical-06547bbd5519b0aee1cefcf6f3423ce4cebc1112a4bb9d54d5ea888271d51a07"></a>

## Direct properties — Property reference / 0166059ffa86 / 3

<a id="canonical-9dc8b42f8743a589860cd232f603d1aee071cb04301a541ba8143a4ecde237ce"></a>

<a id="canonical-cf6bfb3db1a183793ee4015eef176d95c46c78fd77315a894363c8583f2bce71"></a>

## annotations property — Property reference / 0166059ffa86 / 4

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

- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845): complete subsection reference.

<a id="canonical-fd67a336864d3969419898df2d8e5127b85f690b4941f8f89a22f75e885e662a"></a>

<a id="canonical-309b6b05540e2f6a1c7f07d8a650428d302620ffde5bffa6c13de43142ded02b"></a>

## description property — Property reference / 0166059ffa86 / 5

Type: `"string"`. Computed.

Description of the AppType.

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

- [features](data-sources--app_type--reference--group-001.md#canonical-cae6c27b5017fd80d1c34b09a8a775fc30f28f055f2fcf44507511f0cd059816): complete subsection reference.

<a id="canonical-f8f0d5696ac7592d575075483c4fc66beec8ca1ac1eb6a70a88c28531a827b76"></a>

<a id="canonical-36b8da76e29eb102f3a919920d65ae29b62806c52ad802f2109b8e68edea17d5"></a>

## id property — Property reference / 0166059ffa86 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-423de093f1333ca5556cfab659a850b7bc9899f53372effe6b4c89f3e3972706"></a>

<a id="canonical-55a29c6750ac320e4f3fdbecd468dcfc09a0171ca5c4507aab03709f0753df84"></a>

## labels property — Property reference / 0166059ffa86 / 7

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

<a id="canonical-7c60a5477cc12d8837333b6a6b50d5a1762234d182d34e6909add56622be6f72"></a>

<a id="canonical-066334f32f4bc359484e94ac35bda1d607d5e1002bb06333d4a1ceaf9f63f588"></a>

## name property — Property reference / 0166059ffa86 / 8

Type: `"string"`. Required.

Name of the AppType.

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

<a id="canonical-ebbaade649acaee471940e6ab53dda94ce52ce0061a79a09e9ac522a2d9706fa"></a>

<a id="canonical-b324f698486946e0c5f5f9f4ed500374af311fd09e6a2f1b836d3bc42fe9e563"></a>

## namespace property — Property reference / 0166059ffa86 / 9

Type: `"string"`. Required.

Namespace where the AppType exists.

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

<a id="canonical-10f96327b4667b2de16063c61acf59971fa3abc7e8df0f6a79ab6958e81b4e1b"></a>

## All schema paths — Property reference / 0166059ffa86 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--app_type--reference--group-001.md#canonical-9dc8b42f8743a589860cd232f603d1aee071cb04301a541ba8143a4ecde237ce) |
| `business_logic_markup_setting` | [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-b915d618d61b3478fd4e23c88724f00f669152b95c146018ccae1c832013bdde) |
| `business_logic_markup_setting.disable_spec` | [business_logic_markup_setting.disable_spec](data-sources--app_type--reference--group-001.md#canonical-01b51208d8b8512386b443645c2fdd6d56129bfed9dfee85e7c22871a8e04ac7) |
| `business_logic_markup_setting.discovered_api_settings` | [business_logic_markup_setting.discovered_api_settings](data-sources--app_type--reference--group-001.md#canonical-3f4e0e6b7e9f9e1b82aca12bc87c42bface9288ba9ac96583b03cb1bd7bfa9b0) |
| `business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis` | [business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis](data-sources--app_type--reference--group-001.md#canonical-7767ab4895876b420452c6bc772ae78d6c766c7bf1b9c591da23ee5e1b442b36) |
| `business_logic_markup_setting.enable` | [business_logic_markup_setting.enable](data-sources--app_type--reference--group-001.md#canonical-38472e96f2c568f851b71367c2819253f1fd8fb988f24c9db977b60ade41a79a) |
| `description` | [description](data-sources--app_type--reference--group-001.md#canonical-fd67a336864d3969419898df2d8e5127b85f690b4941f8f89a22f75e885e662a) |
| `features` | [features](data-sources--app_type--reference--group-001.md#canonical-2dbd0f604c0a1bf6e8ff6d97470c803d4d02b34ab54f59e6831abea819db842b) |
| `features.type` | [features.type](data-sources--app_type--reference--group-001.md#canonical-3e6d9084aba24a770a0572d9c41f965220153732da2f89f427d4ed9b30570921) |
| `id` | [id](data-sources--app_type--reference--group-001.md#canonical-f8f0d5696ac7592d575075483c4fc66beec8ca1ac1eb6a70a88c28531a827b76) |
| `labels` | [labels](data-sources--app_type--reference--group-001.md#canonical-423de093f1333ca5556cfab659a850b7bc9899f53372effe6b4c89f3e3972706) |
| `name` | [name](data-sources--app_type--reference--group-001.md#canonical-7c60a5477cc12d8837333b6a6b50d5a1762234d182d34e6909add56622be6f72) |
| `namespace` | [namespace](data-sources--app_type--reference--group-001.md#canonical-ebbaade649acaee471940e6ab53dda94ce52ce0061a79a09e9ac522a2d9706fa) |

<a id="canonical-0958c88849b22587f6a6f12d1878f9e45cfbed7fb2b3d19d6409fd5478f943b0"></a>

## Next pages — Property reference / 0166059ffa86 / 11

- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845)
- [features](data-sources--app_type--reference--group-001.md#canonical-cae6c27b5017fd80d1c34b09a8a775fc30f28f055f2fcf44507511f0cd059816)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)

<a id="canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfb19d131ecc5e19152f238d9fb2394e636a1a547f522a70328fc66bd8e91ee8"></a>

## business_logic_markup_setting — business_logic_markup_setting / c709145c05e1 / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- business_logic_markup_setting

<a id="canonical-b915d618d61b3478fd4e23c88724f00f669152b95c146018ccae1c832013bdde"></a>

Type: `"single"`. Computed.

Settings specifying how API Discovery will be performed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-3a99e0740bda15a3bf3b0c278c427423aa02abe63891573099666ce0a84f83a5"></a>

## Direct properties — business_logic_markup_setting / c709145c05e1 / 3

- [disable_spec](data-sources--app_type--reference--group-001.md#canonical-b14d0c4819ea38ce71359a1d3b6a693d5f7531b0567f2ee23983914b17c20317): complete subsection reference.

- [discovered_api_settings](data-sources--app_type--reference--group-001.md#canonical-8eaab9309f49aa5132fa7058004a022a86c4a55207d1844887652e37307c94a9): complete subsection reference.

- [enable](data-sources--app_type--reference--group-001.md#canonical-3dec424e3f7ec84636aa325ddbc595a994d9b4bd5e1bc3613a6318a3ec6965e1): complete subsection reference.

<a id="canonical-ad29429062268517ff7e1254043ef689efabd57752cfaac0aada99c7347d8ffc"></a>

## Next pages — business_logic_markup_setting / c709145c05e1 / 4

- [business_logic_markup_setting.disable_spec](data-sources--app_type--reference--group-001.md#canonical-b14d0c4819ea38ce71359a1d3b6a693d5f7531b0567f2ee23983914b17c20317)
- [business_logic_markup_setting.discovered_api_settings](data-sources--app_type--reference--group-001.md#canonical-8eaab9309f49aa5132fa7058004a022a86c4a55207d1844887652e37307c94a9)
- [business_logic_markup_setting.enable](data-sources--app_type--reference--group-001.md#canonical-3dec424e3f7ec84636aa325ddbc595a994d9b4bd5e1bc3613a6318a3ec6965e1)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)

<a id="canonical-b14d0c4819ea38ce71359a1d3b6a693d5f7531b0567f2ee23983914b17c20317"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd1967e8ee369a802f51d88d582a75329cadd6a6620f9ba01ef59507c27ffe21"></a>

## business_logic_markup_setting.disable_spec — business_logic_markup_setting.disable_spec / 917fb4a7398a / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845)
- business_logic_markup_setting.disable_spec

<a id="canonical-01b51208d8b8512386b443645c2fdd6d56129bfed9dfee85e7c22871a8e04ac7"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-af59fdb02541118fb83d44a3fb22cac366fd07a3d4e8b593af211da08d2fbe8d"></a>

## Direct properties — business_logic_markup_setting.disable_spec / 917fb4a7398a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d516c28dcc48284775e212e590aefbd832f223cf5075624b0d768dcf3bdd3f25"></a>

## Next pages — business_logic_markup_setting.disable_spec / 917fb4a7398a / 4

- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)

<a id="canonical-8eaab9309f49aa5132fa7058004a022a86c4a55207d1844887652e37307c94a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7d392ea1e62fbbf9e80c91bdc6b88071c9f534e20a6cf3ce6127e5e8dc137c0"></a>

## business_logic_markup_setting.discovered_api_settings — business_logic_markup_setting.discovered_api_settings / d9b79e8c2951 / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845)
- business_logic_markup_setting.discovered_api_settings

<a id="canonical-3f4e0e6b7e9f9e1b82aca12bc87c42bface9288ba9ac96583b03cb1bd7bfa9b0"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

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

<a id="canonical-a6b9f648f75cfb8e2f95551ca8bc34591b809b8010e792c2e1326f7c9d1651f6"></a>

## Direct properties — business_logic_markup_setting.discovered_api_settings / d9b79e8c2951 / 3

<a id="canonical-7767ab4895876b420452c6bc772ae78d6c766c7bf1b9c591da23ee5e1b442b36"></a>

<a id="canonical-1420baf56142a995b615930ccba641cf65bfcca6379559dc24452157955bc5ad"></a>

## purge_duration_for_inactive_discovered_apis property — business_logic_markup_setting.discovered_api_settings / d9b79e8c2951 / 4

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-a0c9340bac93bf91aa9b5255f061903689a1a0071868d77da15e0162c42c3291"></a>

## Next pages — business_logic_markup_setting.discovered_api_settings / d9b79e8c2951 / 5

- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)

<a id="canonical-3dec424e3f7ec84636aa325ddbc595a994d9b4bd5e1bc3613a6318a3ec6965e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c7982ddfea5076a842c6d44fbe0ec1b7b69e7b0a25f43cfd8f10218ef5ae789"></a>

## business_logic_markup_setting.enable — business_logic_markup_setting.enable / 68ce7be2c0df / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845)
- business_logic_markup_setting.enable

<a id="canonical-38472e96f2c568f851b71367c2819253f1fd8fb988f24c9db977b60ade41a79a"></a>

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

<a id="canonical-1e48858ef168759924fbc1370cf10173aab742cdc0c407b91f47dc267a4d4e49"></a>

## Direct properties — business_logic_markup_setting.enable / 68ce7be2c0df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-efe38dba6b4e3249835e1275750f241d9587346f95f14a82c0b258fbab237f43"></a>

## Next pages — business_logic_markup_setting.enable / 68ce7be2c0df / 4

- [business_logic_markup_setting](data-sources--app_type--reference--group-001.md#canonical-bb694ab46eb59683fa056548ff864f82039284d6652519d0377921e2f5d47845)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)

<a id="canonical-cae6c27b5017fd80d1c34b09a8a775fc30f28f055f2fcf44507511f0cd059816"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6ae694bcd7175e0cfeabdaad06d9e988702143f42a3f92de9e2c904bc52facd"></a>

## features — features / 1e1e199871d3 / 2

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
- [Property reference](data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- features

<a id="canonical-2dbd0f604c0a1bf6e8ff6d97470c803d4d02b34ab54f59e6831abea819db842b"></a>

Type: `"list"`. Computed.

List of various advanced security features enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e89cdbd471b84c37d23e01dd095e840e99965f9818b1abc95e7dffbd43a6c959"></a>

## Direct properties — features / 1e1e199871d3 / 3

<a id="canonical-3e6d9084aba24a770a0572d9c41f965220153732da2f89f427d4ed9b30570921"></a>

<a id="canonical-bc8982c5a282818d533404f41cd39578977c1d561e0931fc0dfb0c2e073c9b7e"></a>

## type property — features / 1e1e199871d3 / 4

Type: `"string"`. Computed.

\[Enum:
BUSINESS\_LOGIC\_MARKUP|TIMESERIES\_ANOMALY\_DETECTION|PER\_REQ\_ANOMALY\_DETECTION|USER\_BEHAVIOR\_ANALYSIS\]
Enumeration for advanced security features supported API Discovery enables generation of model for
various API interactions between services of App type. Enable analysis of timeseries for various
metric collected like requests, errors, latency etc. Enable anomaly detection per API request, i.e.
Possible values are \`BUSINESS\_LOGIC\_MARKUP\`, \`TIMESERIES\_ANOMALY\_DETECTION\`,
\`PER\_REQ\_ANOMALY\_DETECTION\`, \`USER\_BEHAVIOR\_ANALYSIS\`. Defaults to
\`BUSINESS\_LOGIC\_MARKUP\`.

Upstream description:

Enumeration for advanced security features supported

API Discovery enables generation of model for various API interactions between services of App type.
Enable analysis of timeseries for various metric collected like requests, errors, latency etc.
Enable anomaly detection per API request, i.e. The probability density function (PDF) charts
generation for API endpoints Enable user behavior analysis.

Receipt-pinned upstream constraints:

```json
{
  "default": "BUSINESS_LOGIC_MARKUP",
  "enum": [
    "BUSINESS_LOGIC_MARKUP",
    "TIMESERIES_ANOMALY_DETECTION",
    "PER_REQ_ANOMALY_DETECTION",
    "USER_BEHAVIOR_ANALYSIS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ef636753ddb0999187352f44cbf62e552c123235e5a1471fcc733ae29cc38f4f"></a>

## Next pages — features / 1e1e199871d3 / 5

- [Property reference](data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- [xcsh_app_type](../data-sources/app_type.md#canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c)
