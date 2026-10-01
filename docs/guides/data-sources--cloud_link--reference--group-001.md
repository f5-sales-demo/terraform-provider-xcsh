---
page_title: "xcsh_cloud_link reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link reference."
---

# xcsh_cloud_link reference

<a id="canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47eaf9d294c78617836b7b0d861355d7916b38b2b8cc035a6bdb24357131ab6d"></a>

## Property reference — Property reference / dc807d14d745 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- Property reference

<a id="canonical-abc74a9d3b81f3dba5f11e625a2f037260a7745b66f740e8271dbe5fa1c2690d"></a>

## Direct properties — Property reference / dc807d14d745 / 3

<a id="canonical-f0f6576e02289b0f2e3ae67289f25f86125b6b64d259a88dc8a092dda642fe55"></a>

<a id="canonical-2bd86851d91e76c0ade2b7d44d72d652959fc8872c9c0ae57de57a95359f4797"></a>

## annotations property — Property reference / dc807d14d745 / 4

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

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325): complete subsection reference.

<a id="canonical-f9a4f1c71ddc342be403a4d8dccbaa2ccf6585eb24b768f80ae80483e3d0a215"></a>

<a id="canonical-408272df84d1fc4d1f7cbc3c21811117251df7c6f5b769a7bfd53258ed310210"></a>

## description property — Property reference / dc807d14d745 / 5

Type: `"string"`. Computed.

Description of the CloudLink.

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

- [disabled](data-sources--cloud_link--reference--group-001.md#canonical-59ab90433172a7badf243f644b66d9b88f86e66c64396a669af649e0a797f770): complete subsection reference.

- [enabled](data-sources--cloud_link--reference--group-001.md#canonical-9f957cbd554f7861b0df8bb24192e482faaf8994623ebfc9425c66a6d901e458): complete subsection reference.

- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c): complete subsection reference.

<a id="canonical-40f80c3b31d3d8e49f585d6fe66167394210a058864129a014c84a4bb25eb70f"></a>

<a id="canonical-d489f00b697a3266f261ab3a05b00e4a1aa00f8dd4cdba299b2741b128e71616"></a>

## id property — Property reference / dc807d14d745 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-938867e50808ff520449805f36bbdfd6cac9e64ca1a5eb2066e3af7368185f96"></a>

<a id="canonical-c49dccc69bde52714f620560504276a11645846eabb0a1dd76d9c3c3da05c72a"></a>

## labels property — Property reference / dc807d14d745 / 7

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

<a id="canonical-cf24e6476851a6c3aca02c8a20a3860593a46fe496b01a509305871000b4ea42"></a>

<a id="canonical-1047f19a1244d7931433a096af341eb9d644f8577d6a65dcbc0a9f65e7ee7f37"></a>

## name property — Property reference / dc807d14d745 / 8

Type: `"string"`. Required.

Name of the CloudLink.

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

<a id="canonical-aff06e5f03676039fd34b5223cdd3ced2614fe0345a4d9e41df220ec27617e28"></a>

<a id="canonical-58a51516514ac494fae835f1942c601468b138303f469adb2329fc0f995ed922"></a>

## namespace property — Property reference / dc807d14d745 / 9

Type: `"string"`. Required.

Namespace where the CloudLink exists.

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

<a id="canonical-ff78d52a9501d04a20cca80313c99142bf976e31c0c2ead41afaa7e4767bcee2"></a>

## All schema paths — Property reference / dc807d14d745 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_link--reference--group-001.md#canonical-f0f6576e02289b0f2e3ae67289f25f86125b6b64d259a88dc8a092dda642fe55) |
| `aws` | [aws](data-sources--cloud_link--reference--group-001.md#canonical-cdcdf3d1f5ac4155111c77c20cd72f345150b82343c321837d56b09db4615402) |
| `aws.aws_cred` | [aws.aws_cred](data-sources--cloud_link--reference--group-001.md#canonical-48bce29d2d03e5f51c8cd0a2c499d82ca6b58a17cf53e79a5d90db911f74c558) |
| `aws.aws_cred.name` | [aws.aws_cred.name](data-sources--cloud_link--reference--group-001.md#canonical-5a9df029ddba35e11b99eb61e829a2b05d29c9e8063c346a6e03e79bc4f60edf) |
| `aws.aws_cred.namespace` | [aws.aws_cred.namespace](data-sources--cloud_link--reference--group-001.md#canonical-73066232abcb75abeca44fbf488535b3f6c8df04c6189b541187f04064360ac0) |
| `aws.aws_cred.tenant` | [aws.aws_cred.tenant](data-sources--cloud_link--reference--group-001.md#canonical-54ad41c5e65490d34cc0eb4bfc26f9a994f212351a8f1f32355687a4cb4edfde) |
| `aws.byoc` | [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-c819e024f7e302f07d20b2ae22974601fbdcd4b41c0963d2c90e3cdc4cd715c9) |
| `aws.byoc.connections` | [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-2dc1dfb0e5145bf7b733b20778389df3ae6cdbf24c83f4ebe9e26cac73d14e52) |
| `aws.byoc.connections.auth_key` | [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-d3053f44fa3529edde0467e75990c4188e2361bf123932c5053e0263e2377409) |
| `aws.byoc.connections.auth_key.blindfold_secret_info` | [aws.byoc.connections.auth_key.blindfold_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-432051a554ec737d88e37c1c8df0e7e37ed8770df81a7ba676e1bf7341ad0a8f) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider](data-sources--cloud_link--reference--group-001.md#canonical-f79e10f674814b7daccd3a8a78b0642f290a793d8bffcb02d8455657a08e3e32) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.location` | [aws.byoc.connections.auth_key.blindfold_secret_info.location](data-sources--cloud_link--reference--group-001.md#canonical-8c3ca233eba86a8ecd1fbeb209715a164e5d230fdf17faae5c2c29ec77dc22f3) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.store_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.store_provider](data-sources--cloud_link--reference--group-001.md#canonical-0550d75d353bac20dcc73dd0711999858eeb9e64cc2dc0b993726e595595d71b) |
| `aws.byoc.connections.auth_key.clear_secret_info` | [aws.byoc.connections.auth_key.clear_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-e0cd990c172632428b88da2da8e0196597d48cdbe02425d0c0953dced0b88f03) |
| `aws.byoc.connections.auth_key.clear_secret_info.provider_ref` | [aws.byoc.connections.auth_key.clear_secret_info.provider_ref](data-sources--cloud_link--reference--group-001.md#canonical-973fd0e38b88ed5d1b6ee45f910296e68c11a8beacb98fcd7573edb72a01aa15) |
| `aws.byoc.connections.auth_key.clear_secret_info.url` | [aws.byoc.connections.auth_key.clear_secret_info.url](data-sources--cloud_link--reference--group-001.md#canonical-50aa436676cc6558d91917736f083eb1dca088af8e7a3fa6ec42d45045448b4a) |
| `aws.byoc.connections.bgp_asn` | [aws.byoc.connections.bgp_asn](data-sources--cloud_link--reference--group-001.md#canonical-dfab2a92048a10a57fcf4d83fd5731b52f00bb815cf19affdcdabe5cf3e83ab5) |
| `aws.byoc.connections.connection_id` | [aws.byoc.connections.connection_id](data-sources--cloud_link--reference--group-001.md#canonical-62d4f5333fc7f73c9f9c25bf8c0bbc3a4047f823e96e91750503197175ec0cfc) |
| `aws.byoc.connections.ipv4` | [aws.byoc.connections.ipv4](data-sources--cloud_link--reference--group-001.md#canonical-fc0792909bf465f7885ebb4a76fa175a0dd91a652726b68ea3ee839347121544) |
| `aws.byoc.connections.ipv4.aws_router_peer_address` | [aws.byoc.connections.ipv4.aws_router_peer_address](data-sources--cloud_link--reference--group-001.md#canonical-113502db77253ed1f0a07a25bf28e790e84d84f840eaf670f16dc9a3681a4dce) |
| `aws.byoc.connections.ipv4.router_peer_address` | [aws.byoc.connections.ipv4.router_peer_address](data-sources--cloud_link--reference--group-001.md#canonical-8123d75c4cba54ed09a8ea6de5d8bdc638a20fd7e6513118b1bf8eacc9f6fa1d) |
| `aws.byoc.connections.metadata` | [aws.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-47582190d49d7233bb82f8f2754ea6b8e6ccc78b00b2237095541e773fa67fff) |
| `aws.byoc.connections.metadata.description_spec` | [aws.byoc.connections.metadata.description_spec](data-sources--cloud_link--reference--group-001.md#canonical-7be7223e94230d424ee17e0267658d3b68e478286380ceaf57ab4d8ad3f2700c) |
| `aws.byoc.connections.metadata.name` | [aws.byoc.connections.metadata.name](data-sources--cloud_link--reference--group-001.md#canonical-310d8c030110c8ab251e66d06a3bfff183f7e540541c4db4f4feabb239200191) |
| `aws.byoc.connections.region` | [aws.byoc.connections.region](data-sources--cloud_link--reference--group-001.md#canonical-eb1d0f9f191d2fb629abd9277d0c31fe2a34d59cf03585254d8f491f01a82445) |
| `aws.byoc.connections.system_generated_name` | [aws.byoc.connections.system_generated_name](data-sources--cloud_link--reference--group-001.md#canonical-5b7592250545c059fbeff1fee4e8675f4a875a7bc4cde03dc3cf93a36421befa) |
| `aws.byoc.connections.tags` | [aws.byoc.connections.tags](data-sources--cloud_link--reference--group-001.md#canonical-84ed747ff275cb53f3c269da3d5237930df99e674f2379b17b7d3f63e48028bd) |
| `aws.byoc.connections.user_assigned_name` | [aws.byoc.connections.user_assigned_name](data-sources--cloud_link--reference--group-001.md#canonical-9e4df729d2c85618c14eb05c4493c703931651fc9500b9ffa474804061e8508d) |
| `aws.byoc.connections.virtual_interface_type` | [aws.byoc.connections.virtual_interface_type](data-sources--cloud_link--reference--group-001.md#canonical-8e3c6dcd47038ea1fa3f50fce459a9c9c1f7a661e1dd726665dceebdca7e5981) |
| `aws.byoc.connections.vlan` | [aws.byoc.connections.vlan](data-sources--cloud_link--reference--group-001.md#canonical-6d4adaebb4ee3efe29b9d414e4ee4f756b686b7d85cbfdb9bfd51a576e5c6a03) |
| `aws.custom_asn` | [aws.custom_asn](data-sources--cloud_link--reference--group-001.md#canonical-9bc07a5023f62d1ce3c8ca30e29a7c3b5139a394a68215fcf780abe8cd09ea5e) |
| `description` | [description](data-sources--cloud_link--reference--group-001.md#canonical-f9a4f1c71ddc342be403a4d8dccbaa2ccf6585eb24b768f80ae80483e3d0a215) |
| `disabled` | [disabled](data-sources--cloud_link--reference--group-001.md#canonical-0b9ed49954cedf179ead96deeb2bdd1ea2645072ec404bcbdea38b87bd38e9fc) |
| `enabled` | [enabled](data-sources--cloud_link--reference--group-001.md#canonical-baa67e1570c989b8c48c901c63591c413f6daa344bd409391a6df1204176bc2e) |
| `enabled.cloudlink_network_name` | [enabled.cloudlink_network_name](data-sources--cloud_link--reference--group-001.md#canonical-58d04897e03cf10e205ffa2d480244881b02fb6f87ba8450ffcf68f2d1f3dc41) |
| `gcp` | [gcp](data-sources--cloud_link--reference--group-001.md#canonical-cc86bdd1ff4b43d933c35ad7bc161729f4c26e9b5f7bc06db7ed5dc3cb01e62b) |
| `gcp.byoc` | [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-60abf575e48179a1f85aee43782da6454e336f2a111b30f99ceb40f9cd83bc0d) |
| `gcp.byoc.connections` | [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-0e932775a4bbcdaca7a02313b12a57921a5ac5e390fcb688a1c769dbfa7c5147) |
| `gcp.byoc.connections.interconnect_attachment_name` | [gcp.byoc.connections.interconnect_attachment_name](data-sources--cloud_link--reference--group-001.md#canonical-bb98968dfced8e2821ec6ad154aecbb03892348ef6a7f25e4f4dfaa07899c3c8) |
| `gcp.byoc.connections.metadata` | [gcp.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-ab169ef2d1b1435dc8b1ac4b6c1272fe7e13413ba21311503f587fd53c06b6b3) |
| `gcp.byoc.connections.metadata.description_spec` | [gcp.byoc.connections.metadata.description_spec](data-sources--cloud_link--reference--group-001.md#canonical-412fdb7ec63676efc4bee5e20bc5d8a18adfa2725d8906a15cdac2587533ec3e) |
| `gcp.byoc.connections.metadata.name` | [gcp.byoc.connections.metadata.name](data-sources--cloud_link--reference--group-001.md#canonical-8e6c12152411551eb0151556798b84855ea9a28c827f396858920390ac9baefb) |
| `gcp.byoc.connections.project` | [gcp.byoc.connections.project](data-sources--cloud_link--reference--group-001.md#canonical-7f0b2384420f80ea537f40c16b28812a0743719d4c2e226d97e6130f6e759569) |
| `gcp.byoc.connections.region` | [gcp.byoc.connections.region](data-sources--cloud_link--reference--group-001.md#canonical-f75d05abb38d1aea37ea2a53b0eaca7ab0b22ddacc22215709cfcceee6733312) |
| `gcp.byoc.connections.same_as_credential` | [gcp.byoc.connections.same_as_credential](data-sources--cloud_link--reference--group-001.md#canonical-1bcdcb929b13cb99f175446f4e00969383e039d5345e08efd42e2b2438b5528f) |
| `gcp.gcp_cred` | [gcp.gcp_cred](data-sources--cloud_link--reference--group-001.md#canonical-58923d04ede821d402edd36afc2b12cd07fa1f21be3a1fc6952354b9e402c51b) |
| `gcp.gcp_cred.name` | [gcp.gcp_cred.name](data-sources--cloud_link--reference--group-001.md#canonical-eb228e0d396bcc66e1e516df45e0968ef4af13988ed6aebef302a1630cfe591a) |
| `gcp.gcp_cred.namespace` | [gcp.gcp_cred.namespace](data-sources--cloud_link--reference--group-001.md#canonical-fb3b226bf5f457c4374387377adccae7ad4329fae2dbf8e6edf698711364c572) |
| `gcp.gcp_cred.tenant` | [gcp.gcp_cred.tenant](data-sources--cloud_link--reference--group-001.md#canonical-522104c1663b1f5e687362371539c077483cffbca2299eabca632f7b2492ed77) |
| `id` | [id](data-sources--cloud_link--reference--group-001.md#canonical-40f80c3b31d3d8e49f585d6fe66167394210a058864129a014c84a4bb25eb70f) |
| `labels` | [labels](data-sources--cloud_link--reference--group-001.md#canonical-938867e50808ff520449805f36bbdfd6cac9e64ca1a5eb2066e3af7368185f96) |
| `name` | [name](data-sources--cloud_link--reference--group-001.md#canonical-cf24e6476851a6c3aca02c8a20a3860593a46fe496b01a509305871000b4ea42) |
| `namespace` | [namespace](data-sources--cloud_link--reference--group-001.md#canonical-aff06e5f03676039fd34b5223cdd3ced2614fe0345a4d9e41df220ec27617e28) |

<a id="canonical-2780037ef44844858de63165d29db069643e4099f3be6ddc07bf12bb177378f5"></a>

## Next pages — Property reference / dc807d14d745 / 11

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [disabled](data-sources--cloud_link--reference--group-001.md#canonical-59ab90433172a7badf243f644b66d9b88f86e66c64396a669af649e0a797f770)
- [enabled](data-sources--cloud_link--reference--group-001.md#canonical-9f957cbd554f7861b0df8bb24192e482faaf8994623ebfc9425c66a6d901e458)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ee184d81572e236a82532002a337d007d895c8fdb4471d112fd067bb254ae59"></a>

## aws — aws / 62a5146a2bc2 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- aws

<a id="canonical-cdcdf3d1f5ac4155111c77c20cd72f345150b82343c321837d56b09db4615402"></a>

Type: `"single"`. Computed.

\[OneOf: aws, gcp\] Amazon Web Services(AWS) CloudLink Provider. CloudLink for AWS Cloud Provider.

Upstream description:

CloudLink for AWS Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]",
  "x-ves-oneof-field-direct_connect_gateway_asn_choice": "[\"custom_asn\"]"
}
```

OneOf alternatives in this subsection:

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-cdcdf3d1f5ac4155111c77c20cd72f345150b82343c321837d56b09db4615402)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-cc86bdd1ff4b43d933c35ad7bc161729f4c26e9b5f7bc06db7ed5dc3cb01e62b)

Select alternatives according to the provider validators above.

<a id="canonical-6ec6219a78d2735b0f38e43bcab636ca7116e151eb9146c711f38100232c70b4"></a>

## Direct properties — aws / 62a5146a2bc2 / 3

- [aws_cred](data-sources--cloud_link--reference--group-001.md#canonical-2a0a3ee370c7b78eca1f0676d49af9ffa3c9397313f5c290be3d0671dfd29649): complete subsection reference.

- [byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb): complete subsection reference.

<a id="canonical-9bc07a5023f62d1ce3c8ca30e29a7c3b5139a394a68215fcf780abe8cd09ea5e"></a>

<a id="canonical-c7ccec2268e6f2d2cbb5273586c89d127c86c177e4efafe84228d7dbccbb28f2"></a>

## custom_asn property — aws / 62a5146a2bc2 / 4

Type: `"number"`. Computed.

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Upstream description:

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4294967294,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  }
}
```

<a id="canonical-5d98eedc55f2e3da5312abeeea3eebadf2f889f123fcd050cb02d7835467f19a"></a>

## Next pages — aws / 62a5146a2bc2 / 5

- [aws.aws_cred](data-sources--cloud_link--reference--group-001.md#canonical-2a0a3ee370c7b78eca1f0676d49af9ffa3c9397313f5c290be3d0671dfd29649)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-2a0a3ee370c7b78eca1f0676d49af9ffa3c9397313f5c290be3d0671dfd29649"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3b8186d1aa8c104b1da910dd5ad66a572f68b7d4fdd757d7f311c3ba3e06178"></a>

## aws.aws_cred — aws.aws_cred / fc7b2ac35694 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- aws.aws_cred

<a id="canonical-48bce29d2d03e5f51c8cd0a2c499d82ca6b58a17cf53e79a5d90db911f74c558"></a>

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

<a id="canonical-04bcc952b4a206e833b4b4fdf910998ef0157f67e23c80ba38275688959b2714"></a>

## Direct properties — aws.aws_cred / fc7b2ac35694 / 3

<a id="canonical-5a9df029ddba35e11b99eb61e829a2b05d29c9e8063c346a6e03e79bc4f60edf"></a>

<a id="canonical-5f68e88c0b4dd58057cdaa7600f1e6fc3ca75fb5725aea0615438a457ee106cc"></a>

## name property — aws.aws_cred / fc7b2ac35694 / 4

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

<a id="canonical-73066232abcb75abeca44fbf488535b3f6c8df04c6189b541187f04064360ac0"></a>

<a id="canonical-90b7a5b2959c03a32d3af06f6a432c85de40b6ac5dfc14d4618d95e50ccec787"></a>

## namespace property — aws.aws_cred / fc7b2ac35694 / 5

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

<a id="canonical-54ad41c5e65490d34cc0eb4bfc26f9a994f212351a8f1f32355687a4cb4edfde"></a>

<a id="canonical-20dd47c13ed2a8fb0bd25169adaca5e4c4aea3b740d88bb13832fd76f90c9d93"></a>

## tenant property — aws.aws_cred / fc7b2ac35694 / 6

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

<a id="canonical-f87c2cc403d776d65d575834a1e21dc77986e6dbeffe4ee528841d0d8bfeaa45"></a>

## Next pages — aws.aws_cred / fc7b2ac35694 / 7

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48955a28c8ea933ab955a227f56c1d462df8d284ab3adc10b67ffd229c83718a"></a>

## aws.byoc — aws.byoc / a94e397c1c11 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- aws.byoc

<a id="canonical-c819e024f7e302f07d20b2ae22974601fbdcd4b41c0963d2c90e3cdc4cd715c9"></a>

Type: `"single"`. Computed.

Bring Your Own Connections. List of Bring You Own Connection.

Upstream description:

List of Bring You Own Connection.

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

<a id="canonical-61a695afd2a6862a1bf130e2e172e2a053ada285903790497e29626e72661d1e"></a>

## Direct properties — aws.byoc / a94e397c1c11 / 3

- [connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799): complete subsection reference.

<a id="canonical-fce5562ee7e3d6c3456ec3a169a2cd73cf311d996bb4ad90d2ed3263e396684b"></a>

## Next pages — aws.byoc / a94e397c1c11 / 4

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f111c19cdd83d05f416ce6e39c18e9af1a6a598bf265fbc542df8d79ec4d3080"></a>

## aws.byoc.connections — aws.byoc.connections / 520cea887b24 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- aws.byoc.connections

<a id="canonical-2dc1dfb0e5145bf7b733b20778389df3ae6cdbf24c83f4ebe9e26cac73d14e52"></a>

Type: `"list"`. Computed.

List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but
will be used for connecting sites and REs.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-b7168aaba73d097580ad8feef09f7005702d65b608ba4ca131266d1b887b781b"></a>

## Direct properties — aws.byoc.connections / 520cea887b24 / 3

- [auth_key](data-sources--cloud_link--reference--group-001.md#canonical-c7fa0866a96c84dd0f1f1b518dbd44fd7a3453a21d673d83fe591d364a365391): complete subsection reference.

<a id="canonical-dfab2a92048a10a57fcf4d83fd5731b52f00bb815cf19affdcdabe5cf3e83ab5"></a>

<a id="canonical-5753a0052868d1d4d02f31ee480e1c5f7ce4c37a61117a8964e68b529df6ff40"></a>

## bgp_asn property — aws.byoc.connections / 520cea887b24 / 4

Type: `"number"`. Computed.

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Upstream description:

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-62d4f5333fc7f73c9f9c25bf8c0bbc3a4047f823e96e91750503197175ec0cfc"></a>

<a id="canonical-e0ac43e2fdceea951242192980fd69f709da6cd323794ab24695ab7f0f4b7cda"></a>

## connection_id property — aws.byoc.connections / 520cea887b24 / 5

Type: `"string"`. Computed.

ID of the existing AWS Direct Connect Connection.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [ipv4](data-sources--cloud_link--reference--group-001.md#canonical-6bb7b6603df60cd2cf68efbebc1a8015b7001216623e47c1decd723270e69d68): complete subsection reference.

- [metadata](data-sources--cloud_link--reference--group-001.md#canonical-6b731452736b66da5a17f5c8ebc1af52af53d9c880a22efed62f201bca13eff8): complete subsection reference.

<a id="canonical-eb1d0f9f191d2fb629abd9277d0c31fe2a34d59cf03585254d8f491f01a82445"></a>

<a id="canonical-c262f9b134c12d3d251a22675f3499e79c36f60ff4eebb9d545a84b97b95a616"></a>

## region property — aws.byoc.connections / 520cea887b24 / 6

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
Region. Region where the connection is setup. Possible values are \`ap-northeast-1\`,
\`ap-southeast-1\`, \`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`,
\`us-east-2\`, \`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`,
\`ap-northeast-2\`, \`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`,
\`me-south-1\`, \`us-west-1\`, \`ap-southeast-3\`.

Upstream description:

Region where the connection is setup.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [system_generated_name](data-sources--cloud_link--reference--group-001.md#canonical-a6436305aa8e517b9ab0639897e589c087c57353a091ee655e8531a004119c9c): complete subsection reference.

<a id="canonical-84ed747ff275cb53f3c269da3d5237930df99e674f2379b17b7d3f63e48028bd"></a>

<a id="canonical-0dcc192a00c624f7d996d24f11654e08623e65e4dd301acf9bd1a56863fb3569"></a>

## tags property — aws.byoc.connections / 520cea887b24 / 7

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-9e4df729d2c85618c14eb05c4493c703931651fc9500b9ffa474804061e8508d"></a>

<a id="canonical-fa04e471170210ca28a7a30dea4ee7da7749ebd220b771269a09698b3ac17fae"></a>

## user_assigned_name property — aws.byoc.connections / 520cea887b24 / 8

Type: `"string"`. Computed.

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

Upstream description:

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

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

<a id="canonical-8e3c6dcd47038ea1fa3f50fce459a9c9c1f7a661e1dd726665dceebdca7e5981"></a>

<a id="canonical-06889990516ea9346d429c4b42a3ab96640cc052b562c69b04a726201b0b468d"></a>

## virtual_interface_type property — aws.byoc.connections / 520cea887b24 / 9

Type: `"string"`. Computed.

\[Enum: PRIVATE\] Defines the type of virtual interface that needs to be configured on AWS -
PRIVATE: Private A private virtual interface should be used to access an Amazon VPC using private IP
addresses. - TRANSIT: Transit A transit virtual interface is a VLAN that transports traffic from a
Direct Connect.. The only possible value is \`PRIVATE\`. Defaults to \`PRIVATE\`.

Upstream description:

Defines the type of virtual interface that needs to be configured on AWS

&#8203;- PRIVATE: Private

A private virtual interface should be used to access an Amazon VPC using private IP addresses.
&#8203;- TRANSIT: Transit

A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one
or more transit gateways.

Receipt-pinned upstream constraints:

```json
{
  "default": "PRIVATE",
  "enum": [
    "PRIVATE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6d4adaebb4ee3efe29b9d414e4ee4f756b686b7d85cbfdb9bfd51a576e5c6a03"></a>

<a id="canonical-cdfa60b87edf0df151e2abf654d4231a177d587a1e4dba272b14702f8acdcdc5"></a>

## vlan property — aws.byoc.connections / 520cea887b24 / 10

Type: `"number"`. Computed.

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Upstream description:

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4094,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4094"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4094"
  }
}
```

<a id="canonical-ba708f23fc41ecd14377e00f927fd2dbbf2d1671f166b390e4c10c2f1908eb2f"></a>

## Next pages — aws.byoc.connections / 520cea887b24 / 11

- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-c7fa0866a96c84dd0f1f1b518dbd44fd7a3453a21d673d83fe591d364a365391)
- [aws.byoc.connections.ipv4](data-sources--cloud_link--reference--group-001.md#canonical-6bb7b6603df60cd2cf68efbebc1a8015b7001216623e47c1decd723270e69d68)
- [aws.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-6b731452736b66da5a17f5c8ebc1af52af53d9c880a22efed62f201bca13eff8)
- [aws.byoc.connections.system_generated_name](data-sources--cloud_link--reference--group-001.md#canonical-a6436305aa8e517b9ab0639897e589c087c57353a091ee655e8531a004119c9c)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-c7fa0866a96c84dd0f1f1b518dbd44fd7a3453a21d673d83fe591d364a365391"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b48f1a79a9f24b35ee7049b85c89bea978337f9970ecbabe5f2c0fa20f3ba7e"></a>

## aws.byoc.connections.auth_key — aws.byoc.connections.auth_key / 8527ac325ae9 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- aws.byoc.connections.auth_key

<a id="canonical-d3053f44fa3529edde0467e75990c4188e2361bf123932c5053e0263e2377409"></a>

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

<a id="canonical-d4aea29067d7eab082a9a2bb1edf6e1b68ed579b7bcc63337889cc62716e8fe5"></a>

## Direct properties — aws.byoc.connections.auth_key / 8527ac325ae9 / 3

- [blindfold_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-08392e84a8987cdab2d6e557d6229bd9c1b81c5d52bd46a5660ef70e6bcc5ccd): complete subsection reference.

- [clear_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-52d893e76d7fcb7791bee9928ebed7db8da4f78f22cd248591ad5f8c632649b1): complete subsection reference.

<a id="canonical-fbc9f28b1475e4371366019fd7dbb80e89452711ab074981c64290dabf8dea9c"></a>

## Next pages — aws.byoc.connections.auth_key / 8527ac325ae9 / 4

- [aws.byoc.connections.auth_key.blindfold_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-08392e84a8987cdab2d6e557d6229bd9c1b81c5d52bd46a5660ef70e6bcc5ccd)
- [aws.byoc.connections.auth_key.clear_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-52d893e76d7fcb7791bee9928ebed7db8da4f78f22cd248591ad5f8c632649b1)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-08392e84a8987cdab2d6e557d6229bd9c1b81c5d52bd46a5660ef70e6bcc5ccd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-116ef60e860f11ff2220b4962e62e1f41b0a555b155cde72254c11f9011ebe8c"></a>

## aws.byoc.connections.auth_key.blindfold_secret_info — aws.byoc.connections.auth_key.blindfold_secret_info / 04fcf0cae949 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-c7fa0866a96c84dd0f1f1b518dbd44fd7a3453a21d673d83fe591d364a365391)
- aws.byoc.connections.auth_key.blindfold_secret_info

<a id="canonical-432051a554ec737d88e37c1c8df0e7e37ed8770df81a7ba676e1bf7341ad0a8f"></a>

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

<a id="canonical-ba6f9357240d25533dfd77b39cb959616ee77899adcb560affa54f49e86718f9"></a>

## Direct properties — aws.byoc.connections.auth_key.blindfold_secret_info / 04fcf0cae949 / 3

<a id="canonical-f79e10f674814b7daccd3a8a78b0642f290a793d8bffcb02d8455657a08e3e32"></a>

<a id="canonical-7d47cff5f4ffe80ac733d7a3884fd8733f7423fdb4de63a56297337ddfc376d4"></a>

## decryption_provider property — aws.byoc.connections.auth_key.blindfold_secret_info / 04fcf0cae949 / 4

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

<a id="canonical-8c3ca233eba86a8ecd1fbeb209715a164e5d230fdf17faae5c2c29ec77dc22f3"></a>

<a id="canonical-a62ea76585662aded04696990a5437689e4eec6849154c90dd792c05166e3070"></a>

## location property — aws.byoc.connections.auth_key.blindfold_secret_info / 04fcf0cae949 / 5

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

<a id="canonical-0550d75d353bac20dcc73dd0711999858eeb9e64cc2dc0b993726e595595d71b"></a>

<a id="canonical-85eca8482f38fa942130863608d190b5d0e9a08f8b57c99f278bf96fa7f6e1d6"></a>

## store_provider property — aws.byoc.connections.auth_key.blindfold_secret_info / 04fcf0cae949 / 6

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

<a id="canonical-c207d8791612022f699d94082b743380071e2e986ee2cd9c2fa716712b4706de"></a>

## Next pages — aws.byoc.connections.auth_key.blindfold_secret_info / 04fcf0cae949 / 7

- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-c7fa0866a96c84dd0f1f1b518dbd44fd7a3453a21d673d83fe591d364a365391)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-52d893e76d7fcb7791bee9928ebed7db8da4f78f22cd248591ad5f8c632649b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e39e41dc38fa3b56b8158fd762361b83875736d347cc044b2edd3c62c355cfa6"></a>

## aws.byoc.connections.auth_key.clear_secret_info — aws.byoc.connections.auth_key.clear_secret_info / 8e309dc5e4b1 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-c7fa0866a96c84dd0f1f1b518dbd44fd7a3453a21d673d83fe591d364a365391)
- aws.byoc.connections.auth_key.clear_secret_info

<a id="canonical-e0cd990c172632428b88da2da8e0196597d48cdbe02425d0c0953dced0b88f03"></a>

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

<a id="canonical-e09ed9d5e6c4a9ad6e584c4bde6c0860fcbcaf0eef5646f1865b8c3d4af7326f"></a>

## Direct properties — aws.byoc.connections.auth_key.clear_secret_info / 8e309dc5e4b1 / 3

<a id="canonical-973fd0e38b88ed5d1b6ee45f910296e68c11a8beacb98fcd7573edb72a01aa15"></a>

<a id="canonical-779eaa7199acd571cd5cf9757dab8a1f95861ecea9b98fb86f9a453ca6337fc5"></a>

## provider_ref property — aws.byoc.connections.auth_key.clear_secret_info / 8e309dc5e4b1 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-50aa436676cc6558d91917736f083eb1dca088af8e7a3fa6ec42d45045448b4a"></a>

<a id="canonical-3e61c7191b2ea046fe6f0a22d4a67fa405f038f0053d4447bb2473c8e6210f61"></a>

## url property — aws.byoc.connections.auth_key.clear_secret_info / 8e309dc5e4b1 / 5

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

<a id="canonical-25db1b72ce9bae0f5c707fe38777e820ee914741af5329aec598202a2e590847"></a>

## Next pages — aws.byoc.connections.auth_key.clear_secret_info / 8e309dc5e4b1 / 6

- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-c7fa0866a96c84dd0f1f1b518dbd44fd7a3453a21d673d83fe591d364a365391)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-6bb7b6603df60cd2cf68efbebc1a8015b7001216623e47c1decd723270e69d68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5260c335531cea5fad879ed117ec17dbf764c6850a63f33159115fb0e042d3c"></a>

## aws.byoc.connections.ipv4 — aws.byoc.connections.ipv4 / 02f4740aa8a2 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- aws.byoc.connections.ipv4

<a id="canonical-fc0792909bf465f7885ebb4a76fa175a0dd91a652726b68ea3ee839347121544"></a>

Type: `"single"`. Computed.

Configure BGP IPv4 peering for endpoints.

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

<a id="canonical-4aa29e515e725e870774f4a661419af73e1568413e402b831eb12d7d268e3b8a"></a>

## Direct properties — aws.byoc.connections.ipv4 / 02f4740aa8a2 / 3

<a id="canonical-113502db77253ed1f0a07a25bf28e790e84d84f840eaf670f16dc9a3681a4dce"></a>

<a id="canonical-6d567d632406609e19a842e0d086749b3eed8ac0afcc5cc8f2f73ded211231f1"></a>

## aws_router_peer_address property — aws.byoc.connections.ipv4 / 02f4740aa8a2 / 4

Type: `"string"`. Computed.

The BGP peer IP configured on the AWS endpoint.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```

<a id="canonical-8123d75c4cba54ed09a8ea6de5d8bdc638a20fd7e6513118b1bf8eacc9f6fa1d"></a>

<a id="canonical-0297b14e318d3bcc16b67cbb310023d5e50f27330d7fd1bf4497f3de116195fc"></a>

## router_peer_address property — aws.byoc.connections.ipv4 / 02f4740aa8a2 / 5

Type: `"string"`. Computed.

The BGP peer IP configured on your (customer) endpoint.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```

<a id="canonical-89b6bd6dfb6511f085fbe860981eb37ab7ba0e2729aa5d5317fa9ae30a15778f"></a>

## Next pages — aws.byoc.connections.ipv4 / 02f4740aa8a2 / 6

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-6b731452736b66da5a17f5c8ebc1af52af53d9c880a22efed62f201bca13eff8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba292206ddda17d1acbf542d6f776a35218c45642af1da0524bf989105bd3507"></a>

## aws.byoc.connections.metadata — aws.byoc.connections.metadata / ef3eb50c8af0 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- aws.byoc.connections.metadata

<a id="canonical-47582190d49d7233bb82f8f2754ea6b8e6ccc78b00b2237095541e773fa67fff"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-9a3b8eadcf134df02f5c08fbab68cd7e8d7e860ab412df3aea94e882856c2296"></a>

## Direct properties — aws.byoc.connections.metadata / ef3eb50c8af0 / 3

<a id="canonical-7be7223e94230d424ee17e0267658d3b68e478286380ceaf57ab4d8ad3f2700c"></a>

<a id="canonical-72122f8204a1da185314109cd13083b284c01af546124f8b07e32f7f9bba1846"></a>

## description_spec property — aws.byoc.connections.metadata / ef3eb50c8af0 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-310d8c030110c8ab251e66d06a3bfff183f7e540541c4db4f4feabb239200191"></a>

<a id="canonical-a00c54877781a826b77b9df5c132f7b561a95fe5e880988a4fb2f9c962c68c8c"></a>

## name property — aws.byoc.connections.metadata / ef3eb50c8af0 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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

<a id="canonical-9515ef63577af67af623a9e6075abc8ecad52af0093787c6dc0156af9c0b8f81"></a>

## Next pages — aws.byoc.connections.metadata / ef3eb50c8af0 / 6

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-a6436305aa8e517b9ab0639897e589c087c57353a091ee655e8531a004119c9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26c70a451e380c8563e8d2dd65448ff035842bfe610fc530b91022cecbdc4ea4"></a>

## aws.byoc.connections.system_generated_name — aws.byoc.connections.system_generated_name / 991ff34758dd / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-fe19cda5ee95256baeff97dd3ad8f7664743313733ee8d3004d23000e0512325)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-4f9c1bd0e35a20cc3474a775bc04e7973177312071653efe2e86db2f9fc0c6bb)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- aws.byoc.connections.system_generated_name

<a id="canonical-5b7592250545c059fbeff1fee4e8675f4a875a7bc4cde03dc3cf93a36421befa"></a>

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

<a id="canonical-486067671d7948390ab2dc9badfd8afea88cd6bc2d3dd7f24b49143f5c7857e2"></a>

## Direct properties — aws.byoc.connections.system_generated_name / 991ff34758dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e44e93d0d27611c7bee2a95496229170657db05d4ce2598689ca208090dccac8"></a>

## Next pages — aws.byoc.connections.system_generated_name / 991ff34758dd / 4

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-4efbea7e12206bbf34cdc5a5c553f039c81d40315175b4a3db76918246fef799)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-59ab90433172a7badf243f644b66d9b88f86e66c64396a669af649e0a797f770"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb28b6261947aa01929e0eb245212f2a536bdc97070abe4a11045c7ebea7bc95"></a>

## disabled — disabled / c57f596cfa8f / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- disabled

<a id="canonical-0b9ed49954cedf179ead96deeb2bdd1ea2645072ec404bcbdea38b87bd38e9fc"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disabled, enabled\] Enable this option

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

- [disabled](data-sources--cloud_link--reference--group-001.md#canonical-0b9ed49954cedf179ead96deeb2bdd1ea2645072ec404bcbdea38b87bd38e9fc)
- [enabled](data-sources--cloud_link--reference--group-001.md#canonical-baa67e1570c989b8c48c901c63591c413f6daa344bd409391a6df1204176bc2e)

Select alternatives according to the provider validators above.

<a id="canonical-963c0a681b092e81d8607cc8f27f19a9c76ab698ce1d646ce6822da55dcd1525"></a>

## Direct properties — disabled / c57f596cfa8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afb9eaf955892e90e98fab6bb1c08874a9803c711ea9f73ad0bb0afe71ecea22"></a>

## Next pages — disabled / c57f596cfa8f / 4

- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-9f957cbd554f7861b0df8bb24192e482faaf8994623ebfc9425c66a6d901e458"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78bc964531513c95aacd88448bf59b731f5f0b2a1ed47a48366da2c26d5feb37"></a>

## enabled — enabled / 0d7f5771ef83 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- enabled

<a id="canonical-baa67e1570c989b8c48c901c63591c413f6daa344bd409391a6df1204176bc2e"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

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

<a id="canonical-c971276b3ca22022365620df507643f9b878db7cadf9b23f6f0b93e08f60dd7e"></a>

## Direct properties — enabled / 0d7f5771ef83 / 3

<a id="canonical-58d04897e03cf10e205ffa2d480244881b02fb6f87ba8450ffcf68f2d1f3dc41"></a>

<a id="canonical-3eab79002fdf01bfbb06f9abe8430bb8dd9eb699c9e0db430b368fc44e922eee"></a>

## cloudlink_network_name property — enabled / 0d7f5771ef83 / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3d671d71798604fcf3f7a934c98132bb491bef7abab5b4a0bd59a4e02acc9a10"></a>

## Next pages — enabled / 0d7f5771ef83 / 5

- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1932723b489549a2d0f68cf67902685d95b0e243bd5db5395ac6f9a6f9eceffb"></a>

## gcp — gcp / 28c5e1e4ad06 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- gcp

<a id="canonical-cc86bdd1ff4b43d933c35ad7bc161729f4c26e9b5f7bc06db7ed5dc3cb01e62b"></a>

Type: `"single"`. Computed.

Google Cloud Platform (GCP) CloudLink Provider. CloudLink for GCP Cloud Provider.

Upstream description:

CloudLink for GCP Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]"
}
```

<a id="canonical-04a4db44248ee0fe83dc4c3b94842266c2ed670baf1cbfa54d4729ab094fd8e1"></a>

## Direct properties — gcp / 28c5e1e4ad06 / 3

- [byoc](data-sources--cloud_link--reference--group-001.md#canonical-42947db8b21d9577d29f26791833867f2f82bb15654e539909c26a20bb305c49): complete subsection reference.

- [gcp_cred](data-sources--cloud_link--reference--group-001.md#canonical-b831c38c126de8b5deeca170c7144f1db2e87fe798c3fa4e50d436067db7c286): complete subsection reference.

<a id="canonical-f611d778efd8a97942b108b58da0858ae2063ea8fa03b309e377d69156732ffd"></a>

## Next pages — gcp / 28c5e1e4ad06 / 4

- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-42947db8b21d9577d29f26791833867f2f82bb15654e539909c26a20bb305c49)
- [gcp.gcp_cred](data-sources--cloud_link--reference--group-001.md#canonical-b831c38c126de8b5deeca170c7144f1db2e87fe798c3fa4e50d436067db7c286)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-42947db8b21d9577d29f26791833867f2f82bb15654e539909c26a20bb305c49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85dbcab812e109fd98ccc77f72f84bc777d81e9927b226c8d51f8e623de7c8dc"></a>

## gcp.byoc — gcp.byoc / e7deab18cadf / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- gcp.byoc

<a id="canonical-60abf575e48179a1f85aee43782da6454e336f2a111b30f99ceb40f9cd83bc0d"></a>

Type: `"single"`. Computed.

GCP Bring Your Own Connections. List of GCP Bring You Own Connections.

Upstream description:

List of GCP Bring You Own Connections.

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

<a id="canonical-692a4e420fbc7939477ad7e45a523c8f9f5465370144d6ed2c8b116323c6b6f1"></a>

## Direct properties — gcp.byoc / e7deab18cadf / 3

- [connections](data-sources--cloud_link--reference--group-001.md#canonical-e5fb3be6d553f01fa13c79a93fabda661d163c476a2e672711f090f488efc17f): complete subsection reference.

<a id="canonical-ae8dd31cf36eae0ca6fd5d485682523aff6892527c8575691a76a41e37d022ac"></a>

## Next pages — gcp.byoc / e7deab18cadf / 4

- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-e5fb3be6d553f01fa13c79a93fabda661d163c476a2e672711f090f488efc17f)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-e5fb3be6d553f01fa13c79a93fabda661d163c476a2e672711f090f488efc17f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96af307277b6321874a8f2120b7d1a1e9b952b238b3c4d74f865a2beb7b5dc07"></a>

## gcp.byoc.connections — gcp.byoc.connections / ece97265b27c / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-42947db8b21d9577d29f26791833867f2f82bb15654e539909c26a20bb305c49)
- gcp.byoc.connections

<a id="canonical-0e932775a4bbcdaca7a02313b12a57921a5ac5e390fcb688a1c769dbfa7c5147"></a>

Type: `"list"`. Computed.

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (.

Upstream description:

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud
to facilitate seamless private connectivity.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-940e1bc7fe7a5b11868a140d2c40620d65d1af40aa769ab9b3f52159835eb3c5"></a>

## Direct properties — gcp.byoc.connections / ece97265b27c / 3

<a id="canonical-bb98968dfced8e2821ec6ad154aecbb03892348ef6a7f25e4f4dfaa07899c3c8"></a>

<a id="canonical-1462a952e4c41f3b880428be83b2c61aeeced77407dcdd8f49d07d234faa41b8"></a>

## interconnect_attachment_name property — gcp.byoc.connections / ece97265b27c / 4

Type: `"string"`. Computed.

Name of already-existing GCP Cloud Interconnect Attachment.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [metadata](data-sources--cloud_link--reference--group-001.md#canonical-233850c3d640a2f4248e5ff505b3041565b87bd000dd2d26ad3638aa57b94525): complete subsection reference.

<a id="canonical-7f0b2384420f80ea537f40c16b28812a0743719d4c2e226d97e6130f6e759569"></a>

<a id="canonical-dd8715bc141e1f0c7ab158efb51e82eca2fe83ff92784629cef62dc98076e408"></a>

## project property — gcp.byoc.connections / ece97265b27c / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Upstream description:

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 30,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="canonical-f75d05abb38d1aea37ea2a53b0eaca7ab0b22ddacc22215709cfcceee6733312"></a>

<a id="canonical-ae2501560d091a3ce5a2403896a9b851899db7b1b40604d9e50a6aa63033ce18"></a>

## region property — gcp.byoc.connections / ece97265b27c / 6

Type: `"string"`. Computed.

\[Enum:
asia-east1|asia-east2|asia-northeast1|asia-northeast2|asia-northeast3|asia-southeast1|asia-southeast2|europe-central2|europe-north1|europe-west1|europe-west2|europe-west3|europe-west4|europe-west6|europe-west8|europe-west9|europe-west10|europe-west12|europe-southwest1|me-west1|me-central1|me-central2|northamerica-northeast1|northamerica-northeast2|us-central1|us-east1|us-east4|us-east5|us-south1|us-west1|us-west2|us-west3|us-west4|southamerica-east1|southamerica-west1|australia-southeast1|australia-southeast2|asia-south1|asia-south2\]
GCP Region in which the GCP Cloud Interconnect attachment is configured. Possible values are
\`asia-east1\`, \`asia-east2\`, \`asia-northeast1\`, \`asia-northeast2\`, \`asia-northeast3\`,
\`asia-southeast1\`, \`asia-southeast2\`, \`europe-central2\`, \`europe-north1\`, \`europe-west1\`,
\`europe-west2\`, \`europe-west3\`, \`europe-west4\`, \`europe-west6\`, \`europe-west8\`,
\`europe-west9\`, \`europe-west10\`, \`europe-west12\`, \`europe-southwest1\`, \`me-west1\`,
\`me-central1\`, \`me-central2\`, \`northamerica-northeast1\`, \`northamerica-northeast2\`,
\`us-central1\`, \`us-east1\`, \`us-east4\`, \`us-east5\`, \`us-south1\`, \`us-west1\`,
\`us-west2\`, \`us-west3\`, \`us-west4\`, \`southamerica-east1\`, \`southamerica-west1\`,
\`australia-southeast1\`, \`australia-southeast2\`, \`asia-south1\`, \`asia-south2\`.

Upstream description:

GCP Region in which the GCP Cloud Interconnect attachment is configured.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "asia-east1",
    "asia-east2",
    "asia-northeast1",
    "asia-northeast2",
    "asia-northeast3",
    "asia-southeast1",
    "asia-southeast2",
    "europe-central2",
    "europe-north1",
    "europe-west1",
    "europe-west2",
    "europe-west3",
    "europe-west4",
    "europe-west6",
    "europe-west8",
    "europe-west9",
    "europe-west10",
    "europe-west12",
    "europe-southwest1",
    "me-west1",
    "me-central1",
    "me-central2",
    "northamerica-northeast1",
    "northamerica-northeast2",
    "us-central1",
    "us-east1",
    "us-east4",
    "us-east5",
    "us-south1",
    "us-west1",
    "us-west2",
    "us-west3",
    "us-west4",
    "southamerica-east1",
    "southamerica-west1",
    "australia-southeast1",
    "australia-southeast2",
    "asia-south1",
    "asia-south2"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  }
}
```

- [same_as_credential](data-sources--cloud_link--reference--group-001.md#canonical-2528d26fbb372a66edb6fffbb9bb5b3dea636d32525df2db58314f1d8374cc14): complete subsection reference.

<a id="canonical-98933fa7df391bc828d806e7c8cfef90589486cf0a1ce494a18dc8b8b219f510"></a>

## Next pages — gcp.byoc.connections / ece97265b27c / 7

- [gcp.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-233850c3d640a2f4248e5ff505b3041565b87bd000dd2d26ad3638aa57b94525)
- [gcp.byoc.connections.same_as_credential](data-sources--cloud_link--reference--group-001.md#canonical-2528d26fbb372a66edb6fffbb9bb5b3dea636d32525df2db58314f1d8374cc14)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-42947db8b21d9577d29f26791833867f2f82bb15654e539909c26a20bb305c49)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-233850c3d640a2f4248e5ff505b3041565b87bd000dd2d26ad3638aa57b94525"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1de758440d465259edc53ac761ca1b4d641a1225bf51b9371b921e969356712"></a>

## gcp.byoc.connections.metadata — gcp.byoc.connections.metadata / 5983d2762233 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-42947db8b21d9577d29f26791833867f2f82bb15654e539909c26a20bb305c49)
- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-e5fb3be6d553f01fa13c79a93fabda661d163c476a2e672711f090f488efc17f)
- gcp.byoc.connections.metadata

<a id="canonical-ab169ef2d1b1435dc8b1ac4b6c1272fe7e13413ba21311503f587fd53c06b6b3"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-6833781407c338b42567ce37248043e45985cea6988690beaecfba60fd3d264c"></a>

## Direct properties — gcp.byoc.connections.metadata / 5983d2762233 / 3

<a id="canonical-412fdb7ec63676efc4bee5e20bc5d8a18adfa2725d8906a15cdac2587533ec3e"></a>

<a id="canonical-c396507d981bdae80371268bf022ceb420fd7c9d71365cf1e6d487eda4537bdf"></a>

## description_spec property — gcp.byoc.connections.metadata / 5983d2762233 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-8e6c12152411551eb0151556798b84855ea9a28c827f396858920390ac9baefb"></a>

<a id="canonical-e62290c8919488c36d224545f46bb1e44b0545b805c306d74cd42ef296a5958c"></a>

## name property — gcp.byoc.connections.metadata / 5983d2762233 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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

<a id="canonical-6e7b2e1548e4d794addb4abf7faf46566e81a557941246038e0f418ec2280307"></a>

## Next pages — gcp.byoc.connections.metadata / 5983d2762233 / 6

- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-e5fb3be6d553f01fa13c79a93fabda661d163c476a2e672711f090f488efc17f)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-2528d26fbb372a66edb6fffbb9bb5b3dea636d32525df2db58314f1d8374cc14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4554eecf1e9c942a58b92fa68cb28da52ff36b7839bcb9ede72fa797771f74a"></a>

## gcp.byoc.connections.same_as_credential — gcp.byoc.connections.same_as_credential / 7c283c7cc6cf / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-42947db8b21d9577d29f26791833867f2f82bb15654e539909c26a20bb305c49)
- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-e5fb3be6d553f01fa13c79a93fabda661d163c476a2e672711f090f488efc17f)
- gcp.byoc.connections.same_as_credential

<a id="canonical-1bcdcb929b13cb99f175446f4e00969383e039d5345e08efd42e2b2438b5528f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as credential.

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

<a id="canonical-36824ca68a46139d277b9ae43376827b9f058f6dea796f104b14494e51443a79"></a>

## Direct properties — gcp.byoc.connections.same_as_credential / 7c283c7cc6cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54ce3a23f8c8f827c15c53a2a27610b55fc06613b3c3087fb4f955f6c1057a4f"></a>

## Next pages — gcp.byoc.connections.same_as_credential / 7c283c7cc6cf / 4

- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-e5fb3be6d553f01fa13c79a93fabda661d163c476a2e672711f090f488efc17f)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-b831c38c126de8b5deeca170c7144f1db2e87fe798c3fa4e50d436067db7c286"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1340536e89dc04dd2c52f41051c570fec098e724cf0db851c8fff55a6206e205"></a>

## gcp.gcp_cred — gcp.gcp_cred / 1573a812f979 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- gcp.gcp_cred

<a id="canonical-58923d04ede821d402edd36afc2b12cd07fa1f21be3a1fc6952354b9e402c51b"></a>

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

<a id="canonical-173ac2e25f8dbf8f2eefe4daf6e2c05f07a648ffb70de7b9b900ca2c97e499ec"></a>

## Direct properties — gcp.gcp_cred / 1573a812f979 / 3

<a id="canonical-eb228e0d396bcc66e1e516df45e0968ef4af13988ed6aebef302a1630cfe591a"></a>

<a id="canonical-b1c06ce0972dfa64bf0a1dc353f241c1776cee061933a4c68dc96dbe806c05ec"></a>

## name property — gcp.gcp_cred / 1573a812f979 / 4

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

<a id="canonical-fb3b226bf5f457c4374387377adccae7ad4329fae2dbf8e6edf698711364c572"></a>

<a id="canonical-493cc007c55c34155af5d806f73e5220f58f477cbab2893b70519fd080b5562c"></a>

## namespace property — gcp.gcp_cred / 1573a812f979 / 5

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

<a id="canonical-522104c1663b1f5e687362371539c077483cffbca2299eabca632f7b2492ed77"></a>

<a id="canonical-c9b563b08b4c276d4660ea43fb34c5a6aee0120025979b1d6e50f35a8d26b53c"></a>

## tenant property — gcp.gcp_cred / 1573a812f979 / 6

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

<a id="canonical-6d513c5c59aeb93551dd68d37a3bbaaeaf6be132fbacdd9e1820e3d5dc3e6b38"></a>

## Next pages — gcp.gcp_cred / 1573a812f979 / 7

- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-5e2c96ef3913598b37ced85bc4578014366ec96490bd1636bc23fc0324c62a2c)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
