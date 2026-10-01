---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb0c3ad34005679a4791dd681f842d89eab1cb2797f0194783f471d3fb870fa8"></a>

## Property reference — Property reference / 3287d54a2f26 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- Property reference

<a id="canonical-b58e3d0b7ee1087ab05cf4569fe392abd475e2b0f00995f4827ed0e372c29aab"></a>

## Direct properties — Property reference / 3287d54a2f26 / 3

<a id="canonical-5c4d2c4671b3e8b7be047cc0fbab90a884988475940819fe89a7ae36f04a8527"></a>

<a id="canonical-dbf1574f2033184539e354560418b01f6f26ad73b00771fe9037fac74a44c656"></a>

## annotations property — Property reference / 3287d54a2f26 / 4

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

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63): complete subsection reference.

- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-9f99aa50006ad52bb054b5e32a6546cb75c8176bc2542f884210536a8ad6c8ae): complete subsection reference.

<a id="canonical-e03f2865c682cfcc5c6e4c1ce2e07a1719eb61f20b095154e32663a47f8a24fa"></a>

<a id="canonical-6bac56410a9511232743d48252956404b3742e5597c58efb355463bad9216c39"></a>

## description property — Property reference / 3287d54a2f26 / 5

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

<a id="canonical-7b10832f6848410e0cf053ec3b5c835c1ac384fa6eb36076e268a2ca5a718e86"></a>

<a id="canonical-cf6a31dfac9c933adfed7104e5ed860357161d73ed005baf0e9e772c3717b24d"></a>

## disable property — Property reference / 3287d54a2f26 / 6

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

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369): complete subsection reference.

<a id="canonical-091dfb0cb42273636cd2b82f7246ed381175b79169d9f95f48b3974c59144408"></a>

<a id="canonical-7445f25f575e1b8825165af1e870ee90869bc264ca23f4ad8dcc4b43ff17c1fb"></a>

## id property — Property reference / 3287d54a2f26 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4626cb0e4a457b5f49d07dd61a8f24423a27c55bc53a0654da434c2a7d001b58"></a>

<a id="canonical-4e210dbbe8ae87087c232d72323ca143f7a8ef0686228aa1d3cf6f3be4fda023"></a>

## labels property — Property reference / 3287d54a2f26 / 8

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

- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0): complete subsection reference.

<a id="canonical-27e956ca44dafbe955ab1b47f6d12fe9cefb0ba115abf9e7bbb13993fe0326d5"></a>

<a id="canonical-21a97c7be355f6b17df9369693b66e803b996f9ed15667ff7fbbd0099c0962b8"></a>

## name property — Property reference / 3287d54a2f26 / 9

Type: `"string"`. Required.

Name of the Network Interface. Must be unique within the namespace.

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

<a id="canonical-771ba06ab7838207d041da62e01b61f6473c87734e7711564d263934b453e995"></a>

<a id="canonical-c9867c9afe893a741cea025b4e9857bc60d19eada2c9fdee7d835369d06a2357"></a>

## namespace property — Property reference / 3287d54a2f26 / 10

Type: `"string"`. Required.

Namespace where the Network Interface is created.

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

- [timeouts](resources--network_interface--reference--group-002.md#canonical-737b5fe469f062275aac4bc20ccf512c05ebce12de508b1729aab672e41fe4b7): complete subsection reference.

- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0): complete subsection reference.

<a id="canonical-5b791c39d7e425a145da71dd77f8519343aad1fd5bbd5986ebd173591ca9b5a1"></a>

## All schema paths — Property reference / 3287d54a2f26 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_interface--reference--group-001.md#canonical-5c4d2c4671b3e8b7be047cc0fbab90a884988475940819fe89a7ae36f04a8527) |
| `dedicated_interface` | [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-b53b392cbc7d2577e456e9d0ae806b587b3ee541d9b5bd6809b3d5fba242d3f0) |
| `dedicated_interface.cluster` | [dedicated_interface.cluster](resources--network_interface--reference--group-001.md#canonical-0389958ec42fa822f9da01f4b72355d05e02440e4674ba11e2da4301df2e419f) |
| `dedicated_interface.device` | [dedicated_interface.device](resources--network_interface--reference--group-001.md#canonical-30fd2ac77da1e838a953920921bb7e521b4b1f1c41ee51c7885b8b06a7621a7b) |
| `dedicated_interface.is_primary` | [dedicated_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-a10081edbd5fffad5424a97d32346acd6f883f12a5373a7552d0fe28357bcb67) |
| `dedicated_interface.monitor` | [dedicated_interface.monitor](resources--network_interface--reference--group-001.md#canonical-093f038c92380a7ae5f28befce89ba99de661e277a91011d2cabf95dc928b4ff) |
| `dedicated_interface.monitor_disabled` | [dedicated_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-49c6901bc981c1b0f1b6d759ac4edfd1b91d98e4dbafc623bd9a5b4316cdf46c) |
| `dedicated_interface.mtu` | [dedicated_interface.mtu](resources--network_interface--reference--group-001.md#canonical-a3467b691dd35d7778b002e644810f87defaf6ebeca2f73316eb36f6e826fcb0) |
| `dedicated_interface.node` | [dedicated_interface.node](resources--network_interface--reference--group-001.md#canonical-3e44161c489c63bddb045102a9c6fd3ac67f300c40df038f4c6ff2c4d954a485) |
| `dedicated_interface.not_primary` | [dedicated_interface.not_primary](resources--network_interface--reference--group-001.md#canonical-49d4b82ef9978d7c94f1587aeafe7cfbadb65dcbcc9241459fe56663417793de) |
| `dedicated_interface.priority` | [dedicated_interface.priority](resources--network_interface--reference--group-001.md#canonical-f06226cfc0bcb4aac138eb3208b36c3c0bc1d458fd643f3fd361ac187a0afdcc) |
| `dedicated_management_interface` | [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-219ef4535481f8a9fc1a82fb383574a7d61bf869a42d95fd8c64720c8c323408) |
| `dedicated_management_interface.cluster` | [dedicated_management_interface.cluster](resources--network_interface--reference--group-001.md#canonical-036ae8d507a7df04b170cdb0f8b4afed9aa0e7b2695a14f9e038804546ed94ec) |
| `dedicated_management_interface.device` | [dedicated_management_interface.device](resources--network_interface--reference--group-001.md#canonical-fc16f1f99ba19cfe0b96ecffe4baeeef70736aeb15cf3d7e1bc09107cbf22bbf) |
| `dedicated_management_interface.mtu` | [dedicated_management_interface.mtu](resources--network_interface--reference--group-001.md#canonical-170bcaeccd55c9be40da899f970aed957fd689818d693e94b01c1318085ff567) |
| `dedicated_management_interface.node` | [dedicated_management_interface.node](resources--network_interface--reference--group-001.md#canonical-2edff3601b527a5318776d83bc5bc65c48e0c933ab77da3588b95cc0e8464bed) |
| `description` | [description](resources--network_interface--reference--group-001.md#canonical-e03f2865c682cfcc5c6e4c1ce2e07a1719eb61f20b095154e32663a47f8a24fa) |
| `disable` | [disable](resources--network_interface--reference--group-001.md#canonical-7b10832f6848410e0cf053ec3b5c835c1ac384fa6eb36076e268a2ca5a718e86) |
| `ethernet_interface` | [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-6e0657138526d4bae25542a24bb8e7ba776ee6142d767dc2d18163516f684e89) |
| `ethernet_interface.cluster` | [ethernet_interface.cluster](resources--network_interface--reference--group-001.md#canonical-286e873f464b7834fc4b87a68ae713ba24d44fd6951924dfbc7c846578b03b39) |
| `ethernet_interface.device` | [ethernet_interface.device](resources--network_interface--reference--group-001.md#canonical-ad8b11930cb633c6ffd0cdf7bb58f23c99963ce61ba5c48f3a8434f8dd6f28c7) |
| `ethernet_interface.dhcp_client` | [ethernet_interface.dhcp_client](resources--network_interface--reference--group-001.md#canonical-32197e25be5d5ac066858a2ded06305f48ad6d50ef275dcbcb214abeee1300d3) |
| `ethernet_interface.dhcp_server` | [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-7915f844e72da692deb4e203c026d277b9380bd10cb9a077d609420518aeebf8) |
| `ethernet_interface.dhcp_server.automatic_from_end` | [ethernet_interface.dhcp_server.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-a1389086e0f253cfd6611c869eec2895de0b8e241c7155aebb064ee1862857a2) |
| `ethernet_interface.dhcp_server.automatic_from_start` | [ethernet_interface.dhcp_server.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-3963b691ce8ab57c4284192187d169f72379e1ad578495f58842b697fc247767) |
| `ethernet_interface.dhcp_server.dhcp_networks` | [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-dfc27e6ec4562dc5c33208f7b5a4eefd0e8ac17618e7269799f1489d4dc07af8) |
| `ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [ethernet_interface.dhcp_server.dhcp_networks.dgw_address](resources--network_interface--reference--group-001.md#canonical-0c12d34ad083221af6f88734f0da77d5ad76b4a9acb1a9f149565f979692e6cf) |
| `ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [ethernet_interface.dhcp_server.dhcp_networks.dns_address](resources--network_interface--reference--group-001.md#canonical-b6f82ce1ccfe203f236d0c6f5ee142af2b2c74fc147d623a111a2af22281eccc) |
| `ethernet_interface.dhcp_server.dhcp_networks.first_address` | [ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--network_interface--reference--group-001.md#canonical-f51d24bc743daf61d8f3f80e215bd34f82a9b854c133dc9d542c73c8bcbed4d9) |
| `ethernet_interface.dhcp_server.dhcp_networks.last_address` | [ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--network_interface--reference--group-001.md#canonical-9ef6a54735209c3c85aa73dff81f03724c3578772da3ae050140820c1a951bbb) |
| `ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [ethernet_interface.dhcp_server.dhcp_networks.network_prefix](resources--network_interface--reference--group-001.md#canonical-ea593d83ce801f035965cd79da01d8034307d098860caa27d8b1a076f4befe0f) |
| `ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [ethernet_interface.dhcp_server.dhcp_networks.pool_settings](resources--network_interface--reference--group-001.md#canonical-4fe272c0cbf06036276f20145864ff2f6066cabbefe8104e6e7a76ac474cf8fa) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools` | [ethernet_interface.dhcp_server.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-ab8a2f63fcaed4da5cf579820aa5ad3f5d1f4e4210372bb7667cc321d30640c5) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](resources--network_interface--reference--group-001.md#canonical-32e578cc24b0dd97de5581af53349ed49bd3c288d1762fe326669894fecf3bc4) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](resources--network_interface--reference--group-001.md#canonical-8afa8bbc0805b591a13da147260f95225e7295abab9b68185c621d17ed625658) |
| `ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](resources--network_interface--reference--group-001.md#canonical-46b2636cf331774fc5f4e0fa97a37085c6672a03350cc1601abb267be4016e9b) |
| `ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--network_interface--reference--group-001.md#canonical-aad0e1b6cc4c37da1332dfc9f63b0ff0539a92bd2f643972f39ba463764b49d8) |
| `ethernet_interface.dhcp_server.dhcp_option82_tag` | [ethernet_interface.dhcp_server.dhcp_option82_tag](resources--network_interface--reference--group-001.md#canonical-0e561ae991d832d585d2f2968e78188cfe73f8e006d599e52b292c6aca883b02) |
| `ethernet_interface.dhcp_server.fixed_ip_map` | [ethernet_interface.dhcp_server.fixed_ip_map](resources--network_interface--reference--group-001.md#canonical-6c0819c7d69877e3656f3602f7e66415fd51264f13a2a983122eb1bd49938102) |
| `ethernet_interface.dhcp_server.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-aa45145b5108abf11edf5528e142a3621978db41eb9c9613f7c3025e4b53a68b) |
| `ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-fd0ba71fd7e58bd5b96bae80fb0c04dcd114af16d8ddfd74df5421a15b81d417) |
| `ethernet_interface.ipv6_auto_config` | [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-0383e001aa6430ea05121ee741becd34ed34615bfc522fddb69fbc000d8f2bfc) |
| `ethernet_interface.ipv6_auto_config.host` | [ethernet_interface.ipv6_auto_config.host](resources--network_interface--reference--group-001.md#canonical-2d63b82cc646675e9b627784839a600b6f1e1c540593a29e908f9ed93883ed66) |
| `ethernet_interface.ipv6_auto_config.router` | [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-589d774633953ed07ee758ce5a1024ef59608b53cc9c84bd932098b6110a4dcc) |
| `ethernet_interface.ipv6_auto_config.router.dns_config` | [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-dc53d7e926c8dfd350196bd1ea64c7e835eadc27b224663dede4a9b56616aa79) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--network_interface--reference--group-001.md#canonical-c69670bff4b518bd2539b057799f58abf4ef8b2c8f43dba306676c0b23addc64) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](resources--network_interface--reference--group-001.md#canonical-0fc0b63ee232663ab13e8716b21e785e67284fb2416378f6a66e8ee31a798b47) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-9f98eeb5e176262b2c14a05c871c5c0e481659097b92fb2fd93915fb71c4c136) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](resources--network_interface--reference--group-001.md#canonical-2138a91165a0d6f98a01feef655973e5339049d6a8d5af6a9294be30c873b57b) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--network_interface--reference--group-001.md#canonical-0286f35c933e2f6c0be807ac2975447e36242ba6e3d9016db0ec22a9ff12b30a) |
| `ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--network_interface--reference--group-001.md#canonical-94bc9ace9b2f5e8776d70138d3bea753e9c76e0cf8d30ed8bf9660b6650ddec1) |
| `ethernet_interface.ipv6_auto_config.router.network_prefix` | [ethernet_interface.ipv6_auto_config.router.network_prefix](resources--network_interface--reference--group-001.md#canonical-16936745aca8643384bc1d4579a9493fa57ca2444d0f1938f1ce6fda8b73b810) |
| `ethernet_interface.ipv6_auto_config.router.stateful` | [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-17c38b9f20cdd06f05cfd07ce97855102c6b90ece5bd3c78800f8fec5b534a5a) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-6ca6f12adedaa622904057d41fcf183bed3b220f75f102e27814681a062b1553) |
| `ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-8bf4df97b8e24fd925fd58d569d20a9595e55609b938e8381c67f6ba415b9e50) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-8b0f37b60c5b8865deb3d15af9be8753aa487a374d79337a4aa75e8db5d607fe) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](resources--network_interface--reference--group-001.md#canonical-22175ff499b48401ed21d2d8c84db14ce214635c4631614ba66225003e15c1d9) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](resources--network_interface--reference--group-001.md#canonical-0a747ef544cb43b079530106d8806dfa4b18c99b9e8b488f13feec03bb6ffdc9) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-07227e352a23c33038c8374c3baa62a5eb541ea000cf80e028b03e600895a78f) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](resources--network_interface--reference--group-001.md#canonical-6f011446dce544ad6c7dec656d0c76eca541e857317f21d93d9c4db6c6d02127) |
| `ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](resources--network_interface--reference--group-001.md#canonical-7120feb55e299061997c40d2a160d0f5865804196dcce1ee4b82292de692e4db) |
| `ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](resources--network_interface--reference--group-001.md#canonical-0d4dc6af4c18220c76c2d64fbb35850bdc4387b6f40a2f2a80012029730ce73c) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-8d16e5d6f2f3ce8eae7c8caa200948c6ffa86884f6c53800cf07a8788cd77341) |
| `ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-3e9e67af64c477a87d90da26f09bc231a0599be249bb2118037a5965e9acbb06) |
| `ethernet_interface.is_primary` | [ethernet_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-2a25c8b30670c082b6ccd885febb75c927e6b9dd5b45566ac51d857d95adcbaa) |
| `ethernet_interface.monitor` | [ethernet_interface.monitor](resources--network_interface--reference--group-001.md#canonical-2c2ea94141d5f5c631761897b48784fefa665733c6379d00a7a4e5b646f69c9f) |
| `ethernet_interface.monitor_disabled` | [ethernet_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-00b1b710483d9029260b5e5c65c45e4d6cdcb127f9d18922f127c6796221fe7b) |
| `ethernet_interface.mtu` | [ethernet_interface.mtu](resources--network_interface--reference--group-001.md#canonical-13fe58f9ead768ef7c723e871124f50733863752b281f8091f2ea7ee5ed9c362) |
| `ethernet_interface.no_ipv6_address` | [ethernet_interface.no_ipv6_address](resources--network_interface--reference--group-001.md#canonical-81e8bde0c542ae046890ccd7b6ca97bc9f4c67ec4f003094bbfe2f85610d581b) |
| `ethernet_interface.node` | [ethernet_interface.node](resources--network_interface--reference--group-001.md#canonical-0c231be6cbd72c2b7c982492443ea07e54ea77ad45477ebf4d620759e14a0afb) |
| `ethernet_interface.not_primary` | [ethernet_interface.not_primary](resources--network_interface--reference--group-001.md#canonical-d6feb1ad4a0f6da80f15d4780c38a93c354d8bd742d75c6b8b7bddce772ecfe5) |
| `ethernet_interface.priority` | [ethernet_interface.priority](resources--network_interface--reference--group-001.md#canonical-b4a227d8d1a6c6cc2cd434e4a3f253e81ab115a930ba33ab47a0626182c9de8b) |
| `ethernet_interface.site_local_inside_network` | [ethernet_interface.site_local_inside_network](resources--network_interface--reference--group-001.md#canonical-85f6dd6599ba156de5c47aa2aa4ec8844c6367bb5435184c66788769101f2b66) |
| `ethernet_interface.site_local_network` | [ethernet_interface.site_local_network](resources--network_interface--reference--group-001.md#canonical-83918a5d5cc033b3d4ed328d9c15172284ab234e6456012e467d4851c2d58a34) |
| `ethernet_interface.static_ip` | [ethernet_interface.static_ip](resources--network_interface--reference--group-001.md#canonical-bb93a5143905ed4c3fba99beb8e9a3731ed985a6e0f8a8a48f019d236e243290) |
| `ethernet_interface.static_ip.cluster_static_ip` | [ethernet_interface.static_ip.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-6a939f7a7a5d9fb7934c25d39a50aa46819ab8593b205ac02d66b183b4b61c95) |
| `ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](resources--network_interface--reference--group-002.md#canonical-5737da7b60ec4c327c8026de727f9b79d8d47a924d95eaf2941b4d26179bbdc3) |
| `ethernet_interface.static_ip.node_static_ip` | [ethernet_interface.static_ip.node_static_ip](resources--network_interface--reference--group-002.md#canonical-d70deaa7005b2adac45dd4b06964f96ce201636b4b09ecc1fd8b3a813393f78d) |
| `ethernet_interface.static_ip.node_static_ip.default_gw` | [ethernet_interface.static_ip.node_static_ip.default_gw](resources--network_interface--reference--group-002.md#canonical-2dbaf847402ec4cfc27eef77c4bcc1416f495d74d428cea25f2dc3ee25f0b3ea) |
| `ethernet_interface.static_ip.node_static_ip.dns_server` | [ethernet_interface.static_ip.node_static_ip.dns_server](resources--network_interface--reference--group-002.md#canonical-4f883f9bf03d433e5d41c98daf15eeaa39eb4e15b64725621a048a49728f4b5a) |
| `ethernet_interface.static_ip.node_static_ip.ip_address` | [ethernet_interface.static_ip.node_static_ip.ip_address](resources--network_interface--reference--group-002.md#canonical-7bdce07a4740d26871beccccd5da616e98e84ad4e0e4b1370110e3cd7a63a568) |
| `ethernet_interface.static_ipv6_address` | [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-301fb268d2e32610b857aa8c1512cd5c6b32b1a2a69fdc553ddede6695b27c37) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip` | [ethernet_interface.static_ipv6_address.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-d153016fcdd189f2161cec708be092a601d3d42d2364deec057154f6c45cd7fe) |
| `ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](resources--network_interface--reference--group-002.md#canonical-7061c448dfa6b665754045c4d2b9d339578b0bcbb9066131b9544e2ccd82b066) |
| `ethernet_interface.static_ipv6_address.node_static_ip` | [ethernet_interface.static_ipv6_address.node_static_ip](resources--network_interface--reference--group-002.md#canonical-d20402919e10de0d950f5ea8f899f5d67291e20eb8688d64947f65d8bbe5c7b2) |
| `ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [ethernet_interface.static_ipv6_address.node_static_ip.default_gw](resources--network_interface--reference--group-002.md#canonical-152f91ddc7eae9822fd45692fea0bf829269c2ca52cab80cd2fc5ef2d2c1be1a) |
| `ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [ethernet_interface.static_ipv6_address.node_static_ip.dns_server](resources--network_interface--reference--group-002.md#canonical-30d74b0aa093f79afeccdf7803b6088108362f5c334fcfdfe057a0e44ecae714) |
| `ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [ethernet_interface.static_ipv6_address.node_static_ip.ip_address](resources--network_interface--reference--group-002.md#canonical-7971030a9ac520e6eedecc19034aa7caeac217fb34a13138901deda62dc7580f) |
| `ethernet_interface.storage_network` | [ethernet_interface.storage_network](resources--network_interface--reference--group-002.md#canonical-7ff07e6be90117e3e8bb1aad55041c91828b69ab7b035addabeada48182ad13a) |
| `ethernet_interface.untagged` | [ethernet_interface.untagged](resources--network_interface--reference--group-002.md#canonical-a87a5701bb14fcfda551d90bedf6696e6fa4b1ab0aa87b530b53689cba07b030) |
| `ethernet_interface.vlan_id` | [ethernet_interface.vlan_id](resources--network_interface--reference--group-001.md#canonical-ccc789142af389b5195663f33381a0cbcde4f82c0572a9a8748da98be57e9e9a) |
| `id` | [id](resources--network_interface--reference--group-001.md#canonical-091dfb0cb42273636cd2b82f7246ed381175b79169d9f95f48b3974c59144408) |
| `labels` | [labels](resources--network_interface--reference--group-001.md#canonical-4626cb0e4a457b5f49d07dd61a8f24423a27c55bc53a0654da434c2a7d001b58) |
| `layer2_interface` | [layer2_interface](resources--network_interface--reference--group-002.md#canonical-e83fe3d490b7b1b9797f3df5de4ff22bc8deb68ea422ed8378cb6ac6f9fb18b9) |
| `layer2_interface.l2sriov_interface` | [layer2_interface.l2sriov_interface](resources--network_interface--reference--group-002.md#canonical-2ae9ea6c87a7a85bbc990c5b62b10690956df01954b7faad8f67f41745e24aa3) |
| `layer2_interface.l2sriov_interface.device` | [layer2_interface.l2sriov_interface.device](resources--network_interface--reference--group-002.md#canonical-404c1e6ba76e08387c9d77df9342e24d92f6172d9a00cc70bb43cff11a480ef0) |
| `layer2_interface.l2sriov_interface.untagged` | [layer2_interface.l2sriov_interface.untagged](resources--network_interface--reference--group-002.md#canonical-2b6240834c2b7be8af41328df0235100df6f3a933aa9f74b97b0e4d48c42f46d) |
| `layer2_interface.l2sriov_interface.vlan_id` | [layer2_interface.l2sriov_interface.vlan_id](resources--network_interface--reference--group-002.md#canonical-d00c578d39e778cfd60387e5bdf3f496eb83221cc045beb9f44b51e12630bbac) |
| `layer2_interface.l2vlan_interface` | [layer2_interface.l2vlan_interface](resources--network_interface--reference--group-002.md#canonical-6b68b78583b4ab1bbf657e78184145acc9af43f231669ec56a2789d6d59dc06a) |
| `layer2_interface.l2vlan_interface.device` | [layer2_interface.l2vlan_interface.device](resources--network_interface--reference--group-002.md#canonical-af303c0f4aeecc60838cc6dc0a54f8558208a0fb7b29e336b4ee6cac8fec1de0) |
| `layer2_interface.l2vlan_interface.vlan_id` | [layer2_interface.l2vlan_interface.vlan_id](resources--network_interface--reference--group-002.md#canonical-0bf9ab6fdb2ca90b571c1a8b08fafbbbba6a6578aa527dadc77c04e01f18cc85) |
| `layer2_interface.l2vlan_slo_interface` | [layer2_interface.l2vlan_slo_interface](resources--network_interface--reference--group-002.md#canonical-1e603ca6b506635b704b91fb12737a2a1b5b373ed71bd78f86f38dcad234de0e) |
| `layer2_interface.l2vlan_slo_interface.vlan_id` | [layer2_interface.l2vlan_slo_interface.vlan_id](resources--network_interface--reference--group-002.md#canonical-67387467502d314677eb485651447c1910e7336ad39b4fd5fe202be5cac28714) |
| `name` | [name](resources--network_interface--reference--group-001.md#canonical-27e956ca44dafbe955ab1b47f6d12fe9cefb0ba115abf9e7bbb13993fe0326d5) |
| `namespace` | [namespace](resources--network_interface--reference--group-001.md#canonical-771ba06ab7838207d041da62e01b61f6473c87734e7711564d263934b453e995) |
| `timeouts` | [timeouts](resources--network_interface--reference--group-002.md#canonical-0f1a477aeb1c77ebf2be4f90b99bfacf1b54171a03d304b0c69a4f6c23e114e7) |
| `timeouts.create` | [timeouts.create](resources--network_interface--reference--group-002.md#canonical-efbe76065569c878d92de43cd3da30813ab33955833b53dd17984ae6a1481d8a) |
| `timeouts.delete` | [timeouts.delete](resources--network_interface--reference--group-002.md#canonical-68f60b2a82d750e732d9236e23839690ba7b53dc64a401ab97f41808c6f2f8e2) |
| `timeouts.read` | [timeouts.read](resources--network_interface--reference--group-002.md#canonical-1820c17b3dc561b882b296756e309f8b4c84ce3636eb55d3da302df03cc55dca) |
| `timeouts.update` | [timeouts.update](resources--network_interface--reference--group-002.md#canonical-c1a896e6be4dfe1ecbdfcbfd9b32da7905b4ce1feec017af921e88d22f226a67) |
| `tunnel_interface` | [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-5b16fd24e80b66597695a9d58656efff23e866192b280fee851b04a157e01f46) |
| `tunnel_interface.mtu` | [tunnel_interface.mtu](resources--network_interface--reference--group-002.md#canonical-00466f20e230b514fb5e765632a201547ae0631eecd690ce090f58a44e5fec93) |
| `tunnel_interface.node` | [tunnel_interface.node](resources--network_interface--reference--group-002.md#canonical-b97e10cdb6f73a19ae05dd0dd9baacc1a9e6c7ffd0d4d3520371a9b4da0d4699) |
| `tunnel_interface.priority` | [tunnel_interface.priority](resources--network_interface--reference--group-002.md#canonical-8004221dfacd326c6a35ce6bbc434636bcfa2feee72258cedf9a70acd3557888) |
| `tunnel_interface.site_local_inside_network` | [tunnel_interface.site_local_inside_network](resources--network_interface--reference--group-002.md#canonical-c3a8d12380b8f911d1c554b5e7d05168e525a9aa92cd50c8cd19ab9f8fee4152) |
| `tunnel_interface.site_local_network` | [tunnel_interface.site_local_network](resources--network_interface--reference--group-002.md#canonical-7956099f3a05592b4aa819ed490ff9f8c57dd6c146544af693a9bde8fef480eb) |
| `tunnel_interface.static_ip` | [tunnel_interface.static_ip](resources--network_interface--reference--group-002.md#canonical-c66df417706b2af56978e29f09c01f35fe11f353aaceca85d8f98614b9a23b14) |
| `tunnel_interface.static_ip.cluster_static_ip` | [tunnel_interface.static_ip.cluster_static_ip](resources--network_interface--reference--group-002.md#canonical-d32d6df37cd00022a747281774056c3c2448dc85d3dd169936f9edccbd0dc0a8) |
| `tunnel_interface.static_ip.cluster_static_ip.interface_ip_map` | [tunnel_interface.static_ip.cluster_static_ip.interface_ip_map](resources--network_interface--reference--group-002.md#canonical-c9589b25857b35d35ec1255197c2a76caea57365ba94a268e8f6c31199fbb4e4) |
| `tunnel_interface.static_ip.node_static_ip` | [tunnel_interface.static_ip.node_static_ip](resources--network_interface--reference--group-002.md#canonical-c7f4a5adbb982884dc3c87beb9e6027f8b8b1b2e2fcf5fd12dc0f2ecf74d3a1a) |
| `tunnel_interface.static_ip.node_static_ip.default_gw` | [tunnel_interface.static_ip.node_static_ip.default_gw](resources--network_interface--reference--group-002.md#canonical-593f543f94c55632fb933a6d7fb6819b32a92404ec944c944aede4065bbd5237) |
| `tunnel_interface.static_ip.node_static_ip.dns_server` | [tunnel_interface.static_ip.node_static_ip.dns_server](resources--network_interface--reference--group-002.md#canonical-957a383307f14b2b9837d2629ef9d09b9aa2e247b49dae28fe494f549522dfe3) |
| `tunnel_interface.static_ip.node_static_ip.ip_address` | [tunnel_interface.static_ip.node_static_ip.ip_address](resources--network_interface--reference--group-002.md#canonical-7b3bacbd7966d385fa96314f5958111887a3abebe8eef40d1eae6d56b59f43f2) |
| `tunnel_interface.tunnel` | [tunnel_interface.tunnel](resources--network_interface--reference--group-002.md#canonical-e69e395b805d387e01859c146fbd0e0d20250fbc1c2a2ad34babd72602f38120) |
| `tunnel_interface.tunnel.name` | [tunnel_interface.tunnel.name](resources--network_interface--reference--group-002.md#canonical-e61f2ad9ebabdbf83f94b1596141a831b9e5cecbf4b5c938a01bfb91baaadc73) |
| `tunnel_interface.tunnel.namespace` | [tunnel_interface.tunnel.namespace](resources--network_interface--reference--group-002.md#canonical-60edc75cd60824678ec11b0ade93b552c21c20b6a2fd8c30c79fd4399a4776b2) |
| `tunnel_interface.tunnel.tenant` | [tunnel_interface.tunnel.tenant](resources--network_interface--reference--group-002.md#canonical-79f3b8880bcb9b273cacd00d3bff242f1a1b475c89a996bf494a1b77b9192551) |

<a id="canonical-674f6cae6efecfa28b5126d6a0b4906d67bd1353c5e3f13fa90dfeff057075be"></a>

## Next pages — Property reference / 3287d54a2f26 / 12

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-9f99aa50006ad52bb054b5e32a6546cb75c8176bc2542f884210536a8ad6c8ae)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-1c594b57dc38583cede4438254bed4c980392fdf365967edc1ad24fac341cbd0)
- [timeouts](resources--network_interface--reference--group-002.md#canonical-737b5fe469f062275aac4bc20ccf512c05ebce12de508b1729aab672e41fe4b7)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-0db6b411298b0a3c07d71c63f1bb4732bc353ecea9c44af1f2a38335e20aeda0)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e77b7c10b84341e6107b06875dc2209fe4ba25773097519bccff827497dc441"></a>

## dedicated_interface — dedicated_interface / d159a1b096fd / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- dedicated_interface

<a id="canonical-b53b392cbc7d2577e456e9d0ae806b587b3ee541d9b5bd6809b3d5fba242d3f0"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dedicated\_interface, dedicated\_management\_interface, ethernet\_interface,
layer2\_interface, tunnel\_interface\] Configuration parameter for dedicated interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]"
}
```

OneOf alternatives in this subsection:

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-b53b392cbc7d2577e456e9d0ae806b587b3ee541d9b5bd6809b3d5fba242d3f0)
- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-219ef4535481f8a9fc1a82fb383574a7d61bf869a42d95fd8c64720c8c323408)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-6e0657138526d4bae25542a24bb8e7ba776ee6142d767dc2d18163516f684e89)
- [layer2_interface](resources--network_interface--reference--group-002.md#canonical-e83fe3d490b7b1b9797f3df5de4ff22bc8deb68ea422ed8378cb6ac6f9fb18b9)
- [tunnel_interface](resources--network_interface--reference--group-002.md#canonical-5b16fd24e80b66597695a9d58656efff23e866192b280fee851b04a157e01f46)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dedicated_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2e6a1c076ba2927fd9e24be5b5829e18d73d13d0755ecd034c39fc133c7c2244"></a>

## Direct properties — dedicated_interface / d159a1b096fd / 3

- [cluster](resources--network_interface--reference--group-001.md#canonical-27422342c3f165fc8ac99fdeefd53108887eb6c5b3663331e4f3d411232ad547): complete subsection reference.

<a id="canonical-30fd2ac77da1e838a953920921bb7e521b4b1f1c41ee51c7885b8b06a7621a7b"></a>

<a id="canonical-477b36b2bff727240acc0292a12d56a6ad4edc5203e9eea96c8abe1fe8fec526"></a>

## device property — dedicated_interface / d159a1b096fd / 4

Type: `"string"`. Optional.

Name of the device for which interface is configured. Use wwan0 for 4G/LTE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [is_primary](resources--network_interface--reference--group-001.md#canonical-58a59d3ca74f080847a66d5f91742f825b500d17713a0aa95cdd6cb441456b9f): complete subsection reference.

- [monitor](resources--network_interface--reference--group-001.md#canonical-ece69ee6c836de64b9981232fee7e8bc22dcc44ed73034a5716dbcff4d8c2456): complete subsection reference.

- [monitor_disabled](resources--network_interface--reference--group-001.md#canonical-fc09dbc340f4614e82c9cc6bf163241fb10c2fdf3ba84df2c5ee791fb15ba337): complete subsection reference.

<a id="canonical-a3467b691dd35d7778b002e644810f87defaf6ebeca2f73316eb36f6e826fcb0"></a>

<a id="canonical-db799282fc747efe498bf5fbdf937f4ced2f950194c42c1d0c4154c5715d7643"></a>

## mtu property — dedicated_interface / d159a1b096fd / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-3e44161c489c63bddb045102a9c6fd3ac67f300c40df038f4c6ff2c4d954a485"></a>

<a id="canonical-d793f7451e79984e19e949e339e897739cd99910409f616c7734171bdf143994"></a>

## node property — dedicated_interface / d159a1b096fd / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--network_interface--reference--group-001.md#canonical-7cbbbecfa7b68b984bcc1dc5ca2d01b7ff96cc17f12d1a21b4c2844e56837835): complete subsection reference.

<a id="canonical-f06226cfc0bcb4aac138eb3208b36c3c0bc1d458fd643f3fd361ac187a0afdcc"></a>

<a id="canonical-a0a911670b8c6ae8ad118cfd8ab0ad5b8b3163a4f1f3830a7a131dab184a2ab4"></a>

## priority property — dedicated_interface / d159a1b096fd / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-74b433ff94dbce14fd9dcdc06e75cf06a86a0d2d812ba6635d4e3d17437d6329"></a>

## Next pages — dedicated_interface / d159a1b096fd / 8

- [dedicated_interface.cluster](resources--network_interface--reference--group-001.md#canonical-27422342c3f165fc8ac99fdeefd53108887eb6c5b3663331e4f3d411232ad547)
- [dedicated_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-58a59d3ca74f080847a66d5f91742f825b500d17713a0aa95cdd6cb441456b9f)
- [dedicated_interface.monitor](resources--network_interface--reference--group-001.md#canonical-ece69ee6c836de64b9981232fee7e8bc22dcc44ed73034a5716dbcff4d8c2456)
- [dedicated_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-fc09dbc340f4614e82c9cc6bf163241fb10c2fdf3ba84df2c5ee791fb15ba337)
- [dedicated_interface.not_primary](resources--network_interface--reference--group-001.md#canonical-7cbbbecfa7b68b984bcc1dc5ca2d01b7ff96cc17f12d1a21b4c2844e56837835)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-27422342c3f165fc8ac99fdeefd53108887eb6c5b3663331e4f3d411232ad547"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36cfe1672fa82a161197a094d684de77f19c662b6807cc1a59d6fc33e2cf1b0f"></a>

## dedicated_interface.cluster — dedicated_interface.cluster / 9485621987d8 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- dedicated_interface.cluster

<a id="canonical-0389958ec42fa822f9da01f4b72355d05e02440e4674ba11e2da4301df2e419f"></a>

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
cluster = {}
```

<a id="canonical-d8fe73a27720dbb16c5cd43d3dc8b997bb2940cc5d5fe875df4f4a82c18ae2de"></a>

## Direct properties — dedicated_interface.cluster / 9485621987d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b9ed6cd0380b20e82b2f3e98165dca7077f352178a463778f71dd8666c4871f"></a>

## Next pages — dedicated_interface.cluster / 9485621987d8 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-58a59d3ca74f080847a66d5f91742f825b500d17713a0aa95cdd6cb441456b9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ab1556aa76c2be36b19e0777d16180cdb0a3ca8874d753b78effd2dc703dcae"></a>

## dedicated_interface.is_primary — dedicated_interface.is_primary / 9343a2c8b5cc / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- dedicated_interface.is_primary

<a id="canonical-a10081edbd5fffad5424a97d32346acd6f883f12a5373a7552d0fe28357bcb67"></a>

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
is_primary = {}
```

<a id="canonical-aebea271dd98910c22f55ae845054fc27e78a615e39275db69889f7d5cf070db"></a>

## Direct properties — dedicated_interface.is_primary / 9343a2c8b5cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-936ade05d88a70f4282c3fe34b09fc02e3556f3645df72fe5d0d154d208ccbbc"></a>

## Next pages — dedicated_interface.is_primary / 9343a2c8b5cc / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-ece69ee6c836de64b9981232fee7e8bc22dcc44ed73034a5716dbcff4d8c2456"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d502c3c0c2352a03e16da01dae747f48d1742cbabb43124a7823a4e8e64acb54"></a>

## dedicated_interface.monitor — dedicated_interface.monitor / f0633f122996 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- dedicated_interface.monitor

<a id="canonical-093f038c92380a7ae5f28befce89ba99de661e277a91011d2cabf95dc928b4ff"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

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
monitor = {}
```

<a id="canonical-f99361a3b640bede4ef0e505c934a2ee19f54afe2b97e3cef361646fdf0593d5"></a>

## Direct properties — dedicated_interface.monitor / f0633f122996 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a471328cfc49b891bc637d2a4aeb949e98676d53fbac9ca4c6559610ef8e51c7"></a>

## Next pages — dedicated_interface.monitor / f0633f122996 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-fc09dbc340f4614e82c9cc6bf163241fb10c2fdf3ba84df2c5ee791fb15ba337"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2baf89994523bdc6219db9f6eebb4862b560fce9a58ff12a40118d010a65f92f"></a>

## dedicated_interface.monitor_disabled — dedicated_interface.monitor_disabled / 289dd24058b8 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- dedicated_interface.monitor_disabled

<a id="canonical-49c6901bc981c1b0f1b6d759ac4edfd1b91d98e4dbafc623bd9a5b4316cdf46c"></a>

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
monitor_disabled = {}
```

<a id="canonical-378525457b040c0af79b514a94196498294e4dece3c885dc98feb266a364c5d3"></a>

## Direct properties — dedicated_interface.monitor_disabled / 289dd24058b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2df9f189be5c74f11193b2326a9bd749f01bdad17feb397d3a8eab25aefc9b5"></a>

## Next pages — dedicated_interface.monitor_disabled / 289dd24058b8 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-7cbbbecfa7b68b984bcc1dc5ca2d01b7ff96cc17f12d1a21b4c2844e56837835"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-552b18958a08cad9a823b26af9b162d0ff4685a73618130db2c4a93d7d4efed9"></a>

## dedicated_interface.not_primary — dedicated_interface.not_primary / e6496bba3cc8 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- dedicated_interface.not_primary

<a id="canonical-49d4b82ef9978d7c94f1587aeafe7cfbadb65dcbcc9241459fe56663417793de"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for not primary.

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
not_primary = {}
```

<a id="canonical-e39e2408781d64c996e098213b6ce8ba5951480589426b2b496983df83320397"></a>

## Direct properties — dedicated_interface.not_primary / e6496bba3cc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce4307f65d91b526bff681dd14fd93958d0c724d388d0394d051f0e1aa9677ad"></a>

## Next pages — dedicated_interface.not_primary / e6496bba3cc8 / 4

- [dedicated_interface](resources--network_interface--reference--group-001.md#canonical-a2d9ea4197f661139cf8470e947f4539eddfea1c7053324d7faf90cd97f12b63)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-9f99aa50006ad52bb054b5e32a6546cb75c8176bc2542f884210536a8ad6c8ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e28357377f4ca9baf74013c04b933257c98639eeff62c169b5c3c570ef33fc7"></a>

## dedicated_management_interface — dedicated_management_interface / cc8e986e7761 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- dedicated_management_interface

<a id="canonical-219ef4535481f8a9fc1a82fb383574a7d61bf869a42d95fd8c64720c8c323408"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]"
}
```

Terraform syntax:

```terraform
dedicated_management_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-07f9ad6b409a0799d533ecaf7f86c406a42fe02ecece76f5ebb4313cd976105a"></a>

## Direct properties — dedicated_management_interface / cc8e986e7761 / 3

- [cluster](resources--network_interface--reference--group-001.md#canonical-e663eafd5a4fbdd040e7770cf348a08812d71afecdca8047b8773e20a316e0d8): complete subsection reference.

<a id="canonical-fc16f1f99ba19cfe0b96ecffe4baeeef70736aeb15cf3d7e1bc09107cbf22bbf"></a>

<a id="canonical-ac86a32c53cd87e9ea8531d5d51c47bc5802d4e5d38afb9eb4282b4dbf803c38"></a>

## device property — dedicated_management_interface / cc8e986e7761 / 4

Type: `"string"`. Optional.

Name of the device for which interface is configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-170bcaeccd55c9be40da899f970aed957fd689818d693e94b01c1318085ff567"></a>

<a id="canonical-2ca3075f04d60fdb4f79700ac59d1ae9fbc7028cfc6772245588c8dc3b2f1484"></a>

## mtu property — dedicated_management_interface / cc8e986e7761 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-2edff3601b527a5318776d83bc5bc65c48e0c933ab77da3588b95cc0e8464bed"></a>

<a id="canonical-5f5dcfad22d611d6163d3183f0563f3a483d87152bf1a05006ee42b12d8a2b0a"></a>

## node property — dedicated_management_interface / cc8e986e7761 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-ce55f8f7a7a55bff1e3c877d0ea6d44fdc45af836bde958e4bb08414b70d1e09"></a>

## Next pages — dedicated_management_interface / cc8e986e7761 / 7

- [dedicated_management_interface.cluster](resources--network_interface--reference--group-001.md#canonical-e663eafd5a4fbdd040e7770cf348a08812d71afecdca8047b8773e20a316e0d8)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-e663eafd5a4fbdd040e7770cf348a08812d71afecdca8047b8773e20a316e0d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf7e964e7a03145da3259b3aac2ce528be22f629ad5e1571d5fb2422c9ee99e2"></a>

## dedicated_management_interface.cluster — dedicated_management_interface.cluster / 5f612c79b1d2 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-9f99aa50006ad52bb054b5e32a6546cb75c8176bc2542f884210536a8ad6c8ae)
- dedicated_management_interface.cluster

<a id="canonical-036ae8d507a7df04b170cdb0f8b4afed9aa0e7b2695a14f9e038804546ed94ec"></a>

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
cluster = {}
```

<a id="canonical-3329a90a0e6f52463e095fcf508072d3bba84ac080f8d82eea68dcaa2cebd936"></a>

## Direct properties — dedicated_management_interface.cluster / 5f612c79b1d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-116960a5c6eecfb3b8da04f1ff250fe00a00f2453123615622b8e8128b55a565"></a>

## Next pages — dedicated_management_interface.cluster / 5f612c79b1d2 / 4

- [dedicated_management_interface](resources--network_interface--reference--group-001.md#canonical-9f99aa50006ad52bb054b5e32a6546cb75c8176bc2542f884210536a8ad6c8ae)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77926aa6865fd0aa0e56e975abf2667424a4261b64c6cc8f14506b7dce4fc113"></a>

## ethernet_interface — ethernet_interface / 0483b89fedbc / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- ethernet_interface

<a id="canonical-6e0657138526d4bae25542a24bb8e7ba776ee6142d767dc2d18163516f684e89"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Upstream description:

Ethernet Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("site_local_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("untagged",
    "vlan_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-54c35f6e1c4b7fafd9d965ddef89361ff87be1b4a7d538c8b76f888e57b14da4"></a>

## Direct properties — ethernet_interface / 0483b89fedbc / 3

- [cluster](resources--network_interface--reference--group-001.md#canonical-3d6e609b89c36cad162609dc7fef0ad20e646a9c2b1f76297e3cf0cc40a4c481): complete subsection reference.

<a id="canonical-ad8b11930cb633c6ffd0cdf7bb58f23c99963ce61ba5c48f3a8434f8dd6f28c7"></a>

<a id="canonical-748aa8a63fb19f9109ea315d88fcce354a6e6ac104e6a13548d47e1bf0fc672c"></a>

## device property — ethernet_interface / 0483b89fedbc / 4

Type: `"string"`. Optional.

Interface configuration for the ethernet device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [dhcp_client](resources--network_interface--reference--group-001.md#canonical-9cd6c7a9dfa1e042007ad473dc91abad1b1892c94ecc157b826ea35e4d5fde01): complete subsection reference.

- [dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a): complete subsection reference.

- [ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc): complete subsection reference.

- [is_primary](resources--network_interface--reference--group-001.md#canonical-96d9f13d322d043ffab7b9d4f26f2db3f8534caa14beae66c28f073b0a095cfe): complete subsection reference.

- [monitor](resources--network_interface--reference--group-001.md#canonical-62801a6fa2931fd2b7d35c8817f8b05c3d35b24eb7bb2de84c2cd554935d3567): complete subsection reference.

- [monitor_disabled](resources--network_interface--reference--group-001.md#canonical-46034b21cdd271dec1c6260712d1736070b3232acf27399d469376dad34aa497): complete subsection reference.

<a id="canonical-13fe58f9ead768ef7c723e871124f50733863752b281f8091f2ea7ee5ed9c362"></a>

<a id="canonical-f9f7015430697df32ab0d647a74211251347c5889a6a1cc93db936e865bd1664"></a>

## mtu property — ethernet_interface / 0483b89fedbc / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](resources--network_interface--reference--group-001.md#canonical-7f39517c0d7d7278f51272833c3b4b872306085209e35390e8cae77da081ef67): complete subsection reference.

<a id="canonical-0c231be6cbd72c2b7c982492443ea07e54ea77ad45477ebf4d620759e14a0afb"></a>

<a id="canonical-03e17b481367257044e5fc139c19f2c7d3c46c59157df0c99a0520ddd2f7c454"></a>

## node property — ethernet_interface / 0483b89fedbc / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--network_interface--reference--group-001.md#canonical-60e47d120eb8777337b513994e1d7faaf2e3563d789ab1f932f6ca92123ac8d0): complete subsection reference.

<a id="canonical-b4a227d8d1a6c6cc2cd434e4a3f253e81ab115a930ba33ab47a0626182c9de8b"></a>

<a id="canonical-85411a75be2cb1bc909dd0cad8357a87db6ee1b49dc61a4340241d40de7be427"></a>

## priority property — ethernet_interface / 0483b89fedbc / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](resources--network_interface--reference--group-001.md#canonical-1c24d05bd343dc61c1d309e128a96fcf4e987dd24e1cc1d23bc2a65b370bee1f): complete subsection reference.

- [site_local_network](resources--network_interface--reference--group-001.md#canonical-620cdf9938409a1009b7d079b70e5c3eddb3c79c506a6245e760e69bc9b913a7): complete subsection reference.

- [static_ip](resources--network_interface--reference--group-001.md#canonical-cab9271e1730e6f9466f7224ac0cdb714ed64f8faa8cc1318c98043323baf448): complete subsection reference.

- [static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-82065ed290786f0db2ef4d7e91ef35f55e1eeb3eeb5b3835fbcfd751f3f864a6): complete subsection reference.

- [storage_network](resources--network_interface--reference--group-002.md#canonical-65a5224a94a01239bef6787e02312d7c041201fc3cded5a2a79a9d0992228f05): complete subsection reference.

- [untagged](resources--network_interface--reference--group-002.md#canonical-87f18d894e38d753a597d57b34415cec0096cc21131754751028e636f41b66ce): complete subsection reference.

<a id="canonical-ccc789142af389b5195663f33381a0cbcde4f82c0572a9a8748da98be57e9e9a"></a>

<a id="canonical-4537883e88b0584eef2c66232b85db97ac398b875ce1a87292927fce6abcdc1c"></a>

## vlan_id property — ethernet_interface / 0483b89fedbc / 8

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-252e8a4122724782dafa76f35e85c7f5a480cca98b4ad9afef8cf8dca6c0f79d"></a>

## Next pages — ethernet_interface / 0483b89fedbc / 9

- [ethernet_interface.cluster](resources--network_interface--reference--group-001.md#canonical-3d6e609b89c36cad162609dc7fef0ad20e646a9c2b1f76297e3cf0cc40a4c481)
- [ethernet_interface.dhcp_client](resources--network_interface--reference--group-001.md#canonical-9cd6c7a9dfa1e042007ad473dc91abad1b1892c94ecc157b826ea35e4d5fde01)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.is_primary](resources--network_interface--reference--group-001.md#canonical-96d9f13d322d043ffab7b9d4f26f2db3f8534caa14beae66c28f073b0a095cfe)
- [ethernet_interface.monitor](resources--network_interface--reference--group-001.md#canonical-62801a6fa2931fd2b7d35c8817f8b05c3d35b24eb7bb2de84c2cd554935d3567)
- [ethernet_interface.monitor_disabled](resources--network_interface--reference--group-001.md#canonical-46034b21cdd271dec1c6260712d1736070b3232acf27399d469376dad34aa497)
- [ethernet_interface.no_ipv6_address](resources--network_interface--reference--group-001.md#canonical-7f39517c0d7d7278f51272833c3b4b872306085209e35390e8cae77da081ef67)
- [ethernet_interface.not_primary](resources--network_interface--reference--group-001.md#canonical-60e47d120eb8777337b513994e1d7faaf2e3563d789ab1f932f6ca92123ac8d0)
- [ethernet_interface.site_local_inside_network](resources--network_interface--reference--group-001.md#canonical-1c24d05bd343dc61c1d309e128a96fcf4e987dd24e1cc1d23bc2a65b370bee1f)
- [ethernet_interface.site_local_network](resources--network_interface--reference--group-001.md#canonical-620cdf9938409a1009b7d079b70e5c3eddb3c79c506a6245e760e69bc9b913a7)
- [ethernet_interface.static_ip](resources--network_interface--reference--group-001.md#canonical-cab9271e1730e6f9466f7224ac0cdb714ed64f8faa8cc1318c98043323baf448)
- [ethernet_interface.static_ipv6_address](resources--network_interface--reference--group-002.md#canonical-82065ed290786f0db2ef4d7e91ef35f55e1eeb3eeb5b3835fbcfd751f3f864a6)
- [ethernet_interface.storage_network](resources--network_interface--reference--group-002.md#canonical-65a5224a94a01239bef6787e02312d7c041201fc3cded5a2a79a9d0992228f05)
- [ethernet_interface.untagged](resources--network_interface--reference--group-002.md#canonical-87f18d894e38d753a597d57b34415cec0096cc21131754751028e636f41b66ce)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-3d6e609b89c36cad162609dc7fef0ad20e646a9c2b1f76297e3cf0cc40a4c481"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d37ef7ae68dc09aaf95bb31565fc00dd2fade384e7351ad498f8da8c2fe02e24"></a>

## ethernet_interface.cluster — ethernet_interface.cluster / 5161197a5547 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.cluster

<a id="canonical-286e873f464b7834fc4b87a68ae713ba24d44fd6951924dfbc7c846578b03b39"></a>

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
cluster = {}
```

<a id="canonical-c081d035bb049b7cd7d93fcb8b51e6204f756b6ae8961d70eb155b96fcc02104"></a>

## Direct properties — ethernet_interface.cluster / 5161197a5547 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-705e256aa571ab6ba7b0a912086ec5c99641c059300bf47fff49e2b72565e423"></a>

## Next pages — ethernet_interface.cluster / 5161197a5547 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-9cd6c7a9dfa1e042007ad473dc91abad1b1892c94ecc157b826ea35e4d5fde01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f5028f8830aed562a475b205b08aa06569aaafbb3df23c60ae734f496453668"></a>

## ethernet_interface.dhcp_client — ethernet_interface.dhcp_client / 6a9a6aa0ad1a / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.dhcp_client

<a id="canonical-32197e25be5d5ac066858a2ded06305f48ad6d50ef275dcbcb214abeee1300d3"></a>

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
dhcp_client = {}
```

<a id="canonical-6d6141894c1eb12b865f34d992828e20b759bfa269d0710c01030fe59782ba1b"></a>

## Direct properties — ethernet_interface.dhcp_client / 6a9a6aa0ad1a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bd06a99f44ab0d66590de30064f38bb5aae1aebb7298e642d96665a2172f2b7"></a>

## Next pages — ethernet_interface.dhcp_client / 6a9a6aa0ad1a / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da5ef462573dbf0eb503e265b16f3f4c2cf52858803bdf6367ae6676e3d3a967"></a>

## ethernet_interface.dhcp_server — ethernet_interface.dhcp_server / 7210cc3b9279 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.dhcp_server

<a id="canonical-7915f844e72da692deb4e203c026d277b9380bd10cb9a077d609420518aeebf8"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dhcp server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-e1a84de5f80562e2770e9d9fb3eac85ff3eb23b90cb6da7343852774d201bcc5"></a>

## Direct properties — ethernet_interface.dhcp_server / 7210cc3b9279 / 3

- [automatic_from_end](resources--network_interface--reference--group-001.md#canonical-c2e3feb9a80fa8f8d6cadf3cb61b48b70a60882ce75fe414edcae9f27c8e92b3): complete subsection reference.

- [automatic_from_start](resources--network_interface--reference--group-001.md#canonical-7afe13829c1849a7f4f3881fd164f9794770cfaee750e08d8f0231622f28f008): complete subsection reference.

- [dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a): complete subsection reference.

<a id="canonical-0e561ae991d832d585d2f2968e78188cfe73f8e006d599e52b292c6aca883b02"></a>

<a id="canonical-3179eb822923271040ad1e4ea60dbe5adeda0306324654c51597f4dd04c265ff"></a>

## dhcp_option82_tag property — ethernet_interface.dhcp_server / 7210cc3b9279 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-6c0819c7d69877e3656f3602f7e66415fd51264f13a2a983122eb1bd49938102"></a>

<a id="canonical-a7a72c8c5a1cd293b71a33f1993f55877359c328574d1d8fa4394807794ee0f7"></a>

## fixed_ip_map property — ethernet_interface.dhcp_server / 7210cc3b9279 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--network_interface--reference--group-001.md#canonical-4e527c846b1dd941ea055fc32c71b1bd6cb15aa74b57e13771e95eaed814f37e): complete subsection reference.

<a id="canonical-a604628e694eaa5059da89ef2a80b397a7e73704d1eb914eff605b761d4fcdf2"></a>

## Next pages — ethernet_interface.dhcp_server / 7210cc3b9279 / 6

- [ethernet_interface.dhcp_server.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-c2e3feb9a80fa8f8d6cadf3cb61b48b70a60882ce75fe414edcae9f27c8e92b3)
- [ethernet_interface.dhcp_server.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-7afe13829c1849a7f4f3881fd164f9794770cfaee750e08d8f0231622f28f008)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- [ethernet_interface.dhcp_server.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-4e527c846b1dd941ea055fc32c71b1bd6cb15aa74b57e13771e95eaed814f37e)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-c2e3feb9a80fa8f8d6cadf3cb61b48b70a60882ce75fe414edcae9f27c8e92b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddc568c105b23bae6569d2fdfdc01ad71faa4be28ab1315a32f55f38a84caa24"></a>

## ethernet_interface.dhcp_server.automatic_from_end — ethernet_interface.dhcp_server.automatic_from_end / c61d71907834 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- ethernet_interface.dhcp_server.automatic_from_end

<a id="canonical-a1389086e0f253cfd6611c869eec2895de0b8e241c7155aebb064ee1862857a2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-7c346af988f6e589a9d805b640e0aaad29e11808fe2957431cf2ee011c374f6b"></a>

## Direct properties — ethernet_interface.dhcp_server.automatic_from_end / c61d71907834 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f26fb7b6daa0e90eb2e670302580d3724610fe5857eb50429d75478c41ee3c9"></a>

## Next pages — ethernet_interface.dhcp_server.automatic_from_end / c61d71907834 / 4

- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-7afe13829c1849a7f4f3881fd164f9794770cfaee750e08d8f0231622f28f008"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81fe987f9b70df885b953f01aec31aba9b39c25288afce2a6b342a8841966da9"></a>

## ethernet_interface.dhcp_server.automatic_from_start — ethernet_interface.dhcp_server.automatic_from_start / 45184183eef4 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-3963b691ce8ab57c4284192187d169f72379e1ad578495f58842b697fc247767"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-4c21fe1ca6e4b6d270e4b12e9eac508d8f2535f106debaa70393f24e6e6d92c7"></a>

## Direct properties — ethernet_interface.dhcp_server.automatic_from_start / 45184183eef4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9fdf7eb10c9ab05efd00bb4a8e1893b2d31cf6fea0889d6635ebf0cfa368cf97"></a>

## Next pages — ethernet_interface.dhcp_server.automatic_from_start / 45184183eef4 / 4

- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f10fc2e1f97a2db066684d65f2952e3a492898bfe68396a21272cb2f8ea6aeb8"></a>

## ethernet_interface.dhcp_server.dhcp_networks — ethernet_interface.dhcp_server.dhcp_networks / 383a2aea1d8d / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- ethernet_interface.dhcp_server.dhcp_networks

<a id="canonical-dfc27e6ec4562dc5c33208f7b5a4eefd0e8ac17618e7269799f1489d4dc07af8"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-b8d25d2388a96ad07416dab532e0f73cedb7b5276eb31be5464dac3292452da5"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks / 383a2aea1d8d / 3

<a id="canonical-0c12d34ad083221af6f88734f0da77d5ad76b4a9acb1a9f149565f979692e6cf"></a>

<a id="canonical-30af4d7753af25e192459d6387c494661d569e331f24c36299e1933aa65908c1"></a>

## dgw_address property — ethernet_interface.dhcp_server.dhcp_networks / 383a2aea1d8d / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-b6f82ce1ccfe203f236d0c6f5ee142af2b2c74fc147d623a111a2af22281eccc"></a>

<a id="canonical-4d6f2c36097cff6e76f69589fbe80b8e9e1d794985f37988137aad787a0f569e"></a>

## dns_address property — ethernet_interface.dhcp_server.dhcp_networks / 383a2aea1d8d / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

- [first_address](resources--network_interface--reference--group-001.md#canonical-8d28e0d1887475a8e67ff4cb127e3850f7f8e9e4f145a1792a49d346901a748a): complete subsection reference.

- [last_address](resources--network_interface--reference--group-001.md#canonical-784f62c77f754ca0e3fdde40c11c8c57be7d17bc496be39a3f9c53e6f89393dc): complete subsection reference.

<a id="canonical-ea593d83ce801f035965cd79da01d8034307d098860caa27d8b1a076f4befe0f"></a>

<a id="canonical-375b41cecd2a6fac8fb63a70c8f3fda25611e0ce88a5f8f343f993e29d567786"></a>

## network_prefix property — ethernet_interface.dhcp_server.dhcp_networks / 383a2aea1d8d / 6

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-4fe272c0cbf06036276f20145864ff2f6066cabbefe8104e6e7a76ac474cf8fa"></a>

<a id="canonical-a374f57b464cdbd0610b7937845187ed351af25f4c777a2717510dbddc93f769"></a>

## pool_settings property — ethernet_interface.dhcp_server.dhcp_networks / 383a2aea1d8d / 7

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--network_interface--reference--group-001.md#canonical-0add6de41b5cef175ff4ddde1bbbfbab290d616d108c9b7a579ab65ce3650847): complete subsection reference.

- [same_as_dgw](resources--network_interface--reference--group-001.md#canonical-79fd400ea944817ac09ebd5d551f02c3991f768f47d9f7099665b476975048bf): complete subsection reference.

<a id="canonical-d012c6ca2ce0fc083feaba83ad5246bbb7897f50ee50c961ce3f5b399d4f6b9d"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks / 383a2aea1d8d / 8

- [ethernet_interface.dhcp_server.dhcp_networks.first_address](resources--network_interface--reference--group-001.md#canonical-8d28e0d1887475a8e67ff4cb127e3850f7f8e9e4f145a1792a49d346901a748a)
- [ethernet_interface.dhcp_server.dhcp_networks.last_address](resources--network_interface--reference--group-001.md#canonical-784f62c77f754ca0e3fdde40c11c8c57be7d17bc496be39a3f9c53e6f89393dc)
- [ethernet_interface.dhcp_server.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-0add6de41b5cef175ff4ddde1bbbfbab290d616d108c9b7a579ab65ce3650847)
- [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--network_interface--reference--group-001.md#canonical-79fd400ea944817ac09ebd5d551f02c3991f768f47d9f7099665b476975048bf)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-8d28e0d1887475a8e67ff4cb127e3850f7f8e9e4f145a1792a49d346901a748a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13379aa3573537943e2176fa28f1e3a3dc7d5e620d90221d7d9f499590869118"></a>

## ethernet_interface.dhcp_server.dhcp_networks.first_address — ethernet_interface.dhcp_server.dhcp_networks.first_address / e3c453e40469 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- ethernet_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-f51d24bc743daf61d8f3f80e215bd34f82a9b854c133dc9d542c73c8bcbed4d9"></a>

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
first_address = {}
```

<a id="canonical-41f3178d9a445397e2b8c53618fcfe2c37593799a2de7440bad4c94235beb045"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.first_address / e3c453e40469 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21f743cf5f364c605d2a4ebc54332f3d635739fb93c1a8d0f765be77bf154155"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.first_address / e3c453e40469 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-784f62c77f754ca0e3fdde40c11c8c57be7d17bc496be39a3f9c53e6f89393dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee660479a2ef2009ebc16bc2faeecb3565e97400387637caa7d0257e43298db8"></a>

## ethernet_interface.dhcp_server.dhcp_networks.last_address — ethernet_interface.dhcp_server.dhcp_networks.last_address / 44d8f73e9e78 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- ethernet_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-9ef6a54735209c3c85aa73dff81f03724c3578772da3ae050140820c1a951bbb"></a>

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
last_address = {}
```

<a id="canonical-ccffe0b96218ff239614bb1096d82373fe9117935f9673f679ed7e006d4d407b"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.last_address / 44d8f73e9e78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-637e95eaafe39aea3e2ff6e51c91d13b0e12fac37260a1f1eeebb239b5df2c65"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.last_address / 44d8f73e9e78 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-0add6de41b5cef175ff4ddde1bbbfbab290d616d108c9b7a579ab65ce3650847"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4689ee664c4ca421d154c130178fada6c5bf2387d34f694908d9febb410ed740"></a>

## ethernet_interface.dhcp_server.dhcp_networks.pools — ethernet_interface.dhcp_server.dhcp_networks.pools / 8d7aab9f640f / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-ab8a2f63fcaed4da5cf579820aa5ad3f5d1f4e4210372bb7667cc321d30640c5"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-61133f8b365d7ccd2506c77ec2c90ca9c545ebcbf593e52c7495f15a15c46969"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.pools / 8d7aab9f640f / 3

<a id="canonical-32e578cc24b0dd97de5581af53349ed49bd3c288d1762fe326669894fecf3bc4"></a>

<a id="canonical-32cba509c2180a56ad028cd29e42be8dbb83346c0ffb30966f8b4af854096f7c"></a>

## end_ip property — ethernet_interface.dhcp_server.dhcp_networks.pools / 8d7aab9f640f / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-8afa8bbc0805b591a13da147260f95225e7295abab9b68185c621d17ed625658"></a>

<a id="canonical-0d21d65c3a4b8000c570675a9fab5255b751e2619306aac315713d0396d4dd2c"></a>

## exclude property — ethernet_interface.dhcp_server.dhcp_networks.pools / 8d7aab9f640f / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-46b2636cf331774fc5f4e0fa97a37085c6672a03350cc1601abb267be4016e9b"></a>

<a id="canonical-1d982f7dd20b9dd31f9a92ddcc2d2939e4ce55d90eb8883c06d23892605c0bf3"></a>

## start_ip property — ethernet_interface.dhcp_server.dhcp_networks.pools / 8d7aab9f640f / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-ea6f5b238bcc1f08040dc1691ca13782f98091861715a04bf36f5fca442578f9"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.pools / 8d7aab9f640f / 7

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-79fd400ea944817ac09ebd5d551f02c3991f768f47d9f7099665b476975048bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-298c10067ab98323074415e05323df9c73cd3d865f1038a92fa932cfce591c54"></a>

## ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw — ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw / d148ebb74971 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-aad0e1b6cc4c37da1332dfc9f63b0ff0539a92bd2f643972f39ba463764b49d8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

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
same_as_dgw = {}
```

<a id="canonical-c908dc107abde5f5e447e857b3ebf9510ee4c0d6e85ae00c259afedf3249c686"></a>

## Direct properties — ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw / d148ebb74971 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f81c6cc118978d544e778caf844b1bebd6f7f93da7eb85198a1b782a89ea765"></a>

## Next pages — ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw / d148ebb74971 / 4

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-51b6cf5b180815d210dbacc281fa1aedb8c66c837d3874e0368a1410fa1ddf7a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-4e527c846b1dd941ea055fc32c71b1bd6cb15aa74b57e13771e95eaed814f37e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-103de5f674d8a4b110d521fdc7bf75073c198cd33fd4f7e08b8b4f59e18f9fb0"></a>

## ethernet_interface.dhcp_server.interface_ip_map — ethernet_interface.dhcp_server.interface_ip_map / ac8572f9093d / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-aa45145b5108abf11edf5528e142a3621978db41eb9c9613f7c3025e4b53a68b"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-55aa84396923ed02d83079e229b8c8902b9fa9da0796f790f9c2b8b6a714f1c8"></a>

## Direct properties — ethernet_interface.dhcp_server.interface_ip_map / ac8572f9093d / 3

<a id="canonical-fd0ba71fd7e58bd5b96bae80fb0c04dcd114af16d8ddfd74df5421a15b81d417"></a>

<a id="canonical-69577e55f90a10ede729b0e4f00af30306003ad59fba4469670969f7b8bec930"></a>

## interface_ip_map property — ethernet_interface.dhcp_server.interface_ip_map / ac8572f9093d / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-ba302a682c6b8cea1d42bdfc6dd27838a260057595f012269e02e7fc561ac32d"></a>

## Next pages — ethernet_interface.dhcp_server.interface_ip_map / ac8572f9093d / 5

- [ethernet_interface.dhcp_server](resources--network_interface--reference--group-001.md#canonical-67e6943ba32172c6727e5d13a56aaa4ab92f4c27bd86d90fb9808c923e23e29a)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b6472f569a044f3059770bee3ff8f1f0c0f171dc7a63d6aba8b023c86f47864"></a>

## ethernet_interface.ipv6_auto_config — ethernet_interface.ipv6_auto_config / cd9defb11c18 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.ipv6_auto_config

<a id="canonical-0383e001aa6430ea05121ee741becd34ed34615bfc522fddb69fbc000d8f2bfc"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-21ed3138de944cb4af07497a212b1dfcc56319597dd2ebc1b26c735a4997d3ca"></a>

## Direct properties — ethernet_interface.ipv6_auto_config / cd9defb11c18 / 3

- [host](resources--network_interface--reference--group-001.md#canonical-241c4a149374042008a151dc397daaadc146bb6e20a2229800781607b31d8350): complete subsection reference.

- [router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa): complete subsection reference.

<a id="canonical-7747e594e29d3652b2ef0eaf57a8ba0d1bd13239545ae4d193fdf7424e244270"></a>

## Next pages — ethernet_interface.ipv6_auto_config / cd9defb11c18 / 4

- [ethernet_interface.ipv6_auto_config.host](resources--network_interface--reference--group-001.md#canonical-241c4a149374042008a151dc397daaadc146bb6e20a2229800781607b31d8350)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-241c4a149374042008a151dc397daaadc146bb6e20a2229800781607b31d8350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5374afaa34e6cce097926552faafdbaa4d1cec78f23107a4d305b08bad45fa09"></a>

## ethernet_interface.ipv6_auto_config.host — ethernet_interface.ipv6_auto_config.host / dbe11f277d65 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- ethernet_interface.ipv6_auto_config.host

<a id="canonical-2d63b82cc646675e9b627784839a600b6f1e1c540593a29e908f9ed93883ed66"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-1aba9045ea1ea28dd2ec06a3ca822d8ebdc3a89b7a3a665fcfcd4f2a86c9d2f6"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.host / dbe11f277d65 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab1c52610b26e748103e970e3d94eea4592c461897899f4b76f1d8b024adf9e3"></a>

## Next pages — ethernet_interface.ipv6_auto_config.host / dbe11f277d65 / 4

- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc3c5c36d4ef37dfdf6f85ef671a1f286213b9b28c69104a4e7e70752702399d"></a>

## ethernet_interface.ipv6_auto_config.router — ethernet_interface.ipv6_auto_config.router / fa25fa24c445 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- ethernet_interface.ipv6_auto_config.router

<a id="canonical-589d774633953ed07ee758ce5a1024ef59608b53cc9c84bd932098b6110a4dcc"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-5b9c4be8189d9ec6812d8bac21b178e4bad03484678266cd8956c46ec8ea9cfa"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router / fa25fa24c445 / 3

- [dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875): complete subsection reference.

<a id="canonical-16936745aca8643384bc1d4579a9493fa57ca2444d0f1938f1ce6fda8b73b810"></a>

<a id="canonical-9b863d4fdbed1f16623baac3a4b96910a2fc47c45f597711f5340b50cd0652d0"></a>

## network_prefix property — ethernet_interface.ipv6_auto_config.router / fa25fa24c445 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512): complete subsection reference.

<a id="canonical-e073a6f104ff16466cd38c0e745ff967303827b261b553951608f5d97243a1cc"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router / fa25fa24c445 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31adee625324403ebb7d456d69ef64ce2895f541167ea32762f51ea5665153be"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config — ethernet_interface.ipv6_auto_config.router.dns_config / f7a174abd62b / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- ethernet_interface.ipv6_auto_config.router.dns_config

<a id="canonical-dc53d7e926c8dfd350196bd1ea64c7e835eadc27b224663dede4a9b56616aa79"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-4e332578c23c0797408916c9bec789cf1c7e83e8d61c3c12eb42e37161e44dd3"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config / f7a174abd62b / 3

- [configured_list](resources--network_interface--reference--group-001.md#canonical-2dc56ae3f9cd0a8a042fa5294b80406ccd0543fc0f9d9fceff83d390867a35a8): complete subsection reference.

- [local_dns](resources--network_interface--reference--group-001.md#canonical-61f1effd064a66b98b3936e2dff8e738e9e325705903787a475814b5457cbacf): complete subsection reference.

<a id="canonical-f8283f4820712f09641b5029e605d4af858681d00f56d4658de090db151ce9b4"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config / f7a174abd62b / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--network_interface--reference--group-001.md#canonical-2dc56ae3f9cd0a8a042fa5294b80406ccd0543fc0f9d9fceff83d390867a35a8)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-61f1effd064a66b98b3936e2dff8e738e9e325705903787a475814b5457cbacf)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-2dc56ae3f9cd0a8a042fa5294b80406ccd0543fc0f9d9fceff83d390867a35a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce12b2f921b8dad8172173e4d46fdfbed5ed080033f8edadb8d273dda269cc18"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.configured_list — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / 662b16dc4bc0 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875)
- ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-c69670bff4b518bd2539b057799f58abf4ef8b2c8f43dba306676c0b23addc64"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-72eadc27e096692d20267de777fbffb3f6c3a049ac3f433946b14eee5baaca06"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / 662b16dc4bc0 / 3

<a id="canonical-0fc0b63ee232663ab13e8716b21e785e67284fb2416378f6a66e8ee31a798b47"></a>

<a id="canonical-7aa7ae597c79ea7d56bc4e4e2fb851b75da5eaf7abc05a2ecb0b719b13e900d1"></a>

## dns_list property — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / 662b16dc4bc0 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e603e8a945edab4cb93de2440220ada5b0396c50226c1327538d8cbede2bbe4b"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.configured_list / 662b16dc4bc0 / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-61f1effd064a66b98b3936e2dff8e738e9e325705903787a475814b5457cbacf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43ff920c7eb18ba1478e29b06585321c389f77626edb3d200f6f98536da5e534"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 9d2db61922ce / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-9f98eeb5e176262b2c14a05c871c5c0e481659097b92fb2fd93915fb71c4c136"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-21cb9e6873c8a42cda3ff08e41792754aa5b1952dcd97f9a4c2b2351a8af4d7f"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 9d2db61922ce / 3

<a id="canonical-2138a91165a0d6f98a01feef655973e5339049d6a8d5af6a9294be30c873b57b"></a>

<a id="canonical-72b2d263020051ca20de1b8063571e70d702b05d277855137f6437ca9d7df435"></a>

## configured_address property — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 9d2db61922ce / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

- [first_address](resources--network_interface--reference--group-001.md#canonical-236c0feef46768a1318af3afd99a91ea1e69cfe5dba764569a21bedfeb74e8de): complete subsection reference.

- [last_address](resources--network_interface--reference--group-001.md#canonical-49e72f1fce38021676f4e4c4dc490e5473ef8974506a64d83eadc12677f4cce0): complete subsection reference.

<a id="canonical-3feef39374418ccad6c9bcedd5340bbea4652301f38c5dca11ba3f2ab2f77502"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns / 9d2db61922ce / 5

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--network_interface--reference--group-001.md#canonical-236c0feef46768a1318af3afd99a91ea1e69cfe5dba764569a21bedfeb74e8de)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--network_interface--reference--group-001.md#canonical-49e72f1fce38021676f4e4c4dc490e5473ef8974506a64d83eadc12677f4cce0)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-236c0feef46768a1318af3afd99a91ea1e69cfe5dba764569a21bedfeb74e8de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d0181f7bef08bf658a5bf3d6a9c62c6f9996824a6a38926ef79efe830db673a"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address / 2c1c7f500fd7 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-61f1effd064a66b98b3936e2dff8e738e9e325705903787a475814b5457cbacf)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0286f35c933e2f6c0be807ac2975447e36242ba6e3d9016db0ec22a9ff12b30a"></a>

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
first_address = {}
```

<a id="canonical-44cc325b94f874b17f2223007319e8e7fc57db8f6bee56721a45caaabdcdaabd"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address / 2c1c7f500fd7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d7dac82b0ac606e7d11ec0dc14a4ee4a72bf3cec1fc3905e9b87965b37af39c9"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address / 2c1c7f500fd7 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-61f1effd064a66b98b3936e2dff8e738e9e325705903787a475814b5457cbacf)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-49e72f1fce38021676f4e4c4dc490e5473ef8974506a64d83eadc12677f4cce0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50762bc31c72ab7f7545a8549ffd74c09908003722ccfda21f963bcc20ae74d4"></a>

## ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address / ffb450de4735 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--reference--group-001.md#canonical-95d4d7323a3f8b2d0062f4dde92ec38a688e1d4d03164598e67c60fcd299a875)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-61f1effd064a66b98b3936e2dff8e738e9e325705903787a475814b5457cbacf)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-94bc9ace9b2f5e8776d70138d3bea753e9c76e0cf8d30ed8bf9660b6650ddec1"></a>

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
last_address = {}
```

<a id="canonical-def938d942d9126f962f7b0231dd238fcab0baa5a7900b77c43c197a08ce6f2e"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address / ffb450de4735 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48033654332cfa901f4d1979d012e4385eb90f0975ca62414d4846466887e018"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address / ffb450de4735 / 4

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--reference--group-001.md#canonical-61f1effd064a66b98b3936e2dff8e738e9e325705903787a475814b5457cbacf)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc6d496c0314d18e2775d8ce51fc712a11901794e617d0dfb28bff9b666d85db"></a>

## ethernet_interface.ipv6_auto_config.router.stateful — ethernet_interface.ipv6_auto_config.router.stateful / 1fe5c4834aa8 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-17c38b9f20cdd06f05cfd07ce97855102c6b90ece5bd3c78800f8fec5b534a5a"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-457131f5f86415bc3805549e67d494d95b4c9656e91f03040fc9dd153938883f"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful / 1fe5c4834aa8 / 3

- [automatic_from_end](resources--network_interface--reference--group-001.md#canonical-95df0dad1f7c68a80426201bdf3f81fb6fa4af410573562f75740ea5d500d14c): complete subsection reference.

- [automatic_from_start](resources--network_interface--reference--group-001.md#canonical-4b29dbb99e5a42477d79610557e3c09a06cc66f9bd8c7e644049c5670efe3489): complete subsection reference.

- [dhcp_networks](resources--network_interface--reference--group-001.md#canonical-3fd2f27d5692de3f3a82fbba7d7700b4ff1dfbaeaf704072abdead28c94cc386): complete subsection reference.

<a id="canonical-0d4dc6af4c18220c76c2d64fbb35850bdc4387b6f40a2f2a80012029730ce73c"></a>

<a id="canonical-72a4a8b3d584a386063ff4691de0c08e9ac3e44f8c7730a60b2d679eb2d6da83"></a>

## fixed_ip_map property — ethernet_interface.ipv6_auto_config.router.stateful / 1fe5c4834aa8 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](resources--network_interface--reference--group-001.md#canonical-c2a8d114e51ac8d3347a89aef51510636699cf87567078052584ff3a8a78b008): complete subsection reference.

<a id="canonical-dbb82b4dc815c576caf79763dc6db90f606a6188d025a65a69ed90d66a42c742"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful / 1fe5c4834aa8 / 5

- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--network_interface--reference--group-001.md#canonical-95df0dad1f7c68a80426201bdf3f81fb6fa4af410573562f75740ea5d500d14c)
- [ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--network_interface--reference--group-001.md#canonical-4b29dbb99e5a42477d79610557e3c09a06cc66f9bd8c7e644049c5670efe3489)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-3fd2f27d5692de3f3a82fbba7d7700b4ff1dfbaeaf704072abdead28c94cc386)
- [ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--network_interface--reference--group-001.md#canonical-c2a8d114e51ac8d3347a89aef51510636699cf87567078052584ff3a8a78b008)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-95df0dad1f7c68a80426201bdf3f81fb6fa4af410573562f75740ea5d500d14c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4da46cc78576b099d58476b5326e24e4973aeff3ae24767563c8c262e6048ef7"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end / cee26380845a / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-6ca6f12adedaa622904057d41fcf183bed3b220f75f102e27814681a062b1553"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-ceffd2108409eb1922b6b7f9620f755de20f6bb4b46379f19cfc62a590255141"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end / cee26380845a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0c40a1cecde792c6ba7afce62c34817a7f9b329f1a249488367813ac772941e"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end / cee26380845a / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-4b29dbb99e5a42477d79610557e3c09a06cc66f9bd8c7e644049c5670efe3489"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9b01c7ab7ed0ca0c5815e2e8674003cdae8551c83b896ff83cd91b12f58260d"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start / 7ddd7e8ba6ca / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-8bf4df97b8e24fd925fd58d569d20a9595e55609b938e8381c67f6ba415b9e50"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-670f41150f61e7a24cb99ef85344cd9b69af28655d996bc0f5fe279e63de11bb"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start / 7ddd7e8ba6ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-99de8ddfdddccff7f6dbd41b9ad9d33f7844008627212dbbd5d8f5cee39953af"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start / 7ddd7e8ba6ca / 4

- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-3fd2f27d5692de3f3a82fbba7d7700b4ff1dfbaeaf704072abdead28c94cc386"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-472062aa1f043a112a3a6630e3862ab7efae9aa9882ad26ce3d8d30443c39edb"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / dcdc8cccfba0 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-8b0f37b60c5b8865deb3d15af9be8753aa487a374d79337a4aa75e8db5d607fe"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-e46d6ff36b56b006e2c9ce8b09bdbde7354ceb66f4d5cd4a2b89144d38cd3bde"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / dcdc8cccfba0 / 3

<a id="canonical-22175ff499b48401ed21d2d8c84db14ce214635c4631614ba66225003e15c1d9"></a>

<a id="canonical-8e10f81bb955db0a8788745dc085f6a36ba3771bb4e02c870d9a14cbff5f500e"></a>

## network_prefix property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / dcdc8cccfba0 / 4

Type: `"string"`. Optional.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-0a747ef544cb43b079530106d8806dfa4b18c99b9e8b488f13feec03bb6ffdc9"></a>

<a id="canonical-89209e240aa4c85cc6a34a9855ef6ca491dbbf161b594bb661eb43a3b11e4186"></a>

## pool_settings property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / dcdc8cccfba0 / 5

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--network_interface--reference--group-001.md#canonical-a2ad1340cc8d98107e2988ced8bd45b3999ba012548851d7a9486d5f6eb9b3fa): complete subsection reference.

<a id="canonical-996ed86a5bb4362966a8f8b8db025f94bdb392759ebb31e8965e896454dfb7af"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks / dcdc8cccfba0 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--network_interface--reference--group-001.md#canonical-a2ad1340cc8d98107e2988ced8bd45b3999ba012548851d7a9486d5f6eb9b3fa)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-a2ad1340cc8d98107e2988ced8bd45b3999ba012548851d7a9486d5f6eb9b3fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2c99d7d43de4bd2afc0dfe99882ed0376255b682c10c0f2f4d5080605460a18"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 7b793c4b6f63 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-3fd2f27d5692de3f3a82fbba7d7700b4ff1dfbaeaf704072abdead28c94cc386)
- ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-07227e352a23c33038c8374c3baa62a5eb541ea000cf80e028b03e600895a78f"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-d27bd48fad19e282726fca40eadbe4b69e22d34eb44f88266e66c98504e38016"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 7b793c4b6f63 / 3

<a id="canonical-6f011446dce544ad6c7dec656d0c76eca541e857317f21d93d9c4db6c6d02127"></a>

<a id="canonical-38d49ceca496a3e82fc345b3cd51f6018c8f79887f1e71a217e657aa728d51e2"></a>

## end_ip property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 7b793c4b6f63 / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-7120feb55e299061997c40d2a160d0f5865804196dcce1ee4b82292de692e4db"></a>

<a id="canonical-0a3242e2f0df80f9d1948ed959fd5e986782aff555c7f97142ad2789f5c427f3"></a>

## start_ip property — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 7b793c4b6f63 / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-4f65d3a28574a6a6429bbee09c2cd5865033d5d1fb77c457b96096b7780507b5"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools / 7b793c4b6f63 / 6

- [ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--network_interface--reference--group-001.md#canonical-3fd2f27d5692de3f3a82fbba7d7700b4ff1dfbaeaf704072abdead28c94cc386)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-c2a8d114e51ac8d3347a89aef51510636699cf87567078052584ff3a8a78b008"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b519f07c12b27ab7d87da73b27eadfbac12f47b374ad84c6531a72bd69e81381"></a>

## ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 70ba0c9e5837 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--reference--group-001.md#canonical-d507c80f6145ebfaca69d043fc19fbf9b72f5067c13af30649e79750a254c1cc)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--reference--group-001.md#canonical-9e1e8b4482e969e54bdedbd5f6e51a8ade9180c841dbc9b7e3c9b7b029de24aa)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-8d16e5d6f2f3ce8eae7c8caa200948c6ffa86884f6c53800cf07a8788cd77341"></a>

Type: `"object"`. single nested block, Optional.

Map of Interface IPv6 assignments per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-10dc69ed733c5ce081af750b894db18b9fd27182a6631eb2de33f01826c695a6"></a>

## Direct properties — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 70ba0c9e5837 / 3

<a id="canonical-3e9e67af64c477a87d90da26f09bc231a0599be249bb2118037a5965e9acbb06"></a>

<a id="canonical-152dab58ccc38a9c17e5d694bf18bb102fd6fe277f7f180edcdd5faa4272e6ce"></a>

## interface_ip_map property — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 70ba0c9e5837 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-8e88335be17523e4f9249adbe9cf0c1c5a6fafbc1548309e80a8f422f70959ac"></a>

## Next pages — ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map / 70ba0c9e5837 / 5

- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--reference--group-001.md#canonical-ebd7ef0686c6d804ac98a5d18f5b1a7ab434eea378a3d63b3b9863cfd1cbc512)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-96d9f13d322d043ffab7b9d4f26f2db3f8534caa14beae66c28f073b0a095cfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e33f6b2869637d42e1d5e7c36998e49e95fdc5aa83adadbdfbe60d63c41f16b"></a>

## ethernet_interface.is_primary — ethernet_interface.is_primary / 31fa515fd854 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.is_primary

<a id="canonical-2a25c8b30670c082b6ccd885febb75c927e6b9dd5b45566ac51d857d95adcbaa"></a>

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
is_primary = {}
```

<a id="canonical-e67e6e1c83a6eb56c4d00f849b9f17af99978785bad17a26347a61fdd6de2518"></a>

## Direct properties — ethernet_interface.is_primary / 31fa515fd854 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fffd15411238453a4bd8ec2154f4a2a42e71aa8ba85581eab599c454c3efab00"></a>

## Next pages — ethernet_interface.is_primary / 31fa515fd854 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-62801a6fa2931fd2b7d35c8817f8b05c3d35b24eb7bb2de84c2cd554935d3567"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e09996e2476798dfeab469f77ad947570ac127ba15f6e22f5e2897e4cd36135"></a>

## ethernet_interface.monitor — ethernet_interface.monitor / 06ea60fc43dc / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.monitor

<a id="canonical-2c2ea94141d5f5c631761897b48784fefa665733c6379d00a7a4e5b646f69c9f"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

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
monitor = {}
```

<a id="canonical-77f9b05ff34ce9c827a573d32684a4bd01182b2e9b14b1bab84d471fd44fdfe6"></a>

## Direct properties — ethernet_interface.monitor / 06ea60fc43dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55079c54509e95596685ab8a823b0128b868fbc3394c61bcdd086c2e9c70e9ef"></a>

## Next pages — ethernet_interface.monitor / 06ea60fc43dc / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-46034b21cdd271dec1c6260712d1736070b3232acf27399d469376dad34aa497"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b515f391e23ab3f1ed59180a1e3d018840df7f15b6cecad0152487ac1ed2703"></a>

## ethernet_interface.monitor_disabled — ethernet_interface.monitor_disabled / ab6890d2ecce / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.monitor_disabled

<a id="canonical-00b1b710483d9029260b5e5c65c45e4d6cdcb127f9d18922f127c6796221fe7b"></a>

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
monitor_disabled = {}
```

<a id="canonical-0aa838b204ba92cc156c616f2a85141cff11d672653ed3f5ae1e525c9bd391c9"></a>

## Direct properties — ethernet_interface.monitor_disabled / ab6890d2ecce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9ddfa1c7af3878374efa195bb8c50e9a9ab8bcb25a1bfb3d0e119862a33988a"></a>

## Next pages — ethernet_interface.monitor_disabled / ab6890d2ecce / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-7f39517c0d7d7278f51272833c3b4b872306085209e35390e8cae77da081ef67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3906756b993b84cc089d1d751fca6523557e1ea129e1e99b6c20c4b2c907f186"></a>

## ethernet_interface.no_ipv6_address — ethernet_interface.no_ipv6_address / 7b8371341a9f / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.no_ipv6_address

<a id="canonical-81e8bde0c542ae046890ccd7b6ca97bc9f4c67ec4f003094bbfe2f85610d581b"></a>

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
no_ipv6_address = {}
```

<a id="canonical-69bb123d535792152ac544f2c8c84256e5b2fcb728c53197ea5fc8a2afeb62d1"></a>

## Direct properties — ethernet_interface.no_ipv6_address / 7b8371341a9f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c31831576221050a33c6280c374ceb17eda62dc6233688d8fab6cf1effe91a3"></a>

## Next pages — ethernet_interface.no_ipv6_address / 7b8371341a9f / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-60e47d120eb8777337b513994e1d7faaf2e3563d789ab1f932f6ca92123ac8d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a645b594de83cec881b3f19b40580e0b9bcb8141806e33affea9078cbf2c6ba8"></a>

## ethernet_interface.not_primary — ethernet_interface.not_primary / 1cd48a6eb810 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.not_primary

<a id="canonical-d6feb1ad4a0f6da80f15d4780c38a93c354d8bd742d75c6b8b7bddce772ecfe5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for not primary.

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
not_primary = {}
```

<a id="canonical-f7b564989cdf8b4fee3fe38e8faefd7b340d20130da8fd55b9757cc3d4d158b9"></a>

## Direct properties — ethernet_interface.not_primary / 1cd48a6eb810 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-088024918a8ee2e777ed10ffd4bff061f23f7d5d08066304dcd1f8856f87c41c"></a>

## Next pages — ethernet_interface.not_primary / 1cd48a6eb810 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-1c24d05bd343dc61c1d309e128a96fcf4e987dd24e1cc1d23bc2a65b370bee1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d61bbb4630af05c1cfefda0ffbafe572b5b18f18faaa7722b68eee4bbfbc2035"></a>

## ethernet_interface.site_local_inside_network — ethernet_interface.site_local_inside_network / c8c68dbfe762 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.site_local_inside_network

<a id="canonical-85f6dd6599ba156de5c47aa2aa4ec8844c6367bb5435184c66788769101f2b66"></a>

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
site_local_inside_network = {}
```

<a id="canonical-6af6dad898f212c1690f6496ccbde1d4c5b183ac30f84fa0b7e1fcd3aff62258"></a>

## Direct properties — ethernet_interface.site_local_inside_network / c8c68dbfe762 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7826ff0c64885a0143d12e8af62bb8cf6aba6a691b600598bfb6885fa22ca80"></a>

## Next pages — ethernet_interface.site_local_inside_network / c8c68dbfe762 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-620cdf9938409a1009b7d079b70e5c3eddb3c79c506a6245e760e69bc9b913a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03652211ed3bfaa17c638c13db2ab25290b440a187d762c61ef13ff54cc27746"></a>

## ethernet_interface.site_local_network — ethernet_interface.site_local_network / f7ec9017e460 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.site_local_network

<a id="canonical-83918a5d5cc033b3d4ed328d9c15172284ab234e6456012e467d4851c2d58a34"></a>

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
site_local_network = {}
```

<a id="canonical-05b5f347f2e1729ab2841d3c0f96954f1601e6ee08ca867b08ee9881bb7bd182"></a>

## Direct properties — ethernet_interface.site_local_network / f7ec9017e460 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-548e477405eab682e6f0abd689b7d13e8295f7ce5fd4f31440539e5710ea639d"></a>

## Next pages — ethernet_interface.site_local_network / f7ec9017e460 / 4

- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)

<a id="canonical-cab9271e1730e6f9466f7224ac0cdb714ed64f8faa8cc1318c98043323baf448"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b5639e6e424959bed585cb72568f0d4ac9736acf354554580663d6827f0a56e"></a>

## ethernet_interface.static_ip — ethernet_interface.static_ip / b8d8b7423779 / 2

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4)
- [Property reference](resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [ethernet_interface](resources--network_interface--reference--group-001.md#canonical-261f1ca86182d6d80002e855bcbcf682a31cf4629e1fc0e1672d613963e52369)
- ethernet_interface.static_ip

<a id="canonical-bb93a5143905ed4c3fba99beb8e9a3731ed985a6e0f8a8a48f019d236e243290"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ip {
  # Configure direct properties listed below.
}
```
