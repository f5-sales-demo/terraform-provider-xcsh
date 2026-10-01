---
page_title: "xcsh_dc_cluster_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group reference."
---

# xcsh_dc_cluster_group reference

<a id="canonical-44fdcb354fbb0d410aa13e0c7edd48a9c9ad808c808a37edef7204a79b332e4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-745726831e3c02ef8dcbf376be679c806e828e207ef2c0d58f93e492f13abb84"></a>

## Property reference — Property reference / 9d7a2de59e76 / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
- Property reference

<a id="canonical-0a1257a7c0df4d58ac0b61b588190eeb9e5dc2ca481a55ff3afc6c283dce95a3"></a>

## Direct properties — Property reference / 9d7a2de59e76 / 3

<a id="canonical-61d69b9d7cce0754922db7a673c59752515684983b03559d59ea8d350b6f8f3a"></a>

<a id="canonical-ac4ceb17850a6182d43765b1183b3e428a2dd27c4c19b7c91d2a3f3a8eb73a68"></a>

## annotations property — Property reference / 9d7a2de59e76 / 4

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

<a id="canonical-f460a56bafb6b3045b29aa9d402fc5eb2112aab97b621bcc9557694089845bdf"></a>

<a id="canonical-5b35d67b2ae0738ef4e4f92c80488620c3cccd64ab0f2e9da6adba75f859e9d0"></a>

## description property — Property reference / 9d7a2de59e76 / 5

Type: `"string"`. Computed.

Description of the DcClusterGroup.

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

<a id="canonical-7eb0bb34d7c6933b7c4f4e15e79fbf10876ee5b38bf28f064864f8f4301289e4"></a>

<a id="canonical-6307adef386c06b1a34dafb88ce5d11e6c6e98c4cc9934980e694b0d2cfa1408"></a>

## id property — Property reference / 9d7a2de59e76 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2796a86acf990d8a415036f571f4adafaa5927dc0c874c09d4cd2d1b935b2585"></a>

<a id="canonical-bf2af3ce35136fd56c70e948cb3fb5dae1477c672d3440ad8ad8098c2650263b"></a>

## labels property — Property reference / 9d7a2de59e76 / 7

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

<a id="canonical-ff8db875f4ed154cc1652f9267cb5dbee917939370066f3bcdf7dd8c2a4865fe"></a>

<a id="canonical-871dc19a3b14b41c42106197851f4d3a5396994b7f781ff4d1d210d1ad5449ef"></a>

## name property — Property reference / 9d7a2de59e76 / 8

Type: `"string"`. Required.

Name of the DcClusterGroup.

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

<a id="canonical-c4f77125199f06ee30b824fb6d39add9cfd065f02cc033a5fc28cc7e504d6b0e"></a>

<a id="canonical-be980a8715c2bc7635075631ac979f2635bd5ff6ad22baa3ea3f1c256d40c021"></a>

## namespace property — Property reference / 9d7a2de59e76 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DcClusterGroup exists.

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

- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-99f48f996b2edeb1374646a3186a5162802186db69e5a2de1f8a6a4e3a3c383a): complete subsection reference.

<a id="canonical-14b3deeb8d3af00d6a425ed840d8176f01869023bb8164a57d1313b5abdadd17"></a>

## All schema paths — Property reference / 9d7a2de59e76 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dc_cluster_group--reference--group-001.md#canonical-61d69b9d7cce0754922db7a673c59752515684983b03559d59ea8d350b6f8f3a) |
| `description` | [description](data-sources--dc_cluster_group--reference--group-001.md#canonical-f460a56bafb6b3045b29aa9d402fc5eb2112aab97b621bcc9557694089845bdf) |
| `id` | [id](data-sources--dc_cluster_group--reference--group-001.md#canonical-7eb0bb34d7c6933b7c4f4e15e79fbf10876ee5b38bf28f064864f8f4301289e4) |
| `labels` | [labels](data-sources--dc_cluster_group--reference--group-001.md#canonical-2796a86acf990d8a415036f571f4adafaa5927dc0c874c09d4cd2d1b935b2585) |
| `name` | [name](data-sources--dc_cluster_group--reference--group-001.md#canonical-ff8db875f4ed154cc1652f9267cb5dbee917939370066f3bcdf7dd8c2a4865fe) |
| `namespace` | [namespace](data-sources--dc_cluster_group--reference--group-001.md#canonical-c4f77125199f06ee30b824fb6d39add9cfd065f02cc033a5fc28cc7e504d6b0e) |
| `type` | [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-d5293b3c693a1a901b73d4fce0dca2a6ed3fa371675a41435e659464fa5c3fa8) |
| `type.control_and_data_plane_mesh` | [type.control_and_data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-9b172e0748ea9c8d0ea648e64b57f24787ebc389331abafebf7a6c5b05323610) |
| `type.data_plane_mesh` | [type.data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-240a4d693daf36d644d2fe1da6630944bb6f03c5198854b0582278859f5ef685) |

<a id="canonical-11ff0a00833afcfa75b74eeed91235e9c907808cf39031473cff28b278225658"></a>

## Next pages — Property reference / 9d7a2de59e76 / 11

- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-99f48f996b2edeb1374646a3186a5162802186db69e5a2de1f8a6a4e3a3c383a)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)

<a id="canonical-99f48f996b2edeb1374646a3186a5162802186db69e5a2de1f8a6a4e3a3c383a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a2136e50cd3a1b95a4520c98428f6b1e1d8635818ad5894588f0492435f8109"></a>

## type — type / d9dd194376f4 / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
- [Property reference](data-sources--dc_cluster_group--reference--group-001.md#canonical-44fdcb354fbb0d410aa13e0c7edd48a9c9ad808c808a37edef7204a79b332e4b)
- type

<a id="canonical-d5293b3c693a1a901b73d4fce0dca2a6ed3fa371675a41435e659464fa5c3fa8"></a>

Type: `"single"`. Computed.

DC Cluster Group Mesh Type. Details of DC Cluster Group Mesh Type.

Upstream description:

Details of DC Cluster Group Mesh Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

<a id="canonical-4e82773b3b5d5d35096b48b71bf7540100b8afb335d1ac088ff278fa2a1834c9"></a>

## Direct properties — type / d9dd194376f4 / 3

- [control_and_data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-7975095d19defeb92df7926056319dc3612da3cdba0bab7bee2da452d76d05fc): complete subsection reference.

- [data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-5cb520e76e304a7898fda034f7e92ba514306671446745f8b05316d7c93d7b66): complete subsection reference.

<a id="canonical-320ec4de662e4d19c5e3bc5e9ee7f33a0f9a9de54c2eb6c3c39efa6245647595"></a>

## Next pages — type / d9dd194376f4 / 4

- [type.control_and_data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-7975095d19defeb92df7926056319dc3612da3cdba0bab7bee2da452d76d05fc)
- [type.data_plane_mesh](data-sources--dc_cluster_group--reference--group-001.md#canonical-5cb520e76e304a7898fda034f7e92ba514306671446745f8b05316d7c93d7b66)
- [Property reference](data-sources--dc_cluster_group--reference--group-001.md#canonical-44fdcb354fbb0d410aa13e0c7edd48a9c9ad808c808a37edef7204a79b332e4b)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)

<a id="canonical-7975095d19defeb92df7926056319dc3612da3cdba0bab7bee2da452d76d05fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee8c009d6a49a352faee2ef7b88aa9cce4eb6846dd473bff982e19099ff564dc"></a>

## type.control_and_data_plane_mesh — type.control_and_data_plane_mesh / 72183d308f9d / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
- [Property reference](data-sources--dc_cluster_group--reference--group-001.md#canonical-44fdcb354fbb0d410aa13e0c7edd48a9c9ad808c808a37edef7204a79b332e4b)
- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-99f48f996b2edeb1374646a3186a5162802186db69e5a2de1f8a6a4e3a3c383a)
- type.control_and_data_plane_mesh

<a id="canonical-9b172e0748ea9c8d0ea648e64b57f24787ebc389331abafebf7a6c5b05323610"></a>

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

<a id="canonical-e6135410d65d5ad8fd5b15acacc57f1a6b7432d023ba837850e689c5b365f422"></a>

## Direct properties — type.control_and_data_plane_mesh / 72183d308f9d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fef9fab2c53b2c14c1382e82fabd208743a3fb67b081c99989bfcb309cb5e9f6"></a>

## Next pages — type.control_and_data_plane_mesh / 72183d308f9d / 4

- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-99f48f996b2edeb1374646a3186a5162802186db69e5a2de1f8a6a4e3a3c383a)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)

<a id="canonical-5cb520e76e304a7898fda034f7e92ba514306671446745f8b05316d7c93d7b66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7743aca6b64a0532546ea662d225060e7a83e7c50689e843f8872799ab5d461"></a>

## type.data_plane_mesh — type.data_plane_mesh / 8baacfb3700f / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
- [Property reference](data-sources--dc_cluster_group--reference--group-001.md#canonical-44fdcb354fbb0d410aa13e0c7edd48a9c9ad808c808a37edef7204a79b332e4b)
- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-99f48f996b2edeb1374646a3186a5162802186db69e5a2de1f8a6a4e3a3c383a)
- type.data_plane_mesh

<a id="canonical-240a4d693daf36d644d2fe1da6630944bb6f03c5198854b0582278859f5ef685"></a>

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

<a id="canonical-5f39942f513c533cb6dad3a52815b58b527fec7e3f801f1c754754bfdf454245"></a>

## Direct properties — type.data_plane_mesh / 8baacfb3700f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-daed5a8b601f79fbc3ce459dfb86f68c9e6de820042a9ce32c3462af41f5ba72"></a>

## Next pages — type.data_plane_mesh / 8baacfb3700f / 4

- [type](data-sources--dc_cluster_group--reference--group-001.md#canonical-99f48f996b2edeb1374646a3186a5162802186db69e5a2de1f8a6a4e3a3c383a)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7)
