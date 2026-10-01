---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a3ae9f6fbfc9cd19122415b17381fba81d65b2cd421a0c274de42b833bbeeae"></a>

## Property reference — Property reference / 64ddf1547b9b / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- Property reference

<a id="canonical-8ecd6db955e8b2a1d18bdf511fc883be2eac7ea95bc90a16f8064ea2c9221d0b"></a>

## Direct properties — Property reference / 64ddf1547b9b / 3

<a id="canonical-67dc6dd6494487ab563afbf7d771926648a90e3e404bd87ae44750345eb59025"></a>

<a id="canonical-67005b86e0e177c7e0452711a0f6589d577ce943c3d865e678c420beb902a3da"></a>

## address property — Property reference / 64ddf1547b9b / 4

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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

<a id="canonical-5e3b61d50421b33f0fcd089f21311ff10e1cae7e12264037e51acda31a2a4338"></a>

<a id="canonical-8891844d1195057a2fff96c4b3ebaddcd819f07939b55bc7120034634a8f5c45"></a>

## annotations property — Property reference / 64ddf1547b9b / 5

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

- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268): complete subsection reference.

- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-841de88f947fc15c741935c6d03b445cc3c5554ac154f99139bd0a3a7b677df1): complete subsection reference.

- [coordinates](data-sources--securemesh_site--reference--group-001.md#canonical-453e191bc6a950c24f51c59928005a09198b6c3033f1c9518951b091fc1341ab): complete subsection reference.

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972): complete subsection reference.

- [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-7e085f3ad6c9fb89c8cfc6ff19fb392285a2603ac77d4dd7b0e375dc356a927b): complete subsection reference.

- [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-4e781730aeb3365cbabdb13378ad0597028e4cc344ef384ea9aacd873120700c): complete subsection reference.

<a id="canonical-488a5c135740873aeb444e6d310d88dc76651b1eae8f3cba92d2bbf927a8e139"></a>

<a id="canonical-fd864caa109eba6ab4abf5d1a2f2361de19301d4379128050c5517b1fa542702"></a>

## description property — Property reference / 64ddf1547b9b / 6

Type: `"string"`. Computed.

Description of the SecuremeshSite.

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

<a id="canonical-c09da0338dc0738280e3b21c792a4877df7e5c1302eb2ae60eaff0e7af8fe8e9"></a>

<a id="canonical-51827bfcdfc7bc973074f81235929bfa7e2a64707126e07f413dfdca1f181e57"></a>

## id property — Property reference / 64ddf1547b9b / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-83d8f3c103e9b4e003920cb5bb1f6a494e3b29668f42a5cff516cd52296f4ff6): complete subsection reference.

<a id="canonical-50e31b53298e5a67dc8300b5350c286d5d34d1272ab8ad54039c96c137d3c74b"></a>

<a id="canonical-afa02d82397a2457f7425f81c4e98f0e8374bd410ebe7d97391861d26be48b1b"></a>

## labels property — Property reference / 64ddf1547b9b / 8

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

- [log_receiver](data-sources--securemesh_site--reference--group-004.md#canonical-25a004c0a571194d280d0753b0198bb507fae68b6f61506d561c3989e1223298): complete subsection reference.

- [logs_streaming_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-8cb91c79ca6282c5bcb4daf621b84c6cfcd849040511c71ce6d34405be9cc52f): complete subsection reference.

- [master_node_configuration](data-sources--securemesh_site--reference--group-004.md#canonical-f74288a040c9011d71b7267d7062a4d11f6d05fb98f6b042913290bebbaeef60): complete subsection reference.

<a id="canonical-dbacd0fe3d41ce5bfeaeeb51b84f8f069d78e756a6c16f00887d7362e688cab9"></a>

<a id="canonical-01f98a07e287bdcc3504dacf5d38ef3cad90cfcf5d6e15683767fe99c9b6a11b"></a>

## name property — Property reference / 64ddf1547b9b / 9

Type: `"string"`. Required.

Name of the SecuremeshSite.

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

<a id="canonical-84b5539816aad4d75dccb1d74905e48145aed7e75dc2aa0cb852d85c7e97b466"></a>

<a id="canonical-9dc208aa2171795be9a6712a9b699691192f1c89d0f19662451d09c8e20d828b"></a>

## namespace property — Property reference / 64ddf1547b9b / 10

Type: `"string"`. Required.

Namespace where the SecuremeshSite exists.

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

- [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-435d0c35aadf004489d3208c56f9e62f423f31dcadf2a40c12e51462792c4f94): complete subsection reference.

- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-884e701254c6ef1bbae4be609d2d7d37afeaa4158edbb79039686cc738d6d1b7): complete subsection reference.

- [os](data-sources--securemesh_site--reference--group-004.md#canonical-c664d15c6f04ea3ddc6a580e9d06fe1f0d9bd1fe8e5a553823ec44b74e9d8f51): complete subsection reference.

- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-163cc5bf58deda8edca6ed8d8c88799e7fc5738084b837872603989a33811f3d): complete subsection reference.

- [sw](data-sources--securemesh_site--reference--group-004.md#canonical-513ed30f6d5ab12e1983c05e7d56a9a7cbcb457e3dbd101e2c1084f2f87ea6dd): complete subsection reference.

<a id="canonical-6007d67b80dadc4a497597eda346fb1c822f28e8f08bc112c100ecd24ac44982"></a>

<a id="canonical-fad58454743bb0edd159f725bd7a14d3ea10ca56a0926dd2aebbb06f5b76ac5a"></a>

## volterra_certified_hw property — Property reference / 64ddf1547b9b / 11

Type: `"string"`. Computed.

Name for generic server certified hardware to form this Secure Mesh site.

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

- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-317535c0f3bb7d3c1a957057dca82f817481524abc75232bd2b6ebedf17de214): complete subsection reference.

<a id="canonical-7474ff872265befc125ea46c12987e76ae530d0bfb5e52c9ec005d7123a79d87"></a>

<a id="canonical-dcd784580f4a8782d83ef60226956923f6bc9f8b8a58dddb8bce932e8b2033f1"></a>

## worker_nodes property — Property reference / 64ddf1547b9b / 12

Type: `["list", "string"]`. Computed.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

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

<a id="canonical-2b17359e30458652eda20fb753e309dbd64109b13975ce0565ec16570b40bd91"></a>

## All schema paths — Property reference / 64ddf1547b9b / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--securemesh_site--reference--group-001.md#canonical-67dc6dd6494487ab563afbf7d771926648a90e3e404bd87ae44750345eb59025) |
| `annotations` | [annotations](data-sources--securemesh_site--reference--group-001.md#canonical-5e3b61d50421b33f0fcd089f21311ff10e1cae7e12264037e51acda31a2a4338) |
| `blocked_services` | [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-8e4e361ce83d0a99922e9652e87c92b39c3dda55af0199d15e8a6f7fda460cd3) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-cf6307d696588da0fcfa183389c4017117eb46baa41f41d7ffd68a9edab348b9) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--securemesh_site--reference--group-001.md#canonical-ff5bd209e3bb8e850895db2d26417dc812d59eecd156b15e72d08df9e556d047) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--securemesh_site--reference--group-001.md#canonical-f064f20db9e6ab5c4c70594c76633d51c1f6d4dde21dce7e2b707aa451450bef) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--securemesh_site--reference--group-001.md#canonical-17480858251ddc3be4ff7b194aafd522e575343465501630c3780b6dbab01bcd) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--securemesh_site--reference--group-001.md#canonical-a0fd2e39bb0c157bcb67d9f3aeedcd68a8601f4da0f298afccf8fe7fe08f1f65) |
| `bond_device_list` | [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-874051c27ae46bb8ddf6b8496fe71044327be3c3da300e9cc6e04ba4779c20f9) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-d807b67480088722e3b08a8acb4063e0db2becc4ba4591d63a9b2615ff8f7552) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](data-sources--securemesh_site--reference--group-001.md#canonical-c0ea866a810a6e4d94e7d875d727bd887eba9a05c51ddab869bc226a6ef3f6c1) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](data-sources--securemesh_site--reference--group-001.md#canonical-131602c93a4aba2f735ae91bb41ba4f47bfd653112a1b953779890807ec6d9fb) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](data-sources--securemesh_site--reference--group-001.md#canonical-e8091ac69e985d59de30107862d1f63c47879cb641ff0495e6d103d5cd2b7030) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](data-sources--securemesh_site--reference--group-001.md#canonical-ce1acc22eb5bee222e32e75b69808c107722fde057d824d59b9702466a26a50e) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](data-sources--securemesh_site--reference--group-001.md#canonical-cd4e398297c52253375e04c9f1fd564d65c286eca9dda6a5d4dd63885a541f4b) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](data-sources--securemesh_site--reference--group-001.md#canonical-ea95b21a0e1cc7af03352219fee16404e5d1294540054d8cbf5fc661755e310d) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](data-sources--securemesh_site--reference--group-001.md#canonical-adb116c66ef8191dfe7a354e9ebf3c94f39a0b021a1f8955d1d8e09916f00bf5) |
| `coordinates` | [coordinates](data-sources--securemesh_site--reference--group-001.md#canonical-340afc7e237be13fe98c39020c412e09bb8321b9b3f61d0a8f037e69ca442b77) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--securemesh_site--reference--group-001.md#canonical-fa9f8fab516d31b3cb5f6fbfb4b2c637ddf1bac4ea68d79ffbb755ce64cf4c44) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--securemesh_site--reference--group-001.md#canonical-350b72715ec464ff26517691afce57a561ac2a475ae8e94c8c4b9b2fbd6f061e) |
| `custom_network_config` | [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-afe0919fdf577e1a8fab74825f2ee0a36ffe194811c77c4dc43c4fc6bbafb487) |
| `custom_network_config.active_enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-ada69fa7afbf9f9cfeec838715fe22e8f97dba4b77bc37e06a87c8a28039c98f) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-20b40fd83ad4334e9a9461b3a6a0dff03fe295a0479628f580c3c23041411d79) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--securemesh_site--reference--group-001.md#canonical-7426ece09b3d07e966827a462d9cf5fdbeae9f588b5d549ed1def9cb74954a2d) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--securemesh_site--reference--group-001.md#canonical-a59266089310cb67d3e7cede3eecb4688ab21d8a9db02610f3b5778f2b1d80ac) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--securemesh_site--reference--group-001.md#canonical-d06fe46401086b57420e86aca22323ca67efe520709e0c5e09affce1b8bc5c44) |
| `custom_network_config.active_forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-6500e4c8f249b4fb8f1073ab76052246c94815be714c819e54533e37c69f395c) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-721298d4ca657febcda6afbfe387949f736ffc9b3b0b1f096be45fbcb3b088bf) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--securemesh_site--reference--group-001.md#canonical-c52fe3c7bf62535454ac8da58475cfd2564c763d6f300d8690918e435ccfc7ed) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--securemesh_site--reference--group-001.md#canonical-c604796d12da1760bd61dbd72f518d89c323160fdf890432546db7663da75334) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--securemesh_site--reference--group-001.md#canonical-72f49247a19713bba1b5ae5c1b95ebc029a8802e078ef7bb4cfd0ead94496a5a) |
| `custom_network_config.active_network_policies` | [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-b474513c0c7b58d36cba6153ea656f9ae006f3d83c331e6ac49c94b2685941fe) |
| `custom_network_config.active_network_policies.network_policies` | [custom_network_config.active_network_policies.network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-f46099ba36f76a193eaa98c92aff412e64dd3b2ea8bb7a563298035217ad19c5) |
| `custom_network_config.active_network_policies.network_policies.name` | [custom_network_config.active_network_policies.network_policies.name](data-sources--securemesh_site--reference--group-001.md#canonical-9de80cca2431833d5dc67e9b104fd12740e122098d3a902b11148d2664ce8fc3) |
| `custom_network_config.active_network_policies.network_policies.namespace` | [custom_network_config.active_network_policies.network_policies.namespace](data-sources--securemesh_site--reference--group-001.md#canonical-0b52705b6aff880183316cca3bd7930bc04db51b9996d1a498338114b59f73ce) |
| `custom_network_config.active_network_policies.network_policies.tenant` | [custom_network_config.active_network_policies.network_policies.tenant](data-sources--securemesh_site--reference--group-001.md#canonical-b2894e7011881786def20c0a363484b9e57a5a66d0ce65fd0a8320eceb58e65a) |
| `custom_network_config.default_config` | [custom_network_config.default_config](data-sources--securemesh_site--reference--group-001.md#canonical-3fd5156c0967948fa466d2d4a1f3c09130a52eee88d50ed4b1174ad382bfd9f0) |
| `custom_network_config.default_interface_config` | [custom_network_config.default_interface_config](data-sources--securemesh_site--reference--group-001.md#canonical-aec9251db0470f5c54be1203d30dbd7f6ac74f9cb08b52b569f2e521509bdfd4) |
| `custom_network_config.default_sli_config` | [custom_network_config.default_sli_config](data-sources--securemesh_site--reference--group-001.md#canonical-7fb883b87221c63872193949300ae0f4d6565f1211148cdda2d6fec4c398c5bf) |
| `custom_network_config.forward_proxy_allow_all` | [custom_network_config.forward_proxy_allow_all](data-sources--securemesh_site--reference--group-002.md#canonical-70ad3e3870995693524f7c56a569bbf10190ed058fc8644de96162edee26f0a0) |
| `custom_network_config.global_network_list` | [custom_network_config.global_network_list](data-sources--securemesh_site--reference--group-002.md#canonical-cbd29a30720847e6e4d71e5f8ac1a8d4543351a851c3d97622b66cf7e5151430) |
| `custom_network_config.global_network_list.global_network_connections` | [custom_network_config.global_network_list.global_network_connections](data-sources--securemesh_site--reference--group-002.md#canonical-54813e896457bc12bcf08efa1350f5949ad62890ad46c7748dc46f980a388ab7) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--securemesh_site--reference--group-002.md#canonical-943c42d89c43e1644bc5f4c427991e299c0f9a0b6a3d32470868469da7982cbc) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--securemesh_site--reference--group-002.md#canonical-4958afa643b69bfecef75940aa5bf3622984f4bdf1cf93a78da3517d424b74e4) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--securemesh_site--reference--group-002.md#canonical-cf43c72841d887abaf82de3c063576570ad56b8d1efb7a1541115a5060f40459) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--securemesh_site--reference--group-002.md#canonical-efdd42c78a790d87635d0c3cb4f07ea83de2bf1c96f6370d2c91bc6dcb95ddbd) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--securemesh_site--reference--group-002.md#canonical-e4cd24d946f4f1f3a24f70dff717cc562c582050c54f079aadf7a23b32e6839e) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--securemesh_site--reference--group-002.md#canonical-ab7b164bf4731c6090420ec913288ad537f204aa3b3cd9ea8d302028ec9aafad) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--securemesh_site--reference--group-002.md#canonical-448bd355cc3b88b95c54c46539c435853e7e5e9ea738efc77ea4a2e59d696dca) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--securemesh_site--reference--group-002.md#canonical-98bb51430aeaee9c8a3b7e33e3a51992b1263e6f4546edfcbff5ae4eec38eee9) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--securemesh_site--reference--group-002.md#canonical-148fb56b360ce9a0ef3d19cbb8189b37cc8d4d930ce957c993a21b0cd3e1f500) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--securemesh_site--reference--group-002.md#canonical-057b76988332eb05ceef91c08f52e7b76c48e0cc8f967dd7bb97fbf395d256e5) |
| `custom_network_config.interface_list` | [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d6e87930b49268af6f05d4ef1887490f064fea39a48c527b68a69c141969714e) |
| `custom_network_config.interface_list.interfaces` | [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-17aa407afe3e2e84a67d206b403f8b8d6ce28237a8ec9de6d5ce00a8bdbea953) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](data-sources--securemesh_site--reference--group-002.md#canonical-02652a830b1130d29d2a8961ec0541b294bbaca1042808ab17f9b4619ccf8c02) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](data-sources--securemesh_site--reference--group-002.md#canonical-a8ba284392ff3e617475d194bf9535a12867546ed4e4002c335e0df24689bdbe) |
| `custom_network_config.interface_list.interfaces.dedicated_interface` | [custom_network_config.interface_list.interfaces.dedicated_interface](data-sources--securemesh_site--reference--group-002.md#canonical-ca7f1ec944322173665a689e759f6ccc6af6c81f68d6ddd88eae16db52695cdf) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](data-sources--securemesh_site--reference--group-002.md#canonical-2aec5c4d02c21f27712ef5a7a189bc97a790ccecbaf0a5323724f7898b2387af) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_interface.device](data-sources--securemesh_site--reference--group-002.md#canonical-16c73a937e6390643e44f1fa3eb755a237df008a56955b098225519cfb18ac8f) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.is_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](data-sources--securemesh_site--reference--group-002.md#canonical-9143c6f012ee3df0e805a23e550bd3585edb184b36808bbdb70a84b88ecfca0e) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](data-sources--securemesh_site--reference--group-002.md#canonical-a923b887554e47891296bfd2c0d225fee73d38c72847c416794cc61b30859b13) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](data-sources--securemesh_site--reference--group-002.md#canonical-bfae62ed95c129f420218e86cb17a5d77eeb36dbc70d4e54b4c03db702904701) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_interface.mtu](data-sources--securemesh_site--reference--group-002.md#canonical-e507d65fdcfcb84e625211a5583660b8ae5c6891f316ef225a1b14faee321245) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_interface.node](data-sources--securemesh_site--reference--group-002.md#canonical-ddf59de2ea51582c806a9cac5799f2f3b41b991878dde9877eb6c4d23a438319) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.not_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](data-sources--securemesh_site--reference--group-002.md#canonical-ea42e327b0521d7130d2e2f403419b71675ba6abc5ee19ce4d06edb54e605d8f) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.priority` | [custom_network_config.interface_list.interfaces.dedicated_interface.priority](data-sources--securemesh_site--reference--group-002.md#canonical-4ba77c492d0e6549263cab9202a28d4762458a6720cababe858abeeacfff7c50) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface` | [custom_network_config.interface_list.interfaces.dedicated_management_interface](data-sources--securemesh_site--reference--group-002.md#canonical-e4a42a6f549170d353b50ecbe367473a3814e5528f864690e49d349bf42f51c7) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](data-sources--securemesh_site--reference--group-002.md#canonical-d1213bd5cdb6183084dd15c239ce0448f43223c72d5f6e4e8770af48f636d800) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.device](data-sources--securemesh_site--reference--group-002.md#canonical-d56b5bab4dfd23db14f6c5369e07385357aa1d481e30d6e6e9273a09bf61e409) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu](data-sources--securemesh_site--reference--group-002.md#canonical-baa82d531cda901cde7912a3df1e0dc861dcf80ba6650f42a65bdf60066a0c26) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.node](data-sources--securemesh_site--reference--group-002.md#canonical-1b704d0b49224eb07b1ae607f4d668444dcd67bc4ffd4a6f7308a1da0404546f) |
| `custom_network_config.interface_list.interfaces.description_spec` | [custom_network_config.interface_list.interfaces.description_spec](data-sources--securemesh_site--reference--group-002.md#canonical-6182eab80a1ef273d2fd69d883618d9fb9e2e8e265493fcbe07c8ff21a3a2495) |
| `custom_network_config.interface_list.interfaces.ethernet_interface` | [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-4a9ae331ee7e4ef258f8f066638cd6b90096bc4c6461c8c649f9a20ad48b01a8) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.cluster` | [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](data-sources--securemesh_site--reference--group-002.md#canonical-49222204bc8f74d3deb7e3c671490a486e61c6ee2b42245d25eb750f859d0134) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.device` | [custom_network_config.interface_list.interfaces.ethernet_interface.device](data-sources--securemesh_site--reference--group-002.md#canonical-d470520acd52dc606fbe5dc665ebba7edc8cb9bf283e58956b1a72024291f3a1) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](data-sources--securemesh_site--reference--group-002.md#canonical-c3fefd30711a733ca7c648699dd4987261c47c11d5414ee651efc8aa4e18c665) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--securemesh_site--reference--group-002.md#canonical-4c6b6f6baaa78441dc7799932142b3d410f7f8eaccf0d769952c393329b37f3e) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](data-sources--securemesh_site--reference--group-002.md#canonical-c24716dd5662e1d24f48f3fd626783cf1172d6cefd7c4454f603587acbe33f99) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](data-sources--securemesh_site--reference--group-002.md#canonical-ccc8439696dc4c71b7557c7dab8fa247908d32878712f994ec2bad58e92255a0) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--securemesh_site--reference--group-002.md#canonical-8e1787b983d2502ec39e8ccdb5068176e5677e3464e50ab41ff9b4e20a54e3ea) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address](data-sources--securemesh_site--reference--group-002.md#canonical-6eba56a07e47ebf9db4583c5c3478cc6e78c6b3c5deab965ef89258e052d954f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address](data-sources--securemesh_site--reference--group-002.md#canonical-bbc9879ae95e5f28be4de86943dd2d61c22a6feef9f303a4b196b456a4e9a4df) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site--reference--group-002.md#canonical-5b372d99395cc00b7aa3bf716d2d8094bf6bfd3f0af1c62268bc02874c3f7acb) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site--reference--group-002.md#canonical-d6b181478fd611e9bdf8fb9ab10151bee4c072a70961bdf61c3143c754b2b230) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix](data-sources--securemesh_site--reference--group-002.md#canonical-fbb0a3b4422a7232282bcdd56a58e5ee249414c49583acc7245b233b203e503f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings](data-sources--securemesh_site--reference--group-002.md#canonical-64083918a07e108c869611934094f7f89036bed7cac06654000d9e317a539551) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site--reference--group-002.md#canonical-09246c32509e86faa4a751980332dcc7d4ffb85fb17af1f1b793aa7585772508) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](data-sources--securemesh_site--reference--group-002.md#canonical-2fb6d07d29f84c3685631bd4f4bda4dc3bb2716f7f8d37d192d247bb82410b81) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](data-sources--securemesh_site--reference--group-002.md#canonical-292c84098e62dbf6e72ff744f566d4c4beb343e4acbd167a8e50dc79fba1e959) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](data-sources--securemesh_site--reference--group-002.md#canonical-32dbe02796dcf33bcccdec983adaac74b0ed544658b5b325481703a8c9f6ce68) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site--reference--group-002.md#canonical-daf7a78ab02ef9420399b2f65a6eebea49ef85a6f816c43bbbeaef9dbf13a32a) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag](data-sources--securemesh_site--reference--group-002.md#canonical-52dcdef73357314b4d8b2bd8ab922e58ab937b6d664ca78a43b6b6b6a814cb23) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-96e0a91c1e9fae7a303444889a6359ac7c6d65ac57164367f94c10216279e7a4) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-8f05653f21a06697ad933d772185f89589010c3e19c729863b7ae967bb63fa5e) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-357fc1118dcde228f8182e3144dca426486cb7492ff484450aefe9f41d990143) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--reference--group-002.md#canonical-403a299739fdf397692ff081e5b89d23766864b3a9807db1b1d3edd2b1f07f71) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](data-sources--securemesh_site--reference--group-002.md#canonical-3f2c897f3d015190a38db873bf2b2a975b913f9b00a5420aa4f88ea7dca12b1f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--reference--group-002.md#canonical-e0ef48ee0c86abc83fb854464334a26bdf0e0ddbe90227596bd381ca41ebbef8) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--securemesh_site--reference--group-002.md#canonical-1cb3cb38188c4421871d24d160ce3f15484c839edd12c4f452562a326e0ce2a6) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site--reference--group-002.md#canonical-a163fb509a2e59abdd32ed727f14cb048ef66073c60d9891fa0c760fc5abf1c1) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](data-sources--securemesh_site--reference--group-002.md#canonical-90f285a7fcf42e2caaac2b96d2b0033188081374263dcc3b2128058cc7524268) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site--reference--group-002.md#canonical-07c04afd06ae4b0571c02fbf8e70fa397dd7fe077b4e8cf396997e19ff4b5f4c) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](data-sources--securemesh_site--reference--group-002.md#canonical-fc4ec81cfd3a7752ea18273e8c136ac21bb43317ec25062d2d7ab53d898102a0) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site--reference--group-002.md#canonical-2352d86a7a216df8b13e035ea85cca6827d02d463fed5f81a8e67f7904cec7c7) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site--reference--group-002.md#canonical-935e948b81294ff11ca5e96fe344cc25c80a3f3dbf1880b8285fb886c2d3811e) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix](data-sources--securemesh_site--reference--group-002.md#canonical-a8ee903c801a54df055321575def4aa773191be16266ba45df9e14cbb3447a3e) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-6378141e7f0690dd309caddb0713b2e5d21b58090ba8804f6d35de4e503b8162) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site--reference--group-002.md#canonical-f9953ebb8cd17e09d109b99a1fa1c39975b85b15161fce28e445548dace624b8) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site--reference--group-002.md#canonical-07c452df998ae4e24ef687c83c1bfe4ffdc4b4eaf68314aa268567ee24044ece) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site--reference--group-002.md#canonical-7bda710b572fc7c7c7c03863af887f7bc9ff5b5197e08a7d33ecdf3b929e43c4) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](data-sources--securemesh_site--reference--group-002.md#canonical-b59db00458ebb7bf9079c25197fa4206b636662f3ce5ef1418dd9d6f423aea98) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](data-sources--securemesh_site--reference--group-002.md#canonical-dfdebb05729a88082266679b1cafe743a026bb58124ebdf1355da78baa06ef86) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site--reference--group-003.md#canonical-7f01decd91d1b1087d101dd9015eb34c8a5a45ad51f985d2ba30765318b799f0) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](data-sources--securemesh_site--reference--group-003.md#canonical-bff4039de2f9bd15ec00be8626c2f3e5e1a851b23e32159bd783d56e0610fcaa) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](data-sources--securemesh_site--reference--group-003.md#canonical-a3b36d5a23f8fbfc599d7d5feb8191e3778613e1674246b46e54369ed6081d6d) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-5333a7d660b366958c5347cee77b80642ce8b6eecda350f39013c1c582e84e89) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-3275d37cb0fc1589acfb0f55300d3aa7cb4d58da66957a3467c91c2bd2b773ea) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-693205ba911acb04c2189d205e5ef12b23afbd1e653e54ee084f0bbd83272157) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.is_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](data-sources--securemesh_site--reference--group-003.md#canonical-d9a0c9bddada38952ddea537b0f1e34a4b815df6315e3f1efc5298d236c77d55) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](data-sources--securemesh_site--reference--group-003.md#canonical-83976c7edea019acacc2bd8d4ee018571ab5e7d9553c7dc163c7496ac98f7ea3) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](data-sources--securemesh_site--reference--group-003.md#canonical-4f9cb9f5d9977d35f0dbfd99870317eab9fd37e470d291d0966d48d106d89534) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.mtu` | [custom_network_config.interface_list.interfaces.ethernet_interface.mtu](data-sources--securemesh_site--reference--group-002.md#canonical-6f40277b89311769d4993ad56816ec46368060e2b4283ec0ee31f40c2d8b2179) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-e3ed183dad80ca56240a802e6b25538c39fe3df5eb05bb31d717e682fdbb5682) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.node` | [custom_network_config.interface_list.interfaces.ethernet_interface.node](data-sources--securemesh_site--reference--group-002.md#canonical-986cd314b9a3ed337d682e11615da210070c82c84b052a110a9f5e3c1b87633b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.not_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](data-sources--securemesh_site--reference--group-003.md#canonical-ae4bb61206cb86ad080c93baa19c9863af7feaf4bdd9d70bcd6b2cd46aa0700f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.priority` | [custom_network_config.interface_list.interfaces.ethernet_interface.priority](data-sources--securemesh_site--reference--group-002.md#canonical-27982f54aa129a3321061b222e581bab380f48fec6bbf878e2374cf0c9016654) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](data-sources--securemesh_site--reference--group-003.md#canonical-eb74b147827f3f01c38881184945f081d5e6b75acbc0905bc4d5baf5ea32dff9) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](data-sources--securemesh_site--reference--group-003.md#canonical-584fe5d36c3fe5c93468ac51223c96471f5af955286f8902263f8877e1796f6a) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-3e9be8938e5d2e7109c95df009ae5874e0c0289766f8fe88bfc8348e2741d576) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-ab830cd9c1254e56d774087e21e47eda3b308a09710fa1c368cc0493e80ab883) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-255ddd6acc5bb8e732e90b019097bb5e20e7f99e61b3eff96d0d24960f6b2f1f) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-c97b4054e6ce2a36196d8c9d9c124b97395a4d3c6d142fea1c3706c617bb6671) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw](data-sources--securemesh_site--reference--group-003.md#canonical-2d4525542c8a80b86a356439dd0bc4f1f5151125f999674516b7e1b8e86c8b80) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server](data-sources--securemesh_site--reference--group-003.md#canonical-d88bad542174292a13375003500f3e2a5b8d1b50b17917b871bf58df9514d933) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-9ddebae7912aa71a6c7af82015f648d6ee92f33202f0e25937a1a379e28c5922) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-76309463a9eb4288cc639edd4c97f01b6bfdbc7db497a197f6890df3e5f8183a) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-50344c9373104dedf3571a6c7641ad691725e4a22cd2655b37be84ebcc72412b) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-6900d66811c7d895cee10ea5f5b53670f3b8023285fae310400bc0dc333de83d) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-a1788d446f910e9178656fab814d001e9b9edf0fa4d5f5b5bcccbb40e3f20912) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw](data-sources--securemesh_site--reference--group-003.md#canonical-483587d8a2032d6c121981a6d50399a6853ada690a2d38d4a3fce0e5372ff6af) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server](data-sources--securemesh_site--reference--group-003.md#canonical-682a1d427b71e65cd86bf97a03127068a5ea380cdf95fc9d1ddea9d0735f9bda) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-5864303d62f409df8b2bea6b1c34d083d31bba79e88cdcfaed47e41022567022) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.storage_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](data-sources--securemesh_site--reference--group-003.md#canonical-4a4336861767b4f82e51e07aa8c6c8e49ac8a36e456ebf34f7c40d2d64e222c2) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.untagged` | [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](data-sources--securemesh_site--reference--group-003.md#canonical-9c01f47a493f6625846b721e638783890116d194794269588936503a1ccfde40) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id` | [custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id](data-sources--securemesh_site--reference--group-002.md#canonical-b95c85ba69bec3872b1ebb4376be17c9a5917d69cee000fb00eb3977221b17b5) |
| `custom_network_config.interface_list.interfaces.labels` | [custom_network_config.interface_list.interfaces.labels](data-sources--securemesh_site--reference--group-002.md#canonical-136bba8ce2ea8b50bf9abb30f175a666dbfc86f911ac482efdf165231065f15f) |
| `custom_network_config.no_forward_proxy` | [custom_network_config.no_forward_proxy](data-sources--securemesh_site--reference--group-003.md#canonical-c24c3f389db003ceba4862f4371f5d8e68034899f1f7d78128de7ecbed1a24e7) |
| `custom_network_config.no_global_network` | [custom_network_config.no_global_network](data-sources--securemesh_site--reference--group-003.md#canonical-55f14b72bf209f1a5f886ad8a2b75c5ffd770074f68933fb6a72f616d6b2c421) |
| `custom_network_config.no_network_policy` | [custom_network_config.no_network_policy](data-sources--securemesh_site--reference--group-003.md#canonical-f1b53cebcb3feb816b10b931dfd28d4ed9aaaec0489448dd7f50f527387b2956) |
| `custom_network_config.sli_config` | [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-ad21dacbb4a422228fbbacf96c6fcbac28ae73faf722c956a1d8751301b904ba) |
| `custom_network_config.sli_config.dc_cluster_group` | [custom_network_config.sli_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-d08b3fa380cb3c68b528778ce30e03e8e38ad59511db665b6e98925d55c7bb46) |
| `custom_network_config.sli_config.dc_cluster_group.name` | [custom_network_config.sli_config.dc_cluster_group.name](data-sources--securemesh_site--reference--group-003.md#canonical-73bc5d5ed3ba162f9d8f5e99f9edd1320eaf1aeef9a5bd8af73405d0b7b4155a) |
| `custom_network_config.sli_config.dc_cluster_group.namespace` | [custom_network_config.sli_config.dc_cluster_group.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-cfd541c446cabb90a49b1af7532aafd32b907d12d260b3cf9e29396554e7fd14) |
| `custom_network_config.sli_config.dc_cluster_group.tenant` | [custom_network_config.sli_config.dc_cluster_group.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-a60f093f4c85fae23025969bffe8d0bad09c85818ef28d4b77e6babaed340975) |
| `custom_network_config.sli_config.labels` | [custom_network_config.sli_config.labels](data-sources--securemesh_site--reference--group-003.md#canonical-08b8a1aa7d5ed74d348fd480cd025852f09047d162fdd1024d64bde978395bb3) |
| `custom_network_config.sli_config.nameserver` | [custom_network_config.sli_config.nameserver](data-sources--securemesh_site--reference--group-003.md#canonical-7d6e685d92142ef6dbee013e06baf4ffe4f833ef6290a473dd4ae4fa3ba34929) |
| `custom_network_config.sli_config.no_dc_cluster_group` | [custom_network_config.sli_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-7ec87a5320874fa159378a56ab42222d0a879270b64e560d90effabebccd718b) |
| `custom_network_config.sli_config.no_static_routes` | [custom_network_config.sli_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-ecd077a78286ff354c227bd5b1195f7baa7ba3280bdae5f32744608bdd99d922) |
| `custom_network_config.sli_config.no_v6_static_routes` | [custom_network_config.sli_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-6081515011db362567dd29b8452de982587e54c7dfa225e4f89d6fe9df2cc61a) |
| `custom_network_config.sli_config.static_routes` | [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-adb603cec0bc90888ea38115d2ba3ef6e674a3e3891a00ee33deb0fd70ba9f23) |
| `custom_network_config.sli_config.static_routes.static_routes` | [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-e94cf72a48b636be51b388208261d48d1555cf3cd6221f2db7dc2b2cfc18dc83) |
| `custom_network_config.sli_config.static_routes.static_routes.attrs` | [custom_network_config.sli_config.static_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-003.md#canonical-137659a3f6c579862c49de826bb6ffaf05d31ea9297c048865c33345e24f605e) |
| `custom_network_config.sli_config.static_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-6f1fafee96d79afdcd1a8d5b16d1ce96f96b2e922e1d8fac40369bd832226893) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-16b10c5bba3868fb09944cb521b8c1e31bd28d81d3d6642a8ad95cc1588b999c) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-003.md#canonical-f717bb66faf3feed8f987641003b6cc9dfa4ac6b40ce4a2d3fd1f46c4fd91b06) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-4d10f24d2dca60465a0f9bea7fbba60047c6f0ea006da69d6b30168b45846ff1) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-bc1dc6bf9bfc9a71ab0bf847ebc5532dbe9b98b80c27fc006c4b44b061bcbe54) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-bd3d4ca8f03bfa327d340db87d5ad2ecace4cb2035845a6ddf07ee66ba4a57b2) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-003.md#canonical-10de66f004658d6b2a1fa712d5deb861cfc940921f5e7ac94451b79f674895e8) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-003.md#canonical-c50ae05cc75086f9aeca08672ce72ce206315033839493422f325dfb09de1e57) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-c1f88b730341cedea27532e751d3fd8bfb10920597cf58ff8829a782543d81ca) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-66e151c192cfcc4308bb9a1f37ece792213d18998695fbbd872ac863dbfa45d8) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-003.md#canonical-cf54920870cabc4214cc9241acea1cd596be4c0d09ca5a3eb2f54539c5c8d279) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-003.md#canonical-b18f282dc325325982f045a7f2fe8809334b4c7f1083176da467a89ab33b7670) |
| `custom_network_config.sli_config.static_v6_routes` | [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-45d825941c73e246d7a438069d8cb46ea47af85f4e378089f0df3b7faea7d435) |
| `custom_network_config.sli_config.static_v6_routes.static_routes` | [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-5ffcc4765f8903478318a70ca8888fa5a7548145edec140512fe16a912398771) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.attrs` | [custom_network_config.sli_config.static_v6_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-003.md#canonical-a44136706f17dc64287bcbbe12d82be740321c1860616b4845d120939bbc693e) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-9ff71083b1e631079b742fb238062edf0c62323280f2eb931a4e01b32d89bcd8) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-91e04b48e19227ea98983d404b497c38ae103dc4932559c7ef4e87a2bd0d9ee7) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-003.md#canonical-e711b95653f841089e1f6413b1e9c9eac94661b906724511d92166027f0421a2) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-930b7b967774d9e7a1cb08b85607322f7426fb1794faeeed0791ffeeb6918808) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-ee96d256e035486e4daaef3eed84b4149f5070bfe991248de11237294b076696) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-63505ded2366e6514427473c7b70e7b49489d86dba1fa7c62f8c8f514bb415d2) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-003.md#canonical-e8f2c340d574b8347992f184d2ff75665ec782f74fca740afe859dab8fdd00fd) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-003.md#canonical-58ef9dc6754c094e33ea17c027317eb0632818a3ac865ee90a3674129ce98027) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-35cc441b1b0ffdd6249bc9a29305d78d9c1fa2eb6eeea6e31fdcd4c92a0b2761) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-1c772cc372e25465a7be1379224126cf2712fb9c9c243d57c9ca284c893e2ffa) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-003.md#canonical-29ca2689dcb48c4b23a98309d21ddf971f0a3baec84caa7cf2f735575bf21fb6) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-003.md#canonical-e889913afad0cea90ac71bb496a22b7b9b921401617ef7b3084b5e8c2aada853) |
| `custom_network_config.sli_config.vip` | [custom_network_config.sli_config.vip](data-sources--securemesh_site--reference--group-003.md#canonical-1056bdc637939bcb5ec7a19b4b29cb610aacd9573653ae2af80007fcb994f1ce) |
| `custom_network_config.slo_config` | [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-e3b059f3609c33ca360a0878715c906262327bdd93dcc755e0d2d944d8bde980) |
| `custom_network_config.slo_config.dc_cluster_group` | [custom_network_config.slo_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-967ffd7a08952211431144b72deee745fa776c7dc7f4982e10b6c1d855a65dc1) |
| `custom_network_config.slo_config.dc_cluster_group.name` | [custom_network_config.slo_config.dc_cluster_group.name](data-sources--securemesh_site--reference--group-003.md#canonical-8a5a1baadd41579c709460052eb962a10891c7592340693bbd193f83c2096533) |
| `custom_network_config.slo_config.dc_cluster_group.namespace` | [custom_network_config.slo_config.dc_cluster_group.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-8ecada2131f5da55bf129dc362f57793bf637946c8dc3d2ca4bca5fc3936e956) |
| `custom_network_config.slo_config.dc_cluster_group.tenant` | [custom_network_config.slo_config.dc_cluster_group.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-8116cc256a8031da33403eb44680059285397c78b080d10d19175c29945623b1) |
| `custom_network_config.slo_config.labels` | [custom_network_config.slo_config.labels](data-sources--securemesh_site--reference--group-003.md#canonical-a4ca39a26a53a101141f90aade25f331256c15d97956234ebccb138a0975049f) |
| `custom_network_config.slo_config.nameserver` | [custom_network_config.slo_config.nameserver](data-sources--securemesh_site--reference--group-003.md#canonical-a90e9f45f87aac081027077e55c03013f516c54f313307a96f4e58f3f797ccb9) |
| `custom_network_config.slo_config.no_dc_cluster_group` | [custom_network_config.slo_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-1eda53017d7fb33956a16e356459fe493c8515247e7304e6f8be0a7821b3429e) |
| `custom_network_config.slo_config.no_static_routes` | [custom_network_config.slo_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-fe5dcf9ccefb1c6052017eb4ca8b097c55a51f782e55e52efc2213ef02e6ecf2) |
| `custom_network_config.slo_config.no_v6_static_routes` | [custom_network_config.slo_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-690266c99cf90a29b3be1ceda330b9d5650b8b480a5563db3703130813896583) |
| `custom_network_config.slo_config.static_routes` | [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-a4165b957f93bcbbd02833473aee8ddc6d1628848af38ad91d9c0330c7a034c8) |
| `custom_network_config.slo_config.static_routes.static_routes` | [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-f9f9299c7e689d88d81aefa636d5c8620049a1023d4121e1d9246caebf9c365f) |
| `custom_network_config.slo_config.static_routes.static_routes.attrs` | [custom_network_config.slo_config.static_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-003.md#canonical-5e7c239fa5fd00390db6e7e713d6ee9eb30dd11c9613d6f73e76a709aded9c05) |
| `custom_network_config.slo_config.static_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-f946d448c4109e8859e91b0930de75087abfd7644baf5547474d6ca44729d5d5) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-dfc810eca90e875d5135a166f52b82d77a4e31dc1621fc69e5f7e2e077810107) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-003.md#canonical-01582deba78e34b071fefff2db483562c172c1df5c0f314eb7a6ced6acdc7c72) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-35ac2a106b98ac51d4e79f38451206bba0e4dfbead7e9366db7c4c798403af81) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-a82d2e4d99e655751c07eb6219fc7d4d05630a788e1c9a6e879fe3f6bdc16016) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-6599f5d20b584c63491150bfb78ba7b9fdc9f6119a2bb39ec49184483bb12f5b) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-003.md#canonical-b0480853f4d726a558c5ec360a52d902007630392ad761143b154620950ca901) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-003.md#canonical-3c844f8299751bf4fbc53d16c2e0fde436b2502183b464ed5f39b48fa1012361) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-38069ee7489f90b0ed0242c9f2d84bb668265540f94da535ac3c87c197435dc8) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-e68b9204b155d318d5c9f4c41338b5b04b1e32d5cb69d6c9f9ddebbb1ff42f89) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-003.md#canonical-a3da5e05dd13b314f814e4abbe2904748960b5ad2c6beef72d59d917fee6fe00) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-003.md#canonical-3bdf1d146b591d9d6d9a2dad83a39beb2a3ed0c8c60f9917107239a053792aa4) |
| `custom_network_config.slo_config.static_v6_routes` | [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-662415e19a270a5840fd84198c70e2a7e7022221911c96ba7dd02164251acd12) |
| `custom_network_config.slo_config.static_v6_routes.static_routes` | [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-66072100241f318c771d752b8152e1fee090301578eb1dddf61fff7246758684) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.attrs` | [custom_network_config.slo_config.static_v6_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-003.md#canonical-7250af9de5b3c385a73b3b2ea300a4d48d61cb46dc15f3f52d840f03eab4986e) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-004.md#canonical-ca02337d3abfd41bf1c067b32f8be7cd68fbf0438f8be66f2c26cae1dbfaabc1) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-dac2a660d6add17fe6c48cc8b2fd52b0ab2dd8a4eb104a157b432e420edaa0c6) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-003.md#canonical-9288327cfc719dfb0830ff12abe6f9d07dca161a1911e056d6f78239b0e3eaac) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-de9b507a5e0f6a2fc7389017600c3e3604cca5064ef1fc32ec6ccc5b0093c098) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-004.md#canonical-67d83b0a42574113184dbcc826ecba8a66c62dd3c96131a6f0a036dd1f4fb318) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-004.md#canonical-b2468d5579b75b279aebe81f36c7bf4da956f6d07becb3ef782a0acffdf5e474) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-004.md#canonical-ab26bbb3e363e0ed2d39ffbf052f336bcc63bc4a432079425e0a4d98b2f109ee) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-004.md#canonical-43858256fdafcebbf9d682987aa58a3361bca735c28062070252116e6e0a0baf) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-004.md#canonical-dfcc8e11e63160b4b22db3636c284fce542e071f97a88a13147ca8fe18e63cd6) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-004.md#canonical-14b002ad91a8529a4a3b448df5b5593744d9a7156824ad392c2438eafee6786f) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-004.md#canonical-f736453689699cd9034aba98f63dac63bea97519cf26c6596eb3658ffaffc00f) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-004.md#canonical-f1b3fab39ef5f6da1d142393f5f98a54ca4a382d60755d9dba01c4b582ac2f1d) |
| `custom_network_config.slo_config.vip` | [custom_network_config.slo_config.vip](data-sources--securemesh_site--reference--group-003.md#canonical-e4276c21b447f6d3bbbeaa4cf194ad55017e22fdb81cd0fc4db40edc29c559be) |
| `custom_network_config.sm_connection_public_ip` | [custom_network_config.sm_connection_public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-fe462c077b1f520d592bc01f1eb4f5ccfbc7ec76ec8585021b8425aa3c326f43) |
| `custom_network_config.sm_connection_pvt_ip` | [custom_network_config.sm_connection_pvt_ip](data-sources--securemesh_site--reference--group-004.md#canonical-c87f1aff94c61220c30e5bc564362577031a463d984bae4b5630f36acaecb121) |
| `custom_network_config.tunnel_dead_timeout` | [custom_network_config.tunnel_dead_timeout](data-sources--securemesh_site--reference--group-001.md#canonical-2ebc17566c6a966b0219b5567b5f99a6ac57baa82e1302533f267971561ad68d) |
| `custom_network_config.vip_vrrp_mode` | [custom_network_config.vip_vrrp_mode](data-sources--securemesh_site--reference--group-001.md#canonical-eb1e7318e93f27027561c7d42bca3872efae8ee21fe72d8a9e26f1be49f5a36d) |
| `default_blocked_services` | [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-2adfa522ea1ae16b237ba8c9732d7f12e0b39fc2b0fc5cdc2dabf48d189570b9) |
| `default_network_config` | [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-9a932105e6427a1ebc735056852654b2935ede50f1f9ccba8dcd7479b2406715) |
| `description` | [description](data-sources--securemesh_site--reference--group-001.md#canonical-488a5c135740873aeb444e6d310d88dc76651b1eae8f3cba92d2bbf927a8e139) |
| `id` | [id](data-sources--securemesh_site--reference--group-001.md#canonical-c09da0338dc0738280e3b21c792a4877df7e5c1302eb2ae60eaff0e7af8fe8e9) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-24714c837bcee6f7431b1136149b21aaf0182ba5d75e1ed13694db18028127ef) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-20e3815f478023be630f325c854efb597e3ff81363d8ce9059ea583efca8123b) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-739303cb2556a57c9449ed6fb77eb10b889236db2ee7ae9f2abfd2eed4ed4715) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-91663bccef96a821bfa3f9d5569bbbd0add7c3d6b3694c13e4e214afb530b007) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--securemesh_site--reference--group-004.md#canonical-d754fff8b7badff5435e7baf5d263daddfb36943d105ec725356842dc5555a62) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--securemesh_site--reference--group-004.md#canonical-6e156f21a51dfce9a5a203347aecec75c57f88600ab5b43f1bc26b4f6b838fe2) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--securemesh_site--reference--group-004.md#canonical-723744ac9844f225ba545505a8c02627bb5c1b46373f1d16296cc1303607af43) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-b966f5a873dcb195823121daa060a74dc4487671de8b03132aa85df64146fc0f) |
| `labels` | [labels](data-sources--securemesh_site--reference--group-001.md#canonical-50e31b53298e5a67dc8300b5350c286d5d34d1272ab8ad54039c96c137d3c74b) |
| `log_receiver` | [log_receiver](data-sources--securemesh_site--reference--group-004.md#canonical-1277874e91b435578da6ed916045b1e4c0dcbd289513d017a36372d24c88988b) |
| `log_receiver.name` | [log_receiver.name](data-sources--securemesh_site--reference--group-004.md#canonical-9e1fc0324cba82f72d8290b5ade87844e7f837c45068647ce7f03fd1754981c4) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--securemesh_site--reference--group-004.md#canonical-a71d9516608bd29d8322df0efe919074705af72c175127e6e9c14ea48e34fcc4) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--securemesh_site--reference--group-004.md#canonical-71a25fde02889e7f183e68fc6b62611ee5d27ac548a2c654ea51d517af943de8) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-3eff495aa7f2dffb181b89ef0bb0230c88093c44e09d113b2f54be0d77949319) |
| `master_node_configuration` | [master_node_configuration](data-sources--securemesh_site--reference--group-004.md#canonical-ac22f522a1f3ae0533c5a8944ae1974e4f59f2a3410275ad402e34900be7c167) |
| `master_node_configuration.name` | [master_node_configuration.name](data-sources--securemesh_site--reference--group-004.md#canonical-941df3054b45120a0ddb8c29674b02867e931c8ab4a9e0e041a4f0cc0d34069a) |
| `master_node_configuration.public_ip` | [master_node_configuration.public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-e2516adc0cb7e7c5e3ff16905a1e4874b4aba33969744bd37319efb4b0fbf275) |
| `name` | [name](data-sources--securemesh_site--reference--group-001.md#canonical-dbacd0fe3d41ce5bfeaeeb51b84f8f069d78e756a6c16f00887d7362e688cab9) |
| `namespace` | [namespace](data-sources--securemesh_site--reference--group-001.md#canonical-84b5539816aad4d75dccb1d74905e48145aed7e75dc2aa0cb852d85c7e97b466) |
| `no_bond_devices` | [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-a70e2130921a09dbba2932004821b6e96cf45e6c3da10b712c50cbafdee6a09a) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-bcdbafe6ee2019dd955590339f83fbf3d6f632fed6c34032d2056f0c84790da7) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-dab17e0ab7d25d24a3089afa84cd1131b8cc39973ffba4b06535e1f9a9878a08) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-6d7d4408a5b468294769b95823f2db05b30db107124748115a70d1eda367fa8e) |
| `os` | [os](data-sources--securemesh_site--reference--group-004.md#canonical-3cbb210cec709a4e3421aab971c7a4cd3797d186770931fed5303969bdfc2cc2) |
| `os.default_os_version` | [os.default_os_version](data-sources--securemesh_site--reference--group-004.md#canonical-5a1587ce7721991fc914a8bfe6a58069f6262c65cc9eaa5a17d9d49b2e85615b) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--securemesh_site--reference--group-004.md#canonical-76b6c7b8e2913b3210b8ca5daee126c072da152f6e40aff834b99c29275b1adf) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-a89518f07d6cac0ecbace9cd70a160ca497a25bd938b65b9763de5848ff13208) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-8e572abe8e30c830eeb20aabf1ab51bec7f6fe66675c42f5e1f8b09bcbfdedcc) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-d271976c921212596fac5dbab92cfc2bc68f4d1c6430cdd846faf275628a1618) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-8d83216f757e261273db0b04830b354ef0e1e0338712b6b4dc416ce56b66557d) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-7e815311d4f2447fbae4a9b47773cd07305dba7e7a34b4ba3b9dbbc3f2d8a917) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-eda9f384f5601d5ebd99ddbf49f48f2196f10b1c10726fe86d0789be057664b9) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--securemesh_site--reference--group-004.md#canonical-5725fefb07fd4b3d7e0d5cafc85f92c36cd81f890c8cef8db6af706b8a3dfb64) |
| `sw` | [sw](data-sources--securemesh_site--reference--group-004.md#canonical-0d60f18a8d9cf0659350fe88b4064601e11e25d312df93cd735def01a985c84d) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--securemesh_site--reference--group-004.md#canonical-e4b758dced56dd5c12960e8075669165b08b3072486a22e9e9532b537874d44d) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--securemesh_site--reference--group-004.md#canonical-d65ea5bcdd6c0f729745eb398c2ff31b8fbec422990ec95f4b8812db4ff64879) |
| `volterra_certified_hw` | [volterra_certified_hw](data-sources--securemesh_site--reference--group-001.md#canonical-6007d67b80dadc4a497597eda346fb1c822f28e8f08bc112c100ecd24ac44982) |
| `waf_signatures` | [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-974154b56eedf22d02a42096f67e44e5e93972bd59d93c925c2f58d343d8fddc) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--securemesh_site--reference--group-004.md#canonical-dbe3db56091d741e80ec8586087c63546fdc3f4cb9ea324a66488e25238307fd) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--securemesh_site--reference--group-004.md#canonical-0d801a6bfa34a5619653731b706068a413ab1ca2df1af3d68b09c5b8e216c561) |
| `worker_nodes` | [worker_nodes](data-sources--securemesh_site--reference--group-001.md#canonical-7474ff872265befc125ea46c12987e76ae530d0bfb5e52c9ec005d7123a79d87) |

<a id="canonical-277aeeba8aa3a522de7293a949435096f771840b4fd6aeb668e78ad7f2b0502e"></a>

## Next pages — Property reference / 64ddf1547b9b / 14

- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-841de88f947fc15c741935c6d03b445cc3c5554ac154f99139bd0a3a7b677df1)
- [coordinates](data-sources--securemesh_site--reference--group-001.md#canonical-453e191bc6a950c24f51c59928005a09198b6c3033f1c9518951b091fc1341ab)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-7e085f3ad6c9fb89c8cfc6ff19fb392285a2603ac77d4dd7b0e375dc356a927b)
- [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-4e781730aeb3365cbabdb13378ad0597028e4cc344ef384ea9aacd873120700c)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-83d8f3c103e9b4e003920cb5bb1f6a494e3b29668f42a5cff516cd52296f4ff6)
- [log_receiver](data-sources--securemesh_site--reference--group-004.md#canonical-25a004c0a571194d280d0753b0198bb507fae68b6f61506d561c3989e1223298)
- [logs_streaming_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-8cb91c79ca6282c5bcb4daf621b84c6cfcd849040511c71ce6d34405be9cc52f)
- [master_node_configuration](data-sources--securemesh_site--reference--group-004.md#canonical-f74288a040c9011d71b7267d7062a4d11f6d05fb98f6b042913290bebbaeef60)
- [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-435d0c35aadf004489d3208c56f9e62f423f31dcadf2a40c12e51462792c4f94)
- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-884e701254c6ef1bbae4be609d2d7d37afeaa4158edbb79039686cc738d6d1b7)
- [os](data-sources--securemesh_site--reference--group-004.md#canonical-c664d15c6f04ea3ddc6a580e9d06fe1f0d9bd1fe8e5a553823ec44b74e9d8f51)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-163cc5bf58deda8edca6ed8d8c88799e7fc5738084b837872603989a33811f3d)
- [sw](data-sources--securemesh_site--reference--group-004.md#canonical-513ed30f6d5ab12e1983c05e7d56a9a7cbcb457e3dbd101e2c1084f2f87ea6dd)
- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-317535c0f3bb7d3c1a957057dca82f817481524abc75232bd2b6ebedf17de214)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d72d9cdf1e1f42699a21d8ad8087efb9a57762464c384b44c9fa58ae56424cc9"></a>

## blocked_services — blocked_services / 9698d41b24e5 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- blocked_services

<a id="canonical-8e4e361ce83d0a99922e9652e87c92b39c3dda55af0199d15e8a6f7fda460cd3"></a>

Type: `"single"`. Computed.

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

- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-8e4e361ce83d0a99922e9652e87c92b39c3dda55af0199d15e8a6f7fda460cd3)
- [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-2adfa522ea1ae16b237ba8c9732d7f12e0b39fc2b0fc5cdc2dabf48d189570b9)

Select alternatives according to the provider validators above.

<a id="canonical-49075e7b5afdb88d71d3e488d12127470c37f2b12f19b827e6d2d9de8ed07a3d"></a>

## Direct properties — blocked_services / 9698d41b24e5 / 3

- [blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155): complete subsection reference.

<a id="canonical-4930ea99b49f9bdbaad957bdfa0fea5d63292d92f0009948cc76dd775287c580"></a>

## Next pages — blocked_services / 9698d41b24e5 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4927a8d06d4a7e13fbe9b918b46a128515fd06dcdbc4ac2534859e3e45f6cb4"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / a938f762cd4b / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268)
- blocked_services.blocked_service

<a id="canonical-cf6307d696588da0fcfa183389c4017117eb46baa41f41d7ffd68a9edab348b9"></a>

Type: `"list"`. Computed.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

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

<a id="canonical-6d8cbcec79eb0b649da26a979caf791d487b83b15e932c29edad2a7a513df08b"></a>

## Direct properties — blocked_services.blocked_service / a938f762cd4b / 3

- [dns](data-sources--securemesh_site--reference--group-001.md#canonical-5d2525bdcef7eeebda6a297a27680f801855eece53f040c1906b8493923a7d1f): complete subsection reference.

<a id="canonical-f064f20db9e6ab5c4c70594c76633d51c1f6d4dde21dce7e2b707aa451450bef"></a>

<a id="canonical-002709ff97142d3dcb0420294882965c4e9357c41b81c6dcf753770978853fcb"></a>

## network_type property — blocked_services.blocked_service / a938f762cd4b / 4

Type: `"string"`. Computed.

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

- [ssh](data-sources--securemesh_site--reference--group-001.md#canonical-2e44ddc6bcdac04b7e394ffe204f877c4acedef1de3655daaecf97de0342f015): complete subsection reference.

- [web_user_interface](data-sources--securemesh_site--reference--group-001.md#canonical-35c6e5b7c9265adb65a71be92adfaf980de0d5eb1c3a8116ea622b01794d50e0): complete subsection reference.

<a id="canonical-8b04a307f407ab6603ee35226431e0bf685148a21f0d28b71cdf8abed658c5ea"></a>

## Next pages — blocked_services.blocked_service / a938f762cd4b / 5

- [blocked_services.blocked_service.dns](data-sources--securemesh_site--reference--group-001.md#canonical-5d2525bdcef7eeebda6a297a27680f801855eece53f040c1906b8493923a7d1f)
- [blocked_services.blocked_service.ssh](data-sources--securemesh_site--reference--group-001.md#canonical-2e44ddc6bcdac04b7e394ffe204f877c4acedef1de3655daaecf97de0342f015)
- [blocked_services.blocked_service.web_user_interface](data-sources--securemesh_site--reference--group-001.md#canonical-35c6e5b7c9265adb65a71be92adfaf980de0d5eb1c3a8116ea622b01794d50e0)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-5d2525bdcef7eeebda6a297a27680f801855eece53f040c1906b8493923a7d1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b5ebc8a3bc2dfad35775e9072188a5cda1af5a824bc442e4abc840cf0b3d331"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / bac2900716b6 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268)
- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155)
- blocked_services.blocked_service.dns

<a id="canonical-ff5bd209e3bb8e850895db2d26417dc812d59eecd156b15e72d08df9e556d047"></a>

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

<a id="canonical-93ac3ed34457dc4f02ce7ca3a7bee178ef4f11c1162539dbe906d65a463dc50f"></a>

## Direct properties — blocked_services.blocked_service.dns / bac2900716b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab9c21f6f70745f70535cf121b2458feb491c8f352bb3e0974a4d534015e8bd4"></a>

## Next pages — blocked_services.blocked_service.dns / bac2900716b6 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-2e44ddc6bcdac04b7e394ffe204f877c4acedef1de3655daaecf97de0342f015"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fd507828c6c23a2d664fead6305dc1effca1eea50aaaf1e4df003e5c96ac7b2"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / 10c44686a349 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268)
- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155)
- blocked_services.blocked_service.ssh

<a id="canonical-17480858251ddc3be4ff7b194aafd522e575343465501630c3780b6dbab01bcd"></a>

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

<a id="canonical-38bb020e61c93779adf5b953d5e54cb9cf7fc29b01fa0bc3f15381337c76d8c6"></a>

## Direct properties — blocked_services.blocked_service.ssh / 10c44686a349 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54bf225d9fbbbdd1f1d46b88cf26abb72784bfe35419da5f71e600ed1c7568b2"></a>

## Next pages — blocked_services.blocked_service.ssh / 10c44686a349 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-35c6e5b7c9265adb65a71be92adfaf980de0d5eb1c3a8116ea622b01794d50e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33d93192b830e1541fe5cec2653670637f0e92424ddd36b7f04bcb5c2f31ac62"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / 084c10e4019e / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2c9aaf8e9f9ffac8a39834bcba17b3bf7ac8243dce91cff3121621bc44ffd268)
- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-a0fd2e39bb0c157bcb67d9f3aeedcd68a8601f4da0f298afccf8fe7fe08f1f65"></a>

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

<a id="canonical-aad595cddee8dff89418b1ad1874a9e7ae07278fef0f2df54658ee31c3aaa853"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / 084c10e4019e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-629e3be016c8b52808411c15a6d85aac050c7f6b77deb293594cd8bb0d5b009d"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / 084c10e4019e / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-bed4c6388f6f3eecebbeecf782c77e67a5d1a1796a8586b6eeca52720f22a155)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-841de88f947fc15c741935c6d03b445cc3c5554ac154f99139bd0a3a7b677df1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63937f85fa75e56ad09d878c4f16a024ef31aeaf1d0a18ccf8f854afc17566a8"></a>

## bond_device_list — bond_device_list / bad8bea77708 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- bond_device_list

<a id="canonical-874051c27ae46bb8ddf6b8496fe71044327be3c3da300e9cc6e04ba4779c20f9"></a>

Type: `"single"`. Computed.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

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

- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-874051c27ae46bb8ddf6b8496fe71044327be3c3da300e9cc6e04ba4779c20f9)
- [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-a70e2130921a09dbba2932004821b6e96cf45e6c3da10b712c50cbafdee6a09a)

Select alternatives according to the provider validators above.

<a id="canonical-4116da229ed6fbdc4b90e9b9f58c04eae02a1680b3ed8cf89dbd64b53c3e91d1"></a>

## Direct properties — bond_device_list / bad8bea77708 / 3

- [bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-25016bab60e3edeac3f8dd443719bcfb7751b2e4ec4805b4228f9d388f8d0e0f): complete subsection reference.

<a id="canonical-2a4ae4174d3650656804b3bf30fc0ffd663a2cfa93bbc499faf8246cc02617bf"></a>

## Next pages — bond_device_list / bad8bea77708 / 4

- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-25016bab60e3edeac3f8dd443719bcfb7751b2e4ec4805b4228f9d388f8d0e0f)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-25016bab60e3edeac3f8dd443719bcfb7751b2e4ec4805b4228f9d388f8d0e0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c3e5cc62c5d2210d5a4ebbcb8958fd26d9028465659d2f89f6713f1ad76ca48"></a>

## bond_device_list.bond_devices — bond_device_list.bond_devices / 3c0d534c14c0 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-841de88f947fc15c741935c6d03b445cc3c5554ac154f99139bd0a3a7b677df1)
- bond_device_list.bond_devices

<a id="canonical-d807b67480088722e3b08a8acb4063e0db2becc4ba4591d63a9b2615ff8f7552"></a>

Type: `"list"`. Computed.

Bond Devices. List of bond devices.

Upstream description:

List of bond devices.

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

<a id="canonical-48d0eda0880b5eeb081bc77264c7d68a28d3e4e38e87590a5d48f9a8e073f5b0"></a>

## Direct properties — bond_device_list.bond_devices / 3c0d534c14c0 / 3

- [active_backup](data-sources--securemesh_site--reference--group-001.md#canonical-525712f24d6a701ed6eec081abe54f7719c0f04ec2c1ee52e73f6d0535366013): complete subsection reference.

<a id="canonical-131602c93a4aba2f735ae91bb41ba4f47bfd653112a1b953779890807ec6d9fb"></a>

<a id="canonical-679ab1ea44b608958fa244ba0a737be101173d96cea8b3ebfd09b78b81c8c531"></a>

## devices property — bond_device_list.bond_devices / 3c0d534c14c0 / 4

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

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

- [lacp](data-sources--securemesh_site--reference--group-001.md#canonical-dc5a32c09cb51d8a535813fca2dfd43917df17eb5a27d98dde3b89103e1d3ec4): complete subsection reference.

<a id="canonical-cd4e398297c52253375e04c9f1fd564d65c286eca9dda6a5d4dd63885a541f4b"></a>

<a id="canonical-cd3b7d8936e09171d140c3b020987cbd49505bab372aad73f773792470d829ce"></a>

## link_polling_interval property — bond_device_list.bond_devices / 3c0d534c14c0 / 5

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

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

<a id="canonical-ea95b21a0e1cc7af03352219fee16404e5d1294540054d8cbf5fc661755e310d"></a>

<a id="canonical-4a81d3a52b49d0af03171265b14ae9c1912a362d38173f9255fa340a9c9006d4"></a>

## link_up_delay property — bond_device_list.bond_devices / 3c0d534c14c0 / 6

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

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

<a id="canonical-adb116c66ef8191dfe7a354e9ebf3c94f39a0b021a1f8955d1d8e09916f00bf5"></a>

<a id="canonical-34a93ca8495ecdc90bf110f1aed986bb76ec9e0fe02c96bf5a8333041d7c53af"></a>

## name property — bond_device_list.bond_devices / 3c0d534c14c0 / 7

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

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

<a id="canonical-b245f18509184cfdfdaf660c66bda3a2f09d2f9461f212f92cd371f56fa80b08"></a>

## Next pages — bond_device_list.bond_devices / 3c0d534c14c0 / 8

- [bond_device_list.bond_devices.active_backup](data-sources--securemesh_site--reference--group-001.md#canonical-525712f24d6a701ed6eec081abe54f7719c0f04ec2c1ee52e73f6d0535366013)
- [bond_device_list.bond_devices.lacp](data-sources--securemesh_site--reference--group-001.md#canonical-dc5a32c09cb51d8a535813fca2dfd43917df17eb5a27d98dde3b89103e1d3ec4)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-841de88f947fc15c741935c6d03b445cc3c5554ac154f99139bd0a3a7b677df1)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-525712f24d6a701ed6eec081abe54f7719c0f04ec2c1ee52e73f6d0535366013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e85250cfe1da58c7ce581fae38191038bc13df1be4d44898bd5fd8a942b0cea4"></a>

## bond_device_list.bond_devices.active_backup — bond_device_list.bond_devices.active_backup / 5b894334fcf7 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-841de88f947fc15c741935c6d03b445cc3c5554ac154f99139bd0a3a7b677df1)
- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-25016bab60e3edeac3f8dd443719bcfb7751b2e4ec4805b4228f9d388f8d0e0f)
- bond_device_list.bond_devices.active_backup

<a id="canonical-c0ea866a810a6e4d94e7d875d727bd887eba9a05c51ddab869bc226a6ef3f6c1"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b3034eb806db8f91f29e654a103fc80cae840ff7771ae78fe205c787053e776b"></a>

## Direct properties — bond_device_list.bond_devices.active_backup / 5b894334fcf7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76d356db4c8d6e5e32e96aa4cf9efa47493be099e0727c2889e4fe094f9b3ac3"></a>

## Next pages — bond_device_list.bond_devices.active_backup / 5b894334fcf7 / 4

- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-25016bab60e3edeac3f8dd443719bcfb7751b2e4ec4805b4228f9d388f8d0e0f)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-dc5a32c09cb51d8a535813fca2dfd43917df17eb5a27d98dde3b89103e1d3ec4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-996ca1c677dc154ce5bb728745a55a70aac86f614c77aa0c487166387580cf27"></a>

## bond_device_list.bond_devices.lacp — bond_device_list.bond_devices.lacp / ccc4440fce93 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-841de88f947fc15c741935c6d03b445cc3c5554ac154f99139bd0a3a7b677df1)
- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-25016bab60e3edeac3f8dd443719bcfb7751b2e4ec4805b4228f9d388f8d0e0f)
- bond_device_list.bond_devices.lacp

<a id="canonical-e8091ac69e985d59de30107862d1f63c47879cb641ff0495e6d103d5cd2b7030"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

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

<a id="canonical-3fcabdf912b37723fd72e29d7dc91cd281f31b171131ecf4230814a57a21903a"></a>

## Direct properties — bond_device_list.bond_devices.lacp / ccc4440fce93 / 3

<a id="canonical-ce1acc22eb5bee222e32e75b69808c107722fde057d824d59b9702466a26a50e"></a>

<a id="canonical-bb18d0aebdb86f46a41108a9006679a955d5e0c1706ea34f0967c2dedf5f12b0"></a>

## rate property — bond_device_list.bond_devices.lacp / ccc4440fce93 / 4

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

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

<a id="canonical-c022bf0e943c4b99daee283c1a6efed93130492a34f49ae87631e16f8606471b"></a>

## Next pages — bond_device_list.bond_devices.lacp / ccc4440fce93 / 5

- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-25016bab60e3edeac3f8dd443719bcfb7751b2e4ec4805b4228f9d388f8d0e0f)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-453e191bc6a950c24f51c59928005a09198b6c3033f1c9518951b091fc1341ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11632b69777dc860b3015e1bdb09c53917b9ec498409ed6ac378b73c36ad3d00"></a>

## coordinates — coordinates / 0849f57c0c76 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- coordinates

<a id="canonical-340afc7e237be13fe98c39020c412e09bb8321b9b3f61d0a8f037e69ca442b77"></a>

Type: `"single"`. Computed.

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

<a id="canonical-8e68652aa485a9b5eaa58cd204d8f41e20fe7a711bdb52e8104e5f3836df4f6b"></a>

## Direct properties — coordinates / 0849f57c0c76 / 3

<a id="canonical-fa9f8fab516d31b3cb5f6fbfb4b2c637ddf1bac4ea68d79ffbb755ce64cf4c44"></a>

<a id="canonical-4fc6f9eb30bd6582edd1829b548d995484bc4ab1f9c799cecc8b2ce6c34dc1d4"></a>

## latitude property — coordinates / 0849f57c0c76 / 4

Type: `"number"`. Computed.

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

<a id="canonical-350b72715ec464ff26517691afce57a561ac2a475ae8e94c8c4b9b2fbd6f061e"></a>

<a id="canonical-630af03888b6eaf6dcef0f940d595e83cc7f541101a3e11e2f652bb2bed0d6d6"></a>

## longitude property — coordinates / 0849f57c0c76 / 5

Type: `"number"`. Computed.

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

<a id="canonical-27783682a56dd6f94be25633c5f4f61d9af7aa79bdee5527fd0891265a33cec1"></a>

## Next pages — coordinates / 0849f57c0c76 / 6

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28535728b76829bbb1a6a75cc089bafca46ea68a872e9b95631a07f6843aa00e"></a>

## custom_network_config — custom_network_config / 790c971a63b3 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- custom_network_config

<a id="canonical-afe0919fdf577e1a8fab74825f2ee0a36ffe194811c77c4dc43c4fc6bbafb487"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_network\_config, default\_network\_config; Default: default\_network\_config\]
SmsNetworkConfiguration.

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

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-afe0919fdf577e1a8fab74825f2ee0a36ffe194811c77c4dc43c4fc6bbafb487)
- [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-9a932105e6427a1ebc735056852654b2935ede50f1f9ccba8dcd7479b2406715)

Select alternatives according to the provider validators above.

<a id="canonical-d3380d2a9b6fdafcce1d2cf6f35c77fa3b99f7d43ef4e0de2e4fb914559682e3"></a>

## Direct properties — custom_network_config / 790c971a63b3 / 3

- [active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-9994fb7aa3a2703ea2be6e671c66a2e0f51ed6882d02a985c884c20237afb38f): complete subsection reference.

- [active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-5f74ec54559dd08f664c50b53a9435990bc4672afa6ded964cd0ac85ab07204f): complete subsection reference.

- [active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-afcb1c4b4fdae34b1cdefd3b4c122e9e1c1a2f5ab4ced452cb02f12825b3d8f9): complete subsection reference.

- [default_config](data-sources--securemesh_site--reference--group-001.md#canonical-271154ccae718236ae9af1b31c2aee60a4d7d7b1ec719a60c31a1ffec400bc53): complete subsection reference.

- [default_interface_config](data-sources--securemesh_site--reference--group-001.md#canonical-29f4eaf55eca3676094b275ffa4577b3704afaea4013f12aed6a0a0409f01792): complete subsection reference.

- [default_sli_config](data-sources--securemesh_site--reference--group-001.md#canonical-c8a178dee13f7d11b3adcee5769934294fa9db12a5c0d2a3c4a1d984905c70ee): complete subsection reference.

- [forward_proxy_allow_all](data-sources--securemesh_site--reference--group-001.md#canonical-0952c6c502eb08fba919feab6c1e33ad35fdd2d2403c91bd52326023cfea86e8): complete subsection reference.

- [global_network_list](data-sources--securemesh_site--reference--group-002.md#canonical-6d6de7bdab673c3f80eb9575dce1261c417d12c228ff7772aea448c1467c23d3): complete subsection reference.

- [interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98): complete subsection reference.

- [no_forward_proxy](data-sources--securemesh_site--reference--group-003.md#canonical-a5790b2e3e396057bdf6d186f8e82a3b184aac39d74624bbbab99441acfe18d9): complete subsection reference.

- [no_global_network](data-sources--securemesh_site--reference--group-003.md#canonical-e3b6fd24f6cd0da5263135b5a51a86c63437e29faacb665a8ac4fb03dbc49a79): complete subsection reference.

- [no_network_policy](data-sources--securemesh_site--reference--group-003.md#canonical-25609b502097505129f9ca46cf6a92e8033f27ba9ed7e85f880578f8430ee0ba): complete subsection reference.

- [sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf): complete subsection reference.

- [slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1): complete subsection reference.

- [sm_connection_public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-e2329d10bedb8036b56fb105db5f5a3514befc9dfc0bbb405e9d5a1312b8c9f8): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--securemesh_site--reference--group-004.md#canonical-88e1809caf05a7030e5df09974c5e9b9a9a8f7a97bcf6e2537b80727f142cd55): complete subsection reference.

<a id="canonical-2ebc17566c6a966b0219b5567b5f99a6ac57baa82e1302533f267971561ad68d"></a>

<a id="canonical-00b38b643125bb0360bf5b59e8a01c7687b18fc1a7b67609f405844e68ca9c10"></a>

## tunnel_dead_timeout property — custom_network_config / 790c971a63b3 / 4

Type: `"number"`. Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Upstream description:

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

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

<a id="canonical-eb1e7318e93f27027561c7d42bca3872efae8ee21fe72d8a9e26f1be49f5a36d"></a>

<a id="canonical-245f8d9bf9f227cd9b027514c5b75f2cba12331371283914bae22d275a2d412a"></a>

## vip_vrrp_mode property — custom_network_config / 790c971a63b3 / 5

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

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

<a id="canonical-4b9d7002fe1d611bbd76468f588456ae5563c9cfa45524a2532e9410f6befd00"></a>

## Next pages — custom_network_config / 790c971a63b3 / 6

- [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-9994fb7aa3a2703ea2be6e671c66a2e0f51ed6882d02a985c884c20237afb38f)
- [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-5f74ec54559dd08f664c50b53a9435990bc4672afa6ded964cd0ac85ab07204f)
- [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-afcb1c4b4fdae34b1cdefd3b4c122e9e1c1a2f5ab4ced452cb02f12825b3d8f9)
- [custom_network_config.default_config](data-sources--securemesh_site--reference--group-001.md#canonical-271154ccae718236ae9af1b31c2aee60a4d7d7b1ec719a60c31a1ffec400bc53)
- [custom_network_config.default_interface_config](data-sources--securemesh_site--reference--group-001.md#canonical-29f4eaf55eca3676094b275ffa4577b3704afaea4013f12aed6a0a0409f01792)
- [custom_network_config.default_sli_config](data-sources--securemesh_site--reference--group-001.md#canonical-c8a178dee13f7d11b3adcee5769934294fa9db12a5c0d2a3c4a1d984905c70ee)
- [custom_network_config.forward_proxy_allow_all](data-sources--securemesh_site--reference--group-001.md#canonical-0952c6c502eb08fba919feab6c1e33ad35fdd2d2403c91bd52326023cfea86e8)
- [custom_network_config.global_network_list](data-sources--securemesh_site--reference--group-002.md#canonical-6d6de7bdab673c3f80eb9575dce1261c417d12c228ff7772aea448c1467c23d3)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-d3a6da542f10ec52fdcfec2fdc863fe7c0491e2d5ed371aee67f76060f85ab98)
- [custom_network_config.no_forward_proxy](data-sources--securemesh_site--reference--group-003.md#canonical-a5790b2e3e396057bdf6d186f8e82a3b184aac39d74624bbbab99441acfe18d9)
- [custom_network_config.no_global_network](data-sources--securemesh_site--reference--group-003.md#canonical-e3b6fd24f6cd0da5263135b5a51a86c63437e29faacb665a8ac4fb03dbc49a79)
- [custom_network_config.no_network_policy](data-sources--securemesh_site--reference--group-003.md#canonical-25609b502097505129f9ca46cf6a92e8033f27ba9ed7e85f880578f8430ee0ba)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-d3dc267cb490e71a083063b7ed8272724b62cc0f3ce759c11c9464d8f9ef04cf)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-87175fb3e4027c14d35f03c4c60d6d92c095d1896069ee3b2f818779af9989c1)
- [custom_network_config.sm_connection_public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-e2329d10bedb8036b56fb105db5f5a3514befc9dfc0bbb405e9d5a1312b8c9f8)
- [custom_network_config.sm_connection_pvt_ip](data-sources--securemesh_site--reference--group-004.md#canonical-88e1809caf05a7030e5df09974c5e9b9a9a8f7a97bcf6e2537b80727f142cd55)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-9994fb7aa3a2703ea2be6e671c66a2e0f51ed6882d02a985c884c20237afb38f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64d8ef6f6da4cfc154091f45bb3d5aad9115c0cd31b6cabebe398078da107a09"></a>

## custom_network_config.active_enhanced_firewall_policies — custom_network_config.active_enhanced_firewall_policies / 88644fab90ea / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.active_enhanced_firewall_policies

<a id="canonical-ada69fa7afbf9f9cfeec838715fe22e8f97dba4b77bc37e06a87c8a28039c98f"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

<a id="canonical-da4b8d75a30ab1b74c53c0d9dc6046c0aadf97f9b5826fc22557235920c369db"></a>

## Direct properties — custom_network_config.active_enhanced_firewall_policies / 88644fab90ea / 3

- [enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-d32a7495ad310263ce140057eba53b834bcc9d15877deed41ca1aa65fc227050): complete subsection reference.

<a id="canonical-30513a9ad81137c95f6cd7ef9fe8da0331978e520ff97eab48f86d4c245d8c1d"></a>

## Next pages — custom_network_config.active_enhanced_firewall_policies / 88644fab90ea / 4

- [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-d32a7495ad310263ce140057eba53b834bcc9d15877deed41ca1aa65fc227050)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-d32a7495ad310263ce140057eba53b834bcc9d15877deed41ca1aa65fc227050"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9f624fcb411c6d0b1131323d31947e386d6fd90de533c3e00d08e0ecbcc6c04"></a>

## custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / 98d8fb83f406 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-9994fb7aa3a2703ea2be6e671c66a2e0f51ed6882d02a985c884c20237afb38f)
- custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-20b40fd83ad4334e9a9461b3a6a0dff03fe295a0479628f580c3c23041411d79"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

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

<a id="canonical-f777a914d5eaaff923353fd13e880c238b86c478ebdc2a52214b6b0a9cd00fea"></a>

## Direct properties — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / 98d8fb83f406 / 3

<a id="canonical-7426ece09b3d07e966827a462d9cf5fdbeae9f588b5d549ed1def9cb74954a2d"></a>

<a id="canonical-3a9d1d6cea483b7df5e1e24cecc6cc9f741c6af0824c04ad23ec124e81b39d09"></a>

## name property — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / 98d8fb83f406 / 4

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

<a id="canonical-a59266089310cb67d3e7cede3eecb4688ab21d8a9db02610f3b5778f2b1d80ac"></a>

<a id="canonical-8fd61ca71b537b293a61a23af7459d57bcccba33b5ac498a4205d6076faf88ff"></a>

## namespace property — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / 98d8fb83f406 / 5

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

<a id="canonical-d06fe46401086b57420e86aca22323ca67efe520709e0c5e09affce1b8bc5c44"></a>

<a id="canonical-f651f8ff554dbcf16448902b4c207632aa4bb31e8fb43ba1092efea4251c6016"></a>

## tenant property — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / 98d8fb83f406 / 6

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

<a id="canonical-feeab68bdc4e65b36f7c6a378a165d20475d95da1e95b56f2f564a667ef0e93b"></a>

## Next pages — custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_polici / 98d8fb83f406 / 7

- [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-9994fb7aa3a2703ea2be6e671c66a2e0f51ed6882d02a985c884c20237afb38f)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-5f74ec54559dd08f664c50b53a9435990bc4672afa6ded964cd0ac85ab07204f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24c552d93e1e62bfaabb8a4ae1c28af1a3f9ee96917700b43dcf0d050e0b1c3e"></a>

## custom_network_config.active_forward_proxy_policies — custom_network_config.active_forward_proxy_policies / 452e962a7f6e / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.active_forward_proxy_policies

<a id="canonical-6500e4c8f249b4fb8f1073ab76052246c94815be714c819e54533e37c69f395c"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

<a id="canonical-c80534549a1629d1749cd3668270540346372e7a6ba136bc8f554a00e9982306"></a>

## Direct properties — custom_network_config.active_forward_proxy_policies / 452e962a7f6e / 3

- [forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-ba4853fcf56bab9ad493b088914b6790a0e24846bc0b6e133a80adbf84bce55d): complete subsection reference.

<a id="canonical-eb55f04b0594cd05165b847d07c8be752fe62bb62d8773654d5e5649a7fdae64"></a>

## Next pages — custom_network_config.active_forward_proxy_policies / 452e962a7f6e / 4

- [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-ba4853fcf56bab9ad493b088914b6790a0e24846bc0b6e133a80adbf84bce55d)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-ba4853fcf56bab9ad493b088914b6790a0e24846bc0b6e133a80adbf84bce55d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-729e9627150850217bb84266755e872f3d8c1124c8ae027dea703c5516e0af1b"></a>

## custom_network_config.active_forward_proxy_policies.forward_proxy_policies — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / 38e1bf45137b / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-5f74ec54559dd08f664c50b53a9435990bc4672afa6ded964cd0ac85ab07204f)
- custom_network_config.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-721298d4ca657febcda6afbfe387949f736ffc9b3b0b1f096be45fbcb3b088bf"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

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

<a id="canonical-07d5944d372d3b8f2c842c6d74ac12e634e2111da305cb932dcd1fc76fd56dde"></a>

## Direct properties — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / 38e1bf45137b / 3

<a id="canonical-c52fe3c7bf62535454ac8da58475cfd2564c763d6f300d8690918e435ccfc7ed"></a>

<a id="canonical-16a9c37b5e9014e657193e4962604042c0c152253f1376f67e6f78cb4a48d5a4"></a>

## name property — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / 38e1bf45137b / 4

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

<a id="canonical-c604796d12da1760bd61dbd72f518d89c323160fdf890432546db7663da75334"></a>

<a id="canonical-4d8604240c02bd4628475167c1f9206c9f014c83e323103a8de46406d200fadb"></a>

## namespace property — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / 38e1bf45137b / 5

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

<a id="canonical-72f49247a19713bba1b5ae5c1b95ebc029a8802e078ef7bb4cfd0ead94496a5a"></a>

<a id="canonical-51049ac5b524cfea17cc95710e53552eed0e21f1108f07125a0fc534078b07fc"></a>

## tenant property — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / 38e1bf45137b / 6

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

<a id="canonical-7a7a802282f0ceceaf0ff82517e0b7c3080e37e8492084d705ac5f05e9215333"></a>

## Next pages — custom_network_config.active_forward_proxy_policies.forward_proxy_policies / 38e1bf45137b / 7

- [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-5f74ec54559dd08f664c50b53a9435990bc4672afa6ded964cd0ac85ab07204f)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-afcb1c4b4fdae34b1cdefd3b4c122e9e1c1a2f5ab4ced452cb02f12825b3d8f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd51dffbd38b228ffa6c8a54964376115bb85748e32d23239ca002d3cf7272d9"></a>

## custom_network_config.active_network_policies — custom_network_config.active_network_policies / cb6156d0dcb0 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.active_network_policies

<a id="canonical-b474513c0c7b58d36cba6153ea656f9ae006f3d83c331e6ac49c94b2685941fe"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

<a id="canonical-94e22b23717777e4238d23b8e49096b6e0454764258d1f74be3ce3b59ff13474"></a>

## Direct properties — custom_network_config.active_network_policies / cb6156d0dcb0 / 3

- [network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-dee69532d36c51874bfb0f15af45d3c2baab0023545a75a0aa9c33be897bf6f7): complete subsection reference.

<a id="canonical-f86507117955c5db77c5afe075481c4a74be73a12c287364d5d8b397ff5ae2ce"></a>

## Next pages — custom_network_config.active_network_policies / cb6156d0dcb0 / 4

- [custom_network_config.active_network_policies.network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-dee69532d36c51874bfb0f15af45d3c2baab0023545a75a0aa9c33be897bf6f7)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-dee69532d36c51874bfb0f15af45d3c2baab0023545a75a0aa9c33be897bf6f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7497e5f5f62d0bbc5a06991f3b229d67056cdf3844241b05529c76c9d4215ed4"></a>

## custom_network_config.active_network_policies.network_policies — custom_network_config.active_network_policies.network_policies / 09f16877ca35 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-afcb1c4b4fdae34b1cdefd3b4c122e9e1c1a2f5ab4ced452cb02f12825b3d8f9)
- custom_network_config.active_network_policies.network_policies

<a id="canonical-f46099ba36f76a193eaa98c92aff412e64dd3b2ea8bb7a563298035217ad19c5"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

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

<a id="canonical-70cca8fa1358a60297c9026115414cae5e742a709f73c2b14bf66e4686d3cb85"></a>

## Direct properties — custom_network_config.active_network_policies.network_policies / 09f16877ca35 / 3

<a id="canonical-9de80cca2431833d5dc67e9b104fd12740e122098d3a902b11148d2664ce8fc3"></a>

<a id="canonical-4bbff86e5c0076c37c29e44dea10b0130a80c49a1fb15e3ea19a9c8411eb2bab"></a>

## name property — custom_network_config.active_network_policies.network_policies / 09f16877ca35 / 4

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

<a id="canonical-0b52705b6aff880183316cca3bd7930bc04db51b9996d1a498338114b59f73ce"></a>

<a id="canonical-16172e81923bfbff5ff132eea6748036df16ab280e253ae32d540c9813964200"></a>

## namespace property — custom_network_config.active_network_policies.network_policies / 09f16877ca35 / 5

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

<a id="canonical-b2894e7011881786def20c0a363484b9e57a5a66d0ce65fd0a8320eceb58e65a"></a>

<a id="canonical-6a69acc82cc51bae2f89d27f9165a830b00459e01af2d53967bd2fe17dba93ab"></a>

## tenant property — custom_network_config.active_network_policies.network_policies / 09f16877ca35 / 6

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

<a id="canonical-ea9a17d85045350161ea1d2fbf2d2534e2eadfa0f424c5c6d5005d2b86ab4c3d"></a>

## Next pages — custom_network_config.active_network_policies.network_policies / 09f16877ca35 / 7

- [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-afcb1c4b4fdae34b1cdefd3b4c122e9e1c1a2f5ab4ced452cb02f12825b3d8f9)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-271154ccae718236ae9af1b31c2aee60a4d7d7b1ec719a60c31a1ffec400bc53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19d706f6f13870e6edbfe3b6279925834ba3578032b8b43a16c80ea7da3c43a1"></a>

## custom_network_config.default_config — custom_network_config.default_config / 1b496ff3a982 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.default_config

<a id="canonical-3fd5156c0967948fa466d2d4a1f3c09130a52eee88d50ed4b1174ad382bfd9f0"></a>

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

<a id="canonical-d8420ec688459ef015dc9b302448fe7da5c2ef9ef73f33ff1744466f4794782c"></a>

## Direct properties — custom_network_config.default_config / 1b496ff3a982 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d16d3d1b248151b53f2e041c5ae14a521df2f0a7dbafd68ea7690f5176cefe5c"></a>

## Next pages — custom_network_config.default_config / 1b496ff3a982 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-29f4eaf55eca3676094b275ffa4577b3704afaea4013f12aed6a0a0409f01792"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a459db744d71d14a36ac9e61b145e5f6d13f05508eb320f40f92900fbb132b3"></a>

## custom_network_config.default_interface_config — custom_network_config.default_interface_config / 953f3060c4bb / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.default_interface_config

<a id="canonical-aec9251db0470f5c54be1203d30dbd7f6ac74f9cb08b52b569f2e521509bdfd4"></a>

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

<a id="canonical-58131a51758792d27e9ed03b1b923ae4b1d9147dffccc431150fd89ea3ef468f"></a>

## Direct properties — custom_network_config.default_interface_config / 953f3060c4bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f5219d3184f7ed7ced7d65cfa05569c8872600da09eab268a70235f10b6f8b5a"></a>

## Next pages — custom_network_config.default_interface_config / 953f3060c4bb / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-c8a178dee13f7d11b3adcee5769934294fa9db12a5c0d2a3c4a1d984905c70ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a5f6bd232e218a70fe0cef7d457c4268d4c605a445852ea6a52c77b4fc4bc23"></a>

## custom_network_config.default_sli_config — custom_network_config.default_sli_config / f0744019a595 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- custom_network_config.default_sli_config

<a id="canonical-7fb883b87221c63872193949300ae0f4d6565f1211148cdda2d6fec4c398c5bf"></a>

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

<a id="canonical-4cc71e19a8404eeb360f8af2824eed0883bc06bc4bc84897209b54a9369f8f79"></a>

## Direct properties — custom_network_config.default_sli_config / f0744019a595 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f79f7ee9e7ef32982c368a8e5c6ad158494027c256c18e8234b8761c48eeee8a"></a>

## Next pages — custom_network_config.default_sli_config / f0744019a595 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-157f9157aeb3dfae72a7bc8e1a027b49d34278b21f272faa654cf8cabcb9b972)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-0952c6c502eb08fba919feab6c1e33ad35fdd2d2403c91bd52326023cfea86e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
