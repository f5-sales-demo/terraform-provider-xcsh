---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bd26cbe610f8f8c2020ad578e234240407d62cf400c463389be369631fc84b7"></a>

## Property reference — Property reference / 817c58179c3b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- Property reference

<a id="canonical-ac5b5415775f46fce7e008d6f7d85ed0311dd20bcc4e8f2b2f337d93b123b094"></a>

## Direct properties — Property reference / 817c58179c3b / 3

<a id="canonical-61bfd370e0cf264846f6167aa9e0def64a18311fd812098f810e921e751fe131"></a>

<a id="canonical-1a75bb172fb71ad791f7fde874a801256828b4c7f0b313ac0f9cecb675e8f57c"></a>

## annotations property — Property reference / 817c58179c3b / 4

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

<a id="canonical-f2790568a30eda70fb25c14dce030abd35be1258ee3c8291507eb087e7ca4053"></a>

<a id="canonical-263d6798ba2587868b9d3495c0c68e9f95eee1ee099cb3fb03852f04136d4ea1"></a>

## description property — Property reference / 817c58179c3b / 5

Type: `"string"`. Computed.

Description of the NfvService.

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

- [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-ef39bab5c106466d3615ee776eb28731edc5d8b52f8971794ea7c483b60be792): complete subsection reference.

- [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-693c0f81abe3465901899348b5507a2aeb9286cff3eb73bb93c18d56f2c365a6): complete subsection reference.

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b): complete subsection reference.

- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9): complete subsection reference.

- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56): complete subsection reference.

<a id="canonical-c47a43ca99f82e65054589c25a4c3e01109abeba23cbe47d9b946935fe2ab42c"></a>

<a id="canonical-fa9d058ff1b4182db1a1d477da3ca002a7b9439d25da77bc8d102693442e2fce"></a>

## id property — Property reference / 817c58179c3b / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-050bcef0aeea62ffa38b0e5cd5011893bbe3a35a0f68092ea34b80594fd9cd9a"></a>

<a id="canonical-447727d0ed92d0f839c290a1b9f6c3d362a8feac5eb12584aaf4cb897cf93010"></a>

## labels property — Property reference / 817c58179c3b / 7

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

<a id="canonical-3ba0cbcee4650406f9feb64b67d47bdb3e00e906b67fbfa686e4415a31de3615"></a>

<a id="canonical-562126a543fc0d8aa2dbf4b9c3a8e35617c4cfa58b9ee99076bc96e4c564a596"></a>

## name property — Property reference / 817c58179c3b / 8

Type: `"string"`. Required.

Name of the NfvService.

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

<a id="canonical-b46abce70cce9024a85edde02c4f68f14aa60cdb3da6891e4e51e0c17cb7cc90"></a>

<a id="canonical-cbe531acd87bd42b339f9b8c458be1aa75560ecb6ddf57b7352bb83c63338c36"></a>

## namespace property — Property reference / 817c58179c3b / 9

Type: `"string"`. Required.

Namespace where the NfvService exists.

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

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c): complete subsection reference.

<a id="canonical-6b168e2475d4f89958e08468cbea66dbcaf037a45f2c133337d7d8b86cd9c2a4"></a>

## All schema paths — Property reference / 817c58179c3b / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nfv_service--reference--group-001.md#canonical-61bfd370e0cf264846f6167aa9e0def64a18311fd812098f810e921e751fe131) |
| `description` | [description](data-sources--nfv_service--reference--group-001.md#canonical-f2790568a30eda70fb25c14dce030abd35be1258ee3c8291507eb087e7ca4053) |
| `disable_https_management` | [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-5bb1e34aac0ea8f4e80c234c1c7c9c4f5404c7618423475ae6fd72f747132c2c) |
| `disable_ssh_access` | [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-fac292344ae27a844333e24b4388e6d1bd197e81ca9ae49412159ae88f014b3b) |
| `enabled_ssh_access` | [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-edc881fb1eae592b96392b4c9b6ac03d4dfa8555aee864a12a9d7bf47ca86058) |
| `enabled_ssh_access.advertise_on_sli` | [enabled_ssh_access.advertise_on_sli](data-sources--nfv_service--reference--group-001.md#canonical-c998b07a678c727ed205032113aa530ae2a29f8bddc0ce40da4715ab813f3b24) |
| `enabled_ssh_access.advertise_on_slo` | [enabled_ssh_access.advertise_on_slo](data-sources--nfv_service--reference--group-001.md#canonical-3231214ec869bdbb3f2055892b1941c0a64204ad838184438627461dd15cc7dc) |
| `enabled_ssh_access.advertise_on_slo_sli` | [enabled_ssh_access.advertise_on_slo_sli](data-sources--nfv_service--reference--group-001.md#canonical-dbb30dfbca68bab3e98fcd37180f0b77113fb789a2fd3f62d75aa1fb650bde9f) |
| `enabled_ssh_access.domain_suffix` | [enabled_ssh_access.domain_suffix](data-sources--nfv_service--reference--group-001.md#canonical-a7a96311ff1d4afcaa65f55d15a7c38ebe86fd7b157f3768df578fbabd5b8121) |
| `enabled_ssh_access.node_ssh_ports` | [enabled_ssh_access.node_ssh_ports](data-sources--nfv_service--reference--group-001.md#canonical-f2de1468056c9dcbe891521c40f68a52f68bbb3b65af150e881b3d4e4cf92b28) |
| `enabled_ssh_access.node_ssh_ports.node_name` | [enabled_ssh_access.node_ssh_ports.node_name](data-sources--nfv_service--reference--group-001.md#canonical-65bacc0483f36beb239dd1543e7a7a483b12403b8afdc6f5d3abe28b55479f73) |
| `enabled_ssh_access.node_ssh_ports.ssh_port` | [enabled_ssh_access.node_ssh_ports.ssh_port](data-sources--nfv_service--reference--group-001.md#canonical-a6ab816db195b2df5591c3f06c279966f93a298d53dbfff4be33efef83133522) |
| `f5_big_ip_aws_service` | [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-409aeae2d2bffa70555dfff342646c37211321fc89ce1d653b59a2993c124f8f) |
| `f5_big_ip_aws_service.admin_password` | [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-2371d1bf60289a76f73e93bffd6f5b9d7d73631bcb1995df329c6876fb55c0c3) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-6830b838863bded295fcae57942d03a383a6b495a5671296ae5c70fe00c068f8) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-001.md#canonical-d0bc7a79aa0790a02b563d355136a4534431626c774c7c0d1a8cca45ac2e502c) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.location` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.location](data-sources--nfv_service--reference--group-001.md#canonical-1c5e159993452c8cac67435b4ce42119b34179fd753e502ed36d79cb58b68a0a) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-001.md#canonical-c3d3f4a5316358da393df089f76bdaebb31c230b0f6b0ea43a602fb1a32adc15) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info` | [f5_big_ip_aws_service.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-623547a6e616fc279daf9bc4403ac469e587eff7059dc7ff08fd5583f1219204) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref` | [f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-001.md#canonical-6730eb09c5774e14693822574eab8d51d7c552381413703e0e5475aacd333cfd) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.url` | [f5_big_ip_aws_service.admin_password.clear_secret_info.url](data-sources--nfv_service--reference--group-001.md#canonical-671c8c1f9ca80d922d01b4de96970c747a3c33c2f8352503b072e3374b919ae4) |
| `f5_big_ip_aws_service.admin_username` | [f5_big_ip_aws_service.admin_username](data-sources--nfv_service--reference--group-001.md#canonical-25c885fd32b9429a486817b78af94718b61e04bcd8edf4b0d0a40933d0636d59) |
| `f5_big_ip_aws_service.aws_tgw_site_params` | [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-4b3bc4886fbfb07e3eac4775310fc312d8a650cb4198ee223eb0932178a18339) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](data-sources--nfv_service--reference--group-001.md#canonical-442389f8d37a4a663269557c3f7508dfd2a25fbedcf7221ca5e22dd7bc4667e0) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name](data-sources--nfv_service--reference--group-001.md#canonical-e3649afd64a44f54fb14e4d245b215f3cd77157a9037de5ce235cfd62c44f8fc) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace](data-sources--nfv_service--reference--group-001.md#canonical-99b5566edbe1cb3760e321652de4e3620461107963d4670a3302ae447ac55fbd) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant](data-sources--nfv_service--reference--group-001.md#canonical-980ccf2b4ea292638b1973b05a366be449295309a4bf154be5a43679536c71d0) |
| `f5_big_ip_aws_service.endpoint_service` | [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-92388c7f8da419462feaced3f98164001845422c904bb5695ada9fc8ba281e28) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-ea1eaf3e25ab443e02b1c0bc44717c94bb4d222b4bf0e81b13681b89abf1cfc1) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](data-sources--nfv_service--reference--group-001.md#canonical-64d83f294d262d0b0107f565219049ace56c75637cc3ddf316045b2794336458) |
| `f5_big_ip_aws_service.endpoint_service.automatic_vip` | [f5_big_ip_aws_service.endpoint_service.automatic_vip](data-sources--nfv_service--reference--group-001.md#canonical-a57c33c2fba13afb112fb533226d5500d320a4f9c8ee96d8cc21e473ecee20a3) |
| `f5_big_ip_aws_service.endpoint_service.configured_vip` | [f5_big_ip_aws_service.endpoint_service.configured_vip](data-sources--nfv_service--reference--group-001.md#canonical-7997a93cd36899c8401cad89ac2e468c951388d550265e011a3b289ab50fc760) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-f3979b3b93bd4285bfd68222a1d72f905880244a4e9fcf81ed119d3b9551235b) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports](data-sources--nfv_service--reference--group-001.md#canonical-c34e6aca9e90abed3b916c2aa9c34cf8ff4c21d4e5df65057bffcc35cd0478ac) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-1b7865f0be9fdf408da690050bee1db3c3d3e162d6794afc11b3e8613b1a6f8b) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports](data-sources--nfv_service--reference--group-001.md#canonical-d2636cac11e73ab172707f8d2f926a15118e3ee2d077ad81a60b1d577ff5590e) |
| `f5_big_ip_aws_service.endpoint_service.default_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-83aa9f7fc53961d73f93638077f759c8d7fd57048458e8666bffa9700c469c35) |
| `f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-1568e1c5b5622c2967f003cd86d6fc1a3ad087b9eb96ba33099eb3836568a3f7) |
| `f5_big_ip_aws_service.endpoint_service.http_port` | [f5_big_ip_aws_service.endpoint_service.http_port](data-sources--nfv_service--reference--group-001.md#canonical-24e58c1aaa29ac7d16cef7e37f5e1c978d096a5b630c062c32c8d5af5311462e) |
| `f5_big_ip_aws_service.endpoint_service.https_port` | [f5_big_ip_aws_service.endpoint_service.https_port](data-sources--nfv_service--reference--group-001.md#canonical-8d4ad6c40c545582d32c11945dcf1651ee4e7881f979ebe2b080f17b6b9b2baf) |
| `f5_big_ip_aws_service.endpoint_service.no_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-f155496f8e320613018987573f9a52c3550e076c8964c57d75c9e90642ffd79c) |
| `f5_big_ip_aws_service.endpoint_service.no_udp_ports` | [f5_big_ip_aws_service.endpoint_service.no_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-deaf528072997f8fb6457862c61b228ff3688f527c4576980f6d56fb3cc893ae) |
| `f5_big_ip_aws_service.market_place_image` | [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-6ce22656a4fa3b6575a063afc55c322c1dce7949b299a0037a766672c32816f7) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](data-sources--nfv_service--reference--group-001.md#canonical-8b4d3397af41303e905cc2d3b48a13c75f3a1240e3ba0fb682d410405398d500) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](data-sources--nfv_service--reference--group-002.md#canonical-c89651796244f6551e8b7fa2f0f954726acbbe25ef777ed511ee8b9813d7033b) |
| `f5_big_ip_aws_service.nodes` | [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-cfa47f171548c18963f853af07ad82af5443d6f9c53baa9b0270f9fde9cafe9e) |
| `f5_big_ip_aws_service.nodes.automatic_prefix` | [f5_big_ip_aws_service.nodes.automatic_prefix](data-sources--nfv_service--reference--group-002.md#canonical-f1f10f5d63d8b0a87c2bcfaeb7b5ff662fa103ef1cbc919e3bede6f0b0c37d6e) |
| `f5_big_ip_aws_service.nodes.aws_az_name` | [f5_big_ip_aws_service.nodes.aws_az_name](data-sources--nfv_service--reference--group-002.md#canonical-baaeae1a41544b6e265ed4916e90eed4f8156c3c81c56cf4520e6e6ff16ae66d) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet` | [f5_big_ip_aws_service.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-ee148c512e7de3d604065153a8eeea3f1638d7acba7b4c580ef0c61e5ff0a2f4) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id` | [f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id](data-sources--nfv_service--reference--group-002.md#canonical-86cf4b294be1a38f18b424bb7cc581b6ab37db5d98bc5a1d0950eca1936be808) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--reference--group-002.md#canonical-1e0251c39ea60d7d7d3bade7ddd1477a842a09f676cd5a8344fbc60beeb91c33) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4](data-sources--nfv_service--reference--group-002.md#canonical-c6b31d332e8c66c1368d779a010375bd8c16903056c405e27cd4d1b52a3f96c4) |
| `f5_big_ip_aws_service.nodes.node_name` | [f5_big_ip_aws_service.nodes.node_name](data-sources--nfv_service--reference--group-002.md#canonical-65de6a9a58c40b708e1fa30c47031c84ca734586256aa7a25b284c92d0d18ba9) |
| `f5_big_ip_aws_service.nodes.reserved_mgmt_subnet` | [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-d8ff73141281b0ba0356a72c537de012c1aeb441a312dbc0a1d8bf0fb28038af) |
| `f5_big_ip_aws_service.nodes.tunnel_prefix` | [f5_big_ip_aws_service.nodes.tunnel_prefix](data-sources--nfv_service--reference--group-002.md#canonical-a85099e65fc028729646e684b340c3861670bff168832f2d8c6955ee85cd4ff1) |
| `f5_big_ip_aws_service.ssh_key` | [f5_big_ip_aws_service.ssh_key](data-sources--nfv_service--reference--group-001.md#canonical-ebd505428bc43fe828d10f0a9d30e8da4a6d3302db56868dd8ce55efae868db9) |
| `f5_big_ip_aws_service.tags` | [f5_big_ip_aws_service.tags](data-sources--nfv_service--reference--group-001.md#canonical-5e8d9f6c685cb673f0fb2b2987a412d1ec7de71e79c563f1435ee7910088f652) |
| `https_management` | [https_management](data-sources--nfv_service--reference--group-002.md#canonical-fd4a121fbf60553361382d690ccb12ad338f58b82514814fbcb963bede6c6944) |
| `https_management.advertise_on_internet` | [https_management.advertise_on_internet](data-sources--nfv_service--reference--group-002.md#canonical-3a1ad6a2dd30b631ce8c5e2324e81ec5e61920f3c9042b41aebbc5c663dccdd7) |
| `https_management.advertise_on_internet.public_ip` | [https_management.advertise_on_internet.public_ip](data-sources--nfv_service--reference--group-002.md#canonical-27f5fab27dbaa703d83d1dd4c29801eb5270f826508d99f24d2a97a3518f46ff) |
| `https_management.advertise_on_internet.public_ip.name` | [https_management.advertise_on_internet.public_ip.name](data-sources--nfv_service--reference--group-002.md#canonical-da02994d0e3a88e121b5c4cbcd660c8bba3e5ac2f38e8d6146625c27c26d0d55) |
| `https_management.advertise_on_internet.public_ip.namespace` | [https_management.advertise_on_internet.public_ip.namespace](data-sources--nfv_service--reference--group-002.md#canonical-21112d4949a2381134be679f140bc0fa4eaf63668402eb418b314d13e0ffd36e) |
| `https_management.advertise_on_internet.public_ip.tenant` | [https_management.advertise_on_internet.public_ip.tenant](data-sources--nfv_service--reference--group-002.md#canonical-456fa9da8325d2b45fefb1239fe3a7dee9226c68a883d7f96e0d0d504e60ba62) |
| `https_management.advertise_on_internet_default_vip` | [https_management.advertise_on_internet_default_vip](data-sources--nfv_service--reference--group-002.md#canonical-623487111cbb318eeb6c751ff782ec75b9dbdcbe8e2b002319b42b6ba8bf7f9c) |
| `https_management.advertise_on_sli_vip` | [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-617d3e85a39c22b32f2f0ec63beccf7bd21be0e448654b9e36c63fc829299ded) |
| `https_management.advertise_on_sli_vip.no_mtls` | [https_management.advertise_on_sli_vip.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-ce56e73411d50115233f57f7fb3145a9161b1cf3523babb2aa329de8f5163436) |
| `https_management.advertise_on_sli_vip.tls_certificates` | [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-8217ec8f93128b419d7abc05501050cb0618bb35c0ef3ca940cd0c791fbe9c5f) |
| `https_management.advertise_on_sli_vip.tls_certificates.certificate_url` | [https_management.advertise_on_sli_vip.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-002.md#canonical-b6d9a237a07f0d282d9f38df94ceaa834ec28a7f3f69c04dc2499af41d2c4422) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-9ffce2de81dc9ac671ba3c6bcccfdb7ffff80ad6f5b99b60d6fc7977577e4f3c) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-94a35b1d4b889aa307ffa2a84e7c279696e8b81d8d5b7fef21b61b573ea5faed) |
| `https_management.advertise_on_sli_vip.tls_certificates.description_spec` | [https_management.advertise_on_sli_vip.tls_certificates.description_spec](data-sources--nfv_service--reference--group-002.md#canonical-537fb4f258ae572f9260eda1209b0bd0bccbf44564d9f2874f36ab9ea80f72b5) |
| `https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-3b8984ed91c1a3fab4a4a48347ca8af7449c32f8dbdc1ade9f88f3981784459a) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key` | [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-832c64e9b321e74824904304cf37c3fa52be58d7c88969cddb583ef7d120ebf7) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-db66e505f7d7798670f3cfd191ec33b67906b91796385e6c8645b3fd6a972ebe) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-002.md#canonical-e3c86ce8cf335439ace3c9fb20367fa6a570f916480def3a50bfbb03241f3a45) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-002.md#canonical-92b08cf668151060f13fe1b7c1620b37ef6855535569e9cf8dfad733feccf921) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-002.md#canonical-7b66bf798882c956c3c5b42f1ccc3b4894447fd8d5dad9e80b06c04491abdd06) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-c37fa882674973a7b39a44e2043aa15a70a75a8d5ff1556491b3439960c18953) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-002.md#canonical-a8c07eb14dcff3cce31842faf8ff01b44b9627b84dd22ca67530724017c0f4b1) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-002.md#canonical-ca40d21865bf8b6c90ee0537cf937047739d2eb6d20880f1234b90491d8db8c1) |
| `https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-089f402daf7850da817d582ccf57b5910be4e28628cb320908124d01ef429a9b) |
| `https_management.advertise_on_sli_vip.tls_config` | [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-15606791ee664c18e6e0f3bba67607a29467a2fd8c67475b77d0c908b8cf86b4) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security` | [https_management.advertise_on_sli_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-002.md#canonical-7b8d00c798347ae3999b7956117a75b220cafd519c56d3c64fd042d44dcd71a4) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-002.md#canonical-c74b5331c8c59d2ae05b6c1109a7580ae75a337ac5068cae501ed6343efb277d) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-002.md#canonical-614150add9d866c692b2dfc792e240bc71921fa50dd0747555571056a532ebec) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-002.md#canonical-bb6d9beb59422c65ab4ebc5515112101899a523261b82ece803d4a505304c436) |
| `https_management.advertise_on_sli_vip.tls_config.default_security` | [https_management.advertise_on_sli_vip.tls_config.default_security](data-sources--nfv_service--reference--group-002.md#canonical-055469d52dddda8c20be3ae21d9b5bcf4e649dae634d28541c75b4e55b7fbf29) |
| `https_management.advertise_on_sli_vip.tls_config.low_security` | [https_management.advertise_on_sli_vip.tls_config.low_security](data-sources--nfv_service--reference--group-002.md#canonical-098503c98aef1caf8114de8ac0c07516f3d3adec3edca6df231ec5530dafdbd7) |
| `https_management.advertise_on_sli_vip.tls_config.medium_security` | [https_management.advertise_on_sli_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-002.md#canonical-42e5f4707adbafb3cebbfaeb84a1782da818aa1ed745918b820db86d5b40e984) |
| `https_management.advertise_on_sli_vip.use_mtls` | [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-41282ee243a138e9d2e5aaab686fd7cf00023932b5bab22dd5f1484ec925417e) |
| `https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-002.md#canonical-7902133232464c99e26c2e080d30eb5cfa172a6cc4e7a9c103f74d0b33333396) |
| `https_management.advertise_on_sli_vip.use_mtls.crl` | [https_management.advertise_on_sli_vip.use_mtls.crl](data-sources--nfv_service--reference--group-002.md#canonical-95d43dd626bd334bcd0f533291014e3fe83fc979731fe9d86526a07d243a3ba4) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.name` | [https_management.advertise_on_sli_vip.use_mtls.crl.name](data-sources--nfv_service--reference--group-002.md#canonical-35d34760c3ef015fe50f5ca9c04e8a1249718185f0d5287d336b99d1db6ca4d9) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.namespace` | [https_management.advertise_on_sli_vip.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-002.md#canonical-ebd9bcd892ac7543c2d644c240f700a53029952eca39d6586de0703c30ccce91) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.tenant` | [https_management.advertise_on_sli_vip.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-002.md#canonical-59761ff4c94e8778a89c0577beb1acc078d1bc8393241611546f016cc327cd1e) |
| `https_management.advertise_on_sli_vip.use_mtls.no_crl` | [https_management.advertise_on_sli_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-002.md#canonical-634960c3aae56271ad89b5fca146eed226a07bf85925b8f5e8f953eaf251eda0) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-5748f0476cc3b5a2b34bc1459c614163303ff2ede137c43fcb88db94fa6483d3) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-002.md#canonical-5a38884dcd8d872af4866460320fe8ed090544496bbe42bd7309e17f1a62b7e7) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-002.md#canonical-18af490b8b7f2ec672a5b498dd60dab83ae9471c1fac00b385812c4a8f2241e5) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-002.md#canonical-a769332707a430594b38fb62ef215e70d1f96efa3d6c58a4bb1a4e8a2ce089bc) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-002.md#canonical-57d89f8e3a777dd84a31a5e0b571922206c7e96c588e91342a8dc75d9a5ddb8d) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-2d81f0f2a05035a62e4488c2ff331d3f4e7272f774790ba084750369c884f47f) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-63158af95a0886672e9573b2c2e09caf843ce19c02469968030cefa69b5af26c) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-002.md#canonical-608ba2bcdddc191c1d537768d99de5fa61799385303542b3df6244ed45116489) |
| `https_management.advertise_on_slo_internet_vip` | [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-63d481ba4f4be297c47c9c63cebc2ce640cbf2a29c56ffd45509e4d72f593461) |
| `https_management.advertise_on_slo_internet_vip.no_mtls` | [https_management.advertise_on_slo_internet_vip.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-ee5627be6eab856b4a638ec02eccd2f65a2aae15468260ead52d1632e4dfb473) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates` | [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5c3eec0736bdb6b1f090153cfa8f6e0007d16d339e37ae0bbd49c337df86b731) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-002.md#canonical-448b1f48dae15d619ab980b8a715b546d486e26d4ce4954ab1c814499a66ed9a) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-d5dab0f6be34a79019fc4979707e30e7f9a5b31b43b36c6f706fd2d6ad5ae170) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-cd895daeefeec2a58de6a8bfed8348f9d4e42b3e9e027d88214f11876cb01fe2) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec](data-sources--nfv_service--reference--group-002.md#canonical-c3bff427c7b4f59d4e66a4bc6769abb1e1f1a3171bd89ff2f82040ae906aa7ed) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-97d29eb5c07de3e88070a8d2d9c68afd00c513da19ed985daaabe2b5b8d548be) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-c508d096f3cd25a2fe82533bd0b673caa5e79c9b08b85dc8755d4f5427891f0d) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-6798fab4bbc53734bd1fde479032848f3ee2d8602861d34021f98538fcab3fd5) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-002.md#canonical-0b2cf4fb0593f5c2f7502296379048e92a87f40cf62dc743d09a730c9d82bd95) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-002.md#canonical-c5652f051ebaf3d443f9a02d4bc462088fc67878ad9441e041e74d19b3feedc8) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-002.md#canonical-87a768f1958a8afddf8539377f4124fd1a01810cd85d95bfb56d4a81d2e35c33) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-7390502a546007c1fa440d8c503a9090df2aa93e8bb61d9da1f461d2fe5f9674) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-002.md#canonical-4008344de00fc61c30d75c7dc892e562d8a43477c8ec770d218fef429d1c55f1) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-002.md#canonical-09e5533779e2f6dce31acc31735f1691fdc3631c3bb4d74a2f0a78560de63cac) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-87b6db5582baf32fc5940587dca4c60e6f1e872cb5797bab7c79555c6f324d71) |
| `https_management.advertise_on_slo_internet_vip.tls_config` | [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-9f4794a7f17fad83392fceb46da5239015a93f311924f1cff4209344351af1af) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-002.md#canonical-5092810a7ff67352951137c0fdcaf00fe53864aa4a57a3b23ecf4f7aceee2b6b) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-002.md#canonical-86dffbd531e3d3411162cd922b28aa95ab044c705fbb787711140993f14bc994) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-002.md#canonical-23df65288c70a8b0075425da31f115352859b190f7969bd8f6641e0628caa062) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-002.md#canonical-2c68acb00eed048ae4b679e196ede55d4015fc298043a82c29deb2862729acc7) |
| `https_management.advertise_on_slo_internet_vip.tls_config.default_security` | [https_management.advertise_on_slo_internet_vip.tls_config.default_security](data-sources--nfv_service--reference--group-002.md#canonical-c49cbddb6ee07ecf2f3c0951f28727334a4cee27db21f534935886237dfb6920) |
| `https_management.advertise_on_slo_internet_vip.tls_config.low_security` | [https_management.advertise_on_slo_internet_vip.tls_config.low_security](data-sources--nfv_service--reference--group-002.md#canonical-e656f410b6dd4d0919cb9ea79b14deba73ab36e54194e73cd1459b1dcaedf5fb) |
| `https_management.advertise_on_slo_internet_vip.tls_config.medium_security` | [https_management.advertise_on_slo_internet_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-002.md#canonical-f0334b77f00644aad38200fc7ef1bd243d259ce4ed7c397688c3b5e40151fefd) |
| `https_management.advertise_on_slo_internet_vip.use_mtls` | [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-f74056e015954befeae3133f9c2fff765468ee360b882e0e623580a210df359d) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-002.md#canonical-6bc84d02ceb3904b319d06b34f2f6f133da3c71b33703c0ab9607f4e2d722086) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl](data-sources--nfv_service--reference--group-002.md#canonical-5049403c77ecc9172cbbd8522860ca499467bcf067085314ba31aa8f7d1db448) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.name](data-sources--nfv_service--reference--group-002.md#canonical-66791cf4e7253d028830b2352a5a7d33663d33c762d4aa197803c4ec80f45d26) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-002.md#canonical-946413cd17ad9650bd70489eb7be231003aa9f7861f65a910ef0366cb53fc7a3) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-002.md#canonical-b2062e6e7e3b0136eaa73d7bef9074a753f8b7f623b111bd0de1df625588c4d7) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-002.md#canonical-ff1d63e4c23a1e9aba6434d44e99f8801f5e937c2337767981a30fffdba41ebb) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-f48000760af114e7332429e399979158e311e69b1305778ac49a4997d15d5b84) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-002.md#canonical-c88c0b322bc1bb82a37f003b60dce495fd9ac2be51a4f5e844d263a0bdd72cba) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-002.md#canonical-5fc2ac55442c4fd23f87fd7109a528b9c70521763df8f4cc97e504d6b36385ab) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-002.md#canonical-0cca91838fbb30d4a4da8677d778af4abf6606b95d849b864593b4d395ce92d5) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-002.md#canonical-23c23911c6c389cb744fef52ca147b2556e2566833649e9228196a7a09025762) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-eb6cc07a4dc44dd20153e369d7f25bf58e171e1c2282760700e8aa6f6af62c7a) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-26fa4dcd0e07437257e9e5f91981cbc73fa24ebd88fd73eb389ef40d7ebc7633) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-002.md#canonical-8510ce34115352996363da133753aafdc95377533fb71651a3f13cc0c15f8df6) |
| `https_management.advertise_on_slo_sli` | [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-dd95e1072b6d17e62d99dbf6406caa7f414e7b01b19a15e08fb71b759342f545) |
| `https_management.advertise_on_slo_sli.no_mtls` | [https_management.advertise_on_slo_sli.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-796ad0dc70e9b53048a21e59bf7804f28777b2feada1e7abb8420eb03da0decf) |
| `https_management.advertise_on_slo_sli.tls_certificates` | [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-ea2fdbf237cbafca299b4bd2e0aacd949b2bf3bcd7c6df7a7e3e6537ee2891ad) |
| `https_management.advertise_on_slo_sli.tls_certificates.certificate_url` | [https_management.advertise_on_slo_sli.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-002.md#canonical-04f831a43d83c9399f1616c67cd23f2873ae036d511ca061986ae98b93e03daf) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-79c7b2be0c40275f2b3fa8001ad50d56cbbb9f93c123fdc5a508afb6e29f9326) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-59f315f59c9fb84c0283304aea52cc6f1ea4a04b8bc68cde140b3b91bf937834) |
| `https_management.advertise_on_slo_sli.tls_certificates.description_spec` | [https_management.advertise_on_slo_sli.tls_certificates.description_spec](data-sources--nfv_service--reference--group-003.md#canonical-eff64a6461c9734b85a36752055e5ec8060aa574c95f544a70c2d617d62629db) |
| `https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-09d64b53b40ba47f5c771ddc77ae8a3641e008197e7e855a64a7ec64020ebb0a) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key` | [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-84b843fc9fafde9d423d3df95bb103f91aac130fa0887188fbe62722663366b2) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-c757c94474ed7b4408987fef3e9dc9146b6b69af92c10a390a0f49c8972a10c7) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-003.md#canonical-d6e7dea4072385e58f4e87e27750467467acb584a4884bd4f12c9959e483c0d3) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-003.md#canonical-5a99515a2735d1c4a103d354521001bbeaa903062079e489bd6e3205ef77f3dd) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-003.md#canonical-262e1c90f40693f002f28505b0b8d9d12c53db1ce27771d652cd0ce440f65fbb) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-d28130f973ec7b7d5a1d4fc264c5b68ccfdd189e75c49c967128fd0b53c9386b) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-003.md#canonical-8a60ffc19bc66d11c256b3107a2cac7fa06ed8d0b8beb09b55502dec3e4b8708) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-003.md#canonical-e9a20bd620feb5bda585962058e7e8ee9dbe2a60cd992033cbf4cef0ee5a0fe1) |
| `https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-45c1d4fd51aaeb2e561880a399ea08d07d4b6d2b385bc9a1ab5e8f716ec6f220) |
| `https_management.advertise_on_slo_sli.tls_config` | [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-23f6066444f57c9f6b51ebfed1dad952e36465a9a8c8414c2a98a5a1beff09eb) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security` | [https_management.advertise_on_slo_sli.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-83d1bc8acd2b912d7f71cfd3ae0e288c741ddc3f60cc8ac407fb6144683e4105) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-003.md#canonical-dd45c0227f88504ba703e6ecf9934e65185b9da24d1787e3e42e8e09c7757047) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-003.md#canonical-859d8746c3718b8831461a25f854aae36da394cf0192706fa2f46f978f688d6e) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-003.md#canonical-e660ab03a0fb8ece217f2e0111bfa39791b4470be604ea079dd4a77a277517da) |
| `https_management.advertise_on_slo_sli.tls_config.default_security` | [https_management.advertise_on_slo_sli.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-1925de80a8d1d12c265e54b4fa215125d3b92e55519630fe804c6af12b7ed06f) |
| `https_management.advertise_on_slo_sli.tls_config.low_security` | [https_management.advertise_on_slo_sli.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-e94fe121dc0e9a30082278c233d763868f436ceba8ff36b8f91a0f4c3c966426) |
| `https_management.advertise_on_slo_sli.tls_config.medium_security` | [https_management.advertise_on_slo_sli.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-f0fb041fd4aa111e0cc1de20273e94445f900afab39f2f0709ebdb3da127bb32) |
| `https_management.advertise_on_slo_sli.use_mtls` | [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-28c641ed8d7d9e9db9973ddbb982cfd99d4cd27dfe01a178bae50903d296bfe6) |
| `https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-003.md#canonical-5b9284e702c6b87173e3e3f209dae9fc2f2934b6459570ee32cfb104ed381f45) |
| `https_management.advertise_on_slo_sli.use_mtls.crl` | [https_management.advertise_on_slo_sli.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-4c94d6bcf025780e615d9c073800462905d890de39316978a9ded8594a4cd89b) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.name` | [https_management.advertise_on_slo_sli.use_mtls.crl.name](data-sources--nfv_service--reference--group-003.md#canonical-f64fb6a2d19d8f2f443c8341bd33d59abfe16221d677f8250278dafa148e889d) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.namespace` | [https_management.advertise_on_slo_sli.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-003.md#canonical-97f5f93866e8a7ac7408b7b75f8baa0e53d4d766fe3511fb8c0c0c116c285685) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.tenant` | [https_management.advertise_on_slo_sli.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-003.md#canonical-84cf096ddccb4e8d34ca5884302fbaae8eba9ec859e2279d63f237b385fdbf80) |
| `https_management.advertise_on_slo_sli.use_mtls.no_crl` | [https_management.advertise_on_slo_sli.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-532821a9450cbd0ad155a7152b1032c85de0ca23a70953c6ca3c3556e497c10c) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-bbe167848ce1428b4b53739b4d1165fe0c165974223c5621e330ecd805b6e556) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-003.md#canonical-4e4bbaa4497827f5bef69e37d6b27a252d95326d292a4b0021ac4b7ccb7abdca) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-003.md#canonical-3dd59b5dcd5f2764028c368bc4135247ee00c4179f610043092ec1c0a39d97c5) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-003.md#canonical-ee7ebd9dfc78dc0a4a9921621677c1f0872042f2d61a6b4a429c839a4fd55bd0) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-003.md#canonical-fe465369f2861e04b09594c8f592b7c7c7b53677f9be7bed030447bd5af34131) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-d3d6e282b61c5fe4c3804375233a947d6f8429e7307206c79e316be8079abd1a) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-72ca314c00ac27ddeb6fd4f5e2af73ed5c73cdd9b5dcd4269569544041662b3b) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-003.md#canonical-dbaed32b004001951e8b80b17af6b9dbd4a4f1879a476c1e87895015c66d18be) |
| `https_management.advertise_on_slo_vip` | [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-61c35ba4d8fdbfa689005914e6e26d93208ce25252f7ccb02c4d234ec1ad5ddc) |
| `https_management.advertise_on_slo_vip.no_mtls` | [https_management.advertise_on_slo_vip.no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-6883bf7ece16d5e676022b1e45c5027f0fd7e2e21524dd675cd586663062b76b) |
| `https_management.advertise_on_slo_vip.tls_certificates` | [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1969382a46156908d4321597462fca7de565b3a5ef949c629b20e9d38ad9d033) |
| `https_management.advertise_on_slo_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_vip.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-003.md#canonical-3aedda1d02e16268eef4acb1d3fa4936b1e6b18d0481e0e51b9e094d83a598bd) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-69925e5b27787fb406173f1baef2609ee6e067f1ca515e255d37fc0b24127ec4) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-b29437ece9e8feddc3b33bed83bb8083850bf46c6683a176240e8cd1b9797dbf) |
| `https_management.advertise_on_slo_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_vip.tls_certificates.description_spec](data-sources--nfv_service--reference--group-003.md#canonical-db80fdb39468faa9bdde6f87e9b53e8b0a8fd21d1ebb40dc2e8ca269f6b35d2f) |
| `https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-8be4e97e13eeaf538cacb8152d72b481ec009380da95db2e15a283559a12f194) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-dbb7055fd40b13c2dbbf28743457850e82e070934e285b7a84c66fbde59b9684) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-c5e35e996905f8653c4a7cfc6c72839ca9702f6955873aa1f2d0425f252c2f55) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-003.md#canonical-da9786d03f5632d68af467c44c8b114bacf1de4ee5132618d79c5b5c26501966) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-003.md#canonical-a1cb096441ddff7f2aae058f8c7d1eeb09473dcee578a3a3bd2846dd4c574b85) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-003.md#canonical-7e822f258f5602a099fe9987e6e66a6d627367c9a43a18ef0ea66168607d6f64) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-b95ab7147060d75d7cb9919afccf50201d16bad99c6b7c38ae58e4e37abf7cf6) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-003.md#canonical-e2e6e2255b3c35c97f03c773fa5967a2f371ae160cc1ec1a96ab390b15241947) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-003.md#canonical-19958dde59a9c703ebd5727fa94b60560fc170ce24295aff6655817a7d245798) |
| `https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-9721119e7389eaac5213504c8fc073f4915627506eaef7ecb81e52bd1beb9013) |
| `https_management.advertise_on_slo_vip.tls_config` | [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3dd7a31bdb7d56eb79206dfbdf702d476f6728852704415e85b63d95579761f4) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security` | [https_management.advertise_on_slo_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-f254988aee1cb018793db493c2e36e81c359cc89e2d2501d852be9b125ba5114) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-003.md#canonical-183785fc7836679673f5c643f02a5415807704b8d6c4a425d510f87e846f4f0b) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-003.md#canonical-b556b1faba25c47824033c992f167f7888a6f9c3dd54128ee4b57f691cda8d3f) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-003.md#canonical-925f8b32a62af57168c15cba292df14b863dad907cddd914beeea15cd501c304) |
| `https_management.advertise_on_slo_vip.tls_config.default_security` | [https_management.advertise_on_slo_vip.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-a0122f48601f4dfab730272e19ebc92ab66ad8bb2fff8a18d84e8678e06bd72d) |
| `https_management.advertise_on_slo_vip.tls_config.low_security` | [https_management.advertise_on_slo_vip.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-df48d07c5fb927c20654622535886ae2c0af08bdfa5a57ee756828a5dde978e5) |
| `https_management.advertise_on_slo_vip.tls_config.medium_security` | [https_management.advertise_on_slo_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-d7b18ddd02eeb8b867d8a9a9474499eaec2b20a12cd1da661f34a01898b10df9) |
| `https_management.advertise_on_slo_vip.use_mtls` | [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-001171af2d1cddec42d24fe8a9da51be52e428376aad717462a3817c46e60c4b) |
| `https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-003.md#canonical-6745f143828dd9edb569211e63b50762f4bc02fd614c8bcd6dfa1ddabe2f20fd) |
| `https_management.advertise_on_slo_vip.use_mtls.crl` | [https_management.advertise_on_slo_vip.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-9945360391efe29d27a15861870fd576ec8e144b6df28c2fed6b2553e701b11d) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_vip.use_mtls.crl.name](data-sources--nfv_service--reference--group-003.md#canonical-38da4ff54d9e162628a8d2cddcd47fa778f6cd1d5cc551e6be888a67aa33a893) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_vip.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-003.md#canonical-db2716585fd36806a5031c95dd35965bb22a38bae471016ce7da639197f81844) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_vip.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-003.md#canonical-c81e4a38c10f87f4df26d52f16c19bb5b7fb15f523f12058c6a944ac6af92b72) |
| `https_management.advertise_on_slo_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-85fdbe0771f84eead3becd612d202fa0552582e52665d9d6dd695d25c38d6755) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-da2ff59bbaf9028b80fd1dd08b9662fc2cffdb4027ee8c01a2cdd8a5e1b7c3fa) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-003.md#canonical-2984a992c1c57fa522827ad9afc0da4f7f13238ad7010fe19881de7eb12c46b5) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-003.md#canonical-8a38cec89c617696a58d8a09bbe5897784576e525c0184c05550f242c4fffeea) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-003.md#canonical-9f18439f34fd54485855b1e66390b4e314dbeb525f2efd40378ab42933ef2df7) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-003.md#canonical-5ef247326fbb14b372cc4ffcebe7e73ece2c3c88fc14b90dead0fc92c37cd6d1) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-b04d2b0603b4833e33682a14dbb0b443a5b580896fe4ff99f2b8abb5503aee72) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-ace47e59ac30efa008552a19eca6181b8decd56737a12f83b0f14e13343b1e6c) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-003.md#canonical-f12018f1580a6326a62594baab14b5b918feba7f0a6efe8e91aebb68afb3908e) |
| `https_management.default_https_port` | [https_management.default_https_port](data-sources--nfv_service--reference--group-003.md#canonical-4db773ebbba7e348c67299c69f54f99ca9339fa846c77664a35ecb00d6611edd) |
| `https_management.domain_suffix` | [https_management.domain_suffix](data-sources--nfv_service--reference--group-002.md#canonical-1d1281b9d20046e1a8e87ce545f615f3d362388d1b5fd03291ccc08e23496f7c) |
| `https_management.https_port` | [https_management.https_port](data-sources--nfv_service--reference--group-002.md#canonical-776b0bae7a9d1863d090b2e595ad414c658f2412a5198a6953acceb8671c8e6b) |
| `id` | [id](data-sources--nfv_service--reference--group-001.md#canonical-c47a43ca99f82e65054589c25a4c3e01109abeba23cbe47d9b946935fe2ab42c) |
| `labels` | [labels](data-sources--nfv_service--reference--group-001.md#canonical-050bcef0aeea62ffa38b0e5cd5011893bbe3a35a0f68092ea34b80594fd9cd9a) |
| `name` | [name](data-sources--nfv_service--reference--group-001.md#canonical-3ba0cbcee4650406f9feb64b67d47bdb3e00e906b67fbfa686e4415a31de3615) |
| `namespace` | [namespace](data-sources--nfv_service--reference--group-001.md#canonical-b46abce70cce9024a85edde02c4f68f14aa60cdb3da6891e4e51e0c17cb7cc90) |
| `palo_alto_fw_service` | [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-a6ed4f07d1a8b4c1b3411155b197eb4eb87fa37ad09b24fab219c1fbf7327634) |
| `palo_alto_fw_service.auto_setup` | [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-8514ad49c69ab7c221bd188e1f9ca8e7e02560cb6bac39051fab7bd7e6fd253f) |
| `palo_alto_fw_service.auto_setup.admin_password` | [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-08bc387b290764c36e6fe7c3623b1a4a728a5a0574d23c40ffedfdce9a195b3c) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-4b6d3d34f44bc2e3bd4ccb2183dc9d7311dce62890145eb395b7efbf525d54c8) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-003.md#canonical-6836682ab9541bb6530ee2f81f762168b92a4284e95323319c1bea16b96b307e) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location](data-sources--nfv_service--reference--group-003.md#canonical-a4a1a5141b3978261e743d88bffca3a8d64ac072be672731135dc0e3f5c2291c) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-003.md#canonical-953cf2ac97613d6c194151a5016cb90d3a5a9a291dd8ecd6ef428d1559f38edf) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-dd28916f5379f7ed2fb92cf6e7aeb107527fc2e7e97be7c83b6bcabda635f745) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-003.md#canonical-6086646350fd9991ca9863ac867659adb593d04c9e0410dc1ec439966e6e2e65) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url](data-sources--nfv_service--reference--group-003.md#canonical-ee1fd1b02ea4f5fb3ebbac324eb9ae1724ee92b5fa002fbfb81536c44ad35584) |
| `palo_alto_fw_service.auto_setup.admin_username` | [palo_alto_fw_service.auto_setup.admin_username](data-sources--nfv_service--reference--group-003.md#canonical-39c673559a1b792e4d438814cbad0c5304facda3422d19731e0e9da94ab8a39e) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys` | [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-8d10c552b11c1f65387f1e00ea58de1807f842e41991be40d30aabad007167c7) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-9c5a492f241810769925bc588d3575a9c335198a28cabb1febfb4bde73b49900) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1d468e3937ce025437880db6a737b105d58b87c3d02a40d3d4b2d6a4468e9369) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-003.md#canonical-bfe7248dc6a8be979cfc4207b7e08484e93c84865dc731d92df58194fd806592) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-003.md#canonical-77bdeccdbef7cf07fe2135838faf0690539a9e79f4ce13546119fb5a8c5e0e68) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-003.md#canonical-af7034f6b05cb7264ba3b359650e17b8f85b7a87b51c14db727f28f479c6bb2f) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-735cf0fd70083c7ab03e8fe5c69e9c84e2e6d03ce9ac6d8d8ff409b1b05f9e14) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-003.md#canonical-39749afed7998d06b092e362df07dbb32669a8110ee9a02fc27bfe4b9371c3b9) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-003.md#canonical-8f0ffa6b8b3cc48cb7305e3ddb584a6ec45ff747d8a4de9a3c314596d8965daf) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key](data-sources--nfv_service--reference--group-003.md#canonical-f7ccc85cc22f7b67536bd9d9f07e371d715b8370b3a681a7838ee11930c938e1) |
| `palo_alto_fw_service.aws_tgw_site` | [palo_alto_fw_service.aws_tgw_site](data-sources--nfv_service--reference--group-003.md#canonical-42ada5d9d2f2437c8bccebd5d4087b692a06344f534861244fd447c2b8e3e009) |
| `palo_alto_fw_service.aws_tgw_site.name` | [palo_alto_fw_service.aws_tgw_site.name](data-sources--nfv_service--reference--group-003.md#canonical-8581643397a5cc4e48374a17361b3142b72ddbcda3da87f391788d5cf43f90ec) |
| `palo_alto_fw_service.aws_tgw_site.namespace` | [palo_alto_fw_service.aws_tgw_site.namespace](data-sources--nfv_service--reference--group-003.md#canonical-ef87214a53c6a1e92671d3deace6f7b7d63dd493c9beade6fa8155dc3c886ba4) |
| `palo_alto_fw_service.aws_tgw_site.tenant` | [palo_alto_fw_service.aws_tgw_site.tenant](data-sources--nfv_service--reference--group-003.md#canonical-7b106d42392342088fd0719c7266e48f6135a91a8af0f66e3ef3f3808c3e0c6a) |
| `palo_alto_fw_service.disable_panaroma` | [palo_alto_fw_service.disable_panaroma](data-sources--nfv_service--reference--group-003.md#canonical-e92d000f62de44f6fa7fb62dc087ae08c990e5f264328b8607b7631b5cfb504a) |
| `palo_alto_fw_service.instance_type` | [palo_alto_fw_service.instance_type](data-sources--nfv_service--reference--group-003.md#canonical-dc6036faf89765945cb9f15b8dca63ce0193de201461e237b0b79fcb47621f36) |
| `palo_alto_fw_service.pan_ami_bundle1` | [palo_alto_fw_service.pan_ami_bundle1](data-sources--nfv_service--reference--group-003.md#canonical-3c089318ac95fea53a2af857196c2a909380ba950b8400927ffe840fc6e5ab1f) |
| `palo_alto_fw_service.pan_ami_bundle2` | [palo_alto_fw_service.pan_ami_bundle2](data-sources--nfv_service--reference--group-003.md#canonical-ab5a02ed414e960cb3086b94c5129b4fe732496078b8c4aa0190e55818b82862) |
| `palo_alto_fw_service.panorama_server` | [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-f6db7a352ae40b4857608d9a0b3f02f5d3dd93aad8cc38d39ff38cb070980b26) |
| `palo_alto_fw_service.panorama_server.authorization_key` | [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-7087b7a61ce087440641a34f89da34388592b3099522bef5504188a347e41833) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-ce5ddadb96c4e802f9b962a466b0dfba8b8fecbbfb6738fbaaceb71fbb5a51b4) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-004.md#canonical-e5086876927f7f872a39128513b1e224af6bdf136be0a90cce9be99d336b962d) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-004.md#canonical-76207d0b8f1dea416285800d5e750f151fc39035f48eda3ff837a66e401fbcc6) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-004.md#canonical-fc7822902ef656890c511162923621130b1e96b3e7e9950c409a44e9428e4882) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-fdc86e5a89e70ea638a477ef56eed52b973ae0f0ed3788ac37f0d71b952770ec) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-004.md#canonical-b8877be8dce08bc69f65ba9ede3806ef022d24c2c46146c7cc8d63fea205e6ed) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url](data-sources--nfv_service--reference--group-004.md#canonical-cf495fb23ac807e2ecbbd460459c9c0c9cb7c104e12de0698fdc21588384f2e7) |
| `palo_alto_fw_service.panorama_server.device_group_name` | [palo_alto_fw_service.panorama_server.device_group_name](data-sources--nfv_service--reference--group-004.md#canonical-e07cc48c251ef5305a6604222a929ad16ed2e9b80b82def33af86eed77d819ad) |
| `palo_alto_fw_service.panorama_server.server` | [palo_alto_fw_service.panorama_server.server](data-sources--nfv_service--reference--group-004.md#canonical-eb9b8a4c7f21bc8b114410f1e5266e3d0f5ebea5347241faf2a11a340d6262ed) |
| `palo_alto_fw_service.panorama_server.template_stack_name` | [palo_alto_fw_service.panorama_server.template_stack_name](data-sources--nfv_service--reference--group-004.md#canonical-fb367e3c5894cdd2018afde322234da88c0c378ecf826e2b28fa1b78116ad5fc) |
| `palo_alto_fw_service.service_nodes` | [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-d52ffbccbca930d63feb98fcaf26bd89750e45a621e15f19d7ac3c79bedc6949) |
| `palo_alto_fw_service.service_nodes.nodes` | [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-89a396ded9969282bfcae6b87a1f87e828a9da251b65d7cdae4247aa1b44f2c9) |
| `palo_alto_fw_service.service_nodes.nodes.aws_az_name` | [palo_alto_fw_service.service_nodes.nodes.aws_az_name](data-sources--nfv_service--reference--group-004.md#canonical-668fcc8b1d83f3f22798be801ea32eed23b2aec8f5b429fb02d4d953cfa6e977) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-b5e618de74f566bbb1216eafff1c9b27ba3446737df2192bab91b9651b31a07d) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id](data-sources--nfv_service--reference--group-004.md#canonical-ee97acd810ab2c844cf45f1a8574783af7d16841494c6b01c6757438cb1ae5e1) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--reference--group-004.md#canonical-8137adcaa39053b44102c1a09914f99fa2daf7197dd2f626fb517fdc704b8ebf) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4](data-sources--nfv_service--reference--group-004.md#canonical-70026c5c42fcf3756a5c2a6625a04ec5f9dfdfeae169fb63ab87f4e280535859) |
| `palo_alto_fw_service.service_nodes.nodes.node_name` | [palo_alto_fw_service.service_nodes.nodes.node_name](data-sources--nfv_service--reference--group-004.md#canonical-176272e380ee6430841d508dc4f18522d3c906922a6424ac82942a47e588c183) |
| `palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-ea05d0a89bc1d338ef2e57a04db338d492c86f7e44ce7c8e175b5666e702a218) |
| `palo_alto_fw_service.ssh_key` | [palo_alto_fw_service.ssh_key](data-sources--nfv_service--reference--group-003.md#canonical-67ead7dd3d136ea024d893c92d324b9bef06b7ed53e699550d8d9940ef61b786) |
| `palo_alto_fw_service.tags` | [palo_alto_fw_service.tags](data-sources--nfv_service--reference--group-003.md#canonical-addb9d5c31a1cc4dd4756f48346a8f2d0379ce1b2e25ed4a4346d8ed9ff7e150) |
| `palo_alto_fw_service.version` | [palo_alto_fw_service.version](data-sources--nfv_service--reference--group-003.md#canonical-ad22288ba823d38e7d0adb39ea2e9c88dc3e1558811b0bf2b45f76ac6101444e) |

<a id="canonical-e74f3c5fffce4b34fb0b2cf1b820353e2cead6ae86ffda72fff27c7baf5c2179"></a>

## Next pages — Property reference / 817c58179c3b / 11

- [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-ef39bab5c106466d3615ee776eb28731edc5d8b52f8971794ea7c483b60be792)
- [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-693c0f81abe3465901899348b5507a2aeb9286cff3eb73bb93c18d56f2c365a6)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-1ca81a9ddb9348c620d239771df702d0b475bf02453eb0a2a936c6227d30b60c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-ef39bab5c106466d3615ee776eb28731edc5d8b52f8971794ea7c483b60be792"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fef26cea876cd1ed94da620317b1772a170a4af8d69f4615935804651b9fc74d"></a>

## disable_https_management — disable_https_management / 09377822cdff / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- disable_https_management

<a id="canonical-5bb1e34aac0ea8f4e80c234c1c7c9c4f5404c7618423475ae6fd72f747132c2c"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_https\_management, https\_management; Default: disable\_https\_management\]
Configuration parameter for disable https management.

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

- [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-5bb1e34aac0ea8f4e80c234c1c7c9c4f5404c7618423475ae6fd72f747132c2c)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-fd4a121fbf60553361382d690ccb12ad338f58b82514814fbcb963bede6c6944)

Select alternatives according to the provider validators above.

<a id="canonical-e8d14568820e5cfb0ceeafd677ff50248e4510a2305303360bbce161fef161ee"></a>

## Direct properties — disable_https_management / 09377822cdff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-494dd5d4c48b50d512b8f4c712538a869261f7ae6aef90be3c26f2689bdb1133"></a>

## Next pages — disable_https_management / 09377822cdff / 4

- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-693c0f81abe3465901899348b5507a2aeb9286cff3eb73bb93c18d56f2c365a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c25005feb534d50c98d11410fab3211327b52de6d09d9d9a25793e1e340111f3"></a>

## disable_ssh_access — disable_ssh_access / 6e42770d8c04 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- disable_ssh_access

<a id="canonical-fac292344ae27a844333e24b4388e6d1bd197e81ca9ae49412159ae88f014b3b"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ssh\_access, enabled\_ssh\_access; Default: disable\_ssh\_access\] Configuration
parameter for disable ssh access.

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

- [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-fac292344ae27a844333e24b4388e6d1bd197e81ca9ae49412159ae88f014b3b)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-edc881fb1eae592b96392b4c9b6ac03d4dfa8555aee864a12a9d7bf47ca86058)

Select alternatives according to the provider validators above.

<a id="canonical-f0a257d4a8f6e71f23e20ddf5144c27846fcbf58f32ecaca170612340e2718de"></a>

## Direct properties — disable_ssh_access / 6e42770d8c04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-754ca600defe5f4d38c6dd2e48960679aa4866ecb896a2965c2b279f6957b5b5"></a>

## Next pages — disable_ssh_access / 6e42770d8c04 / 4

- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32f61f1efa1e47da2a95ebf24ae7cb23f33ce5f4b50793797e7fcf8ada626949"></a>

## enabled_ssh_access — enabled_ssh_access / 420385d5c14d / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- enabled_ssh_access

<a id="canonical-edc881fb1eae592b96392b4c9b6ac03d4dfa8555aee864a12a9d7bf47ca86058"></a>

Type: `"single"`. Computed.

Configuration parameter for enabled ssh access.

Upstream description:

SSH based configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_sli\",\"advertise_on_slo\",\"advertise_on_slo_sli\"]"
}
```

<a id="canonical-6275056080fcf86777531fe475bce922a968e53277a1eeb194fea2e003d72f0c"></a>

## Direct properties — enabled_ssh_access / 420385d5c14d / 3

- [advertise_on_sli](data-sources--nfv_service--reference--group-001.md#canonical-0e4d1dd5bdd615c47570a679afb5ebf6b55a9db35d1e334ef6e3674f3e9d03f3): complete subsection reference.

- [advertise_on_slo](data-sources--nfv_service--reference--group-001.md#canonical-71c3aa4b7f2d8f2d04222c9da6d30fcdf105a928728c92d52d23c5db18110d98): complete subsection reference.

- [advertise_on_slo_sli](data-sources--nfv_service--reference--group-001.md#canonical-d134a0c7b8dff1585ff1ab551ae3c6ef19d44d3d87b5362666bf6d3a4fc53692): complete subsection reference.

<a id="canonical-a7a96311ff1d4afcaa65f55d15a7c38ebe86fd7b157f3768df578fbabd5b8121"></a>

<a id="canonical-2de8cca2b0553e9b4f747af183f32685b3d0790254de2c882465ce6e76f52b7f"></a>

## domain_suffix property — enabled_ssh_access / 420385d5c14d / 4

Type: `"string"`. Computed.

Domain suffix will be used along with node name to form the hostname for SSH node management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [node_ssh_ports](data-sources--nfv_service--reference--group-001.md#canonical-0af4e9eab93bd78ad1491da0bf76196d430021f20e72265d6ee99d6d4eb16bcb): complete subsection reference.

<a id="canonical-38028a77b2eb1536b12be565202b7c1f561b9ad19e308b5dda9ecacf03b20401"></a>

## Next pages — enabled_ssh_access / 420385d5c14d / 5

- [enabled_ssh_access.advertise_on_sli](data-sources--nfv_service--reference--group-001.md#canonical-0e4d1dd5bdd615c47570a679afb5ebf6b55a9db35d1e334ef6e3674f3e9d03f3)
- [enabled_ssh_access.advertise_on_slo](data-sources--nfv_service--reference--group-001.md#canonical-71c3aa4b7f2d8f2d04222c9da6d30fcdf105a928728c92d52d23c5db18110d98)
- [enabled_ssh_access.advertise_on_slo_sli](data-sources--nfv_service--reference--group-001.md#canonical-d134a0c7b8dff1585ff1ab551ae3c6ef19d44d3d87b5362666bf6d3a4fc53692)
- [enabled_ssh_access.node_ssh_ports](data-sources--nfv_service--reference--group-001.md#canonical-0af4e9eab93bd78ad1491da0bf76196d430021f20e72265d6ee99d6d4eb16bcb)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-0e4d1dd5bdd615c47570a679afb5ebf6b55a9db35d1e334ef6e3674f3e9d03f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c02fe523b93122c524ff1ffb386489bcb1d1fa11cab1bd91b3cbc8104a5aba2"></a>

## enabled_ssh_access.advertise_on_sli — enabled_ssh_access.advertise_on_sli / 650c850f01b9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- enabled_ssh_access.advertise_on_sli

<a id="canonical-c998b07a678c727ed205032113aa530ae2a29f8bddc0ce40da4715ab813f3b24"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise on sli.

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

<a id="canonical-ee869052a744a0a11795d46ebe31323a6226f9c831b624bc9bb24205ef0b0edd"></a>

## Direct properties — enabled_ssh_access.advertise_on_sli / 650c850f01b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a1a9fdfebc9d3fdf65e8f53346bbf0b280db2609380e13709ba5bcb11f3326e"></a>

## Next pages — enabled_ssh_access.advertise_on_sli / 650c850f01b9 / 4

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-71c3aa4b7f2d8f2d04222c9da6d30fcdf105a928728c92d52d23c5db18110d98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04882de491cfb79fcd287fb66994e6ccabd080c4452014c131e1e4deede6f841"></a>

## enabled_ssh_access.advertise_on_slo — enabled_ssh_access.advertise_on_slo / d06b758d3ee4 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- enabled_ssh_access.advertise_on_slo

<a id="canonical-3231214ec869bdbb3f2055892b1941c0a64204ad838184438627461dd15cc7dc"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise on slo.

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

<a id="canonical-9e9d54364b2a0a4ba9d03f31edcfdf41886e7e37535feb3d22f37ab4a2612d19"></a>

## Direct properties — enabled_ssh_access.advertise_on_slo / d06b758d3ee4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-434474e3d25e79b40803ac1f85f59c1353c05c1aec903f2600be05df5ce0b918"></a>

## Next pages — enabled_ssh_access.advertise_on_slo / d06b758d3ee4 / 4

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-d134a0c7b8dff1585ff1ab551ae3c6ef19d44d3d87b5362666bf6d3a4fc53692"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb36d00535a24faafd8a6e44a03e0eedc0c373adf83ab0858695c66c665f969c"></a>

## enabled_ssh_access.advertise_on_slo_sli — enabled_ssh_access.advertise_on_slo_sli / b81f7201fd40 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- enabled_ssh_access.advertise_on_slo_sli

<a id="canonical-dbb30dfbca68bab3e98fcd37180f0b77113fb789a2fd3f62d75aa1fb650bde9f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise on slo sli.

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

<a id="canonical-dd3668156cfe901f575efebaab4bb9f414f4241c475a8a3efba588866709321a"></a>

## Direct properties — enabled_ssh_access.advertise_on_slo_sli / b81f7201fd40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6606e0c348c2714e460ab0b12f52efc22ad9c4230df7fb7063101e53c15d3852"></a>

## Next pages — enabled_ssh_access.advertise_on_slo_sli / b81f7201fd40 / 4

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-0af4e9eab93bd78ad1491da0bf76196d430021f20e72265d6ee99d6d4eb16bcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f904011a1ac3cc538affd1953a5afbd1c00460c3187087802a6b723caea4c08"></a>

## enabled_ssh_access.node_ssh_ports — enabled_ssh_access.node_ssh_ports / 1b354e469ae9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- enabled_ssh_access.node_ssh_ports

<a id="canonical-f2de1468056c9dcbe891521c40f68a52f68bbb3b65af150e881b3d4e4cf92b28"></a>

Type: `"list"`. Computed.

Management Node SSH Port. Enter TCP port and node name per node.

Upstream description:

Enter TCP port and node name per node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2"
  }
}
```

<a id="canonical-4b4a7dfd6dc27f0fa56965d8d494f5b9c88f2fd322875374ad1b3519f52d1ed1"></a>

## Direct properties — enabled_ssh_access.node_ssh_ports / 1b354e469ae9 / 3

<a id="canonical-65bacc0483f36beb239dd1543e7a7a483b12403b8afdc6f5d3abe28b55479f73"></a>

<a id="canonical-1cdf2335f037ec40728d49e4bf48c570186b877658a66edee646b44475891a4c"></a>

## node_name property — enabled_ssh_access.node_ssh_ports / 1b354e469ae9 / 4

Type: `"string"`. Computed.

Node name will be used to match a particular node with the desired TCP port.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-a6ab816db195b2df5591c3f06c279966f93a298d53dbfff4be33efef83133522"></a>

<a id="canonical-6a491c606850dd5cf16c4b7eaadaa52025cfe7642a76305c6d7ef6e705dec8e6"></a>

## ssh_port property — enabled_ssh_access.node_ssh_ports / 1b354e469ae9 / 5

Type: `"number"`. Computed.

SSH Port. Enter TCP port per node.

Upstream description:

Enter TCP port per node.

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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0cc65dc0dc8be0a3d42459aaf4429b89a527a7b4cf7fa424e2f46209f3bc275c"></a>

## Next pages — enabled_ssh_access.node_ssh_ports / 1b354e469ae9 / 6

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-bb62691dfecf804c98dfc55011048a333bfa9281687bf76c60cd4fe781a8984b)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7dbfddeff958cbb505a60452c703d7a5c890ca109fd5ae80463a15cc3e6b73a"></a>

## f5_big_ip_aws_service — f5_big_ip_aws_service / 889b11a4c85e / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- f5_big_ip_aws_service

<a id="canonical-409aeae2d2bffa70555dfff342646c37211321fc89ce1d653b59a2993c124f8f"></a>

Type: `"single"`. Computed.

\[OneOf: f5\_big\_ip\_aws\_service, palo\_alto\_fw\_service\] Virtual BIG-IP AWS. Virtual BIG-IP
specification for AWS.

Upstream description:

Virtual BIG-IP specification for AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-image_choice": "[\"market_place_image\"]",
  "x-ves-oneof-field-site_type_choice": "[\"aws_tgw_site_params\"]"
}
```

OneOf alternatives in this subsection:

- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-409aeae2d2bffa70555dfff342646c37211321fc89ce1d653b59a2993c124f8f)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-a6ed4f07d1a8b4c1b3411155b197eb4eb87fa37ad09b24fab219c1fbf7327634)

Select alternatives according to the provider validators above.

<a id="canonical-ebdcc1f73019137111565ea7fedb5c58882ec19a8c6a9f7a871dddd50547d489"></a>

## Direct properties — f5_big_ip_aws_service / 889b11a4c85e / 3

- [admin_password](data-sources--nfv_service--reference--group-001.md#canonical-7915b4c32bf14d9115449ce6f12fd36409b39fca1f903a28f9ab12e015675f7d): complete subsection reference.

<a id="canonical-25c885fd32b9429a486817b78af94718b61e04bcd8edf4b0d0a40933d0636d59"></a>

<a id="canonical-3e5cf26c0fcb3f2be35a5978456dcf2ca544107bc0aa0be94fe5eb8c08b714a1"></a>

## admin_username property — f5_big_ip_aws_service / 889b11a4c85e / 4

Type: `"string"`. Computed.

Admin Username. Admin Username for BIG-IP.

Upstream description:

Admin Username for BIG-IP.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-5ce898edb09bf234e4529a6943630f977bd34e1ed0237dfacd22428a94645c3c): complete subsection reference.

- [endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87): complete subsection reference.

- [market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-a436a856fd52a1a49e3f7d06f8669b081d25cc482b3c200ec9cda77881219996): complete subsection reference.

- [nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb): complete subsection reference.

<a id="canonical-ebd505428bc43fe828d10f0a9d30e8da4a6d3302db56868dd8ce55efae868db9"></a>

<a id="canonical-91723eb4de37078ee7dc074c6953fc04b1e6ab2265611c20b280df32040f4110"></a>

## ssh_key property — f5_big_ip_aws_service / 889b11a4c85e / 5

Type: `"string"`. Computed.

Public SSH key for accessing the Big IP nodes.

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

<a id="canonical-5e8d9f6c685cb673f0fb2b2987a412d1ec7de71e79c563f1435ee7910088f652"></a>

<a id="canonical-e35270f9094fd9be97dc35fd135711c5219f4db694a3c74e68c6cb80c5caf879"></a>

## tags property — f5_big_ip_aws_service / 889b11a4c85e / 6

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

<a id="canonical-93fee408633125b2396972b06b3aab1f4bc55c277aaccc8a46b4e3f57ce4aa81"></a>

## Next pages — f5_big_ip_aws_service / 889b11a4c85e / 7

- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-7915b4c32bf14d9115449ce6f12fd36409b39fca1f903a28f9ab12e015675f7d)
- [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-5ce898edb09bf234e4529a6943630f977bd34e1ed0237dfacd22428a94645c3c)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-a436a856fd52a1a49e3f7d06f8669b081d25cc482b3c200ec9cda77881219996)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-7915b4c32bf14d9115449ce6f12fd36409b39fca1f903a28f9ab12e015675f7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67476e533a3211ee6fd96fb6eb1bfd785aed92a85960a7db628faf0070f0f579"></a>

## f5_big_ip_aws_service.admin_password — f5_big_ip_aws_service.admin_password / 80d88e2d1697 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- f5_big_ip_aws_service.admin_password

<a id="canonical-2371d1bf60289a76f73e93bffd6f5b9d7d73631bcb1995df329c6876fb55c0c3"></a>

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

<a id="canonical-651575dca0a783ee984b8b731c68742a724131436a8ade6e01206cba3afb7ee6"></a>

## Direct properties — f5_big_ip_aws_service.admin_password / 80d88e2d1697 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-a5cee6ba6409c755b121215822057f0887fff5e05327fdf28ce41991d0a2d91a): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-05e660b4c7b11c6bbc04e45839da5e12d8e13a44d3ea9cde67b20141a5658721): complete subsection reference.

<a id="canonical-dd1d7be18dd89251b98fc17b6796e74c619f85ad52413c62674fdbdc7bdf6243"></a>

## Next pages — f5_big_ip_aws_service.admin_password / 80d88e2d1697 / 4

- [f5_big_ip_aws_service.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-a5cee6ba6409c755b121215822057f0887fff5e05327fdf28ce41991d0a2d91a)
- [f5_big_ip_aws_service.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-05e660b4c7b11c6bbc04e45839da5e12d8e13a44d3ea9cde67b20141a5658721)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-a5cee6ba6409c755b121215822057f0887fff5e05327fdf28ce41991d0a2d91a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a76332890d260d741d4414972c92c3d35df24cf25e565ab2a639151c15db710a"></a>

## f5_big_ip_aws_service.admin_password.blindfold_secret_info — f5_big_ip_aws_service.admin_password.blindfold_secret_info / e7bffb1891c5 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-7915b4c32bf14d9115449ce6f12fd36409b39fca1f903a28f9ab12e015675f7d)
- f5_big_ip_aws_service.admin_password.blindfold_secret_info

<a id="canonical-6830b838863bded295fcae57942d03a383a6b495a5671296ae5c70fe00c068f8"></a>

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

<a id="canonical-e8b43680d287caa54674ae15af25fcdb83a8e95d0a499a882416ba05e11a5215"></a>

## Direct properties — f5_big_ip_aws_service.admin_password.blindfold_secret_info / e7bffb1891c5 / 3

<a id="canonical-d0bc7a79aa0790a02b563d355136a4534431626c774c7c0d1a8cca45ac2e502c"></a>

<a id="canonical-a0dd3e22189a950804c13b7da5743a23aa56acdce037b0f3d3ed3163421594c4"></a>

## decryption_provider property — f5_big_ip_aws_service.admin_password.blindfold_secret_info / e7bffb1891c5 / 4

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

<a id="canonical-1c5e159993452c8cac67435b4ce42119b34179fd753e502ed36d79cb58b68a0a"></a>

<a id="canonical-74f91ff2a451bf88264ca90dfd4146a3d563262dd97221821f2bcc13cba0c1e9"></a>

## location property — f5_big_ip_aws_service.admin_password.blindfold_secret_info / e7bffb1891c5 / 5

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

<a id="canonical-c3d3f4a5316358da393df089f76bdaebb31c230b0f6b0ea43a602fb1a32adc15"></a>

<a id="canonical-9e9b197cd62bfc7ac77ae3c7b2474f6c9b8625e3ed77051b0f1c3f670aecfdf9"></a>

## store_provider property — f5_big_ip_aws_service.admin_password.blindfold_secret_info / e7bffb1891c5 / 6

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

<a id="canonical-691a38a8892bea6f321cdf625833157bb5d9493f986e83c14529b92e8e8e70ed"></a>

## Next pages — f5_big_ip_aws_service.admin_password.blindfold_secret_info / e7bffb1891c5 / 7

- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-7915b4c32bf14d9115449ce6f12fd36409b39fca1f903a28f9ab12e015675f7d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-05e660b4c7b11c6bbc04e45839da5e12d8e13a44d3ea9cde67b20141a5658721"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5598dbbf05ee5fefde23db01c0cde0931803629854c608cd9ed4d8ecf4e934f3"></a>

## f5_big_ip_aws_service.admin_password.clear_secret_info — f5_big_ip_aws_service.admin_password.clear_secret_info / 78675d8a2736 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-7915b4c32bf14d9115449ce6f12fd36409b39fca1f903a28f9ab12e015675f7d)
- f5_big_ip_aws_service.admin_password.clear_secret_info

<a id="canonical-623547a6e616fc279daf9bc4403ac469e587eff7059dc7ff08fd5583f1219204"></a>

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

<a id="canonical-3367fe5abfab159d8a444b3f4264571d00f6598401a268af68a086502c52fb88"></a>

## Direct properties — f5_big_ip_aws_service.admin_password.clear_secret_info / 78675d8a2736 / 3

<a id="canonical-6730eb09c5774e14693822574eab8d51d7c552381413703e0e5475aacd333cfd"></a>

<a id="canonical-2bcbc19ce9c6141af8f1b7919afd65b88ac324c67adaf39371fffee88ceed006"></a>

## provider_ref property — f5_big_ip_aws_service.admin_password.clear_secret_info / 78675d8a2736 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-671c8c1f9ca80d922d01b4de96970c747a3c33c2f8352503b072e3374b919ae4"></a>

<a id="canonical-29c89a64911c8e3a14794c63f007de580f03ff4b267ce13e5d9d5ce4afcb7565"></a>

## url property — f5_big_ip_aws_service.admin_password.clear_secret_info / 78675d8a2736 / 5

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

<a id="canonical-88c87c66e33a0b352775702ceefbe215cfc6a1a09bb1c20aedeb9b62ea5f68e4"></a>

## Next pages — f5_big_ip_aws_service.admin_password.clear_secret_info / 78675d8a2736 / 6

- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-7915b4c32bf14d9115449ce6f12fd36409b39fca1f903a28f9ab12e015675f7d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-5ce898edb09bf234e4529a6943630f977bd34e1ed0237dfacd22428a94645c3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2286ee5eeef27635952c14e1e0b6b2a062f7f38ebc882d1ecf2c4b2c6e1aa22e"></a>

## f5_big_ip_aws_service.aws_tgw_site_params — f5_big_ip_aws_service.aws_tgw_site_params / 00128992169d / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- f5_big_ip_aws_service.aws_tgw_site_params

<a id="canonical-4b3bc4886fbfb07e3eac4775310fc312d8a650cb4198ee223eb0932178a18339"></a>

Type: `"single"`. Computed.

BIG-IP AWS TGW Site. BIG-IP AWS TGW site specification.

Upstream description:

BIG-IP AWS TGW site specification.

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

<a id="canonical-f3a59691ce5e54e7bebc7ddc23c8d1a5b966533d382c6328ea79f19ebf0391ba"></a>

## Direct properties — f5_big_ip_aws_service.aws_tgw_site_params / 00128992169d / 3

- [aws_tgw_site](data-sources--nfv_service--reference--group-001.md#canonical-880f1c36d15b9e3ebc18baffe61bd5310bfb483d9e65b2459e6f171a2cc83de7): complete subsection reference.

<a id="canonical-1b553ead8c396cc563b1ec560f95cd1d3072e4b97e26424ef0d7fdf00ea81daa"></a>

## Next pages — f5_big_ip_aws_service.aws_tgw_site_params / 00128992169d / 4

- [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](data-sources--nfv_service--reference--group-001.md#canonical-880f1c36d15b9e3ebc18baffe61bd5310bfb483d9e65b2459e6f171a2cc83de7)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-880f1c36d15b9e3ebc18baffe61bd5310bfb483d9e65b2459e6f171a2cc83de7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46b66bf6a641c69ce231fbd2bee7ac4d71136f5d70d1eb00f282ca800fe5c372"></a>

## f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / efc640e2b919 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-5ce898edb09bf234e4529a6943630f977bd34e1ed0237dfacd22428a94645c3c)
- f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site

<a id="canonical-442389f8d37a4a663269557c3f7508dfd2a25fbedcf7221ca5e22dd7bc4667e0"></a>

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

<a id="canonical-c8da2f67c5b3ff9204fa7ce8bbba4e2bb80573d1c7e11a667626a1fffc9f52a6"></a>

## Direct properties — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / efc640e2b919 / 3

<a id="canonical-e3649afd64a44f54fb14e4d245b215f3cd77157a9037de5ce235cfd62c44f8fc"></a>

<a id="canonical-a7006a2a2dae3a0b9ec5a9031ab477ed7167368b7fa184d07cc1d11eafcd922b"></a>

## name property — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / efc640e2b919 / 4

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

<a id="canonical-99b5566edbe1cb3760e321652de4e3620461107963d4670a3302ae447ac55fbd"></a>

<a id="canonical-9a4dfb257d4b7b79530c1f783dd658c92ebdc3801ce8f959d695da987ee62721"></a>

## namespace property — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / efc640e2b919 / 5

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

<a id="canonical-980ccf2b4ea292638b1973b05a366be449295309a4bf154be5a43679536c71d0"></a>

<a id="canonical-16d961b73ffc0d783d5e033eb57537ae8e8853143895ad6c14a69b9b3e6942f8"></a>

## tenant property — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / efc640e2b919 / 6

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

<a id="canonical-b2312e02a5a9de9ee59eeeae24b071ca78d8f1c42f22dd481bebfcc0f17ce4f6"></a>

## Next pages — f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site / efc640e2b919 / 7

- [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-5ce898edb09bf234e4529a6943630f977bd34e1ed0237dfacd22428a94645c3c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6cd62ca1078bdd3c26cd1a8ab64740833401c93fd7f0301a2f066f9c7933a46"></a>

## f5_big_ip_aws_service.endpoint_service — f5_big_ip_aws_service.endpoint_service / 6533842712ad / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- f5_big_ip_aws_service.endpoint_service

<a id="canonical-92388c7f8da419462feaced3f98164001845422c904bb5695ada9fc8ba281e28"></a>

Type: `"single"`. Computed.

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Upstream description:

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-external_vip_choice": "[\"advertise_on_slo_ip\",\"advertise_on_slo_ip_external\",\"disable_advertise_on_slo_ip\"]",
  "x-ves-oneof-field-inside_vip_choice": "[\"automatic_vip\",\"configured_vip\"]",
  "x-ves-oneof-field-tcp_port_choice": "[\"custom_tcp_ports\",\"default_tcp_ports\",\"http_port\",\"https_port\",\"no_tcp_ports\"]",
  "x-ves-oneof-field-udp_port_choice": "[\"custom_udp_ports\",\"no_udp_ports\"]"
}
```

<a id="canonical-723459d16b89aae1dd74697d45e2c5560d907af3d44a1a3ecfa5927793eb7ba1"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service / 6533842712ad / 3

- [advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-dfa034ab15f7765ce9774945f1d63cfe61754bf8669ffead1eb94a6c6c832f52): complete subsection reference.

- [advertise_on_slo_ip_external](data-sources--nfv_service--reference--group-001.md#canonical-ed49148fbf06876f989e4c93c24ead342dacbca51cf6c2f681150255dda14544): complete subsection reference.

- [automatic_vip](data-sources--nfv_service--reference--group-001.md#canonical-3fb3409cb995d14cd76ddcd218e14cb53b09bfb70aa470e0c1a8eac2d05fbd4f): complete subsection reference.

<a id="canonical-7997a93cd36899c8401cad89ac2e468c951388d550265e011a3b289ab50fc760"></a>

<a id="canonical-b9c6805af59db8278c287d23e31b355596a0a8d6141747aa096710d8d085c8ce"></a>

## configured_vip property — f5_big_ip_aws_service.endpoint_service / 6533842712ad / 4

Type: `"string"`. Computed.

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Upstream description:

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  }
}
```

- [custom_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-27ebb766416bae1d9b2b3cf3fbb279dcc6e6c034c4e13bbb633aec8caf44f931): complete subsection reference.

- [custom_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-aa6a44ec5ded6864f13abc42c2914f9b18539edc9ee42c4077f9d154e516a968): complete subsection reference.

- [default_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-15dcd22fd54b15ded2e84ea997e1c3fa1d7cc0cce8973b32c3df7c8253135ecc): complete subsection reference.

- [disable_advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-1510d970a992ccee67f5addcf55e8ffb555d8dda528a8b240209af6439c7e3cf): complete subsection reference.

- [http_port](data-sources--nfv_service--reference--group-001.md#canonical-d8f0709fb3b1a7cb166f8e97b208906dda3e9c7c1314f405ad351ab9c3e83940): complete subsection reference.

- [https_port](data-sources--nfv_service--reference--group-001.md#canonical-9e998b3d0c47eb1ca5b5c8735fd2425e29f881c0e7eccb5b3cf45746b552870a): complete subsection reference.

- [no_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-8daf3be33f03675148dc6773842f5799ca455c8392b6d153cb1ec698e6c77f2c): complete subsection reference.

- [no_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-65f8f89f0f037d33e626d119b8837c276a8c0913ba9a66590c8eb7ece4a14465): complete subsection reference.

<a id="canonical-e46c7d1677c68b7e0c3b4a606d976f8c0c370ecb1251c8b0afb06f4d5ef30b13"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service / 6533842712ad / 5

- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-dfa034ab15f7765ce9774945f1d63cfe61754bf8669ffead1eb94a6c6c832f52)
- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](data-sources--nfv_service--reference--group-001.md#canonical-ed49148fbf06876f989e4c93c24ead342dacbca51cf6c2f681150255dda14544)
- [f5_big_ip_aws_service.endpoint_service.automatic_vip](data-sources--nfv_service--reference--group-001.md#canonical-3fb3409cb995d14cd76ddcd218e14cb53b09bfb70aa470e0c1a8eac2d05fbd4f)
- [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-27ebb766416bae1d9b2b3cf3fbb279dcc6e6c034c4e13bbb633aec8caf44f931)
- [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-aa6a44ec5ded6864f13abc42c2914f9b18539edc9ee42c4077f9d154e516a968)
- [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-15dcd22fd54b15ded2e84ea997e1c3fa1d7cc0cce8973b32c3df7c8253135ecc)
- [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-1510d970a992ccee67f5addcf55e8ffb555d8dda528a8b240209af6439c7e3cf)
- [f5_big_ip_aws_service.endpoint_service.http_port](data-sources--nfv_service--reference--group-001.md#canonical-d8f0709fb3b1a7cb166f8e97b208906dda3e9c7c1314f405ad351ab9c3e83940)
- [f5_big_ip_aws_service.endpoint_service.https_port](data-sources--nfv_service--reference--group-001.md#canonical-9e998b3d0c47eb1ca5b5c8735fd2425e29f881c0e7eccb5b3cf45746b552870a)
- [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-8daf3be33f03675148dc6773842f5799ca455c8392b6d153cb1ec698e6c77f2c)
- [f5_big_ip_aws_service.endpoint_service.no_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-65f8f89f0f037d33e626d119b8837c276a8c0913ba9a66590c8eb7ece4a14465)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-dfa034ab15f7765ce9774945f1d63cfe61754bf8669ffead1eb94a6c6c832f52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a8ce4142a11b927b1a862e0713f14a2a07be10b21b22d69c71b5cde6738b174"></a>

## f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip / 33c05d953d45 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip

<a id="canonical-ea1eaf3e25ab443e02b1c0bc44717c94bb4d222b4bf0e81b13681b89abf1cfc1"></a>

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

<a id="canonical-5dcfc5ae41709be9a996765fb10950070f2379661d1fdb83f3a32f2c70816ed0"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip / 33c05d953d45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d30dd83e12dc21bb906dfd74edd43abbb5a26ea13d64e17f6f11bc059a8fed54"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip / 33c05d953d45 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-ed49148fbf06876f989e4c93c24ead342dacbca51cf6c2f681150255dda14544"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0e8c46d435d2ae52a45fa048b8ba83e76a1245c60e1d7799951f221ddd6ffa9"></a>

## f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external / 4b776f541457 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external

<a id="canonical-64d83f294d262d0b0107f565219049ace56c75637cc3ddf316045b2794336458"></a>

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

<a id="canonical-3886f51663bb198ec04c548283ad1d5a9ff0a54e3c6c96ba81ef33faa659aa2a"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external / 4b776f541457 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b732610c1a77d213e759acdf6dcdef96dc72f37aa40b257e062b16cada8d01df"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external / 4b776f541457 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-3fb3409cb995d14cd76ddcd218e14cb53b09bfb70aa470e0c1a8eac2d05fbd4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-766fb17ac9566bee6426300fa7a0b374ae56e66e8a69691dd155bc3f339e207c"></a>

## f5_big_ip_aws_service.endpoint_service.automatic_vip — f5_big_ip_aws_service.endpoint_service.automatic_vip / 653b0e21b4a0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.automatic_vip

<a id="canonical-a57c33c2fba13afb112fb533226d5500d320a4f9c8ee96d8cc21e473ecee20a3"></a>

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

<a id="canonical-1d7c4653744d650fb3a4fe03ee4642670c5ba65d573c314fc34d25d603ef1adc"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.automatic_vip / 653b0e21b4a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64fbe9300e5cdbf2b6aabb5e3fcab1365f8b375434ad31e48c3b05a839d57a64"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.automatic_vip / 653b0e21b4a0 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-27ebb766416bae1d9b2b3cf3fbb279dcc6e6c034c4e13bbb633aec8caf44f931"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-616028e0c716523b23e0b04b6c063c2b840e5a464db58458e046a5763707a9c6"></a>

## f5_big_ip_aws_service.endpoint_service.custom_tcp_ports — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 465580f2b4a4 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.custom_tcp_ports

<a id="canonical-f3979b3b93bd4285bfd68222a1d72f905880244a4e9fcf81ed119d3b9551235b"></a>

Type: `"single"`. Computed.

Port Range List. List of port ranges.

Upstream description:

List of port ranges.

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

<a id="canonical-47d2ecce152d05c6084aebe10951ded98595860127979c3cac19fec2c79d2dc1"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 465580f2b4a4 / 3

<a id="canonical-c34e6aca9e90abed3b916c2aa9c34cf8ff4c21d4e5df65057bffcc35cd0478ac"></a>

<a id="canonical-0585a6b9d810969a0851cb831a7faaba7ba2b369a0bc18c17e29511baf404483"></a>

## ports property — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 465580f2b4a4 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-5ce10fc3b4cccfc3992223b5689090852fe907b75f99e58edf0fd8fad3571f47"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.custom_tcp_ports / 465580f2b4a4 / 5

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-aa6a44ec5ded6864f13abc42c2914f9b18539edc9ee42c4077f9d154e516a968"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9d1c427ffea34dc8cb30ae5d96d1dfb486fd9846536580a039ddebab84a785d"></a>

## f5_big_ip_aws_service.endpoint_service.custom_udp_ports — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 0fb2030d5595 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.custom_udp_ports

<a id="canonical-1b7865f0be9fdf408da690050bee1db3c3d3e162d6794afc11b3e8613b1a6f8b"></a>

Type: `"single"`. Computed.

Port Range List. List of port ranges.

Upstream description:

List of port ranges.

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

<a id="canonical-302f9c233f32ffc895c7f99a853d7fca9f6b25e30dbd48e75660bdc0befc900e"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 0fb2030d5595 / 3

<a id="canonical-d2636cac11e73ab172707f8d2f926a15118e3ee2d077ad81a60b1d577ff5590e"></a>

<a id="canonical-eddcc9e657f6d56769271f9f05e20d86c89fb4a7a43b432cd9a2f55878147086"></a>

## ports property — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 0fb2030d5595 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-8b7bc5ab695cf99688e8b5643647a2f132ff7f34e2fe85f079f9ec6447c7aa3c"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.custom_udp_ports / 0fb2030d5595 / 5

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-15dcd22fd54b15ded2e84ea997e1c3fa1d7cc0cce8973b32c3df7c8253135ecc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4c6aa9f3b95fc3e1ff224f4404df7cfececd03e6cdfc3a9c0b40d575c733c42"></a>

## f5_big_ip_aws_service.endpoint_service.default_tcp_ports — f5_big_ip_aws_service.endpoint_service.default_tcp_ports / 65e70a4a3e74 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.default_tcp_ports

<a id="canonical-83aa9f7fc53961d73f93638077f759c8d7fd57048458e8666bffa9700c469c35"></a>

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

<a id="canonical-799ad6367bd93b1cd90bf9b06f6a675cc71ec28ceb0e2d6b719ef2bca0c0a737"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.default_tcp_ports / 65e70a4a3e74 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5a80e213a544644f2aa2dd1910652745990ede23876a26c974603287e083ff5"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.default_tcp_ports / 65e70a4a3e74 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1510d970a992ccee67f5addcf55e8ffb555d8dda528a8b240209af6439c7e3cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9b31da296b161eaf6389257ef9860e46d565e335d90c6826601d5a383a9d2ed"></a>

## f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip — f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip / 15cf5d757e70 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip

<a id="canonical-1568e1c5b5622c2967f003cd86d6fc1a3ad087b9eb96ba33099eb3836568a3f7"></a>

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

<a id="canonical-e9f1f81ce6767c0dfd01b1051793e58e3e5fdac1ed13660bbeeebe53f42a8254"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip / 15cf5d757e70 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3252cabcd5414fef9c7222d5925a24a79b4e908dcb0e1e23dbf72e3e02791797"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip / 15cf5d757e70 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-d8f0709fb3b1a7cb166f8e97b208906dda3e9c7c1314f405ad351ab9c3e83940"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3a5ec3b08d5b1a5c58f84869e08a29fc22b0aaad8760a4f83b865746eb63fca"></a>

## f5_big_ip_aws_service.endpoint_service.http_port — f5_big_ip_aws_service.endpoint_service.http_port / d01d6b93ed50 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.http_port

<a id="canonical-24e58c1aaa29ac7d16cef7e37f5e1c978d096a5b630c062c32c8d5af5311462e"></a>

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

<a id="canonical-49187dbe1640c25deab8bdb4fc6a3ea5f754befc340395332a3cedce3983a0cc"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.http_port / d01d6b93ed50 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17c1d4ec4d4f7cc893b2a7a2a5927df65d5f36b775bbf5ba38c0861beca5e5c6"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.http_port / d01d6b93ed50 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-9e998b3d0c47eb1ca5b5c8735fd2425e29f881c0e7eccb5b3cf45746b552870a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4280ade87d69ed227931ccf999f7490c1f4c03ed1dc6e98576fc7106bdc928b0"></a>

## f5_big_ip_aws_service.endpoint_service.https_port — f5_big_ip_aws_service.endpoint_service.https_port / a56e27a34126 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.https_port

<a id="canonical-8d4ad6c40c545582d32c11945dcf1651ee4e7881f979ebe2b080f17b6b9b2baf"></a>

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

<a id="canonical-e11a5242e5bbe117e52f629247bc9821979a75424a29d091361fc53ed0f88bd9"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.https_port / a56e27a34126 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46f15d8b77824d75378ccdbb91e38207555f793b021f629c5891bc67eaa63076"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.https_port / a56e27a34126 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-8daf3be33f03675148dc6773842f5799ca455c8392b6d153cb1ec698e6c77f2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2ab0a54043fe13ae0ac7e5dc4d4b9b648cd965b462af1c29469aee13f100d0e"></a>

## f5_big_ip_aws_service.endpoint_service.no_tcp_ports — f5_big_ip_aws_service.endpoint_service.no_tcp_ports / eca068845b5f / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.no_tcp_ports

<a id="canonical-f155496f8e320613018987573f9a52c3550e076c8964c57d75c9e90642ffd79c"></a>

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

<a id="canonical-8148f8a69d1de75fd56a14cee6b740715d365f27c9d4dfe5029e970d020440a0"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.no_tcp_ports / eca068845b5f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af9faa44e6d86eca42c5ecaa0660269d3660fb59e381659cb82f0de705914230"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.no_tcp_ports / eca068845b5f / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-65f8f89f0f037d33e626d119b8837c276a8c0913ba9a66590c8eb7ece4a14465"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95718d65339910e8dd35eb374de5b51bad52c18000dbe91eceabe4f11cd797e0"></a>

## f5_big_ip_aws_service.endpoint_service.no_udp_ports — f5_big_ip_aws_service.endpoint_service.no_udp_ports / 1dc3457c8c87 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- f5_big_ip_aws_service.endpoint_service.no_udp_ports

<a id="canonical-deaf528072997f8fb6457862c61b228ff3688f527c4576980f6d56fb3cc893ae"></a>

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

<a id="canonical-c2cd7e73ecb499ea6507d3d0d9e0af9c238e4adcd381992a5573b9d94959dce7"></a>

## Direct properties — f5_big_ip_aws_service.endpoint_service.no_udp_ports / 1dc3457c8c87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-859f433ad6ae60c6dfa32ca02bb269b5ce390366d48b5a25c065f37fe7137dcd"></a>

## Next pages — f5_big_ip_aws_service.endpoint_service.no_udp_ports / 1dc3457c8c87 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-50a580f9c75cf028191e3b912b40a6c1126a1262cb79284d0cce6c8a28ce2e87)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-a436a856fd52a1a49e3f7d06f8669b081d25cc482b3c200ec9cda77881219996"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dc2177ebe86b5c2c1f395cbdb2a77606aebb439b64e2206b33d85617c384d40"></a>

## f5_big_ip_aws_service.market_place_image — f5_big_ip_aws_service.market_place_image / d6063f7e1f1b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- f5_big_ip_aws_service.market_place_image

<a id="canonical-6ce22656a4fa3b6575a063afc55c322c1dce7949b299a0037a766672c32816f7"></a>

Type: `"single"`. Computed.

BIG-IP AWS Pay as You Go Image Selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ami_choice": "[\"AWAFPayG200Mbps\",\"AWAFPayG3Gbps\",\"BestPlusPayG200Mbps\",\"best_plus_payg_1gbps\"]"
}
```

<a id="canonical-6cd3bd4e38978b4ed6b7c637dc58b3b2b8eee4afedb65e51d57cd98017c0634e"></a>

## Direct properties — f5_big_ip_aws_service.market_place_image / d6063f7e1f1b / 3

- [awafpay_g200_mbps](data-sources--nfv_service--reference--group-001.md#canonical-238d3dee05cb6879218db54df38028415b35b9bfdb956dd3dc93122a0dbab3fb): complete subsection reference.

- [awafpay_g3_gbps](data-sources--nfv_service--reference--group-001.md#canonical-819d01b1009a129baacf125ba3daef8d988de6a74461791eae545e25f7491e15): complete subsection reference.

<a id="canonical-d68ae27c4f5e8a4b22eff5e79a23cfd677a15f7b260767521cd1651f3de325c7"></a>

## Next pages — f5_big_ip_aws_service.market_place_image / d6063f7e1f1b / 4

- [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](data-sources--nfv_service--reference--group-001.md#canonical-238d3dee05cb6879218db54df38028415b35b9bfdb956dd3dc93122a0dbab3fb)
- [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](data-sources--nfv_service--reference--group-001.md#canonical-819d01b1009a129baacf125ba3daef8d988de6a74461791eae545e25f7491e15)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-238d3dee05cb6879218db54df38028415b35b9bfdb956dd3dc93122a0dbab3fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa9f90a75ac601004e88a910d95627748358094696d1da026e224586591552d6"></a>

## f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps — f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps / 1b410cdb7fb3 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-a436a856fd52a1a49e3f7d06f8669b081d25cc482b3c200ec9cda77881219996)
- f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps

<a id="canonical-8b4d3397af41303e905cc2d3b48a13c75f3a1240e3ba0fb682d410405398d500"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for AWAFPayG200Mbps.

<a id="canonical-02b29a53c243573d35d7f508261d3da6f9370866b26f0c92d2dfe266c641f365"></a>

## Direct properties — f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps / 1b410cdb7fb3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3fc7254b8b8ed64d4932afc5c5ed22b94be86887148bceb68157be793805715"></a>

## Next pages — f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps / 1b410cdb7fb3 / 4

- [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-a436a856fd52a1a49e3f7d06f8669b081d25cc482b3c200ec9cda77881219996)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-819d01b1009a129baacf125ba3daef8d988de6a74461791eae545e25f7491e15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
