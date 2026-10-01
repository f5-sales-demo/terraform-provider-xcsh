---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28912d4b4722b9dcda778a569e427095cca21f34322a8c7caf4a2b411a0e6697"></a>

## Property reference — Property reference / 97ff103600cb / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- Property reference

<a id="canonical-17644dc144a89368692b14c7171a2c01e3404fff0ecc7aa2fb69c0a0eabdae8b"></a>

## Direct properties — Property reference / 97ff103600cb / 3

- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810): complete subsection reference.

<a id="canonical-2bde1a7647158898d46f83fdac72960b7cbb19cc5bd4a2e036d5c9b63979397b"></a>

<a id="canonical-4ce480c3b0bb40d99c63c6752ca532ea5a8a8f78b5a273b1177bfddf5d0c6c96"></a>

## annotations property — Property reference / 97ff103600cb / 4

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

<a id="canonical-bd42a00915d1892bcc3678a2bf425468d505cd2af10cecfdeede4fdcd0ad2348"></a>

<a id="canonical-7e5b2ccf03a2d00ceb416b2dde1ac7263df3a1220733bbc73defae45fe57bc3f"></a>

## description property — Property reference / 97ff103600cb / 5

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

<a id="canonical-7bb1570ce56022dfde45b20c60cde7e6a339a8a0b60a9c0cad8fa4cfa3347497"></a>

<a id="canonical-10983852bf1eda29b67d6843b8bcd820f1989539c4acb174893b93a59388ea76"></a>

## disable property — Property reference / 97ff103600cb / 6

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

<a id="canonical-4d87e6d9e54ad2bf0306203f26d3430410d34edaaaa968ae16149799bcfa626b"></a>

<a id="canonical-52e208e4fa67d917287f68a3d7fda94136966b162847fe6d4ba9f29596055491"></a>

## id property — Property reference / 97ff103600cb / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-281b2e3a325df63df4b03a922c29566b1734df0931f4d77e50f17a6098fe023d"></a>

<a id="canonical-0a6256500e703fbd706bf29c92606b6ae2d9c9933dae9179376c336ac315142b"></a>

## labels property — Property reference / 97ff103600cb / 8

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

<a id="canonical-aaed34e9206ac81683cc71c605f214dc47393d76e529d7c9b20f98f2bb1b5844"></a>

<a id="canonical-2572cac83f8fd828cbb7ef99205b86149d390c9f234be5bb0b18cf78f79da7f9"></a>

## name property — Property reference / 97ff103600cb / 9

Type: `"string"`. Required.

Name of the Secret Management Access. Must be unique within the namespace.

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

<a id="canonical-df3740add5dc38eae19e33c8a805645884daff8247c7e5b41e4a9c9f16d84008"></a>

<a id="canonical-d567e04afc7b7014063703ff7aaf932e1051897bb64e41573292ce7b12ef67c3"></a>

## namespace property — Property reference / 97ff103600cb / 10

Type: `"string"`. Required.

Namespace where the Secret Management Access is created.

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

<a id="canonical-450d360b22bac39cbb20f3bc0eaaca050381df6127b30f1e1c4af74e4d57e733"></a>

<a id="canonical-d512ffb5ec1c3b3ece75db0a4b1085333a5f8b3688596e9e169a52b6b91f643b"></a>

## provider_name property — Property reference / 97ff103600cb / 11

Type: `"string"`. Required.

Name given to this secret management backend. site.provider needs to be unique, and will be
referenced for using this object.

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

- [timeouts](resources--secret_management_access--reference--group-002.md#canonical-fb1d71681608bb290a6f1b56353da27906bf45ff58b32d0cf25c35930489c6f0): complete subsection reference.

- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5): complete subsection reference.

<a id="canonical-39d113bf5db6ce47759e47f5b2dd714a2185111d0b4bb85e9f893b8e7fe91862"></a>

## All schema paths — Property reference / 97ff103600cb / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `access_info` | [access_info](resources--secret_management_access--reference--group-001.md#canonical-6a386163f4ba8ab64d8b981f0adb49fc06c155f218c7a205541ccf1b07dd20a3) |
| `access_info.rest_auth_info` | [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-4d3fec29c58a6e835e039d0788717403e7f50c4b525f19c28c6e4ef6951f91fc) |
| `access_info.rest_auth_info.basic_auth` | [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-b65a0e7b5a06f0f6a1b0a6893e7889a9a149031bd2a4fd724cd580c5823c33aa) |
| `access_info.rest_auth_info.basic_auth.password` | [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-02d62886dc13ef1b772d5bc7ef54d7ee86faf45fdcd81855308d58b1a86aefc9) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-8f82f2b7b25989e6f2fc5a292a18b0a3b1ef54f82ab3468b11f2a14a657bfc45) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-001.md#canonical-7f23d842bcc093b98930be04e2861745c075288affcf0e0800a3e2c992dbc0df) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location](resources--secret_management_access--reference--group-001.md#canonical-edd5a1ad25b6a7437518befff59ff0aa4c8cac51b8314efa75cb91a74234b2da) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-001.md#canonical-0277cb424aecc160756430b967e5f9e6e03ca0040c411a8f4734b959e0b33a25) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-53c3480ad121422b58a1acac10789ac32944c9f1499f3885fabe9059339839c7) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-001.md#canonical-74054866fd7ba945d5765d9fc1db6177214c146f9be1cb27f42c9b4524e44d0b) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.url` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.url](resources--secret_management_access--reference--group-001.md#canonical-111e9d41689846b3e69bfcfd3bc05c5ef82620adaf64449a3c9142d12108cb52) |
| `access_info.rest_auth_info.basic_auth.username` | [access_info.rest_auth_info.basic_auth.username](resources--secret_management_access--reference--group-001.md#canonical-affccdce193a4004c611f1252bd973adfd76d88c91ce6e29edb64fc467d1df83) |
| `access_info.rest_auth_info.headers_auth` | [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-a8614693e0c9254e160eae872446907ea4a1d8fdb4a7daecb0cc8fb5f8657247) |
| `access_info.rest_auth_info.headers_auth.headers` | [access_info.rest_auth_info.headers_auth.headers](resources--secret_management_access--reference--group-001.md#canonical-4028bfdc77a8552e86d71babf15cccb3697d656b0544c17d339e68292a31d2b8) |
| `access_info.rest_auth_info.query_params_auth` | [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-ea223748a8617a0355131584af9f25a287f8c5e91dc29ed0af4b3bed8c87e612) |
| `access_info.rest_auth_info.query_params_auth.query_params` | [access_info.rest_auth_info.query_params_auth.query_params](resources--secret_management_access--reference--group-001.md#canonical-3cd8a9090d53fd0e0fe65f18bca2a7f7a779583429b72c6fb92311f1ec5fa6f2) |
| `access_info.scheme` | [access_info.scheme](resources--secret_management_access--reference--group-001.md#canonical-1659908f09ba394053ce3de5650c61e7a05f99ec3d33edf4fbd8144bc33633d2) |
| `access_info.server_endpoint` | [access_info.server_endpoint](resources--secret_management_access--reference--group-001.md#canonical-7cd0fe591f4d10c96d8a8bf13742e228cb2ece69b06979598174c15b015e76a6) |
| `access_info.tls_config` | [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-428b2c6b28bb1296cdb61db3a046d733f0662e24c6d960a895082cb585f58394) |
| `access_info.tls_config.cert_params` | [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-377302607e854b35f29fbafc8c2634cbc2ae62b289a39e0fc75649c43a70be33) |
| `access_info.tls_config.cert_params.certificates` | [access_info.tls_config.cert_params.certificates](resources--secret_management_access--reference--group-001.md#canonical-dec92ff55d4e043cb8496ec2bac6de8e2b1295bf1f09ea4dd5a33ee2affed711) |
| `access_info.tls_config.cert_params.certificates.kind` | [access_info.tls_config.cert_params.certificates.kind](resources--secret_management_access--reference--group-001.md#canonical-14fb8b18845b0da54e2fdd2716d9e8394371c2a9151ce111d400e43a39c4b577) |
| `access_info.tls_config.cert_params.certificates.name` | [access_info.tls_config.cert_params.certificates.name](resources--secret_management_access--reference--group-001.md#canonical-594b959d8daf9e1ce564f9aefa604d04ff12a17295ea84912d96c9736fe7ae4d) |
| `access_info.tls_config.cert_params.certificates.namespace` | [access_info.tls_config.cert_params.certificates.namespace](resources--secret_management_access--reference--group-001.md#canonical-16de07d482f4232aeff1637acce93c9ec623baf2dc905cd429df94200eb8cc05) |
| `access_info.tls_config.cert_params.certificates.tenant` | [access_info.tls_config.cert_params.certificates.tenant](resources--secret_management_access--reference--group-001.md#canonical-b726e7c7784cb6da6a18c60e5090f01551adf26d645f5ca448256f89d37b6fe4) |
| `access_info.tls_config.cert_params.certificates.uid` | [access_info.tls_config.cert_params.certificates.uid](resources--secret_management_access--reference--group-001.md#canonical-0da8906ee8db07ca070369899d09476ec366618e88f74bcb5559bb56d54821f6) |
| `access_info.tls_config.cert_params.cipher_suites` | [access_info.tls_config.cert_params.cipher_suites](resources--secret_management_access--reference--group-001.md#canonical-d32a7f716f46a80a7c0fc83cce45c2c2bb86f0a958b2a5fe9041113fa2d5208c) |
| `access_info.tls_config.cert_params.maximum_protocol_version` | [access_info.tls_config.cert_params.maximum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-65a64a67ddb216b92ca4fb48f9c388245f11434c97ec33327b353f008408df5f) |
| `access_info.tls_config.cert_params.minimum_protocol_version` | [access_info.tls_config.cert_params.minimum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-50da1e53e6ebc1b3eadd812d3ea50f2bf4370ae5e24b1e22bf9f73a758030d4d) |
| `access_info.tls_config.cert_params.skip_server_verification` | [access_info.tls_config.cert_params.skip_server_verification](resources--secret_management_access--reference--group-001.md#canonical-2ab8f23209c0d1d7e2de4d27f4624fc435387befc542d1c2fd1183dbc0b03b53) |
| `access_info.tls_config.cert_params.tls_validation_params` | [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-50a1f1f904f38843fe108668a17de2e6f0c399e646efc6d6e5eb0993d3fd62a7) |
| `access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification` | [access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification](resources--secret_management_access--reference--group-001.md#canonical-87e20cac3be9040de5b0e45796dcdc03e524696241388853b5ab566a6d6c07d5) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-44ceb21986223670a34cb222b9f53b4b77e3e94c5691fb7b08e3a81395e89740) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-650c1ac0beabe451b603d80a307941efcedd75f96329a768cce70e9d2da2173a) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](resources--secret_management_access--reference--group-001.md#canonical-35f011c3f6859ac20dce523ab8af420a47d7be200c788ac0cc640f2ffe1bc39c) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](resources--secret_management_access--reference--group-001.md#canonical-93bd711c6001e0a8a3551c97b650894e7b1a931c0271b27531cb9c6ee6117771) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](resources--secret_management_access--reference--group-001.md#canonical-d4ff2d671b39df16ee985546189b888f021df35e33da31df7c976070c51476b8) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](resources--secret_management_access--reference--group-001.md#canonical-ad6bafea2bc40847d31891473fa8a41bc37dbaae04db2a2468ccd72847264b79) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](resources--secret_management_access--reference--group-001.md#canonical-620938001305c73abbfcfc9e58ce4eb397126bfd1f570e904fedd3a5c730f57a) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url](resources--secret_management_access--reference--group-001.md#canonical-61902c5c4c10a987461fb1c9beff249669f9ccfd776e511b717a274536e4f530) |
| `access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names` | [access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names](resources--secret_management_access--reference--group-001.md#canonical-f1a59447b5fd9ebb50787190fd7854ee90b02e15ae7a85e4777b37f2e5c88a79) |
| `access_info.tls_config.cert_params.volterra_trusted_ca` | [access_info.tls_config.cert_params.volterra_trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-05adb79ad8d2f365365644d34b04523b1e954d723601af3ebbbad0027400cadb) |
| `access_info.tls_config.common_params` | [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-49dddf179dac50d8f56967ff6c2a5c67a5bd48bd058ab62d7cf401f03522934d) |
| `access_info.tls_config.common_params.cipher_suites` | [access_info.tls_config.common_params.cipher_suites](resources--secret_management_access--reference--group-001.md#canonical-0a0e216b74e3dea2adce8d039db332b8bad386add2e5dac2ac44c236bddc0211) |
| `access_info.tls_config.common_params.maximum_protocol_version` | [access_info.tls_config.common_params.maximum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-54540493f1e8b2c3adcc9691f8e5abbea3d6352956d7e3b4eae6c8f8dd4c9ecc) |
| `access_info.tls_config.common_params.minimum_protocol_version` | [access_info.tls_config.common_params.minimum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-285074ecda796b731bc49ffb9fe6d87d8838c2da35d29eab57c6959c028b6398) |
| `access_info.tls_config.common_params.tls_certificates` | [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-00e836629bc900bc614716e6424c39eeb96ab06050d70ecee38f07818cae8d8f) |
| `access_info.tls_config.common_params.tls_certificates.certificate_url` | [access_info.tls_config.common_params.tls_certificates.certificate_url](resources--secret_management_access--reference--group-001.md#canonical-c71b92c8f122ecf3320a80831c8ba259e9fb464f53cbc2e0e8d59264c946759a) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-474ebb1c837fbb93d3e8c8356f001b2d50dc105392aa8a7eec74480abee74f66) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-42e2bedeaad2807331ab5f23748d500a79bb67a383b6219f0f68e27b1dc02bc0) |
| `access_info.tls_config.common_params.tls_certificates.description_spec` | [access_info.tls_config.common_params.tls_certificates.description_spec](resources--secret_management_access--reference--group-001.md#canonical-82331335243df1069932ec10f2de15d4377872a6a7d9af225217a70f6c2acd57) |
| `access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling` | [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](resources--secret_management_access--reference--group-001.md#canonical-1a5f4b125204d3f2dd24409eff898b54a6a10824a350a3c18459cfd125e13832) |
| `access_info.tls_config.common_params.tls_certificates.private_key` | [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-f3d4ada5d1c348cfa949143382aa597fe46d90abb33123950b60425eaa893a6a) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-1e78b2bcadc0d93ebda6cc1754217fa07b22abdcc8e8bb4bcaec00e34dabaf25) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-001.md#canonical-269dab41634736e89d518ae07424b5bcf3e7a277f246a12cc5e760ae45ea63ec) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--secret_management_access--reference--group-001.md#canonical-76ea84f19295c5298f1f592ecf21622c59b4ca85d429628ea9a6484842b99802) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-001.md#canonical-e3317b6fc4a667ebbda0255acc1c723e8dcd0e1e794cc312cf4f2481d15a5fa9) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-1dbd4247a9733899a87902ddf100f4e9549eb7ae391c344763e31a280f1e0ac0) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-001.md#canonical-5c17f45603b3aeb62c45c8f4cfae180b01ba58cbc66559f28326fb4b00dc1774) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url](resources--secret_management_access--reference--group-001.md#canonical-6bf8bcdef495a879615ad2dfea094259155df7c90f2640e6744c5f1f07d50b72) |
| `access_info.tls_config.common_params.tls_certificates.use_system_defaults` | [access_info.tls_config.common_params.tls_certificates.use_system_defaults](resources--secret_management_access--reference--group-001.md#canonical-2e764bef99485ae4759ff5bcb6a37987113ae0bee1b8de714641dcab22ab96d3) |
| `access_info.tls_config.common_params.validation_params` | [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-c9c42d7f9ddfdc3e660d19b60b29244383e3686dbc01be67c7fc57b640307ab8) |
| `access_info.tls_config.common_params.validation_params.skip_hostname_verification` | [access_info.tls_config.common_params.validation_params.skip_hostname_verification](resources--secret_management_access--reference--group-001.md#canonical-64296554df5cb8164538a42f19ccbe1ed882768e15c3a38edf568c54ccacbdfa) |
| `access_info.tls_config.common_params.validation_params.trusted_ca` | [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2ee1fd55ed3759eb1b61760ef1936e05d4e3975a518125fa6bf27ef33dee02f6) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-9d5bf2ad50aa24c5b187e4335447058f7a0324a9859247083e2eb91308957cf0) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--secret_management_access--reference--group-001.md#canonical-43b3ff77cfd7798769e44ee7da36ee4ab5bf05b3298bbf864407dd5573188b95) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--secret_management_access--reference--group-001.md#canonical-e0c69bcc8cd99e905054eee655b2879388d0c8cbe511aa5271486435f54a6b93) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--secret_management_access--reference--group-001.md#canonical-0a8b5a064951670d206a67b693cb7ed20169bf23805ac4bc4e924e68100f189d) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--secret_management_access--reference--group-001.md#canonical-5bb03d03e187c576bb184cce0db5e3ad690ae10bacb0196cc8d10f5600f4001c) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--secret_management_access--reference--group-001.md#canonical-e6b2759ec771bd27539b1c98936876c3146ad193862e675b5eebc9cdb258f9c5) |
| `access_info.tls_config.common_params.validation_params.trusted_ca_url` | [access_info.tls_config.common_params.validation_params.trusted_ca_url](resources--secret_management_access--reference--group-001.md#canonical-2b19523f5414a7f3f8bf41f05b274b116da89a3a80523bf98d0fde4dc5bb7320) |
| `access_info.tls_config.common_params.validation_params.verify_subject_alt_names` | [access_info.tls_config.common_params.validation_params.verify_subject_alt_names](resources--secret_management_access--reference--group-001.md#canonical-d2d93d26429a521bc80d226bf92e65727fb7685f2b4dc61deca63948d3d8c045) |
| `access_info.tls_config.default_session_key_caching` | [access_info.tls_config.default_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-14125f8f0a1bcdc41349cac3899c0fc2920145a6b9a6fe286e50186c36304702) |
| `access_info.tls_config.disable_session_key_caching` | [access_info.tls_config.disable_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-4b465b869fa6f8a83f734d26dbd938e40a009c63c59ab9dc567d1248df58c89b) |
| `access_info.tls_config.disable_sni` | [access_info.tls_config.disable_sni](resources--secret_management_access--reference--group-001.md#canonical-e84590d96434e90bf5f058e49941df67347553b51686a8b2c08645d6b88093df) |
| `access_info.tls_config.max_session_keys` | [access_info.tls_config.max_session_keys](resources--secret_management_access--reference--group-001.md#canonical-82d136719e1f0b0006652978fefa37cb1b54826430ffb548728f12835a36cf2d) |
| `access_info.tls_config.sni` | [access_info.tls_config.sni](resources--secret_management_access--reference--group-001.md#canonical-5f49f00ca20abb181c9b871b8759f61eff7d7fedac85c152e9962e2ea1a8a6d6) |
| `access_info.tls_config.use_host_header_as_sni` | [access_info.tls_config.use_host_header_as_sni](resources--secret_management_access--reference--group-001.md#canonical-62459b1e52a190bcfa1a73fb2b3bfcf3f627de91eb438d8124c54bfd44a3704a) |
| `access_info.vault_auth_info` | [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-06b5c70b89f1ed821b425a12b8dd3d4e102f09fc18943e9ec8cae91976f8f748) |
| `access_info.vault_auth_info.app_role_auth` | [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-fb3ead27a67aacb3b0c34726b9acfc595c1f56efd8df8c65f47474ddb8f959bc) |
| `access_info.vault_auth_info.app_role_auth.role_id` | [access_info.vault_auth_info.app_role_auth.role_id](resources--secret_management_access--reference--group-002.md#canonical-85b6bcb752bc7104ceee9e7458358ec887a8f617f99c42880b7213af01c84e28) |
| `access_info.vault_auth_info.app_role_auth.secret_id` | [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--reference--group-002.md#canonical-e9c33f4f50d30f02575861665eae5ddd457063888e4b7a78f47b21840288e053) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-aacfe2c02732ef1123fa273595f3ae58e4672a6e34830287a9cdf8d5b053f020) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-002.md#canonical-538657ea8ae7497acb1a678931c0bfc69767046d5aeabd3fd1868e3dba4dc188) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location](resources--secret_management_access--reference--group-002.md#canonical-10ad5833b6b9f6d9153c2040872b362e24ff147f44d52e35654e0a721caef6af) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-002.md#canonical-c04c03f4b68928ea71ad06b1398cfb48dbb1b2287e4320bb3144be30e8ea76a1) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-f19bcb05cc33222b41cf1ee07f8440cfe69a51a14e0ebd6ebaa57bb2f9c70af4) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-002.md#canonical-8546c5faaaf62dd3b434b095352da960fab9e410e96b338578c4b66d8c2a721a) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url](resources--secret_management_access--reference--group-002.md#canonical-6805217a6507363df82fcb5ab041655aa4b8189500d0252f1d8ef6aef018b884) |
| `access_info.vault_auth_info.token` | [access_info.vault_auth_info.token](resources--secret_management_access--reference--group-002.md#canonical-d8e82bc326f4510bf3290e45954a4547faaaebb36a346c30d11e344867b5f48a) |
| `access_info.vault_auth_info.token.blindfold_secret_info` | [access_info.vault_auth_info.token.blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-0a3efc6d4d122e1fdc7ccae9683f73ca8e441e0c986d45bf0296490c657dd93a) |
| `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-002.md#canonical-a00536a4e7c9f02917f1c021cdfa04b5834da3391d7a9962958757c131496858) |
| `access_info.vault_auth_info.token.blindfold_secret_info.location` | [access_info.vault_auth_info.token.blindfold_secret_info.location](resources--secret_management_access--reference--group-002.md#canonical-82a124d87f4b6522b4847ef4ee70c8334c8a1ff74738019c15b07d29d5604244) |
| `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-002.md#canonical-8e362a17e552e114780c7295409684f0909a677645b9f5f55cdec4e6d0edf2d3) |
| `access_info.vault_auth_info.token.clear_secret_info` | [access_info.vault_auth_info.token.clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-eae2c87414c22374b521d9f6defee758cdcb8e31a008e9d3a51c35de3c4ca447) |
| `access_info.vault_auth_info.token.clear_secret_info.provider_ref` | [access_info.vault_auth_info.token.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-002.md#canonical-4d7ef84491e3ef2270afa80ba1371b2863a1184d168033a7353b076d1df93a03) |
| `access_info.vault_auth_info.token.clear_secret_info.url` | [access_info.vault_auth_info.token.clear_secret_info.url](resources--secret_management_access--reference--group-002.md#canonical-d2c4e1409d0add568c498d0889eeb248799a6e14c7f43eca4307264a77371d4a) |
| `annotations` | [annotations](resources--secret_management_access--reference--group-001.md#canonical-2bde1a7647158898d46f83fdac72960b7cbb19cc5bd4a2e036d5c9b63979397b) |
| `description` | [description](resources--secret_management_access--reference--group-001.md#canonical-bd42a00915d1892bcc3678a2bf425468d505cd2af10cecfdeede4fdcd0ad2348) |
| `disable` | [disable](resources--secret_management_access--reference--group-001.md#canonical-7bb1570ce56022dfde45b20c60cde7e6a339a8a0b60a9c0cad8fa4cfa3347497) |
| `id` | [id](resources--secret_management_access--reference--group-001.md#canonical-4d87e6d9e54ad2bf0306203f26d3430410d34edaaaa968ae16149799bcfa626b) |
| `labels` | [labels](resources--secret_management_access--reference--group-001.md#canonical-281b2e3a325df63df4b03a922c29566b1734df0931f4d77e50f17a6098fe023d) |
| `name` | [name](resources--secret_management_access--reference--group-001.md#canonical-aaed34e9206ac81683cc71c605f214dc47393d76e529d7c9b20f98f2bb1b5844) |
| `namespace` | [namespace](resources--secret_management_access--reference--group-001.md#canonical-df3740add5dc38eae19e33c8a805645884daff8247c7e5b41e4a9c9f16d84008) |
| `provider_name` | [provider_name](resources--secret_management_access--reference--group-001.md#canonical-450d360b22bac39cbb20f3bc0eaaca050381df6127b30f1e1c4af74e4d57e733) |
| `timeouts` | [timeouts](resources--secret_management_access--reference--group-002.md#canonical-ef69aa41b68a9ca1441273379d94edb35f6ca400b9dd4678bfab5efaa9272e0e) |
| `timeouts.create` | [timeouts.create](resources--secret_management_access--reference--group-002.md#canonical-97998da037a0860825a4d16d8398dba4b9a0f490073824ab4c96e0085c894a4e) |
| `timeouts.delete` | [timeouts.delete](resources--secret_management_access--reference--group-002.md#canonical-5b75b6826f6f6b4f250546b62c364cfd50a2d754c8c27b822d00e326eacdf755) |
| `timeouts.read` | [timeouts.read](resources--secret_management_access--reference--group-002.md#canonical-421d70ecc20ac1a050cf9da3b212ea80aa12e29cd8368f89daff3dbb932349ad) |
| `timeouts.update` | [timeouts.update](resources--secret_management_access--reference--group-002.md#canonical-4ed5ba7cfc689cb36676cc70132b19dd220d9f901c2b769aa1ce891d21b77f40) |
| `where` | [where](resources--secret_management_access--reference--group-002.md#canonical-aa22c1ca5e2ea26d49e03389a142c04d8cf47423488e4259b50904237a3821e9) |
| `where.site` | [where.site](resources--secret_management_access--reference--group-002.md#canonical-a91c31a37e1617dbb3b8c3f85dc683c5355e51e011d7f46635ff64e730708d70) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-65ad50e31a8a5d63611bb82eb2a3a4d3bd6fbe30b6b31e8316c159016486b008) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-90d4b9a452bacae9a408edc49d766f0692b955b6ff96feeaf40f778355ede91a) |
| `where.site.network_type` | [where.site.network_type](resources--secret_management_access--reference--group-002.md#canonical-9f1e7cb7b86518ae26b1fb8e2280d3752343d67f9bc7879a96c7b15654e53bd4) |
| `where.site.ref` | [where.site.ref](resources--secret_management_access--reference--group-002.md#canonical-781de395688e6ea87d58da1430e2d24d6314196af79947f05ab54fbb7e8e3eab) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--secret_management_access--reference--group-002.md#canonical-a5205831a5bb57a998304ef0346d9a987ee789eef80f2050b3b0731aa9d2ca86) |
| `where.site.ref.name` | [where.site.ref.name](resources--secret_management_access--reference--group-002.md#canonical-ba97ecbc4c4b3e5423302d54832a9cd3361bcd972626b3666002e298ba18cf04) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--secret_management_access--reference--group-002.md#canonical-456cb3738c69680cbae04767549c5f9ec460a78c531cd0e8d744c6dc47aa0d11) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--secret_management_access--reference--group-002.md#canonical-055c56cda9fbaf752a0d8c93b4675ec2719b9141f47a400790a74330dda19d28) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--secret_management_access--reference--group-002.md#canonical-ec508bed7c6581375fd763033992674d525d95ab12d0d3b79ed882aa61e0cdb6) |
| `where.virtual_network` | [where.virtual_network](resources--secret_management_access--reference--group-002.md#canonical-7cfeab1ec7b549a6591ff824c08cbdcb2e70c5ec7bbfdf328ffde63fed41d6a7) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--secret_management_access--reference--group-002.md#canonical-a1e04129c99788db43360efa4d0d0ee086aab7de20f5b6bee6204b93f2e254ee) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--secret_management_access--reference--group-002.md#canonical-37fdea1cb03d29401e0527e4a9b8a4292f7f635f7ac79eff0575aea492b37b27) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--secret_management_access--reference--group-002.md#canonical-e84aac6c4ef7d9e6ef9e0346b398b59b0db2a735687b63e0c055bacf90c12a58) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--secret_management_access--reference--group-002.md#canonical-8ed0f4e57cbdde7438800c364643bd467366d89223d0af7d17621d9c407afcf9) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--secret_management_access--reference--group-002.md#canonical-a7f9d0891a5f490cf7209ab9dd53710c9fb7cd2acdc2ff72e81ea7edf3620aea) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--secret_management_access--reference--group-002.md#canonical-560f10c7ddbdfe437cbe7832ca59f4bd791974a5659f1f0094152df2d83ae3ce) |
| `where.virtual_site` | [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-dc5c57afbd23d619fd55105573448b74aa46a5e6aab529b71fd3e4cd346c7419) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-0c43e402f2667359ac97fe6b56aa56da32f2f81372946a51f50a27871d3c94fa) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-77a7c5e3800f0688a8707a1ec1d219a517d8dd6725ed1a365e99a0dc1b0b38f3) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--secret_management_access--reference--group-002.md#canonical-49054d9f1ae56e265831169149b9619a2eb355024fd48d6e1b916871d06aeb6f) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--secret_management_access--reference--group-002.md#canonical-2ba82298eb47ed20fd17458f3d1a54689dd8d1944eb81efa4729fd3d3a1eb9e3) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--secret_management_access--reference--group-002.md#canonical-22e223960b110165d6cdadff9ad29405d4df000479d48942dbc036c692bd7b08) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--secret_management_access--reference--group-002.md#canonical-cb1efbad658971c542efeb7a52dfe4dd141a7d58ecc5d13e5f3ac559423886a4) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--secret_management_access--reference--group-002.md#canonical-ce099a80cac9ebe68a3f786648301f722775087f9c364abd5eea78dddc06c3a3) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--secret_management_access--reference--group-002.md#canonical-23c49f598afbca47a0a43aca08709aaeaa2ee07fbf35432606560eeb9a1b4155) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--secret_management_access--reference--group-002.md#canonical-6e3627b4f479d2a8d458e3f34f353aba0b07538ea3e78f8ff07f4ce7bdff5cb3) |

<a id="canonical-9ceec4a91e115d4202f9d3be18d97da1f0bece5a46656090d5a61f44f75e5950"></a>

## Next pages — Property reference / 97ff103600cb / 13

- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [timeouts](resources--secret_management_access--reference--group-002.md#canonical-fb1d71681608bb290a6f1b56353da27906bf45ff58b32d0cf25c35930489c6f0)
- [where](resources--secret_management_access--reference--group-002.md#canonical-c4087d0e67a27ee79be7055e4721f539676ce06f807dbea1aea339f3cc0006a5)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a647db15736c49212d17774511b470dcb504973b9d975db381a1b27fdfac9c00"></a>

## access_info — access_info / 40a30c226f54 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- access_info

<a id="canonical-6a386163f4ba8ab64d8b981f0adb49fc06c155f218c7a205541ccf1b07dd20a3"></a>

Type: `"object"`. single nested block, Optional.

HostAccessInfoType contains the information about how to connect to the remote host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_endpoint"),
  validators.ConflictingObjectAttributes("rest_auth_info",
    "vault_auth_info")}
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
  "x-ves-oneof-field-auth_params": "[\"rest_auth_info\",\"vault_auth_info\"]"
}
```

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf9c1efffc1174ee357b3834d19d94289b59b9182dbff47ef6c2688114c77a9d"></a>

## Direct properties — access_info / 40a30c226f54 / 3

- [rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4): complete subsection reference.

<a id="canonical-1659908f09ba394053ce3de5650c61e7a05f99ec3d33edf4fbd8144bc33633d2"></a>

<a id="canonical-28a56f1ca751e248813291ab1355af768978a026dcb68b7bd72559f8738c035b"></a>

## scheme property — access_info / 40a30c226f54 / 4

Type: `"string"`. Optional.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Upstream description:

SchemeType is used to indicate URL scheme

HTTP:// scheme HTTPS:// scheme.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HTTP",
    "HTTPS"),
}
```

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

<a id="canonical-7cd0fe591f4d10c96d8a8bf13742e228cb2ece69b06979598174c15b015e76a6"></a>

<a id="canonical-75e1a332ce9e4c381f29ec899dd8975508d98bfed58d5022a1c0eba002a81396"></a>

## server_endpoint property — access_info / 40a30c226f54 / 5

Type: `"string"`. Optional.

Endpoint to connect to, in host:port format.

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

- [tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99): complete subsection reference.

- [vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036): complete subsection reference.

<a id="canonical-9e4798420e3f5c4ee41558f56658c2f124e0d5146cda2a632b97b8c55ec727ec"></a>

## Next pages — access_info / 40a30c226f54 / 6

- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-001.md#canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41014e11f5a0516a3e7e7378581b3454d85d630f2c85e1089213e09c44f8bbd9"></a>

## access_info.rest_auth_info — access_info.rest_auth_info / 9be9ae285e95 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- access_info.rest_auth_info

<a id="canonical-4d3fec29c58a6e835e039d0788717403e7f50c4b525f19c28c6e4ef6951f91fc"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters for REST based hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("basic_auth",
    "headers_auth"),
  validators.ConflictingObjectAttributes("basic_auth",
    "query_params_auth"),
  validators.ConflictingObjectAttributes("headers_auth",
    "query_params_auth")}
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
  "x-ves-oneof-field-auth_params": "[\"basic_auth\",\"headers_auth\",\"query_params_auth\"]"
}
```

Terraform syntax:

```terraform
rest_auth_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-10582a8e79f2ade44cf738a0af24f3a7e41be85f68a95092155aa18ac2b2658c"></a>

## Direct properties — access_info.rest_auth_info / 9be9ae285e95 / 3

- [basic_auth](resources--secret_management_access--reference--group-001.md#canonical-01e78c62cf8438625a41afb5d2d0a36b91176c639e66a185b2cc0d51a6f8b764): complete subsection reference.

- [headers_auth](resources--secret_management_access--reference--group-001.md#canonical-2e621acbfde59af1897d75dd561c5d277d664286a424c26ed45a30cdd39e12f6): complete subsection reference.

- [query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-4e2f08c9527b8503b328d8bf4d40c193ac3538ef2a18daa556ff4b194cfa725c): complete subsection reference.

<a id="canonical-b99cde297e9b884a503f36270bee056c0dc6ef281b19a6d494b3dba74938516d"></a>

## Next pages — access_info.rest_auth_info / 9be9ae285e95 / 4

- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-01e78c62cf8438625a41afb5d2d0a36b91176c639e66a185b2cc0d51a6f8b764)
- [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-2e621acbfde59af1897d75dd561c5d277d664286a424c26ed45a30cdd39e12f6)
- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-4e2f08c9527b8503b328d8bf4d40c193ac3538ef2a18daa556ff4b194cfa725c)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-01e78c62cf8438625a41afb5d2d0a36b91176c639e66a185b2cc0d51a6f8b764"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b43b941dbb1e845534ce34769e19840b88832c3d55f1c69eb2460bffb396cb6c"></a>

## access_info.rest_auth_info.basic_auth — access_info.rest_auth_info.basic_auth / 2031834c9bce / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- access_info.rest_auth_info.basic_auth

<a id="canonical-b65a0e7b5a06f0f6a1b0a6893e7889a9a149031bd2a4fd724cd580c5823c33aa"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
basic_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad17b6d206c197ba501b59531ea7c020c890ef2028794dadc7d690a28bbd548b"></a>

## Direct properties — access_info.rest_auth_info.basic_auth / 2031834c9bce / 3

- [password](resources--secret_management_access--reference--group-001.md#canonical-06feac9b2faaad5e15461f6618420dd3713c10c09c9cbc7480e98b496f2ab461): complete subsection reference.

<a id="canonical-affccdce193a4004c611f1252bd973adfd76d88c91ce6e29edb64fc467d1df83"></a>

<a id="canonical-2b227e6e599afd42995d067b29a7cfd9872273cc59389a7061e330fb3dc39958"></a>

## username property — access_info.rest_auth_info.basic_auth / 2031834c9bce / 4

Type: `"string"`. Optional.

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

<a id="canonical-e33e0e04188314b9f38f67ca99c0009c5672299587ef5f5292fae3da2c4d1554"></a>

## Next pages — access_info.rest_auth_info.basic_auth / 2031834c9bce / 5

- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-06feac9b2faaad5e15461f6618420dd3713c10c09c9cbc7480e98b496f2ab461)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-06feac9b2faaad5e15461f6618420dd3713c10c09c9cbc7480e98b496f2ab461"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff35fac7a233b2e5001d20f6219f726d1e45181a8e7f70e47aa11fb674daddd2"></a>

## access_info.rest_auth_info.basic_auth.password — access_info.rest_auth_info.basic_auth.password / 264d47f787df / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-01e78c62cf8438625a41afb5d2d0a36b91176c639e66a185b2cc0d51a6f8b764)
- access_info.rest_auth_info.basic_auth.password

<a id="canonical-02d62886dc13ef1b772d5bc7ef54d7ee86faf45fdcd81855308d58b1a86aefc9"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-eeaa72a48e66d26ee631c946f5aea2722858dc4b9d631ca77964325a7bd87e55"></a>

## Direct properties — access_info.rest_auth_info.basic_auth.password / 264d47f787df / 3

- [blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-c52515eaf6d9c6735fc6743b9321ef59acb8894b58845aaefbea5f0a07ca3faf): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-187c1d8525cc621d2c412c709d8eea7022ce966d5f852ef5e542add6ea8466a5): complete subsection reference.

<a id="canonical-3add183aa32537cdc1594e31ba1b31366458948c2564e8f558773dd2b3d9262d"></a>

## Next pages — access_info.rest_auth_info.basic_auth.password / 264d47f787df / 4

- [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-c52515eaf6d9c6735fc6743b9321ef59acb8894b58845aaefbea5f0a07ca3faf)
- [access_info.rest_auth_info.basic_auth.password.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-187c1d8525cc621d2c412c709d8eea7022ce966d5f852ef5e542add6ea8466a5)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-01e78c62cf8438625a41afb5d2d0a36b91176c639e66a185b2cc0d51a6f8b764)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-c52515eaf6d9c6735fc6743b9321ef59acb8894b58845aaefbea5f0a07ca3faf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-881c7a0f4ec7851ff371cec99ed41863171e42c4557930a083689a994df1f18a"></a>

## access_info.rest_auth_info.basic_auth.password.blindfold_secret_info — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / fd99a713f7c4 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-01e78c62cf8438625a41afb5d2d0a36b91176c639e66a185b2cc0d51a6f8b764)
- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-06feac9b2faaad5e15461f6618420dd3713c10c09c9cbc7480e98b496f2ab461)
- access_info.rest_auth_info.basic_auth.password.blindfold_secret_info

<a id="canonical-8f82f2b7b25989e6f2fc5a292a18b0a3b1ef54f82ab3468b11f2a14a657bfc45"></a>

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

<a id="canonical-5522dcbd22a2f84898f995cbaf0bf62d8cefdf4d469614675a12e1fb8b6b005a"></a>

## Direct properties — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / fd99a713f7c4 / 3

<a id="canonical-7f23d842bcc093b98930be04e2861745c075288affcf0e0800a3e2c992dbc0df"></a>

<a id="canonical-b68ee810ca0b693e8fe0259ea59f6fd6610607a3ad5555338f81f8a7c53d9e6c"></a>

## decryption_provider property — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / fd99a713f7c4 / 4

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

<a id="canonical-edd5a1ad25b6a7437518befff59ff0aa4c8cac51b8314efa75cb91a74234b2da"></a>

<a id="canonical-750e2eed412a5be412844ce6c964b765dc39c5dcc653dc25ea7de135918cb602"></a>

## location property — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / fd99a713f7c4 / 5

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

<a id="canonical-0277cb424aecc160756430b967e5f9e6e03ca0040c411a8f4734b959e0b33a25"></a>

<a id="canonical-33f1282219a68849a3499d45e742eb508eab609d815aefe0a0cae3b3ebb11119"></a>

## store_provider property — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / fd99a713f7c4 / 6

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

<a id="canonical-9372e970f70ab8b3b3d00d20f419d8ced40d65a17850e049733e2638abed861b"></a>

## Next pages — access_info.rest_auth_info.basic_auth.password.blindfold_secret_info / fd99a713f7c4 / 7

- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-06feac9b2faaad5e15461f6618420dd3713c10c09c9cbc7480e98b496f2ab461)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-187c1d8525cc621d2c412c709d8eea7022ce966d5f852ef5e542add6ea8466a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8edcdffc7896474ddc8e7b24bcea2a4a681e55341938463572ff9e7ae092a2f5"></a>

## access_info.rest_auth_info.basic_auth.password.clear_secret_info — access_info.rest_auth_info.basic_auth.password.clear_secret_info / f844ceca1e2c / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-01e78c62cf8438625a41afb5d2d0a36b91176c639e66a185b2cc0d51a6f8b764)
- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-06feac9b2faaad5e15461f6618420dd3713c10c09c9cbc7480e98b496f2ab461)
- access_info.rest_auth_info.basic_auth.password.clear_secret_info

<a id="canonical-53c3480ad121422b58a1acac10789ac32944c9f1499f3885fabe9059339839c7"></a>

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

<a id="canonical-baa6ae370f1ae5c4d3b39ca5388969bf1bc33cf6671cefcfc52285e871161076"></a>

## Direct properties — access_info.rest_auth_info.basic_auth.password.clear_secret_info / f844ceca1e2c / 3

<a id="canonical-74054866fd7ba945d5765d9fc1db6177214c146f9be1cb27f42c9b4524e44d0b"></a>

<a id="canonical-d83f1a75156b001a0a1657d913257f41cc5909cd38725b3a384aa15d073fe31b"></a>

## provider_ref property — access_info.rest_auth_info.basic_auth.password.clear_secret_info / f844ceca1e2c / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-111e9d41689846b3e69bfcfd3bc05c5ef82620adaf64449a3c9142d12108cb52"></a>

<a id="canonical-9001589b08e48c6d0e62889228d74232ab6bd53fa69fc2bfdf628c9ec75b36f8"></a>

## url property — access_info.rest_auth_info.basic_auth.password.clear_secret_info / f844ceca1e2c / 5

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

<a id="canonical-8457fdab40a2d5347c1f2cd0c6f58c49fb41851f188598ebaf9d42c0d5888724"></a>

## Next pages — access_info.rest_auth_info.basic_auth.password.clear_secret_info / f844ceca1e2c / 6

- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-06feac9b2faaad5e15461f6618420dd3713c10c09c9cbc7480e98b496f2ab461)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-2e621acbfde59af1897d75dd561c5d277d664286a424c26ed45a30cdd39e12f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3d3a42abc3716b9b0631d56d4ac21cfbb8e7b4e09841f60b22c9f8e1fd046ca"></a>

## access_info.rest_auth_info.headers_auth — access_info.rest_auth_info.headers_auth / 8ccc20bfe31e / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- access_info.rest_auth_info.headers_auth

<a id="canonical-a8614693e0c9254e160eae872446907ea4a1d8fdb4a7daecb0cc8fb5f8657247"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
headers_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d58717eb20848885fd342589c1dd96a60d2d50978f810301757853556ed8c69"></a>

## Direct properties — access_info.rest_auth_info.headers_auth / 8ccc20bfe31e / 3

- [headers](resources--secret_management_access--reference--group-001.md#canonical-e666fa8785b39e61cea21b0d0a432f04c68dc4cf7fcb2ecaa6e4554a59ad854e): complete subsection reference.

<a id="canonical-be4e591441f39cac51d1142921bf751a5c56c3928d7d762abbd61d32fb3ace08"></a>

## Next pages — access_info.rest_auth_info.headers_auth / 8ccc20bfe31e / 4

- [access_info.rest_auth_info.headers_auth.headers](resources--secret_management_access--reference--group-001.md#canonical-e666fa8785b39e61cea21b0d0a432f04c68dc4cf7fcb2ecaa6e4554a59ad854e)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-e666fa8785b39e61cea21b0d0a432f04c68dc4cf7fcb2ecaa6e4554a59ad854e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ee0075c6813e0ae3e03040788c9e64844738bb1f0fed113608377c9a4582bdf"></a>

## access_info.rest_auth_info.headers_auth.headers — access_info.rest_auth_info.headers_auth.headers / 4e0875c4283d / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-2e621acbfde59af1897d75dd561c5d277d664286a424c26ed45a30cdd39e12f6)
- access_info.rest_auth_info.headers_auth.headers

<a id="canonical-4028bfdc77a8552e86d71babf15cccb3697d656b0544c17d339e68292a31d2b8"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
headers {}
```

<a id="canonical-d7c33392480b4a8234b17e9d7ebf0f48b0d2505c6ac124961a3bde573fcc70de"></a>

## Direct properties — access_info.rest_auth_info.headers_auth.headers / 4e0875c4283d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-511eadf6dd056cea3bbcb5e32d9325737f18433022eea85724ea76ebc9e0c558"></a>

## Next pages — access_info.rest_auth_info.headers_auth.headers / 4e0875c4283d / 4

- [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-2e621acbfde59af1897d75dd561c5d277d664286a424c26ed45a30cdd39e12f6)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-4e2f08c9527b8503b328d8bf4d40c193ac3538ef2a18daa556ff4b194cfa725c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7912bcf129691e587adc5ce0fbb25f210df14244404e8165b092244401ba5245"></a>

## access_info.rest_auth_info.query_params_auth — access_info.rest_auth_info.query_params_auth / a571c2abdaea / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- access_info.rest_auth_info.query_params_auth

<a id="canonical-ea223748a8617a0355131584af9f25a287f8c5e91dc29ed0af4b3bed8c87e612"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
query_params_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c4cb0c32a2d974deed9bed260589c59f4daa0c3afd2e6a2d975505e6b436dbf"></a>

## Direct properties — access_info.rest_auth_info.query_params_auth / a571c2abdaea / 3

- [query_params](resources--secret_management_access--reference--group-001.md#canonical-101a5c03969e60a0c5b390b2d2f250d3c87080d702e976aa8611bbdada24e800): complete subsection reference.

<a id="canonical-5c54fb18646781e81bdba446aae7abd20d94ced41461dca7c4fdf66cb3c565c8"></a>

## Next pages — access_info.rest_auth_info.query_params_auth / a571c2abdaea / 4

- [access_info.rest_auth_info.query_params_auth.query_params](resources--secret_management_access--reference--group-001.md#canonical-101a5c03969e60a0c5b390b2d2f250d3c87080d702e976aa8611bbdada24e800)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-101a5c03969e60a0c5b390b2d2f250d3c87080d702e976aa8611bbdada24e800"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f299c61351f83436669b7b7cff0513ce0e6504d17c58bb012922d349a37349a5"></a>

## access_info.rest_auth_info.query_params_auth.query_params — access_info.rest_auth_info.query_params_auth.query_params / e33ce31d47e4 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-f9094db902be9aaeddc83619008d0e588b91c27778101b64dbd80d4010d082b4)
- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-4e2f08c9527b8503b328d8bf4d40c193ac3538ef2a18daa556ff4b194cfa725c)
- access_info.rest_auth_info.query_params_auth.query_params

<a id="canonical-3cd8a9090d53fd0e0fe65f18bca2a7f7a779583429b72c6fb92311f1ec5fa6f2"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
query_params {}
```

<a id="canonical-a9948f0683c542cbe396ed6897bc078a16468b71859fe5f90173164409854675"></a>

## Direct properties — access_info.rest_auth_info.query_params_auth.query_params / e33ce31d47e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7aa0edca63f405d1790f5735484a4683da9467603e9900767593439253dd0f3"></a>

## Next pages — access_info.rest_auth_info.query_params_auth.query_params / e33ce31d47e4 / 4

- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-4e2f08c9527b8503b328d8bf4d40c193ac3538ef2a18daa556ff4b194cfa725c)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61502e2c8c6e49dc96c4bece1e32575eacdd6346b7fecc275c05c0025888895f"></a>

## access_info.tls_config — access_info.tls_config / 3a20d5206528 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- access_info.tls_config

<a id="canonical-428b2c6b28bb1296cdb61db3a046d733f0662e24c6d960a895082cb585f58394"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for upstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cert_params",
    "common_params"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-331985d068dc16be10128f7a1a538b53df79d6431099c842dbfb4cd2cee7f0da"></a>

## Direct properties — access_info.tls_config / 3a20d5206528 / 3

- [cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8): complete subsection reference.

- [common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18): complete subsection reference.

- [default_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-440b457ffd5761009c001add194635315bfd28ae6cb75045bd324238a6943c66): complete subsection reference.

- [disable_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-a30b8903d6e255c86d65f1356f61d1f7c1d6818eea1162139ac18d81d71145f6): complete subsection reference.

- [disable_sni](resources--secret_management_access--reference--group-001.md#canonical-3e601b97c72f8db5926f38ccc8081213cf665e0c33f6b9256dcf2f9ea8896626): complete subsection reference.

<a id="canonical-82d136719e1f0b0006652978fefa37cb1b54826430ffb548728f12835a36cf2d"></a>

<a id="canonical-747dd366d8fd37b20e1f294664cfc54a4c66b9d112d43464ae82bca46e83221e"></a>

## max_session_keys property — access_info.tls_config / 3a20d5206528 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

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

<a id="canonical-5f49f00ca20abb181c9b871b8759f61eff7d7fedac85c152e9962e2ea1a8a6d6"></a>

<a id="canonical-231c956ce01a35c2fff050c1a77329b96f4461b0381472b126dee37c85ed2723"></a>

## sni property — access_info.tls_config / 3a20d5206528 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [use_host_header_as_sni](resources--secret_management_access--reference--group-001.md#canonical-1d7b65d90256b43a6b7ecf321605c56e4aba082e85f7b94ac31daa293ea2d7af): complete subsection reference.

<a id="canonical-66269ff23ee0f1628d18aefe2e16befaa57e45441d578b29a7247d0df2157381"></a>

## Next pages — access_info.tls_config / 3a20d5206528 / 6

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.default_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-440b457ffd5761009c001add194635315bfd28ae6cb75045bd324238a6943c66)
- [access_info.tls_config.disable_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-a30b8903d6e255c86d65f1356f61d1f7c1d6818eea1162139ac18d81d71145f6)
- [access_info.tls_config.disable_sni](resources--secret_management_access--reference--group-001.md#canonical-3e601b97c72f8db5926f38ccc8081213cf665e0c33f6b9256dcf2f9ea8896626)
- [access_info.tls_config.use_host_header_as_sni](resources--secret_management_access--reference--group-001.md#canonical-1d7b65d90256b43a6b7ecf321605c56e4aba082e85f7b94ac31daa293ea2d7af)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b28cb1d6637840edcd1933dd5f631ac8b8fc1f7f62ef94a928c949ea47702ae"></a>

## access_info.tls_config.cert_params — access_info.tls_config.cert_params / 1fb1dfe77036 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- access_info.tls_config.cert_params

<a id="canonical-377302607e854b35f29fbafc8c2634cbc2ae62b289a39e0fc75649c43a70be33"></a>

Type: `"object"`. single nested block, Optional.

Certificate Parameters for authentication, TLS ciphers, and trust store.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "tls_validation_params"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("tls_validation_params",
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
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"tls_validation_params\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3273b21e1e84731b576037d1d8efc72676d7e4ee214fe1130df4e5459570ddac"></a>

## Direct properties — access_info.tls_config.cert_params / 1fb1dfe77036 / 3

- [certificates](resources--secret_management_access--reference--group-001.md#canonical-f6fb0f60fd52830dd3599c53ed72f1d7b0e0f258bde4feb574b2237465d47d5f): complete subsection reference.

<a id="canonical-d32a7f716f46a80a7c0fc83cce45c2c2bb86f0a958b2a5fe9041113fa2d5208c"></a>

<a id="canonical-f98984249f8fd86472c6902509dd6ea9374e877c64ae3faaaba30dfe06d65910"></a>

## cipher_suites property — access_info.tls_config.cert_params / 1fb1dfe77036 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-65a64a67ddb216b92ca4fb48f9c388245f11434c97ec33327b353f008408df5f"></a>

<a id="canonical-a167af0764701f55ad81f1f4c56508ceab0d3ae81ab84c6a677b1b6ca7495d24"></a>

## maximum_protocol_version property — access_info.tls_config.cert_params / 1fb1dfe77036 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-50da1e53e6ebc1b3eadd812d3ea50f2bf4370ae5e24b1e22bf9f73a758030d4d"></a>

<a id="canonical-dfb9a90ec316b4f27f01b32b439f1c24ef8b5ed150ea2e89a14cbed25186421c"></a>

## minimum_protocol_version property — access_info.tls_config.cert_params / 1fb1dfe77036 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

- [skip_server_verification](resources--secret_management_access--reference--group-001.md#canonical-77e7600f896295158f4a9e7ad268f8be08485d4b99ed59579cb34c09b373ca80): complete subsection reference.

- [tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-193785a55b9220aaaabb294c8c94f62da5413d1a54177cc56f75f8a79e6f0ea0): complete subsection reference.

- [volterra_trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-41b69f775bfb542ee5f49a8bf1fde08fcc19a9f4023f93c34ecd493c751b7669): complete subsection reference.

<a id="canonical-65ad44b82d72cec53d46ce040418547c6997d5aa639b9f722336c0e4406d7fc6"></a>

## Next pages — access_info.tls_config.cert_params / 1fb1dfe77036 / 7

- [access_info.tls_config.cert_params.certificates](resources--secret_management_access--reference--group-001.md#canonical-f6fb0f60fd52830dd3599c53ed72f1d7b0e0f258bde4feb574b2237465d47d5f)
- [access_info.tls_config.cert_params.skip_server_verification](resources--secret_management_access--reference--group-001.md#canonical-77e7600f896295158f4a9e7ad268f8be08485d4b99ed59579cb34c09b373ca80)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-193785a55b9220aaaabb294c8c94f62da5413d1a54177cc56f75f8a79e6f0ea0)
- [access_info.tls_config.cert_params.volterra_trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-41b69f775bfb542ee5f49a8bf1fde08fcc19a9f4023f93c34ecd493c751b7669)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-f6fb0f60fd52830dd3599c53ed72f1d7b0e0f258bde4feb574b2237465d47d5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06f6a5f01813a8567710d03680c0378e225660c5884c3ac89c6f8fa070527993"></a>

## access_info.tls_config.cert_params.certificates — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- access_info.tls_config.cert_params.certificates

<a id="canonical-dec92ff55d4e043cb8496ec2bac6de8e2b1295bf1f09ea4dd5a33ee2affed711"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3b4df77a7163a4814674db8c543b621426bb89ce286a9b209c9f57b56ae5155f"></a>

## Direct properties — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 3

<a id="canonical-14fb8b18845b0da54e2fdd2716d9e8394371c2a9151ce111d400e43a39c4b577"></a>

<a id="canonical-6c83b35a5c40c67e0c42c2787b5fa34be0986640af6e1f7b31e1c31d549540d6"></a>

## kind property — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 4

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

<a id="canonical-594b959d8daf9e1ce564f9aefa604d04ff12a17295ea84912d96c9736fe7ae4d"></a>

<a id="canonical-e6b99391fd59a912f71d436a3d68c4ce0e2b9e043d0e8e7750e446f83c425af2"></a>

## name property — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 5

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

<a id="canonical-16de07d482f4232aeff1637acce93c9ec623baf2dc905cd429df94200eb8cc05"></a>

<a id="canonical-fb18f59cc308c4da273bca39dd92933576029bbb5c7263f7e0f36d2d1eab0599"></a>

## namespace property — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 6

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

<a id="canonical-b726e7c7784cb6da6a18c60e5090f01551adf26d645f5ca448256f89d37b6fe4"></a>

<a id="canonical-e2fee36df80d7187373d92c9ab84bbcba7f4f383dbd744d4f636ec56e34c61c7"></a>

## tenant property — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 7

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

<a id="canonical-0da8906ee8db07ca070369899d09476ec366618e88f74bcb5559bb56d54821f6"></a>

<a id="canonical-ac200f9a0287fd40d03b92c6fda19fe48f50bb140e0fa628171b7c415f252b71"></a>

## uid property — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 8

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

<a id="canonical-c126fb8c51ca2ad48408c9b141d18be3a0399ce31c5c6fa8e10c050e1e486756"></a>

## Next pages — access_info.tls_config.cert_params.certificates / a08d56c4b179 / 9

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-77e7600f896295158f4a9e7ad268f8be08485d4b99ed59579cb34c09b373ca80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-590b31986cc3243dbe0e71b9f0152795c8910dba0de7c62f10a4df5837f92892"></a>

## access_info.tls_config.cert_params.skip_server_verification — access_info.tls_config.cert_params.skip_server_verification / c683e47fa8e0 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- access_info.tls_config.cert_params.skip_server_verification

<a id="canonical-2ab8f23209c0d1d7e2de4d27f4624fc435387befc542d1c2fd1183dbc0b03b53"></a>

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
skip_server_verification = {}
```

<a id="canonical-9b105d4ba28519dd883d5d6d4fb0a836b4020bca80c66e009d764ee949cc94ad"></a>

## Direct properties — access_info.tls_config.cert_params.skip_server_verification / c683e47fa8e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-23e343207f15750903a1b260699ff1f1bdbd6b3ac9ab97e06862a40dffacc964"></a>

## Next pages — access_info.tls_config.cert_params.skip_server_verification / c683e47fa8e0 / 4

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-193785a55b9220aaaabb294c8c94f62da5413d1a54177cc56f75f8a79e6f0ea0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0179386c53976a4a12726bb08e6a57a114ca037127ba37ff374af3c3076622b3"></a>

## access_info.tls_config.cert_params.tls_validation_params — access_info.tls_config.cert_params.tls_validation_params / 33e344b1eb1a / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- access_info.tls_config.cert_params.tls_validation_params

<a id="canonical-50a1f1f904f38843fe108668a17de2e6f0c399e646efc6d6e5eb0993d3fd62a7"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
tls_validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-97ecbef5e8d3ed2bc08d7b2323fdc956aa3be1ae43f5b64e8c139778c582d7d8"></a>

## Direct properties — access_info.tls_config.cert_params.tls_validation_params / 33e344b1eb1a / 3

<a id="canonical-87e20cac3be9040de5b0e45796dcdc03e524696241388853b5ab566a6d6c07d5"></a>

<a id="canonical-845cbd3a6f0537fef2b6e0a5279e16d9847004219f984fbd25d015344ddb5909"></a>

## skip_hostname_verification property — access_info.tls_config.cert_params.tls_validation_params / 33e344b1eb1a / 4

Type: `"bool"`. Optional.

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

- [trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8d327e2a0107f7d41cfae6e10a8c2a40a0da8dda131c61cc77c90c492f534079): complete subsection reference.

<a id="canonical-61902c5c4c10a987461fb1c9beff249669f9ccfd776e511b717a274536e4f530"></a>

<a id="canonical-d9cd90b88471d0ae9471a05d1f9949df0eedfab3158cb0d08a3f956abcd6d7ce"></a>

## trusted_ca_url property — access_info.tls_config.cert_params.tls_validation_params / 33e344b1eb1a / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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

<a id="canonical-f1a59447b5fd9ebb50787190fd7854ee90b02e15ae7a85e4777b37f2e5c88a79"></a>

<a id="canonical-4590ccd52aa4a3f4269a8457bddd9e20e8b3907d4493a84c5443f7cd305cee3e"></a>

## verify_subject_alt_names property — access_info.tls_config.cert_params.tls_validation_params / 33e344b1eb1a / 6

Type: `["list", "string"]`. Optional.

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

<a id="canonical-70584be3dd683f0610795117272808f29891322ccb6b8782892d0cc528061500"></a>

## Next pages — access_info.tls_config.cert_params.tls_validation_params / 33e344b1eb1a / 7

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8d327e2a0107f7d41cfae6e10a8c2a40a0da8dda131c61cc77c90c492f534079)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-8d327e2a0107f7d41cfae6e10a8c2a40a0da8dda131c61cc77c90c492f534079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ab7a668417c4aaad81b5f6879c7031f71e6f8e21a94acb6420c67da7513969e"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca — access_info.tls_config.cert_params.tls_validation_params.trusted_ca / 3570147d2c15 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-193785a55b9220aaaabb294c8c94f62da5413d1a54177cc56f75f8a79e6f0ea0)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca

<a id="canonical-44ceb21986223670a34cb222b9f53b4b77e3e94c5691fb7b08e3a81395e89740"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-9ce72880066bee3b6c1436ca50bde86ddc24f22e24c30cecb6d761d2f774823b"></a>

## Direct properties — access_info.tls_config.cert_params.tls_validation_params.trusted_ca / 3570147d2c15 / 3

- [trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-0087e9f5d9a52a84138eabe81097ab227b44571cfe7c5de7b5f8acb34c1e61e4): complete subsection reference.

<a id="canonical-6d81e835eef2506b7939d0d035b3248a974c581337b10521fe01acc13bafbb80"></a>

## Next pages — access_info.tls_config.cert_params.tls_validation_params.trusted_ca / 3570147d2c15 / 4

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-0087e9f5d9a52a84138eabe81097ab227b44571cfe7c5de7b5f8acb34c1e61e4)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-193785a55b9220aaaabb294c8c94f62da5413d1a54177cc56f75f8a79e6f0ea0)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-0087e9f5d9a52a84138eabe81097ab227b44571cfe7c5de7b5f8acb34c1e61e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90655b1817dbd5d2cf1370773a7e3cfd56c5d300fc183f72ae62b0ef2ab7dc9a"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-193785a55b9220aaaabb294c8c94f62da5413d1a54177cc56f75f8a79e6f0ea0)
- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8d327e2a0107f7d41cfae6e10a8c2a40a0da8dda131c61cc77c90c492f534079)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-650c1ac0beabe451b603d80a307941efcedd75f96329a768cce70e9d2da2173a"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-453e17c3b2ed3427a0bd4019ab7592e5b109fe1f530f3d4614de51b7394fa29a"></a>

## Direct properties — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 3

<a id="canonical-35f011c3f6859ac20dce523ab8af420a47d7be200c788ac0cc640f2ffe1bc39c"></a>

<a id="canonical-d6848221e9882925bba159a119d3c870f3b9d620a283932f14f2288b11dc1bfd"></a>

## kind property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 4

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

<a id="canonical-93bd711c6001e0a8a3551c97b650894e7b1a931c0271b27531cb9c6ee6117771"></a>

<a id="canonical-08fd44efe26d45eea226dc3022461c21c206db393a089e2ac92e86e77fce9983"></a>

## name property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 5

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

<a id="canonical-d4ff2d671b39df16ee985546189b888f021df35e33da31df7c976070c51476b8"></a>

<a id="canonical-13c496a52335c2d6177ddb08c6fc9bbb040951c815f93eddd8a5c0686b6e444b"></a>

## namespace property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 6

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

<a id="canonical-ad6bafea2bc40847d31891473fa8a41bc37dbaae04db2a2468ccd72847264b79"></a>

<a id="canonical-8a4069ae3f8a0bb830426b83878982d1af9a5042c147a2e236aa4a7d78dc517a"></a>

## tenant property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 7

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

<a id="canonical-620938001305c73abbfcfc9e58ce4eb397126bfd1f570e904fedd3a5c730f57a"></a>

<a id="canonical-a075be85e5f0bd012e0b2e82bff42e18d40bab2ddd876539cc462e2a15e15c0f"></a>

## uid property — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 8

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

<a id="canonical-162d6641ff73d8a0c5f248c963eb3616a9b301f47660591a2062fdc7187e5eca"></a>

## Next pages — access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_l / 60ad9e513c64 / 9

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8d327e2a0107f7d41cfae6e10a8c2a40a0da8dda131c61cc77c90c492f534079)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-41b69f775bfb542ee5f49a8bf1fde08fcc19a9f4023f93c34ecd493c751b7669"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c34c11fbf25ebf1ccaed3abe3f79d26544fd46e3bd5ff6f640d9f9c8a90c9978"></a>

## access_info.tls_config.cert_params.volterra_trusted_ca — access_info.tls_config.cert_params.volterra_trusted_ca / 059053aa0fe1 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- access_info.tls_config.cert_params.volterra_trusted_ca

<a id="canonical-05adb79ad8d2f365365644d34b04523b1e954d723601af3ebbbad0027400cadb"></a>

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

<a id="canonical-744ca9afffae927666cf1ac026dba4ea5226458e2478f6936ed5f474939faea0"></a>

## Direct properties — access_info.tls_config.cert_params.volterra_trusted_ca / 059053aa0fe1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0fa31da729f9b4a893240d3e665c214ecac2997d76ece56682b679300eca70e8"></a>

## Next pages — access_info.tls_config.cert_params.volterra_trusted_ca / 059053aa0fe1 / 4

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-c1280095a15af9806685342f12d12a7df15f25717d6575cff5b174a92563f1a8)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d9083cb340ae0b54e8782c1e22d28ce9406aae06ab49c1791e02b20c8556590"></a>

## access_info.tls_config.common_params — access_info.tls_config.common_params / 381865855112 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- access_info.tls_config.common_params

<a id="canonical-49dddf179dac50d8f56967ff6c2a5c67a5bd48bd058ab62d7cf401f03522934d"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
common_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-7072959b838e1550a54efa7c48dfb80a49d6b03674438d348ff22a7e0caf61ba"></a>

## Direct properties — access_info.tls_config.common_params / 381865855112 / 3

<a id="canonical-0a0e216b74e3dea2adce8d039db332b8bad386add2e5dac2ac44c236bddc0211"></a>

<a id="canonical-dfb9aeb264d62a16751219c1570fb3c2e2b56a81ac2b2b5ea8495a3eca20b351"></a>

## cipher_suites property — access_info.tls_config.common_params / 381865855112 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-54540493f1e8b2c3adcc9691f8e5abbea3d6352956d7e3b4eae6c8f8dd4c9ecc"></a>

<a id="canonical-353c0b31230e3a4f4f8db9bba49ba87c9af48677e12ab9d5e285d715274ffb6c"></a>

## maximum_protocol_version property — access_info.tls_config.common_params / 381865855112 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-285074ecda796b731bc49ffb9fe6d87d8838c2da35d29eab57c6959c028b6398"></a>

<a id="canonical-19694608db6d90a0a40e5c558e85d35d50963a74b8ca2f5cba82419ba1181399"></a>

## minimum_protocol_version property — access_info.tls_config.common_params / 381865855112 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

- [tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4): complete subsection reference.

- [validation_params](resources--secret_management_access--reference--group-001.md#canonical-902b4d12d094f784797e44fa83b40207d338b61676af569af6f8f6819768173f): complete subsection reference.

<a id="canonical-b4c15300c92849eba1092d51d8203171869d1f9fad6e201b5282bc4e6d61bfda"></a>

## Next pages — access_info.tls_config.common_params / 381865855112 / 7

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-902b4d12d094f784797e44fa83b40207d338b61676af569af6f8f6819768173f)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ad7f468ffab60eebe1e48b36584b556ca37ee7475fdc047a59bb4005745312e"></a>

## access_info.tls_config.common_params.tls_certificates — access_info.tls_config.common_params.tls_certificates / 334f61406f18 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- access_info.tls_config.common_params.tls_certificates

<a id="canonical-00e836629bc900bc614716e6424c39eeb96ab06050d70ecee38f07818cae8d8f"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
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
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-8c7fed74611224921266d251bd79cad3f541f4893e36bc31f65184b0e0658fae"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates / 334f61406f18 / 3

<a id="canonical-c71b92c8f122ecf3320a80831c8ba259e9fb464f53cbc2e0e8d59264c946759a"></a>

<a id="canonical-986fcca2f07b392e330fa4e933565f2d0d0f53f47f1465f4152a1aaa963a1dae"></a>

## certificate_url property — access_info.tls_config.common_params.tls_certificates / 334f61406f18 / 4

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

- [custom_hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-f1a40eaa9d38a3815cdd266426e00835e670496621ab9513cebe211422db5cb4): complete subsection reference.

<a id="canonical-82331335243df1069932ec10f2de15d4377872a6a7d9af225217a70f6c2acd57"></a>

<a id="canonical-a0d0155f9785a1b89d4fb400b57a67f8d7c83277be11216b8382baadaa260a4e"></a>

## description_spec property — access_info.tls_config.common_params.tls_certificates / 334f61406f18 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--secret_management_access--reference--group-001.md#canonical-a5fb19a271c7067788f825b97a777fb37cde544d8305f320c82dce2689763d5e): complete subsection reference.

- [private_key](resources--secret_management_access--reference--group-001.md#canonical-14771f6498f980105348197c89a4e648cb626303e87d76a92b7653298e6e54c7): complete subsection reference.

- [use_system_defaults](resources--secret_management_access--reference--group-001.md#canonical-beaf92f6d28b88523edebf22906fb8c6e3ac4f65c4f2d47c1d697e965dc076ed): complete subsection reference.

<a id="canonical-6760421b39a0889b68b0bb7031dd6996e6e6a68fb62d270ee43b25d810b538b9"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates / 334f61406f18 / 6

- [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-f1a40eaa9d38a3815cdd266426e00835e670496621ab9513cebe211422db5cb4)
- [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](resources--secret_management_access--reference--group-001.md#canonical-a5fb19a271c7067788f825b97a777fb37cde544d8305f320c82dce2689763d5e)
- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-14771f6498f980105348197c89a4e648cb626303e87d76a92b7653298e6e54c7)
- [access_info.tls_config.common_params.tls_certificates.use_system_defaults](resources--secret_management_access--reference--group-001.md#canonical-beaf92f6d28b88523edebf22906fb8c6e3ac4f65c4f2d47c1d697e965dc076ed)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-f1a40eaa9d38a3815cdd266426e00835e670496621ab9513cebe211422db5cb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89ac7b3fa5383b5fe6e07d84df3b2741e9ff94542c184db512aecd2f6c452e0f"></a>

## access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / dbdc2f72419a / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-474ebb1c837fbb93d3e8c8356f001b2d50dc105392aa8a7eec74480abee74f66"></a>

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

<a id="canonical-16d381ee310a2f2f04d2b217a72b9d08d28f06c2dafc806b2633419f28a8a697"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / dbdc2f72419a / 3

<a id="canonical-42e2bedeaad2807331ab5f23748d500a79bb67a383b6219f0f68e27b1dc02bc0"></a>

<a id="canonical-e3728310abea68dd0b3f90bae6d5706da7d870f39a850e8fe96a58c32c294a30"></a>

## hash_algorithms property — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / dbdc2f72419a / 4

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

<a id="canonical-89aab0359a8312f5021df290f3701749944197d33383f5be0c8a0f7b08aa9f23"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms / dbdc2f72419a / 5

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-a5fb19a271c7067788f825b97a777fb37cde544d8305f320c82dce2689763d5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18a29b7a9029ffa7c0a536d8ea74d4d844076894ff342460d1cfae1900bba62b"></a>

## access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling — access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling / da610494c681 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-1a5f4b125204d3f2dd24409eff898b54a6a10824a350a3c18459cfd125e13832"></a>

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

<a id="canonical-47cfa6570b916ec56f47f1958c921d216579fcb53912d4eb07f0e7b3235e1c1c"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling / da610494c681 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9cd4cbd1f3234d2d491628795514caf40c2ba86f265ffffc38fc300152bb318b"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling / da610494c681 / 4

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-14771f6498f980105348197c89a4e648cb626303e87d76a92b7653298e6e54c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-227b46d9546ecbb071f6410049c7c772d1c5a7b6dde73843539eac1fb38ab9ba"></a>

## access_info.tls_config.common_params.tls_certificates.private_key — access_info.tls_config.common_params.tls_certificates.private_key / 6174beb85bae / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- access_info.tls_config.common_params.tls_certificates.private_key

<a id="canonical-f3d4ada5d1c348cfa949143382aa597fe46d90abb33123950b60425eaa893a6a"></a>

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

<a id="canonical-da9417f9db50da774205b1ce29fa778f2ac69518a6bf95a905b773c1a8b74c7f"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.private_key / 6174beb85bae / 3

- [blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-a1ac536244e31ef040f4b3cb04090bca3ccc15b730bcd272ec3b597874533964): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-24670bb96ef1fa58d63cd9863c2a942eea09d524181ad14f55b553c6dd5bea7d): complete subsection reference.

<a id="canonical-9cf1f24ef90c221c3b5b8fac43a2ff4e52f9c401e4edc0527a957c1715f22a1d"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.private_key / 6174beb85bae / 4

- [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-a1ac536244e31ef040f4b3cb04090bca3ccc15b730bcd272ec3b597874533964)
- [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-24670bb96ef1fa58d63cd9863c2a942eea09d524181ad14f55b553c6dd5bea7d)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-a1ac536244e31ef040f4b3cb04090bca3ccc15b730bcd272ec3b597874533964"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da022ac30b659ee2be074af78a10ed14084ce0c2579a6b543e4c8bc2594302bd"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / 07732e52e34e / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-14771f6498f980105348197c89a4e648cb626303e87d76a92b7653298e6e54c7)
- access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1e78b2bcadc0d93ebda6cc1754217fa07b22abdcc8e8bb4bcaec00e34dabaf25"></a>

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

<a id="canonical-4802b07a1b89cfdd1f546ba6276fea13e777a4351db844f3ffcc4900b13f5355"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / 07732e52e34e / 3

<a id="canonical-269dab41634736e89d518ae07424b5bcf3e7a277f246a12cc5e760ae45ea63ec"></a>

<a id="canonical-9152020ffbf69491efa29724d2d175483e5009bd623433ef1c47e6bcd3914566"></a>

## decryption_provider property — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / 07732e52e34e / 4

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

<a id="canonical-76ea84f19295c5298f1f592ecf21622c59b4ca85d429628ea9a6484842b99802"></a>

<a id="canonical-7f66fd5d65832f70b1ca1c1a6d0b1ead25c1d15bfd8c77ed503304bcfbb1c9e6"></a>

## location property — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / 07732e52e34e / 5

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

<a id="canonical-e3317b6fc4a667ebbda0255acc1c723e8dcd0e1e794cc312cf4f2481d15a5fa9"></a>

<a id="canonical-822b6d65cc4f985f11bbc983b5912fda293fe7844d5a24525e63e2b5990a87fd"></a>

## store_provider property — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / 07732e52e34e / 6

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

<a id="canonical-8e76f5cf09574cc67a1caacd6c3a9bba2ea19af1c21e7ac0cdcd571cee704b1e"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secr / 07732e52e34e / 7

- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-14771f6498f980105348197c89a4e648cb626303e87d76a92b7653298e6e54c7)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-24670bb96ef1fa58d63cd9863c2a942eea09d524181ad14f55b553c6dd5bea7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-558cb221d815f51e4231ca662e2931aa54d4d5c0b3da2b5951bea0d467aa1dad"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / f18b648af350 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-14771f6498f980105348197c89a4e648cb626303e87d76a92b7653298e6e54c7)
- access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-1dbd4247a9733899a87902ddf100f4e9549eb7ae391c344763e31a280f1e0ac0"></a>

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

<a id="canonical-30a268c560b8ab1a166c648d908c255cefb25d3db5ba48f458bb931933f06d9a"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / f18b648af350 / 3

<a id="canonical-5c17f45603b3aeb62c45c8f4cfae180b01ba58cbc66559f28326fb4b00dc1774"></a>

<a id="canonical-64a01baa708b30b1ff4fcf8903e9b18ee13596035ff96df587d6c6a3699198e5"></a>

## provider_ref property — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / f18b648af350 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6bf8bcdef495a879615ad2dfea094259155df7c90f2640e6744c5f1f07d50b72"></a>

<a id="canonical-55df15b96295edcd9c7fccceb0206ca1d4c287ba9b1d4f6070e25b975375e903"></a>

## url property — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / f18b648af350 / 5

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

<a id="canonical-cb99eb6c114a4fe4a11d342b240ab45df6327d2322047eb238acab81439cf9de"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_i / f18b648af350 / 6

- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-14771f6498f980105348197c89a4e648cb626303e87d76a92b7653298e6e54c7)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-beaf92f6d28b88523edebf22906fb8c6e3ac4f65c4f2d47c1d697e965dc076ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e93818ed12dcac32050df56ac172d21d53e312272d3683dd29d36ccb23d8f18"></a>

## access_info.tls_config.common_params.tls_certificates.use_system_defaults — access_info.tls_config.common_params.tls_certificates.use_system_defaults / eca05a7dfea3 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- access_info.tls_config.common_params.tls_certificates.use_system_defaults

<a id="canonical-2e764bef99485ae4759ff5bcb6a37987113ae0bee1b8de714641dcab22ab96d3"></a>

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

<a id="canonical-fadde75a05957470abfea88b62a48d078f1c46169857189bce23c38ef1fa1227"></a>

## Direct properties — access_info.tls_config.common_params.tls_certificates.use_system_defaults / eca05a7dfea3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-797a5558f1eea3b6441f334637376d645289f98097175f5ca6211127c5b15344"></a>

## Next pages — access_info.tls_config.common_params.tls_certificates.use_system_defaults / eca05a7dfea3 / 4

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-6ab0dc5943f4ae964a80baacbf603b4c6db33281fa8d50d9ef645ec7ea7b34b4)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-902b4d12d094f784797e44fa83b40207d338b61676af569af6f8f6819768173f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bb57822976a2cb08b9b8521a2a629b8fea6a848ca23e3d770d05a37bfcf06d9"></a>

## access_info.tls_config.common_params.validation_params — access_info.tls_config.common_params.validation_params / 232c8d6f8123 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- access_info.tls_config.common_params.validation_params

<a id="canonical-c9c42d7f9ddfdc3e660d19b60b29244383e3686dbc01be67c7fc57b640307ab8"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-ddb35fa3004b2ca6350ef56f1ee76e7a07776b8f10e3cea326a3690d434784f6"></a>

## Direct properties — access_info.tls_config.common_params.validation_params / 232c8d6f8123 / 3

<a id="canonical-64296554df5cb8164538a42f19ccbe1ed882768e15c3a38edf568c54ccacbdfa"></a>

<a id="canonical-85bc1e3e52636fdd1a7a22156cc54c6f0464b48e1bd8d083a0878b4a1730d58f"></a>

## skip_hostname_verification property — access_info.tls_config.common_params.validation_params / 232c8d6f8123 / 4

Type: `"bool"`. Optional.

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

- [trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8632b568854e9052a284d5cf1af91dc3c132cb42cba1191c7255cd5c5a6f511d): complete subsection reference.

<a id="canonical-2b19523f5414a7f3f8bf41f05b274b116da89a3a80523bf98d0fde4dc5bb7320"></a>

<a id="canonical-328a22cb0eb9f24ddb3b3ff332e1353c5637b1d404194f3c743078993743d265"></a>

## trusted_ca_url property — access_info.tls_config.common_params.validation_params / 232c8d6f8123 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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

<a id="canonical-d2d93d26429a521bc80d226bf92e65727fb7685f2b4dc61deca63948d3d8c045"></a>

<a id="canonical-8499e223b9f90d369c351feead35e5e2af652f0da9225b886bec8005c37572db"></a>

## verify_subject_alt_names property — access_info.tls_config.common_params.validation_params / 232c8d6f8123 / 6

Type: `["list", "string"]`. Optional.

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

<a id="canonical-412fdcec4edf3bcb68109e7b25a492657be71e0207a41cff2811fd1064365f81"></a>

## Next pages — access_info.tls_config.common_params.validation_params / 232c8d6f8123 / 7

- [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8632b568854e9052a284d5cf1af91dc3c132cb42cba1191c7255cd5c5a6f511d)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-8632b568854e9052a284d5cf1af91dc3c132cb42cba1191c7255cd5c5a6f511d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15d04b4614df1ad67db377a0039f11eff272381ef27f6c843ccf0669c4445789"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca — access_info.tls_config.common_params.validation_params.trusted_ca / 81dea0b834fd / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-902b4d12d094f784797e44fa83b40207d338b61676af569af6f8f6819768173f)
- access_info.tls_config.common_params.validation_params.trusted_ca

<a id="canonical-2ee1fd55ed3759eb1b61760ef1936e05d4e3975a518125fa6bf27ef33dee02f6"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d528794aaef5ccb6e72c733690f50cad8f88fabf4dfd88e97c5a42fef73fc43"></a>

## Direct properties — access_info.tls_config.common_params.validation_params.trusted_ca / 81dea0b834fd / 3

- [trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-a237a5e9e5dfd70adc07550db1eb9d8459e57c81e3499a3f5235abde8258f1c7): complete subsection reference.

<a id="canonical-379c2421ac4698bff0e8cc76cdbe8a998353954c38b77a7008bc2427e826104e"></a>

## Next pages — access_info.tls_config.common_params.validation_params.trusted_ca / 81dea0b834fd / 4

- [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-a237a5e9e5dfd70adc07550db1eb9d8459e57c81e3499a3f5235abde8258f1c7)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-902b4d12d094f784797e44fa83b40207d338b61676af569af6f8f6819768173f)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-a237a5e9e5dfd70adc07550db1eb9d8459e57c81e3499a3f5235abde8258f1c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7a2ba71c4644d14d866e6099934ddb91d035670a4d80ea279508aa7e6b569da"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-6ea5052328e6c22c643654628db6133cb0b84af6331fc2bee14e884db0e9ac18)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-902b4d12d094f784797e44fa83b40207d338b61676af569af6f8f6819768173f)
- [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8632b568854e9052a284d5cf1af91dc3c132cb42cba1191c7255cd5c5a6f511d)
- access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-9d5bf2ad50aa24c5b187e4335447058f7a0324a9859247083e2eb91308957cf0"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-356a4e28a747e427951381788d05a734890ba32a53a827a44bfdc4972a44dbe6"></a>

## Direct properties — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 3

<a id="canonical-43b3ff77cfd7798769e44ee7da36ee4ab5bf05b3298bbf864407dd5573188b95"></a>

<a id="canonical-b2979e7b10f751f893310cabae93f34325c0926de7aabf124200ec1e7ba7ec4b"></a>

## kind property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 4

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

<a id="canonical-e0c69bcc8cd99e905054eee655b2879388d0c8cbe511aa5271486435f54a6b93"></a>

<a id="canonical-593ae1c2b6e472f56171e5e5abf981110680918119ae223af7285afd18a1ec52"></a>

## name property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 5

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

<a id="canonical-0a8b5a064951670d206a67b693cb7ed20169bf23805ac4bc4e924e68100f189d"></a>

<a id="canonical-27d35b448eed6e7cf81a8b6f66cc3c6ef9b2c904dbc452b90d93eae153f1305a"></a>

## namespace property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 6

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

<a id="canonical-5bb03d03e187c576bb184cce0db5e3ad690ae10bacb0196cc8d10f5600f4001c"></a>

<a id="canonical-0097e986b537ff19f8e75150fa4f6a9365b49bf207897f57adb5baed0e714bcf"></a>

## tenant property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 7

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

<a id="canonical-e6b2759ec771bd27539b1c98936876c3146ad193862e675b5eebc9cdb258f9c5"></a>

<a id="canonical-553f8de4cde93041b91403111153593080fb59569caf41658879844611ff004b"></a>

## uid property — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 8

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

<a id="canonical-4d201f5c0aead78756025c52379c5010719f0e1028956e5546ddaea63a8a5c18"></a>

## Next pages — access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_lis / c96459c3d8c6 / 9

- [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-8632b568854e9052a284d5cf1af91dc3c132cb42cba1191c7255cd5c5a6f511d)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-440b457ffd5761009c001add194635315bfd28ae6cb75045bd324238a6943c66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d89df047cab13bf1c1bbaa6fe7915bb4b45bc3e5975271b88fa8323319e54d64"></a>

## access_info.tls_config.default_session_key_caching — access_info.tls_config.default_session_key_caching / 4af7d14f1bfc / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- access_info.tls_config.default_session_key_caching

<a id="canonical-14125f8f0a1bcdc41349cac3899c0fc2920145a6b9a6fe286e50186c36304702"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_session_key_caching = {}
```

<a id="canonical-216beb10175f26edf498b049f169348608ed8229d8864cefab27532631345f20"></a>

## Direct properties — access_info.tls_config.default_session_key_caching / 4af7d14f1bfc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5e7f1d36619172fbb34a3c5266fd2cfbc8aa2654a4e2c44fe3a08d92600a9351"></a>

## Next pages — access_info.tls_config.default_session_key_caching / 4af7d14f1bfc / 4

- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-a30b8903d6e255c86d65f1356f61d1f7c1d6818eea1162139ac18d81d71145f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afebdbd4eaa4c6de3a896f0a97063b6f74b73904985a9f4936a395d210405744"></a>

## access_info.tls_config.disable_session_key_caching — access_info.tls_config.disable_session_key_caching / cc1be34542d6 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- access_info.tls_config.disable_session_key_caching

<a id="canonical-4b465b869fa6f8a83f734d26dbd938e40a009c63c59ab9dc567d1248df58c89b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_session_key_caching = {}
```

<a id="canonical-a0ef254ca7025ab541984830c731042ae8c124e51d464310797439473d2c10b2"></a>

## Direct properties — access_info.tls_config.disable_session_key_caching / cc1be34542d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5dad36c26f1257c245c2cb6f7d79076f75ccc9a138cc9fd7e9f52b137bf59f25"></a>

## Next pages — access_info.tls_config.disable_session_key_caching / cc1be34542d6 / 4

- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-3e601b97c72f8db5926f38ccc8081213cf665e0c33f6b9256dcf2f9ea8896626"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5035826ff9d24b8486c2809ebf7445aa10c882eb69420ba5a76790adc5c3f69"></a>

## access_info.tls_config.disable_sni — access_info.tls_config.disable_sni / 20e87817858f / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- access_info.tls_config.disable_sni

<a id="canonical-e84590d96434e90bf5f058e49941df67347553b51686a8b2c08645d6b88093df"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_sni = {}
```

<a id="canonical-6151d16f8d4fab4de40f4efdd7ef9414c3a1e6714302325c2320fbb2eb075f22"></a>

## Direct properties — access_info.tls_config.disable_sni / 20e87817858f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-587fc13ac0299288840f2dd2cbee8ac38f327c8ebbd14b571008e49e9fae1174"></a>

## Next pages — access_info.tls_config.disable_sni / 20e87817858f / 4

- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-1d7b65d90256b43a6b7ecf321605c56e4aba082e85f7b94ac31daa293ea2d7af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4656669992eec42191249b9070dd46c0231e0f19db715f0feba49dc0226b562"></a>

## access_info.tls_config.use_host_header_as_sni — access_info.tls_config.use_host_header_as_sni / 0285c1af8142 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- access_info.tls_config.use_host_header_as_sni

<a id="canonical-62459b1e52a190bcfa1a73fb2b3bfcf3f627de91eb438d8124c54bfd44a3704a"></a>

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
use_host_header_as_sni = {}
```

<a id="canonical-5f4bbae7c5607b49d3c3d5eb8a13a91670b6d4d11c6f43346f29907c665e7caf"></a>

## Direct properties — access_info.tls_config.use_host_header_as_sni / 0285c1af8142 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75e4c9d30f6d04a803f18245b4413875b5d77435bbc256863d9a455e97f72dd8"></a>

## Next pages — access_info.tls_config.use_host_header_as_sni / 0285c1af8142 / 4

- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-33dfe3eb515df9351eed33646d22e0fb7d9e1c20799988deb0717c8a05b81e99)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-b0918fbb2da2dea4ce096b94bf3e5be045867e7018c64c823efb2cddcebd7036"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b72bebcf0c6395b0087f155f4ba0def7f3f818a95ae022a9643f077a038bb667"></a>

## access_info.vault_auth_info — access_info.vault_auth_info / 9d354a40a6e0 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-b230284dcd12752087416741bf1b886d6e4f1996fde152460def8a29652cd810)
- access_info.vault_auth_info

<a id="canonical-06b5c70b89f1ed821b425a12b8dd3d4e102f09fc18943e9ec8cae91976f8f748"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters for Hashicorp Vault hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_role_auth",
    "token")}
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
  "x-ves-oneof-field-auth_params": "[\"app_role_auth\",\"token\"]"
}
```

Terraform syntax:

```terraform
vault_auth_info {
  # Configure direct properties listed below.
}
```
