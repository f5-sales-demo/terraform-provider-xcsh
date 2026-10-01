---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e6fe80bb741517389f08a6d4fc8c8b7f8da57bea5b43545a5798372671f347a"></a>

## Property reference — Property reference / b351774ec556 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- Property reference

<a id="canonical-9ab61596eca48fbd4d537fbc287c80ed3790f7f7092204cf2cc81e231c6cc912"></a>

## Direct properties — Property reference / b351774ec556 / 3

<a id="canonical-abfc99e6667db938e2b8cb7e60ab3a91c327d52a219f1fc2c6544cf9be8c75c9"></a>

<a id="canonical-20bfcfbd5dc1fcb3fc87fbb0df351e8394075523d7f9820da3abfcc2bdb16c3a"></a>

## address property — Property reference / b351774ec556 / 4

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-a74480952c3ed22da8da7d7f1d623279f708017019b1316af8123b9410f3ccad"></a>

<a id="canonical-f30b34a5c5f6a588ec9893a66cb5534d07baff62d7c246eb53ce3e5f6af4e445"></a>

## annotations property — Property reference / b351774ec556 / 5

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

- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d): complete subsection reference.

- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-b46729d24d7c5fedc7541c50c6597214e39effced7311afc2973e351df32216e): complete subsection reference.

- [coordinates](resources--securemesh_site--reference--group-001.md#canonical-6af687786ff749e0698ebec1775e793f9421f4989d96ca808a763d39f55d5db6): complete subsection reference.

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec): complete subsection reference.

- [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-9ce856978b9d1a3834c3b63a67ff7d42db5ed137144b2db86037212335e19331): complete subsection reference.

- [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-6598b3c9c785b28b0be5365f6ae8bc3b1d78ecfffef33c5a8ea18ada89496ed4): complete subsection reference.

<a id="canonical-8b969d40126e45dcdfed605e968a3645963ab803c8e0855177a4807278f3f2c6"></a>

<a id="canonical-b3a323611d07ef8c8690c6dff1fbf40e86287a4e8c132be9947abf6229de6b50"></a>

## description property — Property reference / b351774ec556 / 6

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

<a id="canonical-72ea41d51fe48dcd341143852b4c8d210af090cd6493630a4417bc9a6139773d"></a>

<a id="canonical-4720c7b4e128d64dfe5f7574470c2b890ab02d1a9fcbbc14f8160d969738b06e"></a>

## disable property — Property reference / b351774ec556 / 7

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

<a id="canonical-529d92e98ec3e450cb205a80d12ff3960ed8d9e53a3de28f31891ebf660e21d7"></a>

<a id="canonical-a922a87a8f9e7e89920af7d6b30b93feab747a06cee4da23339c183c342d07de"></a>

## id property — Property reference / b351774ec556 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69): complete subsection reference.

<a id="canonical-e3b790d17ba3356ac782ad59389de40c47daada8cfb1e0610fcbd94d17a361b1"></a>

<a id="canonical-259d29d1cb2539abc2db52c969174af835e1cf85a3ec7d25d7d550e1d466adc2"></a>

## labels property — Property reference / b351774ec556 / 9

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

- [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-64b4390e24c8d3cd7963fefe17327c30c38775d5b9fc0c3a76696addb3c8473b): complete subsection reference.

- [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-31eb5954f742df58c87702640dbcd9722b462484fa62264b963a0128a396fc02): complete subsection reference.

- [master_node_configuration](resources--securemesh_site--reference--group-004.md#canonical-8c2774d9115547c29c565d10884a46a3d040cea98f976355ca3e3dace10ebd9a): complete subsection reference.

<a id="canonical-e70fe104dfb15b5e0ba4617a90e1ddd8a56d0187ffdd65227202ead68fb341bd"></a>

<a id="canonical-d13beeac2e93f68ce234c4b3ba32524cb03728e0e5f1268d3dd3c682e5a79ba2"></a>

## name property — Property reference / b351774ec556 / 10

Type: `"string"`. Required.

Name of the Securemesh Site. Must be unique within the namespace.

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

<a id="canonical-047d58428e59531f5af489dcb2e723200e13fe2bc5b3826da80ecf2b56549e05"></a>

<a id="canonical-d66e8e45954ad8e46bee30a9b11da6ed468bb25edc4f4a4811b43924873d6a02"></a>

## namespace property — Property reference / b351774ec556 / 11

Type: `"string"`. Required.

Namespace where the Securemesh Site is created.

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

- [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-3f591357d9730ab49c161f878602fbcbad1f798979daab263ba21dd5e07fe37c): complete subsection reference.

- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-8876dd170b5fea86ae9ef472fdff18fdf1d5c6c740e6d812f7d8ab5751811350): complete subsection reference.

- [os](resources--securemesh_site--reference--group-004.md#canonical-127eee1dad136acb23d6139da4b0993a34521376b95d09aa218b63fcc060cafa): complete subsection reference.

- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f): complete subsection reference.

- [sw](resources--securemesh_site--reference--group-004.md#canonical-401806d8a416f123ff291c4a5f724bf9513d7fad7f5ce8928e9b0ac494e6f2e7): complete subsection reference.

- [timeouts](resources--securemesh_site--reference--group-004.md#canonical-edc9ffffe4ec4adfdc3b384259fa9b388799b14d92d94a5b7fb79492e7d5346e): complete subsection reference.

<a id="canonical-a1f7959dab3807650db3b45a8128f6488958dc18e2a47dfe5d5bb214d3533d4e"></a>

<a id="canonical-99c5abe15a094f2af9eb00cf344755b252bb7fa0300be2c5bb617bc8f394162a"></a>

## volterra_certified_hw property — Property reference / b351774ec556 / 12

Type: `"string"`. Required.

Name for generic server certified hardware to form this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-e3d3b3dd04a2677e5757751a4242bcaa7ace22229d58fc2ed27e711cfcaac7e0): complete subsection reference.

<a id="canonical-68f1e051f10319f4ee6465087917802319c20bb53b9e68dab85714ac1c7342da"></a>

<a id="canonical-06233a126754524cd9550d526d8f11c3ecfce31d8018cca237d5db78c81028f1"></a>

## worker_nodes property — Property reference / b351774ec556 / 13

Type: `["list", "string"]`. Optional.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3f4e16f54f922acd8f99782195f725da8c6bc294c69263ec46251a5bebdd9a4c"></a>

## All schema paths — Property reference / b351774ec556 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--securemesh_site--reference--group-001.md#canonical-abfc99e6667db938e2b8cb7e60ab3a91c327d52a219f1fc2c6544cf9be8c75c9) |
| `annotations` | [annotations](resources--securemesh_site--reference--group-001.md#canonical-a74480952c3ed22da8da7d7f1d623279f708017019b1316af8123b9410f3ccad) |
| `blocked_services` | [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-7bf7f39c6626a3b36d80b5046bf19787fe5f064b23de1044d9d4e73ed0ddb6f3) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-3c14d732efecd597d23ff647c449b4b29e876f5f6a26d7c2ca92b3907813bd8e) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--securemesh_site--reference--group-001.md#canonical-175332b3a333ba6783bd9d2b0b3011046310b13f8b50d43ecb5899ed792a06d6) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--securemesh_site--reference--group-001.md#canonical-b240b118dcfe47656fc13e5f270ad31853647a045aac4330974508ad98a7654c) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--securemesh_site--reference--group-001.md#canonical-091265a6fd447b4a1b3fd0660e46b0dee1cac4f35ed29919cd9a14cc8ae289da) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--securemesh_site--reference--group-001.md#canonical-3c3dccdb4341cfeea92b3a8a4a165f4fbea8a389679a0913baa39785ef76a964) |
| `bond_device_list` | [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-a065551d6ebc597aa0290d29c4bfe0762ce6d861833f409457ff46b525cd360c) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-c22a2ce54071534b6f7a25647c8babe4e6897fa4c581595de99904f3f1b9ee23) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](resources--securemesh_site--reference--group-001.md#canonical-4713ac5fa7a67c2d086d68d3facd6e5bca8e52c80995969770e8e5b2d2d92c3b) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](resources--securemesh_site--reference--group-001.md#canonical-f08ab8815a1b7860cde80e5561ea3a615119e4d7a0ac5522ea9764c68f3805af) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](resources--securemesh_site--reference--group-001.md#canonical-75d505a004cc4d9d87f41b52ebe41a8d8ab25e26b01cbeee9c5d88f285c1baac) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](resources--securemesh_site--reference--group-001.md#canonical-57d74ad8eb087bddf744c3e57bd40bc90ac77a0e8c41ed0a73147627bc36abba) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](resources--securemesh_site--reference--group-001.md#canonical-d47c86b565c4b959fd298d19a0a9e96c06da64dc3fddec06476524398a501fe0) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](resources--securemesh_site--reference--group-001.md#canonical-6babfec20727d86a1fd514be09eb25361cd230d12e9438f4193a9ff6c8bd55de) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](resources--securemesh_site--reference--group-001.md#canonical-de970fcb77021fce25e2d752aa1f8423f0e857d0940d732bcde2280414a10730) |
| `coordinates` | [coordinates](resources--securemesh_site--reference--group-001.md#canonical-ba39b8067027a2263bb8099e5f6c8f1788d3c104e98248617c0869e8bdd8a217) |
| `coordinates.latitude` | [coordinates.latitude](resources--securemesh_site--reference--group-001.md#canonical-de885a4a3a102461976d0c29be695a3fe1c5e8a14d01dcb0434230956b2577bc) |
| `coordinates.longitude` | [coordinates.longitude](resources--securemesh_site--reference--group-001.md#canonical-46c5522dd710c5084f2f04eccc3bb301a9a6a33cbbdc583d3016c53bfdd5b287) |
| `custom_network_config` | [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-69d51591c455b22c1466f1156dad836e2ee561f2c8525cb8993582e40d0ee8d4) |
| `custom_network_config.active_enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-093a54889d320ab0882b5cc8a346fda6d47641697be16a39e5729187ac995358) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-01dd221a0babe26a838be505d5f4cafcfc9e052ac382f26029d1d3967c72ae5b) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--securemesh_site--reference--group-001.md#canonical-fc37e8809730db5103256bd9bd5ef9cb060438ebab695bcd13ee07e3ab0d14d2) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--securemesh_site--reference--group-001.md#canonical-46a99bb573cc4f83c97e2f5ca9c453ef96a30a6158e21b56bf3dbc2e90a95685) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--securemesh_site--reference--group-001.md#canonical-8c80be15b1205a6854ba69535cfd204cd7374e6b74e5a03b91b91f289730ae68) |
| `custom_network_config.active_forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-8395bcd613c346ca6e727b1036d20f900c3fe1a0a21c4c853b8540a2110086d3) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-d76add8c6125b66eb21fc142b75a264d2d2cae5550f3c71374c0e455f43a7080) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name](resources--securemesh_site--reference--group-001.md#canonical-2c3121db6299ea1169736abd087b6816b55415a03a0e868f72cb2a5480d9a53b) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--securemesh_site--reference--group-001.md#canonical-ad2340b12d69a9142e333e63ab420c0b4648335a6d05a0a2a948394e74547102) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--securemesh_site--reference--group-001.md#canonical-b36e6a6fc6a2def01de7f77904af65dc0c9318eea98d331183907c66594d39dc) |
| `custom_network_config.active_network_policies` | [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-97df10bc598d5cc71fdf2256142e970bad8adb24999bda5ab8e5ec952ccb9b16) |
| `custom_network_config.active_network_policies.network_policies` | [custom_network_config.active_network_policies.network_policies](resources--securemesh_site--reference--group-001.md#canonical-4631c3d357e69c16c1dbb6dffd361dbf0293e08d9d23fa2bf9063e434fc83786) |
| `custom_network_config.active_network_policies.network_policies.name` | [custom_network_config.active_network_policies.network_policies.name](resources--securemesh_site--reference--group-001.md#canonical-c90fc1757fce6055a6d4f90e3a00a218a4f7bd1e8ab34561cc091356dfb4a7a5) |
| `custom_network_config.active_network_policies.network_policies.namespace` | [custom_network_config.active_network_policies.network_policies.namespace](resources--securemesh_site--reference--group-001.md#canonical-3387434dbb481a75c2329351d4a6e118c5cdeaa08d2138a7c0bf7b37f84bbd2d) |
| `custom_network_config.active_network_policies.network_policies.tenant` | [custom_network_config.active_network_policies.network_policies.tenant](resources--securemesh_site--reference--group-001.md#canonical-844bcc9a9fe481ff0bd77c5ddf1c62b4b7023b44ab49939fd99d1097dbeb35c5) |
| `custom_network_config.default_config` | [custom_network_config.default_config](resources--securemesh_site--reference--group-002.md#canonical-1e534a9cf58f0b4ee1ec887f62e5a4ef63f1cbee907039307b51d481feda63c5) |
| `custom_network_config.default_interface_config` | [custom_network_config.default_interface_config](resources--securemesh_site--reference--group-002.md#canonical-e71561cef4573e64c5092d73aa39481f9b032b65f51d103ccef4d92360460822) |
| `custom_network_config.default_sli_config` | [custom_network_config.default_sli_config](resources--securemesh_site--reference--group-002.md#canonical-d81360c0c958eff5dbbbe35fc81fb00b382257e3fbc35199ccb14b95e3500a78) |
| `custom_network_config.forward_proxy_allow_all` | [custom_network_config.forward_proxy_allow_all](resources--securemesh_site--reference--group-002.md#canonical-0f5df94fa0db2c7dae3a0a273ef4eb69d51d2a84d826213f678faeacf76befa8) |
| `custom_network_config.global_network_list` | [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-d250c7951bf45636af16efa0b8b52d03040b37736f400353f5634073af09b061) |
| `custom_network_config.global_network_list.global_network_connections` | [custom_network_config.global_network_list.global_network_connections](resources--securemesh_site--reference--group-002.md#canonical-a022c2908e010fe2f8e84752e55f0f7f7440497fd77d9b66b7c1f1fcf659451a) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-de778b64134ec6c2252c96a4ac8113fe5c4508eabc81f43104c5e006ff8ce01e) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--securemesh_site--reference--group-002.md#canonical-547f96001258bf096bc828cc47ad6f6e8a77cbf9ab8303e95ba115eea5cef13c) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--securemesh_site--reference--group-002.md#canonical-7a1faf243603daf5898dd1ce72cf8f286da8a44a1684bebf667616c507890aeb) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--securemesh_site--reference--group-002.md#canonical-97d5bac898ba4f76c0d256ebd74aeac84efff401a045d73b5ffd362694aeade2) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--securemesh_site--reference--group-002.md#canonical-950c5aef9a2a75b0550574cd0eadd1f5ebd5ed5a91f5a7385e6ff3bb6c7db89d) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](resources--securemesh_site--reference--group-002.md#canonical-5c6a256991d77ea3ae29693088d2fd99a3865b0ab5bbb4e4b575257024a08e29) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--securemesh_site--reference--group-002.md#canonical-1bb1376d28e02f8151eb9778d32622a4ec8b10848bcc6fda640290923f086203) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--securemesh_site--reference--group-002.md#canonical-0e9427d039e0cac78b268ca6e7276fb44884ffbcb84e980137b71836a5bd2edf) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--securemesh_site--reference--group-002.md#canonical-1131d19ffa44d6e3c05acfa6df73b8b0066651fe1147b2e55d2952ef81b09508) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--securemesh_site--reference--group-002.md#canonical-dd189c45cee9198f9ad6e99c03e3d75f6276d75198e74b8d31df32cafa9fc03f) |
| `custom_network_config.interface_list` | [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-cdb774f115e6594834eb9eb8a4097b315999a2e3003aff24728e15ef53224210) |
| `custom_network_config.interface_list.interfaces` | [custom_network_config.interface_list.interfaces](resources--securemesh_site--reference--group-002.md#canonical-8e7d3af48e95ee400a25659c76ee3faa99f6bd7e0d9b404f5d901cb6555b6ae1) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](resources--securemesh_site--reference--group-002.md#canonical-592c599285d9e5b88f92244ff93dfb4e69247ea3c030326a8212755477310d50) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](resources--securemesh_site--reference--group-002.md#canonical-da0c6414501960efb7f07ad4ea4808d20c6a0b9ef731a619247eb50fd1d5b5fe) |
| `custom_network_config.interface_list.interfaces.dedicated_interface` | [custom_network_config.interface_list.interfaces.dedicated_interface](resources--securemesh_site--reference--group-002.md#canonical-878d1d7d65729ecb40cb8bebed40b01c8f5ecc3b2dc717df0e4f30154052eef0) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-d0fa86b3741f5d8715cc7f300ff4d7431905ad1537e89c787d970854c24a7538) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_interface.device](resources--securemesh_site--reference--group-002.md#canonical-55ed3432df8f3d986da835ff67d0ab380edf8931ce374d96b445d5d68db44545) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.is_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](resources--securemesh_site--reference--group-002.md#canonical-c0c6aefb7235b348e2118cc658fa10bad218f4644b683c3e7b5ced31097e66cd) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](resources--securemesh_site--reference--group-002.md#canonical-f6248a25fcab8bf450529231ed8a19c7aee98ba08d15ddc9c3cbef4ff6c5b5e5) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](resources--securemesh_site--reference--group-002.md#canonical-865204eb0c13f6fb722f770b6210c6a287155ac7ef93b7952ce76574a15da5a9) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_interface.mtu](resources--securemesh_site--reference--group-002.md#canonical-ce18c8a2350cbd90f1f7eddd955ff9a6ed7b783b1feb879d67f4117f24487b57) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_interface.node](resources--securemesh_site--reference--group-002.md#canonical-d162a6b42d5b6c441195364c80b509252b2bc708a9247bdeb68d4259a02f461c) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.not_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](resources--securemesh_site--reference--group-002.md#canonical-f9838a6cf30c38226d6820c88d9f872daa26be13024f5594cc8749dd707a3467) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.priority` | [custom_network_config.interface_list.interfaces.dedicated_interface.priority](resources--securemesh_site--reference--group-002.md#canonical-2a47b682541abf9bb5926d49906150629e13b069e1298689dec22a86737162da) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface` | [custom_network_config.interface_list.interfaces.dedicated_management_interface](resources--securemesh_site--reference--group-002.md#canonical-f1ecc03e6177f0b4fca23f9dc48306cb0a04c8b1dd07e9fb06c911b40342218a) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-bbbf7d95be06600d863e2b379d66f3a70d67caa901905c53ee263b265cba764e) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.device](resources--securemesh_site--reference--group-002.md#canonical-3bcbb92d92a9d58f11e09b8600d9e6ddf25fdcf8301d1b461b7ffd85d08f90c0) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu](resources--securemesh_site--reference--group-002.md#canonical-d93604954cc1c351764739c72ffac4eac58af507595621d40208760aef6a9374) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.node](resources--securemesh_site--reference--group-002.md#canonical-4be648ff7efedae5b253282c2da33313296c62743cfac28cff8b86cc2e4641ef) |
| `custom_network_config.interface_list.interfaces.description_spec` | [custom_network_config.interface_list.interfaces.description_spec](resources--securemesh_site--reference--group-002.md#canonical-eca76b3e20bb01d93e6a0c39fab923910404a33d8561c8f96b2034fb8382763c) |
| `custom_network_config.interface_list.interfaces.ethernet_interface` | [custom_network_config.interface_list.interfaces.ethernet_interface](resources--securemesh_site--reference--group-002.md#canonical-4d0e46f90782a2adcd4ccc14279f247ed05b7bfbce80414d0b11a863ce94284b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.cluster` | [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](resources--securemesh_site--reference--group-002.md#canonical-1e74390e3ac68af7e04f6b2e9b0b7d223cbdf6a3e1c89bdae51f853b70c402c0) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.device` | [custom_network_config.interface_list.interfaces.ethernet_interface.device](resources--securemesh_site--reference--group-002.md#canonical-1a79af138f8219ceaef4806cdf8817faa752654a43e2a8df1545cf7ec8ad2992) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](resources--securemesh_site--reference--group-002.md#canonical-e0fe7a4b00a35e5f9b30bb9a6c723e2c71ecdb4fb8fae160fd79f6bdabbb170b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--reference--group-002.md#canonical-5ae89f664dd6f52dce10ea06124f604a25c4a7ae0dbbf0a6a64ff89bcc737855) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](resources--securemesh_site--reference--group-002.md#canonical-91bf0c85f4ceca4a182fbaa91a43e636c4735e7abe14016c275d9276eb10d81e) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](resources--securemesh_site--reference--group-002.md#canonical-12654f3b732dd3c2347736af9540ba09905c5ba91763c81dd9a933615c298277) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--securemesh_site--reference--group-002.md#canonical-5cfe8b437ebcd9cffc5c174a4a267932509b6ae8bc661f7a371c4e5e30a8254b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address](resources--securemesh_site--reference--group-002.md#canonical-38f950bc2f053dce1a1feea5b1fc3e642f540a0f13e078a664c1281d2b737495) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address](resources--securemesh_site--reference--group-002.md#canonical-f28ef46b8b43633e1d2ce89b218109c4f6bb540601a75d5c2bcc57bf4d8b2454) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--securemesh_site--reference--group-002.md#canonical-1db3966d82cb5d9de3554373138d285534265c4658c6084d486da1b9735a2a06) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--securemesh_site--reference--group-002.md#canonical-227aa1033783ddd815729630d55a055defac5e0e4bae8ae75d89222b3d7c4ab6) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix](resources--securemesh_site--reference--group-002.md#canonical-081d896a27c2ab9b7cbfeed601f96ef40b725bc8917058261367efcd71f13c39) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings](resources--securemesh_site--reference--group-002.md#canonical-1de3503d7bbf0bb62844ce0163b9005e5207146c50ef3b3bb17a850e313c35f9) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](resources--securemesh_site--reference--group-002.md#canonical-cc9fd6d386083ab72c503ff51ec9594dd2a31d3ff3d2765d97e54bc01e7161ca) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](resources--securemesh_site--reference--group-002.md#canonical-36698311660fca1c3a1dcc3971d3ef93ce48feb385c8d1d60be7d1fc226c8f98) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](resources--securemesh_site--reference--group-002.md#canonical-dd19c874b6fe7b59f72379b95227dad041a16fa2f89133b33ca7fdb3b3b5c9ee) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](resources--securemesh_site--reference--group-002.md#canonical-4515fdc34a10419bd82854519225f68156e1963a0debdb419e99b8800cf54810) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site--reference--group-002.md#canonical-69948769542b13402a201cd423518919325739520bfddfce5249bbb14f57ecbf) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag](resources--securemesh_site--reference--group-002.md#canonical-0537fdec3ad595f13149e3b3c788d6731be544faf2fc75910cce12fa9b847bae) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map](resources--securemesh_site--reference--group-002.md#canonical-35fda78ef33c34d26dc60305fb4df1c2eaa93e778d98663b64e6bb6487334a39) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](resources--securemesh_site--reference--group-002.md#canonical-26bc16ec998a2d188ec32150e357ae97e895da7fa1f324f5e0fd4d73c69fcbae) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](resources--securemesh_site--reference--group-002.md#canonical-d5249df9e75c606acc564542651598d8daa21f105431f9cae162cb5eda070338) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--reference--group-002.md#canonical-f98f7b21d85a35c530d4857df1bee4fa7d1eafc1f66e50d3e3569582ca9cff9b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](resources--securemesh_site--reference--group-002.md#canonical-fa168e5c5d93f4e46e5cff5cabb60c305d08ecca68022cc2001e5eb0adb8667e) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--securemesh_site--reference--group-002.md#canonical-6f4864eb1c53d4e51b88b0e3fe2b36801a3c87d1e49b17c9dddea0a079398a81) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](resources--securemesh_site--reference--group-002.md#canonical-c9f9b655cadb1154b1a4176021dc4be3d99bdd9877cf8eb43f02aa7bea6e58d0) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site--reference--group-002.md#canonical-d93e1d324e39685a2cb637441980183a4d7a8aefb0709fde64c6aee6e8c33a6e) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](resources--securemesh_site--reference--group-002.md#canonical-a98a689f2c6d0e17616d372a660a35646dc79dac6be38f4339602960beb9b3c5) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site--reference--group-002.md#canonical-a695c7d026815d365ea64e842e1ecd4841881e787c19c037e93278edb4255db5) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](resources--securemesh_site--reference--group-002.md#canonical-80d9e22b43d3bab84c1e557c6c5d172e440765c902dc9db72d0d327d3700de9f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site--reference--group-002.md#canonical-94b3c52daf750462a7c11e927af108ecf4f155c00f37c76efef8a8721e482a39) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site--reference--group-002.md#canonical-3935ea5b1eb69f180c47ae298337cc1d6a30395d99b701edf527f1e977c0d678) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix](resources--securemesh_site--reference--group-002.md#canonical-c40621e02df7f21f7492f896d2b1114f75779734a6f85c328ac783263f976be0) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](resources--securemesh_site--reference--group-002.md#canonical-165da0ebee62bd309af6231b58508ee6248b10c9239a112e5e199748d3220fb8) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site--reference--group-002.md#canonical-e0efc2bf3fedcd0ad58d6628ff3bcb7ffb07c3d7c2fb98fff25f176beb94ee1a) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site--reference--group-003.md#canonical-7c36847faa7ef2bca0554df7cc6f2b399cfb43385ad9db794d5a6d43cd63a6f2) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site--reference--group-003.md#canonical-0e9e087bb0643e3359239efd594dd7ac6fe4c87adb257494be46213d15058117) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](resources--securemesh_site--reference--group-003.md#canonical-c5b9aa80820741a83ae887c6448eb29e02ff625ca1741dbe6421ed3f571d8e76) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](resources--securemesh_site--reference--group-003.md#canonical-4000d2f2cb2e40ee9de995ab61e27e3555026d7ddf8039d30096cf861941268d) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site--reference--group-003.md#canonical-33982099228c03d682234848c5ba4090c970e637db594c8376b048f061ea84bb) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](resources--securemesh_site--reference--group-003.md#canonical-86a3e9208b143264a6aa2eead34433b69d687b04dfd0e29ebab8e30a98c3292d) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](resources--securemesh_site--reference--group-003.md#canonical-b7b001edcfa1ad20f2f25a5b47746f53268f62aab6ff8e260668c2e4f03158e8) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](resources--securemesh_site--reference--group-002.md#canonical-6c1bf3d690a074af57f88587396098b5e6d16a2f26f96693e5f5a6527a4d733d) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-2360b733f6eff730a349c86a5a9b0ee0dc927995d9eb25375b76335a7e018334) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-0ba909893d640c020de022ae9ec53058a105059cac900944aee498a51197dc79) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.is_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](resources--securemesh_site--reference--group-003.md#canonical-9cdb793983f08f0cb224d3e1ccc1db0d96e4b9a98991e9c35de1e01f168d7bf7) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](resources--securemesh_site--reference--group-003.md#canonical-751637a700eaf9f5ba622c8df367881ed1180b6cc2812103f34a2865ae3efa73) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](resources--securemesh_site--reference--group-003.md#canonical-fb659ba168ec1345a738f8f71c0e4442ca50236561262745a9562fe96d504b1a) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.mtu` | [custom_network_config.interface_list.interfaces.ethernet_interface.mtu](resources--securemesh_site--reference--group-002.md#canonical-9168adf4b75080fde63b54f0379be1f8c9569acdd5d5b85bfaec6665015a4833) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-5397d25db26a487f41560e987fb3e7c220e4e717a5f4633085bb3692c7e1357f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.node` | [custom_network_config.interface_list.interfaces.ethernet_interface.node](resources--securemesh_site--reference--group-002.md#canonical-29cee435c514a8c77ad3ed8d0a87a3f31a58bd800fe22d71a57df1c08588d3b8) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.not_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](resources--securemesh_site--reference--group-003.md#canonical-9751428aca492e88ff26d679f2f2dadfa4dbcb4686f5a18100b6899cda4f657b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.priority` | [custom_network_config.interface_list.interfaces.ethernet_interface.priority](resources--securemesh_site--reference--group-002.md#canonical-0d2ea6f1d2aa5329e0ec8fc955754a7be8fc6cc693ab9ac97e1ec82fdef519c8) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](resources--securemesh_site--reference--group-003.md#canonical-d9869451521850210d85bbe6ea7f88540526f6424f4cbdd7b1db06895d1f0084) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](resources--securemesh_site--reference--group-003.md#canonical-88f8d39caff4bee10325da3cbaf497356b1982f559c566623a720e3f94986580) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--reference--group-003.md#canonical-87fcd703c96021008275309ac724284dd2315ab1a5d218dcb90353bef5e39fb4) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-31c9c01ec74ae38f3c366942f3c164a892c7424bdebd82e694d2f038a21ff910) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-8152a8a46c5b550a4f3e36d00e3d07fa7ebc0092d7d3455631c60e61cef3c9a7) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-ea6a6f84054830fc4d1b8039641826d616e7000f70013430a0487b4138222352) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw](resources--securemesh_site--reference--group-003.md#canonical-a5993ca0d993c21bbc9ee6f60a7bfbbea17296bf53e14b8a479d22a5385a3f2f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server](resources--securemesh_site--reference--group-003.md#canonical-d30acaed5ebf601ea2408cd9f830fc4ee9a617262547f75abe6dd1ef33a1ce9b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address](resources--securemesh_site--reference--group-003.md#canonical-d803c4e2cdacd82f3aa86825300ae668ae378d6d5f6d98224e057e8cab1c1721) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--reference--group-003.md#canonical-9d1af69e45a87e3e4790db2bb245e270a409a548d2a765625f3e6c7984381d78) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](resources--securemesh_site--reference--group-003.md#canonical-42f6fdfe8e38ceea71df3c306b45d04c1f330b8df085aa96ebe0f9bb12921a04) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](resources--securemesh_site--reference--group-003.md#canonical-d0b40618cf7bf5434d85aaefae6bac663ee32ee56b04a33d356601df7079dd19) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](resources--securemesh_site--reference--group-003.md#canonical-80eed67f9f45a24e56249a4b05c9c55ad6edbb486f5185453f06b2d4f5bc090c) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw](resources--securemesh_site--reference--group-003.md#canonical-27ff6e157b3d88bbb237bd82c97a91ea0249ed7b0cbb8bc43742c2f5a8c397be) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server](resources--securemesh_site--reference--group-003.md#canonical-92b56288c360e473cbf64c5fe75380e5087185a60121d14430d7e078bb4ddc62) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address](resources--securemesh_site--reference--group-003.md#canonical-b8227539155a575d37038c52af206b8685d91a36a0af3f627e833e4694952a48) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.storage_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](resources--securemesh_site--reference--group-003.md#canonical-93b34b0eaf68cd8e424ecaae95c98f6435ab3d6bec5ae72d6b9ad4e6a98c786c) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.untagged` | [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](resources--securemesh_site--reference--group-003.md#canonical-31f5fb27852e63d4dfb28ac0c6c2c8c1ac987fd33ce0b642132d4c1d7a4b2d97) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id` | [custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id](resources--securemesh_site--reference--group-002.md#canonical-c9889461ed95f0c57c1a036fb0b41c2b2620228f9e457a2eb925a2819e35d28c) |
| `custom_network_config.interface_list.interfaces.labels` | [custom_network_config.interface_list.interfaces.labels](resources--securemesh_site--reference--group-002.md#canonical-1a0718c198f72b70d8debe16bcd40375e4dde15decc1ccda4ccf0f2ec64ca265) |
| `custom_network_config.no_forward_proxy` | [custom_network_config.no_forward_proxy](resources--securemesh_site--reference--group-003.md#canonical-cffd171c355f739b639d2e89711f0d07ed0748c6302fbe5bc0c39b65062b2f55) |
| `custom_network_config.no_global_network` | [custom_network_config.no_global_network](resources--securemesh_site--reference--group-003.md#canonical-d3d85916ac7a92da937dfbf78a6c67a08c5199843b57efc1226ac660ddcbcae4) |
| `custom_network_config.no_network_policy` | [custom_network_config.no_network_policy](resources--securemesh_site--reference--group-003.md#canonical-7ad9c7585770b0ae0b60bcc75cde0dcd86cc6e9192aa93443613f3ec6772fe81) |
| `custom_network_config.sli_config` | [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-69db66f813314271f42815a3276826f0c3fe67327b7087c50c16b3239f68f1c7) |
| `custom_network_config.sli_config.dc_cluster_group` | [custom_network_config.sli_config.dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-4021f90b723e85b013083a6c8ff57f923abba10ea34e45a9ac6ba48938e4436e) |
| `custom_network_config.sli_config.dc_cluster_group.name` | [custom_network_config.sli_config.dc_cluster_group.name](resources--securemesh_site--reference--group-003.md#canonical-24c80775a27d8f828c6837ddc43aae8b111a6807f449ea94d6c3ad53718f4885) |
| `custom_network_config.sli_config.dc_cluster_group.namespace` | [custom_network_config.sli_config.dc_cluster_group.namespace](resources--securemesh_site--reference--group-003.md#canonical-3c2a640940183aa8a7c71a8366786442f02d053eee9be9519d7fe11ee4576f69) |
| `custom_network_config.sli_config.dc_cluster_group.tenant` | [custom_network_config.sli_config.dc_cluster_group.tenant](resources--securemesh_site--reference--group-003.md#canonical-b64c2f5f1468eec15e4f58ea2ab9ea043e8e211594c78892591ee08ddba43731) |
| `custom_network_config.sli_config.labels` | [custom_network_config.sli_config.labels](resources--securemesh_site--reference--group-003.md#canonical-7f9ad52ff71f25ade0ae070135a7ef5d7cd621c86ce5b7f193b13f67e3ba63a7) |
| `custom_network_config.sli_config.nameserver` | [custom_network_config.sli_config.nameserver](resources--securemesh_site--reference--group-003.md#canonical-1a4d5fef52fee6ac22e43ed808c1278047f898f9f393cd47f65e74f2d03f528c) |
| `custom_network_config.sli_config.no_dc_cluster_group` | [custom_network_config.sli_config.no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-56dbaa4ee223f5518163df22aa401130abb57639ee37960041cd9c866d72ee36) |
| `custom_network_config.sli_config.no_static_routes` | [custom_network_config.sli_config.no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-31e557e4eb3a92900e947663296a3522929618999a631433ec09d9d67a5b8957) |
| `custom_network_config.sli_config.no_v6_static_routes` | [custom_network_config.sli_config.no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-8fb72d6a7a919fb56fe155fb4fa345adb8cc80f4d574f5683e66b44559fb80b9) |
| `custom_network_config.sli_config.static_routes` | [custom_network_config.sli_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-2f4f8d25dd89903c17380c2ce1b2ec83ff9e22e3fb9342588597482f434289d5) |
| `custom_network_config.sli_config.static_routes.static_routes` | [custom_network_config.sli_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-6c6cd73e40474708a24c8f28266368d71939e45c3acdab368d85062514f627ec) |
| `custom_network_config.sli_config.static_routes.static_routes.attrs` | [custom_network_config.sli_config.static_routes.static_routes.attrs](resources--securemesh_site--reference--group-003.md#canonical-3b06ee28a9e8227957fc38f451516e6286a198ad9e8f9f19f42874d46a2fbf0d) |
| `custom_network_config.sli_config.static_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-003.md#canonical-10e6798f4851d9534948a6737f2eeaf22e224a5b6f6dffdbee54aef9f9dd0a27) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_routes.static_routes.ip_address](resources--securemesh_site--reference--group-003.md#canonical-ff28d2a8ee119e7f1002b78a2da67a50aed0bf62cf76f517add74dda6e4db429) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-003.md#canonical-e059aa5717451e731a2b75d25563c5548bf1de2ba5c6d101d0f0a72ea3cd102d) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-9dcf8fe527a661e4553720619ecf3abbc5ced309df49cb6d0c11e106e3ae32cb) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-3d0361dd527f5eb5a086bea5f58bd9c29075e0f2f56ff4f6763ca71a1ae4a860) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-003.md#canonical-8c67f4aad130c5fa1566c3a1d96e7dcce2848827dea84159676b06dfc01a5ecc) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-003.md#canonical-e89f37fe293d8523f2dc50a409f61c7aadcac2eda00dc7fa1b81793ffcf74aac) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-003.md#canonical-6954231e3ad4d3f98b0bfb0b6edf1a413c478b36a0455721898f3ec135354d1b) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-003.md#canonical-cf4b02be004ee18a91849ae149155b207459d21566e37d9a5127c248b061dc01) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-003.md#canonical-dac4f6b66b69df225094c8f088b026252086e1d447eb2cff201b7e463959057d) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-003.md#canonical-971ec2584220952ff79ad32653c1bfb47caf5d507368038aac631ababc86bdd3) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-003.md#canonical-7b4797a1fb967999399d111d00af6c73225f72216de6fa2ecbab8c17e312b4c9) |
| `custom_network_config.sli_config.static_v6_routes` | [custom_network_config.sli_config.static_v6_routes](resources--securemesh_site--reference--group-003.md#canonical-95d31947b1dfa80f21e6c969d67d8bc02d10754cf19b10117e2b63490fe83a6d) |
| `custom_network_config.sli_config.static_v6_routes.static_routes` | [custom_network_config.sli_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-0c84a3abe1392d012652be1c980255604ddbfd25f1c96d17df42c764aa9ee971) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.attrs` | [custom_network_config.sli_config.static_v6_routes.static_routes.attrs](resources--securemesh_site--reference--group-003.md#canonical-4641fb09c1649bb4b9e3826866e484c62579f030b11ff347c82f86fa3e9c2d2c) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-003.md#canonical-7f8ea8c7e8bc3e64d992e2407a1dfaad2f8fba5adc9eab8b2be5d47c3cf957d2) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_address](resources--securemesh_site--reference--group-003.md#canonical-809d0b53bd006a2c2b20ebd06375d0b8af37130a6f7388480aee736ac451ccd0) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-003.md#canonical-585b8fc049d7c831699c516dab552422b97203831b918c2e0e585f6f3c520da3) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-e0c1ee972c322e617f313c5bd780149b1fb5454ffddaef845d152db448380aba) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-58b0c1b381991a290a6c1415de7b5ae5ea51862d07bc7006003d9b99e50602ba) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-003.md#canonical-6a0333946c490ec1fba27bcdfe7431811fcd232d056b7ec01842045aaf3b540d) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-003.md#canonical-2e575b9d70be538501ede4983cef6a827dd95813142385abee8ed6fd27554eb0) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-003.md#canonical-616cec34e76c42b195488ee9e1891b437e2bda9d5b77b7223d952bb9c4bb7b3e) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-003.md#canonical-fbae24fe3c0eb722c4dcb92914b6cf23463b5a82996efd39fda2082e6160b165) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-003.md#canonical-8dd99a533f87419ab8e1578bcb43edfe169c6a428cc09e9cbf044bfc2dca763e) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-003.md#canonical-f3b822c49b2d374935f4a417c49a438085fb7cfab78068ca25dbb4c6f8166413) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-003.md#canonical-ccf314806423f08df43143c9f204f348fc4a8452f69bb2709d7c0db7f7b32b65) |
| `custom_network_config.sli_config.vip` | [custom_network_config.sli_config.vip](resources--securemesh_site--reference--group-003.md#canonical-b18955eac1e27c28f9e202b414ef928f7cacf4646695a1aaceab9ae03ee538f7) |
| `custom_network_config.slo_config` | [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-7ef42041475dea229e80cc49fe15352e4bd31e62313ce9a0d9caa7709ad7613a) |
| `custom_network_config.slo_config.dc_cluster_group` | [custom_network_config.slo_config.dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-a9d2a8164dca7d75863e5e0ac4bda35333fa7dabba3bc398b6ea2d1fb16257ac) |
| `custom_network_config.slo_config.dc_cluster_group.name` | [custom_network_config.slo_config.dc_cluster_group.name](resources--securemesh_site--reference--group-003.md#canonical-59695c1f4f8594c56a0fdaba0815d12b47f76894fda8558ffd7f88e80b139970) |
| `custom_network_config.slo_config.dc_cluster_group.namespace` | [custom_network_config.slo_config.dc_cluster_group.namespace](resources--securemesh_site--reference--group-003.md#canonical-01cec6549d1e6aae83c0c2157514b0cb3c189473b280f87e6d6067d542d8efe4) |
| `custom_network_config.slo_config.dc_cluster_group.tenant` | [custom_network_config.slo_config.dc_cluster_group.tenant](resources--securemesh_site--reference--group-003.md#canonical-66cc28a8437edc54f0a9655983b192f81a58d583a6a775874aaaeb42389d28a8) |
| `custom_network_config.slo_config.labels` | [custom_network_config.slo_config.labels](resources--securemesh_site--reference--group-003.md#canonical-b39bd68622a2be067b44ba867613aa3eb95f1e45547b6ad8f8d14942aac7f28a) |
| `custom_network_config.slo_config.nameserver` | [custom_network_config.slo_config.nameserver](resources--securemesh_site--reference--group-003.md#canonical-d428c8c058ac43fbe360a6b608e367b8683f4a003e2e2607c055ba9941451420) |
| `custom_network_config.slo_config.no_dc_cluster_group` | [custom_network_config.slo_config.no_dc_cluster_group](resources--securemesh_site--reference--group-003.md#canonical-b0c1b95d82a6733445bdcef0cfc16d859e30af715d8c3a5543c6c8a7b0ff7e33) |
| `custom_network_config.slo_config.no_static_routes` | [custom_network_config.slo_config.no_static_routes](resources--securemesh_site--reference--group-003.md#canonical-36aa67ee9be53ec8fc6416890148e8c91be5363905c76a2323862512f3ee2ac4) |
| `custom_network_config.slo_config.no_v6_static_routes` | [custom_network_config.slo_config.no_v6_static_routes](resources--securemesh_site--reference--group-003.md#canonical-81228d223e3f5bb7550d4b6809790c610ad3864245467f950f2cda3b4adcd101) |
| `custom_network_config.slo_config.static_routes` | [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-4f0db6c25222f8d47734df423aec3fdd70fb3738c62dc33fc424ebad2731c36f) |
| `custom_network_config.slo_config.static_routes.static_routes` | [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-f0c71922b4bd1d6b394e0b8d9f2b061d77223d5bbceca5f32290f210748fc880) |
| `custom_network_config.slo_config.static_routes.static_routes.attrs` | [custom_network_config.slo_config.static_routes.static_routes.attrs](resources--securemesh_site--reference--group-003.md#canonical-fcc2902fbf9918b6d7b18ddfa1ed9e9262570cb37ab6b793124c4f643e1181ac) |
| `custom_network_config.slo_config.static_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-003.md#canonical-bdfd2e0c3d5aa245d22cf70cb2e5c85bd05823baa00a5c2ff54b9d596924bb15) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_routes.static_routes.ip_address](resources--securemesh_site--reference--group-003.md#canonical-9f5f199188ac0f2062ad0c5a060742cfc06c2a2678a3c7733bbde31fbbeb0a94) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-003.md#canonical-fd21e7da6174be60889b5296aa70ff75ced4310d1b9004d13b3db00ef146d397) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-4d32ca5d1454a0db8c0d79b869bb14bc79c08fabb3e713f45a844020e848eece) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-ddd96e15d5d050dbf713aceffb9c637807b90bc216d44009cdd508b606a85da5) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-004.md#canonical-ef5e7f72db09d7d0d9ea6c1ec43f9c46a2a80a52a150dff2fb1f313d67b2b7fa) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-004.md#canonical-b5b935065e7fcb0bc938168dc97bcfbf221a7da885b9a2f5b2a7696411076d9e) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-004.md#canonical-1e56305fc1ce0bfe476fef1d9ea4311dddc860787a74d33983db13d69721d3c4) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-004.md#canonical-ebcd2551414e597ac198baca04da5bde0cacfb025e204e7a0e466815c440c828) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-004.md#canonical-02f57965552f0b4675eb16a66bb41f58a290bc8230a83ffceee93d4447a3f086) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-004.md#canonical-f3bf5dc3ba5d8b9d572519cd8501fe647393f91925b6a61a8f6b63aa4145b760) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-003.md#canonical-c011cf2db9b7513a0183296dcf2f9130ff1b4f444fe75d0d9e5e829c4510fe15) |
| `custom_network_config.slo_config.static_v6_routes` | [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-a904b1fdd92e2dd12cdf49133e6e80ca5a8a671eee65dfcee29e7fdb725bce9d) |
| `custom_network_config.slo_config.static_v6_routes.static_routes` | [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-3e0246b764b4fb20420a912221315c3c856adb7de9c57605484c6b075dbebd0d) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.attrs` | [custom_network_config.slo_config.static_v6_routes.static_routes.attrs](resources--securemesh_site--reference--group-004.md#canonical-e4abed99cfab9e315924ad1f8f7a5499c0e4380e7132a39cefc9090a004290dd) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-004.md#canonical-70bfff4ba4cd4edba9622c865525b4e6c2b506d7ba8ff70e7f86898a9dbbc29f) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_address](resources--securemesh_site--reference--group-004.md#canonical-acb0fc34bee658e5046e0d5e2650f1018bfeeb54685f4cfad87f277d55571584) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes](resources--securemesh_site--reference--group-004.md#canonical-b4ef4c954332b6bd9881597337ac4b74d15e4ada8dc3ba3dccc6accfb9c49242) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-311cdb0729c05aeedd98201e612d0df0357aac2725d5aace8ffe118e7f39b088) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-a70845ec6315fdd7e64db083d8ea4b145a68cbd8fa6d1f0b7072d74dc6851f1a) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-004.md#canonical-a670428ef76047a7ab2687a261fbe5a25ee7297dd04daddd1463c61d1fd65cd0) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind](resources--securemesh_site--reference--group-004.md#canonical-fd9c8405b29800ca92f5394e175f0362b965d7a4bc7dc1931080e1a464022a5b) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name](resources--securemesh_site--reference--group-004.md#canonical-d3b6e5e0ad420df54c6c4163c73ce1364574b492b0f9cff174404bd5387350c9) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](resources--securemesh_site--reference--group-004.md#canonical-6b653db4e1ad0e1221e577fcae7e68a9fc325f227165e3baeb5563380450f788) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](resources--securemesh_site--reference--group-004.md#canonical-9298f4fa3bdb9fff39ab713cb98e847a89d15c1fd4c5d0e4b0e141e7d8ae8a29) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid](resources--securemesh_site--reference--group-004.md#canonical-dba9bdebb224d568389f8b3e87465a35c0818da8bce1184f65cff70f14667c3f) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node](resources--securemesh_site--reference--group-004.md#canonical-4bfa73f630657482c707d01ea9476f4ae1d20acfc78c96225d64a88faa45e917) |
| `custom_network_config.slo_config.vip` | [custom_network_config.slo_config.vip](resources--securemesh_site--reference--group-003.md#canonical-5951a1862acb803b4eca63ec45e671355feeeaa50bdb5ab6954cac7de9007d25) |
| `custom_network_config.sm_connection_public_ip` | [custom_network_config.sm_connection_public_ip](resources--securemesh_site--reference--group-004.md#canonical-f0fe59ddbc39010fd803346b06a0dfd06ae175e7be0510456de61fd4b394dffb) |
| `custom_network_config.sm_connection_pvt_ip` | [custom_network_config.sm_connection_pvt_ip](resources--securemesh_site--reference--group-004.md#canonical-d911e9af9bc1b8a477d016be97dc2e2b07da59f1c1b166a80fbdf6e0f38c68b4) |
| `custom_network_config.tunnel_dead_timeout` | [custom_network_config.tunnel_dead_timeout](resources--securemesh_site--reference--group-001.md#canonical-81f59a0d3e280e597c287a48e016dca5be0abb312dba95641aec7830cf35921f) |
| `custom_network_config.vip_vrrp_mode` | [custom_network_config.vip_vrrp_mode](resources--securemesh_site--reference--group-001.md#canonical-ab1515d49d6315e9d223e84b756f59735d1c670f8e6911372e538107d30a3757) |
| `default_blocked_services` | [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-73afebe2d256020451a57e09c99766c20af7833ca105c606be8813c42a418f9c) |
| `default_network_config` | [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-871ddd6e9fca923ed67c5ad66b8affcef0f12462f8565e9538776a3e61f8efb5) |
| `description` | [description](resources--securemesh_site--reference--group-001.md#canonical-8b969d40126e45dcdfed605e968a3645963ab803c8e0855177a4807278f3f2c6) |
| `disable` | [disable](resources--securemesh_site--reference--group-001.md#canonical-72ea41d51fe48dcd341143852b4c8d210af090cd6493630a4417bc9a6139773d) |
| `id` | [id](resources--securemesh_site--reference--group-001.md#canonical-529d92e98ec3e450cb205a80d12ff3960ed8d9e53a3de28f31891ebf660e21d7) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-60a04a3b53b7ee4e977a3f20a3bf33612673816cf8e23ee80f0828844a5fefca) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-513cfbb07b012f77e4dd6c86a4b24df90399ee81104edc80f79c05480118e5bc) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-02ae73d7973662f340b89c01cf0507b04bb67b51cec717b99023bdc7c36a5638) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-6a3508f965c6874365e370c5fb1e2ac8ec60a65546e04a9442aa250138faf4c7) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--securemesh_site--reference--group-004.md#canonical-fcd4ed04036b22fa5930d40a646c68fe6e953e2f20a40883af11d9d199ccb565) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--securemesh_site--reference--group-004.md#canonical-a3def8c765fd65a9ba09291c64d9bdd59eb9cb7cc92f2d666e2b16a529d33fa6) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--securemesh_site--reference--group-004.md#canonical-16fe608a27b26482e9cd2a8ee977cc0f1e24f84589141b0262439d51a59f9e59) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-ddc1746b969a8572e3332bf364d17bdfdce24ac48d33b87de34dd70292385418) |
| `labels` | [labels](resources--securemesh_site--reference--group-001.md#canonical-e3b790d17ba3356ac782ad59389de40c47daada8cfb1e0610fcbd94d17a361b1) |
| `log_receiver` | [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-c3429a9023d31ca9356e3d2ceb1e9ca7237c8ccdd903da6dbe8252205cbf9bce) |
| `log_receiver.name` | [log_receiver.name](resources--securemesh_site--reference--group-004.md#canonical-880168d21c2f4f1d2ea5333ecb1cfe4b3c5a8a56f75f8c0a792d01b744070d8c) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--securemesh_site--reference--group-004.md#canonical-45100d480fabe69ccad69c66552f4264173d6cc89b9216b74d3540835f44d7c9) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--securemesh_site--reference--group-004.md#canonical-4a4d52b332d1366e79ecc92abb691e2571efe7b060b54b519a89a10f192e337a) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-1a185866c3136ce73bbc47c579c8cc7ef9d53505e40412b5e387d04158d63237) |
| `master_node_configuration` | [master_node_configuration](resources--securemesh_site--reference--group-004.md#canonical-9c95772323135cd8a70331c22a79957f3c75c30d1a108658788e4bfbfebe52bd) |
| `master_node_configuration.name` | [master_node_configuration.name](resources--securemesh_site--reference--group-004.md#canonical-867de7efda32bacb12c4d4ee3a58b0008c732d8664d3dd811279e65e9019258f) |
| `master_node_configuration.public_ip` | [master_node_configuration.public_ip](resources--securemesh_site--reference--group-004.md#canonical-0d4bdf755625ce5809ddc50eb4b9a3b7cb6e41dd7df006d18b3c1e3844e10169) |
| `name` | [name](resources--securemesh_site--reference--group-001.md#canonical-e70fe104dfb15b5e0ba4617a90e1ddd8a56d0187ffdd65227202ead68fb341bd) |
| `namespace` | [namespace](resources--securemesh_site--reference--group-001.md#canonical-047d58428e59531f5af489dcb2e723200e13fe2bc5b3826da80ecf2b56549e05) |
| `no_bond_devices` | [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-17491a1b8d51d7cc729febc7b394fbeeefd2c87cd44aee4af438dbf46de71db7) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-1a249b8b55ca5b0995a31662206e87d414bb2db0089bef5ad38111871ebfa1a4) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-d04e0a58c12a7d867762a5676ce962f7dda257a005f3923d95a9c87c1a4381db) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-d7f3e9231ade0350555feaf6a50217f60e986fa4d2efd97a4d4fdc5d1768bd60) |
| `os` | [os](resources--securemesh_site--reference--group-004.md#canonical-7ace6872dc9b8c6d170516a8433498d0a83fafab0b7b830465761914927b0946) |
| `os.default_os_version` | [os.default_os_version](resources--securemesh_site--reference--group-004.md#canonical-cd92a3cd16202fd5927a4179689fdfb5264e6f01431245dbaacbd11cad8ab28f) |
| `os.operating_system_version` | [os.operating_system_version](resources--securemesh_site--reference--group-004.md#canonical-2ed03b6e53c24249ca879c576eb32df96eb7ba4e66f2bf60348137d6685136ae) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-2086c93376d1db431dd641b4eb603ae3d64a5c3ec094bca55a212b51437d92f2) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-7d7fc8dfac75ca797b002996d9f713317c55b2048d1dabf7fefaf0311566603e) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--securemesh_site--reference--group-004.md#canonical-e4b0d757ebbe608213d7b5555c10172d9878f23edb5ede4d180314d6a17c4850) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--securemesh_site--reference--group-004.md#canonical-0d2d462520408d21a91197c46a29b53cd671fd5265e0790a89cd258767c230b4) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-47c5c38b2a80d0652acf45c4a150ff6266f29b005c3cea3980315c40aefae9ca) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site--reference--group-004.md#canonical-ead2ae05e2e008a3f6f842395a8b169f244bac2a5031a445000978a3c51866e5) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site--reference--group-004.md#canonical-1e54ae898297ac7a935bc66bcd4f7505fc6a0dcd39bea7265d1cd3c54f168e01) |
| `sw` | [sw](resources--securemesh_site--reference--group-004.md#canonical-ebd1a47213748f6d21b5631cdc1f5ea280664bb65f93be44b8d4b3514c92bfd6) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--securemesh_site--reference--group-004.md#canonical-523cec08581f6fc9d3b4b0997a9da3f9a9e1a54e310c5705622f10e5d12fa976) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--securemesh_site--reference--group-004.md#canonical-7e6cce4a55528ddc5ad808132fde43c6011f1544a0c2e1644739a6f2a1ddc647) |
| `timeouts` | [timeouts](resources--securemesh_site--reference--group-004.md#canonical-3c6815cb7a5ba39b1d2f28d9a4762f6e08d75f8bb1ab429662cf90dcc33579b4) |
| `timeouts.create` | [timeouts.create](resources--securemesh_site--reference--group-004.md#canonical-2f49a683e88c514a953d4084dc9c94e93a9d2bc5505b8ea3137a10481eaced1d) |
| `timeouts.delete` | [timeouts.delete](resources--securemesh_site--reference--group-004.md#canonical-91597a72b6b0c7c2a594e097c780954e8d6603f3585fc6817c2d9590eb41fc40) |
| `timeouts.read` | [timeouts.read](resources--securemesh_site--reference--group-004.md#canonical-a7d1de996d2b3aa8c40fbe5304cd66a23a2798fdf4486f50bddee9877351deda) |
| `timeouts.update` | [timeouts.update](resources--securemesh_site--reference--group-004.md#canonical-e48a8692d84734640045b4b068039e7d89020c20d80b8810a3d6e05b7651e2b6) |
| `volterra_certified_hw` | [volterra_certified_hw](resources--securemesh_site--reference--group-001.md#canonical-a1f7959dab3807650db3b45a8128f6488958dc18e2a47dfe5d5bb214d3533d4e) |
| `waf_signatures` | [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-7ff695956cfab12063a7891a9595a839c9b28064deb96d7610db5b7d001a0c74) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--securemesh_site--reference--group-004.md#canonical-e0a2e0fb03242afe6f34e284dd4fa2b494afe790205f6342d14cacb30cad20b8) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--securemesh_site--reference--group-004.md#canonical-9f796897ea6a52714a7aa563a691507ef56f3439e483504798b73026c23541aa) |
| `worker_nodes` | [worker_nodes](resources--securemesh_site--reference--group-001.md#canonical-68f1e051f10319f4ee6465087917802319c20bb53b9e68dab85714ac1c7342da) |

<a id="canonical-8d984ae3fae3a89262d8fa5acd38f30d37465bc0f0b11003b36752660c2099a0"></a>

## Next pages — Property reference / b351774ec556 / 15

- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-b46729d24d7c5fedc7541c50c6597214e39effced7311afc2973e351df32216e)
- [coordinates](resources--securemesh_site--reference--group-001.md#canonical-6af687786ff749e0698ebec1775e793f9421f4989d96ca808a763d39f55d5db6)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-9ce856978b9d1a3834c3b63a67ff7d42db5ed137144b2db86037212335e19331)
- [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-6598b3c9c785b28b0be5365f6ae8bc3b1d78ecfffef33c5a8ea18ada89496ed4)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69)
- [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-64b4390e24c8d3cd7963fefe17327c30c38775d5b9fc0c3a76696addb3c8473b)
- [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-31eb5954f742df58c87702640dbcd9722b462484fa62264b963a0128a396fc02)
- [master_node_configuration](resources--securemesh_site--reference--group-004.md#canonical-8c2774d9115547c29c565d10884a46a3d040cea98f976355ca3e3dace10ebd9a)
- [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-3f591357d9730ab49c161f878602fbcbad1f798979daab263ba21dd5e07fe37c)
- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-8876dd170b5fea86ae9ef472fdff18fdf1d5c6c740e6d812f7d8ab5751811350)
- [os](resources--securemesh_site--reference--group-004.md#canonical-127eee1dad136acb23d6139da4b0993a34521376b95d09aa218b63fcc060cafa)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- [sw](resources--securemesh_site--reference--group-004.md#canonical-401806d8a416f123ff291c4a5f724bf9513d7fad7f5ce8928e9b0ac494e6f2e7)
- [timeouts](resources--securemesh_site--reference--group-004.md#canonical-edc9ffffe4ec4adfdc3b384259fa9b388799b14d92d94a5b7fb79492e7d5346e)
- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-e3d3b3dd04a2677e5757751a4242bcaa7ace22229d58fc2ed27e711cfcaac7e0)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-929aaa5c89bf002f01fe05fe04275d3881ed8d130e574b0364b832b10f61de4a"></a>

## blocked_services — blocked_services / 62cf1278d6f2 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- blocked_services

<a id="canonical-7bf7f39c6626a3b36d80b5046bf19787fe5f064b23de1044d9d4e73ed0ddb6f3"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: blocked\_services, default\_blocked\_services; Default: default\_blocked\_services\]
Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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

- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-7bf7f39c6626a3b36d80b5046bf19787fe5f064b23de1044d9d4e73ed0ddb6f3)
- [default_blocked_services](resources--securemesh_site--reference--group-004.md#canonical-73afebe2d256020451a57e09c99766c20af7833ca105c606be8813c42a418f9c)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-d3f5c67ee1d30ba3e74be90a7164f8b7fda6b59fe69c2702cfee1e823bcc54af"></a>

## Direct properties — blocked_services / 62cf1278d6f2 / 3

- [blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b): complete subsection reference.

<a id="canonical-2d37a323d63844a22643daa96c1dd38a690c64ae3a8cbfed4668665fc0d7654f"></a>

## Next pages — blocked_services / 62cf1278d6f2 / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8847692dd2573c0711b36b5a7ad05924ef13fe4b321cd662713301a566cc575"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / f557a613bdcc / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d)
- blocked_services.blocked_service

<a id="canonical-3c14d732efecd597d23ff647c449b4b29e876f5f6a26d7c2ca92b3907813bd8e"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d8a282773f40fc2677281203935f78c36b448343024d68e81023a8fe899898b"></a>

## Direct properties — blocked_services.blocked_service / f557a613bdcc / 3

- [dns](resources--securemesh_site--reference--group-001.md#canonical-89681ef8029b412d66bfafe1ffd84ecd7687cdfebf498f06585b33519fbf4c31): complete subsection reference.

<a id="canonical-b240b118dcfe47656fc13e5f270ad31853647a045aac4330974508ad98a7654c"></a>

<a id="canonical-00de336e222690aaf92491a5d06f3e3ec7e266b139820e392d07365176e97e2e"></a>

## network_type property — blocked_services.blocked_service / f557a613bdcc / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ssh](resources--securemesh_site--reference--group-001.md#canonical-074f1c6718d472c8bd04bee8664e5e68c62bc028a3d14e97934f3bc06e567915): complete subsection reference.

- [web_user_interface](resources--securemesh_site--reference--group-001.md#canonical-13cb770212625897a67680a5756e71d3e7d4f45902d7277097ca618b237fc16a): complete subsection reference.

<a id="canonical-c0226106e262305ebb896aaf0ace38b555329a2df8a173418cf1fb23505c4bf5"></a>

## Next pages — blocked_services.blocked_service / f557a613bdcc / 5

- [blocked_services.blocked_service.dns](resources--securemesh_site--reference--group-001.md#canonical-89681ef8029b412d66bfafe1ffd84ecd7687cdfebf498f06585b33519fbf4c31)
- [blocked_services.blocked_service.ssh](resources--securemesh_site--reference--group-001.md#canonical-074f1c6718d472c8bd04bee8664e5e68c62bc028a3d14e97934f3bc06e567915)
- [blocked_services.blocked_service.web_user_interface](resources--securemesh_site--reference--group-001.md#canonical-13cb770212625897a67680a5756e71d3e7d4f45902d7277097ca618b237fc16a)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-89681ef8029b412d66bfafe1ffd84ecd7687cdfebf498f06585b33519fbf4c31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6267ba8d380d78a021785954b06e7d93fa2324cc120fac5a8d02fe2682bf9a94"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / df221f976fda / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d)
- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b)
- blocked_services.blocked_service.dns

<a id="canonical-175332b3a333ba6783bd9d2b0b3011046310b13f8b50d43ecb5899ed792a06d6"></a>

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
dns = {}
```

<a id="canonical-be9d5eac156fa4d17dfd46f8057ea3379ba3aa290bc20fda65bbee9a9abddcbd"></a>

## Direct properties — blocked_services.blocked_service.dns / df221f976fda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-635765fdd31f11935f1d0b2514eb376c92e04ae555905ef02f6d8102b250fe60"></a>

## Next pages — blocked_services.blocked_service.dns / df221f976fda / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-074f1c6718d472c8bd04bee8664e5e68c62bc028a3d14e97934f3bc06e567915"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bde92d9f624c6d2fa4134a93ea02d67f8185857a20cba8b37ca1e84fa83996ab"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / be1ba72e651b / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d)
- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b)
- blocked_services.blocked_service.ssh

<a id="canonical-091265a6fd447b4a1b3fd0660e46b0dee1cac4f35ed29919cd9a14cc8ae289da"></a>

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
ssh = {}
```

<a id="canonical-a2a2c4bcb8f3855cdd498c4fb1d4594583225de5cb20ebd02b636b1f709ba780"></a>

## Direct properties — blocked_services.blocked_service.ssh / be1ba72e651b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ba4147b5f4a1b41fc0bc17d413f901e28162eac7c40723563f1271d9bed882b"></a>

## Next pages — blocked_services.blocked_service.ssh / be1ba72e651b / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-13cb770212625897a67680a5756e71d3e7d4f45902d7277097ca618b237fc16a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8edddd5a3e5058484af84e8bdc2f8d7042ce1f9d4e91fbdc31d5c85f17549542"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / 4082b1b228cc / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [blocked_services](resources--securemesh_site--reference--group-001.md#canonical-c17357bbda7324de7bc0bdb08e153bfeb16f26aab71bbf958fe4d43a8270496d)
- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-3c3dccdb4341cfeea92b3a8a4a165f4fbea8a389679a0913baa39785ef76a964"></a>

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
web_user_interface = {}
```

<a id="canonical-1cee3fe235fb5808add499788ad23ecb456e81f4adcd0db2c1f0b24d8a411bea"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / 4082b1b228cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d0ddd2345cdf76b1321054fdcd60a90334faa938805d00ee5b647faa4c7b38b"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / 4082b1b228cc / 4

- [blocked_services.blocked_service](resources--securemesh_site--reference--group-001.md#canonical-eb784fef3d1f56fed3dd34186abcd71d7bbc0815a2e54fa00c968dcffd56d86b)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-b46729d24d7c5fedc7541c50c6597214e39effced7311afc2973e351df32216e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19cc5f45ce49b1ed97c23899a2f6742ff7f22e6d27117bbf79659edf3b9b0928"></a>

## bond_device_list — bond_device_list / 7d7417e776df / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- bond_device_list

<a id="canonical-a065551d6ebc597aa0290d29c4bfe0762ce6d861833f409457ff46b525cd360c"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bond_devices")}
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

OneOf alternatives in this subsection:

- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-a065551d6ebc597aa0290d29c4bfe0762ce6d861833f409457ff46b525cd360c)
- [no_bond_devices](resources--securemesh_site--reference--group-004.md#canonical-17491a1b8d51d7cc729febc7b394fbeeefd2c87cd44aee4af438dbf46de71db7)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0cf1631a7f1cbf7761c67fa9f152c9e1fe5a4c6c84fa60a9a30a268f3e7100b"></a>

## Direct properties — bond_device_list / 7d7417e776df / 3

- [bond_devices](resources--securemesh_site--reference--group-001.md#canonical-a6fba5275b44298b4a6eb8fc267ae614e34be734b147f7450ebcc2773dd148a0): complete subsection reference.

<a id="canonical-51ccb83e281560069f77014918c51db53a6e4fe14bb5a6c8f9342ddbcaf30074"></a>

## Next pages — bond_device_list / 7d7417e776df / 4

- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-a6fba5275b44298b4a6eb8fc267ae614e34be734b147f7450ebcc2773dd148a0)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-a6fba5275b44298b4a6eb8fc267ae614e34be734b147f7450ebcc2773dd148a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8647fe2a96567d1d66206a7e15332db16d9119d2653bee374a50cfefbc54705e"></a>

## bond_device_list.bond_devices — bond_device_list.bond_devices / 3f2430b060bf / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-b46729d24d7c5fedc7541c50c6597214e39effced7311afc2973e351df32216e)
- bond_device_list.bond_devices

<a id="canonical-c22a2ce54071534b6f7a25647c8babe4e6897fa4c581595de99904f3f1b9ee23"></a>

Type: `"object"`. list nested block, Optional.

Bond Devices. List of bond devices.

Upstream description:

List of bond devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingListObjectAttributes("active_backup",
    "lacp")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
bond_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-1fba2d7103da1aec8c6735b07efafd059b27540f8f79ccd50364cc47385267fb"></a>

## Direct properties — bond_device_list.bond_devices / 3f2430b060bf / 3

- [active_backup](resources--securemesh_site--reference--group-001.md#canonical-ec09c2656c6a6419d6102252b10a1ddb7da6856025ccda324c144f991d784aca): complete subsection reference.

<a id="canonical-f08ab8815a1b7860cde80e5561ea3a615119e4d7a0ac5522ea9764c68f3805af"></a>

<a id="canonical-51569596c835a3bb38535886edd32400a23dbf36dd2b1fab0cc7cf7a0e6702d0"></a>

## devices property — bond_device_list.bond_devices / 3f2430b060bf / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site--reference--group-001.md#canonical-9c287f1919ab6bcace83bb98710a2c620ec68d84d244c635e43dfe29d118b175): complete subsection reference.

<a id="canonical-d47c86b565c4b959fd298d19a0a9e96c06da64dc3fddec06476524398a501fe0"></a>

<a id="canonical-a55b5556f41036ad41192450459641109aa74011a7eba1845f1afd6a649305ae"></a>

## link_polling_interval property — bond_device_list.bond_devices / 3f2430b060bf / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-6babfec20727d86a1fd514be09eb25361cd230d12e9438f4193a9ff6c8bd55de"></a>

<a id="canonical-02c8ab492246fc923f7cae39e195f6ace7c9426618496dd6faab26ade9f05474"></a>

## link_up_delay property — bond_device_list.bond_devices / 3f2430b060bf / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-de970fcb77021fce25e2d752aa1f8423f0e857d0940d732bcde2280414a10730"></a>

<a id="canonical-1f243e13be6551211a3d1b7d378c91610995a0db9b0941dd51fad207ac662128"></a>

## name property — bond_device_list.bond_devices / 3f2430b060bf / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-f1d0ea81dbe0416ca4350746538449f24bf4d96a84696d5a2635cbe15d189d61"></a>

## Next pages — bond_device_list.bond_devices / 3f2430b060bf / 8

- [bond_device_list.bond_devices.active_backup](resources--securemesh_site--reference--group-001.md#canonical-ec09c2656c6a6419d6102252b10a1ddb7da6856025ccda324c144f991d784aca)
- [bond_device_list.bond_devices.lacp](resources--securemesh_site--reference--group-001.md#canonical-9c287f1919ab6bcace83bb98710a2c620ec68d84d244c635e43dfe29d118b175)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-b46729d24d7c5fedc7541c50c6597214e39effced7311afc2973e351df32216e)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-ec09c2656c6a6419d6102252b10a1ddb7da6856025ccda324c144f991d784aca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93b41d9c1d0c0981660218c938fcfd30eaf2d67f4fa4f63bdc150a3c1405b777"></a>

## bond_device_list.bond_devices.active_backup — bond_device_list.bond_devices.active_backup / 810e5c7b8ebc / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-b46729d24d7c5fedc7541c50c6597214e39effced7311afc2973e351df32216e)
- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-a6fba5275b44298b4a6eb8fc267ae614e34be734b147f7450ebcc2773dd148a0)
- bond_device_list.bond_devices.active_backup

<a id="canonical-4713ac5fa7a67c2d086d68d3facd6e5bca8e52c80995969770e8e5b2d2d92c3b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-095f7e280c26f710d71b89cf712e421ae54ed20c8f5459284c6b9a06f3a06d3a"></a>

## Direct properties — bond_device_list.bond_devices.active_backup / 810e5c7b8ebc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb238dfd1d54e9a84b018d02f4c01d0ffa973d75028fe5060a69d28a7a93bed3"></a>

## Next pages — bond_device_list.bond_devices.active_backup / 810e5c7b8ebc / 4

- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-a6fba5275b44298b4a6eb8fc267ae614e34be734b147f7450ebcc2773dd148a0)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-9c287f1919ab6bcace83bb98710a2c620ec68d84d244c635e43dfe29d118b175"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edd79fc9c26d146fd224c7ca738a68a3208ab04f076bcdae44d896d51dbaa496"></a>

## bond_device_list.bond_devices.lacp — bond_device_list.bond_devices.lacp / 050feb8c1ea5 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [bond_device_list](resources--securemesh_site--reference--group-001.md#canonical-b46729d24d7c5fedc7541c50c6597214e39effced7311afc2973e351df32216e)
- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-a6fba5275b44298b4a6eb8fc267ae614e34be734b147f7450ebcc2773dd148a0)
- bond_device_list.bond_devices.lacp

<a id="canonical-75d505a004cc4d9d87f41b52ebe41a8d8ab25e26b01cbeee9c5d88f285c1baac"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-f076c51d0c8bd3f351374a7d19ac7bb54733a4b669b510fcfe2c6a2b04aa850e"></a>

## Direct properties — bond_device_list.bond_devices.lacp / 050feb8c1ea5 / 3

<a id="canonical-57d74ad8eb087bddf744c3e57bd40bc90ac77a0e8c41ed0a73147627bc36abba"></a>

<a id="canonical-6c14377cca8d703d34bc5420ed96cd19175a613a301dba7b29c0d742b552edf2"></a>

## rate property — bond_device_list.bond_devices.lacp / 050feb8c1ea5 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-7d42b2ff4b5c6e97b660e62c2298cf6c5510f5a78e1d9f80f750347a4edeb3af"></a>

## Next pages — bond_device_list.bond_devices.lacp / 050feb8c1ea5 / 5

- [bond_device_list.bond_devices](resources--securemesh_site--reference--group-001.md#canonical-a6fba5275b44298b4a6eb8fc267ae614e34be734b147f7450ebcc2773dd148a0)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-6af687786ff749e0698ebec1775e793f9421f4989d96ca808a763d39f55d5db6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc0b707327af98ae2d7131732c3bf6ce4fb272c5d6ad1ffe84f0e1ab6a059ecf"></a>

## coordinates — coordinates / 5ac3281d0bba / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- coordinates

<a id="canonical-ba39b8067027a2263bb8099e5f6c8f1788d3c104e98248617c0869e8bdd8a217"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

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
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0749a474ab67a2be0177508871e5ffe1fd1b56c27170a7ad289392b73951de77"></a>

## Direct properties — coordinates / 5ac3281d0bba / 3

<a id="canonical-de885a4a3a102461976d0c29be695a3fe1c5e8a14d01dcb0434230956b2577bc"></a>

<a id="canonical-732e636dee594d695502607ff1fd3824c48b22532a41b1f1b69cc1351bd53212"></a>

## latitude property — coordinates / 5ac3281d0bba / 4

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-46c5522dd710c5084f2f04eccc3bb301a9a6a33cbbdc583d3016c53bfdd5b287"></a>

<a id="canonical-316c8feb251e451eebdc0924a4671b8a77bcb73ea568977bc8080bcb1cef3975"></a>

## longitude property — coordinates / 5ac3281d0bba / 5

Type: `"number"`. Optional.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-92929c1199060cf1db8da9d1e42bb99493301bef9bf34cb2e58f910e1542d3f5"></a>

## Next pages — coordinates / 5ac3281d0bba / 6

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31c19b284fc3c469a0e4678478295951f7fdb7cad3f2424104cdc94ff479fa7d"></a>

## custom_network_config — custom_network_config / 4e684775955f / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- custom_network_config

<a id="canonical-69d51591c455b22c1466f1156dad836e2ee561f2c8525cb8993582e40d0ee8d4"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_network\_config, default\_network\_config; Default: default\_network\_config\]
SmsNetworkConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("default_config",
    "slo_config"),
  validators.ConflictingObjectAttributes("default_interface_config",
    "interface_list"),
  validators.ConflictingObjectAttributes("default_sli_config",
    "sli_config"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-interface_choice": "[\"default_interface_config\",\"interface_list\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

OneOf alternatives in this subsection:

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-69d51591c455b22c1466f1156dad836e2ee561f2c8525cb8993582e40d0ee8d4)
- [default_network_config](resources--securemesh_site--reference--group-004.md#canonical-871ddd6e9fca923ed67c5ad66b8affcef0f12462f8565e9538776a3e61f8efb5)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_network_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-9f97304718d3e5342c6172dc9ca9ff28d22b13ac4aa0b59055421f3fef1c27bb"></a>

## Direct properties — custom_network_config / 4e684775955f / 3

- [active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-8237ce8646bf2996c7bc9e78304310a279d920a62fc9942c012104b55417b542): complete subsection reference.

- [active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-a1e4739825741e55d421581251a0fc1070c827d805e56d85c46b7eda0dea62e3): complete subsection reference.

- [active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-cb77d0e98c138071590d3ecfa8a11b6af8b8d8d4ee06cae742265db5c3bc54a8): complete subsection reference.

- [default_config](resources--securemesh_site--reference--group-001.md#canonical-437c3eb3a9d4802d9e21ba918bdb66fb1c2e9a77ac2bb8ce21cd2036a068203e): complete subsection reference.

- [default_interface_config](resources--securemesh_site--reference--group-002.md#canonical-ca87af129f201e4638dba90ae63160ba540ab0904854177b9e1ba46419063216): complete subsection reference.

- [default_sli_config](resources--securemesh_site--reference--group-002.md#canonical-a903e780e5aa99f8360fc293efeaf7d706c7cd1f069dc14f1725c4ff2ac868fa): complete subsection reference.

- [forward_proxy_allow_all](resources--securemesh_site--reference--group-002.md#canonical-4e1016aa4c72bf2a3c382dee5b92076bee343c8febdb7c5f47ee0df59823c791): complete subsection reference.

- [global_network_list](resources--securemesh_site--reference--group-002.md#canonical-a21b20825411f77d59a2f8749e4da7123275e583a9a9b117f0b3e30d6e591f23): complete subsection reference.

- [interface_list](resources--securemesh_site--reference--group-002.md#canonical-24a6f9d27c8d5d66e4a9be21064d404e94269d40c199af0b27aef6aa14dbaf9d): complete subsection reference.

- [no_forward_proxy](resources--securemesh_site--reference--group-003.md#canonical-b359dae03c480794d5cfbb230f6e11334ec1aba331ff4bc612f6eda93d129449): complete subsection reference.

- [no_global_network](resources--securemesh_site--reference--group-003.md#canonical-79226d7055b7fdc8feaeb667d3fb6eb645305b63e628b561ec68085e35d7fd1e): complete subsection reference.

- [no_network_policy](resources--securemesh_site--reference--group-003.md#canonical-41dd35d2403e943695d62d3bb63fec401b465ac44bb3b74ba85e83a6d7f9c4e4): complete subsection reference.

- [sli_config](resources--securemesh_site--reference--group-003.md#canonical-ec44962f36fdb869cf0c6c88f594049bcd841bb5bade24706ed7c0ffb838cea2): complete subsection reference.

- [slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6): complete subsection reference.

- [sm_connection_public_ip](resources--securemesh_site--reference--group-004.md#canonical-9d47714be24e01c0fbd23354a4e2f5bc22bedd857689e3612df4118fc0681053): complete subsection reference.

- [sm_connection_pvt_ip](resources--securemesh_site--reference--group-004.md#canonical-a1caa42f5fc7bc4f4b913c717ce9b31da256bb3e598e725d5049b3208030f6d8): complete subsection reference.

<a id="canonical-81f59a0d3e280e597c287a48e016dca5be0abb312dba95641aec7830cf35921f"></a>

<a id="canonical-f4831e0514721623ba125f2ba5559dcde25b490a5a275ee2032b3ceddab8e273"></a>

## tunnel_dead_timeout property — custom_network_config / 4e684775955f / 4

Type: `"number"`. Optional.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Upstream description:

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 180000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  }
}
```

<a id="canonical-ab1515d49d6315e9d223e84b756f59735d1c670f8e6911372e538107d30a3757"></a>

<a id="canonical-efc123bd3473021083030a633d2c6b89336594a2b50b6108c6a5e375a36b015b"></a>

## vip_vrrp_mode property — custom_network_config / 4e684775955f / 5

Type: `"string"`. Optional.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f672226d80ce420bb414e4f51dfee2f0a3ff7b65f2b8ab009f00a1c27ef9cea9"></a>

## Next pages — custom_network_config / 4e684775955f / 6

- [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-8237ce8646bf2996c7bc9e78304310a279d920a62fc9942c012104b55417b542)
- [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-a1e4739825741e55d421581251a0fc1070c827d805e56d85c46b7eda0dea62e3)
- [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-cb77d0e98c138071590d3ecfa8a11b6af8b8d8d4ee06cae742265db5c3bc54a8)
- [custom_network_config.default_config](resources--securemesh_site--reference--group-001.md#canonical-437c3eb3a9d4802d9e21ba918bdb66fb1c2e9a77ac2bb8ce21cd2036a068203e)
- [custom_network_config.default_interface_config](resources--securemesh_site--reference--group-002.md#canonical-ca87af129f201e4638dba90ae63160ba540ab0904854177b9e1ba46419063216)
- [custom_network_config.default_sli_config](resources--securemesh_site--reference--group-002.md#canonical-a903e780e5aa99f8360fc293efeaf7d706c7cd1f069dc14f1725c4ff2ac868fa)
- [custom_network_config.forward_proxy_allow_all](resources--securemesh_site--reference--group-002.md#canonical-4e1016aa4c72bf2a3c382dee5b92076bee343c8febdb7c5f47ee0df59823c791)
- [custom_network_config.global_network_list](resources--securemesh_site--reference--group-002.md#canonical-a21b20825411f77d59a2f8749e4da7123275e583a9a9b117f0b3e30d6e591f23)
- [custom_network_config.interface_list](resources--securemesh_site--reference--group-002.md#canonical-24a6f9d27c8d5d66e4a9be21064d404e94269d40c199af0b27aef6aa14dbaf9d)
- [custom_network_config.no_forward_proxy](resources--securemesh_site--reference--group-003.md#canonical-b359dae03c480794d5cfbb230f6e11334ec1aba331ff4bc612f6eda93d129449)
- [custom_network_config.no_global_network](resources--securemesh_site--reference--group-003.md#canonical-79226d7055b7fdc8feaeb667d3fb6eb645305b63e628b561ec68085e35d7fd1e)
- [custom_network_config.no_network_policy](resources--securemesh_site--reference--group-003.md#canonical-41dd35d2403e943695d62d3bb63fec401b465ac44bb3b74ba85e83a6d7f9c4e4)
- [custom_network_config.sli_config](resources--securemesh_site--reference--group-003.md#canonical-ec44962f36fdb869cf0c6c88f594049bcd841bb5bade24706ed7c0ffb838cea2)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [custom_network_config.sm_connection_public_ip](resources--securemesh_site--reference--group-004.md#canonical-9d47714be24e01c0fbd23354a4e2f5bc22bedd857689e3612df4118fc0681053)
- [custom_network_config.sm_connection_pvt_ip](resources--securemesh_site--reference--group-004.md#canonical-a1caa42f5fc7bc4f4b913c717ce9b31da256bb3e598e725d5049b3208030f6d8)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-8237ce8646bf2996c7bc9e78304310a279d920a62fc9942c012104b55417b542"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef535ed4dab9ba8641bd12cb8a1e41a56f3cd1d876fce2a797419991e8a63187"></a>

## custom_network_config.active_enhanced_firewall_policies — custom_network_config.active_enhanced_firewall_policies / e4db8149824e / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- custom_network_config.active_enhanced_firewall_policies

<a id="canonical-093a54889d320ab0882b5cc8a346fda6d47641697be16a39e5729187ac995358"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf12976abaa42315ef72364cd08c04f26d40927825e244b877053d0957e67ddc"></a>

## Direct properties — custom_network_config.active_enhanced_firewall_policies / e4db8149824e / 3

- [enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-1b583e935e74e0232f0a731fc3acc5410bc8483d6e5c9884eed915ffddf6aad8): complete subsection reference.

<a id="canonical-37397ccba81d41d802753abed324e112363299a31147e51e2a896711344866e1"></a>

## Next pages — custom_network_config.active_enhanced_firewall_policies / e4db8149824e / 4

- [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-1b583e935e74e0232f0a731fc3acc5410bc8483d6e5c9884eed915ffddf6aad8)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-1b583e935e74e0232f0a731fc3acc5410bc8483d6e5c9884eed915ffddf6aad8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01aa3260b7ab726ca699ed8fcbe37089d4d344ced1b5e85dfd8df300f8d02cce"></a>

## custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / fe923b92736a / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-8237ce8646bf2996c7bc9e78304310a279d920a62fc9942c012104b55417b542)
- custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-01dd221a0babe26a838be505d5f4cafcfc9e052ac382f26029d1d3967c72ae5b"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-c472c3e72a1d5d8b2fa3fa6f62bd90e013f63d5b80ba51458ad2191cc6b80477"></a>

## Direct properties — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / fe923b92736a / 3

<a id="canonical-fc37e8809730db5103256bd9bd5ef9cb060438ebab695bcd13ee07e3ab0d14d2"></a>

<a id="canonical-2a2d3ac7210b629c33ead14994974be540917c2ab1cd60320a3c9a4a8bfd1e15"></a>

## name property — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / fe923b92736a / 4

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

<a id="canonical-46a99bb573cc4f83c97e2f5ca9c453ef96a30a6158e21b56bf3dbc2e90a95685"></a>

<a id="canonical-b9c65a6b822b0b2468f76d1e76e6d3afda35505e0f5ea51cfddb521e0b4a5105"></a>

## namespace property — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / fe923b92736a / 5

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

<a id="canonical-8c80be15b1205a6854ba69535cfd204cd7374e6b74e5a03b91b91f289730ae68"></a>

<a id="canonical-34e00768952a427b32e6001159707479d8b4b095dac47343748f02a956bbc66a"></a>

## tenant property — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / fe923b92736a / 6

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

<a id="canonical-9c5e9f24de6f1ec66797bdb531853bcdf48d78579579d380a818761749c592ba"></a>

## Next pages — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / fe923b92736a / 7

- [custom_network_config.active_enhanced_firewall_policies](resources--securemesh_site--reference--group-001.md#canonical-8237ce8646bf2996c7bc9e78304310a279d920a62fc9942c012104b55417b542)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-a1e4739825741e55d421581251a0fc1070c827d805e56d85c46b7eda0dea62e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d753a2ee2f88eb66e6af4d47fe6eebc1c583305b545e3a7ccc9a7d15ae474fd9"></a>

## custom_network_config.active_forward_proxy_policies — custom_network_config.active_forward_proxy_policies / e0682e6a704a / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- custom_network_config.active_forward_proxy_policies

<a id="canonical-8395bcd613c346ca6e727b1036d20f900c3fe1a0a21c4c853b8540a2110086d3"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-6fea4d50431d447f04369acc142b248adfe83590d069b0488d66166772a94fac"></a>

## Direct properties — custom_network_config.active_forward_proxy_policies / e0682e6a704a / 3

- [forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-0d36a8b01de4cb6e3158baf26f1a0e8216564444b583d51901454e35e9006be2): complete subsection reference.

<a id="canonical-097b963389997e8da043f5c18429249fe689ac602f12c8f5a9b3501757e08620"></a>

## Next pages — custom_network_config.active_forward_proxy_policies / e0682e6a704a / 4

- [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-0d36a8b01de4cb6e3158baf26f1a0e8216564444b583d51901454e35e9006be2)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-0d36a8b01de4cb6e3158baf26f1a0e8216564444b583d51901454e35e9006be2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1febc4b6dbf3bd58300b91ef356769de8ff24cb228ade4b8a99b0c6cb1f11669"></a>

## custom_network_config.active_forward_proxy_policies.forward_proxy_policies — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / c7ed05521dc7 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-a1e4739825741e55d421581251a0fc1070c827d805e56d85c46b7eda0dea62e3)
- custom_network_config.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-d76add8c6125b66eb21fc142b75a264d2d2cae5550f3c71374c0e455f43a7080"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-74289c663272a9e02d5838cfcc5b3c2493a7514df6afe6026ebb585e3f579521"></a>

## Direct properties — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / c7ed05521dc7 / 3

<a id="canonical-2c3121db6299ea1169736abd087b6816b55415a03a0e868f72cb2a5480d9a53b"></a>

<a id="canonical-37ef08792d7613b87a8aea89a1282fcb80b40598fb2d8ff2dc22f642720ce58e"></a>

## name property — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / c7ed05521dc7 / 4

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

<a id="canonical-ad2340b12d69a9142e333e63ab420c0b4648335a6d05a0a2a948394e74547102"></a>

<a id="canonical-c6550ec6cd6b4a355aacf851a50e141a0a6358497f5b22ccfad37c978a998484"></a>

## namespace property — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / c7ed05521dc7 / 5

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

<a id="canonical-b36e6a6fc6a2def01de7f77904af65dc0c9318eea98d331183907c66594d39dc"></a>

<a id="canonical-551af956b71fdf225ced0f33b42094c7eeb6446fbd2547803fd5c5b1e3317234"></a>

## tenant property — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / c7ed05521dc7 / 6

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

<a id="canonical-3a31ecb0e80c6d347cbcaf79455015e52edf6d16c1b5f2925d5d8c93528b0efa"></a>

## Next pages — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / c7ed05521dc7 / 7

- [custom_network_config.active_forward_proxy_policies](resources--securemesh_site--reference--group-001.md#canonical-a1e4739825741e55d421581251a0fc1070c827d805e56d85c46b7eda0dea62e3)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-cb77d0e98c138071590d3ecfa8a11b6af8b8d8d4ee06cae742265db5c3bc54a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38a94db327bf23450fdab44b10f50814760ed643b1f6f3c3e593264f3950de49"></a>

## custom_network_config.active_network_policies — custom_network_config.active_network_policies / a07b97bc0a44 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- custom_network_config.active_network_policies

<a id="canonical-97df10bc598d5cc71fdf2256142e970bad8adb24999bda5ab8e5ec952ccb9b16"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-e88d7e59020815b600d13c505f36fcaed4bf77918bd509150b9101a0c1e75ede"></a>

## Direct properties — custom_network_config.active_network_policies / a07b97bc0a44 / 3

- [network_policies](resources--securemesh_site--reference--group-001.md#canonical-a19bdac0903cd288ad9833f6b79cec810bdde9bfb4b65ddc6accc4a9c8797932): complete subsection reference.

<a id="canonical-24d75daabda1e5a73a4081f8c0ff98bdf8600ea95b01187f3204aa6b6a19059c"></a>

## Next pages — custom_network_config.active_network_policies / a07b97bc0a44 / 4

- [custom_network_config.active_network_policies.network_policies](resources--securemesh_site--reference--group-001.md#canonical-a19bdac0903cd288ad9833f6b79cec810bdde9bfb4b65ddc6accc4a9c8797932)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-a19bdac0903cd288ad9833f6b79cec810bdde9bfb4b65ddc6accc4a9c8797932"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c37b9ae30114f1f87d71b110345b29a2ec1deab3a33b15d52dafd6f2183cf880"></a>

## custom_network_config.active_network_policies.network_policies — custom_network_config.active_network_policies.network_policies / fe5e281b3bf6 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-cb77d0e98c138071590d3ecfa8a11b6af8b8d8d4ee06cae742265db5c3bc54a8)
- custom_network_config.active_network_policies.network_policies

<a id="canonical-4631c3d357e69c16c1dbb6dffd361dbf0293e08d9d23fa2bf9063e434fc83786"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-7f8abc282a4713be3a86b71a577a1fd0c0425ec45cc12724d7c11134c97e21c1"></a>

## Direct properties — custom_network_config.active_network_policies.network_policies / fe5e281b3bf6 / 3

<a id="canonical-c90fc1757fce6055a6d4f90e3a00a218a4f7bd1e8ab34561cc091356dfb4a7a5"></a>

<a id="canonical-b9664f4854269249db0de142b51b6206bcce0f503669ba398328160fd598792f"></a>

## name property — custom_network_config.active_network_policies.network_policies / fe5e281b3bf6 / 4

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

<a id="canonical-3387434dbb481a75c2329351d4a6e118c5cdeaa08d2138a7c0bf7b37f84bbd2d"></a>

<a id="canonical-cc3889ac31f7b0e27fbac2938437e35b13ccd44e497748d575820edfa7253f7b"></a>

## namespace property — custom_network_config.active_network_policies.network_policies / fe5e281b3bf6 / 5

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

<a id="canonical-844bcc9a9fe481ff0bd77c5ddf1c62b4b7023b44ab49939fd99d1097dbeb35c5"></a>

<a id="canonical-e9b02403b5a6b070c22f149d1a275262632a381d70bd988fe50e9fda360cd45b"></a>

## tenant property — custom_network_config.active_network_policies.network_policies / fe5e281b3bf6 / 6

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

<a id="canonical-e21b7fbda94a506b567b66a1518a7bb697bb08fa0a29a7609c8ceac297942b29"></a>

## Next pages — custom_network_config.active_network_policies.network_policies / fe5e281b3bf6 / 7

- [custom_network_config.active_network_policies](resources--securemesh_site--reference--group-001.md#canonical-cb77d0e98c138071590d3ecfa8a11b6af8b8d8d4ee06cae742265db5c3bc54a8)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-437c3eb3a9d4802d9e21ba918bdb66fb1c2e9a77ac2bb8ce21cd2036a068203e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
