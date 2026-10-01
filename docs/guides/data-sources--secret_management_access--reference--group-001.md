---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4b85abded67805d7d49ddb282e4961ac54a244634f77fd6bdb0603cd32f102f"></a>

## Property reference — Property reference / af98700c57c7 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- Property reference

<a id="canonical-111dd2c6644d9242a5e446162c3d8cf33acdaddf9b12d6d4ee249b651de61c7d"></a>

## Direct properties — Property reference / af98700c57c7 / 3

- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f): complete subsection reference.

<a id="canonical-0829fd498a5a6b5ac228a9c0b9c34d2273623e7355d42398ee02aadfe72295f7"></a>

<a id="canonical-50b17806e1906f63953521fa6abc84f25111aac1a52342f119aba5905a9ec0db"></a>

## annotations property — Property reference / af98700c57c7 / 4

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

<a id="canonical-da91059c77ce3e4d8ceeaf793c9072354d8652a50228df0fcd92485eca3e3acd"></a>

<a id="canonical-3d851f047f2e22d898215d3feb39d5d0566bb2069024c3c834784c3b714f3e18"></a>

## description property — Property reference / af98700c57c7 / 5

Type: `"string"`. Computed.

Description of the SecretManagementAccess.

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

<a id="canonical-f0f5fe17c97e99fdfa42bd25041a9f52aec4e12cfe8f2f3eef4df4b4ac445989"></a>

<a id="canonical-de253ea337dfeebc1b5315475a73a396546279b2880c6adf8b4d0aec9e6a0fb0"></a>

## id property — Property reference / af98700c57c7 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6babf5b8731715e21354e717cbdc354180b6ff02a53cb91ef652bbbdea00a216"></a>

<a id="canonical-ab46ec81c4d8297bb9f0900348e3b185c4a87f0fd96567f97286a279414280f2"></a>

## labels property — Property reference / af98700c57c7 / 7

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

<a id="canonical-291b53bb654de6060497310d53cd6041a5bae22ecb5b79a61f1250d2baeb7836"></a>

<a id="canonical-1a41e74433f714d5e9ffb99a2f5ed5b92000f9abcbdd601aac75821c291994ab"></a>

## name property — Property reference / af98700c57c7 / 8

Type: `"string"`. Required.

Name of the SecretManagementAccess.

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

<a id="canonical-65051b51ad26bd866d30db197793ef74d7fcf28fa6a6168aec5e10ebeed5536f"></a>

<a id="canonical-866559b04e49ca254e934426f00bd21842138a7cf43dc1546faf795e445a315e"></a>

## namespace property — Property reference / af98700c57c7 / 9

Type: `"string"`. Required.

Namespace where the SecretManagementAccess exists.

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

<a id="canonical-50fb0ce63239340b864ccd9cc539afe3d048629d11867ba05c020ddcd584de3a"></a>

<a id="canonical-72692f9065a950b1fe934b1070d1151a4fecac364f889c4747ad99ec4476a470"></a>

## provider_name property — Property reference / af98700c57c7 / 10

Type: `"string"`. Computed.

Name given to this secret management backend. site.provider needs to be unique, and will be
referenced for using this object.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091): complete subsection reference.

<a id="canonical-97cb197c08db5ee567e9aae5931428e0eb83ca1a3e556f8eb955436d06c4a05b"></a>

## All schema paths — Property reference / af98700c57c7 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `access_info` | [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-792a99e908515b3dd61eb8e18ad65bf281509f27839f8e7a881a096776738e46) |
| `access_info.rest_auth_info` | [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-78be185161f359b12a203de59dad45887635cdea5151018f9a1b6c0cdb6f9030) |
| `access_info.rest_auth_info.basic_auth` | [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-b80bb0d86dd21fe516f3e7df4ca21f743ee0ddb7a817f4b338522e3cec6da922) |
| `access_info.rest_auth_info.basic_auth.password` | [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-baed24498cc9be2c09b88997c15f0f8012e2b86879b7d4d530b2fdc1029642b8) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-9da2cd9643ebd29f38620cca0fd6b00513ae1d8fef684890d4d064f04b985009) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-001.md#canonical-d0f486d1e812c99460c9cbce0eafb47d460b550ab91d04817cc869fc5168bcf7) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-001.md#canonical-8ad4d0fe6e93e2c1a6e4a51a15d9961f1fde49e6ef0200aed6f68530612dcf3e) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-001.md#canonical-8b2a7690f003e86d04e92c02a6ee773de74df83e1f6e66dd89139d527380553a) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-4b3a3c8803ed0080cf620fdb0d2447dc486e6af1e4af4352a33b09b72521b71b) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-001.md#canonical-0145b00c435a57dc1ef0a1fe3a4466942cb0104582f320d8520a79f70ef51811) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.url` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.url](data-sources--secret_management_access--reference--group-001.md#canonical-8c6a38efdcc55c1f11d3210ea93e3fd7bad891e99e9d6265f4e443eef46729d8) |
| `access_info.rest_auth_info.basic_auth.username` | [access_info.rest_auth_info.basic_auth.username](data-sources--secret_management_access--reference--group-001.md#canonical-2fee73a1568ff52e7fcdcacefac217e220787dd0b3bc0ab5d9dfcc2755bbd3f2) |
| `access_info.rest_auth_info.headers_auth` | [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-855925fb8de47fbe1b55923addc296497eb5cd3c10185b4793e60862d3804b92) |
| `access_info.rest_auth_info.headers_auth.headers` | [access_info.rest_auth_info.headers_auth.headers](data-sources--secret_management_access--reference--group-001.md#canonical-8b5205987afde309b216d3063e11ba1ea21677de1d6aee276d0a99b390d9017b) |
| `access_info.rest_auth_info.query_params_auth` | [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-213eb74c5c3266ce4ca6b631d4b1136f6fadecc3f544071f9bc30627fdb4b442) |
| `access_info.rest_auth_info.query_params_auth.query_params` | [access_info.rest_auth_info.query_params_auth.query_params](data-sources--secret_management_access--reference--group-001.md#canonical-0501aa8ce5b454760846f2dcf9b862406f7d082d669b8cf90b622b691cba9ab4) |
| `access_info.scheme` | [access_info.scheme](data-sources--secret_management_access--reference--group-001.md#canonical-cec8590a59d6fa136ce90a94589d8f517a464b6fa1c33f433e7191250606e924) |
| `access_info.server_endpoint` | [access_info.server_endpoint](data-sources--secret_management_access--reference--group-001.md#canonical-0221be16674d79f758ee64e24ad5c9986ef4e339b5d8b3eb419de732280f18f7) |
| `access_info.tls_config` | [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-a384e74f4940aa90820c9a746440fbf5623ea3a9f771ab76d11b27f76872eb4f) |
| `access_info.tls_config.cert_params` | [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-a76b118600c10db88a94cd701b88c98bd835661b526e14fcfc8e7e766f39f5ea) |
| `access_info.tls_config.cert_params.certificates` | [access_info.tls_config.cert_params.certificates](data-sources--secret_management_access--reference--group-001.md#canonical-9498b8985f31619bb1a9f5aa76608327e79469aaec72f529162abc76c1cb0e49) |
| `access_info.tls_config.cert_params.certificates.kind` | [access_info.tls_config.cert_params.certificates.kind](data-sources--secret_management_access--reference--group-001.md#canonical-17745fb37774a369615629b15fc5b44679cb94f90cdaa1b4b7831cbcfcdae684) |
| `access_info.tls_config.cert_params.certificates.name` | [access_info.tls_config.cert_params.certificates.name](data-sources--secret_management_access--reference--group-001.md#canonical-4a3627d87ac870cd8304613ed6d058153c33eb6624baa957b254f930041ffaa4) |
| `access_info.tls_config.cert_params.certificates.namespace` | [access_info.tls_config.cert_params.certificates.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-b272fdca9f1efc345d1ed082a257e209089c4bf743510c8e0006bec3922cf519) |
| `access_info.tls_config.cert_params.certificates.tenant` | [access_info.tls_config.cert_params.certificates.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-8b7c012195ccece82f0b446419e403207ec5754c6aadfe12090b36287fda16f8) |
| `access_info.tls_config.cert_params.certificates.uid` | [access_info.tls_config.cert_params.certificates.uid](data-sources--secret_management_access--reference--group-001.md#canonical-f1cdc04f07f7b0003f7de81fdf6bc9a282b0d590e547804c1e0933d06b61570a) |
| `access_info.tls_config.cert_params.cipher_suites` | [access_info.tls_config.cert_params.cipher_suites](data-sources--secret_management_access--reference--group-001.md#canonical-48316376c0d238fd9bba607edb1bdb7a39c31b52ee8338d35c20eeaa2ed4ef05) |
| `access_info.tls_config.cert_params.maximum_protocol_version` | [access_info.tls_config.cert_params.maximum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-3041797e0abd524b7885571277426cdf169db298ad9c7c0b29f4e91e214a46e4) |
| `access_info.tls_config.cert_params.minimum_protocol_version` | [access_info.tls_config.cert_params.minimum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-381ffd9c422952b6237d0418e54c2922702f333bd68558dce5173c31b5b310a2) |
| `access_info.tls_config.cert_params.skip_server_verification` | [access_info.tls_config.cert_params.skip_server_verification](data-sources--secret_management_access--reference--group-001.md#canonical-9ec2daaf48a5932529fe3fa31f74df0009a1ac02a5226ff2a18f75f895218596) |
| `access_info.tls_config.cert_params.tls_validation_params` | [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-77787f714ad64acc0032ebedb51a052dd633b1f2a22edef949bb59e5cdceea0e) |
| `access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification` | [access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification](data-sources--secret_management_access--reference--group-001.md#canonical-4ed24bd03c7a386a6549c8fdf246ca46cf88277d2f9e5111a00a90539ddb5b0e) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-c88a9eb2447c9551d39d730d1e00fd9ad3e0af00b9f805a6fa0d1feca7d5670f) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-05fd02548e9bfa4b75068217ef6e7c22a618efb927410b2c99a4b63b1c3583d4) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](data-sources--secret_management_access--reference--group-001.md#canonical-363b1024d329ce8c5d2165c0a12cb35176a10b5e81c68a261b9efd718a24a4e0) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](data-sources--secret_management_access--reference--group-001.md#canonical-74a117621005cb39cff3d59b1412edc295729c6eab726a9e73d78254c14212fd) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-9f5caac418d49731cb010b40b21f6fe3c52f4ad327f7ef5a01c1b4c85c8b4b0c) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-fb4128b9be3488a14fb25fde22ad5126003dfeec7de3e4821f285cf6dc791d16) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](data-sources--secret_management_access--reference--group-001.md#canonical-73ecfbce1cd6902ffd830b69f6cef5a2cef5947618e30270b8265ea8ae41cfe7) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url](data-sources--secret_management_access--reference--group-001.md#canonical-97ab98186155e71ca397829383973e05c3ccbf78a63aba7b9080438c36673b6d) |
| `access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names` | [access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names](data-sources--secret_management_access--reference--group-001.md#canonical-02cbcf2a5cfb0827ef43525574ce614bfa34078d6731341a2c4fdf0afc2141c0) |
| `access_info.tls_config.cert_params.volterra_trusted_ca` | [access_info.tls_config.cert_params.volterra_trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-1b5b4ef1e29651049c95c93bb0561b9d47b0a124dce3e406747dae81bec2b5ef) |
| `access_info.tls_config.common_params` | [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-538da24ac2a26120dd8d06091dab2b567ee4327fb070fa0f804428f5fb4860ba) |
| `access_info.tls_config.common_params.cipher_suites` | [access_info.tls_config.common_params.cipher_suites](data-sources--secret_management_access--reference--group-001.md#canonical-ece869c94b0cfe67241956345a7230a9da03a2555d656fd1f6d0e1e99ccac656) |
| `access_info.tls_config.common_params.maximum_protocol_version` | [access_info.tls_config.common_params.maximum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-3bc513df940e820e3b92afe8b182048590c2e83cddf22ee9044b5691c057bbe3) |
| `access_info.tls_config.common_params.minimum_protocol_version` | [access_info.tls_config.common_params.minimum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-4eecd6531c4f1d554020043b157476617a5d3ef2043de5e21eacb7018182f42d) |
| `access_info.tls_config.common_params.tls_certificates` | [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-fbc82c63b3485aa65a550b1fa2e66a0574c69de800f6da77105863916a0484ed) |
| `access_info.tls_config.common_params.tls_certificates.certificate_url` | [access_info.tls_config.common_params.tls_certificates.certificate_url](data-sources--secret_management_access--reference--group-001.md#canonical-87712c4827117039acc251919492bdf38b037eb635f5fe3c0bb2aefc48a09519) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-07eb2c6eafcd6be232326a2a28178e6c335271ceeee75d3522b3bfef90adf203) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-b028bfec2c4016ec6c1f8900c21822087eb70423507455c1cf25f7b44b0de466) |
| `access_info.tls_config.common_params.tls_certificates.description_spec` | [access_info.tls_config.common_params.tls_certificates.description_spec](data-sources--secret_management_access--reference--group-001.md#canonical-8a797c7655b8bbb369f9321b84fa633e21f58547e2c2370629c60a90c42fad97) |
| `access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling` | [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](data-sources--secret_management_access--reference--group-001.md#canonical-b12e840438a1ae54988c80f53ffc5b844edce0b03f9d69b6a1ec8e8cc70d1471) |
| `access_info.tls_config.common_params.tls_certificates.private_key` | [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-09cb66af96a92c407dad4fc15e38712b7a47644ec45e5e6929df38b57685e69a) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-a611cf9a734fa16acc78b77663026840ccc2372b09c1da947f2ef195a65a6173) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-001.md#canonical-5750f4b07d4c036e13f9181e89767a94b3e750b820c362c1104fbfa6b6976393) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-001.md#canonical-6c2bb81e4a1c54292777e0ea6a915b6deeed89a1992f0d2fd619b9f80a93c103) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-001.md#canonical-219462b49b0f4b419f8f4c59b20862623c7c6fe47dd69c0d1ddf1a66d4071510) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-d2a1b45839fbce48437574c5400bc4cfd62dedce5fcb6590799ae67a5ba5fc89) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-001.md#canonical-18e224f43ba39ad3e19d2c16ee6fa7ab85d6895711033712cef942ae32958eda) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--secret_management_access--reference--group-001.md#canonical-f9999f908cdbf60062a36ac3538fd074afcbfa55983e5b26f93af1ad860d9615) |
| `access_info.tls_config.common_params.tls_certificates.use_system_defaults` | [access_info.tls_config.common_params.tls_certificates.use_system_defaults](data-sources--secret_management_access--reference--group-001.md#canonical-2930506f6874552996c091af31295595e4c688af6715ab632e055398d95ecf7c) |
| `access_info.tls_config.common_params.validation_params` | [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-b48873f832bc572a35909f5a4d30fa6452451f5a5e7df20f9f1499105c7bcb69) |
| `access_info.tls_config.common_params.validation_params.skip_hostname_verification` | [access_info.tls_config.common_params.validation_params.skip_hostname_verification](data-sources--secret_management_access--reference--group-001.md#canonical-0a5ed0be7e2b10568f4055469bc7678947c8e2dc7845f393b741caa2e342e573) |
| `access_info.tls_config.common_params.validation_params.trusted_ca` | [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-162533777383e6c8de6fdbb34dd188a8f713642aa6fade49111a5035f49d2a93) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-53b7e4657ff17c184c198379ddfbc48fb3280d4c0acd564bfa8db94ca07beb8c) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--secret_management_access--reference--group-001.md#canonical-7fdf12e96b7d95220da06bcf3f9ccc200e09e5e610d43736c068c982eadede24) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--secret_management_access--reference--group-001.md#canonical-10d730b65dcd7065ff3b43f55d9d49d094abc99e8241f928832752087fe23111) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-3ade70fb627f0f7e985c1670e76ad92dcf8bb55eef0d0d393cc5627d36eb45fe) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-f5dd89b185c6b2be66892777be41322646e31c265d024ca6fa28d09a46ea2aa6) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--secret_management_access--reference--group-001.md#canonical-3dffa2db891741d3f8c89716dd7e0a66054a19067630abaeba01fbbb0a6d2664) |
| `access_info.tls_config.common_params.validation_params.trusted_ca_url` | [access_info.tls_config.common_params.validation_params.trusted_ca_url](data-sources--secret_management_access--reference--group-001.md#canonical-9f8690272dd70f28161d06c6723ba4177898e05cdf0a8957492a334802b541b4) |
| `access_info.tls_config.common_params.validation_params.verify_subject_alt_names` | [access_info.tls_config.common_params.validation_params.verify_subject_alt_names](data-sources--secret_management_access--reference--group-001.md#canonical-4eb0fb54e1151c2bcc16752768b5145703e116becd35efa8f3c3646a6e66dc83) |
| `access_info.tls_config.default_session_key_caching` | [access_info.tls_config.default_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-9de4e21b2eb640b48d5ca324d9d20bfaff9257dd00d3ddf4bcc4541efb9779ad) |
| `access_info.tls_config.disable_session_key_caching` | [access_info.tls_config.disable_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-8e1082066faf0e83671caea9a291113fd59fef162836b7429635a23f66c740ab) |
| `access_info.tls_config.disable_sni` | [access_info.tls_config.disable_sni](data-sources--secret_management_access--reference--group-001.md#canonical-00f159ddcdac4a649cfb08cf61ed9041185e85f5b8934dccf4d797b11d387bfd) |
| `access_info.tls_config.max_session_keys` | [access_info.tls_config.max_session_keys](data-sources--secret_management_access--reference--group-001.md#canonical-8e1ff9ecf0e2bfe88fbd3891e8d7d513d978c6aea4166d72490e32297b55cb90) |
| `access_info.tls_config.sni` | [access_info.tls_config.sni](data-sources--secret_management_access--reference--group-001.md#canonical-55d987e67ea437faf23c955c25aa66bd9aa4712bdb50068d7a74c8a4d3a29306) |
| `access_info.tls_config.use_host_header_as_sni` | [access_info.tls_config.use_host_header_as_sni](data-sources--secret_management_access--reference--group-001.md#canonical-27217a92c65cccbf874e8bda0dba5c36165d613e3cfe3e9702c02482533f6e84) |
| `access_info.vault_auth_info` | [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-d91d6f254ddc58a55327dca4012ea4ef6e51fffdbcf5724d1ee559ffec403c5e) |
| `access_info.vault_auth_info.app_role_auth` | [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-83c35b0ed1f0a55831e37180fe71177b8768650b899b374fd15b07633ccd80c8) |
| `access_info.vault_auth_info.app_role_auth.role_id` | [access_info.vault_auth_info.app_role_auth.role_id](data-sources--secret_management_access--reference--group-001.md#canonical-c3fda162312388d2733957f63ad6c9ecbb67bd3d4cc134ad66e7627f0a3392e6) |
| `access_info.vault_auth_info.app_role_auth.secret_id` | [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-b17e59d4c3adb33a6b1945d8e48857b67153f16d0f90fed765d20635946ef3a4) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-5b371b4ec88addc4bedec4b999e85052b1c37ff0a8543f4a5f171cbdb6900641) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-002.md#canonical-fea58533dda501a623605e7043e3ac36a35d2217c9290b5d0de944819bb0900a) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-002.md#canonical-861a3856a789e875fff848887cd21c2841f9b044ba7d3df728cabb8b7c70995b) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-002.md#canonical-666974a7b681c8972cf33e0f4ab7ae9b2553f733795df9d6575501a399049edc) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-d0bd1ce0c97981266a5b8df6395ce6ed04206864709699e0c904c7a8e4c6a969) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-002.md#canonical-1553b328697d64b2c8d9502135680f243c592fbed829196333df63f80365ae1a) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url](data-sources--secret_management_access--reference--group-002.md#canonical-316b62ab3833f139e1f5b63e152db5d01c8fb1ee5e37fc952de9b0a4d7e5803f) |
| `access_info.vault_auth_info.token` | [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-0a6b89a8e23e911e2ad9195db9e354d6a24496cdac56ac756acd9ffcb0909bf4) |
| `access_info.vault_auth_info.token.blindfold_secret_info` | [access_info.vault_auth_info.token.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-80458f73ec1a42d7b25b10b0763f1a938f3a5dcc52cd20b52f26e0b2de33123a) |
| `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-002.md#canonical-9b4e236ecbb27fdc953fdd6b0934784f06dc1c67339d5c052f1be0f83e266cd1) |
| `access_info.vault_auth_info.token.blindfold_secret_info.location` | [access_info.vault_auth_info.token.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-002.md#canonical-bc8ff4876123da87d37be029aac6db68502ca3fd4d8c84adf8ddd33f1693c77d) |
| `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-002.md#canonical-548626d85b7161805bb7836cae7d0c6ae5cc5fc1a456c28b61cfa5e6e74f9c0c) |
| `access_info.vault_auth_info.token.clear_secret_info` | [access_info.vault_auth_info.token.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-7a42ccbba549450011eb623f709762e0f9c4b7cc7c8f8814013859d4e720f6f7) |
| `access_info.vault_auth_info.token.clear_secret_info.provider_ref` | [access_info.vault_auth_info.token.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-002.md#canonical-04cee01e2f7b809fe837761f754ecf4a84711bb30335faaf30dff7b2485d0f55) |
| `access_info.vault_auth_info.token.clear_secret_info.url` | [access_info.vault_auth_info.token.clear_secret_info.url](data-sources--secret_management_access--reference--group-002.md#canonical-9b8e627a9c6d09435fd341172a15968797e430b2f4cf639c845071b34432ae47) |
| `annotations` | [annotations](data-sources--secret_management_access--reference--group-001.md#canonical-0829fd498a5a6b5ac228a9c0b9c34d2273623e7355d42398ee02aadfe72295f7) |
| `description` | [description](data-sources--secret_management_access--reference--group-001.md#canonical-da91059c77ce3e4d8ceeaf793c9072354d8652a50228df0fcd92485eca3e3acd) |
| `id` | [id](data-sources--secret_management_access--reference--group-001.md#canonical-f0f5fe17c97e99fdfa42bd25041a9f52aec4e12cfe8f2f3eef4df4b4ac445989) |
| `labels` | [labels](data-sources--secret_management_access--reference--group-001.md#canonical-6babf5b8731715e21354e717cbdc354180b6ff02a53cb91ef652bbbdea00a216) |
| `name` | [name](data-sources--secret_management_access--reference--group-001.md#canonical-291b53bb654de6060497310d53cd6041a5bae22ecb5b79a61f1250d2baeb7836) |
| `namespace` | [namespace](data-sources--secret_management_access--reference--group-001.md#canonical-65051b51ad26bd866d30db197793ef74d7fcf28fa6a6168aec5e10ebeed5536f) |
| `provider_name` | [provider_name](data-sources--secret_management_access--reference--group-001.md#canonical-50fb0ce63239340b864ccd9cc539afe3d048629d11867ba05c020ddcd584de3a) |
| `where` | [where](data-sources--secret_management_access--reference--group-002.md#canonical-c20b6c23fffc096f6c342afeafd03fed39b0fb3c9e05d3ce64b3a0aa1ba51533) |
| `where.site` | [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-fec7f24ae73cbd19402e1fafddc17e6a93e4edb40cc4602def484926589d6a3c) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-86644d4b38ebba5fad9fd82b33cfe96730283206979dcd801d02526128537cf2) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-5569b54acfa7ec67102cda2969b02a0932426d0f816cf080bd1175aa98a69d1b) |
| `where.site.network_type` | [where.site.network_type](data-sources--secret_management_access--reference--group-002.md#canonical-a425bc40fb365dfe718b426af5a24e386af61b20d6407e6c09a57455987da4ca) |
| `where.site.ref` | [where.site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-b82dc783eac399bf1520012a8fec6d35a11fd913cf68b6564c61a93c0d655789) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--secret_management_access--reference--group-002.md#canonical-eb7a5273efaf3a19514a3744c04e849edf55786a3fc7822c50971d162d172012) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--secret_management_access--reference--group-002.md#canonical-0c405da0ce1c2bd7f70f719b40d16dd0eb74b918a2d566cb4cf0daaff59a20ab) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--secret_management_access--reference--group-002.md#canonical-1aafb86024dc8cd81578fc8e305ca917cf3fb65a0f5b4379a73d1cfc75c72c4b) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--secret_management_access--reference--group-002.md#canonical-5fd7b960d74b6ccea4a83681dcbfd8c8930322e38119b7fea21fd8b3e4c79f32) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--secret_management_access--reference--group-002.md#canonical-e581269cc591de2fcd7bcdd9cea5c3b054a66744d47d0deb7b49a876feb48626) |
| `where.virtual_network` | [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-3c5ad0e9ebb1682491b37c0c03b1d9848e619dc246ef8ac374037ea1444b0bbc) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--secret_management_access--reference--group-002.md#canonical-5cbe56e11873703a1592ad788a7ef4f1e674affedddbbdd28278bb0fd0921220) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--secret_management_access--reference--group-002.md#canonical-67006bc5f99913e934e1b8b012b4e039af3023c2c0852575db921c191cbfd6bc) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--secret_management_access--reference--group-002.md#canonical-b47bdbb8fff05ec0db9cefb4d9e69888180c6c29a27d284c863bb03d38bf9a87) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--secret_management_access--reference--group-002.md#canonical-4c74e9e157da85f7e8e54c71810ebe0487fc87d79367803713114ec3cfb31f40) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--secret_management_access--reference--group-002.md#canonical-04b30ab73411a55a433db3a1a5cced269299e2b92fe5899da71a3251d624ef55) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--secret_management_access--reference--group-002.md#canonical-7131ea14e2863efa9eafc94b8c20265ac7a27b5caa0dda3a8a4912f4630529ce) |
| `where.virtual_site` | [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-813f538ac0ebb50de9172bcb70378f2116784ef4a9c64126b11cf6338ee68562) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-5edeaa0cf231d4cb00eb70ae2960303a85dda2b63fcc54bc1e6710822be66463) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-580b2edb06d778c720397c358cb83079f610990b21088ec217cdc5bb31a5fc81) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--secret_management_access--reference--group-002.md#canonical-7f38b6b5d4098b00202675c4d9337933820a1057b112e1e591ea98d47bcf3809) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-6e340a1d3890b7a12f9485fbce02d6f9089da239a90a5e4d0d9e3a07c02a1b67) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--secret_management_access--reference--group-002.md#canonical-a2b4dc75d694b397b23ace282e6ee7ccf772aa10f7c592d097696ba925311f38) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--secret_management_access--reference--group-002.md#canonical-eeb8c1fb2365d8b6ba9b3836dd10340ce72de04043c66b6a65b5af0886197bfc) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--secret_management_access--reference--group-002.md#canonical-49b07b3ea46cc8b82f96007af15e1ea11cb616c16561e4aaad0f154db7d1d008) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--secret_management_access--reference--group-002.md#canonical-2858000a57bb431f3f1aae3ac8f76d441b4f8f4e2bb7b6f0fab8dd624cf01053) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--secret_management_access--reference--group-002.md#canonical-f52875049ff87f0c88b7d714afba0f1bddcdf7d50f1d7fae32e490074bf3c1d3) |

<a id="canonical-ca04bfa05efea198d9c7ec8513acd844242a9370a82383090dd60063038692a9"></a>

## Next pages — Property reference / af98700c57c7 / 12

- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6efc55f735b4e20b8e9ad02cf5853617984fe12ae7c03db7f3098720982fff90"></a>

## access_info — access_info / 80c9d9567a6b / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- access_info

<a id="canonical-792a99e908515b3dd61eb8e18ad65bf281509f27839f8e7a881a096776738e46"></a>

Type: `"single"`. Computed.

HostAccessInfoType contains the information about how to connect to the remote host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"rest_auth_info\",\"vault_auth_info\"]"
}
```

<a id="canonical-91a987226a6671ee5b4361207c26a9ccbacf84f308ec0d9aa8e5fddef8af348b"></a>

## Direct properties — access_info / 80c9d9567a6b / 3

- [rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d): complete subsection reference.

<a id="canonical-cec8590a59d6fa136ce90a94589d8f517a464b6fa1c33f433e7191250606e924"></a>

<a id="canonical-56c85286204578a4ef651067052e2a3255cae778cdd14cd6847de48da83c4615"></a>

## scheme property — access_info / 80c9d9567a6b / 4

Type: `"string"`. Computed.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Upstream description:

SchemeType is used to indicate URL scheme

HTTP:// scheme HTTPS:// scheme.

Receipt-pinned upstream constraints:

```json
{
  "default": "HTTP",
  "enum": [
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0221be16674d79f758ee64e24ad5c9986ef4e339b5d8b3eb419de732280f18f7"></a>

<a id="canonical-7050c5c505c661abe513b494f0214971573f6b06f415c53ab7ecc830893624a3"></a>

## server_endpoint property — access_info / 80c9d9567a6b / 5

Type: `"string"`. Computed.

Endpoint to connect to, in host:port format.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b): complete subsection reference.

- [vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a): complete subsection reference.

<a id="canonical-66fa109b579716bc530483c632fbbe89bc2e475825e129754eb447f096ec267d"></a>

## Next pages — access_info / 80c9d9567a6b / 6

- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f64b31aa2db3f89b0e1b2eb758e1129c2959af5ec763cb20a351355f8a8266a5"></a>

## access_info.rest_auth_info — access_info.rest_auth_info / e4d0e4cab22c / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- access_info.rest_auth_info

<a id="canonical-78be185161f359b12a203de59dad45887635cdea5151018f9a1b6c0cdb6f9030"></a>

Type: `"single"`. Computed.

Authentication parameters for REST based hosts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"basic_auth\",\"headers_auth\",\"query_params_auth\"]"
}
```

<a id="canonical-9bc3d6b00572b096b75cee6bfbaf81541ab9745348a3c09d7b05fd9b6854226b"></a>

## Direct properties — access_info.rest_auth_info / e4d0e4cab22c / 3

- [basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-be92e225aaa97d29118f36639f6984439a8988faf0de911ca85ee3bac3becc31): complete subsection reference.

- [headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-a815f9ddd1755489b832d7eb01e2ef94ec7f42a7cf6b708d4ac86c849b485a4c): complete subsection reference.

- [query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-db9483694fd3b0c1ffe08f84b55b48629af8dbf569a3f814638a67505b936044): complete subsection reference.

<a id="canonical-cab01ea04c9bcd02d93cf6901176e189516cc835c21a7414eb139976c46e4e15"></a>

## Next pages — access_info.rest_auth_info / e4d0e4cab22c / 4

- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-be92e225aaa97d29118f36639f6984439a8988faf0de911ca85ee3bac3becc31)
- [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-a815f9ddd1755489b832d7eb01e2ef94ec7f42a7cf6b708d4ac86c849b485a4c)
- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-db9483694fd3b0c1ffe08f84b55b48629af8dbf569a3f814638a67505b936044)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-be92e225aaa97d29118f36639f6984439a8988faf0de911ca85ee3bac3becc31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59a48fba7a8aee3daf7c818b4be0aa467ee89764954c8e7018c08957a6e4a262"></a>

## access_info.rest_auth_info.basic_auth — access_info.rest_auth_info.basic_auth / d8e707395975 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- access_info.rest_auth_info.basic_auth

<a id="canonical-b80bb0d86dd21fe516f3e7df4ca21f743ee0ddb7a817f4b338522e3cec6da922"></a>

Type: `"single"`. Computed.

AuthnTypeBasicAuth is used for using basic\_auth mode of HTTP authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0cbc063e10ed2fdbebbcf291c71bdafb4098f11f0b670cf095cfbf14f3ae9d85"></a>

## Direct properties — access_info.rest_auth_info.basic_auth / d8e707395975 / 3

- [password](data-sources--secret_management_access--reference--group-001.md#canonical-2e3cfbbef05dfbc8a0992e062649382290216a0b5b3cd5d6ea526bd3117927a6): complete subsection reference.

<a id="canonical-2fee73a1568ff52e7fcdcacefac217e220787dd0b3bc0ab5d9dfcc2755bbd3f2"></a>

<a id="canonical-5bf61a5d48a4e244a770351a06f59d9b4f24f12c7e08b55b9e04307c3da4760a"></a>

## username property — access_info.rest_auth_info.basic_auth / d8e707395975 / 4

Type: `"string"`. Computed.

The username to encode in Basic Auth scheme.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
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

<a id="canonical-d1bf259361d57a102e151140e6794e98f15ae91d48682c31a2a774243c6feca5"></a>

## Next pages — access_info.rest_auth_info.basic_auth / d8e707395975 / 5

- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-2e3cfbbef05dfbc8a0992e062649382290216a0b5b3cd5d6ea526bd3117927a6)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-2e3cfbbef05dfbc8a0992e062649382290216a0b5b3cd5d6ea526bd3117927a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efebd8ac582bcba6a02897969706576d6e031704752019c86f5e828d7ef1169e"></a>

## access_info.rest_auth_info.basic_auth.password — access_info.rest_auth_info.basic_auth.password / c1c5746fac59 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-be92e225aaa97d29118f36639f6984439a8988faf0de911ca85ee3bac3becc31)
- access_info.rest_auth_info.basic_auth.password

<a id="canonical-baed24498cc9be2c09b88997c15f0f8012e2b86879b7d4d530b2fdc1029642b8"></a>

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

<a id="canonical-d4f25551dd8a347dc1ada668a7d4b17419b1371b4891902b9da72be039ea7ded"></a>

## Direct properties — access_info.rest_auth_info.basic_auth.password / c1c5746fac59 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-fc0de5b1e1c94a23b9251268d4783e5933d3b73f2547999706490acdaeb4f9fc): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-c4fc25e2f2efdcb54d65b67cb875cacae4dae8e394c3048b042f71daffcc8c43): complete subsection reference.

<a id="canonical-f02c506512056bcd47df7e3972386b31f45b9b706a572447fcdfc69ef578edc1"></a>

## Next pages — access_info.rest_auth_info.basic_auth.password / c1c5746fac59 / 4

- [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-fc0de5b1e1c94a23b9251268d4783e5933d3b73f2547999706490acdaeb4f9fc)
- [access_info.rest_auth_info.basic_auth.password.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-c4fc25e2f2efdcb54d65b67cb875cacae4dae8e394c3048b042f71daffcc8c43)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-be92e225aaa97d29118f36639f6984439a8988faf0de911ca85ee3bac3becc31)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-fc0de5b1e1c94a23b9251268d4783e5933d3b73f2547999706490acdaeb4f9fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bc29fbde79ff8ed70a535b80000bb22e60554bd6308b25a3fe0df7dd605f41b"></a>

## access_info.rest_auth_info.basic_auth.password.blindfold_secret_info — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / a0c85ce1ee0c / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-be92e225aaa97d29118f36639f6984439a8988faf0de911ca85ee3bac3becc31)
- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-2e3cfbbef05dfbc8a0992e062649382290216a0b5b3cd5d6ea526bd3117927a6)
- access_info.rest_auth_info.basic_auth.password.blindfold_secret_info

<a id="canonical-9da2cd9643ebd29f38620cca0fd6b00513ae1d8fef684890d4d064f04b985009"></a>

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

<a id="canonical-25ae105386602dc263d82611a06f4657ff245a916e9f6360272fc5a8fd0833eb"></a>

## Direct properties — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / a0c85ce1ee0c / 3

<a id="canonical-d0f486d1e812c99460c9cbce0eafb47d460b550ab91d04817cc869fc5168bcf7"></a>

<a id="canonical-da5af5a333af064b94f6be704187e7ec8e2a2f6d3c3712c649ce46cbf8822203"></a>

## decryption_provider property — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / a0c85ce1ee0c / 4

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

<a id="canonical-8ad4d0fe6e93e2c1a6e4a51a15d9961f1fde49e6ef0200aed6f68530612dcf3e"></a>

<a id="canonical-e05f3256410f196089d5639b88583b801cae156bfa7f7dae2373f3dc187bd656"></a>

## location property — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / a0c85ce1ee0c / 5

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

<a id="canonical-8b2a7690f003e86d04e92c02a6ee773de74df83e1f6e66dd89139d527380553a"></a>

<a id="canonical-35a07ba1b4788316dc8d3623e62784b4bbfe5ed08ea8aaf77dda537094231f59"></a>

## store_provider property — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / a0c85ce1ee0c / 6

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

<a id="canonical-335ba15d74a3a1b1c0ddb4b82e053e17656ee7b9cada9e31ec55a63973e1dbd1"></a>

## Next pages — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / a0c85ce1ee0c / 7

- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-2e3cfbbef05dfbc8a0992e062649382290216a0b5b3cd5d6ea526bd3117927a6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-c4fc25e2f2efdcb54d65b67cb875cacae4dae8e394c3048b042f71daffcc8c43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70ccbe539f6594579f7b29278c96dfe1062a33568bb5ae6b001aec3682be043a"></a>

## access_info.rest_auth_info.basic_auth.password.clear_secret_info — access_info.rest_auth_info.basic_auth.password.clear_secret_info / d7ce3f65e0fe / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-be92e225aaa97d29118f36639f6984439a8988faf0de911ca85ee3bac3becc31)
- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-2e3cfbbef05dfbc8a0992e062649382290216a0b5b3cd5d6ea526bd3117927a6)
- access_info.rest_auth_info.basic_auth.password.clear_secret_info

<a id="canonical-4b3a3c8803ed0080cf620fdb0d2447dc486e6af1e4af4352a33b09b72521b71b"></a>

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

<a id="canonical-d240744ae4645f83d663653b92e335f16641c2933d9b9d11c51b2bf171d0ef39"></a>

## Direct properties — access_info.rest_auth_info.basic_auth.password.clear_secret_info / d7ce3f65e0fe / 3

<a id="canonical-0145b00c435a57dc1ef0a1fe3a4466942cb0104582f320d8520a79f70ef51811"></a>

<a id="canonical-a6171f4e86cbc0ebed1f01dde8e506cbba69b9fbe9647acd7c8ca88e58ebdc6f"></a>

## provider_ref property — access_info.rest_auth_info.basic_auth.password.clear_secret_info / d7ce3f65e0fe / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8c6a38efdcc55c1f11d3210ea93e3fd7bad891e99e9d6265f4e443eef46729d8"></a>

<a id="canonical-56af9563d76671060489d4981abbeefa8bbcdd0d93e2eadb7da4506687699fee"></a>

## url property — access_info.rest_auth_info.basic_auth.password.clear_secret_info / d7ce3f65e0fe / 5

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

<a id="canonical-2a7a4a381d2b208e3c6eb9e8ce0069c6937af8eed803dd10f07220ee89301a37"></a>

## Next pages — access_info.rest_auth_info.basic_auth.password.clear_secret_info / d7ce3f65e0fe / 6

- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-2e3cfbbef05dfbc8a0992e062649382290216a0b5b3cd5d6ea526bd3117927a6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-a815f9ddd1755489b832d7eb01e2ef94ec7f42a7cf6b708d4ac86c849b485a4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89f4f34537ecb419d5f6b90e6ef9923a4b032d5fca4b659a4f3d9e3717bde2a9"></a>

## access_info.rest_auth_info.headers_auth — access_info.rest_auth_info.headers_auth / 72e84b3b27d9 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- access_info.rest_auth_info.headers_auth

<a id="canonical-855925fb8de47fbe1b55923addc296497eb5cd3c10185b4793e60862d3804b92"></a>

Type: `"single"`. Computed.

AuthnTypeHeaders is used for setting headers for authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a61be88f892e252962d48b8d0e5befa392208f4e503954212ddc699491d92b14"></a>

## Direct properties — access_info.rest_auth_info.headers_auth / 72e84b3b27d9 / 3

- [headers](data-sources--secret_management_access--reference--group-001.md#canonical-430fdda0fb5d5d5727a8241d59b2703638283d8c85e56b5c0adb2253ec6f5d5a): complete subsection reference.

<a id="canonical-c4ac1fa258383eab6b00cb06e6ffc9271c4f62bbc13e5584a22a591389b027e0"></a>

## Next pages — access_info.rest_auth_info.headers_auth / 72e84b3b27d9 / 4

- [access_info.rest_auth_info.headers_auth.headers](data-sources--secret_management_access--reference--group-001.md#canonical-430fdda0fb5d5d5727a8241d59b2703638283d8c85e56b5c0adb2253ec6f5d5a)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-430fdda0fb5d5d5727a8241d59b2703638283d8c85e56b5c0adb2253ec6f5d5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5ac80a8409bf5ebe928bbb47a4398ee6289aae32840cd8926a132b1eda967ff"></a>

## access_info.rest_auth_info.headers_auth.headers — access_info.rest_auth_info.headers_auth.headers / 01b92b250a54 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-a815f9ddd1755489b832d7eb01e2ef94ec7f42a7cf6b708d4ac86c849b485a4c)
- access_info.rest_auth_info.headers_auth.headers

<a id="canonical-8b5205987afde309b216d3063e11ba1ea21677de1d6aee276d0a99b390d9017b"></a>

Type: `"single"`. Computed.

The set of authentication headers to pass in HTTP request.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-0b43db9a98c7d160ec794240b39593642738dcc8fedcb18b5059115b84abcd59"></a>

## Direct properties — access_info.rest_auth_info.headers_auth.headers / 01b92b250a54 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-259c62af91cde4cfcbb3dfe5ee4c47b91c2390a6bf3cf927e750eb84bc307b1f"></a>

## Next pages — access_info.rest_auth_info.headers_auth.headers / 01b92b250a54 / 4

- [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-a815f9ddd1755489b832d7eb01e2ef94ec7f42a7cf6b708d4ac86c849b485a4c)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-db9483694fd3b0c1ffe08f84b55b48629af8dbf569a3f814638a67505b936044"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7102b09c66bba4923e682f25783f8a97d63fae41453f287c1077ff79d5645632"></a>

## access_info.rest_auth_info.query_params_auth — access_info.rest_auth_info.query_params_auth / 7a0063a28242 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- access_info.rest_auth_info.query_params_auth

<a id="canonical-213eb74c5c3266ce4ca6b631d4b1136f6fadecc3f544071f9bc30627fdb4b442"></a>

Type: `"single"`. Computed.

AuthnTypeQueryParams is used for setting query\_params for authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3e292f652f95f526c42edf47432048ed34f6986dfcfe5fe766eb179d586e8f05"></a>

## Direct properties — access_info.rest_auth_info.query_params_auth / 7a0063a28242 / 3

- [query_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd2da03ba28284497e8dd4f7ac499afd04e9173d1a0bd8b99721ec3abe5996ce): complete subsection reference.

<a id="canonical-8530e03ffe13d90f7cc88c4af35184281d6dc2e5d81d06a958e245015e6a62a3"></a>

## Next pages — access_info.rest_auth_info.query_params_auth / 7a0063a28242 / 4

- [access_info.rest_auth_info.query_params_auth.query_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd2da03ba28284497e8dd4f7ac499afd04e9173d1a0bd8b99721ec3abe5996ce)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-dd2da03ba28284497e8dd4f7ac499afd04e9173d1a0bd8b99721ec3abe5996ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14ed5b834b251dd0044a7a833a4041f8d0cedd57be171725c7673ddc950ca3e4"></a>

## access_info.rest_auth_info.query_params_auth.query_params — access_info.rest_auth_info.query_params_auth.query_params / 3769ab8bedd4 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-48daa251b0a6bc036af041f2426fb9369f9ba6549dc62640e28d6b966182c49d)
- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-db9483694fd3b0c1ffe08f84b55b48629af8dbf569a3f814638a67505b936044)
- access_info.rest_auth_info.query_params_auth.query_params

<a id="canonical-0501aa8ce5b454760846f2dcf9b862406f7d082d669b8cf90b622b691cba9ab4"></a>

Type: `"single"`. Computed.

The set of authentication parameters to be passed as query parameters.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-6d4ca6e25f329a5374458ca16da2e6add55e916dd7a5db36340ba966585fa5ab"></a>

## Direct properties — access_info.rest_auth_info.query_params_auth.query_params / 3769ab8bedd4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b081aa07b66a4255568527c0303a2d428fb73c4fdc5208e86fc227d059f643d9"></a>

## Next pages — access_info.rest_auth_info.query_params_auth.query_params / 3769ab8bedd4 / 4

- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-db9483694fd3b0c1ffe08f84b55b48629af8dbf569a3f814638a67505b936044)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9f31fd0fe9252534985205497187cd07b2da653a58f6c363c9627881c364ea5"></a>

## access_info.tls_config — access_info.tls_config / 96b3f0360e48 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- access_info.tls_config

<a id="canonical-a384e74f4940aa90820c9a746440fbf5623ea3a9f771ab76d11b27f76872eb4f"></a>

Type: `"single"`. Computed.

TLS configuration for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

<a id="canonical-59beb3294cb2018d66ec92597b66d8cab790b15c56fc3923d433498ffafc6efe"></a>

## Direct properties — access_info.tls_config / 96b3f0360e48 / 3

- [cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b): complete subsection reference.

- [common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392): complete subsection reference.

- [default_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-644d66cb750cdff30fe310ac7f09b70bb256131cb5dab058f80090079e10f9d4): complete subsection reference.

- [disable_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-05ef1e82a149a89d9c215e9c7b654582d9a45277c019a1aff70293288722b0b9): complete subsection reference.

- [disable_sni](data-sources--secret_management_access--reference--group-001.md#canonical-f70889dc65a1057184e4455669ac2b3b8684187b0f0b164f25e5f086ec0ba041): complete subsection reference.

<a id="canonical-8e1ff9ecf0e2bfe88fbd3891e8d7d513d978c6aea4166d72490e32297b55cb90"></a>

<a id="canonical-7cc0157e092fd9f50050541f89e92a47bc917f6c29eb736bcfe2a52d009e20cc"></a>

## max_session_keys property — access_info.tls_config / 96b3f0360e48 / 4

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

<a id="canonical-55d987e67ea437faf23c955c25aa66bd9aa4712bdb50068d7a74c8a4d3a29306"></a>

<a id="canonical-ae3dde7a558d0cb6bbcf231efd064792b0ee034abe8a2e29f550ec999abea164"></a>

## sni property — access_info.tls_config / 96b3f0360e48 / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_host_header_as_sni](data-sources--secret_management_access--reference--group-001.md#canonical-7beb298d92f6b0e6ae792e05f778e209c7f54dd934be42e3384306794fe96a86): complete subsection reference.

<a id="canonical-8c79eda385e0a9eeb569d98a3dc108cbb17ea07c0217f6871e952eb538d40865"></a>

## Next pages — access_info.tls_config / 96b3f0360e48 / 6

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.default_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-644d66cb750cdff30fe310ac7f09b70bb256131cb5dab058f80090079e10f9d4)
- [access_info.tls_config.disable_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-05ef1e82a149a89d9c215e9c7b654582d9a45277c019a1aff70293288722b0b9)
- [access_info.tls_config.disable_sni](data-sources--secret_management_access--reference--group-001.md#canonical-f70889dc65a1057184e4455669ac2b3b8684187b0f0b164f25e5f086ec0ba041)
- [access_info.tls_config.use_host_header_as_sni](data-sources--secret_management_access--reference--group-001.md#canonical-7beb298d92f6b0e6ae792e05f778e209c7f54dd934be42e3384306794fe96a86)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55f5c07011339efd606f9e821bf5cf1ec047d065291231f184ceacc20fd68a3a"></a>

## access_info.tls_config.cert_params — access_info.tls_config.cert_params / 8c8a23a93d15 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- access_info.tls_config.cert_params

<a id="canonical-a76b118600c10db88a94cd701b88c98bd835661b526e14fcfc8e7e766f39f5ea"></a>

Type: `"single"`. Computed.

Certificate Parameters for authentication, TLS ciphers, and trust store.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"tls_validation_params\",\"volterra_trusted_ca\"]"
}
```

<a id="canonical-8716bfbf92be296e2f83ac91d5e994fb03258aaaaf9e2a35e3d733b50ee4220c"></a>

## Direct properties — access_info.tls_config.cert_params / 8c8a23a93d15 / 3

- [certificates](data-sources--secret_management_access--reference--group-001.md#canonical-eb532112162684cd4bf09664a2278937109d937aa4da1e73ece758de53281574): complete subsection reference.

<a id="canonical-48316376c0d238fd9bba607edb1bdb7a39c31b52ee8338d35c20eeaa2ed4ef05"></a>

<a id="canonical-1b8258282be86af232e97e622ddb0ec999f877a7a1cc9048aa34cbd7d9c9b938"></a>

## cipher_suites property — access_info.tls_config.cert_params / 8c8a23a93d15 / 4

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3041797e0abd524b7885571277426cdf169db298ad9c7c0b29f4e91e214a46e4"></a>

<a id="canonical-6e08d0cca9e223806b64d143e24db6d7e12f4e35fe855f89d981b8278e4ae947"></a>

## maximum_protocol_version property — access_info.tls_config.cert_params / 8c8a23a93d15 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-381ffd9c422952b6237d0418e54c2922702f333bd68558dce5173c31b5b310a2"></a>

<a id="canonical-d46e7d436dcae0bc6bda8aaec942a6d4cfd98c23536c293776189b4fff709471"></a>

## minimum_protocol_version property — access_info.tls_config.cert_params / 8c8a23a93d15 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [skip_server_verification](data-sources--secret_management_access--reference--group-001.md#canonical-9f25c8f76f5fbb2b37015788a8971f5b0ca006b5d747e4e4a291d1e5aaeb91f0): complete subsection reference.

- [tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-26651825f56890d2fb29c104680f48ab9f7c9826e351e99514166c165c8be7cd): complete subsection reference.

- [volterra_trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-44cc77abe1781298247f0924db3cb45754de730e626f5d83b9928901ed4b57bb): complete subsection reference.

<a id="canonical-66dd9824d31e29ebe8a9790a4533ab0dc39538ddfcc891a719d0a6a9c1a18013"></a>

## Next pages — access_info.tls_config.cert_params / 8c8a23a93d15 / 7

- [access_info.tls_config.cert_params.certificates](data-sources--secret_management_access--reference--group-001.md#canonical-eb532112162684cd4bf09664a2278937109d937aa4da1e73ece758de53281574)
- [access_info.tls_config.cert_params.skip_server_verification](data-sources--secret_management_access--reference--group-001.md#canonical-9f25c8f76f5fbb2b37015788a8971f5b0ca006b5d747e4e4a291d1e5aaeb91f0)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-26651825f56890d2fb29c104680f48ab9f7c9826e351e99514166c165c8be7cd)
- [access_info.tls_config.cert_params.volterra_trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-44cc77abe1781298247f0924db3cb45754de730e626f5d83b9928901ed4b57bb)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-eb532112162684cd4bf09664a2278937109d937aa4da1e73ece758de53281574"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71d96ab7812e3adec3d008a40287b786bcbb6b03faf2a951da1e7f10ba864524"></a>

## access_info.tls_config.cert_params.certificates — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- access_info.tls_config.cert_params.certificates

<a id="canonical-9498b8985f31619bb1a9f5aa76608327e79469aaec72f529162abc76c1cb0e49"></a>

Type: `"list"`. Computed.

Client TLS Certificate required for mTLS authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-789519e3f9e8da4ccc841f6e1ab751b40ef857e733e1c3fb2d475d04d2bef4b4"></a>

## Direct properties — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 3

<a id="canonical-17745fb37774a369615629b15fc5b44679cb94f90cdaa1b4b7831cbcfcdae684"></a>

<a id="canonical-080e3a9219c825759b6691e3171fcdff30314da137ba5f04b602bb12b478c1a0"></a>

## kind property — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 4

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

<a id="canonical-4a3627d87ac870cd8304613ed6d058153c33eb6624baa957b254f930041ffaa4"></a>

<a id="canonical-e1c17d8fb038a6e2919c02b51221674f160741239bb4b57c81639265df5ebc12"></a>

## name property — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 5

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

<a id="canonical-b272fdca9f1efc345d1ed082a257e209089c4bf743510c8e0006bec3922cf519"></a>

<a id="canonical-603d7d954c14106e0a705550462717c411807b3c127daf7efc50360924434d63"></a>

## namespace property — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 6

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

<a id="canonical-8b7c012195ccece82f0b446419e403207ec5754c6aadfe12090b36287fda16f8"></a>

<a id="canonical-6a300dab16f7b38bf54d98ec72e0cadf671dbf5df90bcefb41c1dd0b9e51263c"></a>

## tenant property — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 7

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

<a id="canonical-f1cdc04f07f7b0003f7de81fdf6bc9a282b0d590e547804c1e0933d06b61570a"></a>

<a id="canonical-701500d15631379b80fefa248d60288fc275c9c460450729869ac8cb0ebef0d7"></a>

## uid property — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 8

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

<a id="canonical-358fdfced6cc288443d5d88804ccb3e4f523e6ec768e610174f3deaf4630ccb4"></a>

## Next pages — access_info.tls_config.cert_params.certificates / 2f68e4185072 / 9

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-9f25c8f76f5fbb2b37015788a8971f5b0ca006b5d747e4e4a291d1e5aaeb91f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d32601aa17759d5048c0b72617b188f0542803b07138054715b35273f914b8c"></a>

## access_info.tls_config.cert_params.skip_server_verification — access_info.tls_config.cert_params.skip_server_verification / 0f16aa9d9016 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- access_info.tls_config.cert_params.skip_server_verification

<a id="canonical-9ec2daaf48a5932529fe3fa31f74df0009a1ac02a5226ff2a18f75f895218596"></a>

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

<a id="canonical-50fe0532a8c3952629ceb893ea19dc97c8fa1d9ad41ba7ce685442ee68e6c9c3"></a>

## Direct properties — access_info.tls_config.cert_params.skip_server_verification / 0f16aa9d9016 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0962293e0f70d20501d2441426a57e4345bd9c080a5758d4887569bed98fe84b"></a>

## Next pages — access_info.tls_config.cert_params.skip_server_verification / 0f16aa9d9016 / 4

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-26651825f56890d2fb29c104680f48ab9f7c9826e351e99514166c165c8be7cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54cd86682f3dfcc6e20f77ee42f62162dbee04450e09d8b424779c45d34de93a"></a>

## access_info.tls_config.cert_params.tls_validation_params — access_info.tls_config.cert_params.tls_validation_params / da9bb2529676 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- access_info.tls_config.cert_params.tls_validation_params

<a id="canonical-77787f714ad64acc0032ebedb51a052dd633b1f2a22edef949bb59e5cdceea0e"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-52e3fb43a6fcd2ae7bfc1f7f88e96c22a4d0aab3690e1a739196507fe865f606"></a>

## Direct properties — access_info.tls_config.cert_params.tls_validation_params / da9bb2529676 / 3

<a id="canonical-4ed24bd03c7a386a6549c8fdf246ca46cf88277d2f9e5111a00a90539ddb5b0e"></a>

<a id="canonical-de8160c0350c8803919a8f8e14ef7f76b634897294c051f40171e18f18cfc4cc"></a>

## skip_hostname_verification property — access_info.tls_config.cert_params.tls_validation_params / da9bb2529676 / 4

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-090c6b52574493445bcfc3ab6cc98a0def338aa30ce278a2defa89c0cb7765c6): complete subsection reference.

<a id="canonical-97ab98186155e71ca397829383973e05c3ccbf78a63aba7b9080438c36673b6d"></a>

<a id="canonical-84f941eaf08b08336b4ac5647bb941ecc3c8f1563bb91420546f8118853b94d6"></a>

## trusted_ca_url property — access_info.tls_config.cert_params.tls_validation_params / da9bb2529676 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-02cbcf2a5cfb0827ef43525574ce614bfa34078d6731341a2c4fdf0afc2141c0"></a>

<a id="canonical-d7f9275c7844a632e6236be98bfce5315ae11c2b9c260ba5ec4f936525ffdb6b"></a>

## verify_subject_alt_names property — access_info.tls_config.cert_params.tls_validation_params / da9bb2529676 / 6

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fa75115d26b0652f749b3ff7287d5190b47aae77c7aa1f06043e4b525e2122bf"></a>

## Next pages — access_info.tls_config.cert_params.tls_validation_params / da9bb2529676 / 7

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-090c6b52574493445bcfc3ab6cc98a0def338aa30ce278a2defa89c0cb7765c6)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-090c6b52574493445bcfc3ab6cc98a0def338aa30ce278a2defa89c0cb7765c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2689b98aca82a5f24d4d156bbe4e1e119f23627892a302fead085f13fa6826f2"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca — access_info.tls_config.cert_params.tls_validation_params.trusted_ca / 7b9d3f270ccb / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-26651825f56890d2fb29c104680f48ab9f7c9826e351e99514166c165c8be7cd)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca

<a id="canonical-c88a9eb2447c9551d39d730d1e00fd9ad3e0af00b9f805a6fa0d1feca7d5670f"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-503a1074358510e267abf263b61108928e392de7ed6891ebc49d456e2b563dac"></a>

## Direct properties — access_info.tls_config.cert_params.tls_validation_params.trusted_ca / 7b9d3f270ccb / 3

- [trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-147565656959b3caf1d015273538eaaf28e6856ba2829e3e478378a7ee266793): complete subsection reference.

<a id="canonical-2e2f5294d59ddbfc421d53c7d1d3c4cc93331f9c49d3c2b3910b2b5a80d1cb16"></a>

## Next pages — access_info.tls_config.cert_params.tls_validation_params.trusted_ca / 7b9d3f270ccb / 4

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-147565656959b3caf1d015273538eaaf28e6856ba2829e3e478378a7ee266793)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-26651825f56890d2fb29c104680f48ab9f7c9826e351e99514166c165c8be7cd)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-147565656959b3caf1d015273538eaaf28e6856ba2829e3e478378a7ee266793"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ebf09dfd33b6ecb4dfd15396fee601c39166dedb93be3160f11b963c2f66039"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-26651825f56890d2fb29c104680f48ab9f7c9826e351e99514166c165c8be7cd)
- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-090c6b52574493445bcfc3ab6cc98a0def338aa30ce278a2defa89c0cb7765c6)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-05fd02548e9bfa4b75068217ef6e7c22a618efb927410b2c99a4b63b1c3583d4"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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

<a id="canonical-64caf41443020d25ae9c14fa047dcacd8cc0a001def5550215bb5644b651b2d3"></a>

## Direct properties — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 3

<a id="canonical-363b1024d329ce8c5d2165c0a12cb35176a10b5e81c68a261b9efd718a24a4e0"></a>

<a id="canonical-387f79e739897ce0bfb7b724cce8f261d7a3d161ef997de32fb69bda8ac74b61"></a>

## kind property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 4

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

<a id="canonical-74a117621005cb39cff3d59b1412edc295729c6eab726a9e73d78254c14212fd"></a>

<a id="canonical-6eb3fcfadca9c559c20cc8a9f23259004f37f10f0b418566928bc077dcd41dc8"></a>

## name property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 5

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

<a id="canonical-9f5caac418d49731cb010b40b21f6fe3c52f4ad327f7ef5a01c1b4c85c8b4b0c"></a>

<a id="canonical-fe86a9921e42e272576ba1f8269bfa87d16faf317f203cd46d97e78fc91951b5"></a>

## namespace property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 6

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

<a id="canonical-fb4128b9be3488a14fb25fde22ad5126003dfeec7de3e4821f285cf6dc791d16"></a>

<a id="canonical-e0b0cfdc80ec9dced577ab0bd34f9878d559bdd947ded75f3d7c0b3560b76b3d"></a>

## tenant property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 7

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

<a id="canonical-73ecfbce1cd6902ffd830b69f6cef5a2cef5947618e30270b8265ea8ae41cfe7"></a>

<a id="canonical-2f408223c6961f46884b9bc6295aa69790de7231378eb1dbf7bf07fffc77eab7"></a>

## uid property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 8

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

<a id="canonical-8a21ef0c9c635a6941994a5bb6d90524e98d010efbd6847928189dff6e82d301"></a>

## Next pages — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 4a18f2f63084 / 9

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-090c6b52574493445bcfc3ab6cc98a0def338aa30ce278a2defa89c0cb7765c6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-44cc77abe1781298247f0924db3cb45754de730e626f5d83b9928901ed4b57bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-110f7587b5c24d2c3dfaf6711de18dce4ae05363b2f7b97b057b0340614e9a3e"></a>

## access_info.tls_config.cert_params.volterra_trusted_ca — access_info.tls_config.cert_params.volterra_trusted_ca / b6ebef7fd1de / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- access_info.tls_config.cert_params.volterra_trusted_ca

<a id="canonical-1b5b4ef1e29651049c95c93bb0561b9d47b0a124dce3e406747dae81bec2b5ef"></a>

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

<a id="canonical-52b168b6e5cadea99867199b4532364e86edf536ae78607ca12dd1384429e50b"></a>

## Direct properties — access_info.tls_config.cert_params.volterra_trusted_ca / b6ebef7fd1de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f83c67a51c6d4d8dc4a5f750787fa2fcfca6917d668fcc2b292e61289dc3b69"></a>

## Next pages — access_info.tls_config.cert_params.volterra_trusted_ca / b6ebef7fd1de / 4

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-dd6bdf2921ef138c0357703e4bcededf627bf9e8a44adccea69d2be1f24fa72b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c23eed340cf76fdcb272f0031c99b5906df5b76796d3b5253662be736eab8e02"></a>

## access_info.tls_config.common_params — access_info.tls_config.common_params / e1ac6c5ed645 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- access_info.tls_config.common_params

<a id="canonical-538da24ac2a26120dd8d06091dab2b567ee4327fb070fa0f804428f5fb4860ba"></a>

Type: `"single"`. Computed.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Upstream description:

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3acfbd61ae73959f66dc8b6deadefee50010b01b9e84765f3786735f9b93e19f"></a>

## Direct properties — access_info.tls_config.common_params / e1ac6c5ed645 / 3

<a id="canonical-ece869c94b0cfe67241956345a7230a9da03a2555d656fd1f6d0e1e99ccac656"></a>

<a id="canonical-9d7e52ce1590c290047ff994210fb6e77fa4e345e8b2d2b245b19e9dc9aac331"></a>

## cipher_suites property — access_info.tls_config.common_params / e1ac6c5ed645 / 4

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3bc513df940e820e3b92afe8b182048590c2e83cddf22ee9044b5691c057bbe3"></a>

<a id="canonical-cb6ad3ca93254022dd917304edeb45ef7ac3cf2e0f6d9228dd37a3d840a80649"></a>

## maximum_protocol_version property — access_info.tls_config.common_params / e1ac6c5ed645 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4eecd6531c4f1d554020043b157476617a5d3ef2043de5e21eacb7018182f42d"></a>

<a id="canonical-414c9f97ea662ec0fc2e9e4254493155b677d7a9ce824a5bdaa93fcc889341b8"></a>

## minimum_protocol_version property — access_info.tls_config.common_params / e1ac6c5ed645 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6): complete subsection reference.

- [validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-24dd0516da482900659b52dfd43222fff2673ae8383bfe1e37842faa90d63692): complete subsection reference.

<a id="canonical-06c8ba0e13cc99382a56b382b260bb9b6cf72e4b4913a4b23e241ee9ede1d02f"></a>

## Next pages — access_info.tls_config.common_params / e1ac6c5ed645 / 7

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-24dd0516da482900659b52dfd43222fff2673ae8383bfe1e37842faa90d63692)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-089d0af77a9fd778557fa7e7aac51c0524ba06b1e433e7a2f94712cf29b17f2f"></a>

## access_info.tls_config.common_params.tls_certificates — access_info.tls_config.common_params.tls_certificates / 280bde37d0a9 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- access_info.tls_config.common_params.tls_certificates

<a id="canonical-fbc82c63b3485aa65a550b1fa2e66a0574c69de800f6da77105863916a0484ed"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7b09b8c1717300a5a0aecd1647381dfacc06cbd03e9ccbe83c27e5ebaa45d9d7"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates / 280bde37d0a9 / 3

<a id="canonical-87712c4827117039acc251919492bdf38b037eb635f5fe3c0bb2aefc48a09519"></a>

<a id="canonical-3253f54dc092a4f377ef139e4d4845babec50e2f05f5709cef5a1f854bcdee02"></a>

## certificate_url property — access_info.tls_config.common_params.tls_certificates / 280bde37d0a9 / 4

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

- [custom_hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-143ae8f46afdb61fc77b926820c99775739fc7d005c298c90ecdee2a753c018f): complete subsection reference.

<a id="canonical-8a797c7655b8bbb369f9321b84fa633e21f58547e2c2370629c60a90c42fad97"></a>

<a id="canonical-da6111ee8b2318c7fde537cf7944c3dc2259504721649a93c3c16cb892b74882"></a>

## description_spec property — access_info.tls_config.common_params.tls_certificates / 280bde37d0a9 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--secret_management_access--reference--group-001.md#canonical-032bcf378cba192b220cf610acc94c113defe27ed414fd08ed155d891af5229d): complete subsection reference.

- [private_key](data-sources--secret_management_access--reference--group-001.md#canonical-71b3aab872e5b99c4bfb3ca7cf386582a253e9be055351b8bef5ca0f827c925a): complete subsection reference.

- [use_system_defaults](data-sources--secret_management_access--reference--group-001.md#canonical-dc30aad23b3e3398c84387fb19b358e37fd1c90eef6d676aa4aef7af88b456ea): complete subsection reference.

<a id="canonical-b8b6749b7ab1805aa50cc8c97cb6d97ef3aebbc1a9e0f6c3934700c546ebb291"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates / 280bde37d0a9 / 6

- [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-143ae8f46afdb61fc77b926820c99775739fc7d005c298c90ecdee2a753c018f)
- [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](data-sources--secret_management_access--reference--group-001.md#canonical-032bcf378cba192b220cf610acc94c113defe27ed414fd08ed155d891af5229d)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-71b3aab872e5b99c4bfb3ca7cf386582a253e9be055351b8bef5ca0f827c925a)
- [access_info.tls_config.common_params.tls_certificates.use_system_defaults](data-sources--secret_management_access--reference--group-001.md#canonical-dc30aad23b3e3398c84387fb19b358e37fd1c90eef6d676aa4aef7af88b456ea)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-143ae8f46afdb61fc77b926820c99775739fc7d005c298c90ecdee2a753c018f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bd80a303b224f85777ccc92d214b2785c64ecf1e195bea2ab7d4681658c9a8d"></a>

## access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / 0a49644cb833 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-07eb2c6eafcd6be232326a2a28178e6c335271ceeee75d3522b3bfef90adf203"></a>

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

<a id="canonical-bd9749bf01ab74ed9ca5e71f4257c16452d15a5721dcb83274d65b6580b8ed34"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / 0a49644cb833 / 3

<a id="canonical-b028bfec2c4016ec6c1f8900c21822087eb70423507455c1cf25f7b44b0de466"></a>

<a id="canonical-3bb337709e242c3dd37f0ce3c76ec36421b66e7769c49bc00a53b101929a0392"></a>

## hash_algorithms property — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / 0a49644cb833 / 4

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

<a id="canonical-95c685a309f9d795d06b5bc279c0b39be4b97201a27342ee7199defc6e3a223a"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / 0a49644cb833 / 5

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-032bcf378cba192b220cf610acc94c113defe27ed414fd08ed155d891af5229d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43e65ec20e740b247cb72c7b662278b9f70248e32aab5e064e7b41b87c8f0c10"></a>

## access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling — access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling / 121ae38810c1 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-b12e840438a1ae54988c80f53ffc5b844edce0b03f9d69b6a1ec8e8cc70d1471"></a>

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

<a id="canonical-4d142d3337ed16d04e8fa2c7eaf35f410917d2e5149d01910afe0722cbb9eb9f"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling / 121ae38810c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2355e2e1c6e6a9d39a83962c445da8f9b8e54fa5725a11ad4a4d23ea6424519"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling / 121ae38810c1 / 4

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-71b3aab872e5b99c4bfb3ca7cf386582a253e9be055351b8bef5ca0f827c925a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7617ce9b0cd28191769e75de23ae8b39072c38d66da833419f560c52616f8028"></a>

## access_info.tls_config.common_params.tls_certificates.private_key — access_info.tls_config.common_params.tls_certificates.private_key / 104d61745831 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- access_info.tls_config.common_params.tls_certificates.private_key

<a id="canonical-09cb66af96a92c407dad4fc15e38712b7a47644ec45e5e6929df38b57685e69a"></a>

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

<a id="canonical-c491b52f1d4ca8d2f1db0e99bcef6498024dce317a716982cc81764f1d5fa9e8"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.private_key / 104d61745831 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3883eee39a01ec08d56738a71b074d29667e2eba8e3e9d8fef25ca04194efa2a): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3f25af6ee18b6030513847401d23f987a442212bc75905bff5e26b3a7649fcb8): complete subsection reference.

<a id="canonical-97f5288c3e9279a3f220186a79c8b36a4419f9a3baab0ae726599d018313bd8d"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.private_key / 104d61745831 / 4

- [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3883eee39a01ec08d56738a71b074d29667e2eba8e3e9d8fef25ca04194efa2a)
- [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3f25af6ee18b6030513847401d23f987a442212bc75905bff5e26b3a7649fcb8)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-3883eee39a01ec08d56738a71b074d29667e2eba8e3e9d8fef25ca04194efa2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a235c0bc752b7cfca4211e8586b90fbc38411148a135f8a6b273f909d48bc70c"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / f217793baa4f / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-71b3aab872e5b99c4bfb3ca7cf386582a253e9be055351b8bef5ca0f827c925a)
- access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-a611cf9a734fa16acc78b77663026840ccc2372b09c1da947f2ef195a65a6173"></a>

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

<a id="canonical-5c262fc273c7e982d245f2a35245fb7db353242ef3264344dfd339e706445a28"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / f217793baa4f / 3

<a id="canonical-5750f4b07d4c036e13f9181e89767a94b3e750b820c362c1104fbfa6b6976393"></a>

<a id="canonical-ae7cd8b090767a226dfd52fedd00a814d6ff1ec81d8b4d221da4df2fdb798375"></a>

## decryption_provider property — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / f217793baa4f / 4

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

<a id="canonical-6c2bb81e4a1c54292777e0ea6a915b6deeed89a1992f0d2fd619b9f80a93c103"></a>

<a id="canonical-dee1630a7ebb05621f3fcb88907801a1884cffb6e73f84c5344c1603cde3de1c"></a>

## location property — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / f217793baa4f / 5

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

<a id="canonical-219462b49b0f4b419f8f4c59b20862623c7c6fe47dd69c0d1ddf1a66d4071510"></a>

<a id="canonical-e86dc96fb37a7ea4d97d98159ce684cc803b014aaf0eb4b75b87e89bb6a99875"></a>

## store_provider property — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / f217793baa4f / 6

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

<a id="canonical-f0f16301f9e716df46dae04973ff56dcb68f6adf3c1f3b5dd460fa01dbd6324d"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / f217793baa4f / 7

- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-71b3aab872e5b99c4bfb3ca7cf386582a253e9be055351b8bef5ca0f827c925a)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-3f25af6ee18b6030513847401d23f987a442212bc75905bff5e26b3a7649fcb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e791a093b50fdb4b12efd793d29694674e8f983433b6713f665fa89b2017dbcb"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / 5467aaf8f805 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-71b3aab872e5b99c4bfb3ca7cf386582a253e9be055351b8bef5ca0f827c925a)
- access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-d2a1b45839fbce48437574c5400bc4cfd62dedce5fcb6590799ae67a5ba5fc89"></a>

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

<a id="canonical-505225c17915adf76e5392b2619153e4009ab0a76aa17dd5b793952a4989d392"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / 5467aaf8f805 / 3

<a id="canonical-18e224f43ba39ad3e19d2c16ee6fa7ab85d6895711033712cef942ae32958eda"></a>

<a id="canonical-4728dd22b551bc335334adc1b4be1c77bb0411d998d80907e62229420cc22b4e"></a>

## provider_ref property — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / 5467aaf8f805 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f9999f908cdbf60062a36ac3538fd074afcbfa55983e5b26f93af1ad860d9615"></a>

<a id="canonical-3aa433c0ce4affb6a5412cb3520e08613e92c04e1758079a747f908ae0ea8929"></a>

## url property — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / 5467aaf8f805 / 5

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

<a id="canonical-a21c7055d6f0eef9e86a7d537449d2449c73bcf8e51c8294d9d5d00a98d9b3dd"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / 5467aaf8f805 / 6

- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-71b3aab872e5b99c4bfb3ca7cf386582a253e9be055351b8bef5ca0f827c925a)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-dc30aad23b3e3398c84387fb19b358e37fd1c90eef6d676aa4aef7af88b456ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76848e2965f974315c68ce3eba84bc124a2a11428c9e09e7e523d4cb0e836397"></a>

## access_info.tls_config.common_params.tls_certificates.use_system_defaults — access_info.tls_config.common_params.tls_certificates.use_system_defaults / 10d3461f7074 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- access_info.tls_config.common_params.tls_certificates.use_system_defaults

<a id="canonical-2930506f6874552996c091af31295595e4c688af6715ab632e055398d95ecf7c"></a>

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

<a id="canonical-b259bb9fd6b8db82adbf7a6e5764c57d7af812ae5f2d0dbdf7b0f9318f90f7de"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.use_system_defaults / 10d3461f7074 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5bbcd3536fa31082f96365e4c54b302d507041870374990b8630d63e853bb79"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.use_system_defaults / 10d3461f7074 / 4

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-8f06d0ec0c943eddec13867fdfb1595bd24f7955697b7c740ba8bdf633ae4eb6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-24dd0516da482900659b52dfd43222fff2673ae8383bfe1e37842faa90d63692"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9744248cf88f48fe6792d3f28613df2943b300979b408dfefcd1ecaa5a3f897"></a>

## access_info.tls_config.common_params.validation_params — access_info.tls_config.common_params.validation_params / d395f5323cfe / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- access_info.tls_config.common_params.validation_params

<a id="canonical-b48873f832bc572a35909f5a4d30fa6452451f5a5e7df20f9f1499105c7bcb69"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-23e43977cf9ecdc4acedb089e2a6f6fe3998748565f4a384bc20266b49b373a5"></a>

## Direct properties — access_info.tls_config.common_params.validation_params / d395f5323cfe / 3

<a id="canonical-0a5ed0be7e2b10568f4055469bc7678947c8e2dc7845f393b741caa2e342e573"></a>

<a id="canonical-5250d8414d5f7df7d006f16469f4560884f59d43c5b45b80e987665383c404ac"></a>

## skip_hostname_verification property — access_info.tls_config.common_params.validation_params / d395f5323cfe / 4

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-fb2aeb5922041db3ad414b399c6af453ed43769643aab49c5918ec6e278cb80c): complete subsection reference.

<a id="canonical-9f8690272dd70f28161d06c6723ba4177898e05cdf0a8957492a334802b541b4"></a>

<a id="canonical-dc84a038bfc037074e54edba599ee1d0e9c9cff34270d7d833b9cf0a3bfe1eaa"></a>

## trusted_ca_url property — access_info.tls_config.common_params.validation_params / d395f5323cfe / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-4eb0fb54e1151c2bcc16752768b5145703e116becd35efa8f3c3646a6e66dc83"></a>

<a id="canonical-cb244c6016bda95c50e4c8459222a2873aea3e366ed83785f3df5c48b8864b9e"></a>

## verify_subject_alt_names property — access_info.tls_config.common_params.validation_params / d395f5323cfe / 6

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6c326e7c355e69bd548dd71ce9d71d1bf444810d323b82448d75192fdbb76241"></a>

## Next pages — access_info.tls_config.common_params.validation_params / d395f5323cfe / 7

- [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-fb2aeb5922041db3ad414b399c6af453ed43769643aab49c5918ec6e278cb80c)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-fb2aeb5922041db3ad414b399c6af453ed43769643aab49c5918ec6e278cb80c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-003d9b4ed66c4c0274c2a2399c914032d47f2f41c40a52368cd697d793cbf4d1"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca — access_info.tls_config.common_params.validation_params.trusted_ca / 3908033ccc7f / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-24dd0516da482900659b52dfd43222fff2673ae8383bfe1e37842faa90d63692)
- access_info.tls_config.common_params.validation_params.trusted_ca

<a id="canonical-162533777383e6c8de6fdbb34dd188a8f713642aa6fade49111a5035f49d2a93"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3378eb6b93c03934741bc8d73b5febef427efbf169168e5587e4bbb1f92d6272"></a>

## Direct properties — access_info.tls_config.common_params.validation_params.trusted_ca / 3908033ccc7f / 3

- [trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-e7adef8d0404496bb427a7f429b27f86cb6b39f691f48b4375357c1dbb158b80): complete subsection reference.

<a id="canonical-93f31d5a63c8e96fee30160a8a7bca664aaa8d7d8e359688f7b99a5db996d222"></a>

## Next pages — access_info.tls_config.common_params.validation_params.trusted_ca / 3908033ccc7f / 4

- [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-e7adef8d0404496bb427a7f429b27f86cb6b39f691f48b4375357c1dbb158b80)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-24dd0516da482900659b52dfd43222fff2673ae8383bfe1e37842faa90d63692)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-e7adef8d0404496bb427a7f429b27f86cb6b39f691f48b4375357c1dbb158b80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ed8828af12c21145f3ecd0edc888cc5616d846e3b64083a8df9677d1e938fe8"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-001d4c0652c93f6b8d6edcdc62f825ae34ed020eab2938d905d3c36d5260b392)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-24dd0516da482900659b52dfd43222fff2673ae8383bfe1e37842faa90d63692)
- [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-fb2aeb5922041db3ad414b399c6af453ed43769643aab49c5918ec6e278cb80c)
- access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-53b7e4657ff17c184c198379ddfbc48fb3280d4c0acd564bfa8db94ca07beb8c"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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

<a id="canonical-b3664abd6002b5bdbcc659d529df50d882a797405acd5f97ad03b093812558b4"></a>

## Direct properties — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 3

<a id="canonical-7fdf12e96b7d95220da06bcf3f9ccc200e09e5e610d43736c068c982eadede24"></a>

<a id="canonical-4b88b29730a9fe6d30f57aada6084b8e638b926f8af89fe8f3b73abcbdf5e86f"></a>

## kind property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 4

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

<a id="canonical-10d730b65dcd7065ff3b43f55d9d49d094abc99e8241f928832752087fe23111"></a>

<a id="canonical-084b256a32ed3799cdc0f839d0d5e8bf287809426e0e8adb6959f13ad744d76f"></a>

## name property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 5

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

<a id="canonical-3ade70fb627f0f7e985c1670e76ad92dcf8bb55eef0d0d393cc5627d36eb45fe"></a>

<a id="canonical-926b273ead57e559de88a41eaddd8309ba2f9291085820c2a171e310a4871eab"></a>

## namespace property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 6

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

<a id="canonical-f5dd89b185c6b2be66892777be41322646e31c265d024ca6fa28d09a46ea2aa6"></a>

<a id="canonical-850bde6d0c73dcdb6c2fd6707cdd7cf4bba8f6f27065413f4a11cd497166aa35"></a>

## tenant property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 7

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

<a id="canonical-3dffa2db891741d3f8c89716dd7e0a66054a19067630abaeba01fbbb0a6d2664"></a>

<a id="canonical-e156d186ceecf97b8382c73740acf951733f10c75984d18991745a0812ec979d"></a>

## uid property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 8

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

<a id="canonical-b7a1c383aead0cbee7eb893f7944a2201ab4e3674cab2544838be18af2ced251"></a>

## Next pages — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / 08edb2785545 / 9

- [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-fb2aeb5922041db3ad414b399c6af453ed43769643aab49c5918ec6e278cb80c)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-644d66cb750cdff30fe310ac7f09b70bb256131cb5dab058f80090079e10f9d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2631c63ec29bb6f270b5b35fb6021ed2910c373eee6fb0edcf4e997de62cdf28"></a>

## access_info.tls_config.default_session_key_caching — access_info.tls_config.default_session_key_caching / 633a2c95a462 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- access_info.tls_config.default_session_key_caching

<a id="canonical-9de4e21b2eb640b48d5ca324d9d20bfaff9257dd00d3ddf4bcc4541efb9779ad"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching.

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

<a id="canonical-429ce4d9a2ba2a8b252ea84d32a90f4c637208a79002b3e7eb423a85d42e3a79"></a>

## Direct properties — access_info.tls_config.default_session_key_caching / 633a2c95a462 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34a7e8520282c3f1c6ad3fadfd6992b2aa58ec55a1ff88dae4d73c68296a5661"></a>

## Next pages — access_info.tls_config.default_session_key_caching / 633a2c95a462 / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-05ef1e82a149a89d9c215e9c7b654582d9a45277c019a1aff70293288722b0b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe3472dd1a0bc3c82dc347066333ac560b89a65aa01b8f87b9d5c847a88a10f3"></a>

## access_info.tls_config.disable_session_key_caching — access_info.tls_config.disable_session_key_caching / 3c5eff37fd1a / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- access_info.tls_config.disable_session_key_caching

<a id="canonical-8e1082066faf0e83671caea9a291113fd59fef162836b7429635a23f66c740ab"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-fd0cd3e28eaa39d2dc985a623b78e70ffcf60dee9ca9e9b7716eed7651acf89b"></a>

## Direct properties — access_info.tls_config.disable_session_key_caching / 3c5eff37fd1a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01ba4ef64765f4c28923a9d49d4a3a388ddf90e5b2dd87b6cc6c035a68f04916"></a>

## Next pages — access_info.tls_config.disable_session_key_caching / 3c5eff37fd1a / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-f70889dc65a1057184e4455669ac2b3b8684187b0f0b164f25e5f086ec0ba041"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb4ca510c3860bd5641dafbfabf6fe73be64c57a402001334100e3c88f33e187"></a>

## access_info.tls_config.disable_sni — access_info.tls_config.disable_sni / 88c3f58342db / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- access_info.tls_config.disable_sni

<a id="canonical-00f159ddcdac4a649cfb08cf61ed9041185e85f5b8934dccf4d797b11d387bfd"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-17797214aee4f27c82d463a4a114cd7e698aa4c05a3d062044b685d0c5e24f9a"></a>

## Direct properties — access_info.tls_config.disable_sni / 88c3f58342db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-decb9cc7157e76a931167e71ac0cfb685065b6c0b760bdd3614820eca406ca70"></a>

## Next pages — access_info.tls_config.disable_sni / 88c3f58342db / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-7beb298d92f6b0e6ae792e05f778e209c7f54dd934be42e3384306794fe96a86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83ddcfadaa7bcc35fe4889aa82fe039fd9a4df6967d44221d17175fff15218c2"></a>

## access_info.tls_config.use_host_header_as_sni — access_info.tls_config.use_host_header_as_sni / b0c8b439d8fc / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- access_info.tls_config.use_host_header_as_sni

<a id="canonical-27217a92c65cccbf874e8bda0dba5c36165d613e3cfe3e9702c02482533f6e84"></a>

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

<a id="canonical-00c6a4d3bda764214a71f49fa516f51e44aadeae093f227dccf38d50244d1f21"></a>

## Direct properties — access_info.tls_config.use_host_header_as_sni / b0c8b439d8fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-67a9d25360830fa3dd8fff1276a7de43d3d2941fba572fb6aab191388c7e45a1"></a>

## Next pages — access_info.tls_config.use_host_header_as_sni / b0c8b439d8fc / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-d6c347357ad0e04a5988b8ff076bc99e72e19707416a494ba3e5a59e5c7d832b)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63beb238d85221210401efd5e8fe0950cd370bf594aa62932b54a5c9630cda49"></a>

## access_info.vault_auth_info — access_info.vault_auth_info / dd18ac3aeef4 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- access_info.vault_auth_info

<a id="canonical-d91d6f254ddc58a55327dca4012ea4ef6e51fffdbcf5724d1ee559ffec403c5e"></a>

Type: `"single"`. Computed.

Authentication parameters for Hashicorp Vault hosts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"app_role_auth\",\"token\"]"
}
```

<a id="canonical-5fa4d38a2488afcf0a06d1abd1afa23997846c09397d6cab9470a4126b381e6d"></a>

## Direct properties — access_info.vault_auth_info / dd18ac3aeef4 / 3

- [app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-06c78926d465b5c09ae4a9137a0981c0b52426d1d96d985fce7e04c79f60070a): complete subsection reference.

- [token](data-sources--secret_management_access--reference--group-002.md#canonical-a0a97aa2e4da9e47ecc9104e222afad13599ad2ee0c29b3d97ac4085a25161a3): complete subsection reference.

<a id="canonical-d859109bd039f3f9cae3b57076bfcc4ed8186a6c2dddf3f0e84783ea0cc752de"></a>

## Next pages — access_info.vault_auth_info / dd18ac3aeef4 / 4

- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-06c78926d465b5c09ae4a9137a0981c0b52426d1d96d985fce7e04c79f60070a)
- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-a0a97aa2e4da9e47ecc9104e222afad13599ad2ee0c29b3d97ac4085a25161a3)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-06c78926d465b5c09ae4a9137a0981c0b52426d1d96d985fce7e04c79f60070a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a82922596036a9f59a6810194df90ecc6ee7153d4efa3580e393062bee45effa"></a>

## access_info.vault_auth_info.app_role_auth — access_info.vault_auth_info.app_role_auth / 5ec5678374f5 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- access_info.vault_auth_info.app_role_auth

<a id="canonical-83c35b0ed1f0a55831e37180fe71177b8768650b899b374fd15b07633ccd80c8"></a>

Type: `"single"`. Computed.

AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d7bdecb7231fda062c58021a8b2b1fc1845c5051fded8e0ff770d8d39b60ada1"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth / 5ec5678374f5 / 3

<a id="canonical-c3fda162312388d2733957f63ad6c9ecbb67bd3d4cc134ad66e7627f0a3392e6"></a>

<a id="canonical-170b958aa196e4180eccaa807ba4e2e2118bd040f8ddab8489332e4525c083fe"></a>

## role_id property — access_info.vault_auth_info.app_role_auth / 5ec5678374f5 / 4

Type: `"string"`. Computed.

Role ID. Role-ID to be used for authentication.

Upstream description:

Role-ID to be used for authentication.

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

- [secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-8ce6fbd2b6d65d0662abfee3d6e6541f76975430c54554562baf3898f0fc6c88): complete subsection reference.

<a id="canonical-f214537c920568dde38ffacc73c835bb048e6024a2dc746734eb56921c203871"></a>

## Next pages — access_info.vault_auth_info.app_role_auth / 5ec5678374f5 / 5

- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-8ce6fbd2b6d65d0662abfee3d6e6541f76975430c54554562baf3898f0fc6c88)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-8ce6fbd2b6d65d0662abfee3d6e6541f76975430c54554562baf3898f0fc6c88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ac39b4414c72b5986080a0296d5fdc0ad1404125a042b35b88c918cac71d097"></a>

## access_info.vault_auth_info.app_role_auth.secret_id — access_info.vault_auth_info.app_role_auth.secret_id / 45f6430b5a95 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-06c78926d465b5c09ae4a9137a0981c0b52426d1d96d985fce7e04c79f60070a)
- access_info.vault_auth_info.app_role_auth.secret_id

<a id="canonical-b17e59d4c3adb33a6b1945d8e48857b67153f16d0f90fed765d20635946ef3a4"></a>

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

<a id="canonical-4e94f59a01f0718ecd3d0b8a861f4830936c3963acbd47a969cc1db9e26fb0e1"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth.secret_id / 45f6430b5a95 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-7787e39d9c44b113fb9cd00b252fe4ac504c2aaca285bca733dabaa3487a773c): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-e477db4b9c8b0e26d8a2127c50dde697c089b80bc1d046c914aee8f803511e29): complete subsection reference.
