---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b9c344dda81d85a5002fffe8c0c02da26bad6c52aaaf7c9ca796dee6c6b1f82"></a>

## Property reference — Property reference / ccc795537f3d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- Property reference

<a id="canonical-80361c41eaca747bd278309ec3d3f00972d6ef7e6115d17a93114dbb6bc4f56e"></a>

## Direct properties — Property reference / ccc795537f3d / 3

- [active_enhanced_firewall_policies](resources--securemesh_site_v2--reference--group-003.md#canonical-f4bda8c43ee02c38bd6ed77937dfb6615581f04e02cd0f40c56322415f3f5676): complete subsection reference.

- [active_forward_proxy_policies](resources--securemesh_site_v2--reference--group-003.md#canonical-c4e0a109d30e189596943591d61d4f867bfffa50a82dcd88bb3b58c72767a3ad): complete subsection reference.

- [admin_user_credentials](resources--securemesh_site_v2--reference--group-003.md#canonical-bb58f908971ab8e93abc428e9ae0844492c513c9c2cf611d8617a1ae2d4d7864): complete subsection reference.

<a id="canonical-0f1a1bb7cf2818685c1b2c063ee3804dc6878dd4d6a389b2294aa635aa1248de"></a>

<a id="canonical-c86d33511a24767d114523b7a2e96409513bec4b366d70e587a6e6f8d9caf24b"></a>

## annotations property — Property reference / ccc795537f3d / 4

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

- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-941f3e76980a2bc1ca1b950343489241cf43152ae768f88b43bb60ce3a48fc6b): complete subsection reference.

- [azure](resources--securemesh_site_v2--reference--group-004.md#canonical-d6b794d7e7e25165fa55b5be6623d8b4c48cd27cb72b40d9005227c301983678): complete subsection reference.

- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-a15df73f71a4c0d9ee81c0cd83baade5f542ec9981ed827a060270cd4c587973): complete subsection reference.

- [block_all_services](resources--securemesh_site_v2--reference--group-006.md#canonical-27e1b585a2525e9107ce9a0d76a6b87f87194f36aab936dfbb0c0e3eaeedcdda): complete subsection reference.

- [blocked_services](resources--securemesh_site_v2--reference--group-006.md#canonical-4b2173bd5c5894768869b34a60f087003e187ccc133d462ce04b539f6685f51a): complete subsection reference.

- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-42db0c8c6f959a7a5f722072d54ec680250c2503463c40a3f95b6d6fcf9a37dd): complete subsection reference.

- [custom_proxy_bypass](resources--securemesh_site_v2--reference--group-006.md#canonical-6161a26cc75c3cc436aa4d184ebe978403144855e6afa6e46ff13a68d5a129ac): complete subsection reference.

- [dc_cluster_group_sli](resources--securemesh_site_v2--reference--group-006.md#canonical-48e5fd6f947f59d99bfd0186e0c606e7484dce6a49087fe54221aa92d3fd40bb): complete subsection reference.

- [dc_cluster_group_slo](resources--securemesh_site_v2--reference--group-006.md#canonical-d0ecf0dc72f632936f4646c4e8aba9453e52ed05430bdca584deae0f9e4c281d): complete subsection reference.

<a id="canonical-d0c9da96e01290985074bbfe058d19b8caa7e718d527ebf1f0c463fa4fe240f5"></a>

<a id="canonical-9a25c16b6e332d491d11e8d13d69d2740924c59ac7d1b9bb2fd0bcdf9007f074"></a>

## description property — Property reference / ccc795537f3d / 5

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

<a id="canonical-41dd1e2b4d0842bf543830140bea26e7361466fe56af8625b0e0722888a585ee"></a>

<a id="canonical-101e20fd87ef55e5dff84296954e8456306f999dbff9a0914511521a22c18402"></a>

## disable property — Property reference / ccc795537f3d / 6

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

- [disable_advanced_delivery](resources--securemesh_site_v2--reference--group-006.md#canonical-568c4da41e7354a20aaa23809debe13be4494ac0e006ff516bd71966e04155e4): complete subsection reference.

- [disable_ha](resources--securemesh_site_v2--reference--group-006.md#canonical-ed0b2efe8712bd7e09631916aef586943193383482a3fe628a4dd68bd10d0c42): complete subsection reference.

- [disable_log_anonymization](resources--securemesh_site_v2--reference--group-006.md#canonical-e35b27b3778b5ebe2a2e1be2e92733bd55598e0c50fc35e01fe3de556365ad01): complete subsection reference.

- [disable_management_network](resources--securemesh_site_v2--reference--group-007.md#canonical-dda29ba18b2927698f2c808d11dae7ffe652bce026e8bb699aa866e08235b266): complete subsection reference.

- [disable_url_categorization](resources--securemesh_site_v2--reference--group-007.md#canonical-320e434802af46ef2c7c79a166206d7f16c5c77b0e8c998d4f1711b2db5988ad): complete subsection reference.

- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2620c7c4eb2d479183bba23f1c2d51ee0bc0b06b1d53a6583cd3fd4730872200): complete subsection reference.

- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0): complete subsection reference.

- [enable_advanced_delivery](resources--securemesh_site_v2--reference--group-008.md#canonical-58d87ed581de59f3c7eb9469b207629142c00856112d3a9917b12e7d7ebd180b): complete subsection reference.

- [enable_ha](resources--securemesh_site_v2--reference--group-008.md#canonical-687f3b62d8ed4274294ed0a7726f16941786952f8fe35efd00640ef5dbe8a2de): complete subsection reference.

- [enable_log_anonymization](resources--securemesh_site_v2--reference--group-008.md#canonical-f7fd420a7bf79edf5db7e0403c835c6b9609c913a68a4ed301ebdcc856ff1750): complete subsection reference.

- [enable_management_network](resources--securemesh_site_v2--reference--group-008.md#canonical-d165033fedf36409580e6d7d7e956850c8ba9cc31e82f5d4457d87fd173f4398): complete subsection reference.

- [enable_url_categorization](resources--securemesh_site_v2--reference--group-008.md#canonical-942a730157531e299183f328ff838a8dc1db2259df9c81d356098acc27c24b9d): complete subsection reference.

- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca): complete subsection reference.

- [f5_proxy](resources--securemesh_site_v2--reference--group-009.md#canonical-c9c61f342c718474ad94bfa872822347d516d70f708bcd7ce7358dfea3955e25): complete subsection reference.

- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046): complete subsection reference.

<a id="canonical-98a400cdf0863e8905bd49ba46ed635a9add4eb6f2dba99363dc2ae0931fa226"></a>

<a id="canonical-357e95aee7ffc032df105db5e8050d4a2f342ea0671fa2d9c9b4d66b1610adb2"></a>

## id property — Property reference / ccc795537f3d / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db): complete subsection reference.

<a id="canonical-ea3b4b5838439cb3ce9a8a60704b00f9ba2cd53623b20fe2ddc0e58a84d3a6d3"></a>

<a id="canonical-499a03228c2e95d29e27c9ce9aa837537548fdcbb66ac8cf80450506ace66c0f"></a>

## labels property — Property reference / ccc795537f3d / 8

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

- [load_balancing](resources--securemesh_site_v2--reference--group-011.md#canonical-872e4a68860accaf3e925d83a4ce2b9691840fd9d1ff4d1b969cea185980d877): complete subsection reference.

- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589): complete subsection reference.

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c): complete subsection reference.

- [logs_streaming_disabled](resources--securemesh_site_v2--reference--group-012.md#canonical-a4bf722d3477ef61cdfac72a1f44fbf42bb3d727935964b96c0b646e651df9a9): complete subsection reference.

<a id="canonical-eaf5464ff5b5475859f823a6155f16b433a4033f9795dfe5f302bd536210b9f3"></a>

<a id="canonical-1a7ccb56b4c0356f0700899a26c25a0e59bd7963d608a68704048633872ceb94"></a>

## name property — Property reference / ccc795537f3d / 9

Type: `"string"`. Required.

Name of the Securemesh Site V2. Must be unique within the namespace. Must be at most 63 characters
(DNS-1035).

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
  stringvalidator.LengthAtMost(63),
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

<a id="canonical-5995cf4f224afad06064b3b031e0470ae0520f10d872c45918b2f5106aefafd5"></a>

<a id="canonical-aae1e6a4d7f77bc3f4258fe26c9ca02557da070dc66e66d18ea21625ca4cc742"></a>

## namespace property — Property reference / ccc795537f3d / 10

Type: `"string"`. Optional, Computed.

Namespace for the Securemesh Site V2. The F5 XC API restricts this resource to the system namespace;
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

- [no_forward_proxy](resources--securemesh_site_v2--reference--group-012.md#canonical-ed04394f5786d7ba076683bb8d8c3b295faccdf2dd9bbb6f7127a1214cd25f5d): complete subsection reference.

- [no_network_policy](resources--securemesh_site_v2--reference--group-012.md#canonical-55f68f318cc4eb83391191144f1bc143f9a78ff7aa479bf7e19f3e94c5313c05): complete subsection reference.

- [no_proxy_bypass](resources--securemesh_site_v2--reference--group-012.md#canonical-dbf7f694dafceef77a34b3fcca1e1e924ce856e24e5d48c4539c2cbbc5ae83fa): complete subsection reference.

- [no_s2s_connectivity_sli](resources--securemesh_site_v2--reference--group-012.md#canonical-fa1aa13eac43da651e8e0edc8322e51c0691628168db723973a8397b8861fb63): complete subsection reference.

- [no_s2s_connectivity_slo](resources--securemesh_site_v2--reference--group-012.md#canonical-4e5a9b729b39300b66e4c6167141f27fda0a34d68aa8603d3fba3eb5e17d8722): complete subsection reference.

- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39): complete subsection reference.

- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-7e95a8ee8e9c637536740a9d37bc35c444ce394db4642c61290233eee6269db5): complete subsection reference.

- [offline_survivability_mode](resources--securemesh_site_v2--reference--group-014.md#canonical-81d9cb11de186d5a2c5911e77fc85c81ebf75b02381e9ce3dad56e1778e2f501): complete subsection reference.

- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7): complete subsection reference.

- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a): complete subsection reference.

- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450): complete subsection reference.

- [private_adn](resources--securemesh_site_v2--reference--group-016.md#canonical-5efc800fd9dc85dc5717b76d73721de2fd8a153cee308446056e512d71e786fc): complete subsection reference.

- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-f20247f2bf62bff265884e5308bcb676fb69590f189a42d83f4a0123f809aa1e): complete subsection reference.

- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4): complete subsection reference.

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35): complete subsection reference.

- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f): complete subsection reference.

- [timeouts](resources--securemesh_site_v2--reference--group-017.md#canonical-9d7e81a82ada0db8be76c91fbb07dbb2361aa24927ad606494ddbb6ef6e4e50c): complete subsection reference.

<a id="canonical-df6a4e361ef4d1603603c1066f0900722987e4d62dcb522a1ac2786da45c71d8"></a>

<a id="canonical-669452442007eca132379bca70b32b69773e0b0b774cd0b02a160b83d222a574"></a>

## tunnel_dead_timeout property — Property reference / ccc795537f3d / 11

Type: `"number"`. Optional, Computed.

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

<a id="canonical-0986ee1b37201de5175d0c482d5e265468a31e95621e7a157a64290f61bc061f"></a>

<a id="canonical-5cc62f5162e5526f9b571356b888dc788d0ca403567809815d3875704636c1d0"></a>

## tunnel_type property — Property reference / ccc795537f3d / 12

Type: `"string"`. Optional, Computed.

\[Enum:
SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL|SITE\_TO\_SITE\_TUNNEL\_IPSEC|SITE\_TO\_SITE\_TUNNEL\_SSL\]
Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec
being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are
\`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`, \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\`,
\`SITE\_TO\_SITE\_TUNNEL\_SSL\`. Defaults to \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`.

Upstream description:

Tunnel encapsulation to be used between sites

Tunnel can operate in both IPsec and SSL, with IPsec being preferred over SSL. Tunnel is of type
IPsec Tunnel is of type SSL.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
    "SITE_TO_SITE_TUNNEL_IPSEC",
    "SITE_TO_SITE_TUNNEL_SSL"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
  "enum": [
    "SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
    "SITE_TO_SITE_TUNNEL_IPSEC",
    "SITE_TO_SITE_TUNNEL_SSL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b): complete subsection reference.

- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c): complete subsection reference.
