---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea460c7824a00c51ee3929454a06893b53eb8dcbd2321bf868aca82d4f6a0548"></a>

## Property reference — Property reference / a2ca6f03f464 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- Property reference

<a id="canonical-15e030f23334f77b9e07e1b4966344d7d81a09200129328053e8d3918304b9e3"></a>

## Direct properties — Property reference / a2ca6f03f464 / 3

<a id="canonical-ead45fa5cae9153c4c5a37790e7bc7e31e0c2c6501271d66ffad97dd71c77a0e"></a>

<a id="canonical-36399cb2f1059c63675c3cd4da51d2e79b5191d45045d527e6df1081f9bf07ae"></a>

## annotations property — Property reference / a2ca6f03f464 / 4

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

- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-d30f89abfb3bcd2816487185f399e481a1d04f005852d935b50153fb1f63c465): complete subsection reference.

<a id="canonical-96d175eaee91ca8ea1a2d292d9966cf57c3d25f2848a270516dc55c3961510bf"></a>

<a id="canonical-5b8fa4ea7c863223a125f1c1ad452eab8f2796ca81931f3b39d6bf74188726ac"></a>

## description property — Property reference / a2ca6f03f464 / 5

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

<a id="canonical-45f3b23167012f772dd5d8c46caaf185757d50a73960d4d354a0fb7e29751066"></a>

<a id="canonical-fd6d282151a1aec37347b02f76984161e5e65bae0977b3630deb0efaac804d74"></a>

## disable property — Property reference / a2ca6f03f464 / 6

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

<a id="canonical-5a4aa5d4a120e0b2f9d37636a012e9dceb16fe61e82a35b77d152fdfa73a4977"></a>

<a id="canonical-e2dfe3e056d0ef531e22a576978d96bad835a0b1e9dff53337334af9d3c9c774"></a>

## id property — Property reference / a2ca6f03f464 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-eb6edf2a0a30d6c7331a458cf44f7df21204e147059efebbc4d236468571ae44"></a>

<a id="canonical-70311fe524f38ae0140d81a7b27de61f9141c6ece365b20cd077d8ef13b8508b"></a>

## labels property — Property reference / a2ca6f03f464 / 8

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

<a id="canonical-bb75c8b2866ef908404293ef81adf26231bbf6e268013773dfa014144e716de7"></a>

<a id="canonical-124e87dd2385e2d5053fd96dafd0d237d40da0350d6046c39ad7a12b0e5e6271"></a>

## name property — Property reference / a2ca6f03f464 / 9

Type: `"string"`. Required.

Name of the BGP. Must be unique within the namespace.

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

<a id="canonical-c8e8e6024e5d308e097ccc611278b3b0eb3404034eea68e254958387770f50f9"></a>

<a id="canonical-b8feaa3f3c940fa60a598bd81cb1d5006532687cd72e4c925b2c68ff1aa650d7"></a>

## namespace property — Property reference / a2ca6f03f464 / 10

Type: `"string"`. Required.

Namespace where the BGP is created.

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

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213): complete subsection reference.

- [timeouts](resources--bgp--reference--group-001.md#canonical-f09bbcefdb3d000d17aea209da85cf650c0c55df309b01717b7c4b34ae89314f): complete subsection reference.

- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed): complete subsection reference.

<a id="canonical-09c9009dc9a0001f1c2b8a016bf92be803c072ba420ca6856a77ccc9225fa566"></a>

## All schema paths — Property reference / a2ca6f03f464 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bgp--reference--group-001.md#canonical-ead45fa5cae9153c4c5a37790e7bc7e31e0c2c6501271d66ffad97dd71c77a0e) |
| `bgp_parameters` | [bgp_parameters](resources--bgp--reference--group-001.md#canonical-36ea4f1dee24c0f4814c1ae5d4da83d76afa25f4903c93ec3530e05ff5c5c31b) |
| `bgp_parameters.asn` | [bgp_parameters.asn](resources--bgp--reference--group-001.md#canonical-a414488956bc214e737ed52f072786672b2e212c5b083598bd89db5caa4bdf9f) |
| `bgp_parameters.from_site` | [bgp_parameters.from_site](resources--bgp--reference--group-001.md#canonical-978b1d05b50930adfb079fe1b14b23d3a708d96d2a037d58150ca78ba4c6d0ff) |
| `bgp_parameters.ip_address` | [bgp_parameters.ip_address](resources--bgp--reference--group-001.md#canonical-faa40a4955698771f81f6a3b2d7c2b991e5b7d691d5a83b0c295a6a812e25e55) |
| `bgp_parameters.local_address` | [bgp_parameters.local_address](resources--bgp--reference--group-001.md#canonical-32d767cb782abceb065613e530dae1a67ae4d88e5b93ed873f70d9d34ceba80e) |
| `description` | [description](resources--bgp--reference--group-001.md#canonical-96d175eaee91ca8ea1a2d292d9966cf57c3d25f2848a270516dc55c3961510bf) |
| `disable` | [disable](resources--bgp--reference--group-001.md#canonical-45f3b23167012f772dd5d8c46caaf185757d50a73960d4d354a0fb7e29751066) |
| `id` | [id](resources--bgp--reference--group-001.md#canonical-5a4aa5d4a120e0b2f9d37636a012e9dceb16fe61e82a35b77d152fdfa73a4977) |
| `labels` | [labels](resources--bgp--reference--group-001.md#canonical-eb6edf2a0a30d6c7331a458cf44f7df21204e147059efebbc4d236468571ae44) |
| `name` | [name](resources--bgp--reference--group-001.md#canonical-bb75c8b2866ef908404293ef81adf26231bbf6e268013773dfa014144e716de7) |
| `namespace` | [namespace](resources--bgp--reference--group-001.md#canonical-c8e8e6024e5d308e097ccc611278b3b0eb3404034eea68e254958387770f50f9) |
| `peers` | [peers](resources--bgp--reference--group-001.md#canonical-bd0847406a70497cbe8e3c5a35dadb25e46dd9c9a3445e0c1f40d7bb05285ae4) |
| `peers.bfd_disabled` | [peers.bfd_disabled](resources--bgp--reference--group-001.md#canonical-782499ec191d3b4ecf19e46296b7c4cf09d25be6e400d735180b17133e3d2ffd) |
| `peers.bfd_enabled` | [peers.bfd_enabled](resources--bgp--reference--group-001.md#canonical-0a30eb89eca83e6a826e570c53eb21d8c01fcd67ba6f64497f0cadd987e999bf) |
| `peers.bfd_enabled.multiplier` | [peers.bfd_enabled.multiplier](resources--bgp--reference--group-001.md#canonical-19992874207e5b398b0c30e6504cabdf5a699c6315658aa29a714b8223305a86) |
| `peers.bfd_enabled.receive_interval_milliseconds` | [peers.bfd_enabled.receive_interval_milliseconds](resources--bgp--reference--group-001.md#canonical-25c968a49a62d6cd36a47ba8211c5de20a78bf80eb05d195c161521e382b4226) |
| `peers.bfd_enabled.transmit_interval_milliseconds` | [peers.bfd_enabled.transmit_interval_milliseconds](resources--bgp--reference--group-001.md#canonical-0494500542d0cafcfd6b0904292e4ae67619b7d21df861023bcdefe6086955a6) |
| `peers.disable_spec` | [peers.disable_spec](resources--bgp--reference--group-001.md#canonical-79f746d35a2e2260501f560918f8af59cf20bf5b93d5ea4b9a7851602472ce31) |
| `peers.ebgp_multihop_disabled` | [peers.ebgp_multihop_disabled](resources--bgp--reference--group-001.md#canonical-3dff417b72e32f2cf74395786c49984b309fbef19a2b30ca46db17e652d53d25) |
| `peers.ebgp_multihop_enabled` | [peers.ebgp_multihop_enabled](resources--bgp--reference--group-001.md#canonical-74f601a65d17d93d0c783d4fc18d58e4f0ae371e1b30411e48fd8e2cfa876424) |
| `peers.external` | [peers.external](resources--bgp--reference--group-001.md#canonical-80960c20ffc2cf6696f62c345f5bac78dc548e0942b857ce44c8cdc7f27aa0e4) |
| `peers.external.address` | [peers.external.address](resources--bgp--reference--group-001.md#canonical-9eef0f07c2641e1d30bf04056f9a1a9efc38f671cbc35165710fd0144d50c5a5) |
| `peers.external.address_ipv6` | [peers.external.address_ipv6](resources--bgp--reference--group-001.md#canonical-b2cadc2b616fc98d5322308e6c30e1fa924b141282a41c8dc585e6bcbf0c87ab) |
| `peers.external.asn` | [peers.external.asn](resources--bgp--reference--group-001.md#canonical-67b051c3a6f7b7562f16cab0d31cd1c4d37455a225c90eefccd1d01140650d52) |
| `peers.external.default_gateway` | [peers.external.default_gateway](resources--bgp--reference--group-001.md#canonical-5735f9241bf6e59e72ad6939f21da39ead8d00b60f908e6e29bb10978a00df6c) |
| `peers.external.default_gateway_v6` | [peers.external.default_gateway_v6](resources--bgp--reference--group-001.md#canonical-07bb4a3c1ed6fb99fcb8b97dd5da14bcdd4e282eb76cbbf468574ce472522a88) |
| `peers.external.disable_spec` | [peers.external.disable_spec](resources--bgp--reference--group-001.md#canonical-186c5e7b11aba3cf5dddc42e4440c2a61a7ec764a56fcc8d0194466e11577594) |
| `peers.external.disable_v6` | [peers.external.disable_v6](resources--bgp--reference--group-001.md#canonical-720049b609606087e1b03346494e2232d0052d78c8a17c907b9c53f8bab5139f) |
| `peers.external.external_connector` | [peers.external.external_connector](resources--bgp--reference--group-001.md#canonical-39f01a4b72cfc57f7a265c3dcdd1b9f35f0809060fe70edd32e2e9350194a120) |
| `peers.external.family_inet` | [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-76dcd523905caef67f5925689ed944c4cc8eaa84041e12d2df959512b2fb54f3) |
| `peers.external.family_inet.disable_spec` | [peers.external.family_inet.disable_spec](resources--bgp--reference--group-001.md#canonical-766ad2144a093b3384de6e5da46b35392f19487d7ae9a1203615afec350109b2) |
| `peers.external.family_inet.enable` | [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-cf6e2c354b38afb9f85e6af8949101c6fd85cf83d9ab70ce1dc3db008bc19153) |
| `peers.external.family_inet.enable.aggregation` | [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-a9b90364ed860c3d0acf345da43963e1c2dc42e6658dbe29b54e60ce1ad62350) |
| `peers.external.family_inet.enable.aggregation.ip_prefix` | [peers.external.family_inet.enable.aggregation.ip_prefix](resources--bgp--reference--group-001.md#canonical-232bca7740fc672073b9b650450b904d4a5f93ebfadfc8320303deffa6ce1c3d) |
| `peers.external.family_inet.enable.aggregation.options` | [peers.external.family_inet.enable.aggregation.options](resources--bgp--reference--group-001.md#canonical-ba32da36fafeba658ca91a5051e1270ed9d203b7ab854d47cc995bd5e1a4e5f8) |
| `peers.external.family_inet.enable.aggregation.options.summary_only` | [peers.external.family_inet.enable.aggregation.options.summary_only](resources--bgp--reference--group-001.md#canonical-bb5602cf94e4f44280e2e13480390ff06683179f75d262907e5fb298bb177a03) |
| `peers.external.from_site` | [peers.external.from_site](resources--bgp--reference--group-001.md#canonical-fe12b9930c6c54ca1429d3d4e78765dedbbb4f9acbbbd2528813f1909a60d0ef) |
| `peers.external.from_site_v6` | [peers.external.from_site_v6](resources--bgp--reference--group-001.md#canonical-f102674908c5ccdaff0b2bca2b046b9b4b5a020008c215c28122113189b05c17) |
| `peers.external.interface` | [peers.external.interface](resources--bgp--reference--group-001.md#canonical-aae76cf84158e0c2e3753227dcab6f24c0a211a853b9c471dcd995282643610a) |
| `peers.external.interface.name` | [peers.external.interface.name](resources--bgp--reference--group-001.md#canonical-c4b526bf4631325c26deeeb67cba9bd1d6a65f0c8f14109fb94c52839b40e725) |
| `peers.external.interface.namespace` | [peers.external.interface.namespace](resources--bgp--reference--group-001.md#canonical-cfbae23f2698e2d83fd775c28f7529feed312e64baa46c6973ef703207c155b1) |
| `peers.external.interface.tenant` | [peers.external.interface.tenant](resources--bgp--reference--group-001.md#canonical-06433f66032b508600c00c9043f01998210f0c79ab077f7aec8c260c30f6c651) |
| `peers.external.interface_list` | [peers.external.interface_list](resources--bgp--reference--group-001.md#canonical-c42d003f1eb88514778c2e4a42907204b1fb3b6f8b419baabe522658d1a5f665) |
| `peers.external.interface_list.interfaces` | [peers.external.interface_list.interfaces](resources--bgp--reference--group-001.md#canonical-7ee06c60f5a17fcaec904c4517df799f771be8c18a37d42c61594f3e8e29fe3d) |
| `peers.external.interface_list.interfaces.name` | [peers.external.interface_list.interfaces.name](resources--bgp--reference--group-001.md#canonical-81b414a1f245d75e625992effec830274ce0734d5098f515e1ea1c5b37519dbd) |
| `peers.external.interface_list.interfaces.namespace` | [peers.external.interface_list.interfaces.namespace](resources--bgp--reference--group-001.md#canonical-1d2c174e5a7c5e2248f1729d0eac5d5e01d72c28fc060bacbebf0b93d2012ef0) |
| `peers.external.interface_list.interfaces.tenant` | [peers.external.interface_list.interfaces.tenant](resources--bgp--reference--group-001.md#canonical-b5a5d12df8c7ff52ee00841ae74697844083534b920b9d2006f23269447bbd56) |
| `peers.external.md5_auth_key` | [peers.external.md5_auth_key](resources--bgp--reference--group-001.md#canonical-0a923ffa0774c0002ee23886596ba4e45d3d1e53c1bc7c1afbc03fcc17575d98) |
| `peers.external.no_authentication` | [peers.external.no_authentication](resources--bgp--reference--group-001.md#canonical-fd74d1c01bea9e52c14e5cb1763085819cae5028850bf0d9906cac6d5ed84099) |
| `peers.external.port` | [peers.external.port](resources--bgp--reference--group-001.md#canonical-f659cfcc83450431e5a30f4cb15864c19db82aa14ace3bfdf7bb8985e826c2e8) |
| `peers.external.subnet_begin_offset` | [peers.external.subnet_begin_offset](resources--bgp--reference--group-001.md#canonical-5d2c1b0f9de01746aa51f1897d27d552ec5da048a4287e83b7d4cf42753e849f) |
| `peers.external.subnet_begin_offset_v6` | [peers.external.subnet_begin_offset_v6](resources--bgp--reference--group-001.md#canonical-c6448247d8274d2346f27b6a9b68aed0a490aa69581e277d12f9e5d0b974f1d2) |
| `peers.external.subnet_end_offset` | [peers.external.subnet_end_offset](resources--bgp--reference--group-001.md#canonical-083f0951844e22ef130e12c43c37cdf531790751422513f09c2bf9886924885a) |
| `peers.external.subnet_end_offset_v6` | [peers.external.subnet_end_offset_v6](resources--bgp--reference--group-001.md#canonical-ce2afcba806c7dda78e7da514bccbe5801d1737d3f7e27ffe9d41b412bef74ec) |
| `peers.label` | [peers.label](resources--bgp--reference--group-001.md#canonical-58fab825420856145462647128d426c05012d2c7e169bc8f68dfb79263d5d793) |
| `peers.metadata` | [peers.metadata](resources--bgp--reference--group-001.md#canonical-db25932d7966153cad9eb1b9ad4220526cf6f3f41877a75498c8565e775095f0) |
| `peers.metadata.description_spec` | [peers.metadata.description_spec](resources--bgp--reference--group-001.md#canonical-b9bb4ab58f8a45f9e5b280535a947da5e88102f8875410b8f1d8585cb24f047c) |
| `peers.metadata.name` | [peers.metadata.name](resources--bgp--reference--group-001.md#canonical-a1ca7e6f7df60d76a62943b1aff1f7f61968c3ac0a5a46261008833c8822a3fb) |
| `peers.passive_mode_disabled` | [peers.passive_mode_disabled](resources--bgp--reference--group-001.md#canonical-a622f2e6aaba2d9a88017cf989055c0df7c5afa6486ac0bd8ed2d5b426b5391a) |
| `peers.passive_mode_enabled` | [peers.passive_mode_enabled](resources--bgp--reference--group-001.md#canonical-65b755a4b28214a2ea9993b80fcbd833b40bf795bff24a3268be3c203d770181) |
| `peers.routing_policies` | [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-6a17463248e11b86469bbae9cc1f3ecc50a764dfc750e9d86a4ce19fb1038557) |
| `peers.routing_policies.route_policy` | [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-ca75b3634258ba1de04a148d6e5f580ce5cb3ddba4715bfbf34855890dde02b3) |
| `peers.routing_policies.route_policy.all_nodes` | [peers.routing_policies.route_policy.all_nodes](resources--bgp--reference--group-001.md#canonical-9262c28b3afa865833ac98b7f892add76f39f1a35b7dd0a14e757c4863103990) |
| `peers.routing_policies.route_policy.inbound` | [peers.routing_policies.route_policy.inbound](resources--bgp--reference--group-001.md#canonical-eec26d532d2f85a521590d08bfbd54abc6d249d174a4dfb9d0e259c779af820b) |
| `peers.routing_policies.route_policy.node_name` | [peers.routing_policies.route_policy.node_name](resources--bgp--reference--group-001.md#canonical-90ced0998da2d14ee20793236e2e3b79fe432fced1b4b7dd4c7a950b3a7ac9d9) |
| `peers.routing_policies.route_policy.node_name.node` | [peers.routing_policies.route_policy.node_name.node](resources--bgp--reference--group-001.md#canonical-86f42fd15fe8cd32d40989a16ff8b65e8182f339e9c4599125307dc03672ab95) |
| `peers.routing_policies.route_policy.object_refs` | [peers.routing_policies.route_policy.object_refs](resources--bgp--reference--group-001.md#canonical-2d2b69405f688b5bc69b52c1c82e1836714d861fa2abbaedf5f18da56544c74c) |
| `peers.routing_policies.route_policy.object_refs.kind` | [peers.routing_policies.route_policy.object_refs.kind](resources--bgp--reference--group-001.md#canonical-f3400eb833e33ddb272f851761f9429b8bc9a25f51c6b2af96bcd221fbfcbaaf) |
| `peers.routing_policies.route_policy.object_refs.name` | [peers.routing_policies.route_policy.object_refs.name](resources--bgp--reference--group-001.md#canonical-e7b1248ea2b85ec40da05f8f2832d3e6efe9f29a02c72ce4682fcefcb9b2e271) |
| `peers.routing_policies.route_policy.object_refs.namespace` | [peers.routing_policies.route_policy.object_refs.namespace](resources--bgp--reference--group-001.md#canonical-22ec6834516f8f5c057dc0b4e86308823744e3f87669f5e40930cd8b07438ac8) |
| `peers.routing_policies.route_policy.object_refs.tenant` | [peers.routing_policies.route_policy.object_refs.tenant](resources--bgp--reference--group-001.md#canonical-a0c18ee72b07aeb845e8d5b5822760090bb35baab239e573cfd982bd85abd1e0) |
| `peers.routing_policies.route_policy.object_refs.uid` | [peers.routing_policies.route_policy.object_refs.uid](resources--bgp--reference--group-001.md#canonical-d62cc791180ba59b1fa979fc6d63781ae0888c367d5b2b0eb7701391b8175c4a) |
| `peers.routing_policies.route_policy.outbound` | [peers.routing_policies.route_policy.outbound](resources--bgp--reference--group-001.md#canonical-5ed6c8ea63a04a848966dab022ae0486e2f206f22f2ca58992525562019a14b4) |
| `timeouts` | [timeouts](resources--bgp--reference--group-001.md#canonical-c97be875b057c4e28d91512abb2f0b92591c2069b2571e75709269a95f635267) |
| `timeouts.create` | [timeouts.create](resources--bgp--reference--group-001.md#canonical-016e750f5178b0913139fd09d46b5fcf8267f9238dd28f58429802cb1721ac87) |
| `timeouts.delete` | [timeouts.delete](resources--bgp--reference--group-001.md#canonical-73c1f9a6d70e49512770b28f9c47c40c5d7a63c8fcaea9edcd27e9f91b71cfc9) |
| `timeouts.read` | [timeouts.read](resources--bgp--reference--group-001.md#canonical-278a812ddea3e9d10e7d5c4ecb5ff28a757a2e3cb02d7447fabdf8282497ced7) |
| `timeouts.update` | [timeouts.update](resources--bgp--reference--group-001.md#canonical-65ab51973a1d79a0ce11d55b5cd1b616bec5b30e833e718c4d3aa9e8cdc0be1c) |
| `where` | [where](resources--bgp--reference--group-001.md#canonical-e84d409f7fe08bcff7e23b2848254557635d61705d99b61f0ccea6c10924238b) |
| `where.site` | [where.site](resources--bgp--reference--group-001.md#canonical-ec095faf8b1bea7f30585ec5c9e78519f13dcadcb80170853654f10f2c6fd832) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--bgp--reference--group-001.md#canonical-6b0f5c69f3452ed61c69a9dee73f72073e52d885351337e46fcfb455a7ca6216) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--bgp--reference--group-001.md#canonical-9c048eaf77b5c5d0a752ffb01e734f7e8c786cd0d598ac53642cde0569472438) |
| `where.site.network_type` | [where.site.network_type](resources--bgp--reference--group-001.md#canonical-5d8dbeacbb16c40951031161e287e0cd42da32293ce8faddc77fd908ba5a5497) |
| `where.site.ref` | [where.site.ref](resources--bgp--reference--group-001.md#canonical-4a76677c4616759f5f3a6d9e79e9c10a9f1b11b14389fcb96c9bf1fb675e14a7) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--bgp--reference--group-001.md#canonical-69de9891df396026339015bc8b2ed13c580a3fc07a1e560bb8af53266c52d4bf) |
| `where.site.ref.name` | [where.site.ref.name](resources--bgp--reference--group-001.md#canonical-c03df7694fd63e45ad60f12c426c583d8272536fe03db2ba923a020cf62c4711) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--bgp--reference--group-001.md#canonical-9db9781b72291855a4ce21299d4587b4e8af8231b2c567bf5576e8458eda3562) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--bgp--reference--group-001.md#canonical-19078964e5ad5c8ca4dc392dd2b61238c7ef9ff22f1c0c2e75513e7c69c5d5b0) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--bgp--reference--group-001.md#canonical-f7158375e28df115ad17692809849b1cfbcbdbeda30ebe8f22bf4e50347107ba) |
| `where.virtual_site` | [where.virtual_site](resources--bgp--reference--group-001.md#canonical-5f56f369d2d75506a18b022e3ae39d67d4b132392ca7cf9895147335d97ee051) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--bgp--reference--group-001.md#canonical-3077fd4e2adf942ae9b21122432a0d080a13620b3dcd17fddb9e4fe4e172cf3c) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--bgp--reference--group-001.md#canonical-c3ee5a652e697e87a8c6af9b59d77a064fce1d7f7c7112c5da7b6290a8b67e95) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--bgp--reference--group-001.md#canonical-168aeab8149d664fe4bdde8b541601a3c581024a04247681255205ee77da7da4) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--bgp--reference--group-001.md#canonical-1214283505415477b779106f3d6a5de19e621d7ad23d17ad9e48a2c76b48ceb2) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--bgp--reference--group-002.md#canonical-0e3ffbcc9dcade2a8373f5bc6391e871857e6d774b54154e48198a018a4cc02c) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--bgp--reference--group-002.md#canonical-947b96f34277781a2083f1597c04a27694a4657ec146b18ee43065633b34f60c) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--bgp--reference--group-002.md#canonical-cec6ab3c637feb75beddc9f3e4482918771e9d9009063ef6faa6c26a95491dc4) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--bgp--reference--group-002.md#canonical-1f60238c9fdf7e06f6bd0e6d1285729baaafaa4ca34095d1b0a9bed58d0e0820) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--bgp--reference--group-002.md#canonical-d1419ee4c9cb18225a3b9779ea78d35b5d49b31d8f1ec39d49513caad3eac5fd) |

<a id="canonical-e6ca8eb3e3a8d085ebb9db8c1871dfb07de58f76979374d2f71534d8560eda93"></a>

## Next pages — Property reference / a2ca6f03f464 / 12

- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-d30f89abfb3bcd2816487185f399e481a1d04f005852d935b50153fb1f63c465)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [timeouts](resources--bgp--reference--group-001.md#canonical-f09bbcefdb3d000d17aea209da85cf650c0c55df309b01717b7c4b34ae89314f)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-d30f89abfb3bcd2816487185f399e481a1d04f005852d935b50153fb1f63c465"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ebdcb5f6f1c6a67c5535043c32b7853bf7b2e7ac670b184b4e7b23de0cd293a"></a>

## bgp_parameters — bgp_parameters / 7fcf0bc032da / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- bgp_parameters

<a id="canonical-36ea4f1dee24c0f4814c1ae5d4da83d76afa25f4903c93ec3530e05ff5c5c31b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bgp parameters.

Upstream description:

BGP parameters for the local site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn"),
  validators.ConflictingObjectAttributes("from_site",
    "ip_address"),
  validators.ConflictingObjectAttributes("from_site",
    "local_address"),
  validators.ConflictingObjectAttributes("ip_address",
    "local_address")}
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
  "x-ves-oneof-field-router_id_choice": "[\"from_site\",\"ip_address\",\"local_address\"]"
}
```

Terraform syntax:

```terraform
bgp_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-120f54b5f2aef42c71ef1e01720cfc6a5b751b4bb58fcf2c3c70b0b3477e3e23"></a>

## Direct properties — bgp_parameters / 7fcf0bc032da / 3

<a id="canonical-a414488956bc214e737ed52f072786672b2e212c5b083598bd89db5caa4bdf9f"></a>

<a id="canonical-c7361fcba4edbf657860369304234e16010215ea36de41b8ea1e5e817bde1f74"></a>

## asn property — bgp_parameters / 7fcf0bc032da / 4

Type: `"number"`. Optional.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [from_site](resources--bgp--reference--group-001.md#canonical-63b4d7f9673247ca49d40e7cd5f96e1e5b46a38e96613f1c9333f2eca32d4ae5): complete subsection reference.

<a id="canonical-faa40a4955698771f81f6a3b2d7c2b991e5b7d691d5a83b0c295a6a812e25e55"></a>

<a id="canonical-9df22027a589cd01324e329e6e82d06795ba2e3a97bb7c97c1c7d050b0d03b1e"></a>

## ip_address property — bgp_parameters / 7fcf0bc032da / 5

Type: `"string"`. Optional.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Upstream description:

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

- [local_address](resources--bgp--reference--group-001.md#canonical-3a9bffb1d0854a84e4498f26b96ee050853dc310cecdda3c9b19ded6f9b18e78): complete subsection reference.

<a id="canonical-831cb4bc6d1bf75e6852639c841d1d6fe469813862fb2711b3c7d0bd0c26fc2a"></a>

## Next pages — bgp_parameters / 7fcf0bc032da / 6

- [bgp_parameters.from_site](resources--bgp--reference--group-001.md#canonical-63b4d7f9673247ca49d40e7cd5f96e1e5b46a38e96613f1c9333f2eca32d4ae5)
- [bgp_parameters.local_address](resources--bgp--reference--group-001.md#canonical-3a9bffb1d0854a84e4498f26b96ee050853dc310cecdda3c9b19ded6f9b18e78)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-63b4d7f9673247ca49d40e7cd5f96e1e5b46a38e96613f1c9333f2eca32d4ae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d729a765d36fedb702d587f35b27991004bb96657d59c16dbf133ab3e8807da"></a>

## bgp_parameters.from_site — bgp_parameters.from_site / 91153466b111 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-d30f89abfb3bcd2816487185f399e481a1d04f005852d935b50153fb1f63c465)
- bgp_parameters.from_site

<a id="canonical-978b1d05b50930adfb079fe1b14b23d3a708d96d2a037d58150ca78ba4c6d0ff"></a>

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
from_site = {}
```

<a id="canonical-dabb0cb2be1eebd552e3e44780aa590bc3efe331d7a8a320c65eb5da4891345b"></a>

## Direct properties — bgp_parameters.from_site / 91153466b111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11b904a670d3b1088195f1107bee205de122e4e393b5095243d0a8b803cfb8b9"></a>

## Next pages — bgp_parameters.from_site / 91153466b111 / 4

- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-d30f89abfb3bcd2816487185f399e481a1d04f005852d935b50153fb1f63c465)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-3a9bffb1d0854a84e4498f26b96ee050853dc310cecdda3c9b19ded6f9b18e78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72b74e0765bc91f36127a39e6f5200ab7225d9f5e3ca1b8380dfbdf60ce802e6"></a>

## bgp_parameters.local_address — bgp_parameters.local_address / 1a31b2a0c138 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-d30f89abfb3bcd2816487185f399e481a1d04f005852d935b50153fb1f63c465)
- bgp_parameters.local_address

<a id="canonical-32d767cb782abceb065613e530dae1a67ae4d88e5b93ed873f70d9d34ceba80e"></a>

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
local_address = {}
```

<a id="canonical-ec00fd748fdef96446dce348348befc4e56a0f73d08c75570044220d953114d6"></a>

## Direct properties — bgp_parameters.local_address / 1a31b2a0c138 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f1ee9d211792baf4370141a883946ce2f7b7bbd38d07ce579b6eb32f3ef8332"></a>

## Next pages — bgp_parameters.local_address / 1a31b2a0c138 / 4

- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-d30f89abfb3bcd2816487185f399e481a1d04f005852d935b50153fb1f63c465)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49deaf302e4536d50a127143f9ac65a1e02f6f0eda706f8755515aba5899c7cc"></a>

## peers — peers / 73419caabb8d / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- peers

<a id="canonical-bd0847406a70497cbe8e3c5a35dadb25e46dd9c9a3445e0c1f40d7bb05285ae4"></a>

Type: `"object"`. list nested block, Optional.

Peers. List of peers.

Upstream description:

List of peers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bfd_disabled",
    "bfd_enabled"),
  validators.ConflictingListObjectAttributes("disable_spec",
    "routing_policies"),
  validators.ConflictingListObjectAttributes("ebgp_multihop_disabled",
    "ebgp_multihop_enabled"),
  validators.ConflictingListObjectAttributes("passive_mode_disabled",
    "passive_mode_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
peers {
  # Configure direct properties listed below.
}
```

<a id="canonical-f492d36039696172cc1d29b5f7a24b3f0142f7aa2f6f4d6381bc294b00a8e1f1"></a>

## Direct properties — peers / 73419caabb8d / 3

- [bfd_disabled](resources--bgp--reference--group-001.md#canonical-876749d3b751137ccaf0a115e7081a00e7216a7f0b74acfeb8c5d3a979c8d934): complete subsection reference.

- [bfd_enabled](resources--bgp--reference--group-001.md#canonical-55792024d68e0c83fab75a1962dec885d4d0bddb89824ece282383726ee51219): complete subsection reference.

- [disable_spec](resources--bgp--reference--group-001.md#canonical-ed04fad1dcfdd41a884ce0cba5a2d45546fd25404e40013a5d7be7a4d474d2eb): complete subsection reference.

- [ebgp_multihop_disabled](resources--bgp--reference--group-001.md#canonical-17fb2a9f0433a69a5b3428b8896a647da038255dcb929ea94202b599a2256cd8): complete subsection reference.

- [ebgp_multihop_enabled](resources--bgp--reference--group-001.md#canonical-31cdb70dc13c2e3531f401e112c247b5948c581cfab23cc52190a9f24e39225a): complete subsection reference.

- [external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d): complete subsection reference.

<a id="canonical-58fab825420856145462647128d426c05012d2c7e169bc8f68dfb79263d5d793"></a>

<a id="canonical-3a01f9f8f61da9a524c244e14ac6767bd6b9b6545bd89bb5df598a91a91eb680"></a>

## label property — peers / 73419caabb8d / 4

Type: `"string"`. Optional.

Label. Specify whether this peer should be.

Upstream description:

Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--bgp--reference--group-001.md#canonical-52fb9e3af2530d09851ee01126edb5040e88ed74eb98c11791f84bcdca262c88): complete subsection reference.

- [passive_mode_disabled](resources--bgp--reference--group-001.md#canonical-c1a9a4695efa46c9ef4a837c1c55941d0acb8ded656aec28cfe495c2765cc781): complete subsection reference.

- [passive_mode_enabled](resources--bgp--reference--group-001.md#canonical-2e632c57ee5785748f675de1ccff92ae451e4e90b98917e9ff50ab60fcec0ccd): complete subsection reference.

- [routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237): complete subsection reference.

<a id="canonical-01ddde48a0f79cb4ed4d08e98ec3ae7192db679d735f26b7de39dc6e2127b506"></a>

## Next pages — peers / 73419caabb8d / 5

- [peers.bfd_disabled](resources--bgp--reference--group-001.md#canonical-876749d3b751137ccaf0a115e7081a00e7216a7f0b74acfeb8c5d3a979c8d934)
- [peers.bfd_enabled](resources--bgp--reference--group-001.md#canonical-55792024d68e0c83fab75a1962dec885d4d0bddb89824ece282383726ee51219)
- [peers.disable_spec](resources--bgp--reference--group-001.md#canonical-ed04fad1dcfdd41a884ce0cba5a2d45546fd25404e40013a5d7be7a4d474d2eb)
- [peers.ebgp_multihop_disabled](resources--bgp--reference--group-001.md#canonical-17fb2a9f0433a69a5b3428b8896a647da038255dcb929ea94202b599a2256cd8)
- [peers.ebgp_multihop_enabled](resources--bgp--reference--group-001.md#canonical-31cdb70dc13c2e3531f401e112c247b5948c581cfab23cc52190a9f24e39225a)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [peers.metadata](resources--bgp--reference--group-001.md#canonical-52fb9e3af2530d09851ee01126edb5040e88ed74eb98c11791f84bcdca262c88)
- [peers.passive_mode_disabled](resources--bgp--reference--group-001.md#canonical-c1a9a4695efa46c9ef4a837c1c55941d0acb8ded656aec28cfe495c2765cc781)
- [peers.passive_mode_enabled](resources--bgp--reference--group-001.md#canonical-2e632c57ee5785748f675de1ccff92ae451e4e90b98917e9ff50ab60fcec0ccd)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-876749d3b751137ccaf0a115e7081a00e7216a7f0b74acfeb8c5d3a979c8d934"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec96e28d419e9f2e0bb5280580ff462230986bab534aaa26cce488af467213e2"></a>

## peers.bfd_disabled — peers.bfd_disabled / 32011f1f4360 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.bfd_disabled

<a id="canonical-782499ec191d3b4ecf19e46296b7c4cf09d25be6e400d735180b17133e3d2ffd"></a>

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
bfd_disabled = {}
```

<a id="canonical-448a8fc9e4bf12562f7158869a2535ef58118c6f0d6ef2e58ba9c4a6567ce104"></a>

## Direct properties — peers.bfd_disabled / 32011f1f4360 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-401ae6926719f30e030519cbfcd289d3a5a7d3656ebd418c59c28cac321bfc67"></a>

## Next pages — peers.bfd_disabled / 32011f1f4360 / 4

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-55792024d68e0c83fab75a1962dec885d4d0bddb89824ece282383726ee51219"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a21899e3ea9adb1c2483ed716abe894b20c15d67aa0aa1885b95cd8001dc493"></a>

## peers.bfd_enabled — peers.bfd_enabled / 6bfc80db8bb4 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.bfd_enabled

<a id="canonical-0a30eb89eca83e6a826e570c53eb21d8c01fcd67ba6f64497f0cadd987e999bf"></a>

Type: `"object"`. single nested block, Optional.

BFD. BFD parameters.

Upstream description:

BFD parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("multiplier",
    "receive_interval_milliseconds",
    "transmit_interval_milliseconds")}
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
bfd_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-4cf91a158935e54cc2e1912456d631c4130d3c66629b9b43786a5daa5baeccf4"></a>

## Direct properties — peers.bfd_enabled / 6bfc80db8bb4 / 3

<a id="canonical-19992874207e5b398b0c30e6504cabdf5a699c6315658aa29a714b8223305a86"></a>

<a id="canonical-20427281dc0c656583f76ee4e10cb061931c18b2119e9cbc8a412576468d5fab"></a>

## multiplier property — peers.bfd_enabled / 6bfc80db8bb4 / 4

Type: `"number"`. Optional.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 255),
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
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-25c968a49a62d6cd36a47ba8211c5de20a78bf80eb05d195c161521e382b4226"></a>

<a id="canonical-13548599339a2b88d1334732db02c8e3f7f0943a513cd4f61f2dd2643b4fc8f1"></a>

## receive_interval_milliseconds property — peers.bfd_enabled / 6bfc80db8bb4 / 5

Type: `"number"`. Optional.

BFD receive interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0494500542d0cafcfd6b0904292e4ae67619b7d21df861023bcdefe6086955a6"></a>

<a id="canonical-008d9fd5163860aaa6e3e0b917d12324f0cf34b2a3a58ba61b7b4a11270ccd79"></a>

## transmit_interval_milliseconds property — peers.bfd_enabled / 6bfc80db8bb4 / 6

Type: `"number"`. Optional.

BFD transmit interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1abc415c5a3589b24d1f6eb89b156460397816a0661ba41217405d5051fab027"></a>

## Next pages — peers.bfd_enabled / 6bfc80db8bb4 / 7

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-ed04fad1dcfdd41a884ce0cba5a2d45546fd25404e40013a5d7be7a4d474d2eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a023f48eacfbd9afa39a20c3ffef7305a8c88332a1d129e22a15e4229d2113d0"></a>

## peers.disable_spec — peers.disable_spec / 747ad1d29c14 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.disable_spec

<a id="canonical-79f746d35a2e2260501f560918f8af59cf20bf5b93d5ea4b9a7851602472ce31"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-ee0253267c02b335158c6da0c59da1319dfe9cebc70a1738d96763f4ddfcca44"></a>

## Direct properties — peers.disable_spec / 747ad1d29c14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9dbda64272181145d65bf8cae1b0086126f12a29f56e1c38abc635168d51cb10"></a>

## Next pages — peers.disable_spec / 747ad1d29c14 / 4

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-17fb2a9f0433a69a5b3428b8896a647da038255dcb929ea94202b599a2256cd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89c9b21f6cf58cd53acc0da5486be30f59c5c331a988b2f0b78fe9cc4f7f79b7"></a>

## peers.ebgp_multihop_disabled — peers.ebgp_multihop_disabled / 9a3a63e36663 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.ebgp_multihop_disabled

<a id="canonical-3dff417b72e32f2cf74395786c49984b309fbef19a2b30ca46db17e652d53d25"></a>

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
ebgp_multihop_disabled = {}
```

<a id="canonical-15a411efb90dfbc3aa7511272433ca1b7ef5f3904d62329140a48ae53f2d590e"></a>

## Direct properties — peers.ebgp_multihop_disabled / 9a3a63e36663 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e89cf3af51562150e78ac160d7d7ca363e79a126fa106aa26d28ee8ead2061b"></a>

## Next pages — peers.ebgp_multihop_disabled / 9a3a63e36663 / 4

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-31cdb70dc13c2e3531f401e112c247b5948c581cfab23cc52190a9f24e39225a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5afa2c9e7681a0d4b0248b076d4e7d101e7aaeeb0c9677c00e7b8e8588ed256"></a>

## peers.ebgp_multihop_enabled — peers.ebgp_multihop_enabled / 456f763614a7 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.ebgp_multihop_enabled

<a id="canonical-74f601a65d17d93d0c783d4fc18d58e4f0ae371e1b30411e48fd8e2cfa876424"></a>

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
ebgp_multihop_enabled = {}
```

<a id="canonical-a1b1f06d964f015ef675581e18b4921300a195c131b57f7bc9eabf32895149de"></a>

## Direct properties — peers.ebgp_multihop_enabled / 456f763614a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7176e4cba06bfc6be6be89baf51158a32c720c0f69a459193f4c297f3f39b654"></a>

## Next pages — peers.ebgp_multihop_enabled / 456f763614a7 / 4

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06104b350c7b09de5d55db6c4297031f3e649ce21bb8cb8ec374f356d677d4b5"></a>

## peers.external — peers.external / c09315f81a6c / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.external

<a id="canonical-80960c20ffc2cf6696f62c345f5bac78dc548e0942b857ce44c8cdc7f27aa0e4"></a>

Type: `"object"`. single nested block, Optional.

External BGP Peer. External BGP Peer parameters.

Upstream description:

External BGP Peer parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn",
    "port"),
  validators.ConflictingObjectAttributes("address",
    "default_gateway"),
  validators.ConflictingObjectAttributes("address",
    "disable_spec"),
  validators.ConflictingObjectAttributes("address",
    "external_connector"),
  validators.ConflictingObjectAttributes("address",
    "from_site"),
  validators.ConflictingObjectAttributes("address",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("address",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "default_gateway_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_gateway",
    "external_connector"),
  validators.ConflictingObjectAttributes("default_gateway",
    "from_site"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("disable_spec",
    "external_connector"),
  validators.ConflictingObjectAttributes("disable_spec",
    "from_site"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("disable_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("external_connector",
    "from_site"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("interface",
    "interface_list"),
  validators.ConflictingObjectAttributes("md5_auth_key",
    "no_authentication"),
  validators.ConflictingObjectAttributes("subnet_begin_offset",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("subnet_begin_offset_v6",
    "subnet_end_offset_v6")}
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
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

Terraform syntax:

```terraform
external {
  # Configure direct properties listed below.
}
```

<a id="canonical-6b03af143741be426d67029d79c0a11c35add57a13c683ac752f3746f35a958a"></a>

## Direct properties — peers.external / c09315f81a6c / 3

<a id="canonical-9eef0f07c2641e1d30bf04056f9a1a9efc38f671cbc35165710fd0144d50c5a5"></a>

<a id="canonical-26372e2b65e87d123dae7663680d0ff5181a5336e2ff36d1a509e3faa1f3958f"></a>

## address property — peers.external / c09315f81a6c / 4

Type: `"string"`. Optional.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Upstream description:

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

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

<a id="canonical-b2cadc2b616fc98d5322308e6c30e1fa924b141282a41c8dc585e6bcbf0c87ab"></a>

<a id="canonical-95e0af31c764970c749455984566a513dca1b7373099cf147cac8ba556865c8a"></a>

## address_ipv6 property — peers.external / c09315f81a6c / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Upstream description:

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

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

<a id="canonical-67b051c3a6f7b7562f16cab0d31cd1c4d37455a225c90eefccd1d01140650d52"></a>

<a id="canonical-e5c345bd73adc028cc0d35389c31905240568ab1a1f7451542fd4c2fe4c6551a"></a>

## asn property — peers.external / c09315f81a6c / 6

Type: `"number"`. Optional.

ASN. Autonomous System Number for BGP peer.

Upstream description:

Autonomous System Number for BGP peer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](resources--bgp--reference--group-001.md#canonical-a646f5fc0dd73e913662fbc465e40f8410a709c32712d25f342cd0b504d61e63): complete subsection reference.

- [default_gateway_v6](resources--bgp--reference--group-001.md#canonical-b0f35b8b45c6585728635d4195faec20aeb242f5235fcf0a6ffd2cb70c1f7a1d): complete subsection reference.

- [disable_spec](resources--bgp--reference--group-001.md#canonical-452850a3a85857a1e767da0ce1c98d01bda2bb7b2a6176469de2f56ebcfdb196): complete subsection reference.

- [disable_v6](resources--bgp--reference--group-001.md#canonical-07d74e96c103df0914494917f1d06561253fa6c6128e4df2f9027c8ce1af9541): complete subsection reference.

- [external_connector](resources--bgp--reference--group-001.md#canonical-cf6b5c386f3628f9c02d77b5e0736035d940a68a9344484c3804695e142c9e68): complete subsection reference.

- [family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723): complete subsection reference.

- [from_site](resources--bgp--reference--group-001.md#canonical-1ecb569671cc8878688ef58287367d62058abff3418187699d5ebb0320a65caa): complete subsection reference.

- [from_site_v6](resources--bgp--reference--group-001.md#canonical-1d91b8da6411f7c0b9edb7ac3a926eaa540ead94045b2cad5aa9f9a4d0348afe): complete subsection reference.

- [interface](resources--bgp--reference--group-001.md#canonical-eb06dea74455712226ef475ea1308bdc0fc5d05efb6651ec7f380e69996445f9): complete subsection reference.

- [interface_list](resources--bgp--reference--group-001.md#canonical-2903ae4c579626921f7623ed7d51905b15a3fc16dad5e1da393cde86f4ca0741): complete subsection reference.

<a id="canonical-0a923ffa0774c0002ee23886596ba4e45d3d1e53c1bc7c1afbc03fcc17575d98"></a>

<a id="canonical-d3f24c72f1281af421d01c2a026aad2362aed2bf4591e35da7633c34c3098486"></a>

## md5_auth_key property — peers.external / c09315f81a6c / 7

Type: `"string"`. Optional.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Upstream description:

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385)

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

- [no_authentication](resources--bgp--reference--group-001.md#canonical-3cac5b771dcbb41fae494da03d3dfcce71418891b846b57ad003ac2b0e980d12): complete subsection reference.

<a id="canonical-f659cfcc83450431e5a30f4cb15864c19db82aa14ace3bfdf7bb8985e826c2e8"></a>

<a id="canonical-97486c19c0e26628eb6e7616b497ca1f9ae4599ccb30999e2e5a31c98259ae12"></a>

## port property — peers.external / c09315f81a6c / 8

Type: `"number"`. Optional.

Peer Port. Peer TCP port number.

Upstream description:

Peer TCP port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-5d2c1b0f9de01746aa51f1897d27d552ec5da048a4287e83b7d4cf42753e849f"></a>

<a id="canonical-581b8927b9ba20e6d3ec7e33e05265aab0f8ce219d88f78b400eb78759c662ad"></a>

## subnet_begin_offset property — peers.external / c09315f81a6c / 9

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-c6448247d8274d2346f27b6a9b68aed0a490aa69581e277d12f9e5d0b974f1d2"></a>

<a id="canonical-5ec875434c7db6f88fa2a6935cf4598071528001dd6726921236440c5d67a891"></a>

## subnet_begin_offset_v6 property — peers.external / c09315f81a6c / 10

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-083f0951844e22ef130e12c43c37cdf531790751422513f09c2bf9886924885a"></a>

<a id="canonical-ab7e2cf47b9ecec87fdcb74d45403615beffa9572328f7fae7ea5b4a02dd6850"></a>

## subnet_end_offset property — peers.external / c09315f81a6c / 11

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-ce2afcba806c7dda78e7da514bccbe5801d1737d3f7e27ffe9d41b412bef74ec"></a>

<a id="canonical-f6ce467152aa94c8f5f621e80f00e948f550d712431cbfdbac75ab58cba63f5c"></a>

## subnet_end_offset_v6 property — peers.external / c09315f81a6c / 12

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-334ea779e3f2d125549900d8e8aac2f35e74667c363e115dce131bf74618eba1"></a>

## Next pages — peers.external / c09315f81a6c / 13

- [peers.external.default_gateway](resources--bgp--reference--group-001.md#canonical-a646f5fc0dd73e913662fbc465e40f8410a709c32712d25f342cd0b504d61e63)
- [peers.external.default_gateway_v6](resources--bgp--reference--group-001.md#canonical-b0f35b8b45c6585728635d4195faec20aeb242f5235fcf0a6ffd2cb70c1f7a1d)
- [peers.external.disable_spec](resources--bgp--reference--group-001.md#canonical-452850a3a85857a1e767da0ce1c98d01bda2bb7b2a6176469de2f56ebcfdb196)
- [peers.external.disable_v6](resources--bgp--reference--group-001.md#canonical-07d74e96c103df0914494917f1d06561253fa6c6128e4df2f9027c8ce1af9541)
- [peers.external.external_connector](resources--bgp--reference--group-001.md#canonical-cf6b5c386f3628f9c02d77b5e0736035d940a68a9344484c3804695e142c9e68)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- [peers.external.from_site](resources--bgp--reference--group-001.md#canonical-1ecb569671cc8878688ef58287367d62058abff3418187699d5ebb0320a65caa)
- [peers.external.from_site_v6](resources--bgp--reference--group-001.md#canonical-1d91b8da6411f7c0b9edb7ac3a926eaa540ead94045b2cad5aa9f9a4d0348afe)
- [peers.external.interface](resources--bgp--reference--group-001.md#canonical-eb06dea74455712226ef475ea1308bdc0fc5d05efb6651ec7f380e69996445f9)
- [peers.external.interface_list](resources--bgp--reference--group-001.md#canonical-2903ae4c579626921f7623ed7d51905b15a3fc16dad5e1da393cde86f4ca0741)
- [peers.external.no_authentication](resources--bgp--reference--group-001.md#canonical-3cac5b771dcbb41fae494da03d3dfcce71418891b846b57ad003ac2b0e980d12)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-a646f5fc0dd73e913662fbc465e40f8410a709c32712d25f342cd0b504d61e63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f4bf35c06344cd40cd03262750b312258c1f740b3672dff737a05ecbe66191b"></a>

## peers.external.default_gateway — peers.external.default_gateway / cacc4a75fe82 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.default_gateway

<a id="canonical-5735f9241bf6e59e72ad6939f21da39ead8d00b60f908e6e29bb10978a00df6c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-3c19e081080c74a1757e34e0d84221c4dbf3ab9ddbb0df9a7690f86e4f79be5f"></a>

## Direct properties — peers.external.default_gateway / cacc4a75fe82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85591f6f14c62029021eaad47cfcce7fbe1a3fce06fd33f8dd19ac1f21a73653"></a>

## Next pages — peers.external.default_gateway / cacc4a75fe82 / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-b0f35b8b45c6585728635d4195faec20aeb242f5235fcf0a6ffd2cb70c1f7a1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3299a485ef1428a9b3094945613c4c47fc748a69163fa896c0bd85f359801728"></a>

## peers.external.default_gateway_v6 — peers.external.default_gateway_v6 / 08cd7bcc6dfc / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.default_gateway_v6

<a id="canonical-07bb4a3c1ed6fb99fcb8b97dd5da14bcdd4e282eb76cbbf468574ce472522a88"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway v6.

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
default_gateway_v6 = {}
```

<a id="canonical-48b727b805ba7cf6bcc8761886c95deffd27e02caa84ee1dba0dce28a332b4b2"></a>

## Direct properties — peers.external.default_gateway_v6 / 08cd7bcc6dfc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ee0712fffe9550c7dc7ed8cbb1abed67dc4a633445e2d68e5e6d152c58384a9b"></a>

## Next pages — peers.external.default_gateway_v6 / 08cd7bcc6dfc / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-452850a3a85857a1e767da0ce1c98d01bda2bb7b2a6176469de2f56ebcfdb196"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2518e089dae712c4cc52078e20327f051cfddde984fd42f9b9c26921e29cd71b"></a>

## peers.external.disable_spec — peers.external.disable_spec / a7c95a1c5daf / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.disable_spec

<a id="canonical-186c5e7b11aba3cf5dddc42e4440c2a61a7ec764a56fcc8d0194466e11577594"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-b180997d819ea23455e4c86a86882ce0be98832a1aab0d8afaaf47fc73778e0c"></a>

## Direct properties — peers.external.disable_spec / a7c95a1c5daf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa46cd3b7cf4d41ece8f5795f41729d3d7de5c445b085aac3e2bbd998548c410"></a>

## Next pages — peers.external.disable_spec / a7c95a1c5daf / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-07d74e96c103df0914494917f1d06561253fa6c6128e4df2f9027c8ce1af9541"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d96fa531a463c46f321bc9366e705ff97cb778ab3f5420987396f6d38d1f041e"></a>

## peers.external.disable_v6 — peers.external.disable_v6 / 85e50dd33cc6 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.disable_v6

<a id="canonical-720049b609606087e1b03346494e2232d0052d78c8a17c907b9c53f8bab5139f"></a>

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
disable_v6 = {}
```

<a id="canonical-2ed230164b3f6c3208a16f79d4762eb25bd638673d0d80a4d0af1f3d39222a75"></a>

## Direct properties — peers.external.disable_v6 / 85e50dd33cc6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1af0e0b22896046362ca047f5cac0b4c4ba058d362759f5ede5bbe571fe4ef46"></a>

## Next pages — peers.external.disable_v6 / 85e50dd33cc6 / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-cf6b5c386f3628f9c02d77b5e0736035d940a68a9344484c3804695e142c9e68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c68a3359a878ac315a3221e4d441817165c97af1f9bed17fc78270a58a1fe47"></a>

## peers.external.external_connector — peers.external.external_connector / 4da720193f9b / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.external_connector

<a id="canonical-39f01a4b72cfc57f7a265c3dcdd1b9f35f0809060fe70edd32e2e9350194a120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for external connector.

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
external_connector = {}
```

<a id="canonical-8e0011290431e71033f8e97523363e282d284990269fc6c18e43db977fffdb1f"></a>

## Direct properties — peers.external.external_connector / 4da720193f9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-81aab1021f472102bd8e8f7cdb7684049dda207699f889538be3780c01893fcf"></a>

## Next pages — peers.external.external_connector / 4da720193f9b / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cdbf70b9b0a35b38ad640ed54bb12277873be58c077437a2c4eb98d1bad4f43"></a>

## peers.external.family_inet — peers.external.family_inet / 46ff26154d37 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.family_inet

<a id="canonical-76dcd523905caef67f5925689ed944c4cc8eaa84041e12d2df959512b2fb54f3"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
family_inet {
  # Configure direct properties listed below.
}
```

<a id="canonical-46948acc5aedd39befa641fc5227745a6341e75713b99c378dc988567a4f653a"></a>

## Direct properties — peers.external.family_inet / 46ff26154d37 / 3

- [disable_spec](resources--bgp--reference--group-001.md#canonical-e1e0ab6938bd5c350cbab53cceb0fc531df7a95f6a6b46b895250dc89a70a1ca): complete subsection reference.

- [enable](resources--bgp--reference--group-001.md#canonical-091d6ab92d05531d903ad477b582104d2b7e2233bd1d9b3e7c92fca870c658f1): complete subsection reference.

<a id="canonical-d0fd322ff6a8876c57e645081f7bee62cd4ceadd7aa3f6794124397fb544d23f"></a>

## Next pages — peers.external.family_inet / 46ff26154d37 / 4

- [peers.external.family_inet.disable_spec](resources--bgp--reference--group-001.md#canonical-e1e0ab6938bd5c350cbab53cceb0fc531df7a95f6a6b46b895250dc89a70a1ca)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-091d6ab92d05531d903ad477b582104d2b7e2233bd1d9b3e7c92fca870c658f1)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-e1e0ab6938bd5c350cbab53cceb0fc531df7a95f6a6b46b895250dc89a70a1ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b42e3efafb5afb7330e5153cdbcaab7899624dead6c0b527b8b1b373adb20e5"></a>

## peers.external.family_inet.disable_spec — peers.external.family_inet.disable_spec / e3c068f24224 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- peers.external.family_inet.disable_spec

<a id="canonical-766ad2144a093b3384de6e5da46b35392f19487d7ae9a1203615afec350109b2"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-bde401d5fe11a5c71f76148c59f77c17ec1e34315ccc492e3d2f438e82ceaa95"></a>

## Direct properties — peers.external.family_inet.disable_spec / e3c068f24224 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51960d9ca1946339ad584a0305e1d05fb66275d0f024e57dfcf12f979d5d5b0e"></a>

## Next pages — peers.external.family_inet.disable_spec / e3c068f24224 / 4

- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-091d6ab92d05531d903ad477b582104d2b7e2233bd1d9b3e7c92fca870c658f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b1781456a3ff5edbba7f8ec755f50a58f823576c9d378e3bb7ef2dc0fdbb23d"></a>

## peers.external.family_inet.enable — peers.external.family_inet.enable / 3746048d7137 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- peers.external.family_inet.enable

<a id="canonical-cf6e2c354b38afb9f85e6af8949101c6fd85cf83d9ab70ce1dc3db008bc19153"></a>

Type: `"object"`. single nested block, Optional.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

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
enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-464f704af3fd598d4a000f7d0c5282ee7fa83d22a16811019c5a58d8d976eb19"></a>

## Direct properties — peers.external.family_inet.enable / 3746048d7137 / 3

- [aggregation](resources--bgp--reference--group-001.md#canonical-d995918f335683e8599728563615dc2c0fff0f4c8974343bccd382f88167ffb7): complete subsection reference.

<a id="canonical-922840bd7f235c88a75ebadd9851481482f4ce4f590aefd8dd61819990e4c4c1"></a>

## Next pages — peers.external.family_inet.enable / 3746048d7137 / 4

- [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-d995918f335683e8599728563615dc2c0fff0f4c8974343bccd382f88167ffb7)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-d995918f335683e8599728563615dc2c0fff0f4c8974343bccd382f88167ffb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a23a9d6f48a58219b03340e76009b4d545be259606d540ed46423d36a728813"></a>

## peers.external.family_inet.enable.aggregation — peers.external.family_inet.enable.aggregation / d895a7bc0f42 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-091d6ab92d05531d903ad477b582104d2b7e2233bd1d9b3e7c92fca870c658f1)
- peers.external.family_inet.enable.aggregation

<a id="canonical-a9b90364ed860c3d0acf345da43963e1c2dc42e6658dbe29b54e60ce1ad62350"></a>

Type: `"object"`. list nested block, Optional.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

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

Terraform syntax:

```terraform
aggregation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0fa43c30122b03acb99fbb7e79ebd418412952b49cc2e3822b16f494b0455ca7"></a>

## Direct properties — peers.external.family_inet.enable.aggregation / d895a7bc0f42 / 3

<a id="canonical-232bca7740fc672073b9b650450b904d4a5f93ebfadfc8320303deffa6ce1c3d"></a>

<a id="canonical-7462634d2064efa87fe620c7aaa9ba1f16880461f341386f33ca176491fddc48"></a>

## ip_prefix property — peers.external.family_inet.enable.aggregation / d895a7bc0f42 / 4

Type: `"string"`. Optional.

IP Prefix. Specify IPv4 subnet for aggregation.

Upstream description:

Specify IPv4 subnet for aggregation.

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

- [options](resources--bgp--reference--group-001.md#canonical-e8a42f27a8a35ab8ae3ae44945160c49a4ef723afa024e61f13c15a87c069aef): complete subsection reference.

<a id="canonical-830980c80af9dcecb32e6e68e81cad28f2fdd578627c151bf189187e5d791773"></a>

## Next pages — peers.external.family_inet.enable.aggregation / d895a7bc0f42 / 5

- [peers.external.family_inet.enable.aggregation.options](resources--bgp--reference--group-001.md#canonical-e8a42f27a8a35ab8ae3ae44945160c49a4ef723afa024e61f13c15a87c069aef)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-091d6ab92d05531d903ad477b582104d2b7e2233bd1d9b3e7c92fca870c658f1)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-e8a42f27a8a35ab8ae3ae44945160c49a4ef723afa024e61f13c15a87c069aef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b94f12caf5bde23334f9ce155976785a566d3e014665531b8e0d3f5eec2f3020"></a>

## peers.external.family_inet.enable.aggregation.options — peers.external.family_inet.enable.aggregation.options / 2f010f3e5fab / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-091d6ab92d05531d903ad477b582104d2b7e2233bd1d9b3e7c92fca870c658f1)
- [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-d995918f335683e8599728563615dc2c0fff0f4c8974343bccd382f88167ffb7)
- peers.external.family_inet.enable.aggregation.options

<a id="canonical-ba32da36fafeba658ca91a5051e1270ed9d203b7ab854d47cc995bd5e1a4e5f8"></a>

Type: `"object"`. list nested block, Optional.

Aggregation OPTIONS. Configuration parameter for options

Upstream description:

Configuration parameter for options

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

Terraform syntax:

```terraform
options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3afe14467a298551956fdd52baf818a9c85dd713c75eb37328f993e7933739a4"></a>

## Direct properties — peers.external.family_inet.enable.aggregation.options / 2f010f3e5fab / 3

- [summary_only](resources--bgp--reference--group-001.md#canonical-9dc70108ff6fcaf4998fb3af6fe33aae835f8805a1e16c468267349c5dba50e0): complete subsection reference.

<a id="canonical-abddc1708decef12fd4a6f14ce9ae1f3091bdd455e1e438ebcbb4fb682605181"></a>

## Next pages — peers.external.family_inet.enable.aggregation.options / 2f010f3e5fab / 4

- [peers.external.family_inet.enable.aggregation.options.summary_only](resources--bgp--reference--group-001.md#canonical-9dc70108ff6fcaf4998fb3af6fe33aae835f8805a1e16c468267349c5dba50e0)
- [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-d995918f335683e8599728563615dc2c0fff0f4c8974343bccd382f88167ffb7)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-9dc70108ff6fcaf4998fb3af6fe33aae835f8805a1e16c468267349c5dba50e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be9a98f5f4158bd9826dd332dcb5a718efc597757e5e04c37a7cd3bd4f3eadc4"></a>

## peers.external.family_inet.enable.aggregation.options.summary_only — peers.external.family_inet.enable.aggregation.options.summary_only / dfdaf74270eb / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-17990b54071d42aee51e0f36a42f7682fd212499ef8c8bd09f74c8c345283723)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-091d6ab92d05531d903ad477b582104d2b7e2233bd1d9b3e7c92fca870c658f1)
- [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-d995918f335683e8599728563615dc2c0fff0f4c8974343bccd382f88167ffb7)
- [peers.external.family_inet.enable.aggregation.options](resources--bgp--reference--group-001.md#canonical-e8a42f27a8a35ab8ae3ae44945160c49a4ef723afa024e61f13c15a87c069aef)
- peers.external.family_inet.enable.aggregation.options.summary_only

<a id="canonical-bb5602cf94e4f44280e2e13480390ff06683179f75d262907e5fb298bb177a03"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for summary only.

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
summary_only {}
```

<a id="canonical-fda86d6246d16320e0f8c3d2a76f7289f6a694543d54e969662248764313320e"></a>

## Direct properties — peers.external.family_inet.enable.aggregation.options.summary_only / dfdaf74270eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6d4b6cf3009571702a7c1e697db3d3999641145ff5f17132cfdd325c05f02d4"></a>

## Next pages — peers.external.family_inet.enable.aggregation.options.summary_only / dfdaf74270eb / 4

- [peers.external.family_inet.enable.aggregation.options](resources--bgp--reference--group-001.md#canonical-e8a42f27a8a35ab8ae3ae44945160c49a4ef723afa024e61f13c15a87c069aef)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-1ecb569671cc8878688ef58287367d62058abff3418187699d5ebb0320a65caa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e0ba7280979da20dc1803da66c5be39dae28dc01c870a09709193b0e2a7df13"></a>

## peers.external.from_site — peers.external.from_site / 63c8ddf33eca / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.from_site

<a id="canonical-fe12b9930c6c54ca1429d3d4e78765dedbbb4f9acbbbd2528813f1909a60d0ef"></a>

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
from_site = {}
```

<a id="canonical-31c72f35f9231bccb3f2b29efad6a06fd06349479e3c80da33e6a1831ad1776d"></a>

## Direct properties — peers.external.from_site / 63c8ddf33eca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ddaf649ddf9ae840cca2e9574e84b6ce422815397662492c17a95c080c5180e9"></a>

## Next pages — peers.external.from_site / 63c8ddf33eca / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-1d91b8da6411f7c0b9edb7ac3a926eaa540ead94045b2cad5aa9f9a4d0348afe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be90ef42f13ab50e579dde94cb121ece5b4d68f245e41592e712cd98fb83d105"></a>

## peers.external.from_site_v6 — peers.external.from_site_v6 / 732bffcc1346 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.from_site_v6

<a id="canonical-f102674908c5ccdaff0b2bca2b046b9b4b5a020008c215c28122113189b05c17"></a>

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
from_site_v6 = {}
```

<a id="canonical-3a1b56832aa00bc72c88e8b6a17319cbcdb69b415ccbaee8c40b009e0068941e"></a>

## Direct properties — peers.external.from_site_v6 / 732bffcc1346 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-801e301b90bb092a3b4cf4128a740e84501328ffd0da44dca50b51c4f6336850"></a>

## Next pages — peers.external.from_site_v6 / 732bffcc1346 / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-eb06dea74455712226ef475ea1308bdc0fc5d05efb6651ec7f380e69996445f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d93326f2f0fbe981a1b218711682140a7378c311b4e1839fd87409d0de71ff61"></a>

## peers.external.interface — peers.external.interface / 5fddaa0c81c8 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.interface

<a id="canonical-aae76cf84158e0c2e3753227dcab6f24c0a211a853b9c471dcd995282643610a"></a>

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-89261f05c0e032d3fcd8e784cf737b4ec71d0e1f34f37e107281b2f3c02a2f39"></a>

## Direct properties — peers.external.interface / 5fddaa0c81c8 / 3

<a id="canonical-c4b526bf4631325c26deeeb67cba9bd1d6a65f0c8f14109fb94c52839b40e725"></a>

<a id="canonical-35ca253e6f3333d3cdebab67a5033342b1a52043876c63929099f21367451701"></a>

## name property — peers.external.interface / 5fddaa0c81c8 / 4

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

<a id="canonical-cfbae23f2698e2d83fd775c28f7529feed312e64baa46c6973ef703207c155b1"></a>

<a id="canonical-7f21c45022d620ed54f2d02759027858fb343483faf5f4a79dc27b5df0af97a3"></a>

## namespace property — peers.external.interface / 5fddaa0c81c8 / 5

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

<a id="canonical-06433f66032b508600c00c9043f01998210f0c79ab077f7aec8c260c30f6c651"></a>

<a id="canonical-15d03456011ac44bdad9e1ade256c23224b14a425fcb3d468232f1bc4c0702eb"></a>

## tenant property — peers.external.interface / 5fddaa0c81c8 / 6

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

<a id="canonical-6bd720f44bd63a6413a29b7c12a5849038ae0c7902de2d2abf94cd506918bd05"></a>

## Next pages — peers.external.interface / 5fddaa0c81c8 / 7

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-2903ae4c579626921f7623ed7d51905b15a3fc16dad5e1da393cde86f4ca0741"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b723ba99c2d948f6fa240cce447f325951fe4ba77987f97f5159f7caf9ca404d"></a>

## peers.external.interface_list — peers.external.interface_list / f246874ce3fb / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.interface_list

<a id="canonical-c42d003f1eb88514778c2e4a42907204b1fb3b6f8b419baabe522658d1a5f665"></a>

Type: `"object"`. single nested block, Optional.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-a096571a64f0e8cfc4f84982797265d0aada399921d6cf0f627dffa1135337b5"></a>

## Direct properties — peers.external.interface_list / f246874ce3fb / 3

- [interfaces](resources--bgp--reference--group-001.md#canonical-23a885f7ff70a402c4041f7ebfb79bd58f94eca58d4dee18a51c58adb6495400): complete subsection reference.

<a id="canonical-ba32a12838c2e0242c5b1cd0f24efbdfab8070f1d69a5a816223e5b94f173199"></a>

## Next pages — peers.external.interface_list / f246874ce3fb / 4

- [peers.external.interface_list.interfaces](resources--bgp--reference--group-001.md#canonical-23a885f7ff70a402c4041f7ebfb79bd58f94eca58d4dee18a51c58adb6495400)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-23a885f7ff70a402c4041f7ebfb79bd58f94eca58d4dee18a51c58adb6495400"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8054476af2fc0516cd4115288f3d0a0d2ad20b8215cc47c9e29bad4161c13a32"></a>

## peers.external.interface_list.interfaces — peers.external.interface_list.interfaces / 41c68ee30370 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [peers.external.interface_list](resources--bgp--reference--group-001.md#canonical-2903ae4c579626921f7623ed7d51905b15a3fc16dad5e1da393cde86f4ca0741)
- peers.external.interface_list.interfaces

<a id="canonical-7ee06c60f5a17fcaec904c4517df799f771be8c18a37d42c61594f3e8e29fe3d"></a>

Type: `"object"`. list nested block, Optional.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5e6aef76f35124c49766a17bf1d909f8616a408a0d4f14769611a0992d1373e"></a>

## Direct properties — peers.external.interface_list.interfaces / 41c68ee30370 / 3

<a id="canonical-81b414a1f245d75e625992effec830274ce0734d5098f515e1ea1c5b37519dbd"></a>

<a id="canonical-c73b1aa52c550cfae499ef223a5cd39538e384e3a16a423fb6e19a27462fbf34"></a>

## name property — peers.external.interface_list.interfaces / 41c68ee30370 / 4

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

<a id="canonical-1d2c174e5a7c5e2248f1729d0eac5d5e01d72c28fc060bacbebf0b93d2012ef0"></a>

<a id="canonical-d9391ae539622d68967c9a0f63e2d55ad75206a095715a0f18e1fed4d462ce7c"></a>

## namespace property — peers.external.interface_list.interfaces / 41c68ee30370 / 5

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

<a id="canonical-b5a5d12df8c7ff52ee00841ae74697844083534b920b9d2006f23269447bbd56"></a>

<a id="canonical-866b1fc60b6d9f91872d5832e9281accebc07ec8598c82f173fb2ae9b1130c41"></a>

## tenant property — peers.external.interface_list.interfaces / 41c68ee30370 / 6

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

<a id="canonical-b09dd236eb71ad98dfa05a22704b746f0625980252cce57fc43fa20351704a4c"></a>

## Next pages — peers.external.interface_list.interfaces / 41c68ee30370 / 7

- [peers.external.interface_list](resources--bgp--reference--group-001.md#canonical-2903ae4c579626921f7623ed7d51905b15a3fc16dad5e1da393cde86f4ca0741)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-3cac5b771dcbb41fae494da03d3dfcce71418891b846b57ad003ac2b0e980d12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-540cb2e4203cd09062d358c89f9593752a801e0ff181f7675b36bc65f0035dd2"></a>

## peers.external.no_authentication — peers.external.no_authentication / 209b1d30cbb3 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- peers.external.no_authentication

<a id="canonical-fd74d1c01bea9e52c14e5cb1763085819cae5028850bf0d9906cac6d5ed84099"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authentication.

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
no_authentication = {}
```

<a id="canonical-c67d4c01a6e015581d17409438c98bcda2cfdef8a1537f1099d4143187792e5d"></a>

## Direct properties — peers.external.no_authentication / 209b1d30cbb3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-743b495811cadd10d33bcaf4d4aba59ecc19b54c4e53a9628547d4ca7f21e723"></a>

## Next pages — peers.external.no_authentication / 209b1d30cbb3 / 4

- [peers.external](resources--bgp--reference--group-001.md#canonical-0180206faa4bc2d9ecaf83351c94870bec7d85bdaf798fa86439ddd5dc29c13d)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-52fb9e3af2530d09851ee01126edb5040e88ed74eb98c11791f84bcdca262c88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e17c3a78e95b1721cf20e9de648de64114b34820642e5d982bc7383d8004010"></a>

## peers.metadata — peers.metadata / 5e2d887c5210 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.metadata

<a id="canonical-db25932d7966153cad9eb1b9ad4220526cf6f3f41877a75498c8565e775095f0"></a>

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

<a id="canonical-86d8d06305bc50c08f7028ffb907527a7e0496fa330379537e5c87e17a073696"></a>

## Direct properties — peers.metadata / 5e2d887c5210 / 3

<a id="canonical-b9bb4ab58f8a45f9e5b280535a947da5e88102f8875410b8f1d8585cb24f047c"></a>

<a id="canonical-04984409e8c7780f90be1cba9a01a8ea3c06db8021d69fc19b565d8eaff7ef59"></a>

## description_spec property — peers.metadata / 5e2d887c5210 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-a1ca7e6f7df60d76a62943b1aff1f7f61968c3ac0a5a46261008833c8822a3fb"></a>

<a id="canonical-48bdd0dffc7979535210703ccc5fd86da7314317f824ecc7e380d8c73415f5d5"></a>

## name property — peers.metadata / 5e2d887c5210 / 5

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

<a id="canonical-15458e0d96ef85652ba12d2a20a2e032d58bd7843e7cea00b691df2495bdfe48"></a>

## Next pages — peers.metadata / 5e2d887c5210 / 6

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-c1a9a4695efa46c9ef4a837c1c55941d0acb8ded656aec28cfe495c2765cc781"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8029998f2d878470e2d844d805bee53da9cd0fcc54b5f03456cb2275becbd483"></a>

## peers.passive_mode_disabled — peers.passive_mode_disabled / fceda8716dc4 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.passive_mode_disabled

<a id="canonical-a622f2e6aaba2d9a88017cf989055c0df7c5afa6486ac0bd8ed2d5b426b5391a"></a>

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
passive_mode_disabled = {}
```

<a id="canonical-d44104e6b8dec65e5e15644caab5d9a00a3cf8fc19618308498b2a43b9b7f99f"></a>

## Direct properties — peers.passive_mode_disabled / fceda8716dc4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5d0a946b1b470ed9fa1d7ad45e8854271bf10e1ad22e7a00c03d1a006de1a47"></a>

## Next pages — peers.passive_mode_disabled / fceda8716dc4 / 4

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-2e632c57ee5785748f675de1ccff92ae451e4e90b98917e9ff50ab60fcec0ccd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96c0bd0a3b6ac3b97ab31249cd7f01538524d75307a2d42e3e6fe627d7a0c37e"></a>

## peers.passive_mode_enabled — peers.passive_mode_enabled / ae1c25f570cb / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.passive_mode_enabled

<a id="canonical-65b755a4b28214a2ea9993b80fcbd833b40bf795bff24a3268be3c203d770181"></a>

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
passive_mode_enabled = {}
```

<a id="canonical-abf19c8c14740c16540a71471f45e4d509b5487e0d593214a8fb4a3e524517c1"></a>

## Direct properties — peers.passive_mode_enabled / ae1c25f570cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e729226867e69e614a3a7e11c005a3173056064bcf72efad58a659dd36beae56"></a>

## Next pages — peers.passive_mode_enabled / ae1c25f570cb / 4

- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a88f723c014fac16dac77a1714878a48a177897b0ad7417422904c88382ad76"></a>

## peers.routing_policies — peers.routing_policies / 0bd1157c33ea / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- peers.routing_policies

<a id="canonical-6a17463248e11b86469bbae9cc1f3ecc50a764dfc750e9d86a4ce19fb1038557"></a>

Type: `"object"`. single nested block, Optional.

List of rules which can be applied on all or particular nodes.

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
routing_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-86c899eb9df1cd3405266e1e27cf9c2d9cd731166412162461fe4d3b61abe2a2"></a>

## Direct properties — peers.routing_policies / 0bd1157c33ea / 3

- [route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76): complete subsection reference.

<a id="canonical-253f85d111e82bd6df86cd172e076874dbcf772b5f2d6e114a6dacd3e0d1c119"></a>

## Next pages — peers.routing_policies / 0bd1157c33ea / 4

- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3f1b9f180f253dc4fd8e2ed6bdf1bb16a20ba5a1f3a0ef5ac2fbe68724c84bd"></a>

## peers.routing_policies.route_policy — peers.routing_policies.route_policy / 610dcdddf1fd / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- peers.routing_policies.route_policy

<a id="canonical-ca75b3634258ba1de04a148d6e5f580ce5cb3ddba4715bfbf34855890dde02b3"></a>

Type: `"object"`. list nested block, Optional.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("object_refs"),
  validators.ConflictingListObjectAttributes("all_nodes",
    "node_name"),
  validators.ConflictingListObjectAttributes("inbound",
    "outbound")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
route_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6c5e53acacc7755b2acff4a89a86b0e5a837a64fa9d442cdb683085cfabaf28"></a>

## Direct properties — peers.routing_policies.route_policy / 610dcdddf1fd / 3

- [all_nodes](resources--bgp--reference--group-001.md#canonical-e09b7e2805711fcfbcbf8c2e1bf096fbe4a41cca07f73478c4ad93b6f33d104c): complete subsection reference.

- [inbound](resources--bgp--reference--group-001.md#canonical-3d59fc25285f20f9033218d4fdeaf29bb70c3adbfdda5606a70a328782e9903f): complete subsection reference.

- [node_name](resources--bgp--reference--group-001.md#canonical-2b2f186f07ca3fb790da2505fabaecc04ae601bc7be4aa161a3a843f373e3fe8): complete subsection reference.

- [object_refs](resources--bgp--reference--group-001.md#canonical-cc94b72ad637f0eebfcc9eb69042018ec2d022bd600fb359b6241a3d50c28bdb): complete subsection reference.

- [outbound](resources--bgp--reference--group-001.md#canonical-e9ffb27c3b0021664b3a87141a364cdbe60257453cc3f9fa1c641a915267a079): complete subsection reference.

<a id="canonical-ffa3484522b13a4730c06155716e57bfbc922b82037927134d2a2e54de91d350"></a>

## Next pages — peers.routing_policies.route_policy / 610dcdddf1fd / 4

- [peers.routing_policies.route_policy.all_nodes](resources--bgp--reference--group-001.md#canonical-e09b7e2805711fcfbcbf8c2e1bf096fbe4a41cca07f73478c4ad93b6f33d104c)
- [peers.routing_policies.route_policy.inbound](resources--bgp--reference--group-001.md#canonical-3d59fc25285f20f9033218d4fdeaf29bb70c3adbfdda5606a70a328782e9903f)
- [peers.routing_policies.route_policy.node_name](resources--bgp--reference--group-001.md#canonical-2b2f186f07ca3fb790da2505fabaecc04ae601bc7be4aa161a3a843f373e3fe8)
- [peers.routing_policies.route_policy.object_refs](resources--bgp--reference--group-001.md#canonical-cc94b72ad637f0eebfcc9eb69042018ec2d022bd600fb359b6241a3d50c28bdb)
- [peers.routing_policies.route_policy.outbound](resources--bgp--reference--group-001.md#canonical-e9ffb27c3b0021664b3a87141a364cdbe60257453cc3f9fa1c641a915267a079)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-e09b7e2805711fcfbcbf8c2e1bf096fbe4a41cca07f73478c4ad93b6f33d104c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12bf04853fabfa4ca2e2bd7c265878569b8863afe8ed4b7a46d56fb401b3ab6e"></a>

## peers.routing_policies.route_policy.all_nodes — peers.routing_policies.route_policy.all_nodes / b54c434c7a68 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- peers.routing_policies.route_policy.all_nodes

<a id="canonical-9262c28b3afa865833ac98b7f892add76f39f1a35b7dd0a14e757c4863103990"></a>

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
all_nodes = {}
```

<a id="canonical-84e55040bf7567c91f263a800839d7c68852e61780ba1f7eef078322d6d200f8"></a>

## Direct properties — peers.routing_policies.route_policy.all_nodes / b54c434c7a68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7dc1ec2b6ebee16fece4a0859495a28c2d0ea02b4fb4fa19ab72d540c63e706d"></a>

## Next pages — peers.routing_policies.route_policy.all_nodes / b54c434c7a68 / 4

- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-3d59fc25285f20f9033218d4fdeaf29bb70c3adbfdda5606a70a328782e9903f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa72acfe46c83cdf41b2b12720dc832a0cb137d95751a9150fbc4872d75dab27"></a>

## peers.routing_policies.route_policy.inbound — peers.routing_policies.route_policy.inbound / b24c017a5824 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- peers.routing_policies.route_policy.inbound

<a id="canonical-eec26d532d2f85a521590d08bfbd54abc6d249d174a4dfb9d0e259c779af820b"></a>

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
inbound = {}
```

<a id="canonical-48374e419caf837f3d507a0176cc9a8456a45212b9f7ea3b9cf1fc4a68b5ac9d"></a>

## Direct properties — peers.routing_policies.route_policy.inbound / b24c017a5824 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a44574417905d9b66f01a3df7af401b516dd2bbef293bf8f42309a056b60b0f2"></a>

## Next pages — peers.routing_policies.route_policy.inbound / b24c017a5824 / 4

- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-2b2f186f07ca3fb790da2505fabaecc04ae601bc7be4aa161a3a843f373e3fe8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f598a8427bde37180db773c5fb70101c2b5b22b91e66bda0a2f540c12cc7599"></a>

## peers.routing_policies.route_policy.node_name — peers.routing_policies.route_policy.node_name / 31bb8275684d / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- peers.routing_policies.route_policy.node_name

<a id="canonical-90ced0998da2d14ee20793236e2e3b79fe432fced1b4b7dd4c7a950b3a7ac9d9"></a>

Type: `"object"`. single nested block, Optional.

List of nodes on which BGP routing policy has to be applied.

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
node_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-86263fb1e706fa7bc432d5783d4e10292ce27a187d1f83cca499ebfc56522aca"></a>

## Direct properties — peers.routing_policies.route_policy.node_name / 31bb8275684d / 3

<a id="canonical-86f42fd15fe8cd32d40989a16ff8b65e8182f339e9c4599125307dc03672ab95"></a>

<a id="canonical-c5524cc10a3f57b5e57de15527d46a8ebab53d57d4bdbcf364d28f8fef1b82e2"></a>

## node property — peers.routing_policies.route_policy.node_name / 31bb8275684d / 4

Type: `["list", "string"]`. Optional.

Select BGP Session on which policy will be applied.

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

<a id="canonical-46d732b31dcc87e55efd86d2813a1d3b0c05034e11476cbefb55a19667bf4f57"></a>

## Next pages — peers.routing_policies.route_policy.node_name / 31bb8275684d / 5

- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-cc94b72ad637f0eebfcc9eb69042018ec2d022bd600fb359b6241a3d50c28bdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63e1f37ef1b7de23739145fac44513e6cdd87178fd50aaf36e5c568f985c889f"></a>

## peers.routing_policies.route_policy.object_refs — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- peers.routing_policies.route_policy.object_refs

<a id="canonical-2d2b69405f688b5bc69b52c1c82e1836714d861fa2abbaedf5f18da56544c74c"></a>

Type: `"object"`. list nested block, Optional.

BGP routing policy. Select route policy to apply.

Upstream description:

Select route policy to apply.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
object_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-1627f7e2516e23f894383d70cf75496614a2699036390c53fff5cc2642e47207"></a>

## Direct properties — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 3

<a id="canonical-f3400eb833e33ddb272f851761f9429b8bc9a25f51c6b2af96bcd221fbfcbaaf"></a>

<a id="canonical-e1fd57f9a262c9685cb200046d4244d4e9b0cbed09df772687ce14077f13e50e"></a>

## kind property — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 4

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

<a id="canonical-e7b1248ea2b85ec40da05f8f2832d3e6efe9f29a02c72ce4682fcefcb9b2e271"></a>

<a id="canonical-f31917b98d6e639977f61e61803e1b93f990a381d5a5614f6a1905b288d4a15c"></a>

## name property — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 5

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

<a id="canonical-22ec6834516f8f5c057dc0b4e86308823744e3f87669f5e40930cd8b07438ac8"></a>

<a id="canonical-e0d6c08b032fe9ba38cbc65a92dd7f821a7b936ef28285441c39b57d6718e2c7"></a>

## namespace property — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 6

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

<a id="canonical-a0c18ee72b07aeb845e8d5b5822760090bb35baab239e573cfd982bd85abd1e0"></a>

<a id="canonical-1ffe1a06ba3c0ce4bf13ebac4c41697e9462aa87b7597cdac92b41a318b6ff9b"></a>

## tenant property — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 7

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

<a id="canonical-d62cc791180ba59b1fa979fc6d63781ae0888c367d5b2b0eb7701391b8175c4a"></a>

<a id="canonical-ebfd1365732162eb4b276ec40bf7ac910b34e3c3be0ce8bc2d9fefbc5af884d0"></a>

## uid property — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 8

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

<a id="canonical-8197f6f0487b3ba4cd8713eba3d6af1c0cc301082853070d3a3149d1a52d86b2"></a>

## Next pages — peers.routing_policies.route_policy.object_refs / 8d99b9a22db4 / 9

- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-e9ffb27c3b0021664b3a87141a364cdbe60257453cc3f9fa1c641a915267a079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc69b42869e78a58a43093d8264b75a79086225bb5c33e9bbc1df74bd8f0d65a"></a>

## peers.routing_policies.route_policy.outbound — peers.routing_policies.route_policy.outbound / d31e0dd82cf5 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [peers](resources--bgp--reference--group-001.md#canonical-de0b7ceae426a94f87042f2e1b6567af391a200416720136a14355517801e213)
- [peers.routing_policies](resources--bgp--reference--group-001.md#canonical-2a70b2436216c882f96148c75dc8cde2ccd68b000f8b1380779b5657f827c237)
- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- peers.routing_policies.route_policy.outbound

<a id="canonical-5ed6c8ea63a04a848966dab022ae0486e2f206f22f2ca58992525562019a14b4"></a>

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
outbound = {}
```

<a id="canonical-78e0b1ce8a74891f5a1f76884a54cd0fd1a1e38762d5212cfa3c44405fadc63a"></a>

## Direct properties — peers.routing_policies.route_policy.outbound / d31e0dd82cf5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb3e5f638556e4272a5d501f52d871c032ae536397122b5230dee5c7341f9036"></a>

## Next pages — peers.routing_policies.route_policy.outbound / d31e0dd82cf5 / 4

- [peers.routing_policies.route_policy](resources--bgp--reference--group-001.md#canonical-1a33bcd706123209f579564d8b039790cb950d3419f08f6eb8032eaea031dd76)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-f09bbcefdb3d000d17aea209da85cf650c0c55df309b01717b7c4b34ae89314f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32f639e3b05516f939d7e63d51f777bdb3fa865b9e85b07cd982c2fbbca47176"></a>

## timeouts — timeouts / 5a690096f965 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- timeouts

<a id="canonical-c97be875b057c4e28d91512abb2f0b92591c2069b2571e75709269a95f635267"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-b41840e26ddabc02f23b85ed9b10ef906de9aa4e3e884ccf514cd8ab46598480"></a>

## Direct properties — timeouts / 5a690096f965 / 3

<a id="canonical-016e750f5178b0913139fd09d46b5fcf8267f9238dd28f58429802cb1721ac87"></a>

<a id="canonical-d4a0757867b98676dacfc7001205464b10e15a4b2d95a3ef69f297fb524f2448"></a>

## create property — timeouts / 5a690096f965 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-73c1f9a6d70e49512770b28f9c47c40c5d7a63c8fcaea9edcd27e9f91b71cfc9"></a>

<a id="canonical-d4a288ba16b2b77704e3b26f4a77b8c2045d6dcc0859dc478d843764999f2332"></a>

## delete property — timeouts / 5a690096f965 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-278a812ddea3e9d10e7d5c4ecb5ff28a757a2e3cb02d7447fabdf8282497ced7"></a>

<a id="canonical-5ef3e3793e91a405d51a4cfd38276b0e33f33eae7e025a2695b897484a34dd49"></a>

## read property — timeouts / 5a690096f965 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-65ab51973a1d79a0ce11d55b5cd1b616bec5b30e833e718c4d3aa9e8cdc0be1c"></a>

<a id="canonical-d095fcc4b10811c4a2f8211d951951db3cd42aba90ab0334922b8a5dcddc8fb5"></a>

## update property — timeouts / 5a690096f965 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f7173ed64092ae5b8623c89a2bfe899d8fa74aa62c88037b43ea728d43d281a2"></a>

## Next pages — timeouts / 5a690096f965 / 8

- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbd9c4c42af8bc58678632f358f46e0e19937176e4994d6bc94d6c3e6eefea1a"></a>

## where — where / 1477344f0ecd / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- where

<a id="canonical-e84d409f7fe08bcff7e23b2848254557635d61705d99b61f0ccea6c10924238b"></a>

Type: `"object"`. single nested block, Optional.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Upstream description:

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-75e465698ddd49df448c2cd25e01f0cb4a90747088cd0c27fdb6bea31316ef72"></a>

## Direct properties — where / 1477344f0ecd / 3

- [site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909): complete subsection reference.

- [virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0): complete subsection reference.

<a id="canonical-f23f0e3b8a0a448733096264b4655b56d8d2c7863124b8c3c760b8c181dd7dc3"></a>

## Next pages — where / 1477344f0ecd / 4

- [where.site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909)
- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-369fcf2003a0bc5d39a96883ae7a87f02735333b9d90acf067ff0cc485421016"></a>

## where.site — where.site / 976f93d84ba4 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- where.site

<a id="canonical-ec095faf8b1bea7f30585ec5c9e78519f13dcadcb80170853654f10f2c6fd832"></a>

Type: `"object"`. single nested block, Optional.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea77c82c8bb78d8983f284de517df0ca639b180f036990b526b7d04a3fc1141e"></a>

## Direct properties — where.site / 976f93d84ba4 / 3

- [disable_internet_vip](resources--bgp--reference--group-001.md#canonical-8460a9112de7a8dd3b478bf7cd0039228e25d492e827202bc8c47cfef0c0e4d2): complete subsection reference.

- [enable_internet_vip](resources--bgp--reference--group-001.md#canonical-d7cf1b748d93a1b7e586abe5bb799e18eef9903bc63f04335d974e86700be643): complete subsection reference.

<a id="canonical-5d8dbeacbb16c40951031161e287e0cd42da32293ce8faddc77fd908ba5a5497"></a>

<a id="canonical-5241a122b37b89dd802d88601ba721788576108c8de084e9ef43242ff9ff7241"></a>

## network_type property — where.site / 976f93d84ba4 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

- [ref](resources--bgp--reference--group-001.md#canonical-1a7054091b06c5b365e4b8a10708f53254b6771674433377ab278d03768413a9): complete subsection reference.

<a id="canonical-db4e3adec7bc323e98cf7b90e5d2f95f1d8381589eefe7ff205b4bdca5a46aea"></a>

## Next pages — where.site / 976f93d84ba4 / 5

- [where.site.disable_internet_vip](resources--bgp--reference--group-001.md#canonical-8460a9112de7a8dd3b478bf7cd0039228e25d492e827202bc8c47cfef0c0e4d2)
- [where.site.enable_internet_vip](resources--bgp--reference--group-001.md#canonical-d7cf1b748d93a1b7e586abe5bb799e18eef9903bc63f04335d974e86700be643)
- [where.site.ref](resources--bgp--reference--group-001.md#canonical-1a7054091b06c5b365e4b8a10708f53254b6771674433377ab278d03768413a9)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-8460a9112de7a8dd3b478bf7cd0039228e25d492e827202bc8c47cfef0c0e4d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c9a2487c64fc1722ad15a00e42ab80de6ed70aabf0f1644dd07e01bc8a65649"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / 293e305267a4 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [where.site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909)
- where.site.disable_internet_vip

<a id="canonical-6b0f5c69f3452ed61c69a9dee73f72073e52d885351337e46fcfb455a7ca6216"></a>

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
disable_internet_vip = {}
```

<a id="canonical-8e782c8eb684ab830680aa2c54f8add4b1bd9fd8480eb1cba009c1af8b089eec"></a>

## Direct properties — where.site.disable_internet_vip / 293e305267a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12ce99c1423ad77f04eda72cbc42e45e42e2e71e322d742a6a73502ababfea07"></a>

## Next pages — where.site.disable_internet_vip / 293e305267a4 / 4

- [where.site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-d7cf1b748d93a1b7e586abe5bb799e18eef9903bc63f04335d974e86700be643"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d49118966f8de91b9d4ee0d15762ff8fc2911e1750893f7fc55b8050f693fdf"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / 180b8fd663c3 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [where.site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909)
- where.site.enable_internet_vip

<a id="canonical-9c048eaf77b5c5d0a752ffb01e734f7e8c786cd0d598ac53642cde0569472438"></a>

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
enable_internet_vip = {}
```

<a id="canonical-b4f13eef0172406f460408c70f370e64eee2f0971f9982b608244d5eaaa06ca8"></a>

## Direct properties — where.site.enable_internet_vip / 180b8fd663c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90f08730c139926f56699036787425515d970729e7be4307e2273df005f1eba4"></a>

## Next pages — where.site.enable_internet_vip / 180b8fd663c3 / 4

- [where.site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-1a7054091b06c5b365e4b8a10708f53254b6771674433377ab278d03768413a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b97ac06a438edf1054a76577529bf7db5c586efb63ea9072b2c707f49cd8a35"></a>

## where.site.ref — where.site.ref / 1d86af477471 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [where.site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909)
- where.site.ref

<a id="canonical-4a76677c4616759f5f3a6d9e79e9c10a9f1b11b14389fcb96c9bf1fb675e14a7"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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

<a id="canonical-66b111ee30e7f1c89b46ee0f658e68a333e00f2849236a82ce19b9f49c1f8179"></a>

## Direct properties — where.site.ref / 1d86af477471 / 3

<a id="canonical-69de9891df396026339015bc8b2ed13c580a3fc07a1e560bb8af53266c52d4bf"></a>

<a id="canonical-7d9e71984cf065bf2196ce39e37e13ca1020dbaf76023bc54ceb56ee4eeb87d5"></a>

## kind property — where.site.ref / 1d86af477471 / 4

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

<a id="canonical-c03df7694fd63e45ad60f12c426c583d8272536fe03db2ba923a020cf62c4711"></a>

<a id="canonical-086bea79fee6946aaea379846d95f2d2b9d7f0348912e60f849c92abbd8df3ea"></a>

## name property — where.site.ref / 1d86af477471 / 5

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

<a id="canonical-9db9781b72291855a4ce21299d4587b4e8af8231b2c567bf5576e8458eda3562"></a>

<a id="canonical-9eb4caa6af720863166bafaafafaeec8f3bec376a60431ad063e78afa15611d8"></a>

## namespace property — where.site.ref / 1d86af477471 / 6

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

<a id="canonical-19078964e5ad5c8ca4dc392dd2b61238c7ef9ff22f1c0c2e75513e7c69c5d5b0"></a>

<a id="canonical-5c15f0a20c1b84db012b2ba2070c3866557eba8ffee2545b6bf66169eb59ca5b"></a>

## tenant property — where.site.ref / 1d86af477471 / 7

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

<a id="canonical-f7158375e28df115ad17692809849b1cfbcbdbeda30ebe8f22bf4e50347107ba"></a>

<a id="canonical-953f6f7584e00b17a9f4375370fbd0a5cf048de66a6e8c6c3875353205c0a23d"></a>

## uid property — where.site.ref / 1d86af477471 / 8

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

<a id="canonical-f311e058a18009f90975bc44b0582a6f1e15e8d1ecd5d137e873041d3c0dc8cb"></a>

## Next pages — where.site.ref / 1d86af477471 / 9

- [where.site](resources--bgp--reference--group-001.md#canonical-c34e26c59f4a8d14b773b856c8a8d862ec965ebf78910087c5c083b3f023e909)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fd9b4d043dd818682005069e48af835dea4aa3198f3ba22dc5f0d331d5f5cfd"></a>

## where.virtual_site — where.virtual_site / 6921a1f7c2e2 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- where.virtual_site

<a id="canonical-5f56f369d2d75506a18b022e3ae39d67d4b132392ca7cf9895147335d97ee051"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-298003a927b20154070660f7f2510cb3eb9b1433d787c36a003d5e8562154416"></a>

## Direct properties — where.virtual_site / 6921a1f7c2e2 / 3

- [disable_internet_vip](resources--bgp--reference--group-001.md#canonical-81dbef89654de46c659d68543f2e9e88fde44a0ca528795c285ab8fa380cdd89): complete subsection reference.

- [enable_internet_vip](resources--bgp--reference--group-001.md#canonical-d683d73af230f214d2a939338b74549cd4d0292d709a5b4b3cf61e53945127db): complete subsection reference.

<a id="canonical-168aeab8149d664fe4bdde8b541601a3c581024a04247681255205ee77da7da4"></a>

<a id="canonical-7a134628bac1b9819c66ae244261e242c7fe2a8304a9f5223d50c2355928fad6"></a>

## network_type property — where.virtual_site / 6921a1f7c2e2 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

- [ref](resources--bgp--reference--group-001.md#canonical-94909ec24effc463241234e2357c02b3300c684e22b3b1427adeb09b9159d4f5): complete subsection reference.

<a id="canonical-050fed638c1be2318dc89fb3899bade1d41241ff70c5f0f0b447ef23c8fa024e"></a>

## Next pages — where.virtual_site / 6921a1f7c2e2 / 5

- [where.virtual_site.disable_internet_vip](resources--bgp--reference--group-001.md#canonical-81dbef89654de46c659d68543f2e9e88fde44a0ca528795c285ab8fa380cdd89)
- [where.virtual_site.enable_internet_vip](resources--bgp--reference--group-001.md#canonical-d683d73af230f214d2a939338b74549cd4d0292d709a5b4b3cf61e53945127db)
- [where.virtual_site.ref](resources--bgp--reference--group-001.md#canonical-94909ec24effc463241234e2357c02b3300c684e22b3b1427adeb09b9159d4f5)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-81dbef89654de46c659d68543f2e9e88fde44a0ca528795c285ab8fa380cdd89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46bf0775ff7ce1d69add6455d9e3fe6432622ac6a19ef16299d30f4be239b9a7"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / fe71df0a5969 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0)
- where.virtual_site.disable_internet_vip

<a id="canonical-3077fd4e2adf942ae9b21122432a0d080a13620b3dcd17fddb9e4fe4e172cf3c"></a>

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
disable_internet_vip = {}
```

<a id="canonical-56d57d1b704e0c54df5b87598ea90829377b0d01866bc87e66bda9c26ad53a75"></a>

## Direct properties — where.virtual_site.disable_internet_vip / fe71df0a5969 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be26e7062884f8e5e96241d1fe6bf962efd05e038073d49ea101219e72a16210"></a>

## Next pages — where.virtual_site.disable_internet_vip / fe71df0a5969 / 4

- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-d683d73af230f214d2a939338b74549cd4d0292d709a5b4b3cf61e53945127db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3513262ab3c6b0a446ce220718b9c11ba86a45a7a034215a2d6e8400668b65e5"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / 7bc5c1bf3b55 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0)
- where.virtual_site.enable_internet_vip

<a id="canonical-c3ee5a652e697e87a8c6af9b59d77a064fce1d7f7c7112c5da7b6290a8b67e95"></a>

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
enable_internet_vip = {}
```

<a id="canonical-4becd317ff8932e47a64036fe950046b7db82d4aca900297ac7877950e32e92d"></a>

## Direct properties — where.virtual_site.enable_internet_vip / 7bc5c1bf3b55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d3899c9b16e312ca33386cdf4dd767a16faee65259edc88da2c86e0985eacbf"></a>

## Next pages — where.virtual_site.enable_internet_vip / 7bc5c1bf3b55 / 4

- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-94909ec24effc463241234e2357c02b3300c684e22b3b1427adeb09b9159d4f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7db91af78586816186300c110436f0bd6ece6c374e546f150aaf1c49d25c3177"></a>

## where.virtual_site.ref — where.virtual_site.ref / cc6a2a8af6e1 / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Property reference](resources--bgp--reference--group-001.md#canonical-2a2953839e3572dd9ac64a3b58149f926e162c49c7d745107372b9db7b315ccc)
- [where](resources--bgp--reference--group-001.md#canonical-f45f5139a5be24bac1b963bd3a3898f40d8fd7d799a0392d3301177e6f1c7bed)
- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0)
- where.virtual_site.ref

<a id="canonical-1214283505415477b779106f3d6a5de19e621d7ad23d17ad9e48a2c76b48ceb2"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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
