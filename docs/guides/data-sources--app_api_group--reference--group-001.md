---
page_title: "xcsh_app_api_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group reference."
---

# xcsh_app_api_group reference

<a id="canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e61c85e7fa8548e0de39b173a8260eca750623b32a43af44b62d9c7dff3eadeb"></a>

## Property reference — Property reference / fed42db61da7 / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- Property reference

<a id="canonical-62b96cb8b7c17d8855cb148680fc32c353bcded2eeccbba949d1c07407070a7f"></a>

## Direct properties — Property reference / fed42db61da7 / 3

<a id="canonical-6841edb593f2b299f30e8493975b04ad9094a8a204abd53a9611907c5f082b39"></a>

<a id="canonical-34c8d99185bda3946bbe212a22cff8069c08ef4d8f483f4a3e3c9676dc7f8c5e"></a>

## annotations property — Property reference / fed42db61da7 / 4

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

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-671567299690bf951a05a315c05c2aa92adffd6c16826559e3a0b3d2c8e29ab3): complete subsection reference.

- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-5a9d2b165628351ce8f57aad5203d5f03531155084d3cfcfd3645079f841a415): complete subsection reference.

<a id="canonical-997dbb24e30bf5347539d3d726c72cdf85c02beea360dd48a0e0021b1f386d1f"></a>

<a id="canonical-64932a991243d261d1b9a8446a63cc755869904d124c8827b4d8950ff193c542"></a>

## description property — Property reference / fed42db61da7 / 5

Type: `"string"`. Computed.

Description of the AppAPIGroup.

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

- [elements](data-sources--app_api_group--reference--group-001.md#canonical-625319b4c95fe7c98d85df781951fe0a367709e196ec8057ae8e1ad5cf837b75): complete subsection reference.

- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-46f29d18679c9d0498e0007aefdf0fb3a9a4144a4773b9927bb99f4aa4ae3eba): complete subsection reference.

<a id="canonical-384f82e4f9edba4152af4c057a2806ef3ce33b981b4557cb30778db7f27630d9"></a>

<a id="canonical-4fc63cee8a30da022fbf195d778523e1e6c5278c6102af27823c0ca951efbbaf"></a>

## id property — Property reference / fed42db61da7 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-002eaa67cce5033d024e949e87e2bcf75bc96a552cc2f2e5a22a47cf0d823ca2"></a>

<a id="canonical-9e6b0e48526ec35a3a3ad545dce602e24530353ed0c09922b22aa35cf7c86f29"></a>

## labels property — Property reference / fed42db61da7 / 7

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

<a id="canonical-80ed9d5dd16136ca90da1ef0d78c713cb5a0f13e9578e44b0f959feea7423308"></a>

<a id="canonical-6b8201246c55f6f436fa525a8d71432f825644ce1c9580cc8fd912ec8fa7cb1d"></a>

## name property — Property reference / fed42db61da7 / 8

Type: `"string"`. Required.

Name of the AppAPIGroup.

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

<a id="canonical-913537dd90f0496a413aa991bad7134d86c75af9835a7cbbbe228ff226908c65"></a>

<a id="canonical-4acfdc5c62156cf2a171b904e27554b303457c386e14f4cf610409aef962dac9"></a>

## namespace property — Property reference / fed42db61da7 / 9

Type: `"string"`. Required.

Namespace where the AppAPIGroup exists.

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

<a id="canonical-d1191841137f7898493e9dcea0fd6d01b91a6ae8865fd306656ee11286b6778c"></a>

## All schema paths — Property reference / fed42db61da7 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--app_api_group--reference--group-001.md#canonical-6841edb593f2b299f30e8493975b04ad9094a8a204abd53a9611907c5f082b39) |
| `bigip_virtual_server` | [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-6b2191c2ed096c99aa0434e9046c8e47fecb9b2bf536b4a0e00811f97504f80e) |
| `bigip_virtual_server.bigip_virtual_server` | [bigip_virtual_server.bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-b86b2aba327600eb8471a116b931292d6cedddbd1b408a999d6ad98a2feebbef) |
| `bigip_virtual_server.bigip_virtual_server.name` | [bigip_virtual_server.bigip_virtual_server.name](data-sources--app_api_group--reference--group-001.md#canonical-68fb7ca82906ccb9b05a2b8483499a51c3275c168ef463de27a5b807d667c078) |
| `bigip_virtual_server.bigip_virtual_server.namespace` | [bigip_virtual_server.bigip_virtual_server.namespace](data-sources--app_api_group--reference--group-001.md#canonical-4d2350e9bbe7aa66bef6221df4262579d341f3ffe8962abc248b98a084e99493) |
| `bigip_virtual_server.bigip_virtual_server.tenant` | [bigip_virtual_server.bigip_virtual_server.tenant](data-sources--app_api_group--reference--group-001.md#canonical-c17d86f3762a4c016657b085a0f6460c881c2f581feaf7329975f270e9e7544d) |
| `cdn_loadbalancer` | [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-0313a1950dabaa513a82fb3886c33b9923d0949ec07071e9ad885610a411023b) |
| `cdn_loadbalancer.cdn_loadbalancer` | [cdn_loadbalancer.cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-3a12a96e1fbf310c4323575ff1e9e35d7532ca32e7b12dbcb10e502ec11c8aa2) |
| `cdn_loadbalancer.cdn_loadbalancer.name` | [cdn_loadbalancer.cdn_loadbalancer.name](data-sources--app_api_group--reference--group-001.md#canonical-007f47eb54c9703eeeceade2ca1d5963e0cfacb4fe55c8943a836c0ab75c69a0) |
| `cdn_loadbalancer.cdn_loadbalancer.namespace` | [cdn_loadbalancer.cdn_loadbalancer.namespace](data-sources--app_api_group--reference--group-001.md#canonical-0893c70e188078d1508e293303ed48274ad8df57f9eb865244fbbb16b2d9b9f4) |
| `cdn_loadbalancer.cdn_loadbalancer.tenant` | [cdn_loadbalancer.cdn_loadbalancer.tenant](data-sources--app_api_group--reference--group-001.md#canonical-9e87a9e3b2f3b10b6a170379130eb4b249be27d8476ef70f2c1832bb073ea74f) |
| `description` | [description](data-sources--app_api_group--reference--group-001.md#canonical-997dbb24e30bf5347539d3d726c72cdf85c02beea360dd48a0e0021b1f386d1f) |
| `elements` | [elements](data-sources--app_api_group--reference--group-001.md#canonical-abba819cd315e4c9087b7a483ee811636c4d895c84a348db27862b86496a22a2) |
| `elements.methods` | [elements.methods](data-sources--app_api_group--reference--group-001.md#canonical-75e1aecfa9f734b046b94a0506e91a877da56bbe46f65ae9f5a7753907843807) |
| `elements.path_regex` | [elements.path_regex](data-sources--app_api_group--reference--group-001.md#canonical-1cb009f040d1d9eba6b79c4a7131887966e677da1a1ca33ef4303616c5892ef6) |
| `http_loadbalancer` | [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-14ba826a8b6d2cb21bc6b36e71d4a2bca704af46fffbad296078d349c6900644) |
| `http_loadbalancer.http_loadbalancer` | [http_loadbalancer.http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-7454d72d5a4a6db75c4d2b13b87c8a917414584edd9fcfba9bdb41aa890ce825) |
| `http_loadbalancer.http_loadbalancer.name` | [http_loadbalancer.http_loadbalancer.name](data-sources--app_api_group--reference--group-001.md#canonical-0f952dcf4e5449d876b4b9b448a8d2bec65a4bfbfed5f6b7549135eb21e2318b) |
| `http_loadbalancer.http_loadbalancer.namespace` | [http_loadbalancer.http_loadbalancer.namespace](data-sources--app_api_group--reference--group-001.md#canonical-e2de98f19a227d1a9a14ac713c1ab606aa9bebfbee4cfae8843160a28eb0a3d5) |
| `http_loadbalancer.http_loadbalancer.tenant` | [http_loadbalancer.http_loadbalancer.tenant](data-sources--app_api_group--reference--group-001.md#canonical-5b876687a164bbea545ddf5a9190e92ba90576d7a8ca7e34ba7f5e5e45ffa98a) |
| `id` | [id](data-sources--app_api_group--reference--group-001.md#canonical-384f82e4f9edba4152af4c057a2806ef3ce33b981b4557cb30778db7f27630d9) |
| `labels` | [labels](data-sources--app_api_group--reference--group-001.md#canonical-002eaa67cce5033d024e949e87e2bcf75bc96a552cc2f2e5a22a47cf0d823ca2) |
| `name` | [name](data-sources--app_api_group--reference--group-001.md#canonical-80ed9d5dd16136ca90da1ef0d78c713cb5a0f13e9578e44b0f959feea7423308) |
| `namespace` | [namespace](data-sources--app_api_group--reference--group-001.md#canonical-913537dd90f0496a413aa991bad7134d86c75af9835a7cbbbe228ff226908c65) |

<a id="canonical-5feb52b97b806177284a68215516253518c3633e9fd61541a04b042ea7a47c1b"></a>

## Next pages — Property reference / fed42db61da7 / 11

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-671567299690bf951a05a315c05c2aa92adffd6c16826559e3a0b3d2c8e29ab3)
- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-5a9d2b165628351ce8f57aad5203d5f03531155084d3cfcfd3645079f841a415)
- [elements](data-sources--app_api_group--reference--group-001.md#canonical-625319b4c95fe7c98d85df781951fe0a367709e196ec8057ae8e1ad5cf837b75)
- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-46f29d18679c9d0498e0007aefdf0fb3a9a4144a4773b9927bb99f4aa4ae3eba)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-671567299690bf951a05a315c05c2aa92adffd6c16826559e3a0b3d2c8e29ab3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ecf831fa566e1bbe748a696b57665ef969cbf8e28f2a0ad7263293e002f78d7"></a>

## bigip_virtual_server — bigip_virtual_server / 9d849025aca7 / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- bigip_virtual_server

<a id="canonical-6b2191c2ed096c99aa0434e9046c8e47fecb9b2bf536b4a0e00811f97504f80e"></a>

Type: `"single"`. Computed.

\[OneOf: bigip\_virtual\_server, cdn\_loadbalancer, http\_loadbalancer\] Set the scope of the API
Group to a specific BIG-IP Virtual Server.

Upstream description:

Set the scope of the API Group to a specific BIG-IP Virtual Server.

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

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-6b2191c2ed096c99aa0434e9046c8e47fecb9b2bf536b4a0e00811f97504f80e)
- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-0313a1950dabaa513a82fb3886c33b9923d0949ec07071e9ad885610a411023b)
- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-14ba826a8b6d2cb21bc6b36e71d4a2bca704af46fffbad296078d349c6900644)

Select alternatives according to the provider validators above.

<a id="canonical-4cdfd1fd1cccad426ed775a86a0749b7e17157201420a2fc86a8f6861dab8e94"></a>

## Direct properties — bigip_virtual_server / 9d849025aca7 / 3

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-8d0e3d6c8a7f70378b171f8416e02b90fd4bb38dd87bdeee3aa320729541b98c): complete subsection reference.

<a id="canonical-8c4481811ccd10c055f47d7ce7f87a04f372391aed0352960b48ef813e577eb0"></a>

## Next pages — bigip_virtual_server / 9d849025aca7 / 4

- [bigip_virtual_server.bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-8d0e3d6c8a7f70378b171f8416e02b90fd4bb38dd87bdeee3aa320729541b98c)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-8d0e3d6c8a7f70378b171f8416e02b90fd4bb38dd87bdeee3aa320729541b98c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd2c4c80c0a6b9f3d8a888080aa78fd144299dcfeb92defe10a1836ec458d739"></a>

## bigip_virtual_server.bigip_virtual_server — bigip_virtual_server.bigip_virtual_server / 8d83fd066c1a / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-671567299690bf951a05a315c05c2aa92adffd6c16826559e3a0b3d2c8e29ab3)
- bigip_virtual_server.bigip_virtual_server

<a id="canonical-b86b2aba327600eb8471a116b931292d6cedddbd1b408a999d6ad98a2feebbef"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-94314e06f067386a9fd46260e8615b20b20d133dd7a2041385fe1240d8377c98"></a>

## Direct properties — bigip_virtual_server.bigip_virtual_server / 8d83fd066c1a / 3

<a id="canonical-68fb7ca82906ccb9b05a2b8483499a51c3275c168ef463de27a5b807d667c078"></a>

<a id="canonical-6418e2cbdbc0e4171b6b34d09b2f77c79606a159a69e0fd1af2284a459b02080"></a>

## name property — bigip_virtual_server.bigip_virtual_server / 8d83fd066c1a / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-4d2350e9bbe7aa66bef6221df4262579d341f3ffe8962abc248b98a084e99493"></a>

<a id="canonical-f50e4eb726c001743d8bc26644895668eb60228e2cb759aee789a19af890d4d1"></a>

## namespace property — bigip_virtual_server.bigip_virtual_server / 8d83fd066c1a / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-c17d86f3762a4c016657b085a0f6460c881c2f581feaf7329975f270e9e7544d"></a>

<a id="canonical-a3a34eeeea2d2f52e5bd8e4998e6df522053f7196281ffc47b6454442d5ed1bc"></a>

## tenant property — bigip_virtual_server.bigip_virtual_server / 8d83fd066c1a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-cdd7351d97d1a39de85ea2360f72d2c7bfd78f523e7e29d2f6f7460f910588b0"></a>

## Next pages — bigip_virtual_server.bigip_virtual_server / 8d83fd066c1a / 7

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-671567299690bf951a05a315c05c2aa92adffd6c16826559e3a0b3d2c8e29ab3)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-5a9d2b165628351ce8f57aad5203d5f03531155084d3cfcfd3645079f841a415"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-937b2f22eaa6e4c785125de7a009ad9e32affe0cf224f8ee5305ca71206c7ca8"></a>

## cdn_loadbalancer — cdn_loadbalancer / c4c92ad212a2 / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- cdn_loadbalancer

<a id="canonical-0313a1950dabaa513a82fb3886c33b9923d0949ec07071e9ad885610a411023b"></a>

Type: `"single"`. Computed.

Set the scope of the API Group to a specific CDN Loadbalancer.

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

<a id="canonical-081345ccc3d9ccd626daca4c3b17f063d3ef42c3b2135dbe0b1b5ae74f5ac67b"></a>

## Direct properties — cdn_loadbalancer / c4c92ad212a2 / 3

- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-c6ad1fa946b0d5e1da5690185564805db0b0adbfe57889bdc2dd3618335c718d): complete subsection reference.

<a id="canonical-09ada8bc7e9e1d7fc6309cb931770a6737f09788752eb0d8e2a0c5b4bc29c167"></a>

## Next pages — cdn_loadbalancer / c4c92ad212a2 / 4

- [cdn_loadbalancer.cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-c6ad1fa946b0d5e1da5690185564805db0b0adbfe57889bdc2dd3618335c718d)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-c6ad1fa946b0d5e1da5690185564805db0b0adbfe57889bdc2dd3618335c718d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40ace6e3bfc2f0215b9cb62b134e40a2801f2b9b2b09cd8accac4dcd17a766d6"></a>

## cdn_loadbalancer.cdn_loadbalancer — cdn_loadbalancer.cdn_loadbalancer / 3278a4d76895 / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-5a9d2b165628351ce8f57aad5203d5f03531155084d3cfcfd3645079f841a415)
- cdn_loadbalancer.cdn_loadbalancer

<a id="canonical-3a12a96e1fbf310c4323575ff1e9e35d7532ca32e7b12dbcb10e502ec11c8aa2"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-c94fbde208dda10ef9bbf4a54d6893832f50cb732eaf9de1e834a2f0dc2802b3"></a>

## Direct properties — cdn_loadbalancer.cdn_loadbalancer / 3278a4d76895 / 3

<a id="canonical-007f47eb54c9703eeeceade2ca1d5963e0cfacb4fe55c8943a836c0ab75c69a0"></a>

<a id="canonical-21851977e24eb6c1744f8f4265d9c5c5522f7b8bb580db635b2a040f51199d67"></a>

## name property — cdn_loadbalancer.cdn_loadbalancer / 3278a4d76895 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0893c70e188078d1508e293303ed48274ad8df57f9eb865244fbbb16b2d9b9f4"></a>

<a id="canonical-874aab36afccf469170ea9f22d6d1611a15a494aa8355f61e4bdbf038e975075"></a>

## namespace property — cdn_loadbalancer.cdn_loadbalancer / 3278a4d76895 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-9e87a9e3b2f3b10b6a170379130eb4b249be27d8476ef70f2c1832bb073ea74f"></a>

<a id="canonical-01241271e9a6beca894e5cfcf91ef9f6a33cda1c3e15b944ff06a5e01205fb9a"></a>

## tenant property — cdn_loadbalancer.cdn_loadbalancer / 3278a4d76895 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-10ec43449c3552db24a846e04d8c3371400d90a5aa0aa0e1d76d1db854933564"></a>

## Next pages — cdn_loadbalancer.cdn_loadbalancer / 3278a4d76895 / 7

- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-5a9d2b165628351ce8f57aad5203d5f03531155084d3cfcfd3645079f841a415)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-625319b4c95fe7c98d85df781951fe0a367709e196ec8057ae8e1ad5cf837b75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e977451785ca667a1e90133ef644ab2ec424df55a852fd573b30d4a8a341e8c6"></a>

## elements — elements / 0f38d3007d04 / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- elements

<a id="canonical-abba819cd315e4c9087b7a483ee811636c4d895c84a348db27862b86496a22a2"></a>

Type: `"list"`. Computed.

List of API group elements with methods and path regex for matching requests.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  }
}
```

<a id="canonical-07858dd5d398d2e0d15d9f9b7297fbcd69edbc122b6bcd30c5ffd1f96c121da9"></a>

## Direct properties — elements / 0f38d3007d04 / 3

<a id="canonical-75e1aecfa9f734b046b94a0506e91a877da56bbe46f65ae9f5a7753907843807"></a>

<a id="canonical-5deae24d07bd5316e07459122c6c2ce0330daa6334cf96f3fcba0e51af91dc8a"></a>

## methods property — elements / 0f38d3007d04 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of method values to
match the input request API method against. The match is considered to succeed if the input request
API method is a member of the list. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`,
\`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of method values to match the input request API method against. The match is considered to
succeed if the input request API method is a member of the list.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1cb009f040d1d9eba6b79c4a7131887966e677da1a1ca33ef4303616c5892ef6"></a>

<a id="canonical-6872da0cd0a6a1286ad8c5c6ef50ecb309d2114c350eb17bfddbc3c2d4f2af5e"></a>

## path_regex property — elements / 0f38d3007d04 / 5

Type: `"string"`. Computed.

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regex.

Upstream description:

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regex.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-022d98cd8a13e90c5b4adf57047b02aad0d3b1880b891f893363b3a8a650f977"></a>

## Next pages — elements / 0f38d3007d04 / 6

- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-46f29d18679c9d0498e0007aefdf0fb3a9a4144a4773b9927bb99f4aa4ae3eba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b570c22ce818d739c9516bc73c989bd3b74a69b6395cf02dbead923bbb8f3ed3"></a>

## http_loadbalancer — http_loadbalancer / 71a68d6bb0ac / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- http_loadbalancer

<a id="canonical-14ba826a8b6d2cb21bc6b36e71d4a2bca704af46fffbad296078d349c6900644"></a>

Type: `"single"`. Computed.

Set the scope of the API Group to a specific HTTP Loadbalancer.

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

<a id="canonical-3accfd20fd7141cd46e4d1f9848bcf303d7da054441c25d9146f5d7df48432b6"></a>

## Direct properties — http_loadbalancer / 71a68d6bb0ac / 3

- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-d5e887b29c6a058aa2fbe0037924d9eabd440566945bab591bb5ca5e55e1af39): complete subsection reference.

<a id="canonical-4e2c3235c7818881c012728a97446b29e16db6b4b50d0531de0fb142c7a9bcbc"></a>

## Next pages — http_loadbalancer / 71a68d6bb0ac / 4

- [http_loadbalancer.http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-d5e887b29c6a058aa2fbe0037924d9eabd440566945bab591bb5ca5e55e1af39)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-d5e887b29c6a058aa2fbe0037924d9eabd440566945bab591bb5ca5e55e1af39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71aaa514095533d99634a3c88fdd47081e66a6800757770f0f627bde00b88546"></a>

## http_loadbalancer.http_loadbalancer — http_loadbalancer.http_loadbalancer / c359667efcef / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-dd0ec6121fb9cd90e1fab0ea768ab0167433fd45e59e6a684e9c62f30672102e)
- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-46f29d18679c9d0498e0007aefdf0fb3a9a4144a4773b9927bb99f4aa4ae3eba)
- http_loadbalancer.http_loadbalancer

<a id="canonical-7454d72d5a4a6db75c4d2b13b87c8a917414584edd9fcfba9bdb41aa890ce825"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-f90d806fe49c24ddfe432471583f5f8166ad67d53c6dbbbfbc0fa3b7ad19c00d"></a>

## Direct properties — http_loadbalancer.http_loadbalancer / c359667efcef / 3

<a id="canonical-0f952dcf4e5449d876b4b9b448a8d2bec65a4bfbfed5f6b7549135eb21e2318b"></a>

<a id="canonical-332a0e59a39a2f5d128faff76547aa0c0201b18414bac36e7aa292470ec02eec"></a>

## name property — http_loadbalancer.http_loadbalancer / c359667efcef / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-e2de98f19a227d1a9a14ac713c1ab606aa9bebfbee4cfae8843160a28eb0a3d5"></a>

<a id="canonical-02b725175c269864ad7cfd107a66a0ae47d7d5fb9ffb9487137cbd2b2dfee5c5"></a>

## namespace property — http_loadbalancer.http_loadbalancer / c359667efcef / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-5b876687a164bbea545ddf5a9190e92ba90576d7a8ca7e34ba7f5e5e45ffa98a"></a>

<a id="canonical-e93c7aba6b8d96c045be248d43c422ace2b352fa266b3b1bd0f90a645b2ddc30"></a>

## tenant property — http_loadbalancer.http_loadbalancer / c359667efcef / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-af98c57d90e708137093fe307569a66dcf16700c96a5254e5ef60781fde07d73"></a>

## Next pages — http_loadbalancer.http_loadbalancer / c359667efcef / 7

- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-46f29d18679c9d0498e0007aefdf0fb3a9a4144a4773b9927bb99f4aa4ae3eba)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
