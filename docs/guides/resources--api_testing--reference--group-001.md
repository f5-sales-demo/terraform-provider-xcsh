---
page_title: "xcsh_api_testing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing reference."
---

# xcsh_api_testing reference

<a id="canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abfab74fb470ff02aae3f1410ed982e15027197b522f18b253677ecceb5ab646"></a>

## Property reference — Property reference / be9f69378ddb / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- Property reference

<a id="canonical-9f2b1448cad16a1e4d2f03c43db5278f3e335b12f2a28c6df99e514332e7635f"></a>

## Direct properties — Property reference / be9f69378ddb / 3

<a id="canonical-b2d505784e05d93efdd2d1fc0b961f9cbf0fb3f31a09ec26e87a5cda9eeceefb"></a>

<a id="canonical-201c91e4a473712883dc1b9ba2cd9f9a85115098798871421a294420ad2c2fc9"></a>

## annotations property — Property reference / be9f69378ddb / 4

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

<a id="canonical-876d512e79974636b8d0cc260c1e290a99a6e58fe08208731d8bb35da36b5af1"></a>

<a id="canonical-5486616e6c248cef05a878dcf1de6075e4ef5a16652790ad6009c0435f21fc85"></a>

## custom_header_value property — Property reference / be9f69378ddb / 5

Type: `"string"`. Optional, Computed.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2f0693505f6a1595618e56d20078fd51ff7bd821a98ed51c17988218d618cc11"></a>

<a id="canonical-e43aa2619ddbdacbe833ef999da85dcc36f937353a01fa88e0769c75efe45112"></a>

## description property — Property reference / be9f69378ddb / 6

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

<a id="canonical-c34a3e473e4364b4b328700f63837b2765a65ff071b3ef0cbfa8fc11a4b3c0ec"></a>

<a id="canonical-e6597e1ee37f1d3e7728d94900c1350da0ce1661d848f6d5bbe8ad4e58ce487d"></a>

## disable property — Property reference / be9f69378ddb / 7

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

- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373): complete subsection reference.

- [every_day](resources--api_testing--reference--group-001.md#canonical-147fc8de30935590682ce653c0dd2328db1be0b14c86a08ded65e73c58eeca10): complete subsection reference.

- [every_month](resources--api_testing--reference--group-001.md#canonical-7b622820d86b72ff9602bc01a450f58e4f3721ea22114da4d50de6af92799abf): complete subsection reference.

- [every_week](resources--api_testing--reference--group-001.md#canonical-c3c9c7e68caf9fc72bf690b4168afb4e8af8bafc223971fd966c60fa2ff1a185): complete subsection reference.

<a id="canonical-ad22b0d66925526bf91cf19b5ad671c9e6f81863dfc95f1d6c585bf2d7b4b91d"></a>

<a id="canonical-d4f6bcafb82cf91c15e5065a43c030ffa0e3fbe3204ecda51a3b43883a0c91bd"></a>

## id property — Property reference / be9f69378ddb / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-26f7f6e1de264130c28cca622e432619801dde5a0a44011d551969645c349a19"></a>

<a id="canonical-2748922395185f7f6f43d7653a7063fd505214806911a38ca05c2f0fd2f34f3d"></a>

## labels property — Property reference / be9f69378ddb / 9

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

<a id="canonical-b10bbc9a3b32f011935d4a2ca3a064a09cb3419ef3def3e81cc8a04c8a36806e"></a>

<a id="canonical-4f82886a5d83bf145488e3bd0d129c878b0abe7d7d84c7df0707fad01fbf6efe"></a>

## name property — Property reference / be9f69378ddb / 10

Type: `"string"`. Required.

Name of the API Testing. Must be unique within the namespace.

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

<a id="canonical-452df52b9020a3fc76148c86c86a0dfe0b6b1499989aade1ff3b4ef5b77f9773"></a>

<a id="canonical-6e0a912de75feab5dba3cd0cf3661e265f686f12a7eb16a2f9337c590ad08edf"></a>

## namespace property — Property reference / be9f69378ddb / 11

Type: `"string"`. Required.

Namespace where the API Testing is created.

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

- [timeouts](resources--api_testing--reference--group-001.md#canonical-a8a6ace9796b1491e73969ca972f68e2372185930cf18d95669d078426a98c7c): complete subsection reference.

<a id="canonical-7f459ffb1a585d02785f0b6e1a31e12cb29e52a8cac3cdd325e96b2256dedd10"></a>

## All schema paths — Property reference / be9f69378ddb / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--api_testing--reference--group-001.md#canonical-b2d505784e05d93efdd2d1fc0b961f9cbf0fb3f31a09ec26e87a5cda9eeceefb) |
| `custom_header_value` | [custom_header_value](resources--api_testing--reference--group-001.md#canonical-876d512e79974636b8d0cc260c1e290a99a6e58fe08208731d8bb35da36b5af1) |
| `description` | [description](resources--api_testing--reference--group-001.md#canonical-2f0693505f6a1595618e56d20078fd51ff7bd821a98ed51c17988218d618cc11) |
| `disable` | [disable](resources--api_testing--reference--group-001.md#canonical-c34a3e473e4364b4b328700f63837b2765a65ff071b3ef0cbfa8fc11a4b3c0ec) |
| `domains` | [domains](resources--api_testing--reference--group-001.md#canonical-c051c3aa37a54be3738244202d198d00e73efbd86a7eb1c81bd731e08b8095e2) |
| `domains.allow_destructive_methods` | [domains.allow_destructive_methods](resources--api_testing--reference--group-001.md#canonical-4c398d2096e2804d85fdec4126989928a3a754c17ada403fbe78e6cc20723597) |
| `domains.credentials` | [domains.credentials](resources--api_testing--reference--group-001.md#canonical-f5fc7ab43054b3f7fcf6697141e83bd27823cbfaac29ccdbc8792ec1c930859b) |
| `domains.credentials.admin` | [domains.credentials.admin](resources--api_testing--reference--group-001.md#canonical-68d80addb1877c74fe2f2d9888191deb98a5aede4564f2f287040eddd7dda483) |
| `domains.credentials.api_key` | [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-95bce93468013263dd55d48452a4eb9110140617af8b53e8b9cae98505752815) |
| `domains.credentials.api_key.key` | [domains.credentials.api_key.key](resources--api_testing--reference--group-001.md#canonical-b49fde6ee29517355711674e88f8da5f7a5be7b26377e0bb6216e4f136f364e6) |
| `domains.credentials.api_key.value` | [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-de84068631cdec0414ca1ae45b4ccb8b281f12d26aacccc01d1a3b6cabbdd956) |
| `domains.credentials.api_key.value.blindfold_secret_info` | [domains.credentials.api_key.value.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-2dfd1dbf2a0824f6dbe8154cc49a2829b201e6ae54f1416c16af64ba324f20e7) |
| `domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` | [domains.credentials.api_key.value.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-362762034de1a0e2f731ed6057a53d18b91fc3e4d451996801e84c40ebfb0b04) |
| `domains.credentials.api_key.value.blindfold_secret_info.location` | [domains.credentials.api_key.value.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-87ee9400a82d175e8b98e505e7ca71b52eb0198af8d4331f48ca094b3c31ed02) |
| `domains.credentials.api_key.value.blindfold_secret_info.store_provider` | [domains.credentials.api_key.value.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-3ec5afdf0491c409b8bd30247ab4d96aed561965032b5710289398d18a32d42e) |
| `domains.credentials.api_key.value.clear_secret_info` | [domains.credentials.api_key.value.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-bb10cca514c16da10cb5a0ace5294e9079ea1fa826d61e2821a9534ad3bed9f9) |
| `domains.credentials.api_key.value.clear_secret_info.provider_ref` | [domains.credentials.api_key.value.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-af0d65a9b456c27dcd50d1072deb0be977ba9aab25aead05573b534903284e73) |
| `domains.credentials.api_key.value.clear_secret_info.url` | [domains.credentials.api_key.value.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-6162c7349e4fff1477efc2f9e945be9bd9801dd592b055fe1b537c57b73aa0b4) |
| `domains.credentials.basic_auth` | [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-dacd689c12d7e463ae81df0e430ad07c374d658f5540ae1e470e3f49ebc2426e) |
| `domains.credentials.basic_auth.password` | [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-b5c0c52b2f973b7d0caf278a1898cfd4a041c25a2280f26000e099c855823dca) |
| `domains.credentials.basic_auth.password.blindfold_secret_info` | [domains.credentials.basic_auth.password.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-c4279c1d9fa99cf1f47d5a12912cf56be51561b7432f661912b2ca4c427e3541) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-74295cb7204db17ea80e6a7b7be7b373055b45f3cb32d543c629f99c95e5d539) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.location` | [domains.credentials.basic_auth.password.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-54ce267a0cbff7f285dac27e9c1e837b3821227315a7e51099314b619420329e) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-a1490ae21ef42678d7a83330da30dfaddd1be5a5245f39f4b99d4134e60db1c7) |
| `domains.credentials.basic_auth.password.clear_secret_info` | [domains.credentials.basic_auth.password.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-40bc7c32651bf04162af389adf876ceef642880d32ec0bc42daa54664ddbfa58) |
| `domains.credentials.basic_auth.password.clear_secret_info.provider_ref` | [domains.credentials.basic_auth.password.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-ed249d49e8a1d1a64f3786b8366bbda056a27bfee7e1cd09cdcb5984c8cf9673) |
| `domains.credentials.basic_auth.password.clear_secret_info.url` | [domains.credentials.basic_auth.password.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-45e1e019dc24e3b88675c55429062d02a814ec01c42d59aadd45f407fdf2ed7f) |
| `domains.credentials.basic_auth.user` | [domains.credentials.basic_auth.user](resources--api_testing--reference--group-001.md#canonical-85f7794137464a25e663cb0dad48b43312ed34eeb4683a675a5fd2f9f2161f38) |
| `domains.credentials.bearer_token` | [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-a0da171cca9adc99943c889d400f36bcd026a4c9ec269a41b1139c8b9747c451) |
| `domains.credentials.bearer_token.token` | [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-defc0ca1b2b2db92ce8812aed0f2e3e6b729b0a6e883de2dcfd576770f6e1e9d) |
| `domains.credentials.bearer_token.token.blindfold_secret_info` | [domains.credentials.bearer_token.token.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-af08ca896a7fec6ee65f360413a127ce2955a65e32ea31a69b33bb97442f7453) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-5634a298be30268a15d827ae470b65b3d8b0a7e4f6884176a2664613405f539f) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.location` | [domains.credentials.bearer_token.token.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-7735cf557aa7e87951afb58eed3a669ca2b14204bfa10bf157522c0da74c11d7) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-b9ef80c58eac09289e9fd544077d083a1d7f2d165666060edf05c5ba6b48cff5) |
| `domains.credentials.bearer_token.token.clear_secret_info` | [domains.credentials.bearer_token.token.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-d8c1d5a0c16f6c88c5441feeeb3c836c0cef64a570441b6fc76911a944924f0d) |
| `domains.credentials.bearer_token.token.clear_secret_info.provider_ref` | [domains.credentials.bearer_token.token.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-f33c609cc7a5045f0d38de1fa12a39e99f08b941e60c6005e9d6267e8793e3e2) |
| `domains.credentials.bearer_token.token.clear_secret_info.url` | [domains.credentials.bearer_token.token.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-e99b78b4939924915015fc2694bc743df508c756e3d827c4de6de27ab2a5db3e) |
| `domains.credentials.credential_name` | [domains.credentials.credential_name](resources--api_testing--reference--group-001.md#canonical-52ac0f2490d95cb424fd21cf4da17c0d5da51990bb8bba6d05f5cf3ad5ede48e) |
| `domains.credentials.login_endpoint` | [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-0b249610947199f552c61805864680b6eccfd33b03c9c06d0a2222516d4409ab) |
| `domains.credentials.login_endpoint.json_payload` | [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-faa626646c17da12dc47587cf2581b5426255e2c07bec3a1dea891ec505f89d7) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-0ac4fa7a7d06df022100d3be5033efe8402399c8c7f4fc3ec030cadd5b0b04b6) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider](resources--api_testing--reference--group-001.md#canonical-d5fc04bfd07faf5fffb2fde04133f24353ec457f4f92c0b142639aa49392b09b) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location](resources--api_testing--reference--group-001.md#canonical-c2257da7bae99a6d82dc094248cfc6341a11f6220862bbe9d5a0995d8534a99c) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider](resources--api_testing--reference--group-001.md#canonical-d4edfffae09f73c4b64b53c5070a2425da49808ee9afe213a9b5ca85a375084e) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info` | [domains.credentials.login_endpoint.json_payload.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-c6031f2d83a1c92696c546bb1238f2d6dfc6b583bbcb7325c2fd5aa441d51e28) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref](resources--api_testing--reference--group-001.md#canonical-9ed0ff313e3d84ee45802bf67812496fab41631befe48bf82c7610642ad8dcae) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.url` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.url](resources--api_testing--reference--group-001.md#canonical-c4995ddb4074658728880d95287b18d573871d6f3a9180d0dc7fbe7467722ab2) |
| `domains.credentials.login_endpoint.method` | [domains.credentials.login_endpoint.method](resources--api_testing--reference--group-001.md#canonical-ba006d693fd326f7b03c89b13db984e9ecd3408647685dab438240eec827614f) |
| `domains.credentials.login_endpoint.path` | [domains.credentials.login_endpoint.path](resources--api_testing--reference--group-001.md#canonical-da4afd3902d5fcd677360f3c4b5e408b467683999b395ac0a9060ce7f5a6e93c) |
| `domains.credentials.login_endpoint.token_response_key` | [domains.credentials.login_endpoint.token_response_key](resources--api_testing--reference--group-001.md#canonical-c9889134c79879d343f0ccf1c1a8b2eb5191756493592ad3736bac46dd874ed1) |
| `domains.credentials.standard` | [domains.credentials.standard](resources--api_testing--reference--group-001.md#canonical-68c8b589c9f1aca58fb61ba0eb5784797cf25b38fe9c2979c8f0e150848028ab) |
| `domains.domain` | [domains.domain](resources--api_testing--reference--group-001.md#canonical-c9c57438be97624fce0de604631d7bbb5b5074ccffa8c3e5612273f0482fc73f) |
| `every_day` | [every_day](resources--api_testing--reference--group-001.md#canonical-593caf78d7818b562092d1fa4172cdfa11762ed044487ff6e0059581f96eca0b) |
| `every_month` | [every_month](resources--api_testing--reference--group-001.md#canonical-42519e6f1425876cdaa2006544b601bc1eb3e5d5ff15e49665c7e6be74a6be5a) |
| `every_week` | [every_week](resources--api_testing--reference--group-001.md#canonical-3d07c97bfc5c9ee6669c00303a440bea546a0a85626d44577407a55a6d45cc55) |
| `id` | [id](resources--api_testing--reference--group-001.md#canonical-ad22b0d66925526bf91cf19b5ad671c9e6f81863dfc95f1d6c585bf2d7b4b91d) |
| `labels` | [labels](resources--api_testing--reference--group-001.md#canonical-26f7f6e1de264130c28cca622e432619801dde5a0a44011d551969645c349a19) |
| `name` | [name](resources--api_testing--reference--group-001.md#canonical-b10bbc9a3b32f011935d4a2ca3a064a09cb3419ef3def3e81cc8a04c8a36806e) |
| `namespace` | [namespace](resources--api_testing--reference--group-001.md#canonical-452df52b9020a3fc76148c86c86a0dfe0b6b1499989aade1ff3b4ef5b77f9773) |
| `timeouts` | [timeouts](resources--api_testing--reference--group-001.md#canonical-b50793f1f21f2bd09a063c0f2dfaf807924999d97b3d96506d3e53f749585e33) |
| `timeouts.create` | [timeouts.create](resources--api_testing--reference--group-001.md#canonical-9d7b0c5b4de85600bf2da11c9a68d1f7ff174cdb4303b83557e4235892cfa0c9) |
| `timeouts.delete` | [timeouts.delete](resources--api_testing--reference--group-001.md#canonical-f1cfca34f0eacbcad36d592bb652915d47b28f6047e77602b6652cb3559852fa) |
| `timeouts.read` | [timeouts.read](resources--api_testing--reference--group-001.md#canonical-8ef61c7160bc072d04738684e481f641febc9f4290980170659a9d89e472d5c4) |
| `timeouts.update` | [timeouts.update](resources--api_testing--reference--group-001.md#canonical-a7f2500133f5a5b2969c9e215242884b172abfddb89e0883c3bc701e6a69debc) |

<a id="canonical-ce2f9d5429a7f74e06155eddee094ed156fc7ae3e7746b0040f1a7a3f836f9d4"></a>

## Next pages — Property reference / be9f69378ddb / 13

- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [every_day](resources--api_testing--reference--group-001.md#canonical-147fc8de30935590682ce653c0dd2328db1be0b14c86a08ded65e73c58eeca10)
- [every_month](resources--api_testing--reference--group-001.md#canonical-7b622820d86b72ff9602bc01a450f58e4f3721ea22114da4d50de6af92799abf)
- [every_week](resources--api_testing--reference--group-001.md#canonical-c3c9c7e68caf9fc72bf690b4168afb4e8af8bafc223971fd966c60fa2ff1a185)
- [timeouts](resources--api_testing--reference--group-001.md#canonical-a8a6ace9796b1491e73969ca972f68e2372185930cf18d95669d078426a98c7c)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7269ac172c5ebe3dce3a326ec7f50c935a80d3dc23ebd173d40a11e58ce1cfba"></a>

## domains — domains / 3061c817c16a / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- domains

<a id="canonical-c051c3aa37a54be3738244202d198d00e73efbd86a7eb1c81bd731e08b8095e2"></a>

Type: `"object"`. list nested block, Optional.

Add and configure testing domains and credentials.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credentials",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-ae4b0743ea46236f99496f2de767d56ba2769ed7a497fa5204d83c308e1e47f5"></a>

## Direct properties — domains / 3061c817c16a / 3

<a id="canonical-4c398d2096e2804d85fdec4126989928a3a754c17ada403fbe78e6cc20723597"></a>

<a id="canonical-867a392d19b9b70da90047a952e8562445222b9aa5258f26221130a3ac06ce89"></a>

## allow_destructive_methods property — domains / 3061c817c16a / 4

Type: `"bool"`. Optional.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71): complete subsection reference.

<a id="canonical-c9c57438be97624fce0de604631d7bbb5b5074ccffa8c3e5612273f0482fc73f"></a>

<a id="canonical-465c895f12def223ec2054344ba6e897324dfc3d5adf20da546b9defb68fb0c4"></a>

## domain property — domains / 3061c817c16a / 5

Type: `"string"`. Optional.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

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
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-82f30b3953336b6261c16cfb87d6f5e11e072932713b283f877d5823909c1208"></a>

## Next pages — domains / 3061c817c16a / 6

- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e25ff4b6da3fdcfe10785ddcabf9831c9ac5946e88a3669a15d31e44282f547"></a>

## domains.credentials — domains.credentials / 5500777ddb00 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- domains.credentials

<a id="canonical-f5fc7ab43054b3f7fcf6697141e83bd27823cbfaac29ccdbc8792ec1c930859b"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("credential_name"),
  validators.ConflictingListObjectAttributes("admin",
    "standard"),
  validators.ConflictingListObjectAttributes("api_key",
    "basic_auth"),
  validators.ConflictingListObjectAttributes("api_key",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("api_key",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "bearer_token"),
  validators.ConflictingListObjectAttributes("basic_auth",
    "login_endpoint"),
  validators.ConflictingListObjectAttributes("bearer_token",
    "login_endpoint")}
```

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
credentials {
  # Configure direct properties listed below.
}
```

<a id="canonical-0867f3c721b3fc9e3969b1fd1e85860c3f2c161ab7cab2176c871b27b2289ef1"></a>

## Direct properties — domains.credentials / 5500777ddb00 / 3

- [admin](resources--api_testing--reference--group-001.md#canonical-da9e255ea881db2a1f7e5c9443fec116712b73f9f7ed840b3b24ee2bb1bf0466): complete subsection reference.

- [api_key](resources--api_testing--reference--group-001.md#canonical-74ebdde0985ab0ff36d72b19a7255c4dba8938475505646717fe70a28930504a): complete subsection reference.

- [basic_auth](resources--api_testing--reference--group-001.md#canonical-8f96bed7bd1a828f6db40ae9afa438894ff9ad15e83b89bf6292f16d60f5a62e): complete subsection reference.

- [bearer_token](resources--api_testing--reference--group-001.md#canonical-9200370c944c66dcb9de731f30b9665f1a5e1b00509cf5537b5a83e199eee6bb): complete subsection reference.

<a id="canonical-52ac0f2490d95cb424fd21cf4da17c0d5da51990bb8bba6d05f5cf3ad5ede48e"></a>

<a id="canonical-f131f5056e364e4106dd1dd899f49b97283b83592085d5f63dce3672b431915c"></a>

## credential_name property — domains.credentials / 5500777ddb00 / 4

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [login_endpoint](resources--api_testing--reference--group-001.md#canonical-76d284b2134e4eaa9c0da193234c49bfe3cc57d78b6b7ba888f416557e4e6335): complete subsection reference.

- [standard](resources--api_testing--reference--group-001.md#canonical-392a3767cd1701d1c291803a560cf29188b82dd2b98a3d276c1bf22aaf8098e4): complete subsection reference.

<a id="canonical-4fb645a9e00359a907f33e4503b201528236c8533331f617868bd497bca6ddfb"></a>

## Next pages — domains.credentials / 5500777ddb00 / 5

- [domains.credentials.admin](resources--api_testing--reference--group-001.md#canonical-da9e255ea881db2a1f7e5c9443fec116712b73f9f7ed840b3b24ee2bb1bf0466)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-74ebdde0985ab0ff36d72b19a7255c4dba8938475505646717fe70a28930504a)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-8f96bed7bd1a828f6db40ae9afa438894ff9ad15e83b89bf6292f16d60f5a62e)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-9200370c944c66dcb9de731f30b9665f1a5e1b00509cf5537b5a83e199eee6bb)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-76d284b2134e4eaa9c0da193234c49bfe3cc57d78b6b7ba888f416557e4e6335)
- [domains.credentials.standard](resources--api_testing--reference--group-001.md#canonical-392a3767cd1701d1c291803a560cf29188b82dd2b98a3d276c1bf22aaf8098e4)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-da9e255ea881db2a1f7e5c9443fec116712b73f9f7ed840b3b24ee2bb1bf0466"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df5e3a0275956a13b67ec251f5b80b61253203dab5e0672cdfd051c765d78bac"></a>

## domains.credentials.admin — domains.credentials.admin / d9b898eefc94 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- domains.credentials.admin

<a id="canonical-68d80addb1877c74fe2f2d9888191deb98a5aede4564f2f287040eddd7dda483"></a>

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
admin = {}
```

<a id="canonical-bae7bc66edb4b9d9d7f023c81a066c29ccc621007c0bc50dac9f923a2711a7d2"></a>

## Direct properties — domains.credentials.admin / d9b898eefc94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64131f0ead6f6843c003ac14670bf9ed002ccbbad53ec3215b28893a2da79478"></a>

## Next pages — domains.credentials.admin / d9b898eefc94 / 4

- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-74ebdde0985ab0ff36d72b19a7255c4dba8938475505646717fe70a28930504a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75d182dcb617621119cc4e8734755c1a269c47520e527b72ee0369af97c6d025"></a>

## domains.credentials.api_key — domains.credentials.api_key / 1d5d1dc2a660 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- domains.credentials.api_key

<a id="canonical-95bce93468013263dd55d48452a4eb9110140617af8b53e8b9cae98505752815"></a>

Type: `"object"`. single nested block, Optional.

API Key

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("key")}
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
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-93e0a43107c123aceb799968f7f37119d139cde01f3a14c58cd1c09232169aba"></a>

## Direct properties — domains.credentials.api_key / 1d5d1dc2a660 / 3

<a id="canonical-b49fde6ee29517355711674e88f8da5f7a5be7b26377e0bb6216e4f136f364e6"></a>

<a id="canonical-f1109d25c21ce247c28fa1e5a7b8b91519dfd26904c74b043d3296473ed25aa0"></a>

## key property — domains.credentials.api_key / 1d5d1dc2a660 / 4

Type: `"string"`. Optional.

Key. Cryptographic key material

Upstream description:

Cryptographic key material

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](resources--api_testing--reference--group-001.md#canonical-d6db0922f50596ed1afbbe9feda391031764ce0da84e31b8cc67e9e0f685425b): complete subsection reference.

<a id="canonical-fef08413bc6eb9396dff91b66f0dd5b524f04d731a461d525edd8faacd194494"></a>

## Next pages — domains.credentials.api_key / 1d5d1dc2a660 / 5

- [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-d6db0922f50596ed1afbbe9feda391031764ce0da84e31b8cc67e9e0f685425b)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-d6db0922f50596ed1afbbe9feda391031764ce0da84e31b8cc67e9e0f685425b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b35b3f7d3bb6216f077ba3a5ca6761f490623064006c80098b13842a68205150"></a>

## domains.credentials.api_key.value — domains.credentials.api_key.value / 10f0bc085e16 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-74ebdde0985ab0ff36d72b19a7255c4dba8938475505646717fe70a28930504a)
- domains.credentials.api_key.value

<a id="canonical-de84068631cdec0414ca1ae45b4ccb8b281f12d26aacccc01d1a3b6cabbdd956"></a>

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
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-05ff5ec76b77ddfceb0c57246b825102228204a6d97115f440d54ee283cac8c5"></a>

## Direct properties — domains.credentials.api_key.value / 10f0bc085e16 / 3

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-fdac7567087db6d7de2d1044dd50561b449b1c831fc22325f6ad58157fe5dbd2): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-c17c663070cf8fafd200abcfbc92684bce9c5f6a1bfa31fa8000079583bf0065): complete subsection reference.

<a id="canonical-e391cc02128e13989ec37932a03d03c11938e3a91948749c0607eb1214286540"></a>

## Next pages — domains.credentials.api_key.value / 10f0bc085e16 / 4

- [domains.credentials.api_key.value.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-fdac7567087db6d7de2d1044dd50561b449b1c831fc22325f6ad58157fe5dbd2)
- [domains.credentials.api_key.value.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-c17c663070cf8fafd200abcfbc92684bce9c5f6a1bfa31fa8000079583bf0065)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-74ebdde0985ab0ff36d72b19a7255c4dba8938475505646717fe70a28930504a)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-fdac7567087db6d7de2d1044dd50561b449b1c831fc22325f6ad58157fe5dbd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aefdc2438d66b73ff9d02927a2bed9b1babfb0913a875a83bd09eda217fe9af6"></a>

## domains.credentials.api_key.value.blindfold_secret_info — domains.credentials.api_key.value.blindfold_secret_info / 43fe54a41b05 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-74ebdde0985ab0ff36d72b19a7255c4dba8938475505646717fe70a28930504a)
- [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-d6db0922f50596ed1afbbe9feda391031764ce0da84e31b8cc67e9e0f685425b)
- domains.credentials.api_key.value.blindfold_secret_info

<a id="canonical-2dfd1dbf2a0824f6dbe8154cc49a2829b201e6ae54f1416c16af64ba324f20e7"></a>

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

<a id="canonical-9c4fe7171faa62867ea46a237be02b2ee3402aa288d486b3912e0ec01a6d8ada"></a>

## Direct properties — domains.credentials.api_key.value.blindfold_secret_info / 43fe54a41b05 / 3

<a id="canonical-362762034de1a0e2f731ed6057a53d18b91fc3e4d451996801e84c40ebfb0b04"></a>

<a id="canonical-4ff905ba4f2a97773191a2e6befd09722b953f0b38a1437c8206a4e28181e2e7"></a>

## decryption_provider property — domains.credentials.api_key.value.blindfold_secret_info / 43fe54a41b05 / 4

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

<a id="canonical-87ee9400a82d175e8b98e505e7ca71b52eb0198af8d4331f48ca094b3c31ed02"></a>

<a id="canonical-c70e028c7215f1dac7370a03f5495f1d6e9a131f7c149786e13c1728a521ffe6"></a>

## location property — domains.credentials.api_key.value.blindfold_secret_info / 43fe54a41b05 / 5

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

<a id="canonical-3ec5afdf0491c409b8bd30247ab4d96aed561965032b5710289398d18a32d42e"></a>

<a id="canonical-ea2c77913ca928ff0ecdd91ff43f57a93b4f64082c862b306fae362d75e72da4"></a>

## store_provider property — domains.credentials.api_key.value.blindfold_secret_info / 43fe54a41b05 / 6

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

<a id="canonical-a20d34db64cf5b113d22d1ba9a0e6f1b714eff55204ab66875fa9a33dc173bfd"></a>

## Next pages — domains.credentials.api_key.value.blindfold_secret_info / 43fe54a41b05 / 7

- [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-d6db0922f50596ed1afbbe9feda391031764ce0da84e31b8cc67e9e0f685425b)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-c17c663070cf8fafd200abcfbc92684bce9c5f6a1bfa31fa8000079583bf0065"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a187759619ef9eea18ccda94b3ef6e3183397d6a22a7ef417869c652f0221e2"></a>

## domains.credentials.api_key.value.clear_secret_info — domains.credentials.api_key.value.clear_secret_info / 88ef836886f3 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.api_key](resources--api_testing--reference--group-001.md#canonical-74ebdde0985ab0ff36d72b19a7255c4dba8938475505646717fe70a28930504a)
- [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-d6db0922f50596ed1afbbe9feda391031764ce0da84e31b8cc67e9e0f685425b)
- domains.credentials.api_key.value.clear_secret_info

<a id="canonical-bb10cca514c16da10cb5a0ace5294e9079ea1fa826d61e2821a9534ad3bed9f9"></a>

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

<a id="canonical-718ff9bc4596d8d56fab5882bf65148782244980ccb44c2bb8256bb7b4129495"></a>

## Direct properties — domains.credentials.api_key.value.clear_secret_info / 88ef836886f3 / 3

<a id="canonical-af0d65a9b456c27dcd50d1072deb0be977ba9aab25aead05573b534903284e73"></a>

<a id="canonical-217df66d7a642d4b77d2fbdfb341be1e2a5e58b02ded2ccbda6dda4e50db782d"></a>

## provider_ref property — domains.credentials.api_key.value.clear_secret_info / 88ef836886f3 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6162c7349e4fff1477efc2f9e945be9bd9801dd592b055fe1b537c57b73aa0b4"></a>

<a id="canonical-573e5c69a66ce3458cb14cd3db4b37b40ece6be400f10824fc6821a2a1a732f2"></a>

## url property — domains.credentials.api_key.value.clear_secret_info / 88ef836886f3 / 5

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

<a id="canonical-49403139557517d84b17ce3ac6f5fefdfe65b72d977ce1c49eaa3729fcd9a077"></a>

## Next pages — domains.credentials.api_key.value.clear_secret_info / 88ef836886f3 / 6

- [domains.credentials.api_key.value](resources--api_testing--reference--group-001.md#canonical-d6db0922f50596ed1afbbe9feda391031764ce0da84e31b8cc67e9e0f685425b)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-8f96bed7bd1a828f6db40ae9afa438894ff9ad15e83b89bf6292f16d60f5a62e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-468f4cb2ea05fe7d2c69f315bf62ffba31be94e1d7e2c0520f16fa5009d8eb5d"></a>

## domains.credentials.basic_auth — domains.credentials.basic_auth / 21836fd81175 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- domains.credentials.basic_auth

<a id="canonical-dacd689c12d7e463ae81df0e430ad07c374d658f5540ae1e470e3f49ebc2426e"></a>

Type: `"object"`. single nested block, Optional.

Basic Authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("user")}
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
basic_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-3aae6af415af296f15fee65a1dcef8958cb061caa01834c04a6a48560994b5c6"></a>

## Direct properties — domains.credentials.basic_auth / 21836fd81175 / 3

- [password](resources--api_testing--reference--group-001.md#canonical-3d057dac139b6643f1762ba40c3b1f6c802ebe3eabd2d151391e6bdc3d540cb9): complete subsection reference.

<a id="canonical-85f7794137464a25e663cb0dad48b43312ed34eeb4683a675a5fd2f9f2161f38"></a>

<a id="canonical-d42987cafcf208626414337630e1c5c35173b1808299148f85250f3a00313d6c"></a>

## user property — domains.credentials.basic_auth / 21836fd81175 / 4

Type: `"string"`. Optional.

User. Configuration parameter for user

Upstream description:

Configuration parameter for user

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-8dd0e8ce10ff4967b0a3e6497219a941e0f6f1c09bfbeddaa7b7cea9f84c698a"></a>

## Next pages — domains.credentials.basic_auth / 21836fd81175 / 5

- [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-3d057dac139b6643f1762ba40c3b1f6c802ebe3eabd2d151391e6bdc3d540cb9)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-3d057dac139b6643f1762ba40c3b1f6c802ebe3eabd2d151391e6bdc3d540cb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cd940d2ae1887796d5d2368ea49b8ac2bdaf64abe8dfeb0b5e62eb70333da7e"></a>

## domains.credentials.basic_auth.password — domains.credentials.basic_auth.password / 1ac6c0442d83 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-8f96bed7bd1a828f6db40ae9afa438894ff9ad15e83b89bf6292f16d60f5a62e)
- domains.credentials.basic_auth.password

<a id="canonical-b5c0c52b2f973b7d0caf278a1898cfd4a041c25a2280f26000e099c855823dca"></a>

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

<a id="canonical-4fa347b8e52cba11fcf20a693f55c8f534e0013f3f4de62ebaa9682d314fa019"></a>

## Direct properties — domains.credentials.basic_auth.password / 1ac6c0442d83 / 3

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-696a553382ff058088693731be0e0f0316e84c19e9425e81d62d098224f94e6f): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-e6a136320c7c0ce0bd76fc293a354812cd6ffc2d99c25582dbc209b596ff5028): complete subsection reference.

<a id="canonical-a3a639b2798104932529e8e76c1118e65e96ea7f43551ff2373df89c0d77c7c4"></a>

## Next pages — domains.credentials.basic_auth.password / 1ac6c0442d83 / 4

- [domains.credentials.basic_auth.password.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-696a553382ff058088693731be0e0f0316e84c19e9425e81d62d098224f94e6f)
- [domains.credentials.basic_auth.password.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-e6a136320c7c0ce0bd76fc293a354812cd6ffc2d99c25582dbc209b596ff5028)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-8f96bed7bd1a828f6db40ae9afa438894ff9ad15e83b89bf6292f16d60f5a62e)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-696a553382ff058088693731be0e0f0316e84c19e9425e81d62d098224f94e6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61b086100da979aca9179e90d4d08bc45767a6f81d69f91513a0130c6cb0651b"></a>

## domains.credentials.basic_auth.password.blindfold_secret_info — domains.credentials.basic_auth.password.blindfold_secret_info / 5348bad6f5f0 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-8f96bed7bd1a828f6db40ae9afa438894ff9ad15e83b89bf6292f16d60f5a62e)
- [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-3d057dac139b6643f1762ba40c3b1f6c802ebe3eabd2d151391e6bdc3d540cb9)
- domains.credentials.basic_auth.password.blindfold_secret_info

<a id="canonical-c4279c1d9fa99cf1f47d5a12912cf56be51561b7432f661912b2ca4c427e3541"></a>

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

<a id="canonical-55ce40f274fe19e349dc01256817b0f6be5daf876c760fef50275b8e9aa83669"></a>

## Direct properties — domains.credentials.basic_auth.password.blindfold_secret_info / 5348bad6f5f0 / 3

<a id="canonical-74295cb7204db17ea80e6a7b7be7b373055b45f3cb32d543c629f99c95e5d539"></a>

<a id="canonical-69114ac75051b6d86eaaa03a64206a229af4a736b72a215660253c64674fa6e2"></a>

## decryption_provider property — domains.credentials.basic_auth.password.blindfold_secret_info / 5348bad6f5f0 / 4

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

<a id="canonical-54ce267a0cbff7f285dac27e9c1e837b3821227315a7e51099314b619420329e"></a>

<a id="canonical-4ef81fe2b968d4833cfb3ed1cb2d4d1ecc2697fc4fc79287f9b8170b5aa90341"></a>

## location property — domains.credentials.basic_auth.password.blindfold_secret_info / 5348bad6f5f0 / 5

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

<a id="canonical-a1490ae21ef42678d7a83330da30dfaddd1be5a5245f39f4b99d4134e60db1c7"></a>

<a id="canonical-6e38f84876116ce697257421ca8eeb150d17763c25931e20386247f3a5ed1220"></a>

## store_provider property — domains.credentials.basic_auth.password.blindfold_secret_info / 5348bad6f5f0 / 6

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

<a id="canonical-8421411d1713763171bfd64c497b788dec4e49b66dc786a1e5d09135d9995a04"></a>

## Next pages — domains.credentials.basic_auth.password.blindfold_secret_info / 5348bad6f5f0 / 7

- [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-3d057dac139b6643f1762ba40c3b1f6c802ebe3eabd2d151391e6bdc3d540cb9)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-e6a136320c7c0ce0bd76fc293a354812cd6ffc2d99c25582dbc209b596ff5028"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-552916af249d6937c39e1e24fa4d53294c87a92b3261232e7bff7ca587577949"></a>

## domains.credentials.basic_auth.password.clear_secret_info — domains.credentials.basic_auth.password.clear_secret_info / cf2886ae97a7 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.basic_auth](resources--api_testing--reference--group-001.md#canonical-8f96bed7bd1a828f6db40ae9afa438894ff9ad15e83b89bf6292f16d60f5a62e)
- [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-3d057dac139b6643f1762ba40c3b1f6c802ebe3eabd2d151391e6bdc3d540cb9)
- domains.credentials.basic_auth.password.clear_secret_info

<a id="canonical-40bc7c32651bf04162af389adf876ceef642880d32ec0bc42daa54664ddbfa58"></a>

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

<a id="canonical-98206802a31ffebf30110a8f46c6ec98e57a0ac41409921ba139b472f24255fd"></a>

## Direct properties — domains.credentials.basic_auth.password.clear_secret_info / cf2886ae97a7 / 3

<a id="canonical-ed249d49e8a1d1a64f3786b8366bbda056a27bfee7e1cd09cdcb5984c8cf9673"></a>

<a id="canonical-b986e4af5e133b306a6d1cf432252fca58a413671600ac20551a43281178b7d1"></a>

## provider_ref property — domains.credentials.basic_auth.password.clear_secret_info / cf2886ae97a7 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-45e1e019dc24e3b88675c55429062d02a814ec01c42d59aadd45f407fdf2ed7f"></a>

<a id="canonical-26b4484e5d5dafa8f4da148bc7d6215bbc3b4120171702e3202a5dcafde93bbd"></a>

## url property — domains.credentials.basic_auth.password.clear_secret_info / cf2886ae97a7 / 5

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

<a id="canonical-aadebcd6febf08bbe33f0918e0e20db831f1f0da1b3683b7960e4aec4491c550"></a>

## Next pages — domains.credentials.basic_auth.password.clear_secret_info / cf2886ae97a7 / 6

- [domains.credentials.basic_auth.password](resources--api_testing--reference--group-001.md#canonical-3d057dac139b6643f1762ba40c3b1f6c802ebe3eabd2d151391e6bdc3d540cb9)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-9200370c944c66dcb9de731f30b9665f1a5e1b00509cf5537b5a83e199eee6bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f53300339ec140ce3b7062bcc1cefc4fca1791662592253dbff45a8f77c1f91a"></a>

## domains.credentials.bearer_token — domains.credentials.bearer_token / 2b0ef1d22ba3 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- domains.credentials.bearer_token

<a id="canonical-a0da171cca9adc99943c889d400f36bcd026a4c9ec269a41b1139c8b9747c451"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

Receipt-pinned upstream constraints:

```json
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
bearer_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-03deee5174ae9d1799b4d689e2f883b4b4f2091faabe9baa56eeeb95e1370970"></a>

## Direct properties — domains.credentials.bearer_token / 2b0ef1d22ba3 / 3

- [token](resources--api_testing--reference--group-001.md#canonical-29d3cae017945904b6685ae065b3bba3cef022d83e098e26fb0a4ffa501224d0): complete subsection reference.

<a id="canonical-c525e7de1498c548190cc4807b472c8cba5d85e3eb2010c1828f55ee2ecb6e67"></a>

## Next pages — domains.credentials.bearer_token / 2b0ef1d22ba3 / 4

- [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-29d3cae017945904b6685ae065b3bba3cef022d83e098e26fb0a4ffa501224d0)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-29d3cae017945904b6685ae065b3bba3cef022d83e098e26fb0a4ffa501224d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cfa26ee44b37a95bbd12a5c65378d7c266d8d220250cf82fe58cd936ed0f9fb"></a>

## domains.credentials.bearer_token.token — domains.credentials.bearer_token.token / c4896ff992f6 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-9200370c944c66dcb9de731f30b9665f1a5e1b00509cf5537b5a83e199eee6bb)
- domains.credentials.bearer_token.token

<a id="canonical-defc0ca1b2b2db92ce8812aed0f2e3e6b729b0a6e883de2dcfd576770f6e1e9d"></a>

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
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-71939263895683cd0a8a20837b3a9a11acfaf9403b124a44c7c3346ee3a21a8f"></a>

## Direct properties — domains.credentials.bearer_token.token / c4896ff992f6 / 3

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-4e70b92cf8e022698d6a551bc927f445731d5b272c7875e43c2f6a3f47d930eb): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-dc6336fc4b26f4ab18fc79189a818350698fccacbf5c3f40cba63713e0ea1783): complete subsection reference.

<a id="canonical-2ced3327c49cb7fe98ce860ea69a55587fb3f240feb6c29f6f3dbff980718446"></a>

## Next pages — domains.credentials.bearer_token.token / c4896ff992f6 / 4

- [domains.credentials.bearer_token.token.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-4e70b92cf8e022698d6a551bc927f445731d5b272c7875e43c2f6a3f47d930eb)
- [domains.credentials.bearer_token.token.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-dc6336fc4b26f4ab18fc79189a818350698fccacbf5c3f40cba63713e0ea1783)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-9200370c944c66dcb9de731f30b9665f1a5e1b00509cf5537b5a83e199eee6bb)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-4e70b92cf8e022698d6a551bc927f445731d5b272c7875e43c2f6a3f47d930eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8aec6876e0171b5b62ab86a1c9d3137c624d1156b447084c43d54cb46fa71e4"></a>

## domains.credentials.bearer_token.token.blindfold_secret_info — domains.credentials.bearer_token.token.blindfold_secret_info / 68cb5da4d3fb / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-9200370c944c66dcb9de731f30b9665f1a5e1b00509cf5537b5a83e199eee6bb)
- [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-29d3cae017945904b6685ae065b3bba3cef022d83e098e26fb0a4ffa501224d0)
- domains.credentials.bearer_token.token.blindfold_secret_info

<a id="canonical-af08ca896a7fec6ee65f360413a127ce2955a65e32ea31a69b33bb97442f7453"></a>

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

<a id="canonical-4eb6b6b058b563797811f8c9110da3a6de5d59c95eca35c6aaa096d9b1bd2f8e"></a>

## Direct properties — domains.credentials.bearer_token.token.blindfold_secret_info / 68cb5da4d3fb / 3

<a id="canonical-5634a298be30268a15d827ae470b65b3d8b0a7e4f6884176a2664613405f539f"></a>

<a id="canonical-75217d169596d670c860f4f60164f98cd401465759b87551527d13b27c87025c"></a>

## decryption_provider property — domains.credentials.bearer_token.token.blindfold_secret_info / 68cb5da4d3fb / 4

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

<a id="canonical-7735cf557aa7e87951afb58eed3a669ca2b14204bfa10bf157522c0da74c11d7"></a>

<a id="canonical-6b652a7e6bde745001f6dd81e8bec76f74bee572287649c210b82e1961f2b147"></a>

## location property — domains.credentials.bearer_token.token.blindfold_secret_info / 68cb5da4d3fb / 5

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

<a id="canonical-b9ef80c58eac09289e9fd544077d083a1d7f2d165666060edf05c5ba6b48cff5"></a>

<a id="canonical-b1c2fcdf9d252b70a54892742796140befc7c6abc6bd819d88d466eaf19185ac"></a>

## store_provider property — domains.credentials.bearer_token.token.blindfold_secret_info / 68cb5da4d3fb / 6

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

<a id="canonical-ace047017673d39a497cf6f8541efd7563b84cb177ec4f2c97c1a7e43c2ac133"></a>

## Next pages — domains.credentials.bearer_token.token.blindfold_secret_info / 68cb5da4d3fb / 7

- [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-29d3cae017945904b6685ae065b3bba3cef022d83e098e26fb0a4ffa501224d0)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-dc6336fc4b26f4ab18fc79189a818350698fccacbf5c3f40cba63713e0ea1783"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcf3159d7df05f6b37615b2fa540f815a95bfa45597bc9246facb5b80b64b3bc"></a>

## domains.credentials.bearer_token.token.clear_secret_info — domains.credentials.bearer_token.token.clear_secret_info / bcf92675bfa1 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.bearer_token](resources--api_testing--reference--group-001.md#canonical-9200370c944c66dcb9de731f30b9665f1a5e1b00509cf5537b5a83e199eee6bb)
- [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-29d3cae017945904b6685ae065b3bba3cef022d83e098e26fb0a4ffa501224d0)
- domains.credentials.bearer_token.token.clear_secret_info

<a id="canonical-d8c1d5a0c16f6c88c5441feeeb3c836c0cef64a570441b6fc76911a944924f0d"></a>

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

<a id="canonical-13fa00383fb369d53974281be0511d299f50d4b5a4e1c7466a4c40d89f2de695"></a>

## Direct properties — domains.credentials.bearer_token.token.clear_secret_info / bcf92675bfa1 / 3

<a id="canonical-f33c609cc7a5045f0d38de1fa12a39e99f08b941e60c6005e9d6267e8793e3e2"></a>

<a id="canonical-f0c1a2931826d25105a208603344af653c5e355db6bd3e8852242ca86ebca31e"></a>

## provider_ref property — domains.credentials.bearer_token.token.clear_secret_info / bcf92675bfa1 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e99b78b4939924915015fc2694bc743df508c756e3d827c4de6de27ab2a5db3e"></a>

<a id="canonical-18bbe9c2826335c80ee2d66d13f9c477f76f88cdb40a13c8a67499d592c29e2f"></a>

## url property — domains.credentials.bearer_token.token.clear_secret_info / bcf92675bfa1 / 5

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

<a id="canonical-4216b6705d4862be1ac63aef7cc28bc3fb8dba2dd01c0aa3cd32329903b035c3"></a>

## Next pages — domains.credentials.bearer_token.token.clear_secret_info / bcf92675bfa1 / 6

- [domains.credentials.bearer_token.token](resources--api_testing--reference--group-001.md#canonical-29d3cae017945904b6685ae065b3bba3cef022d83e098e26fb0a4ffa501224d0)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-76d284b2134e4eaa9c0da193234c49bfe3cc57d78b6b7ba888f416557e4e6335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4553b8cbb5d84b6b163c4086c6bb74108155364e358ec3cbc156efcb70f17d61"></a>

## domains.credentials.login_endpoint — domains.credentials.login_endpoint / a8ea1a8cb867 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- domains.credentials.login_endpoint

<a id="canonical-0b249610947199f552c61805864680b6eccfd33b03c9c06d0a2222516d4409ab"></a>

Type: `"object"`. single nested block, Optional.

Login Endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path",
    "token_response_key")}
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
login_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-f59a05c568ad35eb4520b676879817769fe94b5bb66182a3ac169e74588a2366"></a>

## Direct properties — domains.credentials.login_endpoint / a8ea1a8cb867 / 3

- [json_payload](resources--api_testing--reference--group-001.md#canonical-8a267a21aff51f9f76bf1548d2b55839b3b7d5f918a357eca143e1a68504119b): complete subsection reference.

<a id="canonical-ba006d693fd326f7b03c89b13db984e9ecd3408647685dab438240eec827614f"></a>

<a id="canonical-7135cc50e5744ca17b53f29882660bc95658f0ba3a76ec38280c65848fbfe5b3"></a>

## method property — domains.credentials.login_endpoint / a8ea1a8cb867 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-da4afd3902d5fcd677360f3c4b5e408b467683999b395ac0a9060ce7f5a6e93c"></a>

<a id="canonical-614f96adc476222ae6d41a8bb5337b851fab8e7a6bb20f244aceb9dce9417919"></a>

## path property — domains.credentials.login_endpoint / a8ea1a8cb867 / 5

Type: `"string"`. Optional.

Path. URL path for the endpoint

Upstream description:

URL path for the endpoint

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-c9889134c79879d343f0ccf1c1a8b2eb5191756493592ad3736bac46dd874ed1"></a>

<a id="canonical-580f98f85edce975550b9f8e33c40ba0da4111342dc9e56a45d028d669280979"></a>

## token_response_key property — domains.credentials.login_endpoint / a8ea1a8cb867 / 6

Type: `"string"`. Optional.

Configuration parameter for token response key.

Upstream description:

Configuration parameter for token response key

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

<a id="canonical-a9f9faadb2771f5649a1cdd16c7633cb4485c5117f9099b8ccaf7eac4c53b292"></a>

## Next pages — domains.credentials.login_endpoint / a8ea1a8cb867 / 7

- [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-8a267a21aff51f9f76bf1548d2b55839b3b7d5f918a357eca143e1a68504119b)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-8a267a21aff51f9f76bf1548d2b55839b3b7d5f918a357eca143e1a68504119b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a0af24b1f4b96d5608397034993c54e69d3e4b3e060987f808545790bddbeb8"></a>

## domains.credentials.login_endpoint.json_payload — domains.credentials.login_endpoint.json_payload / f091ca13779e / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-76d284b2134e4eaa9c0da193234c49bfe3cc57d78b6b7ba888f416557e4e6335)
- domains.credentials.login_endpoint.json_payload

<a id="canonical-faa626646c17da12dc47587cf2581b5426255e2c07bec3a1dea891ec505f89d7"></a>

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
json_payload {
  # Configure direct properties listed below.
}
```

<a id="canonical-089a15ec3d8cb88885cc36d04e581d6a075a91e485719350acd8487c7f041f31"></a>

## Direct properties — domains.credentials.login_endpoint.json_payload / f091ca13779e / 3

- [blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-4d7218e39cf4bb676db01415d8b19968a74eb8050c5678b1e022bba51c64cde6): complete subsection reference.

- [clear_secret_info](resources--api_testing--reference--group-001.md#canonical-50279cf9273a95f1ad27aed00edc139b597cce4ff93fb368dda49799cd97df93): complete subsection reference.

<a id="canonical-1e49f8077a0f6916c2204c0e17e6f48cd1913e3804776d14d87f581a489c48ef"></a>

## Next pages — domains.credentials.login_endpoint.json_payload / f091ca13779e / 4

- [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](resources--api_testing--reference--group-001.md#canonical-4d7218e39cf4bb676db01415d8b19968a74eb8050c5678b1e022bba51c64cde6)
- [domains.credentials.login_endpoint.json_payload.clear_secret_info](resources--api_testing--reference--group-001.md#canonical-50279cf9273a95f1ad27aed00edc139b597cce4ff93fb368dda49799cd97df93)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-76d284b2134e4eaa9c0da193234c49bfe3cc57d78b6b7ba888f416557e4e6335)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-4d7218e39cf4bb676db01415d8b19968a74eb8050c5678b1e022bba51c64cde6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a890fcac39c3aed7e4953b0c64a0178eda64310fc36c2d95e47bff84efde0e80"></a>

## domains.credentials.login_endpoint.json_payload.blindfold_secret_info — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / d9e1fdc3fb66 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-76d284b2134e4eaa9c0da193234c49bfe3cc57d78b6b7ba888f416557e4e6335)
- [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-8a267a21aff51f9f76bf1548d2b55839b3b7d5f918a357eca143e1a68504119b)
- domains.credentials.login_endpoint.json_payload.blindfold_secret_info

<a id="canonical-0ac4fa7a7d06df022100d3be5033efe8402399c8c7f4fc3ec030cadd5b0b04b6"></a>

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

<a id="canonical-76728b446b388c7bad7735c85a59d1c7d77c941fe4eb6c14e9cd3ccf50f7be38"></a>

## Direct properties — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / d9e1fdc3fb66 / 3

<a id="canonical-d5fc04bfd07faf5fffb2fde04133f24353ec457f4f92c0b142639aa49392b09b"></a>

<a id="canonical-d20b6b99dabad549aaab6b66cd0a42147949dae963a9b1e41cc339f9f0252721"></a>

## decryption_provider property — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / d9e1fdc3fb66 / 4

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

<a id="canonical-c2257da7bae99a6d82dc094248cfc6341a11f6220862bbe9d5a0995d8534a99c"></a>

<a id="canonical-4b872b070549ddf955a39b055b0467dadfd992043abfd7816e4f5dbf3f525b33"></a>

## location property — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / d9e1fdc3fb66 / 5

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

<a id="canonical-d4edfffae09f73c4b64b53c5070a2425da49808ee9afe213a9b5ca85a375084e"></a>

<a id="canonical-628071cc1f800b5a5652762d2dcd352cd1326edd331010a8bc300149a9688ba6"></a>

## store_provider property — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / d9e1fdc3fb66 / 6

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

<a id="canonical-f33e7c93cf1c8b46a3fa9fc01a906eb35bd83c403ecf2bb6c67110a6a26dfd0a"></a>

## Next pages — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / d9e1fdc3fb66 / 7

- [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-8a267a21aff51f9f76bf1548d2b55839b3b7d5f918a357eca143e1a68504119b)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-50279cf9273a95f1ad27aed00edc139b597cce4ff93fb368dda49799cd97df93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6235170aec0b861bb1756039ab3259a676b79cd1f7bdb729cc0a7d5ce7660520"></a>

## domains.credentials.login_endpoint.json_payload.clear_secret_info — domains.credentials.login_endpoint.json_payload.clear_secret_info / 4fa707daffcc / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [domains.credentials.login_endpoint](resources--api_testing--reference--group-001.md#canonical-76d284b2134e4eaa9c0da193234c49bfe3cc57d78b6b7ba888f416557e4e6335)
- [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-8a267a21aff51f9f76bf1548d2b55839b3b7d5f918a357eca143e1a68504119b)
- domains.credentials.login_endpoint.json_payload.clear_secret_info

<a id="canonical-c6031f2d83a1c92696c546bb1238f2d6dfc6b583bbcb7325c2fd5aa441d51e28"></a>

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

<a id="canonical-cf05d11e81a44b04068bf07278976948b4291687c1bd7d3d2b69ff52f55aabcf"></a>

## Direct properties — domains.credentials.login_endpoint.json_payload.clear_secret_info / 4fa707daffcc / 3

<a id="canonical-9ed0ff313e3d84ee45802bf67812496fab41631befe48bf82c7610642ad8dcae"></a>

<a id="canonical-4406d284e2842ed0b81b963c1ff425e66d2bb0e2d48f253844772d67f98eecdb"></a>

## provider_ref property — domains.credentials.login_endpoint.json_payload.clear_secret_info / 4fa707daffcc / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c4995ddb4074658728880d95287b18d573871d6f3a9180d0dc7fbe7467722ab2"></a>

<a id="canonical-aa32cfd34251d1202c9b933e8b516c71c3b480bb4e73dec70f755aa777c9bc6d"></a>

## url property — domains.credentials.login_endpoint.json_payload.clear_secret_info / 4fa707daffcc / 5

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

<a id="canonical-7f8624be0830e816eeb0ee81b6af763f288b21a3853bbd1336e7afcb39fb3a27"></a>

## Next pages — domains.credentials.login_endpoint.json_payload.clear_secret_info / 4fa707daffcc / 6

- [domains.credentials.login_endpoint.json_payload](resources--api_testing--reference--group-001.md#canonical-8a267a21aff51f9f76bf1548d2b55839b3b7d5f918a357eca143e1a68504119b)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-392a3767cd1701d1c291803a560cf29188b82dd2b98a3d276c1bf22aaf8098e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20a08ca3442b76a832725c922643d51535fa6c2865060c042e87e7b1b8a78633"></a>

## domains.credentials.standard — domains.credentials.standard / e68512c22698 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [domains](resources--api_testing--reference--group-001.md#canonical-ab28fb1e5008d84fb0c90fb00284bdddcfcfa022e3505f6759aea345b888e373)
- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- domains.credentials.standard

<a id="canonical-68c8b589c9f1aca58fb61ba0eb5784797cf25b38fe9c2979c8f0e150848028ab"></a>

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
standard = {}
```

<a id="canonical-b2305a3cd01c94a61eeabac84bf7d3f7cc5d6e8a58810483f0229eb632d5e584"></a>

## Direct properties — domains.credentials.standard / e68512c22698 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a68ac993891276da8eaf3f41e02b4ec5bacb33d54555f52c39cace5ba77c4a48"></a>

## Next pages — domains.credentials.standard / e68512c22698 / 4

- [domains.credentials](resources--api_testing--reference--group-001.md#canonical-63a1523326979513ae3b05501a2375217a1db81ede2f96cadfa110d32aae5e71)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-147fc8de30935590682ce653c0dd2328db1be0b14c86a08ded65e73c58eeca10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce890c2c43c806fdc99e01f0c6670b789a827fb3a1b2d19054de94293db1f4c0"></a>

## every_day — every_day / bd3eb1d5de9b / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- every_day

<a id="canonical-593caf78d7818b562092d1fa4172cdfa11762ed044487ff6e0059581f96eca0b"></a>

Type: `["object", {}]`. Optional.

\[OneOf: every\_day, every\_month, every\_week\] Enable this option

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

- [every_day](resources--api_testing--reference--group-001.md#canonical-593caf78d7818b562092d1fa4172cdfa11762ed044487ff6e0059581f96eca0b)
- [every_month](resources--api_testing--reference--group-001.md#canonical-42519e6f1425876cdaa2006544b601bc1eb3e5d5ff15e49665c7e6be74a6be5a)
- [every_week](resources--api_testing--reference--group-001.md#canonical-3d07c97bfc5c9ee6669c00303a440bea546a0a85626d44577407a55a6d45cc55)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
every_day = {}
```

<a id="canonical-156c691cee67d194adb08eb8c653f676ae5af9bf9c28ee18a5790c1b55ee023d"></a>

## Direct properties — every_day / bd3eb1d5de9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9d1711a882f2ff6e5136dbab591192b356ab0726aacbdc4ee5356e0711b403d"></a>

## Next pages — every_day / bd3eb1d5de9b / 4

- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-7b622820d86b72ff9602bc01a450f58e4f3721ea22114da4d50de6af92799abf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86913acc808c077e9e00daccf802db66c180c8888d5cd37ec11fe1fadb94ea49"></a>

## every_month — every_month / 4abbafd2e2df / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- every_month

<a id="canonical-42519e6f1425876cdaa2006544b601bc1eb3e5d5ff15e49665c7e6be74a6be5a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for every month.

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
every_month = {}
```

<a id="canonical-6d66967a055be5ecde1c80a8ff87b31f66c80452252c1f89e4695d795ef153e0"></a>

## Direct properties — every_month / 4abbafd2e2df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd447184fec0edb5869237edb3120c99719b9ff517e81cd6ee94a326202dcc5d"></a>

## Next pages — every_month / 4abbafd2e2df / 4

- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-c3c9c7e68caf9fc72bf690b4168afb4e8af8bafc223971fd966c60fa2ff1a185"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28e59da17af3e69fecda255106892293cd41086f2b1871248cdb4690fad178e2"></a>

## every_week — every_week / 355a9b951e29 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- every_week

<a id="canonical-3d07c97bfc5c9ee6669c00303a440bea546a0a85626d44577407a55a6d45cc55"></a>

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
every_week = {}
```

<a id="canonical-c89ae1f761bf3c2452be5706250de7f9fb6a285fc2cc2e8a4feb5b93aaa8a427"></a>

## Direct properties — every_week / 355a9b951e29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73dc17c50b34efd8d7887e5c77be8e9f4994a0494139e3aee4573b3f2836a786"></a>

## Next pages — every_week / 355a9b951e29 / 4

- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-a8a6ace9796b1491e73969ca972f68e2372185930cf18d95669d078426a98c7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db12692108bf05f1ab51a116d5a8acb44bdde2562f8c71293236af17be38981d"></a>

## timeouts — timeouts / e156b6b3d717 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- timeouts

<a id="canonical-b50793f1f21f2bd09a063c0f2dfaf807924999d97b3d96506d3e53f749585e33"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-c873e702b3c7089ea26346e47b1acc5a76272390cfb02f3b02c800200c6d2c7c"></a>

## Direct properties — timeouts / e156b6b3d717 / 3

<a id="canonical-9d7b0c5b4de85600bf2da11c9a68d1f7ff174cdb4303b83557e4235892cfa0c9"></a>

<a id="canonical-ea7aba9e9b195881e8c583c88161c00d7574d451f767a8f043c714103c789223"></a>

## create property — timeouts / e156b6b3d717 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f1cfca34f0eacbcad36d592bb652915d47b28f6047e77602b6652cb3559852fa"></a>

<a id="canonical-04a5d1f59549b85bbc46757cfd6a6ead029d7831c2d962c14d9625fb8eec5089"></a>

## delete property — timeouts / e156b6b3d717 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-8ef61c7160bc072d04738684e481f641febc9f4290980170659a9d89e472d5c4"></a>

<a id="canonical-51357fbf070cc7f83081f07050f772369f50cc857f1a7397710fcb9f471dcfad"></a>

## read property — timeouts / e156b6b3d717 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a7f2500133f5a5b2969c9e215242884b172abfddb89e0883c3bc701e6a69debc"></a>

<a id="canonical-cff14e6b32f94a8f43da20c00b48544ac9adba9a199ab157110f3876b2bf6cf1"></a>

## update property — timeouts / e156b6b3d717 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f9200e0f2ce220827f2516e5bb4a6cc9cc2ffba0a72491193db35ad24327462d"></a>

## Next pages — timeouts / e156b6b3d717 / 8

- [Property reference](resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
