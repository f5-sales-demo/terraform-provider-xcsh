---
page_title: "xcsh_network_policy_view reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view reference."
---

# xcsh_network_policy_view reference

<a id="canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b4f69ff20fb17d18c99f2ddf48e21ab5ddce3aa4b32569c8123d48cf1845d8d"></a>

## Property reference — Property reference / dff8fb964119 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- Property reference

<a id="canonical-9ca00472e9ee567232685a6e7f0d3876648a9fb5204102724403600086fa1e10"></a>

## Direct properties — Property reference / dff8fb964119 / 3

<a id="canonical-162f07811d3a69e6a36c193eee484dd77230af0e5cbc1855b8c34bcec25ce5b6"></a>

<a id="canonical-65a216228eb854c0dc293a6768d354439a805bc55f3ed0df1fab38ea6db83f61"></a>

## annotations property — Property reference / dff8fb964119 / 4

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

<a id="canonical-4622727cd330a49a20aa163d44e7f3cc020eb71e4718bd61d500801d05bcb22b"></a>

<a id="canonical-de82e0e0194f33d78458d59e91ac28e5647902990529515d0776f833a87a184a"></a>

## description property — Property reference / dff8fb964119 / 5

Type: `"string"`. Computed.

Description of the NetworkPolicyView.

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

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca): complete subsection reference.

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618): complete subsection reference.

<a id="canonical-66b56dc70c221d91ae83f083c9aad44d6a34286f21b2d671126200ce40b2f53a"></a>

<a id="canonical-080c60a7c14a487ff7660b5f318dc17585756524fe8d4c4a3ad52088a478de7c"></a>

## id property — Property reference / dff8fb964119 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec): complete subsection reference.

<a id="canonical-f1b0f668df2c4556775f435d152c4aeb88707eebc326fb955b04dc0c17545b13"></a>

<a id="canonical-77e388db648e50175282bf226f5f7954b4d1c69ec25603d94fcc30561b014d60"></a>

## labels property — Property reference / dff8fb964119 / 7

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

<a id="canonical-8da1a3be47f1c973760ed27e101c320549f54a6296e5ff4ae469f66e2a57510a"></a>

<a id="canonical-f5d8709c82457519290c5081aa0bafc24832d26fc8f3bd33462e36745bc8f4b3"></a>

## name property — Property reference / dff8fb964119 / 8

Type: `"string"`. Required.

Name of the NetworkPolicyView.

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

<a id="canonical-436c4d28256450e3c90b3bb648865678eb897a0fc5f9b09c9c0fa981feea7204"></a>

<a id="canonical-d9839bd469afb9dfcc2420af29e9abc881e248201221c35a1061fe6ab576914a"></a>

## namespace property — Property reference / dff8fb964119 / 9

Type: `"string"`. Optional, Computed.

Namespace where the NetworkPolicyView exists.

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

<a id="canonical-50364e4624c9f4713eeccf7af9f0481d4e280c8b4e1503ddfb8b93271cf63f93"></a>

## All schema paths — Property reference / dff8fb964119 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_policy_view--reference--group-001.md#canonical-162f07811d3a69e6a36c193eee484dd77230af0e5cbc1855b8c34bcec25ce5b6) |
| `description` | [description](data-sources--network_policy_view--reference--group-001.md#canonical-4622727cd330a49a20aa163d44e7f3cc020eb71e4718bd61d500801d05bcb22b) |
| `egress_rules` | [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-241aad1923c44e0c565a43f0f840ca358c23fe6866360b1e61962c9668018a4c) |
| `egress_rules.action` | [egress_rules.action](data-sources--network_policy_view--reference--group-001.md#canonical-13e3ca94b09052513126903e18594a835bf4b730d6c85644f15b5abd667791d6) |
| `egress_rules.adv_action` | [egress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-a790e60fbdbe6b09131935f204f5ea4f541cde2c5112086de8f8d5cc1ef1072c) |
| `egress_rules.adv_action.action` | [egress_rules.adv_action.action](data-sources--network_policy_view--reference--group-001.md#canonical-44a960189cb13281320ffe26ab24690027ad28c1c51650c41147bd24a5490e73) |
| `egress_rules.all_tcp_traffic` | [egress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-7753d7de7f94950ea954c2fba030addae7a44613d76bbbeb1c5b980b82cd8a78) |
| `egress_rules.all_traffic` | [egress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-1671be26e28c5ce0922992a96a6692b83bf87e2228a699101b3951f9c5711475) |
| `egress_rules.all_udp_traffic` | [egress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3a9e1a47a7dd8d3b2316ed2e122570f621207eef969cf3e5208c5483123f258f) |
| `egress_rules.any` | [egress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-a231850c23c2695ae6ea04a0afadb5bd6267243fb4133492abc584f0a470d011) |
| `egress_rules.applications` | [egress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-6b85944213e9ea8991a90b45c402cd342037cf31b88d09529a31a974fb123753) |
| `egress_rules.applications.applications` | [egress_rules.applications.applications](data-sources--network_policy_view--reference--group-001.md#canonical-90be7f8d58dd6b37e3a3464b8b36c8bd002e8044dca26248a961d412373ed4ee) |
| `egress_rules.inside_endpoints` | [egress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-34cb5edcb56142801f4afee2042fc95013299ea990c94ff75739c533284ed205) |
| `egress_rules.ip_prefix_set` | [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-7d0793200c67c616b9e4a24d4d14a370f2cdafe7522be4efe5e736530bad5d6a) |
| `egress_rules.ip_prefix_set.ref` | [egress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-df60cd0bbe247cdd8dd0306756f65224cd665ef11829413a4b75f0ab11567649) |
| `egress_rules.ip_prefix_set.ref.kind` | [egress_rules.ip_prefix_set.ref.kind](data-sources--network_policy_view--reference--group-001.md#canonical-945e3a83ea13c86f46275060d9d9c83e834e2fffedb2a787f19efc836f4a7e52) |
| `egress_rules.ip_prefix_set.ref.name` | [egress_rules.ip_prefix_set.ref.name](data-sources--network_policy_view--reference--group-001.md#canonical-e5071663cbd8fc2097422c65f99433d3465e6f240078ae9c84d8d643ba8c67dc) |
| `egress_rules.ip_prefix_set.ref.namespace` | [egress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy_view--reference--group-001.md#canonical-0cb2f5e933e9830054af06482e61ba9423a22080e1b3fe3e77608310ea2efe56) |
| `egress_rules.ip_prefix_set.ref.tenant` | [egress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy_view--reference--group-001.md#canonical-5b7b34de259127d3c3248438c04bd8541384fb235648c6d3c8efae6a00af2ffc) |
| `egress_rules.ip_prefix_set.ref.uid` | [egress_rules.ip_prefix_set.ref.uid](data-sources--network_policy_view--reference--group-001.md#canonical-8b7ac39d4fc67a32c70d0317eab987c46f6d99ff4e0592f0457573d1ba91d320) |
| `egress_rules.label_matcher` | [egress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-5b4fc5afdb180ed5dae02cbb530c64f333dc8d84c37cac324121720851440e10) |
| `egress_rules.label_matcher.keys` | [egress_rules.label_matcher.keys](data-sources--network_policy_view--reference--group-001.md#canonical-eff3cbf1357bb37205b5cf6edd0855aaa5b4078dd6953e201768d2928dfe4f95) |
| `egress_rules.label_selector` | [egress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-965f23bebe1f77fee33b8f825f566463b860c4f0615c31a317947c9e20e64eec) |
| `egress_rules.label_selector.expressions` | [egress_rules.label_selector.expressions](data-sources--network_policy_view--reference--group-001.md#canonical-3e049f1933878505b6df6a1dfd2f4aaa8ee24c5ad042c0940f067fec7b50294d) |
| `egress_rules.metadata` | [egress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-0f81b7ba46506111a4965e0162d10b99bb4e1ed02eca667a3b486b978eef529e) |
| `egress_rules.metadata.description_spec` | [egress_rules.metadata.description_spec](data-sources--network_policy_view--reference--group-001.md#canonical-220ca34e1df08cd4f78f64bf136ef26b7930a0ea88b5dda2b89484f8953c1f55) |
| `egress_rules.metadata.name` | [egress_rules.metadata.name](data-sources--network_policy_view--reference--group-001.md#canonical-ba65536dfa57e196bd938cca647b53f8fc4b7858cc4346c71de292d65d257ab8) |
| `egress_rules.outside_endpoints` | [egress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-9aba30228e71dc54c20d5e452938defe4aedc0718550a69f0e3a1dac384f9aa3) |
| `egress_rules.prefix_list` | [egress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-917cd3b2fe6d4fdada8cff5a871662f2ab23f856a550edb75f202a2493fa7bc8) |
| `egress_rules.prefix_list.prefixes` | [egress_rules.prefix_list.prefixes](data-sources--network_policy_view--reference--group-001.md#canonical-ac5b4a793cc511de7d404d29c096e1a76a8371607a6652e5967fa396b938ddb6) |
| `egress_rules.protocol_port_range` | [egress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-169bf78850740e2274d65b8a06b0ce173cd2187639a7aeed32ba7d7fc206ec2f) |
| `egress_rules.protocol_port_range.port_ranges` | [egress_rules.protocol_port_range.port_ranges](data-sources--network_policy_view--reference--group-001.md#canonical-2d56c8bb3beac748aa9a4f719cde092b2f17e7a130ae9e6c8b013d9b9dae1aef) |
| `egress_rules.protocol_port_range.protocol` | [egress_rules.protocol_port_range.protocol](data-sources--network_policy_view--reference--group-001.md#canonical-c56731af2355ef50228a850c7d53b34ec4aecf6e3521019871eef751dad4506a) |
| `endpoint` | [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-9fc117fd858f033c30558a1e57ac340c3938e769ee840d41ce09eb5b3df74367) |
| `endpoint.any` | [endpoint.any](data-sources--network_policy_view--reference--group-001.md#canonical-03ce34a677408971c782d4de18da43070cff038f40bb572b96ef74d5198d37ce) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-8150987b59945e9a1378657f64b7e653ad3d87c179bcbf85dd387d6e73205dfa) |
| `endpoint.label_selector` | [endpoint.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-e0cd3c54f2b54fb2fd0141a2b8def438eece02326ec7aad655dc789a03de94b4) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](data-sources--network_policy_view--reference--group-001.md#canonical-58f146997c7581013e8188cc6c1f19a4ae105ecdb276c15fe6b56975a5d90162) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-81656058b4a4810ed7bfdfb27fec2b9556a120bac8d366e1ac699eb65f505b3a) |
| `endpoint.prefix_list` | [endpoint.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-02c33c1accf4c3699cb4e3a9de5b86f6d6194c8f52fcbc8b386ab2c8efeb4acd) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](data-sources--network_policy_view--reference--group-001.md#canonical-2e54b97c8c18ac6684cd4d687aa4b4b42d5bc851123d2c4988e685693f659246) |
| `id` | [id](data-sources--network_policy_view--reference--group-001.md#canonical-66b56dc70c221d91ae83f083c9aad44d6a34286f21b2d671126200ce40b2f53a) |
| `ingress_rules` | [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-d73b6c96cfde7400d1a4825ff20e5c1b7082c17a7964d334966e686148e53eac) |
| `ingress_rules.action` | [ingress_rules.action](data-sources--network_policy_view--reference--group-001.md#canonical-2b1e6134a7e7aa5616794a7e68f5f701d42695a866a25a8b0028671bfc98a89a) |
| `ingress_rules.adv_action` | [ingress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-b42882a1755276de991f7ca5143941770cec9a63d12a1bc2b474a228f6c41660) |
| `ingress_rules.adv_action.action` | [ingress_rules.adv_action.action](data-sources--network_policy_view--reference--group-001.md#canonical-15423e7d21a45f52ce4a8cd7a0ab57bdf20335933cffc637e3d2f234ef731028) |
| `ingress_rules.all_tcp_traffic` | [ingress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-690ca9abecdb3ec7ce40a33993d3c5ba870c7f6bd987aac04133098eb7889622) |
| `ingress_rules.all_traffic` | [ingress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-5ea8a450ee0902298b29bbc245788385d1eb8eae9107f6189933800ecaf69b32) |
| `ingress_rules.all_udp_traffic` | [ingress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0c4b8026e6807f16a02623031d756b030ad10bb26589b7e94f5deb090fd44fe0) |
| `ingress_rules.any` | [ingress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-5a389005983de25fb28de257c7d60b1c681dc032cc1bd2ecb163f9eabe375496) |
| `ingress_rules.applications` | [ingress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-d73177f1ea41684be6edd35794e55ff31d07aa0bda7fd9cf439f82f4dc0adab2) |
| `ingress_rules.applications.applications` | [ingress_rules.applications.applications](data-sources--network_policy_view--reference--group-001.md#canonical-bbf52bb5a43f7a67b425966eda123313216331c973a62d6d82557efa0aa54d6d) |
| `ingress_rules.inside_endpoints` | [ingress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-7599af7b7e8d3646578bf98f5502f754d1eb12fe7a3fa640638fdf808e482fb0) |
| `ingress_rules.ip_prefix_set` | [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-fde060c9b4236fccb1fa1c73929919a7e8519fd69b6932bd7296cfbe05d4a83b) |
| `ingress_rules.ip_prefix_set.ref` | [ingress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-15c96c070378e95998a1c5d616fb659f9c9500244ebfa2af7edf6d77c35c5582) |
| `ingress_rules.ip_prefix_set.ref.kind` | [ingress_rules.ip_prefix_set.ref.kind](data-sources--network_policy_view--reference--group-001.md#canonical-af5b5dca4063dff126a4bc820703f236dadb2f759c454c124ca43d9d9020ca05) |
| `ingress_rules.ip_prefix_set.ref.name` | [ingress_rules.ip_prefix_set.ref.name](data-sources--network_policy_view--reference--group-001.md#canonical-3afbd449fd22ba2aefb881608fc0eef4c839f276e83cac9f7ec39c7214f978be) |
| `ingress_rules.ip_prefix_set.ref.namespace` | [ingress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy_view--reference--group-001.md#canonical-cc43883b84cda2c713983ae11dcd447197c5f924126e92880290025f70c94256) |
| `ingress_rules.ip_prefix_set.ref.tenant` | [ingress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy_view--reference--group-001.md#canonical-c0fc609e4baf571171ae143ad0b07b6581f1fd22001220c3bb659c33c432610e) |
| `ingress_rules.ip_prefix_set.ref.uid` | [ingress_rules.ip_prefix_set.ref.uid](data-sources--network_policy_view--reference--group-001.md#canonical-19bf2c858bc592ad3d0286504fba634013e8cada04dea00e906559052eae228a) |
| `ingress_rules.label_matcher` | [ingress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-d202640fcc7b5d634641cf712ffc8887a09a14e89cf886d47952fe6bde85fcf7) |
| `ingress_rules.label_matcher.keys` | [ingress_rules.label_matcher.keys](data-sources--network_policy_view--reference--group-001.md#canonical-459ef1356f3ef99cd4e683bfeadd1b78bd6dce36b8d341fbdcba0b0b4ab4d30c) |
| `ingress_rules.label_selector` | [ingress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-942408488dd008726390c11046266d0430101d1f70818f41203d9f539c773f32) |
| `ingress_rules.label_selector.expressions` | [ingress_rules.label_selector.expressions](data-sources--network_policy_view--reference--group-001.md#canonical-3239945ccb1abb4ff151511d62382ea86a74563bc49e2d4bf4549b118a58821b) |
| `ingress_rules.metadata` | [ingress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-9f77aff7530895024e8ef6c990da9c7a0a827d42f0e8262dc9dc1b27f6396bfd) |
| `ingress_rules.metadata.description_spec` | [ingress_rules.metadata.description_spec](data-sources--network_policy_view--reference--group-001.md#canonical-7e87efce1850edf87be48ffe1d35c9b616b8b666d434d20846d383795f6b6dcb) |
| `ingress_rules.metadata.name` | [ingress_rules.metadata.name](data-sources--network_policy_view--reference--group-001.md#canonical-8d55e1163859ccb2a1725b897240df63fc8f0022cd3da287b6b07dbf73cbe7df) |
| `ingress_rules.outside_endpoints` | [ingress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-c53e021c424979978413d038908e0bae65850656d0fccde9cb111eea46b8ed0c) |
| `ingress_rules.prefix_list` | [ingress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-7bcece0f7573d318268bf7648075d88aa843d961673ee1de2163d6b4a713b28f) |
| `ingress_rules.prefix_list.prefixes` | [ingress_rules.prefix_list.prefixes](data-sources--network_policy_view--reference--group-001.md#canonical-85ddd6497a11b08a6d8c5cdfb1ead328a913dd6bb1b24b067ca35c171dc85215) |
| `ingress_rules.protocol_port_range` | [ingress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-ae691ab5baa92d2ca446c272f8941e9b20720ce358d3402488d50baca6ba3b0b) |
| `ingress_rules.protocol_port_range.port_ranges` | [ingress_rules.protocol_port_range.port_ranges](data-sources--network_policy_view--reference--group-001.md#canonical-19942d874b5d2f8a91baf0910a07ec88180167debe0a193371e083aabdc2a41a) |
| `ingress_rules.protocol_port_range.protocol` | [ingress_rules.protocol_port_range.protocol](data-sources--network_policy_view--reference--group-001.md#canonical-aca65bef7c59f8e8ef2377385b69ecdef9570489144642e985481f2f9ae27101) |
| `labels` | [labels](data-sources--network_policy_view--reference--group-001.md#canonical-f1b0f668df2c4556775f435d152c4aeb88707eebc326fb955b04dc0c17545b13) |
| `name` | [name](data-sources--network_policy_view--reference--group-001.md#canonical-8da1a3be47f1c973760ed27e101c320549f54a6296e5ff4ae469f66e2a57510a) |
| `namespace` | [namespace](data-sources--network_policy_view--reference--group-001.md#canonical-436c4d28256450e3c90b3bb648865678eb897a0fc5f9b09c9c0fa981feea7204) |

<a id="canonical-c3c3746bd8da526b97effd126a174acdcd6d9bcbcdf5a5a05832146f4df6bf11"></a>

## Next pages — Property reference / dff8fb964119 / 11

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7376c014204cfef4df85fe6b4fe496ee274872d9423b20c26fd75d565ee3dbd1"></a>

## egress_rules — egress_rules / 879837e5de54 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- egress_rules

<a id="canonical-241aad1923c44e0c565a43f0f840ca358c23fe6866360b1e61962c9668018a4c"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections from policy endpoints.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-bf5cf0f9afc4320d5d8b5b45f11ed6d7dc46bff9dc122b8207302ec7fdbd0831"></a>

## Direct properties — egress_rules / 879837e5de54 / 3

<a id="canonical-13e3ca94b09052513126903e18594a835bf4b730d6c85644f15b5abd667791d6"></a>

<a id="canonical-764b55514f9b9828cd28e36c8023aa2ab9689cd61054dbff605f6c4e8f17174d"></a>

## action property — egress_rules / 879837e5de54 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-cf200590a98ddd50344f5e57307c793324052d397e743dc6a0d36047283f2fbe): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-dda340993f8a3cd0282187fce172be5eeb7af6e639b64dfa84e1ab7a35bb4c47): complete subsection reference.

- [all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-cec185fbf45a3c52d40c048a5bdc85f339af9b1383c3cc85619fc4f56105e5b1): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-20d67ff97cbbf690ec46b655ea0e5dc079e5628d040f5de47626c43732283dbf): complete subsection reference.

- [any](data-sources--network_policy_view--reference--group-001.md#canonical-7e5ad4599fb519910a2b7db0a80c70159d04b54f0daf811360d1424f09329f3b): complete subsection reference.

- [applications](data-sources--network_policy_view--reference--group-001.md#canonical-2ce34d76af010f9fbeb41f669dbe9a0a315eb799e16614d573c1d1fe546a6ed6): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-d2d8529c358fa58a1506b02610a1f59f21333ceaf762e59fd749171fa65deff2): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-b045796a3963401730125fbf11f538e071b3d67f3df06c00c4321b8739b148bb): complete subsection reference.

- [label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-3b74ac4ce04a1fe653f21769dcd546f3bdb918ed7f52e978209c46dce306645c): complete subsection reference.

- [label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-b235aa8859847557e30a5d132b8fab5db789d2ee80f2542843707b8a29703bb8): complete subsection reference.

- [metadata](data-sources--network_policy_view--reference--group-001.md#canonical-7aeaa8936d49381de489e4297bc58a86cd9e6a7a5ea6c2588cb13eb97fd01736): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-4f0f731763cd92331cbe865821359d03b27022d5a8c0ce306ee99d1eb8cf7f57): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-702d791676d0231007304e65c02a438288c0d7a5b0599e7c7d215efe1aa7dbfa): complete subsection reference.

- [protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-6a2ae796a854076542d60f79eeb302cca3052c4dd8ca7a4f1116f14dc650bf72): complete subsection reference.

<a id="canonical-b641b83b5e9c4de041928d79f7379ab330b1fc7155f97ec8a2179b0384cf2ef1"></a>

## Next pages — egress_rules / 879837e5de54 / 5

- [egress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-cf200590a98ddd50344f5e57307c793324052d397e743dc6a0d36047283f2fbe)
- [egress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-dda340993f8a3cd0282187fce172be5eeb7af6e639b64dfa84e1ab7a35bb4c47)
- [egress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-cec185fbf45a3c52d40c048a5bdc85f339af9b1383c3cc85619fc4f56105e5b1)
- [egress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-20d67ff97cbbf690ec46b655ea0e5dc079e5628d040f5de47626c43732283dbf)
- [egress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-7e5ad4599fb519910a2b7db0a80c70159d04b54f0daf811360d1424f09329f3b)
- [egress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-2ce34d76af010f9fbeb41f669dbe9a0a315eb799e16614d573c1d1fe546a6ed6)
- [egress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-d2d8529c358fa58a1506b02610a1f59f21333ceaf762e59fd749171fa65deff2)
- [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-b045796a3963401730125fbf11f538e071b3d67f3df06c00c4321b8739b148bb)
- [egress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-3b74ac4ce04a1fe653f21769dcd546f3bdb918ed7f52e978209c46dce306645c)
- [egress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-b235aa8859847557e30a5d132b8fab5db789d2ee80f2542843707b8a29703bb8)
- [egress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-7aeaa8936d49381de489e4297bc58a86cd9e6a7a5ea6c2588cb13eb97fd01736)
- [egress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-4f0f731763cd92331cbe865821359d03b27022d5a8c0ce306ee99d1eb8cf7f57)
- [egress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-702d791676d0231007304e65c02a438288c0d7a5b0599e7c7d215efe1aa7dbfa)
- [egress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-6a2ae796a854076542d60f79eeb302cca3052c4dd8ca7a4f1116f14dc650bf72)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-cf200590a98ddd50344f5e57307c793324052d397e743dc6a0d36047283f2fbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c17e9407a508259c94f5a10053d76ca4a57a58a72319dc3f0db25d20e90f5a6"></a>

## egress_rules.adv_action — egress_rules.adv_action / 6d021650f672 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.adv_action

<a id="canonical-a790e60fbdbe6b09131935f204f5ea4f541cde2c5112086de8f8d5cc1ef1072c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-6f38159e03e873bcc159f3f4bfa46c0f752cae396af51ddc49eddf9fd7e6af86"></a>

## Direct properties — egress_rules.adv_action / 6d021650f672 / 3

<a id="canonical-44a960189cb13281320ffe26ab24690027ad28c1c51650c41147bd24a5490e73"></a>

<a id="canonical-6f046b72c01899e86845f2e0d290328dc311c60516707bf417dba1664907f446"></a>

## action property — egress_rules.adv_action / 6d021650f672 / 4

Type: `"string"`. Computed.

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

<a id="canonical-f4b0a9caf118c458309f5a77189bcdb4ed8b67eb76f81e6863028f73655d0fd8"></a>

## Next pages — egress_rules.adv_action / 6d021650f672 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-dda340993f8a3cd0282187fce172be5eeb7af6e639b64dfa84e1ab7a35bb4c47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc90baf250839a4e697f930b40fbedfc7e3d353c6945f5da513334b4281c02c7"></a>

## egress_rules.all_tcp_traffic — egress_rules.all_tcp_traffic / d452e7ca6e3d / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.all_tcp_traffic

<a id="canonical-7753d7de7f94950ea954c2fba030addae7a44613d76bbbeb1c5b980b82cd8a78"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-bdf655eab89d62cea2ba52835826de9fd2a25b13b4842c487c6d741521412441"></a>

## Direct properties — egress_rules.all_tcp_traffic / d452e7ca6e3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d431a4d6e80fcd0dfb42354b0ef86448cc55af51c227d9237f5d541f42cc94b4"></a>

## Next pages — egress_rules.all_tcp_traffic / d452e7ca6e3d / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-cec185fbf45a3c52d40c048a5bdc85f339af9b1383c3cc85619fc4f56105e5b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eef676d8b5dddfb2bfc246ce6d01f52b213ee93b87fae158019e94c35fec3ce5"></a>

## egress_rules.all_traffic — egress_rules.all_traffic / aed9b5999bda / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.all_traffic

<a id="canonical-1671be26e28c5ce0922992a96a6692b83bf87e2228a699101b3951f9c5711475"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0059b96a60b1dea6cc93e87214b7450a78fe4e20458bdb1dc9ea5201c364d5ff"></a>

## Direct properties — egress_rules.all_traffic / aed9b5999bda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68e5831d349574ee2677c98bfe8e044d684e49bb58eb2ee9d56b6dab2ce8fbd8"></a>

## Next pages — egress_rules.all_traffic / aed9b5999bda / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-20d67ff97cbbf690ec46b655ea0e5dc079e5628d040f5de47626c43732283dbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66ddf46f819120f35138bba12f5e6ba18bbcce640425922106b56d89e3a3b3fe"></a>

## egress_rules.all_udp_traffic — egress_rules.all_udp_traffic / db785a9da51b / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.all_udp_traffic

<a id="canonical-3a9e1a47a7dd8d3b2316ed2e122570f621207eef969cf3e5208c5483123f258f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d4b5903405aef230751c9b1ec82aebcf81bb687ea2ba99416c354224cb3e9d4e"></a>

## Direct properties — egress_rules.all_udp_traffic / db785a9da51b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57326ea7c0761aa85fc75d1915087f6f797cbce51d0ef15b15364370236a3d51"></a>

## Next pages — egress_rules.all_udp_traffic / db785a9da51b / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-7e5ad4599fb519910a2b7db0a80c70159d04b54f0daf811360d1424f09329f3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6cae8e79cd0065ee7a8015680a5747eb942100a5673f79e889fa8aa6636ce64"></a>

## egress_rules.any — egress_rules.any / 068ceb40cdcd / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.any

<a id="canonical-a231850c23c2695ae6ea04a0afadb5bd6267243fb4133492abc584f0a470d011"></a>

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

<a id="canonical-85c876b1416d7a6ac415ef9b0f972b8a0a1e651c3f66e8f7cb35d51fc0124428"></a>

## Direct properties — egress_rules.any / 068ceb40cdcd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-600e4b3b62633524f557c10e50aa4c316b06b618b6be96ea4262f458c82b78d7"></a>

## Next pages — egress_rules.any / 068ceb40cdcd / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-2ce34d76af010f9fbeb41f669dbe9a0a315eb799e16614d573c1d1fe546a6ed6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e84c5b7b66ea11482d6339213da4e02a3ca69306ed955b6c5a7544fe565457a4"></a>

## egress_rules.applications — egress_rules.applications / e2f871ef80b9 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.applications

<a id="canonical-6b85944213e9ea8991a90b45c402cd342037cf31b88d09529a31a974fb123753"></a>

Type: `"single"`. Computed.

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

<a id="canonical-18f45365990d924cf915e5396ecd89c07afeb3c029d6dfe1cb6887eb79a6435d"></a>

## Direct properties — egress_rules.applications / e2f871ef80b9 / 3

<a id="canonical-90be7f8d58dd6b37e3a3464b8b36c8bd002e8044dca26248a961d412373ed4ee"></a>

<a id="canonical-442eb1c222be8d2aeb6fb578f3981d792d47e3d09c336bd72cfac4d64f6e0c93"></a>

## applications property — egress_rules.applications / e2f871ef80b9 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-828293ede4d873fc7ba598f672d990161a6d332ac602b30085d5f2d58f24dcbe"></a>

## Next pages — egress_rules.applications / e2f871ef80b9 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-d2d8529c358fa58a1506b02610a1f59f21333ceaf762e59fd749171fa65deff2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61ab8f6c04cd1cecedf1dcd7e609a619ac9e106c54df292f704900c30ea63807"></a>

## egress_rules.inside_endpoints — egress_rules.inside_endpoints / da77bd0acaa5 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.inside_endpoints

<a id="canonical-34cb5edcb56142801f4afee2042fc95013299ea990c94ff75739c533284ed205"></a>

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

<a id="canonical-49a705ad7fa37d51cc0b45e71b1c49e21b401bdce27130fe31cce652b6e930da"></a>

## Direct properties — egress_rules.inside_endpoints / da77bd0acaa5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80193d46cc24f5b7448a29fe24922d0ae62fd50022d1980d72c9970b23633771"></a>

## Next pages — egress_rules.inside_endpoints / da77bd0acaa5 / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-b045796a3963401730125fbf11f538e071b3d67f3df06c00c4321b8739b148bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c196352efe95a9989bb277b347303078637562bacf6975e1286a495ad2f99634"></a>

## egress_rules.ip_prefix_set — egress_rules.ip_prefix_set / b4d7b5865310 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.ip_prefix_set

<a id="canonical-7d0793200c67c616b9e4a24d4d14a370f2cdafe7522be4efe5e736530bad5d6a"></a>

Type: `"single"`. Computed.

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

<a id="canonical-dbac6933db685570ef91b74889be1171dbf171e508c4f423963115af8fc4e560"></a>

## Direct properties — egress_rules.ip_prefix_set / b4d7b5865310 / 3

- [ref](data-sources--network_policy_view--reference--group-001.md#canonical-1a38f9e7bcb6fd83289cd6e0fb656e3452dac86b31035e4de3549bcb6b63786d): complete subsection reference.

<a id="canonical-63adddb95c213743ebad56592d1641063249b89d7a1ededb00a9de4a35da3609"></a>

## Next pages — egress_rules.ip_prefix_set / b4d7b5865310 / 4

- [egress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-1a38f9e7bcb6fd83289cd6e0fb656e3452dac86b31035e4de3549bcb6b63786d)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-1a38f9e7bcb6fd83289cd6e0fb656e3452dac86b31035e4de3549bcb6b63786d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-193a83f1886e5db7f082020d85e819fd090cb60b783c2610b29d570e0799607b"></a>

## egress_rules.ip_prefix_set.ref — egress_rules.ip_prefix_set.ref / b87de347195c / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-b045796a3963401730125fbf11f538e071b3d67f3df06c00c4321b8739b148bb)
- egress_rules.ip_prefix_set.ref

<a id="canonical-df60cd0bbe247cdd8dd0306756f65224cd665ef11829413a4b75f0ab11567649"></a>

Type: `"list"`. Computed.

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

<a id="canonical-4e82218313d4ae3ca687681be1879ffe1aace1506fbd214e1c9ce365e37e2f42"></a>

## Direct properties — egress_rules.ip_prefix_set.ref / b87de347195c / 3

<a id="canonical-945e3a83ea13c86f46275060d9d9c83e834e2fffedb2a787f19efc836f4a7e52"></a>

<a id="canonical-cd51d7a5db28cb7ff187e2f1d6bc8c3078acd62cad4dcbcd3dd6205ba37d8ace"></a>

## kind property — egress_rules.ip_prefix_set.ref / b87de347195c / 4

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

<a id="canonical-e5071663cbd8fc2097422c65f99433d3465e6f240078ae9c84d8d643ba8c67dc"></a>

<a id="canonical-8575b57f7173e2eb0e51c17fd97ac10be3f0b772ebe30a56930483391a544e38"></a>

## name property — egress_rules.ip_prefix_set.ref / b87de347195c / 5

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

<a id="canonical-0cb2f5e933e9830054af06482e61ba9423a22080e1b3fe3e77608310ea2efe56"></a>

<a id="canonical-388e5287399b06d304ea77e7f07118e7e6d1a3242e1ee1762664bbe4805528dc"></a>

## namespace property — egress_rules.ip_prefix_set.ref / b87de347195c / 6

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

<a id="canonical-5b7b34de259127d3c3248438c04bd8541384fb235648c6d3c8efae6a00af2ffc"></a>

<a id="canonical-4bcd22795a6b8083e4da3683dd68cf91f86d7124be60dcb585c21f17890d5f2d"></a>

## tenant property — egress_rules.ip_prefix_set.ref / b87de347195c / 7

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

<a id="canonical-8b7ac39d4fc67a32c70d0317eab987c46f6d99ff4e0592f0457573d1ba91d320"></a>

<a id="canonical-4eb866dcf0395319ed94b2e0ce1dbcb644257c0ec5b9fcb32d0f53552718b7f6"></a>

## uid property — egress_rules.ip_prefix_set.ref / b87de347195c / 8

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

<a id="canonical-93f37422497794dd836792f07800d5577f648207aa2f67458cd1f7b6f7e82b2c"></a>

## Next pages — egress_rules.ip_prefix_set.ref / b87de347195c / 9

- [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-b045796a3963401730125fbf11f538e071b3d67f3df06c00c4321b8739b148bb)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-3b74ac4ce04a1fe653f21769dcd546f3bdb918ed7f52e978209c46dce306645c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87e23b1eb3e05d3fa1cbbe069281e4f3d8fb66885086bdbfc8f2ccbd328e8d8e"></a>

## egress_rules.label_matcher — egress_rules.label_matcher / e5481881116f / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.label_matcher

<a id="canonical-5b4fc5afdb180ed5dae02cbb530c64f333dc8d84c37cac324121720851440e10"></a>

Type: `"single"`. Computed.

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

<a id="canonical-9a0f4114168353f829fdc17d82c9ad631dba6fb231cddb154b71d61b688ae37e"></a>

## Direct properties — egress_rules.label_matcher / e5481881116f / 3

<a id="canonical-eff3cbf1357bb37205b5cf6edd0855aaa5b4078dd6953e201768d2928dfe4f95"></a>

<a id="canonical-f46b7eea32f5cf291b67ad7eb093b46442941899d93a2c53c7829eef2d1106c5"></a>

## keys property — egress_rules.label_matcher / e5481881116f / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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

<a id="canonical-aef58fbe5a5d2c12171198b2e67fe99559fc936d5482b435404ce467fde7624e"></a>

## Next pages — egress_rules.label_matcher / e5481881116f / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-b235aa8859847557e30a5d132b8fab5db789d2ee80f2542843707b8a29703bb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b35521eda55fda182603028c45cd59134e38906cd98d6edfeb3c0ca63004ff2b"></a>

## egress_rules.label_selector — egress_rules.label_selector / eb704867a939 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.label_selector

<a id="canonical-965f23bebe1f77fee33b8f825f566463b860c4f0615c31a317947c9e20e64eec"></a>

Type: `"single"`. Computed.

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

<a id="canonical-b7124dc5a85c17d284ee695f48eb732bddc212417c40243a9f0881d6adcb49f3"></a>

## Direct properties — egress_rules.label_selector / eb704867a939 / 3

<a id="canonical-3e049f1933878505b6df6a1dfd2f4aaa8ee24c5ad042c0940f067fec7b50294d"></a>

<a id="canonical-6005edd6ce1129854df623f9be59115d8b842ed61d8e391f4bd9fb1955fc9a81"></a>

## expressions property — egress_rules.label_selector / eb704867a939 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-a997e6ae1bc5cd7a8b16ad10ffa462672d85a62122e3eb568141e69b6f55bc6a"></a>

## Next pages — egress_rules.label_selector / eb704867a939 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-7aeaa8936d49381de489e4297bc58a86cd9e6a7a5ea6c2588cb13eb97fd01736"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d4ca9f183643ac2cf23e99293d9044f1b94df176bde201fae2e8389e8b5b568"></a>

## egress_rules.metadata — egress_rules.metadata / 69aabddcd4b6 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.metadata

<a id="canonical-0f81b7ba46506111a4965e0162d10b99bb4e1ed02eca667a3b486b978eef529e"></a>

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

<a id="canonical-f5de99d3854b4d429b0c8383a316c7ba2e8280b49038db846417d8376af6b5be"></a>

## Direct properties — egress_rules.metadata / 69aabddcd4b6 / 3

<a id="canonical-220ca34e1df08cd4f78f64bf136ef26b7930a0ea88b5dda2b89484f8953c1f55"></a>

<a id="canonical-f359e6fb91ec989a77fd28e87e4dd3547098214cf0ed14683b176c0e0cd11f51"></a>

## description_spec property — egress_rules.metadata / 69aabddcd4b6 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-ba65536dfa57e196bd938cca647b53f8fc4b7858cc4346c71de292d65d257ab8"></a>

<a id="canonical-fa92a004898e61b7a5728fe26bf322aa5672d3941ba052d534822fb976a8ee10"></a>

## name property — egress_rules.metadata / 69aabddcd4b6 / 5

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

<a id="canonical-d832caa5453069524bd7b07c7f1ad42e4b1b208acd0196c1969a4a0ad7020f7b"></a>

## Next pages — egress_rules.metadata / 69aabddcd4b6 / 6

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-4f0f731763cd92331cbe865821359d03b27022d5a8c0ce306ee99d1eb8cf7f57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-737947352818299764e65d2a41d5793a652927d8a45baa66dada68d64a4a0849"></a>

## egress_rules.outside_endpoints — egress_rules.outside_endpoints / 9425e711f10d / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.outside_endpoints

<a id="canonical-9aba30228e71dc54c20d5e452938defe4aedc0718550a69f0e3a1dac384f9aa3"></a>

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

<a id="canonical-d319f95e36bd7c890a61c5c01355b5618acc47d5c0bfc6d68e477ca6a4351638"></a>

## Direct properties — egress_rules.outside_endpoints / 9425e711f10d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fff2aac8e5cee2e886282c3c5f23a0b4881c9a8df83729ce9a653a78a58496ae"></a>

## Next pages — egress_rules.outside_endpoints / 9425e711f10d / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-702d791676d0231007304e65c02a438288c0d7a5b0599e7c7d215efe1aa7dbfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad6cd59de9aa872a122be6c095482636cb188a5f86c2b4e65e7db922d1237ad4"></a>

## egress_rules.prefix_list — egress_rules.prefix_list / d6cb7837cb2e / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.prefix_list

<a id="canonical-917cd3b2fe6d4fdada8cff5a871662f2ab23f856a550edb75f202a2493fa7bc8"></a>

Type: `"single"`. Computed.

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

<a id="canonical-45acb41e91d9f98fef2e66ae58fec5639f8d7c80a1c7ce424d5d057494620687"></a>

## Direct properties — egress_rules.prefix_list / d6cb7837cb2e / 3

<a id="canonical-ac5b4a793cc511de7d404d29c096e1a76a8371607a6652e5967fa396b938ddb6"></a>

<a id="canonical-887abecc297e5ee70a9014484e11cf9ca8d5651c0b3d2ecdb4980df32494336b"></a>

## prefixes property — egress_rules.prefix_list / d6cb7837cb2e / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-3a1b73f8b277de28c0cd0f2760a9e80408d7c6c58c6fa1bae04890d18a248c4a"></a>

## Next pages — egress_rules.prefix_list / d6cb7837cb2e / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-6a2ae796a854076542d60f79eeb302cca3052c4dd8ca7a4f1116f14dc650bf72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3250ba162eb344c74f031706e8f2fe008ada183eebb3e037cd135d0648447e08"></a>

## egress_rules.protocol_port_range — egress_rules.protocol_port_range / 2af34ba6e827 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- egress_rules.protocol_port_range

<a id="canonical-169bf78850740e2274d65b8a06b0ce173cd2187639a7aeed32ba7d7fc206ec2f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-fba5b2c2255595366d4aae62b6b358d00dc7be4d9748bff4b6921049f33c5fe1"></a>

## Direct properties — egress_rules.protocol_port_range / 2af34ba6e827 / 3

<a id="canonical-2d56c8bb3beac748aa9a4f719cde092b2f17e7a130ae9e6c8b013d9b9dae1aef"></a>

<a id="canonical-f887329bf4ec9c080c0d2b8e7196e35b1d169918c57d52db5aa04cca2de65f47"></a>

## port_ranges property — egress_rules.protocol_port_range / 2af34ba6e827 / 4

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

<a id="canonical-c56731af2355ef50228a850c7d53b34ec4aecf6e3521019871eef751dad4506a"></a>

<a id="canonical-94e06676470362cf5bbffd7e0e4d3844988522032f7beb9d7dd27c4bcfdc1223"></a>

## protocol property — egress_rules.protocol_port_range / 2af34ba6e827 / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

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

<a id="canonical-059142999489e00a26c4c9647ce1531c42bfc75a82b105a8ebfd91f0b4d38613"></a>

## Next pages — egress_rules.protocol_port_range / 2af34ba6e827 / 6

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-6bd2ed2376a32e0242bee0c33fdf15de0198a685466a96825039cf6ccfaad1ca)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97bc98801b367053131afc600e5e57ab5c301aa550425e022137352c12bf394e"></a>

## endpoint — endpoint / 7e87374a895d / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- endpoint

<a id="canonical-9fc117fd858f033c30558a1e57ac340c3938e769ee840d41ce09eb5b3df74367"></a>

Type: `"single"`. Computed.

Shape of the endpoint choices for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"any\",\"inside_endpoints\",\"label_selector\",\"outside_endpoints\",\"prefix_list\"]"
}
```

<a id="canonical-7dd59e5b241fa9c911737074b8a881ca6ff10283d0002fcedc4be7ae920ebbd5"></a>

## Direct properties — endpoint / 7e87374a895d / 3

- [any](data-sources--network_policy_view--reference--group-001.md#canonical-1dff2fd6b4675267a545fbcda84dd656344dd7e2fca50eed177c026c1eb3f3ff): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-a2e63320ca8ba04edb81c8ede5d9fefdad0fe0123b4cbfa80e7a7d37327b825a): complete subsection reference.

- [label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-4cec9b09377827e68fa86887897d016764cca3cb69327ba673505f222bcfa844): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-e85148ee88f957b182dbfb55bce8e5b55a6aa9aa2584b46a11d987c11ca771ba): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-72480f3ed28046057f77dd0e438b0fac5ba7dca5a51482d248c0b12de5604834): complete subsection reference.

<a id="canonical-c1e686d551a8fc6bdadaa6940e15dfe1753a45f653f9a928d396adf4efc582f1"></a>

## Next pages — endpoint / 7e87374a895d / 4

- [endpoint.any](data-sources--network_policy_view--reference--group-001.md#canonical-1dff2fd6b4675267a545fbcda84dd656344dd7e2fca50eed177c026c1eb3f3ff)
- [endpoint.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-a2e63320ca8ba04edb81c8ede5d9fefdad0fe0123b4cbfa80e7a7d37327b825a)
- [endpoint.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-4cec9b09377827e68fa86887897d016764cca3cb69327ba673505f222bcfa844)
- [endpoint.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-e85148ee88f957b182dbfb55bce8e5b55a6aa9aa2584b46a11d987c11ca771ba)
- [endpoint.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-72480f3ed28046057f77dd0e438b0fac5ba7dca5a51482d248c0b12de5604834)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-1dff2fd6b4675267a545fbcda84dd656344dd7e2fca50eed177c026c1eb3f3ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f0667ecfa59fadef540b53a7c5e9ae3ccd299e26480d9d403fa3b5a4609a627"></a>

## endpoint.any — endpoint.any / b394299fb73f / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- endpoint.any

<a id="canonical-03ce34a677408971c782d4de18da43070cff038f40bb572b96ef74d5198d37ce"></a>

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

<a id="canonical-09d266fcfe572d3ef4213e0b3eef00a61eaa8216087d416cb24cbad5c935823b"></a>

## Direct properties — endpoint.any / b394299fb73f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-297fec948e9ad44926fdf835fee10af915796e32fdd92fd5a219be58ea775d68"></a>

## Next pages — endpoint.any / b394299fb73f / 4

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-a2e63320ca8ba04edb81c8ede5d9fefdad0fe0123b4cbfa80e7a7d37327b825a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c605c463f97f816981ccaaae6d310c159e3ed6523ae9e6b8a711b820bd92b372"></a>

## endpoint.inside_endpoints — endpoint.inside_endpoints / 189bc6e7ae60 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- endpoint.inside_endpoints

<a id="canonical-8150987b59945e9a1378657f64b7e653ad3d87c179bcbf85dd387d6e73205dfa"></a>

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

<a id="canonical-9a019fe9b1260c1f95bad6061ce70fe7eea50b77b3448762487d1906818fb852"></a>

## Direct properties — endpoint.inside_endpoints / 189bc6e7ae60 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc7589c8db18ef880e76f44e521789ab8a513d35b6cfd1ea223b2d0a3e575dd8"></a>

## Next pages — endpoint.inside_endpoints / 189bc6e7ae60 / 4

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-4cec9b09377827e68fa86887897d016764cca3cb69327ba673505f222bcfa844"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08914e4f1e98953c0d7bb488d5c90c3ba5cb00c4a5711af231f5b004bc4ad137"></a>

## endpoint.label_selector — endpoint.label_selector / e3c5d10df7df / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- endpoint.label_selector

<a id="canonical-e0cd3c54f2b54fb2fd0141a2b8def438eece02326ec7aad655dc789a03de94b4"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c6d49485d3f5b73a5e36a4aefab6c144c7ffb36847bf96bde2b2f07f805e3735"></a>

## Direct properties — endpoint.label_selector / e3c5d10df7df / 3

<a id="canonical-58f146997c7581013e8188cc6c1f19a4ae105ecdb276c15fe6b56975a5d90162"></a>

<a id="canonical-c78ffe4db078b9d6bb0629ee36e379127b18d8c5066223ceae13fc6e6ac2b419"></a>

## expressions property — endpoint.label_selector / e3c5d10df7df / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-49494852eb5fa1971496ae5ee4c711eff07beedfda40d350cf1d14ae81308d23"></a>

## Next pages — endpoint.label_selector / e3c5d10df7df / 5

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-e85148ee88f957b182dbfb55bce8e5b55a6aa9aa2584b46a11d987c11ca771ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daa3b2da875ea620d18270a3b39327d3cffdcae255821894012ca42623f685f2"></a>

## endpoint.outside_endpoints — endpoint.outside_endpoints / 1eddc23f8035 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- endpoint.outside_endpoints

<a id="canonical-81656058b4a4810ed7bfdfb27fec2b9556a120bac8d366e1ac699eb65f505b3a"></a>

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

<a id="canonical-ff65aa34fbd5bbfc367725394242d17787f4638ea3313ba3821c5441d305cc3b"></a>

## Direct properties — endpoint.outside_endpoints / 1eddc23f8035 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2df78247fea188fdfc788ab8dd07c022fae4de35ce58d914ef661c316633af1"></a>

## Next pages — endpoint.outside_endpoints / 1eddc23f8035 / 4

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-72480f3ed28046057f77dd0e438b0fac5ba7dca5a51482d248c0b12de5604834"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9b4ad9d55ec60c11b2fe091385e5e3a6f3d30862483f588ff5b0cb5b401cce9"></a>

## endpoint.prefix_list — endpoint.prefix_list / b854e799a21f / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- endpoint.prefix_list

<a id="canonical-02c33c1accf4c3699cb4e3a9de5b86f6d6194c8f52fcbc8b386ab2c8efeb4acd"></a>

Type: `"single"`. Computed.

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

<a id="canonical-a18dbfd3179ca127029af792d2d214cc8a3355b29b38018faa08ac9240d44e32"></a>

## Direct properties — endpoint.prefix_list / b854e799a21f / 3

<a id="canonical-2e54b97c8c18ac6684cd4d687aa4b4b42d5bc851123d2c4988e685693f659246"></a>

<a id="canonical-2dc57d1e019cad1e9d58f74c145f596ff3de5a8308b17a881982e83814efa7b0"></a>

## prefixes property — endpoint.prefix_list / b854e799a21f / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-e0e0f0e30185b050f58561c1bd82d9651b00eaa840f4873274b62af8c094694e"></a>

## Next pages — endpoint.prefix_list / b854e799a21f / 5

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-873ec0cd9fa6db92022bb39fb919de3516ff18d417312b95f4688d4ae957a618)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e43364bd00d9c4023e2971374c9842708482176fbfcdbb91079b9881855f6cda"></a>

## ingress_rules — ingress_rules / 168bbc6f13a8 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- ingress_rules

<a id="canonical-d73b6c96cfde7400d1a4825ff20e5c1b7082c17a7964d334966e686148e53eac"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections to policy endpoints.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-68a7a09a6d034957a918c54a96251086cca8ad8ec9fd446fe48d633c12055345"></a>

## Direct properties — ingress_rules / 168bbc6f13a8 / 3

<a id="canonical-2b1e6134a7e7aa5616794a7e68f5f701d42695a866a25a8b0028671bfc98a89a"></a>

<a id="canonical-7eb8849c384f813d300ee9ddb481cce21f6dc2f15ecc70c5a27ad9d9037336f8"></a>

## action property — ingress_rules / 168bbc6f13a8 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-cc53986e7f4d5249476b254e908f4c7e4e6d2c3758cf418f427206aa5e437bc4): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3f213c097c03a620d54fca8f0fa1e43047cd5b1815365e5b56b40338dc144385): complete subsection reference.

- [all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-49bbaa72b5f519e3e3ed81b507d51c582333273e692e3028cd2d1758f37723b8): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-dc0149a47db54d523466ab0f08133023995410687f9bd464d653e10a10e74f84): complete subsection reference.

- [any](data-sources--network_policy_view--reference--group-001.md#canonical-ff2f18793515105ea083818eff623d23a18fc6cf3dfd86cc08262185cc8c17c8): complete subsection reference.

- [applications](data-sources--network_policy_view--reference--group-001.md#canonical-fe41903b0befa8b1819b59a0d88ffde1f77665cbe97d39539e1f7eb1e6f3e632): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-c7914377d9bb9e6639003fbad749d70efad879131e124fcde673cd46002f521c): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-d6136e3ac56eee4ba39de181079adbe1c8f2198ddbb59d8aad9ed28885626388): complete subsection reference.

- [label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-029e3b0a474510887d7afb51e6b08002c00b3412fe3ec2fb502a28eb37dfc8f3): complete subsection reference.

- [label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-038de6ac8c8758486e9fe7caaae702b903d5e5ff51a4c22ee94fe2216ae2e234): complete subsection reference.

- [metadata](data-sources--network_policy_view--reference--group-001.md#canonical-d280a0bb125d30291ce2e6a0917f472c1f516fc982b3523ec021139115f3cef5): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-26b746269d6b40913df66d8e4dc07b41ced9b664983d4eb03bf0a7ffc6021c7a): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-0e53adf97b02aeb25e57a81c5298690803a59ac0108c0eebd77aaec77d0bad4a): complete subsection reference.

- [protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-93130476972c41bf865580628e47d6ab7f662c88a152b02ac4cc907f89e13702): complete subsection reference.

<a id="canonical-675ca45cfa231e4cb0c5835d12952495316937ef28c814a5134bad0cf93c56ad"></a>

## Next pages — ingress_rules / 168bbc6f13a8 / 5

- [ingress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-cc53986e7f4d5249476b254e908f4c7e4e6d2c3758cf418f427206aa5e437bc4)
- [ingress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3f213c097c03a620d54fca8f0fa1e43047cd5b1815365e5b56b40338dc144385)
- [ingress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-49bbaa72b5f519e3e3ed81b507d51c582333273e692e3028cd2d1758f37723b8)
- [ingress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-dc0149a47db54d523466ab0f08133023995410687f9bd464d653e10a10e74f84)
- [ingress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-ff2f18793515105ea083818eff623d23a18fc6cf3dfd86cc08262185cc8c17c8)
- [ingress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-fe41903b0befa8b1819b59a0d88ffde1f77665cbe97d39539e1f7eb1e6f3e632)
- [ingress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-c7914377d9bb9e6639003fbad749d70efad879131e124fcde673cd46002f521c)
- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-d6136e3ac56eee4ba39de181079adbe1c8f2198ddbb59d8aad9ed28885626388)
- [ingress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-029e3b0a474510887d7afb51e6b08002c00b3412fe3ec2fb502a28eb37dfc8f3)
- [ingress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-038de6ac8c8758486e9fe7caaae702b903d5e5ff51a4c22ee94fe2216ae2e234)
- [ingress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-d280a0bb125d30291ce2e6a0917f472c1f516fc982b3523ec021139115f3cef5)
- [ingress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-26b746269d6b40913df66d8e4dc07b41ced9b664983d4eb03bf0a7ffc6021c7a)
- [ingress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-0e53adf97b02aeb25e57a81c5298690803a59ac0108c0eebd77aaec77d0bad4a)
- [ingress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-93130476972c41bf865580628e47d6ab7f662c88a152b02ac4cc907f89e13702)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-cc53986e7f4d5249476b254e908f4c7e4e6d2c3758cf418f427206aa5e437bc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23be2d8d32d428412c981106c9e87ee573ca29d14f3daec996ec6ef7946e88ed"></a>

## ingress_rules.adv_action — ingress_rules.adv_action / 7f83d7d51719 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.adv_action

<a id="canonical-b42882a1755276de991f7ca5143941770cec9a63d12a1bc2b474a228f6c41660"></a>

Type: `"single"`. Computed.

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

<a id="canonical-f8cf7cb3c770e089fcc74c9e374e5321471fc92d875fce64d8c49d834d174476"></a>

## Direct properties — ingress_rules.adv_action / 7f83d7d51719 / 3

<a id="canonical-15423e7d21a45f52ce4a8cd7a0ab57bdf20335933cffc637e3d2f234ef731028"></a>

<a id="canonical-6953e2f53d6f4571816ea222b1454e4667ec34c86539e74ff741933f9e78f2b8"></a>

## action property — ingress_rules.adv_action / 7f83d7d51719 / 4

Type: `"string"`. Computed.

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

<a id="canonical-bee659d7015e0b10579ccd510d24cf2aa5decb0fbfa41cd83fd191c593535639"></a>

## Next pages — ingress_rules.adv_action / 7f83d7d51719 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-3f213c097c03a620d54fca8f0fa1e43047cd5b1815365e5b56b40338dc144385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0a0628e5e1293c3e4bc56c8fc2d076ad171b8da1cf07f17b49d11177e13e39b"></a>

## ingress_rules.all_tcp_traffic — ingress_rules.all_tcp_traffic / 2b259534e1df / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.all_tcp_traffic

<a id="canonical-690ca9abecdb3ec7ce40a33993d3c5ba870c7f6bd987aac04133098eb7889622"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-99aaafa31ffd2d73d6ce21fb073b688ea8c9cf6d97082d38e0d4ad19f2225abe"></a>

## Direct properties — ingress_rules.all_tcp_traffic / 2b259534e1df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f13e4d6ae98395a9f7ad73d744121aa298ec2b9e5dcac596dc401f3bea107906"></a>

## Next pages — ingress_rules.all_tcp_traffic / 2b259534e1df / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-49bbaa72b5f519e3e3ed81b507d51c582333273e692e3028cd2d1758f37723b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3eda9fac24a12679cc92944081805fe40d04ace1a41c9e1a590390dfca6f796"></a>

## ingress_rules.all_traffic — ingress_rules.all_traffic / 933567e07769 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.all_traffic

<a id="canonical-5ea8a450ee0902298b29bbc245788385d1eb8eae9107f6189933800ecaf69b32"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-a41ebeaa5e32bff0c1e4da8374b2b07597172e10bbaa3f6317ce4d9ef6c867e6"></a>

## Direct properties — ingress_rules.all_traffic / 933567e07769 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b68042c68a9782e0798ae5382bebc70deb15215c2054d7c093d2dd80cc4ba475"></a>

## Next pages — ingress_rules.all_traffic / 933567e07769 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-dc0149a47db54d523466ab0f08133023995410687f9bd464d653e10a10e74f84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-686e5d46b5b71328a30466070453b5486557209405dcc34ee25fe3cf11e889ca"></a>

## ingress_rules.all_udp_traffic — ingress_rules.all_udp_traffic / 94ca0a8bbcb3 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.all_udp_traffic

<a id="canonical-0c4b8026e6807f16a02623031d756b030ad10bb26589b7e94f5deb090fd44fe0"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-e76e539c54f77bde01d71cb8744e664dee15760c84b961da78254936a1e6e2d8"></a>

## Direct properties — ingress_rules.all_udp_traffic / 94ca0a8bbcb3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3c5e0065922a68ccc99660c4c07d7e0476e99c916240faff52411ac54d7f7ba"></a>

## Next pages — ingress_rules.all_udp_traffic / 94ca0a8bbcb3 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-ff2f18793515105ea083818eff623d23a18fc6cf3dfd86cc08262185cc8c17c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21482f6876e2627c7983ab5d3ad566bf5820b65ae965f15d0188833b78dbc901"></a>

## ingress_rules.any — ingress_rules.any / a2702ded3a3c / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.any

<a id="canonical-5a389005983de25fb28de257c7d60b1c681dc032cc1bd2ecb163f9eabe375496"></a>

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

<a id="canonical-001accc904757583330301a0762160d584b5059671fae96e8bdbe0b7e8a291af"></a>

## Direct properties — ingress_rules.any / a2702ded3a3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02484c528a34da1b7230671c3c517d3edbbeaa82891febf47ff3f2b0a4377162"></a>

## Next pages — ingress_rules.any / a2702ded3a3c / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-fe41903b0befa8b1819b59a0d88ffde1f77665cbe97d39539e1f7eb1e6f3e632"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2b52c84d2f65ba6ca72e73218f927382689e6612d7dc3870358704eae2433ca"></a>

## ingress_rules.applications — ingress_rules.applications / 88c41380b861 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.applications

<a id="canonical-d73177f1ea41684be6edd35794e55ff31d07aa0bda7fd9cf439f82f4dc0adab2"></a>

Type: `"single"`. Computed.

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

<a id="canonical-f675ade4dc60b06e7b2168a6db436e6384d7a4b6756daf91988672406fa632fc"></a>

## Direct properties — ingress_rules.applications / 88c41380b861 / 3

<a id="canonical-bbf52bb5a43f7a67b425966eda123313216331c973a62d6d82557efa0aa54d6d"></a>

<a id="canonical-4fd6de35b7be923363eea931d00f61d9b10fa47890dcec16cd6fda7a4fbfe429"></a>

## applications property — ingress_rules.applications / 88c41380b861 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-62e60e7e0080eebb2097ddbf4c73757e75ef89bedb6efaee044f6dd3f955a0e4"></a>

## Next pages — ingress_rules.applications / 88c41380b861 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-c7914377d9bb9e6639003fbad749d70efad879131e124fcde673cd46002f521c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcaf347978ed7edddc66bf15f11802f82e07cf3548b16bfd8c9c9e30ceea4dca"></a>

## ingress_rules.inside_endpoints — ingress_rules.inside_endpoints / 8ccbfc635bbd / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.inside_endpoints

<a id="canonical-7599af7b7e8d3646578bf98f5502f754d1eb12fe7a3fa640638fdf808e482fb0"></a>

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

<a id="canonical-70bb3923ff2c51f11283a69d9a4c15319f8482da7ee2f2e081f5848697754c20"></a>

## Direct properties — ingress_rules.inside_endpoints / 8ccbfc635bbd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f275ac7713ac8300c9f068e1456bcc08b16e6e677216e9160c6f82bcc03968b8"></a>

## Next pages — ingress_rules.inside_endpoints / 8ccbfc635bbd / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-d6136e3ac56eee4ba39de181079adbe1c8f2198ddbb59d8aad9ed28885626388"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f9e2f70485e043c87b7ecd2795a02b9c9ee6d216dd1a0d85e2a022ed7766929"></a>

## ingress_rules.ip_prefix_set — ingress_rules.ip_prefix_set / a7d996ae8a7e / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.ip_prefix_set

<a id="canonical-fde060c9b4236fccb1fa1c73929919a7e8519fd69b6932bd7296cfbe05d4a83b"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c967ee1a0dce2a3859c7932c8c94661196a45b0a6edccb398cecc6336c05e455"></a>

## Direct properties — ingress_rules.ip_prefix_set / a7d996ae8a7e / 3

- [ref](data-sources--network_policy_view--reference--group-001.md#canonical-a3ce7a01ecd34834090218be5acaa4433125622d44a74891f2889f5db7ac272b): complete subsection reference.

<a id="canonical-1fe49fe6009d496cfcc18f52eeef73c084bec253d4920c0195116f3ca7a6743b"></a>

## Next pages — ingress_rules.ip_prefix_set / a7d996ae8a7e / 4

- [ingress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-a3ce7a01ecd34834090218be5acaa4433125622d44a74891f2889f5db7ac272b)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-a3ce7a01ecd34834090218be5acaa4433125622d44a74891f2889f5db7ac272b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76ed89c194be006e38f75a9cf071721410c5b68b886a6aff218ec932938eed39"></a>

## ingress_rules.ip_prefix_set.ref — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-d6136e3ac56eee4ba39de181079adbe1c8f2198ddbb59d8aad9ed28885626388)
- ingress_rules.ip_prefix_set.ref

<a id="canonical-15c96c070378e95998a1c5d616fb659f9c9500244ebfa2af7edf6d77c35c5582"></a>

Type: `"list"`. Computed.

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

<a id="canonical-641f266fd516ce143169313c6f0fe9161a78e0f72e31307a65dbda0971eaa304"></a>

## Direct properties — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 3

<a id="canonical-af5b5dca4063dff126a4bc820703f236dadb2f759c454c124ca43d9d9020ca05"></a>

<a id="canonical-36725fe85590915e4bce2a040eea8da87dc7ce38c67dfab0b0da039768c98997"></a>

## kind property — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 4

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

<a id="canonical-3afbd449fd22ba2aefb881608fc0eef4c839f276e83cac9f7ec39c7214f978be"></a>

<a id="canonical-c3c65f035edc01619b8b691dbbd2cf25c3a91ce1a96899d3613fe2c29ad4e8a9"></a>

## name property — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 5

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

<a id="canonical-cc43883b84cda2c713983ae11dcd447197c5f924126e92880290025f70c94256"></a>

<a id="canonical-e018ad4de6a586ba0444d1d996bdb99931a247374650aba53652204470708887"></a>

## namespace property — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 6

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

<a id="canonical-c0fc609e4baf571171ae143ad0b07b6581f1fd22001220c3bb659c33c432610e"></a>

<a id="canonical-ee9cc5d8c5d33e337d559780e51694bc0a7e8b56a759b61c7bbdc35b20fb714a"></a>

## tenant property — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 7

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

<a id="canonical-19bf2c858bc592ad3d0286504fba634013e8cada04dea00e906559052eae228a"></a>

<a id="canonical-f7559a5c9e7b5f7de83e33160ee77709bece0a24b8224fa29a09d0fa7403ee2e"></a>

## uid property — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 8

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

<a id="canonical-e95fe22dfa85aead297d084f22282248dca10e240da40acf58a4bf50a5de0641"></a>

## Next pages — ingress_rules.ip_prefix_set.ref / bc01a84db241 / 9

- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-d6136e3ac56eee4ba39de181079adbe1c8f2198ddbb59d8aad9ed28885626388)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-029e3b0a474510887d7afb51e6b08002c00b3412fe3ec2fb502a28eb37dfc8f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db8c06c6bc09293dd18c2fbd89dc686146ef91079919b7f238ca5780d633ef0a"></a>

## ingress_rules.label_matcher — ingress_rules.label_matcher / d70c5a1c16a7 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.label_matcher

<a id="canonical-d202640fcc7b5d634641cf712ffc8887a09a14e89cf886d47952fe6bde85fcf7"></a>

Type: `"single"`. Computed.

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

<a id="canonical-234222058c0bf5d76c9ff11ae34185e310cb1a063d24cf6f0b4afcc69749b806"></a>

## Direct properties — ingress_rules.label_matcher / d70c5a1c16a7 / 3

<a id="canonical-459ef1356f3ef99cd4e683bfeadd1b78bd6dce36b8d341fbdcba0b0b4ab4d30c"></a>

<a id="canonical-ce196e2ed2052945eda05baf5ad68b5e02091f7502b3a1cd691ec5711cae3de8"></a>

## keys property — ingress_rules.label_matcher / d70c5a1c16a7 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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

<a id="canonical-2e453f24fc4e7de0e430be3506884a763e31bdb0b47e98658c2a8f7af48aee58"></a>

## Next pages — ingress_rules.label_matcher / d70c5a1c16a7 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-038de6ac8c8758486e9fe7caaae702b903d5e5ff51a4c22ee94fe2216ae2e234"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a270a5d7872c6667e7e6d1c097beba3cc0eed38bbc70ea4ad15539d7a5655474"></a>

## ingress_rules.label_selector — ingress_rules.label_selector / ef03a58b6d72 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.label_selector

<a id="canonical-942408488dd008726390c11046266d0430101d1f70818f41203d9f539c773f32"></a>

Type: `"single"`. Computed.

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

<a id="canonical-9510eb1f47d5c959408fe2f4c0d1a22bae1e80cb9a0d71519c7846eb8c694f02"></a>

## Direct properties — ingress_rules.label_selector / ef03a58b6d72 / 3

<a id="canonical-3239945ccb1abb4ff151511d62382ea86a74563bc49e2d4bf4549b118a58821b"></a>

<a id="canonical-6c44ac00d9bcd1ed0cbd7dfd5c57726a282188a5c3362e04900ccbbea0444f9c"></a>

## expressions property — ingress_rules.label_selector / ef03a58b6d72 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-f70243a6dfeebee4982ee6c550b013f55b33775fa9c1d0b469a4f92e56b71345"></a>

## Next pages — ingress_rules.label_selector / ef03a58b6d72 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-d280a0bb125d30291ce2e6a0917f472c1f516fc982b3523ec021139115f3cef5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eee2d777e0f0d817da3d5a0780979a1ed3827e0cbd7d6963a148d8c1a8701042"></a>

## ingress_rules.metadata — ingress_rules.metadata / 344d7358a5aa / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.metadata

<a id="canonical-9f77aff7530895024e8ef6c990da9c7a0a827d42f0e8262dc9dc1b27f6396bfd"></a>

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

<a id="canonical-1d506588ed15d92ddaa4e41797163e00fbc8f9f01961889bb6823471836f7278"></a>

## Direct properties — ingress_rules.metadata / 344d7358a5aa / 3

<a id="canonical-7e87efce1850edf87be48ffe1d35c9b616b8b666d434d20846d383795f6b6dcb"></a>

<a id="canonical-f4c673a0a2cd99f8759d13d180f139fe77a65b3f86fb20efcada85f08899e591"></a>

## description_spec property — ingress_rules.metadata / 344d7358a5aa / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-8d55e1163859ccb2a1725b897240df63fc8f0022cd3da287b6b07dbf73cbe7df"></a>

<a id="canonical-fc05a14d9bcbdd6a00c4f51967c7e1e8038af899547251bcd20b0ce358da2a83"></a>

## name property — ingress_rules.metadata / 344d7358a5aa / 5

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

<a id="canonical-04f732bf9ce26c004b54ad21c9e0325e30addea2b9055f3c15c45c784526b38f"></a>

## Next pages — ingress_rules.metadata / 344d7358a5aa / 6

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-26b746269d6b40913df66d8e4dc07b41ced9b664983d4eb03bf0a7ffc6021c7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9efe4bd586ec50b8f55f432e132eb84b5649582fddd4146c2251eadc036a7c36"></a>

## ingress_rules.outside_endpoints — ingress_rules.outside_endpoints / c043da20ee0a / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.outside_endpoints

<a id="canonical-c53e021c424979978413d038908e0bae65850656d0fccde9cb111eea46b8ed0c"></a>

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

<a id="canonical-5b041abd0c9c85974f2249d882e76d8c028d7c081dea0859cdc5929ff858d422"></a>

## Direct properties — ingress_rules.outside_endpoints / c043da20ee0a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7628cbdde7a725b18e82eec48b52a84dcaf26897c0ba1421733e7d19340bfe1e"></a>

## Next pages — ingress_rules.outside_endpoints / c043da20ee0a / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-0e53adf97b02aeb25e57a81c5298690803a59ac0108c0eebd77aaec77d0bad4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375c8fcbb80adf83fa372bbebd103560660491a0baea36a94fe1da29e036cd95"></a>

## ingress_rules.prefix_list — ingress_rules.prefix_list / e28ce5e0e0b3 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.prefix_list

<a id="canonical-7bcece0f7573d318268bf7648075d88aa843d961673ee1de2163d6b4a713b28f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2a98334231151125ffa1a8389acd437b11fda662548126554398e5f736881906"></a>

## Direct properties — ingress_rules.prefix_list / e28ce5e0e0b3 / 3

<a id="canonical-85ddd6497a11b08a6d8c5cdfb1ead328a913dd6bb1b24b067ca35c171dc85215"></a>

<a id="canonical-7a69bdb8b153f2ecdff8473420738fb6e53cfb6e67dd92eb5b9a2155cacf985e"></a>

## prefixes property — ingress_rules.prefix_list / e28ce5e0e0b3 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-280527e8611035d1fa73249a9b71b6b4e28b14e75e756da77e28ffe4c269636f"></a>

## Next pages — ingress_rules.prefix_list / e28ce5e0e0b3 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-93130476972c41bf865580628e47d6ab7f662c88a152b02ac4cc907f89e13702"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbbfaf328dc52dea47028cca14b5dcf2b89ea124f04f7f8220ee434c60ef7925"></a>

## ingress_rules.protocol_port_range — ingress_rules.protocol_port_range / 9b6161616a3f / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- ingress_rules.protocol_port_range

<a id="canonical-ae691ab5baa92d2ca446c272f8941e9b20720ce358d3402488d50baca6ba3b0b"></a>

Type: `"single"`. Computed.

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

<a id="canonical-f75805f504c2974ea45057a1e9fa8c2e845e6b0d1d2fb0ff4e3711cf3ebd7cea"></a>

## Direct properties — ingress_rules.protocol_port_range / 9b6161616a3f / 3

<a id="canonical-19942d874b5d2f8a91baf0910a07ec88180167debe0a193371e083aabdc2a41a"></a>

<a id="canonical-391745f3efff07152e86e1f21d86f2539f298641d557548a2e43ee19cd00c74b"></a>

## port_ranges property — ingress_rules.protocol_port_range / 9b6161616a3f / 4

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

<a id="canonical-aca65bef7c59f8e8ef2377385b69ecdef9570489144642e985481f2f9ae27101"></a>

<a id="canonical-3249dcb5762907cb7a747b93a81d22feccfa5463ae54d9d2987790f79b32b992"></a>

## protocol property — ingress_rules.protocol_port_range / 9b6161616a3f / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

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

<a id="canonical-f2f29d8b031d37ed26d53db2cbe382971cb1dbcc06557b8d081b311250272ac8"></a>

## Next pages — ingress_rules.protocol_port_range / 9b6161616a3f / 6

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-136634ba34b9f662a5efd64537c1a515a1aa979072ff324c5cdacb82ba47f0ec)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
