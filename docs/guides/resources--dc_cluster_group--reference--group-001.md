---
page_title: "xcsh_dc_cluster_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group reference."
---

# xcsh_dc_cluster_group reference

<a id="canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c62278f834009fb2da3b863ac975d5350a911d2f10b700d5f27844cf6215bab"></a>

## Property reference — Property reference / 19961649a2c1 / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
- Property reference

<a id="canonical-313b6a141c8f90181168496d57b98bbb3110ac4cc8e2bb819bed104cf5e53dad"></a>

## Direct properties — Property reference / 19961649a2c1 / 3

<a id="canonical-1fafcdf648e9ef32f71372fa8a11472194017dea1a53ad441fc39334f6711d6c"></a>

<a id="canonical-d4c3799f301ca8d87d178a54c04a6962e9af71a42c4abe21c547c5946ce557e1"></a>

## annotations property — Property reference / 19961649a2c1 / 4

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

<a id="canonical-2e946a96d0f89f2829abe4eaf20b8ba9a204fc261ed03aa6f9df56bb75fabbed"></a>

<a id="canonical-4c87bf63829fb282a7409e5b0ff51dee343fc29643c534257c43929eada81eb5"></a>

## description property — Property reference / 19961649a2c1 / 5

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

<a id="canonical-ff8cbc11f0a5fc35208846f0540d44351e8828b5857bc788c78fa94b50c4c1d6"></a>

<a id="canonical-08edb7577fa7a7581cad50bcac915bc29e405c766b5918a68f7fd909516154bf"></a>

## disable property — Property reference / 19961649a2c1 / 6

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

<a id="canonical-55d59a6f6711fdb0989709118918f8742d58d77a07dbe4dd9467275cf63b4340"></a>

<a id="canonical-3e7f7f893beefd543d29969c170ce46a1398ad7e3b6a81523e8ccc77841d40e4"></a>

## id property — Property reference / 19961649a2c1 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-8dede1a6d12e48b349b6686b4804b6651e281be8a632af47526bbc488e8b2af9"></a>

<a id="canonical-08550330d4fdeb1e55ef3e088f72c57bc839b728e0d560f8087c76e96f8ece0d"></a>

## labels property — Property reference / 19961649a2c1 / 8

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

<a id="canonical-68eff493a35fdd1698ac2fb9f7fbacc437bf281ac10158426260809846c44768"></a>

<a id="canonical-fd83ccb20d50f7f7f03141bf387a630736fea6cd6194877693030e032b32e3f7"></a>

## name property — Property reference / 19961649a2c1 / 9

Type: `"string"`. Required.

Name of the Dc Cluster Group. Must be unique within the namespace.

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

<a id="canonical-5937e0a60e69dcdda81b1809f7055f246d875a23f4b8764cde93ae9b9077e840"></a>

<a id="canonical-46478053daa2d5683b527bb740e15612b94c97436a4efd68243bf162c7fbd2a5"></a>

## namespace property — Property reference / 19961649a2c1 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Dc Cluster Group. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [timeouts](resources--dc_cluster_group--reference--group-001.md#canonical-e93ce1c02e37015bd5509a140fa46c5689972a9e79ab4e4c0069e939e8d1e986): complete subsection reference.

- [type](resources--dc_cluster_group--reference--group-001.md#canonical-d478544eb3933ee81badc565a66fb52287f88551fb794c4562d4177dd00e75a9): complete subsection reference.

<a id="canonical-534669ee9de533c5ff345faa3f870d13dc6c8dfa104e6fb8b35d66f2bf3cbefb"></a>

## All schema paths — Property reference / 19961649a2c1 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dc_cluster_group--reference--group-001.md#canonical-1fafcdf648e9ef32f71372fa8a11472194017dea1a53ad441fc39334f6711d6c) |
| `description` | [description](resources--dc_cluster_group--reference--group-001.md#canonical-2e946a96d0f89f2829abe4eaf20b8ba9a204fc261ed03aa6f9df56bb75fabbed) |
| `disable` | [disable](resources--dc_cluster_group--reference--group-001.md#canonical-ff8cbc11f0a5fc35208846f0540d44351e8828b5857bc788c78fa94b50c4c1d6) |
| `id` | [id](resources--dc_cluster_group--reference--group-001.md#canonical-55d59a6f6711fdb0989709118918f8742d58d77a07dbe4dd9467275cf63b4340) |
| `labels` | [labels](resources--dc_cluster_group--reference--group-001.md#canonical-8dede1a6d12e48b349b6686b4804b6651e281be8a632af47526bbc488e8b2af9) |
| `name` | [name](resources--dc_cluster_group--reference--group-001.md#canonical-68eff493a35fdd1698ac2fb9f7fbacc437bf281ac10158426260809846c44768) |
| `namespace` | [namespace](resources--dc_cluster_group--reference--group-001.md#canonical-5937e0a60e69dcdda81b1809f7055f246d875a23f4b8764cde93ae9b9077e840) |
| `timeouts` | [timeouts](resources--dc_cluster_group--reference--group-001.md#canonical-459e5c6b18bb9d1390e4148f76f5407d19252965a7aa325757d582a37f128052) |
| `timeouts.create` | [timeouts.create](resources--dc_cluster_group--reference--group-001.md#canonical-604f41eb7db3c63680b01d2fe9581e7232385f278e4d1c60009d37635e36e65e) |
| `timeouts.delete` | [timeouts.delete](resources--dc_cluster_group--reference--group-001.md#canonical-50d150df5f30ee9522fe29fae68c1dd5f415417117a94e62be34e23a65db9082) |
| `timeouts.read` | [timeouts.read](resources--dc_cluster_group--reference--group-001.md#canonical-85ffa5c59262006798419ee6bcaf77e74ee3025e9a496c13b640f65382ca2566) |
| `timeouts.update` | [timeouts.update](resources--dc_cluster_group--reference--group-001.md#canonical-a0202f313060b143aba6b76938e6631bcc616d56ddf64d270c79de79031c50c2) |
| `type` | [type](resources--dc_cluster_group--reference--group-001.md#canonical-e4bf5c31e6fd2f1764b47115ce4836812329ffee0376d0dc04a9719f6ba1f56e) |
| `type.control_and_data_plane_mesh` | [type.control_and_data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-5c9b87a949823265f1a0f95441f72195d9c0e53ed7a6d2e6754b838a8e07921d) |
| `type.data_plane_mesh` | [type.data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-ee152879ff4232b58bfee6e6b40c1dda46baf44f478782af17fe14401c8b04b0) |

<a id="canonical-eab3797c2e99b27ddab6df1ee1d9a31a0fde580e3ecf61997b0e915cdba84908"></a>

## Next pages — Property reference / 19961649a2c1 / 12

- [timeouts](resources--dc_cluster_group--reference--group-001.md#canonical-e93ce1c02e37015bd5509a140fa46c5689972a9e79ab4e4c0069e939e8d1e986)
- [type](resources--dc_cluster_group--reference--group-001.md#canonical-d478544eb3933ee81badc565a66fb52287f88551fb794c4562d4177dd00e75a9)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)

<a id="canonical-e93ce1c02e37015bd5509a140fa46c5689972a9e79ab4e4c0069e939e8d1e986"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae67ea2f9a67458e1ee3d1519f567be2f1d0f2306ffbd246e4ee0476a05c4984"></a>

## timeouts — timeouts / afbd206ff97b / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9)
- timeouts

<a id="canonical-459e5c6b18bb9d1390e4148f76f5407d19252965a7aa325757d582a37f128052"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-124323f7fcf88205105dc380ab64419e82566f0eacf5e3f14f211050b4e0a3f3"></a>

## Direct properties — timeouts / afbd206ff97b / 3

<a id="canonical-604f41eb7db3c63680b01d2fe9581e7232385f278e4d1c60009d37635e36e65e"></a>

<a id="canonical-cfac40db609130c3b53ea5175bf9ee33531085320499c713998562ef76416375"></a>

## create property — timeouts / afbd206ff97b / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-50d150df5f30ee9522fe29fae68c1dd5f415417117a94e62be34e23a65db9082"></a>

<a id="canonical-1406b1f2183eee50585eeed28042c1dcd3fff8f6a42066c755c3fa4c1b5f4b6a"></a>

## delete property — timeouts / afbd206ff97b / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-85ffa5c59262006798419ee6bcaf77e74ee3025e9a496c13b640f65382ca2566"></a>

<a id="canonical-10b5ab627eb14ff3501dcc05e19327fac482009ec72348d76de3f2f7b0073a04"></a>

## read property — timeouts / afbd206ff97b / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a0202f313060b143aba6b76938e6631bcc616d56ddf64d270c79de79031c50c2"></a>

<a id="canonical-6dcdb805a360a3788ab9024ddd43d2ba2ef779bad653968ba05def70fdbc6621"></a>

## update property — timeouts / afbd206ff97b / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a0107059c9c1262e16e44cccb1da511ef597bbd76ff03d21da7a51da83c27111"></a>

## Next pages — timeouts / afbd206ff97b / 8

- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)

<a id="canonical-d478544eb3933ee81badc565a66fb52287f88551fb794c4562d4177dd00e75a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da3e37fc12c08bfdb34fd53e18b0823893d24a6cff4fb970ca18d362b2bd11a1"></a>

## type — type / d21ea912406b / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9)
- type

<a id="canonical-e4bf5c31e6fd2f1764b47115ce4836812329ffee0376d0dc04a9719f6ba1f56e"></a>

Type: `"object"`. single nested block, Optional.

DC Cluster Group Mesh Type. Details of DC Cluster Group Mesh Type.

Upstream description:

Details of DC Cluster Group Mesh Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("control_and_data_plane_mesh",
    "data_plane_mesh")}
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
  "x-ves-oneof-field-dc_cluster_group_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
type {
  # Configure direct properties listed below.
}
```

<a id="canonical-bbd2bf765cc378e7ba4acc443ae9ddd63066552f926910d8bf50ddc0f4103364"></a>

## Direct properties — type / d21ea912406b / 3

- [control_and_data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-20d65b51ed3620dfd950d5e48be6f71bc2f8a52294b3e465344174891a29eab5): complete subsection reference.

- [data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-b8b9e036445c30607529a8daa5de32a9d668f278f78c02bd7866f269f98ace27): complete subsection reference.

<a id="canonical-86aee97fcf76fddef7b5b206e9dfeddbbdfa64a5831b9c7f8b11bbbf9d75c01a"></a>

## Next pages — type / d21ea912406b / 4

- [type.control_and_data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-20d65b51ed3620dfd950d5e48be6f71bc2f8a52294b3e465344174891a29eab5)
- [type.data_plane_mesh](resources--dc_cluster_group--reference--group-001.md#canonical-b8b9e036445c30607529a8daa5de32a9d668f278f78c02bd7866f269f98ace27)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)

<a id="canonical-20d65b51ed3620dfd950d5e48be6f71bc2f8a52294b3e465344174891a29eab5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c287af66c037d24a85bd645ead9a303b1b7a53fd13954db3212a0e17fabbf1d"></a>

## type.control_and_data_plane_mesh — type.control_and_data_plane_mesh / 5423d3ee8a9b / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9)
- [type](resources--dc_cluster_group--reference--group-001.md#canonical-d478544eb3933ee81badc565a66fb52287f88551fb794c4562d4177dd00e75a9)
- type.control_and_data_plane_mesh

<a id="canonical-5c9b87a949823265f1a0f95441f72195d9c0e53ed7a6d2e6754b838a8e07921d"></a>

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
control_and_data_plane_mesh = {}
```

<a id="canonical-871cef9b48733a8d14aeae4c4297167d66105210ed5b9230ff508d5274ec0b60"></a>

## Direct properties — type.control_and_data_plane_mesh / 5423d3ee8a9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-083537c234871a20468e0f5e374ec543896d7277668776341c72003ebf53eb08"></a>

## Next pages — type.control_and_data_plane_mesh / 5423d3ee8a9b / 4

- [type](resources--dc_cluster_group--reference--group-001.md#canonical-d478544eb3933ee81badc565a66fb52287f88551fb794c4562d4177dd00e75a9)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)

<a id="canonical-b8b9e036445c30607529a8daa5de32a9d668f278f78c02bd7866f269f98ace27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9292c093b2f84c5b3279892c0c154be0961fcf6ac9b1b362eaa09821b9289b9e"></a>

## type.data_plane_mesh — type.data_plane_mesh / 4c2202e7bcb9 / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
- [Property reference](resources--dc_cluster_group--reference--group-001.md#canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9)
- [type](resources--dc_cluster_group--reference--group-001.md#canonical-d478544eb3933ee81badc565a66fb52287f88551fb794c4562d4177dd00e75a9)
- type.data_plane_mesh

<a id="canonical-ee152879ff4232b58bfee6e6b40c1dda46baf44f478782af17fe14401c8b04b0"></a>

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
data_plane_mesh = {}
```

<a id="canonical-52b918e129960af45cbbac028bb76e69c0063ba67f5f0cea7172632d174de295"></a>

## Direct properties — type.data_plane_mesh / 4c2202e7bcb9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-703c73bc9c30a57b0442fa82643567654e7493e0c789ada2bfbbe73365be87b4"></a>

## Next pages — type.data_plane_mesh / 4c2202e7bcb9 / 4

- [type](resources--dc_cluster_group--reference--group-001.md#canonical-d478544eb3933ee81badc565a66fb52287f88551fb794c4562d4177dd00e75a9)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
