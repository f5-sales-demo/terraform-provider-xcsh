---
page_title: "xcsh_app_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall reference."
---

# xcsh_app_firewall reference

<a id="canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b54c420d89fd87900faf43556901691458ff8a875ef3cdab9f6f9a48ab901adb"></a>

## Property reference — Property reference / 6d3207546445 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- Property reference

<a id="canonical-2fe50ae6e7ad51cb2d2f16c9005436ab944917e9034d33e8076289080a6e4d66"></a>

## Direct properties — Property reference / 6d3207546445 / 3

- [allow_all_response_codes](resources--app_firewall--reference--group-001.md#canonical-99e75007717b48e03533e9b898700c624862621444d75258e3499b4a886b6a53): complete subsection reference.

- [allowed_response_codes](resources--app_firewall--reference--group-001.md#canonical-18d5fde757750c5ae1f21da40bf6c3348e1ce30d1a5e9499aff566ee1c05c9f6): complete subsection reference.

<a id="canonical-b7a1643f43fcf910d6689f41da4417f85af63199400dee49cf718c7c4b60f455"></a>

<a id="canonical-92aff7f7d410dcf20c05d9d2cb63b8adfb965d5daa73bcbb0e5068152d732d11"></a>

## annotations property — Property reference / 6d3207546445 / 4

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

- [blocking](resources--app_firewall--reference--group-001.md#canonical-693b041ad0bdd1cbc9ed189a8d1f0c51264fa57d2eb99321044f6af82b046836): complete subsection reference.

- [blocking_page](resources--app_firewall--reference--group-001.md#canonical-072958fc9ffb537d907339d78644e0e909d2fafbeefef201fd6997bb27655858): complete subsection reference.

- [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-f561d7c60d03776bdafc8bcd1ff3cc2e20c3d873eba27b07550e9f5f35938d55): complete subsection reference.

- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543): complete subsection reference.

- [default_anonymization](resources--app_firewall--reference--group-001.md#canonical-6947b4b6ae9be6828bcabf9af8ec58df4d8d915bc460cd5ac7aca0837968974a): complete subsection reference.

- [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-44fe272375c06e26a85aa2b82dfb265fe7083495451fe26afb286e1f38d67f28): complete subsection reference.

- [default_detection_settings](resources--app_firewall--reference--group-001.md#canonical-5be5ae1e2e5176410bb44105286f6d85a52b42ab0b6197700dabfddfb431af0c): complete subsection reference.

<a id="canonical-3d21e40a84f93db709171c9798997e75a875a6d311e710043fd64bba0d7812f5"></a>

<a id="canonical-3ae06fb56aa853c508135fba65887f9c65d5263fde191bfb99b286e41afc27ad"></a>

## description property — Property reference / 6d3207546445 / 5

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

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25): complete subsection reference.

<a id="canonical-c58f2081977e182263854051ec1c1e744b42708f528b422eecda62e72e4d772e"></a>

<a id="canonical-786e9d019022cee289e72c232f46a51a95ce45ab48ff82d07a0e0bbf574cab5c"></a>

## disable property — Property reference / 6d3207546445 / 6

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

- [disable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-47affa490998093df338643b50e1c90112344a03acdef78edc4281137a077690): complete subsection reference.

- [disable_anonymization](resources--app_firewall--reference--group-001.md#canonical-ad9e53637860a64860af7cc627b9c06ca47bdd01c64d32907874af14cf12e9d2): complete subsection reference.

- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-66e746a4d30169eb67cd7aafdfbc6d21e2dacb0a438c4640b7615cc3ccc360a2): complete subsection reference.

<a id="canonical-546bffcc5a98df820cef681f79c54e8fbb2021f6627bc641cd2b3d747880777c"></a>

<a id="canonical-c76253db749eb705b3bb31439230cdc79b12fa561eb4c09dd9f5c091370ee1cd"></a>

## id property — Property reference / 6d3207546445 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-f9ecf3241c5fe2e22223abc5890f1dc6389c762ee61e0ba08170f5907c1adbc4"></a>

<a id="canonical-9a1c7b71aa2d9fde58038f072531a5dd91625890da56821a8ff73f54474fe2f8"></a>

## labels property — Property reference / 6d3207546445 / 8

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

- [monitoring](resources--app_firewall--reference--group-001.md#canonical-0af4bd45812502713e359638e33b95e7b7a75e36c5e4cbd05bd2e0e32d361a5b): complete subsection reference.

<a id="canonical-99212eb54ab584464d55fc5bf7cf01634d13507bb4c5bc463aa5ff3d7e231c34"></a>

<a id="canonical-36961cabddccd4aa3f6a562561e95de57acf39322eb3c1aa83206dba8803ef34"></a>

## name property — Property reference / 6d3207546445 / 9

Type: `"string"`. Required.

Name of the App Firewall. Must be unique within the namespace.

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

<a id="canonical-a54162f6577c422d428f88851d5f89fa10cb10fa0cd22120ae1a827807fb3cea"></a>

<a id="canonical-a579149576edf3fbf70a3bba5b63f5629faf4bb4ebe413a6be50899c7e6ab2e7"></a>

## namespace property — Property reference / 6d3207546445 / 10

Type: `"string"`. Required.

Namespace where the App Firewall is created.

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

- [timeouts](resources--app_firewall--reference--group-001.md#canonical-67d4d2bfa69993c60e75b7447ba2cd302087271ff43cb4137ef0271be5a07774): complete subsection reference.

- [use_default_blocking_page](resources--app_firewall--reference--group-001.md#canonical-bd4acad954f14442cc39f60e29182a6ea4cb4aa9fb62403d2ff142704a95d1b8): complete subsection reference.

<a id="canonical-7a3ae8b0f15e8f178732576897d35d46c52ee02bc5c47c58d8816c96d836c271"></a>

## All schema paths — Property reference / 6d3207546445 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_response_codes` | [allow_all_response_codes](resources--app_firewall--reference--group-001.md#canonical-0093ee8cc6de21bb83a3f78596af64d49bb19d6ab92460094070aabbc5749613) |
| `allowed_response_codes` | [allowed_response_codes](resources--app_firewall--reference--group-001.md#canonical-3e9dfe7b6e88eab4c936e2ebfce9b9582dc5748f12447b8c047d49b430bb7837) |
| `allowed_response_codes.response_code` | [allowed_response_codes.response_code](resources--app_firewall--reference--group-001.md#canonical-6757f06bebfda4b6a343e4c5f9acbd069e27b7f8241a55d7c834c29ccb89730f) |
| `annotations` | [annotations](resources--app_firewall--reference--group-001.md#canonical-b7a1643f43fcf910d6689f41da4417f85af63199400dee49cf718c7c4b60f455) |
| `blocking` | [blocking](resources--app_firewall--reference--group-001.md#canonical-58e805d65ae44d54805832aeca95deddbca2b1956a92f7e8eb87aeb324a70814) |
| `blocking_page` | [blocking_page](resources--app_firewall--reference--group-001.md#canonical-c8c1215d4c70797a3a77938684a5f8d751a0ad448f1a1c1da3bdacba0b7df8b4) |
| `blocking_page.blocking_page` | [blocking_page.blocking_page](resources--app_firewall--reference--group-001.md#canonical-c9be7abfb1bea31b8fb282253da6ebaf71a296d82c542d853eec60609ed4dbd1) |
| `blocking_page.response_code` | [blocking_page.response_code](resources--app_firewall--reference--group-001.md#canonical-c08ca5feb293b8a2aa283ffdd1519c6d38838054eb00b1157b52c397c1afb527) |
| `bot_protection_setting` | [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-7d88d8a4dfc187f0fffb82258f500d93dc7f8cfb2d283ed4b351d5030857ed5b) |
| `bot_protection_setting.good_bot_action` | [bot_protection_setting.good_bot_action](resources--app_firewall--reference--group-001.md#canonical-ee93d5abf625ecd35f90918c21d720333d71bfae02bb6c973d91dc83584e9daf) |
| `bot_protection_setting.malicious_bot_action` | [bot_protection_setting.malicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-697e342ea786153b9d7c6da9aaf989d32d1d66cced3ab7fc8f1d3ff7fa938872) |
| `bot_protection_setting.suspicious_bot_action` | [bot_protection_setting.suspicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-282564e2353269d561b4a048ffd6eb8c8c70592fc0631c05b80c04645fd95608) |
| `custom_anonymization` | [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-741ecda829962b0b8635898004a8cbd61109562fd7819b8e05f810517ca27d20) |
| `custom_anonymization.anonymization_config` | [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-6dbce352c9627857a0ed8a9a419f6d2e93613e43a593f71309b5ec8c8666527e) |
| `custom_anonymization.anonymization_config.cookie` | [custom_anonymization.anonymization_config.cookie](resources--app_firewall--reference--group-001.md#canonical-1def22ce36f36bfc84ece4df8a1ae7884cf0f81c624bdc6234fb6b5a8db40e67) |
| `custom_anonymization.anonymization_config.cookie.cookie_name` | [custom_anonymization.anonymization_config.cookie.cookie_name](resources--app_firewall--reference--group-001.md#canonical-343642986599e61e941c8fa19b3412cc618ea4068dfbac32feb9cb397092f749) |
| `custom_anonymization.anonymization_config.http_header` | [custom_anonymization.anonymization_config.http_header](resources--app_firewall--reference--group-001.md#canonical-a5b965ed3475960fa7f8de69e76ebb6ea87defede64c6beb28c26e103d2a157a) |
| `custom_anonymization.anonymization_config.http_header.header_name` | [custom_anonymization.anonymization_config.http_header.header_name](resources--app_firewall--reference--group-001.md#canonical-76904d0f31c4a8ae29e701c79b8e9de85f96901d0d399e09896ae4f524b73ed9) |
| `custom_anonymization.anonymization_config.query_parameter` | [custom_anonymization.anonymization_config.query_parameter](resources--app_firewall--reference--group-001.md#canonical-0da37484036f70e3d444ce2161dca30d7c4dc97ff8214069744d6c571f71e8c5) |
| `custom_anonymization.anonymization_config.query_parameter.query_param_name` | [custom_anonymization.anonymization_config.query_parameter.query_param_name](resources--app_firewall--reference--group-001.md#canonical-dbc2ae71eb1f18f023f740464710b21d27438b932ce51bb53948a8614d62ba00) |
| `default_anonymization` | [default_anonymization](resources--app_firewall--reference--group-001.md#canonical-9727bb489265bcd50444538b1c399efb7f1ec5c9d729251c5ca7a089a0159b36) |
| `default_bot_setting` | [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-b1aa2df863c14a99978e209770f55885d172db05011f1d5bffc299ec084cf773) |
| `default_detection_settings` | [default_detection_settings](resources--app_firewall--reference--group-001.md#canonical-011347eb16b655bc5769fc2731f7d0b1271ef3d185d5d3e33c64a8ac2620af67) |
| `description` | [description](resources--app_firewall--reference--group-001.md#canonical-3d21e40a84f93db709171c9798997e75a875a6d311e710043fd64bba0d7812f5) |
| `detection_settings` | [detection_settings](resources--app_firewall--reference--group-001.md#canonical-70e24531aa649a904098fc49ac0f87afdee27ab5c32a4807db2f01fc7db71c86) |
| `detection_settings.bot_protection_setting` | [detection_settings.bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-2e539c1f5dd6d10e13ebd0c8dfeab8171604f5e6ad670df5b80d3468358229aa) |
| `detection_settings.bot_protection_setting.good_bot_action` | [detection_settings.bot_protection_setting.good_bot_action](resources--app_firewall--reference--group-001.md#canonical-b9e46cc87ecb11251e778681ad1275e7f55e968783020c5ec53088412be2e109) |
| `detection_settings.bot_protection_setting.malicious_bot_action` | [detection_settings.bot_protection_setting.malicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-f4185cfe588df1e2c30a18a52c996eb18a163dbcdb94a02ca49ba0155c604e64) |
| `detection_settings.bot_protection_setting.suspicious_bot_action` | [detection_settings.bot_protection_setting.suspicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-c046bd7e395c179029f862f5a2fc54dcf418edbc00503ab2d084a5c24cded679) |
| `detection_settings.default_bot_setting` | [detection_settings.default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-052654dd081696fcb90eacd51a34c594aabe4d8dc9a33a5332eacf00949bc438) |
| `detection_settings.default_violation_settings` | [detection_settings.default_violation_settings](resources--app_firewall--reference--group-001.md#canonical-f0efe0e359fec05e4e6ae40aae6327dfe2e91bcc03ae70a2db2733150f12619b) |
| `detection_settings.disable_staging` | [detection_settings.disable_staging](resources--app_firewall--reference--group-001.md#canonical-feed9c592978514c4d1c3486c4996059d1b71c0fe7558373c015c2fd11fecd37) |
| `detection_settings.disable_suppression` | [detection_settings.disable_suppression](resources--app_firewall--reference--group-001.md#canonical-46b71052d29ddf541afd4c2c354ab3aef53f4be247196b17db74250c7ebd98e4) |
| `detection_settings.disable_threat_campaigns` | [detection_settings.disable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-82600db76196568ac4cc94ae85310b9d788dc57d4ee7d7b66bd4cc9e078e4210) |
| `detection_settings.enable_suppression` | [detection_settings.enable_suppression](resources--app_firewall--reference--group-001.md#canonical-2e59e7bd929fd1c7c588ef351964479e5310c321f7440cffb1bf7b90d926c313) |
| `detection_settings.enable_threat_campaigns` | [detection_settings.enable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-d7aa96b917ea9d1d57eb860a533966dddc3b9270ee0196654c723d739917e87f) |
| `detection_settings.signature_selection_setting` | [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-6976908045cc6c7da57931a88b71a65112f814ffa5380d7ca55267cbfbbeb884) |
| `detection_settings.signature_selection_setting.attack_type_settings` | [detection_settings.signature_selection_setting.attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-c95c6ae8811b03c9949497fd0d6d726c042236c7d8613acfa18ba8be9d388f88) |
| `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` | [detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types](resources--app_firewall--reference--group-001.md#canonical-708ef1383dfae0537c4e076fd150c0dc64035f53ddb94e8085db5eceba1fe7ed) |
| `detection_settings.signature_selection_setting.default_attack_type_settings` | [detection_settings.signature_selection_setting.default_attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-bbbdabd84189eac9eeae26ca7d6ccbb9a0df2695d790c4e2b731594bc4a6bafd) |
| `detection_settings.signature_selection_setting.default_signature_setting` | [detection_settings.signature_selection_setting.default_signature_setting](resources--app_firewall--reference--group-001.md#canonical-d8f33b74e4afcb5dfa5c9453becc57f3565e7c72adb60a21b046b53d2b20efd9) |
| `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-8fb306a310398f4b2077ab8537ec0a63ca4548834ab5ccee746b9ea716272093) |
| `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-c08d7ecc04fe0d75f924b942a2d5f98beb3443080f1af8b6e4d1fc6c4e178e16) |
| `detection_settings.signature_selection_setting.only_high_accuracy_signatures` | [detection_settings.signature_selection_setting.only_high_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-d2db926a6f5e6a583208fb8605f262dfed09a03b69d6269ac4a2132949a5d838) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy](resources--app_firewall--reference--group-001.md#canonical-099e881eed27d69bfe04a2779cedbc06a114d21c72e0c05e3fa8c4da0520dfe8) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action](resources--app_firewall--reference--group-001.md#canonical-85e1b1061ff13259a7c3a169d9402ff0b08b081a1abc423e243c361c4c7a56cc) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action](resources--app_firewall--reference--group-001.md#canonical-df89ef43e46359f1907cd3679a8b5f13aa547de1c0d1811348d5f5c5bdc1f525) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action](resources--app_firewall--reference--group-001.md#canonical-0c55dc8bdb3d7af8e7f9d66fbe0705463fa6e3ec63691034db80dc52fe3c5b40) |
| `detection_settings.stage_new_and_updated_signatures` | [detection_settings.stage_new_and_updated_signatures](resources--app_firewall--reference--group-001.md#canonical-bbaba6e34215d14a44f66fee78d56bd02f329e3723e21cdd8d7050ec67500cc6) |
| `detection_settings.stage_new_and_updated_signatures.staging_period` | [detection_settings.stage_new_and_updated_signatures.staging_period](resources--app_firewall--reference--group-001.md#canonical-952a06ffbd8220618ee7049bb07e6b03a5af0d9d0334601164b1d96684902448) |
| `detection_settings.stage_new_signatures` | [detection_settings.stage_new_signatures](resources--app_firewall--reference--group-001.md#canonical-9d824bf82d0196a7d0fab123fecf2f801a967b7f20fd448596e42580784dda31) |
| `detection_settings.stage_new_signatures.staging_period` | [detection_settings.stage_new_signatures.staging_period](resources--app_firewall--reference--group-001.md#canonical-208e3356fcf2e95e0e56d565a1b5122c96c5d16e38b3e47575d9cd333109f25e) |
| `detection_settings.violation_settings` | [detection_settings.violation_settings](resources--app_firewall--reference--group-001.md#canonical-c07dd07d2a73094aff2168ea28189477899c6c5e1ab0d1fc9e2841109d25d67e) |
| `detection_settings.violation_settings.disabled_violation_types` | [detection_settings.violation_settings.disabled_violation_types](resources--app_firewall--reference--group-001.md#canonical-923be1899177f5decb21c31d560e7a494e73cd280672cfc75a773da9eb68795c) |
| `detection_settings.violations_view` | [detection_settings.violations_view](resources--app_firewall--reference--group-001.md#canonical-7f0a74733caaa88496fd424dfdef7b1adfb648f4ff4c5f1d6db248c495238952) |
| `detection_settings.violations_view.description_spec` | [detection_settings.violations_view.description_spec](resources--app_firewall--reference--group-001.md#canonical-ce6e5851b29207632de5a17adb35feaa76d94c51375fea58c44f988871f065ee) |
| `detection_settings.violations_view.enabled` | [detection_settings.violations_view.enabled](resources--app_firewall--reference--group-001.md#canonical-1168e33b0f4401c9fa8508a4fd372785ae477b83af999f2569debebbf602089a) |
| `detection_settings.violations_view.enabled_by_default` | [detection_settings.violations_view.enabled_by_default](resources--app_firewall--reference--group-001.md#canonical-26e1a0adcf4de1ac46b06ffb548bfa24471a96987e8d3ee91297ff34f42fef68) |
| `detection_settings.violations_view.name` | [detection_settings.violations_view.name](resources--app_firewall--reference--group-001.md#canonical-db944b9c5d9495300161bdc6192ad461bc591ad0adf368180f8b2cb431c3c983) |
| `detection_settings.violations_view.title` | [detection_settings.violations_view.title](resources--app_firewall--reference--group-001.md#canonical-3cc3878cf49fc38de8b7581378ef6e8d06141f5dfcb65bb50e3a5b726e6d44a9) |
| `disable` | [disable](resources--app_firewall--reference--group-001.md#canonical-c58f2081977e182263854051ec1c1e744b42708f528b422eecda62e72e4d772e) |
| `disable_ai_enhancements` | [disable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-51ba662e5b8c9b65413be9c8dee479c2763e325af11635dc4b22e4f83ea5c45a) |
| `disable_anonymization` | [disable_anonymization](resources--app_firewall--reference--group-001.md#canonical-9fd50e3757171707ef410094d05c48cdc26d6219a8dc9e01f3b8c422e7ce77b8) |
| `enable_ai_enhancements` | [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-e15da73ee066596a786073362d2a50a22d6966b73494d3f0c43141af69049afa) |
| `enable_ai_enhancements.mitigate_high_medium_risk_action` | [enable_ai_enhancements.mitigate_high_medium_risk_action](resources--app_firewall--reference--group-001.md#canonical-371f99c7ae9884367f103e01278a0cf12c001c84402fe278e61c4c9ba9f24265) |
| `enable_ai_enhancements.mitigate_high_risk_action` | [enable_ai_enhancements.mitigate_high_risk_action](resources--app_firewall--reference--group-001.md#canonical-d81b31ab2b5ecd6aff09d59ca5cb165227e173bb77c8ed865540db462e02c07b) |
| `id` | [id](resources--app_firewall--reference--group-001.md#canonical-546bffcc5a98df820cef681f79c54e8fbb2021f6627bc641cd2b3d747880777c) |
| `labels` | [labels](resources--app_firewall--reference--group-001.md#canonical-f9ecf3241c5fe2e22223abc5890f1dc6389c762ee61e0ba08170f5907c1adbc4) |
| `monitoring` | [monitoring](resources--app_firewall--reference--group-001.md#canonical-31173fc53755e36aeb57e91c8ade2b3b3306a03640ed7372256752b2c0954a62) |
| `name` | [name](resources--app_firewall--reference--group-001.md#canonical-99212eb54ab584464d55fc5bf7cf01634d13507bb4c5bc463aa5ff3d7e231c34) |
| `namespace` | [namespace](resources--app_firewall--reference--group-001.md#canonical-a54162f6577c422d428f88851d5f89fa10cb10fa0cd22120ae1a827807fb3cea) |
| `timeouts` | [timeouts](resources--app_firewall--reference--group-001.md#canonical-703aca3e43c29ba6f2037b305f0b1855c8291afa5341773a9db32d0aa641c8fe) |
| `timeouts.create` | [timeouts.create](resources--app_firewall--reference--group-001.md#canonical-a75a6126eead8c4028458166ebf57c9c5fbbb69deb4311cc971d5d767991e243) |
| `timeouts.delete` | [timeouts.delete](resources--app_firewall--reference--group-001.md#canonical-87220a77c2a5c536ce774df4f386897a74e6fa12376b0d1a8a3d706280ce638a) |
| `timeouts.read` | [timeouts.read](resources--app_firewall--reference--group-001.md#canonical-11ebb6dc003ff2679f692eb6c7bffcd330114368d5daadc0a7c249df476bf2e7) |
| `timeouts.update` | [timeouts.update](resources--app_firewall--reference--group-001.md#canonical-3c4857464273f6055f0de427e517afafde9bd128851136582116c96484327c97) |
| `use_default_blocking_page` | [use_default_blocking_page](resources--app_firewall--reference--group-001.md#canonical-286e9e1d7913821e1199805378122627c59e4f6f9f7f7e45acc6c68e38bf0077) |

<a id="canonical-a68e69504a5f7e0d230f68a202c1c76fda676d279233bda811dc6844c138a57e"></a>

## Next pages — Property reference / 6d3207546445 / 12

- [allow_all_response_codes](resources--app_firewall--reference--group-001.md#canonical-99e75007717b48e03533e9b898700c624862621444d75258e3499b4a886b6a53)
- [allowed_response_codes](resources--app_firewall--reference--group-001.md#canonical-18d5fde757750c5ae1f21da40bf6c3348e1ce30d1a5e9499aff566ee1c05c9f6)
- [blocking](resources--app_firewall--reference--group-001.md#canonical-693b041ad0bdd1cbc9ed189a8d1f0c51264fa57d2eb99321044f6af82b046836)
- [blocking_page](resources--app_firewall--reference--group-001.md#canonical-072958fc9ffb537d907339d78644e0e909d2fafbeefef201fd6997bb27655858)
- [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-f561d7c60d03776bdafc8bcd1ff3cc2e20c3d873eba27b07550e9f5f35938d55)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543)
- [default_anonymization](resources--app_firewall--reference--group-001.md#canonical-6947b4b6ae9be6828bcabf9af8ec58df4d8d915bc460cd5ac7aca0837968974a)
- [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-44fe272375c06e26a85aa2b82dfb265fe7083495451fe26afb286e1f38d67f28)
- [default_detection_settings](resources--app_firewall--reference--group-001.md#canonical-5be5ae1e2e5176410bb44105286f6d85a52b42ab0b6197700dabfddfb431af0c)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [disable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-47affa490998093df338643b50e1c90112344a03acdef78edc4281137a077690)
- [disable_anonymization](resources--app_firewall--reference--group-001.md#canonical-ad9e53637860a64860af7cc627b9c06ca47bdd01c64d32907874af14cf12e9d2)
- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-66e746a4d30169eb67cd7aafdfbc6d21e2dacb0a438c4640b7615cc3ccc360a2)
- [monitoring](resources--app_firewall--reference--group-001.md#canonical-0af4bd45812502713e359638e33b95e7b7a75e36c5e4cbd05bd2e0e32d361a5b)
- [timeouts](resources--app_firewall--reference--group-001.md#canonical-67d4d2bfa69993c60e75b7447ba2cd302087271ff43cb4137ef0271be5a07774)
- [use_default_blocking_page](resources--app_firewall--reference--group-001.md#canonical-bd4acad954f14442cc39f60e29182a6ea4cb4aa9fb62403d2ff142704a95d1b8)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-99e75007717b48e03533e9b898700c624862621444d75258e3499b4a886b6a53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99cf4dfefe23c3ffbc0044a6414051f416e2cf7a59e8a4e1851bb05aaddb993b"></a>

## allow_all_response_codes — allow_all_response_codes / 33d3a5705dcb / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- allow_all_response_codes

<a id="canonical-0093ee8cc6de21bb83a3f78596af64d49bb19d6ab92460094070aabbc5749613"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: allow\_all\_response\_codes, allowed\_response\_codes\] Configuration parameter for allow
all response codes. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [allow_all_response_codes](resources--app_firewall--reference--group-001.md#canonical-0093ee8cc6de21bb83a3f78596af64d49bb19d6ab92460094070aabbc5749613)
- [allowed_response_codes](resources--app_firewall--reference--group-001.md#canonical-3e9dfe7b6e88eab4c936e2ebfce9b9582dc5748f12447b8c047d49b430bb7837)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_response_codes = {}
```

<a id="canonical-81a0132dbf636f5a84a993fae9e2e0f74cdac8b7e22a2c189f900f4250dfc8a7"></a>

## Direct properties — allow_all_response_codes / 33d3a5705dcb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bcb4751a34dcd010da46c7e9cc03ef10741b2941565a5eb6e92f4cd47f9a598c"></a>

## Next pages — allow_all_response_codes / 33d3a5705dcb / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-18d5fde757750c5ae1f21da40bf6c3348e1ce30d1a5e9499aff566ee1c05c9f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1e45d7c77331a6cf861459921556263630c8546d4cf44bc0607638f81a5ebcb"></a>

## allowed_response_codes — allowed_response_codes / c9304dca4b90 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- allowed_response_codes

<a id="canonical-3e9dfe7b6e88eab4c936e2ebfce9b9582dc5748f12447b8c047d49b430bb7837"></a>

Type: `"object"`. single nested block, Optional.

List of HTTP response status codes that are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
allowed_response_codes {
  # Configure direct properties listed below.
}
```

<a id="canonical-d92141d3a979f0ec4f5c73d78e520a52fcaae6a1f99f2fc997a6f64ec69ee58c"></a>

## Direct properties — allowed_response_codes / c9304dca4b90 / 3

<a id="canonical-6757f06bebfda4b6a343e4c5f9acbd069e27b7f8241a55d7c834c29ccb89730f"></a>

<a id="canonical-49127397e90b3d6f10286344ae8ab6ddf7cfc09600c01ad8a7a360a544f1ae16"></a>

## response_code property — allowed_response_codes / c9304dca4b90 / 4

Type: `["list", "number"]`. Optional.

List of HTTP response status codes that are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 48),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 48,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 48,
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
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8da669620f79594ac4d4ed93613caf0a91bb0ea3ded97c0bd533e71696bd8d00"></a>

## Next pages — allowed_response_codes / c9304dca4b90 / 5

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-693b041ad0bdd1cbc9ed189a8d1f0c51264fa57d2eb99321044f6af82b046836"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01f34e27dbf4aeb4b9d1fefa6938fb627a6988f6a852e3b1d39846aee06886f4"></a>

## blocking — blocking / ad305bf7d2a9 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- blocking

<a id="canonical-58e805d65ae44d54805832aeca95deddbca2b1956a92f7e8eb87aeb324a70814"></a>

Type: `["object", {}]`. Optional.

\[OneOf: blocking, monitoring\] Enable this option

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

- [blocking](resources--app_firewall--reference--group-001.md#canonical-58e805d65ae44d54805832aeca95deddbca2b1956a92f7e8eb87aeb324a70814)
- [monitoring](resources--app_firewall--reference--group-001.md#canonical-31173fc53755e36aeb57e91c8ade2b3b3306a03640ed7372256752b2c0954a62)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocking = {}
```

<a id="canonical-730c585fe4f63b1266aaedfa2bb81d1802fac366e393f82a6fc10e77ce586f6a"></a>

## Direct properties — blocking / ad305bf7d2a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cfabf6cd9208754e2a0d139beccccacfe58782ac8418887e0f77a2c22d161b0f"></a>

## Next pages — blocking / ad305bf7d2a9 / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-072958fc9ffb537d907339d78644e0e909d2fafbeefef201fd6997bb27655858"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e31396fea64f4d5b8c821af15f64d21cbc910418a83ec3459100f42aea0cddd6"></a>

## blocking_page — blocking_page / 3227d13a87fe / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- blocking_page

<a id="canonical-c8c1215d4c70797a3a77938684a5f8d751a0ad448f1a1c1da3bdacba0b7df8b4"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: blocking\_page, use\_default\_blocking\_page; Default: use\_default\_blocking\_page\]
Custom Blocking Response Page. Custom blocking response page body.

Upstream description:

Custom blocking response page body.

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

- [blocking_page](resources--app_firewall--reference--group-001.md#canonical-c8c1215d4c70797a3a77938684a5f8d751a0ad448f1a1c1da3bdacba0b7df8b4)
- [use_default_blocking_page](resources--app_firewall--reference--group-001.md#canonical-286e9e1d7913821e1199805378122627c59e4f6f9f7f7e45acc6c68e38bf0077)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocking_page {
  # Configure direct properties listed below.
}
```

<a id="canonical-1267d83362f7248b9fc4ba3f674e8db09f8b981f3e6257807da516e69aa82352"></a>

## Direct properties — blocking_page / 3227d13a87fe / 3

<a id="canonical-c9be7abfb1bea31b8fb282253da6ebaf71a296d82c542d853eec60609ed4dbd1"></a>

<a id="canonical-0ad1f459609886072b94a94179c0722ee21ab3e4bc0e6503dd3f0f7207a3cd51"></a>

## blocking_page property — blocking_page / 3227d13a87fe / 4

Type: `"string"`. Optional.

Define the content of the response page (e.g., an HTML document or a JSON object), use the
\{\{request\_id\}\} placeholder to provide users with a unique identifier to be able to trace the
blocked request in the logs. The maximum allowed size of response body is 4096 bytes after base64
encoding..

Upstream description:

Define the content of the response page (e.g., an HTML document or a JSON object), use the
\{\{request\_id\}\} placeholder to provide users with a unique identifier to be able to trace the
blocked request in the logs. The maximum allowed size of response body is 4096 bytes after base64
encoding, which would be about 3070 bytes in plain text.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 4096,
    "metadata": {
      "category": "content",
      "confidence": 0.99,
      "note": "Must be a valid URI (uri_ref). API rejects inline HTML despite the description example.",
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-c08ca5feb293b8a2aa283ffdd1519c6d38838054eb00b1157b52c397c1afb527"></a>

<a id="canonical-e2cfcb864a47e263c4132dd9a7e74f6bec7e06d7f418e2fb1fc40b30cc2cce4e"></a>

## response_code property — blocking_page / 3227d13a87fe / 5

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-238bac0a62425929ae25efb506cf95edc9585c814a502212c8edc3b53190522b"></a>

## Next pages — blocking_page / 3227d13a87fe / 6

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-f561d7c60d03776bdafc8bcd1ff3cc2e20c3d873eba27b07550e9f5f35938d55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58b229e4efca7b04fc5769a5315873d94855e63d8a22bcd9bbdc43af3d7915ae"></a>

## bot_protection_setting — bot_protection_setting / bcb36319714f / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- bot_protection_setting

<a id="canonical-7d88d8a4dfc187f0fffb82258f500d93dc7f8cfb2d283ed4b351d5030857ed5b"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bot\_protection\_setting, default\_bot\_setting; Default: default\_bot\_setting\]
Configuration parameter for bot protection setting.

Upstream description:

Configuration of WAF Bot Protection.

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

- [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-7d88d8a4dfc187f0fffb82258f500d93dc7f8cfb2d283ed4b351d5030857ed5b)
- [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-b1aa2df863c14a99978e209770f55885d172db05011f1d5bffc299ec084cf773)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bot_protection_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5b4848d34db6f021cb6f693f5478f7aa40110a9252f0f5e774b085b9059d961"></a>

## Direct properties — bot_protection_setting / bcb36319714f / 3

<a id="canonical-ee93d5abf625ecd35f90918c21d720333d71bfae02bb6c973d91dc83584e9daf"></a>

<a id="canonical-5dc53110553f03aef645d9a8d977a052ee7ad62a51ec5f537549e3a666afd567"></a>

## good_bot_action property — bot_protection_setting / bcb36319714f / 4

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-697e342ea786153b9d7c6da9aaf989d32d1d66cced3ab7fc8f1d3ff7fa938872"></a>

<a id="canonical-a70cc90f465f0ae3670de350766c98b2fa6ec3798d25088530f34762d2a90424"></a>

## malicious_bot_action property — bot_protection_setting / bcb36319714f / 5

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-282564e2353269d561b4a048ffd6eb8c8c70592fc0631c05b80c04645fd95608"></a>

<a id="canonical-ebe18b4ccffe801ddb9078c3c12afd25ebdf69575edc063d2ea8d5fbc6fcd6c8"></a>

## suspicious_bot_action property — bot_protection_setting / bcb36319714f / 6

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d458ad42651610bfb720636526321729439addb20c5560a7a4f53c69223d9718"></a>

## Next pages — bot_protection_setting / bcb36319714f / 7

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e97fc7abfa273969547d46976d3da4130c83d15b52b697ee9f8e5d62c4e9b7a4"></a>

## custom_anonymization — custom_anonymization / f99f26e0f452 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- custom_anonymization

<a id="canonical-741ecda829962b0b8635898004a8cbd61109562fd7819b8e05f810517ca27d20"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_anonymization, default\_anonymization, disable\_anonymization; Default:
default\_anonymization\] Anonymization settings which is a list of HTTP headers, parameters and
cookies.

Upstream description:

Anonymization settings which is a list of HTTP headers, parameters and cookies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("anonymization_config")}
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

OneOf alternatives in this subsection:

- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-741ecda829962b0b8635898004a8cbd61109562fd7819b8e05f810517ca27d20)
- [default_anonymization](resources--app_firewall--reference--group-001.md#canonical-9727bb489265bcd50444538b1c399efb7f1ec5c9d729251c5ca7a089a0159b36)
- [disable_anonymization](resources--app_firewall--reference--group-001.md#canonical-9fd50e3757171707ef410094d05c48cdc26d6219a8dc9e01f3b8c422e7ce77b8)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_anonymization {
  # Configure direct properties listed below.
}
```

<a id="canonical-204107807a1c7d3bcd73401d1e0f9e00435b580e7ad5942e129e54f8a18398c6"></a>

## Direct properties — custom_anonymization / f99f26e0f452 / 3

- [anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b): complete subsection reference.

<a id="canonical-22ffa2389776abfe7495afb72ebc0a7347cc3f528ace6bd71d5a66b650ece1ec"></a>

## Next pages — custom_anonymization / f99f26e0f452 / 4

- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8bf859467a6ac59981409c4db8a3b13f9a8abe7b775cb182ced942d69de2b8a"></a>

## custom_anonymization.anonymization_config — custom_anonymization.anonymization_config / 333af49b9222 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543)
- custom_anonymization.anonymization_config

<a id="canonical-6dbce352c9627857a0ed8a9a419f6d2e93613e43a593f71309b5ec8c8666527e"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP headers, cookies and query parameters whose values will be masked.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "http_header"),
  validators.ConflictingListObjectAttributes("cookie",
    "query_parameter"),
  validators.ConflictingListObjectAttributes("http_header",
    "query_parameter")}
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
anonymization_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-121002d4f6cc703799e0ede207a1353ef132e64e7f73fd50f20fd52be9f6cc94"></a>

## Direct properties — custom_anonymization.anonymization_config / 333af49b9222 / 3

- [cookie](resources--app_firewall--reference--group-001.md#canonical-9102c0c6bce576d63ad09219a14416ef8a7b890f61ad06ef19611e41d91fb1cc): complete subsection reference.

- [http_header](resources--app_firewall--reference--group-001.md#canonical-f96aa2981e881752c2f8ac0bc98be331780cd9f60d3e7caacfc556b734e92c8f): complete subsection reference.

- [query_parameter](resources--app_firewall--reference--group-001.md#canonical-6076a2d2e0b48d82902073dc80fad277d99b87e9aedeb1e826685eb5cefd242b): complete subsection reference.

<a id="canonical-c2e5062bc0a98be1d2e38181235e08d7597ee7cdfd0ae726ffcc67cfa0415a93"></a>

## Next pages — custom_anonymization.anonymization_config / 333af49b9222 / 4

- [custom_anonymization.anonymization_config.cookie](resources--app_firewall--reference--group-001.md#canonical-9102c0c6bce576d63ad09219a14416ef8a7b890f61ad06ef19611e41d91fb1cc)
- [custom_anonymization.anonymization_config.http_header](resources--app_firewall--reference--group-001.md#canonical-f96aa2981e881752c2f8ac0bc98be331780cd9f60d3e7caacfc556b734e92c8f)
- [custom_anonymization.anonymization_config.query_parameter](resources--app_firewall--reference--group-001.md#canonical-6076a2d2e0b48d82902073dc80fad277d99b87e9aedeb1e826685eb5cefd242b)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-9102c0c6bce576d63ad09219a14416ef8a7b890f61ad06ef19611e41d91fb1cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2fc2bec6f7dca11a025411c73c393e7864f73c30b4da65f072ae0345bfe46b7"></a>

## custom_anonymization.anonymization_config.cookie — custom_anonymization.anonymization_config.cookie / 17a79e71f053 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543)
- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b)
- custom_anonymization.anonymization_config.cookie

<a id="canonical-1def22ce36f36bfc84ece4df8a1ae7884cf0f81c624bdc6234fb6b5a8db40e67"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Cookies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_name")}
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
cookie {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8f81f919525dafa8efa974d6e2198f422014687fb28af24e84ede4976bffdb7"></a>

## Direct properties — custom_anonymization.anonymization_config.cookie / 17a79e71f053 / 3

<a id="canonical-343642986599e61e941c8fa19b3412cc618ea4068dfbac32feb9cb397092f749"></a>

<a id="canonical-687b5a8c3c25371c5bd9bff36c33adaed9beacb82a34b438237c51fcbd77c451"></a>

## cookie_name property — custom_anonymization.anonymization_config.cookie / 17a79e71f053 / 4

Type: `"string"`. Optional.

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

Upstream description:

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-b41f461f31ea0b01c6eb46dce9755a22c7971a63190205b7146f3114048824b2"></a>

## Next pages — custom_anonymization.anonymization_config.cookie / 17a79e71f053 / 5

- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-f96aa2981e881752c2f8ac0bc98be331780cd9f60d3e7caacfc556b734e92c8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03f3202f13078151bf1c95bd3e31e4bb76bef74b9e14ee4c7098a64d6a084761"></a>

## custom_anonymization.anonymization_config.http_header — custom_anonymization.anonymization_config.http_header / 352e45846606 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543)
- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b)
- custom_anonymization.anonymization_config.http_header

<a id="canonical-a5b965ed3475960fa7f8de69e76ebb6ea87defede64c6beb28c26e103d2a157a"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("header_name")}
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd991839410b4357645d39c3202bad4e691951eba5ec043213740edee5c8ad35"></a>

## Direct properties — custom_anonymization.anonymization_config.http_header / 352e45846606 / 3

<a id="canonical-76904d0f31c4a8ae29e701c79b8e9de85f96901d0d399e09896ae4f524b73ed9"></a>

<a id="canonical-b2ac0f23b517ce959740730a6053ef5c858eedcdc1a8952e39b2fd2e4efb00fc"></a>

## header_name property — custom_anonymization.anonymization_config.http_header / 352e45846606 / 4

Type: `"string"`. Optional.

Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can
be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (\*), or by using
only an asterisk to match any HTTP header name.

Upstream description:

Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can
be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (\*), or by using
only an asterisk to match any HTTP header name.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true"
  }
}
```

<a id="canonical-80c4a55de84c7b7e8f792ddd1e9102fb325820026417e054c3ba31aace553428"></a>

## Next pages — custom_anonymization.anonymization_config.http_header / 352e45846606 / 5

- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-6076a2d2e0b48d82902073dc80fad277d99b87e9aedeb1e826685eb5cefd242b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64caef20979e707e7417257e7b0cee9d5d97bddaa3eaf6d7873f102096dc3c9b"></a>

## custom_anonymization.anonymization_config.query_parameter — custom_anonymization.anonymization_config.query_parameter / f90894db925e / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-254ad5b9f3bac7f5008304f7b3e6554c47f4cf5d029dca329d4725b019ff6543)
- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b)
- custom_anonymization.anonymization_config.query_parameter

<a id="canonical-0da37484036f70e3d444ce2161dca30d7c4dc97ff8214069744d6c571f71e8c5"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("query_param_name")}
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
query_parameter {
  # Configure direct properties listed below.
}
```

<a id="canonical-88f980597533d67f4c0613752875cad79771bfea501ef8ebef14a5bf1a0c0410"></a>

## Direct properties — custom_anonymization.anonymization_config.query_parameter / f90894db925e / 3

<a id="canonical-dbc2ae71eb1f18f023f740464710b21d27438b932ce51bb53948a8614d62ba00"></a>

<a id="canonical-1bce7962f9a3f1aad7a6c0df1c69c0ee844dd52b551fd8f2de7ac14ba87c2d48"></a>

## query_param_name property — custom_anonymization.anonymization_config.query_parameter / f90894db925e / 4

Type: `"string"`. Optional.

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Upstream description:

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-55cccf44a84ad4428d158e78bb1f711b996724736bd0b40bf69322e52069853b"></a>

## Next pages — custom_anonymization.anonymization_config.query_parameter / f90894db925e / 5

- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-384b9603bbef09e9364c3baf8ce50d98ff77e7a4a68541ed934cb58dfb00409b)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-6947b4b6ae9be6828bcabf9af8ec58df4d8d915bc460cd5ac7aca0837968974a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26cb6f0032b55319146954cd10f298570e5c5cb9a1ca893505e5acf879025e4f"></a>

## default_anonymization — default_anonymization / 2b2baa10d9c5 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- default_anonymization

<a id="canonical-9727bb489265bcd50444538b1c399efb7f1ec5c9d729251c5ca7a089a0159b36"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default anonymization. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
default_anonymization = {}
```

<a id="canonical-e47a896de7aec0992efd117756fd911a134b1f4093a48d1053823304758d2ea0"></a>

## Direct properties — default_anonymization / 2b2baa10d9c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d206968a0d38d882fe6b02803ae5f0034c7648d21fb8271b5bc6ce155161ffc3"></a>

## Next pages — default_anonymization / 2b2baa10d9c5 / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-44fe272375c06e26a85aa2b82dfb265fe7083495451fe26afb286e1f38d67f28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-565be0598c77a17e55179a052f568e5796ba47bb41cb1c72e93ecedc67e2a5eb"></a>

## default_bot_setting — default_bot_setting / 01c266bbbc41 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- default_bot_setting

<a id="canonical-b1aa2df863c14a99978e209770f55885d172db05011f1d5bffc299ec084cf773"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default bot setting. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
default_bot_setting = {}
```

<a id="canonical-08abfbe6caae49f643bde1bcf2c0b280749ce860b84fb1d465a63191da15d353"></a>

## Direct properties — default_bot_setting / 01c266bbbc41 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf7d9ec654a966295e3bdf42a2d0c621d8f1001a73259199ac853741cd6ce1af"></a>

## Next pages — default_bot_setting / 01c266bbbc41 / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-5be5ae1e2e5176410bb44105286f6d85a52b42ab0b6197700dabfddfb431af0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdb5784d06423f186cfd9294a040b584cc76fcabfbb117a18ca40a756ef57080"></a>

## default_detection_settings — default_detection_settings / c830419680da / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- default_detection_settings

<a id="canonical-011347eb16b655bc5769fc2731f7d0b1271ef3d185d5d3e33c64a8ac2620af67"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: default\_detection\_settings, detection\_settings; Default: default\_detection\_settings\]
Configuration parameter for default detection settings. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

- [default_detection_settings](resources--app_firewall--reference--group-001.md#canonical-011347eb16b655bc5769fc2731f7d0b1271ef3d185d5d3e33c64a8ac2620af67)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-70e24531aa649a904098fc49ac0f87afdee27ab5c32a4807db2f01fc7db71c86)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_detection_settings = {}
```

<a id="canonical-c4a6fe4404e5609f32d59bfc0a392b3322a2c7a970af2da2b4be88b037ba29f8"></a>

## Direct properties — default_detection_settings / c830419680da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae5a5f50ba3b6eba478730e0e89236450ac0c3e48ba8e9c43569f53f0350abf7"></a>

## Next pages — default_detection_settings / c830419680da / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df99547fd0cad612093e828497607f49514170e9370853776289e6f42c606c2d"></a>

## detection_settings — detection_settings / e61eeb46775c / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- detection_settings

<a id="canonical-70e24531aa649a904098fc49ac0f87afdee27ab5c32a4807db2f01fc7db71c86"></a>

Type: `"object"`. single nested block, Optional.

Specifies detection settings to be used by WAF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_protection_setting",
    "default_bot_setting"),
  validators.ConflictingObjectAttributes("default_violation_settings",
    "violation_settings"),
  validators.ConflictingObjectAttributes("disable_staging",
    "stage_new_and_updated_signatures"),
  validators.ConflictingObjectAttributes("disable_staging",
    "stage_new_signatures"),
  validators.ConflictingObjectAttributes("disable_suppression",
    "enable_suppression"),
  validators.ConflictingObjectAttributes("disable_threat_campaigns",
    "enable_threat_campaigns"),
  validators.ConflictingObjectAttributes("stage_new_and_updated_signatures",
    "stage_new_signatures")}
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
  "x-ves-oneof-field-bot_protection_choice": "[\"bot_protection_setting\",\"default_bot_setting\"]",
  "x-ves-oneof-field-false_positive_suppression": "[\"disable_suppression\",\"enable_suppression\"]",
  "x-ves-oneof-field-signatures_staging_settings": "[\"disable_staging\",\"stage_new_and_updated_signatures\",\"stage_new_signatures\"]",
  "x-ves-oneof-field-threat_campaign_choice": "[\"disable_threat_campaigns\",\"enable_threat_campaigns\"]",
  "x-ves-oneof-field-violation_detection_setting": "[\"default_violation_settings\",\"violation_settings\"]"
}
```

Terraform syntax:

```terraform
detection_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-957784985fae10beeac31c90a797e97ca4f9d8ab046e7119a7871828c7673367"></a>

## Direct properties — detection_settings / e61eeb46775c / 3

- [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-07e4081987613c3a68980e41480869d3b95480690c55cdf843fe753d9b8a0f94): complete subsection reference.

- [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-61400fdff8c519472ed7502c419146d944ca7ce0743105519028759c0b1883dd): complete subsection reference.

- [default_violation_settings](resources--app_firewall--reference--group-001.md#canonical-33ccc62329659e64238c15080b95c40c70b5b66de1a68bc38b8fb0d4b030376b): complete subsection reference.

- [disable_staging](resources--app_firewall--reference--group-001.md#canonical-7d1abdd1ef436ee0f6799a670d1f884d5ce9d256d520498060d98f8bbbcee75b): complete subsection reference.

- [disable_suppression](resources--app_firewall--reference--group-001.md#canonical-8b2148a68294a1f9a612a5cb93c9d69beba44bb613585ed238bad6175bfab0fa): complete subsection reference.

- [disable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-afbe76045c2e51be79cbee95d021919cf376b542760ae7b7031456ecd222101e): complete subsection reference.

- [enable_suppression](resources--app_firewall--reference--group-001.md#canonical-de472f3861e768e6c0f4e40a5c1f33e8043d012a33d6e5b2900e19d20798d20c): complete subsection reference.

- [enable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-04f450b9f5e7a8b3ae8a3e60228be0b433c0f0f0535417c3134aa1c07f30ec9f): complete subsection reference.

- [signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33): complete subsection reference.

- [stage_new_and_updated_signatures](resources--app_firewall--reference--group-001.md#canonical-acae56c05cba1eccdb7cf8c63de412288fb23866cdc8ba1d0dc7213366c70cd2): complete subsection reference.

- [stage_new_signatures](resources--app_firewall--reference--group-001.md#canonical-6228cc79e4b6a4b20fd635d6e2070432f1eb909c57f4fa9881be0f5024297671): complete subsection reference.

- [violation_settings](resources--app_firewall--reference--group-001.md#canonical-1c74050138fcc23ffa3ba856e17b5f13cbed18739b345f2dc49e893b31d7c510): complete subsection reference.

- [violations_view](resources--app_firewall--reference--group-001.md#canonical-ab92e2105163d3b386c6dbeb464717f7f31d562bce448ba3491a24b09e752dc9): complete subsection reference.

<a id="canonical-7c90266023023684b6102cf4077c239259de3d6d171f65bffd87bc35b33cd946"></a>

## Next pages — detection_settings / e61eeb46775c / 4

- [detection_settings.bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-07e4081987613c3a68980e41480869d3b95480690c55cdf843fe753d9b8a0f94)
- [detection_settings.default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-61400fdff8c519472ed7502c419146d944ca7ce0743105519028759c0b1883dd)
- [detection_settings.default_violation_settings](resources--app_firewall--reference--group-001.md#canonical-33ccc62329659e64238c15080b95c40c70b5b66de1a68bc38b8fb0d4b030376b)
- [detection_settings.disable_staging](resources--app_firewall--reference--group-001.md#canonical-7d1abdd1ef436ee0f6799a670d1f884d5ce9d256d520498060d98f8bbbcee75b)
- [detection_settings.disable_suppression](resources--app_firewall--reference--group-001.md#canonical-8b2148a68294a1f9a612a5cb93c9d69beba44bb613585ed238bad6175bfab0fa)
- [detection_settings.disable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-afbe76045c2e51be79cbee95d021919cf376b542760ae7b7031456ecd222101e)
- [detection_settings.enable_suppression](resources--app_firewall--reference--group-001.md#canonical-de472f3861e768e6c0f4e40a5c1f33e8043d012a33d6e5b2900e19d20798d20c)
- [detection_settings.enable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-04f450b9f5e7a8b3ae8a3e60228be0b433c0f0f0535417c3134aa1c07f30ec9f)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [detection_settings.stage_new_and_updated_signatures](resources--app_firewall--reference--group-001.md#canonical-acae56c05cba1eccdb7cf8c63de412288fb23866cdc8ba1d0dc7213366c70cd2)
- [detection_settings.stage_new_signatures](resources--app_firewall--reference--group-001.md#canonical-6228cc79e4b6a4b20fd635d6e2070432f1eb909c57f4fa9881be0f5024297671)
- [detection_settings.violation_settings](resources--app_firewall--reference--group-001.md#canonical-1c74050138fcc23ffa3ba856e17b5f13cbed18739b345f2dc49e893b31d7c510)
- [detection_settings.violations_view](resources--app_firewall--reference--group-001.md#canonical-ab92e2105163d3b386c6dbeb464717f7f31d562bce448ba3491a24b09e752dc9)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-07e4081987613c3a68980e41480869d3b95480690c55cdf843fe753d9b8a0f94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82f9bff738c415de242b558e409a78f2df6074bc282ff648f11f78bc0ac3c227"></a>

## detection_settings.bot_protection_setting — detection_settings.bot_protection_setting / 9c6ba444d07c / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.bot_protection_setting

<a id="canonical-2e539c1f5dd6d10e13ebd0c8dfeab8171604f5e6ad670df5b80d3468358229aa"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bot protection setting.

Upstream description:

Configuration of WAF Bot Protection.

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
bot_protection_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d7d4095db21945c3ae653584257dd3fe9d1c14589282740048168e1d17cd4ff"></a>

## Direct properties — detection_settings.bot_protection_setting / 9c6ba444d07c / 3

<a id="canonical-b9e46cc87ecb11251e778681ad1275e7f55e968783020c5ec53088412be2e109"></a>

<a id="canonical-40884f28b1a36e48764cb410db9192d425fe02646e2a1e493b1ae346e52570b0"></a>

## good_bot_action property — detection_settings.bot_protection_setting / 9c6ba444d07c / 4

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f4185cfe588df1e2c30a18a52c996eb18a163dbcdb94a02ca49ba0155c604e64"></a>

<a id="canonical-ab93f212c2f7969e0734e632f0dfb47e0d707aa5ab312a5c653718ebd2068c11"></a>

## malicious_bot_action property — detection_settings.bot_protection_setting / 9c6ba444d07c / 5

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c046bd7e395c179029f862f5a2fc54dcf418edbc00503ab2d084a5c24cded679"></a>

<a id="canonical-5a5b49df28d614ae7a19140cee40eb7fbf5ad0b2e05ff1b064eb542cfd78bbdb"></a>

## suspicious_bot_action property — detection_settings.bot_protection_setting / 9c6ba444d07c / 6

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-dd8e1745da56455e5b1f1c3ba22a0a65fd672d8b0c316a484f73ee2c7ff3bbae"></a>

## Next pages — detection_settings.bot_protection_setting / 9c6ba444d07c / 7

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-61400fdff8c519472ed7502c419146d944ca7ce0743105519028759c0b1883dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8de092f6f89c85847857a0fe5175f578b4461d9123ffc955217c6bfb92ea4004"></a>

## detection_settings.default_bot_setting — detection_settings.default_bot_setting / d573a4c16558 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.default_bot_setting

<a id="canonical-052654dd081696fcb90eacd51a34c594aabe4d8dc9a33a5332eacf00949bc438"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default bot setting.

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
default_bot_setting = {}
```

<a id="canonical-68d85989e029fa4eb25861185b680d7348e5bce28928248a1f113c041955b840"></a>

## Direct properties — detection_settings.default_bot_setting / d573a4c16558 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44225b6c7a6a552ee9a68494ddcc9bf55ada28cd55aca89ebfa03aee1e2da467"></a>

## Next pages — detection_settings.default_bot_setting / d573a4c16558 / 4

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-33ccc62329659e64238c15080b95c40c70b5b66de1a68bc38b8fb0d4b030376b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e150f0922f631aa36f7560102cb147c7087f0b5153fc5bb8e7763f0301fa99e"></a>

## detection_settings.default_violation_settings — detection_settings.default_violation_settings / 81473f327bd1 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.default_violation_settings

<a id="canonical-f0efe0e359fec05e4e6ae40aae6327dfe2e91bcc03ae70a2db2733150f12619b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default violation settings.

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
default_violation_settings = {}
```

<a id="canonical-63acd5c85e92af0d21317224f691e9667870594c69a5060338edef0e858c1784"></a>

## Direct properties — detection_settings.default_violation_settings / 81473f327bd1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d561d7ea70f2e432363860ffe12cc9c791d13b6137d0949571a629bbd9f8efc"></a>

## Next pages — detection_settings.default_violation_settings / 81473f327bd1 / 4

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-7d1abdd1ef436ee0f6799a670d1f884d5ce9d256d520498060d98f8bbbcee75b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e02c76c5150434694d25ca8a76f078de2c42228b29ff8b95c76b5d894b426d3"></a>

## detection_settings.disable_staging — detection_settings.disable_staging / e477ea861700 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.disable_staging

<a id="canonical-feed9c592978514c4d1c3486c4996059d1b71c0fe7558373c015c2fd11fecd37"></a>

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
disable_staging = {}
```

<a id="canonical-81cad19d980b3322a057dd76e48bf5cb243e189a36af12cd83faf4c77eaf99ba"></a>

## Direct properties — detection_settings.disable_staging / e477ea861700 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6e618ebb9fb71ef0f818de154145828ee78a2182b27b5a2a3c4e288cb7174f7"></a>

## Next pages — detection_settings.disable_staging / e477ea861700 / 4

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-8b2148a68294a1f9a612a5cb93c9d69beba44bb613585ed238bad6175bfab0fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb403be6aab73b6e9b6c3f8f00c8f4da32e1e66770a26f2e900a0300ccd5d2ea"></a>

## detection_settings.disable_suppression — detection_settings.disable_suppression / b88f30a8de92 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.disable_suppression

<a id="canonical-46b71052d29ddf541afd4c2c354ab3aef53f4be247196b17db74250c7ebd98e4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable suppression.

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
disable_suppression = {}
```

<a id="canonical-db3f39b94d03cbd47584e1ac2e0cfc9013ee198e8ebce28268b0b387f1e3c145"></a>

## Direct properties — detection_settings.disable_suppression / b88f30a8de92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8ce1852e7d00c5189175e3112029297e5e466b3f59e98a3fd0bd883868ff8c28"></a>

## Next pages — detection_settings.disable_suppression / b88f30a8de92 / 4

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-afbe76045c2e51be79cbee95d021919cf376b542760ae7b7031456ecd222101e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a244da2cf8b11898764adab94d1a990a1e10f3af123117587febfc8198d5c5d2"></a>

## detection_settings.disable_threat_campaigns — detection_settings.disable_threat_campaigns / 423def178c9e / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.disable_threat_campaigns

<a id="canonical-82600db76196568ac4cc94ae85310b9d788dc57d4ee7d7b66bd4cc9e078e4210"></a>

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
disable_threat_campaigns = {}
```

<a id="canonical-52eedc22e35e5e0916fda4715f6e4eb8099aa535cd5e3272676e5688bbd234e0"></a>

## Direct properties — detection_settings.disable_threat_campaigns / 423def178c9e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f95ea32929b2b9227516ea88c79860de62a9e11e6169b5c967f3913fbd914fdb"></a>

## Next pages — detection_settings.disable_threat_campaigns / 423def178c9e / 4

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-de472f3861e768e6c0f4e40a5c1f33e8043d012a33d6e5b2900e19d20798d20c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2dfa4e1eded18dd6dcc4443b26ba1130abe40492a0e8e67c3546f417d292d388"></a>

## detection_settings.enable_suppression — detection_settings.enable_suppression / faf3367f1391 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.enable_suppression

<a id="canonical-2e59e7bd929fd1c7c588ef351964479e5310c321f7440cffb1bf7b90d926c313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable suppression.

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
enable_suppression = {}
```

<a id="canonical-70fd5829b8b54d3f16e4f19c7fb4efd94fa0fa1ce08bece7418ef8ed3bebfbd9"></a>

## Direct properties — detection_settings.enable_suppression / faf3367f1391 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80a36d851132b448cf79801ef523606337d4ba322327e29fc88a5dcdbf53a9fe"></a>

## Next pages — detection_settings.enable_suppression / faf3367f1391 / 4

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-04f450b9f5e7a8b3ae8a3e60228be0b433c0f0f0535417c3134aa1c07f30ec9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac285ecb6ce8408f042e78ba5aad5a56ee34424940f2454512cf690524e46c53"></a>

## detection_settings.enable_threat_campaigns — detection_settings.enable_threat_campaigns / 2024a27ec524 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.enable_threat_campaigns

<a id="canonical-d7aa96b917ea9d1d57eb860a533966dddc3b9270ee0196654c723d739917e87f"></a>

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
enable_threat_campaigns = {}
```

<a id="canonical-22732233ecf871b797511fe9cee8883f1b0c0777fbbb66e9b6f28f99b70e9cbf"></a>

## Direct properties — detection_settings.enable_threat_campaigns / 2024a27ec524 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51a5cb7ee0d5bb7f6ad59a9163c0c4e05e2db4d22fcff3ba92433cdd5660e9f0"></a>

## Next pages — detection_settings.enable_threat_campaigns / 2024a27ec524 / 4

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b9ae9040f7fc45285136a5a05dccef8d3165ffca6ec7e14dd8436039fb5871e"></a>

## detection_settings.signature_selection_setting — detection_settings.signature_selection_setting / 737620959ad9 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.signature_selection_setting

<a id="canonical-6976908045cc6c7da57931a88b71a65112f814ffa5380d7ca55267cbfbbeb884"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures are patterns that identify attacks on a web application and its components.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("attack_type_settings",
    "default_attack_type_settings"),
  validators.ConflictingObjectAttributes("default_signature_setting",
    "signature_settings_by_accuracy"),
  validators.ConflictingObjectAttributes("high_medium_accuracy_signatures",
    "high_medium_low_accuracy_signatures"),
  validators.ConflictingObjectAttributes("high_medium_accuracy_signatures",
    "only_high_accuracy_signatures"),
  validators.ConflictingObjectAttributes("high_medium_low_accuracy_signatures",
    "only_high_accuracy_signatures")}
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
  "x-ves-oneof-field-attack_type_setting": "[\"attack_type_settings\",\"default_attack_type_settings\"]",
  "x-ves-oneof-field-signature_protection_choice": "[\"default_signature_setting\",\"signature_settings_by_accuracy\"]",
  "x-ves-oneof-field-signature_selection_by_accuracy": "[\"high_medium_accuracy_signatures\",\"high_medium_low_accuracy_signatures\",\"only_high_accuracy_signatures\"]"
}
```

Terraform syntax:

```terraform
signature_selection_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-b1f89fdec8c8440f5c0befdff8e02ed0befc17b6f9d42ba9dfcdf5c1a9d72fb3"></a>

## Direct properties — detection_settings.signature_selection_setting / 737620959ad9 / 3

- [attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-d830b2bd8e56c4239afad917b9c06d58c6802547cc7278c29f0517f2a0d1a265): complete subsection reference.

- [default_attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-4e3128b25b020a369fb5db44d1c1c74152cc50c63cbd064173367125256db1fc): complete subsection reference.

- [default_signature_setting](resources--app_firewall--reference--group-001.md#canonical-6bad3e48fd059dd6d772349a730a63963df031f70ef2d8f4785cfa94e19c1260): complete subsection reference.

- [high_medium_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-bb6226e7837586665243f18f816dbdda47c2f34571fb2768fa4656ef656eeb99): complete subsection reference.

- [high_medium_low_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-fdc01c597c155d8cb1e023adc57d79079eca989a1e1b6671384e9de988983cb0): complete subsection reference.

- [only_high_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-6d7cbd4a2e34e25ef3e1beb749e285fc430dbe9110c6f3fe524481385d5999d8): complete subsection reference.

- [signature_settings_by_accuracy](resources--app_firewall--reference--group-001.md#canonical-9f0c788d4b5c717de7c53a638dfd375c86d984d8f8d162d1303a49a1c141703e): complete subsection reference.

<a id="canonical-f9f8b293cb214f29bde6ab557e967241a8a34772b43159be16357bdfb8ee61f8"></a>

## Next pages — detection_settings.signature_selection_setting / 737620959ad9 / 4

- [detection_settings.signature_selection_setting.attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-d830b2bd8e56c4239afad917b9c06d58c6802547cc7278c29f0517f2a0d1a265)
- [detection_settings.signature_selection_setting.default_attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-4e3128b25b020a369fb5db44d1c1c74152cc50c63cbd064173367125256db1fc)
- [detection_settings.signature_selection_setting.default_signature_setting](resources--app_firewall--reference--group-001.md#canonical-6bad3e48fd059dd6d772349a730a63963df031f70ef2d8f4785cfa94e19c1260)
- [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-bb6226e7837586665243f18f816dbdda47c2f34571fb2768fa4656ef656eeb99)
- [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-fdc01c597c155d8cb1e023adc57d79079eca989a1e1b6671384e9de988983cb0)
- [detection_settings.signature_selection_setting.only_high_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-6d7cbd4a2e34e25ef3e1beb749e285fc430dbe9110c6f3fe524481385d5999d8)
- [detection_settings.signature_selection_setting.signature_settings_by_accuracy](resources--app_firewall--reference--group-001.md#canonical-9f0c788d4b5c717de7c53a638dfd375c86d984d8f8d162d1303a49a1c141703e)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-d830b2bd8e56c4239afad917b9c06d58c6802547cc7278c29f0517f2a0d1a265"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-867f27055b0bbceb194503de3007c6fab68170471b56c4a4680a3422f72a1b0d"></a>

## detection_settings.signature_selection_setting.attack_type_settings — detection_settings.signature_selection_setting.attack_type_settings / 3dd880767747 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- detection_settings.signature_selection_setting.attack_type_settings

<a id="canonical-c95c6ae8811b03c9949497fd0d6d726c042236c7d8613acfa18ba8be9d388f88"></a>

Type: `"object"`. single nested block, Optional.

Specifies attack-type settings to be used by WAF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("disabled_attack_types")}
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
attack_type_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-8fa29943602d5c2118555d57c263ee4c299d6162aa128ff450ac61958a161590"></a>

## Direct properties — detection_settings.signature_selection_setting.attack_type_settings / 3dd880767747 / 3

<a id="canonical-708ef1383dfae0537c4e076fd150c0dc64035f53ddb94e8085db5eceba1fe7ed"></a>

<a id="canonical-bfd1bc7677daa52469ed8ab019f8907592f0d273ac32b6c42b581f9b17046c67"></a>

## disabled_attack_types property — detection_settings.signature_selection_setting.attack_type_settings / 3dd880767747 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of Attack Types that will be ignored and not trigger a detection. Possible values are
\`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

List of Attack Types that will be ignored and not trigger a detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(22),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 22,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "22",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "22",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-673cd6cfd68fd53660fc5de87388003d0c8998649dd531bd6dc62e2b0534b4eb"></a>

## Next pages — detection_settings.signature_selection_setting.attack_type_settings / 3dd880767747 / 5

- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-4e3128b25b020a369fb5db44d1c1c74152cc50c63cbd064173367125256db1fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5212cb7eff88a56c897aebe5994292586424eff67850b6067c2b185d22d7ead3"></a>

## detection_settings.signature_selection_setting.default_attack_type_settings — detection_settings.signature_selection_setting.default_attack_type_settings / a1733e5b49c7 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- detection_settings.signature_selection_setting.default_attack_type_settings

<a id="canonical-bbbdabd84189eac9eeae26ca7d6ccbb9a0df2695d790c4e2b731594bc4a6bafd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default attack type settings.

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
default_attack_type_settings = {}
```

<a id="canonical-f7c4de2ca4299373f08828ae33bedbd8dc3fa94b1d1465e88bab4c99e7489880"></a>

## Direct properties — detection_settings.signature_selection_setting.default_attack_type_settings / a1733e5b49c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2abc6d93a122ccd47d578d9a988bdbd54f9fd1f54125f28b06c504b6f45f5e67"></a>

## Next pages — detection_settings.signature_selection_setting.default_attack_type_settings / a1733e5b49c7 / 4

- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-6bad3e48fd059dd6d772349a730a63963df031f70ef2d8f4785cfa94e19c1260"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24daacb738b5a8cf8c6e60f91e8d04e396ba63bfc1f989648707439aa45f3af1"></a>

## detection_settings.signature_selection_setting.default_signature_setting — detection_settings.signature_selection_setting.default_signature_setting / dc42e56bb331 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- detection_settings.signature_selection_setting.default_signature_setting

<a id="canonical-d8f33b74e4afcb5dfa5c9453becc57f3565e7c72adb60a21b046b53d2b20efd9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default signature setting.

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
default_signature_setting = {}
```

<a id="canonical-0b88d7bcbc3f244d9b4750fce24f97a69ce8be1974abd58211a149163334134c"></a>

## Direct properties — detection_settings.signature_selection_setting.default_signature_setting / dc42e56bb331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d053ea729da9e82e41ea7e60ef1d0e3b3d820d5fe5abdc7487c9b20131cc16c7"></a>

## Next pages — detection_settings.signature_selection_setting.default_signature_setting / dc42e56bb331 / 4

- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-bb6226e7837586665243f18f816dbdda47c2f34571fb2768fa4656ef656eeb99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-710938acef1a82741c37e4155c98900fd1a906bf816f9cd773ee8631c78fc72c"></a>

## detection_settings.signature_selection_setting.high_medium_accuracy_signatures — detection_settings.signature_selection_setting.high_medium_accuracy_signatures / 94e448eff80b / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- detection_settings.signature_selection_setting.high_medium_accuracy_signatures

<a id="canonical-8fb306a310398f4b2077ab8537ec0a63ca4548834ab5ccee746b9ea716272093"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for high medium accuracy signatures.

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
high_medium_accuracy_signatures = {}
```

<a id="canonical-22629c13713e94401ff7a6b6dcd53a045df9e6fd3d4e6d80a6cc160dce37d899"></a>

## Direct properties — detection_settings.signature_selection_setting.high_medium_accuracy_signatures / 94e448eff80b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1bd3f3bd7c7a823bcc95ba075523c8ed49bbfa78dc661d384aa0808cb52bb528"></a>

## Next pages — detection_settings.signature_selection_setting.high_medium_accuracy_signatures / 94e448eff80b / 4

- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-fdc01c597c155d8cb1e023adc57d79079eca989a1e1b6671384e9de988983cb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0911540598837649eabac998174642bb1ea24c5f72b96d6be80f47f86a9deaa2"></a>

## detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures — detection_settings.signature_selection_setting.high_medium_low_accuracy_signatur / 6b11a5dcd210 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures

<a id="canonical-c08d7ecc04fe0d75f924b942a2d5f98beb3443080f1af8b6e4d1fc6c4e178e16"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for high medium low accuracy signatures.

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
high_medium_low_accuracy_signatures = {}
```

<a id="canonical-2d825a120b79e8d5148f6a17b7a0c09aecf5f9b78d4ee7c6cbe33cdc0ca5c3cf"></a>

## Direct properties — detection_settings.signature_selection_setting.high_medium_low_accuracy_signatur / 6b11a5dcd210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36bc8f7d5835c16174f11b6f83158407998baff145c7b3ce2cc1f34d20dd97bf"></a>

## Next pages — detection_settings.signature_selection_setting.high_medium_low_accuracy_signatur / 6b11a5dcd210 / 4

- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-6d7cbd4a2e34e25ef3e1beb749e285fc430dbe9110c6f3fe524481385d5999d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-274117524bc1b9bdf802b983dda3190d4af00f8fee694fba15959126ac3d7109"></a>

## detection_settings.signature_selection_setting.only_high_accuracy_signatures — detection_settings.signature_selection_setting.only_high_accuracy_signatures / 6251915c7826 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- detection_settings.signature_selection_setting.only_high_accuracy_signatures

<a id="canonical-d2db926a6f5e6a583208fb8605f262dfed09a03b69d6269ac4a2132949a5d838"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for only high accuracy signatures.

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
only_high_accuracy_signatures = {}
```

<a id="canonical-69c399423e8eaeb2e23ec1275c1e2a8a56a04642a41d26f28bfac25190a39eba"></a>

## Direct properties — detection_settings.signature_selection_setting.only_high_accuracy_signatures / 6251915c7826 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26a7744e8badc2aa0f8e8790912ea3c06f0eeb5dabaf70d993faabaf482c8164"></a>

## Next pages — detection_settings.signature_selection_setting.only_high_accuracy_signatures / 6251915c7826 / 4

- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-9f0c788d4b5c717de7c53a638dfd375c86d984d8f8d162d1303a49a1c141703e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd402e7b50fa82a2f190526bdd9c56a3a5f40e65ee8f68c2c91bc4899f157743"></a>

## detection_settings.signature_selection_setting.signature_settings_by_accuracy — detection_settings.signature_selection_setting.signature_settings_by_accuracy / f961367ce143 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- detection_settings.signature_selection_setting.signature_settings_by_accuracy

<a id="canonical-099e881eed27d69bfe04a2779cedbc06a114d21c72e0c05e3fa8c4da0520dfe8"></a>

Type: `"object"`. single nested block, Optional.

Configuration of WAF Signature Protection.

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
signature_settings_by_accuracy {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b40c1b6a7e7b41262507c3d27a9912a19a545419dc0b4669fc9c94dc92ece5d"></a>

## Direct properties — detection_settings.signature_selection_setting.signature_settings_by_accuracy / f961367ce143 / 3

<a id="canonical-85e1b1061ff13259a7c3a169d9402ff0b08b081a1abc423e243c361c4c7a56cc"></a>

<a id="canonical-b8f8fb5b97770c30313c75583a122bf51de7f32442540dc54226af66549dac32"></a>

## high_accuracy_action property — detection_settings.signature_selection_setting.signature_settings_by_accuracy / f961367ce143 / 4

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-df89ef43e46359f1907cd3679a8b5f13aa547de1c0d1811348d5f5c5bdc1f525"></a>

<a id="canonical-b20db08c73e009a4c4d2d6a397c3ec157a0e4746adde7401860bd19814b96956"></a>

## low_accuracy_action property — detection_settings.signature_selection_setting.signature_settings_by_accuracy / f961367ce143 / 5

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0c55dc8bdb3d7af8e7f9d66fbe0705463fa6e3ec63691034db80dc52fe3c5b40"></a>

<a id="canonical-e7fd83d7a71c1ed5035fd6eebec6c87c126a3eea80b0f3f44006cdfeb4e310c6"></a>

## medium_accuracy_action property — detection_settings.signature_selection_setting.signature_settings_by_accuracy / f961367ce143 / 6

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5318fc6ee6620c96774efcbca0fa2546a826ea4232672036c49bfbcbbe9e0ecd"></a>

## Next pages — detection_settings.signature_selection_setting.signature_settings_by_accuracy / f961367ce143 / 7

- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-313da9987467171ecc9db013a10506b3c4d3d44eb638b1333fdd318f34125e33)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-acae56c05cba1eccdb7cf8c63de412288fb23866cdc8ba1d0dc7213366c70cd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffa2d903b9508da3b5194f34677f20804f5100e99096422dabbe07c1bd294d4b"></a>

## detection_settings.stage_new_and_updated_signatures — detection_settings.stage_new_and_updated_signatures / de79d5434094 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.stage_new_and_updated_signatures

<a id="canonical-bbaba6e34215d14a44f66fee78d56bd02f329e3723e21cdd8d7050ec67500cc6"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures staging configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("staging_period")}
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
stage_new_and_updated_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-09d31715ff50b12333a24f2e920ec34f858fc048aae82772cdc4f8acaa605b7d"></a>

## Direct properties — detection_settings.stage_new_and_updated_signatures / de79d5434094 / 3

<a id="canonical-952a06ffbd8220618ee7049bb07e6b03a5af0d9d0334601164b1d96684902448"></a>

<a id="canonical-ab6244baa3a0bf93bb7d3b292e87b1e38c00bd83f39ecc25ed14f9e17b236227"></a>

## staging_period property — detection_settings.stage_new_and_updated_signatures / de79d5434094 / 4

Type: `"number"`. Optional.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Upstream description:

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Staging period in days. Default 7, max 20. Applies to both stage_new_and_updated_signatures and stage_new_signatures.",
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-86efe36536a48f5f0a3ce6fbea53ea22915fdd4daadf267ff098b3e84dba2690"></a>

## Next pages — detection_settings.stage_new_and_updated_signatures / de79d5434094 / 5

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-6228cc79e4b6a4b20fd635d6e2070432f1eb909c57f4fa9881be0f5024297671"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0f4026d29de9f332cbd8ab9bddab12bf325bcb6667592f50777a92ea75615ac"></a>

## detection_settings.stage_new_signatures — detection_settings.stage_new_signatures / 9d682627f5c4 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.stage_new_signatures

<a id="canonical-9d824bf82d0196a7d0fab123fecf2f801a967b7f20fd448596e42580784dda31"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures staging configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("staging_period")}
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
stage_new_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-a820f72fb14f57cbab5a7d630a3b8335742828eee1172b1e467374fc96327e79"></a>

## Direct properties — detection_settings.stage_new_signatures / 9d682627f5c4 / 3

<a id="canonical-208e3356fcf2e95e0e56d565a1b5122c96c5d16e38b3e47575d9cd333109f25e"></a>

<a id="canonical-d44cabd65ac42d0b57e18184aed3a0b57959ef6832afe2fb10fbe5bbbb27271e"></a>

## staging_period property — detection_settings.stage_new_signatures / 9d682627f5c4 / 4

Type: `"number"`. Optional.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Upstream description:

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Staging period in days. Default 7, max 20. Applies to both stage_new_and_updated_signatures and stage_new_signatures.",
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-bab6e80393f3954faa1b53fb7a6c145e2b0f068d37e57aa8942cad2d7beb9cba"></a>

## Next pages — detection_settings.stage_new_signatures / 9d682627f5c4 / 5

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-1c74050138fcc23ffa3ba856e17b5f13cbed18739b345f2dc49e893b31d7c510"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de7be1301f83726887258657add1e8e9b0b5a90fedd31654ac23917a81989e88"></a>

## detection_settings.violation_settings — detection_settings.violation_settings / 6326142f53fd / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.violation_settings

<a id="canonical-c07dd07d2a73094aff2168ea28189477899c6c5e1ab0d1fc9e2841109d25d67e"></a>

Type: `"object"`. single nested block, Optional.

Specifies violation settings to be used by WAF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("disabled_violation_types")}
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
violation_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-62b3e1f0714727b5e482edf3da0985272e2f0d0ccd8daffd036d9d4f3bc11890"></a>

## Direct properties — detection_settings.violation_settings / 6326142f53fd / 3

<a id="canonical-923be1899177f5decb21c31d560e7a494e73cd280672cfc75a773da9eb68795c"></a>

<a id="canonical-c89704b34a283af1213471480964aab41fc3485a009ca39b5c496fe423ff6917"></a>

## disabled_violation_types property — detection_settings.violation_settings / 6326142f53fd / 4

Type: `["list", "string"]`. Optional, Deprecated.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
Disabled Violations. List of violations to be excluded. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of violations to be excluded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "40",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "40",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cd940646a29d6af192f5697cdc3396a0cb2608b8a4fc079374caaee6749487a6"></a>

## Next pages — detection_settings.violation_settings / 6326142f53fd / 5

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-ab92e2105163d3b386c6dbeb464717f7f31d562bce448ba3491a24b09e752dc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4c56f742d6f391b8cb23d5ed42fc3c7eb13a0276d3591e749f58b633c98bea0"></a>

## detection_settings.violations_view — detection_settings.violations_view / ad6d2e15e9c9 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- detection_settings.violations_view

<a id="canonical-7f0a74733caaa88496fd424dfdef7b1adfb648f4ff4c5f1d6db248c495238952"></a>

Type: `"object"`. list nested block, Optional.

List of violation checks that are performed on HTTP request to ensure the requests are properly
formatted, detection of evasion techniques and other violations.

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
violations_view {
  # Configure direct properties listed below.
}
```

<a id="canonical-e9454545b27ea84848632da93fc3586e615457dd4b2b64378ec2c94b55dc9e7b"></a>

## Direct properties — detection_settings.violations_view / ad6d2e15e9c9 / 3

<a id="canonical-ce6e5851b29207632de5a17adb35feaa76d94c51375fea58c44f988871f065ee"></a>

<a id="canonical-8e8752361fc25d7bd016899117ab369f9168d76e1649d2272ae47bec0290f1ad"></a>

## description_spec property — detection_settings.violations_view / ad6d2e15e9c9 / 4

Type: `"string"`. Optional.

Description. Human-readable description text

<a id="canonical-1168e33b0f4401c9fa8508a4fd372785ae477b83af999f2569debebbf602089a"></a>

<a id="canonical-b03dc133d905b2b8622a22f4a6fe43db71c354b67c9a5c0559b0c89b08f8e3d3"></a>

## enabled property — detection_settings.violations_view / ad6d2e15e9c9 / 5

Type: `"bool"`. Optional.

State. Enable or disable the feature

Upstream description:

Enable or disable the feature

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

<a id="canonical-26e1a0adcf4de1ac46b06ffb548bfa24471a96987e8d3ee91297ff34f42fef68"></a>

<a id="canonical-28f20bda9b5ab0febf9e1b4fd141a48daf585a0f4845726b066eaad7d08d9cd0"></a>

## enabled_by_default property — detection_settings.violations_view / ad6d2e15e9c9 / 6

Type: `"string"`. Optional.

Violations that are enabled by default by F5 are advisable to leave enabled.

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

<a id="canonical-db944b9c5d9495300161bdc6192ad461bc591ad0adf368180f8b2cb431c3c983"></a>

<a id="canonical-9b083927957016fc3b552ecf39e4d62bd77fd7c9f4972761b03b91fd81d16142"></a>

## name property — detection_settings.violations_view / ad6d2e15e9c9 / 7

Type: `"string"`. Optional.

Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3cc3878cf49fc38de8b7581378ef6e8d06141f5dfcb65bb50e3a5b726e6d44a9"></a>

<a id="canonical-8e050f09b206dfd7563b254700452a59813ec731e96ea8f45402a131fea8ed73"></a>

## title property — detection_settings.violations_view / ad6d2e15e9c9 / 8

Type: `"string"`. Optional.

Title. Human-readable title for the resource

Upstream description:

Human-readable title for the resource

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

<a id="canonical-e282cd0f4cb6d834bdd53278262bf2198e6592ce5a677ad7352bbf210c1064c4"></a>

## Next pages — detection_settings.violations_view / ad6d2e15e9c9 / 9

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-9141dcda73075ecf630f673963f34fe0753de9929ce97279cbb5a120b46b0a25)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-47affa490998093df338643b50e1c90112344a03acdef78edc4281137a077690"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9276478517290eb478d54f09d726a69e18295e88ad899bb156e339956ecb28c2"></a>

## disable_ai_enhancements — disable_ai_enhancements / d6e1aed0fefc / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- disable_ai_enhancements

<a id="canonical-51ba662e5b8c9b65413be9c8dee479c2763e325af11635dc4b22e4f83ea5c45a"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_ai\_enhancements, enable\_ai\_enhancements; Default: disable\_ai\_enhancements\]
Configuration parameter for disable ai enhancements. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

- [disable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-51ba662e5b8c9b65413be9c8dee479c2763e325af11635dc4b22e4f83ea5c45a)
- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-e15da73ee066596a786073362d2a50a22d6966b73494d3f0c43141af69049afa)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ai_enhancements = {}
```

<a id="canonical-df90b47d16e50bb87b3b8aae35707b157f503f82a44ca5f6597e4e392f54786e"></a>

## Direct properties — disable_ai_enhancements / d6e1aed0fefc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1458485a13966dc20c819548ec9d9889e92e04439542030457350bb227dc1b02"></a>

## Next pages — disable_ai_enhancements / d6e1aed0fefc / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-ad9e53637860a64860af7cc627b9c06ca47bdd01c64d32907874af14cf12e9d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f80b8d1489b89e3dcd361da40fbefe38c482e756390efd403643b804da7542e"></a>

## disable_anonymization — disable_anonymization / 62d82577242b / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- disable_anonymization

<a id="canonical-9fd50e3757171707ef410094d05c48cdc26d6219a8dc9e01f3b8c422e7ce77b8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable anonymization.

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
disable_anonymization = {}
```

<a id="canonical-6e8af3278749c1ef5225e06112ca27e16cc2a7db485adc3113ba1bc0b5a06ffe"></a>

## Direct properties — disable_anonymization / 62d82577242b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a104e4a9b83079aee20b4fdf7fdf7032801e994f25b229bf74777815b762ae5a"></a>

## Next pages — disable_anonymization / 62d82577242b / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-66e746a4d30169eb67cd7aafdfbc6d21e2dacb0a438c4640b7615cc3ccc360a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e58893ecf3b79d3efca60e574b6dfb2cdf256ed0bc09ce314673612814543a35"></a>

## enable_ai_enhancements — enable_ai_enhancements / 2d4616639cca / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- enable_ai_enhancements

<a id="canonical-e15da73ee066596a786073362d2a50a22d6966b73494d3f0c43141af69049afa"></a>

Type: `"object"`. single nested block, Optional.

Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("mitigate_high_medium_risk_action",
    "mitigate_high_risk_action")}
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
  "x-ves-oneof-field-risk_score_action_choice": "[\"mitigate_high_medium_risk_action\",\"mitigate_high_risk_action\"]"
}
```

Terraform syntax:

```terraform
enable_ai_enhancements {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ec5b1fe8d9a386fcd93f84f8d12a4e69aa94da6868ad664f3c3488623ec0cc9"></a>

## Direct properties — enable_ai_enhancements / 2d4616639cca / 3

- [mitigate_high_medium_risk_action](resources--app_firewall--reference--group-001.md#canonical-7e04a89fa71d47eb31388636e9f6adf45eae625c79dd5465414fa7c9c4579c6c): complete subsection reference.

- [mitigate_high_risk_action](resources--app_firewall--reference--group-001.md#canonical-ebbebd2fbc73f6427b4f50d10b16503ef0fb1060968325d55ef08db4176f52e9): complete subsection reference.

<a id="canonical-d2dbb1e5b31fadb3ff19d5c74a2734cb074675dce2e281c1b3bb10bb5056dcf2"></a>

## Next pages — enable_ai_enhancements / 2d4616639cca / 4

- [enable_ai_enhancements.mitigate_high_medium_risk_action](resources--app_firewall--reference--group-001.md#canonical-7e04a89fa71d47eb31388636e9f6adf45eae625c79dd5465414fa7c9c4579c6c)
- [enable_ai_enhancements.mitigate_high_risk_action](resources--app_firewall--reference--group-001.md#canonical-ebbebd2fbc73f6427b4f50d10b16503ef0fb1060968325d55ef08db4176f52e9)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-7e04a89fa71d47eb31388636e9f6adf45eae625c79dd5465414fa7c9c4579c6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ff96340227554424b2ad84c7a61a717d6b9f143ed887cf7d007d205e5c20401"></a>

## enable_ai_enhancements.mitigate_high_medium_risk_action — enable_ai_enhancements.mitigate_high_medium_risk_action / 88e5d812976a / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-66e746a4d30169eb67cd7aafdfbc6d21e2dacb0a438c4640b7615cc3ccc360a2)
- enable_ai_enhancements.mitigate_high_medium_risk_action

<a id="canonical-371f99c7ae9884367f103e01278a0cf12c001c84402fe278e61c4c9ba9f24265"></a>

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
mitigate_high_medium_risk_action = {}
```

<a id="canonical-f0c723b056bca3f65ae2cc838c20e079d7700e90299fbe8cb0314838636f0d5d"></a>

## Direct properties — enable_ai_enhancements.mitigate_high_medium_risk_action / 88e5d812976a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87de9352ee5cb83271d8f47a539780acb403b48126d958d88edc03934726f7f9"></a>

## Next pages — enable_ai_enhancements.mitigate_high_medium_risk_action / 88e5d812976a / 4

- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-66e746a4d30169eb67cd7aafdfbc6d21e2dacb0a438c4640b7615cc3ccc360a2)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-ebbebd2fbc73f6427b4f50d10b16503ef0fb1060968325d55ef08db4176f52e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-560f93d9c04c014137a40dcb0f1c5e54c80b1be559744eb67100ca75e9a7bcfc"></a>

## enable_ai_enhancements.mitigate_high_risk_action — enable_ai_enhancements.mitigate_high_risk_action / aea7d2086183 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-66e746a4d30169eb67cd7aafdfbc6d21e2dacb0a438c4640b7615cc3ccc360a2)
- enable_ai_enhancements.mitigate_high_risk_action

<a id="canonical-d81b31ab2b5ecd6aff09d59ca5cb165227e173bb77c8ed865540db462e02c07b"></a>

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
mitigate_high_risk_action = {}
```

<a id="canonical-49e911ad976a709a3e22b07e0359d5973572f1162aa3d8c953b17489aa0e62b7"></a>

## Direct properties — enable_ai_enhancements.mitigate_high_risk_action / aea7d2086183 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4410a0fe7d6f96205cf01ff0fba6174065e91aed369cec7bd5c81fff4eddb3e"></a>

## Next pages — enable_ai_enhancements.mitigate_high_risk_action / aea7d2086183 / 4

- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-66e746a4d30169eb67cd7aafdfbc6d21e2dacb0a438c4640b7615cc3ccc360a2)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-0af4bd45812502713e359638e33b95e7b7a75e36c5e4cbd05bd2e0e32d361a5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a09a090661581079a2933a02fd1d58c7f14fcdfc73b1604200be185824a03f7a"></a>

## monitoring — monitoring / a4cc60ab6b05 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- monitoring

<a id="canonical-31173fc53755e36aeb57e91c8ade2b3b3306a03640ed7372256752b2c0954a62"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
monitoring = {}
```

<a id="canonical-eb26073ec86816d5dcfc6d867d838253ec021be70a8ddaed66c88867185430d4"></a>

## Direct properties — monitoring / a4cc60ab6b05 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3eb0a40f1e4765e1c9aace918e724263fb63f581b33fc4f50d7bda1db0564975"></a>

## Next pages — monitoring / a4cc60ab6b05 / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-67d4d2bfa69993c60e75b7447ba2cd302087271ff43cb4137ef0271be5a07774"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98b2ae7ecc959c574937b537f8e053b4ffaf9a3e84f555b396663ca3d29e5ef1"></a>

## timeouts — timeouts / 4b22d9416aeb / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- timeouts

<a id="canonical-703aca3e43c29ba6f2037b305f0b1855c8291afa5341773a9db32d0aa641c8fe"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-6776de9be760fe71c1836f2a437e94805892633e71e191d384fefab4d692b046"></a>

## Direct properties — timeouts / 4b22d9416aeb / 3

<a id="canonical-a75a6126eead8c4028458166ebf57c9c5fbbb69deb4311cc971d5d767991e243"></a>

<a id="canonical-5353649f685745e8bdfbc056946d50d1e1a7e50e1e3c388039845036f885d23d"></a>

## create property — timeouts / 4b22d9416aeb / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-87220a77c2a5c536ce774df4f386897a74e6fa12376b0d1a8a3d706280ce638a"></a>

<a id="canonical-4bb69663a9f708240c4a58aba18bd7429979d01e87ea913a6f608332472c881b"></a>

## delete property — timeouts / 4b22d9416aeb / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-11ebb6dc003ff2679f692eb6c7bffcd330114368d5daadc0a7c249df476bf2e7"></a>

<a id="canonical-bd0c50da0b91b211da0ec9c1a695ab62012c2e04bdc0758a3332cc369d34e938"></a>

## read property — timeouts / 4b22d9416aeb / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3c4857464273f6055f0de427e517afafde9bd128851136582116c96484327c97"></a>

<a id="canonical-10ebbdc7c5fc0299738e5377426e5727b180c6402402cde55580903319d4a110"></a>

## update property — timeouts / 4b22d9416aeb / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-6762c2b6b990f6177eeb4172bec1e60d217174210fe08635ca0e8ffdd4453a18"></a>

## Next pages — timeouts / 4b22d9416aeb / 8

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)

<a id="canonical-bd4acad954f14442cc39f60e29182a6ea4cb4aa9fb62403d2ff142704a95d1b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bea1ed002b6df1cbde5c985a8858e378f42f5fa1d40c25031fdffaed19c0ebbb"></a>

## use_default_blocking_page — use_default_blocking_page / 35c6be34c736 / 2

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- use_default_blocking_page

<a id="canonical-286e9e1d7913821e1199805378122627c59e4f6f9f7f7e45acc6c68e38bf0077"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
use_default_blocking_page = {}
```

<a id="canonical-b5938a9aa144e3e29e1ffc94d6e83da2e2b10bafe519612b943fbca05cf26022"></a>

## Direct properties — use_default_blocking_page / 35c6be34c736 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10f816dd9a23a3dfca80fc04c2d05ca8bfa3995a6c816247157f21597354fbe4"></a>

## Next pages — use_default_blocking_page / 35c6be34c736 / 4

- [Property reference](resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [xcsh_app_firewall](../resources/app_firewall.md#canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf)
