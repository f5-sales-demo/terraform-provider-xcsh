---
page_title: "xcsh_alert_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver reference."
---

# xcsh_alert_receiver reference

<a id="canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0d823387e15f42fa5588b6a5eab984f2be9f1dba0e10af4ab31b2635bce7bfe"></a>

## Property reference — Property reference / f158f3f9a88d / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- Property reference

<a id="canonical-2fcc5d1aacc212aae3b429b5390345c49a2e73b5c0690be0a451f457d6e7d0c0"></a>

## Direct properties — Property reference / f158f3f9a88d / 3

<a id="canonical-c5e51a7d6a43f8ff78e8e2f118dac3811e90ec160258e5e5469f523bb3a188e9"></a>

<a id="canonical-2e1213f905a74ca82155607a87f4f512e54a28f4755897ceb0ce5289a5011ea9"></a>

## annotations property — Property reference / f158f3f9a88d / 4

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

<a id="canonical-00ff8e5edd9712a25ca54703138286a29ed898633a1a0d6feffa53f9ded61149"></a>

<a id="canonical-46f6221f2b6ac965730aa9e83165a4412149c14f97c36404dfcdeb0213e7c4f7"></a>

## description property — Property reference / f158f3f9a88d / 5

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

<a id="canonical-9ec5b518e4cd1429aed6556809ad0053ea0a68f5f20fa38163f61946d863e7b4"></a>

<a id="canonical-76de5618403f3880de09760ef9504fbe34c7527fd28728586016ac13a689632b"></a>

## disable property — Property reference / f158f3f9a88d / 6

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

- [email](resources--alert_receiver--reference--group-001.md#canonical-7b28ebb37ce3eda385800dfb1d8d3f9e21d3271d44d00c71ce9d23fa98b2438e): complete subsection reference.

<a id="canonical-de7b2120ef763eda8006cdfbc3731fb32fdc15457b5d0da24a34ed3555cb73ae"></a>

<a id="canonical-40584c7e246135ab271bde44e13dbe0049b8ed80256fe31841a243662ae73582"></a>

## id property — Property reference / f158f3f9a88d / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-9d771a6d90af1b428b5246040084529a59745e1a8a6b3a8ce1fb84d98f876e9c"></a>

<a id="canonical-14e68e3777a0ba63a5d322b5fc7593c2d29b98599db7f0779f87a6be25cdaedb"></a>

## labels property — Property reference / f158f3f9a88d / 8

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

<a id="canonical-706f72e66a990b99998297fcf9457561d33fb93ad731205b6533f14856fc7652"></a>

<a id="canonical-9141f15d97e421622578be92eb46fe6dcddd633062022187014f5945c1237a70"></a>

## name property — Property reference / f158f3f9a88d / 9

Type: `"string"`. Required.

Name of the Alert Receiver. Must be unique within the namespace.

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

<a id="canonical-50a42af193406f0b15e48db913275150716a98a74a29f8a405e842a66e34244f"></a>

<a id="canonical-5ebc44916ea7f2cef1ef7b6cb477b6f70487234949266ad86bcb20812df17b8e"></a>

## namespace property — Property reference / f158f3f9a88d / 10

Type: `"string"`. Required.

Namespace where the Alert Receiver is created.

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

- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-cd2eed9162ec23e59e3b4e74752ca5bf1bb34a4c615ab1ad9f522e726fa5a999): complete subsection reference.

- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-8c16c3910cf89b252a779eb7cbbaf7be56b04684d35e6da07e68c85237cac141): complete subsection reference.

- [slack](resources--alert_receiver--reference--group-001.md#canonical-5eb2462c49807895168fa3c04b9823bd1e6412c90d7ac8213a8e7c7553d860c5): complete subsection reference.

- [sms](resources--alert_receiver--reference--group-001.md#canonical-b9ebc6bafb47a6223e03da567f4c7441fc83f2d6db862699593f75feb6f44fa2): complete subsection reference.

- [timeouts](resources--alert_receiver--reference--group-001.md#canonical-4d698c906228b09a3dcf057c8ff469cf2ca0acb5241d136bad8e5ef1549afca9): complete subsection reference.

- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693): complete subsection reference.

<a id="canonical-29bf3bb301bc17808fabcdeb068e5e54ebd362cc870201f51f1d20837d02bc0d"></a>

## All schema paths — Property reference / f158f3f9a88d / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--alert_receiver--reference--group-001.md#canonical-c5e51a7d6a43f8ff78e8e2f118dac3811e90ec160258e5e5469f523bb3a188e9) |
| `description` | [description](resources--alert_receiver--reference--group-001.md#canonical-00ff8e5edd9712a25ca54703138286a29ed898633a1a0d6feffa53f9ded61149) |
| `disable` | [disable](resources--alert_receiver--reference--group-001.md#canonical-9ec5b518e4cd1429aed6556809ad0053ea0a68f5f20fa38163f61946d863e7b4) |
| `email` | [email](resources--alert_receiver--reference--group-001.md#canonical-0c0b086032c3186aef56a69309a187dfae6aa2a461a53e44b7190d1c0fa2a693) |
| `email.email` | [email.email](resources--alert_receiver--reference--group-001.md#canonical-18ffc9c8599cc811e9f4536b7f49f7cae58568d3ed70c802ba9b224af58467d6) |
| `id` | [id](resources--alert_receiver--reference--group-001.md#canonical-de7b2120ef763eda8006cdfbc3731fb32fdc15457b5d0da24a34ed3555cb73ae) |
| `labels` | [labels](resources--alert_receiver--reference--group-001.md#canonical-9d771a6d90af1b428b5246040084529a59745e1a8a6b3a8ce1fb84d98f876e9c) |
| `name` | [name](resources--alert_receiver--reference--group-001.md#canonical-706f72e66a990b99998297fcf9457561d33fb93ad731205b6533f14856fc7652) |
| `namespace` | [namespace](resources--alert_receiver--reference--group-001.md#canonical-50a42af193406f0b15e48db913275150716a98a74a29f8a405e842a66e34244f) |
| `opsgenie` | [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-d575cd8c91e051635db8285d5ffef6424cc678a5629dd5982381464c63f0887b) |
| `opsgenie.api_key` | [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-d9268800146255c55df554e5c62bd51c5d72f894f6c5878c62ea449c5d8c2c28) |
| `opsgenie.api_key.blindfold_secret_info` | [opsgenie.api_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-334ed8f22a7373bdb10df072bff36a970caad5daf1fc21dedfddef4b408ed949) |
| `opsgenie.api_key.blindfold_secret_info.decryption_provider` | [opsgenie.api_key.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-2bb7102503c46d10b85b1911a72294e745d8fe3726f50a44fe1034dc1eccb2f7) |
| `opsgenie.api_key.blindfold_secret_info.location` | [opsgenie.api_key.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-94d4682170c881d1a53d13bda11204374a8f6eff4a1026b34b9ff95b70e2d14a) |
| `opsgenie.api_key.blindfold_secret_info.store_provider` | [opsgenie.api_key.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-8a556db32f6243820f9e3bdcc58ba13f6171348e9dd333cb3a1ed089de401c62) |
| `opsgenie.api_key.clear_secret_info` | [opsgenie.api_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-9c36ea25a1e28d746192db26f9ea0267ddcf708a0659579c74c0b51d8a3770d2) |
| `opsgenie.api_key.clear_secret_info.provider_ref` | [opsgenie.api_key.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-2baac4d5f5e018acda16744d8781d83b6de97fa96747a7e71da299db4cc26ba5) |
| `opsgenie.api_key.clear_secret_info.url` | [opsgenie.api_key.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-52269b3cae3f9b4f3a28ebe0e64642435a2c104465631186e2bf27f597d720d1) |
| `opsgenie.url` | [opsgenie.url](resources--alert_receiver--reference--group-001.md#canonical-66b50dfba699c17a64f94ad0213c11383c8056513455ce75b1ea66ce137b9abf) |
| `pagerduty` | [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-d0c0e1bf76a7586625466c63a5853cda0ba616ee56505f7bdef5ca0d79939fb3) |
| `pagerduty.routing_key` | [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-3662dae72c6c72091a7013b3709e5eee40b02b7b79882a6f358955bf80f0bc31) |
| `pagerduty.routing_key.blindfold_secret_info` | [pagerduty.routing_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0af03fc81dd6fc881e68eb7181a635e9af42630be3413e2dc046e75a88462876) |
| `pagerduty.routing_key.blindfold_secret_info.decryption_provider` | [pagerduty.routing_key.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-cfddea51c1af2c75319e3e0096c65f2c60638e88e6e1d43d4905f6ff4a9b4b63) |
| `pagerduty.routing_key.blindfold_secret_info.location` | [pagerduty.routing_key.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-a9469712d9031daa5977baec373aa1c73ac8c1124e8bacb7c15627121e52cf6d) |
| `pagerduty.routing_key.blindfold_secret_info.store_provider` | [pagerduty.routing_key.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-b821c2adb3554cc71f14f46a5c214717ed772d67642ac7bf3b9347b1b7f8f75a) |
| `pagerduty.routing_key.clear_secret_info` | [pagerduty.routing_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-efff300feb16cd9845174b56db28ef52348c0b8f3a30955e5bb826f655814881) |
| `pagerduty.routing_key.clear_secret_info.provider_ref` | [pagerduty.routing_key.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-7d52dba217102bd60669668808192c50ba81eaab2c1eb89590d5c7d0812a0cfe) |
| `pagerduty.routing_key.clear_secret_info.url` | [pagerduty.routing_key.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-a848534f2041b2863f97282862f14f0a2a1778ffd72f344c366eb3da307f7361) |
| `pagerduty.url` | [pagerduty.url](resources--alert_receiver--reference--group-001.md#canonical-2e099ed47f71b5d35aac664929e5183e48105349af938e7f5405abba4cb5761e) |
| `slack` | [slack](resources--alert_receiver--reference--group-001.md#canonical-78c1a11caea868b00c64c0f333aefea848b50f06456150ca377bfdaf90b0a34a) |
| `slack.channel` | [slack.channel](resources--alert_receiver--reference--group-001.md#canonical-471d3cda3bef1e5da94ad01f74c974be01d634320c2e8d7f40ca4a1e237493ad) |
| `slack.url` | [slack.url](resources--alert_receiver--reference--group-001.md#canonical-5d3d3017e55db32da581effe7f234662c3c3e166098326e31383ad0366fc2de5) |
| `slack.url.blindfold_secret_info` | [slack.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-223fb6ac017ef82a6054190736c929b3158ccedbc2b097f667b7ac6bcb4047ce) |
| `slack.url.blindfold_secret_info.decryption_provider` | [slack.url.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-1acf43f716b217eb4bdd915bcd6558dd3d0b707dfe6731f9500707b47fc53b4f) |
| `slack.url.blindfold_secret_info.location` | [slack.url.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-05d02fb1dd2e4475c7fd69eee58b431d3beda090737e364eb50616c0a680723a) |
| `slack.url.blindfold_secret_info.store_provider` | [slack.url.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-139ce373580fa6337444e4188b48ae8d06b0b6d3921c37afdba873faa84bbd79) |
| `slack.url.clear_secret_info` | [slack.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-1a65a7e2cbaceb15427e4cf5ad314ca4eae52eaf5d55213449df8ca163ded0e0) |
| `slack.url.clear_secret_info.provider_ref` | [slack.url.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-5cf1d29f0f1bbd720f26f1e0477d3a9eb940c63657b64a2a6347cfdb1c4a6a32) |
| `slack.url.clear_secret_info.url` | [slack.url.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-fc2f00b68e3775de1d4d799a9fc9af0c7f94469f512b65a473efbec8b48048b1) |
| `sms` | [sms](resources--alert_receiver--reference--group-001.md#canonical-818d75cc41c7615ffc51b808095b929afa600fbf24a1d15d007ac8680835dece) |
| `sms.contact_number` | [sms.contact_number](resources--alert_receiver--reference--group-001.md#canonical-08afdce794c88332e83456874a7d218e0ba5d8ffc60c5a0b804a9e89cf3c4cd4) |
| `timeouts` | [timeouts](resources--alert_receiver--reference--group-001.md#canonical-d51010d031ebfa17cc954fbd40ddb2d940de1b635c14c4b34113544aa9496280) |
| `timeouts.create` | [timeouts.create](resources--alert_receiver--reference--group-001.md#canonical-9f57e544a4597ee910ea2dc5baeda562185dc57ef7ab7518be8e58e212bc6429) |
| `timeouts.delete` | [timeouts.delete](resources--alert_receiver--reference--group-001.md#canonical-8380363a531107199203134fed44d17bc3f97eb23e3d71e551d057ac7eb71ce8) |
| `timeouts.read` | [timeouts.read](resources--alert_receiver--reference--group-001.md#canonical-dfecf90f45b0d39e5a8f18762de54f2386a7168346e67760fdd4fc017e550279) |
| `timeouts.update` | [timeouts.update](resources--alert_receiver--reference--group-001.md#canonical-aad6c9e6dfd1f93f57c2ed3e153663d9adf7b8cb1f9008d2bb2843e2c9c9b499) |
| `webhook` | [webhook](resources--alert_receiver--reference--group-001.md#canonical-7ff5070337af3f8c75f04742c6019aba8c74123aa2a4c7d99940445a32db63b7) |
| `webhook.http_config` | [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-8460ffda1e1d529ca784751fdb26c85d858dd234fb91b458721a96a09035c36b) |
| `webhook.http_config.auth_token` | [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-47bfcf94135899c9fc64e8c5a5be54d1a3dee4020fd67233b5a05366a8013f2a) |
| `webhook.http_config.auth_token.token` | [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-8ac484a30937470771d8ba5fc7bafe6d85f7c2a0daf42a1690e55f541b8ed670) |
| `webhook.http_config.auth_token.token.blindfold_secret_info` | [webhook.http_config.auth_token.token.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-c44ee64301c895771d9c4ab20ca24abb2c941ebae718e79d7e817c72d15b2fed) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-3eabb976317008628418892a9f37f21f072d1d2b7e6b666556fd87e066697439) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.location` | [webhook.http_config.auth_token.token.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-e5e608f6ee30b4aafbd5f5be0f5ecabce74410754cc0e557beffc2ca90ba4f9b) |
| `webhook.http_config.auth_token.token.blindfold_secret_info.store_provider` | [webhook.http_config.auth_token.token.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-2cb3bcd010f1bbc66f6fa02aedad45051c7be7690ae192ea129f92f685e48286) |
| `webhook.http_config.auth_token.token.clear_secret_info` | [webhook.http_config.auth_token.token.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-a5cbb0ec6197e3ebfef1350f26752e514559a7a8a1d133cd0001bf6811e36671) |
| `webhook.http_config.auth_token.token.clear_secret_info.provider_ref` | [webhook.http_config.auth_token.token.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-20fea35d856fcf979433545b8d3c3f97d4273585bd3728276edd6bcdf56d6339) |
| `webhook.http_config.auth_token.token.clear_secret_info.url` | [webhook.http_config.auth_token.token.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-a3bae3dfa106b040b4105898937ea70d58da7649ca80dd45134a8e6357260e20) |
| `webhook.http_config.basic_auth` | [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-9ee54ca45b085ce2ede8a84d81d682bca11777c53eaf38d18d0d7368d988ee9b) |
| `webhook.http_config.basic_auth.password` | [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-5d3dba706544beb0f0b94acb66cfde5dd12f7a982c959e86c62b8e6e103e4ca3) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info` | [webhook.http_config.basic_auth.password.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-ecf2f5fb88efd9b4a24ee5a8e186323f5248fd79d0d2c7563a9722db7bff6d5a) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-02bc2d854fda381fccae6a4b1d356b6628b9e8cfc04bcae547a32502cb029e61) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.location` | [webhook.http_config.basic_auth.password.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-b484674ea081e2403c9f7d65c50a5be7b8deb30049701e6efde275c5a052e92a) |
| `webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider` | [webhook.http_config.basic_auth.password.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-22995b2adec3e39c806b4e6a16bd328cca493b5535fcefc519f8c9a9a242e41e) |
| `webhook.http_config.basic_auth.password.clear_secret_info` | [webhook.http_config.basic_auth.password.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-6fb3ed599d684198009398e557a0354895af040c1143394caa61e062191c0340) |
| `webhook.http_config.basic_auth.password.clear_secret_info.provider_ref` | [webhook.http_config.basic_auth.password.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-17eb314d25cc59a62a0d3ce32de0cea391ad31eecaad6d34c21c595b3ac0d8ec) |
| `webhook.http_config.basic_auth.password.clear_secret_info.url` | [webhook.http_config.basic_auth.password.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-c409e032d11bff461ba0adb56b4c12882ca45284ffdeaf04be8559d300c26e0f) |
| `webhook.http_config.basic_auth.user_name` | [webhook.http_config.basic_auth.user_name](resources--alert_receiver--reference--group-001.md#canonical-36f64fc4e87e6cb6273f969fd39580da780b0918f7786921b16aad8172312242) |
| `webhook.http_config.client_cert_obj` | [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-d786a3df5ed0f861c9027620fee05a67668fae06742f45bdc840abf06e487f3b) |
| `webhook.http_config.client_cert_obj.use_tls_obj` | [webhook.http_config.client_cert_obj.use_tls_obj](resources--alert_receiver--reference--group-001.md#canonical-649adc1d5d30f774c59cad4d62f84e982593a804f0473a54cb26144985923bf1) |
| `webhook.http_config.client_cert_obj.use_tls_obj.kind` | [webhook.http_config.client_cert_obj.use_tls_obj.kind](resources--alert_receiver--reference--group-001.md#canonical-ab01f1630dc657da32b6b27a34a84606378f09e2f34355d0a3ba0dc1549e557e) |
| `webhook.http_config.client_cert_obj.use_tls_obj.name` | [webhook.http_config.client_cert_obj.use_tls_obj.name](resources--alert_receiver--reference--group-001.md#canonical-19e402718a8448a70e5bfad95055574609d40bbbd10a65c2b8c2fa612b0e4277) |
| `webhook.http_config.client_cert_obj.use_tls_obj.namespace` | [webhook.http_config.client_cert_obj.use_tls_obj.namespace](resources--alert_receiver--reference--group-001.md#canonical-407441b9089acde5736e8f1b0a41bdee7f882a260bc718a2a6ba376d1fe3063d) |
| `webhook.http_config.client_cert_obj.use_tls_obj.tenant` | [webhook.http_config.client_cert_obj.use_tls_obj.tenant](resources--alert_receiver--reference--group-001.md#canonical-027207ecd1952f39912e50ac9f124ed642d7e146371e63e499f1bb061a48befc) |
| `webhook.http_config.client_cert_obj.use_tls_obj.uid` | [webhook.http_config.client_cert_obj.use_tls_obj.uid](resources--alert_receiver--reference--group-001.md#canonical-2cb755c86f034c8676e03756e1f528db89f07e9b1c4c23198faf239243c0beb2) |
| `webhook.http_config.enable_http2` | [webhook.http_config.enable_http2](resources--alert_receiver--reference--group-001.md#canonical-e97e76520a3f422e976f50405eec11b0f09ca0d6ab4173275867d0937943a918) |
| `webhook.http_config.follow_redirects` | [webhook.http_config.follow_redirects](resources--alert_receiver--reference--group-001.md#canonical-8f037bc307cb8b8261989f987e9570950020c369e8487437bbdaa8dc42bb3223) |
| `webhook.http_config.no_authorization` | [webhook.http_config.no_authorization](resources--alert_receiver--reference--group-001.md#canonical-dac7ad8effc278e5ab98cf9fa2620b94fd589c9a410007fb0e4617d7f820c7bd) |
| `webhook.http_config.no_tls` | [webhook.http_config.no_tls](resources--alert_receiver--reference--group-001.md#canonical-c03c5294cfcfda533ae0d8f19c20af73fda848d0e1d28d58477a0ba8da008b03) |
| `webhook.http_config.use_tls` | [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-e30ec7dd9ea96fb2c10803fd4d5a235acf297aa1201f0c4fa44152d6add636f4) |
| `webhook.http_config.use_tls.disable_sni` | [webhook.http_config.use_tls.disable_sni](resources--alert_receiver--reference--group-001.md#canonical-afa31575c1bbbd8cde912b94b852af147f6be50590b4f550b7ac9d15a6ca00b5) |
| `webhook.http_config.use_tls.max_version` | [webhook.http_config.use_tls.max_version](resources--alert_receiver--reference--group-001.md#canonical-9ed5e534ba7b00fb86c2721602f7bb6df00724450026471ed2d574434d097929) |
| `webhook.http_config.use_tls.min_version` | [webhook.http_config.use_tls.min_version](resources--alert_receiver--reference--group-001.md#canonical-55558a20167ed63e4df3629ab3784779f338f0685fea64da01a8cca3efd302cd) |
| `webhook.http_config.use_tls.sni` | [webhook.http_config.use_tls.sni](resources--alert_receiver--reference--group-001.md#canonical-53e5538791e98a40310da5ae85e133eb6aee70ef73e277743b745ae6e674287d) |
| `webhook.http_config.use_tls.use_server_verification` | [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-5c17a1d4745d9298c95d8f4adeaa0cbef44def513b2507ee56fd88f98d3e726d) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-acbe8542d41360cfe8d873e1601eff6c94b2f51e34ec57c2fdfe5ea2012155aa) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-5520a8b763a0040db998b6bd2c0f57f8cfc0ea2536a9c575c52f3a7d8accc181) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.kind](resources--alert_receiver--reference--group-001.md#canonical-a73e7ab7b9217e0577836c9a323541ad18ccb9826a3568f891fd98a8de713627) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.name](resources--alert_receiver--reference--group-001.md#canonical-6856aaccef1b7bd2bf6b12ddbaed2efc188d588cec133aa741ad89569e9e0ccb) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.namespace](resources--alert_receiver--reference--group-001.md#canonical-bfae57937a16b43893442c5891bf81de5a8642291feb6c2084ad8c213d1151ef) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.tenant](resources--alert_receiver--reference--group-001.md#canonical-52ebc2bc312b470276e1785dba2745a4a08a2efcd359a0ec7f4bd4dfbb92dbb9) |
| `webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid` | [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca.uid](resources--alert_receiver--reference--group-001.md#canonical-9fb782087788f01acb4317f578c97fd1621dddfdfb40126b7f28508beeab0b95) |
| `webhook.http_config.use_tls.volterra_trusted_ca` | [webhook.http_config.use_tls.volterra_trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-e4985a4ddb855154bbe168acee47c808373202310814c7ecaea1eec56432f912) |
| `webhook.url` | [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-201a2eb8601bd06a5248dfbde10a1cfa4a9aa4da45b9f447c97e144eb62bdb33) |
| `webhook.url.blindfold_secret_info` | [webhook.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-25b8b44dd2b44a1948216eb8fec3578adaad92f17e08f91c9e0083779d1938b1) |
| `webhook.url.blindfold_secret_info.decryption_provider` | [webhook.url.blindfold_secret_info.decryption_provider](resources--alert_receiver--reference--group-001.md#canonical-922999a21f723853e56393d2fa54da37aa99d1d818b1fb22e63c4d8f33774c68) |
| `webhook.url.blindfold_secret_info.location` | [webhook.url.blindfold_secret_info.location](resources--alert_receiver--reference--group-001.md#canonical-b46a6877dee9e1bb25c9dbfbd86f86cbaa660b5c0d0e5a48a95e61ae026487e9) |
| `webhook.url.blindfold_secret_info.store_provider` | [webhook.url.blindfold_secret_info.store_provider](resources--alert_receiver--reference--group-001.md#canonical-e4d32cb4c5a8f723ff762beece7379ce734bfe51d9f0a3f69005d4820ed249bb) |
| `webhook.url.clear_secret_info` | [webhook.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-39bf30a429874541d8d870bed6e8d6e6afa368a5a10c22a9fd97594d07e6ba18) |
| `webhook.url.clear_secret_info.provider_ref` | [webhook.url.clear_secret_info.provider_ref](resources--alert_receiver--reference--group-001.md#canonical-660966b3de3191a31d59dbf32eac47b1fed1f198aacb5b3c2c08893ebffb5843) |
| `webhook.url.clear_secret_info.url` | [webhook.url.clear_secret_info.url](resources--alert_receiver--reference--group-001.md#canonical-2ed52239b58c64a5c172748317889bf3c6407392123738c7ede2a7c2bb547aab) |

<a id="canonical-1db5a8d35935eb13ab80a8780ed37d7e39dbbd2567079f1902ca22353553346d"></a>

## Next pages — Property reference / f158f3f9a88d / 12

- [email](resources--alert_receiver--reference--group-001.md#canonical-7b28ebb37ce3eda385800dfb1d8d3f9e21d3271d44d00c71ce9d23fa98b2438e)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-cd2eed9162ec23e59e3b4e74752ca5bf1bb34a4c615ab1ad9f522e726fa5a999)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-8c16c3910cf89b252a779eb7cbbaf7be56b04684d35e6da07e68c85237cac141)
- [slack](resources--alert_receiver--reference--group-001.md#canonical-5eb2462c49807895168fa3c04b9823bd1e6412c90d7ac8213a8e7c7553d860c5)
- [sms](resources--alert_receiver--reference--group-001.md#canonical-b9ebc6bafb47a6223e03da567f4c7441fc83f2d6db862699593f75feb6f44fa2)
- [timeouts](resources--alert_receiver--reference--group-001.md#canonical-4d698c906228b09a3dcf057c8ff469cf2ca0acb5241d136bad8e5ef1549afca9)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-7b28ebb37ce3eda385800dfb1d8d3f9e21d3271d44d00c71ce9d23fa98b2438e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bea8900fc1254a373cc17378f7de34da698f4f28aa00ec0d0cc95a2bb4629814"></a>

## email — email / 40cb1690c4c9 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- email

<a id="canonical-0c0b086032c3186aef56a69309a187dfae6aa2a461a53e44b7190d1c0fa2a693"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: email, opsgenie, pagerduty, slack, sms, webhook\] Email Configuration.

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

- [email](resources--alert_receiver--reference--group-001.md#canonical-0c0b086032c3186aef56a69309a187dfae6aa2a461a53e44b7190d1c0fa2a693)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-d575cd8c91e051635db8285d5ffef6424cc678a5629dd5982381464c63f0887b)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-d0c0e1bf76a7586625466c63a5853cda0ba616ee56505f7bdef5ca0d79939fb3)
- [slack](resources--alert_receiver--reference--group-001.md#canonical-78c1a11caea868b00c64c0f333aefea848b50f06456150ca377bfdaf90b0a34a)
- [sms](resources--alert_receiver--reference--group-001.md#canonical-818d75cc41c7615ffc51b808095b929afa600fbf24a1d15d007ac8680835dece)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-7ff5070337af3f8c75f04742c6019aba8c74123aa2a4c7d99940445a32db63b7)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
email {
  # Configure direct properties listed below.
}
```

<a id="canonical-29d4ab1137c4c2290e1f86995e89e5eb75e029722503416bef9ea84c7025ec0f"></a>

## Direct properties — email / 40cb1690c4c9 / 3

<a id="canonical-18ffc9c8599cc811e9f4536b7f49f7cae58568d3ed70c802ba9b224af58467d6"></a>

<a id="canonical-68d6edc71d261a4be31ec30f059f73f38936988f37f11f540e95e25f8956099a"></a>

## email property — email / 40cb1690c4c9 / 4

Type: `"string"`. Optional.

Email. Email ID of the user.

Upstream description:

Email ID of the user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

<a id="canonical-91d20f3a1d35196d1de35eac893f6aa39140aa2e0d36f89fa86f2d3ec54a258f"></a>

## Next pages — email / 40cb1690c4c9 / 5

- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-cd2eed9162ec23e59e3b4e74752ca5bf1bb34a4c615ab1ad9f522e726fa5a999"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43cb56c9a8dcfc3f5c145cb1804a973f86e58872ce962ff8c73c18a9c43f0ecc"></a>

## opsgenie — opsgenie / 6a516f8757e8 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- opsgenie

<a id="canonical-d575cd8c91e051635db8285d5ffef6424cc678a5629dd5982381464c63f0887b"></a>

Type: `"object"`. single nested block, Optional.

OpsGenie configuration to send alert notifications.

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
opsgenie {
  # Configure direct properties listed below.
}
```

<a id="canonical-442cc0ee9ccf8a2f421cb147162101af804049db00a357168c347fe3e3c0af1e"></a>

## Direct properties — opsgenie / 6a516f8757e8 / 3

- [api_key](resources--alert_receiver--reference--group-001.md#canonical-5abe45456d939b9e9871abec53c6efabcb8d5494d768807d5ba1401551f425d2): complete subsection reference.

<a id="canonical-66b50dfba699c17a64f94ad0213c11383c8056513455ce75b1ea66ce137b9abf"></a>

<a id="canonical-fe379fd52f8f41082ad5ae90eec8e1142b33ab5731abce9172964995443a8229"></a>

## url property — opsgenie / 6a516f8757e8 / 4

Type: `"string"`. Optional.

API URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
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

<a id="canonical-4b6b4281a2d96b80f3ebdcbec10123c42a0712d127d56ed1a791e0e36aa41e28"></a>

## Next pages — opsgenie / 6a516f8757e8 / 5

- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-5abe45456d939b9e9871abec53c6efabcb8d5494d768807d5ba1401551f425d2)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-5abe45456d939b9e9871abec53c6efabcb8d5494d768807d5ba1401551f425d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-640fed12127b5fded2b81904f58f70b060e4e9fbcd98f1af6e88c0622a782f9d"></a>

## opsgenie.api_key — opsgenie.api_key / c188b6096835 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-cd2eed9162ec23e59e3b4e74752ca5bf1bb34a4c615ab1ad9f522e726fa5a999)
- opsgenie.api_key

<a id="canonical-d9268800146255c55df554e5c62bd51c5d72f894f6c5878c62ea449c5d8c2c28"></a>

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
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-eaf4c6ea292e2952d0d8b68d73fcb018199331a4a3bc358f43a91844af70f6ec"></a>

## Direct properties — opsgenie.api_key / c188b6096835 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-7653e1739c004447f9dfc621b3a5793750eda2511e9b57d6ea045cf106035f31): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-063bb92ea044c9746d0f23f92e634e456dc387b32213e083e3abf9796364d008): complete subsection reference.

<a id="canonical-137b780c6522f8d0d982c6ce35493c99626c6fdaffcbc946183eb31338b828c1"></a>

## Next pages — opsgenie.api_key / c188b6096835 / 4

- [opsgenie.api_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-7653e1739c004447f9dfc621b3a5793750eda2511e9b57d6ea045cf106035f31)
- [opsgenie.api_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-063bb92ea044c9746d0f23f92e634e456dc387b32213e083e3abf9796364d008)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-cd2eed9162ec23e59e3b4e74752ca5bf1bb34a4c615ab1ad9f522e726fa5a999)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-7653e1739c004447f9dfc621b3a5793750eda2511e9b57d6ea045cf106035f31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b40ffb5ceb5abc5f51fd1694cb2353363a39b93eb76c33de5822ceb74e1c33e1"></a>

## opsgenie.api_key.blindfold_secret_info — opsgenie.api_key.blindfold_secret_info / f2ea035ff82c / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-cd2eed9162ec23e59e3b4e74752ca5bf1bb34a4c615ab1ad9f522e726fa5a999)
- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-5abe45456d939b9e9871abec53c6efabcb8d5494d768807d5ba1401551f425d2)
- opsgenie.api_key.blindfold_secret_info

<a id="canonical-334ed8f22a7373bdb10df072bff36a970caad5daf1fc21dedfddef4b408ed949"></a>

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

<a id="canonical-3866af7d568dc5c0822e72b1e6e7d154b7401ecdf0c696de1fad54f55cb1ba85"></a>

## Direct properties — opsgenie.api_key.blindfold_secret_info / f2ea035ff82c / 3

<a id="canonical-2bb7102503c46d10b85b1911a72294e745d8fe3726f50a44fe1034dc1eccb2f7"></a>

<a id="canonical-68637f3c6a38805f4d63b37c6d8029f1c09cf87fd9f680639c974c6eab59c788"></a>

## decryption_provider property — opsgenie.api_key.blindfold_secret_info / f2ea035ff82c / 4

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

<a id="canonical-94d4682170c881d1a53d13bda11204374a8f6eff4a1026b34b9ff95b70e2d14a"></a>

<a id="canonical-8f1f7d21a8a7d4c7ad8bc295d7ad8b976031b97b8e8c538847fa3413d08ed274"></a>

## location property — opsgenie.api_key.blindfold_secret_info / f2ea035ff82c / 5

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

<a id="canonical-8a556db32f6243820f9e3bdcc58ba13f6171348e9dd333cb3a1ed089de401c62"></a>

<a id="canonical-ce0cd32ca4439c6353b6046b122f0855f217392534a94f12334590df186cf9c6"></a>

## store_provider property — opsgenie.api_key.blindfold_secret_info / f2ea035ff82c / 6

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

<a id="canonical-dd0a62f2777ca8335a33dd92e77b3b2af3873c1076aa899ee2b27a3867b8818e"></a>

## Next pages — opsgenie.api_key.blindfold_secret_info / f2ea035ff82c / 7

- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-5abe45456d939b9e9871abec53c6efabcb8d5494d768807d5ba1401551f425d2)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-063bb92ea044c9746d0f23f92e634e456dc387b32213e083e3abf9796364d008"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddb4e659f68bf333af36eb5c308e056f44d2f0f043c63c5ff287f1916159c5ec"></a>

## opsgenie.api_key.clear_secret_info — opsgenie.api_key.clear_secret_info / b66c17f9c26b / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [opsgenie](resources--alert_receiver--reference--group-001.md#canonical-cd2eed9162ec23e59e3b4e74752ca5bf1bb34a4c615ab1ad9f522e726fa5a999)
- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-5abe45456d939b9e9871abec53c6efabcb8d5494d768807d5ba1401551f425d2)
- opsgenie.api_key.clear_secret_info

<a id="canonical-9c36ea25a1e28d746192db26f9ea0267ddcf708a0659579c74c0b51d8a3770d2"></a>

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

<a id="canonical-d5bc1082e4982fe67b6fee2a06779f4dee0a65c76f0bb7688677f748aeb2ef4e"></a>

## Direct properties — opsgenie.api_key.clear_secret_info / b66c17f9c26b / 3

<a id="canonical-2baac4d5f5e018acda16744d8781d83b6de97fa96747a7e71da299db4cc26ba5"></a>

<a id="canonical-94da338aabb1806d73b6d0b475389c723fa31a412fb9b748dd3f03723059b3d1"></a>

## provider_ref property — opsgenie.api_key.clear_secret_info / b66c17f9c26b / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-52269b3cae3f9b4f3a28ebe0e64642435a2c104465631186e2bf27f597d720d1"></a>

<a id="canonical-32af0b40dd665eb06118341866ab3fc958ab2875a6e43d82c5b4280f04db09c9"></a>

## url property — opsgenie.api_key.clear_secret_info / b66c17f9c26b / 5

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

<a id="canonical-4903420eb1d8399dc3b1804a2e88faeb16b829102052bdb28c033e5fc7add902"></a>

## Next pages — opsgenie.api_key.clear_secret_info / b66c17f9c26b / 6

- [opsgenie.api_key](resources--alert_receiver--reference--group-001.md#canonical-5abe45456d939b9e9871abec53c6efabcb8d5494d768807d5ba1401551f425d2)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-8c16c3910cf89b252a779eb7cbbaf7be56b04684d35e6da07e68c85237cac141"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edc064c1924dc7899b4d4d8867a7c58335f457ee75629089984f2b2ee9a185a3"></a>

## pagerduty — pagerduty / 113f5dfd028f / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- pagerduty

<a id="canonical-d0c0e1bf76a7586625466c63a5853cda0ba616ee56505f7bdef5ca0d79939fb3"></a>

Type: `"object"`. single nested block, Optional.

PagerDuty configuration to send alert notifications.

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
pagerduty {
  # Configure direct properties listed below.
}
```

<a id="canonical-f865bd4ee2da11acee8e5c243733010a9aabff83e653a19fbbb7ef977a8bc8de"></a>

## Direct properties — pagerduty / 113f5dfd028f / 3

- [routing_key](resources--alert_receiver--reference--group-001.md#canonical-12eabe0d86bc622df7304e8916abe6044845bb3a42005403b226b825e24c4f55): complete subsection reference.

<a id="canonical-2e099ed47f71b5d35aac664929e5183e48105349af938e7f5405abba4cb5761e"></a>

<a id="canonical-f5ef57923d067f967f679028314eca2f6c9a0f03ff3223ec3071a317d2d35978"></a>

## url property — pagerduty / 113f5dfd028f / 4

Type: `"string"`. Optional.

Pager Duty URL. URL to send API requests to.

Upstream description:

URL to send API requests to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
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

<a id="canonical-66b769022151dab2cfa6ff06218cb1152fa52d6093c1d5d9da718dab33c51f43"></a>

## Next pages — pagerduty / 113f5dfd028f / 5

- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-12eabe0d86bc622df7304e8916abe6044845bb3a42005403b226b825e24c4f55)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-12eabe0d86bc622df7304e8916abe6044845bb3a42005403b226b825e24c4f55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d05b9d866640cde9ebcfef143b03e6906ae8057a5e5419f8cd214d16528c179"></a>

## pagerduty.routing_key — pagerduty.routing_key / 51b334c54572 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-8c16c3910cf89b252a779eb7cbbaf7be56b04684d35e6da07e68c85237cac141)
- pagerduty.routing_key

<a id="canonical-3662dae72c6c72091a7013b3709e5eee40b02b7b79882a6f358955bf80f0bc31"></a>

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
routing_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b3a2e937a9f4f886078989956bfce4cabd57ec6cae5db287e4468ed0bec5087"></a>

## Direct properties — pagerduty.routing_key / 51b334c54572 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-644aa62ca9500d87bbe15aaa1fb1cba0aefc758c78f2889ec859e4c40c8e869e): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-069bd8123bcbf3dd87221ba63fc8228f46feb95a8aef77c0c5d9d109e89ff196): complete subsection reference.

<a id="canonical-011fd5d848480908df6f73cfc1c62575b3c6fe43d1ee1f8332cafbf1f036cc5f"></a>

## Next pages — pagerduty.routing_key / 51b334c54572 / 4

- [pagerduty.routing_key.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-644aa62ca9500d87bbe15aaa1fb1cba0aefc758c78f2889ec859e4c40c8e869e)
- [pagerduty.routing_key.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-069bd8123bcbf3dd87221ba63fc8228f46feb95a8aef77c0c5d9d109e89ff196)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-8c16c3910cf89b252a779eb7cbbaf7be56b04684d35e6da07e68c85237cac141)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-644aa62ca9500d87bbe15aaa1fb1cba0aefc758c78f2889ec859e4c40c8e869e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-349666030518837d60928ead49c8e197f6711dc463ee19f6a2c8273ae77110f0"></a>

## pagerduty.routing_key.blindfold_secret_info — pagerduty.routing_key.blindfold_secret_info / 0100de205dab / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-8c16c3910cf89b252a779eb7cbbaf7be56b04684d35e6da07e68c85237cac141)
- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-12eabe0d86bc622df7304e8916abe6044845bb3a42005403b226b825e24c4f55)
- pagerduty.routing_key.blindfold_secret_info

<a id="canonical-0af03fc81dd6fc881e68eb7181a635e9af42630be3413e2dc046e75a88462876"></a>

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

<a id="canonical-19d465726854b3e244194c6ed361a0a3a1a4684a6594b8750783b9e54a84ec0c"></a>

## Direct properties — pagerduty.routing_key.blindfold_secret_info / 0100de205dab / 3

<a id="canonical-cfddea51c1af2c75319e3e0096c65f2c60638e88e6e1d43d4905f6ff4a9b4b63"></a>

<a id="canonical-4c56f48df08a8626337f1a031d32c49a08e782fad741ee8ab93f3a588c482cf5"></a>

## decryption_provider property — pagerduty.routing_key.blindfold_secret_info / 0100de205dab / 4

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

<a id="canonical-a9469712d9031daa5977baec373aa1c73ac8c1124e8bacb7c15627121e52cf6d"></a>

<a id="canonical-ac6615c8ff030556f74023c021131957561b74e9d2c8af3a217973089801f650"></a>

## location property — pagerduty.routing_key.blindfold_secret_info / 0100de205dab / 5

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

<a id="canonical-b821c2adb3554cc71f14f46a5c214717ed772d67642ac7bf3b9347b1b7f8f75a"></a>

<a id="canonical-10724576f87df2f34d793cd581b70e228ee67b21efc2c2468d9de06fbbc34e71"></a>

## store_provider property — pagerduty.routing_key.blindfold_secret_info / 0100de205dab / 6

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

<a id="canonical-5754985d15b3a261eca9dbb13e47f8c2fa118691d1430628a64b109de89be612"></a>

## Next pages — pagerduty.routing_key.blindfold_secret_info / 0100de205dab / 7

- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-12eabe0d86bc622df7304e8916abe6044845bb3a42005403b226b825e24c4f55)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-069bd8123bcbf3dd87221ba63fc8228f46feb95a8aef77c0c5d9d109e89ff196"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-582998cdb443cc3749ddda0d12d84417f98d49a65fbde6d1e3a40a2365a9504c"></a>

## pagerduty.routing_key.clear_secret_info — pagerduty.routing_key.clear_secret_info / fb9766e89338 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [pagerduty](resources--alert_receiver--reference--group-001.md#canonical-8c16c3910cf89b252a779eb7cbbaf7be56b04684d35e6da07e68c85237cac141)
- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-12eabe0d86bc622df7304e8916abe6044845bb3a42005403b226b825e24c4f55)
- pagerduty.routing_key.clear_secret_info

<a id="canonical-efff300feb16cd9845174b56db28ef52348c0b8f3a30955e5bb826f655814881"></a>

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

<a id="canonical-ea6eb913e3bd1d75c39a8eebc4883302cd1c7e4343880faca82965f575023aa4"></a>

## Direct properties — pagerduty.routing_key.clear_secret_info / fb9766e89338 / 3

<a id="canonical-7d52dba217102bd60669668808192c50ba81eaab2c1eb89590d5c7d0812a0cfe"></a>

<a id="canonical-3c9baf2a4e70fb6a814a0b6b29aad117ff73af81747cc31e5b10672f8440d6cf"></a>

## provider_ref property — pagerduty.routing_key.clear_secret_info / fb9766e89338 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a848534f2041b2863f97282862f14f0a2a1778ffd72f344c366eb3da307f7361"></a>

<a id="canonical-5a5c9e519d5b1ac07f56c3ae918d68ec28d9483f132e36213ca224b97154a1fa"></a>

## url property — pagerduty.routing_key.clear_secret_info / fb9766e89338 / 5

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

<a id="canonical-19ee0e7bd504a538619ae518cf6027d74b084725d5bf37bd42e1d30a7cc7766d"></a>

## Next pages — pagerduty.routing_key.clear_secret_info / fb9766e89338 / 6

- [pagerduty.routing_key](resources--alert_receiver--reference--group-001.md#canonical-12eabe0d86bc622df7304e8916abe6044845bb3a42005403b226b825e24c4f55)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-5eb2462c49807895168fa3c04b9823bd1e6412c90d7ac8213a8e7c7553d860c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62d716ded52fd6766e7a57cd4afb40a68fcda7fc6db7ffd496e9cc88337ca67b"></a>

## slack — slack / f5d1e14fd534 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- slack

<a id="canonical-78c1a11caea868b00c64c0f333aefea848b50f06456150ca377bfdaf90b0a34a"></a>

Type: `"object"`. single nested block, Optional.

Slack configuration to send alert notifications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("channel")}
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
slack {
  # Configure direct properties listed below.
}
```

<a id="canonical-934f0576c04134fcea50510d0adbb1e8934a280ccafaff04d6aad42c80ba4089"></a>

## Direct properties — slack / f5d1e14fd534 / 3

<a id="canonical-471d3cda3bef1e5da94ad01f74c974be01d634320c2e8d7f40ca4a1e237493ad"></a>

<a id="canonical-27a672b96e1b05297279dcc0b24ff5a596cafd18f724d20737e4a94ad10beb38"></a>

## channel property — slack / f5d1e14fd534 / 4

Type: `"string"`. Optional.

Channel or user to send notifications to.

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
    "pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9-_]{1,80}$"
  }
}
```

- [url](resources--alert_receiver--reference--group-001.md#canonical-b3f57f9141adeccb5d712d48ac2928db06ef2d4d77d5ef95e62f2c45fc7ff00f): complete subsection reference.

<a id="canonical-93277b376978082837c9dd708d65ced97bf7d118e6070f01a8896340e05477e0"></a>

## Next pages — slack / f5d1e14fd534 / 5

- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-b3f57f9141adeccb5d712d48ac2928db06ef2d4d77d5ef95e62f2c45fc7ff00f)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-b3f57f9141adeccb5d712d48ac2928db06ef2d4d77d5ef95e62f2c45fc7ff00f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c17505ae6bb11e3da7bed985b601369e4428bcab2b54b7fbfe1ba362afb94b05"></a>

## slack.url — slack.url / 6d26fd6a1502 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [slack](resources--alert_receiver--reference--group-001.md#canonical-5eb2462c49807895168fa3c04b9823bd1e6412c90d7ac8213a8e7c7553d860c5)
- slack.url

<a id="canonical-5d3d3017e55db32da581effe7f234662c3c3e166098326e31383ad0366fc2de5"></a>

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
url {
  # Configure direct properties listed below.
}
```

<a id="canonical-92fd4c71ff24fe5b24ccce7145c177c50ac747e52c6c09a55c7928aaf0daed3a"></a>

## Direct properties — slack.url / 6d26fd6a1502 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-10d075979975ece24080ea6b7fdbfbf38d8db8a600303f4a58bd7581a4e2976b): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2dc26148c4c3e622e16372793e725190c27092c2f3be3c76163ef739b9ba7fca): complete subsection reference.

<a id="canonical-596abde0afc51969bf465939a07be616153201aba8e70e1c02e331f6130a649c"></a>

## Next pages — slack.url / 6d26fd6a1502 / 4

- [slack.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-10d075979975ece24080ea6b7fdbfbf38d8db8a600303f4a58bd7581a4e2976b)
- [slack.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2dc26148c4c3e622e16372793e725190c27092c2f3be3c76163ef739b9ba7fca)
- [slack](resources--alert_receiver--reference--group-001.md#canonical-5eb2462c49807895168fa3c04b9823bd1e6412c90d7ac8213a8e7c7553d860c5)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-10d075979975ece24080ea6b7fdbfbf38d8db8a600303f4a58bd7581a4e2976b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56fceccfc05e7c7e049dccc2baa1205a0163965596d5cde37439766a6ab0c504"></a>

## slack.url.blindfold_secret_info — slack.url.blindfold_secret_info / d978c2c7c169 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [slack](resources--alert_receiver--reference--group-001.md#canonical-5eb2462c49807895168fa3c04b9823bd1e6412c90d7ac8213a8e7c7553d860c5)
- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-b3f57f9141adeccb5d712d48ac2928db06ef2d4d77d5ef95e62f2c45fc7ff00f)
- slack.url.blindfold_secret_info

<a id="canonical-223fb6ac017ef82a6054190736c929b3158ccedbc2b097f667b7ac6bcb4047ce"></a>

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

<a id="canonical-ce3130d22dbdd95bed6ff53604ef256285bc42470242ac523ef85eb59fee748c"></a>

## Direct properties — slack.url.blindfold_secret_info / d978c2c7c169 / 3

<a id="canonical-1acf43f716b217eb4bdd915bcd6558dd3d0b707dfe6731f9500707b47fc53b4f"></a>

<a id="canonical-b3addd7dcb9e373a1901aefd74dbfb7056e4d9e4a7cbfd62c31df7d59b96cabc"></a>

## decryption_provider property — slack.url.blindfold_secret_info / d978c2c7c169 / 4

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

<a id="canonical-05d02fb1dd2e4475c7fd69eee58b431d3beda090737e364eb50616c0a680723a"></a>

<a id="canonical-4a77970a31be2abf5104115467aeeeedf8c0ebbfdb3d93e132ca03a6addefd12"></a>

## location property — slack.url.blindfold_secret_info / d978c2c7c169 / 5

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

<a id="canonical-139ce373580fa6337444e4188b48ae8d06b0b6d3921c37afdba873faa84bbd79"></a>

<a id="canonical-21bcddf085a54304e5a85e2437882c742f9fe8f357cac5f9b54f2c04f3608292"></a>

## store_provider property — slack.url.blindfold_secret_info / d978c2c7c169 / 6

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

<a id="canonical-c4c45e05075f7fd4b07ca4e8a4bf569a52c9842c9b195d56bd87495876647af8"></a>

## Next pages — slack.url.blindfold_secret_info / d978c2c7c169 / 7

- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-b3f57f9141adeccb5d712d48ac2928db06ef2d4d77d5ef95e62f2c45fc7ff00f)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-2dc26148c4c3e622e16372793e725190c27092c2f3be3c76163ef739b9ba7fca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbcf6d385708e77e1d2ec2de360c459a6ff77bd9985c05934e360b86b90df40e"></a>

## slack.url.clear_secret_info — slack.url.clear_secret_info / 98a83b9197d7 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [slack](resources--alert_receiver--reference--group-001.md#canonical-5eb2462c49807895168fa3c04b9823bd1e6412c90d7ac8213a8e7c7553d860c5)
- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-b3f57f9141adeccb5d712d48ac2928db06ef2d4d77d5ef95e62f2c45fc7ff00f)
- slack.url.clear_secret_info

<a id="canonical-1a65a7e2cbaceb15427e4cf5ad314ca4eae52eaf5d55213449df8ca163ded0e0"></a>

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

<a id="canonical-dfdec6c358e44dfb08935bd689347cf3c8e1804a5b46995491fac122bfcbe1eb"></a>

## Direct properties — slack.url.clear_secret_info / 98a83b9197d7 / 3

<a id="canonical-5cf1d29f0f1bbd720f26f1e0477d3a9eb940c63657b64a2a6347cfdb1c4a6a32"></a>

<a id="canonical-007a00ba6ebdc6f7a564d6806a3606b24610a71b272786beb00dedf12e478382"></a>

## provider_ref property — slack.url.clear_secret_info / 98a83b9197d7 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-fc2f00b68e3775de1d4d799a9fc9af0c7f94469f512b65a473efbec8b48048b1"></a>

<a id="canonical-83b411912d0dda462c9767e6020c5095b6c65aec1aa6e8da336471049c55a955"></a>

## url property — slack.url.clear_secret_info / 98a83b9197d7 / 5

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

<a id="canonical-32bf6b070195607c75edc0ead0c9be3bb369399aa7110ad52b76483b2dcad8fc"></a>

## Next pages — slack.url.clear_secret_info / 98a83b9197d7 / 6

- [slack.url](resources--alert_receiver--reference--group-001.md#canonical-b3f57f9141adeccb5d712d48ac2928db06ef2d4d77d5ef95e62f2c45fc7ff00f)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-b9ebc6bafb47a6223e03da567f4c7441fc83f2d6db862699593f75feb6f44fa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b21aae3866b49eac321cc5ba33e4de30d9ba7a06daa547f730aa75ba1aaf363"></a>

## sms — sms / 0e192ee56f8c / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- sms

<a id="canonical-818d75cc41c7615ffc51b808095b929afa600fbf24a1d15d007ac8680835dece"></a>

Type: `"object"`. single nested block, Optional.

SMS Configuration.

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
sms {
  # Configure direct properties listed below.
}
```

<a id="canonical-8441945766436d377321a1a2938c0ddc2a656ca8aa4aed9419c2668ab1d2611c"></a>

## Direct properties — sms / 0e192ee56f8c / 3

<a id="canonical-08afdce794c88332e83456874a7d218e0ba5d8ffc60c5a0b804a9e89cf3c4cd4"></a>

<a id="canonical-4e4b5ee78388039f65f48de16f49c14c149ed0b4e64a20f2bdff0957d76c981c"></a>

## contact_number property — sms / 0e192ee56f8c / 4

Type: `"string"`. Optional.

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\].

Upstream description:

Contact number of the user in ITU E.164 format \[+\]\[country code\]\[subscriber number including
area code\]

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
    "ves.io.schema.rules.string.phone_number": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.phone_number": "true"
  }
}
```

<a id="canonical-15dd7e9502326e9ae2a533bc8a3392656e774c16519e575e10aa755f1e5957df"></a>

## Next pages — sms / 0e192ee56f8c / 5

- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-4d698c906228b09a3dcf057c8ff469cf2ca0acb5241d136bad8e5ef1549afca9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecc6966fdeefba275464fa5220dece9248a23ecc1b139a3d48ad1f8a3e4cea5d"></a>

## timeouts — timeouts / b33283a67f15 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- timeouts

<a id="canonical-d51010d031ebfa17cc954fbd40ddb2d940de1b635c14c4b34113544aa9496280"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-968b6c06a5d2a2e5b5bb94e25fc4c3073f068202951dea891dae4f0767423979"></a>

## Direct properties — timeouts / b33283a67f15 / 3

<a id="canonical-9f57e544a4597ee910ea2dc5baeda562185dc57ef7ab7518be8e58e212bc6429"></a>

<a id="canonical-8bc64389b7c8cf0fcc5954e3a2bbb171e367eaf909b7d04b5abfb409d241810d"></a>

## create property — timeouts / b33283a67f15 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8380363a531107199203134fed44d17bc3f97eb23e3d71e551d057ac7eb71ce8"></a>

<a id="canonical-348dcec21af230f426cac5f949a43fbdaeda82d6c23f1cfbd229438442465bab"></a>

## delete property — timeouts / b33283a67f15 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-dfecf90f45b0d39e5a8f18762de54f2386a7168346e67760fdd4fc017e550279"></a>

<a id="canonical-f6661ac0729131cdd62600d99a6433327b21ceb8beaa4b77eff85b4b48b08bdc"></a>

## read property — timeouts / b33283a67f15 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-aad6c9e6dfd1f93f57c2ed3e153663d9adf7b8cb1f9008d2bb2843e2c9c9b499"></a>

<a id="canonical-b6a05694b6b6b453b280001a86d770f0a4cd19e3d712d4226c70e6abdbf8c4f4"></a>

## update property — timeouts / b33283a67f15 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7bde05191d74575ffef144872f49296003ac02cdb2098e11ce09c911fd3de9c5"></a>

## Next pages — timeouts / b33283a67f15 / 8

- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc120e5966fca84198841652bb9f899f5e3f589f3077c223ec71c38465ffc270"></a>

## webhook — webhook / 82c9ae40e14a / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- webhook

<a id="canonical-7ff5070337af3f8c75f04742c6019aba8c74123aa2a4c7d99940445a32db63b7"></a>

Type: `"object"`. single nested block, Optional.

Webhook configuration to send alert notifications.

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
webhook {
  # Configure direct properties listed below.
}
```

<a id="canonical-f36552f9fba01fa623fddc1ea3047eb44ef9fe0dd53a6032e04971eae5897749"></a>

## Direct properties — webhook / 82c9ae40e14a / 3

- [http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca): complete subsection reference.

- [url](resources--alert_receiver--reference--group-001.md#canonical-8f53b3d3defe000318d20a2bfb8fbcfed1e4c956d96d83913ab6c2568d1844d4): complete subsection reference.

<a id="canonical-51941ec387fbb762872170f812942d6c0d30e47e419ab903b81396bdf0586b88"></a>

## Next pages — webhook / 82c9ae40e14a / 4

- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-8f53b3d3defe000318d20a2bfb8fbcfed1e4c956d96d83913ab6c2568d1844d4)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-052667a46f49f2dd67267990f30dd4531c01eaf9ef2b49d265f5408d55d50e0a"></a>

## webhook.http_config — webhook.http_config / 162c21102224 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- webhook.http_config

<a id="canonical-8460ffda1e1d529ca784751fdb26c85d858dd234fb91b458721a96a09035c36b"></a>

Type: `"object"`. single nested block, Optional.

HTTP Configuration. Configuration for HTTP endpoint.

Upstream description:

Configuration for HTTP endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_token",
    "basic_auth"),
  validators.ConflictingObjectAttributes("auth_token",
    "client_cert_obj"),
  validators.ConflictingObjectAttributes("auth_token",
    "no_authorization"),
  validators.ConflictingObjectAttributes("basic_auth",
    "client_cert_obj"),
  validators.ConflictingObjectAttributes("basic_auth",
    "no_authorization"),
  validators.ConflictingObjectAttributes("client_cert_obj",
    "no_authorization"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-auth_choice": "[\"auth_token\",\"basic_auth\",\"client_cert_obj\",\"no_authorization\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
http_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-961d319801c388aa2905e6f03e06fb5d5575b76cc11152b7b3ccca5768866d8f"></a>

## Direct properties — webhook.http_config / 162c21102224 / 3

- [auth_token](resources--alert_receiver--reference--group-001.md#canonical-c8d5d9cd48adf846e4f23e55b28c810855cc376cbd805165b7d5c99b285843f6): complete subsection reference.

- [basic_auth](resources--alert_receiver--reference--group-001.md#canonical-a364201a82c4621cde1b507357ddff199d9062f317e4b912c6dc8cfa7fc3856f): complete subsection reference.

- [client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-017de9d074bf9aadcf47e6fdbd7e81bf6f30744a099f90d999f42b1ec522ee50): complete subsection reference.

<a id="canonical-e97e76520a3f422e976f50405eec11b0f09ca0d6ab4173275867d0937943a918"></a>

<a id="canonical-58989c62ef06ddcba460ce2e761bea987641a5c90a50548a9f3ccee5f1991d30"></a>

## enable_http2 property — webhook.http_config / 162c21102224 / 4

Type: `"bool"`. Optional.

Enable HTTP2. Configure to use HTTP2 protocol.

Upstream description:

Configure to use HTTP2 protocol.

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

<a id="canonical-8f037bc307cb8b8261989f987e9570950020c369e8487437bbdaa8dc42bb3223"></a>

<a id="canonical-a133693574445a8e15c45f73f6a7e990d7514bb6957693d7be03aacaf944523b"></a>

## follow_redirects property — webhook.http_config / 162c21102224 / 5

Type: `"bool"`. Optional.

Configure whether HTTP requests follow HTTP 3xx redirects.

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

- [no_authorization](resources--alert_receiver--reference--group-001.md#canonical-c1f131b5b866bf2c85c0cf9eee52a3185cb85f666732afbded8f9e5c8d6864ae): complete subsection reference.

- [no_tls](resources--alert_receiver--reference--group-001.md#canonical-fa500a2b821e84abb1d1e3145de4ce9ccbb97404c72a09a6133caa3b83d009dc): complete subsection reference.

- [use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367): complete subsection reference.

<a id="canonical-580026fe7fd0c2a4584af9fca4b5d562f42c6481d84125b5e99c1b9bacf571ab"></a>

## Next pages — webhook.http_config / 162c21102224 / 6

- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-c8d5d9cd48adf846e4f23e55b28c810855cc376cbd805165b7d5c99b285843f6)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-a364201a82c4621cde1b507357ddff199d9062f317e4b912c6dc8cfa7fc3856f)
- [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-017de9d074bf9aadcf47e6fdbd7e81bf6f30744a099f90d999f42b1ec522ee50)
- [webhook.http_config.no_authorization](resources--alert_receiver--reference--group-001.md#canonical-c1f131b5b866bf2c85c0cf9eee52a3185cb85f666732afbded8f9e5c8d6864ae)
- [webhook.http_config.no_tls](resources--alert_receiver--reference--group-001.md#canonical-fa500a2b821e84abb1d1e3145de4ce9ccbb97404c72a09a6133caa3b83d009dc)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-c8d5d9cd48adf846e4f23e55b28c810855cc376cbd805165b7d5c99b285843f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c4485f412c835169b911b9534d5b554d0d7d8bfbacebdecf033372123e25c71"></a>

## webhook.http_config.auth_token — webhook.http_config.auth_token / 3d50bb9c62a3 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- webhook.http_config.auth_token

<a id="canonical-47bfcf94135899c9fc64e8c5a5be54d1a3dee4020fd67233b5a05366a8013f2a"></a>

Type: `"object"`. single nested block, Optional.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

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
auth_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-313ab3bbda6c45f7f192002180eab3822b1be073b7459ce157f423c9c32f1033"></a>

## Direct properties — webhook.http_config.auth_token / 3d50bb9c62a3 / 3

- [token](resources--alert_receiver--reference--group-001.md#canonical-cda41fa632a26bb435acff8b34ff0228080598f61e79d8c5d3c42883f8d6850a): complete subsection reference.

<a id="canonical-ce0f47c9bd95f34098e95d6be4f1d774692dc769b1b133f3f37e1c9b93c01d02"></a>

## Next pages — webhook.http_config.auth_token / 3d50bb9c62a3 / 4

- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-cda41fa632a26bb435acff8b34ff0228080598f61e79d8c5d3c42883f8d6850a)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-cda41fa632a26bb435acff8b34ff0228080598f61e79d8c5d3c42883f8d6850a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef5853b0b04f2712ec9eef12a9371a648b4deba48b093c07702ebc0af24782a9"></a>

## webhook.http_config.auth_token.token — webhook.http_config.auth_token.token / 98b541f549c7 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-c8d5d9cd48adf846e4f23e55b28c810855cc376cbd805165b7d5c99b285843f6)
- webhook.http_config.auth_token.token

<a id="canonical-8ac484a30937470771d8ba5fc7bafe6d85f7c2a0daf42a1690e55f541b8ed670"></a>

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

<a id="canonical-97e1f275de268016fbc715e7a7af4980185dba0f138138c14108986b8c5c8080"></a>

## Direct properties — webhook.http_config.auth_token.token / 98b541f549c7 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-63df323fc10d441b51ef48e481870cc32156b972c98dfcace97a1d77b36b9354): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-bc8b5a53a7f44367a882d2191675856a512978fba30d6aba6c7c9b21438a9b6c): complete subsection reference.

<a id="canonical-84963a710c9caf3261eecf6ebd607af0645f8ed6fedbd05de0c898da8bbcb7b1"></a>

## Next pages — webhook.http_config.auth_token.token / 98b541f549c7 / 4

- [webhook.http_config.auth_token.token.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-63df323fc10d441b51ef48e481870cc32156b972c98dfcace97a1d77b36b9354)
- [webhook.http_config.auth_token.token.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-bc8b5a53a7f44367a882d2191675856a512978fba30d6aba6c7c9b21438a9b6c)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-c8d5d9cd48adf846e4f23e55b28c810855cc376cbd805165b7d5c99b285843f6)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-63df323fc10d441b51ef48e481870cc32156b972c98dfcace97a1d77b36b9354"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b0941e81b4b042f8171e35b1d1973fe7365b250bdee69080cb0df08e07c1089"></a>

## webhook.http_config.auth_token.token.blindfold_secret_info — webhook.http_config.auth_token.token.blindfold_secret_info / c1b932e54467 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-c8d5d9cd48adf846e4f23e55b28c810855cc376cbd805165b7d5c99b285843f6)
- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-cda41fa632a26bb435acff8b34ff0228080598f61e79d8c5d3c42883f8d6850a)
- webhook.http_config.auth_token.token.blindfold_secret_info

<a id="canonical-c44ee64301c895771d9c4ab20ca24abb2c941ebae718e79d7e817c72d15b2fed"></a>

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

<a id="canonical-f41f0a6bdc53b4ca9c99de240f0975f2157ef6d5cf97384e115dcb1cbd7691e7"></a>

## Direct properties — webhook.http_config.auth_token.token.blindfold_secret_info / c1b932e54467 / 3

<a id="canonical-3eabb976317008628418892a9f37f21f072d1d2b7e6b666556fd87e066697439"></a>

<a id="canonical-d83492bee04ac98bba4aafbf7ce7d7c7b444b847526018623d064bef3e74957e"></a>

## decryption_provider property — webhook.http_config.auth_token.token.blindfold_secret_info / c1b932e54467 / 4

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

<a id="canonical-e5e608f6ee30b4aafbd5f5be0f5ecabce74410754cc0e557beffc2ca90ba4f9b"></a>

<a id="canonical-13744f0aed5e2e2e8e8f8f94ae87785445969dcc54d4d5325d5426435668b1ba"></a>

## location property — webhook.http_config.auth_token.token.blindfold_secret_info / c1b932e54467 / 5

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

<a id="canonical-2cb3bcd010f1bbc66f6fa02aedad45051c7be7690ae192ea129f92f685e48286"></a>

<a id="canonical-29c33ab554f105e43437d6bf8fd1f81d885ada88fa21948a3253212079da97d1"></a>

## store_provider property — webhook.http_config.auth_token.token.blindfold_secret_info / c1b932e54467 / 6

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

<a id="canonical-ddf26455031d5df824ff2562d57e81abeb5979a5031ee74b8cb6c2e7de54cf18"></a>

## Next pages — webhook.http_config.auth_token.token.blindfold_secret_info / c1b932e54467 / 7

- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-cda41fa632a26bb435acff8b34ff0228080598f61e79d8c5d3c42883f8d6850a)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-bc8b5a53a7f44367a882d2191675856a512978fba30d6aba6c7c9b21438a9b6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92ae6bc31ca9fa383a0e7bc797699227bc40c56134f3993b257e64ab6cfecb9c"></a>

## webhook.http_config.auth_token.token.clear_secret_info — webhook.http_config.auth_token.token.clear_secret_info / 0455c4ae6feb / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.auth_token](resources--alert_receiver--reference--group-001.md#canonical-c8d5d9cd48adf846e4f23e55b28c810855cc376cbd805165b7d5c99b285843f6)
- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-cda41fa632a26bb435acff8b34ff0228080598f61e79d8c5d3c42883f8d6850a)
- webhook.http_config.auth_token.token.clear_secret_info

<a id="canonical-a5cbb0ec6197e3ebfef1350f26752e514559a7a8a1d133cd0001bf6811e36671"></a>

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

<a id="canonical-c751a579e7921152fe6ffc5fd61cdcbf4b7c5a574300143c1316fb46b2a5c882"></a>

## Direct properties — webhook.http_config.auth_token.token.clear_secret_info / 0455c4ae6feb / 3

<a id="canonical-20fea35d856fcf979433545b8d3c3f97d4273585bd3728276edd6bcdf56d6339"></a>

<a id="canonical-471df28c843955845aa85e866318b7abca256dde50b0ccd05f94db8c6abc8e5e"></a>

## provider_ref property — webhook.http_config.auth_token.token.clear_secret_info / 0455c4ae6feb / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-a3bae3dfa106b040b4105898937ea70d58da7649ca80dd45134a8e6357260e20"></a>

<a id="canonical-77669b008218363cd069fecce985ce059bd83620b371acc7a21f768eb13a8327"></a>

## url property — webhook.http_config.auth_token.token.clear_secret_info / 0455c4ae6feb / 5

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

<a id="canonical-9b94588bb884c4d703ec064c49b3d14f990e87bb88a1ac7317e2fe7ead6324ee"></a>

## Next pages — webhook.http_config.auth_token.token.clear_secret_info / 0455c4ae6feb / 6

- [webhook.http_config.auth_token.token](resources--alert_receiver--reference--group-001.md#canonical-cda41fa632a26bb435acff8b34ff0228080598f61e79d8c5d3c42883f8d6850a)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-a364201a82c4621cde1b507357ddff199d9062f317e4b912c6dc8cfa7fc3856f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1090a2a86e5e90834b70be82bdfb9f2c3de9c27fdae5ce318eb67750d0e092f"></a>

## webhook.http_config.basic_auth — webhook.http_config.basic_auth / d4edbe8321a6 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- webhook.http_config.basic_auth

<a id="canonical-9ee54ca45b085ce2ede8a84d81d682bca11777c53eaf38d18d0d7368d988ee9b"></a>

Type: `"object"`. single nested block, Optional.

Authorization parameters to access HTPP alert Receiver Endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("user_name")}
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

<a id="canonical-43416741a2350e55a94e8da1c32f5e3ed6c09e2817d6d485b90948ed2a905600"></a>

## Direct properties — webhook.http_config.basic_auth / d4edbe8321a6 / 3

- [password](resources--alert_receiver--reference--group-001.md#canonical-671032a2898d0ec0e41d1a0124b8f33b27141136ec00b4f58991f2e0255f302f): complete subsection reference.

<a id="canonical-36f64fc4e87e6cb6273f969fd39580da780b0918f7786921b16aad8172312242"></a>

<a id="canonical-874830260c881922b755f441123662dfe3e9840374237f1b1d698de0d4d7d47a"></a>

## user_name property — webhook.http_config.basic_auth / d4edbe8321a6 / 4

Type: `"string"`. Optional.

User Name. HTTP Basic Auth User Name.

Upstream description:

HTTP Basic Auth User Name.

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

<a id="canonical-ce72c1cbddddbc4aae12469066f123fd9d3a5f487fb0606a081b4ee2e8f35d93"></a>

## Next pages — webhook.http_config.basic_auth / d4edbe8321a6 / 5

- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-671032a2898d0ec0e41d1a0124b8f33b27141136ec00b4f58991f2e0255f302f)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-671032a2898d0ec0e41d1a0124b8f33b27141136ec00b4f58991f2e0255f302f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2575e651f1dbde433440abbcf1d80e9ecb5a988ca40f49964fb1becf448ded7"></a>

## webhook.http_config.basic_auth.password — webhook.http_config.basic_auth.password / 82b8604342e9 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-a364201a82c4621cde1b507357ddff199d9062f317e4b912c6dc8cfa7fc3856f)
- webhook.http_config.basic_auth.password

<a id="canonical-5d3dba706544beb0f0b94acb66cfde5dd12f7a982c959e86c62b8e6e103e4ca3"></a>

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

<a id="canonical-1977245e3df9bfe837493e74a39cf7d7013e1899024b5dd3c2111762ad5d2e62"></a>

## Direct properties — webhook.http_config.basic_auth.password / 82b8604342e9 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0f1605e014083681f329d7b609ea5acd604f4c957b6b1d47b31fae670fdf2948): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2e107bcd89907de682049f57cd9deb310af2cf0014dd3d33def36723fbeb31b3): complete subsection reference.

<a id="canonical-db22a2d714b663cecb4374e15e370d9991f48b6b61cfaa0e61b88330d7ad6a8e"></a>

## Next pages — webhook.http_config.basic_auth.password / 82b8604342e9 / 4

- [webhook.http_config.basic_auth.password.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-0f1605e014083681f329d7b609ea5acd604f4c957b6b1d47b31fae670fdf2948)
- [webhook.http_config.basic_auth.password.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-2e107bcd89907de682049f57cd9deb310af2cf0014dd3d33def36723fbeb31b3)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-a364201a82c4621cde1b507357ddff199d9062f317e4b912c6dc8cfa7fc3856f)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-0f1605e014083681f329d7b609ea5acd604f4c957b6b1d47b31fae670fdf2948"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-861bcfb2eeecaf663c6f83f29bc62915c926986513acdfafb2d30d7ed3110078"></a>

## webhook.http_config.basic_auth.password.blindfold_secret_info — webhook.http_config.basic_auth.password.blindfold_secret_info / 10be2d69c4f0 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-a364201a82c4621cde1b507357ddff199d9062f317e4b912c6dc8cfa7fc3856f)
- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-671032a2898d0ec0e41d1a0124b8f33b27141136ec00b4f58991f2e0255f302f)
- webhook.http_config.basic_auth.password.blindfold_secret_info

<a id="canonical-ecf2f5fb88efd9b4a24ee5a8e186323f5248fd79d0d2c7563a9722db7bff6d5a"></a>

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

<a id="canonical-c4cd88a3f39e7ee374e43c59724b509ebaee121ab98e89df8c34c36adeda3270"></a>

## Direct properties — webhook.http_config.basic_auth.password.blindfold_secret_info / 10be2d69c4f0 / 3

<a id="canonical-02bc2d854fda381fccae6a4b1d356b6628b9e8cfc04bcae547a32502cb029e61"></a>

<a id="canonical-25d0fa1dd136bb041ce95c0b69fab5486cdfc7362b69956c6d076a4b428491ba"></a>

## decryption_provider property — webhook.http_config.basic_auth.password.blindfold_secret_info / 10be2d69c4f0 / 4

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

<a id="canonical-b484674ea081e2403c9f7d65c50a5be7b8deb30049701e6efde275c5a052e92a"></a>

<a id="canonical-a7f2b6dcec1607f9e7aab76d86c8ff01f00ddc77943ab9135b5dc7111d7f930f"></a>

## location property — webhook.http_config.basic_auth.password.blindfold_secret_info / 10be2d69c4f0 / 5

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

<a id="canonical-22995b2adec3e39c806b4e6a16bd328cca493b5535fcefc519f8c9a9a242e41e"></a>

<a id="canonical-a45dfaf533898ea7bb8eddc3f258a6814116e50e61a25060283c70a1bbc29292"></a>

## store_provider property — webhook.http_config.basic_auth.password.blindfold_secret_info / 10be2d69c4f0 / 6

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

<a id="canonical-5435f9a2e9dbdf5b652adcde25724ced482beaebee285b838fb1a537eb8cb7ed"></a>

## Next pages — webhook.http_config.basic_auth.password.blindfold_secret_info / 10be2d69c4f0 / 7

- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-671032a2898d0ec0e41d1a0124b8f33b27141136ec00b4f58991f2e0255f302f)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-2e107bcd89907de682049f57cd9deb310af2cf0014dd3d33def36723fbeb31b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45972f238b4355a3279e5feb77a1ac54fa65f67755b7ea7a4eaa95aaee0d5997"></a>

## webhook.http_config.basic_auth.password.clear_secret_info — webhook.http_config.basic_auth.password.clear_secret_info / c5e0066b727a / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.basic_auth](resources--alert_receiver--reference--group-001.md#canonical-a364201a82c4621cde1b507357ddff199d9062f317e4b912c6dc8cfa7fc3856f)
- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-671032a2898d0ec0e41d1a0124b8f33b27141136ec00b4f58991f2e0255f302f)
- webhook.http_config.basic_auth.password.clear_secret_info

<a id="canonical-6fb3ed599d684198009398e557a0354895af040c1143394caa61e062191c0340"></a>

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

<a id="canonical-ef6448c63ba0dedc0935c4a876ca27eeb2380491dcb21af21884a6c23065a41d"></a>

## Direct properties — webhook.http_config.basic_auth.password.clear_secret_info / c5e0066b727a / 3

<a id="canonical-17eb314d25cc59a62a0d3ce32de0cea391ad31eecaad6d34c21c595b3ac0d8ec"></a>

<a id="canonical-8f9393c8c53b3faacf807f0aa938cb3c0ba722d5a080d29659d823a359d47a22"></a>

## provider_ref property — webhook.http_config.basic_auth.password.clear_secret_info / c5e0066b727a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c409e032d11bff461ba0adb56b4c12882ca45284ffdeaf04be8559d300c26e0f"></a>

<a id="canonical-1d45e6a6cc9598be297ad2ff91fa5439a9525b59854298b37c6fc11ea3ff41ea"></a>

## url property — webhook.http_config.basic_auth.password.clear_secret_info / c5e0066b727a / 5

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

<a id="canonical-c933c6ae749924ac63cbaf56342da28003403b003bf72161aeace8e58e22a665"></a>

## Next pages — webhook.http_config.basic_auth.password.clear_secret_info / c5e0066b727a / 6

- [webhook.http_config.basic_auth.password](resources--alert_receiver--reference--group-001.md#canonical-671032a2898d0ec0e41d1a0124b8f33b27141136ec00b4f58991f2e0255f302f)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-017de9d074bf9aadcf47e6fdbd7e81bf6f30744a099f90d999f42b1ec522ee50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3fee0a9417eb9cf172b1a8e9f502abdec1147b1f2249340a782fbaec105a0e3"></a>

## webhook.http_config.client_cert_obj — webhook.http_config.client_cert_obj / c3da152c31af / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- webhook.http_config.client_cert_obj

<a id="canonical-d786a3df5ed0f861c9027620fee05a67668fae06742f45bdc840abf06e487f3b"></a>

Type: `"object"`. single nested block, Optional.

Client Certificate Object. Configuration for client certificate.

Upstream description:

Configuration for client certificate.

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
client_cert_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-1bf57373b2bb222a8b93633aa7c382df086caf9c2c2e642a693d277085bab8c9"></a>

## Direct properties — webhook.http_config.client_cert_obj / c3da152c31af / 3

- [use_tls_obj](resources--alert_receiver--reference--group-001.md#canonical-f77f77edc21fb48612ee2fc680a840065b8b9a7db4e04dcef8666452eb2df7b4): complete subsection reference.

<a id="canonical-64c4d792907076be65f62501308e094710dea65ea410585f80c6aeecfc8e41ea"></a>

## Next pages — webhook.http_config.client_cert_obj / c3da152c31af / 4

- [webhook.http_config.client_cert_obj.use_tls_obj](resources--alert_receiver--reference--group-001.md#canonical-f77f77edc21fb48612ee2fc680a840065b8b9a7db4e04dcef8666452eb2df7b4)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-f77f77edc21fb48612ee2fc680a840065b8b9a7db4e04dcef8666452eb2df7b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-873d9568d7028c0242875a057fedab72e666cc01c21e7687b3c705636c7c61e5"></a>

## webhook.http_config.client_cert_obj.use_tls_obj — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-017de9d074bf9aadcf47e6fdbd7e81bf6f30744a099f90d999f42b1ec522ee50)
- webhook.http_config.client_cert_obj.use_tls_obj

<a id="canonical-649adc1d5d30f774c59cad4d62f84e982593a804f0473a54cb26144985923bf1"></a>

Type: `"object"`. list nested block, Optional.

Certificate Object. Reference to client certificate object.

Upstream description:

Reference to client certificate object.

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
use_tls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-00faa0ab63475ea043ebeba67d1ec3a77a04703277a78342cb86528fd1413165"></a>

## Direct properties — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 3

<a id="canonical-ab01f1630dc657da32b6b27a34a84606378f09e2f34355d0a3ba0dc1549e557e"></a>

<a id="canonical-536f9993473f123081ea60bb0609e2621b8673625c002abbb7113cbaa6cec70c"></a>

## kind property — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 4

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

<a id="canonical-19e402718a8448a70e5bfad95055574609d40bbbd10a65c2b8c2fa612b0e4277"></a>

<a id="canonical-83f6cc53c1ce34918e1604f11df2335b12c1d777f748aac3270e4e56195f92c4"></a>

## name property — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 5

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

<a id="canonical-407441b9089acde5736e8f1b0a41bdee7f882a260bc718a2a6ba376d1fe3063d"></a>

<a id="canonical-8ac7d444894f5ce93417622e365f33ffd55a28899a7495c9c441f715c513ebf8"></a>

## namespace property — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 6

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

<a id="canonical-027207ecd1952f39912e50ac9f124ed642d7e146371e63e499f1bb061a48befc"></a>

<a id="canonical-c3dd8505cae076b4f12b237f4ba43be47ceeb59ff371fa70d19b80372275e008"></a>

## tenant property — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 7

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

<a id="canonical-2cb755c86f034c8676e03756e1f528db89f07e9b1c4c23198faf239243c0beb2"></a>

<a id="canonical-fa56bc8ec65f6f2ddc1ab6c3ca4cbbd6ce59cd049f517e10f6331d336b6b2ae9"></a>

## uid property — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 8

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

<a id="canonical-1c204c3143b6b7a03b97a3225a222b93cc0f910e9cce4ae3c97d8932c5801a55"></a>

## Next pages — webhook.http_config.client_cert_obj.use_tls_obj / 2bc6851f2de0 / 9

- [webhook.http_config.client_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-017de9d074bf9aadcf47e6fdbd7e81bf6f30744a099f90d999f42b1ec522ee50)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-c1f131b5b866bf2c85c0cf9eee52a3185cb85f666732afbded8f9e5c8d6864ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40157bfba9fce47072ade57249df3a45be29c87c2da7a7c9291a8290a017f7e6"></a>

## webhook.http_config.no_authorization — webhook.http_config.no_authorization / a001d2d54c25 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- webhook.http_config.no_authorization

<a id="canonical-dac7ad8effc278e5ab98cf9fa2620b94fd589c9a410007fb0e4617d7f820c7bd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authorization.

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
no_authorization = {}
```

<a id="canonical-67411cdd23ce3c5bcf4d656cb17ac1054657fa60082a847aa7e48cff41899ff5"></a>

## Direct properties — webhook.http_config.no_authorization / a001d2d54c25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-209078145dce61a63422c3f9ea7ece5a711972d47714bef8dad3b71868ce2e74"></a>

## Next pages — webhook.http_config.no_authorization / a001d2d54c25 / 4

- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-fa500a2b821e84abb1d1e3145de4ce9ccbb97404c72a09a6133caa3b83d009dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-256656be54fc47afe660af783b7a4e4ddc6472a7be0a35ff85d0d741fa5daaac"></a>

## webhook.http_config.no_tls — webhook.http_config.no_tls / e5dba5bc86d5 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- webhook.http_config.no_tls

<a id="canonical-c03c5294cfcfda533ae0d8f19c20af73fda848d0e1d28d58477a0ba8da008b03"></a>

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
no_tls = {}
```

<a id="canonical-7c90227e5a3fc799c14d51f0b8aa0ba50cef7f954e7feb99b0b302f090a96eef"></a>

## Direct properties — webhook.http_config.no_tls / e5dba5bc86d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5dd93984e5417f3ecf5a25e57f477607143229d8a5cd8c583ff008c41e1012a"></a>

## Next pages — webhook.http_config.no_tls / e5dba5bc86d5 / 4

- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc384cb98f8b4da1c9495d0db54b8710f7e0c4e3f0146676d431cd510787d365"></a>

## webhook.http_config.use_tls — webhook.http_config.use_tls / 694b6f0dc2cb / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- webhook.http_config.use_tls

<a id="canonical-e30ec7dd9ea96fb2c10803fd4d5a235acf297aa1201f0c4fa44152d6add636f4"></a>

Type: `"object"`. single nested block, Optional.

Configures the token request's TLS settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("use_server_verification",
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
  "x-ves-oneof-field-server_validation_choice": "[\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-00bc414f89ab5c4974e3be5c518c75fc9b60543e949b028a6e0ae80072128063"></a>

## Direct properties — webhook.http_config.use_tls / 694b6f0dc2cb / 3

- [disable_sni](resources--alert_receiver--reference--group-001.md#canonical-3832f334473485bc7a63142d8f37605d9bc1c022db61e9f91ad03ae0d22aaa23): complete subsection reference.

<a id="canonical-9ed5e534ba7b00fb86c2721602f7bb6df00724450026471ed2d574434d097929"></a>

<a id="canonical-98d533af440e1b4f7873bb4bfde4f8249d5031a07d9aaca2a7c5751fbdb4b982"></a>

## max_version property — webhook.http_config.use_tls / 694b6f0dc2cb / 4

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

<a id="canonical-55558a20167ed63e4df3629ab3784779f338f0685fea64da01a8cca3efd302cd"></a>

<a id="canonical-f69c18461130a86d7f35264c2978d44dc98e9edfc674943828f70ffd68d02bed"></a>

## min_version property — webhook.http_config.use_tls / 694b6f0dc2cb / 5

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

<a id="canonical-53e5538791e98a40310da5ae85e133eb6aee70ef73e277743b745ae6e674287d"></a>

<a id="canonical-0e1c97d38a1cc087239a87eeff9a0b4ea5b5b0a550ac32bbd73942b3f82dc4bb"></a>

## sni property — webhook.http_config.use_tls / 694b6f0dc2cb / 6

Type: `"string"`. Optional.

Exclusive with \[disable\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni\] SNI value to be used.

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

- [use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-9da04c70dbe6e87ba788578b4722a7c0493fd138af4bb637719767e1e6b20c5f): complete subsection reference.

- [volterra_trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-edc2c09e5ab37c9fd5aeb38f4a5cdfe4344ff63691a418063d4a26cacc946952): complete subsection reference.

<a id="canonical-3654882dc25a93c03842f9a04b855c2d72de45cfff400d4e78d9e814dd56eddb"></a>

## Next pages — webhook.http_config.use_tls / 694b6f0dc2cb / 7

- [webhook.http_config.use_tls.disable_sni](resources--alert_receiver--reference--group-001.md#canonical-3832f334473485bc7a63142d8f37605d9bc1c022db61e9f91ad03ae0d22aaa23)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-9da04c70dbe6e87ba788578b4722a7c0493fd138af4bb637719767e1e6b20c5f)
- [webhook.http_config.use_tls.volterra_trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-edc2c09e5ab37c9fd5aeb38f4a5cdfe4344ff63691a418063d4a26cacc946952)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-3832f334473485bc7a63142d8f37605d9bc1c022db61e9f91ad03ae0d22aaa23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cc1a65b3c9e09884d194d241817e6f03952da9acb306dee5a40547118169d3f"></a>

## webhook.http_config.use_tls.disable_sni — webhook.http_config.use_tls.disable_sni / c8c90bca8d2c / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- webhook.http_config.use_tls.disable_sni

<a id="canonical-afa31575c1bbbd8cde912b94b852af147f6be50590b4f550b7ac9d15a6ca00b5"></a>

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

<a id="canonical-9bbd49d19bbf82c662c4f1cae627767d0c63a1e1d1a0168437fa53fedd3a0edb"></a>

## Direct properties — webhook.http_config.use_tls.disable_sni / c8c90bca8d2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9102474a9eef89009e3aaea26a966246a339b2750296363052ca39409920b779"></a>

## Next pages — webhook.http_config.use_tls.disable_sni / c8c90bca8d2c / 4

- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-9da04c70dbe6e87ba788578b4722a7c0493fd138af4bb637719767e1e6b20c5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84918ea99d60b359336fbc4aaa443d12a6f81f2b8212e412aa8309025485b193"></a>

## webhook.http_config.use_tls.use_server_verification — webhook.http_config.use_tls.use_server_verification / cc3d72fbdcde / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- webhook.http_config.use_tls.use_server_verification

<a id="canonical-5c17a1d4745d9298c95d8f4adeaa0cbef44def513b2507ee56fd88f98d3e726d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

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
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-4eb4751201a93ffe07064742ed2030518fc4d08fdcb765c96b8a7bc3014915cf"></a>

## Direct properties — webhook.http_config.use_tls.use_server_verification / cc3d72fbdcde / 3

- [ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-43c0bf3651b2461eefe62bb4c1f7c74efb02a34bc83439a24a8b17150c59ff09): complete subsection reference.

<a id="canonical-79e1d2f2609d65caadd6c835d2026bf45101f0a4a0c64e4152b0e78cfa610dde"></a>

## Next pages — webhook.http_config.use_tls.use_server_verification / cc3d72fbdcde / 4

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-43c0bf3651b2461eefe62bb4c1f7c74efb02a34bc83439a24a8b17150c59ff09)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-43c0bf3651b2461eefe62bb4c1f7c74efb02a34bc83439a24a8b17150c59ff09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90747069e1033374c058f8364be5c86abf07d5b1d0d4114afcd48a0744b54241"></a>

## webhook.http_config.use_tls.use_server_verification.ca_cert_obj — webhook.http_config.use_tls.use_server_verification.ca_cert_obj / bdd0f2c32265 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-9da04c70dbe6e87ba788578b4722a7c0493fd138af4bb637719767e1e6b20c5f)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj

<a id="canonical-acbe8542d41360cfe8d873e1601eff6c94b2f51e34ec57c2fdfe5ea2012155aa"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ca cert obj.

Upstream description:

Configuration for CA certificate.

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
ca_cert_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7d67528309f5c31118d55194cb48c066035e32da7d03675d419e9ee90209c5f"></a>

## Direct properties — webhook.http_config.use_tls.use_server_verification.ca_cert_obj / bdd0f2c32265 / 3

- [trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-7d86ae073ebbea8d9a29ab2130856d03f7481273c9aab32165952421806350da): complete subsection reference.

<a id="canonical-40823480d01e5b05933fab11c682487f35f830b93b5f005c9c88070b3af45508"></a>

## Next pages — webhook.http_config.use_tls.use_server_verification.ca_cert_obj / bdd0f2c32265 / 4

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](resources--alert_receiver--reference--group-001.md#canonical-7d86ae073ebbea8d9a29ab2130856d03f7481273c9aab32165952421806350da)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-9da04c70dbe6e87ba788578b4722a7c0493fd138af4bb637719767e1e6b20c5f)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-7d86ae073ebbea8d9a29ab2130856d03f7481273c9aab32165952421806350da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88dedbda4e59a8d9674aed473fa0b10986fd14fea7af5e8b4528f8b8cf9e71f5"></a>

## webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--reference--group-001.md#canonical-9da04c70dbe6e87ba788578b4722a7c0493fd138af4bb637719767e1e6b20c5f)
- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-43c0bf3651b2461eefe62bb4c1f7c74efb02a34bc83439a24a8b17150c59ff09)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca

<a id="canonical-5520a8b763a0040db998b6bd2c0f57f8cfc0ea2536a9c575c52f3a7d8accc181"></a>

Type: `"object"`. list nested block, Optional.

Certificate Object. Reference to client certificate object.

Upstream description:

Reference to client certificate object.

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-673d9fa02a5763d8fdc20cebeb254c7dcdd5c2004f1f1b2ed2995c9c7918db9a"></a>

## Direct properties — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 3

<a id="canonical-a73e7ab7b9217e0577836c9a323541ad18ccb9826a3568f891fd98a8de713627"></a>

<a id="canonical-eb933dc8b0f0266189e0d6477aa14b7beef8997357d554bc23273514b4b6f601"></a>

## kind property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 4

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

<a id="canonical-6856aaccef1b7bd2bf6b12ddbaed2efc188d588cec133aa741ad89569e9e0ccb"></a>

<a id="canonical-749672a629da223407648093838aa8422e56d8e3358285f235f0e0a2e18eaff7"></a>

## name property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 5

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

<a id="canonical-bfae57937a16b43893442c5891bf81de5a8642291feb6c2084ad8c213d1151ef"></a>

<a id="canonical-1f373bf6425fcc3e7a7425510b9e44a71f5f383fbdc348b6c7a15c1bddb2d48a"></a>

## namespace property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 6

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

<a id="canonical-52ebc2bc312b470276e1785dba2745a4a08a2efcd359a0ec7f4bd4dfbb92dbb9"></a>

<a id="canonical-efac0d108e0da38e58861c943b2822c37e0b8ed758305498cd3c075ef40cc746"></a>

## tenant property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 7

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

<a id="canonical-9fb782087788f01acb4317f578c97fd1621dddfdfb40126b7f28508beeab0b95"></a>

<a id="canonical-fb6b988968a60a353415e6bcc0fb34f6e73c73be7bd2444d5ca4ae9336761aea"></a>

## uid property — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 8

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

<a id="canonical-e1f026243af20b0d08d37c615079382aedbd02fab9b8f1e837c1a57e81021754"></a>

## Next pages — webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca / a1b2bf99e825 / 9

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](resources--alert_receiver--reference--group-001.md#canonical-43c0bf3651b2461eefe62bb4c1f7c74efb02a34bc83439a24a8b17150c59ff09)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-edc2c09e5ab37c9fd5aeb38f4a5cdfe4344ff63691a418063d4a26cacc946952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-362d4964a77cde911c0d863dd4c94503f7eb2d425540e068b482c96616b01e85"></a>

## webhook.http_config.use_tls.volterra_trusted_ca — webhook.http_config.use_tls.volterra_trusted_ca / 1fb1f103a216 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.http_config](resources--alert_receiver--reference--group-001.md#canonical-4c8315b92c996648f3eff070110f17d19b6cea140654bc85bca87712d33854ca)
- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- webhook.http_config.use_tls.volterra_trusted_ca

<a id="canonical-e4985a4ddb855154bbe168acee47c808373202310814c7ecaea1eec56432f912"></a>

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

<a id="canonical-f20141bb9fc8beec512703ce5c2070a952e8e70ba3e4895e6085eec9f9dbd9e6"></a>

## Direct properties — webhook.http_config.use_tls.volterra_trusted_ca / 1fb1f103a216 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ad59b8f14b62ecb6fa279715ac44bac58194785efb988797970271129d3bc4ea"></a>

## Next pages — webhook.http_config.use_tls.volterra_trusted_ca / 1fb1f103a216 / 4

- [webhook.http_config.use_tls](resources--alert_receiver--reference--group-001.md#canonical-af81d88b9f251016cf43c9374acc99033f007e031c1c90104aa039f9eb194367)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-8f53b3d3defe000318d20a2bfb8fbcfed1e4c956d96d83913ab6c2568d1844d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba1713056f9bc4302f2a3b56861ae1076532e4d4495d937379b0bce8f41b1f41"></a>

## webhook.url — webhook.url / b1b74bc98d50 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- webhook.url

<a id="canonical-201a2eb8601bd06a5248dfbde10a1cfa4a9aa4da45b9f447c97e144eb62bdb33"></a>

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
url {
  # Configure direct properties listed below.
}
```

<a id="canonical-2735f809cc3c5d4685abdec4682fbc7c0065101837abdef5d08a679961b62e54"></a>

## Direct properties — webhook.url / b1b74bc98d50 / 3

- [blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-394bc295b38809875ed5ffe90029668a1452d172167517460e738e89432563c3): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-88e54669c7c01e113b5af839f0a2ec06a7f04bf862ea04a2a67875590c5ef264): complete subsection reference.

<a id="canonical-211813b1ca7d2fe16097ce39df3fd8893ae3ad9088189e2b45a4507ae5c38f26"></a>

## Next pages — webhook.url / b1b74bc98d50 / 4

- [webhook.url.blindfold_secret_info](resources--alert_receiver--reference--group-001.md#canonical-394bc295b38809875ed5ffe90029668a1452d172167517460e738e89432563c3)
- [webhook.url.clear_secret_info](resources--alert_receiver--reference--group-001.md#canonical-88e54669c7c01e113b5af839f0a2ec06a7f04bf862ea04a2a67875590c5ef264)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-394bc295b38809875ed5ffe90029668a1452d172167517460e738e89432563c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-980d93e55b5de98c92622dc32f710f1b6f83859c0f5095dcd59b56f19acf1a68"></a>

## webhook.url.blindfold_secret_info — webhook.url.blindfold_secret_info / 6371ee8da044 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-8f53b3d3defe000318d20a2bfb8fbcfed1e4c956d96d83913ab6c2568d1844d4)
- webhook.url.blindfold_secret_info

<a id="canonical-25b8b44dd2b44a1948216eb8fec3578adaad92f17e08f91c9e0083779d1938b1"></a>

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

<a id="canonical-27634fd87060809aa01d46aeab65b11433d6c87414b5d105c1d855b088178aeb"></a>

## Direct properties — webhook.url.blindfold_secret_info / 6371ee8da044 / 3

<a id="canonical-922999a21f723853e56393d2fa54da37aa99d1d818b1fb22e63c4d8f33774c68"></a>

<a id="canonical-e85fc30cadcb21b839c1d69d56eeb75d0d0e86edbf676b1cd58f6517a15af3c0"></a>

## decryption_provider property — webhook.url.blindfold_secret_info / 6371ee8da044 / 4

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

<a id="canonical-b46a6877dee9e1bb25c9dbfbd86f86cbaa660b5c0d0e5a48a95e61ae026487e9"></a>

<a id="canonical-bafde4d87d85a3143e087f18386bb241d532e69161c7d040552074be1bfe9a01"></a>

## location property — webhook.url.blindfold_secret_info / 6371ee8da044 / 5

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

<a id="canonical-e4d32cb4c5a8f723ff762beece7379ce734bfe51d9f0a3f69005d4820ed249bb"></a>

<a id="canonical-c5eb6d79a2eeb330fc35568fbb05e73855a3dcc25dd17d955e8285c2137d953c"></a>

## store_provider property — webhook.url.blindfold_secret_info / 6371ee8da044 / 6

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

<a id="canonical-793a636e4c34832505dbda7758dc80dfc7e43e48dea8c5daf6be6ca062ede9ad"></a>

## Next pages — webhook.url.blindfold_secret_info / 6371ee8da044 / 7

- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-8f53b3d3defe000318d20a2bfb8fbcfed1e4c956d96d83913ab6c2568d1844d4)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-88e54669c7c01e113b5af839f0a2ec06a7f04bf862ea04a2a67875590c5ef264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-955f2def0b4566a4a9100cd1389ac1f0a0fe32744066ed18f2bce121340fa459"></a>

## webhook.url.clear_secret_info — webhook.url.clear_secret_info / 4c7bda824a54 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Property reference](resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [webhook](resources--alert_receiver--reference--group-001.md#canonical-230c90d6bd8132f21a7255c651f55966ffea0d4c681723cb30e3d6890168d693)
- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-8f53b3d3defe000318d20a2bfb8fbcfed1e4c956d96d83913ab6c2568d1844d4)
- webhook.url.clear_secret_info

<a id="canonical-39bf30a429874541d8d870bed6e8d6e6afa368a5a10c22a9fd97594d07e6ba18"></a>

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

<a id="canonical-01d882ef928ad2b56e95e862b2420dcb85cc6edbefae76c7cf0ffde62cab38ce"></a>

## Direct properties — webhook.url.clear_secret_info / 4c7bda824a54 / 3

<a id="canonical-660966b3de3191a31d59dbf32eac47b1fed1f198aacb5b3c2c08893ebffb5843"></a>

<a id="canonical-8479c980b57ce68adcd1010b44c6015c7fa5f0075d2c65d8e1f84500e0a7b387"></a>

## provider_ref property — webhook.url.clear_secret_info / 4c7bda824a54 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2ed52239b58c64a5c172748317889bf3c6407392123738c7ede2a7c2bb547aab"></a>

<a id="canonical-d4d153a27642d5d27729d5bd32eb528792db62e64fac4891c70b9addbc4e6666"></a>

## url property — webhook.url.clear_secret_info / 4c7bda824a54 / 5

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

<a id="canonical-cfcdc593035d2e23f15e6c4c424bac47e047d73c7766e227d76e7c521045ed41"></a>

## Next pages — webhook.url.clear_secret_info / 4c7bda824a54 / 6

- [webhook.url](resources--alert_receiver--reference--group-001.md#canonical-8f53b3d3defe000318d20a2bfb8fbcfed1e4c956d96d83913ab6c2568d1844d4)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
