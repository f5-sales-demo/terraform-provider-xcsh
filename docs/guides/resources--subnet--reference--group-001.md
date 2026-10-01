---
page_title: "xcsh_subnet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet reference."
---

# xcsh_subnet reference

<a id="canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72f4ff7552bf0bc40be9b0229a85648b0f76098b2f377d5e3227415fd6f53728"></a>

## Property reference — Property reference / f3e6cad424a5 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- Property reference

<a id="canonical-dec4d19eed3a6087069b647d3dd9381807ce7c0a9ad21259c6ab53dcf3352e1c"></a>

## Direct properties — Property reference / f3e6cad424a5 / 3

<a id="canonical-f93e41dd5f3480f19839b3a2e5fdba8d730a1da01104c3cda579ee29f1597f78"></a>

<a id="canonical-1cf061409ab40c3667ec9a4285ce4bde27cd54741622eb466433c42dcce5cca4"></a>

## annotations property — Property reference / f3e6cad424a5 / 4

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

- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-e5e31f9e35a62cfe043ad9308079cf11738a3a9d6defc94439bc2dcc488afa6a): complete subsection reference.

- [connect_to_slo](resources--subnet--reference--group-001.md#canonical-a9582cfb760ca8ec54b1a5617289ee2b5efc90825f7dea6d28d3751d0dae5df0): complete subsection reference.

<a id="canonical-2a12cb8b8485e8e8364740edc2438500b7f0d7fdec458551ab48056d065f098a"></a>

<a id="canonical-e9d08cef453e274c7b12b5152739b867cb789d53e0ebffc2197ef6e532f4a34f"></a>

## description property — Property reference / f3e6cad424a5 / 5

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

<a id="canonical-bf77f4a4d90629a7fdbad46b193126195631b39aae6ad948c0f1d2d9e57dd6e1"></a>

<a id="canonical-27e2a6d578d04754c2f969e721c714f1b316cf1c28af5947ee80db56b344856f"></a>

## disable property — Property reference / f3e6cad424a5 / 6

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

<a id="canonical-2dbe7a7c7b9bb182e5c19e5ad4596b3b6a5f353f14c93870c0b304b83c6cc2d0"></a>

<a id="canonical-f38a27c260d309c290b0fee3abb3e66848791399ab39c5e92e083c583637bd44"></a>

## id property — Property reference / f3e6cad424a5 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated_nw](resources--subnet--reference--group-001.md#canonical-1371e14bd89a726283dc0310ab7e6a9c256dc0373013492eb7698da169d10d70): complete subsection reference.

<a id="canonical-5a91c640d94a94d61192ce9468a9ad94b1f4c33dfca8da4f607944222ddf46ca"></a>

<a id="canonical-6f060b10372e74132b9baa0d0db14ec79ac3854baf9f022b7328dd0956061f16"></a>

## labels property — Property reference / f3e6cad424a5 / 8

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

<a id="canonical-9e2816357b8b407206d2ba71da8e77f9e467578b5ff46cb66c173fd4a7f9ff53"></a>

<a id="canonical-a71a9f43140c9034a55afbf89ebccce5b32252ef77a9b12c9af13c93f2b220a5"></a>

## name property — Property reference / f3e6cad424a5 / 9

Type: `"string"`. Required.

Name of the Subnet. Must be unique within the namespace.

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

<a id="canonical-26aabe97760dfcd73c3191e56244ccc2608d50798fa8b0aa3a0a1d5db8fa4693"></a>

<a id="canonical-d27c1e1d25ca8f34cf5524d368374206a53872882f6a9f05d44c201800cca745"></a>

## namespace property — Property reference / f3e6cad424a5 / 10

Type: `"string"`. Required.

Namespace where the Subnet is created.

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

- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07): complete subsection reference.

- [timeouts](resources--subnet--reference--group-001.md#canonical-b8a99f5e7c9237c1dee4eec2a987d37e971d8d352d728faf67c4fc3afe6f9985): complete subsection reference.

<a id="canonical-60723556ea12d501f0d8aeb2fbe044a22c920c6209bb2b70cb75156689559b86"></a>

## All schema paths — Property reference / f3e6cad424a5 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--subnet--reference--group-001.md#canonical-f93e41dd5f3480f19839b3a2e5fdba8d730a1da01104c3cda579ee29f1597f78) |
| `connect_to_layer2` | [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-177fb2b491bf84ff25227f48fe21e89e847f83338124e580698c145e5e746c54) |
| `connect_to_layer2.layer2_intf_ref` | [connect_to_layer2.layer2_intf_ref](resources--subnet--reference--group-001.md#canonical-d001432831972f359fac78681500ad4f728e21653032cb1855f361cbaa97a8c4) |
| `connect_to_layer2.layer2_intf_ref.name` | [connect_to_layer2.layer2_intf_ref.name](resources--subnet--reference--group-001.md#canonical-e239b919e8f3d0934bde88519ce1b80a97ba47b6d03ebdb18a114d954cf854b2) |
| `connect_to_layer2.layer2_intf_ref.namespace` | [connect_to_layer2.layer2_intf_ref.namespace](resources--subnet--reference--group-001.md#canonical-618dfe486abaebda3bd1acb2eac27af5294e35dfc3ea93f225d109d52af456ce) |
| `connect_to_layer2.layer2_intf_ref.tenant` | [connect_to_layer2.layer2_intf_ref.tenant](resources--subnet--reference--group-001.md#canonical-38495da98676f8f5cd22e7d77b8f600af2db95fc06178d639a4a04831af300e4) |
| `connect_to_slo` | [connect_to_slo](resources--subnet--reference--group-001.md#canonical-75878c7f3c8de8c76167228de54aa55b7f15b319e240065c603f350b555bc680) |
| `description` | [description](resources--subnet--reference--group-001.md#canonical-2a12cb8b8485e8e8364740edc2438500b7f0d7fdec458551ab48056d065f098a) |
| `disable` | [disable](resources--subnet--reference--group-001.md#canonical-bf77f4a4d90629a7fdbad46b193126195631b39aae6ad948c0f1d2d9e57dd6e1) |
| `id` | [id](resources--subnet--reference--group-001.md#canonical-2dbe7a7c7b9bb182e5c19e5ad4596b3b6a5f353f14c93870c0b304b83c6cc2d0) |
| `isolated_nw` | [isolated_nw](resources--subnet--reference--group-001.md#canonical-2de2f5710b3427fd948deadcef910268001112fb29403a2e91d54052bb19edc2) |
| `labels` | [labels](resources--subnet--reference--group-001.md#canonical-5a91c640d94a94d61192ce9468a9ad94b1f4c33dfca8da4f607944222ddf46ca) |
| `name` | [name](resources--subnet--reference--group-001.md#canonical-9e2816357b8b407206d2ba71da8e77f9e467578b5ff46cb66c173fd4a7f9ff53) |
| `namespace` | [namespace](resources--subnet--reference--group-001.md#canonical-26aabe97760dfcd73c3191e56244ccc2608d50798fa8b0aa3a0a1d5db8fa4693) |
| `site_subnet_params` | [site_subnet_params](resources--subnet--reference--group-001.md#canonical-2651de736895205154cbe6ecf6199654b5927cc7a0e2819283d449de50f256b3) |
| `site_subnet_params.dhcp` | [site_subnet_params.dhcp](resources--subnet--reference--group-001.md#canonical-102166a5738ce4858544f3a4ca668b74cc6035f4aaa62e41ade97cadbbf0272e) |
| `site_subnet_params.site` | [site_subnet_params.site](resources--subnet--reference--group-001.md#canonical-43e465669b32d157341ad8218ba2a46e9ccf90a8ce67ec1b8950889ac422fa60) |
| `site_subnet_params.site.name` | [site_subnet_params.site.name](resources--subnet--reference--group-001.md#canonical-eff90ccc1dae1bf9a390f7d0e977a0a2dd974f8298f11b426dbcf5014890c77c) |
| `site_subnet_params.site.namespace` | [site_subnet_params.site.namespace](resources--subnet--reference--group-001.md#canonical-fcd95523beb3f1c525def23a634f718519f55f30c6c3bfaafd71ccde9f9e3d4a) |
| `site_subnet_params.site.tenant` | [site_subnet_params.site.tenant](resources--subnet--reference--group-001.md#canonical-dac2ecc595e5251f6d852f10e60a7062da372a3b2f5ab8776c4af4549d7f706a) |
| `site_subnet_params.static_ip` | [site_subnet_params.static_ip](resources--subnet--reference--group-001.md#canonical-fbd720a80b4692574f7e262ca32f96a9ae49b8cd3f676ce9aa8dd3db53b91f78) |
| `site_subnet_params.subnet_dhcp_server_params` | [site_subnet_params.subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-d67171df47c63a740dc23920e5450f5c25573f8e9770c0529971eff4a66b1fd6) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](resources--subnet--reference--group-001.md#canonical-afd521966e8a733327a53b67f50b2e90b93e062170aba84d0332e71542a1a1af) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix](resources--subnet--reference--group-001.md#canonical-44b3e1750ad9251ea0e54d1bda3163edf712d419009f2190c3e94cb866a424ec) |
| `timeouts` | [timeouts](resources--subnet--reference--group-001.md#canonical-c7842296a76b7d246642b3595ec11ce87cf8f87d74bc1e0211cf5d2ae931ffe9) |
| `timeouts.create` | [timeouts.create](resources--subnet--reference--group-001.md#canonical-3cedf4e4d53c5bc4037899a8598d57729ed112495222370f683f3c5be3974dc8) |
| `timeouts.delete` | [timeouts.delete](resources--subnet--reference--group-001.md#canonical-27217832a5154be15fcc09387fa94d2b78b44f4ce86256de6d78d95a3c58edf6) |
| `timeouts.read` | [timeouts.read](resources--subnet--reference--group-001.md#canonical-7b86f7671bedcec89bccdd98faec98581a528b0aa6906a15be19ddb2a433ea99) |
| `timeouts.update` | [timeouts.update](resources--subnet--reference--group-001.md#canonical-6059d26ce8b789a5b19d1e72e2a0c20366c48654687517cb7033b8fddb6352a4) |

<a id="canonical-62741929d816f5a63462243874941cb91ccbb410b90196db6090c123100d2781"></a>

## Next pages — Property reference / f3e6cad424a5 / 12

- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-e5e31f9e35a62cfe043ad9308079cf11738a3a9d6defc94439bc2dcc488afa6a)
- [connect_to_slo](resources--subnet--reference--group-001.md#canonical-a9582cfb760ca8ec54b1a5617289ee2b5efc90825f7dea6d28d3751d0dae5df0)
- [isolated_nw](resources--subnet--reference--group-001.md#canonical-1371e14bd89a726283dc0310ab7e6a9c256dc0373013492eb7698da169d10d70)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- [timeouts](resources--subnet--reference--group-001.md#canonical-b8a99f5e7c9237c1dee4eec2a987d37e971d8d352d728faf67c4fc3afe6f9985)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-e5e31f9e35a62cfe043ad9308079cf11738a3a9d6defc94439bc2dcc488afa6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4e8182ddcf047150b1d2fbdd373449c1b996cba4bbfa9b864aae64dd2daeb06"></a>

## connect_to_layer2 — connect_to_layer2 / f9d9713e721f / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- connect_to_layer2

<a id="canonical-177fb2b491bf84ff25227f48fe21e89e847f83338124e580698c145e5e746c54"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: connect\_to\_layer2, connect\_to\_slo, isolated\_nw\] Configuration parameter for connect
to layer2.

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

- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-177fb2b491bf84ff25227f48fe21e89e847f83338124e580698c145e5e746c54)
- [connect_to_slo](resources--subnet--reference--group-001.md#canonical-75878c7f3c8de8c76167228de54aa55b7f15b319e240065c603f350b555bc680)
- [isolated_nw](resources--subnet--reference--group-001.md#canonical-2de2f5710b3427fd948deadcef910268001112fb29403a2e91d54052bb19edc2)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
connect_to_layer2 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2a250d7032d0982a88adf86034aadb2485d1b61d04fd4935565ba7617f2ad69c"></a>

## Direct properties — connect_to_layer2 / f9d9713e721f / 3

- [layer2_intf_ref](resources--subnet--reference--group-001.md#canonical-19850f44889484caa6308670edabec07f246ee08f653ed30e975fa6496d77230): complete subsection reference.

<a id="canonical-0a14ca5f2a222a7cd8692267d69ef69c781b569a8fd9bcddaac31a80a548a1b4"></a>

## Next pages — connect_to_layer2 / f9d9713e721f / 4

- [connect_to_layer2.layer2_intf_ref](resources--subnet--reference--group-001.md#canonical-19850f44889484caa6308670edabec07f246ee08f653ed30e975fa6496d77230)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-19850f44889484caa6308670edabec07f246ee08f653ed30e975fa6496d77230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baadeec0cb833628b835a89ee64b751d2d4169eca115bd1563c99ee7d7c4f80f"></a>

## connect_to_layer2.layer2_intf_ref — connect_to_layer2.layer2_intf_ref / 2ac7242b1691 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-e5e31f9e35a62cfe043ad9308079cf11738a3a9d6defc94439bc2dcc488afa6a)
- connect_to_layer2.layer2_intf_ref

<a id="canonical-d001432831972f359fac78681500ad4f728e21653032cb1855f361cbaa97a8c4"></a>

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
layer2_intf_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-25c134e297c298cbe247f481457a5cec748d899feda05d6239f265faa2c62eb0"></a>

## Direct properties — connect_to_layer2.layer2_intf_ref / 2ac7242b1691 / 3

<a id="canonical-e239b919e8f3d0934bde88519ce1b80a97ba47b6d03ebdb18a114d954cf854b2"></a>

<a id="canonical-cbd96c8b787df697e2ce5bc19cf581fca48e8a3eea976ce79758bd6c7e0bd0c0"></a>

## name property — connect_to_layer2.layer2_intf_ref / 2ac7242b1691 / 4

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

<a id="canonical-618dfe486abaebda3bd1acb2eac27af5294e35dfc3ea93f225d109d52af456ce"></a>

<a id="canonical-def326b02eb5f48862b1d06a61230397ed7c974c582d15072109d23cf7814ca1"></a>

## namespace property — connect_to_layer2.layer2_intf_ref / 2ac7242b1691 / 5

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

<a id="canonical-38495da98676f8f5cd22e7d77b8f600af2db95fc06178d639a4a04831af300e4"></a>

<a id="canonical-5c4a9266b1d4ddbd307c6724ed3c67597d64416ae0f83cb82b1f46cd7494a641"></a>

## tenant property — connect_to_layer2.layer2_intf_ref / 2ac7242b1691 / 6

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

<a id="canonical-d3701612d2eecc2a5feb2b0c9b1aa275b5f12de1695ea055f954aee90096a5fc"></a>

## Next pages — connect_to_layer2.layer2_intf_ref / 2ac7242b1691 / 7

- [connect_to_layer2](resources--subnet--reference--group-001.md#canonical-e5e31f9e35a62cfe043ad9308079cf11738a3a9d6defc94439bc2dcc488afa6a)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-a9582cfb760ca8ec54b1a5617289ee2b5efc90825f7dea6d28d3751d0dae5df0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a6eb95430601fd453c0ff60075b06723c48d5633b4ea4ad9d0765cb60281e68"></a>

## connect_to_slo — connect_to_slo / c66de15635df / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- connect_to_slo

<a id="canonical-75878c7f3c8de8c76167228de54aa55b7f15b319e240065c603f350b555bc680"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for connect to slo.

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
connect_to_slo = {}
```

<a id="canonical-daf9d9d8a971bc6b767cec67fd70ad1df5ce0c2334adad7da40b5f538e2eccdf"></a>

## Direct properties — connect_to_slo / c66de15635df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7eaa563ca69e2064255455e7bad17ae31284c4508770fd3f42be287ed4cdd9b4"></a>

## Next pages — connect_to_slo / c66de15635df / 4

- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-1371e14bd89a726283dc0310ab7e6a9c256dc0373013492eb7698da169d10d70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5d643b30d306aa6acf40254b4add2ac1490f207512de8dc03bbc5a6275f0ad8"></a>

## isolated_nw — isolated_nw / 24adf093b175 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- isolated_nw

<a id="canonical-2de2f5710b3427fd948deadcef910268001112fb29403a2e91d54052bb19edc2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for isolated nw.

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
isolated_nw = {}
```

<a id="canonical-4979e19c5cd3d328ac274b6715e474022cd0e0b6dbf332cd9ce0e4bbe6c89d0b"></a>

## Direct properties — isolated_nw / 24adf093b175 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f68b0b33e9b1a369d2d5442bf7a56f1a5e4fc1839f6cd374faae66839f55c08f"></a>

## Next pages — isolated_nw / 24adf093b175 / 4

- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e29a4c76ef4b7118224fe5328d8cfc0cc8b70b637332bd7c183b3c5e8bcc8a1c"></a>

## site_subnet_params — site_subnet_params / 036df84074a4 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- site_subnet_params

<a id="canonical-2651de736895205154cbe6ecf6199654b5927cc7a0e2819283d449de50f256b3"></a>

Type: `"object"`. list nested block, Optional.

Site Subnet Parameters. Configure subnet parameters per site.

Upstream description:

Configure subnet parameters per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dhcp",
    "static_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
site_subnet_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-123fcfad4cde986a6366b298b950d934ba5f95e36808129ee5b314422617760a"></a>

## Direct properties — site_subnet_params / 036df84074a4 / 3

- [dhcp](resources--subnet--reference--group-001.md#canonical-f3bc19ba66a279be70c80922fe5344738e836f5b80f60cde13148dc26b408596): complete subsection reference.

- [site](resources--subnet--reference--group-001.md#canonical-d22109e70234e89a219e7a15a2cba985ce8031fd67371c29d430cb7028f31247): complete subsection reference.

- [static_ip](resources--subnet--reference--group-001.md#canonical-0063064d8d06464dbb2aea5807d07bc0ea6d54761a27d188d6fa569b466f99e0): complete subsection reference.

- [subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-ea423b760296ec89e9c123afdff76e84136cc137a2201050bc96f2cfc3a776a3): complete subsection reference.

<a id="canonical-9ded17abf80c0a123bdef78d05dfc4476fc121f4ef200b9fd406b09965d88bc6"></a>

## Next pages — site_subnet_params / 036df84074a4 / 4

- [site_subnet_params.dhcp](resources--subnet--reference--group-001.md#canonical-f3bc19ba66a279be70c80922fe5344738e836f5b80f60cde13148dc26b408596)
- [site_subnet_params.site](resources--subnet--reference--group-001.md#canonical-d22109e70234e89a219e7a15a2cba985ce8031fd67371c29d430cb7028f31247)
- [site_subnet_params.static_ip](resources--subnet--reference--group-001.md#canonical-0063064d8d06464dbb2aea5807d07bc0ea6d54761a27d188d6fa569b466f99e0)
- [site_subnet_params.subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-ea423b760296ec89e9c123afdff76e84136cc137a2201050bc96f2cfc3a776a3)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-f3bc19ba66a279be70c80922fe5344738e836f5b80f60cde13148dc26b408596"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9df28feea0aaf05d8163cfc8c766870fbea06946c16a6b4913aad0c94d5184c"></a>

## site_subnet_params.dhcp — site_subnet_params.dhcp / fd16e85df51a / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- site_subnet_params.dhcp

<a id="canonical-102166a5738ce4858544f3a4ca668b74cc6035f4aaa62e41ade97cadbbf0272e"></a>

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
dhcp = {}
```

<a id="canonical-16bf20e4ddffa64e45f5ad74c05d4012e8aa95afb916f6662abc4698a7c3ca92"></a>

## Direct properties — site_subnet_params.dhcp / fd16e85df51a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0d9131b70154935f6580e450be71bfcd943740cd4ef9c3e20bd9e6fac15b7537"></a>

## Next pages — site_subnet_params.dhcp / fd16e85df51a / 4

- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-d22109e70234e89a219e7a15a2cba985ce8031fd67371c29d430cb7028f31247"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0aa620f9737372acc2f94b9fc4543f28c671185593bdc870fd0355b26568859"></a>

## site_subnet_params.site — site_subnet_params.site / a86454f076f3 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- site_subnet_params.site

<a id="canonical-43e465669b32d157341ad8218ba2a46e9ccf90a8ce67ec1b8950889ac422fa60"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-c158aed9d914e2d24612884d963139246f2ac07590fa1da85068235f7f35aae0"></a>

## Direct properties — site_subnet_params.site / a86454f076f3 / 3

<a id="canonical-eff90ccc1dae1bf9a390f7d0e977a0a2dd974f8298f11b426dbcf5014890c77c"></a>

<a id="canonical-31faf23f6fdc0641d84aad3206f26bed649dc7e772f3a020edd08c06ef7b3e66"></a>

## name property — site_subnet_params.site / a86454f076f3 / 4

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

<a id="canonical-fcd95523beb3f1c525def23a634f718519f55f30c6c3bfaafd71ccde9f9e3d4a"></a>

<a id="canonical-1948e4afe2b515ea9b52483e7a79adeb5c6517fb6af282959c395b537ea99704"></a>

## namespace property — site_subnet_params.site / a86454f076f3 / 5

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

<a id="canonical-dac2ecc595e5251f6d852f10e60a7062da372a3b2f5ab8776c4af4549d7f706a"></a>

<a id="canonical-829cb2803bead0708d757540331355567c6d359084c3fa1e140f0c4c0a057ee3"></a>

## tenant property — site_subnet_params.site / a86454f076f3 / 6

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

<a id="canonical-8e6cfb680caddd9fb7bed262c0c6fc863384cc060c9633af08b4d0823b68a0dd"></a>

## Next pages — site_subnet_params.site / a86454f076f3 / 7

- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-0063064d8d06464dbb2aea5807d07bc0ea6d54761a27d188d6fa569b466f99e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17c34e499e8c3be0c86abd722b85e23328b1708fa16c03fc20674ed6ab871cd4"></a>

## site_subnet_params.static_ip — site_subnet_params.static_ip / 4cf2623c177f / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- site_subnet_params.static_ip

<a id="canonical-fbd720a80b4692574f7e262ca32f96a9ae49b8cd3f676ce9aa8dd3db53b91f78"></a>

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
static_ip = {}
```

<a id="canonical-f18f8c6cbd9d44e9ea6f7f3e2862de456d8d1f7144ce61ca7adddb838634b9a6"></a>

## Direct properties — site_subnet_params.static_ip / 4cf2623c177f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc82a45abc9316ee980bbb1a8f031d7b47cf61c571a737e3b13109c1caf28e8b"></a>

## Next pages — site_subnet_params.static_ip / 4cf2623c177f / 4

- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-ea423b760296ec89e9c123afdff76e84136cc137a2201050bc96f2cfc3a776a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7dde8e06f745e26373d56ddf653d2af36fb961f224d3d30831a1c0f4d778f27"></a>

## site_subnet_params.subnet_dhcp_server_params — site_subnet_params.subnet_dhcp_server_params / f004c8b7d086 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- site_subnet_params.subnet_dhcp_server_params

<a id="canonical-d67171df47c63a740dc23920e5450f5c25573f8e9770c0529971eff4a66b1fd6"></a>

Type: `"object"`. single nested block, Optional.

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

Upstream description:

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

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
subnet_dhcp_server_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2cbbfbde6c92a67be872d36a20e31d4943fd25c13bff28e915b0baea86c13e95"></a>

## Direct properties — site_subnet_params.subnet_dhcp_server_params / f004c8b7d086 / 3

- [dhcp_networks](resources--subnet--reference--group-001.md#canonical-66eb4058179dbeead82491d262b598e029f11099fcaefe8d247eb9a4d03c57ed): complete subsection reference.

<a id="canonical-41484cb039f55244ff173f384fb6af57797a4eb8e91ccff74d387e86fbcb0c40"></a>

## Next pages — site_subnet_params.subnet_dhcp_server_params / f004c8b7d086 / 4

- [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](resources--subnet--reference--group-001.md#canonical-66eb4058179dbeead82491d262b598e029f11099fcaefe8d247eb9a4d03c57ed)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-66eb4058179dbeead82491d262b598e029f11099fcaefe8d247eb9a4d03c57ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d9c1c2b1ecd072b28c9c2ad7c33cf420242a1ffc8be06c7496a87e67d6f0a23"></a>

## site_subnet_params.subnet_dhcp_server_params.dhcp_networks — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / a19478f73615 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [site_subnet_params](resources--subnet--reference--group-001.md#canonical-3509922f191306f8fed1a1d4413b3945d8c1d8ab5ab3c4445ffd32053f182f07)
- [site_subnet_params.subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-ea423b760296ec89e9c123afdff76e84136cc137a2201050bc96f2cfc3a776a3)
- site_subnet_params.subnet_dhcp_server_params.dhcp_networks

<a id="canonical-afd521966e8a733327a53b67f50b2e90b93e062170aba84d0332e71542a1a1af"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-fb2338213d2e665128fd0470fcd26ba9ff2c4f5dc73ea83d33ac31a0734b26fb"></a>

## Direct properties — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / a19478f73615 / 3

<a id="canonical-44b3e1750ad9251ea0e54d1bda3163edf712d419009f2190c3e94cb866a424ec"></a>

<a id="canonical-725162802d89341d5e0063d101926dfc8c1959ee4523493ed90b26185eed6c80"></a>

## network_prefix property — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / a19478f73615 / 4

Type: `"string"`. Optional.

Exclusive with \[\] Network prefix for subnet.

Upstream description:

Exclusive with \[\] Network prefix for subnet.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-9c181ac40bebf5532e4beba76c3806a614010c3b072c7a91bd0d90072f555b2a"></a>

## Next pages — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / a19478f73615 / 5

- [site_subnet_params.subnet_dhcp_server_params](resources--subnet--reference--group-001.md#canonical-ea423b760296ec89e9c123afdff76e84136cc137a2201050bc96f2cfc3a776a3)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-b8a99f5e7c9237c1dee4eec2a987d37e971d8d352d728faf67c4fc3afe6f9985"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ae704981aaa763ba30fafe58c3adc209b33c76a372abc63448f31a45258bb27"></a>

## timeouts — timeouts / 1c99646801be / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- timeouts

<a id="canonical-c7842296a76b7d246642b3595ec11ce87cf8f87d74bc1e0211cf5d2ae931ffe9"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-9e6409e1677589580d875ceee6e8623a11905b28fc51c1f02f6bd968c66624d0"></a>

## Direct properties — timeouts / 1c99646801be / 3

<a id="canonical-3cedf4e4d53c5bc4037899a8598d57729ed112495222370f683f3c5be3974dc8"></a>

<a id="canonical-30175bbee95b792ec775cea09dea987884bfaf59e09fcb6d3ce53dd16b857b76"></a>

## create property — timeouts / 1c99646801be / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-27217832a5154be15fcc09387fa94d2b78b44f4ce86256de6d78d95a3c58edf6"></a>

<a id="canonical-8db87e4fba67f767749d733836e915cabfecfd2dccf3a097e10ee1bb6bd0dd2f"></a>

## delete property — timeouts / 1c99646801be / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-7b86f7671bedcec89bccdd98faec98581a528b0aa6906a15be19ddb2a433ea99"></a>

<a id="canonical-020126cbb20448e00c3a29efe658abce4d214351a514d737c6ecbc9592fa5bf0"></a>

## read property — timeouts / 1c99646801be / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-6059d26ce8b789a5b19d1e72e2a0c20366c48654687517cb7033b8fddb6352a4"></a>

<a id="canonical-203689c323069b80ba20410f58c306efd798cc7df71f1e859831c5f882b9531b"></a>

## update property — timeouts / 1c99646801be / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-25ce6d26911f6e4f22735be2de40732e0d1e9437127968c7d101dc01deca2d74"></a>

## Next pages — timeouts / 1c99646801be / 8

- [Property reference](resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
