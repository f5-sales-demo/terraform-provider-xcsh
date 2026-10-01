---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2f7efee74d314a7e7e0c88afa9294aabacf1045b20e35d0c827af0217c678cad"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present / c58df7fe34f4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-81809ff78d5325bbe571dc7339a6d0da75cd1969352e735a7c8e5989a2ed5c71"></a>

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

<a id="canonical-d9bc56f09e01b83a173bc2e811010761c90a49ea47e786f0d491240d2b386127"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present / c58df7fe34f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0628e7433146d75aaee754ea70548a1eafc40b1586d56ef4ace5d549c1c95d00"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present / c58df7fe34f4 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b9f4c1587f31a2d4bbca087510f70ae93e27d43ff67f16f0221fab1a6c059690"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69c753d6725778be08ee700a92c94974ba0ce4f0ff94c07ec6f96c01e4e946be"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item / 695c1567017d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-360cd921a248aa2cb3a425e3b702072964c573a9eab8c45bba9b77580abc35f0"></a>

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

<a id="canonical-3b5258414da6aa4e5aca312e77dcf27decce2e08772c41dd2065ed97ebbb985f"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item / 695c1567017d / 3

<a id="canonical-fb26b3df1c6139b561d1654ad2895b245b383dbfd4af30bd0bc8754ee2f02c69"></a>

<a id="canonical-c719aaca26b9f6565d8922b55037793aa9f6b35f2b87bc3f8fac82539a7242c0"></a>

## exact_values property — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item / 695c1567017d / 4

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

<a id="canonical-474d77d38474868f49da7aa039dc9583d2a1ccfc1e3b89459dc65b96f8852d83"></a>

<a id="canonical-f02377d124f6a9800d727c3eb1d854cf665c43438e20917781aaec203bc03185"></a>

## regex_values property — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item / 695c1567017d / 5

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

<a id="canonical-7acce3961fc845090f36035ef97228db1f2ac9a776df150aedff86496a0a12cc"></a>

<a id="canonical-d9b56867b17c9fba01a0dcfa63b4bce8d26f525848ecc92bf4aa8d581b85bf98"></a>

## transformers property — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item / 695c1567017d / 6

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

<a id="canonical-fc6c1dd2c3bbce2b5cb17d0b2dd5f6467f16a940f22af1db8f648fb3990ef0ff"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item / 695c1567017d / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-64f0d6fbf93f46e0fba4f63c325e65566d1bbd8abff050f92508495d6fca05b4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcffce00d3363176911c27d2a5ad90dae79b87e1a4828c00a6878073f31e741a"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params — api_rate_limit.api_endpoint_rules.request_matcher.query_params / ca3641fdbbb1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params

<a id="canonical-23c9c32e33afef7a6459f4f9a7c5aca7c8c53516d3dd097273eb5ed7608a48d2"></a>

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

<a id="canonical-5a34ceb2ae576be2b6d78acd99dd8f4459367a03fc966dc8f7a99d25d80045f2"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.query_params / ca3641fdbbb1 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-ef14d0f99c73d7d350454515d2beee5936fc90ba49a6995dc0f37f4580a89d09): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-503c85f0ef78a344638059978699fbc0483487de40fa31ebbababb57526f2ab9): complete subsection reference.

<a id="canonical-be67673fb490ac58f8af7563290ae4864aa45fbfaac8a975950f5c9ac70a51ea"></a>

<a id="canonical-ac0f126d5368e5e8b0af53df0ddd5101da486015d850d2a39533024cc24c6e06"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.request_matcher.query_params / ca3641fdbbb1 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-def466a16b6e19c0c4c4b4bc93e41d26631beed367850205c59e998332d40f6b): complete subsection reference.

<a id="canonical-ef191015d4748b7ae48f3d3bed04219e6db3497a47b342203ecf596c073f29fd"></a>

<a id="canonical-1d5da1cdfbe6a5bb619e593dbaaf9cc2313854896838fe24d09ede2e56df31ed"></a>

## key property — api_rate_limit.api_endpoint_rules.request_matcher.query_params / ca3641fdbbb1 / 5

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

<a id="canonical-e56cdae3f2085a2f9cfefa2738cff3796eaa55bad63a2ad8b5d46216ce62e0b1"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.query_params / ca3641fdbbb1 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-ef14d0f99c73d7d350454515d2beee5936fc90ba49a6995dc0f37f4580a89d09)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-503c85f0ef78a344638059978699fbc0483487de40fa31ebbababb57526f2ab9)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params.item](data-sources--http_loadbalancer--reference--group-007.md#canonical-def466a16b6e19c0c4c4b4bc93e41d26631beed367850205c59e998332d40f6b)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ef14d0f99c73d7d350454515d2beee5936fc90ba49a6995dc0f37f4580a89d09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7d84a53af38f0d7131f4a903c17732e71378057e3f31603d2704461c13b9bcf"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present — api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present / aed11e1f10b3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-06ac338c7b0418f73e86af3ca4e9f9657c7827b31093e15591aa354ba0fc3dab"></a>

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

<a id="canonical-5917963d8957f9752e0f567b615b7cc2b3d92fc102c7f6116786082e2f8edbed"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present / aed11e1f10b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-05de524ba10049e8e601d37f971f95cbf42291be79cdf174a14acd92b1c89fc1"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present / aed11e1f10b3 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-503c85f0ef78a344638059978699fbc0483487de40fa31ebbababb57526f2ab9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09b2791b725cce22d25fadce88dc22d9eebe89894f5238b620846d298b5a8824"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present — api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present / e1ed28b10b94 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-a4167f377f81a2d4fa62ae00f9591eb0cfddcaf2776fdd476031cc9e75e7f3e9"></a>

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

<a id="canonical-9dfb456c93f47b13d5abadaa95c35b856e697136fc7560ed1061ecb9d14ab633"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present / e1ed28b10b94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-409dd03c5469383c665dc6e60588f4d4b770f7441598ee7b8a4affd29eb6eed8"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present / e1ed28b10b94 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-def466a16b6e19c0c4c4b4bc93e41d26631beed367850205c59e998332d40f6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e6dbe106e3a9cf715a5cd16d837847488bd9bc6123fdede73d2da0475af1fbf"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params.item — api_rate_limit.api_endpoint_rules.request_matcher.query_params.item / 38bb01884407 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-f1e3f2cf03bcdb75bc9d9180eee3dd714af4b2d43ad2686b1dc793258a8c88a2)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2d4cd407344e9f292d286ed9674a89622c1915dc900d67c72747425dfb2f9670)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-5579e40808ccbede6eb5ff774b8bb98fe3e54fba5052b47406349ecab830f013"></a>

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

<a id="canonical-a7048a30524a977c82afe739d69c6c4a4a5ed21e285d27d774298e47f9fdebfb"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.query_params.item / 38bb01884407 / 3

<a id="canonical-a0297172d9824523427a1c629c603de97050c28f65dcdfd59abe35ad34865875"></a>

<a id="canonical-ae2af2179ecc69967dbb7bc020c1c3ebc4772a2db1fd4178b92813b4f13261bc"></a>

## exact_values property — api_rate_limit.api_endpoint_rules.request_matcher.query_params.item / 38bb01884407 / 4

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

<a id="canonical-7f23d7fcf21f1fde0c25224822db41b1f36e3ede2fac0752679df96bf0385fa9"></a>

<a id="canonical-487cb8d113eb4d6fe99b7321f086b788af23f77c74c5a2da3508df104198579e"></a>

## regex_values property — api_rate_limit.api_endpoint_rules.request_matcher.query_params.item / 38bb01884407 / 5

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

<a id="canonical-526e407a68f664e811d635b0389f23a1e78ed327c6ba40d5e61fe42e373cf42c"></a>

<a id="canonical-5cf7d9ebc00a53a0252cec2561e23245c3edae7f92daf27ee5cc2d976c6a5c87"></a>

## transformers property — api_rate_limit.api_endpoint_rules.request_matcher.query_params.item / 38bb01884407 / 6

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

<a id="canonical-d84a805165b62706b76b1ef4ea336c4568a4287fcdc25ca8b05f7023265535b1"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.query_params.item / 38bb01884407 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-d22b5abe1590e1c5c111e681b584f8aa2092393654821def12ae09a1dc2d62d2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-965cfc117289f2adb714f268ce6f4ae9e48f3fa512a1f626382fe2cae9cdbe8c"></a>

## api_rate_limit.bypass_rate_limiting_rules — api_rate_limit.bypass_rate_limiting_rules / eb82af86169c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- api_rate_limit.bypass_rate_limiting_rules

<a id="canonical-61ecc4146885004e3e36b119534fa8665e0c9c17e028309a800ab642f66c724d"></a>

Type: `"single"`. Computed.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8868092747100f52f61bb5857a23a095fc0f881dfcc881cfb8c249e9dc7f8d22"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules / eb82af86169c / 3

- [bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385): complete subsection reference.

<a id="canonical-8f81267fcb9b2bd308e1b3762998f3ed18c9897f0e0996dbe005cd7f7e8fb893"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules / eb82af86169c / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da2f24a7c1729ede040dfad795ac4e50a825f7bbadfa006af501a568a2fa4e17"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / 6eb756d7c17e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

<a id="canonical-f5eab770a193ce00d03d644a718089e03d87bd4a5ec96d03f318319ad5018128"></a>

Type: `"list"`. Computed.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

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
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-a1b644cddd25fc9ec88d7039735df03f29744ad4ec1fe9ce4701eb55709345e3"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / 6eb756d7c17e / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-007.md#canonical-35d157b62074dc6fb49262829f2b0888f72f28b1ff4fe17b678b2f853f56aa8e): complete subsection reference.

- [any_url](data-sources--http_loadbalancer--reference--group-007.md#canonical-50f88f9f603da695bd6502b0f61c1f3f76e307d4df0180959ad7fc37b1857e07): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-007.md#canonical-be702ca2a1b64ced536852f817bd3fdaa0f454592248f4649d6aab0c7bf88fbe): complete subsection reference.

- [api_groups](data-sources--http_loadbalancer--reference--group-007.md#canonical-c6d06bcc7f254387f0bbebf7463edefffb3c72d4f8e2209fab38135cc0e25c08): complete subsection reference.

<a id="canonical-06cc6dd9c90cb7d6ef4006e57c1eed82f2b0bc63ad529727cc69fd8e466d095f"></a>

<a id="canonical-8ced12bb20310bf36580285a502bd94b629e5deb5db06f3352046c90be78da78"></a>

## base_path property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / 6eb756d7c17e / 4

Type: `"string"`. Computed.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Upstream description:

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a): complete subsection reference.

<a id="canonical-08fbf534d5c7877622e55ba64eba293cebcc6d0d36f6ec0911dfc61f68f5279f"></a>

<a id="canonical-fcb28673e0656901a99cfded8f7c98f0823a049777774b663c9875a803890a21"></a>

## specific_domain property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / 6eb756d7c17e / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

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

<a id="canonical-3ded693b317810995a7b54f72c8cfdf9f08791477c72600a555624b4a707d430"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / 6eb756d7c17e / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain](data-sources--http_loadbalancer--reference--group-007.md#canonical-35d157b62074dc6fb49262829f2b0888f72f28b1ff4fe17b678b2f853f56aa8e)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url](data-sources--http_loadbalancer--reference--group-007.md#canonical-50f88f9f603da695bd6502b0f61c1f3f76e307d4df0180959ad7fc37b1857e07)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint](data-sources--http_loadbalancer--reference--group-007.md#canonical-be702ca2a1b64ced536852f817bd3fdaa0f454592248f4649d6aab0c7bf88fbe)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups](data-sources--http_loadbalancer--reference--group-007.md#canonical-c6d06bcc7f254387f0bbebf7463edefffb3c72d4f8e2209fab38135cc0e25c08)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-35d157b62074dc6fb49262829f2b0888f72f28b1ff4fe17b678b2f853f56aa8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c4faa293d5a1b70d5822b4c39815d14609bd53ba8b1f4ef81c12001c1fd4867"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain / 090529d50638 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-906c2e482a2e57ffbaf5358a66a12bb5e47325db825a6058458d66f91e6ab675"></a>

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

<a id="canonical-0e2e1ec9212ca7da53d3cb932d8b191e23236b68e77826b2402f35b19b3e3e3f"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain / 090529d50638 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e501ad0de79e62f900dac92a5eeeeac08500508e1027c78cd12f9b062f3d9f3b"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain / 090529d50638 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-50f88f9f603da695bd6502b0f61c1f3f76e307d4df0180959ad7fc37b1857e07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3683986df50b62963cbad2b547f8dcf98e79391fcf0271734c9d6ae4d7e0294b"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url / ae75dd62ecd2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-addc696c104b18dcc9fb1d230ec868a2bd44e5c0aefb5ea9a794c8c8bfef185c"></a>

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

<a id="canonical-a9a00a7e47175a6965060fee1690725ba83027a94edb6e96647aea03c02b3c0b"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url / ae75dd62ecd2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d22c1ae74a9b4966e715c3d57219c1a194bb91ed88c1034b5ef01b119b02482d"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url / ae75dd62ecd2 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-be702ca2a1b64ced536852f817bd3fdaa0f454592248f4649d6aab0c7bf88fbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f572a06f04a7e739a1712621b810cb4c761482d0fbce84a9b8bf2e76716d17f6"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 78d5ec3c85d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-087d2a0c0629ad12a3f9747521e7e5c404adb128e08740ce8a8b6c46351282d7"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b424fb7be849b0a79cb66d3d7d962e995cba2910a4cdd8151865351523758140"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 78d5ec3c85d9 / 3

<a id="canonical-0d5d83e260a7395055d580ad99ba538e21c0ea08722f755cf84f4e9228e1c518"></a>

<a id="canonical-f322132f8ffc8f34f06d41149962fefd48ef2e6fdb231cea94a7b22ae145670b"></a>

## methods property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 78d5ec3c85d9 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-92e8caef8647ab2df80eb29bbaf9f09e1f4d59970f92630466dab8c20c906750"></a>

<a id="canonical-57731c0428a3ed803da556bd362efe14909c22b3db2f3dc5a5735399cbe59453"></a>

## path property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 78d5ec3c85d9 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-97ba8a74c13c991928427e9eecc2e52a007d7a315a81c94e1820b8e06881fa75"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 78d5ec3c85d9 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c6d06bcc7f254387f0bbebf7463edefffb3c72d4f8e2209fab38135cc0e25c08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f83ef15b864bd78172beca9b5031e16e6c9950755e02b9ea68a09a6708d1f52"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / b87220e05abb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-bc7e10d87cf3ec460782271feda266ba1ad5bf842ccb921d8a60b219a52d0dcb"></a>

Type: `"single"`. Computed.

API Groups.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f6b96d085cb4d5a2271a59d95d891a7eb04e2b3c1831503f86617ba0352dd717"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / b87220e05abb / 3

<a id="canonical-3ec400b945cbb96464a04ab7bea95bc0be51d250ec0fa3ffdd9d1662dad35c49"></a>

<a id="canonical-e81e751b465388d07a822f37f3243284622a9139022547594cab4888afe30218"></a>

## api_groups property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / b87220e05abb / 4

Type: `["list", "string"]`. Computed.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f9572f5b3466dfe5258aff2262f8c3158ac5dd90ff1f0242f78c86acb80192f2"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / b87220e05abb / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a20990c42d4ad492ffddd4a482fc6defa84a9d16105b450d8fd957e72fff7f92"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80d377a66617 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-e5abc59e34fee7bbb7f238a58f250a081783bf95d5144518f13e09c8462993af"></a>

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

<a id="canonical-0060bb0e5fa1ed264eeb7409b14e9ea53157410eb993b817c13247c1778c9c0e"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80d377a66617 / 3

- [any_client](data-sources--http_loadbalancer--reference--group-007.md#canonical-2d5b41e366591540c9d4fa72194cdea672deae3b6bfaa45610c3c4afaae54c52): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-007.md#canonical-acbcc2c2d320a75e49677b9c71f808fe8568192670608857562261949916b0c1): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-2cfa4d3246f3934cd228a208e5d34453522422e84e21676005263d643fe2c798): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-662664f5ce0f05f79d8e8644a3ed6fa666addc2b2898b26f3272d702b449a7aa): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-007.md#canonical-442b85a9110dedd8cb1e1fdf0f2a10bc62c69f2328b28428f8b7276d80d4855a): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-46d4cd002a848b746a6b13fab0b6d5aee5e470414a1f98ef9de7afcc9e86cb66): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-d27aa9b95605c6b1d807a37b210674e207d6a9b76f38a325e952f35a469940ab): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-3973ec55c243012832744d4c601e3459428a82c47cdad7a36c516e4c45f5826e): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-87ed5af3abbbaf534a521b568c195a6cf6e3eeecc4749db562fb6438e53647e1): complete subsection reference.

<a id="canonical-29c823d19103a6b7bf8662a3c278a5174d766e4e97fd02dae60d31c2f2d624cd"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80d377a66617 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client](data-sources--http_loadbalancer--reference--group-007.md#canonical-2d5b41e366591540c9d4fa72194cdea672deae3b6bfaa45610c3c4afaae54c52)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip](data-sources--http_loadbalancer--reference--group-007.md#canonical-acbcc2c2d320a75e49677b9c71f808fe8568192670608857562261949916b0c1)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-2cfa4d3246f3934cd228a208e5d34453522422e84e21676005263d643fe2c798)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-662664f5ce0f05f79d8e8644a3ed6fa666addc2b2898b26f3272d702b449a7aa)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector](data-sources--http_loadbalancer--reference--group-007.md#canonical-442b85a9110dedd8cb1e1fdf0f2a10bc62c69f2328b28428f8b7276d80d4855a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-46d4cd002a848b746a6b13fab0b6d5aee5e470414a1f98ef9de7afcc9e86cb66)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-d27aa9b95605c6b1d807a37b210674e207d6a9b76f38a325e952f35a469940ab)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-3973ec55c243012832744d4c601e3459428a82c47cdad7a36c516e4c45f5826e)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-87ed5af3abbbaf534a521b568c195a6cf6e3eeecc4749db562fb6438e53647e1)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2d5b41e366591540c9d4fa72194cdea672deae3b6bfaa45610c3c4afaae54c52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae7220fe596370adc9559327e81773cb230f5667283a698dba5e1b77cc7d213f"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 5810c051d873 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-2f20957f6442844f2bbf44224698c9a52800c3851b9d323274de7fd033207a2d"></a>

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

<a id="canonical-0d7a1d3c8c3f07b1b60579ea8f83a5eafd8ad5f154a14f5b4f6aba3d67ff4bef"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 5810c051d873 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ae5f3f377d548016f33fd03fafba0f664267d4dc858eca7fbb2cd768e81c7ee"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 5810c051d873 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-acbcc2c2d320a75e49677b9c71f808fe8568192670608857562261949916b0c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92e953a2ed63683a06319a528a688c2752ac5cd364ff837f6a40d08a9dd0891b"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 2ab581a94ab9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-8748ebd29459ab7e4922d72c8decfe2dc7a225201cda275119be0e67e8618b00"></a>

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

<a id="canonical-80d561a170a310b3c56e88147c871b9d4cde10cd680f72770a032a54a368bb8e"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 2ab581a94ab9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd18459c0218825d89dffd64a8819e6da0c3aba8ad376320bfe0549b55118502"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 2ab581a94ab9 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2cfa4d3246f3934cd228a208e5d34453522422e84e21676005263d643fe2c798"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2a72e4743483f806304d9edb4db05394f45100c16b3b5024d1fdf02d1d2f1cb"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 29eab9dd7945 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-462a056f1ca974cba21ce0793730af185d055685281bf6827b9597601afa8422"></a>

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

<a id="canonical-7a585b2de9e55e169591423a3111549220c8a9c2063d43820daa7d47730ad587"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 29eab9dd7945 / 3

<a id="canonical-19060c0f0471ed9a97e222bd825e6b328a134e8a00a9d9f5221ca0b526511421"></a>

<a id="canonical-9d5b3dac0a192d75017b9da9dfb91f72ffeff3a2f9507aae58753bc770f5d4d9"></a>

## as_numbers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 29eab9dd7945 / 4

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

<a id="canonical-48d67f41860481fe4fc6f9d63d0b99a9934d1f1d390e3c17b084f941a70d49fe"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 29eab9dd7945 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-662664f5ce0f05f79d8e8644a3ed6fa666addc2b2898b26f3272d702b449a7aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e36dc60288b4af7407b01942b3518dc489da6fad0e6f3d507810d00ef7f3cb72"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3a8fc909024e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-b61688d4faab0e947e63864922d4e5441a5a5d14b0b3ce02558c13cf5f80b5ce"></a>

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

<a id="canonical-8b89c3a9816fe0ea22c68f042651397ccdcf1f2596617fe5d076498842eaafbe"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3a8fc909024e / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-007.md#canonical-1c3d37ead305ecdbca0e9552940fb4ee9534090e2c05cc32820aabf3086a810b): complete subsection reference.

<a id="canonical-7892635d8e203e38c86bc270095b2e4af11462105cf2ad4d6257354c7aec3f1c"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3a8fc909024e / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-007.md#canonical-1c3d37ead305ecdbca0e9552940fb4ee9534090e2c05cc32820aabf3086a810b)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1c3d37ead305ecdbca0e9552940fb4ee9534090e2c05cc32820aabf3086a810b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74f445c9c3ca312f7a2fdf39c24bf1eaf810a91618f7c7f7ad7e9738dc807978"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-662664f5ce0f05f79d8e8644a3ed6fa666addc2b2898b26f3272d702b449a7aa)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-e4c407ae12b9a340af2c13b25824092681781990c27b532e5f0c88c77821a8df"></a>

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

<a id="canonical-6b2750f53c331835879f6036ac3191a3e889b338aa47aebd1f26d8e87d55c7b3"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 3

<a id="canonical-ae6f4104ce12cfd5c297e3961cc9aa1939788b094f678708157cd818ed2ed745"></a>

<a id="canonical-d8899b0d07860416707de41a20e34aa1a6dfe51a39df4bb4ec005eca21f96841"></a>

## kind property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 4

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

<a id="canonical-5dd11cbff113672c12a10d5f06766b3b54bced4bbf3ed03dba3d252a51afe512"></a>

<a id="canonical-24410effa56adc128a23436ec1ebc648cf5efc1da831741dd86897305d1fdc60"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 5

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

<a id="canonical-4ce9a1a1eb5fb1eeceb68895714da6a020c1a73138d443a0559a168cac63adaa"></a>

<a id="canonical-c6a6a69233200d7faabc7fa5dcec5176a8709b208d01b43eccf8815dbc96659d"></a>

## namespace property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 6

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

<a id="canonical-d347eaffd9779e0702727cebb2970977508dc2220ad8ca9da7ed993a657682ba"></a>

<a id="canonical-9d74366f0fcff59c045271e3296385c6da3ead36c1343bebe1a80ff7815bd6f3"></a>

## tenant property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 7

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

<a id="canonical-bc4fad8826bf5f87b4b54927b46c97c8356bc3e3a3b8e4bf896b6c6990624fd6"></a>

<a id="canonical-8292d1a386043c0cdd14fa9eabb2a5f7c203095402d2968bc95e36a257332a9d"></a>

## uid property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 8

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

<a id="canonical-fdc947b0c39630b1f93e8ae52bc85366c16d7311e750b3afebb7b4993c487b8d"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 041c9eb8296f / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-662664f5ce0f05f79d8e8644a3ed6fa666addc2b2898b26f3272d702b449a7aa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-442b85a9110dedd8cb1e1fdf0f2a10bc62c69f2328b28428f8b7276d80d4855a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ac8c2db284d286669b7d65091ce6b40384656c753e89b2cb72cebf094cfa83d"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 568e28f95455 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-7f2389320d5336e5cd40657bc854c25515f1a0556408d85f876a905666109bbb"></a>

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

<a id="canonical-5e465652c57daecaf99f1a5456fcdc322facb10295f830f3d7aaeae62df16e54"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 568e28f95455 / 3

<a id="canonical-4b634f6d06e1062e34a47d942e153a48e5c6f4f0da36e48ce1be687063fd077f"></a>

<a id="canonical-1c4a3c574fa957d9255e75ae31207a6baea11b94e32ceca0b0154b93193c999e"></a>

## expressions property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 568e28f95455 / 4

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

<a id="canonical-92a1476021c10f29a5de584929ef9579d321a29a8421a7495b8bbbcf0ddd4178"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 568e28f95455 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-46d4cd002a848b746a6b13fab0b6d5aee5e470414a1f98ef9de7afcc9e86cb66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f50b2e84f009579a20d50fea4a98b48a710a3fe288792b3d6aa3a70db697d20a"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 9841c6fe2d17 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-c9f7e073634ed6d4c7f67a13ad387a6005f8d06c41b975bac65f9dd217ece023"></a>

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

<a id="canonical-8c8a0a799731f34d435c866c93267160ab075f84b3827eaa9723eb9e4a20be62"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 9841c6fe2d17 / 3

<a id="canonical-ea355c69d475930a705e663c2438d5b5f49c53bdcc2d64fa2f87b1884399c193"></a>

<a id="canonical-34d2e22d6206329e18c6bd3115f3dd56b02cb9a246e14e09e118014dffda4eca"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 9841c6fe2d17 / 4

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-007.md#canonical-3a569a91a17826b2be2b1e942c2f2f09f67cea39a00c40cd2bf067919b0424cd): complete subsection reference.

<a id="canonical-1c0e31dc560a6e7b4ec7cdd4a47f99fd1846930b806fdd6d31f4fbc7514573d7"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 9841c6fe2d17 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-007.md#canonical-3a569a91a17826b2be2b1e942c2f2f09f67cea39a00c40cd2bf067919b0424cd)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3a569a91a17826b2be2b1e942c2f2f09f67cea39a00c40cd2bf067919b0424cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de720261dc226974c7812996e6d168e2a8fe64fa9e96481de948184843a570b3"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-46d4cd002a848b746a6b13fab0b6d5aee5e470414a1f98ef9de7afcc9e86cb66)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-fcc018108c21d603a3bb615094a4739595b71cbdf951dc3ddb69def7e43194bf"></a>

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

<a id="canonical-fa3cacf4e9c61206dd9b088a8e5de2adaad9cbf05e90e067534c41c618d03a4f"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 3

<a id="canonical-df502731f6e7f4ffa7c5254bbd986600c244aa399eb201ddad5c19e8bb3c641c"></a>

<a id="canonical-1eb7d73ebefe15d3274c3a42c12f975e3124c484702e02915f3e9f5c92fab557"></a>

## kind property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 4

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

<a id="canonical-836e16cdf78b6780378b2a3510007c2961d141c83c2fc0119dfbd3bfc33490ee"></a>

<a id="canonical-0408db19c2b2433c788b468e77bbfc34dfe30f481e8ab60e859ed48cce5f9f07"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 5

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

<a id="canonical-695c266a123448538116f6e06d44ea942787142f8e5e69aacba69ff1ed669c4b"></a>

<a id="canonical-5b8c80a243f061977c55ab230f188d7908e88a551ef0e98480894b79f248bd82"></a>

## namespace property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 6

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

<a id="canonical-b3b280bb055c28366f1a74657eac71abad593f5fd574268658e263ff51022074"></a>

<a id="canonical-f183971f0eca05acf4cb09eb9f6628654bd28930c71b123998228dd3f67de908"></a>

## tenant property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 7

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

<a id="canonical-c669c3d61d135f468362c5fd504b67c3bd80f8bb0385f5319b3ed9be775e7f3f"></a>

<a id="canonical-0e7b0eda70b06a62e3935baea2420c84e77cd34076aa1e2245ebbfad65bda299"></a>

## uid property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 8

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

<a id="canonical-0e4f4419acbca840867489ce347f10fcb5f61d7293876791c1f852edee942fdd"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 685b9240b346 / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-46d4cd002a848b746a6b13fab0b6d5aee5e470414a1f98ef9de7afcc9e86cb66)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d27aa9b95605c6b1d807a37b210674e207d6a9b76f38a325e952f35a469940ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd380cf0b9c076bd18f61e424795e54a3175e9ae51b9f2538147ab9858e6373e"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ddb0904e7a9e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-dcc85c861fa9adaf3a5f919dbcd2583ce23dfa3fbb5976a01baa8fdb428e7aae"></a>

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

<a id="canonical-f3cc6fa6cf3445d1cfeab12b11fea2aee68440c15156a9e6a55bed571894c36b"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ddb0904e7a9e / 3

<a id="canonical-2b10579dd610fff1e20d8ceab5cd7c94d52bcd940c934749465208f28481c031"></a>

<a id="canonical-a61992f9586201271b093e0adc63ee51696135e9b7483cb7f1d4ea35385bbeec"></a>

## invert_match property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ddb0904e7a9e / 4

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

<a id="canonical-1a6747b33e3be732af0afea4966aa500e4cc6f1e792872e5a6d11f05705ae3de"></a>

<a id="canonical-770ba49cf9de5cd3eeeec9a994cfc11cf393256d3b44f6fe7093e50380384565"></a>

## ip_prefixes property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ddb0904e7a9e / 5

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

<a id="canonical-278023d8043809acdd13d90769cb438eba7a18ee771ef0721e569ccf30b92c31"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ddb0904e7a9e / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3973ec55c243012832744d4c601e3459428a82c47cdad7a36c516e4c45f5826e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bee0aad4a62c286c1744908f93a2fbe34b846321469a74f7e453a4a4626526d3"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 4778236da167 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-f58ed6e34cf7c0485b94b54353ece6b5636bf3c1522e05a5fb338ca3dc62c373"></a>

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

<a id="canonical-fb7871c4490323e374f7bf00e0bf81f5a81cf6ae1264996bd3b1a83099baeb93"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 4778236da167 / 3

<a id="canonical-1e4768b08ec458fead07426e44401d5fe13ede3f000b5e9dd62fd676200d0b6a"></a>

<a id="canonical-8bf5ce74e8e04ea283c291f70edf0fa111d184656f993f2e750d040f4f6ae5bf"></a>

## ip_threat_categories property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 4778236da167 / 4

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

<a id="canonical-c012af8c9bcd72419d71f9be064e27bf45c2a8ec0042d216abf3658be77fe74b"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 4778236da167 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-87ed5af3abbbaf534a521b568c195a6cf6e3eeecc4749db562fb6438e53647e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac1062643872d8e20162a23b627344b29ea72618d1399bc79ebce76a4a0ec940"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 93eee5002de8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-86cb5de6b3f66daad1c6b2fbbf663e4c663fbac6f918df9f0e995bad48174bfd"></a>

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

<a id="canonical-92042a8179c7064bb86c24214a8924a03fffcb4af5af73c15a78683385c1a1fb"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 93eee5002de8 / 3

<a id="canonical-c6613082e116427d983a02a2da794a4d20bd275ee33e55c0d991e9a94a4db2ef"></a>

<a id="canonical-d68942ddc4bfa78bfef5c3c9934da4ebdfc0cfc27276fa5282634f6a74db1ef9"></a>

## classes property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 93eee5002de8 / 4

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

<a id="canonical-9157d14708ce5ed2264940627244784ef31b0010b3cbfcb726598bcb8e9c1a68"></a>

<a id="canonical-e648706045e09acf72095a2641d616180d59833ad5765dcf7293f5a45d53279d"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 93eee5002de8 / 5

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

<a id="canonical-0bb7f9679d9150a976fc1b2a1eb8710b8e9eea65030ef093da624a446a0e3563"></a>

<a id="canonical-5a8cb1d0cd56e2f17ccd6ea254573399852731a3ee0e392ba5f714f4fe199896"></a>

## excluded_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 93eee5002de8 / 6

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

<a id="canonical-eda42f9687eb4072a91dd0401510a4d580b4f3d9f79c692131bcaf05718ab17e"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 93eee5002de8 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3f6c127a0d2417a14cabd818ac5ae01fe28254f408f9fb4417f60c403f599899)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ee2652adb23096f2d57e40b33a177b5343bce460842b8056eea9e6bc7e88ea6"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 3a127bf189e9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-37e1ba3f97f3318182c8a75b5d415ce99d3de676ea7c34fee7ec40f21730e87e"></a>

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

<a id="canonical-a4ae3e6132971ed3cc902fc6f8f65fccf4f217d841a693fa684a8e4a6be03939"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 3a127bf189e9 / 3

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674): complete subsection reference.

<a id="canonical-c62a8405ba6bb5411ea533a99aa6846b08fd4b849c887adf991c4411f2bf8ac5"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 3a127bf189e9 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6365be376463f8da1caffc56ae4ce09ee7cd5d62ce223a39c5841bec42d7ecdc"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 5de845c357e5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-50c01c7738eec2b3c325ebb526941bff90a2b6aa1f8d349bb43eec48505f524c"></a>

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

<a id="canonical-6581b53bc69e96b93f2a7f4ad997b7c45cfcfc638e5ed207ed079e3b7edb252e"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 5de845c357e5 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-d0e9f0517e35365375abca2568fcb7c0d52f7539d9dcacb40ec94a3841cf7341): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-04723234c5e0707ec7a5bf772bf7335365b2f040ff4183e0572e77a59933f15b): complete subsection reference.

<a id="canonical-c7d7e5731ae35a8094af5fe4aabd53d23b3c0838d1397fb7a34b5c1c5d50a539"></a>

<a id="canonical-a89b52f70dec54a3fa2eab8267a7f528353ca3967f68b09d913ba4fcdb9dc8ae"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 5de845c357e5 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-94dd1ad664140144b5f58c7c9e55f90c3256af63ca3735590ac001ff1257c5e0): complete subsection reference.

<a id="canonical-7456f60abcb5facc33de865779973475e7eddb236d57c8966355f53ca1556451"></a>

<a id="canonical-fe335a6084c13acd97281e15fa9a5ab9f3b1bf3d0cdd27eaa8f809fac874fb27"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 5de845c357e5 / 5

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

<a id="canonical-3ab26cfe19fa37c0882c0d279b2ffe551f86e3b8e14756ef4408f01a07aa9930"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 5de845c357e5 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-d0e9f0517e35365375abca2568fcb7c0d52f7539d9dcacb40ec94a3841cf7341)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-04723234c5e0707ec7a5bf772bf7335365b2f040ff4183e0572e77a59933f15b)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-007.md#canonical-94dd1ad664140144b5f58c7c9e55f90c3256af63ca3735590ac001ff1257c5e0)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d0e9f0517e35365375abca2568fcb7c0d52f7539d9dcacb40ec94a3841cf7341"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e75ce34e38cbff16670ac8428ce0b917ecf04dbcdbfaa176d149dc498bd4ea89"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1d87a15ed1d6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-cea0f03d956b18e418870a6d80eb02acd0c82ba106d56db441689d94ce4c1f14"></a>

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

<a id="canonical-4618f000e502b96d33459f30eafac7e3e4f653693b555cd24f9f719e2a760d30"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1d87a15ed1d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1819c003ae5e7d51c816a134abd6313b731cd12a11d0ba6de23bd67d79e06632"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1d87a15ed1d6 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-04723234c5e0707ec7a5bf772bf7335365b2f040ff4183e0572e77a59933f15b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb9a7f7980bf6517553398bd533716f1bdea3a6c7043e00aae16b5eba1bdcfd6"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f1cd9ec66e06 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-b591481b676c51c8b7c2a0c4ec324a706881b76fdcf17477e5ca0a47bfb81f52"></a>

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

<a id="canonical-d068d54de081b02b6cc3158c886433e66db43138152bbfb6fb88300a4698902d"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f1cd9ec66e06 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7bff7c47e585343bf68abaa724c0f6870f2164d109a4e640d3a38ac808f8aeaf"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f1cd9ec66e06 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-94dd1ad664140144b5f58c7c9e55f90c3256af63ca3735590ac001ff1257c5e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19112f2e92cf66249ec4cdefab06f4a3305bc1460c48d5eee155c459df37eb8f"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 86b515866cc7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-59456d0a70c99dca08402547f4cf54c82dc02d2e7885d7533541c027b0723338"></a>

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

<a id="canonical-d50e1cf54aee219a16654a7f9c5e9f38d82733b9f1d4a4746335c74d07113de7"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 86b515866cc7 / 3

<a id="canonical-d8b2ef3badccc840671a6d0ad8e726948df40602c80601895df70c66eabf4204"></a>

<a id="canonical-b702525833338736456b568ace2bf02cda249c82d0dc3a03244d34c6e0474254"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 86b515866cc7 / 4

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

<a id="canonical-29414c40334172e155bc1b4b18ce31a5ec76533cdc23baf2dbeacfcf7a7f65a3"></a>

<a id="canonical-e868dbc4a5fef0ff0c1a797aa42f0ba8633261b4376f842529fe71732138f4e5"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 86b515866cc7 / 5

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

<a id="canonical-f8b2964df8e004b9f4ead2a6d277804284d4eb2ba5a5e3e13b0855f08b968cdd"></a>

<a id="canonical-fe6d6291484187679cb1c9ba1c81cf61b3ec88a74a9600ef6bde467b03b974a8"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 86b515866cc7 / 6

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

<a id="canonical-da5241789c1e724861945bad6b0b0f2ce4406c5c272f566779bd2a3fd49c08bd"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 86b515866cc7 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-6fd5e18bcbb23763763a99d03a7b352a5fa710574ba9e9462454ec6504da26af)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fefe662a1d72582ab7d6c8fd102247dfb31d51c57aefda7646be6a438f9f64b"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 7a5447780d69 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-770b3fc234b8119009f6d295194020d29bdd898f31767e78eff8481eb4a478d4"></a>

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

<a id="canonical-84862f38ffe2ecd150e415c5f33235e41a64e196b3b520e4c8ec817a873cd0ba"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 7a5447780d69 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-21b5a2ff17818012b896b6154c050020259267951c39054db673783ded898fee): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-3cf5da69d9bc6832e07895a757d246fd5a9b3e4a44a6215ee93925704e16a0bd): complete subsection reference.

<a id="canonical-8ed166f4005ec2e97a05d18f120ea94c311a70a6f4d7070c5d3daaff1257bef0"></a>

<a id="canonical-fd2def3ee9dabc7b9cfbee84468202929bbd97f859136069a16138bc45531bab"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 7a5447780d69 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba10afc03ad4180d3a8b0ca43655353176296c141f8f6384b340dfdbcf6e7d04): complete subsection reference.

<a id="canonical-86374783c5c6368a071f530f6ec16bd6a9b4ebfcf77f696f7ea9681bcbf2661f"></a>

<a id="canonical-9b1317032f98a95d12b94c982fbe9f8ce8ef4c016fd3c3fa74486e0d1385320a"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 7a5447780d69 / 5

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

<a id="canonical-4cd96b4b36418094531f4e606d768eafe97a9f26f9c1d3d033d344ad24b3bd72"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 7a5447780d69 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-21b5a2ff17818012b896b6154c050020259267951c39054db673783ded898fee)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-3cf5da69d9bc6832e07895a757d246fd5a9b3e4a44a6215ee93925704e16a0bd)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba10afc03ad4180d3a8b0ca43655353176296c141f8f6384b340dfdbcf6e7d04)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-21b5a2ff17818012b896b6154c050020259267951c39054db673783ded898fee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1eda4b9a4c5f4788c569a286d3b1511a7fbefc6d9f425897412cec9e999257e4"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fd1465e4bc87 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-115df22b1852c192e176e5fd85dd83c862066583c6335851971a3328676d12b9"></a>

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

<a id="canonical-172848b80724cd818d09bda1219110f09afb0a134cea83a43848324f4861950b"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fd1465e4bc87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3c9b7abb67cdb66fb9fd1240a217874ef082735c94b62f21afdced0b67ae04b"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fd1465e4bc87 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3cf5da69d9bc6832e07895a757d246fd5a9b3e4a44a6215ee93925704e16a0bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d313a6de200d8cad2d13c3ba0fb1a7ae096ba1cf7d839c089642cd7e6c93a88"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 9778289931b0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-ea48888113da86339941d48e79ed102c14c696f6e8d568f3fc33427227d64c85"></a>

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

<a id="canonical-4534ed5d630157e0a35a2e9789a17fb3a10467d7145e8836ac4728af7a6a78de"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 9778289931b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-91ac01d68bac9045cd33bf95631176769458a35124a6701a0f02d31c652e16c7"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 9778289931b0 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ba10afc03ad4180d3a8b0ca43655353176296c141f8f6384b340dfdbcf6e7d04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6f7d67b87857cc14a01fa0bb56ab08618385b5bd00b0384f30168011031228c"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fb8069a1d889 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-2ded1c520b4f0c86111c365641548a2f06b0ea796cb3b2e57ea43399f90734f8"></a>

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

<a id="canonical-37e39ddb099aade88d93a76c952eae7f1b9c1c6f1753bdf40ca7efa0bbcadfce"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fb8069a1d889 / 3

<a id="canonical-86adc52b43178ed50f62339561eb47642042d7ff2943159c423fd282df2cbbff"></a>

<a id="canonical-f026a7572b8c652fcaeedae717c88ed6461f230533be14a9677e1a6f78024266"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fb8069a1d889 / 4

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

<a id="canonical-474242ad8947a0eb782fca801d575ca8db3cf47769dbb4d418963499f614c149"></a>

<a id="canonical-b01449a59dd7d78be497ef19ad24f2eb6f4f51b470e973cd7aeec76e747edbd3"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fb8069a1d889 / 5

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

<a id="canonical-becf5a980f8b3fba23457eb1583b6360ae209759f17ad22774c3033772b548e2"></a>

<a id="canonical-88eddc984293d0c12223fa33786071b0a83df5bd32d354ddc294b5896d89c059"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fb8069a1d889 / 6

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

<a id="canonical-77f2158f6c33a49932c1e80913d5c54e1c6ecca0e76ac0fb3a70e237b6f14fd2"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / fb8069a1d889 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-007.md#canonical-2a364224a6397dbfaddb6f90a5c303e60ad74997b676b341d80d365fb8a0ad42)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-105cc84dde602ed686a21cd010e70a043ca68e0d1680c9b5e1b4c4d81e9a33f7"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 067ea2908dc0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-b45903da3e3061fee3e4865472067128ba1ddb4cf5dfd55f0503e700243e1704"></a>

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

<a id="canonical-5bf3a4e00de5f18e4a225ce8570f7a9af0f9d24dd2316f1dec3c5a4cabd41000"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 067ea2908dc0 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-f80140f612dab5a3c4897dcb2319c321177abd2afd2341a6f93737fc08195cbe): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-06c9247b397b6d0f0e0ab5e051afe2612f3cd14f6dcb95ae103d47140edf813d): complete subsection reference.

<a id="canonical-f0818917d35de67c19f05916e1244f09609a19053a6090bf6f4b1096c1f2d36d"></a>

<a id="canonical-e7a005fa5410d86129f4ba0059c6c7585c8ccf8b6fadc776c1dbf5e6288ba55c"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 067ea2908dc0 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-b4820d7820f6f638628fba7443036daf2b6b28189eb99ef5550ea8ddab5e23b9): complete subsection reference.

<a id="canonical-55a440a3c71f33b2c1a898eaa6ccbf178627085cf79c5bc55d0e29577df77f04"></a>

<a id="canonical-d6b9da138bbfec5aae113964dcbc9c59a5e27667d5454bfdb3d496c661c3b725"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 067ea2908dc0 / 5

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

<a id="canonical-18b0e11c149dbe96f58551683d87b16cae5bbf015ceb7cd5b60a6aaaf9376c0d"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 067ea2908dc0 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-f80140f612dab5a3c4897dcb2319c321177abd2afd2341a6f93737fc08195cbe)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-06c9247b397b6d0f0e0ab5e051afe2612f3cd14f6dcb95ae103d47140edf813d)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item](data-sources--http_loadbalancer--reference--group-007.md#canonical-b4820d7820f6f638628fba7443036daf2b6b28189eb99ef5550ea8ddab5e23b9)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f80140f612dab5a3c4897dcb2319c321177abd2afd2341a6f93737fc08195cbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a8569ad60136ab64af23daab55d2c81c4cfdd05edf196d7dff041b3be266c80"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / ade780ef6c30 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-b08702d658cc9365121bee578e3afda13fdf7cf8227944c31b89eb22246a13fc"></a>

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

<a id="canonical-acbe6017a824522f2ff6a39d5c3dcf99b6b00aa3df06324d922ad741bd8cc376"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / ade780ef6c30 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c78ecba1c7600ea8cfdc9295142f80e5c42b2c90ebe6201b3683949d8dfa5d94"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / ade780ef6c30 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-06c9247b397b6d0f0e0ab5e051afe2612f3cd14f6dcb95ae103d47140edf813d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf532cb77cc846ee117daf5f89e694724bb3e031b9ec5dae1e1b098424bc08b7"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / d7631440f617 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-e4d09957a80e315609e0d73d3b03ad103d12bd17935caacb123ce26a115e3c87"></a>

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

<a id="canonical-2b129d466fd285e9f0a3a7df88b3d4c9d10e6efa9f8973f2255ce57281028ffb"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / d7631440f617 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b71f702742390d8b51ebe6a38a5a4a8c75de8ee7e7a9bc55d4d8d3f0f4b319d6"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / d7631440f617 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b4820d7820f6f638628fba7443036daf2b6b28189eb99ef5550ea8ddab5e23b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c528fc9e0c286f645cac419526670101582606b11c60b2d51988f0d85028d97"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 97c7a58b6e33 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-f433b832c425e794984e38404d716543a20bc65b7c4653c8b7fd58076f2a5e3f"></a>

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

<a id="canonical-1f3956fd019cb809332c15e28380b463678c4cc93a70b197f8523524999578c7"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 97c7a58b6e33 / 3

<a id="canonical-88fae6044ebe9695357fbd529faa0d2c5e048d0ad1a2a81494d41771be1e46b0"></a>

<a id="canonical-6bd915ec5f19497127ce90dae673c2399fddfe3a26da103c2b88e89416771725"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 97c7a58b6e33 / 4

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

<a id="canonical-55e472169b1e0e8f4767f79f919b9e61104efc31abbf1938628e347837fc0f26"></a>

<a id="canonical-f927689090d16f56c4a8c5c459df80fc641dfc70c4b32136ff47af585bedf654"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 97c7a58b6e33 / 5

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

<a id="canonical-17f1bbab45234da01e9f0a2bb2ad13fda829eda26efd896b601171fa179ad961"></a>

<a id="canonical-7b4996d59cd6f3df10162945f621e1a3142453d1965f6c2809ffdcc7af2235ca"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 97c7a58b6e33 / 6

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

<a id="canonical-f19734950b2d9294ac412392723f9eca20b809433fa5fc595dd9bcdf17bdd5bc"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 97c7a58b6e33 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-ce6c672127cd32b48b168979e1768fa414ae728ebc62fdd7714d8c04d6d91248)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cee93a3a9dbb7f366e03229a4bcc3461025f7b3220f2df4ac743e73cbe70f237"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1382dd05d376 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-dc20be9da0379b4f2886fd346f6bd4f9c6970ee5362eb4b1793c04b5e6a678d4"></a>

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

<a id="canonical-db6889648eae3b9147c200afcb896332f676aab6526723c55b90ce55fbdb9ce4"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1382dd05d376 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-d79ecf1118cab75b4cb258ed2171d8402dbb5ae46eb810997b03b17d7e76b27d): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-5d57446f7f48aa2a0a80b5d156d5fae850662e2d88f6cef4e881930dd157bdca): complete subsection reference.

<a id="canonical-9ff83e2a92beaa0bb631d9fefe84fabce52099d8951ba352fb51ce1503a5bac3"></a>

<a id="canonical-516e7d3e24b63ddfe90d47d47c9eb4cd0329a532180a9dd21181911d7cdf6435"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1382dd05d376 / 4

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-7b9acf60e802cd7b8424f81e0fe2831fc98ce634e2a0b9d1767d0310d9c47c88): complete subsection reference.

<a id="canonical-d9d7a8d3f7ca1a6fc5cd4679eb1d8ef230783281ae1343114ea0a6592adedf11"></a>

<a id="canonical-5abb2f57fc35b587b12caf8acd8df907658f7dceab54999ce0b432215bdcf2ce"></a>

## key property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1382dd05d376 / 5

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

<a id="canonical-696edf4887a0cd97ca041190608fe6ae8b068b72d0f235bf1d1cacf9d95b5027"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1382dd05d376 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-d79ecf1118cab75b4cb258ed2171d8402dbb5ae46eb810997b03b17d7e76b27d)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-5d57446f7f48aa2a0a80b5d156d5fae850662e2d88f6cef4e881930dd157bdca)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item](data-sources--http_loadbalancer--reference--group-007.md#canonical-7b9acf60e802cd7b8424f81e0fe2831fc98ce634e2a0b9d1767d0310d9c47c88)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d79ecf1118cab75b4cb258ed2171d8402dbb5ae46eb810997b03b17d7e76b27d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-396d0d6a2b1ce98a4bf96eefe6c3a7f9f716239264d5d6242223d1b1b593c450"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 424b7a9b4b27 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-db642d4f710c20c7eb4a5a03c27f47634a84ed982ce9d5d364fcd2cf4fb9355e"></a>

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

<a id="canonical-dcef0886426ad265fee6c9012286f6630fd6ad94cfa419359a241c336d95a4f2"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 424b7a9b4b27 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bff7999e94019ebd606e10e2cce83e90c34931eb28b8c18acb3334b994c0b8c"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 424b7a9b4b27 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5d57446f7f48aa2a0a80b5d156d5fae850662e2d88f6cef4e881930dd157bdca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3707140bca38eb51c58aefed6056abba1c29c4dd7347cca1255ee671ef7ddd89"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bb437243f806 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-f0e587362fb3a2cfbac06cd4e45c75a522bd44f7a00fce8fbd1a96d9b4c7a60e"></a>

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

<a id="canonical-90cdd8208a5ee15391d4f46f54af801e59692925827ae89f0c3202c595461249"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bb437243f806 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fddcab6ba82cb90e74b4883bc6c4d3c658ef08d8509ec5c97df2911c110556d9"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bb437243f806 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7b9acf60e802cd7b8424f81e0fe2831fc98ce634e2a0b9d1767d0310d9c47c88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6544579f93261bc389240ccdb68ce14d8621576e98688e510b1378242b861ca"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bef79e6f925a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-e6b80ddaea5f54ef4f19e11526cf6a439e3f195aa855bca14a510a5d98d0e596)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-8314802432b04d348f907a61e1cd98e9e0961caeac63b09ef7f402055501a385)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-be59f4e0afab0865cc29098f2aeceda158162ec0b9ec97642cff4f74a714cc9a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-edbda80b9c78b769bc906add67be293b3f21e60efdc515e71e2fb028fd918200"></a>

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

<a id="canonical-3e33c7de0f21765c2cbf66b0332c74a1bcca5a40e5bf8da4dda546dee29d0418"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bef79e6f925a / 3

<a id="canonical-c9ad00a4db349185bcdf4ba4d4e133675e50b25e3098638cb888faa04f83f12b"></a>

<a id="canonical-7e41428de4d358560077caac8f8d23d28c2c9bb9dfb1cf9c8c2558a3cc9bed82"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bef79e6f925a / 4

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

<a id="canonical-e372d2266fc43875bae6bdadd3598959675f7279a584b35c8ae6532cd9501026"></a>

<a id="canonical-1c8647cddc4369fe716cc31cca4420c099934f0ca89f5fe85b2bf4450e3408fc"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bef79e6f925a / 5

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

<a id="canonical-aad591b8f006e0a67c5f9c7dccf2c77eaabbad07b6a2288e198ec3ebe8731021"></a>

<a id="canonical-39ada721928f15a09c199f0c3df6a3656fe86544964762aeed3b37b805876d3a"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bef79e6f925a / 6

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

<a id="canonical-5aac505f54f32a4b3f318f0a0f7a2ab633b232592dbdf0b2560bddc8f4718de8"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bef79e6f925a / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-ba5024c2d69a7600bd266bf1c10399a34e32d935d571663d5bc8d873e4008674)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c715e7544dcff217a1f6dac72a2bd057247ccf7107eb5c68b4f07cb612bc936f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35e1c714454acfc2383d28eeb53198d7622df7e62e139230ea19b500b72c31c7"></a>

## api_rate_limit.custom_ip_allowed_list — api_rate_limit.custom_ip_allowed_list / 6a7ce8aa999a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-c94ad9b5ca667665daec13974f29ee2b59b11301af82407f9037cf24ccf1a20f"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-91c8db833f319da3690419f3d505c3f9e9c6a516bf8a5a9255dc4e775aa6db38"></a>

## Direct properties — api_rate_limit.custom_ip_allowed_list / 6a7ce8aa999a / 3

- [rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--reference--group-007.md#canonical-0d95cf68214b6e1213c4e1532161a442bde2d70e28be24f889bdf89e15eaef95): complete subsection reference.

<a id="canonical-2ef27de8a2d5a0379222f756d725e6a271540b99e6c2a3b42331b856b85728a5"></a>

## Next pages — api_rate_limit.custom_ip_allowed_list / 6a7ce8aa999a / 4

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--reference--group-007.md#canonical-0d95cf68214b6e1213c4e1532161a442bde2d70e28be24f889bdf89e15eaef95)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0d95cf68214b6e1213c4e1532161a442bde2d70e28be24f889bdf89e15eaef95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-657832d7f1a0a9f37a6bc8f8f9a8a8d4eef72d27996b59182551434c9ad5ccff"></a>

## api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 2f7d58ccc173 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-82d6e8342baeeea67458651fac504c258e5ed9e3c2df4f193cbede248617ecf5)
- [api_rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-c715e7544dcff217a1f6dac72a2bd057247ccf7107eb5c68b4f07cb612bc936f)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-9b15959c9d49320287f79a60d468869b6e1c125e3a2a6c7a6c4c83525a5f592f"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-274b8d7676923c2117d08601a321a9663d602ee602e09c1ae4883e84c3cba8c9"></a>

## Direct properties — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 2f7d58ccc173 / 3

<a id="canonical-7e9b29d9da040962ea52435d210eb26b2349ce853e3fa67476b76519b71c0775"></a>

<a id="canonical-3857e2832679189ee21b030cda7c45c6ba8427a889158b1d95a667718463b735"></a>

## name property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 2f7d58ccc173 / 4

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

<a id="canonical-cc14c54e4a9aa117450477c0780d3f400779d7c9912d18709fa7bd84e2792c50"></a>

<a id="canonical-245913b550661b8b50ad837d7804324167391b6654b167ff29cbb9f561d74fde"></a>

## namespace property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 2f7d58ccc173 / 5

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

<a id="canonical-a829330ee52586f8c89498674d3100f031a55ef011d4d40ade8b03c6fe847340"></a>

<a id="canonical-0022565946d07ded41326ad75c8deb481f1a448d104f3b12af247db5e11419ad"></a>

## tenant property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 2f7d58ccc173 / 6

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

<a id="canonical-7ad45c76e43dab7070f9d581a2529a770f3c3dfc6b24864a4278958e8e5a32c4"></a>

## Next pages — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 2f7d58ccc173 / 7

- [api_rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-c715e7544dcff217a1f6dac72a2bd057247ccf7107eb5c68b4f07cb612bc936f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d57cb4bf5849c5b7d343d37b5411422264ecfd7929bb36f15c9a460c7e4057f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
