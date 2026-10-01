---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-e0c3d91142e395cd8b3a0e841bfbcd805999d09a5606408dd312cb32f8eb17d6"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_pres / 5895f2ed1fea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7795f4786e89c9439c284c313f904914887bf1142105542d2c9b95f4c5211307"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_pres / 5895f2ed1fea / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b051e8a43c9083ef46f46183bcdd9f2a79a23f30ba4190210dc9e419a4febc81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55e04b46307d9f6044de14e3f8d207ff98d4275b1b45a4f2ffbfca687b7630eb"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / b009384d8fc6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item

<a id="canonical-9c2921fae51120d14f4917bf3a7be6b8c24ca76fac03a2a7a44fbc700bb1c3e5"></a>

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

<a id="canonical-c2d710132ebb4e80826637331a61d1dee99cf8d1faa1c003801723ba458aa4e1"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / b009384d8fc6 / 3

<a id="canonical-030b78265ec9807297254f639ae7ef113ab4e00d364637522bd2f06b72bc314b"></a>

<a id="canonical-fe6433bb45cae8df991a464fc16c0e7273ebe0d42ce5b7159e337ec45d4a2dc9"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / b009384d8fc6 / 4

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

<a id="canonical-13cada11ebb2314bbc512b0bb33c577149148b7f46c882745b8ad7294aab5bcf"></a>

<a id="canonical-6d72590a8b2b88d54073552c3b330f4ed2f54c212aab3933b2f157cf11a599f3"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / b009384d8fc6 / 5

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

<a id="canonical-0006f5f0181d8e014b4f1ae00b93a6c7c40c1dbb52d6f19d7f1a1148a9bc997e"></a>

<a id="canonical-9edc8d3bb10259d2c5a7b29aa31236900db92ee9cfc93c88018a5986a42de6c1"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / b009384d8fc6 / 6

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

<a id="canonical-c8d145cd29327a080c09b4ad829d21936ba919e8d0bbdb4f377471044a4e5411"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / b009384d8fc6 / 7

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-e35159d06bb46f012f8c93b49dcaac00cf5704f4556cc3e091ab22aeef5df1f2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74153d24f943b80f8d03da8b4a75ceff2aacf7aa089dc7de2467489d0bdb0bd4"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers — api_protection_rules.api_groups_rules.request_matcher.headers / c029a8fad6dc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- api_protection_rules.api_groups_rules.request_matcher.headers

<a id="canonical-9655fc12679a926bbc441bb8fa4d4838d86b6309f2de94d7ca69cd0ac06b82fd"></a>

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

<a id="canonical-97d6ae8b7e012b3fd24381f673a806000896464ed4965686851713efc860ddf0"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers / c029a8fad6dc / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-44dac61b1f21f15ca47ae43d165a4ff69bad90e02dea685df28f42582f8f3958): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-61af8ae7489b0d95612de9cfc40b9ef19dc6bdc6c33da00e13d5b0b35dce25d6): complete subsection reference.

<a id="canonical-9e52ff61ca19c1b023497d0ca88b0e67b82759d0859098941c891548fff69067"></a>

<a id="canonical-94baa230c11664a4446d11153cf36a71dbf3f7efcb0683f9ffafd4ed5b555313"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.headers / c029a8fad6dc / 4

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-7da503f8c6c2ab4d9fe55b68acd3c9c6eeca84b3716e01a0944e1e6e2973b674): complete subsection reference.

<a id="canonical-2ebc8a52c1e81d1fcc7af52adcdb03984f5f9afb569706fdb157511ece7c8a35"></a>

<a id="canonical-b9689240aff4f319705cc492c095f8a0089c29a0eafdf8a40c6dfe3bb6f4fb92"></a>

## name property — api_protection_rules.api_groups_rules.request_matcher.headers / c029a8fad6dc / 5

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

<a id="canonical-8042128adf841f55b6f015c5a0ef7558c63bad496dd9fca69cd1a2ce56b921f1"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers / c029a8fad6dc / 6

- [api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-44dac61b1f21f15ca47ae43d165a4ff69bad90e02dea685df28f42582f8f3958)
- [api_protection_rules.api_groups_rules.request_matcher.headers.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-61af8ae7489b0d95612de9cfc40b9ef19dc6bdc6c33da00e13d5b0b35dce25d6)
- [api_protection_rules.api_groups_rules.request_matcher.headers.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-7da503f8c6c2ab4d9fe55b68acd3c9c6eeca84b3716e01a0944e1e6e2973b674)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-44dac61b1f21f15ca47ae43d165a4ff69bad90e02dea685df28f42582f8f3958"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec1c9e6a7646fc2c56a5b7abacafc2e6f0de8f0774126ec864873ef769678860"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present — api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present / 65ee5bdf1a84 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present

<a id="canonical-de96ff7bf4908f05e2ffbc7efbaedd891e166efa3089ae7760c542c170306e61"></a>

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

<a id="canonical-82fb597cd55a59f74f44d66b59ca08baa9f4bc09e56ed6dd54450333590f0b2e"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present / 65ee5bdf1a84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e33d7ff1bd5f21045ef6c6fbecbfb2d23bef44b9f3d06f99657ebc5827bfdd7"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present / 65ee5bdf1a84 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-61af8ae7489b0d95612de9cfc40b9ef19dc6bdc6c33da00e13d5b0b35dce25d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0661fe708c11905884ac9275e73a933c2b8fc911d4f95056c809a0cbcd334f73"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_present — api_protection_rules.api_groups_rules.request_matcher.headers.check_present / aaca21b40343 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_present

<a id="canonical-03b95324fdce9938f0163474b4a4d44bf32e2c407ff3256c240533515fa0f8fb"></a>

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

<a id="canonical-7b2bb5557bd01e60cd4d5e6bc5250ac3d4464743c1b76a0de3711de94ee50dc0"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers.check_present / aaca21b40343 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-756ba2bc8a0a3eb242df960e691bce9705f5967301a73a92365488ea6b863c37"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers.check_present / aaca21b40343 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7da503f8c6c2ab4d9fe55b68acd3c9c6eeca84b3716e01a0944e1e6e2973b674"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-940a487f6fe527bd36efcc77d6631bfc7e25037a467ca7bd7e6f90136d82ebf5"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.item — api_protection_rules.api_groups_rules.request_matcher.headers.item / fa78ab32d098 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd)
- api_protection_rules.api_groups_rules.request_matcher.headers.item

<a id="canonical-9f4f5707bc508f2de837b4996e41233c572d5e6ead09ba5d3081360f4175ee3f"></a>

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

<a id="canonical-3894cb6589a8f2c0cad4530085c5bf0fdec996056866169f98b0c25be0b75c08"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers.item / fa78ab32d098 / 3

<a id="canonical-293ab0d84e7bf119e3bbc21af0dcf84aede1e6dcbb36102085e64658d7e40589"></a>

<a id="canonical-c724738b1ff4a64aac9494ab6c19e5a7731d849f5664585db08b976163d68092"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.headers.item / fa78ab32d098 / 4

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

<a id="canonical-dbc133091aafa7f011f4efe12c7c0c66ac715bc31f900ec2cdaffdd1b92f2b5a"></a>

<a id="canonical-94acd8237a92f379f25740ff1c67c91a434cb062a4f5876f5bf46ea756a69803"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.headers.item / fa78ab32d098 / 5

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

<a id="canonical-8903cfe8b3a6e7dddefd2b540c8efa0439ed6e443a64d355417ad13e23071d22"></a>

<a id="canonical-98956f45df18bec08de44d1a7c7a00074eb2cc941b610706617ce6ded8908569"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.headers.item / fa78ab32d098 / 6

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

<a id="canonical-60943362ebed1801faff9296944cb412b94b89c9a9d0ab68849579ea6474a2d1"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers.item / fa78ab32d098 / 7

- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-2e992deebbcd5dfb1b352c1da1094f5bc08dd4fc5e22fc8ce2faa2c5f01a50bd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-489bcd6b25fadccfe55fb0f3dca4772db72cc480b6de0142a8fb1d934abc6478"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / b253d3641ade / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims

<a id="canonical-cf75c7dfa19b0d9b686a619d10dc929dc3b205f3ed4e58b94eb51fdebc715666"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

<a id="canonical-cb98c935a92341c0bc1a0e28639eba6a39884cb58b89b95d3ef47529f45e26bd"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / b253d3641ade / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3c6eb85168b4027e05c1db70cefd9b66b34137322d1dbb77a9aa524a8d26f3eb): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-dc86c1aba1fde39579b0a7b61df652766ca8221a9d76205be021337346b5ed98): complete subsection reference.

<a id="canonical-8cd5be1d13790348585ba5c4097502eedbc8ff62386af7228d9c3cb6a7f11629"></a>

<a id="canonical-dc94bf2f3bf006936faf6b617398bf3d6a0a61f0dabaebf531c2ec063b5786b0"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / b253d3641ade / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-bc83f671917636f6410d402bc416651e8f9b7602fccc8f4607f502792b7be585): complete subsection reference.

<a id="canonical-13534ca6d7ce577a0837f7ad9e156046224115ccd575f97f3269ca44cb268f12"></a>

<a id="canonical-95300bc678c570aacefd87d67ddfaa5a235f9a794133bb2ff0d4b48fcdbd96fa"></a>

## name property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / b253d3641ade / 5

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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

<a id="canonical-a93a1a520db4f9148a5de8a20cee78c25f20903364519b4db0e5c92975e47535"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / b253d3641ade / 6

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3c6eb85168b4027e05c1db70cefd9b66b34137322d1dbb77a9aa524a8d26f3eb)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-dc86c1aba1fde39579b0a7b61df652766ca8221a9d76205be021337346b5ed98)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-bc83f671917636f6410d402bc416651e8f9b7602fccc8f4607f502792b7be585)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3c6eb85168b4027e05c1db70cefd9b66b34137322d1dbb77a9aa524a8d26f3eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44fa9c95b35283fd2e6e195778d28badbe612e21f4c074e27d47aca572e12116"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_prese / e1e5a02bfa29 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-a0cace09ecb252ba8a35373dee5202bd814bd14d4ac0ce2657d89c4e6672accd"></a>

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

<a id="canonical-9ee60e991f50ab6f38a7a1c416b1a63bd1f7f863e503a5e8a28bb133b1db31ff"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_prese / e1e5a02bfa29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8dedf7959767fffc384208e79b7e0741c1fd01c624cc493e3d6874a3178e6da8"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_prese / e1e5a02bfa29 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dc86c1aba1fde39579b0a7b61df652766ca8221a9d76205be021337346b5ed98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a48481f420cd2e1c4c1327a28c040db21816273ba5de8fbacf271607b00c4d04"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present / 58eb3b678cf1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present

<a id="canonical-37e5d64a54b14c0cb12a07c273ec352a33153a8c110473c8a90c59fa45bcedb8"></a>

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

<a id="canonical-0fc60b47989990332562a94958897061db568f8b5bce9e07b9bc521e00cd1490"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present / 58eb3b678cf1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef78acf794ddb9ddedf8a19eec7eab8204e32721f728a6fafb8132466e2e1c70"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present / 58eb3b678cf1 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bc83f671917636f6410d402bc416651e8f9b7602fccc8f4607f502792b7be585"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d745a8ffaee9753a06b99b82e216d630a9aa2726101e4ac6ff742a79f2d83e19"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / b65f0f5accef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item

<a id="canonical-0f9baf11cfcfb4130e195e68ede1c05c46a60026d81a8885ef4a841c6b1f3f1e"></a>

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

<a id="canonical-9aed77d1d7a420fd546d2124dd28e245278d332567e6555ee108b857697d0f6f"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / b65f0f5accef / 3

<a id="canonical-901e60d3e73bb0d5928c645ad4cbe1c084c7b0a65f23cdf46e7cc5d0c9d76559"></a>

<a id="canonical-c45c9dfdf0494d058f36509739dd24bccd74cf4b4952c3660da0a559c389eae9"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / b65f0f5accef / 4

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

<a id="canonical-c83e76ee30c33850b6df36e5ca0e6f75cdb2a9f093573ff2cf6d4df2959d8680"></a>

<a id="canonical-534d6a71126d88233806b7f6e0308ec6515ed6e1fbec4858fe19354300ded912"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / b65f0f5accef / 5

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

<a id="canonical-d23dad9d22d5910e5070a4541c94b1d182d74beef5b8069c001b1ea0528a942c"></a>

<a id="canonical-e3f3f93c23c5e41f98c512f256b119cfbedb0ddb8e8a43a9522084b4305a976c"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / b65f0f5accef / 6

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

<a id="canonical-7d38dc4744ebc7ec47d3ecb7b41493c088dfd9c96bcd5820b94862a9a41965b9"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / b65f0f5accef / 7

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-57e6f0503ee00e535beacca7e3fc592b9c63a2aea6f567472d36593f7ac3d5d0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8becf3b870267a5feae99a9b29ae882e9eba86ba98afff04037ad092725dc0c3"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params — api_protection_rules.api_groups_rules.request_matcher.query_params / 06e1ba60d6e9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- api_protection_rules.api_groups_rules.request_matcher.query_params

<a id="canonical-777222519f874274841022ab3cc60f9bc0e4d4186c22043b5f40dffc97a46235"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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

<a id="canonical-91a5dd406c68b2e1050a264c34ca9142567fdfb1d9a563468797024c7bac1cc1"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params / 06e1ba60d6e9 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-e3024023626baf71042a50974d157c0acdf640c7c042c59d8ce785953905b14f): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-67ed1c5ff7f24ed2f1d51ff496d5bf0cb897008f22b5f505cd0f11068d51fdfa): complete subsection reference.

<a id="canonical-e5f4a4aae2076b3590090bd8916a83525bbe7d05fa754b4954de915a8b45ea34"></a>

<a id="canonical-a1392f3bbab671d5ba93dc7a3380019f2b59768552bbcd0d61c4bfcfa08e33f2"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.query_params / 06e1ba60d6e9 / 4

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-ec1b354e49f1c06a6707b7ee34e9a2173d28356c42388d7c61e8ac655eb276e1): complete subsection reference.

<a id="canonical-6db7e4516a4d8de57a1cd4b3a9ead5c3a8b8ebf6a25d5d46644ba87a54425a8c"></a>

<a id="canonical-d7596b666ab32e97711c6da5f557fd55a594c963ea73768771b3d0c20ac2203f"></a>

## key property — api_protection_rules.api_groups_rules.request_matcher.query_params / 06e1ba60d6e9 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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

<a id="canonical-bb40b792329071296c226678c51abdf96e95de669f57e3c71f5ff8e7f137917d"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params / 06e1ba60d6e9 / 6

- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-e3024023626baf71042a50974d157c0acdf640c7c042c59d8ce785953905b14f)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-67ed1c5ff7f24ed2f1d51ff496d5bf0cb897008f22b5f505cd0f11068d51fdfa)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-ec1b354e49f1c06a6707b7ee34e9a2173d28356c42388d7c61e8ac655eb276e1)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e3024023626baf71042a50974d157c0acdf640c7c042c59d8ce785953905b14f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3880e392df87f8b63c250cc452a8caa59c1d5f28464feabb105379043fd0cbc"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present — api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_pre / 40e4600673e1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present

<a id="canonical-596cee4cf12da09f609ca6639f297de644b4b4974230b6f57e5084d6b895f515"></a>

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

<a id="canonical-441a4847b6533d8d0cec899105af43d9e525106b52b18547665d04f9ba7e6570"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_pre / 40e4600673e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de65f76dd90e03245c67bb118bfc4a1c5653135916e878703797684ec9aaace5"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_pre / 40e4600673e1 / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-67ed1c5ff7f24ed2f1d51ff496d5bf0cb897008f22b5f505cd0f11068d51fdfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffab739bd6ceac4fb6b937aa3b7709e1077862a1fe13d567375aaefe51977f20"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_present — api_protection_rules.api_groups_rules.request_matcher.query_params.check_present / e048f6ee12cf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_present

<a id="canonical-27d07e6b9c3bfd3c0c485f348c9ea46404a6f9da459d01c9522d50c9ecb73b82"></a>

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

<a id="canonical-4b17e0282929da505dfacfe51bfc97b3e47ce75174dc206dc7b774af5a9a0c12"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params.check_present / e048f6ee12cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-32e4c94ca2ca0a938f316dbf6638877614fc84961012a43e6faaf8595618be40"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params.check_present / e048f6ee12cf / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ec1b354e49f1c06a6707b7ee34e9a2173d28356c42388d7c61e8ac655eb276e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6eb8c10d752dd89083805b97a7bdf69d8187e41c20cd49557b201fd4e15cb88"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.item — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 8cc18c203256 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-415db5a0c9b555a56c81eacd00ebe6899c44c8547e4545b1bf69b6617218b318)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-82b8bd5c4f0bfd4d3edf41b2fa947b76e00a1a17ce867325e634eca0c5a54eb7)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1799983b63e67772d3b71c4156c2cd18db14ba86436ab2971e1a87048d4ba462)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14)
- api_protection_rules.api_groups_rules.request_matcher.query_params.item

<a id="canonical-8df95aeac89c3a9e1ef50414217dd3a5d0d5f3f8fe84b98b2e27a702375d33f0"></a>

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

<a id="canonical-dc88680efa8d0422b9562b0c2eed01ae7d1b46f5790b79c942161761dd6e7eaf"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 8cc18c203256 / 3

<a id="canonical-2911113aec463c37879604e5efaf9046a04d6495de260ffa6fce60eea2c289db"></a>

<a id="canonical-b554b8ae4d03b99fb2fd930996eb49d92781a9c009e3d1eb946cc7f0db887a7d"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 8cc18c203256 / 4

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

<a id="canonical-454fee0a33a4e74de4b18051656453f0cdd7c9d0c465a6f34c200f59f20a228c"></a>

<a id="canonical-d12439ee91e569bc1ae210d1f62eb6727714b5fe9264fe1b413646975a11d862"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 8cc18c203256 / 5

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

<a id="canonical-b9bc69a9e3976f4b8a7ac621235f8450a61ed74672903a1967781f757d37157d"></a>

<a id="canonical-114a1e295a3b89cdec0a3ec4875c95f37d67eb03ccbb1f9b1eb1d42780a83e87"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 8cc18c203256 / 6

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

<a id="canonical-2853505b6941a5da93386eacd5186ba98bfaab1cbd8d87df4e1260d32ef38a5e"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 8cc18c203256 / 7

- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-56d69a7fed65a4971608d20e3406eea39bed0b6487892eb7ef89edd532e78f14)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e134e851dde4d9d53a9e30501ad1cd938e8d06bf26f59fea0d6282dd99d1349"></a>

## api_rate_limit — api_rate_limit / 3ab8d2bafa04 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- api_rate_limit

<a id="canonical-422d912fa2d0eb414f9415712c4c6c017f389bfcf4a444184813e1bbe6ab56f9"></a>

Type: `"single"`. Computed.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\] Path-
or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and choose
inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Upstream description:

Path- or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and
choose inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-422d912fa2d0eb414f9415712c4c6c017f389bfcf4a444184813e1bbe6ab56f9)
- [disable_rate_limit](data-sources--http_loadbalancer--reference--group-017.md#canonical-b2480f6fd6f9a92a4645c5b143d29c07d358fd24e7d4b7bd9fe593cb83a47bd3)
- [rate_limit](data-sources--http_loadbalancer--reference--group-022.md#canonical-155da133b1ab0a39650fe0d779ba2bed147e2b177752a6c36db1e3d419936057)

Select alternatives according to the provider validators above.

<a id="canonical-3d80c2d5809a9d2bc3b5f53a182f37ea56017695599df5badcfc1ab81c897948"></a>

## Direct properties — api_rate_limit / 3ab8d2bafa04 / 3

- [api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2): complete subsection reference.

- [bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596): complete subsection reference.

- [custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-c715e7544dcff217a1f6dac72a2bd057247ccf7107eb5c68b4f07cb612bc936f): complete subsection reference.

- [ip_allowed_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-d57cb4bf5849c5b7d343d37b5411422264ecfd7929bb36f15c9a460c7e4057f8): complete subsection reference.

- [no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-c6953626a1af97898911862c167c8d2f988e255dd59e4089d14bb78c6831fb36): complete subsection reference.

- [server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61): complete subsection reference.

<a id="canonical-fd99edf714ea66a54084c8a6386153b49cd0a68251fea411869fa508f67e6b85"></a>

## Next pages — api_rate_limit / 3ab8d2bafa04 / 4

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-c715e7544dcff217a1f6dac72a2bd057247ccf7107eb5c68b4f07cb612bc936f)
- [api_rate_limit.ip_allowed_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-d57cb4bf5849c5b7d343d37b5411422264ecfd7929bb36f15c9a460c7e4057f8)
- [api_rate_limit.no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-c6953626a1af97898911862c167c8d2f988e255dd59e4089d14bb78c6831fb36)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-9e6b24ec02dd53ecf0edbe6ff0678b336acbac6fd497789e545a8a7c13d66e61)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9b93c337bdfaf9124e4139116f5e69a1ea927211a73bf6ebeedd39eb08341ec"></a>

## api_rate_limit.api_endpoint_rules — api_rate_limit.api_endpoint_rules / 4e24b3becaf7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- api_rate_limit.api_endpoint_rules

<a id="canonical-1aacf6b907187f0d9b750c3b5d10a1af99ac4bac9944d3a672c226cf4378309b"></a>

Type: `"list"`. Computed.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-c2b9a783e1b4b87e78d5e78d4f3af583af263d8a98890065419e772b3d1d4146"></a>

## Direct properties — api_rate_limit.api_endpoint_rules / 4e24b3becaf7 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-006.md#canonical-1d3bb6f675cce707689acbc68ea42081a4b4a25bda7b4e9460d413b83454233c): complete subsection reference.

- [api_endpoint_method](data-sources--http_loadbalancer--reference--group-006.md#canonical-1816e0cddcf421267981f073199813d9f9072c7f328f9cfad9320ce73cf0643a): complete subsection reference.

<a id="canonical-035a61cb6dba34dedfec290cdd8db2e1b78da5a7b0c4e11e6f02d2b448e80117"></a>

<a id="canonical-b701b8dd0a45164c16880cb8fb3b1acc49767fe40ab7e50cf0bd47eb21f246f6"></a>

## api_endpoint_path property — api_rate_limit.api_endpoint_rules / 4e24b3becaf7 / 4

Type: `"string"`. Computed.

API Endpoint. The endpoint (path) of the request.

Upstream description:

The endpoint (path) of the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0): complete subsection reference.

- [inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-4142a078c5f22dc8f4209015b3f12a8aa0a063657fdc460771f360eb82573bb9): complete subsection reference.

- [ref_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-f6af1af010e2d67da04cd9879f8ed6b8e0a2bdb57a14b5f239928fff0784581b): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670): complete subsection reference.

<a id="canonical-ff5013a3260f9d2244ce9aa17e5f113067c99634c976e6ac326517cf01d31136"></a>

<a id="canonical-843bae2f8432c8fc53a203708f8204388f541bd2f90c24f33da699fcdad61722"></a>

## specific_domain property — api_rate_limit.api_endpoint_rules / 4e24b3becaf7 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-d8dbad2cba87ed9ea82d26ce88383ef41d86cebca3dfa8c98f4672d7a8b92fa2"></a>

## Next pages — api_rate_limit.api_endpoint_rules / 4e24b3becaf7 / 6

- [api_rate_limit.api_endpoint_rules.any_domain](data-sources--http_loadbalancer--reference--group-006.md#canonical-1d3bb6f675cce707689acbc68ea42081a4b4a25bda7b4e9460d413b83454233c)
- [api_rate_limit.api_endpoint_rules.api_endpoint_method](data-sources--http_loadbalancer--reference--group-006.md#canonical-1816e0cddcf421267981f073199813d9f9072c7f328f9cfad9320ce73cf0643a)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-4142a078c5f22dc8f4209015b3f12a8aa0a063657fdc460771f360eb82573bb9)
- [api_rate_limit.api_endpoint_rules.ref_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-f6af1af010e2d67da04cd9879f8ed6b8e0a2bdb57a14b5f239928fff0784581b)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1d3bb6f675cce707689acbc68ea42081a4b4a25bda7b4e9460d413b83454233c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e80714f7a1bd8aef9fe935cd0db70e0acf01e3ea23ce8a0c732a6857bbfb9c0d"></a>

## api_rate_limit.api_endpoint_rules.any_domain — api_rate_limit.api_endpoint_rules.any_domain / 0c1801b33fab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-92a52dcdac2844b762d292e3c797c1873d66532961d4526e229ccef4dbec3ccd"></a>

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

<a id="canonical-5f2b57901f16f9726d42a42d067bb2e01bbcbfd5e322624985f07783a65b5ca3"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.any_domain / 0c1801b33fab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-925bea547635619219a147aef077c2c187920e6f03b19c916ec593e1ff72a781"></a>

## Next pages — api_rate_limit.api_endpoint_rules.any_domain / 0c1801b33fab / 4

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1816e0cddcf421267981f073199813d9f9072c7f328f9cfad9320ce73cf0643a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb3335a306eddf8fe87b8b3108a8e859bed8b090044a183ed0294d08e069e180"></a>

## api_rate_limit.api_endpoint_rules.api_endpoint_method — api_rate_limit.api_endpoint_rules.api_endpoint_method / 124d38c3bf7e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-e5bdc27532e5b36dc386ae171325077e38ecbc5f65f0bf1bfcbed6cd81416708"></a>

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

<a id="canonical-7626defe71e29d982e01ffbe858258748f18046f734ad575176830cb8652f77b"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.api_endpoint_method / 124d38c3bf7e / 3

<a id="canonical-b8be7c12a834ca04821353e402a7de3012455f7008dc4ed6aac33be60045c7a8"></a>

<a id="canonical-53182a03afa71196d3c6491e9388b237134355b5f8d5e01d17da4d11df660ad8"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.api_endpoint_method / 124d38c3bf7e / 4

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

<a id="canonical-5b5b156a7bca4eb2aa49a657249f26cbef46f083ae8fc70a30ab0e47817d6411"></a>

<a id="canonical-aedc1e49bcca1a99081affcd1f45605c25c16c729fbc769067dc910dd50c120e"></a>

## methods property — api_rate_limit.api_endpoint_rules.api_endpoint_method / 124d38c3bf7e / 5

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

<a id="canonical-80e1073d26b790b4d0482e8184dcc3ad44f64ae3933054896c633a60ef9624bd"></a>

## Next pages — api_rate_limit.api_endpoint_rules.api_endpoint_method / 124d38c3bf7e / 6

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e1798490460688f7575f46d193c8ae7b3086f46bf47c21fe6b6f3eb9d853764"></a>

## api_rate_limit.api_endpoint_rules.client_matcher — api_rate_limit.api_endpoint_rules.client_matcher / 1496a2d0b6ff / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-6db83a0e1ff5c5372aca27ed553844991d54e83a05ec1f2003afc80b50b0d67d"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

<a id="canonical-82c5f83a2c8b81173923415b0c57238202d68bc69f833231a14da7480d663ac2"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher / 1496a2d0b6ff / 3

- [any_client](data-sources--http_loadbalancer--reference--group-006.md#canonical-4e1564639c8bc842a176ebb1b9838f93539c03fcb09b27e361685f79e6d9fb0d): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-006.md#canonical-5b2ee7a88558d8a69e5b8f55a64f6f8ab82a473cf2dd2e6d38f5d056ac7c122f): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-a8117e93b3836718878462e09ed52e63cb26ee5b4403a09d704614b7cc752bd4): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3cbd89c0e2263048bddf6dfc15ccd2162066d23ab62edfb63835ab65b744d1ee): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-006.md#canonical-6d22f27c50796667dfafc46d279be20524064751ae2062866b1e521772c5245a): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-342c2520ae6643eb518924a585730ea80185d885dbe7c1fc7ec13f18d59fdd0d): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-4446c3987e3088f434cedd87375f2146f16eee176afbd8aef48b687eba7f1dc9): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-e2b15f1453a11d911f6bf79b1e09b6004de37d0a798523018a5c344a1cb13b8e): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-d9461102144bb52a1aceb64da206d1f9807be67415d842a0cc94f1d668d7b603): complete subsection reference.

<a id="canonical-02f9debbb58f71b13c672ff3f58317772431addc2bcb0512da83a389bc7247cb"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher / 1496a2d0b6ff / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.any_client](data-sources--http_loadbalancer--reference--group-006.md#canonical-4e1564639c8bc842a176ebb1b9838f93539c03fcb09b27e361685f79e6d9fb0d)
- [api_rate_limit.api_endpoint_rules.client_matcher.any_ip](data-sources--http_loadbalancer--reference--group-006.md#canonical-5b2ee7a88558d8a69e5b8f55a64f6f8ab82a473cf2dd2e6d38f5d056ac7c122f)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-a8117e93b3836718878462e09ed52e63cb26ee5b4403a09d704614b7cc752bd4)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3cbd89c0e2263048bddf6dfc15ccd2162066d23ab62edfb63835ab65b744d1ee)
- [api_rate_limit.api_endpoint_rules.client_matcher.client_selector](data-sources--http_loadbalancer--reference--group-006.md#canonical-6d22f27c50796667dfafc46d279be20524064751ae2062866b1e521772c5245a)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-342c2520ae6643eb518924a585730ea80185d885dbe7c1fc7ec13f18d59fdd0d)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-4446c3987e3088f434cedd87375f2146f16eee176afbd8aef48b687eba7f1dc9)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-e2b15f1453a11d911f6bf79b1e09b6004de37d0a798523018a5c344a1cb13b8e)
- [api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-d9461102144bb52a1aceb64da206d1f9807be67415d842a0cc94f1d668d7b603)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4e1564639c8bc842a176ebb1b9838f93539c03fcb09b27e361685f79e6d9fb0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-965999c9160e45e09eecd0be588a8de2cc363eb4dc6fd8c8125db774c8f718f2"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_client — api_rate_limit.api_endpoint_rules.client_matcher.any_client / daaf7a41a9cd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-573ac1de73242c01ca29e8a54238a28e6376391ad51b143a96e59098cc5ffff5"></a>

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

<a id="canonical-1c7eac1493de928b615074d130c8fc637437a5df7d1f0a3aa8264e21af6820a6"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.any_client / daaf7a41a9cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ad4c0cc9c08d5180ee9499480c85671379473a483a99664382acb8e5877d1c3"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.any_client / daaf7a41a9cd / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5b2ee7a88558d8a69e5b8f55a64f6f8ab82a473cf2dd2e6d38f5d056ac7c122f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1976bb917265e03f5fd413171a2bc2fdba09baf981fe77994f915b9e100d1b31"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_ip — api_rate_limit.api_endpoint_rules.client_matcher.any_ip / 345ae2c0541b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-88f76a3396112d60ab0fd9c76beefaf91c0ca6bc7fdb8df9e1921e3b41551731"></a>

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

<a id="canonical-cd22f0d51f8347c6a4e9ac46ff6d30a3b828d86d23f75047c8d34e1335ba62e7"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.any_ip / 345ae2c0541b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ad8e59b2c5802c429c1f3740c16efcae17275b5bb337707748cd0867ffd7f3fc"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.any_ip / 345ae2c0541b / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a8117e93b3836718878462e09ed52e63cb26ee5b4403a09d704614b7cc752bd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f4a3710877a7a5bd3807dfc1c647335a18de781a227554a1b560e6ef7520ce7"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_list — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 79c55c9e2c09 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-00b118be67be94b58332c62c6877d16919bd383479a32202712d5d63213d8fbe"></a>

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

<a id="canonical-77bcf0d71ddfcec8412eb03bc5daaa950b4c87da53948c06320280c4529171f9"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 79c55c9e2c09 / 3

<a id="canonical-f27d3bef6774396b3ebdb670f52ebd2447f5d4868d2e504a71fbcaa255701629"></a>

<a id="canonical-7d5fdc2038930e09c5bc4d8f7d8833ffe9d150ae145a4e2e4fe39a0cb55d12b7"></a>

## as_numbers property — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 79c55c9e2c09 / 4

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

<a id="canonical-7e0e05fd243993fda93ece13780d31f0b395eaa2bbf1526af83e88c972209686"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 79c55c9e2c09 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3cbd89c0e2263048bddf6dfc15ccd2162066d23ab62edfb63835ab65b744d1ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfc1933a3bd80760e2ada29ad7709661755d0cb90af9c76ad4a74b50a6988b77"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher / badc482ebb7d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-fbd017ceacdbae2aa6e5fe7e02f01e6bd13b0e7c0c8125ba12130e9b8e625c51"></a>

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

<a id="canonical-0f8ecd9903e3f0d03010f43780709aee97dfeac966829a46caf6b8d8564aaa28"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher / badc482ebb7d / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-e61ea7686e7f5d577dd3f7c7b28817c72d017e61a3aa9e0e224eb52c340c9bda): complete subsection reference.

<a id="canonical-5498bc10abeda9564ae3fa991c5042aa2ef5f8f844a08c3a24dec744073040f1"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher / badc482ebb7d / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-e61ea7686e7f5d577dd3f7c7b28817c72d017e61a3aa9e0e224eb52c340c9bda)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e61ea7686e7f5d577dd3f7c7b28817c72d017e61a3aa9e0e224eb52c340c9bda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12630810fa8224609b9a788bcba5b8c12e3183edbbc5180b23efa6f4c56c30ef"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3cbd89c0e2263048bddf6dfc15ccd2162066d23ab62edfb63835ab65b744d1ee)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-effed9255b56815c42289a79b54a2b262f5fb0c25877ad92e16c9ab18558cc06"></a>

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

<a id="canonical-8c85ebe75cf8cb7c95b855e28d1fd23b8811621c23fd9223e402d028faaaaaa8"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 3

<a id="canonical-5fefa4987d270d7b8c9142198613c80eff09d941484753eec1fe47ea5e2662b0"></a>

<a id="canonical-907d457e3bb9c1572cacbdfaee501fe874cbf706a66cb4478da428fab0a12091"></a>

## kind property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 4

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

<a id="canonical-af3941f1dc8b3c152b05835cf08c23b918199bf5e4c3aae72af5a3529c95112f"></a>

<a id="canonical-cd93c0bc540647fd14077991cfa3ef8847865004b27e2fe85bff5d2b19991d7b"></a>

## name property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 5

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

<a id="canonical-7a67ad6e1d84177535371d5913dfdb4cbd93373ebb341b53b95c2f5fa042f690"></a>

<a id="canonical-669a04459ee5a38bdd2b07bcb9d46bdee5ee14892f6706e860bdb2e87e60b04a"></a>

## namespace property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 6

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

<a id="canonical-cfa0ee6ac2d4c6ec0d09def0caa25fef3e9644caf13400cf1d63fc7c0a0b8b05"></a>

<a id="canonical-50bdf9e03d6a14b6b96d6e649ae71c80c5e453be93eca20201039c17d95e3d3a"></a>

## tenant property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 7

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

<a id="canonical-896870cfef39b8544d5c37d3166e666065ff0816742450d84e0c2cbeb5e02b16"></a>

<a id="canonical-ce9afde4a912729165b22b5b025c76e2716491b963a2c276d14115e7c014ba0f"></a>

## uid property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 8

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

<a id="canonical-9643535bc26345c8e0d56cc9a3bfada9f4e04ee8718d362fb3763372e5a16cdf"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / c6c069a2abdd / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3cbd89c0e2263048bddf6dfc15ccd2162066d23ab62edfb63835ab65b744d1ee)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6d22f27c50796667dfafc46d279be20524064751ae2062866b1e521772c5245a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4c09b941c5f8e441e91b2ee4466ddbf488d37f2eb81ef61c4910de447b9c46c"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.client_selector — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / 7302bd198c46 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-ad43c44400182ebe88fe9e05924631eab77e0f0fcb9f84cc650efc1128d06af3"></a>

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

<a id="canonical-a726694762938db163580833dc58327bcc88b9c0189e6cd282ec39f527b6eeb3"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / 7302bd198c46 / 3

<a id="canonical-afeb729f59484ede2995f167668a3761c4e833d176165f5adf7c813fa9d7ce7f"></a>

<a id="canonical-372ca623cf99167d61cf9971718ce02aa251577162ae0dcfce7aeab83231d0d4"></a>

## expressions property — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / 7302bd198c46 / 4

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

<a id="canonical-ce3b26bdaaa3222761e0b2ca20ecf266408fecc2041c08a2cac0888c02d874b4"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / 7302bd198c46 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-342c2520ae6643eb518924a585730ea80185d885dbe7c1fc7ec13f18d59fdd0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2574e435ec0bfd799b7293eb3f381829c580d86fcd0173a84e748bbea48be6c"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 6a721b4f2128 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-6547136b8eae2b1c153e2a9ed05d02e190cd4ecbd87decfc077bb40ec325518c"></a>

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

<a id="canonical-ec030c385399a57869c7e06b17e97ee4654fefb4f0f9a0939062795e7946242f"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 6a721b4f2128 / 3

<a id="canonical-70e7c4b94a8dfbf69051100ce02a8acf06f3e81875342b97b639df2d3961995c"></a>

<a id="canonical-f5267f840c8e2f42ee77300c283575f8337ad6943fa248d45e37a24aad2c84d5"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 6a721b4f2128 / 4

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-4412d0dd50da14aabc98c17f95b3fb3ca0716c475e382619b4cd9e9f7ad3c80b): complete subsection reference.

<a id="canonical-87133c8f6972b14fbfae6516790f7fdfae4ab4259ed9b36da4eee2fe8eb1150b"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 6a721b4f2128 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-4412d0dd50da14aabc98c17f95b3fb3ca0716c475e382619b4cd9e9f7ad3c80b)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4412d0dd50da14aabc98c17f95b3fb3ca0716c475e382619b4cd9e9f7ad3c80b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b1800cb3bf4e6d00103f4536ad3afd17e43c65c0512de91e574f4593a924b08"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-342c2520ae6643eb518924a585730ea80185d885dbe7c1fc7ec13f18d59fdd0d)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-b29f9c6b3ee56ee017aaa2b85aeabc673c629ae2f8fdb4fee6124621be294baf"></a>

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

<a id="canonical-29f32c3ff0407f36ee3f4532b8d34e29d800964a472f4ab246da57c4a78222ec"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 3

<a id="canonical-810d7413a0b23bdf35998e0dadf315317daa3469eb82f74bbc9f156dee16398d"></a>

<a id="canonical-4060b61fedee7ddc89b8299cb80c7e9852222ff716bf14cbd61f12ea7ffee063"></a>

## kind property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 4

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

<a id="canonical-874085ba5f89dfaae00ecd1d1e46bcaebab3c293441f5a4e23af04eac448d667"></a>

<a id="canonical-86a58bdbe9200c2fc37e5ccdc115c8672f5ac54aa2c460176ca082123dca1e79"></a>

## name property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 5

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

<a id="canonical-b5c10f9d49d5e72aa4817ef5f7620db5fcf1a0bfff3b786ce12658336610fd59"></a>

<a id="canonical-0d2173455840d1d3e0263f112686b6ab9de69647c012100c08ce80777a451189"></a>

## namespace property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 6

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

<a id="canonical-e6dbffa633d61b368761e68a842466fb490d09cb046e6dd8b59f933024d0bf48"></a>

<a id="canonical-c6f50d6a50754c1ba1c56328849caad222bcd0e03f89e63dcc0614bffcf41e42"></a>

## tenant property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 7

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

<a id="canonical-201c16312de1e1bcc30418e7d869989ac2d7cba1ede108bca14a8ad82d1dffd6"></a>

<a id="canonical-345bdc19489f541095c3b6d2ee342703f982179d3a59c536b3ba4d9de701021d"></a>

## uid property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 8

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

<a id="canonical-32c34da1df89d9bcc3eacd973a43e0bf3af7b39fbc3e941734e8754f7c3fde13"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 9224aba9bd6c / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-342c2520ae6643eb518924a585730ea80185d885dbe7c1fc7ec13f18d59fdd0d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4446c3987e3088f434cedd87375f2146f16eee176afbd8aef48b687eba7f1dc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ee4f9b6bfc0c198ec816198b3a8603902ff2ade51a306de5b81040a97f794f1"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 02bf1eb69bd5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-f3e7d0e5f14c09d47e8cd2b8cabec5cf32ac933778798aae07e005636226040d"></a>

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

<a id="canonical-8ee243bfa60892ab150b9aeddae4d5a261b7aea1bee8358a22c5537829cc11bc"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 02bf1eb69bd5 / 3

<a id="canonical-1765a1d45b3f8e98d9043ed46433dfda4a8a45a784552cc658b3cb83005175ad"></a>

<a id="canonical-9116e8bd4e94dc7cf496b789362cdcef412d223c17a123b5375d4a15dd52fcbe"></a>

## invert_match property — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 02bf1eb69bd5 / 4

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

<a id="canonical-51afd4b62d7c6d94277d73497db85867051f4e86440cbea83efeae68ba39386b"></a>

<a id="canonical-db98f2078436df0714be1abea0ca566827a9ebbb45c5fa7fe1e3be78f9b7825b"></a>

## ip_prefixes property — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 02bf1eb69bd5 / 5

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

<a id="canonical-e8bb0f2d293a32e134a4e2c65064e76ded3ff4089cd8734571b00159bb2d3a68"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 02bf1eb69bd5 / 6

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e2b15f1453a11d911f6bf79b1e09b6004de37d0a798523018a5c344a1cb13b8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc28091fd10f9bbfb3e50db098b4abc1793f322f36105db2580e95c9cea2ef8f"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / c2a58d5b52fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-b8ded0436495526b75c4cc32a04bd97fc74dcd19c6481355d79ba404a0682f6c"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-8133b647bb1ade47e41852b9bacad0da9551331c861a624199072470db727aca"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / c2a58d5b52fb / 3

<a id="canonical-1d874cfa92b5119ff001a443743aa8cbb2b1bc62c98c9e27fdaf772342319d10"></a>

<a id="canonical-09fddcd64a0d95c9ebd3e18d81ba696bd4510cf6b52375fd4a994a252180be2d"></a>

## ip_threat_categories property — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / c2a58d5b52fb / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ec58837a29f0100b00a94d6ca70ad5a59eb7191fc5813581a72fc18c4cc5ea78"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / c2a58d5b52fb / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d9461102144bb52a1aceb64da206d1f9807be67415d842a0cc94f1d668d7b603"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-416f963ee52b1978c41b2927a81353895ef477dc5824058b52b34e42c2c85e9a"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / a84d8d6ffa58 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-856d8ed3bda775af769a110bde24ac9062004244aa1049715dd1fc7e637de039"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-385fc17b349d6a71842b379437cfece27b689f2f960f7c2368b776d7b1a04a42"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / a84d8d6ffa58 / 3

<a id="canonical-b28ad6f5ae3409c42b3115d35fa0aa8c792347b226c10c1d9f4a34e1b60b42b2"></a>

<a id="canonical-a70933ba39899f0cdb289f6e31b7efc95add706aa1e805fb4fae166457848f3a"></a>

## classes property — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / a84d8d6ffa58 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e4d3acc91b507c544e0529c8f2c35a85e04887e1048c15f83154a3ec11cc6034"></a>

<a id="canonical-ad1bd28ca91cb0b77b132b8569222f1683431715e61973aafff2568510a5ef34"></a>

## exact_values property — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / a84d8d6ffa58 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1ec798dd6cce2928eee98f40213d81edc05665239c635563f9277ed0012b89ac"></a>

<a id="canonical-c1656378e3f2c126bbd1c4f8eb8053c9efa02cbcb720e8b68353929208f09a51"></a>

## excluded_values property — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / a84d8d6ffa58 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cf48cd69ffe0c7a3a193c2de3aec17e109b729513160d3b18c4ecf642341a111"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / a84d8d6ffa58 / 7

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-11b837fb69d2e3772c8dd713cb5547a9cb0dbc994749b63b20a88436d3a3b4c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4142a078c5f22dc8f4209015b3f12a8aa0a063657fdc460771f360eb82573bb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-064c9e9283b2603ae05a3f92c7ae6f94659dc859aa749a437a3c6a4f71039337"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter — api_rate_limit.api_endpoint_rules.inline_rate_limiter / bd07f5a5e95d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-d14895005e59bde0edc3d657794cfb3b82ffae55ad2c0dcb79a933aa6b4e2138"></a>

Type: `"single"`. Computed.

Configuration parameter for inline rate limiter.

Upstream description:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

<a id="canonical-e6deb596aef8f9eb0f3a6531b53112918b6250a3b355acba88d1d29070f8b85c"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.inline_rate_limiter / bd07f5a5e95d / 3

- [ref_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-d1e06bf5d2a11168a43dd7ce966f9daf0990e92132962fd809a1adafbaa7cbe4): complete subsection reference.

<a id="canonical-50d70f1f18f4534534119075b329f27db381b8867b7989bd21d1f6fe100372b2"></a>

<a id="canonical-d3f345676d11cd39c0e7355e350af46c9a9cf0f6558b49131b6b45882f036896"></a>

## threshold property — api_rate_limit.api_endpoint_rules.inline_rate_limiter / bd07f5a5e95d / 4

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-4e74c8b34cb037ca11f96810338577e553c9244a6e490f745fcbd8c72120f99a"></a>

<a id="canonical-a544c570b1501deb8c0cb05a49d4fe60476510292ec043e2b84e5ea95f5b78c7"></a>

## unit property — api_rate_limit.api_endpoint_rules.inline_rate_limiter / bd07f5a5e95d / 5

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-49f1f0b9c732c022995e584ceb11e4fbd8466b58ed68388e00683db62ef2117e): complete subsection reference.

<a id="canonical-edaa1d4df8a02eeaaeeab1c05c662d5b06a98a9c91f6a4f6fc8c311a8b408973"></a>

## Next pages — api_rate_limit.api_endpoint_rules.inline_rate_limiter / bd07f5a5e95d / 6

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-d1e06bf5d2a11168a43dd7ce966f9daf0990e92132962fd809a1adafbaa7cbe4)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-49f1f0b9c732c022995e584ceb11e4fbd8466b58ed68388e00683db62ef2117e)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d1e06bf5d2a11168a43dd7ce966f9daf0990e92132962fd809a1adafbaa7cbe4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a1209746cfc276e1e163b23c293f181b6f94cecb2a45db28237e40d9a5c6567"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 2855d4a4c848 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-4142a078c5f22dc8f4209015b3f12a8aa0a063657fdc460771f360eb82573bb9)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-79e2194d3aa77e32eef1374da80cb1df3f920914501b77b335cc7475a5333d4f"></a>

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

<a id="canonical-474c22f27255c148e3fd57cc699715870a4a278c42ac48160597c2e804db28f2"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 2855d4a4c848 / 3

<a id="canonical-5af62397440759395909c8383bae20ecf62953031c65974a46a3932eea771a42"></a>

<a id="canonical-6dd13f2fc1425a8f16926bf357b7181f7de203f1b7fb19a57fb7ffef50913827"></a>

## name property — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 2855d4a4c848 / 4

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

<a id="canonical-43aad0ace21cc55c4929b21fde2662378393bafa6bac792e093987337d89d129"></a>

<a id="canonical-37d9f1ab1e2a2b77401a9e1f3db594415f58cf6ca6e02b24e0c53bbba098df18"></a>

## namespace property — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 2855d4a4c848 / 5

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

<a id="canonical-354216623acc9a20a7e01e6446cc6f53e5e9df6f687021d0bcc6d9dfa06534ea"></a>

<a id="canonical-cf21ae11f00483253db06dfc2b5198c8b0c9f1c2bb35c47a07f308383625c171"></a>

## tenant property — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 2855d4a4c848 / 6

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

<a id="canonical-0fcd376a4a4a10e563ac91f60f37ff0aadab3e9f2b3aaaa314b041d157fb3a02"></a>

## Next pages — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 2855d4a4c848 / 7

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-4142a078c5f22dc8f4209015b3f12a8aa0a063657fdc460771f360eb82573bb9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-49f1f0b9c732c022995e584ceb11e4fbd8466b58ed68388e00683db62ef2117e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44cf0c2aed4d2856f1508ce52696aaa5a0a1b4e765ab7a5fe512cbb2da5c71a6"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id — api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id / 97e2ead35c9a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-4142a078c5f22dc8f4209015b3f12a8aa0a063657fdc460771f360eb82573bb9)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-89f9efcb83646b917d19f8a834111ea554cad24eba4957284c86ea769f1b4b0d"></a>

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

<a id="canonical-132609f69061fc3371fded00dd85cfee274a3e34c6c0f3e376010a3897289e91"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id / 97e2ead35c9a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-70acac02af20046002ffab9e02c982018365b26550cc00125fbd8b98a4be19e6"></a>

## Next pages — api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id / 97e2ead35c9a / 4

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-4142a078c5f22dc8f4209015b3f12a8aa0a063657fdc460771f360eb82573bb9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f6af1af010e2d67da04cd9879f8ed6b8e0a2bdb57a14b5f239928fff0784581b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5390157a53be6398c750e6c033efd86911327fdb83999bab31f54e63d14a09d5"></a>

## api_rate_limit.api_endpoint_rules.ref_rate_limiter — api_rate_limit.api_endpoint_rules.ref_rate_limiter / aaef8c4e356e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-0d84b595b75b03c009a84ffdf7d88beb6de78828c86e5824d56c977937c62485"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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

<a id="canonical-8a20503276c6ca553adc26923501a9eaf127801cd4d4b3943ad8880ba84e89ee"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.ref_rate_limiter / aaef8c4e356e / 3

<a id="canonical-d6c5bbfa2b10ea0e1eedf3dbc58dce5bbdcfed9f1f7f3f51824b59b232e18bb4"></a>

<a id="canonical-755d2e0442c6b8033ef14e1401de92de60f6aa41af928656d8f1481f1de93099"></a>

## name property — api_rate_limit.api_endpoint_rules.ref_rate_limiter / aaef8c4e356e / 4

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

<a id="canonical-2c921e8b4c5b7b4d77da16383a67162791460f7b9aff9242467b7e9b02d10139"></a>

<a id="canonical-8a75cb84230e4e6c3282798b39f94fda205e4cd64a2da2c9cf004e1d338c7025"></a>

## namespace property — api_rate_limit.api_endpoint_rules.ref_rate_limiter / aaef8c4e356e / 5

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

<a id="canonical-6603cbf9a404a132117fda3fd8fbb1d94c41be491a625bfb9edc109368fa7d64"></a>

<a id="canonical-78b510adf04b2b740d9fc8db2cdd526c746e3e9e43676a8cbd1285947a619322"></a>

## tenant property — api_rate_limit.api_endpoint_rules.ref_rate_limiter / aaef8c4e356e / 6

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

<a id="canonical-af164af173d26f4c729abf74f2c21c0c31a32ce5f6daaf68b7466d0161c25a9f"></a>

## Next pages — api_rate_limit.api_endpoint_rules.ref_rate_limiter / aaef8c4e356e / 7

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6af3513b5521fddb4ab89b81072eb8d87459696c5d8994938d173548386091f8"></a>

## api_rate_limit.api_endpoint_rules.request_matcher — api_rate_limit.api_endpoint_rules.request_matcher / cce7c691fdfe / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-9e6ae2bd41c6fedb41d021d3676b48d05159ec87ce09401a5885c4e39339080a"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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

<a id="canonical-75af60f22a999b106e4768894734fada2726139d2db813056679cd659d878f8f"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher / cce7c691fdfe / 3

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2): complete subsection reference.

<a id="canonical-4b603b19c670acfe9f1d62316d5a2d7f96eea9d2312b53d06b10df1b6daabb05"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher / cce7c691fdfe / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd553370069195da89ada68b7b30864835778179fa8a017fd870789f23ace9c1"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / ebb5d4d95da6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-feb42130d689dac315e46ecce2242397b44f569fd881e5206f87749bcc1efe7f"></a>

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

<a id="canonical-2019113214edbdbce07962a42d05e7705a327aa66b79f21b13fd5dccbf59496f"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / ebb5d4d95da6 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-f4b52cb527f29fe159468269e11b9f63027f3e1508db7db6cc4c521597faeff7): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-e9961551e66ffc0802ab0c34bd7c3b61683744e778fe0deb14399668460c5471): complete subsection reference.

<a id="canonical-0ac95f4b2dfbaa39e06da9bc704467de5917ab7b075e0480ee3f22bc24ba7006"></a>

<a id="canonical-dd673426c2bc360f02aa38fe49e8786fac7d9fedba30747f8bc7eda4bb15d203"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / ebb5d4d95da6 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-6c9b5aa9eab9fa7c248172d96f91f8779c814949475f35178ca214e5158067be): complete subsection reference.

<a id="canonical-331a993c3e40a6fd103c2a7eda89a38abda30950d69d8c1cfaeb17fc760210b3"></a>

<a id="canonical-94c01993e5069235d065dc817a5fc71d17afb7a6a1c5030f2c046283851b09a7"></a>

## name property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / ebb5d4d95da6 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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

<a id="canonical-1ef953f9adf2e852f6f6fe66267c37662d0d25d88b110be3c08c3147d5707ac1"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / ebb5d4d95da6 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-f4b52cb527f29fe159468269e11b9f63027f3e1508db7db6cc4c521597faeff7)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-e9961551e66ffc0802ab0c34bd7c3b61683744e778fe0deb14399668460c5471)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-6c9b5aa9eab9fa7c248172d96f91f8779c814949475f35178ca214e5158067be)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f4b52cb527f29fe159468269e11b9f63027f3e1508db7db6cc4c521597faeff7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7af70a860824375e1b54bdc443742128e7b9ebde7860732f8b61b5184afaa3d0"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_pres / 8255c9ea2963 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-e1fa3568c22a2c50194e598e09e470f9fdbd37be5d046e86b405cdf513c567cb"></a>

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

<a id="canonical-f63a7208e49b6bd1c02e5b2fb4e518084f675225881fc96116757df4d265425d"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_pres / 8255c9ea2963 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5dc9327abacfdece9fa190fca2ddd8f9f008da76bf12f051a8503e424173e353"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_pres / 8255c9ea2963 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e9961551e66ffc0802ab0c34bd7c3b61683744e778fe0deb14399668460c5471"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ab5bb325c38b1420eba58324815ae75a1d9d05159a795525a623a1c2353926e"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present / d295b1ecc5f4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-d2e606fc405100c60194bec86ab2c841cd29b29113814ba3bbbfd9bc24ebffbe"></a>

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

<a id="canonical-e04728b1ab00d8c7d9d03df5536a9649b922402e6b6a02259b387cbe5479d77a"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present / d295b1ecc5f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed08588e2ed3ec4de886f9ac01dc69a67f5be15078d64f41c6ea0d553c37712a"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present / d295b1ecc5f4 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6c9b5aa9eab9fa7c248172d96f91f8779c814949475f35178ca214e5158067be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd6c91974169055aea74146f1e041abe823aeff01f94be8292f319e1ad7e5006"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / e8e31b8dbbc7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-992492bd98b68cb2a319b009c0edd8bbe327645323efc86b7e1770f70aeb8edf"></a>

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

<a id="canonical-79e1d834041bc7d5c5ae96e40c6776f5f3d04e5d0aa710c6432c6f0c29a09982"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / e8e31b8dbbc7 / 3

<a id="canonical-8729252610309ba933ad3734979e78b0bd4b2bc0d8db33078ccd13658a590917"></a>

<a id="canonical-34160419dc65a14a9ab4d931fcf716eafd8c611af4ef7bf111af8a7f7f251187"></a>

## exact_values property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / e8e31b8dbbc7 / 4

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

<a id="canonical-7d8c2160f326cb6823a5c6a0b822201346e4e16031b1c0fef476f35c88bce477"></a>

<a id="canonical-84779846448495e85780702f68bc1ce81439a8948be36f61caff93d26fb64793"></a>

## regex_values property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / e8e31b8dbbc7 / 5

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

<a id="canonical-7a94da997ca74219c773e89a23bda0201653f3a55a1c45667c3378a5b401ca8b"></a>

<a id="canonical-b1718c4239e12696305433bf38a08eec1daecffaed534138dc224491c26fa0cf"></a>

## transformers property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / e8e31b8dbbc7 / 6

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

<a id="canonical-3522bc9b127820cc65bf44960d6a04c8ded46848201b2d4e835699a0493f2def"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / e8e31b8dbbc7 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-4fa59e3fe7e470bf4f7301634cbe2b531c4e5944bd2ab3ce8c160856d09e591a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a1117ef5a19f81c5ab1fa4dc137e694576d705aae1949602a33ea13d5783bb5"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers — api_rate_limit.api_endpoint_rules.request_matcher.headers / 849e258dca86 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-2db3f91ea1a08001a43b47de0f0e869d7009ca59f215e9e0349cf90be1178982"></a>

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

<a id="canonical-ef268de9322f6f295f8ff76f9110aec4cc8884d4ad024f3530bb77af24957947"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.headers / 849e258dca86 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-9c182e42d49ec5e6ca38f90750fdc9e2f88245b56fde39a7228897fa6c4a240b): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1aa89cb235bfd1c469dd41cc8a681f4a682b7b0e210333113a9611b4ea002f80): complete subsection reference.

<a id="canonical-c9b76436bf9cea493fefec603a3d4764dfaff3bc5060b1d0d09e8ff4dbf870b9"></a>

<a id="canonical-5305e60a3e8fb5924e747b6896780b8b63d9788bb49dd83e82d3915fe078ccb0"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.request_matcher.headers / 849e258dca86 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-73b6a80eafb77605a9b8b5aea9d6960b9f4b19047848bd04f4bc35e5b59478a6): complete subsection reference.

<a id="canonical-5c637a2f2cd89aceeeb421fab63b0b130b169094b2bb5f3138db4aa6ffffcb6e"></a>

<a id="canonical-ee12529d7fe2a1c26904167ad31be99354623448c6b77834d60645f8a3d4d8f3"></a>

## name property — api_rate_limit.api_endpoint_rules.request_matcher.headers / 849e258dca86 / 5

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

<a id="canonical-214b75c8d6b909928ed9d42fa18a4d784c775fc606985ddd700368e7ee9eba5a"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.headers / 849e258dca86 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-9c182e42d49ec5e6ca38f90750fdc9e2f88245b56fde39a7228897fa6c4a240b)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1aa89cb235bfd1c469dd41cc8a681f4a682b7b0e210333113a9611b4ea002f80)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-73b6a80eafb77605a9b8b5aea9d6960b9f4b19047848bd04f4bc35e5b59478a6)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9c182e42d49ec5e6ca38f90750fdc9e2f88245b56fde39a7228897fa6c4a240b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a529562faeccc3121536f836ded5c01176db50f11127a75e91d41181ba7972d5"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present — api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present / e548dd736b62 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-c478bb4897436d190631c002f7836937d20013fdb7c44b94150e929005e9299b"></a>

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

<a id="canonical-0c465806e8f5f593eeee1b9896542a32a98c80db68d50dd1ae193450ea299e3b"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present / e548dd736b62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe7e7eb73871e64421b30feebbd27a0d3ef4ede5bd0b76b38e45db1f55df9709"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present / e548dd736b62 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1aa89cb235bfd1c469dd41cc8a681f4a682b7b0e210333113a9611b4ea002f80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a2da5c98328a3b0d89f58341599e51fe97d18fc1a52f537e40df23319ed14f3"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present — api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present / 1c757bf32e24 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-bac0f3ae2d292833b5ecf41c4ac845f3d2ef9ba390f67a4a9614d983a8f8487c"></a>

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

<a id="canonical-fa20b0593120005febf1c3a22b7049758553323b0564edbb5975a53513695b5e"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present / 1c757bf32e24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-00f1f7189af6b0ea605ab81a31d46e6cdf59a08358057ee5801b4f5b179e2923"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present / 1c757bf32e24 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-73b6a80eafb77605a9b8b5aea9d6960b9f4b19047848bd04f4bc35e5b59478a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5792e4f301678fc450f3273789242e94eb66c4b36dfd2a1e1a883228793c1895"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers.item — api_rate_limit.api_endpoint_rules.request_matcher.headers.item / 606f061a0e0b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-ff6f711aa0376448ad9301b13a07e5bdc645884d6c255d2a76d26183ea6f7e8e"></a>

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

<a id="canonical-f22464caf73ac048470607bc3c222de59a984f4bfe592cd8343d041a4e466a4d"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.headers.item / 606f061a0e0b / 3

<a id="canonical-e0271cd30c5494173afab79c6d4cfdf9081925ecd4e1b6270d179251dcdcb34d"></a>

<a id="canonical-098e99b3b5f6f1c79fe95753209f89c2f8aa3a83d2e567668f83a7afbe503e09"></a>

## exact_values property — api_rate_limit.api_endpoint_rules.request_matcher.headers.item / 606f061a0e0b / 4

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

<a id="canonical-024f65335341614c6f2cd1e754cb1399439b269f528b95e1cd21505db3dbb9ca"></a>

<a id="canonical-f31ec3ce9867394f9f8a6da24ab431f42e1e47943dbff703b5b659d5a6430360"></a>

## regex_values property — api_rate_limit.api_endpoint_rules.request_matcher.headers.item / 606f061a0e0b / 5

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

<a id="canonical-d18bdd6eb36eb76afbbebe50314c338a3ad906c39b2f6ebf177127830fb28a2c"></a>

<a id="canonical-cfceb590d0781f7e7593fda19db9c6bad325e3a03331f16e78df457b35bced75"></a>

## transformers property — api_rate_limit.api_endpoint_rules.request_matcher.headers.item / 606f061a0e0b / 6

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

<a id="canonical-448587c9292ce1a12500ee18daf0e2c525b1b7b15bc1c106397dfb2ae302589e"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.headers.item / 606f061a0e0b / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-f5d717eaae77de227f4ddfe945a4d38671bf848a67b95e3bef3ff653a7624cbb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5905444057f20679e5958680811ee33c893290f86a59a5009fc3e952be3dfc12"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims / da28bd77b9ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-0f89d30fa6827649f1aa71ca49815af8b38a03823fccd548172146ec497d7e0a"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

<a id="canonical-5a9db792cd689b9d1fe70f0b4b4e280576e9505a89e77a344273d1db43dd7dc7"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims / da28bd77b9ef / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-a26d9ffe22201144f6dc76c105304d4666d72944d4f41ebbf2fb91b7f700153f): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-f67b6bf8261829748a5c9a99f20f65d1a7de9b59e9ac45ce3222f6574d964d80): complete subsection reference.

<a id="canonical-0dcd0d4381bfe66e559bbfda3f43f5f64f49c267b3ec781c033d022d73e01471"></a>

<a id="canonical-4afd268a12258e8408a3428637106b3b680666b2d9f6448030f93c07c4bdcd78"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims / da28bd77b9ef / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-b9f4c1587f31a2d4bbca087510f70ae93e27d43ff67f16f0221fab1a6c059690): complete subsection reference.

<a id="canonical-a18bf0572f437fdd26d172dbed5f92799ec09ecd8419e024ac0cc8dee93b54ff"></a>

<a id="canonical-12b3dce24da196e31f789a51f6cd5fde3bac7cf6e2041e4af23f75aace544cfa"></a>

## name property — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims / da28bd77b9ef / 5

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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

<a id="canonical-3a8a5cae063571e49ba21139428807d39bfc1f218176d4900479c1b63cd27f81"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims / da28bd77b9ef / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-a26d9ffe22201144f6dc76c105304d4666d72944d4f41ebbf2fb91b7f700153f)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-f67b6bf8261829748a5c9a99f20f65d1a7de9b59e9ac45ce3222f6574d964d80)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item](data-sources--http_loadbalancer--reference--group-007.md#canonical-b9f4c1587f31a2d4bbca087510f70ae93e27d43ff67f16f0221fab1a6c059690)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a26d9ffe22201144f6dc76c105304d4666d72944d4f41ebbf2fb91b7f700153f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-459d9deea42d7d1bbb00d461dbf3ad9f0cddf2db4d43d39dbfceb2b392bfd70f"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present / d5c1ce1ecea3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-85601584255a3886616b47d2d5ab6dddae4ab50569f37e312c895863d6bd87c8"></a>

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

<a id="canonical-ef950b8a91cfea4cf34278ce808fd5e08dd904a705dbc934c19da081504a732e"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present / d5c1ce1ecea3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c912b4d33f0ec90053e479deb5c4bc1d5ded5c6bb6364b366fa15719d825bcc"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present / d5c1ce1ecea3 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f67b6bf8261829748a5c9a99f20f65d1a7de9b59e9ac45ce3222f6574d964d80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
