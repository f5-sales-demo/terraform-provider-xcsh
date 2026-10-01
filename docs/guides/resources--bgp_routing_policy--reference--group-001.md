---
page_title: "xcsh_bgp_routing_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy reference."
---

# xcsh_bgp_routing_policy reference

<a id="canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-262bf249aef212af65bda357089cc622aa26307ab92b3ed82a7d1fd3fe4a1cbc"></a>

## Property reference — Property reference / d047a8a4f446 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- Property reference

<a id="canonical-ae19a4cab6c3e4d35bc7714fd0e7d5b00af32209ab29e38874847e8ff0fabd66"></a>

## Direct properties — Property reference / d047a8a4f446 / 3

<a id="canonical-7456ccfd62bcb1b9ba02c0139d0c2f109e21d8bad4ae799534f8933711a88184"></a>

<a id="canonical-33761ae35589d5e4c574741e32e2834dcbd7536d577b520286f5bad2d0beb67d"></a>

## annotations property — Property reference / d047a8a4f446 / 4

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

<a id="canonical-c3367eeadd85347ae90d01ad52bca647af7806e6bda50e7faeb33e2ca2bb51d1"></a>

<a id="canonical-5a42c1f483fbfc5deafaf122479c478f9f70101c036d16ec01eeeeff85d02e1b"></a>

## description property — Property reference / d047a8a4f446 / 5

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

<a id="canonical-ada8b210ad685d67b764e47ff981704931a8b9a7b1a53aa762b9e9372c9e3257"></a>

<a id="canonical-40386317ae480d9aa0da6ce6c4fcbd7402c9ea54301350a35a90a2319a8e29f5"></a>

## disable property — Property reference / d047a8a4f446 / 6

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

<a id="canonical-673083bb30282c1a95368bc07120327c580ee54735e133777bbc763f35e1663d"></a>

<a id="canonical-088cb12868a8d0de9ac2725de1da0da12b7d413aeeb2672be320c3b2bcd58cea"></a>

## id property — Property reference / d047a8a4f446 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-488491ab7acf80739377d8caca30b97a3c5dd55a757b5ad3525cc51611773147"></a>

<a id="canonical-4c14dba90538d9f7dd536cb4ae4d2e5213689e0954d789d7505d9c8e05395c64"></a>

## labels property — Property reference / d047a8a4f446 / 8

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

<a id="canonical-5053eebcf3931cdaecf7b3fdf1ac327a060ed2ffa3e2f2f1877851512d8163b7"></a>

<a id="canonical-6c2937b3ce616107968b60e0c0ae2ea994dbd6164a1694310e39f8bec5f43b4d"></a>

## name property — Property reference / d047a8a4f446 / 9

Type: `"string"`. Required.

Name of the BGP Routing Policy. Must be unique within the namespace.

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

<a id="canonical-9cb158619a6204364761415fb9a316ae9e1333d917802e041d27cfc4a3381a45"></a>

<a id="canonical-acc9e9c88040eb9b5b2dfe2dd5e9873aa3b57527d074208b3de29b94135fe5f6"></a>

## namespace property — Property reference / d047a8a4f446 / 10

Type: `"string"`. Required.

Namespace where the BGP Routing Policy is created.

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

- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9): complete subsection reference.

- [timeouts](resources--bgp_routing_policy--reference--group-001.md#canonical-c9a860ff79fefdf349a65cfb4f7a31c8f86b3cac9cc17e7425d6754034b3424b): complete subsection reference.

<a id="canonical-630ed538b7387c667f1a8777b6598c31543706f0d66e1202336a3d73fd686924"></a>

## All schema paths — Property reference / d047a8a4f446 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bgp_routing_policy--reference--group-001.md#canonical-7456ccfd62bcb1b9ba02c0139d0c2f109e21d8bad4ae799534f8933711a88184) |
| `description` | [description](resources--bgp_routing_policy--reference--group-001.md#canonical-c3367eeadd85347ae90d01ad52bca647af7806e6bda50e7faeb33e2ca2bb51d1) |
| `disable` | [disable](resources--bgp_routing_policy--reference--group-001.md#canonical-ada8b210ad685d67b764e47ff981704931a8b9a7b1a53aa762b9e9372c9e3257) |
| `id` | [id](resources--bgp_routing_policy--reference--group-001.md#canonical-673083bb30282c1a95368bc07120327c580ee54735e133777bbc763f35e1663d) |
| `labels` | [labels](resources--bgp_routing_policy--reference--group-001.md#canonical-488491ab7acf80739377d8caca30b97a3c5dd55a757b5ad3525cc51611773147) |
| `name` | [name](resources--bgp_routing_policy--reference--group-001.md#canonical-5053eebcf3931cdaecf7b3fdf1ac327a060ed2ffa3e2f2f1877851512d8163b7) |
| `namespace` | [namespace](resources--bgp_routing_policy--reference--group-001.md#canonical-9cb158619a6204364761415fb9a316ae9e1333d917802e041d27cfc4a3381a45) |
| `rules` | [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-e986fce8c23ad5ecddf817b7ef91dae83225a7de012b7a68eaeeed647ffe4449) |
| `rules.action` | [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-b8bd1c790b010381bb702a1e82d38ab97f8a51bd1d94d8c368f9c094efd81314) |
| `rules.action.allow` | [rules.action.allow](resources--bgp_routing_policy--reference--group-001.md#canonical-f299f88510b8b2d8574540672e57083989d01ef0817f3fd9926f6bf3badf7964) |
| `rules.action.as_path` | [rules.action.as_path](resources--bgp_routing_policy--reference--group-001.md#canonical-76d84835f4d1bd3cc9f0cc1b32c912cc0dd2507921f8d8f8747579a33eddb709) |
| `rules.action.community` | [rules.action.community](resources--bgp_routing_policy--reference--group-001.md#canonical-4a85bb2a2cd16070b7f5b7e601ae9a34cbf702075c642a109ef610f2c7d128d9) |
| `rules.action.community.community` | [rules.action.community.community](resources--bgp_routing_policy--reference--group-001.md#canonical-8ea58b18fdf97c142da14d25a54cb9ae4d2f28176ac7be5bf3dced119e2c3a51) |
| `rules.action.deny` | [rules.action.deny](resources--bgp_routing_policy--reference--group-001.md#canonical-16f3b60204c540bd343b9359773e6bad3c52bb05ec0a64115e81c13f7c77110b) |
| `rules.action.local_preference` | [rules.action.local_preference](resources--bgp_routing_policy--reference--group-001.md#canonical-b15489816388984332dc0fdc2d57189a73bca9ffa67bbc2e751ce4f91b0fd5b9) |
| `rules.action.metric` | [rules.action.metric](resources--bgp_routing_policy--reference--group-001.md#canonical-fb04f68c5114c3582d817fb0c87f18f32d8134075c24ce295e1980ce6144cd34) |
| `rules.match` | [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-1fbe514c4ed589640871f1e7461309cff108fc2605d22a5df07eb9b7413a55c6) |
| `rules.match.as_path` | [rules.match.as_path](resources--bgp_routing_policy--reference--group-001.md#canonical-fcb4bcb65b15764c74353316a0e002457176e7057fe65a592fe7faad565f5d7c) |
| `rules.match.community` | [rules.match.community](resources--bgp_routing_policy--reference--group-001.md#canonical-548a169912b0a2dcf7f46ae8774307e2a110089078fc6500d87e1e098d104ef4) |
| `rules.match.community.community` | [rules.match.community.community](resources--bgp_routing_policy--reference--group-001.md#canonical-947fefc4cccd3b75883b85d59589727d67abc68503cb444c8bd66636db33d9b7) |
| `rules.match.ip_prefixes` | [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-f7542c575a47b81660843c88e8de4627d36ef2126434342f812a7b3a136c1f0f) |
| `rules.match.ip_prefixes.prefixes` | [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-594e2f8d2507d0df1af1db9b86a9e222590b52022e35b2c9ffe84f2581379d04) |
| `rules.match.ip_prefixes.prefixes.equal_or_longer_than` | [rules.match.ip_prefixes.prefixes.equal_or_longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-c334e13fa88b9312ea63a79c3370d72d1976bda3af32de6eafeee7317ff80534) |
| `rules.match.ip_prefixes.prefixes.exact_match` | [rules.match.ip_prefixes.prefixes.exact_match](resources--bgp_routing_policy--reference--group-001.md#canonical-5ead52d33ddf8fedddba8d722ad1490f776b93c004986531f5210c27bd961fd9) |
| `rules.match.ip_prefixes.prefixes.ip_prefixes` | [rules.match.ip_prefixes.prefixes.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-14073c7ede6e74aeaefa3f03869567ccbb1756bf2de274800d800004f034c959) |
| `rules.match.ip_prefixes.prefixes.longer_than` | [rules.match.ip_prefixes.prefixes.longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-770e41769121abc3af6d1ec3955be37b2288ab658ced136b8650256dec40a0de) |
| `timeouts` | [timeouts](resources--bgp_routing_policy--reference--group-001.md#canonical-868da87c22b9a277e317a3f6ae974632bf6f95e5808898dcbcf8aafd39d21cf1) |
| `timeouts.create` | [timeouts.create](resources--bgp_routing_policy--reference--group-001.md#canonical-cbd1dc2be7fd3e786e4b3269b6fc065c1e5ceafd47eb4e0255ce474ba3a74f69) |
| `timeouts.delete` | [timeouts.delete](resources--bgp_routing_policy--reference--group-001.md#canonical-c5a02843cf3ba7128f9e6585adc4b9bc659683416790749088919132c8239953) |
| `timeouts.read` | [timeouts.read](resources--bgp_routing_policy--reference--group-001.md#canonical-3970684b4ed02e3c99062089764f1abac323325f3ebebf4c3f55de532660fae0) |
| `timeouts.update` | [timeouts.update](resources--bgp_routing_policy--reference--group-001.md#canonical-2e65f5b732abffa9552876c22d2c6255654f4662033b3630a5124e24adecfc9a) |

<a id="canonical-cf9bf97e035ace2328eab24677efb8238e3dcf73281dfe1fdebe4a2eaaa40979"></a>

## Next pages — Property reference / d047a8a4f446 / 12

- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [timeouts](resources--bgp_routing_policy--reference--group-001.md#canonical-c9a860ff79fefdf349a65cfb4f7a31c8f86b3cac9cc17e7425d6754034b3424b)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06134db78d356c471de4366016c1f67edceeb12e2a3f8e9964cb808a222542c0"></a>

## rules — rules / db5cb0907c92 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- rules

<a id="canonical-e986fce8c23ad5ecddf817b7ef91dae83225a7de012b7a68eaeeed647ffe4449"></a>

Type: `"object"`. list nested block, Optional.

BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Upstream description:

A BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

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
    "uniqueItems": false
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-8245030568359036e0dc0fd0ab39f3ffac84f38a91110c0b5127f0cb438debdd"></a>

## Direct properties — rules / db5cb0907c92 / 3

- [action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b): complete subsection reference.

- [match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55): complete subsection reference.

<a id="canonical-fabdfcedcce076bdd179e4d23106fde03970d50a18759e6733a02b2c3b2038ed"></a>

## Next pages — rules / db5cb0907c92 / 4

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d01a4de850dd7a14618108afd9d191c74f5ed97a20d8acc3fea367e3ecf6be56"></a>

## rules.action — rules.action / 1511ced9efcc / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- rules.action

<a id="canonical-b8bd1c790b010381bb702a1e82d38ab97f8a51bd1d94d8c368f9c094efd81314"></a>

Type: `"object"`. single nested block, Optional.

Action to be enforced if the BGP route matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "as_path"),
  validators.ConflictingObjectAttributes("allow",
    "community"),
  validators.ConflictingObjectAttributes("allow",
    "deny"),
  validators.ConflictingObjectAttributes("allow",
    "local_preference"),
  validators.ConflictingObjectAttributes("allow",
    "metric"),
  validators.ConflictingObjectAttributes("as_path",
    "community"),
  validators.ConflictingObjectAttributes("as_path",
    "deny"),
  validators.ConflictingObjectAttributes("as_path",
    "local_preference"),
  validators.ConflictingObjectAttributes("as_path",
    "metric"),
  validators.ConflictingObjectAttributes("community",
    "deny"),
  validators.ConflictingObjectAttributes("community",
    "local_preference"),
  validators.ConflictingObjectAttributes("community",
    "metric"),
  validators.ConflictingObjectAttributes("deny",
    "local_preference"),
  validators.ConflictingObjectAttributes("deny",
    "metric"),
  validators.ConflictingObjectAttributes("local_preference",
    "metric")}
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
  "x-ves-oneof-field-action_type": "[\"allow\",\"as_path\",\"community\",\"deny\",\"local_preference\",\"metric\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-8804140db5bd0bf652bfe5a34c698e701202a0d3b5dcb30714519519d0f622be"></a>

## Direct properties — rules.action / 1511ced9efcc / 3

- [allow](resources--bgp_routing_policy--reference--group-001.md#canonical-0972ecae3711514da6109bd2327efe60669f3436cbf78a8a28bbbf30e158a8e7): complete subsection reference.

<a id="canonical-76d84835f4d1bd3cc9f0cc1b32c912cc0dd2507921f8d8f8747579a33eddb709"></a>

<a id="canonical-bc3cbbc39c0f370ac80f3bde5948de4ab79cf36b25e5ce953fbecaa51c5ad227"></a>

## as_path property — rules.action / 1511ced9efcc / 4

Type: `"string"`. Optional.

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

Upstream description:

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

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

- [community](resources--bgp_routing_policy--reference--group-001.md#canonical-3f51aedacbcc72c2a2cdd2342ee8ac48f0a63d604683a27d28aa1ad37b521ddc): complete subsection reference.

- [deny](resources--bgp_routing_policy--reference--group-001.md#canonical-d3cf61fba3ed3d52903f95b38435d118930feac74ffc376bf3d044d350d62445): complete subsection reference.

<a id="canonical-b15489816388984332dc0fdc2d57189a73bca9ffa67bbc2e751ce4f91b0fd5b9"></a>

<a id="canonical-9761d14fdca561fd15ff207058992ecbc921f38e8e0a2bb58b56ebb8763c7e27"></a>

## local_preference property — rules.action / 1511ced9efcc / 5

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

Upstream description:

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

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

<a id="canonical-fb04f68c5114c3582d817fb0c87f18f32d8134075c24ce295e1980ce6144cd34"></a>

<a id="canonical-dc10d1c56e892fbeb50dbdbbbee9ee54203e4c5be9e6cd2266d57342bd425dc5"></a>

## metric property — rules.action / 1511ced9efcc / 6

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

Upstream description:

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

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

<a id="canonical-0d0a20355c0dbf814f720ce2bbfae078374bf6f2e07b1d3a9f3920df80f964bf"></a>

## Next pages — rules.action / 1511ced9efcc / 7

- [rules.action.allow](resources--bgp_routing_policy--reference--group-001.md#canonical-0972ecae3711514da6109bd2327efe60669f3436cbf78a8a28bbbf30e158a8e7)
- [rules.action.community](resources--bgp_routing_policy--reference--group-001.md#canonical-3f51aedacbcc72c2a2cdd2342ee8ac48f0a63d604683a27d28aa1ad37b521ddc)
- [rules.action.deny](resources--bgp_routing_policy--reference--group-001.md#canonical-d3cf61fba3ed3d52903f95b38435d118930feac74ffc376bf3d044d350d62445)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-0972ecae3711514da6109bd2327efe60669f3436cbf78a8a28bbbf30e158a8e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cd572fa2286addc778f986f47d3dd2317b1f441c927801abfd207be65b147fd"></a>

## rules.action.allow — rules.action.allow / df23eb244037 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b)
- rules.action.allow

<a id="canonical-f299f88510b8b2d8574540672e57083989d01ef0817f3fd9926f6bf3badf7964"></a>

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
allow = {}
```

<a id="canonical-1437c9a85327c1fb9950a5614a621489dac5b9a1e4e147573c1c061ac7850fb4"></a>

## Direct properties — rules.action.allow / df23eb244037 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf20d95c74788c8ecfc6d773b354d1dfb274b61bcbadc4e6d5330671b95ca888"></a>

## Next pages — rules.action.allow / df23eb244037 / 4

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-3f51aedacbcc72c2a2cdd2342ee8ac48f0a63d604683a27d28aa1ad37b521ddc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6eb5245c20be5a1239cd1d7b12b4114349a763e01d179deee0f39aefdc8ad3c1"></a>

## rules.action.community — rules.action.community / 226f08f5de0d / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b)
- rules.action.community

<a id="canonical-4a85bb2a2cd16070b7f5b7e601ae9a34cbf702075c642a109ef610f2c7d128d9"></a>

Type: `"object"`. single nested block, Optional.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("community")}
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
community {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d27a83259d1e1cf8c9e95fefa5788284fbd152e568e656ab65c790fde7618fb"></a>

## Direct properties — rules.action.community / 226f08f5de0d / 3

<a id="canonical-8ea58b18fdf97c142da14d25a54cb9ae4d2f28176ac7be5bf3dced119e2c3a51"></a>

<a id="canonical-4db24fb2863a6126b461046365a67990eb0783ff2064be600afe6fe60c309e71"></a>

## community property — rules.action.community / 226f08f5de0d / 4

Type: `["list", "string"]`. Optional.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-41b535a49bd55c4561cc66bc5ed7a2168c18866f4adc3a335d1f792524898682"></a>

## Next pages — rules.action.community / 226f08f5de0d / 5

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-d3cf61fba3ed3d52903f95b38435d118930feac74ffc376bf3d044d350d62445"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9eace889cb861f19fb618a4055519e9c2edf14f4f1ef202c6d941d7ed47ee63a"></a>

## rules.action.deny — rules.action.deny / bdddbe9a5256 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b)
- rules.action.deny

<a id="canonical-16f3b60204c540bd343b9359773e6bad3c52bb05ec0a64115e81c13f7c77110b"></a>

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
deny = {}
```

<a id="canonical-7d4c8bb7e3d1f4ec4d177b270abdc0df2b0a1b3903eb927630668ca33f4f8254"></a>

## Direct properties — rules.action.deny / bdddbe9a5256 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5df882dffdd6d1dcf09608641d834afc03caaf90225ff3f6a3f0b19608faa6d2"></a>

## Next pages — rules.action.deny / bdddbe9a5256 / 4

- [rules.action](resources--bgp_routing_policy--reference--group-001.md#canonical-997cbbaacae511c8d52a466868b798e103c930a784e99d9ce2282a8935f3760b)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-362ee9ad8d493fc754ff6795314c3cd13ffc12ae40e869392fff99fc3449e02d"></a>

## rules.match — rules.match / 19cef3727bad / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- rules.match

<a id="canonical-1fbe514c4ed589640871f1e7461309cff108fc2605d22a5df07eb9b7413a55c6"></a>

Type: `"object"`. single nested block, Optional.

Predicates which have to match information in route for action to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("as_path",
    "community"),
  validators.ConflictingObjectAttributes("as_path",
    "ip_prefixes"),
  validators.ConflictingObjectAttributes("community",
    "ip_prefixes")}
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
  "x-ves-oneof-field-type_of_match": "[\"as_path\",\"community\",\"ip_prefixes\"]"
}
```

Terraform syntax:

```terraform
match {
  # Configure direct properties listed below.
}
```

<a id="canonical-2651234068c4d52bb8455d7c8ce68e516229fb3cedca0a706d50ae572ea53dba"></a>

## Direct properties — rules.match / 19cef3727bad / 3

<a id="canonical-fcb4bcb65b15764c74353316a0e002457176e7057fe65a592fe7faad565f5d7c"></a>

<a id="canonical-8bd986f1cac556e59cf03809184cc9b4474683e9c6e40ec38cd027effda29eb3"></a>

## as_path property — rules.match / 19cef3727bad / 4

Type: `"string"`. Optional.

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

Upstream description:

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

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

- [community](resources--bgp_routing_policy--reference--group-001.md#canonical-6f178a053ed67531db794caf20be076078e006a0e33ba7537e8c2c928ae0231c): complete subsection reference.

- [ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8): complete subsection reference.

<a id="canonical-dd8a48e28baa8f54fe71eb7fc5c46b2ca59798e5edb604f89cc5fcbb51a73e10"></a>

## Next pages — rules.match / 19cef3727bad / 5

- [rules.match.community](resources--bgp_routing_policy--reference--group-001.md#canonical-6f178a053ed67531db794caf20be076078e006a0e33ba7537e8c2c928ae0231c)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-6f178a053ed67531db794caf20be076078e006a0e33ba7537e8c2c928ae0231c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04c8ad1ec81a9c895a733e2afb582675ab1744b56629ac3bef62224c2d536bf4"></a>

## rules.match.community — rules.match.community / a553d190b7f5 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- rules.match.community

<a id="canonical-548a169912b0a2dcf7f46ae8774307e2a110089078fc6500d87e1e098d104ef4"></a>

Type: `"object"`. single nested block, Optional.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("community")}
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
community {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ef3e66f0fefd302f46c2c6959b1f41218c24bbe7284ce8657a75fc27779a967"></a>

## Direct properties — rules.match.community / a553d190b7f5 / 3

<a id="canonical-947fefc4cccd3b75883b85d59589727d67abc68503cb444c8bd66636db33d9b7"></a>

<a id="canonical-ef2051b04cd2c9e2cbddc973b4cbd957ba60b0ecb869d44d25a1dd2d007eb645"></a>

## community property — rules.match.community / a553d190b7f5 / 4

Type: `["list", "string"]`. Optional.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e6d6e3c876b72672d2cc9db27bce304e63bb40d8a794992464658c344a1bb058"></a>

## Next pages — rules.match.community / a553d190b7f5 / 5

- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8034b42e4140d659fc2ac338421b969a0ff7ae60db72209e25e7ad494a92aae"></a>

## rules.match.ip_prefixes — rules.match.ip_prefixes / 1f38119bc601 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- rules.match.ip_prefixes

<a id="canonical-f7542c575a47b81660843c88e8de4627d36ef2126434342f812a7b3a136c1f0f"></a>

Type: `"object"`. single nested block, Optional.

List of IP prefix and prefix length range match condition.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefixes")}
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
ip_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f88be43e6f085ef4bd66bd3e36a6cad59d0b697a69b5d8dabf7c61dfdd2b4f5"></a>

## Direct properties — rules.match.ip_prefixes / 1f38119bc601 / 3

- [prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b): complete subsection reference.

<a id="canonical-e0842cbbe1ed208af793855235b10ea733e4abbde0cdf70850a71e4414a08803"></a>

## Next pages — rules.match.ip_prefixes / 1f38119bc601 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-983cf2bfc7f271211e6bccbeed64027b051f4ea5896a5ff2842d36829991080c"></a>

## rules.match.ip_prefixes.prefixes — rules.match.ip_prefixes.prefixes / 38965fc7da35 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8)
- rules.match.ip_prefixes.prefixes

<a id="canonical-594e2f8d2507d0df1af1db9b86a9e222590b52022e35b2c9ffe84f2581379d04"></a>

Type: `"object"`. list nested block, Optional.

Prefix list. List of IP prefix.

Upstream description:

List of IP prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("equal_or_longer_than",
    "exact_match"),
  validators.ConflictingListObjectAttributes("equal_or_longer_than",
    "longer_than"),
  validators.ConflictingListObjectAttributes("exact_match",
    "longer_than")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-9abc53460f94d2cc5ab2696c2f562e41874ef751cc3f4b9e81554aa9c21ef4ef"></a>

## Direct properties — rules.match.ip_prefixes.prefixes / 38965fc7da35 / 3

- [equal_or_longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-ac6c12959712b3909067aa491e6fa6733c1c6d3e2274be46a8366d1cf925b0b1): complete subsection reference.

- [exact_match](resources--bgp_routing_policy--reference--group-001.md#canonical-e1aea0312f8218a84dbc758dd290e8ddf3a75178e87c2d280eace748b6cec184): complete subsection reference.

<a id="canonical-14073c7ede6e74aeaefa3f03869567ccbb1756bf2de274800d800004f034c959"></a>

<a id="canonical-cb571d59c973748b037abf9e2eea2d33f41100af43cf9d88d16f7646b7f878ed"></a>

## ip_prefixes property — rules.match.ip_prefixes.prefixes / 38965fc7da35 / 4

Type: `"string"`. Optional.

IP Prefix. IP prefix to match on BGP route.

Upstream description:

IP prefix to match on BGP route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-c6c1ba946585e2fcdf3cb05478f229356277c4a333b75f1b3b71796cec42f3c4): complete subsection reference.

<a id="canonical-f21f157698454c34878a538f85f8e647a5c04ac4b149daa9a0914603829641d3"></a>

## Next pages — rules.match.ip_prefixes.prefixes / 38965fc7da35 / 5

- [rules.match.ip_prefixes.prefixes.equal_or_longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-ac6c12959712b3909067aa491e6fa6733c1c6d3e2274be46a8366d1cf925b0b1)
- [rules.match.ip_prefixes.prefixes.exact_match](resources--bgp_routing_policy--reference--group-001.md#canonical-e1aea0312f8218a84dbc758dd290e8ddf3a75178e87c2d280eace748b6cec184)
- [rules.match.ip_prefixes.prefixes.longer_than](resources--bgp_routing_policy--reference--group-001.md#canonical-c6c1ba946585e2fcdf3cb05478f229356277c4a333b75f1b3b71796cec42f3c4)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-ac6c12959712b3909067aa491e6fa6733c1c6d3e2274be46a8366d1cf925b0b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5b99162bdf9bf7ff2d2a714a13e5302bc81271027f86217a4b653db515cfdf1"></a>

## rules.match.ip_prefixes.prefixes.equal_or_longer_than — rules.match.ip_prefixes.prefixes.equal_or_longer_than / ba78904022e1 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8)
- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b)
- rules.match.ip_prefixes.prefixes.equal_or_longer_than

<a id="canonical-c334e13fa88b9312ea63a79c3370d72d1976bda3af32de6eafeee7317ff80534"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for equal or longer than.

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
equal_or_longer_than = {}
```

<a id="canonical-f2728d3a5312f956e4ca92f88fe88a104cd51d0ed15fb91f860714794e0b3d5b"></a>

## Direct properties — rules.match.ip_prefixes.prefixes.equal_or_longer_than / ba78904022e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7261d3c937ad718f59fa7ea2951be6a68c8613ab59f6d801eea6591c169057a"></a>

## Next pages — rules.match.ip_prefixes.prefixes.equal_or_longer_than / ba78904022e1 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-e1aea0312f8218a84dbc758dd290e8ddf3a75178e87c2d280eace748b6cec184"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-164ba80892bd8b3a08c78763f880d47ef06250eff09b0ebef0671f4e1bf70f77"></a>

## rules.match.ip_prefixes.prefixes.exact_match — rules.match.ip_prefixes.prefixes.exact_match / aeb6b1314504 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8)
- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b)
- rules.match.ip_prefixes.prefixes.exact_match

<a id="canonical-5ead52d33ddf8fedddba8d722ad1490f776b93c004986531f5210c27bd961fd9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exact match.

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
exact_match = {}
```

<a id="canonical-25ac805be537a89340dcbec275919bd90e155585cbef01f6760de8d1a4e1257d"></a>

## Direct properties — rules.match.ip_prefixes.prefixes.exact_match / aeb6b1314504 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c37a87187d6f4b73f9dea62464f830e5998fb8129d08c8cf811bcb61b24eb01f"></a>

## Next pages — rules.match.ip_prefixes.prefixes.exact_match / aeb6b1314504 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-c6c1ba946585e2fcdf3cb05478f229356277c4a333b75f1b3b71796cec42f3c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01d317980bfe45da6117d664460b867f7e6c19812f8542c428f3bba0a4b9a25c"></a>

## rules.match.ip_prefixes.prefixes.longer_than — rules.match.ip_prefixes.prefixes.longer_than / db0a4821b136 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [rules](resources--bgp_routing_policy--reference--group-001.md#canonical-d082d86afd3213490a5bb95a24a42989f337fe40113b0a82668bad930713e4b9)
- [rules.match](resources--bgp_routing_policy--reference--group-001.md#canonical-fad2747985aecdb6862ec68f1b404a4f099c3d739f02bbdeea07c74565df8b55)
- [rules.match.ip_prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-28425111739c40f576f9d6f33817f5964461121745e092b60e34a0a94ffb8bc8)
- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b)
- rules.match.ip_prefixes.prefixes.longer_than

<a id="canonical-770e41769121abc3af6d1ec3955be37b2288ab658ced136b8650256dec40a0de"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for longer than.

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
longer_than = {}
```

<a id="canonical-c60bc58879b3cd052c8a885f9cbc11440576f116101e2a6248c4389ae867f9bd"></a>

## Direct properties — rules.match.ip_prefixes.prefixes.longer_than / db0a4821b136 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6021c2eb7b102c36ad1ce6e134fce8c3901aa26b276b23905e2ddc1234971643"></a>

## Next pages — rules.match.ip_prefixes.prefixes.longer_than / db0a4821b136 / 4

- [rules.match.ip_prefixes.prefixes](resources--bgp_routing_policy--reference--group-001.md#canonical-1e63c1fc6639a531a8d3e786c79b5a7dadbcd05713e07dd21dc90b1ef1a7d24b)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)

<a id="canonical-c9a860ff79fefdf349a65cfb4f7a31c8f86b3cac9cc17e7425d6754034b3424b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02b56fcfcc55becbb2ed1b8ea76f9ccd47e1cb888b7a1159af2a1550671accfe"></a>

## timeouts — timeouts / 907c553db6e6 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- timeouts

<a id="canonical-868da87c22b9a277e317a3f6ae974632bf6f95e5808898dcbcf8aafd39d21cf1"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3a1c2acbbdae3707d687ca317d95118c5a1ecc8e6e151a73f7a2ca54756929b7"></a>

## Direct properties — timeouts / 907c553db6e6 / 3

<a id="canonical-cbd1dc2be7fd3e786e4b3269b6fc065c1e5ceafd47eb4e0255ce474ba3a74f69"></a>

<a id="canonical-146afc1b4503e5ca0ed60ba8535dea72780880bb4457886eae8d2273c6c56516"></a>

## create property — timeouts / 907c553db6e6 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c5a02843cf3ba7128f9e6585adc4b9bc659683416790749088919132c8239953"></a>

<a id="canonical-d458f4022bd51a6544cf56307e634c3e0d32e6816e7790fe7c45d33e98b7cd0d"></a>

## delete property — timeouts / 907c553db6e6 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3970684b4ed02e3c99062089764f1abac323325f3ebebf4c3f55de532660fae0"></a>

<a id="canonical-d26a7cb2b286dcee1402df0acfee4d5187a7ef2339a93a29037d6626f8e790f3"></a>

## read property — timeouts / 907c553db6e6 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2e65f5b732abffa9552876c22d2c6255654f4662033b3630a5124e24adecfc9a"></a>

<a id="canonical-517727ee71696a8783700578cff4798dae9cc59f3035c0ba223e565edd462eae"></a>

## update property — timeouts / 907c553db6e6 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0982fe8a8577bdbe4f97dd414d646902f8604a8837a8bfbbaf03c6472b91e5f5"></a>

## Next pages — timeouts / 907c553db6e6 / 8

- [Property reference](resources--bgp_routing_policy--reference--group-001.md#canonical-c72ce71fac4abd1b7af13fe660e6b9827ab44900fa751595f774bec024eaacf2)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-a52b2ec832e07dcfd0c3d20b81b2af1ed61314bb6d9237e5a98ee665cc2f23d5)
