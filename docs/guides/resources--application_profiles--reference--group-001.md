---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f06f0de5fdb7a41e131aefdd251de8881e759726a75f415ee57ac3f8ae6972ae"></a>

## Property reference — Property reference / 9580f5c2d307 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- Property reference

<a id="canonical-28a0dae10c3340ec3cc453a33ade6b55e20add436492d7fe9a06cb18d0d3a050"></a>

## Direct properties — Property reference / 9580f5c2d307 / 3

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-1219c9680f1cb17608da1bff14b3321209c387abae26d7ec1cfd56a03083cefa): complete subsection reference.

<a id="canonical-c3e31bff8bfa15fddccc24b1dccfe2698b073de424d50fdca8147b192f1be3a3"></a>

<a id="canonical-f8850b07eb8d5b9fa7d781ce873a094d4668073f46fc9cfea58b9302fed54a09"></a>

## annotations property — Property reference / 9580f5c2d307 / 4

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

- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-a2f8dadbb1c84e46454b50aa5033edcdf5f94245ad7c84259a87b214799bc630): complete subsection reference.

<a id="canonical-4aec4599e2d39de21dfafca0cb857b0827611f7bb2f83b052b50759e53c34099"></a>

<a id="canonical-70f9276b568b25e88e7cdb7bd9fdb93cd090c850c5564e88c84a039cf4652281"></a>

## description property — Property reference / 9580f5c2d307 / 5

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

<a id="canonical-4258643eb76a4199c29b4e7f8edb3bcace2685a9fa946685166d1eeff1ce62ab"></a>

<a id="canonical-91ad4edb3a059eece64c9a84a6d895a7cb93dee5e0b1c22c693369e0fd63576b"></a>

## disable property — Property reference / 9580f5c2d307 / 6

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

<a id="canonical-5054148ade0a5f07062cc5915334bad222f6d6f5a06600106ad0c40c9c98643e"></a>

<a id="canonical-3df5b45690bd8f89890dc2ed13fb60b02c2272c01d04710ee1be5a96cb8d13a7"></a>

## id property — Property reference / 9580f5c2d307 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--application_profiles--reference--group-001.md#canonical-3f967a8395804c6c158083e0546c4f390cfd910da7685ac8449c86b36c6d564f): complete subsection reference.

<a id="canonical-e80b43352d103cef5a019abda317d6be15e5c4120b6add542c5ac92a428b569d"></a>

<a id="canonical-81fb5c274e8a5de62907fdc70f206469d3726555dd134d6476938e5e35741e2d"></a>

## labels property — Property reference / 9580f5c2d307 / 8

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

<a id="canonical-79787089d2cf6c4720736677543960c95062a8c37daa29d3d0b54a09e26570f1"></a>

<a id="canonical-cf69283f00702066e48b9c4db23935a78d772319f15a0b758167618fec2b3506"></a>

## name property — Property reference / 9580f5c2d307 / 9

Type: `"string"`. Required.

Name of the Application Profiles. Must be unique within the namespace.

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

<a id="canonical-42fc6a89ca5c2bbf2acf540aff6430ddc567e517ae5b6f41cda716b92dff7179"></a>

<a id="canonical-7dcf0206aab68a777ccb3b39f914eb986ce5739e4bbf4752440b8f2e08a826b5"></a>

## namespace property — Property reference / 9580f5c2d307 / 10

Type: `"string"`. Required.

Namespace where the Application Profiles is created.

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

- [timeouts](resources--application_profiles--reference--group-001.md#canonical-90d932bd8f79925ec5d6b9945630df17fac61100dbca671694a7e9d6a95b4b80): complete subsection reference.

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276): complete subsection reference.

<a id="canonical-66f9e0c878917df686b85fbc4648f2a5e1cacc95f45d40b8399ae21c397f1f7d"></a>

## All schema paths — Property reference / 9580f5c2d307 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_tcp_profile` | [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-1d8de3942644bd448e525f5b3e8973c3c2acbe1a4131c05656735be246d65fe8) |
| `advanced_tcp_profile.disable_tcp_advanced_profile` | [advanced_tcp_profile.disable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-a7aa6aadbea989fa9c84e8755cf9319f534fce2ae068c6ce8eb2f81c9cb26cae) |
| `advanced_tcp_profile.enable_tcp_advanced_profile` | [advanced_tcp_profile.enable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-d555dd2392205dab61e67e011a61deadb475f83d0601eaa2af474e9f3628cc7e) |
| `annotations` | [annotations](resources--application_profiles--reference--group-001.md#canonical-c3e31bff8bfa15fddccc24b1dccfe2698b073de424d50fdca8147b192f1be3a3) |
| `ddos_profile` | [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-6be19b43f9ccd108b6cbe05345c64ef57df0553c4e165e968ee557232295110c) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-79e5240a52fe43be5885723bdb8963780455b31fc307e57613a2b771fd99ed7c) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-6849bab2862656d1d4ec98685734fabc09cf6ac33dd9ff2284663a83ae915455) |
| `description` | [description](resources--application_profiles--reference--group-001.md#canonical-4aec4599e2d39de21dfafca0cb857b0827611f7bb2f83b052b50759e53c34099) |
| `disable` | [disable](resources--application_profiles--reference--group-001.md#canonical-4258643eb76a4199c29b4e7f8edb3bcace2685a9fa946685166d1eeff1ce62ab) |
| `id` | [id](resources--application_profiles--reference--group-001.md#canonical-5054148ade0a5f07062cc5915334bad222f6d6f5a06600106ad0c40c9c98643e) |
| `irules` | [irules](resources--application_profiles--reference--group-001.md#canonical-e1db3b24fc664083e1a28057e5d664756b66cc8cbfd4b08ca336a4c2332248b2) |
| `irules.kind` | [irules.kind](resources--application_profiles--reference--group-001.md#canonical-27c84041f7c637c4b3bfa8f5b1c2b8b2d871558f7fa402d21b1a2eb84583134b) |
| `irules.name` | [irules.name](resources--application_profiles--reference--group-001.md#canonical-2eb245aa564bd4c56ebb3c2b3948fc7e656b2d0a825dbf43c6c1f76d2c458bd2) |
| `irules.namespace` | [irules.namespace](resources--application_profiles--reference--group-001.md#canonical-dec890cbfdfddc1025695a1197e9567644e9fb3579375ed015d575df97751a5b) |
| `irules.tenant` | [irules.tenant](resources--application_profiles--reference--group-001.md#canonical-e1ca62d1466dc7876d743046c1737b0a198f95898d298cc540605112cab8b5f1) |
| `irules.uid` | [irules.uid](resources--application_profiles--reference--group-001.md#canonical-1a0c8e98670b9569218c3d71f4b4229c9e7ed831070939e814d5af1ea29b38b1) |
| `labels` | [labels](resources--application_profiles--reference--group-001.md#canonical-e80b43352d103cef5a019abda317d6be15e5c4120b6add542c5ac92a428b569d) |
| `name` | [name](resources--application_profiles--reference--group-001.md#canonical-79787089d2cf6c4720736677543960c95062a8c37daa29d3d0b54a09e26570f1) |
| `namespace` | [namespace](resources--application_profiles--reference--group-001.md#canonical-42fc6a89ca5c2bbf2acf540aff6430ddc567e517ae5b6f41cda716b92dff7179) |
| `timeouts` | [timeouts](resources--application_profiles--reference--group-001.md#canonical-cb0af81e6944e89a2340a75f3fb13ec59f670fb9a174ac02d1412f1fb99ac4b4) |
| `timeouts.create` | [timeouts.create](resources--application_profiles--reference--group-001.md#canonical-d0934290f09dcf77a29a38d373d1a237e8888ed3d4ccdce72d07235ed843cc49) |
| `timeouts.delete` | [timeouts.delete](resources--application_profiles--reference--group-001.md#canonical-2fb371a595a959d394d59ea724fb423e9a9e40dc73ad4578f73bc5ba39e9cb55) |
| `timeouts.read` | [timeouts.read](resources--application_profiles--reference--group-001.md#canonical-792417dab7cd03d3ba512bae84641e342972c7e59afb6692a30bab5819d82094) |
| `timeouts.update` | [timeouts.update](resources--application_profiles--reference--group-001.md#canonical-5c92e8f0a032781e74f7f2030a149133f98f0c71f11046eb32c87af411bfb3ab) |
| `virtual_server` | [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2b6f30603d702213b1c00e30bb081d82075e8fc771b28ae0759548e7ea8d2691) |
| `virtual_server.access_profile` | [virtual_server.access_profile](resources--application_profiles--reference--group-001.md#canonical-ab8de66f1cb320c80a79ce667b7c4b28df588d86a7d0e349a923824a7a228fcf) |
| `virtual_server.access_profile.kind` | [virtual_server.access_profile.kind](resources--application_profiles--reference--group-001.md#canonical-b6637b1a137e11f0c8c464f1980484805cce6b0491e0b328af0946aec4b6bf1c) |
| `virtual_server.access_profile.name` | [virtual_server.access_profile.name](resources--application_profiles--reference--group-001.md#canonical-ff009ada6de3ebcbf9e89b335c1b7790cfb20f93edaa0fd949be60bcec4fdb2c) |
| `virtual_server.access_profile.namespace` | [virtual_server.access_profile.namespace](resources--application_profiles--reference--group-001.md#canonical-12fd29faff78f39c3a0ff63692b615ba4f7dc4231e281cf0b443aa3746c9c881) |
| `virtual_server.access_profile.tenant` | [virtual_server.access_profile.tenant](resources--application_profiles--reference--group-001.md#canonical-75b57a8a21de63a301b0be6f0fa9e976cfe6cf248c098ab34bcd60196c859a56) |
| `virtual_server.access_profile.uid` | [virtual_server.access_profile.uid](resources--application_profiles--reference--group-001.md#canonical-09d53c318b2dd95fdb443055e732ecc1c1079aaca916c4fec7c1788686c7b70c) |
| `virtual_server.address_translation` | [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-303b6501ee26d1cf62be8306849f86e75608671428768441e8e3dd8105f33ffa) |
| `virtual_server.address_translation.address_translation_disable` | [virtual_server.address_translation.address_translation_disable](resources--application_profiles--reference--group-001.md#canonical-066399375ec4ab4dde58ddfc517ea3d6a6ac51bbcbaf2dcc4c86725685c8e82c) |
| `virtual_server.address_translation.address_translation_enable` | [virtual_server.address_translation.address_translation_enable](resources--application_profiles--reference--group-001.md#canonical-3dbdb8919f290a174f662ab33a6c59db45b7cf8af8d15d12cf486da14bce7f6a) |
| `virtual_server.auto_last_hop` | [virtual_server.auto_last_hop](resources--application_profiles--reference--group-001.md#canonical-920ff090d872f321ac2c53da5aa0d40d0219faec680099acb7cb7566e692f88a) |
| `virtual_server.auto_last_hop.auto_last_hop_default` | [virtual_server.auto_last_hop.auto_last_hop_default](resources--application_profiles--reference--group-002.md#canonical-20d93a775356004e690a8174165eb52fefe84f63ec1e369a5c1af4bdb7b4addd) |
| `virtual_server.auto_last_hop.auto_last_hop_disable` | [virtual_server.auto_last_hop.auto_last_hop_disable](resources--application_profiles--reference--group-002.md#canonical-ffef5c294e93b45dd41f9ce82143ec9b2ad25e1070e18fb56b033880bd3aca45) |
| `virtual_server.auto_last_hop.auto_last_hop_enable` | [virtual_server.auto_last_hop.auto_last_hop_enable](resources--application_profiles--reference--group-002.md#canonical-9b45997e5b3c9d9d28b0e3f876629c0a37c29f46ec5a5d4b3e263e14db7177b5) |
| `virtual_server.clone_pool_client` | [virtual_server.clone_pool_client](resources--application_profiles--reference--group-002.md#canonical-4f32c96249ffb1bffba76f45f86bd5a3a8639bad04a58ee474df140aa0c3a592) |
| `virtual_server.clone_pool_client.kind` | [virtual_server.clone_pool_client.kind](resources--application_profiles--reference--group-002.md#canonical-d8ddbd66c0b4c88c64348d18ff924fafe27df1694c244ce698daf9328b888210) |
| `virtual_server.clone_pool_client.name` | [virtual_server.clone_pool_client.name](resources--application_profiles--reference--group-002.md#canonical-135a7c6cff945e9c2df93d486d97c3138406b97345d6d1046d4109f62f6bb2c5) |
| `virtual_server.clone_pool_client.namespace` | [virtual_server.clone_pool_client.namespace](resources--application_profiles--reference--group-002.md#canonical-31dcf0d39d2e0da67158a1cea59da4a85a284ea48de174ffc6f08e025337354b) |
| `virtual_server.clone_pool_client.tenant` | [virtual_server.clone_pool_client.tenant](resources--application_profiles--reference--group-002.md#canonical-5d81578f653f5e480ae2cc99a22c902813764e6e81d0638f3513acb4d4a28a1d) |
| `virtual_server.clone_pool_client.uid` | [virtual_server.clone_pool_client.uid](resources--application_profiles--reference--group-002.md#canonical-dbbf9f383d3dd3a4ee837ab18cc9ac6da30ca5df90f49b9aca79f1eb21f144f5) |
| `virtual_server.clone_pool_server` | [virtual_server.clone_pool_server](resources--application_profiles--reference--group-002.md#canonical-47a5db7314ee190971fb53ab7f2cc07890a098cd220bbeebce46f6ee980238d9) |
| `virtual_server.clone_pool_server.kind` | [virtual_server.clone_pool_server.kind](resources--application_profiles--reference--group-002.md#canonical-6b67401d0ac43159a9693c970c5432566cca4a5106c1e3e202efdcc68cf0d94e) |
| `virtual_server.clone_pool_server.name` | [virtual_server.clone_pool_server.name](resources--application_profiles--reference--group-002.md#canonical-b211efb1a5b692ec54515b4c72ff51ab1cd55f39ba3eb1b734fa6ca338f1169f) |
| `virtual_server.clone_pool_server.namespace` | [virtual_server.clone_pool_server.namespace](resources--application_profiles--reference--group-002.md#canonical-cc498cbd1c084e3cfee6f54277f2f7359ae2fca5e53318b9c90872bbe88b649d) |
| `virtual_server.clone_pool_server.tenant` | [virtual_server.clone_pool_server.tenant](resources--application_profiles--reference--group-002.md#canonical-d7f0ce26a4bf58e2b96158a52d5442ad2479eec13b09b10de59c81361c6ecbc0) |
| `virtual_server.clone_pool_server.uid` | [virtual_server.clone_pool_server.uid](resources--application_profiles--reference--group-002.md#canonical-88d880e0da81aa81dda10bba89b8e7e243f96e9ab75b78d37e2e759543db3ed2) |
| `virtual_server.connection_limit` | [virtual_server.connection_limit](resources--application_profiles--reference--group-001.md#canonical-f2123abf1bf98cac5f0620f786256465a818118895a1f9289b876f1c605ae8ac) |
| `virtual_server.connection_rate_limit` | [virtual_server.connection_rate_limit](resources--application_profiles--reference--group-001.md#canonical-e645b6dba2cd08dfd3b100fa336d24202c0dc6ba95cbd582d4e195e6a38ace73) |
| `virtual_server.connection_rate_limit_mode` | [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-f12ffc0ea41f370ffe30ade4757819abca1d954e29cddf9a896e6e0729f88ff5) |
| `virtual_server.connection_rate_limit_mode.per_destination_address` | [virtual_server.connection_rate_limit_mode.per_destination_address](resources--application_profiles--reference--group-002.md#canonical-f7bb2e027a752d76e27842cdb9709607837e07079ae039b9c81704c5f2a8b783) |
| `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-45b7be542810657d01d8e725ac182298a3b7a85183395f542ad2101d423688ff) |
| `virtual_server.connection_rate_limit_mode.per_source_address` | [virtual_server.connection_rate_limit_mode.per_source_address](resources--application_profiles--reference--group-002.md#canonical-4ede28fda7e31a42a5911e68df6ecea871921f9d39371668bea639b57b71877b) |
| `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-cc7bdaedb5659d490bf0bba89ce43c2239edcbd9bfc60ebf9eed6483492fa613) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-59e44cbde41897e17503e54b129bfd78727c3a7ffec5da767ae718f7a0ed7883) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-dc63645a71444fe3799628fdef2876367609e9b3d3896566b4a257e16afb4550) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-dd5b8c7751ba55c5c165d07211d8eed54001d74ea2d80fa05d2ef509cd655ef5) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server` | [virtual_server.connection_rate_limit_mode.per_virtual_server](resources--application_profiles--reference--group-002.md#canonical-dc5cc92ef6ae7b0cbae6271372691bfd18d1ee03440c64a3432c274c4afcaec6) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](resources--application_profiles--reference--group-002.md#canonical-86297c17b1c57158886d5cda4475f8086642a0ca94745b1bb3920bd470ac7b08) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-40152993d777b3704fe5fd9cc5559db30efbc4e79cc6a0d22e73477a06bcffef) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](resources--application_profiles--reference--group-002.md#canonical-91cdc9156849b5112df4b13563747134cb7a9003ce81e870eacc96a01ff18f41) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-3fb663284937d54439962a85e9a0f58eaf6a4cf81e9858b6ebf73b03215946ff) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-3dd042b4b29ceae96bb83850c913e4b5a571750a2b045ac3e6d2be660bb9cded) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-5d495f98dd778763e3609e692f66a4e3a7d7c065902d0fb6df1581a8f95ac4b7) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-6d3e74a60870f69f3ca92699a5eddb0620851548489b4debfa84a60240d9cd97) |
| `virtual_server.default_persistence_profile` | [virtual_server.default_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-c3f266ca622f2848aa75cca8d358b4ee181c9d6471604a3f0db6332f810194d8) |
| `virtual_server.default_persistence_profile.kind` | [virtual_server.default_persistence_profile.kind](resources--application_profiles--reference--group-002.md#canonical-4e317c6fc6d38154b9d796dc443eb954c2f0d4244f46d14e09f71aa776010812) |
| `virtual_server.default_persistence_profile.name` | [virtual_server.default_persistence_profile.name](resources--application_profiles--reference--group-002.md#canonical-47c8043030afd40cb1303077e59bf83128a0f2c99e63d936e1269f212d4e5571) |
| `virtual_server.default_persistence_profile.namespace` | [virtual_server.default_persistence_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-1b51fd0bbca2e168845f6e646758cac58d4b09f1fd1f933b2651b2718a714dc7) |
| `virtual_server.default_persistence_profile.tenant` | [virtual_server.default_persistence_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-9e1152c3744f0d150ba9ef3ce31017c156351831f6164c82614b67509d69e570) |
| `virtual_server.default_persistence_profile.uid` | [virtual_server.default_persistence_profile.uid](resources--application_profiles--reference--group-002.md#canonical-b260929e476da1c2a7bd7d2ecf6278124e0db0cf935cf4c0ebcd997e7865babd) |
| `virtual_server.default_pool` | [virtual_server.default_pool](resources--application_profiles--reference--group-002.md#canonical-fbc13bd0a0bf3d715abae8a270861bd5f8804324c19c8178bca9c16b07f287c2) |
| `virtual_server.default_pool.kind` | [virtual_server.default_pool.kind](resources--application_profiles--reference--group-002.md#canonical-ef17257bc4f6a6fa97a5786630448ec0d7dcde5c1d301bcb27420a23ac0f4821) |
| `virtual_server.default_pool.name` | [virtual_server.default_pool.name](resources--application_profiles--reference--group-002.md#canonical-427a86545c3880cf823d97ae982cf0ed0a522d2cfe2854dd27e415dc20a4134e) |
| `virtual_server.default_pool.namespace` | [virtual_server.default_pool.namespace](resources--application_profiles--reference--group-002.md#canonical-142f32fcf47873278529e2b2eca27bf7717fbfd18f8fe5b0ff8042c9271c194a) |
| `virtual_server.default_pool.tenant` | [virtual_server.default_pool.tenant](resources--application_profiles--reference--group-002.md#canonical-f26dbda878eb706711ef122922dbf72d545b812e60521c25469e8d230f11bfa8) |
| `virtual_server.default_pool.uid` | [virtual_server.default_pool.uid](resources--application_profiles--reference--group-002.md#canonical-062ff8028c23a9fbc6a87fe969b801c8ea9b1dcb7752bf4272debc54fec29013) |
| `virtual_server.fallback_persistence_profile` | [virtual_server.fallback_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-c9c0b08180f15ac53d1ca12cfcdd7795fb3ae243dac4768ba64ab15cf0715a7d) |
| `virtual_server.fallback_persistence_profile.kind` | [virtual_server.fallback_persistence_profile.kind](resources--application_profiles--reference--group-002.md#canonical-7c5ac21109d36a7d1365668824e5656bd307b5388e932e6f05e34b4b4c3ec2a3) |
| `virtual_server.fallback_persistence_profile.name` | [virtual_server.fallback_persistence_profile.name](resources--application_profiles--reference--group-002.md#canonical-8ca41731da345d233347037400797e646b67271786be0c0cb13963c016446b35) |
| `virtual_server.fallback_persistence_profile.namespace` | [virtual_server.fallback_persistence_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-5bee1641877f188617494eb0fd77a60ceda9e1cabb03e7cd7109152cd9f2c484) |
| `virtual_server.fallback_persistence_profile.tenant` | [virtual_server.fallback_persistence_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-7fef36efa67cf3ce98b926c26809a8025bd252e74d5c21a9185d8ac4550cd7ec) |
| `virtual_server.fallback_persistence_profile.uid` | [virtual_server.fallback_persistence_profile.uid](resources--application_profiles--reference--group-002.md#canonical-f19171b8a6e3c03c45980af572a81a168ea30789191bd83d1d2ebcb2fe2a30f2) |
| `virtual_server.fix_profile` | [virtual_server.fix_profile](resources--application_profiles--reference--group-002.md#canonical-a8f2bace83804e3ade89b419a563803a4761cd4dd15614fff57193789a9f2fed) |
| `virtual_server.fix_profile.kind` | [virtual_server.fix_profile.kind](resources--application_profiles--reference--group-002.md#canonical-be2be6c707b486066d1212d19e08952b01d960b27cfdd36964e0ccfc9a94f82f) |
| `virtual_server.fix_profile.name` | [virtual_server.fix_profile.name](resources--application_profiles--reference--group-002.md#canonical-af284ecda887a8874260113a0487cd1120610fc97f331ab4cf472eca39fdbe90) |
| `virtual_server.fix_profile.namespace` | [virtual_server.fix_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-2339c035b99b13f1bddc5a004339cf2cfcd2dfead9bc28fdad9f099ae7f1637b) |
| `virtual_server.fix_profile.tenant` | [virtual_server.fix_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-9966a40f58a26578b4c627186a84c34a1c6d876d91ac3a6131e0aa22fb52e4ba) |
| `virtual_server.fix_profile.uid` | [virtual_server.fix_profile.uid](resources--application_profiles--reference--group-002.md#canonical-62afb4a00640259e068307a977a3c7ee9c2d8bc2ac99240dcbff62f90a81e12e) |
| `virtual_server.http` | [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-be100f8814f66fb0fcd60f20ea004b7769d73f28377ab2adf17337410f4ef9e2) |
| `virtual_server.http.client_ssl_profile` | [virtual_server.http.client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-83a5eafd66fee84be5d7fd0da9895c697b907b3ef6a31a7033e484a39a614bff) |
| `virtual_server.http.client_ssl_profile.kind` | [virtual_server.http.client_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-f0f870b6e4571355c6891318e5ec776e76a81e34a54df4d510331b423b90a184) |
| `virtual_server.http.client_ssl_profile.name` | [virtual_server.http.client_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-8d554e2f0fcebd80079c7510814b45f290d4caf320fe8d11eaa3029a9db25e97) |
| `virtual_server.http.client_ssl_profile.namespace` | [virtual_server.http.client_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-ff3be08a0b90a53ee2099f21db6c1e1fc0a476ab6aa62f4de54f61355166048f) |
| `virtual_server.http.client_ssl_profile.tenant` | [virtual_server.http.client_ssl_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-21dbd1eac9bd5724345c0c27b46acbd96593f90d388ce9e0763637a4a02b718a) |
| `virtual_server.http.client_ssl_profile.uid` | [virtual_server.http.client_ssl_profile.uid](resources--application_profiles--reference--group-002.md#canonical-8ee7036152e40941809dc7b4edf37d9cc6affc9bd2e07a85c2be8c9a6642a834) |
| `virtual_server.http.http2_client_profile` | [virtual_server.http.http2_client_profile](resources--application_profiles--reference--group-002.md#canonical-4ad96fa02aadc14b65e7b8e46a224d4b5815142e8b5b6c9f0681d8fc6e5ab5c7) |
| `virtual_server.http.http2_client_profile.kind` | [virtual_server.http.http2_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-d10bded6ab7fd72ad5edb6ebd5273372679c4e73f00fbdf3d1e9cd6213719efe) |
| `virtual_server.http.http2_client_profile.name` | [virtual_server.http.http2_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-fa74e65c37fa21cc4ef58de74582d52753a855e647d5dd634ec0903f8fc3a2c4) |
| `virtual_server.http.http2_client_profile.namespace` | [virtual_server.http.http2_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3ff3ab5d5be9d38db5df1294549520c1e560b5feafbf0f7772ca40f8716525bc) |
| `virtual_server.http.http2_client_profile.tenant` | [virtual_server.http.http2_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-93e11d74477322f9520e849a1e079995ca324869e5f791ca45f4356d0ee2cbc2) |
| `virtual_server.http.http2_client_profile.uid` | [virtual_server.http.http2_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-5956e30db75f9238e5eb54c239d3d331e1b5092cc28e9f7f8130a6bcb928ff3a) |
| `virtual_server.http.http2_server_profile` | [virtual_server.http.http2_server_profile](resources--application_profiles--reference--group-002.md#canonical-495ebf722889df649b7989f1f499186895bfd64f0a6bd655e1025bdbd826ee2e) |
| `virtual_server.http.http2_server_profile.kind` | [virtual_server.http.http2_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-a92c8af517f5ab83db86995e00f1640177d58194800545fc96289efef4288f0e) |
| `virtual_server.http.http2_server_profile.name` | [virtual_server.http.http2_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-ce17310893223087104f3a01a3bde402b6c8ceb81e9f5121a5bf21ccb2f3a99f) |
| `virtual_server.http.http2_server_profile.namespace` | [virtual_server.http.http2_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-c14ca3ec7f0537554aa9e0d017941b4d99e58053e80d3d0f919a056d44a67cd9) |
| `virtual_server.http.http2_server_profile.tenant` | [virtual_server.http.http2_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-4da36a1b45bd07b46cf5b4756f62884febd9185c0d0a54b77611753115129227) |
| `virtual_server.http.http2_server_profile.uid` | [virtual_server.http.http2_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-3c73e4a4543caa060d7b85c8f3a389059e914dac83ba58f8d32e5e8244a3f731) |
| `virtual_server.http.http_client_profile` | [virtual_server.http.http_client_profile](resources--application_profiles--reference--group-002.md#canonical-5742a25d85d2410a182d1e903c8b5e40de82e5ecbb504bc2f9d61bb06a5ef62f) |
| `virtual_server.http.http_client_profile.kind` | [virtual_server.http.http_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-1127febaf37a1d6c9034ed9b7f53655fa864564871742fca77732aafab6dddc2) |
| `virtual_server.http.http_client_profile.name` | [virtual_server.http.http_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-e8cf1acb9b95036820bd514055afa0fb10c63079df512b14919a41842cdf1ec5) |
| `virtual_server.http.http_client_profile.namespace` | [virtual_server.http.http_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-4c5734bd72e68fc16111f252e7ceb905ed0f6f127ef7ad3ca195e11776548bca) |
| `virtual_server.http.http_client_profile.tenant` | [virtual_server.http.http_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-f042f1d24613808df27421aa64d17c04283ae12e53743b853984ac809630dbd0) |
| `virtual_server.http.http_client_profile.uid` | [virtual_server.http.http_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-357a3eb52f7018706f29397f9f1f822f15e337cea67966842fa07773822b0a06) |
| `virtual_server.http.http_server_profile` | [virtual_server.http.http_server_profile](resources--application_profiles--reference--group-002.md#canonical-6753002b9ae7cae9db085027097b0f820e28bacbdff0e42970234e7a2bea424e) |
| `virtual_server.http.http_server_profile.kind` | [virtual_server.http.http_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-1b59b620122683e16072547c4e201c730e1fd7319c9845e956895bb95bb9b7ef) |
| `virtual_server.http.http_server_profile.name` | [virtual_server.http.http_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-7056d90e48104745c6d2221776949adbdf055dbe21c00b851fe07f76b27d1083) |
| `virtual_server.http.http_server_profile.namespace` | [virtual_server.http.http_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-d9bba788fe2e1f9b2f8ce98b281a72fab717fb77b7056b83f0aaa348d047a62f) |
| `virtual_server.http.http_server_profile.tenant` | [virtual_server.http.http_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-bbdbb0f7cfdbdea5e3df12299e56e9837febab06fe5cc9ea2dae52e6199306ca) |
| `virtual_server.http.http_server_profile.uid` | [virtual_server.http.http_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1a7f607d4ad231322f193fa77584ad6c165ba6ad2d13c755234259dcae89bea2) |
| `virtual_server.http.ocsp_profile` | [virtual_server.http.ocsp_profile](resources--application_profiles--reference--group-002.md#canonical-d0c423297f7e5f3ecb9e725c4c40141bade3df05939bcd103e717df2bb72802b) |
| `virtual_server.http.ocsp_profile.kind` | [virtual_server.http.ocsp_profile.kind](resources--application_profiles--reference--group-002.md#canonical-279c02cc8d9f3570b1147da8ef33cffdb3e2f5abb29c86958a7430eb2666f886) |
| `virtual_server.http.ocsp_profile.name` | [virtual_server.http.ocsp_profile.name](resources--application_profiles--reference--group-002.md#canonical-1c631150e0ae6646a20720de3312fe2e05ad926f311c6d0083dcc04f67412fe0) |
| `virtual_server.http.ocsp_profile.namespace` | [virtual_server.http.ocsp_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-308af13da595e90ad3b629cfbf6025ae9279f970fd4ba3580c3fc1e7ea86c25b) |
| `virtual_server.http.ocsp_profile.tenant` | [virtual_server.http.ocsp_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-35c58cf0fa9f1c469ccdd3b10e8bb068666fdf10112865b9c80e62f0308641a5) |
| `virtual_server.http.ocsp_profile.uid` | [virtual_server.http.ocsp_profile.uid](resources--application_profiles--reference--group-002.md#canonical-c1f589fc6ab6eae054244c141cef53c70762fa99db1621eed50d6d42904d5da9) |
| `virtual_server.http.server_ssl_profile` | [virtual_server.http.server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-c47f3c05ea5b0cf861ac86bc9f8211069b4da62259cf9fa2f8dbd6e280ff2318) |
| `virtual_server.http.server_ssl_profile.kind` | [virtual_server.http.server_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-46d308eea4d4d7c37a32ea3b8103136dd82b7ecf8ac5f9effc01160c33c064a4) |
| `virtual_server.http.server_ssl_profile.name` | [virtual_server.http.server_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-1770ff03814a771267f66a919e28eea6b40e10e2f9adc65d79534f6fd3dbef8a) |
| `virtual_server.http.server_ssl_profile.namespace` | [virtual_server.http.server_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-d3d6aebeb74776bff3d68b2d0695d49ca74396acfe22e3d94d163d4f1a1664b4) |
| `virtual_server.http.server_ssl_profile.tenant` | [virtual_server.http.server_ssl_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-35f79c698e7d7973969edc54214d4a20c33449c55ffa2a9f94973a3ff6d1c05f) |
| `virtual_server.http.server_ssl_profile.uid` | [virtual_server.http.server_ssl_profile.uid](resources--application_profiles--reference--group-002.md#canonical-c00aafa71cbf34a6d92e642d8925b506738f78665cb87c648564375336fb38dc) |
| `virtual_server.http.stream_profile` | [virtual_server.http.stream_profile](resources--application_profiles--reference--group-002.md#canonical-d21cc50250a46eccf61cf593c62f068a1e974484f108e94a940910f64ff451a6) |
| `virtual_server.http.stream_profile.kind` | [virtual_server.http.stream_profile.kind](resources--application_profiles--reference--group-002.md#canonical-0fefb597d9f7e69575623d80e35d92627a4df209e3b2110601ac1613aeac1e48) |
| `virtual_server.http.stream_profile.name` | [virtual_server.http.stream_profile.name](resources--application_profiles--reference--group-002.md#canonical-bfbd6f92080fb13659a7ea75069eb4a3b5e8f23f9c1c78e5f7b69c131e80e0df) |
| `virtual_server.http.stream_profile.namespace` | [virtual_server.http.stream_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-8872ef279147b998040378c7824e20dcef2c61fe9a49c8fbb992a6b8c59174d6) |
| `virtual_server.http.stream_profile.tenant` | [virtual_server.http.stream_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-dfc538eee365180bdfef87295901851d1d9a11952876a2b2af67ea4fb0c1d137) |
| `virtual_server.http.stream_profile.uid` | [virtual_server.http.stream_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1274afb571d1ce4711eb833a1cfd18f1bd30f13db488e92db12543ad1dd28966) |
| `virtual_server.http.tcp_client_profile` | [virtual_server.http.tcp_client_profile](resources--application_profiles--reference--group-002.md#canonical-aa4c29ea798fdc8a66ff14456e5a95dd896fa903fc1ee8e97b954264e3d92e5f) |
| `virtual_server.http.tcp_client_profile.kind` | [virtual_server.http.tcp_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-d09004da5c753588f33ae3b06fa3fae70bef7b426310a106416ed603fddfa546) |
| `virtual_server.http.tcp_client_profile.name` | [virtual_server.http.tcp_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-9153f46f706f37ea6252caf279489c23c2d00343baaccb082267d7184bd52850) |
| `virtual_server.http.tcp_client_profile.namespace` | [virtual_server.http.tcp_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-07d61ef6a5db6d8a4ac953027d5181a60bc0d8a08edf96f54c3ee87e08613b00) |
| `virtual_server.http.tcp_client_profile.tenant` | [virtual_server.http.tcp_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-3982bc911074ed552a88f10ccfbff1027eea6c8ab5144690c8f7487a9eb41f38) |
| `virtual_server.http.tcp_client_profile.uid` | [virtual_server.http.tcp_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-83133d070d905b1f397ed78ad49a644ca024fa6ee7d25d0931072b7cd6eb976d) |
| `virtual_server.http.tcp_server_profile` | [virtual_server.http.tcp_server_profile](resources--application_profiles--reference--group-002.md#canonical-0ee9802d979a3ddafae8462df15575498db37e90a8a77149f36745de0ac10d0e) |
| `virtual_server.http.tcp_server_profile.kind` | [virtual_server.http.tcp_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-fcb618c34daa2d3218a088bbbaec6385f5c18c582ac7268fb1e4b06b05022040) |
| `virtual_server.http.tcp_server_profile.name` | [virtual_server.http.tcp_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-d1973ace6d238a2bca234e8cdfa7c2116106365e918a9fad1321de9cc608162e) |
| `virtual_server.http.tcp_server_profile.namespace` | [virtual_server.http.tcp_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-c93850fa7ef2822f4b68f7567661b80b27e07dc7de0f754ca719f22628a04739) |
| `virtual_server.http.tcp_server_profile.tenant` | [virtual_server.http.tcp_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-2b159c0ba6ba7de8a2b0db8b7eaf82cea5f262c27ac43dc547fbc68b7740de2b) |
| `virtual_server.http.tcp_server_profile.uid` | [virtual_server.http.tcp_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-c653ab7189a30c21354007e8ce38ff099ee6f347ee891f0786a3ba7a8b6e11ba) |
| `virtual_server.http.websocket_client_profile` | [virtual_server.http.websocket_client_profile](resources--application_profiles--reference--group-002.md#canonical-ab52863167afe6887859ffff2194488d4a78df703e3a2ec67389ae7eada58a0e) |
| `virtual_server.http.websocket_client_profile.kind` | [virtual_server.http.websocket_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-cc384775414ebd3b051e96c7074953ba8bb0fcf7cf74f4dd070178517727752a) |
| `virtual_server.http.websocket_client_profile.name` | [virtual_server.http.websocket_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-5acbae8192e39259d4f28faa053bceb435d9614c0398b0c5d35c707cd87943d1) |
| `virtual_server.http.websocket_client_profile.namespace` | [virtual_server.http.websocket_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-d149b88e974b35dcc8dba306d1dbe28d50defcff5dd09c7aaf38516c013d44c6) |
| `virtual_server.http.websocket_client_profile.tenant` | [virtual_server.http.websocket_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-e08dc767a53583c353a414b02284c1df0b85259d540d754ef103de4ec4cfaa2a) |
| `virtual_server.http.websocket_client_profile.uid` | [virtual_server.http.websocket_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-7d5716a44b3fbf99d6dbc0ae6b926a58be306ba752bbb4f88eea10a2e9347e3b) |
| `virtual_server.http.websocket_server_profile` | [virtual_server.http.websocket_server_profile](resources--application_profiles--reference--group-002.md#canonical-62bb237307ee6103880620b853a7d2a6179ca55a92f1519f0a75ace1131a7d25) |
| `virtual_server.http.websocket_server_profile.kind` | [virtual_server.http.websocket_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-9d26aeb27458a785569c7db62be0d3029123905c595f989a6681a378b54c723d) |
| `virtual_server.http.websocket_server_profile.name` | [virtual_server.http.websocket_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-6d5d7238429c66dcd2f69ed99173dba58103f4b20e4900a4649e5cccf49732ac) |
| `virtual_server.http.websocket_server_profile.namespace` | [virtual_server.http.websocket_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-ff46203d3de3fad834e7f2d87c3cb15bb08d516099710b109fa70f97ad2e892b) |
| `virtual_server.http.websocket_server_profile.tenant` | [virtual_server.http.websocket_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-559f676202010b9ab13d80566f2fb8d12c116bf23924d5ad5b3e9af5369797fd) |
| `virtual_server.http.websocket_server_profile.uid` | [virtual_server.http.websocket_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-0d86c4a9f9546601ebcff13fb3b602d5ff49e0b200c6e1c92dbd935ac99ed745) |
| `virtual_server.http3` | [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-82b5c1d978044ed9bdcb8dfaa9217bbe579e55214ab015d15ec5eeaad83ab05d) |
| `virtual_server.http3.client_ssl_profile` | [virtual_server.http3.client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-58533851710533e3814eb95b4be0e1fc5640ade8c51078e6d46817abd516b63f) |
| `virtual_server.http3.client_ssl_profile.kind` | [virtual_server.http3.client_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-a08234d2df1cc8ac498c20e994fca195d4c87e7c7dc9100375a5cc207f8eb355) |
| `virtual_server.http3.client_ssl_profile.name` | [virtual_server.http3.client_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-e97bc76ebece692614ae4301654202f09e024e345ffaa0e907dce86319be3de8) |
| `virtual_server.http3.client_ssl_profile.namespace` | [virtual_server.http3.client_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-4deecedef31873b3c37e964a8a00088ce356e82a75625b8f09d7349ec946e7b9) |
| `virtual_server.http3.client_ssl_profile.tenant` | [virtual_server.http3.client_ssl_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-5f49d14f35877413faefddb22b0132aeb84f226405bc2ab0e2e4692ee922be2c) |
| `virtual_server.http3.client_ssl_profile.uid` | [virtual_server.http3.client_ssl_profile.uid](resources--application_profiles--reference--group-002.md#canonical-9f312a09fd7f3bf6043c72cb691600882e8f7c4809d2a91de3635669c40e579e) |
| `virtual_server.http3.http3_profile` | [virtual_server.http3.http3_profile](resources--application_profiles--reference--group-002.md#canonical-3a80373490037174f45158df3afc4ca797446eaac925a9947883fc0a44434eb6) |
| `virtual_server.http3.http3_profile.kind` | [virtual_server.http3.http3_profile.kind](resources--application_profiles--reference--group-002.md#canonical-2b2fd240ab34430bb52716415ca80143b37c25dc12cbdc0aa26a0baf697eacb7) |
| `virtual_server.http3.http3_profile.name` | [virtual_server.http3.http3_profile.name](resources--application_profiles--reference--group-002.md#canonical-6a82dd5051c80b0cad951473b457c3cbaa24dcdea6fc234c30cbcfa67f10ebe4) |
| `virtual_server.http3.http3_profile.namespace` | [virtual_server.http3.http3_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3a58979fc93842068bc42c15a2af05760449cf7ef66ba79be9aac76fad7a13d1) |
| `virtual_server.http3.http3_profile.tenant` | [virtual_server.http3.http3_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-4af151ebd2d757944610c0f118994ad5c64410b0d1636609154e8cc6ddca1e4a) |
| `virtual_server.http3.http3_profile.uid` | [virtual_server.http3.http3_profile.uid](resources--application_profiles--reference--group-002.md#canonical-7aee224d8611077faf367e7f1b649b9afdfdb1619fb9dd3f1418a2e88c2b0071) |
| `virtual_server.http3.http_client_profile` | [virtual_server.http3.http_client_profile](resources--application_profiles--reference--group-002.md#canonical-95b581bc6fa87226870e13fdf297d62d7778add7141b79861d632a97af27a839) |
| `virtual_server.http3.http_client_profile.kind` | [virtual_server.http3.http_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-6eef6db82383712a22446626f827d63ae686cb331ee4cc96be7a71b170a6ce25) |
| `virtual_server.http3.http_client_profile.name` | [virtual_server.http3.http_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-dfed9098284043495ecbe1d801c4957beaee6db23d9d3fce1f0134ae700fea9b) |
| `virtual_server.http3.http_client_profile.namespace` | [virtual_server.http3.http_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-5b7c310649dff87743d71d97df60bc63af669f49a385116fd813e787e555a77c) |
| `virtual_server.http3.http_client_profile.tenant` | [virtual_server.http3.http_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-39e1c252fc672dca41d8b80fcbeaa0d469014c70dba98e4fc018a15f241bfbf7) |
| `virtual_server.http3.http_client_profile.uid` | [virtual_server.http3.http_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-521cd7bb02d66e7f054e424a90ef7e6f62d78762460f44d5978a97a6cca40854) |
| `virtual_server.http3.http_server_profile` | [virtual_server.http3.http_server_profile](resources--application_profiles--reference--group-002.md#canonical-9f73257b35f97112a32719a9f50d76e1821daa7423a4618b23049ce346a7b58a) |
| `virtual_server.http3.http_server_profile.kind` | [virtual_server.http3.http_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-b2e5bf35850dd7bf3be44890096e9d32ce194e293d9ea80eee0d3e0c586cad55) |
| `virtual_server.http3.http_server_profile.name` | [virtual_server.http3.http_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-a7d2b8be19e15b61c33670e3bfbc311337010875071eb529abf0fac442376d7b) |
| `virtual_server.http3.http_server_profile.namespace` | [virtual_server.http3.http_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-4ae050bf0bf3e86cd73c696b48878350c609629ccda2ac993f177311c7278239) |
| `virtual_server.http3.http_server_profile.tenant` | [virtual_server.http3.http_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-8db6451b236d063dc0d6998958a8bb9560ecb9fecce73fc624fc9dea4d5565fc) |
| `virtual_server.http3.http_server_profile.uid` | [virtual_server.http3.http_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-fcd1e9023285094ec6da22ba254ca40a34dfe4e7f2ab56abf61064020cdd3a67) |
| `virtual_server.http3.quic_profile` | [virtual_server.http3.quic_profile](resources--application_profiles--reference--group-002.md#canonical-72bb975309553837bd1700d301397f539bc2ef2e223fe0c5319880afe34cb023) |
| `virtual_server.http3.quic_profile.kind` | [virtual_server.http3.quic_profile.kind](resources--application_profiles--reference--group-002.md#canonical-93c74c73a7033f3e6437d6662563ab4482ee8cfcfaf6b7e88900a3a420060783) |
| `virtual_server.http3.quic_profile.name` | [virtual_server.http3.quic_profile.name](resources--application_profiles--reference--group-002.md#canonical-45776d93be50cda97e95600f9a60b8bc0c8459f36b46c8b35249910a47899cc0) |
| `virtual_server.http3.quic_profile.namespace` | [virtual_server.http3.quic_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-ed88856e8b99b64a07adb934404d2c9f41953c9a6fb20afbca0addacb839ee43) |
| `virtual_server.http3.quic_profile.tenant` | [virtual_server.http3.quic_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-3db4d404de4e4e140a6b4877f9fb3412c5f9f1d83859b5b634443ae3446f14a8) |
| `virtual_server.http3.quic_profile.uid` | [virtual_server.http3.quic_profile.uid](resources--application_profiles--reference--group-002.md#canonical-44b14a7e629edeb1a543305f6f2ae052c742eac05d952043c8402b61779914bc) |
| `virtual_server.http3.server_ssl_profile` | [virtual_server.http3.server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-c520e61eb27e45e3adf20a0885c1ada89cd2172a636e574409dd8475ecfe4275) |
| `virtual_server.http3.server_ssl_profile.kind` | [virtual_server.http3.server_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-db45e6355992cf7d6f6f55c5fa584d15021ba67dc88fa39ed7bf9f22769c585e) |
| `virtual_server.http3.server_ssl_profile.name` | [virtual_server.http3.server_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-f377152149cd0ecf118b9011e47b9ae6e067eebb4f0b35348296934a9ee36827) |
| `virtual_server.http3.server_ssl_profile.namespace` | [virtual_server.http3.server_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-fcab7c3b70cd54ea66d2f7d4b47f3c3604efe746f8884dc875b55ad82cb6d77f) |
| `virtual_server.http3.server_ssl_profile.tenant` | [virtual_server.http3.server_ssl_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-bcdd6d51bbf7a95ecedf95579f017882b0a3ead4ac706bcd8408422a033f8422) |
| `virtual_server.http3.server_ssl_profile.uid` | [virtual_server.http3.server_ssl_profile.uid](resources--application_profiles--reference--group-002.md#canonical-230e6816f4c31d25bbf3297dfa51ebda4324cfd54673a7b1030f5959014cd2b8) |
| `virtual_server.http3.tcp_server_profile` | [virtual_server.http3.tcp_server_profile](resources--application_profiles--reference--group-002.md#canonical-962a230b107cab3903eaf8466fc9fe9d691985c570c7a4a1e8bae1b0b796bb1d) |
| `virtual_server.http3.tcp_server_profile.kind` | [virtual_server.http3.tcp_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-ef9aa2178731a82676248fafa7777cf68a3e61a07eb543aabc5a8378639d487d) |
| `virtual_server.http3.tcp_server_profile.name` | [virtual_server.http3.tcp_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-77f501c64c36bd49c6c7d572d213112abdce32851bcc97c475395a8f79c70015) |
| `virtual_server.http3.tcp_server_profile.namespace` | [virtual_server.http3.tcp_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-6ad069c2db474d1373429324c44a36fceb3ecc75f691ea702bc6ade538781ab1) |
| `virtual_server.http3.tcp_server_profile.tenant` | [virtual_server.http3.tcp_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-f2e4364773e29463fc60c403faa267b59d93d2ae9b0225155696b0bed8053cf1) |
| `virtual_server.http3.tcp_server_profile.uid` | [virtual_server.http3.tcp_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-85b5bc2554f633dbdff6f946f28292372241df10f54951e7899009f972f07a26) |
| `virtual_server.http3.udp_client_profile` | [virtual_server.http3.udp_client_profile](resources--application_profiles--reference--group-002.md#canonical-71bad483fd1c0bdbe4b7988a7b8da92f524f4b1fde5dba2ea1284b6e1805a4c5) |
| `virtual_server.http3.udp_client_profile.kind` | [virtual_server.http3.udp_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-c7ad6412c7055d7b6d092d6b052cd2fce743c27a167139b9ec1f572b639ad911) |
| `virtual_server.http3.udp_client_profile.name` | [virtual_server.http3.udp_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-59e08dbe54b773c9f3909eb0928fba3449212b95ed0512695fb613857adced00) |
| `virtual_server.http3.udp_client_profile.namespace` | [virtual_server.http3.udp_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-ba4a35764fdbbc5c83d9c6251c0bbe67845d82fffcef0e6f165e381f4c800574) |
| `virtual_server.http3.udp_client_profile.tenant` | [virtual_server.http3.udp_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-53152a76783bd6dae7bc9bd42baf6a0b9822c1798d7797582ac6ae02e4bdbebc) |
| `virtual_server.http3.udp_client_profile.uid` | [virtual_server.http3.udp_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-452d77c134db57604ea49d2c278c68e82e7775e20c84ed83f5958d0ee289fca6) |
| `virtual_server.http3.udp_server_profile` | [virtual_server.http3.udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-91ce9335ff70b3b99fa637571a1094c1c6e8c27b4961a023375409aa0e299f38) |
| `virtual_server.http3.udp_server_profile.kind` | [virtual_server.http3.udp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-00769b48749abda73d1de11475bf273bc7837e1deb61e19e586172c669bea849) |
| `virtual_server.http3.udp_server_profile.name` | [virtual_server.http3.udp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-b0b3eb554762f4c6239e2f806e039b15a6d122e3812c0c65d88fccbb3c6a6f0f) |
| `virtual_server.http3.udp_server_profile.namespace` | [virtual_server.http3.udp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0301e122e4dfa173b4ef80b43ac3444e8cd7fdd489e98c70e3d7001ac98fe5bb) |
| `virtual_server.http3.udp_server_profile.tenant` | [virtual_server.http3.udp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-aadec0fd8f3813c538fe0ff35c687e890897687780c5a61cdaa28fc486c86cb8) |
| `virtual_server.http3.udp_server_profile.uid` | [virtual_server.http3.udp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-231443a49da8ace5e9a63eabecc3c84144c9925f146d4fc2ae015f4006b8ba60) |
| `virtual_server.https` | [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-c12f767912d5f44f0b3184d9bc70757f9b5d715c229a3b8a5bc245574b61efea) |
| `virtual_server.https.client_ssl_profile` | [virtual_server.https.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-346592f022e0465aac97e141e23a6da552e8014118356b040391f2bbfce06a2d) |
| `virtual_server.https.client_ssl_profile.kind` | [virtual_server.https.client_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-ceb5453b4fe4933a1bc7305cebf1c205e961e21e6ba4236ba413880ef71d6c44) |
| `virtual_server.https.client_ssl_profile.name` | [virtual_server.https.client_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-030f102a31424332ded3239128e3cc41c9a049bfddf95d4228e82d2bcea0b825) |
| `virtual_server.https.client_ssl_profile.namespace` | [virtual_server.https.client_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-75ee85239f6c9daf8b38db97e307d1441e73c4af51c11a3d4d8f17afe112b424) |
| `virtual_server.https.client_ssl_profile.tenant` | [virtual_server.https.client_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-8a8b8c9b5feabe366cf80e9bca586afaa6e7cd3850e80bbba957cf940950b18c) |
| `virtual_server.https.client_ssl_profile.uid` | [virtual_server.https.client_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-abd120550a81fe102dc071773511f302e53a877b779f6b52fde282e6a4cae6ed) |
| `virtual_server.https.http2_client_profile` | [virtual_server.https.http2_client_profile](resources--application_profiles--reference--group-003.md#canonical-436f0da56064a67c4727b5d52442b880b80dd2b9aba4470690cb8e1a9fd56958) |
| `virtual_server.https.http2_client_profile.kind` | [virtual_server.https.http2_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-08939e498c5bea659814aaec35fc4cccbc6d16b0686d63ae5b28375f7d5e0f09) |
| `virtual_server.https.http2_client_profile.name` | [virtual_server.https.http2_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-da093a1f1d2d42b1396b2c9703d6fd3ffe38f3c74e117dfc3a1b9b3430ab483d) |
| `virtual_server.https.http2_client_profile.namespace` | [virtual_server.https.http2_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-65f7ded224eb0bf5cf08649759b5581892a1855d69148cc69275c97c1be5e060) |
| `virtual_server.https.http2_client_profile.tenant` | [virtual_server.https.http2_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-7ac3469a0378b3924fc576fb897fd1065d1e625abf2f051963cb192472f20e0f) |
| `virtual_server.https.http2_client_profile.uid` | [virtual_server.https.http2_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-53737348eb9297b48dbe4927a0b988fbc1ba6e9bdb944ef8720ccdb64123b6a1) |
| `virtual_server.https.http2_server_profile` | [virtual_server.https.http2_server_profile](resources--application_profiles--reference--group-003.md#canonical-71d4c299e04fca28e4bd687804cd9b6b37a97408d672141afa6c4ac842ced0a3) |
| `virtual_server.https.http2_server_profile.kind` | [virtual_server.https.http2_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-5796dbce2add113cbd542b582e834486ec8395d9783d7cf5656c8cdc1b3dd01b) |
| `virtual_server.https.http2_server_profile.name` | [virtual_server.https.http2_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-cbbcf21d79d5c549bd19398b07c5f80da5906f0afac2024609b446aec9b421ef) |
| `virtual_server.https.http2_server_profile.namespace` | [virtual_server.https.http2_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-b1e3e6c571dd8d7a009472f52b238df74234ab799fdca2e093492e1954d5b394) |
| `virtual_server.https.http2_server_profile.tenant` | [virtual_server.https.http2_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-38a478a384a7d503e37a0c95c8cf1bdce2b7b7018b1966ba08e148d648bb1503) |
| `virtual_server.https.http2_server_profile.uid` | [virtual_server.https.http2_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-13c530ac0dd187503d291ef1e5838038144392d8bf71783974a57473f0b987af) |
| `virtual_server.https.http_client_profile` | [virtual_server.https.http_client_profile](resources--application_profiles--reference--group-003.md#canonical-0eef7685571edbc66585e9e5580c3b8957f4d2fe839d7660339cc5d329d91163) |
| `virtual_server.https.http_client_profile.kind` | [virtual_server.https.http_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-8b5e4800205aa9d52d8dcc091d4c6c736f88ee59aae8661c7003b273939e9969) |
| `virtual_server.https.http_client_profile.name` | [virtual_server.https.http_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-33ee89c880a974ec2be065a4ec2436c4dd21417dc1a0c1b95d9713047db5fbaf) |
| `virtual_server.https.http_client_profile.namespace` | [virtual_server.https.http_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0abfd85db6ee4d3f6fc6ce86ede7a2b6d8403132c283b81328e465e931787f11) |
| `virtual_server.https.http_client_profile.tenant` | [virtual_server.https.http_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-89ddb871e37e5990782f46d225cac1303b188c289c7d3f4837e43aaf532838b8) |
| `virtual_server.https.http_client_profile.uid` | [virtual_server.https.http_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-af3d4d3db3887e57d6c294ca2bb9fa411cef2d8387403fe453e4fe2e3ab041b9) |
| `virtual_server.https.http_server_profile` | [virtual_server.https.http_server_profile](resources--application_profiles--reference--group-003.md#canonical-0a5c06e974d71222cfc2cda8d64f7623e86e342dafb61f1c879932ecb3483d86) |
| `virtual_server.https.http_server_profile.kind` | [virtual_server.https.http_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-e3991582b273e401b6e8534bdb7cdd6e0dfdd36f52714e2f024b11eea274ac22) |
| `virtual_server.https.http_server_profile.name` | [virtual_server.https.http_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-23e0d69506efbe7c2e2637267f1ef732e985ec1296f5a57878e9a60d73873468) |
| `virtual_server.https.http_server_profile.namespace` | [virtual_server.https.http_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-31d4473e87ad4f43b2addf04407d1d9fcf192e23777de39db0258ea7ff8fb775) |
| `virtual_server.https.http_server_profile.tenant` | [virtual_server.https.http_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-4fe9cffff49b64326a653a12c5e7306f8499771324ccff92819f86c11845138e) |
| `virtual_server.https.http_server_profile.uid` | [virtual_server.https.http_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-3e4adfa944cdcf115b07ef8fb2d5e9e5af4be2b4fea6cff44fef3bca0c56bb74) |
| `virtual_server.https.ocsp_profile` | [virtual_server.https.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-5e18bda1a9fbe6751673f7300838c82ff71f6ce67c2a7316b75f38651366c0a3) |
| `virtual_server.https.ocsp_profile.kind` | [virtual_server.https.ocsp_profile.kind](resources--application_profiles--reference--group-003.md#canonical-3dcf425ed83a0ac8ff11fe133fbee859bad0700d94ac8e5ca1f8312bb962d356) |
| `virtual_server.https.ocsp_profile.name` | [virtual_server.https.ocsp_profile.name](resources--application_profiles--reference--group-003.md#canonical-ecafa25f3abf950b937457ae2c64302ecffb7fff4bfaef4ab59a0a07c3ae5a25) |
| `virtual_server.https.ocsp_profile.namespace` | [virtual_server.https.ocsp_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-200150ae1dd4248747080e86f10f1c8ca45014bd2e2965e57822a9123b42ee76) |
| `virtual_server.https.ocsp_profile.tenant` | [virtual_server.https.ocsp_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-8aefd7bc18f5394b2fe911443561a33aab3fa094827832828faf9da885d6a023) |
| `virtual_server.https.ocsp_profile.uid` | [virtual_server.https.ocsp_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2116f561c3af4ed03122b1602bfd66fa58e71f472a179060b4f0c2d339c5a1f2) |
| `virtual_server.https.server_ssl_profile` | [virtual_server.https.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-7274ce87cb31a97f73e61974cba60f4bf10415aefae6a39cdd2881ce85c3e836) |
| `virtual_server.https.server_ssl_profile.kind` | [virtual_server.https.server_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-19f1db1a9e8e8d4ea39a99a26df65bf96e28d7ca84e5ad78f10ba9bc974dc68c) |
| `virtual_server.https.server_ssl_profile.name` | [virtual_server.https.server_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-c530a3ac6e0778ca486c987d8a16814383a104f86acaa43df2472d8208c3601c) |
| `virtual_server.https.server_ssl_profile.namespace` | [virtual_server.https.server_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-05a5e04656d7cf1e3a3722338295809e06ec3cebbb0c006365930ac60cfdfd58) |
| `virtual_server.https.server_ssl_profile.tenant` | [virtual_server.https.server_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-34f1a30d11ab5e024c51ef7bb3ddca98d5445cfd19df1be1781507f247395c9e) |
| `virtual_server.https.server_ssl_profile.uid` | [virtual_server.https.server_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-324613bde8ef8d8a046130fd3603dae3d1b4c7b4d7eb4dfaf4a9293f68ab32ae) |
| `virtual_server.https.stream_profile` | [virtual_server.https.stream_profile](resources--application_profiles--reference--group-003.md#canonical-7b7e9b1165a094bd3556f71afd636c6afb817abc42116819b88593d54deea18a) |
| `virtual_server.https.stream_profile.kind` | [virtual_server.https.stream_profile.kind](resources--application_profiles--reference--group-003.md#canonical-5098cebd6f266741e34cfa9436d44ec0c33df4670e1ec93e854e8aa9f001bd7c) |
| `virtual_server.https.stream_profile.name` | [virtual_server.https.stream_profile.name](resources--application_profiles--reference--group-003.md#canonical-e5f5393ba9fbfa2de14687c2c3f3c867fef9e21403b6c35985e51633ce835599) |
| `virtual_server.https.stream_profile.namespace` | [virtual_server.https.stream_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-d1771fd482dcebb3f940a55e6a13e4b112aab7293cf0638424f90ae04c73d25c) |
| `virtual_server.https.stream_profile.tenant` | [virtual_server.https.stream_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-b37ade79bbb8cf1a0ffcc042eddb412dd80a166f373cc0ebf3eacb3c15f8c72e) |
| `virtual_server.https.stream_profile.uid` | [virtual_server.https.stream_profile.uid](resources--application_profiles--reference--group-003.md#canonical-b094ac5875ec5b65be23438d8b8ff8c191cb79fc6a65b4f2b4105b83b16faf48) |
| `virtual_server.https.tcp_client_profile` | [virtual_server.https.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-f3231e170758e2607aa8b03298743fc33977baedd5b4d2e7f4e91354f57c0007) |
| `virtual_server.https.tcp_client_profile.kind` | [virtual_server.https.tcp_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-cec26f11880ce06ab9beb3010da2f6264f9715d41cc68cfa022ccf4bee1ab35e) |
| `virtual_server.https.tcp_client_profile.name` | [virtual_server.https.tcp_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-217c450d943ca852ce71e5fe397ec6c26564a697ac3eb2c98ed72d6cc4f79f13) |
| `virtual_server.https.tcp_client_profile.namespace` | [virtual_server.https.tcp_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-83a5c58c9be0974b41f3d532b51d1f348f2289de52812dac11658ee7b067211e) |
| `virtual_server.https.tcp_client_profile.tenant` | [virtual_server.https.tcp_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-ca81222692db9467fe6a7630cc84f9f473b4d07bdb95680a13cf9ff9fd940670) |
| `virtual_server.https.tcp_client_profile.uid` | [virtual_server.https.tcp_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-5ef0790081802244a3299e17392289902ce36a60c888fcdfe869f54b0dc6f0cf) |
| `virtual_server.https.tcp_server_profile` | [virtual_server.https.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-a4f8d051ce5f5ab415b4312e3faa5f4adca3cc353b6ff2e7dab0d73f3ac7c222) |
| `virtual_server.https.tcp_server_profile.kind` | [virtual_server.https.tcp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-9b4788bf6cdd5603f99f0bc33f248a771ac4aa4985cb553cb450cdbd62f25f2d) |
| `virtual_server.https.tcp_server_profile.name` | [virtual_server.https.tcp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-659fe76bf938d6b572b196c9547b7b85c9e39f0f2c031169cb3a47e80f79ac12) |
| `virtual_server.https.tcp_server_profile.namespace` | [virtual_server.https.tcp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2b11f28303d27b9c834ffced3c02517e7d50f82d9793ae4485046a78a0f3497a) |
| `virtual_server.https.tcp_server_profile.tenant` | [virtual_server.https.tcp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-fec493f5811dd05474d3dedf3fdff55348b10bb4e72d96536f079f1c9a2d7f8f) |
| `virtual_server.https.tcp_server_profile.uid` | [virtual_server.https.tcp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-b3d34e7e850d6c00ab6cfeac09bd7b225556766340217145708fb26f113c92ca) |
| `virtual_server.https.websocket_client_profile` | [virtual_server.https.websocket_client_profile](resources--application_profiles--reference--group-003.md#canonical-3979bc2e7735f3e2d8cc8e2793517d62d1c16c29b2b540fbd5fa60bc2714d176) |
| `virtual_server.https.websocket_client_profile.kind` | [virtual_server.https.websocket_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-6c4536a7a8e99e55d40511a2270fe2c901e616722d1dad28feb56102a371ad72) |
| `virtual_server.https.websocket_client_profile.name` | [virtual_server.https.websocket_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-cf4a4344a849f6881dd488b04ed0009d807995da7d50b6339ea2f2a436f54224) |
| `virtual_server.https.websocket_client_profile.namespace` | [virtual_server.https.websocket_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-6bc98672809a8db34704b2757caec71c7059ac0348a2eb1cd4fd45c124d39421) |
| `virtual_server.https.websocket_client_profile.tenant` | [virtual_server.https.websocket_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0dbccd77ef194e199f2298cb39a3043cc01176576d7ebe03dde33ea4dfc27c90) |
| `virtual_server.https.websocket_client_profile.uid` | [virtual_server.https.websocket_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-22427b8a9609c50e6fe35670f0d868feba5378682e493a2598dca4cf3204cc15) |
| `virtual_server.https.websocket_server_profile` | [virtual_server.https.websocket_server_profile](resources--application_profiles--reference--group-003.md#canonical-654f3287c560d8b6f61345373d9da715135d8e42102e4b8ae75c4734f39313ef) |
| `virtual_server.https.websocket_server_profile.kind` | [virtual_server.https.websocket_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-8279ec2c21d1e98f7a58b976cb7034130eb0b1f2e273a7e10ecd6717e8940e2e) |
| `virtual_server.https.websocket_server_profile.name` | [virtual_server.https.websocket_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-d346a6c528997cbd2339e786e76e516885bcfb909d2969805b84d146b3b83934) |
| `virtual_server.https.websocket_server_profile.namespace` | [virtual_server.https.websocket_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-270afc9d0097ea62f1ac0ac1feac6f00f6093e6fcf689c31213b48f53e32289f) |
| `virtual_server.https.websocket_server_profile.tenant` | [virtual_server.https.websocket_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-4b7405f33c3baf66e4965fad838bc643f10145eca86904a8bc86801dc5646706) |
| `virtual_server.https.websocket_server_profile.uid` | [virtual_server.https.websocket_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-ddaf3e7f06a75149cd005ce515dc8efa4c214a0edfa19d0fada3f20ab6134382) |
| `virtual_server.immediate_action_on_service_down` | [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-ec9cea6a07fdf90f6a787dd6ba4045377cddd5942c66d1842e3c1e9f3489e62d) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](resources--application_profiles--reference--group-003.md#canonical-691f984d154b0ab1f511d9cc01826c7e202854dc6c33ac416448f2233412ebc6) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](resources--application_profiles--reference--group-003.md#canonical-8374076fd0af922f614197f2e08f815516788b0188f7785601ae47d68d6bace3) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](resources--application_profiles--reference--group-003.md#canonical-f294896154e9c1112b7c80c75beba8c9f6ede37ce021ecb600943fd6cfb325a3) |
| `virtual_server.last_hop_pool` | [virtual_server.last_hop_pool](resources--application_profiles--reference--group-003.md#canonical-dc4369260886abbb12afc2c793664126983cbba3b6bfd4436061a7e5ee3bc0d4) |
| `virtual_server.last_hop_pool.kind` | [virtual_server.last_hop_pool.kind](resources--application_profiles--reference--group-003.md#canonical-1fa3e5a3988fd2f3134e105008d323777a679a94abb39220bbd12f5c94e8445f) |
| `virtual_server.last_hop_pool.name` | [virtual_server.last_hop_pool.name](resources--application_profiles--reference--group-003.md#canonical-1d6c072281ff1cf4f47d7da6a5f7160dc9221fe64b0f4139d2ef411773cb936c) |
| `virtual_server.last_hop_pool.namespace` | [virtual_server.last_hop_pool.namespace](resources--application_profiles--reference--group-003.md#canonical-d4a04fff1e52e838bfc7b1a4460b4e0a64c88cdd3616ef52e38715eca713724c) |
| `virtual_server.last_hop_pool.tenant` | [virtual_server.last_hop_pool.tenant](resources--application_profiles--reference--group-003.md#canonical-48661e5539ab3a35cda180df4bbec286897faba5c271c7e4371d1b4d7462bacd) |
| `virtual_server.last_hop_pool.uid` | [virtual_server.last_hop_pool.uid](resources--application_profiles--reference--group-003.md#canonical-83307141818367bcc6adacd693a9b05fe6f4937d6a1304bc1c4a0bdc10afafa7) |
| `virtual_server.nat64` | [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-4346f65ce486e2d6f472232980a36fdd53c585a2ae5144e5c908fc0eecaaae0c) |
| `virtual_server.nat64.nat64_disable` | [virtual_server.nat64.nat64_disable](resources--application_profiles--reference--group-003.md#canonical-0e82a5dbebd31e94d2ef4c793be3161ee7c6abbf73ee30316c0b9c5720ad0793) |
| `virtual_server.nat64.nat64_enable` | [virtual_server.nat64.nat64_enable](resources--application_profiles--reference--group-003.md#canonical-e7d81cb31bc12f4185be25e6cb7852352a6cc6d54beb514343c78efe6d81aefc) |
| `virtual_server.port_translation` | [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-f095dcca27f2dc833b94bad6297a8a3659fd99e49ff6e6be04c31bb61a640999) |
| `virtual_server.port_translation.port_translation_disable` | [virtual_server.port_translation.port_translation_disable](resources--application_profiles--reference--group-003.md#canonical-3d87c5a8fa021cabf2ee51f4452afc39128f1d04315128e02a7698771c256c2c) |
| `virtual_server.port_translation.port_translation_enable` | [virtual_server.port_translation.port_translation_enable](resources--application_profiles--reference--group-003.md#canonical-9493a6c550b1ee2511d94f2104de2634b2c317adf3804f14486107eb3c16fb49) |
| `virtual_server.request_logging_profile` | [virtual_server.request_logging_profile](resources--application_profiles--reference--group-003.md#canonical-c94395b16fc46d6036f0b56081c48cf5afc9c78d972d11d51f26f9fbeedbc9cb) |
| `virtual_server.request_logging_profile.kind` | [virtual_server.request_logging_profile.kind](resources--application_profiles--reference--group-003.md#canonical-afb8c2f7df8bb51dddeae68a486f983fb995effc3c91a880aebd9236597f6df9) |
| `virtual_server.request_logging_profile.name` | [virtual_server.request_logging_profile.name](resources--application_profiles--reference--group-003.md#canonical-bc5da30fffd28d7e2b2519f37b4f9a59e197667f0f391774be2ac7cbfaade0e8) |
| `virtual_server.request_logging_profile.namespace` | [virtual_server.request_logging_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2f6134327585451158f6e06400669c19981cd0d127f6f9ba04eac6bfdb6fe1a4) |
| `virtual_server.request_logging_profile.tenant` | [virtual_server.request_logging_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-234a5addd3487780dc6e4cf590d07e8b83e708b211079bb22e8f9400b41d4f01) |
| `virtual_server.request_logging_profile.uid` | [virtual_server.request_logging_profile.uid](resources--application_profiles--reference--group-003.md#canonical-4d23b12874549cdd6573161418cf2f36318acf941a955208bb4977fc289cf9ad) |
| `virtual_server.source_port` | [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-a613eadd95ee70e9da9f32f488fddebb138504424812e277158e224aab781e27) |
| `virtual_server.source_port.source_port_change` | [virtual_server.source_port.source_port_change](resources--application_profiles--reference--group-003.md#canonical-75f3b24ea41ffff27dc4372303fe8840ddf46180873b274cb3897683fb5e15b8) |
| `virtual_server.source_port.source_port_preserve` | [virtual_server.source_port.source_port_preserve](resources--application_profiles--reference--group-003.md#canonical-4adc65163883520b628ec14c48429bb07a04b018dbdeb1a633acca5b5265e1fa) |
| `virtual_server.source_port.source_port_preserve_strict` | [virtual_server.source_port.source_port_preserve_strict](resources--application_profiles--reference--group-003.md#canonical-a654341548db2fd370475774355a37928e746b675b49ad0db436c1311c08d1b0) |
| `virtual_server.statistics_profile` | [virtual_server.statistics_profile](resources--application_profiles--reference--group-003.md#canonical-8f35735e0b193c1f57d5ef898184c4aed550735b7acf117ec911e8dad9b30fc0) |
| `virtual_server.statistics_profile.kind` | [virtual_server.statistics_profile.kind](resources--application_profiles--reference--group-003.md#canonical-432884df9d55f0fd3c3d8833a132033ab3e420a5228d387ab4f582f08a8a48cd) |
| `virtual_server.statistics_profile.name` | [virtual_server.statistics_profile.name](resources--application_profiles--reference--group-003.md#canonical-70f6775064b6be64084d9ff6d172a12ede8a5d51c41757da0d47436772347dfb) |
| `virtual_server.statistics_profile.namespace` | [virtual_server.statistics_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-ebef21e9a6a78d5710d8a0913ce3e46110be1e894cd96d2e05294dc9f02e582b) |
| `virtual_server.statistics_profile.tenant` | [virtual_server.statistics_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-a97d56edd44712d6ed0c70069ab11533e61eca0caeb8f1a86bbd620d0233fd59) |
| `virtual_server.statistics_profile.uid` | [virtual_server.statistics_profile.uid](resources--application_profiles--reference--group-003.md#canonical-1ee1a0e52db5afbd43137d279bc5ac6d3afa9d0a236e56e86a3727e960f5a2d4) |
| `virtual_server.tcp` | [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-7491386678c5e86f64807b6b51b0179bf23ee15ebee76a46bb7a9f661ceeebf1) |
| `virtual_server.tcp.client_ssl_profile` | [virtual_server.tcp.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-46f674b5b5a3857a1beefa61338dc0a28019c5051d7c8de2956b94f67ed194f7) |
| `virtual_server.tcp.client_ssl_profile.kind` | [virtual_server.tcp.client_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-6f1b3f4271c2b70864494fc4614ce3ec4cb21b7dbcb39bea9bde0d32c481cf61) |
| `virtual_server.tcp.client_ssl_profile.name` | [virtual_server.tcp.client_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-f627ea79a430a4716b4af5d374a19886224d3cba388efd26ba9abe8f6f33b339) |
| `virtual_server.tcp.client_ssl_profile.namespace` | [virtual_server.tcp.client_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-ad2582b2cd63a1a8db8ac2e261bcac46d666015bfdadd2e26cd3bb2ee8134955) |
| `virtual_server.tcp.client_ssl_profile.tenant` | [virtual_server.tcp.client_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-d898654e474efa846303d9945f48dbb5d9aeedff2fa73b9f3bb7d52861ac9027) |
| `virtual_server.tcp.client_ssl_profile.uid` | [virtual_server.tcp.client_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-59bc497dcda88b59b601a1d468a7242ce0e47a36fd83c786c1d02d77fa6bb78f) |
| `virtual_server.tcp.ocsp_profile` | [virtual_server.tcp.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-a1f12e75d7498d531e030c1baef6d5890cd6e3fc4cfe8f5faf9a98c0a8247e9c) |
| `virtual_server.tcp.ocsp_profile.kind` | [virtual_server.tcp.ocsp_profile.kind](resources--application_profiles--reference--group-003.md#canonical-6149962a4aab398e3a8f0be1b968ef657635b8b0db7490510ddb6129cedf0702) |
| `virtual_server.tcp.ocsp_profile.name` | [virtual_server.tcp.ocsp_profile.name](resources--application_profiles--reference--group-003.md#canonical-e80226f19fd88d9aa705d679b06b98f57f72f867632d151f35708ca09fa23494) |
| `virtual_server.tcp.ocsp_profile.namespace` | [virtual_server.tcp.ocsp_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-4cf1d22dc6ca64b21daafa29b7a68f2457fec62beb64ece4caa06a47c9be6d75) |
| `virtual_server.tcp.ocsp_profile.tenant` | [virtual_server.tcp.ocsp_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0eb03436592e5cadd8358ae0a6296ff798721d215e4ab5bc33c7e0b50188b0fb) |
| `virtual_server.tcp.ocsp_profile.uid` | [virtual_server.tcp.ocsp_profile.uid](resources--application_profiles--reference--group-003.md#canonical-a92d86e4c73d88906a8edb8524bb4b23f980c7e0902bad20d0c68644582876f5) |
| `virtual_server.tcp.server_ssl_profile` | [virtual_server.tcp.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-37f95f02b022cd2dad04acf92ff97a23ed1d35099dddfbd25f50738edca4daf2) |
| `virtual_server.tcp.server_ssl_profile.kind` | [virtual_server.tcp.server_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-c0c673082eaf892b508cb555a06fcb1b285d76880feb26a0488d774d8fef8a5b) |
| `virtual_server.tcp.server_ssl_profile.name` | [virtual_server.tcp.server_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-3d175c13b323714b3bd5d07eae19bfa66746c1d4a3b86953365892a49d8bbf56) |
| `virtual_server.tcp.server_ssl_profile.namespace` | [virtual_server.tcp.server_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-a5de5ab96071fd4179ae36f6ba1c60cc91e0d3afc975c154f8b6350f9479c90f) |
| `virtual_server.tcp.server_ssl_profile.tenant` | [virtual_server.tcp.server_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-38cff4667005b8e658302ecdce3f85c2ae37f7834f2963fa701ddd8b64091fa3) |
| `virtual_server.tcp.server_ssl_profile.uid` | [virtual_server.tcp.server_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-b1a0b3fc8b7789c39db27e3570a9795d56f5a65ca532833f96cba851e1c6bfe9) |
| `virtual_server.tcp.tcp_client_profile` | [virtual_server.tcp.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-bb16b3cac76eed33a02eff1c9418eeb8a503a0fb5ef1730a64511eb885792ab0) |
| `virtual_server.tcp.tcp_client_profile.kind` | [virtual_server.tcp.tcp_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-62fdf1eb39c2e7b0fafbeafaf4424b80b701bdca1f34634c153c15cb7f2df59f) |
| `virtual_server.tcp.tcp_client_profile.name` | [virtual_server.tcp.tcp_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-b7e62883a54f8355789f45c16d8766d61f3c75e7c1dfc75fa00431e3afdbded8) |
| `virtual_server.tcp.tcp_client_profile.namespace` | [virtual_server.tcp.tcp_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-95618de0ea44a30471207585ad63b4d884e09cfcca5f7f912cce37ddadee2136) |
| `virtual_server.tcp.tcp_client_profile.tenant` | [virtual_server.tcp.tcp_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-4d5787a320183948c94ff7d3a49b84474013b2f3801154512c57d89bed6063c7) |
| `virtual_server.tcp.tcp_client_profile.uid` | [virtual_server.tcp.tcp_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-8d1fc82b1f3a8c9feb92bba254f6864927a41e4f16225623eef7482858be6e3d) |
| `virtual_server.tcp.tcp_server_profile` | [virtual_server.tcp.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-da4955a924121eed898094164813bd0e45351a385500691b04f1cdab6dba869f) |
| `virtual_server.tcp.tcp_server_profile.kind` | [virtual_server.tcp.tcp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-73031e72fb2d73419d5458885717ae51d1db154590308370a9ab51020c6e79f2) |
| `virtual_server.tcp.tcp_server_profile.name` | [virtual_server.tcp.tcp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-440a9eada095842aebeea3900c00f6e361196375619c513f2a3f0e388e610e7c) |
| `virtual_server.tcp.tcp_server_profile.namespace` | [virtual_server.tcp.tcp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-987a403b9295cda36b8894219d7c2cf9e36d7a46a7dde05f527a1b824938852c) |
| `virtual_server.tcp.tcp_server_profile.tenant` | [virtual_server.tcp.tcp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-8c730f605f4c9b733610fbd87ee71b308252c5a6842b20361af48122cd38415f) |
| `virtual_server.tcp.tcp_server_profile.uid` | [virtual_server.tcp.tcp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-eaab5ff199a86544424a598eb93b5c8bb67b1269872d36838bf2fde9e2dbdc65) |
| `virtual_server.udp` | [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-22d0d88487b3c3e8e360cc38efd40f86bf000b27c5b177b9b53f087cb8ef72f3) |
| `virtual_server.udp.client_ssl_profile` | [virtual_server.udp.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-cf36b72ed7adc83c897fa39d9928abef14c1d9fe699c42548fb0f44f710d390e) |
| `virtual_server.udp.client_ssl_profile.kind` | [virtual_server.udp.client_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-16a7cf0f1b42af170076b3edda04fac7b7767653cebfe5251609fcb506f9a666) |
| `virtual_server.udp.client_ssl_profile.name` | [virtual_server.udp.client_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-f40d5d851453282c74de895578529bbe6e7f61e24ee6645def081fec033a002c) |
| `virtual_server.udp.client_ssl_profile.namespace` | [virtual_server.udp.client_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-a324fee4dd00c609fee1aa074ebfb80d5be9e727ee4e49883656aac1d3155444) |
| `virtual_server.udp.client_ssl_profile.tenant` | [virtual_server.udp.client_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-f3b30e0bb6944a1cd826fa36b54e7e8138b77c97e05234526ca1179ed583bf02) |
| `virtual_server.udp.client_ssl_profile.uid` | [virtual_server.udp.client_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-49b36101fbabc1eb3687c7e634c34f3b5c686a18c75ded5ab8032a084ae2e883) |
| `virtual_server.udp.server_ssl_profile` | [virtual_server.udp.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-c30e16a22b12eecc18145b78a48be5565b3152522d619f01634c8f5b030f3eb2) |
| `virtual_server.udp.server_ssl_profile.kind` | [virtual_server.udp.server_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-fc151e7f926a7fbafd6e3d067861f93c9d5e43df09dfa70c89a0080aa5268393) |
| `virtual_server.udp.server_ssl_profile.name` | [virtual_server.udp.server_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-77e98811014058dcaa2cdf867794552c90e8aeeeb0d4e20addc5939ff29d8705) |
| `virtual_server.udp.server_ssl_profile.namespace` | [virtual_server.udp.server_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-690d8ab3976fbf1658b8a3dba50eae296060ac558b2469597a2c51e10adf1117) |
| `virtual_server.udp.server_ssl_profile.tenant` | [virtual_server.udp.server_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2a64cae754748a3efbea5e02dd3ff2951eb2a7c893ccbea8ca242b5af216485d) |
| `virtual_server.udp.server_ssl_profile.uid` | [virtual_server.udp.server_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-c77d25a478656bc55f734b9f4ee2524bf024af0be94e911b521d71f3c3806c28) |
| `virtual_server.udp.udp_client_profile` | [virtual_server.udp.udp_client_profile](resources--application_profiles--reference--group-003.md#canonical-d2baaecdbb6d1a58ac2ce6534560657553e415aa317f5bd427436a891a1f557c) |
| `virtual_server.udp.udp_client_profile.kind` | [virtual_server.udp.udp_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-8c332d1ae54035a06b8d61cc57d8fe139bf3df332541833d2d96499d57b14043) |
| `virtual_server.udp.udp_client_profile.name` | [virtual_server.udp.udp_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-d27a75fc5003afbe27b487c907de84a3ff6e2ccb03754221536ed611c657fde5) |
| `virtual_server.udp.udp_client_profile.namespace` | [virtual_server.udp.udp_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-b9549f576e68f1d0ff83fa2401e929207dfd56dc8f18230d784a28a33b205d86) |
| `virtual_server.udp.udp_client_profile.tenant` | [virtual_server.udp.udp_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0e43f8532805921fa9e0fb36a483b58b8161c32732a970c50bd3162cf933c64a) |
| `virtual_server.udp.udp_client_profile.uid` | [virtual_server.udp.udp_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-7019f2e20d4ec4d7886ba091d0408788af586f3321fa03e29c55a6f1312d4a13) |
| `virtual_server.udp.udp_server_profile` | [virtual_server.udp.udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-2c6e3a7a3c1a9b7639513be87575dccd26e8ce6e41089edb7fedcef1339c6ec5) |
| `virtual_server.udp.udp_server_profile.kind` | [virtual_server.udp.udp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-ffacc8f4ca1e60d368672b248bdc71b616122230320839f9b5607d1acbd82168) |
| `virtual_server.udp.udp_server_profile.name` | [virtual_server.udp.udp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-980d55c1edcb03c796936c7accacf2ef901d89cad5b3d1bd8a2021970ff6ed13) |
| `virtual_server.udp.udp_server_profile.namespace` | [virtual_server.udp.udp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-012c65d41fc10bf5a36ac1ab00a095a4c2588dcdba93f935b15d0ff763c41343) |
| `virtual_server.udp.udp_server_profile.tenant` | [virtual_server.udp.udp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2dd5d48f60ae83deb1b0095c79ba0fa20cb3423b86379dc2d7d028688a1f8d0a) |
| `virtual_server.udp.udp_server_profile.uid` | [virtual_server.udp.udp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-10cff189a038bbf899d9907a0ddcb9ac6216d503f13f5caef8fe8deb86ea117e) |
| `virtual_server.virtual_server_state` | [virtual_server.virtual_server_state](resources--application_profiles--reference--group-003.md#canonical-1acda12d7935f75c62c8c4fdd53ef74bd1232373ca5bb82f83205675a70218fa) |
| `virtual_server.virtual_server_state.state_disabled` | [virtual_server.virtual_server_state.state_disabled](resources--application_profiles--reference--group-004.md#canonical-18c5c0192391a95f49791b7311f8abc99f33c9f4f1f2eabd2234ab3059eb1d7b) |
| `virtual_server.virtual_server_state.state_enabled` | [virtual_server.virtual_server_state.state_enabled](resources--application_profiles--reference--group-004.md#canonical-8c829cd1611df0305ed4c031de1b0f209f2d33a08eb8c1c67e55a4e6e15d9c91) |
| `virtual_server.vs_score` | [virtual_server.vs_score](resources--application_profiles--reference--group-001.md#canonical-9e75536c70da03cb2dc957c89c7787e627a6533987623c9b7c52dee36235fe72) |

<a id="canonical-af6c41fe29e43e3e788fdf99875c3d766e24613277de7f1ad5e08f0a7cfb6b3e"></a>

## Next pages — Property reference / 9580f5c2d307 / 12

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-1219c9680f1cb17608da1bff14b3321209c387abae26d7ec1cfd56a03083cefa)
- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-a2f8dadbb1c84e46454b50aa5033edcdf5f94245ad7c84259a87b214799bc630)
- [irules](resources--application_profiles--reference--group-001.md#canonical-3f967a8395804c6c158083e0546c4f390cfd910da7685ac8449c86b36c6d564f)
- [timeouts](resources--application_profiles--reference--group-001.md#canonical-90d932bd8f79925ec5d6b9945630df17fac61100dbca671694a7e9d6a95b4b80)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-1219c9680f1cb17608da1bff14b3321209c387abae26d7ec1cfd56a03083cefa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d02aae2fbb19f3ddc72fe708d11dc8e84c605ec0ea8e6c2203789b75d1ff212"></a>

## advanced_tcp_profile — advanced_tcp_profile / 0d4a65f50336 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- advanced_tcp_profile

<a id="canonical-1d8de3942644bd448e525f5b3e8973c3c2acbe1a4131c05656735be246d65fe8"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for advanced tcp profile.

Upstream description:

BIG-IP Advanced TCP Profile.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_tcp_advanced_profile",
    "enable_tcp_advanced_profile")}
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
  "x-ves-oneof-field-tcp_advanced_profile_choice": "[\"disable_tcp_advanced_profile\",\"enable_tcp_advanced_profile\"]"
}
```

Terraform syntax:

```terraform
advanced_tcp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0bc23005260cac3e84981ad00b524b69b97ce4bb15c2fb7d0b956937d8300f1b"></a>

## Direct properties — advanced_tcp_profile / 0d4a65f50336 / 3

- [disable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-21a721d08773a73860cd75aa768fcf528dee9bd000258c15f3f6fabafcd5317c): complete subsection reference.

- [enable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-10606cc54ebebfd6a1d7a83a1e9cca943a77b6a604a414375fcf759f6a112f32): complete subsection reference.

<a id="canonical-b70ded10700ccbc2168deeafade0586e584074d3502247c060f01233bfb324b2"></a>

## Next pages — advanced_tcp_profile / 0d4a65f50336 / 4

- [advanced_tcp_profile.disable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-21a721d08773a73860cd75aa768fcf528dee9bd000258c15f3f6fabafcd5317c)
- [advanced_tcp_profile.enable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-10606cc54ebebfd6a1d7a83a1e9cca943a77b6a604a414375fcf759f6a112f32)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-21a721d08773a73860cd75aa768fcf528dee9bd000258c15f3f6fabafcd5317c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63f72bfde7262221c5b73cf7c70381e64028c7d2b0d914a4237eb1de0a1a96fa"></a>

## advanced_tcp_profile.disable_tcp_advanced_profile — advanced_tcp_profile.disable_tcp_advanced_profile / 77e27ce1ed95 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-1219c9680f1cb17608da1bff14b3321209c387abae26d7ec1cfd56a03083cefa)
- advanced_tcp_profile.disable_tcp_advanced_profile

<a id="canonical-a7aa6aadbea989fa9c84e8755cf9319f534fce2ae068c6ce8eb2f81c9cb26cae"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tcp advanced profile.

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
disable_tcp_advanced_profile = {}
```

<a id="canonical-e0461a6caae948e73975ed70efb2c9f11ed798f9a3902681dbf2882b242a37a0"></a>

## Direct properties — advanced_tcp_profile.disable_tcp_advanced_profile / 77e27ce1ed95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42d88764e11951189513eb103fa1204b7d8dfcd2c23777f2e59d08b8a121474d"></a>

## Next pages — advanced_tcp_profile.disable_tcp_advanced_profile / 77e27ce1ed95 / 4

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-1219c9680f1cb17608da1bff14b3321209c387abae26d7ec1cfd56a03083cefa)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-10606cc54ebebfd6a1d7a83a1e9cca943a77b6a604a414375fcf759f6a112f32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e73a53b1a984c2b1bfabd7a4add0931a400b1a622792fcb0bdef4eb16e5c52e9"></a>

## advanced_tcp_profile.enable_tcp_advanced_profile — advanced_tcp_profile.enable_tcp_advanced_profile / 1e231544999f / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-1219c9680f1cb17608da1bff14b3321209c387abae26d7ec1cfd56a03083cefa)
- advanced_tcp_profile.enable_tcp_advanced_profile

<a id="canonical-d555dd2392205dab61e67e011a61deadb475f83d0601eaa2af474e9f3628cc7e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tcp advanced profile.

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
enable_tcp_advanced_profile = {}
```

<a id="canonical-7b1aa14fb5c70d8a9c42891cf664d72ded70e3798999daa8573c0a81151ad47d"></a>

## Direct properties — advanced_tcp_profile.enable_tcp_advanced_profile / 1e231544999f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-365a715edbaad6e079422e0c0372ab6012eb673a16697215058ec512de1df873"></a>

## Next pages — advanced_tcp_profile.enable_tcp_advanced_profile / 1e231544999f / 4

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-1219c9680f1cb17608da1bff14b3321209c387abae26d7ec1cfd56a03083cefa)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-a2f8dadbb1c84e46454b50aa5033edcdf5f94245ad7c84259a87b214799bc630"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f0551744058350ae881175fbdda7b3e9a0ff309d874343d91dc70d840816fd4"></a>

## ddos_profile — ddos_profile / a5def6f1199d / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- ddos_profile

<a id="canonical-6be19b43f9ccd108b6cbe05345c64ef57df0553c4e165e968ee557232295110c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Upstream description:

BIG-IP DDoS Protection Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
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
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5b0b802dc25fd1a2e0d174576b1c7948f6b3dcdda05e7513a98c3dd6b54a3f9"></a>

## Direct properties — ddos_profile / a5def6f1199d / 3

- [disable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-a907a9ca539795d14c5324019b610a65e3753cfd43824313b0eab3182621dc29): complete subsection reference.

- [enable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-c6974194f497a7db4435d353cb915d82c96af31c875e0a859b6d454144554a4a): complete subsection reference.

<a id="canonical-6beb82a46f8e01e6d8995d05b941b9011455d07e1869003a705e60b029918bec"></a>

## Next pages — ddos_profile / a5def6f1199d / 4

- [ddos_profile.disable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-a907a9ca539795d14c5324019b610a65e3753cfd43824313b0eab3182621dc29)
- [ddos_profile.enable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-c6974194f497a7db4435d353cb915d82c96af31c875e0a859b6d454144554a4a)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-a907a9ca539795d14c5324019b610a65e3753cfd43824313b0eab3182621dc29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-807aaa0a8171a43516bd38b9a5d05bcee4e3389cc134f4e38b6d11ec1bd1ae72"></a>

## ddos_profile.disable_ddos_mitigation — ddos_profile.disable_ddos_mitigation / c3903598fd6b / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-a2f8dadbb1c84e46454b50aa5033edcdf5f94245ad7c84259a87b214799bc630)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-79e5240a52fe43be5885723bdb8963780455b31fc307e57613a2b771fd99ed7c"></a>

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
disable_ddos_mitigation = {}
```

<a id="canonical-4600624e558ada422899fe174b128dafb0bb29f926416fafc6ae2e2b6e093dcc"></a>

## Direct properties — ddos_profile.disable_ddos_mitigation / c3903598fd6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71b0aef764c342fe186adcfc8a3f6f16fda7f6325d03b175af5532516be7eaae"></a>

## Next pages — ddos_profile.disable_ddos_mitigation / c3903598fd6b / 4

- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-a2f8dadbb1c84e46454b50aa5033edcdf5f94245ad7c84259a87b214799bc630)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-c6974194f497a7db4435d353cb915d82c96af31c875e0a859b6d454144554a4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff5731f6a48f14d0f9359170bd1469631ac3b1a6aa5ecfc3e7e3f000a58eab90"></a>

## ddos_profile.enable_ddos_mitigation — ddos_profile.enable_ddos_mitigation / db8a456f9cd2 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-a2f8dadbb1c84e46454b50aa5033edcdf5f94245ad7c84259a87b214799bc630)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-6849bab2862656d1d4ec98685734fabc09cf6ac33dd9ff2284663a83ae915455"></a>

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
enable_ddos_mitigation = {}
```

<a id="canonical-bb6d72287fd19a8f9e51918458fac147bf328693da5483f0a83e555957830b34"></a>

## Direct properties — ddos_profile.enable_ddos_mitigation / db8a456f9cd2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ad570c41b87e0677f721f7e140c1362dfc6b01b0ffbffbf86f4ac5dc5bf4c6a"></a>

## Next pages — ddos_profile.enable_ddos_mitigation / db8a456f9cd2 / 4

- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-a2f8dadbb1c84e46454b50aa5033edcdf5f94245ad7c84259a87b214799bc630)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-3f967a8395804c6c158083e0546c4f390cfd910da7685ac8449c86b36c6d564f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caebaa3bd6aa3a1484487b4cf1bdc7e4db60a428cdc7522c09243109567fd375"></a>

## irules — irules / 74662a2151fe / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- irules

<a id="canonical-e1db3b24fc664083e1a28057e5d664756b66cc8cbfd4b08ca336a4c2332248b2"></a>

Type: `"object"`. list nested block, Optional.

OPTIONS for attaching iRules to BIG-IP Proxy.

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
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-bff3d14eacec89f67fc2cb50966f76b9a702ce55c3979ab7d50cab7737e470b4"></a>

## Direct properties — irules / 74662a2151fe / 3

<a id="canonical-27c84041f7c637c4b3bfa8f5b1c2b8b2d871558f7fa402d21b1a2eb84583134b"></a>

<a id="canonical-82e1ba40108fd0d911622db96becb9209fd4358df23e2c17bf186e4edb535362"></a>

## kind property — irules / 74662a2151fe / 4

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

<a id="canonical-2eb245aa564bd4c56ebb3c2b3948fc7e656b2d0a825dbf43c6c1f76d2c458bd2"></a>

<a id="canonical-8b11d47c6fee95067110bafd76476dec9f87803d635ad32506ce10ef58272579"></a>

## name property — irules / 74662a2151fe / 5

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

<a id="canonical-dec890cbfdfddc1025695a1197e9567644e9fb3579375ed015d575df97751a5b"></a>

<a id="canonical-3a6c38a9e6c7fc71fbc9b697906661c0a393dfe852f30508b687e17788af43c0"></a>

## namespace property — irules / 74662a2151fe / 6

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

<a id="canonical-e1ca62d1466dc7876d743046c1737b0a198f95898d298cc540605112cab8b5f1"></a>

<a id="canonical-3b1450a8ef96837b8e1b75b4d27563f5a377f92f6ca171c36845dc5b7fce1e3c"></a>

## tenant property — irules / 74662a2151fe / 7

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

<a id="canonical-1a0c8e98670b9569218c3d71f4b4229c9e7ed831070939e814d5af1ea29b38b1"></a>

<a id="canonical-a073f175ddb28559cc03e231306997a54a53c84402a994f46eb808efd6d746e5"></a>

## uid property — irules / 74662a2151fe / 8

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

<a id="canonical-177e90c37186dd9d5130602a0433e29d7766b77799b9a98d2fd1271d8f5c2176"></a>

## Next pages — irules / 74662a2151fe / 9

- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-90d932bd8f79925ec5d6b9945630df17fac61100dbca671694a7e9d6a95b4b80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1506fb2c8633d851775572602903b1209ac7cfa66db822fa821f88eb2bec2f2"></a>

## timeouts — timeouts / 128eafdb20e4 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- timeouts

<a id="canonical-cb0af81e6944e89a2340a75f3fb13ec59f670fb9a174ac02d1412f1fb99ac4b4"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-935992047b6efa60ff9a034eadeffa24507e79c209a03cc1805f7f1284fa6477"></a>

## Direct properties — timeouts / 128eafdb20e4 / 3

<a id="canonical-d0934290f09dcf77a29a38d373d1a237e8888ed3d4ccdce72d07235ed843cc49"></a>

<a id="canonical-9fdf07e0c20cd80a3461f018c8a83b552ad2802b4326e6e5a87620f3ea5947e2"></a>

## create property — timeouts / 128eafdb20e4 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2fb371a595a959d394d59ea724fb423e9a9e40dc73ad4578f73bc5ba39e9cb55"></a>

<a id="canonical-e3ada56fcd5c399acf0f22098df5506f0e44e8b55cb3e06e7f42daaa99120d65"></a>

## delete property — timeouts / 128eafdb20e4 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-792417dab7cd03d3ba512bae84641e342972c7e59afb6692a30bab5819d82094"></a>

<a id="canonical-d20e375901ddb5eca001b7258695f97eda67a90adfe657a5a80786d6ce5a1548"></a>

## read property — timeouts / 128eafdb20e4 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-5c92e8f0a032781e74f7f2030a149133f98f0c71f11046eb32c87af411bfb3ab"></a>

<a id="canonical-097b75ecf7033c45d26f249a5026bab4b87849c4dc0274b49808264bd6c5cf9c"></a>

## update property — timeouts / 128eafdb20e4 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-91db893f1a19b3b3435ce77ad007760403d7a307f5a2aa510502cafec76e84e9"></a>

## Next pages — timeouts / 128eafdb20e4 / 8

- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7371675c28fa4445edf39a041f2c83e87023cb48395b6055a1c95611833fb51b"></a>

## virtual_server — virtual_server / 010ec9ebeb8e / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- virtual_server

<a id="canonical-2b6f30603d702213b1c00e30bb081d82075e8fc771b28ae0759548e7ea8d2691"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration related to virtual server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http",
    "http3"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "tcp"),
  validators.ConflictingObjectAttributes("http",
    "udp"),
  validators.ConflictingObjectAttributes("http3",
    "https"),
  validators.ConflictingObjectAttributes("http3",
    "tcp"),
  validators.ConflictingObjectAttributes("http3",
    "udp"),
  validators.ConflictingObjectAttributes("https",
    "tcp"),
  validators.ConflictingObjectAttributes("https",
    "udp"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
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
  "x-ves-oneof-field-virtual_server_type": "[\"http\",\"http3\",\"https\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
virtual_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-c1a2eb17ee4d1b637aaa4a4187093736293f71d3fd7b1ca42d2571cae25e4d6f"></a>

## Direct properties — virtual_server / 010ec9ebeb8e / 3

- [access_profile](resources--application_profiles--reference--group-001.md#canonical-53a12df85327519c6a2adf959ce627f6de90e337ae3e8e8cde8013a4f139d1c2): complete subsection reference.

- [address_translation](resources--application_profiles--reference--group-001.md#canonical-c95851c1ed1a2083a5bf039f7729054276dd26634edb5d8da143c244124c75de): complete subsection reference.

- [auto_last_hop](resources--application_profiles--reference--group-001.md#canonical-19d0309cc3fa1af260f82783bfedeb541e2c98bc72ce8a8f4fd7c86fcb2ee49b): complete subsection reference.

- [clone_pool_client](resources--application_profiles--reference--group-002.md#canonical-49cbc9e2b3c6de12a5eda13f89c3ee8de6363bb835533295772a40cff05d459a): complete subsection reference.

- [clone_pool_server](resources--application_profiles--reference--group-002.md#canonical-14337ae926d7e352e80382de174238e4b0edd55e7898e86390befbf655a3561e): complete subsection reference.

<a id="canonical-f2123abf1bf98cac5f0620f786256465a818118895a1f9289b876f1c605ae8ac"></a>

<a id="canonical-d657b63e737166178e2928702d13d608304edfc0de12746aab8f8268ff2dc0d2"></a>

## connection_limit property — virtual_server / 010ec9ebeb8e / 4

Type: `"number"`. Optional.

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The.

Upstream description:

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The default is 0.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-e645b6dba2cd08dfd3b100fa336d24202c0dc6ba95cbd582d4e195e6a38ace73"></a>

<a id="canonical-8430b5c7ca5c053b00047782e3e7b647dd2ef0ee23ab9988c566bd9eea99cb2a"></a>

## connection_rate_limit property — virtual_server / 010ec9ebeb8e / 5

Type: `"number"`. Optional.

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where..

Upstream description:

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where
connection requests flood a virtual server. Setting this to 0 turns off connection limits. The
default is 0.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-64284b227bed9b4910c57bb9b2bf4f69629f9f33efdb84033e61b1cd75400e22): complete subsection reference.

- [default_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-5c818cde74f48e2aabfb3f4241d7397d5caca4fdfe3454393414b1733c6c78c9): complete subsection reference.

- [default_pool](resources--application_profiles--reference--group-002.md#canonical-289b8d2f0832c12551802b0bef3cb35ab80d0ed170b002c07a8b8958118a4f35): complete subsection reference.

- [fallback_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-d3e4a68978dff826f462df05c47e7623aa4f2ca9c502652b078ab1d26c387bd8): complete subsection reference.

- [fix_profile](resources--application_profiles--reference--group-002.md#canonical-303b196af6db71478aaec73bd47f0a1f59e11259b3e297a556ce916d881a3110): complete subsection reference.

- [http](resources--application_profiles--reference--group-002.md#canonical-9ea109f422920dc2372b6e6469314452e306debd1b830028804da4f13c33907a): complete subsection reference.

- [http3](resources--application_profiles--reference--group-002.md#canonical-039be1a3ec0416bbb07feba4ff81a7ea46e0ddfcf8a1c6b56c629e6e8776f55e): complete subsection reference.

- [https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d): complete subsection reference.

- [immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061): complete subsection reference.

- [last_hop_pool](resources--application_profiles--reference--group-003.md#canonical-5fc75b0574c39ea49487a9c012d59d535f61f433ca326badc03a12230bc11b50): complete subsection reference.

- [nat64](resources--application_profiles--reference--group-003.md#canonical-8237d6e97da7de4bd5b9a151386a463726e940692ca8d2762352f764c7fc6f42): complete subsection reference.

- [port_translation](resources--application_profiles--reference--group-003.md#canonical-d966d04dfcc007c24db2a54d14f1ad8f8d88ddfbd24e0c21ac1f126fd90a2b4c): complete subsection reference.

- [request_logging_profile](resources--application_profiles--reference--group-003.md#canonical-e48b0df1348ee94f794c3241260b6d7e56f733fdfb00cdf93002db17e8b5c4c8): complete subsection reference.

- [source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da): complete subsection reference.

- [statistics_profile](resources--application_profiles--reference--group-003.md#canonical-acf84624ce84de1878e7e0416c9389ebf7710cc4cca4e71825ff53f6c74fa497): complete subsection reference.

- [tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d): complete subsection reference.

- [udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7): complete subsection reference.

- [virtual_server_state](resources--application_profiles--reference--group-003.md#canonical-3b26eca3fc9a3a7609a95519ac21cca375b3782df408a853bdc25fcec9eabc05): complete subsection reference.

<a id="canonical-9e75536c70da03cb2dc957c89c7787e627a6533987623c9b7c52dee36235fe72"></a>

<a id="canonical-b4824612e3cf75b6f4270335af94d67429ca62cc947196d52e13789164c9f615"></a>

## vs_score property — virtual_server / 010ec9ebeb8e / 6

Type: `"number"`. Optional.

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The , meaning that no additional metric is applied
for the virtual server.

Upstream description:

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The default is 0, meaning that no additional
metric is applied for the virtual server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-cbb44fffd51d9dc555efaa79f3614d2bf933699e8ff15897011098bf0dae2eba"></a>

## Next pages — virtual_server / 010ec9ebeb8e / 7

- [virtual_server.access_profile](resources--application_profiles--reference--group-001.md#canonical-53a12df85327519c6a2adf959ce627f6de90e337ae3e8e8cde8013a4f139d1c2)
- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-c95851c1ed1a2083a5bf039f7729054276dd26634edb5d8da143c244124c75de)
- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-001.md#canonical-19d0309cc3fa1af260f82783bfedeb541e2c98bc72ce8a8f4fd7c86fcb2ee49b)
- [virtual_server.clone_pool_client](resources--application_profiles--reference--group-002.md#canonical-49cbc9e2b3c6de12a5eda13f89c3ee8de6363bb835533295772a40cff05d459a)
- [virtual_server.clone_pool_server](resources--application_profiles--reference--group-002.md#canonical-14337ae926d7e352e80382de174238e4b0edd55e7898e86390befbf655a3561e)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-64284b227bed9b4910c57bb9b2bf4f69629f9f33efdb84033e61b1cd75400e22)
- [virtual_server.default_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-5c818cde74f48e2aabfb3f4241d7397d5caca4fdfe3454393414b1733c6c78c9)
- [virtual_server.default_pool](resources--application_profiles--reference--group-002.md#canonical-289b8d2f0832c12551802b0bef3cb35ab80d0ed170b002c07a8b8958118a4f35)
- [virtual_server.fallback_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-d3e4a68978dff826f462df05c47e7623aa4f2ca9c502652b078ab1d26c387bd8)
- [virtual_server.fix_profile](resources--application_profiles--reference--group-002.md#canonical-303b196af6db71478aaec73bd47f0a1f59e11259b3e297a556ce916d881a3110)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-9ea109f422920dc2372b6e6469314452e306debd1b830028804da4f13c33907a)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-039be1a3ec0416bbb07feba4ff81a7ea46e0ddfcf8a1c6b56c629e6e8776f55e)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-1604b9a6fd4a35a1e191f43ff0fa50db7cabb21b48dfaa60407b47ab5ca9294d)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-98dc507f3e79f5a58e5e29e05fff8918180d965354bc090cb7e5c1b9311ac061)
- [virtual_server.last_hop_pool](resources--application_profiles--reference--group-003.md#canonical-5fc75b0574c39ea49487a9c012d59d535f61f433ca326badc03a12230bc11b50)
- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-8237d6e97da7de4bd5b9a151386a463726e940692ca8d2762352f764c7fc6f42)
- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-d966d04dfcc007c24db2a54d14f1ad8f8d88ddfbd24e0c21ac1f126fd90a2b4c)
- [virtual_server.request_logging_profile](resources--application_profiles--reference--group-003.md#canonical-e48b0df1348ee94f794c3241260b6d7e56f733fdfb00cdf93002db17e8b5c4c8)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-6eb3c4d745ed621496ebffc7f25a44e98228e6e80b82cbdc92312e13076d60da)
- [virtual_server.statistics_profile](resources--application_profiles--reference--group-003.md#canonical-acf84624ce84de1878e7e0416c9389ebf7710cc4cca4e71825ff53f6c74fa497)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1f199c7572eac3126faf4e9ea1aa153894aa6f33692fa82498e1437b23a34b2d)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-109ad4c08d75d204c352938626c782959fb21bf4af9139a2eb8780b57a457ab7)
- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-003.md#canonical-3b26eca3fc9a3a7609a95519ac21cca375b3782df408a853bdc25fcec9eabc05)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-53a12df85327519c6a2adf959ce627f6de90e337ae3e8e8cde8013a4f139d1c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f5acc8c1d24f9231af5fa22cc0ffed6e6fedc41bd92ea36d4799f483031c52b"></a>

## virtual_server.access_profile — virtual_server.access_profile / ec56fc07f612 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.access_profile

<a id="canonical-ab8de66f1cb320c80a79ce667b7c4b28df588d86a7d0e349a923824a7a228fcf"></a>

Type: `"object"`. list nested block, Optional.

Specifies an access policy that determines the authentication rules and access controls applied to
user sessions for this virtual server.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
access_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-368f1c8f6d00d4548b22082414f2224b85b5f0cb6aa95ac71a3e9f5a193ad1cc"></a>

## Direct properties — virtual_server.access_profile / ec56fc07f612 / 3

<a id="canonical-b6637b1a137e11f0c8c464f1980484805cce6b0491e0b328af0946aec4b6bf1c"></a>

<a id="canonical-50e84a15fedf8526a7bb48803a4802f66846b2ec151c21314e1536c05c4e7b84"></a>

## kind property — virtual_server.access_profile / ec56fc07f612 / 4

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

<a id="canonical-ff009ada6de3ebcbf9e89b335c1b7790cfb20f93edaa0fd949be60bcec4fdb2c"></a>

<a id="canonical-ca7626cb18256180777feeca12323cb8029aebf0ea18f51c1a3a5b75419fb041"></a>

## name property — virtual_server.access_profile / ec56fc07f612 / 5

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

<a id="canonical-12fd29faff78f39c3a0ff63692b615ba4f7dc4231e281cf0b443aa3746c9c881"></a>

<a id="canonical-f4066d62a97db14c6be611b1611597d01726b7f0062a8833a077b97b1e1f2aed"></a>

## namespace property — virtual_server.access_profile / ec56fc07f612 / 6

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

<a id="canonical-75b57a8a21de63a301b0be6f0fa9e976cfe6cf248c098ab34bcd60196c859a56"></a>

<a id="canonical-791a418a5ca3715b1a64dbe2b774945f73b24dc626f4433ecd08f677071fc9c5"></a>

## tenant property — virtual_server.access_profile / ec56fc07f612 / 7

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

<a id="canonical-09d53c318b2dd95fdb443055e732ecc1c1079aaca916c4fec7c1788686c7b70c"></a>

<a id="canonical-939d6c7c7783c5c86c9d18737e570c98ec4d09b9acc0b3a4e50c92571c50fc6e"></a>

## uid property — virtual_server.access_profile / ec56fc07f612 / 8

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

<a id="canonical-0b54d39d3966e42da7a015f1421524abb9cad2765912de06585327cfd4551d54"></a>

## Next pages — virtual_server.access_profile / ec56fc07f612 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-c95851c1ed1a2083a5bf039f7729054276dd26634edb5d8da143c244124c75de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09173ff2f0640d60ed207742350b82a4263507fa9f75e44aa8b54c61211fae8c"></a>

## virtual_server.address_translation — virtual_server.address_translation / 153cddaa2f75 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.address_translation

<a id="canonical-303b6501ee26d1cf62be8306849f86e75608671428768441e8e3dd8105f33ffa"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address.

Upstream description:

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address. The default is
enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("address_translation_disable",
    "address_translation_enable")}
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
  "x-ves-oneof-field-address_translation_choice": "[\"address_translation_disable\",\"address_translation_enable\"]"
}
```

Terraform syntax:

```terraform
address_translation {
  # Configure direct properties listed below.
}
```

<a id="canonical-cf9083e961c996c4e65c405ae3ed6d9b85e440e958ec5b7b2355c7e8e15fbae1"></a>

## Direct properties — virtual_server.address_translation / 153cddaa2f75 / 3

- [address_translation_disable](resources--application_profiles--reference--group-001.md#canonical-806a4695197db2b3d74d2be4b8c6766a272cf92ce191f3af748bb9f5b9f613b2): complete subsection reference.

- [address_translation_enable](resources--application_profiles--reference--group-001.md#canonical-c6845d6206e2848b432c3ae45e3ec37535e5d6ef64540d863a53172c635cd98e): complete subsection reference.

<a id="canonical-71d15481ba237ffa0ecbf507a8d788ff11e911a421b27d6f85aef563707ab928"></a>

## Next pages — virtual_server.address_translation / 153cddaa2f75 / 4

- [virtual_server.address_translation.address_translation_disable](resources--application_profiles--reference--group-001.md#canonical-806a4695197db2b3d74d2be4b8c6766a272cf92ce191f3af748bb9f5b9f613b2)
- [virtual_server.address_translation.address_translation_enable](resources--application_profiles--reference--group-001.md#canonical-c6845d6206e2848b432c3ae45e3ec37535e5d6ef64540d863a53172c635cd98e)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-806a4695197db2b3d74d2be4b8c6766a272cf92ce191f3af748bb9f5b9f613b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-024c6e60b54ff85102c80487b2e62cc9dbf6d98ee07342c8cb8a48f1efe9d92e"></a>

## virtual_server.address_translation.address_translation_disable — virtual_server.address_translation.address_translation_disable / 9969d1f55f03 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-c95851c1ed1a2083a5bf039f7729054276dd26634edb5d8da143c244124c75de)
- virtual_server.address_translation.address_translation_disable

<a id="canonical-066399375ec4ab4dde58ddfc517ea3d6a6ac51bbcbaf2dcc4c86725685c8e82c"></a>

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
address_translation_disable = {}
```

<a id="canonical-e667ec52351900ae51bd3813a1d49e53ea78178dbe67a78638c245aeddc95e24"></a>

## Direct properties — virtual_server.address_translation.address_translation_disable / 9969d1f55f03 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5d30d09f6980670724d30ad8bffa81f8c4f0d6e5a594ec32c2da676e0024fda"></a>

## Next pages — virtual_server.address_translation.address_translation_disable / 9969d1f55f03 / 4

- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-c95851c1ed1a2083a5bf039f7729054276dd26634edb5d8da143c244124c75de)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-c6845d6206e2848b432c3ae45e3ec37535e5d6ef64540d863a53172c635cd98e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa2c82b19e5f3f0877c569d781fa3ebf9c7f1b7620c1ec9db53a5493d1705f2e"></a>

## virtual_server.address_translation.address_translation_enable — virtual_server.address_translation.address_translation_enable / 470303e63190 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-c95851c1ed1a2083a5bf039f7729054276dd26634edb5d8da143c244124c75de)
- virtual_server.address_translation.address_translation_enable

<a id="canonical-3dbdb8919f290a174f662ab33a6c59db45b7cf8af8d15d12cf486da14bce7f6a"></a>

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
address_translation_enable = {}
```

<a id="canonical-a0a7b5955960e98466411aa2bdea57ec9cecd6f7237d915b5cced65448772c3e"></a>

## Direct properties — virtual_server.address_translation.address_translation_enable / 470303e63190 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4968bd56b9ecd1ee9d4e27870099babe78fdf5524e0d0f1d02c75b75cdc8bed"></a>

## Next pages — virtual_server.address_translation.address_translation_enable / 470303e63190 / 4

- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-c95851c1ed1a2083a5bf039f7729054276dd26634edb5d8da143c244124c75de)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-19d0309cc3fa1af260f82783bfedeb541e2c98bc72ce8a8f4fd7c86fcb2ee49b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-358f5b7479949979c13b021017c528e0e7278f31372d32b8deec187778840088"></a>

## virtual_server.auto_last_hop — virtual_server.auto_last_hop / cd99f96182fb / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-4002983ebe21e2fc570e5fe9278d8848d8c452bfdc37b1991b0315d9993712d6)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-8ac0e20ace645fbe7b29e445dd7f617c04c454d0c860cf74f250194f06f11276)
- virtual_server.auto_last_hop

<a id="canonical-920ff090d872f321ac2c53da5aa0d40d0219faec680099acb7cb7566e692f88a"></a>

Type: `"object"`. single nested block, Optional.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_last_hop_default",
    "auto_last_hop_disable"),
  validators.ConflictingObjectAttributes("auto_last_hop_default",
    "auto_last_hop_enable"),
  validators.ConflictingObjectAttributes("auto_last_hop_disable",
    "auto_last_hop_enable")}
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
  "x-ves-oneof-field-auto_last_hop_choice": "[\"auto_last_hop_default\",\"auto_last_hop_disable\",\"auto_last_hop_enable\"]"
}
```

Terraform syntax:

```terraform
auto_last_hop {
  # Configure direct properties listed below.
}
```
