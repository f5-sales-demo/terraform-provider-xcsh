---
page_title: "xcsh_network_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy reference."
---

# xcsh_network_policy reference

<a id="canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f169f089693cc5ac07a09fe4e42372860260b9e492ec4eed8811202e2f3217f3"></a>

## Property reference — Property reference / 73acb7ea7b8c / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- Property reference

<a id="canonical-141e84fdad1ab70823850c136d432f9059b5e587dc6078751f654c95ed4ea576"></a>

## Direct properties — Property reference / 73acb7ea7b8c / 3

<a id="canonical-5de619a91eba445b61c33f0383993c130e47c5899941071b3efad93aee21dec1"></a>

<a id="canonical-5e2d9268fd6e7604eb85122af9d77aecbaaecff9e8d5d59b25201a7a4a949894"></a>

## annotations property — Property reference / 73acb7ea7b8c / 4

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

<a id="canonical-07b3ce448c666d5edcc92809a6ad6f207838184a6eb3981ab805cbc65ee9657e"></a>

<a id="canonical-cbf8b72437ef17eb9219c076060451dcb443c82f281fa36787fd14328258fcc5"></a>

## description property — Property reference / 73acb7ea7b8c / 5

Type: `"string"`. Computed.

Description of the NetworkPolicy.

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

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe): complete subsection reference.

<a id="canonical-c4dd954bc40ba498c664b2c24149bcc1f7ed73db177f3506c929ba49211261b2"></a>

<a id="canonical-b475007b8f67a75c3d4d48afc967f7e1a34989d1f380089aa8823e07baba3378"></a>

## id property — Property reference / 73acb7ea7b8c / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0dc84ca1958c6cc17b02c6de9afd0dc032b45697dc67040b5e1da02b9b45fb63"></a>

<a id="canonical-5f8e70b8716f1308b72a00ac4181d03552467b812a7096a7b12ce32e4444dd8f"></a>

## labels property — Property reference / 73acb7ea7b8c / 7

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

<a id="canonical-10305b1a16e84e8c8acc453d80d39953c439031af503f27e8e7712e36c2941a5"></a>

<a id="canonical-6b070c881971a8f6856f05b25149a6ec14110588479057563d9fde59c677ff48"></a>

## name property — Property reference / 73acb7ea7b8c / 8

Type: `"string"`. Required.

Name of the NetworkPolicy.

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

<a id="canonical-1d72b1d532e7a5b52e6cfdf2104de6cf69d6492ea0b10421fb69ba29f23be942"></a>

<a id="canonical-19d323fba8c4152efa20e8c2d3eaac691489908a61c26342fedac61018ec5a8b"></a>

## namespace property — Property reference / 73acb7ea7b8c / 9

Type: `"string"`. Required.

Namespace where the NetworkPolicy exists.

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

- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26): complete subsection reference.

<a id="canonical-db4ff0795e10ad9fd8e0a54bfe4feed2689a49bcfa9f554474facd16b3fe99d1"></a>

## All schema paths — Property reference / 73acb7ea7b8c / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_policy--reference--group-001.md#canonical-5de619a91eba445b61c33f0383993c130e47c5899941071b3efad93aee21dec1) |
| `description` | [description](data-sources--network_policy--reference--group-001.md#canonical-07b3ce448c666d5edcc92809a6ad6f207838184a6eb3981ab805cbc65ee9657e) |
| `endpoint` | [endpoint](data-sources--network_policy--reference--group-001.md#canonical-eea0342caf7bf6bb954b892af36b22cba252b08788b6b11b9c4d635b60b90357) |
| `endpoint.any` | [endpoint.any](data-sources--network_policy--reference--group-001.md#canonical-ee1c68e606c4abc83dd503ef27359f04bb5466a34f5519eba83115888bf79891) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-81232ea0444f95d6d664d2deed8f1ee7806030416237b9ba149570ec446cf811) |
| `endpoint.label_selector` | [endpoint.label_selector](data-sources--network_policy--reference--group-001.md#canonical-71792c139b6373f043f928aad7ee903264030cc989b39f516f96b0bd53e81bc9) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](data-sources--network_policy--reference--group-001.md#canonical-6bb373b5b2789c5fbada3827c2adf6d29bc80a73ce7daff199d6e22ccbd20d48) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-df5fc2c8903517a2c42503f0d52314892036f54ed02716faab3b3628f0805741) |
| `endpoint.prefix_list` | [endpoint.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-92958dc4d817973d7d9d37682ae42b3904f5f466c10db183d9f8180ac255f001) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](data-sources--network_policy--reference--group-001.md#canonical-d1a58e411af3228f6faa547621cd58812a055e8295ba91a3e921657496b7ff6f) |
| `id` | [id](data-sources--network_policy--reference--group-001.md#canonical-c4dd954bc40ba498c664b2c24149bcc1f7ed73db177f3506c929ba49211261b2) |
| `labels` | [labels](data-sources--network_policy--reference--group-001.md#canonical-0dc84ca1958c6cc17b02c6de9afd0dc032b45697dc67040b5e1da02b9b45fb63) |
| `name` | [name](data-sources--network_policy--reference--group-001.md#canonical-10305b1a16e84e8c8acc453d80d39953c439031af503f27e8e7712e36c2941a5) |
| `namespace` | [namespace](data-sources--network_policy--reference--group-001.md#canonical-1d72b1d532e7a5b52e6cfdf2104de6cf69d6492ea0b10421fb69ba29f23be942) |
| `rules` | [rules](data-sources--network_policy--reference--group-001.md#canonical-b3cd9cd411d253ba6a2f801964b6623cacb8844b6212586d0ef6922fdff6e8df) |
| `rules.egress_rules` | [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-c87f7c38587a4fcc3b12c8e1f1f4b8509c241ba99e3ca6289b6e262d6a34c7cc) |
| `rules.egress_rules.action` | [rules.egress_rules.action](data-sources--network_policy--reference--group-001.md#canonical-a50208f958d53dbbbae1cf2c4c01647e1b633a3c5d927d161fb0d12089a55477) |
| `rules.egress_rules.adv_action` | [rules.egress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-933fe72a647d22a6f740984ed6cf4fec14623a14cecdbc87d5acab96ec24caa3) |
| `rules.egress_rules.adv_action.action` | [rules.egress_rules.adv_action.action](data-sources--network_policy--reference--group-001.md#canonical-d48a3ef4f8399c578d11c80908faca0b18c796b89123590aa46630777fff60f8) |
| `rules.egress_rules.all_tcp_traffic` | [rules.egress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-d95d1fb1037cc6a8331eaf677bf4c214c5706b8c8261563c23e3d2b3862b3af6) |
| `rules.egress_rules.all_traffic` | [rules.egress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-d56d340bf81bb6ef0faf4109d0aebfeea9eba960af1f6d43c7066771899206f6) |
| `rules.egress_rules.all_udp_traffic` | [rules.egress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-8bb01550726d18051b2d79737f279b7b9c6a92eba2776e7d21b2ff8b31e3875f) |
| `rules.egress_rules.any` | [rules.egress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-f8ee290f972c9b93704f865085f9a2c0ee86fdb8fb40dd74a853ef101c2cb0b5) |
| `rules.egress_rules.applications` | [rules.egress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-a701a973484a9a05701c44f0ba7c42a95be6d1bca52682d1e68eb97ca587b004) |
| `rules.egress_rules.applications.applications` | [rules.egress_rules.applications.applications](data-sources--network_policy--reference--group-001.md#canonical-1aa67850296bf610c540514a06b9e16bcd1cbc2b7aaa4a508368a6503cead5eb) |
| `rules.egress_rules.inside_endpoints` | [rules.egress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-19798de2269c957e5ecd4b12e6252ab4bc904efb9e20f7b1c455a2c76cc1e925) |
| `rules.egress_rules.ip_prefix_set` | [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-6e813d9a3a9743fe54d766ae6f77edf55d582f1d6882978ddba8685c7377063d) |
| `rules.egress_rules.ip_prefix_set.ref` | [rules.egress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-14beb053eac489635b5194be4064277a36497874c8c3d66887ed59ca3b141a2b) |
| `rules.egress_rules.ip_prefix_set.ref.kind` | [rules.egress_rules.ip_prefix_set.ref.kind](data-sources--network_policy--reference--group-001.md#canonical-5c9234d4e33d601478d7dd03c04a5c0a8543730ace3703dacbd8a1aeb5d1c317) |
| `rules.egress_rules.ip_prefix_set.ref.name` | [rules.egress_rules.ip_prefix_set.ref.name](data-sources--network_policy--reference--group-001.md#canonical-d458602bce2b2601043ec6595f2632342a19811e5f70f09940a2e288cd683f87) |
| `rules.egress_rules.ip_prefix_set.ref.namespace` | [rules.egress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy--reference--group-001.md#canonical-18ab5f9f5ec707a5908a208437a3661c1bb97ed05127f90c958cf637d67b43a1) |
| `rules.egress_rules.ip_prefix_set.ref.tenant` | [rules.egress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy--reference--group-001.md#canonical-a537266f3a2964965b2969a65da566b5adda592b4e2ac2a572be8adaa5493ae7) |
| `rules.egress_rules.ip_prefix_set.ref.uid` | [rules.egress_rules.ip_prefix_set.ref.uid](data-sources--network_policy--reference--group-001.md#canonical-46c6231a867f1a0f7d104b19a135495db687c6efd3c0af50ca1dd304735163f1) |
| `rules.egress_rules.label_matcher` | [rules.egress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-df27ad6fd539cc0ad1db85d2e530f7ce3ec78c6aa2546fcea813075bba43655e) |
| `rules.egress_rules.label_matcher.keys` | [rules.egress_rules.label_matcher.keys](data-sources--network_policy--reference--group-001.md#canonical-006f2e3d3ef3090329cdf3af531686623afbd11c680a3479c2be50d5b653524e) |
| `rules.egress_rules.label_selector` | [rules.egress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-579ac2dedd2152d0de46c67689b9fccd7ccf93dc152bf4a749c90a71bee9a4fc) |
| `rules.egress_rules.label_selector.expressions` | [rules.egress_rules.label_selector.expressions](data-sources--network_policy--reference--group-001.md#canonical-c342083d5b6832d689a86419a2807e0c7fed8548b6c4357b0f7e145b1256679e) |
| `rules.egress_rules.metadata` | [rules.egress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-ef1a859c9eeb2fd765d1222e2557d03494513d73cd1f11d27f46961fe241469e) |
| `rules.egress_rules.metadata.description_spec` | [rules.egress_rules.metadata.description_spec](data-sources--network_policy--reference--group-001.md#canonical-52637feccce34ea10980fbd8e7832fa2fe2e8bbb008f9ce956911e7a4af8bc45) |
| `rules.egress_rules.metadata.name` | [rules.egress_rules.metadata.name](data-sources--network_policy--reference--group-001.md#canonical-1804e4f8d5c66542835f2ebb2e66e896256d19d1549b0ea9909169af819ce196) |
| `rules.egress_rules.outside_endpoints` | [rules.egress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-05d4f7068c498e764090243830971bbfe708b991c9803e963913e98105d03b10) |
| `rules.egress_rules.prefix_list` | [rules.egress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-36c7dbe4ef760013396af538370efd22d000401301a8ba66c18262bb45376147) |
| `rules.egress_rules.prefix_list.prefixes` | [rules.egress_rules.prefix_list.prefixes](data-sources--network_policy--reference--group-001.md#canonical-7f2592aa50eebb17b29098721d52940305f5bd57240d61318fe27b54a735aae6) |
| `rules.egress_rules.protocol_port_range` | [rules.egress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-5439c6aca8c38ede23cd727c71e8d9b18eb1133a06dc168e966c08e3553058a8) |
| `rules.egress_rules.protocol_port_range.port_ranges` | [rules.egress_rules.protocol_port_range.port_ranges](data-sources--network_policy--reference--group-001.md#canonical-bd90aa02feb18a9c82659904db4197c8a27cef40adc0734e228c163ad486ea66) |
| `rules.egress_rules.protocol_port_range.protocol` | [rules.egress_rules.protocol_port_range.protocol](data-sources--network_policy--reference--group-001.md#canonical-002773e2b30e2f49617c51cb9c56fb25ff7e61501ac103994317a9af88e06865) |
| `rules.ingress_rules` | [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-2a5516510991f6cf289290609c42563b1f730802b624b8d701aa7be17f8eaaa6) |
| `rules.ingress_rules.action` | [rules.ingress_rules.action](data-sources--network_policy--reference--group-001.md#canonical-96edf5952f78e1fbcdf331c5ea752f7b9792d7558d016f6687d1f4c6913bae9d) |
| `rules.ingress_rules.adv_action` | [rules.ingress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-e5351c3dfe6d2615a5cebf1d3b67e508825c17a62742bc0dbb0c5cf4fd633780) |
| `rules.ingress_rules.adv_action.action` | [rules.ingress_rules.adv_action.action](data-sources--network_policy--reference--group-001.md#canonical-bfd16925b753aabdbff0754f5dd2e8fea655d0af51f670cba6b380f6bd53f444) |
| `rules.ingress_rules.all_tcp_traffic` | [rules.ingress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-97bb353675013d86cd1eccb72c439c79a14dd42a29e9fac2b32642bf7ec05591) |
| `rules.ingress_rules.all_traffic` | [rules.ingress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-3bf620acc20a291b4e65ae1347452fcef623f0a6de719b60734ec1bb04fef17a) |
| `rules.ingress_rules.all_udp_traffic` | [rules.ingress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-9cf95f9037cb40a322b2f4bd0e4a29be9586187bb84e8b98e663126eadb5f9e2) |
| `rules.ingress_rules.any` | [rules.ingress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-0b4f1665819c09b5b30b89e74b934963bdaae5796c48c0a18ff444243c83c35a) |
| `rules.ingress_rules.applications` | [rules.ingress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-2dfe52b2b0d53cdc4448dbd9d6434e6aac4c2ea6f1b51fb4878f70c0824dd682) |
| `rules.ingress_rules.applications.applications` | [rules.ingress_rules.applications.applications](data-sources--network_policy--reference--group-001.md#canonical-ce0ebb6d86993fc672eb2b349524ed7950c91b91c614bee821af64f2e6bcf860) |
| `rules.ingress_rules.inside_endpoints` | [rules.ingress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-ce372acc3ca38a461cb6364bfd8e07ab5dc0fd4cc30b4f43d9b9ebd43ae90110) |
| `rules.ingress_rules.ip_prefix_set` | [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-4d7693e2827731f880943e7bc86e2a08d7fc251b501bae3e1e63ef276beaba26) |
| `rules.ingress_rules.ip_prefix_set.ref` | [rules.ingress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-134ed8949b5a873dd958cd3ee5bdf56fd945645c5628581b13e915e06117cc15) |
| `rules.ingress_rules.ip_prefix_set.ref.kind` | [rules.ingress_rules.ip_prefix_set.ref.kind](data-sources--network_policy--reference--group-001.md#canonical-250e9db41c4a6f851c67c7081a5eba23508bea75b915fe56c9a4146f0c38b5e3) |
| `rules.ingress_rules.ip_prefix_set.ref.name` | [rules.ingress_rules.ip_prefix_set.ref.name](data-sources--network_policy--reference--group-001.md#canonical-2f0d3469b920ba504c46306ca8860034f948a90f57c24c146db9fc236bdfa92c) |
| `rules.ingress_rules.ip_prefix_set.ref.namespace` | [rules.ingress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy--reference--group-001.md#canonical-844e5075e896f4b6f9d12744ede833fd7c4c9bfe46168aec540dbf50d0d349aa) |
| `rules.ingress_rules.ip_prefix_set.ref.tenant` | [rules.ingress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy--reference--group-001.md#canonical-b1a3fcbd1f01969c702e3db7186611cfe62bdef64fd551139f38d9fda1047fdf) |
| `rules.ingress_rules.ip_prefix_set.ref.uid` | [rules.ingress_rules.ip_prefix_set.ref.uid](data-sources--network_policy--reference--group-001.md#canonical-99fee300caff6cf5fc00757bbe5eaf6deab3f7675aab48a8063a71c9afa775c2) |
| `rules.ingress_rules.label_matcher` | [rules.ingress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-47dcf6dcd413655a0921336deb40dbe4806aa4f846505ee7341322483135a4a4) |
| `rules.ingress_rules.label_matcher.keys` | [rules.ingress_rules.label_matcher.keys](data-sources--network_policy--reference--group-001.md#canonical-ec67ae6a315bfc655c1647beee4e84be5448257dd20796c943996b0be88b503d) |
| `rules.ingress_rules.label_selector` | [rules.ingress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-5e26c97bc85da7d028781488ecb027469ab937b97a927e02e3beaaec2d6cb162) |
| `rules.ingress_rules.label_selector.expressions` | [rules.ingress_rules.label_selector.expressions](data-sources--network_policy--reference--group-001.md#canonical-d1861fef6ba8f48b5e897d332e5b220c5adfd0452e57bacb6dadd5cb3bf309ff) |
| `rules.ingress_rules.metadata` | [rules.ingress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-2369e0bb4fe4b72586d309e4a3e36d9194d08fd9946f3bf56792f6222f157644) |
| `rules.ingress_rules.metadata.description_spec` | [rules.ingress_rules.metadata.description_spec](data-sources--network_policy--reference--group-001.md#canonical-45279ddc76cc60228275d2d576ac2290d2ac9f5d190d469cc7fc41d064f9f70d) |
| `rules.ingress_rules.metadata.name` | [rules.ingress_rules.metadata.name](data-sources--network_policy--reference--group-001.md#canonical-026f4a0c945817edbc0c9f2114b687465f015285745df3b4ca9d52ce4bf2793c) |
| `rules.ingress_rules.outside_endpoints` | [rules.ingress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-4aae885c912224649b977c366fd4adcc29aadaf9f3f07946795da4cb53a4ba3a) |
| `rules.ingress_rules.prefix_list` | [rules.ingress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-a9223642006d6dcd38a02739a7302439226c3a9692b9e3ff5604e4c715ff0a86) |
| `rules.ingress_rules.prefix_list.prefixes` | [rules.ingress_rules.prefix_list.prefixes](data-sources--network_policy--reference--group-001.md#canonical-ba431253a568f86fa95ec4aa125227ad4d525f594d06e8d78e50ff4bfc4f9e2b) |
| `rules.ingress_rules.protocol_port_range` | [rules.ingress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-71b48295d166ddef6c73127df7ace4d4712df5bdd3b40808143bf356ab638961) |
| `rules.ingress_rules.protocol_port_range.port_ranges` | [rules.ingress_rules.protocol_port_range.port_ranges](data-sources--network_policy--reference--group-001.md#canonical-8486feef459ea4e76f051bb9b6d72da367d0ab310e1aed5c465fb584e19ea186) |
| `rules.ingress_rules.protocol_port_range.protocol` | [rules.ingress_rules.protocol_port_range.protocol](data-sources--network_policy--reference--group-001.md#canonical-83741b5d200bb89f25ad828847304159f46d0a486f3a4fccce5bbc4b97f2bdc4) |

<a id="canonical-0fd3fb9474c15beb54657df51de96234831770567b6a1ff3ef8b301bcce632d4"></a>

## Next pages — Property reference / 73acb7ea7b8c / 11

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f9f5b331c051ed74732c893bc0a769b14d6cf3f6aa58782da61ceeae34305a9"></a>

## endpoint — endpoint / d4cc85ecb73d / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- endpoint

<a id="canonical-eea0342caf7bf6bb954b892af36b22cba252b08788b6b11b9c4d635b60b90357"></a>

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

<a id="canonical-a15b7fd8064798d71ced78fd91aa567b93f1d0dfbb8e3d7a8aff743e89ed6e6e"></a>

## Direct properties — endpoint / d4cc85ecb73d / 3

- [any](data-sources--network_policy--reference--group-001.md#canonical-50a39d41fb47657ce2506798a62dfcaf587e52cc8cc7e1c01307349e91a8a719): complete subsection reference.

- [inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-c4e6b1708353a516a8e8cdae1a6723fe32349312f1150f412c73a2f8afe6856d): complete subsection reference.

- [label_selector](data-sources--network_policy--reference--group-001.md#canonical-8c8df0b43d3d9600043ed818f9ac559b26b56d30aff813497205a8428e272c0d): complete subsection reference.

- [outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-8207789040ba8897674ff5fcd7074e3612ba31419323b8b57d8286a3bdc613b1): complete subsection reference.

- [prefix_list](data-sources--network_policy--reference--group-001.md#canonical-8506c2bb398101237dca209fa27590971ddfb899f33b3040abf7bcedd97bec3d): complete subsection reference.

<a id="canonical-fff57083706dfaf1ba05f44ca0e9dffffbc4bb6a0ec61a404a574bbd69300666"></a>

## Next pages — endpoint / d4cc85ecb73d / 4

- [endpoint.any](data-sources--network_policy--reference--group-001.md#canonical-50a39d41fb47657ce2506798a62dfcaf587e52cc8cc7e1c01307349e91a8a719)
- [endpoint.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-c4e6b1708353a516a8e8cdae1a6723fe32349312f1150f412c73a2f8afe6856d)
- [endpoint.label_selector](data-sources--network_policy--reference--group-001.md#canonical-8c8df0b43d3d9600043ed818f9ac559b26b56d30aff813497205a8428e272c0d)
- [endpoint.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-8207789040ba8897674ff5fcd7074e3612ba31419323b8b57d8286a3bdc613b1)
- [endpoint.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-8506c2bb398101237dca209fa27590971ddfb899f33b3040abf7bcedd97bec3d)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-50a39d41fb47657ce2506798a62dfcaf587e52cc8cc7e1c01307349e91a8a719"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7503f0706c1050810187b4e9824902eb5dc5b7b3e3d902157fcf12ef9e741715"></a>

## endpoint.any — endpoint.any / f18b2481d25d / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- endpoint.any

<a id="canonical-ee1c68e606c4abc83dd503ef27359f04bb5466a34f5519eba83115888bf79891"></a>

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

<a id="canonical-70133d12d215e25382d5975fc3bbc3ceb16d4b8dbf7101bdf659887d28a2ea0c"></a>

## Direct properties — endpoint.any / f18b2481d25d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-99f0d856f9ec164d7f8614debe7885b9c29a1b570337f3436047fb63c147060d"></a>

## Next pages — endpoint.any / f18b2481d25d / 4

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-c4e6b1708353a516a8e8cdae1a6723fe32349312f1150f412c73a2f8afe6856d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47da97c7125ff6a3a41f7d1c74013066337f6c75c35236af232cc721eff45172"></a>

## endpoint.inside_endpoints — endpoint.inside_endpoints / 0dfc37fcec20 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- endpoint.inside_endpoints

<a id="canonical-81232ea0444f95d6d664d2deed8f1ee7806030416237b9ba149570ec446cf811"></a>

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

<a id="canonical-84e417c31cb74bc0a69b478dacaff666987c400819d59e1f71c9d4e5b927f358"></a>

## Direct properties — endpoint.inside_endpoints / 0dfc37fcec20 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3623823796388dfa026e3437384dac52d984a0d648bb40517cda29c93a56b105"></a>

## Next pages — endpoint.inside_endpoints / 0dfc37fcec20 / 4

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-8c8df0b43d3d9600043ed818f9ac559b26b56d30aff813497205a8428e272c0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-578ce9380119d9da5755bd6dd96f67b452db7334bb366453bdcd92aff9998b76"></a>

## endpoint.label_selector — endpoint.label_selector / d01f068f13a5 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- endpoint.label_selector

<a id="canonical-71792c139b6373f043f928aad7ee903264030cc989b39f516f96b0bd53e81bc9"></a>

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

<a id="canonical-c089b7357ba082107e37408eec83b348a5ceca1636cdb0a28b4b219c548f2102"></a>

## Direct properties — endpoint.label_selector / d01f068f13a5 / 3

<a id="canonical-6bb373b5b2789c5fbada3827c2adf6d29bc80a73ce7daff199d6e22ccbd20d48"></a>

<a id="canonical-6bda65d8d6842d0406ab07d5e7b7f2f4121d56f40621b21b0b7a32b01cd0aadd"></a>

## expressions property — endpoint.label_selector / d01f068f13a5 / 4

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

<a id="canonical-b85b48ee26a0e0c0367ba8449dd7cec88765cafebe36a2b980d518bb664392d7"></a>

## Next pages — endpoint.label_selector / d01f068f13a5 / 5

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-8207789040ba8897674ff5fcd7074e3612ba31419323b8b57d8286a3bdc613b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4d1aa0744fbcd0e6b81acf2b2a6ac7aeb35c3c7b070de772b5771f3759e1157"></a>

## endpoint.outside_endpoints — endpoint.outside_endpoints / 6264970ae301 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- endpoint.outside_endpoints

<a id="canonical-df5fc2c8903517a2c42503f0d52314892036f54ed02716faab3b3628f0805741"></a>

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

<a id="canonical-e48df01dd345e3a8f2c18141e4e246a264a39958030cee9ec4244bf308a82d27"></a>

## Direct properties — endpoint.outside_endpoints / 6264970ae301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b231622e26b4d14a61730558392eb1f63a0cfdf4e4404d1f7a80db914af55b68"></a>

## Next pages — endpoint.outside_endpoints / 6264970ae301 / 4

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-8506c2bb398101237dca209fa27590971ddfb899f33b3040abf7bcedd97bec3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-999fe5801dc88bdeaa81feafadcd621dc897f64ad9c7b9bab1b464e9494382cd"></a>

## endpoint.prefix_list — endpoint.prefix_list / 728e64fa6e54 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- endpoint.prefix_list

<a id="canonical-92958dc4d817973d7d9d37682ae42b3904f5f466c10db183d9f8180ac255f001"></a>

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

<a id="canonical-1728e0998989160f12bca99a02c1be19f81d98084c6ad2a4b58f80caafadd627"></a>

## Direct properties — endpoint.prefix_list / 728e64fa6e54 / 3

<a id="canonical-d1a58e411af3228f6faa547621cd58812a055e8295ba91a3e921657496b7ff6f"></a>

<a id="canonical-95e5cb140f12b80b11eb85e525fdaac4bc6bacf4154eb5454df05a51b0f67df5"></a>

## prefixes property — endpoint.prefix_list / 728e64fa6e54 / 4

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

<a id="canonical-63ef8ab8c9a75a109ccde3231ec5cb154f183a0a26de78462617134f7ebf63bd"></a>

## Next pages — endpoint.prefix_list / 728e64fa6e54 / 5

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-841a1bb8841faf9a079dd84fc2707b9ce839fda8f80eb662e06ec8afec9728fe)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7422d6e9c88ede89d69f81df94ea94d77e2121ada62f99d89d55d9091aff5af2"></a>

## rules — rules / ea6d06b146b1 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- rules

<a id="canonical-b3cd9cd411d253ba6a2f801964b6623cacb8844b6212586d0ef6922fdff6e8df"></a>

Type: `"single"`. Computed.

Rule Choice. Shape of Rule Choice.

Upstream description:

Shape of Rule Choice.

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

<a id="canonical-abb811dee3dfe4815ba2fbe6447afdd02ec6cb0f76c92ee57c8caff77bf7dd51"></a>

## Direct properties — rules / ea6d06b146b1 / 3

- [egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b): complete subsection reference.

- [ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362): complete subsection reference.

<a id="canonical-9ed6aae2a3a7d8811949499e026c1e07f1d16010b66d8b85fb893a341f973b20"></a>

## Next pages — rules / ea6d06b146b1 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6aa7fad10c2669035857b33dd815e3ca77bca07796c12ea86ef16bd2b819b3d5"></a>

## rules.egress_rules — rules.egress_rules / 7d0fd88ca199 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- rules.egress_rules

<a id="canonical-c87f7c38587a4fcc3b12c8e1f1f4b8509c241ba99e3ca6289b6e262d6a34c7cc"></a>

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

<a id="canonical-074389814e6d9140a37efd47a1430379611bffeac7c77e86fc6d0cf377eee144"></a>

## Direct properties — rules.egress_rules / 7d0fd88ca199 / 3

<a id="canonical-a50208f958d53dbbbae1cf2c4c01647e1b633a3c5d927d161fb0d12089a55477"></a>

<a id="canonical-03823fdbc0c986daff103d0241ce73cf4fca22021023c6bae7e69d1f2e3f1a80"></a>

## action property — rules.egress_rules / 7d0fd88ca199 / 4

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

- [adv_action](data-sources--network_policy--reference--group-001.md#canonical-1e27a4d78b806500ce6b524a21bd3e9786280e682752d33a1fe6c43bd9369d06): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-cc8c3407a5b5d47bc5a8f81d16d2ce2f6f8b9144ccf8b4b446d6f934d19131ed): complete subsection reference.

- [all_traffic](data-sources--network_policy--reference--group-001.md#canonical-3db6e7f5badc948c8c243ed74b294c02c3e3833f145d60630b0e80e63d91527d): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-0a5b2cbd5faa47c52931187bf02a888b85dd9962623577457ec9a99d33d5ec76): complete subsection reference.

- [any](data-sources--network_policy--reference--group-001.md#canonical-60717f205c2b0c58f08386a81785975f6025c17568190ffafe23ec7729387331): complete subsection reference.

- [applications](data-sources--network_policy--reference--group-001.md#canonical-d26ddbcd98457b1708268ae331f1009dea61fe692810b3873dfbbd22f8090a0d): complete subsection reference.

- [inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-e4e8c16a27ece7f70b8fd8bd9626ca149802b8a0ab9e6b68808f35914c813ffb): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-76d94b8d63c573f08eb66a4d87bb52253830466bfeb294fd55ff3ec3f8a20928): complete subsection reference.

- [label_matcher](data-sources--network_policy--reference--group-001.md#canonical-d3622f68c8ab8253edb727a129558133e1fc9d429a0bd3a039241b33ecc55fe4): complete subsection reference.

- [label_selector](data-sources--network_policy--reference--group-001.md#canonical-1f9d74db5dfda525d778aebd4a2ec61c51cb99f7831e2ea23a8ae89b5ba05efc): complete subsection reference.

- [metadata](data-sources--network_policy--reference--group-001.md#canonical-3442990451a12946cf8b8391065c7cfe6c8f45ec5bcd54e5d63121b0fd50373a): complete subsection reference.

- [outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-083dee444c34b393ca039532a459f70d9a68c17469536359a16999c0fd8f274c): complete subsection reference.

- [prefix_list](data-sources--network_policy--reference--group-001.md#canonical-90289b181fecdcbe30ba901da9bf0b3b4fda347d72005a06298a368f0c066c90): complete subsection reference.

- [protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-0dcaa8e615e76752e40ee528100cef854ac5797bab86b9a3d55410b3aaedfd43): complete subsection reference.

<a id="canonical-b022ee5a9df08963094e374c044ccc8bf3904e035116125b87926d45f7d55bc1"></a>

## Next pages — rules.egress_rules / 7d0fd88ca199 / 5

- [rules.egress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-1e27a4d78b806500ce6b524a21bd3e9786280e682752d33a1fe6c43bd9369d06)
- [rules.egress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-cc8c3407a5b5d47bc5a8f81d16d2ce2f6f8b9144ccf8b4b446d6f934d19131ed)
- [rules.egress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-3db6e7f5badc948c8c243ed74b294c02c3e3833f145d60630b0e80e63d91527d)
- [rules.egress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-0a5b2cbd5faa47c52931187bf02a888b85dd9962623577457ec9a99d33d5ec76)
- [rules.egress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-60717f205c2b0c58f08386a81785975f6025c17568190ffafe23ec7729387331)
- [rules.egress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-d26ddbcd98457b1708268ae331f1009dea61fe692810b3873dfbbd22f8090a0d)
- [rules.egress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-e4e8c16a27ece7f70b8fd8bd9626ca149802b8a0ab9e6b68808f35914c813ffb)
- [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-76d94b8d63c573f08eb66a4d87bb52253830466bfeb294fd55ff3ec3f8a20928)
- [rules.egress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-d3622f68c8ab8253edb727a129558133e1fc9d429a0bd3a039241b33ecc55fe4)
- [rules.egress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-1f9d74db5dfda525d778aebd4a2ec61c51cb99f7831e2ea23a8ae89b5ba05efc)
- [rules.egress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-3442990451a12946cf8b8391065c7cfe6c8f45ec5bcd54e5d63121b0fd50373a)
- [rules.egress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-083dee444c34b393ca039532a459f70d9a68c17469536359a16999c0fd8f274c)
- [rules.egress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-90289b181fecdcbe30ba901da9bf0b3b4fda347d72005a06298a368f0c066c90)
- [rules.egress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-0dcaa8e615e76752e40ee528100cef854ac5797bab86b9a3d55410b3aaedfd43)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-1e27a4d78b806500ce6b524a21bd3e9786280e682752d33a1fe6c43bd9369d06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da43cf155897de3c14c223f1d55926988aa36b8bf3662903d959b09f457656fb"></a>

## rules.egress_rules.adv_action — rules.egress_rules.adv_action / c01aab85f3f9 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.adv_action

<a id="canonical-933fe72a647d22a6f740984ed6cf4fec14623a14cecdbc87d5acab96ec24caa3"></a>

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

<a id="canonical-b1bebce8288d4b0d0c31ce253356a0360ddb200d36b915034b0e569f852123dd"></a>

## Direct properties — rules.egress_rules.adv_action / c01aab85f3f9 / 3

<a id="canonical-d48a3ef4f8399c578d11c80908faca0b18c796b89123590aa46630777fff60f8"></a>

<a id="canonical-857af35fa01614cd616ad5313b0e6f51319e0a909e728569b5e7b4b0e27252ae"></a>

## action property — rules.egress_rules.adv_action / c01aab85f3f9 / 4

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

<a id="canonical-4a387bb2996bfbba0f9ab0ef29142e6274ef8b6521de5abbbba4cc3dfc10ccf1"></a>

## Next pages — rules.egress_rules.adv_action / c01aab85f3f9 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-cc8c3407a5b5d47bc5a8f81d16d2ce2f6f8b9144ccf8b4b446d6f934d19131ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78fc1cb1235d0e367509b15918ef8a6fdf0643f77f385d59be8536d258f853b9"></a>

## rules.egress_rules.all_tcp_traffic — rules.egress_rules.all_tcp_traffic / f800cd2ecc13 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.all_tcp_traffic

<a id="canonical-d95d1fb1037cc6a8331eaf677bf4c214c5706b8c8261563c23e3d2b3862b3af6"></a>

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

<a id="canonical-b2213e7252102d21446ee71fee9f352bf755269a1c637bfc32749050837c7908"></a>

## Direct properties — rules.egress_rules.all_tcp_traffic / f800cd2ecc13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a72f7e34b8d9c01a4a589aaf936419322b6fba630025e36ee79582933392442"></a>

## Next pages — rules.egress_rules.all_tcp_traffic / f800cd2ecc13 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-3db6e7f5badc948c8c243ed74b294c02c3e3833f145d60630b0e80e63d91527d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8678d814f5b053c4973efa998a3927aa7db0e7b5c6dc1a05d815cc9cb93db67"></a>

## rules.egress_rules.all_traffic — rules.egress_rules.all_traffic / 6dd5ef7c0cf8 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.all_traffic

<a id="canonical-d56d340bf81bb6ef0faf4109d0aebfeea9eba960af1f6d43c7066771899206f6"></a>

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

<a id="canonical-008535736787085282ce981ad8188768b9b198bada76af096ec3272b51f815e8"></a>

## Direct properties — rules.egress_rules.all_traffic / 6dd5ef7c0cf8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f6d86b61b70c9fb253e5025fb1ad71bcf7bb7c51712e646f9d2f4a9d6c72330"></a>

## Next pages — rules.egress_rules.all_traffic / 6dd5ef7c0cf8 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-0a5b2cbd5faa47c52931187bf02a888b85dd9962623577457ec9a99d33d5ec76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31eeaf5f713a554ec4172ea1edfa370059729a51f45b9681b0c71ce547bcec51"></a>

## rules.egress_rules.all_udp_traffic — rules.egress_rules.all_udp_traffic / 67bb9ea6f890 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.all_udp_traffic

<a id="canonical-8bb01550726d18051b2d79737f279b7b9c6a92eba2776e7d21b2ff8b31e3875f"></a>

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

<a id="canonical-be8f2ec98165d95231c6f9dcb44a1746bd358d351e94efb072143beef09196f0"></a>

## Direct properties — rules.egress_rules.all_udp_traffic / 67bb9ea6f890 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb21490b4c38159ce2c4e355e2527acccd84ebdb95b27cf775ff550609522672"></a>

## Next pages — rules.egress_rules.all_udp_traffic / 67bb9ea6f890 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-60717f205c2b0c58f08386a81785975f6025c17568190ffafe23ec7729387331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57369968e04808358f3d46a47a466b3932ddb9365b8bdd92b1a51f1cceda6f06"></a>

## rules.egress_rules.any — rules.egress_rules.any / 76bc92ac153e / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.any

<a id="canonical-f8ee290f972c9b93704f865085f9a2c0ee86fdb8fb40dd74a853ef101c2cb0b5"></a>

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

<a id="canonical-471db5ae91363bfb68ac085c481501915f8c2b1894531cf68999d4ce16462bff"></a>

## Direct properties — rules.egress_rules.any / 76bc92ac153e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9f7699c5bc72e5ff1951df3409657227a9156b5ffcb131cedb838e3261dc8ac"></a>

## Next pages — rules.egress_rules.any / 76bc92ac153e / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-d26ddbcd98457b1708268ae331f1009dea61fe692810b3873dfbbd22f8090a0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-818ea3b02e86c6364b8a4b71184f74bb90a60eee2100bd935d10679e91223cc5"></a>

## rules.egress_rules.applications — rules.egress_rules.applications / e3125e066b6b / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.applications

<a id="canonical-a701a973484a9a05701c44f0ba7c42a95be6d1bca52682d1e68eb97ca587b004"></a>

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

<a id="canonical-6fe3f46cb668f2fc08a47be2de79c9d3e41b680c914d72420fb27bd882d6302a"></a>

## Direct properties — rules.egress_rules.applications / e3125e066b6b / 3

<a id="canonical-1aa67850296bf610c540514a06b9e16bcd1cbc2b7aaa4a508368a6503cead5eb"></a>

<a id="canonical-6a6636329241ed10624ae7201a0f6b324fa79ed0ae8543412e22322a95a06fda"></a>

## applications property — rules.egress_rules.applications / e3125e066b6b / 4

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

<a id="canonical-d38f212a5fbb823d3b9963e7da7d2862cd0fcd5bb945f6208d595f7f2b94cfd8"></a>

## Next pages — rules.egress_rules.applications / e3125e066b6b / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-e4e8c16a27ece7f70b8fd8bd9626ca149802b8a0ab9e6b68808f35914c813ffb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53847cd10f1d995239f5ac26538a7cca4df8c1d4980afe1d650e18d38665a6a1"></a>

## rules.egress_rules.inside_endpoints — rules.egress_rules.inside_endpoints / 29e443f23eaf / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.inside_endpoints

<a id="canonical-19798de2269c957e5ecd4b12e6252ab4bc904efb9e20f7b1c455a2c76cc1e925"></a>

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

<a id="canonical-9fb37740f839b45f20bdf703e1bd01109567bfc5b9c032112b948df613645c00"></a>

## Direct properties — rules.egress_rules.inside_endpoints / 29e443f23eaf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c37005ae3cddc054f94f192ea256ca22a512639fbbcba1d8f0e52e0d8b45c82"></a>

## Next pages — rules.egress_rules.inside_endpoints / 29e443f23eaf / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-76d94b8d63c573f08eb66a4d87bb52253830466bfeb294fd55ff3ec3f8a20928"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b74a0e7d0a5ac94fcc6e67edfbfc37acbfe49bd50915719cef18ab0df0bdb25"></a>

## rules.egress_rules.ip_prefix_set — rules.egress_rules.ip_prefix_set / 41faff431e97 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.ip_prefix_set

<a id="canonical-6e813d9a3a9743fe54d766ae6f77edf55d582f1d6882978ddba8685c7377063d"></a>

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

<a id="canonical-c50d641db0776064580b2a4a1192ffd417847889b05ae0d08b7228d8f1db1d61"></a>

## Direct properties — rules.egress_rules.ip_prefix_set / 41faff431e97 / 3

- [ref](data-sources--network_policy--reference--group-001.md#canonical-aa37c8b0e197ca2ab5c37bdedec7f0e9eb3047f3c02d6631300736260500f571): complete subsection reference.

<a id="canonical-f58ef091e0c2e4e00e18f6914f3eafada16e26cf61d2e1cbacbf8f592667d39c"></a>

## Next pages — rules.egress_rules.ip_prefix_set / 41faff431e97 / 4

- [rules.egress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-aa37c8b0e197ca2ab5c37bdedec7f0e9eb3047f3c02d6631300736260500f571)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-aa37c8b0e197ca2ab5c37bdedec7f0e9eb3047f3c02d6631300736260500f571"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-854d0ecb7670fdaa69d894740b13fa1504ee289df808991c86f060620444e44c"></a>

## rules.egress_rules.ip_prefix_set.ref — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-76d94b8d63c573f08eb66a4d87bb52253830466bfeb294fd55ff3ec3f8a20928)
- rules.egress_rules.ip_prefix_set.ref

<a id="canonical-14beb053eac489635b5194be4064277a36497874c8c3d66887ed59ca3b141a2b"></a>

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

<a id="canonical-1dc3e1b37872fc8dbc704a0a2a4b1f358ed0a395fb32fccffab7b9a0f7065ff2"></a>

## Direct properties — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 3

<a id="canonical-5c9234d4e33d601478d7dd03c04a5c0a8543730ace3703dacbd8a1aeb5d1c317"></a>

<a id="canonical-9c0ea92e95bdfaf68fefdb5aaa99821419e652e1eafeda00214886380953df1b"></a>

## kind property — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 4

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

<a id="canonical-d458602bce2b2601043ec6595f2632342a19811e5f70f09940a2e288cd683f87"></a>

<a id="canonical-e354109e97673df831a34f4ce89e3aa07bc136ec18bc2fc71472c7fc61251a62"></a>

## name property — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 5

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

<a id="canonical-18ab5f9f5ec707a5908a208437a3661c1bb97ed05127f90c958cf637d67b43a1"></a>

<a id="canonical-59fbd7da725f1a8cc4af468af1ce76a89d2e218333273644a29bd2dba4e95948"></a>

## namespace property — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 6

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

<a id="canonical-a537266f3a2964965b2969a65da566b5adda592b4e2ac2a572be8adaa5493ae7"></a>

<a id="canonical-e8b525a930cfed9e3c7a37937e1564124e6d9f208a480a99c09ba071836cfe8e"></a>

## tenant property — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 7

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

<a id="canonical-46c6231a867f1a0f7d104b19a135495db687c6efd3c0af50ca1dd304735163f1"></a>

<a id="canonical-cf20f8ec7e12e76b363546a4c6a797093a7bfbba44516d3d332d65b2370bf70c"></a>

## uid property — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 8

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

<a id="canonical-5178d921f6c3d5cb8daa4cf3f6a403cb71a5d9433022b7409429948033b5178e"></a>

## Next pages — rules.egress_rules.ip_prefix_set.ref / e993782aa418 / 9

- [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-76d94b8d63c573f08eb66a4d87bb52253830466bfeb294fd55ff3ec3f8a20928)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-d3622f68c8ab8253edb727a129558133e1fc9d429a0bd3a039241b33ecc55fe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c617d8c778b2f27c7b95c9be5e04161ff73b4d21ef797b0100fef3734790916"></a>

## rules.egress_rules.label_matcher — rules.egress_rules.label_matcher / fa5deba00b9c / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.label_matcher

<a id="canonical-df27ad6fd539cc0ad1db85d2e530f7ce3ec78c6aa2546fcea813075bba43655e"></a>

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

<a id="canonical-a9f77b6e346ae5636d6c36f18f1f07e8d96ed565d3833e0927edfbbcf1f6a5fa"></a>

## Direct properties — rules.egress_rules.label_matcher / fa5deba00b9c / 3

<a id="canonical-006f2e3d3ef3090329cdf3af531686623afbd11c680a3479c2be50d5b653524e"></a>

<a id="canonical-a9ddde9010192849ffcaf21de792a10b2971448580af4c3e4a9a8e1b4567d7ee"></a>

## keys property — rules.egress_rules.label_matcher / fa5deba00b9c / 4

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

<a id="canonical-741166154b19a1be64e1266895665145d6e0f4c7acb31905e4ed41077a98769f"></a>

## Next pages — rules.egress_rules.label_matcher / fa5deba00b9c / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-1f9d74db5dfda525d778aebd4a2ec61c51cb99f7831e2ea23a8ae89b5ba05efc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d08d6b7889d119f37109b30f63830ee51726507e8776680314ccf8aaf8d0e8c"></a>

## rules.egress_rules.label_selector — rules.egress_rules.label_selector / 45d21f4a9199 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.label_selector

<a id="canonical-579ac2dedd2152d0de46c67689b9fccd7ccf93dc152bf4a749c90a71bee9a4fc"></a>

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

<a id="canonical-f666822008670334dd91ab9c345a67a02c9342f20381d1c7bf3a27471fea3ca6"></a>

## Direct properties — rules.egress_rules.label_selector / 45d21f4a9199 / 3

<a id="canonical-c342083d5b6832d689a86419a2807e0c7fed8548b6c4357b0f7e145b1256679e"></a>

<a id="canonical-eda99a26b20ae687d16497696d3b993b4452acce8be7e7c72ef249d59e1079d0"></a>

## expressions property — rules.egress_rules.label_selector / 45d21f4a9199 / 4

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

<a id="canonical-7974c4f63cfc696cb10d4631cc60fd273829f068b62ef0f4e09cd5e7067c722b"></a>

## Next pages — rules.egress_rules.label_selector / 45d21f4a9199 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-3442990451a12946cf8b8391065c7cfe6c8f45ec5bcd54e5d63121b0fd50373a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bd1f6d121d6ef457260d81e7aa02fa9117ea61341931e92fdfa0bd60f1f37f9"></a>

## rules.egress_rules.metadata — rules.egress_rules.metadata / 370e7b65b049 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.metadata

<a id="canonical-ef1a859c9eeb2fd765d1222e2557d03494513d73cd1f11d27f46961fe241469e"></a>

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

<a id="canonical-a317bb3440fb02cb48b63aecf5be941695d02cc01df55c076967e9e575528a92"></a>

## Direct properties — rules.egress_rules.metadata / 370e7b65b049 / 3

<a id="canonical-52637feccce34ea10980fbd8e7832fa2fe2e8bbb008f9ce956911e7a4af8bc45"></a>

<a id="canonical-878909008802e70c7eadc6b7488c17f04af257ee30fe0947f0d8422e40055cc9"></a>

## description_spec property — rules.egress_rules.metadata / 370e7b65b049 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1804e4f8d5c66542835f2ebb2e66e896256d19d1549b0ea9909169af819ce196"></a>

<a id="canonical-6ea22399f199fa72708c8501ebae0ca018da4bfa16a2cb1a2ed6a26d3661eacc"></a>

## name property — rules.egress_rules.metadata / 370e7b65b049 / 5

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

<a id="canonical-05788a22ddc50bc0d6dcf30671b4fbd16b4a2f822f1e4848ffa9cffaba37436e"></a>

## Next pages — rules.egress_rules.metadata / 370e7b65b049 / 6

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-083dee444c34b393ca039532a459f70d9a68c17469536359a16999c0fd8f274c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-226412f20b28d83c749399f9fa83f336a8228161e1d158261c0c823b18c373a1"></a>

## rules.egress_rules.outside_endpoints — rules.egress_rules.outside_endpoints / c9cd85604274 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.outside_endpoints

<a id="canonical-05d4f7068c498e764090243830971bbfe708b991c9803e963913e98105d03b10"></a>

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

<a id="canonical-58637b200e3b3e39ec47c4ebabf5894055c5b71e85dad177ec8cb5a7edd7183c"></a>

## Direct properties — rules.egress_rules.outside_endpoints / c9cd85604274 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-255c5e6ec34851656cf9c219c59da414dda8edcbe3857ca971d29e06f843cca3"></a>

## Next pages — rules.egress_rules.outside_endpoints / c9cd85604274 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-90289b181fecdcbe30ba901da9bf0b3b4fda347d72005a06298a368f0c066c90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88657092ecaadb0344fd6f6497bf12fa88215d8b0b7747218f1e5710b9d51492"></a>

## rules.egress_rules.prefix_list — rules.egress_rules.prefix_list / 3cd10a4271c2 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.prefix_list

<a id="canonical-36c7dbe4ef760013396af538370efd22d000401301a8ba66c18262bb45376147"></a>

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

<a id="canonical-a2c035df86edb2a07f413a5f4c6af4324163fa5f0423fe163efee000f02205cd"></a>

## Direct properties — rules.egress_rules.prefix_list / 3cd10a4271c2 / 3

<a id="canonical-7f2592aa50eebb17b29098721d52940305f5bd57240d61318fe27b54a735aae6"></a>

<a id="canonical-ec63e1d3ae6d59b92658943a5da73cbe70b1faf74ec529fe5df846d720f2d8de"></a>

## prefixes property — rules.egress_rules.prefix_list / 3cd10a4271c2 / 4

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

<a id="canonical-d36dfc421f93ba7196eeb20cdfda157c2a2c1b755a029280a39e7c98bbf764ba"></a>

## Next pages — rules.egress_rules.prefix_list / 3cd10a4271c2 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-0dcaa8e615e76752e40ee528100cef854ac5797bab86b9a3d55410b3aaedfd43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c310b58e22f13177c34a01947c1c94103d30dab740958f42b47b5adb66862b4d"></a>

## rules.egress_rules.protocol_port_range — rules.egress_rules.protocol_port_range / 642657971328 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- rules.egress_rules.protocol_port_range

<a id="canonical-5439c6aca8c38ede23cd727c71e8d9b18eb1133a06dc168e966c08e3553058a8"></a>

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

<a id="canonical-d498b6529e519900f69db8ad8be472df6aa11df86ae79d9e1d4d00b34991020f"></a>

## Direct properties — rules.egress_rules.protocol_port_range / 642657971328 / 3

<a id="canonical-bd90aa02feb18a9c82659904db4197c8a27cef40adc0734e228c163ad486ea66"></a>

<a id="canonical-11bc769f0b95a85c46442b51e685cba512283fa44d4fb10a51eaf064d646a957"></a>

## port_ranges property — rules.egress_rules.protocol_port_range / 642657971328 / 4

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

<a id="canonical-002773e2b30e2f49617c51cb9c56fb25ff7e61501ac103994317a9af88e06865"></a>

<a id="canonical-37c3d28b1fed04328d44ca01956006fa3ba9267654ab1b2e9b8ac72127bf2f43"></a>

## protocol property — rules.egress_rules.protocol_port_range / 642657971328 / 5

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

<a id="canonical-9a798002e8b936a95c1c57ca27bf09aea4333a275290919ca11647b79a2ede6e"></a>

## Next pages — rules.egress_rules.protocol_port_range / 642657971328 / 6

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0a2d5db104673d20964cf539335c1d2761d74270f7f6f25fc497b39552f8703b)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64a3abff958198451308849baf10e7a93648421628dd261756af89821b91599d"></a>

## rules.ingress_rules — rules.ingress_rules / 12bf282e813d / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- rules.ingress_rules

<a id="canonical-2a5516510991f6cf289290609c42563b1f730802b624b8d701aa7be17f8eaaa6"></a>

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

<a id="canonical-1ed732db882960331ef6bd0cb43eaacc0c018c60dbc5846eb7d23db72722c7dd"></a>

## Direct properties — rules.ingress_rules / 12bf282e813d / 3

<a id="canonical-96edf5952f78e1fbcdf331c5ea752f7b9792d7558d016f6687d1f4c6913bae9d"></a>

<a id="canonical-97c40ecbc9a49cc72a3e8da29ed050c785109f7801a9f9e066253555d4b0fca0"></a>

## action property — rules.ingress_rules / 12bf282e813d / 4

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

- [adv_action](data-sources--network_policy--reference--group-001.md#canonical-3eb378261771282a62b2874e951ffa03db3810e409fc77537e08df951dc23f31): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-41b389c3b6d2a02e29841780159d2e56701f669c050dec97c17b2f0d5b264f0a): complete subsection reference.

- [all_traffic](data-sources--network_policy--reference--group-001.md#canonical-f746170673b956d87283860398427ee72fa643ef988824fc758ee4e4f48b1264): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-a408899042d659c6ccdca6d5a1081a0cc1032ebe99a74021571369ec3c499a77): complete subsection reference.

- [any](data-sources--network_policy--reference--group-001.md#canonical-98c05864c8e446985f0c866f0bf966990c6599842c4574f225553defc475086a): complete subsection reference.

- [applications](data-sources--network_policy--reference--group-001.md#canonical-74fa0607b68f1e6034951d35fb97640b2aecd91c6686856d77f7ea91d88a9f10): complete subsection reference.

- [inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-a3de765a765ccf3c9a0fcc2da5b94e6ba4a9fdc70843a53ec5b4c7c56345aad0): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-38977cabc31c0d2a02db2fec07b03e54311cffec0729fee794f34c93ca5e8a66): complete subsection reference.

- [label_matcher](data-sources--network_policy--reference--group-001.md#canonical-71bf275ec560d79e6cdca72137ae789903abe5a65e59217b5ed455cf058b6f34): complete subsection reference.

- [label_selector](data-sources--network_policy--reference--group-001.md#canonical-d7a24f07dd4ec56799af683d0007b1f737780ee6b32212e3f8b80f0c29134c24): complete subsection reference.

- [metadata](data-sources--network_policy--reference--group-001.md#canonical-0ed36d464e3c23b38fb93d609ce7499671a879700e65bfc4352d6846b349eabe): complete subsection reference.

- [outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-762b5fcada68e2adeb3bafcfa902a48fc244b26483c5d1855f531cd190724e17): complete subsection reference.

- [prefix_list](data-sources--network_policy--reference--group-001.md#canonical-e1eb5970f48ec99146e22c4db25038206953350a3256e9dfff54289c0303fdb0): complete subsection reference.

- [protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-d1448ab984fb119569c3cfb0e3608da7525b16979fddba0b5af0f1e8a7290615): complete subsection reference.

<a id="canonical-d7af3c9a9f07c5096edc00f24a04059e3c0b6347c97d221b80081341c4d66d23"></a>

## Next pages — rules.ingress_rules / 12bf282e813d / 5

- [rules.ingress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-3eb378261771282a62b2874e951ffa03db3810e409fc77537e08df951dc23f31)
- [rules.ingress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-41b389c3b6d2a02e29841780159d2e56701f669c050dec97c17b2f0d5b264f0a)
- [rules.ingress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-f746170673b956d87283860398427ee72fa643ef988824fc758ee4e4f48b1264)
- [rules.ingress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-a408899042d659c6ccdca6d5a1081a0cc1032ebe99a74021571369ec3c499a77)
- [rules.ingress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-98c05864c8e446985f0c866f0bf966990c6599842c4574f225553defc475086a)
- [rules.ingress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-74fa0607b68f1e6034951d35fb97640b2aecd91c6686856d77f7ea91d88a9f10)
- [rules.ingress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-a3de765a765ccf3c9a0fcc2da5b94e6ba4a9fdc70843a53ec5b4c7c56345aad0)
- [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-38977cabc31c0d2a02db2fec07b03e54311cffec0729fee794f34c93ca5e8a66)
- [rules.ingress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-71bf275ec560d79e6cdca72137ae789903abe5a65e59217b5ed455cf058b6f34)
- [rules.ingress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-d7a24f07dd4ec56799af683d0007b1f737780ee6b32212e3f8b80f0c29134c24)
- [rules.ingress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-0ed36d464e3c23b38fb93d609ce7499671a879700e65bfc4352d6846b349eabe)
- [rules.ingress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-762b5fcada68e2adeb3bafcfa902a48fc244b26483c5d1855f531cd190724e17)
- [rules.ingress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-e1eb5970f48ec99146e22c4db25038206953350a3256e9dfff54289c0303fdb0)
- [rules.ingress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-d1448ab984fb119569c3cfb0e3608da7525b16979fddba0b5af0f1e8a7290615)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-3eb378261771282a62b2874e951ffa03db3810e409fc77537e08df951dc23f31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f95ecfc93be8b0091052b8b19dddd46fd8ecdb133a040fba657f7f978235f685"></a>

## rules.ingress_rules.adv_action — rules.ingress_rules.adv_action / c9ac439c0162 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.adv_action

<a id="canonical-e5351c3dfe6d2615a5cebf1d3b67e508825c17a62742bc0dbb0c5cf4fd633780"></a>

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

<a id="canonical-3a65b0d1cd3d84abfa0dd17e800207ded35b0cbc9cb5bc3496b182ab358a4ae3"></a>

## Direct properties — rules.ingress_rules.adv_action / c9ac439c0162 / 3

<a id="canonical-bfd16925b753aabdbff0754f5dd2e8fea655d0af51f670cba6b380f6bd53f444"></a>

<a id="canonical-9cf63c3e256a19faa98cf6eb624d19f8c503169f967636176546fa8a22c09245"></a>

## action property — rules.ingress_rules.adv_action / c9ac439c0162 / 4

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

<a id="canonical-66da23e69868d54cb9970590ce68eccc432ffa7d794cc36ad3f66d0ec169033e"></a>

## Next pages — rules.ingress_rules.adv_action / c9ac439c0162 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-41b389c3b6d2a02e29841780159d2e56701f669c050dec97c17b2f0d5b264f0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8a4ccfac85c36911fdfefcb8f5353444db0755e177e8d24f17c2b8886494267"></a>

## rules.ingress_rules.all_tcp_traffic — rules.ingress_rules.all_tcp_traffic / 9054e24be46f / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.all_tcp_traffic

<a id="canonical-97bb353675013d86cd1eccb72c439c79a14dd42a29e9fac2b32642bf7ec05591"></a>

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

<a id="canonical-2725fbf911f46a26e028b70c3f163211f4ac1216d936d6ad82116f90264c5b51"></a>

## Direct properties — rules.ingress_rules.all_tcp_traffic / 9054e24be46f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b52805da627dac4a8949d8538c358e53058c5e818754bc1627a247a6e34a2c32"></a>

## Next pages — rules.ingress_rules.all_tcp_traffic / 9054e24be46f / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-f746170673b956d87283860398427ee72fa643ef988824fc758ee4e4f48b1264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-316b8eca0180cbee2c73874bde5487abd730da61adff4a6cfad43cf222055bad"></a>

## rules.ingress_rules.all_traffic — rules.ingress_rules.all_traffic / 591b613fb7c3 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.all_traffic

<a id="canonical-3bf620acc20a291b4e65ae1347452fcef623f0a6de719b60734ec1bb04fef17a"></a>

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

<a id="canonical-38e66a65819443fd89cdbef7f008cbf35cdba1e94278531fa62ed5c12a420e33"></a>

## Direct properties — rules.ingress_rules.all_traffic / 591b613fb7c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1fd0efbd0833cea4ac8ab4617685e8b86ff56055b3499b1dba905668b450996"></a>

## Next pages — rules.ingress_rules.all_traffic / 591b613fb7c3 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-a408899042d659c6ccdca6d5a1081a0cc1032ebe99a74021571369ec3c499a77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d220017f5022d1231c41696f0f95bc60a9ce2975664b8649a6ac179caf54c377"></a>

## rules.ingress_rules.all_udp_traffic — rules.ingress_rules.all_udp_traffic / 8bc2c46429c1 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.all_udp_traffic

<a id="canonical-9cf95f9037cb40a322b2f4bd0e4a29be9586187bb84e8b98e663126eadb5f9e2"></a>

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

<a id="canonical-6d3f0d247131436c5a8954363ad35b2649eb50039eec795b46e028798cfe2834"></a>

## Direct properties — rules.ingress_rules.all_udp_traffic / 8bc2c46429c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d8c6ab99fc991b6980184b4ab8d131cc264f9f59d24df2d928076f6c3813a80c"></a>

## Next pages — rules.ingress_rules.all_udp_traffic / 8bc2c46429c1 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-98c05864c8e446985f0c866f0bf966990c6599842c4574f225553defc475086a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-424728f682631849791ff2c56b12f4d2c9a179af27c9a1a7a6d495fdcd9d477c"></a>

## rules.ingress_rules.any — rules.ingress_rules.any / 0c8e47835207 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.any

<a id="canonical-0b4f1665819c09b5b30b89e74b934963bdaae5796c48c0a18ff444243c83c35a"></a>

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

<a id="canonical-41e5a1e24ac8f2ea0bb9ce0157ab8fafa51455d08671b338129502efb4c53ce0"></a>

## Direct properties — rules.ingress_rules.any / 0c8e47835207 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fd4c12a7b0c6a6b3cff58d6121264b0077888ff084380190867c1ff0fe9c7f3"></a>

## Next pages — rules.ingress_rules.any / 0c8e47835207 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-74fa0607b68f1e6034951d35fb97640b2aecd91c6686856d77f7ea91d88a9f10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98665938c8a314fac7cf14becfcb1cc3fd14a38f7826e14964d952125b02de97"></a>

## rules.ingress_rules.applications — rules.ingress_rules.applications / 30eb4feab529 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.applications

<a id="canonical-2dfe52b2b0d53cdc4448dbd9d6434e6aac4c2ea6f1b51fb4878f70c0824dd682"></a>

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

<a id="canonical-728b662a51e9313bae6f7dcaa0311ad1b5209d02a2fefac2f2d8230f477dad73"></a>

## Direct properties — rules.ingress_rules.applications / 30eb4feab529 / 3

<a id="canonical-ce0ebb6d86993fc672eb2b349524ed7950c91b91c614bee821af64f2e6bcf860"></a>

<a id="canonical-e7fb0b134566702f5043a632820d65bd041b81e96dc3101b85dfa86a4c7092e6"></a>

## applications property — rules.ingress_rules.applications / 30eb4feab529 / 4

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

<a id="canonical-0fa7a694a6ef33ae2032b4d1fa7b1d7ad4bf494249a440c292a2339d737d631e"></a>

## Next pages — rules.ingress_rules.applications / 30eb4feab529 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-a3de765a765ccf3c9a0fcc2da5b94e6ba4a9fdc70843a53ec5b4c7c56345aad0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b029e5c41de7c14631862db544e4942379adbb014dbec38c27cd9a37f01c9d13"></a>

## rules.ingress_rules.inside_endpoints — rules.ingress_rules.inside_endpoints / 3cffe21bf4d5 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.inside_endpoints

<a id="canonical-ce372acc3ca38a461cb6364bfd8e07ab5dc0fd4cc30b4f43d9b9ebd43ae90110"></a>

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

<a id="canonical-c4885b0bbe43a905f25272519d7f40bfbf93deb2c27427ddba706438652857ed"></a>

## Direct properties — rules.ingress_rules.inside_endpoints / 3cffe21bf4d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7df813a7e6d62d77fb363ed08436aff85a54c726ecfdcad9bf4d2a72a2972037"></a>

## Next pages — rules.ingress_rules.inside_endpoints / 3cffe21bf4d5 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-38977cabc31c0d2a02db2fec07b03e54311cffec0729fee794f34c93ca5e8a66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f95f07aecbbb393da92be85ada6d85276adc072d119afe8fa8ca3a028f6dec1f"></a>

## rules.ingress_rules.ip_prefix_set — rules.ingress_rules.ip_prefix_set / d0219eefeae5 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.ip_prefix_set

<a id="canonical-4d7693e2827731f880943e7bc86e2a08d7fc251b501bae3e1e63ef276beaba26"></a>

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

<a id="canonical-69eff28c332cc896b398791cb7a646ece54626313b97e446da7ece9cf65a2b97"></a>

## Direct properties — rules.ingress_rules.ip_prefix_set / d0219eefeae5 / 3

- [ref](data-sources--network_policy--reference--group-001.md#canonical-b5eb4f96b218be7f57413b5a8711a732a8da474e38480befb0bb85ad23e8beab): complete subsection reference.

<a id="canonical-00836a8f5126a13c5a863aa77a4a810e0dff5f66beb62b156a001e34cf3fefb7"></a>

## Next pages — rules.ingress_rules.ip_prefix_set / d0219eefeae5 / 4

- [rules.ingress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-b5eb4f96b218be7f57413b5a8711a732a8da474e38480befb0bb85ad23e8beab)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-b5eb4f96b218be7f57413b5a8711a732a8da474e38480befb0bb85ad23e8beab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-792de1e7b31ca4e6d254aadcd4dd62df84cbe47c92bd5e5517814dccb25cab29"></a>

## rules.ingress_rules.ip_prefix_set.ref — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-38977cabc31c0d2a02db2fec07b03e54311cffec0729fee794f34c93ca5e8a66)
- rules.ingress_rules.ip_prefix_set.ref

<a id="canonical-134ed8949b5a873dd958cd3ee5bdf56fd945645c5628581b13e915e06117cc15"></a>

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

<a id="canonical-7f882178bf7bcae01699a3f31d9e87490e5281e8f11e1e1b5495e1150e6fa230"></a>

## Direct properties — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 3

<a id="canonical-250e9db41c4a6f851c67c7081a5eba23508bea75b915fe56c9a4146f0c38b5e3"></a>

<a id="canonical-bc9463647f3046885995b7cc76b1a5657dd34864ae85a45fb41a82b2c5f310ec"></a>

## kind property — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 4

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

<a id="canonical-2f0d3469b920ba504c46306ca8860034f948a90f57c24c146db9fc236bdfa92c"></a>

<a id="canonical-3ea7552ace32f41a6218aeee38427eac22c59b1dd8accf773b1096f41ea54635"></a>

## name property — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 5

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

<a id="canonical-844e5075e896f4b6f9d12744ede833fd7c4c9bfe46168aec540dbf50d0d349aa"></a>

<a id="canonical-07df69b39c57fbd312fe6f8fa1e90fb5b528842785f34db960fa373105b5697a"></a>

## namespace property — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 6

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

<a id="canonical-b1a3fcbd1f01969c702e3db7186611cfe62bdef64fd551139f38d9fda1047fdf"></a>

<a id="canonical-4d497a8bde2af3c3182837a80a5971b226a89c347957641318c0669151693055"></a>

## tenant property — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 7

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

<a id="canonical-99fee300caff6cf5fc00757bbe5eaf6deab3f7675aab48a8063a71c9afa775c2"></a>

<a id="canonical-8f5a800c170741b051bbe83a03ddfc672708edf2567b73bb4bf0a6786deca5e8"></a>

## uid property — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 8

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

<a id="canonical-75c688766cd015f02424b47f335e59204cb3642958d0612e6588e05503f2b99b"></a>

## Next pages — rules.ingress_rules.ip_prefix_set.ref / 6660e889ec6c / 9

- [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-38977cabc31c0d2a02db2fec07b03e54311cffec0729fee794f34c93ca5e8a66)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-71bf275ec560d79e6cdca72137ae789903abe5a65e59217b5ed455cf058b6f34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db1f3ae0f045b7cb79c08a3ac4d728a6a3f6c5a2c749bae8a0265a281497c57d"></a>

## rules.ingress_rules.label_matcher — rules.ingress_rules.label_matcher / 907f87dee0bf / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.label_matcher

<a id="canonical-47dcf6dcd413655a0921336deb40dbe4806aa4f846505ee7341322483135a4a4"></a>

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

<a id="canonical-e5fc339ae30d559dfb440362d03a6d1e15552ac6939b3b762236f9739545d515"></a>

## Direct properties — rules.ingress_rules.label_matcher / 907f87dee0bf / 3

<a id="canonical-ec67ae6a315bfc655c1647beee4e84be5448257dd20796c943996b0be88b503d"></a>

<a id="canonical-96fb34f64364474e0b8945853fc6d6dce23471c3b77829e50f2d399192ea1ddc"></a>

## keys property — rules.ingress_rules.label_matcher / 907f87dee0bf / 4

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

<a id="canonical-0da3a72cf94eca5d716e87a6449447efd121010c600fdc52f3766ae89609dac9"></a>

## Next pages — rules.ingress_rules.label_matcher / 907f87dee0bf / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-d7a24f07dd4ec56799af683d0007b1f737780ee6b32212e3f8b80f0c29134c24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19003a3e5dac2134cb506d32b448c4e370432cbca47e02446dd71c7f53c4bc4d"></a>

## rules.ingress_rules.label_selector — rules.ingress_rules.label_selector / 2a7b473fc016 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.label_selector

<a id="canonical-5e26c97bc85da7d028781488ecb027469ab937b97a927e02e3beaaec2d6cb162"></a>

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

<a id="canonical-a2ce5d48fd62711b3e5706f5329813e6ca348ebba2c8372002062f0d84080e85"></a>

## Direct properties — rules.ingress_rules.label_selector / 2a7b473fc016 / 3

<a id="canonical-d1861fef6ba8f48b5e897d332e5b220c5adfd0452e57bacb6dadd5cb3bf309ff"></a>

<a id="canonical-5fd4e092e32b702f027baebe79f79e23ff59e5d22e00d61499394a569e5e8bdd"></a>

## expressions property — rules.ingress_rules.label_selector / 2a7b473fc016 / 4

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

<a id="canonical-78e9090c696d559ae155ff23fdfc6851a22e1da5d19a217e14fbbf6df4224389"></a>

## Next pages — rules.ingress_rules.label_selector / 2a7b473fc016 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-0ed36d464e3c23b38fb93d609ce7499671a879700e65bfc4352d6846b349eabe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c5078286d6edceb04e6863626b08abbadee1114f5db98ea05918fb79512578e"></a>

## rules.ingress_rules.metadata — rules.ingress_rules.metadata / b5cf7a28f5b9 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.metadata

<a id="canonical-2369e0bb4fe4b72586d309e4a3e36d9194d08fd9946f3bf56792f6222f157644"></a>

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

<a id="canonical-32be845a8ecda59f05ac75dcd34ed6a877c5e27308a56d92f5ad7042c0a749db"></a>

## Direct properties — rules.ingress_rules.metadata / b5cf7a28f5b9 / 3

<a id="canonical-45279ddc76cc60228275d2d576ac2290d2ac9f5d190d469cc7fc41d064f9f70d"></a>

<a id="canonical-d9f1d6ecb4721a3c47d5b9ac00b064fbb765df21ffbc0b04d55bc7d66cbcdf40"></a>

## description_spec property — rules.ingress_rules.metadata / b5cf7a28f5b9 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-026f4a0c945817edbc0c9f2114b687465f015285745df3b4ca9d52ce4bf2793c"></a>

<a id="canonical-f6cc6c1ecccd9a7fa47bd7f49a6724d4aa62bd4d4c8f1498b42d79fd926196a5"></a>

## name property — rules.ingress_rules.metadata / b5cf7a28f5b9 / 5

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

<a id="canonical-bf964ad198b0684d80e27c9e4b9cfd48d9b8d2c74cdc3d22443ffe66442ad795"></a>

## Next pages — rules.ingress_rules.metadata / b5cf7a28f5b9 / 6

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-762b5fcada68e2adeb3bafcfa902a48fc244b26483c5d1855f531cd190724e17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12488639013eee0a38fd47398773c840d2f63076ee694df4923035e5da9abe50"></a>

## rules.ingress_rules.outside_endpoints — rules.ingress_rules.outside_endpoints / 0cbb04811487 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.outside_endpoints

<a id="canonical-4aae885c912224649b977c366fd4adcc29aadaf9f3f07946795da4cb53a4ba3a"></a>

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

<a id="canonical-7b20c12fc63832d0f601eef975b705f280344a37a640d206988c6dd1b5bfee1e"></a>

## Direct properties — rules.ingress_rules.outside_endpoints / 0cbb04811487 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0da8b1f6f4fb8b122134187f65015f86bfd05665ce00416298fbb6ca2b9a1016"></a>

## Next pages — rules.ingress_rules.outside_endpoints / 0cbb04811487 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-e1eb5970f48ec99146e22c4db25038206953350a3256e9dfff54289c0303fdb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dab6797e406636e30ef5464b8fac213301b98645fa552d8bfc094e511d7c080f"></a>

## rules.ingress_rules.prefix_list — rules.ingress_rules.prefix_list / ca8903d594c6 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.prefix_list

<a id="canonical-a9223642006d6dcd38a02739a7302439226c3a9692b9e3ff5604e4c715ff0a86"></a>

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

<a id="canonical-44d5fb654dad30f7c1afd5783508364b5e24f982b98eb9d5f7b5f5ade6dee8e9"></a>

## Direct properties — rules.ingress_rules.prefix_list / ca8903d594c6 / 3

<a id="canonical-ba431253a568f86fa95ec4aa125227ad4d525f594d06e8d78e50ff4bfc4f9e2b"></a>

<a id="canonical-a2c43f001c072c4b10a1bc97ede3e96be780b1d19a0e1f90dd7fbc0836289822"></a>

## prefixes property — rules.ingress_rules.prefix_list / ca8903d594c6 / 4

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

<a id="canonical-6ed9276c9c9c22c0f2722eeb8152725e65eac473aba14ad9ae24a2194cdbbfa5"></a>

## Next pages — rules.ingress_rules.prefix_list / ca8903d594c6 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-d1448ab984fb119569c3cfb0e3608da7525b16979fddba0b5af0f1e8a7290615"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48f794b88c6899c43cbaf6e8909f5402cca04535ef4650958a1b08ce48355988"></a>

## rules.ingress_rules.protocol_port_range — rules.ingress_rules.protocol_port_range / efcfd5195c6a / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-3c8336a0b3cef6221fd29314cb6eea9929fcb521cba3f978c1a975673479df26)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- rules.ingress_rules.protocol_port_range

<a id="canonical-71b48295d166ddef6c73127df7ace4d4712df5bdd3b40808143bf356ab638961"></a>

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

<a id="canonical-0edca4da94aa3ffe057e60a05319f9dc58b02b9c850b3296870c934b50904042"></a>

## Direct properties — rules.ingress_rules.protocol_port_range / efcfd5195c6a / 3

<a id="canonical-8486feef459ea4e76f051bb9b6d72da367d0ab310e1aed5c465fb584e19ea186"></a>

<a id="canonical-5b2e46d3f370e477d586ce4c974ba8c70c03db8a5c50c93956c9523b1b3118ed"></a>

## port_ranges property — rules.ingress_rules.protocol_port_range / efcfd5195c6a / 4

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

<a id="canonical-83741b5d200bb89f25ad828847304159f46d0a486f3a4fccce5bbc4b97f2bdc4"></a>

<a id="canonical-0025e1d827cec1a99c3c0f9d6cfb538fbb6e5dc74b28ce5f4e42ae488964208a"></a>

## protocol property — rules.ingress_rules.protocol_port_range / efcfd5195c6a / 5

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

<a id="canonical-6cd53f4abb9eaa370d2c8d9da7fd3d492dca3e9ef1d8170703900a8433ea6aa8"></a>

## Next pages — rules.ingress_rules.protocol_port_range / efcfd5195c6a / 6

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-1b2ced5562aa846dbd6a0ab570ef3a601c293f13b287016f64f318a1775a0362)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
