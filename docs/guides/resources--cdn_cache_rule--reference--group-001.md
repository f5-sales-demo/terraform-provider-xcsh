---
page_title: "xcsh_cdn_cache_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule reference."
---

# xcsh_cdn_cache_rule reference

<a id="canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a12e216c7bd0aa207075a08fddfd46be4e9491a031cec1009c3b8220e6112be9"></a>

## Property reference — Property reference / 616950b09b03 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- Property reference

<a id="canonical-ea3beba3a6b9f7ba411951b5c6b7dac64fd986dc0ee3ef1f12e5b5044d3ac5a8"></a>

## Direct properties — Property reference / 616950b09b03 / 3

<a id="canonical-1569f9e7c9da44e33a2c92ab10602e4dbfa3e79ade7d9f2100911841712cd568"></a>

<a id="canonical-fd6fa38684807f70a856cd6bae3dbde7044093aa4d4af6b8555d660ddd421f35"></a>

## annotations property — Property reference / 616950b09b03 / 4

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

- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e): complete subsection reference.

<a id="canonical-1dcdc2ccf6188a8815020cb6f1124a9b4bcc65fd7ea6c35f729a5f0441f70e67"></a>

<a id="canonical-0c20e1f75b63fc5e6d3c8f14dca284fcfa1e8f8f0e5cd0c9e84c79cc098bad77"></a>

## description property — Property reference / 616950b09b03 / 5

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

<a id="canonical-5242aca943748c1e6d851c5d73e6923ce2e9a2cb951b8c83f9aeb96e814cc7c1"></a>

<a id="canonical-06283b8a90f8bc33ae267807ad4743855a56f28f3e1fb0a6a961bb1a6f9d0f57"></a>

## disable property — Property reference / 616950b09b03 / 6

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

<a id="canonical-fdd5ba85973ae7ff449624fb4013fd43788de7cf575ed92bbba353e06f4ac12e"></a>

<a id="canonical-de449a3598d6f21e799af10631673d7fe5d028e4d46aa17de8841eb4c65a4fe2"></a>

## id property — Property reference / 616950b09b03 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-b2680fabaacf78cbe0a1a6cc9225af4e387346e4d9355f3ad9da4108032fdf34"></a>

<a id="canonical-423ccc4e3dbac2c9959e8be9b49145f60ebe514a2b2601543a13eccb7808c5d6"></a>

## labels property — Property reference / 616950b09b03 / 8

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

<a id="canonical-6a64dde82d0d134533b731465b5c4300952708e4224623c8fa7ccaa8a2631635"></a>

<a id="canonical-319b01123b76ad03e1f33974aa52ec277e63feeec5f8a491c1aad479985644ba"></a>

## name property — Property reference / 616950b09b03 / 9

Type: `"string"`. Required.

Name of the CDN Cache Rule. Must be unique within the namespace.

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

<a id="canonical-6360402ed01cb85e62ada773158071c0d91e2b02659724704296c646cd2d50af"></a>

<a id="canonical-c8d84b613ad9d094c25159d5d8e5e1f277690d2d2c444ddbfcd686e36283be32"></a>

## namespace property — Property reference / 616950b09b03 / 10

Type: `"string"`. Required.

Namespace where the CDN Cache Rule is created.

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

- [timeouts](resources--cdn_cache_rule--reference--group-001.md#canonical-663cb96ecde51e43bfa858046ce1016c061e90dd3d1d52335ae31d2a63ed7b6c): complete subsection reference.

<a id="canonical-2e84784d05a3b6843d1d714240bf584e60a21acbf5096d88f98bbccffe7e58cd"></a>

## All schema paths — Property reference / 616950b09b03 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cdn_cache_rule--reference--group-001.md#canonical-1569f9e7c9da44e33a2c92ab10602e4dbfa3e79ade7d9f2100911841712cd568) |
| `cache_rules` | [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-5afc149ab9dcbda183181b20e36cd2af09baa98e75cd67c89c4908ec036310bd) |
| `cache_rules.cache_bypass` | [cache_rules.cache_bypass](resources--cdn_cache_rule--reference--group-001.md#canonical-01e592a1218b78453bcd4d652d3ac543e1f137dd76aa476189c69517930127be) |
| `cache_rules.eligible_for_cache` | [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-c273a99a60c7b6c618116eef1eea638a40e78823bdfd2831104787f8b8da1c6e) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-d85f21e5a7acd632dbebdf3cef72b03473b0e73905925d7088cd7b75c912f69f) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override](resources--cdn_cache_rule--reference--group-001.md#canonical-28f8c4f0f2497ca115081eaa5ce21d539db4fd40440165eda18bb5bf459cc171) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl](resources--cdn_cache_rule--reference--group-001.md#canonical-2ac656b4c3328119ccba5524568dacecf7d80ceb2316ef9f7758181b38d1de99) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie](resources--cdn_cache_rule--reference--group-001.md#canonical-7f26a7f366f991a5b1f422dd286d9d994800f02bbb719b2989f36533696f1d4d) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-8070fa9ef069534a684bc742e00e1cf55b8527febb389f6f49a789a80eb351fa) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override](resources--cdn_cache_rule--reference--group-001.md#canonical-f0eff59d3d9d34ba5a689278c491f3bc3734da0aa4be1ea88d98a4c16f9cf46d) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl](resources--cdn_cache_rule--reference--group-001.md#canonical-210688d569242332ca455f5043824c312145599a9367b4e962d6e6ec6db2247a) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie](resources--cdn_cache_rule--reference--group-001.md#canonical-81e130c2f3b0bc0e638f8da4e33ecc43bce3b28ab23b923afb8c274ae246b571) |
| `cache_rules.rule_expression_list` | [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-85425e671a4ffce2105a46e72522c42400f9f543f04bb4eea1cf0711b0c17623) |
| `cache_rules.rule_expression_list.cache_rule_expression` | [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-99522e78351d70677f1ad7fb7d6cc54553df8f336256ca67d62ccd6e6f56c8fd) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-ceed7b8c2ca0d7c0dc6e033aa26987152b82a402be9f73d512596af0f7cd1ae9) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name](resources--cdn_cache_rule--reference--group-001.md#canonical-0239f7d2a30e795d70d2b1db61f2d5f5b03ebe678d1cb8282e649aeddced7033) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-d05227e641efe763d07f1724ac56e6a7bda5fddf8594273d2b8fcc719951595f) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-f18a4c70f18da4df31ec46e1b2e7fda0ca4fbfcc8f2e60085a385362862c54d2) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-fc322c9c2cc06df20de1528ece60b3631c67235652ed8cbfad1c409be95d79a6) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-936f2b931b46d19c074e8a9a2088a5fae17f35ccacef689068943806f7264e39) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-1a1e9f4c8af7aa4e04beb56f0351c84e627d998cfadf116511b672623e80e71a) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-e4083cc075003ab6506f677ca88277755d2a617b45c09cb8986374de82650020) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-8649f8e4972424926846efc696debc8298a31acb0ef931b0bb9254555b7e09b1) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-e1b71edb958c1e05f09f00f7516a1d6c42a5270a3ad577d1968d1c98ff4d439e) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-a2e405f55870e5281a198420d830a55a267fd0abe47309c309390ff6e0470c0e) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-0f1763deadd03e3ecdac756cc3094e8eabdd6a85f8145e7c44eff1663a8033ae) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-d90bd0aba97ceda9b760af95925e9f71e05194c944c54979bb8b91b9ba0f5ce2) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name](resources--cdn_cache_rule--reference--group-001.md#canonical-8c02612e8c5b4853baedd726897ffeaac40321ac291c34d708b72064a2174046) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-a7ce510e23fba4dbe2c02582b82ed6ea9932f2522b0d4d3eb2b8ecfdd4f7b17c) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-b0132b6287b3a5bafecf64786e1290696a43b15b3a3dcab943d33d95290cbfe9) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-99b390b1f596396d09ea49abc677d2eab26e8143690ebb5ee2966dbf39c6b7c9) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-38503109d4f206d4c7941dadbc42dbc4c02b49f3dc82dc0786d7daf16169483f) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-80798151163cb4f5ff45553159923eedac944050254ccdea4ef6d34a41354196) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-c19e759eed9dd3c1aca1f2b0265d087a92ee586079b1a732175746647d98b565) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-9749ed5f0b41d2577022d23d8d4d92c55e8b4469204ed43f32e003ed5a1fd54a) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-1140b64dd09510b28bbe959d423d4c4d6a38ebaca23d934471d831f94cc242ec) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-61b676711d71962f3e4249887cdcd7f1821bc8566d7c7108c3d13e32ad9a85da) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-cef8396b9e26f2aad66156266558d0cb021e403ace37b18b85994d326a28ae7f) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match` | [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-ac2414fdca8fbe59ac1a50552e1e08bd954e1375c2650e5af9c46afa6713d280) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-97080f2ef4fe51d49de860751347e6f65c960c19e8f2765c88981f0ec28230a9) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-2780694750eb385f126de6e2f595443a337c4cc65d6bb8c5383ad9c8e51718db) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-f82b9a1bc74e7d1cb8e2c985b54d39bd45e3f15896938b3b57662619dec53bb3) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-5dd917e5331e5724abac6a304785338fb9327b0c26632e0788566fa8bc80e9f5) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-147c2f488f551f3f056f5b2085b7c0d7c616f6ace2d4cdea8c99352f7f0a249e) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-b703c063d59349dc3e6b949834637b56f487619797e412f8f54d4d98bc9235c7) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-ca46f6830b240ac9da363a2179f4c0ad1fd961a4656e9706069fb4d44e97361e) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-5a1748e3638b84d1326d34ebf8d3fa9d2a16a6f727e39ed063bc442e0cc71a63) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-ef58be7319efc3ab9dad821b8d2537c2af3debc6794c4782aadcd6045c2bdd07) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-5f392aea846a547ddb2a0a5040a6dcb55d7ef2e7b16172a304f1bb330698c09d) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-b5770c6fb04ca3a11ab2d076cb01173fc5921f56206095de7381b468da8eb4be) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key](resources--cdn_cache_rule--reference--group-001.md#canonical-a5f1d3b15fd729f9f38f784e81c2c217c50fd0ea147454f6e3a5a0bf69b54c7d) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-75365384f1a39c26db9608d27aadfd85259aa13f64d9a66437429a690f8fa1e8) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-7e8aa7f00fd8791ab35428d57386535ce46a3c9810d66b979261c87adf43eefb) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-c5ebeccbe286b7100cb3ca622864f8a941867283e0c77f862b59ed587f542981) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-0505d6caf37d12d2003c02bed140fef2ae19c6869d5b3a20984a9b0eede3b60d) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-4716e9ef00b2432bab902171bf2dc06d1388f2ca39a1c44617b041d1804e5921) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-b3bd86873295d1ee37c2df47ff90628c20c0a301e632589f145a11690addbab5) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-35f37fc82b0934ed464c4ab429ad8a8be9c523c7e474510bddbd93e1a42560ce) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-c47bae5aa6861944e82bcea02bb9a543da82daa4d700d9a1e0f18ccb03db01c9) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-f5785559c6cb4ed6594e1aeee96c3f705f7ae6c17baf8ebf3c28e0f2151ffca0) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-b4f6eaf217eb3fe40670437f52771cbaf15de494bd3a0080db04a6028de1280a) |
| `cache_rules.rule_expression_list.expression_name` | [cache_rules.rule_expression_list.expression_name](resources--cdn_cache_rule--reference--group-001.md#canonical-7700b8dac73ea01e4fecfe50d45560a03e096ca7caf268ac720577f8a25ae26d) |
| `cache_rules.rule_name` | [cache_rules.rule_name](resources--cdn_cache_rule--reference--group-001.md#canonical-c843afd6613a9e6adb3fdc3fdf4eba05e4ea1dff1cfe072a25c04cb34c9de8a7) |
| `description` | [description](resources--cdn_cache_rule--reference--group-001.md#canonical-1dcdc2ccf6188a8815020cb6f1124a9b4bcc65fd7ea6c35f729a5f0441f70e67) |
| `disable` | [disable](resources--cdn_cache_rule--reference--group-001.md#canonical-5242aca943748c1e6d851c5d73e6923ce2e9a2cb951b8c83f9aeb96e814cc7c1) |
| `id` | [id](resources--cdn_cache_rule--reference--group-001.md#canonical-fdd5ba85973ae7ff449624fb4013fd43788de7cf575ed92bbba353e06f4ac12e) |
| `labels` | [labels](resources--cdn_cache_rule--reference--group-001.md#canonical-b2680fabaacf78cbe0a1a6cc9225af4e387346e4d9355f3ad9da4108032fdf34) |
| `name` | [name](resources--cdn_cache_rule--reference--group-001.md#canonical-6a64dde82d0d134533b731465b5c4300952708e4224623c8fa7ccaa8a2631635) |
| `namespace` | [namespace](resources--cdn_cache_rule--reference--group-001.md#canonical-6360402ed01cb85e62ada773158071c0d91e2b02659724704296c646cd2d50af) |
| `timeouts` | [timeouts](resources--cdn_cache_rule--reference--group-001.md#canonical-43ee98d5b7c3c2f607cbf601911a8af9b7990f3a8de411d6c884a32c565f9039) |
| `timeouts.create` | [timeouts.create](resources--cdn_cache_rule--reference--group-001.md#canonical-a457eab41a1ac3b922bd953eed8b5f67187078ac18edde77ee5e2b7853268f4d) |
| `timeouts.delete` | [timeouts.delete](resources--cdn_cache_rule--reference--group-001.md#canonical-e68778ca0c3ed2e5d8f45e96c1ea0d1c219e83eae43ace3006cac4ddd525f07c) |
| `timeouts.read` | [timeouts.read](resources--cdn_cache_rule--reference--group-001.md#canonical-1e1535909b9c412e57afaf60b1472b865a61f3aa7431d16527a93940577f49a6) |
| `timeouts.update` | [timeouts.update](resources--cdn_cache_rule--reference--group-001.md#canonical-d14c588390847ac4898eed1d9d116b2c6a55d3072f6e4a1dc2d8e689cacb382d) |

<a id="canonical-f54c3823144e409bb843eea531cc2a8b05a856ca8dd5d5e5c3a07e56ee95f287"></a>

## Next pages — Property reference / 616950b09b03 / 12

- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [timeouts](resources--cdn_cache_rule--reference--group-001.md#canonical-663cb96ecde51e43bfa858046ce1016c061e90dd3d1d52335ae31d2a63ed7b6c)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4de794d49d4a365152d0fe1a40d19fe77553ab563b7f498221db9f4c4bc7146"></a>

## cache_rules — cache_rules / 5a33461324f2 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- cache_rules

<a id="canonical-5afc149ab9dcbda183181b20e36cd2af09baa98e75cd67c89c4908ec036310bd"></a>

Type: `"object"`. single nested block, Optional.

Cache Rule. This defines a CDN Cache Rule.

Upstream description:

This defines a CDN Cache Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule_expression_list",
    "rule_name"),
  validators.ConflictingObjectAttributes("cache_bypass",
    "eligible_for_cache")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_bypass\",\"eligible_for_cache\"]"
}
```

Terraform syntax:

```terraform
cache_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc01963e1cf8ce2f8240f752377839c48431b5a3ce91901b13bf53e87602361b"></a>

## Direct properties — cache_rules / 5a33461324f2 / 3

- [cache_bypass](resources--cdn_cache_rule--reference--group-001.md#canonical-9a1dcb9f1512c822a2e7410523809c593e7bc00ffe3297e1320d2d46144546e6): complete subsection reference.

- [eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-ebbb44b9c3fd7b4317bd0acf420193e284f7cf77ecc4c44582dc3abd38b94090): complete subsection reference.

- [rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e): complete subsection reference.

<a id="canonical-c843afd6613a9e6adb3fdc3fdf4eba05e4ea1dff1cfe072a25c04cb34c9de8a7"></a>

<a id="canonical-28c0b7f810142c579892e31fff9a1de8a051f3fdece5461083cb219c76936aec"></a>

## rule_name property — cache_rules / 5a33461324f2 / 4

Type: `"string"`. Optional.

Rule Name. Name of the Cache Rule.

Upstream description:

Name of the Cache Rule.

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

<a id="canonical-bb1648ba7f76edf0507cd3578d1077ceb5c935e842acdbb81a6e2dde966c7b57"></a>

## Next pages — cache_rules / 5a33461324f2 / 5

- [cache_rules.cache_bypass](resources--cdn_cache_rule--reference--group-001.md#canonical-9a1dcb9f1512c822a2e7410523809c593e7bc00ffe3297e1320d2d46144546e6)
- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-ebbb44b9c3fd7b4317bd0acf420193e284f7cf77ecc4c44582dc3abd38b94090)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-9a1dcb9f1512c822a2e7410523809c593e7bc00ffe3297e1320d2d46144546e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c663fd9fe197e8944427bcb21bac44deee2f63796fff00ea9b1ff2bd1aeca559"></a>

## cache_rules.cache_bypass — cache_rules.cache_bypass / a0342f04e67b / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- cache_rules.cache_bypass

<a id="canonical-01e592a1218b78453bcd4d652d3ac543e1f137dd76aa476189c69517930127be"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for cache bypass.

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
cache_bypass = {}
```

<a id="canonical-14d2a0321e36ab7a763c4bafadb3683ce38b70270643ca851e28d485a1e535d5"></a>

## Direct properties — cache_rules.cache_bypass / a0342f04e67b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5f607e1eff1e53b0c37507d30f5277dff9fdd869e3c259e9a1f7df180cd056f4"></a>

## Next pages — cache_rules.cache_bypass / a0342f04e67b / 4

- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-ebbb44b9c3fd7b4317bd0acf420193e284f7cf77ecc4c44582dc3abd38b94090"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6f826a295bde67e2205dc6af2dcdd61d809aef16141ddcea0f03caaee068e70"></a>

## cache_rules.eligible_for_cache — cache_rules.eligible_for_cache / 613699e4dde2 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- cache_rules.eligible_for_cache

<a id="canonical-c273a99a60c7b6c618116eef1eea638a40e78823bdfd2831104787f8b8da1c6e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eligible for cache.

Upstream description:

List of OPTIONS for Cache Action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("scheme_proxy_host_request_uri",
    "scheme_proxy_host_uri")}
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
  "x-ves-oneof-field-eligible_for_cache": "[\"scheme_proxy_host_request_uri\",\"scheme_proxy_host_uri\"]"
}
```

Terraform syntax:

```terraform
eligible_for_cache {
  # Configure direct properties listed below.
}
```

<a id="canonical-69b3724b16669ad8b44810de09b80816118285f56b775335a2e868fa5f6fb246"></a>

## Direct properties — cache_rules.eligible_for_cache / 613699e4dde2 / 3

- [scheme_proxy_host_request_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-7bb19ae02699d3c749b7f402452383f0dd8b3356c652db7eeaef4bcdb66d94c6): complete subsection reference.

- [scheme_proxy_host_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-dd207207d50bc998fdede2eb95dc74888d484a8c9941d179d05454ab110ff766): complete subsection reference.

<a id="canonical-113f20f5a8341ada2dc43ef53f7fb600ca3f59379d19e510cbb82308f2e2ff6d"></a>

## Next pages — cache_rules.eligible_for_cache / 613699e4dde2 / 4

- [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-7bb19ae02699d3c749b7f402452383f0dd8b3356c652db7eeaef4bcdb66d94c6)
- [cache_rules.eligible_for_cache.scheme_proxy_host_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-dd207207d50bc998fdede2eb95dc74888d484a8c9941d179d05454ab110ff766)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-7bb19ae02699d3c749b7f402452383f0dd8b3356c652db7eeaef4bcdb66d94c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd84f28810e123ffad8d21a8b60800a18262fce155ac2d6a67803dd65611ab04"></a>

## cache_rules.eligible_for_cache.scheme_proxy_host_request_uri — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / 53076a9dd696 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-ebbb44b9c3fd7b4317bd0acf420193e284f7cf77ecc4c44582dc3abd38b94090)
- cache_rules.eligible_for_cache.scheme_proxy_host_request_uri

<a id="canonical-d85f21e5a7acd632dbebdf3cef72b03473b0e73905925d7088cd7b75c912f69f"></a>

Type: `"object"`. single nested block, Optional.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_ttl")}
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
scheme_proxy_host_request_uri {
  # Configure direct properties listed below.
}
```

<a id="canonical-18df878a4f8eb542987ebfbebb79fe08c16e55837fcfb2fc9dd566240fdc4091"></a>

## Direct properties — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / 53076a9dd696 / 3

<a id="canonical-28f8c4f0f2497ca115081eaa5ce21d539db4fd40440165eda18bb5bf459cc171"></a>

<a id="canonical-59c30af35ed47e4dcffc6f3f94adcb697c4671f49737d6f9551d6824c18d35ec"></a>

## cache_override property — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / 53076a9dd696 / 4

Type: `"bool"`. Optional.

Cache Override. Honour Cache Override.

Upstream description:

Honour Cache Override.

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

<a id="canonical-2ac656b4c3328119ccba5524568dacecf7d80ceb2316ef9f7758181b38d1de99"></a>

<a id="canonical-2d9b69f1b0f0d876c99e711f5ff4dbdc070387bed3a5512f2f66b7632a997423"></a>

## cache_ttl property — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / 53076a9dd696 / 5

Type: `"string"`. Optional.

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

Upstream description:

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-7f26a7f366f991a5b1f422dd286d9d994800f02bbb719b2989f36533696f1d4d"></a>

<a id="canonical-500e153c76ec743e3d20bfc55d832a5f1b65f3ea086bec80a623313b9ef9f33b"></a>

## ignore_response_cookie property — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / 53076a9dd696 / 6

Type: `"bool"`. Optional.

By default, response will not be cached if set-cookie header is present. This option will override
the behavior and cache response even with set-cookie header present.

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

<a id="canonical-2102524b33d695b7b77ee3b4624a37c6f8583ae15cecc7a49de9226ce18998b2"></a>

## Next pages — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / 53076a9dd696 / 7

- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-ebbb44b9c3fd7b4317bd0acf420193e284f7cf77ecc4c44582dc3abd38b94090)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-dd207207d50bc998fdede2eb95dc74888d484a8c9941d179d05454ab110ff766"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22e8d7d3fd9c0309c743f44b19a7c6d8e435a80b12cb381911751f7677ac4925"></a>

## cache_rules.eligible_for_cache.scheme_proxy_host_uri — cache_rules.eligible_for_cache.scheme_proxy_host_uri / 7c1c4b9f7450 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-ebbb44b9c3fd7b4317bd0acf420193e284f7cf77ecc4c44582dc3abd38b94090)
- cache_rules.eligible_for_cache.scheme_proxy_host_uri

<a id="canonical-8070fa9ef069534a684bc742e00e1cf55b8527febb389f6f49a789a80eb351fa"></a>

Type: `"object"`. single nested block, Optional.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_ttl")}
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
scheme_proxy_host_uri {
  # Configure direct properties listed below.
}
```

<a id="canonical-6b72bdf7e721b05f18acff128d8f9b514d442a14ae774aa372ebc9bc1db3fc40"></a>

## Direct properties — cache_rules.eligible_for_cache.scheme_proxy_host_uri / 7c1c4b9f7450 / 3

<a id="canonical-f0eff59d3d9d34ba5a689278c491f3bc3734da0aa4be1ea88d98a4c16f9cf46d"></a>

<a id="canonical-aa09ff72adcb50e107da074989ebc2d5a339cc577114452c8bf4911189199531"></a>

## cache_override property — cache_rules.eligible_for_cache.scheme_proxy_host_uri / 7c1c4b9f7450 / 4

Type: `"bool"`. Optional.

Cache Override. Honour Cache Override.

Upstream description:

Honour Cache Override.

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

<a id="canonical-210688d569242332ca455f5043824c312145599a9367b4e962d6e6ec6db2247a"></a>

<a id="canonical-9f21267c801bf0239c315248abff80de87ebf83a488334bc9a1d2d8506694021"></a>

## cache_ttl property — cache_rules.eligible_for_cache.scheme_proxy_host_uri / 7c1c4b9f7450 / 5

Type: `"string"`. Optional.

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

Upstream description:

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-81e130c2f3b0bc0e638f8da4e33ecc43bce3b28ab23b923afb8c274ae246b571"></a>

<a id="canonical-ad0f51ff614ca53b0705241ab2ebe69a73a9aa4e10c8c46834f480c14e152b7e"></a>

## ignore_response_cookie property — cache_rules.eligible_for_cache.scheme_proxy_host_uri / 7c1c4b9f7450 / 6

Type: `"bool"`. Optional.

By default, response will not be cached if set-cookie header is present. This option will override
the behavior and cache response even with set-cookie header present.

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

<a id="canonical-71290d92592c4d72754d16abcec54b1e4ba296fda0c0226e1700ebd2159f63ec"></a>

## Next pages — cache_rules.eligible_for_cache.scheme_proxy_host_uri / 7c1c4b9f7450 / 7

- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-ebbb44b9c3fd7b4317bd0acf420193e284f7cf77ecc4c44582dc3abd38b94090)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-454c2ffc81ca59d09441879d0144c98738029e00ca4fa12470a6d36259571886"></a>

## cache_rules.rule_expression_list — cache_rules.rule_expression_list / 285910683d2b / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- cache_rules.rule_expression_list

<a id="canonical-85425e671a4ffce2105a46e72522c42400f9f543f04bb4eea1cf0711b0c17623"></a>

Type: `"object"`. list nested block, Optional.

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs.

Upstream description:

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs..

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cache_rule_expression",
    "expression_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rule_expression_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-6ea52c05cf33b314ac855432ec93fcfc5cb34362c7c41c389bf25dca030041ce"></a>

## Direct properties — cache_rules.rule_expression_list / 285910683d2b / 3

- [cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5): complete subsection reference.

<a id="canonical-7700b8dac73ea01e4fecfe50d45560a03e096ca7caf268ac720577f8a25ae26d"></a>

<a id="canonical-2a6dff5a803e7295ea059c8d862cb4bdc011a167f100c9abaccf7a3bbfab8c86"></a>

## expression_name property — cache_rules.rule_expression_list / 285910683d2b / 4

Type: `"string"`. Optional.

Name of the Expressions items that are ANDed.

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

<a id="canonical-46e1d32c12d9b30c521e7428d55ff9865a4dd565386ffc8893a116aa3ed94022"></a>

## Next pages — cache_rules.rule_expression_list / 285910683d2b / 5

- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f51bf8699b14b2962ea23c4835babc36507268162500a86baa6cc9027c1b257a"></a>

## cache_rules.rule_expression_list.cache_rule_expression — cache_rules.rule_expression_list.cache_rule_expression / e1c3622c2250 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- cache_rules.rule_expression_list.cache_rule_expression

<a id="canonical-99522e78351d70677f1ad7fb7d6cc54553df8f336256ca67d62ccd6e6f56c8fd"></a>

Type: `"object"`. list nested block, Optional.

The Cache Rule Expression Terms that are ANDed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cache_rule_expression {
  # Configure direct properties listed below.
}
```

<a id="canonical-20158a8c9967fe12a3741e17aa6286c0d81203e6686d0862c200fa49dcc6f1e8"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression / e1c3622c2250 / 3

- [cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-0986b7db601d9bbc10f0249d7b2aab0bff189fbc3f6b38cfad5fd7ace17bb9e8): complete subsection reference.

- [cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-fdab340a8ec0b9e25e1f8f6e84ac9ca4526af1cec6f1db355af281236495975d): complete subsection reference.

- [path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-42434708340d777587211e8f1879f43a3ed3e21f55faed301bcc2accfc877827): complete subsection reference.

- [query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-5e7f47d34af0ab20e193109b18292e36903ac8e35cf30c238aed231f3a4e68b4): complete subsection reference.

<a id="canonical-c99a343ff03c49247180da684ebd4b5879bac1e50e6592f68c9d2c334d246787"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression / e1c3622c2250 / 4

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-0986b7db601d9bbc10f0249d7b2aab0bff189fbc3f6b38cfad5fd7ace17bb9e8)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-fdab340a8ec0b9e25e1f8f6e84ac9ca4526af1cec6f1db355af281236495975d)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-42434708340d777587211e8f1879f43a3ed3e21f55faed301bcc2accfc877827)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-5e7f47d34af0ab20e193109b18292e36903ac8e35cf30c238aed231f3a4e68b4)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-0986b7db601d9bbc10f0249d7b2aab0bff189fbc3f6b38cfad5fd7ace17bb9e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c508f8462f0397b329cbe488652c29b5a3164a864f9dca331745db4fbe8c560e"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cache_headers — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / 0d52afb60c11 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers

<a id="canonical-ceed7b8c2ca0d7c0dc6e033aa26987152b82a402be9f73d512596af0f7cd1ae9"></a>

Type: `"object"`. list nested block, Optional.

Configure cache rule headers to match the criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cache_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-92866f6ccf0ef941ffe6826380c9aaac72c747243dbb4c6dbe7f2c34644ce0fd"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / 0d52afb60c11 / 3

<a id="canonical-0239f7d2a30e795d70d2b1db61f2d5f5b03ebe678d1cb8282e649aeddced7033"></a>

<a id="canonical-3fe57d17ef6ce7174b7d3c2dd53c109618b810cd25b543750d2b21fd34f06892"></a>

## name property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / 0d52afb60c11 / 4

Type: `"string"`. Optional.

\[Enum: PROXY\_HOST|REFERER|SCHEME|USER\_AGENT\] - PROXY\_HOST: Proxy Host Name of the proxied
server - REFERER: Referer This is the address of the previous web page from which a link to the
currently requested page was followed - SCHEME: Scheme The HTTP scheme used: HTTP or HTTPS -
USER\_AGENT: User Agent The user agent string of the user agent. Possible values are
\`PROXY\_HOST\`, \`REFERER\`, \`SCHEME\`, \`USER\_AGENT\`. Defaults to \`PROXY\_HOST\`.

Upstream description:

&#8203;- PROXY\_HOST: Proxy Host

Name of the proxied server &#8203;- REFERER: Referer

This is the address of the previous web page from which a link to the currently requested page was
followed &#8203;- SCHEME: Scheme

The HTTP scheme used: HTTP or HTTPS &#8203;- USER\_AGENT: User Agent

The user agent string of the user agent.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROXY_HOST",
    "REFERER",
    "SCHEME",
    "USER_AGENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROXY_HOST",
  "enum": [
    "PROXY_HOST",
    "REFERER",
    "SCHEME",
    "USER_AGENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-5834bcb6a558eeb6bcd5e49eae32dc2d355180e2614d762a8a14f99edf3c4407): complete subsection reference.

<a id="canonical-94946aaba5580e60812089ad0bee3eff957b457462ff0663095ee2bf8a698104"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / 0d52afb60c11 / 5

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-5834bcb6a558eeb6bcd5e49eae32dc2d355180e2614d762a8a14f99edf3c4407)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-5834bcb6a558eeb6bcd5e49eae32dc2d355180e2614d762a8a14f99edf3c4407"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f344b8fd215828e4675878abc7a25a85fb302973d0260c77c51986cb0def367"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-0986b7db601d9bbc10f0249d7b2aab0bff189fbc3f6b38cfad5fd7ace17bb9e8)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator

<a id="canonical-d05227e641efe763d07f1724ac56e6a7bda5fddf8594273d2b8fcc719951595f"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0a0df019d652b3a56476b5ab9852d0b20bd24d8daff4b543865193b2d3a0de7e"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 3

<a id="canonical-f18a4c70f18da4df31ec46e1b2e7fda0ca4fbfcc8f2e60085a385362862c54d2"></a>

<a id="canonical-fecd46d63e4409a6792d23afeba9f6cc8bd3793236da26f2e3c6fae01d312099"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 4

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The header value must include the specified value as a substring.

<a id="canonical-fc322c9c2cc06df20de1528ece60b3631c67235652ed8cbfad1c409be95d79a6"></a>

<a id="canonical-585526912bcc3470f490cb48fa97bd386c3a240464ca5a1be23c8d75fe000401"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 5

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not include the specified value as a substring.

<a id="canonical-936f2b931b46d19c074e8a9a2088a5fae17f35ccacef689068943806f7264e39"></a>

<a id="canonical-9b14a682ef3765b738c011ed0af8f42f704fc97934c927ffc56a500c245fd760"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 6

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not end with the specified value.

<a id="canonical-1a1e9f4c8af7aa4e04beb56f0351c84e627d998cfadf116511b672623e80e71a"></a>

<a id="canonical-5877b1b42222f37516565e51226210ad5598659f92e188bc19dc22ca84749691"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 7

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not match the specified value.

<a id="canonical-e4083cc075003ab6506f677ca88277755d2a617b45c09cb8986374de82650020"></a>

<a id="canonical-5d0b9f79e70d4e2040e720ee55a44b808c20698d964e175250fa6e2a5327a87c"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 8

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The header value must not begin with the specified value.

<a id="canonical-8649f8e4972424926846efc696debc8298a31acb0ef931b0bb9254555b7e09b1"></a>

<a id="canonical-0aa5bd3213756739de1a529119144004f7f8de7d59acfe8423447f0ea19e8a8d"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 9

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The header value must end with the specified value.

<a id="canonical-e1b71edb958c1e05f09f00f7516a1d6c42a5270a3ad577d1968d1c98ff4d439e"></a>

<a id="canonical-8e054ead18abe5bd062d9f5419b762c569dcacd5ce690f4ae859844fca6fe0fd"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 10

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The header value must exactly match the specified value.

<a id="canonical-a2e405f55870e5281a198420d830a55a267fd0abe47309c309390ff6e0470c0e"></a>

<a id="canonical-b42d2516214b3b93c63fc7fc7930faec9e0104fda350d902c3c2d5e5629197f7"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 11

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The header value must match the specified regular expression pattern.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-0f1763deadd03e3ecdac756cc3094e8eabdd6a85f8145e7c44eff1663a8033ae"></a>

<a id="canonical-699fe98b3eeb89b8b4701325c44c14693aec68715feb8aba6e6712b556dfd771"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 12

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The header value must begin with the specified value.

<a id="canonical-1ebc4912c78f626aafb1018780f3562ea1ea5bc1ece83087d894435aa66e0738"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 781fa1fc8723 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-0986b7db601d9bbc10f0249d7b2aab0bff189fbc3f6b38cfad5fd7ace17bb9e8)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-fdab340a8ec0b9e25e1f8f6e84ac9ca4526af1cec6f1db355af281236495975d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2005a9ebf517f8a74f303df7e800cd0925251213cfe4c536d3adb60b6ba19f5b"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / be084bd9b4af / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher

<a id="canonical-d90bd0aba97ceda9b760af95925e9f71e05194c944c54979bb8b91b9ba0f5ce2"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cookie_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-92a676e0a4037decf39af9b809d04687ce25649cf54bcbe4e1e4cfe4fc6f591d"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / be084bd9b4af / 3

<a id="canonical-8c02612e8c5b4853baedd726897ffeaac40321ac291c34d708b72064a2174046"></a>

<a id="canonical-f2f15f20061248cf0a97ad0f5a07832ac842a537ee05ff7fe276b4ab5f8ddcbe"></a>

## name property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / be084bd9b4af / 4

Type: `"string"`. Optional.

Cookie Name. Enter the name of the cookie to match.

Upstream description:

Enter the name of the cookie to match.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-ffe694caed217d7f8c46c6a4fe916b74939a200d9d12b4f68d83189dc440f5bf): complete subsection reference.

<a id="canonical-bdbf658ccb856342bfa2924710abbbb5f8d8d082ac8e9bc31a80ba74d3c05fdc"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / be084bd9b4af / 5

- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-ffe694caed217d7f8c46c6a4fe916b74939a200d9d12b4f68d83189dc440f5bf)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-ffe694caed217d7f8c46c6a4fe916b74939a200d9d12b4f68d83189dc440f5bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8f86fd24446fe161b67ecc268156623d0171c7267550d194cd1b16c7f2bd8d8"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-fdab340a8ec0b9e25e1f8f6e84ac9ca4526af1cec6f1db355af281236495975d)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator

<a id="canonical-a7ce510e23fba4dbe2c02582b82ed6ea9932f2522b0d4d3eb2b8ecfdd4f7b17c"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7ac067a4c89f6ec88b2ef56857a5f1da839be19e0a62c879d07c640b4cab10e"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 3

<a id="canonical-b0132b6287b3a5bafecf64786e1290696a43b15b3a3dcab943d33d95290cbfe9"></a>

<a id="canonical-661302d074c7125ffb08903e242215fea40b846e09d24f6970a3addeabb0c822"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 4

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The cookie value must include the specified value as a substring.

<a id="canonical-99b390b1f596396d09ea49abc677d2eab26e8143690ebb5ee2966dbf39c6b7c9"></a>

<a id="canonical-2dd2a4eedcf36f8fba6310251ff875472db4f79becb768c3154c7584d5b534f4"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 5

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not include the specified value as a substring.

<a id="canonical-38503109d4f206d4c7941dadbc42dbc4c02b49f3dc82dc0786d7daf16169483f"></a>

<a id="canonical-67b2a11cf1d016c0815a32cbebcc51386d7665031b149fb30ce06fd36ae4ac37"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 6

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not end with the specified value.

<a id="canonical-80798151163cb4f5ff45553159923eedac944050254ccdea4ef6d34a41354196"></a>

<a id="canonical-c32547dceaa6cea383dfedea22fa705789f6adff21dbe9fee8ec0891e5cc5694"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 7

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not match the specified value.

<a id="canonical-c19e759eed9dd3c1aca1f2b0265d087a92ee586079b1a732175746647d98b565"></a>

<a id="canonical-57e20664bc6de20c11498ea5e29eb0c3525d6e0895bf2c01d53eef042af6a30d"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 8

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The cookie value must not begin with the specified value.

<a id="canonical-9749ed5f0b41d2577022d23d8d4d92c55e8b4469204ed43f32e003ed5a1fd54a"></a>

<a id="canonical-3db6ca25dea57a72c8e4e697719a239f7d9f0ccdc7386da806882d5c5aa9551a"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 9

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The cookie value must end with the specified value.

<a id="canonical-1140b64dd09510b28bbe959d423d4c4d6a38ebaca23d934471d831f94cc242ec"></a>

<a id="canonical-cfd4a34cd3c78fa5d989ee874b12a294574196543c2b9ae51d410471d06bf16b"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 10

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The cookie value must exactly match the specified value.

<a id="canonical-61b676711d71962f3e4249887cdcd7f1821bc8566d7c7108c3d13e32ad9a85da"></a>

<a id="canonical-6b5a8015db7236fc75755c73bc1d0b2677174d86a0786e976a57f09c1ae642ab"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 11

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The cookie value must match the specified regular expression pattern in PCRE
format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-cef8396b9e26f2aad66156266558d0cb021e403ace37b18b85994d326a28ae7f"></a>

<a id="canonical-05da840a8396a0b245fe5b529e630a2c9937d1e1f925a0ed2375d508c9eb18a3"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 12

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The cookie value must begin with the specified value.

<a id="canonical-557bbaf3fd07105dd0317989e1bea4a306a6b5091bc519051b919112c566293c"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / cec36d588b11 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-fdab340a8ec0b9e25e1f8f6e84ac9ca4526af1cec6f1db355af281236495975d)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-42434708340d777587211e8f1879f43a3ed3e21f55faed301bcc2accfc877827"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-492028b53528297318948736f9b9ac3b280f67f5c5b8187917357493b7ead3fc"></a>

## cache_rules.rule_expression_list.cache_rule_expression.path_match — cache_rules.rule_expression_list.cache_rule_expression.path_match / 8ec95dd151c4 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- cache_rules.rule_expression_list.cache_rule_expression.path_match

<a id="canonical-ac2414fdca8fbe59ac1a50552e1e08bd954e1375c2650e5af9c46afa6713d280"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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
path_match {
  # Configure direct properties listed below.
}
```

<a id="canonical-a735370ea94083786735d86b2e821fc05b2df9db4824c7f77f5b03db3ddaebc0"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.path_match / 8ec95dd151c4 / 3

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-b224d5eab5f11dee639fa776bd1f51131e2b9cb1117c07c649815e93512d62dd): complete subsection reference.

<a id="canonical-5685f28375af20403613a26f299c47c32bda68bfb8867bf91033d70f0983689a"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.path_match / 8ec95dd151c4 / 4

- [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-b224d5eab5f11dee639fa776bd1f51131e2b9cb1117c07c649815e93512d62dd)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-b224d5eab5f11dee639fa776bd1f51131e2b9cb1117c07c649815e93512d62dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a206ea0a6e38a0923e7125f63761762f1545c48780b475d8db8f253546bb18b"></a>

## cache_rules.rule_expression_list.cache_rule_expression.path_match.operator — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-42434708340d777587211e8f1879f43a3ed3e21f55faed301bcc2accfc877827)
- cache_rules.rule_expression_list.cache_rule_expression.path_match.operator

<a id="canonical-97080f2ef4fe51d49de860751347e6f65c960c19e8f2765c88981f0ec28230a9"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-731c1ef7aacb96dc4e766ae0e6eb435af246b45d820e502256d03798fbc92201"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 3

<a id="canonical-2780694750eb385f126de6e2f595443a337c4cc65d6bb8c5383ad9c8e51718db"></a>

<a id="canonical-2f07a0ce83687b52445948038fb79ced58145eecad6bb5d5f3406da38b7f4ca6"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 4

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The path must include the specified value as a substring, up to the
filename.

<a id="canonical-f82b9a1bc74e7d1cb8e2c985b54d39bd45e3f15896938b3b57662619dec53bb3"></a>

<a id="canonical-53b46d1cd04d984ef7c320f5af67f5acea81b8aef015877fa552fbc95090209a"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 5

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not include the specified value as a substring, up to the filename.

<a id="canonical-5dd917e5331e5724abac6a304785338fb9327b0c26632e0788566fa8bc80e9f5"></a>

<a id="canonical-40be5bd9ae7db6b21b96547688bb57e284e07277c1ab07dd6f05782f68a03a7d"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 6

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not end with the specified value, up to the filename.

<a id="canonical-147c2f488f551f3f056f5b2085b7c0d7c616f6ace2d4cdea8c99352f7f0a249e"></a>

<a id="canonical-6494e720493af011dfa74f794a87198b650bbccb8f41c4ece4bb17946bd8587a"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 7

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not match the specified value, up to the filename.

<a id="canonical-b703c063d59349dc3e6b949834637b56f487619797e412f8f54d4d98bc9235c7"></a>

<a id="canonical-6fd8255f45ced69cf03436a15c28fc25ebdeed8ec360390f2bb6b1ea12a07032"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 8

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The path must not begin with the specified value, up to the filename.

<a id="canonical-ca46f6830b240ac9da363a2179f4c0ad1fd961a4656e9706069fb4d44e97361e"></a>

<a id="canonical-d657c0719f12b736f14b7d42ab89b05dce3aaccdde5912e5c28b66e419a03bef"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 9

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The path must end with the specified value, up to the filename.

<a id="canonical-5a1748e3638b84d1326d34ebf8d3fa9d2a16a6f727e39ed063bc442e0cc71a63"></a>

<a id="canonical-3c8df0dce87489cd62f007438ec377c71f38bb720ce1a0805b359aada8b7acc3"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 10

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The path must exactly match the specified value, up to the filename.

<a id="canonical-ef58be7319efc3ab9dad821b8d2537c2af3debc6794c4782aadcd6045c2bdd07"></a>

<a id="canonical-065dbc4dc30b690a954e5c076e887648696d8c570d17e8a9b2f0b66c09b18380"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 11

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The path must match the specified regular expression pattern in PCRE format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-5f392aea846a547ddb2a0a5040a6dcb55d7ef2e7b16172a304f1bb330698c09d"></a>

<a id="canonical-3e36dfe0c30c3002d1567f3c197218d12256beaf04161e8d99a9cbb17780a21b"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 12

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The path must begin with the specified value, up to the filename.

<a id="canonical-ad8c8413bf80592a9bb0fc5022f02e0de89fbe921c1e477cc3fc81d070052bf8"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / c5715576c385 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-42434708340d777587211e8f1879f43a3ed3e21f55faed301bcc2accfc877827)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-5e7f47d34af0ab20e193109b18292e36903ac8e35cf30c238aed231f3a4e68b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bcf507525530ec9a3caddba55045621c8e66181102f9d8e345f5bf263471ec8"></a>

## cache_rules.rule_expression_list.cache_rule_expression.query_parameters — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / b666190e1b04 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters

<a id="canonical-b5770c6fb04ca3a11ab2d076cb01173fc5921f56206095de7381b468da8eb4be"></a>

Type: `"object"`. list nested block, Optional.

Query Parameters. List of (key, value) query parameters.

Upstream description:

List of (key, value) query parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc53370ea272ff14bc28f988def24f4cb3d902c4cce954de4535553d3b974847"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / b666190e1b04 / 3

<a id="canonical-a5f1d3b15fd729f9f38f784e81c2c217c50fd0ea147454f6e3a5a0bf69b54c7d"></a>

<a id="canonical-e1913692021f20b88af5c390279af81d581a09abfc31d387d41f29309afbd6e6"></a>

## key property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / b666190e1b04 / 4

Type: `"string"`. Optional.

The name of the query parameter to match.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-890d36bde9c41cf0e1ad1400ecdec59a7d76f3edfef48d5e3fa968cafa55d22d): complete subsection reference.

<a id="canonical-ce8b868018eb5dfc4acfe5a283357a51c678213104597c5f47901b674a37caf9"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / b666190e1b04 / 5

- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-890d36bde9c41cf0e1ad1400ecdec59a7d76f3edfef48d5e3fa968cafa55d22d)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-890d36bde9c41cf0e1ad1400ecdec59a7d76f3edfef48d5e3fa968cafa55d22d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edb757e10f91156b999f676d6d1acff61e7e38a7566e50e41cdb1a8c45dc61b4"></a>

## cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-9038feb6b4b2bd3acd855e50198087c999e4b6ef507b31cdf45a768e4ec7504e)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-d4483c9db2d9f910cd4ad73455317199bf6bff3336fd3fc087659dddee9d2d6e)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-cff8e3de8ceaaff4cdabd1c026998ff6257c2cc914d2c03aeab3aac2b926e2a5)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-5e7f47d34af0ab20e193109b18292e36903ac8e35cf30c238aed231f3a4e68b4)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator

<a id="canonical-75365384f1a39c26db9608d27aadfd85259aa13f64d9a66437429a690f8fa1e8"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-9feed4ff9a90a06c9910f633151082f72c1295cc950f715397b91b05c36d8f58"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 3

<a id="canonical-7e8aa7f00fd8791ab35428d57386535ce46a3c9810d66b979261c87adf43eefb"></a>

<a id="canonical-9e97224eb66fa681038f25dba6571fc01e170af1b4671254012133ee0912eca7"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 4

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The query parameter value must include the specified value as a substring.

<a id="canonical-c5ebeccbe286b7100cb3ca622864f8a941867283e0c77f862b59ed587f542981"></a>

<a id="canonical-4003407946b7d3c4475304b84e745dd1f200d6413f1f28d90d9204792bec4f68"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 5

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not include the specified value as a substring.

<a id="canonical-0505d6caf37d12d2003c02bed140fef2ae19c6869d5b3a20984a9b0eede3b60d"></a>

<a id="canonical-e9f56f50961ba7fe73647dd93b7270537c92cf8444c1922589776a251183d0a4"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 6

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not end with the specified value.

<a id="canonical-4716e9ef00b2432bab902171bf2dc06d1388f2ca39a1c44617b041d1804e5921"></a>

<a id="canonical-add301a8443a749be8045128dd369e56c2e9a00494b865a729f623d02093090c"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 7

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not match the specified value.

<a id="canonical-b3bd86873295d1ee37c2df47ff90628c20c0a301e632589f145a11690addbab5"></a>

<a id="canonical-33f99e7c7781fd6dedb55cd9fff5c1939e8b82a42b43a2bdabc238e568cf9834"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 8

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The query parameter value must not begin with the specified value.

<a id="canonical-35f37fc82b0934ed464c4ab429ad8a8be9c523c7e474510bddbd93e1a42560ce"></a>

<a id="canonical-38965001197dfe6a81cb33ab2a65410aaeaad3947c7a3e38e8a513a026f3da0d"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 9

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The query parameter value must end with the specified value.

<a id="canonical-c47bae5aa6861944e82bcea02bb9a543da82daa4d700d9a1e0f18ccb03db01c9"></a>

<a id="canonical-37c4e52023ab7559b3a7e4a5f4c89c113471300ca96f9113794e131e412777df"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 10

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The query parameter value must exactly match the specified value.

<a id="canonical-f5785559c6cb4ed6594e1aeee96c3f705f7ae6c17baf8ebf3c28e0f2151ffca0"></a>

<a id="canonical-8c9515ca9403e3fb121741ce625092d9fcf97b4dadacbda1d068c8477792f7d7"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 11

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The query parameter value must match the specified regular expression pattern in
PCRE format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-b4f6eaf217eb3fe40670437f52771cbaf15de494bd3a0080db04a6028de1280a"></a>

<a id="canonical-c89ed2758e677bbde2496de4a52a03a2ba0081c79eefed6cfcfb55531995c26f"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 12

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The query parameter value must begin with the specified value.

<a id="canonical-c3f2b511617ac223571026a4174a1a5c9386bf281768fc116ff844575a1f56ef"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / 9ac1bd475041 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-5e7f47d34af0ab20e193109b18292e36903ac8e35cf30c238aed231f3a4e68b4)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-663cb96ecde51e43bfa858046ce1016c061e90dd3d1d52335ae31d2a63ed7b6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e7a9fdfb68429419a40c281c732fb43b58e2e2d5b97e0a70b5bf7b7170f1ef8"></a>

## timeouts — timeouts / 658babed0e04 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- timeouts

<a id="canonical-43ee98d5b7c3c2f607cbf601911a8af9b7990f3a8de411d6c884a32c565f9039"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-9506bc264f1c385b09922b3850d9e69270e57a2012d104524f8cecabc8222d92"></a>

## Direct properties — timeouts / 658babed0e04 / 3

<a id="canonical-a457eab41a1ac3b922bd953eed8b5f67187078ac18edde77ee5e2b7853268f4d"></a>

<a id="canonical-e6578a51e1b961e3d93fc2e12c9dffb602514f169d5de347f214ccebecc41bed"></a>

## create property — timeouts / 658babed0e04 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e68778ca0c3ed2e5d8f45e96c1ea0d1c219e83eae43ace3006cac4ddd525f07c"></a>

<a id="canonical-9de239a0a67ca131d5d232531573d36aa8671d21f7a7558e1f54a6fe76ea4871"></a>

## delete property — timeouts / 658babed0e04 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1e1535909b9c412e57afaf60b1472b865a61f3aa7431d16527a93940577f49a6"></a>

<a id="canonical-36b0b219eff7ca36e58ed0e4031203be8510f45a2cc91edba1e9a030938826da"></a>

## read property — timeouts / 658babed0e04 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-d14c588390847ac4898eed1d9d116b2c6a55d3072f6e4a1dc2d8e689cacb382d"></a>

<a id="canonical-d9231815561c396412c02ff6a4ac6f2ce3badd4eb52072440edcf8da530c2ec9"></a>

## update property — timeouts / 658babed0e04 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f763fe21ade38f60e654649cc3b40116dc985126bb3245a8c783c5567fba3a84"></a>

## Next pages — timeouts / 658babed0e04 / 8

- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0debae215b3c4fc4c6dda2dd02f05cf0a752629e49c17285897ce53e06ecc481)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
