---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-299a976f4685ee0e3af8b5871b1ee5abc430412549e2ab16ac1ce4e1303b9db7"></a>

## Property reference — Property reference / 53c4f33afa42 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- Property reference

<a id="canonical-532827f850ceeab392aadce45767044521555553ba2a1a4c0f4f20b3c19326d8"></a>

## Direct properties — Property reference / 53c4f33afa42 / 3

<a id="canonical-bc4ce0e48835bd8bb9109b9ae0030c95322e24aeeb9a5079f12aca716623270d"></a>

<a id="canonical-aa07a65a0ac64d8e5cffe81dcd75f346b9d2a4d1cc50093a5eeab4a52db9d8ea"></a>

## annotations property — Property reference / 53c4f33afa42 / 4

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

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872): complete subsection reference.

- [block_all_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-d97df5820ffa9083872d636b6f12ccfba661e8ceb4ced94125bf20220097cbf6): complete subsection reference.

- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-9065ec37d6ec0b7dfacc5a66d3ba150f8b71e54a085c55e401b3d26e0a27ee89): complete subsection reference.

- [coordinates](data-sources--aws_tgw_site--reference--group-002.md#canonical-f88da05b7925af8930063e804ed423f19943a3a16a6194a1a4426d8c4bd499f4): complete subsection reference.

- [custom_dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-45aeb23832ca3dbd757bdef1e8b6d6caf4010551ba9796e0e81c452aa73c6471): complete subsection reference.

- [default_blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-667795826786ba0811684ca6551f464c08907f534c98d05ab49463b40ce4439b): complete subsection reference.

<a id="canonical-c9340580743f91e1f7378d8ab447525930792d67015a2e1e1a2a2203b9cbffa1"></a>

<a id="canonical-420b63904b9af7254b330d3c62c876b3e846e747c63ae70877758212675d89aa"></a>

## description property — Property reference / 53c4f33afa42 / 5

Type: `"string"`. Computed.

Description of the AWSTGWSite.

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

- [direct_connect_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-1a0d78279f45609ad32ac2a33f4d1a3dc9c80eefd47cbaa48fc5f984d0320ac8): complete subsection reference.

- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-365a515e3f4c937f5dc99dac97b4ad775620553cf7b15c224d3cfe294ed537f0): complete subsection reference.

<a id="canonical-e4d0c041778bc2b2323f1046475a62b45b66f1a75866e36b4e34e733d271042c"></a>

<a id="canonical-6df27117627e2913de54dea5530449edfbeee3ce983abb80eb788b8cf07405d6"></a>

## id property — Property reference / 53c4f33afa42 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-3a4e0518667a945acda23b4526ee7eb358936563b025790a8f5da140ba6136ac): complete subsection reference.

<a id="canonical-4c20cce67d2db4c8b51d70cc468fbce0c4976673e4c4455f8cd4b038d0e0942a"></a>

<a id="canonical-77d9d470a1fd50b94d6bff9b2cf4f7baab994a5a0c7f289f86260902528d3c28"></a>

## labels property — Property reference / 53c4f33afa42 / 7

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

- [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-88271efdc13eb8fbbe3ea8a8f53c3e724d1e1af8bffc05f7401a3babc291c965): complete subsection reference.

- [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-57257a31846eb027c65f9471e86ce1b4007facd84f738bbd709e95af6ee482c0): complete subsection reference.

<a id="canonical-05893da18f07bdfd290a5878ecbed8951c3ba91fe0ddc201b3c8c59f3f946852"></a>

<a id="canonical-b1dd9309222338220bc2fee456ed01bdd805117e8388ef695b947e8befd75554"></a>

## name property — Property reference / 53c4f33afa42 / 8

Type: `"string"`. Required.

Name of the AWSTGWSite.

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

<a id="canonical-2cc9dcc893e15b630308e49328ae7627a108dc56a2c7dba66762a4047bb7e2cd"></a>

<a id="canonical-6f9ea180d6f1378ba22445ecd6ea27d5077a288e4de564624d8ff9bb7cd0c2cf"></a>

## namespace property — Property reference / 53c4f33afa42 / 9

Type: `"string"`. Required.

Namespace where the AWSTGWSite exists.

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

- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-4244e169f7d7eb4f78609fb64c52633e88576fb2a39de494147f6f4d3c29e2e4): complete subsection reference.

- [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-5d94edace7908538b2948f16b2cb5560db576aa216e7674b5fb8d1085d8f999c): complete subsection reference.

- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-3ce8e047006358cedc3c59b57fc4b80562ac72af3eefd5b6af40ab014380dc3e): complete subsection reference.

- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-b1d3abf5c88782ae24b64db8240510981bb2c45c61bfbc2ead4671d31690aa3c): complete subsection reference.

- [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-2274ef167fa95f7b95100502348384d42db2101bdd2013e74f0674421accdbd3): complete subsection reference.

<a id="canonical-322d346a9d567023210f03933cbbee637008d53bb4973e5d2e856c0c746d0265"></a>

<a id="canonical-a85e9e284bb9fb86b4bfb0937811cfd47b0334d6b430ff046ff274670317fc11"></a>

## tags property — Property reference / 53c4f33afa42 / 10

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-2e1ea519eb9353ad3ea4668d9e236a9eb32c420e919b7afa630cf37e55e12526): complete subsection reference.

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c): complete subsection reference.

- [vpc_attachments](data-sources--aws_tgw_site--reference--group-003.md#canonical-411569971dc5e487a1b7eb5dd5026c6a951d476dc5d2fda3bedd09ca3b953633): complete subsection reference.

- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0f930540b03ce3ec08b48a152be09afd0635e044541923a65add395e816e0f26): complete subsection reference.

<a id="canonical-1e7d4c2c455f0646600b42e5b97509545989f3b83553d127fc456f8352cc6371"></a>

## All schema paths — Property reference / 53c4f33afa42 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--aws_tgw_site--reference--group-001.md#canonical-bc4ce0e48835bd8bb9109b9ae0030c95322e24aeeb9a5079f12aca716623270d) |
| `aws_parameters` | [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-fe5668507c4b696604c2d7479fcef313aa239e41af57f2f4e596596174b17227) |
| `aws_parameters.admin_password` | [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-7cb4c6aa01a1eef88e82635af4d54c59805530cb271e63926bb45a0990951904) |
| `aws_parameters.admin_password.blindfold_secret_info` | [aws_parameters.admin_password.blindfold_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-8ea0b7b3e9f28e6192573cdb90cc0425ccd133d9c5fe462a49181302ec667344) |
| `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` | [aws_parameters.admin_password.blindfold_secret_info.decryption_provider](data-sources--aws_tgw_site--reference--group-001.md#canonical-5bf3dde0b7ad07db23f4fe349db8909a3cfd65ec6670b2878f5d5d463ac47f46) |
| `aws_parameters.admin_password.blindfold_secret_info.location` | [aws_parameters.admin_password.blindfold_secret_info.location](data-sources--aws_tgw_site--reference--group-001.md#canonical-77c3ca25058e70b854041518ac49ab4f5fc701c7dca7372d16dcb584e3a0305b) |
| `aws_parameters.admin_password.blindfold_secret_info.store_provider` | [aws_parameters.admin_password.blindfold_secret_info.store_provider](data-sources--aws_tgw_site--reference--group-001.md#canonical-cc232a1c77647b6335e39380e9fc1fe20c4ee377a4fcaf2f9ce709fb3e5a4d53) |
| `aws_parameters.admin_password.clear_secret_info` | [aws_parameters.admin_password.clear_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-12f7a2144dd802528fa48305f82261aeec79a9680206d01754585d730d67a58d) |
| `aws_parameters.admin_password.clear_secret_info.provider_ref` | [aws_parameters.admin_password.clear_secret_info.provider_ref](data-sources--aws_tgw_site--reference--group-001.md#canonical-590996c1f9e78b5429c242299a78aabd53ea0d91540aa717f2b9f976cefaffda) |
| `aws_parameters.admin_password.clear_secret_info.url` | [aws_parameters.admin_password.clear_secret_info.url](data-sources--aws_tgw_site--reference--group-001.md#canonical-f628fdf83a8229e4d41c41e358ceabc52dff7a1518f2c9a3ab997c08aa3cc0c4) |
| `aws_parameters.aws_cred` | [aws_parameters.aws_cred](data-sources--aws_tgw_site--reference--group-001.md#canonical-e87d2bbe5bfacc142975eb93a83741e9fd979cff5d5d039055649cbc83e50c78) |
| `aws_parameters.aws_cred.name` | [aws_parameters.aws_cred.name](data-sources--aws_tgw_site--reference--group-001.md#canonical-14ed388e5566984123500eb9315250425e37f4f05eb893714bee224a137b0c7d) |
| `aws_parameters.aws_cred.namespace` | [aws_parameters.aws_cred.namespace](data-sources--aws_tgw_site--reference--group-001.md#canonical-a0fc8964422db1f4ffeb618164193045f6556725a9b11b827f53733bb4b76b3a) |
| `aws_parameters.aws_cred.tenant` | [aws_parameters.aws_cred.tenant](data-sources--aws_tgw_site--reference--group-001.md#canonical-f27b231f78ea415949cd405553b461d56c4df610385c5b247860121d0628ff28) |
| `aws_parameters.aws_region` | [aws_parameters.aws_region](data-sources--aws_tgw_site--reference--group-001.md#canonical-394d08a7419c1a368f8a8bba3fb88609246d210ae0ce9acf5b56da4d5e495786) |
| `aws_parameters.az_nodes` | [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-5157eae2c0af635e1567709493298dffeef608a7ab4fc8c7ce77fe4d94d727d1) |
| `aws_parameters.az_nodes.aws_az_name` | [aws_parameters.az_nodes.aws_az_name](data-sources--aws_tgw_site--reference--group-001.md#canonical-a093ac2c1981421a35bde14f3c3f2297adca08fbc0d33ea9d9e3ce82fb011e41) |
| `aws_parameters.az_nodes.inside_subnet` | [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-367640f658465a59d0a6cf90493040a902a2b57727858269535452641f8b28c3) |
| `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.inside_subnet.existing_subnet_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-5dfef55b61fd03397fd08485f3fcca6b052a243d309e142d9b76c0af8990c938) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param` | [aws_parameters.az_nodes.inside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-fa1c1585c9c18e4d1f1b1e5f36c251e933a8a162bd6783d88518bd503397c424) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--reference--group-001.md#canonical-bfb5e73bb366f45f82cb20eec32db08c5604326886d353f4a1e397b60b6d46c6) |
| `aws_parameters.az_nodes.outside_subnet` | [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-cac0c3dc598ea0f0ac491ebbd90a76874b057534f5e8d9e201b06fa542587c0f) |
| `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.outside_subnet.existing_subnet_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-e5ae1d93f74386581aefd18c0a724644df651426c2c8589e90099aad87cfd10c) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param` | [aws_parameters.az_nodes.outside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-5ce9dd41425ebdef288463b9f349e5e20fcf6a39525215614c67cfcc854ad05c) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--reference--group-001.md#canonical-426685a193734c756187149b8741e0ff6026e292eabb699db8b6ba6b18e66797) |
| `aws_parameters.az_nodes.reserved_inside_subnet` | [aws_parameters.az_nodes.reserved_inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-a2e534f4996fc48abf0b04e943da5b4fb8fcaac2d4915f88b58c325b88bf7a8f) |
| `aws_parameters.az_nodes.workload_subnet` | [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-953bae85b93c126fa2024ba7af826cb0ec7aeabb2e94d75c8bd5a09866f82d02) |
| `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` | [aws_parameters.az_nodes.workload_subnet.existing_subnet_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-bb61dbd927df321946313d17772d6036c665375caed8ea3a8d4c4e594525e044) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param` | [aws_parameters.az_nodes.workload_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-1ef20757846b40763cc41ebe6e255a0af06591db1daefdf026a241095f3273df) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--reference--group-001.md#canonical-c6544431c63a5aa3b0aea28ce44fbee7bfb41fb708dc0ae29805243e0dadef89) |
| `aws_parameters.custom_security_group` | [aws_parameters.custom_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-254cde3f8697513bf6705d9c062ac1962c6b8cdfafd6f8cb9d02fe4549760877) |
| `aws_parameters.custom_security_group.inside_security_group_id` | [aws_parameters.custom_security_group.inside_security_group_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-508be9d27736d9dc5d61ccd51643159181edbbb0f708b16a8fc7dd063d2226a6) |
| `aws_parameters.custom_security_group.outside_security_group_id` | [aws_parameters.custom_security_group.outside_security_group_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-83623756b5c51f7dcffbeb84055e61cc581acfb73a55f3497392becc791d6eba) |
| `aws_parameters.disable_encryption` | [aws_parameters.disable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-5084d065d2220181f508a8b978fe9a624818957cb63040edae30057de73412fa) |
| `aws_parameters.disable_internet_vip` | [aws_parameters.disable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-b19ce2d7ebdc7bac9a0205e43f4f5cac1c41ec3772106f59bd31c78cc5d3a238) |
| `aws_parameters.disk_size` | [aws_parameters.disk_size](data-sources--aws_tgw_site--reference--group-001.md#canonical-c67841f7914a909ac77c7c2f373bf4e2d0fc12448a1ee0c431273ab0d5ead3e5) |
| `aws_parameters.enable_encryption` | [aws_parameters.enable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-4b01140c1e2eb68110f85355f798063e02818959ee3a03bc06e867d0f590aad7) |
| `aws_parameters.enable_encryption.kms_key_id` | [aws_parameters.enable_encryption.kms_key_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-6b9486d97e651cb65cce97f9cb5c903e8ab362c962b66b073913c90f74bcd386) |
| `aws_parameters.enable_internet_vip` | [aws_parameters.enable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-36841cb8f2076efe557bc13aaaa8788453d36bf8b6e9dd52beb25acdf44d81dd) |
| `aws_parameters.existing_tgw` | [aws_parameters.existing_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-f194ae324b82d34e0086dafe19c745c2d80556de95fc17e4770ecf2933b8e680) |
| `aws_parameters.existing_tgw.tgw_asn` | [aws_parameters.existing_tgw.tgw_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-8a571f0fb53ba8fca48c7b80f8fffd34109e808d86475150faff0ffad9a89cff) |
| `aws_parameters.existing_tgw.tgw_id` | [aws_parameters.existing_tgw.tgw_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-2e3685af6971d03b71b3c9b9bc774357b39e3cac1dd111b99fed4fb512e1e06c) |
| `aws_parameters.existing_tgw.volterra_site_asn` | [aws_parameters.existing_tgw.volterra_site_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-cefc7d17d1229641e8241879aa780ae337dbae3d03541889e7d9a8ac48f4295a) |
| `aws_parameters.f5xc_security_group` | [aws_parameters.f5xc_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-3f9ba2046da1f8ded40a732e09c45e4df1a7fb8685a1b2b580328e987ce12198) |
| `aws_parameters.instance_type` | [aws_parameters.instance_type](data-sources--aws_tgw_site--reference--group-001.md#canonical-4c1b421db516c58fe0ec84977bf0b54b1f6848fbe135dd6b74de78ab3fa549d4) |
| `aws_parameters.new_tgw` | [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-cc0802a381820efad57a1e1ef762d4ca5bda3385cb3f9542f02c82f79a9f79e0) |
| `aws_parameters.new_tgw.system_generated` | [aws_parameters.new_tgw.system_generated](data-sources--aws_tgw_site--reference--group-001.md#canonical-0087dde7c8d7547be4b723dc6e593188815cd466a32c639e2a26c97c829872e3) |
| `aws_parameters.new_tgw.user_assigned` | [aws_parameters.new_tgw.user_assigned](data-sources--aws_tgw_site--reference--group-001.md#canonical-520a902fa30b9ec99f7046ac8b778c19b62cf5366c36447d8515cad5775c230f) |
| `aws_parameters.new_tgw.user_assigned.tgw_asn` | [aws_parameters.new_tgw.user_assigned.tgw_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-56a447e649734880658ec3b551dfced844778533d91d45f3c5f7920b00a3e607) |
| `aws_parameters.new_tgw.user_assigned.volterra_site_asn` | [aws_parameters.new_tgw.user_assigned.volterra_site_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-795b3ec8bc1b1c9266a6bd022986d8cd1c2eb1e26cc5ad4e3de5d226f2d8662c) |
| `aws_parameters.new_vpc` | [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-16b56bd50f11ce4baf6b405c4f8c106dfdbc3f6002b7eebc6efd0eda2c42ebd2) |
| `aws_parameters.new_vpc.autogenerate` | [aws_parameters.new_vpc.autogenerate](data-sources--aws_tgw_site--reference--group-001.md#canonical-be5e52acc93b86578b572c18119b32c142190908de72976b54061ba62abbb063) |
| `aws_parameters.new_vpc.name_tag` | [aws_parameters.new_vpc.name_tag](data-sources--aws_tgw_site--reference--group-001.md#canonical-85a979e6bcd3e87797b1f0fca7809fa9a3550963b234653e5c6c2e03597efddd) |
| `aws_parameters.new_vpc.primary_ipv4` | [aws_parameters.new_vpc.primary_ipv4](data-sources--aws_tgw_site--reference--group-001.md#canonical-019370d02e077cd8bfb540ab198e18da745e4a61b7b45f921b0d38d7838322a9) |
| `aws_parameters.no_worker_nodes` | [aws_parameters.no_worker_nodes](data-sources--aws_tgw_site--reference--group-002.md#canonical-e30f89b165d927bf08e064dfde580f599ea96128813da4db04b1745441b79a5a) |
| `aws_parameters.nodes_per_az` | [aws_parameters.nodes_per_az](data-sources--aws_tgw_site--reference--group-001.md#canonical-4d0ef69002946c60e37ff3f4014db5e54810e3b52e24a1087d5e74a8f8af7fb4) |
| `aws_parameters.reserved_tgw_cidr` | [aws_parameters.reserved_tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-c3b18a6b8a4647502ae7a09ee260cf5ea0cd89b6ba948a14639692325ced6a55) |
| `aws_parameters.ssh_key` | [aws_parameters.ssh_key](data-sources--aws_tgw_site--reference--group-001.md#canonical-a6ebda838e9e80b9bd70e6aa18652a86692d69d1fc1370ded0379b06af984d11) |
| `aws_parameters.tgw_cidr` | [aws_parameters.tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-7778f101a01d68d4e5271212e47a75a41b551b95480eb21dc624dc4c23b99a29) |
| `aws_parameters.tgw_cidr.ipv4` | [aws_parameters.tgw_cidr.ipv4](data-sources--aws_tgw_site--reference--group-002.md#canonical-af703f8994aa3e399a06aa008d8ba3271f15f3237d0f7cc393f3172cda1419ed) |
| `aws_parameters.total_nodes` | [aws_parameters.total_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-03a6af61a13ff9f9ac1752c3be013a569c4ffca3106cc0f16cd707528f5ba7d0) |
| `aws_parameters.vpc_id` | [aws_parameters.vpc_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-94f24fe2a39d83fd49219566da999520dd5c8c1258c02c51e072f3c7a023a384) |
| `block_all_services` | [block_all_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-3766da8bfe172b4cc1e1bc73fcbae23ec559174988408593706501df993e626a) |
| `blocked_services` | [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-b62dd04141fc37bf897564aa63de84f0637a5c555d180bcc6616583d827974f8) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-7029df65a8cb874bce846aa5b1f20875846efef7432de0ce3adda190b8063fb4) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-fbee3b4603325d868dec895f2cfcb0101b076325ee4d98ea6811941f3f6009f6) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--aws_tgw_site--reference--group-002.md#canonical-53ce3745666b0448d2b8f501836796c7126722119ed140596199eee700ad72ba) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--aws_tgw_site--reference--group-002.md#canonical-58a69c59816e4085ab39d5cc99d047e68af2f19fa892ddd8165d2ffa2faaa039) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--aws_tgw_site--reference--group-002.md#canonical-188855fa2512378acd9d691f82cca8c8660261d38ba4ecae35547c5470e15769) |
| `coordinates` | [coordinates](data-sources--aws_tgw_site--reference--group-002.md#canonical-116c209ee9f3ff9064729c8470a1a02330f0d5d261e1beed0b980e47e810bda2) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--aws_tgw_site--reference--group-002.md#canonical-1c314c7394c84c546f70877bb0bf87ea288cbc8230dd79427910d61a1940a1c2) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--aws_tgw_site--reference--group-002.md#canonical-98c44c96859703fe20bdb66f20f7316c4fd8219afe88b713ce8f6f4865a051d5) |
| `custom_dns` | [custom_dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-a9ca6ea96be2929428fa0846f6f8fb3ccfb1699b671733f8e39713872d123323) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--aws_tgw_site--reference--group-002.md#canonical-f07ea26be9d798b733b6f5a02a80a2b19e612b9350f620d79f166b220791d1b2) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--aws_tgw_site--reference--group-002.md#canonical-327f0ff2d31699f4bba1d56da8fa905a489332e27136519de5ecc787d206f51e) |
| `default_blocked_services` | [default_blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-b6314feef46afa79caa60ae2ebf054c2aa71a7ae43a3b5a6c784ff9ffc38c2e0) |
| `description` | [description](data-sources--aws_tgw_site--reference--group-001.md#canonical-c9340580743f91e1f7378d8ab447525930792d67015a2e1e1a2a2203b9cbffa1) |
| `direct_connect_disabled` | [direct_connect_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-2715929d3672c3b0a45214c5ef899639c9701a7aa80cfb066c0dc41e8b4cadb0) |
| `direct_connect_enabled` | [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-d575fac80a60125bf86ae623a04a267c90979b7ec81b3be89b69110003c7455f) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](data-sources--aws_tgw_site--reference--group-002.md#canonical-cb34fd946b400df2ec53918b29dcec1126df1ecb009b0cd162c6aef8c487026d) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](data-sources--aws_tgw_site--reference--group-002.md#canonical-78262d7f24874bc49eba8b43ff761af999e625026738c612cdb5a90d8c5c7204) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-de3efc5cfa421a91e1e1738c1e58aa2ef35e0b0cdfc1bea18d9adf1fc738d28e) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_tgw_site--reference--group-002.md#canonical-b9443c781e21ffde21d112b27311e37cfe3928d33ba201a5b6264a3208d7c900) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](data-sources--aws_tgw_site--reference--group-002.md#canonical-bb8e7905c71e7391a25d467cb632ebdefc96d1d273ad0406efd43e99c3805494) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_tgw_site--reference--group-002.md#canonical-10dd478ce8fa5e6d2a364a98138c1878086b535e94097151fc8638e3ccd8c007) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_tgw_site--reference--group-002.md#canonical-3f64dba17eba0cb4419268a25b192d1a47ac01270c2696bfaf07dff39d8a85ae) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](data-sources--aws_tgw_site--reference--group-002.md#canonical-ad9058fe779b4eea6c40c01aad436a7690a5d3c6f0b99cbc7a4d12d4a28eaecb) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_tgw_site--reference--group-002.md#canonical-4583cefcbbe03be71ddcfc3291adf71bd7c024535c79599d8a46ff8ec5530d88) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](data-sources--aws_tgw_site--reference--group-002.md#canonical-8283cf35ac9e515ad68d82af5228a16da2ff16291f7c500e6c42a532c8ee5df8) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-0978ea2af73219e7971de06b7a0bd45f048606cb245e4ada1ad02b0721bef14b) |
| `id` | [id](data-sources--aws_tgw_site--reference--group-001.md#canonical-e4d0c041778bc2b2323f1046475a62b45b66f1a75866e36b4e34e733d271042c) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-5a6bd9a8cea5c719370e5a46ac3c04ad84c57a259dafb9f84e7c6c3465ef506a) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-499b12d741e8d2dbbd03e09aa44e3eb338a76d45d996dc23b79bf4a687d4dc39) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-f30f6ed9ecc6e25b0581c8feb78b3f4469575b782ebc10f8db32422b9657f562) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-8619c01352fcd773810d6319afe1b3d562f7c6ebaa455e4579bd0b316b0d12a9) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--aws_tgw_site--reference--group-002.md#canonical-47da0566726e55b7382f621ecbb4d862ea63b0ae62e122cf419648fa15e4e3d7) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--aws_tgw_site--reference--group-002.md#canonical-ec7cdc6e6af3bd3d46141cd18c12a00190c48e5a691498172e899be8a8fa0a20) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--aws_tgw_site--reference--group-002.md#canonical-2be1bdf063823133cfb66d37325d2558a1217a17657853b741105b6bd616504e) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-46f7218da63c1f4c82c539720d5b5b4aa56955a3f993f5b2cb1d3261d5e9cdc1) |
| `labels` | [labels](data-sources--aws_tgw_site--reference--group-001.md#canonical-4c20cce67d2db4c8b51d70cc468fbce0c4976673e4c4455f8cd4b038d0e0942a) |
| `log_receiver` | [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-ccb103b7a012db3466ce2d941b1cddf8a1907268299a3ce4d3832d992716e9fe) |
| `log_receiver.name` | [log_receiver.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-523db926650e303e0dbb923ccb9226bd1a838fa8a441b302818ad4280054a6a5) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-2e016b51a4a7af80de93a8376fd076989cd67f67e9372d742dadbcbf7a8bb945) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-08c287e826dde451d080d6e851f8ba6aadbf5c44791ac9867e8df37b5d58eb95) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-7536ec541197f206b04c1eed30a585d9154e8a9b2a600bbd4e7682484d02828d) |
| `name` | [name](data-sources--aws_tgw_site--reference--group-001.md#canonical-05893da18f07bdfd290a5878ecbed8951c3ba91fe0ddc201b3c8c59f3f946852) |
| `namespace` | [namespace](data-sources--aws_tgw_site--reference--group-001.md#canonical-2cc9dcc893e15b630308e49328ae7627a108dc56a2c7dba66762a4047bb7e2cd) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-19e928a07d49176f6b9a6bbf03a32a06ea8f5aa7069a3fcde0256747ece4b9dc) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1218404ce8869390987f7dde1b705515d102994104dad252d041c38b1e308335) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-adaab5c3c05e38d546c436a9005fdee6a3124e9c89ef62b193c3f04b12aada73) |
| `os` | [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-6e1bb0da18cfe7dba8df326ece77fd8dbef791a71c9c8c0abe2cea9527e6d431) |
| `os.default_os_version` | [os.default_os_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-8b2f3a87e50e62e60cf8a4f17a2cce0f1b9ae70fbdd4624f5ef04d0bd3d3a2a0) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-c0396ccc426480ca0f36c179d5f313829ef9c89ff186778e461afdd2faeaef44) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-51219babe41431c02bf82278a1764fff874295b28d1e2807797242a0e8ba2fb0) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-ccbc8e229b639276f1134fa336198006883c2545dc86a902daf54498bea4dab5) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-58400fb527c8d77e9c20f3e67938d5d3b6f9521e61d9e4109d40ab0bd5fd1a00) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-b3e615db345b89d00ac7ab07a63556222a5a0be53b9070e3e701b37a463fe48f) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-37bd473846a561d8788543e63479b57b98b35be2a1860d6b748983e2fb95613e) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-e39cb458a766bcc70951012b367660dd91309427bdc4cc8f6469517105843fd5) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-09c25ca368935d5f27a483dc1220b3df9d1adcf2982403aef0bc8e41560b8a19) |
| `private_connectivity` | [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-bcd73d00e1443e2a7e2664d4018e55d4d0a80b52e497dbad2dd1efcea11d19a3) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--aws_tgw_site--reference--group-002.md#canonical-0c133052e546f4444d0d7387031367147768509b0918d1a33a3a5aeec071d67a) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-54b34d7ccc9d3f215e5d030df7602a6c374cc681d0201f3a51ddfc4df0b6c320) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-5c2a44a2d9ed8c62e1f6311379ad30b7d43de2b152c26c7bdcb3213bb55fec44) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-330755e1900e7f1695b1f0b14599ef731dd6be97efd08423e03e916ee5899273) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--aws_tgw_site--reference--group-002.md#canonical-9702074bcc3410048d33f3af5ce357f9417a85a8d6224802abbe55e51d4fa748) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--aws_tgw_site--reference--group-002.md#canonical-74efddfcb79ed45242980b2b32bee58b81151b56c5f6802b7270cfd8e05adf63) |
| `sw` | [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-d900a93a59f92e75bcfb4981c8b1813fef63aa5de931419f7614ee7250ef3799) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-276a85c797a0ac0760bb762399c4c3e9c878776934159585cb776d5d18b29428) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-3c27282eb49626e67d8a15d89d4804ab743e98f9733c98afacc1fca7098f6450) |
| `tags` | [tags](data-sources--aws_tgw_site--reference--group-001.md#canonical-322d346a9d567023210f03933cbbee637008d53bb4973e5d2e856c0c746d0265) |
| `tgw_security` | [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0b40ac3270c53733c920e23a7f781b2793eab0f88fa478d95f0ac36fd2577d75) |
| `tgw_security.active_east_west_service_policies` | [tgw_security.active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-7c6649f174abb3a2fbd35488a055020e1cfde551ad39d64c7487e05499199395) |
| `tgw_security.active_east_west_service_policies.service_policies` | [tgw_security.active_east_west_service_policies.service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0bc22b1e675c25a562d74bde36afc887ecceb601c07219555c182e48d2d5539f) |
| `tgw_security.active_east_west_service_policies.service_policies.name` | [tgw_security.active_east_west_service_policies.service_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-b092bedef709ce214c8429ed33c61041d5c03638c8755c1129a1a4132ee24817) |
| `tgw_security.active_east_west_service_policies.service_policies.namespace` | [tgw_security.active_east_west_service_policies.service_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-1caac7596384ba6204ab24249db760f0becd7ab776227d054d5179475f3aad4c) |
| `tgw_security.active_east_west_service_policies.service_policies.tenant` | [tgw_security.active_east_west_service_policies.service_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-0a0e3feaa47275b870e334c4898d5e833096d8d9051bd7049a0e4a8357d152fc) |
| `tgw_security.active_enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-d06b06e12f91147e8e7f4fa130e2bfbce0bf3fc94c504a97394f7e62fc8b8f24) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-f59e1f9de76d9f3d706752e47eaf87b12aee0a9615a07e1acbc1db8d4053fba2) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-c4ef60e36cfa6f1ee383747347e73b72736844c95d8448ba3cd4ad94045e9048) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-098707924f07fc4c90ab319086c1841fc0c15f622658bee9800549ee9c656ef7) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-bd4dc13bdc55e83c6859626c77446f2c3a6f8c22fa8329df262c63849484d934) |
| `tgw_security.active_forward_proxy_policies` | [tgw_security.active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0f4b9f8dbe4b166b37ba52f47fe6a9eabb9332506e6915b2807ef6a627621281) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-96e6fbf67d24fc071949d502029edac092544902a440681e007ab95af5c90d13) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-11b97357f34f98b6f9d4cb0124c9f157f80ebfb589ab8fdf2181e2f68799f1a5) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-b8674e6e5571d8dc5ac77b3e00d2f71a2c51ee451e0cdeb7c6dc92e05c05cdfd) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-8a0967fa01ee41869a74a2cf8cefa3fbeca0c8addd35f5519d3669f56dce533a) |
| `tgw_security.active_network_policies` | [tgw_security.active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-6acc2b74a1277980ac23ce7f134f1ec47583fc1d83fc4ab26f8499d3078a8b78) |
| `tgw_security.active_network_policies.network_policies` | [tgw_security.active_network_policies.network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-59c91c012206d2565fb8ee8fa6ab677707a5f73e6259f8181401ce5ba76566f7) |
| `tgw_security.active_network_policies.network_policies.name` | [tgw_security.active_network_policies.network_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-aa4e134489590bbbb62cda06068bc0dec5dd8fdf75b44213cbcd043e8fb8edd7) |
| `tgw_security.active_network_policies.network_policies.namespace` | [tgw_security.active_network_policies.network_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-9c284d2e079f1ad67f2e08f9f03304b7fdd8fbf37488eb69b55c785f09d04ea3) |
| `tgw_security.active_network_policies.network_policies.tenant` | [tgw_security.active_network_policies.network_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-0046bd6cd61e9c99c3b7697a847a875a934286cf5c22c49798fc79f898b83e20) |
| `tgw_security.east_west_service_policy_allow_all` | [tgw_security.east_west_service_policy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-fd3ec50d2c10429cd1443d4615684d7daeec7ce0f9adef8172cb201382d3b1cd) |
| `tgw_security.forward_proxy_allow_all` | [tgw_security.forward_proxy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-1694b1f6fcae807b2e75f392aa79f0d2933d3ba2c39de90618e169d7cfe1f768) |
| `tgw_security.no_east_west_policy` | [tgw_security.no_east_west_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-bd0a653e80435640f5d321416979b093711e8fb490e3caa5de1364db600b3988) |
| `tgw_security.no_forward_proxy` | [tgw_security.no_forward_proxy](data-sources--aws_tgw_site--reference--group-002.md#canonical-4bd6d3770f44b98bfe845052513d9bdaefcdbba935f67265247824998adc1914) |
| `tgw_security.no_network_policy` | [tgw_security.no_network_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-4fb31cd84d99cf035ef69af63dd6df1a3ac8ebde55f4d6dd9a735d142edd15e5) |
| `vn_config` | [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-5e7267df70b77cd2c6d68a4cc265fe04d7674c2569452eab997aa2fda735719a) |
| `vn_config.allowed_vip_port` | [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-002.md#canonical-89dc9ceffbec7deb3f90f5846e6f4081aae417ec48229aa2004108cc42495b7b) |
| `vn_config.allowed_vip_port.custom_ports` | [vn_config.allowed_vip_port.custom_ports](data-sources--aws_tgw_site--reference--group-002.md#canonical-f73ed2fa11662d3cccc3ee7fa26dc47d2af26c76482b63a66b9b485ee1c067c2) |
| `vn_config.allowed_vip_port.custom_ports.port_ranges` | [vn_config.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_tgw_site--reference--group-002.md#canonical-365f32dd32b3ec329d7019b3cc4ca684ff49376e81883063601b7afb43388b9c) |
| `vn_config.allowed_vip_port.disable_allowed_vip_port` | [vn_config.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-002.md#canonical-33e4932ccac95eedff94b296c938e275f3f930c0407bb5afb21e05db144c5bea) |
| `vn_config.allowed_vip_port.use_http_https_port` | [vn_config.allowed_vip_port.use_http_https_port](data-sources--aws_tgw_site--reference--group-002.md#canonical-df1d6134f4613791d9850bda2eea504f5818524b2111fbceaa346e32ec3edd04) |
| `vn_config.allowed_vip_port.use_http_port` | [vn_config.allowed_vip_port.use_http_port](data-sources--aws_tgw_site--reference--group-002.md#canonical-bf6fa266462f7a5b735b50deaee7d5a8e084a096171e0ae1af7a338d101122c9) |
| `vn_config.allowed_vip_port.use_https_port` | [vn_config.allowed_vip_port.use_https_port](data-sources--aws_tgw_site--reference--group-002.md#canonical-60e50673b17b1e1b8c14ccbe62fa7771a945ef3d96cba8951a468c4edee1757e) |
| `vn_config.allowed_vip_port_sli` | [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-86114df20d333a69267033a41329347c5917e8e4fab5f068240efd85a199e896) |
| `vn_config.allowed_vip_port_sli.custom_ports` | [vn_config.allowed_vip_port_sli.custom_ports](data-sources--aws_tgw_site--reference--group-003.md#canonical-9e8db0b2ca68c49489ed9ce4c305ff6b83b561045c87db275c4adab5a8b2796b) |
| `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` | [vn_config.allowed_vip_port_sli.custom_ports.port_ranges](data-sources--aws_tgw_site--reference--group-003.md#canonical-a0097a7edab31d767fcbbd9597517b6dd499163a5c7607ff194e4077802a3e5a) |
| `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` | [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-df2e9c3a0ab5d181c8f5cefda3c71c540726e9b585a967e9518415a1d644b743) |
| `vn_config.allowed_vip_port_sli.use_http_https_port` | [vn_config.allowed_vip_port_sli.use_http_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-18e65baba2dd97984b46eee7f81f110c93dbaa6669ca3c2bfe4aece4213a0419) |
| `vn_config.allowed_vip_port_sli.use_http_port` | [vn_config.allowed_vip_port_sli.use_http_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2582d626f45d902c0835897f699be2bf9ea5df58d3564530ab63d0c3904f8f80) |
| `vn_config.allowed_vip_port_sli.use_https_port` | [vn_config.allowed_vip_port_sli.use_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-71f76f1657caa45ff53300830a5a0713f98932198d2c452dbb26ce6a62e9449d) |
| `vn_config.dc_cluster_group_inside_vn` | [vn_config.dc_cluster_group_inside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-07a495ad6894a04082761fe9a7e87646827c3aba90ff1e864cbd384b219d98c9) |
| `vn_config.dc_cluster_group_inside_vn.name` | [vn_config.dc_cluster_group_inside_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-1c7061deb0bee5932cacd13b179aef2f77d52916d17e2e8ea8f33762d978cdb7) |
| `vn_config.dc_cluster_group_inside_vn.namespace` | [vn_config.dc_cluster_group_inside_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-b338afe8f8f729e23f5430480e947ea1052c3f556feef7c9826ed8ffe8f2fb71) |
| `vn_config.dc_cluster_group_inside_vn.tenant` | [vn_config.dc_cluster_group_inside_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-8f946d5c0338a5345e2bc37e35e44f8e73b3d9ac2014c8647d5c1c12cbc45fd0) |
| `vn_config.dc_cluster_group_outside_vn` | [vn_config.dc_cluster_group_outside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-2f7c595e9d8e123aec5deacc82185a031b73a22f98562ad758131e0c88110b46) |
| `vn_config.dc_cluster_group_outside_vn.name` | [vn_config.dc_cluster_group_outside_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-9f4094dd7c3171c0f83ea52e2ed78c4d77a84e98dd9a2fe35931a04af21af16e) |
| `vn_config.dc_cluster_group_outside_vn.namespace` | [vn_config.dc_cluster_group_outside_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-8d8624961aac15ee3efa59df58d5576e4af525de215b7eab987ecad07363121b) |
| `vn_config.dc_cluster_group_outside_vn.tenant` | [vn_config.dc_cluster_group_outside_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-8cbeef19c903b9afdc8cf2f1826fd7c8ed1bdeabaea2d6cc03194b89f7e15b26) |
| `vn_config.global_network_list` | [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-19d95146fb7ff7e344dcf1febec3f9f7944a495af462f8785a9a0c5cf6472991) |
| `vn_config.global_network_list.global_network_connections` | [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ba255be613fdae79944d031ba220f074560c4e3b5e77ab84c8d0e271597a63a4) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-510d66e0eb26ac744a950f15b1d48f8a63f5957f5068d37401d79f5660a9d232) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-2116b8fa7f9acf945028f6db8bd39335d521829e52ddb15ce328f244d486f69a) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-29e5a2cde78b1c2b0799dde611a9f9a393536a808db1e9a04772ae8d91b08a00) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-17212703e6995ee7f1bd649a885c858de288bb67d936b5b135cf57fe097743f4) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-eaa69a0b13d951c2f19e2daa3cea24c6f6d5ad0b1036652e93fe56cc2f650255) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-ac51addef5cf435bb85972d4b82d0b5840ec976ccf9a73cf1324c4ced8f2cab9) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-178efda810336a5ac12f165c4f82ada8c554933cb1932ee53c1838458fb2a8bd) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-1724bc704cae11616e438d898b66c68d30a025d7a5ac652e92407ee9a497a0d5) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-963bea0b5f37cb209d77e36af97ef3f76fdb380188ca26f0bd308d8928d6335b) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-ef760c3856b8d7cb1b2ba01d517fd016a1610f2514c6b267463f23b69f34969f) |
| `vn_config.inside_static_routes` | [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-77395b6a6d1b2077e2a00a4c4fe80f620c5109175a808bb0feefcbd75ab195ea) |
| `vn_config.inside_static_routes.static_route_list` | [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-38085c9a9a136e3727de8ea4db5a4bdecec927d3c207bba0232ddfffb2912686) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route` | [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-b4255be746d21e9ed6b08e0b35a042e1df40bca48660545fe7c12f98caa72555) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_tgw_site--reference--group-003.md#canonical-7e00f4cb2cf22dcb0ab3089a6fe8a8bcb096821f34af841c6d1f1351569e9b89) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-9069cdda869b1fe17184519b84a4132f1420d6e6510a6b8c47b5b3ce0afb5e67) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-d844ff7876b3a4e16e74b13b4f551e8d81572abc8ad5af2ab0618eee7e0f9ff9) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-bcaab18cbe117d2aeb6ae9fed6826f280a6cc1791c193b3ca965e08f3282543c) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_tgw_site--reference--group-003.md#canonical-6a4b44118e914655bdf6bfab8f74e5ad94c2d7858d3939eaec13a20337800f55) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-c65ea6904265eea7906d8664bd6593111f6732c56241ab0e5459e4295865ff97) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-d1c533fa3087382233a6443a246466df15bdcb3ec44a3d9c964827139238b2a5) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-64cd0fec1eb02bd38deaee2a55b57f867a678ca2bdbf16c54d5f418e84c2cb94) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_tgw_site--reference--group-003.md#canonical-74f5a08924f780c8188c773a93af11ae384c9838156f9c56c5a03923d538fd7a) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-47ab70ce259ee88e9f92ad6411ab8af9204dcd63798e954b520a9370b3c99815) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-e6510a9e39310b7d8e64871fd3bfbd899048c86a8022bcce9ee94a6d47cd5fe9) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ed1d2721c33e535342aa72d9127c462565b0436b6534d6b7f6b154ca126ae3) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-58948e3494d73ed114b5bd32289ed06f42bfa4ea73437b49321da44c01546ab0) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-aa45d2dd132d13bc261310aebec11b9d2256461a6a4c86951490ba93adab0545) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-85a74f1b64336852f64711f242b90875854dd0cbb237327134090769d866454e) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-dcfe85c3a29229032a6e3af58f765803738dda3367a0282ed947b50fcd0981bf) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-89632ca4de1d22027116222352b7d46014bf05c0db7b0c34518a944bc5353625) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-098f6bfd4ff7eaf789d3590e54b8852d0e214a89609b804ab123bb1fbced3da5) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-deff29e8bceba55cde7ff455bce3c7098143da6d5693222ffda77b0ed322d1f0) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_tgw_site--reference--group-003.md#canonical-34ca86cd22203aec34245d4b7cc1259937b97a3decab967410b4b09e8f10f8f4) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-e01dea2cc991d5254ae6e0cca8a5bbb84e5ecd6fe7a11eb59154755d3a88a3ea) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-72a3f012479b3f64f6eaa51459dd0e569882e2870688b4d3b2fb254105999e54) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_tgw_site--reference--group-003.md#canonical-4ba09d1fefd192e35bdde26e127b4f917ea00b7e69169593919181bd870c2971) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_tgw_site--reference--group-003.md#canonical-6cd5ea2b3284f4da4ac2db8229f9dca2e531c9429903aee90f20a3b784654b3a) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-784c97cca1b5cd990e897e7c22f3a881e4d2737286ab6711b9478f2537a1010c) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_tgw_site--reference--group-003.md#canonical-9f9c58234560fd4eda94dc0e660bdc0bd149fbfa3dc809b526b12692a8283c1d) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_tgw_site--reference--group-003.md#canonical-1b08f9c4ee3119852207fc5ab0c22e6c34a8c672b79ec71fd1ab6e0ed3028660) |
| `vn_config.inside_static_routes.static_route_list.simple_static_route` | [vn_config.inside_static_routes.static_route_list.simple_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-0028e82b9552f84599d5eb24689043c6224ab65cc2104dfcc96ff0918860ad36) |
| `vn_config.no_dc_cluster_group` | [vn_config.no_dc_cluster_group](data-sources--aws_tgw_site--reference--group-003.md#canonical-6978794e9f73e35b082bd014485a0732a9f0c6f4a003e7a1fb7e7e20823fa2d1) |
| `vn_config.no_global_network` | [vn_config.no_global_network](data-sources--aws_tgw_site--reference--group-003.md#canonical-2e70c28e74d3d84bfc1a13aba5a46cce722710230d861c949626e86c16724d7f) |
| `vn_config.no_inside_static_routes` | [vn_config.no_inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e59e2f2fc6dd976c8795649b8d25aa62451ca5093ada723e77bdbd8738525058) |
| `vn_config.no_outside_static_routes` | [vn_config.no_outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-13725b68d17fcc374360eff82993dc4a0f178f1f2129ff4bc62c27557f1f81da) |
| `vn_config.outside_static_routes` | [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-f6f7be2e5b45e94a34a64cfd268a67b4b45d00c3ab5f6e632f95cf0b2241b25d) |
| `vn_config.outside_static_routes.static_route_list` | [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-17bf8c248d588df892025a7d223fa406821a7fada821333d5c0a40d7dea51e77) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route` | [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-a799550dc83d4c1feadab4463d9ea02630adc6f321e81987f44e189f1d4d6c11) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_tgw_site--reference--group-003.md#canonical-01cf3966c1567ce2ea27048453b53a656005ec6f97d7fa19d789fad261a392af) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-608eb731ef84811cdcf6d89f7760720015133a9702dc278e18fb450f0d9abcae) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-33cd500dd4a4cf86e5caca5d8ba721f8fca5f772df6a127ba99cf7df5ffbb8fe) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-62f231cd56dfd0e515cdeaf0854cf85079b39495c125ee2534e43d5f591a6b33) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_tgw_site--reference--group-003.md#canonical-5ac00e2e95eebf245814692708b32b74c433d187daf0865e87d80adb712ae01d) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-431065f1a56cd191ed8a6c94c966c6a59b24215ce7cb387abc36cdd081629475) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-6357f8a6ae93b36378dfc39bfb130abb1bbf951e3534a8411d657c413e68ed28) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-75acb681d3d734b1f29a1736201faf04b277988fb51445195b73cd94e138a114) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_tgw_site--reference--group-003.md#canonical-114486fd1c13e3f46824f44ffe9b42220d23b83ee03fd8cb677133319dc10ca9) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-afbee053028e331a9b9a34afe97677ca96c2d4b0c6553054ae6ecf482435c839) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-230b9464ad6f42ae941927e1fb7070d3c510e939892fcbd18e056c980006225c) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-34b26846d372d405a551bae592fa1971e6ba05a941ac1abcaf3a432297aac262) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-ab4ca23c72524d8bbbe482c9b5b800878716857f31ae6f3bd0861701579bc67f) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-694346bb645b208d5bffa02256a1e4239d88dc8fdad162ea58a8d25047fc2b5a) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-b8e7b8cbb6865243bb1d468576e64e19023e6162a8ae171afa3ddf8b890407ea) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-049644de7be974ec2f8a6f012b0ca586c50ce02991851156bb1e17cf31da0eb8) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-61ed817f48c82cb1d03628f758bc6cae9013d03b52877104a1ddc0a54404515b) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-e8524d59ca1a2589da615a1d8b801491b528eea7c174b9d0cb9f1ac7d1ffc2d1) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-e159d9239d4a539a927fd0a8037ce7cf1bcd691c030aeffed118ffbac69b827b) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_tgw_site--reference--group-003.md#canonical-3e638279180a6560392271bff24b20a2c3a46ee4a1348ade28257d8a2640cde9) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-257a7a9bf7672c034fcb584f91c93f7b6c15687e3a59fc313d97f79488b11eba) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-947838441b2543891d4cc5e76533764d64d49ef410211c0b8c831a86c99acae9) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_tgw_site--reference--group-003.md#canonical-d1322a713490a10ca60ef016badf9b194c7a63f2a0f7f5496cfeeb6c543894e7) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_tgw_site--reference--group-003.md#canonical-421022fea38be249e8397b96885c2ac47976114a9f26c6ae758ffc632ffb7ee7) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-4742df12e337e9d413117b4083bc236cff530ad764f079401da931c0bb5cd79c) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_tgw_site--reference--group-003.md#canonical-7e774c8b919b9a72b5e8b4624b2001c0b4cc4aaeb8f2008673fe19168d2b6d33) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_tgw_site--reference--group-003.md#canonical-40995b42f8baa88050c1865e3ec47b5f2fe82b8b900b7c1b220dc2eb7698cdcf) |
| `vn_config.outside_static_routes.static_route_list.simple_static_route` | [vn_config.outside_static_routes.static_route_list.simple_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-32d3678a730ba3bf075b4615917a6a5a4b6e89274bb0709e71be35cd38becf58) |
| `vn_config.sm_connection_public_ip` | [vn_config.sm_connection_public_ip](data-sources--aws_tgw_site--reference--group-003.md#canonical-50d3cf912cfaea739a7162aa74e95e3da5282f223ab577100a1506394880c9d2) |
| `vn_config.sm_connection_pvt_ip` | [vn_config.sm_connection_pvt_ip](data-sources--aws_tgw_site--reference--group-003.md#canonical-031c93baa4c02b1e6d0859ccda05cd8ddd86b0f57287727d6be962a578e3127b) |
| `vpc_attachments` | [vpc_attachments](data-sources--aws_tgw_site--reference--group-003.md#canonical-8e6516a8c16e051650a02da73fb22e1d1c2ba7b9f3bbbaeb4b90e1519cfa4d43) |
| `vpc_attachments.vpc_list` | [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-a09dee683a033d7161a8faea75075e45b4e60f0c239204b6225ace0cd6ec5dc6) |
| `vpc_attachments.vpc_list.labels` | [vpc_attachments.vpc_list.labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-d56f24b645f0a3f71bb0d60fe00ac8d648fa5b80898c10f1790f8d71bf2958c7) |
| `vpc_attachments.vpc_list.vpc_id` | [vpc_attachments.vpc_list.vpc_id](data-sources--aws_tgw_site--reference--group-004.md#canonical-03200909b2d67a7d6a89cba0b72e76e1afc4f4a7a40607498f0f13fb676a0ef1) |
| `waf_signatures` | [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-8267bda250f978b920e3a7bc785efd78537994e0f9db1cbed1c9399489be0684) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--aws_tgw_site--reference--group-004.md#canonical-73bdd94497da37556150b81291e358fff53c59a6811b69b53ae9180233ecf461) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--aws_tgw_site--reference--group-004.md#canonical-620c2c68cecd3fd166330a8ebc02629b1c1ae990fd91f5fc227b5a2b136e5877) |

<a id="canonical-f6b020c42300cef5067fdbbf9c35738d9996f113390755917bd0010804f77436"></a>

## Next pages — Property reference / 53c4f33afa42 / 12

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [block_all_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-d97df5820ffa9083872d636b6f12ccfba661e8ceb4ced94125bf20220097cbf6)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-9065ec37d6ec0b7dfacc5a66d3ba150f8b71e54a085c55e401b3d26e0a27ee89)
- [coordinates](data-sources--aws_tgw_site--reference--group-002.md#canonical-f88da05b7925af8930063e804ed423f19943a3a16a6194a1a4426d8c4bd499f4)
- [custom_dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-45aeb23832ca3dbd757bdef1e8b6d6caf4010551ba9796e0e81c452aa73c6471)
- [default_blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-667795826786ba0811684ca6551f464c08907f534c98d05ab49463b40ce4439b)
- [direct_connect_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-1a0d78279f45609ad32ac2a33f4d1a3dc9c80eefd47cbaa48fc5f984d0320ac8)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-365a515e3f4c937f5dc99dac97b4ad775620553cf7b15c224d3cfe294ed537f0)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-3a4e0518667a945acda23b4526ee7eb358936563b025790a8f5da140ba6136ac)
- [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-88271efdc13eb8fbbe3ea8a8f53c3e724d1e1af8bffc05f7401a3babc291c965)
- [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-57257a31846eb027c65f9471e86ce1b4007facd84f738bbd709e95af6ee482c0)
- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-4244e169f7d7eb4f78609fb64c52633e88576fb2a39de494147f6f4d3c29e2e4)
- [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-5d94edace7908538b2948f16b2cb5560db576aa216e7674b5fb8d1085d8f999c)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-3ce8e047006358cedc3c59b57fc4b80562ac72af3eefd5b6af40ab014380dc3e)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-b1d3abf5c88782ae24b64db8240510981bb2c45c61bfbc2ead4671d31690aa3c)
- [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-2274ef167fa95f7b95100502348384d42db2101bdd2013e74f0674421accdbd3)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-2e1ea519eb9353ad3ea4668d9e236a9eb32c420e919b7afa630cf37e55e12526)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-003.md#canonical-411569971dc5e487a1b7eb5dd5026c6a951d476dc5d2fda3bedd09ca3b953633)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0f930540b03ce3ec08b48a152be09afd0635e044541923a65add395e816e0f26)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae5b82737623bf747ae514b4b8dae1d9e30db884efd1bd9f95d5e4357ae37820"></a>

## aws_parameters — aws_parameters / d1a8fd58e181 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- aws_parameters

<a id="canonical-fe5668507c4b696604c2d7479fcef313aa239e41af57f2f4e596596174b17227"></a>

Type: `"single"`. Computed.

Setup AWS services VPC, transit gateway and site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deployment": "[\"aws_cred\"]",
  "x-ves-oneof-field-encryption_choice": "[\"disable_encryption\",\"enable_encryption\"]",
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]",
  "x-ves-oneof-field-security_group_choice": "[\"custom_security_group\",\"f5xc_security_group\"]",
  "x-ves-oneof-field-service_vpc_choice": "[\"new_vpc\",\"vpc_id\"]",
  "x-ves-oneof-field-tgw_choice": "[\"existing_tgw\",\"new_tgw\"]",
  "x-ves-oneof-field-tgw_cidr_choice": "[\"reserved_tgw_cidr\",\"tgw_cidr\"]",
  "x-ves-oneof-field-worker_nodes": "[\"no_worker_nodes\",\"nodes_per_az\",\"total_nodes\"]"
}
```

<a id="canonical-2ddab3f0caefd5efccf242f058e161523425fc1a7cc85a89417cd2f5e0409a89"></a>

## Direct properties — aws_parameters / d1a8fd58e181 / 3

- [admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-427fd9512fafcbf2e602a41a442924d9b16b2a85b3965f597e03498ba4f40d81): complete subsection reference.

- [aws_cred](data-sources--aws_tgw_site--reference--group-001.md#canonical-a17fd7dc76c41860eb770c2e67e95585b5e165f22361b47f7d15c8da27b5798d): complete subsection reference.

<a id="canonical-394d08a7419c1a368f8a8bba3fb88609246d210ae0ce9acf5b56da4d5e495786"></a>

<a id="canonical-d7fbfcade3529b9dcda2f7e0cde245a9b30959337d3aa5a6bc5c7d26c5fae70f"></a>

## aws_region property — aws_parameters / d1a8fd58e181 / 4

Type: `"string"`. Computed.

AWS Region of your services VPC, where F5XC site will be deployed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2): complete subsection reference.

- [custom_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-60f382ec3a23f7ed7f880aa170917c5dcaa87378bb0d17d629f9258968a389ba): complete subsection reference.

- [disable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-beb858327d992a81f23bfefeff3d4f7b5b31abdd4b38af3f3c411fe48f78d9ec): complete subsection reference.

- [disable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-4ba2140ab942a6f4f0bd961b3a1e03755aaf8ef7dc6c8486634f349722c518ff): complete subsection reference.

<a id="canonical-c67841f7914a909ac77c7c2f373bf4e2d0fc12448a1ee0c431273ab0d5ead3e5"></a>

<a id="canonical-9c05db35b3c66a7083e3c65e6674103f3a896804eb5fecd8a696b4a1fceee8ce"></a>

## disk_size property — aws_parameters / d1a8fd58e181 / 5

Type: `"number"`. Computed.

Node disk size for all node in the F5XC site. Unit is GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-c4c8261093089e3cd43d2d8384e21ffbec226771bcb7fc88616dc90630abc2a7): complete subsection reference.

- [enable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-1a83b1e4b1d83b65d1d7185a98b222103b9b83073ec919f0add90ea2a77947f4): complete subsection reference.

- [existing_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-8a2bcf5e62c4068a1db1e653b03e1cf153d8621975eb1c64399538af270f6bcc): complete subsection reference.

- [f5xc_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-958d863e35e9da10743f095c325bcf760046b81fe41e2e8211017e5ff8270d08): complete subsection reference.

<a id="canonical-4c1b421db516c58fe0ec84977bf0b54b1f6848fbe135dd6b74de78ab3fa549d4"></a>

<a id="canonical-5df1bea42179d3369f4fecd7d5636c89170fa576ee51ffabc5bf105045cb0700"></a>

## instance_type property — aws_parameters / d1a8fd58e181 / 6

Type: `"string"`. Computed.

AWS Instance Type for Node. Instance size based on the performance.

Upstream description:

Instance size based on the performance.

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

- [new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-6ecdcdc30012a777ac3b186af4be6961737177618a676c001871e0b5fc334c53): complete subsection reference.

- [new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-a3d4cc21785dd411c6e93fdfc38f79feb4c8bdc2c791c0768ffa790c0cf3ac2f): complete subsection reference.

- [no_worker_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-2d42d69e895bc6d802c1728cb112a4af67cae50fca5e22033ac294bbaf84a020): complete subsection reference.

<a id="canonical-4d0ef69002946c60e37ff3f4014db5e54810e3b52e24a1087d5e74a8f8af7fb4"></a>

<a id="canonical-50b3a42e8614cb3b211f41583ccbfc9dc6c917e8509b9567a186eaf227dfee0e"></a>

## nodes_per_az property — aws_parameters / d1a8fd58e181 / 7

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
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
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [reserved_tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-cfe0573a85ccf3c134889c4ced0af096b34c0d143e14ec1b381ce226bd5179c8): complete subsection reference.

<a id="canonical-a6ebda838e9e80b9bd70e6aa18652a86692d69d1fc1370ded0379b06af984d11"></a>

<a id="canonical-641c9f1a6b00f9bf6c95c69904f72e46bf27fe19a0391442956e4ea12a066fe2"></a>

## ssh_key property — aws_parameters / d1a8fd58e181 / 8

Type: `"string"`. Computed.

Public SSH key for accessing nodes of the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-e225954e4e5d68f745371e584d73b44b5b5d096c0ccc5755683b8ec9c9532ec2): complete subsection reference.

<a id="canonical-03a6af61a13ff9f9ac1752c3be013a569c4ffca3106cc0f16cd707528f5ba7d0"></a>

<a id="canonical-b512255f02c8bad70cc02f37ab4eb4fdcae3636be908560390e4794baa02a61f"></a>

## total_nodes property — aws_parameters / d1a8fd58e181 / 9

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
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
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

<a id="canonical-94f24fe2a39d83fd49219566da999520dd5c8c1258c02c51e072f3c7a023a384"></a>

<a id="canonical-3e035faff7a84f949618c44829a82bebb4173e270c9b5ddd7fec29196dbf603b"></a>

## vpc_id property — aws_parameters / d1a8fd58e181 / 10

Type: `"string"`. Computed.

Exclusive with \[new\_vpc\] Existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Existing VPC ID.

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
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-0d1101ae4da59d28e31dab413a445cf815e304bab1bee87563eac43b11811e62"></a>

## Next pages — aws_parameters / d1a8fd58e181 / 11

- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-427fd9512fafcbf2e602a41a442924d9b16b2a85b3965f597e03498ba4f40d81)
- [aws_parameters.aws_cred](data-sources--aws_tgw_site--reference--group-001.md#canonical-a17fd7dc76c41860eb770c2e67e95585b5e165f22361b47f7d15c8da27b5798d)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [aws_parameters.custom_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-60f382ec3a23f7ed7f880aa170917c5dcaa87378bb0d17d629f9258968a389ba)
- [aws_parameters.disable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-beb858327d992a81f23bfefeff3d4f7b5b31abdd4b38af3f3c411fe48f78d9ec)
- [aws_parameters.disable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-4ba2140ab942a6f4f0bd961b3a1e03755aaf8ef7dc6c8486634f349722c518ff)
- [aws_parameters.enable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-c4c8261093089e3cd43d2d8384e21ffbec226771bcb7fc88616dc90630abc2a7)
- [aws_parameters.enable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-1a83b1e4b1d83b65d1d7185a98b222103b9b83073ec919f0add90ea2a77947f4)
- [aws_parameters.existing_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-8a2bcf5e62c4068a1db1e653b03e1cf153d8621975eb1c64399538af270f6bcc)
- [aws_parameters.f5xc_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-958d863e35e9da10743f095c325bcf760046b81fe41e2e8211017e5ff8270d08)
- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-6ecdcdc30012a777ac3b186af4be6961737177618a676c001871e0b5fc334c53)
- [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-a3d4cc21785dd411c6e93fdfc38f79feb4c8bdc2c791c0768ffa790c0cf3ac2f)
- [aws_parameters.no_worker_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-2d42d69e895bc6d802c1728cb112a4af67cae50fca5e22033ac294bbaf84a020)
- [aws_parameters.reserved_tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-cfe0573a85ccf3c134889c4ced0af096b34c0d143e14ec1b381ce226bd5179c8)
- [aws_parameters.tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-e225954e4e5d68f745371e584d73b44b5b5d096c0ccc5755683b8ec9c9532ec2)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-427fd9512fafcbf2e602a41a442924d9b16b2a85b3965f597e03498ba4f40d81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04723327d84c1f358e79d7dab8b810ce6903b8ea8a41d34b7c58cfad41c6dbb8"></a>

## aws_parameters.admin_password — aws_parameters.admin_password / b6d46b9ff6ae / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.admin_password

<a id="canonical-7cb4c6aa01a1eef88e82635af4d54c59805530cb271e63926bb45a0990951904"></a>

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

<a id="canonical-91ccd89d687a0fe637e90fd69e72c84a11226b3737372f8b2c5ad0b2cc25f77b"></a>

## Direct properties — aws_parameters.admin_password / b6d46b9ff6ae / 3

- [blindfold_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-446fa155df53ed513c7c923a31c70e67936c6545f646ad2b8d53b0d07091255f): complete subsection reference.

- [clear_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-dd4aee074fce03b02f7a708ba0d1b8b72b8876e271d3db9e1d0bb0ff7c903d59): complete subsection reference.

<a id="canonical-afd90d8614e664c2691dc2fab278e01ac144cde70637eab539e3bff2efee5777"></a>

## Next pages — aws_parameters.admin_password / b6d46b9ff6ae / 4

- [aws_parameters.admin_password.blindfold_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-446fa155df53ed513c7c923a31c70e67936c6545f646ad2b8d53b0d07091255f)
- [aws_parameters.admin_password.clear_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-dd4aee074fce03b02f7a708ba0d1b8b72b8876e271d3db9e1d0bb0ff7c903d59)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-446fa155df53ed513c7c923a31c70e67936c6545f646ad2b8d53b0d07091255f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3164e264b58dfe630f4a2b13bb5b20ca3cd1df8b3b04b9fae05559e8cb56efb"></a>

## aws_parameters.admin_password.blindfold_secret_info — aws_parameters.admin_password.blindfold_secret_info / e82dc96e43e6 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-427fd9512fafcbf2e602a41a442924d9b16b2a85b3965f597e03498ba4f40d81)
- aws_parameters.admin_password.blindfold_secret_info

<a id="canonical-8ea0b7b3e9f28e6192573cdb90cc0425ccd133d9c5fe462a49181302ec667344"></a>

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

<a id="canonical-1d337bca949fe88560d75b7a5c028607c8feca5963bcce788d4554d91e835f3f"></a>

## Direct properties — aws_parameters.admin_password.blindfold_secret_info / e82dc96e43e6 / 3

<a id="canonical-5bf3dde0b7ad07db23f4fe349db8909a3cfd65ec6670b2878f5d5d463ac47f46"></a>

<a id="canonical-c92eb258bc235ea1fd8b5538351dfbf2db9ce119679b89609f985d797fc6c200"></a>

## decryption_provider property — aws_parameters.admin_password.blindfold_secret_info / e82dc96e43e6 / 4

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

<a id="canonical-77c3ca25058e70b854041518ac49ab4f5fc701c7dca7372d16dcb584e3a0305b"></a>

<a id="canonical-18ca09dc4194a88113b2c4e3c30f95e0e9671c9ef2f98ef34d9b19fecd4603b3"></a>

## location property — aws_parameters.admin_password.blindfold_secret_info / e82dc96e43e6 / 5

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

<a id="canonical-cc232a1c77647b6335e39380e9fc1fe20c4ee377a4fcaf2f9ce709fb3e5a4d53"></a>

<a id="canonical-91da61dfdbde42e6aa996f598a759e966fb6a5600e5e8123eab32da7a3557efd"></a>

## store_provider property — aws_parameters.admin_password.blindfold_secret_info / e82dc96e43e6 / 6

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

<a id="canonical-5479f7ed1525e465eb251a628599000b39f1d00af72da1187e97de24761fcb09"></a>

## Next pages — aws_parameters.admin_password.blindfold_secret_info / e82dc96e43e6 / 7

- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-427fd9512fafcbf2e602a41a442924d9b16b2a85b3965f597e03498ba4f40d81)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-dd4aee074fce03b02f7a708ba0d1b8b72b8876e271d3db9e1d0bb0ff7c903d59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ceb085005d1d012fc41684e3e037b433a8445c0ad0906e5a12ecf0c33272165c"></a>

## aws_parameters.admin_password.clear_secret_info — aws_parameters.admin_password.clear_secret_info / 659cb56e582b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-427fd9512fafcbf2e602a41a442924d9b16b2a85b3965f597e03498ba4f40d81)
- aws_parameters.admin_password.clear_secret_info

<a id="canonical-12f7a2144dd802528fa48305f82261aeec79a9680206d01754585d730d67a58d"></a>

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

<a id="canonical-159bcc0e7bcc31563c74fbb64b0b2bca81fb7e711e05f123313d6631cb0fcbb0"></a>

## Direct properties — aws_parameters.admin_password.clear_secret_info / 659cb56e582b / 3

<a id="canonical-590996c1f9e78b5429c242299a78aabd53ea0d91540aa717f2b9f976cefaffda"></a>

<a id="canonical-5bb9e82c3f7646d3057d0ba6c578ae2d19768bc6b3c2305c6f994f2b84c897ef"></a>

## provider_ref property — aws_parameters.admin_password.clear_secret_info / 659cb56e582b / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f628fdf83a8229e4d41c41e358ceabc52dff7a1518f2c9a3ab997c08aa3cc0c4"></a>

<a id="canonical-a1dc216e35b2ba289e2e1d25826574f6221e0cada6da9239a2df423acbac82aa"></a>

## url property — aws_parameters.admin_password.clear_secret_info / 659cb56e582b / 5

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

<a id="canonical-3066897f1032d97b67703718195a8943fb69f5b5c23adda22da5cc0fe0bc40be"></a>

## Next pages — aws_parameters.admin_password.clear_secret_info / 659cb56e582b / 6

- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-427fd9512fafcbf2e602a41a442924d9b16b2a85b3965f597e03498ba4f40d81)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-a17fd7dc76c41860eb770c2e67e95585b5e165f22361b47f7d15c8da27b5798d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1c8bef3ecf419c5fabd9ada6e3587b261054eced3d5e3cf16a1741658406c29"></a>

## aws_parameters.aws_cred — aws_parameters.aws_cred / 8490e8938d2b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.aws_cred

<a id="canonical-e87d2bbe5bfacc142975eb93a83741e9fd979cff5d5d039055649cbc83e50c78"></a>

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

<a id="canonical-cd22c20d27ea582d45119c6a3a27bbee116bd8910254cf8e660c9d5d027ed5f4"></a>

## Direct properties — aws_parameters.aws_cred / 8490e8938d2b / 3

<a id="canonical-14ed388e5566984123500eb9315250425e37f4f05eb893714bee224a137b0c7d"></a>

<a id="canonical-d8c4974100edfb3d6a243b4650d4fc885f21231e46bcc0c58eebd5ed8e6a83ba"></a>

## name property — aws_parameters.aws_cred / 8490e8938d2b / 4

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

<a id="canonical-a0fc8964422db1f4ffeb618164193045f6556725a9b11b827f53733bb4b76b3a"></a>

<a id="canonical-fb420a620251c0059d4e559ec2026f0d902393f792dcd24c7dac059382431063"></a>

## namespace property — aws_parameters.aws_cred / 8490e8938d2b / 5

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

<a id="canonical-f27b231f78ea415949cd405553b461d56c4df610385c5b247860121d0628ff28"></a>

<a id="canonical-1f72a136daeb719e8dfa6a731ab13d383d1be43e4a650a866a13425b2c8c77bb"></a>

## tenant property — aws_parameters.aws_cred / 8490e8938d2b / 6

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

<a id="canonical-f8975bdd1110ca9dfc17b834b364836fba01d9c560dc465ce90314e0baac827e"></a>

## Next pages — aws_parameters.aws_cred / 8490e8938d2b / 7

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42ee01ec7ed19eccd8c4064df1630e3a4f11e7d0be5987b51f81768889dc4644"></a>

## aws_parameters.az_nodes — aws_parameters.az_nodes / f6279c113c27 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.az_nodes

<a id="canonical-5157eae2c0af635e1567709493298dffeef608a7ab4fc8c7ce77fe4d94d727d1"></a>

Type: `"list"`. Computed.

Only Single AZ or Three AZ(s) nodes are supported currently.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

<a id="canonical-ecc30fc7eb427289b3064261afce54e8946dd3987e969b53250cc895598d27f8"></a>

## Direct properties — aws_parameters.az_nodes / f6279c113c27 / 3

<a id="canonical-a093ac2c1981421a35bde14f3c3f2297adca08fbc0d33ea9d9e3ce82fb011e41"></a>

<a id="canonical-3d8519d32b1682a7d1783e3168cbb71a7ef14ea66ade8852f00299ef7c4e92eb"></a>

## aws_az_name property — aws_parameters.az_nodes / f6279c113c27 / 4

Type: `"string"`. Computed.

AWS availability zone, must be consistent with the selected AWS region.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-4e95a8a46200e19b81e3af77ec55fcea9753f5261737df6557bd989f8025e91d): complete subsection reference.

- [outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-b0ed83977cd935f41cf061ac32e42a8da212d31fe452c07184063a341ba67f1a): complete subsection reference.

- [reserved_inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-5b3dea02ce99ee0d99b1d46392a2529328d85b15d542e6eadcd62d2c09d6ec18): complete subsection reference.

- [workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-812ff56d70ff67609b483f3079ab741330778b9f3fdef94b3be21ba7e74e40a3): complete subsection reference.

<a id="canonical-faeb41324089f178c77f3856a49f22a6b7181aa91ae1595b25eb7ce93be839e2"></a>

## Next pages — aws_parameters.az_nodes / f6279c113c27 / 5

- [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-4e95a8a46200e19b81e3af77ec55fcea9753f5261737df6557bd989f8025e91d)
- [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-b0ed83977cd935f41cf061ac32e42a8da212d31fe452c07184063a341ba67f1a)
- [aws_parameters.az_nodes.reserved_inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-5b3dea02ce99ee0d99b1d46392a2529328d85b15d542e6eadcd62d2c09d6ec18)
- [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-812ff56d70ff67609b483f3079ab741330778b9f3fdef94b3be21ba7e74e40a3)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-4e95a8a46200e19b81e3af77ec55fcea9753f5261737df6557bd989f8025e91d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92f9905a31846a6eb0549a4bf14c1f62e672c99d96ee7daa41ac6075db678916"></a>

## aws_parameters.az_nodes.inside_subnet — aws_parameters.az_nodes.inside_subnet / 67066111e560 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- aws_parameters.az_nodes.inside_subnet

<a id="canonical-367640f658465a59d0a6cf90493040a902a2b57727858269535452641f8b28c3"></a>

Type: `"single"`. Computed.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-5578c75a4b0d8558709cdfb57aae1d11ddfe3e00197fcdff63b486afa6090f47"></a>

## Direct properties — aws_parameters.az_nodes.inside_subnet / 67066111e560 / 3

<a id="canonical-5dfef55b61fd03397fd08485f3fcca6b052a243d309e142d9b76c0af8990c938"></a>

<a id="canonical-a454afef48e3140f29040f14f2a0c214e92450029a81d25073c08cf3c2b9f11d"></a>

## existing_subnet_id property — aws_parameters.az_nodes.inside_subnet / 67066111e560 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-f02762e48c495008c570cb705ca2f2ab53ce550298eacf1a1d31e5aed4440b36): complete subsection reference.

<a id="canonical-f8f8960cb29d549c27d976325f0d9cd3ab8ce7f4dca08d7ad83c9a16c609e067"></a>

## Next pages — aws_parameters.az_nodes.inside_subnet / 67066111e560 / 5

- [aws_parameters.az_nodes.inside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-f02762e48c495008c570cb705ca2f2ab53ce550298eacf1a1d31e5aed4440b36)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-f02762e48c495008c570cb705ca2f2ab53ce550298eacf1a1d31e5aed4440b36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3eef87c8bba61d5085ef6f3d565f3808533af380b16e7f4a6d1455979d39d232"></a>

## aws_parameters.az_nodes.inside_subnet.subnet_param — aws_parameters.az_nodes.inside_subnet.subnet_param / 416a8e804579 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-4e95a8a46200e19b81e3af77ec55fcea9753f5261737df6557bd989f8025e91d)
- aws_parameters.az_nodes.inside_subnet.subnet_param

<a id="canonical-fa1c1585c9c18e4d1f1b1e5f36c251e933a8a162bd6783d88518bd503397c424"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-8762e9c32754424a1c72a854e99d0a5710901452cc7997bd7c746b958dc32e57"></a>

## Direct properties — aws_parameters.az_nodes.inside_subnet.subnet_param / 416a8e804579 / 3

<a id="canonical-bfb5e73bb366f45f82cb20eec32db08c5604326886d353f4a1e397b60b6d46c6"></a>

<a id="canonical-40aaca85d7b314cdaa8bec92b8d05ae0edbda28a365048125788640f4b05f1da"></a>

## ipv4 property — aws_parameters.az_nodes.inside_subnet.subnet_param / 416a8e804579 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-d98d0f92607f8046c86968c3cb4d92961083dd39edc43782f9dab0b675332386"></a>

## Next pages — aws_parameters.az_nodes.inside_subnet.subnet_param / 416a8e804579 / 5

- [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-4e95a8a46200e19b81e3af77ec55fcea9753f5261737df6557bd989f8025e91d)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-b0ed83977cd935f41cf061ac32e42a8da212d31fe452c07184063a341ba67f1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ae61e23c0b2ddc9fb4636c35c6499670af7d96efeb046b0f020bb455c6ec8a7"></a>

## aws_parameters.az_nodes.outside_subnet — aws_parameters.az_nodes.outside_subnet / ec4c2859bc01 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- aws_parameters.az_nodes.outside_subnet

<a id="canonical-cac0c3dc598ea0f0ac491ebbd90a76874b057534f5e8d9e201b06fa542587c0f"></a>

Type: `"single"`. Computed.

Configuration parameter for outside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-4092f03f9f8276320a24b56a8f7686711c338d80fc32fdde727af48c41f2e1a7"></a>

## Direct properties — aws_parameters.az_nodes.outside_subnet / ec4c2859bc01 / 3

<a id="canonical-e5ae1d93f74386581aefd18c0a724644df651426c2c8589e90099aad87cfd10c"></a>

<a id="canonical-20693a722a6e15b773ebec711ce7a1a1a0c3d4f1a5553202f15604314de381bd"></a>

## existing_subnet_id property — aws_parameters.az_nodes.outside_subnet / ec4c2859bc01 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-6abd1265903a78fff682e102d59e86cbc22257cad9892c9c08c4ebc5737d9db7): complete subsection reference.

<a id="canonical-c4f7976565091e6b1840aade8a0a770bf5891cf24f68093177cd20865f3057f2"></a>

## Next pages — aws_parameters.az_nodes.outside_subnet / ec4c2859bc01 / 5

- [aws_parameters.az_nodes.outside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-6abd1265903a78fff682e102d59e86cbc22257cad9892c9c08c4ebc5737d9db7)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-6abd1265903a78fff682e102d59e86cbc22257cad9892c9c08c4ebc5737d9db7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43090945ccefb6dd493af54332388450945bbb63597136b2bcbbdc11a5a836a2"></a>

## aws_parameters.az_nodes.outside_subnet.subnet_param — aws_parameters.az_nodes.outside_subnet.subnet_param / 6649662a9bb3 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-b0ed83977cd935f41cf061ac32e42a8da212d31fe452c07184063a341ba67f1a)
- aws_parameters.az_nodes.outside_subnet.subnet_param

<a id="canonical-5ce9dd41425ebdef288463b9f349e5e20fcf6a39525215614c67cfcc854ad05c"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-6d41e4430dee61e61e392bcca8a02fec988293399676aa6c7c4d55edc7c15c58"></a>

## Direct properties — aws_parameters.az_nodes.outside_subnet.subnet_param / 6649662a9bb3 / 3

<a id="canonical-426685a193734c756187149b8741e0ff6026e292eabb699db8b6ba6b18e66797"></a>

<a id="canonical-9ef04ecd5570117a702b45ce45ec3737772c0d09c0ae8e2e7e959c56fe7a51b9"></a>

## ipv4 property — aws_parameters.az_nodes.outside_subnet.subnet_param / 6649662a9bb3 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-4ebea843f2ced717115ed960a48ccf8b9c07420098bf2659eaaeea8e80a7f31a"></a>

## Next pages — aws_parameters.az_nodes.outside_subnet.subnet_param / 6649662a9bb3 / 5

- [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-b0ed83977cd935f41cf061ac32e42a8da212d31fe452c07184063a341ba67f1a)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-5b3dea02ce99ee0d99b1d46392a2529328d85b15d542e6eadcd62d2c09d6ec18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76fa4aa77cda31f714b436e314a17f2c04e4fb64d683468d05b77d0c9bfbb595"></a>

## aws_parameters.az_nodes.reserved_inside_subnet — aws_parameters.az_nodes.reserved_inside_subnet / 975ee5654d47 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- aws_parameters.az_nodes.reserved_inside_subnet

<a id="canonical-a2e534f4996fc48abf0b04e943da5b4fb8fcaac2d4915f88b58c325b88bf7a8f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved inside subnet.

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

<a id="canonical-21eac3219a099e29826839b0a30d3c30b7a6f7469d73c1a2224e4d0050e69985"></a>

## Direct properties — aws_parameters.az_nodes.reserved_inside_subnet / 975ee5654d47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-575d7911c4d26c232206ab77c279f482e1b8c5fc7fe403b2ba86b27cc9b93133"></a>

## Next pages — aws_parameters.az_nodes.reserved_inside_subnet / 975ee5654d47 / 4

- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-812ff56d70ff67609b483f3079ab741330778b9f3fdef94b3be21ba7e74e40a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75978b2451e5b0da5c6c58f8c4ca963259a10b14120dc87b211a5665ae1e52e4"></a>

## aws_parameters.az_nodes.workload_subnet — aws_parameters.az_nodes.workload_subnet / 4c33bc0a0f19 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- aws_parameters.az_nodes.workload_subnet

<a id="canonical-953bae85b93c126fa2024ba7af826cb0ec7aeabb2e94d75c8bd5a09866f82d02"></a>

Type: `"single"`. Computed.

Configuration parameter for workload subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-de0a0e0ce4a72b77e682314513d6cc619793a2e8b7ae6a293966a089112856f2"></a>

## Direct properties — aws_parameters.az_nodes.workload_subnet / 4c33bc0a0f19 / 3

<a id="canonical-bb61dbd927df321946313d17772d6036c665375caed8ea3a8d4c4e594525e044"></a>

<a id="canonical-9cd274ffff8e4cf7e85fef4f93405d3b3a820a3fc1033e70c8414cbc67ea3786"></a>

## existing_subnet_id property — aws_parameters.az_nodes.workload_subnet / 4c33bc0a0f19 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-b0d2cc2d0c90194cb082b92ba5431e970c452ea774d6e75507b086eb1a268ddc): complete subsection reference.

<a id="canonical-2168ab2d88a96afcaa643d3697140d604266b379ad0e066e727876a16e3b769b"></a>

## Next pages — aws_parameters.az_nodes.workload_subnet / 4c33bc0a0f19 / 5

- [aws_parameters.az_nodes.workload_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-b0d2cc2d0c90194cb082b92ba5431e970c452ea774d6e75507b086eb1a268ddc)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-b0d2cc2d0c90194cb082b92ba5431e970c452ea774d6e75507b086eb1a268ddc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48a9c842a04fe384992d3bd366e2234a31c2a01baf9d7e78b2b4b2f9f70ac320"></a>

## aws_parameters.az_nodes.workload_subnet.subnet_param — aws_parameters.az_nodes.workload_subnet.subnet_param / a317563660ea / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-30aff105170222d3890f8e954b7715e0529c7f856735cbf3b5e798dc919fdfa2)
- [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-812ff56d70ff67609b483f3079ab741330778b9f3fdef94b3be21ba7e74e40a3)
- aws_parameters.az_nodes.workload_subnet.subnet_param

<a id="canonical-1ef20757846b40763cc41ebe6e255a0af06591db1daefdf026a241095f3273df"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-9eb9a6f042b43aa4d866e061b244673ac7f995f47f1c693f6006ff6eabbdd340"></a>

## Direct properties — aws_parameters.az_nodes.workload_subnet.subnet_param / a317563660ea / 3

<a id="canonical-c6544431c63a5aa3b0aea28ce44fbee7bfb41fb708dc0ae29805243e0dadef89"></a>

<a id="canonical-ded355f4fd34004c2f487548d1eec634c8b05d531e06475ce7d946c017a74d89"></a>

## ipv4 property — aws_parameters.az_nodes.workload_subnet.subnet_param / a317563660ea / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-ec688d4288557cbe983f1c867d44361026b37e905526997831174b8123b12406"></a>

## Next pages — aws_parameters.az_nodes.workload_subnet.subnet_param / a317563660ea / 5

- [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-812ff56d70ff67609b483f3079ab741330778b9f3fdef94b3be21ba7e74e40a3)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-60f382ec3a23f7ed7f880aa170917c5dcaa87378bb0d17d629f9258968a389ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41682f9ba24decfe13d597a7c6efa48417805ddda53e711866b640fe9fdd08b5"></a>

## aws_parameters.custom_security_group — aws_parameters.custom_security_group / c2064a46e26d / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.custom_security_group

<a id="canonical-254cde3f8697513bf6705d9c062ac1962c6b8cdfafd6f8cb9d02fe4549760877"></a>

Type: `"single"`. Computed.

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

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

<a id="canonical-41443cf76833b786c3b105142e0b760251c80abc1cd80c894c635888135acca5"></a>

## Direct properties — aws_parameters.custom_security_group / c2064a46e26d / 3

<a id="canonical-508be9d27736d9dc5d61ccd51643159181edbbb0f708b16a8fc7dd063d2226a6"></a>

<a id="canonical-1209d9f45ff5c97f173416893ec504096c72fcc1b1bb0c8f0c7ebc63c94bd43f"></a>

## inside_security_group_id property — aws_parameters.custom_security_group / c2064a46e26d / 4

Type: `"string"`. Computed.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-83623756b5c51f7dcffbeb84055e61cc581acfb73a55f3497392becc791d6eba"></a>

<a id="canonical-abdbe07f37db32dfef4a410db0e7dba9db9ffd951aa07aab1dd1dacbc595a781"></a>

## outside_security_group_id property — aws_parameters.custom_security_group / c2064a46e26d / 5

Type: `"string"`. Computed.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-950141b5b76ae0c4c5c3384346fedf4c40faad050f80990256f8098fd0286144"></a>

## Next pages — aws_parameters.custom_security_group / c2064a46e26d / 6

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-beb858327d992a81f23bfefeff3d4f7b5b31abdd4b38af3f3c411fe48f78d9ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71c3d61e478a36ccf1d7b78c4b34cc66cb50d04861f23ad1f7e4899a4e07c8d3"></a>

## aws_parameters.disable_encryption — aws_parameters.disable_encryption / 32afe166a9ff / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.disable_encryption

<a id="canonical-5084d065d2220181f508a8b978fe9a624818957cb63040edae30057de73412fa"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable encryption.

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

<a id="canonical-47eb5ff00262e551af019fc01185f0bed76a65aeb480fc03af7e822cd428db2a"></a>

## Direct properties — aws_parameters.disable_encryption / 32afe166a9ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f9ae020abd64cc1fe0e90e4198b1dc3cc906d4f626d9a356a737665675b883a"></a>

## Next pages — aws_parameters.disable_encryption / 32afe166a9ff / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-4ba2140ab942a6f4f0bd961b3a1e03755aaf8ef7dc6c8486634f349722c518ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-127eb9fcb4ebbc87dccb7b7833ed13499223954aed4da7eaa5745cd7c19a8d6b"></a>

## aws_parameters.disable_internet_vip — aws_parameters.disable_internet_vip / 14e9fbe3f9dc / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.disable_internet_vip

<a id="canonical-b19ce2d7ebdc7bac9a0205e43f4f5cac1c41ec3772106f59bd31c78cc5d3a238"></a>

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

<a id="canonical-399654ff8f535c1fce2f95eec36f00a75216fec85adb781c71ca5bb0c16659bb"></a>

## Direct properties — aws_parameters.disable_internet_vip / 14e9fbe3f9dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68b44e81632fe0e8dabbe37f8f3e9256d33691d5c8cb7d329dabb48da5703568"></a>

## Next pages — aws_parameters.disable_internet_vip / 14e9fbe3f9dc / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-c4c8261093089e3cd43d2d8384e21ffbec226771bcb7fc88616dc90630abc2a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f710ab7afbaa8a8b21f705e9c23a3a4f16f8ec3fc91a4f2082df028e1afe6ec"></a>

## aws_parameters.enable_encryption — aws_parameters.enable_encryption / a08a7465fc33 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.enable_encryption

<a id="canonical-4b01140c1e2eb68110f85355f798063e02818959ee3a03bc06e867d0f590aad7"></a>

Type: `"single"`. Computed.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

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

<a id="canonical-b5ce42c3ed6ac03de258f21d6d9704cd018bffea3d8169799820f8583cd16220"></a>

## Direct properties — aws_parameters.enable_encryption / a08a7465fc33 / 3

<a id="canonical-6b9486d97e651cb65cce97f9cb5c903e8ab362c962b66b073913c90f74bcd386"></a>

<a id="canonical-405cf5efaca36fa51eff06cbf2ffaf9951b90c265aea6908818790321204c6f7"></a>

## kms_key_id property — aws_parameters.enable_encryption / a08a7465fc33 / 4

Type: `"string"`. Computed.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-a07dd16f1a1004a59aa1f85444271a4af2b6ec66c5a4948d50a323dd06920cb0"></a>

## Next pages — aws_parameters.enable_encryption / a08a7465fc33 / 5

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-1a83b1e4b1d83b65d1d7185a98b222103b9b83073ec919f0add90ea2a77947f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33d795ff1853ee49e5e59812f2b5f85d5737d932509d0dfa9c6d28c208abb855"></a>

## aws_parameters.enable_internet_vip — aws_parameters.enable_internet_vip / 02bb0fcebb00 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.enable_internet_vip

<a id="canonical-36841cb8f2076efe557bc13aaaa8788453d36bf8b6e9dd52beb25acdf44d81dd"></a>

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

<a id="canonical-664934af7d45c385d782f2ac06e44b8313e6ee2ba30ddd9464f3e216f294a743"></a>

## Direct properties — aws_parameters.enable_internet_vip / 02bb0fcebb00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9627b13b1562f4f0a237f01290f53afdf566b6ee80171a97d605e40badf89f68"></a>

## Next pages — aws_parameters.enable_internet_vip / 02bb0fcebb00 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-8a2bcf5e62c4068a1db1e653b03e1cf153d8621975eb1c64399538af270f6bcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc89c064364049629d621ad7862e02a4267fe7c8875f1319b3ab95f83b46bbc8"></a>

## aws_parameters.existing_tgw — aws_parameters.existing_tgw / e396aab96816 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.existing_tgw

<a id="canonical-f194ae324b82d34e0086dafe19c745c2d80556de95fc17e4770ecf2933b8e680"></a>

Type: `"single"`. Computed.

Configuration parameter for existing tgw.

Upstream description:

Information needed for existing TGW.

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

<a id="canonical-079b03bd608783f4d949deac9db550bc8eb06f294926d9b494f6554d17de4649"></a>

## Direct properties — aws_parameters.existing_tgw / e396aab96816 / 3

<a id="canonical-8a571f0fb53ba8fca48c7b80f8fffd34109e808d86475150faff0ffad9a89cff"></a>

<a id="canonical-b17a552f515065bf32135daf3abb71923a6a04db5652760971538a71443d59f6"></a>

## tgw_asn property — aws_parameters.existing_tgw / e396aab96816 / 4

Type: `"number"`. Computed.

Enter TGW ASN. TGW ASN.

Upstream description:

TGW ASN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2e3685af6971d03b71b3c9b9bc774357b39e3cac1dd111b99fed4fb512e1e06c"></a>

<a id="canonical-3b0fc679f28f5abc141c7cb974f60c11031624fc59d16696eb0237ea584d8caf"></a>

## tgw_id property — aws_parameters.existing_tgw / e396aab96816 / 5

Type: `"string"`. Computed.

Existing TGW ID. Existing TGW ID.

Upstream description:

Existing TGW ID.

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
    "pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-cefc7d17d1229641e8241879aa780ae337dbae3d03541889e7d9a8ac48f4295a"></a>

<a id="canonical-a1a335a66b78d3bb0e00ebaad8136fd2558b776d1d4dfc0083da58c5458a5848"></a>

## volterra_site_asn property — aws_parameters.existing_tgw / e396aab96816 / 6

Type: `"number"`. Computed.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-b4c03e721b783f6f99991c949bcfa43c297f49465fd219079b694373d523fdf4"></a>

## Next pages — aws_parameters.existing_tgw / e396aab96816 / 7

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-958d863e35e9da10743f095c325bcf760046b81fe41e2e8211017e5ff8270d08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf9ed7f8eb6da73bb6e832121e3cf10143ee88ed3c61a2d33761a0a3a0e82a8a"></a>

## aws_parameters.f5xc_security_group — aws_parameters.f5xc_security_group / f5dea9c91149 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.f5xc_security_group

<a id="canonical-3f9ba2046da1f8ded40a732e09c45e4df1a7fb8685a1b2b580328e987ce12198"></a>

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

<a id="canonical-c24f85b90cb22b862d59bdc1b4b227289cca82e6c6e5e64e55ebe2ff464109f0"></a>

## Direct properties — aws_parameters.f5xc_security_group / f5dea9c91149 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ecb27185a24c75ee38021d4c2b9cdf6f9c65d904deff5a3fd86fd93c36acf05e"></a>

## Next pages — aws_parameters.f5xc_security_group / f5dea9c91149 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-6ecdcdc30012a777ac3b186af4be6961737177618a676c001871e0b5fc334c53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e7bc5486b5df1ad2cf51e3e2252bfc04d7c71c7ea9b3b8c5ff73beaf6854030"></a>

## aws_parameters.new_tgw — aws_parameters.new_tgw / 2bb86b477d3c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.new_tgw

<a id="canonical-cc0802a381820efad57a1e1ef762d4ca5bda3385cb3f9542f02c82f79a9f79e0"></a>

Type: `"single"`. Computed.

TGWParamsType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"system_generated\",\"user_assigned\"]"
}
```

<a id="canonical-2fc9ac3d1ddd1c405dc34cb1c982a5bef0987472bde6b9f3f5ee0b61c8c80bbe"></a>

## Direct properties — aws_parameters.new_tgw / 2bb86b477d3c / 3

- [system_generated](data-sources--aws_tgw_site--reference--group-001.md#canonical-2b5f9950fd3dd4c08a26bea2beab1a381d246d732b175b719af75956798f2402): complete subsection reference.

- [user_assigned](data-sources--aws_tgw_site--reference--group-001.md#canonical-cf002520cf2afccaf09871559284c0bd352c742129384851652b0bfdfd649955): complete subsection reference.

<a id="canonical-beadee0d72bb01dcffb7fcb732f09097626248998a12df1429053b73b38927a6"></a>

## Next pages — aws_parameters.new_tgw / 2bb86b477d3c / 4

- [aws_parameters.new_tgw.system_generated](data-sources--aws_tgw_site--reference--group-001.md#canonical-2b5f9950fd3dd4c08a26bea2beab1a381d246d732b175b719af75956798f2402)
- [aws_parameters.new_tgw.user_assigned](data-sources--aws_tgw_site--reference--group-001.md#canonical-cf002520cf2afccaf09871559284c0bd352c742129384851652b0bfdfd649955)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-2b5f9950fd3dd4c08a26bea2beab1a381d246d732b175b719af75956798f2402"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a85c161daea80bd92777efe63d36a3bb1dc02545f06a87c005400f5af8d0483b"></a>

## aws_parameters.new_tgw.system_generated — aws_parameters.new_tgw.system_generated / 863e7230e1f4 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-6ecdcdc30012a777ac3b186af4be6961737177618a676c001871e0b5fc334c53)
- aws_parameters.new_tgw.system_generated

<a id="canonical-0087dde7c8d7547be4b723dc6e593188815cd466a32c639e2a26c97c829872e3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for system generated.

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

<a id="canonical-ad90469ccd88813fbbcc75ba8a8fa168125ac4e951271f517d81735f77ee9385"></a>

## Direct properties — aws_parameters.new_tgw.system_generated / 863e7230e1f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c2cf4a079c56a504f8de038c53b95298f3b2599186b8eb3295e798cb72ff422"></a>

## Next pages — aws_parameters.new_tgw.system_generated / 863e7230e1f4 / 4

- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-6ecdcdc30012a777ac3b186af4be6961737177618a676c001871e0b5fc334c53)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-cf002520cf2afccaf09871559284c0bd352c742129384851652b0bfdfd649955"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4ce59157a108370d220f3e49a93a1c5de90b874985d2c459a06818dc7d77cf6"></a>

## aws_parameters.new_tgw.user_assigned — aws_parameters.new_tgw.user_assigned / a72c627cc4ed / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-6ecdcdc30012a777ac3b186af4be6961737177618a676c001871e0b5fc334c53)
- aws_parameters.new_tgw.user_assigned

<a id="canonical-520a902fa30b9ec99f7046ac8b778c19b62cf5366c36447d8515cad5775c230f"></a>

Type: `"single"`. Computed.

Information needed when ASNs are assigned by the user.

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

<a id="canonical-d2f3c15183b7ccd56b1a8c3b8fdcd3adc2b120de3e3b930f15ad45e7265a34db"></a>

## Direct properties — aws_parameters.new_tgw.user_assigned / a72c627cc4ed / 3

<a id="canonical-56a447e649734880658ec3b551dfced844778533d91d45f3c5f7920b00a3e607"></a>

<a id="canonical-5c45aaf9e457bb6db6d6d6669c790d236ca22b96a829772779d6a5a484961d9b"></a>

## tgw_asn property — aws_parameters.new_tgw.user_assigned / a72c627cc4ed / 4

Type: `"number"`. Computed.

TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65534,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 64513
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  }
}
```

<a id="canonical-795b3ec8bc1b1c9266a6bd022986d8cd1c2eb1e26cc5ad4e3de5d226f2d8662c"></a>

<a id="canonical-5f019260453334499991425f265b9fcb10a1b83f2a9a283e1535e724c3eee0a0"></a>

## volterra_site_asn property — aws_parameters.new_tgw.user_assigned / a72c627cc4ed / 5

Type: `"number"`. Computed.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-34422d9c6faf1c369f103ff75a08ac69f95372666d88eeba263ac3ce647d419f"></a>

## Next pages — aws_parameters.new_tgw.user_assigned / a72c627cc4ed / 6

- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-6ecdcdc30012a777ac3b186af4be6961737177618a676c001871e0b5fc334c53)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-a3d4cc21785dd411c6e93fdfc38f79feb4c8bdc2c791c0768ffa790c0cf3ac2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72321f67dbb3f56e62d4852406ca9960650745c3831313d6a01d6aae94acd17e"></a>

## aws_parameters.new_vpc — aws_parameters.new_vpc / 7f25e769572b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- aws_parameters.new_vpc

<a id="canonical-16b56bd50f11ce4baf6b405c4f8c106dfdbc3f6002b7eebc6efd0eda2c42ebd2"></a>

Type: `"single"`. Computed.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name_tag\"]"
}
```

<a id="canonical-519da093719a20b290e25af55b29aa6c9b0a7558136949e6dcaae3bcb19d5a10"></a>

## Direct properties — aws_parameters.new_vpc / 7f25e769572b / 3

- [autogenerate](data-sources--aws_tgw_site--reference--group-001.md#canonical-b134399d405f509f2ccc812729edfc81a54911a399cbd20c36f88a49d4627a5b): complete subsection reference.

<a id="canonical-85a979e6bcd3e87797b1f0fca7809fa9a3550963b234653e5c6c2e03597efddd"></a>

<a id="canonical-25316b0b9c189f1154771bf464637c60bef445253b1b53c755cadbcbbdcd7e3a"></a>

## name_tag property — aws_parameters.new_vpc / 7f25e769572b / 4

Type: `"string"`. Computed.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-019370d02e077cd8bfb540ab198e18da745e4a61b7b45f921b0d38d7838322a9"></a>

<a id="canonical-5fcfff4ad3426c287bf73fd38df2ea9e96dc0d2a13bdcfbaf541af9895af100b"></a>

## primary_ipv4 property — aws_parameters.new_vpc / 7f25e769572b / 5

Type: `"string"`. Computed.

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

Upstream description:

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  }
}
```

<a id="canonical-b3e341a7589d005756a724eb8dfc097195fe83318bcea57e6ce950169c0e291f"></a>

## Next pages — aws_parameters.new_vpc / 7f25e769572b / 6

- [aws_parameters.new_vpc.autogenerate](data-sources--aws_tgw_site--reference--group-001.md#canonical-b134399d405f509f2ccc812729edfc81a54911a399cbd20c36f88a49d4627a5b)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-b134399d405f509f2ccc812729edfc81a54911a399cbd20c36f88a49d4627a5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0a287bb3ebb4dfdee302e99ccf8e412f0c43cec0d94f7598bdc3465bedd70d7"></a>

## aws_parameters.new_vpc.autogenerate — aws_parameters.new_vpc.autogenerate / b2575d2451f3 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-f561136116f0ff553c19f73d9aad6adf704bb06caab43534d606cbee804f5872)
- [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-a3d4cc21785dd411c6e93fdfc38f79feb4c8bdc2c791c0768ffa790c0cf3ac2f)
- aws_parameters.new_vpc.autogenerate

<a id="canonical-be5e52acc93b86578b572c18119b32c142190908de72976b54061ba62abbb063"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for autogenerate.

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

<a id="canonical-e3adfa3ff7b45fd07e09c603fec1c97c9f89a9bf3480a34212dca243e30476aa"></a>

## Direct properties — aws_parameters.new_vpc.autogenerate / b2575d2451f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-28c91fb3a3412bb29b1dbf6031be4e0e33e27b379d6382ca7e3e3f9a79c781ef"></a>

## Next pages — aws_parameters.new_vpc.autogenerate / b2575d2451f3 / 4

- [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-a3d4cc21785dd411c6e93fdfc38f79feb4c8bdc2c791c0768ffa790c0cf3ac2f)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-2d42d69e895bc6d802c1728cb112a4af67cae50fca5e22033ac294bbaf84a020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
