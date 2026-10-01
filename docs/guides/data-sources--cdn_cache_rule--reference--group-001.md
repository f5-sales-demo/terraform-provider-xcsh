---
page_title: "xcsh_cdn_cache_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule reference."
---

# xcsh_cdn_cache_rule reference

<a id="canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1977ff3301300d5007062e5288bc317cce854912ba381a1abb680e5353590b6f"></a>

## Property reference — Property reference / b4a4edc533d9 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- Property reference

<a id="canonical-1c3f7aa8294c7aec445f551d7d370439a679bc3a36154a334844dd6fe83f72f6"></a>

## Direct properties — Property reference / b4a4edc533d9 / 3

<a id="canonical-9bed956e3cb87f05c1f13063d2da4cd3853d3ef40fcac542b40e49274b4cc473"></a>

<a id="canonical-85dfe3d6d5c10eedf683047385a306fa47aa737b3782a9ca3df651cc45178973"></a>

## annotations property — Property reference / b4a4edc533d9 / 4

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

- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377): complete subsection reference.

<a id="canonical-75421786553b9b275751b18f4931fb4927a0be1b813aeb5bf600525b24a48969"></a>

<a id="canonical-8a02af432113201f4841ba7632e8de030d215c73e9936955a8fe41240589f67d"></a>

## description property — Property reference / b4a4edc533d9 / 5

Type: `"string"`. Computed.

Description of the CDNCacheRule.

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

<a id="canonical-d7848b50d7c11d384584281930a77757e8891d2160245bfb433a55e2a45627d3"></a>

<a id="canonical-41119d2b4375430d9c252ccf16aac264d082ff448efcd4d49105329610b22b00"></a>

## id property — Property reference / b4a4edc533d9 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-fee6dc75034483ec65872741df24deed30dfeead02112e763a4c94f291a28a41"></a>

<a id="canonical-95784fa43f2cac0b5771bf60dea87c5602627e49898147972acf0e622aa93f72"></a>

## labels property — Property reference / b4a4edc533d9 / 7

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

<a id="canonical-9cab5bc90ccedb554101581f4770dc2b116ba985ae87e5505093304bd5923aaf"></a>

<a id="canonical-4f339b795d8109af648c083983bc7ef4011447e71d2d61e5cda46331f9f6d6b5"></a>

## name property — Property reference / b4a4edc533d9 / 8

Type: `"string"`. Required.

Name of the CDNCacheRule.

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

<a id="canonical-112e24e1947f0510f1cd9ee22962a11da0a20803d6d077a1ef4dcc10b35dabdd"></a>

<a id="canonical-bcf169ed8e30db6e86ed16bd47698346127f4702f181bf00e8c8cd33fadc3f7c"></a>

## namespace property — Property reference / b4a4edc533d9 / 9

Type: `"string"`. Required.

Namespace where the CDNCacheRule exists.

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

<a id="canonical-455ce666a007d4453a2ec29b932c6b792174cacefe1e25229d135c2304c880b5"></a>

## All schema paths — Property reference / b4a4edc533d9 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cdn_cache_rule--reference--group-001.md#canonical-9bed956e3cb87f05c1f13063d2da4cd3853d3ef40fcac542b40e49274b4cc473) |
| `cache_rules` | [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-994fce3e07b292af0d9d720b6e4a24d1f7c0ccffb72c6520b3aca534f85f238e) |
| `cache_rules.cache_bypass` | [cache_rules.cache_bypass](data-sources--cdn_cache_rule--reference--group-001.md#canonical-d2bc44a1ec34dcb15e20b4e7bc8d26df375ab2560c3b6d90cf5dc4483b2a5c14) |
| `cache_rules.eligible_for_cache` | [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-68986a34b10d5e6dafc64408bfde65bc351804d5cdd4419240edbe71e57f3f7e) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-cfb0d23494853244d1d12c1b886c61926b0eeb336169f236fa7f6e576121b77b) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override](data-sources--cdn_cache_rule--reference--group-001.md#canonical-af699daf7bd653abcfa8bd54fa8b3d17ef5790abf8f015b841dd6dfcbbe650ee) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl](data-sources--cdn_cache_rule--reference--group-001.md#canonical-130f2178732313c635cbae63612ddadd87016adfc2812f48af3e8fa500401443) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0923a2ffd43d597c31ef297386ec270ff90d15538bf4e16be81b4dc3fe638589) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-7bf6bb131e5de1577fdf7c4409010abdd88d8a7108df7db0b172006193a8ac7d) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override](data-sources--cdn_cache_rule--reference--group-001.md#canonical-fde0d001ac0adc18fd9afd5d9e204e0620360fffee857361d52b2f35259080ae) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl](data-sources--cdn_cache_rule--reference--group-001.md#canonical-119686ab7e189017b4bc89d47499fe2cec11d5d4880ff0a1d6b092a0fba37333) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3fce2accd643a7c79620a14c2c467ea7f75f1d945b9d637e8c97c2849e94133a) |
| `cache_rules.rule_expression_list` | [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-09166e07b8ad119ed3e48cd5d09873d7b671e170eaea24c2f7f3d6f4b5b4e44e) |
| `cache_rules.rule_expression_list.cache_rule_expression` | [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-78f69ab06d01c94bf9b448ef9e5fa5c3e664204cf74396ea6f7d628ca60f66e2) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-cd08c3322e3c0a67766466bc27a329937f82228ccb58156247bc26be3600195b) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-d08c140859739d40a0cdb6298b519b737786a3cbf0fd641c8a6a1360ff4e447a) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-635d736a2278d72602909e7c7c075755b85224a53467a0910e6125662b4c5751) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-4218ace1aa6bdb7aa3f56cd7d15d4bb1fffd7bf5c39c748f57f72f7fbaa54c16) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-7715892860acbfe1eb2c4636c1524006848c846e09f70c259d6f2b3637ee5433) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-4802029a393c810d02773264b6470d892357905c37f9c644affd58bd50ff09f5) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-c07315d61d739be183fd761eeaf48c0fdcba34ec4fddbf239bcb9b53f70a6fcd) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-7098cace312476275cc56c6668fda6840e9dee631b9fa44f93efaf243533c124) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-ff9b7e2d85fe74bba9c90b656d71fc1fbdc1a3e9e0bbbcc5047ad7495610c3cd) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6a5d746b75c6baa7b13d0b6bb5e30927821fe0c06335ddaac5033b209f2653eb) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-338c72d5461ea7cec84292f1e3ca664b15a8f6782f7beb11fbafc7e493ab1c30) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-ccb69688e0c748dbc7a800ab3066f4563ff1b80409c6f4d251a969b375e326e9) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0ba20526102b86bd6e169b9141300eede789fc32292bbc3ab164913041e231d9) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3c866a49e45bad3b276812d623bf016b2ac64c16ed2c5468e32015b5a3ce0a29) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3f97f73855a98f3b03672c4943568dc0de12d0d07dfe217f3b5d83abc1b1e0c9) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-4d935730f15e65fb938b1dba2b1c721411b41962beca36b36749ec9c0bbb2a9c) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-49e43b2229757ed961aa32c901a6d37d13726c8ace67bfec9819bf996ae77e43) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-065946fdfde1af56d80c57586673a6d95f022e9f7ac35d22010638f0d7bf5687) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-d9a6b94d9d9c19a9a0c22698d4782f6e252281148d0658181dfc98e440f72f9f) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-564e85793b90163146f07226817da5fe2d2b602c38b40feaabb529d7f1e10664) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1685163865f7f676e2aef3b38addf436108ec7ada0b8e362d0d7d32c466d1be8) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b9cf344b4b5e5046e4933a211c0baf323ed306c13d53b652e5ef7d03861c59d6) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b133279fa97cde483d0ec140955752954d476d198cc21bd57bee2e118ac19623) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-224b1e56f106c17e7180b1206478ca73b9af49fba740fd5b613589d1e16fad0e) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match` | [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1920f10bdb1e65d4fab4441d7bf6d37889be50e7144088c6c9d15750cd19fb6f) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-993a3d8b3ca950cce9d9b56e55d7741cf125e1b9276226f37e828f50534ebafa) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-27d74ccf3bc1461209a52ecc81d7f19f1b1269976a04f7f7c02b5f13df750be9) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-cbe05a949492ff5fb865c1baa99b255ef8ae22b993fb3f19ebc4047bc7674d1c) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-84e883c57d5aa3ed5f175b0053d3cca3b93f2f9c9ddec01ceb6ba90856fa58c4) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-123d3946151a502c82c60ca237d849fa4770a732e09615f3fb469cc4152ff6a0) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-326ea478869cdac211858efeeafd661938a2e5327188b3e7dc2199acb9479e40) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e28c40a8169a31a8b45e80fb493c2cac81c09ba8b6c53df5c64b1aac9ef6df73) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-a9301afd3361e66b5fdb956232166d5308a1f8799f1ccfece23f0f1aaf2cc561) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-f85b9a5dce3a91c2a4feb112f3e8410c6ff0627d96c066b06effe3d595ceaa62) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-4a872adfbafe9b7227d19eb0681b15a96fbb0d829f294c575fad919acdcee3e2) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-9e42e6d8fb2099ed45c1e7d79933f1853f7918557734d40fc55ef8489c7478d5) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1e609d2e261e9f41de5e393bfd66257be5fb8cdc8809c01fd547afd5eb29dab6) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1d0b8ed7e83eafc5a13c84753b68f9ae85a16dd7c401cff7eb49e02c5b8dd3c4) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-fab713a64873526085b40e268588d378e8af749411974654f3add980983a9956) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-237aeb2c038fcd23f014746dcc7f1609d210e2e3a261250330d1dce89f71f461) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b191674420d5db99e3cd9b8ec80093ce43171f0d61950498a5701d1692c00106) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-ccb314fcc76d1130206324c9286851444c2813561681724027488b595eaec3e4) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-7cf3c2dc978ae88999ca4b53505d6d2d2672cfa6355b28d7e09c527ba7946d8c) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5dbb7702542321adc4915cd6af0c0b02c368c8aa5760a35a2afb131fd2a5360f) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b51de005b5c321bf406f16becef6b28bf7d6b82cceda40186a65a1793adb7fef) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-9a90f41a2378f0f06c7d7ae61dccbd010df5a63d8056f5dfd30be872bd766c7a) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-11812bf99e90fb2088ba4cc0643ab98abd9f70d06de425294c8b948d8969321a) |
| `cache_rules.rule_expression_list.expression_name` | [cache_rules.rule_expression_list.expression_name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-f48c7aa4ca5a2816bdef685483c8763c60a76cff8cbe11c4812e31908e363278) |
| `cache_rules.rule_name` | [cache_rules.rule_name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5850c9974d3a91adb6a3c0da0a06ca4575c7fcbb2860f38aca1e5a174ebf0975) |
| `description` | [description](data-sources--cdn_cache_rule--reference--group-001.md#canonical-75421786553b9b275751b18f4931fb4927a0be1b813aeb5bf600525b24a48969) |
| `id` | [id](data-sources--cdn_cache_rule--reference--group-001.md#canonical-d7848b50d7c11d384584281930a77757e8891d2160245bfb433a55e2a45627d3) |
| `labels` | [labels](data-sources--cdn_cache_rule--reference--group-001.md#canonical-fee6dc75034483ec65872741df24deed30dfeead02112e763a4c94f291a28a41) |
| `name` | [name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-9cab5bc90ccedb554101581f4770dc2b116ba985ae87e5505093304bd5923aaf) |
| `namespace` | [namespace](data-sources--cdn_cache_rule--reference--group-001.md#canonical-112e24e1947f0510f1cd9ee22962a11da0a20803d6d077a1ef4dcc10b35dabdd) |

<a id="canonical-5bf1e6a9365eaa778185949f3eb68602480c383ce4722d65908cc200f76fb302"></a>

## Next pages — Property reference / b4a4edc533d9 / 11

- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89ff049c4f0226038e470d1da80dae9f880af55081c5b74571f8787c2b1f1a61"></a>

## cache_rules — cache_rules / 4e488c466d0f / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- cache_rules

<a id="canonical-994fce3e07b292af0d9d720b6e4a24d1f7c0ccffb72c6520b3aca534f85f238e"></a>

Type: `"single"`. Computed.

Cache Rule. This defines a CDN Cache Rule.

Upstream description:

This defines a CDN Cache Rule.

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

<a id="canonical-2ffa164e0f490f371e9d903b57b90f030206725235d700c5ed4c97129aad0e65"></a>

## Direct properties — cache_rules / 4e488c466d0f / 3

- [cache_bypass](data-sources--cdn_cache_rule--reference--group-001.md#canonical-60d218f139f95f3519c91374265f53e17878d3f639ce27a7490ef2952067a724): complete subsection reference.

- [eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-19aa73572255739cc4312784d53bdbd7c70ace784e8a517880e7de416afb68db): complete subsection reference.

- [rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0): complete subsection reference.

<a id="canonical-5850c9974d3a91adb6a3c0da0a06ca4575c7fcbb2860f38aca1e5a174ebf0975"></a>

<a id="canonical-8ebf54f4467e3799460cb5be36d348320e1c999039f7e006125844304573ce67"></a>

## rule_name property — cache_rules / 4e488c466d0f / 4

Type: `"string"`. Computed.

Rule Name. Name of the Cache Rule.

Upstream description:

Name of the Cache Rule.

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

<a id="canonical-54132b158acb8b1f07f21c3fe5b906069ee77dd9d2346fdb918dbfd3af459811"></a>

## Next pages — cache_rules / 4e488c466d0f / 5

- [cache_rules.cache_bypass](data-sources--cdn_cache_rule--reference--group-001.md#canonical-60d218f139f95f3519c91374265f53e17878d3f639ce27a7490ef2952067a724)
- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-19aa73572255739cc4312784d53bdbd7c70ace784e8a517880e7de416afb68db)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-60d218f139f95f3519c91374265f53e17878d3f639ce27a7490ef2952067a724"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-459d14c5fd2c2a822b08f7d29d9c8bc2f004978b7137289c401e6196dc8c9421"></a>

## cache_rules.cache_bypass — cache_rules.cache_bypass / 1a4c17d028df / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- cache_rules.cache_bypass

<a id="canonical-d2bc44a1ec34dcb15e20b4e7bc8d26df375ab2560c3b6d90cf5dc4483b2a5c14"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-707714976f324dc338ae0b7956d9e6e42a64c211583426fd73b288fcb76b02f7"></a>

## Direct properties — cache_rules.cache_bypass / 1a4c17d028df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9cc3b5fdad1a96853bb15b783cb1bd08c9c9ee6ead74406e3c136c0795b5347"></a>

## Next pages — cache_rules.cache_bypass / 1a4c17d028df / 4

- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-19aa73572255739cc4312784d53bdbd7c70ace784e8a517880e7de416afb68db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e8971c20915614c14a85d4635a4985117772fc600b6b991a6febda797f3c248"></a>

## cache_rules.eligible_for_cache — cache_rules.eligible_for_cache / c4bab1aa467f / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- cache_rules.eligible_for_cache

<a id="canonical-68986a34b10d5e6dafc64408bfde65bc351804d5cdd4419240edbe71e57f3f7e"></a>

Type: `"single"`. Computed.

Configuration parameter for eligible for cache.

Upstream description:

List of OPTIONS for Cache Action.

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

<a id="canonical-7e8f2faff365cd8e2a39607ad7124f0c5bc6ef1004f78d61f70258f113063401"></a>

## Direct properties — cache_rules.eligible_for_cache / c4bab1aa467f / 3

- [scheme_proxy_host_request_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5f654fe992c5daaa51fd1b7e40fe749000b767cc7045e3fa44c6d056fd39f972): complete subsection reference.

- [scheme_proxy_host_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-baeecf0e4745c8fa82d50896ecfe9f01c54e0e007e5ffa728705b004a475f477): complete subsection reference.

<a id="canonical-8880f365ed9f49193b54bb72333dbcba74d157f489b7aa36ef878cb6f674ab36"></a>

## Next pages — cache_rules.eligible_for_cache / c4bab1aa467f / 4

- [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5f654fe992c5daaa51fd1b7e40fe749000b767cc7045e3fa44c6d056fd39f972)
- [cache_rules.eligible_for_cache.scheme_proxy_host_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-baeecf0e4745c8fa82d50896ecfe9f01c54e0e007e5ffa728705b004a475f477)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-5f654fe992c5daaa51fd1b7e40fe749000b767cc7045e3fa44c6d056fd39f972"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69b3bfa87d87b669392003e2ce70a90c29dcf49f9184c3c74b85c4302682afe9"></a>

## cache_rules.eligible_for_cache.scheme_proxy_host_request_uri — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / e1e399ede098 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-19aa73572255739cc4312784d53bdbd7c70ace784e8a517880e7de416afb68db)
- cache_rules.eligible_for_cache.scheme_proxy_host_request_uri

<a id="canonical-cfb0d23494853244d1d12c1b886c61926b0eeb336169f236fa7f6e576121b77b"></a>

Type: `"single"`. Computed.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

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

<a id="canonical-1672f1e018deec6ea49210834902c273992f65f388a793d45089b39699e4d4f9"></a>

## Direct properties — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / e1e399ede098 / 3

<a id="canonical-af699daf7bd653abcfa8bd54fa8b3d17ef5790abf8f015b841dd6dfcbbe650ee"></a>

<a id="canonical-a9b11e4bd33c227b062da74d93075a42385dd261b3b810492923232f5d96b004"></a>

## cache_override property — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / e1e399ede098 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-130f2178732313c635cbae63612ddadd87016adfc2812f48af3e8fa500401443"></a>

<a id="canonical-47f3a38c73bcd687703e57ae6b02f64981abfcdad16256bf3124c7afca0646c9"></a>

## cache_ttl property — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / e1e399ede098 / 5

Type: `"string"`. Computed.

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

<a id="canonical-0923a2ffd43d597c31ef297386ec270ff90d15538bf4e16be81b4dc3fe638589"></a>

<a id="canonical-faecf5f62a4e62daa5bd321c6cced9e69205ac53b54db156cdfd54f58f8c0f27"></a>

## ignore_response_cookie property — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / e1e399ede098 / 6

Type: `"bool"`. Computed.

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

<a id="canonical-6ea16cd8f4271eb1bb270428625aad955d7e4071ba913cc6628b1aef8baf36f3"></a>

## Next pages — cache_rules.eligible_for_cache.scheme_proxy_host_request_uri / e1e399ede098 / 7

- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-19aa73572255739cc4312784d53bdbd7c70ace784e8a517880e7de416afb68db)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-baeecf0e4745c8fa82d50896ecfe9f01c54e0e007e5ffa728705b004a475f477"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c50b76316199d642cab71c3b715bd764d92b5112dc393151ee1c33b609d5642"></a>

## cache_rules.eligible_for_cache.scheme_proxy_host_uri — cache_rules.eligible_for_cache.scheme_proxy_host_uri / e59a0d049f07 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-19aa73572255739cc4312784d53bdbd7c70ace784e8a517880e7de416afb68db)
- cache_rules.eligible_for_cache.scheme_proxy_host_uri

<a id="canonical-7bf6bb131e5de1577fdf7c4409010abdd88d8a7108df7db0b172006193a8ac7d"></a>

Type: `"single"`. Computed.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

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

<a id="canonical-fe05ca10a04a7f1d387866c1536938a078ed8c679257f65cfcad97fc3b198625"></a>

## Direct properties — cache_rules.eligible_for_cache.scheme_proxy_host_uri / e59a0d049f07 / 3

<a id="canonical-fde0d001ac0adc18fd9afd5d9e204e0620360fffee857361d52b2f35259080ae"></a>

<a id="canonical-64b15a3b3c31d40dfb7c76494da983e90c3d69775fa85e059acbd5d7f6007634"></a>

## cache_override property — cache_rules.eligible_for_cache.scheme_proxy_host_uri / e59a0d049f07 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-119686ab7e189017b4bc89d47499fe2cec11d5d4880ff0a1d6b092a0fba37333"></a>

<a id="canonical-f5aeed6c940917cec40f7962a278f8e3bfebcf4606ef7952d5c05f0175b5ad0d"></a>

## cache_ttl property — cache_rules.eligible_for_cache.scheme_proxy_host_uri / e59a0d049f07 / 5

Type: `"string"`. Computed.

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

<a id="canonical-3fce2accd643a7c79620a14c2c467ea7f75f1d945b9d637e8c97c2849e94133a"></a>

<a id="canonical-55ff078fa0439be6944072205ffc7fbc7190909a4798ac602ea113eaee184296"></a>

## ignore_response_cookie property — cache_rules.eligible_for_cache.scheme_proxy_host_uri / e59a0d049f07 / 6

Type: `"bool"`. Computed.

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

<a id="canonical-54596a0c1c0b6dd90b52b4496850c4407e9aecce1d37bac31ebe8d9877b214fb"></a>

## Next pages — cache_rules.eligible_for_cache.scheme_proxy_host_uri / e59a0d049f07 / 7

- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-19aa73572255739cc4312784d53bdbd7c70ace784e8a517880e7de416afb68db)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fca5d83e4b994f1a5cae13db3d668d242d48d4d460a0f5f0b5c28b3a0fe676e3"></a>

## cache_rules.rule_expression_list — cache_rules.rule_expression_list / 6a9ba7397ec2 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- cache_rules.rule_expression_list

<a id="canonical-09166e07b8ad119ed3e48cd5d09873d7b671e170eaea24c2f7f3d6f4b5b4e44e"></a>

Type: `"list"`. Computed.

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs.

Upstream description:

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs..

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

<a id="canonical-285d614c172cf21f99ab786860febb51c28c38a04d265d573c3c03ef94727a1f"></a>

## Direct properties — cache_rules.rule_expression_list / 6a9ba7397ec2 / 3

- [cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878): complete subsection reference.

<a id="canonical-f48c7aa4ca5a2816bdef685483c8763c60a76cff8cbe11c4812e31908e363278"></a>

<a id="canonical-4b6fd970c7f34f2dae141cd18e9d32dfa380691a13f8289bf27cf6b9c261001c"></a>

## expression_name property — cache_rules.rule_expression_list / 6a9ba7397ec2 / 4

Type: `"string"`. Computed.

Name of the Expressions items that are ANDed.

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

<a id="canonical-12a0ea124bf57cece8fd9c82ae07468d4d4c8cbac9ebfa42e4238e9e0ce4aa4b"></a>

## Next pages — cache_rules.rule_expression_list / 6a9ba7397ec2 / 5

- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a94656adbad30bfe114fdbede23fae5b71cc218dc6ee3f75a202f28985baeef0"></a>

## cache_rules.rule_expression_list.cache_rule_expression — cache_rules.rule_expression_list.cache_rule_expression / 3a3ecf4fad7c / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- cache_rules.rule_expression_list.cache_rule_expression

<a id="canonical-78f69ab06d01c94bf9b448ef9e5fa5c3e664204cf74396ea6f7d628ca60f66e2"></a>

Type: `"list"`. Computed.

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

<a id="canonical-5845b37addbd6cf8cee17e1ed166945457e9050823391aabc99c1b2766d679d1"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression / 3a3ecf4fad7c / 3

- [cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b44b9ecf649bbefe629636ac897d75a2430d93dbd65ac84ec5445b2d05d68b33): complete subsection reference.

- [cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6da0ba1942373bcf16efe7bbc24ec712daf2b01130424df6617aff12379bb504): complete subsection reference.

- [path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6a7b38d27e0b9077d1029effe81d4cc0c79e91d5c4e897ae7aa6728053c25c81): complete subsection reference.

- [query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-ecc9368bcc358b04d1acfd7ccf9504fe886fee0daacf2decc3116833465ebe56): complete subsection reference.

<a id="canonical-83e93e05e841c4d9ed64e09d86deb4f4858c1371f59c2fdd25d69152ec07dd00"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression / 3a3ecf4fad7c / 4

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b44b9ecf649bbefe629636ac897d75a2430d93dbd65ac84ec5445b2d05d68b33)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6da0ba1942373bcf16efe7bbc24ec712daf2b01130424df6617aff12379bb504)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6a7b38d27e0b9077d1029effe81d4cc0c79e91d5c4e897ae7aa6728053c25c81)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-ecc9368bcc358b04d1acfd7ccf9504fe886fee0daacf2decc3116833465ebe56)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-b44b9ecf649bbefe629636ac897d75a2430d93dbd65ac84ec5445b2d05d68b33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6124f73bec3cf0c19f39c85aa49a10cf4be64c14fcd318786ec8e9c109d274d7"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cache_headers — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / bd2b5dab269a / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers

<a id="canonical-cd08c3322e3c0a67766466bc27a329937f82228ccb58156247bc26be3600195b"></a>

Type: `"list"`. Computed.

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

<a id="canonical-1b164488a1e1aa944bc8dafa1d842c51c30ffa6ed58fc8abb6e260d524184428"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / bd2b5dab269a / 3

<a id="canonical-d08c140859739d40a0cdb6298b519b737786a3cbf0fd641c8a6a1360ff4e447a"></a>

<a id="canonical-9d5a9165dff7b008f5460ef6d16c2a54a07338b1e9e7a790ff77b34c89866772"></a>

## name property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / bd2b5dab269a / 4

Type: `"string"`. Computed.

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

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-209ffde0d1eed62cc6411b45b4a5c84263058300ae5714e5695d9999646217fa): complete subsection reference.

<a id="canonical-70c051f8a8a82950da9c4e174901cc3cbcc2cad4949fbbd1acc2323b20ee0d20"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cache_headers / bd2b5dab269a / 5

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-209ffde0d1eed62cc6411b45b4a5c84263058300ae5714e5695d9999646217fa)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-209ffde0d1eed62cc6411b45b4a5c84263058300ae5714e5695d9999646217fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb19e0e32be2f3aa0c30e1ec922e2028ca3b6bf0e4bfb99e9145804fa0b0124b"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b44b9ecf649bbefe629636ac897d75a2430d93dbd65ac84ec5445b2d05d68b33)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator

<a id="canonical-635d736a2278d72602909e7c7c075755b85224a53467a0910e6125662b4c5751"></a>

Type: `"single"`. Computed.

Operator

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

<a id="canonical-f3c2d2e892908251c03c98bf6fb1b86e0852f30620decb5dbad2e5517303d0c1"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 3

<a id="canonical-4218ace1aa6bdb7aa3f56cd7d15d4bb1fffd7bf5c39c748f57f72f7fbaa54c16"></a>

<a id="canonical-7d88964eb261f55ea08b138605ea33a65abc730ee637d51f0e747045078a1db8"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The header value must include the specified value as a substring.

<a id="canonical-7715892860acbfe1eb2c4636c1524006848c846e09f70c259d6f2b3637ee5433"></a>

<a id="canonical-8f46f7573155e7ef31a8ad0cf85fe4b6ed8d58774dffae69ff60746f4673fd69"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not include the specified value as a substring.

<a id="canonical-4802029a393c810d02773264b6470d892357905c37f9c644affd58bd50ff09f5"></a>

<a id="canonical-e63437fd65b17560fd772cdc3f57984b9e3f2632d7bcec991b9f5ba90ad35cd0"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not end with the specified value.

<a id="canonical-c07315d61d739be183fd761eeaf48c0fdcba34ec4fddbf239bcb9b53f70a6fcd"></a>

<a id="canonical-b23014d357a4fcf50f0de1e09e8705ab72bf7fba40f16afd1dbf4f6a6c9b0ed8"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not match the specified value.

<a id="canonical-7098cace312476275cc56c6668fda6840e9dee631b9fa44f93efaf243533c124"></a>

<a id="canonical-16f27353de508668c76238d9462fce87028e97625cc00af478ee31b58c954075"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The header value must not begin with the specified value.

<a id="canonical-ff9b7e2d85fe74bba9c90b656d71fc1fbdc1a3e9e0bbbcc5047ad7495610c3cd"></a>

<a id="canonical-b3ecfe4747d703e8bae52f0fa9e954eda21c3745ec07419a1b8d59f6d0781e05"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The header value must end with the specified value.

<a id="canonical-6a5d746b75c6baa7b13d0b6bb5e30927821fe0c06335ddaac5033b209f2653eb"></a>

<a id="canonical-e6861e29e8cf40623909d9de30f34db7d8270057a92a0647956be67dbe9b57bc"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The header value must exactly match the specified value.

<a id="canonical-338c72d5461ea7cec84292f1e3ca664b15a8f6782f7beb11fbafc7e493ab1c30"></a>

<a id="canonical-f6767b3d9073a63a50964507de6dadc4248102c53ddc3b936670e0c5b9485946"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The header value must match the specified regular expression pattern.

<a id="canonical-ccb69688e0c748dbc7a800ab3066f4563ff1b80409c6f4d251a969b375e326e9"></a>

<a id="canonical-da6cc3d4fb1b0de9600265ff561ed13916dfb783aa2fb1b5f367fd795ea04d7c"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The header value must begin with the specified value.

<a id="canonical-a70a6383c8c7374f059b18e21ce8afbfdfe70ee017484dc98764cce26a68bd5d"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator / 0a2ef4fa42dd / 13

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b44b9ecf649bbefe629636ac897d75a2430d93dbd65ac84ec5445b2d05d68b33)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-6da0ba1942373bcf16efe7bbc24ec712daf2b01130424df6617aff12379bb504"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0817c9504e87754df13f52ed38a7959e5c764cc1ecc0a0bed01cd23adb153907"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / a37fefca4b2b / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher

<a id="canonical-0ba20526102b86bd6e169b9141300eede789fc32292bbc3ab164913041e231d9"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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

<a id="canonical-2f08712ea81fb3cb29ccc02942f0b49d39d5404b3612f4266853d12528d5a0c2"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / a37fefca4b2b / 3

<a id="canonical-3c866a49e45bad3b276812d623bf016b2ac64c16ed2c5468e32015b5a3ce0a29"></a>

<a id="canonical-41d7f6b10c5ec100e4bb7236d1e278064f5342c160ef4f5bef217439f5d37ed9"></a>

## name property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / a37fefca4b2b / 4

Type: `"string"`. Computed.

Cookie Name. Enter the name of the cookie to match.

Upstream description:

Enter the name of the cookie to match.

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

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-454dd3f8548de10092986c8f6d1fca0b828e82536e6a3dfca44c45b3fb0edd00): complete subsection reference.

<a id="canonical-edd74a0d5d16aed769502cfeda8f5b6aad46e4d87e0d98e50ead634cc9908151"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher / a37fefca4b2b / 5

- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-454dd3f8548de10092986c8f6d1fca0b828e82536e6a3dfca44c45b3fb0edd00)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-454dd3f8548de10092986c8f6d1fca0b828e82536e6a3dfca44c45b3fb0edd00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26f7f0a530c20c6335890b343a00de4c22ef254978a4a9b2ddc4ac0054bfb1b7"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6da0ba1942373bcf16efe7bbc24ec712daf2b01130424df6617aff12379bb504)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator

<a id="canonical-3f97f73855a98f3b03672c4943568dc0de12d0d07dfe217f3b5d83abc1b1e0c9"></a>

Type: `"single"`. Computed.

Operator

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

<a id="canonical-2d44ff803b184d279b3fb46e0b0f1542ecc6f4854fe4a0aa57b22a9340434b13"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 3

<a id="canonical-4d935730f15e65fb938b1dba2b1c721411b41962beca36b36749ec9c0bbb2a9c"></a>

<a id="canonical-415f7fde293714a42820d2e8ae30abd87a3de7e8d4d24fae68d5dee917102a44"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The cookie value must include the specified value as a substring.

<a id="canonical-49e43b2229757ed961aa32c901a6d37d13726c8ace67bfec9819bf996ae77e43"></a>

<a id="canonical-87d7a0e607f41983177cde81e95e03bf9ca05b4514747a479a4474e08cf033fd"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not include the specified value as a substring.

<a id="canonical-065946fdfde1af56d80c57586673a6d95f022e9f7ac35d22010638f0d7bf5687"></a>

<a id="canonical-046282dc75cbc72d9d71c5799b957f1e48e48f22c6e32b4d58bf74155909ddab"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not end with the specified value.

<a id="canonical-d9a6b94d9d9c19a9a0c22698d4782f6e252281148d0658181dfc98e440f72f9f"></a>

<a id="canonical-1b6f7e299340c920d7646b8627b88df03783fa4612bfddb52dc93691cc6a9d42"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not match the specified value.

<a id="canonical-564e85793b90163146f07226817da5fe2d2b602c38b40feaabb529d7f1e10664"></a>

<a id="canonical-5277f43f0fa33e6ae6f90efa219fb9947d549d5ace2f4e0f72f257e802a2a5bf"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The cookie value must not begin with the specified value.

<a id="canonical-1685163865f7f676e2aef3b38addf436108ec7ada0b8e362d0d7d32c466d1be8"></a>

<a id="canonical-8b2d1d6fabf9abfb05e95ac9be78b10cd3643053baa532256653f326dbe589a0"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The cookie value must end with the specified value.

<a id="canonical-b9cf344b4b5e5046e4933a211c0baf323ed306c13d53b652e5ef7d03861c59d6"></a>

<a id="canonical-f1e53cbfffd099bb4fb6eb46e83c53c9a459a6dfb103c56047b6f1fd3027841f"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The cookie value must exactly match the specified value.

<a id="canonical-b133279fa97cde483d0ec140955752954d476d198cc21bd57bee2e118ac19623"></a>

<a id="canonical-6c7ada01cf04d2a11522968d7f908c41355da537b92e77783422838c66ac1b47"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The cookie value must match the specified regular expression pattern in PCRE
format.

<a id="canonical-224b1e56f106c17e7180b1206478ca73b9af49fba740fd5b613589d1e16fad0e"></a>

<a id="canonical-7e00f2cae902a96df226225b42cdb2aed6f0a1b507370682b44da578fcbab8b2"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The cookie value must begin with the specified value.

<a id="canonical-e43b676e9100ccb3c1a4515a25aee79c161bb962fae7db3717176c200c1ceb44"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator / b624ad09eef5 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6da0ba1942373bcf16efe7bbc24ec712daf2b01130424df6617aff12379bb504)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-6a7b38d27e0b9077d1029effe81d4cc0c79e91d5c4e897ae7aa6728053c25c81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07d15a7b2838d275c1a99584653ce92c1f0b989b7a250f079d8fccf2df4cdbaf"></a>

## cache_rules.rule_expression_list.cache_rule_expression.path_match — cache_rules.rule_expression_list.cache_rule_expression.path_match / cc114fdbd56f / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- cache_rules.rule_expression_list.cache_rule_expression.path_match

<a id="canonical-1920f10bdb1e65d4fab4441d7bf6d37889be50e7144088c6c9d15750cd19fb6f"></a>

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
  }
}
```

<a id="canonical-e647ee51f50feea9a7aaff2004c8fc4b8c459c7d4bb548235c4240a14487ee9d"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.path_match / cc114fdbd56f / 3

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-44b3439c84ccb76c1c5e7b712dfdd38facc1707b1bc1733adb101b2dfee10748): complete subsection reference.

<a id="canonical-fd384348421a9e850d03b2f0d2b64ab798dc37d8105566db98de0b8ade88268f"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.path_match / cc114fdbd56f / 4

- [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-44b3439c84ccb76c1c5e7b712dfdd38facc1707b1bc1733adb101b2dfee10748)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-44b3439c84ccb76c1c5e7b712dfdd38facc1707b1bc1733adb101b2dfee10748"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac695632b7958cb06d3528d260d9ec8a911e8766b8e8652958fa757fe1995369"></a>

## cache_rules.rule_expression_list.cache_rule_expression.path_match.operator — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6a7b38d27e0b9077d1029effe81d4cc0c79e91d5c4e897ae7aa6728053c25c81)
- cache_rules.rule_expression_list.cache_rule_expression.path_match.operator

<a id="canonical-993a3d8b3ca950cce9d9b56e55d7741cf125e1b9276226f37e828f50534ebafa"></a>

Type: `"single"`. Computed.

Operator

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

<a id="canonical-35375ed1558f0f7d4f5fa7910a3c6148d18343dba775d20a18c42ad5f1a7bfb7"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 3

<a id="canonical-27d74ccf3bc1461209a52ecc81d7f19f1b1269976a04f7f7c02b5f13df750be9"></a>

<a id="canonical-8f2f50cb140dd96073701c31f469864de4d402aa491c0d28d4fea3bd5591ae71"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The path must include the specified value as a substring, up to the
filename.

<a id="canonical-cbe05a949492ff5fb865c1baa99b255ef8ae22b993fb3f19ebc4047bc7674d1c"></a>

<a id="canonical-7ca5e380904dc47a48a018c4549f59c18277b42b29241ad1b744b8a54ec88048"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not include the specified value as a substring, up to the filename.

<a id="canonical-84e883c57d5aa3ed5f175b0053d3cca3b93f2f9c9ddec01ceb6ba90856fa58c4"></a>

<a id="canonical-3c9444139dfac168427f343c07711bf90516a7bc706f941b32d5126dfc7e85d4"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not end with the specified value, up to the filename.

<a id="canonical-123d3946151a502c82c60ca237d849fa4770a732e09615f3fb469cc4152ff6a0"></a>

<a id="canonical-8c05b359118dc65f792f29abd07a3ff6b33eff65351478f6b9590c44a06b1b57"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not match the specified value, up to the filename.

<a id="canonical-326ea478869cdac211858efeeafd661938a2e5327188b3e7dc2199acb9479e40"></a>

<a id="canonical-624165c5a6f69c257b9336340d7e841cd9d874ad489c2654378e93dcde52c7d5"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The path must not begin with the specified value, up to the filename.

<a id="canonical-e28c40a8169a31a8b45e80fb493c2cac81c09ba8b6c53df5c64b1aac9ef6df73"></a>

<a id="canonical-a38cc22150418376dc52761e28a6b9476000cc77e92d6b6fcc1c8b9f47a066ff"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The path must end with the specified value, up to the filename.

<a id="canonical-a9301afd3361e66b5fdb956232166d5308a1f8799f1ccfece23f0f1aaf2cc561"></a>

<a id="canonical-04cc7dc4a2b0a6e70bdbff72d14ebc3a95534ea0f566cb4c747cae74a83af22a"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The path must exactly match the specified value, up to the filename.

<a id="canonical-f85b9a5dce3a91c2a4feb112f3e8410c6ff0627d96c066b06effe3d595ceaa62"></a>

<a id="canonical-0d51d387a327accd78f7edc653006aa6759465de02a9bbdc74e80b63b45289e8"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The path must match the specified regular expression pattern in PCRE format.

<a id="canonical-4a872adfbafe9b7227d19eb0681b15a96fbb0d829f294c575fad919acdcee3e2"></a>

<a id="canonical-c5b01e73c6d9c7ce53836ed8b19b30b08d8f578342271cf5f7e12dfcb3368b52"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The path must begin with the specified value, up to the filename.

<a id="canonical-6d1df8458d8ad2da3ffe454643c70a070f48d1ec6c143ac2eafaa53a3d038f7e"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.path_match.operator / 9d6abce85517 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-6a7b38d27e0b9077d1029effe81d4cc0c79e91d5c4e897ae7aa6728053c25c81)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-ecc9368bcc358b04d1acfd7ccf9504fe886fee0daacf2decc3116833465ebe56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b5bb87bc1035bd621fbc6b7685a153db84c243bbc098b2822dded9c590129d4"></a>

## cache_rules.rule_expression_list.cache_rule_expression.query_parameters — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / d25868b230c0 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters

<a id="canonical-9e42e6d8fb2099ed45c1e7d79933f1853f7918557734d40fc55ef8489c7478d5"></a>

Type: `"list"`. Computed.

Query Parameters. List of (key, value) query parameters.

Upstream description:

List of (key, value) query parameters.

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

<a id="canonical-5d74910278931fe8b6c7f5f601163cf12c5f53cc31bc319774d8736ac51bf8f4"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / d25868b230c0 / 3

<a id="canonical-1e609d2e261e9f41de5e393bfd66257be5fb8cdc8809c01fd547afd5eb29dab6"></a>

<a id="canonical-ca6fee4974cfb63f1199f1d97af53a4fb648b7b033fc8add2f5542f8b5ab7981"></a>

## key property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / d25868b230c0 / 4

Type: `"string"`. Computed.

The name of the query parameter to match.

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

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-397cb46bffafca180d70b1536447fbffc71f50f43002c942e787d9705c488e17): complete subsection reference.

<a id="canonical-159d419136998b534d1ac047fde9b3cd16d43bc3b11cff3a344ef42ba84148e5"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.query_parameters / d25868b230c0 / 5

- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-397cb46bffafca180d70b1536447fbffc71f50f43002c942e787d9705c488e17)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-397cb46bffafca180d70b1536447fbffc71f50f43002c942e787d9705c488e17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-382767d27d0e01e1de7d4224297956fdaf687f57c2284505b57279472287ca0a"></a>

## cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-b8dd202c6c3d15a64786c181af5b951bb789335c4a000437e6946559b2cf93d9)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-eec090ceb54d3cf3a6b3e295c2135beff14a58ee47a6a32547052a09d38e6377)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-5bf474b4afb8d0cbbe83dec5ba601e74213f4d6d5ed71959fc6928435ae7c3b0)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-e74a8ae28e57b1495761777e5d158ef82e0538a9b47b0c2aa09cc29bc3e87878)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-ecc9368bcc358b04d1acfd7ccf9504fe886fee0daacf2decc3116833465ebe56)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator

<a id="canonical-1d0b8ed7e83eafc5a13c84753b68f9ae85a16dd7c401cff7eb49e02c5b8dd3c4"></a>

Type: `"single"`. Computed.

Operator

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

<a id="canonical-3bb7d8ee79430b8f73105ba1e677c66110aee8b7a45abc3ef3e78ac99db6bbce"></a>

## Direct properties — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 3

<a id="canonical-fab713a64873526085b40e268588d378e8af749411974654f3add980983a9956"></a>

<a id="canonical-2461a08e20912e00d2b65e1f86c944ff9f6c673857d73cdad39fd6bd6570cf80"></a>

## contains property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The query parameter value must include the specified value as a substring.

<a id="canonical-237aeb2c038fcd23f014746dcc7f1609d210e2e3a261250330d1dce89f71f461"></a>

<a id="canonical-a7abeaa19f90c895fffca71d5ba3af763488d04bcdec0674f448038a82787184"></a>

## does_not_contain property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not include the specified value as a substring.

<a id="canonical-b191674420d5db99e3cd9b8ec80093ce43171f0d61950498a5701d1692c00106"></a>

<a id="canonical-d9a114b843cad7519012a7777570c34a54e833c3136ba2077cfbdfe9eeeb2f54"></a>

## does_not_end_with property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not end with the specified value.

<a id="canonical-ccb314fcc76d1130206324c9286851444c2813561681724027488b595eaec3e4"></a>

<a id="canonical-63f5dfd578f2378ebbc91be2433190b3bf031a78fb54da21c492cf87c9b9f8ca"></a>

## does_not_equal property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not match the specified value.

<a id="canonical-7cf3c2dc978ae88999ca4b53505d6d2d2672cfa6355b28d7e09c527ba7946d8c"></a>

<a id="canonical-60194345b2dd1b6506c82f322eec91468e9226aaa0c4f11db7cfe7e05fec1ae3"></a>

## does_not_start_with property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The query parameter value must not begin with the specified value.

<a id="canonical-5dbb7702542321adc4915cd6af0c0b02c368c8aa5760a35a2afb131fd2a5360f"></a>

<a id="canonical-da7ad370c7931588c9d635e359ab3f461b0cf891f13ddcb9403e8b2bdd862e0e"></a>

## endswith property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The query parameter value must end with the specified value.

<a id="canonical-b51de005b5c321bf406f16becef6b28bf7d6b82cceda40186a65a1793adb7fef"></a>

<a id="canonical-b10618c8385cb2710f59f06e6640463cb39e3843b8611ae8c44fe339b335a6ae"></a>

## equals property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The query parameter value must exactly match the specified value.

<a id="canonical-9a90f41a2378f0f06c7d7ae61dccbd010df5a63d8056f5dfd30be872bd766c7a"></a>

<a id="canonical-c5ed9b9cdf7ebbf85871d238cb8744c0042785ce103fc4cb43d624757ccd4aee"></a>

## match_regex property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The query parameter value must match the specified regular expression pattern in
PCRE format.

<a id="canonical-11812bf99e90fb2088ba4cc0643ab98abd9f70d06de425294c8b948d8969321a"></a>

<a id="canonical-8ef9490fc5a44a5bd8dc7513da831cae6ed67220712e56ee8fa5a872727bcabe"></a>

## startswith property — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The query parameter value must begin with the specified value.

<a id="canonical-e77df739170bae8679e3cb22521f295eacfaae9b517e0c8deffc173303f4f442"></a>

## Next pages — cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator / cc8b1d49aa5d / 13

- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-ecc9368bcc358b04d1acfd7ccf9504fe886fee0daacf2decc3116833465ebe56)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
