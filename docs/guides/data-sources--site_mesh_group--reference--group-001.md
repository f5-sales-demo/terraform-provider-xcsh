---
page_title: "xcsh_site_mesh_group reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group reference."
---

# xcsh_site_mesh_group reference

<a id="canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37559a450143e59fc9e193b98a7ff680724e7a7841a70553a8053fa0a103fadf"></a>

## Property reference — Property reference / 11e9a49d41e8 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- Property reference

<a id="canonical-b28a1728fb20e26fcad68dc347a17d6c6c241b70401d5920876cb24b10416804"></a>

## Direct properties — Property reference / 11e9a49d41e8 / 3

<a id="canonical-d6a4b2277c08a8f2c100c65ef3cca907e302127c3cbb658acd51bc37c983f0b4"></a>

<a id="canonical-c1b93249a7ae8ac86dae1889cb74cad1854d2754f8bcf8cb814457daae8b87da"></a>

## annotations property — Property reference / 11e9a49d41e8 / 4

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

- [bfd_disabled](data-sources--site_mesh_group--reference--group-001.md#canonical-181016123fd35f48774b63a20fba48ea9d8c3bc5ceadd60ae9771323fbe43908): complete subsection reference.

- [bfd_enabled](data-sources--site_mesh_group--reference--group-001.md#canonical-e797489225116b86a1f9140200fb5ff7bd90f0352826b0791707084229268088): complete subsection reference.

<a id="canonical-6fb94e8f0faf86c6fe414946512e358de45945dac227f9b40b1da515b952bcdb"></a>

<a id="canonical-c412c465dd1a0eea6f14212a0bdcc055ae81d8d38f77522f84952ff1105eb0c0"></a>

## description property — Property reference / 11e9a49d41e8 / 5

Type: `"string"`. Computed.

Description of the SiteMeshGroup.

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

- [disable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-06df6d27d32aace149b0ac37c919b333c3bf80a615c142bc0db33b31436dd76a): complete subsection reference.

- [enable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-440f8af49585155ef5e4749e73de16fcf41404741bd29a0ff592286dc9d2efc0): complete subsection reference.

- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-85d44cf5d1ff6dbc399d97dd1ea896712545d638bcda608438092f6b90606c26): complete subsection reference.

- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f455434d0dd86ee67e53abeed9b659387855409cb4510673909d35bdf46604bb): complete subsection reference.

<a id="canonical-49a767c020359a11c08ecbadcb54fc5cfb6f237a0d0b65b09498700f63ee842a"></a>

<a id="canonical-bfd40609fee6db9abfb7ba532c05a0dc6888807d44c3ed37f9eede6e669612a9"></a>

## id property — Property reference / 11e9a49d41e8 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-fc00bbf4ae152de43821d08fd152c973bd7c2e09e1ddeddce98149119490ab4e"></a>

<a id="canonical-5037334cc155287b6c636568c2fd1f57a2382357d150880dcf454912fa29a6a7"></a>

## labels property — Property reference / 11e9a49d41e8 / 7

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

<a id="canonical-8d7e0efc867ceb2c6552e312f084b7097075fd52e4ddc707e10ee871e008e13b"></a>

<a id="canonical-d200704199fd631f11ef9134aa50d2f0c9cefe4b8df2a6dd1477097745b00657"></a>

## name property — Property reference / 11e9a49d41e8 / 8

Type: `"string"`. Required.

Name of the SiteMeshGroup.

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

<a id="canonical-25d3a10b79d8b7bd2910a7b2425c34bb1eadf128a079575ac5c0df133c5f220b"></a>

<a id="canonical-bd3bca1dadbca153bb067192c373534a4faf03dea13dfb53b9d95dc2d4de3707"></a>

## namespace property — Property reference / 11e9a49d41e8 / 9

Type: `"string"`. Required.

Namespace where the SiteMeshGroup exists.

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

- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647): complete subsection reference.

- [virtual_site](data-sources--site_mesh_group--reference--group-001.md#canonical-a9e329125b7a1d7844be73cac31dcf65b489fd929a01b43d4a2c5b62c555a0ae): complete subsection reference.

<a id="canonical-b2c70b84aea3da1b02a61feb3210cc36f142209fe025ed2e36aa25b45561775e"></a>

## All schema paths — Property reference / 11e9a49d41e8 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--site_mesh_group--reference--group-001.md#canonical-d6a4b2277c08a8f2c100c65ef3cca907e302127c3cbb658acd51bc37c983f0b4) |
| `bfd_disabled` | [bfd_disabled](data-sources--site_mesh_group--reference--group-001.md#canonical-6184342a27952146ad3492aa88b79fae8a963a7b5f90794b3bfefeaf88230184) |
| `bfd_enabled` | [bfd_enabled](data-sources--site_mesh_group--reference--group-001.md#canonical-f9d1e1d3e26828878a3a259dc35f45e196a4d405084719dd1e6342dd133e9db9) |
| `bfd_enabled.multiplier` | [bfd_enabled.multiplier](data-sources--site_mesh_group--reference--group-001.md#canonical-724710e21252a39a10c405f377c0baddd921720ff1549423add32e6081766669) |
| `bfd_enabled.receive_interval_milliseconds` | [bfd_enabled.receive_interval_milliseconds](data-sources--site_mesh_group--reference--group-001.md#canonical-1fa5358c9a66f90d49e63969845050016db078c99d176c037d811f6897a129be) |
| `bfd_enabled.transmit_interval_milliseconds` | [bfd_enabled.transmit_interval_milliseconds](data-sources--site_mesh_group--reference--group-001.md#canonical-19ca8d36ea09b2b8a3c24b5ec6d959698e53e320fbbd6d161a21dcee84b3cd5f) |
| `description` | [description](data-sources--site_mesh_group--reference--group-001.md#canonical-6fb94e8f0faf86c6fe414946512e358de45945dac227f9b40b1da515b952bcdb) |
| `disable_re_fallback` | [disable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-bba38a5ec121288dfe8f6195bdab4bdbe6cefa73e13d26a4c167e645d0d22fca) |
| `enable_re_fallback` | [enable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-4ad1fe475bb25cde2ee55aecd8fb15d87e48442f34cd6b6b5e9cea4d48da34a9) |
| `full_mesh` | [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f7739600ff27c71cf742280b86a77930b9e479a8f1fea372a5d29103d731003e) |
| `full_mesh.control_and_data_plane_mesh` | [full_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-55d593bf0d68a5c033538ebd8f860a665da2e8625d0d046949b80fa4ada969b7) |
| `full_mesh.data_plane_mesh` | [full_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-6a93d8043b77aba2e17a6169d29cfa027325976c7c2e2d08ce2d2216924b8237) |
| `hub_mesh` | [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-253a37fa61e54ee6fe4845ca76cf541120f94acd51cd9c687ef75cf816cab27c) |
| `hub_mesh.control_and_data_plane_mesh` | [hub_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-15b5b2a9f1a92a7acaaa5e4ead409fd9d9cb494ff0f060ac9972d2bb48d025ce) |
| `hub_mesh.data_plane_mesh` | [hub_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-454228b77d3d3c6643d462b19d4d9a0fba15d4454ea2235ae2caccf94d8f6819) |
| `id` | [id](data-sources--site_mesh_group--reference--group-001.md#canonical-49a767c020359a11c08ecbadcb54fc5cfb6f237a0d0b65b09498700f63ee842a) |
| `labels` | [labels](data-sources--site_mesh_group--reference--group-001.md#canonical-fc00bbf4ae152de43821d08fd152c973bd7c2e09e1ddeddce98149119490ab4e) |
| `name` | [name](data-sources--site_mesh_group--reference--group-001.md#canonical-8d7e0efc867ceb2c6552e312f084b7097075fd52e4ddc707e10ee871e008e13b) |
| `namespace` | [namespace](data-sources--site_mesh_group--reference--group-001.md#canonical-25d3a10b79d8b7bd2910a7b2425c34bb1eadf128a079575ac5c0df133c5f220b) |
| `spoke_mesh` | [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-a1e25e9fbb20e21385d137101d6a1723d401ea33cc06cd2aabbe8603860a3f62) |
| `spoke_mesh.control_and_data_plane_mesh` | [spoke_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-da25d44ac7877f1053b7563b03906181c1c136038e9a78f941b431ac86ccebc2) |
| `spoke_mesh.data_plane_mesh` | [spoke_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-96332c5fab40c6fb154e6f3ba469ee0b2b3c9a14beaca8be27fda12f343dd80d) |
| `spoke_mesh.hub_mesh_group` | [spoke_mesh.hub_mesh_group](data-sources--site_mesh_group--reference--group-001.md#canonical-5abe78c44df1a87c00f1f67a316686e95acf2dfed6f025ab09871ce53e313d7e) |
| `spoke_mesh.hub_mesh_group.name` | [spoke_mesh.hub_mesh_group.name](data-sources--site_mesh_group--reference--group-001.md#canonical-6aee091a3b6e30b0d5b7ca3da81a0629165844304c1bfd449658883b418dc081) |
| `spoke_mesh.hub_mesh_group.namespace` | [spoke_mesh.hub_mesh_group.namespace](data-sources--site_mesh_group--reference--group-001.md#canonical-a94c016d0ec2321d0693873fa43c20f8c335a9e48a28bed0585a52ea9211af31) |
| `spoke_mesh.hub_mesh_group.tenant` | [spoke_mesh.hub_mesh_group.tenant](data-sources--site_mesh_group--reference--group-001.md#canonical-3fe836229bad5ac2b65e3121133d521b55fcfa52579bfb58d39cdda7afb08dd5) |
| `virtual_site` | [virtual_site](data-sources--site_mesh_group--reference--group-001.md#canonical-6398a30466d2b176dfe8ff00a0299bd4451b3242d6f4fdb848f0d0cecfe853e1) |
| `virtual_site.kind` | [virtual_site.kind](data-sources--site_mesh_group--reference--group-001.md#canonical-e652f3f9c2c5c02793847ff824f3dc451bbd57c1fe4b89d2fe7008e82b1a2e2c) |
| `virtual_site.name` | [virtual_site.name](data-sources--site_mesh_group--reference--group-001.md#canonical-c581871ba7cbb9cc9ad0a4bf78b8d4ac42f34212ce9ba664c10bbbd682c5ab6b) |
| `virtual_site.namespace` | [virtual_site.namespace](data-sources--site_mesh_group--reference--group-001.md#canonical-2160e97177870789e69bd43cf5a0a01441408e6aa7b114da18e9b7294f3a8096) |
| `virtual_site.tenant` | [virtual_site.tenant](data-sources--site_mesh_group--reference--group-001.md#canonical-b2cf00b0a425edb1d71ee5ddbb4ba3c3e650a7b01c6b72387652d4ea7b1358e3) |
| `virtual_site.uid` | [virtual_site.uid](data-sources--site_mesh_group--reference--group-001.md#canonical-2c93c81f386ecc16f2bdd4bf5ba3c7842ad10a65ab13b4cb9112a372adcfb7e3) |

<a id="canonical-59e4578b3215abe10f4b1a240c9703d065f0fc93c42a1a1047e3b247ae79f053"></a>

## Next pages — Property reference / 11e9a49d41e8 / 11

- [bfd_disabled](data-sources--site_mesh_group--reference--group-001.md#canonical-181016123fd35f48774b63a20fba48ea9d8c3bc5ceadd60ae9771323fbe43908)
- [bfd_enabled](data-sources--site_mesh_group--reference--group-001.md#canonical-e797489225116b86a1f9140200fb5ff7bd90f0352826b0791707084229268088)
- [disable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-06df6d27d32aace149b0ac37c919b333c3bf80a615c142bc0db33b31436dd76a)
- [enable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-440f8af49585155ef5e4749e73de16fcf41404741bd29a0ff592286dc9d2efc0)
- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-85d44cf5d1ff6dbc399d97dd1ea896712545d638bcda608438092f6b90606c26)
- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f455434d0dd86ee67e53abeed9b659387855409cb4510673909d35bdf46604bb)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647)
- [virtual_site](data-sources--site_mesh_group--reference--group-001.md#canonical-a9e329125b7a1d7844be73cac31dcf65b489fd929a01b43d4a2c5b62c555a0ae)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-181016123fd35f48774b63a20fba48ea9d8c3bc5ceadd60ae9771323fbe43908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89695b37f22fcc7a53468b6d8b6ac82b360b1cb12d0d8fff6926a4ac484f8539"></a>

## bfd_disabled — bfd_disabled / fdb90511dbdd / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- bfd_disabled

<a id="canonical-6184342a27952146ad3492aa88b79fae8a963a7b5f90794b3bfefeaf88230184"></a>

Type: `["object", {}]`. Computed.

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

- [bfd_disabled](data-sources--site_mesh_group--reference--group-001.md#canonical-6184342a27952146ad3492aa88b79fae8a963a7b5f90794b3bfefeaf88230184)
- [bfd_enabled](data-sources--site_mesh_group--reference--group-001.md#canonical-f9d1e1d3e26828878a3a259dc35f45e196a4d405084719dd1e6342dd133e9db9)

Select alternatives according to the provider validators above.

<a id="canonical-a4700433be035a491ccd06b00a3be04f5fa9b9ee1b20d4cd01b0892041cc4d27"></a>

## Direct properties — bfd_disabled / fdb90511dbdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1928c14c70de35de3127b160b9f96815506037f242ca775631692429324c7d09"></a>

## Next pages — bfd_disabled / fdb90511dbdd / 4

- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-e797489225116b86a1f9140200fb5ff7bd90f0352826b0791707084229268088"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-456b2e41cafb3659b0faca6c68a1932d17e6b66ca93bd650a54a7f662eac8327"></a>

## bfd_enabled — bfd_enabled / cd8b74770180 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- bfd_enabled

<a id="canonical-f9d1e1d3e26828878a3a259dc35f45e196a4d405084719dd1e6342dd133e9db9"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

Upstream description:

BFD parameters.

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

<a id="canonical-cc705e3c61da257e19e5a3833e9bfe0356914899d03cc9e90d7d04910847ae1d"></a>

## Direct properties — bfd_enabled / cd8b74770180 / 3

<a id="canonical-724710e21252a39a10c405f377c0baddd921720ff1549423add32e6081766669"></a>

<a id="canonical-5958cf5fd5973bfb513b3f9e2ef4aaa7da7a91b9cf7df96fbe1247561629a18c"></a>

## multiplier property — bfd_enabled / cd8b74770180 / 4

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

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

<a id="canonical-1fa5358c9a66f90d49e63969845050016db078c99d176c037d811f6897a129be"></a>

<a id="canonical-6adb9be03475773e13a0c15cebc2b4d17ad89e1bf9133c7384229abb63d1134c"></a>

## receive_interval_milliseconds property — bfd_enabled / cd8b74770180 / 5

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

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

<a id="canonical-19ca8d36ea09b2b8a3c24b5ec6d959698e53e320fbbd6d161a21dcee84b3cd5f"></a>

<a id="canonical-25d78c8edce2d8c48e188e599e44f97a3eb1bebe53b2765ec4fd6d9ff3a79024"></a>

## transmit_interval_milliseconds property — bfd_enabled / cd8b74770180 / 6

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

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

<a id="canonical-681e283cc90fdcfd30adef293c99e40d8358a598dd709d1a5e542ef6fe6be741"></a>

## Next pages — bfd_enabled / cd8b74770180 / 7

- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-06df6d27d32aace149b0ac37c919b333c3bf80a615c142bc0db33b31436dd76a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5602cb3203b7c6da9ea19777559be71518996141efd258fcf001e6812ebdc838"></a>

## disable_re_fallback — disable_re_fallback / cb93383a4c85 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- disable_re_fallback

<a id="canonical-bba38a5ec121288dfe8f6195bdab4bdbe6cefa73e13d26a4c167e645d0d22fca"></a>

Type: `["object", {}]`. Computed.

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

- [disable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-bba38a5ec121288dfe8f6195bdab4bdbe6cefa73e13d26a4c167e645d0d22fca)
- [enable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-4ad1fe475bb25cde2ee55aecd8fb15d87e48442f34cd6b6b5e9cea4d48da34a9)

Select alternatives according to the provider validators above.

<a id="canonical-d0c33d09f7187fbf8524436fd44dcf968e89d98d996e1f8f5e1177278381c64c"></a>

## Direct properties — disable_re_fallback / cb93383a4c85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-937cdc0148bb4219e5cc36d101286321590f09709fbff31998f1e889da042319"></a>

## Next pages — disable_re_fallback / cb93383a4c85 / 4

- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-440f8af49585155ef5e4749e73de16fcf41404741bd29a0ff592286dc9d2efc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90cbeab22d3baa0cbfc6d7b19564bb10cd60adc03f2470c621df3f0ab3bb8e41"></a>

## enable_re_fallback — enable_re_fallback / 91170f903c4a / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- enable_re_fallback

<a id="canonical-4ad1fe475bb25cde2ee55aecd8fb15d87e48442f34cd6b6b5e9cea4d48da34a9"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-df6014d839e55c7282d9b66b87fb38348a8f9df01b3a2ebb12a28d3dec2eada8"></a>

## Direct properties — enable_re_fallback / 91170f903c4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7584715ab7b038f1e0306b076a955d23af021175aff58cc2a5e35872ed5a04e"></a>

## Next pages — enable_re_fallback / 91170f903c4a / 4

- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-85d44cf5d1ff6dbc399d97dd1ea896712545d638bcda608438092f6b90606c26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37ac233fc7a9bc18dbd2c08066d69444391f3764412c19b072a51d33ea94b259"></a>

## full_mesh — full_mesh / 9e90eb5819e4 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- full_mesh

<a id="canonical-f7739600ff27c71cf742280b86a77930b9e479a8f1fea372a5d29103d731003e"></a>

Type: `"single"`. Computed.

\[OneOf: full\_mesh, hub\_mesh, spoke\_mesh\] Full Mesh. Details of Full Mesh Group Type.

Upstream description:

Details of Full Mesh Group Type.

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

- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f7739600ff27c71cf742280b86a77930b9e479a8f1fea372a5d29103d731003e)
- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-253a37fa61e54ee6fe4845ca76cf541120f94acd51cd9c687ef75cf816cab27c)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-a1e25e9fbb20e21385d137101d6a1723d401ea33cc06cd2aabbe8603860a3f62)

Select alternatives according to the provider validators above.

<a id="canonical-108ee13622d6ea7f94bb35ef8cff58e9c31a06b8ed7a3861666c84dc645accd9"></a>

## Direct properties — full_mesh / 9e90eb5819e4 / 3

- [control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-b1856f51fe5783fc08223111e14f02c494c50ddab64acf4942004db1583da1c4): complete subsection reference.

- [data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-0806c6872152c7c2e53d7058d3cef3deafca224af5fce3ffc4e9d002462942bc): complete subsection reference.

<a id="canonical-4bd0a465375febb041ec8c3c8c14e96d6961f5922baf9bfe5f8d7dcd70880b62"></a>

## Next pages — full_mesh / 9e90eb5819e4 / 4

- [full_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-b1856f51fe5783fc08223111e14f02c494c50ddab64acf4942004db1583da1c4)
- [full_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-0806c6872152c7c2e53d7058d3cef3deafca224af5fce3ffc4e9d002462942bc)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-b1856f51fe5783fc08223111e14f02c494c50ddab64acf4942004db1583da1c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72e6730cc9b336cffc3fcc820ef603c7e70e2741af546d24f0a3a6d1c72749b1"></a>

## full_mesh.control_and_data_plane_mesh — full_mesh.control_and_data_plane_mesh / 5a76df1a76ca / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-85d44cf5d1ff6dbc399d97dd1ea896712545d638bcda608438092f6b90606c26)
- full_mesh.control_and_data_plane_mesh

<a id="canonical-55d593bf0d68a5c033538ebd8f860a665da2e8625d0d046949b80fa4ada969b7"></a>

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

<a id="canonical-3ef6277081c9aeb06cb09f08ab8c89da719202a7f0b91a5a5241f870c583d1cc"></a>

## Direct properties — full_mesh.control_and_data_plane_mesh / 5a76df1a76ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-941368ed9cdc7fc7c4e6b3d2fb1b675cf218629b9f0f645c5a2b9a20d12a705b"></a>

## Next pages — full_mesh.control_and_data_plane_mesh / 5a76df1a76ca / 4

- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-85d44cf5d1ff6dbc399d97dd1ea896712545d638bcda608438092f6b90606c26)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-0806c6872152c7c2e53d7058d3cef3deafca224af5fce3ffc4e9d002462942bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec922798e43b339c5ab183f1f61917f0f2c29f5d4611a42a42335967ce6fe16a"></a>

## full_mesh.data_plane_mesh — full_mesh.data_plane_mesh / f1f95f36f7b8 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-85d44cf5d1ff6dbc399d97dd1ea896712545d638bcda608438092f6b90606c26)
- full_mesh.data_plane_mesh

<a id="canonical-6a93d8043b77aba2e17a6169d29cfa027325976c7c2e2d08ce2d2216924b8237"></a>

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

<a id="canonical-4e68a996d72227665c0132913ccce77c77a3407fc03933beb225992efbac2db0"></a>

## Direct properties — full_mesh.data_plane_mesh / f1f95f36f7b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0a2116e69bb93e3c341ce1af46a2116f4e38ffe8981b268d3014e488cb109b5a"></a>

## Next pages — full_mesh.data_plane_mesh / f1f95f36f7b8 / 4

- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-85d44cf5d1ff6dbc399d97dd1ea896712545d638bcda608438092f6b90606c26)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-f455434d0dd86ee67e53abeed9b659387855409cb4510673909d35bdf46604bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89a8c067c9d4c82fe8beecc5bb6228c1f15e34ba3c82e1830408e20fdddfa47f"></a>

## hub_mesh — hub_mesh / 4a29b6a7c1ce / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- hub_mesh

<a id="canonical-253a37fa61e54ee6fe4845ca76cf541120f94acd51cd9c687ef75cf816cab27c"></a>

Type: `"single"`. Computed.

Hub Full Mesh. Details of Hub Full Mesh Group Type.

Upstream description:

Details of Hub Full Mesh Group Type.

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

<a id="canonical-ded401a506f7b4d8055340c213d98482ee57ce9b706022aabc0e9053d60dd213"></a>

## Direct properties — hub_mesh / 4a29b6a7c1ce / 3

- [control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-83a63ac6b46ee8fed364d1321c0de5fed9ab75b2c41c70ea7cddc12ef949dd09): complete subsection reference.

- [data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-cbb5bc331d2238c25876f28ee752b4e2e2169776940f8fba03bfee134782c638): complete subsection reference.

<a id="canonical-03b3af7be6cf2723386ba52b14d78280014d0c18833f1fc109377f8480fbf2e7"></a>

## Next pages — hub_mesh / 4a29b6a7c1ce / 4

- [hub_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-83a63ac6b46ee8fed364d1321c0de5fed9ab75b2c41c70ea7cddc12ef949dd09)
- [hub_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-cbb5bc331d2238c25876f28ee752b4e2e2169776940f8fba03bfee134782c638)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-83a63ac6b46ee8fed364d1321c0de5fed9ab75b2c41c70ea7cddc12ef949dd09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cd61b233a02ad7d28c08223303ef8f7e5f11d37e0a97a8bde85dba235cd2d83"></a>

## hub_mesh.control_and_data_plane_mesh — hub_mesh.control_and_data_plane_mesh / c648d6305a77 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f455434d0dd86ee67e53abeed9b659387855409cb4510673909d35bdf46604bb)
- hub_mesh.control_and_data_plane_mesh

<a id="canonical-15b5b2a9f1a92a7acaaa5e4ead409fd9d9cb494ff0f060ac9972d2bb48d025ce"></a>

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

<a id="canonical-abbbd2008b94b227288fa5d641bc433d3f2268e9f60bf19a957bb56bf53265c9"></a>

## Direct properties — hub_mesh.control_and_data_plane_mesh / c648d6305a77 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-df55995d49aa3d53b2bd6377e5105cbffe37d0b9d2e37408123f82effbf7e1c8"></a>

## Next pages — hub_mesh.control_and_data_plane_mesh / c648d6305a77 / 4

- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f455434d0dd86ee67e53abeed9b659387855409cb4510673909d35bdf46604bb)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-cbb5bc331d2238c25876f28ee752b4e2e2169776940f8fba03bfee134782c638"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-284edb2508ee05704ed69dd40aa280c4a575c4353ec5f57908199268c3712e9b"></a>

## hub_mesh.data_plane_mesh — hub_mesh.data_plane_mesh / a3d9b2bdb992 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f455434d0dd86ee67e53abeed9b659387855409cb4510673909d35bdf46604bb)
- hub_mesh.data_plane_mesh

<a id="canonical-454228b77d3d3c6643d462b19d4d9a0fba15d4454ea2235ae2caccf94d8f6819"></a>

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

<a id="canonical-0139e9430ca72326864076e6ce15c6af3e486ebfb3831b84d4097639dd686f55"></a>

## Direct properties — hub_mesh.data_plane_mesh / a3d9b2bdb992 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8db8a051f18c4789508130fca2c8c259b8f56660baf54c455e2dc4f13ad2248d"></a>

## Next pages — hub_mesh.data_plane_mesh / a3d9b2bdb992 / 4

- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f455434d0dd86ee67e53abeed9b659387855409cb4510673909d35bdf46604bb)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97c5f3018834947d3ef2b46c93da465d9ea266ace7b453a236fd1312b6970ea4"></a>

## spoke_mesh — spoke_mesh / 19e0efbd6232 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- spoke_mesh

<a id="canonical-a1e25e9fbb20e21385d137101d6a1723d401ea33cc06cd2aabbe8603860a3f62"></a>

Type: `"single"`. Computed.

Spoke. Details of Spoke Mesh Group Type.

Upstream description:

Details of Spoke Mesh Group Type.

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

<a id="canonical-b1388187b3ccf6c43f0eef5210f4ee8e775947604d5d2aee6698c21525308b27"></a>

## Direct properties — spoke_mesh / 19e0efbd6232 / 3

- [control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-4c704162ddf50ab064adf54d9293a43f79257a60fe18ab0663bcf83bd9279330): complete subsection reference.

- [data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-a78192e24c579a4f256febdfc7e8bcd028857b65eb7a3772d82df00cc3b830ec): complete subsection reference.

- [hub_mesh_group](data-sources--site_mesh_group--reference--group-001.md#canonical-7ecc9164160c98b68d850d67628225a12f42e8c82292a34d22f553da688db832): complete subsection reference.

<a id="canonical-50429f985f26a67a90dab2bb4d48cf5d29ae6562bd72bef4e000230845f8788f"></a>

## Next pages — spoke_mesh / 19e0efbd6232 / 4

- [spoke_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-4c704162ddf50ab064adf54d9293a43f79257a60fe18ab0663bcf83bd9279330)
- [spoke_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-a78192e24c579a4f256febdfc7e8bcd028857b65eb7a3772d82df00cc3b830ec)
- [spoke_mesh.hub_mesh_group](data-sources--site_mesh_group--reference--group-001.md#canonical-7ecc9164160c98b68d850d67628225a12f42e8c82292a34d22f553da688db832)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-4c704162ddf50ab064adf54d9293a43f79257a60fe18ab0663bcf83bd9279330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad5ffa57c2089da770e816f799b16a40f6d4ef148b57d71ab9a38bad88a32330"></a>

## spoke_mesh.control_and_data_plane_mesh — spoke_mesh.control_and_data_plane_mesh / 015b8b88b7af / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647)
- spoke_mesh.control_and_data_plane_mesh

<a id="canonical-da25d44ac7877f1053b7563b03906181c1c136038e9a78f941b431ac86ccebc2"></a>

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

<a id="canonical-d3a01d03dbc26f432568b9d3b959e7bb76f43358d998140fbb19d2e44c73f1f9"></a>

## Direct properties — spoke_mesh.control_and_data_plane_mesh / 015b8b88b7af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a10a22fbb1a4d1f3b6e9224d2a20b4abf3445595a9f0c0e525bca23153d3984"></a>

## Next pages — spoke_mesh.control_and_data_plane_mesh / 015b8b88b7af / 4

- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-a78192e24c579a4f256febdfc7e8bcd028857b65eb7a3772d82df00cc3b830ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fd45f61a8d38c0c699800e445cb905c5bd378508c684d17bf6b57b2f00e920b"></a>

## spoke_mesh.data_plane_mesh — spoke_mesh.data_plane_mesh / 5f6dd4760986 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647)
- spoke_mesh.data_plane_mesh

<a id="canonical-96332c5fab40c6fb154e6f3ba469ee0b2b3c9a14beaca8be27fda12f343dd80d"></a>

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

<a id="canonical-9d35275036223ac4a7f6a6b61c0059154367c98d79901f050eac9e34245f0aa7"></a>

## Direct properties — spoke_mesh.data_plane_mesh / 5f6dd4760986 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3365bb161870cb75b0aa8b4100697bb89aade130b40a94697d96dd81a7d4bdff"></a>

## Next pages — spoke_mesh.data_plane_mesh / 5f6dd4760986 / 4

- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-7ecc9164160c98b68d850d67628225a12f42e8c82292a34d22f553da688db832"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32dbe283965ccc2f89b8cf87cafd573af264cc21d5fa6462d3eb64a474a700c7"></a>

## spoke_mesh.hub_mesh_group — spoke_mesh.hub_mesh_group / 2f5817578862 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647)
- spoke_mesh.hub_mesh_group

<a id="canonical-5abe78c44df1a87c00f1f67a316686e95acf2dfed6f025ab09871ce53e313d7e"></a>

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

<a id="canonical-c2d20940df7f1478114f7286bd68c924216025593df03453009b1cb9bc192c98"></a>

## Direct properties — spoke_mesh.hub_mesh_group / 2f5817578862 / 3

<a id="canonical-6aee091a3b6e30b0d5b7ca3da81a0629165844304c1bfd449658883b418dc081"></a>

<a id="canonical-751814e2e0e4e84b4fa7101760583f54c910e335ff8a6b62565960c062907684"></a>

## name property — spoke_mesh.hub_mesh_group / 2f5817578862 / 4

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

<a id="canonical-a94c016d0ec2321d0693873fa43c20f8c335a9e48a28bed0585a52ea9211af31"></a>

<a id="canonical-ce4f5060a708803e3c69f96bd8cd02013352edd9d4698fa7dfcdf0c1e1badc0a"></a>

## namespace property — spoke_mesh.hub_mesh_group / 2f5817578862 / 5

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

<a id="canonical-3fe836229bad5ac2b65e3121133d521b55fcfa52579bfb58d39cdda7afb08dd5"></a>

<a id="canonical-7c2fc74a9c431251fcc3c01aba57984aa0a70bce380f230ba9bad5e4d0902687"></a>

## tenant property — spoke_mesh.hub_mesh_group / 2f5817578862 / 6

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

<a id="canonical-9964e9f99ac6d7ec854375bc78bee2eafef1e76a88a7f030189b5a1d7f0c7426"></a>

## Next pages — spoke_mesh.hub_mesh_group / 2f5817578862 / 7

- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-f3b049be1b404fa59d54932cf1f1394c51633bb62126853b237bd2388b6c1647)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-a9e329125b7a1d7844be73cac31dcf65b489fd929a01b43d4a2c5b62c555a0ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da37bba9eb64bb0a329b560eba6f3b19eede7b7007ea2bed96018559fc0cc3cf"></a>

## virtual_site — virtual_site / 87818aa503ca / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- virtual_site

<a id="canonical-6398a30466d2b176dfe8ff00a0299bd4451b3242d6f4fdb848f0d0cecfe853e1"></a>

Type: `"list"`. Computed.

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

<a id="canonical-3e3979cc33adeeca869ea7bd4ddca78cfe6b5a44c3e7b467e7f5ce1d1cb4542d"></a>

## Direct properties — virtual_site / 87818aa503ca / 3

<a id="canonical-e652f3f9c2c5c02793847ff824f3dc451bbd57c1fe4b89d2fe7008e82b1a2e2c"></a>

<a id="canonical-cb1a999ce6775b0c3ee49dde425840e600431b4f7f10194e662aab72c47ae4ac"></a>

## kind property — virtual_site / 87818aa503ca / 4

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

<a id="canonical-c581871ba7cbb9cc9ad0a4bf78b8d4ac42f34212ce9ba664c10bbbd682c5ab6b"></a>

<a id="canonical-a8f672124935f874cda7c851361b9a6fc31422afe6e14446a1cea4512f60060a"></a>

## name property — virtual_site / 87818aa503ca / 5

Type: `"string"`. Computed.

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

<a id="canonical-2160e97177870789e69bd43cf5a0a01441408e6aa7b114da18e9b7294f3a8096"></a>

<a id="canonical-eb35d4f1f0ce8c4e87bd4b9e5abe1fad6db3ed56281d0a19572f4e99c355631c"></a>

## namespace property — virtual_site / 87818aa503ca / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-b2cf00b0a425edb1d71ee5ddbb4ba3c3e650a7b01c6b72387652d4ea7b1358e3"></a>

<a id="canonical-81877663d3cfcb3847d39999ae60a92ac0121e7d07629f7a2e775d3604ef435d"></a>

## tenant property — virtual_site / 87818aa503ca / 7

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

<a id="canonical-2c93c81f386ecc16f2bdd4bf5ba3c7842ad10a65ab13b4cb9112a372adcfb7e3"></a>

<a id="canonical-7ebda801c6d753892057fd1472dbb357c0773d6a083e953b4a41a260808a0611"></a>

## uid property — virtual_site / 87818aa503ca / 8

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

<a id="canonical-1150ac9621ffdb3798c485adb8dd661c22bb44bb005d273eaaa8b9f4aa02d3ae"></a>

## Next pages — virtual_site / 87818aa503ca / 9

- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
