---
page_title: "xcsh_rate_limiter_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy reference."
---

# xcsh_rate_limiter_policy reference

<a id="canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffbc18f930de3ae70a3bddd3e12049764baac0bf27563335f5f3ccc031d29847"></a>

## Property reference — Property reference / 32b78a761a0e / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- Property reference

<a id="canonical-e9f8099b999216003c534d76450f0d190c2cf38dea6ce255e01dc0ee14fc7fee"></a>

## Direct properties — Property reference / 32b78a761a0e / 3

<a id="canonical-b7243dc70fbabccdbcb7c16046eb349fdb9289a38733ebf7d7aa615122b64b1a"></a>

<a id="canonical-e65855e1fff1d69ad09f4403b183d876c714457902c033ab7feadef3fdc579f6"></a>

## annotations property — Property reference / 32b78a761a0e / 4

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

- [any_server](resources--rate_limiter_policy--reference--group-001.md#canonical-259eae0b31626643a9d6ca160de383c88e28187307228398093a6b5c43ff6439): complete subsection reference.

<a id="canonical-4d1b7877be1ac239f23165b69875631f6129e6db76adf8cafc979830f230c49a"></a>

<a id="canonical-b26fa2d94653f37e5a16400abded73523588ce1424dc19aafac7657bba1bedf1"></a>

## description property — Property reference / 32b78a761a0e / 5

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

<a id="canonical-3a25c25f05ecd086b61e2b39e26387984df44d219128cf35ecac6aedce8ffeb8"></a>

<a id="canonical-12bdcd15e5830f69146bdf4db01a3a895d32868a1d02fac203843a5a22a23025"></a>

## disable property — Property reference / 32b78a761a0e / 6

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

<a id="canonical-528a81a34ff4abe6f2b3d27b0c24420b9ee02471296b9725b27efca59b44b7f0"></a>

<a id="canonical-dc78bd9966c8b73d79dd4d28514ad7dfbd6a7a1bfc2882a6cf5d153663f428e0"></a>

## id property — Property reference / 32b78a761a0e / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d08c161b4234134e31793ec106abb85818748825d2c88640780b84984f0864ef"></a>

<a id="canonical-3820549e50ab43062b47ad02c60e1d7bb7c925733da27922cdd8aba07985e544"></a>

## labels property — Property reference / 32b78a761a0e / 8

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

<a id="canonical-5ee11f5764c14be4bd1530f74a5f0fbc3535e81155cf864c5c92447dd6c68a66"></a>

<a id="canonical-2dd965b4ec2730f314b56757ddb22d21558b85f7b61475bd475a518321a42c4e"></a>

## name property — Property reference / 32b78a761a0e / 9

Type: `"string"`. Required.

Name of the Rate Limiter Policy. Must be unique within the namespace.

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

<a id="canonical-744a277bf3067c58f1d393c0418c8f512909af7396d5f21619c4a433abfc910f"></a>

<a id="canonical-4735333694b21fcf6da5afac857eb7ba639f94d0dee351c1f26a0dcf508ba4c3"></a>

## namespace property — Property reference / 32b78a761a0e / 10

Type: `"string"`. Required.

Namespace where the Rate Limiter Policy is created.

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

- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3): complete subsection reference.

<a id="canonical-105c5085315080237edf5a6df9cafbf6959e9c476b7c8611cee03bc87fd34fc6"></a>

<a id="canonical-6e29c36e90e1a6250cdd4e095775eab2d358bbd4d2e6cec63b6ba589fce7bab8"></a>

## server_name property — Property reference / 32b78a761a0e / 11

Type: `"string"`. Optional, Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

Upstream description:

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3e759f0ab8dc83a52fb91efa326ee68914c628ce86919f316e19a1c4368373a3): complete subsection reference.

- [server_selector](resources--rate_limiter_policy--reference--group-001.md#canonical-a76ecf72963cf5b3f31077a4892876c381eb73eebf600858b6dd88334e80a6c7): complete subsection reference.

- [timeouts](resources--rate_limiter_policy--reference--group-001.md#canonical-596b33ac3e31c789c6ff45eb3514ad49a37cf7f206057468810e3be146744ecf): complete subsection reference.

<a id="canonical-827de3540a63ff28b8c5ce06d350ca34d277b4f98d177933981eef9b7b4f928d"></a>

## All schema paths — Property reference / 32b78a761a0e / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--rate_limiter_policy--reference--group-001.md#canonical-b7243dc70fbabccdbcb7c16046eb349fdb9289a38733ebf7d7aa615122b64b1a) |
| `any_server` | [any_server](resources--rate_limiter_policy--reference--group-001.md#canonical-2777f1c8cd4230c9d6cd44d7996500aecadd2474f006fd9c13a28e0c52f9b947) |
| `description` | [description](resources--rate_limiter_policy--reference--group-001.md#canonical-4d1b7877be1ac239f23165b69875631f6129e6db76adf8cafc979830f230c49a) |
| `disable` | [disable](resources--rate_limiter_policy--reference--group-001.md#canonical-3a25c25f05ecd086b61e2b39e26387984df44d219128cf35ecac6aedce8ffeb8) |
| `id` | [id](resources--rate_limiter_policy--reference--group-001.md#canonical-528a81a34ff4abe6f2b3d27b0c24420b9ee02471296b9725b27efca59b44b7f0) |
| `labels` | [labels](resources--rate_limiter_policy--reference--group-001.md#canonical-d08c161b4234134e31793ec106abb85818748825d2c88640780b84984f0864ef) |
| `name` | [name](resources--rate_limiter_policy--reference--group-001.md#canonical-5ee11f5764c14be4bd1530f74a5f0fbc3535e81155cf864c5c92447dd6c68a66) |
| `namespace` | [namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-744a277bf3067c58f1d393c0418c8f512909af7396d5f21619c4a433abfc910f) |
| `rules` | [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-a666a5c6864db8036f73005a2e011f3de7b9d71bef9b4afe5643f44e971e06cf) |
| `rules.metadata` | [rules.metadata](resources--rate_limiter_policy--reference--group-001.md#canonical-ce6aa3058b8d07a250989801e13c2d2b495a02d63a5d61307efa40a464ca392c) |
| `rules.metadata.description_spec` | [rules.metadata.description_spec](resources--rate_limiter_policy--reference--group-001.md#canonical-55642899b2f7d41d1e55ed64659b6532171b012fdeaddecd88e4f7f403c8a1c6) |
| `rules.metadata.name` | [rules.metadata.name](resources--rate_limiter_policy--reference--group-001.md#canonical-a9f45e916d72131a8d4eb4689adf11c46b95cf2724d868756047d31ad7fc330b) |
| `rules.spec` | [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-9f1aef64cd75db85245caad89c90efa36c3eaadf1e11e8ae9265721b03bad7a1) |
| `rules.spec.any_asn` | [rules.spec.any_asn](resources--rate_limiter_policy--reference--group-001.md#canonical-4733ad3b5dc4f5afe21afd45d5a3855006cdb5ccfd71f9ca3ca5670276380a9b) |
| `rules.spec.any_country` | [rules.spec.any_country](resources--rate_limiter_policy--reference--group-001.md#canonical-c79c15bcd26aef1495ecef47840a27f4d2454c4460257b4f51a6f11e05c7fda6) |
| `rules.spec.any_ip` | [rules.spec.any_ip](resources--rate_limiter_policy--reference--group-001.md#canonical-c2d835d0ddccae21fe28c37ade04017f8ed9eec4d89e73d9e631825eec175b1a) |
| `rules.spec.apply_rate_limiter` | [rules.spec.apply_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-dea393d208d5da742a7cb3f68d6b2701839da7e19bc98115ec68c709ca246d32) |
| `rules.spec.asn_list` | [rules.spec.asn_list](resources--rate_limiter_policy--reference--group-001.md#canonical-e067111f75e87f22fb33bed9da886ad51ff351bbaaf7e5272c25ca90883fde92) |
| `rules.spec.asn_list.as_numbers` | [rules.spec.asn_list.as_numbers](resources--rate_limiter_policy--reference--group-001.md#canonical-998f95aa6ffcbb30ad25b284f1c5df079f22f9649e230ea3e43dae064baa8dab) |
| `rules.spec.asn_matcher` | [rules.spec.asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-485e4cabf3bcfa154eb79cde868d438933262be6238d62300972f6ca4b5d6032) |
| `rules.spec.asn_matcher.asn_sets` | [rules.spec.asn_matcher.asn_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-0727ef6b81a4bd2df78e7f2140f6dd5ec99c1df10034b83bacaa2bab98d01a7d) |
| `rules.spec.asn_matcher.asn_sets.kind` | [rules.spec.asn_matcher.asn_sets.kind](resources--rate_limiter_policy--reference--group-001.md#canonical-c76b03864d4ae9c81a218256cb2abfa387a74ba9e6c1d99bbe025cca66d96349) |
| `rules.spec.asn_matcher.asn_sets.name` | [rules.spec.asn_matcher.asn_sets.name](resources--rate_limiter_policy--reference--group-001.md#canonical-87ccada235a7ccd8ae9bb07c7f2e545b3b030026eddf24e63143dfde60623229) |
| `rules.spec.asn_matcher.asn_sets.namespace` | [rules.spec.asn_matcher.asn_sets.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-99a97c4f0a4146320675bdcd33aa00675cd8258db028abd4825a1fd3ead7316a) |
| `rules.spec.asn_matcher.asn_sets.tenant` | [rules.spec.asn_matcher.asn_sets.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-369567d1f161d8eddb1a0069219851febdd01eb506dc726a5182b31d19c4ba79) |
| `rules.spec.asn_matcher.asn_sets.uid` | [rules.spec.asn_matcher.asn_sets.uid](resources--rate_limiter_policy--reference--group-001.md#canonical-23078be93a299abdbdfed134d1bc7fd1deed90e5d204dfaff5ad3a3832b04b98) |
| `rules.spec.bypass_rate_limiter` | [rules.spec.bypass_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-bfb8a18fe30af2faccabafffea980236d5fc268c2994689669ee948a47ab0b28) |
| `rules.spec.country_list` | [rules.spec.country_list](resources--rate_limiter_policy--reference--group-001.md#canonical-317c15156ef6dd7ebc600fa04e58613c5fb117310eebefa4f12a0f9137fae361) |
| `rules.spec.country_list.country_codes` | [rules.spec.country_list.country_codes](resources--rate_limiter_policy--reference--group-001.md#canonical-568a052ca95d77075982372c19e764e7f067f4fb45d5423003c354b9009df160) |
| `rules.spec.country_list.invert_match` | [rules.spec.country_list.invert_match](resources--rate_limiter_policy--reference--group-001.md#canonical-9b8c2b9608267155cbef5d175b83a6b94a536d94397a15c985b45ecf5b6492e6) |
| `rules.spec.custom_rate_limiter` | [rules.spec.custom_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-ae950a5e84bd3c63cf2aa3292bb67359aa6eeba5f5aee43bccec4482bdbae914) |
| `rules.spec.custom_rate_limiter.name` | [rules.spec.custom_rate_limiter.name](resources--rate_limiter_policy--reference--group-001.md#canonical-fee094cd36995f5f3861566670cd02713d63ce17f12ce63e971e493844bcf21b) |
| `rules.spec.custom_rate_limiter.namespace` | [rules.spec.custom_rate_limiter.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-9a4dde97bd54c8d42ef78310b425d6eb1a1e1413ed802ce6d5239f347d16d17f) |
| `rules.spec.custom_rate_limiter.tenant` | [rules.spec.custom_rate_limiter.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-bd472a365ba7a730c0d1459d35a69b2dbb7cb7a63bf470280755f00f9a063400) |
| `rules.spec.domain_matcher` | [rules.spec.domain_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-cfb12d152f296937b2bee6bd3a300327f617d6fe675137739885fb050cd3ebc1) |
| `rules.spec.domain_matcher.exact_values` | [rules.spec.domain_matcher.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-251b5685fcaa28597def692b021aabd7313d8d3b0b60f872a5a1c94946aad309) |
| `rules.spec.domain_matcher.regex_values` | [rules.spec.domain_matcher.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-5c66025d528010a0a66adb3333f8c243bb13b0b9976ab7f1a9fcd6083a43f7db) |
| `rules.spec.headers` | [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-e815c8e086298db976d17dfc7ff4a945368a1bc6ba70f20a9e2a5195caecb2fe) |
| `rules.spec.headers.check_not_present` | [rules.spec.headers.check_not_present](resources--rate_limiter_policy--reference--group-001.md#canonical-25e2a05352d877e60b60a6cdb0794c60cee00c30a822473bf3c7887b88d1e5a1) |
| `rules.spec.headers.check_present` | [rules.spec.headers.check_present](resources--rate_limiter_policy--reference--group-001.md#canonical-2ff954b6dcdffda6a01d5c441691a87a47e42f441065f29c3f7c22c5b53702bb) |
| `rules.spec.headers.invert_matcher` | [rules.spec.headers.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-6af5f5cefd3e3c66b13e51663ca8728596e507bcb1914487ab1832c01137188b) |
| `rules.spec.headers.item` | [rules.spec.headers.item](resources--rate_limiter_policy--reference--group-001.md#canonical-e8f7c97f85f70bb0c693204456f173036733e7d396307f1bf20ae72aaa61609c) |
| `rules.spec.headers.item.exact_values` | [rules.spec.headers.item.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-ae4427adda1d33ca62be995582435a960dd8d821dbe586016aacefe916515d88) |
| `rules.spec.headers.item.regex_values` | [rules.spec.headers.item.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-84c315d9aa96f7a6bf3fe0b4c62b0432d65e862eef8aaf5898e132b4d6fc022d) |
| `rules.spec.headers.item.transformers` | [rules.spec.headers.item.transformers](resources--rate_limiter_policy--reference--group-001.md#canonical-38577224f5144bae2eddeb0531f38c6fce9ae20bcf4e657912582ed5bc05f54f) |
| `rules.spec.headers.name` | [rules.spec.headers.name](resources--rate_limiter_policy--reference--group-001.md#canonical-e98090a71ea9f0ae4c3d2b01c77d8a821f33ec49970569d0733ee3290c457c93) |
| `rules.spec.http_method` | [rules.spec.http_method](resources--rate_limiter_policy--reference--group-001.md#canonical-fd4881419232303cb7de408fe3937bd2d8d61ae6983f006b7f9451cafff93c43) |
| `rules.spec.http_method.invert_matcher` | [rules.spec.http_method.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-6e0d3ed0fe34351f9d005012c9a3df296bede49ff0a822da7008321d49436d82) |
| `rules.spec.http_method.methods` | [rules.spec.http_method.methods](resources--rate_limiter_policy--reference--group-001.md#canonical-ebd0add567dcd231d59e7550ddc4caab9af7d6456c06e45e600263269135a117) |
| `rules.spec.ip_matcher` | [rules.spec.ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-dd2405117e2e1481d3156c991db002c19d105c51d46a651ea1106dcc88d82b06) |
| `rules.spec.ip_matcher.invert_matcher` | [rules.spec.ip_matcher.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-ec4c3798d8505f85afe6a8dcfeb2c110eee7193864b658edc9fbdb39ddede8ab) |
| `rules.spec.ip_matcher.prefix_sets` | [rules.spec.ip_matcher.prefix_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-c220bc84e50dc0481a98f0a97e737ea255fd86effbe1f6d07400bf9844ff29bd) |
| `rules.spec.ip_matcher.prefix_sets.kind` | [rules.spec.ip_matcher.prefix_sets.kind](resources--rate_limiter_policy--reference--group-001.md#canonical-a9cab52b4d140fd0d510951090720f06998ad570b2a046bd278bb8eaeb5b7dde) |
| `rules.spec.ip_matcher.prefix_sets.name` | [rules.spec.ip_matcher.prefix_sets.name](resources--rate_limiter_policy--reference--group-001.md#canonical-3408ad5828adfecb7cd7a197bff258ff73a5c3175bcc4864f4b6698c228f6aa3) |
| `rules.spec.ip_matcher.prefix_sets.namespace` | [rules.spec.ip_matcher.prefix_sets.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-b01292bb9ba69aab6859e3615b9a046948eab7277678ade269f13a966bd2618c) |
| `rules.spec.ip_matcher.prefix_sets.tenant` | [rules.spec.ip_matcher.prefix_sets.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-b7865ce258eab3866782b50aa4e19d7a700d49e3ab903c9cf9a3a00267b9b3a7) |
| `rules.spec.ip_matcher.prefix_sets.uid` | [rules.spec.ip_matcher.prefix_sets.uid](resources--rate_limiter_policy--reference--group-001.md#canonical-2e92b3dc69c42eeb4198e946abda79c8c9c8bdd50c8f1d769e0b4824c54ae00d) |
| `rules.spec.ip_prefix_list` | [rules.spec.ip_prefix_list](resources--rate_limiter_policy--reference--group-001.md#canonical-d95bd8df95cc76d2652840c5bac1b95b8323ddceb6f38cd5e8b71f572dd3f21e) |
| `rules.spec.ip_prefix_list.invert_match` | [rules.spec.ip_prefix_list.invert_match](resources--rate_limiter_policy--reference--group-001.md#canonical-56001455a98d66b9917bed339d919d4af024efec1975528564a49a1a27a5cf51) |
| `rules.spec.ip_prefix_list.ip_prefixes` | [rules.spec.ip_prefix_list.ip_prefixes](resources--rate_limiter_policy--reference--group-001.md#canonical-cd522a3e2a77e375ad71ac82ef51fec817ff52b10520d4a4cf4edca32a211e57) |
| `rules.spec.path` | [rules.spec.path](resources--rate_limiter_policy--reference--group-001.md#canonical-d15878353d46fdaa60e87cc0e5cf470d8e416da048e9b46ed11a90169ff2e64c) |
| `rules.spec.path.encoded_path_matcher` | [rules.spec.path.encoded_path_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-8ecbb81cdb0ce896d5df239b5a7e4ff6ce03da8476a4570ca9ed4adc10463dd0) |
| `rules.spec.path.exact_values` | [rules.spec.path.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-115ad9b6c35225208621773de12e29d58d82e3d61ac76a72f515812ef9a40c39) |
| `rules.spec.path.invert_matcher` | [rules.spec.path.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-231fe6b8cb279e5ca3c1ff9ea05ad0e968b9338ad7e1cb32e3a20a54e860746a) |
| `rules.spec.path.prefix_values` | [rules.spec.path.prefix_values](resources--rate_limiter_policy--reference--group-001.md#canonical-856bfc5bf35473c2f543e5e8c0ac45ce257473885ce18312c528286ec4cfba99) |
| `rules.spec.path.regex_values` | [rules.spec.path.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-0c65eac1ddf00eb0f50ed24c8cb87475423f60d9d4f405d7a19489cd606d191b) |
| `rules.spec.path.suffix_values` | [rules.spec.path.suffix_values](resources--rate_limiter_policy--reference--group-001.md#canonical-5e9139945192e780e7330d5d10b0c8453a5cdd9c0be52a35bb31ac0fc99a76e3) |
| `rules.spec.path.transformers` | [rules.spec.path.transformers](resources--rate_limiter_policy--reference--group-001.md#canonical-1e8803222f0b824f67b4a1fe4bf745cb71d31efa20a35d1b26e0668ffbe03a2e) |
| `rules.spec.segment_policy` | [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-781fd778c2cacb5613afc0ed675c94459bbb508504ceeefcd11c177aa56fa352) |
| `rules.spec.segment_policy.dst_any` | [rules.spec.segment_policy.dst_any](resources--rate_limiter_policy--reference--group-001.md#canonical-739a91dd1c9814931d45ebdc0123ff121d327b8656f5511586a49272f11a0682) |
| `rules.spec.segment_policy.dst_segments` | [rules.spec.segment_policy.dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0e77093791998a02cfc1ec64be667266bfbf510a248a95a8383297f76888ad02) |
| `rules.spec.segment_policy.dst_segments.segments` | [rules.spec.segment_policy.dst_segments.segments](resources--rate_limiter_policy--reference--group-001.md#canonical-9c5f01a5a645d91ffa78b6f962864f6935aa16b744254d7f8c2ce0ed0d79b9ba) |
| `rules.spec.segment_policy.dst_segments.segments.name` | [rules.spec.segment_policy.dst_segments.segments.name](resources--rate_limiter_policy--reference--group-001.md#canonical-9051f274165d37adec43a2b168bcf66ffc163bc22029c5f2e3242c3f98a4b1d2) |
| `rules.spec.segment_policy.dst_segments.segments.namespace` | [rules.spec.segment_policy.dst_segments.segments.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-68998df1997b46fe64aa8a1941270057f11434c4bf58eeadcdcb653600e811ff) |
| `rules.spec.segment_policy.dst_segments.segments.tenant` | [rules.spec.segment_policy.dst_segments.segments.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-231601d8704b1ca7963064ecf1c420fb67f895896b1f24f6424841106a123aeb) |
| `rules.spec.segment_policy.intra_segment` | [rules.spec.segment_policy.intra_segment](resources--rate_limiter_policy--reference--group-001.md#canonical-dfbde1bff989374e0597b32c6f9ddf0b40e0a32ec03d33c4f2a2ebdf3d0bb937) |
| `rules.spec.segment_policy.src_any` | [rules.spec.segment_policy.src_any](resources--rate_limiter_policy--reference--group-001.md#canonical-e589f7d37ae7a8350015177737bdaa094b0d96405ff8281726d0e0b4c02aef91) |
| `rules.spec.segment_policy.src_segments` | [rules.spec.segment_policy.src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-014adb043eff200a4110234c5eac80becc1abe5697e96a048488c2811ff9ea69) |
| `rules.spec.segment_policy.src_segments.segments` | [rules.spec.segment_policy.src_segments.segments](resources--rate_limiter_policy--reference--group-001.md#canonical-f42a3b5879455b4c3f9b07e561f446acffb70f0421b9892b148c55554656fdba) |
| `rules.spec.segment_policy.src_segments.segments.name` | [rules.spec.segment_policy.src_segments.segments.name](resources--rate_limiter_policy--reference--group-001.md#canonical-09a4568eebdc7063dcae8d9065ea9b4b49ffa6d792b6b1b9fa879653927ece27) |
| `rules.spec.segment_policy.src_segments.segments.namespace` | [rules.spec.segment_policy.src_segments.segments.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-8b4777d17b6eba8e569794489073d3c956af7c3e534806017af84891c914d838) |
| `rules.spec.segment_policy.src_segments.segments.tenant` | [rules.spec.segment_policy.src_segments.segments.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-2d44c9ff909f35acf43dc7b30f55d2323664db04c6d24b4ff29462dee50420ac) |
| `server_name` | [server_name](resources--rate_limiter_policy--reference--group-001.md#canonical-105c5085315080237edf5a6df9cafbf6959e9c476b7c8611cee03bc87fd34fc6) |
| `server_name_matcher` | [server_name_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-0147585cf633a964d017070090c9913886726c790787618a9b68a08dcfde66be) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-15169173abd7dea0b9a1e59c12c1c196293c265a6f742072d41e3bad5652ad3f) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-c0cbadb432d16ca9bfe413e49bdc1056fad3ab2a57d12b5999e914e7a93f3290) |
| `server_selector` | [server_selector](resources--rate_limiter_policy--reference--group-001.md#canonical-ffa987c4a6cd7dd45162b28bcff8f3f61c7eea22399e91b5f5c65e51281ab0ff) |
| `server_selector.expressions` | [server_selector.expressions](resources--rate_limiter_policy--reference--group-001.md#canonical-7531f93e5218c6ef5cfc94dec0a8e35941c972f5e63f798b098cdfed6d53ac5c) |
| `timeouts` | [timeouts](resources--rate_limiter_policy--reference--group-001.md#canonical-e64fdbe480ab2dce371d1204d4479595cd2d309062fa091ee6e9a42651337c11) |
| `timeouts.create` | [timeouts.create](resources--rate_limiter_policy--reference--group-001.md#canonical-ceb26058658ae462a01806a290bb8a27599deb7100576ccc23db2a85de7f6008) |
| `timeouts.delete` | [timeouts.delete](resources--rate_limiter_policy--reference--group-001.md#canonical-0cf6abf7c561d284351fc50b203b71843e267b1383eb513d8fab93375be4d1e0) |
| `timeouts.read` | [timeouts.read](resources--rate_limiter_policy--reference--group-001.md#canonical-b843b0c266f574d3838f340f6f2c242ff51ffc007cbd8813c1d662ea27bc4701) |
| `timeouts.update` | [timeouts.update](resources--rate_limiter_policy--reference--group-001.md#canonical-a132780cf5d0c457c55bf71df8e68eeebef895f33ee0af1aee5684d302cf14da) |

<a id="canonical-a7c054ef9d1e38cd41c2a91da26dc587fa82f798579589825f6a64056f454ab9"></a>

## Next pages — Property reference / 32b78a761a0e / 13

- [any_server](resources--rate_limiter_policy--reference--group-001.md#canonical-259eae0b31626643a9d6ca160de383c88e28187307228398093a6b5c43ff6439)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [server_name_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3e759f0ab8dc83a52fb91efa326ee68914c628ce86919f316e19a1c4368373a3)
- [server_selector](resources--rate_limiter_policy--reference--group-001.md#canonical-a76ecf72963cf5b3f31077a4892876c381eb73eebf600858b6dd88334e80a6c7)
- [timeouts](resources--rate_limiter_policy--reference--group-001.md#canonical-596b33ac3e31c789c6ff45eb3514ad49a37cf7f206057468810e3be146744ecf)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-259eae0b31626643a9d6ca160de383c88e28187307228398093a6b5c43ff6439"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebd4ad9f58db9397313a807e57658d628e5ecc6c27f7c349b78db8115bf66e88"></a>

## any_server — any_server / 6ee83cf49f47 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- any_server

<a id="canonical-2777f1c8cd4230c9d6cd44d7996500aecadd2474f006fd9c13a28e0c52f9b947"></a>

Type: `["object", {}]`. Optional.

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

- [any_server](resources--rate_limiter_policy--reference--group-001.md#canonical-2777f1c8cd4230c9d6cd44d7996500aecadd2474f006fd9c13a28e0c52f9b947)
- [server_name](resources--rate_limiter_policy--reference--group-001.md#canonical-105c5085315080237edf5a6df9cafbf6959e9c476b7c8611cee03bc87fd34fc6)
- [server_name_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-0147585cf633a964d017070090c9913886726c790787618a9b68a08dcfde66be)
- [server_selector](resources--rate_limiter_policy--reference--group-001.md#canonical-ffa987c4a6cd7dd45162b28bcff8f3f61c7eea22399e91b5f5c65e51281ab0ff)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_server = {}
```

<a id="canonical-631d8b92dd4d9527c8381adbc5397ec5195ebeef79994a8f594a7197df452bce"></a>

## Direct properties — any_server / 6ee83cf49f47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-96f42ef20c9dab814e5e469069da1e306b691c89561d565679768f42e8df6aff"></a>

## Next pages — any_server / 6ee83cf49f47 / 4

- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6b75149175630b2c34ba3c1dfd8321c52cfc06f3df0889ffd7493e244cff2bc"></a>

## rules — rules / e12c3811f5eb / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- rules

<a id="canonical-a666a5c6864db8036f73005a2e011f3de7b9d71bef9b4afe5643f44e971e06cf"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-575e9fd6e0f6712d5d0d25ed17fd277be624d3f53f15f55cc4b69ee541962d2a"></a>

## Direct properties — rules / e12c3811f5eb / 3

- [metadata](resources--rate_limiter_policy--reference--group-001.md#canonical-cb184bbb4b4eb32ba866b98d2233c63de63866fd23acb411420fdad83a9c6909): complete subsection reference.

- [spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d): complete subsection reference.

<a id="canonical-aea5bc41cc2fb8dda44a2407f84b0c4682b534e95d320da61abcf47a7be25dc2"></a>

## Next pages — rules / e12c3811f5eb / 4

- [rules.metadata](resources--rate_limiter_policy--reference--group-001.md#canonical-cb184bbb4b4eb32ba866b98d2233c63de63866fd23acb411420fdad83a9c6909)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-cb184bbb4b4eb32ba866b98d2233c63de63866fd23acb411420fdad83a9c6909"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52f0d1a501b5cc727af50c964c8cddc810a8d87666d4751221c26f4cadb5f525"></a>

## rules.metadata — rules.metadata / 771a8367e0ab / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- rules.metadata

<a id="canonical-ce6aa3058b8d07a250989801e13c2d2b495a02d63a5d61307efa40a464ca392c"></a>

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

<a id="canonical-fc9137754acb32900b900828ef513e1a94bdaceae6ce0c418255bf556a4ed871"></a>

## Direct properties — rules.metadata / 771a8367e0ab / 3

<a id="canonical-55642899b2f7d41d1e55ed64659b6532171b012fdeaddecd88e4f7f403c8a1c6"></a>

<a id="canonical-390443dcd61eb2ca4671aca6627e204666e7db3fad973fe760c289293469a223"></a>

## description_spec property — rules.metadata / 771a8367e0ab / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-a9f45e916d72131a8d4eb4689adf11c46b95cf2724d868756047d31ad7fc330b"></a>

<a id="canonical-18ac334cc2d9396fd36824a034bd28b54acce2248e74e1d545b99a5a602f66c1"></a>

## name property — rules.metadata / 771a8367e0ab / 5

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

<a id="canonical-e711d9b0be7b2869dec0f88a7d57fc30c2f61748410c265b3932cffa773ad497"></a>

## Next pages — rules.metadata / 771a8367e0ab / 6

- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6045d2d7b671a46958ae4fe1a12711d63524d8085f504aa51608fbe13f73543"></a>

## rules.spec — rules.spec / 8058917806b0 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- rules.spec

<a id="canonical-9f1aef64cd75db85245caad89c90efa36c3eaadf1e11e8ae9265721b03bad7a1"></a>

Type: `"object"`. single nested block, Optional.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Upstream description:

Shape of Rate Limiter Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_country",
    "country_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("apply_rate_limiter",
    "bypass_rate_limiter"),
  validators.ConflictingObjectAttributes("apply_rate_limiter",
    "custom_rate_limiter"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("bypass_rate_limiter",
    "custom_rate_limiter"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-action_choice": "[\"apply_rate_limiter\",\"bypass_rate_limiter\",\"custom_rate_limiter\"]",
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-country_choice": "[\"any_country\",\"country_list\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-7edd6a0260636e9721083cc2f5d87b96207e2f88ee124262ad75ab22196ca2d0"></a>

## Direct properties — rules.spec / 8058917806b0 / 3

- [any_asn](resources--rate_limiter_policy--reference--group-001.md#canonical-47ee8d152eaa107fdf880f2048212d8c305941e544559f0877007d7187b05334): complete subsection reference.

- [any_country](resources--rate_limiter_policy--reference--group-001.md#canonical-dca07798ed30757cb9c1d037ebc3c46c43b0e4e30586a55b25fe7ace81febf88): complete subsection reference.

- [any_ip](resources--rate_limiter_policy--reference--group-001.md#canonical-5eb6b3ad4ae1fb7abf934aa2a1e13ccda0597b7126df6eafc62c7095d5bf7fd9): complete subsection reference.

- [apply_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-f9d1d88f699829893780352b1014b093c69b2dd59ee22311715f3763ca4a4775): complete subsection reference.

- [asn_list](resources--rate_limiter_policy--reference--group-001.md#canonical-c44f7988511a2bde7cf0fc0ab5d202cfe58ebb9df8c9136d2570c8e66dad94fe): complete subsection reference.

- [asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-590b56d01ac236ebc17628be45b58b1c3142b3c9d7cc487a725b1ed726d674a9): complete subsection reference.

- [bypass_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-caea7f2f2d8902fde0a991538fb4578eefa8883c74dd7ca4bf193d334c5e412f): complete subsection reference.

- [country_list](resources--rate_limiter_policy--reference--group-001.md#canonical-414147638adda44dc92258c785358393ad91edbac6fe96f3d8524ee47ef2f48d): complete subsection reference.

- [custom_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-9c34d0e1d9269ddbe6f3e78dd4ce10950a8590d8df38a31ea4b8837fb3214d65): complete subsection reference.

- [domain_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-f027e7ff63ee4fda7e8e21c1a0e75a29562b85d8f9f87d8bf40088d9e31ce282): complete subsection reference.

- [headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7): complete subsection reference.

- [http_method](resources--rate_limiter_policy--reference--group-001.md#canonical-bd7095efea16a74b537ad6b4950d3ea63ab245f2d546746c0e78f4b7222e97d6): complete subsection reference.

- [ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-d5746e1b12da1121317c0a7ce76c63b6537db7d3ac8d2fc98d62c1291c669686): complete subsection reference.

- [ip_prefix_list](resources--rate_limiter_policy--reference--group-001.md#canonical-4a0d4d9a1df889b0ae2fe780c8e25e302c470b0e71cebb501890e3aa507f0970): complete subsection reference.

- [path](resources--rate_limiter_policy--reference--group-001.md#canonical-9d8f11d8ce1ee7122604889577d8166719df102f378555ef3b2c8e65d46c3fb9): complete subsection reference.

- [segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076): complete subsection reference.

<a id="canonical-74bb89836700bcf65f377e3092f1726eef453277471bcee6b7e349e8ace106a1"></a>

## Next pages — rules.spec / 8058917806b0 / 4

- [rules.spec.any_asn](resources--rate_limiter_policy--reference--group-001.md#canonical-47ee8d152eaa107fdf880f2048212d8c305941e544559f0877007d7187b05334)
- [rules.spec.any_country](resources--rate_limiter_policy--reference--group-001.md#canonical-dca07798ed30757cb9c1d037ebc3c46c43b0e4e30586a55b25fe7ace81febf88)
- [rules.spec.any_ip](resources--rate_limiter_policy--reference--group-001.md#canonical-5eb6b3ad4ae1fb7abf934aa2a1e13ccda0597b7126df6eafc62c7095d5bf7fd9)
- [rules.spec.apply_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-f9d1d88f699829893780352b1014b093c69b2dd59ee22311715f3763ca4a4775)
- [rules.spec.asn_list](resources--rate_limiter_policy--reference--group-001.md#canonical-c44f7988511a2bde7cf0fc0ab5d202cfe58ebb9df8c9136d2570c8e66dad94fe)
- [rules.spec.asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-590b56d01ac236ebc17628be45b58b1c3142b3c9d7cc487a725b1ed726d674a9)
- [rules.spec.bypass_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-caea7f2f2d8902fde0a991538fb4578eefa8883c74dd7ca4bf193d334c5e412f)
- [rules.spec.country_list](resources--rate_limiter_policy--reference--group-001.md#canonical-414147638adda44dc92258c785358393ad91edbac6fe96f3d8524ee47ef2f48d)
- [rules.spec.custom_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-9c34d0e1d9269ddbe6f3e78dd4ce10950a8590d8df38a31ea4b8837fb3214d65)
- [rules.spec.domain_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-f027e7ff63ee4fda7e8e21c1a0e75a29562b85d8f9f87d8bf40088d9e31ce282)
- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7)
- [rules.spec.http_method](resources--rate_limiter_policy--reference--group-001.md#canonical-bd7095efea16a74b537ad6b4950d3ea63ab245f2d546746c0e78f4b7222e97d6)
- [rules.spec.ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-d5746e1b12da1121317c0a7ce76c63b6537db7d3ac8d2fc98d62c1291c669686)
- [rules.spec.ip_prefix_list](resources--rate_limiter_policy--reference--group-001.md#canonical-4a0d4d9a1df889b0ae2fe780c8e25e302c470b0e71cebb501890e3aa507f0970)
- [rules.spec.path](resources--rate_limiter_policy--reference--group-001.md#canonical-9d8f11d8ce1ee7122604889577d8166719df102f378555ef3b2c8e65d46c3fb9)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-47ee8d152eaa107fdf880f2048212d8c305941e544559f0877007d7187b05334"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ffb9cac3a8b95cb0b9d49757bd4b58f749f7ec2a23b89fd1b481b0d3b36fc8c"></a>

## rules.spec.any_asn — rules.spec.any_asn / e6fbae337d08 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.any_asn

<a id="canonical-4733ad3b5dc4f5afe21afd45d5a3855006cdb5ccfd71f9ca3ca5670276380a9b"></a>

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
any_asn = {}
```

<a id="canonical-14e580594d8a5332887f24800263ef0a6cd85f65b1e013669938b3362d623354"></a>

## Direct properties — rules.spec.any_asn / e6fbae337d08 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7912b880bcd6729d000a0810191c63067c0fd9b92b2638cf1a24257d8fdb3dc0"></a>

## Next pages — rules.spec.any_asn / e6fbae337d08 / 4

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-dca07798ed30757cb9c1d037ebc3c46c43b0e4e30586a55b25fe7ace81febf88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b4962a1e5a811bacb756812ee0bf2690f5e2578d83b66a7a26dab77e48850ec"></a>

## rules.spec.any_country — rules.spec.any_country / 324216391e91 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.any_country

<a id="canonical-c79c15bcd26aef1495ecef47840a27f4d2454c4460257b4f51a6f11e05c7fda6"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_country = {}
```

<a id="canonical-c0d9bc01c1417136f56453d3c30ae002fbeae7c7d34da5e405c68e0590110603"></a>

## Direct properties — rules.spec.any_country / 324216391e91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87e5fdae73072f4ac3723f7f6f4862fe3ae62ffaaf4a44c1cfe8e3b0405ef837"></a>

## Next pages — rules.spec.any_country / 324216391e91 / 4

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-5eb6b3ad4ae1fb7abf934aa2a1e13ccda0597b7126df6eafc62c7095d5bf7fd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c6fc8d844136c5ba315ffda04d009aced8f2c9f2123816ed5f536cac3a374ff"></a>

## rules.spec.any_ip — rules.spec.any_ip / aaebfa165b6e / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.any_ip

<a id="canonical-c2d835d0ddccae21fe28c37ade04017f8ed9eec4d89e73d9e631825eec175b1a"></a>

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
any_ip = {}
```

<a id="canonical-05a42cbdb92ca5385d70b2c6925f5ad561cc02294a6a87b96839c09bfdf04ade"></a>

## Direct properties — rules.spec.any_ip / aaebfa165b6e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e2a9f7a455f90e34573cffd4fe711bc922e7cbc665c3144849969caf2a08a50f"></a>

## Next pages — rules.spec.any_ip / aaebfa165b6e / 4

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-f9d1d88f699829893780352b1014b093c69b2dd59ee22311715f3763ca4a4775"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dc55e9651a1fff40f7b0d9cb0ed053f7b897f6f5d5b567a89a09b03ba77a0f7"></a>

## rules.spec.apply_rate_limiter — rules.spec.apply_rate_limiter / ede3f3fab7f6 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.apply_rate_limiter

<a id="canonical-dea393d208d5da742a7cb3f68d6b2701839da7e19bc98115ec68c709ca246d32"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
apply_rate_limiter = {}
```

<a id="canonical-f6f41f78af7671eee128620427fd59dd9cebae2311328abd517eefa9642b2659"></a>

## Direct properties — rules.spec.apply_rate_limiter / ede3f3fab7f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a6b3a1651cc4f629e67864fcd9e06d15fcd0b87008c24d219792e27be31eedd4"></a>

## Next pages — rules.spec.apply_rate_limiter / ede3f3fab7f6 / 4

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-c44f7988511a2bde7cf0fc0ab5d202cfe58ebb9df8c9136d2570c8e66dad94fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-855fb1025b38e76d51b86cb602abede88f63decb687eb1ea3b809a1549b9911c"></a>

## rules.spec.asn_list — rules.spec.asn_list / 2109cf623128 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.asn_list

<a id="canonical-e067111f75e87f22fb33bed9da886ad51ff351bbaaf7e5272c25ca90883fde92"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-62c296f012198e783d10fd9b7c298a454ce169310a928ed37d24314feb8bdd4c"></a>

## Direct properties — rules.spec.asn_list / 2109cf623128 / 3

<a id="canonical-998f95aa6ffcbb30ad25b284f1c5df079f22f9649e230ea3e43dae064baa8dab"></a>

<a id="canonical-b4cc62e8f9dc73a34b52e66f7d4d1002f7bd48440c804a150c70284e7f5ae361"></a>

## as_numbers property — rules.spec.asn_list / 2109cf623128 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-806a7bf1a2e949ef278f7c1f90c0ecb2ca4ad41601e81250bbc7413ce018b70b"></a>

## Next pages — rules.spec.asn_list / 2109cf623128 / 5

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-590b56d01ac236ebc17628be45b58b1c3142b3c9d7cc487a725b1ed726d674a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d86ea908c8c4c0517ee4d0e20622903bf8dc12eb475b626ed470108ce5ed71d"></a>

## rules.spec.asn_matcher — rules.spec.asn_matcher / c9fbfd056366 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.asn_matcher

<a id="canonical-485e4cabf3bcfa154eb79cde868d438933262be6238d62300972f6ca4b5d6032"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-f75de12f5fb734b36e2b45429e24fd62295002c267db9fce3fa3915518b4f45d"></a>

## Direct properties — rules.spec.asn_matcher / c9fbfd056366 / 3

- [asn_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-eac27a4760c128b6d95f7ea758328247ceb07ebcb87e7c3c03746ed23b16b5a1): complete subsection reference.

<a id="canonical-aee2153386c5b47aa4cd3e6940b970c73fbd14b271e27810ad0c0ebba5b72805"></a>

## Next pages — rules.spec.asn_matcher / c9fbfd056366 / 4

- [rules.spec.asn_matcher.asn_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-eac27a4760c128b6d95f7ea758328247ceb07ebcb87e7c3c03746ed23b16b5a1)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-eac27a4760c128b6d95f7ea758328247ceb07ebcb87e7c3c03746ed23b16b5a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-178015957be9fdbcc5b0e3945a6e9bedf0d51442ac0ddc2ced331fe4de3cfcdf"></a>

## rules.spec.asn_matcher.asn_sets — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-590b56d01ac236ebc17628be45b58b1c3142b3c9d7cc487a725b1ed726d674a9)
- rules.spec.asn_matcher.asn_sets

<a id="canonical-0727ef6b81a4bd2df78e7f2140f6dd5ec99c1df10034b83bacaa2bab98d01a7d"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-9752c07371161e64c2aadaae92bf58e053eb3adc8dab41c6f06c96a8de023acb"></a>

## Direct properties — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 3

<a id="canonical-c76b03864d4ae9c81a218256cb2abfa387a74ba9e6c1d99bbe025cca66d96349"></a>

<a id="canonical-8c044df408263ae644508b956a5e31e8ed9aae28eb0abd8f73ab7046e751d107"></a>

## kind property — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 4

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

<a id="canonical-87ccada235a7ccd8ae9bb07c7f2e545b3b030026eddf24e63143dfde60623229"></a>

<a id="canonical-1fa38189ee0e570b0919cf5f6f5ae110bc9e7c0019a1c1e3ef711be231e317dc"></a>

## name property — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 5

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

<a id="canonical-99a97c4f0a4146320675bdcd33aa00675cd8258db028abd4825a1fd3ead7316a"></a>

<a id="canonical-3053a26c117ead053a964d7471975c12656e17660519e92363457831167c5f14"></a>

## namespace property — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 6

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

<a id="canonical-369567d1f161d8eddb1a0069219851febdd01eb506dc726a5182b31d19c4ba79"></a>

<a id="canonical-050619a0192ccd5b0a36b0dd02f15fed228111912552ed3fd3b0d53038d0a365"></a>

## tenant property — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 7

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

<a id="canonical-23078be93a299abdbdfed134d1bc7fd1deed90e5d204dfaff5ad3a3832b04b98"></a>

<a id="canonical-45ef5e9f3e88860300590694c70b28fee9056431dfe39d90ad23627fa16eee56"></a>

## uid property — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 8

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

<a id="canonical-df19035304664e29d4a7ff0fa6d57194a87826f9466ed9f912bff6feab48e4ee"></a>

## Next pages — rules.spec.asn_matcher.asn_sets / 977bf9a35c8c / 9

- [rules.spec.asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-590b56d01ac236ebc17628be45b58b1c3142b3c9d7cc487a725b1ed726d674a9)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-caea7f2f2d8902fde0a991538fb4578eefa8883c74dd7ca4bf193d334c5e412f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa15f178f0d9017116d92c5a7b8fe5dd5ca31c68fa8035912885368b209137d9"></a>

## rules.spec.bypass_rate_limiter — rules.spec.bypass_rate_limiter / 6b6f9c292a32 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.bypass_rate_limiter

<a id="canonical-bfb8a18fe30af2faccabafffea980236d5fc268c2994689669ee948a47ab0b28"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
bypass_rate_limiter = {}
```

<a id="canonical-0b9682b09bb56526b148ee4b6de9a6d1b399387224ae7f25ac8682e9574ee26d"></a>

## Direct properties — rules.spec.bypass_rate_limiter / 6b6f9c292a32 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b54ef61d5759128586cab0b3664b2a0dcc20eecd361959d2e51527015fa204e2"></a>

## Next pages — rules.spec.bypass_rate_limiter / 6b6f9c292a32 / 4

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-414147638adda44dc92258c785358393ad91edbac6fe96f3d8524ee47ef2f48d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c551e170317f7b5ce9a2cc172050546287d3927f88fa6cd21d8e33c26ffbfa67"></a>

## rules.spec.country_list — rules.spec.country_list / dcf2d54018e4 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.country_list

<a id="canonical-317c15156ef6dd7ebc600fa04e58613c5fb117310eebefa4f12a0f9137fae361"></a>

Type: `"object"`. single nested block, Optional.

Country Codes List. List of Country Codes to match against.

Upstream description:

List of Country Codes to match against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("country_codes")}
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
country_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc5daebc591a93736492a404fb2917aed2ff6b21865e779ed926b614bc416b12"></a>

## Direct properties — rules.spec.country_list / dcf2d54018e4 / 3

<a id="canonical-568a052ca95d77075982372c19e764e7f067f4fb45d5423003c354b9009df160"></a>

<a id="canonical-9c9803ba60a4e3422a1a9c6e64201b507b4ce0f81b132a955c932f03549bd77f"></a>

## country_codes property — rules.spec.country_list / dcf2d54018e4 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

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

<a id="canonical-9b8c2b9608267155cbef5d175b83a6b94a536d94397a15c985b45ecf5b6492e6"></a>

<a id="canonical-b623c991e9bab0359353d61515cd8ca3c3aa8e8b921197929c7603ea7bb668a6"></a>

## invert_match property — rules.spec.country_list / dcf2d54018e4 / 5

Type: `"bool"`. Optional.

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

<a id="canonical-ceac445d522fc4ea4a4ba628d1dd984478a3046522cbcc586cdf5c2601f423d9"></a>

## Next pages — rules.spec.country_list / dcf2d54018e4 / 6

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-9c34d0e1d9269ddbe6f3e78dd4ce10950a8590d8df38a31ea4b8837fb3214d65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b559b6710e9e9a08be0bd7199d8d4ceb92e5a7c0bb1a73e99e80249a9d34b558"></a>

## rules.spec.custom_rate_limiter — rules.spec.custom_rate_limiter / 125e946469f0 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.custom_rate_limiter

<a id="canonical-ae950a5e84bd3c63cf2aa3292bb67359aa6eeba5f5aee43bccec4482bdbae914"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
custom_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-0de6b428f1a538f5a827d714d023d02124c25f48948abf18b3a99b1ac615f66b"></a>

## Direct properties — rules.spec.custom_rate_limiter / 125e946469f0 / 3

<a id="canonical-fee094cd36995f5f3861566670cd02713d63ce17f12ce63e971e493844bcf21b"></a>

<a id="canonical-ff71b79e2719dabee22359e4525d38edb330b60273e55577147880ce93d2a682"></a>

## name property — rules.spec.custom_rate_limiter / 125e946469f0 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-9a4dde97bd54c8d42ef78310b425d6eb1a1e1413ed802ce6d5239f347d16d17f"></a>

<a id="canonical-8b903ef92cb954733a2f806ed3c31caeffd1852fb9b3be5f1b8b76bec075ae38"></a>

## namespace property — rules.spec.custom_rate_limiter / 125e946469f0 / 5

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
}
```

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

<a id="canonical-bd472a365ba7a730c0d1459d35a69b2dbb7cb7a63bf470280755f00f9a063400"></a>

<a id="canonical-5d605caac2a96d3acbdaa497ed204d3c4f7ebffe4f6e59c543b8950045d04d15"></a>

## tenant property — rules.spec.custom_rate_limiter / 125e946469f0 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-f7491ac330f4667f29c4d2ed92b0555232293e1ca7d36d41ac91bec59d5e4e2d"></a>

## Next pages — rules.spec.custom_rate_limiter / 125e946469f0 / 7

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-f027e7ff63ee4fda7e8e21c1a0e75a29562b85d8f9f87d8bf40088d9e31ce282"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4360af0755af5dc1489587f5ef5ce38107c0c5cb6ce27ebe8afc29bbf40803c6"></a>

## rules.spec.domain_matcher — rules.spec.domain_matcher / 406006698d3b / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.domain_matcher

<a id="canonical-cfb12d152f296937b2bee6bd3a300327f617d6fe675137739885fb050cd3ebc1"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-6b99ad38f7c2401638daa3d7d278a2395c282e3830d5ad22e5e2f93ff23e7291"></a>

## Direct properties — rules.spec.domain_matcher / 406006698d3b / 3

<a id="canonical-251b5685fcaa28597def692b021aabd7313d8d3b0b60f872a5a1c94946aad309"></a>

<a id="canonical-053f924bdcfe1f4b60feb1a27165168194aacede5cec9fbc0f4b52c83ab65c20"></a>

## exact_values property — rules.spec.domain_matcher / 406006698d3b / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-5c66025d528010a0a66adb3333f8c243bb13b0b9976ab7f1a9fcd6083a43f7db"></a>

<a id="canonical-a9cf72d966fd8292a7008830f79255d8396b65e654f8c6150412f53f0d1eab61"></a>

## regex_values property — rules.spec.domain_matcher / 406006698d3b / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-da6b23fdf02a8df857caf0e0e9b97242d3aae9b95264e5f913b292a89e6dcb83"></a>

## Next pages — rules.spec.domain_matcher / 406006698d3b / 6

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46cd77906f8ae06f8eaa25e8a343a9314a0bd17b4f3643955241c86994329b8c"></a>

## rules.spec.headers — rules.spec.headers / 9daacacb083c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.headers

<a id="canonical-e815c8e086298db976d17dfc7ff4a945368a1bc6ba70f20a9e2a5195caecb2fe"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

<a id="canonical-fba8145c1c1ecdbe4423f2755bacb7c76dec80594935a40c3eb163f2c1830584"></a>

## Direct properties — rules.spec.headers / 9daacacb083c / 3

- [check_not_present](resources--rate_limiter_policy--reference--group-001.md#canonical-1facfbd9842f319838413e6ae8d7a34357a941201c44e3364b48fcff19cdc69a): complete subsection reference.

- [check_present](resources--rate_limiter_policy--reference--group-001.md#canonical-549b15d795dcbf494521cfcfddda4307cac0ff2b57d31d47935cf4ee9cc75e98): complete subsection reference.

<a id="canonical-6af5f5cefd3e3c66b13e51663ca8728596e507bcb1914487ab1832c01137188b"></a>

<a id="canonical-9d47e110caa2da0127e2c992b00aba5ea313aac8c54c2526f06e5ff8cfed971e"></a>

## invert_matcher property — rules.spec.headers / 9daacacb083c / 4

Type: `"bool"`. Optional.

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

- [item](resources--rate_limiter_policy--reference--group-001.md#canonical-bba37585fa00728f49090629b546175980858f1eff1e4e1c63416bee5cf40abe): complete subsection reference.

<a id="canonical-e98090a71ea9f0ae4c3d2b01c77d8a821f33ec49970569d0733ee3290c457c93"></a>

<a id="canonical-9d18617af4ef8438bc3bef9e1c851e1650ce2a99a9632702ffa2a42a35edff24"></a>

## name property — rules.spec.headers / 9daacacb083c / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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

<a id="canonical-3f9fdc04e8f8483d4b87fde5ceed04e682b7b2062f968814a756e813c4ddebe9"></a>

## Next pages — rules.spec.headers / 9daacacb083c / 6

- [rules.spec.headers.check_not_present](resources--rate_limiter_policy--reference--group-001.md#canonical-1facfbd9842f319838413e6ae8d7a34357a941201c44e3364b48fcff19cdc69a)
- [rules.spec.headers.check_present](resources--rate_limiter_policy--reference--group-001.md#canonical-549b15d795dcbf494521cfcfddda4307cac0ff2b57d31d47935cf4ee9cc75e98)
- [rules.spec.headers.item](resources--rate_limiter_policy--reference--group-001.md#canonical-bba37585fa00728f49090629b546175980858f1eff1e4e1c63416bee5cf40abe)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-1facfbd9842f319838413e6ae8d7a34357a941201c44e3364b48fcff19cdc69a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d6413e84269fa87ff4f0da12b37cbb705fdb26dac8eb8bd142b7595992a664b"></a>

## rules.spec.headers.check_not_present — rules.spec.headers.check_not_present / c7e640737a15 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7)
- rules.spec.headers.check_not_present

<a id="canonical-25e2a05352d877e60b60a6cdb0794c60cee00c30a822473bf3c7887b88d1e5a1"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-66f33a7bd39802dc6d5f0ed06998e0a4715b0224205fac2be642accd741160dd"></a>

## Direct properties — rules.spec.headers.check_not_present / c7e640737a15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12df52f34c30909251b2e4cd3a5ca0ddfb59c1edc53ec0e4e355215c51473d8b"></a>

## Next pages — rules.spec.headers.check_not_present / c7e640737a15 / 4

- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-549b15d795dcbf494521cfcfddda4307cac0ff2b57d31d47935cf4ee9cc75e98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-405370c950fd0c4876f910efef887152400434751795c7b4da33d0e956c7dc6b"></a>

## rules.spec.headers.check_present — rules.spec.headers.check_present / 558dfb2b66eb / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7)
- rules.spec.headers.check_present

<a id="canonical-2ff954b6dcdffda6a01d5c441691a87a47e42f441065f29c3f7c22c5b53702bb"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-e42f0bd5de7261b4245a7a077fa4ac9504784ff57951ddad21ddb2b93f4f9abc"></a>

## Direct properties — rules.spec.headers.check_present / 558dfb2b66eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d82356ad5388fbe99285a9feb111bc87847678529a9f5bdc99f35ff7bd6e7caa"></a>

## Next pages — rules.spec.headers.check_present / 558dfb2b66eb / 4

- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-bba37585fa00728f49090629b546175980858f1eff1e4e1c63416bee5cf40abe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24a424e4fc3fd1aa2304252b5ce542a857cd5cb7f2fa87eed49230748201ec99"></a>

## rules.spec.headers.item — rules.spec.headers.item / a349a2cfb09d / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7)
- rules.spec.headers.item

<a id="canonical-e8f7c97f85f70bb0c693204456f173036733e7d396307f1bf20ae72aaa61609c"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-ebb30b92ce2a160f0b3fa26361c21c942adcc7dbb2f76f6dcb43ff3fac801972"></a>

## Direct properties — rules.spec.headers.item / a349a2cfb09d / 3

<a id="canonical-ae4427adda1d33ca62be995582435a960dd8d821dbe586016aacefe916515d88"></a>

<a id="canonical-5693aa7b40323d9363c3c0601dfeaaf45df8a5d0bede8f8d60c72246a4edaaf8"></a>

## exact_values property — rules.spec.headers.item / a349a2cfb09d / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-84c315d9aa96f7a6bf3fe0b4c62b0432d65e862eef8aaf5898e132b4d6fc022d"></a>

<a id="canonical-3ee21de0162094928faebd60e8023e19984e0046d4fdc66ee75a75ad1f6601b5"></a>

## regex_values property — rules.spec.headers.item / a349a2cfb09d / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-38577224f5144bae2eddeb0531f38c6fce9ae20bcf4e657912582ed5bc05f54f"></a>

<a id="canonical-f3bfeff97ec8f3cd21e1dbbc9bbd1451dbcbcbc19fa20bb8db97ccbdb263bf32"></a>

## transformers property — rules.spec.headers.item / a349a2cfb09d / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-bac9555ceb8542640c93384542d78a6ac0d95884bc1230cc3ddd728ba46a5d25"></a>

## Next pages — rules.spec.headers.item / a349a2cfb09d / 7

- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-a9e39c39bce6533a5bde3d222b021fdeb028c72356f22cac7d860044999a73c7)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-bd7095efea16a74b537ad6b4950d3ea63ab245f2d546746c0e78f4b7222e97d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f13e188e92424709f1f684d226cc73953c6774113041456df9184959f2b5cb27"></a>

## rules.spec.http_method — rules.spec.http_method / d8920273d146 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.http_method

<a id="canonical-fd4881419232303cb7de408fe3937bd2d8d61ae6983f006b7f9451cafff93c43"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-70ecda39117e3587aefac3e1e349fd691a048d5c7fc253b544b4308ef128ad26"></a>

## Direct properties — rules.spec.http_method / d8920273d146 / 3

<a id="canonical-6e0d3ed0fe34351f9d005012c9a3df296bede49ff0a822da7008321d49436d82"></a>

<a id="canonical-5c70fabbdd7ad66fefae99b06809a92e9ed1f67b6ec763186c4054aebe7a44bb"></a>

## invert_matcher property — rules.spec.http_method / d8920273d146 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-ebd0add567dcd231d59e7550ddc4caab9af7d6456c06e45e600263269135a117"></a>

<a id="canonical-d45e8c27cd12d701c1ca9a4f1455bcbae5e1963e1747d17414cb10b93e35e4fb"></a>

## methods property — rules.spec.http_method / d8920273d146 / 5

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-8cab250eb85031aaad4c18fd321d41a4a1cc2f8ea4adba6aeb3aafa12d5ac673"></a>

## Next pages — rules.spec.http_method / d8920273d146 / 6

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-d5746e1b12da1121317c0a7ce76c63b6537db7d3ac8d2fc98d62c1291c669686"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-411ab0923ff3097c32422d5ba56a3620e56cfe69e2b1eb0b8d1388bb2ccac62c"></a>

## rules.spec.ip_matcher — rules.spec.ip_matcher / 06ca0ba0ad43 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.ip_matcher

<a id="canonical-dd2405117e2e1481d3156c991db002c19d105c51d46a651ea1106dcc88d82b06"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-b9f536371d85cc765d05d9d54229cef180fcd1ecdc79106b3054752cd5b3ae1b"></a>

## Direct properties — rules.spec.ip_matcher / 06ca0ba0ad43 / 3

<a id="canonical-ec4c3798d8505f85afe6a8dcfeb2c110eee7193864b658edc9fbdb39ddede8ab"></a>

<a id="canonical-633efc91dc803a198bc6dad405b2b74e6ac3a101dced5f9abf973710b74084ba"></a>

## invert_matcher property — rules.spec.ip_matcher / 06ca0ba0ad43 / 4

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-5a6e845e213cce22be5101c538644fd5ae17f7ef2e8032d898ac5ecb7db3256f): complete subsection reference.

<a id="canonical-cca11fbc71b77b75899ca9ab0a381cf3984a3a29addd381870e07d3b1572da24"></a>

## Next pages — rules.spec.ip_matcher / 06ca0ba0ad43 / 5

- [rules.spec.ip_matcher.prefix_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-5a6e845e213cce22be5101c538644fd5ae17f7ef2e8032d898ac5ecb7db3256f)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-5a6e845e213cce22be5101c538644fd5ae17f7ef2e8032d898ac5ecb7db3256f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cee0a1bf3c0f8ac59404b17fc9ff62d3c6fd621a58139f394c86875f96e14036"></a>

## rules.spec.ip_matcher.prefix_sets — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-d5746e1b12da1121317c0a7ce76c63b6537db7d3ac8d2fc98d62c1291c669686)
- rules.spec.ip_matcher.prefix_sets

<a id="canonical-c220bc84e50dc0481a98f0a97e737ea255fd86effbe1f6d07400bf9844ff29bd"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ceef4560c5734f55044fe4a3ad53796cede4b7fac0a525ed45f943d46db2319"></a>

## Direct properties — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 3

<a id="canonical-a9cab52b4d140fd0d510951090720f06998ad570b2a046bd278bb8eaeb5b7dde"></a>

<a id="canonical-1935242ef020e1b047466d8c5dce72a35e45e07692463ffaf0201b462bfff1cd"></a>

## kind property — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 4

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

<a id="canonical-3408ad5828adfecb7cd7a197bff258ff73a5c3175bcc4864f4b6698c228f6aa3"></a>

<a id="canonical-c5959d6f81ee35a9c4d9d32c135d410ad5c3645959e9ab91df4c42588b40b811"></a>

## name property — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 5

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

<a id="canonical-b01292bb9ba69aab6859e3615b9a046948eab7277678ade269f13a966bd2618c"></a>

<a id="canonical-a2ee2a5756ab659e54b1ffb98613d3e529f98010923b49955fd74a7533c11bc6"></a>

## namespace property — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 6

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

<a id="canonical-b7865ce258eab3866782b50aa4e19d7a700d49e3ab903c9cf9a3a00267b9b3a7"></a>

<a id="canonical-717b898746695300982f00081a38ed0b549d96944351e7da4f41970b9553a7ed"></a>

## tenant property — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 7

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

<a id="canonical-2e92b3dc69c42eeb4198e946abda79c8c9c8bdd50c8f1d769e0b4824c54ae00d"></a>

<a id="canonical-274c5df251c658dbad50c755402b8d50aafc9f431e1a9e95660d3c28758cc2b8"></a>

## uid property — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 8

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

<a id="canonical-cf402e2000c66657907bc143812334a7bd7b545dc09592e4055017c5fe53bf64"></a>

## Next pages — rules.spec.ip_matcher.prefix_sets / a2172cab935c / 9

- [rules.spec.ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-d5746e1b12da1121317c0a7ce76c63b6537db7d3ac8d2fc98d62c1291c669686)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-4a0d4d9a1df889b0ae2fe780c8e25e302c470b0e71cebb501890e3aa507f0970"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ab34e2f7966234c4659d3317b38e1bee2d3441188c0fbb7016aa60a72e268af"></a>

## rules.spec.ip_prefix_list — rules.spec.ip_prefix_list / 562fb2c90b51 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.ip_prefix_list

<a id="canonical-d95bd8df95cc76d2652840c5bac1b95b8323ddceb6f38cd5e8b71f572dd3f21e"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-193717bf047e6de45744776eefeeb5fb3392d26af268d3c10c0042c0fa8a466b"></a>

## Direct properties — rules.spec.ip_prefix_list / 562fb2c90b51 / 3

<a id="canonical-56001455a98d66b9917bed339d919d4af024efec1975528564a49a1a27a5cf51"></a>

<a id="canonical-0df2537b63ca5c218ee94747b86fbffc99e14f9e385fa87416d296af4abe5dc9"></a>

## invert_match property — rules.spec.ip_prefix_list / 562fb2c90b51 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-cd522a3e2a77e375ad71ac82ef51fec817ff52b10520d4a4cf4edca32a211e57"></a>

<a id="canonical-b810d6c3eb726c9bcb398b0c9ac5d95d7d573f8cef6b4689c6038785b203ee34"></a>

## ip_prefixes property — rules.spec.ip_prefix_list / 562fb2c90b51 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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

<a id="canonical-0f4e5e99d037fcb146804704e9ee2394b8da7e73fbb8743d7a665f29d8d3b09e"></a>

## Next pages — rules.spec.ip_prefix_list / 562fb2c90b51 / 6

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-9d8f11d8ce1ee7122604889577d8166719df102f378555ef3b2c8e65d46c3fb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41f030aa655536b0024cbc5d33fc0296c75e250551cd19149223399b967d7d30"></a>

## rules.spec.path — rules.spec.path / 6a0147c63f1f / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.path

<a id="canonical-d15878353d46fdaa60e87cc0e5cf470d8e416da048e9b46ed11a90169ff2e64c"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-d80fe19d237aa2eb63addbe67a45d1f5110b4954d992eb1b94705f61b653b3a3"></a>

## Direct properties — rules.spec.path / 6a0147c63f1f / 3

<a id="canonical-8ecbb81cdb0ce896d5df239b5a7e4ff6ce03da8476a4570ca9ed4adc10463dd0"></a>

<a id="canonical-081555f5e85cb0fc9f31a5b86a068601fec83d19cb4e44c9d018e37d1512fd58"></a>

## encoded_path_matcher property — rules.spec.path / 6a0147c63f1f / 4

Type: `"bool"`. Optional.

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

<a id="canonical-115ad9b6c35225208621773de12e29d58d82e3d61ac76a72f515812ef9a40c39"></a>

<a id="canonical-5f09896db10f00bf3756b4d49e6dde7ca521258fba38a9d281aed91c14d7dee0"></a>

## exact_values property — rules.spec.path / 6a0147c63f1f / 5

Type: `["list", "string"]`. Optional.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-231fe6b8cb279e5ca3c1ff9ea05ad0e968b9338ad7e1cb32e3a20a54e860746a"></a>

<a id="canonical-adad3823d1765503b94496c63d73f033bda67b7e8c09634be53b4458e927b186"></a>

## invert_matcher property — rules.spec.path / 6a0147c63f1f / 6

Type: `"bool"`. Optional.

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

<a id="canonical-856bfc5bf35473c2f543e5e8c0ac45ce257473885ce18312c528286ec4cfba99"></a>

<a id="canonical-40ef7bef8e954d57e7140c69162c5038a6caa4b1305df81edc49bf2d506c1420"></a>

## prefix_values property — rules.spec.path / 6a0147c63f1f / 7

Type: `["list", "string"]`. Optional.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-0c65eac1ddf00eb0f50ed24c8cb87475423f60d9d4f405d7a19489cd606d191b"></a>

<a id="canonical-84df71df7cbc0fc9df5c69c5c800965c6540e2a2184461303751d0725222c457"></a>

## regex_values property — rules.spec.path / 6a0147c63f1f / 8

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-5e9139945192e780e7330d5d10b0c8453a5cdd9c0be52a35bb31ac0fc99a76e3"></a>

<a id="canonical-0c5e443ed5fc819c27bde97c9f73107b5c2ef0285aaab20cab7a546c3ffc578b"></a>

## suffix_values property — rules.spec.path / 6a0147c63f1f / 9

Type: `["list", "string"]`. Optional.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-1e8803222f0b824f67b4a1fe4bf745cb71d31efa20a35d1b26e0668ffbe03a2e"></a>

<a id="canonical-70261c355e5e88f24dfa71bcca69331978f0d4ebc3253172e42b0db4ed82f58a"></a>

## transformers property — rules.spec.path / 6a0147c63f1f / 10

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-533321be627bfe0ae6db1e7c1ce9d8d3243090a4c563db86882b8d2084d3b3fa"></a>

## Next pages — rules.spec.path / 6a0147c63f1f / 11

- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-406858017b115eed88fc578f45bec3e898f378209c9e0d7209f27bd28e0954dd"></a>

## rules.spec.segment_policy — rules.spec.segment_policy / d83710f2cf68 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- rules.spec.segment_policy

<a id="canonical-781fd778c2cacb5613afc0ed675c94459bbb508504ceeefcd11c177aa56fa352"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dst_any",
    "dst_segments"),
  validators.ConflictingObjectAttributes("dst_any",
    "intra_segment"),
  validators.ConflictingObjectAttributes("dst_segments",
    "intra_segment"),
  validators.ConflictingObjectAttributes("src_any",
    "src_segments")}
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
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

Terraform syntax:

```terraform
segment_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-878dbc6e99f5f2f1885d25329542d2567cfb3470e45733ed50fe1aba9f53cc3e"></a>

## Direct properties — rules.spec.segment_policy / d83710f2cf68 / 3

- [dst_any](resources--rate_limiter_policy--reference--group-001.md#canonical-a9211447bacc30d92ad514eca3eb9dc32d01192c7b412d5f79d98054188ae1f6): complete subsection reference.

- [dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-fa814690238b49eaf28c847c8cca661ff09a7fef6b3f9233c26be7d163984913): complete subsection reference.

- [intra_segment](resources--rate_limiter_policy--reference--group-001.md#canonical-4d56abddcc378ab74501d96d260911f3979e4215c5ef40231fd7c2efa67a3d00): complete subsection reference.

- [src_any](resources--rate_limiter_policy--reference--group-001.md#canonical-78da93710f169ce915de496926601e9372965a825f7ef7c9b4acb755fcbd7eb8): complete subsection reference.

- [src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0ad1c964aa9286bb9ae4b7fe8d62c33f7ff10dee82797c142061c20ed1dec4fb): complete subsection reference.

<a id="canonical-a78d188ee2759b50bef9ca126895f2de63f3db1186dc47d239830555536ff79a"></a>

## Next pages — rules.spec.segment_policy / d83710f2cf68 / 4

- [rules.spec.segment_policy.dst_any](resources--rate_limiter_policy--reference--group-001.md#canonical-a9211447bacc30d92ad514eca3eb9dc32d01192c7b412d5f79d98054188ae1f6)
- [rules.spec.segment_policy.dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-fa814690238b49eaf28c847c8cca661ff09a7fef6b3f9233c26be7d163984913)
- [rules.spec.segment_policy.intra_segment](resources--rate_limiter_policy--reference--group-001.md#canonical-4d56abddcc378ab74501d96d260911f3979e4215c5ef40231fd7c2efa67a3d00)
- [rules.spec.segment_policy.src_any](resources--rate_limiter_policy--reference--group-001.md#canonical-78da93710f169ce915de496926601e9372965a825f7ef7c9b4acb755fcbd7eb8)
- [rules.spec.segment_policy.src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0ad1c964aa9286bb9ae4b7fe8d62c33f7ff10dee82797c142061c20ed1dec4fb)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-a9211447bacc30d92ad514eca3eb9dc32d01192c7b412d5f79d98054188ae1f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a3ee91e301fc69b03acb1473830d969e82295808736c809f675b18af8aca189"></a>

## rules.spec.segment_policy.dst_any — rules.spec.segment_policy.dst_any / 0686744f49f0 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- rules.spec.segment_policy.dst_any

<a id="canonical-739a91dd1c9814931d45ebdc0123ff121d327b8656f5511586a49272f11a0682"></a>

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
dst_any = {}
```

<a id="canonical-747827d611b609133cdc333526643e44a8ca27a046c014f639c2ba70ed325b7d"></a>

## Direct properties — rules.spec.segment_policy.dst_any / 0686744f49f0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f9a09378c248d1ae200a987ad260ffb7663569c61450ce8b60082dc21709ccc"></a>

## Next pages — rules.spec.segment_policy.dst_any / 0686744f49f0 / 4

- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-fa814690238b49eaf28c847c8cca661ff09a7fef6b3f9233c26be7d163984913"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f1493b9261f3cd8cbd9174642514ee45bb37512f3bba42c5a0604c1078ad9bf"></a>

## rules.spec.segment_policy.dst_segments — rules.spec.segment_policy.dst_segments / dd492db19a17 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- rules.spec.segment_policy.dst_segments

<a id="canonical-0e77093791998a02cfc1ec64be667266bfbf510a248a95a8383297f76888ad02"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
dst_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-c1b2ce1af8d2c1dd8f21914dc9d2eabf1b36c31e3bb08de7b52d8aeb06f7de13"></a>

## Direct properties — rules.spec.segment_policy.dst_segments / dd492db19a17 / 3

- [segments](resources--rate_limiter_policy--reference--group-001.md#canonical-1fdaa6c12f62065d150113767fd5561199e1b530d169cd73c110c16f04d1f40a): complete subsection reference.

<a id="canonical-40f2c81dcd22a0ac0d8b481262d32218155ddecccfe8106823202c7cc2438ffa"></a>

## Next pages — rules.spec.segment_policy.dst_segments / dd492db19a17 / 4

- [rules.spec.segment_policy.dst_segments.segments](resources--rate_limiter_policy--reference--group-001.md#canonical-1fdaa6c12f62065d150113767fd5561199e1b530d169cd73c110c16f04d1f40a)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-1fdaa6c12f62065d150113767fd5561199e1b530d169cd73c110c16f04d1f40a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a773dbfc71f3fa09712ecab06dbac1ae58a99f32c64ffcc2619c86952da97867"></a>

## rules.spec.segment_policy.dst_segments.segments — rules.spec.segment_policy.dst_segments.segments / a5cb6eeb15a8 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [rules.spec.segment_policy.dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-fa814690238b49eaf28c847c8cca661ff09a7fef6b3f9233c26be7d163984913)
- rules.spec.segment_policy.dst_segments.segments

<a id="canonical-9c5f01a5a645d91ffa78b6f962864f6935aa16b744254d7f8c2ce0ed0d79b9ba"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-0a22820466f3c7cb77d87dc3ab98e05ec29e5a991133fa9690cec241f5a84d01"></a>

## Direct properties — rules.spec.segment_policy.dst_segments.segments / a5cb6eeb15a8 / 3

<a id="canonical-9051f274165d37adec43a2b168bcf66ffc163bc22029c5f2e3242c3f98a4b1d2"></a>

<a id="canonical-b82f5fd7be270db862b75081df84a5c9b6286c5fb8479d8bd120a0ddca3c8a6d"></a>

## name property — rules.spec.segment_policy.dst_segments.segments / a5cb6eeb15a8 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-68998df1997b46fe64aa8a1941270057f11434c4bf58eeadcdcb653600e811ff"></a>

<a id="canonical-3128e926dceffeae7428dcb86ebe5f40db1c3e091d3276c54f01c7f09989626c"></a>

## namespace property — rules.spec.segment_policy.dst_segments.segments / a5cb6eeb15a8 / 5

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
}
```

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

<a id="canonical-231601d8704b1ca7963064ecf1c420fb67f895896b1f24f6424841106a123aeb"></a>

<a id="canonical-cccb0b8d850cf1e4e84117df46137118bc493b712e141b624cf75414f8337870"></a>

## tenant property — rules.spec.segment_policy.dst_segments.segments / a5cb6eeb15a8 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-a5b4d7b31c9671861be3c4b79f8f4b6d1c0839cacdbe53bae28d6c0443996c39"></a>

## Next pages — rules.spec.segment_policy.dst_segments.segments / a5cb6eeb15a8 / 7

- [rules.spec.segment_policy.dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-fa814690238b49eaf28c847c8cca661ff09a7fef6b3f9233c26be7d163984913)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-4d56abddcc378ab74501d96d260911f3979e4215c5ef40231fd7c2efa67a3d00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-585beb30d84de493bce533b181e8b9c7ba458ea371065267c48f5cb37ccd1e02"></a>

## rules.spec.segment_policy.intra_segment — rules.spec.segment_policy.intra_segment / 351659f3eb15 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- rules.spec.segment_policy.intra_segment

<a id="canonical-dfbde1bff989374e0597b32c6f9ddf0b40e0a32ec03d33c4f2a2ebdf3d0bb937"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
intra_segment = {}
```

<a id="canonical-3169523de4d1fb0ac881bb8c53782c6b854875898c726fbd37a2068ac135d06a"></a>

## Direct properties — rules.spec.segment_policy.intra_segment / 351659f3eb15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1f45b0b0f22f68b25594bce4ad29729f98e687fff534546f5d276627a647d48"></a>

## Next pages — rules.spec.segment_policy.intra_segment / 351659f3eb15 / 4

- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-78da93710f169ce915de496926601e9372965a825f7ef7c9b4acb755fcbd7eb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eaca9519a958d5b565f5690d4634d95cb39c990926db96a8dadc4672a12eed5"></a>

## rules.spec.segment_policy.src_any — rules.spec.segment_policy.src_any / c2c010bcf822 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- rules.spec.segment_policy.src_any

<a id="canonical-e589f7d37ae7a8350015177737bdaa094b0d96405ff8281726d0e0b4c02aef91"></a>

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
src_any = {}
```

<a id="canonical-b50d1d407e2dfd7001e6892fef8810f9a4e8cda1225e588b27f95549c2a30df5"></a>

## Direct properties — rules.spec.segment_policy.src_any / c2c010bcf822 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3ea001bc35a31c2093a9dad04d39e8981921923e1a5cd7552caad0d6b1cf24a"></a>

## Next pages — rules.spec.segment_policy.src_any / c2c010bcf822 / 4

- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-0ad1c964aa9286bb9ae4b7fe8d62c33f7ff10dee82797c142061c20ed1dec4fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d1cd049b835d84bc6c84095293542643b4040bbec6e0c249831ee77d157d869"></a>

## rules.spec.segment_policy.src_segments — rules.spec.segment_policy.src_segments / 4e0e42152e94 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- rules.spec.segment_policy.src_segments

<a id="canonical-014adb043eff200a4110234c5eac80becc1abe5697e96a048488c2811ff9ea69"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
src_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-f570780b6f3dffed6e84fda2c9e99b0c9f2bd17184321df0f84b195055fb3ed7"></a>

## Direct properties — rules.spec.segment_policy.src_segments / 4e0e42152e94 / 3

- [segments](resources--rate_limiter_policy--reference--group-001.md#canonical-bfdec0745c2ce47268f52623529aa8e40a1138e81f5131f3b1d29f0061a9efeb): complete subsection reference.

<a id="canonical-e1db9bd663e97f07cc133a579ebf47a98bb874dea53d466a78b234acc2b3cf76"></a>

## Next pages — rules.spec.segment_policy.src_segments / 4e0e42152e94 / 4

- [rules.spec.segment_policy.src_segments.segments](resources--rate_limiter_policy--reference--group-001.md#canonical-bfdec0745c2ce47268f52623529aa8e40a1138e81f5131f3b1d29f0061a9efeb)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-bfdec0745c2ce47268f52623529aa8e40a1138e81f5131f3b1d29f0061a9efeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c8c3d20c1b3f29779fafd5f39f08f8e6013ea9e33597b774e881d65ac1177f0"></a>

## rules.spec.segment_policy.src_segments.segments — rules.spec.segment_policy.src_segments.segments / 9cc2c24659ba / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-aa7701ea8c21daf9315a908b937c206522e47231237f100f4f082156c9dd23f3)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1e887952bc93507838e56084fa54673b060cf65765337dfa8c282fe88f39ed8d)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-43c357549a29a9032f3adcb6180de01fbae977c3baf5da9dd62ee06f0ca4a076)
- [rules.spec.segment_policy.src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0ad1c964aa9286bb9ae4b7fe8d62c33f7ff10dee82797c142061c20ed1dec4fb)
- rules.spec.segment_policy.src_segments.segments

<a id="canonical-f42a3b5879455b4c3f9b07e561f446acffb70f0421b9892b148c55554656fdba"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-8dfecf50aa040afdc1a86fadf64aec3eb62ec53c922238460b3831b4f53b1020"></a>

## Direct properties — rules.spec.segment_policy.src_segments.segments / 9cc2c24659ba / 3

<a id="canonical-09a4568eebdc7063dcae8d9065ea9b4b49ffa6d792b6b1b9fa879653927ece27"></a>

<a id="canonical-708353a6c8bc2319897f7d6c66a8bf38169ba0f23e4ff8889a44a9ff3abb8789"></a>

## name property — rules.spec.segment_policy.src_segments.segments / 9cc2c24659ba / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-8b4777d17b6eba8e569794489073d3c956af7c3e534806017af84891c914d838"></a>

<a id="canonical-b61f1ee2f2d92e6023f8764e6d75c0ba62c9369b67e93b020b2fd7c011deb1c2"></a>

## namespace property — rules.spec.segment_policy.src_segments.segments / 9cc2c24659ba / 5

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
}
```

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

<a id="canonical-2d44c9ff909f35acf43dc7b30f55d2323664db04c6d24b4ff29462dee50420ac"></a>

<a id="canonical-115a09bf8b0b500e65bb5268ac3fd61c04d5300f12eceea02cdefeeb05098b84"></a>

## tenant property — rules.spec.segment_policy.src_segments.segments / 9cc2c24659ba / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-5d360bcc26e6cae38dd7476baa47bb746fe5855185e06305b68445ecf20ff0f0"></a>

## Next pages — rules.spec.segment_policy.src_segments.segments / 9cc2c24659ba / 7

- [rules.spec.segment_policy.src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0ad1c964aa9286bb9ae4b7fe8d62c33f7ff10dee82797c142061c20ed1dec4fb)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-3e759f0ab8dc83a52fb91efa326ee68914c628ce86919f316e19a1c4368373a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0741d4d8eb12efe304a0318f67ac10521798f327551b02c54f9ac50f16aaf8eb"></a>

## server_name_matcher — server_name_matcher / bb79bb1f2aa0 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- server_name_matcher

<a id="canonical-0147585cf633a964d017070090c9913886726c790787618a9b68a08dcfde66be"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
server_name_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5129445ab4ae621adc105eb8c18ac8e4894deffb26dfd648d581b73ee6f9073"></a>

## Direct properties — server_name_matcher / bb79bb1f2aa0 / 3

<a id="canonical-15169173abd7dea0b9a1e59c12c1c196293c265a6f742072d41e3bad5652ad3f"></a>

<a id="canonical-101819eb00ba0a4f1091191776092124cf801ae2c127112ff2c2f9b5ca3f5784"></a>

## exact_values property — server_name_matcher / bb79bb1f2aa0 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-c0cbadb432d16ca9bfe413e49bdc1056fad3ab2a57d12b5999e914e7a93f3290"></a>

<a id="canonical-68b3598631d423f7ab61d0a638aed9adfdf32a622f5c4a63efad2c494a6d677a"></a>

## regex_values property — server_name_matcher / bb79bb1f2aa0 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-0dd8bf84ec6370b3eabef35e2f8677e234cb103ca73049fd3121d8cb6726ebe4"></a>

## Next pages — server_name_matcher / bb79bb1f2aa0 / 6

- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-a76ecf72963cf5b3f31077a4892876c381eb73eebf600858b6dd88334e80a6c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3620eaf03a900f70c2bf85838628e61221bd37b99f60fc23b18a7e1d14942272"></a>

## server_selector — server_selector / 3909e049ea66 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- server_selector

<a id="canonical-ffa987c4a6cd7dd45162b28bcff8f3f61c7eea22399e91b5f5c65e51281ab0ff"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
server_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-03195de478f438aa223571cbc3c4074a606b6c197bcd60379872230b67733d16"></a>

## Direct properties — server_selector / 3909e049ea66 / 3

<a id="canonical-7531f93e5218c6ef5cfc94dec0a8e35941c972f5e63f798b098cdfed6d53ac5c"></a>

<a id="canonical-6dd1c1622eeb0d40e5ea87136797090987f2ff16ab4c8ec699b10dbfaa100895"></a>

## expressions property — server_selector / 3909e049ea66 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-c0f9540fba47edce1b92ef25ddee849c7e604142bb738d635124415ad9785400"></a>

## Next pages — server_selector / 3909e049ea66 / 5

- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-596b33ac3e31c789c6ff45eb3514ad49a37cf7f206057468810e3be146744ecf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e9a9b4546f3826f73a8bcf3db52c7e139cee40661ccda4b66ee47947ad5b09c"></a>

## timeouts — timeouts / db58ab44f225 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- timeouts

<a id="canonical-e64fdbe480ab2dce371d1204d4479595cd2d309062fa091ee6e9a42651337c11"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-c41c92a094df12c7a72acfd53f01fd24b6af1763c12c8d1d6843d3a3594adc90"></a>

## Direct properties — timeouts / db58ab44f225 / 3

<a id="canonical-ceb26058658ae462a01806a290bb8a27599deb7100576ccc23db2a85de7f6008"></a>

<a id="canonical-8dab9b84f0d27fee8d62a741cbf111c458600026e651c659dce75e7d8e786668"></a>

## create property — timeouts / db58ab44f225 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0cf6abf7c561d284351fc50b203b71843e267b1383eb513d8fab93375be4d1e0"></a>

<a id="canonical-183a7e37b1a7b4ad332e5105f38349c1a6c47e6285a999fddb101af87d1506b5"></a>

## delete property — timeouts / db58ab44f225 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-b843b0c266f574d3838f340f6f2c242ff51ffc007cbd8813c1d662ea27bc4701"></a>

<a id="canonical-eea69cb41503648901946329932622471327333c630a3ae1f91149c21b33f4b1"></a>

## read property — timeouts / db58ab44f225 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a132780cf5d0c457c55bf71df8e68eeebef895f33ee0af1aee5684d302cf14da"></a>

<a id="canonical-59da3b96e3bf05fdb608f6ff6b88d42a6fb6c8d37ad8403d47b88d6ebee03892"></a>

## update property — timeouts / db58ab44f225 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-90a2d466dfc0b96c21e7b099410b2fbe1f4111d00773d44b1af34f8d37c484d3"></a>

## Next pages — timeouts / db58ab44f225 / 8

- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-3009b5c32e16c1cc4a64b67e74efcc24077dd4d05d49fe3fd30373d616c8f64b)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
