---
page_title: "xcsh_api_testing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing reference."
---

# xcsh_api_testing reference

<a id="canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a0a32f74eb65e7f9737bd220e7f9e62372b48b1889a7b34b84bbb2b546e72e5"></a>

## Property reference — Property reference / 18f76cd140e6 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- Property reference

<a id="canonical-f75a2a0eb2d8dca73326e73cfea1b65ee30fc1280b050950cf71d628bf04e3f2"></a>

## Direct properties — Property reference / 18f76cd140e6 / 3

<a id="canonical-53209365466041be3ebb53d25861cee9dd6811e0c154640465ca82c8f975455a"></a>

<a id="canonical-f07f55ddbcbe1052291fd96ee72414908a59f80eb7361686cc0f3db15709b802"></a>

## annotations property — Property reference / 18f76cd140e6 / 4

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

<a id="canonical-e107922b9bb36d43dd0b2a058ad1f53207506becb0b7e63fc4c5f9406952339f"></a>

<a id="canonical-f6858149d3d18058cfe4e72eeb4ac9d6b5c60abd636cf99a5c87dbc78db5189c"></a>

## custom_header_value property — Property reference / 18f76cd140e6 / 5

Type: `"string"`. Computed.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

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

<a id="canonical-33766848693712b8d74b7f3a77a22d83075cefdaced6d7fb956379cbbdc876cd"></a>

<a id="canonical-08aefe94e167e7483e9a4252def4fcfe41fec1976dacf39682242523341e64a1"></a>

## description property — Property reference / 18f76cd140e6 / 6

Type: `"string"`. Computed.

Description of the APITesting.

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

- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006): complete subsection reference.

- [every_day](data-sources--api_testing--reference--group-001.md#canonical-bef0046bbe3fd9de7166290fd32fe2257e363fcf53500939061071a13739a1e6): complete subsection reference.

- [every_month](data-sources--api_testing--reference--group-001.md#canonical-1ad13f49880acec74448818804753c339eaf3e09b75d15858a684b11d51fc55e): complete subsection reference.

- [every_week](data-sources--api_testing--reference--group-001.md#canonical-5791b434f21e166201743305c608ca8fd93d0ee5060e234906071d1c9ffab591): complete subsection reference.

<a id="canonical-d00b18ca373a487e4bb605fb54e5f6645dd708c565e8ccacacd1a5eb95bdef17"></a>

<a id="canonical-4e13abb7c3a90de6eb00e7e2c13927aed45e0a8814f591e5beb3059beb0d4980"></a>

## id property — Property reference / 18f76cd140e6 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d2e0741c67469a304398444e4155ac6c8e9a6f305b52c52bb5bbc09fabb358cd"></a>

<a id="canonical-34e72efa882fa09ce32bcdc4433e144229ed59a4c15a0c91bedb4c16173d76ba"></a>

## labels property — Property reference / 18f76cd140e6 / 8

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

<a id="canonical-267ea68249ae3f983ca053cf2d237cfc65c3798aa2b90b79bb58b8ecb942b390"></a>

<a id="canonical-5881aa45c6912c65963a388bc480c9cd95dd12fd9776bb451d8d610d9768f98e"></a>

## name property — Property reference / 18f76cd140e6 / 9

Type: `"string"`. Required.

Name of the APITesting.

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

<a id="canonical-6dde5aa1a5f7d1f9a0e66f7a0cd6026e970a09c0b0bbc88960ff641782d117ae"></a>

<a id="canonical-9445d901d8cc7b72199d98020b3544d5f1b6e8a7421dc99ff0f6e83b6b7a939f"></a>

## namespace property — Property reference / 18f76cd140e6 / 10

Type: `"string"`. Required.

Namespace where the APITesting exists.

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

<a id="canonical-3dd026eddb7122c2eb305653a2df4926e48adafbc501383051dfaebc374669c1"></a>

## All schema paths — Property reference / 18f76cd140e6 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_testing--reference--group-001.md#canonical-53209365466041be3ebb53d25861cee9dd6811e0c154640465ca82c8f975455a) |
| `custom_header_value` | [custom_header_value](data-sources--api_testing--reference--group-001.md#canonical-e107922b9bb36d43dd0b2a058ad1f53207506becb0b7e63fc4c5f9406952339f) |
| `description` | [description](data-sources--api_testing--reference--group-001.md#canonical-33766848693712b8d74b7f3a77a22d83075cefdaced6d7fb956379cbbdc876cd) |
| `domains` | [domains](data-sources--api_testing--reference--group-001.md#canonical-4775e02a36d5d211ac2327feb21b01833aa33a2f9c427e3dbee727a42915c7d3) |
| `domains.allow_destructive_methods` | [domains.allow_destructive_methods](data-sources--api_testing--reference--group-001.md#canonical-9e2ddb8ffa5c4534daad0ee72f76cf65729e5378eee1c753ec51c6b75986a042) |
| `domains.credentials` | [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-97648a54a6e259729407374d4b7baed2c5f1cc703131b3f05d5611d71d4268aa) |
| `domains.credentials.admin` | [domains.credentials.admin](data-sources--api_testing--reference--group-001.md#canonical-cb4e15fc9a20b796a0f0439634ab8548c81e3b75ef54ef846042128be38d1cb9) |
| `domains.credentials.api_key` | [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-acf47378f207f6bd715792e8de99558f42ab32e01dc8982e5c91bbd8467513e9) |
| `domains.credentials.api_key.key` | [domains.credentials.api_key.key](data-sources--api_testing--reference--group-001.md#canonical-f88f7e2773f6413de263c06dbf570343cec23d1e4150ad8cc3c3449c4ad12258) |
| `domains.credentials.api_key.value` | [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-aec66fbcf6942d4e207c407a46a833c0395a67241ec0de8dece45082439cfa47) |
| `domains.credentials.api_key.value.blindfold_secret_info` | [domains.credentials.api_key.value.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-58122be3ef348ce91618590c975d6020ebc2ff3eff047ed5f14c32a629da97d7) |
| `domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` | [domains.credentials.api_key.value.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-89b26554abbe7a2d3c75ffffbd341c2ccc6a72730b25651bcc6ece798235aa58) |
| `domains.credentials.api_key.value.blindfold_secret_info.location` | [domains.credentials.api_key.value.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-a7aef8bc2dd99b7a4f23d4470580f57f1feeabb535beecf6aca68289bd1ce759) |
| `domains.credentials.api_key.value.blindfold_secret_info.store_provider` | [domains.credentials.api_key.value.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-9b68dbf288ba1d14ff62feea7d4bec9fc010d561c17cfa629559dda724501061) |
| `domains.credentials.api_key.value.clear_secret_info` | [domains.credentials.api_key.value.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-55ba0ac2f61066716e0f195a4284937eeef442ec1b125f7af49e84ac9171eae4) |
| `domains.credentials.api_key.value.clear_secret_info.provider_ref` | [domains.credentials.api_key.value.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-4c2bd685331d6613b29baf5ac8ce2fdcdc5e6dc9b831cc707b2b92bbfe6c528a) |
| `domains.credentials.api_key.value.clear_secret_info.url` | [domains.credentials.api_key.value.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-1f14136bc1c9fe42d982b09de23d710777d7d437ebeff704684c0e9a80a4d178) |
| `domains.credentials.basic_auth` | [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-347752d902a9daa73e3e8af6c9befad43251d675268af7977a7bd2e05442e9db) |
| `domains.credentials.basic_auth.password` | [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-cfe149e7672c76aa6cfedbe4be16ba9a894b82d378e8f782f4ca598fa84f6a6d) |
| `domains.credentials.basic_auth.password.blindfold_secret_info` | [domains.credentials.basic_auth.password.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-7a2afaf5de856f3ae7d33c80c3d3418916fef5a1e2651fba29158072998c50d4) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-e8fa1fff7ed6847f4cba9e7a8e87d9f7d281dfc598620ef5b3a4047f056e352e) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.location` | [domains.credentials.basic_auth.password.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-e435a79193f4f1817156c8ea11b313b6b8f17ea7d7a38a3a1b977d65df5e8efa) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-276aff44cb3c0aeb8f93e75e95895fb9b99be7d65625bc91ce3e4bec4dca70ef) |
| `domains.credentials.basic_auth.password.clear_secret_info` | [domains.credentials.basic_auth.password.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-90fcd383ce6d69488db7782663576ae06fb34e861c24cd746f7dcc3868b7930a) |
| `domains.credentials.basic_auth.password.clear_secret_info.provider_ref` | [domains.credentials.basic_auth.password.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-67ef00aa356d18cc9cdd365ca9a9825a4bea9f4ad09205d7ec1670c9ac9e6529) |
| `domains.credentials.basic_auth.password.clear_secret_info.url` | [domains.credentials.basic_auth.password.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-92d405c204197a5037b4cba23c2c1b6e55a7e7ff4208179b5cc66ed929506245) |
| `domains.credentials.basic_auth.user` | [domains.credentials.basic_auth.user](data-sources--api_testing--reference--group-001.md#canonical-fe118e8258b7e94f80b0ffecb437deb4d6a287bf45c4f121b3815af3433d8260) |
| `domains.credentials.bearer_token` | [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-ffc8ffb233a33aba085a2e828371c3308281cf29144250d14bd799cbb469c304) |
| `domains.credentials.bearer_token.token` | [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-488a8d6290d2a374afe26473e24eb3b757e8ea8614601be6e7a2c15971b9656a) |
| `domains.credentials.bearer_token.token.blindfold_secret_info` | [domains.credentials.bearer_token.token.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-895bc50829d1e8aa4842fa90cbb132c689cf8a25941829eb0b8ff766a00b5f6b) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-f443a955c985a7ccedbff820610d3ad56688cc67d23f4e33e2720b87bffc3daa) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.location` | [domains.credentials.bearer_token.token.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-60d0912f296663e867d3a86b80351aa998bfb18c9b9894f9b1970cf143369235) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-2e92a3a81cecf4ca00ee2e01c8ef6364a8de0aa02fdecea26b42654d834d3120) |
| `domains.credentials.bearer_token.token.clear_secret_info` | [domains.credentials.bearer_token.token.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-64dbe330d7a1bc0ceaf40e3da3f3a13ac7203f6a4a716cda53a03cb7ff5a2da0) |
| `domains.credentials.bearer_token.token.clear_secret_info.provider_ref` | [domains.credentials.bearer_token.token.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-cf6ca5cdb1f2eb34300495a5e1ef2c9d62d5df77ca00bb5676ca547adb7c886a) |
| `domains.credentials.bearer_token.token.clear_secret_info.url` | [domains.credentials.bearer_token.token.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-bfdc1cd630ca3cb90fd085401b145a40253b33452df2751208fa6facc41bf75d) |
| `domains.credentials.credential_name` | [domains.credentials.credential_name](data-sources--api_testing--reference--group-001.md#canonical-923a4f78a0339e1a7495ccc6d7396e3333be2b8954bb1bf248408f72215d12fe) |
| `domains.credentials.login_endpoint` | [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-76d2082432a3c25969046d19139f23c7699e7e855bdc27524f17a41cae9d07bb) |
| `domains.credentials.login_endpoint.json_payload` | [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-b368e647550d160547b700b6ee7dbff43ae01388b529e938a857f67a1c4d4357) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-006f71947036c2d56fbb94766dce8da15fd780c028c981f941a7bbdb4e175b3f) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-2ea618e088f1a1fca65ac2be103f546b489a3097a8efb2069401d3f6cdae98a2) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-a811dc068cf14615d50d084a1426d3b6f19932a8cacf401c57dbbbfcfd6769cd) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-c17ba9c602b148470f4a7f66d3868e8ef5e6b1eb0007ab1eb3b57bb1276cc2d1) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info` | [domains.credentials.login_endpoint.json_payload.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-5a156ef9e281f4e51df3adf952e28519eb9da42875a2e2181855e239bac60ced) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-b0d46b7cb2f76804c2cc363c6064b3e741e195faced77a4f6ea51c718f3b33b8) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.url` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-711cfa1a8bdff8537064777757ed6ab0a97ec3a0e79f3990b0c042b5d17c61ba) |
| `domains.credentials.login_endpoint.method` | [domains.credentials.login_endpoint.method](data-sources--api_testing--reference--group-001.md#canonical-3e3722b371c00e1e7996d5167674fa6de3aa0f842e1d02a03f176d6be9167667) |
| `domains.credentials.login_endpoint.path` | [domains.credentials.login_endpoint.path](data-sources--api_testing--reference--group-001.md#canonical-5ff3d1e543f341a83da36dde9bbd694d4fcd6d95bad0ee9d8c44f8b0aaf1fb30) |
| `domains.credentials.login_endpoint.token_response_key` | [domains.credentials.login_endpoint.token_response_key](data-sources--api_testing--reference--group-001.md#canonical-1ca535b73265ce25386078d025332ec958c93fed6d146aa85d28a1a627d4fb8a) |
| `domains.credentials.standard` | [domains.credentials.standard](data-sources--api_testing--reference--group-001.md#canonical-814bd195197a95f42db15f95e7e3b63e1c1d183c549d97acaed0cd33263a157a) |
| `domains.domain` | [domains.domain](data-sources--api_testing--reference--group-001.md#canonical-61050b46d79e30e7a1bbdb16e274f1ab26f52333d4bd0f2680d45c189a5a4111) |
| `every_day` | [every_day](data-sources--api_testing--reference--group-001.md#canonical-4259fe1d574ebfac934632b5dc810af17bd0bdf54c44cab7710c21561a3dfbb3) |
| `every_month` | [every_month](data-sources--api_testing--reference--group-001.md#canonical-bae3add6cd132f6b276d2de75d79ab1cbe63782cd2c27286b6f54c81c96b08f4) |
| `every_week` | [every_week](data-sources--api_testing--reference--group-001.md#canonical-1ac2e3c3805c0ecfdd81d3c61d3c4d8fb869aa8be8b9769c37a994973f940cbc) |
| `id` | [id](data-sources--api_testing--reference--group-001.md#canonical-d00b18ca373a487e4bb605fb54e5f6645dd708c565e8ccacacd1a5eb95bdef17) |
| `labels` | [labels](data-sources--api_testing--reference--group-001.md#canonical-d2e0741c67469a304398444e4155ac6c8e9a6f305b52c52bb5bbc09fabb358cd) |
| `name` | [name](data-sources--api_testing--reference--group-001.md#canonical-267ea68249ae3f983ca053cf2d237cfc65c3798aa2b90b79bb58b8ecb942b390) |
| `namespace` | [namespace](data-sources--api_testing--reference--group-001.md#canonical-6dde5aa1a5f7d1f9a0e66f7a0cd6026e970a09c0b0bbc88960ff641782d117ae) |

<a id="canonical-70521ec7e277a13dda8d85ee26924ee3884aedb0007478a5833c9246ca01ee9e"></a>

## Next pages — Property reference / 18f76cd140e6 / 12

- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [every_day](data-sources--api_testing--reference--group-001.md#canonical-bef0046bbe3fd9de7166290fd32fe2257e363fcf53500939061071a13739a1e6)
- [every_month](data-sources--api_testing--reference--group-001.md#canonical-1ad13f49880acec74448818804753c339eaf3e09b75d15858a684b11d51fc55e)
- [every_week](data-sources--api_testing--reference--group-001.md#canonical-5791b434f21e166201743305c608ca8fd93d0ee5060e234906071d1c9ffab591)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e191a533a33b635f51faf0a4371c14bf2167b33e93b9b40dfd9df93d37ee2c2"></a>

## domains — domains / 048a855a99e1 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- domains

<a id="canonical-4775e02a36d5d211ac2327feb21b01833aa33a2f9c427e3dbee727a42915c7d3"></a>

Type: `"list"`. Computed.

Add and configure testing domains and credentials.

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

<a id="canonical-1f85d195c90cf672a9999f891bf756bb2343df809f502b4949732c9bfb05118e"></a>

## Direct properties — domains / 048a855a99e1 / 3

<a id="canonical-9e2ddb8ffa5c4534daad0ee72f76cf65729e5378eee1c753ec51c6b75986a042"></a>

<a id="canonical-aa83908964a573ca4b72c3b14891faa8e41f3f7f053e120df6dad50dcb3ab14f"></a>

## allow_destructive_methods property — domains / 048a855a99e1 / 4

Type: `"bool"`. Computed.

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

- [credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16): complete subsection reference.

<a id="canonical-61050b46d79e30e7a1bbdb16e274f1ab26f52333d4bd0f2680d45c189a5a4111"></a>

<a id="canonical-4406801b590f8794348dc3822669fdfe384682ae6197cdf0c9c99bec11c9247f"></a>

## domain property — domains / 048a855a99e1 / 5

Type: `"string"`. Computed.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

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

<a id="canonical-11489be0bcba4e5e8596d6d2435f391e63183eaf82e2ba281398ced25803f2cd"></a>

## Next pages — domains / 048a855a99e1 / 6

- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6535ac57306639de72faed44d956ebad97670bed8591943ba9770a26e18b39c4"></a>

## domains.credentials — domains.credentials / d9654f121690 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- domains.credentials

<a id="canonical-97648a54a6e259729407374d4b7baed2c5f1cc703131b3f05d5611d71d4268aa"></a>

Type: `"list"`. Computed.

Add credentials for API testing to use in the selected environment.

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

<a id="canonical-aa3f874ee17e368a4a1aa70fc6e70d0686f96f44651e65e3dfa35d27ba01043a"></a>

## Direct properties — domains.credentials / d9654f121690 / 3

- [admin](data-sources--api_testing--reference--group-001.md#canonical-7b09d5e96ca7aae8728ddd699a7e10fa05efc7bb75e53e6d112b71cc3e78a5a2): complete subsection reference.

- [api_key](data-sources--api_testing--reference--group-001.md#canonical-bb18e2f78b56c65bc61907f9d449d5c6761bd04bd2e255099b9d70b15ad83c8b): complete subsection reference.

- [basic_auth](data-sources--api_testing--reference--group-001.md#canonical-aa193809b175f75541433c2376e607e356bcc1744f714132083d7d7b53cb4bda): complete subsection reference.

- [bearer_token](data-sources--api_testing--reference--group-001.md#canonical-42c1f538fe9382495545dc4b2d9c6cb81a8b208524924917ce08f2754d87cac7): complete subsection reference.

<a id="canonical-923a4f78a0339e1a7495ccc6d7396e3333be2b8954bb1bf248408f72215d12fe"></a>

<a id="canonical-a1ad2af42d8a666b08b3f9bab18f376aa9d14de47d1984405cbb05f2b9e5fc03"></a>

## credential_name property — domains.credentials / d9654f121690 / 4

Type: `"string"`. Computed.

Enter a unique name for the credentials used in API testing.

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

- [login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-2957691b8432fafac009d886cf355cf706257dce6e55943c9eb25f56ad52668c): complete subsection reference.

- [standard](data-sources--api_testing--reference--group-001.md#canonical-c230dbede3e63935fb094f7700976ae8bd6b622e30d85cc9b74808729a4304c0): complete subsection reference.

<a id="canonical-09107ef2ce7601da65611e1098dbf967a9d4df376f1a796b964431fe9bfac81c"></a>

## Next pages — domains.credentials / d9654f121690 / 5

- [domains.credentials.admin](data-sources--api_testing--reference--group-001.md#canonical-7b09d5e96ca7aae8728ddd699a7e10fa05efc7bb75e53e6d112b71cc3e78a5a2)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-bb18e2f78b56c65bc61907f9d449d5c6761bd04bd2e255099b9d70b15ad83c8b)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-aa193809b175f75541433c2376e607e356bcc1744f714132083d7d7b53cb4bda)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-42c1f538fe9382495545dc4b2d9c6cb81a8b208524924917ce08f2754d87cac7)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-2957691b8432fafac009d886cf355cf706257dce6e55943c9eb25f56ad52668c)
- [domains.credentials.standard](data-sources--api_testing--reference--group-001.md#canonical-c230dbede3e63935fb094f7700976ae8bd6b622e30d85cc9b74808729a4304c0)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-7b09d5e96ca7aae8728ddd699a7e10fa05efc7bb75e53e6d112b71cc3e78a5a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2ed68e4c8908573e30f5e3265c77ebaf99dfa463b241e4ca706f94737ef94c6"></a>

## domains.credentials.admin — domains.credentials.admin / c00c8d2c3a4c / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- domains.credentials.admin

<a id="canonical-cb4e15fc9a20b796a0f0439634ab8548c81e3b75ef54ef846042128be38d1cb9"></a>

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

<a id="canonical-94491df3f6616eebe103a7726436551c191551f61d1d589079d73ec8441421b7"></a>

## Direct properties — domains.credentials.admin / c00c8d2c3a4c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51d5e87e3d623d22ab24aeb701665a012a550d67e7ca27721dc7548b8db9d91a"></a>

## Next pages — domains.credentials.admin / c00c8d2c3a4c / 4

- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-bb18e2f78b56c65bc61907f9d449d5c6761bd04bd2e255099b9d70b15ad83c8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31f1c3be6cc9542c8a5536a4a694d99e8a02f6aa4e362c7aee2590075f0e432f"></a>

## domains.credentials.api_key — domains.credentials.api_key / fc0743371779 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- domains.credentials.api_key

<a id="canonical-acf47378f207f6bd715792e8de99558f42ab32e01dc8982e5c91bbd8467513e9"></a>

Type: `"single"`. Computed.

API Key

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

<a id="canonical-fb345d749b11572b21176d2e165619812ba77ed895f5c632968a13d4e48ddcd6"></a>

## Direct properties — domains.credentials.api_key / fc0743371779 / 3

<a id="canonical-f88f7e2773f6413de263c06dbf570343cec23d1e4150ad8cc3c3449c4ad12258"></a>

<a id="canonical-8a683cdb265970ffa0334a0e148fe74e643a8e2ae7c572faa5558d5ba5670012"></a>

## key property — domains.credentials.api_key / fc0743371779 / 4

Type: `"string"`. Computed.

Key. Cryptographic key material

Upstream description:

Cryptographic key material

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

- [value](data-sources--api_testing--reference--group-001.md#canonical-8f7e9253bb51aa8813c086d7adcb907952e08b577e54fd52e1cb56e2d1db3bb8): complete subsection reference.

<a id="canonical-fa975690e1b4ad9eb5490faab2c414429d26dda7e90ed63229f7d0a5e4b1ed68"></a>

## Next pages — domains.credentials.api_key / fc0743371779 / 5

- [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-8f7e9253bb51aa8813c086d7adcb907952e08b577e54fd52e1cb56e2d1db3bb8)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-8f7e9253bb51aa8813c086d7adcb907952e08b577e54fd52e1cb56e2d1db3bb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a42a40d59ea93fa66a4e872469ed75be2aa16a08bebeef62a7bbb02f72fac8f7"></a>

## domains.credentials.api_key.value — domains.credentials.api_key.value / 909ea6e5f59c / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-bb18e2f78b56c65bc61907f9d449d5c6761bd04bd2e255099b9d70b15ad83c8b)
- domains.credentials.api_key.value

<a id="canonical-aec66fbcf6942d4e207c407a46a833c0395a67241ec0de8dece45082439cfa47"></a>

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

<a id="canonical-8c262b689997b2f6f762989d181777d71db2a7b843817849edbb34f4538c853f"></a>

## Direct properties — domains.credentials.api_key.value / 909ea6e5f59c / 3

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-e25e6d716fb6da26123cd906fe7a1b66cfe3dc854acf6e9e5939a8c3b38e2b65): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-77e3c4081e9e22461f07a8b9aacdb85eca8b99196a37be39d9de5f9b21d3b012): complete subsection reference.

<a id="canonical-291534b2b8e6b095538fadb53bfcf956f116bae6c89399a25e40bb28e2cf3ee7"></a>

## Next pages — domains.credentials.api_key.value / 909ea6e5f59c / 4

- [domains.credentials.api_key.value.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-e25e6d716fb6da26123cd906fe7a1b66cfe3dc854acf6e9e5939a8c3b38e2b65)
- [domains.credentials.api_key.value.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-77e3c4081e9e22461f07a8b9aacdb85eca8b99196a37be39d9de5f9b21d3b012)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-bb18e2f78b56c65bc61907f9d449d5c6761bd04bd2e255099b9d70b15ad83c8b)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-e25e6d716fb6da26123cd906fe7a1b66cfe3dc854acf6e9e5939a8c3b38e2b65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43d15ba1cb5a965e83a9f23c0ed141dd52cb6e349e9e4e188ecfae3d89d87792"></a>

## domains.credentials.api_key.value.blindfold_secret_info — domains.credentials.api_key.value.blindfold_secret_info / a72a9c170349 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-bb18e2f78b56c65bc61907f9d449d5c6761bd04bd2e255099b9d70b15ad83c8b)
- [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-8f7e9253bb51aa8813c086d7adcb907952e08b577e54fd52e1cb56e2d1db3bb8)
- domains.credentials.api_key.value.blindfold_secret_info

<a id="canonical-58122be3ef348ce91618590c975d6020ebc2ff3eff047ed5f14c32a629da97d7"></a>

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

<a id="canonical-e8f4222d24231b2b722613a6e1c267be1e4d4fe6d7cac4913539ec01bfbab9b1"></a>

## Direct properties — domains.credentials.api_key.value.blindfold_secret_info / a72a9c170349 / 3

<a id="canonical-89b26554abbe7a2d3c75ffffbd341c2ccc6a72730b25651bcc6ece798235aa58"></a>

<a id="canonical-a3f55d645ca84443f9579b0352947624a50970fb575399e8a745ba151f3738b2"></a>

## decryption_provider property — domains.credentials.api_key.value.blindfold_secret_info / a72a9c170349 / 4

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

<a id="canonical-a7aef8bc2dd99b7a4f23d4470580f57f1feeabb535beecf6aca68289bd1ce759"></a>

<a id="canonical-b41676d393678c406393aa3b2d983beda6e366e4f1ecc68122283edb50cd1bf0"></a>

## location property — domains.credentials.api_key.value.blindfold_secret_info / a72a9c170349 / 5

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

<a id="canonical-9b68dbf288ba1d14ff62feea7d4bec9fc010d561c17cfa629559dda724501061"></a>

<a id="canonical-9cbadb262447aa6c9d3ab96c6cb26884208dcc449195437accea6a6039353e52"></a>

## store_provider property — domains.credentials.api_key.value.blindfold_secret_info / a72a9c170349 / 6

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

<a id="canonical-9e3404ad26ba3762f7245383a62847ce58d22145fd6147898d39f1ba08b84c45"></a>

## Next pages — domains.credentials.api_key.value.blindfold_secret_info / a72a9c170349 / 7

- [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-8f7e9253bb51aa8813c086d7adcb907952e08b577e54fd52e1cb56e2d1db3bb8)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-77e3c4081e9e22461f07a8b9aacdb85eca8b99196a37be39d9de5f9b21d3b012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41142632f5a50e70719f330e6c8ebc4ee047c2d02e261aa62a3e1f3e349e242c"></a>

## domains.credentials.api_key.value.clear_secret_info — domains.credentials.api_key.value.clear_secret_info / 8a4f8f625a8b / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-bb18e2f78b56c65bc61907f9d449d5c6761bd04bd2e255099b9d70b15ad83c8b)
- [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-8f7e9253bb51aa8813c086d7adcb907952e08b577e54fd52e1cb56e2d1db3bb8)
- domains.credentials.api_key.value.clear_secret_info

<a id="canonical-55ba0ac2f61066716e0f195a4284937eeef442ec1b125f7af49e84ac9171eae4"></a>

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

<a id="canonical-bced9398ad05ab23592fd3d8d056e55af7a62751b069ef2a753d9e9870866c41"></a>

## Direct properties — domains.credentials.api_key.value.clear_secret_info / 8a4f8f625a8b / 3

<a id="canonical-4c2bd685331d6613b29baf5ac8ce2fdcdc5e6dc9b831cc707b2b92bbfe6c528a"></a>

<a id="canonical-3398fdd5b65c288a6e88ec89b9678853d44a4b267d7b787beda0b917e2277ee9"></a>

## provider_ref property — domains.credentials.api_key.value.clear_secret_info / 8a4f8f625a8b / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1f14136bc1c9fe42d982b09de23d710777d7d437ebeff704684c0e9a80a4d178"></a>

<a id="canonical-a815b7bab69fd5e1e028af897040ced6caff41db3069130c80723ca6b5b64157"></a>

## url property — domains.credentials.api_key.value.clear_secret_info / 8a4f8f625a8b / 5

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

<a id="canonical-e877def43dbbd82344450c2983510b1f1264f878422dd0e9ff35ed2a6b76ff19"></a>

## Next pages — domains.credentials.api_key.value.clear_secret_info / 8a4f8f625a8b / 6

- [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-8f7e9253bb51aa8813c086d7adcb907952e08b577e54fd52e1cb56e2d1db3bb8)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-aa193809b175f75541433c2376e607e356bcc1744f714132083d7d7b53cb4bda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a8d02275941d6570ca137ec043498dba913b51c97b06682035703823db19c26"></a>

## domains.credentials.basic_auth — domains.credentials.basic_auth / 2b344b4e0485 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- domains.credentials.basic_auth

<a id="canonical-347752d902a9daa73e3e8af6c9befad43251d675268af7977a7bd2e05442e9db"></a>

Type: `"single"`. Computed.

Basic Authentication.

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

<a id="canonical-47a41864509c0768b546961a993b5e4c666b5cfa96562bae72c528d94a6faca2"></a>

## Direct properties — domains.credentials.basic_auth / 2b344b4e0485 / 3

- [password](data-sources--api_testing--reference--group-001.md#canonical-540ce2da8085482c38448b0fb51138b309b3fdf9d78336632733f2e25e24514d): complete subsection reference.

<a id="canonical-fe118e8258b7e94f80b0ffecb437deb4d6a287bf45c4f121b3815af3433d8260"></a>

<a id="canonical-855bd0226e5e748d809a204e5a7a590728c692b1ce41d19221b597b60bd0cfff"></a>

## user property — domains.credentials.basic_auth / 2b344b4e0485 / 4

Type: `"string"`. Computed.

User. Configuration parameter for user

Upstream description:

Configuration parameter for user

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

<a id="canonical-2685b7085c3091dd6fb7a71d09c8f526d299c86652a6940cb00fc64f73d8508b"></a>

## Next pages — domains.credentials.basic_auth / 2b344b4e0485 / 5

- [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-540ce2da8085482c38448b0fb51138b309b3fdf9d78336632733f2e25e24514d)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-540ce2da8085482c38448b0fb51138b309b3fdf9d78336632733f2e25e24514d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87fb54bf3377cfaabcfca79d6302ec16640d7a1eed259316d497b3117dd30e9e"></a>

## domains.credentials.basic_auth.password — domains.credentials.basic_auth.password / 0b624c5cdb61 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-aa193809b175f75541433c2376e607e356bcc1744f714132083d7d7b53cb4bda)
- domains.credentials.basic_auth.password

<a id="canonical-cfe149e7672c76aa6cfedbe4be16ba9a894b82d378e8f782f4ca598fa84f6a6d"></a>

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

<a id="canonical-d1a81e2431776820971052656e69a885e31412b007b3ed7cf31872580db1647f"></a>

## Direct properties — domains.credentials.basic_auth.password / 0b624c5cdb61 / 3

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-728998777fee1aeaff0869f98175ab0e401ecf3622bfd2c93fa84e91b068a76a): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-5d373485e2c49be30053f9bd2b8484fa72b80a3ede7d6788c4162332f3da69da): complete subsection reference.

<a id="canonical-4c8675f0e589afdc738111984b7934352e8253126d06425df015d1a98877a45f"></a>

## Next pages — domains.credentials.basic_auth.password / 0b624c5cdb61 / 4

- [domains.credentials.basic_auth.password.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-728998777fee1aeaff0869f98175ab0e401ecf3622bfd2c93fa84e91b068a76a)
- [domains.credentials.basic_auth.password.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-5d373485e2c49be30053f9bd2b8484fa72b80a3ede7d6788c4162332f3da69da)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-aa193809b175f75541433c2376e607e356bcc1744f714132083d7d7b53cb4bda)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-728998777fee1aeaff0869f98175ab0e401ecf3622bfd2c93fa84e91b068a76a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ed5f4539901d37b0700832441826560f0c2a5ffcd3f5ed9ef7808bc2e62cee2"></a>

## domains.credentials.basic_auth.password.blindfold_secret_info — domains.credentials.basic_auth.password.blindfold_secret_info / 78314846bc10 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-aa193809b175f75541433c2376e607e356bcc1744f714132083d7d7b53cb4bda)
- [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-540ce2da8085482c38448b0fb51138b309b3fdf9d78336632733f2e25e24514d)
- domains.credentials.basic_auth.password.blindfold_secret_info

<a id="canonical-7a2afaf5de856f3ae7d33c80c3d3418916fef5a1e2651fba29158072998c50d4"></a>

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

<a id="canonical-95f08512cc0c145370c1f4453a39a576616d527fdff6ebd493f25888013defc6"></a>

## Direct properties — domains.credentials.basic_auth.password.blindfold_secret_info / 78314846bc10 / 3

<a id="canonical-e8fa1fff7ed6847f4cba9e7a8e87d9f7d281dfc598620ef5b3a4047f056e352e"></a>

<a id="canonical-39cadbead421ceee47fafe1f234821bf1d7627ac1f68bee55e461d8ad08b78ca"></a>

## decryption_provider property — domains.credentials.basic_auth.password.blindfold_secret_info / 78314846bc10 / 4

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

<a id="canonical-e435a79193f4f1817156c8ea11b313b6b8f17ea7d7a38a3a1b977d65df5e8efa"></a>

<a id="canonical-46bcf6d9aac04e9f19d1db9ba86e4fa4681f4d09958a1622ac72f6725541afa2"></a>

## location property — domains.credentials.basic_auth.password.blindfold_secret_info / 78314846bc10 / 5

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

<a id="canonical-276aff44cb3c0aeb8f93e75e95895fb9b99be7d65625bc91ce3e4bec4dca70ef"></a>

<a id="canonical-6507801c4f513dd9a37b977b374a62e1b3601a48f4d3960bd42d0f7224dbb22d"></a>

## store_provider property — domains.credentials.basic_auth.password.blindfold_secret_info / 78314846bc10 / 6

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

<a id="canonical-2c5c70071c43b73a70cc8d7f4c2d1a32473709f974bea6c79af665b9d5025819"></a>

## Next pages — domains.credentials.basic_auth.password.blindfold_secret_info / 78314846bc10 / 7

- [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-540ce2da8085482c38448b0fb51138b309b3fdf9d78336632733f2e25e24514d)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-5d373485e2c49be30053f9bd2b8484fa72b80a3ede7d6788c4162332f3da69da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-739cf09f119df621c734de0e552878ff91f79e1a9e9e42fcf9b3deb7e264fdda"></a>

## domains.credentials.basic_auth.password.clear_secret_info — domains.credentials.basic_auth.password.clear_secret_info / 678a88138f9f / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-aa193809b175f75541433c2376e607e356bcc1744f714132083d7d7b53cb4bda)
- [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-540ce2da8085482c38448b0fb51138b309b3fdf9d78336632733f2e25e24514d)
- domains.credentials.basic_auth.password.clear_secret_info

<a id="canonical-90fcd383ce6d69488db7782663576ae06fb34e861c24cd746f7dcc3868b7930a"></a>

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

<a id="canonical-9c88d9de477f8450ce147640f3cb29755376f3b3ed9e0bf10ad8379fa4e40aac"></a>

## Direct properties — domains.credentials.basic_auth.password.clear_secret_info / 678a88138f9f / 3

<a id="canonical-67ef00aa356d18cc9cdd365ca9a9825a4bea9f4ad09205d7ec1670c9ac9e6529"></a>

<a id="canonical-0a59d80aa7d1711050cdd9772f4eada51a2c05d4cae764688bbbbf7096703088"></a>

## provider_ref property — domains.credentials.basic_auth.password.clear_secret_info / 678a88138f9f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-92d405c204197a5037b4cba23c2c1b6e55a7e7ff4208179b5cc66ed929506245"></a>

<a id="canonical-d239142d894be05bea7d7ac94f58f26447537d9a26894b6b9a9ae7f59426bfa7"></a>

## url property — domains.credentials.basic_auth.password.clear_secret_info / 678a88138f9f / 5

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

<a id="canonical-a155ec99fb1aab54d6464cc4b684086a8662f776168ed48c1e6bf328dda396c3"></a>

## Next pages — domains.credentials.basic_auth.password.clear_secret_info / 678a88138f9f / 6

- [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-540ce2da8085482c38448b0fb51138b309b3fdf9d78336632733f2e25e24514d)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-42c1f538fe9382495545dc4b2d9c6cb81a8b208524924917ce08f2754d87cac7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f254a6c4b927e56bf960d9f380692dc4b97a2340eb95c5da02b6a7ee022e073b"></a>

## domains.credentials.bearer_token — domains.credentials.bearer_token / 8429906f1960 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- domains.credentials.bearer_token

<a id="canonical-ffc8ffb233a33aba085a2e828371c3308281cf29144250d14bd799cbb469c304"></a>

Type: `"single"`. Computed.

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

<a id="canonical-b4acddc1089535ca9619f7d9f7f24348bee5b7200563746914c07441565a15e0"></a>

## Direct properties — domains.credentials.bearer_token / 8429906f1960 / 3

- [token](data-sources--api_testing--reference--group-001.md#canonical-b9c0315dfa1e94130d291735d1634b820552f4625ebffb808c726180ca1dbb1b): complete subsection reference.

<a id="canonical-226cc280313ed21ca3096999d4de7c15c9e012836a7bda7e306807a2386c1a2c"></a>

## Next pages — domains.credentials.bearer_token / 8429906f1960 / 4

- [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-b9c0315dfa1e94130d291735d1634b820552f4625ebffb808c726180ca1dbb1b)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-b9c0315dfa1e94130d291735d1634b820552f4625ebffb808c726180ca1dbb1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4387c6df05937330673c078a12f98682ccd490aa1a6eda6597f34390daf3b3c"></a>

## domains.credentials.bearer_token.token — domains.credentials.bearer_token.token / 28f309ab8ab3 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-42c1f538fe9382495545dc4b2d9c6cb81a8b208524924917ce08f2754d87cac7)
- domains.credentials.bearer_token.token

<a id="canonical-488a8d6290d2a374afe26473e24eb3b757e8ea8614601be6e7a2c15971b9656a"></a>

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

<a id="canonical-640d4409f5c38c7953324c4abfb5936e99801524273c220be7c328b4075533ab"></a>

## Direct properties — domains.credentials.bearer_token.token / 28f309ab8ab3 / 3

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-4e7d0b05e0f947a218ccb1a74e7e84a89760c9241cc042fd87253c6c5ed43ae8): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-c9b7622b7f52d8bec4b82605686be386e578a7b57b312579cffcbb9f8a70f6f3): complete subsection reference.

<a id="canonical-4b33f10e3708130338ef011c0e9b751a96029124476b0b3ee8eff20fdf57629e"></a>

## Next pages — domains.credentials.bearer_token.token / 28f309ab8ab3 / 4

- [domains.credentials.bearer_token.token.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-4e7d0b05e0f947a218ccb1a74e7e84a89760c9241cc042fd87253c6c5ed43ae8)
- [domains.credentials.bearer_token.token.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-c9b7622b7f52d8bec4b82605686be386e578a7b57b312579cffcbb9f8a70f6f3)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-42c1f538fe9382495545dc4b2d9c6cb81a8b208524924917ce08f2754d87cac7)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-4e7d0b05e0f947a218ccb1a74e7e84a89760c9241cc042fd87253c6c5ed43ae8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5578a4e3ec1438e5e896665bdb20d5a22a20b0cd647bba8a47332280ce6ae2cd"></a>

## domains.credentials.bearer_token.token.blindfold_secret_info — domains.credentials.bearer_token.token.blindfold_secret_info / 791ab7a540cf / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-42c1f538fe9382495545dc4b2d9c6cb81a8b208524924917ce08f2754d87cac7)
- [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-b9c0315dfa1e94130d291735d1634b820552f4625ebffb808c726180ca1dbb1b)
- domains.credentials.bearer_token.token.blindfold_secret_info

<a id="canonical-895bc50829d1e8aa4842fa90cbb132c689cf8a25941829eb0b8ff766a00b5f6b"></a>

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

<a id="canonical-33b993f02bc1ffc52531582706ca981157333a29f33a0f6fa78b37a67f6eedd8"></a>

## Direct properties — domains.credentials.bearer_token.token.blindfold_secret_info / 791ab7a540cf / 3

<a id="canonical-f443a955c985a7ccedbff820610d3ad56688cc67d23f4e33e2720b87bffc3daa"></a>

<a id="canonical-158f1591e6a36ee7b68da4cb2e5462d323d3f5d1cefe04b1cecc0aac9a25d7fb"></a>

## decryption_provider property — domains.credentials.bearer_token.token.blindfold_secret_info / 791ab7a540cf / 4

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

<a id="canonical-60d0912f296663e867d3a86b80351aa998bfb18c9b9894f9b1970cf143369235"></a>

<a id="canonical-f9b78647c724fe7ffdc844600b30e8dbd58389953e78c748d71a821b7a1b93f5"></a>

## location property — domains.credentials.bearer_token.token.blindfold_secret_info / 791ab7a540cf / 5

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

<a id="canonical-2e92a3a81cecf4ca00ee2e01c8ef6364a8de0aa02fdecea26b42654d834d3120"></a>

<a id="canonical-cf92f3b4aeaabb52f44a27be771c3583e3e6f3966c2bcc255db9f122746c4ba5"></a>

## store_provider property — domains.credentials.bearer_token.token.blindfold_secret_info / 791ab7a540cf / 6

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

<a id="canonical-06bc81b478fb1977af980d52d08e60a8d59c9225352c6d78381c7a82bd874058"></a>

## Next pages — domains.credentials.bearer_token.token.blindfold_secret_info / 791ab7a540cf / 7

- [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-b9c0315dfa1e94130d291735d1634b820552f4625ebffb808c726180ca1dbb1b)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-c9b7622b7f52d8bec4b82605686be386e578a7b57b312579cffcbb9f8a70f6f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9691975961e23e39319ca32b98d8084339333bac4dd7991b7ce3b384131a98ab"></a>

## domains.credentials.bearer_token.token.clear_secret_info — domains.credentials.bearer_token.token.clear_secret_info / 7ca478bd374a / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-42c1f538fe9382495545dc4b2d9c6cb81a8b208524924917ce08f2754d87cac7)
- [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-b9c0315dfa1e94130d291735d1634b820552f4625ebffb808c726180ca1dbb1b)
- domains.credentials.bearer_token.token.clear_secret_info

<a id="canonical-64dbe330d7a1bc0ceaf40e3da3f3a13ac7203f6a4a716cda53a03cb7ff5a2da0"></a>

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

<a id="canonical-f91f22b1d43bb83792e309b5ed7d6094efe72dbf9165cca665a07630d1cc2942"></a>

## Direct properties — domains.credentials.bearer_token.token.clear_secret_info / 7ca478bd374a / 3

<a id="canonical-cf6ca5cdb1f2eb34300495a5e1ef2c9d62d5df77ca00bb5676ca547adb7c886a"></a>

<a id="canonical-359fc62bdb7f369e055560fd70fb28d1751ea5354ad2fef202b14b5e41c66cb7"></a>

## provider_ref property — domains.credentials.bearer_token.token.clear_secret_info / 7ca478bd374a / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-bfdc1cd630ca3cb90fd085401b145a40253b33452df2751208fa6facc41bf75d"></a>

<a id="canonical-926da5cce566d3c3b07a929d25a4d257ef3c98c81b9a15a0ef7767703cd86b4e"></a>

## url property — domains.credentials.bearer_token.token.clear_secret_info / 7ca478bd374a / 5

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

<a id="canonical-840eb6be840559ce96d645683c2d4c96d3994ae0062837044d37476b8f5a4f8b"></a>

## Next pages — domains.credentials.bearer_token.token.clear_secret_info / 7ca478bd374a / 6

- [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-b9c0315dfa1e94130d291735d1634b820552f4625ebffb808c726180ca1dbb1b)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-2957691b8432fafac009d886cf355cf706257dce6e55943c9eb25f56ad52668c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02b75203f2b7ef58a01d1c8133cf2b49e40638835335ded47d996eac71be7b50"></a>

## domains.credentials.login_endpoint — domains.credentials.login_endpoint / 046e5f81caf3 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- domains.credentials.login_endpoint

<a id="canonical-76d2082432a3c25969046d19139f23c7699e7e855bdc27524f17a41cae9d07bb"></a>

Type: `"single"`. Computed.

Login Endpoint.

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

<a id="canonical-1f30e51577c30e0d944e643f4c546ba007da200575913f57d33d398bd9cf6b76"></a>

## Direct properties — domains.credentials.login_endpoint / 046e5f81caf3 / 3

- [json_payload](data-sources--api_testing--reference--group-001.md#canonical-4e213a146be541001c0e10dd21caa4a63a3f4bae702911cb8b098b7ce54eee5f): complete subsection reference.

<a id="canonical-3e3722b371c00e1e7996d5167674fa6de3aa0f842e1d02a03f176d6be9167667"></a>

<a id="canonical-63e352c797b07185819bc97743608464c8e7f579584aae749a0957a4fe0d4c31"></a>

## method property — domains.credentials.login_endpoint / 046e5f81caf3 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

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

<a id="canonical-5ff3d1e543f341a83da36dde9bbd694d4fcd6d95bad0ee9d8c44f8b0aaf1fb30"></a>

<a id="canonical-166ce0f29edc708a687fed085c30d0335ee2f990242e85d993a7c4c796d198f1"></a>

## path property — domains.credentials.login_endpoint / 046e5f81caf3 / 5

Type: `"string"`. Computed.

Path. URL path for the endpoint

Upstream description:

URL path for the endpoint

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

<a id="canonical-1ca535b73265ce25386078d025332ec958c93fed6d146aa85d28a1a627d4fb8a"></a>

<a id="canonical-ae484a9d126e61b83f81979bfd13493a3bb5135a98370b24c37501b017554c89"></a>

## token_response_key property — domains.credentials.login_endpoint / 046e5f81caf3 / 6

Type: `"string"`. Computed.

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

<a id="canonical-a781bf9ada24cb84271fbf3b6a6bb4ba2ac5f19082f277d2f8af48ce5e3ab767"></a>

## Next pages — domains.credentials.login_endpoint / 046e5f81caf3 / 7

- [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-4e213a146be541001c0e10dd21caa4a63a3f4bae702911cb8b098b7ce54eee5f)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-4e213a146be541001c0e10dd21caa4a63a3f4bae702911cb8b098b7ce54eee5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fda0099edac2a385d2731aafebb0907419e26f5ccb3f524bbb70edc0ec94061"></a>

## domains.credentials.login_endpoint.json_payload — domains.credentials.login_endpoint.json_payload / 9eeb28ce47cc / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-2957691b8432fafac009d886cf355cf706257dce6e55943c9eb25f56ad52668c)
- domains.credentials.login_endpoint.json_payload

<a id="canonical-b368e647550d160547b700b6ee7dbff43ae01388b529e938a857f67a1c4d4357"></a>

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

<a id="canonical-824058d7a918bd299e54ee40f5000502ff6148a9654edb38b9e58fb595a76577"></a>

## Direct properties — domains.credentials.login_endpoint.json_payload / 9eeb28ce47cc / 3

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-fda5e4062e48ca7a1f3c3643f4100b615e0f646e993a828552c26a804da9318f): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-8da733b8cee0cf4397df033d053756848c476d2340cc23c96e0545d8c23e28a4): complete subsection reference.

<a id="canonical-61827bced9893e78cb7f13598a97874183cedaa21027bb52029972ecd2953040"></a>

## Next pages — domains.credentials.login_endpoint.json_payload / 9eeb28ce47cc / 4

- [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-fda5e4062e48ca7a1f3c3643f4100b615e0f646e993a828552c26a804da9318f)
- [domains.credentials.login_endpoint.json_payload.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-8da733b8cee0cf4397df033d053756848c476d2340cc23c96e0545d8c23e28a4)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-2957691b8432fafac009d886cf355cf706257dce6e55943c9eb25f56ad52668c)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-fda5e4062e48ca7a1f3c3643f4100b615e0f646e993a828552c26a804da9318f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a34d71071dc679e38211e8307b22c94d3c54211ddc817a14f1feabe324ff75a3"></a>

## domains.credentials.login_endpoint.json_payload.blindfold_secret_info — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / 353ab21936d3 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-2957691b8432fafac009d886cf355cf706257dce6e55943c9eb25f56ad52668c)
- [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-4e213a146be541001c0e10dd21caa4a63a3f4bae702911cb8b098b7ce54eee5f)
- domains.credentials.login_endpoint.json_payload.blindfold_secret_info

<a id="canonical-006f71947036c2d56fbb94766dce8da15fd780c028c981f941a7bbdb4e175b3f"></a>

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

<a id="canonical-9a2b0f95dbaeb2c315f2b28d25fa53d47dd4b28884c2d3c7f47582d647704356"></a>

## Direct properties — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / 353ab21936d3 / 3

<a id="canonical-2ea618e088f1a1fca65ac2be103f546b489a3097a8efb2069401d3f6cdae98a2"></a>

<a id="canonical-884cc5668d8958ef35e0b3bdd79aabdadd63a93db9e4420c6e02ba070ed45b1f"></a>

## decryption_provider property — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / 353ab21936d3 / 4

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

<a id="canonical-a811dc068cf14615d50d084a1426d3b6f19932a8cacf401c57dbbbfcfd6769cd"></a>

<a id="canonical-924850b4877ccfdd46fae7943155e90783bd754278012c3dfa22a1f90335fd8f"></a>

## location property — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / 353ab21936d3 / 5

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

<a id="canonical-c17ba9c602b148470f4a7f66d3868e8ef5e6b1eb0007ab1eb3b57bb1276cc2d1"></a>

<a id="canonical-5ebb513cdd137e9bcb04009e4949fcbd550b5a660509646cf9c4d4e17f127c03"></a>

## store_provider property — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / 353ab21936d3 / 6

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

<a id="canonical-da91641c48679b5194a5a8857daeb13e2b8d112283032b37a1ce957b777e4080"></a>

## Next pages — domains.credentials.login_endpoint.json_payload.blindfold_secret_info / 353ab21936d3 / 7

- [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-4e213a146be541001c0e10dd21caa4a63a3f4bae702911cb8b098b7ce54eee5f)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-8da733b8cee0cf4397df033d053756848c476d2340cc23c96e0545d8c23e28a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0773fe22d316f4564f342fb8c550a1c4c2c35fb3d633173f7e330b9174008c3d"></a>

## domains.credentials.login_endpoint.json_payload.clear_secret_info — domains.credentials.login_endpoint.json_payload.clear_secret_info / b1ccd629d27e / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-2957691b8432fafac009d886cf355cf706257dce6e55943c9eb25f56ad52668c)
- [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-4e213a146be541001c0e10dd21caa4a63a3f4bae702911cb8b098b7ce54eee5f)
- domains.credentials.login_endpoint.json_payload.clear_secret_info

<a id="canonical-5a156ef9e281f4e51df3adf952e28519eb9da42875a2e2181855e239bac60ced"></a>

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

<a id="canonical-1aa791ebbe5543d2d8301f5643ce988966f6f7e713bfb3e0c65c3437795ad172"></a>

## Direct properties — domains.credentials.login_endpoint.json_payload.clear_secret_info / b1ccd629d27e / 3

<a id="canonical-b0d46b7cb2f76804c2cc363c6064b3e741e195faced77a4f6ea51c718f3b33b8"></a>

<a id="canonical-8de4524f96d1561d73472cb7b661df884e86dab99e472961eab3b4592b9e7561"></a>

## provider_ref property — domains.credentials.login_endpoint.json_payload.clear_secret_info / b1ccd629d27e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-711cfa1a8bdff8537064777757ed6ab0a97ec3a0e79f3990b0c042b5d17c61ba"></a>

<a id="canonical-072947000f1e8eb05a22b76b89383dd7822bdd289e6787b8ab3b471906787b86"></a>

## url property — domains.credentials.login_endpoint.json_payload.clear_secret_info / b1ccd629d27e / 5

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

<a id="canonical-17ad9bd8ba35a083b6b5c6e5a685da1f1f31377c8545af7735e4a9c515ef39a3"></a>

## Next pages — domains.credentials.login_endpoint.json_payload.clear_secret_info / b1ccd629d27e / 6

- [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-4e213a146be541001c0e10dd21caa4a63a3f4bae702911cb8b098b7ce54eee5f)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-c230dbede3e63935fb094f7700976ae8bd6b622e30d85cc9b74808729a4304c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b7d3ef0fd540ace2cc09c7af02de687859b26d6e1fa8f405085b18bfbf50423"></a>

## domains.credentials.standard — domains.credentials.standard / 0512b8aa2ef2 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-dd35e25fc4cca7608f8ec30bdd41c63863903a9448583be39f70f94f051a0006)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- domains.credentials.standard

<a id="canonical-814bd195197a95f42db15f95e7e3b63e1c1d183c549d97acaed0cd33263a157a"></a>

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

<a id="canonical-94b30a5ed530dd908be0af9a149a139111ed8ff780c161cc7f60a5685cb5dcfd"></a>

## Direct properties — domains.credentials.standard / 0512b8aa2ef2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-faac76b2d55a41fb0625c9da862844e982f49a1860aa6aab7519b6c4f3224601"></a>

## Next pages — domains.credentials.standard / 0512b8aa2ef2 / 4

- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-f0c8917b6f14fba20175b2a31713f14189c5c21e2be61dcaf01ccf48e9ffac16)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-bef0046bbe3fd9de7166290fd32fe2257e363fcf53500939061071a13739a1e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aab908eaca77406b72abefc0a89aa18b9ca547b295ee98a0ea426e66b12bec2"></a>

## every_day — every_day / 3df03bbc3a8a / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- every_day

<a id="canonical-4259fe1d574ebfac934632b5dc810af17bd0bdf54c44cab7710c21561a3dfbb3"></a>

Type: `["object", {}]`. Computed.

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

- [every_day](data-sources--api_testing--reference--group-001.md#canonical-4259fe1d574ebfac934632b5dc810af17bd0bdf54c44cab7710c21561a3dfbb3)
- [every_month](data-sources--api_testing--reference--group-001.md#canonical-bae3add6cd132f6b276d2de75d79ab1cbe63782cd2c27286b6f54c81c96b08f4)
- [every_week](data-sources--api_testing--reference--group-001.md#canonical-1ac2e3c3805c0ecfdd81d3c61d3c4d8fb869aa8be8b9769c37a994973f940cbc)

Select alternatives according to the provider validators above.

<a id="canonical-b118afbf1c262fd4940e1f4f96273fb3d9b5a5dc68155eb5958d40be0ef9a6ea"></a>

## Direct properties — every_day / 3df03bbc3a8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-96a736d0dfad20ac8cc6e94277db6781e74acbe6dfd98a6246fd3ed336b1e3c6"></a>

## Next pages — every_day / 3df03bbc3a8a / 4

- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-1ad13f49880acec74448818804753c339eaf3e09b75d15858a684b11d51fc55e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7da7359fbd9e04d8bbbab68bc204cc0ed8701c2d59af72de7bd865f94d6cf186"></a>

## every_month — every_month / 9088fa984cba / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- every_month

<a id="canonical-bae3add6cd132f6b276d2de75d79ab1cbe63782cd2c27286b6f54c81c96b08f4"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-35ef1cf1e38541b8bf284483d507dfb4b70849fa183b2644492dcfccdd6a6d6f"></a>

## Direct properties — every_month / 9088fa984cba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4f376e718b9266521e026814b0de8a831f8d799cd8d9b498777a7feddef45ee"></a>

## Next pages — every_month / 9088fa984cba / 4

- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-5791b434f21e166201743305c608ca8fd93d0ee5060e234906071d1c9ffab591"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3beb35f119829ace5fa5a9ae7d5d36cc259925204a3b7a6bcd5e1da85d350f3f"></a>

## every_week — every_week / c8fc54206d84 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- every_week

<a id="canonical-1ac2e3c3805c0ecfdd81d3c61d3c4d8fb869aa8be8b9769c37a994973f940cbc"></a>

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

<a id="canonical-ec18387ca4d098c87df354463fb048d8b5e8acac47478f36f33501e5588e0649"></a>

## Direct properties — every_week / c8fc54206d84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce46f0cced65d24c6d291951047b7d89afc201667d97218696abf11a9674506a"></a>

## Next pages — every_week / c8fc54206d84 / 4

- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
