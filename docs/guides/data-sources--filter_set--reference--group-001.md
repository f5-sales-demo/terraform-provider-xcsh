---
page_title: "xcsh_filter_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set reference."
---

# xcsh_filter_set reference

<a id="canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6a4bf30616a0139e8010fcbfcb492b73478d04a4904ab8577216fdf94e5ac2c"></a>

## Property reference — Property reference / 2827e5a2f266 / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- Property reference

<a id="canonical-0042989af35ec6b24c2fab3cb79c8841458385d18ea073535eba8f7bccd34bd4"></a>

## Direct properties — Property reference / 2827e5a2f266 / 3

<a id="canonical-ba3127f056b5b1b7d3679c1379463efdaf20acb2d4f06aa498d788de87fcbbd6"></a>

<a id="canonical-ae8a185527ad9240bac3a95b2af44f82fdee8064f822dcbdb163ddf51e763394"></a>

## annotations property — Property reference / 2827e5a2f266 / 4

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

<a id="canonical-e7d65b00244c08719c7fce50e6159aff808843213fb3800a6975771f9ba493a7"></a>

<a id="canonical-2d5c3a2128d9a348e05a61e56277a370f5163e11bc94b20882d477b23c2b2578"></a>

## context_key property — Property reference / 2827e5a2f266 / 5

Type: `"string"`. Computed.

Indexable context key that identifies a page or page type for which the FilterSet is applicable.

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

<a id="canonical-ef9493caa7add116974bf6835ad74440041c27a8f73505a4af9130c2039816d0"></a>

<a id="canonical-693469f4e34244e3d40ff6bb8375dd0bbdc9dfcf40d5e8b2be4ef18b6e27b0ac"></a>

## description property — Property reference / 2827e5a2f266 / 6

Type: `"string"`. Computed.

Description of the FilterSet.

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

- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd): complete subsection reference.

<a id="canonical-eae653038bd91616cc811e5d6c63a61d3d6640322e5cf0eec94788ba394f7bdc"></a>

<a id="canonical-29d132c209bb433a27575026914e13b5b61219466757280586bdfb99d6b74d8f"></a>

## id property — Property reference / 2827e5a2f266 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6626e032a4c968fdc224591221d21cef615dc037d997f14f5135763f7bb9a617"></a>

<a id="canonical-f4fa22f48d1e9cc5c37e97b6e8423296b1ca3de354cee4a3e45f3d3485d54196"></a>

## labels property — Property reference / 2827e5a2f266 / 8

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

<a id="canonical-e40f62c1e8272801e2f315a6fa65830144357935ef8b4ea78e7424e8b404728d"></a>

<a id="canonical-dd7b447e72d85a7eb38b507ea3c5baf855d8272bb045119ef0041346c5cf2475"></a>

## name property — Property reference / 2827e5a2f266 / 9

Type: `"string"`. Required.

Name of the FilterSet.

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

<a id="canonical-88151d075e2e41b5298a95630b7e37e66612aba4d3b2acafad93f77826c6d78b"></a>

<a id="canonical-224de29c62c7da28a83c98c41daa0584fc569b64adfc44b53f1be320b8eba2d5"></a>

## namespace property — Property reference / 2827e5a2f266 / 10

Type: `"string"`. Required.

Namespace where the FilterSet exists.

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

<a id="canonical-f54eef756099bc0ed84b8709f7a3b07e8475224fd8b9f6f5b48a6708df74803d"></a>

## All schema paths — Property reference / 2827e5a2f266 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--filter_set--reference--group-001.md#canonical-ba3127f056b5b1b7d3679c1379463efdaf20acb2d4f06aa498d788de87fcbbd6) |
| `context_key` | [context_key](data-sources--filter_set--reference--group-001.md#canonical-e7d65b00244c08719c7fce50e6159aff808843213fb3800a6975771f9ba493a7) |
| `description` | [description](data-sources--filter_set--reference--group-001.md#canonical-ef9493caa7add116974bf6835ad74440041c27a8f73505a4af9130c2039816d0) |
| `filter_fields` | [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-8dd3277ef9c65aa59609f0007a7e13a75cb709a9866f1589ed87e05ac831b44e) |
| `filter_fields.date_field` | [filter_fields.date_field](data-sources--filter_set--reference--group-001.md#canonical-bdddb2132e104ce1ca859cde132ad0dfbdf5f250d7627d086604eb0c551844c8) |
| `filter_fields.date_field.absolute` | [filter_fields.date_field.absolute](data-sources--filter_set--reference--group-001.md#canonical-730cfc42924b7e2695d5a1219a1b21cbc9edc2949e50c8b98a038f9ead1b5fc1) |
| `filter_fields.date_field.absolute.end_date` | [filter_fields.date_field.absolute.end_date](data-sources--filter_set--reference--group-001.md#canonical-6b23f0de9e15bc20ae1ad89044c75a53a555e9d5582838ddcee6b7130af14e61) |
| `filter_fields.date_field.absolute.start_date` | [filter_fields.date_field.absolute.start_date](data-sources--filter_set--reference--group-001.md#canonical-459d010ce497bd4eba0668f8ba7fb55d56137e51d4f7f73cc091106a085e3903) |
| `filter_fields.date_field.relative` | [filter_fields.date_field.relative](data-sources--filter_set--reference--group-001.md#canonical-39fd97ef988b3571fd3fb6a5cc03dce89fd37cca381b94765dd2580982da4fab) |
| `filter_fields.field_id` | [filter_fields.field_id](data-sources--filter_set--reference--group-001.md#canonical-56af4553bd8f611bc2ba3823614c3eecf291c2d2cf37432f00a8a645010ed68f) |
| `filter_fields.filter_expression_field` | [filter_fields.filter_expression_field](data-sources--filter_set--reference--group-001.md#canonical-c590483f28b8bdf365672c74f94df8591a324a1f186db150ec7226597f287e2f) |
| `filter_fields.filter_expression_field.expression` | [filter_fields.filter_expression_field.expression](data-sources--filter_set--reference--group-001.md#canonical-24b6885edfb8dc672fbbad9cff68586d662cc69ff846d44baa5e6250cf6c9b77) |
| `filter_fields.string_field` | [filter_fields.string_field](data-sources--filter_set--reference--group-001.md#canonical-0d690a3cf7cc41e14eefa6c400be4a260f4226fb438442703246c781251d5264) |
| `filter_fields.string_field.field_values` | [filter_fields.string_field.field_values](data-sources--filter_set--reference--group-001.md#canonical-27246d39fce95c9ffcc54681cfa242b38a9e6cfeb0a4b85a9552e5713fe713d0) |
| `id` | [id](data-sources--filter_set--reference--group-001.md#canonical-eae653038bd91616cc811e5d6c63a61d3d6640322e5cf0eec94788ba394f7bdc) |
| `labels` | [labels](data-sources--filter_set--reference--group-001.md#canonical-6626e032a4c968fdc224591221d21cef615dc037d997f14f5135763f7bb9a617) |
| `name` | [name](data-sources--filter_set--reference--group-001.md#canonical-e40f62c1e8272801e2f315a6fa65830144357935ef8b4ea78e7424e8b404728d) |
| `namespace` | [namespace](data-sources--filter_set--reference--group-001.md#canonical-88151d075e2e41b5298a95630b7e37e66612aba4d3b2acafad93f77826c6d78b) |

<a id="canonical-9b9d49c2487a09ca0ab00aedb97665f28eff51b5ae26dd94511a41ab32e305e3"></a>

## Next pages — Property reference / 2827e5a2f266 / 12

- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)

<a id="canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9f3f7f7cf6f195900e388ac6c3517272e122664f72c52371d6215a15d3cabce"></a>

## filter_fields — filter_fields / 274e3d514e08 / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599)
- filter_fields

<a id="canonical-8dd3277ef9c65aa59609f0007a7e13a75cb709a9866f1589ed87e05ac831b44e"></a>

Type: `"list"`. Computed.

List of fields and their values selected by the user.

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

<a id="canonical-ab320f852d7129da833d66276fa01b30684e4bd43bbb9c5e3448fea0cd9449f9"></a>

## Direct properties — filter_fields / 274e3d514e08 / 3

- [date_field](data-sources--filter_set--reference--group-001.md#canonical-c1db6181b6a0bc8dd9c6af370091d27a11039465a7fee685f89ff35cfd710ff3): complete subsection reference.

<a id="canonical-56af4553bd8f611bc2ba3823614c3eecf291c2d2cf37432f00a8a645010ed68f"></a>

<a id="canonical-8116ed70f1d441bac3ae2a6c7dceb8667d7ff233a88142d7162ca6d19b853f82"></a>

## field_id property — filter_fields / 274e3d514e08 / 4

Type: `"string"`. Computed.

Identifier for the field that maps to some UI filter component.

Upstream description:

An identifier for the field that maps to some UI filter component.

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

- [filter_expression_field](data-sources--filter_set--reference--group-001.md#canonical-864dc06cba6474f210941aacf008a01892e13d58985e6f7dcebb6139566be13e): complete subsection reference.

- [string_field](data-sources--filter_set--reference--group-001.md#canonical-8db38053e7b40f18e57363b5b63c298513e852eaee6351fb9a0ff4baedd7442c): complete subsection reference.

<a id="canonical-8e0d122bccceccdbaf782b68a7bb846f6f216a7dacc1a9d710de571514eae0dc"></a>

## Next pages — filter_fields / 274e3d514e08 / 5

- [filter_fields.date_field](data-sources--filter_set--reference--group-001.md#canonical-c1db6181b6a0bc8dd9c6af370091d27a11039465a7fee685f89ff35cfd710ff3)
- [filter_fields.filter_expression_field](data-sources--filter_set--reference--group-001.md#canonical-864dc06cba6474f210941aacf008a01892e13d58985e6f7dcebb6139566be13e)
- [filter_fields.string_field](data-sources--filter_set--reference--group-001.md#canonical-8db38053e7b40f18e57363b5b63c298513e852eaee6351fb9a0ff4baedd7442c)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)

<a id="canonical-c1db6181b6a0bc8dd9c6af370091d27a11039465a7fee685f89ff35cfd710ff3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b64e66a47c51d5bbe6b417c32b04be09fd0c0f32c76403c4c7c19ccf033db804"></a>

## filter_fields.date_field — filter_fields.date_field / 52d9ec39dfc6 / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- filter_fields.date_field

<a id="canonical-bdddb2132e104ce1ca859cde132ad0dfbdf5f250d7627d086604eb0c551844c8"></a>

Type: `"single"`. Computed.

Either an absolute time range or a relative time interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-range_type": "[\"absolute\",\"relative\"]"
}
```

<a id="canonical-7aedabf071697f1928cf3b4c3154a17428824d6ea9df52c22ab71ba4067511b7"></a>

## Direct properties — filter_fields.date_field / 52d9ec39dfc6 / 3

- [absolute](data-sources--filter_set--reference--group-001.md#canonical-76fa92f42aed4881d2531a7fea0fa91b59a8e0254e2f76f57152bfb93352158a): complete subsection reference.

<a id="canonical-39fd97ef988b3571fd3fb6a5cc03dce89fd37cca381b94765dd2580982da4fab"></a>

<a id="canonical-0029415ebb367764160a388c58939031181abea3221aeaeff639ef023de8e46a"></a>

## relative property — filter_fields.date_field / 52d9ec39dfc6 / 4

Type: `"string"`. Computed.

Exclusive with \[absolute\] relative time duration.

Upstream description:

Exclusive with \[absolute\] relative time duration.

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

<a id="canonical-c1e0ee3cfe9c3de3e31f7d42cf47ab77bb037921e47b90995e68fc745270596c"></a>

## Next pages — filter_fields.date_field / 52d9ec39dfc6 / 5

- [filter_fields.date_field.absolute](data-sources--filter_set--reference--group-001.md#canonical-76fa92f42aed4881d2531a7fea0fa91b59a8e0254e2f76f57152bfb93352158a)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)

<a id="canonical-76fa92f42aed4881d2531a7fea0fa91b59a8e0254e2f76f57152bfb93352158a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26a1dd263de14ae0c65663e4322281ad43ad08268c768437c692d205e719141a"></a>

## filter_fields.date_field.absolute — filter_fields.date_field.absolute / 6a0d43e9c988 / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- [filter_fields.date_field](data-sources--filter_set--reference--group-001.md#canonical-c1db6181b6a0bc8dd9c6af370091d27a11039465a7fee685f89ff35cfd710ff3)
- filter_fields.date_field.absolute

<a id="canonical-730cfc42924b7e2695d5a1219a1b21cbc9edc2949e50c8b98a038f9ead1b5fc1"></a>

Type: `"single"`. Computed.

Date range is for selecting a date range.

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

<a id="canonical-4fd5bef0db0c6f43bedc220a4e300c4230626a896676588a8d0339c359ccd13e"></a>

## Direct properties — filter_fields.date_field.absolute / 6a0d43e9c988 / 3

<a id="canonical-6b23f0de9e15bc20ae1ad89044c75a53a555e9d5582838ddcee6b7130af14e61"></a>

<a id="canonical-6e44dc56ac8574226c859b1370ae7fd071b98c1df7c8a4bdb82f09be6f8c0fe6"></a>

## end_date property — filter_fields.date_field.absolute / 6a0d43e9c988 / 4

Type: `"string"`. Computed.

End Date. Contains end date.

Upstream description:

Contains end date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

<a id="canonical-459d010ce497bd4eba0668f8ba7fb55d56137e51d4f7f73cc091106a085e3903"></a>

<a id="canonical-6b93e75a69bfe7deab92f470c7c9d03a834e23c460f0d54489b09386b8f4465c"></a>

## start_date property — filter_fields.date_field.absolute / 6a0d43e9c988 / 5

Type: `"string"`. Computed.

Start Date. Contains start date.

Upstream description:

Contains start date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

<a id="canonical-3bc510fe37a9b680a58d1469842f18e41c0a7eed33406edffe87d810a70eb5e3"></a>

## Next pages — filter_fields.date_field.absolute / 6a0d43e9c988 / 6

- [filter_fields.date_field](data-sources--filter_set--reference--group-001.md#canonical-c1db6181b6a0bc8dd9c6af370091d27a11039465a7fee685f89ff35cfd710ff3)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)

<a id="canonical-864dc06cba6474f210941aacf008a01892e13d58985e6f7dcebb6139566be13e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbf03f18c2d02208133eeeaca1625583b5657206c8b268a8c775e0fdaaab716d"></a>

## filter_fields.filter_expression_field — filter_fields.filter_expression_field / a09e4e95e79d / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- filter_fields.filter_expression_field

<a id="canonical-c590483f28b8bdf365672c74f94df8591a324a1f186db150ec7226597f287e2f"></a>

Type: `"single"`. Computed.

Filter Expression Field.

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

<a id="canonical-b0ce475ab80e4ed63554de79e8eb1c95ed6320ec40e5c0173b1a43ec409bcc47"></a>

## Direct properties — filter_fields.filter_expression_field / a09e4e95e79d / 3

<a id="canonical-24b6885edfb8dc672fbbad9cff68586d662cc69ff846d44baa5e6250cf6c9b77"></a>

<a id="canonical-c953d4f7509b64a2f950b341782d5a4deb47e585a3fe62b5f5c95687b5cfc70c"></a>

## expression property — filter_fields.filter_expression_field / a09e4e95e79d / 4

Type: `"string"`. Computed.

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

Upstream description:

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

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

<a id="canonical-b490d44769bb7d0e7f5a1674dbd8fcf406b076252f33b4aa97694a89c59622ab"></a>

## Next pages — filter_fields.filter_expression_field / a09e4e95e79d / 5

- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)

<a id="canonical-8db38053e7b40f18e57363b5b63c298513e852eaee6351fb9a0ff4baedd7442c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dac129cab2ae2d666a0d45079ca6c710bd97f5e80b077954fc0141ce8cb20049"></a>

## filter_fields.string_field — filter_fields.string_field / cfd2744b3bde / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- [Property reference](data-sources--filter_set--reference--group-001.md#canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599)
- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- filter_fields.string_field

<a id="canonical-0d690a3cf7cc41e14eefa6c400be4a260f4226fb438442703246c781251d5264"></a>

Type: `"single"`. Computed.

Filter String Field.

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

<a id="canonical-f55904dd247cddc0bed15c29073ce7282f810ae7cd8aff2427c0e07da16c7012"></a>

## Direct properties — filter_fields.string_field / cfd2744b3bde / 3

<a id="canonical-27246d39fce95c9ffcc54681cfa242b38a9e6cfeb0a4b85a9552e5713fe713d0"></a>

<a id="canonical-c83a36a2ba0c76f3162c5f015b6f74fac6c6b4e18dd186b821ef38e6a88ca699"></a>

## field_values property — filter_fields.string_field / cfd2744b3bde / 4

Type: `["list", "string"]`. Computed.

String Value(s). Field specification or configuration

Upstream description:

Field specification or configuration

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-3602fdf23de84bbc40a62caa83857a7ffcba77db5c37c83f930d6343f87f116e"></a>

## Next pages — filter_fields.string_field / cfd2744b3bde / 5

- [filter_fields](data-sources--filter_set--reference--group-001.md#canonical-6de5af2ede9f9b39eec1c3b2c07ab0ae250c24ba502f5d25eb321341b4717bfd)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
