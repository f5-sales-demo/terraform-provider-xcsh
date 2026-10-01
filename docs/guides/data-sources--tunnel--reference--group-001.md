---
page_title: "xcsh_tunnel reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel reference."
---

# xcsh_tunnel reference

<a id="canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc701a353717f1d7159a3fb92dbcfb11d4354bfc6c3dee9a221308d4b2a22250"></a>

## Property reference — Property reference / f7303f87cc5e / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- Property reference

<a id="canonical-5659abc7cfb87709423251c86c3e0947fcc2fbf9535eb598e76bf99130586ee8"></a>

## Direct properties — Property reference / f7303f87cc5e / 3

<a id="canonical-64dce7103097983aed789777227bdafd399df7e1a8892f4d505ba5458e29bfc7"></a>

<a id="canonical-85c4dfe52b2eef36c4dbf7aca8bc93ec3c7bf52a5d75289c9a01869456924bfc"></a>

## annotations property — Property reference / f7303f87cc5e / 4

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

<a id="canonical-b7e846863ad9ce986f8ef24bcdbecd760d3da1f35846ca6dfbc6536ac2f08529"></a>

<a id="canonical-6b13d645666dd695cf45b0c0fcbc7e243e4ea8cf14c8c9f16a9bba3743f03871"></a>

## description property — Property reference / f7303f87cc5e / 5

Type: `"string"`. Computed.

Description of the Tunnel.

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

<a id="canonical-b60644962501fc193f3c69651cfc5b0db6a7b13bf7794b2c426c05f79f96e875"></a>

<a id="canonical-b40bb5eb22600500cf3d034d5d8ae1908c9e50a120c325e4564ea6830326b19c"></a>

## id property — Property reference / f7303f87cc5e / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-056da7e7e81d050bef72a24c29ea0066862094b073225c9887fc861564ef187c"></a>

<a id="canonical-9351808e3f7df89055b16704048f65d3e8e493a912ef73a8d0586a0ec0d6e9e8"></a>

## labels property — Property reference / f7303f87cc5e / 7

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

- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6): complete subsection reference.

<a id="canonical-774dc044926ef324538e8e289a9a45ef7c110dc7d36b1940228a8d7838eeb2f1"></a>

<a id="canonical-6421d756d20b55da5cf8168125892e18c19d9b78370bf58e4f581286a08141a8"></a>

## name property — Property reference / f7303f87cc5e / 8

Type: `"string"`. Required.

Name of the Tunnel.

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

<a id="canonical-feb1dea4311f58e1ea94f8310642ecf5207fc7e66234248169dbb6a966a89fe5"></a>

<a id="canonical-a7d620388af362e3e76099c799c0d2edb97c2682b07a8a977bb5809a20d25604"></a>

## namespace property — Property reference / f7303f87cc5e / 9

Type: `"string"`. Required.

Namespace where the Tunnel exists.

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

- [params](data-sources--tunnel--reference--group-001.md#canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7): complete subsection reference.

- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8): complete subsection reference.

<a id="canonical-e68c239f6704b909862be98b8fcd40d485911adca9b3e783291f588182ea7c92"></a>

<a id="canonical-8690b9e3c4f51b7d05edc4c78758255aa39981eebdec42638da4ea4e6ae71d3c"></a>

## tunnel_type property — Property reference / f7303f87cc5e / 10

Type: `"string"`. Computed.

\[Enum: IPSEC\_PSK|GRE\] Supported tunnel types are IPsec IPsec tunnel type with PSK GRE tunnel
type. Possible values are \`IPSEC\_PSK\`, \`GRE\`. Defaults to \`IPSEC\_PSK\`.

Upstream description:

Supported tunnel types are IPsec

IPsec tunnel type with PSK GRE tunnel type.

Receipt-pinned upstream constraints:

```json
{
  "default": "IPSEC_PSK",
  "enum": [
    "IPSEC_PSK",
    "GRE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9db157148f7cfac8bca0f416d9109ebeacc2d9157c7e43069b2f6bd8054d1ce4"></a>

## All schema paths — Property reference / f7303f87cc5e / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--tunnel--reference--group-001.md#canonical-64dce7103097983aed789777227bdafd399df7e1a8892f4d505ba5458e29bfc7) |
| `description` | [description](data-sources--tunnel--reference--group-001.md#canonical-b7e846863ad9ce986f8ef24bcdbecd760d3da1f35846ca6dfbc6536ac2f08529) |
| `id` | [id](data-sources--tunnel--reference--group-001.md#canonical-b60644962501fc193f3c69651cfc5b0db6a7b13bf7794b2c426c05f79f96e875) |
| `labels` | [labels](data-sources--tunnel--reference--group-001.md#canonical-056da7e7e81d050bef72a24c29ea0066862094b073225c9887fc861564ef187c) |
| `local_ip` | [local_ip](data-sources--tunnel--reference--group-001.md#canonical-97edccce064646a89b060c31597eb265b3631d9bfe01bbfbf0b29cd46b049359) |
| `local_ip.intf` | [local_ip.intf](data-sources--tunnel--reference--group-001.md#canonical-0d9d6d0c4bc10557bff9a79200ae3e8c936d5f8a0493de7b7b9281e27d8f3965) |
| `local_ip.intf.local_intf` | [local_ip.intf.local_intf](data-sources--tunnel--reference--group-001.md#canonical-ac0588a9159a3f358659236ad624fb80b5b42cb3fb4b55cd88f07d28b5f71394) |
| `local_ip.intf.local_intf.kind` | [local_ip.intf.local_intf.kind](data-sources--tunnel--reference--group-001.md#canonical-eac68472faaf99302019458a47a451e75875e62f410d887a25c7602e7f1a8d63) |
| `local_ip.intf.local_intf.name` | [local_ip.intf.local_intf.name](data-sources--tunnel--reference--group-001.md#canonical-c17d15da48d2a51c9245919527f088405aaf5074b3fa73e79db226f6ee1018b6) |
| `local_ip.intf.local_intf.namespace` | [local_ip.intf.local_intf.namespace](data-sources--tunnel--reference--group-001.md#canonical-de057c9e11d7e2544da6f97379948b4b5d057b79945c9718f9f064d4be0749dd) |
| `local_ip.intf.local_intf.tenant` | [local_ip.intf.local_intf.tenant](data-sources--tunnel--reference--group-001.md#canonical-814e0929a58b96fdd3ce35e8da74f7ed31ca5075a50425c845b79f14cc52c991) |
| `local_ip.intf.local_intf.uid` | [local_ip.intf.local_intf.uid](data-sources--tunnel--reference--group-001.md#canonical-b370a4a4d5d56dca6e1d04a61929a378ffa2e4e66f287744fcc56aca36b1b9ac) |
| `local_ip.ip_address` | [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-457e42c308dd721289a5d970f7c238e4c3ce61dce33a3743ced3e6e18607d301) |
| `local_ip.ip_address.auto` | [local_ip.ip_address.auto](data-sources--tunnel--reference--group-001.md#canonical-f065d949e28effcbf0dc732e0a5a6bf9e0f88ed7f6b64a0919fab47dc33dd2de) |
| `local_ip.ip_address.ip_address` | [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-aeb66a88c8d9e4cab300c5e42f514e962160d07bad6b04eb1eac2a44185123c9) |
| `local_ip.ip_address.ip_address.dual_stack` | [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-d53a937e5448a7c6f14aaee267117b674133028732fd14bd6c0f392099c9d035) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4` | [local_ip.ip_address.ip_address.dual_stack.ipv4](data-sources--tunnel--reference--group-001.md#canonical-02dd41ed074f7c0401c74c232a7139e26451ba7237035a3df2319b87504d468f) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-c7224b3f4043f3d76e859e2a27b6a6c5822806f0a8e8d1f4d404898988066cc9) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6` | [local_ip.ip_address.ip_address.dual_stack.ipv6](data-sources--tunnel--reference--group-001.md#canonical-8777401ef071fa507d8745a2a6f05a8052d6a318ddb44ec319df477e562f938c) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-83e866a6094afa1e5f88e8294d9200927dbe860afb185916633b3f32d5d6cd1d) |
| `local_ip.ip_address.ip_address.ipv4` | [local_ip.ip_address.ip_address.ipv4](data-sources--tunnel--reference--group-001.md#canonical-19193575f0508dc0eb5db3b80c45ce58d3139af83d0a96f9bdf83485178b486f) |
| `local_ip.ip_address.ip_address.ipv4.addr` | [local_ip.ip_address.ip_address.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-fa6b573acd31c75823292dd039a64181218810ffb7c94cf75a9e6552081efb58) |
| `local_ip.ip_address.ip_address.ipv6` | [local_ip.ip_address.ip_address.ipv6](data-sources--tunnel--reference--group-001.md#canonical-7708e91430f874ec26587829d366b74aecfe538d0f89934c748d97fce9710280) |
| `local_ip.ip_address.ip_address.ipv6.addr` | [local_ip.ip_address.ip_address.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-388330648884fd47554841ddc5bd91a0063260ddbc3667fc3f7eb15c11331a90) |
| `local_ip.ip_address.virtual_network_type` | [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-a2581433c74c4f5530784982d656b0e5c4191fba11f09364cf3774ca2fb99bec) |
| `local_ip.ip_address.virtual_network_type.public` | [local_ip.ip_address.virtual_network_type.public](data-sources--tunnel--reference--group-001.md#canonical-b351961f88421d860b7c32baa99b6ca3b1a299e89c04606591469495bc8903fa) |
| `local_ip.ip_address.virtual_network_type.site_local` | [local_ip.ip_address.virtual_network_type.site_local](data-sources--tunnel--reference--group-001.md#canonical-b7fdefa7501ee4699787be7132e424a92d425d0f4c6ee15a11539f3dfc3ef3f4) |
| `local_ip.ip_address.virtual_network_type.site_local_inside` | [local_ip.ip_address.virtual_network_type.site_local_inside](data-sources--tunnel--reference--group-001.md#canonical-e0412700a479502eebf51c92fe2303f293ebe3d0bc6f7c01917d1f8da9255f11) |
| `name` | [name](data-sources--tunnel--reference--group-001.md#canonical-774dc044926ef324538e8e289a9a45ef7c110dc7d36b1940228a8d7838eeb2f1) |
| `namespace` | [namespace](data-sources--tunnel--reference--group-001.md#canonical-feb1dea4311f58e1ea94f8310642ecf5207fc7e66234248169dbb6a966a89fe5) |
| `params` | [params](data-sources--tunnel--reference--group-001.md#canonical-276eb55911f744152522eeca6d51fb040033f68912c7a1b3704201354cc8cfbb) |
| `params.ipsec` | [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0a5dbe551b6a638d172f31ca5d72b0f2ee61c11ad8852d68373f3c5f2ff3b203) |
| `params.ipsec.ipsec_psk` | [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-3072e45defe96852aef770a5e67c8675627d46ef9d4a43e1f64bda9d29ef0cca) |
| `params.ipsec.ipsec_psk.blindfold_secret_info` | [params.ipsec.ipsec_psk.blindfold_secret_info](data-sources--tunnel--reference--group-001.md#canonical-c8911cd69a3eb5697fc0a774cfd06a23c625ba83e4b2265d619167b4a60fdf9d) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider](data-sources--tunnel--reference--group-001.md#canonical-a698cead90411a229049d96a3fe8cbd04e3bf0b00d648bc2498145bcdd66f37d) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.location` | [params.ipsec.ipsec_psk.blindfold_secret_info.location](data-sources--tunnel--reference--group-001.md#canonical-039c56970d816df16b0cfe448e915ab56832bb5f3acace3cdd6817a2989cf31b) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.store_provider](data-sources--tunnel--reference--group-001.md#canonical-4d635fb860477e1b490ea09348fc6dc24a4624042534a66bebda2ff8c5adb2e3) |
| `params.ipsec.ipsec_psk.clear_secret_info` | [params.ipsec.ipsec_psk.clear_secret_info](data-sources--tunnel--reference--group-001.md#canonical-f884f6e950470e0985a1420196375fdb39f92d6c286e5a572e7b9dbcd3c931d4) |
| `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` | [params.ipsec.ipsec_psk.clear_secret_info.provider_ref](data-sources--tunnel--reference--group-001.md#canonical-a0c0258e520f5fa926bf6319f766e866db9ec95d677e068211208bd0449f832b) |
| `params.ipsec.ipsec_psk.clear_secret_info.url` | [params.ipsec.ipsec_psk.clear_secret_info.url](data-sources--tunnel--reference--group-001.md#canonical-29877ceeba063acfea29ae6f7d8ffdeaf18f7917f30df50d6530ca7babeb294d) |
| `remote_ip` | [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-02a5367f71638ef96d936e8808e528505889812b9d456f8951d136329cfcdf40) |
| `remote_ip.endpoints` | [remote_ip.endpoints](data-sources--tunnel--reference--group-001.md#canonical-2d5c02d399a9c98678344b38366645a1359ed710dac88e2d559ca9c9ae3bdcbd) |
| `remote_ip.endpoints.endpoints` | [remote_ip.endpoints.endpoints](data-sources--tunnel--reference--group-001.md#canonical-6d822f8e73bfa0ebd00be833417b31ad2ca4f0eb43573847fc8bffe41f8a4e1f) |
| `remote_ip.ip` | [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-2fa17b42fb9af711287d3f50f3738b56a90785e33c74dadcfa5f5ebca6b489c1) |
| `remote_ip.ip.dual_stack` | [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-05797c15a32e3822b01ec39eda969a200ad279e6ac9ca54c2a4871bcbf04b091) |
| `remote_ip.ip.dual_stack.ipv4` | [remote_ip.ip.dual_stack.ipv4](data-sources--tunnel--reference--group-001.md#canonical-096eb01083a5c23f3fbf1cb529059cb8369fa96cfb0bde86a7364eca2bb993b8) |
| `remote_ip.ip.dual_stack.ipv4.addr` | [remote_ip.ip.dual_stack.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-afa96ed458053941559f0683debce71e43939b2b3ceee3c43cbf3cecce0c2995) |
| `remote_ip.ip.dual_stack.ipv6` | [remote_ip.ip.dual_stack.ipv6](data-sources--tunnel--reference--group-001.md#canonical-107de1005df9150acc5cc2716f3469ebf4fb69765adfbe4b89f46c921d80d10b) |
| `remote_ip.ip.dual_stack.ipv6.addr` | [remote_ip.ip.dual_stack.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-2f2d5280a68a26c847cc7f5e174e7086d20078ecaea717af6dd410f003474afe) |
| `remote_ip.ip.ipv4` | [remote_ip.ip.ipv4](data-sources--tunnel--reference--group-001.md#canonical-cdb5e4640784426a861598723f4a19d7537a5848a2bfefbe75d01e2e0b7c3b62) |
| `remote_ip.ip.ipv4.addr` | [remote_ip.ip.ipv4.addr](data-sources--tunnel--reference--group-001.md#canonical-f70edbb63ddef7bb3d3e64c2ba0e3267c68f1a7ba802b2c63b2c3b4bd0fc64ee) |
| `remote_ip.ip.ipv6` | [remote_ip.ip.ipv6](data-sources--tunnel--reference--group-001.md#canonical-40a257728a45439f1ff87aada0adb1355f86e1e547f8526af95a143b08b64cd3) |
| `remote_ip.ip.ipv6.addr` | [remote_ip.ip.ipv6.addr](data-sources--tunnel--reference--group-001.md#canonical-92ea45de2e8e94a05f9dda9b1d36a13364ff8f912bf2cc33a61d89df5fba9a75) |
| `tunnel_type` | [tunnel_type](data-sources--tunnel--reference--group-001.md#canonical-e68c239f6704b909862be98b8fcd40d485911adca9b3e783291f588182ea7c92) |

<a id="canonical-74f947d1ab2f1c417e0679dd2064af3fd62749bd100e503ee95a264e682caac9"></a>

## Next pages — Property reference / f7303f87cc5e / 12

- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [params](data-sources--tunnel--reference--group-001.md#canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f82ed1647582af2d788156cedbb72119d26dadc1cc2e505bafdf7dc285d9637"></a>

## local_ip — local_ip / 4d3debf39832 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- local_ip

<a id="canonical-97edccce064646a89b060c31597eb265b3631d9bfe01bbfbf0b29cd46b049359"></a>

Type: `"single"`. Computed.

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - 1. Local Interface - Network Interface from which IP address and network will be
selected 2. IP Address - IP address and network can be configured explicitly.

Upstream description:

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - &#8203;1. Local Interface - Network Interface from which IP address and network will
be selected &#8203;2. IP Address - IP address and network can be configured explicitly.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"intf\",\"ip_address\"]"
}
```

<a id="canonical-f0c60d1763152e7ba71aeb569e3c0a82461b1f79dff625c8c27931924fe9a1d7"></a>

## Direct properties — local_ip / 4d3debf39832 / 3

- [intf](data-sources--tunnel--reference--group-001.md#canonical-13b6ee0fbf61da8b91b7dd8588bef4171f07110d4c5250de46c54c3b76b21b17): complete subsection reference.

- [ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14): complete subsection reference.

<a id="canonical-a40c4341f70397fa8bf472db8665b374cb80ac9f233b7d9be0ae70e732da6a6e"></a>

## Next pages — local_ip / 4d3debf39832 / 4

- [local_ip.intf](data-sources--tunnel--reference--group-001.md#canonical-13b6ee0fbf61da8b91b7dd8588bef4171f07110d4c5250de46c54c3b76b21b17)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-13b6ee0fbf61da8b91b7dd8588bef4171f07110d4c5250de46c54c3b76b21b17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2717fe3a1d9dfd848a5d54e33142d6213bc39c7c1f4e805aca4fa7fe6ee8f650"></a>

## local_ip.intf — local_ip.intf / 97da62339ce1 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- local_ip.intf

<a id="canonical-0d9d6d0c4bc10557bff9a79200ae3e8c936d5f8a0493de7b7b9281e27d8f3965"></a>

Type: `"single"`. Computed.

Provides the local interface to pick up source IP and network for transporting encapsulated packet.

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

<a id="canonical-17938849b67505b0973740498574ad4c51a13c04d8228f7d5796253cf7e2f5bd"></a>

## Direct properties — local_ip.intf / 97da62339ce1 / 3

- [local_intf](data-sources--tunnel--reference--group-001.md#canonical-65fa02f80777413e0b25900e7d3daddbce6c18218976657765d63d7976649a94): complete subsection reference.

<a id="canonical-c98f9970cc11ec3fcbdf46489f868412a20fe57553a8702aa394f0333b285cb1"></a>

## Next pages — local_ip.intf / 97da62339ce1 / 4

- [local_ip.intf.local_intf](data-sources--tunnel--reference--group-001.md#canonical-65fa02f80777413e0b25900e7d3daddbce6c18218976657765d63d7976649a94)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-65fa02f80777413e0b25900e7d3daddbce6c18218976657765d63d7976649a94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28b3405519f84a91393715b289b4cdf4b1c4714b5be5832b00fb036be76f2375"></a>

## local_ip.intf.local_intf — local_ip.intf.local_intf / 82bb6c6485e1 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.intf](data-sources--tunnel--reference--group-001.md#canonical-13b6ee0fbf61da8b91b7dd8588bef4171f07110d4c5250de46c54c3b76b21b17)
- local_ip.intf.local_intf

<a id="canonical-ac0588a9159a3f358659236ad624fb80b5b42cb3fb4b55cd88f07d28b5f71394"></a>

Type: `"list"`. Computed.

Local interface to be used for filling in source information of IP and network for transport.

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

<a id="canonical-26dd5e65a2e6b459f553a66ff48f78e6921a09b7cfdbf2e9db65ac113470a6bd"></a>

## Direct properties — local_ip.intf.local_intf / 82bb6c6485e1 / 3

<a id="canonical-eac68472faaf99302019458a47a451e75875e62f410d887a25c7602e7f1a8d63"></a>

<a id="canonical-dd1054cceff1d14eaafe6907ff376bf1741306e7c47b2c1d5dc0fbeb6a3f84e7"></a>

## kind property — local_ip.intf.local_intf / 82bb6c6485e1 / 4

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

<a id="canonical-c17d15da48d2a51c9245919527f088405aaf5074b3fa73e79db226f6ee1018b6"></a>

<a id="canonical-f6a34ef6507850aee8ed213c74d2ded8a0e506a7b7136f7230b0391777614480"></a>

## name property — local_ip.intf.local_intf / 82bb6c6485e1 / 5

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

<a id="canonical-de057c9e11d7e2544da6f97379948b4b5d057b79945c9718f9f064d4be0749dd"></a>

<a id="canonical-75bb52573967a97ae879c14280174a3d748399be2c8af1b6329abe28039eece7"></a>

## namespace property — local_ip.intf.local_intf / 82bb6c6485e1 / 6

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

<a id="canonical-814e0929a58b96fdd3ce35e8da74f7ed31ca5075a50425c845b79f14cc52c991"></a>

<a id="canonical-d20c8dae4087560d17e571d273ef0f09f7c7dbe9d6dfc69e91e5ccb7b300a9b1"></a>

## tenant property — local_ip.intf.local_intf / 82bb6c6485e1 / 7

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

<a id="canonical-b370a4a4d5d56dca6e1d04a61929a378ffa2e4e66f287744fcc56aca36b1b9ac"></a>

<a id="canonical-90ac13be1ac5297c831a544da08f91811f470765d82969916f9e5dec901e969b"></a>

## uid property — local_ip.intf.local_intf / 82bb6c6485e1 / 8

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

<a id="canonical-7baaa346495617db71199d214fc4c960a81b3d974fd5a4e6104a4130d2d69362"></a>

## Next pages — local_ip.intf.local_intf / 82bb6c6485e1 / 9

- [local_ip.intf](data-sources--tunnel--reference--group-001.md#canonical-13b6ee0fbf61da8b91b7dd8588bef4171f07110d4c5250de46c54c3b76b21b17)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-600d54cd9d26c87314eeadf34346291a984165f64e532cf07cb5faf73979d085"></a>

## local_ip.ip_address — local_ip.ip_address / 1de9eba7e281 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- local_ip.ip_address

<a id="canonical-457e42c308dd721289a5d970f7c238e4c3ce61dce33a3743ced3e6e18607d301"></a>

Type: `"single"`. Computed.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"auto\",\"ip_address\"]"
}
```

<a id="canonical-b27272760843e9268e9fd4d554e88b55f03d37779f853531d4c082eeb05b30e0"></a>

## Direct properties — local_ip.ip_address / 1de9eba7e281 / 3

- [auto](data-sources--tunnel--reference--group-001.md#canonical-739aa444c473c43bd2a2ddf19df73ab3a3bcb9a392b0a26cfc16927a7464b640): complete subsection reference.

- [ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d): complete subsection reference.

- [virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c): complete subsection reference.

<a id="canonical-e036c9a7d0683e528023539910b1ba94ad1b56a20d74369197771a1c7dc3809a"></a>

## Next pages — local_ip.ip_address / 1de9eba7e281 / 4

- [local_ip.ip_address.auto](data-sources--tunnel--reference--group-001.md#canonical-739aa444c473c43bd2a2ddf19df73ab3a3bcb9a392b0a26cfc16927a7464b640)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-739aa444c473c43bd2a2ddf19df73ab3a3bcb9a392b0a26cfc16927a7464b640"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfed64bd2f355aa0dab0ca041a0fc5720e0c02d8cf9d7fde2d38ec4cb6c83e71"></a>

## local_ip.ip_address.auto — local_ip.ip_address.auto / 940978f67359 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- local_ip.ip_address.auto

<a id="canonical-f065d949e28effcbf0dc732e0a5a6bf9e0f88ed7f6b64a0919fab47dc33dd2de"></a>

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

<a id="canonical-0ae53d80d83da8f498adce7186498df189abcbeea312b4c7c88606241fce331f"></a>

## Direct properties — local_ip.ip_address.auto / 940978f67359 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57b6e18d238ac6e623f23bbd05ba0f9d0e2f94281cd323edafe14803b301656a"></a>

## Next pages — local_ip.ip_address.auto / 940978f67359 / 4

- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f13debc8ed4fdb1662acdeebfbb0e1e337c468e0bff6f4794bb780eac6824e8"></a>

## local_ip.ip_address.ip_address — local_ip.ip_address.ip_address / 16713403038b / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- local_ip.ip_address.ip_address

<a id="canonical-aeb66a88c8d9e4cab300c5e42f514e962160d07bad6b04eb1eac2a44185123c9"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-d2eb1003dcd1eac9f7c7c8e89919cb9906cc476ee71a70266f21229177e35add"></a>

## Direct properties — local_ip.ip_address.ip_address / 16713403038b / 3

- [dual_stack](data-sources--tunnel--reference--group-001.md#canonical-c372de488c6c7de5b20ccf16b7369bb53d8a7b531b81af10223dd0f537bb56df): complete subsection reference.

- [ipv4](data-sources--tunnel--reference--group-001.md#canonical-0e55a8e3314b676ef867601b41a475272b6b765bda2da6f6fbbcd6456a0429c9): complete subsection reference.

- [ipv6](data-sources--tunnel--reference--group-001.md#canonical-c4b63b296e8d3c45c113d0c53875bcb2b5eb4f7cf36d858182b6e7c7acea7365): complete subsection reference.

<a id="canonical-ecd09edda0417ec25edfe15e5fb2c4b082ce8bf52f3e61e5985059c9e40fcc53"></a>

## Next pages — local_ip.ip_address.ip_address / 16713403038b / 4

- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-c372de488c6c7de5b20ccf16b7369bb53d8a7b531b81af10223dd0f537bb56df)
- [local_ip.ip_address.ip_address.ipv4](data-sources--tunnel--reference--group-001.md#canonical-0e55a8e3314b676ef867601b41a475272b6b765bda2da6f6fbbcd6456a0429c9)
- [local_ip.ip_address.ip_address.ipv6](data-sources--tunnel--reference--group-001.md#canonical-c4b63b296e8d3c45c113d0c53875bcb2b5eb4f7cf36d858182b6e7c7acea7365)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-c372de488c6c7de5b20ccf16b7369bb53d8a7b531b81af10223dd0f537bb56df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7740362ae0597ab77ebd822177f90abcc875ffae3c1e40b88b6377f1ad3f689"></a>

## local_ip.ip_address.ip_address.dual_stack — local_ip.ip_address.ip_address.dual_stack / 0209d565bb9c / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- local_ip.ip_address.ip_address.dual_stack

<a id="canonical-d53a937e5448a7c6f14aaee267117b674133028732fd14bd6c0f392099c9d035"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

<a id="canonical-d54e1e66e174acd71d30344ab7db685f37e1720d403dec27bc0e5aa742767b91"></a>

## Direct properties — local_ip.ip_address.ip_address.dual_stack / 0209d565bb9c / 3

- [ipv4](data-sources--tunnel--reference--group-001.md#canonical-68550169060043bc4321f5a4a0d179e77816468e3af0ddde6454db3889d55228): complete subsection reference.

- [ipv6](data-sources--tunnel--reference--group-001.md#canonical-da8f9f26cb21845b0fed7ca1f3c5e5ae63bbb69f1592511e7238cd34939e58e2): complete subsection reference.

<a id="canonical-605e3a2db6aa25e779c0c9607c20d98e7218935fa7c95abd8f9b5ff8535cd343"></a>

## Next pages — local_ip.ip_address.ip_address.dual_stack / 0209d565bb9c / 4

- [local_ip.ip_address.ip_address.dual_stack.ipv4](data-sources--tunnel--reference--group-001.md#canonical-68550169060043bc4321f5a4a0d179e77816468e3af0ddde6454db3889d55228)
- [local_ip.ip_address.ip_address.dual_stack.ipv6](data-sources--tunnel--reference--group-001.md#canonical-da8f9f26cb21845b0fed7ca1f3c5e5ae63bbb69f1592511e7238cd34939e58e2)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-68550169060043bc4321f5a4a0d179e77816468e3af0ddde6454db3889d55228"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01fa777855935aeb5e7b5357c8fe3c6425c04ccbf869c5b8a1e3ebc58cc369d1"></a>

## local_ip.ip_address.ip_address.dual_stack.ipv4 — local_ip.ip_address.ip_address.dual_stack.ipv4 / 55ba0eeec931 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-c372de488c6c7de5b20ccf16b7369bb53d8a7b531b81af10223dd0f537bb56df)
- local_ip.ip_address.ip_address.dual_stack.ipv4

<a id="canonical-02dd41ed074f7c0401c74c232a7139e26451ba7237035a3df2319b87504d468f"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-520e0c88072b641f86cac0929d9c17d91fa071c1551b652d82d3bcc18b63b3a1"></a>

## Direct properties — local_ip.ip_address.ip_address.dual_stack.ipv4 / 55ba0eeec931 / 3

<a id="canonical-c7224b3f4043f3d76e859e2a27b6a6c5822806f0a8e8d1f4d404898988066cc9"></a>

<a id="canonical-bd3367a93ba190fc5140439a161c404c99b2b9bb3743f428741d568138b4e9d3"></a>

## addr property — local_ip.ip_address.ip_address.dual_stack.ipv4 / 55ba0eeec931 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-ecbc0aa78cdcdf075f36fca94b936b23aa2b945bfe9c433cff5d155ac0145e10"></a>

## Next pages — local_ip.ip_address.ip_address.dual_stack.ipv4 / 55ba0eeec931 / 5

- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-c372de488c6c7de5b20ccf16b7369bb53d8a7b531b81af10223dd0f537bb56df)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-da8f9f26cb21845b0fed7ca1f3c5e5ae63bbb69f1592511e7238cd34939e58e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08814b4308130a4b77d8d025f4fd441706b7701d9b03d8ecbd0853faedf1fc02"></a>

## local_ip.ip_address.ip_address.dual_stack.ipv6 — local_ip.ip_address.ip_address.dual_stack.ipv6 / baaa799ec36a / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-c372de488c6c7de5b20ccf16b7369bb53d8a7b531b81af10223dd0f537bb56df)
- local_ip.ip_address.ip_address.dual_stack.ipv6

<a id="canonical-8777401ef071fa507d8745a2a6f05a8052d6a318ddb44ec319df477e562f938c"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-0d2c9f36a0230a6c2ccb76a4f239454fdaf7b839e8bd7ee8a5c68f16e8792127"></a>

## Direct properties — local_ip.ip_address.ip_address.dual_stack.ipv6 / baaa799ec36a / 3

<a id="canonical-83e866a6094afa1e5f88e8294d9200927dbe860afb185916633b3f32d5d6cd1d"></a>

<a id="canonical-ddb7997dfe7334c2a11f0dda41feea58948e07ab5b8a60e68e492b51e3e5dcfc"></a>

## addr property — local_ip.ip_address.ip_address.dual_stack.ipv6 / baaa799ec36a / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-40a7e99ac22f34dedffb9f09653503a594a999fa1a5b3c918d754830c12c9fbf"></a>

## Next pages — local_ip.ip_address.ip_address.dual_stack.ipv6 / baaa799ec36a / 5

- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-c372de488c6c7de5b20ccf16b7369bb53d8a7b531b81af10223dd0f537bb56df)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-0e55a8e3314b676ef867601b41a475272b6b765bda2da6f6fbbcd6456a0429c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98c4d83e51db7ccc6f93c5adcc4c30394f2a8e2be2e094d886c4490e5aaf453b"></a>

## local_ip.ip_address.ip_address.ipv4 — local_ip.ip_address.ip_address.ipv4 / 24b0cb136993 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- local_ip.ip_address.ip_address.ipv4

<a id="canonical-19193575f0508dc0eb5db3b80c45ce58d3139af83d0a96f9bdf83485178b486f"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-05397052e272f73c17495545dc0dd8e4b53bf8b071d3284b2e47d50cc73b058f"></a>

## Direct properties — local_ip.ip_address.ip_address.ipv4 / 24b0cb136993 / 3

<a id="canonical-fa6b573acd31c75823292dd039a64181218810ffb7c94cf75a9e6552081efb58"></a>

<a id="canonical-5ba23a381dd3f8e44f081048b1bf057c6c99d0f80a820744eb0d1d54cdc26863"></a>

## addr property — local_ip.ip_address.ip_address.ipv4 / 24b0cb136993 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-40635f28d73a790eb498d05fccb18055c3f0b97949c4ea6e00438f81d6ea3e8a"></a>

## Next pages — local_ip.ip_address.ip_address.ipv4 / 24b0cb136993 / 5

- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-c4b63b296e8d3c45c113d0c53875bcb2b5eb4f7cf36d858182b6e7c7acea7365"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8c5a0d906bb3c3151a2e7f1fc3b2566ac0acd0533b2d06c6fd799e79565e790"></a>

## local_ip.ip_address.ip_address.ipv6 — local_ip.ip_address.ip_address.ipv6 / 7506be522f2f / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- local_ip.ip_address.ip_address.ipv6

<a id="canonical-7708e91430f874ec26587829d366b74aecfe538d0f89934c748d97fce9710280"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-84ef9491e89416bbf7d9b4cfe1f49aca85b1b8b83649a79934c000527c6fd36b"></a>

## Direct properties — local_ip.ip_address.ip_address.ipv6 / 7506be522f2f / 3

<a id="canonical-388330648884fd47554841ddc5bd91a0063260ddbc3667fc3f7eb15c11331a90"></a>

<a id="canonical-20540824775a2c505ddb65b199d41a6a3c30041be0809d68ca6e27ecb044c3d0"></a>

## addr property — local_ip.ip_address.ip_address.ipv6 / 7506be522f2f / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-6f4a82bfb29bc4c720c31ec238fa03ca8a599594abc6dd912b0327c89b5bebd2"></a>

## Next pages — local_ip.ip_address.ip_address.ipv6 / 7506be522f2f / 5

- [local_ip.ip_address.ip_address](data-sources--tunnel--reference--group-001.md#canonical-38c1975bef68cfbba9ba411b4e3232e3503f630b94d9b634e1a112a6f461801d)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42175f8c80fd8e46007bec83592927d0f2649ed40172130240cd4c41a7a0c587"></a>

## local_ip.ip_address.virtual_network_type — local_ip.ip_address.virtual_network_type / 6382f9934034 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- local_ip.ip_address.virtual_network_type

<a id="canonical-a2581433c74c4f5530784982d656b0e5c4191fba11f09364cf3774ca2fb99bec"></a>

Type: `"single"`. Computed.

Different types of virtual networks understood by the system.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vn_type_choice": "[\"public\",\"site_local\",\"site_local_inside\"]"
}
```

<a id="canonical-7ed6bb33e9a40376ef3242c544ee50d0c5ae5f9d3e52aca36791d6072579f4c8"></a>

## Direct properties — local_ip.ip_address.virtual_network_type / 6382f9934034 / 3

- [public](data-sources--tunnel--reference--group-001.md#canonical-45f755b030ad68602d6898b29c3b64e9c6383df59e2e6cd9c73c26a2222399fe): complete subsection reference.

- [site_local](data-sources--tunnel--reference--group-001.md#canonical-3c7366ab64d862a6b63109155f494d00edbf16e8bf30447e48b3418534816f67): complete subsection reference.

- [site_local_inside](data-sources--tunnel--reference--group-001.md#canonical-82d35872af11be6445313ca3938a789d758afae9d037061658f1399a3909d688): complete subsection reference.

<a id="canonical-05c2178e7ebedf59bf8a11b4d09968ccfb985f450de4eb59cfb092eb8ce3d270"></a>

## Next pages — local_ip.ip_address.virtual_network_type / 6382f9934034 / 4

- [local_ip.ip_address.virtual_network_type.public](data-sources--tunnel--reference--group-001.md#canonical-45f755b030ad68602d6898b29c3b64e9c6383df59e2e6cd9c73c26a2222399fe)
- [local_ip.ip_address.virtual_network_type.site_local](data-sources--tunnel--reference--group-001.md#canonical-3c7366ab64d862a6b63109155f494d00edbf16e8bf30447e48b3418534816f67)
- [local_ip.ip_address.virtual_network_type.site_local_inside](data-sources--tunnel--reference--group-001.md#canonical-82d35872af11be6445313ca3938a789d758afae9d037061658f1399a3909d688)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-45f755b030ad68602d6898b29c3b64e9c6383df59e2e6cd9c73c26a2222399fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0faec95a719b199aa5df32acaee1393dee9db8052edda71ca66f281ef63f68f9"></a>

## local_ip.ip_address.virtual_network_type.public — local_ip.ip_address.virtual_network_type.public / 08eaeb185fe7 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c)
- local_ip.ip_address.virtual_network_type.public

<a id="canonical-b351961f88421d860b7c32baa99b6ca3b1a299e89c04606591469495bc8903fa"></a>

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

<a id="canonical-6629e82922f8b600db332a5761a3c1f406e7335963c989af0a06139e3c709f7f"></a>

## Direct properties — local_ip.ip_address.virtual_network_type.public / 08eaeb185fe7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39c33ff44f6d21c72b6efcfb87d852b78f6ff2d7bcf95ddfc2c101f82c04ddcc"></a>

## Next pages — local_ip.ip_address.virtual_network_type.public / 08eaeb185fe7 / 4

- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-3c7366ab64d862a6b63109155f494d00edbf16e8bf30447e48b3418534816f67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-944021d8c06bc26d4ee8ea37df03b335c3cff4d295ea01a17302aef1cf4fb11b"></a>

## local_ip.ip_address.virtual_network_type.site_local — local_ip.ip_address.virtual_network_type.site_local / ce73ad5c4be7 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c)
- local_ip.ip_address.virtual_network_type.site_local

<a id="canonical-b7fdefa7501ee4699787be7132e424a92d425d0f4c6ee15a11539f3dfc3ef3f4"></a>

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

<a id="canonical-92440066bdcd65ac9856c55797bf068008d545948f1e89ef76028760596230ea"></a>

## Direct properties — local_ip.ip_address.virtual_network_type.site_local / ce73ad5c4be7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46d2773c80e54e7c2cd0c57dd35653c336fad73a4a5cc48c7a6074236ec2a449"></a>

## Next pages — local_ip.ip_address.virtual_network_type.site_local / ce73ad5c4be7 / 4

- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-82d35872af11be6445313ca3938a789d758afae9d037061658f1399a3909d688"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-466944d0f6e9cc482fed0e3c950ed6f92b736a635f228126401ee0a2988530e1"></a>

## local_ip.ip_address.virtual_network_type.site_local_inside — local_ip.ip_address.virtual_network_type.site_local_inside / 5b7761d20eb1 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [local_ip](data-sources--tunnel--reference--group-001.md#canonical-192d4b9b34165d95ce406423076ee0a197879b897c636a75a3e92d3b74b05cd6)
- [local_ip.ip_address](data-sources--tunnel--reference--group-001.md#canonical-28ec14bd27d4cf1595ba79154f63c34025f25af67de1ce883cc351e9204e0e14)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c)
- local_ip.ip_address.virtual_network_type.site_local_inside

<a id="canonical-e0412700a479502eebf51c92fe2303f293ebe3d0bc6f7c01917d1f8da9255f11"></a>

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

<a id="canonical-1229bdafa5cc6d9fe8a654da569bcdd87f027940189acdf37602e631e9442149"></a>

## Direct properties — local_ip.ip_address.virtual_network_type.site_local_inside / 5b7761d20eb1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64436bc65d53942cdd98bcfb12a1bba4f8152a12820cd38c07231d904166dc74"></a>

## Next pages — local_ip.ip_address.virtual_network_type.site_local_inside / 5b7761d20eb1 / 4

- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--reference--group-001.md#canonical-fc027401b483e6848ddb6bd4bab7d1423bd4aebf5b34544100132e4c2965760c)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5caa0449c4a47e52274835fb66ddd28d2e828df3a898c6ed73cd363ef45806f8"></a>

## params — params / cf0deadb35d9 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- params

<a id="canonical-276eb55911f744152522eeca6d51fb040033f68912c7a1b3704201354cc8cfbb"></a>

Type: `"single"`. Computed.

Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which
PSK can be configured.

Upstream description:

Tunnel configuration parameters for supported encapsulation &#8203;1. IPsec is supported with PSK
for which PSK can be configured.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"ipsec\"]"
}
```

<a id="canonical-743428b2db8ba76333bcb60afcf7e8eaaa4ff7d80e54849f03c00199d00abde7"></a>

## Direct properties — params / cf0deadb35d9 / 3

- [ipsec](data-sources--tunnel--reference--group-001.md#canonical-0cee72aeb6bf8429a85a9c72044b8aa0321f084a1f9e356477ee65cc52e3cf61): complete subsection reference.

<a id="canonical-9974d92dcd376307fdc52c1115353c8f15a6d4ee3e9406ae7abd46a12f7da1a4"></a>

## Next pages — params / cf0deadb35d9 / 4

- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0cee72aeb6bf8429a85a9c72044b8aa0321f084a1f9e356477ee65cc52e3cf61)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-0cee72aeb6bf8429a85a9c72044b8aa0321f084a1f9e356477ee65cc52e3cf61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-659acc9b32eea69d1d28b56ed5040f0b4eaf1ee6a5781e2da27e9eb8ac158cbb"></a>

## params.ipsec — params.ipsec / 4993cdfcb1c1 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [params](data-sources--tunnel--reference--group-001.md#canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7)
- params.ipsec

<a id="canonical-0a5dbe551b6a638d172f31ca5d72b0f2ee61c11ad8852d68373f3c5f2ff3b203"></a>

Type: `"single"`. Computed.

Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.

Upstream description:

Configuration for IPsec encapsulation are: &#8203;1. PSK - pre shared key to be used by IKE.

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

<a id="canonical-722a20e687b03506cda34543ebcb52cbbffd690eba09e5e4adcb8a022bce92fe"></a>

## Direct properties — params.ipsec / 4993cdfcb1c1 / 3

- [ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-4a914378f202cb1a296a823af53db901510b323cd94b9f7ac703ea7a5bf36d49): complete subsection reference.

<a id="canonical-6b6fa8b488e03a0d0c4f3c7ee8a8658da457f782fc6785e7a9f65875e725b063"></a>

## Next pages — params.ipsec / 4993cdfcb1c1 / 4

- [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-4a914378f202cb1a296a823af53db901510b323cd94b9f7ac703ea7a5bf36d49)
- [params](data-sources--tunnel--reference--group-001.md#canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-4a914378f202cb1a296a823af53db901510b323cd94b9f7ac703ea7a5bf36d49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b942ed6ad92caac7673d04b9a0fdf67cbf8dc2041cb777ddcfb33f3aabb34be1"></a>

## params.ipsec.ipsec_psk — params.ipsec.ipsec_psk / 23cff15e41e9 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [params](data-sources--tunnel--reference--group-001.md#canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7)
- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0cee72aeb6bf8429a85a9c72044b8aa0321f084a1f9e356477ee65cc52e3cf61)
- params.ipsec.ipsec_psk

<a id="canonical-3072e45defe96852aef770a5e67c8675627d46ef9d4a43e1f64bda9d29ef0cca"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1df649ebe1031d718b85444cf6ceab20c031b5a13ff7a7f01caa68624403489f"></a>

## Direct properties — params.ipsec.ipsec_psk / 23cff15e41e9 / 3

- [blindfold_secret_info](data-sources--tunnel--reference--group-001.md#canonical-0c2a531f57f040af62c7a4a78a5e7e0ca615b569cb5f2c82ebc5c2d8e3f04b1b): complete subsection reference.

- [clear_secret_info](data-sources--tunnel--reference--group-001.md#canonical-7f348478e76c5e7acaaca771bd7c6d069fda99b984d38bea7e84eca40a4a9b21): complete subsection reference.

<a id="canonical-eb03ade9cacad0275707c2f896512101f9274456ac41c1ad444e2bb3602dbd34"></a>

## Next pages — params.ipsec.ipsec_psk / 23cff15e41e9 / 4

- [params.ipsec.ipsec_psk.blindfold_secret_info](data-sources--tunnel--reference--group-001.md#canonical-0c2a531f57f040af62c7a4a78a5e7e0ca615b569cb5f2c82ebc5c2d8e3f04b1b)
- [params.ipsec.ipsec_psk.clear_secret_info](data-sources--tunnel--reference--group-001.md#canonical-7f348478e76c5e7acaaca771bd7c6d069fda99b984d38bea7e84eca40a4a9b21)
- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0cee72aeb6bf8429a85a9c72044b8aa0321f084a1f9e356477ee65cc52e3cf61)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-0c2a531f57f040af62c7a4a78a5e7e0ca615b569cb5f2c82ebc5c2d8e3f04b1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92bf4128dafc93769e6f08d67d90083d5f2e0c10ca0f773e8ee21ca2b8a1c863"></a>

## params.ipsec.ipsec_psk.blindfold_secret_info — params.ipsec.ipsec_psk.blindfold_secret_info / 608350013063 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [params](data-sources--tunnel--reference--group-001.md#canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7)
- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0cee72aeb6bf8429a85a9c72044b8aa0321f084a1f9e356477ee65cc52e3cf61)
- [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-4a914378f202cb1a296a823af53db901510b323cd94b9f7ac703ea7a5bf36d49)
- params.ipsec.ipsec_psk.blindfold_secret_info

<a id="canonical-c8911cd69a3eb5697fc0a774cfd06a23c625ba83e4b2265d619167b4a60fdf9d"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-65209c3c30d9f476e8adbf42b0d82150f7d580d4901cb83ee3baafad9f56e31c"></a>

## Direct properties — params.ipsec.ipsec_psk.blindfold_secret_info / 608350013063 / 3

<a id="canonical-a698cead90411a229049d96a3fe8cbd04e3bf0b00d648bc2498145bcdd66f37d"></a>

<a id="canonical-18d7a1ea6c6fcba19a5ce31b2e44da4f55e999772436bc5d51c42ca6918c2a49"></a>

## decryption_provider property — params.ipsec.ipsec_psk.blindfold_secret_info / 608350013063 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-039c56970d816df16b0cfe448e915ab56832bb5f3acace3cdd6817a2989cf31b"></a>

<a id="canonical-ea058245d6ea34d9cabb1aafc6818e118e8909ac43817db7f844b72543126049"></a>

## location property — params.ipsec.ipsec_psk.blindfold_secret_info / 608350013063 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-4d635fb860477e1b490ea09348fc6dc24a4624042534a66bebda2ff8c5adb2e3"></a>

<a id="canonical-09f29b9edb18fb8ba226a10eb01ee66a71adaed2af8ffde726cc3a8bbe634e15"></a>

## store_provider property — params.ipsec.ipsec_psk.blindfold_secret_info / 608350013063 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-a8235f45c31f21e43d85826f86697d4c3677723958ed54f6895c488d3e596508"></a>

## Next pages — params.ipsec.ipsec_psk.blindfold_secret_info / 608350013063 / 7

- [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-4a914378f202cb1a296a823af53db901510b323cd94b9f7ac703ea7a5bf36d49)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-7f348478e76c5e7acaaca771bd7c6d069fda99b984d38bea7e84eca40a4a9b21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-134c203e56ad8d9136c074923a6418da153734498986509599a57d2f966be6f6"></a>

## params.ipsec.ipsec_psk.clear_secret_info — params.ipsec.ipsec_psk.clear_secret_info / ed06d9f707c4 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [params](data-sources--tunnel--reference--group-001.md#canonical-f8c9f3a088304f9aa2b70f4c6bdedf5b32eb3d3917441e4e63db13eecf1eedb7)
- [params.ipsec](data-sources--tunnel--reference--group-001.md#canonical-0cee72aeb6bf8429a85a9c72044b8aa0321f084a1f9e356477ee65cc52e3cf61)
- [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-4a914378f202cb1a296a823af53db901510b323cd94b9f7ac703ea7a5bf36d49)
- params.ipsec.ipsec_psk.clear_secret_info

<a id="canonical-f884f6e950470e0985a1420196375fdb39f92d6c286e5a572e7b9dbcd3c931d4"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-9a1f87846754baac710d70ac08f238bded6e60529ae371fc5f12c62e480ad917"></a>

## Direct properties — params.ipsec.ipsec_psk.clear_secret_info / ed06d9f707c4 / 3

<a id="canonical-a0c0258e520f5fa926bf6319f766e866db9ec95d677e068211208bd0449f832b"></a>

<a id="canonical-5b604c031c33babb4f1abeec90baed57fc9db41fa7c0b22a6c9ca551c258d2da"></a>

## provider_ref property — params.ipsec.ipsec_psk.clear_secret_info / ed06d9f707c4 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-29877ceeba063acfea29ae6f7d8ffdeaf18f7917f30df50d6530ca7babeb294d"></a>

<a id="canonical-31c3613d50333d0e42d6896a7444e629e6cc6bb1b4f18215b15519937923b376"></a>

## url property — params.ipsec.ipsec_psk.clear_secret_info / ed06d9f707c4 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-6f255d6939c4920f8657399846b486d9b8eabd16d5d0edccba2bdb030bb6065d"></a>

## Next pages — params.ipsec.ipsec_psk.clear_secret_info / ed06d9f707c4 / 6

- [params.ipsec.ipsec_psk](data-sources--tunnel--reference--group-001.md#canonical-4a914378f202cb1a296a823af53db901510b323cd94b9f7ac703ea7a5bf36d49)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-704e0c55474635578898401d6f88a98d8d981a0fb797b3684b7ece521c8eb49b"></a>

## remote_ip — remote_ip / dcbcd9d59c17 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- remote_ip

<a id="canonical-02a5367f71638ef96d936e8808e528505889812b9d456f8951d136329cfcdf40"></a>

Type: `"single"`. Computed.

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP
Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map
of IP address on per ver node basis.

Upstream description:

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - &#8203;1.
IP Address - Specifies the remote IP to which tunnel has to be connected &#8203;2. Remote endpoint -
Is a map of IP address on per ver node basis.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"endpoints\",\"ip\"]"
}
```

<a id="canonical-09776cb2bbc9d37278b25326f57c483e75ab31f8b92cee8cf1779b94e18841c1"></a>

## Direct properties — remote_ip / dcbcd9d59c17 / 3

- [endpoints](data-sources--tunnel--reference--group-001.md#canonical-c30431ad88374fcf585cb33e9d9e840ffda3cea27ea511cc5992c4b83ef4f6f7): complete subsection reference.

- [ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be): complete subsection reference.

<a id="canonical-0a453c978db200eabae8333e4a343a6b0be90f9ef37dc02a6f1338102c5bcc82"></a>

## Next pages — remote_ip / dcbcd9d59c17 / 4

- [remote_ip.endpoints](data-sources--tunnel--reference--group-001.md#canonical-c30431ad88374fcf585cb33e9d9e840ffda3cea27ea511cc5992c4b83ef4f6f7)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-c30431ad88374fcf585cb33e9d9e840ffda3cea27ea511cc5992c4b83ef4f6f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a66c295927a182657950237dbe854b09b3a0010ef92005c2a8c5c3d5fec753d"></a>

## remote_ip.endpoints — remote_ip.endpoints / 0dd41990ab82 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- remote_ip.endpoints

<a id="canonical-2d5c02d399a9c98678344b38366645a1359ed710dac88e2d559ca9c9ae3bdcbd"></a>

Type: `"single"`. Computed.

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

Upstream description:

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

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

<a id="canonical-12f1097463fbd2ec0db5ab23de9509c06cc9b1649bc494ecd117bbffd9269c9a"></a>

## Direct properties — remote_ip.endpoints / 0dd41990ab82 / 3

- [endpoints](data-sources--tunnel--reference--group-001.md#canonical-012719fe14e6f30fc38e36aa69c30f30f9364074f35e728f59eed2f83cfe4f0f): complete subsection reference.

<a id="canonical-cc6b3cbbf0fd7716aeadd8c9cd0cebc39b017e48446692bd8bc6913198e502c8"></a>

## Next pages — remote_ip.endpoints / 0dd41990ab82 / 4

- [remote_ip.endpoints.endpoints](data-sources--tunnel--reference--group-001.md#canonical-012719fe14e6f30fc38e36aa69c30f30f9364074f35e728f59eed2f83cfe4f0f)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-012719fe14e6f30fc38e36aa69c30f30f9364074f35e728f59eed2f83cfe4f0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c69aebc6766d552088a70285837cc1fa2ebfcb6150ded8cd0e2410e8b1140e9a"></a>

## remote_ip.endpoints.endpoints — remote_ip.endpoints.endpoints / f41d8de7f206 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [remote_ip.endpoints](data-sources--tunnel--reference--group-001.md#canonical-c30431ad88374fcf585cb33e9d9e840ffda3cea27ea511cc5992c4b83ef4f6f7)
- remote_ip.endpoints.endpoints

<a id="canonical-6d822f8e73bfa0ebd00be833417b31ad2ca4f0eb43573847fc8bffe41f8a4e1f"></a>

Type: `"single"`. Computed.

Map of remote attributes to which tunnel will be established on per site node basis Every node can
have a different attributes and IP address to connect to Key is ver node name and value is Remote
node attributes.

Upstream description:

Map of remote attributes to which tunnel will be established on per site node basis Every node can
have a different attributes and IP address to connect to Key is ver node name and value is Remote
node attributes.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-738e1324f4e056c8c2656cdd3bfee33e760a8b92c6283fdc3e75e6d40dc29f35"></a>

## Direct properties — remote_ip.endpoints.endpoints / f41d8de7f206 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f36b8b60cf1c43fe8a5d22292eaecb74ff95389fcd256b897d93e505194fe980"></a>

## Next pages — remote_ip.endpoints.endpoints / f41d8de7f206 / 4

- [remote_ip.endpoints](data-sources--tunnel--reference--group-001.md#canonical-c30431ad88374fcf585cb33e9d9e840ffda3cea27ea511cc5992c4b83ef4f6f7)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dd7d87023ec969c718569d37f4bf9cddfcea2b4ce969c9a22e142431fd99927"></a>

## remote_ip.ip — remote_ip.ip / ef6eea23467b / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- remote_ip.ip

<a id="canonical-2fa17b42fb9af711287d3f50f3738b56a90785e33c74dadcfa5f5ebca6b489c1"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-213b4bcff7d14acb29ded271b9bf37f1e2dd64b9863a8ded39b2fca0eee6bd06"></a>

## Direct properties — remote_ip.ip / ef6eea23467b / 3

- [dual_stack](data-sources--tunnel--reference--group-001.md#canonical-5855512a836c624be15a2d56b6bf15f30a97d11b3b0c791cf661b0eba1c26359): complete subsection reference.

- [ipv4](data-sources--tunnel--reference--group-001.md#canonical-d4be86d8a15de3e3f16137c98065fa475d414a6eb0b795d863612dc4c7f5152b): complete subsection reference.

- [ipv6](data-sources--tunnel--reference--group-001.md#canonical-4aa9930fab654f04d8cbf02da4608afdaaa031042d480d77a7879ecb18c50b11): complete subsection reference.

<a id="canonical-35f93738801f6c15a7dbd1647cadd6ff8fabd89d6fefecc7919ea0931754e653"></a>

## Next pages — remote_ip.ip / ef6eea23467b / 4

- [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-5855512a836c624be15a2d56b6bf15f30a97d11b3b0c791cf661b0eba1c26359)
- [remote_ip.ip.ipv4](data-sources--tunnel--reference--group-001.md#canonical-d4be86d8a15de3e3f16137c98065fa475d414a6eb0b795d863612dc4c7f5152b)
- [remote_ip.ip.ipv6](data-sources--tunnel--reference--group-001.md#canonical-4aa9930fab654f04d8cbf02da4608afdaaa031042d480d77a7879ecb18c50b11)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-5855512a836c624be15a2d56b6bf15f30a97d11b3b0c791cf661b0eba1c26359"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f676e8b54d0eb4e132e0e40d4d3466cb0192f985bb3e64ada88377aa448d14ab"></a>

## remote_ip.ip.dual_stack — remote_ip.ip.dual_stack / 7f3525c82977 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- remote_ip.ip.dual_stack

<a id="canonical-05797c15a32e3822b01ec39eda969a200ad279e6ac9ca54c2a4871bcbf04b091"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

<a id="canonical-72ecf289aedd7bef959cc334f00f7d9ed2092ec4e7214b053e369891e48a9e4a"></a>

## Direct properties — remote_ip.ip.dual_stack / 7f3525c82977 / 3

- [ipv4](data-sources--tunnel--reference--group-001.md#canonical-23b9d6317a9e36a91f62d1396f9ff181f474102cec5c5d5f248050b9a5efce58): complete subsection reference.

- [ipv6](data-sources--tunnel--reference--group-001.md#canonical-3e4a02ac30f28a5ed631deb4e0b1669fc93eb57164e17c3b8ffd617d93f1adfe): complete subsection reference.

<a id="canonical-024be355f1189bbac4949f870a22a0040c665e1907c5eadaf1203a3b8a8c49f5"></a>

## Next pages — remote_ip.ip.dual_stack / 7f3525c82977 / 4

- [remote_ip.ip.dual_stack.ipv4](data-sources--tunnel--reference--group-001.md#canonical-23b9d6317a9e36a91f62d1396f9ff181f474102cec5c5d5f248050b9a5efce58)
- [remote_ip.ip.dual_stack.ipv6](data-sources--tunnel--reference--group-001.md#canonical-3e4a02ac30f28a5ed631deb4e0b1669fc93eb57164e17c3b8ffd617d93f1adfe)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-23b9d6317a9e36a91f62d1396f9ff181f474102cec5c5d5f248050b9a5efce58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21999e9935afc0598aebef877d26b512deb3d033c1efcd9f410569f1b34c5ccf"></a>

## remote_ip.ip.dual_stack.ipv4 — remote_ip.ip.dual_stack.ipv4 / ce82034891f4 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-5855512a836c624be15a2d56b6bf15f30a97d11b3b0c791cf661b0eba1c26359)
- remote_ip.ip.dual_stack.ipv4

<a id="canonical-096eb01083a5c23f3fbf1cb529059cb8369fa96cfb0bde86a7364eca2bb993b8"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-c02802088cf8f982bf734acb14d17a0d73a58f3c82cf476f2403b1b2bd5cb958"></a>

## Direct properties — remote_ip.ip.dual_stack.ipv4 / ce82034891f4 / 3

<a id="canonical-afa96ed458053941559f0683debce71e43939b2b3ceee3c43cbf3cecce0c2995"></a>

<a id="canonical-a1d812e1b8c30f201c6b55c13b59a4c9d112ea4271a692471aaab861f32e3611"></a>

## addr property — remote_ip.ip.dual_stack.ipv4 / ce82034891f4 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3cd397b181c3f17bdf6d177b096766d755d1ffcb84ddafc1cf7278d05f56c684"></a>

## Next pages — remote_ip.ip.dual_stack.ipv4 / ce82034891f4 / 5

- [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-5855512a836c624be15a2d56b6bf15f30a97d11b3b0c791cf661b0eba1c26359)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-3e4a02ac30f28a5ed631deb4e0b1669fc93eb57164e17c3b8ffd617d93f1adfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59f816775a9b212c7984692b4b1d8907f57de2995fcd4c8478f0060ea56bcba9"></a>

## remote_ip.ip.dual_stack.ipv6 — remote_ip.ip.dual_stack.ipv6 / 086b527db3f8 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-5855512a836c624be15a2d56b6bf15f30a97d11b3b0c791cf661b0eba1c26359)
- remote_ip.ip.dual_stack.ipv6

<a id="canonical-107de1005df9150acc5cc2716f3469ebf4fb69765adfbe4b89f46c921d80d10b"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-369cd2bef93fb99f96f0f9eeb24376ff6574126db3f7486d2e6eb948908d9c66"></a>

## Direct properties — remote_ip.ip.dual_stack.ipv6 / 086b527db3f8 / 3

<a id="canonical-2f2d5280a68a26c847cc7f5e174e7086d20078ecaea717af6dd410f003474afe"></a>

<a id="canonical-295742b28c8d339717cefa4ff614acf4d99f7d51440b6c10a2bc77e8876fa30b"></a>

## addr property — remote_ip.ip.dual_stack.ipv6 / 086b527db3f8 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-ddec3b1122caeb5140646c926affc492bbe04f1d3b4b175515879d5f671a65b7"></a>

## Next pages — remote_ip.ip.dual_stack.ipv6 / 086b527db3f8 / 5

- [remote_ip.ip.dual_stack](data-sources--tunnel--reference--group-001.md#canonical-5855512a836c624be15a2d56b6bf15f30a97d11b3b0c791cf661b0eba1c26359)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-d4be86d8a15de3e3f16137c98065fa475d414a6eb0b795d863612dc4c7f5152b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ad71e7aeea669414c27dafcd9e7d2e565da21a305c8c594e2c10c37dd83d9af"></a>

## remote_ip.ip.ipv4 — remote_ip.ip.ipv4 / 16fe90b25f7f / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- remote_ip.ip.ipv4

<a id="canonical-cdb5e4640784426a861598723f4a19d7537a5848a2bfefbe75d01e2e0b7c3b62"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-61e882948d9473fc9f250f6ea4a69a56073c863c69078cd013a03a2315d6d5ed"></a>

## Direct properties — remote_ip.ip.ipv4 / 16fe90b25f7f / 3

<a id="canonical-f70edbb63ddef7bb3d3e64c2ba0e3267c68f1a7ba802b2c63b2c3b4bd0fc64ee"></a>

<a id="canonical-53a5cbc2113b227566b077180fb7d1c9e56aef610741d397f4910f311e486fd1"></a>

## addr property — remote_ip.ip.ipv4 / 16fe90b25f7f / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-49bce22e9389efa9552cb51ace8506e589bba493a7d09c887fff20b8c515986e"></a>

## Next pages — remote_ip.ip.ipv4 / 16fe90b25f7f / 5

- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-4aa9930fab654f04d8cbf02da4608afdaaa031042d480d77a7879ecb18c50b11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fe278748ed8e31b1a00a8a95fbd8edd2ef3cd5ef0c1a09f83c14f827135c059"></a>

## remote_ip.ip.ipv6 — remote_ip.ip.ipv6 / 9133a2f2595b / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Property reference](data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [remote_ip](data-sources--tunnel--reference--group-001.md#canonical-ded8bb06c750268dde569c61caf44eeb6ec7078daa89392dad684971b24f1fd8)
- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- remote_ip.ip.ipv6

<a id="canonical-40a257728a45439f1ff87aada0adb1355f86e1e547f8526af95a143b08b64cd3"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-d3a82d4fed1858b2757486eb6e4555d66fc03edcc7201b4ad4cc675322e044b9"></a>

## Direct properties — remote_ip.ip.ipv6 / 9133a2f2595b / 3

<a id="canonical-92ea45de2e8e94a05f9dda9b1d36a13364ff8f912bf2cc33a61d89df5fba9a75"></a>

<a id="canonical-899e9a0d57a7eb479b902b0e6fa47632f06a460635c97243d23dc99b729859dd"></a>

## addr property — remote_ip.ip.ipv6 / 9133a2f2595b / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-70b43b4454d3b5d044551c8cce88270ac01fd982e55413ba2eeec9d30f277d4a"></a>

## Next pages — remote_ip.ip.ipv6 / 9133a2f2595b / 5

- [remote_ip.ip](data-sources--tunnel--reference--group-001.md#canonical-1029a659f4ff1f4a154c1ada94e749c3b1320a715135d7ca39f8335fb99459be)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
