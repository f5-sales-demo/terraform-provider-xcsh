---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87fc88d0c6a44f87197fcbf4aff222b6f58259105a3b061982e253970bb44ad9"></a>

## Property reference — Property reference / 0f3a7fc24589 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- Property reference

<a id="canonical-ddb664b6a76c7cee0a2027af699c89218871beb06035cb25437003edb4063743"></a>

## Direct properties — Property reference / 0f3a7fc24589 / 3

- [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-9d682ea065487da43edc1ea5e3d025aa52b3425b72cdc1e324d288ee618f49e3): complete subsection reference.

<a id="canonical-2f9446a8948518071d51f4b91af87432ec3be9cdc4f307295e637d963dcf3e76"></a>

<a id="canonical-7fea36548b4278295fbbae1f0b646483a44350e440bbd65faa855cdf43dd5aff"></a>

## annotations property — Property reference / 0f3a7fc24589 / 4

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

- [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-d99334616dce20db6de192b4c1846fe58475d4fb2cf375c0ef2dd7a5fe8c22c0): complete subsection reference.

- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd): complete subsection reference.

- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074): complete subsection reference.

- [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-9d2981ecfbe9d0b8d6732e231eb0af76e2721080347b0d3c532913f663fd36da): complete subsection reference.

<a id="canonical-f6fc69e53bb386d3206759f5e059b4f9a1ae3eb72d5598eeb79374af95c7fa3f"></a>

<a id="canonical-c8646f2cd7d0a70675c49597851dec9b56c973b77caa54924391c1634d63fc40"></a>

## description property — Property reference / 0f3a7fc24589 / 5

Type: `"string"`. Computed.

Description of the ProtectedApplication.

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

- [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-a9e6f84ab6c0e0483f50d3fd1f9d0e2837df1e409dc26e4065cb0801b072cd71): complete subsection reference.

<a id="canonical-95054b0d91fe8ed7e509a293def72a4265b577cafd8bf518e8e6ff21e5f6e52f"></a>

<a id="canonical-4d8978314defc5d0a7691385a721d2c5510a7dee5da71860688e76e41315bf2c"></a>

## id property — Property reference / 0f3a7fc24589 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-9eb30875d5d3c0aad8b10e894ace27531bc4c582bba0c1f95bbd3ff70366353c"></a>

<a id="canonical-d599cf11cd53ddcb176f14d3aafe0ae46b7a360b9a4bc48c4849528a5bd162c9"></a>

## labels property — Property reference / 0f3a7fc24589 / 7

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

<a id="canonical-0b6dbb265397a55b0418e5a7bb0b55f098c9e81adfbefdea4a09581241be5da9"></a>

<a id="canonical-8f3d06382d0e09d8ffa4b0bfcc1678b3035e55401990f7155be85d3a7384918f"></a>

## name property — Property reference / 0f3a7fc24589 / 8

Type: `"string"`. Required.

Name of the ProtectedApplication.

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

<a id="canonical-744d246b2c27b9b9df9ac66a7b7460ba6c1e010b6db4338f1b243463d00642ff"></a>

<a id="canonical-6e586b7aa6fbb621cf41e586915fffd9191b5119ffdd41738b33586452b5ebf9"></a>

## namespace property — Property reference / 0f3a7fc24589 / 9

Type: `"string"`. Required.

Namespace where the ProtectedApplication exists.

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

<a id="canonical-044811827b8d59d08ea0aea8f8a16827dede567e0d076b604d381dfaffd1571a"></a>

<a id="canonical-a161fe87fb475b67b762f9514730ce433c6949ae0de65af45c2a142cb6862fb7"></a>

## region property — Property reference / 0f3a7fc24589 / 10

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- US: US

United States of America &#8203;- EU: EU

European Union &#8203;- ASIA: ASIA

Asia &#8203;- CA: CA

Canada.

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA",
    "CA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-f04c990851b7f2ef626b63087a3226d19a94d9cc23bebf2bfd39037ad13c276a): complete subsection reference.

<a id="canonical-4b2560d4c664067fb89c2dad9b731a363e4d8873c9d2f77c9d77a44795f76451"></a>

## All schema paths — Property reference / 0f3a7fc24589 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `adobe_commerce_connector` | [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-05d44ea52d6e9a0ea9b12220784d2aa7d2de2fed92a0e90256b9b4c4c4f7fc75) |
| `annotations` | [annotations](data-sources--protected_application--reference--group-001.md#canonical-2f9446a8948518071d51f4b91af87432ec3be9cdc4f307295e637d963dcf3e76) |
| `big_ip_iapp` | [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-b41324199086a3d71911ffeafdc42fb683d57222a1b5f896f0d313f812bace4f) |
| `cloudflare` | [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-cb54f9a5556bc8b1065d3f18d6c059da55fb4e52914ad0a2f8fcbed50b51b501) |
| `cloudflare.continue_mitigation_action_hdr` | [cloudflare.continue_mitigation_action_hdr](data-sources--protected_application--reference--group-001.md#canonical-a283efd73921dd0b8ad59198064b74cee4001f9bfdcfc4b048aaca72d785edfb) |
| `cloudflare.disable_js_insert` | [cloudflare.disable_js_insert](data-sources--protected_application--reference--group-001.md#canonical-1b91afdb539e1ab18e2ef6648838b17f0b109895489d8994ddef75ced0357940) |
| `cloudflare.disable_mobile_sdk` | [cloudflare.disable_mobile_sdk](data-sources--protected_application--reference--group-001.md#canonical-18b083946f720fe56ba43de935374a8750643fb44a02e25384f4ced09fcba7b0) |
| `cloudflare.js_insertion_rules` | [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-cf726e6368994bdfe7b8ebe402190ecb3cd9dad8d19d888dd1dffbb54f029563) |
| `cloudflare.js_insertion_rules.exclude_list` | [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-9120827327b2d0f0d431cce3aa97acaa52d672f9e3cca51575f6170bccaad263) |
| `cloudflare.js_insertion_rules.exclude_list.any_domain` | [cloudflare.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-001.md#canonical-953602b5c5c9f4a553d6ad63a77c298b29376ac70b44f8d08a5b564a21a2bc00) |
| `cloudflare.js_insertion_rules.exclude_list.domain` | [cloudflare.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-001.md#canonical-9b1584a5a0218fc819d0c156a65c04f53b1b58aaadee0700a925e86381053ca7) |
| `cloudflare.js_insertion_rules.exclude_list.domain.exact_value` | [cloudflare.js_insertion_rules.exclude_list.domain.exact_value](data-sources--protected_application--reference--group-001.md#canonical-8c5523888ad901a12a619f30a328d6b3b0320e24d84162650f51620d6f9166c9) |
| `cloudflare.js_insertion_rules.exclude_list.domain.regex_value` | [cloudflare.js_insertion_rules.exclude_list.domain.regex_value](data-sources--protected_application--reference--group-001.md#canonical-bf8d1ad3a802269f334d254aee3aef8f7fa84a6c0c554212302e40aec3ec0a19) |
| `cloudflare.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudflare.js_insertion_rules.exclude_list.domain.suffix_value](data-sources--protected_application--reference--group-001.md#canonical-74ac002bfc060d53a1baa5fab4484aaf71b736140db9dc5e3e8adbba8afff7d0) |
| `cloudflare.js_insertion_rules.exclude_list.metadata` | [cloudflare.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-001.md#canonical-432df8a928168b6462fcbd765de52d4c6a191b8178ca7286411e00f1ebb435f1) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudflare.js_insertion_rules.exclude_list.metadata.description_spec](data-sources--protected_application--reference--group-001.md#canonical-7509691e54b5540d924013c40d78e91dd1d3e08f1ba303f1f62cf4884fbe263a) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.name` | [cloudflare.js_insertion_rules.exclude_list.metadata.name](data-sources--protected_application--reference--group-001.md#canonical-e75126b66119ea451d512bb063ef904f724e7bc932d31d35f99c57c0e6480728) |
| `cloudflare.js_insertion_rules.exclude_list.path` | [cloudflare.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-001.md#canonical-53ac78942289025c8d09b6394594c3f6b92eee9e963ef724d0ce5ad2463d5910) |
| `cloudflare.js_insertion_rules.exclude_list.path.path` | [cloudflare.js_insertion_rules.exclude_list.path.path](data-sources--protected_application--reference--group-001.md#canonical-10a1a2cf04ad370a6e645459343edfe7834d4bd6078108eab38f600e02b678dc) |
| `cloudflare.js_insertion_rules.exclude_list.path.prefix` | [cloudflare.js_insertion_rules.exclude_list.path.prefix](data-sources--protected_application--reference--group-001.md#canonical-374b13a718795c306d41d3fd02d66055b6e981c825ce306a76a27c2f1066b509) |
| `cloudflare.js_insertion_rules.exclude_list.path.regex` | [cloudflare.js_insertion_rules.exclude_list.path.regex](data-sources--protected_application--reference--group-001.md#canonical-70d7bd03045d145eb3076a10fa03bb42abbad6c10c38fc7a0b43237382f51bba) |
| `cloudflare.js_insertion_rules.javascript_location` | [cloudflare.js_insertion_rules.javascript_location](data-sources--protected_application--reference--group-001.md#canonical-b2ec916d86b0823c2126f1598d0dbe6aee00f76ce585df2a6c482a0632da35d3) |
| `cloudflare.js_insertion_rules.js_download_path` | [cloudflare.js_insertion_rules.js_download_path](data-sources--protected_application--reference--group-001.md#canonical-44e1ad65d7ebf434b92e72b4bb62ee00e0388d8693e70eb377e5b478a680ffab) |
| `cloudflare.js_insertion_rules.rules` | [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-3142f7fa3b9bf4e6cb93ff5f5ed3a938ca6ee56b3fc7800ebfb4f15e87ba8207) |
| `cloudflare.js_insertion_rules.rules.any_domain` | [cloudflare.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-001.md#canonical-8f2ed1547e9e98d3f7f3c5e2e0137864ca7f0650a4bb2eefae44784d8774c9f8) |
| `cloudflare.js_insertion_rules.rules.domain` | [cloudflare.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-001.md#canonical-16d797752145174f3dcaa34e88cece05fc7063f4cc00cb417fbfc822ab8ce479) |
| `cloudflare.js_insertion_rules.rules.domain.exact_value` | [cloudflare.js_insertion_rules.rules.domain.exact_value](data-sources--protected_application--reference--group-001.md#canonical-88a6b7469d3dbb92ecaee3f2f4e47423b28a9db421fca524fe00651cd8844c31) |
| `cloudflare.js_insertion_rules.rules.domain.regex_value` | [cloudflare.js_insertion_rules.rules.domain.regex_value](data-sources--protected_application--reference--group-001.md#canonical-850bc0f9cd3c157f90c4bf7cba399c58201e059d3712cad20c3b988228858690) |
| `cloudflare.js_insertion_rules.rules.domain.suffix_value` | [cloudflare.js_insertion_rules.rules.domain.suffix_value](data-sources--protected_application--reference--group-001.md#canonical-b79b22fffa02d37fd50fa753be75d3d41d61f9cd64952d208e09b1d33e3d2181) |
| `cloudflare.js_insertion_rules.rules.exact_path` | [cloudflare.js_insertion_rules.rules.exact_path](data-sources--protected_application--reference--group-001.md#canonical-abd0166f86a670a232009855a32deb84081bd74feb974a8bfc7789ca9c65bfad) |
| `cloudflare.js_insertion_rules.rules.glob` | [cloudflare.js_insertion_rules.rules.glob](data-sources--protected_application--reference--group-001.md#canonical-4797809dca93df35a48458ce19252f4589b3ee9301247b1566df4d2e0b547c75) |
| `cloudflare.js_insertion_rules.rules.metadata` | [cloudflare.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-001.md#canonical-6de971da694f973c43ae7a72c24eed15d5bbdd5b98919264dd99843021c783d8) |
| `cloudflare.js_insertion_rules.rules.metadata.description_spec` | [cloudflare.js_insertion_rules.rules.metadata.description_spec](data-sources--protected_application--reference--group-001.md#canonical-6e87cbe51a484d4115689b7f517c8c82732ece6b7033f903d870ba57c081177a) |
| `cloudflare.js_insertion_rules.rules.metadata.name` | [cloudflare.js_insertion_rules.rules.metadata.name](data-sources--protected_application--reference--group-001.md#canonical-6ceb6700da034965f62be5542c369f33b966954dab65fcc91e7f7413ee64576d) |
| `cloudflare.js_insertion_rules.rules.prefix` | [cloudflare.js_insertion_rules.rules.prefix](data-sources--protected_application--reference--group-001.md#canonical-db124372f7f4fdc79aa9640fb89e56bd811433c8904987faae12e542918991d2) |
| `cloudflare.loglevel` | [cloudflare.loglevel](data-sources--protected_application--reference--group-001.md#canonical-aeea0c3086c2f4da5f5ead949e034be0efa0eae2fa4331c0dcfb6f3a423e7d56) |
| `cloudflare.manual_js_insert` | [cloudflare.manual_js_insert](data-sources--protected_application--reference--group-001.md#canonical-c45398b4d2684259668c778485de845d926bea00bc86b37b9637af70dfc77e44) |
| `cloudflare.manual_js_insert.js_download_path` | [cloudflare.manual_js_insert.js_download_path](data-sources--protected_application--reference--group-001.md#canonical-efd714f974ed904485b0f9c9c4443d6e786a63932687e4c0790b2a88ab71bd7d) |
| `cloudflare.mobile_sdk_config` | [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-76ec3f5ba91aeb95484bc39e05a1b2320c69322fb4f52746f36f9c6dfec13282) |
| `cloudflare.mobile_sdk_config.mobile_identifier` | [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-0c0e9f173618eec54381cd0daff5bf6e4b66edb86d291d3394cd77637aa773a6) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers` | [cloudflare.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-001.md#canonical-63f1cdd1c66d2eecdfafa1ee89d00fe52382f88f19d95fb7864602013a8c41cf) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.exact](data-sources--protected_application--reference--group-001.md#canonical-40383724c736d35fdd5c8a314c4ea152a6d302dc061ed94264fb48064a4fee36) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.name` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.name](data-sources--protected_application--reference--group-001.md#canonical-2809e36e0abe8897a1c7410876d6f2bb8e02ae91b62304f8e21ef3a5742ce42e) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.regex](data-sources--protected_application--reference--group-001.md#canonical-215c0cd1b564aaa735611c2e8e5269e6cec4a3b69899e03fd53324f966b8d6b3) |
| `cloudflare.protected_endpoints` | [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-b06731213efa4a56bf9b6b10c576725c84fe8dc855a7e59249f58f701b3c6240) |
| `cloudflare.protected_endpoints.any_domain` | [cloudflare.protected_endpoints.any_domain](data-sources--protected_application--reference--group-001.md#canonical-ebadff18fef6e9d651df4a7971dd0bd363fa66e6cf82e47d23cfe0fd5fbe3843) |
| `cloudflare.protected_endpoints.domain` | [cloudflare.protected_endpoints.domain](data-sources--protected_application--reference--group-001.md#canonical-2b984be80640971b67dbd8f34af829cc13d3a289cc497fe0e77f390deeb7803c) |
| `cloudflare.protected_endpoints.domain.exact_value` | [cloudflare.protected_endpoints.domain.exact_value](data-sources--protected_application--reference--group-001.md#canonical-8517bf640e5fb50c54faa2053c3a19749554f4c16ebb3598c180ad6c6069499e) |
| `cloudflare.protected_endpoints.domain.regex_value` | [cloudflare.protected_endpoints.domain.regex_value](data-sources--protected_application--reference--group-001.md#canonical-96fc2f012f1c1afcc7f5718e9ae87d8d39749d4e236f9becf4f2db62eac04da9) |
| `cloudflare.protected_endpoints.domain.suffix_value` | [cloudflare.protected_endpoints.domain.suffix_value](data-sources--protected_application--reference--group-001.md#canonical-0f7f2e7adfe090c9c679b89b03418c045425967b233d31d72e0695ed244de600) |
| `cloudflare.protected_endpoints.http_methods` | [cloudflare.protected_endpoints.http_methods](data-sources--protected_application--reference--group-001.md#canonical-70069403fa976335a8d6ff04601461e917b0c5ee6ad8fdd1f1b853eb056f7298) |
| `cloudflare.protected_endpoints.metadata` | [cloudflare.protected_endpoints.metadata](data-sources--protected_application--reference--group-001.md#canonical-e9298dc67363cdef08e52e5c152b4e45bb87b5184635dab9e044f6c78d73f466) |
| `cloudflare.protected_endpoints.metadata.description_spec` | [cloudflare.protected_endpoints.metadata.description_spec](data-sources--protected_application--reference--group-001.md#canonical-2ea10bbebe554f58914e3401c79d97fb80a417328db6a8d74cf840608bb7503c) |
| `cloudflare.protected_endpoints.metadata.name` | [cloudflare.protected_endpoints.metadata.name](data-sources--protected_application--reference--group-001.md#canonical-7a9c480d1b00a68439a857a3ee414c713bc15dc9a043b7df1714b274541ee1e9) |
| `cloudflare.protected_endpoints.mobile_client` | [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-44a3f339c558b2dc801835a9e1b3bdc14b284ff96b7afe1f7237897101e4873c) |
| `cloudflare.protected_endpoints.mobile_client.block` | [cloudflare.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-002.md#canonical-05339b8fd6f300eaa3dc04ba0267a908a6e6148ce57c57a7f01c9143bf992e41) |
| `cloudflare.protected_endpoints.mobile_client.block.body` | [cloudflare.protected_endpoints.mobile_client.block.body](data-sources--protected_application--reference--group-002.md#canonical-c298267a3dd71c950610a5cdac92916ac3818abe7a13c278fe33039633112d2b) |
| `cloudflare.protected_endpoints.mobile_client.block.content_type` | [cloudflare.protected_endpoints.mobile_client.block.content_type](data-sources--protected_application--reference--group-002.md#canonical-5e42d1fccba4c34090ed512e453f5fe6fbe4bd96a887eaf3ee0afa0dc02d834a) |
| `cloudflare.protected_endpoints.mobile_client.block.status` | [cloudflare.protected_endpoints.mobile_client.block.status](data-sources--protected_application--reference--group-002.md#canonical-3c3ff138d0085e0a9c42c3b06afa619b0f133f21e5bddc6340434bee6d55523a) |
| `cloudflare.protected_endpoints.mobile_client.continue` | [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-4baf9d24fd83c5b786f03d59fd9dbe205f34854dc366dc427740c7442e9332b4) |
| `cloudflare.protected_endpoints.mobile_client.continue.add_header` | [cloudflare.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-af7687d1f896dc7006082819793062e5758ce2cfc83b1be4f220ab9fd12d3b02) |
| `cloudflare.protected_endpoints.mobile_client.continue.no_header` | [cloudflare.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-dbe90f6da0c999382a59fad6c5ecea6785957b6fc957bb6198a5de56fc09064d) |
| `cloudflare.protected_endpoints.path` | [cloudflare.protected_endpoints.path](data-sources--protected_application--reference--group-002.md#canonical-407ba3d6edeb64866af26299302743e4c72506d6734feb6d9f4cede60c5ace03) |
| `cloudflare.protected_endpoints.path.caseinsensitive` | [cloudflare.protected_endpoints.path.caseinsensitive](data-sources--protected_application--reference--group-002.md#canonical-36ba7cbdb8db839d58080aaf37bfa9ed4dafabe2ccfde5dc60b4bc45d86b0b71) |
| `cloudflare.protected_endpoints.path.path` | [cloudflare.protected_endpoints.path.path](data-sources--protected_application--reference--group-002.md#canonical-7c9da84c0f9714ed7436ca3b4993a8ef1b1e62f185e035225300da90b9f249fd) |
| `cloudflare.protected_endpoints.query` | [cloudflare.protected_endpoints.query](data-sources--protected_application--reference--group-001.md#canonical-017a14b06058f069c022f2573ea2f88e80f13f0d2621935c84ed2ebbe31b145d) |
| `cloudflare.protected_endpoints.web_client` | [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-9b777c9ee64a564f5efbe0990817775526e74e777d7b7c79dbaf61a4be45fabe) |
| `cloudflare.protected_endpoints.web_client.block` | [cloudflare.protected_endpoints.web_client.block](data-sources--protected_application--reference--group-002.md#canonical-14e52a69c3b183602ac8194aff8bba90ea78e7d239263ac4caa3d9009db31e23) |
| `cloudflare.protected_endpoints.web_client.block.body` | [cloudflare.protected_endpoints.web_client.block.body](data-sources--protected_application--reference--group-002.md#canonical-550670b1f51e061ae78e83465e3bf694dac25465de8a6f31c33c893065de5b84) |
| `cloudflare.protected_endpoints.web_client.block.content_type` | [cloudflare.protected_endpoints.web_client.block.content_type](data-sources--protected_application--reference--group-002.md#canonical-4700cce8fea3916e54f86e77a48feb73c3f4a61524469c686e6808e888573162) |
| `cloudflare.protected_endpoints.web_client.block.status` | [cloudflare.protected_endpoints.web_client.block.status](data-sources--protected_application--reference--group-002.md#canonical-bbf240d50533d61745e35bdaa16d9a8d1d718ef3277636a21e8707561db9fa34) |
| `cloudflare.protected_endpoints.web_client.continue` | [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-cfad53267b1331c3e0f69ee9513f0da4f2bd3107ae965d008181ed677875b58a) |
| `cloudflare.protected_endpoints.web_client.continue.add_header` | [cloudflare.protected_endpoints.web_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-b2aa5c68cdd1274d53d2937e5c9b6aa73fad637f0bd225f42bbc7599387accc1) |
| `cloudflare.protected_endpoints.web_client.continue.no_header` | [cloudflare.protected_endpoints.web_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-eac1d24b4d4a3bf5d0ef5f053c5cb8cabb723df20fbe871d82840dddeaad82c7) |
| `cloudflare.protected_endpoints.web_client.redirect` | [cloudflare.protected_endpoints.web_client.redirect](data-sources--protected_application--reference--group-002.md#canonical-e78d54d2783cb48ddc19a19071ba9220b81b4804276d1dc1d81d36541a34f19c) |
| `cloudflare.protected_endpoints.web_client.redirect.location` | [cloudflare.protected_endpoints.web_client.redirect.location](data-sources--protected_application--reference--group-002.md#canonical-c3d89f9afde72303f2cc7a8845d14303842e64aab789bf8a9c2fa56c5dabea98) |
| `cloudflare.protected_endpoints.web_client.redirect.status` | [cloudflare.protected_endpoints.web_client.redirect.status](data-sources--protected_application--reference--group-002.md#canonical-3dbc7e6dac6b97c12941f9f9fdb88352bde8b4170d01f1ff540d20638ce27f8e) |
| `cloudflare.protected_endpoints.web_mobile_client` | [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-965b182fadd91510c266cf34ecd9e54d3f60bed9bec6ad0c2678b981a959c435) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--reference--group-002.md#canonical-81214094e1b82cd285fd45d95bc4ea2c7c42c61aff04151e582fa926b12052d9) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.body](data-sources--protected_application--reference--group-002.md#canonical-fa4a5241c578a70d7823767f9c21d6498cef49d617104be30a6d31a6c740e6f5) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type](data-sources--protected_application--reference--group-002.md#canonical-c37430ca101b42399bcae70454e5e745ca6e7f79b3c760c96611ae941d932a34) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.status](data-sources--protected_application--reference--group-002.md#canonical-7676a379f406c8aec645ee3a7e0bc3e60b81b09261ffdff3d9c0cf4dc91f0931) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web` | [cloudflare.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--reference--group-002.md#canonical-97bb0772834ee927c72c685d774aac0d21cc534aebb8ff89a92940763eb57c0b) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.body` | [cloudflare.protected_endpoints.web_mobile_client.block_web.body](data-sources--protected_application--reference--group-002.md#canonical-9040c472e8877d8c6f52d3dc57384e005497475bd6783d1eed12e8c0ed9ebef0) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_web.content_type](data-sources--protected_application--reference--group-002.md#canonical-0b6ef0b8cf5e01c0fcdb463c5fb191096d5341052783e8748792221114cd3a53) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.status` | [cloudflare.protected_endpoints.web_mobile_client.block_web.status](data-sources--protected_application--reference--group-002.md#canonical-6b574f802af64b5b71120bbc26330238dfedaca018f65d2b9d71268b12626f53) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-299042d5ad2bc29c51ce6cbacfb7eb1b909141158140b14ad95ce602ec801ad3) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](data-sources--protected_application--reference--group-002.md#canonical-3dcb00e18481a8ad1431bb06fed0b5f0f7226aea0abd7c8b6cf5ad0af6eff116) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](data-sources--protected_application--reference--group-002.md#canonical-38c8198083d0c3aa035ad8b60c493bb8ca07eeaa59ba2f0208406da79f953b08) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web` | [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-685d88d93972f54f20153209adfe4896260085867082be4c8091400bba24435d) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](data-sources--protected_application--reference--group-002.md#canonical-73dd28072a79af3c8af32375d52cae17052733043b17e5e4288b0cd9674967e0) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](data-sources--protected_application--reference--group-002.md#canonical-bbd6a5999745d7a174d9df064f91c8f2284d3e3ccb51702ee72c4b208430509c) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--reference--group-002.md#canonical-fad10e9e853c5430a34e7319022f65787b76c5be2a40c63e8f0a56f71c0c3a0e) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.location](data-sources--protected_application--reference--group-002.md#canonical-2b4d0c535ca191a8ff6d5b63dca95a2fe201f68ce100d3e9d626cb6092b37594) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.status](data-sources--protected_application--reference--group-002.md#canonical-79dfd52e93cc687e2eca564b9175dad7ebb096453ea22e0528f8585fc893d96f) |
| `cloudflare.timeout` | [cloudflare.timeout](data-sources--protected_application--reference--group-001.md#canonical-6aa70ad8d77f28876c3a67584ebd088fa809c218b48b355ede2d67c433b7a632) |
| `cloudflare.trusted_clients` | [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1f4ba568f5b810052c2413444a99be5a5b72f8d77bf24d48d4ef7ab1587171ef) |
| `cloudflare.trusted_clients.http_header` | [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-fd7090a059b3616d3d43b6ff8d4fb53a6b1cfc1c81bdc5424984eabd81abea53) |
| `cloudflare.trusted_clients.http_header.headers` | [cloudflare.trusted_clients.http_header.headers](data-sources--protected_application--reference--group-002.md#canonical-e1f85056e09d0943782537343853132337dd0df7979bcd06076a1fd369be4834) |
| `cloudflare.trusted_clients.http_header.headers.exact` | [cloudflare.trusted_clients.http_header.headers.exact](data-sources--protected_application--reference--group-002.md#canonical-8d0741159ba82ff1aaeb2b8f9c2233f06dbdb62584384d12fce218c46aac968f) |
| `cloudflare.trusted_clients.http_header.headers.name` | [cloudflare.trusted_clients.http_header.headers.name](data-sources--protected_application--reference--group-002.md#canonical-4ace4b7c3059cbd06beef9a26dd5161484b283c7cec50e2c9696a8099b638b5d) |
| `cloudflare.trusted_clients.http_header.headers.regex` | [cloudflare.trusted_clients.http_header.headers.regex](data-sources--protected_application--reference--group-002.md#canonical-554fcdaecbabf577c7a99a1eee36097b102f6dabb32a19b717ef3267a238b82d) |
| `cloudflare.trusted_clients.ip_prefix` | [cloudflare.trusted_clients.ip_prefix](data-sources--protected_application--reference--group-002.md#canonical-df162deaade587bc90f361967843e3d752710cfe6614567845f6d2ff38462c4b) |
| `cloudflare.trusted_clients.metadata` | [cloudflare.trusted_clients.metadata](data-sources--protected_application--reference--group-002.md#canonical-241c504f935e1fb347922dc3e7016667eef6b5a676d2f9ae170a8da20814d535) |
| `cloudflare.trusted_clients.metadata.description_spec` | [cloudflare.trusted_clients.metadata.description_spec](data-sources--protected_application--reference--group-002.md#canonical-d281605032e15770c12e2f2d91fbc8928814ccf68d888e10e12e4370f3bd37e4) |
| `cloudflare.trusted_clients.metadata.name` | [cloudflare.trusted_clients.metadata.name](data-sources--protected_application--reference--group-002.md#canonical-c8462c2e1351e2ea157a7f071e0bff778435e8710ea83ca573720c36c210e381) |
| `cloudfront` | [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-41b247837668c727fb5f889326104b14970f4fbba0b3be090903fb9277c9aa42) |
| `cloudfront.aws_configuration_id_selector` | [cloudfront.aws_configuration_id_selector](data-sources--protected_application--reference--group-002.md#canonical-281482fa84436ea871a92c64da4fc88ca3419ce012cf56ba6884d8549ef99d97) |
| `cloudfront.aws_configuration_id_selector.ids` | [cloudfront.aws_configuration_id_selector.ids](data-sources--protected_application--reference--group-002.md#canonical-267370a3bc2ae2d297fa45d49848460d8a53afcb496b14ba721d7de1b82659e0) |
| `cloudfront.aws_configuration_tag_selector` | [cloudfront.aws_configuration_tag_selector](data-sources--protected_application--reference--group-002.md#canonical-d999ac44092e4dca866dd6341fb01d9a72d1e14544cceefb582299f38d5e622c) |
| `cloudfront.aws_configuration_tag_selector.tags` | [cloudfront.aws_configuration_tag_selector.tags](data-sources--protected_application--reference--group-002.md#canonical-3460d7ebd7f21fb37d2de3bb8ba771986d3d2611b5ac18da5cadbd2261db1640) |
| `cloudfront.continue_mitigation_action_hdr` | [cloudfront.continue_mitigation_action_hdr](data-sources--protected_application--reference--group-002.md#canonical-713410c7a973ab47d007e8cda92fc79796d0fddae34445652281b07626945bc9) |
| `cloudfront.data_sample` | [cloudfront.data_sample](data-sources--protected_application--reference--group-002.md#canonical-5e4bbb61a2123d8eaafeaa5deac0a34501b8bc6d91161be11bf64617d69d7454) |
| `cloudfront.disable_aws_configuration` | [cloudfront.disable_aws_configuration](data-sources--protected_application--reference--group-002.md#canonical-a7a5404e40bc394a545f752ff1115e216544eab82b2047af9ffc2a51c04e7415) |
| `cloudfront.disable_js_insert` | [cloudfront.disable_js_insert](data-sources--protected_application--reference--group-002.md#canonical-d466dec7844bf89f6b0fb4014585c22aa9de343484f1224bda2c010c706665bb) |
| `cloudfront.disable_mobile_sdk` | [cloudfront.disable_mobile_sdk](data-sources--protected_application--reference--group-002.md#canonical-627efefb2af39b1220864f7c68d2eed00d00d106fd05cce7caeed9a4515e07f2) |
| `cloudfront.js_insertion_rules` | [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-aaa791313e2b8d45bd9b9f914d117d86e4bd94b409b8b2783562d1dc675554d7) |
| `cloudfront.js_insertion_rules.exclude_list` | [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-4a512b977af7af8241c94d710e66a36606bd68d50e7cd67e5855dc6f0dde1dc3) |
| `cloudfront.js_insertion_rules.exclude_list.any_domain` | [cloudfront.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-002.md#canonical-28b101c07f751195be7101f3edd2c87a01306c8d08c96dd8c89934e7f2a528ae) |
| `cloudfront.js_insertion_rules.exclude_list.domain` | [cloudfront.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-002.md#canonical-b41624e898361b588d542e9891927e462336f9140bd2a805fc66db14385c019c) |
| `cloudfront.js_insertion_rules.exclude_list.domain.exact_value` | [cloudfront.js_insertion_rules.exclude_list.domain.exact_value](data-sources--protected_application--reference--group-002.md#canonical-760c42599d7af5a3717f5f37721406e36d48982b655ebdca14d6f78e38b36643) |
| `cloudfront.js_insertion_rules.exclude_list.domain.regex_value` | [cloudfront.js_insertion_rules.exclude_list.domain.regex_value](data-sources--protected_application--reference--group-002.md#canonical-9627afbc9d50255f9d8e0ce6a6ac8036cb2d65afe86aa9c43e86de699a993c1f) |
| `cloudfront.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudfront.js_insertion_rules.exclude_list.domain.suffix_value](data-sources--protected_application--reference--group-002.md#canonical-e24d19077aaf13739506ecdb1c421525d571df9f95a24e644c0769dd30ed8fe7) |
| `cloudfront.js_insertion_rules.exclude_list.metadata` | [cloudfront.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-002.md#canonical-154919c36c5d536d22949ae7daeb800e957f368d0b11fc652cf9a15888cec202) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudfront.js_insertion_rules.exclude_list.metadata.description_spec](data-sources--protected_application--reference--group-002.md#canonical-66e97ee1ff7605a9b0caa9dbad244b024f06c40f3af3770c62064bc07509c3e4) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.name` | [cloudfront.js_insertion_rules.exclude_list.metadata.name](data-sources--protected_application--reference--group-002.md#canonical-25599158161290e25d755509bc26b032bb178237d06c3df0abe5ceeb9b745aea) |
| `cloudfront.js_insertion_rules.exclude_list.path` | [cloudfront.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-002.md#canonical-d982962766ecddb6187b35fae8b9dfe3385cfc3ffc8d83ddac01e90f4975e77c) |
| `cloudfront.js_insertion_rules.exclude_list.path.path` | [cloudfront.js_insertion_rules.exclude_list.path.path](data-sources--protected_application--reference--group-002.md#canonical-2f00889807c9f8e73cb4a2c245389ae1f9e7c2541ac0c05bcedc48ad1520cfdd) |
| `cloudfront.js_insertion_rules.exclude_list.path.prefix` | [cloudfront.js_insertion_rules.exclude_list.path.prefix](data-sources--protected_application--reference--group-002.md#canonical-8264e3e265a8bbb66b1bde98d84736995f5300a19e67fd95c02b46d1f7c80c03) |
| `cloudfront.js_insertion_rules.exclude_list.path.regex` | [cloudfront.js_insertion_rules.exclude_list.path.regex](data-sources--protected_application--reference--group-002.md#canonical-84376cc2e551bfdc04c071ef075dcf492cf3f48f7646773f7bfea84283a42776) |
| `cloudfront.js_insertion_rules.javascript_location` | [cloudfront.js_insertion_rules.javascript_location](data-sources--protected_application--reference--group-002.md#canonical-a79cdb66038cc502a8ddd958af4012cfc64d04be01b31219c027c4be829372ea) |
| `cloudfront.js_insertion_rules.javascript_mode` | [cloudfront.js_insertion_rules.javascript_mode](data-sources--protected_application--reference--group-002.md#canonical-008c89b19f62052e90ef19ab0ef245fa07e8739ad451c1d00551562bdda6de78) |
| `cloudfront.js_insertion_rules.js_download_path` | [cloudfront.js_insertion_rules.js_download_path](data-sources--protected_application--reference--group-002.md#canonical-e5937f84a4f9200c04ad58d2b142d4dea108d4095d82694b31b551df6f2cbb84) |
| `cloudfront.js_insertion_rules.rules` | [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-d697b8066bebbad9694874a0ad7d1f49b004ebff2d63c4e46693ec2846d772b9) |
| `cloudfront.js_insertion_rules.rules.any_domain` | [cloudfront.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-002.md#canonical-2a0ea93be25528e8529786273f5766e507591f17ab3e44edd6a9aad684184447) |
| `cloudfront.js_insertion_rules.rules.domain` | [cloudfront.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-002.md#canonical-fda66c1f6b9733f38473a25a126867748588329b07d6ba4e124afe2c9e9a8d74) |
| `cloudfront.js_insertion_rules.rules.domain.exact_value` | [cloudfront.js_insertion_rules.rules.domain.exact_value](data-sources--protected_application--reference--group-002.md#canonical-9bb3c970022d1b988669b77827feb196bd43cbdc279dc3c41ad016862ec0259d) |
| `cloudfront.js_insertion_rules.rules.domain.regex_value` | [cloudfront.js_insertion_rules.rules.domain.regex_value](data-sources--protected_application--reference--group-002.md#canonical-757b6d535f98a937925fe45f5af508c6dd7e693b2668aa2ab3834cc8e90cf022) |
| `cloudfront.js_insertion_rules.rules.domain.suffix_value` | [cloudfront.js_insertion_rules.rules.domain.suffix_value](data-sources--protected_application--reference--group-002.md#canonical-e7d1a35bf8ddc34c5f7a8a2e4b8790d9353b0128ae34f8fe66c3842c9bfb5c8e) |
| `cloudfront.js_insertion_rules.rules.exact_path` | [cloudfront.js_insertion_rules.rules.exact_path](data-sources--protected_application--reference--group-002.md#canonical-bc2ac7a5e90a05e05075b8b42443c34629589f7e8546c9fa99b8f6753e57ea8a) |
| `cloudfront.js_insertion_rules.rules.glob` | [cloudfront.js_insertion_rules.rules.glob](data-sources--protected_application--reference--group-002.md#canonical-270d24f12d7f3e7404ac0a100d3cfb73344717ff15c209f9fdd99606ed1b3a59) |
| `cloudfront.js_insertion_rules.rules.metadata` | [cloudfront.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-002.md#canonical-0ad324cd3db5233fd8e6ad7608743f55a63fbc7772ace366bd70a8087af47d51) |
| `cloudfront.js_insertion_rules.rules.metadata.description_spec` | [cloudfront.js_insertion_rules.rules.metadata.description_spec](data-sources--protected_application--reference--group-002.md#canonical-c1b76d5976b41e0903dd36e63cd965336d0f945df1b3357f6d31869e9d3a8ec8) |
| `cloudfront.js_insertion_rules.rules.metadata.name` | [cloudfront.js_insertion_rules.rules.metadata.name](data-sources--protected_application--reference--group-002.md#canonical-9b7c699fb48a8d6f27b0bcb9e1c39c07e48f6954d65e4589733265cb8c5232e5) |
| `cloudfront.js_insertion_rules.rules.prefix` | [cloudfront.js_insertion_rules.rules.prefix](data-sources--protected_application--reference--group-002.md#canonical-525ba94c541e770863e22cbd15d507a2852f20cc944a43eacad1633218b2292c) |
| `cloudfront.loglevel` | [cloudfront.loglevel](data-sources--protected_application--reference--group-002.md#canonical-a137552ac01ee7d9899aa4e4fd078ba03a62b5663650fc1c6ef7c4ccd64aebc5) |
| `cloudfront.manual_js_insert` | [cloudfront.manual_js_insert](data-sources--protected_application--reference--group-002.md#canonical-6d15f12bdab3a2a9c9a83bfd91d555c1b73cfd789122b57fa76df88a26b6299e) |
| `cloudfront.manual_js_insert.javascript_mode` | [cloudfront.manual_js_insert.javascript_mode](data-sources--protected_application--reference--group-002.md#canonical-8dadc1b8422e83d95bce96e093e4d387eb4d1718ec96ede38c8a80f68521f76e) |
| `cloudfront.manual_js_insert.js_download_path` | [cloudfront.manual_js_insert.js_download_path](data-sources--protected_application--reference--group-002.md#canonical-37f4c7e64e57b3b31024e528ee55634f1d1632023c7ad19bec98a51821cc108b) |
| `cloudfront.mobile_sdk_config` | [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-762646863dc0921bac8fd7e3697d69ccb149a9e92a7546d4f9978932f6097e1e) |
| `cloudfront.mobile_sdk_config.mobile_identifier` | [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-7c3bbb9078731bacdba2a8b1b38ed7556c2e18a1d80700b4d91aebd768527152) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers` | [cloudfront.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-002.md#canonical-7ad94d8fb817f77c9b710c033f24b9a2bfdbcd93d8ef08998b49df1cb05accaf) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.exact](data-sources--protected_application--reference--group-002.md#canonical-baacc2690f779e2d3dac529ad26922169813269b498246bbce0d3f4617ee749e) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.name` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.name](data-sources--protected_application--reference--group-002.md#canonical-12b541bc13d3795287bca29f18d1550ebc1d804ade6b671f64970563d94f1516) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.regex](data-sources--protected_application--reference--group-002.md#canonical-3ef5f4f2c67efd2f3e65179a6104ca8c0635c6ff8c16e5dcddba126be8d03392) |
| `cloudfront.protected_endpoints` | [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-2e9c89bb4831a4dead28d748a15fdc3315c79f219a235fbcf3393c6b81c28894) |
| `cloudfront.protected_endpoints.any_domain` | [cloudfront.protected_endpoints.any_domain](data-sources--protected_application--reference--group-002.md#canonical-3269411c8d4d34403208a9080ab224ebf73dfcd1771c88c4629cb55b7b0c50e5) |
| `cloudfront.protected_endpoints.domain` | [cloudfront.protected_endpoints.domain](data-sources--protected_application--reference--group-002.md#canonical-b2b612f4b7c9002adb7a18a5375f6bb5ffa3a1e305808990249e1fa822b1bf93) |
| `cloudfront.protected_endpoints.domain.exact_value` | [cloudfront.protected_endpoints.domain.exact_value](data-sources--protected_application--reference--group-003.md#canonical-ae7ff992545d343aebe6778094b2c91057c9f12a6bcdc7c19e167ac5c56950fc) |
| `cloudfront.protected_endpoints.domain.regex_value` | [cloudfront.protected_endpoints.domain.regex_value](data-sources--protected_application--reference--group-003.md#canonical-922007086cd6ab4393ae21d589d7cdd9cf251e9c2c2c5d4c955f68be808f6ef8) |
| `cloudfront.protected_endpoints.domain.suffix_value` | [cloudfront.protected_endpoints.domain.suffix_value](data-sources--protected_application--reference--group-003.md#canonical-d0f631e795bdbe877b66ce34dc5e342ace8c9ded8cf212cbd8b57182273dca53) |
| `cloudfront.protected_endpoints.flow_label` | [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-883f954fa2a9f6c89d7e69fb2a41235f8f386662b78ba726f427dd4ac7aa4023) |
| `cloudfront.protected_endpoints.flow_label.account_management` | [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-e0141c194617023468bc7c89e1ab3c76e503a7fde1d813686ac1118ba8ee9854) |
| `cloudfront.protected_endpoints.flow_label.account_management.create` | [cloudfront.protected_endpoints.flow_label.account_management.create](data-sources--protected_application--reference--group-003.md#canonical-6197666565d2ca11a8a7398abd5576a568cada100b933657a64c7a8fe2459d18) |
| `cloudfront.protected_endpoints.flow_label.account_management.password_reset` | [cloudfront.protected_endpoints.flow_label.account_management.password_reset](data-sources--protected_application--reference--group-003.md#canonical-c908cd8075655561a494b1c9df3d2db2d7765dab3b2c3418a2e0e9eb4022d72e) |
| `cloudfront.protected_endpoints.flow_label.authentication` | [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1f27742923d590ed1bd11ed4cfffe6ef8937b5be1eb8d0939377d3100824f07c) |
| `cloudfront.protected_endpoints.flow_label.authentication.login` | [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-464df341b620fbed08e03dc2c827ad2822acc63f219aba91269467ea8a5fcd1a) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--protected_application--reference--group-003.md#canonical-1a8250c21c9baef1526bd57f4a6f5d82ff4362ce994b1dda948fd62f2022cbad) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-a654a7bea255a8d5d1b7b3f8480abe73d554969ff3d31f47c75916c6b4009e9e) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](data-sources--protected_application--reference--group-003.md#canonical-f4a2f2b198bbf3d5ab9d76ad1fe1ca2cfffda35b6ac239e78a1c1ae52743c815) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name](data-sources--protected_application--reference--group-003.md#canonical-38b0d0b3aa7ac998b832774b782aaa8bb18abc35a892d529c42cc63ad65a4762) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values](data-sources--protected_application--reference--group-003.md#canonical-638c64ee4cabac18e0bd322b988edbfbdb8a90fb2eda0729a485ab1437dc140c) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status](data-sources--protected_application--reference--group-003.md#canonical-abd850b442e540ecf7cf1b78434a2dd6f7734e99f3af7cf33dcacbc6b517da98) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](data-sources--protected_application--reference--group-003.md#canonical-c6a4fb6019cb14454f6be6cb051481590c31ef833c85e8b682f0cb5db7dae632) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name](data-sources--protected_application--reference--group-003.md#canonical-c3683e548dbc68a3c904b28e30a3c0138a85bcd469b926c4bd9ff247e1687cf5) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values](data-sources--protected_application--reference--group-003.md#canonical-059ffdd8aca654751c5659d3fd2be063ba25a5719da4758e6592b7a9b93fff78) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status](data-sources--protected_application--reference--group-003.md#canonical-83980c7397239f6665d78c52beccc64a81c96bf099233840a13c64893b7c5fae) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_mfa` | [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](data-sources--protected_application--reference--group-003.md#canonical-e83d761391be4ce942bb16f6392dbbbb1919473f2b70eeebac2d68aeea3fe97d) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_partner` | [cloudfront.protected_endpoints.flow_label.authentication.login_partner](data-sources--protected_application--reference--group-003.md#canonical-322c6bc7e900b8048d13c22e42e2890b4381c78b5d79525897f0e25438f6ee41) |
| `cloudfront.protected_endpoints.flow_label.authentication.logout` | [cloudfront.protected_endpoints.flow_label.authentication.logout](data-sources--protected_application--reference--group-003.md#canonical-42dc2d513d755c7b1fbb70f6aad2a45b0b4738e51f4c71127c142eec52565cf1) |
| `cloudfront.protected_endpoints.flow_label.authentication.token_refresh` | [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](data-sources--protected_application--reference--group-003.md#canonical-b1b5693effe5abf629f3bf60eff031c09eff0f6af8f6d1d0998fb63d28a30342) |
| `cloudfront.protected_endpoints.flow_label.financial_services` | [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-aa2e51b746515ebd61cb823548072da32c7883bd2d2655239f8355b350bcd737) |
| `cloudfront.protected_endpoints.flow_label.financial_services.apply` | [cloudfront.protected_endpoints.flow_label.financial_services.apply](data-sources--protected_application--reference--group-003.md#canonical-5e489ca71ae4c2da7c25320cf293a7bc3f836c3f06cb124cf3db9f045be01c92) |
| `cloudfront.protected_endpoints.flow_label.financial_services.money_transfer` | [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](data-sources--protected_application--reference--group-003.md#canonical-f2875c697c098aa377837b07cc2d8a01028c1a451040980c6542a54295bb9d09) |
| `cloudfront.protected_endpoints.flow_label.flight` | [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-6e6b2bfc27fa95ddb845c400568c53676a66649c3c5c1b68de12471e51d18c2f) |
| `cloudfront.protected_endpoints.flow_label.flight.checkin` | [cloudfront.protected_endpoints.flow_label.flight.checkin](data-sources--protected_application--reference--group-003.md#canonical-e20c01147f4214772185dd2a0f6c2817ab0a523d844f64f215da94f1d2a05e3a) |
| `cloudfront.protected_endpoints.flow_label.profile_management` | [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-b35149c42e05f472ff4405a8188b28fdb034afa5bff913c21d8d1078b8d4216e) |
| `cloudfront.protected_endpoints.flow_label.profile_management.create` | [cloudfront.protected_endpoints.flow_label.profile_management.create](data-sources--protected_application--reference--group-003.md#canonical-cfb910dc1c64a4b7c3be4fdafaeec819cacf5709eb297ae4fedc87b1596458fe) |
| `cloudfront.protected_endpoints.flow_label.profile_management.update` | [cloudfront.protected_endpoints.flow_label.profile_management.update](data-sources--protected_application--reference--group-003.md#canonical-8b08b61a39ed216409d3c11e7905c9bf7b54caaf2e9f455b38ee7f449792b0d2) |
| `cloudfront.protected_endpoints.flow_label.profile_management.view` | [cloudfront.protected_endpoints.flow_label.profile_management.view](data-sources--protected_application--reference--group-003.md#canonical-0597fdc9f02ecef57bb4b875d113670f4cb7e953666384d5663e0b81c3ba7093) |
| `cloudfront.protected_endpoints.flow_label.search` | [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-ac6949da5461d381ef9164f087327928840b79a9b5bdeaf48c399b08c11e298b) |
| `cloudfront.protected_endpoints.flow_label.search.flight_search` | [cloudfront.protected_endpoints.flow_label.search.flight_search](data-sources--protected_application--reference--group-003.md#canonical-f83b5130562a452c2d4bc6ca0880565a692d9cc7c49e70aa826d34230fcbb467) |
| `cloudfront.protected_endpoints.flow_label.search.product_search` | [cloudfront.protected_endpoints.flow_label.search.product_search](data-sources--protected_application--reference--group-003.md#canonical-86dadaece91ab3eb51f530d34a9541ae63be8f795db62e96c0bb7cdceb1f7f81) |
| `cloudfront.protected_endpoints.flow_label.search.reservation_search` | [cloudfront.protected_endpoints.flow_label.search.reservation_search](data-sources--protected_application--reference--group-003.md#canonical-d52676fb7ddcb3baeae9a929385c7b72d5ec0964ac594bd836307f5750caec76) |
| `cloudfront.protected_endpoints.flow_label.search.room_search` | [cloudfront.protected_endpoints.flow_label.search.room_search](data-sources--protected_application--reference--group-003.md#canonical-8a56a68d4beb4547e2777a4ce603759e30038fb5216966f38f5d99e377526773) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-61a1490e2aaa8f28b8fec43669981cf45f641be0ef8dff43ce74e853919b7544) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](data-sources--protected_application--reference--group-003.md#canonical-d2dad06badf63953db7ba722d5ac344dc821d8e59f09a59bc38b803ef165c7c6) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](data-sources--protected_application--reference--group-003.md#canonical-9c8524837626cff600660b1ff1edd82531451a35b0367dc1cef11198d4caa330) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](data-sources--protected_application--reference--group-003.md#canonical-13158fcf2d383b852fe24bb064399af4642afd17967288a75af323217741ec72) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](data-sources--protected_application--reference--group-003.md#canonical-1d5aaaf9e8665e61911441bd1722764c46e0d81fd41c2989da3164af8100e169) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](data-sources--protected_application--reference--group-003.md#canonical-4c29775288e5ad20377fc86afc15019e0b680dcf07a3dc5b4f1b256a1c927f15) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](data-sources--protected_application--reference--group-003.md#canonical-2c47a3a877286ed8705961b0662336f8bc36d3cdab5c752452a6fc69bc71cebe) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](data-sources--protected_application--reference--group-003.md#canonical-66e6293e57a0bf41fb9aa4e690b7dc8ac103458132183fc81496ec5a72a210a8) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](data-sources--protected_application--reference--group-003.md#canonical-3a88340827e4e858bf075078189c8f064e32606401206d4655e6cc2a536c22b8) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](data-sources--protected_application--reference--group-003.md#canonical-de7a35d7ab5f151498def5fc023797e051fdd138e6d531782aafc2e84af42be4) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](data-sources--protected_application--reference--group-003.md#canonical-64da04ae6ff76e58057624039a8f3fd5b80c60f10374434a119dec27b48b9bb2) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](data-sources--protected_application--reference--group-003.md#canonical-ca1ba5fe987f4d28cb0d58d49c89db3f8285621a1e1e53fd74a67e31962d3876) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](data-sources--protected_application--reference--group-003.md#canonical-2721103d17b010921d8151fb2c0ebda589d2e6480af7c55457536893c3da84fd) |
| `cloudfront.protected_endpoints.http_methods` | [cloudfront.protected_endpoints.http_methods](data-sources--protected_application--reference--group-002.md#canonical-f38f68cd18ce9af63d327e3c7cd1a9d52e9cd5bc9201ac7c7945093ec017a7ee) |
| `cloudfront.protected_endpoints.metadata` | [cloudfront.protected_endpoints.metadata](data-sources--protected_application--reference--group-003.md#canonical-5405021cd33c592ad0ff76d3fafdb26f39f8798d114e2f1c59bd531bbc0450ba) |
| `cloudfront.protected_endpoints.metadata.description_spec` | [cloudfront.protected_endpoints.metadata.description_spec](data-sources--protected_application--reference--group-003.md#canonical-0cb739c676bc6fa512e066079f94ad71ecbf538cc7b28162fd015fa4a49afaba) |
| `cloudfront.protected_endpoints.metadata.name` | [cloudfront.protected_endpoints.metadata.name](data-sources--protected_application--reference--group-003.md#canonical-c9dde01e59931b906aa510c09e13935c44e57e57390190c5cf603073395871ea) |
| `cloudfront.protected_endpoints.mobile_client` | [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0659ed68d7590abeb654a95a10740d0b4c19ace0614b96c3f403485aeac5242e) |
| `cloudfront.protected_endpoints.mobile_client.block` | [cloudfront.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-003.md#canonical-345dbf3c47564f8a4dd6216651a48a8726c404916eec97827a02dc82cf86f60f) |
| `cloudfront.protected_endpoints.mobile_client.block.body` | [cloudfront.protected_endpoints.mobile_client.block.body](data-sources--protected_application--reference--group-003.md#canonical-2654cf1701eef24bb3ad724d6ecea17fc4889186b2a641f0689834f1d71a6671) |
| `cloudfront.protected_endpoints.mobile_client.block.content_type` | [cloudfront.protected_endpoints.mobile_client.block.content_type](data-sources--protected_application--reference--group-003.md#canonical-91fbb745c401c5b30ab351a739bd0c082a09849608fe6f26b3836312f57fbadb) |
| `cloudfront.protected_endpoints.mobile_client.block.status` | [cloudfront.protected_endpoints.mobile_client.block.status](data-sources--protected_application--reference--group-003.md#canonical-cb901675392f067fafe2effef9a3f9d41938c3c63f5846e1360654450357e84e) |
| `cloudfront.protected_endpoints.mobile_client.continue` | [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-a8798c5dcd12bb07048698624756cf76491d09124bfaf7cbc791f7d48b7d9be7) |
| `cloudfront.protected_endpoints.mobile_client.continue.add_header` | [cloudfront.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-003.md#canonical-b4ef99d56c6718413b4b1816d50c31d1d9ce63c32c9acf85109ec2bcea3b70e9) |
| `cloudfront.protected_endpoints.mobile_client.continue.no_header` | [cloudfront.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-003.md#canonical-e3b924cbc32b6e2e3930d19dfbcac1e780f352f89f6f9c3db51a05b4a04e4852) |
| `cloudfront.protected_endpoints.path` | [cloudfront.protected_endpoints.path](data-sources--protected_application--reference--group-002.md#canonical-9a1919d7a9860cd29a148be229f9fa3d24b63e419b4f0a5ddbe87f64275557be) |
| `cloudfront.protected_endpoints.query` | [cloudfront.protected_endpoints.query](data-sources--protected_application--reference--group-002.md#canonical-cebdb8166a09e3883df2a7707d5e0987dfc6fe580cefcd2456382c962d0ee851) |
| `cloudfront.protected_endpoints.undefined_flow_label` | [cloudfront.protected_endpoints.undefined_flow_label](data-sources--protected_application--reference--group-003.md#canonical-5d2db764f5baadc9931c395646f500c04ea88ffea0c4806ff58885bccf450dd7) |
| `cloudfront.protected_endpoints.web_client` | [cloudfront.protected_endpoints.web_client](data-sources--protected_application--reference--group-003.md#canonical-1e1c463393e17da6849f73ae17baeb59e61fe8234427995474009bea100d9afd) |
| `cloudfront.protected_endpoints.web_client.block` | [cloudfront.protected_endpoints.web_client.block](data-sources--protected_application--reference--group-003.md#canonical-097aa2b346d55f7792de416891dce767a2cc3638c68fef38436c52a2df9f9586) |
| `cloudfront.protected_endpoints.web_client.block.body` | [cloudfront.protected_endpoints.web_client.block.body](data-sources--protected_application--reference--group-003.md#canonical-3702bb34cd6e1b30465c1f859a3ba0578aaaefba09e954552fccc697bcaf9b23) |
| `cloudfront.protected_endpoints.web_client.block.content_type` | [cloudfront.protected_endpoints.web_client.block.content_type](data-sources--protected_application--reference--group-003.md#canonical-366fcc37dddc44c60826b4ec30301c340d1d9ddaecd0358c0946dfe463da7122) |
| `cloudfront.protected_endpoints.web_client.block.status` | [cloudfront.protected_endpoints.web_client.block.status](data-sources--protected_application--reference--group-003.md#canonical-59dcfee5a792b2908d53b12b1f89a8c434569284d5aea07e6318f1656f4647a2) |
| `cloudfront.protected_endpoints.web_client.continue` | [cloudfront.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-004.md#canonical-ae481a813648f84eab2bd2ad9d14c5c3ff9451b148431c348a9bf585d19157ce) |
| `cloudfront.protected_endpoints.web_client.continue.add_header` | [cloudfront.protected_endpoints.web_client.continue.add_header](data-sources--protected_application--reference--group-004.md#canonical-f6bd119999a3d056c16452db0b6d54c8c85691cf630189e7728a38acaa93db4a) |
| `cloudfront.protected_endpoints.web_client.continue.no_header` | [cloudfront.protected_endpoints.web_client.continue.no_header](data-sources--protected_application--reference--group-004.md#canonical-c2cb58a363e85f1c5df06b022e241bdda315e6bd740c17663e5d93dce1d540a3) |
| `cloudfront.protected_endpoints.web_client.redirect` | [cloudfront.protected_endpoints.web_client.redirect](data-sources--protected_application--reference--group-004.md#canonical-b7edf2bbed4b9883a8bdec8b6932ce3dcf18a805e6c925d46727da82f38e3882) |
| `cloudfront.protected_endpoints.web_client.redirect.location` | [cloudfront.protected_endpoints.web_client.redirect.location](data-sources--protected_application--reference--group-004.md#canonical-fef0f40ea02542c50bc0b80408bcf0812f0e3c6fc81954a5873392e807a3b1b2) |
| `cloudfront.protected_endpoints.web_client.redirect.status` | [cloudfront.protected_endpoints.web_client.redirect.status](data-sources--protected_application--reference--group-004.md#canonical-421cb19b920728600b829c76e7520bdf62a53c0a1d1f0665744dfc17af87a9a6) |
| `cloudfront.protected_endpoints.web_mobile_client` | [cloudfront.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-004.md#canonical-2427779dd4ee05d2a9d4efc08ae5f98e1a08afcdf1dceaa1fce1f487ca4c3109) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--reference--group-004.md#canonical-c10c034301d531b424e06208b8c880ccd4b8a83d2a78d11e79bf2b69a43f103f) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.body](data-sources--protected_application--reference--group-004.md#canonical-99106d13ebe432dbacfa11965754665aae1febbd15edc1f3099078d4b410be1d) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type](data-sources--protected_application--reference--group-004.md#canonical-dee064cd542bb2cf0b81087ad642545da8331192a2ab0643f5ace3181c081699) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.status](data-sources--protected_application--reference--group-004.md#canonical-d107ee8378ddb483c97c341c34d4c0b30d89e9e15917e51fb983186dc5ac8c08) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web` | [cloudfront.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--reference--group-004.md#canonical-1cccd90b5ef09d67f5e05e5b20e4cac17dfdfc74c726fbfed89955b509d5e0a2) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.body` | [cloudfront.protected_endpoints.web_mobile_client.block_web.body](data-sources--protected_application--reference--group-004.md#canonical-7548f72cdbbb96c3978462f06e502b90cd8cbd30ac08c4ee948dff76052f30bb) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_web.content_type](data-sources--protected_application--reference--group-004.md#canonical-ab50835175b821ba0c2b0d4762f1ac9f6470394bdd51df22b29d0e8f5ec34fce) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.status` | [cloudfront.protected_endpoints.web_mobile_client.block_web.status](data-sources--protected_application--reference--group-004.md#canonical-e3c3197fcafe3d0bfbf076793d56d10ed07d63b2b08faa1434559ade5d930d17) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-004.md#canonical-f5d12eb2efb3fbb53ce5e775320341ac4b380aa30809c8f29553fc824b958b35) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header](data-sources--protected_application--reference--group-004.md#canonical-3684d192b3724f6394369cd5a44adf917b8d2b390e8b6c39ce693ab39c8c6f58) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header](data-sources--protected_application--reference--group-004.md#canonical-11d17202ec6412d80a005d9d1449267e3aa7631beb089c8512e7a86b63ded416) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web` | [cloudfront.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-004.md#canonical-bff03ffdd43a0c53bb01b314fa90106a723ce8bb7a0bfbe916d710e3334e3733) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header](data-sources--protected_application--reference--group-004.md#canonical-6419619662bbc8e30fff1792a5e1bfe29c3219d31a92219ebd397383af2982c0) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header](data-sources--protected_application--reference--group-004.md#canonical-f9a308bc0c6014efa19144261d8267d22ca9c2da0aa47ed49ad7261936b72d05) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--reference--group-004.md#canonical-c9abae19a6ba156d9beec87e27349f3acc912efdaf82a8b186dc048545642992) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.location](data-sources--protected_application--reference--group-004.md#canonical-d8c08ec64e206a70a62bead1d41e9789081a48d91be13d6bae2f31754ca0de07) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.status](data-sources--protected_application--reference--group-004.md#canonical-715d47f2325005319f1bf3311c802fef8a721c569287998cdc59b362eebbd601) |
| `cloudfront.timeout` | [cloudfront.timeout](data-sources--protected_application--reference--group-002.md#canonical-71b05a4526574b2d94f6e8635f1401b5389ce9105efbc2e299f0127a034b5507) |
| `cloudfront.trusted_clients` | [cloudfront.trusted_clients](data-sources--protected_application--reference--group-004.md#canonical-5c129facbbeecd766988282fcf68ef70958b4ff9525ff043fd4a8ff334f505f2) |
| `cloudfront.trusted_clients.http_header` | [cloudfront.trusted_clients.http_header](data-sources--protected_application--reference--group-004.md#canonical-044a8fa385607e81a2189f7b25fe1e52c369ce8af46e69fa222fdef96cc0cc79) |
| `cloudfront.trusted_clients.http_header.headers` | [cloudfront.trusted_clients.http_header.headers](data-sources--protected_application--reference--group-004.md#canonical-2e4ed04dfec2931ccb2ce22bd781204277fe16ec1afdb88df83eec384681977b) |
| `cloudfront.trusted_clients.http_header.headers.exact` | [cloudfront.trusted_clients.http_header.headers.exact](data-sources--protected_application--reference--group-004.md#canonical-14facd3f1527189607a26194980cc1c4b19b67e7f2be18caa83ecd7179faef5e) |
| `cloudfront.trusted_clients.http_header.headers.name` | [cloudfront.trusted_clients.http_header.headers.name](data-sources--protected_application--reference--group-004.md#canonical-8e0120e129c106820aa91400a5f644a1179a2ff8fbcd7e8561f4e3b5a1b8a52c) |
| `cloudfront.trusted_clients.http_header.headers.regex` | [cloudfront.trusted_clients.http_header.headers.regex](data-sources--protected_application--reference--group-004.md#canonical-ac1933faeac60116447ba6f4f93cfe165ef56a615d3bec3b4c3dc51117678353) |
| `cloudfront.trusted_clients.ip_prefix` | [cloudfront.trusted_clients.ip_prefix](data-sources--protected_application--reference--group-004.md#canonical-2bca50f016e9cbaaa363526c15b3115eab71be31501268448afc7642c65c38f2) |
| `cloudfront.trusted_clients.metadata` | [cloudfront.trusted_clients.metadata](data-sources--protected_application--reference--group-004.md#canonical-0489e58841e00c392a9c519b7aecf6928a39f30154288b8279ba972f6f3eabfa) |
| `cloudfront.trusted_clients.metadata.description_spec` | [cloudfront.trusted_clients.metadata.description_spec](data-sources--protected_application--reference--group-004.md#canonical-edcb5632ce40c8dd4063270474d77d7a4e3fc23e7e753cd60fafd47986672271) |
| `cloudfront.trusted_clients.metadata.name` | [cloudfront.trusted_clients.metadata.name](data-sources--protected_application--reference--group-004.md#canonical-276372877f37e101f0e4d4f644ae5b2f845690b5f6f379a25cdd1172e8a888bb) |
| `custom_connector` | [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-92f3a3a5136fe39cb234315cd2391ba48d23c21e40cc2de12899c1146be4da19) |
| `description` | [description](data-sources--protected_application--reference--group-001.md#canonical-f6fc69e53bb386d3206759f5e059b4f9a1ae3eb72d5598eeb79374af95c7fa3f) |
| `f5_big_ip` | [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-3eca30b37bb1b70a27409289ace2c65b39eb6eccc7fa19a06bfedf577d9ca1b4) |
| `id` | [id](data-sources--protected_application--reference--group-001.md#canonical-95054b0d91fe8ed7e509a293def72a4265b577cafd8bf518e8e6ff21e5f6e52f) |
| `labels` | [labels](data-sources--protected_application--reference--group-001.md#canonical-9eb30875d5d3c0aad8b10e894ace27531bc4c582bba0c1f95bbd3ff70366353c) |
| `name` | [name](data-sources--protected_application--reference--group-001.md#canonical-0b6dbb265397a55b0418e5a7bb0b55f098c9e81adfbefdea4a09581241be5da9) |
| `namespace` | [namespace](data-sources--protected_application--reference--group-001.md#canonical-744d246b2c27b9b9df9ac66a7b7460ba6c1e010b6db4338f1b243463d00642ff) |
| `region` | [region](data-sources--protected_application--reference--group-001.md#canonical-044811827b8d59d08ea0aea8f8a16827dede567e0d076b604d381dfaffd1571a) |
| `salesforce_commerce_connector` | [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-96d4494e546e5a8b45ea34e894ce261208b12830b63e726037b8cc904db6381d) |

<a id="canonical-bd15445060c4a2b11d27433f94aa7088d02245b5229782a3619b3ae66f58d709"></a>

## Next pages — Property reference / 0f3a7fc24589 / 12

- [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-9d682ea065487da43edc1ea5e3d025aa52b3425b72cdc1e324d288ee618f49e3)
- [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-d99334616dce20db6de192b4c1846fe58475d4fb2cf375c0ef2dd7a5fe8c22c0)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-5b21b41a1814c9565dadc62c0ab6a5ca595ed279fbcde9e9276830e2b5eea074)
- [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-9d2981ecfbe9d0b8d6732e231eb0af76e2721080347b0d3c532913f663fd36da)
- [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-a9e6f84ab6c0e0483f50d3fd1f9d0e2837df1e409dc26e4065cb0801b072cd71)
- [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-f04c990851b7f2ef626b63087a3226d19a94d9cc23bebf2bfd39037ad13c276a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-9d682ea065487da43edc1ea5e3d025aa52b3425b72cdc1e324d288ee618f49e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfa9dbcccf644be7d632aa34da7b93fd6d79521e7e23db8948068392e5f32bb1"></a>

## adobe_commerce_connector — adobe_commerce_connector / e0a69f9f5f6a / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- adobe_commerce_connector

<a id="canonical-05d44ea52d6e9a0ea9b12220784d2aa7d2de2fed92a0e90256b9b4c4c4f7fc75"></a>

Type: `["object", {}]`. Computed.

\[OneOf: adobe\_commerce\_connector, big\_ip\_iapp, cloudflare, cloudfront, custom\_connector,
f5\_big\_ip, salesforce\_commerce\_connector\] Configuration parameter for adobe commerce connector.

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

- [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-05d44ea52d6e9a0ea9b12220784d2aa7d2de2fed92a0e90256b9b4c4c4f7fc75)
- [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-b41324199086a3d71911ffeafdc42fb683d57222a1b5f896f0d313f812bace4f)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-cb54f9a5556bc8b1065d3f18d6c059da55fb4e52914ad0a2f8fcbed50b51b501)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-41b247837668c727fb5f889326104b14970f4fbba0b3be090903fb9277c9aa42)
- [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-92f3a3a5136fe39cb234315cd2391ba48d23c21e40cc2de12899c1146be4da19)
- [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-3eca30b37bb1b70a27409289ace2c65b39eb6eccc7fa19a06bfedf577d9ca1b4)
- [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-96d4494e546e5a8b45ea34e894ce261208b12830b63e726037b8cc904db6381d)

Select alternatives according to the provider validators above.

<a id="canonical-b91cd3de479743acec041afca7a447114617dca22bba82f8e5813b901eaf3f9c"></a>

## Direct properties — adobe_commerce_connector / e0a69f9f5f6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6068402fbe017c742d5f47ff06c3f82d5bc39ec3c70791fd0689ea52335997c7"></a>

## Next pages — adobe_commerce_connector / e0a69f9f5f6a / 4

- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-d99334616dce20db6de192b4c1846fe58475d4fb2cf375c0ef2dd7a5fe8c22c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b399a7977a388429d6a9fd97dc5b46da7b37bb250086a1faa89730bd39479ee"></a>

## big_ip_iapp — big_ip_iapp / 72cc5166d7f7 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- big_ip_iapp

<a id="canonical-b41324199086a3d71911ffeafdc42fb683d57222a1b5f896f0d313f812bace4f"></a>

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

<a id="canonical-30db0ddacacc7c86b7b681150727d95c8190caa4c93588eedd43f05fcb8e1e70"></a>

## Direct properties — big_ip_iapp / 72cc5166d7f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-011ac363711b46cbf372ec11fd4ab13f7e93ba5e063543a2e24e53d099eb594b"></a>

## Next pages — big_ip_iapp / 72cc5166d7f7 / 4

- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0f6f878c8bda7b1580cc965564c8a454e05c31a811ffba44a1f099de6f62d17"></a>

## cloudflare — cloudflare / a8fa1d90b067 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- cloudflare

<a id="canonical-cb54f9a5556bc8b1065d3f18d6c059da55fb4e52914ad0a2f8fcbed50b51b501"></a>

Type: `"single"`. Computed.

Bot Defense policy configuration for Cloudflare.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

<a id="canonical-738302a85ac28f82fdbc66e30ac9d0fd955e8b8c61195aba7409d5c4ea465324"></a>

## Direct properties — cloudflare / a8fa1d90b067 / 3

<a id="canonical-a283efd73921dd0b8ad59198064b74cee4001f9bfdcfc4b048aaca72d785edfb"></a>

<a id="canonical-53586aeba2316b25989e4a4c1eaa3d4f726e4ced1d09fb2cb04216fb680e436a"></a>

## continue_mitigation_action_hdr property — cloudflare / a8fa1d90b067 / 4

Type: `"string"`. Computed.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [disable_js_insert](data-sources--protected_application--reference--group-001.md#canonical-2800a972dcea40694121868298086458e9088c7c07f172903484158c17e79351): complete subsection reference.

- [disable_mobile_sdk](data-sources--protected_application--reference--group-001.md#canonical-47a8e6dcd64752cef2a6301eb3f94160159651fb26cf9335d5db18e7d79ee0d2): complete subsection reference.

- [js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6): complete subsection reference.

<a id="canonical-aeea0c3086c2f4da5f5ead949e034be0efa0eae2fa4331c0dcfb6f3a423e7d56"></a>

<a id="canonical-a98332cc1155211594bf1bb2b4db03ca53d4f0bbea20b683fee42230d34fb1f6"></a>

## loglevel property — cloudflare / a8fa1d90b067 / 5

Type: `"string"`. Computed.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

Upstream description:

Select the level of logging desired. Levels are cumulative (e.g. Debug includes Error, Warning, and
Informational)

&#8203;- LOG\_UNDEFINED: Undefined

&#8203;- LOG\_ERROR: Error

Log only errors &#8203;- LOG\_WARNING: Warning

Log malicious requests &#8203;- LOG\_INFO: Info

Log all requests &#8203;- LOG\_DEBUG: Debug

Log debugging data.

Receipt-pinned upstream constraints:

```json
{
  "default": "LOG_UNDEFINED",
  "enum": [
    "LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [manual_js_insert](data-sources--protected_application--reference--group-001.md#canonical-178e9de5d6de69bfb8d329aca7439ec74c627ec7324dbb0b9b197898f2b515ee): complete subsection reference.

- [mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-46f2f356aae713cdf2bdd290be3b00e2b80d5af30dece59c88f1517f93f9546d): complete subsection reference.

- [protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a): complete subsection reference.

<a id="canonical-6aa70ad8d77f28876c3a67584ebd088fa809c218b48b355ede2d67c433b7a632"></a>

<a id="canonical-304d253972ede6d83a314028e642921abb7ba328ae9791bdbd42178fef61d79c"></a>

## timeout property — cloudflare / a8fa1d90b067 / 6

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8): complete subsection reference.

<a id="canonical-27a072487ea5cb0bd0d5ef27e2711c7e82859686703610071ab6e579d7245a65"></a>

## Next pages — cloudflare / a8fa1d90b067 / 7

- [cloudflare.disable_js_insert](data-sources--protected_application--reference--group-001.md#canonical-2800a972dcea40694121868298086458e9088c7c07f172903484158c17e79351)
- [cloudflare.disable_mobile_sdk](data-sources--protected_application--reference--group-001.md#canonical-47a8e6dcd64752cef2a6301eb3f94160159651fb26cf9335d5db18e7d79ee0d2)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.manual_js_insert](data-sources--protected_application--reference--group-001.md#canonical-178e9de5d6de69bfb8d329aca7439ec74c627ec7324dbb0b9b197898f2b515ee)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-46f2f356aae713cdf2bdd290be3b00e2b80d5af30dece59c88f1517f93f9546d)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-7bbe474329711bf4a1b6de708dcce03ec8d0f1ce8a041b030989bc9841b46aa8)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-2800a972dcea40694121868298086458e9088c7c07f172903484158c17e79351"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29e229a2268222ae746d024f1038a90fae6922551fce3b53e031fefb3bd5733f"></a>

## cloudflare.disable_js_insert — cloudflare.disable_js_insert / ee3750f94c2c / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- cloudflare.disable_js_insert

<a id="canonical-1b91afdb539e1ab18e2ef6648838b17f0b109895489d8994ddef75ced0357940"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

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

<a id="canonical-65d123f53059a8f73280300a5595fc227b4bd7c5e09bb70a0e82a7560c7a1c08"></a>

## Direct properties — cloudflare.disable_js_insert / ee3750f94c2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e9a3ad940d24f3e2f5a28fa5fc98f0015afdb330099744bba5a685179540ce3"></a>

## Next pages — cloudflare.disable_js_insert / ee3750f94c2c / 4

- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-47a8e6dcd64752cef2a6301eb3f94160159651fb26cf9335d5db18e7d79ee0d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-006e19b0efc7af2a0c1dfe2e54350299e8fcc303f57ba5a541282ab09dd6da2f"></a>

## cloudflare.disable_mobile_sdk — cloudflare.disable_mobile_sdk / 797d70f717d4 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- cloudflare.disable_mobile_sdk

<a id="canonical-18b083946f720fe56ba43de935374a8750643fb44a02e25384f4ced09fcba7b0"></a>

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

<a id="canonical-11f1bc32ec3706f8868a3fd07fbffe5c34363a54656a5c51330fcad83a5d57bb"></a>

## Direct properties — cloudflare.disable_mobile_sdk / 797d70f717d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-effd9dd5b6918778c752e89d442d87c8dc8772c31929b3f8bcbb8d2f26879a65"></a>

## Next pages — cloudflare.disable_mobile_sdk / 797d70f717d4 / 4

- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9605828cb35b7829a36d25fe264443008767da05253447d56bbfed4ee5de1e58"></a>

## cloudflare.js_insertion_rules — cloudflare.js_insertion_rules / 2f55df086705 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- cloudflare.js_insertion_rules

<a id="canonical-cf726e6368994bdfe7b8ebe402190ecb3cd9dad8d19d888dd1dffbb54f029563"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-25f98e595754757c4fa950f49dcdfdde7c74abc84a6cfe09c0f47631ed8fd96e"></a>

## Direct properties — cloudflare.js_insertion_rules / 2f55df086705 / 3

- [exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf): complete subsection reference.

<a id="canonical-b2ec916d86b0823c2126f1598d0dbe6aee00f76ce585df2a6c482a0632da35d3"></a>

<a id="canonical-f52a2536ea3a11c285602056ca69b049e644a62ea85b2dbebffca31232fb7fdd"></a>

## javascript_location property — cloudflare.js_insertion_rules / 2f55df086705 / 4

Type: `"string"`. Computed.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

Upstream description:

All inside networks.

&#8203;- JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED

Undefined Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag.
Insert JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "JAVA_SCRIPT_LOCATION_UNDEFINED",
  "enum": [
    "JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-44e1ad65d7ebf434b92e72b4bb62ee00e0388d8693e70eb377e5b478a680ffab"></a>

<a id="canonical-90785d9bcd81e5e067a3db26257784e9d34ea55143e8787221b30797ea644f28"></a>

## js_download_path property — cloudflare.js_insertion_rules / 2f55df086705 / 5

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/common.js’.

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39): complete subsection reference.

<a id="canonical-2f5f557083538708cd7092464a130483c7cca6ff00c52297106e5c41af9e8f20"></a>

## Next pages — cloudflare.js_insertion_rules / 2f55df086705 / 6

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fde2397d5e3123f146ecc409d79b3a41dd77494cb0a23f4db063c14185a71da"></a>

## cloudflare.js_insertion_rules.exclude_list — cloudflare.js_insertion_rules.exclude_list / 8b65d28f5d00 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- cloudflare.js_insertion_rules.exclude_list

<a id="canonical-9120827327b2d0f0d431cce3aa97acaa52d672f9e3cca51575f6170bccaad263"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-00b56489c4f5682300b40debfb6aa82f9448560c8ce612381ca386307fa0ff25"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list / 8b65d28f5d00 / 3

- [any_domain](data-sources--protected_application--reference--group-001.md#canonical-aa0056b99a076fb7c5fedd374fc053f0bbfa7d95f03e3a72c876db669f923600): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-001.md#canonical-b4a19c2760d443006dec6ec78698c5ff8a9fd7394c6609ab0acb4bc16f3eb8b9): complete subsection reference.

- [metadata](data-sources--protected_application--reference--group-001.md#canonical-8757b5f15ec9517bf1e5e740b3ca543abbd63fd5fb190707be52a0655428828b): complete subsection reference.

- [path](data-sources--protected_application--reference--group-001.md#canonical-9ada4be0df65de090d6aaa63c831ae72f10ea5bf5379cb35fffb123893134554): complete subsection reference.

<a id="canonical-07e3d9c825234945844c375a8dd6028dfd94c832e14c99181ef5c760489d1ffb"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list / 8b65d28f5d00 / 4

- [cloudflare.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-001.md#canonical-aa0056b99a076fb7c5fedd374fc053f0bbfa7d95f03e3a72c876db669f923600)
- [cloudflare.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-001.md#canonical-b4a19c2760d443006dec6ec78698c5ff8a9fd7394c6609ab0acb4bc16f3eb8b9)
- [cloudflare.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-001.md#canonical-8757b5f15ec9517bf1e5e740b3ca543abbd63fd5fb190707be52a0655428828b)
- [cloudflare.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-001.md#canonical-9ada4be0df65de090d6aaa63c831ae72f10ea5bf5379cb35fffb123893134554)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-aa0056b99a076fb7c5fedd374fc053f0bbfa7d95f03e3a72c876db669f923600"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a778b0730fbe19ba32ccef9da33dbf12b7edc251e429745fa7502d90644bb3b8"></a>

## cloudflare.js_insertion_rules.exclude_list.any_domain — cloudflare.js_insertion_rules.exclude_list.any_domain / efb179529e88 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- cloudflare.js_insertion_rules.exclude_list.any_domain

<a id="canonical-953602b5c5c9f4a553d6ad63a77c298b29376ac70b44f8d08a5b564a21a2bc00"></a>

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

<a id="canonical-245c71aaf298f1399d78802db658102d96ec22e859cb3d13127fae2629baa110"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.any_domain / efb179529e88 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39dbbcc5cb7170956c593f06a9260be93da10c5ef3a01134c9f6d59a011fee23"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.any_domain / efb179529e88 / 4

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-b4a19c2760d443006dec6ec78698c5ff8a9fd7394c6609ab0acb4bc16f3eb8b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd5c33d8b4b3d326867f7c4a1df8d6d60bbcdf312f80d88aaa60ba32c0a18e0e"></a>

## cloudflare.js_insertion_rules.exclude_list.domain — cloudflare.js_insertion_rules.exclude_list.domain / c4fd79778c9b / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- cloudflare.js_insertion_rules.exclude_list.domain

<a id="canonical-9b1584a5a0218fc819d0c156a65c04f53b1b58aaadee0700a925e86381053ca7"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

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

<a id="canonical-f07044135fa17310cb197bdb9785a8fa579284b809914ad652c0c2353cfa7e3d"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.domain / c4fd79778c9b / 3

<a id="canonical-8c5523888ad901a12a619f30a328d6b3b0320e24d84162650f51620d6f9166c9"></a>

<a id="canonical-6f6cb35d115f7a636d313e6153204bebe551106fad05b56418391aa3ef02d2f3"></a>

## exact_value property — cloudflare.js_insertion_rules.exclude_list.domain / c4fd79778c9b / 4

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

<a id="canonical-bf8d1ad3a802269f334d254aee3aef8f7fa84a6c0c554212302e40aec3ec0a19"></a>

<a id="canonical-79ed3c6353c1f90e1cbd29dfb6953f3b95eeb5ff348a849aa718db9e5bb37560"></a>

## regex_value property — cloudflare.js_insertion_rules.exclude_list.domain / c4fd79778c9b / 5

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

<a id="canonical-74ac002bfc060d53a1baa5fab4484aaf71b736140db9dc5e3e8adbba8afff7d0"></a>

<a id="canonical-cf44b3a89dece465e59de78468413aac8b36115d8d40b33d5ae678606386d761"></a>

## suffix_value property — cloudflare.js_insertion_rules.exclude_list.domain / c4fd79778c9b / 6

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

<a id="canonical-4835782ca44a32ffe3a3e58523fd5cde3821b93c4d65c9f3d25926696a9f61cd"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.domain / c4fd79778c9b / 7

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-8757b5f15ec9517bf1e5e740b3ca543abbd63fd5fb190707be52a0655428828b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56cec511c1116ff385328ed8b3efb8690991ef667f5b7e432304eb8b38868602"></a>

## cloudflare.js_insertion_rules.exclude_list.metadata — cloudflare.js_insertion_rules.exclude_list.metadata / aed194c9fd24 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- cloudflare.js_insertion_rules.exclude_list.metadata

<a id="canonical-432df8a928168b6462fcbd765de52d4c6a191b8178ca7286411e00f1ebb435f1"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7e4ec076ed483594b3cec8fb9f943d826ca0d0a72ca69ef482212c82e6c4653c"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.metadata / aed194c9fd24 / 3

<a id="canonical-7509691e54b5540d924013c40d78e91dd1d3e08f1ba303f1f62cf4884fbe263a"></a>

<a id="canonical-3ea5b3a63a1b5b53bd7952b14affb1ba422d3ff78d8167c8c6c39788ab78e9a8"></a>

## description_spec property — cloudflare.js_insertion_rules.exclude_list.metadata / aed194c9fd24 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-e75126b66119ea451d512bb063ef904f724e7bc932d31d35f99c57c0e6480728"></a>

<a id="canonical-e528b907aa7b2484d2dceff58475a864fd027c699648507c30c214357cbd2972"></a>

## name property — cloudflare.js_insertion_rules.exclude_list.metadata / aed194c9fd24 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-4279a18fa6c997a383da011f1286e73ceae6e411841cb84a80bc55f2f62cd8e1"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.metadata / aed194c9fd24 / 6

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-9ada4be0df65de090d6aaa63c831ae72f10ea5bf5379cb35fffb123893134554"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec914850bad035b7f9f014ba7f35a3f2bd550a3e361afcaecffa79bcd646d2ca"></a>

## cloudflare.js_insertion_rules.exclude_list.path — cloudflare.js_insertion_rules.exclude_list.path / a8433726c99b / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- cloudflare.js_insertion_rules.exclude_list.path

<a id="canonical-53ac78942289025c8d09b6394594c3f6b92eee9e963ef724d0ce5ad2463d5910"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-718b362492fb83f7cf10589ef61136d8035f4fb03f4cea3dc85fbc8b9e628eda"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.path / a8433726c99b / 3

<a id="canonical-10a1a2cf04ad370a6e645459343edfe7834d4bd6078108eab38f600e02b678dc"></a>

<a id="canonical-1ebc6e2c7a4e066000a218689d68e511b97e395bcbe70e962c282374a5250645"></a>

## path property — cloudflare.js_insertion_rules.exclude_list.path / a8433726c99b / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-374b13a718795c306d41d3fd02d66055b6e981c825ce306a76a27c2f1066b509"></a>

<a id="canonical-b3eee832887a449e310b8646b2d51e77d246b0d234c0d0362143c4e9975d42f0"></a>

## prefix property — cloudflare.js_insertion_rules.exclude_list.path / a8433726c99b / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-70d7bd03045d145eb3076a10fa03bb42abbad6c10c38fc7a0b43237382f51bba"></a>

<a id="canonical-1e56e754adc9d67fd7cb8ab6fcc312779c996828413fa256acf122ac1104d33b"></a>

## regex property — cloudflare.js_insertion_rules.exclude_list.path / a8433726c99b / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-e056908dd9a7d482f6b1a9fc7bc6d2ad0b1791429b4937212a284da1190003ca"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.path / a8433726c99b / 7

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-81fd77ff381cd339562fc9a727fecb023504892a886720d2abbb511fd0848bbf)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49a82a81639a542c8b1ed6bfd72a6825c73c1134d48bd7d8b90c85f5ca52fe5c"></a>

## cloudflare.js_insertion_rules.rules — cloudflare.js_insertion_rules.rules / ec3ac16a4710 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- cloudflare.js_insertion_rules.rules

<a id="canonical-3142f7fa3b9bf4e6cb93ff5f5ed3a938ca6ee56b3fc7800ebfb4f15e87ba8207"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b6c92f9e25a586d1fd716dd777942bd96d2d2d25c5df24145bd538283237ac2a"></a>

## Direct properties — cloudflare.js_insertion_rules.rules / ec3ac16a4710 / 3

- [any_domain](data-sources--protected_application--reference--group-001.md#canonical-c0628b49e3dc31040405ae2474552c735a9c8fb42582ca503c0fd06bdd3064fc): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-001.md#canonical-7e598f83efd8c3ae1d663abff621b1339cbdc4cdc8f6e742f76d79030481d1c4): complete subsection reference.

<a id="canonical-abd0166f86a670a232009855a32deb84081bd74feb974a8bfc7789ca9c65bfad"></a>

<a id="canonical-81dd2fec6b2919804cc60ab37df8bbd92340411ccfd43bf3c440f9f403c2a0ff"></a>

## exact_path property — cloudflare.js_insertion_rules.rules / ec3ac16a4710 / 4

Type: `"string"`. Computed.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-4797809dca93df35a48458ce19252f4589b3ee9301247b1566df4d2e0b547c75"></a>

<a id="canonical-d5d730414b70fecd2ae7fbac9e87f829ad765450bd1af28a67aedd203aaf027f"></a>

## glob property — cloudflare.js_insertion_rules.rules / ec3ac16a4710 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-001.md#canonical-ade2aef94aadafdec2f5d650c1878ea9d00b84bfb9e86499fd406999fd6c6631): complete subsection reference.

<a id="canonical-db124372f7f4fdc79aa9640fb89e56bd811433c8904987faae12e542918991d2"></a>

<a id="canonical-c257f09abca06d043acca3287b0549d6fdf7b9aa0ca7e738abcc5ec8467a58e1"></a>

## prefix property — cloudflare.js_insertion_rules.rules / ec3ac16a4710 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-e1a085f3d90b575c4f0ba633a9cbd634b6f51a308d1420bd5f5c0a8b3e12920e"></a>

## Next pages — cloudflare.js_insertion_rules.rules / ec3ac16a4710 / 7

- [cloudflare.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-001.md#canonical-c0628b49e3dc31040405ae2474552c735a9c8fb42582ca503c0fd06bdd3064fc)
- [cloudflare.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-001.md#canonical-7e598f83efd8c3ae1d663abff621b1339cbdc4cdc8f6e742f76d79030481d1c4)
- [cloudflare.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-001.md#canonical-ade2aef94aadafdec2f5d650c1878ea9d00b84bfb9e86499fd406999fd6c6631)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-c0628b49e3dc31040405ae2474552c735a9c8fb42582ca503c0fd06bdd3064fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4520b2a9d27c8ee7437c842bf239d3266226e9ca9468d89b9436e56c457c412d"></a>

## cloudflare.js_insertion_rules.rules.any_domain — cloudflare.js_insertion_rules.rules.any_domain / 71a6b0d869d7 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39)
- cloudflare.js_insertion_rules.rules.any_domain

<a id="canonical-8f2ed1547e9e98d3f7f3c5e2e0137864ca7f0650a4bb2eefae44784d8774c9f8"></a>

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

<a id="canonical-fe5b0ce402eda453bce34f7be55b9f891f67e01ae4ee1b36ba2d2505dcda1dc7"></a>

## Direct properties — cloudflare.js_insertion_rules.rules.any_domain / 71a6b0d869d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ed4a6385c6c9b8d32f0c8698731bbbc5a033069f68cb7d574d34c43ff99ae41"></a>

## Next pages — cloudflare.js_insertion_rules.rules.any_domain / 71a6b0d869d7 / 4

- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-7e598f83efd8c3ae1d663abff621b1339cbdc4cdc8f6e742f76d79030481d1c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eafa9af492dea42945a3644055bd084aec4c7f7004b640b39e4639c662ef019c"></a>

## cloudflare.js_insertion_rules.rules.domain — cloudflare.js_insertion_rules.rules.domain / 82117049efc1 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39)
- cloudflare.js_insertion_rules.rules.domain

<a id="canonical-16d797752145174f3dcaa34e88cece05fc7063f4cc00cb417fbfc822ab8ce479"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

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

<a id="canonical-655921618303e44523662ff9ead9007ac9530924bc06dcc0f6d81a3d39762864"></a>

## Direct properties — cloudflare.js_insertion_rules.rules.domain / 82117049efc1 / 3

<a id="canonical-88a6b7469d3dbb92ecaee3f2f4e47423b28a9db421fca524fe00651cd8844c31"></a>

<a id="canonical-7dd2ab9b5ca2d17c12f88593f8cd77b563033e0d3bea53ba87c154e858eda047"></a>

## exact_value property — cloudflare.js_insertion_rules.rules.domain / 82117049efc1 / 4

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

<a id="canonical-850bc0f9cd3c157f90c4bf7cba399c58201e059d3712cad20c3b988228858690"></a>

<a id="canonical-8356257c0150e6392be32ae54e995d42093941c5ee992cf4945216ea4c408111"></a>

## regex_value property — cloudflare.js_insertion_rules.rules.domain / 82117049efc1 / 5

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

<a id="canonical-b79b22fffa02d37fd50fa753be75d3d41d61f9cd64952d208e09b1d33e3d2181"></a>

<a id="canonical-ac2b5aa6f2dd68498fe10be8a9e88f7397136457bdf6c3dd1d7be84ae7062542"></a>

## suffix_value property — cloudflare.js_insertion_rules.rules.domain / 82117049efc1 / 6

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

<a id="canonical-fc7f20be2625034ec75372adad2791099bee4f907926427360b65d8ecaeb5140"></a>

## Next pages — cloudflare.js_insertion_rules.rules.domain / 82117049efc1 / 7

- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-ade2aef94aadafdec2f5d650c1878ea9d00b84bfb9e86499fd406999fd6c6631"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a42d0ce976eec6a6d87c572ae235979758bebd2a8b545813697413c1ee0e8da"></a>

## cloudflare.js_insertion_rules.rules.metadata — cloudflare.js_insertion_rules.rules.metadata / 538978d1d3a3 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-6dd6aa0f150f32f06f21035da62f2beacbbfd3667ee61c4c5a6ea373b6ff9ef6)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39)
- cloudflare.js_insertion_rules.rules.metadata

<a id="canonical-6de971da694f973c43ae7a72c24eed15d5bbdd5b98919264dd99843021c783d8"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-92bd09d342b5c4934d9f3760ffcdc9c703295de508cb8cc1c7813183dc9618ed"></a>

## Direct properties — cloudflare.js_insertion_rules.rules.metadata / 538978d1d3a3 / 3

<a id="canonical-6e87cbe51a484d4115689b7f517c8c82732ece6b7033f903d870ba57c081177a"></a>

<a id="canonical-77f058730f5d27c4fe92ef20c595fdf882d8accfe034eb7a04e5a10833bda4e1"></a>

## description_spec property — cloudflare.js_insertion_rules.rules.metadata / 538978d1d3a3 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-6ceb6700da034965f62be5542c369f33b966954dab65fcc91e7f7413ee64576d"></a>

<a id="canonical-ea0afce0d5521017e8cea1a4336366d8ffa97471d77f944a298d2ad09dbf366f"></a>

## name property — cloudflare.js_insertion_rules.rules.metadata / 538978d1d3a3 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-4e21df42c0c90199913a33f9d9b09b44caf74433d16312f14abefa87c35e7ddd"></a>

## Next pages — cloudflare.js_insertion_rules.rules.metadata / 538978d1d3a3 / 6

- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-50c88b884f0c4be6aeb1d9d3f736a92216dbb64209a605f52954f4391cb8fe39)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-178e9de5d6de69bfb8d329aca7439ec74c627ec7324dbb0b9b197898f2b515ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8753eeda679a6bfcb5d3bced717793cd74cc91c173368d8b98bba47d144d26e3"></a>

## cloudflare.manual_js_insert — cloudflare.manual_js_insert / cdea65a0257b / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- cloudflare.manual_js_insert

<a id="canonical-c45398b4d2684259668c778485de845d926bea00bc86b37b9637af70dfc77e44"></a>

Type: `"single"`. Computed.

Insert JavaScript Manually. Insert JavaScript manually.

Upstream description:

Insert JavaScript manually.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a3cce30a57dd91ec59c88125b7a36663f156dd2db952d76d01b90f01e492294d"></a>

## Direct properties — cloudflare.manual_js_insert / cdea65a0257b / 3

<a id="canonical-efd714f974ed904485b0f9c9c4443d6e786a63932687e4c0790b2a88ab71bd7d"></a>

<a id="canonical-06c9f4b0b653830b7049e854a70b759f2f945762775323aa16d4fc4fd4c5000f"></a>

## js_download_path property — cloudflare.manual_js_insert / cdea65a0257b / 4

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/common.js’.

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

<a id="canonical-fde575152496d88cbdd5eb221d93dbbb2972ec2fcd79a9629f2cfd7a45f902d7"></a>

## Next pages — cloudflare.manual_js_insert / cdea65a0257b / 5

- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-46f2f356aae713cdf2bdd290be3b00e2b80d5af30dece59c88f1517f93f9546d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-611800b92f4a95fac39175df3be1c3997e9e72143e4c9eb9a5b2de9e3fb39388"></a>

## cloudflare.mobile_sdk_config — cloudflare.mobile_sdk_config / 6632d75d75b8 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- cloudflare.mobile_sdk_config

<a id="canonical-76ec3f5ba91aeb95484bc39e05a1b2320c69322fb4f52746f36f9c6dfec13282"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-444f1df5333260f63ce1fb67de28711c415016a34ead5e413606ca6a3533abd2"></a>

## Direct properties — cloudflare.mobile_sdk_config / 6632d75d75b8 / 3

- [mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-a96b88b20a77e5a7d5b0d5818f3477de892f06b874d3bf42c9f13434121dca69): complete subsection reference.

<a id="canonical-989b3ccc4c2ef0837392db015a916fc88c14525df91478fc03ab60c504598919"></a>

## Next pages — cloudflare.mobile_sdk_config / 6632d75d75b8 / 4

- [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-a96b88b20a77e5a7d5b0d5818f3477de892f06b874d3bf42c9f13434121dca69)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-a96b88b20a77e5a7d5b0d5818f3477de892f06b874d3bf42c9f13434121dca69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91695a8d61a527e99e78e36f3f4f3565138c779604c043105842f6ff24de2768"></a>

## cloudflare.mobile_sdk_config.mobile_identifier — cloudflare.mobile_sdk_config.mobile_identifier / c9368312fbb7 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-46f2f356aae713cdf2bdd290be3b00e2b80d5af30dece59c88f1517f93f9546d)
- cloudflare.mobile_sdk_config.mobile_identifier

<a id="canonical-0c0e9f173618eec54381cd0daff5bf6e4b66edb86d291d3394cd77637aa773a6"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9efa08d0a8bd6595a345b3cbdb8ba8fcd305ec401bec984ed5354139c2f91ec2"></a>

## Direct properties — cloudflare.mobile_sdk_config.mobile_identifier / c9368312fbb7 / 3

- [headers](data-sources--protected_application--reference--group-001.md#canonical-fe7f3d023a6485e6d858365eaaa197dfdd63a7378cfe6414800a7299684626a8): complete subsection reference.

<a id="canonical-12de2624f8acb2e3f00e69da7e2b9ab9799de512478585fd6271ec34e5f31b0c"></a>

## Next pages — cloudflare.mobile_sdk_config.mobile_identifier / c9368312fbb7 / 4

- [cloudflare.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-001.md#canonical-fe7f3d023a6485e6d858365eaaa197dfdd63a7378cfe6414800a7299684626a8)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-46f2f356aae713cdf2bdd290be3b00e2b80d5af30dece59c88f1517f93f9546d)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-fe7f3d023a6485e6d858365eaaa197dfdd63a7378cfe6414800a7299684626a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-760c4284b94eb86babb5566b96c64aca36f35b5d7adf2a74829c7510f075c678"></a>

## cloudflare.mobile_sdk_config.mobile_identifier.headers — cloudflare.mobile_sdk_config.mobile_identifier.headers / 585e079954fc / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-46f2f356aae713cdf2bdd290be3b00e2b80d5af30dece59c88f1517f93f9546d)
- [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-a96b88b20a77e5a7d5b0d5818f3477de892f06b874d3bf42c9f13434121dca69)
- cloudflare.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-63f1cdd1c66d2eecdfafa1ee89d00fe52382f88f19d95fb7864602013a8c41cf"></a>

Type: `"list"`. Computed.

List of headers that can be used to identify mobile traffic.

Upstream description:

A list of headers that can be used to identify mobile traffic.

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
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1631444f3feee7f7e0a34b5fa0b36bfb876a1b8e72a790ba207fac43cfeff61a"></a>

## Direct properties — cloudflare.mobile_sdk_config.mobile_identifier.headers / 585e079954fc / 3

<a id="canonical-40383724c736d35fdd5c8a314c4ea152a6d302dc061ed94264fb48064a4fee36"></a>

<a id="canonical-46ebd6264b11d7a9813dd989be7400575fe44444f1e36d42cc8359ed163200d3"></a>

## exact property — cloudflare.mobile_sdk_config.mobile_identifier.headers / 585e079954fc / 4

Type: `"string"`. Computed.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2809e36e0abe8897a1c7410876d6f2bb8e02ae91b62304f8e21ef3a5742ce42e"></a>

<a id="canonical-a69dec9f48b1aa855e9b065d88ae3f6ae32df0c491064785056fd1e0db16abf0"></a>

## name property — cloudflare.mobile_sdk_config.mobile_identifier.headers / 585e079954fc / 5

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-215c0cd1b564aaa735611c2e8e5269e6cec4a3b69899e03fd53324f966b8d6b3"></a>

<a id="canonical-b4671113c942b0e051ff9e24f887e0095c5d5ccf6929904ba5ec0c73be5ce2b7"></a>

## regex property — cloudflare.mobile_sdk_config.mobile_identifier.headers / 585e079954fc / 6

Type: `"string"`. Computed.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-cc294f2b7379dfcb05088c6fef97cf9ddbd45910aeeba2c5e2b4dbdfbc4c2737"></a>

## Next pages — cloudflare.mobile_sdk_config.mobile_identifier.headers / 585e079954fc / 7

- [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-a96b88b20a77e5a7d5b0d5818f3477de892f06b874d3bf42c9f13434121dca69)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ec6fc3c9892e9f834b8ba930c1e18c38b44f57fb89a52a1d6938a9e04ff04ba"></a>

## cloudflare.protected_endpoints — cloudflare.protected_endpoints / 2171643a5054 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- cloudflare.protected_endpoints

<a id="canonical-b06731213efa4a56bf9b6b10c576725c84fe8dc855a7e59249f58f701b3c6240"></a>

Type: `"list"`. Computed.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-26c19bfc3a0192254087689bcb1a6e8f3aabdced883c341dbd1a685842829666"></a>

## Direct properties — cloudflare.protected_endpoints / 2171643a5054 / 3

- [any_domain](data-sources--protected_application--reference--group-001.md#canonical-8005dad8075bd8f494e220512675b9c2469ef05cc6a2f04e8cab3b14af78627d): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-001.md#canonical-0ffb441a3a820b078ced7af109cb1c93c8ad2cd0a79e4bbb265597369b2061b5): complete subsection reference.

<a id="canonical-70069403fa976335a8d6ff04601461e917b0c5ee6ad8fdd1f1b853eb056f7298"></a>

<a id="canonical-ae05ef48ede2d47d942d6c7a0cc96753f3f74f498d1fc902e4db149474cf9d75"></a>

## http_methods property — cloudflare.protected_endpoints / 2171643a5054 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-001.md#canonical-9de81935223170e1aa599825ec8823a90bc7e0ae5fcd785c1716908ecc47174f): complete subsection reference.

- [mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab): complete subsection reference.

- [path](data-sources--protected_application--reference--group-002.md#canonical-ee8706ed2b875ad7520a81346808930b223ea48f0e68b428b74abd39f4ef1069): complete subsection reference.

<a id="canonical-017a14b06058f069c022f2573ea2f88e80f13f0d2621935c84ed2ebbe31b145d"></a>

<a id="canonical-3b50364cab297b8a033c1b5f71586000cb6151ea711c1a20dc97bdef1b82fd2b"></a>

## query property — cloudflare.protected_endpoints / 2171643a5054 / 5

Type: `"string"`. Computed.

Enter a regular expression to match your query parameters of interest.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835): complete subsection reference.

- [web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210): complete subsection reference.

<a id="canonical-85d955f34f74e0f838915a9a5ac73ba092f6a75b927cdae9eeffeca593a4029f"></a>

## Next pages — cloudflare.protected_endpoints / 2171643a5054 / 6

- [cloudflare.protected_endpoints.any_domain](data-sources--protected_application--reference--group-001.md#canonical-8005dad8075bd8f494e220512675b9c2469ef05cc6a2f04e8cab3b14af78627d)
- [cloudflare.protected_endpoints.domain](data-sources--protected_application--reference--group-001.md#canonical-0ffb441a3a820b078ced7af109cb1c93c8ad2cd0a79e4bbb265597369b2061b5)
- [cloudflare.protected_endpoints.metadata](data-sources--protected_application--reference--group-001.md#canonical-9de81935223170e1aa599825ec8823a90bc7e0ae5fcd785c1716908ecc47174f)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-001.md#canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab)
- [cloudflare.protected_endpoints.path](data-sources--protected_application--reference--group-002.md#canonical-ee8706ed2b875ad7520a81346808930b223ea48f0e68b428b74abd39f4ef1069)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-eca6ded4801e7f8bc55abe484b8f36d2c0f3e5b4de2f1b56197cef4f7d805835)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-a0ec7456043b9f858643bb8ea0f04771b63fbd5e44fce0d439c7175626a0a210)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-8005dad8075bd8f494e220512675b9c2469ef05cc6a2f04e8cab3b14af78627d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0b81e06fb679fda72603c8669c29207e1218df3842d5a166043b61eeded939a"></a>

## cloudflare.protected_endpoints.any_domain — cloudflare.protected_endpoints.any_domain / 2f4ff7d8853b / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- cloudflare.protected_endpoints.any_domain

<a id="canonical-ebadff18fef6e9d651df4a7971dd0bd363fa66e6cf82e47d23cfe0fd5fbe3843"></a>

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

<a id="canonical-de4c34ed44a560d7518cdbdfec4fc17ff8244cb50ea0ba62fe46d794a223e505"></a>

## Direct properties — cloudflare.protected_endpoints.any_domain / 2f4ff7d8853b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-538a84c453c43d9b279e672d884729c90256aa866cdb78752404278caf5bdb98"></a>

## Next pages — cloudflare.protected_endpoints.any_domain / 2f4ff7d8853b / 4

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-0ffb441a3a820b078ced7af109cb1c93c8ad2cd0a79e4bbb265597369b2061b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e66d0a68f28b914f9f045eacc29db8a9b32d6bdb9b40e8dc9bcacbc8d7fd0e28"></a>

## cloudflare.protected_endpoints.domain — cloudflare.protected_endpoints.domain / d8d367298bf7 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- cloudflare.protected_endpoints.domain

<a id="canonical-2b984be80640971b67dbd8f34af829cc13d3a289cc497fe0e77f390deeb7803c"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

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

<a id="canonical-dc5bbc61306ce379fce3b9f13196e4221340e0b93e276428a97f2ae7c2b5d3db"></a>

## Direct properties — cloudflare.protected_endpoints.domain / d8d367298bf7 / 3

<a id="canonical-8517bf640e5fb50c54faa2053c3a19749554f4c16ebb3598c180ad6c6069499e"></a>

<a id="canonical-46df00fdbf562ea1d1cd47ec824a33d952d0377eeb418693cc0d9da0abca9753"></a>

## exact_value property — cloudflare.protected_endpoints.domain / d8d367298bf7 / 4

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

<a id="canonical-96fc2f012f1c1afcc7f5718e9ae87d8d39749d4e236f9becf4f2db62eac04da9"></a>

<a id="canonical-01d83e6ca2d2a95a41199f864193b72d8d662fad8540e41e04e8b979ac790443"></a>

## regex_value property — cloudflare.protected_endpoints.domain / d8d367298bf7 / 5

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

<a id="canonical-0f7f2e7adfe090c9c679b89b03418c045425967b233d31d72e0695ed244de600"></a>

<a id="canonical-ee4a3ade840ef0e6fb04e8153efec89be674d140c4558647dd748af1e708bd52"></a>

## suffix_value property — cloudflare.protected_endpoints.domain / d8d367298bf7 / 6

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

<a id="canonical-73d07caea97b2dfcffad19760ebf31100d60bce78adbc6b4f99d709db40790fe"></a>

## Next pages — cloudflare.protected_endpoints.domain / d8d367298bf7 / 7

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-9de81935223170e1aa599825ec8823a90bc7e0ae5fcd785c1716908ecc47174f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21377bb5e867312ca57d2b4d2adb125e3ea3d0a393691ba0d0f9952ae86fe65d"></a>

## cloudflare.protected_endpoints.metadata — cloudflare.protected_endpoints.metadata / c63575fbb49a / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [cloudflare](data-sources--protected_application--reference--group-001.md#canonical-93312ce8ce02f39775d8f1492335268b3644069e90c40675ab96abf4210555bd)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- cloudflare.protected_endpoints.metadata

<a id="canonical-e9298dc67363cdef08e52e5c152b4e45bb87b5184635dab9e044f6c78d73f466"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9d9c53e31bc9a268bb5ac8c2ebf87ab4d4d61dd96cd6c99b57b54e69b7d11ae5"></a>

## Direct properties — cloudflare.protected_endpoints.metadata / c63575fbb49a / 3

<a id="canonical-2ea10bbebe554f58914e3401c79d97fb80a417328db6a8d74cf840608bb7503c"></a>

<a id="canonical-cfe23ec270cd1cb853410ec4bde7268d508ae6c8c1a9d4452b5d8912b059acd3"></a>

## description_spec property — cloudflare.protected_endpoints.metadata / c63575fbb49a / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-7a9c480d1b00a68439a857a3ee414c713bc15dc9a043b7df1714b274541ee1e9"></a>

<a id="canonical-49ef84ec567143ef26aa7a9b0c27d100de80a7f8f7f32d2cb9601622e04fe239"></a>

## name property — cloudflare.protected_endpoints.metadata / c63575fbb49a / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-d43619547187864474fc16ec9dea172f3a3953e18d9e4e021d9b0de58d7db4a5"></a>

## Next pages — cloudflare.protected_endpoints.metadata / c63575fbb49a / 6

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-4ca2254153515c507a4adbdd1b4f7333af01221dd3eb7e481a7bcbd67055a05a)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-8c2cc444745ff3d2472c3f90ef963dbc0f050d668c120cb24ba2781a2b7377ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
