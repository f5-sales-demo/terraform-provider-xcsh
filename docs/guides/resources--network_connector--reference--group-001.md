---
page_title: "xcsh_network_connector reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector reference."
---

# xcsh_network_connector reference

<a id="canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-079b0dd3279f76ea685e5e50f0e06549a6999752647141051605741c15295dc3"></a>

## Property reference — Property reference / 14987134c10c / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- Property reference

<a id="canonical-dc1be983b6c2361cca597aed183e6ec0f4ff1ea23146aec86c32f62366cb0667"></a>

## Direct properties — Property reference / 14987134c10c / 3

<a id="canonical-148eda2b095a14275ab07f9677ef2f3b899f1a3bbd022f6dc1bfac2e6464dc9c"></a>

<a id="canonical-61c352ed59ebd05031b709df6b6837cb3347d10a7fb14c5b33f30ec56b1ceb39"></a>

## annotations property — Property reference / 14987134c10c / 4

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

<a id="canonical-946106397d8bf0ad8b5d74dc3ceb1fe2da0c04a4c62b5fcffb99bc2b71390cd3"></a>

<a id="canonical-86d212fa411af5ac0ebc8781edb81caff253bda73f3bc36b5975bd82a2a06b03"></a>

## description property — Property reference / 14987134c10c / 5

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

<a id="canonical-0fba9c0081b51b40ceac5f30883569a0001b2e55a9b08a2594437b15cc748761"></a>

<a id="canonical-5ab2edeec4fb56384e373bbd7dec94f905683f97ad869ac1532e8da63cc742c5"></a>

## disable property — Property reference / 14987134c10c / 6

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

- [disable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-95ffc30042358376339ad7cd10495b9bba44536c1858bf15ec1ea0949eeace59): complete subsection reference.

- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48): complete subsection reference.

<a id="canonical-efd0992c56ce46a73946bdfa55a27505e0c3cf533f3ccd5fcd9ed01916d33302"></a>

<a id="canonical-750ebf8f932009bfc7cde72a845b170bc809c4fc911be0b49aff09ef9fc88fed"></a>

## id property — Property reference / 14987134c10c / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-380c7d878494bb34ef8b168f92a7254fdcb34dc1e9098a792d5d6a4d5853eb31"></a>

<a id="canonical-8821d8004d2aafdeab65604b67c785781e9a2bed10621ad0e25bdf7f226067d4"></a>

## labels property — Property reference / 14987134c10c / 8

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

<a id="canonical-46536f9406816cc8daeec451ce09d716c9e60b9fe7507e14f3b60a6b939c4c4f"></a>

<a id="canonical-45601559051f830b69083f2fb3d6f7c85c15624cecda641d17a03f6de8935f8b"></a>

## name property — Property reference / 14987134c10c / 9

Type: `"string"`. Required.

Name of the Network Connector. Must be unique within the namespace.

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

<a id="canonical-adcdcb2848cdb45bfdce436c03b4c689aebd57d32374d3030614eb37698eefe8"></a>

<a id="canonical-9392e3971f95410def699d0ea17624e1a538f68ef80d5ff4315c34a58cf8a96e"></a>

## namespace property — Property reference / 14987134c10c / 10

Type: `"string"`. Required.

Namespace where the Network Connector is created.

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

- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-7ba0153a60944bef2faf88462dc67925427df3cb0c4cf89522c40b3b34f2d1b8): complete subsection reference.

- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-35b27d234276bca17116a2d1a66b0e4993dea66284127d3bd54048641abe1489): complete subsection reference.

- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-fcd5ed60be849b0d67ff6027dbdf14f6d887bb7ff3b43e5abb8a217e7a6e1315): complete subsection reference.

- [timeouts](resources--network_connector--reference--group-001.md#canonical-e9cc6a09d7f9a114af512eb9694817eae4f4c08a5d7c900822027349c63c8049): complete subsection reference.

<a id="canonical-3d72c3bdf5d096e954e077b40a7a05dd5366b48411b57f209ff87de331b6e581"></a>

## All schema paths — Property reference / 14987134c10c / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_connector--reference--group-001.md#canonical-148eda2b095a14275ab07f9677ef2f3b899f1a3bbd022f6dc1bfac2e6464dc9c) |
| `description` | [description](resources--network_connector--reference--group-001.md#canonical-946106397d8bf0ad8b5d74dc3ceb1fe2da0c04a4c62b5fcffb99bc2b71390cd3) |
| `disable` | [disable](resources--network_connector--reference--group-001.md#canonical-0fba9c0081b51b40ceac5f30883569a0001b2e55a9b08a2594437b15cc748761) |
| `disable_forward_proxy` | [disable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-80b24b57c07800ec8f31dc7ba59b322a3cc454c13bf5e00a6da0cc2cac698591) |
| `enable_forward_proxy` | [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-a433f6d5a0c93f69dd0b4f24de7eae59b92714eb08f8f25e784229687c5578de) |
| `enable_forward_proxy.connection_timeout` | [enable_forward_proxy.connection_timeout](resources--network_connector--reference--group-001.md#canonical-a073ff3cac8c243c6ec3127dd4ee22365a24cf6b40aad9f9634dc4eb30509603) |
| `enable_forward_proxy.max_connect_attempts` | [enable_forward_proxy.max_connect_attempts](resources--network_connector--reference--group-001.md#canonical-3068449d4eaa9564d06cdc57c3050defa0a8ad4d1bdfa55edcd7c2554444998f) |
| `enable_forward_proxy.no_interception` | [enable_forward_proxy.no_interception](resources--network_connector--reference--group-001.md#canonical-8055bdaa3de5886959996c069a53d604fe1b7dd3032ee442a4c2b7e624ab53c2) |
| `enable_forward_proxy.tls_intercept` | [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-33f4dff0acfc492921ffa7ca389ed54886c04e25143b817fddd0687cf2ea12ea) |
| `enable_forward_proxy.tls_intercept.custom_certificate` | [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-3bfb37711785bf64a54b5300e01ed3c6a243ab1716e3ff8f843b8e83235f4038) |
| `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` | [enable_forward_proxy.tls_intercept.custom_certificate.certificate_url](resources--network_connector--reference--group-001.md#canonical-0b7ab67912a236eaa0b089757e3fa059ff2c3f39fb01e693dd5471f09ca29c12) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](resources--network_connector--reference--group-001.md#canonical-269fff1e9f6859551f3b62b7980b1d563f38469632c3c7fcd68e9eebed06cbea) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](resources--network_connector--reference--group-001.md#canonical-e649d3b41470295eccbb33cb6a40cbf4ddf27a619a47ed567c20af177c9df407) |
| `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` | [enable_forward_proxy.tls_intercept.custom_certificate.description_spec](resources--network_connector--reference--group-001.md#canonical-09389b2a14fdce4ca2943ee109394bcb1c31b4c738810cfeaf7a2dabd90b6fa6) |
| `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` | [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](resources--network_connector--reference--group-001.md#canonical-20ce506e7d7ef6b46363af38ce8a8e5418b9c5ce9666350d5f37fb464b018725) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-d5c9829501d4a17e6dcbcf70a2fc38f443283731bfeba42bfca5434fdfc82ced) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--network_connector--reference--group-001.md#canonical-3ce2e5076236aec1e10329473f7a08007fdafe554fc055b5810f717fbdd64ba1) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](resources--network_connector--reference--group-001.md#canonical-16a6a0abf03c2d62e4f4b8e079c177ce8809ca7dd82612a483a80b11423e5635) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](resources--network_connector--reference--group-001.md#canonical-db72773d4726e007f9ab00cce295e432a2280911a07b67817999be770cb46df5) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](resources--network_connector--reference--group-001.md#canonical-c33a7a3b597de5b071f238519df1d24d406b5dece4263c1ea746384ce2471326) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](resources--network_connector--reference--group-001.md#canonical-e9068d07b9f5ee276dfe7f4ba8420c86bc84101a7c4b11f86f49cba1dbf940a5) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](resources--network_connector--reference--group-001.md#canonical-b7c8a8f697dee3e326ec8bbb0af65e6b04e2c7a551f12fed7a82e77fe4d91dbd) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url](resources--network_connector--reference--group-001.md#canonical-159b5fff11d08a4d33ddb1a04777a117431585a8d68e09bf63e5f710daed90a6) |
| `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` | [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](resources--network_connector--reference--group-001.md#canonical-881f7d177765805af46c3ed69504e99f494e1a5afa09a47eabd4e8440b677671) |
| `enable_forward_proxy.tls_intercept.enable_for_all_domains` | [enable_forward_proxy.tls_intercept.enable_for_all_domains](resources--network_connector--reference--group-001.md#canonical-972bd990316c16256f7c14530251b93c61bf1767e1336423ada15446d7b538f3) |
| `enable_forward_proxy.tls_intercept.policy` | [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-002cd13b4f1def10788bcaf9aa33a585c2ddff9337bffaa4922cdae7a37253c6) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules` | [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-746bf6d9211f9ce79f9753a197afbaa814fd086233dac9da66425e568d69fc1a) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](resources--network_connector--reference--group-001.md#canonical-df9c73a728504e9993a66b6f63769b9349ae77b64b2af7245bdbfdfc8a200d8f) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](resources--network_connector--reference--group-001.md#canonical-741df57b10998b9b855a54a95746f5d38c912b8b81c74f940fc027234e2e64db) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value](resources--network_connector--reference--group-001.md#canonical-c88b91203e7df2b5c086432046e82a27943a6edf5975e1f92656d8f1e2724aeb) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value](resources--network_connector--reference--group-001.md#canonical-e7979cac5549dac2b1d05256b88f8e39154111de4065ff099ba754b3e3de61aa) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value](resources--network_connector--reference--group-001.md#canonical-91ac5444d552aa6d00795a94746d9a404173fe0bff2bf1d6df50276f0cff13ca) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](resources--network_connector--reference--group-001.md#canonical-01c0915f92e1fa24b34594795c26d731dbd193616c619e351d3901ebb6eb597c) |
| `enable_forward_proxy.tls_intercept.trusted_ca_url` | [enable_forward_proxy.tls_intercept.trusted_ca_url](resources--network_connector--reference--group-001.md#canonical-1d1fcb80776a89e161d0c5922b286bf5a80e46d8d4fc9b8f29d29054f9c934d1) |
| `enable_forward_proxy.tls_intercept.volterra_certificate` | [enable_forward_proxy.tls_intercept.volterra_certificate](resources--network_connector--reference--group-001.md#canonical-6fc0ec2d69d2cd7aa7be54756166e77b7362cc62db2facf535737978b805cff5) |
| `enable_forward_proxy.tls_intercept.volterra_trusted_ca` | [enable_forward_proxy.tls_intercept.volterra_trusted_ca](resources--network_connector--reference--group-001.md#canonical-a6f52b683ccdae05b8eeb69656d8a5a0ef1f8b9f46de3ab2aa285fac6fc40602) |
| `enable_forward_proxy.white_listed_ports` | [enable_forward_proxy.white_listed_ports](resources--network_connector--reference--group-001.md#canonical-077c590ed29c98d71f3d26ac8b2fc8129d2058abbee61d85b22035e936cd2517) |
| `enable_forward_proxy.white_listed_prefixes` | [enable_forward_proxy.white_listed_prefixes](resources--network_connector--reference--group-001.md#canonical-6eb4eacff1508a8f5bfe3c53b207b72fc5e459a1b28e0b6295a8533eadab8b07) |
| `id` | [id](resources--network_connector--reference--group-001.md#canonical-efd0992c56ce46a73946bdfa55a27505e0c3cf533f3ccd5fcd9ed01916d33302) |
| `labels` | [labels](resources--network_connector--reference--group-001.md#canonical-380c7d878494bb34ef8b168f92a7254fdcb34dc1e9098a792d5d6a4d5853eb31) |
| `name` | [name](resources--network_connector--reference--group-001.md#canonical-46536f9406816cc8daeec451ce09d716c9e60b9fe7507e14f3b60a6b939c4c4f) |
| `namespace` | [namespace](resources--network_connector--reference--group-001.md#canonical-adcdcb2848cdb45bfdce436c03b4c689aebd57d32374d3030614eb37698eefe8) |
| `sli_to_global_dr` | [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-dddd546acf32c1b16ad0fdfc09b6a93192e79b6cd991255f57e87b5bd083d5a2) |
| `sli_to_global_dr.global_vn` | [sli_to_global_dr.global_vn](resources--network_connector--reference--group-001.md#canonical-a1bd917e7ef7d2eba215b6b1965b8c412bbf5662fab9cdbd6b25785487b45a2e) |
| `sli_to_global_dr.global_vn.name` | [sli_to_global_dr.global_vn.name](resources--network_connector--reference--group-001.md#canonical-fd749c9c27657fd50752c74eed8343bf628e56459288170f8ff622cc9953fa02) |
| `sli_to_global_dr.global_vn.namespace` | [sli_to_global_dr.global_vn.namespace](resources--network_connector--reference--group-001.md#canonical-4f5fed7331f07d8d31af3eef83c57bc6cac99a1994820f454ea87b60b7101d00) |
| `sli_to_global_dr.global_vn.tenant` | [sli_to_global_dr.global_vn.tenant](resources--network_connector--reference--group-001.md#canonical-12999551141cc474a7f0eb804e63c43e618b38983668aa89b194aefb47b1d7bc) |
| `sli_to_slo_snat` | [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-b1c8a6864018866d3cbb36eb67b3ab024ec044d69d025bceb19409e4bc08796e) |
| `sli_to_slo_snat.default_gw_snat` | [sli_to_slo_snat.default_gw_snat](resources--network_connector--reference--group-001.md#canonical-537cec77c5fff3b8222c411f991c18505b4de52f0279b8a32d17747e3cb9af5a) |
| `sli_to_slo_snat.interface_ip` | [sli_to_slo_snat.interface_ip](resources--network_connector--reference--group-001.md#canonical-1ea2269efc3590842da750c01a3c9a5db5222cbf058888cca862ca365e54190e) |
| `slo_to_global_dr` | [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-3a266d5cae820174cd7e4180ae374bdf6eb94750daec738d1aa5ef874e156a63) |
| `slo_to_global_dr.global_vn` | [slo_to_global_dr.global_vn](resources--network_connector--reference--group-001.md#canonical-8d7075e4450446d60b2552feada053a61b0c2e04f46f60ac5ef3de7477bfe24c) |
| `slo_to_global_dr.global_vn.name` | [slo_to_global_dr.global_vn.name](resources--network_connector--reference--group-001.md#canonical-9af0cb8d0cafc3c87c13b78b2a1f564d8b8af4a59f28569a733b035c56a986c1) |
| `slo_to_global_dr.global_vn.namespace` | [slo_to_global_dr.global_vn.namespace](resources--network_connector--reference--group-001.md#canonical-8f93b2fa1e6fed7c1bbfdb5dcff5dd1f018a09a749c11a8ca7d71352110f4707) |
| `slo_to_global_dr.global_vn.tenant` | [slo_to_global_dr.global_vn.tenant](resources--network_connector--reference--group-001.md#canonical-b00b91762de98a27c4343de26e3f444e17ca4cfb61b1da316f00ba5a7268d955) |
| `timeouts` | [timeouts](resources--network_connector--reference--group-001.md#canonical-66170410bdbeffd11beee7aa52dd12c5fbb1900676ac141d427958c92f2e9f2a) |
| `timeouts.create` | [timeouts.create](resources--network_connector--reference--group-001.md#canonical-ca2f195192197a0de1b2ec49f589900e572dd921f6fb19c6cab63bc1fb281a1f) |
| `timeouts.delete` | [timeouts.delete](resources--network_connector--reference--group-001.md#canonical-733429a8b77dcd2c4ad70700ed49b0701378652f520def15616dde6d37231d11) |
| `timeouts.read` | [timeouts.read](resources--network_connector--reference--group-001.md#canonical-588b5ff72b960f12c033ec06436d61a1ce99a6013b806f83cb82936a59550695) |
| `timeouts.update` | [timeouts.update](resources--network_connector--reference--group-001.md#canonical-d8ba203706ea3c38f701210d91457950937d095999da3e98b2cccac015b4efee) |

<a id="canonical-21263871aadc599a69e245e430f323484a56d85821cae3f1646012cd5b6809db"></a>

## Next pages — Property reference / 14987134c10c / 12

- [disable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-95ffc30042358376339ad7cd10495b9bba44536c1858bf15ec1ea0949eeace59)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-7ba0153a60944bef2faf88462dc67925427df3cb0c4cf89522c40b3b34f2d1b8)
- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-35b27d234276bca17116a2d1a66b0e4993dea66284127d3bd54048641abe1489)
- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-fcd5ed60be849b0d67ff6027dbdf14f6d887bb7ff3b43e5abb8a217e7a6e1315)
- [timeouts](resources--network_connector--reference--group-001.md#canonical-e9cc6a09d7f9a114af512eb9694817eae4f4c08a5d7c900822027349c63c8049)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-95ffc30042358376339ad7cd10495b9bba44536c1858bf15ec1ea0949eeace59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f210c82ddf977dd7f34ae282fb8f107674af7696aa72eeef60941794dc005339"></a>

## disable_forward_proxy — disable_forward_proxy / 9b03c1040167 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- disable_forward_proxy

<a id="canonical-80b24b57c07800ec8f31dc7ba59b322a3cc454c13bf5e00a6da0cc2cac698591"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_forward\_proxy, enable\_forward\_proxy; Default: disable\_forward\_proxy\]
Configuration parameter for disable forward proxy.

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

- [disable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-80b24b57c07800ec8f31dc7ba59b322a3cc454c13bf5e00a6da0cc2cac698591)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-a433f6d5a0c93f69dd0b4f24de7eae59b92714eb08f8f25e784229687c5578de)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_forward_proxy = {}
```

<a id="canonical-7f74705b51384500626ee9157fb30bc547c74319f82a1995ab9645290646a28f"></a>

## Direct properties — disable_forward_proxy / 9b03c1040167 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be0e769717f5627979b8deaf4a780fccd00890ce9656ce1ad14efcb7b624db9a"></a>

## Next pages — disable_forward_proxy / 9b03c1040167 / 4

- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2572dae11ee07fd8b0b3dc331fad934b38ca8e5e5613cce63680b0247ea9814"></a>

## enable_forward_proxy — enable_forward_proxy / b5b0e0b6ca39 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- enable_forward_proxy

<a id="canonical-a433f6d5a0c93f69dd0b4f24de7eae59b92714eb08f8f25e784229687c5578de"></a>

Type: `"object"`. single nested block, Optional.

Fine tune forward proxy behavior Few configurations allowed are White listed ports and IP prefixes:
Forward proxy does application protocol detection and server name(SNI) detection by peeking into the
traffic on the incoming downstream connection. Few protocols doesn't have client sending the..

Upstream description:

Fine tune forward proxy behavior

Few configurations allowed are

White listed ports and IP prefixes: Forward proxy does application protocol detection and server
name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols
doesn't have client sending the first data. In such cases, protocol and SNI detection fails. This
configuration allows, skipping protocol and SNI detection for whitelisted IP-prefix-list and ports
connection\_timeout: The timeout for new network connections to upstream server.
Max\_connect\_attempts: Maximum number of attempts made to make new network connection to upstream
server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_interception",
    "tls_intercept")}
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
  "x-ves-oneof-field-tls_interception_choice": "[\"no_interception\",\"tls_intercept\"]"
}
```

Terraform syntax:

```terraform
enable_forward_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2deaaa1b1b729fced5b0c131c8e0673d6517e43076e636132d57cd8ff05ebbd"></a>

## Direct properties — enable_forward_proxy / b5b0e0b6ca39 / 3

<a id="canonical-a073ff3cac8c243c6ec3127dd4ee22365a24cf6b40aad9f9634dc4eb30509603"></a>

<a id="canonical-fc50c6653528a609ea650740a15949f1b98b23454c049b332b84e2ca5d2ce334"></a>

## connection_timeout property — enable_forward_proxy / b5b0e0b6ca39 / 4

Type: `"number"`. Optional.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3068449d4eaa9564d06cdc57c3050defa0a8ad4d1bdfa55edcd7c2554444998f"></a>

<a id="canonical-4a3873e04aa423e1f59ae8ba5b492ea48b6b4a5aaf6e51d61a67e2eff57d4409"></a>

## max_connect_attempts property — enable_forward_proxy / b5b0e0b6ca39 / 5

Type: `"number"`. Optional.

Specifies the allowed number of retries on connect failure to upstream server. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

- [no_interception](resources--network_connector--reference--group-001.md#canonical-68158da90f64d9b55c9b6411f226e64d6d24093eb50d70acc2de7a52947fa3bf): complete subsection reference.

- [tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697): complete subsection reference.

<a id="canonical-077c590ed29c98d71f3d26ac8b2fc8129d2058abbee61d85b22035e936cd2517"></a>

<a id="canonical-f7e1a8efacecc7beda8c0210d66afb4617c93c3f217788afde40020dee8177b5"></a>

## white_listed_ports property — enable_forward_proxy / b5b0e0b6ca39 / 6

Type: `["list", "number"]`. Optional.

Traffic to these destination TCP ports is not subjected to protocol parsing Example 'tmate' server
port.

Upstream description:

Traffic to these destination TCP ports is not subjected to protocol parsing Example "tmate" server
port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="canonical-6eb4eacff1508a8f5bfe3c53b207b72fc5e459a1b28e0b6295a8533eadab8b07"></a>

<a id="canonical-2df1355c1199819a36ad42d12ccedcd9845a14206042de0c55594445d9f748e1"></a>

## white_listed_prefixes property — enable_forward_proxy / b5b0e0b6ca39 / 7

Type: `["list", "string"]`. Optional.

Traffic to these destination IP prefixes is not subjected to protocol parsing Example 'tmate' server
IP.

Upstream description:

Traffic to these destination IP prefixes is not subjected to protocol parsing Example "tmate" server
IP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c42f7aa1e482aed2dda592c9c3d8cfed9e738a48cfb7796e6935425c47fe0653"></a>

## Next pages — enable_forward_proxy / b5b0e0b6ca39 / 8

- [enable_forward_proxy.no_interception](resources--network_connector--reference--group-001.md#canonical-68158da90f64d9b55c9b6411f226e64d6d24093eb50d70acc2de7a52947fa3bf)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-68158da90f64d9b55c9b6411f226e64d6d24093eb50d70acc2de7a52947fa3bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f6cf77b70156dca25172e038dde23cce9a7ba4f94ad82e9de93e60a1c63d2f8"></a>

## enable_forward_proxy.no_interception — enable_forward_proxy.no_interception / 303f9a6e05c6 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- enable_forward_proxy.no_interception

<a id="canonical-8055bdaa3de5886959996c069a53d604fe1b7dd3032ee442a4c2b7e624ab53c2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no interception.

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
no_interception = {}
```

<a id="canonical-ce7ea21a6d194e968d126ca485f9a7812e2f737e9e9901226ace05853ee90d4b"></a>

## Direct properties — enable_forward_proxy.no_interception / 303f9a6e05c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0c67835e0d988fdd294925f1aff434066e9c7a4d4c5a1c011cd5c9f58938ef3"></a>

## Next pages — enable_forward_proxy.no_interception / 303f9a6e05c6 / 4

- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ba98ddd5f0e99cf7448495216e4b447ed117e558209d8e03b990e8eeab7174e"></a>

## enable_forward_proxy.tls_intercept — enable_forward_proxy.tls_intercept / 2d47baa4a5b6 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- enable_forward_proxy.tls_intercept

<a id="canonical-33f4dff0acfc492921ffa7ca389ed54886c04e25143b817fddd0687cf2ea12ea"></a>

Type: `"object"`. single nested block, Optional.

Configuration to enable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_certificate",
    "volterra_certificate"),
  validators.ConflictingObjectAttributes("enable_for_all_domains",
    "policy"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
tls_intercept {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb7867c1bcecb9b6888bff30741eb8c52953d068455064bff476b4067cc4360a"></a>

## Direct properties — enable_forward_proxy.tls_intercept / 2d47baa4a5b6 / 3

- [custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef): complete subsection reference.

- [enable_for_all_domains](resources--network_connector--reference--group-001.md#canonical-34a3a496cb716ab5261674a35b0e6426d62ea6c990724e20dde6e4a5f01645a3): complete subsection reference.

- [policy](resources--network_connector--reference--group-001.md#canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697): complete subsection reference.

<a id="canonical-1d1fcb80776a89e161d0c5922b286bf5a80e46d8d4fc9b8f29d29054f9c934d1"></a>

<a id="canonical-adee9fae2a9cff25ddf663c4ec2f3f99350af06ead26937bf065036437b2859f"></a>

## trusted_ca_url property — enable_forward_proxy.tls_intercept / 2d47baa4a5b6 / 4

Type: `"string"`. Optional.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](resources--network_connector--reference--group-001.md#canonical-1c32883c7fa89f245fc065a03ba8a8eba7c80b2d9246aab37530cea789583fa3): complete subsection reference.

- [volterra_trusted_ca](resources--network_connector--reference--group-001.md#canonical-b96a37769a8b8de74a47ece1d0faad8e85504cfff33cc302d7b361e97420223c): complete subsection reference.

<a id="canonical-6e5261467cb2faf90104388ec07f1bbb542ad7179e18ffede9262ed6b4c2241b"></a>

## Next pages — enable_forward_proxy.tls_intercept / 2d47baa4a5b6 / 5

- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- [enable_forward_proxy.tls_intercept.enable_for_all_domains](resources--network_connector--reference--group-001.md#canonical-34a3a496cb716ab5261674a35b0e6426d62ea6c990724e20dde6e4a5f01645a3)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697)
- [enable_forward_proxy.tls_intercept.volterra_certificate](resources--network_connector--reference--group-001.md#canonical-1c32883c7fa89f245fc065a03ba8a8eba7c80b2d9246aab37530cea789583fa3)
- [enable_forward_proxy.tls_intercept.volterra_trusted_ca](resources--network_connector--reference--group-001.md#canonical-b96a37769a8b8de74a47ece1d0faad8e85504cfff33cc302d7b361e97420223c)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3dfc422be5fe1e8b0d0abbc9ba988e8e8c24eae469f72d574ee2d6df2235284"></a>

## enable_forward_proxy.tls_intercept.custom_certificate — enable_forward_proxy.tls_intercept.custom_certificate / 9d3d45123d95 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- enable_forward_proxy.tls_intercept.custom_certificate

<a id="canonical-3bfb37711785bf64a54b5300e01ed3c6a243ab1716e3ff8f843b8e83235f4038"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

Terraform syntax:

```terraform
custom_certificate {
  # Configure direct properties listed below.
}
```

<a id="canonical-624b67b28b91d2ef96e8f80a25fc99467fee6eba6d7e05641467633240e5afae"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate / 9d3d45123d95 / 3

<a id="canonical-0b7ab67912a236eaa0b089757e3fa059ff2c3f39fb01e693dd5471f09ca29c12"></a>

<a id="canonical-3c756f421744e002e2496049ad5dd679f5621de2c5f9131342e46f56a130d94c"></a>

## certificate_url property — enable_forward_proxy.tls_intercept.custom_certificate / 9d3d45123d95 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--network_connector--reference--group-001.md#canonical-6aa296799653c48d8598e372f09a66145be85e77b014c2fe1e71fa4230b90823): complete subsection reference.

<a id="canonical-09389b2a14fdce4ca2943ee109394bcb1c31b4c738810cfeaf7a2dabd90b6fa6"></a>

<a id="canonical-7cfe496aaf7007826c64557a9052e56a10c0f42ffa3d2ca3a2a4c7560b4d3127"></a>

## description_spec property — enable_forward_proxy.tls_intercept.custom_certificate / 9d3d45123d95 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--network_connector--reference--group-001.md#canonical-c98f9d2adbbabdd0b614d25c17d359854de979e945e276b006664993640dcda5): complete subsection reference.

- [private_key](resources--network_connector--reference--group-001.md#canonical-a9952a7a1919cece98f4ee174a3419b4ecb530b5e2c835333d46360f23315069): complete subsection reference.

- [use_system_defaults](resources--network_connector--reference--group-001.md#canonical-b2ec67c0cbebc681952b7b6b07a58b913a12db46ed1679814f081f7b6182a592): complete subsection reference.

<a id="canonical-9769ec205f0c01cb0661e56145e5c1ff03f90b3d2c5669de73800aae4f670ed2"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate / 9d3d45123d95 / 6

- [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](resources--network_connector--reference--group-001.md#canonical-6aa296799653c48d8598e372f09a66145be85e77b014c2fe1e71fa4230b90823)
- [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](resources--network_connector--reference--group-001.md#canonical-c98f9d2adbbabdd0b614d25c17d359854de979e945e276b006664993640dcda5)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-a9952a7a1919cece98f4ee174a3419b4ecb530b5e2c835333d46360f23315069)
- [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](resources--network_connector--reference--group-001.md#canonical-b2ec67c0cbebc681952b7b6b07a58b913a12db46ed1679814f081f7b6182a592)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-6aa296799653c48d8598e372f09a66145be85e77b014c2fe1e71fa4230b90823"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-083e12bafce56b530648d5e4d3fb9ee857cf10277f822f5ed2a17c03c364f45b"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 0ffba7cc7b8a / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-269fff1e9f6859551f3b62b7980b1d563f38469632c3c7fcd68e9eebed06cbea"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-4bf93ccaea4026b97009fb1770b5a81a3cd80370cbde6d6c7f642d21df205ac3"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 0ffba7cc7b8a / 3

<a id="canonical-e649d3b41470295eccbb33cb6a40cbf4ddf27a619a47ed567c20af177c9df407"></a>

<a id="canonical-7a1386b00fd551717de666cd55ac5208333dbe41ef4c78d851fdec33472217d3"></a>

## hash_algorithms property — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 0ffba7cc7b8a / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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

<a id="canonical-5adcef6ace5912d5961f5482d244cecefac22342edf811bceb12fbd170b72828"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 0ffba7cc7b8a / 5

- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-c98f9d2adbbabdd0b614d25c17d359854de979e945e276b006664993640dcda5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3af848c6c60807adfea2e26040c054b32c3785b6cd9ee2a649877db379d4c26"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling — enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling / bceb4a31f3c8 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-20ce506e7d7ef6b46363af38ce8a8e5418b9c5ce9666350d5f37fb464b018725"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-39eb5896de77c70fb8f94f192c17f5ae67fe582a0d6cf1bf6b48077334f0928a"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling / bceb4a31f3c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-153d5d3175382479a2e78a1a87cb5ac7dc1540aa0e0c121b4b6ef2c1bbd0dc69"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling / bceb4a31f3c8 / 4

- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-a9952a7a1919cece98f4ee174a3419b4ecb530b5e2c835333d46360f23315069"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e5ef73f49ba0d9fa10779dc6e7c2ebfb785910483344f5f7bf0d018365048d6"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.private_key — enable_forward_proxy.tls_intercept.custom_certificate.private_key / a56fbf66291c / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key

<a id="canonical-d5c9829501d4a17e6dcbcf70a2fc38f443283731bfeba42bfca5434fdfc82ced"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-65200e75d366eac0eb0a9becce450d8e6c0e06ad5c66f4ed7f94467924f85fc1"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.private_key / a56fbf66291c / 3

- [blindfold_secret_info](resources--network_connector--reference--group-001.md#canonical-31951a4398accd55cb5572c267006e59d21d6e17964e34b3ff57f21468c3047b): complete subsection reference.

- [clear_secret_info](resources--network_connector--reference--group-001.md#canonical-0e3c74aef34275c9af8617d891e6c229c038920a55cd2fecf5997bb877b6380e): complete subsection reference.

<a id="canonical-6077782fb1894375fc18d8b516d23129696c7b0befa43d0a2f856b0a938c18c2"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.private_key / a56fbf66291c / 4

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--network_connector--reference--group-001.md#canonical-31951a4398accd55cb5572c267006e59d21d6e17964e34b3ff57f21468c3047b)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](resources--network_connector--reference--group-001.md#canonical-0e3c74aef34275c9af8617d891e6c229c038920a55cd2fecf5997bb877b6380e)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-31951a4398accd55cb5572c267006e59d21d6e17964e34b3ff57f21468c3047b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59ae476088a649af8e83b3c10d47bfab58faf744b939ba96da55bc77cb6ca508"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / a8f46e83780c / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-a9952a7a1919cece98f4ee174a3419b4ecb530b5e2c835333d46360f23315069)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-3ce2e5076236aec1e10329473f7a08007fdafe554fc055b5810f717fbdd64ba1"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-cd2c99a19320caf3e8061c457531c42f94382f435fc7b7727eb38b04a4580be0"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / a8f46e83780c / 3

<a id="canonical-16a6a0abf03c2d62e4f4b8e079c177ce8809ca7dd82612a483a80b11423e5635"></a>

<a id="canonical-ab04de2423b109e663dc382e465f8e7536e75f6e03c79818da9f9967e4836259"></a>

## decryption_provider property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / a8f46e83780c / 4

Type: `"string"`. Optional.

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

<a id="canonical-db72773d4726e007f9ab00cce295e432a2280911a07b67817999be770cb46df5"></a>

<a id="canonical-ccaef16bc63e95fd937b08998c8ab45c1cb8a6c8d9f219283afcb4d164abf7d9"></a>

## location property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / a8f46e83780c / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-c33a7a3b597de5b071f238519df1d24d406b5dece4263c1ea746384ce2471326"></a>

<a id="canonical-23b519ca70a2d0c2023dd531b2f8e2026331fcf24c56d1303db9de409138caa5"></a>

## store_provider property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / a8f46e83780c / 6

Type: `"string"`. Optional.

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

<a id="canonical-d42f3771078179b6b2a184bc5e099592de7e37bca47063d90ce57c9fc8ad10d0"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / a8f46e83780c / 7

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-a9952a7a1919cece98f4ee174a3419b4ecb530b5e2c835333d46360f23315069)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-0e3c74aef34275c9af8617d891e6c229c038920a55cd2fecf5997bb877b6380e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ca4a85c526d61af17a88f5fe28cf5efe44ef5d8f924f6bf39df8a08a4b2b1ee"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / b7c32247c510 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-a9952a7a1919cece98f4ee174a3419b4ecb530b5e2c835333d46360f23315069)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-e9068d07b9f5ee276dfe7f4ba8420c86bc84101a7c4b11f86f49cba1dbf940a5"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-662723055a95aa7563a67ae9a813493c2eaff602e9b4d9e6893413eebfafd83b"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / b7c32247c510 / 3

<a id="canonical-b7c8a8f697dee3e326ec8bbb0af65e6b04e2c7a551f12fed7a82e77fe4d91dbd"></a>

<a id="canonical-d14271792371083ddc96a4d6af8bd3bf1624148949e5c6f48db534ca7e64537a"></a>

## provider_ref property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / b7c32247c510 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-159b5fff11d08a4d33ddb1a04777a117431585a8d68e09bf63e5f710daed90a6"></a>

<a id="canonical-5bbf59bfad90e27ff0c7a352d3fb5db48b6f2248a192ae1b3d4a13290f823ef1"></a>

## url property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / b7c32247c510 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-552dda8a86b5f8df0a839e6a2171b621f0b2fe0368c6459152cc133463178ae6"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / b7c32247c510 / 6

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-a9952a7a1919cece98f4ee174a3419b4ecb530b5e2c835333d46360f23315069)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-b2ec67c0cbebc681952b7b6b07a58b913a12db46ed1679814f081f7b6182a592"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-591b3d581f18d95d48c1f383eab79fd06de3c719d702592cd20365f1b1536107"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults — enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults / 7ad99337ac51 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-881f7d177765805af46c3ed69504e99f494e1a5afa09a47eabd4e8440b677671"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-ae514164ec69d51a4a1a355789755ab91461d38ac5e9ab2395f08373ff3c75e0"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults / 7ad99337ac51 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9454224f3e0afa720ba43e20d5d5b0c529d4ab1d352bd0fea5d732e62398471"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults / 7ad99337ac51 / 4

- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-c1bd5735be06dd0e066775b0a168d8f70f013ab9211204526e3307facc479aef)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-34a3a496cb716ab5261674a35b0e6426d62ea6c990724e20dde6e4a5f01645a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-519049f73802a7ca0f889ad044fcb931c9dfa349ad2f07bdcf9d22dafd40a1d3"></a>

## enable_forward_proxy.tls_intercept.enable_for_all_domains — enable_forward_proxy.tls_intercept.enable_for_all_domains / 13bfb29058a1 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- enable_forward_proxy.tls_intercept.enable_for_all_domains

<a id="canonical-972bd990316c16256f7c14530251b93c61bf1767e1336423ada15446d7b538f3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable for all domains.

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
enable_for_all_domains = {}
```

<a id="canonical-d175816fae7e0c5d0734b36cb490906a7738ab99b210e3e673eaf1ba3a692248"></a>

## Direct properties — enable_forward_proxy.tls_intercept.enable_for_all_domains / 13bfb29058a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-93b8ae821255450f8abb2d58c49d8f9d5e61138c91849368be6d323049d6de16"></a>

## Next pages — enable_forward_proxy.tls_intercept.enable_for_all_domains / 13bfb29058a1 / 4

- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81018110764fee01f1f697c0be14da54af3050379f7c71efc665dfea7efa7296"></a>

## enable_forward_proxy.tls_intercept.policy — enable_forward_proxy.tls_intercept.policy / ccab85f030bf / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- enable_forward_proxy.tls_intercept.policy

<a id="canonical-002cd13b4f1def10788bcaf9aa33a585c2ddff9337bffaa4922cdae7a37253c6"></a>

Type: `"object"`. single nested block, Optional.

Policy to enable or disable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interception_rules")}
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
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-959684426ba3385a516362a664193c209639efb51b8d0bfbf802081da1289a13"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy / ccab85f030bf / 3

- [interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d): complete subsection reference.

<a id="canonical-79669f469eeeb2e5ab90133b80de75120d22092f6b64cd304dd24e5f3d0320e6"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy / ccab85f030bf / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faaa621035ce62f7f3fc59befeaa527045c4e85a5697e217543fcac43a59d168"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules — enable_forward_proxy.tls_intercept.policy.interception_rules / e55c6761ca5b / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697)
- enable_forward_proxy.tls_intercept.policy.interception_rules

<a id="canonical-746bf6d9211f9ce79f9753a197afbaa814fd086233dac9da66425e568d69fc1a"></a>

Type: `"object"`. list nested block, Optional.

List of ordered rules to enable or disable for TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("disable_interception",
    "enable_interception")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interception_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-19168fc72c37145a7508080db7aade374ffde592c78716ae065548fc7a5a8828"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules / e55c6761ca5b / 3

- [disable_interception](resources--network_connector--reference--group-001.md#canonical-f2416378ced9d34f394a0094cdf7559450227304a4bc392240c0e5532c6bd7d3): complete subsection reference.

- [domain_match](resources--network_connector--reference--group-001.md#canonical-34805416e8984c1abb2c8cb52bffc51a72b0706a2ce8c614fe04caf1ac965b29): complete subsection reference.

- [enable_interception](resources--network_connector--reference--group-001.md#canonical-afe58353a9f0f2e8f3b7d5c58ec0c497b808167a7835e07cdaa597c1a909b280): complete subsection reference.

<a id="canonical-368b6b93efa79ddfa64d470c3191a0b048c48006a53c7b2fa84e53c097f68a60"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules / e55c6761ca5b / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](resources--network_connector--reference--group-001.md#canonical-f2416378ced9d34f394a0094cdf7559450227304a4bc392240c0e5532c6bd7d3)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](resources--network_connector--reference--group-001.md#canonical-34805416e8984c1abb2c8cb52bffc51a72b0706a2ce8c614fe04caf1ac965b29)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](resources--network_connector--reference--group-001.md#canonical-afe58353a9f0f2e8f3b7d5c58ec0c497b808167a7835e07cdaa597c1a909b280)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-f2416378ced9d34f394a0094cdf7559450227304a4bc392240c0e5532c6bd7d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7074fc940e07fc6113f2aa8dac70a86832dd935e87e4f39e5eb47dc821b167b3"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception — enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interceptio / b9203782087a / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d)
- enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-df9c73a728504e9993a66b6f63769b9349ae77b64b2af7245bdbfdfc8a200d8f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable interception.

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
disable_interception = {}
```

<a id="canonical-951b2438af70eec1f678f59a31ddf79678adbbe7f4f35b24591c8ebbc111244a"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interceptio / b9203782087a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fac6e8357bc47cf69b60dbda482b2d65beddae56389681ecdb5d4780d03c9d5a"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interceptio / b9203782087a / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-34805416e8984c1abb2c8cb52bffc51a72b0706a2ce8c614fe04caf1ac965b29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b68aaba3b52b9d3c3184df75be6000cfbf76a1951c01227e237d68a838fd12b"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / be032eaea3aa / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d)
- enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match

<a id="canonical-741df57b10998b9b855a54a95746f5d38c912b8b81c74f940fc027234e2e64db"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for domain match.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain_match {
  # Configure direct properties listed below.
}
```

<a id="canonical-d2bd525517fbd6d7b06bfbebc4158d4548453f5999a2478fd78900cd8a364d61"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / be032eaea3aa / 3

<a id="canonical-c88b91203e7df2b5c086432046e82a27943a6edf5975e1f92656d8f1e2724aeb"></a>

<a id="canonical-416772293c5c42a77cdb07431a420e3827f3c0e1bf77cac9ba395034603a4dd9"></a>

## exact_value property — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / be032eaea3aa / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-e7979cac5549dac2b1d05256b88f8e39154111de4065ff099ba754b3e3de61aa"></a>

<a id="canonical-3918ef4e315063efdef70e6638d906a24338c4822552f91e956ddb92f2496efd"></a>

## regex_value property — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / be032eaea3aa / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-91ac5444d552aa6d00795a94746d9a404173fe0bff2bf1d6df50276f0cff13ca"></a>

<a id="canonical-d7bf1ba9c17081c5e994bbe281843a3b4714af343b8f6f1fddeb1f3a03fcf192"></a>

## suffix_value property — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / be032eaea3aa / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-eee9e0c68a4c11e9566d6ca18e94195b159051e06752040dfddaf57a7a2d0aa1"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / be032eaea3aa / 7

- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-afe58353a9f0f2e8f3b7d5c58ec0c497b808167a7835e07cdaa597c1a909b280"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be41c0623915d98b5cc6abf31411e01cdcf5de8604274d0757456864d0142af6"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception — enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception / 49395cc3e110 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-e5995e21aa688e27787eba7b01d5c556902c278237ea59a5a4116166c0903697)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d)
- enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-01c0915f92e1fa24b34594795c26d731dbd193616c619e351d3901ebb6eb597c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable interception.

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
enable_interception = {}
```

<a id="canonical-921cde98cbdd8cec77c1d9fd2adddcec4e675a9c1ebf5698ae035f3006235e2f"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception / 49395cc3e110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be785e27c18a2744f0f5c48c6929485c580cf225897906ecb399e886aa9bbb7e"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception / 49395cc3e110 / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-a67aa36b8a5dda3aff3b5adc6325de4f71bf918a164d4c12df392b37f931298d)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-1c32883c7fa89f245fc065a03ba8a8eba7c80b2d9246aab37530cea789583fa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebf0f7057f398b9b279dad111474ec24f59533ecf9ec0fe3369d873a234e7fb4"></a>

## enable_forward_proxy.tls_intercept.volterra_certificate — enable_forward_proxy.tls_intercept.volterra_certificate / ba9fc81f3f4b / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- enable_forward_proxy.tls_intercept.volterra_certificate

<a id="canonical-6fc0ec2d69d2cd7aa7be54756166e77b7362cc62db2facf535737978b805cff5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra certificate.

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
volterra_certificate = {}
```

<a id="canonical-99973da48816cc6fe4a54a541f4e0fce5638e4be1f34dc861b58f8ce0a3f8498"></a>

## Direct properties — enable_forward_proxy.tls_intercept.volterra_certificate / ba9fc81f3f4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-350013a944a0aad4c89d15c8b738b5be0819cd49a84adfa3af4d3c904ccb6d3d"></a>

## Next pages — enable_forward_proxy.tls_intercept.volterra_certificate / ba9fc81f3f4b / 4

- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-b96a37769a8b8de74a47ece1d0faad8e85504cfff33cc302d7b361e97420223c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-598e9fd98fdc699562f8985fafe883fb4760e9530a3a83fc28b10ed0d2ca5b3f"></a>

## enable_forward_proxy.tls_intercept.volterra_trusted_ca — enable_forward_proxy.tls_intercept.volterra_trusted_ca / 38fcd7df6ee8 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-5150cdf7929b3c15387072a2f595b6b4a04ee971ebab0176a26c05a7a8d61e48)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- enable_forward_proxy.tls_intercept.volterra_trusted_ca

<a id="canonical-a6f52b683ccdae05b8eeb69656d8a5a0ef1f8b9f46de3ab2aa285fac6fc40602"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

<a id="canonical-1fede5c78fbf6b6f198079397db1469628eed9d9fd6c6b666d0531d794c87f55"></a>

## Direct properties — enable_forward_proxy.tls_intercept.volterra_trusted_ca / 38fcd7df6ee8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a97081b88adc7effbe3bbe1b581c77dc55cc88b334cd0b9768b6f8507346f622"></a>

## Next pages — enable_forward_proxy.tls_intercept.volterra_trusted_ca / 38fcd7df6ee8 / 4

- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-53c08bc1ef49e5e210b53d959f63a154e31803bbcfb1bfc780a56575b26ce697)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-7ba0153a60944bef2faf88462dc67925427df3cb0c4cf89522c40b3b34f2d1b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9728249d689bd7418057d4af3ec3be9b9faa65c83cd5fd2c905b551cd1615b1e"></a>

## sli_to_global_dr — sli_to_global_dr / a58442aadd2e / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- sli_to_global_dr

<a id="canonical-dddd546acf32c1b16ad0fdfc09b6a93192e79b6cd991255f57e87b5bd083d5a2"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: sli\_to\_global\_dr, sli\_to\_slo\_snat, slo\_to\_global\_dr\] Global network reference for
direct connection.

Upstream description:

Global network reference for direct connection.

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

- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-dddd546acf32c1b16ad0fdfc09b6a93192e79b6cd991255f57e87b5bd083d5a2)
- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-b1c8a6864018866d3cbb36eb67b3ab024ec044d69d025bceb19409e4bc08796e)
- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-3a266d5cae820174cd7e4180ae374bdf6eb94750daec738d1aa5ef874e156a63)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1c6ea00abc9242d1be327422f319f6323971911a2fb9a2cc228ec3840c20fbbc"></a>

## Direct properties — sli_to_global_dr / a58442aadd2e / 3

- [global_vn](resources--network_connector--reference--group-001.md#canonical-68ecb93ab5d952512f999e649e4326364898fe5f60b6fad1e6ed3dd8ed46e00a): complete subsection reference.

<a id="canonical-4ea9476d9f9ed7de44ff112cfe6ad6d550b45a625b3bf425534c3942c7ac88dd"></a>

## Next pages — sli_to_global_dr / a58442aadd2e / 4

- [sli_to_global_dr.global_vn](resources--network_connector--reference--group-001.md#canonical-68ecb93ab5d952512f999e649e4326364898fe5f60b6fad1e6ed3dd8ed46e00a)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-68ecb93ab5d952512f999e649e4326364898fe5f60b6fad1e6ed3dd8ed46e00a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7393d535b0c6b5f629262b008a6bae1aa96e19f6c0072725df923f54451b199f"></a>

## sli_to_global_dr.global_vn — sli_to_global_dr.global_vn / 58d51bee2b50 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-7ba0153a60944bef2faf88462dc67925427df3cb0c4cf89522c40b3b34f2d1b8)
- sli_to_global_dr.global_vn

<a id="canonical-a1bd917e7ef7d2eba215b6b1965b8c412bbf5662fab9cdbd6b25785487b45a2e"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-8968a202033e605c2d0c990caad84de7365a46e60dcf65b8c9146aa0db244de6"></a>

## Direct properties — sli_to_global_dr.global_vn / 58d51bee2b50 / 3

<a id="canonical-fd749c9c27657fd50752c74eed8343bf628e56459288170f8ff622cc9953fa02"></a>

<a id="canonical-ffbfafb6dcf44288fe128652317afcb9f34b500b86280784735c3ec0730d359d"></a>

## name property — sli_to_global_dr.global_vn / 58d51bee2b50 / 4

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

<a id="canonical-4f5fed7331f07d8d31af3eef83c57bc6cac99a1994820f454ea87b60b7101d00"></a>

<a id="canonical-3f38022706f68c71efc21b172fc9e1708bfafce06b4a7f9ef752cde27a661418"></a>

## namespace property — sli_to_global_dr.global_vn / 58d51bee2b50 / 5

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

<a id="canonical-12999551141cc474a7f0eb804e63c43e618b38983668aa89b194aefb47b1d7bc"></a>

<a id="canonical-547dab5a9ba64b126dc7d8851eb8f8af40d490a6da6b8bd13bfe8cc022e55145"></a>

## tenant property — sli_to_global_dr.global_vn / 58d51bee2b50 / 6

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

<a id="canonical-80d36c87dad861c4a9803cdacfd9aa3d24fcf4df6575378bfa5c75e714f00fcc"></a>

## Next pages — sli_to_global_dr.global_vn / 58d51bee2b50 / 7

- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-7ba0153a60944bef2faf88462dc67925427df3cb0c4cf89522c40b3b34f2d1b8)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-35b27d234276bca17116a2d1a66b0e4993dea66284127d3bd54048641abe1489"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-109b6455d6265460531c21a3220139162ee566571cf6c84c8674e118b04b2550"></a>

## sli_to_slo_snat — sli_to_slo_snat / 4dfe4e4a7bdf / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- sli_to_slo_snat

<a id="canonical-b1c8a6864018866d3cbb36eb67b3ab024ec044d69d025bceb19409e4bc08796e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sli to slo snat.

Upstream description:

X-example: "" description.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"interface_ip\"]",
  "x-ves-oneof-field-routing_choice": "[\"default_gw_snat\"]"
}
```

Terraform syntax:

```terraform
sli_to_slo_snat {
  # Configure direct properties listed below.
}
```

<a id="canonical-cf094541428d106bf81f8afb4bb12a7744b89305b44e540895c64f0469b3398b"></a>

## Direct properties — sli_to_slo_snat / 4dfe4e4a7bdf / 3

- [default_gw_snat](resources--network_connector--reference--group-001.md#canonical-3c57c172982d1ae281345719ca71ddce4591cef0e80d4c1224e6f0423d91f90e): complete subsection reference.

- [interface_ip](resources--network_connector--reference--group-001.md#canonical-0f82fdf8421fa05e859cd2fcc3c515aa59a75788a11fa3f3684e5b249f8c48c5): complete subsection reference.

<a id="canonical-7ead877921a11c450d303f701b1b57194dc3dd2c695d06315db9e590bdc190f8"></a>

## Next pages — sli_to_slo_snat / 4dfe4e4a7bdf / 4

- [sli_to_slo_snat.default_gw_snat](resources--network_connector--reference--group-001.md#canonical-3c57c172982d1ae281345719ca71ddce4591cef0e80d4c1224e6f0423d91f90e)
- [sli_to_slo_snat.interface_ip](resources--network_connector--reference--group-001.md#canonical-0f82fdf8421fa05e859cd2fcc3c515aa59a75788a11fa3f3684e5b249f8c48c5)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-3c57c172982d1ae281345719ca71ddce4591cef0e80d4c1224e6f0423d91f90e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93c6af6da52db10e88eaf2601cdbf46c703e57d631224cece71600db5527199b"></a>

## sli_to_slo_snat.default_gw_snat — sli_to_slo_snat.default_gw_snat / 5825dba20684 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-35b27d234276bca17116a2d1a66b0e4993dea66284127d3bd54048641abe1489)
- sli_to_slo_snat.default_gw_snat

<a id="canonical-537cec77c5fff3b8222c411f991c18505b4de52f0279b8a32d17747e3cb9af5a"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default gw snat.

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
default_gw_snat {}
```

<a id="canonical-0e721fb2ef1110d60b8b981ca96506fcee0ea2310ef70666b55466c825c55221"></a>

## Direct properties — sli_to_slo_snat.default_gw_snat / 5825dba20684 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3ba8af185a54ec046aafcbd7116c6f61a30e2d87148bcbade6789ac4dfd4fa5"></a>

## Next pages — sli_to_slo_snat.default_gw_snat / 5825dba20684 / 4

- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-35b27d234276bca17116a2d1a66b0e4993dea66284127d3bd54048641abe1489)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-0f82fdf8421fa05e859cd2fcc3c515aa59a75788a11fa3f3684e5b249f8c48c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-313a1983dd3551ba880a61607cc50e4d029af94b50d083f8db1ead8faad6bd0d"></a>

## sli_to_slo_snat.interface_ip — sli_to_slo_snat.interface_ip / 1fc0cb488a74 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-35b27d234276bca17116a2d1a66b0e4993dea66284127d3bd54048641abe1489)
- sli_to_slo_snat.interface_ip

<a id="canonical-1ea2269efc3590842da750c01a3c9a5db5222cbf058888cca862ca365e54190e"></a>

Type: `"object"`. single nested block, Optional.

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
interface_ip {}
```

<a id="canonical-db6c5374128f988853dc5a0b92b3914095d6c7c9c2364e39f4d4eb31a1a6193b"></a>

## Direct properties — sli_to_slo_snat.interface_ip / 1fc0cb488a74 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89a2c104d7a4cb16d37d7596f03b01e7e37a2964ef60c5ab110ad6f77b251dde"></a>

## Next pages — sli_to_slo_snat.interface_ip / 1fc0cb488a74 / 4

- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-35b27d234276bca17116a2d1a66b0e4993dea66284127d3bd54048641abe1489)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-fcd5ed60be849b0d67ff6027dbdf14f6d887bb7ff3b43e5abb8a217e7a6e1315"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4aa1f0a6f23f0f459c3a34f643e4c66817d30084ff72d60ebf4ba10ba77887e4"></a>

## slo_to_global_dr — slo_to_global_dr / e75634185401 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- slo_to_global_dr

<a id="canonical-3a266d5cae820174cd7e4180ae374bdf6eb94750daec738d1aa5ef874e156a63"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-282c389312f627862569e380ede7f325229d4a52fa8cb02edce7b25368d06e51"></a>

## Direct properties — slo_to_global_dr / e75634185401 / 3

- [global_vn](resources--network_connector--reference--group-001.md#canonical-a77ea49944f3b6116358f1cb22c1f617738fc39731d10dc17bc8193f6e36bb86): complete subsection reference.

<a id="canonical-6362cb3b849acc73929408a8833055d678395a0fc34d9aaf323d9343984011b1"></a>

## Next pages — slo_to_global_dr / e75634185401 / 4

- [slo_to_global_dr.global_vn](resources--network_connector--reference--group-001.md#canonical-a77ea49944f3b6116358f1cb22c1f617738fc39731d10dc17bc8193f6e36bb86)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-a77ea49944f3b6116358f1cb22c1f617738fc39731d10dc17bc8193f6e36bb86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e56c3c31fefed0c75b3299cc4454351a834c70c88bdace64c3fae455e906f7e3"></a>

## slo_to_global_dr.global_vn — slo_to_global_dr.global_vn / 868801a498d4 / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-fcd5ed60be849b0d67ff6027dbdf14f6d887bb7ff3b43e5abb8a217e7a6e1315)
- slo_to_global_dr.global_vn

<a id="canonical-8d7075e4450446d60b2552feada053a61b0c2e04f46f60ac5ef3de7477bfe24c"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-94f0b9c20e9afd8e9b4f8ec5dcbdcc0f98167062b021972e0f3dc72da67dd972"></a>

## Direct properties — slo_to_global_dr.global_vn / 868801a498d4 / 3

<a id="canonical-9af0cb8d0cafc3c87c13b78b2a1f564d8b8af4a59f28569a733b035c56a986c1"></a>

<a id="canonical-8aee6ab6bfb398662ef061c5a91501a0fe51b52db0fc61ffa0d063d838ebd3e9"></a>

## name property — slo_to_global_dr.global_vn / 868801a498d4 / 4

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

<a id="canonical-8f93b2fa1e6fed7c1bbfdb5dcff5dd1f018a09a749c11a8ca7d71352110f4707"></a>

<a id="canonical-af1dc99172f49edce17497fc20364d6deb748d631a3676064f71502ec7310f17"></a>

## namespace property — slo_to_global_dr.global_vn / 868801a498d4 / 5

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

<a id="canonical-b00b91762de98a27c4343de26e3f444e17ca4cfb61b1da316f00ba5a7268d955"></a>

<a id="canonical-82f64bd95e4b1b62de40eaa52a3828d441c4d0dbd4b598426d83217cca9bb57b"></a>

## tenant property — slo_to_global_dr.global_vn / 868801a498d4 / 6

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

<a id="canonical-215497bb50bb511edaa2279e0fd478b5536a90cc13551c388e0159f06d071dd9"></a>

## Next pages — slo_to_global_dr.global_vn / 868801a498d4 / 7

- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-fcd5ed60be849b0d67ff6027dbdf14f6d887bb7ff3b43e5abb8a217e7a6e1315)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)

<a id="canonical-e9cc6a09d7f9a114af512eb9694817eae4f4c08a5d7c900822027349c63c8049"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3229c7a9e5f28704dced954ec9b6e417d1018522b8a4e2f8fa63eb3b35fded2"></a>

## timeouts — timeouts / a1a2e44515af / 2

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- timeouts

<a id="canonical-66170410bdbeffd11beee7aa52dd12c5fbb1900676ac141d427958c92f2e9f2a"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-09c8bd5801bedc0b839199fd91928d96b2222088e7159e4bd282351ddc596a9a"></a>

## Direct properties — timeouts / a1a2e44515af / 3

<a id="canonical-ca2f195192197a0de1b2ec49f589900e572dd921f6fb19c6cab63bc1fb281a1f"></a>

<a id="canonical-e5c4e887cbaca9938aaaf1d40e6c2b0bb9df1c9e40649eaaa11790c35b878cd0"></a>

## create property — timeouts / a1a2e44515af / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-733429a8b77dcd2c4ad70700ed49b0701378652f520def15616dde6d37231d11"></a>

<a id="canonical-4ab02307eeb12e84d492dadce71e105369c17b9c5a73705d122b2949497cf1e6"></a>

## delete property — timeouts / a1a2e44515af / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-588b5ff72b960f12c033ec06436d61a1ce99a6013b806f83cb82936a59550695"></a>

<a id="canonical-0afd9aa37dbddde4c41fc4f904548965b9919186cc8e815d799579f77ecec5dc"></a>

## read property — timeouts / a1a2e44515af / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-d8ba203706ea3c38f701210d91457950937d095999da3e98b2cccac015b4efee"></a>

<a id="canonical-6e2183363fc4a9bfbb6f3c949f59dd280b3fe0a4989a62058e931dd8ba9f0c05"></a>

## update property — timeouts / a1a2e44515af / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-37cbc0dd0bfc4f8f4dbb5f066ed38d9e236834a969deec8ffd2cf9dbab437220"></a>

## Next pages — timeouts / a1a2e44515af / 8

- [Property reference](resources--network_connector--reference--group-001.md#canonical-ec878df0d04ea9f1eac68f80926f55ca939df9125352efce9b144b7bb4ad30fd)
- [xcsh_network_connector](../resources/network_connector.md#canonical-0cb56818b44775bd16a78d7c4430851038331e4d21680537b2f6be63dce8e010)
