---
page_title: "xcsh_site_mesh_group reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group reference."
---

# xcsh_site_mesh_group reference

<a id="canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6851c4b9efc498454e22b2ac2222f63ec857343f712c1ff0fe297abc0f6ee153"></a>

## Property reference — Property reference / 7af8da5159d7 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- Property reference

<a id="canonical-17315cba7958f6876ea42677c04dd028a97231068aa9f752a70727013fcf5862"></a>

## Direct properties — Property reference / 7af8da5159d7 / 3

<a id="canonical-5da85d2cc3ec006f722c233e5d56052a8372d85b3ed997e5bfb74eb069f77729"></a>

<a id="canonical-8e599e0163ad5557c2a41595c70ce86c352e540371118a2f4d9e4b9d06a04b1c"></a>

## annotations property — Property reference / 7af8da5159d7 / 4

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

- [bfd_disabled](resources--site_mesh_group--reference--group-001.md#canonical-45bbb5d96de227952d2ad634a4e53dd21d7ce874373fd69d893cf28c1124e6e2): complete subsection reference.

- [bfd_enabled](resources--site_mesh_group--reference--group-001.md#canonical-0cd9eb96b353fae38fcab75c9bd0b3ec774c7113052ea4f8fc3f058e2df37292): complete subsection reference.

<a id="canonical-4e0b4c1f4bd55824fa4c22534e7794db8647c3191d679837700d9a9f71ce20a4"></a>

<a id="canonical-7e22099f05ed89424b89f2a196b1f8fe88043e2049d5f80aff1208c31d373f4e"></a>

## description property — Property reference / 7af8da5159d7 / 5

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

<a id="canonical-0355cd68fa5f9a74ea48e93cd57711e80baad2c13aecc9f7eb4bf8244414a276"></a>

<a id="canonical-1b5a5b9afde1f004edb9e98f93c53699e1e63ce46ff367483b0e84d644ed9603"></a>

## disable property — Property reference / 7af8da5159d7 / 6

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

- [disable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-b04b5d135e3d6f152e44e74f8486eea8ac4662b70ccedcdf95e01a93baa3cd5c): complete subsection reference.

- [enable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-e4220a2e1835c209cab9ace1e2cb0b7ea99cd0e1150da8337e99e4a641495ee1): complete subsection reference.

- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8d064e8b4e9544faa84d785a97da2c70dcedcd1866290e0e00925efa7fec05a5): complete subsection reference.

- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-457c1d87b67af5367ecaa4ac5a73da4519e7370c15bff7fc9617b0b8a4e3ba16): complete subsection reference.

<a id="canonical-3854a485bc059e8acaf72d4cec2fc818d977babf4bee02d6e94b32d932a6c146"></a>

<a id="canonical-7604b2a8667da2781c18b0b434278973bb19d3d25b6b728c9415b541e9f3f683"></a>

## id property — Property reference / 7af8da5159d7 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-fcf7d84c74360451d678207ce04e36504b91f3631550bc4a88326743bbd2c56d"></a>

<a id="canonical-4fec0be7710c9ca2f68d06410d42fb2f355fb4091729cd985da25edb12395cf8"></a>

## labels property — Property reference / 7af8da5159d7 / 8

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

<a id="canonical-d51ec5c9f46567811f954e83d90e426dcea42fd59a3efbd25d50536bef80df3f"></a>

<a id="canonical-d9ae59577873857f575cbe8a8eac122dd443a821d3a192b39df18a71949b8922"></a>

## name property — Property reference / 7af8da5159d7 / 9

Type: `"string"`. Required.

Name of the Site Mesh Group. Must be unique within the namespace.

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

<a id="canonical-3930491e1c6d61d593c1c209c1b4e3e8cd7c6df09d93e01734b20f9d7dc1d9bb"></a>

<a id="canonical-7a79ba4ef0339add5a8dc2081e60d405dedbb1180172b128ebcfc548f9625511"></a>

## namespace property — Property reference / 7af8da5159d7 / 10

Type: `"string"`. Required.

Namespace where the Site Mesh Group is created.

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

- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759): complete subsection reference.

- [timeouts](resources--site_mesh_group--reference--group-001.md#canonical-0542d8e111378fd85fbf83a05a84465e981936714c770f3c62ad507b27c82d77): complete subsection reference.

- [virtual_site](resources--site_mesh_group--reference--group-001.md#canonical-68decd4772a312af42a6edde36a861c9725bcc37ee1cf23ac547d5e3b6f5eb00): complete subsection reference.

<a id="canonical-3f0b3242b3f32c556c635f28bc5e51a9fdd85b9a759d9076fad6c96d0c2bf368"></a>

## All schema paths — Property reference / 7af8da5159d7 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--site_mesh_group--reference--group-001.md#canonical-5da85d2cc3ec006f722c233e5d56052a8372d85b3ed997e5bfb74eb069f77729) |
| `bfd_disabled` | [bfd_disabled](resources--site_mesh_group--reference--group-001.md#canonical-11dd523a9760dc9d7ea2b0365d95287da95517aa7d0e120cda29ee61b13a5e2c) |
| `bfd_enabled` | [bfd_enabled](resources--site_mesh_group--reference--group-001.md#canonical-b91e37d4c5bcb24728445341616025600b7f0ea55cd804e5f7bd38e073945183) |
| `bfd_enabled.multiplier` | [bfd_enabled.multiplier](resources--site_mesh_group--reference--group-001.md#canonical-459246038480df69dcb24935831560fcafb5b634e84bdfd09514d505d8adbdaa) |
| `bfd_enabled.receive_interval_milliseconds` | [bfd_enabled.receive_interval_milliseconds](resources--site_mesh_group--reference--group-001.md#canonical-ba93b46c00cc42fdb97a21c1f1539274f51b01712d21dd4be44e76d7932b8e6d) |
| `bfd_enabled.transmit_interval_milliseconds` | [bfd_enabled.transmit_interval_milliseconds](resources--site_mesh_group--reference--group-001.md#canonical-806fe3e497a01b042616f156228974f7f964bc21efb662152d638a04eb435cdf) |
| `description` | [description](resources--site_mesh_group--reference--group-001.md#canonical-4e0b4c1f4bd55824fa4c22534e7794db8647c3191d679837700d9a9f71ce20a4) |
| `disable` | [disable](resources--site_mesh_group--reference--group-001.md#canonical-0355cd68fa5f9a74ea48e93cd57711e80baad2c13aecc9f7eb4bf8244414a276) |
| `disable_re_fallback` | [disable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-652eb6d4b136fc8f3da74b0857b00a4706d40f39dd23616ed2535c48d54dae21) |
| `enable_re_fallback` | [enable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-5d518332a8cff794427e6264f3fc46b91ed0c9338cd2fbd242efe5fb30958ff4) |
| `full_mesh` | [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-1ad3b7706acddf848f39c545835e9c75ec9999e8a47723b8e6f22481558674b4) |
| `full_mesh.control_and_data_plane_mesh` | [full_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-c235aaa8218e994c4d8936d7464307fda51b945cf44fee605ba0e04aacbb4fd9) |
| `full_mesh.data_plane_mesh` | [full_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8dfd4e34ea381c706a48eddd23d87dd26a1cf28f930a1895cf7e05d0d08ad14b) |
| `hub_mesh` | [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-e55a7c5d9e213d375305bb2089b9710bf37a00c384d1418cac4d9e3bfc86f914) |
| `hub_mesh.control_and_data_plane_mesh` | [hub_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-5831cc095e56f0d863fddb9e7562b2cec30edc6b31d3971675ec1d9914980290) |
| `hub_mesh.data_plane_mesh` | [hub_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2d1623f8935fab931a77722f5421665bd349158a0caa67c03562e3376a7623b9) |
| `id` | [id](resources--site_mesh_group--reference--group-001.md#canonical-3854a485bc059e8acaf72d4cec2fc818d977babf4bee02d6e94b32d932a6c146) |
| `labels` | [labels](resources--site_mesh_group--reference--group-001.md#canonical-fcf7d84c74360451d678207ce04e36504b91f3631550bc4a88326743bbd2c56d) |
| `name` | [name](resources--site_mesh_group--reference--group-001.md#canonical-d51ec5c9f46567811f954e83d90e426dcea42fd59a3efbd25d50536bef80df3f) |
| `namespace` | [namespace](resources--site_mesh_group--reference--group-001.md#canonical-3930491e1c6d61d593c1c209c1b4e3e8cd7c6df09d93e01734b20f9d7dc1d9bb) |
| `spoke_mesh` | [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-b93d470f538648d33f1eedd5beff41b21e6d3c4305cd72f0549e85915008d3ca) |
| `spoke_mesh.control_and_data_plane_mesh` | [spoke_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-e2dd58baa02ff24462e7e14d2f74c4ad139f697801d1784fdb60866cab16f477) |
| `spoke_mesh.data_plane_mesh` | [spoke_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8e3b0ac41986739a0b4a0e38bf741e960dbb7a077adaaa343ea20f35b4333e07) |
| `spoke_mesh.hub_mesh_group` | [spoke_mesh.hub_mesh_group](resources--site_mesh_group--reference--group-001.md#canonical-877e83a2be83ea49241e2316179c26ed847edd7c605d605e893d8e818a946683) |
| `spoke_mesh.hub_mesh_group.name` | [spoke_mesh.hub_mesh_group.name](resources--site_mesh_group--reference--group-001.md#canonical-33d80a4940a58a42d13acbfac43dc983e214976607fe16a757cbf4f2c021e7e2) |
| `spoke_mesh.hub_mesh_group.namespace` | [spoke_mesh.hub_mesh_group.namespace](resources--site_mesh_group--reference--group-001.md#canonical-5dcce9bf2735f4baeabdf16b3cb213b6df53d16429e32df4334fc854caa5f444) |
| `spoke_mesh.hub_mesh_group.tenant` | [spoke_mesh.hub_mesh_group.tenant](resources--site_mesh_group--reference--group-001.md#canonical-35ca7521e24b3bff8587a3c03218e146367ea962295492ce54bf61a68f6bc328) |
| `timeouts` | [timeouts](resources--site_mesh_group--reference--group-001.md#canonical-7edf4a61bef6426ddc46bdc883de1e09ff9d061d2ebfbb5b210458b0c7565144) |
| `timeouts.create` | [timeouts.create](resources--site_mesh_group--reference--group-001.md#canonical-f23b020ed9c1dc44434cc9680f8ddf854482357c2e53688a80181411b14ce541) |
| `timeouts.delete` | [timeouts.delete](resources--site_mesh_group--reference--group-001.md#canonical-063c4ceba9e1a2022949b3cd4aedf9884b4e31a581fe8f88f8f8273a763ee78f) |
| `timeouts.read` | [timeouts.read](resources--site_mesh_group--reference--group-001.md#canonical-737b82dea66707553be834c0f7a36f15edb74c3d14632da9d6052bb99e962cf0) |
| `timeouts.update` | [timeouts.update](resources--site_mesh_group--reference--group-001.md#canonical-79d911aeac33f92340a8ac2e1edfe260d534e10f3f76197e302b468c20e940ed) |
| `virtual_site` | [virtual_site](resources--site_mesh_group--reference--group-001.md#canonical-6fecdbc856b80e5d01714a2d5e2c05595b8ca633b5999eded04946133b83eecb) |
| `virtual_site.kind` | [virtual_site.kind](resources--site_mesh_group--reference--group-001.md#canonical-fff5b37b6f76f5075fe9771eee6cb4e397f2b2dc8820069e2ae76da31e0cc503) |
| `virtual_site.name` | [virtual_site.name](resources--site_mesh_group--reference--group-001.md#canonical-034b27cb7a3684ecdd51a3b3520a376a5ed33e880da87c0697ae460ab0a09f8f) |
| `virtual_site.namespace` | [virtual_site.namespace](resources--site_mesh_group--reference--group-001.md#canonical-351e675a0a66271bf24b95e44cf402123e5f3b1df12fc82151d819307ab7df60) |
| `virtual_site.tenant` | [virtual_site.tenant](resources--site_mesh_group--reference--group-001.md#canonical-e9aa50647bbafc6426cfffded3996614eeb66388b674127c43e7838b2db91484) |
| `virtual_site.uid` | [virtual_site.uid](resources--site_mesh_group--reference--group-001.md#canonical-310f4185401ab62f9ca2fa82b773283e546b7c95f3bf95895a040c2a1d094eb6) |

<a id="canonical-9f4340f15a957bbf210bc875d1196a78bd55d6c3d46065cdcc5af6102140b6f4"></a>

## Next pages — Property reference / 7af8da5159d7 / 12

- [bfd_disabled](resources--site_mesh_group--reference--group-001.md#canonical-45bbb5d96de227952d2ad634a4e53dd21d7ce874373fd69d893cf28c1124e6e2)
- [bfd_enabled](resources--site_mesh_group--reference--group-001.md#canonical-0cd9eb96b353fae38fcab75c9bd0b3ec774c7113052ea4f8fc3f058e2df37292)
- [disable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-b04b5d135e3d6f152e44e74f8486eea8ac4662b70ccedcdf95e01a93baa3cd5c)
- [enable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-e4220a2e1835c209cab9ace1e2cb0b7ea99cd0e1150da8337e99e4a641495ee1)
- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8d064e8b4e9544faa84d785a97da2c70dcedcd1866290e0e00925efa7fec05a5)
- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-457c1d87b67af5367ecaa4ac5a73da4519e7370c15bff7fc9617b0b8a4e3ba16)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759)
- [timeouts](resources--site_mesh_group--reference--group-001.md#canonical-0542d8e111378fd85fbf83a05a84465e981936714c770f3c62ad507b27c82d77)
- [virtual_site](resources--site_mesh_group--reference--group-001.md#canonical-68decd4772a312af42a6edde36a861c9725bcc37ee1cf23ac547d5e3b6f5eb00)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-45bbb5d96de227952d2ad634a4e53dd21d7ce874373fd69d893cf28c1124e6e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a8f43495173f951b4809f404507b43ed6b634c90658e97f6794c79cea06ed45"></a>

## bfd_disabled — bfd_disabled / 8a5327f7fdfe / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- bfd_disabled

<a id="canonical-11dd523a9760dc9d7ea2b0365d95287da95517aa7d0e120cda29ee61b13a5e2c"></a>

Type: `["object", {}]`. Optional.

\[OneOf: bfd\_disabled, bfd\_enabled\] Enable this option

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

OneOf alternatives in this subsection:

- [bfd_disabled](resources--site_mesh_group--reference--group-001.md#canonical-11dd523a9760dc9d7ea2b0365d95287da95517aa7d0e120cda29ee61b13a5e2c)
- [bfd_enabled](resources--site_mesh_group--reference--group-001.md#canonical-b91e37d4c5bcb24728445341616025600b7f0ea55cd804e5f7bd38e073945183)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bfd_disabled = {}
```

<a id="canonical-07ca63764168947b76bc788e359ad5814daa656f09803139b3ff342bd1350227"></a>

## Direct properties — bfd_disabled / 8a5327f7fdfe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ccd9e9fcfef0751660e8b3cbbf39591641a0f797b806b42ee6c3da3fc040afdd"></a>

## Next pages — bfd_disabled / 8a5327f7fdfe / 4

- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-0cd9eb96b353fae38fcab75c9bd0b3ec774c7113052ea4f8fc3f058e2df37292"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c9b9f1ea56ad7e2234c2fc2b9c5292a8eb2481880db0b95c73d208247e40059"></a>

## bfd_enabled — bfd_enabled / 5617fa0673ea / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- bfd_enabled

<a id="canonical-b91e37d4c5bcb24728445341616025600b7f0ea55cd804e5f7bd38e073945183"></a>

Type: `"object"`. single nested block, Optional.

BFD. BFD parameters.

Upstream description:

BFD parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("multiplier",
    "receive_interval_milliseconds",
    "transmit_interval_milliseconds")}
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
bfd_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-7dda6c826dacde0170e6855809f4217aa84841c0b694fbf86e651e9217bfd819"></a>

## Direct properties — bfd_enabled / 5617fa0673ea / 3

<a id="canonical-459246038480df69dcb24935831560fcafb5b634e84bdfd09514d505d8adbdaa"></a>

<a id="canonical-a7aa1db6cbe4bf601439799ee19e0ad0d66ebede089487b87cd8897049f60589"></a>

## multiplier property — bfd_enabled / 5617fa0673ea / 4

Type: `"number"`. Optional.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-ba93b46c00cc42fdb97a21c1f1539274f51b01712d21dd4be44e76d7932b8e6d"></a>

<a id="canonical-e309e2a38ebe0862a7d30d179903cb26b164615427754767d84764428b6c829e"></a>

## receive_interval_milliseconds property — bfd_enabled / 5617fa0673ea / 5

Type: `"number"`. Optional.

BFD receive interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-806fe3e497a01b042616f156228974f7f964bc21efb662152d638a04eb435cdf"></a>

<a id="canonical-470212175411d37ec146cb51c82688abb41d4a893e968cd431f29d7fa34a1da8"></a>

## transmit_interval_milliseconds property — bfd_enabled / 5617fa0673ea / 6

Type: `"number"`. Optional.

BFD transmit interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-813c7048ef5a36e4930195da5bb6cceaf5a8e810d85e464dbce2618c7ec4e9ab"></a>

## Next pages — bfd_enabled / 5617fa0673ea / 7

- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-b04b5d135e3d6f152e44e74f8486eea8ac4662b70ccedcdf95e01a93baa3cd5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf3d332f28d1a42a50e63e1a9332e8a0028d742449d583fd3fe68b476bfa62b4"></a>

## disable_re_fallback — disable_re_fallback / 386e38be2bc4 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- disable_re_fallback

<a id="canonical-652eb6d4b136fc8f3da74b0857b00a4706d40f39dd23616ed2535c48d54dae21"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_re\_fallback, enable\_re\_fallback; Default: disable\_re\_fallback\] Configuration
parameter for disable re fallback.

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

OneOf alternatives in this subsection:

- [disable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-652eb6d4b136fc8f3da74b0857b00a4706d40f39dd23616ed2535c48d54dae21)
- [enable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-5d518332a8cff794427e6264f3fc46b91ed0c9338cd2fbd242efe5fb30958ff4)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_re_fallback = {}
```

<a id="canonical-0a3f46b023713ebef83a1fc6ed7fecb252e4f0a0539913a4698e720bfb95b8d4"></a>

## Direct properties — disable_re_fallback / 386e38be2bc4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdb3d99f2c35f9399e1897141e081fcc0c835664140d55aee968597d4f5c3297"></a>

## Next pages — disable_re_fallback / 386e38be2bc4 / 4

- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-e4220a2e1835c209cab9ace1e2cb0b7ea99cd0e1150da8337e99e4a641495ee1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-980a969c4d257248c11cf98bc5e2395f41288d64cd03ae5af56862bccef65f6a"></a>

## enable_re_fallback — enable_re_fallback / fd36b7fa87a3 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- enable_re_fallback

<a id="canonical-5d518332a8cff794427e6264f3fc46b91ed0c9338cd2fbd242efe5fb30958ff4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable re fallback.

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
enable_re_fallback = {}
```

<a id="canonical-90b111988f2f0b870bad08c8077d824a8ccba0ff8b94507ee8520cfe9d898745"></a>

## Direct properties — enable_re_fallback / fd36b7fa87a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cca93e3bc72cca226d40d31d77a97070ccbff1d9dc4a4023c6cf7d92d733c24e"></a>

## Next pages — enable_re_fallback / fd36b7fa87a3 / 4

- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-8d064e8b4e9544faa84d785a97da2c70dcedcd1866290e0e00925efa7fec05a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2972b545ff648259cd5538ebfd82727a525e1234f89d9e95dadce931277376a6"></a>

## full_mesh — full_mesh / 2983206f40e2 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- full_mesh

<a id="canonical-1ad3b7706acddf848f39c545835e9c75ec9999e8a47723b8e6f22481558674b4"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: full\_mesh, hub\_mesh, spoke\_mesh\] Full Mesh. Details of Full Mesh Group Type.

Upstream description:

Details of Full Mesh Group Type.

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
  "x-ves-oneof-field-full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

OneOf alternatives in this subsection:

- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-1ad3b7706acddf848f39c545835e9c75ec9999e8a47723b8e6f22481558674b4)
- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-e55a7c5d9e213d375305bb2089b9710bf37a00c384d1418cac4d9e3bfc86f914)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-b93d470f538648d33f1eedd5beff41b21e6d3c4305cd72f0549e85915008d3ca)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
full_mesh {
  # Configure direct properties listed below.
}
```

<a id="canonical-bbbef7c86abb0931753af1845770ed1d9e46fca1c830f6665e7cec91e97c5514"></a>

## Direct properties — full_mesh / 2983206f40e2 / 3

- [control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-d28ebd06c0a27122b0391c2b25356519b74bea57ea009b554d7668cda4c217ec): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-be64825f67d81fda091ff3fbdb6259eb5f87914c65ecf68e274b296faf44a12e): complete subsection reference.

<a id="canonical-099925731513747d0a2a6e58e1d4aedd5a03bf477b48b8809655fe94aa400009"></a>

## Next pages — full_mesh / 2983206f40e2 / 4

- [full_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-d28ebd06c0a27122b0391c2b25356519b74bea57ea009b554d7668cda4c217ec)
- [full_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-be64825f67d81fda091ff3fbdb6259eb5f87914c65ecf68e274b296faf44a12e)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-d28ebd06c0a27122b0391c2b25356519b74bea57ea009b554d7668cda4c217ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e66a1fc3dc24f18e3ac52cccd3c3be2faa34784d5fcffbfae0f98d968ef093c1"></a>

## full_mesh.control_and_data_plane_mesh — full_mesh.control_and_data_plane_mesh / 8b88e898866c / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8d064e8b4e9544faa84d785a97da2c70dcedcd1866290e0e00925efa7fec05a5)
- full_mesh.control_and_data_plane_mesh

<a id="canonical-c235aaa8218e994c4d8936d7464307fda51b945cf44fee605ba0e04aacbb4fd9"></a>

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

<a id="canonical-2a4d24efe8d1e571ecef15581ebaa039ce1b847bc5247dd54a37c14e9dbc005f"></a>

## Direct properties — full_mesh.control_and_data_plane_mesh / 8b88e898866c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-469f48da88b44b6257e3a5f7d90668f66febd88f7c4d80a8bd21a6d1d2ad355f"></a>

## Next pages — full_mesh.control_and_data_plane_mesh / 8b88e898866c / 4

- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8d064e8b4e9544faa84d785a97da2c70dcedcd1866290e0e00925efa7fec05a5)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-be64825f67d81fda091ff3fbdb6259eb5f87914c65ecf68e274b296faf44a12e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85fb6f718a9c26ed798d615097b25f51eafe2095accd191c671caddafcee3aba"></a>

## full_mesh.data_plane_mesh — full_mesh.data_plane_mesh / 87f581a65a3c / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8d064e8b4e9544faa84d785a97da2c70dcedcd1866290e0e00925efa7fec05a5)
- full_mesh.data_plane_mesh

<a id="canonical-8dfd4e34ea381c706a48eddd23d87dd26a1cf28f930a1895cf7e05d0d08ad14b"></a>

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

<a id="canonical-f18a1f9482d79cfbbdf76f552a042b0f0ede29e2a765bb8d65cd6cb98d698395"></a>

## Direct properties — full_mesh.data_plane_mesh / 87f581a65a3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6e681a66e4d9cf0886e1782483f152511cd0413c73060f4e69b01efc99baa082"></a>

## Next pages — full_mesh.data_plane_mesh / 87f581a65a3c / 4

- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-8d064e8b4e9544faa84d785a97da2c70dcedcd1866290e0e00925efa7fec05a5)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-457c1d87b67af5367ecaa4ac5a73da4519e7370c15bff7fc9617b0b8a4e3ba16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb6e0f278c27fdfe8e307fe08ccde9e67baffabd266a593b094590fa76354ce7"></a>

## hub_mesh — hub_mesh / 58caf4ee3379 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- hub_mesh

<a id="canonical-e55a7c5d9e213d375305bb2089b9710bf37a00c384d1418cac4d9e3bfc86f914"></a>

Type: `"object"`. single nested block, Optional.

Hub Full Mesh. Details of Hub Full Mesh Group Type.

Upstream description:

Details of Hub Full Mesh Group Type.

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
  "x-ves-oneof-field-hub_full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
hub_mesh {
  # Configure direct properties listed below.
}
```

<a id="canonical-4fdedad5f662b4835697ba4d623335dd442d8cae097f72514a182827b0ab17c9"></a>

## Direct properties — hub_mesh / 58caf4ee3379 / 3

- [control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-be1781921e9ee57802c014a5c394e6fa6a17b4d027e605e870b1344415ddc639): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-cdcb41cf4e09c47493e927478a6213391c63cadcc7aa0181215a8d4291fa8a80): complete subsection reference.

<a id="canonical-5cc2a3f56491244dc497c6cd55f3797e3659a808b583ba1fff6004c92acb2917"></a>

## Next pages — hub_mesh / 58caf4ee3379 / 4

- [hub_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-be1781921e9ee57802c014a5c394e6fa6a17b4d027e605e870b1344415ddc639)
- [hub_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-cdcb41cf4e09c47493e927478a6213391c63cadcc7aa0181215a8d4291fa8a80)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-be1781921e9ee57802c014a5c394e6fa6a17b4d027e605e870b1344415ddc639"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34a2db0cab7afbf4e890cca2e75ac026b2279451084ee778a88f08ce40d111f4"></a>

## hub_mesh.control_and_data_plane_mesh — hub_mesh.control_and_data_plane_mesh / 5abd65819ed6 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-457c1d87b67af5367ecaa4ac5a73da4519e7370c15bff7fc9617b0b8a4e3ba16)
- hub_mesh.control_and_data_plane_mesh

<a id="canonical-5831cc095e56f0d863fddb9e7562b2cec30edc6b31d3971675ec1d9914980290"></a>

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

<a id="canonical-fb48a69c966251e3531c13cdf31bb2d6665f7a5f1f2fac4318c2d3268476d02e"></a>

## Direct properties — hub_mesh.control_and_data_plane_mesh / 5abd65819ed6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0829e09e552c3249498a96a3d9d5a411904062cc06e955c1ac105bc8b1fee2db"></a>

## Next pages — hub_mesh.control_and_data_plane_mesh / 5abd65819ed6 / 4

- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-457c1d87b67af5367ecaa4ac5a73da4519e7370c15bff7fc9617b0b8a4e3ba16)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-cdcb41cf4e09c47493e927478a6213391c63cadcc7aa0181215a8d4291fa8a80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d553b7786feffdb825a068ce0a7a3d000db3385a0d291fb5e2209a196081e0db"></a>

## hub_mesh.data_plane_mesh — hub_mesh.data_plane_mesh / e507e216d95f / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-457c1d87b67af5367ecaa4ac5a73da4519e7370c15bff7fc9617b0b8a4e3ba16)
- hub_mesh.data_plane_mesh

<a id="canonical-2d1623f8935fab931a77722f5421665bd349158a0caa67c03562e3376a7623b9"></a>

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

<a id="canonical-4a0ccc30e321d368cda3a72c31bb7e0cb6f1c412596a2f75616341a6757e2430"></a>

## Direct properties — hub_mesh.data_plane_mesh / e507e216d95f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-baf21f60d353ba9abf93c8fe26557d1217ea7092ccd3dae8acd15a077a908a7a"></a>

## Next pages — hub_mesh.data_plane_mesh / e507e216d95f / 4

- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-457c1d87b67af5367ecaa4ac5a73da4519e7370c15bff7fc9617b0b8a4e3ba16)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef0da51d127cc13046d4651e1cfcbb46dde292b2898251a3f0924bc5cd47124f"></a>

## spoke_mesh — spoke_mesh / 85eafaffa48d / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- spoke_mesh

<a id="canonical-b93d470f538648d33f1eedd5beff41b21e6d3c4305cd72f0549e85915008d3ca"></a>

Type: `"object"`. single nested block, Optional.

Spoke. Details of Spoke Mesh Group Type.

Upstream description:

Details of Spoke Mesh Group Type.

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
  "x-ves-oneof-field-spoke_hub_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
spoke_mesh {
  # Configure direct properties listed below.
}
```

<a id="canonical-fcf737a8ce7f8143912255508da9fbae89f7bb839c04e4d9450c59f8bb964f02"></a>

## Direct properties — spoke_mesh / 85eafaffa48d / 3

- [control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-b65054a5fb962d61efaab1f4381da8469155215e024f8433e91c2f72bc96fc0d): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-f771edb7fb9f69d096c86319a6cc8192f03977d0b59e49e7cf06dc863b796cf7): complete subsection reference.

- [hub_mesh_group](resources--site_mesh_group--reference--group-001.md#canonical-a76d323f867ae60baa56331d4cbaa9d255ec39cc18771343fd4037530240c333): complete subsection reference.

<a id="canonical-d6653421aa3de19c647d15c5ff814c50b4c890935f3cc027e9aaeb21745f9bb4"></a>

## Next pages — spoke_mesh / 85eafaffa48d / 4

- [spoke_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-b65054a5fb962d61efaab1f4381da8469155215e024f8433e91c2f72bc96fc0d)
- [spoke_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-f771edb7fb9f69d096c86319a6cc8192f03977d0b59e49e7cf06dc863b796cf7)
- [spoke_mesh.hub_mesh_group](resources--site_mesh_group--reference--group-001.md#canonical-a76d323f867ae60baa56331d4cbaa9d255ec39cc18771343fd4037530240c333)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-b65054a5fb962d61efaab1f4381da8469155215e024f8433e91c2f72bc96fc0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cf391f76747ff3a82eb21bb214abdcf303013d61043ec21384a89732cc9fd7b"></a>

## spoke_mesh.control_and_data_plane_mesh — spoke_mesh.control_and_data_plane_mesh / 25acc4113be7 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759)
- spoke_mesh.control_and_data_plane_mesh

<a id="canonical-e2dd58baa02ff24462e7e14d2f74c4ad139f697801d1784fdb60866cab16f477"></a>

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

<a id="canonical-202fef6fd334c84d50f8525a3a544465683de3daa0cbe3385df3dd8372ff335b"></a>

## Direct properties — spoke_mesh.control_and_data_plane_mesh / 25acc4113be7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9244b334a33662b358c03988e17d7509dfa83d512655e44f895b4ce4f32a9c13"></a>

## Next pages — spoke_mesh.control_and_data_plane_mesh / 25acc4113be7 / 4

- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-f771edb7fb9f69d096c86319a6cc8192f03977d0b59e49e7cf06dc863b796cf7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8af69c871e7350b0bd8f4b20fb03959af9bcf3e88fdc898d44b2fcb2e83ea29"></a>

## spoke_mesh.data_plane_mesh — spoke_mesh.data_plane_mesh / e58b075026f2 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759)
- spoke_mesh.data_plane_mesh

<a id="canonical-8e3b0ac41986739a0b4a0e38bf741e960dbb7a077adaaa343ea20f35b4333e07"></a>

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

<a id="canonical-17b45339fc1d7b6d25df0e9d54b4f5c847da21e5bed3f451715c70719cabe5b7"></a>

## Direct properties — spoke_mesh.data_plane_mesh / e58b075026f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f910af533eebc88b582bb3fc61a767c0fb432c0dad6f24be18b5c5238765dac"></a>

## Next pages — spoke_mesh.data_plane_mesh / e58b075026f2 / 4

- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-a76d323f867ae60baa56331d4cbaa9d255ec39cc18771343fd4037530240c333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b0eedc0709eefd7f2348aa59ef107a2bc57335942aebb03037818641c5072d9"></a>

## spoke_mesh.hub_mesh_group — spoke_mesh.hub_mesh_group / 1d05b3d18e77 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759)
- spoke_mesh.hub_mesh_group

<a id="canonical-877e83a2be83ea49241e2316179c26ed847edd7c605d605e893d8e818a946683"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
hub_mesh_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-0ef4fd3d918626a0a1cc061ae6730e7724213c31c0887d91d597092055147d68"></a>

## Direct properties — spoke_mesh.hub_mesh_group / 1d05b3d18e77 / 3

<a id="canonical-33d80a4940a58a42d13acbfac43dc983e214976607fe16a757cbf4f2c021e7e2"></a>

<a id="canonical-f5c0c3072356a94cf2f4317a619b21412731d8838f8ed3ede3a8b7a22fcbf556"></a>

## name property — spoke_mesh.hub_mesh_group / 1d05b3d18e77 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-5dcce9bf2735f4baeabdf16b3cb213b6df53d16429e32df4334fc854caa5f444"></a>

<a id="canonical-c2d2f84ad619eec53361f35e7b39c996688e2da9672771855a88f94eb5847609"></a>

## namespace property — spoke_mesh.hub_mesh_group / 1d05b3d18e77 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-35ca7521e24b3bff8587a3c03218e146367ea962295492ce54bf61a68f6bc328"></a>

<a id="canonical-1ac21c6f660038113f471352329b545032e50e778aebd5ee62157d4e176bc7c3"></a>

## tenant property — spoke_mesh.hub_mesh_group / 1d05b3d18e77 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-f5f6bf2f23efdcc8725c26926fd5af5db0369394f810a8eb9928f5ef68559a8e"></a>

## Next pages — spoke_mesh.hub_mesh_group / 1d05b3d18e77 / 7

- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-88ea125a1d177a9911285c72c2f5546eec1fdf0b0ed995b0ba53873286c88759)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-0542d8e111378fd85fbf83a05a84465e981936714c770f3c62ad507b27c82d77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bccd89d03323d783645ea967e6266ff7b343da8bc816e273169f550ccd4067c7"></a>

## timeouts — timeouts / 8fbc8a9a1683 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- timeouts

<a id="canonical-7edf4a61bef6426ddc46bdc883de1e09ff9d061d2ebfbb5b210458b0c7565144"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-43f6b13cd80380a084c35f811f23e910c2d848c886bd52dc7c185a6a4891edb3"></a>

## Direct properties — timeouts / 8fbc8a9a1683 / 3

<a id="canonical-f23b020ed9c1dc44434cc9680f8ddf854482357c2e53688a80181411b14ce541"></a>

<a id="canonical-d1bea0fa3376884b4e5738699c6e33d26f19cd69eb4daa1d5e3aca219a637cbf"></a>

## create property — timeouts / 8fbc8a9a1683 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-063c4ceba9e1a2022949b3cd4aedf9884b4e31a581fe8f88f8f8273a763ee78f"></a>

<a id="canonical-ac36dc79079dbf56741296c78ae256d501ab2879851676e9acb13dd1d2d9194f"></a>

## delete property — timeouts / 8fbc8a9a1683 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-737b82dea66707553be834c0f7a36f15edb74c3d14632da9d6052bb99e962cf0"></a>

<a id="canonical-9f9605ab151f61e63ff9de4e33a1e1c5336381f265ecd3173f3e99524a37d889"></a>

## read property — timeouts / 8fbc8a9a1683 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-79d911aeac33f92340a8ac2e1edfe260d534e10f3f76197e302b468c20e940ed"></a>

<a id="canonical-6b31bae8fd47ee86a7a951f61a115eaf92148b5a96aa99ff46fca5c6002b2d14"></a>

## update property — timeouts / 8fbc8a9a1683 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8fa138c97614fc312158431efdda56656decb7321813fb452a82a04c26e9eb42"></a>

## Next pages — timeouts / 8fbc8a9a1683 / 8

- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-68decd4772a312af42a6edde36a861c9725bcc37ee1cf23ac547d5e3b6f5eb00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-247d23bcb2ec04333088b59b7774ad8469cd1254a76744cd50fa408696fc46ab"></a>

## virtual_site — virtual_site / 29726e1a1fe6 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- virtual_site

<a id="canonical-6fecdbc856b80e5d01714a2d5e2c05595b8ca633b5999eded04946133b83eecb"></a>

Type: `"object"`. list nested block, Optional.

Set of sites for which this mesh group config is valid. If 'Type' is Spoke, then it gives set of
spoke sites. If 'Type' is Hub, then it gives set of hub sites.

Upstream description:

Set of sites for which this mesh group config is valid. If 'Type' is Spoke, then it gives set of
spoke sites. If 'Type' is Hub, then it gives set of hub sites. If 'Type' is Full Mesh, then it gives
set of sites that are connected in full mesh.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-544b5b8e3b1b1229c339e19f055e1903da9be67772f68d29c03affc9e3aacc59"></a>

## Direct properties — virtual_site / 29726e1a1fe6 / 3

<a id="canonical-fff5b37b6f76f5075fe9771eee6cb4e397f2b2dc8820069e2ae76da31e0cc503"></a>

<a id="canonical-2405661d4c8c3f162265283f2d9166ea03dc454e827d5036593de1c3b798a998"></a>

## kind property — virtual_site / 29726e1a1fe6 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-034b27cb7a3684ecdd51a3b3520a376a5ed33e880da87c0697ae460ab0a09f8f"></a>

<a id="canonical-6cb179c21c85f9af5e5ea0b22efb225f42f90f84b34f49938688dfdce15ed70b"></a>

## name property — virtual_site / 29726e1a1fe6 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-351e675a0a66271bf24b95e44cf402123e5f3b1df12fc82151d819307ab7df60"></a>

<a id="canonical-1ab220e9e2ba52a103cd262b1a33ce2c7d95d06b09746b6c5552249d68dfe44e"></a>

## namespace property — virtual_site / 29726e1a1fe6 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-e9aa50647bbafc6426cfffded3996614eeb66388b674127c43e7838b2db91484"></a>

<a id="canonical-36888c4e2d449ee0e13c58698b5f3ecffcc3fc57d443e6d132dba36ddeb27bd7"></a>

## tenant property — virtual_site / 29726e1a1fe6 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-310f4185401ab62f9ca2fa82b773283e546b7c95f3bf95895a040c2a1d094eb6"></a>

<a id="canonical-e0a7815438b045d54f9f20402174e0b79c824c3dbe151e832426314f5f9cb25f"></a>

## uid property — virtual_site / 29726e1a1fe6 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-8b6849a9c2b6313cefa77fbf20c3757a8c7b8c09e62c4606eee315d450458312"></a>

## Next pages — virtual_site / 29726e1a1fe6 / 9

- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
