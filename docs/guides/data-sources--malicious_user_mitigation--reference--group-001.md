---
page_title: "xcsh_malicious_user_mitigation reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation reference."
---

# xcsh_malicious_user_mitigation reference

<a id="canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9bdf0a0ab291b9c60df931ced2d03f7d4b700689a93c923e051016115856478"></a>

## Property reference — Property reference / e5b4f5e531f7 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- Property reference

<a id="canonical-437c708a1f995aed0f70a3ec72d12e0c2f26bbc463b633e26d0e96825b71df0a"></a>

## Direct properties — Property reference / e5b4f5e531f7 / 3

<a id="canonical-b1b03d9deaac1c09c50492591ef3f0da44a39dd7c4ba3b83c589a305a878eee1"></a>

<a id="canonical-bb58ecf7d01a1db548dcbf9fc2370444fb40e206921b20d7f9061c6cdbca5941"></a>

## annotations property — Property reference / e5b4f5e531f7 / 4

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

<a id="canonical-25031100934239a985409dd5f4dd15f936581f8380e3bccbf02fc313989b6a9d"></a>

<a id="canonical-337a40da20f0807624fb39c98dc6a3ed178cd1e0f2037dd9f2210141fb8c7395"></a>

## description property — Property reference / e5b4f5e531f7 / 5

Type: `"string"`. Computed.

Description of the MaliciousUserMitigation.

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

<a id="canonical-adb322cd3cb6b936f082147924b740c28876c07ae02f304c16e25eb325eb0aa7"></a>

<a id="canonical-018eb167abef30af157f708536a70f36e5724d3af5f56a0c2e542e913c92df9b"></a>

## id property — Property reference / e5b4f5e531f7 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-be054063594008e9d3135b1959a125fe65b1ce342e0ce553d4e932b3f70ea97b"></a>

<a id="canonical-c4a94722cdfcd651bc67828cb9848d842c85fa1384fb61eb4847740cd6998144"></a>

## labels property — Property reference / e5b4f5e531f7 / 7

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

- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465): complete subsection reference.

<a id="canonical-fd54dafad4233a2b85f159a9b10164dc9ffc4bd5b6dfefde7d13e2163ae8ce72"></a>

<a id="canonical-faf3e56045298c1d5584e0b39a09cdeb1e002f92c62a6ead014fe5e40ed7acd2"></a>

## name property — Property reference / e5b4f5e531f7 / 8

Type: `"string"`. Required.

Name of the MaliciousUserMitigation.

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

<a id="canonical-887a143a955026ffc52d7a051b99955db052085f779d606230844016b2f65144"></a>

<a id="canonical-0765d37537b8063b743e5a17cbf50be2a0029d4922b5c5d755855c95117768b6"></a>

## namespace property — Property reference / e5b4f5e531f7 / 9

Type: `"string"`. Required.

Namespace where the MaliciousUserMitigation exists.

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

<a id="canonical-2fa85b85ab41dda7ddb0ea29fa739852bc052b82c9a9ea4ac3ca660308ccaefe"></a>

## All schema paths — Property reference / e5b4f5e531f7 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-b1b03d9deaac1c09c50492591ef3f0da44a39dd7c4ba3b83c589a305a878eee1) |
| `description` | [description](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-25031100934239a985409dd5f4dd15f936581f8380e3bccbf02fc313989b6a9d) |
| `id` | [id](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-adb322cd3cb6b936f082147924b740c28876c07ae02f304c16e25eb325eb0aa7) |
| `labels` | [labels](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-be054063594008e9d3135b1959a125fe65b1ce342e0ce553d4e932b3f70ea97b) |
| `mitigation_type` | [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-83205429f107e631449a53bcca75436567995a4fe8af9e90e1416ed6aba95b72) |
| `mitigation_type.rules` | [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-8c30c8842a13a388b52d16c8e68450e03a1907745cf9989e800786640fa68997) |
| `mitigation_type.rules.mitigation_action` | [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-7cb20b02932cea086a9bdb88a803b86e0e30c37d63d277be9dbc60a5d5f2e5f9) |
| `mitigation_type.rules.mitigation_action.block_temporarily` | [mitigation_type.rules.mitigation_action.block_temporarily](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-25abb4e9cb52997cb07010206fdb204d7fbdd9768366528358f7a7e1ac2b3074) |
| `mitigation_type.rules.mitigation_action.captcha_challenge` | [mitigation_type.rules.mitigation_action.captcha_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-9a71554985ba15f286a7d0610b15a0d231ff7dc4dd0a91fc827cc00c39e32338) |
| `mitigation_type.rules.mitigation_action.javascript_challenge` | [mitigation_type.rules.mitigation_action.javascript_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-1fda519edff2b5d89b9cc261f3bc584fecc635473907cdf537c1b82144b1084f) |
| `mitigation_type.rules.threat_level` | [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e054a9a7754897b6ab98e817a44577b1cfd6e22b1a368ace46c4eb414889d8a9) |
| `mitigation_type.rules.threat_level.high` | [mitigation_type.rules.threat_level.high](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-1b52378241bd44ea1a07beeb3481152b66e1fb1544ba18d76de17434add802b9) |
| `mitigation_type.rules.threat_level.low` | [mitigation_type.rules.threat_level.low](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-16b56b5280de9a95b3dff35620363addb8b462985cbf7f37e8c50c7df0347827) |
| `mitigation_type.rules.threat_level.medium` | [mitigation_type.rules.threat_level.medium](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-790c260771846d4cf0b27b3160f1d0cc7f8551737404567b160ad146e72a791f) |
| `name` | [name](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-fd54dafad4233a2b85f159a9b10164dc9ffc4bd5b6dfefde7d13e2163ae8ce72) |
| `namespace` | [namespace](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-887a143a955026ffc52d7a051b99955db052085f779d606230844016b2f65144) |

<a id="canonical-234ff0f3acee2200d25aefe9b29ca46aee96344443723c55848c93e8d9b1968d"></a>

## Next pages — Property reference / e5b4f5e531f7 / 11

- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6f837a85fa17921b8976076923fef1e584adaa4032aa3bd6a4ba63821e5709a"></a>

## mitigation_type — mitigation_type / c2aaed7f314a / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- mitigation_type

<a id="canonical-83205429f107e631449a53bcca75436567995a4fe8af9e90e1416ed6aba95b72"></a>

Type: `"single"`. Computed.

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. Server applies default when omitted.

Upstream description:

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. The settings defined in malicious user
mitigation specify what mitigation actions to take for user determined to be at different threat
levels.

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

<a id="canonical-f1f1da093d0632a1146709f3b596995ce815568d0dc4efacf28a36e1b0045764"></a>

## Direct properties — mitigation_type / c2aaed7f314a / 3

- [rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9): complete subsection reference.

<a id="canonical-99539c21af63d6e239763c2ac44ef1a5fe9fd045ee27daa09ced333e1a0c9d7f"></a>

## Next pages — mitigation_type / c2aaed7f314a / 4

- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ccb4e375bf54be13df6717ae245df7f78230c93a4c10d39a7ccdbacee3036a2"></a>

## mitigation_type.rules — mitigation_type.rules / 0230f7eb6601 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- mitigation_type.rules

<a id="canonical-8c30c8842a13a388b52d16c8e68450e03a1907745cf9989e800786640fa68997"></a>

Type: `"list"`. Computed.

Define the threat levels and the corresponding mitigation actions to be taken.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
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
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  }
}
```

<a id="canonical-a8e1d6697b41c5570dad60c6523ab70c4a4bbd448f194435c4e7de11009ef30b"></a>

## Direct properties — mitigation_type.rules / 0230f7eb6601 / 3

- [mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13): complete subsection reference.

- [threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a): complete subsection reference.

<a id="canonical-37cbc7757e74f758206775ee2674f51c6de75c92525bf48882415b3a479dd2ee"></a>

## Next pages — mitigation_type.rules / 0230f7eb6601 / 4

- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-adb0b68e8aa741e2aeb244a90b90acd6ddbbff76c3d3d4d7f8145224651d1286"></a>

## mitigation_type.rules.mitigation_action — mitigation_type.rules.mitigation_action / 14e1bba3bd48 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- mitigation_type.rules.mitigation_action

<a id="canonical-7cb20b02932cea086a9bdb88a803b86e0e30c37d63d277be9dbc60a5d5f2e5f9"></a>

Type: `"single"`. Computed.

Supported actions that can be taken to mitigate malicious activity from a user.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation_action": "[\"block_temporarily\",\"captcha_challenge\",\"javascript_challenge\"]"
}
```

<a id="canonical-5cca6376dad9886f34b847cecfc2f05050206718eb5a1a3db4197960d9b6de7b"></a>

## Direct properties — mitigation_type.rules.mitigation_action / 14e1bba3bd48 / 3

- [block_temporarily](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-f0df6aa413c0146809cbcff68ab0b93c2aa50cfcd93c3fe556c132ac3a4ede9d): complete subsection reference.

- [captcha_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-f5f5c709206eac82dcf7d297b75f379f0eb3ac897a255de2e200590d6eb25d6d): complete subsection reference.

- [javascript_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-79f1bfa4a9e07227e01c98b1b16107d7a2f8d3e62255151ab537097773af8385): complete subsection reference.

<a id="canonical-e0d8fc7dedf5f35cf52fee0bbcadadda3815d8f8a3c11ff9e6c42daca2845f1e"></a>

## Next pages — mitigation_type.rules.mitigation_action / 14e1bba3bd48 / 4

- [mitigation_type.rules.mitigation_action.block_temporarily](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-f0df6aa413c0146809cbcff68ab0b93c2aa50cfcd93c3fe556c132ac3a4ede9d)
- [mitigation_type.rules.mitigation_action.captcha_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-f5f5c709206eac82dcf7d297b75f379f0eb3ac897a255de2e200590d6eb25d6d)
- [mitigation_type.rules.mitigation_action.javascript_challenge](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-79f1bfa4a9e07227e01c98b1b16107d7a2f8d3e62255151ab537097773af8385)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-f0df6aa413c0146809cbcff68ab0b93c2aa50cfcd93c3fe556c132ac3a4ede9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b3afe3e451089b9db36066cabd686afa8aa8a5f56f4bee602dccc8710d663bd"></a>

## mitigation_type.rules.mitigation_action.block_temporarily — mitigation_type.rules.mitigation_action.block_temporarily / 8714585f18fe / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13)
- mitigation_type.rules.mitigation_action.block_temporarily

<a id="canonical-25abb4e9cb52997cb07010206fdb204d7fbdd9768366528358f7a7e1ac2b3074"></a>

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

<a id="canonical-1392c9c0ca33b72fa0c15809285cc55626de53f9a75f2c44ae5f78d49def75e5"></a>

## Direct properties — mitigation_type.rules.mitigation_action.block_temporarily / 8714585f18fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0709ec5a271f5ce11c5adbbc02ca969a692d18793ad310a9486aae3e94eb141"></a>

## Next pages — mitigation_type.rules.mitigation_action.block_temporarily / 8714585f18fe / 4

- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-f5f5c709206eac82dcf7d297b75f379f0eb3ac897a255de2e200590d6eb25d6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f782d81c943bc342efae8211cf50fb13c3935d0432bfd63135ee5273f1469121"></a>

## mitigation_type.rules.mitigation_action.captcha_challenge — mitigation_type.rules.mitigation_action.captcha_challenge / 5639fe274ac8 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13)
- mitigation_type.rules.mitigation_action.captcha_challenge

<a id="canonical-9a71554985ba15f286a7d0610b15a0d231ff7dc4dd0a91fc827cc00c39e32338"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for captcha challenge.

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

<a id="canonical-31906fc360a63114e80d09988d720a88f8f5a27329dbb4c5d918a7a0e3ce1c5c"></a>

## Direct properties — mitigation_type.rules.mitigation_action.captcha_challenge / 5639fe274ac8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-488dc5e99bcce2337f57b3553255d8ea313a1fecf002734d0a6e1a299bae125b"></a>

## Next pages — mitigation_type.rules.mitigation_action.captcha_challenge / 5639fe274ac8 / 4

- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-79f1bfa4a9e07227e01c98b1b16107d7a2f8d3e62255151ab537097773af8385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3761e392ff85f718dd0163421b9478ee4aeec0c0d4a062aac4917c10e88052d"></a>

## mitigation_type.rules.mitigation_action.javascript_challenge — mitigation_type.rules.mitigation_action.javascript_challenge / ecbff166d64e / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13)
- mitigation_type.rules.mitigation_action.javascript_challenge

<a id="canonical-1fda519edff2b5d89b9cc261f3bc584fecc635473907cdf537c1b82144b1084f"></a>

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

<a id="canonical-e12726663af01fa0b951945790cac0eac91eac475e9b054951acb6f5f820d761"></a>

## Direct properties — mitigation_type.rules.mitigation_action.javascript_challenge / ecbff166d64e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26f611ea4e21e1d8bff5ebd2e16bcb9d738d699200b5797cf3708cacce4b186c"></a>

## Next pages — mitigation_type.rules.mitigation_action.javascript_challenge / ecbff166d64e / 4

- [mitigation_type.rules.mitigation_action](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0c333210358d0668fda2ff2911c1f0bf59b14a5d71a13014a4b1037160782e13)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b54030b9833f615d7b2b74341ccd0c801dc3d8ff2373539ab0b9a24ca8af03bf"></a>

## mitigation_type.rules.threat_level — mitigation_type.rules.threat_level / 6196c9ba6387 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- mitigation_type.rules.threat_level

<a id="canonical-e054a9a7754897b6ab98e817a44577b1cfd6e22b1a368ace46c4eb414889d8a9"></a>

Type: `"single"`. Computed.

Threat level estimated for each user based on the user's activity and reputation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-threat_level": "[\"high\",\"low\",\"medium\"]"
}
```

<a id="canonical-b5fad76b30d0ef53112f39e8062c6b83f87c76631a1266d1408a84f309e0cd1c"></a>

## Direct properties — mitigation_type.rules.threat_level / 6196c9ba6387 / 3

- [high](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-5f5a849b32ca89d64c6cba83e34a8cb36c8d2b52b25120763fdbf0ecba6f36e1): complete subsection reference.

- [low](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-a61be0d4b279542f3d137d5fb557fb4930864fff987b1a1d4370b91f138e931e): complete subsection reference.

- [medium](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-9c7424a7a164df7f96c891dd0f14eac0ea43019bfdf27b98a7cfe96b943a2187): complete subsection reference.

<a id="canonical-d77d8e4436066871d0a40e7de1eab787bbaf61264cf4fb125f421d3742cd4249"></a>

## Next pages — mitigation_type.rules.threat_level / 6196c9ba6387 / 4

- [mitigation_type.rules.threat_level.high](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-5f5a849b32ca89d64c6cba83e34a8cb36c8d2b52b25120763fdbf0ecba6f36e1)
- [mitigation_type.rules.threat_level.low](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-a61be0d4b279542f3d137d5fb557fb4930864fff987b1a1d4370b91f138e931e)
- [mitigation_type.rules.threat_level.medium](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-9c7424a7a164df7f96c891dd0f14eac0ea43019bfdf27b98a7cfe96b943a2187)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-5f5a849b32ca89d64c6cba83e34a8cb36c8d2b52b25120763fdbf0ecba6f36e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4aef325e6ec1c1f35019c934eb92f9ffdb1fb91a3f32a39c920491bbc0f21de"></a>

## mitigation_type.rules.threat_level.high — mitigation_type.rules.threat_level.high / 679fff4f0b80 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a)
- mitigation_type.rules.threat_level.high

<a id="canonical-1b52378241bd44ea1a07beeb3481152b66e1fb1544ba18d76de17434add802b9"></a>

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

<a id="canonical-8c354a7697a004abd81c068a736207b69c5d011725987421c3e25a134a8a2445"></a>

## Direct properties — mitigation_type.rules.threat_level.high / 679fff4f0b80 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73af803160be0948cbcac2fb4a840d22f29aa7a58ad210f62e6c3253aa3b767b"></a>

## Next pages — mitigation_type.rules.threat_level.high / 679fff4f0b80 / 4

- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-a61be0d4b279542f3d137d5fb557fb4930864fff987b1a1d4370b91f138e931e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0a2f67e948f937c0e78b1022e69e72da22556b4450043578ec598d8c2788d61"></a>

## mitigation_type.rules.threat_level.low — mitigation_type.rules.threat_level.low / 8ce030cae14d / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a)
- mitigation_type.rules.threat_level.low

<a id="canonical-16b56b5280de9a95b3dff35620363addb8b462985cbf7f37e8c50c7df0347827"></a>

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

<a id="canonical-964defc5befdbef77dd587b8be8be19365b0798cad6ed99cc3b3f79df33945e9"></a>

## Direct properties — mitigation_type.rules.threat_level.low / 8ce030cae14d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c0347a05129a5a8194b9ea1bc0c950a8a35bf78befa03d94a1a6a71f1eaf19f"></a>

## Next pages — mitigation_type.rules.threat_level.low / 8ce030cae14d / 4

- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-9c7424a7a164df7f96c891dd0f14eac0ea43019bfdf27b98a7cfe96b943a2187"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ef23111c054c60f041165c8b2ae5dde34c0bb9441f70b0590581a8f3a15b4ae"></a>

## mitigation_type.rules.threat_level.medium — mitigation_type.rules.threat_level.medium / a419bf39d386 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Property reference](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [mitigation_type](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-0e292f557270a642d33635fb973dc88751fdb9ebfd4cd045b2b719c6eea3a465)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-048981929fca1e2b1513ed88ee17b104d6390b2613036ae4c2969bf132fb56d9)
- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a)
- mitigation_type.rules.threat_level.medium

<a id="canonical-790c260771846d4cf0b27b3160f1d0cc7f8551737404567b160ad146e72a791f"></a>

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

<a id="canonical-02a397fde3c83b0cd0896864420bf0e2bdba678472854ba17b55edf417f17247"></a>

## Direct properties — mitigation_type.rules.threat_level.medium / a419bf39d386 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f131a5aee1f93b0a036055e27e6832b83117b48636ad84d240849a29ced3286"></a>

## Next pages — mitigation_type.rules.threat_level.medium / a419bf39d386 / 4

- [mitigation_type.rules.threat_level](data-sources--malicious_user_mitigation--reference--group-001.md#canonical-916654872d54db88ac337c66bf7a6d5b8a7d4731d4d349cad7ec1af90a237c8a)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
