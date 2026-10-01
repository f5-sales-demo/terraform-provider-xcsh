---
page_title: "xcsh_filter_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set reference."
---

# xcsh_filter_set reference

<a id="canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5504fc8034cc610d87da66d7b3c47fc763a8197dfe7f08bebd1bb37e79191c9"></a>

## Property reference — Property reference / 6f95a6946c41 / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- Property reference

<a id="canonical-327ba83989be977147e760f19583f1c43e5795e64ec46448ded7914f98a9a6a7"></a>

## Direct properties — Property reference / 6f95a6946c41 / 3

<a id="canonical-c347a83f218e9311726bb2b78b019a00932b046baaf95939c1474e0e85dfda8b"></a>

<a id="canonical-f419f7b577fd32df46a358096d1f68fcb04f258970e5b078443a641662a69687"></a>

## annotations property — Property reference / 6f95a6946c41 / 4

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

<a id="canonical-78addda485ecc54483d7bc76789f5decc98ddcabbacf891cd0368132c3e184ca"></a>

<a id="canonical-553e7d931ce2a2967195ac1fff3f6763e526bd39df2c7627382ce40a3393f677"></a>

## context_key property — Property reference / 6f95a6946c41 / 5

Type: `"string"`. Required.

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

<a id="canonical-6f1fe71750c8dbb3b1416668fefb8e929b0b1386a3ae721edaa3719e53197cb1"></a>

<a id="canonical-13f9cb2f9bba323a22900070b86cb8a8b10f7c3feb0f2c7ed916cd9b0bdd0157"></a>

## description property — Property reference / 6f95a6946c41 / 6

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

<a id="canonical-cc9af6fba0eaf34038c5177391e91d6ba398ae112e8019a216d64a9adb6939e4"></a>

<a id="canonical-e8e19e66b781d4f24b81370d7f9656088f30251cc0cb906efa2f700a1e0da1f8"></a>

## disable property — Property reference / 6f95a6946c41 / 7

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

- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be): complete subsection reference.

<a id="canonical-e7f0f3c58d2a4f31238673d19b50f1185adad4569a801ca1eb1ce11813638bad"></a>

<a id="canonical-2dd65e6a8ee16e2f7528b1390d902821c88a3c326370fb73c1e4ab9e462de84f"></a>

## id property — Property reference / 6f95a6946c41 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-e9176c6d787ed90b8f2f2236221f58d62c2dafc3b6e79264e2f9a0ca8ab5061a"></a>

<a id="canonical-412f7ef2fdbd9786a2f4ecc675509d5491e4f36289e28a39efb0660f3e3fc36a"></a>

## labels property — Property reference / 6f95a6946c41 / 9

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

<a id="canonical-308ab8f99894bbc5957185c9d60a392dfbcf449f3fdd0f62d562a14ff553966a"></a>

<a id="canonical-40f72115aa7ed859edf9c7c912ad4e0af8f4f39278cec01111eaf793fbc1d7a3"></a>

## name property — Property reference / 6f95a6946c41 / 10

Type: `"string"`. Required.

Name of the Filter Set. Must be unique within the namespace.

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

<a id="canonical-372f8c418a6a5eeb6aeaf151b74bdf1c3927ab18b9710b2ba207f2c180dd64b5"></a>

<a id="canonical-20c2383bbc4c2b043dd593d94918b57d82dc3060ce62684ffb8846af9a11a1ca"></a>

## namespace property — Property reference / 6f95a6946c41 / 11

Type: `"string"`. Required.

Namespace where the Filter Set is created.

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

- [timeouts](resources--filter_set--reference--group-001.md#canonical-05449447b77dc9e1d32ce75ac2241551e4beb41538ab2808ca34fbd3112319cd): complete subsection reference.

<a id="canonical-aa87c4ed7b0a8fc38af4582acc71244df950976312ac4f903e896bdc356a907d"></a>

## All schema paths — Property reference / 6f95a6946c41 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--filter_set--reference--group-001.md#canonical-c347a83f218e9311726bb2b78b019a00932b046baaf95939c1474e0e85dfda8b) |
| `context_key` | [context_key](resources--filter_set--reference--group-001.md#canonical-78addda485ecc54483d7bc76789f5decc98ddcabbacf891cd0368132c3e184ca) |
| `description` | [description](resources--filter_set--reference--group-001.md#canonical-6f1fe71750c8dbb3b1416668fefb8e929b0b1386a3ae721edaa3719e53197cb1) |
| `disable` | [disable](resources--filter_set--reference--group-001.md#canonical-cc9af6fba0eaf34038c5177391e91d6ba398ae112e8019a216d64a9adb6939e4) |
| `filter_fields` | [filter_fields](resources--filter_set--reference--group-001.md#canonical-fdd1c23d21b9befa6bc0d3e529acc80f0a1c2ceee7d9bfb4d616d3d46e8b8c31) |
| `filter_fields.date_field` | [filter_fields.date_field](resources--filter_set--reference--group-001.md#canonical-dc6b5d98d7f013cde0e2f176206feb6b431c2a326fbd9bc055711b3cc8dda01e) |
| `filter_fields.date_field.absolute` | [filter_fields.date_field.absolute](resources--filter_set--reference--group-001.md#canonical-3f49888d36bb75de67c50709eb135111ea6a597c7b09dab8cff6c0c60c75d31b) |
| `filter_fields.date_field.absolute.end_date` | [filter_fields.date_field.absolute.end_date](resources--filter_set--reference--group-001.md#canonical-722b19f15512ca0721b55ef185dafa3cedfbcb87dd0173bea9804db9d08b4ceb) |
| `filter_fields.date_field.absolute.start_date` | [filter_fields.date_field.absolute.start_date](resources--filter_set--reference--group-001.md#canonical-0550fd10d27b73cab7aaff25adcd8e1ac93215bd3ed3bf045af5958102ae012c) |
| `filter_fields.date_field.relative` | [filter_fields.date_field.relative](resources--filter_set--reference--group-001.md#canonical-d987148f8371a8e044f3e939b7f1b51963a2c251c567d54911584cc1c2d8b894) |
| `filter_fields.field_id` | [filter_fields.field_id](resources--filter_set--reference--group-001.md#canonical-b6358cb433719346071bdbc67cc44f4875a498eaf0ffa89df5fc6cfe3d8e04d7) |
| `filter_fields.filter_expression_field` | [filter_fields.filter_expression_field](resources--filter_set--reference--group-001.md#canonical-b3942a3389ff4e180d7428b3e17744376c0e12c78c966d0e47931d83520918eb) |
| `filter_fields.filter_expression_field.expression` | [filter_fields.filter_expression_field.expression](resources--filter_set--reference--group-001.md#canonical-d2109691d952dc110da06a0d4184848d16c7b17be3cd02bc8964f2341c126dec) |
| `filter_fields.string_field` | [filter_fields.string_field](resources--filter_set--reference--group-001.md#canonical-21f464710328ed2c9d110eafaa3b50cfcabff3b1c1659038fbcef3244181a46f) |
| `filter_fields.string_field.field_values` | [filter_fields.string_field.field_values](resources--filter_set--reference--group-001.md#canonical-12a58b2d0547bb88c5cd426158d2c4771a8386350f0607c6cfcf2ca72692681e) |
| `id` | [id](resources--filter_set--reference--group-001.md#canonical-e7f0f3c58d2a4f31238673d19b50f1185adad4569a801ca1eb1ce11813638bad) |
| `labels` | [labels](resources--filter_set--reference--group-001.md#canonical-e9176c6d787ed90b8f2f2236221f58d62c2dafc3b6e79264e2f9a0ca8ab5061a) |
| `name` | [name](resources--filter_set--reference--group-001.md#canonical-308ab8f99894bbc5957185c9d60a392dfbcf449f3fdd0f62d562a14ff553966a) |
| `namespace` | [namespace](resources--filter_set--reference--group-001.md#canonical-372f8c418a6a5eeb6aeaf151b74bdf1c3927ab18b9710b2ba207f2c180dd64b5) |
| `timeouts` | [timeouts](resources--filter_set--reference--group-001.md#canonical-e17c8e82b1c1ba21178659a69ff9a313653bd9e0c4517db78eb70516541063f3) |
| `timeouts.create` | [timeouts.create](resources--filter_set--reference--group-001.md#canonical-b09ecd1d60a4672915c37e7c6b866a93b1367beb5857d5f4685f3f2da5015d82) |
| `timeouts.delete` | [timeouts.delete](resources--filter_set--reference--group-001.md#canonical-c6b9498b4c8ed8c2b69f089ddcd0e7a55af441ed3537156c9132b8d9ef7d042e) |
| `timeouts.read` | [timeouts.read](resources--filter_set--reference--group-001.md#canonical-fb04330a5092613012b4d4a8badda216d6920684454623034c6abbc5e2aee1b1) |
| `timeouts.update` | [timeouts.update](resources--filter_set--reference--group-001.md#canonical-14386a2acb4bd6d67868fcf6ade51c04264ee1a54bdce5c61d6d0642f8203b56) |

<a id="canonical-205a66317476e841dfffca4286d40d322104211c4efd35f00440374f2306eaf9"></a>

## Next pages — Property reference / 6f95a6946c41 / 13

- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- [timeouts](resources--filter_set--reference--group-001.md#canonical-05449447b77dc9e1d32ce75ac2241551e4beb41538ab2808ca34fbd3112319cd)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)

<a id="canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c26af3daac47a3ca0bad4d04fe05af426be1f6b8f26e08d775efd99f6067aa7"></a>

## filter_fields — filter_fields / c7aa81d47a9a / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- filter_fields

<a id="canonical-fdd1c23d21b9befa6bc0d3e529acc80f0a1c2ceee7d9bfb4d616d3d46e8b8c31"></a>

Type: `"object"`. list nested block, Optional.

List of fields and their values selected by the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("field_id"),
  validators.ConflictingListObjectAttributes("date_field",
    "filter_expression_field"),
  validators.ConflictingListObjectAttributes("date_field",
    "string_field"),
  validators.ConflictingListObjectAttributes("filter_expression_field",
    "string_field")}
```

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

Terraform syntax:

```terraform
filter_fields {
  # Configure direct properties listed below.
}
```

<a id="canonical-9ab3479fde3f06e3fa9ef05abc56334e1bb94e0e134a6d14778f1ecafbbfa9b9"></a>

## Direct properties — filter_fields / c7aa81d47a9a / 3

- [date_field](resources--filter_set--reference--group-001.md#canonical-3584b3f6cde7a5b9c74a8c3fdbd59a9f7efdebfa072925b76bf4e995e9b41a8e): complete subsection reference.

<a id="canonical-b6358cb433719346071bdbc67cc44f4875a498eaf0ffa89df5fc6cfe3d8e04d7"></a>

<a id="canonical-f2213d556b5a1719ca4995a91ec7924292f87ac4373149c86474d6e333c56d6d"></a>

## field_id property — filter_fields / c7aa81d47a9a / 4

Type: `"string"`. Optional.

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

- [filter_expression_field](resources--filter_set--reference--group-001.md#canonical-479ae4830bf6a022d3e6116176596b7e52c3935ceaeed1e3ea0b087c8749479a): complete subsection reference.

- [string_field](resources--filter_set--reference--group-001.md#canonical-1664ef7f5cdaf92926f2b9b09716dcbbd68984a04fa21dffbd5992443bb8c427): complete subsection reference.

<a id="canonical-c61e2b37ebd7d6223f086641c7bb4953168f0d5526cd87d36882c190f8b7bd1f"></a>

## Next pages — filter_fields / c7aa81d47a9a / 5

- [filter_fields.date_field](resources--filter_set--reference--group-001.md#canonical-3584b3f6cde7a5b9c74a8c3fdbd59a9f7efdebfa072925b76bf4e995e9b41a8e)
- [filter_fields.filter_expression_field](resources--filter_set--reference--group-001.md#canonical-479ae4830bf6a022d3e6116176596b7e52c3935ceaeed1e3ea0b087c8749479a)
- [filter_fields.string_field](resources--filter_set--reference--group-001.md#canonical-1664ef7f5cdaf92926f2b9b09716dcbbd68984a04fa21dffbd5992443bb8c427)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)

<a id="canonical-3584b3f6cde7a5b9c74a8c3fdbd59a9f7efdebfa072925b76bf4e995e9b41a8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cede5c4f5020ee22079f379c03846cab411a1171fc0259147a3ac80affd9f1d8"></a>

## filter_fields.date_field — filter_fields.date_field / 107d6d08f5bc / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- filter_fields.date_field

<a id="canonical-dc6b5d98d7f013cde0e2f176206feb6b431c2a326fbd9bc055711b3cc8dda01e"></a>

Type: `"object"`. single nested block, Optional.

Either an absolute time range or a relative time interval.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("absolute",
    "relative")}
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
  "x-ves-oneof-field-range_type": "[\"absolute\",\"relative\"]"
}
```

Terraform syntax:

```terraform
date_field {
  # Configure direct properties listed below.
}
```

<a id="canonical-58ef63528ff2ca0b1b6de2ea62ef1b2e65597de824897bbbc850531276afce76"></a>

## Direct properties — filter_fields.date_field / 107d6d08f5bc / 3

- [absolute](resources--filter_set--reference--group-001.md#canonical-0911bc9ee7c82af12da0b47322c0a8cb22382a13a3877d06cf8bde75d0a5aff3): complete subsection reference.

<a id="canonical-d987148f8371a8e044f3e939b7f1b51963a2c251c567d54911584cc1c2d8b894"></a>

<a id="canonical-953446160da1b9db8ef66f70cb228fb618e27c4cebad0d859a78dc705c25a3a5"></a>

## relative property — filter_fields.date_field / 107d6d08f5bc / 4

Type: `"string"`. Optional.

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

<a id="canonical-173967f2c481f2f17e8066011d0a2353eab36367f44ba1f95899ef9a31b8227f"></a>

## Next pages — filter_fields.date_field / 107d6d08f5bc / 5

- [filter_fields.date_field.absolute](resources--filter_set--reference--group-001.md#canonical-0911bc9ee7c82af12da0b47322c0a8cb22382a13a3877d06cf8bde75d0a5aff3)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)

<a id="canonical-0911bc9ee7c82af12da0b47322c0a8cb22382a13a3877d06cf8bde75d0a5aff3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b6db01e18f49f7008104d431c44ebde43465bdf4c30fcaeb626b110359147e6"></a>

## filter_fields.date_field.absolute — filter_fields.date_field.absolute / a262366fa470 / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- [filter_fields.date_field](resources--filter_set--reference--group-001.md#canonical-3584b3f6cde7a5b9c74a8c3fdbd59a9f7efdebfa072925b76bf4e995e9b41a8e)
- filter_fields.date_field.absolute

<a id="canonical-3f49888d36bb75de67c50709eb135111ea6a597c7b09dab8cff6c0c60c75d31b"></a>

Type: `"object"`. single nested block, Optional.

Date range is for selecting a date range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("end_date",
    "start_date")}
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
absolute {
  # Configure direct properties listed below.
}
```

<a id="canonical-1257db80fbecac542bf637395022f2dceab17d565a30dcb8448a2e706d23cd60"></a>

## Direct properties — filter_fields.date_field.absolute / a262366fa470 / 3

<a id="canonical-722b19f15512ca0721b55ef185dafa3cedfbcb87dd0173bea9804db9d08b4ceb"></a>

<a id="canonical-219696bd251dd15639cdac9f893fb1e85b4641ee44ba0fd875b2fba6c88bf5d8"></a>

## end_date property — filter_fields.date_field.absolute / a262366fa470 / 4

Type: `"string"`. Optional.

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

<a id="canonical-0550fd10d27b73cab7aaff25adcd8e1ac93215bd3ed3bf045af5958102ae012c"></a>

<a id="canonical-f7295db70e869a82215635dd453edef8bee715f081cecfd0a9d389fcf20df805"></a>

## start_date property — filter_fields.date_field.absolute / a262366fa470 / 5

Type: `"string"`. Optional.

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

<a id="canonical-7ba47cc6c6b011486aec568491fe1744bb9dd17cb63df3cd9c5aec473892d334"></a>

## Next pages — filter_fields.date_field.absolute / a262366fa470 / 6

- [filter_fields.date_field](resources--filter_set--reference--group-001.md#canonical-3584b3f6cde7a5b9c74a8c3fdbd59a9f7efdebfa072925b76bf4e995e9b41a8e)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)

<a id="canonical-479ae4830bf6a022d3e6116176596b7e52c3935ceaeed1e3ea0b087c8749479a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efd662f45b8e61cd1e13fff795b23c3e657039c3c36cf8e63d0bea7297588625"></a>

## filter_fields.filter_expression_field — filter_fields.filter_expression_field / 604d5fc8a3d3 / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- filter_fields.filter_expression_field

<a id="canonical-b3942a3389ff4e180d7428b3e17744376c0e12c78c966d0e47931d83520918eb"></a>

Type: `"object"`. single nested block, Optional.

Filter Expression Field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expression")}
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
filter_expression_field {
  # Configure direct properties listed below.
}
```

<a id="canonical-edc7bd2a658a5ab82bc45c9a9e69894fc3d9408feb27f1ba1dba5b69f324230a"></a>

## Direct properties — filter_fields.filter_expression_field / 604d5fc8a3d3 / 3

<a id="canonical-d2109691d952dc110da06a0d4184848d16c7b17be3cd02bc8964f2341c126dec"></a>

<a id="canonical-c91c40adbe59617d7b628b7bdb80c5de9d1d536971336d6c94c0becbaf37ae5c"></a>

## expression property — filter_fields.filter_expression_field / 604d5fc8a3d3 / 4

Type: `"string"`. Optional.

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

<a id="canonical-48f1804cb4fa04e5d676e8bddba10eaa131025d066c88927fb5ec49d2cb2f43a"></a>

## Next pages — filter_fields.filter_expression_field / 604d5fc8a3d3 / 5

- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)

<a id="canonical-1664ef7f5cdaf92926f2b9b09716dcbbd68984a04fa21dffbd5992443bb8c427"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68153201dbdac9fe2347f23e48da142901e66ed309014649f7ba527c714fb93f"></a>

## filter_fields.string_field — filter_fields.string_field / 949e5bea4673 / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- filter_fields.string_field

<a id="canonical-21f464710328ed2c9d110eafaa3b50cfcabff3b1c1659038fbcef3244181a46f"></a>

Type: `"object"`. single nested block, Optional.

Filter String Field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("field_values")}
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
string_field {
  # Configure direct properties listed below.
}
```

<a id="canonical-20693ac00c2e30c6df4351e5f2528026986ab9651dd11e357cbc04246448f644"></a>

## Direct properties — filter_fields.string_field / 949e5bea4673 / 3

<a id="canonical-12a58b2d0547bb88c5cd426158d2c4771a8386350f0607c6cfcf2ca72692681e"></a>

<a id="canonical-44594c34cebe370917946024900324517739724dcbeb7a53269eee1b00690ed4"></a>

## field_values property — filter_fields.string_field / 949e5bea4673 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-53c79a0091aca95a2ccf50300f2a46831b1c3816c87fa8f3795ff575699b7441"></a>

## Next pages — filter_fields.string_field / 949e5bea4673 / 5

- [filter_fields](resources--filter_set--reference--group-001.md#canonical-62bf68c8b4a5e4bf620094aa29d75e8868aa2dcf89563a6b18b58751a136f8be)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)

<a id="canonical-05449447b77dc9e1d32ce75ac2241551e4beb41538ab2808ca34fbd3112319cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71aee3c46e4dd8dae78d4c2a5e15de061dc4a68a7aa11cb85cf22a286bd567c6"></a>

## timeouts — timeouts / ac54e8c3e3fe / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- timeouts

<a id="canonical-e17c8e82b1c1ba21178659a69ff9a313653bd9e0c4517db78eb70516541063f3"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-10cc8e96a6de59b7ce7b2a5c1aa324f096d8d12b6a49701c85ed99ea4ed0debf"></a>

## Direct properties — timeouts / ac54e8c3e3fe / 3

<a id="canonical-b09ecd1d60a4672915c37e7c6b866a93b1367beb5857d5f4685f3f2da5015d82"></a>

<a id="canonical-1a460329f5d2fce5c716c4c3414674cda421f4afe0f5882434c382c734c43677"></a>

## create property — timeouts / ac54e8c3e3fe / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c6b9498b4c8ed8c2b69f089ddcd0e7a55af441ed3537156c9132b8d9ef7d042e"></a>

<a id="canonical-4f62942a584ffa4b4d490ee9dc6b91579f39f8b9f30f658eee98dddce03434a1"></a>

## delete property — timeouts / ac54e8c3e3fe / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-fb04330a5092613012b4d4a8badda216d6920684454623034c6abbc5e2aee1b1"></a>

<a id="canonical-4d1d4bc80f2d6905040cd499e90da1f92052ce4776539b6c41b6d2a6d6d8efca"></a>

## read property — timeouts / ac54e8c3e3fe / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-14386a2acb4bd6d67868fcf6ade51c04264ee1a54bdce5c61d6d0642f8203b56"></a>

<a id="canonical-642bed77f03faca2186f8466f71ed9b8ae8389c262c2a63612134141f98c7e9c"></a>

## update property — timeouts / ac54e8c3e3fe / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-98b90bde03e8ece61c7d9cdfe66e3a67fe358e2eff5507850f62ffcbad830a2d"></a>

## Next pages — timeouts / ac54e8c3e3fe / 8

- [Property reference](resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
