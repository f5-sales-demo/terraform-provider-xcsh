---
page_title: "xcsh_network_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule reference."
---

# xcsh_network_policy_rule reference

<a id="canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87918eaad6a7f40b4581d7ef6e200eafc0628e84a60a4153aca69fc61e0696a7"></a>

## Property reference — Property reference / 2b8c09173707 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- Property reference

<a id="canonical-379342eb09a583bb847e12987692b52c96e44d1410e918c5911b8421d5bc8f36"></a>

## Direct properties — Property reference / 2b8c09173707 / 3

<a id="canonical-257e44cdb10114696f70dc6c014e74a17b4ebaa874839924102f6b38f74aadda"></a>

<a id="canonical-dab6a6181f2e9c46fbecb3db9c7dab2e43f00cdc00d621a999b0d7e6061c5adc"></a>

## action property — Property reference / 2b8c09173707 / 4

Type: `"string"`. Optional, Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

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

- [advanced_action](resources--network_policy_rule--reference--group-001.md#canonical-a10e18c24124edfd00ba310200b43212c18e4b5aea3485ee8a46cf5beed10dc8): complete subsection reference.

<a id="canonical-77cc731a9987955767eca0ba225d59746ffa9cae396d9e55b4a6fc63fbf2006d"></a>

<a id="canonical-42d2bf529cec2972cb61944d7e4d3f9a41e5f3370eef03f2cfecdff6ddb7e30a"></a>

## annotations property — Property reference / 2b8c09173707 / 5

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

<a id="canonical-ae745135e5526c56746538f5301994aa85809c2a5b3104a998163b4c91c5ca2d"></a>

<a id="canonical-ec5395ffccca87aeaf4d233f7c9886caffaa8689252ebbebbfe97787607fa6a0"></a>

## description property — Property reference / 2b8c09173707 / 6

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

<a id="canonical-c5ce96fa86214708e5bfc5dfcbe4a20923f71de098e2d27d5ca20b86ea0bf6fd"></a>

<a id="canonical-3207c4be003d4f11f391578b5d27d8dba14f38818cb6e31e337be40b291e081c"></a>

## disable property — Property reference / 2b8c09173707 / 7

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

<a id="canonical-100139c68248fb32703af9cf399df0be47a03546e7b3e6830809334dbccad59c"></a>

<a id="canonical-a290b23ea5041d9c0807a9e54c2221f38a9677a3723e30e014f83dec122ad7e2"></a>

## id property — Property reference / 2b8c09173707 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-7c2c3f3518e262c0c55f62a6f99970fa7a334b35089daf828e8a27331115bade): complete subsection reference.

- [label_matcher](resources--network_policy_rule--reference--group-001.md#canonical-ea976c6c270e875f535f143a76118f210e21f63aac1da101c337ece80419b30d): complete subsection reference.

<a id="canonical-80e237a8b83dbe2ebdad82bd2f391fcca72d91f4291a7d18cc6a0afe0d7d488e"></a>

<a id="canonical-3503df75cecf5c0498dba9008add13974bda829bd244d30aad3a3d55346564a8"></a>

## labels property — Property reference / 2b8c09173707 / 9

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

<a id="canonical-5b62eba67a42726183c209233457d78a7f919022271d93593e130111302ea077"></a>

<a id="canonical-7e04dcb31b2b9a8de19975693f5452b017b190922dc3601d249bfb186edee286"></a>

## name property — Property reference / 2b8c09173707 / 10

Type: `"string"`. Required.

Name of the Network Policy Rule. Must be unique within the namespace.

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

<a id="canonical-5c5879669d364726c093d847cbe03196a254295493b5ce405966ad00ccac40ab"></a>

<a id="canonical-d05fb98765e7d7a303c6141c8de13a1d56d5de08ef319e38e50d4ed58990350a"></a>

## namespace property — Property reference / 2b8c09173707 / 11

Type: `"string"`. Required.

Namespace where the Network Policy Rule is created.

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

<a id="canonical-273315acb6aadc6f1e20aa7a84495c959a2019ac37dc0cd8bbed7dcc28b2a8b5"></a>

<a id="canonical-31125635630fd18ef8e46d72a28867a0aa4f9032e3c214c649722b0c4e34ebc1"></a>

## ports property — Property reference / 2b8c09173707 / 12

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [prefix](resources--network_policy_rule--reference--group-001.md#canonical-7af9e1ee3edd74185a365689e2c8bec40e32924c70a802ba6bdc6d3578be337a): complete subsection reference.

- [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-1b7cffa87ec89849753531eb1e30753d46110141e59aa439e14a3256fe6862aa): complete subsection reference.

<a id="canonical-6761467e5d9f902375a358d1367d0c31b1b578fc3ea129bfadf0ae5da815381e"></a>

<a id="canonical-89739581b3cbbd2d56a6d0e5237e82724b7d11a6364753052ddb3da3818f7d00"></a>

## protocol property — Property reference / 2b8c09173707 / 13

Type: `"string"`. Optional, Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

- [timeouts](resources--network_policy_rule--reference--group-001.md#canonical-703888f5013cea6ff95d0565ce5e0c2ab084fb288f170ea6fa25d794b19070be): complete subsection reference.

<a id="canonical-0334f46ab49a4e72405806b371fe3369647b0d7103fead422df1dac482362d79"></a>

## All schema paths — Property reference / 2b8c09173707 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--network_policy_rule--reference--group-001.md#canonical-257e44cdb10114696f70dc6c014e74a17b4ebaa874839924102f6b38f74aadda) |
| `advanced_action` | [advanced_action](resources--network_policy_rule--reference--group-001.md#canonical-675b61176f215131ccdc71f89204ff789ebc7c9a2975e79c8914ce7afba33a85) |
| `advanced_action.action` | [advanced_action.action](resources--network_policy_rule--reference--group-001.md#canonical-8ca0cb984de2f36076431b77cd5db6bd6c9a795edccef4cf461492da4375c125) |
| `annotations` | [annotations](resources--network_policy_rule--reference--group-001.md#canonical-77cc731a9987955767eca0ba225d59746ffa9cae396d9e55b4a6fc63fbf2006d) |
| `description` | [description](resources--network_policy_rule--reference--group-001.md#canonical-ae745135e5526c56746538f5301994aa85809c2a5b3104a998163b4c91c5ca2d) |
| `disable` | [disable](resources--network_policy_rule--reference--group-001.md#canonical-c5ce96fa86214708e5bfc5dfcbe4a20923f71de098e2d27d5ca20b86ea0bf6fd) |
| `id` | [id](resources--network_policy_rule--reference--group-001.md#canonical-100139c68248fb32703af9cf399df0be47a03546e7b3e6830809334dbccad59c) |
| `ip_prefix_set` | [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-68f6943f9a6b8264b464cb4e673f0767b63f1a01c1f899223905855e4d67a0d8) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](resources--network_policy_rule--reference--group-001.md#canonical-545e0196de7b770a40ad599ce2a8c5d87485aa99bb434f5369026e2404e922f1) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](resources--network_policy_rule--reference--group-001.md#canonical-e6e974294dfa1ad7a4781a11f5dcd31a1109ca5ab153bc0757f9a3a9866fc56d) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](resources--network_policy_rule--reference--group-001.md#canonical-dd79bad8dd774f82295b57ee37aaca1cb5a8f7a8e0ab37f2de2ebe6aae68e7c2) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](resources--network_policy_rule--reference--group-001.md#canonical-0d9674b9085aa36560b7b1390918deaf1d58b22b08d5e00ef066128ec38c1395) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](resources--network_policy_rule--reference--group-001.md#canonical-93f3a5461a3cdce93645f0ad7fa1d3fd6bceff12b81cf4431fa8103b4823c2d5) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](resources--network_policy_rule--reference--group-001.md#canonical-8b86badb9f8a57df9f68cd914128afb4ae651777d25c869397047d0b9e60f897) |
| `label_matcher` | [label_matcher](resources--network_policy_rule--reference--group-001.md#canonical-ff584515e1c16a86330c9db56e412d30ef28371fb875d27a25be9811e4930372) |
| `label_matcher.keys` | [label_matcher.keys](resources--network_policy_rule--reference--group-001.md#canonical-16d1d40eab227ca1a42d52d0e23aee83542a6bea1c34485a9148bfa08ba6e81e) |
| `labels` | [labels](resources--network_policy_rule--reference--group-001.md#canonical-80e237a8b83dbe2ebdad82bd2f391fcca72d91f4291a7d18cc6a0afe0d7d488e) |
| `name` | [name](resources--network_policy_rule--reference--group-001.md#canonical-5b62eba67a42726183c209233457d78a7f919022271d93593e130111302ea077) |
| `namespace` | [namespace](resources--network_policy_rule--reference--group-001.md#canonical-5c5879669d364726c093d847cbe03196a254295493b5ce405966ad00ccac40ab) |
| `ports` | [ports](resources--network_policy_rule--reference--group-001.md#canonical-273315acb6aadc6f1e20aa7a84495c959a2019ac37dc0cd8bbed7dcc28b2a8b5) |
| `prefix` | [prefix](resources--network_policy_rule--reference--group-001.md#canonical-1c3223a82730129099ebca198c964bbbde44b343b491587dfc65f2277f7ce453) |
| `prefix.prefix` | [prefix.prefix](resources--network_policy_rule--reference--group-001.md#canonical-0ad0d5ecf972398709e3050d2ec51b7e3cc5a727f70d0e111221f76a03fd6431) |
| `prefix_selector` | [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-d294f1c6ac42af1b205963664f1156a993b6203bc672ef34090ac2af0e8b073f) |
| `prefix_selector.expressions` | [prefix_selector.expressions](resources--network_policy_rule--reference--group-001.md#canonical-73be43bfb4141a3f338578c1615c95f7bdcb5e623ce8c996e2f781defcaf3193) |
| `protocol` | [protocol](resources--network_policy_rule--reference--group-001.md#canonical-6761467e5d9f902375a358d1367d0c31b1b578fc3ea129bfadf0ae5da815381e) |
| `timeouts` | [timeouts](resources--network_policy_rule--reference--group-001.md#canonical-8e99b83f750a484fc71f3945ae75bf3d729986c099b01ff0fa300ace2e646508) |
| `timeouts.create` | [timeouts.create](resources--network_policy_rule--reference--group-001.md#canonical-2aff1e5a726beda0142fa33a7a9b8ea0a7702a3c6ccc64bc6f69aedbb5c02872) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy_rule--reference--group-001.md#canonical-843d7c00ed81f128a14b7d8266dd5845f754ca9d7716af63bef2d9a8d8162b0e) |
| `timeouts.read` | [timeouts.read](resources--network_policy_rule--reference--group-001.md#canonical-799504527bf549f11e2d18420111e479a0bb4778bff4ae95ff07e059e3a2e104) |
| `timeouts.update` | [timeouts.update](resources--network_policy_rule--reference--group-001.md#canonical-65ef87e1cde572d19cbfcdbaf5800c6a999408c44a2141adfb8377ac89379679) |

<a id="canonical-3a96d540fa4747a70da799a76ed2cd7746853566e4ee9cb60c30c97cd69f45b3"></a>

## Next pages — Property reference / 2b8c09173707 / 15

- [advanced_action](resources--network_policy_rule--reference--group-001.md#canonical-a10e18c24124edfd00ba310200b43212c18e4b5aea3485ee8a46cf5beed10dc8)
- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-7c2c3f3518e262c0c55f62a6f99970fa7a334b35089daf828e8a27331115bade)
- [label_matcher](resources--network_policy_rule--reference--group-001.md#canonical-ea976c6c270e875f535f143a76118f210e21f63aac1da101c337ece80419b30d)
- [prefix](resources--network_policy_rule--reference--group-001.md#canonical-7af9e1ee3edd74185a365689e2c8bec40e32924c70a802ba6bdc6d3578be337a)
- [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-1b7cffa87ec89849753531eb1e30753d46110141e59aa439e14a3256fe6862aa)
- [timeouts](resources--network_policy_rule--reference--group-001.md#canonical-703888f5013cea6ff95d0565ce5e0c2ab084fb288f170ea6fa25d794b19070be)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-a10e18c24124edfd00ba310200b43212c18e4b5aea3485ee8a46cf5beed10dc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d63bbf24e1ce0ba9f6e34e32494a43e2af9bfe54a511b70998c5b33726ee2b3"></a>

## advanced_action — advanced_action / 03856a84d276 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- advanced_action

<a id="canonical-675b61176f215131ccdc71f89204ff789ebc7c9a2975e79c8914ce7afba33a85"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
advanced_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2c0793fb39a72e4780166ff6c8d1fc323ca27ef5400ed7c431ba8614ded7b08"></a>

## Direct properties — advanced_action / 03856a84d276 / 3

<a id="canonical-8ca0cb984de2f36076431b77cd5db6bd6c9a795edccef4cf461492da4375c125"></a>

<a id="canonical-935a17f9006975f25d5c8aabbb4c7f33058f9fa712c92f58b1387994a5e95b90"></a>

## action property — advanced_action / 03856a84d276 / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-46f7030ae66ba03ad2e1b6ca7cfb76cc9bc8147193159355cb68e2a1f7244673"></a>

## Next pages — advanced_action / 03856a84d276 / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-7c2c3f3518e262c0c55f62a6f99970fa7a334b35089daf828e8a27331115bade"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d5d9d9052072a20231700ca9379a25f94070fa47625e58958c0266c02df4a93"></a>

## ip_prefix_set — ip_prefix_set / a733778fdb3f / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- ip_prefix_set

<a id="canonical-68f6943f9a6b8264b464cb4e673f0767b63f1a01c1f899223905855e4d67a0d8"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ip\_prefix\_set, prefix, prefix\_selector\] List of references to ip\_prefix\_set objects.

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

OneOf alternatives in this subsection:

- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-68f6943f9a6b8264b464cb4e673f0767b63f1a01c1f899223905855e4d67a0d8)
- [prefix](resources--network_policy_rule--reference--group-001.md#canonical-1c3223a82730129099ebca198c964bbbde44b343b491587dfc65f2277f7ce453)
- [prefix_selector](resources--network_policy_rule--reference--group-001.md#canonical-d294f1c6ac42af1b205963664f1156a993b6203bc672ef34090ac2af0e8b073f)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-27aab77664368e8739f9031bcc5a53cb4b4cae2da009d89178473d557aa08940"></a>

## Direct properties — ip_prefix_set / a733778fdb3f / 3

- [ref](resources--network_policy_rule--reference--group-001.md#canonical-5094f4690055511ecd7868f774e1e12110d856dbb76e5f8cfacbad72e6d4232b): complete subsection reference.

<a id="canonical-87580c2593059283e3fd7ec80979e7be7deaa90c03762d6ec46e982b0a703586"></a>

## Next pages — ip_prefix_set / a733778fdb3f / 4

- [ip_prefix_set.ref](resources--network_policy_rule--reference--group-001.md#canonical-5094f4690055511ecd7868f774e1e12110d856dbb76e5f8cfacbad72e6d4232b)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-5094f4690055511ecd7868f774e1e12110d856dbb76e5f8cfacbad72e6d4232b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5bc725fe851044ebf353db34081f91162ae494f169aeffd92b25d23b0dfdb50"></a>

## ip_prefix_set.ref — ip_prefix_set.ref / ef7f561c11b9 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-7c2c3f3518e262c0c55f62a6f99970fa7a334b35089daf828e8a27331115bade)
- ip_prefix_set.ref

<a id="canonical-545e0196de7b770a40ad599ce2a8c5d87485aa99bb434f5369026e2404e922f1"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2e9e6243c2827efc76d597d6429a1ebd5f5ac1306ad07ed0366469e789148d6"></a>

## Direct properties — ip_prefix_set.ref / ef7f561c11b9 / 3

<a id="canonical-e6e974294dfa1ad7a4781a11f5dcd31a1109ca5ab153bc0757f9a3a9866fc56d"></a>

<a id="canonical-1d2eb644f0f572216fcc2258e1a2c3c9a8e77831a729ca931f6aca61ec47e15c"></a>

## kind property — ip_prefix_set.ref / ef7f561c11b9 / 4

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

<a id="canonical-dd79bad8dd774f82295b57ee37aaca1cb5a8f7a8e0ab37f2de2ebe6aae68e7c2"></a>

<a id="canonical-8f43cb626f91d85e083a228ceb6ceeb600fa2dd7c9099944efbcb0fd0aed145b"></a>

## name property — ip_prefix_set.ref / ef7f561c11b9 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0d9674b9085aa36560b7b1390918deaf1d58b22b08d5e00ef066128ec38c1395"></a>

<a id="canonical-47289dbe791fac52b3152f470485b2649d70963bedd52d37f8d3672ec5af8e99"></a>

## namespace property — ip_prefix_set.ref / ef7f561c11b9 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-93f3a5461a3cdce93645f0ad7fa1d3fd6bceff12b81cf4431fa8103b4823c2d5"></a>

<a id="canonical-0ec5cd8d7d3d7b10a83c07dbe8e68ecf7894031076329c0cf2ddd22024593c40"></a>

## tenant property — ip_prefix_set.ref / ef7f561c11b9 / 7

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

<a id="canonical-8b86badb9f8a57df9f68cd914128afb4ae651777d25c869397047d0b9e60f897"></a>

<a id="canonical-129a1b429d6b10adc75aeb235d364bcd0a97d8928ad075c03857e379127f1b61"></a>

## uid property — ip_prefix_set.ref / ef7f561c11b9 / 8

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

<a id="canonical-8bf9c5d075d72dfd7fd36c40752c034ed9f68f67be62f9c1f2d0331b0e9309d1"></a>

## Next pages — ip_prefix_set.ref / ef7f561c11b9 / 9

- [ip_prefix_set](resources--network_policy_rule--reference--group-001.md#canonical-7c2c3f3518e262c0c55f62a6f99970fa7a334b35089daf828e8a27331115bade)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-ea976c6c270e875f535f143a76118f210e21f63aac1da101c337ece80419b30d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0a64fe1c563001d8c6ac3da7d53b4a3a0a08bacccb03dcdfe5a25b870e20df5"></a>

## label_matcher — label_matcher / 86f185b9d532 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- label_matcher

<a id="canonical-ff584515e1c16a86330c9db56e412d30ef28371fb875d27a25be9811e4930372"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-29d6f4f08fed13a49d898bccf9551ecf02c2e4bb9a4b4ae3b0a9dc32d811825e"></a>

## Direct properties — label_matcher / 86f185b9d532 / 3

<a id="canonical-16d1d40eab227ca1a42d52d0e23aee83542a6bea1c34485a9148bfa08ba6e81e"></a>

<a id="canonical-0f9e55cb08ce9955f717cbe1a360d0e115ecc5147f1028a6ed8f29b303a1c31e"></a>

## keys property — label_matcher / 86f185b9d532 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-300ff452bd756c8ed464e940b455379622626bb9261c9d85e75982a5f3c0406a"></a>

## Next pages — label_matcher / 86f185b9d532 / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-7af9e1ee3edd74185a365689e2c8bec40e32924c70a802ba6bdc6d3578be337a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-535ca0cba6376d3699e2d37dec0141677bf24a61b521b98ee878596befd22dea"></a>

## prefix — prefix / 36a5630bc4a3 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- prefix

<a id="canonical-1c3223a82730129099ebca198c964bbbde44b343b491587dfc65f2277f7ce453"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
prefix {
  # Configure direct properties listed below.
}
```

<a id="canonical-80a2c7858a49488642b1a527b3c279f4ba2e258327f5a9ac49b1d5a2a64fcfe7"></a>

## Direct properties — prefix / 36a5630bc4a3 / 3

<a id="canonical-0ad0d5ecf972398709e3050d2ec51b7e3cc5a727f70d0e111221f76a03fd6431"></a>

<a id="canonical-09a0ad746b27ad6bc3903aa7f7d8869bede1e4f8397731ab1d3482a57e357860"></a>

## prefix property — prefix / 36a5630bc4a3 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

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

<a id="canonical-4b79d46357c3d8f45ceb18932db7093880990dea27de00eaf67151c6cf60131c"></a>

## Next pages — prefix / 36a5630bc4a3 / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-1b7cffa87ec89849753531eb1e30753d46110141e59aa439e14a3256fe6862aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-963a362af939cf90e35952024dcedd22d0190484212614d768a773eae6fc97e8"></a>

## prefix_selector — prefix_selector / dac1d3135d0d / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- prefix_selector

<a id="canonical-d294f1c6ac42af1b205963664f1156a993b6203bc672ef34090ac2af0e8b073f"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
prefix_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-b32e649e13010a4dd03e2c94af7915a7e0a3a5cff48a9a308d6992e71e713684"></a>

## Direct properties — prefix_selector / dac1d3135d0d / 3

<a id="canonical-73be43bfb4141a3f338578c1615c95f7bdcb5e623ce8c996e2f781defcaf3193"></a>

<a id="canonical-aa5b31d460516a3a7e544c553ee36fc45f13d7f9ace10a350e4b695252df331d"></a>

## expressions property — prefix_selector / dac1d3135d0d / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-34ace00a168ea4e5ac18524b03a701cf1063fbac9c6f8a258131f0b90775509e"></a>

## Next pages — prefix_selector / dac1d3135d0d / 5

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-703888f5013cea6ff95d0565ce5e0c2ab084fb288f170ea6fa25d794b19070be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e31cef2371f2de0ef4352cb656494b8189aa5fb519e5c4bff735d5b847528edb"></a>

## timeouts — timeouts / 54a136260beb / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- timeouts

<a id="canonical-8e99b83f750a484fc71f3945ae75bf3d729986c099b01ff0fa300ace2e646508"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-86372b5fe6fec4ace71813e0d80f9cdb0ecb4734ba8ca8cf831dc8578b21185e"></a>

## Direct properties — timeouts / 54a136260beb / 3

<a id="canonical-2aff1e5a726beda0142fa33a7a9b8ea0a7702a3c6ccc64bc6f69aedbb5c02872"></a>

<a id="canonical-3eb94ef0b661e432957158c09afdcc1eb8260028628dd92999ce01665314a346"></a>

## create property — timeouts / 54a136260beb / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-843d7c00ed81f128a14b7d8266dd5845f754ca9d7716af63bef2d9a8d8162b0e"></a>

<a id="canonical-bddf439b31fa51ec67cead1c5fd5d035487ab896eafa2ae7a18b3e3a760389ca"></a>

## delete property — timeouts / 54a136260beb / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-799504527bf549f11e2d18420111e479a0bb4778bff4ae95ff07e059e3a2e104"></a>

<a id="canonical-6385e8be45ab6b5399e3fa0e32e3206cbcae381d4e3efb2b4718d2569fffce5c"></a>

## read property — timeouts / 54a136260beb / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-65ef87e1cde572d19cbfcdbaf5800c6a999408c44a2141adfb8377ac89379679"></a>

<a id="canonical-f77bb1e2c1c831ba17c97a75d16cb00eb3bf7a1bdad28fb627f0dc19081b9e11"></a>

## update property — timeouts / 54a136260beb / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-491c80b6d0b2cdfaaede482e3a43a59f02aea9017151dc7945a375be63ceaae1"></a>

## Next pages — timeouts / 54a136260beb / 8

- [Property reference](resources--network_policy_rule--reference--group-001.md#canonical-c6cf66c6a64fe4824aab545eb09e94fea44fac2a10f7b8a1bea9efc1b39518f0)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
