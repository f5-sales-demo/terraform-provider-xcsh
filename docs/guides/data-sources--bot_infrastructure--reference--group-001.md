---
page_title: "xcsh_bot_infrastructure reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure reference."
---

# xcsh_bot_infrastructure reference

<a id="canonical-e3ce34ec019c1f41a1f0eae669da3664626b6c7209669ae24fa287835acfefee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55bd18c3ecc7853a8f736c5d7148d8bf3c9d1de1ec270cd61aa06fb43d0f144b"></a>

## Property reference — Property reference / 61dd550e09fd / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
- Property reference

<a id="canonical-1071651c19e3e9123696d5ca00bc34c52f3070e8035bf63daed57c34bfac8573"></a>

## Direct properties — Property reference / 61dd550e09fd / 3

<a id="canonical-0d87c9f5951f1fbc80ab28bc48763a27dddb1af980c77a0abddaf7cfe7f594c6"></a>

<a id="canonical-1cdd1a907bc2aee818867301f93e066d826ac049655f45043bf03c62c454b583"></a>

## annotations property — Property reference / 61dd550e09fd / 4

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

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-42a6a07091a3a1551f90bee9ebfad0c1cf0e9b122612821a38744c0af4c83859): complete subsection reference.

<a id="canonical-62dfd0ae5759890aeb0384a058f51110f98991c3bb6c749fad6bfb7996b85383"></a>

<a id="canonical-5ff6f8f637f8414700ab37f727f89039b9fb0220b2ac66c59d68dea4ffb11d1a"></a>

## description property — Property reference / 61dd550e09fd / 5

Type: `"string"`. Computed.

Description of the BotInfrastructure.

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

<a id="canonical-bbc1ad6a5a7978b165a5b3644bf387f5d8f3d2cdd02f56c412076ffcb71d39d3"></a>

<a id="canonical-d6ec138849eb81763a58810449b147b8072ff8e085f7b4a04d0ded1313a4a9df"></a>

## id property — Property reference / 61dd550e09fd / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-124cbbf6faa8e2d3eb128553b991bfef7631b6c9a2e8bbd70ae5230804c75a44"></a>

<a id="canonical-afac7598c704582f9ea1685100c6b72af8f6e3d3cec24aee02a42b20c380eeec"></a>

## labels property — Property reference / 61dd550e09fd / 7

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

<a id="canonical-83a1660b2644a001805c619ee27dea8707ac7f1f26205cd2e65df10a17d302c3"></a>

<a id="canonical-14d257fd9c394ca67896ea3b065ff7975c8a1102a0b1a5c308d9b40044dd1949"></a>

## name property — Property reference / 61dd550e09fd / 8

Type: `"string"`. Required.

Name of the BotInfrastructure.

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

<a id="canonical-c6cf03be6dcc6633c791f47e142c831128195365f1b6f8031136cd2cd9986a07"></a>

<a id="canonical-54b0241b00c4d8e2a4d39a4fd7beae69e5878e617aa2a98138189e08ec84d2e7"></a>

## namespace property — Property reference / 61dd550e09fd / 9

Type: `"string"`. Required.

Namespace where the BotInfrastructure exists.

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

<a id="canonical-21c353327f6c67384ff18af7e0744d17c02ec201e3faefec7f42944bceeef854"></a>

<a id="canonical-40d545643c306884e51da43a0218131775e36192944ebbf97aff13b85fda053a"></a>

## traffic_type property — Property reference / 61dd550e09fd / 10

Type: `"string"`. Computed.

\[Enum: WEB|MOBILE\] The type of traffic that is routed to and processed by this infrastructure (Web
or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through
this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense
SDK are routed.. Possible values are \`WEB\`, \`MOBILE\`. Defaults to \`WEB\`.

Upstream description:

The type of traffic that is routed to and processed by this infrastructure (Web or Mobile).

Only web traffic, including browser-based traffic from mobile devices, is routed through this Bot
Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense SDK are
routed through this Bot Defense infrastructure.

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

<a id="canonical-50b5c1fc158d05ed500ce1cea301936cf8c3a5c2a90c0ef8225b8ef98149350f"></a>

## All schema paths — Property reference / 61dd550e09fd / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bot_infrastructure--reference--group-001.md#canonical-0d87c9f5951f1fbc80ab28bc48763a27dddb1af980c77a0abddaf7cfe7f594c6) |
| `create_cloud_hosted` | [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-219b6a4ceae65a60930e39ed6bc882f7ba591826b9b4395f84b00e50d79b7b1f) |
| `create_cloud_hosted.ip_addresses` | [create_cloud_hosted.ip_addresses](data-sources--bot_infrastructure--reference--group-001.md#canonical-ad7517bd8fd1e63b2356934ef2d4308dfb62a1e924a7904b0be9df4914275c61) |
| `create_cloud_hosted.production` | [create_cloud_hosted.production](data-sources--bot_infrastructure--reference--group-001.md#canonical-92f7d3cd495a8631bffd6532d7f51397c54f66273bec6694f25cd88417f920b6) |
| `create_cloud_hosted.production.region_1` | [create_cloud_hosted.production.region_1](data-sources--bot_infrastructure--reference--group-001.md#canonical-7e443cfc26b46b278d67c95773eb11bf4caec77a2cc70898121da4bca2b83ed0) |
| `create_cloud_hosted.production.region_2` | [create_cloud_hosted.production.region_2](data-sources--bot_infrastructure--reference--group-001.md#canonical-2e296ce9307a55068c190038f116e2cbac216df88f134b545649d89b1191ce6e) |
| `create_cloud_hosted.testing` | [create_cloud_hosted.testing](data-sources--bot_infrastructure--reference--group-001.md#canonical-294ec0c659d51b638eb60aaadbd854663be508b101ea97793f5fea2f3c148fb0) |
| `create_cloud_hosted.testing.region_1` | [create_cloud_hosted.testing.region_1](data-sources--bot_infrastructure--reference--group-001.md#canonical-2e4e4f6cac3deb3e974186a37629ce504626336f50c44bc8bb77d206851a258d) |
| `description` | [description](data-sources--bot_infrastructure--reference--group-001.md#canonical-62dfd0ae5759890aeb0384a058f51110f98991c3bb6c749fad6bfb7996b85383) |
| `id` | [id](data-sources--bot_infrastructure--reference--group-001.md#canonical-bbc1ad6a5a7978b165a5b3644bf387f5d8f3d2cdd02f56c412076ffcb71d39d3) |
| `labels` | [labels](data-sources--bot_infrastructure--reference--group-001.md#canonical-124cbbf6faa8e2d3eb128553b991bfef7631b6c9a2e8bbd70ae5230804c75a44) |
| `name` | [name](data-sources--bot_infrastructure--reference--group-001.md#canonical-83a1660b2644a001805c619ee27dea8707ac7f1f26205cd2e65df10a17d302c3) |
| `namespace` | [namespace](data-sources--bot_infrastructure--reference--group-001.md#canonical-c6cf03be6dcc6633c791f47e142c831128195365f1b6f8031136cd2cd9986a07) |
| `traffic_type` | [traffic_type](data-sources--bot_infrastructure--reference--group-001.md#canonical-21c353327f6c67384ff18af7e0744d17c02ec201e3faefec7f42944bceeef854) |

<a id="canonical-e7b1e3164084a0e066d79d2d8b6b0f92c725228e42fd4926c104605cb29e68dd"></a>

## Next pages — Property reference / 61dd550e09fd / 12

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-42a6a07091a3a1551f90bee9ebfad0c1cf0e9b122612821a38744c0af4c83859)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)

<a id="canonical-42a6a07091a3a1551f90bee9ebfad0c1cf0e9b122612821a38744c0af4c83859"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-836fa5e6437eeae7e53c80993c33d3d24c4d50eed02aeb67131e12de2ee9b32c"></a>

## create_cloud_hosted — create_cloud_hosted / d17d7e23e2bf / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-e3ce34ec019c1f41a1f0eae669da3664626b6c7209669ae24fa287835acfefee)
- create_cloud_hosted

<a id="canonical-219b6a4ceae65a60930e39ed6bc882f7ba591826b9b4395f84b00e50d79b7b1f"></a>

Type: `"single"`. Computed.

F5 Cloud Hosted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_choice": "[\"production\",\"testing\"]"
}
```

<a id="canonical-04854ee103ada719c6c7eceb68e851aad9b904cba7203f9de43d3d29479ac0d2"></a>

## Direct properties — create_cloud_hosted / d17d7e23e2bf / 3

<a id="canonical-ad7517bd8fd1e63b2356934ef2d4308dfb62a1e924a7904b0be9df4914275c61"></a>

<a id="canonical-f8962195f05fdad65e8607229363eed1f867ce909304b515092577bfca9a3463"></a>

## ip_addresses property — create_cloud_hosted / d17d7e23e2bf / 4

Type: `["list", "string"]`. Computed.

Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [production](data-sources--bot_infrastructure--reference--group-001.md#canonical-d35c64a6542461bbb575b00a0d437294a7cc69f4910ad7c5880bf1813f0e651d): complete subsection reference.

- [testing](data-sources--bot_infrastructure--reference--group-001.md#canonical-27d7594f1809924f458da66029ebc370d680a7094060f50b28910f0838b73e32): complete subsection reference.

<a id="canonical-5981ab55d27b28a685b3df223043fe66f0d0554ab4ee1063c10eb069a505ec27"></a>

## Next pages — create_cloud_hosted / d17d7e23e2bf / 5

- [create_cloud_hosted.production](data-sources--bot_infrastructure--reference--group-001.md#canonical-d35c64a6542461bbb575b00a0d437294a7cc69f4910ad7c5880bf1813f0e651d)
- [create_cloud_hosted.testing](data-sources--bot_infrastructure--reference--group-001.md#canonical-27d7594f1809924f458da66029ebc370d680a7094060f50b28910f0838b73e32)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-e3ce34ec019c1f41a1f0eae669da3664626b6c7209669ae24fa287835acfefee)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)

<a id="canonical-d35c64a6542461bbb575b00a0d437294a7cc69f4910ad7c5880bf1813f0e651d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2baed5e015834485ff485943da580c7c772105eb1fa92c4e97dfb3164becfbbe"></a>

## create_cloud_hosted.production — create_cloud_hosted.production / 177347938b24 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-e3ce34ec019c1f41a1f0eae669da3664626b6c7209669ae24fa287835acfefee)
- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-42a6a07091a3a1551f90bee9ebfad0c1cf0e9b122612821a38744c0af4c83859)
- create_cloud_hosted.production

<a id="canonical-92f7d3cd495a8631bffd6532d7f51397c54f66273bec6694f25cd88417f920b6"></a>

Type: `"single"`. Computed.

Production.

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

<a id="canonical-34874451fb800cdef5dde6e434e91c53f380f0178b82f1e3ddd2c3b3de54aee7"></a>

## Direct properties — create_cloud_hosted.production / 177347938b24 / 3

<a id="canonical-7e443cfc26b46b278d67c95773eb11bf4caec77a2cc70898121da4bca2b83ed0"></a>

<a id="canonical-b66ffb04f67eaa8d10093e799d557c78cdb27720d00d824a6780b433b99c7a9c"></a>

## region_1 property — create_cloud_hosted.production / 177347938b24 / 4

Type: `"string"`. Computed.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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

<a id="canonical-2e296ce9307a55068c190038f116e2cbac216df88f134b545649d89b1191ce6e"></a>

<a id="canonical-f988dc471f0e10b574ea43dcfe449dfde160ee5428550668056475386d371c00"></a>

## region_2 property — create_cloud_hosted.production / 177347938b24 / 5

Type: `"string"`. Computed.

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

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

<a id="canonical-06c8d1b7edb9e47414383977a7639816e5e3834092e7f7cb139ced71c6e97a4d"></a>

## Next pages — create_cloud_hosted.production / 177347938b24 / 6

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-42a6a07091a3a1551f90bee9ebfad0c1cf0e9b122612821a38744c0af4c83859)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)

<a id="canonical-27d7594f1809924f458da66029ebc370d680a7094060f50b28910f0838b73e32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf7f48ac1de7f4a21771910d1c616fb311869bd6f72a39cd6b9f46b0b523c3f9"></a>

## create_cloud_hosted.testing — create_cloud_hosted.testing / 95d6b84b315a / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
- [Property reference](data-sources--bot_infrastructure--reference--group-001.md#canonical-e3ce34ec019c1f41a1f0eae669da3664626b6c7209669ae24fa287835acfefee)
- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-42a6a07091a3a1551f90bee9ebfad0c1cf0e9b122612821a38744c0af4c83859)
- create_cloud_hosted.testing

<a id="canonical-294ec0c659d51b638eb60aaadbd854663be508b101ea97793f5fea2f3c148fb0"></a>

Type: `"single"`. Computed.

Testing

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

<a id="canonical-8843e31457bdb5e37729f544468945633fcb3a682e3d32f2d551ba0971aa5c3a"></a>

## Direct properties — create_cloud_hosted.testing / 95d6b84b315a / 3

<a id="canonical-2e4e4f6cac3deb3e974186a37629ce504626336f50c44bc8bb77d206851a258d"></a>

<a id="canonical-e7e35373ba28b04db405ad4d1212fef2a59eaa38b0f81cfaccdc337bdc1557b1"></a>

## region_1 property — create_cloud_hosted.testing / 95d6b84b315a / 4

Type: `"string"`. Computed.

Active-Passive Infrastructure configuration where traffic is routed to a single region.

Upstream description:

This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.

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

<a id="canonical-0a747fe7d668d2188697576038a0ed326acac8764acc0854edd3917282ca5f83"></a>

## Next pages — create_cloud_hosted.testing / 95d6b84b315a / 5

- [create_cloud_hosted](data-sources--bot_infrastructure--reference--group-001.md#canonical-42a6a07091a3a1551f90bee9ebfad0c1cf0e9b122612821a38744c0af4c83859)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
