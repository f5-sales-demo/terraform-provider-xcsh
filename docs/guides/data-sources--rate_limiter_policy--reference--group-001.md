---
page_title: "xcsh_rate_limiter_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy reference."
---

# xcsh_rate_limiter_policy reference

<a id="canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a05d7636d5204c39764218ba383f694b756da5b777c2ebaf05c3cd29f160b555"></a>

## Property reference — Property reference / 52d63fd75df3 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- Property reference

<a id="canonical-054d8dcb73292901cc64ce7fa0ff45994f59f2997321edb8987c596565312846"></a>

## Direct properties — Property reference / 52d63fd75df3 / 3

<a id="canonical-dc1af2b4da93859c7e3cfa0b006f058138fd14608691f65c84ddda61fb5e3060"></a>

<a id="canonical-d904227750683b6c8bf646d89557a802d2e26715e4b86c38bacb9ffe8747d1bf"></a>

## annotations property — Property reference / 52d63fd75df3 / 4

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

- [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-99007c6de9f1e79410e4b7e6abc47b52b4a496329e04f2ad5d363fdd55c265f9): complete subsection reference.

<a id="canonical-99b2a41a14815005605d1301899ab177104f3063ba3245bb9f16ea3fcd3f2453"></a>

<a id="canonical-6a432074f4c783b5d599a4cca4167acdf86d16758ab7051d19124e25db2056ca"></a>

## description property — Property reference / 52d63fd75df3 / 5

Type: `"string"`. Computed.

Description of the RateLimiterPolicy.

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

<a id="canonical-2ba1d430d0ef94bcff9f428ba3d71f1a3274f326340a4ec592ea3b062fe40631"></a>

<a id="canonical-0c59bc35104d09af30e218785c692b76d177edc92dc060ac87dc2594376edff9"></a>

## id property — Property reference / 52d63fd75df3 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-e5be53631499f7d3dc6c55768d1b6f65d26a7cb986b2c345588e56ab6f79694d"></a>

<a id="canonical-ad0f5e794167530e678876cf3c9da6cbd3cee2ba6059cd212194ff00433b57e7"></a>

## labels property — Property reference / 52d63fd75df3 / 7

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

<a id="canonical-0b84b357802e1d431d6db41795dc2f9f246127aa0c0392c9f928338b9dfa976c"></a>

<a id="canonical-2211d1d54a48ce63a34fad2ddb3c2e36de408d63288067926c9af750caba3d47"></a>

## name property — Property reference / 52d63fd75df3 / 8

Type: `"string"`. Required.

Name of the RateLimiterPolicy.

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

<a id="canonical-85e24df12c278fe56d90885e689c268e53c4ca360c2146a72b865dff332df80e"></a>

<a id="canonical-38b69d017b2c46af6389b0cf6ca0372f48a0efc46c4bff727b384cac81aa148d"></a>

## namespace property — Property reference / 52d63fd75df3 / 9

Type: `"string"`. Required.

Namespace where the RateLimiterPolicy exists.

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

- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745): complete subsection reference.

<a id="canonical-5d7f5c770a159307202cf728091de193d8946519a8b57c319f06e8c4371f4a1a"></a>

<a id="canonical-23de934df583f90f990ed3efbbe5b219680084e3f8732d24da665e3905a570c6"></a>

## server_name property — Property reference / 52d63fd75df3 / 10

Type: `"string"`. Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

Upstream description:

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-da88e17b6725eae1eb6021003f187e187ed2d9df64e8a3c8f4a257b0efe94706): complete subsection reference.

- [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b0a5251ead8f0ceabde02cf071dd72a814033cc432b5c122b67826a1cfa78600): complete subsection reference.

<a id="canonical-9f111a01e3c3310c07ec160e2e2019cf973839302cf0464e90a39fff9b3a4645"></a>

## All schema paths — Property reference / 52d63fd75df3 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--rate_limiter_policy--reference--group-001.md#canonical-dc1af2b4da93859c7e3cfa0b006f058138fd14608691f65c84ddda61fb5e3060) |
| `any_server` | [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-81e6642340d78b5211aeedcb709ea41bdecca75e748d1e2c776f926fdd0ee283) |
| `description` | [description](data-sources--rate_limiter_policy--reference--group-001.md#canonical-99b2a41a14815005605d1301899ab177104f3063ba3245bb9f16ea3fcd3f2453) |
| `id` | [id](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2ba1d430d0ef94bcff9f428ba3d71f1a3274f326340a4ec592ea3b062fe40631) |
| `labels` | [labels](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e5be53631499f7d3dc6c55768d1b6f65d26a7cb986b2c345588e56ab6f79694d) |
| `name` | [name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0b84b357802e1d431d6db41795dc2f9f246127aa0c0392c9f928338b9dfa976c) |
| `namespace` | [namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-85e24df12c278fe56d90885e689c268e53c4ca360c2146a72b865dff332df80e) |
| `rules` | [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1aca56cb26db14beb2595b15c8132579d848edbd98140accfaf490e2b012ef3a) |
| `rules.metadata` | [rules.metadata](data-sources--rate_limiter_policy--reference--group-001.md#canonical-11af1eb14797628a96d654d17ac9405e14ee01490debeb8044f8c43ce29939e2) |
| `rules.metadata.description_spec` | [rules.metadata.description_spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-82e5677ebc8f0b77690880399631e68e69bd17e9bd1c41a0b19c0467e86f97f6) |
| `rules.metadata.name` | [rules.metadata.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e01aeb883340b8dd48416d5c8edad44d96cb153d23062eb6346320961907998f) |
| `rules.spec` | [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-07820e7b9cbf5fb23dd7f944dbcd25de0ad4a0e5d9173d53564bad44f9092b89) |
| `rules.spec.any_asn` | [rules.spec.any_asn](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a57e0f0c3c1099e77d9506466f838f1495f836019f02b0dc70780d6d88f513d6) |
| `rules.spec.any_country` | [rules.spec.any_country](data-sources--rate_limiter_policy--reference--group-001.md#canonical-65361bcc9931561134cba09bafbdfeaee42245af48f1fcae97cfd630d427af89) |
| `rules.spec.any_ip` | [rules.spec.any_ip](data-sources--rate_limiter_policy--reference--group-001.md#canonical-354ae2b48480fbd6abdcf5532b2b61942783468f8daf17923750145c7a4f83de) |
| `rules.spec.apply_rate_limiter` | [rules.spec.apply_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b8b26e531e3a23507bc42543aeafea61da15a30b783f1565cb94828e39908f39) |
| `rules.spec.asn_list` | [rules.spec.asn_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-db8cb47c9157f12683c89b123908b9e9ec1f32db8c51d31999960e6a46a1d46b) |
| `rules.spec.asn_list.as_numbers` | [rules.spec.asn_list.as_numbers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-bc221bd2c237a33486f3eff76d66d93e215e42e8a8fd7d8b24872157861283a9) |
| `rules.spec.asn_matcher` | [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-f5dca5422109da3fb45626e795d9afaa8e40f74b4b17770150154078d81b4b54) |
| `rules.spec.asn_matcher.asn_sets` | [rules.spec.asn_matcher.asn_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-32d52801a0b95e315166b5ae1fcacac9947921887bb90e998cfc9a3efdfecf5d) |
| `rules.spec.asn_matcher.asn_sets.kind` | [rules.spec.asn_matcher.asn_sets.kind](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2bb1cf041e33d9bed9aa576c05074e4cbda090d567f40e002c045a5995c46361) |
| `rules.spec.asn_matcher.asn_sets.name` | [rules.spec.asn_matcher.asn_sets.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b09324f454f630d30ae206425b9646be77156d28a551979809ba7b3286f52d32) |
| `rules.spec.asn_matcher.asn_sets.namespace` | [rules.spec.asn_matcher.asn_sets.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b25fc703c1ae9b018000410dec574ef5cd1a7ecbe5acb1b6073e0fd358df547f) |
| `rules.spec.asn_matcher.asn_sets.tenant` | [rules.spec.asn_matcher.asn_sets.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-303cb32176bc092828f87d628ade51b002f87cce95ad713ad97ab29fb49a4906) |
| `rules.spec.asn_matcher.asn_sets.uid` | [rules.spec.asn_matcher.asn_sets.uid](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c815fc0cd27ea558af5370aae5e8429faccb73153175e564ca19ac18ff0b659b) |
| `rules.spec.bypass_rate_limiter` | [rules.spec.bypass_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-33d0507d71419ca21a42d68aa4f37214e83bffa5c02696554db2158e2cc85f25) |
| `rules.spec.country_list` | [rules.spec.country_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-cbf51acb90e0540960de12ed4c27edfcc0d8464c9d7585c5585bddbe30cb3bdc) |
| `rules.spec.country_list.country_codes` | [rules.spec.country_list.country_codes](data-sources--rate_limiter_policy--reference--group-001.md#canonical-f9ec88ba1c77547eb9134d515fe14996a48ffe3d6a56543fcb176de7b23b4684) |
| `rules.spec.country_list.invert_match` | [rules.spec.country_list.invert_match](data-sources--rate_limiter_policy--reference--group-001.md#canonical-fb25caf29dbc1b2fa7b587048dc656e66e2ea3267286761a10522a91297c7b65) |
| `rules.spec.custom_rate_limiter` | [rules.spec.custom_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-17bbfa538b81224702b234dba88e6063a449103e33501809982dcb60f1660399) |
| `rules.spec.custom_rate_limiter.name` | [rules.spec.custom_rate_limiter.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-633217d825d629e85d4161f20122950eeca4203da9b2396cec98d0ab74fc4d19) |
| `rules.spec.custom_rate_limiter.namespace` | [rules.spec.custom_rate_limiter.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-659e130839fa8e0d2a6117183e0586bd3e3a96a21f878e948960e2082a95dfab) |
| `rules.spec.custom_rate_limiter.tenant` | [rules.spec.custom_rate_limiter.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-d440d04b94adefb54e715f06bf97830c9dc04680133dcb2fe92d2651c2f54d6c) |
| `rules.spec.domain_matcher` | [rules.spec.domain_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-382b4b7b01ae3fcf83465164e437d8b0a1967d99540920e589925afc489df82f) |
| `rules.spec.domain_matcher.exact_values` | [rules.spec.domain_matcher.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c5ef4a1284e7e8c7a95d5a0d1fa665c201e6e2f4d4fad0c93e7f066c00fb7f40) |
| `rules.spec.domain_matcher.regex_values` | [rules.spec.domain_matcher.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-62c23e6e23c983ff54d1c11cf0750a6481892b353a925134e1e7e2b449fbcb11) |
| `rules.spec.headers` | [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-54845f984d490cd84363e316efce6e244d98e1ad637ae61eee21ef5dd27ddf31) |
| `rules.spec.headers.check_not_present` | [rules.spec.headers.check_not_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a33ea02a6eb0979dccec8380812565bd215ef611d67b8b685b897253a6c9535e) |
| `rules.spec.headers.check_present` | [rules.spec.headers.check_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-85dfec420a7b92087d381f85c5de156413a5eb315ad4519e2414dceb68540021) |
| `rules.spec.headers.invert_matcher` | [rules.spec.headers.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a685f060844b8064dba73647dabec325fbed36688e0741e27644d321f04d0da9) |
| `rules.spec.headers.item` | [rules.spec.headers.item](data-sources--rate_limiter_policy--reference--group-001.md#canonical-d94d73796e96e6d28feafe035b0d031e04af85f2ebeedc1e42e4d7b4af3b8189) |
| `rules.spec.headers.item.exact_values` | [rules.spec.headers.item.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-5ee29a611f7dec0fedf9154004551b852a42644e706847f424fcd3f673fcacc7) |
| `rules.spec.headers.item.regex_values` | [rules.spec.headers.item.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-73b3803cfc77851425f63cdfce2d9c25527e9e52c87d86bace47f8637e0ee312) |
| `rules.spec.headers.item.transformers` | [rules.spec.headers.item.transformers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-36f966366b110a7528d97ccf519999f5d18fadb251ea7730aeff698bca905cc5) |
| `rules.spec.headers.name` | [rules.spec.headers.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-d4215e5fbe4490ea0865188ac9645bfc0065907402dc1c334bc9b5c08804cc26) |
| `rules.spec.http_method` | [rules.spec.http_method](data-sources--rate_limiter_policy--reference--group-001.md#canonical-200c86f8b198e44b7080df0b7e8097d1a452ef13bdf7351818f57f00dd0a2f2a) |
| `rules.spec.http_method.invert_matcher` | [rules.spec.http_method.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-171aadfccd16326f71cf7161ca27f9a350df2a5740c87e11601c3904a537133d) |
| `rules.spec.http_method.methods` | [rules.spec.http_method.methods](data-sources--rate_limiter_policy--reference--group-001.md#canonical-de1869a81a01a8a42efe6f3b16c6d06a82e44ad00afe502b80e599827ddf7a43) |
| `rules.spec.ip_matcher` | [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-96220d91a70d418bd5d422ded9c178b506568363074a3c8db7e320eb8f76ad8e) |
| `rules.spec.ip_matcher.invert_matcher` | [rules.spec.ip_matcher.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-8336594f0b8fed01cf8b001b66fe2b4e54a3229c6c479ecfd481802df3b16f58) |
| `rules.spec.ip_matcher.prefix_sets` | [rules.spec.ip_matcher.prefix_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3c7efd96db222315d1a29049b057693cca38884c4b1c00df283e2ad5176757a0) |
| `rules.spec.ip_matcher.prefix_sets.kind` | [rules.spec.ip_matcher.prefix_sets.kind](data-sources--rate_limiter_policy--reference--group-001.md#canonical-6aa0174cac1a5a0432139dd68e1bc24f69729c0d931b16458c2016abaff91c53) |
| `rules.spec.ip_matcher.prefix_sets.name` | [rules.spec.ip_matcher.prefix_sets.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-991c32eb1344bd876777438e9b4feb3479ba80af3f9afb8a03fdb8e86750c34f) |
| `rules.spec.ip_matcher.prefix_sets.namespace` | [rules.spec.ip_matcher.prefix_sets.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-cdda4fa821bd44462660886436fa988ad055292f5f6e7ade45b75ea18dd36a9a) |
| `rules.spec.ip_matcher.prefix_sets.tenant` | [rules.spec.ip_matcher.prefix_sets.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-ac81757938449049b8bb28f0c851c44f910edadc39a15b2d252767366075ed3b) |
| `rules.spec.ip_matcher.prefix_sets.uid` | [rules.spec.ip_matcher.prefix_sets.uid](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1b5b8ed6c45ba4439ad372e26be5ba7def045ecffe7ccdab5a52b8497eb4d300) |
| `rules.spec.ip_prefix_list` | [rules.spec.ip_prefix_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a7d49cb7f0d74f1a2698fb57c5dd9b357c15c8942255a9ed49747e037e77a860) |
| `rules.spec.ip_prefix_list.invert_match` | [rules.spec.ip_prefix_list.invert_match](data-sources--rate_limiter_policy--reference--group-001.md#canonical-09f66b983a2118cfef92579847666a5445f0992321e8bf47874b4785367b551e) |
| `rules.spec.ip_prefix_list.ip_prefixes` | [rules.spec.ip_prefix_list.ip_prefixes](data-sources--rate_limiter_policy--reference--group-001.md#canonical-d62e84bc6ee44b40a9b80ac6db2ad9b60d90f3cae2c6db7fee61f63525c66dc5) |
| `rules.spec.path` | [rules.spec.path](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b51be94676bfcaed2e45a5e684593d7505bb5a2262f507b62d17b48ee93a9f74) |
| `rules.spec.path.encoded_path_matcher` | [rules.spec.path.encoded_path_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9faec31d05d1787c93789650174ac350231402d80369a59df556252a8db6122b) |
| `rules.spec.path.exact_values` | [rules.spec.path.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-ca26e11be0180f21fc9c084ccbfdbd234e2857de198b9d7dd4f8b0d62a84c8a0) |
| `rules.spec.path.invert_matcher` | [rules.spec.path.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c0e617f4380e4848f97efd5c065938fd89b502b30e1e32ef3bce1d56d6f30835) |
| `rules.spec.path.prefix_values` | [rules.spec.path.prefix_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b7b7bad05e1aabfc8eb0ff22b88ec127157e2f5e821f20321a0d86347da479ee) |
| `rules.spec.path.regex_values` | [rules.spec.path.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1123fbe260af568060000a6211cd3d773068795642ac0c2c644ab42a92287add) |
| `rules.spec.path.suffix_values` | [rules.spec.path.suffix_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1ae2712b95c2bb80a0f1b1edbb26c52fd190a1b164bff457c85ce03c556d5e8f) |
| `rules.spec.path.transformers` | [rules.spec.path.transformers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-63059170d8f594b81f3df7af24cd01afd7a6eb6e6c3c48bffc453e28d2e5a5ee) |
| `rules.spec.segment_policy` | [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c531ef2668ace0cd52b668e7a9c1cf03213df9ff0bcf3f53185c64631f378762) |
| `rules.spec.segment_policy.dst_any` | [rules.spec.segment_policy.dst_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-99deccf19023e230909fd0cf84903549b2c3f5170bc4cb8f127d7ea88722152c) |
| `rules.spec.segment_policy.dst_segments` | [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4da3129c8aefea988d8d10313c1ab2a3ec9465b323bfb2bb16c08d4a39921f6e) |
| `rules.spec.segment_policy.dst_segments.segments` | [rules.spec.segment_policy.dst_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0e1172bffddb97f5f37a7803cc382f28f8670193132cff2966052e7fa176fe30) |
| `rules.spec.segment_policy.dst_segments.segments.name` | [rules.spec.segment_policy.dst_segments.segments.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-79b4813639de5db0d123feb3d76fb01fd0bc8b201c0a763eba4cf38dda1d44af) |
| `rules.spec.segment_policy.dst_segments.segments.namespace` | [rules.spec.segment_policy.dst_segments.segments.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1f052427da8ba8219a9e3b0347ef7842a8f8bd4585b270aa71abb377995915bc) |
| `rules.spec.segment_policy.dst_segments.segments.tenant` | [rules.spec.segment_policy.dst_segments.segments.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-dd5548e9c29c3d669756ba09cc3d75609ae03503c7a60c0f4e471eec7d10842c) |
| `rules.spec.segment_policy.intra_segment` | [rules.spec.segment_policy.intra_segment](data-sources--rate_limiter_policy--reference--group-001.md#canonical-91a05fc189a0e1bf3212a0745cee430a054cf22add595a6a6bc6255cd5f98e72) |
| `rules.spec.segment_policy.src_any` | [rules.spec.segment_policy.src_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9f4e9472842a3e3b42410d57c89e635abd816af17761eefcacd3858592269792) |
| `rules.spec.segment_policy.src_segments` | [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-58f718dbbc3a1d7b0a427bd4ee1e15fa2937ba2070bc0c23d3356f5fd59d01d2) |
| `rules.spec.segment_policy.src_segments.segments` | [rules.spec.segment_policy.src_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c3174a8273d72dd558dbb3971169b0ac4a6e3a32b786f5b6219101ec88ece78e) |
| `rules.spec.segment_policy.src_segments.segments.name` | [rules.spec.segment_policy.src_segments.segments.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0b3704592cd7d9df919b3f20263920b7f571746aa3d75e8997b861d0422550e7) |
| `rules.spec.segment_policy.src_segments.segments.namespace` | [rules.spec.segment_policy.src_segments.segments.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4d9a231809a58eeb1b04675bd0b0431593606216f4e4a7e8c838d42ace6e3681) |
| `rules.spec.segment_policy.src_segments.segments.tenant` | [rules.spec.segment_policy.src_segments.segments.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-d0a4439c9438d62ccf945e5178b18d6e0dcf4d05ab16909b20507f2464cdeacb) |
| `server_name` | [server_name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-5d7f5c770a159307202cf728091de193d8946519a8b57c319f06e8c4371f4a1a) |
| `server_name_matcher` | [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-58d9103af9cdde757afa45aac2fdce10fe308323b11a0c0d1a0f56a0d6391164) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-8f21946468005067c81be35ebc8f7b3a5df5d1ff43222cb041b122744cdcf3a1) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-ef2d14061cc47af239873b7a287845acc4c0b4fcb34ddddc110e36f652cd11fb) |
| `server_selector` | [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-7cf788cf70ecfcc93b52e2c1f6d7e5b53b40c6b638a3875879cfaf100fe4685a) |
| `server_selector.expressions` | [server_selector.expressions](data-sources--rate_limiter_policy--reference--group-001.md#canonical-6bd95e858c9db6b4cd0e5870ce32ca296989813022fc6c2357eb70c7937cdc1f) |

<a id="canonical-58b246b1f4c58295278a86acdfcdcbe9aa2dab64ab2a00ba6b756f9f6fc98736"></a>

## Next pages — Property reference / 52d63fd75df3 / 12

- [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-99007c6de9f1e79410e4b7e6abc47b52b4a496329e04f2ad5d363fdd55c265f9)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-da88e17b6725eae1eb6021003f187e187ed2d9df64e8a3c8f4a257b0efe94706)
- [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b0a5251ead8f0ceabde02cf071dd72a814033cc432b5c122b67826a1cfa78600)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-99007c6de9f1e79410e4b7e6abc47b52b4a496329e04f2ad5d363fdd55c265f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b1f64cbb78bfc3dd283c2006f4c685840e0618dee3b9204ab1fa711493d3604"></a>

## any_server — any_server / 75af26acc0d3 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- any_server

<a id="canonical-81e6642340d78b5211aeedcb709ea41bdecca75e748d1e2c776f926fdd0ee283"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option

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

- [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-81e6642340d78b5211aeedcb709ea41bdecca75e748d1e2c776f926fdd0ee283)
- [server_name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-5d7f5c770a159307202cf728091de193d8946519a8b57c319f06e8c4371f4a1a)
- [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-58d9103af9cdde757afa45aac2fdce10fe308323b11a0c0d1a0f56a0d6391164)
- [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-7cf788cf70ecfcc93b52e2c1f6d7e5b53b40c6b638a3875879cfaf100fe4685a)

Select alternatives according to the provider validators above.

<a id="canonical-40b9b7c1bf56091f9927805a4db6404f2146a624d3a4b9c64ceb671b295a69b9"></a>

## Direct properties — any_server / 75af26acc0d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bcd9080d867c60789f4625c81c125a8ae731eb2d5d871ebbc3d070c83bd328ee"></a>

## Next pages — any_server / 75af26acc0d3 / 4

- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7285f0ad43123dc4b6e764f6bc98aa56cf9be93627bd6e58e001f34b805580be"></a>

## rules — rules / 6c84149a0758 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- rules

<a id="canonical-1aca56cb26db14beb2595b15c8132579d848edbd98140accfaf490e2b012ef3a"></a>

Type: `"list"`. Computed.

List of RateLimiterRules that are evaluated sequentially till a matching rule is identified.
Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-89ac222470de692a65b905cb183bd093c17d28ec95266ee3b6982e0caf416428"></a>

## Direct properties — rules / 6c84149a0758 / 3

- [metadata](data-sources--rate_limiter_policy--reference--group-001.md#canonical-718bb85d437d6c1acb17bb30841310ac7daf573ea9afdec4b18f549e610bed6e): complete subsection reference.

- [spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb): complete subsection reference.

<a id="canonical-b857a3625cf1ca57813bb8b2a9c2f02eeb9d7fc24f675fd66f2e2c99cedacc5b"></a>

## Next pages — rules / 6c84149a0758 / 4

- [rules.metadata](data-sources--rate_limiter_policy--reference--group-001.md#canonical-718bb85d437d6c1acb17bb30841310ac7daf573ea9afdec4b18f549e610bed6e)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-718bb85d437d6c1acb17bb30841310ac7daf573ea9afdec4b18f549e610bed6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44efa90785831f7e11b3bdffd3b2aeecc87006cd46a9728e3baaa47fb01f987d"></a>

## rules.metadata — rules.metadata / b0392d059a46 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- rules.metadata

<a id="canonical-11af1eb14797628a96d654d17ac9405e14ee01490debeb8044f8c43ce29939e2"></a>

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

<a id="canonical-ad434033a755f80b84d3d14834d718705a517cf6230b18b16da518d16c5d7a45"></a>

## Direct properties — rules.metadata / b0392d059a46 / 3

<a id="canonical-82e5677ebc8f0b77690880399631e68e69bd17e9bd1c41a0b19c0467e86f97f6"></a>

<a id="canonical-3b04fe3e99056068095d9c65c53c57034f8a26f391641eee2bc98164da7be445"></a>

## description_spec property — rules.metadata / b0392d059a46 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-e01aeb883340b8dd48416d5c8edad44d96cb153d23062eb6346320961907998f"></a>

<a id="canonical-ebf6c548c202d28ecb88fd3bcef7031654a40af3dcbf59d4a274c210452cdf3f"></a>

## name property — rules.metadata / b0392d059a46 / 5

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

<a id="canonical-723c1e90fc0a2bf3aec57e33ddc25df8faffe7eeb100c4dff2c05d961795dd20"></a>

## Next pages — rules.metadata / b0392d059a46 / 6

- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18abff8827638cf0b54f326c26453fa24a17bfb4da883c6c4bf719d9a994d2b1"></a>

## rules.spec — rules.spec / 3a0dc6e38d2c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- rules.spec

<a id="canonical-07820e7b9cbf5fb23dd7f944dbcd25de0ad4a0e5d9173d53564bad44f9092b89"></a>

Type: `"single"`. Computed.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Upstream description:

Shape of Rate Limiter Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"apply_rate_limiter\",\"bypass_rate_limiter\",\"custom_rate_limiter\"]",
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-country_choice": "[\"any_country\",\"country_list\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

<a id="canonical-67cfe98e0f415df76093d0f7364572dca1e9b4bce7090b3857122ccc395a0882"></a>

## Direct properties — rules.spec / 3a0dc6e38d2c / 3

- [any_asn](data-sources--rate_limiter_policy--reference--group-001.md#canonical-733347969eeb02595c22ed7386b0ad21c507d7ae2a38ec9b9632e601526798ff): complete subsection reference.

- [any_country](data-sources--rate_limiter_policy--reference--group-001.md#canonical-8e795816266b8e529c93cf2b40fe8fc83aef498f7f08ed303378f5985717db28): complete subsection reference.

- [any_ip](data-sources--rate_limiter_policy--reference--group-001.md#canonical-ccc00d1ec71f88c5dd3487db8f64dbf1ea30de267491f050736fc702fa6fdb75): complete subsection reference.

- [apply_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1a7da9ecc8b998bfd418085d3be95fd72e7508229863234d5204fd1824debfec): complete subsection reference.

- [asn_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c6807bad165a66ace990bf368ee22b11fb9e7ebf2c6c8448742e331addea9fc8): complete subsection reference.

- [asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c2a59573bd8a37247ef48825c45a55060e2faef3971d4991982dca313a9df380): complete subsection reference.

- [bypass_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9f196193e9f6eef014b1197b63ca9d14dd6e9990a3ac4a269cd5cd19b131f35a): complete subsection reference.

- [country_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e70e33512c9b6a6f3221ea4479e5a339c8476fba0d62662ed02581db79c85fda): complete subsection reference.

- [custom_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-257edeb3357ed9b3970ba83f2e5a127af81958915c34ed88d6204d319e17d2fe): complete subsection reference.

- [domain_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b367e1556e829c82ca38c05c2cad5acaea334fb94490878623dfced6aa17624e): complete subsection reference.

- [headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1): complete subsection reference.

- [http_method](data-sources--rate_limiter_policy--reference--group-001.md#canonical-435585446265ce9d3352f9ab1d78c43a5873f3d35a7fd7bf76f3f55bd3bcaea6): complete subsection reference.

- [ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-960ba7eaffa1f453120be0bca57bd87e5b281334855f5ccd6e8101bd37219185): complete subsection reference.

- [ip_prefix_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-334776643454dd11e8b97ae93d3873abedc016fbc8c4a02b040a350caa3319a2): complete subsection reference.

- [path](data-sources--rate_limiter_policy--reference--group-001.md#canonical-19a6e8088a3282f708a1c21d9e0f3b4bd020c2ac9a4d937d56d229cb70370abe): complete subsection reference.

- [segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede): complete subsection reference.

<a id="canonical-890ca56590f1e96c0cf327694efb9fbe6d63c3fddb33972449931f1eabfedc8d"></a>

## Next pages — rules.spec / 3a0dc6e38d2c / 4

- [rules.spec.any_asn](data-sources--rate_limiter_policy--reference--group-001.md#canonical-733347969eeb02595c22ed7386b0ad21c507d7ae2a38ec9b9632e601526798ff)
- [rules.spec.any_country](data-sources--rate_limiter_policy--reference--group-001.md#canonical-8e795816266b8e529c93cf2b40fe8fc83aef498f7f08ed303378f5985717db28)
- [rules.spec.any_ip](data-sources--rate_limiter_policy--reference--group-001.md#canonical-ccc00d1ec71f88c5dd3487db8f64dbf1ea30de267491f050736fc702fa6fdb75)
- [rules.spec.apply_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1a7da9ecc8b998bfd418085d3be95fd72e7508229863234d5204fd1824debfec)
- [rules.spec.asn_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c6807bad165a66ace990bf368ee22b11fb9e7ebf2c6c8448742e331addea9fc8)
- [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c2a59573bd8a37247ef48825c45a55060e2faef3971d4991982dca313a9df380)
- [rules.spec.bypass_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9f196193e9f6eef014b1197b63ca9d14dd6e9990a3ac4a269cd5cd19b131f35a)
- [rules.spec.country_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e70e33512c9b6a6f3221ea4479e5a339c8476fba0d62662ed02581db79c85fda)
- [rules.spec.custom_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-257edeb3357ed9b3970ba83f2e5a127af81958915c34ed88d6204d319e17d2fe)
- [rules.spec.domain_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-b367e1556e829c82ca38c05c2cad5acaea334fb94490878623dfced6aa17624e)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1)
- [rules.spec.http_method](data-sources--rate_limiter_policy--reference--group-001.md#canonical-435585446265ce9d3352f9ab1d78c43a5873f3d35a7fd7bf76f3f55bd3bcaea6)
- [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-960ba7eaffa1f453120be0bca57bd87e5b281334855f5ccd6e8101bd37219185)
- [rules.spec.ip_prefix_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-334776643454dd11e8b97ae93d3873abedc016fbc8c4a02b040a350caa3319a2)
- [rules.spec.path](data-sources--rate_limiter_policy--reference--group-001.md#canonical-19a6e8088a3282f708a1c21d9e0f3b4bd020c2ac9a4d937d56d229cb70370abe)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-733347969eeb02595c22ed7386b0ad21c507d7ae2a38ec9b9632e601526798ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2191903e7a2449f4f458ea66818b9b0c493a941e3fa0df8acf3618a75e6a4e77"></a>

## rules.spec.any_asn — rules.spec.any_asn / aaa3fa1191b8 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.any_asn

<a id="canonical-a57e0f0c3c1099e77d9506466f838f1495f836019f02b0dc70780d6d88f513d6"></a>

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

<a id="canonical-762f5b4f7dde6cd9555f761c8bfb57f23830a96df82c35cb6e36865e5679944a"></a>

## Direct properties — rules.spec.any_asn / aaa3fa1191b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80ef60462997e6839c672ab90678aaf3e8270051ba215d9daafcb886bbb5e994"></a>

## Next pages — rules.spec.any_asn / aaa3fa1191b8 / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-8e795816266b8e529c93cf2b40fe8fc83aef498f7f08ed303378f5985717db28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-920fd16cd796965cdfd8bbf335c6ff31ec1bed48d919e4e2f92c5d5f3c709bc7"></a>

## rules.spec.any_country — rules.spec.any_country / fa18637d9f2c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.any_country

<a id="canonical-65361bcc9931561134cba09bafbdfeaee42245af48f1fcae97cfd630d427af89"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for any country.

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

<a id="canonical-823d3fab19d42625d986bb36101df240b0a7d719032c5d7f0e6836e43495b851"></a>

## Direct properties — rules.spec.any_country / fa18637d9f2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7348665f6fa7402755f287013fba67c3204cde02f27e0fca3159b4c60db4192c"></a>

## Next pages — rules.spec.any_country / fa18637d9f2c / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-ccc00d1ec71f88c5dd3487db8f64dbf1ea30de267491f050736fc702fa6fdb75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c95d46d935212fb87136526bbe795b46576cc3884c6f6aac90d9a60a1f431261"></a>

## rules.spec.any_ip — rules.spec.any_ip / 0edb7832aa2a / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.any_ip

<a id="canonical-354ae2b48480fbd6abdcf5532b2b61942783468f8daf17923750145c7a4f83de"></a>

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

<a id="canonical-7732ca4b4788b1580e42cd359bb0ea0ec20f9a31f453d9a561a95f72a8ec1c56"></a>

## Direct properties — rules.spec.any_ip / 0edb7832aa2a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d17d834819a7105092e43b62ee82cb7b0c90e6c52dc1c3fa8cca61a714322da"></a>

## Next pages — rules.spec.any_ip / 0edb7832aa2a / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-1a7da9ecc8b998bfd418085d3be95fd72e7508229863234d5204fd1824debfec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e24c58076e0adfafcee371cdc2d39341e2dc29326ef776f80947ff6db7b69fbf"></a>

## rules.spec.apply_rate_limiter — rules.spec.apply_rate_limiter / 565c7b1a9df8 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.apply_rate_limiter

<a id="canonical-b8b26e531e3a23507bc42543aeafea61da15a30b783f1565cb94828e39908f39"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for apply rate limiter.

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

<a id="canonical-daa757f917d549a9eb0171f5b1cea4a829a83591a617506087f359c5c11a92ea"></a>

## Direct properties — rules.spec.apply_rate_limiter / 565c7b1a9df8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84ff0249d51cf94fa7d920157f89772eec57b715bfc769269376919c8b07872f"></a>

## Next pages — rules.spec.apply_rate_limiter / 565c7b1a9df8 / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-c6807bad165a66ace990bf368ee22b11fb9e7ebf2c6c8448742e331addea9fc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e878fe843e5d24e9c13dc30ff6d82cb1270acad46d2df734a2cc98b7571f7e3a"></a>

## rules.spec.asn_list — rules.spec.asn_list / 98bcd7257d95 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.asn_list

<a id="canonical-db8cb47c9157f12683c89b123908b9e9ec1f32db8c51d31999960e6a46a1d46b"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-05afa90b1ad77de36109f8757c27c47a675afaa39cb2ce2612896a3954148bf3"></a>

## Direct properties — rules.spec.asn_list / 98bcd7257d95 / 3

<a id="canonical-bc221bd2c237a33486f3eff76d66d93e215e42e8a8fd7d8b24872157861283a9"></a>

<a id="canonical-9c474017cca473044fc6982f4f566245d85baabb7f6b4c8a42c5294f7fc734c6"></a>

## as_numbers property — rules.spec.asn_list / 98bcd7257d95 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-63f88a36f602a135ed5d92da3795aa63f79f5a22d0b47c909ea0e4bdb06eff7b"></a>

## Next pages — rules.spec.asn_list / 98bcd7257d95 / 5

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-c2a59573bd8a37247ef48825c45a55060e2faef3971d4991982dca313a9df380"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a8a33b7b4eade94e760bf91fe833afd5d07e01a3d31b2cfa399880550522f0f"></a>

## rules.spec.asn_matcher — rules.spec.asn_matcher / 829d853c5b87 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.asn_matcher

<a id="canonical-f5dca5422109da3fb45626e795d9afaa8e40f74b4b17770150154078d81b4b54"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-5c2c25f0244a3a9aa9df246b43c0ff9be2ae043ae45924b8ccfafb049ee3be8d"></a>

## Direct properties — rules.spec.asn_matcher / 829d853c5b87 / 3

- [asn_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-720d459d7fd7b958e91ceafa81c8b32e8f9c115bab19b85dcc29f4a78a13db41): complete subsection reference.

<a id="canonical-21bab2a19d6a8b3bbc4d88f5208be464f1a1784c8386af0734f04e9ee7d640ec"></a>

## Next pages — rules.spec.asn_matcher / 829d853c5b87 / 4

- [rules.spec.asn_matcher.asn_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-720d459d7fd7b958e91ceafa81c8b32e8f9c115bab19b85dcc29f4a78a13db41)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-720d459d7fd7b958e91ceafa81c8b32e8f9c115bab19b85dcc29f4a78a13db41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41ebbc81dbbda16ff50a535e33ff299fb36f82545636c2eb0f51ddf6a93a4eb5"></a>

## rules.spec.asn_matcher.asn_sets — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c2a59573bd8a37247ef48825c45a55060e2faef3971d4991982dca313a9df380)
- rules.spec.asn_matcher.asn_sets

<a id="canonical-32d52801a0b95e315166b5ae1fcacac9947921887bb90e998cfc9a3efdfecf5d"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

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

<a id="canonical-9bfe8871c4921e3205764b941e48dd464ea9e0d2bcbfa439abc6a7e59ced633f"></a>

## Direct properties — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 3

<a id="canonical-2bb1cf041e33d9bed9aa576c05074e4cbda090d567f40e002c045a5995c46361"></a>

<a id="canonical-659d7057c7d9f3512532197ce3ba636384035af19e2dcd75bc87ad59d89dc099"></a>

## kind property — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 4

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

<a id="canonical-b09324f454f630d30ae206425b9646be77156d28a551979809ba7b3286f52d32"></a>

<a id="canonical-e6ca85f3952d8e6b3963ebd376d3879b94404378175f1428534316e8663ced34"></a>

## name property — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 5

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

<a id="canonical-b25fc703c1ae9b018000410dec574ef5cd1a7ecbe5acb1b6073e0fd358df547f"></a>

<a id="canonical-54095a1da0d7076e529bbfdacbd1aa0970327990421d78a4303f307fbcaf3779"></a>

## namespace property — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 6

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

<a id="canonical-303cb32176bc092828f87d628ade51b002f87cce95ad713ad97ab29fb49a4906"></a>

<a id="canonical-339ccfba957d44408a745a49c78ac471af312f4bbfce4df022ffce7a328fbe3f"></a>

## tenant property — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 7

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

<a id="canonical-c815fc0cd27ea558af5370aae5e8429faccb73153175e564ca19ac18ff0b659b"></a>

<a id="canonical-f83af4122fd6f496f1bb9c66af6f587936b0088b8b5a78ec15bd8192e6cdb3b4"></a>

## uid property — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 8

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

<a id="canonical-1a8d10758676b31b01f321603acf22f51f00a797dd271e5cf024c0ec015ff99a"></a>

## Next pages — rules.spec.asn_matcher.asn_sets / c8003b9324a9 / 9

- [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c2a59573bd8a37247ef48825c45a55060e2faef3971d4991982dca313a9df380)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-9f196193e9f6eef014b1197b63ca9d14dd6e9990a3ac4a269cd5cd19b131f35a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e51657ecce5feadf6c9153e9724b7e4b16cb938421d2cba5345b83afa6a9324"></a>

## rules.spec.bypass_rate_limiter — rules.spec.bypass_rate_limiter / 7a701e54811f / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.bypass_rate_limiter

<a id="canonical-33d0507d71419ca21a42d68aa4f37214e83bffa5c02696554db2158e2cc85f25"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for bypass rate limiter.

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

<a id="canonical-90415bd518240d4cc0d618d1abad82ab1c51337aedbe0c060a1b5c241c2ca2f6"></a>

## Direct properties — rules.spec.bypass_rate_limiter / 7a701e54811f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d826beb3d685d95d69ecc0c417eb173c91ce78bb986631611b733a6e43efacd5"></a>

## Next pages — rules.spec.bypass_rate_limiter / 7a701e54811f / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-e70e33512c9b6a6f3221ea4479e5a339c8476fba0d62662ed02581db79c85fda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8572c17d899199a000363793070d5d849e4c920d4a4e76d5c4eb897eb06168e5"></a>

## rules.spec.country_list — rules.spec.country_list / 515b57ba30ee / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.country_list

<a id="canonical-cbf51acb90e0540960de12ed4c27edfcc0d8464c9d7585c5585bddbe30cb3bdc"></a>

Type: `"single"`. Computed.

Country Codes List. List of Country Codes to match against.

Upstream description:

List of Country Codes to match against.

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

<a id="canonical-b4d0bc0d5e1e487a2c5a08eaea5d312b91eb9e7c34918c674008e212e4c996bc"></a>

## Direct properties — rules.spec.country_list / 515b57ba30ee / 3

<a id="canonical-f9ec88ba1c77547eb9134d515fe14996a48ffe3d6a56543fcb176de7b23b4684"></a>

<a id="canonical-b2b7d66f216bc6a8d6dabdfc3a24468d6d186da083020e58f4cf0b1907b0efb5"></a>

## country_codes property — rules.spec.country_list / 515b57ba30ee / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

List of Country Codes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fb25caf29dbc1b2fa7b587048dc656e66e2ea3267286761a10522a91297c7b65"></a>

<a id="canonical-9b4d6e7542fd3c94b3c46e2a97ee4033d6f30f00bf5a2199c8832bd1e6eed45f"></a>

## invert_match property — rules.spec.country_list / 515b57ba30ee / 5

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-ffd085558de6c0cf4a2695158031bee56f98a7b49ea1209cd4978bc182c1b0b6"></a>

## Next pages — rules.spec.country_list / 515b57ba30ee / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-257edeb3357ed9b3970ba83f2e5a127af81958915c34ed88d6204d319e17d2fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ca2f09dd17627df5470ffcda4d27ab179cadf8a3e7eacb1a384d24b9c0baa94"></a>

## rules.spec.custom_rate_limiter — rules.spec.custom_rate_limiter / e2a96681539e / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.custom_rate_limiter

<a id="canonical-17bbfa538b81224702b234dba88e6063a449103e33501809982dcb60f1660399"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-e40bfeb1905d4d68986f24c6e6baa7467eb0e246703e17a8243a9fa4b0becbcd"></a>

## Direct properties — rules.spec.custom_rate_limiter / e2a96681539e / 3

<a id="canonical-633217d825d629e85d4161f20122950eeca4203da9b2396cec98d0ab74fc4d19"></a>

<a id="canonical-9fc853743772d659d4b86b4dfe946971d93c3104ff2f8d4d32eed5d0ae3abc48"></a>

## name property — rules.spec.custom_rate_limiter / e2a96681539e / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-659e130839fa8e0d2a6117183e0586bd3e3a96a21f878e948960e2082a95dfab"></a>

<a id="canonical-bef5085166c21552fac8b2ca3fa9a1c64647bad95484982b71130b7788a96154"></a>

## namespace property — rules.spec.custom_rate_limiter / e2a96681539e / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-d440d04b94adefb54e715f06bf97830c9dc04680133dcb2fe92d2651c2f54d6c"></a>

<a id="canonical-bbe3f2385cbf4b108c5813f979a18304e23fd480ed36b13549b82c9342da6fc6"></a>

## tenant property — rules.spec.custom_rate_limiter / e2a96681539e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-09ace41b3197ccb16e875865048cdc9f4eb8062a4b8aa23264de48765513bf5d"></a>

## Next pages — rules.spec.custom_rate_limiter / e2a96681539e / 7

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-b367e1556e829c82ca38c05c2cad5acaea334fb94490878623dfced6aa17624e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f5cf2a950614982c331ce1b1c90e31800dc6217bd3c6735533099a7e1321152"></a>

## rules.spec.domain_matcher — rules.spec.domain_matcher / 89aae1a7ac86 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.domain_matcher

<a id="canonical-382b4b7b01ae3fcf83465164e437d8b0a1967d99540920e589925afc489df82f"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-f6134b7e9747ec18d11b976801d864f635639205280e7b8fd6be32df9aeed254"></a>

## Direct properties — rules.spec.domain_matcher / 89aae1a7ac86 / 3

<a id="canonical-c5ef4a1284e7e8c7a95d5a0d1fa665c201e6e2f4d4fad0c93e7f066c00fb7f40"></a>

<a id="canonical-0e354a57b276af9dea7b034d7259712185b05b9978f8d062762f9aa8f4787c24"></a>

## exact_values property — rules.spec.domain_matcher / 89aae1a7ac86 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-62c23e6e23c983ff54d1c11cf0750a6481892b353a925134e1e7e2b449fbcb11"></a>

<a id="canonical-b2d3bdcdbbe697e04e303c3cdb29fa6f704d8a880d473c0d3acf5c3719bd34bc"></a>

## regex_values property — rules.spec.domain_matcher / 89aae1a7ac86 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f219ffaf49205b7bb58dbad6dae162d1d8ab94d8e79b36c595726a5a6f985f71"></a>

## Next pages — rules.spec.domain_matcher / 89aae1a7ac86 / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7959a52c788149a4d22da1535f2edec15acfef4a44372d19bb432d41577dfbbb"></a>

## rules.spec.headers — rules.spec.headers / cfd96b60358c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.headers

<a id="canonical-54845f984d490cd84363e316efce6e244d98e1ad637ae61eee21ef5dd27ddf31"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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

<a id="canonical-7eebd4aeae48ed66f1183904b3d97993c0543abbadf5d314e7e1ff67d12dcbce"></a>

## Direct properties — rules.spec.headers / cfd96b60358c / 3

- [check_not_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3c4cf9e79cbbc6b47d7771abf3aba76924aa29ab1e835acc23e4ce7763069a03): complete subsection reference.

- [check_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-baad7f1d43deba571688444a3b76f3a925dbee6dc1a34d30743adf10ebc10646): complete subsection reference.

<a id="canonical-a685f060844b8064dba73647dabec325fbed36688e0741e27644d321f04d0da9"></a>

<a id="canonical-57a1ebc2580b2b737d24257a2bc24f5a5128032f6f80545c269cfbb02eef2100"></a>

## invert_matcher property — rules.spec.headers / cfd96b60358c / 4

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](data-sources--rate_limiter_policy--reference--group-001.md#canonical-76a73d06db749f3b7b75bffa3ede7b74d02cb4c3bc3e1705828799efce95cb71): complete subsection reference.

<a id="canonical-d4215e5fbe4490ea0865188ac9645bfc0065907402dc1c334bc9b5c08804cc26"></a>

<a id="canonical-fe378a2b440c18d0c9252ba508f717424fd04cb8046eece1a7d84ecaaf36397a"></a>

## name property — rules.spec.headers / cfd96b60358c / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-fb88ce6bfc10ce54bfe44bc912adde5d0cfa2276b24c42a2eeaa9ffe23a0797d"></a>

## Next pages — rules.spec.headers / cfd96b60358c / 6

- [rules.spec.headers.check_not_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3c4cf9e79cbbc6b47d7771abf3aba76924aa29ab1e835acc23e4ce7763069a03)
- [rules.spec.headers.check_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-baad7f1d43deba571688444a3b76f3a925dbee6dc1a34d30743adf10ebc10646)
- [rules.spec.headers.item](data-sources--rate_limiter_policy--reference--group-001.md#canonical-76a73d06db749f3b7b75bffa3ede7b74d02cb4c3bc3e1705828799efce95cb71)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-3c4cf9e79cbbc6b47d7771abf3aba76924aa29ab1e835acc23e4ce7763069a03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69225cd0b519299800ca475314d57ad1c1947ab85e00e59cb877a7ae8d4393c5"></a>

## rules.spec.headers.check_not_present — rules.spec.headers.check_not_present / d3ba9178f0ff / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1)
- rules.spec.headers.check_not_present

<a id="canonical-a33ea02a6eb0979dccec8380812565bd215ef611d67b8b685b897253a6c9535e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-c9665c916b64346b0c112358833d6bc1a156d2e843959881ce4fc85cb10c2d91"></a>

## Direct properties — rules.spec.headers.check_not_present / d3ba9178f0ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6aa93ab176536118ba58b769b7b917fd38091b9bb706afa1c1313e8ad78195a0"></a>

## Next pages — rules.spec.headers.check_not_present / d3ba9178f0ff / 4

- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-baad7f1d43deba571688444a3b76f3a925dbee6dc1a34d30743adf10ebc10646"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-adbf3ede823bb33ab27df54a11a514fef6e983ef369be5b355339a8037cdce94"></a>

## rules.spec.headers.check_present — rules.spec.headers.check_present / e0d32e503b2c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1)
- rules.spec.headers.check_present

<a id="canonical-85dfec420a7b92087d381f85c5de156413a5eb315ad4519e2414dceb68540021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-c2086c14cccd098bc9eb46ebb11878a1b618a77a558a42b5508f8c36e52284ad"></a>

## Direct properties — rules.spec.headers.check_present / e0d32e503b2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de1db3b4fa8705b54447fd109f155157dc2b589f53b75b9782fa6e7c0d316f51"></a>

## Next pages — rules.spec.headers.check_present / e0d32e503b2c / 4

- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-76a73d06db749f3b7b75bffa3ede7b74d02cb4c3bc3e1705828799efce95cb71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15d029554c4cf31f120c574bbdcb46884279d7ea69ea879283a54a5f182f6285"></a>

## rules.spec.headers.item — rules.spec.headers.item / 5092bac423bf / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1)
- rules.spec.headers.item

<a id="canonical-d94d73796e96e6d28feafe035b0d031e04af85f2ebeedc1e42e4d7b4af3b8189"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-0ce37e5d10918818ef64e690f51c7e4d4b3dcc58ecf7e434c8d0a7fbfa159fb8"></a>

## Direct properties — rules.spec.headers.item / 5092bac423bf / 3

<a id="canonical-5ee29a611f7dec0fedf9154004551b852a42644e706847f424fcd3f673fcacc7"></a>

<a id="canonical-902cb872f40fd5f5ce3c423ea273d11890dd55f1fadb312f966162618aec7a6d"></a>

## exact_values property — rules.spec.headers.item / 5092bac423bf / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-73b3803cfc77851425f63cdfce2d9c25527e9e52c87d86bace47f8637e0ee312"></a>

<a id="canonical-f4136c4a1b1c36da7d98986a6bd45771f18cd1cc8c0ac62806de671c058bf2e5"></a>

## regex_values property — rules.spec.headers.item / 5092bac423bf / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-36f966366b110a7528d97ccf519999f5d18fadb251ea7730aeff698bca905cc5"></a>

<a id="canonical-80d49ac1f6bdb5d54a11c266ffee1e22e53e51bb5d4b5fff6bae9f83df89d99c"></a>

## transformers property — rules.spec.headers.item / 5092bac423bf / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bd3037cf260bffaee09a4b7375ed44d3cd520466b1f86c2e252bf80685f0e232"></a>

## Next pages — rules.spec.headers.item / 5092bac423bf / 7

- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-e1f155b3252ac2374279a9a5abcf4896ab5b13992321df63dc0184a80cfa5ae1)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-435585446265ce9d3352f9ab1d78c43a5873f3d35a7fd7bf76f3f55bd3bcaea6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df1d734c3dca1b86a4f9246ea559ec1ae4e143080ae9fa4174798941272868b2"></a>

## rules.spec.http_method — rules.spec.http_method / 5afe02e172d0 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.http_method

<a id="canonical-200c86f8b198e44b7080df0b7e8097d1a452ef13bdf7351818f57f00dd0a2f2a"></a>

Type: `"single"`. Computed.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-10e46dc3f75cc0c5e77cc146fc2b0bbdeb845fef19f21e84a746ed4b9d92584a"></a>

## Direct properties — rules.spec.http_method / 5afe02e172d0 / 3

<a id="canonical-171aadfccd16326f71cf7161ca27f9a350df2a5740c87e11601c3904a537133d"></a>

<a id="canonical-0dda96ba1e3e94506424356f575dde14965834e35da66fb7be628d9c72fe7f9d"></a>

## invert_matcher property — rules.spec.http_method / 5afe02e172d0 / 4

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-de1869a81a01a8a42efe6f3b16c6d06a82e44ad00afe502b80e599827ddf7a43"></a>

<a id="canonical-cbc33bce66c0656939a17af613137ec5a9eff1aa281dab2c317de7f41f7900bc"></a>

## methods property — rules.spec.http_method / 5afe02e172d0 / 5

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d4d044f7835170b05c6a1f601fa9ed6d90be3f575a05269dfd48a0ff0ab00246"></a>

## Next pages — rules.spec.http_method / 5afe02e172d0 / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-960ba7eaffa1f453120be0bca57bd87e5b281334855f5ccd6e8101bd37219185"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c31b6d9933b9adfd274017d9ff1d462d68ab5e02b91ead8b62ce98a4a700502"></a>

## rules.spec.ip_matcher — rules.spec.ip_matcher / b4902230c3bb / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.ip_matcher

<a id="canonical-96220d91a70d418bd5d422ded9c178b506568363074a3c8db7e320eb8f76ad8e"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-8cafd4409795996190d90d0fbf5a3ece79048e6dc82f78ce317367affe87f828"></a>

## Direct properties — rules.spec.ip_matcher / b4902230c3bb / 3

<a id="canonical-8336594f0b8fed01cf8b001b66fe2b4e54a3229c6c479ecfd481802df3b16f58"></a>

<a id="canonical-3765be4bb182297fefaf36202933d70d2a15d9bc060347dfd4078753d1b6588d"></a>

## invert_matcher property — rules.spec.ip_matcher / b4902230c3bb / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-17e540e2c5d2d67193d4ddd9d4efb2be80e486a619d2d8b05eb51915419ab10a): complete subsection reference.

<a id="canonical-a1a8fc82ff85b96b755943ca8cfa7ee2aa3de99116a7be1ecb8a25dad2883ac5"></a>

## Next pages — rules.spec.ip_matcher / b4902230c3bb / 5

- [rules.spec.ip_matcher.prefix_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-17e540e2c5d2d67193d4ddd9d4efb2be80e486a619d2d8b05eb51915419ab10a)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-17e540e2c5d2d67193d4ddd9d4efb2be80e486a619d2d8b05eb51915419ab10a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bba1395da1d3d6f94316fce236eef8e3d41e69625870b4cc445c562c722c6cdb"></a>

## rules.spec.ip_matcher.prefix_sets — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-960ba7eaffa1f453120be0bca57bd87e5b281334855f5ccd6e8101bd37219185)
- rules.spec.ip_matcher.prefix_sets

<a id="canonical-3c7efd96db222315d1a29049b057693cca38884c4b1c00df283e2ad5176757a0"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-1f25bbbf8e616a988e28499dd1948040de12b3be6e23ea20c9322262f32662d6"></a>

## Direct properties — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 3

<a id="canonical-6aa0174cac1a5a0432139dd68e1bc24f69729c0d931b16458c2016abaff91c53"></a>

<a id="canonical-dc37f44dc87b9b70208fa8db2ded873ca849f9bfecc326544eb3e23c3ed6ce8f"></a>

## kind property — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 4

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

<a id="canonical-991c32eb1344bd876777438e9b4feb3479ba80af3f9afb8a03fdb8e86750c34f"></a>

<a id="canonical-d517bd96183d5ef2935c1ffecf11eb9b742bf389576c7b27d9c1b392427a72e9"></a>

## name property — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 5

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

<a id="canonical-cdda4fa821bd44462660886436fa988ad055292f5f6e7ade45b75ea18dd36a9a"></a>

<a id="canonical-fd4a6c1f420605feda00aaac884b5a302e27a567cc96ed5639c8e88cabc0af48"></a>

## namespace property — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 6

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

<a id="canonical-ac81757938449049b8bb28f0c851c44f910edadc39a15b2d252767366075ed3b"></a>

<a id="canonical-8807bad3cad03c8e64df3e74f4853692e3ef63c9f1842744d61a0a6201ddd769"></a>

## tenant property — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 7

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

<a id="canonical-1b5b8ed6c45ba4439ad372e26be5ba7def045ecffe7ccdab5a52b8497eb4d300"></a>

<a id="canonical-51655eb025eb538eef00cf73227dac17187c3219b70222f5d1109d7a70cf4b14"></a>

## uid property — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 8

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

<a id="canonical-e705bd583fbb8bc0f5f176223d721f7249c47642af618e1c0f7866dd34c69c66"></a>

## Next pages — rules.spec.ip_matcher.prefix_sets / ccf837a933b5 / 9

- [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-960ba7eaffa1f453120be0bca57bd87e5b281334855f5ccd6e8101bd37219185)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-334776643454dd11e8b97ae93d3873abedc016fbc8c4a02b040a350caa3319a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68c13af32250441379234c06102060b3e680f9a1fa215fb3b51b5761f9da8853"></a>

## rules.spec.ip_prefix_list — rules.spec.ip_prefix_list / 2db182281f2b / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.ip_prefix_list

<a id="canonical-a7d49cb7f0d74f1a2698fb57c5dd9b357c15c8942255a9ed49747e037e77a860"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-e82ab5d57851c24953181c79a314db069039c8fccde06a3a985ff5b446700813"></a>

## Direct properties — rules.spec.ip_prefix_list / 2db182281f2b / 3

<a id="canonical-09f66b983a2118cfef92579847666a5445f0992321e8bf47874b4785367b551e"></a>

<a id="canonical-ac604e8957a90355af0ae69ce287597ac3130589fa8c5b077ceb933b14a12d2e"></a>

## invert_match property — rules.spec.ip_prefix_list / 2db182281f2b / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-d62e84bc6ee44b40a9b80ac6db2ad9b60d90f3cae2c6db7fee61f63525c66dc5"></a>

<a id="canonical-a8d11b3879efc1663968fbc0679d3d10c8dae667d8780a23233bfd94f3afe3dd"></a>

## ip_prefixes property — rules.spec.ip_prefix_list / 2db182281f2b / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5f2f45620be959ff2843a7d6023deb7d64b7cca9c7b2f6e2bcf9dbd8d162e0f6"></a>

## Next pages — rules.spec.ip_prefix_list / 2db182281f2b / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-19a6e8088a3282f708a1c21d9e0f3b4bd020c2ac9a4d937d56d229cb70370abe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b75a7cfb2037f0c9cd678fe2df07412b9d2b8956dfd01d149e6f33069be25261"></a>

## rules.spec.path — rules.spec.path / 5837780b85ee / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.path

<a id="canonical-b51be94676bfcaed2e45a5e684593d7505bb5a2262f507b62d17b48ee93a9f74"></a>

Type: `"single"`. Computed.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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

<a id="canonical-0dca0ce2fd4f5f856f0f10af741787d7391918d6aafd64a0f878c05d4161222c"></a>

## Direct properties — rules.spec.path / 5837780b85ee / 3

<a id="canonical-9faec31d05d1787c93789650174ac350231402d80369a59df556252a8db6122b"></a>

<a id="canonical-3c5e3ced11a28c07762b4fc76f63675e7678075d75f31393039aaa44055cbc03"></a>

## encoded_path_matcher property — rules.spec.path / 5837780b85ee / 4

Type: `"bool"`. Computed.

Match against the encoded, escaped path.

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

<a id="canonical-ca26e11be0180f21fc9c084ccbfdbd234e2857de198b9d7dd4f8b0d62a84c8a0"></a>

<a id="canonical-4fc7f487db657d6f4f8462973f4ec25c2f0a62174c90b2ac25f2f80756f1a2f7"></a>

## exact_values property — rules.spec.path / 5837780b85ee / 5

Type: `["list", "string"]`. Computed.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c0e617f4380e4848f97efd5c065938fd89b502b30e1e32ef3bce1d56d6f30835"></a>

<a id="canonical-4ec476ad5d7b15c0b38fd9c19fdb5a9f2623ee78863d22169ce7b518e82fffdb"></a>

## invert_matcher property — rules.spec.path / 5837780b85ee / 6

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-b7b7bad05e1aabfc8eb0ff22b88ec127157e2f5e821f20321a0d86347da479ee"></a>

<a id="canonical-5b08495224301ca9952315085a39b247ee382137fd3b9437cd0564bdac91042d"></a>

## prefix_values property — rules.spec.path / 5837780b85ee / 7

Type: `["list", "string"]`. Computed.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1123fbe260af568060000a6211cd3d773068795642ac0c2c644ab42a92287add"></a>

<a id="canonical-c99dbb5f024d3d0ec6f8f2b9022bfd70664b7c78afeebd65cae08d90463cd974"></a>

## regex_values property — rules.spec.path / 5837780b85ee / 8

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1ae2712b95c2bb80a0f1b1edbb26c52fd190a1b164bff457c85ce03c556d5e8f"></a>

<a id="canonical-942819146b6b33895c2f41c6edc60de8a63c8b1bfa791ec1b2dd1476da3e06dc"></a>

## suffix_values property — rules.spec.path / 5837780b85ee / 9

Type: `["list", "string"]`. Computed.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-63059170d8f594b81f3df7af24cd01afd7a6eb6e6c3c48bffc453e28d2e5a5ee"></a>

<a id="canonical-8f6da2aab421c4d5f0f92304ce981d6a39edec72076c4fc862d9b00085ea2171"></a>

## transformers property — rules.spec.path / 5837780b85ee / 10

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4b84ea2a84072037a99993546f797d10ed07a4f2300a381a905f44f4255ad48c"></a>

## Next pages — rules.spec.path / 5837780b85ee / 11

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49f087ece0baea95820070a7b68e47d823c243212012a10ba013e0949d466dd7"></a>

## rules.spec.segment_policy — rules.spec.segment_policy / 06b5fd35ba87 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- rules.spec.segment_policy

<a id="canonical-c531ef2668ace0cd52b668e7a9c1cf03213df9ff0bcf3f53185c64631f378762"></a>

Type: `"single"`. Computed.

Configure source and destination segment for policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

<a id="canonical-0d6d25335468107a8698210eea68aad4d9e89467de32913a39b178007ff65948"></a>

## Direct properties — rules.spec.segment_policy / 06b5fd35ba87 / 3

- [dst_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-348af57686df5939eef1bf162d95086fc931ec121b336329d66f3a01b0d6a744): complete subsection reference.

- [dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c19a231b853b313ec558ea2feadd59f72b6858dc75b518b5a018010d77e04269): complete subsection reference.

- [intra_segment](data-sources--rate_limiter_policy--reference--group-001.md#canonical-60eb90bf452474ee5921a22acf9ed96fa55d5c6d00e6144d920ebb007d7ee591): complete subsection reference.

- [src_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-cadf9b92dde0135c932452efd5cd45747cb50f320de33bf0b905fc93364effb9): complete subsection reference.

- [src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-88cb979867c95e6c6b5a21c185975c30853a002de6f4cc6530aac81e7815cc23): complete subsection reference.

<a id="canonical-084f026f279d00ac4f7a35e34189a8540380247f561323e760065fde503aa9b7"></a>

## Next pages — rules.spec.segment_policy / 06b5fd35ba87 / 4

- [rules.spec.segment_policy.dst_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-348af57686df5939eef1bf162d95086fc931ec121b336329d66f3a01b0d6a744)
- [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c19a231b853b313ec558ea2feadd59f72b6858dc75b518b5a018010d77e04269)
- [rules.spec.segment_policy.intra_segment](data-sources--rate_limiter_policy--reference--group-001.md#canonical-60eb90bf452474ee5921a22acf9ed96fa55d5c6d00e6144d920ebb007d7ee591)
- [rules.spec.segment_policy.src_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-cadf9b92dde0135c932452efd5cd45747cb50f320de33bf0b905fc93364effb9)
- [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-88cb979867c95e6c6b5a21c185975c30853a002de6f4cc6530aac81e7815cc23)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-348af57686df5939eef1bf162d95086fc931ec121b336329d66f3a01b0d6a744"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-554add2aeb58e635aa8af19a09bfa3b66879122c8d0a96595405c89be7011da3"></a>

## rules.spec.segment_policy.dst_any — rules.spec.segment_policy.dst_any / 9822cf5e9c31 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- rules.spec.segment_policy.dst_any

<a id="canonical-99deccf19023e230909fd0cf84903549b2c3f5170bc4cb8f127d7ea88722152c"></a>

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

<a id="canonical-0dfcbe666550d3f6c6731ab9dd4b210cfa8be543359b299370fd508c3b844a26"></a>

## Direct properties — rules.spec.segment_policy.dst_any / 9822cf5e9c31 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06100a54eb271b7487600deafca80b2f8a73000d390666e03354928bfef923f2"></a>

## Next pages — rules.spec.segment_policy.dst_any / 9822cf5e9c31 / 4

- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-c19a231b853b313ec558ea2feadd59f72b6858dc75b518b5a018010d77e04269"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2739b69113518d53ec1c9d436b14095bd96bde6f0a017e527a373b62b1e4084c"></a>

## rules.spec.segment_policy.dst_segments — rules.spec.segment_policy.dst_segments / 717d43879559 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- rules.spec.segment_policy.dst_segments

<a id="canonical-4da3129c8aefea988d8d10313c1ab2a3ec9465b323bfb2bb16c08d4a39921f6e"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

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

<a id="canonical-ce5370cd7d6ddb499556c7167c3f37d76ebbaa3dd0c1755439e0473b2bcca65f"></a>

## Direct properties — rules.spec.segment_policy.dst_segments / 717d43879559 / 3

- [segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-647ad0df9eeb976cfe55a80b2729a996aefdfaf7046868129e0281bf7d55e304): complete subsection reference.

<a id="canonical-13278bb4db354ce6feda6f7eff5acd0aa6e7b2e989f77e690a0feeee38a9e895"></a>

## Next pages — rules.spec.segment_policy.dst_segments / 717d43879559 / 4

- [rules.spec.segment_policy.dst_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-647ad0df9eeb976cfe55a80b2729a996aefdfaf7046868129e0281bf7d55e304)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-647ad0df9eeb976cfe55a80b2729a996aefdfaf7046868129e0281bf7d55e304"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dbb0cf5b42b1ffbd2c5857862fd3dd00b847000fc92c8294a2d724e783f9bbf"></a>

## rules.spec.segment_policy.dst_segments.segments — rules.spec.segment_policy.dst_segments.segments / 99f4039424a4 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c19a231b853b313ec558ea2feadd59f72b6858dc75b518b5a018010d77e04269)
- rules.spec.segment_policy.dst_segments.segments

<a id="canonical-0e1172bffddb97f5f37a7803cc382f28f8670193132cff2966052e7fa176fe30"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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

<a id="canonical-c147e214c95b39ed775df45b4ddb8ab4fe1792cf4847264626a84b3f3d53e5a6"></a>

## Direct properties — rules.spec.segment_policy.dst_segments.segments / 99f4039424a4 / 3

<a id="canonical-79b4813639de5db0d123feb3d76fb01fd0bc8b201c0a763eba4cf38dda1d44af"></a>

<a id="canonical-e14e225ecdfe9d35dd12386f84a0bc3ec2c66467ea1967480385eb1de5f8ab84"></a>

## name property — rules.spec.segment_policy.dst_segments.segments / 99f4039424a4 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1f052427da8ba8219a9e3b0347ef7842a8f8bd4585b270aa71abb377995915bc"></a>

<a id="canonical-cbbd3129d44a5d66ece76ff793a17ec5ba6d07dc9a329f8a7a8da5044939366f"></a>

## namespace property — rules.spec.segment_policy.dst_segments.segments / 99f4039424a4 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-dd5548e9c29c3d669756ba09cc3d75609ae03503c7a60c0f4e471eec7d10842c"></a>

<a id="canonical-c5dcdd66d4dae000666339709aa80e849545dde987c4d12fde5eed9a3885856b"></a>

## tenant property — rules.spec.segment_policy.dst_segments.segments / 99f4039424a4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1a0264c77731ac4cb47403b3d3f46cef4e69f403f65875cc52b1eea6abac8313"></a>

## Next pages — rules.spec.segment_policy.dst_segments.segments / 99f4039424a4 / 7

- [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-c19a231b853b313ec558ea2feadd59f72b6858dc75b518b5a018010d77e04269)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-60eb90bf452474ee5921a22acf9ed96fa55d5c6d00e6144d920ebb007d7ee591"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68bbb1ab1d1c7d431dd0d51770fb65609fdc1c7a9df8621a42f89916806eee4d"></a>

## rules.spec.segment_policy.intra_segment — rules.spec.segment_policy.intra_segment / 926185958027 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- rules.spec.segment_policy.intra_segment

<a id="canonical-91a05fc189a0e1bf3212a0745cee430a054cf22add595a6a6bc6255cd5f98e72"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for intra segment.

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

<a id="canonical-3e192b77c1cbfb8cafd2248afeec4da70242a0945146ef12c56698f121078f5e"></a>

## Direct properties — rules.spec.segment_policy.intra_segment / 926185958027 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6bbfdc779bff84a786e818eeb0ce23f769317f5a9dd12a887c4af68df9271b55"></a>

## Next pages — rules.spec.segment_policy.intra_segment / 926185958027 / 4

- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-cadf9b92dde0135c932452efd5cd45747cb50f320de33bf0b905fc93364effb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ec7e7170e0642fc7aba94347b4c65c6fe7bdfac920dcadd6e4490f44d95224f"></a>

## rules.spec.segment_policy.src_any — rules.spec.segment_policy.src_any / 47dabbfbce25 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- rules.spec.segment_policy.src_any

<a id="canonical-9f4e9472842a3e3b42410d57c89e635abd816af17761eefcacd3858592269792"></a>

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

<a id="canonical-9a7cd4b60aa3fdafa56dc5b7abd204c62679bff4a90b4747f44aa775d96c6921"></a>

## Direct properties — rules.spec.segment_policy.src_any / 47dabbfbce25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64cccce4e33ec42548822630d9750f65a6a8a231e094c396632a89360b303bd7"></a>

## Next pages — rules.spec.segment_policy.src_any / 47dabbfbce25 / 4

- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-88cb979867c95e6c6b5a21c185975c30853a002de6f4cc6530aac81e7815cc23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54737c3fc283d551f7ec808546db157052b775644f72e1870a731ad2a8ef50bf"></a>

## rules.spec.segment_policy.src_segments — rules.spec.segment_policy.src_segments / f0daf36d892c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- rules.spec.segment_policy.src_segments

<a id="canonical-58f718dbbc3a1d7b0a427bd4ee1e15fa2937ba2070bc0c23d3356f5fd59d01d2"></a>

Type: `"single"`. Computed.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

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

<a id="canonical-121afc14c4e0edd0bb797a97a644853406eb8f6a62a5148e915ca2ad2b88a77b"></a>

## Direct properties — rules.spec.segment_policy.src_segments / f0daf36d892c / 3

- [segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-6454be239eeb241cde28f378abbc7712694471f99024c67d5b22675b61c86894): complete subsection reference.

<a id="canonical-03d0dbfada3ced51dedabb29d246a2cda4efff4ef535746e4d07036f7e77104a"></a>

## Next pages — rules.spec.segment_policy.src_segments / f0daf36d892c / 4

- [rules.spec.segment_policy.src_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-6454be239eeb241cde28f378abbc7712694471f99024c67d5b22675b61c86894)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-6454be239eeb241cde28f378abbc7712694471f99024c67d5b22675b61c86894"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c9ae664ca54798f0f112f1cffa9d253078534d88cf91d3e2b84f00579475468"></a>

## rules.spec.segment_policy.src_segments.segments — rules.spec.segment_policy.src_segments.segments / 48736c71095e / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2c97bbeed0fe94a846cf4a1e6a56604bc5f966894d7c709af0d8ca87eb93b745)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-4560be4b285997876dcd130279929cb1dae8e29f58d7e78f3de35e5d490d8bbb)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-a10c0b1ad76f622b1606ae09fc21664e8b620fdc5eaba8ac6921885f0646bede)
- [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-88cb979867c95e6c6b5a21c185975c30853a002de6f4cc6530aac81e7815cc23)
- rules.spec.segment_policy.src_segments.segments

<a id="canonical-c3174a8273d72dd558dbb3971169b0ac4a6e3a32b786f5b6219101ec88ece78e"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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

<a id="canonical-5a81a5a6faad39feee49deba98b411b2088402e7580e67b8592984583c9c7de1"></a>

## Direct properties — rules.spec.segment_policy.src_segments.segments / 48736c71095e / 3

<a id="canonical-0b3704592cd7d9df919b3f20263920b7f571746aa3d75e8997b861d0422550e7"></a>

<a id="canonical-930e0a3cb3d78c697ae4a5c75e2342eec22276ce156c6beefadae2027d628f4e"></a>

## name property — rules.spec.segment_policy.src_segments.segments / 48736c71095e / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-4d9a231809a58eeb1b04675bd0b0431593606216f4e4a7e8c838d42ace6e3681"></a>

<a id="canonical-f7e0a2e69bcb75aa80de40d88247d4d0e84579aa07bec5b410cf989416128aee"></a>

## namespace property — rules.spec.segment_policy.src_segments.segments / 48736c71095e / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-d0a4439c9438d62ccf945e5178b18d6e0dcf4d05ab16909b20507f2464cdeacb"></a>

<a id="canonical-ddd1e97fbe0611c0e2fdf9c0b06db41e4f82923368fe97d42ba646272558b894"></a>

## tenant property — rules.spec.segment_policy.src_segments.segments / 48736c71095e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1c8252e310ca73dcd63fdd30a55e38a948d31628a92e8ed9ad2630a7c03f93a1"></a>

## Next pages — rules.spec.segment_policy.src_segments.segments / 48736c71095e / 7

- [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-88cb979867c95e6c6b5a21c185975c30853a002de6f4cc6530aac81e7815cc23)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-da88e17b6725eae1eb6021003f187e187ed2d9df64e8a3c8f4a257b0efe94706"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bac4cb477a264e438f987d228d5b065be02954536edae01f268d46ac9660926"></a>

## server_name_matcher — server_name_matcher / 7e9bb0f38c0f / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- server_name_matcher

<a id="canonical-58d9103af9cdde757afa45aac2fdce10fe308323b11a0c0d1a0f56a0d6391164"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-4239ee0797968fd20d9e2ad72b7988b6bc0e82621ad830c1218377f7d1468761"></a>

## Direct properties — server_name_matcher / 7e9bb0f38c0f / 3

<a id="canonical-8f21946468005067c81be35ebc8f7b3a5df5d1ff43222cb041b122744cdcf3a1"></a>

<a id="canonical-bb6f197a9b8db3b77d2912670abd73aa257afd70bfda4babe6f75019fdabaad0"></a>

## exact_values property — server_name_matcher / 7e9bb0f38c0f / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ef2d14061cc47af239873b7a287845acc4c0b4fcb34ddddc110e36f652cd11fb"></a>

<a id="canonical-c960705b67d3fa49e8104c23a85b8f47117163ca7b64fde802947a8a6151ceeb"></a>

## regex_values property — server_name_matcher / 7e9bb0f38c0f / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f9e48c05bf29be7fc96ae079bd388b3df4e44ddf7b5a4b18fa91243b6640eca8"></a>

## Next pages — server_name_matcher / 7e9bb0f38c0f / 6

- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-b0a5251ead8f0ceabde02cf071dd72a814033cc432b5c122b67826a1cfa78600"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a11ce237eefd1fdb5eb7c9e5edab4e53652716c9d1be531c8019a524ce395a2c"></a>

## server_selector — server_selector / 0bec82b1781c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- server_selector

<a id="canonical-7cf788cf70ecfcc93b52e2c1f6d7e5b53b40c6b638a3875879cfaf100fe4685a"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-c69487686583f2182c3179c78d2b8e08e3dd30f6ff7bdd828ee69a1d681f253e"></a>

## Direct properties — server_selector / 0bec82b1781c / 3

<a id="canonical-6bd95e858c9db6b4cd0e5870ce32ca296989813022fc6c2357eb70c7937cdc1f"></a>

<a id="canonical-cbb3eb68dc2821272ac48bf20bb33d2c796457f5aa839779535fc8f7c4301527"></a>

## expressions property — server_selector / 0bec82b1781c / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-a995e1e6eb83ff5d6c4020da826bfd0e7707ae13394cd3500a1f3ecb4f17921a"></a>

## Next pages — server_selector / 0bec82b1781c / 5

- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-9d6a07279db3ea5414f61669cbab5bebf9ed356fff30a5e35a33d069400bad3c)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
