---
page_title: "xcsh_alert_policy reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy reference."
---

# xcsh_alert_policy reference

<a id="canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68a14858dc819472709412f81594b348af7e5d5cbb896db546d78ee480128663"></a>

## Property reference — Property reference / 35fc36ba6c27 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- Property reference

<a id="canonical-8e45b24170d5b1c4c28a81390c28ed05586d0d735a39c6346ae95f20f42589af"></a>

## Direct properties — Property reference / 35fc36ba6c27 / 3

<a id="canonical-e84f7535f08cf6400c2a856bb6c4a876985e573d2ebd6279a40026c1d6504e3b"></a>

<a id="canonical-30086fd840791f8ef0eb5ccbc3076e2549296ab0294bc647f2809ad165da7a2e"></a>

## annotations property — Property reference / 35fc36ba6c27 / 4

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

<a id="canonical-3469c7b1c679d7b5b80aca2a77bdf5f8917a8f1f47fd25277b29863532045999"></a>

<a id="canonical-15a762e4e7d9f2e48fb3f15a55cc72fb032471f33aec5379aaee6da8cc1c46f6"></a>

## description property — Property reference / 35fc36ba6c27 / 5

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

<a id="canonical-955ba6b5a77eb86fc953d27f667ca4173e8e8dab1ff7eeee2631088b310a9801"></a>

<a id="canonical-babc1249fab3dbb67f24426e1170c050d08bdff3279c307c2b3c7d4423296010"></a>

## disable property — Property reference / 35fc36ba6c27 / 6

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

<a id="canonical-e05b6f0003195fdf70c779bc0b87e976894864e5da77a551dd7710b1ee4fc3c5"></a>

<a id="canonical-395fe3dc79fa9a5cbaa49f331354a5bf7a113baf1830d7e2f147094df67b89b0"></a>

## id property — Property reference / 35fc36ba6c27 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4366ce6339b58f47a1a60c54e574acbd1ae614f23c7e01ea22dc1117e31f89b4"></a>

<a id="canonical-77eb34917f25339b0afcebf7a07afba5100b6d7bc961152aaecbc680e8b6dbc8"></a>

## labels property — Property reference / 35fc36ba6c27 / 8

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

<a id="canonical-362623d5a957a7d173872686cc2846c80732d1361e122a6365b9988594813470"></a>

<a id="canonical-f89695ae529211187761711338aa696ce9a1a08a56495793ca51cb127c0103f1"></a>

## name property — Property reference / 35fc36ba6c27 / 9

Type: `"string"`. Required.

Name of the Alert Policy. Must be unique within the namespace.

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

<a id="canonical-7f04d65394d46bfe33fc2940a75f6732623e620266ea173018603746d668604f"></a>

<a id="canonical-c384afc731c92fbe7e8217d0fa9e340472bc10d2dcac33c918c6b9ad305da633"></a>

## namespace property — Property reference / 35fc36ba6c27 / 10

Type: `"string"`. Required.

Namespace where the Alert Policy is created.

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

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989): complete subsection reference.

- [receivers](resources--alert_policy--reference--group-001.md#canonical-ddebc42c9d6a73646d69696ac8c1305a9c894c0c0109a249a4223cf69593e762): complete subsection reference.

- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f): complete subsection reference.

- [timeouts](resources--alert_policy--reference--group-001.md#canonical-2ed7e2c972f5b95c7ffdc6ac774fd54ef2b45c05425cc0c15c7db6a7bcc7d7d9): complete subsection reference.

<a id="canonical-5647396ac39aacab652ba43c57c93fa906dce6c979833df3450939e5117dfd7c"></a>

## All schema paths — Property reference / 35fc36ba6c27 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--alert_policy--reference--group-001.md#canonical-e84f7535f08cf6400c2a856bb6c4a876985e573d2ebd6279a40026c1d6504e3b) |
| `description` | [description](resources--alert_policy--reference--group-001.md#canonical-3469c7b1c679d7b5b80aca2a77bdf5f8917a8f1f47fd25277b29863532045999) |
| `disable` | [disable](resources--alert_policy--reference--group-001.md#canonical-955ba6b5a77eb86fc953d27f667ca4173e8e8dab1ff7eeee2631088b310a9801) |
| `id` | [id](resources--alert_policy--reference--group-001.md#canonical-e05b6f0003195fdf70c779bc0b87e976894864e5da77a551dd7710b1ee4fc3c5) |
| `labels` | [labels](resources--alert_policy--reference--group-001.md#canonical-4366ce6339b58f47a1a60c54e574acbd1ae614f23c7e01ea22dc1117e31f89b4) |
| `name` | [name](resources--alert_policy--reference--group-001.md#canonical-362623d5a957a7d173872686cc2846c80732d1361e122a6365b9988594813470) |
| `namespace` | [namespace](resources--alert_policy--reference--group-001.md#canonical-7f04d65394d46bfe33fc2940a75f6732623e620266ea173018603746d668604f) |
| `notification_parameters` | [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-cf879c5adcc066280dfab0cec63f6c4bc7e56bc3aac397021707aaaa8f5eba3b) |
| `notification_parameters.custom` | [notification_parameters.custom](resources--alert_policy--reference--group-001.md#canonical-e6d3b452c59669e10a32d4923be4c7a33168cc99a054dd92a758eac5077f56d2) |
| `notification_parameters.custom.labels` | [notification_parameters.custom.labels](resources--alert_policy--reference--group-001.md#canonical-e70d4a4b962243b2d3893288ce5689dc104c9db674b75611e9ee9292f0f4462a) |
| `notification_parameters.default` | [notification_parameters.default](resources--alert_policy--reference--group-001.md#canonical-fa3255dd7872528f1a7d68e8746a874d8e97fed47e59d7aeb5991e19a481edca) |
| `notification_parameters.group_interval` | [notification_parameters.group_interval](resources--alert_policy--reference--group-001.md#canonical-d0caec1c5ad9192566af69d525cd3a5d62df21359ced8bf9b152f29980a56354) |
| `notification_parameters.group_wait` | [notification_parameters.group_wait](resources--alert_policy--reference--group-001.md#canonical-8707babea75f8604583bb7cb549b26b3ba024e154c755dc501bcaac2d67ba1ad) |
| `notification_parameters.individual` | [notification_parameters.individual](resources--alert_policy--reference--group-001.md#canonical-ce5152bf4be29ad0baba65249a63bb6d0a7b1cad0daff75584ef8cad4a2d4e23) |
| `notification_parameters.repeat_interval` | [notification_parameters.repeat_interval](resources--alert_policy--reference--group-001.md#canonical-969969b5ee7a1acbb3e05dcffb203e1482935e0070b2763738de4ce3979c0017) |
| `notification_parameters.ves_io_group` | [notification_parameters.ves_io_group](resources--alert_policy--reference--group-001.md#canonical-c3b4b033d118dd3e3d77ee9a1fa3580f14efa5a9762d39bcad901dce5633a2ea) |
| `receivers` | [receivers](resources--alert_policy--reference--group-001.md#canonical-ca502c6677ddf0f1487cf190dfe2cbf56d2f523a03cc6e5b2a5310d6af92eeee) |
| `receivers.kind` | [receivers.kind](resources--alert_policy--reference--group-001.md#canonical-456e4fa7ae4c5b202063b051474c10b39a1c885f763c68ff8a85b94d74465a14) |
| `receivers.name` | [receivers.name](resources--alert_policy--reference--group-001.md#canonical-7e3114dd636151bcd89d6924f31af3c92f07af784231d8006378d5088f4394f9) |
| `receivers.namespace` | [receivers.namespace](resources--alert_policy--reference--group-001.md#canonical-63f4df31ba49dae6ba8c186cac4bf3cb2e998a5b4db2299734adac0ca025b22d) |
| `receivers.tenant` | [receivers.tenant](resources--alert_policy--reference--group-001.md#canonical-49f872180287e28ca1d330864dc12503c4321a27d498ba36a451b6907b0bb69e) |
| `receivers.uid` | [receivers.uid](resources--alert_policy--reference--group-001.md#canonical-73e44f999ea49c85ef0772ca21bfa5cae69f7de21ce109b6a51033b7012ff724) |
| `routes` | [routes](resources--alert_policy--reference--group-001.md#canonical-31e6b78e36589cd87c0dc813534cfc61f453140ffe8aa7f3b20a55e545010fa0) |
| `routes.alertname` | [routes.alertname](resources--alert_policy--reference--group-001.md#canonical-38356d65ce2bea9fd65ab473a35e7adae52b18bf5e68157c7c7483567ab2eeb2) |
| `routes.alertname_regex` | [routes.alertname_regex](resources--alert_policy--reference--group-001.md#canonical-5240f736df355c74ec9da635ef6e30153d1014e2885eec5edd1dadd806dc5133) |
| `routes.any` | [routes.any](resources--alert_policy--reference--group-001.md#canonical-82078830fd6e83172223bef3e40c7b0bea09d956fef9a5b821265811c48614f9) |
| `routes.custom` | [routes.custom](resources--alert_policy--reference--group-001.md#canonical-19c17b0ef4872eb5199067f166a1ef5ae22093f0b11602f7ecc3b65c8e650d03) |
| `routes.custom.alertlabel` | [routes.custom.alertlabel](resources--alert_policy--reference--group-001.md#canonical-cd2f000b3535422b0ff9cfff95fe46ccf725bfe8bc16b5abd7e5d78e5052d583) |
| `routes.custom.alertname` | [routes.custom.alertname](resources--alert_policy--reference--group-001.md#canonical-bf457c276336c727b69c82c8abe36eb900ebb1b4842555ddb592cc6e6f372266) |
| `routes.custom.alertname.exact_match` | [routes.custom.alertname.exact_match](resources--alert_policy--reference--group-001.md#canonical-8f58b18ec20bbb4513718c71e707b55712b2f1af8329a46cb9a6add2c5c2d260) |
| `routes.custom.alertname.regex_match` | [routes.custom.alertname.regex_match](resources--alert_policy--reference--group-001.md#canonical-50ae7d040a9bdd9103e0339615baf15b539f91c9d40157dbaf0d02ea3a92ae4c) |
| `routes.custom.group` | [routes.custom.group](resources--alert_policy--reference--group-001.md#canonical-756da3439d317c3b8a77df6772bf2c00765dd50338caa6a3ecbf55fb38a6cb48) |
| `routes.custom.group.exact_match` | [routes.custom.group.exact_match](resources--alert_policy--reference--group-001.md#canonical-ec75a716e6368f4a6fbdacf3e374daaf9696e84d631ff05796c52070b05e88fe) |
| `routes.custom.group.regex_match` | [routes.custom.group.regex_match](resources--alert_policy--reference--group-001.md#canonical-53d1e621967dfa33d772c1940ef1d1bfdc12a4864fa33fadffedeef4d6621752) |
| `routes.custom.severity` | [routes.custom.severity](resources--alert_policy--reference--group-001.md#canonical-cabea6ebae22d75eed65c56135620838b071ded16869f7ecbe6deaea494e325a) |
| `routes.custom.severity.exact_match` | [routes.custom.severity.exact_match](resources--alert_policy--reference--group-001.md#canonical-907689b35f84d598c131e17974dcdce516f14550fbfe6b034aadbdfe9cb9a392) |
| `routes.custom.severity.regex_match` | [routes.custom.severity.regex_match](resources--alert_policy--reference--group-001.md#canonical-ec79e551c698861b047422ec5e1d8cc655e9b95167753a8c9fe87f547dac0c75) |
| `routes.dont_send` | [routes.dont_send](resources--alert_policy--reference--group-001.md#canonical-0703143905a8bae78a865ac9e98b6bb0b56556ad954f91c6ea6c1197f3435936) |
| `routes.group` | [routes.group](resources--alert_policy--reference--group-001.md#canonical-02f8e90efed32ad005f88a0b9df8db4dea1451d789175cbfcc019a1664c92138) |
| `routes.group.groups` | [routes.group.groups](resources--alert_policy--reference--group-001.md#canonical-545ba5a1aa3886d784e76fd78dd468342efbaaa0185c4c20d4ed5e1ac8720ab5) |
| `routes.notification_parameters` | [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-ff0b26198dab0e39efba9ac00ad38e30139b438c2333239e71f9c77df467d513) |
| `routes.notification_parameters.custom` | [routes.notification_parameters.custom](resources--alert_policy--reference--group-001.md#canonical-2f627e920348b8527f88fb691d96c091e21ece6e50a2070b37637f3b858dc5f2) |
| `routes.notification_parameters.custom.labels` | [routes.notification_parameters.custom.labels](resources--alert_policy--reference--group-001.md#canonical-68671e96323c0e5e50e61d97982b5dd45899b1800ae3465557c08c5ab9ea87b7) |
| `routes.notification_parameters.default` | [routes.notification_parameters.default](resources--alert_policy--reference--group-001.md#canonical-460ba4aa49d100ba139863b08fa5b2013a439832a8b973ec35e958cab9815be2) |
| `routes.notification_parameters.group_interval` | [routes.notification_parameters.group_interval](resources--alert_policy--reference--group-001.md#canonical-c04ce99080b20ae5caf199760d83eef1bfd11f02b26a2dd5f85312eafc525bdb) |
| `routes.notification_parameters.group_wait` | [routes.notification_parameters.group_wait](resources--alert_policy--reference--group-001.md#canonical-47bf5c10d7a6c5bb91984672e31e8a62b62bf6a73eb5ee81c1242d8dae069611) |
| `routes.notification_parameters.individual` | [routes.notification_parameters.individual](resources--alert_policy--reference--group-001.md#canonical-7d072597343cda700082ebb1514bd09afaeaab273a9d4d490706a299380546d4) |
| `routes.notification_parameters.repeat_interval` | [routes.notification_parameters.repeat_interval](resources--alert_policy--reference--group-001.md#canonical-c78f96396893376273e675e21ec4df97d94b3873d659c1263438ddc859ed3966) |
| `routes.notification_parameters.ves_io_group` | [routes.notification_parameters.ves_io_group](resources--alert_policy--reference--group-001.md#canonical-b73c816c163e9433e7c99fe7e7ddba9184217b40c118ba7c319e9ff791f8e330) |
| `routes.send` | [routes.send](resources--alert_policy--reference--group-001.md#canonical-2bdc08ee42ec16d8ca260de40ba127ef8decc4d83f997d0c45cc2a2a0218af48) |
| `routes.severity` | [routes.severity](resources--alert_policy--reference--group-001.md#canonical-305b9fbc9efa5fc86135a53aa5528031af976b1b387dc2dbcb69bb520dde5c31) |
| `routes.severity.severities` | [routes.severity.severities](resources--alert_policy--reference--group-001.md#canonical-cf0993fd803a4db468277015356010b6a6a2fca4c4de25327148af42d670cfc9) |
| `timeouts` | [timeouts](resources--alert_policy--reference--group-001.md#canonical-d490587c4bcc0e9baf19af72c5609850e1bc405c2af382d498ddb8e2d35ac595) |
| `timeouts.create` | [timeouts.create](resources--alert_policy--reference--group-001.md#canonical-f68ef5375777433fae28907ba5e8307aa2d5c67f9926481c1d7f05c5019fb907) |
| `timeouts.delete` | [timeouts.delete](resources--alert_policy--reference--group-001.md#canonical-33328e550fee5d12ecfcf158c1e09df2c1c8f95499bbe1e3f84d6df3d8e23620) |
| `timeouts.read` | [timeouts.read](resources--alert_policy--reference--group-001.md#canonical-62ecbf8efde71fd46015dd656fede8a255939ac0a2a723b26c83098ec211734c) |
| `timeouts.update` | [timeouts.update](resources--alert_policy--reference--group-001.md#canonical-c23df81a4da2430950004b118f5ba9488daed13bba0d0cd2876f854ada445147) |

<a id="canonical-79f1f696a727ca29b2df8e0b53e3a15ccd9bc7bedbfb4cb367902e670e7d41fd"></a>

## Next pages — Property reference / 35fc36ba6c27 / 12

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- [receivers](resources--alert_policy--reference--group-001.md#canonical-ddebc42c9d6a73646d69696ac8c1305a9c894c0c0109a249a4223cf69593e762)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [timeouts](resources--alert_policy--reference--group-001.md#canonical-2ed7e2c972f5b95c7ffdc6ac774fd54ef2b45c05425cc0c15c7db6a7bcc7d7d9)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7c3b25cc438515f0b1524d0ae0064e2148afb282c8a66b603fa8c16f19277a3"></a>

## notification_parameters — notification_parameters / 82679e48d075 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- notification_parameters

<a id="canonical-cf879c5adcc066280dfab0cec63f6c4bc7e56bc3aac397021707aaaa8f5eba3b"></a>

Type: `"object"`. single nested block, Optional.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom",
    "default"),
  validators.ConflictingObjectAttributes("custom",
    "individual"),
  validators.ConflictingObjectAttributes("custom",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("default",
    "individual"),
  validators.ConflictingObjectAttributes("default",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("individual",
    "ves_io_group")}
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
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

Terraform syntax:

```terraform
notification_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-057c9f162a644132880ef7050528798b3f6d0589aaad525927f912247ac462e2"></a>

## Direct properties — notification_parameters / 82679e48d075 / 3

- [custom](resources--alert_policy--reference--group-001.md#canonical-d7fd006e619b7ccd1949ffb5ad480cb68f17b550fe0bbb9849890998fd4d94ca): complete subsection reference.

- [default](resources--alert_policy--reference--group-001.md#canonical-9b631c58d71278f7a8922e12328635927f8dc1bbd6138bd6126b8223c3b5e10e): complete subsection reference.

<a id="canonical-d0caec1c5ad9192566af69d525cd3a5d62df21359ced8bf9b152f29980a56354"></a>

<a id="canonical-cbcf5c4b2f622835d4e64e2462dd88f737251d26b7e8cd91d2244258539604c3"></a>

## group_interval property — notification_parameters / 82679e48d075 / 4

Type: `"string"`. Optional.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-8707babea75f8604583bb7cb549b26b3ba024e154c755dc501bcaac2d67ba1ad"></a>

<a id="canonical-bf42943722b7217aae5e57cfa88597c28b704f9a166df5fe78a4525ce1d64787"></a>

## group_wait property — notification_parameters / 82679e48d075 / 5

Type: `"string"`. Optional.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

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
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](resources--alert_policy--reference--group-001.md#canonical-a563258899566a413be77332945f0d55537101bda4dabe9784ba0fa1ddf1d89f): complete subsection reference.

<a id="canonical-969969b5ee7a1acbb3e05dcffb203e1482935e0070b2763738de4ce3979c0017"></a>

<a id="canonical-141f859ec6b20e6e0f391efbe7c6b70ff0e375d402b5de5eb37c1651a47c92bf"></a>

## repeat_interval property — notification_parameters / 82679e48d075 / 6

Type: `"string"`. Optional.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

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
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](resources--alert_policy--reference--group-001.md#canonical-9499d5fa3a0ac8e8b1da965eb0666b4c59fe2761583749157f1a65d79f39b350): complete subsection reference.

<a id="canonical-4974a7580204212cbe872939d6f4a4da21689e0f34accbbc93fb5421a3d79a9a"></a>

## Next pages — notification_parameters / 82679e48d075 / 7

- [notification_parameters.custom](resources--alert_policy--reference--group-001.md#canonical-d7fd006e619b7ccd1949ffb5ad480cb68f17b550fe0bbb9849890998fd4d94ca)
- [notification_parameters.default](resources--alert_policy--reference--group-001.md#canonical-9b631c58d71278f7a8922e12328635927f8dc1bbd6138bd6126b8223c3b5e10e)
- [notification_parameters.individual](resources--alert_policy--reference--group-001.md#canonical-a563258899566a413be77332945f0d55537101bda4dabe9784ba0fa1ddf1d89f)
- [notification_parameters.ves_io_group](resources--alert_policy--reference--group-001.md#canonical-9499d5fa3a0ac8e8b1da965eb0666b4c59fe2761583749157f1a65d79f39b350)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-d7fd006e619b7ccd1949ffb5ad480cb68f17b550fe0bbb9849890998fd4d94ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-251fccd19f029a04cc08fd2fbc049b74060338c3a152fdb943fddbf6fda51642"></a>

## notification_parameters.custom — notification_parameters.custom / 8cbe5ea830a5 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- notification_parameters.custom

<a id="canonical-e6d3b452c59669e10a32d4923be4c7a33168cc99a054dd92a758eac5077f56d2"></a>

Type: `"object"`. single nested block, Optional.

Specify list of custom labels to group/aggregate the alerts.

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
custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-ae4dcdd2f675e9b69de423ace2eec1da963fcb68a4798263d6d7918e88d42023"></a>

## Direct properties — notification_parameters.custom / 8cbe5ea830a5 / 3

<a id="canonical-e70d4a4b962243b2d3893288ce5689dc104c9db674b75611e9ee9292f0f4462a"></a>

<a id="canonical-cda04280016f06a53f94435c335730d1de5d1b65f15bb1895cee981ffc9c6120"></a>

## labels property — notification_parameters.custom / 8cbe5ea830a5 / 4

Type: `["list", "string"]`. Optional.

Name of labels to group/aggregate the alerts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d930be1b53d8b9d2212626ee2f7ba2dcadafb9ac778650d9c55d094601e2f9a3"></a>

## Next pages — notification_parameters.custom / 8cbe5ea830a5 / 5

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-9b631c58d71278f7a8922e12328635927f8dc1bbd6138bd6126b8223c3b5e10e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-515641e36ef30bf84cc7e10e997ff415689a5ccfc4016c4b71e7f20a8d7bdbac"></a>

## notification_parameters.default — notification_parameters.default / 96de9088b085 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- notification_parameters.default

<a id="canonical-fa3255dd7872528f1a7d68e8746a874d8e97fed47e59d7aeb5991e19a481edca"></a>

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
default = {}
```

<a id="canonical-bddf5571934d19d2a1c221fb0ed256ab56c73251c2e209108ce0dd40cf676cba"></a>

## Direct properties — notification_parameters.default / 96de9088b085 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ca566f66cf485a1264f722bdcbaaab0f620c6bc272843ce8a3c6d8b189dd2a7"></a>

## Next pages — notification_parameters.default / 96de9088b085 / 4

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-a563258899566a413be77332945f0d55537101bda4dabe9784ba0fa1ddf1d89f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8b19548076495d2025a203633e6e4bc0da827146cee31126620da7a0c2177a5"></a>

## notification_parameters.individual — notification_parameters.individual / ee0baf1544be / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- notification_parameters.individual

<a id="canonical-ce5152bf4be29ad0baba65249a63bb6d0a7b1cad0daff75584ef8cad4a2d4e23"></a>

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
individual = {}
```

<a id="canonical-b931f61174f2596177faf3ae85e89c49394fda569768adc9da9d715a0c48bd4a"></a>

## Direct properties — notification_parameters.individual / ee0baf1544be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e323e1683ac054c6b048ae03df95653e8f3ecc034b330202470873699b495725"></a>

## Next pages — notification_parameters.individual / ee0baf1544be / 4

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-9499d5fa3a0ac8e8b1da965eb0666b4c59fe2761583749157f1a65d79f39b350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61c211eaf470b8839c0f6bb1be3afd1205dc47d449f2ad4efd5626aa44c6f502"></a>

## notification_parameters.ves_io_group — notification_parameters.ves_io_group / 657e5c7f6c5c / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- notification_parameters.ves_io_group

<a id="canonical-c3b4b033d118dd3e3d77ee9a1fa3580f14efa5a9762d39bcad901dce5633a2ea"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ves io group.

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
ves_io_group = {}
```

<a id="canonical-21b707cecf079be4b21a667bb1431c98d973907580e099781da2c471da6330f8"></a>

## Direct properties — notification_parameters.ves_io_group / 657e5c7f6c5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0951f01e4b2b3a861ccaddc0b066294030362a3a0ffda070d12dc055fbc4195"></a>

## Next pages — notification_parameters.ves_io_group / 657e5c7f6c5c / 4

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-b2ae5a04dffeb6fb38d62f64ec8653df2807a130bb3e9f4ab95286c821906989)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-ddebc42c9d6a73646d69696ac8c1305a9c894c0c0109a249a4223cf69593e762"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6dc434bb9fd325ab9ecc3780c369033fe0b4015291a31e95cd9c9c88f4cc057"></a>

## receivers — receivers / 57ff05f23ca7 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- receivers

<a id="canonical-ca502c6677ddf0f1487cf190dfe2cbf56d2f523a03cc6e5b2a5310d6af92eeee"></a>

Type: `"object"`. list nested block, Optional.

List of Alert Receivers where the alerts will be sent.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
receivers {
  # Configure direct properties listed below.
}
```

<a id="canonical-5ce8e248d4695ffcb37a0e8ce89f898e47132b50b1c0d6ce80a85d03cac89a61"></a>

## Direct properties — receivers / 57ff05f23ca7 / 3

<a id="canonical-456e4fa7ae4c5b202063b051474c10b39a1c885f763c68ff8a85b94d74465a14"></a>

<a id="canonical-cc5bfb396dd0a25df329b901d2ea7ee588378cfacf82ad6a8ff35ce72db880e0"></a>

## kind property — receivers / 57ff05f23ca7 / 4

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

<a id="canonical-7e3114dd636151bcd89d6924f31af3c92f07af784231d8006378d5088f4394f9"></a>

<a id="canonical-0054c056cec568613deb8c93cefd0aa8b337a6dec0269ef0edb4b9d1cd5423ea"></a>

## name property — receivers / 57ff05f23ca7 / 5

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

<a id="canonical-63f4df31ba49dae6ba8c186cac4bf3cb2e998a5b4db2299734adac0ca025b22d"></a>

<a id="canonical-85d245c8d6d96160367bf8660619459fb6018d5006de3333a132d9991dc53d73"></a>

## namespace property — receivers / 57ff05f23ca7 / 6

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

<a id="canonical-49f872180287e28ca1d330864dc12503c4321a27d498ba36a451b6907b0bb69e"></a>

<a id="canonical-ae87565b9db1ef4682fbe9d634104f949fef05c68a25924343e61e85cca4fb4c"></a>

## tenant property — receivers / 57ff05f23ca7 / 7

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

<a id="canonical-73e44f999ea49c85ef0772ca21bfa5cae69f7de21ce109b6a51033b7012ff724"></a>

<a id="canonical-ec2ffb3a646a2e989fff1417bc0865c3d240c565509bd8268360c7a57903ee0f"></a>

## uid property — receivers / 57ff05f23ca7 / 8

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

<a id="canonical-94d659174fe429a13bae4cc128c79f831261e5579d137e765b70e5ccbd888a71"></a>

## Next pages — receivers / 57ff05f23ca7 / 9

- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5188cd897f8ffc21a75ec2f852a3caac44ebf4bc5de145d1f707549efc38dcb2"></a>

## routes — routes / 0cf9d943e05e / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- routes

<a id="canonical-31e6b78e36589cd87c0dc813534cfc61f453140ffe8aa7f3b20a55e545010fa0"></a>

Type: `"object"`. list nested block, Optional.

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Upstream description:

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("alertname",
    "alertname_regex"),
  validators.ConflictingListObjectAttributes("alertname",
    "any"),
  validators.ConflictingListObjectAttributes("alertname",
    "custom"),
  validators.ConflictingListObjectAttributes("alertname",
    "group"),
  validators.ConflictingListObjectAttributes("alertname",
    "severity"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "any"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "custom"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "group"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "severity"),
  validators.ConflictingListObjectAttributes("any",
    "custom"),
  validators.ConflictingListObjectAttributes("any",
    "group"),
  validators.ConflictingListObjectAttributes("any",
    "severity"),
  validators.ConflictingListObjectAttributes("custom",
    "group"),
  validators.ConflictingListObjectAttributes("custom",
    "severity"),
  validators.ConflictingListObjectAttributes("dont_send",
    "send"),
  validators.ConflictingListObjectAttributes("group",
    "severity")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-9280ee8af93ca97db267834a6b59fd1b56ef8eb113d44a440b524a86a528b6a6"></a>

## Direct properties — routes / 0cf9d943e05e / 3

<a id="canonical-38356d65ce2bea9fd65ab473a35e7adae52b18bf5e68157c7c7483567ab2eeb2"></a>

<a id="canonical-e8940bc66a861b6db36575a642123a60ce759217a68e6ebdd0b9458af9cfef14"></a>

## alertname property — routes / 0cf9d943e05e / 4

Type: `"string"`. Optional.

\[Enum:
SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN|SITE\_PHYSICAL\_INTERFACE\_DOWN|TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN|SERVICE\_SERVER\_ERROR|SERVICE\_CLIENT\_ERROR|SERVICE\_HEALTH\_LOW|SERVICE\_UNAVAILABLE|SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE|SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL|MALICIOUS\_USER\_DETECTED|WAF\_TOO\_MANY\_ATTACKS|API\_SECURITY\_TOO\_MANY\_ATTACKS|SERVICE\_POLICY\_TOO\_MANY\_ATTACKS|WAF\_TOO\_MANY\_MALICIOUS\_BOTS|BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS|THREAT\_CAMPAIGN|VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN|VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING|TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON|TLS\_CUSTOM\_CERTIFICATE\_EXPIRED|L7DDOS|DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD|API\_SECURITY\_UNUSED\_API\_DETECTED|API\_SECURITY\_SHADOW\_API\_DETECTED|API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED|API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED|ROUTED\_DDOS\_ALERT\_NOTIFICATION|ROUTED\_DDOS\_MITIGATION\_NOTIFICATION|ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION|L7\_DDOS\_AUTO\_MITIGATION\]
List of Alert Names Customer tunnel interface down Physical Interface down Tunnel Interfaces to
Customer Site Down Virtual Host server error Virtual Host client error Service Health Low Service
Unavailable Virtual Host server error Virtual Host client error Endpoint Healthcheck failure
Synthetic.. Possible values are \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`,
\`SITE\_PHYSICAL\_INTERFACE\_DOWN\`, \`TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN\`,
\`SERVICE\_SERVER\_ERROR\`, \`SERVICE\_CLIENT\_ERROR\`, \`SERVICE\_HEALTH\_LOW\`,
\`SERVICE\_UNAVAILABLE\`, \`SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE\`,
\`SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE\`, \`SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE\`,
\`SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL\`, \`MALICIOUS\_USER\_DETECTED\`,
\`WAF\_TOO\_MANY\_ATTACKS\`, \`API\_SECURITY\_TOO\_MANY\_ATTACKS\`,
\`SERVICE\_POLICY\_TOO\_MANY\_ATTACKS\`, \`WAF\_TOO\_MANY\_MALICIOUS\_BOTS\`,
\`BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS\`, \`THREAT\_CAMPAIGN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING\`, \`TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\`, \`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRED\`, \`L7DDOS\`, \`DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD\`,
\`API\_SECURITY\_UNUSED\_API\_DETECTED\`, \`API\_SECURITY\_SHADOW\_API\_DETECTED\`,
\`API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED\`,
\`API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED\`, \`ROUTED\_DDOS\_ALERT\_NOTIFICATION\`,
\`ROUTED\_DDOS\_MITIGATION\_NOTIFICATION\`, \`ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION\`,
\`L7\_DDOS\_AUTO\_MITIGATION\`. Defaults to \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`.

Upstream description:

List of Alert Names

Customer tunnel interface down Physical Interface down Tunnel Interfaces to Customer Site Down
Virtual Host server error Virtual Host client error Service Health Low Service Unavailable Virtual
Host server error Virtual Host client error Endpoint Healthcheck failure Synthetic monitor health
critical Malicious user detected Virtual Host WAF security events detected Virtual Host API security
events detected Virtual Host Service Policy security events detected Virtual Host Many Malicious
Bots based WAF security events detected Virtual Host Many Malicious Bots based Bot Defense security
events detected Virtual Host Many Threat campaign based WAF security events detected Suspicious
domain identified by Client-Side Defense service Client-Side Defense has identified a suspicious
script that is reading sensitive form field TLS Automatic Certificate renewal is failing TLS
Automatic Certificate renewal is still failing after multiple retries TLS Automatic Certificate has
expired TLS Custom Certificate will expire in less than 28 days TLS Custom Certificate will expire
in less than 15 days TLS Custom Certificate has expired DDoS security event detected DNS Zone
Ignored a Duplicate Record Create Request Unused APIs Detected Shadow APIs Detected Endpoints With
Sensitive Data In Response Detected High Risk Score Endpoints Detected A routed DDoS traffic anomaly
has been detected A routed DDoS mitigation has been implemented to block malicious traffic A routed
DDoS tunnel status has been changed L7 DDoS attack was detected, automatic mitigation is taking
place.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
    "SITE_PHYSICAL_INTERFACE_DOWN",
    "TUNNELS_TO_CUSTOMER_SITE_DOWN",
    "SERVICE_SERVER_ERROR",
    "SERVICE_CLIENT_ERROR",
    "SERVICE_HEALTH_LOW",
    "SERVICE_UNAVAILABLE",
    "SERVICE_SERVER_ERROR_PER_SOURCE_SITE",
    "SERVICE_CLIENT_ERROR_PER_SOURCE_SITE",
    "SERVICE_ENDPOINT_HEALTHCHECK_FAILURE",
    "SYNTHETIC_MONITOR_HEALTH_CRITICAL",
    "MALICIOUS_USER_DETECTED",
    "WAF_TOO_MANY_ATTACKS",
    "API_SECURITY_TOO_MANY_ATTACKS",
    "SERVICE_POLICY_TOO_MANY_ATTACKS",
    "WAF_TOO_MANY_MALICIOUS_BOTS",
    "BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS",
    "THREAT_CAMPAIGN",
    "VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN",
    "VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING",
    "TLS_AUTOMATIC_CERTIFICATE_EXPIRED",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON",
    "TLS_CUSTOM_CERTIFICATE_EXPIRED",
    "L7DDOS",
    "DNS_ZONE_IGNORED_DUPLICATE_RECORD",
    "API_SECURITY_UNUSED_API_DETECTED",
    "API_SECURITY_SHADOW_API_DETECTED",
    "API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED",
    "API_SECURITY_RISK_SCORE_HIGH_DETECTED",
    "ROUTED_DDOS_ALERT_NOTIFICATION",
    "ROUTED_DDOS_MITIGATION_NOTIFICATION",
    "ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION",
    "L7_DDOS_AUTO_MITIGATION"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
  "enum": [
    "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
    "SITE_PHYSICAL_INTERFACE_DOWN",
    "TUNNELS_TO_CUSTOMER_SITE_DOWN",
    "SERVICE_SERVER_ERROR",
    "SERVICE_CLIENT_ERROR",
    "SERVICE_HEALTH_LOW",
    "SERVICE_UNAVAILABLE",
    "SERVICE_SERVER_ERROR_PER_SOURCE_SITE",
    "SERVICE_CLIENT_ERROR_PER_SOURCE_SITE",
    "SERVICE_ENDPOINT_HEALTHCHECK_FAILURE",
    "SYNTHETIC_MONITOR_HEALTH_CRITICAL",
    "MALICIOUS_USER_DETECTED",
    "WAF_TOO_MANY_ATTACKS",
    "API_SECURITY_TOO_MANY_ATTACKS",
    "SERVICE_POLICY_TOO_MANY_ATTACKS",
    "WAF_TOO_MANY_MALICIOUS_BOTS",
    "BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS",
    "THREAT_CAMPAIGN",
    "VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN",
    "VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING",
    "TLS_AUTOMATIC_CERTIFICATE_EXPIRED",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON",
    "TLS_CUSTOM_CERTIFICATE_EXPIRED",
    "L7DDOS",
    "DNS_ZONE_IGNORED_DUPLICATE_RECORD",
    "API_SECURITY_UNUSED_API_DETECTED",
    "API_SECURITY_SHADOW_API_DETECTED",
    "API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED",
    "API_SECURITY_RISK_SCORE_HIGH_DETECTED",
    "ROUTED_DDOS_ALERT_NOTIFICATION",
    "ROUTED_DDOS_MITIGATION_NOTIFICATION",
    "ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION",
    "L7_DDOS_AUTO_MITIGATION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5240f736df355c74ec9da635ef6e30153d1014e2885eec5edd1dadd806dc5133"></a>

<a id="canonical-7d3df1933a0e4a674e0a15ef33fa6a152000f10b9c0739e7aa140dc562cfa4a1"></a>

## alertname_regex property — routes / 0cf9d943e05e / 5

Type: `"string"`. Optional.

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

Upstream description:

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

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

- [any](resources--alert_policy--reference--group-001.md#canonical-5ee1317a7d2e9e58f0b1958b9dc08fff2c2e3091de257ed797f380badfa3ac7f): complete subsection reference.

- [custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1): complete subsection reference.

- [dont_send](resources--alert_policy--reference--group-001.md#canonical-8f79c61c87f14aa5b6acef1219d7dcf7809637f5f59e996635196a6f7d924b86): complete subsection reference.

- [group](resources--alert_policy--reference--group-001.md#canonical-d91d19911d5528ef3d077cdaf231f4dd9fd2e8cf3db92274f00c957faaf1e0a8): complete subsection reference.

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d): complete subsection reference.

- [send](resources--alert_policy--reference--group-001.md#canonical-94d3715956f56bb9bed5c59211bfccea8eebbaf1d0e785962cb8aeaf622e4e2b): complete subsection reference.

- [severity](resources--alert_policy--reference--group-001.md#canonical-8f06144aa977deba57a9f5fce5accba2cf72b47adf0217e3c0770a3c4bf0d077): complete subsection reference.

<a id="canonical-fee3613a64e35ad2176e3d5d4424eff6b8efa8fa9528a1f8ef61232a89c3421a"></a>

## Next pages — routes / 0cf9d943e05e / 6

- [routes.any](resources--alert_policy--reference--group-001.md#canonical-5ee1317a7d2e9e58f0b1958b9dc08fff2c2e3091de257ed797f380badfa3ac7f)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- [routes.dont_send](resources--alert_policy--reference--group-001.md#canonical-8f79c61c87f14aa5b6acef1219d7dcf7809637f5f59e996635196a6f7d924b86)
- [routes.group](resources--alert_policy--reference--group-001.md#canonical-d91d19911d5528ef3d077cdaf231f4dd9fd2e8cf3db92274f00c957faaf1e0a8)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- [routes.send](resources--alert_policy--reference--group-001.md#canonical-94d3715956f56bb9bed5c59211bfccea8eebbaf1d0e785962cb8aeaf622e4e2b)
- [routes.severity](resources--alert_policy--reference--group-001.md#canonical-8f06144aa977deba57a9f5fce5accba2cf72b47adf0217e3c0770a3c4bf0d077)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-5ee1317a7d2e9e58f0b1958b9dc08fff2c2e3091de257ed797f380badfa3ac7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4936e07edf4b89f8837a8ad55d65b7afa3abd9ebe718e4c13c8e87e9b8020da5"></a>

## routes.any — routes.any / 1905b6982221 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- routes.any

<a id="canonical-82078830fd6e83172223bef3e40c7b0bea09d956fef9a5b821265811c48614f9"></a>

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
any = {}
```

<a id="canonical-7f997b2252ccfa34ad52eacfd0528ae4ee7efe12e4431f48b23b7de3e740f916"></a>

## Direct properties — routes.any / 1905b6982221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4631ec34bc51be3b4887eaef9a3542d8ef195d8a2d65c8aade13b415dbd33b4"></a>

## Next pages — routes.any / 1905b6982221 / 4

- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dec4bac3a8774c5c616a11b81f574d6bcdfe6ce521763521ff5ef9045c05fef6"></a>

## routes.custom — routes.custom / bf52cdc354d2 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- routes.custom

<a id="canonical-19c17b0ef4872eb5199067f166a1ef5ae22093f0b11602f7ecc3b65c8e650d03"></a>

Type: `"object"`. single nested block, Optional.

Set of matchers an alert has to fulfill to match the route.

Upstream description:

A set of matchers an alert has to fulfill to match the route.

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
custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-bbe925e11b52debc905ee8db1c6b6278a62223211fd30f45f3f84f5ae527ed7c"></a>

## Direct properties — routes.custom / bf52cdc354d2 / 3

- [alertlabel](resources--alert_policy--reference--group-001.md#canonical-e786a7a5399a4653d5e0303babf4bbfb610c21a1932578def652735b2395ed51): complete subsection reference.

- [alertname](resources--alert_policy--reference--group-001.md#canonical-881bed8a6c5b6576d5bd4daabc0b617561e33a6579b59e1a713dc3c8ef4095f3): complete subsection reference.

- [group](resources--alert_policy--reference--group-001.md#canonical-85117263bf02446cced615334797aeea57a30d9621b87e0fc6d294a50d0630dc): complete subsection reference.

- [severity](resources--alert_policy--reference--group-001.md#canonical-c4fd94b3517690011de9b2df1cf5f8738992f3dee09e19f602d8e0f36e5fb9c3): complete subsection reference.

<a id="canonical-3835f27d0aca0e43f62df40126c7f09764731e438c7f3feaaef10d70fe4e1066"></a>

## Next pages — routes.custom / bf52cdc354d2 / 4

- [routes.custom.alertlabel](resources--alert_policy--reference--group-001.md#canonical-e786a7a5399a4653d5e0303babf4bbfb610c21a1932578def652735b2395ed51)
- [routes.custom.alertname](resources--alert_policy--reference--group-001.md#canonical-881bed8a6c5b6576d5bd4daabc0b617561e33a6579b59e1a713dc3c8ef4095f3)
- [routes.custom.group](resources--alert_policy--reference--group-001.md#canonical-85117263bf02446cced615334797aeea57a30d9621b87e0fc6d294a50d0630dc)
- [routes.custom.severity](resources--alert_policy--reference--group-001.md#canonical-c4fd94b3517690011de9b2df1cf5f8738992f3dee09e19f602d8e0f36e5fb9c3)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-e786a7a5399a4653d5e0303babf4bbfb610c21a1932578def652735b2395ed51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6955d9ce18a0dc5092384ad7e4f7149b7042852a9170a190061e6f412b42b92"></a>

## routes.custom.alertlabel — routes.custom.alertlabel / f15c61bfad3f / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- routes.custom.alertlabel

<a id="canonical-cd2f000b3535422b0ff9cfff95fe46ccf725bfe8bc16b5abd7e5d78e5052d583"></a>

Type: `"object"`. single nested block, Optional.

AlertLabel to configure the alert policy rule.

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
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  }
}
```

Terraform syntax:

```terraform
alertlabel {}
```

<a id="canonical-6d101076d19e5df3faeed9d36d750f4b52aa6f2a43e72fd951c4cd73b73a0a0a"></a>

## Direct properties — routes.custom.alertlabel / f15c61bfad3f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cf80fab33c5cbaaafe1004cd8323b15815d5ba57a93bf4ff44e06d8dbc87113"></a>

## Next pages — routes.custom.alertlabel / f15c61bfad3f / 4

- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-881bed8a6c5b6576d5bd4daabc0b617561e33a6579b59e1a713dc3c8ef4095f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c664be3e4f00e0c1ce54a523e5481087e08813bfcac07492a0b2de09d10dee04"></a>

## routes.custom.alertname — routes.custom.alertname / 1c6a396d5f13 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- routes.custom.alertname

<a id="canonical-bf457c276336c727b69c82c8abe36eb900ebb1b4842555ddb592cc6e6f372266"></a>

Type: `"object"`. single nested block, Optional.

Label Matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_match",
    "regex_match")}
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
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

Terraform syntax:

```terraform
alertname {
  # Configure direct properties listed below.
}
```

<a id="canonical-50694179aab1ce823f59efb209be2fb31949410edbff952a069294869add6e44"></a>

## Direct properties — routes.custom.alertname / 1c6a396d5f13 / 3

<a id="canonical-8f58b18ec20bbb4513718c71e707b55712b2f1af8329a46cb9a6add2c5c2d260"></a>

<a id="canonical-8da585eb7766a6a2d769406da5daba672eb7cfb78c9f66ffb9c89822362d1c91"></a>

## exact_match property — routes.custom.alertname / 1c6a396d5f13 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-50ae7d040a9bdd9103e0339615baf15b539f91c9d40157dbaf0d02ea3a92ae4c"></a>

<a id="canonical-fcdd22daaac8b84de3c3233167f0dce19efaf2262001d908acb7ac62f4787e0e"></a>

## regex_match property — routes.custom.alertname / 1c6a396d5f13 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-71465715467492d8fb8e8a831fe99988e3b83755b9395e3954bdf45a2e2ec482"></a>

## Next pages — routes.custom.alertname / 1c6a396d5f13 / 6

- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-85117263bf02446cced615334797aeea57a30d9621b87e0fc6d294a50d0630dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7a3014d6c04f7bc5ad7ad2d8b53ab815e2d5b374412835c754a054ceb8d4513"></a>

## routes.custom.group — routes.custom.group / 600326d12724 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- routes.custom.group

<a id="canonical-756da3439d317c3b8a77df6772bf2c00765dd50338caa6a3ecbf55fb38a6cb48"></a>

Type: `"object"`. single nested block, Optional.

Label Matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_match",
    "regex_match")}
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
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

Terraform syntax:

```terraform
group {
  # Configure direct properties listed below.
}
```

<a id="canonical-41ced4731abedd6b5143fe6308859e15423b997084bb1716f986aae9075cfd90"></a>

## Direct properties — routes.custom.group / 600326d12724 / 3

<a id="canonical-ec75a716e6368f4a6fbdacf3e374daaf9696e84d631ff05796c52070b05e88fe"></a>

<a id="canonical-1d9c26ca7cce49698058750a5a90669bb0bc943d5c5c1ecf8c3b06afae658637"></a>

## exact_match property — routes.custom.group / 600326d12724 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-53d1e621967dfa33d772c1940ef1d1bfdc12a4864fa33fadffedeef4d6621752"></a>

<a id="canonical-67788853968b351a745e1015904a6682cbc41331b96895a84b236e22dd4ae60e"></a>

## regex_match property — routes.custom.group / 600326d12724 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-85370615b517f6e72c86c3d57fee1bf195e37446fcf53f819552c608ee83c9c2"></a>

## Next pages — routes.custom.group / 600326d12724 / 6

- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-c4fd94b3517690011de9b2df1cf5f8738992f3dee09e19f602d8e0f36e5fb9c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b973166ecb0e92882d3d36e396f694b4a390635a01c239f5591ac2ca416ce2d"></a>

## routes.custom.severity — routes.custom.severity / afd3bc8be971 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- routes.custom.severity

<a id="canonical-cabea6ebae22d75eed65c56135620838b071ded16869f7ecbe6deaea494e325a"></a>

Type: `"object"`. single nested block, Optional.

Label Matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_match",
    "regex_match")}
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
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

Terraform syntax:

```terraform
severity {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d6591fe896b8e18b5aa6ef5e427f575a2bcf2ea20ed8a58aeeeabae52eee534"></a>

## Direct properties — routes.custom.severity / afd3bc8be971 / 3

<a id="canonical-907689b35f84d598c131e17974dcdce516f14550fbfe6b034aadbdfe9cb9a392"></a>

<a id="canonical-caeb8cabbe87cf16195fb16230f145e81a065b7d7a9c195ee6d55ce163f8afa1"></a>

## exact_match property — routes.custom.severity / afd3bc8be971 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-ec79e551c698861b047422ec5e1d8cc655e9b95167753a8c9fe87f547dac0c75"></a>

<a id="canonical-2f58b5269fc3dae8c32bbaeeb28d269a8cb40648b0834b5ca8cea65fe7745262"></a>

## regex_match property — routes.custom.severity / afd3bc8be971 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-94aa2c1adb48f7a805b67b4864a4ac3a40cb1aa8e29fb4492a9e92ce23a3bb7c"></a>

## Next pages — routes.custom.severity / afd3bc8be971 / 6

- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-ec860c14c496a413e1b2d48d671539de347ab44204bdbfa48c2629a2cd2336d1)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-8f79c61c87f14aa5b6acef1219d7dcf7809637f5f59e996635196a6f7d924b86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb73dab8449e1a1c16a7e3637dd0d156d45ee7bf0f7c3ff4bef01c75558f7c0d"></a>

## routes.dont_send — routes.dont_send / 1d4e0a3b4fb8 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- routes.dont_send

<a id="canonical-0703143905a8bae78a865ac9e98b6bb0b56556ad954f91c6ea6c1197f3435936"></a>

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
dont_send = {}
```

<a id="canonical-0c08ab6e1697d8a55e6d98730795f313fd18996b8ce16b5e1136787327361bd7"></a>

## Direct properties — routes.dont_send / 1d4e0a3b4fb8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be643e53dd2439bbc937f3058e9e0f9109296db00704232f31d2529a74027a06"></a>

## Next pages — routes.dont_send / 1d4e0a3b4fb8 / 4

- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-d91d19911d5528ef3d077cdaf231f4dd9fd2e8cf3db92274f00c957faaf1e0a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c69a7f9ab7828017343fdfa8aa5f18a1bb5a6fa041157024a3963c3b5183f68"></a>

## routes.group — routes.group / 2042314a2612 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- routes.group

<a id="canonical-02f8e90efed32ad005f88a0b9df8db4dea1451d789175cbfcc019a1664c92138"></a>

Type: `"object"`. single nested block, Optional.

Select one or more known group names to match the incoming alert.

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
group {
  # Configure direct properties listed below.
}
```

<a id="canonical-a39d04d360f7da7b6bc0a3f32cc58732c50988eb55f90717443c7ff1cfecaac2"></a>

## Direct properties — routes.group / 2042314a2612 / 3

<a id="canonical-545ba5a1aa3886d784e76fd78dd468342efbaaa0185c4c20d4ed5e1ac8720ab5"></a>

<a id="canonical-11dbc8e0dbc87c2c3d41494fe1262d4394f194c2be0c63c2f7a520d954e5b376"></a>

## groups property — routes.group / 2042314a2612 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
INFRASTRUCTURE|IAAS\_CAAS|VIRTUAL\_HOST|VOLT\_SHARE|UAM|SECURITY|TIMESERIES\_ANOMALY|SHAPE\_SECURITY|SECURITY\_CSD|CDN|SYNTHETIC\_MONITORS|TLS|SECURITY\_BOT\_DEFENSE|CLOUD\_LINK|DNS|ROUTED\_DDOS\]
Groups. Name of groups to match the alert. Possible values are \`INFRASTRUCTURE\`, \`IAAS\_CAAS\`,
\`VIRTUAL\_HOST\`, \`VOLT\_SHARE\`, \`UAM\`, \`SECURITY\`, \`TIMESERIES\_ANOMALY\`,
\`SHAPE\_SECURITY\`, \`SECURITY\_CSD\`, \`CDN\`, \`SYNTHETIC\_MONITORS\`, \`TLS\`,
\`SECURITY\_BOT\_DEFENSE\`, \`CLOUD\_LINK\`, \`DNS\`, \`ROUTED\_DDOS\`. Defaults to
\`INFRASTRUCTURE\`.

Upstream description:

Name of groups to match the alert.

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

<a id="canonical-f84a2b816efd4ee8b06c95cdedae984a98da7ed7eb5140152ef15e0b4b70edf2"></a>

## Next pages — routes.group / 2042314a2612 / 5

- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b9b2aae1e0a8e6ddd975e9c1df065156f09c65f0cc651b8a8b413928b39c6f9"></a>

## routes.notification_parameters — routes.notification_parameters / e2cfb0a18f4d / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- routes.notification_parameters

<a id="canonical-ff0b26198dab0e39efba9ac00ad38e30139b438c2333239e71f9c77df467d513"></a>

Type: `"object"`. single nested block, Optional.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom",
    "default"),
  validators.ConflictingObjectAttributes("custom",
    "individual"),
  validators.ConflictingObjectAttributes("custom",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("default",
    "individual"),
  validators.ConflictingObjectAttributes("default",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("individual",
    "ves_io_group")}
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
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

Terraform syntax:

```terraform
notification_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-c95168521c4a04a9b915c3c8e292667a90330d4a6293da86a34845c06dabff7f"></a>

## Direct properties — routes.notification_parameters / e2cfb0a18f4d / 3

- [custom](resources--alert_policy--reference--group-001.md#canonical-e7ec174ac546f524852ca0c6c0655badebe668e8e549198520aa1ac8f02175c5): complete subsection reference.

- [default](resources--alert_policy--reference--group-001.md#canonical-705bcec178acf4b024d1068f167c94fc8de1ddfcc0a2ccf886e1e6c61b3f8272): complete subsection reference.

<a id="canonical-c04ce99080b20ae5caf199760d83eef1bfd11f02b26a2dd5f85312eafc525bdb"></a>

<a id="canonical-db796cc4aa9ceda4c0d3b458a1f69cbb60252352c6cf4b58059e74f97757a29d"></a>

## group_interval property — routes.notification_parameters / e2cfb0a18f4d / 4

Type: `"string"`. Optional.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-47bf5c10d7a6c5bb91984672e31e8a62b62bf6a73eb5ee81c1242d8dae069611"></a>

<a id="canonical-80942f9c05336dfcc003c1d595521a81179dca095fd29cf49d2ec7f96567999e"></a>

## group_wait property — routes.notification_parameters / e2cfb0a18f4d / 5

Type: `"string"`. Optional.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

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
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](resources--alert_policy--reference--group-001.md#canonical-b7d43e0cd76a0d736556fa6da75baab97fcaa30847b3fb5947a99f8bb4f2bcab): complete subsection reference.

<a id="canonical-c78f96396893376273e675e21ec4df97d94b3873d659c1263438ddc859ed3966"></a>

<a id="canonical-9f3ad3bf11326ac2b2c89427db5b6871f2c69827471027f4923be12c307e0122"></a>

## repeat_interval property — routes.notification_parameters / e2cfb0a18f4d / 6

Type: `"string"`. Optional.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

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
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](resources--alert_policy--reference--group-001.md#canonical-86f711fa2644264d3184e1b9e4029d7073c7bc810947912d1b1eb760c79c3e0b): complete subsection reference.

<a id="canonical-b7cfa55fdf7514c8f23cebd7fdd336f75275f9fea7ca34cfcfc127285849309d"></a>

## Next pages — routes.notification_parameters / e2cfb0a18f4d / 7

- [routes.notification_parameters.custom](resources--alert_policy--reference--group-001.md#canonical-e7ec174ac546f524852ca0c6c0655badebe668e8e549198520aa1ac8f02175c5)
- [routes.notification_parameters.default](resources--alert_policy--reference--group-001.md#canonical-705bcec178acf4b024d1068f167c94fc8de1ddfcc0a2ccf886e1e6c61b3f8272)
- [routes.notification_parameters.individual](resources--alert_policy--reference--group-001.md#canonical-b7d43e0cd76a0d736556fa6da75baab97fcaa30847b3fb5947a99f8bb4f2bcab)
- [routes.notification_parameters.ves_io_group](resources--alert_policy--reference--group-001.md#canonical-86f711fa2644264d3184e1b9e4029d7073c7bc810947912d1b1eb760c79c3e0b)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-e7ec174ac546f524852ca0c6c0655badebe668e8e549198520aa1ac8f02175c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-541640717a7f174cabcb7195485313f0db3c1ea632c889f332b941d6ee71942a"></a>

## routes.notification_parameters.custom — routes.notification_parameters.custom / 1c95c22bff3d / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- routes.notification_parameters.custom

<a id="canonical-2f627e920348b8527f88fb691d96c091e21ece6e50a2070b37637f3b858dc5f2"></a>

Type: `"object"`. single nested block, Optional.

Specify list of custom labels to group/aggregate the alerts.

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
custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-92649cc44062821954e7bc0a4b2ee4ae39b303dfc9daee38c7e8e209379123aa"></a>

## Direct properties — routes.notification_parameters.custom / 1c95c22bff3d / 3

<a id="canonical-68671e96323c0e5e50e61d97982b5dd45899b1800ae3465557c08c5ab9ea87b7"></a>

<a id="canonical-1b301e395a7c82b66cccc4ce96d61608d4f3840a92c42c1463387376d155e69c"></a>

## labels property — routes.notification_parameters.custom / 1c95c22bff3d / 4

Type: `["list", "string"]`. Optional.

Name of labels to group/aggregate the alerts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ff6aca7a6205649c05ea801a3d8395b32d28ee8abb2a162a24d9474167b59dff"></a>

## Next pages — routes.notification_parameters.custom / 1c95c22bff3d / 5

- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-705bcec178acf4b024d1068f167c94fc8de1ddfcc0a2ccf886e1e6c61b3f8272"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ea2d91842ce399871e1db069cce51b958a3301356034c8f7f74d27f22d3565c"></a>

## routes.notification_parameters.default — routes.notification_parameters.default / 3186c57ff12e / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- routes.notification_parameters.default

<a id="canonical-460ba4aa49d100ba139863b08fa5b2013a439832a8b973ec35e958cab9815be2"></a>

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
default = {}
```

<a id="canonical-66e92335540374b7e4a244cb35b719e7b9c7eacd4414dc15c5ec83615d669941"></a>

## Direct properties — routes.notification_parameters.default / 3186c57ff12e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-795cf7780770a83149c5baaca12e772735cc9db0d34fa9f130c28d736bdc2347"></a>

## Next pages — routes.notification_parameters.default / 3186c57ff12e / 4

- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-b7d43e0cd76a0d736556fa6da75baab97fcaa30847b3fb5947a99f8bb4f2bcab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc378d9f82aecfe61ed24e165baeabb4e7eaa80b0553b3ef6685b692941fa3ff"></a>

## routes.notification_parameters.individual — routes.notification_parameters.individual / 7225495c9a13 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- routes.notification_parameters.individual

<a id="canonical-7d072597343cda700082ebb1514bd09afaeaab273a9d4d490706a299380546d4"></a>

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
individual = {}
```

<a id="canonical-47cc56ffc6b01e6b54179b43022e46e9d55a9bd968565db4a9f2e0272aaed08d"></a>

## Direct properties — routes.notification_parameters.individual / 7225495c9a13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e8d551e3a4258b38cedadbee50dcd1be3e17929b89fb6e4eb5622b655b13e54"></a>

## Next pages — routes.notification_parameters.individual / 7225495c9a13 / 4

- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-86f711fa2644264d3184e1b9e4029d7073c7bc810947912d1b1eb760c79c3e0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26ace04ee0a2e7fffbfc9265e6cfa392df9979fc9e7b7c6cbd04d7f45079cba1"></a>

## routes.notification_parameters.ves_io_group — routes.notification_parameters.ves_io_group / 4668768ea3a7 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- routes.notification_parameters.ves_io_group

<a id="canonical-b73c816c163e9433e7c99fe7e7ddba9184217b40c118ba7c319e9ff791f8e330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ves io group.

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
ves_io_group = {}
```

<a id="canonical-f087c22fff1147185be7cfbb48387da16f740215ac41d419c521176f2d14a3e7"></a>

## Direct properties — routes.notification_parameters.ves_io_group / 4668768ea3a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10121d94f5e001e2ad57d2bf9670ccee3afe0538a7335d0d07c7c88b2e567482"></a>

## Next pages — routes.notification_parameters.ves_io_group / 4668768ea3a7 / 4

- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-5e4aa018762e4989a64cd883b311f73d8171584535e3eb38ac848cd3a8fb5e7d)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-94d3715956f56bb9bed5c59211bfccea8eebbaf1d0e785962cb8aeaf622e4e2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-516996b935a7c531873019b6c0506da67a9268be1d1bbfb2e842d005510a6a65"></a>

## routes.send — routes.send / bc19d4015cb9 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- routes.send

<a id="canonical-2bdc08ee42ec16d8ca260de40ba127ef8decc4d83f997d0c45cc2a2a0218af48"></a>

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
send = {}
```

<a id="canonical-6356288d7c0d6b8c179d4c8c0c05737fd51eb7646c76dbd9b38930623b043e55"></a>

## Direct properties — routes.send / bc19d4015cb9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c9c63dcf39940cc819fb91d6f2f85d083ca834860f89d933ca4fb38addccd5e"></a>

## Next pages — routes.send / bc19d4015cb9 / 4

- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-8f06144aa977deba57a9f5fce5accba2cf72b47adf0217e3c0770a3c4bf0d077"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76a10ed171b404684e2902b50ba37b03f00ddc6eb65e88c71a773b289654a291"></a>

## routes.severity — routes.severity / 69466fe776ad / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- routes.severity

<a id="canonical-305b9fbc9efa5fc86135a53aa5528031af976b1b387dc2dbcb69bb520dde5c31"></a>

Type: `"object"`. single nested block, Optional.

Select one or more severity levels to match the incoming alert.

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
severity {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d044afa39780b7c9023db86f00212375253e816097b381f7e4c7b6021ed97d1"></a>

## Direct properties — routes.severity / 69466fe776ad / 3

<a id="canonical-cf0993fd803a4db468277015356010b6a6a2fca4c4de25327148af42d670cfc9"></a>

<a id="canonical-24b25ab27e51669e7c1c589a5b6939a782a677735624b18b03023649384a8f81"></a>

## severities property — routes.severity / 69466fe776ad / 4

Type: `["list", "string"]`. Optional.

\[Enum: MINOR|MAJOR|CRITICAL\] Severities. List of severity levels. Possible values are \`MINOR\`,
\`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of severity levels.

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

<a id="canonical-557a99b05c21869dc0d931b3014b925af4a511eeb8f9a0b27d0683fbf28bc752"></a>

## Next pages — routes.severity / 69466fe776ad / 5

- [routes](resources--alert_policy--reference--group-001.md#canonical-9ce1a579e5e57297b23e4d47b8d7720039964d4ab190e21fa0756be1c139080f)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-2ed7e2c972f5b95c7ffdc6ac774fd54ef2b45c05425cc0c15c7db6a7bcc7d7d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fc75edf73f13423914b3e7647a11ce9470e5b889a55c094db3bdd0fc18f6370"></a>

## timeouts — timeouts / e4218cd4171b / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- timeouts

<a id="canonical-d490587c4bcc0e9baf19af72c5609850e1bc405c2af382d498ddb8e2d35ac595"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a92e5b866869c3ef97b6f77bd6fbb564fd6f1d4cdc1bb949a8b5aa427719d327"></a>

## Direct properties — timeouts / e4218cd4171b / 3

<a id="canonical-f68ef5375777433fae28907ba5e8307aa2d5c67f9926481c1d7f05c5019fb907"></a>

<a id="canonical-d5fec164fec286dcbb8a1acf839f7cd75745228f0f1df541f5e421858b24a76b"></a>

## create property — timeouts / e4218cd4171b / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-33328e550fee5d12ecfcf158c1e09df2c1c8f95499bbe1e3f84d6df3d8e23620"></a>

<a id="canonical-67183640898019022a6e083e103993942acb6e3465632d8a99aab019e6604210"></a>

## delete property — timeouts / e4218cd4171b / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-62ecbf8efde71fd46015dd656fede8a255939ac0a2a723b26c83098ec211734c"></a>

<a id="canonical-e1aa9b8fa6f16bb181d1a75d9d664dc67b4249e3abf220979889e01a2ba6f097"></a>

## read property — timeouts / e4218cd4171b / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-c23df81a4da2430950004b118f5ba9488daed13bba0d0cd2876f854ada445147"></a>

<a id="canonical-097e6cd5015c2fb4626380fa5b2349d2008ba831ad90206269314174dc228df0"></a>

## update property — timeouts / e4218cd4171b / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7c55d3367bbd88f0e17709b576beb2572ef7614b1f145ab73c989f8e60d06650"></a>

## Next pages — timeouts / e4218cd4171b / 8

- [Property reference](resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
