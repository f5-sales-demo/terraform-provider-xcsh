---
page_title: "xcsh_bgp_routing_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy reference."
---

# xcsh_bgp_routing_policy reference

<a id="canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8ed183f664e7dd2b8981cad2770db79bb811d9108fc933d31147fc000e4c7d1"></a>

## Property reference — Property reference / d94cd5d65739 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- Property reference

<a id="canonical-c7a02da7a1257883f6273c0a0c92f87458748011bc945dd98017eafed589b10c"></a>

## Direct properties — Property reference / d94cd5d65739 / 3

<a id="canonical-253e34656549593e0d29c22c26488cb6ab15f935a7cf514130060980c4cd0d1b"></a>

<a id="canonical-409c72da1904dccdb062c88d3e7dd4f1dab7591772a8c1322e1a11e171d2ec80"></a>

## annotations property — Property reference / d94cd5d65739 / 4

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

<a id="canonical-10bbe594d50d060ea9c8b1d51ff943466f2722d17c95ae5af823e7627d8fa07d"></a>

<a id="canonical-ffae414c9b08dad860659fb32e3dc9be57cb6ce233626fe159b4bc775f3ee9db"></a>

## description property — Property reference / d94cd5d65739 / 5

Type: `"string"`. Computed.

Description of the BGPRoutingPolicy.

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

<a id="canonical-36a9d8d8f7eb7f0aa6e910674f7ada16eae39007549c6a7bcb7668086ad8526e"></a>

<a id="canonical-f40e356a25a430169fc71164e9f56521c4f0ad9dd95f4badc781f5053170b457"></a>

## id property — Property reference / d94cd5d65739 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-335fb68da4d6a4cde6f192dc897ccd70c372e1d355ae3123820913b6ee8551e0"></a>

<a id="canonical-15085faf937911028ff1f2da12fd7710efc86bed296f4bb968a317d1142ee49c"></a>

## labels property — Property reference / d94cd5d65739 / 7

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

<a id="canonical-3a4b846bd5a4d65b66871f7c608f2977c9c6b4efe3d8423fd4db76b2df65bdb9"></a>

<a id="canonical-78d7b16c1b0749d6a28d1fefb6b06688096fd0b800e29c2f7ce4075f76913958"></a>

## name property — Property reference / d94cd5d65739 / 8

Type: `"string"`. Required.

Name of the BGPRoutingPolicy.

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

<a id="canonical-696911af45dcaa7afcc431791e2a1152369c1da65b52c177e5cac3edca926ef4"></a>

<a id="canonical-b77d735a63b924724fe4136d8663e4a30355205f65b4c63eee0d4f8b491ae4f2"></a>

## namespace property — Property reference / d94cd5d65739 / 9

Type: `"string"`. Required.

Namespace where the BGPRoutingPolicy exists.

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

- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951): complete subsection reference.

<a id="canonical-bc6fe7bf94ec165d7ff83bbc61cfeb3fc9e6184e30c1aa54b32af436c049e803"></a>

## All schema paths — Property reference / d94cd5d65739 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp_routing_policy--reference--group-001.md#canonical-253e34656549593e0d29c22c26488cb6ab15f935a7cf514130060980c4cd0d1b) |
| `description` | [description](data-sources--bgp_routing_policy--reference--group-001.md#canonical-10bbe594d50d060ea9c8b1d51ff943466f2722d17c95ae5af823e7627d8fa07d) |
| `id` | [id](data-sources--bgp_routing_policy--reference--group-001.md#canonical-36a9d8d8f7eb7f0aa6e910674f7ada16eae39007549c6a7bcb7668086ad8526e) |
| `labels` | [labels](data-sources--bgp_routing_policy--reference--group-001.md#canonical-335fb68da4d6a4cde6f192dc897ccd70c372e1d355ae3123820913b6ee8551e0) |
| `name` | [name](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3a4b846bd5a4d65b66871f7c608f2977c9c6b4efe3d8423fd4db76b2df65bdb9) |
| `namespace` | [namespace](data-sources--bgp_routing_policy--reference--group-001.md#canonical-696911af45dcaa7afcc431791e2a1152369c1da65b52c177e5cac3edca926ef4) |
| `rules` | [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-4fd36201b3a45f6143ef759bc82261e15f288984b7fe361fa70966f6335fdbc9) |
| `rules.action` | [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-d38270c336efb3dfc71cc9fd85e9b72b9c590aa37fb410bb713a9d9b92dd0f91) |
| `rules.action.allow` | [rules.action.allow](data-sources--bgp_routing_policy--reference--group-001.md#canonical-15aa16910d27fd182ef9a3076eba1c4148e90b00e276441369c621cdb9712381) |
| `rules.action.as_path` | [rules.action.as_path](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3bd14a7ea7e333157733e16b62e0ac7d31b4cea71d1bc8e0cbc15181ac7405ba) |
| `rules.action.community` | [rules.action.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e0a6a1ef9568fe6bd09f4e4c79f34aacf906742270932266a8f396979a46521e) |
| `rules.action.community.community` | [rules.action.community.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-5b8a207d192c84c9eb80353bd95b04ef84c7c7630c74f22ef18fa16c7605c3f2) |
| `rules.action.deny` | [rules.action.deny](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7e0725ed38303ca3b839582da45b6f9bff43ab16508be4a3b16e0b0d10b6515d) |
| `rules.action.local_preference` | [rules.action.local_preference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-5e8aadee288969ee08517a66aa8f1de0b70630fb76ae016915c2b131008dbc54) |
| `rules.action.metric` | [rules.action.metric](data-sources--bgp_routing_policy--reference--group-001.md#canonical-30102e2c078f8bc7359e7c0b2b792d456fa27271649e4595c5a2d5667da0ee83) |
| `rules.match` | [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1b7c19545b97da8e5796a0a1fed8907b797809f815f44eecef0af798aea66c5f) |
| `rules.match.as_path` | [rules.match.as_path](data-sources--bgp_routing_policy--reference--group-001.md#canonical-04d572ff39d27588f8291a9a8d1c438ad4796db196e3abcb353d3f070f08c0c5) |
| `rules.match.community` | [rules.match.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-439c86bf83baba9da99f7ce4cb95d872ffdee04e5da61e54a32c46f523b2fb3a) |
| `rules.match.community.community` | [rules.match.community.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e4370937430457ae9c2b40c0113c7694281e6329d14883c6e04018bc9179ee28) |
| `rules.match.ip_prefixes` | [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-d193d8be7cf856b2536e5adb8c77605d6faf5b2cac56705bdd62121123260c47) |
| `rules.match.ip_prefixes.prefixes` | [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-01b62943be263e1cff6593eb5c0bdd5eb3afea5817c8e06b5a20104da394e622) |
| `rules.match.ip_prefixes.prefixes.equal_or_longer_than` | [rules.match.ip_prefixes.prefixes.equal_or_longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-348837fd5e7ee9ab29c3137ac1cbcaa39f5bb58ef074a6806b8afcddee202dce) |
| `rules.match.ip_prefixes.prefixes.exact_match` | [rules.match.ip_prefixes.prefixes.exact_match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8f8db897da562eacabf5ed43ddf4d113e0ac9611efeadcc9600645d7765e1c32) |
| `rules.match.ip_prefixes.prefixes.ip_prefixes` | [rules.match.ip_prefixes.prefixes.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ec9644c8b530611e486546ce5aa4fed2fd256354e0af68baa0462c445b343cb3) |
| `rules.match.ip_prefixes.prefixes.longer_than` | [rules.match.ip_prefixes.prefixes.longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ed89b5b23a518ae48574dbc6aa4019fa3bfb18ec39c1438f985218be9d1dc11f) |

<a id="canonical-a37dbc48d1ef975cf2b569d20935433d085013ad04b02831723913d32e4ba220"></a>

## Next pages — Property reference / d94cd5d65739 / 11

- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c94c8c17341f02b89e6fa310b7338561597b41f28ea99823ca8a46fdb4177a97"></a>

## rules — rules / 5fd31b5cfc31 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- rules

<a id="canonical-4fd36201b3a45f6143ef759bc82261e15f288984b7fe361fa70966f6335fdbc9"></a>

Type: `"list"`. Computed.

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

<a id="canonical-5b194611a053e1409182717790795e07c5743e196adb873addaaaa0849e3be5e"></a>

## Direct properties — rules / 5fd31b5cfc31 / 3

- [action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79): complete subsection reference.

- [match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486): complete subsection reference.

<a id="canonical-a289bab3f9d1e47b989166c9d154a29eadcf3f466e3cae4fe6f19e7169774715"></a>

## Next pages — rules / 5fd31b5cfc31 / 4

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9a1de9dac5f66c4a01f7abf5e4ebd970d1e1e3f0d40e2afeeb5ee805e3e9721"></a>

## rules.action — rules.action / 4bb9d09a586c / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- rules.action

<a id="canonical-d38270c336efb3dfc71cc9fd85e9b72b9c590aa37fb410bb713a9d9b92dd0f91"></a>

Type: `"single"`. Computed.

Action to be enforced if the BGP route matches the rule.

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

<a id="canonical-a1bc1339f5291089b06f8a61e0efde82624c2df46cfd2477a0c4cc82b49e206a"></a>

## Direct properties — rules.action / 4bb9d09a586c / 3

- [allow](data-sources--bgp_routing_policy--reference--group-001.md#canonical-55b079c2bbb0436d6a2351a6914fc21b07ce30872fb43efe44cfc6734588b51b): complete subsection reference.

<a id="canonical-3bd14a7ea7e333157733e16b62e0ac7d31b4cea71d1bc8e0cbc15181ac7405ba"></a>

<a id="canonical-cc6f36caa982dd221c56c9bcdd33a454095c89102806c49e288fe32707820f5a"></a>

## as_path property — rules.action / 4bb9d09a586c / 4

Type: `"string"`. Computed.

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

- [community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0f5390329ecf19897ef44bb8d1f967753ae7dfe1791019dfeac2f42f68e94a4d): complete subsection reference.

- [deny](data-sources--bgp_routing_policy--reference--group-001.md#canonical-762fd518bfafef5cd0c0e3dc8758121d4f7de322303d7f72de07a78856841e91): complete subsection reference.

<a id="canonical-5e8aadee288969ee08517a66aa8f1de0b70630fb76ae016915c2b131008dbc54"></a>

<a id="canonical-acd4a7dfaac97a73d0e6c2563b875565c5ebabd40ad9df8b364cf2be62209576"></a>

## local_preference property — rules.action / 4bb9d09a586c / 5

Type: `"number"`. Computed.

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

<a id="canonical-30102e2c078f8bc7359e7c0b2b792d456fa27271649e4595c5a2d5667da0ee83"></a>

<a id="canonical-f2c7e8257a3fc2d4e676217a9234ba9c4c5200cfbbb139c795c77a9b7d5bc4e9"></a>

## metric property — rules.action / 4bb9d09a586c / 6

Type: `"number"`. Computed.

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

<a id="canonical-0ae046e6b9eeb2fd61a159531d7c534d08c3f12bad99cc78a4b41d6acaa4008b"></a>

## Next pages — rules.action / 4bb9d09a586c / 7

- [rules.action.allow](data-sources--bgp_routing_policy--reference--group-001.md#canonical-55b079c2bbb0436d6a2351a6914fc21b07ce30872fb43efe44cfc6734588b51b)
- [rules.action.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0f5390329ecf19897ef44bb8d1f967753ae7dfe1791019dfeac2f42f68e94a4d)
- [rules.action.deny](data-sources--bgp_routing_policy--reference--group-001.md#canonical-762fd518bfafef5cd0c0e3dc8758121d4f7de322303d7f72de07a78856841e91)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-55b079c2bbb0436d6a2351a6914fc21b07ce30872fb43efe44cfc6734588b51b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53155cf5f4ea2985a04bbdb451a1987b2867ed7d8078d3363ab4e0568bbc2ae4"></a>

## rules.action.allow — rules.action.allow / 8178f7e29db9 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79)
- rules.action.allow

<a id="canonical-15aa16910d27fd182ef9a3076eba1c4148e90b00e276441369c621cdb9712381"></a>

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

<a id="canonical-989e6de2efafea6effc508f6642c7f894f3d45ef0ddd0f78ef6849a97c8876f9"></a>

## Direct properties — rules.action.allow / 8178f7e29db9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb807657fd215f6fb0376f88899005321aed9fe0492b5bffecbe807015b59ba2"></a>

## Next pages — rules.action.allow / 8178f7e29db9 / 4

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-0f5390329ecf19897ef44bb8d1f967753ae7dfe1791019dfeac2f42f68e94a4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b0c5fb5b05603c691a0abf908ce1e0d28c9e62c420d0682f176ac921b9ab61b"></a>

## rules.action.community — rules.action.community / 828ff1c4c510 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79)
- rules.action.community

<a id="canonical-e0a6a1ef9568fe6bd09f4e4c79f34aacf906742270932266a8f396979a46521e"></a>

Type: `"single"`. Computed.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

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

<a id="canonical-3aef7b3d3557600e63b855167732c6b3883ecf9dff56291e5e9806459f9d9de8"></a>

## Direct properties — rules.action.community / 828ff1c4c510 / 3

<a id="canonical-5b8a207d192c84c9eb80353bd95b04ef84c7c7630c74f22ef18fa16c7605c3f2"></a>

<a id="canonical-ec66b00c5d6bd74d514e98f2937e4f05feafd9b85c14e2ad6bafdf1d73a9c241"></a>

## community property — rules.action.community / 828ff1c4c510 / 4

Type: `["list", "string"]`. Computed.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

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

<a id="canonical-c19ba15b733ea5285fe50f64592d691fbbb45758d1b2248365577ac2cdb45cc1"></a>

## Next pages — rules.action.community / 828ff1c4c510 / 5

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-762fd518bfafef5cd0c0e3dc8758121d4f7de322303d7f72de07a78856841e91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79d1c609ced2b392d4db7d0ebb5ec776e00c90ff5791d5defdf738f87d5bf76c"></a>

## rules.action.deny — rules.action.deny / 907ad414313c / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79)
- rules.action.deny

<a id="canonical-7e0725ed38303ca3b839582da45b6f9bff43ab16508be4a3b16e0b0d10b6515d"></a>

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

<a id="canonical-2803e83ea42167eca061d527c7fc7a1c1ea9d31dc2b20c8c66ec468130966ec3"></a>

## Direct properties — rules.action.deny / 907ad414313c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f59bbf7204e070ac04ced396bf38b3e4f5a3346cce35d908c74a9b9322fb8a6d"></a>

## Next pages — rules.action.deny / 907ad414313c / 4

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-e712e3deaf698e92a1dcc8c03fee9a466f2c60fc36d83e1e78f19c98945ded79)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86b4baf63ee6957369733a5cf67b01bd0cd4909fe53522adf691dc88bfeb19af"></a>

## rules.match — rules.match / 563ddeea0972 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- rules.match

<a id="canonical-1b7c19545b97da8e5796a0a1fed8907b797809f815f44eecef0af798aea66c5f"></a>

Type: `"single"`. Computed.

Predicates which have to match information in route for action to be applied.

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

<a id="canonical-5756498cc870b25760b93d36635d3dbe7a199c0d7883a3c8b4a54a06c71c8098"></a>

## Direct properties — rules.match / 563ddeea0972 / 3

<a id="canonical-04d572ff39d27588f8291a9a8d1c438ad4796db196e3abcb353d3f070f08c0c5"></a>

<a id="canonical-a49227e2c0df4a8c463b1498eba32dbce21444d01da1b4008549712d1e4feaf4"></a>

## as_path property — rules.match / 563ddeea0972 / 4

Type: `"string"`. Computed.

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

- [community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0189d5b9673209efa9b4a7c9f3a53f50e3afe8f11b6e06ec537e643167d8f93a): complete subsection reference.

- [ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb): complete subsection reference.

<a id="canonical-bbb43cde489c76d6bd9ab7af835d9ae39161bedd76e384bf01346f8ffabbf66a"></a>

## Next pages — rules.match / 563ddeea0972 / 5

- [rules.match.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0189d5b9673209efa9b4a7c9f3a53f50e3afe8f11b6e06ec537e643167d8f93a)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-0189d5b9673209efa9b4a7c9f3a53f50e3afe8f11b6e06ec537e643167d8f93a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdaa7f8f0a87c13387e7e1bd695d190c152319f931c189ce82fd8bfedc9ecdfc"></a>

## rules.match.community — rules.match.community / 87bda63111a7 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- rules.match.community

<a id="canonical-439c86bf83baba9da99f7ce4cb95d872ffdee04e5da61e54a32c46f523b2fb3a"></a>

Type: `"single"`. Computed.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

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

<a id="canonical-04cf841d3c3a715e76dd2974e4059e8408289c61d5b2f6995d797a459cfaf134"></a>

## Direct properties — rules.match.community / 87bda63111a7 / 3

<a id="canonical-e4370937430457ae9c2b40c0113c7694281e6329d14883c6e04018bc9179ee28"></a>

<a id="canonical-4e733ff9f6e6bc20e9704d7e7fe1dd77203612a36aed7aa4547be859e9fcad7a"></a>

## community property — rules.match.community / 87bda63111a7 / 4

Type: `["list", "string"]`. Computed.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

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

<a id="canonical-6c6da32f6f70f0aa1c88e806ede9e62d6773131fec2f1aa781b708aad4b3c9ec"></a>

## Next pages — rules.match.community / 87bda63111a7 / 5

- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6de711c7f7a14b2e19dd95b245bc7197320fa3e64856fff577462cdce62dba81"></a>

## rules.match.ip_prefixes — rules.match.ip_prefixes / 501ba286fc02 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- rules.match.ip_prefixes

<a id="canonical-d193d8be7cf856b2536e5adb8c77605d6faf5b2cac56705bdd62121123260c47"></a>

Type: `"single"`. Computed.

List of IP prefix and prefix length range match condition.

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

<a id="canonical-17ee6efb6abb0236470728963e77f8b77b5786c0f2d5f74d62f771948e897734"></a>

## Direct properties — rules.match.ip_prefixes / 501ba286fc02 / 3

- [prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080): complete subsection reference.

<a id="canonical-952defec12b9ee8993fa5153eaef2adf04f071324cbda59928a6ebcc07fbf41e"></a>

## Next pages — rules.match.ip_prefixes / 501ba286fc02 / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce6a1ca530fbf99e7ac55e3651717887f40ff2011b09007cf6c1acb874747994"></a>

## rules.match.ip_prefixes.prefixes — rules.match.ip_prefixes.prefixes / 68f688b78781 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb)
- rules.match.ip_prefixes.prefixes

<a id="canonical-01b62943be263e1cff6593eb5c0bdd5eb3afea5817c8e06b5a20104da394e622"></a>

Type: `"list"`. Computed.

Prefix list. List of IP prefix.

Upstream description:

List of IP prefix.

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

<a id="canonical-7e19810e029cdb97196cc323e21ce3414dd668d4500a4f2fff5de77aea5c7f9f"></a>

## Direct properties — rules.match.ip_prefixes.prefixes / 68f688b78781 / 3

- [equal_or_longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-db4ef830b8dc475b6796ae3eeb857c387228fe2a543b750225cbbf2767580bd2): complete subsection reference.

- [exact_match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-9db612d7026c309c2fd99ef7d0984de4728bf1633b08296460fa6dffd372021c): complete subsection reference.

<a id="canonical-ec9644c8b530611e486546ce5aa4fed2fd256354e0af68baa0462c445b343cb3"></a>

<a id="canonical-64d6fb4066d482c84b086ea2e980c68ace64807879296f16c72bc350a473b871"></a>

## ip_prefixes property — rules.match.ip_prefixes.prefixes / 68f688b78781 / 4

Type: `"string"`. Computed.

IP Prefix. IP prefix to match on BGP route.

Upstream description:

IP prefix to match on BGP route.

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

- [longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-385fc56f8802bcd42e3c848f539e7da58e332efa90b882034f0a1510ca297b7f): complete subsection reference.

<a id="canonical-c5805ae519fa02dc15eca9d718067f02bc731ae402b4f981ccd646e9854a7e98"></a>

## Next pages — rules.match.ip_prefixes.prefixes / 68f688b78781 / 5

- [rules.match.ip_prefixes.prefixes.equal_or_longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-db4ef830b8dc475b6796ae3eeb857c387228fe2a543b750225cbbf2767580bd2)
- [rules.match.ip_prefixes.prefixes.exact_match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-9db612d7026c309c2fd99ef7d0984de4728bf1633b08296460fa6dffd372021c)
- [rules.match.ip_prefixes.prefixes.longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-385fc56f8802bcd42e3c848f539e7da58e332efa90b882034f0a1510ca297b7f)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-db4ef830b8dc475b6796ae3eeb857c387228fe2a543b750225cbbf2767580bd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8477b53bf432acf106d417a86f45dd14479114770b324b60a24d6a748a148aa9"></a>

## rules.match.ip_prefixes.prefixes.equal_or_longer_than — rules.match.ip_prefixes.prefixes.equal_or_longer_than / daad3a9d0115 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb)
- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080)
- rules.match.ip_prefixes.prefixes.equal_or_longer_than

<a id="canonical-348837fd5e7ee9ab29c3137ac1cbcaa39f5bb58ef074a6806b8afcddee202dce"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-f9e0a51d7d7510a894e59d898e5af4a3fe9d07cd05e5f1d93ab8d9a3aed21625"></a>

## Direct properties — rules.match.ip_prefixes.prefixes.equal_or_longer_than / daad3a9d0115 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de798aa7d6d80f15004d6e407dd274af40a9fbfaa480a719bc9b25fe910eb06c"></a>

## Next pages — rules.match.ip_prefixes.prefixes.equal_or_longer_than / daad3a9d0115 / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-9db612d7026c309c2fd99ef7d0984de4728bf1633b08296460fa6dffd372021c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-169e7362a222df7cf9b2ebc115558e05fc4f49f496bd06d94e31fc7c0c92ba64"></a>

## rules.match.ip_prefixes.prefixes.exact_match — rules.match.ip_prefixes.prefixes.exact_match / 538225dd35ab / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb)
- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080)
- rules.match.ip_prefixes.prefixes.exact_match

<a id="canonical-8f8db897da562eacabf5ed43ddf4d113e0ac9611efeadcc9600645d7765e1c32"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-fac8a314f79aa00a3ff1a795feda0f49743109b286f892d60f1c4613685fae78"></a>

## Direct properties — rules.match.ip_prefixes.prefixes.exact_match / 538225dd35ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f75b380c1e7052f077ca484ef94230534c6235490a4fd818d973aafafe5cba33"></a>

## Next pages — rules.match.ip_prefixes.prefixes.exact_match / 538225dd35ab / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-385fc56f8802bcd42e3c848f539e7da58e332efa90b882034f0a1510ca297b7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ec0e6499aaf0383129f86e628c803efdcbd67384847518326eff49c89ce8f99"></a>

## rules.match.ip_prefixes.prefixes.longer_than — rules.match.ip_prefixes.prefixes.longer_than / d3a35f6b0471 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-eccae3ab92e4281f4d8846e4b097e7e7a2471f1c57d3a01c64e029bb315272fb)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-8645fddcb320ea5df9f71f29e53a090de381987c451840d9a46964b40702e951)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-7a468153371fe867bb031bcbec617872c1ee69358c861d7b6d07b1073d029486)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-ee6387786956d03d8f36008aad6e2771cb9a4e96dd88c8939539526e0641e4fb)
- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080)
- rules.match.ip_prefixes.prefixes.longer_than

<a id="canonical-ed89b5b23a518ae48574dbc6aa4019fa3bfb18ec39c1438f985218be9d1dc11f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-86e70da9531cc372412b8849cf5307eee4141bcac931f0eb4645c15b88add88e"></a>

## Direct properties — rules.match.ip_prefixes.prefixes.longer_than / d3a35f6b0471 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e3479a9850c3724fa8e2297d6d1b5d3e033a74bbcceaacd0522935eb8d86d1f"></a>

## Next pages — rules.match.ip_prefixes.prefixes.longer_than / d3a35f6b0471 / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-b2adceaddb8027e1a43db5cffbc0f1bbabad9882b3097e17a535005d13841080)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
