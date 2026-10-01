---
page_title: "xcsh_data_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group reference."
---

# xcsh_data_group reference

<a id="canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64a775346b4db59baf534e7fd3b68a3edfbd249d26a0582971bd6dfec507bddc"></a>

## Property reference — Property reference / a2355e4abb27 / 2

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
- Property reference

<a id="canonical-d09899735fb96aaf4e8883a775c6d1dadbf949a13be9783ad6543dde5ba6b74c"></a>

## Direct properties — Property reference / a2355e4abb27 / 3

- [address_records](data-sources--data_group--reference--group-001.md#canonical-00b5d8308363bc8fbaa542dcf2ab2cf9f205b8df91414df0267228fdc983060a): complete subsection reference.

<a id="canonical-c24c030deb47569174ea3cce8ee80d3f983df7599def5a5d050acdb8d7097ba0"></a>

<a id="canonical-052ea29900a59bb5f173b1a0cea335b0e545a03011ed40b1ace1d5082dd0888d"></a>

## annotations property — Property reference / a2355e4abb27 / 4

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

<a id="canonical-6500a44c468ec1365a3a51015fc0f6e4ac6db2ac7dd81ea246fe8476d0f7eb53"></a>

<a id="canonical-cd736946a1248f0a5ed0506e25cee5bff61e2daf59a33364b103967df254d48e"></a>

## description property — Property reference / a2355e4abb27 / 5

Type: `"string"`. Computed.

Description of the DataGroup.

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

<a id="canonical-05531dc1adb84271b8de0e9c06b56c705d6ffa0a94cd6e2d23ad50797f4677a6"></a>

<a id="canonical-d5805ec346c861555caa4e7ae14cff5e8064ad91e41f95dd2e65e599586fde2f"></a>

## id property — Property reference / a2355e4abb27 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [integer_records](data-sources--data_group--reference--group-001.md#canonical-dc4c9144667daa4c1f1b6e0a97212c44d6cb24966534b7266e8c48d5a65ea367): complete subsection reference.

<a id="canonical-0982a12a49a01e20e78c54392882d0c1dce08414bd86bfe296b47b9dc3223b08"></a>

<a id="canonical-1841dd48911ae46f647208167938390bea8bcb8542bec54f74815224cb2b4be8"></a>

## labels property — Property reference / a2355e4abb27 / 7

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

<a id="canonical-8d1dfebc35e0af1b675cc17405b5a738ed2ee0fa99f2d5ac0593834955808971"></a>

<a id="canonical-a90b81c54ec2b3418fe83db85ef77b74a98032292c87fed9a0e712932643e1c4"></a>

## name property — Property reference / a2355e4abb27 / 8

Type: `"string"`. Required.

Name of the DataGroup.

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

<a id="canonical-4345facbb42696051adce7dc967a514d46ed545e15439138129222b06b115ef0"></a>

<a id="canonical-dc623bb3fc8894bcbaa627d33e18c616da04b89b2d4fac376f285e6cfcb65ae9"></a>

## namespace property — Property reference / a2355e4abb27 / 9

Type: `"string"`. Required.

Namespace where the DataGroup exists.

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

- [string_records](data-sources--data_group--reference--group-001.md#canonical-c084520d30fe20c74e9bfb845c50e3e5c7a37deff56c9b2ff1beeb09721bbe47): complete subsection reference.

<a id="canonical-18269c28abc46b85ceaeaba56876fcb532cca66236e1216e398c57f2f8ff7af7"></a>

## All schema paths — Property reference / a2355e4abb27 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_records` | [address_records](data-sources--data_group--reference--group-001.md#canonical-a5ef1ae65e5e5c660a84c967aae60ac621ad67597a94832168731004496e319f) |
| `address_records.records` | [address_records.records](data-sources--data_group--reference--group-001.md#canonical-cd53a80a9a6f5ec9e3e92566d26fb38cdfe2374dd414cf6aefa702afc0fefba6) |
| `annotations` | [annotations](data-sources--data_group--reference--group-001.md#canonical-c24c030deb47569174ea3cce8ee80d3f983df7599def5a5d050acdb8d7097ba0) |
| `description` | [description](data-sources--data_group--reference--group-001.md#canonical-6500a44c468ec1365a3a51015fc0f6e4ac6db2ac7dd81ea246fe8476d0f7eb53) |
| `id` | [id](data-sources--data_group--reference--group-001.md#canonical-05531dc1adb84271b8de0e9c06b56c705d6ffa0a94cd6e2d23ad50797f4677a6) |
| `integer_records` | [integer_records](data-sources--data_group--reference--group-001.md#canonical-3108733e1b25002987d1b27cc56331c3e0f9f1132bcf93613fcfb0fb1af42b8c) |
| `integer_records.records` | [integer_records.records](data-sources--data_group--reference--group-001.md#canonical-ca5950c9211fd9ce453b573b8f58797b1df777459ed98add088ce87d41f29a10) |
| `labels` | [labels](data-sources--data_group--reference--group-001.md#canonical-0982a12a49a01e20e78c54392882d0c1dce08414bd86bfe296b47b9dc3223b08) |
| `name` | [name](data-sources--data_group--reference--group-001.md#canonical-8d1dfebc35e0af1b675cc17405b5a738ed2ee0fa99f2d5ac0593834955808971) |
| `namespace` | [namespace](data-sources--data_group--reference--group-001.md#canonical-4345facbb42696051adce7dc967a514d46ed545e15439138129222b06b115ef0) |
| `string_records` | [string_records](data-sources--data_group--reference--group-001.md#canonical-cd3e15185134ca4c9159fef732c5790c62a4d4803b8908cc5b411a8208de5655) |
| `string_records.records` | [string_records.records](data-sources--data_group--reference--group-001.md#canonical-b1495cc4201660ce8372c97641687d5e9a052ae5bdce4514ccad02989269cf4b) |

<a id="canonical-5df4459309e02dc1d71d5557107ef8bdb31183fb9c3fad49c248661cf36b74b9"></a>

## Next pages — Property reference / a2355e4abb27 / 11

- [address_records](data-sources--data_group--reference--group-001.md#canonical-00b5d8308363bc8fbaa542dcf2ab2cf9f205b8df91414df0267228fdc983060a)
- [integer_records](data-sources--data_group--reference--group-001.md#canonical-dc4c9144667daa4c1f1b6e0a97212c44d6cb24966534b7266e8c48d5a65ea367)
- [string_records](data-sources--data_group--reference--group-001.md#canonical-c084520d30fe20c74e9bfb845c50e3e5c7a37deff56c9b2ff1beeb09721bbe47)
- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)

<a id="canonical-00b5d8308363bc8fbaa542dcf2ab2cf9f205b8df91414df0267228fdc983060a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d776202aec3e2f4cd3f2b40c7063e291c844d8ea54b54234a5c58be9b3ff044b"></a>

## address_records — address_records / 6b28e2d20d7e / 2

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
- [Property reference](data-sources--data_group--reference--group-001.md#canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81)
- address_records

<a id="canonical-a5ef1ae65e5e5c660a84c967aae60ac621ad67597a94832168731004496e319f"></a>

Type: `"single"`. Computed.

\[OneOf: address\_records, integer\_records, string\_records\] Address Record. Data group with
address record List.

Upstream description:

Data group with address record List.

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

- [address_records](data-sources--data_group--reference--group-001.md#canonical-a5ef1ae65e5e5c660a84c967aae60ac621ad67597a94832168731004496e319f)
- [integer_records](data-sources--data_group--reference--group-001.md#canonical-3108733e1b25002987d1b27cc56331c3e0f9f1132bcf93613fcfb0fb1af42b8c)
- [string_records](data-sources--data_group--reference--group-001.md#canonical-cd3e15185134ca4c9159fef732c5790c62a4d4803b8908cc5b411a8208de5655)

Select alternatives according to the provider validators above.

<a id="canonical-6adc72f8dd041fa2b3b34689884269327bacec086df28f9716c0ba409d1c3b7a"></a>

## Direct properties — address_records / 6b28e2d20d7e / 3

<a id="canonical-cd53a80a9a6f5ec9e3e92566d26fb38cdfe2374dd414cf6aefa702afc0fefba6"></a>

<a id="canonical-90c139004de2568c6153b82fa6982289bbae99cc3c7a7482aaf909bec2722a7b"></a>

## records property — address_records / 6b28e2d20d7e / 4

Type: `["map", "string"]`. Computed.

Address records. Configuration parameter for records

Upstream description:

Configuration parameter for records

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
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

<a id="canonical-1af615a816aaab1fb90e9d00c8e836e3a97ee29b38cb227ee1e24eabc124d477"></a>

## Next pages — address_records / 6b28e2d20d7e / 5

- [Property reference](data-sources--data_group--reference--group-001.md#canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81)
- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)

<a id="canonical-dc4c9144667daa4c1f1b6e0a97212c44d6cb24966534b7266e8c48d5a65ea367"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d0bc6baa6aa0bf412e9fd4462981eefb3f9aeaed49c7c3f2227a39005f6d5ee"></a>

## integer_records — integer_records / 77841320d50e / 2

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
- [Property reference](data-sources--data_group--reference--group-001.md#canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81)
- integer_records

<a id="canonical-3108733e1b25002987d1b27cc56331c3e0f9f1132bcf93613fcfb0fb1af42b8c"></a>

Type: `"single"`. Computed.

Configuration parameter for integer records.

Upstream description:

Data group with integer record List.

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

<a id="canonical-1af8089fb32036573d3f2b2f5974ebbdfcd7e34815ad131174679714dcfe39dc"></a>

## Direct properties — integer_records / 77841320d50e / 3

<a id="canonical-ca5950c9211fd9ce453b573b8f58797b1df777459ed98add088ce87d41f29a10"></a>

<a id="canonical-b6f2da93ea1810004b3b2ef1341b341b4c83e4445c9bba1451e76ab7f89fbf70"></a>

## records property — integer_records / 77841320d50e / 4

Type: `["map", "string"]`. Computed.

Integer records. Configuration parameter for records

Upstream description:

Configuration parameter for records

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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  }
}
```

<a id="canonical-89dcf1ce509e212f237886eabce7ad77400cfffbbbbdff9276909041e4e8e66b"></a>

## Next pages — integer_records / 77841320d50e / 5

- [Property reference](data-sources--data_group--reference--group-001.md#canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81)
- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)

<a id="canonical-c084520d30fe20c74e9bfb845c50e3e5c7a37deff56c9b2ff1beeb09721bbe47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8734f915007e85a0a499c34699d350d78c92611fe5a47715590ddb45217c9a55"></a>

## string_records — string_records / e8f2f5384c1b / 2

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
- [Property reference](data-sources--data_group--reference--group-001.md#canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81)
- string_records

<a id="canonical-cd3e15185134ca4c9159fef732c5790c62a4d4803b8908cc5b411a8208de5655"></a>

Type: `"single"`. Computed.

Configuration parameter for string records.

Upstream description:

Data group with strings record List.

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

<a id="canonical-da7c4770fb4a511db4f6783441874fa8c523231ead7b842302d0fc0cf6fedfc4"></a>

## Direct properties — string_records / e8f2f5384c1b / 3

<a id="canonical-b1495cc4201660ce8372c97641687d5e9a052ae5bdce4514ccad02989269cf4b"></a>

<a id="canonical-f31cb73913e7d25935983b56fa1d371b582c50effd0d1f302a00425f6f4ae058"></a>

## records property — string_records / e8f2f5384c1b / 4

Type: `["map", "string"]`. Computed.

String records. Configuration parameter for records

Upstream description:

Configuration parameter for records

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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

<a id="canonical-11297f767a937a34301725fa9c0577e025cfe83fdcb9602643abb0756ade9b59"></a>

## Next pages — string_records / e8f2f5384c1b / 5

- [Property reference](data-sources--data_group--reference--group-001.md#canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81)
- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
