---
page_title: "xcsh_network_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule reference."
---

# xcsh_network_policy_rule reference

<a id="canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-180682577f0eff7007d499e4306a92b05e56cca234fece2204bf7cecc405427a"></a>

## Property reference — Property reference / 6fe57eb4c426 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- Property reference

<a id="canonical-c17993a844ad004bfe30cdab160d6f771ee322b67a1e27230b4b2fd7c5709159"></a>

## Direct properties — Property reference / 6fe57eb4c426 / 3

<a id="canonical-95738905de5060ad5585e4c2cae2c10c525044f372c2689da3f753e37b8fbfbf"></a>

<a id="canonical-e298b48c5fb1f19d271e6a0759a03929e375bfc38c76daee126a510b90abdcf3"></a>

## action property — Property reference / 6fe57eb4c426 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

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

- [advanced_action](data-sources--network_policy_rule--reference--group-001.md#canonical-4c8276e597e2ae3d06f6d222e3d050fba274be7ece00d1957903eab83cd77fe4): complete subsection reference.

<a id="canonical-ee35cfd8013cf34b6690a9082df4d8045efd8ae39dd05f77b91c74c0aa38e7fc"></a>

<a id="canonical-62c3cebc021254b21fa2541c1609fe7f25cc628a60358b23d4ac981484689c09"></a>

## annotations property — Property reference / 6fe57eb4c426 / 5

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

<a id="canonical-e30f01563afbb1453248849fd7fd7ed87d21efd544d3c558f031b8567663b91d"></a>

<a id="canonical-e6f2ebf576fc2ceb68631db4356a7a091fda7522369ea8e50844788cd0895c3f"></a>

## description property — Property reference / 6fe57eb4c426 / 6

Type: `"string"`. Computed.

Description of the NetworkPolicyRule.

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

<a id="canonical-963b94849a7ed49904461f4596a721961b9980d0056c10a73618b6088a1f3677"></a>

<a id="canonical-279ed1070cf8fffe370df1c0eeb1c0040c2a392635687fd7a20378e034887c1c"></a>

## id property — Property reference / 6fe57eb4c426 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-bb1f54d857565fa232b182406e5509f286e2bb5bcddb11fab559b5cc187fc2b5): complete subsection reference.

- [label_matcher](data-sources--network_policy_rule--reference--group-001.md#canonical-85d959014511e9940b84c748d7499fb7c6c00e770bfde7caeaa0d51a215a398e): complete subsection reference.

<a id="canonical-ca501e3a4c6bd592c3acbf081539e35a6f1678a50e9d923d7e04a74b6e4ac47d"></a>

<a id="canonical-98d4d22dacaf69b54e3cb5f3e31903d97eae0f38899c925319873b55108c0fd3"></a>

## labels property — Property reference / 6fe57eb4c426 / 8

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

<a id="canonical-adf1bc3f25e6c1a0a01f24afd6aec977e28ccc8d22d7bd066c8c62daa10bb648"></a>

<a id="canonical-9e9d7f0988262ec166505b4e9e98f14eb00390978b4d13bdce63e792bce148d1"></a>

## name property — Property reference / 6fe57eb4c426 / 9

Type: `"string"`. Required.

Name of the NetworkPolicyRule.

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

<a id="canonical-94cafd0684b975b15245e21a237821e704f0feca164169e48a14407deaf11275"></a>

<a id="canonical-fae1d8d7481fc7656b665fb1993e21d335c2384da776c23e0241bcbe8cb48bd9"></a>

## namespace property — Property reference / 6fe57eb4c426 / 10

Type: `"string"`. Required.

Namespace where the NetworkPolicyRule exists.

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

<a id="canonical-4de0c155338e136a44b8cbcee727046163300ebfa42a4ec958baf9223c276a5a"></a>

<a id="canonical-f7e178d28bfea0878d538b14a7928eed68fd9858b079d094ab7508c8a3f10ad3"></a>

## ports property — Property reference / 6fe57eb4c426 / 11

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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

- [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-46b3291d46378e8d568c496c3c731940508cb40af3205737616395ae09ec177c): complete subsection reference.

- [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-d37e491dd5b228024b0bcbe44ad289eb39ecff0c232213194d7173a533c9e506): complete subsection reference.

<a id="canonical-3961a9e8bed6eac0aa5b4cf5ed94be5115e6c309f8aab4dfa79a4050bb14cea5"></a>

<a id="canonical-d4c7befe50d8971b42279aa682f38a175d056a5ce2d39e494102c6738c57d352"></a>

## protocol property — Property reference / 6fe57eb4c426 / 12

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

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

<a id="canonical-47d705181356e22c2148c8ce7bc35f4839c2d5963e32a44693f13a2c7553f5d3"></a>

## All schema paths — Property reference / 6fe57eb4c426 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--network_policy_rule--reference--group-001.md#canonical-95738905de5060ad5585e4c2cae2c10c525044f372c2689da3f753e37b8fbfbf) |
| `advanced_action` | [advanced_action](data-sources--network_policy_rule--reference--group-001.md#canonical-e0f1976a784459c103ffa710a715f11c9f05efaa064b16d63af9a19b4cfa06bf) |
| `advanced_action.action` | [advanced_action.action](data-sources--network_policy_rule--reference--group-001.md#canonical-14d84afbe70c2c66b07b82a78f63ff1feb5a30be2d8eda479def50f4c78177cf) |
| `annotations` | [annotations](data-sources--network_policy_rule--reference--group-001.md#canonical-ee35cfd8013cf34b6690a9082df4d8045efd8ae39dd05f77b91c74c0aa38e7fc) |
| `description` | [description](data-sources--network_policy_rule--reference--group-001.md#canonical-e30f01563afbb1453248849fd7fd7ed87d21efd544d3c558f031b8567663b91d) |
| `id` | [id](data-sources--network_policy_rule--reference--group-001.md#canonical-963b94849a7ed49904461f4596a721961b9980d0056c10a73618b6088a1f3677) |
| `ip_prefix_set` | [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-e0b84a44c0ab8f012829dcbb1bddddaa53a5bc5cacc45311c4a21359c0545dfd) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](data-sources--network_policy_rule--reference--group-001.md#canonical-060e1967164ff119093cbc26701d99012b84f644a8c929b5094bd788ec9f0bba) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](data-sources--network_policy_rule--reference--group-001.md#canonical-0e83605a2608035e25a7fc80894723710dea1036989c76cdbeb3e8ca5197850f) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](data-sources--network_policy_rule--reference--group-001.md#canonical-eb3ba6a16b9a0215f87a9eceb27978e78f0073608854e84d523290a83a4f527b) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](data-sources--network_policy_rule--reference--group-001.md#canonical-1804cc2f963b2ab87c9c1c19779803c13fd324c5a118f5353ed4d4c451230a35) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](data-sources--network_policy_rule--reference--group-001.md#canonical-dfc20697f3af1d8786c4b7acc96fb944e3e9dc1d62c483b449156bfe5ee8b9d6) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](data-sources--network_policy_rule--reference--group-001.md#canonical-c64de26cdd2ec49062d71b19fb6a97be24666d831dbabd2db595f10f241c1b4f) |
| `label_matcher` | [label_matcher](data-sources--network_policy_rule--reference--group-001.md#canonical-d0096075e571112b87d517e2008c655c2b97ebaaaa333021621d2669bf145ccb) |
| `label_matcher.keys` | [label_matcher.keys](data-sources--network_policy_rule--reference--group-001.md#canonical-7b6a5a1934a3dab56c01071d35e6240bc106c5bc2346815845e1c44ffd35e421) |
| `labels` | [labels](data-sources--network_policy_rule--reference--group-001.md#canonical-ca501e3a4c6bd592c3acbf081539e35a6f1678a50e9d923d7e04a74b6e4ac47d) |
| `name` | [name](data-sources--network_policy_rule--reference--group-001.md#canonical-adf1bc3f25e6c1a0a01f24afd6aec977e28ccc8d22d7bd066c8c62daa10bb648) |
| `namespace` | [namespace](data-sources--network_policy_rule--reference--group-001.md#canonical-94cafd0684b975b15245e21a237821e704f0feca164169e48a14407deaf11275) |
| `ports` | [ports](data-sources--network_policy_rule--reference--group-001.md#canonical-4de0c155338e136a44b8cbcee727046163300ebfa42a4ec958baf9223c276a5a) |
| `prefix` | [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-14d983635da9fa81ecef2c9a1fade6abd6a37a618b84cc62a80c6ead4fa76d05) |
| `prefix.prefix` | [prefix.prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-4841e22007ff7e68bd4c6f34f11ff9d0a119668b7ec0e5702609db2374dc5341) |
| `prefix_selector` | [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-449bedc477fa9654c741e34722d00e50c41c995d3cf3de9b092c71694fec7c90) |
| `prefix_selector.expressions` | [prefix_selector.expressions](data-sources--network_policy_rule--reference--group-001.md#canonical-5f5853753b8f03b55ecc91556b23c3adb9871d0c0b9978df920917037e3f7716) |
| `protocol` | [protocol](data-sources--network_policy_rule--reference--group-001.md#canonical-3961a9e8bed6eac0aa5b4cf5ed94be5115e6c309f8aab4dfa79a4050bb14cea5) |

<a id="canonical-ce041658353ef1eeaf7c4831f8499170a2d12c3d17b08b326babbe2ceb45bcc5"></a>

## Next pages — Property reference / 6fe57eb4c426 / 14

- [advanced_action](data-sources--network_policy_rule--reference--group-001.md#canonical-4c8276e597e2ae3d06f6d222e3d050fba274be7ece00d1957903eab83cd77fe4)
- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-bb1f54d857565fa232b182406e5509f286e2bb5bcddb11fab559b5cc187fc2b5)
- [label_matcher](data-sources--network_policy_rule--reference--group-001.md#canonical-85d959014511e9940b84c748d7499fb7c6c00e770bfde7caeaa0d51a215a398e)
- [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-46b3291d46378e8d568c496c3c731940508cb40af3205737616395ae09ec177c)
- [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-d37e491dd5b228024b0bcbe44ad289eb39ecff0c232213194d7173a533c9e506)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)

<a id="canonical-4c8276e597e2ae3d06f6d222e3d050fba274be7ece00d1957903eab83cd77fe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f2f41274d0a2aa8ef2567240834c2f06b504e80a4c349ec9a4504aa483dc380"></a>

## advanced_action — advanced_action / 1d38d8ec00b0 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- advanced_action

<a id="canonical-e0f1976a784459c103ffa710a715f11c9f05efaa064b16d63af9a19b4cfa06bf"></a>

Type: `"single"`. Computed.

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

<a id="canonical-4b242c0be319ffc83532f85cba1d9cc98311582b873cbe9d251c52636e232f3b"></a>

## Direct properties — advanced_action / 1d38d8ec00b0 / 3

<a id="canonical-14d84afbe70c2c66b07b82a78f63ff1feb5a30be2d8eda479def50f4c78177cf"></a>

<a id="canonical-02bbc567551dcf05c5a6ce48f0a34851fbfd86f25df8c45690fb8a34cc424184"></a>

## action property — advanced_action / 1d38d8ec00b0 / 4

Type: `"string"`. Computed.

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

<a id="canonical-df6f21e0a2b2724c3cf6deca2f7305af47b60c9c873dad0cd405da6c9d0b1a78"></a>

## Next pages — advanced_action / 1d38d8ec00b0 / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)

<a id="canonical-bb1f54d857565fa232b182406e5509f286e2bb5bcddb11fab559b5cc187fc2b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ec0ff653593f2c4d14e7a57e700b818fda2ebea7a9097c1ba5d5dd0e64573c9"></a>

## ip_prefix_set — ip_prefix_set / ca815c68c70c / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- ip_prefix_set

<a id="canonical-e0b84a44c0ab8f012829dcbb1bddddaa53a5bc5cacc45311c4a21359c0545dfd"></a>

Type: `"single"`. Computed.

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

- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-e0b84a44c0ab8f012829dcbb1bddddaa53a5bc5cacc45311c4a21359c0545dfd)
- [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-14d983635da9fa81ecef2c9a1fade6abd6a37a618b84cc62a80c6ead4fa76d05)
- [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-449bedc477fa9654c741e34722d00e50c41c995d3cf3de9b092c71694fec7c90)

Select alternatives according to the provider validators above.

<a id="canonical-928cead7f7497e37b9990ebbd94a935458019e469d6c21cf98a4b045e3ec3574"></a>

## Direct properties — ip_prefix_set / ca815c68c70c / 3

- [ref](data-sources--network_policy_rule--reference--group-001.md#canonical-54459fb9676daacd515e64821ccb747ec318f0e008e19be3eeff9fe9e2784f74): complete subsection reference.

<a id="canonical-aafbdad20a8178da5e8b05220c080c95f97b18ce81afd6b552c9d268789df954"></a>

## Next pages — ip_prefix_set / ca815c68c70c / 4

- [ip_prefix_set.ref](data-sources--network_policy_rule--reference--group-001.md#canonical-54459fb9676daacd515e64821ccb747ec318f0e008e19be3eeff9fe9e2784f74)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)

<a id="canonical-54459fb9676daacd515e64821ccb747ec318f0e008e19be3eeff9fe9e2784f74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21670ab5c0a51c912a00c106dbfa53fdc131dc6f0774f0b2dce8032443b69d67"></a>

## ip_prefix_set.ref — ip_prefix_set.ref / 03c58c04039c / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-bb1f54d857565fa232b182406e5509f286e2bb5bcddb11fab559b5cc187fc2b5)
- ip_prefix_set.ref

<a id="canonical-060e1967164ff119093cbc26701d99012b84f644a8c929b5094bd788ec9f0bba"></a>

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

<a id="canonical-a38d80c35c40cb9a708023635c53272072e297b750c513b5f9497dfcc066b8a3"></a>

## Direct properties — ip_prefix_set.ref / 03c58c04039c / 3

<a id="canonical-0e83605a2608035e25a7fc80894723710dea1036989c76cdbeb3e8ca5197850f"></a>

<a id="canonical-d5e1a9c9658762032f4e8fce0af4d3b60260bd2395cb2e932da38e0c1f609712"></a>

## kind property — ip_prefix_set.ref / 03c58c04039c / 4

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

<a id="canonical-eb3ba6a16b9a0215f87a9eceb27978e78f0073608854e84d523290a83a4f527b"></a>

<a id="canonical-5d5ddab4506dd287a33f94c6138dc3aaa78bc28f55b96f31800d541bad9b9ccc"></a>

## name property — ip_prefix_set.ref / 03c58c04039c / 5

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

<a id="canonical-1804cc2f963b2ab87c9c1c19779803c13fd324c5a118f5353ed4d4c451230a35"></a>

<a id="canonical-cb3bffacf56baeabd29d3e03e31fb25d76a059f0b24fd1f6e333f29fa84d293a"></a>

## namespace property — ip_prefix_set.ref / 03c58c04039c / 6

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

<a id="canonical-dfc20697f3af1d8786c4b7acc96fb944e3e9dc1d62c483b449156bfe5ee8b9d6"></a>

<a id="canonical-87f17a776bcbb025a61de9f057a7ee65c04d0b25bd06ce7ccd00aa4b60797e44"></a>

## tenant property — ip_prefix_set.ref / 03c58c04039c / 7

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

<a id="canonical-c64de26cdd2ec49062d71b19fb6a97be24666d831dbabd2db595f10f241c1b4f"></a>

<a id="canonical-4abcb19f4456df925128a0f2dd6accb72fbb4a815cab5e337d77c9254c755165"></a>

## uid property — ip_prefix_set.ref / 03c58c04039c / 8

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

<a id="canonical-bc3c0119db7caa5e59de01ae9fe3108b5cf92dc415d27ebf567b3c5f6cb86e65"></a>

## Next pages — ip_prefix_set.ref / 03c58c04039c / 9

- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-bb1f54d857565fa232b182406e5509f286e2bb5bcddb11fab559b5cc187fc2b5)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)

<a id="canonical-85d959014511e9940b84c748d7499fb7c6c00e770bfde7caeaa0d51a215a398e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d609d134bb25951df29734cd9aa70806ee56a8bf7040075dd77e82cdf0e3db1"></a>

## label_matcher — label_matcher / 6e4e1fec28a3 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- label_matcher

<a id="canonical-d0096075e571112b87d517e2008c655c2b97ebaaaa333021621d2669bf145ccb"></a>

Type: `"single"`. Computed.

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

<a id="canonical-d15cfc199e95c6e237eccabe13950a9f464705cea59556410243b73e02c503ad"></a>

## Direct properties — label_matcher / 6e4e1fec28a3 / 3

<a id="canonical-7b6a5a1934a3dab56c01071d35e6240bc106c5bc2346815845e1c44ffd35e421"></a>

<a id="canonical-d584c524f5be8ed7aca39033f39a2274534f0a65d73b6e29567fbb271f8882a7"></a>

## keys property — label_matcher / 6e4e1fec28a3 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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

<a id="canonical-81fcdec996d2353d6cb50f1504ce9bc058806a8d01d2b5a7017def19d79b8585"></a>

## Next pages — label_matcher / 6e4e1fec28a3 / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)

<a id="canonical-46b3291d46378e8d568c496c3c731940508cb40af3205737616395ae09ec177c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f319c5509cf43a3af81777e3bfc5fc1b47c0ee5569b3ac85a14fc09f566a23d6"></a>

## prefix — prefix / 9d84f8fdccec / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- prefix

<a id="canonical-14d983635da9fa81ecef2c9a1fade6abd6a37a618b84cc62a80c6ead4fa76d05"></a>

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

<a id="canonical-e4955e35339c952b1c612f8c85d5d4ec34fd6e9b26c4ceff905635e554311b8f"></a>

## Direct properties — prefix / 9d84f8fdccec / 3

<a id="canonical-4841e22007ff7e68bd4c6f34f11ff9d0a119668b7ec0e5702609db2374dc5341"></a>

<a id="canonical-84ebaef9b48422a98c84f17c980d9fd7db0ef607dd614d67ca77ab30994a7f29"></a>

## prefix property — prefix / 9d84f8fdccec / 4

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

<a id="canonical-ad6869d7c4c4f45c1c99356618a475172b500918168c49d2001260c0d0d5f930"></a>

## Next pages — prefix / 9d84f8fdccec / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)

<a id="canonical-d37e491dd5b228024b0bcbe44ad289eb39ecff0c232213194d7173a533c9e506"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4beabbba40dbd41e6c6e0f7b5e2efe4f1bb3f7e9d4bc524dfbfda35438c628a1"></a>

## prefix_selector — prefix_selector / 8e2467e928cb / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- prefix_selector

<a id="canonical-449bedc477fa9654c741e34722d00e50c41c995d3cf3de9b092c71694fec7c90"></a>

Type: `"single"`. Computed.

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

<a id="canonical-976de8776a001585d20f015278842dfef66b5195c44c40f51d691b06ba97ff3c"></a>

## Direct properties — prefix_selector / 8e2467e928cb / 3

<a id="canonical-5f5853753b8f03b55ecc91556b23c3adb9871d0c0b9978df920917037e3f7716"></a>

<a id="canonical-5f2e46def5390179f41156f76c18175b795a872efe181613b2564464cc54a703"></a>

## expressions property — prefix_selector / 8e2467e928cb / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-61deeba66d8c36aaed598edab0d9ebb4a74e6e4b0744c5a30f80cb04d2f7a44b"></a>

## Next pages — prefix_selector / 8e2467e928cb / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
