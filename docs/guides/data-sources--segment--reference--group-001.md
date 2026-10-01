---
page_title: "xcsh_segment reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment reference."
---

# xcsh_segment reference

<a id="canonical-542cdf7a24561b3f7a476758c7d287b3eb9814b546e6df99ab410ec927020b6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95064b027b444408c6595477712a9307aa2eba72d091c467c228705938c1d8ad"></a>

## Property reference — Property reference / 5a395c09aec3 / 2

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)
- Property reference

<a id="canonical-2e8332782c5b181533b82c3367564434e8bbe9c4a2d925541e97cdbcc0c74fc1"></a>

## Direct properties — Property reference / 5a395c09aec3 / 3

<a id="canonical-cd679e1afaa60513ad0f838a3a00f188a1132e3e6437d00c54b5e500044fc5d3"></a>

<a id="canonical-3f6a951db76cf1e4ed2c3acbd24807b667b3f185bd35159958be30d4370a4216"></a>

## annotations property — Property reference / 5a395c09aec3 / 4

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

<a id="canonical-2c877193fcd7d1310f0345e6b94c835e699d434384e39bc94cda98d53516eaab"></a>

<a id="canonical-6e58ff30dded0cd28285c6ec986c84c402fa49fe7c35aa10457ae92de4f46f72"></a>

## description property — Property reference / 5a395c09aec3 / 5

Type: `"string"`. Computed.

Description of the Segment.

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

- [disable_spec](data-sources--segment--reference--group-001.md#canonical-91c019419654b3cb4ae6aca1437f56f89a6acf02abf55a2b84ef760f8d34f40d): complete subsection reference.

- [enable](data-sources--segment--reference--group-001.md#canonical-7659587c492b8ced9343e8a7cec81e0b9d9047f6f8f31b4ea5acf17875904a26): complete subsection reference.

<a id="canonical-305404dec3d6407619a699eb32bb1629a5b086debc0ad55af72b2362b5b48ee1"></a>

<a id="canonical-29a6fbe456858aa2486e565f3265be7cd991dd7d7e15a41d3e157193c8e26a9c"></a>

## id property — Property reference / 5a395c09aec3 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-88da734f23ac2c999201ff7659d2caddaf01c52761d5f7e7592cae3e2ce158c8"></a>

<a id="canonical-bd4a42db9645b6ad443e9461eaf6d7e116527091296083323dba44dd62770288"></a>

## labels property — Property reference / 5a395c09aec3 / 7

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

<a id="canonical-c0bd91bfee3d92f07367e6cf688c6150e6a02b3605545bdfafca34a50f3d6a1f"></a>

<a id="canonical-f8a6e68417196af1f8b9612ade829e1106a7cbcc1ed91ae8927cc840561e7dc2"></a>

## name property — Property reference / 5a395c09aec3 / 8

Type: `"string"`. Required.

Name of the Segment.

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

<a id="canonical-6e84770b3edc0aa6d75845b64b1a148a02d941798e89b7c9eb270fcd12403b0c"></a>

<a id="canonical-8163d3d252d7e51013c91733381c3f714155675df78b2fce317a26acb836da1a"></a>

## namespace property — Property reference / 5a395c09aec3 / 9

Type: `"string"`. Optional, Computed.

Namespace where the Segment exists.

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

<a id="canonical-63649b63f4241edf5314ae9b38936c9ab25e214288376aba1ad7903675d64465"></a>

## All schema paths — Property reference / 5a395c09aec3 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--segment--reference--group-001.md#canonical-cd679e1afaa60513ad0f838a3a00f188a1132e3e6437d00c54b5e500044fc5d3) |
| `description` | [description](data-sources--segment--reference--group-001.md#canonical-2c877193fcd7d1310f0345e6b94c835e699d434384e39bc94cda98d53516eaab) |
| `disable_spec` | [disable_spec](data-sources--segment--reference--group-001.md#canonical-f985a91263d9fe2e6babd7c77cc7fb83c721e273281c5f0d94113548fdea28df) |
| `enable` | [enable](data-sources--segment--reference--group-001.md#canonical-e1754ee91d95a02a5e9d3af9c3d9466affc051bc41e63a2e3fcb842ea37adf57) |
| `id` | [id](data-sources--segment--reference--group-001.md#canonical-305404dec3d6407619a699eb32bb1629a5b086debc0ad55af72b2362b5b48ee1) |
| `labels` | [labels](data-sources--segment--reference--group-001.md#canonical-88da734f23ac2c999201ff7659d2caddaf01c52761d5f7e7592cae3e2ce158c8) |
| `name` | [name](data-sources--segment--reference--group-001.md#canonical-c0bd91bfee3d92f07367e6cf688c6150e6a02b3605545bdfafca34a50f3d6a1f) |
| `namespace` | [namespace](data-sources--segment--reference--group-001.md#canonical-6e84770b3edc0aa6d75845b64b1a148a02d941798e89b7c9eb270fcd12403b0c) |

<a id="canonical-6e00e4cfee71c1cdfb8d50b442261055821ddf3525407c4bb8fdb74b9778c983"></a>

## Next pages — Property reference / 5a395c09aec3 / 11

- [disable_spec](data-sources--segment--reference--group-001.md#canonical-91c019419654b3cb4ae6aca1437f56f89a6acf02abf55a2b84ef760f8d34f40d)
- [enable](data-sources--segment--reference--group-001.md#canonical-7659587c492b8ced9343e8a7cec81e0b9d9047f6f8f31b4ea5acf17875904a26)
- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)

<a id="canonical-91c019419654b3cb4ae6aca1437f56f89a6acf02abf55a2b84ef760f8d34f40d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb9ee09544a9743e7b5f6be1b71264fba18651d955a2b19e071bdd80ba4137c7"></a>

## disable_spec — disable_spec / e9bd7881edab / 2

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)
- [Property reference](data-sources--segment--reference--group-001.md#canonical-542cdf7a24561b3f7a476758c7d287b3eb9814b546e6df99ab410ec927020b6f)
- disable_spec

<a id="canonical-f985a91263d9fe2e6babd7c77cc7fb83c721e273281c5f0d94113548fdea28df"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable, enable\] Enable this option

OneOf alternatives in this subsection:

- `disable`
- [enable](data-sources--segment--reference--group-001.md#canonical-e1754ee91d95a02a5e9d3af9c3d9466affc051bc41e63a2e3fcb842ea37adf57)

Select alternatives according to the provider validators above.

<a id="canonical-5d5d08827ca97a03c51ba455ba00ae5b35d098f447c49af5375d699a2bf0a94e"></a>

## Direct properties — disable_spec / e9bd7881edab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03c0a636e46a02c7d347851df229877f3206e3fbb5a0245ce4a5eb0b2c345923"></a>

## Next pages — disable_spec / e9bd7881edab / 4

- [Property reference](data-sources--segment--reference--group-001.md#canonical-542cdf7a24561b3f7a476758c7d287b3eb9814b546e6df99ab410ec927020b6f)
- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)

<a id="canonical-7659587c492b8ced9343e8a7cec81e0b9d9047f6f8f31b4ea5acf17875904a26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b4b3f8800d1b2263dfd73e649f6a13a7b60074439a53d37e2087ba3ce0677a1"></a>

## enable — enable / d14e5a8a31c9 / 2

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)
- [Property reference](data-sources--segment--reference--group-001.md#canonical-542cdf7a24561b3f7a476758c7d287b3eb9814b546e6df99ab410ec927020b6f)
- enable

<a id="canonical-e1754ee91d95a02a5e9d3af9c3d9466affc051bc41e63a2e3fcb842ea37adf57"></a>

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

<a id="canonical-3b6ea76b756884797b171b2e4a9a3402f0bea5144592c1a01eaffa78ed1662ec"></a>

## Direct properties — enable / d14e5a8a31c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39121b4c36051492717221060f3f818c464a5d12206b7db2048b65f388b38594"></a>

## Next pages — enable / d14e5a8a31c9 / 4

- [Property reference](data-sources--segment--reference--group-001.md#canonical-542cdf7a24561b3f7a476758c7d287b3eb9814b546e6df99ab410ec927020b6f)
- [xcsh_segment](../data-sources/segment.md#canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650)
