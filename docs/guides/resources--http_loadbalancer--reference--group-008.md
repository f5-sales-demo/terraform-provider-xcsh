---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-299350ba1cb0571a65315c42006879bdb5df9899040106b166aeb1bf62b40334"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 59fed3544b1f / 6

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

<a id="canonical-ccdf4a6e1a7f11123b59fed902c1db5ed92c86056a0ae13616e964136dd90722"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 59fed3544b1f / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-a51935220b8c60cb4adb5b2a62bb0fba8d5a90f68b89705ce3e15298b04d6e20)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-74f6cab9ed823ccc813bf81b0abfac42571822ba9604bd80ad10cc4b9944f505"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f411b62fcbf11f1e15f67acef58a35cb80c03890ad3634c4fcddb8f0b57b5918"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 94cb059628ed / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-83438b206ab5978b7c2a91280ed12652707aa076d66ece929570a47a3bb96981)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0b0c3a368200742fd86e17dfe267aaceb313ea2d549903d400f4855f2519558f)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-1603810536d3938f9034c28a6970085a4a2624a284912172a335541f8ff3866a)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-2c4322c7815cbc146e60c00581b6a7a22fff8bed8b02598a3f90f825699c2323"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-7eb7f98f81cb34fc8831d70940e895fd27adb1c133d5f5fb233c0585bc3d4c2a"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 94cb059628ed / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-9ff689845d200c26008a1abc923daae9c2f26d8614ecf061bb64a3bb99a493dd): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-4cd98e950fb394b8652df35db3eae029e4586f53a4289a1a323e6fca84ee971b): complete subsection reference.

<a id="canonical-e30fb13a4808f2cbb9ea25d7de40a8257995b080d15547d455a8d2b065d1fe04"></a>

<a id="canonical-19c079c7055ec35410c6566b08eee0b0386c30ae659d8f6d82f462d3a83c9277"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 94cb059628ed / 4

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-e2c31d8d08325aa5376706758d99de7852757850f6db17f3dc1d557a57ebf93f): complete subsection reference.

<a id="canonical-8e460a279af6a3145b4f66b67fac4e32e54acd82a64a6da2e53f1b8d9fa3000b"></a>

<a id="canonical-ad2bbef374468128c12f38cf78976e46009065198ab2b54a1d6623b8c2b31063"></a>

## key property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 94cb059628ed / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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

<a id="canonical-8a89369852c661b93a248686084ee71303debf20bb897c3a696a0614be538e9f"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 94cb059628ed / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-9ff689845d200c26008a1abc923daae9c2f26d8614ecf061bb64a3bb99a493dd)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-4cd98e950fb394b8652df35db3eae029e4586f53a4289a1a323e6fca84ee971b)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-008.md#canonical-e2c31d8d08325aa5376706758d99de7852757850f6db17f3dc1d557a57ebf93f)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-1603810536d3938f9034c28a6970085a4a2624a284912172a335541f8ff3866a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9ff689845d200c26008a1abc923daae9c2f26d8614ecf061bb64a3bb99a493dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-972eb89be246ed326412b53a0cefefbe2fbf6348076950f9c0de798f0dd7c2aa"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 63388e9efcd9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-83438b206ab5978b7c2a91280ed12652707aa076d66ece929570a47a3bb96981)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0b0c3a368200742fd86e17dfe267aaceb313ea2d549903d400f4855f2519558f)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-1603810536d3938f9034c28a6970085a4a2624a284912172a335541f8ff3866a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-74f6cab9ed823ccc813bf81b0abfac42571822ba9604bd80ad10cc4b9944f505)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-f254bc0f265a0692bd6e7e1d65224ff24a693dd6ec7c8c769c874bf68210f85a"></a>

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

<a id="canonical-fa815978b25fa6e87827c4eb58b6d4125c8fc9e835d089c9c717e1b1397a390c"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 63388e9efcd9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb78dcf87729079cde1d7f144c6535beb47b6e3943ab96e31ad74a1f19d2a73a"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 63388e9efcd9 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-74f6cab9ed823ccc813bf81b0abfac42571822ba9604bd80ad10cc4b9944f505)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4cd98e950fb394b8652df35db3eae029e4586f53a4289a1a323e6fca84ee971b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23163844f0c71e00ce2536b2c0c9d4225b2591001883fb4639978bd028a3bae8"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 3609db573e75 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-83438b206ab5978b7c2a91280ed12652707aa076d66ece929570a47a3bb96981)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0b0c3a368200742fd86e17dfe267aaceb313ea2d549903d400f4855f2519558f)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-1603810536d3938f9034c28a6970085a4a2624a284912172a335541f8ff3866a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-74f6cab9ed823ccc813bf81b0abfac42571822ba9604bd80ad10cc4b9944f505)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-d5207896b8c6edf060fff2fbca0247e1149af6a3845ff96994fe54ebaef18c38"></a>

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

<a id="canonical-03bd8704a158504367483b53e430e29f4dec255aaced7ab25acc02afdfad32a6"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 3609db573e75 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f8348bf7540bc953753a0dd5dccd03aa06ad071d286e0cff4661c7251fcef93"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 3609db573e75 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-74f6cab9ed823ccc813bf81b0abfac42571822ba9604bd80ad10cc4b9944f505)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e2c31d8d08325aa5376706758d99de7852757850f6db17f3dc1d557a57ebf93f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d937ae6288294b39fb6287bab00cc9c4f3f198e6efac210678469463defe5834"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f57a5be398fc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-83438b206ab5978b7c2a91280ed12652707aa076d66ece929570a47a3bb96981)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0b0c3a368200742fd86e17dfe267aaceb313ea2d549903d400f4855f2519558f)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-1603810536d3938f9034c28a6970085a4a2624a284912172a335541f8ff3866a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-74f6cab9ed823ccc813bf81b0abfac42571822ba9604bd80ad10cc4b9944f505)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-2aadf6c35c6ef208002d7386e0ddc3bf158575249a9857a459d4354474381815"></a>

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

<a id="canonical-94d7fd7b123f3eefe0fbb88bc4203fad47500ace519f32d34ef49b8f30806fcc"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f57a5be398fc / 3

<a id="canonical-2c670527de326af97314263e40e6ed890749a3186d35022b46a893e23398d193"></a>

<a id="canonical-3144eee4780c9ba517b94e7c3213fea1619c1928b90fddec80bab10612e9b7e1"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f57a5be398fc / 4

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

<a id="canonical-23c520f3aabf7f1499233433ca36d7b0afa9d459ffe2e3e43205ce7b47916886"></a>

<a id="canonical-738120c8bd8db41ec75f09af2c01cd44dc4aa10ef7af5924fcc7c33b322af89d"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f57a5be398fc / 5

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

<a id="canonical-db12b7ec706bb7c307b080830bbf6c138d49930d7868e11061765ff416f9061a"></a>

<a id="canonical-6a55c39d74b58be478b04200ec84949183e765a855b93bba909e70b476e08bcd"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f57a5be398fc / 6

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

<a id="canonical-92d62130630c1a6f747737b3cbccd9dad0528f1231682e1133137ae419bb73ef"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f57a5be398fc / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-74f6cab9ed823ccc813bf81b0abfac42571822ba9604bd80ad10cc4b9944f505)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b7b28053e6c70126c76df78f5c8d05a459dce130378187bcac769f9645feebfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0e10a28c0202fdb1ed7ad27cbd12ac2485a207e8b863869f2e95255a3a9f62d"></a>

## api_rate_limit.custom_ip_allowed_list — api_rate_limit.custom_ip_allowed_list / 65a3b39b239c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-2be669abfb3f1b06746601f52bccf0d314cf123c67a2f6c7a5f360e579294768"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3c4bbe19bd4388c9b6ced52fd2ef60b9e4d4166944da2c9c6609225f4c292438"></a>

## Direct properties — api_rate_limit.custom_ip_allowed_list / 65a3b39b239c / 3

- [rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-008.md#canonical-449445d93dd195cf8c34e2bc49e81c2eed133334a63faab5cbe856f8098a1ca0): complete subsection reference.

<a id="canonical-aca2eb666039095af03e51a1d06ce4219fb6a1920ccd221a92f73dd6a65b743e"></a>

## Next pages — api_rate_limit.custom_ip_allowed_list / 65a3b39b239c / 4

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-008.md#canonical-449445d93dd195cf8c34e2bc49e81c2eed133334a63faab5cbe856f8098a1ca0)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-449445d93dd195cf8c34e2bc49e81c2eed133334a63faab5cbe856f8098a1ca0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b34dff439f81064d02f002e09653dbec851a98b8dbd614896b44312a83f8ef8f"></a>

## api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 6e97b63dc70b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-b7b28053e6c70126c76df78f5c8d05a459dce130378187bcac769f9645feebfa)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-f65373249db057ad4488f7fe90928f824cce5a76dd9d31be7e6ab7e7d515d824"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-69197ab3e121ad1aff482e70b3cfb6294dd401c1d83cb3568d7901faa038ce0d"></a>

## Direct properties — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 6e97b63dc70b / 3

<a id="canonical-ba5c8bf247076c4497856c03e3a8db607a15b4f39dae20b64bfb7c07d927112a"></a>

<a id="canonical-da4df854e46818b0b980e5b98a97b6e07c90bf1d60c2769b2eaa849423f1a976"></a>

## name property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 6e97b63dc70b / 4

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

<a id="canonical-8409bfcc96a67aecf18b3e82e21a3b5131b23f045c51a7175f9e196e10e4444f"></a>

<a id="canonical-8d55ac2c7610c5bcd9aaabd8624226ba842f8fc5c7fe6c29e65252671741d0d1"></a>

## namespace property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 6e97b63dc70b / 5

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

<a id="canonical-83e7cafd4978ec992615bee131da459cab38a4af136387456f42b29c0976864b"></a>

<a id="canonical-f684192fb2349775f77ceb6908f0a45ac8e7783012090cc7c47954f412760a84"></a>

## tenant property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 6e97b63dc70b / 6

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

<a id="canonical-80313ed0e82b0d3b97fcba0a1c02710b00cdc51ad6185720aeb175eb4f9e855a"></a>

## Next pages — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 6e97b63dc70b / 7

- [api_rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-b7b28053e6c70126c76df78f5c8d05a459dce130378187bcac769f9645feebfa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-61932de3a5c2e6783be5d0bb3d6ba270ac39bdc0fc21e0b0cc57a72fdcd06ef0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dcbe1ebc85d6d2896945ffa8da48b2f60997747121654530973ead836151c32"></a>

## api_rate_limit.ip_allowed_list — api_rate_limit.ip_allowed_list / 30589caebe5f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- api_rate_limit.ip_allowed_list

<a id="canonical-c6ff6eac66229789d573c24da2146dc8170a4ef63e6f177feb43867a38aae2c1"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-870c17d2827dc7743045cb7879158d78dae2b3ba64cc8c1694c52ea7b7408f41"></a>

## Direct properties — api_rate_limit.ip_allowed_list / 30589caebe5f / 3

<a id="canonical-de70cd1b9a610f6d06a01bbc832a8e5584a9b5a7658fdff22edd2c676480cf27"></a>

<a id="canonical-27362b2a4da4d1263c8a2f693b27babf80a5ce2e9be5f17bf3a0a1a81f64abbc"></a>

## prefixes property — api_rate_limit.ip_allowed_list / 30589caebe5f / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f0d7d12096cb3ba41db9a07a232723e2a4de27b5d34c162e9882167cda8a02d2"></a>

## Next pages — api_rate_limit.ip_allowed_list / 30589caebe5f / 5

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4cc948a3cab3c32c9bbacd1a87339619d2654cd958da12277ae5daeb5cbcf3f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b24bd1bb291de442e83ea6c9b567817d5a0cf927567f8a5cd83ae157b07a497"></a>

## api_rate_limit.no_ip_allowed_list — api_rate_limit.no_ip_allowed_list / e8fb3cc6712b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-819914044df6feb8f0e7aec577a4da175532428132865c81fcd22f832a523fa5"></a>

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
no_ip_allowed_list = {}
```

<a id="canonical-c9c432520d53fd1720514922f86601e646b8f5a89d8ba25e2888b45583374aa9"></a>

## Direct properties — api_rate_limit.no_ip_allowed_list / e8fb3cc6712b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42d1bce6cd30fb2ac69e7227f88c4c49b0e99fd294904895c292b0c7205b86e3"></a>

## Next pages — api_rate_limit.no_ip_allowed_list / e8fb3cc6712b / 4

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d677a7d63550cba049c20584bcfda40fcddaa8a8ca30fadc9e17103368f482d3"></a>

## api_rate_limit.server_url_rules — api_rate_limit.server_url_rules / ad4f9b858ad7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- api_rate_limit.server_url_rules

<a id="canonical-4ae821967fb2ceb6dcabb6a36cda4aed351b13862ffdd9d68355b93294b0c190"></a>

Type: `"object"`. list nested block, Optional.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
```

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

Terraform syntax:

```terraform
server_url_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-92a71bfeda567fe271fef36f95c6bdafb02e8c2b164d0ddefdfdcbf3443632c4"></a>

## Direct properties — api_rate_limit.server_url_rules / ad4f9b858ad7 / 3

- [any_domain](resources--http_loadbalancer--reference--group-008.md#canonical-262de909ed1f42fcfe1e2669020d10ded96b2560e75149fc0f52d210d73ce4cd): complete subsection reference.

<a id="canonical-48f9df5e9c770a025112d9ddf470ec303ba103b07772723a7bca48c60d4b5a5e"></a>

<a id="canonical-348e10096794fd88f1e88059861e7392e94575857e4b5f418fbdec77f180e0b6"></a>

## api_group property — api_rate_limit.server_url_rules / ad4f9b858ad7 / 4

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-f7d1afc5db53e7eb98147cb73a45f3c6f546b4e42eae2bf9039def2f6f49dd8f"></a>

<a id="canonical-af8a917a928078163b1ad5568434e0e7852d5c164756b52ffbeae134b4e7fb5c"></a>

## base_path property — api_rate_limit.server_url_rules / ad4f9b858ad7 / 5

Type: `"string"`. Optional.

Base Path. Prefix of the request path.

Upstream description:

Prefix of the request path.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-ec150c21c9eae1de98717a12c0dbcc339e9a023eb144db259b671feb65c6d026): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-432fda8fa91b8eff33c364390100f048dfa030569dd4e96c431572d411e108d6): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d): complete subsection reference.

<a id="canonical-dbff17c8a083d7f475b710f89d693254ff92a2ba5a4e0857730d84516363cb13"></a>

<a id="canonical-475abf43e385b177aa991a60c2db055cd67f1b1780e97c7d1abc59d7f56ae876"></a>

## specific_domain property — api_rate_limit.server_url_rules / ad4f9b858ad7 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

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

<a id="canonical-a1c5d28bbe27bab63516b867311e4f3d7314a79e7b099b652f39439770d9b2a2"></a>

## Next pages — api_rate_limit.server_url_rules / ad4f9b858ad7 / 7

- [api_rate_limit.server_url_rules.any_domain](resources--http_loadbalancer--reference--group-008.md#canonical-262de909ed1f42fcfe1e2669020d10ded96b2560e75149fc0f52d210d73ce4cd)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-ec150c21c9eae1de98717a12c0dbcc339e9a023eb144db259b671feb65c6d026)
- [api_rate_limit.server_url_rules.ref_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-432fda8fa91b8eff33c364390100f048dfa030569dd4e96c431572d411e108d6)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-262de909ed1f42fcfe1e2669020d10ded96b2560e75149fc0f52d210d73ce4cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6b0bad4cd3249de455c4ef485de7470a384fa07b0b40c161810bab5a208d944"></a>

## api_rate_limit.server_url_rules.any_domain — api_rate_limit.server_url_rules.any_domain / 965af6aaf098 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-e3e7818305a0675c387e0b3082da745dad532250275f2418753facae98c0fad8"></a>

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

<a id="canonical-dc6b0e6036cca81ddd5ecb42aa5e2d32e960d82834bf9339352db39eef96683e"></a>

## Direct properties — api_rate_limit.server_url_rules.any_domain / 965af6aaf098 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f5d56f548d5ea3602e943f4830500d46f2214fe015f53bce9da8b6d219f52da"></a>

## Next pages — api_rate_limit.server_url_rules.any_domain / 965af6aaf098 / 4

- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36fa3ed33e95eff7fe4ad6e7e1c3eacde48b51a4137921297dfc0c778d4236a3"></a>

## api_rate_limit.server_url_rules.client_matcher — api_rate_limit.server_url_rules.client_matcher / e50f6373462d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-675a653beac2774d29a9a64f169534fb38637830fcd5dffd654900c4c29d9f4b"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-95a4e6271e3749bb603f78eaf79a5f118f9ac1313b1e249d8e59780288b86776"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher / e50f6373462d / 3

- [any_client](resources--http_loadbalancer--reference--group-008.md#canonical-e1dd8ab6ee0e9c9435708d7631688d9ad1ba51f08d91d903e0e4731f05112a92): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-008.md#canonical-0a1624c93f6f4a499e13f5bbf6f3c4bcc7951a3d6ab44aa2a063810524653a72): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-008.md#canonical-5f3cd45df2b1b6ca5aec529138f5ba903111cb22c7c8fc5bfc9cfa7123ca7f59): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-8009d72337eebd8b9ea32723fae9384d6b8cb67489135e751c72e3c8a0961aa7): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-008.md#canonical-e9baaf4787fd1bb365667e31105dedcfbd8e37c38d262a217905323edb98b9a6): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-61d962e8d449f570e0905c7652a7bb46e1f20d88295ea9fb26df680451d1842e): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-008.md#canonical-a365dc6fa3850e80a89b162e8b32950681960679755b2de87fcf2acecf3b3fa2): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-008.md#canonical-6df25eb85ae0c5457cc761a2b46b965426ed5126f78a0b0d2aac9fcd9cb56a0a): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-954abda8d323cd13d0dc77543add92aa0cda10dc0b78d37d616b49e8e4182796): complete subsection reference.

<a id="canonical-df1b4e5ef2dd1018290917fbb30025f1e031d5617989fff8623f9a0417c359d9"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher / e50f6373462d / 4

- [api_rate_limit.server_url_rules.client_matcher.any_client](resources--http_loadbalancer--reference--group-008.md#canonical-e1dd8ab6ee0e9c9435708d7631688d9ad1ba51f08d91d903e0e4731f05112a92)
- [api_rate_limit.server_url_rules.client_matcher.any_ip](resources--http_loadbalancer--reference--group-008.md#canonical-0a1624c93f6f4a499e13f5bbf6f3c4bcc7951a3d6ab44aa2a063810524653a72)
- [api_rate_limit.server_url_rules.client_matcher.asn_list](resources--http_loadbalancer--reference--group-008.md#canonical-5f3cd45df2b1b6ca5aec529138f5ba903111cb22c7c8fc5bfc9cfa7123ca7f59)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-8009d72337eebd8b9ea32723fae9384d6b8cb67489135e751c72e3c8a0961aa7)
- [api_rate_limit.server_url_rules.client_matcher.client_selector](resources--http_loadbalancer--reference--group-008.md#canonical-e9baaf4787fd1bb365667e31105dedcfbd8e37c38d262a217905323edb98b9a6)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-61d962e8d449f570e0905c7652a7bb46e1f20d88295ea9fb26df680451d1842e)
- [api_rate_limit.server_url_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--reference--group-008.md#canonical-a365dc6fa3850e80a89b162e8b32950681960679755b2de87fcf2acecf3b3fa2)
- [api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--reference--group-008.md#canonical-6df25eb85ae0c5457cc761a2b46b965426ed5126f78a0b0d2aac9fcd9cb56a0a)
- [api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-954abda8d323cd13d0dc77543add92aa0cda10dc0b78d37d616b49e8e4182796)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e1dd8ab6ee0e9c9435708d7631688d9ad1ba51f08d91d903e0e4731f05112a92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae7236aea9542b6b1c479d8fe4b8a53cc60403ad3516f00922aed9d00432998c"></a>

## api_rate_limit.server_url_rules.client_matcher.any_client — api_rate_limit.server_url_rules.client_matcher.any_client / b2355a48fc7f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-17fac3879e2125bf00f0f0da61b05dbbec1dce866a897b5a75f1c6fa73caf9c5"></a>

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
any_client = {}
```

<a id="canonical-9521585d412512e96c608c8295579a16dafca7420d96ed3ce66ae4832bc15e66"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.any_client / b2355a48fc7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea03fe3903724b34f13f0e46bc97bd09e1de2cef9074cfb954ab86396507b3e5"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.any_client / b2355a48fc7f / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0a1624c93f6f4a499e13f5bbf6f3c4bcc7951a3d6ab44aa2a063810524653a72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61405278b7f3d1506c691af5802114fa5aa6373ad4b3d673d3c4d9176ca6bdef"></a>

## api_rate_limit.server_url_rules.client_matcher.any_ip — api_rate_limit.server_url_rules.client_matcher.any_ip / 56c68a49ff30 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-cfa1063749376167debeb68868b7ea4a41c5d0d56354841b9db237ed01de7a1c"></a>

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

<a id="canonical-cf65e7af6b2e435209a6a3769c662f233869e62997d1dfe53998eb3df50fbc39"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.any_ip / 56c68a49ff30 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac26c265e853b585199e9427b53929d3b39c40d2a220735a31021e1ebd7908d2"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.any_ip / 56c68a49ff30 / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5f3cd45df2b1b6ca5aec529138f5ba903111cb22c7c8fc5bfc9cfa7123ca7f59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eee6d3323017aad2bc3dc02911ffc6dd1fcc8aeba4de3308fde6cdac3cfeea2c"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_list — api_rate_limit.server_url_rules.client_matcher.asn_list / 11f70f2d5adf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-3881c0a9ee7e34824623205f7ff43bac14413f8abce018c4c82e895bf6074a68"></a>

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

<a id="canonical-79c4393a4d47aac74a6285c446c6ef0a81983758a8f3c6ce3260cd2cc3a4214d"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_list / 11f70f2d5adf / 3

<a id="canonical-9aed5c04766a7f04f0c55e0440402dfc61c97c9f20dbfcf748e3ca34a0a15a88"></a>

<a id="canonical-93c4d3b2ff6bd359b1a4e7d94b7b5db388796e0a919e443bd8bb204e9c61f7cf"></a>

## as_numbers property — api_rate_limit.server_url_rules.client_matcher.asn_list / 11f70f2d5adf / 4

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

<a id="canonical-f1f5119edae7747a58bc58c46c3e2cc1ed1a371a643b46bff19fb64b643d4d3e"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_list / 11f70f2d5adf / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8009d72337eebd8b9ea32723fae9384d6b8cb67489135e751c72e3c8a0961aa7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4017bf3f59917a4daaaeba514e9bfc0b7ffce8803a7289bcca2d4773b2abe1b9"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 638c6a62c47f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-1fc2bab4c1f154c3a2de37f97e011b26d14ae24de864984e563efbdaa278d75d"></a>

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

<a id="canonical-ddb4098ed8df995db71404c134b388b07b69db3af65450314a0354cdea76faca"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 638c6a62c47f / 3

- [asn_sets](resources--http_loadbalancer--reference--group-008.md#canonical-eae927132153fcad18f7b896e33e0eec833adc50d3ca4b702021c8182b970f5d): complete subsection reference.

<a id="canonical-0a21bc679cb024ad4fa8a3cce3f7a2b880e30742792faa6b1a0dabaacf28c80c"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_matcher / 638c6a62c47f / 4

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-008.md#canonical-eae927132153fcad18f7b896e33e0eec833adc50d3ca4b702021c8182b970f5d)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eae927132153fcad18f7b896e33e0eec833adc50d3ca4b702021c8182b970f5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5e2dfce3a046720d767d9df0a9f485e9ce0f18cd3807cddf2ccd77a88579be9"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-8009d72337eebd8b9ea32723fae9384d6b8cb67489135e751c72e3c8a0961aa7)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-6cf4ada43a3f7d3f34297297ea2dd92f0aefef2073e2ff04b0a79c66c922e1f2"></a>

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

<a id="canonical-797fcba93d75d266cf660fe9b2e8875d88363b24ea849f6f1b568e4ff4bfea03"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 3

<a id="canonical-769a8964bd566a9e218435657ad8076e3eb1d1f1bd3ed2c68b66d1a14db817fe"></a>

<a id="canonical-86e302d149450c1a9fa178c5234351618120710879088591cf2c390b5b204271"></a>

## kind property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 4

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

<a id="canonical-5d6ebe932cf79a9f919ed6d8ef046e06293888d185ec5c03ceed98f99df1409a"></a>

<a id="canonical-fd2fa94b3beada81316d73ff46bc2a7a21682c34193dff65081fd344f54c5061"></a>

## name property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 5

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

<a id="canonical-b4b970a298c7e7f008cc2b449cb082407519251428c3fcbdf9649e118ff48e93"></a>

<a id="canonical-f08829b09b402c928534e3588acca1189dd67474684f075fb153a577cbe26b33"></a>

## namespace property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 6

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

<a id="canonical-bfcaa0092d394097521cb4000b2693743b00f6bf187881dcd3cad187855c35b0"></a>

<a id="canonical-5e0699f53e73eccd1fa46469274c247c9367470d0de643aa9eed09d0ebc1d202"></a>

## tenant property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 7

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

<a id="canonical-f75eb717d9234fc93a53258c3b8dbaaf79897a84aa86db9b4b955c71e88a96c7"></a>

<a id="canonical-07dd9d7c62ebc73879aeaada59ad1c5ab65fb205b2acba83e430d168d6256793"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 8

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

<a id="canonical-d462a7afd8d498b7b26f807d1916d300a7cb77259ddf3e34a68c3a174e40e911"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets / 652d35f9eafe / 9

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-8009d72337eebd8b9ea32723fae9384d6b8cb67489135e751c72e3c8a0961aa7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e9baaf4787fd1bb365667e31105dedcfbd8e37c38d262a217905323edb98b9a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba6ed2e47917e9f002178dc6025170e00061e686509b75e1d939d5e89bca0098"></a>

## api_rate_limit.server_url_rules.client_matcher.client_selector — api_rate_limit.server_url_rules.client_matcher.client_selector / 5ac32079c878 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-eeb12833db5aca8a5c1fd749e44954ee09356c2e10cf3bf4647860ef858bce1e"></a>

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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-9182f72b1e31b998c517f74ba8cecdf23305f497238985bf66df4d4a6372e716"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.client_selector / 5ac32079c878 / 3

<a id="canonical-4c7222fa1fa128aa28a6b2f9bfadda2d3e6d0c5381ea139b612d5e418e8676f3"></a>

<a id="canonical-dcf6cc23b0dc501c69f768cf070fd6da056dbb2c700ac4b1edcd4647e9b93932"></a>

## expressions property — api_rate_limit.server_url_rules.client_matcher.client_selector / 5ac32079c878 / 4

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

<a id="canonical-aa26176122f7586d9e12eb74de6146cc2b89945191b46502643629681db4a550"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.client_selector / 5ac32079c878 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-61d962e8d449f570e0905c7652a7bb46e1f20d88295ea9fb26df680451d1842e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cee9977af24d541bd831d4221382e42edddca57afc4e36c0a91e2314263e3ecf"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher — api_rate_limit.server_url_rules.client_matcher.ip_matcher / d403dbcbfa92 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-d5986f51a8816f1f86cde1756a7a089147aeaa0b0dd8ce4311e1e38a15026f26"></a>

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

<a id="canonical-0aa34f72a5c13bd41a1c5ca041f61b3f27e7424e8096ea047d6983652659cead"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher / d403dbcbfa92 / 3

<a id="canonical-29ab81ead4a1b8a3662f4326dff6d455d11e1a8158d20c5b9500ad5beaa9e536"></a>

<a id="canonical-d9bcac8b437d529c71cea2f6851faed665c66ecd00337b44b89237d2a7d9d5ef"></a>

## invert_matcher property — api_rate_limit.server_url_rules.client_matcher.ip_matcher / d403dbcbfa92 / 4

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

- [prefix_sets](resources--http_loadbalancer--reference--group-008.md#canonical-55a296b15c3e57de93a4ceec41e0c580ceec90bb036ada7cc0cac24cce5ae84c): complete subsection reference.

<a id="canonical-ac35223fbd92bb48f61949511e862fce463d6cae1a388785036b162fe29f657f"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher / d403dbcbfa92 / 5

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-008.md#canonical-55a296b15c3e57de93a4ceec41e0c580ceec90bb036ada7cc0cac24cce5ae84c)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-55a296b15c3e57de93a4ceec41e0c580ceec90bb036ada7cc0cac24cce5ae84c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19011f631d6f63a392857dcb0a53fd443bbc434e8d713d9b74c050ad8b522986"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-61d962e8d449f570e0905c7652a7bb46e1f20d88295ea9fb26df680451d1842e)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-c56fcaf332ca2e114b88ddd036797d7b3cfcc657bf7240ac5171ac686ebcc88a"></a>

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

<a id="canonical-11a49093011d6c57047d8991b257d20686d77b69aac4223d087b39e752a3e321"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 3

<a id="canonical-06677f6e7011b668ad619adbb202763b115a2a6ebe1f43389c51b4af48eaf668"></a>

<a id="canonical-90a10c6985879304cbb87a8bc51a3d0e67fa74d1adde823e73bea49324357cc3"></a>

## kind property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 4

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

<a id="canonical-fd841507c28060b7a710c96e7ad39ce750d6f419d3724910593f009b926d4d46"></a>

<a id="canonical-f3eea9f14f6987f3382c254e3e46be0992205f18bd4cb81fdd8b72590a2d4603"></a>

## name property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 5

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

<a id="canonical-fbde7d8f42cded8c21b41563b51bf4144caf15fd3fb515df5e92364c4339f02d"></a>

<a id="canonical-756aecfb338013c8d1c22505d24a0b686e76f80fe64fde100e90c4e8c6e75e6d"></a>

## namespace property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 6

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

<a id="canonical-4e07a5daf354a74fe9582fa75b9de977ea2bb20e9ce4465a1c12186c0e07c5d6"></a>

<a id="canonical-b9ddd92f53a478a8ee81faa1f53ad07dd210b72ddfdb774e8ef00d535b19233e"></a>

## tenant property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 7

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

<a id="canonical-236f8c58134d3187d67789c9e915f91dffeffaba3fc51e7eaff6dea65cff5416"></a>

<a id="canonical-625669b7569a13ffd53d7713368cdf5fec03733e6a709360dc85f9be1b2cf9c6"></a>

## uid property — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 8

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

<a id="canonical-a9a8e0574e6669410a65dfd255c445603096048da6e1903a962008a69b98efbc"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets / 59ee461ea29d / 9

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-61d962e8d449f570e0905c7652a7bb46e1f20d88295ea9fb26df680451d1842e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a365dc6fa3850e80a89b162e8b32950681960679755b2de87fcf2acecf3b3fa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe18dfa6ccbf6a797b340a02f7e39a33dc9b345e82d632234227544e32ad5287"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_prefix_list — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / b7b012e80073 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-cb6857adb30052c7b5e68ecad36a799189c4d0e71afe7c75860d90625b9036e8"></a>

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

<a id="canonical-722e5cfc0ec2ce6b59dae2e924297056bd0abbf24530f1a4abed7b5c5503813d"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / b7b012e80073 / 3

<a id="canonical-532c18e02d6c03413b72ca8dc9ec686d69163ae9442288eaf970a77729fb66c3"></a>

<a id="canonical-9ae4ab0954720c382bfcc24a3a61ebcee266d4220a45834828092070f687dee0"></a>

## invert_match property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / b7b012e80073 / 4

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

<a id="canonical-fdf9290c60439125a4fa2a11b5166f87fc91672ee88ca512cfb649f6f334e681"></a>

<a id="canonical-277c1367a074b532ad26880aa228341de39cb317ea6e2f57bb27a60b2b2716c8"></a>

## ip_prefixes property — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / b7b012e80073 / 5

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

<a id="canonical-367fba7d6de3876649e1a6b55e4e12623b8ea73a2b736cba464eac322833679c"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_prefix_list / b7b012e80073 / 6

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6df25eb85ae0c5457cc761a2b46b965426ed5126f78a0b0d2aac9fcd9cb56a0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-340cdbb7edcd6ec0f0a1aac14ad1ddc63e0083383af43764973f4a38bd324185"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 0664e2a08e40 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-705f601e11e93f315202d3034a065325fe5ce43375f16d07845c1270d328223c"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-c8c4ee5b653e417a4cf1b70e7c34d71871454f37ed41006b3c094e2edd685797"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 0664e2a08e40 / 3

<a id="canonical-0af05364a6f069ac3a03c7717d8527ebaf760d0801ed3fc4408ea42670f0b145"></a>

<a id="canonical-2248dbc782219cc17ba6e7f3f354f6490688bc1bcdda23f95966825f2f416ca2"></a>

## ip_threat_categories property — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 0664e2a08e40 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-a12f917bcbe4a59edab3909bcd3d413ae8356585d8cc4b1d76bdf23f971d4de6"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list / 0664e2a08e40 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-954abda8d323cd13d0dc77543add92aa0cda10dc0b78d37d616b49e8e4182796"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b69f3686cd02484f5243215a8a7fc15dfffe83f2db69df169f26769f32ca39b5"></a>

## api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 9c12e04b4832 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-725e038b34018965e208c5e7fba8ea8d86074252a428c1c4f57e190f9f419824"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-c91c1899091e7c5cf2c49460dbb28e4a1e4ddbbb0d15e82685f73148a8a2e3df"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 9c12e04b4832 / 3

<a id="canonical-d01aedeea57e523f6022a8a44fc08c270eb398513411d2a9346bb4e7542debc2"></a>

<a id="canonical-dc7201cd66ccf25623dd815c28d90fcc0562334e888b809a966236f749063df8"></a>

## classes property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 9c12e04b4832 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2addd00aa9616cecec64a1fea6a0d46c6956496968225951a6a8f609ec08d171"></a>

<a id="canonical-95cd6bc6a3a58b38233d6179c1dcd3a9af06e3833372dcf2534acd5c3b80d269"></a>

## exact_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 9c12e04b4832 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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

<a id="canonical-1d59195bdcc8fa18ba0b60fadf8e17fa7d49b988e1271a171d82562a34a0a834"></a>

<a id="canonical-8d5f7b1b7f5dda33ca4f9f84e74f13bf1fa01b480c4e7b54a57a42176e5f3d8e"></a>

## excluded_values property — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 9c12e04b4832 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-e523beb75d9848a9efcd8745408c2b5461f35632492483591f813ce791afd0be"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher / 9c12e04b4832 / 7

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-71dbcfe1a323693e4dc7cdb38cddf03c2f2a53080b512a4f7707fc27e45f29f5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ec150c21c9eae1de98717a12c0dbcc339e9a023eb144db259b671feb65c6d026"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-813ffb660af164667251d99e127a09180e96ede923543a00918b7bfbbf844a70"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter — api_rate_limit.server_url_rules.inline_rate_limiter / 79816ac977c9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-77026d777dd8d6f5c54d4e70c1d63531091947f9ce1d22aae64b80e25df0da5f"></a>

Type: `"object"`. single nested block, Optional.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
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
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-711543103be2c8555a93569f91b07a411f1abc25051f192efe2207b4abe9864d"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter / 79816ac977c9 / 3

- [ref_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-da51225d2415e2b18b540fee8d7f96e22576a7acc7b4f1bc3d3107e83cbfc063): complete subsection reference.

<a id="canonical-9a451498c4f9b9603e60604a9bc26ddf87bc9aca0cc1106b8b8774c30f800972"></a>

<a id="canonical-e69af8bcc8d99db0fc9563c319be80ecf4cc2e34d3e1d3243f3484b2014e2593"></a>

## threshold property — api_rate_limit.server_url_rules.inline_rate_limiter / 79816ac977c9 / 4

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

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

<a id="canonical-8c95acf8917a786e07d41ae4870b88ad70d546b7580e4d1703b81bb0d98ba8d3"></a>

<a id="canonical-2e9d5757c07313d7dce4eefa34e9780e1534e363b012eb195b29ee27d2739174"></a>

## unit property — api_rate_limit.server_url_rules.inline_rate_limiter / 79816ac977c9 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

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

- [use_http_lb_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-155a759d1e12ca16b72f8a06b945d615f7a05c2644ae67e3c6a9ae75bd13707a): complete subsection reference.

<a id="canonical-3530d4681bf4debf4bedfb474fec215b608dabec8350f6ccd063af5ffc6317bd"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter / 79816ac977c9 / 6

- [api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-da51225d2415e2b18b540fee8d7f96e22576a7acc7b4f1bc3d3107e83cbfc063)
- [api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-155a759d1e12ca16b72f8a06b945d615f7a05c2644ae67e3c6a9ae75bd13707a)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-da51225d2415e2b18b540fee8d7f96e22576a7acc7b4f1bc3d3107e83cbfc063"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5fea75932b39812a1bd9d4caf258308ba183bbf824ddc828acd10832780e725"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 2f5b5599118e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-ec150c21c9eae1de98717a12c0dbcc339e9a023eb144db259b671feb65c6d026)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-36e59fa1d7c22072eae8c758673d7538c164a6a54effcca2ce73fa3b032c163e"></a>

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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-ba5bda7f2e1ab4b4de34c695e89d6e7614368f7c4bfc4cef44b78ccc814b63f4"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 2f5b5599118e / 3

<a id="canonical-c908d28c8dcefe164738162cdda6ec418ed2de860d73e63ea42316356048e3fe"></a>

<a id="canonical-99986442736a79c29dfeba71f3bca5baafaf031e2e63a112aaac62f874c9f104"></a>

## name property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 2f5b5599118e / 4

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

<a id="canonical-624cc3ecc7436198985bc2542c00b27d58c3fd600a104f42cd1f6d01a315839c"></a>

<a id="canonical-a115a7dc72f3f20769f22c76d4330add9018be3b70e6e547455d245fd6b1bcb8"></a>

## namespace property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 2f5b5599118e / 5

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

<a id="canonical-bcf5e2889ad25c064c032a3e8513144825d996c705433eb39fe2216ece9c9f98"></a>

<a id="canonical-1290faef7495e549425effa777e0d612a7c9c6add8931622dd874f7f2e8dfc31"></a>

## tenant property — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 2f5b5599118e / 6

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

<a id="canonical-fa3e3a917c6fd5c4c1dc1f714d565a6596d90cb88f3ef98a87b93f62f7c99cef"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id / 2f5b5599118e / 7

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-ec150c21c9eae1de98717a12c0dbcc339e9a023eb144db259b671feb65c6d026)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-155a759d1e12ca16b72f8a06b945d615f7a05c2644ae67e3c6a9ae75bd13707a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e64bef745770485364b486cdcc7b9468639d26dedd75db9cfd3a91beae6281a"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / 0142a01605ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-ec150c21c9eae1de98717a12c0dbcc339e9a023eb144db259b671feb65c6d026)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-b90eebf30e0e16ee6b171d54aeb71da85fb26c8858db7da056d7f8b357266ac1"></a>

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
use_http_lb_user_id = {}
```

<a id="canonical-f9ccaa888a3979602946b9eac1665104d572bc780fca4f9fb73f0568dde993f7"></a>

## Direct properties — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / 0142a01605ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb0572940129a09bbd1984b9258293af32d385f2e49b5e6d2a46e65a0ae8b803"></a>

## Next pages — api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id / 0142a01605ef / 4

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-ec150c21c9eae1de98717a12c0dbcc339e9a023eb144db259b671feb65c6d026)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-432fda8fa91b8eff33c364390100f048dfa030569dd4e96c431572d411e108d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-701c5ada30e347e9154af3bad788f761f631ff2ec0a33ad7f39e2aa2fff5e1c6"></a>

## api_rate_limit.server_url_rules.ref_rate_limiter — api_rate_limit.server_url_rules.ref_rate_limiter / c038ab53a3e6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-7c363adbd9648409571be27319f6765ab8950a001a711273f02c8a6b010b3be3"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-a693c83b6d0799e1246e0344df87d118d45b9de4f00d3aa963a0398e3fb4c93e"></a>

## Direct properties — api_rate_limit.server_url_rules.ref_rate_limiter / c038ab53a3e6 / 3

<a id="canonical-741a1d4b21f9cfa5a517343b89989d4b5fe98021c72c64105b59d0f49cadb17d"></a>

<a id="canonical-a9c6c15444742db885d61bb97a6399716b401f5af2d2bb23c7402b222d654ac7"></a>

## name property — api_rate_limit.server_url_rules.ref_rate_limiter / c038ab53a3e6 / 4

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

<a id="canonical-304c099581e11827391e4557b6cf5c065698b416542f98cdd75641f5099a6f41"></a>

<a id="canonical-4314301ac15cae745978ef2f77a1223f0c497f215011a7e1223fa84804d64f15"></a>

## namespace property — api_rate_limit.server_url_rules.ref_rate_limiter / c038ab53a3e6 / 5

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

<a id="canonical-65718982b28dcc3b130787fe9937578361d90eb74fd8196bba009f40e91c16a7"></a>

<a id="canonical-f6dee7f0b20a5e104fdcf9e27c84334091c1a81656b9960d78912447760e2e79"></a>

## tenant property — api_rate_limit.server_url_rules.ref_rate_limiter / c038ab53a3e6 / 6

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

<a id="canonical-869e5f51d3b16193859a16aac1d3a26e45c0c0755df446c805ea4fd89bdfe463"></a>

## Next pages — api_rate_limit.server_url_rules.ref_rate_limiter / c038ab53a3e6 / 7

- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76149f420a8cca99c47a44fde8e99d4cd57a0b4d127b11fbe5d92871815981cd"></a>

## api_rate_limit.server_url_rules.request_matcher — api_rate_limit.server_url_rules.request_matcher / ddc6965ccd6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-5773a46dbd6359414c16ff546d77a9f18727ac7667bac8b41ae439841cf29af1"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-f72bf8aa564930c6b11f01876e2cb35f1f2266f2574f96233f4fcc51f5b8ce0a"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher / ddc6965ccd6c / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85): complete subsection reference.

<a id="canonical-12829acf032b912870e00f4199e1bcea01b5a065bd8f9164c2b90c8da46fd84e"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher / ddc6965ccd6c / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-767445f31ee1f8e40234c4ebb5f987a31ae1a3189ab51b6c7cac55b3b0e726e4"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 05fed821904c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-dfda0e4336fac4096b460dccb013bdfccc9898c9093db5b6d6ab3140cc916df8"></a>

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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-bfbaff591a2fd01af892a4c31229f6d9f7c193ee33f03a78eaeb516af2be605a"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 05fed821904c / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-a264af58893cdd43a1080864d79f9937bc625f68727a8c01ec24b18f68977f84): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-531a3e9efef1eda9c322eeccc4a8e65d86c48036ac22633d8f14ca9afcb4712c): complete subsection reference.

<a id="canonical-73c4b20afc48fb1ce33d3b028712bfe29f8a202481b579fa41c0226420c002c8"></a>

<a id="canonical-4d259caa8b14124b0db26797b823b8bf9e041a8df98f17f849686dd8f4b8414c"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 05fed821904c / 4

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-2f7bcab03a8e70db54faca3581a87c571744d01c88aebb6fd9f4e1d6fffb5783): complete subsection reference.

<a id="canonical-4a4c0e8b6db6bc0a8a2e408d742f82ff61ba39401d3465cc346538bd49e17732"></a>

<a id="canonical-7cd2434a04a3c55ed8a97df744344423c04bb7503ebeb6cb67be70c95cb047a8"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 05fed821904c / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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

<a id="canonical-3e68c6dfbd7f6834463ecec033e6311b17a190a499d2bff66ced9e3b3cad0c3a"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers / 05fed821904c / 6

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-a264af58893cdd43a1080864d79f9937bc625f68727a8c01ec24b18f68977f84)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-531a3e9efef1eda9c322eeccc4a8e65d86c48036ac22633d8f14ca9afcb4712c)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-008.md#canonical-2f7bcab03a8e70db54faca3581a87c571744d01c88aebb6fd9f4e1d6fffb5783)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a264af58893cdd43a1080864d79f9937bc625f68727a8c01ec24b18f68977f84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b68130f9aa3ad9b463ca6c2cadcb24d26271d49356f0bcfa36cb2964cce9aae"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / 1ffeb2a59425 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-ee7681aece02054ca5b30326dbca0de2dd0e08b1f3e3022ed66c38a939f13408"></a>

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

<a id="canonical-d5632da1957f006cf86fab29733ce509d751371483670f210a095e060e17585d"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / 1ffeb2a59425 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e784d43a1b5cc2441e11f648447d6dffa2652f1025e7769889c9b1e711d28d7e"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_presen / 1ffeb2a59425 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-531a3e9efef1eda9c322eeccc4a8e65d86c48036ac22633d8f14ca9afcb4712c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94c1aaeb39b0ad04581038eb2ed33b44714eea33ff906157685e656d275a83dd"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / c0a7850c2aba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-cc34e8c62294ca6e05401f3d1d37e5e16ffbe557d45d322d3db5114543f0b985"></a>

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

<a id="canonical-a78cfab77262e239c323ddfdc81d4506520a9b10daeaa7fcedc142c173e01c82"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / c0a7850c2aba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6cbfe0b0bdf803fbb973de019c201a2322c91a801ff65da62a48ef8e2a6c479"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present / c0a7850c2aba / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2f7bcab03a8e70db54faca3581a87c571744d01c88aebb6fd9f4e1d6fffb5783"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e026dabe6b9b8be773dc7cf9bf62ef29072f8139c1cb7876f0968536f3c3dc0"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / c21616cc6b27 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-f5a792fbd2d3347f7b23db5518bc2790dc80d9a8048bf71d63616738ffac6984"></a>

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

<a id="canonical-4f5d33d1aa645dd72c768b620c38118deb69e993f37b14601dc12cf8465c4075"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / c21616cc6b27 / 3

<a id="canonical-d15e7fc26409d9b51c765def69522b55a07955fe10e9f661ae15c05f5e0b1378"></a>

<a id="canonical-f6ce2974207bc95cb9980b89df31b833bc76728153b6a47cc3e7c70853f61cf6"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / c21616cc6b27 / 4

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

<a id="canonical-0e43bb1c0c3484e22dff8a3cda313fe48b11a3813e76ce41ce394529d5260ba8"></a>

<a id="canonical-7ca17c3ef2d50124e8be41376c7c34df844609dd3f3c25aa98785dbc3c96060f"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / c21616cc6b27 / 5

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

<a id="canonical-72a97dddec1b7583116d0085b4ead3bf07c797f4e2079bd5b8797c8ef247397f"></a>

<a id="canonical-2775c6776dda57abafaee9133b4c43838ea345ea58e1b0cef81320f8214f080c"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / c21616cc6b27 / 6

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

<a id="canonical-3cc7bee3f1bb5e8ab2b0582edfd08585b43f97e31b3c083c6cf88be71d671485"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item / c21616cc6b27 / 7

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-cce443b6d7220939450ce0f3b2f660287214f29394b6b0c28c4672df20d25275)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2541429edaf7586ca8f55496376e0382eb9be7969b0150df8c63119f4bf8fad"></a>

## api_rate_limit.server_url_rules.request_matcher.headers — api_rate_limit.server_url_rules.request_matcher.headers / 4e01e3e6615a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-bdf9de01e9bb9f98322d9019910724a437b9746dcc574b40d18d1398d1c072e4"></a>

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

<a id="canonical-935b2c9eb4550b398e5b252182090d7c6d673e2eec6143bbad994a31b432748e"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers / 4e01e3e6615a / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-4338e3a3f62753409f27b20baa08f07e8cf5a9e68caac26dfe176533ecf5e0d3): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-42a3bbd9d7326a341ed38d9a49edb3966a89c871c119d2972af7fb3033c14933): complete subsection reference.

<a id="canonical-19688e4ea08104424274ff0f2b2707bb25abdf7a7f029fd5296a2c07ec42a695"></a>

<a id="canonical-2603cd421614b71d0dad704e75eb65a58a448849347ecdde37c8048966172e42"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.headers / 4e01e3e6615a / 4

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-d136e3c4cf8c812d77c5d7ce31db5da5355f9e79e7fff71cb5c386533e146704): complete subsection reference.

<a id="canonical-0b2dd58858585bbfb97aba718bb41e74ed9ebde01361ff519ff7f3e7b097b5c2"></a>

<a id="canonical-2ad69d79ac599b4a2c4ecfdf22a1b97926693619608e73aa65e43c32b83e9cfc"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.headers / 4e01e3e6615a / 5

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

<a id="canonical-b83aa7684a89c6ad26cc74898761df33217717cf2ed20c0a6abf74ea95f96eb1"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers / 4e01e3e6615a / 6

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-4338e3a3f62753409f27b20baa08f07e8cf5a9e68caac26dfe176533ecf5e0d3)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-42a3bbd9d7326a341ed38d9a49edb3966a89c871c119d2972af7fb3033c14933)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-008.md#canonical-d136e3c4cf8c812d77c5d7ce31db5da5355f9e79e7fff71cb5c386533e146704)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4338e3a3f62753409f27b20baa08f07e8cf5a9e68caac26dfe176533ecf5e0d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9ee8893ddd2ccb15b741139921e85e3c4c94337566e8f45adf49d5f9ad5c9ab"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_not_present — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 21a492e0548e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-1ff87b7d27b9dfc2d55db3503b8182c74ae16d9aff32fe022f60ad22856e1616"></a>

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

<a id="canonical-1216bac50eb78963dfe8a8aba0053d1fe1b8610fb5ad6d0864d7acdb4951f4e2"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 21a492e0548e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86194fffc1f7a30a3f65063e3130fb94452300d5efe62cd3220b1e5ff3ec79c1"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_not_present / 21a492e0548e / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-42a3bbd9d7326a341ed38d9a49edb3966a89c871c119d2972af7fb3033c14933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b18dd0e28dfb63827ef805bd20f877c71e38207322df6634e1410f441c8538ee"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_present — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 6ba4ef4a2072 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-8993384eabdc4101fb8b103fb95164959e1557dc459b920a770e140d0ffa9c10"></a>

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

<a id="canonical-4d98f55c72dc4ad838b64a8479808d08f7c2b434868528e70868bff56ff4262a"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 6ba4ef4a2072 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-418f574ed2493b3231ae7eba245374ee583640bc655b419724254f011217b4e2"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.check_present / 6ba4ef4a2072 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d136e3c4cf8c812d77c5d7ce31db5da5355f9e79e7fff71cb5c386533e146704"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a88e60b57d230692449ac84a2c5c4ba4137b325a81732755cc3ad3d5bc672c7f"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.item — api_rate_limit.server_url_rules.request_matcher.headers.item / de7898811c86 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-ac9a20c240f4cbb2972fb1e347112dfa6b63a1b42dbbb30aae48172a5bb20f41"></a>

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

<a id="canonical-24c0f05693a4dbacc8dabb841c155517f196034a0ab9d38d7aa53f27f6aef22f"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.headers.item / de7898811c86 / 3

<a id="canonical-6424715c8a234e8fcf4951c28e6be1172ef55fcc0e26610ac7209978ee4e1d51"></a>

<a id="canonical-8a2e212f19c6ace4a601058d55cc4fc3a37ef8cf0318f59c06653137f2b703b8"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / de7898811c86 / 4

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

<a id="canonical-559ea1a61ee2595129b1cd5a6fd1beb1a13709f15e2e93f4f091060f142815ce"></a>

<a id="canonical-d31b2c1d5f76d6d363de18259baee1478feee3a676328a96004574d8030a88ea"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.headers.item / de7898811c86 / 5

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

<a id="canonical-238ceffb236221a76e68ee10c7b55513a2f74fb15e6fbbb9e38de461f65b2c95"></a>

<a id="canonical-13aad0db623c7b9446508422ac29323f071933d1c8e30b2a9933074afa426f4c"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.headers.item / de7898811c86 / 6

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

<a id="canonical-d506d7fcd53fd785939ba6a42bda4cef7bb24304446e74aaab40334e06678d4a"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.headers.item / de7898811c86 / 7

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-9410343bea1ed764241248e9732452a1a124a8b40c34fa23d4f875bf120ace02)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a996cf819e1aea1bb89c991e2cd842305a83cc0d1aa545b8e35e37b4423b1781"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims — api_rate_limit.server_url_rules.request_matcher.jwt_claims / ae88be7c9775 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-9a89b6123cd0ad5c04e5875360d9b2d016c1fe46125e308a778e724342808416"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-e329a1743bd0e58b00266a4e50c08bf14c0cbab6b09f1e4961c8f896f720b896"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims / ae88be7c9775 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-d4983531069e1fbbad1c3e5aeab77956dbb7ac0dcfdfd50ac53e3a2b7ab39446): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-956c052e1572c58a7227343c12330bfc2c8898bc5cfebedf969e63667f5b04ef): complete subsection reference.

<a id="canonical-651488ead28ac6203e900097b7e348656f287214f55314226663155b215a7a10"></a>

<a id="canonical-dcf4d65c42f6acb33390e6414623458d95e3038835cd2c356f6f4d8fdfaf702a"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / ae88be7c9775 / 4

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-958ae9d48b92b1b9e1aa5ba61a2487d48f38620378bb99fa7c743bc400d5b736): complete subsection reference.

<a id="canonical-74760c44cbcc3c5f004405c97d22d6cdf7d553af786c287996fa2e9e54d87609"></a>

<a id="canonical-a6f6505f429bd69896e9c7e64950674e39024c789d5d5cf1fcf4669f5d1ca945"></a>

## name property — api_rate_limit.server_url_rules.request_matcher.jwt_claims / ae88be7c9775 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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

<a id="canonical-0929c5b6ebc25668da7cff49774bf3051a83d35b0a834904c4b773655c500a9d"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims / ae88be7c9775 / 6

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-d4983531069e1fbbad1c3e5aeab77956dbb7ac0dcfdfd50ac53e3a2b7ab39446)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-956c052e1572c58a7227343c12330bfc2c8898bc5cfebedf969e63667f5b04ef)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-008.md#canonical-958ae9d48b92b1b9e1aa5ba61a2487d48f38620378bb99fa7c743bc400d5b736)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d4983531069e1fbbad1c3e5aeab77956dbb7ac0dcfdfd50ac53e3a2b7ab39446"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63e7835be2d3f2b506e39fbe683b3c43fbc095f44b53f5638fb7f6fd9d5dff9c"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 85bcee0c0ea9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-72708ec7f5343d080288358ff52139076cd07bcd5f85822e956cb4c2323e01b1"></a>

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

<a id="canonical-6986cf34bc76db3f6dd2dd7a1f10e4e7d97db0e2db38637434d7f99bf6ff838b"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 85bcee0c0ea9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-455994b02a34efedf56316fa78c9d638b4d7046abb3c9f88c82148ab5bea3234"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present / 85bcee0c0ea9 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-956c052e1572c58a7227343c12330bfc2c8898bc5cfebedf969e63667f5b04ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c1fa3a41d70b4be3a552b000e94a075f4064b6279e746f48a84a93d0e7e9cc1"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / 0c9a5081b464 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-1cf8fc1e215756e828f63bf9e6611ab7fb4d75ef9d12e46f059ca5798acf4321"></a>

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

<a id="canonical-073aa30c4f1c22158d0231415686041063a867fa513bc016ed8009dc7f5ca8d7"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / 0c9a5081b464 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ccea694ab99df624642216b060731648a73b71425ae8d9fcb9426dceb1eeb63"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present / 0c9a5081b464 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-958ae9d48b92b1b9e1aa5ba61a2487d48f38620378bb99fa7c743bc400d5b736"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75e21657caf9df2d2dcc3f515b00605248eab944efb12a0d138f67ff0a29b858"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.item — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 2b24df3bd261 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-3270ec3fb70082bacd3db7fcc4c7cfc3eb5dba608d4078c278e1a9e6ac75be47"></a>

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

<a id="canonical-9cdb3e39eb90d5692ba6d3899e9edf9119d3e6d75bc1782e12efbd04802e6fa1"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 2b24df3bd261 / 3

<a id="canonical-799d9ca310db07369a646b249dbc38336c743a91c4704d1e6292b8dce8257758"></a>

<a id="canonical-b0aef11e74924553ac5e3a6e40d4c8d69b31d33a17ab1f837835bf664bfb714e"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 2b24df3bd261 / 4

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

<a id="canonical-06c0989587945a6a4073ece5dbf9444f712f0fd0cc9d0e4c04025d2cd473667b"></a>

<a id="canonical-800ec638c9a5885cf115fac358cb762e89013cab52bebec31ad447d47e7307a5"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 2b24df3bd261 / 5

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

<a id="canonical-030eebb04ce909ac8ea140f836de3b38fdbffc794b561848cabce13cbefc547b"></a>

<a id="canonical-f09cbde9da857c387cd8f942bdd699718c37870c8b5cb800f3d5e7e0fa77baaf"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 2b24df3bd261 / 6

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

<a id="canonical-37ad3051ec53078bbb55d2036bed68d85b1da36300ffd0241862d8e0609293cf"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.jwt_claims.item / 2b24df3bd261 / 7

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-3248bbdc956cb9e3618d8480e23b73afa336cb0aa13acf49c175eb217b21d785)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58201c7705c4d23a6002f18e63559197eaea3d9253bf2059b7fd2356f7eb381d"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params — api_rate_limit.server_url_rules.request_matcher.query_params / 87904c912efe / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-07910874e8dd54b97413e3070e9c6a675f364019b2a8a2deed5123d8636f3466"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-c270b4a9290dee9a5afac49c6a628a3d6bf3521fa4b2e40a34e3e321c1ad7966"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params / 87904c912efe / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-c96030029a7de80f954286b95c6b6b995bea7fc6aa9e7a9a3dbcb7938c233ce9): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-b85c3f3f4bc3f200ca481baa768feb6cd39f848d1aff883f5107b4ee0f8f2696): complete subsection reference.

<a id="canonical-3e49ed1a32370afecc6aa4913fa04d8e27800a362e27f8e9d05509aac01d9f50"></a>

<a id="canonical-db6cddc751a231dba5cfdc7d0188e3bdec66c027de0fd2e8b3ca45e95f1356bd"></a>

## invert_matcher property — api_rate_limit.server_url_rules.request_matcher.query_params / 87904c912efe / 4

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-71ad445be214227e250d0190d0f12d57ab7ffa9a3b93f476d38814ff42953a03): complete subsection reference.

<a id="canonical-3649737b23bb4b84097076ce7dc97e78c82d748f6bf476591015d5aeb46a509b"></a>

<a id="canonical-5362ad54790e00c3c027d9cf78df9e692505c809b168382a16326e5223c7544a"></a>

## key property — api_rate_limit.server_url_rules.request_matcher.query_params / 87904c912efe / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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

<a id="canonical-b33602ab54efe569325ba3bfb1b5242b537c33caa8ce07a379ead04a98c71728"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params / 87904c912efe / 6

- [api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-c96030029a7de80f954286b95c6b6b995bea7fc6aa9e7a9a3dbcb7938c233ce9)
- [api_rate_limit.server_url_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-b85c3f3f4bc3f200ca481baa768feb6cd39f848d1aff883f5107b4ee0f8f2696)
- [api_rate_limit.server_url_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-008.md#canonical-71ad445be214227e250d0190d0f12d57ab7ffa9a3b93f476d38814ff42953a03)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c96030029a7de80f954286b95c6b6b995bea7fc6aa9e7a9a3dbcb7938c233ce9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cff730f53b88fde871ec92717301fcf1cd5306e7c9f927d9e20be503bce68f5b"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 455255d6c5a7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-f36dbeb28746b018f0c8a970bb8cd5c4f19b9d110c8596e9621c839149344e16"></a>

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

<a id="canonical-d34ebdc149307ebdb889f09d39dedb8f06a02d87239cb0f6ff4e91ec91b6e16e"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 455255d6c5a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fed51089345ad2fef1875c8a35be0065c7b3222a30162a6ef527073d233c6734"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present / 455255d6c5a7 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b85c3f3f4bc3f200ca481baa768feb6cd39f848d1aff883f5107b4ee0f8f2696"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-685d2f7be2be35d17ccaeabdf762b688653c30c227bfb2b7e82d037f8dc40591"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_present — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / d36114f2bd6f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-a0e1b2c53092eade89208de7009b372f3a2b5b7e2acf3047d07821eaaa02ee4f"></a>

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

<a id="canonical-2cacb0466b0b46893852bb9ab5caaa4436f5ed3db4d757b5ba7b0ef0ff5191d4"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / d36114f2bd6f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-81ca07de6ecb041c68238a0411cf8c089c4cc4dea427734d3bf349502014f1b6"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.check_present / d36114f2bd6f / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-71ad445be214227e250d0190d0f12d57ab7ffa9a3b93f476d38814ff42953a03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-558862870420ec29c08ad826bfaa3ea5b3e77bd07a5f97fc4d9d898c89eeca18"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.item — api_rate_limit.server_url_rules.request_matcher.query_params.item / 939cc7481961 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-25567186ffe6acfdf1ecd23deee8c54bb5a06a5ee2305927e225ce77d233fe0d)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-49cf1bad1f60be2d9dbca17297c4006e0a6b5610cba022bf9fa16af203d731e3"></a>

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

<a id="canonical-87ff7363e24a147a93f454044717c388c0f5f56b494656ba8c279cacbdd8390e"></a>

## Direct properties — api_rate_limit.server_url_rules.request_matcher.query_params.item / 939cc7481961 / 3

<a id="canonical-ce1ba953bcd3677cd002d6a686451c6550938f48e7ffbe303658d73c59c54930"></a>

<a id="canonical-7030d47034aeedc614a7d6ca3f68dab236ccf8d64ea938de1459e6d75a4ebb3e"></a>

## exact_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / 939cc7481961 / 4

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

<a id="canonical-2767738f87fe1f5fe7a49bec3045e0ddfb9d477571484d1de357bc4adba64031"></a>

<a id="canonical-9f4e1174eec561083bb09852d6e9ae4e46f36fc73ca4573870a6a2b47cdae7fa"></a>

## regex_values property — api_rate_limit.server_url_rules.request_matcher.query_params.item / 939cc7481961 / 5

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

<a id="canonical-a0ab4dfa937bd3b3d800d40345c569c274871609433851efa3fd14b4706a2fae"></a>

<a id="canonical-d56d171a61f9e6aa1b66587b47ca65a3fc6ffb188fd921c7bffdceeb2c6bbcd1"></a>

## transformers property — api_rate_limit.server_url_rules.request_matcher.query_params.item / 939cc7481961 / 6

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

<a id="canonical-eabb839349ae14f9ead2dea2206cb7e884f1b2a258745bd3b0178e363d724bd6"></a>

## Next pages — api_rate_limit.server_url_rules.request_matcher.query_params.item / 939cc7481961 / 7

- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-dee7740dafdf076b2f20b766649cef3ef80513d47dd5f31d68ef2df0014c5c85)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ef837ffe8345df9687ec248a1b52c2b3b8ba8e509baf1807cabd79f65bccffb"></a>

## api_specification — api_specification / 389ec9405489 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- api_specification

<a id="canonical-c649651f501ebe8e3dcd532d5b9af4577f039254f64bb23dd8eb42a1771c843d"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_custom_list"),
  validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_disabled"),
  validators.ConflictingObjectAttributes("validation_custom_list",
    "validation_disabled")}
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
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-c649651f501ebe8e3dcd532d5b9af4577f039254f64bb23dd8eb42a1771c843d)
- [disable_api_definition](resources--http_loadbalancer--reference--group-017.md#canonical-199baa9cd32513f95b003364d3600c1e425b2483c83890657c8fe21f70f09bf7)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

<a id="canonical-8c62f058cab3b1170d1036b963ef957e6160d87eb2138cb29490cffea6955f16"></a>

## Direct properties — api_specification / 389ec9405489 / 3

- [api_definition](resources--http_loadbalancer--reference--group-008.md#canonical-03803bdbaf57c6f52870ce3d1ae6e0e025756465a6870b7051b4887cd8cb8393): complete subsection reference.

- [validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c): complete subsection reference.

- [validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4): complete subsection reference.

- [validation_disabled](resources--http_loadbalancer--reference--group-010.md#canonical-37e8d4458cfcaef85a98370e5f94e240f8bd8191af346c1e16886e800d749ece): complete subsection reference.

<a id="canonical-6be3a72d22549ee9ae2cde6d57966e9ff62296bd8dd1e1c447cc72c61cd28ad4"></a>

## Next pages — api_specification / 389ec9405489 / 4

- [api_specification.api_definition](resources--http_loadbalancer--reference--group-008.md#canonical-03803bdbaf57c6f52870ce3d1ae6e0e025756465a6870b7051b4887cd8cb8393)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_disabled](resources--http_loadbalancer--reference--group-010.md#canonical-37e8d4458cfcaef85a98370e5f94e240f8bd8191af346c1e16886e800d749ece)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-03803bdbaf57c6f52870ce3d1ae6e0e025756465a6870b7051b4887cd8cb8393"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d64aedb6807fe730d286e226a8d2b784b8f288486cf0b9b165150e70f48df5c7"></a>

## api_specification.api_definition — api_specification.api_definition / ab9a6f8bd191 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- api_specification.api_definition

<a id="canonical-c117d11a858def8adbda7e203a8f9c77993c92cdcd80f0bd8062c604fdb343ea"></a>

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
api_definition {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb8825e8e001ec2063f2df2ff1b0d25065ff789fff5b51af6492f3327ead1c42"></a>

## Direct properties — api_specification.api_definition / ab9a6f8bd191 / 3

<a id="canonical-4d1208e115c00e601db8c24d05e2e7784e939e9ba30b20abba21c74193207a20"></a>
