---
page_title: "xcsh_network_connector reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector reference."
---

# xcsh_network_connector reference

<a id="canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d90a5f2f65f39dab9ed73e52203a84a2498d3b15149389847598d173d4e0bd06"></a>

## Property reference — Property reference / fbfe390f9ddd / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- Property reference

<a id="canonical-a7cedbd21d04e51dd52ba5c265fb431fa134e35a2d54a990de119afb04dfd592"></a>

## Direct properties — Property reference / fbfe390f9ddd / 3

<a id="canonical-87d4e21c54443e56e9b2c6779446782914bee66a75402848d32b9a3063a610ad"></a>

<a id="canonical-4ab2c4507fc86d4e01cc903716fd363fb3a451cc153f57a101ce81e3c1531484"></a>

## annotations property — Property reference / fbfe390f9ddd / 4

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

<a id="canonical-a1be5cfaa3f633e76b948f501d61e00023a1835cb6662a7b482d3920681aec13"></a>

<a id="canonical-3fcd0f69a9c151a09269c341ecba486cddf73df1c0daca2d2b0d111bb5557760"></a>

## description property — Property reference / fbfe390f9ddd / 5

Type: `"string"`. Computed.

Description of the NetworkConnector.

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

- [disable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-e5cdd6d68c5befd208842254f5433d34fc9a7ec5178dff4ea516bbc8da5a07d4): complete subsection reference.

- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61): complete subsection reference.

<a id="canonical-570df1cd3b0a65c3742fe35b3f995737950efbc9a0ed35c411b4de95d6227e05"></a>

<a id="canonical-e8bd5a2f967c6ebcccb7c56ef7f45787b5dbd74a6155b9c05200804debf1ae6f"></a>

## id property — Property reference / fbfe390f9ddd / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-28574ee33b0a5f44c9c8537ff00868a7dbbbf21844f140dbdc1e6e92c6bed955"></a>

<a id="canonical-00d1fa72500fb0f878269c2e7cfd8a17879503950b636bab3cf161afc26bc5fb"></a>

## labels property — Property reference / fbfe390f9ddd / 7

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

<a id="canonical-4d562be0f6d48d8f9cb74e97c2b714b635e92a771d2d82cac306849ba2277f91"></a>

<a id="canonical-d64c7af0127166d220f696d213dacb15acf75dafac91f1d10748770f17545ba9"></a>

## name property — Property reference / fbfe390f9ddd / 8

Type: `"string"`. Required.

Name of the NetworkConnector.

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

<a id="canonical-913ac07ba5d39d60536f6b625c29050612ee05ece5a64fb17ff29dc7d724c52d"></a>

<a id="canonical-f032b5b27e382944e34066377df9bd56dcf4c361438f40ea6957bd53dbbfc60c"></a>

## namespace property — Property reference / fbfe390f9ddd / 9

Type: `"string"`. Required.

Namespace where the NetworkConnector exists.

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

- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e20c03fc1ee46ed9a1097eab71c5722e73933ac25415cfce2237c5da1019b6ee): complete subsection reference.

- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-196f8edbe7e940a790fca00398917d957d7d922e26f48f544d6a35d2214aa70d): complete subsection reference.

- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e210fe40b954ecc3d4242335bd73ed6f5a6d086948c433a6cb9c2e376e6e384b): complete subsection reference.

<a id="canonical-91cb029e6b2c0ca37bdbed845ecaf033f479310b349798993fbe08fe53688645"></a>

## All schema paths — Property reference / fbfe390f9ddd / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_connector--reference--group-001.md#canonical-87d4e21c54443e56e9b2c6779446782914bee66a75402848d32b9a3063a610ad) |
| `description` | [description](data-sources--network_connector--reference--group-001.md#canonical-a1be5cfaa3f633e76b948f501d61e00023a1835cb6662a7b482d3920681aec13) |
| `disable_forward_proxy` | [disable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-b09ac8af8ffe6a7c1f7433ce9ecce9facb6407eeeb10c420c4917c325c30e876) |
| `enable_forward_proxy` | [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-173e25f3fe5a74026c5b7d186c4f1e7adf1411d6cb2f96dfc6ce2721e8654f1b) |
| `enable_forward_proxy.connection_timeout` | [enable_forward_proxy.connection_timeout](data-sources--network_connector--reference--group-001.md#canonical-088cbd2c44c16ed402a2509252b2d57b91e8d0cc7d7e7f37dddf110f378c03f9) |
| `enable_forward_proxy.max_connect_attempts` | [enable_forward_proxy.max_connect_attempts](data-sources--network_connector--reference--group-001.md#canonical-f99e705c672f3ade419d5e30b9826e46b1d5b1864e9245148920bc9f3921aa4d) |
| `enable_forward_proxy.no_interception` | [enable_forward_proxy.no_interception](data-sources--network_connector--reference--group-001.md#canonical-07fa7d01874ff9f2adaed2c17fb3cc2c6bdc8e726e627234ac96c74104d878a4) |
| `enable_forward_proxy.tls_intercept` | [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-173bb03372f87ec7cb78b23cfa75a632c997180ab60decce15f42c70fe653a25) |
| `enable_forward_proxy.tls_intercept.custom_certificate` | [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-b3efaf03f181611785c1f0d9398ea2e4bdc4c80fee6f019f9b097e7c0c2ea1b4) |
| `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` | [enable_forward_proxy.tls_intercept.custom_certificate.certificate_url](data-sources--network_connector--reference--group-001.md#canonical-ee241f06d88304f807e4df29087afa79c22fca7d7c0df4c751e20408e0207c45) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--network_connector--reference--group-001.md#canonical-f45c27c9ba15f7725a8b64138bedce6063e6b9d8b6c8f85cb3798d18ce5fb6c2) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](data-sources--network_connector--reference--group-001.md#canonical-a19a954b27ed0d3c3c635b0880ac46f780cb8ee420156a7e616b0da31ae01a3a) |
| `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` | [enable_forward_proxy.tls_intercept.custom_certificate.description_spec](data-sources--network_connector--reference--group-001.md#canonical-f8ea2b6aa78ea7dd2908698f62648c18fb693ccb79316858fad07af427d0b2b5) |
| `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` | [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--network_connector--reference--group-001.md#canonical-03733f448b67e0ad584881ced1f6cdf9d1d25e5b184b7f593a2748967e4f3cf0) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-58b9095bd322d38cc4dc4f81314a84df849f0ffc9257397dc68eca4cff998eb8) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--network_connector--reference--group-001.md#canonical-1f21988c04c4cc806083c00820bf2701db51494c2681453738dd301653b5852e) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](data-sources--network_connector--reference--group-001.md#canonical-40471daba87df3e98b064a72ddc31f14512fbc00d240eb14074df8db09ba28f7) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](data-sources--network_connector--reference--group-001.md#canonical-d1c7bfd1c54a51787ab06bb62bbfd62edd08e08990e53e64d3962f8e3ddbfc88) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](data-sources--network_connector--reference--group-001.md#canonical-43faaaa644bcb9cd2caf405ba2833151414c58bfe4f19a9b176a03fb18cadc0f) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--network_connector--reference--group-001.md#canonical-c1de6d3e123956bc55765bbee3ab62915bf3620d0d7e76a2ea3f4bc09b46c599) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](data-sources--network_connector--reference--group-001.md#canonical-baf1f24545150d00efc74f9a808a020ecbe4980aa2da6c7e51ac015a0fcd9dfa) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url](data-sources--network_connector--reference--group-001.md#canonical-35f37b2a4bd63d1ec9761e132c2cecba2605aa1fd3c950a3b6836eeedb6237a2) |
| `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` | [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](data-sources--network_connector--reference--group-001.md#canonical-649ad8d7cce05c740946d24fe6c440f5d438375cfe82c2482f1f47131f9799a9) |
| `enable_forward_proxy.tls_intercept.enable_for_all_domains` | [enable_forward_proxy.tls_intercept.enable_for_all_domains](data-sources--network_connector--reference--group-001.md#canonical-dbac5c50e2df7579980465693be5dcc8fbc15d9c69bcb8c574e89b9b7543e7c2) |
| `enable_forward_proxy.tls_intercept.policy` | [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-c580d27782405dc07b1fa87461e93fd59a9e41766d152f1b485324b75a97a104) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules` | [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-91c74741dd713e71e280f1eaf1e5ded82ebefb32196ad7b397fa7f9908858819) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](data-sources--network_connector--reference--group-001.md#canonical-2d650301e44b9f133e8f79cc43d3bd8700178a351d053a6d12949de5b8441055) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](data-sources--network_connector--reference--group-001.md#canonical-ba7bfaea561277e42b678ed534b4525e617dcc96e6f2249bbe84942287c82ba4) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value](data-sources--network_connector--reference--group-001.md#canonical-6307980f1471c6e22352cd4f7d5e26853e918229b335aa5a007f8a49c947c0b3) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value](data-sources--network_connector--reference--group-001.md#canonical-c7545f95445a311f4ef9a0b614a8095bb0cdf0eba13a06196f89d00ee5416602) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value](data-sources--network_connector--reference--group-001.md#canonical-41e30f749a040b18419762f724fb2b1ab780f9695df1e985e9a3c31bd5ec69ef) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](data-sources--network_connector--reference--group-001.md#canonical-86cf46daaf0a94209f931a523db3adc86fd40e180a467b95991e800412446d5d) |
| `enable_forward_proxy.tls_intercept.trusted_ca_url` | [enable_forward_proxy.tls_intercept.trusted_ca_url](data-sources--network_connector--reference--group-001.md#canonical-8188dc38f5a4f48964edaeb81f3d4940cfe809b13ce943819b4a184a69573ce0) |
| `enable_forward_proxy.tls_intercept.volterra_certificate` | [enable_forward_proxy.tls_intercept.volterra_certificate](data-sources--network_connector--reference--group-001.md#canonical-cb323c666de4c09bf9f415d8db982f7c6af330b577b964627fad3aa6005136d7) |
| `enable_forward_proxy.tls_intercept.volterra_trusted_ca` | [enable_forward_proxy.tls_intercept.volterra_trusted_ca](data-sources--network_connector--reference--group-001.md#canonical-457b9ea96c64f392b25ebc2aaf0ad2dc6aa39af0f8a8e10754275df17cad9bcf) |
| `enable_forward_proxy.white_listed_ports` | [enable_forward_proxy.white_listed_ports](data-sources--network_connector--reference--group-001.md#canonical-65eecc3515d24703795abb0e98197cdf29e78b6c0f7285d6cb10ec8dcdfdce60) |
| `enable_forward_proxy.white_listed_prefixes` | [enable_forward_proxy.white_listed_prefixes](data-sources--network_connector--reference--group-001.md#canonical-cab92d4149ee992108625e27f3282975a11abaa2dc47d72d9f426c379204172e) |
| `id` | [id](data-sources--network_connector--reference--group-001.md#canonical-570df1cd3b0a65c3742fe35b3f995737950efbc9a0ed35c411b4de95d6227e05) |
| `labels` | [labels](data-sources--network_connector--reference--group-001.md#canonical-28574ee33b0a5f44c9c8537ff00868a7dbbbf21844f140dbdc1e6e92c6bed955) |
| `name` | [name](data-sources--network_connector--reference--group-001.md#canonical-4d562be0f6d48d8f9cb74e97c2b714b635e92a771d2d82cac306849ba2277f91) |
| `namespace` | [namespace](data-sources--network_connector--reference--group-001.md#canonical-913ac07ba5d39d60536f6b625c29050612ee05ece5a64fb17ff29dc7d724c52d) |
| `sli_to_global_dr` | [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-832177e12d058c23da8c5bce4c3732c8032f9536c2b5c2bd4817faf31a080995) |
| `sli_to_global_dr.global_vn` | [sli_to_global_dr.global_vn](data-sources--network_connector--reference--group-001.md#canonical-faa689fbeb794d5f70fc4b1728206214f3fff07130dc1a2d77cc2326c0c632cf) |
| `sli_to_global_dr.global_vn.name` | [sli_to_global_dr.global_vn.name](data-sources--network_connector--reference--group-001.md#canonical-07fc227208e774d6d769da4af5cc526dda2fca530fc16bd9db06469db2d106f3) |
| `sli_to_global_dr.global_vn.namespace` | [sli_to_global_dr.global_vn.namespace](data-sources--network_connector--reference--group-001.md#canonical-63d41e74dc87158af4c4e3230587a5c7a94c6e922f01a2ddfd1ea8df534d3050) |
| `sli_to_global_dr.global_vn.tenant` | [sli_to_global_dr.global_vn.tenant](data-sources--network_connector--reference--group-001.md#canonical-bfad9e82eb73555d676d63cc91bc585e56eb811800a003d8010652daedeac201) |
| `sli_to_slo_snat` | [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-980d7501e4b1973daa520a60f9f41d71bb28e59a27ca92163b39010bdf539fe7) |
| `sli_to_slo_snat.default_gw_snat` | [sli_to_slo_snat.default_gw_snat](data-sources--network_connector--reference--group-001.md#canonical-b4d9fefcfb058e50ee223a5e8d99e001939d7da075a879d28867c43452fa99e6) |
| `sli_to_slo_snat.interface_ip` | [sli_to_slo_snat.interface_ip](data-sources--network_connector--reference--group-001.md#canonical-add916a20ac5b6cfec43c057dfeca867f225fa3c1f932986cd9d31e4ec13f9c6) |
| `slo_to_global_dr` | [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-d788fe9acb9fcec352e96d84c696226160c897f7de12824fe84ae34ea9b32c10) |
| `slo_to_global_dr.global_vn` | [slo_to_global_dr.global_vn](data-sources--network_connector--reference--group-001.md#canonical-065d1f8269b39f4ccde8aa07e0aa57cbde8eed95500f041ac9b270fa1c06c7aa) |
| `slo_to_global_dr.global_vn.name` | [slo_to_global_dr.global_vn.name](data-sources--network_connector--reference--group-001.md#canonical-4942b7c23681c5149f5b8e714ae6f32d99427423068d445c58325430b2375540) |
| `slo_to_global_dr.global_vn.namespace` | [slo_to_global_dr.global_vn.namespace](data-sources--network_connector--reference--group-001.md#canonical-965b9df3817c1cf15d1bc6b8dd0baf69fb26f2293c3fd552affb579a31be9445) |
| `slo_to_global_dr.global_vn.tenant` | [slo_to_global_dr.global_vn.tenant](data-sources--network_connector--reference--group-001.md#canonical-c8eaa3634a3c95d47a74fa177c4752e2d07aabce1955a84ba0cb8212eb643e65) |

<a id="canonical-1576577c8bc9e8862d24d2a11818e3dd18c8121e970bffeed25c6674eb21ad46"></a>

## Next pages — Property reference / fbfe390f9ddd / 11

- [disable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-e5cdd6d68c5befd208842254f5433d34fc9a7ec5178dff4ea516bbc8da5a07d4)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e20c03fc1ee46ed9a1097eab71c5722e73933ac25415cfce2237c5da1019b6ee)
- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-196f8edbe7e940a790fca00398917d957d7d922e26f48f544d6a35d2214aa70d)
- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e210fe40b954ecc3d4242335bd73ed6f5a6d086948c433a6cb9c2e376e6e384b)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-e5cdd6d68c5befd208842254f5433d34fc9a7ec5178dff4ea516bbc8da5a07d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90c0efe87198de03ec157d807666737f1dc4740b7f4255c6f240a4cac8691586"></a>

## disable_forward_proxy — disable_forward_proxy / 0a0e6471765c / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- disable_forward_proxy

<a id="canonical-b09ac8af8ffe6a7c1f7433ce9ecce9facb6407eeeb10c420c4917c325c30e876"></a>

Type: `["object", {}]`. Computed.

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

- [disable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-b09ac8af8ffe6a7c1f7433ce9ecce9facb6407eeeb10c420c4917c325c30e876)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-173e25f3fe5a74026c5b7d186c4f1e7adf1411d6cb2f96dfc6ce2721e8654f1b)

Select alternatives according to the provider validators above.

<a id="canonical-59c9f8e044264234cf3c18a86b3ea754b36f006e2b8ecc7dc8bab59b47441883"></a>

## Direct properties — disable_forward_proxy / 0a0e6471765c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d118b336ea0f14c1a9ad4ea81a08473e7b36e42adb6bd4247bf91c181f8e852"></a>

## Next pages — disable_forward_proxy / 0a0e6471765c / 4

- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d67253a03cde14b0cf2175f9100b0f450da65cf5149ab63ca12359a2baf09b5"></a>

## enable_forward_proxy — enable_forward_proxy / 9ee84e59b828 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- enable_forward_proxy

<a id="canonical-173e25f3fe5a74026c5b7d186c4f1e7adf1411d6cb2f96dfc6ce2721e8654f1b"></a>

Type: `"single"`. Computed.

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

<a id="canonical-ec47c35ab82ccdd77c440ab6c3201cbc921b439159026ad7f20ef03a5dee29c6"></a>

## Direct properties — enable_forward_proxy / 9ee84e59b828 / 3

<a id="canonical-088cbd2c44c16ed402a2509252b2d57b91e8d0cc7d7e7f37dddf110f378c03f9"></a>

<a id="canonical-144afd696deb7085f2ff952c1417afd5bc944af5d20fdace65b5a6fe39a623fe"></a>

## connection_timeout property — enable_forward_proxy / 9ee84e59b828 / 4

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

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

<a id="canonical-f99e705c672f3ade419d5e30b9826e46b1d5b1864e9245148920bc9f3921aa4d"></a>

<a id="canonical-d5a482e36502325420286b51fb5c239180251884407fc5067069b006e411cc8e"></a>

## max_connect_attempts property — enable_forward_proxy / 9ee84e59b828 / 5

Type: `"number"`. Computed.

Specifies the allowed number of retries on connect failure to upstream server. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1.

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

- [no_interception](data-sources--network_connector--reference--group-001.md#canonical-c0b74462f4d3c94397f9b73d29fbbd00e420e8b59d20efdd09b743c33f6f3137): complete subsection reference.

- [tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8): complete subsection reference.

<a id="canonical-65eecc3515d24703795abb0e98197cdf29e78b6c0f7285d6cb10ec8dcdfdce60"></a>

<a id="canonical-b45fcf2e21f230f615cd5a9e3c98eaaff07f93caeca9e8e9dfbecfb9a7235332"></a>

## white_listed_ports property — enable_forward_proxy / 9ee84e59b828 / 6

Type: `["list", "number"]`. Computed.

Traffic to these destination TCP ports is not subjected to protocol parsing Example 'tmate' server
port.

Upstream description:

Traffic to these destination TCP ports is not subjected to protocol parsing Example "tmate" server
port.

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

<a id="canonical-cab92d4149ee992108625e27f3282975a11abaa2dc47d72d9f426c379204172e"></a>

<a id="canonical-5a668cf84f3d05d7f2698e6cf250b1f187b5aa3192df1cc19bb89f568d2509e0"></a>

## white_listed_prefixes property — enable_forward_proxy / 9ee84e59b828 / 7

Type: `["list", "string"]`. Computed.

Traffic to these destination IP prefixes is not subjected to protocol parsing Example 'tmate' server
IP.

Upstream description:

Traffic to these destination IP prefixes is not subjected to protocol parsing Example "tmate" server
IP.

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

<a id="canonical-7e6f52cbe0ac2c811a0c1fac056b06f5ea0ee1cfde3c0b4cd6e7588f04357238"></a>

## Next pages — enable_forward_proxy / 9ee84e59b828 / 8

- [enable_forward_proxy.no_interception](data-sources--network_connector--reference--group-001.md#canonical-c0b74462f4d3c94397f9b73d29fbbd00e420e8b59d20efdd09b743c33f6f3137)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-c0b74462f4d3c94397f9b73d29fbbd00e420e8b59d20efdd09b743c33f6f3137"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be7b0e405fd5a9d22ac123d8462745615e88f2948a67df4a2fad212a7128da3d"></a>

## enable_forward_proxy.no_interception — enable_forward_proxy.no_interception / f1d7532dc2f4 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- enable_forward_proxy.no_interception

<a id="canonical-07fa7d01874ff9f2adaed2c17fb3cc2c6bdc8e726e627234ac96c74104d878a4"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-21487a5251323a92fe60b2977ab49f03dfca2318481206afdd8b28a32dde92bd"></a>

## Direct properties — enable_forward_proxy.no_interception / f1d7532dc2f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b5586d5b4dc837c6aae46f9933a1ee5a59cb3488183fbbd42c5c319405cd32b"></a>

## Next pages — enable_forward_proxy.no_interception / f1d7532dc2f4 / 4

- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-717e0dd6f90303187aeb06871bbb90faef358dfb8b2d06de6b9c185a44de723e"></a>

## enable_forward_proxy.tls_intercept — enable_forward_proxy.tls_intercept / 666b57583c40 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- enable_forward_proxy.tls_intercept

<a id="canonical-173bb03372f87ec7cb78b23cfa75a632c997180ab60decce15f42c70fe653a25"></a>

Type: `"single"`. Computed.

Configuration to enable TLS interception.

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

<a id="canonical-003db841e3f91d28a55eee89c6bdbd93b763ffecfa7fcb6eae1485c05c0d3a7e"></a>

## Direct properties — enable_forward_proxy.tls_intercept / 666b57583c40 / 3

- [custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87): complete subsection reference.

- [enable_for_all_domains](data-sources--network_connector--reference--group-001.md#canonical-1262917e34c495c80b36230b5872d92ece579f1edea631b56b6ea57947e5cb5c): complete subsection reference.

- [policy](data-sources--network_connector--reference--group-001.md#canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575): complete subsection reference.

<a id="canonical-8188dc38f5a4f48964edaeb81f3d4940cfe809b13ce943819b4a184a69573ce0"></a>

<a id="canonical-f348a8b735609d166eb1dd17efb3136407a4a881096616b1af2d9a3125d9a32b"></a>

## trusted_ca_url property — enable_forward_proxy.tls_intercept / 666b57583c40 / 4

Type: `"string"`. Computed.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

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

- [volterra_certificate](data-sources--network_connector--reference--group-001.md#canonical-6b421305ab578d4570dfe7d40cde67f29e3a0f6fed873797975507d50cb744b8): complete subsection reference.

- [volterra_trusted_ca](data-sources--network_connector--reference--group-001.md#canonical-4a5cd7bacfe95191e6a80a35e0a2a564b7b5a29af534493fb7f5fbfe7f5db3c8): complete subsection reference.

<a id="canonical-42462a4bb8f8a7181760ac05c560cb3050121dc74bca05d2265d06485da605d1"></a>

## Next pages — enable_forward_proxy.tls_intercept / 666b57583c40 / 5

- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- [enable_forward_proxy.tls_intercept.enable_for_all_domains](data-sources--network_connector--reference--group-001.md#canonical-1262917e34c495c80b36230b5872d92ece579f1edea631b56b6ea57947e5cb5c)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575)
- [enable_forward_proxy.tls_intercept.volterra_certificate](data-sources--network_connector--reference--group-001.md#canonical-6b421305ab578d4570dfe7d40cde67f29e3a0f6fed873797975507d50cb744b8)
- [enable_forward_proxy.tls_intercept.volterra_trusted_ca](data-sources--network_connector--reference--group-001.md#canonical-4a5cd7bacfe95191e6a80a35e0a2a564b7b5a29af534493fb7f5fbfe7f5db3c8)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-004006c65e86985e32a442f9e9ff0cfb672855e77180b6ec747b1521759f8019"></a>

## enable_forward_proxy.tls_intercept.custom_certificate — enable_forward_proxy.tls_intercept.custom_certificate / 9663419efef5 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- enable_forward_proxy.tls_intercept.custom_certificate

<a id="canonical-b3efaf03f181611785c1f0d9398ea2e4bdc4c80fee6f019f9b097e7c0c2ea1b4"></a>

Type: `"single"`. Computed.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

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

<a id="canonical-1c4603a884dada58131b713661f1cde1ca3bb0f103cdcd295aef835ff9ee1856"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate / 9663419efef5 / 3

<a id="canonical-ee241f06d88304f807e4df29087afa79c22fca7d7c0df4c751e20408e0207c45"></a>

<a id="canonical-53ba0c73549d209f4d0427f8165b8f97dc8623fab5ce701c27ac367b5d1a758b"></a>

## certificate_url property — enable_forward_proxy.tls_intercept.custom_certificate / 9663419efef5 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--network_connector--reference--group-001.md#canonical-75eedcac18b9a3b00d1410118582232acbd0af003dfb5c59460fc6f150bb84cc): complete subsection reference.

<a id="canonical-f8ea2b6aa78ea7dd2908698f62648c18fb693ccb79316858fad07af427d0b2b5"></a>

<a id="canonical-899570e66e1fec4f662eafd8534d377dd12516a439f77dfff2ba1379d86a2fb4"></a>

## description_spec property — enable_forward_proxy.tls_intercept.custom_certificate / 9663419efef5 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--network_connector--reference--group-001.md#canonical-aef888f912b42a96353f4788b6ae6c4758fc1da754bebba28fffa371badd90c5): complete subsection reference.

- [private_key](data-sources--network_connector--reference--group-001.md#canonical-24657f5aabbdea27c8370ca63bf989451516802089a18d848059fdb45c27ff66): complete subsection reference.

- [use_system_defaults](data-sources--network_connector--reference--group-001.md#canonical-b12235093966f4d61a38f31cae6a0b73048bcf9797c783bcb39d776f117521e0): complete subsection reference.

<a id="canonical-e055fc3bb2b7f1391483f38c4a27f1928f66a4948e58120cc75d30bfa4eeef33"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate / 9663419efef5 / 6

- [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--network_connector--reference--group-001.md#canonical-75eedcac18b9a3b00d1410118582232acbd0af003dfb5c59460fc6f150bb84cc)
- [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--network_connector--reference--group-001.md#canonical-aef888f912b42a96353f4788b6ae6c4758fc1da754bebba28fffa371badd90c5)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-24657f5aabbdea27c8370ca63bf989451516802089a18d848059fdb45c27ff66)
- [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](data-sources--network_connector--reference--group-001.md#canonical-b12235093966f4d61a38f31cae6a0b73048bcf9797c783bcb39d776f117521e0)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-75eedcac18b9a3b00d1410118582232acbd0af003dfb5c59460fc6f150bb84cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d393e776d13e605ace067649bbb750b89d47d36dcdc0a156421ae85bc3c51613"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 5d4eeece2ca8 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-f45c27c9ba15f7725a8b64138bedce6063e6b9d8b6c8f85cb3798d18ce5fb6c2"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-0186a74414dc8acf8d3baa9c6824463ae53465eee8cf4a153fe9ab2071b94067"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 5d4eeece2ca8 / 3

<a id="canonical-a19a954b27ed0d3c3c635b0880ac46f780cb8ee420156a7e616b0da31ae01a3a"></a>

<a id="canonical-13904b8d8957d56ecd6c8d13ce4b633ea75316c4dbbcda28e76c65e0f1205819"></a>

## hash_algorithms property — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 5d4eeece2ca8 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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

<a id="canonical-1f90475ffce8eec38cd3c4d2c60f4054c27be6485f9104be0a090ef152d03c79"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms / 5d4eeece2ca8 / 5

- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-aef888f912b42a96353f4788b6ae6c4758fc1da754bebba28fffa371badd90c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d9bb70fab8c1a2da3746dbda6c99d482faf2220f90a9c9b90e972b2e4d0addb"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling — enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling / 329a1985d32d / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-03733f448b67e0ad584881ced1f6cdf9d1d25e5b184b7f593a2748967e4f3cf0"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-203ac80786f1acfa60b705939211fc327f054531ba403892ebd04834b9021b58"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling / 329a1985d32d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cd700bd8dc5cedd0a4aaaf848cd2d1a011a3d3ec4d9b67e1aa03f2392deca4f"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling / 329a1985d32d / 4

- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-24657f5aabbdea27c8370ca63bf989451516802089a18d848059fdb45c27ff66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d95571de6836a85ca465e85d3fb6efd23d3546710094621f107230c27d45a4a7"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.private_key — enable_forward_proxy.tls_intercept.custom_certificate.private_key / 142fd94bd02b / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key

<a id="canonical-58b9095bd322d38cc4dc4f81314a84df849f0ffc9257397dc68eca4cff998eb8"></a>

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

<a id="canonical-8218b8a592eb78969d513e3b33130388eaaa54e66a22030eb9c2cdf067fd1e74"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.private_key / 142fd94bd02b / 3

- [blindfold_secret_info](data-sources--network_connector--reference--group-001.md#canonical-9f2b82ccba5e2d18ddf9741246c17ea606886b123c86e1f5034ce0bd78119255): complete subsection reference.

- [clear_secret_info](data-sources--network_connector--reference--group-001.md#canonical-0afbba09746d6f2d8c94e9ce34079c41c7fc1f1bb226ace8c4e53294291e9c6f): complete subsection reference.

<a id="canonical-dd6e16389b0c15d661e898270f79c25b6c5812add5c4853f64ad2d2938fce8e9"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.private_key / 142fd94bd02b / 4

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--network_connector--reference--group-001.md#canonical-9f2b82ccba5e2d18ddf9741246c17ea606886b123c86e1f5034ce0bd78119255)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--network_connector--reference--group-001.md#canonical-0afbba09746d6f2d8c94e9ce34079c41c7fc1f1bb226ace8c4e53294291e9c6f)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-9f2b82ccba5e2d18ddf9741246c17ea606886b123c86e1f5034ce0bd78119255"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4507821c944308f47cb25eacdd07127d9ba66c995c4866e25022aab7673012f6"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / 0cccd5cf469c / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-24657f5aabbdea27c8370ca63bf989451516802089a18d848059fdb45c27ff66)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-1f21988c04c4cc806083c00820bf2701db51494c2681453738dd301653b5852e"></a>

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

<a id="canonical-af39bc50359130073ec636ba3817f8b9b108333125d3e2f58e9169bc41b77f15"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / 0cccd5cf469c / 3

<a id="canonical-40471daba87df3e98b064a72ddc31f14512fbc00d240eb14074df8db09ba28f7"></a>

<a id="canonical-815d701a51b3314cf68dcfe667c93c5c604b7996560ce42c539a45f9e9013887"></a>

## decryption_provider property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / 0cccd5cf469c / 4

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

<a id="canonical-d1c7bfd1c54a51787ab06bb62bbfd62edd08e08990e53e64d3962f8e3ddbfc88"></a>

<a id="canonical-d9140e6e617efad61eb71847bf9a3089f152ebdb34a30368e0fde0061447bd5e"></a>

## location property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / 0cccd5cf469c / 5

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

<a id="canonical-43faaaa644bcb9cd2caf405ba2833151414c58bfe4f19a9b176a03fb18cadc0f"></a>

<a id="canonical-3dd988a36d8f9a2944744eb16d14d852b287b25994328a84f2c3a93f45279f3b"></a>

## store_provider property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / 0cccd5cf469c / 6

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

<a id="canonical-654421b174219877a2887be63fa4c35b5da76b5306786a7e9bdef4e4a00291a6"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secr / 0cccd5cf469c / 7

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-24657f5aabbdea27c8370ca63bf989451516802089a18d848059fdb45c27ff66)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-0afbba09746d6f2d8c94e9ce34079c41c7fc1f1bb226ace8c4e53294291e9c6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b766fa773365cde06a7e71c53050e2448fb151c8026d185af76f1edd33a00007"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / fecba5c96d6d / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-24657f5aabbdea27c8370ca63bf989451516802089a18d848059fdb45c27ff66)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-c1de6d3e123956bc55765bbee3ab62915bf3620d0d7e76a2ea3f4bc09b46c599"></a>

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

<a id="canonical-086537cd0c88fd0a7b108c11143e3c618c77a094c982abe4e14051df00125553"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / fecba5c96d6d / 3

<a id="canonical-baf1f24545150d00efc74f9a808a020ecbe4980aa2da6c7e51ac015a0fcd9dfa"></a>

<a id="canonical-0f159001b7d6c150b6cbe15118b84105624bfdd030b0289867d2032095057b34"></a>

## provider_ref property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / fecba5c96d6d / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-35f37b2a4bd63d1ec9761e132c2cecba2605aa1fd3c950a3b6836eeedb6237a2"></a>

<a id="canonical-0566c7fc85b65e3a6ba18c343040e20a228c749f534fcfcfb7f5509abb326e4e"></a>

## url property — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / fecba5c96d6d / 5

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

<a id="canonical-e52429b19cd782d6330793b15aea30862261a3079c6136b6347cf1b6195d559d"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_i / fecba5c96d6d / 6

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-24657f5aabbdea27c8370ca63bf989451516802089a18d848059fdb45c27ff66)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-b12235093966f4d61a38f31cae6a0b73048bcf9797c783bcb39d776f117521e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-969eb62af4d2b664a0e0504d043a773b64b93e34a174aebc7bddfd129f0fb0c9"></a>

## enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults — enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults / 7fd8851e3df8 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-649ad8d7cce05c740946d24fe6c440f5d438375cfe82c2482f1f47131f9799a9"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-383982d2510b4215f4f76d02b35a5e8d8fcc3fc2077d8b53f7dbb32ff678275c"></a>

## Direct properties — enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults / 7fd8851e3df8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d42bc0f054e882529405f9020eb9d70e4f952bd198dd3bc39cb7ce5cb137a22"></a>

## Next pages — enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults / 7fd8851e3df8 / 4

- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-123f693636bd77e9ba69687ca80becd11d7c80aeb1b062a87e01edc00ed3fb87)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-1262917e34c495c80b36230b5872d92ece579f1edea631b56b6ea57947e5cb5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e2a395269dee4eced974fcff93f6428358e5006ba36398f62f0d96b0cfa4dc2"></a>

## enable_forward_proxy.tls_intercept.enable_for_all_domains — enable_forward_proxy.tls_intercept.enable_for_all_domains / 8bae0c29cda1 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- enable_forward_proxy.tls_intercept.enable_for_all_domains

<a id="canonical-dbac5c50e2df7579980465693be5dcc8fbc15d9c69bcb8c574e89b9b7543e7c2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b771d2e004bf761988c8196be218c4c1ae33a26b57ecc78f9054db80ed658995"></a>

## Direct properties — enable_forward_proxy.tls_intercept.enable_for_all_domains / 8bae0c29cda1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d029c978761672be172c8ea8463e387ba2b9142aa61c8be6edf087d6f6dc958"></a>

## Next pages — enable_forward_proxy.tls_intercept.enable_for_all_domains / 8bae0c29cda1 / 4

- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9e65b7429fb42585bc72bb2dfa92a6b75ea45a5adcf019180faa56bdb714d23"></a>

## enable_forward_proxy.tls_intercept.policy — enable_forward_proxy.tls_intercept.policy / 214d2beb1473 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- enable_forward_proxy.tls_intercept.policy

<a id="canonical-c580d27782405dc07b1fa87461e93fd59a9e41766d152f1b485324b75a97a104"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

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

<a id="canonical-9616972383a7940e4b0b55481d758226274c1ad66162897f83b14f54933d7a59"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy / 214d2beb1473 / 3

- [interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264): complete subsection reference.

<a id="canonical-543ab89f3b9a639c209cef84b51a254b9cf82454e920867dfae78c412944cc53"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy / 214d2beb1473 / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4031a5ef6512f5666f97a8edd84e262c71669515c1ebac258abdde5dddec3273"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules — enable_forward_proxy.tls_intercept.policy.interception_rules / f5ae57041692 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575)
- enable_forward_proxy.tls_intercept.policy.interception_rules

<a id="canonical-91c74741dd713e71e280f1eaf1e5ded82ebefb32196ad7b397fa7f9908858819"></a>

Type: `"list"`. Computed.

List of ordered rules to enable or disable for TLS interception.

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

<a id="canonical-e6c13ce8fe54e8c3578059bbaa28872a2467f8cbcc3d764c1e19e5b37f5358da"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules / f5ae57041692 / 3

- [disable_interception](data-sources--network_connector--reference--group-001.md#canonical-b9c129af14fda5de32f6db1aa10ab3018b71f2a4083d6fd6e8c64fdf98606bba): complete subsection reference.

- [domain_match](data-sources--network_connector--reference--group-001.md#canonical-1a6ce2824cb8bf0dd3fa7e43a435cac98917513c20b8d71f169810f749a9c336): complete subsection reference.

- [enable_interception](data-sources--network_connector--reference--group-001.md#canonical-55d6a1cffe9e2bf6e164c50945ad330b7e66218404a0fa93968df358b7de7b59): complete subsection reference.

<a id="canonical-30d33a6b0b97c00645b0446af088d1468ddc3ebfb2c709cd833cade117280009"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules / f5ae57041692 / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](data-sources--network_connector--reference--group-001.md#canonical-b9c129af14fda5de32f6db1aa10ab3018b71f2a4083d6fd6e8c64fdf98606bba)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](data-sources--network_connector--reference--group-001.md#canonical-1a6ce2824cb8bf0dd3fa7e43a435cac98917513c20b8d71f169810f749a9c336)
- [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](data-sources--network_connector--reference--group-001.md#canonical-55d6a1cffe9e2bf6e164c50945ad330b7e66218404a0fa93968df358b7de7b59)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-b9c129af14fda5de32f6db1aa10ab3018b71f2a4083d6fd6e8c64fdf98606bba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf743239bc40568c227a9d0ac5cd999350db1ae7fccd2af0540300aac541455a"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception — enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interceptio / 7a0183b46587 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264)
- enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-2d650301e44b9f133e8f79cc43d3bd8700178a351d053a6d12949de5b8441055"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-71cf40966018452de3ef32223befc5894a64cd410821a4ad6d26749848035101"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interceptio / 7a0183b46587 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c65265e7546e437ce2a58ccf7c843cd0a0a3590a1b2b1b1471f510375f6dd56d"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interceptio / 7a0183b46587 / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-1a6ce2824cb8bf0dd3fa7e43a435cac98917513c20b8d71f169810f749a9c336"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-133cc8ba36c2b07904d15ace8f27eab4ef6670e919f2a1f9dd554ce5d73ab514"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / 2a8e1aa9a0a3 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264)
- enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match

<a id="canonical-ba7bfaea561277e42b678ed534b4525e617dcc96e6f2249bbe84942287c82ba4"></a>

Type: `"single"`. Computed.

Configuration parameter for domain match.

Upstream description:

Domains names.

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

<a id="canonical-b880821b14d2811cf93a7cab6a6c9b92a06edeb35c2708992e24b5fca3741093"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / 2a8e1aa9a0a3 / 3

<a id="canonical-6307980f1471c6e22352cd4f7d5e26853e918229b335aa5a007f8a49c947c0b3"></a>

<a id="canonical-6a702fae658db10d982c35ced65090cfca4461857b2024569e25316558a58ce3"></a>

## exact_value property — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / 2a8e1aa9a0a3 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-c7545f95445a311f4ef9a0b614a8095bb0cdf0eba13a06196f89d00ee5416602"></a>

<a id="canonical-1f90832eb89b67cb8bd22d699141c1f57c805fa80cbc65f787ab142c08859e06"></a>

## regex_value property — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / 2a8e1aa9a0a3 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-41e30f749a040b18419762f724fb2b1ab780f9695df1e985e9a3c31bd5ec69ef"></a>

<a id="canonical-113b23a71613d3149dfd62506b3ddab2ed78e10694d3a0a4b1d5368d2aea6845"></a>

## suffix_value property — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / 2a8e1aa9a0a3 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-85d5a372921fee6e4a544f6f66c3bf48212faec3e4707df8d4acaffe28048275"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match / 2a8e1aa9a0a3 / 7

- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-55d6a1cffe9e2bf6e164c50945ad330b7e66218404a0fa93968df358b7de7b59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e2513c5088527fc6ae0b5d75d9b5f315c9987827dc7194a4fe6890d3e7c883c"></a>

## enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception — enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception / f920c00edfb7 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-3aaec5bfa9f31cbac1a4b65c720c67d7c4031a69dc553adc0648c24b4ef68575)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264)
- enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-86cf46daaf0a94209f931a523db3adc86fd40e180a467b95991e800412446d5d"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-cc7e51d1eb1e97bf9ed8c1c5340a84fa543ffe8c8995e744b67f3a3df782df8a"></a>

## Direct properties — enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception / f920c00edfb7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c398ddfdae873b2b6440806e49b1498efa9fe8d7e49928b861153f58a6d55211"></a>

## Next pages — enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception / f920c00edfb7 / 4

- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-d0336ee968bdf829af8ef92a18912ed066690d8a16fc6effdaa0a1ce9ef70264)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-6b421305ab578d4570dfe7d40cde67f29e3a0f6fed873797975507d50cb744b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9c1143848c4cf06be09399fef77e9f4b559850831880799de5b5a33c8ebe82d"></a>

## enable_forward_proxy.tls_intercept.volterra_certificate — enable_forward_proxy.tls_intercept.volterra_certificate / 0dc940cf91bf / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- enable_forward_proxy.tls_intercept.volterra_certificate

<a id="canonical-cb323c666de4c09bf9f415d8db982f7c6af330b577b964627fad3aa6005136d7"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-60a13eb35f5d399c344259c47584482a26d19beabcbd524622d1b251f09b3b46"></a>

## Direct properties — enable_forward_proxy.tls_intercept.volterra_certificate / 0dc940cf91bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-998c85c512ab401bad5e86beb6206267374e7d553bb7186177874261aba58ea2"></a>

## Next pages — enable_forward_proxy.tls_intercept.volterra_certificate / 0dc940cf91bf / 4

- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-4a5cd7bacfe95191e6a80a35e0a2a564b7b5a29af534493fb7f5fbfe7f5db3c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09b5994322c2374d0ef152041b9d77048ec45bfe55a38097a3552ef5e0f24ff7"></a>

## enable_forward_proxy.tls_intercept.volterra_trusted_ca — enable_forward_proxy.tls_intercept.volterra_trusted_ca / 539ce14e1c05 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-a754d64dc3eb31650827dcb427dadfea229b73a80219d3433a45a81271918b61)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- enable_forward_proxy.tls_intercept.volterra_trusted_ca

<a id="canonical-457b9ea96c64f392b25ebc2aaf0ad2dc6aa39af0f8a8e10754275df17cad9bcf"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1b5ea5b6bb7c087910bb619476041daea7fa7fe90d9af7af655c538d914e966b"></a>

## Direct properties — enable_forward_proxy.tls_intercept.volterra_trusted_ca / 539ce14e1c05 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d0ef839685baad210876894aa8f1c71c72dfd3d53eadc32c3ae68bb42255bf2"></a>

## Next pages — enable_forward_proxy.tls_intercept.volterra_trusted_ca / 539ce14e1c05 / 4

- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-3edacebda6f32b7e330ebd439fbd8917d55dd40c03a1d431ff1b1fc02cded6f8)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-e20c03fc1ee46ed9a1097eab71c5722e73933ac25415cfce2237c5da1019b6ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f7f654a4c7b46ab6ef56ac6359aea78741560254f13b401d55394c188836ea3"></a>

## sli_to_global_dr — sli_to_global_dr / 106a9a3c6af3 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- sli_to_global_dr

<a id="canonical-832177e12d058c23da8c5bce4c3732c8032f9536c2b5c2bd4817faf31a080995"></a>

Type: `"single"`. Computed.

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

- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-832177e12d058c23da8c5bce4c3732c8032f9536c2b5c2bd4817faf31a080995)
- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-980d7501e4b1973daa520a60f9f41d71bb28e59a27ca92163b39010bdf539fe7)
- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-d788fe9acb9fcec352e96d84c696226160c897f7de12824fe84ae34ea9b32c10)

Select alternatives according to the provider validators above.

<a id="canonical-3649944c7ce3cbe4ae87a3fff144b1b766750568702ec6d2964bac33c73826df"></a>

## Direct properties — sli_to_global_dr / 106a9a3c6af3 / 3

- [global_vn](data-sources--network_connector--reference--group-001.md#canonical-7a7fd5f75b79255b02a36ee457714cf856100b1aa4612edddba85cc7726a4ccc): complete subsection reference.

<a id="canonical-455c9ce48414cf02414987ea8e68bddc04087402694e91117f5c966e860e3850"></a>

## Next pages — sli_to_global_dr / 106a9a3c6af3 / 4

- [sli_to_global_dr.global_vn](data-sources--network_connector--reference--group-001.md#canonical-7a7fd5f75b79255b02a36ee457714cf856100b1aa4612edddba85cc7726a4ccc)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-7a7fd5f75b79255b02a36ee457714cf856100b1aa4612edddba85cc7726a4ccc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb9f741f7b4ce68c7e55ffe2e1073f2b41c91eb96d4e256c5e22da9beb8ce91e"></a>

## sli_to_global_dr.global_vn — sli_to_global_dr.global_vn / 0243bf9af3d9 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e20c03fc1ee46ed9a1097eab71c5722e73933ac25415cfce2237c5da1019b6ee)
- sli_to_global_dr.global_vn

<a id="canonical-faa689fbeb794d5f70fc4b1728206214f3fff07130dc1a2d77cc2326c0c632cf"></a>

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

<a id="canonical-e1c297b2ea4d6e0f8a8a19c3ca05f967d338ee0b3a57b5e5bd025e35132344c1"></a>

## Direct properties — sli_to_global_dr.global_vn / 0243bf9af3d9 / 3

<a id="canonical-07fc227208e774d6d769da4af5cc526dda2fca530fc16bd9db06469db2d106f3"></a>

<a id="canonical-47d36686bb641f1b3f7b03e39260512757d29eeb4d7897a4c3977c8eac2ee2a9"></a>

## name property — sli_to_global_dr.global_vn / 0243bf9af3d9 / 4

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

<a id="canonical-63d41e74dc87158af4c4e3230587a5c7a94c6e922f01a2ddfd1ea8df534d3050"></a>

<a id="canonical-cf8c60fb4e1cf84a22a78a488b60c707c42432e42ce2c38649a161d81a6dc46d"></a>

## namespace property — sli_to_global_dr.global_vn / 0243bf9af3d9 / 5

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

<a id="canonical-bfad9e82eb73555d676d63cc91bc585e56eb811800a003d8010652daedeac201"></a>

<a id="canonical-de318521947e01c1cb87b198ad160d9901dbb25922eb846547143a6202b5d3ee"></a>

## tenant property — sli_to_global_dr.global_vn / 0243bf9af3d9 / 6

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

<a id="canonical-731cefae9fd7da20119fbd8d852bb3f93145088719c833dd62fdc7fabbbba32e"></a>

## Next pages — sli_to_global_dr.global_vn / 0243bf9af3d9 / 7

- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e20c03fc1ee46ed9a1097eab71c5722e73933ac25415cfce2237c5da1019b6ee)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-196f8edbe7e940a790fca00398917d957d7d922e26f48f544d6a35d2214aa70d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ec3e644b860ec552fcda3eb98d62711ac5a5a9afd8e4197bcbb74293db622ca"></a>

## sli_to_slo_snat — sli_to_slo_snat / b1b1130ef706 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- sli_to_slo_snat

<a id="canonical-980d7501e4b1973daa520a60f9f41d71bb28e59a27ca92163b39010bdf539fe7"></a>

Type: `"single"`. Computed.

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

<a id="canonical-22b11a53c3a01923d93f2f2e68418a259e89a4b1135d779d5124d6f708becf04"></a>

## Direct properties — sli_to_slo_snat / b1b1130ef706 / 3

- [default_gw_snat](data-sources--network_connector--reference--group-001.md#canonical-62c527fd9165cf729a28adc5792259aba753b6c8ec665bacd1221b15bb8151d5): complete subsection reference.

- [interface_ip](data-sources--network_connector--reference--group-001.md#canonical-76fd24d444a05553b3ad16a43b359b0ea124aa396b377a83bd264666dbe81043): complete subsection reference.

<a id="canonical-7c600bedd09a0e4e689b5d7dab60fa6c41c6c03df48e7b4f3d8010c913a22d44"></a>

## Next pages — sli_to_slo_snat / b1b1130ef706 / 4

- [sli_to_slo_snat.default_gw_snat](data-sources--network_connector--reference--group-001.md#canonical-62c527fd9165cf729a28adc5792259aba753b6c8ec665bacd1221b15bb8151d5)
- [sli_to_slo_snat.interface_ip](data-sources--network_connector--reference--group-001.md#canonical-76fd24d444a05553b3ad16a43b359b0ea124aa396b377a83bd264666dbe81043)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-62c527fd9165cf729a28adc5792259aba753b6c8ec665bacd1221b15bb8151d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a339c4aa81401e9bd6b4ebfdd9ef78649b7eabd365175d1fb648327e49108b0e"></a>

## sli_to_slo_snat.default_gw_snat — sli_to_slo_snat.default_gw_snat / 96d65c556023 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-196f8edbe7e940a790fca00398917d957d7d922e26f48f544d6a35d2214aa70d)
- sli_to_slo_snat.default_gw_snat

<a id="canonical-b4d9fefcfb058e50ee223a5e8d99e001939d7da075a879d28867c43452fa99e6"></a>

Type: `"single"`. Computed.

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

<a id="canonical-fd3f671060b34a8057d17231dac577efa276d47175e09ffdc4d65582902acd7f"></a>

## Direct properties — sli_to_slo_snat.default_gw_snat / 96d65c556023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af66229e4f34c92410519d926af85191166ab30935c4c85eb52fe1daebbaae7b"></a>

## Next pages — sli_to_slo_snat.default_gw_snat / 96d65c556023 / 4

- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-196f8edbe7e940a790fca00398917d957d7d922e26f48f544d6a35d2214aa70d)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-76fd24d444a05553b3ad16a43b359b0ea124aa396b377a83bd264666dbe81043"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39ddf85e47a12a3bd58d8cbcb5655c3cfd7299590d4691e186c5ead3b1787ff5"></a>

## sli_to_slo_snat.interface_ip — sli_to_slo_snat.interface_ip / f680dfd900e5 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-196f8edbe7e940a790fca00398917d957d7d922e26f48f544d6a35d2214aa70d)
- sli_to_slo_snat.interface_ip

<a id="canonical-add916a20ac5b6cfec43c057dfeca867f225fa3c1f932986cd9d31e4ec13f9c6"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0aa1caa063e7f4f8b1a0ab2c3a8cc662cc4a0f2ab6f15fb514c054715a0c3ede"></a>

## Direct properties — sli_to_slo_snat.interface_ip / f680dfd900e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d85471d3f70d9a5a5a2074929942ad95b0f915aab7170446106585f284957e4d"></a>

## Next pages — sli_to_slo_snat.interface_ip / f680dfd900e5 / 4

- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-196f8edbe7e940a790fca00398917d957d7d922e26f48f544d6a35d2214aa70d)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-e210fe40b954ecc3d4242335bd73ed6f5a6d086948c433a6cb9c2e376e6e384b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71f68c05dcbfe91820f80b2706f2a7c10a2a3532c56f67a53d93e2f9a526eadf"></a>

## slo_to_global_dr — slo_to_global_dr / dddc4b59f6f8 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- slo_to_global_dr

<a id="canonical-d788fe9acb9fcec352e96d84c696226160c897f7de12824fe84ae34ea9b32c10"></a>

Type: `"single"`. Computed.

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

<a id="canonical-aaabbbfea882b16aacfc2a2eaa87865957804e1136cb5df57f96a52608f4add7"></a>

## Direct properties — slo_to_global_dr / dddc4b59f6f8 / 3

- [global_vn](data-sources--network_connector--reference--group-001.md#canonical-cd02c22f385866a4105e462e93718c4c7180a2bba4fb2b2647df6f307afe49b2): complete subsection reference.

<a id="canonical-86b5acc1ff7a459d63bab81a51052909a43bc244e12c0b8c3d45385f979502ea"></a>

## Next pages — slo_to_global_dr / dddc4b59f6f8 / 4

- [slo_to_global_dr.global_vn](data-sources--network_connector--reference--group-001.md#canonical-cd02c22f385866a4105e462e93718c4c7180a2bba4fb2b2647df6f307afe49b2)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-cd02c22f385866a4105e462e93718c4c7180a2bba4fb2b2647df6f307afe49b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7c5d3a303178d6e3cf1874241cd29400be75edd5a3fc20e1d4cc27715565a48"></a>

## slo_to_global_dr.global_vn — slo_to_global_dr.global_vn / 7d7f2c3fb94a / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e210fe40b954ecc3d4242335bd73ed6f5a6d086948c433a6cb9c2e376e6e384b)
- slo_to_global_dr.global_vn

<a id="canonical-065d1f8269b39f4ccde8aa07e0aa57cbde8eed95500f041ac9b270fa1c06c7aa"></a>

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

<a id="canonical-565c3921b0590670a41f8d6ba2f138535d32d253843a420208816d0ce2b20d52"></a>

## Direct properties — slo_to_global_dr.global_vn / 7d7f2c3fb94a / 3

<a id="canonical-4942b7c23681c5149f5b8e714ae6f32d99427423068d445c58325430b2375540"></a>

<a id="canonical-cef41b4436afb5d66bb1c3fc703d18d20a59c5dda28c77a1249434c851b65709"></a>

## name property — slo_to_global_dr.global_vn / 7d7f2c3fb94a / 4

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

<a id="canonical-965b9df3817c1cf15d1bc6b8dd0baf69fb26f2293c3fd552affb579a31be9445"></a>

<a id="canonical-58e91590cfa9869811b9b30d2e2227b78814bb0b8b0fa0fecff80ac0b85c791b"></a>

## namespace property — slo_to_global_dr.global_vn / 7d7f2c3fb94a / 5

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

<a id="canonical-c8eaa3634a3c95d47a74fa177c4752e2d07aabce1955a84ba0cb8212eb643e65"></a>

<a id="canonical-b4e93144fc4d784dfecc3e2e8dc1538e68f45250e04ff92c25243fc60457bab4"></a>

## tenant property — slo_to_global_dr.global_vn / 7d7f2c3fb94a / 6

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

<a id="canonical-16923e6a6befe320ecad698d6f15407de3a5dd2e836855a6decdad30a5c2e412"></a>

## Next pages — slo_to_global_dr.global_vn / 7d7f2c3fb94a / 7

- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-e210fe40b954ecc3d4242335bd73ed6f5a6d086948c433a6cb9c2e376e6e384b)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
