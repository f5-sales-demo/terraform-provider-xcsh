---
page_title: "xcsh_enhanced_firewall_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy reference."
---

# xcsh_enhanced_firewall_policy reference

<a id="canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7ef2d82e5dcfed20d1288b602045bb5323c8e02337f411272d2ac10627ba31a"></a>

## Property reference — Property reference / 74125bcbd8b9 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- Property reference

<a id="canonical-0a903f15a5a375cab431a308bdcb3c80ffcd14d6a28834c5b2005b26ad3b4ca8"></a>

## Direct properties — Property reference / 74125bcbd8b9 / 3

- [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-5276b4ed6998fd5ea6e1f5109bbe832b3805bbec286bef8a4d15eb9680715106): complete subsection reference.

- [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-48f0d33799af6b101ccc88cd26961c20f7bf92e78d0a66e580c9da909f017aef): complete subsection reference.

- [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2f4458e29d46b0b4734a51bb236c769f346e73e3aceade0dcade5f04c2b388ed): complete subsection reference.

<a id="canonical-b367b45634aa09e6395f4cd9b0fad1f944afbb68930a8396bb4dfe095ea224f4"></a>

<a id="canonical-6a30ff7c33e3c7aac2966d6e01e5f173a976b8b6bc5230a128610a3a896dfbf7"></a>

## annotations property — Property reference / 74125bcbd8b9 / 4

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

- [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-bcca1d680f02b6fa02f6f34a433cf602ade748faf3185cf50fef774b41767f2e): complete subsection reference.

- [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-cce21b83f7898464681e349c686b5fc293412893c52c5164fa3168b500801426): complete subsection reference.

- [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b5b3a88eddb7e5e1af7a0b8e9a401a6b018ccc4a11ed7a0dc68b0d8e66a61f0c): complete subsection reference.

<a id="canonical-a504ebcfea08e1ae1db0af68deff9ea440beca822e12e73218daf861fcd49206"></a>

<a id="canonical-4c1305f10fcbbb519096e3eb1656203d723f5e13ab847ce022af572059738475"></a>

## description property — Property reference / 74125bcbd8b9 / 5

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

<a id="canonical-5e39a22574abd5e921437b8117f26c7b6979078e7d687ddeac09383fd9e91bf0"></a>

<a id="canonical-5be811ed2632edfe551243e7a223cb438cfc3ea49d70b9d1dc9ecdf5ff6ac635"></a>

## disable property — Property reference / 74125bcbd8b9 / 6

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

<a id="canonical-853788457a89e94ab6e721099c01b7438ff9f49589d934495dcbaa5572ec121e"></a>

<a id="canonical-9e73bea752ace7631871b12fc4f8bcc77aaafedd965cdf674311e6539b5b0117"></a>

## id property — Property reference / 74125bcbd8b9 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-62c0546161669d9f2319613de60a5bdd40cbf2a33a539501082bfc6360a7e3f1"></a>

<a id="canonical-373d273904e78c1757bece2ded7dfd13b9cc72fe8701b5378a5df877fd534477"></a>

## labels property — Property reference / 74125bcbd8b9 / 8

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

<a id="canonical-6575073b0b9efd4ffe6e79a5836000cc67bbe5b7ff7a390aa369b409a9c02ffc"></a>

<a id="canonical-53cc8e04be04019418657b076fcdecc07d983371c6536a93527bfc7ad75ed541"></a>

## name property — Property reference / 74125bcbd8b9 / 9

Type: `"string"`. Required.

Name of the Enhanced Firewall Policy. Must be unique within the namespace.

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

<a id="canonical-479b749f5d858dd412a40c64e78cbaaf70f71d6b7458b4d93094592b17919fa8"></a>

<a id="canonical-8c397581c011430bbb3d00acb1759124bf8164b7699bb4b03472c91a4d69ce60"></a>

## namespace property — Property reference / 74125bcbd8b9 / 10

Type: `"string"`. Required.

Namespace where the Enhanced Firewall Policy is created.

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

- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf): complete subsection reference.

- [timeouts](resources--enhanced_firewall_policy--reference--group-001.md#canonical-507b69e00e66568c208f5606958c9ee2fb8c1eed5b5c2940ca5eb5b7549f3400): complete subsection reference.

<a id="canonical-ff737f345b70bed948d35279ef8048c5ed79aa59c73feaddef5ffe3ca9c83559"></a>

## All schema paths — Property reference / 74125bcbd8b9 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-9faf31af7a963e12ff7c34a18edd805e416f4e6d504b3e69695f2745ee285076) |
| `allowed_destinations` | [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-639af01b66a943cc0c70087d18c371e78b980cb19b097c015ef5c3f6cd89939c) |
| `allowed_destinations.prefix` | [allowed_destinations.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-81ed56df4de829c94d7841f0ec22f2c06318a1210ed1fd163257f2e761478e18) |
| `allowed_sources` | [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d4ffc522fb73a6622fca90d51ddbc44e8d79a62740b54c1d9e4cb36995db98b5) |
| `allowed_sources.prefix` | [allowed_sources.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-cc45abb9e97b63c32df3da679481a4724030a2048bfa88113d531e9747874956) |
| `annotations` | [annotations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b367b45634aa09e6395f4cd9b0fad1f944afbb68930a8396bb4dfe095ea224f4) |
| `denied_destinations` | [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-25ff2dfe97aa0cbb8b5aa73de8e20c019b8e32a3583d1813339405266758694a) |
| `denied_destinations.prefix` | [denied_destinations.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-040740875f57bb0080ef074706da67ab2a839fc7fbf793593cfbefdf1c0cd45b) |
| `denied_sources` | [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-9114780aa936a69ec7524513b52ee583bd246a71caf023b4c67fdc5f5a8f5f49) |
| `denied_sources.prefix` | [denied_sources.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-441cbe84386d242252abcc49ae2baa1f6108dc7d2108297d42aed77d9dba587c) |
| `deny_all` | [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-805eced645ffcf73c665dcc28322d2e8ac4c6bcc3355df92a6ad00ea910cff67) |
| `description` | [description](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a504ebcfea08e1ae1db0af68deff9ea440beca822e12e73218daf861fcd49206) |
| `disable` | [disable](resources--enhanced_firewall_policy--reference--group-001.md#canonical-5e39a22574abd5e921437b8117f26c7b6979078e7d687ddeac09383fd9e91bf0) |
| `id` | [id](resources--enhanced_firewall_policy--reference--group-001.md#canonical-853788457a89e94ab6e721099c01b7438ff9f49589d934495dcbaa5572ec121e) |
| `labels` | [labels](resources--enhanced_firewall_policy--reference--group-001.md#canonical-62c0546161669d9f2319613de60a5bdd40cbf2a33a539501082bfc6360a7e3f1) |
| `name` | [name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-6575073b0b9efd4ffe6e79a5836000cc67bbe5b7ff7a390aa369b409a9c02ffc) |
| `namespace` | [namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-479b749f5d858dd412a40c64e78cbaaf70f71d6b7458b4d93094592b17919fa8) |
| `rule_list` | [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a2173d314aa1b1f4950102ace47c14aa6cc1566fd49bf837a27ab149f9a50355) |
| `rule_list.rules` | [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-f1429792d32130de3aebf6d92b0d57f39ccf88fa2a28002428d9f99e2d73f24b) |
| `rule_list.rules.advanced_action` | [rule_list.rules.advanced_action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-e15a8d03863e0f7d3011f828a73aa3e3446c152bf17666f1e989b37acb410d3e) |
| `rule_list.rules.advanced_action.action` | [rule_list.rules.advanced_action.action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-69ecddbad856b6f68823de187992afbc00b2e6dc1fc6cecd097388c8b123ed9d) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-9b0e97754e6a04ef6d0dd806359b354589fe0554615afa5385b63b5e2bcf8898) |
| `rule_list.rules.all_sli_vips` | [rule_list.rules.all_sli_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-9f6de4b09d03134b5e161ed2e0daedfdfc1307e007df3a7bc88475bcd8769226) |
| `rule_list.rules.all_slo_vips` | [rule_list.rules.all_slo_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-5b9cc39d2a21f0f0a2a1e7fb6404ed876d12568c55c7b6f9d81d5ecea8e45e3d) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b414e6d4c388ece624be15a8da857fa217add4e83b99cecca973dc99ae0dc5d0) |
| `rule_list.rules.all_tcp_traffic` | [rule_list.rules.all_tcp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-61964d4e8778f040e6657a53df4848901e823d0068a3104395e1f4ad3eb7f965) |
| `rule_list.rules.all_traffic` | [rule_list.rules.all_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3e3256e04a30f4650def0e2e368febef5293a755aa6faf05619a6a5cae1c794b) |
| `rule_list.rules.all_udp_traffic` | [rule_list.rules.all_udp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-92b3a1ed390e1b311e5cc66a74069086311504e517bf538a83d76337eeb4e440) |
| `rule_list.rules.allow` | [rule_list.rules.allow](resources--enhanced_firewall_policy--reference--group-001.md#canonical-650d2c009a8d4374e2ecfba078c224d218b1546825558f3a8bd2a5d080f249b8) |
| `rule_list.rules.applications` | [rule_list.rules.applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-15f202ca085a25e43e2a80d9be2a039582ddf302a09f882853fb6112feafa25e) |
| `rule_list.rules.applications.applications` | [rule_list.rules.applications.applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3524aa71a0884ba161171ae7f6bc5f0511e287ecb6dbb724d1453d7cfbb4d625) |
| `rule_list.rules.deny` | [rule_list.rules.deny](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3a0e3f106ddcec9451ceb07a188ac42e1905da0b607cacac4c4c11ad5e45eb61) |
| `rule_list.rules.destination_aws_vpc_ids` | [rule_list.rules.destination_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-27ace91d66e071e0371ae741f7a7cf9763bc4c749dd163cee6a0495c809fb7a0) |
| `rule_list.rules.destination_aws_vpc_ids.vpc_id` | [rule_list.rules.destination_aws_vpc_ids.vpc_id](resources--enhanced_firewall_policy--reference--group-001.md#canonical-25353afa2f55215632c9a32330ffee115492292ec74a5b6df5ad70b99499e1ae) |
| `rule_list.rules.destination_ip_prefix_set` | [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2d0421ea370b3327b76b263abf00b0e2309cdac69f40b806c9334667f8c4e7b9) |
| `rule_list.rules.destination_ip_prefix_set.ref` | [rule_list.rules.destination_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ef95a667ce70817641e2a89fbbac8c6b4537df1abcf604782fc4c146aafda680) |
| `rule_list.rules.destination_ip_prefix_set.ref.kind` | [rule_list.rules.destination_ip_prefix_set.ref.kind](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8c91abd2c13c19b9389426ff361aab7d261632618e4be242efa82eec3e930a64) |
| `rule_list.rules.destination_ip_prefix_set.ref.name` | [rule_list.rules.destination_ip_prefix_set.ref.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ef2b97d2eac0edaba0a454487d895ffa9e7ff8632ac11170a49ae304426c127a) |
| `rule_list.rules.destination_ip_prefix_set.ref.namespace` | [rule_list.rules.destination_ip_prefix_set.ref.namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0a90da958902b192d054247c3376b3f2936a8b8e4da25c22416f27e04e066b6b) |
| `rule_list.rules.destination_ip_prefix_set.ref.tenant` | [rule_list.rules.destination_ip_prefix_set.ref.tenant](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1e86056f714e8138f6521e06f30fb2b0c036d3e9d7c616572b7b4accc50306f0) |
| `rule_list.rules.destination_ip_prefix_set.ref.uid` | [rule_list.rules.destination_ip_prefix_set.ref.uid](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c1884b2b4db474ab89f7af4416bec0d74ee04e9683c196c25336f48b1111045f) |
| `rule_list.rules.destination_label_selector` | [rule_list.rules.destination_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-400e7043e67357286c01c0151a3c8bd107775d7b2f98fe19ee8d7ad74549a274) |
| `rule_list.rules.destination_label_selector.expressions` | [rule_list.rules.destination_label_selector.expressions](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ae86c865ec2512011411eb2814bad24ae8563f1f9a17cce013eaf8a22c8e764c) |
| `rule_list.rules.destination_prefix_list` | [rule_list.rules.destination_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-12453a2bdbaa8a11b02c289383236078706d8ece05862d089eefa7be20c2b88f) |
| `rule_list.rules.destination_prefix_list.prefixes` | [rule_list.rules.destination_prefix_list.prefixes](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2f9e2406791e89390749848e98da1898d58b815ad2cc5a2b14a3bcd03e4d6d9b) |
| `rule_list.rules.insert_service` | [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3cb99e68ac3efdbfcc2616cea1a79ec38940a5431c314f62e29355253918336e) |
| `rule_list.rules.insert_service.nfv_service` | [rule_list.rules.insert_service.nfv_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-691e517290e7b0efa5d615acd90b388bbdff7ff4ba130b3a60bd818bf34d1faf) |
| `rule_list.rules.insert_service.nfv_service.name` | [rule_list.rules.insert_service.nfv_service.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-5e16ad124ef4bb72c89fe250fe7563264b4ff16a2a3419910cb42d177267899d) |
| `rule_list.rules.insert_service.nfv_service.namespace` | [rule_list.rules.insert_service.nfv_service.namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ae46c2ed8ef5b916136df5226a6a3e349d3d3ad3c20410c8effa0c328281c9af) |
| `rule_list.rules.insert_service.nfv_service.tenant` | [rule_list.rules.insert_service.nfv_service.tenant](resources--enhanced_firewall_policy--reference--group-001.md#canonical-cead5103a841d878f9b9780db3f99e9ae4225ad33b48bfc4445e7f40f0a2ae5d) |
| `rule_list.rules.inside_destinations` | [rule_list.rules.inside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-5798d14d3b926b08bb93784235c7cf33148191678897981b2f773720593ea7c5) |
| `rule_list.rules.inside_sources` | [rule_list.rules.inside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-cfc814ea2b12b916687aad60bf5be35f5b688904c94d733cb9cd4d8962d231e3) |
| `rule_list.rules.label_matcher` | [rule_list.rules.label_matcher](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a2f1c78ed06416d050edade323eab72a4be0838b9bc08d7ab1230b3922659f3a) |
| `rule_list.rules.label_matcher.keys` | [rule_list.rules.label_matcher.keys](resources--enhanced_firewall_policy--reference--group-001.md#canonical-93a79c704876d20ab4201c39c03b2bf270cb333cb7950386bc5c34df6e941517) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--enhanced_firewall_policy--reference--group-001.md#canonical-6ac2f0cc6db8aacdeeed2cc6b5b37b47ebb1689157ee3c97b487468f2df94eca) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--enhanced_firewall_policy--reference--group-001.md#canonical-9438bda3b1ca25f15aa9831e46325bfb175a91096c0540e15f6324884839de9a) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-89049dfdce8fc4646f5ff88b29ec5021b64ab432a8ce254425199c97b29baf0e) |
| `rule_list.rules.outside_destinations` | [rule_list.rules.outside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c1655578dcff8913053b1666b67c03340bff5ead56f383ddf0fde633245b2292) |
| `rule_list.rules.outside_sources` | [rule_list.rules.outside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b4545d22c4660b29c684a14e10c6e1fa96cf76809ef61e65d92ae658fed92d60) |
| `rule_list.rules.protocol_port_range` | [rule_list.rules.protocol_port_range](resources--enhanced_firewall_policy--reference--group-001.md#canonical-27943e54491b481302c46c379b8bd092ea5a072f4474bd3fa33bef20f6a597d4) |
| `rule_list.rules.protocol_port_range.port_ranges` | [rule_list.rules.protocol_port_range.port_ranges](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ed8b0a0be79e76a1d37f3781ed41b96594fbda5babb7e08f055a4ec7a1d4768e) |
| `rule_list.rules.protocol_port_range.protocol` | [rule_list.rules.protocol_port_range.protocol](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d9cf0de8df6bef566e9ff47ad4e6fe2bd0862a9bd56b4a5a9eb2848c733f1955) |
| `rule_list.rules.source_aws_vpc_ids` | [rule_list.rules.source_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-f211ce5b827cadb9d42eeefe07463bef893256c137fe89cf28897b1f094c6c1a) |
| `rule_list.rules.source_aws_vpc_ids.vpc_id` | [rule_list.rules.source_aws_vpc_ids.vpc_id](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ef786556bccd94472fea928b293c15cc86457629c00435713e99216b00ed2e78) |
| `rule_list.rules.source_ip_prefix_set` | [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-07f0ed96a8bf58cef01ae1fec6efbfd7fe2eb85d5ddc35ac679e437ade8029d7) |
| `rule_list.rules.source_ip_prefix_set.ref` | [rule_list.rules.source_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ce287f79f7675b50a93c454623a6b002c06ca0abbbc17e21ed39c1066cdb448d) |
| `rule_list.rules.source_ip_prefix_set.ref.kind` | [rule_list.rules.source_ip_prefix_set.ref.kind](resources--enhanced_firewall_policy--reference--group-001.md#canonical-f43c3c157f8caf6c025209c67dda2aefdaa4499f7d96ad3acd9ef43de3705031) |
| `rule_list.rules.source_ip_prefix_set.ref.name` | [rule_list.rules.source_ip_prefix_set.ref.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-342966d73a299729b3c9f67f57251e073f1dcc5af914cfc03403993a2aebdc22) |
| `rule_list.rules.source_ip_prefix_set.ref.namespace` | [rule_list.rules.source_ip_prefix_set.ref.namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d30c25f3165f0b4239d55a919b408179720de3a249b626f7a37e31299e2cd79a) |
| `rule_list.rules.source_ip_prefix_set.ref.tenant` | [rule_list.rules.source_ip_prefix_set.ref.tenant](resources--enhanced_firewall_policy--reference--group-001.md#canonical-afbb854760e44da5ad3edc949080a0debec533770059e920dcdf2d381b709dea) |
| `rule_list.rules.source_ip_prefix_set.ref.uid` | [rule_list.rules.source_ip_prefix_set.ref.uid](resources--enhanced_firewall_policy--reference--group-001.md#canonical-396a07c5304d78c05492c5cd312775f3dfb20b9aa95a6a4f3cbae5047d2b2162) |
| `rule_list.rules.source_label_selector` | [rule_list.rules.source_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-4541f9ca68cc7f2db46c31038c319e61f59cef0ccaaf5d9887ac2ef5fd677735) |
| `rule_list.rules.source_label_selector.expressions` | [rule_list.rules.source_label_selector.expressions](resources--enhanced_firewall_policy--reference--group-001.md#canonical-350fe43b6babc908337405ac68d8ad6a869a44616e52f399d90d385417db55ae) |
| `rule_list.rules.source_prefix_list` | [rule_list.rules.source_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-daab1d2cbe6cb6c0cd1ab8da89f3364ff9e3adb719b4416bc8f6f95899145395) |
| `rule_list.rules.source_prefix_list.prefixes` | [rule_list.rules.source_prefix_list.prefixes](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8f5c50164a133dd5dd5f7cd1eeba01b67aad5b61f9539cb7980d4c7b66975636) |
| `timeouts` | [timeouts](resources--enhanced_firewall_policy--reference--group-001.md#canonical-725b7fc87126b784415a555ec1a8eaf4ecd5a221d6180b09dc1643e973c3a315) |
| `timeouts.create` | [timeouts.create](resources--enhanced_firewall_policy--reference--group-001.md#canonical-21c58c0f83589e0c0c298f49d04d34cf850cc6fada3d011a4cc82aafd20b5dc5) |
| `timeouts.delete` | [timeouts.delete](resources--enhanced_firewall_policy--reference--group-001.md#canonical-39af0e9e57ee9f48d5b01aac07f0ffa4a69dab9fc7dcaeaef4107af708969686) |
| `timeouts.read` | [timeouts.read](resources--enhanced_firewall_policy--reference--group-001.md#canonical-e68e532451226b72249aff8ae6c66fcd2ca071267c846f563108bc15dbd928c1) |
| `timeouts.update` | [timeouts.update](resources--enhanced_firewall_policy--reference--group-001.md#canonical-ab88d4accdd5cc66e5e3375d18acca1986c84e503b52dbf4fa360b42eafa545c) |

<a id="canonical-ebe165aca6cc9ce80e0fc49c5e5ce76cab7691cbf6d3d90c00c02b633d02d50f"></a>

## Next pages — Property reference / 74125bcbd8b9 / 12

- [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-5276b4ed6998fd5ea6e1f5109bbe832b3805bbec286bef8a4d15eb9680715106)
- [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-48f0d33799af6b101ccc88cd26961c20f7bf92e78d0a66e580c9da909f017aef)
- [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2f4458e29d46b0b4734a51bb236c769f346e73e3aceade0dcade5f04c2b388ed)
- [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-bcca1d680f02b6fa02f6f34a433cf602ade748faf3185cf50fef774b41767f2e)
- [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-cce21b83f7898464681e349c686b5fc293412893c52c5164fa3168b500801426)
- [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b5b3a88eddb7e5e1af7a0b8e9a401a6b018ccc4a11ed7a0dc68b0d8e66a61f0c)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [timeouts](resources--enhanced_firewall_policy--reference--group-001.md#canonical-507b69e00e66568c208f5606958c9ee2fb8c1eed5b5c2940ca5eb5b7549f3400)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-5276b4ed6998fd5ea6e1f5109bbe832b3805bbec286bef8a4d15eb9680715106"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-998aeeb1cf5127fa1fb05ec109adefde6260b2fb99d2812bbe2bad148bd321b4"></a>

## allow_all — allow_all / c805bddd3fc3 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- allow_all

<a id="canonical-9faf31af7a963e12ff7c34a18edd805e416f4e6d504b3e69695f2745ee285076"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: allow\_all, allowed\_destinations, allowed\_sources, denied\_destinations, denied\_sources,
deny\_all, rule\_list\] Enable this option. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

- [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-9faf31af7a963e12ff7c34a18edd805e416f4e6d504b3e69695f2745ee285076)
- [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-639af01b66a943cc0c70087d18c371e78b980cb19b097c015ef5c3f6cd89939c)
- [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d4ffc522fb73a6622fca90d51ddbc44e8d79a62740b54c1d9e4cb36995db98b5)
- [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-25ff2dfe97aa0cbb8b5aa73de8e20c019b8e32a3583d1813339405266758694a)
- [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-9114780aa936a69ec7524513b52ee583bd246a71caf023b4c67fdc5f5a8f5f49)
- [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-805eced645ffcf73c665dcc28322d2e8ac4c6bcc3355df92a6ad00ea910cff67)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a2173d314aa1b1f4950102ace47c14aa6cc1566fd49bf837a27ab149f9a50355)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all = {}
```

<a id="canonical-710da7dde2eef3d03774748a61ad7b7067ac64b2815ecd8b23329377bbabeb5a"></a>

## Direct properties — allow_all / c805bddd3fc3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1c84092f3923ba767430fa20ddea7a906f2a15b8910fa76e051c5909c572cf0"></a>

## Next pages — allow_all / c805bddd3fc3 / 4

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-48f0d33799af6b101ccc88cd26961c20f7bf92e78d0a66e580c9da909f017aef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0d0a61a36ac3f24afb22766d343d8df2bfed1f0491be291c3cce7a7568c71b5"></a>

## allowed_destinations — allowed_destinations / 9d13e285e3e9 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- allowed_destinations

<a id="canonical-639af01b66a943cc0c70087d18c371e78b980cb19b097c015ef5c3f6cd89939c"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
allowed_destinations {
  # Configure direct properties listed below.
}
```

<a id="canonical-782a3aa2ff39b3c41c9b4b677cce36ea06abe7f2bdbb150ddbba16d3af3a07eb"></a>

## Direct properties — allowed_destinations / 9d13e285e3e9 / 3

<a id="canonical-81ed56df4de829c94d7841f0ec22f2c06318a1210ed1fd163257f2e761478e18"></a>

<a id="canonical-4ad8ae7b0ab29cd73fcbdf033bcb13addf48cc27bf826105a8f15216574ae1ac"></a>

## prefix property — allowed_destinations / 9d13e285e3e9 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3b56ba102f027e0ebca384413a623dadcc277790be32a02c7e84e685ec7af3a6"></a>

## Next pages — allowed_destinations / 9d13e285e3e9 / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-2f4458e29d46b0b4734a51bb236c769f346e73e3aceade0dcade5f04c2b388ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09b422752ea7a532deda917c6ff5d2f482fcc4ac5b4c347a315a63234ed9e5d8"></a>

## allowed_sources — allowed_sources / 4f5c5415328f / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- allowed_sources

<a id="canonical-d4ffc522fb73a6622fca90d51ddbc44e8d79a62740b54c1d9e4cb36995db98b5"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
allowed_sources {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c4bce1048338178599a2913282f779cb640360f34d39ab979bb78db293d8f7e"></a>

## Direct properties — allowed_sources / 4f5c5415328f / 3

<a id="canonical-cc45abb9e97b63c32df3da679481a4724030a2048bfa88113d531e9747874956"></a>

<a id="canonical-8fbca1e88c4e69bae72d9894a9e989c3c643efc35a2cd475f57653d013ec5ae1"></a>

## prefix property — allowed_sources / 4f5c5415328f / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3c9281c8663870776a013feeba9d073d640572e70fe035cbc43df6d4bc62c5f5"></a>

## Next pages — allowed_sources / 4f5c5415328f / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-bcca1d680f02b6fa02f6f34a433cf602ade748faf3185cf50fef774b41767f2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ea179708c271e3c2fb5aad647103f78b0fc025b16b152e0bb89aa6e24fb11ed"></a>

## denied_destinations — denied_destinations / d44ebdcf0e2a / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- denied_destinations

<a id="canonical-25ff2dfe97aa0cbb8b5aa73de8e20c019b8e32a3583d1813339405266758694a"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
denied_destinations {
  # Configure direct properties listed below.
}
```

<a id="canonical-124445b73a6406dc7a637b49ca6a5f71a1707ddc96b2ab0befd3e1ec4c3b8c2f"></a>

## Direct properties — denied_destinations / d44ebdcf0e2a / 3

<a id="canonical-040740875f57bb0080ef074706da67ab2a839fc7fbf793593cfbefdf1c0cd45b"></a>

<a id="canonical-15e17e1bfd761fa5a8937c0f75d6a1b60037745a0669dc05dc63d4132f867b67"></a>

## prefix property — denied_destinations / d44ebdcf0e2a / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-9e23f8f786f3a93f3db164e9384c0d57a349c1683e8991468db8d8ecadefa40f"></a>

## Next pages — denied_destinations / d44ebdcf0e2a / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-cce21b83f7898464681e349c686b5fc293412893c52c5164fa3168b500801426"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f4154df6248446f3b40af50ce0058ad0c6f13a367adc8e887153a9665dd71ae"></a>

## denied_sources — denied_sources / a012fab515a2 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- denied_sources

<a id="canonical-9114780aa936a69ec7524513b52ee583bd246a71caf023b4c67fdc5f5a8f5f49"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
denied_sources {
  # Configure direct properties listed below.
}
```

<a id="canonical-92a97a38510709c97232c1a9ebfd054b02f12543eb2003e9455d2aa83b0f63b0"></a>

## Direct properties — denied_sources / a012fab515a2 / 3

<a id="canonical-441cbe84386d242252abcc49ae2baa1f6108dc7d2108297d42aed77d9dba587c"></a>

<a id="canonical-75211a0fee9dfc8ac512976b41fb5cf94e07202a25228be9fb45d48820b97e5a"></a>

## prefix property — denied_sources / a012fab515a2 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-4148b286cdf745737bee72430977d0260f5d8035a6fac1671da6f2861636d990"></a>

## Next pages — denied_sources / a012fab515a2 / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-b5b3a88eddb7e5e1af7a0b8e9a401a6b018ccc4a11ed7a0dc68b0d8e66a61f0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e7941934e4f8fbcd13fc6a9bda108a78a9ae412a3f3bb8f52aef90809c3cb62"></a>

## deny_all — deny_all / af5614c49021 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- deny_all

<a id="canonical-805eced645ffcf73c665dcc28322d2e8ac4c6bcc3355df92a6ad00ea910cff67"></a>

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
deny_all = {}
```

<a id="canonical-3d58ba159f7ffebd18ceecdaf1cbeb57d7d22b66fd020475cb9d118ca0f82575"></a>

## Direct properties — deny_all / af5614c49021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1825774b6a16632a39ee04c5a58bb9429084778ed067820d8fe7c06a162a7f5"></a>

## Next pages — deny_all / af5614c49021 / 4

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98df3070ede703fff705da3b96cfdc2872c2634db1a42df7bd7394920c2f805e"></a>

## rule_list — rule_list / 9896fa2eab09 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- rule_list

<a id="canonical-a2173d314aa1b1f4950102ace47c14aa6cc1566fd49bf837a27ab149f9a50355"></a>

Type: `"object"`. single nested block, Optional.

Custom Enhanced Firewall Policy Rules. Custom Enhanced Firewall Policy Rules.

Upstream description:

Custom Enhanced Firewall Policy Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-7770d84ffeea5595929d7fa4d5384bbdec5ca126b08411f6db8d04c8095a0f3a"></a>

## Direct properties — rule_list / 9896fa2eab09 / 3

- [rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306): complete subsection reference.

<a id="canonical-4242132699055e64398a9631708a147a89b0109a89659f4c7283cb70993857b1"></a>

## Next pages — rule_list / 9896fa2eab09 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8031c42b0b8ee7ffb8d9231f780f93c3bc85389883a50c4b7855d2e269aa7b56"></a>

## rule_list.rules — rule_list.rules / 456638466867 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- rule_list.rules

<a id="canonical-f1429792d32130de3aebf6d92b0d57f39ccf88fa2a28002428d9f99e2d73f24b"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policy Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_destinations",
    "all_sli_vips"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "all_slo_vips"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "all_slo_vips"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_sources",
    "inside_sources"),
  validators.ConflictingListObjectAttributes("all_sources",
    "outside_sources"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("allow",
    "deny"),
  validators.ConflictingListObjectAttributes("allow",
    "insert_service"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("deny",
    "insert_service"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_prefix_list",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_prefix_list",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("inside_destinations",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "outside_sources"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_ip_prefix_set",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("source_ip_prefix_set",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_label_selector",
    "source_prefix_list")}
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
    "minItems": 0,
    "uniqueItems": false
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
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0b225d2a1ae2130e26c40ca7223000350edb3af682ce9c7546da237e4010514"></a>

## Direct properties — rule_list.rules / 456638466867 / 3

- [advanced_action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-7911bc82b88127fce476d5d3a62341e4f577ee358011e7ef6ee9cea840980fd4): complete subsection reference.

- [all_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-feac365310bba950efbb710095a86a53f2863b86e13fcdd23bd4f72617e0cae5): complete subsection reference.

- [all_sli_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-389326855adb0ad656e848c126b631903c2f0e55ed3c4bc2731dcd642442f1f5): complete subsection reference.

- [all_slo_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a413e3a55e4e115f275d5eb264bb7a92fd55666ead7c3d1a8daba6be41586bd1): complete subsection reference.

- [all_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-14f6c5bcd2fb2ca0865b68b370adf126fdb18f29360402d668c9572912914e2a): complete subsection reference.

- [all_tcp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2fd3d827fccb6ca1c6c887a9047c58940de4224a04f0648effb782b6d3e52a94): complete subsection reference.

- [all_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a9e5893b22c8547ebc0d65f4ce885b372f98d7a41879b5070b3358d9db62f4c2): complete subsection reference.

- [all_udp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b7494586b0d978db6af307302051723f92c2067b35cbda1c4822e7ee121c786d): complete subsection reference.

- [allow](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1a0739453a584bb8cd535bcd4f145f72bff8ce4a9724458e4d5d3db66a02fb1d): complete subsection reference.

- [applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3c3621f63b087e6477f432d07c513515286db03edf8ad7adf1f0be66fd79e0c3): complete subsection reference.

- [deny](resources--enhanced_firewall_policy--reference--group-001.md#canonical-6524a6e42e40eec0a3d9a82df635616a4e67e7bc1d60fe15c882f929bb017124): complete subsection reference.

- [destination_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2355f8cb7c5c2bbd14ce30275a33f3aa4327c4c87c5c6d3123dca294636576a1): complete subsection reference.

- [destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-7dd314109557553a4999bf23321518e0d81c6c79436eec77e2cd1167bbad557c): complete subsection reference.

- [destination_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-919db6f89e2bf6c9cee6f98b669ea10206e5e8b7c3b020c2a2bf80e94393c00f): complete subsection reference.

- [destination_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b7021620c82431914a3dbc5362419ea369fc906189c68eb49567ea7b74cd3f8f): complete subsection reference.

- [insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-88c3d91eac73e24400a532b8af3bb7be0e4be20753dca7999f7d8d2a9ea6a93c): complete subsection reference.

- [inside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-246f37df1965000ccbbba1fe98d11432871844b3bdc15f0848d3f03a7aa9b83e): complete subsection reference.

- [inside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-e28b62cf4d2cc678691b7c78e0beaac269c3a8329b6ced81dba9df1f409679c3): complete subsection reference.

- [label_matcher](resources--enhanced_firewall_policy--reference--group-001.md#canonical-33073d9336381a731e4f5689b03b8217476f7646970a9b9c6bf4beed25e9c7c8): complete subsection reference.

- [metadata](resources--enhanced_firewall_policy--reference--group-001.md#canonical-014bf24ff4756de0f817fa67f2eaadd7285119a718f0a50340338efb90206ec3): complete subsection reference.

- [outside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-da339f47c1f01cffb1826b6c97b7ffa46de6931a1254b699f453d99771ed9536): complete subsection reference.

- [outside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c120ab4c728d95943c413639652a742e4c007a04e79dbec21162839f8d0daf78): complete subsection reference.

- [protocol_port_range](resources--enhanced_firewall_policy--reference--group-001.md#canonical-824d55f03e48f08f9ec80a4be95614331488f1523fdb16f65ac66ffcfd52ecae): complete subsection reference.

- [source_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3801314a92ee470a785fa24850812665d7e00c4eea6491a0a655f925c585a529): complete subsection reference.

- [source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c60fc4401d68a6a85982077c994f65617ee19e79aa412ddb7599981404855369): complete subsection reference.

- [source_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b5ed050a0436f48b2b5678c8d7c72eaef87f43a649d1eb5cf9811a2de6df516d): complete subsection reference.

- [source_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-cbb81deacfdecc97ef232d739e71879dee93c01187de6a35cf5b544464f11e75): complete subsection reference.

<a id="canonical-ad25cfeb1c1e141af73ac5ce7490041af81d53f094bd8905841e3e4c72892ee9"></a>

## Next pages — rule_list.rules / 456638466867 / 4

- [rule_list.rules.advanced_action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-7911bc82b88127fce476d5d3a62341e4f577ee358011e7ef6ee9cea840980fd4)
- [rule_list.rules.all_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-feac365310bba950efbb710095a86a53f2863b86e13fcdd23bd4f72617e0cae5)
- [rule_list.rules.all_sli_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-389326855adb0ad656e848c126b631903c2f0e55ed3c4bc2731dcd642442f1f5)
- [rule_list.rules.all_slo_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a413e3a55e4e115f275d5eb264bb7a92fd55666ead7c3d1a8daba6be41586bd1)
- [rule_list.rules.all_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-14f6c5bcd2fb2ca0865b68b370adf126fdb18f29360402d668c9572912914e2a)
- [rule_list.rules.all_tcp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2fd3d827fccb6ca1c6c887a9047c58940de4224a04f0648effb782b6d3e52a94)
- [rule_list.rules.all_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-a9e5893b22c8547ebc0d65f4ce885b372f98d7a41879b5070b3358d9db62f4c2)
- [rule_list.rules.all_udp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b7494586b0d978db6af307302051723f92c2067b35cbda1c4822e7ee121c786d)
- [rule_list.rules.allow](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1a0739453a584bb8cd535bcd4f145f72bff8ce4a9724458e4d5d3db66a02fb1d)
- [rule_list.rules.applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3c3621f63b087e6477f432d07c513515286db03edf8ad7adf1f0be66fd79e0c3)
- [rule_list.rules.deny](resources--enhanced_firewall_policy--reference--group-001.md#canonical-6524a6e42e40eec0a3d9a82df635616a4e67e7bc1d60fe15c882f929bb017124)
- [rule_list.rules.destination_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2355f8cb7c5c2bbd14ce30275a33f3aa4327c4c87c5c6d3123dca294636576a1)
- [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-7dd314109557553a4999bf23321518e0d81c6c79436eec77e2cd1167bbad557c)
- [rule_list.rules.destination_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-919db6f89e2bf6c9cee6f98b669ea10206e5e8b7c3b020c2a2bf80e94393c00f)
- [rule_list.rules.destination_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b7021620c82431914a3dbc5362419ea369fc906189c68eb49567ea7b74cd3f8f)
- [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-88c3d91eac73e24400a532b8af3bb7be0e4be20753dca7999f7d8d2a9ea6a93c)
- [rule_list.rules.inside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-246f37df1965000ccbbba1fe98d11432871844b3bdc15f0848d3f03a7aa9b83e)
- [rule_list.rules.inside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-e28b62cf4d2cc678691b7c78e0beaac269c3a8329b6ced81dba9df1f409679c3)
- [rule_list.rules.label_matcher](resources--enhanced_firewall_policy--reference--group-001.md#canonical-33073d9336381a731e4f5689b03b8217476f7646970a9b9c6bf4beed25e9c7c8)
- [rule_list.rules.metadata](resources--enhanced_firewall_policy--reference--group-001.md#canonical-014bf24ff4756de0f817fa67f2eaadd7285119a718f0a50340338efb90206ec3)
- [rule_list.rules.outside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-da339f47c1f01cffb1826b6c97b7ffa46de6931a1254b699f453d99771ed9536)
- [rule_list.rules.outside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c120ab4c728d95943c413639652a742e4c007a04e79dbec21162839f8d0daf78)
- [rule_list.rules.protocol_port_range](resources--enhanced_firewall_policy--reference--group-001.md#canonical-824d55f03e48f08f9ec80a4be95614331488f1523fdb16f65ac66ffcfd52ecae)
- [rule_list.rules.source_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3801314a92ee470a785fa24850812665d7e00c4eea6491a0a655f925c585a529)
- [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c60fc4401d68a6a85982077c994f65617ee19e79aa412ddb7599981404855369)
- [rule_list.rules.source_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-b5ed050a0436f48b2b5678c8d7c72eaef87f43a649d1eb5cf9811a2de6df516d)
- [rule_list.rules.source_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-cbb81deacfdecc97ef232d739e71879dee93c01187de6a35cf5b544464f11e75)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-7911bc82b88127fce476d5d3a62341e4f577ee358011e7ef6ee9cea840980fd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-696abdf03fe6a34021c655944002e87d2383f6aa6056fbe63b03aeaee86fa35c"></a>

## rule_list.rules.advanced_action — rule_list.rules.advanced_action / d69ea0141d21 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.advanced_action

<a id="canonical-e15a8d03863e0f7d3011f828a73aa3e3446c152bf17666f1e989b37acb410d3e"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
advanced_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3b637708a46395a83a60b643c367c3c6da571c770a18bba6ea123f73e6d55413"></a>

## Direct properties — rule_list.rules.advanced_action / d69ea0141d21 / 3

<a id="canonical-69ecddbad856b6f68823de187992afbc00b2e6dc1fc6cecd097388c8b123ed9d"></a>

<a id="canonical-b9bff75ba4d1401f90fdf14ab39029de0f08fccac1c3c60fa3c36f1b5c5f799e"></a>

## action property — rule_list.rules.advanced_action / d69ea0141d21 / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a95cb5db37d6b4e7367df1e342145d87028bc263759cb9f37dcac61b64b6d2d1"></a>

## Next pages — rule_list.rules.advanced_action / d69ea0141d21 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-feac365310bba950efbb710095a86a53f2863b86e13fcdd23bd4f72617e0cae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb6773ecb7c9190368b445f7135cdc91b5f11e1f058c9a7c7a9507574e13c2c2"></a>

## rule_list.rules.all_destinations — rule_list.rules.all_destinations / cdeadc0bc83d / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.all_destinations

<a id="canonical-9b0e97754e6a04ef6d0dd806359b354589fe0554615afa5385b63b5e2bcf8898"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

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
all_destinations = {}
```

<a id="canonical-ec9d8a60f421e17bf283403702fbb7130877a64a080bf7ffbe8c4dd720caa55b"></a>

## Direct properties — rule_list.rules.all_destinations / cdeadc0bc83d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae67b22dc93a4fa3f54ed6233168d3958cdfda43e682401a69308d11e64040f0"></a>

## Next pages — rule_list.rules.all_destinations / cdeadc0bc83d / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-389326855adb0ad656e848c126b631903c2f0e55ed3c4bc2731dcd642442f1f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f2fa532f2c663bd3366e10aa3eda5a33e621fcfeda1ccaf345753b8cec3f209"></a>

## rule_list.rules.all_sli_vips — rule_list.rules.all_sli_vips / 5f3e708f7fac / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.all_sli_vips

<a id="canonical-9f6de4b09d03134b5e161ed2e0daedfdfc1307e007df3a7bc88475bcd8769226"></a>

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
all_sli_vips = {}
```

<a id="canonical-312e118e4b6d6130c1a80bbc94d18aac9102dbe71ba79b3f54b1ee34dbe4f94f"></a>

## Direct properties — rule_list.rules.all_sli_vips / 5f3e708f7fac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae48d9ef32a5622ff9e203740ea456c1cb66db624cdef09e03446d1d53fafad6"></a>

## Next pages — rule_list.rules.all_sli_vips / 5f3e708f7fac / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-a413e3a55e4e115f275d5eb264bb7a92fd55666ead7c3d1a8daba6be41586bd1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d8e26928ce48cffb03349c1ddbe7909d6c95c4ef0cf99e9f700dc7624ea9e14"></a>

## rule_list.rules.all_slo_vips — rule_list.rules.all_slo_vips / 6ec448406995 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.all_slo_vips

<a id="canonical-5b9cc39d2a21f0f0a2a1e7fb6404ed876d12568c55c7b6f9d81d5ecea8e45e3d"></a>

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
all_slo_vips = {}
```

<a id="canonical-4319f28bb9837d7139c5081fde8ce90a5026cf7e38b1369728d3e61bb299161c"></a>

## Direct properties — rule_list.rules.all_slo_vips / 6ec448406995 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ea721778d82d78d49d14784cbaf042632b1d972fa2f7ed070e44e401c258cbc"></a>

## Next pages — rule_list.rules.all_slo_vips / 6ec448406995 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-14f6c5bcd2fb2ca0865b68b370adf126fdb18f29360402d668c9572912914e2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64613bd9c306b712f796239c6371cecc2343e4afd4af47ac3175ad09151aa602"></a>

## rule_list.rules.all_sources — rule_list.rules.all_sources / 8f1adb494b73 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.all_sources

<a id="canonical-b414e6d4c388ece624be15a8da857fa217add4e83b99cecca973dc99ae0dc5d0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all sources.

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
all_sources = {}
```

<a id="canonical-cfdd4833998860262e38ccbd606bd07ad8bd4bfe879366dd3518b3d20021d75f"></a>

## Direct properties — rule_list.rules.all_sources / 8f1adb494b73 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57a2f25dc101bdb4975ecd907a53bb92ae4a9c3b90ae30311bcc98fa8d39eb00"></a>

## Next pages — rule_list.rules.all_sources / 8f1adb494b73 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-2fd3d827fccb6ca1c6c887a9047c58940de4224a04f0648effb782b6d3e52a94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2fa28151bfa6783e88eeaba8e5a48ae0b94f66bd34a647fd15a424613da85af"></a>

## rule_list.rules.all_tcp_traffic — rule_list.rules.all_tcp_traffic / ceb76efe957f / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.all_tcp_traffic

<a id="canonical-61964d4e8778f040e6657a53df4848901e823d0068a3104395e1f4ad3eb7f965"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

<a id="canonical-e908263cdb441da67dcd3da85a590c166267e017da39e295a64d46e24f8c99c9"></a>

## Direct properties — rule_list.rules.all_tcp_traffic / ceb76efe957f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0658e51d8234217bfbb08edb6db12e1ab5252f9b5564090d0792e363de52f62"></a>

## Next pages — rule_list.rules.all_tcp_traffic / ceb76efe957f / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-a9e5893b22c8547ebc0d65f4ce885b372f98d7a41879b5070b3358d9db62f4c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a682f7d368e253abce245b80c9501c2bf001d6fe59143a1b8c17409d5ab8d739"></a>

## rule_list.rules.all_traffic — rule_list.rules.all_traffic / 0a98ae2b59d3 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.all_traffic

<a id="canonical-3e3256e04a30f4650def0e2e368febef5293a755aa6faf05619a6a5cae1c794b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

<a id="canonical-5c63e3f32ba467de59244ed2e73401aaba8ffa6f6b2c63c5be043c3858501c32"></a>

## Direct properties — rule_list.rules.all_traffic / 0a98ae2b59d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4702d6a5288e3a84723cf33ae6211189f62e9b094d44d447302b431c21d39f2a"></a>

## Next pages — rule_list.rules.all_traffic / 0a98ae2b59d3 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-b7494586b0d978db6af307302051723f92c2067b35cbda1c4822e7ee121c786d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-901511dd7a28d1af717b31413c90a16b94bdbd7e01add4b0425fcf951a2adf32"></a>

## rule_list.rules.all_udp_traffic — rule_list.rules.all_udp_traffic / af2812ed813c / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.all_udp_traffic

<a id="canonical-92b3a1ed390e1b311e5cc66a74069086311504e517bf538a83d76337eeb4e440"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

<a id="canonical-44a124b11902cff91d3ef5f3a42986420578ccf32ca738d66c45776913dc4fd8"></a>

## Direct properties — rule_list.rules.all_udp_traffic / af2812ed813c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ded3c71f69e1163d054dcbb1384416a1294ee1ddeb6cb1e917ed0beeaf68b831"></a>

## Next pages — rule_list.rules.all_udp_traffic / af2812ed813c / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-1a0739453a584bb8cd535bcd4f145f72bff8ce4a9724458e4d5d3db66a02fb1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0edf52926fb0efc2fc224e55563b30f4bc92a6005be91de3e6e2feabf81a3f26"></a>

## rule_list.rules.allow — rule_list.rules.allow / d7516275bcf7 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.allow

<a id="canonical-650d2c009a8d4374e2ecfba078c224d218b1546825558f3a8bd2a5d080f249b8"></a>

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
allow = {}
```

<a id="canonical-ac528f3169194d17e17fb62cbed60848c93a0bab2f6cc5831c44dbdf3acef607"></a>

## Direct properties — rule_list.rules.allow / d7516275bcf7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc1a664e894f0490b1ea1a987d0036d49741346072423a411e719b6021dd3d26"></a>

## Next pages — rule_list.rules.allow / d7516275bcf7 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-3c3621f63b087e6477f432d07c513515286db03edf8ad7adf1f0be66fd79e0c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bcc024ac04704eda65dd63d45f0ded53ecb73510a33876341e8cc7ded4f6b06"></a>

## rule_list.rules.applications — rule_list.rules.applications / 1f8760fd3f65 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.applications

<a id="canonical-15f202ca085a25e43e2a80d9be2a039582ddf302a09f882853fb6112feafa25e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-c0ab1401b471288da4b44b8d257880ad80bb2ad6e6a57f4bde036fd7474f50f1"></a>

## Direct properties — rule_list.rules.applications / 1f8760fd3f65 / 3

<a id="canonical-3524aa71a0884ba161171ae7f6bc5f0511e287ecb6dbb724d1453d7cfbb4d625"></a>

<a id="canonical-65e09982189795f6a636b50c2b23fd07208d4f962c9277b4abfa6b65196e22f8"></a>

## applications property — rule_list.rules.applications / 1f8760fd3f65 / 4

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-c701ba94bdd3d2ce8979527b7fbd2c4ab0c4f489fedae315fd051f97e88146f6"></a>

## Next pages — rule_list.rules.applications / 1f8760fd3f65 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-6524a6e42e40eec0a3d9a82df635616a4e67e7bc1d60fe15c882f929bb017124"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ca9e5ac5d101b2d018f20d5cdae21d866480e91675d4d4b16ddafd7900e96e3"></a>

## rule_list.rules.deny — rule_list.rules.deny / 94c0f914db40 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.deny

<a id="canonical-3a0e3f106ddcec9451ceb07a188ac42e1905da0b607cacac4c4c11ad5e45eb61"></a>

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
deny = {}
```

<a id="canonical-ad2f99c2f7770393e7070c6d066a8586aa879c392bc6e4277e21493481fda0cc"></a>

## Direct properties — rule_list.rules.deny / 94c0f914db40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b24818232b65fd98c92eaa92bcd4df41cfdc642df03fe0e9ffb21b0d9cbf746"></a>

## Next pages — rule_list.rules.deny / 94c0f914db40 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-2355f8cb7c5c2bbd14ce30275a33f3aa4327c4c87c5c6d3123dca294636576a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d716699ddfea3084511d219c12bfa38bc9e36cf25d43a0feb0f71594cfa4e6c5"></a>

## rule_list.rules.destination_aws_vpc_ids — rule_list.rules.destination_aws_vpc_ids / 0fe7c20b3ec0 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.destination_aws_vpc_ids

<a id="canonical-27ace91d66e071e0371ae741f7a7cf9763bc4c749dd163cee6a0495c809fb7a0"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for destination aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("vpc_id")}
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
destination_aws_vpc_ids {
  # Configure direct properties listed below.
}
```

<a id="canonical-73341a4358ef314de57e9ab0c94a85faa8a6a95bea7503007b317707d7775a2e"></a>

## Direct properties — rule_list.rules.destination_aws_vpc_ids / 0fe7c20b3ec0 / 3

<a id="canonical-25353afa2f55215632c9a32330ffee115492292ec74a5b6df5ad70b99499e1ae"></a>

<a id="canonical-dfc3a9714ec82a97447db08af8b55f1f61117ce6a7254804a99525a2d5788b6e"></a>

## vpc_id property — rule_list.rules.destination_aws_vpc_ids / 0fe7c20b3ec0 / 4

Type: `["list", "string"]`. Optional.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0ed7bf7e83c4fa02e43933d3f74345bc5aa87ae2d0268b7c941449a42e7809ed"></a>

## Next pages — rule_list.rules.destination_aws_vpc_ids / 0fe7c20b3ec0 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-7dd314109557553a4999bf23321518e0d81c6c79436eec77e2cd1167bbad557c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-417887921aad70f7a89166ce35f6aac56ec97951ecf04ba6ea969cf6cb9c063d"></a>

## rule_list.rules.destination_ip_prefix_set — rule_list.rules.destination_ip_prefix_set / 232314e24b9d / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.destination_ip_prefix_set

<a id="canonical-2d0421ea370b3327b76b263abf00b0e2309cdac69f40b806c9334667f8c4e7b9"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
destination_ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a2817ab937cdbf754af508a74197289eb4e3121d882fcced9e0fa29a8b14160"></a>

## Direct properties — rule_list.rules.destination_ip_prefix_set / 232314e24b9d / 3

- [ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-fc94c8e6933ac31dad174a436d98bccee444b7100855f48abe6ba45041dc7bcc): complete subsection reference.

<a id="canonical-ae93d7bd40ec7b8e4494b911baf8562ae90f33818860d381a1590573b923fa19"></a>

## Next pages — rule_list.rules.destination_ip_prefix_set / 232314e24b9d / 4

- [rule_list.rules.destination_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-fc94c8e6933ac31dad174a436d98bccee444b7100855f48abe6ba45041dc7bcc)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-fc94c8e6933ac31dad174a436d98bccee444b7100855f48abe6ba45041dc7bcc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09a067e0a5c33c58d91723ff8bfdcbb3a0114ed4d5f2d2fe3689c46770653d5c"></a>

## rule_list.rules.destination_ip_prefix_set.ref — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-7dd314109557553a4999bf23321518e0d81c6c79436eec77e2cd1167bbad557c)
- rule_list.rules.destination_ip_prefix_set.ref

<a id="canonical-ef95a667ce70817641e2a89fbbac8c6b4537df1abcf604782fc4c146aafda680"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-74454ce683bd17938a6e051726f715ea680df3c2d5183439f27267c47003f00c"></a>

## Direct properties — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 3

<a id="canonical-8c91abd2c13c19b9389426ff361aab7d261632618e4be242efa82eec3e930a64"></a>

<a id="canonical-c6436c406db5043294b6e461600a9aa856fc0d653dbcfa44b6c7815a1f3f2684"></a>

## kind property — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 4

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

<a id="canonical-ef2b97d2eac0edaba0a454487d895ffa9e7ff8632ac11170a49ae304426c127a"></a>

<a id="canonical-373e968d22240dbad15407aac55665cd4df38a215f3157ed797160aabe7d3089"></a>

## name property — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 5

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

<a id="canonical-0a90da958902b192d054247c3376b3f2936a8b8e4da25c22416f27e04e066b6b"></a>

<a id="canonical-9e50024384b489924a30aea9b4a4e5c59750c954be8383b8dbd0433196e4ae89"></a>

## namespace property — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 6

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

<a id="canonical-1e86056f714e8138f6521e06f30fb2b0c036d3e9d7c616572b7b4accc50306f0"></a>

<a id="canonical-3c38fa70624dbecfd0c5fb30e3b8dd4455be29815f629c654a644a679671b598"></a>

## tenant property — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 7

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

<a id="canonical-c1884b2b4db474ab89f7af4416bec0d74ee04e9683c196c25336f48b1111045f"></a>

<a id="canonical-09bb42df6a6a236fc712555bef20a6891b1311c407b1f7a59fd1777cfd534ad1"></a>

## uid property — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 8

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

<a id="canonical-8f9aa0f7a78d160f5ddb4620954d900fc15ab617fffcff5b3eab08101bb20e8e"></a>

## Next pages — rule_list.rules.destination_ip_prefix_set.ref / ec182d14f648 / 9

- [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-7dd314109557553a4999bf23321518e0d81c6c79436eec77e2cd1167bbad557c)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-919db6f89e2bf6c9cee6f98b669ea10206e5e8b7c3b020c2a2bf80e94393c00f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1acbc118ab97c2a544f63d3c79d86da0a21e17a782782b6a0bb2ca175b8ab79a"></a>

## rule_list.rules.destination_label_selector — rule_list.rules.destination_label_selector / 5be18b50e740 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.destination_label_selector

<a id="canonical-400e7043e67357286c01c0151a3c8bd107775d7b2f98fe19ee8d7ad74549a274"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
destination_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-52a3e0b00dae9a9c1b20697839a42732ccd9f10ed99a4ecf25d57c887b5dc60e"></a>

## Direct properties — rule_list.rules.destination_label_selector / 5be18b50e740 / 3

<a id="canonical-ae86c865ec2512011411eb2814bad24ae8563f1f9a17cce013eaf8a22c8e764c"></a>

<a id="canonical-cc9f904551aa9bd1b2b6cae5f567d65643cc2d7f0907132c241ff90ebc82eb16"></a>

## expressions property — rule_list.rules.destination_label_selector / 5be18b50e740 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-09f3e9be1b0db56a68ab3caee3367b760303e62cf4dbf1e799c852dece2b8e0f"></a>

## Next pages — rule_list.rules.destination_label_selector / 5be18b50e740 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-b7021620c82431914a3dbc5362419ea369fc906189c68eb49567ea7b74cd3f8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4531925547849c546df20a8a1534374b4a21f27dc0b50bf65fee60468550f587"></a>

## rule_list.rules.destination_prefix_list — rule_list.rules.destination_prefix_list / c2a3b6b840ba / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.destination_prefix_list

<a id="canonical-12453a2bdbaa8a11b02c289383236078706d8ece05862d089eefa7be20c2b88f"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
destination_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-e26646c6478a285ed7723c3620fb8ffb7ecaaf2bb1404980fa3cf4a144a3b347"></a>

## Direct properties — rule_list.rules.destination_prefix_list / c2a3b6b840ba / 3

<a id="canonical-2f9e2406791e89390749848e98da1898d58b815ad2cc5a2b14a3bcd03e4d6d9b"></a>

<a id="canonical-e64997ba052cba7ea976b8d65fc02eb59f5db10cc02c2b3a4457af4c801d9349"></a>

## prefixes property — rule_list.rules.destination_prefix_list / c2a3b6b840ba / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-136df690ddc93e911e43caf7c128a8516aeab9e002ba6d70c35a5558b292ac81"></a>

## Next pages — rule_list.rules.destination_prefix_list / c2a3b6b840ba / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-88c3d91eac73e24400a532b8af3bb7be0e4be20753dca7999f7d8d2a9ea6a93c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-696d3212a0573035c819a41b5d4af1c53cdfdf4325d0d998f4e2b81b4fdd11d2"></a>

## rule_list.rules.insert_service — rule_list.rules.insert_service / 694df9d23a87 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.insert_service

<a id="canonical-3cb99e68ac3efdbfcc2616cea1a79ec38940a5431c314f62e29355253918336e"></a>

Type: `"object"`. single nested block, Optional.

Action to forward traffic to external service.

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
insert_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e7e4aa3ba2cd3869791aa6237750353089f8f64df630b47d5285c5cb9e6e2d4"></a>

## Direct properties — rule_list.rules.insert_service / 694df9d23a87 / 3

- [nfv_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-e93d1bd8bec29d2ea3167ac08a7beccafa0ca3453553ecda2b08f32ad89917db): complete subsection reference.

<a id="canonical-563b00339390090162eedf29551ee205d6e573fa404e6676afc70f11f699cae1"></a>

## Next pages — rule_list.rules.insert_service / 694df9d23a87 / 4

- [rule_list.rules.insert_service.nfv_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-e93d1bd8bec29d2ea3167ac08a7beccafa0ca3453553ecda2b08f32ad89917db)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-e93d1bd8bec29d2ea3167ac08a7beccafa0ca3453553ecda2b08f32ad89917db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4520dc9a6520472bf2d91be7f71263ee13fcee5c6d63b70f0b9e19047b3e0d5"></a>

## rule_list.rules.insert_service.nfv_service — rule_list.rules.insert_service.nfv_service / 8298726bc71d / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-88c3d91eac73e24400a532b8af3bb7be0e4be20753dca7999f7d8d2a9ea6a93c)
- rule_list.rules.insert_service.nfv_service

<a id="canonical-691e517290e7b0efa5d615acd90b388bbdff7ff4ba130b3a60bd818bf34d1faf"></a>

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
nfv_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-865a50f276573fa71f7aa03d9b2889191e5597bf866ec2f25956d0cc6c5a93c3"></a>

## Direct properties — rule_list.rules.insert_service.nfv_service / 8298726bc71d / 3

<a id="canonical-5e16ad124ef4bb72c89fe250fe7563264b4ff16a2a3419910cb42d177267899d"></a>

<a id="canonical-741ed632b4cc70567f3a122ef1d32c3bc7942dfcf1b493eebc39d0ac66e04aa5"></a>

## name property — rule_list.rules.insert_service.nfv_service / 8298726bc71d / 4

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

<a id="canonical-ae46c2ed8ef5b916136df5226a6a3e349d3d3ad3c20410c8effa0c328281c9af"></a>

<a id="canonical-2f9d30bbcb3ccfba13f6105a0dc23d0cc0369815349637d05655aeaf7105382f"></a>

## namespace property — rule_list.rules.insert_service.nfv_service / 8298726bc71d / 5

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

<a id="canonical-cead5103a841d878f9b9780db3f99e9ae4225ad33b48bfc4445e7f40f0a2ae5d"></a>

<a id="canonical-2a31cc0bf8715f0ec0367fd38cdcba7575c7c908a697404781daffc8c91f3423"></a>

## tenant property — rule_list.rules.insert_service.nfv_service / 8298726bc71d / 6

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

<a id="canonical-bea4b72dc811f5a863d83b23c430a2ff34a021182f8781a7a4307bb7a4cad5d5"></a>

## Next pages — rule_list.rules.insert_service.nfv_service / 8298726bc71d / 7

- [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-88c3d91eac73e24400a532b8af3bb7be0e4be20753dca7999f7d8d2a9ea6a93c)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-246f37df1965000ccbbba1fe98d11432871844b3bdc15f0848d3f03a7aa9b83e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-677a16d30b965b12573da88e7a930fa3ee8314d7d246d14d4c9802240803d33b"></a>

## rule_list.rules.inside_destinations — rule_list.rules.inside_destinations / 5631112c4c52 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.inside_destinations

<a id="canonical-5798d14d3b926b08bb93784235c7cf33148191678897981b2f773720593ea7c5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside destinations.

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
inside_destinations = {}
```

<a id="canonical-a94eee60904677130df8849b8b77ba2f52a719c2769b882f9c827a24a1c2ebc3"></a>

## Direct properties — rule_list.rules.inside_destinations / 5631112c4c52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-385ded4d61eca9676c00b6ab911f6f082e4810cff0ef3057efe57b44e0b303a1"></a>

## Next pages — rule_list.rules.inside_destinations / 5631112c4c52 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-e28b62cf4d2cc678691b7c78e0beaac269c3a8329b6ced81dba9df1f409679c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9fd13514375796e25f4c462cab684d17e6d30149ed5efec22dd78e298d6a692"></a>

## rule_list.rules.inside_sources — rule_list.rules.inside_sources / 353219ed5624 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.inside_sources

<a id="canonical-cfc814ea2b12b916687aad60bf5be35f5b688904c94d733cb9cd4d8962d231e3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside sources.

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
inside_sources = {}
```

<a id="canonical-389ac3e450ea62a1bd972ccc053323603b4c62cc62620fec48f0acac0b89438b"></a>

## Direct properties — rule_list.rules.inside_sources / 353219ed5624 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c8b0ec4b20dcc733c4c57400a3da5750f1aef960619fa4a1db3381aa72ccb77"></a>

## Next pages — rule_list.rules.inside_sources / 353219ed5624 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-33073d9336381a731e4f5689b03b8217476f7646970a9b9c6bf4beed25e9c7c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5443676f1b580f486ddbb0fb10d15999e6283a7b2db5df9d4f37ae6716d82010"></a>

## rule_list.rules.label_matcher — rule_list.rules.label_matcher / 840bddea22c1 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.label_matcher

<a id="canonical-a2f1c78ed06416d050edade323eab72a4be0838b9bc08d7ab1230b3922659f3a"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2ac0ac810f270eb031c570ba2f78841c833a35f37d40b48f419063b27eb9050d"></a>

## Direct properties — rule_list.rules.label_matcher / 840bddea22c1 / 3

<a id="canonical-93a79c704876d20ab4201c39c03b2bf270cb333cb7950386bc5c34df6e941517"></a>

<a id="canonical-ba99856f2f07b0a3dfd9ed8c020b0e165aad9ec98e53205dde623c24b328e5d6"></a>

## keys property — rule_list.rules.label_matcher / 840bddea22c1 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-215cb480ce715d54493cc51208bfea3eda6b7466cf8f3599a1b5bd9cf17deac1"></a>

## Next pages — rule_list.rules.label_matcher / 840bddea22c1 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-014bf24ff4756de0f817fa67f2eaadd7285119a718f0a50340338efb90206ec3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32f6436c14d0aa326866cd0ce0b838109deaf9779a73513118e749ca70ec9ee6"></a>

## rule_list.rules.metadata — rule_list.rules.metadata / aa9e5b351c31 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.metadata

<a id="canonical-6ac2f0cc6db8aacdeeed2cc6b5b37b47ebb1689157ee3c97b487468f2df94eca"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-b604eb5faed591b60a9ed4a87bb57439ab64d5a48e9e05dd8dfa9610dfac883b"></a>

## Direct properties — rule_list.rules.metadata / aa9e5b351c31 / 3

<a id="canonical-9438bda3b1ca25f15aa9831e46325bfb175a91096c0540e15f6324884839de9a"></a>

<a id="canonical-c98995b98e7ecd11450c3e8f84f7c5fa8788a747521e031ff0a6ca0c51f0ae34"></a>

## description_spec property — rule_list.rules.metadata / aa9e5b351c31 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-89049dfdce8fc4646f5ff88b29ec5021b64ab432a8ce254425199c97b29baf0e"></a>

<a id="canonical-5d28089c513ffed9a6c40cf8168e5da3ec5bf54f2ef95a80ed8ddb9627ebab2c"></a>

## name property — rule_list.rules.metadata / aa9e5b351c31 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-169d29b4b53af13ef00df8f43da1ecd15d35779be7adbe6f76dc61b4c6dc8db7"></a>

## Next pages — rule_list.rules.metadata / aa9e5b351c31 / 6

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-da339f47c1f01cffb1826b6c97b7ffa46de6931a1254b699f453d99771ed9536"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d13cdbba8be479cae3dce72e35ccc8d88c272d8ba6fb2e8b594558d19efa25eb"></a>

## rule_list.rules.outside_destinations — rule_list.rules.outside_destinations / a2843ad5e8db / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.outside_destinations

<a id="canonical-c1655578dcff8913053b1666b67c03340bff5ead56f383ddf0fde633245b2292"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside destinations.

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
outside_destinations = {}
```

<a id="canonical-c66f1eff635944bd19819176e90a1475e5841091821312f148f01b4cb18bbbc5"></a>

## Direct properties — rule_list.rules.outside_destinations / a2843ad5e8db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-884aa0a2fbbe39fbd7e8660ad6293894abaf8e335184adf43582392e0fb556e8"></a>

## Next pages — rule_list.rules.outside_destinations / a2843ad5e8db / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-c120ab4c728d95943c413639652a742e4c007a04e79dbec21162839f8d0daf78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23471b3c4a16bad0c3b142698e35af8c866574d5cca3439a9510b32ee3d28388"></a>

## rule_list.rules.outside_sources — rule_list.rules.outside_sources / e91398f3a73b / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.outside_sources

<a id="canonical-b4545d22c4660b29c684a14e10c6e1fa96cf76809ef61e65d92ae658fed92d60"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside sources.

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
outside_sources = {}
```

<a id="canonical-97e3ad3f0a344a4e43a4e8344715d6f0ba49c97747b79cf607983616d076f962"></a>

## Direct properties — rule_list.rules.outside_sources / e91398f3a73b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3fee335b947574a81abc3aa3cd58b7f107f9355034f79d64d607e9575d810e90"></a>

## Next pages — rule_list.rules.outside_sources / e91398f3a73b / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-824d55f03e48f08f9ec80a4be95614331488f1523fdb16f65ac66ffcfd52ecae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fed5ec36ef6a1486f83e40fd61c96d8f97d99e27f9b66d8f60b1c3b4893619e"></a>

## rule_list.rules.protocol_port_range — rule_list.rules.protocol_port_range / 5313586f3f2e / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.protocol_port_range

<a id="canonical-27943e54491b481302c46c379b8bd092ea5a072f4474bd3fa33bef20f6a597d4"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-821b47707f739f58f0bd4277e4f0bb71edc24e45789a53881e1f9125b4e0d2c3"></a>

## Direct properties — rule_list.rules.protocol_port_range / 5313586f3f2e / 3

<a id="canonical-ed8b0a0be79e76a1d37f3781ed41b96594fbda5babb7e08f055a4ec7a1d4768e"></a>

<a id="canonical-853c5f4ccdcc6d0513c96bc85ad1482098185d22f0b060b445a95736b094f048"></a>

## port_ranges property — rule_list.rules.protocol_port_range / 5313586f3f2e / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-d9cf0de8df6bef566e9ff47ad4e6fe2bd0862a9bd56b4a5a9eb2848c733f1955"></a>

<a id="canonical-1bf6a30dae245de209f2f9c1109c83a21d112e8e384d96738da02370d9acce35"></a>

## protocol property — rule_list.rules.protocol_port_range / 5313586f3f2e / 5

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-3c6aa93879022122d459239b8983856b545db8b710deddf2ca4581923a5e2d16"></a>

## Next pages — rule_list.rules.protocol_port_range / 5313586f3f2e / 6

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-3801314a92ee470a785fa24850812665d7e00c4eea6491a0a655f925c585a529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a9ee6fdbc753a26cfc443ace7422ed460b58e2acf67c32683472b91760329e6"></a>

## rule_list.rules.source_aws_vpc_ids — rule_list.rules.source_aws_vpc_ids / e09c86ba386a / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.source_aws_vpc_ids

<a id="canonical-f211ce5b827cadb9d42eeefe07463bef893256c137fe89cf28897b1f094c6c1a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for source aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("vpc_id")}
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
source_aws_vpc_ids {
  # Configure direct properties listed below.
}
```

<a id="canonical-cfe98e560d43898a6df5a3d7c79c4ec4ad3892317b0da61ac449d76389f9f259"></a>

## Direct properties — rule_list.rules.source_aws_vpc_ids / e09c86ba386a / 3

<a id="canonical-ef786556bccd94472fea928b293c15cc86457629c00435713e99216b00ed2e78"></a>

<a id="canonical-94f49ec75055f7d75955a41c57e8a7b1022eb5929c50530707e99d420b9e0f3a"></a>

## vpc_id property — rule_list.rules.source_aws_vpc_ids / e09c86ba386a / 4

Type: `["list", "string"]`. Optional.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f51f4fc6d0a669bc7845ac5ea08ca35de681b7677c563ebb18ee06a403ed9962"></a>

## Next pages — rule_list.rules.source_aws_vpc_ids / e09c86ba386a / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-c60fc4401d68a6a85982077c994f65617ee19e79aa412ddb7599981404855369"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-214c69fbca9d5b9183cdb9d29915200346b50aea2e8c4c52119dd255ec015778"></a>

## rule_list.rules.source_ip_prefix_set — rule_list.rules.source_ip_prefix_set / 2d1301f04015 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.source_ip_prefix_set

<a id="canonical-07f0ed96a8bf58cef01ae1fec6efbfd7fe2eb85d5ddc35ac679e437ade8029d7"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
source_ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7516fd7f9b5cbfa117cb3898b28ae8e2ee7619201e93fd03dcd39f5f7df697f"></a>

## Direct properties — rule_list.rules.source_ip_prefix_set / 2d1301f04015 / 3

- [ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0af954fbe9a58b0133be04a3e35c553f19ec37e7f7befb47fa2fb0a6676a5093): complete subsection reference.

<a id="canonical-626e236418ce1d57d5c28de715a774c5f90e45980a5f6e387d41c2a5ccd30c9e"></a>

## Next pages — rule_list.rules.source_ip_prefix_set / 2d1301f04015 / 4

- [rule_list.rules.source_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0af954fbe9a58b0133be04a3e35c553f19ec37e7f7befb47fa2fb0a6676a5093)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-0af954fbe9a58b0133be04a3e35c553f19ec37e7f7befb47fa2fb0a6676a5093"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7250a47655d9da82b367a3d32abf9ec140185cc933c9a4fcea184c42fa5f9af8"></a>

## rule_list.rules.source_ip_prefix_set.ref — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c60fc4401d68a6a85982077c994f65617ee19e79aa412ddb7599981404855369)
- rule_list.rules.source_ip_prefix_set.ref

<a id="canonical-ce287f79f7675b50a93c454623a6b002c06ca0abbbc17e21ed39c1066cdb448d"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-87edb75895897d48722bb7c3a01a8cf7fc1bfb821b17e878cf91c2fed54d485b"></a>

## Direct properties — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 3

<a id="canonical-f43c3c157f8caf6c025209c67dda2aefdaa4499f7d96ad3acd9ef43de3705031"></a>

<a id="canonical-30b91798a84205e5c68b84b950fd8ab0757304872713c584b66bd8e119a5e242"></a>

## kind property — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 4

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

<a id="canonical-342966d73a299729b3c9f67f57251e073f1dcc5af914cfc03403993a2aebdc22"></a>

<a id="canonical-31e0b5ee5a0a1c85747a78cd1b390d3f620b1ae4ac6a63c204937b287bfbeb45"></a>

## name property — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 5

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

<a id="canonical-d30c25f3165f0b4239d55a919b408179720de3a249b626f7a37e31299e2cd79a"></a>

<a id="canonical-aa9555e5198166dbf04da72efac1eeeddb292427b47f48e409f15a291b0ac2c9"></a>

## namespace property — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 6

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

<a id="canonical-afbb854760e44da5ad3edc949080a0debec533770059e920dcdf2d381b709dea"></a>

<a id="canonical-c363804be2ffde17d9a94a638ccd91d0a6ea9425fd4ef80843b472f53b1ab763"></a>

## tenant property — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 7

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

<a id="canonical-396a07c5304d78c05492c5cd312775f3dfb20b9aa95a6a4f3cbae5047d2b2162"></a>

<a id="canonical-5834c6398e716a1c8cc3b65027314d87bb4b0b7ef62eb3bb351b55f45eae3648"></a>

## uid property — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 8

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

<a id="canonical-24f9732790df69bee4a6bf1c5b37d09597de0016c57d9bc715ea8f9be1d9fbd4"></a>

## Next pages — rule_list.rules.source_ip_prefix_set.ref / 976c5c88592d / 9

- [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-c60fc4401d68a6a85982077c994f65617ee19e79aa412ddb7599981404855369)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-b5ed050a0436f48b2b5678c8d7c72eaef87f43a649d1eb5cf9811a2de6df516d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dd886e57d90a62827e0667514ecdf2edf9225d8431505a34d108a11099b6a6d"></a>

## rule_list.rules.source_label_selector — rule_list.rules.source_label_selector / 1e315f4d3341 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.source_label_selector

<a id="canonical-4541f9ca68cc7f2db46c31038c319e61f59cef0ccaaf5d9887ac2ef5fd677735"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
source_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-61f19a963f5932e1f91a83a30cadc605c3d3fba2eb9d944aba66350827ccdb23"></a>

## Direct properties — rule_list.rules.source_label_selector / 1e315f4d3341 / 3

<a id="canonical-350fe43b6babc908337405ac68d8ad6a869a44616e52f399d90d385417db55ae"></a>

<a id="canonical-710eab5dc07722a24de3ee9d5bbeea4e819f8090fe9ec97e9eb8aa443ad81367"></a>

## expressions property — rule_list.rules.source_label_selector / 1e315f4d3341 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-d278c421fb1f84c25c9629a7f775a767fb45ddf50dbbd800dd47a9678a72eea6"></a>

## Next pages — rule_list.rules.source_label_selector / 1e315f4d3341 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-cbb81deacfdecc97ef232d739e71879dee93c01187de6a35cf5b544464f11e75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a91d188f60582c8d3c99ecf53bedb3d4d4a437b57c04f0b0b944cd33380a8c97"></a>

## rule_list.rules.source_prefix_list — rule_list.rules.source_prefix_list / 15231edbbe56 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-8bac01521e851f7c57a71347c46cbaebc93bf3da961f8f78c751106655a22eaf)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- rule_list.rules.source_prefix_list

<a id="canonical-daab1d2cbe6cb6c0cd1ab8da89f3364ff9e3adb719b4416bc8f6f95899145395"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
source_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-27d383b603140f73ba22d3955bf0e50266546d93fb5f904a4f2f3e345dc29f2c"></a>

## Direct properties — rule_list.rules.source_prefix_list / 15231edbbe56 / 3

<a id="canonical-8f5c50164a133dd5dd5f7cd1eeba01b67aad5b61f9539cb7980d4c7b66975636"></a>

<a id="canonical-968a61230e21598f20e29cb0e2c6b8543fd2931031dc6f3669ed09b666a65252"></a>

## prefixes property — rule_list.rules.source_prefix_list / 15231edbbe56 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a9b6fbd28214777c9c51207c89ef8296eff6023e33684748cd39f7a10778fa64"></a>

## Next pages — rule_list.rules.source_prefix_list / 15231edbbe56 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-22d09bdceeacc008a576c814d152745ef0eac8fcba5a8b530e579228b5ca5306)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-507b69e00e66568c208f5606958c9ee2fb8c1eed5b5c2940ca5eb5b7549f3400"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00ec47cd553fe1e241134d46fc6fb7a1ab2f8e014270db9b29187f36007e2dad"></a>

## timeouts — timeouts / b3e2a097d23e / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- timeouts

<a id="canonical-725b7fc87126b784415a555ec1a8eaf4ecd5a221d6180b09dc1643e973c3a315"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-5698ee79d4ab60e546e9671a83106f6e401fec25860e6cffedbcd0721061e62f"></a>

## Direct properties — timeouts / b3e2a097d23e / 3

<a id="canonical-21c58c0f83589e0c0c298f49d04d34cf850cc6fada3d011a4cc82aafd20b5dc5"></a>

<a id="canonical-7be9e83ecf2a8d6dc6a5903b26f182abdc4727e622acbe9958034d815a37a6a4"></a>

## create property — timeouts / b3e2a097d23e / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-39af0e9e57ee9f48d5b01aac07f0ffa4a69dab9fc7dcaeaef4107af708969686"></a>

<a id="canonical-b7eb9f37137d95af8115c1cfce3b8cf64246728f2fd656d5160a6ddb31fd3cc1"></a>

## delete property — timeouts / b3e2a097d23e / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-e68e532451226b72249aff8ae6c66fcd2ca071267c846f563108bc15dbd928c1"></a>

<a id="canonical-78d8272ee1acb2c95d35a0bda2c24f9c626da9aa03719641b35343920ec64c36"></a>

## read property — timeouts / b3e2a097d23e / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-ab88d4accdd5cc66e5e3375d18acca1986c84e503b52dbf4fa360b42eafa545c"></a>

<a id="canonical-66779206cc97de48486443239a03e86cea83173ebdd0ac3d6c3a7537b241c3be"></a>

## update property — timeouts / b3e2a097d23e / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a7abaec4585d8a792d762e2e30d10d182bb5aea38473b9d58058ac2db5083b11"></a>

## Next pages — timeouts / b3e2a097d23e / 8

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
