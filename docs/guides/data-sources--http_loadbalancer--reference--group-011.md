---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-bc07bc176d105baa735b4529c461c301c07d7962d36c800a7c4eed918f6e6536"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.path — bot_defense.policy.js_insertion_rules.exclude_list.path / 0c60077d9b25 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-b473a9279ba92c3e3089600693d3acea482dbf183ca47a55a3dc6008e1025e80"></a>

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

<a id="canonical-2a45510771ff9f8a40e1593242d3f5300d6654b4c88e32a4389db500e094df71"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.path / 0c60077d9b25 / 3

<a id="canonical-9e9e7ceb6da4420195955728c7c6f1b6ce62b6c6a18774b4b505a118edd0d32f"></a>

<a id="canonical-68190abf86e0e8eb970268977a91940a23f15f6a032575097c576ebfafe23675"></a>

## path property — bot_defense.policy.js_insertion_rules.exclude_list.path / 0c60077d9b25 / 4

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

<a id="canonical-e91eb6704cabdf8a9e6c57d5985ba623f884a555345ceb0d8eed1339109fb525"></a>

<a id="canonical-1f7503f07b4b43a28d0904e9fb504173f5333eb280c224fd684e71c3ca3e88fc"></a>

## prefix property — bot_defense.policy.js_insertion_rules.exclude_list.path / 0c60077d9b25 / 5

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

<a id="canonical-9ba3fcecd7263cf49ef62552ade5a68689b7d80533de9c9f944df4eea1e62cfd"></a>

<a id="canonical-649431339b2d24b1691e5758521cdc704bbc43c18fdabb7d5b1699f770844399"></a>

## regex property — bot_defense.policy.js_insertion_rules.exclude_list.path / 0c60077d9b25 / 6

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

<a id="canonical-06ac73c9a17e10a770c1dbf654f18ed1e290e5dce7a1dfddd26716923d696ac9"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.path / 0c60077d9b25 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69367a545e09310dc6b00a31e0305184fdd825ae8a6dc6b890dc0988516e0e7e"></a>

## bot_defense.policy.js_insertion_rules.rules — bot_defense.policy.js_insertion_rules.rules / e44173411e3f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-7bff8c3ea6461cffb2b5a4da02b93fc2f2f5441fd1e754f0cdacd148d9835441"></a>

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

<a id="canonical-ee8946bb12c724d34c0a76de1eefc1fea2c5a987edc34c4afdea7d6fb2387870"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules / e44173411e3f / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-7efb92ccccf9df1deb3ab8bb65eda01296bb698a3d8787e86f10c58f460e8490): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-6adc29c58596a416b1669103f67e39e010cef23ebc2e01eaff2aeb71122e3054): complete subsection reference.

<a id="canonical-316111709a552aaa64f5827c61d447a41e40958a0d2f47e9096ca32dc40d86a3"></a>

<a id="canonical-befef78d3b91e793d4ac6e77d443c081402fb9870c958aa59abc40e7dc08599b"></a>

## javascript_location property — bot_defense.policy.js_insertion_rules.rules / e44173411e3f / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
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

- [metadata](data-sources--http_loadbalancer--reference--group-011.md#canonical-16049955386160d8a7713459b7681ff5b5cfe36dd2653b980febc153b4f8fdfe): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-011.md#canonical-2f58ba22050a273bb1e6499f4d24279469287bbdae90d39937c03345e3147bbc): complete subsection reference.

<a id="canonical-539aa1c6984a1d68d25be3947673e7613206bed660eb00a538551f247cc40d5e"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules / e44173411e3f / 5

- [bot_defense.policy.js_insertion_rules.rules.any_domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-7efb92ccccf9df1deb3ab8bb65eda01296bb698a3d8787e86f10c58f460e8490)
- [bot_defense.policy.js_insertion_rules.rules.domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-6adc29c58596a416b1669103f67e39e010cef23ebc2e01eaff2aeb71122e3054)
- [bot_defense.policy.js_insertion_rules.rules.metadata](data-sources--http_loadbalancer--reference--group-011.md#canonical-16049955386160d8a7713459b7681ff5b5cfe36dd2653b980febc153b4f8fdfe)
- [bot_defense.policy.js_insertion_rules.rules.path](data-sources--http_loadbalancer--reference--group-011.md#canonical-2f58ba22050a273bb1e6499f4d24279469287bbdae90d39937c03345e3147bbc)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7efb92ccccf9df1deb3ab8bb65eda01296bb698a3d8787e86f10c58f460e8490"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd9f113629172c8a1458dc2582c878287dc096b8e4e81d579d1c85e3950c19ea"></a>

## bot_defense.policy.js_insertion_rules.rules.any_domain — bot_defense.policy.js_insertion_rules.rules.any_domain / b899b7c10bdb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-667e70223af1e37c56cfd4b06858851356063d6cd3fa44a2423fdd80601f4192"></a>

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

<a id="canonical-242f321ecfbd6bc82646475ce6a51fbda3350f2f9998d1ebf469c0fe024deb95"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.any_domain / b899b7c10bdb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1b4e55f5dfa7d7f03198d5394bab0e32f6c2a0eeeb31d72ff265b1210a2b112"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.any_domain / b899b7c10bdb / 4

- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6adc29c58596a416b1669103f67e39e010cef23ebc2e01eaff2aeb71122e3054"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f56c169c97a192089b1d7b03024c606de65e5857bebdc00556a5113e7d6d257"></a>

## bot_defense.policy.js_insertion_rules.rules.domain — bot_defense.policy.js_insertion_rules.rules.domain / b4c95745bdb7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-a0f5c1b7c024849aac50a86fbc1984021dfd47f0b4430ab91132c6d51ca93c27"></a>

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

<a id="canonical-7b2e609b8f3d01d92a48846b1a3aa1248e7ec795ff9a1d3ca44b57ca657a56e3"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.domain / b4c95745bdb7 / 3

<a id="canonical-ad9207dab97a32e56a1443f949ccc2a0f5c2fb3a1f69133d8de9bd109f67c1e1"></a>

<a id="canonical-00815c35c9155268b248ba0c2f60995c94e645e70e50bf2d8bda8911bc2e5793"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.rules.domain / b4c95745bdb7 / 4

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

<a id="canonical-aa8faddb37a7273a6638da88baac45e6a507f5607dbd3cd947e5194ed6f45fb9"></a>

<a id="canonical-0398bb1bf40bfac5537258ace122906d77da82d73747f3305d37da13a338547e"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.rules.domain / b4c95745bdb7 / 5

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

<a id="canonical-cdef33ca108e87467641c84fada0fbd0609f3c53e3af97942603f9f388dbd2e3"></a>

<a id="canonical-6694a54d4c41f44f68a0aa3fb74be86f00132c7d7ddbd02eeecf31413b066d4e"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.rules.domain / b4c95745bdb7 / 6

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

<a id="canonical-576ae4a052d7f6a98ba7a7618bf1af39379413c03652199be6024ace1b3f8728"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.domain / b4c95745bdb7 / 7

- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-16049955386160d8a7713459b7681ff5b5cfe36dd2653b980febc153b4f8fdfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d63d839ee9d4d4f05ae1e7af68c1155ce1c753b7947535545da8de0d5e26cf0"></a>

## bot_defense.policy.js_insertion_rules.rules.metadata — bot_defense.policy.js_insertion_rules.rules.metadata / 17f1d90ffcef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-5626649cb2f5a64c1b462b143864e5829cee83f06ba35179eb7bd316e629b521"></a>

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

<a id="canonical-48040fddcb2dce6f2c8c87163cb349228acc7bd722a0b4adf507e9eaf941e806"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.metadata / 17f1d90ffcef / 3

<a id="canonical-87a4c1b2c3ceee7ec6ed6e60c52f8bf37694a46461aaec59dcb1d13926a5e630"></a>

<a id="canonical-d345c19bc5b41400c3f3cc2dea55a9313cf120dfde0763ac9182c604b5712ec8"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.rules.metadata / 17f1d90ffcef / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-19c63452f967b5ed10d9d7697a18d65d9cabd0e141f08d34eedeebc543d8415b"></a>

<a id="canonical-fea103a3e95f9f278a03e6381a1c550cbd37295cf404078bf71be0b4bf5e9341"></a>

## name property — bot_defense.policy.js_insertion_rules.rules.metadata / 17f1d90ffcef / 5

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

<a id="canonical-b38e4a295b91e735f5a82f2f21d9ac3fd7f658bba0f8a234885fe4005259935c"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.metadata / 17f1d90ffcef / 6

- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2f58ba22050a273bb1e6499f4d24279469287bbdae90d39937c03345e3147bbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0db8af2e1e604b6b8355fa632f716c03aa4fa9c7619918f7f69282409e8b9db6"></a>

## bot_defense.policy.js_insertion_rules.rules.path — bot_defense.policy.js_insertion_rules.rules.path / 235aea23719f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-0b6549bb6e3dcb664f590ed50fcce0890d62483658accfa63ca474cc9cec0ba6"></a>

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

<a id="canonical-0765bd731aa45c73e2b8c458caabfea4ef9e271603fb22094e9099c5b548499c"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.path / 235aea23719f / 3

<a id="canonical-e7cda29708d3214b666a1cbdff781bcaef556067506c46bff10bcbb7bab4e92a"></a>

<a id="canonical-806cbd472fde559f0c821b93aaba400300936da106f599decef03570c8ae652d"></a>

## path property — bot_defense.policy.js_insertion_rules.rules.path / 235aea23719f / 4

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

<a id="canonical-d473bb78f1595ddd215db6ed8ea506b32ec489c1c4b3e347b9cfc8a155582f3c"></a>

<a id="canonical-47d465d80c7f9f3ab4f17543b6a6839533ad182e889eee18acb5780015abae2d"></a>

## prefix property — bot_defense.policy.js_insertion_rules.rules.path / 235aea23719f / 5

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

<a id="canonical-4b46f53b67f2d69ecb641de50b038f9d13b0a933ce8c163008d905730aa272b1"></a>

<a id="canonical-68fcc1b73bf80b80e107bef700d572291fb62ebaf7ecd24ba139b428cea8a3f7"></a>

## regex property — bot_defense.policy.js_insertion_rules.rules.path / 235aea23719f / 6

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

<a id="canonical-1a7e0b51800856e0581e473da21eb437018ff2722a64b83f76a24725283f7df3"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.path / 235aea23719f / 7

- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a9b7464ddb126cccea594228e35443c197be27c70b58a1799cf864601e03a68"></a>

## bot_defense.policy.mobile_sdk_config — bot_defense.policy.mobile_sdk_config / 124a3a5b0cc9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-7727fa2cf98617c5ed26c7f93854e447e0eee4990d9a873c5f5140aec67d8026"></a>

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

<a id="canonical-1b54eb607430f507def70fd27874ae808d45200aaef3abe7f8bae1a6d101e107"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config / 124a3a5b0cc9 / 3

- [mobile_identifier](data-sources--http_loadbalancer--reference--group-011.md#canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703): complete subsection reference.

<a id="canonical-d79b8de13b6fb8fb487f78e29462fda77782b354e3da4ba586632607aa151440"></a>

## Next pages — bot_defense.policy.mobile_sdk_config / 124a3a5b0cc9 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-011.md#canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5c967c58c59016caaca7f28478f105abf4cf8ad99fff6d15407e832a6a549cc"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier — bot_defense.policy.mobile_sdk_config.mobile_identifier / 37b7445dbcd2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-b6e45bf21d7215a573247b99b8eb14e2cc52d7632da47dc884d6ec33167f0a0a"></a>

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

<a id="canonical-47d96579e2dc9fe6848516923b0d0f75e773618b13bdf9ebffc62fe6235d5ee4"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier / 37b7445dbcd2 / 3

- [headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b): complete subsection reference.

<a id="canonical-aed5e887b63eed574841a7b47c30d1e73a4562bc69274fd9c2dbc7d4764bd228"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier / 37b7445dbcd2 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afb58c770abec421b4d2f626b2a6e86f2201ed8a45ca6b55287baf733345a16a"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3a9b68127ee2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-011.md#canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-f192132eb2754e4e27264d075e1a168c09bc25dd32a92a2587da976498f8e8ba"></a>

Type: `"list"`. Computed.

Headers that can be used to identify mobile traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-896c6dc19b18ecb6f4282c8ae6942a1a53d42b0ace5cc7c22a5d211f897c002d"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3a9b68127ee2 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-011.md#canonical-c0f7a6ea4707bd1f2a60c7617b6f3cd6739db125520163c2f1d472291daab2b6): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-011.md#canonical-72d11f59302d7069ee8b19f593b147ee4bba3d8498aac7144d69f9f04b6e8be3): complete subsection reference.

- [item](data-sources--http_loadbalancer--reference--group-011.md#canonical-2760fd053606a8d6d01d688ab69d618fee297e0769833d54db6bcf3b49fa4971): complete subsection reference.

<a id="canonical-c1cede6e90b90a185115abbf4f53112390c72d52882166bd3de9a969a24a6941"></a>

<a id="canonical-bb7b5295ceebebcd45dfe8dbd5c6e47a53f94f4fcf0f4e8af1f0b9ab1c7fadb0"></a>

## name property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3a9b68127ee2 / 4

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

<a id="canonical-0768347ec6fba8cee07de25ec18be40b34878e4da4aa1709f8722e335a9e8663"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / 3a9b68127ee2 / 5

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present](data-sources--http_loadbalancer--reference--group-011.md#canonical-c0f7a6ea4707bd1f2a60c7617b6f3cd6739db125520163c2f1d472291daab2b6)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present](data-sources--http_loadbalancer--reference--group-011.md#canonical-72d11f59302d7069ee8b19f593b147ee4bba3d8498aac7144d69f9f04b6e8be3)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item](data-sources--http_loadbalancer--reference--group-011.md#canonical-2760fd053606a8d6d01d688ab69d618fee297e0769833d54db6bcf3b49fa4971)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-011.md#canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c0f7a6ea4707bd1f2a60c7617b6f3cd6739db125520163c2f1d472291daab2b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf4635af8995a884ae31ee5798b1e222d5ff37ba97166194fcdc084dcd93c260"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 1b560ab40198 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-011.md#canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-9d9edf79d6794b9bbf32d3f02ac27620e5be7347dc83dfb5e62365517aff99f8"></a>

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

<a id="canonical-76e89ab479a03acf4576789b6dedfa7dc95a9ec8ddacc93daef1fa1a89939abb"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 1b560ab40198 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9bbaf79084a30a04c426759da2fb1e2dda8b613bf2ee7b67f06bb30222c9fe8"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 1b560ab40198 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-72d11f59302d7069ee8b19f593b147ee4bba3d8498aac7144d69f9f04b6e8be3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d944601b96b48f286c9e2e1155d30ac918dc9f4e49b63b0a29c1f2361f8e84f1"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 12cea4baa877 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-011.md#canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-24ee5ab5c55722c90fa1e91f035cb6399cac33381ee926d79265d0fd73516f71"></a>

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

<a id="canonical-1a7955f0ee86bfcbe61e85f8d7c93deaf29f57e5f964e7f656766e403759cc81"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 12cea4baa877 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1e1643ebd197760efb3a24408a5aee2d09921632c55eaee29a191554a9c2f66"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / 12cea4baa877 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2760fd053606a8d6d01d688ab69d618fee297e0769833d54db6bcf3b49fa4971"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88934150a62d820d1af0779330c4bc35bd6a30855168e77c5adf5c4dc9528b8c"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 70e20ce76219 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-011.md#canonical-cb732e001668b47c807655725c94e19d571443179d93f974205cd653c4259703)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-623a49bcbcd635317e9e77fe9e900a47f738217c331f3567b3089209687ed30f"></a>

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

<a id="canonical-473301940f31af9dda807d6623ab4e7ce9eda8ddd9004ec06e8719ae5aec2e72"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 70e20ce76219 / 3

<a id="canonical-56f6ee49f62b402fc7a95e9ca8ea8aa3610d825cd2c9b1bb2d1c52032a2380ad"></a>

<a id="canonical-ac9924e0fbe5bc70174a2a0d0b15bf3751ca749a1c0774eefd429a3b7d2ccb8e"></a>

## exact_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 70e20ce76219 / 4

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

<a id="canonical-dab49151b9013781ad552519273843176570bbc56fb33c29fbaa5f924a4fbcbd"></a>

<a id="canonical-5d209b6e75852b63eb6f4d23f6388ca5a290a449afc035683e394117a44e85d4"></a>

## regex_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 70e20ce76219 / 5

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

<a id="canonical-acd59d3f9f02e8b9ff667f9712e4f9f04e35882f4b6acafae78675a9c31ea418"></a>

<a id="canonical-91de94a720da6b8fc452c1c3f19bea0f759699642f62c1e7d56678c023817951"></a>

## transformers property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 70e20ce76219 / 6

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

<a id="canonical-20e1affb01524df9e6a8368a8bc8ffe5a6aa7424b3bffa0d0ff8e4933284bf35"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / 70e20ce76219 / 7

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-011.md#canonical-e04a70061d9366f61214ec4cb38bec8e1639f9a9499cf658850a570a5f71fe6b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8171696d06761e6b7066560c16aca36c5f5ce9a35e6fad3b7119b0dc143d45a"></a>

## bot_defense.policy.protected_app_endpoints — bot_defense.policy.protected_app_endpoints / 65326cb3c6db / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-43f5d702c2d2290a3a5e69e5bc07df249ee6077eab416d4aadf0ef98c21204d2"></a>

Type: `"list"`. Computed.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

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

<a id="canonical-d0e428aa20de2f432e1ef3a4683d96de73a7c50ca7c0b6996f7d8ee5a1df9275"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints / 65326cb3c6db / 3

- [allow_good_bots](data-sources--http_loadbalancer--reference--group-011.md#canonical-e7c6cea2552e0ac1a7da4ae242eb4ad2649dfcbdb2cb5d53cd42676e885ce8f3): complete subsection reference.

- [any_domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-8b85849d665303bf0a0ebf7af8b40316c0b4cc906417ed6e8cc2089b449d8478): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-1e5f0b510b9da14a66f5f2bfb979d6b28ee04aaf616dbbdce0cf164008ed98f8): complete subsection reference.

- [flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-012.md#canonical-456066ae1ec37d2ff34d06c59f71438b3acfc00de7c67030d34d2a38469c7b20): complete subsection reference.

<a id="canonical-cd7ba06dd507608849de295096167d5bbe007f2a34fb8da66e45658086489c67"></a>

<a id="canonical-82fe12e3808f6bd6aa0281f59a279942100a6646e42752a7cf09daaeddb549ed"></a>

## http_methods property — bot_defense.policy.protected_app_endpoints / 65326cb3c6db / 4

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
  "maxItems": 5,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-012.md#canonical-2d67645af0f8346adbd8fee782a2fa4ccfab5e057f484738345f25f6f1473bff): complete subsection reference.

- [mitigate_good_bots](data-sources--http_loadbalancer--reference--group-012.md#canonical-28c6b3cf79c8ca304614ab122ceb15fdfbc1558ad66d8feb13f708c1519c5429): complete subsection reference.

- [mitigation](data-sources--http_loadbalancer--reference--group-012.md#canonical-d7641e888b47263fb6e1f0cfdd38fa1e43d7f2024d548f8a116092fc4efc8a60): complete subsection reference.

- [mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-7c69c3a89c3972d763e6f5daba5dc8e4dfc782804050ee29cb8a708c7826e5da): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-012.md#canonical-bcab9662ee76b16e1ffc96170988b3218f721350f0ce386259b5a35471e4552e): complete subsection reference.

<a id="canonical-4ce75a37a0655c207231fffb7c116033f27d07049d386fc350df930f36ca3076"></a>

<a id="canonical-70c2c97e055c368c3303bc310581aa893e776d037243effc9c6c40bfd10dc6cb"></a>

## protocol property — bot_defense.policy.protected_app_endpoints / 65326cb3c6db / 5

Type: `"string"`. Computed.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](data-sources--http_loadbalancer--reference--group-012.md#canonical-1e7c9504396fb479a5e8e25024bd56636f014fba14e2613ffc3dad7edf11c937): complete subsection reference.

- [undefined_flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-06fd9f0f0798df46669ecc99fec4da356b6d23e56ecbc2b990e1712a47635ae0): complete subsection reference.

- [web](data-sources--http_loadbalancer--reference--group-012.md#canonical-aec11cae23598ae8b46c6e38b65ffb3edaa113d25e6fbdcba2b5fc286ad9773c): complete subsection reference.

- [web_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-dbb528cb4cc1701263144ad68be9c472d2a2e3d53340731280615e7e06906063): complete subsection reference.

<a id="canonical-953b9cb94fc78127f25f38a32c5db51d98df0f75490b90b7514099210a1eacd9"></a>

## Next pages — bot_defense.policy.protected_app_endpoints / 65326cb3c6db / 6

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](data-sources--http_loadbalancer--reference--group-011.md#canonical-e7c6cea2552e0ac1a7da4ae242eb4ad2649dfcbdb2cb5d53cd42676e885ce8f3)
- [bot_defense.policy.protected_app_endpoints.any_domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-8b85849d665303bf0a0ebf7af8b40316c0b4cc906417ed6e8cc2089b449d8478)
- [bot_defense.policy.protected_app_endpoints.domain](data-sources--http_loadbalancer--reference--group-011.md#canonical-1e5f0b510b9da14a66f5f2bfb979d6b28ee04aaf616dbbdce0cf164008ed98f8)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.headers](data-sources--http_loadbalancer--reference--group-012.md#canonical-456066ae1ec37d2ff34d06c59f71438b3acfc00de7c67030d34d2a38469c7b20)
- [bot_defense.policy.protected_app_endpoints.metadata](data-sources--http_loadbalancer--reference--group-012.md#canonical-2d67645af0f8346adbd8fee782a2fa4ccfab5e057f484738345f25f6f1473bff)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](data-sources--http_loadbalancer--reference--group-012.md#canonical-28c6b3cf79c8ca304614ab122ceb15fdfbc1558ad66d8feb13f708c1519c5429)
- [bot_defense.policy.protected_app_endpoints.mitigation](data-sources--http_loadbalancer--reference--group-012.md#canonical-d7641e888b47263fb6e1f0cfdd38fa1e43d7f2024d548f8a116092fc4efc8a60)
- [bot_defense.policy.protected_app_endpoints.mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-7c69c3a89c3972d763e6f5daba5dc8e4dfc782804050ee29cb8a708c7826e5da)
- [bot_defense.policy.protected_app_endpoints.path](data-sources--http_loadbalancer--reference--group-012.md#canonical-bcab9662ee76b16e1ffc96170988b3218f721350f0ce386259b5a35471e4552e)
- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--http_loadbalancer--reference--group-012.md#canonical-1e7c9504396fb479a5e8e25024bd56636f014fba14e2613ffc3dad7edf11c937)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-06fd9f0f0798df46669ecc99fec4da356b6d23e56ecbc2b990e1712a47635ae0)
- [bot_defense.policy.protected_app_endpoints.web](data-sources--http_loadbalancer--reference--group-012.md#canonical-aec11cae23598ae8b46c6e38b65ffb3edaa113d25e6fbdcba2b5fc286ad9773c)
- [bot_defense.policy.protected_app_endpoints.web_mobile](data-sources--http_loadbalancer--reference--group-012.md#canonical-dbb528cb4cc1701263144ad68be9c472d2a2e3d53340731280615e7e06906063)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e7c6cea2552e0ac1a7da4ae242eb4ad2649dfcbdb2cb5d53cd42676e885ce8f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef96749b32ca71ba2ba48de81f4c6414828643fa8ab27ea6d4792051c45ee6d3"></a>

## bot_defense.policy.protected_app_endpoints.allow_good_bots — bot_defense.policy.protected_app_endpoints.allow_good_bots / 937591c474be / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-f2993d9ae6b6e62a0ffac1a945d75901d48d3fe1789f31ae62851bbef59058d3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow good bots.

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

<a id="canonical-e9159f5e2a80d07d4159a3b522a59874157e27a652c87c14c2c82fa02faba44a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.allow_good_bots / 937591c474be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd4ade00a3eda0bef86ce707daef6f64ff908785d3dda747197e4bb70c05f424"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.allow_good_bots / 937591c474be / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8b85849d665303bf0a0ebf7af8b40316c0b4cc906417ed6e8cc2089b449d8478"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10201fa7fe2af1d530ac881fc80eeec5705d5f2dda603cbfd3d49f54d08d5d2b"></a>

## bot_defense.policy.protected_app_endpoints.any_domain — bot_defense.policy.protected_app_endpoints.any_domain / d58dce8c2a25 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-220dc62ce9fd1301e793420362bc8b8e74d10c15bbed68947d4e430766bf0057"></a>

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

<a id="canonical-082d4e5fc7a7bd946ac1f308d8a16519fae2dd26489a9d629ef11644b32dc9f2"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.any_domain / d58dce8c2a25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24f684dee5be160b0051c8f4b9f97ef29bd28aae3c1e42f15cf73a94ee070938"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.any_domain / d58dce8c2a25 / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1e5f0b510b9da14a66f5f2bfb979d6b28ee04aaf616dbbdce0cf164008ed98f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c0338677e72689f3c886bbf961d23a4387993f57ecb57bf7d68db96e6ef1fcd"></a>

## bot_defense.policy.protected_app_endpoints.domain — bot_defense.policy.protected_app_endpoints.domain / 540553ff393a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-3dbdcf9709206993ed61f38550868c2cd116137375e5fb33efc3ae1909d425ab"></a>

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

<a id="canonical-e0d6b4dd44b0f3fff0b157cc32bf62c828ed11c7380eac9ffb39bd8f8ae8f515"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.domain / 540553ff393a / 3

<a id="canonical-348fe408ef8b748c62ca3a9e0368f4afef5b93f57ef01c1658fafb9742a9e280"></a>

<a id="canonical-1182161cd9817f19b792586459f71e29a435ef69c289ee9d0ca6afdba4da6e16"></a>

## exact_value property — bot_defense.policy.protected_app_endpoints.domain / 540553ff393a / 4

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

<a id="canonical-ebd27182c8e6053fa619b3f5330d7c7364324b3010f247ac3eb849825dc6a64c"></a>

<a id="canonical-6e162bc38f4e8ce3e347b374634e757ca1113aa6393f113410f49ab624e4d403"></a>

## regex_value property — bot_defense.policy.protected_app_endpoints.domain / 540553ff393a / 5

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

<a id="canonical-7a6cbf116108d6a190f328bf7b22c030f9693817bd1ae3853f9e4d8c32f47556"></a>

<a id="canonical-de53c424b6cc889bc23143c85b12210451ba42602ca3b04092e2d3dc54e940ef"></a>

## suffix_value property — bot_defense.policy.protected_app_endpoints.domain / 540553ff393a / 6

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

<a id="canonical-05bc5d364c1ca3932f1a53e5f882b3404e6b17c87131d78008121087fbf92fa8"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.domain / 540553ff393a / 7

- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-231cca42f3cfe542ec5748fb67389a7429734b86e4e9c8f5b2bcb0871e1a5305"></a>

## bot_defense.policy.protected_app_endpoints.flow_label — bot_defense.policy.protected_app_endpoints.flow_label / a67f2d4f9c06 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-9536712ed542f66c399e6fa32bd955ebe0a314dcf7c09887789fd6776d1de3c9"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

<a id="canonical-6204195066f186033419a3db9f1a87cc66e7349d421308c8ba92fb4e1397c9f3"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label / a67f2d4f9c06 / 3

- [account_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-2a0fc6230b3e698f0148d207a24f0bc378ea29062bc2a0c3691c4c2107941864): complete subsection reference.

- [authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a): complete subsection reference.

- [financial_services](data-sources--http_loadbalancer--reference--group-011.md#canonical-7280d433821056445915875bb74a212a6f667280ec0c54e339c1ba345f775b3e): complete subsection reference.

- [flight](data-sources--http_loadbalancer--reference--group-011.md#canonical-4a541c8519d6264b5c03337bb67df59f72193e74389fa07692ffe56c6b8d7e95): complete subsection reference.

- [profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d): complete subsection reference.

- [search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b): complete subsection reference.

- [shopping_gift_cards](data-sources--http_loadbalancer--reference--group-011.md#canonical-bca13a6c757bd474155422fd9d31efa518e780212f6ac3fb532fa794fb8dc2ab): complete subsection reference.

<a id="canonical-44b9e0d39786b76945d4cb8207d229bd0b9262b0713636c28f07f33bc7cdb78f"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label / a67f2d4f9c06 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-2a0fc6230b3e698f0148d207a24f0bc378ea29062bc2a0c3691c4c2107941864)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--reference--group-011.md#canonical-7280d433821056445915875bb74a212a6f667280ec0c54e339c1ba345f775b3e)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--http_loadbalancer--reference--group-011.md#canonical-4a541c8519d6264b5c03337bb67df59f72193e74389fa07692ffe56c6b8d7e95)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-011.md#canonical-bca13a6c757bd474155422fd9d31efa518e780212f6ac3fb532fa794fb8dc2ab)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2a0fc6230b3e698f0148d207a24f0bc378ea29062bc2a0c3691c4c2107941864"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-112c302e443be71608c4e4799295ace9591c6ac415147de19da1150b6f353c29"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management — bot_defense.policy.protected_app_endpoints.flow_label.account_management / b4501b8e4fde / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-728e355a9581d687b86547466610f3bcccbccd8382611c280105f4594858624f"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

<a id="canonical-bb43067fd0b2205f324a373e07c86a7632c4688e4ca52f4f29ad12a9305724ad"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management / b4501b8e4fde / 3

- [create](data-sources--http_loadbalancer--reference--group-011.md#canonical-9b09ae9beb9df3e53c1c90deabdf47ee030201155ee11fad3819676679224099): complete subsection reference.

- [password_reset](data-sources--http_loadbalancer--reference--group-011.md#canonical-94f14e024ae73481912b1dad43013b0674b81f8c10d201f330f7d0b71234c164): complete subsection reference.

<a id="canonical-895733b1ad4b4dc0d3f8a1ef32349e3b0a3b0182e8b770ab0ce3874fd8b2398e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management / b4501b8e4fde / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](data-sources--http_loadbalancer--reference--group-011.md#canonical-9b09ae9beb9df3e53c1c90deabdf47ee030201155ee11fad3819676679224099)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](data-sources--http_loadbalancer--reference--group-011.md#canonical-94f14e024ae73481912b1dad43013b0674b81f8c10d201f330f7d0b71234c164)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9b09ae9beb9df3e53c1c90deabdf47ee030201155ee11fad3819676679224099"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3427f0eaacd02eb1fc93552ac1d3e45cc708c44e399046d6d5f29c45d40baa09"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.create — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 40da352f8279 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-2a0fc6230b3e698f0148d207a24f0bc378ea29062bc2a0c3691c4c2107941864)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-4e84f088432ad557c12b9793dc7df0f95793a582911dd0528b3ed6cd7aad8e94"></a>

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

<a id="canonical-4449832dc8691dbce55aeef4e76a9a23d75ce6e6d0b26129bc5738de03f92e9e"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 40da352f8279 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-28568116d3a31a5985f527cd71557d1ebd433fddee54c328a1fb33e3d17fab90"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 40da352f8279 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-2a0fc6230b3e698f0148d207a24f0bc378ea29062bc2a0c3691c4c2107941864)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-94f14e024ae73481912b1dad43013b0674b81f8c10d201f330f7d0b71234c164"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1add3f23230b144dc838d6770e382226117b9edb63eea0feb8de391d2c4c8021"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / bd16d6e74ba5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-2a0fc6230b3e698f0148d207a24f0bc378ea29062bc2a0c3691c4c2107941864)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-015789aac2a8d8ce13df3bcae38c8490a32a2ed199992e8b8eb18fa065889bc8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for password reset.

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

<a id="canonical-c604926d696525ed122052041b7a95f8268ff2630437afa60adbb158cde34e5b"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / bd16d6e74ba5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e2d5a85fee7950c92c1d45011e48b52844c27873d72b31b7fb82faf3f4d58d57"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / bd16d6e74ba5 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-2a0fc6230b3e698f0148d207a24f0bc378ea29062bc2a0c3691c4c2107941864)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09e6b2109ad3f8e4a30850427e9f84fb10cd3b73f49a18f12c6111db51260f21"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 5e272790e6ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-c02978f331132231fe922cab4c2da8c3878fcf7671c10f2f8eb536f6269084d7"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-1c1cb32c72c7711b50dde7d8d0f232fb7efd7e7888e1253250ffcb06d3b72e9f"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 5e272790e6ef / 3

- [login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46): complete subsection reference.

- [login_mfa](data-sources--http_loadbalancer--reference--group-011.md#canonical-e83527d4c382a722a99ab05fbd4ea8390a15e7ee983b86292bfba2570d5ad046): complete subsection reference.

- [login_partner](data-sources--http_loadbalancer--reference--group-011.md#canonical-c45e92f3b4a419457476219940a19f3a90221d67a28d3e86a12dda5c7f40fc8c): complete subsection reference.

- [logout](data-sources--http_loadbalancer--reference--group-011.md#canonical-fdd4a4629e33bbe5c6fb6d8e056ba4d98e703551b4865a1a3d9207dcb3859e79): complete subsection reference.

- [token_refresh](data-sources--http_loadbalancer--reference--group-011.md#canonical-f722267c61aaf758828c8fdd7a6f1a1a3d1243d63a834a9ce7771a2f1b00c63f): complete subsection reference.

<a id="canonical-0e4e3a22e84c8afb450bc895bdc3782f089989774df095ba579b08f8bdbc5311"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 5e272790e6ef / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](data-sources--http_loadbalancer--reference--group-011.md#canonical-e83527d4c382a722a99ab05fbd4ea8390a15e7ee983b86292bfba2570d5ad046)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](data-sources--http_loadbalancer--reference--group-011.md#canonical-c45e92f3b4a419457476219940a19f3a90221d67a28d3e86a12dda5c7f40fc8c)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](data-sources--http_loadbalancer--reference--group-011.md#canonical-fdd4a4629e33bbe5c6fb6d8e056ba4d98e703551b4865a1a3d9207dcb3859e79)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](data-sources--http_loadbalancer--reference--group-011.md#canonical-f722267c61aaf758828c8fdd7a6f1a1a3d1243d63a834a9ce7771a2f1b00c63f)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5daf55c09989cf472328fa4b47dba3ebc6c5e861a3ad610be8e42b59d88dff2a"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 26a20444784f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-ff3114fe07e43e47ffba810198d48edf08f028fbde064eaaeba954f00e1ca988"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-5713d7cba7c9571347056ff91fb2b0b584ebae85a571d0a588873b069fed1f67"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 26a20444784f / 3

- [disable_transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-362c00ab81971ef6a6e9eebb13074bf9d7b70620584ee6beed87df95bda9bf89): complete subsection reference.

- [transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-6accd76879460d8d155e5cc37889d672d2ca5f55a09e259cfa97552205e52984): complete subsection reference.

<a id="canonical-2e86968029ab665e981dff3d3ce62ffca6e8a39af0d9fcf8b6fea6ceb9156a0d"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / 26a20444784f / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-362c00ab81971ef6a6e9eebb13074bf9d7b70620584ee6beed87df95bda9bf89)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-6accd76879460d8d155e5cc37889d672d2ca5f55a09e259cfa97552205e52984)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-362c00ab81971ef6a6e9eebb13074bf9d7b70620584ee6beed87df95bda9bf89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a1498484c544f82a08f29d03840e70940183b5f871cc2047c10833ec8cd7c46"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 3fddd80b3a89 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-263a5b33a312ef5298c255637d60e41629c0b5e78e31127f6ac33101abaf37f4"></a>

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

<a id="canonical-899e1bf0fce1a5a0ad5ddc60bddc276966c77a6f7cb4d1aaa6204905c1472480"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 3fddd80b3a89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b7f40f81cb366d598c063783822d3ae9829c5424f4c362b3ebec79771fd13ae"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 3fddd80b3a89 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6accd76879460d8d155e5cc37889d672d2ca5f55a09e259cfa97552205e52984"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7a527a2a22550f6bf971c2fc26afba6d5b5fd8fb03d9760584e8366380ca1b0"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f6275c800d68 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-c9ed27b36cf752e43c4d672b65959986b37683a1409217ead9165c18e62091d4"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

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

<a id="canonical-6f37f9ef6bc421b1adbeccca482de0353e53f09231922f76c6261cfa30332ecf"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f6275c800d68 / 3

- [failure_conditions](data-sources--http_loadbalancer--reference--group-011.md#canonical-a292a3fb838186817b2ee03be4989132603b2195243abe410af454f5eff51d62): complete subsection reference.

- [success_conditions](data-sources--http_loadbalancer--reference--group-011.md#canonical-bea66eb4f46b25eaad2d30f8b3f5f32b0efcb73f3ceed3565a95b11fd66287ae): complete subsection reference.

<a id="canonical-704bc98f338dcb28a8a46fc38a02955f5c537028c8be63cae4b5f24957e4a6de"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f6275c800d68 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](data-sources--http_loadbalancer--reference--group-011.md#canonical-a292a3fb838186817b2ee03be4989132603b2195243abe410af454f5eff51d62)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions](data-sources--http_loadbalancer--reference--group-011.md#canonical-bea66eb4f46b25eaad2d30f8b3f5f32b0efcb73f3ceed3565a95b11fd66287ae)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a292a3fb838186817b2ee03be4989132603b2195243abe410af454f5eff51d62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bcb6bdac608cca38df15f1f86b631d7f11b24f90bdcb56b10d383ea35b5993a"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / db527b14fd07 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-6accd76879460d8d155e5cc37889d672d2ca5f55a09e259cfa97552205e52984)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-766ac98627c7ed26d76d446545eb67c217c8dd340c3ae6096c0a4cdddf075c7d"></a>

Type: `"list"`. Computed.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-370d0e4abe170bf201596de0c34f90dd19e1bb158f4984d2cf5d9082bcd832af"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / db527b14fd07 / 3

<a id="canonical-c4b545f4f8f5ef7f24104da18b68d21fce6909818130798ac1dbac7806228b3c"></a>

<a id="canonical-7579347caebe2ea26eb8ee7d82b0530abe1659ff32136cee5df4234f9d7f524f"></a>

## name property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / db527b14fd07 / 4

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

<a id="canonical-aaa3065fd4861ab397b6235ceda795517c9e77f51042713fca686fda754f2a19"></a>

<a id="canonical-28ee5a2d074bf51f4ecab243088fba92f06ccda5ff6e684f9472a7a46239c0a0"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / db527b14fd07 / 5

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b635c6928461561c2d5edb15acae0d04ea9cb7100566133862a3a2345a7c9603"></a>

<a id="canonical-4ecbecd446a5fa9b8526cf510a5e54b8342eba09bc7d67938588c2752e380d66"></a>

## status property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / db527b14fd07 / 6

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

<a id="canonical-40f68bcae94f8b99bc42c5f55d0d08613b018392bb83e0d517f904658bfb0f07"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / db527b14fd07 / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-6accd76879460d8d155e5cc37889d672d2ca5f55a09e259cfa97552205e52984)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bea66eb4f46b25eaad2d30f8b3f5f32b0efcb73f3ceed3565a95b11fd66287ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f79a13ce89da704a22aad75cc738c9c038fc56057438c74a052f4086a7e2796c"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f4e88c8951b9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-011.md#canonical-f59d98675d7c48eb1c380c969cfffc9a9849d248cfa93925654e744c97496b46)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-6accd76879460d8d155e5cc37889d672d2ca5f55a09e259cfa97552205e52984)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-1ed86f8df422054b5b5c7e6c73b9955d6a9f299133ac05f6499217af3838e8be"></a>

Type: `"list"`. Computed.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a35f15bd4efaa3629639cd7fc6e121a3398783f7dad58c6d99a5578e9556fe58"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f4e88c8951b9 / 3

<a id="canonical-a15fab14bc0091cd463ecc6bcf32da27a15292a76c7e876f35708fa3e62f57fe"></a>

<a id="canonical-ee0a20dc1153a88209374e402f551afa6a533988abd50941fd6bbb55f3c3fb0a"></a>

## name property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f4e88c8951b9 / 4

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

<a id="canonical-448a6d5feb102a332bae6170047a341b15c435fa02a15b277e048c4b80291f30"></a>

<a id="canonical-0439b6b60554e985b027c19d65687f01fe4ab024bc9a188b6dd36436e8ce85c7"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f4e88c8951b9 / 5

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3844d1059d7da44a6df28965e2ea7391126c55428b0df84b9178f9e674c43fab"></a>

<a id="canonical-2725260fe33bd5a880aa02a48986a103a6673d35e153967aea572d9f957da369"></a>

## status property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f4e88c8951b9 / 6

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

<a id="canonical-edbf150df184f609192dcf422e17e9efa64b20ea59e37eee122f532efa84785b"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / f4e88c8951b9 / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--http_loadbalancer--reference--group-011.md#canonical-6accd76879460d8d155e5cc37889d672d2ca5f55a09e259cfa97552205e52984)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e83527d4c382a722a99ab05fbd4ea8390a15e7ee983b86292bfba2570d5ad046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b0f2ed2ed62d7de7f876be3d57049ff5d04d2a12be46a0b652ceab3fbdf8da2"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / da8851588dd4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-08671bbb02f7c932a1d2b6f018da28e60bb317162a7d84f9c11625b676ec61da"></a>

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

<a id="canonical-cb7d92100b03cbe4471596f2e83a4e1590d01fa14b6befb3e5cc1f5159bc6de1"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / da8851588dd4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-551c342c785f96249f29dc2e9f1d400c09f333994ca16e89a6c1120e25fd7419"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / da8851588dd4 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c45e92f3b4a419457476219940a19f3a90221d67a28d3e86a12dda5c7f40fc8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-681fb71aff8d7c25ae7892c9a5d0e641aa30e0a7cbec31f26ab43e88bdade9cd"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / 06368ceaba89 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-458a53b47566dff47098fc6e619718cbd974d416ce0eca178e55e622fa104613"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for login partner.

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

<a id="canonical-05a48f8b6875c0e249b64f34a01f1c4d16bf420bb51d6dbec5ca74da7e54d2e5"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / 06368ceaba89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7bbe9b17796bf612f260cd556e449c0e75c3540e7c8c15386b197012e670365d"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / 06368ceaba89 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fdd4a4629e33bbe5c6fb6d8e056ba4d98e703551b4865a1a3d9207dcb3859e79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19356be0459d2359896eacb6c2818cbfc6af012ded0f432674b68ac6c8e347bc"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout — bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout / 811970ba1b4d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-221455c1652d2064ae410ea0bfcc1ac48637df6532a87066fbf9a73090344794"></a>

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

<a id="canonical-00d0d2ccc4872ca0fcfb58365d55f14ba8f07ecd4cfccfbdb97e3f91227c9016"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout / 811970ba1b4d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7dbaa05cb4d6b982eb703880ebfee8914ecacd6d7fcade0da8b3afd17764d81"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout / 811970ba1b4d / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f722267c61aaf758828c8fdd7a6f1a1a3d1243d63a834a9ce7771a2f1b00c63f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22467259edc714aa47bdb45977df041f2e8802d7681fa9055c5dfaf648948484"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh — bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refre / 7d8b3d758e4a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-c2329f02c6fc2d10e98b038cc60d4662243eacbd285193e3dcef8ff1d50a1364"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for token refresh.

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

<a id="canonical-305a77f4c5851076c5875ee63c5ae1cd89e860e61ebdb8cfee78f7571d63141b"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refre / 7d8b3d758e4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-77ed49a0d5778303df891fc6d8ee0095116b87f09f2884b842bad6962623a2b3"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refre / 7d8b3d758e4a / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-011.md#canonical-8f19c1c6cdc186399d3dd6ded4a14f996f46f927be701af073f3b16fb787612a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7280d433821056445915875bb74a212a6f667280ec0c54e339c1ba345f775b3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5d91aad829c21dd0ae8c5cbd7732253a8e28df432f061ebc157bce0f59b0188"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services — bot_defense.policy.protected_app_endpoints.flow_label.financial_services / 897e1a492993 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-31fb2c6089f287e9f77a77add60f3bd83d7da0a6c52198c8db679eadbb43ea66"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

<a id="canonical-5a1127c3f6df0d498c62121132a04e0623b5d427cb5e585d103f2d9a64b32efa"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.financial_services / 897e1a492993 / 3

- [apply](data-sources--http_loadbalancer--reference--group-011.md#canonical-afc6ff6a012275eeaa87e3f7f1f1327353617d38baf8e5163032638ed6d5f497): complete subsection reference.

- [money_transfer](data-sources--http_loadbalancer--reference--group-011.md#canonical-c0b2cfc17a245ce93e13e0a00877f66eb0a5d9f4b6285aed0697fc18d6c094f4): complete subsection reference.

<a id="canonical-a894ca97836f6703a0e8183ee49be1b91e37fb150c17e57676a81e469e4f0521"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.financial_services / 897e1a492993 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply](data-sources--http_loadbalancer--reference--group-011.md#canonical-afc6ff6a012275eeaa87e3f7f1f1327353617d38baf8e5163032638ed6d5f497)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer](data-sources--http_loadbalancer--reference--group-011.md#canonical-c0b2cfc17a245ce93e13e0a00877f66eb0a5d9f4b6285aed0697fc18d6c094f4)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-afc6ff6a012275eeaa87e3f7f1f1327353617d38baf8e5163032638ed6d5f497"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-532b35aff363e5b884b89d750112878d10dbc37425282c301d66bc96cbc0ab44"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply / 8291e13bc590 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--reference--group-011.md#canonical-7280d433821056445915875bb74a212a6f667280ec0c54e339c1ba345f775b3e)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-9d6a8cd05bfefc97b77f93da29dd39e29f95e297306c9256a4e51b4aafd181c6"></a>

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

<a id="canonical-32b22ffb1269764940fa44a260a25821c703648ad6b72444b3ff617424c8ec2a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply / 8291e13bc590 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78a2688f6d02e674d07e8947566dfb0d0788b8dc47cbcd3a1d4a37e00ad3acdd"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply / 8291e13bc590 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--reference--group-011.md#canonical-7280d433821056445915875bb74a212a6f667280ec0c54e339c1ba345f775b3e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c0b2cfc17a245ce93e13e0a00877f66eb0a5d9f4b6285aed0697fc18d6c094f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-196d55003cea4f3d95be342d71401afd43d72ab3854f6f9074d8d23decd24077"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_t / a6e5b9a11099 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--reference--group-011.md#canonical-7280d433821056445915875bb74a212a6f667280ec0c54e339c1ba345f775b3e)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-e3277fc02c78e978b10a88e60bb1dba1c6bf40d3eccfa77e7eb5ae08401203ce"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for money transfer.

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

<a id="canonical-be73f49d5b0281e55e9628afc9537a7eb8fc1c2e36a7a4a1fc0d404c1ea98b4b"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_t / a6e5b9a11099 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9289865351dfcb239f0c4926c33ffedb2639c55dfd22feea889dd674efc4b8a7"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_t / a6e5b9a11099 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--reference--group-011.md#canonical-7280d433821056445915875bb74a212a6f667280ec0c54e339c1ba345f775b3e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4a541c8519d6264b5c03337bb67df59f72193e74389fa07692ffe56c6b8d7e95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a5378266bd79879b551c3e36c4431818f6eb29e070308fd7ed747f9a4b0596d"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight — bot_defense.policy.protected_app_endpoints.flow_label.flight / 39f16b6378c6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-a35b888e289697cb176c756bd6d950cdf6d521d28138d1defe8df07fadeb4e56"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

<a id="canonical-6956c5e3b556e0983608a9fb2bf8edc8e4d460bd1c90f8031db6b9a2c413a34d"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.flight / 39f16b6378c6 / 3

- [checkin](data-sources--http_loadbalancer--reference--group-011.md#canonical-6aed3afec9f6d29f4e65c9e5792a0139469463303f81524b1a91727847fa3a95): complete subsection reference.

<a id="canonical-bf60afb77803e5bf4f7130b74ec2e0b3e7434b3ce315cfd68d8690d06458fab8"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.flight / 39f16b6378c6 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin](data-sources--http_loadbalancer--reference--group-011.md#canonical-6aed3afec9f6d29f4e65c9e5792a0139469463303f81524b1a91727847fa3a95)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6aed3afec9f6d29f4e65c9e5792a0139469463303f81524b1a91727847fa3a95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ce652b5e440dfd57e56b882c812a51652a47849cb5c33b414a61de638b507f8"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin — bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin / 3fcedc0649d2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--http_loadbalancer--reference--group-011.md#canonical-4a541c8519d6264b5c03337bb67df59f72193e74389fa07692ffe56c6b8d7e95)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-8d80c64fd7afe4b8b2cf8086bb16a4ac088040c6ccb4e6510ff529c9a9f97985"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1507298fab26793d64279ba9b7f23444113f9605e606affd73fa4a3254011c8f"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin / 3fcedc0649d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8bf3a9d381448075c40f54f8a9ac7c5aeef1e5dff3661292b3114bad8d3dfedc"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin / 3fcedc0649d2 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--http_loadbalancer--reference--group-011.md#canonical-4a541c8519d6264b5c03337bb67df59f72193e74389fa07692ffe56c6b8d7e95)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2afc7f11005f937a899b7bdacee9fa738b8d231862a20f0e2270862a14819831"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management — bot_defense.policy.protected_app_endpoints.flow_label.profile_management / 70e08ff5e857 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-67174f18bf556134ffffdf285d1932aa6634c726ac73b51f12d82a74d94a88a1"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Profile Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

<a id="canonical-bab2dcd9ed322eb59cedb55586082225f12c2bf868b2f3d1c260377a49f24f03"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management / 70e08ff5e857 / 3

- [create](data-sources--http_loadbalancer--reference--group-011.md#canonical-8075399283b9ce2da77e57beb784722c7bb40dd24e32fc535efeb6dd254203e6): complete subsection reference.

- [update](data-sources--http_loadbalancer--reference--group-011.md#canonical-acd846cacffff111c194c197986712f08742d60fb256396ab364c5f3a4cc5f27): complete subsection reference.

- [view](data-sources--http_loadbalancer--reference--group-011.md#canonical-56982cdbd98874cd339a37355bfd4ab1f7a8a5c52692eebb520dea38f265fbad): complete subsection reference.

<a id="canonical-847acc8537b9d3e282c3456b99ac4c7d9d35c0d9856c1a5a0d2c189d91789b88"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management / 70e08ff5e857 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create](data-sources--http_loadbalancer--reference--group-011.md#canonical-8075399283b9ce2da77e57beb784722c7bb40dd24e32fc535efeb6dd254203e6)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update](data-sources--http_loadbalancer--reference--group-011.md#canonical-acd846cacffff111c194c197986712f08742d60fb256396ab364c5f3a4cc5f27)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view](data-sources--http_loadbalancer--reference--group-011.md#canonical-56982cdbd98874cd339a37355bfd4ab1f7a8a5c52692eebb520dea38f265fbad)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8075399283b9ce2da77e57beb784722c7bb40dd24e32fc535efeb6dd254203e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fd3d8546c3afd6e9be777a5e7e1f8c2996a76ac0f2d92fd71c7902ec925642c"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create / 7b22d5c87027 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-789c65fc7cf7539291eafc6621d230896772250e3b877aa6bdf331b132c5f078"></a>

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

<a id="canonical-19d6b31c1e5e54f96aeca7a46b1f5841753726a66610073d0baf98182460cf92"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create / 7b22d5c87027 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6cba14aa8bb057601bbeff53c016e34f23bc41104e92bf97f7ec24431a7c9c7b"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create / 7b22d5c87027 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-acd846cacffff111c194c197986712f08742d60fb256396ab364c5f3a4cc5f27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb9c36a1e84fd779f168c98d26e0b1851727b5429931adeab57c6be23f246902"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update / c9369af6acda / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-e9d1947aa1d99863f52e40182602042f9cbfa79cd1831748d87a88684d8d3cc4"></a>

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

<a id="canonical-48a5ac104f4c556f37884ee534181042213d5b59f35e32ce6ed083d5529c394b"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update / c9369af6acda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e02cc7dda91921f386159b427544ece9e86a425d18504c2c5384978d2cbe79a"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update / c9369af6acda / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-56982cdbd98874cd339a37355bfd4ab1f7a8a5c52692eebb520dea38f265fbad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-915b058da659487415a460ec14ef3754602a65b5ff530fcb121ed83d70a3dbbb"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view / 404461d3b3f6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-26a42d2a5a13eea2ab6562ff2cb23861edcb652934b22dd68e9baf452e1fce55"></a>

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

<a id="canonical-59290d1ceb061812e6057724d5e65a24447cedbe8c422064ffc9b0a00dd362dc"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view / 404461d3b3f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a69a7b3d64518e7c2ba1cf83b8a4a303f2a91e13c1922061850d974e835f1b4"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view / 404461d3b3f6 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-011.md#canonical-f4a488eac44c0c42c5cc909ee6ed184d322bb970bff83f0fa17f06fc57d6709d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b88df8542ad1d87ce061fc9b33ef7250ac6303604b1eff44260b23192588f559"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search — bot_defense.policy.protected_app_endpoints.flow_label.search / ff7fd85a40a6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-1db2beef98f563f93029e6d490a7d7726e10ee5957b27ecdfd21c8ce98c36850"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

<a id="canonical-31ec449616292aefb2269de102bbc244d2023338199e44e8a8d0e730a914cbf3"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search / ff7fd85a40a6 / 3

- [flight_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-41126f604507ec3a7fe8d7773a6ce76390db85fc8c709b5c251f71d3a8ead5e2): complete subsection reference.

- [product_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-48ae55c1461484a73c18d1cb4ee8c1392499a25801f7dd0215471d1840cfb868): complete subsection reference.

- [reservation_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d0e12de90f3c7af61ea7b702f9a239ab497347a1081c2ada968262b1e992ed66): complete subsection reference.

- [room_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-e6dcb830b88c3f0b951f0197a4b9e5285e8a9f07611895c17c9c19a6dbd04212): complete subsection reference.

<a id="canonical-8c809b65d01b0295a1d3090c4a8afb20cfa668053bf5eaad8d363dfd76f3649c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search / ff7fd85a40a6 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-41126f604507ec3a7fe8d7773a6ce76390db85fc8c709b5c251f71d3a8ead5e2)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.product_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-48ae55c1461484a73c18d1cb4ee8c1392499a25801f7dd0215471d1840cfb868)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d0e12de90f3c7af61ea7b702f9a239ab497347a1081c2ada968262b1e992ed66)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.room_search](data-sources--http_loadbalancer--reference--group-011.md#canonical-e6dcb830b88c3f0b951f0197a4b9e5285e8a9f07611895c17c9c19a6dbd04212)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-41126f604507ec3a7fe8d7773a6ce76390db85fc8c709b5c251f71d3a8ead5e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-264428e7ea342f9e96961696f3f84ecb327175f6b9051e20fdc2ee261d670830"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search — bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search / dac4ffe970a5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-fdb4924d5c5e7c8aa58a8e11e7e8649e12d7d0921d2f7310a47be8b93dc2dcc0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for flight search.

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

<a id="canonical-c358681e22ccf2d046ddedde4b623e0811dac4fe44a99a4d657e4ec6797f473f"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search / dac4ffe970a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf74ed97bea7b86c183321d8dcdb7c001919fb295834f908e6144fe9cb6f99d6"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search / dac4ffe970a5 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-48ae55c1461484a73c18d1cb4ee8c1392499a25801f7dd0215471d1840cfb868"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a02f91340f7b48da2252e5d59a4700b197ea21d55bbac06279ea99bd491b8807"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.product_search — bot_defense.policy.protected_app_endpoints.flow_label.search.product_search / 29cdfaa45bc0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-0986b8eab30b5412774a8cec6fc12f9e24a065063a81973fabd9270c77fe1518"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for product search.

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

<a id="canonical-5e72b6368c2561414cca56767b2b564db95f5d72e5941d85de0a2ef3f6c6d565"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.product_search / 29cdfaa45bc0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac37e3cf928681c7dceba9d856cefa1cff75edf133574f546c3caa024c2ec88e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.product_search / 29cdfaa45bc0 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d0e12de90f3c7af61ea7b702f9a239ab497347a1081c2ada968262b1e992ed66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9eae5254fed73c9c826b0d4ac9a32b293ba7fa1baff06822c386d5d5d5bea55"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search — bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search / a9945a1065ac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-ec248d78a9065ecaa58db7b1df7401dcaff267e9471c30ab6d64045286b52005"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reservation search.

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

<a id="canonical-6ea4bbee1d8f5fc6c7aeb7368684bc653f91f179b396ce2ec49e06a6654cb652"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search / a9945a1065ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf87441f375cd9a6aa2fb7e4c2d20e63cfeed79d2e9e02cd199260d12eb73535"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search / a9945a1065ac / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e6dcb830b88c3f0b951f0197a4b9e5285e8a9f07611895c17c9c19a6dbd04212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30627726ea4be336ab255b61c6e807f58fc13c1f06e726ae89d4fb779740e366"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.room_search — bot_defense.policy.protected_app_endpoints.flow_label.search.room_search / b07414a21879 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-d14d94bcfc8795be11904ff25b468d7c744d76d2b42efd6dde01749530fd3dbb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for room search.

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

<a id="canonical-f4a604b76fdc89e0e1a4f7f6fc9bcab92661365b7c483892bb31cd091339fe39"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.room_search / b07414a21879 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6eb14059af2abc11738a48c02bdd68d565af9cc12d3e57d23db88d323a69b775"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.room_search / b07414a21879 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-011.md#canonical-d67cbd924e0fa53bdda37a83f2dc66e592cfd2f883afeffa9b25c551bf24999b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bca13a6c757bd474155422fd9d31efa518e780212f6ac3fb532fa794fb8dc2ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec4c20038b03966b90ee0f646dd11a471411af242441683f74c814dcdcfbee7e"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards / ea844e3fc195 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-011.md#canonical-d049e92c51cfc14dc82de70abe42704f404b9f738bd719ddcf1e02ee4c686896)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-9f737ffab15268af55411d214372e939b89fabe9acd18dc77fbe2fdc072a1b42"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```
