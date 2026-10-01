---
page_title: "xcsh_virtual_network reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network reference."
---

# xcsh_virtual_network reference

<a id="canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9abaf06cb94f1917b64e113d0bde14b3ddbff1cf4db4d1901a854fdf2bed64a"></a>

## Property reference — Property reference / 25787300d35f / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- Property reference

<a id="canonical-b1be83a2561624f1acafeb237445efb06e2b09ee8b25334857d4828e607f0902"></a>

## Direct properties — Property reference / 25787300d35f / 3

<a id="canonical-e23548693676bf1228d82be9330931d59e2b8e1321a6c4038c99609bffda68d6"></a>

<a id="canonical-7efc723decf5b79c466929fc03b3724e312d4ab6cc4dd7fb649ed57cd2c77e0e"></a>

## annotations property — Property reference / 25787300d35f / 4

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

<a id="canonical-93e93d7ee08eeaced86169806bfce54592a6106a75aef73575c816c6d274ee82"></a>

<a id="canonical-af63f82b1da6f775f2deb7658cb7751c96851c0ff64f2a7e8636c1aec96712a4"></a>

## description property — Property reference / 25787300d35f / 5

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

<a id="canonical-a5ffa04d01e1aacd11a42382b1e6258c4f0ea68471128908e5439237b84f5d17"></a>

<a id="canonical-63a8d8fdfde656d3726d8442e93b94cb14c57661f19dfc5f0a8f553b2a96b9f4"></a>

## disable property — Property reference / 25787300d35f / 6

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

- [global_network](resources--virtual_network--reference--group-001.md#canonical-8a430583688976415d0c9c9dddd161a709a151ff33fd58d8cd75da8c86de1558): complete subsection reference.

<a id="canonical-88b18e21093f9a5bb2a61b4d6754c11f298a15e9c1b982b91c57cbb905a8836e"></a>

<a id="canonical-4fce205d4ba2d55b1aadf898db62e1cb3f13e0fd0a15ff5492af2bf6711417ba"></a>

## id property — Property reference / 25787300d35f / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1bad323426049b78712fe7b5c4613d4ddd684b1aa201b4fbeacce6bfbf4181a2"></a>

<a id="canonical-3cf9712d31a61114fc9bb6f9447ffad601cd89548d97dc57ebfa1e6138f4882a"></a>

## labels property — Property reference / 25787300d35f / 8

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

<a id="canonical-6f4d4f6c080e4f41b01f11f3f865167223e2d4c4afa50d95310672305dd8dd59"></a>

<a id="canonical-9a36eb014e8cfd9dcccd3b21d5688e56eddcf9a036b33d65d1cbb47b2e673627"></a>

## name property — Property reference / 25787300d35f / 9

Type: `"string"`. Required.

Name of the Virtual Network. Must be unique within the namespace.

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

<a id="canonical-f65cc2b7cbb131e2257f0a11291c365e780937f3aadcf093cf6e782d01ada4a1"></a>

<a id="canonical-7bf902ca5a0a59b47f109543f661392129e649e5a35a0c74f03c5d1ac22fa40e"></a>

## namespace property — Property reference / 25787300d35f / 10

Type: `"string"`. Optional, Computed.

Namespace for the Virtual Network. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

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

- [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-56fa2f88f867dd3f0c5a84bbca2a536495909bf050c5f92c93a6d87e26dc7851): complete subsection reference.

- [site_local_network](resources--virtual_network--reference--group-001.md#canonical-a9a63da282e92c824b5587b1fd0fd1e94d78085ac5b369ccc7c8ae0ad6e654b5): complete subsection reference.

- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5): complete subsection reference.

- [timeouts](resources--virtual_network--reference--group-001.md#canonical-33cd972095a87e93a9a12c9064ff94de12937675103eebb6485a244ca2e1623e): complete subsection reference.

<a id="canonical-13f07fc373e6e15393cc3cc2e7e80c753b852dcabef0c5b8f70c3534d1ee8975"></a>

## All schema paths — Property reference / 25787300d35f / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--virtual_network--reference--group-001.md#canonical-e23548693676bf1228d82be9330931d59e2b8e1321a6c4038c99609bffda68d6) |
| `description` | [description](resources--virtual_network--reference--group-001.md#canonical-93e93d7ee08eeaced86169806bfce54592a6106a75aef73575c816c6d274ee82) |
| `disable` | [disable](resources--virtual_network--reference--group-001.md#canonical-a5ffa04d01e1aacd11a42382b1e6258c4f0ea68471128908e5439237b84f5d17) |
| `global_network` | [global_network](resources--virtual_network--reference--group-001.md#canonical-2cf7c0e0eea8d30a71d60cabb1b0e096be0cc05356f7c464ad6ac78c805ce93a) |
| `id` | [id](resources--virtual_network--reference--group-001.md#canonical-88b18e21093f9a5bb2a61b4d6754c11f298a15e9c1b982b91c57cbb905a8836e) |
| `labels` | [labels](resources--virtual_network--reference--group-001.md#canonical-1bad323426049b78712fe7b5c4613d4ddd684b1aa201b4fbeacce6bfbf4181a2) |
| `name` | [name](resources--virtual_network--reference--group-001.md#canonical-6f4d4f6c080e4f41b01f11f3f865167223e2d4c4afa50d95310672305dd8dd59) |
| `namespace` | [namespace](resources--virtual_network--reference--group-001.md#canonical-f65cc2b7cbb131e2257f0a11291c365e780937f3aadcf093cf6e782d01ada4a1) |
| `site_local_inside_network` | [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-7919454faf1075153cc14df7ee707ed4b6956ce2c1eb4503998427c2a101b3ea) |
| `site_local_network` | [site_local_network](resources--virtual_network--reference--group-001.md#canonical-a474b32fa78a1746bb052cef8d90de2bc3673acc519f0955dc0445bc6b7a84a7) |
| `static_routes` | [static_routes](resources--virtual_network--reference--group-001.md#canonical-d1df9615befeb3fcbc3e0c18f4e251feaabfaf90bf69db1e6b6a550213d3eaaf) |
| `static_routes.attrs` | [static_routes.attrs](resources--virtual_network--reference--group-001.md#canonical-23ef178f219545b941028b2c16d4defee734b08f0e9085904e5db50a08a05d1b) |
| `static_routes.default_gateway` | [static_routes.default_gateway](resources--virtual_network--reference--group-001.md#canonical-dc2f75ace09b84fe43fb4689b33431316ee831ec01c0e60a32c0d7a2ed7012e6) |
| `static_routes.ip_address` | [static_routes.ip_address](resources--virtual_network--reference--group-001.md#canonical-0e51d8e5b6a816eac3901edabeb05bc46f9ca82d4abbf11ea0d4f859e0c94ac5) |
| `static_routes.ip_prefixes` | [static_routes.ip_prefixes](resources--virtual_network--reference--group-001.md#canonical-745d041a2294974754ba4e9c31eefd2d2d560553d6c7995536c35777d5bd926e) |
| `static_routes.node_interface` | [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-5e286e6c79e4b7df2902fa0f4b5700863a9bce665bb557fb0a1a763fbd190c4b) |
| `static_routes.node_interface.list` | [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-b32ed2cef97e61fd98b21989d2dc0cd20f7f441936a122c444e9880673956ea8) |
| `static_routes.node_interface.list.interface` | [static_routes.node_interface.list.interface](resources--virtual_network--reference--group-001.md#canonical-c52d806ce80dc7c6a61f05d3e6b4f01aadfd3dad050ecf9841ec170f54ba6969) |
| `static_routes.node_interface.list.interface.kind` | [static_routes.node_interface.list.interface.kind](resources--virtual_network--reference--group-001.md#canonical-29379fd600176ad56624e2c66becba59f9db3b5b700a1ee175787509b12c049a) |
| `static_routes.node_interface.list.interface.name` | [static_routes.node_interface.list.interface.name](resources--virtual_network--reference--group-001.md#canonical-81c74c8be041df3302b4f11617e67680a7a1fda5f176ad524900a5ddfb190458) |
| `static_routes.node_interface.list.interface.namespace` | [static_routes.node_interface.list.interface.namespace](resources--virtual_network--reference--group-001.md#canonical-f0ea0b257f7b6a738cf503c24c6641ff02da20e1b8f37893e0036b6be2e0add9) |
| `static_routes.node_interface.list.interface.tenant` | [static_routes.node_interface.list.interface.tenant](resources--virtual_network--reference--group-001.md#canonical-97ed1f6ac69a697a229a930ecdc22aeb0b5c5a7b477905833d504f5528cd2128) |
| `static_routes.node_interface.list.interface.uid` | [static_routes.node_interface.list.interface.uid](resources--virtual_network--reference--group-001.md#canonical-e112e53cc64e295072531e86592c03cd79c2c652f0d397777cb27e72c846c551) |
| `static_routes.node_interface.list.node` | [static_routes.node_interface.list.node](resources--virtual_network--reference--group-001.md#canonical-c4c4c1a4ed81180bdfe4ebdcacd4ff368bb1feaf5e48509735eaf6427f832a1a) |
| `timeouts` | [timeouts](resources--virtual_network--reference--group-001.md#canonical-643ac6dd4e13c3a2db8e0233ece8b42f0bb6947fd999568b2ad35d2040c1d3ba) |
| `timeouts.create` | [timeouts.create](resources--virtual_network--reference--group-001.md#canonical-61bc9c778c2be219f7ba57896f8528aaa107240185abbed1df712eb16f7641d6) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_network--reference--group-001.md#canonical-9afc41a47315591c2fa25794229651c21868848bfa744ab06ce221703fa7816f) |
| `timeouts.read` | [timeouts.read](resources--virtual_network--reference--group-001.md#canonical-2059fef375acaf56798e2ca742b31772d2548ce1a3cda9fb98f11c8f38d418a3) |
| `timeouts.update` | [timeouts.update](resources--virtual_network--reference--group-001.md#canonical-7c31fd2b958f32a11a09f02f9decf035258f7553f670dc3ac7d5671965c88ab8) |

<a id="canonical-4751f81aab9d3486a6d8a3e99d063a1e7b2b25964d65294df6402aae61b76586"></a>

## Next pages — Property reference / 25787300d35f / 12

- [global_network](resources--virtual_network--reference--group-001.md#canonical-8a430583688976415d0c9c9dddd161a709a151ff33fd58d8cd75da8c86de1558)
- [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-56fa2f88f867dd3f0c5a84bbca2a536495909bf050c5f92c93a6d87e26dc7851)
- [site_local_network](resources--virtual_network--reference--group-001.md#canonical-a9a63da282e92c824b5587b1fd0fd1e94d78085ac5b369ccc7c8ae0ad6e654b5)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5)
- [timeouts](resources--virtual_network--reference--group-001.md#canonical-33cd972095a87e93a9a12c9064ff94de12937675103eebb6485a244ca2e1623e)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-8a430583688976415d0c9c9dddd161a709a151ff33fd58d8cd75da8c86de1558"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a909df64cb57e94b33824ff0ddad8bd5d270c906ed2b7229e069aae4f49cb956"></a>

## global_network — global_network / deb658358363 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- global_network

<a id="canonical-2cf7c0e0eea8d30a71d60cabb1b0e096be0cc05356f7c464ad6ac78c805ce93a"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

Upstream description:

Select the global virtual-network scope for connectivity across participating sites.

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

- [global_network](resources--virtual_network--reference--group-001.md#canonical-2cf7c0e0eea8d30a71d60cabb1b0e096be0cc05356f7c464ad6ac78c805ce93a)
- [site_local_inside_network](resources--virtual_network--reference--group-001.md#canonical-7919454faf1075153cc14df7ee707ed4b6956ce2c1eb4503998427c2a101b3ea)
- [site_local_network](resources--virtual_network--reference--group-001.md#canonical-a474b32fa78a1746bb052cef8d90de2bc3673acc519f0955dc0445bc6b7a84a7)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_network = {}
```

<a id="canonical-cbdee37b2d2f7e05aae7721ccbb79190ba5d96c75b5674837897d10771b2474e"></a>

## Direct properties — global_network / deb658358363 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df09a70c2963bf56d99acf8065d4a7550abf656182a4067930b69cb7d7bb91c5"></a>

## Next pages — global_network / deb658358363 / 4

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-56fa2f88f867dd3f0c5a84bbca2a536495909bf050c5f92c93a6d87e26dc7851"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a14aed09cfeb1e0763c727a226f5e0939625c83811f9410e8407e85aebf04bd"></a>

## site_local_inside_network — site_local_inside_network / 65ab3b1911f2 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- site_local_inside_network

<a id="canonical-7919454faf1075153cc14df7ee707ed4b6956ce2c1eb4503998427c2a101b3ea"></a>

Type: `["object", {}]`. Optional.

Select the site-local inside network for site-internal connectivity.

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
site_local_inside_network = {}
```

<a id="canonical-5fa6583beff57230557770e03a8cae8ecb24506c156beb01423df2231763d3c2"></a>

## Direct properties — site_local_inside_network / 65ab3b1911f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc06fd0936472a8cf0266788c8a3673a8b9326dd571de77b15c829e4fe1c0c69"></a>

## Next pages — site_local_inside_network / 65ab3b1911f2 / 4

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-a9a63da282e92c824b5587b1fd0fd1e94d78085ac5b369ccc7c8ae0ad6e654b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed743f4f95d5935c426a4403c3a09d038a0a2b0dbb24382188debf8e9cccf7a5"></a>

## site_local_network — site_local_network / f9f8c0ba86c2 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- site_local_network

<a id="canonical-a474b32fa78a1746bb052cef8d90de2bc3673acc519f0955dc0445bc6b7a84a7"></a>

Type: `["object", {}]`. Optional.

Select a site-local virtual network when connectivity must remain within one site.

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
site_local_network = {}
```

<a id="canonical-11b7d839610d13464ce5e3ca27d1a55c198ba7a2aa3fadabc5a85bed2527e1c0"></a>

## Direct properties — site_local_network / f9f8c0ba86c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60135907000abcbe44ade853add5ce7b34acf70c5929dcda0fdf4233d498acff"></a>

## Next pages — site_local_network / f9f8c0ba86c2 / 4

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-253b70231aa6e3ff6b8768c413d6a43970d882d3a5c30fc38d8838214eabd40f"></a>

## static_routes — static_routes / d8ef0e9d04be / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- static_routes

<a id="canonical-d1df9615befeb3fcbc3e0c18f4e251feaabfaf90bf69db1e6b6a550213d3eaaf"></a>

Type: `"object"`. list nested block, Optional.

List of static routes on the virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 165,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 165,
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
    "ves.io.schema.rules.repeated.max_items": "165",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "165",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a6d95d732cb8e398a97757b0f15991b9e74ade6208c675c360f67d2e124a028"></a>

## Direct properties — static_routes / d8ef0e9d04be / 3

<a id="canonical-23ef178f219545b941028b2c16d4defee734b08f0e9085904e5db50a08a05d1b"></a>

<a id="canonical-6fb1dbecdffca2ec5624a0fb47cca55d18d7798c8d0e0ba70afafee00b895e29"></a>

## attrs property — static_routes / d8ef0e9d04be / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--virtual_network--reference--group-001.md#canonical-e98e91201b6e77de3c4896d7813f80c1a7457f17800cf09dc5f52035873ee2ac): complete subsection reference.

<a id="canonical-0e51d8e5b6a816eac3901edabeb05bc46f9ca82d4abbf11ea0d4f859e0c94ac5"></a>

<a id="canonical-af757fad21de4c49b67f0c834349837360b9d2bd1902dc4815d46fc3259ffafb"></a>

## ip_address property — static_routes / d8ef0e9d04be / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-745d041a2294974754ba4e9c31eefd2d2d560553d6c7995536c35777d5bd926e"></a>

<a id="canonical-2365ee638cca7dfe6d513a47a0b34e6f7b293e173e06a9480c53ceb4c85270b8"></a>

## ip_prefixes property — static_routes / d8ef0e9d04be / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--virtual_network--reference--group-001.md#canonical-f255c6201376cbd351b22f92293ce85b4a2db42b1156e4f2ae6ceb5190a74e08): complete subsection reference.

<a id="canonical-c1d875b34669430feead5f01cdb6ddeb40e988b3e7d0f6d5579d265301ded800"></a>

## Next pages — static_routes / d8ef0e9d04be / 7

- [static_routes.default_gateway](resources--virtual_network--reference--group-001.md#canonical-e98e91201b6e77de3c4896d7813f80c1a7457f17800cf09dc5f52035873ee2ac)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-f255c6201376cbd351b22f92293ce85b4a2db42b1156e4f2ae6ceb5190a74e08)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-e98e91201b6e77de3c4896d7813f80c1a7457f17800cf09dc5f52035873ee2ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8e00d5fa9d3338ad741ab261578efeec2420fc982b9d0385745cf9240e2dfe2"></a>

## static_routes.default_gateway — static_routes.default_gateway / 0524535b8283 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5)
- static_routes.default_gateway

<a id="canonical-dc2f75ace09b84fe43fb4689b33431316ee831ec01c0e60a32c0d7a2ed7012e6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-d547729fdd6bfbaed57dc3a66d49007c0ff37219e6463a12461842711a230531"></a>

## Direct properties — static_routes.default_gateway / 0524535b8283 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4452b0e1adc1b82f0b9ea32df4ba91fb88d61cad33ad20c78a42b0b24c909749"></a>

## Next pages — static_routes.default_gateway / 0524535b8283 / 4

- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-f255c6201376cbd351b22f92293ce85b4a2db42b1156e4f2ae6ceb5190a74e08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5395b3cef27759255e74bc20e5549598b839d2087a13bc1d5f95812cdda2d4c8"></a>

## static_routes.node_interface — static_routes.node_interface / 2e170ea66cfe / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5)
- static_routes.node_interface

<a id="canonical-5e286e6c79e4b7df2902fa0f4b5700863a9bce665bb557fb0a1a763fbd190c4b"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-9b96be35203b1403afca0c05b979bc00b42f65c6b7efd33140f7a53a829aa2d8"></a>

## Direct properties — static_routes.node_interface / 2e170ea66cfe / 3

- [list](resources--virtual_network--reference--group-001.md#canonical-d7314f295f1550a322a7cfa08a610b6018859f0d9157097b6fd24ed519496667): complete subsection reference.

<a id="canonical-c49d60335fb618979dae2b77b85fa001071e368ec9b6604985089ba8c4fbf886"></a>

## Next pages — static_routes.node_interface / 2e170ea66cfe / 4

- [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-d7314f295f1550a322a7cfa08a610b6018859f0d9157097b6fd24ed519496667)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-d7314f295f1550a322a7cfa08a610b6018859f0d9157097b6fd24ed519496667"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f89a42a38a696684f8b82cd1a79c686c58bb32da030a2a8bbb656c2df8a73f21"></a>

## static_routes.node_interface.list — static_routes.node_interface.list / d4f6b233ab13 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-f255c6201376cbd351b22f92293ce85b4a2db42b1156e4f2ae6ceb5190a74e08)
- static_routes.node_interface.list

<a id="canonical-b32ed2cef97e61fd98b21989d2dc0cd20f7f441936a122c444e9880673956ea8"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-d2d130413567db47ecc1c2f570f56222c09954e0c6f30180c81409e9265bd7b6"></a>

## Direct properties — static_routes.node_interface.list / d4f6b233ab13 / 3

- [interface](resources--virtual_network--reference--group-001.md#canonical-6cbde57b2b5ee45934ad2364af3cfaaeeb87cb7a0a25aec9d0d1ab7e01d17607): complete subsection reference.

<a id="canonical-c4c4c1a4ed81180bdfe4ebdcacd4ff368bb1feaf5e48509735eaf6427f832a1a"></a>

<a id="canonical-4facbf1b0c2ebc0b5e9679e0754d8d2d6f066518ff65caa87b8c6c46fe2b9378"></a>

## node property — static_routes.node_interface.list / d4f6b233ab13 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-40403f730af9697e96caa67495583fb45503b9e64405b79556d5f2cde273f173"></a>

## Next pages — static_routes.node_interface.list / d4f6b233ab13 / 5

- [static_routes.node_interface.list.interface](resources--virtual_network--reference--group-001.md#canonical-6cbde57b2b5ee45934ad2364af3cfaaeeb87cb7a0a25aec9d0d1ab7e01d17607)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-f255c6201376cbd351b22f92293ce85b4a2db42b1156e4f2ae6ceb5190a74e08)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-6cbde57b2b5ee45934ad2364af3cfaaeeb87cb7a0a25aec9d0d1ab7e01d17607"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9978bbe2149570f834f8fe937d7aae56eb6c1a048f57a3fa9b02e93d24be032"></a>

## static_routes.node_interface.list.interface — static_routes.node_interface.list.interface / 100ba03b0d77 / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [static_routes](resources--virtual_network--reference--group-001.md#canonical-a72f9113af181478efc09762952fc80e3a171e63acc578fb1cf39f6f8484a0c5)
- [static_routes.node_interface](resources--virtual_network--reference--group-001.md#canonical-f255c6201376cbd351b22f92293ce85b4a2db42b1156e4f2ae6ceb5190a74e08)
- [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-d7314f295f1550a322a7cfa08a610b6018859f0d9157097b6fd24ed519496667)
- static_routes.node_interface.list.interface

<a id="canonical-c52d806ce80dc7c6a61f05d3e6b4f01aadfd3dad050ecf9841ec170f54ba6969"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-09f9775401d35248dd2e7939008bcdd6c6d16a5ce19eacf192e264323b751b19"></a>

## Direct properties — static_routes.node_interface.list.interface / 100ba03b0d77 / 3

<a id="canonical-29379fd600176ad56624e2c66becba59f9db3b5b700a1ee175787509b12c049a"></a>

<a id="canonical-a82f1cf3b830c49211ba904bbac7e8f5ef6123207dc0ea0373fd29347d524531"></a>

## kind property — static_routes.node_interface.list.interface / 100ba03b0d77 / 4

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

<a id="canonical-81c74c8be041df3302b4f11617e67680a7a1fda5f176ad524900a5ddfb190458"></a>

<a id="canonical-a9adb819deb3fe1e2f0116e77315f06f03b2ff135c91c8f1a48bdcf909b65e41"></a>

## name property — static_routes.node_interface.list.interface / 100ba03b0d77 / 5

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

<a id="canonical-f0ea0b257f7b6a738cf503c24c6641ff02da20e1b8f37893e0036b6be2e0add9"></a>

<a id="canonical-11ea4af18c922e2fd1e5409b902a4049d2dc6987a36e927cda2fb9ea6afc80e1"></a>

## namespace property — static_routes.node_interface.list.interface / 100ba03b0d77 / 6

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

<a id="canonical-97ed1f6ac69a697a229a930ecdc22aeb0b5c5a7b477905833d504f5528cd2128"></a>

<a id="canonical-fec0ee13994580c606a5c3a3371b422cac46de4b8aaabc9d771103e59c5c609c"></a>

## tenant property — static_routes.node_interface.list.interface / 100ba03b0d77 / 7

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

<a id="canonical-e112e53cc64e295072531e86592c03cd79c2c652f0d397777cb27e72c846c551"></a>

<a id="canonical-38ac75c459a6d40a7936e07d25cc326d75a6e8ea0c4dab1cbc3a060504b421d5"></a>

## uid property — static_routes.node_interface.list.interface / 100ba03b0d77 / 8

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

<a id="canonical-5cd5bf79b0029f3bd340783e52ab9a449d5718cdd272a879d493c5e8eaa06a5b"></a>

## Next pages — static_routes.node_interface.list.interface / 100ba03b0d77 / 9

- [static_routes.node_interface.list](resources--virtual_network--reference--group-001.md#canonical-d7314f295f1550a322a7cfa08a610b6018859f0d9157097b6fd24ed519496667)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)

<a id="canonical-33cd972095a87e93a9a12c9064ff94de12937675103eebb6485a244ca2e1623e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-658630d43e75973d8fad77dbcf3671aaf1777886a6d9c983eedbc0bbbfc5c8c1"></a>

## timeouts — timeouts / fc66c4643e1c / 2

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- timeouts

<a id="canonical-643ac6dd4e13c3a2db8e0233ece8b42f0bb6947fd999568b2ad35d2040c1d3ba"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-cbdc5898c1521c7aba2f70be2f80e5376f93562b2acc5dbea812652b1730770a"></a>

## Direct properties — timeouts / fc66c4643e1c / 3

<a id="canonical-61bc9c778c2be219f7ba57896f8528aaa107240185abbed1df712eb16f7641d6"></a>

<a id="canonical-3f90cd83c772c57765414390308b9a9135467626c149b84a958a97142c0baed6"></a>

## create property — timeouts / fc66c4643e1c / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9afc41a47315591c2fa25794229651c21868848bfa744ab06ce221703fa7816f"></a>

<a id="canonical-d1b62621d0f749600bcda1a0a1b8cea4509a34a4abd0313494889716d43b8bb3"></a>

## delete property — timeouts / fc66c4643e1c / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2059fef375acaf56798e2ca742b31772d2548ce1a3cda9fb98f11c8f38d418a3"></a>

<a id="canonical-755e836ca44d20f222385d3102ce6489b0868aad0a937e64a7fef7b7adeb92e9"></a>

## read property — timeouts / fc66c4643e1c / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-7c31fd2b958f32a11a09f02f9decf035258f7553f670dc3ac7d5671965c88ab8"></a>

<a id="canonical-cfdad1842b79c166fb0f12e909eee4fb6ab2f5a89a5f807ac400e246997ed893"></a>

## update property — timeouts / fc66c4643e1c / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a161cfd4aae3e784def474276d7d51c290909c33be1bb7971c19d92031332860"></a>

## Next pages — timeouts / fc66c4643e1c / 8

- [Property reference](resources--virtual_network--reference--group-001.md#canonical-3ea74abf82db9443d83a671d5f5b828b070ed4def044e50a05a69761fc5a0ccf)
- [xcsh_virtual_network](../resources/virtual_network.md#canonical-21f8cd52cc88675de24bb43345d87da0d8c4a374c7bc1f8c33fd0f7b6ebd5181)
