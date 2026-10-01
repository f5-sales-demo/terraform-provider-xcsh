---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e1a7926bf79ad279c3ed317607fbd11617117346b190d7c9d075d697f50186d"></a>

## Property reference — Property reference / d9f064e5a984 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- Property reference

<a id="canonical-80f8c3f54e4f38e37e4fd65ef2074d4c6aa1d9661d488b1568a77d52511d7147"></a>

## Direct properties — Property reference / d9f064e5a984 / 3

- [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-fe2a351667b89f7b378f73a2056b469b91391f4341f2c2db4e6ba53162e128bb): complete subsection reference.

<a id="canonical-a9f5a13d7ab593a78e412249ea322b32ed5b7a80760e05d3be93f4b83e4146c5"></a>

<a id="canonical-e7eeedf3234ec71ca71c11b3c06a6a68fe656bf6bc55c528fb38874c78cebcaf"></a>

## annotations property — Property reference / d9f064e5a984 / 4

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

- [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-33502cea2351445a45c0ed492ff21356bb1e679d0d4bd40e91eeda5bd42a1048): complete subsection reference.

- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355): complete subsection reference.

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138): complete subsection reference.

- [custom_connector](resources--protected_application--reference--group-004.md#canonical-015119427834384a4021619b34741d838615efc7b8a44226c1650dcff1fc786e): complete subsection reference.

<a id="canonical-343242c033c262941fe0200be051a82eb401a25a15c1f623d52f328bbcf71c87"></a>

<a id="canonical-205790a902233cd316e06653dcf986160578b0254eda22d2f36bceb32852fc38"></a>

## description property — Property reference / d9f064e5a984 / 5

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

<a id="canonical-fd8a1d72d09eb7e196967f28f53a432b12852174c75d95b03b4fbc2d18c758d2"></a>

<a id="canonical-3caf344440b7765c12e9037293839d9154be81ac57b2e2cdb0bf1600f83acdeb"></a>

## disable property — Property reference / d9f064e5a984 / 6

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

- [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-09aaf1806353d50ca14d6e132e2cb3694e4b277d578cb3e7f1afd534500fa0d9): complete subsection reference.

<a id="canonical-474610d9da6d88b4e10faacb72b941949b532e133a726b2a8f5466527542a61c"></a>

<a id="canonical-1c1d36ef1ac231712d8f9ed49242af7378769579b1d85c57dd44a6a770c5f322"></a>

## id property — Property reference / d9f064e5a984 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-30e16e31fe9f1f51dba05ccdec1cef02d71c987ba919197b077dd135fb1a9dec"></a>

<a id="canonical-7ae54bd14fad6d3401fe38238e72b6b0783f51105f39a33730c81b14f3ed9fda"></a>

## labels property — Property reference / d9f064e5a984 / 8

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

<a id="canonical-5c45cbd15c1de0c1b19c48bc0546ae79f458a1b93d4bdb69b4643a52fa624d2d"></a>

<a id="canonical-ccf71193a0c7a8ff1a87a5b5d2f5ac12c3d1fce3da9af98cb22f644aa0db7672"></a>

## name property — Property reference / d9f064e5a984 / 9

Type: `"string"`. Required.

Name of the Protected Application. Must be unique within the namespace.

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

<a id="canonical-4bf2001d58108b565aea91862ce0fde7cbc537d0b93f1a74d52b184671e149d9"></a>

<a id="canonical-adfdfa999ae8ae0b4fb2a5f9fe90cc07c461ccdd662c4c89ea70dcbcfb48a049"></a>

## namespace property — Property reference / d9f064e5a984 / 10

Type: `"string"`. Required.

Namespace where the Protected Application is created.

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

<a id="canonical-911d26eac05cfd0abd1fe59170d3e5296d281df833932c37e87195c246c8f2ae"></a>

<a id="canonical-957c7650b7583bf8606374f3412c32f318366a5af026c3c7759fb8d0d020edd4"></a>

## region property — Property reference / d9f064e5a984 / 11

Type: `"string"`. Optional, Computed.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

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

- [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-be3e87bf4eaaa0be35a3e894bbcd2e36a0721ae85f61c76da9372a3145122e64): complete subsection reference.

- [timeouts](resources--protected_application--reference--group-004.md#canonical-795c21823eea76f40b96ac6f5e098873ab1262b372f3bba87d702754c6a10784): complete subsection reference.

<a id="canonical-a61bdf85f8fab3ceda87007fbc908b5074d1060336b82160f86f2a14470cc35f"></a>

## All schema paths — Property reference / d9f064e5a984 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `adobe_commerce_connector` | [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-44ed9fb3a940fcfbacc5846c7a1219a377bf559ca3ee9766ebb2a0434aa7f6f6) |
| `annotations` | [annotations](resources--protected_application--reference--group-001.md#canonical-a9f5a13d7ab593a78e412249ea322b32ed5b7a80760e05d3be93f4b83e4146c5) |
| `big_ip_iapp` | [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-7f964c9ca22d58e9d9b04d8d9f1648d340c40b26732d9797e486d37d0d503c83) |
| `cloudflare` | [cloudflare](resources--protected_application--reference--group-001.md#canonical-0f81c386d6c92420e2f6257571a58fc5781085d50a24397716e076b967b06085) |
| `cloudflare.continue_mitigation_action_hdr` | [cloudflare.continue_mitigation_action_hdr](resources--protected_application--reference--group-001.md#canonical-2e694fc6bcb74c3b0d4117c2d52ef283dceea002ade6baba3100dc568295850c) |
| `cloudflare.disable_js_insert` | [cloudflare.disable_js_insert](resources--protected_application--reference--group-001.md#canonical-1c81f87475c4560e27937b2a1cdba2c3391e795c221f4925a646452491b7da33) |
| `cloudflare.disable_mobile_sdk` | [cloudflare.disable_mobile_sdk](resources--protected_application--reference--group-001.md#canonical-c8cdce542cb64c8b475dd745e742575fc9b7a62ac7d1b465b8cff4e925e52206) |
| `cloudflare.js_insertion_rules` | [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-2aee71ee08354bc82a2ab095cab18af6c18e59ed8d88303fcb5aecc7a0b8084e) |
| `cloudflare.js_insertion_rules.exclude_list` | [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-3c9f030f54321fdb742dca1c059f1e8ba861b6cf2c4f2b8dce97a896774b8ff0) |
| `cloudflare.js_insertion_rules.exclude_list.any_domain` | [cloudflare.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-001.md#canonical-8df15559495c8d3bb36f47bcc4a70c29018c6edadd494a9de5b97a8a4c218a61) |
| `cloudflare.js_insertion_rules.exclude_list.domain` | [cloudflare.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-001.md#canonical-cf0c2079ce2ca36f9d0412e495ee16a716436efab09271d8d617d1169872db01) |
| `cloudflare.js_insertion_rules.exclude_list.domain.exact_value` | [cloudflare.js_insertion_rules.exclude_list.domain.exact_value](resources--protected_application--reference--group-001.md#canonical-7c6eefdacd57cdee5fcdb12724c415fe45ecd535ead39981b274d7b9babb1241) |
| `cloudflare.js_insertion_rules.exclude_list.domain.regex_value` | [cloudflare.js_insertion_rules.exclude_list.domain.regex_value](resources--protected_application--reference--group-001.md#canonical-2e0f77fbad3f395d1dcbfa7350c6d20cf9d58944e0f5ec0684b401ef71a372cd) |
| `cloudflare.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudflare.js_insertion_rules.exclude_list.domain.suffix_value](resources--protected_application--reference--group-001.md#canonical-b64730264974c306efb8bf83f8852737267a01efc6987689cf64e70f67c23426) |
| `cloudflare.js_insertion_rules.exclude_list.metadata` | [cloudflare.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-001.md#canonical-b93aa4a8334e458e823387085d6e25429bb71aed5e5d3abaed2e28e097153c20) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudflare.js_insertion_rules.exclude_list.metadata.description_spec](resources--protected_application--reference--group-001.md#canonical-790528c11968367eba673346d4976ea2c247ba29618c0c64b953e61808f651db) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.name` | [cloudflare.js_insertion_rules.exclude_list.metadata.name](resources--protected_application--reference--group-001.md#canonical-ee7815a81de32560b127894af0d239e431abbe922c2c033c26a10098386bc12c) |
| `cloudflare.js_insertion_rules.exclude_list.path` | [cloudflare.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-001.md#canonical-9e6ad2692f11c619ee9836f647fe74cda6c91144a457f43b469fcb0ae20b05f2) |
| `cloudflare.js_insertion_rules.exclude_list.path.path` | [cloudflare.js_insertion_rules.exclude_list.path.path](resources--protected_application--reference--group-001.md#canonical-902f6e3b96e9a84a7c27b396340818c9053caa24b0b8004000fbd59cf94a9df9) |
| `cloudflare.js_insertion_rules.exclude_list.path.prefix` | [cloudflare.js_insertion_rules.exclude_list.path.prefix](resources--protected_application--reference--group-001.md#canonical-8056e800537f36720fe1473b3a3f4395bb1ce1aeee07f6eeb63c13f2524d5513) |
| `cloudflare.js_insertion_rules.exclude_list.path.regex` | [cloudflare.js_insertion_rules.exclude_list.path.regex](resources--protected_application--reference--group-001.md#canonical-4897b90c8dcc5b4e5fac828ca7a37a74684366cae074b89131fed2d4cd190fe4) |
| `cloudflare.js_insertion_rules.javascript_location` | [cloudflare.js_insertion_rules.javascript_location](resources--protected_application--reference--group-001.md#canonical-c24fa9337bdd62b818ed3b4aff18ef636afed1104c1fd92374497525e4098c1c) |
| `cloudflare.js_insertion_rules.js_download_path` | [cloudflare.js_insertion_rules.js_download_path](resources--protected_application--reference--group-001.md#canonical-5ef59e573155e8b914c9debfc83c05d39e98e37d2ede39c02c79612c048def6a) |
| `cloudflare.js_insertion_rules.rules` | [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-27b56b2c81717809b9dfa802c3fb059cc5ec3b19762e31619ae7e3b9fcebf82d) |
| `cloudflare.js_insertion_rules.rules.any_domain` | [cloudflare.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-001.md#canonical-c05ea790d972408c8646b75b83f681b72383ad883bcbabcc3036b8b76dcad8d6) |
| `cloudflare.js_insertion_rules.rules.domain` | [cloudflare.js_insertion_rules.rules.domain](resources--protected_application--reference--group-001.md#canonical-7621e46460f522d606a9d10ff86164c4936ae0fc6623ca129bbef951ee5eba01) |
| `cloudflare.js_insertion_rules.rules.domain.exact_value` | [cloudflare.js_insertion_rules.rules.domain.exact_value](resources--protected_application--reference--group-001.md#canonical-755d3a106ed14187b9ed917409cdc2c6c83667ff0c44c44d19181ac976f665f1) |
| `cloudflare.js_insertion_rules.rules.domain.regex_value` | [cloudflare.js_insertion_rules.rules.domain.regex_value](resources--protected_application--reference--group-001.md#canonical-e3fbd3e600e13afa87bb191ea1c13fcb255b96312da6b79564ea0ead23cb667c) |
| `cloudflare.js_insertion_rules.rules.domain.suffix_value` | [cloudflare.js_insertion_rules.rules.domain.suffix_value](resources--protected_application--reference--group-001.md#canonical-9d69707d50a189a1032ac25bb5c81e95b098bed4b09fef57d0022f728482ed7d) |
| `cloudflare.js_insertion_rules.rules.exact_path` | [cloudflare.js_insertion_rules.rules.exact_path](resources--protected_application--reference--group-001.md#canonical-3fe103456f07c4aeb3fe8fc25ba64fc3987db4889c7cd910ed6359abf42724ef) |
| `cloudflare.js_insertion_rules.rules.glob` | [cloudflare.js_insertion_rules.rules.glob](resources--protected_application--reference--group-001.md#canonical-b06e0855b3f2f59c4f1ce1ab6c2dfd66940c4832ee6e6d47835c624b1b505b74) |
| `cloudflare.js_insertion_rules.rules.metadata` | [cloudflare.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-001.md#canonical-908a4ccc3956dcbf13151f85e7d6f9ff161cb08e99bba25c9e284f6669a42eb1) |
| `cloudflare.js_insertion_rules.rules.metadata.description_spec` | [cloudflare.js_insertion_rules.rules.metadata.description_spec](resources--protected_application--reference--group-001.md#canonical-b23f3d9742c11a43010df50653ba882d9e84f14a569b11556413c01397bb8275) |
| `cloudflare.js_insertion_rules.rules.metadata.name` | [cloudflare.js_insertion_rules.rules.metadata.name](resources--protected_application--reference--group-001.md#canonical-e09ba90830db68c26db84de0ed5e148b15cc7a0ea48c4a142523adf21568cd14) |
| `cloudflare.js_insertion_rules.rules.prefix` | [cloudflare.js_insertion_rules.rules.prefix](resources--protected_application--reference--group-001.md#canonical-4ded4c4a0f9ff306de7967ded5d6864106af1da180087754e5f1ff94d220ac6d) |
| `cloudflare.loglevel` | [cloudflare.loglevel](resources--protected_application--reference--group-001.md#canonical-86ee6dbf4c5524b315f697a199bb9df3b56f80db32b352342d4953cf25d08aae) |
| `cloudflare.manual_js_insert` | [cloudflare.manual_js_insert](resources--protected_application--reference--group-001.md#canonical-a4819691f510f124a0cac4f50f52d74818e5fa6bd910bcd3be3772b6013e0658) |
| `cloudflare.manual_js_insert.js_download_path` | [cloudflare.manual_js_insert.js_download_path](resources--protected_application--reference--group-001.md#canonical-d0b9e40758c071b22b6f6d6b4fce8148ebeabfffdada5dff2bd261dd638ed054) |
| `cloudflare.mobile_sdk_config` | [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-272b23d770180fa02d6dd9597871ff74b653182416c08a8bec5b4fe62003a4ab) |
| `cloudflare.mobile_sdk_config.mobile_identifier` | [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-49d861b862a269662521ebe9a145c25676ef2bc12d83b0284d05e46d338f6f5f) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers` | [cloudflare.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-001.md#canonical-61b4c4262f75d13221b69bf075aeb7862644c32244b29cf7421b46eecbaf5add) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.exact](resources--protected_application--reference--group-001.md#canonical-c47c0f4bbf63c3662886ec5e0f2eee1091cd7f6d06a99ea17f1529e45945e6f9) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.name` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.name](resources--protected_application--reference--group-001.md#canonical-f9424bc143c6dffd36b0a5474d8eccb9f05f6dd5050f02b312ada7260b06f177) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.regex](resources--protected_application--reference--group-001.md#canonical-97aa85cbb4258bcf8b234e17d93af49b1f2010b330c02221a212e81797c12e26) |
| `cloudflare.protected_endpoints` | [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-e63c920461d5cfdb3bb29b5c64453dad91344c21bfb0dad14c9731e318f14be1) |
| `cloudflare.protected_endpoints.any_domain` | [cloudflare.protected_endpoints.any_domain](resources--protected_application--reference--group-001.md#canonical-1749db8d76ec60ba5ece1869d63fc37cba6e0625c4acf127a2853a6230eb2cd8) |
| `cloudflare.protected_endpoints.domain` | [cloudflare.protected_endpoints.domain](resources--protected_application--reference--group-001.md#canonical-7fc8abaf98c3453fa81932828441c6d87b18664502c89c1a2b3efc02a4ca773a) |
| `cloudflare.protected_endpoints.domain.exact_value` | [cloudflare.protected_endpoints.domain.exact_value](resources--protected_application--reference--group-001.md#canonical-d11db7cb06a6b8382178debb9b96ee2dd4ceebe19e32eb73b03e0305008c65ab) |
| `cloudflare.protected_endpoints.domain.regex_value` | [cloudflare.protected_endpoints.domain.regex_value](resources--protected_application--reference--group-001.md#canonical-e58e6a544fac15e1766ba2579348c1945373bb9be5d6c4b7795f2047d17c9a47) |
| `cloudflare.protected_endpoints.domain.suffix_value` | [cloudflare.protected_endpoints.domain.suffix_value](resources--protected_application--reference--group-002.md#canonical-3cb1822b2899b1a9bfc647ccd6f00f0851d91c35b386a9ceac3f644b4b6dcb66) |
| `cloudflare.protected_endpoints.http_methods` | [cloudflare.protected_endpoints.http_methods](resources--protected_application--reference--group-001.md#canonical-8cfda0e6b89c65290096330f84fe3879691cfd7557f405f0d4838f5f4eacf676) |
| `cloudflare.protected_endpoints.metadata` | [cloudflare.protected_endpoints.metadata](resources--protected_application--reference--group-002.md#canonical-06b9a63e266e0610e31e6c5d7d6e25d5334956ba6865b14bcf04878722211a26) |
| `cloudflare.protected_endpoints.metadata.description_spec` | [cloudflare.protected_endpoints.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-e96b8d51d208c7d6e3849e558a7fe022fc9f81b6982b5638ace030fd543f68de) |
| `cloudflare.protected_endpoints.metadata.name` | [cloudflare.protected_endpoints.metadata.name](resources--protected_application--reference--group-002.md#canonical-8cd7a3fb096e73444923d0b4c131273c410ec38b89d40ea823673961a0a295cb) |
| `cloudflare.protected_endpoints.mobile_client` | [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-9fbba17190ee8a82d4b592374f60b1a75a3818281a9d80d3ebcfb85bfe7e9b4e) |
| `cloudflare.protected_endpoints.mobile_client.block` | [cloudflare.protected_endpoints.mobile_client.block](resources--protected_application--reference--group-002.md#canonical-f75831ebdba17d982aaf85633595b628ce6edb2f449ea627d15aebeb20886761) |
| `cloudflare.protected_endpoints.mobile_client.block.body` | [cloudflare.protected_endpoints.mobile_client.block.body](resources--protected_application--reference--group-002.md#canonical-de0e156bd0d81b3ec1d06909ef4543c928ed015541fdd64eddaf89515f89ecce) |
| `cloudflare.protected_endpoints.mobile_client.block.content_type` | [cloudflare.protected_endpoints.mobile_client.block.content_type](resources--protected_application--reference--group-002.md#canonical-3fcd1efeecde9e4f9214a15be5f38b3bd4d22e8cad41631d0260d8d45f5b3fe7) |
| `cloudflare.protected_endpoints.mobile_client.block.status` | [cloudflare.protected_endpoints.mobile_client.block.status](resources--protected_application--reference--group-002.md#canonical-67da709b01321e010e6afa09ac20c2018a5a1e5486a3df9db346b250439d51ec) |
| `cloudflare.protected_endpoints.mobile_client.continue` | [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-56e59604f0d05a45820bdc46d147e3f3ccc56d1e8c0ce5ff094083bbb81f6ffe) |
| `cloudflare.protected_endpoints.mobile_client.continue.add_header` | [cloudflare.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-d56dc4ecf0f68eac81b58d74c7796af6f54d56cbc6b14e6f7e000194ac1ff9bd) |
| `cloudflare.protected_endpoints.mobile_client.continue.no_header` | [cloudflare.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-7caf9af1ff22cda054dfd08f573f3710a83077cd8013bdc7abd544d4efb00905) |
| `cloudflare.protected_endpoints.path` | [cloudflare.protected_endpoints.path](resources--protected_application--reference--group-002.md#canonical-3f15db54066a22f881af4237f32bdf04a2c7c5ebe403e4be2511c09e6dfaab7d) |
| `cloudflare.protected_endpoints.path.caseinsensitive` | [cloudflare.protected_endpoints.path.caseinsensitive](resources--protected_application--reference--group-002.md#canonical-3f39f639ad8695b6416adbdba62aabf5f931cf1c05435c78f58c9d2900550720) |
| `cloudflare.protected_endpoints.path.path` | [cloudflare.protected_endpoints.path.path](resources--protected_application--reference--group-002.md#canonical-4db42f87134ba20eaff689d5fb4c02397ce14705860cda7792a2aec7997a0cc1) |
| `cloudflare.protected_endpoints.query` | [cloudflare.protected_endpoints.query](resources--protected_application--reference--group-001.md#canonical-cf957d83019359eca0220ebeea1bd0d9c086c5a36a3125e75edc6e9b15077e7b) |
| `cloudflare.protected_endpoints.web_client` | [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-d9d2b4d531138187d850c7a8e2b6cf124bba186738d207ff5a851e443fb8c2aa) |
| `cloudflare.protected_endpoints.web_client.block` | [cloudflare.protected_endpoints.web_client.block](resources--protected_application--reference--group-002.md#canonical-d7b0dd9e47fcc53453851dd184216c26e93db3920f198eda9e1be94d5e18994b) |
| `cloudflare.protected_endpoints.web_client.block.body` | [cloudflare.protected_endpoints.web_client.block.body](resources--protected_application--reference--group-002.md#canonical-40cd0c5877941547e3563dc784aeb4f130a69420c7475dc04c3b8e73e09a0681) |
| `cloudflare.protected_endpoints.web_client.block.content_type` | [cloudflare.protected_endpoints.web_client.block.content_type](resources--protected_application--reference--group-002.md#canonical-4f16cda1a30cf438eb57c4bc617e493d4f891627e7ed987590a57c34616a6e3f) |
| `cloudflare.protected_endpoints.web_client.block.status` | [cloudflare.protected_endpoints.web_client.block.status](resources--protected_application--reference--group-002.md#canonical-a5e14c83d47648a6d583b7046131e3132819ac27ca9a4f7b73accfd2fb306b51) |
| `cloudflare.protected_endpoints.web_client.continue` | [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-13c40c6e9289f06dd1112d121ee5b1f6af5cade9ba248276b863bc6ac48e1f92) |
| `cloudflare.protected_endpoints.web_client.continue.add_header` | [cloudflare.protected_endpoints.web_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-b9cf3159d2fd7e2212cf5ef268bbfe85aa740bcbfb32e90dacfac49c002603f0) |
| `cloudflare.protected_endpoints.web_client.continue.no_header` | [cloudflare.protected_endpoints.web_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-35b8ed51e1f4a7aa35b3f3ba30fa28f1a125c586f6825473db38b7c06d8af8cc) |
| `cloudflare.protected_endpoints.web_client.redirect` | [cloudflare.protected_endpoints.web_client.redirect](resources--protected_application--reference--group-002.md#canonical-bcf5b26eb495170ee86bfeeb66648457d84f50c3bac0e7e4851c3180bd4cc0e5) |
| `cloudflare.protected_endpoints.web_client.redirect.location` | [cloudflare.protected_endpoints.web_client.redirect.location](resources--protected_application--reference--group-002.md#canonical-f3e8b0e828e6b5d49ed4863bd2bcda858aa9e2c35c600326a0166c7172ef3d23) |
| `cloudflare.protected_endpoints.web_client.redirect.status` | [cloudflare.protected_endpoints.web_client.redirect.status](resources--protected_application--reference--group-002.md#canonical-e0c129b1ae365c6da4d790175f47ff57ebea807b3e07bfeaf7077038ae6df792) |
| `cloudflare.protected_endpoints.web_mobile_client` | [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-9b748861d6d0c22a4062dd4599f06ff03a225ccdffa45ad2cc3be32a18b8d158) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--reference--group-002.md#canonical-1f177f475b234766682e7580a454987b45c40cb5160999617d53d76bee5d4884) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.body](resources--protected_application--reference--group-002.md#canonical-336877c6758134bd4f11cbc04f7d4c7ec252b740c1d2ef0c23ef0ef87c9c8fbf) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type](resources--protected_application--reference--group-002.md#canonical-03151bc852efb1572583eb5afff29edb51fdc9d093efa5d1132435dd8d9e3dd6) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.status](resources--protected_application--reference--group-002.md#canonical-554b2cee32fd0d2412ec1dadcd51aa9236b1f0aa13e758628a6b9359f9a2a61d) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web` | [cloudflare.protected_endpoints.web_mobile_client.block_web](resources--protected_application--reference--group-002.md#canonical-99701a6507543c569f93387d5fd0aa7140e79d33b56d9104a482a780ccb21b26) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.body` | [cloudflare.protected_endpoints.web_mobile_client.block_web.body](resources--protected_application--reference--group-002.md#canonical-b25296dbfe972d13197c0381890b04dc823b337145256c5055ce5506b3c8d040) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_web.content_type](resources--protected_application--reference--group-002.md#canonical-2cb2add3da283bf1100c7c474759cd7c9b2141b17e6ff88e9fb219a987b72ad9) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.status` | [cloudflare.protected_endpoints.web_mobile_client.block_web.status](resources--protected_application--reference--group-002.md#canonical-dbf7a2200ce1ef20e07a0beb1cff986d3bbfeafac7743604c1b98c42b3a3efe3) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-08d882d43e4a8e60ecc084cc5e6b19a12f9d911de32b20151b54e4f692611187) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--reference--group-002.md#canonical-622a48451b9f319f091e9f5f5cbf0df3c68d86864ade1b4ddfb25ba2a59bf5f3) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--reference--group-002.md#canonical-f3c9eb8a7d3ddd7299d5986a8c972a9bf26de7bf1d0e3f21c8f2a8623279ff33) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web` | [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-6d5cc2928083ab9ab9d9b201173e2f263c546f28339eab534421ef2992e29196) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--reference--group-002.md#canonical-d66d1a7a2bcd24ccec8bc4c76f9fc8d244beedbe0168b483c992964bc1d74d12) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--reference--group-002.md#canonical-25c7e0c85fec933c04640822b41944dc009254c23d3b3b9fe8807b4a02ff211c) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--reference--group-002.md#canonical-f9647c23a9c2241c8165c0d3ab51ebd70c5fc926cfa6c145c9bd7357c146185b) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.location](resources--protected_application--reference--group-002.md#canonical-4f4b771966560e369310acf738941121ad6aa9365c0a1ca84a1338784af418f6) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.status](resources--protected_application--reference--group-002.md#canonical-6d9baf5782474f90c24c57532fdc1b85176c1dc9b1d5b3d4e098fadcb52f0981) |
| `cloudflare.timeout` | [cloudflare.timeout](resources--protected_application--reference--group-001.md#canonical-97e012b0d77532681f43aea2c21444db7ac4b9f78b82a118ad4f8496b5741593) |
| `cloudflare.trusted_clients` | [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-5dfde5e2b1cf80ffb447f5765569159e2fe4cfff8466709bef2882a5e8900244) |
| `cloudflare.trusted_clients.http_header` | [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-a72319ecab94e3cf22e2709be43a2ab0e24e98dae777d0804684a413e02d14d7) |
| `cloudflare.trusted_clients.http_header.headers` | [cloudflare.trusted_clients.http_header.headers](resources--protected_application--reference--group-002.md#canonical-fcb4ade7eb3c48cd378384820a2d4cc3329bedb497068064ef2d5bdb02765d34) |
| `cloudflare.trusted_clients.http_header.headers.exact` | [cloudflare.trusted_clients.http_header.headers.exact](resources--protected_application--reference--group-002.md#canonical-7fc26feef11b995253f93216e71771e86b8906107dad18c272ca45840a0bd6d5) |
| `cloudflare.trusted_clients.http_header.headers.name` | [cloudflare.trusted_clients.http_header.headers.name](resources--protected_application--reference--group-002.md#canonical-b9ab3334a43295db8e451d2743f87d1808d987cdb3a8e7d1fd7c0f9c2b8503ad) |
| `cloudflare.trusted_clients.http_header.headers.regex` | [cloudflare.trusted_clients.http_header.headers.regex](resources--protected_application--reference--group-002.md#canonical-ac65196f0ef59e4f305d4df1199a58807a786c0ff81f4ef30a5f7bdfd58f5187) |
| `cloudflare.trusted_clients.ip_prefix` | [cloudflare.trusted_clients.ip_prefix](resources--protected_application--reference--group-002.md#canonical-a1a8eb4b50557e803404a10112511757528039f712c76f0b6adf3663a952deb8) |
| `cloudflare.trusted_clients.metadata` | [cloudflare.trusted_clients.metadata](resources--protected_application--reference--group-002.md#canonical-a1ea7ad629d33e0d344e7752f4c75f25cc3e7212f7577ddd4fc06dfd8befeae3) |
| `cloudflare.trusted_clients.metadata.description_spec` | [cloudflare.trusted_clients.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-15970282e564761db1d7dacd86ad21c271e6eac70c4b94eed0a3cf319be22d39) |
| `cloudflare.trusted_clients.metadata.name` | [cloudflare.trusted_clients.metadata.name](resources--protected_application--reference--group-002.md#canonical-7b7d7fbc44c7804ee1b0be73deebaf34fe72fe8563e4be86a3d7faa10458a323) |
| `cloudfront` | [cloudfront](resources--protected_application--reference--group-002.md#canonical-3b163d79026cd25596af8a276de79a6ad6020516bd2678119c801b88b42eff3c) |
| `cloudfront.aws_configuration_id_selector` | [cloudfront.aws_configuration_id_selector](resources--protected_application--reference--group-002.md#canonical-e328fb8ff021e8b85230e8fc8302d915f641ec34db21267cc5c7511d22342dc9) |
| `cloudfront.aws_configuration_id_selector.ids` | [cloudfront.aws_configuration_id_selector.ids](resources--protected_application--reference--group-002.md#canonical-f3a6fbde38fe24eceb4ac92e6f13e3a29078c0517fedea7c8724f4cb7ef7c9a0) |
| `cloudfront.aws_configuration_tag_selector` | [cloudfront.aws_configuration_tag_selector](resources--protected_application--reference--group-002.md#canonical-d582bbeb60e2497c2c0f73bb4941a679d32f6001ff09cd8a68465f25f3efcfe0) |
| `cloudfront.aws_configuration_tag_selector.tags` | [cloudfront.aws_configuration_tag_selector.tags](resources--protected_application--reference--group-002.md#canonical-c82f2f14959863604f7bce9b358cbbc232f3234642193249a8cb1bfcbc8b7a4b) |
| `cloudfront.continue_mitigation_action_hdr` | [cloudfront.continue_mitigation_action_hdr](resources--protected_application--reference--group-002.md#canonical-cb4904b08850cb9ba80ce4401dd09f5bee5c6f326d37f46b35af7ce6a2202a7f) |
| `cloudfront.data_sample` | [cloudfront.data_sample](resources--protected_application--reference--group-002.md#canonical-9acc36ecd0fb6e1810881fbfcde54e76db191ba716628b7b833938ed30d2b1c4) |
| `cloudfront.disable_aws_configuration` | [cloudfront.disable_aws_configuration](resources--protected_application--reference--group-002.md#canonical-c4fad582262a682012df3d9ff968216d37ccb8daf70f4bdff2cac1d81b5c117e) |
| `cloudfront.disable_js_insert` | [cloudfront.disable_js_insert](resources--protected_application--reference--group-002.md#canonical-6c4abbca2784e82299b1cd302146adb3e02d2f13d4e3de54b859632125bd9293) |
| `cloudfront.disable_mobile_sdk` | [cloudfront.disable_mobile_sdk](resources--protected_application--reference--group-002.md#canonical-0b22b8e5eae9c12a735a317a02338769ed0dfad44792788df88f3a18949abf28) |
| `cloudfront.js_insertion_rules` | [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-e4bcdfb97ac85043699329beb8615ed70d01931692aaf002e0e02a111ccd016b) |
| `cloudfront.js_insertion_rules.exclude_list` | [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-513efa0f572c03400cc9a9a78acf3f2e9680a4fdd0bf4ca673ee5be70d167a9b) |
| `cloudfront.js_insertion_rules.exclude_list.any_domain` | [cloudfront.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-002.md#canonical-f25ea3f5bca89f0c6ebac56cded2527733089337b2131f02f2f4286a500e4bc0) |
| `cloudfront.js_insertion_rules.exclude_list.domain` | [cloudfront.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-002.md#canonical-142d157c900d1864e8931dd6da19be1da307d8c300682248665880090e150a8d) |
| `cloudfront.js_insertion_rules.exclude_list.domain.exact_value` | [cloudfront.js_insertion_rules.exclude_list.domain.exact_value](resources--protected_application--reference--group-002.md#canonical-ff98dfa46f0621a4401a4176258dedf14c3998b122fadc3824aa50279cdb3238) |
| `cloudfront.js_insertion_rules.exclude_list.domain.regex_value` | [cloudfront.js_insertion_rules.exclude_list.domain.regex_value](resources--protected_application--reference--group-002.md#canonical-66d873d7f651efb02de1f01606eba923c8616b02d6aaf9eeafb908f93db53fda) |
| `cloudfront.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudfront.js_insertion_rules.exclude_list.domain.suffix_value](resources--protected_application--reference--group-002.md#canonical-c4d88e39c99c6eaccd2cfbddb9f067260563a46b6e6b564e769e51f54fb93809) |
| `cloudfront.js_insertion_rules.exclude_list.metadata` | [cloudfront.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-002.md#canonical-3d743622f969dc2198f396b1d3431b8200af63d2d9c28a0356f798cf50db4b5f) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudfront.js_insertion_rules.exclude_list.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-9001d0f68a54c800d56cb30b80ed40c9dec1871acee561052c54805d7fe2faa4) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.name` | [cloudfront.js_insertion_rules.exclude_list.metadata.name](resources--protected_application--reference--group-002.md#canonical-ad57ea524a551dce4b56c2962f495647fcd720c626e63973dea8671aa3e45de6) |
| `cloudfront.js_insertion_rules.exclude_list.path` | [cloudfront.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-002.md#canonical-66fe972d83c54877aefed9fa4b067c9b71161971a89d5979ea73f0e0e94c43bb) |
| `cloudfront.js_insertion_rules.exclude_list.path.path` | [cloudfront.js_insertion_rules.exclude_list.path.path](resources--protected_application--reference--group-002.md#canonical-7eb00272f4d133184b00ae7ef484b14592c535540f59a017d7a303cd3a56a608) |
| `cloudfront.js_insertion_rules.exclude_list.path.prefix` | [cloudfront.js_insertion_rules.exclude_list.path.prefix](resources--protected_application--reference--group-002.md#canonical-36254fe61f2c6729f7707f002b12647dd5091e39ac8c3b79fa158226238b23a8) |
| `cloudfront.js_insertion_rules.exclude_list.path.regex` | [cloudfront.js_insertion_rules.exclude_list.path.regex](resources--protected_application--reference--group-002.md#canonical-7c9020a8e03f56f777f1cf2688df44c845f3c997d82cd512525a5d4a2f11e037) |
| `cloudfront.js_insertion_rules.javascript_location` | [cloudfront.js_insertion_rules.javascript_location](resources--protected_application--reference--group-002.md#canonical-a25ff865223ff51cda59e5b8123fc5ae1876478afbf225730e5a82c12c4ca22d) |
| `cloudfront.js_insertion_rules.javascript_mode` | [cloudfront.js_insertion_rules.javascript_mode](resources--protected_application--reference--group-002.md#canonical-5c5d0b0fea4339238f3cf948ad93e0d449828925ca9dcab005fd20b7c974f73d) |
| `cloudfront.js_insertion_rules.js_download_path` | [cloudfront.js_insertion_rules.js_download_path](resources--protected_application--reference--group-002.md#canonical-686438b2692c84f2ff7fd14b97322884d821d96bd45cc834e4bec92779eae114) |
| `cloudfront.js_insertion_rules.rules` | [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-aec65af9e99ebc34af80f0920491c90d525b65be7a05fab4f2425f3b8197f7bd) |
| `cloudfront.js_insertion_rules.rules.any_domain` | [cloudfront.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-002.md#canonical-9ed4c5145b9015ac58df873ec7168017cc18345d59033d4997f7d8fffafcc53e) |
| `cloudfront.js_insertion_rules.rules.domain` | [cloudfront.js_insertion_rules.rules.domain](resources--protected_application--reference--group-002.md#canonical-0740a08332ffd2e98195fe85cf195b171a76df40b27f6f7064031f13160151b3) |
| `cloudfront.js_insertion_rules.rules.domain.exact_value` | [cloudfront.js_insertion_rules.rules.domain.exact_value](resources--protected_application--reference--group-002.md#canonical-374ea9e9b8e32426631fab5663189d9d488ef9efc260f4039ccc74e574e17899) |
| `cloudfront.js_insertion_rules.rules.domain.regex_value` | [cloudfront.js_insertion_rules.rules.domain.regex_value](resources--protected_application--reference--group-002.md#canonical-60be98993a2be1af5bac84faeb50932e575db617963213c4f4954f9ddf639b29) |
| `cloudfront.js_insertion_rules.rules.domain.suffix_value` | [cloudfront.js_insertion_rules.rules.domain.suffix_value](resources--protected_application--reference--group-002.md#canonical-e4e0470128574b7bc07f42ea1f4fb19e5aeccea04c19ec84486adcccbb90e352) |
| `cloudfront.js_insertion_rules.rules.exact_path` | [cloudfront.js_insertion_rules.rules.exact_path](resources--protected_application--reference--group-002.md#canonical-3bd4732ad653bf3b2f2fefe9bdc03ab7c7b70e5fa552da20d314be3640993448) |
| `cloudfront.js_insertion_rules.rules.glob` | [cloudfront.js_insertion_rules.rules.glob](resources--protected_application--reference--group-002.md#canonical-8b20ee2e59fd61580deb396841cfd2a25fc3d35abc8a972d9b46f956469d0170) |
| `cloudfront.js_insertion_rules.rules.metadata` | [cloudfront.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-002.md#canonical-4d11e7564ae09b9ddd9fa38c5af7cfc69a5fe0d6a8fdda220a696055f37c6f72) |
| `cloudfront.js_insertion_rules.rules.metadata.description_spec` | [cloudfront.js_insertion_rules.rules.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-46e95e2e243dda92d7812bfce9584dba5bbf2dfe6005123dad2b3d81917018a1) |
| `cloudfront.js_insertion_rules.rules.metadata.name` | [cloudfront.js_insertion_rules.rules.metadata.name](resources--protected_application--reference--group-002.md#canonical-51abe93bb3c9c92c170d3b9d3cac8322e0a6b4f5f323403ee6fc894d29d62d3b) |
| `cloudfront.js_insertion_rules.rules.prefix` | [cloudfront.js_insertion_rules.rules.prefix](resources--protected_application--reference--group-002.md#canonical-eb363bce831a18a6ae7d01efe26dcba2bcdc788047942742b26b93ae3566ae94) |
| `cloudfront.loglevel` | [cloudfront.loglevel](resources--protected_application--reference--group-002.md#canonical-209ea247869b3bd007a8fb169b4eeccdfe39e98fa7d2c0c21b974478be4ac759) |
| `cloudfront.manual_js_insert` | [cloudfront.manual_js_insert](resources--protected_application--reference--group-002.md#canonical-d5d40e8c0a4002482e6e5f8ee9d336d892f567a2ddf0df8bbe8ce1a3920726ac) |
| `cloudfront.manual_js_insert.javascript_mode` | [cloudfront.manual_js_insert.javascript_mode](resources--protected_application--reference--group-002.md#canonical-0c068ab569fea3f3df56aa2a55c5a2b2cb6f602db17de20425955fd51b0cadad) |
| `cloudfront.manual_js_insert.js_download_path` | [cloudfront.manual_js_insert.js_download_path](resources--protected_application--reference--group-002.md#canonical-050f1a9274bd8da4191c2dec0ba0c80f4a6c84ce0b640dd9c9a847dcf258e7ff) |
| `cloudfront.mobile_sdk_config` | [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-002.md#canonical-034acbfa5b732cb6741535d55072adf9ceae85dc60b62d644bd712cabc1d028b) |
| `cloudfront.mobile_sdk_config.mobile_identifier` | [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-002.md#canonical-4bb491bbc9a218e4e6db240e6273455e23c1d540f024b712373d12452552e839) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers` | [cloudfront.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-003.md#canonical-04c7d044f238d79f67100a3bb6da6bf1bccbcdd0147d41012babbfcce53865bf) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.exact](resources--protected_application--reference--group-003.md#canonical-54da92f86f18fa1389e2a1f6dda96c02da1f7c4a98e6babedeacb820105de352) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.name` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.name](resources--protected_application--reference--group-003.md#canonical-5f045cf278bff1e8dcacb7c0bf30d751cebf46f56e6b50623032bef587269576) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.regex](resources--protected_application--reference--group-003.md#canonical-22227d571b6ce6051d3a223bf45f1c36754b6a888a351b6bac696466e118cb2c) |
| `cloudfront.protected_endpoints` | [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-1ce96a8cc4d2f9ec2a090b4716d1f65513a8d3a73fb0c4e9e7f8c43e0cce7267) |
| `cloudfront.protected_endpoints.any_domain` | [cloudfront.protected_endpoints.any_domain](resources--protected_application--reference--group-003.md#canonical-12f7886fee20e40a393d9b183b378bdfda4fc514af54f5aa97c8c574c4453917) |
| `cloudfront.protected_endpoints.domain` | [cloudfront.protected_endpoints.domain](resources--protected_application--reference--group-003.md#canonical-7753624b1b32d0d177d2d1120ebf9d4d78c31bcab5308dcb7fcf1cb4c413e161) |
| `cloudfront.protected_endpoints.domain.exact_value` | [cloudfront.protected_endpoints.domain.exact_value](resources--protected_application--reference--group-003.md#canonical-d86af2565dd76157e11c899080d0051ede3c400d0ec9a6bd13ffb23bc1987193) |
| `cloudfront.protected_endpoints.domain.regex_value` | [cloudfront.protected_endpoints.domain.regex_value](resources--protected_application--reference--group-003.md#canonical-63c643201218fcc4969e3319d4cab743a28c5530eb37aca0fdb7f8bc0ac9eacb) |
| `cloudfront.protected_endpoints.domain.suffix_value` | [cloudfront.protected_endpoints.domain.suffix_value](resources--protected_application--reference--group-003.md#canonical-e3bd38293b1fddad1c275c3ff18c2ebef6281b7ec602ef32a54d8df168670eda) |
| `cloudfront.protected_endpoints.flow_label` | [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83a32201ab330d6813e10fcd015cafec407fb4e2b3e4ba52bbd7cca2bf628849) |
| `cloudfront.protected_endpoints.flow_label.account_management` | [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-aaafefc4910e854afdc5341410ece28265762c8b45eb839430b885a3e1b7dbd7) |
| `cloudfront.protected_endpoints.flow_label.account_management.create` | [cloudfront.protected_endpoints.flow_label.account_management.create](resources--protected_application--reference--group-003.md#canonical-5b2eec6a67db95b939f1df26b638896f266c915709d8d38967130892093f6a76) |
| `cloudfront.protected_endpoints.flow_label.account_management.password_reset` | [cloudfront.protected_endpoints.flow_label.account_management.password_reset](resources--protected_application--reference--group-003.md#canonical-d7e490e20403f21d675bfa0c1d9dd59386ee8b4044b44a28eb89d5d520c1085a) |
| `cloudfront.protected_endpoints.flow_label.authentication` | [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-981640247631decf7b8e60c415325a4720cdb8a78510eef204adaa3dc8378ee4) |
| `cloudfront.protected_endpoints.flow_label.authentication.login` | [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-90a0c7cdaa218792b38dc262258b372c0b748f22bd1c9aa0b75fab7d79833622) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](resources--protected_application--reference--group-003.md#canonical-0bcf11dd58d6f1cda1b51c0f0f6e49e3932bb51aee0ab4eb83f440d66fcc6920) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-76e450f93d4aed1e27042927ea04995f11db83eb979a54cfe2199fd93d96a1fd) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--protected_application--reference--group-003.md#canonical-ba22f3d76edf5a8ab172867e782c71a67911ce13aee9b967fe2620c4397303b2) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name](resources--protected_application--reference--group-003.md#canonical-83562465fb7330083906df91ce4f3a36d4d271465876f2771027b4306ba8080f) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values](resources--protected_application--reference--group-003.md#canonical-4520995a8bec4ab356e3381c938b712acf7fd80e1070753d9e20869caf4381a0) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status](resources--protected_application--reference--group-003.md#canonical-5fde6390d722bd3660a4c58de7cf4d29ba000f0f5dac78dd12dc1b70d1699162) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--protected_application--reference--group-003.md#canonical-b3bf6d8ffeadb6b039701c31ea2d83531c33c518ccfbf403d12c7f2f371ade04) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name](resources--protected_application--reference--group-003.md#canonical-daaeb317717971fccec44d21f4227b9823b6ead94375c085920375b266e396be) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values](resources--protected_application--reference--group-003.md#canonical-aaa8013f464a1b5b642be3ca1c5cf9c5f94187bfb762dd0e12e3ef1dd5eea65f) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status](resources--protected_application--reference--group-003.md#canonical-3b304cd79a56bc1bf19f2647c26e763bd2c8ff8af69811fe8384a49720a74ecd) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_mfa` | [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](resources--protected_application--reference--group-003.md#canonical-c3c229936db01d04708e709d4bd88d4ad13708cc71ff8bba7f926cb5c5b79200) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_partner` | [cloudfront.protected_endpoints.flow_label.authentication.login_partner](resources--protected_application--reference--group-003.md#canonical-5ecde7bfc6c6e927eee1665462e36ac6ecc0ebc189945b3f5c6e577e7a69afcf) |
| `cloudfront.protected_endpoints.flow_label.authentication.logout` | [cloudfront.protected_endpoints.flow_label.authentication.logout](resources--protected_application--reference--group-003.md#canonical-e4ebf30528b3cd72350e16c6ad342db14adc908a317c941f8e3303b866827ecd) |
| `cloudfront.protected_endpoints.flow_label.authentication.token_refresh` | [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](resources--protected_application--reference--group-003.md#canonical-1f21eb274eb0d8a92f7cb73696ec56b8697cd96c195cd4cab8519c2cfe9b870d) |
| `cloudfront.protected_endpoints.flow_label.financial_services` | [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-f52da908064dc74796cc44d32b471d572c9e5f56fa2eed181db323598f14bb5f) |
| `cloudfront.protected_endpoints.flow_label.financial_services.apply` | [cloudfront.protected_endpoints.flow_label.financial_services.apply](resources--protected_application--reference--group-003.md#canonical-548f04d5e0eb1c07d1620d23d357ace6f4fdfc453fc93cf58e62e341cc940cfe) |
| `cloudfront.protected_endpoints.flow_label.financial_services.money_transfer` | [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](resources--protected_application--reference--group-003.md#canonical-06f3edbc9a417c4daa83744ea0efd9208ecca9a9653f13ea6f9a4c7752acc660) |
| `cloudfront.protected_endpoints.flow_label.flight` | [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-f1e5d2f732f171ccf4b9432accdeed88783aeecb7e8ce55ccc2d29f8623e1fcb) |
| `cloudfront.protected_endpoints.flow_label.flight.checkin` | [cloudfront.protected_endpoints.flow_label.flight.checkin](resources--protected_application--reference--group-003.md#canonical-993837945eb512d856b6536a553919c657f448769b4e93eab7737a0b337be032) |
| `cloudfront.protected_endpoints.flow_label.profile_management` | [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e186dcf62e2fb13c8c6c4f40d55c6ed52ccb66a956f989798eba1d63c5a0c50c) |
| `cloudfront.protected_endpoints.flow_label.profile_management.create` | [cloudfront.protected_endpoints.flow_label.profile_management.create](resources--protected_application--reference--group-003.md#canonical-020f739e4d703d2859d1baf7d88833e6d9c2241fa183fd6e46936833dd99f668) |
| `cloudfront.protected_endpoints.flow_label.profile_management.update` | [cloudfront.protected_endpoints.flow_label.profile_management.update](resources--protected_application--reference--group-003.md#canonical-6a6aa4101a346c9d54234a91ae2ae2cafcbcc17a4fa030fd1ed11a910895291c) |
| `cloudfront.protected_endpoints.flow_label.profile_management.view` | [cloudfront.protected_endpoints.flow_label.profile_management.view](resources--protected_application--reference--group-003.md#canonical-568e79bd9d85ae1a26c6036211e49a4dbc137e523252874c4de28fab7567a9df) |
| `cloudfront.protected_endpoints.flow_label.search` | [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-4f6801c3605c073434c3c88ce315256de1e15b32842c675f50e268fb9ae8ad4b) |
| `cloudfront.protected_endpoints.flow_label.search.flight_search` | [cloudfront.protected_endpoints.flow_label.search.flight_search](resources--protected_application--reference--group-003.md#canonical-8e6394182be2d3201328935e3a97bd9bd411531b8f7d1afb7b9d1a8f89aa3f12) |
| `cloudfront.protected_endpoints.flow_label.search.product_search` | [cloudfront.protected_endpoints.flow_label.search.product_search](resources--protected_application--reference--group-003.md#canonical-d98943c38434994273e9761c98bc47e789231c6147282989d59cc7d26f60ec19) |
| `cloudfront.protected_endpoints.flow_label.search.reservation_search` | [cloudfront.protected_endpoints.flow_label.search.reservation_search](resources--protected_application--reference--group-003.md#canonical-783dcc97359da602704553e2534861ff1b870be1f7589f6d60601906a2217988) |
| `cloudfront.protected_endpoints.flow_label.search.room_search` | [cloudfront.protected_endpoints.flow_label.search.room_search](resources--protected_application--reference--group-003.md#canonical-ba0d7067ffddfec03099a3977cbb7d777d31db6a4981549974a5cfc717a87f08) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-7ec8beac847eb2676a0ef58023d81a67d916def0edf6196c3de59c04ee603a62) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--protected_application--reference--group-003.md#canonical-a93be5c799e2c476be09cea734be5be424c6f18f0e8985a1446ebaad16c48a6f) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--protected_application--reference--group-003.md#canonical-287a7913cbe6a94d19ebdd56d57a3117b725a7bbf47486635d9b1efdc3a7fdcb) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--protected_application--reference--group-003.md#canonical-4636ce24ac98300c174284ab1c663d23d05b83fbb32e95ab49f889d36654c2bd) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--protected_application--reference--group-003.md#canonical-a2787939120d9f84b59531b7b7651aeabaadd37f0d621c2fb4da184ecd8d7e05) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--protected_application--reference--group-003.md#canonical-d3d7bff1125d764be0f06fec4081752f5be017db3b5fbfc6e80785e6744f5f8f) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--protected_application--reference--group-003.md#canonical-11413064cb9a7f6da1ebe0839433fe4456cc14b82883c972613d7778f0b4e316) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--protected_application--reference--group-003.md#canonical-0b1295607c98d7176bac09ebb1c8fa98cc4d3b136929f52f7fa5270285212a5b) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](resources--protected_application--reference--group-003.md#canonical-22d8a702fa7af6cac41f2fe59d891a98dbff5d5c5a26ccfa8181a5e2c914d998) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--protected_application--reference--group-003.md#canonical-4ffdae5f1255b93a2809abb379c73ace557941bb77cd834e369817d652c71140) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--protected_application--reference--group-003.md#canonical-c3b5459a019322c5efb0ce74746ec65134b2d7eb266c9c13cac779db95920d62) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--protected_application--reference--group-003.md#canonical-df49a6a411dc26eb472b1065af19fed7759596674a61c71a836f9b23fcc48ffa) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--protected_application--reference--group-003.md#canonical-af700b52fba5ecab03f005e141697da1ffad66ebaf867f57eef165d4ea460a17) |
| `cloudfront.protected_endpoints.http_methods` | [cloudfront.protected_endpoints.http_methods](resources--protected_application--reference--group-003.md#canonical-f61764d96305487347982735124b327c670e50226cb024ee9a9f5df51c2b6e2e) |
| `cloudfront.protected_endpoints.metadata` | [cloudfront.protected_endpoints.metadata](resources--protected_application--reference--group-003.md#canonical-c6b6e83335af3e580b478be18a158c2bfb7ab8fc68da0367d04656e46f1cf240) |
| `cloudfront.protected_endpoints.metadata.description_spec` | [cloudfront.protected_endpoints.metadata.description_spec](resources--protected_application--reference--group-003.md#canonical-e67d8edf25607cdbf19d8ae18b4cab2558ecdee15605b59262c4a48693b808c9) |
| `cloudfront.protected_endpoints.metadata.name` | [cloudfront.protected_endpoints.metadata.name](resources--protected_application--reference--group-003.md#canonical-f66637c2c29e79a015031118243dac74bf362bc1771a6e92388017c387413ba7) |
| `cloudfront.protected_endpoints.mobile_client` | [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-58b0690a51f840faa3198ae2ad73ed65b9a4795b90053764d6a608f1a661c75e) |
| `cloudfront.protected_endpoints.mobile_client.block` | [cloudfront.protected_endpoints.mobile_client.block](resources--protected_application--reference--group-003.md#canonical-dc003019d30a72221594d38b63313e14941e1572e76bc5526726147f816e5c72) |
| `cloudfront.protected_endpoints.mobile_client.block.body` | [cloudfront.protected_endpoints.mobile_client.block.body](resources--protected_application--reference--group-003.md#canonical-c1f1dbd01e5e13d86586b994d36a4272122dcd2748513e516252545649e50500) |
| `cloudfront.protected_endpoints.mobile_client.block.content_type` | [cloudfront.protected_endpoints.mobile_client.block.content_type](resources--protected_application--reference--group-003.md#canonical-a15d067ecce83d131919fddf2c544aef8db51d1bc619f61658b5ec191a508aca) |
| `cloudfront.protected_endpoints.mobile_client.block.status` | [cloudfront.protected_endpoints.mobile_client.block.status](resources--protected_application--reference--group-003.md#canonical-890d778f22938af229acc87deb6f75271460c5e6efd3bf7e3fc8cb14095208a4) |
| `cloudfront.protected_endpoints.mobile_client.continue` | [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-004.md#canonical-1a4962871cef4997128adf0e4ea5a58a161f329e7850cf385ba1910ac9ea3ce2) |
| `cloudfront.protected_endpoints.mobile_client.continue.add_header` | [cloudfront.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--reference--group-004.md#canonical-98542123d7004c5fdd6bce13686a595f5157ebb9d16faadb35792315366c7037) |
| `cloudfront.protected_endpoints.mobile_client.continue.no_header` | [cloudfront.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--reference--group-004.md#canonical-222b671b76d4141a0ef64b956c6d835a9f39d00ccbb700c8a8cbebeb729ce63f) |
| `cloudfront.protected_endpoints.path` | [cloudfront.protected_endpoints.path](resources--protected_application--reference--group-003.md#canonical-269779a9ada28bbbb11d02b36ea09b58d89ef7317462dd5d78d59183e395dbdd) |
| `cloudfront.protected_endpoints.query` | [cloudfront.protected_endpoints.query](resources--protected_application--reference--group-003.md#canonical-5de944d0b4b591bde214466448d5a5d86e37a89b1803bc66193995438d04af12) |
| `cloudfront.protected_endpoints.undefined_flow_label` | [cloudfront.protected_endpoints.undefined_flow_label](resources--protected_application--reference--group-004.md#canonical-f4177e4a9b9942cc7a3a0ea3aa3e25c2a7da6c84bb03d9616518d8a99799ffe3) |
| `cloudfront.protected_endpoints.web_client` | [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-0d38b0aff2bb4d072a105a35662d782c723c056242112a0250a9a091f3602048) |
| `cloudfront.protected_endpoints.web_client.block` | [cloudfront.protected_endpoints.web_client.block](resources--protected_application--reference--group-004.md#canonical-c92e7edb1e3ea7dc79b72699bfd0a855a1844bfc59d5a9b80862fa2163c472b4) |
| `cloudfront.protected_endpoints.web_client.block.body` | [cloudfront.protected_endpoints.web_client.block.body](resources--protected_application--reference--group-004.md#canonical-147c0ec59d1b8f190dd63d865bdb6a32a2adf9462c865b26d6a0301fe4d912cd) |
| `cloudfront.protected_endpoints.web_client.block.content_type` | [cloudfront.protected_endpoints.web_client.block.content_type](resources--protected_application--reference--group-004.md#canonical-ae0c6ea6180bff2950bb0ea033e69f77412f1041d44ae6b2bea105f8a4750906) |
| `cloudfront.protected_endpoints.web_client.block.status` | [cloudfront.protected_endpoints.web_client.block.status](resources--protected_application--reference--group-004.md#canonical-a1bb8c6c0c338365c2dff0ddc1fdeb83e70dd63cb1aca853d7795804775f8876) |
| `cloudfront.protected_endpoints.web_client.continue` | [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--reference--group-004.md#canonical-b54cec048df2da1e56d767276c7fd2567be771c967bc854f13db8b00b77364f4) |
| `cloudfront.protected_endpoints.web_client.continue.add_header` | [cloudfront.protected_endpoints.web_client.continue.add_header](resources--protected_application--reference--group-004.md#canonical-8d5ce116bfa7dd86223e04bf5720e3b97e9388b0e42f0156a5b1f185ad76c25d) |
| `cloudfront.protected_endpoints.web_client.continue.no_header` | [cloudfront.protected_endpoints.web_client.continue.no_header](resources--protected_application--reference--group-004.md#canonical-2aa36aeb3554cc761a6a9e600d9ab6304a3fdb9204cafc30281899f7c22fc737) |
| `cloudfront.protected_endpoints.web_client.redirect` | [cloudfront.protected_endpoints.web_client.redirect](resources--protected_application--reference--group-004.md#canonical-5dab7753c34966e046df80b3aeb3eadd7feda8d8ccbd766553c4ed8646ae2d8c) |
| `cloudfront.protected_endpoints.web_client.redirect.location` | [cloudfront.protected_endpoints.web_client.redirect.location](resources--protected_application--reference--group-004.md#canonical-59e6ea4596e9a6568b1f1eb456b5cea1d9728a1e25c4a19d838702030fff3060) |
| `cloudfront.protected_endpoints.web_client.redirect.status` | [cloudfront.protected_endpoints.web_client.redirect.status](resources--protected_application--reference--group-004.md#canonical-4d73f65dccac6a0febe32c364a8348585543c3607a0b15c53ceb9aec090ef2ec) |
| `cloudfront.protected_endpoints.web_mobile_client` | [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-78e9983ea813aa1fbc01c557eb90cd5cac0201c60bdc9d62c95e57f989388d61) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--reference--group-004.md#canonical-355c7d4a112835c13cc51cf72abfd0b138f7aaf1504ac56705e4ea0fe242805c) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.body](resources--protected_application--reference--group-004.md#canonical-c881e7c7a4d6d8ff6a19a8c73f569f1945e3d8b4ece08bab1975b24a5c85609d) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type](resources--protected_application--reference--group-004.md#canonical-3430efca028dc8a8b0994449a73032f0f4c4444a3f77832395b795bbd35567a0) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.status](resources--protected_application--reference--group-004.md#canonical-87fe4d5bf1c041f04e0edffc9c4b7d8e6e04af47b251d41887d6bbbec78e1e5f) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web` | [cloudfront.protected_endpoints.web_mobile_client.block_web](resources--protected_application--reference--group-004.md#canonical-9756ef73f8d4ccdc6d6ffa7990fb0f72525ab5b58203dcae0e0d251cbafa1da7) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.body` | [cloudfront.protected_endpoints.web_mobile_client.block_web.body](resources--protected_application--reference--group-004.md#canonical-b8d2c5cbf66fba6c867f9a43da8a0ea6e2ecff736179de5ab6645df7d42a7bfc) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_web.content_type](resources--protected_application--reference--group-004.md#canonical-70c901cc0271f936a8052d4abf3c35653936143ecced290b537d5837f0d17b47) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.status` | [cloudfront.protected_endpoints.web_mobile_client.block_web.status](resources--protected_application--reference--group-004.md#canonical-c670a6cc6fae35c57f2dc989a7d9ad35ed96ac91f048529c44aa3b7881a37e86) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-004.md#canonical-f0b09b05329b881c2abb55d9fa562424632fab19948b410b07c67d70f8d7ddae) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--reference--group-004.md#canonical-cf833703d156a0b9f7d659fd6a197dc5edbf6169864257626a28a47671c821c7) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--reference--group-004.md#canonical-b05d688663a64d0dcfec0c1b6e82b821843dcea1bc095a19de84dadee386e79b) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web` | [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-004.md#canonical-f2db42c531c317702b6ab42fcc30b6e7d2df4559da1adb1dbea0c55a49f01e87) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--reference--group-004.md#canonical-dc7d707d3fc01e7159ae14954236a310fa52eb9762ed69a5d7d6586bb8693998) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--reference--group-004.md#canonical-40974e0c4df7379efecec45dfd8c0668d481303024a8b3ffdffb470235c91544) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--reference--group-004.md#canonical-93b9a5c554524e95f155fa2a5c9abadec4a8b8424fb1a27e3b2bea0a623ea3ae) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.location](resources--protected_application--reference--group-004.md#canonical-567106317eae5b703adbfb1f479d0f7882c85f0d3b9293b6ddfb765170fc7476) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.status](resources--protected_application--reference--group-004.md#canonical-831fe9fb1f88b7bdb61ab5bc614506d8625a2ae76de26741c2a37bff3a71fffa) |
| `cloudfront.timeout` | [cloudfront.timeout](resources--protected_application--reference--group-002.md#canonical-490a991dff05e269cd674b1ef43bc6cb95b6bd31151ce3c1d57a1029842bd68c) |
| `cloudfront.trusted_clients` | [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-5e6288851f7fb1da6f26e7a1d20eefbd704751e5bc0364c966365561d58b1cbd) |
| `cloudfront.trusted_clients.http_header` | [cloudfront.trusted_clients.http_header](resources--protected_application--reference--group-004.md#canonical-a56ae78fb57a1ba1089d033cc7cb3d0a4a8973513f1d9de3ca2f476f4df13d3d) |
| `cloudfront.trusted_clients.http_header.headers` | [cloudfront.trusted_clients.http_header.headers](resources--protected_application--reference--group-004.md#canonical-337d92f8b6862f67d84de9ddc22b82dbe29ac39ea0de81879db26a8a2945b4a4) |
| `cloudfront.trusted_clients.http_header.headers.exact` | [cloudfront.trusted_clients.http_header.headers.exact](resources--protected_application--reference--group-004.md#canonical-fb61a8222aae52320769c67c25648188c12708300661b8f1f1aa422269bb6872) |
| `cloudfront.trusted_clients.http_header.headers.name` | [cloudfront.trusted_clients.http_header.headers.name](resources--protected_application--reference--group-004.md#canonical-3fde174f2baa92798f502549930cde2a9812b3efe1e86e6eb9f5a706c5fea0a1) |
| `cloudfront.trusted_clients.http_header.headers.regex` | [cloudfront.trusted_clients.http_header.headers.regex](resources--protected_application--reference--group-004.md#canonical-1ab2eb65bc622782eb0f92ad6a2c042f2be1d9ab8044094ba420d4d13a41dfe0) |
| `cloudfront.trusted_clients.ip_prefix` | [cloudfront.trusted_clients.ip_prefix](resources--protected_application--reference--group-004.md#canonical-2b114f1806faa6414877b8ce6332ea856e43463c3451338697d53443365aa663) |
| `cloudfront.trusted_clients.metadata` | [cloudfront.trusted_clients.metadata](resources--protected_application--reference--group-004.md#canonical-3eadabc7bb13231be0c8f4a7cc39932338d7836a73fe444937adf3f39c937127) |
| `cloudfront.trusted_clients.metadata.description_spec` | [cloudfront.trusted_clients.metadata.description_spec](resources--protected_application--reference--group-004.md#canonical-b4206a2f865ae2ede2df301e976ef86a86921d58ec9ba17fa4b718649bc8b442) |
| `cloudfront.trusted_clients.metadata.name` | [cloudfront.trusted_clients.metadata.name](resources--protected_application--reference--group-004.md#canonical-5da5ff0e4c6ed1c20ea7dec93fe2d3987c5642abd0d9cb0c97e17cc2005c92be) |
| `custom_connector` | [custom_connector](resources--protected_application--reference--group-004.md#canonical-7aa10585b9f37614e34760a4e1a8a5437bfa0cb3f8c89cc28da45f215976a824) |
| `description` | [description](resources--protected_application--reference--group-001.md#canonical-343242c033c262941fe0200be051a82eb401a25a15c1f623d52f328bbcf71c87) |
| `disable` | [disable](resources--protected_application--reference--group-001.md#canonical-fd8a1d72d09eb7e196967f28f53a432b12852174c75d95b03b4fbc2d18c758d2) |
| `f5_big_ip` | [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-08cf5923e93d3f704c81fba4e5beb3f6d8d9fa99aeb9e828733fbbf75e44bf1e) |
| `id` | [id](resources--protected_application--reference--group-001.md#canonical-474610d9da6d88b4e10faacb72b941949b532e133a726b2a8f5466527542a61c) |
| `labels` | [labels](resources--protected_application--reference--group-001.md#canonical-30e16e31fe9f1f51dba05ccdec1cef02d71c987ba919197b077dd135fb1a9dec) |
| `name` | [name](resources--protected_application--reference--group-001.md#canonical-5c45cbd15c1de0c1b19c48bc0546ae79f458a1b93d4bdb69b4643a52fa624d2d) |
| `namespace` | [namespace](resources--protected_application--reference--group-001.md#canonical-4bf2001d58108b565aea91862ce0fde7cbc537d0b93f1a74d52b184671e149d9) |
| `region` | [region](resources--protected_application--reference--group-001.md#canonical-911d26eac05cfd0abd1fe59170d3e5296d281df833932c37e87195c246c8f2ae) |
| `salesforce_commerce_connector` | [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-3d84769a6a83f8cdb4493bbe723fe45f7247bdbb72d85de52ba9827eb3ed0bb5) |
| `timeouts` | [timeouts](resources--protected_application--reference--group-004.md#canonical-7058d92adc2561c625a2c5a22bb04e060c8b8e68e6b0c1f6ed1cb748ae27499e) |
| `timeouts.create` | [timeouts.create](resources--protected_application--reference--group-004.md#canonical-9580f11f96e011c089aea5f5cf0674403107c6f0f1973410bc39cf0e6d179ce5) |
| `timeouts.delete` | [timeouts.delete](resources--protected_application--reference--group-004.md#canonical-1b7ce3533ae0424b633140459ca5a2e5659e35ea2246c911e379abae84987f60) |
| `timeouts.read` | [timeouts.read](resources--protected_application--reference--group-004.md#canonical-5c10c88110fbaede537d5184a594e583f93765af4cd899809f9c99e08c0185a1) |
| `timeouts.update` | [timeouts.update](resources--protected_application--reference--group-004.md#canonical-9091fba4ccdaac05861adc5a958998ae608393d30389e90adc737fa6af1e9910) |

<a id="canonical-9e02d8334113f99c203edcf36eaf7d6705038e6242fd7b4758e65fd4fbe7a1b3"></a>

## Next pages — Property reference / d9f064e5a984 / 13

- [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-fe2a351667b89f7b378f73a2056b469b91391f4341f2c2db4e6ba53162e128bb)
- [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-33502cea2351445a45c0ed492ff21356bb1e679d0d4bd40e91eeda5bd42a1048)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [custom_connector](resources--protected_application--reference--group-004.md#canonical-015119427834384a4021619b34741d838615efc7b8a44226c1650dcff1fc786e)
- [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-09aaf1806353d50ca14d6e132e2cb3694e4b277d578cb3e7f1afd534500fa0d9)
- [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-be3e87bf4eaaa0be35a3e894bbcd2e36a0721ae85f61c76da9372a3145122e64)
- [timeouts](resources--protected_application--reference--group-004.md#canonical-795c21823eea76f40b96ac6f5e098873ab1262b372f3bba87d702754c6a10784)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-fe2a351667b89f7b378f73a2056b469b91391f4341f2c2db4e6ba53162e128bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20ad77fc90c034f3a78402fb82f63bd9cd68f4063ffde1cdd1b6ea03ac0404fb"></a>

## adobe_commerce_connector — adobe_commerce_connector / 68dc73c6935f / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- adobe_commerce_connector

<a id="canonical-44ed9fb3a940fcfbacc5846c7a1219a377bf559ca3ee9766ebb2a0434aa7f6f6"></a>

Type: `["object", {}]`. Optional.

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

- [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-44ed9fb3a940fcfbacc5846c7a1219a377bf559ca3ee9766ebb2a0434aa7f6f6)
- [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-7f964c9ca22d58e9d9b04d8d9f1648d340c40b26732d9797e486d37d0d503c83)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-0f81c386d6c92420e2f6257571a58fc5781085d50a24397716e076b967b06085)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3b163d79026cd25596af8a276de79a6ad6020516bd2678119c801b88b42eff3c)
- [custom_connector](resources--protected_application--reference--group-004.md#canonical-7aa10585b9f37614e34760a4e1a8a5437bfa0cb3f8c89cc28da45f215976a824)
- [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-08cf5923e93d3f704c81fba4e5beb3f6d8d9fa99aeb9e828733fbbf75e44bf1e)
- [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-3d84769a6a83f8cdb4493bbe723fe45f7247bdbb72d85de52ba9827eb3ed0bb5)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
adobe_commerce_connector = {}
```

<a id="canonical-e5023aaffcf32d4eb3cb025d85059f0cd11b0858967f3d7beacbc6858cb1ae18"></a>

## Direct properties — adobe_commerce_connector / 68dc73c6935f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a487089448730fad7d9617e5dafe250c2d4361fbdaec259556f54ce9d0bbb23c"></a>

## Next pages — adobe_commerce_connector / 68dc73c6935f / 4

- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-33502cea2351445a45c0ed492ff21356bb1e679d0d4bd40e91eeda5bd42a1048"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7046fe7378e4105611a11e419eb40dcbf0c198f8894ab9a92db720434a2b780a"></a>

## big_ip_iapp — big_ip_iapp / a1f2a03b3aaf / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- big_ip_iapp

<a id="canonical-7f964c9ca22d58e9d9b04d8d9f1648d340c40b26732d9797e486d37d0d503c83"></a>

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
big_ip_iapp = {}
```

<a id="canonical-31ef24ef3f936c13cd95144bf0f2b03acaf492a51bc7a7350ae866f23f397aa5"></a>

## Direct properties — big_ip_iapp / a1f2a03b3aaf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-921dddb7cca4d9d7e31afeb424b26d6d9fc4ae473d025599ac15d35abf181ac9"></a>

## Next pages — big_ip_iapp / a1f2a03b3aaf / 4

- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e7663ad87e7a7ca25c113f2b2239a8be4bf0fd37ee72f9951df434b0c44d2ec"></a>

## cloudflare — cloudflare / 2e086d4921ac / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- cloudflare

<a id="canonical-0f81c386d6c92420e2f6257571a58fc5781085d50a24397716e076b967b06085"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense policy configuration for Cloudflare.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "manual_js_insert"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insertion_rules",
    "manual_js_insert")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
cloudflare {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a08ee0103df1afba9c176ca97ca31f0ffd1ef75d6ba40e64f8ad4e41c609ae4"></a>

## Direct properties — cloudflare / 2e086d4921ac / 3

<a id="canonical-2e694fc6bcb74c3b0d4117c2d52ef283dceea002ade6baba3100dc568295850c"></a>

<a id="canonical-333749376a526f303b24b1844417486648b2d82752ec16cc9300a9e526837268"></a>

## continue_mitigation_action_hdr property — cloudflare / 2e086d4921ac / 4

Type: `"string"`. Optional.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

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

- [disable_js_insert](resources--protected_application--reference--group-001.md#canonical-37a928af93b3dff667da9d36463ddd6e686632c4dd847799228df3dc50cf3e52): complete subsection reference.

- [disable_mobile_sdk](resources--protected_application--reference--group-001.md#canonical-af1f61aa3f4326a04169ba400fab34675fc1dcccd87401a253813f4dc4c89aa2): complete subsection reference.

- [js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551): complete subsection reference.

<a id="canonical-86ee6dbf4c5524b315f697a199bb9df3b56f80db32b352342d4953cf25d08aae"></a>

<a id="canonical-385b940b292c23d7310441fa86fe7682b0d844e1539278dac95232bdbd4c8eb1"></a>

## loglevel property — cloudflare / 2e086d4921ac / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"),
}
```

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

- [manual_js_insert](resources--protected_application--reference--group-001.md#canonical-b2d1f2806c7191e6fb62f6bca969ea019a05f6b29e13708bc6dd8723dbb98a32): complete subsection reference.

- [mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-1a9b3c8ceb69aa8b64bf6361f79136f93fc5abafb2b52a1bb17e2221fe3c0a5d): complete subsection reference.

- [protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0): complete subsection reference.

<a id="canonical-97e012b0d77532681f43aea2c21444db7ac4b9f78b82a118ad4f8496b5741593"></a>

<a id="canonical-643fefdf604b20ec45418ac4245c98b6c094b3e18de1286b5af5bbe293aff6b9"></a>

## timeout property — cloudflare / 2e086d4921ac / 6

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

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

- [trusted_clients](resources--protected_application--reference--group-002.md#canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763): complete subsection reference.

<a id="canonical-a91e58852b97d2b168e3718faaf6ebcd4d1c703998f5d47206d761a85f702118"></a>

## Next pages — cloudflare / 2e086d4921ac / 7

- [cloudflare.disable_js_insert](resources--protected_application--reference--group-001.md#canonical-37a928af93b3dff667da9d36463ddd6e686632c4dd847799228df3dc50cf3e52)
- [cloudflare.disable_mobile_sdk](resources--protected_application--reference--group-001.md#canonical-af1f61aa3f4326a04169ba400fab34675fc1dcccd87401a253813f4dc4c89aa2)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.manual_js_insert](resources--protected_application--reference--group-001.md#canonical-b2d1f2806c7191e6fb62f6bca969ea019a05f6b29e13708bc6dd8723dbb98a32)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-1a9b3c8ceb69aa8b64bf6361f79136f93fc5abafb2b52a1bb17e2221fe3c0a5d)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-969690ac1edf20e3cb49b9a391b479ee78b6c6841ec5a89abba8ba8166920763)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-37a928af93b3dff667da9d36463ddd6e686632c4dd847799228df3dc50cf3e52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad2d1d24ffccc23ddb8ad35f812daeb448fe7061b339e95a07412fa8eaa95438"></a>

## cloudflare.disable_js_insert — cloudflare.disable_js_insert / 9b53cfefaa22 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- cloudflare.disable_js_insert

<a id="canonical-1c81f87475c4560e27937b2a1cdba2c3391e795c221f4925a646452491b7da33"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_js_insert = {}
```

<a id="canonical-6b3a7401de855807420cb0f4c5231d97328955ce9743a4dc1d6c980e54674d79"></a>

## Direct properties — cloudflare.disable_js_insert / 9b53cfefaa22 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e47cc966d1116d75577b7d7f48bfa9bf178ebce0968a69b6d313e116bd67773f"></a>

## Next pages — cloudflare.disable_js_insert / 9b53cfefaa22 / 4

- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-af1f61aa3f4326a04169ba400fab34675fc1dcccd87401a253813f4dc4c89aa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f03f7b97ac3304743ae22569f9e77dc28debf57b2a0df5dcd030b56c77a66b10"></a>

## cloudflare.disable_mobile_sdk — cloudflare.disable_mobile_sdk / 07d0d1038805 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- cloudflare.disable_mobile_sdk

<a id="canonical-c8cdce542cb64c8b475dd745e742575fc9b7a62ac7d1b465b8cff4e925e52206"></a>

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
disable_mobile_sdk = {}
```

<a id="canonical-a45a444d91dbac279f99cf3788bbefca9e7db304ec7ed1a7bf4dbd1e9276ded5"></a>

## Direct properties — cloudflare.disable_mobile_sdk / 07d0d1038805 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72a27479571f688e97752442654576e61b9331cc26a7ad8f3588aea020252119"></a>

## Next pages — cloudflare.disable_mobile_sdk / 07d0d1038805 / 4

- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e136a7e2286bb21289881c2ebe53690f390facda85a18bd6aa6e04ea665ccc61"></a>

## cloudflare.js_insertion_rules — cloudflare.js_insertion_rules / 15c6b5f4907b / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- cloudflare.js_insertion_rules

<a id="canonical-2aee71ee08354bc82a2ab095cab18af6c18e59ed8d88303fcb5aecc7a0b8084e"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c7cdca2fc0ce493c0b595961436d5b0678db93d55cefb3a5178f5a0a079809b"></a>

## Direct properties — cloudflare.js_insertion_rules / 15c6b5f4907b / 3

- [exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2): complete subsection reference.

<a id="canonical-c24fa9337bdd62b818ed3b4aff18ef636afed1104c1fd92374497525e4098c1c"></a>

<a id="canonical-1ccfdacc8e4bbfb9a80d9a5f1a4c078557f9fc89048be0c44c4ae9e9ef25fec6"></a>

## javascript_location property — cloudflare.js_insertion_rules / 15c6b5f4907b / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

<a id="canonical-5ef59e573155e8b914c9debfc83c05d39e98e37d2ede39c02c79612c048def6a"></a>

<a id="canonical-26d79f71597e996c9e66587c5315609c71cc227748a05ca2417ba81fd5f21151"></a>

## js_download_path property — cloudflare.js_insertion_rules / 15c6b5f4907b / 5

Type: `"string"`. Optional.

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

- [rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149): complete subsection reference.

<a id="canonical-5fdb753325381ad4179c5d145737422b1db4941e89fdadaf612ac443be069c98"></a>

## Next pages — cloudflare.js_insertion_rules / 15c6b5f4907b / 6

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0a60741c3613fb3a4d00ca77e21a1f47a337c2fafab1daae15296cff088e006"></a>

## cloudflare.js_insertion_rules.exclude_list — cloudflare.js_insertion_rules.exclude_list / 10da7e588ddb / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- cloudflare.js_insertion_rules.exclude_list

<a id="canonical-3c9f030f54321fdb742dca1c059f1e8ba861b6cf2c4f2b8dce97a896774b8ff0"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd345260ab1401aeb1b1120b4e6b2130f8b6d3a89532f20c175bd24ac38ab784"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list / 10da7e588ddb / 3

- [any_domain](resources--protected_application--reference--group-001.md#canonical-19877f357976e4a4467e6c7ecd0b6eb4b581183dd5c02d776bb5670bd12787b3): complete subsection reference.

- [domain](resources--protected_application--reference--group-001.md#canonical-2606618787331be72f8b76768242b5a9ca3e9680a5db21d3702acc9d51fab31a): complete subsection reference.

- [metadata](resources--protected_application--reference--group-001.md#canonical-bdc5e52547fe4732d7d2713aa0951989adf47d6c83214ce1df932ad980e28cca): complete subsection reference.

- [path](resources--protected_application--reference--group-001.md#canonical-1a80f23f6f8d9d19f015ccea49335baf2e34610cc1f10f7c423c7894f3a9662d): complete subsection reference.

<a id="canonical-495f7e6ea3a7beca4f0d90f85aecf3bca2afef065beb46cffecec8a4dff5806d"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list / 10da7e588ddb / 4

- [cloudflare.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-001.md#canonical-19877f357976e4a4467e6c7ecd0b6eb4b581183dd5c02d776bb5670bd12787b3)
- [cloudflare.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-001.md#canonical-2606618787331be72f8b76768242b5a9ca3e9680a5db21d3702acc9d51fab31a)
- [cloudflare.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-001.md#canonical-bdc5e52547fe4732d7d2713aa0951989adf47d6c83214ce1df932ad980e28cca)
- [cloudflare.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-001.md#canonical-1a80f23f6f8d9d19f015ccea49335baf2e34610cc1f10f7c423c7894f3a9662d)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-19877f357976e4a4467e6c7ecd0b6eb4b581183dd5c02d776bb5670bd12787b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69e2e6e3433a6146484b8dc4222352a89c1da7b4c7d6bb5b85bec1ec4a1b99f7"></a>

## cloudflare.js_insertion_rules.exclude_list.any_domain — cloudflare.js_insertion_rules.exclude_list.any_domain / ff622110398e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- cloudflare.js_insertion_rules.exclude_list.any_domain

<a id="canonical-8df15559495c8d3bb36f47bcc4a70c29018c6edadd494a9de5b97a8a4c218a61"></a>

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
any_domain = {}
```

<a id="canonical-08e4d1a8cbf04a2ee2bec7789f322a1dbc950a0f4cbbf07eedf106b6a23b00dd"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.any_domain / ff622110398e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1bc423f278b44f6eb2c43c38f46f5f99db40a985547cb6259618f1f68a81adc0"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.any_domain / ff622110398e / 4

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-2606618787331be72f8b76768242b5a9ca3e9680a5db21d3702acc9d51fab31a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bd07114807cf58dcb0469ab7409581bd5044124a577694219d96481fc8cac62"></a>

## cloudflare.js_insertion_rules.exclude_list.domain — cloudflare.js_insertion_rules.exclude_list.domain / d30a0ddfbe11 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- cloudflare.js_insertion_rules.exclude_list.domain

<a id="canonical-cf0c2079ce2ca36f9d0412e495ee16a716436efab09271d8d617d1169872db01"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-b9691823ac03b84f4079b531620ee0c02c3fa5c3da549b7a578fd3ef88706f0a"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.domain / d30a0ddfbe11 / 3

<a id="canonical-7c6eefdacd57cdee5fcdb12724c415fe45ecd535ead39981b274d7b9babb1241"></a>

<a id="canonical-2f09f216f2778535c848c240cfb7a71a182ac36c3b45df2f585b0b0b319efd52"></a>

## exact_value property — cloudflare.js_insertion_rules.exclude_list.domain / d30a0ddfbe11 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-2e0f77fbad3f395d1dcbfa7350c6d20cf9d58944e0f5ec0684b401ef71a372cd"></a>

<a id="canonical-df67d9af9a7edf7412cd8283dd420ec4b4e95f858546fb8c75292b6d0a135f46"></a>

## regex_value property — cloudflare.js_insertion_rules.exclude_list.domain / d30a0ddfbe11 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-b64730264974c306efb8bf83f8852737267a01efc6987689cf64e70f67c23426"></a>

<a id="canonical-7f1a6f55db4fc76be939c0d7db8af0ddf4e8fadeef2d223071edf010b057a5dc"></a>

## suffix_value property — cloudflare.js_insertion_rules.exclude_list.domain / d30a0ddfbe11 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-763fb919c50945e5ec666d2b0dc1e79ae5c1bd40fc3faeb9f3713703001acc24"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.domain / d30a0ddfbe11 / 7

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-bdc5e52547fe4732d7d2713aa0951989adf47d6c83214ce1df932ad980e28cca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4c394fc429e3f72db974f86231910b6542f9c77117f887d996117b8c4dd4ee3"></a>

## cloudflare.js_insertion_rules.exclude_list.metadata — cloudflare.js_insertion_rules.exclude_list.metadata / e3ad0e463318 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- cloudflare.js_insertion_rules.exclude_list.metadata

<a id="canonical-b93aa4a8334e458e823387085d6e25429bb71aed5e5d3abaed2e28e097153c20"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3201003ecf9ed3b2644a3b6d426f70ac7fb2cb8ef28d9e14e4f2e8fcea46ef6"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.metadata / e3ad0e463318 / 3

<a id="canonical-790528c11968367eba673346d4976ea2c247ba29618c0c64b953e61808f651db"></a>

<a id="canonical-c40732589dc5762533113ca9fb6e77831753c9a29667bbd9745ba96b8e7b3222"></a>

## description_spec property — cloudflare.js_insertion_rules.exclude_list.metadata / e3ad0e463318 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-ee7815a81de32560b127894af0d239e431abbe922c2c033c26a10098386bc12c"></a>

<a id="canonical-898596831df6ec64d4380d058f0b8936fd787fb5fdba7c1acc0fd257bd2156d5"></a>

## name property — cloudflare.js_insertion_rules.exclude_list.metadata / e3ad0e463318 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-9ea9989a7088178d2e90dc1243cb0ac5b567ea040eba315caeafb5cf3b2f6a0f"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.metadata / e3ad0e463318 / 6

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-1a80f23f6f8d9d19f015ccea49335baf2e34610cc1f10f7c423c7894f3a9662d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ce9f1317d8275bfdbba1c3ed6c4aa7ce90d293399a46768c8118bc9608ee06c"></a>

## cloudflare.js_insertion_rules.exclude_list.path — cloudflare.js_insertion_rules.exclude_list.path / 62357c5bb2cb / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- cloudflare.js_insertion_rules.exclude_list.path

<a id="canonical-9e6ad2692f11c619ee9836f647fe74cda6c91144a457f43b469fcb0ae20b05f2"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-994bf2319be9b47e03b9a1c2a550259d326105a50ce63b24125460dcdbeff896"></a>

## Direct properties — cloudflare.js_insertion_rules.exclude_list.path / 62357c5bb2cb / 3

<a id="canonical-902f6e3b96e9a84a7c27b396340818c9053caa24b0b8004000fbd59cf94a9df9"></a>

<a id="canonical-2c33a4078a449fa09be35c065513abf41599fa5593281ac797fbfdb5d79c27d8"></a>

## path property — cloudflare.js_insertion_rules.exclude_list.path / 62357c5bb2cb / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-8056e800537f36720fe1473b3a3f4395bb1ce1aeee07f6eeb63c13f2524d5513"></a>

<a id="canonical-b8a8d6147af3db0208b7f293730e8e72b60fb31a2d4bd11b65ded81054336246"></a>

## prefix property — cloudflare.js_insertion_rules.exclude_list.path / 62357c5bb2cb / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-4897b90c8dcc5b4e5fac828ca7a37a74684366cae074b89131fed2d4cd190fe4"></a>

<a id="canonical-2e514126c36d8eef720ce7b574f50882b38cab8cfecd435bc313ad1add46b24a"></a>

## regex property — cloudflare.js_insertion_rules.exclude_list.path / 62357c5bb2cb / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-51c73b42adf79fc18dae0ed726f1b303101e828d7c5a966683a5692ef7301a0f"></a>

## Next pages — cloudflare.js_insertion_rules.exclude_list.path / 62357c5bb2cb / 7

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-6a2f36a1bb6bcdfc2f12b3b54b33f7a841ccda02fa8414b88c2e807893ba0df2)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b45412bfd2a416bcb528c229d2ef71c6a82a6c719bc8bacc193a9a8f6d257f8e"></a>

## cloudflare.js_insertion_rules.rules — cloudflare.js_insertion_rules.rules / 9087d69a97ef / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- cloudflare.js_insertion_rules.rules

<a id="canonical-27b56b2c81717809b9dfa802c3fb059cc5ec3b19762e31619ae7e3b9fcebf82d"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("exact_path",
    "glob"),
  validators.ConflictingListObjectAttributes("exact_path",
    "prefix"),
  validators.ConflictingListObjectAttributes("glob",
    "prefix")}
```

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-942bcdc62acaa69fb2aaf0787c1c3e51e73e0a618c1552cc948f3cbc0cfc366e"></a>

## Direct properties — cloudflare.js_insertion_rules.rules / 9087d69a97ef / 3

- [any_domain](resources--protected_application--reference--group-001.md#canonical-0b27bcd3105ad28adf5f139a73cd936e2e802c2fe0726155f67580cbbf7928a8): complete subsection reference.

- [domain](resources--protected_application--reference--group-001.md#canonical-4be657c0bbd196ea31e70410e3056bd18825ec06846fb031ee645af0f78c229d): complete subsection reference.

<a id="canonical-3fe103456f07c4aeb3fe8fc25ba64fc3987db4889c7cd910ed6359abf42724ef"></a>

<a id="canonical-0c1c6ca727150525461dd6b01c4285cd2fc879573324da93c9622be68233c539"></a>

## exact_path property — cloudflare.js_insertion_rules.rules / 9087d69a97ef / 4

Type: `"string"`. Optional.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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

<a id="canonical-b06e0855b3f2f59c4f1ce1ab6c2dfd66940c4832ee6e6d47835c624b1b505b74"></a>

<a id="canonical-b9abade2fe1d1b40a3610957e8b260d9dbe78861c7180fa086f7d3d402a2371f"></a>

## glob property — cloudflare.js_insertion_rules.rules / 9087d69a97ef / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

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

- [metadata](resources--protected_application--reference--group-001.md#canonical-abb1294bfa89691a05f8f74a11086baa61b8c43b3628eb2d5568437c2bf9e5db): complete subsection reference.

<a id="canonical-4ded4c4a0f9ff306de7967ded5d6864106af1da180087754e5f1ff94d220ac6d"></a>

<a id="canonical-31e2b888ee9236cded4c60116af2c2abd9733bae46c75a6fdab141167cc9a433"></a>

## prefix property — cloudflare.js_insertion_rules.rules / 9087d69a97ef / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-d028ef8e9104228de1ac33c8908987268206741e003ba09395c71e6dac53b055"></a>

## Next pages — cloudflare.js_insertion_rules.rules / 9087d69a97ef / 7

- [cloudflare.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-001.md#canonical-0b27bcd3105ad28adf5f139a73cd936e2e802c2fe0726155f67580cbbf7928a8)
- [cloudflare.js_insertion_rules.rules.domain](resources--protected_application--reference--group-001.md#canonical-4be657c0bbd196ea31e70410e3056bd18825ec06846fb031ee645af0f78c229d)
- [cloudflare.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-001.md#canonical-abb1294bfa89691a05f8f74a11086baa61b8c43b3628eb2d5568437c2bf9e5db)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-0b27bcd3105ad28adf5f139a73cd936e2e802c2fe0726155f67580cbbf7928a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa63c67264f85eab897776ac94f3f0a26c1d96475b5140d581c68cbd1a07580d"></a>

## cloudflare.js_insertion_rules.rules.any_domain — cloudflare.js_insertion_rules.rules.any_domain / a51bde8022a9 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149)
- cloudflare.js_insertion_rules.rules.any_domain

<a id="canonical-c05ea790d972408c8646b75b83f681b72383ad883bcbabcc3036b8b76dcad8d6"></a>

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
any_domain = {}
```

<a id="canonical-2db6f122c16e3c02843148bc8b838885ef24e0d92e595a2a74b01c7b28d9c98a"></a>

## Direct properties — cloudflare.js_insertion_rules.rules.any_domain / a51bde8022a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19763f35df789cbbbdec46386d9e4f1e9e92179d2f5c2d1f1cf222865cd75bcb"></a>

## Next pages — cloudflare.js_insertion_rules.rules.any_domain / a51bde8022a9 / 4

- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-4be657c0bbd196ea31e70410e3056bd18825ec06846fb031ee645af0f78c229d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c21ffd3607d10bb0fbd162474c7025e4441372869c611837305e66fa87e584c1"></a>

## cloudflare.js_insertion_rules.rules.domain — cloudflare.js_insertion_rules.rules.domain / f55ece7442a0 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149)
- cloudflare.js_insertion_rules.rules.domain

<a id="canonical-7621e46460f522d606a9d10ff86164c4936ae0fc6623ca129bbef951ee5eba01"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-e413384a58564ebaf6aae00eabbbf5b1a4e2ce182af9f4b8becc1b8f0d104235"></a>

## Direct properties — cloudflare.js_insertion_rules.rules.domain / f55ece7442a0 / 3

<a id="canonical-755d3a106ed14187b9ed917409cdc2c6c83667ff0c44c44d19181ac976f665f1"></a>

<a id="canonical-509a362aa561101e538c8dea061b844a249a3bc2220698d23cbcaaa9118e65ef"></a>

## exact_value property — cloudflare.js_insertion_rules.rules.domain / f55ece7442a0 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-e3fbd3e600e13afa87bb191ea1c13fcb255b96312da6b79564ea0ead23cb667c"></a>

<a id="canonical-3556aa23d587f537f26e4f242380e2407e3f1dd337518e2212946616aae7571f"></a>

## regex_value property — cloudflare.js_insertion_rules.rules.domain / f55ece7442a0 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-9d69707d50a189a1032ac25bb5c81e95b098bed4b09fef57d0022f728482ed7d"></a>

<a id="canonical-0b8c0d1cb4dc94928c4095bcf3a414691065c1a9480a1d4a244d5240e1144e15"></a>

## suffix_value property — cloudflare.js_insertion_rules.rules.domain / f55ece7442a0 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-a984ab9918df3d6517acf7c2a27be0c006a91d861bb76e373b8750d6e2071c10"></a>

## Next pages — cloudflare.js_insertion_rules.rules.domain / f55ece7442a0 / 7

- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-abb1294bfa89691a05f8f74a11086baa61b8c43b3628eb2d5568437c2bf9e5db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f15cc788e2fd06b890f1dfc78c0c424fd1b1906dee6a60dafa8db1885f80631"></a>

## cloudflare.js_insertion_rules.rules.metadata — cloudflare.js_insertion_rules.rules.metadata / d88e358d86a4 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-7a089b70c4c7f64d7a5064d688166d389aa5326265ffd1d9e2b71d0addc7b551)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149)
- cloudflare.js_insertion_rules.rules.metadata

<a id="canonical-908a4ccc3956dcbf13151f85e7d6f9ff161cb08e99bba25c9e284f6669a42eb1"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-a9444881dd1c69be55ac24d24064f0c75b688e03bd960526babeec1eb3b1a6a0"></a>

## Direct properties — cloudflare.js_insertion_rules.rules.metadata / d88e358d86a4 / 3

<a id="canonical-b23f3d9742c11a43010df50653ba882d9e84f14a569b11556413c01397bb8275"></a>

<a id="canonical-9bd72666e2f0d635f108eda664467ef6e4830398ddeda049687f3275102ddc77"></a>

## description_spec property — cloudflare.js_insertion_rules.rules.metadata / d88e358d86a4 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-e09ba90830db68c26db84de0ed5e148b15cc7a0ea48c4a142523adf21568cd14"></a>

<a id="canonical-d938938a7574d0ca72a0a3e5da876923478f8f12ccd36bbbd907887a11ff6639"></a>

## name property — cloudflare.js_insertion_rules.rules.metadata / d88e358d86a4 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-db1975109eca8b2e5b821b9f54a223aad368da9f19f5b2a80b660f19c3963e24"></a>

## Next pages — cloudflare.js_insertion_rules.rules.metadata / d88e358d86a4 / 6

- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-bc38c968b943fdaab36244592199887c4bf5bf4abac28ddf68cad68941cb5149)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-b2d1f2806c7191e6fb62f6bca969ea019a05f6b29e13708bc6dd8723dbb98a32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1275f11e40822a00938acc9c2f475bda3730708af3088d7298e11421ca608b7a"></a>

## cloudflare.manual_js_insert — cloudflare.manual_js_insert / 07c5dd375289 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- cloudflare.manual_js_insert

<a id="canonical-a4819691f510f124a0cac4f50f52d74818e5fa6bd910bcd3be3772b6013e0658"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
manual_js_insert {
  # Configure direct properties listed below.
}
```

<a id="canonical-bfaa94394b5c1f3a91a479dc6e9ae08c18cf64e149eb84edd4d864445f704d95"></a>

## Direct properties — cloudflare.manual_js_insert / 07c5dd375289 / 3

<a id="canonical-d0b9e40758c071b22b6f6d6b4fce8148ebeabfffdada5dff2bd261dd638ed054"></a>

<a id="canonical-f904f2d6e406df30591306f2c270c8b6eac7f195261b5ad44ba327310918099e"></a>

## js_download_path property — cloudflare.manual_js_insert / 07c5dd375289 / 4

Type: `"string"`. Optional.

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

<a id="canonical-64335b2753ccd7de61c16aa3648b99f04fe9ab6e01bed6675c84bb2d3ad6d6d2"></a>

## Next pages — cloudflare.manual_js_insert / 07c5dd375289 / 5

- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-1a9b3c8ceb69aa8b64bf6361f79136f93fc5abafb2b52a1bb17e2221fe3c0a5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2255d1e226e8be2081bc9e902227c5570da472747bc01167a95aeb8d6c99206"></a>

## cloudflare.mobile_sdk_config — cloudflare.mobile_sdk_config / 4f469f2e94db / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- cloudflare.mobile_sdk_config

<a id="canonical-272b23d770180fa02d6dd9597871ff74b653182416c08a8bec5b4fe62003a4ab"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-bbb1afe655da40db38db703cd1209a6d0ff6121575a01158664cdc8ce4f44da6"></a>

## Direct properties — cloudflare.mobile_sdk_config / 4f469f2e94db / 3

- [mobile_identifier](resources--protected_application--reference--group-001.md#canonical-542a451f3268fc7d6e3929b75abbd4cdce8a6a898031499408f4978615654b31): complete subsection reference.

<a id="canonical-84b550c03a638c64d805f1a0926b948cb5beeebeeaa1b1f355211f2259131415"></a>

## Next pages — cloudflare.mobile_sdk_config / 4f469f2e94db / 4

- [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-542a451f3268fc7d6e3929b75abbd4cdce8a6a898031499408f4978615654b31)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-542a451f3268fc7d6e3929b75abbd4cdce8a6a898031499408f4978615654b31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-567c1aee4529eb9e88c39cb588548bf335b9663d8b321499b6013c6860c540df"></a>

## cloudflare.mobile_sdk_config.mobile_identifier — cloudflare.mobile_sdk_config.mobile_identifier / affa4cf2ac4e / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-1a9b3c8ceb69aa8b64bf6361f79136f93fc5abafb2b52a1bb17e2221fe3c0a5d)
- cloudflare.mobile_sdk_config.mobile_identifier

<a id="canonical-49d861b862a269662521ebe9a145c25676ef2bc12d83b0284d05e46d338f6f5f"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-ede513358604a682ad39fe1cb85245c77f1dad074357f4d450a12accccf54d32"></a>

## Direct properties — cloudflare.mobile_sdk_config.mobile_identifier / affa4cf2ac4e / 3

- [headers](resources--protected_application--reference--group-001.md#canonical-623dab63d5c690a30dc47d3d3448e1efe3ee3b9051907c6e9339620d0940b3fe): complete subsection reference.

<a id="canonical-534fe348dff8fad6f2a5a134fb738e639604549b569fe7eaaa9d58d811247dd4"></a>

## Next pages — cloudflare.mobile_sdk_config.mobile_identifier / affa4cf2ac4e / 4

- [cloudflare.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-001.md#canonical-623dab63d5c690a30dc47d3d3448e1efe3ee3b9051907c6e9339620d0940b3fe)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-1a9b3c8ceb69aa8b64bf6361f79136f93fc5abafb2b52a1bb17e2221fe3c0a5d)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-623dab63d5c690a30dc47d3d3448e1efe3ee3b9051907c6e9339620d0940b3fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c421b409359a9a3978c4737f70d93a0bf5cf4763d76abe262808c916f44b38f4"></a>

## cloudflare.mobile_sdk_config.mobile_identifier.headers — cloudflare.mobile_sdk_config.mobile_identifier.headers / ecad74849a54 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-1a9b3c8ceb69aa8b64bf6361f79136f93fc5abafb2b52a1bb17e2221fe3c0a5d)
- [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-542a451f3268fc7d6e3929b75abbd4cdce8a6a898031499408f4978615654b31)
- cloudflare.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-61b4c4262f75d13221b69bf075aeb7862644c32244b29cf7421b46eecbaf5add"></a>

Type: `"object"`. list nested block, Optional.

List of headers that can be used to identify mobile traffic.

Upstream description:

A list of headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "regex")}
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-e206a16aa875a009c36e41c4a2915d7a79c22433a7212bea69bc60c891306de0"></a>

## Direct properties — cloudflare.mobile_sdk_config.mobile_identifier.headers / ecad74849a54 / 3

<a id="canonical-c47c0f4bbf63c3662886ec5e0f2eee1091cd7f6d06a99ea17f1529e45945e6f9"></a>

<a id="canonical-3d06c9cb05b9e6f0e1685d753f0dafd679d631bd39551fd287b003ddb7149ea6"></a>

## exact property — cloudflare.mobile_sdk_config.mobile_identifier.headers / ecad74849a54 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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

<a id="canonical-f9424bc143c6dffd36b0a5474d8eccb9f05f6dd5050f02b312ada7260b06f177"></a>

<a id="canonical-fc6e9061265db82e283d9d8831b771f72b80db4b86aa3b0e250e64f687c974b1"></a>

## name property — cloudflare.mobile_sdk_config.mobile_identifier.headers / ecad74849a54 / 5

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-97aa85cbb4258bcf8b234e17d93af49b1f2010b330c02221a212e81797c12e26"></a>

<a id="canonical-4aa774a8f4aac74b17feb2db5dfeae4772abb34d724c436d103f111392b288b1"></a>

## regex property — cloudflare.mobile_sdk_config.mobile_identifier.headers / ecad74849a54 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

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

<a id="canonical-6ed35a81da78cbe7a9ef78df6a68ca214af0a8be2ead3ac9f3a523a824b9404b"></a>

## Next pages — cloudflare.mobile_sdk_config.mobile_identifier.headers / ecad74849a54 / 7

- [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-542a451f3268fc7d6e3929b75abbd4cdce8a6a898031499408f4978615654b31)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8532bfe60f8e0e87b0ceab0f70286e95eceb6403a25db533b7fc22ad5d86e09f"></a>

## cloudflare.protected_endpoints — cloudflare.protected_endpoints / 56d45472408a / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- cloudflare.protected_endpoints

<a id="canonical-e63c920461d5cfdb3bb29b5c64453dad91344c21bfb0dad14c9731e318f14be1"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_client"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_mobile_client"),
  validators.ConflictingListObjectAttributes("web_client",
    "web_mobile_client")}
```

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

Terraform syntax:

```terraform
protected_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-9390269d8806574fdc74b3f52769429a34e1474ce9e6aa6a2fe8d54c0e25e62c"></a>

## Direct properties — cloudflare.protected_endpoints / 56d45472408a / 3

- [any_domain](resources--protected_application--reference--group-001.md#canonical-4fd609e1fe3e97559a46218c28fafea298943bf5e7ce4fe6a30c028877033dae): complete subsection reference.

- [domain](resources--protected_application--reference--group-001.md#canonical-caff12fe88a6569ea8f703c2e4b1aeda31adc3c254bb84e09e974ae13b0317fd): complete subsection reference.

<a id="canonical-8cfda0e6b89c65290096330f84fe3879691cfd7557f405f0d4838f5f4eacf676"></a>

<a id="canonical-1df85cd9fecd6617cc946a935ec0619f66d636e84b7e179b4dde60e69048a0a9"></a>

## http_methods property — cloudflare.protected_endpoints / 56d45472408a / 4

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

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

- [metadata](resources--protected_application--reference--group-002.md#canonical-371da26b9e490b11558667c8d31d2c8e0426ba3d621105d0c366ad9fb3d40efb): complete subsection reference.

- [mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf): complete subsection reference.

- [path](resources--protected_application--reference--group-002.md#canonical-41c264bde6f1257bd08480e52e48086ca57e3355de7ceceff66d8eca33a99e66): complete subsection reference.

<a id="canonical-cf957d83019359eca0220ebeea1bd0d9c086c5a36a3125e75edc6e9b15077e7b"></a>

<a id="canonical-a69fb3d747e9c55efebe5f22a04440627bbbd92a0a2c885e82b14561e74e78ab"></a>

## query property — cloudflare.protected_endpoints / 56d45472408a / 5

Type: `"string"`. Optional.

Enter a regular expression to match your query parameters of interest.

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

- [web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e): complete subsection reference.

- [web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68): complete subsection reference.

<a id="canonical-5a19ac34bc15ba6e0767cd2cfb389034a39f6404e6550bcf8fa1475219e0e62c"></a>

## Next pages — cloudflare.protected_endpoints / 56d45472408a / 6

- [cloudflare.protected_endpoints.any_domain](resources--protected_application--reference--group-001.md#canonical-4fd609e1fe3e97559a46218c28fafea298943bf5e7ce4fe6a30c028877033dae)
- [cloudflare.protected_endpoints.domain](resources--protected_application--reference--group-001.md#canonical-caff12fe88a6569ea8f703c2e4b1aeda31adc3c254bb84e09e974ae13b0317fd)
- [cloudflare.protected_endpoints.metadata](resources--protected_application--reference--group-002.md#canonical-371da26b9e490b11558667c8d31d2c8e0426ba3d621105d0c366ad9fb3d40efb)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-ad8c26be01537d3bdbddaf3b1f6d7c31320adf4c5ec833f75af7d569b876d1cf)
- [cloudflare.protected_endpoints.path](resources--protected_application--reference--group-002.md#canonical-41c264bde6f1257bd08480e52e48086ca57e3355de7ceceff66d8eca33a99e66)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-278aebc500da600797402b54b10019ee0d03771d7f674de2a1e6c3491b2e617e)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-57d6ee3f30611e2394ad25a78765f950f961ec65880ed35f7bd6ec734a227d68)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-4fd609e1fe3e97559a46218c28fafea298943bf5e7ce4fe6a30c028877033dae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bee2801065a153bc0cb45c01b8f66d738a1f4632cc6ed1487530e9dbc7f7900"></a>

## cloudflare.protected_endpoints.any_domain — cloudflare.protected_endpoints.any_domain / 19e41968bfaa / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- cloudflare.protected_endpoints.any_domain

<a id="canonical-1749db8d76ec60ba5ece1869d63fc37cba6e0625c4acf127a2853a6230eb2cd8"></a>

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
any_domain = {}
```

<a id="canonical-fa31e8b020da627165ce16e0c098f32df6da9122657bfc524d0c7917a0446ae2"></a>

## Direct properties — cloudflare.protected_endpoints.any_domain / 19e41968bfaa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c112f38702a5ede2f13ae2c2ccf6bdf4874301de1fa204e708d48671c48512b8"></a>

## Next pages — cloudflare.protected_endpoints.any_domain / 19e41968bfaa / 4

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-caff12fe88a6569ea8f703c2e4b1aeda31adc3c254bb84e09e974ae13b0317fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13ffebbce3d8838bfcaa749f6cdac4a6dda975a11e74b4f4a56ba131f616b8d1"></a>

## cloudflare.protected_endpoints.domain — cloudflare.protected_endpoints.domain / 9101070ad78b / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudflare](resources--protected_application--reference--group-001.md#canonical-709db75a4cd6bc845ba429b05b6a712054f6faf4e53e8360ffcb060c083f0355)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-dac09ed0ace08f88f0738ea13321225b58278a1cdf01e9571fdc3aa4968ea3a0)
- cloudflare.protected_endpoints.domain

<a id="canonical-7fc8abaf98c3453fa81932828441c6d87b18664502c89c1a2b3efc02a4ca773a"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe06816683f505c4c0e2d61393e19d5ce102a3cc370272c8dd882e5a22f3afd1"></a>

## Direct properties — cloudflare.protected_endpoints.domain / 9101070ad78b / 3

<a id="canonical-d11db7cb06a6b8382178debb9b96ee2dd4ceebe19e32eb73b03e0305008c65ab"></a>

<a id="canonical-b91e970eae07b442a4dcca548b5ff6ea12d17f98393f1cc7eeef5d58f2b65e95"></a>

## exact_value property — cloudflare.protected_endpoints.domain / 9101070ad78b / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-e58e6a544fac15e1766ba2579348c1945373bb9be5d6c4b7795f2047d17c9a47"></a>
