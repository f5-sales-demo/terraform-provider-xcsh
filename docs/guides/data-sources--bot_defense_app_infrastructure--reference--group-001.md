---
page_title: "xcsh_bot_defense_app_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure reference."
---

# xcsh_bot_defense_app_infrastructure reference

<a id="canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ba1c9079345b3c89db6aae19b6adb55c3c460b239c8c37fa277519ec9a663a9"></a>

## Property reference — Property reference / f25028580ae2 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- Property reference

<a id="canonical-b66f1666cedfdff689856c20719b95910b77eca92ecc943844191529c43af347"></a>

## Direct properties — Property reference / f25028580ae2 / 3

<a id="canonical-ffd27212a099253229146f9b6f35847497b06dba4a34c41e13f0533a589721d9"></a>

<a id="canonical-b18267c572c4d9749b96a14e5016685f9d43a137660cdd173639a15526324aef"></a>

## annotations property — Property reference / f25028580ae2 / 4

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

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-53d1123273bfe5112de2ab7b2fa472073f689702764dbcc1c9aa5cadb7d26e93): complete subsection reference.

- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c8c7e430dfa42bfedd29a6c2cba0ca04db5392abd33cb2990253542e52b2a5c1): complete subsection reference.

<a id="canonical-bd82ed4c1470136e7384b007dd5c69b72283f5563f895fb0f8a4f1a0502c7d56"></a>

<a id="canonical-e5d1a49b253d54f3c9b691b67c8948ddaff379bffd6e1bcb659c54972a8b2183"></a>

## description property — Property reference / f25028580ae2 / 5

Type: `"string"`. Computed.

Description of the BotDefenseAppInfrastructure.

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

<a id="canonical-422f3703060fe9f441c8e6cddae3d97266dcdf572f968cea40baa08d52e3233b"></a>

<a id="canonical-dfbc0d81a9ca83d1b8d0f8a8850ab8bf96447b95995edec497909e1263387323"></a>

## environment_type property — Property reference / f25028580ae2 / 6

Type: `"string"`. Computed.

\[Enum: PRODUCTION|TESTING\] Environment Type Production environment Testing environment. Possible
values are \`PRODUCTION\`, \`TESTING\`. Defaults to \`PRODUCTION\`.

Upstream description:

Environment Type

Production environment Testing environment.

Receipt-pinned upstream constraints:

```json
{
  "default": "PRODUCTION",
  "enum": [
    "PRODUCTION",
    "TESTING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-084cb0745e3e4d803f314c36a3653f03934f1b299e80283dc2a380bbb8323846"></a>

<a id="canonical-7b5862db8eab6f0f63fc70682c536eb888ef943175127bddf3e7f4195c456fdf"></a>

## id property — Property reference / f25028580ae2 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-521068ce244e9a1807f63e490f5b2a05a8f58487849de8a74ea4f863fae9f35e"></a>

<a id="canonical-cb198970324d2453de0a9b58803759c12761ba5e8ccad75dc1104972fc09446a"></a>

## labels property — Property reference / f25028580ae2 / 8

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

<a id="canonical-44ba561d45b5b015df3c277d9dfe56e54412c994a85d7b4c63f1f6c6b037d967"></a>

<a id="canonical-2161eaabadc7bebca57c075ca103c1000ce8f1cd02765d4529ab722992946b7f"></a>

## name property — Property reference / f25028580ae2 / 9

Type: `"string"`. Required.

Name of the BotDefenseAppInfrastructure.

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

<a id="canonical-ff7a717ff4b8d12dfeb9149361ff1a69df0ab77751cddc241b7f60d83f0d4631"></a>

<a id="canonical-20cc8e010c9d58e7efc122db08911135e132bdcf0bea5084efc5c1cb7ab192d8"></a>

## namespace property — Property reference / f25028580ae2 / 10

Type: `"string"`. Required.

Namespace where the BotDefenseAppInfrastructure exists.

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

<a id="canonical-8d7d4f7ecdaff307eb7630635706c9d1189aba28f12576d945da27f7ca585b08"></a>

<a id="canonical-8e8fd838d05fe236508257fe831715c649e298cb1c6af44b36ca95dbda780b91"></a>

## traffic_type property — Property reference / f25028580ae2 / 11

Type: `"string"`. Computed.

\[Enum: WEB|MOBILE\] Traffic Type Web traffic Mobile traffic. Possible values are \`WEB\`,
\`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

Traffic Type

Web traffic Mobile traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "WEB",
  "enum": [
    "WEB",
    "MOBILE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d0881b505339700738c8bc880b199bcfc1fee776abd1bab8b256581374a09758"></a>

## All schema paths — Property reference / f25028580ae2 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ffd27212a099253229146f9b6f35847497b06dba4a34c41e13f0533a589721d9) |
| `cloud_hosted` | [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-cfa5167ab51567c5424fee5bda172ff7a1a316f510d0b9378d129f56dbdefbc0) |
| `cloud_hosted.egress` | [cloud_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-bcdfdf0fd2fc56027e8eae4e12e5543606cde4abedc917b88f18eb738ea29193) |
| `cloud_hosted.egress.ip_address` | [cloud_hosted.egress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0bcd587440c659c79bd839ecc95c3e7c8d3c0d5b6b6b9feb232e25f71bab7eca) |
| `cloud_hosted.egress.location` | [cloud_hosted.egress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-98ce21d78bd2f6cd5a6191e9f2d2c276804d3bdb5574a3ffe9ca76733ad9879b) |
| `cloud_hosted.infra_host_name` | [cloud_hosted.infra_host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ba6d23dab81988e9aeea091d87742f9b18796b5247517b73ecf7ac568922c073) |
| `cloud_hosted.ingress` | [cloud_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-8678ef86becde6e019c1e66b9371e487adb7db3fe9ce393d460da0ec6b201a96) |
| `cloud_hosted.ingress.host_name` | [cloud_hosted.ingress.host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-745535e9922a0972cb2faa9e1e7621f8d04c66fead3413a69c8413c7a741492f) |
| `cloud_hosted.ingress.ip_address` | [cloud_hosted.ingress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-205ffd37dc8546d9131c89281e3c03270db88b9e5921a495a038ce0f88879b80) |
| `cloud_hosted.ingress.location` | [cloud_hosted.ingress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-4ed2d7ef384a795d4003abf7b33ad89413acf85173e0902a12b80313501d12fc) |
| `cloud_hosted.region` | [cloud_hosted.region](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-6ea6eddfcdca9cf26eadd97b756faeb65603dce44c121114eaaa267f6c290a0f) |
| `data_center_hosted` | [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-145b898837b03238cadeb1802c63cb37fdff4185d0119b6dbbf33b39bd16e5d0) |
| `data_center_hosted.egress` | [data_center_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-64d3b2c6692feec4cc82a4dea624449b4291093a93e6bfd04033be7474a7ad92) |
| `data_center_hosted.egress.ip_address` | [data_center_hosted.egress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ea24c5b3939dcbc3000ea62d3b6ecd439fbce8e28cc10cc92376924f809e3883) |
| `data_center_hosted.egress.location` | [data_center_hosted.egress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-8bdf4670c6fd8ad4ed737e1a04de7d678072032be3cadc684ef459155bc5f942) |
| `data_center_hosted.infra_host_name` | [data_center_hosted.infra_host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-bfa32c64e68ccb7fc4edddd86cb46e94c7e33fdf36699dfe1876ba1944f21894) |
| `data_center_hosted.ingress` | [data_center_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c55b2924867ce96e3e0f29bb1fedc4d448e2dd196034e94466ee7064802f6e01) |
| `data_center_hosted.ingress.host_name` | [data_center_hosted.ingress.host_name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-129269a0e0bbc00fdacb1df3ceb739bac39aec4b5c7442aa978c84e8b8edb67e) |
| `data_center_hosted.ingress.ip_address` | [data_center_hosted.ingress.ip_address](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2f2affb4f8b9cc37ac51f0de74e876c9bb36731c053329f9ca8e24cfad4e2243) |
| `data_center_hosted.ingress.location` | [data_center_hosted.ingress.location](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0fe9feff2af52bda3f73f63d98497767f9b40cc38f12e420fa67532cfdda3d57) |
| `data_center_hosted.region` | [data_center_hosted.region](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-bbf49bf41a838375d5a55d5fa7b335d010ea16cbccc6af0372a1b121698a8544) |
| `description` | [description](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-bd82ed4c1470136e7384b007dd5c69b72283f5563f895fb0f8a4f1a0502c7d56) |
| `environment_type` | [environment_type](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-422f3703060fe9f441c8e6cddae3d97266dcdf572f968cea40baa08d52e3233b) |
| `id` | [id](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-084cb0745e3e4d803f314c36a3653f03934f1b299e80283dc2a380bbb8323846) |
| `labels` | [labels](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-521068ce244e9a1807f63e490f5b2a05a8f58487849de8a74ea4f863fae9f35e) |
| `name` | [name](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-44ba561d45b5b015df3c277d9dfe56e54412c994a85d7b4c63f1f6c6b037d967) |
| `namespace` | [namespace](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-ff7a717ff4b8d12dfeb9149361ff1a69df0ab77751cddc241b7f60d83f0d4631) |
| `traffic_type` | [traffic_type](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-8d7d4f7ecdaff307eb7630635706c9d1189aba28f12576d945da27f7ca585b08) |

<a id="canonical-5ca361d72bf3c05eec696af1078f183ab7aeaf97f2901182e09876ef14d18eec"></a>

## Next pages — Property reference / f25028580ae2 / 13

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-53d1123273bfe5112de2ab7b2fa472073f689702764dbcc1c9aa5cadb7d26e93)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c8c7e430dfa42bfedd29a6c2cba0ca04db5392abd33cb2990253542e52b2a5c1)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)

<a id="canonical-53d1123273bfe5112de2ab7b2fa472073f689702764dbcc1c9aa5cadb7d26e93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ec22a79232ca7867b849c13796d8e1c9826ea2732caa4fe2e9c89b64ebf5c3a"></a>

## cloud_hosted — cloud_hosted / 0b7a9565258a / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- cloud_hosted

<a id="canonical-cfa5167ab51567c5424fee5bda172ff7a1a316f510d0b9378d129f56dbdefbc0"></a>

Type: `"single"`. Computed.

\[OneOf: cloud\_hosted, data\_center\_hosted\] F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

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

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-cfa5167ab51567c5424fee5bda172ff7a1a316f510d0b9378d129f56dbdefbc0)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-145b898837b03238cadeb1802c63cb37fdff4185d0119b6dbbf33b39bd16e5d0)

Select alternatives according to the provider validators above.

<a id="canonical-28165d37a7872988dc0de4d2c33f2640e6c298681e2105e76f99d078e8126a74"></a>

## Direct properties — cloud_hosted / 0b7a9565258a / 3

- [egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0bdb5f9b2949a9cce096734b53554c5c5e06bd0a30e8753d690ca7014f0d6bc5): complete subsection reference.

<a id="canonical-ba6d23dab81988e9aeea091d87742f9b18796b5247517b73ecf7ac568922c073"></a>

<a id="canonical-5bf107c4b5d761d8b04c77339b61ad614a14e1e7e5b0b12d95294f5bbe9aabdc"></a>

## infra_host_name property — cloud_hosted / 0b7a9565258a / 4

Type: `"string"`. Computed.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-e67f939d52af905fe5b9401d0870c6301c475286634a3492596f48e581053c21): complete subsection reference.

<a id="canonical-6ea6eddfcdca9cf26eadd97b756faeb65603dce44c121114eaaa267f6c290a0f"></a>

<a id="canonical-c30be4b0cd036078ea16c318a8fac459818809b34e396b8f57c5509003c37a0d"></a>

## region property — cloud_hosted / 0b7a9565258a / 5

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bdd72b882dac4eed33e38580ee2119ab9e5df222ce793a748daae7444e891fe8"></a>

## Next pages — cloud_hosted / 0b7a9565258a / 6

- [cloud_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-0bdb5f9b2949a9cce096734b53554c5c5e06bd0a30e8753d690ca7014f0d6bc5)
- [cloud_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-e67f939d52af905fe5b9401d0870c6301c475286634a3492596f48e581053c21)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)

<a id="canonical-0bdb5f9b2949a9cce096734b53554c5c5e06bd0a30e8753d690ca7014f0d6bc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cef14909487864378d03f3776afa1a9afcf44a385ce37ab12fbe25d2ad012f2"></a>

## cloud_hosted.egress — cloud_hosted.egress / ccb44d6ed157 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-53d1123273bfe5112de2ab7b2fa472073f689702764dbcc1c9aa5cadb7d26e93)
- cloud_hosted.egress

<a id="canonical-bcdfdf0fd2fc56027e8eae4e12e5543606cde4abedc917b88f18eb738ea29193"></a>

Type: `"list"`. Computed.

Egress. Egress

Upstream description:

Egress

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-933e7fa37a987a9049da07164a7215b1255fe2d8c68d658de4085a7c45694cab"></a>

## Direct properties — cloud_hosted.egress / ccb44d6ed157 / 3

<a id="canonical-0bcd587440c659c79bd839ecc95c3e7c8d3c0d5b6b6b9feb232e25f71bab7eca"></a>

<a id="canonical-2c1c563d650157387b78d7372dc24fc88e46f971fba4301dc23e4839e59da945"></a>

## ip_address property — cloud_hosted.egress / ccb44d6ed157 / 4

Type: `"string"`. Computed.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-98ce21d78bd2f6cd5a6191e9f2d2c276804d3bdb5574a3ffe9ca76733ad9879b"></a>

<a id="canonical-1d9d4ab374413cd4e366aacc9920368f2cb9c347cf7effb651f6535126838160"></a>

## location property — cloud_hosted.egress / ccb44d6ed157 / 5

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-eacfc916bd2915cc23208812f9818dceb469ce3c86b03c653394115cb0a6b787"></a>

## Next pages — cloud_hosted.egress / ccb44d6ed157 / 6

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-53d1123273bfe5112de2ab7b2fa472073f689702764dbcc1c9aa5cadb7d26e93)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)

<a id="canonical-e67f939d52af905fe5b9401d0870c6301c475286634a3492596f48e581053c21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d94b11ab8e1d8530b3a45b0615ee9899e307961ba42693c512a81c039bae65e8"></a>

## cloud_hosted.ingress — cloud_hosted.ingress / 34fe9d9a4aec / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-53d1123273bfe5112de2ab7b2fa472073f689702764dbcc1c9aa5cadb7d26e93)
- cloud_hosted.ingress

<a id="canonical-8678ef86becde6e019c1e66b9371e487adb7db3fe9ce393d460da0ec6b201a96"></a>

Type: `"list"`. Computed.

Ingress. Ingress

Upstream description:

Ingress

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f12a4999d562247c1d52c78048100b3b141045582bf6c2264e9cb439d0d3b951"></a>

## Direct properties — cloud_hosted.ingress / 34fe9d9a4aec / 3

<a id="canonical-745535e9922a0972cb2faa9e1e7621f8d04c66fead3413a69c8413c7a741492f"></a>

<a id="canonical-f7766fcb1030b7e7a394f28e3a109776d27817492be9d5e3babdbde61905ddab"></a>

## host_name property — cloud_hosted.ingress / 34fe9d9a4aec / 4

Type: `"string"`. Computed.

Exclusive with \[ip\_address\] Ingress Host Name.

Upstream description:

Exclusive with \[ip\_address\] Ingress Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-205ffd37dc8546d9131c89281e3c03270db88b9e5921a495a038ce0f88879b80"></a>

<a id="canonical-9460626c29d9dcc4f984799ccd53c284eba7b6c44f6c2a583693bcef2a089e9d"></a>

## ip_address property — cloud_hosted.ingress / 34fe9d9a4aec / 5

Type: `"string"`. Computed.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-4ed2d7ef384a795d4003abf7b33ad89413acf85173e0902a12b80313501d12fc"></a>

<a id="canonical-97a7afe48553af5ee86dd47ed17c6290cf9bb046f84357a968a030ee5b3a16ee"></a>

## location property — cloud_hosted.ingress / 34fe9d9a4aec / 6

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8d2027f8afaef6e1f03431df02c6d00037e939360a4dd42106d5dbea3b38faae"></a>

## Next pages — cloud_hosted.ingress / 34fe9d9a4aec / 7

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-53d1123273bfe5112de2ab7b2fa472073f689702764dbcc1c9aa5cadb7d26e93)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)

<a id="canonical-c8c7e430dfa42bfedd29a6c2cba0ca04db5392abd33cb2990253542e52b2a5c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d084de11a1646a0ce84ef584aa27f350ae02ed699e764acb89a8565d56d17468"></a>

## data_center_hosted — data_center_hosted / 3909d7033c83 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- data_center_hosted

<a id="canonical-145b898837b03238cadeb1802c63cb37fdff4185d0119b6dbbf33b39bd16e5d0"></a>

Type: `"single"`. Computed.

F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

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

<a id="canonical-9c08bc5ae8a8d7e081b0ac9393d7055e615e405d1b644729abba2b0623fcb679"></a>

## Direct properties — data_center_hosted / 3909d7033c83 / 3

- [egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-85af8fec0c32977a69727935b8a7e57d491b636a24fcbc5369f6a142483cd100): complete subsection reference.

<a id="canonical-bfa32c64e68ccb7fc4edddd86cb46e94c7e33fdf36699dfe1876ba1944f21894"></a>

<a id="canonical-f9c3943dbd2e3fec72110fc3d7865714805b2c004097eb8912ed4db06e1c6e89"></a>

## infra_host_name property — data_center_hosted / 3909d7033c83 / 4

Type: `"string"`. Computed.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-be7a55db38263deef270cd96c21067d15e6cffb64a3f37028043f25d99564ab7): complete subsection reference.

<a id="canonical-bbf49bf41a838375d5a55d5fa7b335d010ea16cbccc6af0372a1b121698a8544"></a>

<a id="canonical-bd57961534143375e0553dcaae626765024ee34ebd6adcb73a229a0584f91250"></a>

## region property — data_center_hosted / 3909d7033c83 / 5

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5da6643928bbbe317e5b3a15e4bb80c74a8128a28ef11f51929a7a5598765fdd"></a>

## Next pages — data_center_hosted / 3909d7033c83 / 6

- [data_center_hosted.egress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-85af8fec0c32977a69727935b8a7e57d491b636a24fcbc5369f6a142483cd100)
- [data_center_hosted.ingress](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-be7a55db38263deef270cd96c21067d15e6cffb64a3f37028043f25d99564ab7)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)

<a id="canonical-85af8fec0c32977a69727935b8a7e57d491b636a24fcbc5369f6a142483cd100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9b290aa6ef695a4f725b8c32d7d302c69772205f12badbd90fa6fe58e6ed322"></a>

## data_center_hosted.egress — data_center_hosted.egress / 725e618871f3 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c8c7e430dfa42bfedd29a6c2cba0ca04db5392abd33cb2990253542e52b2a5c1)
- data_center_hosted.egress

<a id="canonical-64d3b2c6692feec4cc82a4dea624449b4291093a93e6bfd04033be7474a7ad92"></a>

Type: `"list"`. Computed.

Egress. Egress

Upstream description:

Egress

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-334a6342cc6260797d2fe4bc99861122e96f85a65d49266dd03f5d5788a3c4b3"></a>

## Direct properties — data_center_hosted.egress / 725e618871f3 / 3

<a id="canonical-ea24c5b3939dcbc3000ea62d3b6ecd439fbce8e28cc10cc92376924f809e3883"></a>

<a id="canonical-f45c9dc636d11709e8c88bc4cee4b012ed1330122bd59160cad7218960990ea9"></a>

## ip_address property — data_center_hosted.egress / 725e618871f3 / 4

Type: `"string"`. Computed.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-8bdf4670c6fd8ad4ed737e1a04de7d678072032be3cadc684ef459155bc5f942"></a>

<a id="canonical-1c66af2e8c3b3cf51b797724baf3a3245d8b6531fa8dcd22a67fd3d1f08b35e5"></a>

## location property — data_center_hosted.egress / 725e618871f3 / 5

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8443cc1272b0c13990170ec1f177521aa7817544a5caf917bb936cc481575436"></a>

## Next pages — data_center_hosted.egress / 725e618871f3 / 6

- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c8c7e430dfa42bfedd29a6c2cba0ca04db5392abd33cb2990253542e52b2a5c1)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)

<a id="canonical-be7a55db38263deef270cd96c21067d15e6cffb64a3f37028043f25d99564ab7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91d2541ef2dc50e313d3c0f4e4043f114069afc9435f30d83f35728b219b1638"></a>

## data_center_hosted.ingress — data_center_hosted.ingress / 103611ed3a9f / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-9ed1ea7dcbecfc8b88df10bacaa7a6abc8c9e7cef807d0f3b9cbbe1690640bda)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c8c7e430dfa42bfedd29a6c2cba0ca04db5392abd33cb2990253542e52b2a5c1)
- data_center_hosted.ingress

<a id="canonical-c55b2924867ce96e3e0f29bb1fedc4d448e2dd196034e94466ee7064802f6e01"></a>

Type: `"list"`. Computed.

Ingress. Ingress

Upstream description:

Ingress

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8a134954b8ccd6d4a774be3d80497150e8470817c2a942c42a188a8a5265625f"></a>

## Direct properties — data_center_hosted.ingress / 103611ed3a9f / 3

<a id="canonical-129269a0e0bbc00fdacb1df3ceb739bac39aec4b5c7442aa978c84e8b8edb67e"></a>

<a id="canonical-0c9c822565f314f3824edd5b70e7a0fb7e2528c5c1e28e7757fbdee56c31b5bc"></a>

## host_name property — data_center_hosted.ingress / 103611ed3a9f / 4

Type: `"string"`. Computed.

Exclusive with \[ip\_address\] Ingress Host Name.

Upstream description:

Exclusive with \[ip\_address\] Ingress Host Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2f2affb4f8b9cc37ac51f0de74e876c9bb36731c053329f9ca8e24cfad4e2243"></a>

<a id="canonical-f5560438c33295d59b85cf6f95fa62918a618253c25155132cdb01a3e125b877"></a>

## ip_address property — data_center_hosted.ingress / 103611ed3a9f / 5

Type: `"string"`. Computed.

Exclusive with \[host\_name\] Ingress IP Address.

Upstream description:

Exclusive with \[host\_name\] Ingress IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0fe9feff2af52bda3f73f63d98497767f9b40cc38f12e420fa67532cfdda3d57"></a>

<a id="canonical-58558a44b73e8d968d15bbb929cc47481bfab15c720a3e641ce8af175a850a2f"></a>

## location property — data_center_hosted.ingress / 103611ed3a9f / 6

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Upstream description:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-20f2dd8feda3514821a848b6b9b5ea1e98bd17f9ce695b53d42945d9be7d37d6"></a>

## Next pages — data_center_hosted.ingress / 103611ed3a9f / 7

- [data_center_hosted](data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-c8c7e430dfa42bfedd29a6c2cba0ca04db5392abd33cb2990253542e52b2a5c1)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-f36ab4b52244feda08ea56f1901fc5744d9285c8e182df39ed067d65a8637425)
