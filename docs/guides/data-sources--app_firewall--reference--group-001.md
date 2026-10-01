---
page_title: "xcsh_app_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall reference."
---

# xcsh_app_firewall reference

<a id="canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90137b1a5c909fa40485af12a1896f0e839fbc9d234914121b650c1e9c9e8053"></a>

## Property reference — Property reference / 619f3a4a423a / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- Property reference

<a id="canonical-7383a0cf10e19d6620ae9b68956975030f41e3735ab2f02f4d6d01933021f3fd"></a>

## Direct properties — Property reference / 619f3a4a423a / 3

- [allow_all_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-b7077d063193988d29112a0b46c27a029b4a9acd421fc1c457024be9954eaa07): complete subsection reference.

- [allowed_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-982f10529300907bec5f5acbe8468680b0535a0099774e109cd92de6a035a752): complete subsection reference.

<a id="canonical-5360ae27f9fbf968cc628ad18586a667e5bf1040b90fc5e92d15c701eda7b7c4"></a>

<a id="canonical-6c6480df19885c36e876ffeaae839836a5615573a69e093ce032494c1db36fcb"></a>

## annotations property — Property reference / 619f3a4a423a / 4

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

- [blocking](data-sources--app_firewall--reference--group-001.md#canonical-bcd3456b05507a43dc7ef59a1e7e2d4ff4371f0bbdaa7cfed1ec234776ad8f87): complete subsection reference.

- [blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-b8edf21322eff78b2dfdd77ce400c552507d006ee15b9411f9150f91ef30d34d): complete subsection reference.

- [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-932a9382cdb86d30a1222ced128b734f897e01133cb737083c962a94c98ed0e8): complete subsection reference.

- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945): complete subsection reference.

- [default_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-6cfe415781aec97010e0a7f2d9be34b5aef8d17eca28633e075c603f119e3a28): complete subsection reference.

- [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-62357eb47ae3cebef5f2aaeed586e8c444e2ee73e66cdcaadccbc0f3921f63e5): complete subsection reference.

- [default_detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-5cefd983e2bbcfaa75926052f8ab411809c5ecf04d9124b0239c355e7996ea59): complete subsection reference.

<a id="canonical-511baca66baefdd132cdd60a0832b0a8d6364c8f2f85c10f2c7a97f12fffcdf0"></a>

<a id="canonical-6c09005ad2a572aec967298ae625ae7c59e0078d3c88482d757640b20f5b6a31"></a>

## description property — Property reference / 619f3a4a423a / 5

Type: `"string"`. Computed.

Description of the AppFirewall.

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

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4): complete subsection reference.

- [disable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-44e035f938c10b60f96cf72524a61411b2d81a63d6ed271c32afa1919755c343): complete subsection reference.

- [disable_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-f3e7b68ac4045b02148757c9c852efeac2f9b18e7778c684d1a47c8d6e8412ba): complete subsection reference.

- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-fcec69530d75cc1a8525164bd619b25bedc1adfd5198f98ef2eaf9aee0e68ae5): complete subsection reference.

<a id="canonical-0a271e0fe9a10068f165379aaf01d931b29716e498eed68131703d467e3d674a"></a>

<a id="canonical-4e3c13d8dd75aaa8e261af8f61350171a42a4f27089c9da5a7f0d8f66212ee9b"></a>

## id property — Property reference / 619f3a4a423a / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4d0b40c416d0f2030b49f24ac51a4491aa3ad0282aa49c9b81b3dcf9a6c56226"></a>

<a id="canonical-6abac0f1d8debd615d92ee0c1da23dae5aac60c019a8ebf402d414a86438f45a"></a>

## labels property — Property reference / 619f3a4a423a / 7

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

- [monitoring](data-sources--app_firewall--reference--group-001.md#canonical-1ef81513048703b410f3651ef1afd3b69d8bd753f8118b1431462ebc01179857): complete subsection reference.

<a id="canonical-4b18fb1f9af03c508816e454566ed3b2e38456c3854eabf357f0ad39950514ca"></a>

<a id="canonical-ddaf8188546d0b344b1887f4e9608d22ceebb3720f7cb9058fd63ef7caeca7aa"></a>

## name property — Property reference / 619f3a4a423a / 8

Type: `"string"`. Required.

Name of the AppFirewall.

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

<a id="canonical-4fa7323cbbd078d7297fbc9b1b36218729655af9601d59c5b39ee1c768578539"></a>

<a id="canonical-0b0e51abeccf4d57281595312da16fdd111892423276b89ac33cd62deaba837d"></a>

## namespace property — Property reference / 619f3a4a423a / 9

Type: `"string"`. Required.

Namespace where the AppFirewall exists.

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

- [use_default_blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-9f246fc996a562416e707b574f84dce284240997ed47a6fd1bbf0bb3f4725ffc): complete subsection reference.

<a id="canonical-9e209aff235ae437012704e0d6f4792152eff9a464dc6d42f003de515c2f172e"></a>

## All schema paths — Property reference / 619f3a4a423a / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_response_codes` | [allow_all_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-8c62e7a01e8d767cb883b6479d0a14d289caf226853fa9d708add1cadb4e207b) |
| `allowed_response_codes` | [allowed_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-ba093a9546347115c5bbc38383fdee7a4e6e125879db70404082e77726d85911) |
| `allowed_response_codes.response_code` | [allowed_response_codes.response_code](data-sources--app_firewall--reference--group-001.md#canonical-382b3ea3af303c6fe2d7f9d2a62594b7690a7df8860b2b01800ef0146c3d83e2) |
| `annotations` | [annotations](data-sources--app_firewall--reference--group-001.md#canonical-5360ae27f9fbf968cc628ad18586a667e5bf1040b90fc5e92d15c701eda7b7c4) |
| `blocking` | [blocking](data-sources--app_firewall--reference--group-001.md#canonical-3009f5e7a3924817b4a9e0c3b6ce39e6ace5a7ca3dc8660dade064c6ead87cf2) |
| `blocking_page` | [blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-2b7d5052f77b4cc0ffd886feed4400fcab25452d97e22256d122fc3b9afc3209) |
| `blocking_page.blocking_page` | [blocking_page.blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-03ef2e4ebc6b385a6f4d82998c2aa317565c1324c0275e74fe892838beab614e) |
| `blocking_page.response_code` | [blocking_page.response_code](data-sources--app_firewall--reference--group-001.md#canonical-3d36093bb4d43ee63a22d908b6dc38de24cda3c4183182e399eb2b953e40064a) |
| `bot_protection_setting` | [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-c2b6b993051856063a601a77254cd927b4c28b3204c2ca81ebb28505c4f385af) |
| `bot_protection_setting.good_bot_action` | [bot_protection_setting.good_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-71616fd4209e5c2998ba8184183a0052959bba6829fdeea071c5ba5cdd318867) |
| `bot_protection_setting.malicious_bot_action` | [bot_protection_setting.malicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-44a058a27c966252960b7f85e80a202b329bdcfd34f4becf9ca664f76a8d9597) |
| `bot_protection_setting.suspicious_bot_action` | [bot_protection_setting.suspicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-084adfb8118161c32514e5ae5869da27cde974ece785c76223522242195c8404) |
| `custom_anonymization` | [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-8647e57eefa234b828f17d38b943e9b7c283570a24058676bb2ecd9d75ba2fad) |
| `custom_anonymization.anonymization_config` | [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-088777cee70e830bfd28ee1eacee3f7f58fd1eea56fe927a01b483966071d5cd) |
| `custom_anonymization.anonymization_config.cookie` | [custom_anonymization.anonymization_config.cookie](data-sources--app_firewall--reference--group-001.md#canonical-30eb906df4730a117b6265640c9ade52040d66731cb9b4a9567247b98d7e2399) |
| `custom_anonymization.anonymization_config.cookie.cookie_name` | [custom_anonymization.anonymization_config.cookie.cookie_name](data-sources--app_firewall--reference--group-001.md#canonical-6cd9bc5d2848fef28dab064deb0cd96d6da31c49673067de15388a35d719c119) |
| `custom_anonymization.anonymization_config.http_header` | [custom_anonymization.anonymization_config.http_header](data-sources--app_firewall--reference--group-001.md#canonical-109fa2fa3ad563ffe761803f1a0ceb50c93f8c7c98696fe064ffbb458aee0158) |
| `custom_anonymization.anonymization_config.http_header.header_name` | [custom_anonymization.anonymization_config.http_header.header_name](data-sources--app_firewall--reference--group-001.md#canonical-16f0eca78c4630b3ad1326ae38148695a01efe3ceba0968d48bfd40699896e96) |
| `custom_anonymization.anonymization_config.query_parameter` | [custom_anonymization.anonymization_config.query_parameter](data-sources--app_firewall--reference--group-001.md#canonical-c88ab4f3a2897dde658909b6792f77efcd32cbdb8fb4426d8c2981a9f18ab197) |
| `custom_anonymization.anonymization_config.query_parameter.query_param_name` | [custom_anonymization.anonymization_config.query_parameter.query_param_name](data-sources--app_firewall--reference--group-001.md#canonical-7aae01a0630a2af52f740c23d738c616c9872d291719a3685a1eac1d11547065) |
| `default_anonymization` | [default_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-cfd9ce7185449c7a623412d0d0814cdb98eb920c24de7060b36117ebbbbea32b) |
| `default_bot_setting` | [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-4e9cb75daf676a9c58bfb9dd87aa25d3fe4a0e907e348a949e73036d5664809c) |
| `default_detection_settings` | [default_detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-005dd7654e80c77b9ce466e5862b7ca2b78aef9318ce17b0c6a77ce0f39a9731) |
| `description` | [description](data-sources--app_firewall--reference--group-001.md#canonical-511baca66baefdd132cdd60a0832b0a8d6364c8f2f85c10f2c7a97f12fffcdf0) |
| `detection_settings` | [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-092c60de7d22883fab85adabaa2f3df8aca6fbfc8d195c539d4ceb638b3e6586) |
| `detection_settings.bot_protection_setting` | [detection_settings.bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-d4a109492762005d822826e1677e8197492b62ba439699fa9dabb52fa48a45d1) |
| `detection_settings.bot_protection_setting.good_bot_action` | [detection_settings.bot_protection_setting.good_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-992fe12389781d769a98beb3097e74594a6c63cd28f816debe143c0483dcb34c) |
| `detection_settings.bot_protection_setting.malicious_bot_action` | [detection_settings.bot_protection_setting.malicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-856c95edcbfc0bf54281f43780e41690f7e4b76da6a7e3f7f05f41b3ba7d379c) |
| `detection_settings.bot_protection_setting.suspicious_bot_action` | [detection_settings.bot_protection_setting.suspicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-f05fa4855f56f63603ccf693c8e076e3bd9f3e1df36b725e7d11b5282afd3c88) |
| `detection_settings.default_bot_setting` | [detection_settings.default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-61a532cc5d14566888c0e162f30c7c60582f5ace46832f597826593437ff06f9) |
| `detection_settings.default_violation_settings` | [detection_settings.default_violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-a6ab11418a639f52f644b82c828315ce37458d7759936b19904c590e555f3fe1) |
| `detection_settings.disable_staging` | [detection_settings.disable_staging](data-sources--app_firewall--reference--group-001.md#canonical-14e4c6e8c649312be1e6ef14cffd558c209310273e0654c6e417448c9f074e84) |
| `detection_settings.disable_suppression` | [detection_settings.disable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-9f06702df29c70bb711dbb06a4e94b0911bd26427323fb44ac60f533bb7f4198) |
| `detection_settings.disable_threat_campaigns` | [detection_settings.disable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-b2329bb5d06c95129c948ca88a91ec0b4ac375507c29f22ced68304459a64406) |
| `detection_settings.enable_suppression` | [detection_settings.enable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-f73f8721eb191ec129787c5299a9ef87124b0f12d58cd61dfd6b10537798d9ec) |
| `detection_settings.enable_threat_campaigns` | [detection_settings.enable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-c4dc94aad78f38c54a75c6ebd0f792e441e80da38955bcb67bcc72e026edad6a) |
| `detection_settings.signature_selection_setting` | [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-546710318a465f09ecd0bad573b3e882f78deb4b0e112474d4abd6c0a2ee70a5) |
| `detection_settings.signature_selection_setting.attack_type_settings` | [detection_settings.signature_selection_setting.attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-fcd9c4d91715eacc45571e17feb2778df0d5eb2720c585566f16bc6ea6def6d2) |
| `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` | [detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types](data-sources--app_firewall--reference--group-001.md#canonical-89360c68f05bef3cf0d3db23e13d719e6a12c044cfbbb5d26b1627b473309927) |
| `detection_settings.signature_selection_setting.default_attack_type_settings` | [detection_settings.signature_selection_setting.default_attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-a74d028baba263c769a907c7471e674bac2a988ac9fabaa7ca4d4cbd21cc159d) |
| `detection_settings.signature_selection_setting.default_signature_setting` | [detection_settings.signature_selection_setting.default_signature_setting](data-sources--app_firewall--reference--group-001.md#canonical-f168ecde6a98d100e49d410cfa4ae3bc16a17a37275cec876bbc309c4c2331eb) |
| `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-be7f45889d5e01a9fb23dba8a937013529f8195d818ed3186dba0bd438e99179) |
| `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-f8698191f8efd7a24fa005cbcb01a1eb8f343031cbf65d4264ce24ae16272c3d) |
| `detection_settings.signature_selection_setting.only_high_accuracy_signatures` | [detection_settings.signature_selection_setting.only_high_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-25a4fd6c4318ab011401b8f02780644f9c3722beec905f72e8559cb4acefb5f8) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy](data-sources--app_firewall--reference--group-001.md#canonical-c0beabbb8882569593da08385cf267a4c1eb49c181c622a07c87bd8342abc898) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action](data-sources--app_firewall--reference--group-001.md#canonical-55bcb971244bae864ef7a21de0c0a7fd930f248ec4a2296c872dcca76e3f5707) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action](data-sources--app_firewall--reference--group-001.md#canonical-82c923b03efce4e15f683d9a7c6ea77be1e12ef81cf8130b4d3b764961591c8c) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action](data-sources--app_firewall--reference--group-001.md#canonical-55a64bf0b549df476be1fc1efdeed72ba2656b6123adf74d7e187bf4743ebbdd) |
| `detection_settings.stage_new_and_updated_signatures` | [detection_settings.stage_new_and_updated_signatures](data-sources--app_firewall--reference--group-001.md#canonical-a1c67aa6ad9ddfe476edb6e19fc091ddc4151e64ff282111ead19f3d74bdac43) |
| `detection_settings.stage_new_and_updated_signatures.staging_period` | [detection_settings.stage_new_and_updated_signatures.staging_period](data-sources--app_firewall--reference--group-001.md#canonical-bdea08690aea9bb583149ce7ea7e7bee7e422e431621e7e24f5d977c35bb8e79) |
| `detection_settings.stage_new_signatures` | [detection_settings.stage_new_signatures](data-sources--app_firewall--reference--group-001.md#canonical-2cce129b6a32c732cecf82c10af56fcb6c74e9222a5844f8f2bbbb58fb9e08e9) |
| `detection_settings.stage_new_signatures.staging_period` | [detection_settings.stage_new_signatures.staging_period](data-sources--app_firewall--reference--group-001.md#canonical-6d8ea3698c36e850d067c9c9dc76a7197626a1faff402aab0532aed41baf2f03) |
| `detection_settings.violation_settings` | [detection_settings.violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-67110e6af0928c20d1252aa63dce223c5ba469085c72ab6ec981cef7a09d338c) |
| `detection_settings.violation_settings.disabled_violation_types` | [detection_settings.violation_settings.disabled_violation_types](data-sources--app_firewall--reference--group-001.md#canonical-0212ce5b34c9c5f9a99e195a9a6d03a5f9ff2a5da3c161e1c2fe16f88108bb5d) |
| `detection_settings.violations_view` | [detection_settings.violations_view](data-sources--app_firewall--reference--group-001.md#canonical-63fb9a7795c8bdf159545de3cd99be7bc29e7f89752a9169f97df7bad3c033e9) |
| `detection_settings.violations_view.description_spec` | [detection_settings.violations_view.description_spec](data-sources--app_firewall--reference--group-001.md#canonical-ef30bb24fd313f55b717c6811c9284b3a021c84b8de6d89d6921481c0ae6d531) |
| `detection_settings.violations_view.enabled` | [detection_settings.violations_view.enabled](data-sources--app_firewall--reference--group-001.md#canonical-a18afff24908577a9b13e9455c1b51e7e69efd844935c367307efc2175b6a721) |
| `detection_settings.violations_view.enabled_by_default` | [detection_settings.violations_view.enabled_by_default](data-sources--app_firewall--reference--group-001.md#canonical-dab606ca3cfc914b5b298827e7d4769625a28ba7a294c831f9d7c78b6cf7e962) |
| `detection_settings.violations_view.name` | [detection_settings.violations_view.name](data-sources--app_firewall--reference--group-001.md#canonical-eadf44bdb18fb8a5603c3a6940db5503ac54e4b93b431143bc631c9da94e6a03) |
| `detection_settings.violations_view.title` | [detection_settings.violations_view.title](data-sources--app_firewall--reference--group-001.md#canonical-f67751e868a154ae6ead434bbcc46ba211164428c0fdc8bc0b7dfa8f2b41b2a3) |
| `disable_ai_enhancements` | [disable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-b013e1fa336f759368c1763e2e71c77f78e2297ea9a8443d7a594ce1fa8cef8a) |
| `disable_anonymization` | [disable_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-078327a6c8e4882459102e191b325bfd490e7adfb21f1bf0bd96f6baf399c0a8) |
| `enable_ai_enhancements` | [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-a8799cac70fd7916220a34f1a08ca9ebe81ab22890cc715c33a811727229c0e7) |
| `enable_ai_enhancements.mitigate_high_medium_risk_action` | [enable_ai_enhancements.mitigate_high_medium_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-6f1be2a13759d970c93f905046ab1d7913286d9c68788f62f39cea6cd5b9ff11) |
| `enable_ai_enhancements.mitigate_high_risk_action` | [enable_ai_enhancements.mitigate_high_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-d3fbc5deeeafff8a77a82ca06535339612cb655e05e900125deb268ac2e31f4e) |
| `id` | [id](data-sources--app_firewall--reference--group-001.md#canonical-0a271e0fe9a10068f165379aaf01d931b29716e498eed68131703d467e3d674a) |
| `labels` | [labels](data-sources--app_firewall--reference--group-001.md#canonical-4d0b40c416d0f2030b49f24ac51a4491aa3ad0282aa49c9b81b3dcf9a6c56226) |
| `monitoring` | [monitoring](data-sources--app_firewall--reference--group-001.md#canonical-f0a4fcc34a2ee4417ff037db52f9184b895611e2a88c9f7cbb0b2780ed8400f7) |
| `name` | [name](data-sources--app_firewall--reference--group-001.md#canonical-4b18fb1f9af03c508816e454566ed3b2e38456c3854eabf357f0ad39950514ca) |
| `namespace` | [namespace](data-sources--app_firewall--reference--group-001.md#canonical-4fa7323cbbd078d7297fbc9b1b36218729655af9601d59c5b39ee1c768578539) |
| `use_default_blocking_page` | [use_default_blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-396dd05afe9fdf2565125ad9bd2c54a87f5d956d26d832c0d873b3011c86ef43) |

<a id="canonical-b36cfd0aee2c4b2c3fc9aba3a46178d273dd11a673d95f216f482a5e85801b13"></a>

## Next pages — Property reference / 619f3a4a423a / 11

- [allow_all_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-b7077d063193988d29112a0b46c27a029b4a9acd421fc1c457024be9954eaa07)
- [allowed_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-982f10529300907bec5f5acbe8468680b0535a0099774e109cd92de6a035a752)
- [blocking](data-sources--app_firewall--reference--group-001.md#canonical-bcd3456b05507a43dc7ef59a1e7e2d4ff4371f0bbdaa7cfed1ec234776ad8f87)
- [blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-b8edf21322eff78b2dfdd77ce400c552507d006ee15b9411f9150f91ef30d34d)
- [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-932a9382cdb86d30a1222ced128b734f897e01133cb737083c962a94c98ed0e8)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945)
- [default_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-6cfe415781aec97010e0a7f2d9be34b5aef8d17eca28633e075c603f119e3a28)
- [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-62357eb47ae3cebef5f2aaeed586e8c444e2ee73e66cdcaadccbc0f3921f63e5)
- [default_detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-5cefd983e2bbcfaa75926052f8ab411809c5ecf04d9124b0239c355e7996ea59)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [disable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-44e035f938c10b60f96cf72524a61411b2d81a63d6ed271c32afa1919755c343)
- [disable_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-f3e7b68ac4045b02148757c9c852efeac2f9b18e7778c684d1a47c8d6e8412ba)
- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-fcec69530d75cc1a8525164bd619b25bedc1adfd5198f98ef2eaf9aee0e68ae5)
- [monitoring](data-sources--app_firewall--reference--group-001.md#canonical-1ef81513048703b410f3651ef1afd3b69d8bd753f8118b1431462ebc01179857)
- [use_default_blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-9f246fc996a562416e707b574f84dce284240997ed47a6fd1bbf0bb3f4725ffc)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-b7077d063193988d29112a0b46c27a029b4a9acd421fc1c457024be9954eaa07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b3e6aa35a5415d1cf8c40865930d80572f6f752b6374e5bf8bc114710bcd598"></a>

## allow_all_response_codes — allow_all_response_codes / 1b6c54cecd64 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- allow_all_response_codes

<a id="canonical-8c62e7a01e8d767cb883b6479d0a14d289caf226853fa9d708add1cadb4e207b"></a>

Type: `["object", {}]`. Computed.

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

- [allow_all_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-8c62e7a01e8d767cb883b6479d0a14d289caf226853fa9d708add1cadb4e207b)
- [allowed_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-ba093a9546347115c5bbc38383fdee7a4e6e125879db70404082e77726d85911)

Select alternatives according to the provider validators above.

<a id="canonical-16b8b55c888ebeea45d28468b1ff833b0a4cf406e44928ad52df1f3bc1758448"></a>

## Direct properties — allow_all_response_codes / 1b6c54cecd64 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2dc25bc75276f0addb2fc90d314a3d22c65f950de8914178932f74b8dd49e194"></a>

## Next pages — allow_all_response_codes / 1b6c54cecd64 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-982f10529300907bec5f5acbe8468680b0535a0099774e109cd92de6a035a752"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a23136cc546cbd988509e85edc244a334249d7d04007fee24affda501a1d2907"></a>

## allowed_response_codes — allowed_response_codes / 060b40913785 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- allowed_response_codes

<a id="canonical-ba093a9546347115c5bbc38383fdee7a4e6e125879db70404082e77726d85911"></a>

Type: `"single"`. Computed.

List of HTTP response status codes that are allowed.

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

<a id="canonical-7ed221f667a95f9a92a8d8d332156c5796b998bbb90542fbd598695305d3381c"></a>

## Direct properties — allowed_response_codes / 060b40913785 / 3

<a id="canonical-382b3ea3af303c6fe2d7f9d2a62594b7690a7df8860b2b01800ef0146c3d83e2"></a>

<a id="canonical-399b0684c8d3ea9d0c1286dd2615779b7fcc1c9e9b7fad7af02462d89d298d70"></a>

## response_code property — allowed_response_codes / 060b40913785 / 4

Type: `["list", "number"]`. Computed.

List of HTTP response status codes that are allowed.

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

<a id="canonical-26e4868eb8e393a547745d172450fe6824ffeb57cf283846bf1ffb25761e2a1e"></a>

## Next pages — allowed_response_codes / 060b40913785 / 5

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-bcd3456b05507a43dc7ef59a1e7e2d4ff4371f0bbdaa7cfed1ec234776ad8f87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6253c660478f4c8773eaf79301c350f7c872c7d6545933e888c3b2bea7fb74a6"></a>

## blocking — blocking / f2d46d773723 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- blocking

<a id="canonical-3009f5e7a3924817b4a9e0c3b6ce39e6ace5a7ca3dc8660dade064c6ead87cf2"></a>

Type: `["object", {}]`. Computed.

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

- [blocking](data-sources--app_firewall--reference--group-001.md#canonical-3009f5e7a3924817b4a9e0c3b6ce39e6ace5a7ca3dc8660dade064c6ead87cf2)
- [monitoring](data-sources--app_firewall--reference--group-001.md#canonical-f0a4fcc34a2ee4417ff037db52f9184b895611e2a88c9f7cbb0b2780ed8400f7)

Select alternatives according to the provider validators above.

<a id="canonical-b7d8388050890685f4cf687e47d2605d255d893429cfb8fbc82894ffcb0affa0"></a>

## Direct properties — blocking / f2d46d773723 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0dd5cd939d4d1a9c27705de9e39aa831d55bb1dd49ad1449791c3567f8922a5c"></a>

## Next pages — blocking / f2d46d773723 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-b8edf21322eff78b2dfdd77ce400c552507d006ee15b9411f9150f91ef30d34d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-119976dd259855aaf318e888f80e47886daddffb919fdcff19ce144bad4c08cb"></a>

## blocking_page — blocking_page / 7c82c6463f77 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- blocking_page

<a id="canonical-2b7d5052f77b4cc0ffd886feed4400fcab25452d97e22256d122fc3b9afc3209"></a>

Type: `"single"`. Computed.

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

- [blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-2b7d5052f77b4cc0ffd886feed4400fcab25452d97e22256d122fc3b9afc3209)
- [use_default_blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-396dd05afe9fdf2565125ad9bd2c54a87f5d956d26d832c0d873b3011c86ef43)

Select alternatives according to the provider validators above.

<a id="canonical-88b4400a7f89d461f8bff97bb34762db6095f894429a9508e77be1f26941d886"></a>

## Direct properties — blocking_page / 7c82c6463f77 / 3

<a id="canonical-03ef2e4ebc6b385a6f4d82998c2aa317565c1324c0275e74fe892838beab614e"></a>

<a id="canonical-56ea6fe1e0c092df00c5c0c0fc50d9aeddadaf62e75e87430702c713ac0f8452"></a>

## blocking_page property — blocking_page / 7c82c6463f77 / 4

Type: `"string"`. Computed.

Define the content of the response page (e.g., an HTML document or a JSON object), use the
\{\{request\_id\}\} placeholder to provide users with a unique identifier to be able to trace the
blocked request in the logs. The maximum allowed size of response body is 4096 bytes after base64
encoding..

Upstream description:

Define the content of the response page (e.g., an HTML document or a JSON object), use the
\{\{request\_id\}\} placeholder to provide users with a unique identifier to be able to trace the
blocked request in the logs. The maximum allowed size of response body is 4096 bytes after base64
encoding, which would be about 3070 bytes in plain text.

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

<a id="canonical-3d36093bb4d43ee63a22d908b6dc38de24cda3c4183182e399eb2b953e40064a"></a>

<a id="canonical-34641c654ded40d1e5e99c09eff3cf38c4e21777e695e7a5a7a8895187de6d34"></a>

## response_code property — blocking_page / 7c82c6463f77 / 5

Type: `"string"`. Computed.

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

<a id="canonical-c1ebdb81a94d7944a0e876c52db9374e4c115e8b538f1a7d246f94299ad06448"></a>

## Next pages — blocking_page / 7c82c6463f77 / 6

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-932a9382cdb86d30a1222ced128b734f897e01133cb737083c962a94c98ed0e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d6ad42725d1d74854b3e5325d8287ab96596c7078a3a83e2bd6b99721ab115a"></a>

## bot_protection_setting — bot_protection_setting / ac24493d00cb / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- bot_protection_setting

<a id="canonical-c2b6b993051856063a601a77254cd927b4c28b3204c2ca81ebb28505c4f385af"></a>

Type: `"single"`. Computed.

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

- [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-c2b6b993051856063a601a77254cd927b4c28b3204c2ca81ebb28505c4f385af)
- [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-4e9cb75daf676a9c58bfb9dd87aa25d3fe4a0e907e348a949e73036d5664809c)

Select alternatives according to the provider validators above.

<a id="canonical-8fa05436a5b66cf7d54d223db0824d2d4b23bd4b5a87a555b7eb4de1fd0f90b6"></a>

## Direct properties — bot_protection_setting / ac24493d00cb / 3

<a id="canonical-71616fd4209e5c2998ba8184183a0052959bba6829fdeea071c5ba5cdd318867"></a>

<a id="canonical-aedbe0400f89937f2566c5827e4d2167e5e7db69941120792e42209a2384adc5"></a>

## good_bot_action property — bot_protection_setting / ac24493d00cb / 4

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-44a058a27c966252960b7f85e80a202b329bdcfd34f4becf9ca664f76a8d9597"></a>

<a id="canonical-588f4003b77f93d31ace4966b0aa1dfb9b67c2430eb1891143e58a1e9e532204"></a>

## malicious_bot_action property — bot_protection_setting / ac24493d00cb / 5

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-084adfb8118161c32514e5ae5869da27cde974ece785c76223522242195c8404"></a>

<a id="canonical-fbc7480b840f996b7fc639e9dbbfefe72035f9936bbd657a9fa349d22d922fa6"></a>

## suspicious_bot_action property — bot_protection_setting / ac24493d00cb / 6

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-d6522149c4cb261830be79f09c5fdad31af8e5ca198955bcb20d735ddb4ca6d3"></a>

## Next pages — bot_protection_setting / ac24493d00cb / 7

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4197c6ae29796939fa14db396f12f4b2cc79c5adedfef37d87be5f6f51402c01"></a>

## custom_anonymization — custom_anonymization / cce1b8a8b94e / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- custom_anonymization

<a id="canonical-8647e57eefa234b828f17d38b943e9b7c283570a24058676bb2ecd9d75ba2fad"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_anonymization, default\_anonymization, disable\_anonymization; Default:
default\_anonymization\] Anonymization settings which is a list of HTTP headers, parameters and
cookies.

Upstream description:

Anonymization settings which is a list of HTTP headers, parameters and cookies.

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

- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-8647e57eefa234b828f17d38b943e9b7c283570a24058676bb2ecd9d75ba2fad)
- [default_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-cfd9ce7185449c7a623412d0d0814cdb98eb920c24de7060b36117ebbbbea32b)
- [disable_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-078327a6c8e4882459102e191b325bfd490e7adfb21f1bf0bd96f6baf399c0a8)

Select alternatives according to the provider validators above.

<a id="canonical-01b2f050838e1e3bc43c31b431d4f7d5fa20ca315c1a4f1c7480379857d0593c"></a>

## Direct properties — custom_anonymization / cce1b8a8b94e / 3

- [anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f): complete subsection reference.

<a id="canonical-3178aabfcb257c0a96ba7b41c3b31b1bebb46689f7ad16363cd9b1c807a4bc5a"></a>

## Next pages — custom_anonymization / cce1b8a8b94e / 4

- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89f7cd8149cf8cab17e9f975c78c13859cb22b9b3b7a3c698b1058663e7e2ac6"></a>

## custom_anonymization.anonymization_config — custom_anonymization.anonymization_config / 0b4b34664834 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945)
- custom_anonymization.anonymization_config

<a id="canonical-088777cee70e830bfd28ee1eacee3f7f58fd1eea56fe927a01b483966071d5cd"></a>

Type: `"list"`. Computed.

List of HTTP headers, cookies and query parameters whose values will be masked.

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

<a id="canonical-f8e77522a13c74565aa5d121807ea8e690732ddc4c24737c00c3d312252174a1"></a>

## Direct properties — custom_anonymization.anonymization_config / 0b4b34664834 / 3

- [cookie](data-sources--app_firewall--reference--group-001.md#canonical-02359bcffd9de07a1d04a58b9dfe0974d41dc05bdb52d069d7a42bb5a2aee165): complete subsection reference.

- [http_header](data-sources--app_firewall--reference--group-001.md#canonical-10e6a15abd181e5ca18170bf6e52b112eeb2f040d02fb150ae0678b64c38d788): complete subsection reference.

- [query_parameter](data-sources--app_firewall--reference--group-001.md#canonical-ebe75b05f230925f641abef59603af2c7daf037a587021b9e383b6d4ea6025f2): complete subsection reference.

<a id="canonical-aa065fe30c7f79cb2df480bcd2ebae95c8bdf676fd775893a83fc8b6a631acf4"></a>

## Next pages — custom_anonymization.anonymization_config / 0b4b34664834 / 4

- [custom_anonymization.anonymization_config.cookie](data-sources--app_firewall--reference--group-001.md#canonical-02359bcffd9de07a1d04a58b9dfe0974d41dc05bdb52d069d7a42bb5a2aee165)
- [custom_anonymization.anonymization_config.http_header](data-sources--app_firewall--reference--group-001.md#canonical-10e6a15abd181e5ca18170bf6e52b112eeb2f040d02fb150ae0678b64c38d788)
- [custom_anonymization.anonymization_config.query_parameter](data-sources--app_firewall--reference--group-001.md#canonical-ebe75b05f230925f641abef59603af2c7daf037a587021b9e383b6d4ea6025f2)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-02359bcffd9de07a1d04a58b9dfe0974d41dc05bdb52d069d7a42bb5a2aee165"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9af56875ee961e2e8809efb8ac5a901a3b5acbe9a3859cf6fb1e57ac563b4b01"></a>

## custom_anonymization.anonymization_config.cookie — custom_anonymization.anonymization_config.cookie / a71b362ec8af / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945)
- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f)
- custom_anonymization.anonymization_config.cookie

<a id="canonical-30eb906df4730a117b6265640c9ade52040d66731cb9b4a9567247b98d7e2399"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Cookies.

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

<a id="canonical-285aa8ad59cc2abeb0ca50adf3cb9d84b6202fbb54ccfbe1679293b08ce1bc28"></a>

## Direct properties — custom_anonymization.anonymization_config.cookie / a71b362ec8af / 3

<a id="canonical-6cd9bc5d2848fef28dab064deb0cd96d6da31c49673067de15388a35d719c119"></a>

<a id="canonical-02449735f114c45b9e199050581db5c8237b5186a1a7b50107ea89f4ab3105bd"></a>

## cookie_name property — custom_anonymization.anonymization_config.cookie / a71b362ec8af / 4

Type: `"string"`. Computed.

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

Upstream description:

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

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

<a id="canonical-7cb93dafb8d2b62d9639a2da9605454c6dbd50c437606716fc3d0b8b337f2f21"></a>

## Next pages — custom_anonymization.anonymization_config.cookie / a71b362ec8af / 5

- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-10e6a15abd181e5ca18170bf6e52b112eeb2f040d02fb150ae0678b64c38d788"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18ba977ba469cd121a4f97697e0118c5b27df3a17d097537320972cfdd3537fb"></a>

## custom_anonymization.anonymization_config.http_header — custom_anonymization.anonymization_config.http_header / c9a1801e0e77 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945)
- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f)
- custom_anonymization.anonymization_config.http_header

<a id="canonical-109fa2fa3ad563ffe761803f1a0ceb50c93f8c7c98696fe064ffbb458aee0158"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Headers.

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

<a id="canonical-1893ced8a5ed6fb5b8f10448bd837aa1e0dcd3f9d96e9c4c233b12373421f00c"></a>

## Direct properties — custom_anonymization.anonymization_config.http_header / c9a1801e0e77 / 3

<a id="canonical-16f0eca78c4630b3ad1326ae38148695a01efe3ceba0968d48bfd40699896e96"></a>

<a id="canonical-7a37ae4e1e0f8407136b3a7b4161d32fd7e2d8cb2a6459df4dded1db281c4128"></a>

## header_name property — custom_anonymization.anonymization_config.http_header / c9a1801e0e77 / 4

Type: `"string"`. Computed.

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

<a id="canonical-963c127680d4f79ac12aaa6d6d2de8b8af2d5a1085e2589beda98853797ece92"></a>

## Next pages — custom_anonymization.anonymization_config.http_header / c9a1801e0e77 / 5

- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-ebe75b05f230925f641abef59603af2c7daf037a587021b9e383b6d4ea6025f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33d777e820c21ae9a411b091f5093fef17e9e34b254cceada40952a456bdd952"></a>

## custom_anonymization.anonymization_config.query_parameter — custom_anonymization.anonymization_config.query_parameter / 699391063b71 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0b22a1d228a6f45df42c85e2a3780c2df4815c1a6abefdf31d1abd2e657a7945)
- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f)
- custom_anonymization.anonymization_config.query_parameter

<a id="canonical-c88ab4f3a2897dde658909b6792f77efcd32cbdb8fb4426d8c2981a9f18ab197"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Parameters.

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

<a id="canonical-e9dbb3b041648907d65e5ddca6d5ce42b21279ad804ffe4459d33dbb913c190a"></a>

## Direct properties — custom_anonymization.anonymization_config.query_parameter / 699391063b71 / 3

<a id="canonical-7aae01a0630a2af52f740c23d738c616c9872d291719a3685a1eac1d11547065"></a>

<a id="canonical-4f0dc068eccb00ec89d1167dbf80fb1a01a68b5da60199638149585b0f4f7207"></a>

## query_param_name property — custom_anonymization.anonymization_config.query_parameter / 699391063b71 / 4

Type: `"string"`. Computed.

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Upstream description:

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

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

<a id="canonical-1f7e8b6f0ce00ec1f6cb0efc54ee63578ed130bda613db4dbdfd24cb2099032d"></a>

## Next pages — custom_anonymization.anonymization_config.query_parameter / 699391063b71 / 5

- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-f74cb506d66f6936bfe8e4330237f0f93c1283fcaed853ac02dabc27c4520a2f)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-6cfe415781aec97010e0a7f2d9be34b5aef8d17eca28633e075c603f119e3a28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c5bd7c597b5a0baf0011304d64f44f1313340c9ae9df3010638e5f5721a4dd0"></a>

## default_anonymization — default_anonymization / 58c55e37a3f8 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- default_anonymization

<a id="canonical-cfd9ce7185449c7a623412d0d0814cdb98eb920c24de7060b36117ebbbbea32b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-ba0c8226d25e0f677a2b9c7d00d2f62bc1bf6267554ec6ddea41fc376c6a8913"></a>

## Direct properties — default_anonymization / 58c55e37a3f8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82dab7baa801223625c9e1853ccdae6c1840b978e6f77583ec85e39eb683194c"></a>

## Next pages — default_anonymization / 58c55e37a3f8 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-62357eb47ae3cebef5f2aaeed586e8c444e2ee73e66cdcaadccbc0f3921f63e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2586fd86576d690f65db09438a77b80d56e701ba830047210a7529d080c71eb5"></a>

## default_bot_setting — default_bot_setting / db655ca1f2c6 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- default_bot_setting

<a id="canonical-4e9cb75daf676a9c58bfb9dd87aa25d3fe4a0e907e348a949e73036d5664809c"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-337cab6bf3d02c7b127120a27cc62e8f6c2320fd633297986c9f6cd7c9166c48"></a>

## Direct properties — default_bot_setting / db655ca1f2c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31ebca6883c8b0177ba743a501b4f7a8e6ba2516ab4571a53063b2e9ab847090"></a>

## Next pages — default_bot_setting / db655ca1f2c6 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-5cefd983e2bbcfaa75926052f8ab411809c5ecf04d9124b0239c355e7996ea59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2e720721f1d8b32ec2ac466dd430b7a77b7bd146b09025c8603f19f4c836910"></a>

## default_detection_settings — default_detection_settings / be39300b0ec8 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- default_detection_settings

<a id="canonical-005dd7654e80c77b9ce466e5862b7ca2b78aef9318ce17b0c6a77ce0f39a9731"></a>

Type: `["object", {}]`. Computed.

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

- [default_detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-005dd7654e80c77b9ce466e5862b7ca2b78aef9318ce17b0c6a77ce0f39a9731)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-092c60de7d22883fab85adabaa2f3df8aca6fbfc8d195c539d4ceb638b3e6586)

Select alternatives according to the provider validators above.

<a id="canonical-043e0889d4c8e3282e093f244b241e87477516e40bcfad90f6a682ede5c4a1c5"></a>

## Direct properties — default_detection_settings / be39300b0ec8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c72337a563f0cf5e138788d262e1f8d23c2114955d65eebcabaef2067f0ac95e"></a>

## Next pages — default_detection_settings / be39300b0ec8 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab059de303b886d1aedb86b532aa7ab2c459a2015b606ff9d0955f482d6880b3"></a>

## detection_settings — detection_settings / f7f2c29bdf92 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- detection_settings

<a id="canonical-092c60de7d22883fab85adabaa2f3df8aca6fbfc8d195c539d4ceb638b3e6586"></a>

Type: `"single"`. Computed.

Specifies detection settings to be used by WAF.

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

<a id="canonical-30e5d39abfec899b5d697514c82aa6be93e3fcb845953748ac6fff7245e89264"></a>

## Direct properties — detection_settings / f7f2c29bdf92 / 3

- [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-d8bc851f02e353f77014cff3bb5e6c236b376240d6acda051a24be7d1da9becb): complete subsection reference.

- [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-1efc721c522706b4f9744818f3ac0030014c39cd57a54a2d7bc1f5637608eac9): complete subsection reference.

- [default_violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-09a87f2f24c9c76bce521860c6c95e70c0ffc995304975191d90161dc15795f2): complete subsection reference.

- [disable_staging](data-sources--app_firewall--reference--group-001.md#canonical-e76b36c5971befafe3a6aea2f1056a10bf3f41eacf284488a496705f244eea98): complete subsection reference.

- [disable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-26d638764a488f7871aa913db6147afd9681f852e6afa8a50d67e58a84123860): complete subsection reference.

- [disable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-88e02403676a1f1ddfd8d916b24f101c7812206a56e97b2f047cc57afd72b47e): complete subsection reference.

- [enable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-1b824099e1438bae6b26fcc515a5d96e04a90868e6be4afdbc92d81b7ee37ca9): complete subsection reference.

- [enable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-71c80a7d4c30cbe6d0957be2f6c2739c470da4de53927b46469e691e259aefd4): complete subsection reference.

- [signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec): complete subsection reference.

- [stage_new_and_updated_signatures](data-sources--app_firewall--reference--group-001.md#canonical-ee826af0f82d4fb0bde2520d8c49c0282b585d3961dbd785e570706f14f21f85): complete subsection reference.

- [stage_new_signatures](data-sources--app_firewall--reference--group-001.md#canonical-11cd9cbaab4552932b57a4e1bdc033a604ec96b00c1df82cab84861f1967e993): complete subsection reference.

- [violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-59b777e240fa8bb798a751970904e94a184807cdaf91d7c8a1455b69fe20678d): complete subsection reference.

- [violations_view](data-sources--app_firewall--reference--group-001.md#canonical-d05829dd59400aa86a661ccecf1b306a92f35d1c494f1096291e3d10ba680875): complete subsection reference.

<a id="canonical-3249da3eb8f693a0b4eaee00048271fe2e3c43746dd1fa365ff818d4cd65f53e"></a>

## Next pages — detection_settings / f7f2c29bdf92 / 4

- [detection_settings.bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-d8bc851f02e353f77014cff3bb5e6c236b376240d6acda051a24be7d1da9becb)
- [detection_settings.default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-1efc721c522706b4f9744818f3ac0030014c39cd57a54a2d7bc1f5637608eac9)
- [detection_settings.default_violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-09a87f2f24c9c76bce521860c6c95e70c0ffc995304975191d90161dc15795f2)
- [detection_settings.disable_staging](data-sources--app_firewall--reference--group-001.md#canonical-e76b36c5971befafe3a6aea2f1056a10bf3f41eacf284488a496705f244eea98)
- [detection_settings.disable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-26d638764a488f7871aa913db6147afd9681f852e6afa8a50d67e58a84123860)
- [detection_settings.disable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-88e02403676a1f1ddfd8d916b24f101c7812206a56e97b2f047cc57afd72b47e)
- [detection_settings.enable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-1b824099e1438bae6b26fcc515a5d96e04a90868e6be4afdbc92d81b7ee37ca9)
- [detection_settings.enable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-71c80a7d4c30cbe6d0957be2f6c2739c470da4de53927b46469e691e259aefd4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [detection_settings.stage_new_and_updated_signatures](data-sources--app_firewall--reference--group-001.md#canonical-ee826af0f82d4fb0bde2520d8c49c0282b585d3961dbd785e570706f14f21f85)
- [detection_settings.stage_new_signatures](data-sources--app_firewall--reference--group-001.md#canonical-11cd9cbaab4552932b57a4e1bdc033a604ec96b00c1df82cab84861f1967e993)
- [detection_settings.violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-59b777e240fa8bb798a751970904e94a184807cdaf91d7c8a1455b69fe20678d)
- [detection_settings.violations_view](data-sources--app_firewall--reference--group-001.md#canonical-d05829dd59400aa86a661ccecf1b306a92f35d1c494f1096291e3d10ba680875)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-d8bc851f02e353f77014cff3bb5e6c236b376240d6acda051a24be7d1da9becb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07a0f50f002ce0a2fee33a6593f69cbe221504ffdd3443a1f097093ddd95bf92"></a>

## detection_settings.bot_protection_setting — detection_settings.bot_protection_setting / f3bd1d4ecc9c / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.bot_protection_setting

<a id="canonical-d4a109492762005d822826e1677e8197492b62ba439699fa9dabb52fa48a45d1"></a>

Type: `"single"`. Computed.

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

<a id="canonical-cf1a3dc90cada553e4d6fcc6057216ec9d6770d8e78cbd846c0da494388b0871"></a>

## Direct properties — detection_settings.bot_protection_setting / f3bd1d4ecc9c / 3

<a id="canonical-992fe12389781d769a98beb3097e74594a6c63cd28f816debe143c0483dcb34c"></a>

<a id="canonical-c4df0c8532f55b46a2d6a550f0cebdb6303e46477217ff134a507f1d11da3b95"></a>

## good_bot_action property — detection_settings.bot_protection_setting / f3bd1d4ecc9c / 4

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-856c95edcbfc0bf54281f43780e41690f7e4b76da6a7e3f7f05f41b3ba7d379c"></a>

<a id="canonical-9447de3e994ff37b3cb822f1703d4d33bcf45c4475a58b77cb215aba15ced508"></a>

## malicious_bot_action property — detection_settings.bot_protection_setting / f3bd1d4ecc9c / 5

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-f05fa4855f56f63603ccf693c8e076e3bd9f3e1df36b725e7d11b5282afd3c88"></a>

<a id="canonical-05cec9c3af581ae510748c86d805c1ae61feb456a722f5d1309837c405891e1d"></a>

## suspicious_bot_action property — detection_settings.bot_protection_setting / f3bd1d4ecc9c / 6

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-2b8ac952be9744780dabe3276e9e908d5f207a0c92105c28b707930a2f96dc26"></a>

## Next pages — detection_settings.bot_protection_setting / f3bd1d4ecc9c / 7

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-1efc721c522706b4f9744818f3ac0030014c39cd57a54a2d7bc1f5637608eac9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa94a54132ad2e45388a4d15bcf56ecf4cd39596bddabf6d2332263fc470f4e6"></a>

## detection_settings.default_bot_setting — detection_settings.default_bot_setting / 66ff3545861a / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.default_bot_setting

<a id="canonical-61a532cc5d14566888c0e162f30c7c60582f5ace46832f597826593437ff06f9"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3407b1aa0a094de23141dd50500d6480bee04a2dc89557d8eb6e5f59c4b5b439"></a>

## Direct properties — detection_settings.default_bot_setting / 66ff3545861a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c462825e05dbd6e997e4c2e076b85687a76975d4b957fcb28f22ec2c63167bff"></a>

## Next pages — detection_settings.default_bot_setting / 66ff3545861a / 4

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-09a87f2f24c9c76bce521860c6c95e70c0ffc995304975191d90161dc15795f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6bf5555bcb140ad02df813c1b7dc12352afc3ef69058ef142950019f8592b7c"></a>

## detection_settings.default_violation_settings — detection_settings.default_violation_settings / 820e227ed498 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.default_violation_settings

<a id="canonical-a6ab11418a639f52f644b82c828315ce37458d7759936b19904c590e555f3fe1"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-7bc0519e932a05f20854e376a55a76e31e32fa69dbaea222a5b0f0f88e05146b"></a>

## Direct properties — detection_settings.default_violation_settings / 820e227ed498 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1925035fbb51c510da06c0167ec8bcd7ceacde2f6202ea704828874d254c69c"></a>

## Next pages — detection_settings.default_violation_settings / 820e227ed498 / 4

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-e76b36c5971befafe3a6aea2f1056a10bf3f41eacf284488a496705f244eea98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07299a6397193e2f40bc093d1fefc1c7e4377be30296f753067fc645ce49d170"></a>

## detection_settings.disable_staging — detection_settings.disable_staging / 700ca1510bcd / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.disable_staging

<a id="canonical-14e4c6e8c649312be1e6ef14cffd558c209310273e0654c6e417448c9f074e84"></a>

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

<a id="canonical-dd1e48b9d4454e2f69631e800204dd87fd516a8fc59e04f40a7f88bf65be023d"></a>

## Direct properties — detection_settings.disable_staging / 700ca1510bcd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-988b627bdcf922de925e593f7a00531bf1544cbbb127208600b9b79a0e42dcd4"></a>

## Next pages — detection_settings.disable_staging / 700ca1510bcd / 4

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-26d638764a488f7871aa913db6147afd9681f852e6afa8a50d67e58a84123860"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4470166c00d0ed5a5844a4141e1da1b0d26e3abbb68bfc23dd0e6b52f0f967f6"></a>

## detection_settings.disable_suppression — detection_settings.disable_suppression / a1f63fed4145 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.disable_suppression

<a id="canonical-9f06702df29c70bb711dbb06a4e94b0911bd26427323fb44ac60f533bb7f4198"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-bfd38f7c2b4daf160c9bc4f81f2a120314149fa121d7ebebfdd74b04c33a242c"></a>

## Direct properties — detection_settings.disable_suppression / a1f63fed4145 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e17686f5bf9bf66924dc75110078645a20396b7b5f04935315b63870946c534"></a>

## Next pages — detection_settings.disable_suppression / a1f63fed4145 / 4

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-88e02403676a1f1ddfd8d916b24f101c7812206a56e97b2f047cc57afd72b47e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3eb654e986da5411df90049b68fbc3d21aec036db64ab9126decaf0c1b10a31"></a>

## detection_settings.disable_threat_campaigns — detection_settings.disable_threat_campaigns / eff827f6b8c1 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.disable_threat_campaigns

<a id="canonical-b2329bb5d06c95129c948ca88a91ec0b4ac375507c29f22ced68304459a64406"></a>

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

<a id="canonical-63773aba882ed141e2fdede0971761b8cde2d87e529fd2da2ae626e526b52c79"></a>

## Direct properties — detection_settings.disable_threat_campaigns / eff827f6b8c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-13a5851bf0f96dddfe20c64932e44e0e935146317883d2d3e7b55c4629029ce3"></a>

## Next pages — detection_settings.disable_threat_campaigns / eff827f6b8c1 / 4

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-1b824099e1438bae6b26fcc515a5d96e04a90868e6be4afdbc92d81b7ee37ca9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14e2a94bcb6c2b9632c0f14d27088c2890a3994c58e9d467fb3d28ad706eedf9"></a>

## detection_settings.enable_suppression — detection_settings.enable_suppression / 70290bb67f8e / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.enable_suppression

<a id="canonical-f73f8721eb191ec129787c5299a9ef87124b0f12d58cd61dfd6b10537798d9ec"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-8e8225a73812ebae2bcb3645f1dde27c4949b268283087735f0191bbccbf6fb0"></a>

## Direct properties — detection_settings.enable_suppression / 70290bb67f8e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f1278ec0fe7f6a6fc705c4619151ead5e22726e483fd51fb641ba88dc8351a0"></a>

## Next pages — detection_settings.enable_suppression / 70290bb67f8e / 4

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-71c80a7d4c30cbe6d0957be2f6c2739c470da4de53927b46469e691e259aefd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b81de12f113faf00b335145ad504c821661cce743fed70cb30cd0ea00f0f19e1"></a>

## detection_settings.enable_threat_campaigns — detection_settings.enable_threat_campaigns / 844a1b69e73a / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.enable_threat_campaigns

<a id="canonical-c4dc94aad78f38c54a75c6ebd0f792e441e80da38955bcb67bcc72e026edad6a"></a>

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

<a id="canonical-d0627c866b5c4e850c3f85017c2f13ad9ea696acd4aff0bfdc793749e14557ff"></a>

## Direct properties — detection_settings.enable_threat_campaigns / 844a1b69e73a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11127ac7100025eab4e20deb96704bd47dc1d3f36f190c8a4d58c13b682307fc"></a>

## Next pages — detection_settings.enable_threat_campaigns / 844a1b69e73a / 4

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5db71885311207b4badee1a4b6111bc39a6d5dc67c6ec8eae041977492688ab9"></a>

## detection_settings.signature_selection_setting — detection_settings.signature_selection_setting / cff0ff2b0700 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.signature_selection_setting

<a id="canonical-546710318a465f09ecd0bad573b3e882f78deb4b0e112474d4abd6c0a2ee70a5"></a>

Type: `"single"`. Computed.

Attack Signatures are patterns that identify attacks on a web application and its components.

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

<a id="canonical-557e5a25e403f1e966dd0d5d56e81e00687497e35e20baab0f8f0cd36733edfe"></a>

## Direct properties — detection_settings.signature_selection_setting / cff0ff2b0700 / 3

- [attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-e98da775d056c075649dbcef34aa7f3f75b82f972b9778d4d7cca60df3d0178b): complete subsection reference.

- [default_attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-3c62aeaa359e50cfc7757ca87378c42c6469422f47f7b49aa5825550efd3b4da): complete subsection reference.

- [default_signature_setting](data-sources--app_firewall--reference--group-001.md#canonical-2c5f530a51f77a8247c9cb9ece4ee1906741208195819088e1c5f8e25370e4b1): complete subsection reference.

- [high_medium_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-211d3681f10883eac5ad93b8628aa9b40fcd82f8002eee4270ec9cf8652904d2): complete subsection reference.

- [high_medium_low_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-fce5cb603b7a7c23e98f57ce5a8a4d8fb888afeef34b2005e9e8654066e2cf2f): complete subsection reference.

- [only_high_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-20cc71434616842282c671ecb22bcc507f994acc2958d8f35cbfd656125ec2ca): complete subsection reference.

- [signature_settings_by_accuracy](data-sources--app_firewall--reference--group-001.md#canonical-2bd07ea4216146b2e4502921f872c58ec5c7c42a945c1ae01ffc62e989884f4e): complete subsection reference.

<a id="canonical-8e33b03acad73c11dfc8b8935fc93dfef4756b52feca0eecb40e89e5db4b4911"></a>

## Next pages — detection_settings.signature_selection_setting / cff0ff2b0700 / 4

- [detection_settings.signature_selection_setting.attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-e98da775d056c075649dbcef34aa7f3f75b82f972b9778d4d7cca60df3d0178b)
- [detection_settings.signature_selection_setting.default_attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-3c62aeaa359e50cfc7757ca87378c42c6469422f47f7b49aa5825550efd3b4da)
- [detection_settings.signature_selection_setting.default_signature_setting](data-sources--app_firewall--reference--group-001.md#canonical-2c5f530a51f77a8247c9cb9ece4ee1906741208195819088e1c5f8e25370e4b1)
- [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-211d3681f10883eac5ad93b8628aa9b40fcd82f8002eee4270ec9cf8652904d2)
- [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-fce5cb603b7a7c23e98f57ce5a8a4d8fb888afeef34b2005e9e8654066e2cf2f)
- [detection_settings.signature_selection_setting.only_high_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-20cc71434616842282c671ecb22bcc507f994acc2958d8f35cbfd656125ec2ca)
- [detection_settings.signature_selection_setting.signature_settings_by_accuracy](data-sources--app_firewall--reference--group-001.md#canonical-2bd07ea4216146b2e4502921f872c58ec5c7c42a945c1ae01ffc62e989884f4e)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-e98da775d056c075649dbcef34aa7f3f75b82f972b9778d4d7cca60df3d0178b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab00c98757408b95d173a9f7cbffbc836bf336b06d4994342412115a1121b44b"></a>

## detection_settings.signature_selection_setting.attack_type_settings — detection_settings.signature_selection_setting.attack_type_settings / 1ab22d753013 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- detection_settings.signature_selection_setting.attack_type_settings

<a id="canonical-fcd9c4d91715eacc45571e17feb2778df0d5eb2720c585566f16bc6ea6def6d2"></a>

Type: `"single"`. Computed.

Specifies attack-type settings to be used by WAF.

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

<a id="canonical-9274505044d6c3c402a58677e1e9829b8424463a982130c96795b4f3fc2a92ee"></a>

## Direct properties — detection_settings.signature_selection_setting.attack_type_settings / 1ab22d753013 / 3

<a id="canonical-89360c68f05bef3cf0d3db23e13d719e6a12c044cfbbb5d26b1627b473309927"></a>

<a id="canonical-6e713e1fe0c1f28c03b03c0a0cd1488c384bea92d9c5b391d9fdc3e9c2257129"></a>

## disabled_attack_types property — detection_settings.signature_selection_setting.attack_type_settings / 1ab22d753013 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-bd4d591f85a180cb735cf6f99b52db52f5e0a97b2af8d720f9d7e310b6ea0b84"></a>

## Next pages — detection_settings.signature_selection_setting.attack_type_settings / 1ab22d753013 / 5

- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-3c62aeaa359e50cfc7757ca87378c42c6469422f47f7b49aa5825550efd3b4da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79eeb04fae74a4f269262fbb872b0d81c02c9fee7858bf34e05b659269f45cab"></a>

## detection_settings.signature_selection_setting.default_attack_type_settings — detection_settings.signature_selection_setting.default_attack_type_settings / e60563f6eab1 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- detection_settings.signature_selection_setting.default_attack_type_settings

<a id="canonical-a74d028baba263c769a907c7471e674bac2a988ac9fabaa7ca4d4cbd21cc159d"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-8b8b0fac30a8572f8c67ccccb24a57eec1b8cdd48efaa7b36c40be58f1964742"></a>

## Direct properties — detection_settings.signature_selection_setting.default_attack_type_settings / e60563f6eab1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fbd8808eeede047f5b548b19eedfae94aee297cf2f032a0face02489fc07d84e"></a>

## Next pages — detection_settings.signature_selection_setting.default_attack_type_settings / e60563f6eab1 / 4

- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-2c5f530a51f77a8247c9cb9ece4ee1906741208195819088e1c5f8e25370e4b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-058e2003d1940c0f1da15cd13b7d64a97ac1457b777bce00f90f03e955863fa8"></a>

## detection_settings.signature_selection_setting.default_signature_setting — detection_settings.signature_selection_setting.default_signature_setting / 0b2ea20ed889 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- detection_settings.signature_selection_setting.default_signature_setting

<a id="canonical-f168ecde6a98d100e49d410cfa4ae3bc16a17a37275cec876bbc309c4c2331eb"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-01d518cd28821d3d5747e69b26b89726713ca134524979b2a5250501cc170b18"></a>

## Direct properties — detection_settings.signature_selection_setting.default_signature_setting / 0b2ea20ed889 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2cc0af06d9bab33d42ed3508f09b25900be4d3453e25eb965f33538ef7a508c1"></a>

## Next pages — detection_settings.signature_selection_setting.default_signature_setting / 0b2ea20ed889 / 4

- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-211d3681f10883eac5ad93b8628aa9b40fcd82f8002eee4270ec9cf8652904d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67283780e14efc5e96b584119d40aeb950621ac35b99c04d468b96eede7721f6"></a>

## detection_settings.signature_selection_setting.high_medium_accuracy_signatures — detection_settings.signature_selection_setting.high_medium_accuracy_signatures / de39e32f553c / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- detection_settings.signature_selection_setting.high_medium_accuracy_signatures

<a id="canonical-be7f45889d5e01a9fb23dba8a937013529f8195d818ed3186dba0bd438e99179"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d2cf0846447a9121304dc65526f5b0c785cd14dfbe043a1cf4406b2d18d61f00"></a>

## Direct properties — detection_settings.signature_selection_setting.high_medium_accuracy_signatures / de39e32f553c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83377f1d76605f6d0ed75403f038a076f04ebaa24976f157761bbef857c7dd4b"></a>

## Next pages — detection_settings.signature_selection_setting.high_medium_accuracy_signatures / de39e32f553c / 4

- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-fce5cb603b7a7c23e98f57ce5a8a4d8fb888afeef34b2005e9e8654066e2cf2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f903537e087e7f1d681047ac84ac8ff4e3e8756fba6524dcfde933352d2e4ea"></a>

## detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures — detection_settings.signature_selection_setting.high_medium_low_accuracy_signatur / f07781becb1b / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures

<a id="canonical-f8698191f8efd7a24fa005cbcb01a1eb8f343031cbf65d4264ce24ae16272c3d"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-6a401699b6e85b14c410ae90ac9fe269e21f146760f2962400f7f49bfec688aa"></a>

## Direct properties — detection_settings.signature_selection_setting.high_medium_low_accuracy_signatur / f07781becb1b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d685e2d328f05b895e05acbdcdd4c1de02fa3205be121ea75ce3f2b20f310753"></a>

## Next pages — detection_settings.signature_selection_setting.high_medium_low_accuracy_signatur / f07781becb1b / 4

- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-20cc71434616842282c671ecb22bcc507f994acc2958d8f35cbfd656125ec2ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a6d8d535bbf3b0fe07bf0470e617bb09b745d1500dd8958918c71134a7c6823"></a>

## detection_settings.signature_selection_setting.only_high_accuracy_signatures — detection_settings.signature_selection_setting.only_high_accuracy_signatures / 5bab593fb910 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- detection_settings.signature_selection_setting.only_high_accuracy_signatures

<a id="canonical-25a4fd6c4318ab011401b8f02780644f9c3722beec905f72e8559cb4acefb5f8"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-9054af6fee807da17dbc11d55a3d82c8fc0ecac13fd9ced547b451cdc101ebae"></a>

## Direct properties — detection_settings.signature_selection_setting.only_high_accuracy_signatures / 5bab593fb910 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62c9391cf67a494cfc2ae22ea8c4bec2e5eb36a71993deb71a74b0db8c83c382"></a>

## Next pages — detection_settings.signature_selection_setting.only_high_accuracy_signatures / 5bab593fb910 / 4

- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-2bd07ea4216146b2e4502921f872c58ec5c7c42a945c1ae01ffc62e989884f4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27b32c80bf2f3aa9c578fffc678b8db286583adcb1d42bdc0951c758d1bbef4a"></a>

## detection_settings.signature_selection_setting.signature_settings_by_accuracy — detection_settings.signature_selection_setting.signature_settings_by_accuracy / 3b1914642213 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- detection_settings.signature_selection_setting.signature_settings_by_accuracy

<a id="canonical-c0beabbb8882569593da08385cf267a4c1eb49c181c622a07c87bd8342abc898"></a>

Type: `"single"`. Computed.

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

<a id="canonical-450accaca0fd7267d4e0c78a328e7087383c7a1a231d386297401bfc758f4636"></a>

## Direct properties — detection_settings.signature_selection_setting.signature_settings_by_accuracy / 3b1914642213 / 3

<a id="canonical-55bcb971244bae864ef7a21de0c0a7fd930f248ec4a2296c872dcca76e3f5707"></a>

<a id="canonical-c9fa64f68ce8269c5732e86bb1a52e688438345210ee3c72efb35b67ef15f8f8"></a>

## high_accuracy_action property — detection_settings.signature_selection_setting.signature_settings_by_accuracy / 3b1914642213 / 4

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-82c923b03efce4e15f683d9a7c6ea77be1e12ef81cf8130b4d3b764961591c8c"></a>

<a id="canonical-166a10bbaff10b5c00b5ec740dfa590fe47ed5d7e8d5edaed0ff8639786c6fd4"></a>

## low_accuracy_action property — detection_settings.signature_selection_setting.signature_settings_by_accuracy / 3b1914642213 / 5

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-55a64bf0b549df476be1fc1efdeed72ba2656b6123adf74d7e187bf4743ebbdd"></a>

<a id="canonical-a3c2156a717fc3f28293e578c09444ceb13157656100f969ce4638748ba214d9"></a>

## medium_accuracy_action property — detection_settings.signature_selection_setting.signature_settings_by_accuracy / 3b1914642213 / 6

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

<a id="canonical-0b26022eefc23efdf6b42d163dfac081b04f934173ce5ce06c7ebe6b854ddc2a"></a>

## Next pages — detection_settings.signature_selection_setting.signature_settings_by_accuracy / 3b1914642213 / 7

- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-e5c273bdcd16bc12d0b8eee7ca1c5fd1bee0d554f1eb008477bcd24887c1c2ec)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-ee826af0f82d4fb0bde2520d8c49c0282b585d3961dbd785e570706f14f21f85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59be785a4162c8ffbcda71304108c712fb2bacbac040739385a69b7afb1512bf"></a>

## detection_settings.stage_new_and_updated_signatures — detection_settings.stage_new_and_updated_signatures / dc2bd9521a84 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.stage_new_and_updated_signatures

<a id="canonical-a1c67aa6ad9ddfe476edb6e19fc091ddc4151e64ff282111ead19f3d74bdac43"></a>

Type: `"single"`. Computed.

Attack Signatures staging configuration.

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

<a id="canonical-2f78f332ed34369c0edbdef202dbeabeaab6fcb08c30727826cbcefa54866fb2"></a>

## Direct properties — detection_settings.stage_new_and_updated_signatures / dc2bd9521a84 / 3

<a id="canonical-bdea08690aea9bb583149ce7ea7e7bee7e422e431621e7e24f5d977c35bb8e79"></a>

<a id="canonical-ba7ae24965b9321448897cd07915be1b026d6804863934a01f52e2f255efab52"></a>

## staging_period property — detection_settings.stage_new_and_updated_signatures / dc2bd9521a84 / 4

Type: `"number"`. Computed.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Upstream description:

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

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

<a id="canonical-0496aa0729b369d9cee5b976ca3a9ffd8b815bcf2a0b1de36e1b31de9900d012"></a>

## Next pages — detection_settings.stage_new_and_updated_signatures / dc2bd9521a84 / 5

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-11cd9cbaab4552932b57a4e1bdc033a604ec96b00c1df82cab84861f1967e993"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db819f87c74db180801b0b6b13ef712b37c7cd6066c291ffbd79869cd11615e8"></a>

## detection_settings.stage_new_signatures — detection_settings.stage_new_signatures / 5ba239994a10 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.stage_new_signatures

<a id="canonical-2cce129b6a32c732cecf82c10af56fcb6c74e9222a5844f8f2bbbb58fb9e08e9"></a>

Type: `"single"`. Computed.

Attack Signatures staging configuration.

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

<a id="canonical-cc32c8fdf2f5fbe13d41d45a7cc77326d191cae06031fbb84cbf577958077b4d"></a>

## Direct properties — detection_settings.stage_new_signatures / 5ba239994a10 / 3

<a id="canonical-6d8ea3698c36e850d067c9c9dc76a7197626a1faff402aab0532aed41baf2f03"></a>

<a id="canonical-2d04a791bf8857f4704df40f86c7d99985091884e914e17217003d4e6db695df"></a>

## staging_period property — detection_settings.stage_new_signatures / 5ba239994a10 / 4

Type: `"number"`. Computed.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Upstream description:

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

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

<a id="canonical-918c170d5d8f098ce3b7f9bab773945d69b54478c11b54da1443d6e302f457ee"></a>

## Next pages — detection_settings.stage_new_signatures / 5ba239994a10 / 5

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-59b777e240fa8bb798a751970904e94a184807cdaf91d7c8a1455b69fe20678d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49b1fbecc2840ebc0ad38687df861e9cd2a6ee85cf97a6806a2dd1e20f495587"></a>

## detection_settings.violation_settings — detection_settings.violation_settings / af918413522f / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.violation_settings

<a id="canonical-67110e6af0928c20d1252aa63dce223c5ba469085c72ab6ec981cef7a09d338c"></a>

Type: `"single"`. Computed.

Specifies violation settings to be used by WAF.

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

<a id="canonical-12588c9dd88ff2838b12f3497ca5a3b1157c106e4b0e571077d24919c57d91c8"></a>

## Direct properties — detection_settings.violation_settings / af918413522f / 3

<a id="canonical-0212ce5b34c9c5f9a99e195a9a6d03a5f9ff2a5da3c161e1c2fe16f88108bb5d"></a>

<a id="canonical-ed9c8baf220f2ba4af222128f294e4c941dac6ffef4879583d5e711965d7179c"></a>

## disabled_violation_types property — detection_settings.violation_settings / af918413522f / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-6c0b04f340345b5e1f6bd472efba8b713dcf6e1165bee238f09ddba561b7e2be"></a>

## Next pages — detection_settings.violation_settings / af918413522f / 5

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-d05829dd59400aa86a661ccecf1b306a92f35d1c494f1096291e3d10ba680875"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa3eb12b2335597bb4390700e68ea9cda01d873d2320abee25c3d5515421f640"></a>

## detection_settings.violations_view — detection_settings.violations_view / bf5d0db1042a / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- detection_settings.violations_view

<a id="canonical-63fb9a7795c8bdf159545de3cd99be7bc29e7f89752a9169f97df7bad3c033e9"></a>

Type: `"list"`. Computed.

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

<a id="canonical-13923c738c5b4b87c074aaf61e7e1245a56bcb941001a12deffb981467ec7c37"></a>

## Direct properties — detection_settings.violations_view / bf5d0db1042a / 3

<a id="canonical-ef30bb24fd313f55b717c6811c9284b3a021c84b8de6d89d6921481c0ae6d531"></a>

<a id="canonical-9aff7cd9257d2722c1d945ac26afc1870431f62d2794496f60dbde37a8baa077"></a>

## description_spec property — detection_settings.violations_view / bf5d0db1042a / 4

Type: `"string"`. Computed.

Description. Human-readable description text

<a id="canonical-a18afff24908577a9b13e9455c1b51e7e69efd844935c367307efc2175b6a721"></a>

<a id="canonical-eca3e9614c5f7556feef58585cb13707f4ab190642129580f235d3ba38256b1b"></a>

## enabled property — detection_settings.violations_view / bf5d0db1042a / 5

Type: `"bool"`. Computed.

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

<a id="canonical-dab606ca3cfc914b5b298827e7d4769625a28ba7a294c831f9d7c78b6cf7e962"></a>

<a id="canonical-e54230e4931ffe49d53a168d7a38f73f577e513ff45a948ceaa26332e4ae48e6"></a>

## enabled_by_default property — detection_settings.violations_view / bf5d0db1042a / 6

Type: `"string"`. Computed.

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

<a id="canonical-eadf44bdb18fb8a5603c3a6940db5503ac54e4b93b431143bc631c9da94e6a03"></a>

<a id="canonical-e980c8f74d2d1bf6d488ab7edb7449b25a08be28b210816d8883dcecec7fa9b7"></a>

## name property — detection_settings.violations_view / bf5d0db1042a / 7

Type: `"string"`. Computed.

Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-f67751e868a154ae6ead434bbcc46ba211164428c0fdc8bc0b7dfa8f2b41b2a3"></a>

<a id="canonical-29337b3f8794bd8372b36fac597e67bed64898cfaee8d02eb4c26b1edb615224"></a>

## title property — detection_settings.violations_view / bf5d0db1042a / 8

Type: `"string"`. Computed.

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

<a id="canonical-55e50cc3b577dda2f38a45265747211fae75feda7afedcfae9a8bf0a9a48b6f6"></a>

## Next pages — detection_settings.violations_view / bf5d0db1042a / 9

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-bfb5a7fd498a61f3607c0c33bfd0c74f15a78222597e6140b1725b84bc6533a4)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-44e035f938c10b60f96cf72524a61411b2d81a63d6ed271c32afa1919755c343"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77feccf89b73b7a424929c5108e669f035e7981b86bc19fd239bd9ac5cf3221f"></a>

## disable_ai_enhancements — disable_ai_enhancements / 19a6333713fa / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- disable_ai_enhancements

<a id="canonical-b013e1fa336f759368c1763e2e71c77f78e2297ea9a8443d7a594ce1fa8cef8a"></a>

Type: `["object", {}]`. Computed.

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

- [disable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-b013e1fa336f759368c1763e2e71c77f78e2297ea9a8443d7a594ce1fa8cef8a)
- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-a8799cac70fd7916220a34f1a08ca9ebe81ab22890cc715c33a811727229c0e7)

Select alternatives according to the provider validators above.

<a id="canonical-3f20b8fa0e3be9b68e6620e77bd6e37ea816e369923667bde03331afdb9765d2"></a>

## Direct properties — disable_ai_enhancements / 19a6333713fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-231e63eeb433e98002f7a0f022fa1e4a727cb0fad67bc53ee81a6124221254ba"></a>

## Next pages — disable_ai_enhancements / 19a6333713fa / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-f3e7b68ac4045b02148757c9c852efeac2f9b18e7778c684d1a47c8d6e8412ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d90f757062fcdae37870462b2cb8251d3f29bb12ff1dff2fcb62c30fd677668"></a>

## disable_anonymization — disable_anonymization / a9b7254a0482 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- disable_anonymization

<a id="canonical-078327a6c8e4882459102e191b325bfd490e7adfb21f1bf0bd96f6baf399c0a8"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-8241fe777ac263cce47052ba82836773c9ea4c5e388d4f37c61ed29be042d42b"></a>

## Direct properties — disable_anonymization / a9b7254a0482 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da2cae1f589fdcd6d265ff32799de4c52462b620692df060ecf54cac2aebe98e"></a>

## Next pages — disable_anonymization / a9b7254a0482 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-fcec69530d75cc1a8525164bd619b25bedc1adfd5198f98ef2eaf9aee0e68ae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23d1b3c2fd4faa6030a034536e5fc53281cd0bde8bf7185c4d4e5687c5015910"></a>

## enable_ai_enhancements — enable_ai_enhancements / 912c698a85d5 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- enable_ai_enhancements

<a id="canonical-a8799cac70fd7916220a34f1a08ca9ebe81ab22890cc715c33a811727229c0e7"></a>

Type: `"single"`. Computed.

Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.

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

<a id="canonical-46f7586f75b6fddc3378e78359f7bc79d0a29ebfa2acbd15fcbc1298c59399ff"></a>

## Direct properties — enable_ai_enhancements / 912c698a85d5 / 3

- [mitigate_high_medium_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-b35092ffd29e7469d4c8a7e50fe58f6ad90d86855e76199bb18615df5623d85b): complete subsection reference.

- [mitigate_high_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-5d37a815b3fae36d42e4fe89daf129b4a18f04841c72485fb8e7f2f5af512e2f): complete subsection reference.

<a id="canonical-52ab197eb773b6237a1ed36275afe4213c2d2a673d9379b32fb807629d27b904"></a>

## Next pages — enable_ai_enhancements / 912c698a85d5 / 4

- [enable_ai_enhancements.mitigate_high_medium_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-b35092ffd29e7469d4c8a7e50fe58f6ad90d86855e76199bb18615df5623d85b)
- [enable_ai_enhancements.mitigate_high_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-5d37a815b3fae36d42e4fe89daf129b4a18f04841c72485fb8e7f2f5af512e2f)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-b35092ffd29e7469d4c8a7e50fe58f6ad90d86855e76199bb18615df5623d85b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07c4e54309d513be85dd2834f3e800f9a98e9e92155f497abdfbea17f3661655"></a>

## enable_ai_enhancements.mitigate_high_medium_risk_action — enable_ai_enhancements.mitigate_high_medium_risk_action / fbeeaffe5978 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-fcec69530d75cc1a8525164bd619b25bedc1adfd5198f98ef2eaf9aee0e68ae5)
- enable_ai_enhancements.mitigate_high_medium_risk_action

<a id="canonical-6f1be2a13759d970c93f905046ab1d7913286d9c68788f62f39cea6cd5b9ff11"></a>

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

<a id="canonical-039ffaeced45d28bc4ac1b501d19f51d1ceb5c27e36cfa558135b423fe65479e"></a>

## Direct properties — enable_ai_enhancements.mitigate_high_medium_risk_action / fbeeaffe5978 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e70eba40709521402e48deb81c8be1fcd6f1893251e8f51e076c40cfd026973c"></a>

## Next pages — enable_ai_enhancements.mitigate_high_medium_risk_action / fbeeaffe5978 / 4

- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-fcec69530d75cc1a8525164bd619b25bedc1adfd5198f98ef2eaf9aee0e68ae5)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-5d37a815b3fae36d42e4fe89daf129b4a18f04841c72485fb8e7f2f5af512e2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67cbf489e02346c690378b4dc530f3aa616da9d9357702820b756551046bc7ff"></a>

## enable_ai_enhancements.mitigate_high_risk_action — enable_ai_enhancements.mitigate_high_risk_action / b8c8fe48b491 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-fcec69530d75cc1a8525164bd619b25bedc1adfd5198f98ef2eaf9aee0e68ae5)
- enable_ai_enhancements.mitigate_high_risk_action

<a id="canonical-d3fbc5deeeafff8a77a82ca06535339612cb655e05e900125deb268ac2e31f4e"></a>

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

<a id="canonical-19d30cb0f6084a0e9be446b7f0df8b7d212bb97276cc9e1ec32282c1f7827935"></a>

## Direct properties — enable_ai_enhancements.mitigate_high_risk_action / b8c8fe48b491 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b4e1ebe3f6d5b9ba21e0bfa672f9f8a8e98279d21fd2907e016abcbdc135d1bf"></a>

## Next pages — enable_ai_enhancements.mitigate_high_risk_action / b8c8fe48b491 / 4

- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-fcec69530d75cc1a8525164bd619b25bedc1adfd5198f98ef2eaf9aee0e68ae5)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-1ef81513048703b410f3651ef1afd3b69d8bd753f8118b1431462ebc01179857"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd3bbb8f4760b7c3523985d09233d28560e772ed06520b2c2931b07fe16eb48c"></a>

## monitoring — monitoring / 930eb34f9c76 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- monitoring

<a id="canonical-f0a4fcc34a2ee4417ff037db52f9184b895611e2a88c9f7cbb0b2780ed8400f7"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-8692b19c91057e619dec3d9f6695d31bdd343ac0b89d66d0e11da0f71d9debc2"></a>

## Direct properties — monitoring / 930eb34f9c76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31ba264b5500c7e42b99092c30cbd010d7830d92a84fd045ceb793b5b39c478c"></a>

## Next pages — monitoring / 930eb34f9c76 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-9f246fc996a562416e707b574f84dce284240997ed47a6fd1bbf0bb3f4725ffc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c85de6f98e2d6a40351980bb101b022c30d3c6644d4c05a0b48432948c51e382"></a>

## use_default_blocking_page — use_default_blocking_page / 743b8d9919b4 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- use_default_blocking_page

<a id="canonical-396dd05afe9fdf2565125ad9bd2c54a87f5d956d26d832c0d873b3011c86ef43"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0461f3d1c5088017edc4262128219a075c40ae07bd1905abcf0ca4ea61ee2423"></a>

## Direct properties — use_default_blocking_page / 743b8d9919b4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1512062e24cd87d03d1151c61375aab31df94f8e4ca8ccf13e2b084fcb374f4"></a>

## Next pages — use_default_blocking_page / 743b8d9919b4 / 4

- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
