---
page_title: "xcsh_ip_prefix_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set reference."
---

# xcsh_ip_prefix_set reference

<a id="canonical-24ad13b3f3a6891fd4f419a489a2eede859b177520a1a420539bf19661644552"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ac47e07d143e724c06c3d82890444c2eca5b096955ca5a233cba36eb2d3fdde"></a>

## Property reference — Property reference / f0297699888f / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)
- Property reference

<a id="canonical-464d970ca98ed664af7af6203f420d00569fb9d85fb3783fcee21fb325ed0e3e"></a>

## Direct properties — Property reference / f0297699888f / 3

<a id="canonical-0a1eabbad264ec24cb3838ce8baf20a1cb13cfa86bcaaf63766f816158a802e6"></a>

<a id="canonical-94e97cd2034059cfb13018ae08a6376a41c6398a670f74cee403e760b744bd62"></a>

## annotations property — Property reference / f0297699888f / 4

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

<a id="canonical-ceb2babd5a3275485ecdb80db48610e7bab6709a96c397dc358992b55ad22ee6"></a>

<a id="canonical-71c96f8c259ea1459930e5637adfeebdfac5792a4f4e22fb13234213f17debb4"></a>

## description property — Property reference / f0297699888f / 5

Type: `"string"`. Computed.

Description of the IPPrefixSet.

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

<a id="canonical-f8075cd93c39c4ffba5bedf6be4736fff2f131ff1dd959294b00e90c3856607f"></a>

<a id="canonical-f40a2faf5b6ee52c3e96b8ab59e977741b95e9a040f6e62a4406c8c90d9688fb"></a>

## id property — Property reference / f0297699888f / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4_prefixes](data-sources--ip_prefix_set--reference--group-001.md#canonical-ac891904bd02133e39754853dba518767b461eba0b28ea7580dafaf21959af6f): complete subsection reference.

<a id="canonical-ee186fd6269b20c5d8fec66af98f31d3718e15ca7536dc8c8148a9c9f5e31ca9"></a>

<a id="canonical-1e9621b9f999792b5d7b740569d2d4520201a0cc6313e78eb3be441d968b6b26"></a>

## labels property — Property reference / f0297699888f / 7

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

<a id="canonical-290123e12e698ae2b873dfc4536952ffcb31a4043d6d8b8cce7d2215d48654d4"></a>

<a id="canonical-3a6c3cb2ef30dc201e87129dbaac2c1fde181ccbac19f715a828eea669fc2a1d"></a>

## name property — Property reference / f0297699888f / 8

Type: `"string"`. Required.

Name of the IPPrefixSet.

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

<a id="canonical-10bd176a9330e9479aa97458ccfa70672fd304fbaa040c95ab591120e4d9938a"></a>

<a id="canonical-1b880a5b14fe0ada02f24b1e078175060921d66a96290efaf686724796923856"></a>

## namespace property — Property reference / f0297699888f / 9

Type: `"string"`. Required.

Namespace where the IPPrefixSet exists.

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

<a id="canonical-16c13ab49d419124b884485cbf8a3fe7cf82fa5c3a4eb87f979eed46661da716"></a>

## All schema paths — Property reference / f0297699888f / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ip_prefix_set--reference--group-001.md#canonical-0a1eabbad264ec24cb3838ce8baf20a1cb13cfa86bcaaf63766f816158a802e6) |
| `description` | [description](data-sources--ip_prefix_set--reference--group-001.md#canonical-ceb2babd5a3275485ecdb80db48610e7bab6709a96c397dc358992b55ad22ee6) |
| `id` | [id](data-sources--ip_prefix_set--reference--group-001.md#canonical-f8075cd93c39c4ffba5bedf6be4736fff2f131ff1dd959294b00e90c3856607f) |
| `ipv4_prefixes` | [ipv4_prefixes](data-sources--ip_prefix_set--reference--group-001.md#canonical-f9226f88eab3e4f06beefacc827a05360feee0fc0f4b9e2b40ca1acc42c50749) |
| `ipv4_prefixes.description_spec` | [ipv4_prefixes.description_spec](data-sources--ip_prefix_set--reference--group-001.md#canonical-a326497d444478590c7ac0e96d5d0206435e959cc20292adbc51901ec36e0341) |
| `ipv4_prefixes.ipv4_prefix` | [ipv4_prefixes.ipv4_prefix](data-sources--ip_prefix_set--reference--group-001.md#canonical-a3cf402057bb02fa5fc65b3d18a17b35213ebd5176b5da7c2621f7c672607826) |
| `labels` | [labels](data-sources--ip_prefix_set--reference--group-001.md#canonical-ee186fd6269b20c5d8fec66af98f31d3718e15ca7536dc8c8148a9c9f5e31ca9) |
| `name` | [name](data-sources--ip_prefix_set--reference--group-001.md#canonical-290123e12e698ae2b873dfc4536952ffcb31a4043d6d8b8cce7d2215d48654d4) |
| `namespace` | [namespace](data-sources--ip_prefix_set--reference--group-001.md#canonical-10bd176a9330e9479aa97458ccfa70672fd304fbaa040c95ab591120e4d9938a) |

<a id="canonical-44697071e788b9e2d3e39f2f7449fa3cc240dccc64a48591930ded13c71a42c9"></a>

## Next pages — Property reference / f0297699888f / 11

- [ipv4_prefixes](data-sources--ip_prefix_set--reference--group-001.md#canonical-ac891904bd02133e39754853dba518767b461eba0b28ea7580dafaf21959af6f)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)

<a id="canonical-ac891904bd02133e39754853dba518767b461eba0b28ea7580dafaf21959af6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b30af5b42c3db69ea0c94f8e3f518a331fc0315727d46482fd45a4e041123226"></a>

## ipv4_prefixes — ipv4_prefixes / f0573c96274c / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)
- [Property reference](data-sources--ip_prefix_set--reference--group-001.md#canonical-24ad13b3f3a6891fd4f419a489a2eede859b177520a1a420539bf19661644552)
- ipv4_prefixes

<a id="canonical-f9226f88eab3e4f06beefacc827a05360feee0fc0f4b9e2b40ca1acc42c50749"></a>

Type: `"list"`. Computed.

IPv4 Prefixes. List of IPv4 prefixes with description.

Upstream description:

List of IPv4 prefixes with description.

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

<a id="canonical-ae4ae84d2535293459476aaae28879f1b90bec803c0636a0361145091677613e"></a>

## Direct properties — ipv4_prefixes / f0573c96274c / 3

<a id="canonical-a326497d444478590c7ac0e96d5d0206435e959cc20292adbc51901ec36e0341"></a>

<a id="canonical-2744f943b8c8d070e01520f4ad54debf9f19327d01d8d43911d976f98e5362c2"></a>

## description_spec property — ipv4_prefixes / f0573c96274c / 4

Type: `"string"`. Computed.

Description. Human-readable description text

<a id="canonical-a3cf402057bb02fa5fc65b3d18a17b35213ebd5176b5da7c2621f7c672607826"></a>

<a id="canonical-80e791062e7821967765606bd8bbbcd9fd4816f668864d5e1542054593f204ff"></a>

## ipv4_prefix property — ipv4_prefixes / f0573c96274c / 5

Type: `"string"`. Computed.

IPv4 Prefix. IP address configuration

Upstream description:

IP address configuration

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-e71ba7c7a8f0144678ea663ba6e214312d93711b53aa26afa82558e7150d83e0"></a>

## Next pages — ipv4_prefixes / f0573c96274c / 6

- [Property reference](data-sources--ip_prefix_set--reference--group-001.md#canonical-24ad13b3f3a6891fd4f419a489a2eede859b177520a1a420539bf19661644552)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)
