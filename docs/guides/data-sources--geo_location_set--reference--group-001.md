---
page_title: "xcsh_geo_location_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set reference."
---

# xcsh_geo_location_set reference

<a id="canonical-5bf5901a30d1ac937b093d57aa613f518f7da14cd2317e1a7f4c81cd568647fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe6f22996bb3f944d690176acd919d59e7eec34ecf270f584656ee184b718faf"></a>

## Property reference — Property reference / 84a1df121946 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)
- Property reference

<a id="canonical-f0d0f51ec899c6b33c496a1030c5c10a6183e1e0cd88c2234429e449bb3f20a3"></a>

## Direct properties — Property reference / 84a1df121946 / 3

<a id="canonical-6b601063678e8b960f6afa8d33086ad1469ff7c5a6e98752e9ece8380980c554"></a>

<a id="canonical-7d1652e0a587b00b4cad30c6a2efc213195bd1fbc1a4ee9326033e4b16bd8fb6"></a>

## annotations property — Property reference / 84a1df121946 / 4

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

- [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-88f90956cf984341df2a68d96303f539cec517f172e64f3a259196b932f70fb5): complete subsection reference.

<a id="canonical-d0d3bbc14bcb467766e5126a6dc5d253ab42fce6820e9c3cdc311dcef06c4c54"></a>

<a id="canonical-9f17ce3a99c117cce31712d144797b32b7ee236417e2ef18ef8a17bd9ada843b"></a>

## description property — Property reference / 84a1df121946 / 5

Type: `"string"`. Computed.

Description of the GeoLocationSet.

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

- [global](data-sources--geo_location_set--reference--group-001.md#canonical-3b825e0f915ecd1aa2bcd31e0b64b5c30f5df72ceaaa37ba535ea1999bd96cfb): complete subsection reference.

<a id="canonical-9ea96208a246a631b32179b3a2090fe8f9359ce4ee8f792701d4041b5bb3f90b"></a>

<a id="canonical-a96336f5a0a97158736cadcce4cb15dabfebe6cdef278ec87493b66a847b17ca"></a>

## id property — Property reference / 84a1df121946 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6320ce318c674fc9cb549ebd451d0fa94d627006d1c7b87d83e8a0bb94615510"></a>

<a id="canonical-632683c6ac1a6a8d2a9a3533d42abce691e96b55d696e2312541a7f2776bc7f6"></a>

## labels property — Property reference / 84a1df121946 / 7

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

<a id="canonical-64334c77ba366dbe4a356b4ac4a9d9cd44c56568fd34eae7a1c6311b230d6901"></a>

<a id="canonical-837cf99da9c77516e5ffeb188c2197282f4feb6a0f445a21718d63f3cb9518b0"></a>

## name property — Property reference / 84a1df121946 / 8

Type: `"string"`. Required.

Name of the GeoLocationSet.

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

<a id="canonical-8e7ddc575e9cb1493ba9533e4ddc4a151662cbdef4766be9253239896aef95b9"></a>

<a id="canonical-baa25f963781f5ba021545c39748defda4bf93eaa6694949f66c3a55d2091de5"></a>

## namespace property — Property reference / 84a1df121946 / 9

Type: `"string"`. Optional, Computed.

Namespace where the GeoLocationSet exists.

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

<a id="canonical-9f849c7e75c7b55462ec5ee026f5c879c4baca905782c501dfd20a36eeed82cc"></a>

## All schema paths — Property reference / 84a1df121946 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--geo_location_set--reference--group-001.md#canonical-6b601063678e8b960f6afa8d33086ad1469ff7c5a6e98752e9ece8380980c554) |
| `custom_geo_location_selector` | [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-ba4e9706cf3b77bb8ba168970416a45a74ddab3ce575015fa2b0b7574d3e7be7) |
| `custom_geo_location_selector.expressions` | [custom_geo_location_selector.expressions](data-sources--geo_location_set--reference--group-001.md#canonical-f13491fb2c9d891e2aed94e09365c277703bda621bf87752d66b38cb81d6c6eb) |
| `description` | [description](data-sources--geo_location_set--reference--group-001.md#canonical-d0d3bbc14bcb467766e5126a6dc5d253ab42fce6820e9c3cdc311dcef06c4c54) |
| `global` | [global](data-sources--geo_location_set--reference--group-001.md#canonical-0f8ae06f0526d462a2f4af44817f2d9cc6ff9d6a4c5b6133aecff50da5c1ef51) |
| `id` | [id](data-sources--geo_location_set--reference--group-001.md#canonical-9ea96208a246a631b32179b3a2090fe8f9359ce4ee8f792701d4041b5bb3f90b) |
| `labels` | [labels](data-sources--geo_location_set--reference--group-001.md#canonical-6320ce318c674fc9cb549ebd451d0fa94d627006d1c7b87d83e8a0bb94615510) |
| `name` | [name](data-sources--geo_location_set--reference--group-001.md#canonical-64334c77ba366dbe4a356b4ac4a9d9cd44c56568fd34eae7a1c6311b230d6901) |
| `namespace` | [namespace](data-sources--geo_location_set--reference--group-001.md#canonical-8e7ddc575e9cb1493ba9533e4ddc4a151662cbdef4766be9253239896aef95b9) |

<a id="canonical-a7d570054b6cca6d60bf599016ab35c0ff2cf34419a49c29eb45b3f4bd2290b2"></a>

## Next pages — Property reference / 84a1df121946 / 11

- [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-88f90956cf984341df2a68d96303f539cec517f172e64f3a259196b932f70fb5)
- [global](data-sources--geo_location_set--reference--group-001.md#canonical-3b825e0f915ecd1aa2bcd31e0b64b5c30f5df72ceaaa37ba535ea1999bd96cfb)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)

<a id="canonical-88f90956cf984341df2a68d96303f539cec517f172e64f3a259196b932f70fb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88e336772b36b01dbf66942ee082da625c4a24f25570a3ac8ed32dede37a9f6d"></a>

## custom_geo_location_selector — custom_geo_location_selector / 23e01a7f62dc / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)
- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-5bf5901a30d1ac937b093d57aa613f518f7da14cd2317e1a7f4c81cd568647fc)
- custom_geo_location_selector

<a id="canonical-ba4e9706cf3b77bb8ba168970416a45a74ddab3ce575015fa2b0b7574d3e7be7"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_geo\_location\_selector, global\] Type can be used to establish a 'selector
reference' from one object(called selector) to a set of other objects(called selectees) based on the
value of expressions. A label selector is a label query over a set of resources. An empty label
selector matches all objects.

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

OneOf alternatives in this subsection:

- [custom_geo_location_selector](data-sources--geo_location_set--reference--group-001.md#canonical-ba4e9706cf3b77bb8ba168970416a45a74ddab3ce575015fa2b0b7574d3e7be7)
- [global](data-sources--geo_location_set--reference--group-001.md#canonical-0f8ae06f0526d462a2f4af44817f2d9cc6ff9d6a4c5b6133aecff50da5c1ef51)

Select alternatives according to the provider validators above.

<a id="canonical-8f756ea448d96f0480728067c33da245cdc24eb6ff7dd1a6742dd42fe85f8584"></a>

## Direct properties — custom_geo_location_selector / 23e01a7f62dc / 3

<a id="canonical-f13491fb2c9d891e2aed94e09365c277703bda621bf87752d66b38cb81d6c6eb"></a>

<a id="canonical-c8d6d0c8e0e6e682877e8156bae7550ee715a259ad9568e1d109abb627a43bea"></a>

## expressions property — custom_geo_location_selector / 23e01a7f62dc / 4

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

<a id="canonical-98231172e1c3416858f43b84d9d9584d9330d70ff24152bda220d5c884ccec9b"></a>

## Next pages — custom_geo_location_selector / 23e01a7f62dc / 5

- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-5bf5901a30d1ac937b093d57aa613f518f7da14cd2317e1a7f4c81cd568647fc)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)

<a id="canonical-3b825e0f915ecd1aa2bcd31e0b64b5c30f5df72ceaaa37ba535ea1999bd96cfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f276149925ebef4586f1951e8c00213153dbbb25896fbf5e463da8eecd5f1fa"></a>

## global — global / 4043edde7c7b / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)
- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-5bf5901a30d1ac937b093d57aa613f518f7da14cd2317e1a7f4c81cd568647fc)
- global

<a id="canonical-0f8ae06f0526d462a2f4af44817f2d9cc6ff9d6a4c5b6133aecff50da5c1ef51"></a>

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

<a id="canonical-da5559db8672cd1f1a0d537ff84153582a45a860485476de6f7a35e4903d47e2"></a>

## Direct properties — global / 4043edde7c7b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4605f61ce52a7fea68003e825d57defba1997f9bc10f561db7d393cbde98bb6"></a>

## Next pages — global / 4043edde7c7b / 4

- [Property reference](data-sources--geo_location_set--reference--group-001.md#canonical-5bf5901a30d1ac937b093d57aa613f518f7da14cd2317e1a7f4c81cd568647fc)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)
